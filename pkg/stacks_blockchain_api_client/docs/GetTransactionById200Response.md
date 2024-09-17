# GetTransactionById200Response

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
**TxStatus** | [**TokenTransferMempoolTransaction1TxStatus**](TokenTransferMempoolTransaction1TxStatus.md) |  | 
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
**SmartContract** | [**SmartContractTransactionSmartContract**](SmartContractTransactionSmartContract.md) |  | 
**ContractCall** | [**ContractCallTransactionContractCall**](ContractCallTransactionContractCall.md) |  | 
**PoisonMicroblock** | [**PoisonMicroblockTransactionPoisonMicroblock**](PoisonMicroblockTransactionPoisonMicroblock.md) |  | 
**CoinbasePayload** | [**CoinbaseTransactionCoinbasePayload**](CoinbaseTransactionCoinbasePayload.md) |  | 
**TenureChangePayload** | [**TenureChangeTransactionTenureChangePayload**](TenureChangeTransactionTenureChangePayload.md) |  | 
**ReceiptTime** | **int32** | A unix timestamp (in seconds) indicating when the transaction broadcast was received by the node. | 
**ReceiptTimeIso** | **string** | An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) timestamp indicating when the transaction broadcast was received by the node. | 

## Methods

### NewGetTransactionById200Response

`func NewGetTransactionById200Response(txId string, nonce int32, feeRate string, senderAddress string, sponsored bool, postConditionMode TokenTransferTransactionPostConditionMode, postConditions []TokenTransferTransactionPostConditionsInner, anchorMode TokenTransferTransactionAnchorMode, blockHash string, blockHeight int32, blockTime float32, blockTimeIso string, burnBlockTime int32, burnBlockHeight int32, burnBlockTimeIso string, parentBurnBlockTime int32, parentBurnBlockTimeIso string, canonical bool, txIndex int32, txStatus TokenTransferMempoolTransaction1TxStatus, txResult TokenTransferTransactionTxResult, eventCount int32, parentBlockHash string, isUnanchored bool, microblockHash string, microblockSequence int32, microblockCanonical bool, executionCostReadCount int32, executionCostReadLength int32, executionCostRuntime int32, executionCostWriteCount int32, executionCostWriteLength int32, events []TokenTransferTransactionEventsInner, txType string, tokenTransfer TokenTransferTransactionTokenTransfer, smartContract SmartContractTransactionSmartContract, contractCall ContractCallTransactionContractCall, poisonMicroblock PoisonMicroblockTransactionPoisonMicroblock, coinbasePayload CoinbaseTransactionCoinbasePayload, tenureChangePayload TenureChangeTransactionTenureChangePayload, receiptTime int32, receiptTimeIso string, ) *GetTransactionById200Response`

NewGetTransactionById200Response instantiates a new GetTransactionById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionById200ResponseWithDefaults

`func NewGetTransactionById200ResponseWithDefaults() *GetTransactionById200Response`

NewGetTransactionById200ResponseWithDefaults instantiates a new GetTransactionById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxId

`func (o *GetTransactionById200Response) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *GetTransactionById200Response) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *GetTransactionById200Response) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetNonce

`func (o *GetTransactionById200Response) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *GetTransactionById200Response) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *GetTransactionById200Response) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetFeeRate

`func (o *GetTransactionById200Response) GetFeeRate() string`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *GetTransactionById200Response) GetFeeRateOk() (*string, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *GetTransactionById200Response) SetFeeRate(v string)`

SetFeeRate sets FeeRate field to given value.


### GetSenderAddress

`func (o *GetTransactionById200Response) GetSenderAddress() string`

GetSenderAddress returns the SenderAddress field if non-nil, zero value otherwise.

### GetSenderAddressOk

`func (o *GetTransactionById200Response) GetSenderAddressOk() (*string, bool)`

GetSenderAddressOk returns a tuple with the SenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAddress

`func (o *GetTransactionById200Response) SetSenderAddress(v string)`

SetSenderAddress sets SenderAddress field to given value.


### GetSponsorNonce

`func (o *GetTransactionById200Response) GetSponsorNonce() int32`

GetSponsorNonce returns the SponsorNonce field if non-nil, zero value otherwise.

### GetSponsorNonceOk

`func (o *GetTransactionById200Response) GetSponsorNonceOk() (*int32, bool)`

GetSponsorNonceOk returns a tuple with the SponsorNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorNonce

`func (o *GetTransactionById200Response) SetSponsorNonce(v int32)`

SetSponsorNonce sets SponsorNonce field to given value.

### HasSponsorNonce

`func (o *GetTransactionById200Response) HasSponsorNonce() bool`

HasSponsorNonce returns a boolean if a field has been set.

### GetSponsored

`func (o *GetTransactionById200Response) GetSponsored() bool`

GetSponsored returns the Sponsored field if non-nil, zero value otherwise.

### GetSponsoredOk

`func (o *GetTransactionById200Response) GetSponsoredOk() (*bool, bool)`

GetSponsoredOk returns a tuple with the Sponsored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsored

`func (o *GetTransactionById200Response) SetSponsored(v bool)`

SetSponsored sets Sponsored field to given value.


### GetSponsorAddress

`func (o *GetTransactionById200Response) GetSponsorAddress() string`

GetSponsorAddress returns the SponsorAddress field if non-nil, zero value otherwise.

### GetSponsorAddressOk

`func (o *GetTransactionById200Response) GetSponsorAddressOk() (*string, bool)`

GetSponsorAddressOk returns a tuple with the SponsorAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorAddress

`func (o *GetTransactionById200Response) SetSponsorAddress(v string)`

SetSponsorAddress sets SponsorAddress field to given value.

### HasSponsorAddress

`func (o *GetTransactionById200Response) HasSponsorAddress() bool`

HasSponsorAddress returns a boolean if a field has been set.

### GetPostConditionMode

`func (o *GetTransactionById200Response) GetPostConditionMode() TokenTransferTransactionPostConditionMode`

GetPostConditionMode returns the PostConditionMode field if non-nil, zero value otherwise.

### GetPostConditionModeOk

`func (o *GetTransactionById200Response) GetPostConditionModeOk() (*TokenTransferTransactionPostConditionMode, bool)`

GetPostConditionModeOk returns a tuple with the PostConditionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostConditionMode

`func (o *GetTransactionById200Response) SetPostConditionMode(v TokenTransferTransactionPostConditionMode)`

SetPostConditionMode sets PostConditionMode field to given value.


### GetPostConditions

`func (o *GetTransactionById200Response) GetPostConditions() []TokenTransferTransactionPostConditionsInner`

GetPostConditions returns the PostConditions field if non-nil, zero value otherwise.

### GetPostConditionsOk

`func (o *GetTransactionById200Response) GetPostConditionsOk() (*[]TokenTransferTransactionPostConditionsInner, bool)`

GetPostConditionsOk returns a tuple with the PostConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostConditions

`func (o *GetTransactionById200Response) SetPostConditions(v []TokenTransferTransactionPostConditionsInner)`

SetPostConditions sets PostConditions field to given value.


### GetAnchorMode

`func (o *GetTransactionById200Response) GetAnchorMode() TokenTransferTransactionAnchorMode`

GetAnchorMode returns the AnchorMode field if non-nil, zero value otherwise.

### GetAnchorModeOk

`func (o *GetTransactionById200Response) GetAnchorModeOk() (*TokenTransferTransactionAnchorMode, bool)`

GetAnchorModeOk returns a tuple with the AnchorMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorMode

`func (o *GetTransactionById200Response) SetAnchorMode(v TokenTransferTransactionAnchorMode)`

SetAnchorMode sets AnchorMode field to given value.


### GetBlockHash

`func (o *GetTransactionById200Response) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *GetTransactionById200Response) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *GetTransactionById200Response) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.


### GetBlockHeight

`func (o *GetTransactionById200Response) GetBlockHeight() int32`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *GetTransactionById200Response) GetBlockHeightOk() (*int32, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *GetTransactionById200Response) SetBlockHeight(v int32)`

SetBlockHeight sets BlockHeight field to given value.


### GetBlockTime

`func (o *GetTransactionById200Response) GetBlockTime() float32`

GetBlockTime returns the BlockTime field if non-nil, zero value otherwise.

### GetBlockTimeOk

`func (o *GetTransactionById200Response) GetBlockTimeOk() (*float32, bool)`

GetBlockTimeOk returns a tuple with the BlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTime

`func (o *GetTransactionById200Response) SetBlockTime(v float32)`

SetBlockTime sets BlockTime field to given value.


### GetBlockTimeIso

`func (o *GetTransactionById200Response) GetBlockTimeIso() string`

GetBlockTimeIso returns the BlockTimeIso field if non-nil, zero value otherwise.

### GetBlockTimeIsoOk

`func (o *GetTransactionById200Response) GetBlockTimeIsoOk() (*string, bool)`

GetBlockTimeIsoOk returns a tuple with the BlockTimeIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTimeIso

`func (o *GetTransactionById200Response) SetBlockTimeIso(v string)`

SetBlockTimeIso sets BlockTimeIso field to given value.


### GetBurnBlockTime

`func (o *GetTransactionById200Response) GetBurnBlockTime() int32`

GetBurnBlockTime returns the BurnBlockTime field if non-nil, zero value otherwise.

### GetBurnBlockTimeOk

`func (o *GetTransactionById200Response) GetBurnBlockTimeOk() (*int32, bool)`

GetBurnBlockTimeOk returns a tuple with the BurnBlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockTime

`func (o *GetTransactionById200Response) SetBurnBlockTime(v int32)`

SetBurnBlockTime sets BurnBlockTime field to given value.


### GetBurnBlockHeight

`func (o *GetTransactionById200Response) GetBurnBlockHeight() int32`

GetBurnBlockHeight returns the BurnBlockHeight field if non-nil, zero value otherwise.

### GetBurnBlockHeightOk

`func (o *GetTransactionById200Response) GetBurnBlockHeightOk() (*int32, bool)`

GetBurnBlockHeightOk returns a tuple with the BurnBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockHeight

`func (o *GetTransactionById200Response) SetBurnBlockHeight(v int32)`

SetBurnBlockHeight sets BurnBlockHeight field to given value.


### GetBurnBlockTimeIso

`func (o *GetTransactionById200Response) GetBurnBlockTimeIso() string`

GetBurnBlockTimeIso returns the BurnBlockTimeIso field if non-nil, zero value otherwise.

### GetBurnBlockTimeIsoOk

`func (o *GetTransactionById200Response) GetBurnBlockTimeIsoOk() (*string, bool)`

GetBurnBlockTimeIsoOk returns a tuple with the BurnBlockTimeIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockTimeIso

`func (o *GetTransactionById200Response) SetBurnBlockTimeIso(v string)`

SetBurnBlockTimeIso sets BurnBlockTimeIso field to given value.


### GetParentBurnBlockTime

`func (o *GetTransactionById200Response) GetParentBurnBlockTime() int32`

GetParentBurnBlockTime returns the ParentBurnBlockTime field if non-nil, zero value otherwise.

### GetParentBurnBlockTimeOk

`func (o *GetTransactionById200Response) GetParentBurnBlockTimeOk() (*int32, bool)`

GetParentBurnBlockTimeOk returns a tuple with the ParentBurnBlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBurnBlockTime

`func (o *GetTransactionById200Response) SetParentBurnBlockTime(v int32)`

SetParentBurnBlockTime sets ParentBurnBlockTime field to given value.


### GetParentBurnBlockTimeIso

`func (o *GetTransactionById200Response) GetParentBurnBlockTimeIso() string`

GetParentBurnBlockTimeIso returns the ParentBurnBlockTimeIso field if non-nil, zero value otherwise.

### GetParentBurnBlockTimeIsoOk

`func (o *GetTransactionById200Response) GetParentBurnBlockTimeIsoOk() (*string, bool)`

GetParentBurnBlockTimeIsoOk returns a tuple with the ParentBurnBlockTimeIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBurnBlockTimeIso

`func (o *GetTransactionById200Response) SetParentBurnBlockTimeIso(v string)`

SetParentBurnBlockTimeIso sets ParentBurnBlockTimeIso field to given value.


### GetCanonical

`func (o *GetTransactionById200Response) GetCanonical() bool`

GetCanonical returns the Canonical field if non-nil, zero value otherwise.

### GetCanonicalOk

`func (o *GetTransactionById200Response) GetCanonicalOk() (*bool, bool)`

GetCanonicalOk returns a tuple with the Canonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonical

`func (o *GetTransactionById200Response) SetCanonical(v bool)`

SetCanonical sets Canonical field to given value.


### GetTxIndex

`func (o *GetTransactionById200Response) GetTxIndex() int32`

GetTxIndex returns the TxIndex field if non-nil, zero value otherwise.

### GetTxIndexOk

`func (o *GetTransactionById200Response) GetTxIndexOk() (*int32, bool)`

GetTxIndexOk returns a tuple with the TxIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxIndex

`func (o *GetTransactionById200Response) SetTxIndex(v int32)`

SetTxIndex sets TxIndex field to given value.


### GetTxStatus

`func (o *GetTransactionById200Response) GetTxStatus() TokenTransferMempoolTransaction1TxStatus`

GetTxStatus returns the TxStatus field if non-nil, zero value otherwise.

### GetTxStatusOk

`func (o *GetTransactionById200Response) GetTxStatusOk() (*TokenTransferMempoolTransaction1TxStatus, bool)`

GetTxStatusOk returns a tuple with the TxStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxStatus

`func (o *GetTransactionById200Response) SetTxStatus(v TokenTransferMempoolTransaction1TxStatus)`

SetTxStatus sets TxStatus field to given value.


### GetTxResult

`func (o *GetTransactionById200Response) GetTxResult() TokenTransferTransactionTxResult`

GetTxResult returns the TxResult field if non-nil, zero value otherwise.

### GetTxResultOk

`func (o *GetTransactionById200Response) GetTxResultOk() (*TokenTransferTransactionTxResult, bool)`

GetTxResultOk returns a tuple with the TxResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxResult

`func (o *GetTransactionById200Response) SetTxResult(v TokenTransferTransactionTxResult)`

SetTxResult sets TxResult field to given value.


### GetEventCount

`func (o *GetTransactionById200Response) GetEventCount() int32`

GetEventCount returns the EventCount field if non-nil, zero value otherwise.

### GetEventCountOk

`func (o *GetTransactionById200Response) GetEventCountOk() (*int32, bool)`

GetEventCountOk returns a tuple with the EventCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCount

`func (o *GetTransactionById200Response) SetEventCount(v int32)`

SetEventCount sets EventCount field to given value.


### GetParentBlockHash

`func (o *GetTransactionById200Response) GetParentBlockHash() string`

GetParentBlockHash returns the ParentBlockHash field if non-nil, zero value otherwise.

### GetParentBlockHashOk

`func (o *GetTransactionById200Response) GetParentBlockHashOk() (*string, bool)`

GetParentBlockHashOk returns a tuple with the ParentBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBlockHash

`func (o *GetTransactionById200Response) SetParentBlockHash(v string)`

SetParentBlockHash sets ParentBlockHash field to given value.


### GetIsUnanchored

`func (o *GetTransactionById200Response) GetIsUnanchored() bool`

GetIsUnanchored returns the IsUnanchored field if non-nil, zero value otherwise.

### GetIsUnanchoredOk

`func (o *GetTransactionById200Response) GetIsUnanchoredOk() (*bool, bool)`

GetIsUnanchoredOk returns a tuple with the IsUnanchored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnanchored

`func (o *GetTransactionById200Response) SetIsUnanchored(v bool)`

SetIsUnanchored sets IsUnanchored field to given value.


### GetMicroblockHash

`func (o *GetTransactionById200Response) GetMicroblockHash() string`

GetMicroblockHash returns the MicroblockHash field if non-nil, zero value otherwise.

### GetMicroblockHashOk

`func (o *GetTransactionById200Response) GetMicroblockHashOk() (*string, bool)`

GetMicroblockHashOk returns a tuple with the MicroblockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroblockHash

`func (o *GetTransactionById200Response) SetMicroblockHash(v string)`

SetMicroblockHash sets MicroblockHash field to given value.


### GetMicroblockSequence

`func (o *GetTransactionById200Response) GetMicroblockSequence() int32`

GetMicroblockSequence returns the MicroblockSequence field if non-nil, zero value otherwise.

### GetMicroblockSequenceOk

`func (o *GetTransactionById200Response) GetMicroblockSequenceOk() (*int32, bool)`

GetMicroblockSequenceOk returns a tuple with the MicroblockSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroblockSequence

`func (o *GetTransactionById200Response) SetMicroblockSequence(v int32)`

SetMicroblockSequence sets MicroblockSequence field to given value.


### GetMicroblockCanonical

`func (o *GetTransactionById200Response) GetMicroblockCanonical() bool`

GetMicroblockCanonical returns the MicroblockCanonical field if non-nil, zero value otherwise.

### GetMicroblockCanonicalOk

`func (o *GetTransactionById200Response) GetMicroblockCanonicalOk() (*bool, bool)`

GetMicroblockCanonicalOk returns a tuple with the MicroblockCanonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroblockCanonical

`func (o *GetTransactionById200Response) SetMicroblockCanonical(v bool)`

SetMicroblockCanonical sets MicroblockCanonical field to given value.


### GetExecutionCostReadCount

`func (o *GetTransactionById200Response) GetExecutionCostReadCount() int32`

GetExecutionCostReadCount returns the ExecutionCostReadCount field if non-nil, zero value otherwise.

### GetExecutionCostReadCountOk

`func (o *GetTransactionById200Response) GetExecutionCostReadCountOk() (*int32, bool)`

GetExecutionCostReadCountOk returns a tuple with the ExecutionCostReadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCostReadCount

`func (o *GetTransactionById200Response) SetExecutionCostReadCount(v int32)`

SetExecutionCostReadCount sets ExecutionCostReadCount field to given value.


### GetExecutionCostReadLength

`func (o *GetTransactionById200Response) GetExecutionCostReadLength() int32`

GetExecutionCostReadLength returns the ExecutionCostReadLength field if non-nil, zero value otherwise.

### GetExecutionCostReadLengthOk

`func (o *GetTransactionById200Response) GetExecutionCostReadLengthOk() (*int32, bool)`

GetExecutionCostReadLengthOk returns a tuple with the ExecutionCostReadLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCostReadLength

`func (o *GetTransactionById200Response) SetExecutionCostReadLength(v int32)`

SetExecutionCostReadLength sets ExecutionCostReadLength field to given value.


### GetExecutionCostRuntime

`func (o *GetTransactionById200Response) GetExecutionCostRuntime() int32`

GetExecutionCostRuntime returns the ExecutionCostRuntime field if non-nil, zero value otherwise.

### GetExecutionCostRuntimeOk

`func (o *GetTransactionById200Response) GetExecutionCostRuntimeOk() (*int32, bool)`

GetExecutionCostRuntimeOk returns a tuple with the ExecutionCostRuntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCostRuntime

`func (o *GetTransactionById200Response) SetExecutionCostRuntime(v int32)`

SetExecutionCostRuntime sets ExecutionCostRuntime field to given value.


### GetExecutionCostWriteCount

`func (o *GetTransactionById200Response) GetExecutionCostWriteCount() int32`

GetExecutionCostWriteCount returns the ExecutionCostWriteCount field if non-nil, zero value otherwise.

### GetExecutionCostWriteCountOk

`func (o *GetTransactionById200Response) GetExecutionCostWriteCountOk() (*int32, bool)`

GetExecutionCostWriteCountOk returns a tuple with the ExecutionCostWriteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCostWriteCount

`func (o *GetTransactionById200Response) SetExecutionCostWriteCount(v int32)`

SetExecutionCostWriteCount sets ExecutionCostWriteCount field to given value.


### GetExecutionCostWriteLength

`func (o *GetTransactionById200Response) GetExecutionCostWriteLength() int32`

GetExecutionCostWriteLength returns the ExecutionCostWriteLength field if non-nil, zero value otherwise.

### GetExecutionCostWriteLengthOk

`func (o *GetTransactionById200Response) GetExecutionCostWriteLengthOk() (*int32, bool)`

GetExecutionCostWriteLengthOk returns a tuple with the ExecutionCostWriteLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCostWriteLength

`func (o *GetTransactionById200Response) SetExecutionCostWriteLength(v int32)`

SetExecutionCostWriteLength sets ExecutionCostWriteLength field to given value.


### GetEvents

`func (o *GetTransactionById200Response) GetEvents() []TokenTransferTransactionEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GetTransactionById200Response) GetEventsOk() (*[]TokenTransferTransactionEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GetTransactionById200Response) SetEvents(v []TokenTransferTransactionEventsInner)`

SetEvents sets Events field to given value.


### GetTxType

`func (o *GetTransactionById200Response) GetTxType() string`

GetTxType returns the TxType field if non-nil, zero value otherwise.

### GetTxTypeOk

`func (o *GetTransactionById200Response) GetTxTypeOk() (*string, bool)`

GetTxTypeOk returns a tuple with the TxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxType

`func (o *GetTransactionById200Response) SetTxType(v string)`

SetTxType sets TxType field to given value.


### GetTokenTransfer

`func (o *GetTransactionById200Response) GetTokenTransfer() TokenTransferTransactionTokenTransfer`

GetTokenTransfer returns the TokenTransfer field if non-nil, zero value otherwise.

### GetTokenTransferOk

`func (o *GetTransactionById200Response) GetTokenTransferOk() (*TokenTransferTransactionTokenTransfer, bool)`

GetTokenTransferOk returns a tuple with the TokenTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTransfer

`func (o *GetTransactionById200Response) SetTokenTransfer(v TokenTransferTransactionTokenTransfer)`

SetTokenTransfer sets TokenTransfer field to given value.


### GetSmartContract

`func (o *GetTransactionById200Response) GetSmartContract() SmartContractTransactionSmartContract`

GetSmartContract returns the SmartContract field if non-nil, zero value otherwise.

### GetSmartContractOk

`func (o *GetTransactionById200Response) GetSmartContractOk() (*SmartContractTransactionSmartContract, bool)`

GetSmartContractOk returns a tuple with the SmartContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartContract

`func (o *GetTransactionById200Response) SetSmartContract(v SmartContractTransactionSmartContract)`

SetSmartContract sets SmartContract field to given value.


### GetContractCall

`func (o *GetTransactionById200Response) GetContractCall() ContractCallTransactionContractCall`

GetContractCall returns the ContractCall field if non-nil, zero value otherwise.

### GetContractCallOk

`func (o *GetTransactionById200Response) GetContractCallOk() (*ContractCallTransactionContractCall, bool)`

GetContractCallOk returns a tuple with the ContractCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractCall

`func (o *GetTransactionById200Response) SetContractCall(v ContractCallTransactionContractCall)`

SetContractCall sets ContractCall field to given value.


### GetPoisonMicroblock

`func (o *GetTransactionById200Response) GetPoisonMicroblock() PoisonMicroblockTransactionPoisonMicroblock`

GetPoisonMicroblock returns the PoisonMicroblock field if non-nil, zero value otherwise.

### GetPoisonMicroblockOk

`func (o *GetTransactionById200Response) GetPoisonMicroblockOk() (*PoisonMicroblockTransactionPoisonMicroblock, bool)`

GetPoisonMicroblockOk returns a tuple with the PoisonMicroblock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoisonMicroblock

`func (o *GetTransactionById200Response) SetPoisonMicroblock(v PoisonMicroblockTransactionPoisonMicroblock)`

SetPoisonMicroblock sets PoisonMicroblock field to given value.


### GetCoinbasePayload

`func (o *GetTransactionById200Response) GetCoinbasePayload() CoinbaseTransactionCoinbasePayload`

GetCoinbasePayload returns the CoinbasePayload field if non-nil, zero value otherwise.

### GetCoinbasePayloadOk

`func (o *GetTransactionById200Response) GetCoinbasePayloadOk() (*CoinbaseTransactionCoinbasePayload, bool)`

GetCoinbasePayloadOk returns a tuple with the CoinbasePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinbasePayload

`func (o *GetTransactionById200Response) SetCoinbasePayload(v CoinbaseTransactionCoinbasePayload)`

SetCoinbasePayload sets CoinbasePayload field to given value.


### GetTenureChangePayload

`func (o *GetTransactionById200Response) GetTenureChangePayload() TenureChangeTransactionTenureChangePayload`

GetTenureChangePayload returns the TenureChangePayload field if non-nil, zero value otherwise.

### GetTenureChangePayloadOk

`func (o *GetTransactionById200Response) GetTenureChangePayloadOk() (*TenureChangeTransactionTenureChangePayload, bool)`

GetTenureChangePayloadOk returns a tuple with the TenureChangePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenureChangePayload

`func (o *GetTransactionById200Response) SetTenureChangePayload(v TenureChangeTransactionTenureChangePayload)`

SetTenureChangePayload sets TenureChangePayload field to given value.


### GetReceiptTime

`func (o *GetTransactionById200Response) GetReceiptTime() int32`

GetReceiptTime returns the ReceiptTime field if non-nil, zero value otherwise.

### GetReceiptTimeOk

`func (o *GetTransactionById200Response) GetReceiptTimeOk() (*int32, bool)`

GetReceiptTimeOk returns a tuple with the ReceiptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptTime

`func (o *GetTransactionById200Response) SetReceiptTime(v int32)`

SetReceiptTime sets ReceiptTime field to given value.


### GetReceiptTimeIso

`func (o *GetTransactionById200Response) GetReceiptTimeIso() string`

GetReceiptTimeIso returns the ReceiptTimeIso field if non-nil, zero value otherwise.

### GetReceiptTimeIsoOk

`func (o *GetTransactionById200Response) GetReceiptTimeIsoOk() (*string, bool)`

GetReceiptTimeIsoOk returns a tuple with the ReceiptTimeIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptTimeIso

`func (o *GetTransactionById200Response) SetReceiptTimeIso(v string)`

SetReceiptTimeIso sets ReceiptTimeIso field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


