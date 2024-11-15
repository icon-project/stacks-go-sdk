package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn       *websocket.Conn
	idCounter  uint64
	pending    map[uint64]chan *jsonrpcResponse
	pendingMu  sync.Mutex
	handlers   map[string][]chan json.RawMessage
	handlersMu sync.Mutex
	ctx        context.Context
	cancel     context.CancelFunc
}

func NewClient(apiURL string) (*Client, error) {
	u, err := url.Parse(apiURL)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	switch u.Scheme {
	case "http":
		u.Scheme = "ws"
	case "https":
		u.Scheme = "wss"
	}

	if u.Path == "" || u.Path == "/" {
		u.Path = "/extended/v1/ws"
	}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %w", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	client := &Client{
		conn:     conn,
		pending:  make(map[uint64]chan *jsonrpcResponse),
		handlers: make(map[string][]chan json.RawMessage),
		ctx:      ctx,
		cancel:   cancel,
	}

	go client.listen()
	go client.startHeartbeat()

	return client, nil
}

func (c *Client) SubscribeBlocks(ctx context.Context) (<-chan BlockEvent, error) {
	blockChan := make(chan BlockEvent, 10)
	rawChan := make(chan json.RawMessage, 10)

	c.handlersMu.Lock()
	if c.handlers["block"] == nil {
		c.handlers["block"] = make([]chan json.RawMessage, 0)
	}
	c.handlers["block"] = append(c.handlers["block"], rawChan)
	c.handlersMu.Unlock()

	params := BlockSubscriptionParams{
		Event: "block",
	}

	if err := c.subscribe(ctx, params); err != nil {
		return nil, err
	}

	go func() {
		defer close(blockChan)
		for {
			select {
			case raw := <-rawChan:
				var event BlockEvent
				if err := json.Unmarshal(raw, &event); err != nil {
					log.Printf("Failed to unmarshal block event: %v", err)
					continue
				}
				select {
				case blockChan <- event:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return blockChan, nil
}

func (c *Client) SubscribeMempool(ctx context.Context) (<-chan MempoolEvent, error) {
	mempoolChan := make(chan MempoolEvent, 10)
	rawChan := make(chan json.RawMessage, 10)

	c.handlersMu.Lock()
	if c.handlers["mempool"] == nil {
		c.handlers["mempool"] = make([]chan json.RawMessage, 0)
	}
	c.handlers["mempool"] = append(c.handlers["mempool"], rawChan)
	c.handlersMu.Unlock()

	params := MempoolSubscriptionParams{
		Event: "mempool",
	}

	if err := c.subscribe(ctx, params); err != nil {
		return nil, err
	}

	go func() {
		defer close(mempoolChan)
		for {
			select {
			case raw := <-rawChan:
				var event MempoolEvent
				if err := json.Unmarshal(raw, &event); err != nil {
					log.Printf("Failed to unmarshal mempool event: %v", err)
					continue
				}
				select {
				case mempoolChan <- event:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return mempoolChan, nil
}

func (c *Client) SubscribeAddressTransactions(ctx context.Context, address string) (<-chan AddressTxUpdateEvent, error) {
	addressTxChan := make(chan AddressTxUpdateEvent, 10)
	rawChan := make(chan json.RawMessage, 10)

	c.handlersMu.Lock()
	if c.handlers["address_tx_update"] == nil {
		c.handlers["address_tx_update"] = make([]chan json.RawMessage, 0)
	}
	c.handlers["address_tx_update"] = append(c.handlers["address_tx_update"], rawChan)
	c.handlersMu.Unlock()

	params := AddressTransactionSubscriptionParams{
		Event:   "address_tx_update",
		Address: address,
	}

	if err := c.subscribe(ctx, params); err != nil {
		return nil, err
	}

	go func() {
		defer close(addressTxChan)
		for {
			select {
			case raw := <-rawChan:
				var event AddressTxUpdateEvent
				if err := json.Unmarshal(raw, &event); err != nil {
					log.Printf("Failed to unmarshal address tx update event: %v", err)
					continue
				}
				select {
				case addressTxChan <- event:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return addressTxChan, nil
}

func (c *Client) subscribe(ctx context.Context, params SubscriptionParams) error {
	id := atomic.AddUint64(&c.idCounter, 1)
	respChan := make(chan *jsonrpcResponse, 1)

	c.pendingMu.Lock()
	c.pending[id] = respChan
	c.pendingMu.Unlock()

	req := jsonrpcRequest{
		JSONRPC: "2.0",
		ID:      id,
		Method:  "subscribe",
		Params:  params,
	}

	if err := c.conn.WriteJSON(req); err != nil {
		return fmt.Errorf("failed to send subscription request: %w", err)
	}

	select {
	case resp := <-respChan:
		if resp.Error != nil {
			return fmt.Errorf("subscription error: %v", resp.Error)
		}
		return nil
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(10 * time.Second):
		return fmt.Errorf("subscription timeout")
	}
}

func (c *Client) listen() {
	defer c.conn.Close()

	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			var msg json.RawMessage
			if err := c.conn.ReadJSON(&msg); err != nil {
				log.Printf("WebSocket read error: %v", err)
				c.cancel()
				return
			}

			var resp jsonrpcResponse
			if err := json.Unmarshal(msg, &resp); err == nil {
				if resp.ID != 0 {
					c.pendingMu.Lock()
					if ch, ok := c.pending[resp.ID]; ok {
						ch <- &resp
						delete(c.pending, resp.ID)
					}
					c.pendingMu.Unlock()
					continue
				}
			}

			var event struct {
				Method string          `json:"method"`
				Params json.RawMessage `json:"params"`
			}
			if err := json.Unmarshal(msg, &event); err != nil {
				log.Printf("Failed to parse event: %v", err)
				continue
			}

			c.handlersMu.Lock()
			handlers := c.handlers[event.Method]
			c.handlersMu.Unlock()

			for _, handler := range handlers {
				select {
				case handler <- msg:
				default:
					log.Printf("Handler channel full, dropping message for %s", event.Method)
				}
			}
		}
	}
}

func (c *Client) Close() error {
	c.cancel()
	return c.conn.Close()
}

func (c *Client) startHeartbeat() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-c.ctx.Done():
			return
		case <-ticker.C:
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Printf("Failed to send ping: %v", err)
				c.cancel()
				return
			}
		}
	}
}
