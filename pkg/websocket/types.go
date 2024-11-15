package websocket

import "encoding/json"

type SubscriptionParams interface {
	GetEvent() string
}

type BlockSubscriptionParams struct {
	Event string `json:"event"`
}

func (p BlockSubscriptionParams) GetEvent() string {
	return p.Event
}

type MempoolSubscriptionParams struct {
	Event string `json:"event"`
}

func (p MempoolSubscriptionParams) GetEvent() string {
	return p.Event
}

type AddressTransactionSubscriptionParams struct {
	Event   string `json:"event"`
	Address string `json:"address"`
}

func (p AddressTransactionSubscriptionParams) GetEvent() string {
	return p.Event
}

type AddressTxUpdate struct {
	Address      string          `json:"address"`
	TxID         string          `json:"tx_id"`
	TxStatus     string          `json:"tx_status"`
	TxType       string          `json:"tx_type"`
	Tx           TransactionInfo `json:"tx"`
	StxSent      string          `json:"stx_sent"`
	StxReceived  string          `json:"stx_received"`
	StxTransfers []StxTransfer   `json:"stx_transfers"`
}

type StacksEvent struct {
	Type  string          `json:"type"`
	Value json.RawMessage `json:"value"`
}

type ContractLogEvent struct {
	ContractID string          `json:"contract_id"`
	Topic      string          `json:"topic"`
	Value      json.RawMessage `json:"value"`
}

type TransactionInfo struct {
	TxID                string        `json:"tx_id"`
	Nonce               int           `json:"nonce"`
	FeeRate             string        `json:"fee_rate"`
	SenderAddress       string        `json:"sender_address"`
	Sponsored           bool          `json:"sponsored"`
	PostConditionMode   string        `json:"post_condition_mode"`
	PostConditions      []interface{} `json:"post_conditions"`
	AnchorMode          string        `json:"anchor_mode"`
	IsUnanchored        bool          `json:"is_unanchored"`
	BlockHash           string        `json:"block_hash"`
	ParentBlockHash     string        `json:"parent_block_hash"`
	BlockHeight         int           `json:"block_height"`
	BlockTime           int64         `json:"block_time"`
	BlockTimeISO        string        `json:"block_time_iso"`
	BurnBlockHeight     int           `json:"burn_block_height"`
	BurnBlockTime       int64         `json:"burn_block_time"`
	BurnBlockTimeISO    string        `json:"burn_block_time_iso"`
	Canonical           bool          `json:"canonical"`
	TxIndex             int           `json:"tx_index"`
	TxStatus            string        `json:"tx_status"`
	TxResult            TxResult      `json:"tx_result"`
	MicroblockHash      string        `json:"microblock_hash"`
	MicroblockSequence  int           `json:"microblock_sequence"`
	MicroblockCanonical bool          `json:"microblock_canonical"`
	EventCount          int           `json:"event_count"`
	Events              []StacksEvent `json:"events"`
	ExecutionCost       ExecutionCost `json:"execution_cost"`
	TxType              string        `json:"tx_type"`
	ContractCall        *ContractCall `json:"contract_call,omitempty"`
}

type TxResult struct {
	Hex  string `json:"hex"`
	Repr string `json:"repr"`
}

type ExecutionCost struct {
	ReadCount   int `json:"execution_cost_read_count"`
	ReadLength  int `json:"execution_cost_read_length"`
	Runtime     int `json:"execution_cost_runtime"`
	WriteCount  int `json:"execution_cost_write_count"`
	WriteLength int `json:"execution_cost_write_length"`
}

type ContractCall struct {
	ContractID        string        `json:"contract_id"`
	FunctionName      string        `json:"function_name"`
	FunctionSignature string        `json:"function_signature"`
	FunctionArgs      []FunctionArg `json:"function_args"`
}

type FunctionArg struct {
	Hex  string `json:"hex"`
	Repr string `json:"repr"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type StxTransfer struct {
	Amount    string `json:"amount"`
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
}

type AddressTxUpdateEvent struct {
	Method string          `json:"method"`
	Params AddressTxUpdate `json:"params"`
}

type Block struct {
	Hash                     string         `json:"hash"`
	BurnBlockHash            string         `json:"burn_block_hash"`
	ParentBlockHash          string         `json:"parent_block_hash"`
	Height                   int            `json:"height"`
	BurnBlockTime            int            `json:"burn_block_time"`
	BlockTime                int            `json:"block_time"`
	Canonical                bool           `json:"canonical"`
	TenureHeight             int            `json:"tenure_height"`
	IndexBlockHash           string         `json:"index_block_hash"`
	ParentMicroblockHash     string         `json:"parent_microblock_hash"`
	ParentMicroblockSequence int            `json:"parent_microblock_sequence"`
	Txs                      []string       `json:"txs"`
	MicroblocksAccepted      []interface{}  `json:"microblocks_accepted"`
	MicroblocksStreamed      []interface{}  `json:"microblocks_streamed"`
	ExecutionCostReadCount   int            `json:"execution_cost_read_count"`
	ExecutionCostReadLength  int            `json:"execution_cost_read_length"`
	ExecutionCostRuntime     int            `json:"execution_cost_runtime"`
	ExecutionCostWriteCount  int            `json:"execution_cost_write_count"`
	ExecutionCostWriteLength int            `json:"execution_cost_write_length"`
	MicroblockTxCount        map[string]int `json:"microblock_tx_count"`
}

type Microblock struct {
	MicroblockHash string   `json:"microblock_hash"`
	ParentHash     string   `json:"parent_hash"`
	Sequence       int      `json:"sequence"`
	TxIDs          []string `json:"txs"`
}

type Transaction struct {
	TxID        string `json:"tx_id"`
	TxStatus    string `json:"tx_status"`
	BlockHeight int    `json:"block_height"`
}

type MempoolTransaction struct {
	TxID     string `json:"tx_id"`
	TxStatus string `json:"tx_status"`
}

type AddressTxNotification struct {
	Address     string      `json:"address"`
	Transaction Transaction `json:"tx"`
}

type AddressBalanceNotification struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
}

type NFTEvent struct {
	EventType       string `json:"event_type"`
	AssetIdentifier string `json:"asset_identifier"`
	Sender          string `json:"sender"`
	Recipient       string `json:"recipient"`
	Value           struct {
		Hex  string `json:"hex"`
		Repr string `json:"repr"`
		Type string `json:"type"`
	} `json:"value"`
}

type BlockEvent struct {
	Method string `json:"method"`
	Params struct {
		Canonical       bool     `json:"canonical"`
		Height          int      `json:"height"`
		Hash            string   `json:"hash"`
		Txs             []string `json:"txs"`
		BlockTime       int64    `json:"block_time"`
		BurnBlockHeight int      `json:"burn_block_height"`
	} `json:"params"`
}

type MempoolEvent struct {
	Method string `json:"method"`
	Params struct {
		TxID        string `json:"tx_id"`
		Status      string `json:"tx_status"`
		ReceiptTime int64  `json:"receipt_time"`
	} `json:"params"`
}

type jsonrpcRequest struct {
	JSONRPC string             `json:"jsonrpc"`
	ID      uint64             `json:"id"`
	Method  string             `json:"method"`
	Params  SubscriptionParams `json:"params,omitempty"`
}

type jsonrpcResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      uint64          `json:"id"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *jsonrpcError   `json:"error,omitempty"`
	Method  string          `json:"method,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
}

type jsonrpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
