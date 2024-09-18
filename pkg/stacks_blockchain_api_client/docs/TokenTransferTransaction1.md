# TokenTransferTransaction1

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

## Methods

### NewTokenTransferTransaction1

`func NewTokenTransferTransaction1(txId string, nonce int32, feeRate string, senderAddress string, sponsored bool, postConditionMode TokenTransferTransaction1PostConditionMode, postConditions []TokenTransferTransaction1PostConditionsInner, anchorMode TokenTransferTransaction1AnchorMode, blockHash string, blockHeight int32, blockTime float32, blockTimeIso string, burnBlockTime int32, burnBlockHeight int32, burnBlockTimeIso string, parentBurnBlockTime int32, parentBurnBlockTimeIso string, canonical bool, txIndex int32, txStatus TokenTransferTransaction1TxStatus, txResult TokenTransferTransactionTxResult, eventCount int32, parentBlockHash string, isUnanchored bool, microblockHash string, microblockSequence int32, microblockCanonical bool, executionCostReadCount int32, executionCostReadLength int32, executionCostRuntime int32, executionCostWriteCount int32, executionCostWriteLength int32, events []TokenTransferTransaction1EventsInner, txType string, tokenTransfer TokenTransferTransactionTokenTransfer, ) *TokenTransferTransaction1`

NewTokenTransferTransaction1 instantiates a new TokenTransferTransaction1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenTransferTransaction1WithDefaults

`func NewTokenTransferTransaction1WithDefaults() *TokenTransferTransaction1`

NewTokenTransferTransaction1WithDefaults instantiates a new TokenTransferTransaction1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxId

`func (o *TokenTransferTransaction1) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *TokenTransferTransaction1) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *TokenTransferTransaction1) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetNonce

`func (o *TokenTransferTransaction1) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *TokenTransferTransaction1) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *TokenTransferTransaction1) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetFeeRate

`func (o *TokenTransferTransaction1) GetFeeRate() string`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *TokenTransferTransaction1) GetFeeRateOk() (*string, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *TokenTransferTransaction1) SetFeeRate(v string)`

SetFeeRate sets FeeRate field to given value.


### GetSenderAddress

`func (o *TokenTransferTransaction1) GetSenderAddress() string`

GetSenderAddress returns the SenderAddress field if non-nil, zero value otherwise.

### GetSenderAddressOk

`func (o *TokenTransferTransaction1) GetSenderAddressOk() (*string, bool)`

GetSenderAddressOk returns a tuple with the SenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAddress

`func (o *TokenTransferTransaction1) SetSenderAddress(v string)`

SetSenderAddress sets SenderAddress field to given value.


### GetSponsorNonce

`func (o *TokenTransferTransaction1) GetSponsorNonce() int32`

GetSponsorNonce returns the SponsorNonce field if non-nil, zero value otherwise.

### GetSponsorNonceOk

`func (o *TokenTransferTransaction1) GetSponsorNonceOk() (*int32, bool)`

GetSponsorNonceOk returns a tuple with the SponsorNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorNonce

`func (o *TokenTransferTransaction1) SetSponsorNonce(v int32)`

SetSponsorNonce sets SponsorNonce field to given value.

### HasSponsorNonce

`func (o *TokenTransferTransaction1) HasSponsorNonce() bool`

HasSponsorNonce returns a boolean if a field has been set.

### GetSponsored

`func (o *TokenTransferTransaction1) GetSponsored() bool`

GetSponsored returns the Sponsored field if non-nil, zero value otherwise.

### GetSponsoredOk

`func (o *TokenTransferTransaction1) GetSponsoredOk() (*bool, bool)`

GetSponsoredOk returns a tuple with the Sponsored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsored

`func (o *TokenTransferTransaction1) SetSponsored(v bool)`

SetSponsored sets Sponsored field to given value.


### GetSponsorAddress

`func (o *TokenTransferTransaction1) GetSponsorAddress() string`

GetSponsorAddress returns the SponsorAddress field if non-nil, zero value otherwise.

### GetSponsorAddressOk

`func (o *TokenTransferTransaction1) GetSponsorAddressOk() (*string, bool)`

GetSponsorAddressOk returns a tuple with the SponsorAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorAddress

`func (o *TokenTransferTransaction1) SetSponsorAddress(v string)`

SetSponsorAddress sets SponsorAddress field to given value.

### HasSponsorAddress

`func (o *TokenTransferTransaction1) HasSponsorAddress() bool`

HasSponsorAddress returns a boolean if a field has been set.

### GetPostConditionMode

`func (o *TokenTransferTransaction1) GetPostConditionMode() TokenTransferTransaction1PostConditionMode`

GetPostConditionMode returns the PostConditionMode field if non-nil, zero value otherwise.

### GetPostConditionModeOk

`func (o *TokenTransferTransaction1) GetPostConditionModeOk() (*TokenTransferTransaction1PostConditionMode, bool)`

GetPostConditionModeOk returns a tuple with the PostConditionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostConditionMode

`func (o *TokenTransferTransaction1) SetPostConditionMode(v TokenTransferTransaction1PostConditionMode)`

SetPostConditionMode sets PostConditionMode field to given value.


### GetPostConditions

`func (o *TokenTransferTransaction1) GetPostConditions() []TokenTransferTransaction1PostConditionsInner`

GetPostConditions returns the PostConditions field if non-nil, zero value otherwise.

### GetPostConditionsOk

`func (o *TokenTransferTransaction1) GetPostConditionsOk() (*[]TokenTransferTransaction1PostConditionsInner, bool)`

GetPostConditionsOk returns a tuple with the PostConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostConditions

`func (o *TokenTransferTransaction1) SetPostConditions(v []TokenTransferTransaction1PostConditionsInner)`

SetPostConditions sets PostConditions field to given value.


### GetAnchorMode

`func (o *TokenTransferTransaction1) GetAnchorMode() TokenTransferTransaction1AnchorMode`

GetAnchorMode returns the AnchorMode field if non-nil, zero value otherwise.

### GetAnchorModeOk

`func (o *TokenTransferTransaction1) GetAnchorModeOk() (*TokenTransferTransaction1AnchorMode, bool)`

GetAnchorModeOk returns a tuple with the AnchorMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorMode

`func (o *TokenTransferTransaction1) SetAnchorMode(v TokenTransferTransaction1AnchorMode)`

SetAnchorMode sets AnchorMode field to given value.


### GetBlockHash

`func (o *TokenTransferTransaction1) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *TokenTransferTransaction1) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *TokenTransferTransaction1) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.


### GetBlockHeight

`func (o *TokenTransferTransaction1) GetBlockHeight() int32`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *TokenTransferTransaction1) GetBlockHeightOk() (*int32, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *TokenTransferTransaction1) SetBlockHeight(v int32)`

SetBlockHeight sets BlockHeight field to given value.


### GetBlockTime

`func (o *TokenTransferTransaction1) GetBlockTime() float32`

GetBlockTime returns the BlockTime field if non-nil, zero value otherwise.

### GetBlockTimeOk

`func (o *TokenTransferTransaction1) GetBlockTimeOk() (*float32, bool)`

GetBlockTimeOk returns a tuple with the BlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTime

`func (o *TokenTransferTransaction1) SetBlockTime(v float32)`

SetBlockTime sets BlockTime field to given value.


### GetBlockTimeIso

`func (o *TokenTransferTransaction1) GetBlockTimeIso() string`

GetBlockTimeIso returns the BlockTimeIso field if non-nil, zero value otherwise.

### GetBlockTimeIsoOk

`func (o *TokenTransferTransaction1) GetBlockTimeIsoOk() (*string, bool)`

GetBlockTimeIsoOk returns a tuple with the BlockTimeIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTimeIso

`func (o *TokenTransferTransaction1) SetBlockTimeIso(v string)`

SetBlockTimeIso sets BlockTimeIso field to given value.


### GetBurnBlockTime

`func (o *TokenTransferTransaction1) GetBurnBlockTime() int32`

GetBurnBlockTime returns the BurnBlockTime field if non-nil, zero value otherwise.

### GetBurnBlockTimeOk

`func (o *TokenTransferTransaction1) GetBurnBlockTimeOk() (*int32, bool)`

GetBurnBlockTimeOk returns a tuple with the BurnBlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockTime

`func (o *TokenTransferTransaction1) SetBurnBlockTime(v int32)`

SetBurnBlockTime sets BurnBlockTime field to given value.


### GetBurnBlockHeight

`func (o *TokenTransferTransaction1) GetBurnBlockHeight() int32`

GetBurnBlockHeight returns the BurnBlockHeight field if non-nil, zero value otherwise.

### GetBurnBlockHeightOk

`func (o *TokenTransferTransaction1) GetBurnBlockHeightOk() (*int32, bool)`

GetBurnBlockHeightOk returns a tuple with the BurnBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockHeight

`func (o *TokenTransferTransaction1) SetBurnBlockHeight(v int32)`

SetBurnBlockHeight sets BurnBlockHeight field to given value.


### GetBurnBlockTimeIso

`func (o *TokenTransferTransaction1) GetBurnBlockTimeIso() string`

GetBurnBlockTimeIso returns the BurnBlockTimeIso field if non-nil, zero value otherwise.

### GetBurnBlockTimeIsoOk

`func (o *TokenTransferTransaction1) GetBurnBlockTimeIsoOk() (*string, bool)`

GetBurnBlockTimeIsoOk returns a tuple with the BurnBlockTimeIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockTimeIso

`func (o *TokenTransferTransaction1) SetBurnBlockTimeIso(v string)`

SetBurnBlockTimeIso sets BurnBlockTimeIso field to given value.


### GetParentBurnBlockTime

`func (o *TokenTransferTransaction1) GetParentBurnBlockTime() int32`

GetParentBurnBlockTime returns the ParentBurnBlockTime field if non-nil, zero value otherwise.

### GetParentBurnBlockTimeOk

`func (o *TokenTransferTransaction1) GetParentBurnBlockTimeOk() (*int32, bool)`

GetParentBurnBlockTimeOk returns a tuple with the ParentBurnBlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBurnBlockTime

`func (o *TokenTransferTransaction1) SetParentBurnBlockTime(v int32)`

SetParentBurnBlockTime sets ParentBurnBlockTime field to given value.


### GetParentBurnBlockTimeIso

`func (o *TokenTransferTransaction1) GetParentBurnBlockTimeIso() string`

GetParentBurnBlockTimeIso returns the ParentBurnBlockTimeIso field if non-nil, zero value otherwise.

### GetParentBurnBlockTimeIsoOk

`func (o *TokenTransferTransaction1) GetParentBurnBlockTimeIsoOk() (*string, bool)`

GetParentBurnBlockTimeIsoOk returns a tuple with the ParentBurnBlockTimeIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBurnBlockTimeIso

`func (o *TokenTransferTransaction1) SetParentBurnBlockTimeIso(v string)`

SetParentBurnBlockTimeIso sets ParentBurnBlockTimeIso field to given value.


### GetCanonical

`func (o *TokenTransferTransaction1) GetCanonical() bool`

GetCanonical returns the Canonical field if non-nil, zero value otherwise.

### GetCanonicalOk

`func (o *TokenTransferTransaction1) GetCanonicalOk() (*bool, bool)`

GetCanonicalOk returns a tuple with the Canonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonical

`func (o *TokenTransferTransaction1) SetCanonical(v bool)`

SetCanonical sets Canonical field to given value.


### GetTxIndex

`func (o *TokenTransferTransaction1) GetTxIndex() int32`

GetTxIndex returns the TxIndex field if non-nil, zero value otherwise.

### GetTxIndexOk

`func (o *TokenTransferTransaction1) GetTxIndexOk() (*int32, bool)`

GetTxIndexOk returns a tuple with the TxIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxIndex

`func (o *TokenTransferTransaction1) SetTxIndex(v int32)`

SetTxIndex sets TxIndex field to given value.


### GetTxStatus

`func (o *TokenTransferTransaction1) GetTxStatus() TokenTransferTransaction1TxStatus`

GetTxStatus returns the TxStatus field if non-nil, zero value otherwise.

### GetTxStatusOk

`func (o *TokenTransferTransaction1) GetTxStatusOk() (*TokenTransferTransaction1TxStatus, bool)`

GetTxStatusOk returns a tuple with the TxStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxStatus

`func (o *TokenTransferTransaction1) SetTxStatus(v TokenTransferTransaction1TxStatus)`

SetTxStatus sets TxStatus field to given value.


### GetTxResult

`func (o *TokenTransferTransaction1) GetTxResult() TokenTransferTransactionTxResult`

GetTxResult returns the TxResult field if non-nil, zero value otherwise.

### GetTxResultOk

`func (o *TokenTransferTransaction1) GetTxResultOk() (*TokenTransferTransactionTxResult, bool)`

GetTxResultOk returns a tuple with the TxResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxResult

`func (o *TokenTransferTransaction1) SetTxResult(v TokenTransferTransactionTxResult)`

SetTxResult sets TxResult field to given value.


### GetEventCount

`func (o *TokenTransferTransaction1) GetEventCount() int32`

GetEventCount returns the EventCount field if non-nil, zero value otherwise.

### GetEventCountOk

`func (o *TokenTransferTransaction1) GetEventCountOk() (*int32, bool)`

GetEventCountOk returns a tuple with the EventCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCount

`func (o *TokenTransferTransaction1) SetEventCount(v int32)`

SetEventCount sets EventCount field to given value.


### GetParentBlockHash

`func (o *TokenTransferTransaction1) GetParentBlockHash() string`

GetParentBlockHash returns the ParentBlockHash field if non-nil, zero value otherwise.

### GetParentBlockHashOk

`func (o *TokenTransferTransaction1) GetParentBlockHashOk() (*string, bool)`

GetParentBlockHashOk returns a tuple with the ParentBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBlockHash

`func (o *TokenTransferTransaction1) SetParentBlockHash(v string)`

SetParentBlockHash sets ParentBlockHash field to given value.


### GetIsUnanchored

`func (o *TokenTransferTransaction1) GetIsUnanchored() bool`

GetIsUnanchored returns the IsUnanchored field if non-nil, zero value otherwise.

### GetIsUnanchoredOk

`func (o *TokenTransferTransaction1) GetIsUnanchoredOk() (*bool, bool)`

GetIsUnanchoredOk returns a tuple with the IsUnanchored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnanchored

`func (o *TokenTransferTransaction1) SetIsUnanchored(v bool)`

SetIsUnanchored sets IsUnanchored field to given value.


### GetMicroblockHash

`func (o *TokenTransferTransaction1) GetMicroblockHash() string`

GetMicroblockHash returns the MicroblockHash field if non-nil, zero value otherwise.

### GetMicroblockHashOk

`func (o *TokenTransferTransaction1) GetMicroblockHashOk() (*string, bool)`

GetMicroblockHashOk returns a tuple with the MicroblockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroblockHash

`func (o *TokenTransferTransaction1) SetMicroblockHash(v string)`

SetMicroblockHash sets MicroblockHash field to given value.


### GetMicroblockSequence

`func (o *TokenTransferTransaction1) GetMicroblockSequence() int32`

GetMicroblockSequence returns the MicroblockSequence field if non-nil, zero value otherwise.

### GetMicroblockSequenceOk

`func (o *TokenTransferTransaction1) GetMicroblockSequenceOk() (*int32, bool)`

GetMicroblockSequenceOk returns a tuple with the MicroblockSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroblockSequence

`func (o *TokenTransferTransaction1) SetMicroblockSequence(v int32)`

SetMicroblockSequence sets MicroblockSequence field to given value.


### GetMicroblockCanonical

`func (o *TokenTransferTransaction1) GetMicroblockCanonical() bool`

GetMicroblockCanonical returns the MicroblockCanonical field if non-nil, zero value otherwise.

### GetMicroblockCanonicalOk

`func (o *TokenTransferTransaction1) GetMicroblockCanonicalOk() (*bool, bool)`

GetMicroblockCanonicalOk returns a tuple with the MicroblockCanonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroblockCanonical

`func (o *TokenTransferTransaction1) SetMicroblockCanonical(v bool)`

SetMicroblockCanonical sets MicroblockCanonical field to given value.


### GetExecutionCostReadCount

`func (o *TokenTransferTransaction1) GetExecutionCostReadCount() int32`

GetExecutionCostReadCount returns the ExecutionCostReadCount field if non-nil, zero value otherwise.

### GetExecutionCostReadCountOk

`func (o *TokenTransferTransaction1) GetExecutionCostReadCountOk() (*int32, bool)`

GetExecutionCostReadCountOk returns a tuple with the ExecutionCostReadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCostReadCount

`func (o *TokenTransferTransaction1) SetExecutionCostReadCount(v int32)`

SetExecutionCostReadCount sets ExecutionCostReadCount field to given value.


### GetExecutionCostReadLength

`func (o *TokenTransferTransaction1) GetExecutionCostReadLength() int32`

GetExecutionCostReadLength returns the ExecutionCostReadLength field if non-nil, zero value otherwise.

### GetExecutionCostReadLengthOk

`func (o *TokenTransferTransaction1) GetExecutionCostReadLengthOk() (*int32, bool)`

GetExecutionCostReadLengthOk returns a tuple with the ExecutionCostReadLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCostReadLength

`func (o *TokenTransferTransaction1) SetExecutionCostReadLength(v int32)`

SetExecutionCostReadLength sets ExecutionCostReadLength field to given value.


### GetExecutionCostRuntime

`func (o *TokenTransferTransaction1) GetExecutionCostRuntime() int32`

GetExecutionCostRuntime returns the ExecutionCostRuntime field if non-nil, zero value otherwise.

### GetExecutionCostRuntimeOk

`func (o *TokenTransferTransaction1) GetExecutionCostRuntimeOk() (*int32, bool)`

GetExecutionCostRuntimeOk returns a tuple with the ExecutionCostRuntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCostRuntime

`func (o *TokenTransferTransaction1) SetExecutionCostRuntime(v int32)`

SetExecutionCostRuntime sets ExecutionCostRuntime field to given value.


### GetExecutionCostWriteCount

`func (o *TokenTransferTransaction1) GetExecutionCostWriteCount() int32`

GetExecutionCostWriteCount returns the ExecutionCostWriteCount field if non-nil, zero value otherwise.

### GetExecutionCostWriteCountOk

`func (o *TokenTransferTransaction1) GetExecutionCostWriteCountOk() (*int32, bool)`

GetExecutionCostWriteCountOk returns a tuple with the ExecutionCostWriteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCostWriteCount

`func (o *TokenTransferTransaction1) SetExecutionCostWriteCount(v int32)`

SetExecutionCostWriteCount sets ExecutionCostWriteCount field to given value.


### GetExecutionCostWriteLength

`func (o *TokenTransferTransaction1) GetExecutionCostWriteLength() int32`

GetExecutionCostWriteLength returns the ExecutionCostWriteLength field if non-nil, zero value otherwise.

### GetExecutionCostWriteLengthOk

`func (o *TokenTransferTransaction1) GetExecutionCostWriteLengthOk() (*int32, bool)`

GetExecutionCostWriteLengthOk returns a tuple with the ExecutionCostWriteLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCostWriteLength

`func (o *TokenTransferTransaction1) SetExecutionCostWriteLength(v int32)`

SetExecutionCostWriteLength sets ExecutionCostWriteLength field to given value.


### GetEvents

`func (o *TokenTransferTransaction1) GetEvents() []TokenTransferTransaction1EventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *TokenTransferTransaction1) GetEventsOk() (*[]TokenTransferTransaction1EventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *TokenTransferTransaction1) SetEvents(v []TokenTransferTransaction1EventsInner)`

SetEvents sets Events field to given value.


### GetTxType

`func (o *TokenTransferTransaction1) GetTxType() string`

GetTxType returns the TxType field if non-nil, zero value otherwise.

### GetTxTypeOk

`func (o *TokenTransferTransaction1) GetTxTypeOk() (*string, bool)`

GetTxTypeOk returns a tuple with the TxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxType

`func (o *TokenTransferTransaction1) SetTxType(v string)`

SetTxType sets TxType field to given value.


### GetTokenTransfer

`func (o *TokenTransferTransaction1) GetTokenTransfer() TokenTransferTransactionTokenTransfer`

GetTokenTransfer returns the TokenTransfer field if non-nil, zero value otherwise.

### GetTokenTransferOk

`func (o *TokenTransferTransaction1) GetTokenTransferOk() (*TokenTransferTransactionTokenTransfer, bool)`

GetTokenTransferOk returns a tuple with the TokenTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTransfer

`func (o *TokenTransferTransaction1) SetTokenTransfer(v TokenTransferTransactionTokenTransfer)`

SetTokenTransfer sets TokenTransfer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


