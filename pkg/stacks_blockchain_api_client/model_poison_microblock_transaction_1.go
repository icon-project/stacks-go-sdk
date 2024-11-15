/*
Stacks Blockchain API

Welcome to the API reference overview for the [Stacks Blockchain API](https://docs.hiro.so/stacks-blockchain-api).        [Download Postman collection](https://hirosystems.github.io/stacks-blockchain-api/collection.json)

API version: v8.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stacks_blockchain_api_client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PoisonMicroblockTransaction1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PoisonMicroblockTransaction1{}

// PoisonMicroblockTransaction1 struct for PoisonMicroblockTransaction1
type PoisonMicroblockTransaction1 struct {
	// Transaction ID
	TxId string `json:"tx_id"`
	// Used for ordering the transactions originating from and paying from an account. The nonce ensures that a transaction is processed at most once. The nonce counts the number of times an account's owner(s) have authorized a transaction. The first transaction from an account will have a nonce value equal to 0, the second will have a nonce value equal to 1, and so on.
	Nonce int32 `json:"nonce"`
	// Transaction fee as Integer string (64-bit unsigned integer).
	FeeRate string `json:"fee_rate"`
	// Address of the transaction initiator
	SenderAddress string `json:"sender_address"`
	SponsorNonce *int32 `json:"sponsor_nonce,omitempty"`
	// Denotes whether the originating account is the same as the paying account
	Sponsored bool `json:"sponsored"`
	SponsorAddress *string `json:"sponsor_address,omitempty"`
	PostConditionMode TokenTransferTransaction1PostConditionMode `json:"post_condition_mode"`
	PostConditions []TokenTransferTransaction1PostConditionsInner `json:"post_conditions"`
	AnchorMode TokenTransferTransaction1AnchorMode `json:"anchor_mode"`
	// Hash of the blocked this transactions was associated with
	BlockHash string `json:"block_hash"`
	// Height of the block this transactions was associated with
	BlockHeight int32 `json:"block_height"`
	// Unix timestamp (in seconds) indicating when this block was mined.
	BlockTime float32 `json:"block_time"`
	// An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) indicating when this block was mined.
	BlockTimeIso string `json:"block_time_iso"`
	// Unix timestamp (in seconds) indicating when this block was mined.
	BurnBlockTime int32 `json:"burn_block_time"`
	// Height of the anchor burn block.
	BurnBlockHeight int32 `json:"burn_block_height"`
	// An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) timestamp indicating when this block was mined.
	BurnBlockTimeIso string `json:"burn_block_time_iso"`
	// Unix timestamp (in seconds) indicating when this parent block was mined
	ParentBurnBlockTime int32 `json:"parent_burn_block_time"`
	// An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) timestamp indicating when this parent block was mined.
	ParentBurnBlockTimeIso string `json:"parent_burn_block_time_iso"`
	// Set to `true` if block corresponds to the canonical chain tip
	Canonical bool `json:"canonical"`
	// Index of the transaction, indicating the order. Starts at `0` and increases with each transaction
	TxIndex int32 `json:"tx_index"`
	TxStatus TokenTransferTransaction1TxStatus `json:"tx_status"`
	TxResult TokenTransferTransactionTxResult `json:"tx_result"`
	// Number of transaction events
	EventCount int32 `json:"event_count"`
	// Hash of the previous block.
	ParentBlockHash string `json:"parent_block_hash"`
	// True if the transaction is included in a microblock that has not been confirmed by an anchor block.
	IsUnanchored bool `json:"is_unanchored"`
	// The microblock hash that this transaction was streamed in. If the transaction was batched in an anchor block (not included within a microblock) then this value will be an empty string.
	MicroblockHash string `json:"microblock_hash"`
	// The microblock sequence number that this transaction was streamed in. If the transaction was batched in an anchor block (not included within a microblock) then this value will be 2147483647 (0x7fffffff, the max int32 value), this value preserves logical transaction ordering on (block_height, microblock_sequence, tx_index).
	MicroblockSequence int32 `json:"microblock_sequence"`
	// Set to `true` if microblock is anchored in the canonical chain tip, `false` if the transaction was orphaned in a micro-fork.
	MicroblockCanonical bool `json:"microblock_canonical"`
	// Execution cost read count.
	ExecutionCostReadCount int32 `json:"execution_cost_read_count"`
	// Execution cost read length.
	ExecutionCostReadLength int32 `json:"execution_cost_read_length"`
	// Execution cost runtime.
	ExecutionCostRuntime int32 `json:"execution_cost_runtime"`
	// Execution cost write count.
	ExecutionCostWriteCount int32 `json:"execution_cost_write_count"`
	// Execution cost write length.
	ExecutionCostWriteLength int32 `json:"execution_cost_write_length"`
	Events []TokenTransferTransaction1EventsInner `json:"events"`
	TxType string `json:"tx_type"`
	PoisonMicroblock PoisonMicroblockTransactionPoisonMicroblock `json:"poison_microblock"`
}

type _PoisonMicroblockTransaction1 PoisonMicroblockTransaction1

// NewPoisonMicroblockTransaction1 instantiates a new PoisonMicroblockTransaction1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoisonMicroblockTransaction1(txId string, nonce int32, feeRate string, senderAddress string, sponsored bool, postConditionMode TokenTransferTransaction1PostConditionMode, postConditions []TokenTransferTransaction1PostConditionsInner, anchorMode TokenTransferTransaction1AnchorMode, blockHash string, blockHeight int32, blockTime float32, blockTimeIso string, burnBlockTime int32, burnBlockHeight int32, burnBlockTimeIso string, parentBurnBlockTime int32, parentBurnBlockTimeIso string, canonical bool, txIndex int32, txStatus TokenTransferTransaction1TxStatus, txResult TokenTransferTransactionTxResult, eventCount int32, parentBlockHash string, isUnanchored bool, microblockHash string, microblockSequence int32, microblockCanonical bool, executionCostReadCount int32, executionCostReadLength int32, executionCostRuntime int32, executionCostWriteCount int32, executionCostWriteLength int32, events []TokenTransferTransaction1EventsInner, txType string, poisonMicroblock PoisonMicroblockTransactionPoisonMicroblock) *PoisonMicroblockTransaction1 {
	this := PoisonMicroblockTransaction1{}
	this.TxId = txId
	this.Nonce = nonce
	this.FeeRate = feeRate
	this.SenderAddress = senderAddress
	this.Sponsored = sponsored
	this.PostConditionMode = postConditionMode
	this.PostConditions = postConditions
	this.AnchorMode = anchorMode
	this.BlockHash = blockHash
	this.BlockHeight = blockHeight
	this.BlockTime = blockTime
	this.BlockTimeIso = blockTimeIso
	this.BurnBlockTime = burnBlockTime
	this.BurnBlockHeight = burnBlockHeight
	this.BurnBlockTimeIso = burnBlockTimeIso
	this.ParentBurnBlockTime = parentBurnBlockTime
	this.ParentBurnBlockTimeIso = parentBurnBlockTimeIso
	this.Canonical = canonical
	this.TxIndex = txIndex
	this.TxStatus = txStatus
	this.TxResult = txResult
	this.EventCount = eventCount
	this.ParentBlockHash = parentBlockHash
	this.IsUnanchored = isUnanchored
	this.MicroblockHash = microblockHash
	this.MicroblockSequence = microblockSequence
	this.MicroblockCanonical = microblockCanonical
	this.ExecutionCostReadCount = executionCostReadCount
	this.ExecutionCostReadLength = executionCostReadLength
	this.ExecutionCostRuntime = executionCostRuntime
	this.ExecutionCostWriteCount = executionCostWriteCount
	this.ExecutionCostWriteLength = executionCostWriteLength
	this.Events = events
	this.TxType = txType
	this.PoisonMicroblock = poisonMicroblock
	return &this
}

// NewPoisonMicroblockTransaction1WithDefaults instantiates a new PoisonMicroblockTransaction1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoisonMicroblockTransaction1WithDefaults() *PoisonMicroblockTransaction1 {
	this := PoisonMicroblockTransaction1{}
	return &this
}

// GetTxId returns the TxId field value
func (o *PoisonMicroblockTransaction1) GetTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *PoisonMicroblockTransaction1) SetTxId(v string) {
	o.TxId = v
}

// GetNonce returns the Nonce field value
func (o *PoisonMicroblockTransaction1) GetNonce() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetNonceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *PoisonMicroblockTransaction1) SetNonce(v int32) {
	o.Nonce = v
}

// GetFeeRate returns the FeeRate field value
func (o *PoisonMicroblockTransaction1) GetFeeRate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeeRate
}

// GetFeeRateOk returns a tuple with the FeeRate field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetFeeRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeRate, true
}

// SetFeeRate sets field value
func (o *PoisonMicroblockTransaction1) SetFeeRate(v string) {
	o.FeeRate = v
}

// GetSenderAddress returns the SenderAddress field value
func (o *PoisonMicroblockTransaction1) GetSenderAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SenderAddress
}

// GetSenderAddressOk returns a tuple with the SenderAddress field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetSenderAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenderAddress, true
}

// SetSenderAddress sets field value
func (o *PoisonMicroblockTransaction1) SetSenderAddress(v string) {
	o.SenderAddress = v
}

// GetSponsorNonce returns the SponsorNonce field value if set, zero value otherwise.
func (o *PoisonMicroblockTransaction1) GetSponsorNonce() int32 {
	if o == nil || IsNil(o.SponsorNonce) {
		var ret int32
		return ret
	}
	return *o.SponsorNonce
}

// GetSponsorNonceOk returns a tuple with the SponsorNonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetSponsorNonceOk() (*int32, bool) {
	if o == nil || IsNil(o.SponsorNonce) {
		return nil, false
	}
	return o.SponsorNonce, true
}

// HasSponsorNonce returns a boolean if a field has been set.
func (o *PoisonMicroblockTransaction1) HasSponsorNonce() bool {
	if o != nil && !IsNil(o.SponsorNonce) {
		return true
	}

	return false
}

// SetSponsorNonce gets a reference to the given int32 and assigns it to the SponsorNonce field.
func (o *PoisonMicroblockTransaction1) SetSponsorNonce(v int32) {
	o.SponsorNonce = &v
}

// GetSponsored returns the Sponsored field value
func (o *PoisonMicroblockTransaction1) GetSponsored() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Sponsored
}

// GetSponsoredOk returns a tuple with the Sponsored field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetSponsoredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sponsored, true
}

// SetSponsored sets field value
func (o *PoisonMicroblockTransaction1) SetSponsored(v bool) {
	o.Sponsored = v
}

// GetSponsorAddress returns the SponsorAddress field value if set, zero value otherwise.
func (o *PoisonMicroblockTransaction1) GetSponsorAddress() string {
	if o == nil || IsNil(o.SponsorAddress) {
		var ret string
		return ret
	}
	return *o.SponsorAddress
}

// GetSponsorAddressOk returns a tuple with the SponsorAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetSponsorAddressOk() (*string, bool) {
	if o == nil || IsNil(o.SponsorAddress) {
		return nil, false
	}
	return o.SponsorAddress, true
}

// HasSponsorAddress returns a boolean if a field has been set.
func (o *PoisonMicroblockTransaction1) HasSponsorAddress() bool {
	if o != nil && !IsNil(o.SponsorAddress) {
		return true
	}

	return false
}

// SetSponsorAddress gets a reference to the given string and assigns it to the SponsorAddress field.
func (o *PoisonMicroblockTransaction1) SetSponsorAddress(v string) {
	o.SponsorAddress = &v
}

// GetPostConditionMode returns the PostConditionMode field value
func (o *PoisonMicroblockTransaction1) GetPostConditionMode() TokenTransferTransaction1PostConditionMode {
	if o == nil {
		var ret TokenTransferTransaction1PostConditionMode
		return ret
	}

	return o.PostConditionMode
}

// GetPostConditionModeOk returns a tuple with the PostConditionMode field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetPostConditionModeOk() (*TokenTransferTransaction1PostConditionMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostConditionMode, true
}

// SetPostConditionMode sets field value
func (o *PoisonMicroblockTransaction1) SetPostConditionMode(v TokenTransferTransaction1PostConditionMode) {
	o.PostConditionMode = v
}

// GetPostConditions returns the PostConditions field value
func (o *PoisonMicroblockTransaction1) GetPostConditions() []TokenTransferTransaction1PostConditionsInner {
	if o == nil {
		var ret []TokenTransferTransaction1PostConditionsInner
		return ret
	}

	return o.PostConditions
}

// GetPostConditionsOk returns a tuple with the PostConditions field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetPostConditionsOk() ([]TokenTransferTransaction1PostConditionsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostConditions, true
}

// SetPostConditions sets field value
func (o *PoisonMicroblockTransaction1) SetPostConditions(v []TokenTransferTransaction1PostConditionsInner) {
	o.PostConditions = v
}

// GetAnchorMode returns the AnchorMode field value
func (o *PoisonMicroblockTransaction1) GetAnchorMode() TokenTransferTransaction1AnchorMode {
	if o == nil {
		var ret TokenTransferTransaction1AnchorMode
		return ret
	}

	return o.AnchorMode
}

// GetAnchorModeOk returns a tuple with the AnchorMode field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetAnchorModeOk() (*TokenTransferTransaction1AnchorMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnchorMode, true
}

// SetAnchorMode sets field value
func (o *PoisonMicroblockTransaction1) SetAnchorMode(v TokenTransferTransaction1AnchorMode) {
	o.AnchorMode = v
}

// GetBlockHash returns the BlockHash field value
func (o *PoisonMicroblockTransaction1) GetBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockHash
}

// GetBlockHashOk returns a tuple with the BlockHash field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHash, true
}

// SetBlockHash sets field value
func (o *PoisonMicroblockTransaction1) SetBlockHash(v string) {
	o.BlockHash = v
}

// GetBlockHeight returns the BlockHeight field value
func (o *PoisonMicroblockTransaction1) GetBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BlockHeight
}

// GetBlockHeightOk returns a tuple with the BlockHeight field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHeight, true
}

// SetBlockHeight sets field value
func (o *PoisonMicroblockTransaction1) SetBlockHeight(v int32) {
	o.BlockHeight = v
}

// GetBlockTime returns the BlockTime field value
func (o *PoisonMicroblockTransaction1) GetBlockTime() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BlockTime
}

// GetBlockTimeOk returns a tuple with the BlockTime field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetBlockTimeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockTime, true
}

// SetBlockTime sets field value
func (o *PoisonMicroblockTransaction1) SetBlockTime(v float32) {
	o.BlockTime = v
}

// GetBlockTimeIso returns the BlockTimeIso field value
func (o *PoisonMicroblockTransaction1) GetBlockTimeIso() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockTimeIso
}

// GetBlockTimeIsoOk returns a tuple with the BlockTimeIso field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetBlockTimeIsoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockTimeIso, true
}

// SetBlockTimeIso sets field value
func (o *PoisonMicroblockTransaction1) SetBlockTimeIso(v string) {
	o.BlockTimeIso = v
}

// GetBurnBlockTime returns the BurnBlockTime field value
func (o *PoisonMicroblockTransaction1) GetBurnBlockTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BurnBlockTime
}

// GetBurnBlockTimeOk returns a tuple with the BurnBlockTime field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetBurnBlockTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockTime, true
}

// SetBurnBlockTime sets field value
func (o *PoisonMicroblockTransaction1) SetBurnBlockTime(v int32) {
	o.BurnBlockTime = v
}

// GetBurnBlockHeight returns the BurnBlockHeight field value
func (o *PoisonMicroblockTransaction1) GetBurnBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BurnBlockHeight
}

// GetBurnBlockHeightOk returns a tuple with the BurnBlockHeight field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetBurnBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockHeight, true
}

// SetBurnBlockHeight sets field value
func (o *PoisonMicroblockTransaction1) SetBurnBlockHeight(v int32) {
	o.BurnBlockHeight = v
}

// GetBurnBlockTimeIso returns the BurnBlockTimeIso field value
func (o *PoisonMicroblockTransaction1) GetBurnBlockTimeIso() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BurnBlockTimeIso
}

// GetBurnBlockTimeIsoOk returns a tuple with the BurnBlockTimeIso field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetBurnBlockTimeIsoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnBlockTimeIso, true
}

// SetBurnBlockTimeIso sets field value
func (o *PoisonMicroblockTransaction1) SetBurnBlockTimeIso(v string) {
	o.BurnBlockTimeIso = v
}

// GetParentBurnBlockTime returns the ParentBurnBlockTime field value
func (o *PoisonMicroblockTransaction1) GetParentBurnBlockTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ParentBurnBlockTime
}

// GetParentBurnBlockTimeOk returns a tuple with the ParentBurnBlockTime field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetParentBurnBlockTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentBurnBlockTime, true
}

// SetParentBurnBlockTime sets field value
func (o *PoisonMicroblockTransaction1) SetParentBurnBlockTime(v int32) {
	o.ParentBurnBlockTime = v
}

// GetParentBurnBlockTimeIso returns the ParentBurnBlockTimeIso field value
func (o *PoisonMicroblockTransaction1) GetParentBurnBlockTimeIso() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentBurnBlockTimeIso
}

// GetParentBurnBlockTimeIsoOk returns a tuple with the ParentBurnBlockTimeIso field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetParentBurnBlockTimeIsoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentBurnBlockTimeIso, true
}

// SetParentBurnBlockTimeIso sets field value
func (o *PoisonMicroblockTransaction1) SetParentBurnBlockTimeIso(v string) {
	o.ParentBurnBlockTimeIso = v
}

// GetCanonical returns the Canonical field value
func (o *PoisonMicroblockTransaction1) GetCanonical() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Canonical
}

// GetCanonicalOk returns a tuple with the Canonical field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetCanonicalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Canonical, true
}

// SetCanonical sets field value
func (o *PoisonMicroblockTransaction1) SetCanonical(v bool) {
	o.Canonical = v
}

// GetTxIndex returns the TxIndex field value
func (o *PoisonMicroblockTransaction1) GetTxIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TxIndex
}

// GetTxIndexOk returns a tuple with the TxIndex field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetTxIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxIndex, true
}

// SetTxIndex sets field value
func (o *PoisonMicroblockTransaction1) SetTxIndex(v int32) {
	o.TxIndex = v
}

// GetTxStatus returns the TxStatus field value
func (o *PoisonMicroblockTransaction1) GetTxStatus() TokenTransferTransaction1TxStatus {
	if o == nil {
		var ret TokenTransferTransaction1TxStatus
		return ret
	}

	return o.TxStatus
}

// GetTxStatusOk returns a tuple with the TxStatus field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetTxStatusOk() (*TokenTransferTransaction1TxStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxStatus, true
}

// SetTxStatus sets field value
func (o *PoisonMicroblockTransaction1) SetTxStatus(v TokenTransferTransaction1TxStatus) {
	o.TxStatus = v
}

// GetTxResult returns the TxResult field value
func (o *PoisonMicroblockTransaction1) GetTxResult() TokenTransferTransactionTxResult {
	if o == nil {
		var ret TokenTransferTransactionTxResult
		return ret
	}

	return o.TxResult
}

// GetTxResultOk returns a tuple with the TxResult field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetTxResultOk() (*TokenTransferTransactionTxResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxResult, true
}

// SetTxResult sets field value
func (o *PoisonMicroblockTransaction1) SetTxResult(v TokenTransferTransactionTxResult) {
	o.TxResult = v
}

// GetEventCount returns the EventCount field value
func (o *PoisonMicroblockTransaction1) GetEventCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventCount
}

// GetEventCountOk returns a tuple with the EventCount field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetEventCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventCount, true
}

// SetEventCount sets field value
func (o *PoisonMicroblockTransaction1) SetEventCount(v int32) {
	o.EventCount = v
}

// GetParentBlockHash returns the ParentBlockHash field value
func (o *PoisonMicroblockTransaction1) GetParentBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentBlockHash
}

// GetParentBlockHashOk returns a tuple with the ParentBlockHash field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetParentBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentBlockHash, true
}

// SetParentBlockHash sets field value
func (o *PoisonMicroblockTransaction1) SetParentBlockHash(v string) {
	o.ParentBlockHash = v
}

// GetIsUnanchored returns the IsUnanchored field value
func (o *PoisonMicroblockTransaction1) GetIsUnanchored() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsUnanchored
}

// GetIsUnanchoredOk returns a tuple with the IsUnanchored field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetIsUnanchoredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsUnanchored, true
}

// SetIsUnanchored sets field value
func (o *PoisonMicroblockTransaction1) SetIsUnanchored(v bool) {
	o.IsUnanchored = v
}

// GetMicroblockHash returns the MicroblockHash field value
func (o *PoisonMicroblockTransaction1) GetMicroblockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicroblockHash
}

// GetMicroblockHashOk returns a tuple with the MicroblockHash field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetMicroblockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicroblockHash, true
}

// SetMicroblockHash sets field value
func (o *PoisonMicroblockTransaction1) SetMicroblockHash(v string) {
	o.MicroblockHash = v
}

// GetMicroblockSequence returns the MicroblockSequence field value
func (o *PoisonMicroblockTransaction1) GetMicroblockSequence() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MicroblockSequence
}

// GetMicroblockSequenceOk returns a tuple with the MicroblockSequence field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetMicroblockSequenceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicroblockSequence, true
}

// SetMicroblockSequence sets field value
func (o *PoisonMicroblockTransaction1) SetMicroblockSequence(v int32) {
	o.MicroblockSequence = v
}

// GetMicroblockCanonical returns the MicroblockCanonical field value
func (o *PoisonMicroblockTransaction1) GetMicroblockCanonical() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MicroblockCanonical
}

// GetMicroblockCanonicalOk returns a tuple with the MicroblockCanonical field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetMicroblockCanonicalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicroblockCanonical, true
}

// SetMicroblockCanonical sets field value
func (o *PoisonMicroblockTransaction1) SetMicroblockCanonical(v bool) {
	o.MicroblockCanonical = v
}

// GetExecutionCostReadCount returns the ExecutionCostReadCount field value
func (o *PoisonMicroblockTransaction1) GetExecutionCostReadCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExecutionCostReadCount
}

// GetExecutionCostReadCountOk returns a tuple with the ExecutionCostReadCount field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetExecutionCostReadCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionCostReadCount, true
}

// SetExecutionCostReadCount sets field value
func (o *PoisonMicroblockTransaction1) SetExecutionCostReadCount(v int32) {
	o.ExecutionCostReadCount = v
}

// GetExecutionCostReadLength returns the ExecutionCostReadLength field value
func (o *PoisonMicroblockTransaction1) GetExecutionCostReadLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExecutionCostReadLength
}

// GetExecutionCostReadLengthOk returns a tuple with the ExecutionCostReadLength field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetExecutionCostReadLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionCostReadLength, true
}

// SetExecutionCostReadLength sets field value
func (o *PoisonMicroblockTransaction1) SetExecutionCostReadLength(v int32) {
	o.ExecutionCostReadLength = v
}

// GetExecutionCostRuntime returns the ExecutionCostRuntime field value
func (o *PoisonMicroblockTransaction1) GetExecutionCostRuntime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExecutionCostRuntime
}

// GetExecutionCostRuntimeOk returns a tuple with the ExecutionCostRuntime field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetExecutionCostRuntimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionCostRuntime, true
}

// SetExecutionCostRuntime sets field value
func (o *PoisonMicroblockTransaction1) SetExecutionCostRuntime(v int32) {
	o.ExecutionCostRuntime = v
}

// GetExecutionCostWriteCount returns the ExecutionCostWriteCount field value
func (o *PoisonMicroblockTransaction1) GetExecutionCostWriteCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExecutionCostWriteCount
}

// GetExecutionCostWriteCountOk returns a tuple with the ExecutionCostWriteCount field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetExecutionCostWriteCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionCostWriteCount, true
}

// SetExecutionCostWriteCount sets field value
func (o *PoisonMicroblockTransaction1) SetExecutionCostWriteCount(v int32) {
	o.ExecutionCostWriteCount = v
}

// GetExecutionCostWriteLength returns the ExecutionCostWriteLength field value
func (o *PoisonMicroblockTransaction1) GetExecutionCostWriteLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExecutionCostWriteLength
}

// GetExecutionCostWriteLengthOk returns a tuple with the ExecutionCostWriteLength field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetExecutionCostWriteLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionCostWriteLength, true
}

// SetExecutionCostWriteLength sets field value
func (o *PoisonMicroblockTransaction1) SetExecutionCostWriteLength(v int32) {
	o.ExecutionCostWriteLength = v
}

// GetEvents returns the Events field value
func (o *PoisonMicroblockTransaction1) GetEvents() []TokenTransferTransaction1EventsInner {
	if o == nil {
		var ret []TokenTransferTransaction1EventsInner
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetEventsOk() ([]TokenTransferTransaction1EventsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *PoisonMicroblockTransaction1) SetEvents(v []TokenTransferTransaction1EventsInner) {
	o.Events = v
}

// GetTxType returns the TxType field value
func (o *PoisonMicroblockTransaction1) GetTxType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxType
}

// GetTxTypeOk returns a tuple with the TxType field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetTxTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxType, true
}

// SetTxType sets field value
func (o *PoisonMicroblockTransaction1) SetTxType(v string) {
	o.TxType = v
}

// GetPoisonMicroblock returns the PoisonMicroblock field value
func (o *PoisonMicroblockTransaction1) GetPoisonMicroblock() PoisonMicroblockTransactionPoisonMicroblock {
	if o == nil {
		var ret PoisonMicroblockTransactionPoisonMicroblock
		return ret
	}

	return o.PoisonMicroblock
}

// GetPoisonMicroblockOk returns a tuple with the PoisonMicroblock field value
// and a boolean to check if the value has been set.
func (o *PoisonMicroblockTransaction1) GetPoisonMicroblockOk() (*PoisonMicroblockTransactionPoisonMicroblock, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoisonMicroblock, true
}

// SetPoisonMicroblock sets field value
func (o *PoisonMicroblockTransaction1) SetPoisonMicroblock(v PoisonMicroblockTransactionPoisonMicroblock) {
	o.PoisonMicroblock = v
}

func (o PoisonMicroblockTransaction1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoisonMicroblockTransaction1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tx_id"] = o.TxId
	toSerialize["nonce"] = o.Nonce
	toSerialize["fee_rate"] = o.FeeRate
	toSerialize["sender_address"] = o.SenderAddress
	if !IsNil(o.SponsorNonce) {
		toSerialize["sponsor_nonce"] = o.SponsorNonce
	}
	toSerialize["sponsored"] = o.Sponsored
	if !IsNil(o.SponsorAddress) {
		toSerialize["sponsor_address"] = o.SponsorAddress
	}
	toSerialize["post_condition_mode"] = o.PostConditionMode
	toSerialize["post_conditions"] = o.PostConditions
	toSerialize["anchor_mode"] = o.AnchorMode
	toSerialize["block_hash"] = o.BlockHash
	toSerialize["block_height"] = o.BlockHeight
	toSerialize["block_time"] = o.BlockTime
	toSerialize["block_time_iso"] = o.BlockTimeIso
	toSerialize["burn_block_time"] = o.BurnBlockTime
	toSerialize["burn_block_height"] = o.BurnBlockHeight
	toSerialize["burn_block_time_iso"] = o.BurnBlockTimeIso
	toSerialize["parent_burn_block_time"] = o.ParentBurnBlockTime
	toSerialize["parent_burn_block_time_iso"] = o.ParentBurnBlockTimeIso
	toSerialize["canonical"] = o.Canonical
	toSerialize["tx_index"] = o.TxIndex
	toSerialize["tx_status"] = o.TxStatus
	toSerialize["tx_result"] = o.TxResult
	toSerialize["event_count"] = o.EventCount
	toSerialize["parent_block_hash"] = o.ParentBlockHash
	toSerialize["is_unanchored"] = o.IsUnanchored
	toSerialize["microblock_hash"] = o.MicroblockHash
	toSerialize["microblock_sequence"] = o.MicroblockSequence
	toSerialize["microblock_canonical"] = o.MicroblockCanonical
	toSerialize["execution_cost_read_count"] = o.ExecutionCostReadCount
	toSerialize["execution_cost_read_length"] = o.ExecutionCostReadLength
	toSerialize["execution_cost_runtime"] = o.ExecutionCostRuntime
	toSerialize["execution_cost_write_count"] = o.ExecutionCostWriteCount
	toSerialize["execution_cost_write_length"] = o.ExecutionCostWriteLength
	toSerialize["events"] = o.Events
	toSerialize["tx_type"] = o.TxType
	toSerialize["poison_microblock"] = o.PoisonMicroblock
	return toSerialize, nil
}

func (o *PoisonMicroblockTransaction1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tx_id",
		"nonce",
		"fee_rate",
		"sender_address",
		"sponsored",
		"post_condition_mode",
		"post_conditions",
		"anchor_mode",
		"block_hash",
		"block_height",
		"block_time",
		"block_time_iso",
		"burn_block_time",
		"burn_block_height",
		"burn_block_time_iso",
		"parent_burn_block_time",
		"parent_burn_block_time_iso",
		"canonical",
		"tx_index",
		"tx_status",
		"tx_result",
		"event_count",
		"parent_block_hash",
		"is_unanchored",
		"microblock_hash",
		"microblock_sequence",
		"microblock_canonical",
		"execution_cost_read_count",
		"execution_cost_read_length",
		"execution_cost_runtime",
		"execution_cost_write_count",
		"execution_cost_write_length",
		"events",
		"tx_type",
		"poison_microblock",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPoisonMicroblockTransaction1 := _PoisonMicroblockTransaction1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPoisonMicroblockTransaction1)

	if err != nil {
		return err
	}

	*o = PoisonMicroblockTransaction1(varPoisonMicroblockTransaction1)

	return err
}

type NullablePoisonMicroblockTransaction1 struct {
	value *PoisonMicroblockTransaction1
	isSet bool
}

func (v NullablePoisonMicroblockTransaction1) Get() *PoisonMicroblockTransaction1 {
	return v.value
}

func (v *NullablePoisonMicroblockTransaction1) Set(val *PoisonMicroblockTransaction1) {
	v.value = val
	v.isSet = true
}

func (v NullablePoisonMicroblockTransaction1) IsSet() bool {
	return v.isSet
}

func (v *NullablePoisonMicroblockTransaction1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoisonMicroblockTransaction1(val *PoisonMicroblockTransaction1) *NullablePoisonMicroblockTransaction1 {
	return &NullablePoisonMicroblockTransaction1{value: val, isSet: true}
}

func (v NullablePoisonMicroblockTransaction1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoisonMicroblockTransaction1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


