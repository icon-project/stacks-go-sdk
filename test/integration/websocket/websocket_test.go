package websocket_test

import (
	"context"
	"testing"
	"time"

	"github.com/icon-project/stacks-go-sdk/pkg/websocket"
)

func TestWebSocketClient(t *testing.T) {
	t.Log("Starting WebSocket client test")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	t.Log("Connecting to WebSocket...")
	client, err := websocket.NewClient("wss://api.testnet.hiro.so")
	if err != nil {
		t.Fatalf("Failed to create WebSocket client: %v", err)
	}
	defer func() {
		t.Log("Closing WebSocket connection")
		client.Close()
	}()

	t.Log("Subscribing to block events...")
	blockChan, err := client.SubscribeBlocks(ctx)
	if err != nil {
		t.Fatalf("Failed to subscribe to blocks: %v", err)
	}

	t.Log("Subscribing to mempool events...")
	mempoolChan, err := client.SubscribeMempool(ctx)
	if err != nil {
		t.Fatalf("Failed to subscribe to mempool: %v", err)
	}

	t.Log("Subscribing to address transaction events...")
	testAddress := "ST1FTM84RHDZ3AB21MNYKA3XQEFB090HZBB81DSFE"
	addressTxChan, err := client.SubscribeAddressTransactions(ctx, testAddress)
	if err != nil {
		t.Fatalf("Failed to subscribe to address transactions: %v", err)
	}

	blockReceived := make(chan struct{})
	mempoolReceived := make(chan struct{})
	addressTxReceived := make(chan struct{})

	go func() {
		select {
		case block := <-blockChan:
			t.Logf("Block received! Height=%d Hash=%s Time=%s TxCount=%d",
				block.Params.Height,
				block.Params.Hash,
				time.Unix(block.Params.BlockTime, 0).Format(time.RFC3339),
				len(block.Params.Txs))
			close(blockReceived)
		case <-ctx.Done():
			t.Log("Context cancelled while waiting for block")
			return
		}
	}()

	go func() {
		select {
		case tx := <-mempoolChan:
			t.Logf("Mempool transaction received! TxID=%s Status=%s Time=%s",
				tx.Params.TxID,
				tx.Params.Status,
				time.Unix(tx.Params.ReceiptTime, 0).Format(time.RFC3339))
			close(mempoolReceived)
		case <-ctx.Done():
			t.Log("Context cancelled while waiting for mempool tx")
			return
		}
	}()

	go func() {
		select {
		case tx := <-addressTxChan:
			t.Logf("Address transaction received! Address=%s TxID=%s Type=%s Status=%s",
				tx.Params.Address,
				tx.Params.TxID,
				tx.Params.TxType,
				tx.Params.TxStatus)
			if tx.Params.Tx.ContractCall != nil {
				t.Logf("Contract call details: Contract=%s Function=%s",
					tx.Params.Tx.ContractCall.ContractID,
					tx.Params.Tx.ContractCall.FunctionName)
			}
			if len(tx.Params.StxTransfers) > 0 {
				t.Logf("STX transfers: Count=%d First transfer: Amount=%s From=%s To=%s",
					len(tx.Params.StxTransfers),
					tx.Params.StxTransfers[0].Amount,
					tx.Params.StxTransfers[0].Sender,
					tx.Params.StxTransfers[0].Recipient)
			}
			close(addressTxReceived)
		case <-ctx.Done():
			t.Log("Context cancelled while waiting for address transaction")
			return
		}
	}()

	t.Log("Waiting for events...")

	select {
	case <-blockReceived:
		t.Log("Successfully received and processed block event")
	case <-time.After(45 * time.Second):
		t.Error("Timeout waiting for block event")
	}

	select {
	case <-mempoolReceived:
		t.Log("Successfully received and processed mempool event")
	case <-time.After(5 * time.Second):
		t.Log("Note: No mempool events received (this is not an error)")
	}

	select {
	case <-addressTxReceived:
		t.Log("Successfully received and processed address transaction event")
	case <-time.After(45 * time.Second):
		t.Log("Note: No address transaction events received (this is not an error)")
	}

	t.Log("Test completed successfully")
}
