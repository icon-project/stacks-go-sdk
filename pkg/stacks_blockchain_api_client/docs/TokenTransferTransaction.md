# TokenTransferTransaction

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
**PostConditionMode** | [**TokenTransferTransactionPostConditionMode**](TokenTransferTransactionPostConditionMode.md) |  | 
**PostConditions** | [**[]TokenTransferTransactionPostConditionsInner**](TokenTransferTransactionPostConditionsInner.md) |  | 
**AnchorMode** | [**TokenTransferTransactionAnchorMode**](TokenTransferTransactionAnchorMode.md) |  | 
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
**TxStatus** | [**TokenTransferTransactionTxStatus**](TokenTransferTransactionTxStatus.md) |  | 
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
**Events** | [**[]TokenTransferTransactionEventsInner**](TokenTransferTransactionEventsInner.md) |  | 
**TxType** | **string** |  | 
**TokenTransfer** | [**TokenTransferTransactionTokenTransfer**](TokenTransferTransactionTokenTransfer.md) |  | 

## Methods

### NewTokenTransferTransaction

`func NewTokenTransferTransaction(txId string, nonce int32, feeRate string, senderAddress string, sponsored bool, postConditionMode TokenTransferTransactionPostConditionMode, postConditions []TokenTransferTransactionPostConditionsInner, anchorMode TokenTransferTransactionAnchorMode, blockHash string, blockHeight int32, blockTime float32, blockTimeIso string, burnBlockTime int32, burnBlockHeight int32, burnBlockTimeIso string, parentBurnBlockTime int32, parentBurnBlockTimeIso string, canonical bool, txIndex int32, txStatus TokenTransferTransactionTxStatus, txResult TokenTransferTransactionTxResult, eventCount int32, parentBlockHash string, isUnanchored bool, microblockHash string, microblockSequence int32, microblockCanonical bool, executionCostReadCount int32, executionCostReadLength int32, executionCostRuntime int32, executionCostWriteCount int32, executionCostWriteLength int32, events []TokenTransferTransactionEventsInner, txType string, tokenTransfer TokenTransferTransactionTokenTransfer, ) *TokenTransferTransaction`

NewTokenTransferTransaction instantiates a new TokenTransferTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenTransferTransactionWithDefaults

`func NewTokenTransferTransactionWithDefaults() *TokenTransferTransaction`

NewTokenTransferTransactionWithDefaults instantiates a new TokenTransferTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxId

`func (o *TokenTransferTransaction) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *TokenTransferTransaction) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *TokenTransferTransaction) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetNonce

`func (o *TokenTransferTransaction) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *TokenTransferTransaction) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *TokenTransferTransaction) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetFeeRate

`func (o *TokenTransferTransaction) GetFeeRate() string`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *TokenTransferTransaction) GetFeeRateOk() (*string, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *TokenTransferTransaction) SetFeeRate(v string)`

SetFeeRate sets FeeRate field to given value.


### GetSenderAddress

`func (o *TokenTransferTransaction) GetSenderAddress() string`

GetSenderAddress returns the SenderAddress field if non-nil, zero value otherwise.

### GetSenderAddressOk

`func (o *TokenTransferTransaction) GetSenderAddressOk() (*string, bool)`

GetSenderAddressOk returns a tuple with the SenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAddress

`func (o *TokenTransferTransaction) SetSenderAddress(v string)`

SetSenderAddress sets SenderAddress field to given value.


### GetSponsorNonce

`func (o *TokenTransferTransaction) GetSponsorNonce() int32`

GetSponsorNonce returns the SponsorNonce field if non-nil, zero value otherwise.

### GetSponsorNonceOk

`func (o *TokenTransferTransaction) GetSponsorNonceOk() (*int32, bool)`

GetSponsorNonceOk returns a tuple with the SponsorNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorNonce

`func (o *TokenTransferTransaction) SetSponsorNonce(v int32)`

SetSponsorNonce sets SponsorNonce field to given value.

### HasSponsorNonce

`func (o *TokenTransferTransaction) HasSponsorNonce() bool`

HasSponsorNonce returns a boolean if a field has been set.

### GetSponsored

`func (o *TokenTransferTransaction) GetSponsored() bool`

GetSponsored returns the Sponsored field if non-nil, zero value otherwise.

### GetSponsoredOk

`func (o *TokenTransferTransaction) GetSponsoredOk() (*bool, bool)`

GetSponsoredOk returns a tuple with the Sponsored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsored

`func (o *TokenTransferTransaction) SetSponsored(v bool)`

SetSponsored sets Sponsored field to given value.


### GetSponsorAddress

`func (o *TokenTransferTransaction) GetSponsorAddress() string`

GetSponsorAddress returns the SponsorAddress field if non-nil, zero value otherwise.

### GetSponsorAddressOk

`func (o *TokenTransferTransaction) GetSponsorAddressOk() (*string, bool)`

GetSponsorAddressOk returns a tuple with the SponsorAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorAddress

`func (o *TokenTransferTransaction) SetSponsorAddress(v string)`

SetSponsorAddress sets SponsorAddress field to given value.

### HasSponsorAddress

`func (o *TokenTransferTransaction) HasSponsorAddress() bool`

HasSponsorAddress returns a boolean if a field has been set.

### GetPostConditionMode

`func (o *TokenTransferTransaction) GetPostConditionMode() TokenTransferTransactionPostConditionMode`

GetPostConditionMode returns the PostConditionMode field if non-nil, zero value otherwise.

### GetPostConditionModeOk

`func (o *TokenTransferTransaction) GetPostConditionModeOk() (*TokenTransferTransactionPostConditionMode, bool)`

GetPostConditionModeOk returns a tuple with the PostConditionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostConditionMode

`func (o *TokenTransferTransaction) SetPostConditionMode(v TokenTransferTransactionPostConditionMode)`

SetPostConditionMode sets PostConditionMode field to given value.


### GetPostConditions

`func (o *TokenTransferTransaction) GetPostConditions() []TokenTransferTransactionPostConditionsInner`

GetPostConditions returns the PostConditions field if non-nil, zero value otherwise.

### GetPostConditionsOk

`func (o *TokenTransferTransaction) GetPostConditionsOk() (*[]TokenTransferTransactionPostConditionsInner, bool)`

GetPostConditionsOk returns a tuple with the PostConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostConditions

`func (o *TokenTransferTransaction) SetPostConditions(v []TokenTransferTransactionPostConditionsInner)`

SetPostConditions sets PostConditions field to given value.


### GetAnchorMode

`func (o *TokenTransferTransaction) GetAnchorMode() TokenTransferTransactionAnchorMode`

GetAnchorMode returns the AnchorMode field if non-nil, zero value otherwise.

### GetAnchorModeOk

`func (o *TokenTransferTransaction) GetAnchorModeOk() (*TokenTransferTransactionAnchorMode, bool)`

GetAnchorModeOk returns a tuple with the AnchorMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorMode

`func (o *TokenTransferTransaction) SetAnchorMode(v TokenTransferTransactionAnchorMode)`

SetAnchorMode sets AnchorMode field to given value.


### GetBlockHash

`func (o *TokenTransferTransaction) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *TokenTransferTransaction) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *TokenTransferTransaction) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.


### GetBlockHeight

`func (o *TokenTransferTransaction) GetBlockHeight() int32`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *TokenTransferTransaction) GetBlockHeightOk() (*int32, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *TokenTransferTransaction) SetBlockHeight(v int32)`

SetBlockHeight sets BlockHeight field to given value.


### GetBlockTime

`func (o *TokenTransferTransaction) GetBlockTime() float32`

GetBlockTime returns the BlockTime field if non-nil, zero value otherwise.

### GetBlockTimeOk

`func (o *TokenTransferTransaction) GetBlockTimeOk() (*float32, bool)`

GetBlockTimeOk returns a tuple with the BlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTime

`func (o *TokenTransferTransaction) SetBlockTime(v float32)`

SetBlockTime sets BlockTime field to given value.


### GetBlockTimeIso

`func (o *TokenTransferTransaction) GetBlockTimeIso() string`

GetBlockTimeIso returns the BlockTimeIso field if non-nil, zero value otherwise.

### GetBlockTimeIsoOk

`func (o *TokenTransferTransaction) GetBlockTimeIsoOk() (*string, bool)`

GetBlockTimeIsoOk returns a tuple with the BlockTimeIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTimeIso

`func (o *TokenTransferTransaction) SetBlockTimeIso(v string)`

SetBlockTimeIso sets BlockTimeIso field to given value.


### GetBurnBlockTime

`func (o *TokenTransferTransaction) GetBurnBlockTime() int32`

GetBurnBlockTime returns the BurnBlockTime field if non-nil, zero value otherwise.

### GetBurnBlockTimeOk

`func (o *TokenTransferTransaction) GetBurnBlockTimeOk() (*int32, bool)`

GetBurnBlockTimeOk returns a tuple with the BurnBlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockTime

`func (o *TokenTransferTransaction) SetBurnBlockTime(v int32)`

SetBurnBlockTime sets BurnBlockTime field to given value.


### GetBurnBlockHeight

`func (o *TokenTransferTransaction) GetBurnBlockHeight() int32`

GetBurnBlockHeight returns the BurnBlockHeight field if non-nil, zero value otherwise.

### GetBurnBlockHeightOk

`func (o *TokenTransferTransaction) GetBurnBlockHeightOk() (*int32, bool)`

GetBurnBlockHeightOk returns a tuple with the BurnBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockHeight

`func (o *TokenTransferTransaction) SetBurnBlockHeight(v int32)`

SetBurnBlockHeight sets BurnBlockHeight field to given value.


### GetBurnBlockTimeIso

`func (o *TokenTransferTransaction) GetBurnBlockTimeIso() string`

GetBurnBlockTimeIso returns the BurnBlockTimeIso field if non-nil, zero value otherwise.

### GetBurnBlockTimeIsoOk

`func (o *TokenTransferTransaction) GetBurnBlockTimeIsoOk() (*string, bool)`

GetBurnBlockTimeIsoOk returns a tuple with the BurnBlockTimeIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockTimeIso

`func (o *TokenTransferTransaction) SetBurnBlockTimeIso(v string)`

SetBurnBlockTimeIso sets BurnBlockTimeIso field to given value.


### GetParentBurnBlockTime

`func (o *TokenTransferTransaction) GetParentBurnBlockTime() int32`

GetParentBurnBlockTime returns the ParentBurnBlockTime field if non-nil, zero value otherwise.

### GetParentBurnBlockTimeOk

`func (o *TokenTransferTransaction) GetParentBurnBlockTimeOk() (*int32, bool)`

GetParentBurnBlockTimeOk returns a tuple with the ParentBurnBlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBurnBlockTime

`func (o *TokenTransferTransaction) SetParentBurnBlockTime(v int32)`

SetParentBurnBlockTime sets ParentBurnBlockTime field to given value.


### GetParentBurnBlockTimeIso

`func (o *TokenTransferTransaction) GetParentBurnBlockTimeIso() string`

GetParentBurnBlockTimeIso returns the ParentBurnBlockTimeIso field if non-nil, zero value otherwise.

### GetParentBurnBlockTimeIsoOk

`func (o *TokenTransferTransaction) GetParentBurnBlockTimeIsoOk() (*string, bool)`

GetParentBurnBlockTimeIsoOk returns a tuple with the ParentBurnBlockTimeIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBurnBlockTimeIso

`func (o *TokenTransferTransaction) SetParentBurnBlockTimeIso(v string)`

SetParentBurnBlockTimeIso sets ParentBurnBlockTimeIso field to given value.


### GetCanonical

`func (o *TokenTransferTransaction) GetCanonical() bool`

GetCanonical returns the Canonical field if non-nil, zero value otherwise.

### GetCanonicalOk

`func (o *TokenTransferTransaction) GetCanonicalOk() (*bool, bool)`

GetCanonicalOk returns a tuple with the Canonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonical

`func (o *TokenTransferTransaction) SetCanonical(v bool)`

SetCanonical sets Canonical field to given value.


### GetTxIndex

`func (o *TokenTransferTransaction) GetTxIndex() int32`

GetTxIndex returns the TxIndex field if non-nil, zero value otherwise.

### GetTxIndexOk

`func (o *TokenTransferTransaction) GetTxIndexOk() (*int32, bool)`

GetTxIndexOk returns a tuple with the TxIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxIndex

`func (o *TokenTransferTransaction) SetTxIndex(v int32)`

SetTxIndex sets TxIndex field to given value.


### GetTxStatus

`func (o *TokenTransferTransaction) GetTxStatus() TokenTransferTransactionTxStatus`

GetTxStatus returns the TxStatus field if non-nil, zero value otherwise.

### GetTxStatusOk

`func (o *TokenTransferTransaction) GetTxStatusOk() (*TokenTransferTransactionTxStatus, bool)`

GetTxStatusOk returns a tuple with the TxStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxStatus

`func (o *TokenTransferTransaction) SetTxStatus(v TokenTransferTransactionTxStatus)`

SetTxStatus sets TxStatus field to given value.


### GetTxResult

`func (o *TokenTransferTransaction) GetTxResult() TokenTransferTransactionTxResult`

GetTxResult returns the TxResult field if non-nil, zero value otherwise.

### GetTxResultOk

`func (o *TokenTransferTransaction) GetTxResultOk() (*TokenTransferTransactionTxResult, bool)`

GetTxResultOk returns a tuple with the TxResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxResult

`func (o *TokenTransferTransaction) SetTxResult(v TokenTransferTransactionTxResult)`

SetTxResult sets TxResult field to given value.


### GetEventCount

`func (o *TokenTransferTransaction) GetEventCount() int32`

GetEventCount returns the EventCount field if non-nil, zero value otherwise.

### GetEventCountOk

`func (o *TokenTransferTransaction) GetEventCountOk() (*int32, bool)`

GetEventCountOk returns a tuple with the EventCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCount

`func (o *TokenTransferTransaction) SetEventCount(v int32)`

SetEventCount sets EventCount field to given value.


### GetParentBlockHash

`func (o *TokenTransferTransaction) GetParentBlockHash() string`

GetParentBlockHash returns the ParentBlockHash field if non-nil, zero value otherwise.

### GetParentBlockHashOk

`func (o *TokenTransferTransaction) GetParentBlockHashOk() (*string, bool)`

GetParentBlockHashOk returns a tuple with the ParentBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBlockHash

`func (o *TokenTransferTransaction) SetParentBlockHash(v string)`

SetParentBlockHash sets ParentBlockHash field to given value.


### GetIsUnanchored

`func (o *TokenTransferTransaction) GetIsUnanchored() bool`

GetIsUnanchored returns the IsUnanchored field if non-nil, zero value otherwise.

### GetIsUnanchoredOk

`func (o *TokenTransferTransaction) GetIsUnanchoredOk() (*bool, bool)`

GetIsUnanchoredOk returns a tuple with the IsUnanchored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnanchored

`func (o *TokenTransferTransaction) SetIsUnanchored(v bool)`

SetIsUnanchored sets IsUnanchored field to given value.


### GetMicroblockHash

`func (o *TokenTransferTransaction) GetMicroblockHash() string`

GetMicroblockHash returns the MicroblockHash field if non-nil, zero value otherwise.

### GetMicroblockHashOk

`func (o *TokenTransferTransaction) GetMicroblockHashOk() (*string, bool)`

GetMicroblockHashOk returns a tuple with the MicroblockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroblockHash

`func (o *TokenTransferTransaction) SetMicroblockHash(v string)`

SetMicroblockHash sets MicroblockHash field to given value.


### GetMicroblockSequence

`func (o *TokenTransferTransaction) GetMicroblockSequence() int32`

GetMicroblockSequence returns the MicroblockSequence field if non-nil, zero value otherwise.

### GetMicroblockSequenceOk

`func (o *TokenTransferTransaction) GetMicroblockSequenceOk() (*int32, bool)`

GetMicroblockSequenceOk returns a tuple with the MicroblockSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroblockSequence

`func (o *TokenTransferTransaction) SetMicroblockSequence(v int32)`

SetMicroblockSequence sets MicroblockSequence field to given value.


### GetMicroblockCanonical

`func (o *TokenTransferTransaction) GetMicroblockCanonical() bool`

GetMicroblockCanonical returns the MicroblockCanonical field if non-nil, zero value otherwise.

### GetMicroblockCanonicalOk

`func (o *TokenTransferTransaction) GetMicroblockCanonicalOk() (*bool, bool)`

GetMicroblockCanonicalOk returns a tuple with the MicroblockCanonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroblockCanonical

`func (o *TokenTransferTransaction) SetMicroblockCanonical(v bool)`

SetMicroblockCanonical sets MicroblockCanonical field to given value.


### GetExecutionCostReadCount

`func (o *TokenTransferTransaction) GetExecutionCostReadCount() int32`

GetExecutionCostReadCount returns the ExecutionCostReadCount field if non-nil, zero value otherwise.

### GetExecutionCostReadCountOk

`func (o *TokenTransferTransaction) GetExecutionCostReadCountOk() (*int32, bool)`

GetExecutionCostReadCountOk returns a tuple with the ExecutionCostReadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCostReadCount

`func (o *TokenTransferTransaction) SetExecutionCostReadCount(v int32)`

SetExecutionCostReadCount sets ExecutionCostReadCount field to given value.


### GetExecutionCostReadLength

`func (o *TokenTransferTransaction) GetExecutionCostReadLength() int32`

GetExecutionCostReadLength returns the ExecutionCostReadLength field if non-nil, zero value otherwise.

### GetExecutionCostReadLengthOk

`func (o *TokenTransferTransaction) GetExecutionCostReadLengthOk() (*int32, bool)`

GetExecutionCostReadLengthOk returns a tuple with the ExecutionCostReadLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCostReadLength

`func (o *TokenTransferTransaction) SetExecutionCostReadLength(v int32)`

SetExecutionCostReadLength sets ExecutionCostReadLength field to given value.


### GetExecutionCostRuntime

`func (o *TokenTransferTransaction) GetExecutionCostRuntime() int32`

GetExecutionCostRuntime returns the ExecutionCostRuntime field if non-nil, zero value otherwise.

### GetExecutionCostRuntimeOk

`func (o *TokenTransferTransaction) GetExecutionCostRuntimeOk() (*int32, bool)`

GetExecutionCostRuntimeOk returns a tuple with the ExecutionCostRuntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCostRuntime

`func (o *TokenTransferTransaction) SetExecutionCostRuntime(v int32)`

SetExecutionCostRuntime sets ExecutionCostRuntime field to given value.


### GetExecutionCostWriteCount

`func (o *TokenTransferTransaction) GetExecutionCostWriteCount() int32`

GetExecutionCostWriteCount returns the ExecutionCostWriteCount field if non-nil, zero value otherwise.

### GetExecutionCostWriteCountOk

`func (o *TokenTransferTransaction) GetExecutionCostWriteCountOk() (*int32, bool)`

GetExecutionCostWriteCountOk returns a tuple with the ExecutionCostWriteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCostWriteCount

`func (o *TokenTransferTransaction) SetExecutionCostWriteCount(v int32)`

SetExecutionCostWriteCount sets ExecutionCostWriteCount field to given value.


### GetExecutionCostWriteLength

`func (o *TokenTransferTransaction) GetExecutionCostWriteLength() int32`

GetExecutionCostWriteLength returns the ExecutionCostWriteLength field if non-nil, zero value otherwise.

### GetExecutionCostWriteLengthOk

`func (o *TokenTransferTransaction) GetExecutionCostWriteLengthOk() (*int32, bool)`

GetExecutionCostWriteLengthOk returns a tuple with the ExecutionCostWriteLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCostWriteLength

`func (o *TokenTransferTransaction) SetExecutionCostWriteLength(v int32)`

SetExecutionCostWriteLength sets ExecutionCostWriteLength field to given value.


### GetEvents

`func (o *TokenTransferTransaction) GetEvents() []TokenTransferTransactionEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *TokenTransferTransaction) GetEventsOk() (*[]TokenTransferTransactionEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *TokenTransferTransaction) SetEvents(v []TokenTransferTransactionEventsInner)`

SetEvents sets Events field to given value.


### GetTxType

`func (o *TokenTransferTransaction) GetTxType() string`

GetTxType returns the TxType field if non-nil, zero value otherwise.

### GetTxTypeOk

`func (o *TokenTransferTransaction) GetTxTypeOk() (*string, bool)`

GetTxTypeOk returns a tuple with the TxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxType

`func (o *TokenTransferTransaction) SetTxType(v string)`

SetTxType sets TxType field to given value.


### GetTokenTransfer

`func (o *TokenTransferTransaction) GetTokenTransfer() TokenTransferTransactionTokenTransfer`

GetTokenTransfer returns the TokenTransfer field if non-nil, zero value otherwise.

### GetTokenTransferOk

`func (o *TokenTransferTransaction) GetTokenTransferOk() (*TokenTransferTransactionTokenTransfer, bool)`

GetTokenTransferOk returns a tuple with the TokenTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTransfer

`func (o *TokenTransferTransaction) SetTokenTransfer(v TokenTransferTransactionTokenTransfer)`

SetTokenTransfer sets TokenTransfer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


