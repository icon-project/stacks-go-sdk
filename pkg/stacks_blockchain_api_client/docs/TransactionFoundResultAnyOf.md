# TransactionFoundResultAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxId** | **string** | Transaction ID | 
**Nonce** | **int32** | Used for ordering the transactions originating from and paying from an account. The nonce ensures that a transaction is processed at most once. The nonce counts the number of times an account&#39;s owner(s) have authorized a transaction. The first transaction from an account will have a nonce value equal to 0, the second will have a nonce value equal to 1, and so on. | 
**FeeRate** | **string** | Transaction fee as Integer string (64-bit unsigned integer). | 
**SenderAddress** | **string** | Address of the transaction initiator | 
**SponsorNonce** | Pointer to **int32** |  | [optional] 
**Sponsored** | **bool** | Denotes whether the originating account is the same as the paying account | 
**SponsorAddress** | Pointer to **string** |  | [optional] 
**PostConditionMode** | [**TokenTransferTransaction1PostConditionMode**](TokenTransferTransaction1PostConditionMode.md) |  | 
**PostConditions** | [**[]TokenTransferTransaction1PostConditionsInner**](TokenTransferTransaction1PostConditionsInner.md) |  | 
**AnchorMode** | [**TokenTransferTransaction1AnchorMode**](TokenTransferTransaction1AnchorMode.md) |  | 
**BlockHash** | **string** | Hash of the blocked this transactions was associated with | 
**BlockHeight** | **int32** | Height of the block this transactions was associated with | 
**BlockTime** | **float32** | Unix timestamp (in seconds) indicating when this block was mined. | 
**BlockTimeIso** | **string** | An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) indicating when this block was mined. | 
**BurnBlockTime** | **int32** | Unix timestamp (in seconds) indicating when this block was mined. | 
**BurnBlockHeight** | **int32** | Height of the anchor burn block. | 
**BurnBlockTimeIso** | **string** | An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) timestamp indicating when this block was mined. | 
**ParentBurnBlockTime** | **int32** | Unix timestamp (in seconds) indicating when this parent block was mined | 
**ParentBurnBlockTimeIso** | **string** | An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) timestamp indicating when this parent block was mined. | 
**Canonical** | **bool** | Set to &#x60;true&#x60; if block corresponds to the canonical chain tip | 
**TxIndex** | **int32** | Index of the transaction, indicating the order. Starts at &#x60;0&#x60; and increases with each transaction | 
**TxStatus** | [**TokenTransferTransaction1TxStatus**](TokenTransferTransaction1TxStatus.md) |  | 
**TxResult** | [**TokenTransferTransactionTxResult**](TokenTransferTransactionTxResult.md) |  | 
**EventCount** | **int32** | Number of transaction events | 
**ParentBlockHash** | **string** | Hash of the previous block. | 
**IsUnanchored** | **bool** | True if the transaction is included in a microblock that has not been confirmed by an anchor block. | 
**MicroblockHash** | **string** | The microblock hash that this transaction was streamed in. If the transaction was batched in an anchor block (not included within a microblock) then this value will be an empty string. | 
**MicroblockSequence** | **int32** | The microblock sequence number that this transaction was streamed in. If the transaction was batched in an anchor block (not included within a microblock) then this value will be 2147483647 (0x7fffffff, the max int32 value), this value preserves logical transaction ordering on (block_height, microblock_sequence, tx_index). | 
**MicroblockCanonical** | **bool** | Set to &#x60;true&#x60; if microblock is anchored in the canonical chain tip, &#x60;false&#x60; if the transaction was orphaned in a micro-fork. | 
**ExecutionCostReadCount** | **int32** | Execution cost read count. | 
**ExecutionCostReadLength** | **int32** | Execution cost read length. | 
**ExecutionCostRuntime** | **int32** | Execution cost runtime. | 
**ExecutionCostWriteCount** | **int32** | Execution cost write count. | 
**ExecutionCostWriteLength** | **int32** | Execution cost write length. | 
**Events** | [**[]TokenTransferTransaction1EventsInner**](TokenTransferTransaction1EventsInner.md) |  | 
**TxType** | **string** |  | 
**TokenTransfer** | [**TokenTransferTransactionTokenTransfer**](TokenTransferTransactionTokenTransfer.md) |  | 
**SmartContract** | [**SmartContractTransactionSmartContract**](SmartContractTransactionSmartContract.md) |  | 
**ContractCall** | [**ContractCallTransactionContractCall**](ContractCallTransactionContractCall.md) |  | 
**PoisonMicroblock** | [**PoisonMicroblockTransactionPoisonMicroblock**](PoisonMicroblockTransactionPoisonMicroblock.md) |  | 
**CoinbasePayload** | [**CoinbaseTransactionCoinbasePayload**](CoinbaseTransactionCoinbasePayload.md) |  | 
**TenureChangePayload** | [**TenureChangeTransaction1TenureChangePayload**](TenureChangeTransaction1TenureChangePayload.md) |  | 

## Methods

### NewTransactionFoundResultAnyOf

`func NewTransactionFoundResultAnyOf(txId string, nonce int32, feeRate string, senderAddress string, sponsored bool, postConditionMode TokenTransferTransaction1PostConditionMode, postConditions []TokenTransferTransaction1PostConditionsInner, anchorMode TokenTransferTransaction1AnchorMode, blockHash string, blockHeight int32, blockTime float32, blockTimeIso string, burnBlockTime int32, burnBlockHeight int32, burnBlockTimeIso string, parentBurnBlockTime int32, parentBurnBlockTimeIso string, canonical bool, txIndex int32, txStatus TokenTransferTransaction1TxStatus, txResult TokenTransferTransactionTxResult, eventCount int32, parentBlockHash string, isUnanchored bool, microblockHash string, microblockSequence int32, microblockCanonical bool, executionCostReadCount int32, executionCostReadLength int32, executionCostRuntime int32, executionCostWriteCount int32, executionCostWriteLength int32, events []TokenTransferTransaction1EventsInner, txType string, tokenTransfer TokenTransferTransactionTokenTransfer, smartContract SmartContractTransactionSmartContract, contractCall ContractCallTransactionContractCall, poisonMicroblock PoisonMicroblockTransactionPoisonMicroblock, coinbasePayload CoinbaseTransactionCoinbasePayload, tenureChangePayload TenureChangeTransaction1TenureChangePayload, ) *TransactionFoundResultAnyOf`

NewTransactionFoundResultAnyOf instantiates a new TransactionFoundResultAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionFoundResultAnyOfWithDefaults

`func NewTransactionFoundResultAnyOfWithDefaults() *TransactionFoundResultAnyOf`

NewTransactionFoundResultAnyOfWithDefaults instantiates a new TransactionFoundResultAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxId

`func (o *TransactionFoundResultAnyOf) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *TransactionFoundResultAnyOf) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *TransactionFoundResultAnyOf) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetNonce

`func (o *TransactionFoundResultAnyOf) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *TransactionFoundResultAnyOf) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *TransactionFoundResultAnyOf) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetFeeRate

`func (o *TransactionFoundResultAnyOf) GetFeeRate() string`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *TransactionFoundResultAnyOf) GetFeeRateOk() (*string, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *TransactionFoundResultAnyOf) SetFeeRate(v string)`

SetFeeRate sets FeeRate field to given value.


### GetSenderAddress

`func (o *TransactionFoundResultAnyOf) GetSenderAddress() string`

GetSenderAddress returns the SenderAddress field if non-nil, zero value otherwise.

### GetSenderAddressOk

`func (o *TransactionFoundResultAnyOf) GetSenderAddressOk() (*string, bool)`

GetSenderAddressOk returns a tuple with the SenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAddress

`func (o *TransactionFoundResultAnyOf) SetSenderAddress(v string)`

SetSenderAddress sets SenderAddress field to given value.


### GetSponsorNonce

`func (o *TransactionFoundResultAnyOf) GetSponsorNonce() int32`

GetSponsorNonce returns the SponsorNonce field if non-nil, zero value otherwise.

### GetSponsorNonceOk

`func (o *TransactionFoundResultAnyOf) GetSponsorNonceOk() (*int32, bool)`

GetSponsorNonceOk returns a tuple with the SponsorNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorNonce

`func (o *TransactionFoundResultAnyOf) SetSponsorNonce(v int32)`

SetSponsorNonce sets SponsorNonce field to given value.

### HasSponsorNonce

`func (o *TransactionFoundResultAnyOf) HasSponsorNonce() bool`

HasSponsorNonce returns a boolean if a field has been set.

### GetSponsored

`func (o *TransactionFoundResultAnyOf) GetSponsored() bool`

GetSponsored returns the Sponsored field if non-nil, zero value otherwise.

### GetSponsoredOk

`func (o *TransactionFoundResultAnyOf) GetSponsoredOk() (*bool, bool)`

GetSponsoredOk returns a tuple with the Sponsored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsored

`func (o *TransactionFoundResultAnyOf) SetSponsored(v bool)`

SetSponsored sets Sponsored field to given value.


### GetSponsorAddress

`func (o *TransactionFoundResultAnyOf) GetSponsorAddress() string`

GetSponsorAddress returns the SponsorAddress field if non-nil, zero value otherwise.

### GetSponsorAddressOk

`func (o *TransactionFoundResultAnyOf) GetSponsorAddressOk() (*string, bool)`

GetSponsorAddressOk returns a tuple with the SponsorAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorAddress

`func (o *TransactionFoundResultAnyOf) SetSponsorAddress(v string)`

SetSponsorAddress sets SponsorAddress field to given value.

### HasSponsorAddress

`func (o *TransactionFoundResultAnyOf) HasSponsorAddress() bool`

HasSponsorAddress returns a boolean if a field has been set.

### GetPostConditionMode

`func (o *TransactionFoundResultAnyOf) GetPostConditionMode() TokenTransferTransaction1PostConditionMode`

GetPostConditionMode returns the PostConditionMode field if non-nil, zero value otherwise.

### GetPostConditionModeOk

`func (o *TransactionFoundResultAnyOf) GetPostConditionModeOk() (*TokenTransferTransaction1PostConditionMode, bool)`

GetPostConditionModeOk returns a tuple with the PostConditionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostConditionMode

`func (o *TransactionFoundResultAnyOf) SetPostConditionMode(v TokenTransferTransaction1PostConditionMode)`

SetPostConditionMode sets PostConditionMode field to given value.


### GetPostConditions

`func (o *TransactionFoundResultAnyOf) GetPostConditions() []TokenTransferTransaction1PostConditionsInner`

GetPostConditions returns the PostConditions field if non-nil, zero value otherwise.

### GetPostConditionsOk

`func (o *TransactionFoundResultAnyOf) GetPostConditionsOk() (*[]TokenTransferTransaction1PostConditionsInner, bool)`

GetPostConditionsOk returns a tuple with the PostConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostConditions

`func (o *TransactionFoundResultAnyOf) SetPostConditions(v []TokenTransferTransaction1PostConditionsInner)`

SetPostConditions sets PostConditions field to given value.


### GetAnchorMode

`func (o *TransactionFoundResultAnyOf) GetAnchorMode() TokenTransferTransaction1AnchorMode`

GetAnchorMode returns the AnchorMode field if non-nil, zero value otherwise.

### GetAnchorModeOk

`func (o *TransactionFoundResultAnyOf) GetAnchorModeOk() (*TokenTransferTransaction1AnchorMode, bool)`

GetAnchorModeOk returns a tuple with the AnchorMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorMode

`func (o *TransactionFoundResultAnyOf) SetAnchorMode(v TokenTransferTransaction1AnchorMode)`

SetAnchorMode sets AnchorMode field to given value.


### GetBlockHash

`func (o *TransactionFoundResultAnyOf) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *TransactionFoundResultAnyOf) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *TransactionFoundResultAnyOf) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.


### GetBlockHeight

`func (o *TransactionFoundResultAnyOf) GetBlockHeight() int32`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *TransactionFoundResultAnyOf) GetBlockHeightOk() (*int32, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *TransactionFoundResultAnyOf) SetBlockHeight(v int32)`

SetBlockHeight sets BlockHeight field to given value.


### GetBlockTime

`func (o *TransactionFoundResultAnyOf) GetBlockTime() float32`

GetBlockTime returns the BlockTime field if non-nil, zero value otherwise.

### GetBlockTimeOk

`func (o *TransactionFoundResultAnyOf) GetBlockTimeOk() (*float32, bool)`

GetBlockTimeOk returns a tuple with the BlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTime

`func (o *TransactionFoundResultAnyOf) SetBlockTime(v float32)`

SetBlockTime sets BlockTime field to given value.


### GetBlockTimeIso

`func (o *TransactionFoundResultAnyOf) GetBlockTimeIso() string`

GetBlockTimeIso returns the BlockTimeIso field if non-nil, zero value otherwise.

### GetBlockTimeIsoOk

`func (o *TransactionFoundResultAnyOf) GetBlockTimeIsoOk() (*string, bool)`

GetBlockTimeIsoOk returns a tuple with the BlockTimeIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTimeIso

`func (o *TransactionFoundResultAnyOf) SetBlockTimeIso(v string)`

SetBlockTimeIso sets BlockTimeIso field to given value.


### GetBurnBlockTime

`func (o *TransactionFoundResultAnyOf) GetBurnBlockTime() int32`

GetBurnBlockTime returns the BurnBlockTime field if non-nil, zero value otherwise.

### GetBurnBlockTimeOk

`func (o *TransactionFoundResultAnyOf) GetBurnBlockTimeOk() (*int32, bool)`

GetBurnBlockTimeOk returns a tuple with the BurnBlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockTime

`func (o *TransactionFoundResultAnyOf) SetBurnBlockTime(v int32)`

SetBurnBlockTime sets BurnBlockTime field to given value.


### GetBurnBlockHeight

`func (o *TransactionFoundResultAnyOf) GetBurnBlockHeight() int32`

GetBurnBlockHeight returns the BurnBlockHeight field if non-nil, zero value otherwise.

### GetBurnBlockHeightOk

`func (o *TransactionFoundResultAnyOf) GetBurnBlockHeightOk() (*int32, bool)`

GetBurnBlockHeightOk returns a tuple with the BurnBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockHeight

`func (o *TransactionFoundResultAnyOf) SetBurnBlockHeight(v int32)`

SetBurnBlockHeight sets BurnBlockHeight field to given value.


### GetBurnBlockTimeIso

`func (o *TransactionFoundResultAnyOf) GetBurnBlockTimeIso() string`

GetBurnBlockTimeIso returns the BurnBlockTimeIso field if non-nil, zero value otherwise.

### GetBurnBlockTimeIsoOk

`func (o *TransactionFoundResultAnyOf) GetBurnBlockTimeIsoOk() (*string, bool)`

GetBurnBlockTimeIsoOk returns a tuple with the BurnBlockTimeIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockTimeIso

`func (o *TransactionFoundResultAnyOf) SetBurnBlockTimeIso(v string)`

SetBurnBlockTimeIso sets BurnBlockTimeIso field to given value.


### GetParentBurnBlockTime

`func (o *TransactionFoundResultAnyOf) GetParentBurnBlockTime() int32`

GetParentBurnBlockTime returns the ParentBurnBlockTime field if non-nil, zero value otherwise.

### GetParentBurnBlockTimeOk

`func (o *TransactionFoundResultAnyOf) GetParentBurnBlockTimeOk() (*int32, bool)`

GetParentBurnBlockTimeOk returns a tuple with the ParentBurnBlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBurnBlockTime

`func (o *TransactionFoundResultAnyOf) SetParentBurnBlockTime(v int32)`

SetParentBurnBlockTime sets ParentBurnBlockTime field to given value.


### GetParentBurnBlockTimeIso

`func (o *TransactionFoundResultAnyOf) GetParentBurnBlockTimeIso() string`

GetParentBurnBlockTimeIso returns the ParentBurnBlockTimeIso field if non-nil, zero value otherwise.

### GetParentBurnBlockTimeIsoOk

`func (o *TransactionFoundResultAnyOf) GetParentBurnBlockTimeIsoOk() (*string, bool)`

GetParentBurnBlockTimeIsoOk returns a tuple with the ParentBurnBlockTimeIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBurnBlockTimeIso

`func (o *TransactionFoundResultAnyOf) SetParentBurnBlockTimeIso(v string)`

SetParentBurnBlockTimeIso sets ParentBurnBlockTimeIso field to given value.


### GetCanonical

`func (o *TransactionFoundResultAnyOf) GetCanonical() bool`

GetCanonical returns the Canonical field if non-nil, zero value otherwise.

### GetCanonicalOk

`func (o *TransactionFoundResultAnyOf) GetCanonicalOk() (*bool, bool)`

GetCanonicalOk returns a tuple with the Canonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonical

`func (o *TransactionFoundResultAnyOf) SetCanonical(v bool)`

SetCanonical sets Canonical field to given value.


### GetTxIndex

`func (o *TransactionFoundResultAnyOf) GetTxIndex() int32`

GetTxIndex returns the TxIndex field if non-nil, zero value otherwise.

### GetTxIndexOk

`func (o *TransactionFoundResultAnyOf) GetTxIndexOk() (*int32, bool)`

GetTxIndexOk returns a tuple with the TxIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxIndex

`func (o *TransactionFoundResultAnyOf) SetTxIndex(v int32)`

SetTxIndex sets TxIndex field to given value.


### GetTxStatus

`func (o *TransactionFoundResultAnyOf) GetTxStatus() TokenTransferTransaction1TxStatus`

GetTxStatus returns the TxStatus field if non-nil, zero value otherwise.

### GetTxStatusOk

`func (o *TransactionFoundResultAnyOf) GetTxStatusOk() (*TokenTransferTransaction1TxStatus, bool)`

GetTxStatusOk returns a tuple with the TxStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxStatus

`func (o *TransactionFoundResultAnyOf) SetTxStatus(v TokenTransferTransaction1TxStatus)`

SetTxStatus sets TxStatus field to given value.


### GetTxResult

`func (o *TransactionFoundResultAnyOf) GetTxResult() TokenTransferTransactionTxResult`

GetTxResult returns the TxResult field if non-nil, zero value otherwise.

### GetTxResultOk

`func (o *TransactionFoundResultAnyOf) GetTxResultOk() (*TokenTransferTransactionTxResult, bool)`

GetTxResultOk returns a tuple with the TxResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxResult

`func (o *TransactionFoundResultAnyOf) SetTxResult(v TokenTransferTransactionTxResult)`

SetTxResult sets TxResult field to given value.


### GetEventCount

`func (o *TransactionFoundResultAnyOf) GetEventCount() int32`

GetEventCount returns the EventCount field if non-nil, zero value otherwise.

### GetEventCountOk

`func (o *TransactionFoundResultAnyOf) GetEventCountOk() (*int32, bool)`

GetEventCountOk returns a tuple with the EventCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCount

`func (o *TransactionFoundResultAnyOf) SetEventCount(v int32)`

SetEventCount sets EventCount field to given value.


### GetParentBlockHash

`func (o *TransactionFoundResultAnyOf) GetParentBlockHash() string`

GetParentBlockHash returns the ParentBlockHash field if non-nil, zero value otherwise.

### GetParentBlockHashOk

`func (o *TransactionFoundResultAnyOf) GetParentBlockHashOk() (*string, bool)`

GetParentBlockHashOk returns a tuple with the ParentBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBlockHash

`func (o *TransactionFoundResultAnyOf) SetParentBlockHash(v string)`

SetParentBlockHash sets ParentBlockHash field to given value.


### GetIsUnanchored

`func (o *TransactionFoundResultAnyOf) GetIsUnanchored() bool`

GetIsUnanchored returns the IsUnanchored field if non-nil, zero value otherwise.

### GetIsUnanchoredOk

`func (o *TransactionFoundResultAnyOf) GetIsUnanchoredOk() (*bool, bool)`

GetIsUnanchoredOk returns a tuple with the IsUnanchored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnanchored

`func (o *TransactionFoundResultAnyOf) SetIsUnanchored(v bool)`

SetIsUnanchored sets IsUnanchored field to given value.


### GetMicroblockHash

`func (o *TransactionFoundResultAnyOf) GetMicroblockHash() string`

GetMicroblockHash returns the MicroblockHash field if non-nil, zero value otherwise.

### GetMicroblockHashOk

`func (o *TransactionFoundResultAnyOf) GetMicroblockHashOk() (*string, bool)`

GetMicroblockHashOk returns a tuple with the MicroblockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroblockHash

`func (o *TransactionFoundResultAnyOf) SetMicroblockHash(v string)`

SetMicroblockHash sets MicroblockHash field to given value.


### GetMicroblockSequence

`func (o *TransactionFoundResultAnyOf) GetMicroblockSequence() int32`

GetMicroblockSequence returns the MicroblockSequence field if non-nil, zero value otherwise.

### GetMicroblockSequenceOk

`func (o *TransactionFoundResultAnyOf) GetMicroblockSequenceOk() (*int32, bool)`

GetMicroblockSequenceOk returns a tuple with the MicroblockSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroblockSequence

`func (o *TransactionFoundResultAnyOf) SetMicroblockSequence(v int32)`

SetMicroblockSequence sets MicroblockSequence field to given value.


### GetMicroblockCanonical

`func (o *TransactionFoundResultAnyOf) GetMicroblockCanonical() bool`

GetMicroblockCanonical returns the MicroblockCanonical field if non-nil, zero value otherwise.

### GetMicroblockCanonicalOk

`func (o *TransactionFoundResultAnyOf) GetMicroblockCanonicalOk() (*bool, bool)`

GetMicroblockCanonicalOk returns a tuple with the MicroblockCanonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroblockCanonical

`func (o *TransactionFoundResultAnyOf) SetMicroblockCanonical(v bool)`

SetMicroblockCanonical sets MicroblockCanonical field to given value.


### GetExecutionCostReadCount

`func (o *TransactionFoundResultAnyOf) GetExecutionCostReadCount() int32`

GetExecutionCostReadCount returns the ExecutionCostReadCount field if non-nil, zero value otherwise.

### GetExecutionCostReadCountOk

`func (o *TransactionFoundResultAnyOf) GetExecutionCostReadCountOk() (*int32, bool)`

GetExecutionCostReadCountOk returns a tuple with the ExecutionCostReadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCostReadCount

`func (o *TransactionFoundResultAnyOf) SetExecutionCostReadCount(v int32)`

SetExecutionCostReadCount sets ExecutionCostReadCount field to given value.


### GetExecutionCostReadLength

`func (o *TransactionFoundResultAnyOf) GetExecutionCostReadLength() int32`

GetExecutionCostReadLength returns the ExecutionCostReadLength field if non-nil, zero value otherwise.

### GetExecutionCostReadLengthOk

`func (o *TransactionFoundResultAnyOf) GetExecutionCostReadLengthOk() (*int32, bool)`

GetExecutionCostReadLengthOk returns a tuple with the ExecutionCostReadLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCostReadLength

`func (o *TransactionFoundResultAnyOf) SetExecutionCostReadLength(v int32)`

SetExecutionCostReadLength sets ExecutionCostReadLength field to given value.


### GetExecutionCostRuntime

`func (o *TransactionFoundResultAnyOf) GetExecutionCostRuntime() int32`

GetExecutionCostRuntime returns the ExecutionCostRuntime field if non-nil, zero value otherwise.

### GetExecutionCostRuntimeOk

`func (o *TransactionFoundResultAnyOf) GetExecutionCostRuntimeOk() (*int32, bool)`

GetExecutionCostRuntimeOk returns a tuple with the ExecutionCostRuntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCostRuntime

`func (o *TransactionFoundResultAnyOf) SetExecutionCostRuntime(v int32)`

SetExecutionCostRuntime sets ExecutionCostRuntime field to given value.


### GetExecutionCostWriteCount

`func (o *TransactionFoundResultAnyOf) GetExecutionCostWriteCount() int32`

GetExecutionCostWriteCount returns the ExecutionCostWriteCount field if non-nil, zero value otherwise.

### GetExecutionCostWriteCountOk

`func (o *TransactionFoundResultAnyOf) GetExecutionCostWriteCountOk() (*int32, bool)`

GetExecutionCostWriteCountOk returns a tuple with the ExecutionCostWriteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCostWriteCount

`func (o *TransactionFoundResultAnyOf) SetExecutionCostWriteCount(v int32)`

SetExecutionCostWriteCount sets ExecutionCostWriteCount field to given value.


### GetExecutionCostWriteLength

`func (o *TransactionFoundResultAnyOf) GetExecutionCostWriteLength() int32`

GetExecutionCostWriteLength returns the ExecutionCostWriteLength field if non-nil, zero value otherwise.

### GetExecutionCostWriteLengthOk

`func (o *TransactionFoundResultAnyOf) GetExecutionCostWriteLengthOk() (*int32, bool)`

GetExecutionCostWriteLengthOk returns a tuple with the ExecutionCostWriteLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCostWriteLength

`func (o *TransactionFoundResultAnyOf) SetExecutionCostWriteLength(v int32)`

SetExecutionCostWriteLength sets ExecutionCostWriteLength field to given value.


### GetEvents

`func (o *TransactionFoundResultAnyOf) GetEvents() []TokenTransferTransaction1EventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *TransactionFoundResultAnyOf) GetEventsOk() (*[]TokenTransferTransaction1EventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *TransactionFoundResultAnyOf) SetEvents(v []TokenTransferTransaction1EventsInner)`

SetEvents sets Events field to given value.


### GetTxType

`func (o *TransactionFoundResultAnyOf) GetTxType() string`

GetTxType returns the TxType field if non-nil, zero value otherwise.

### GetTxTypeOk

`func (o *TransactionFoundResultAnyOf) GetTxTypeOk() (*string, bool)`

GetTxTypeOk returns a tuple with the TxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxType

`func (o *TransactionFoundResultAnyOf) SetTxType(v string)`

SetTxType sets TxType field to given value.


### GetTokenTransfer

`func (o *TransactionFoundResultAnyOf) GetTokenTransfer() TokenTransferTransactionTokenTransfer`

GetTokenTransfer returns the TokenTransfer field if non-nil, zero value otherwise.

### GetTokenTransferOk

`func (o *TransactionFoundResultAnyOf) GetTokenTransferOk() (*TokenTransferTransactionTokenTransfer, bool)`

GetTokenTransferOk returns a tuple with the TokenTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTransfer

`func (o *TransactionFoundResultAnyOf) SetTokenTransfer(v TokenTransferTransactionTokenTransfer)`

SetTokenTransfer sets TokenTransfer field to given value.


### GetSmartContract

`func (o *TransactionFoundResultAnyOf) GetSmartContract() SmartContractTransactionSmartContract`

GetSmartContract returns the SmartContract field if non-nil, zero value otherwise.

### GetSmartContractOk

`func (o *TransactionFoundResultAnyOf) GetSmartContractOk() (*SmartContractTransactionSmartContract, bool)`

GetSmartContractOk returns a tuple with the SmartContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartContract

`func (o *TransactionFoundResultAnyOf) SetSmartContract(v SmartContractTransactionSmartContract)`

SetSmartContract sets SmartContract field to given value.


### GetContractCall

`func (o *TransactionFoundResultAnyOf) GetContractCall() ContractCallTransactionContractCall`

GetContractCall returns the ContractCall field if non-nil, zero value otherwise.

### GetContractCallOk

`func (o *TransactionFoundResultAnyOf) GetContractCallOk() (*ContractCallTransactionContractCall, bool)`

GetContractCallOk returns a tuple with the ContractCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractCall

`func (o *TransactionFoundResultAnyOf) SetContractCall(v ContractCallTransactionContractCall)`

SetContractCall sets ContractCall field to given value.


### GetPoisonMicroblock

`func (o *TransactionFoundResultAnyOf) GetPoisonMicroblock() PoisonMicroblockTransactionPoisonMicroblock`

GetPoisonMicroblock returns the PoisonMicroblock field if non-nil, zero value otherwise.

### GetPoisonMicroblockOk

`func (o *TransactionFoundResultAnyOf) GetPoisonMicroblockOk() (*PoisonMicroblockTransactionPoisonMicroblock, bool)`

GetPoisonMicroblockOk returns a tuple with the PoisonMicroblock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoisonMicroblock

`func (o *TransactionFoundResultAnyOf) SetPoisonMicroblock(v PoisonMicroblockTransactionPoisonMicroblock)`

SetPoisonMicroblock sets PoisonMicroblock field to given value.


### GetCoinbasePayload

`func (o *TransactionFoundResultAnyOf) GetCoinbasePayload() CoinbaseTransactionCoinbasePayload`

GetCoinbasePayload returns the CoinbasePayload field if non-nil, zero value otherwise.

### GetCoinbasePayloadOk

`func (o *TransactionFoundResultAnyOf) GetCoinbasePayloadOk() (*CoinbaseTransactionCoinbasePayload, bool)`

GetCoinbasePayloadOk returns a tuple with the CoinbasePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinbasePayload

`func (o *TransactionFoundResultAnyOf) SetCoinbasePayload(v CoinbaseTransactionCoinbasePayload)`

SetCoinbasePayload sets CoinbasePayload field to given value.


### GetTenureChangePayload

`func (o *TransactionFoundResultAnyOf) GetTenureChangePayload() TenureChangeTransaction1TenureChangePayload`

GetTenureChangePayload returns the TenureChangePayload field if non-nil, zero value otherwise.

### GetTenureChangePayloadOk

`func (o *TransactionFoundResultAnyOf) GetTenureChangePayloadOk() (*TenureChangeTransaction1TenureChangePayload, bool)`

GetTenureChangePayloadOk returns a tuple with the TenureChangePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenureChangePayload

`func (o *TransactionFoundResultAnyOf) SetTenureChangePayload(v TenureChangeTransaction1TenureChangePayload)`

SetTenureChangePayload sets TenureChangePayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


