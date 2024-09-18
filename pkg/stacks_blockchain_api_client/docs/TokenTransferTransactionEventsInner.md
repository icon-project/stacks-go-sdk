# TokenTransferTransactionEventsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventIndex** | **int32** |  | 
**EventType** | **string** |  | 
**TxId** | **string** |  | 
**ContractLog** | [**SmartContractLogTransactionEventAllOfContractLog**](SmartContractLogTransactionEventAllOfContractLog.md) |  | 
**StxLockEvent** | [**StxLockTransactionEventAllOfStxLockEvent**](StxLockTransactionEventAllOfStxLockEvent.md) |  | 
**Asset** | [**NonFungibleTokenAssetTransactionEventAllOfAsset**](NonFungibleTokenAssetTransactionEventAllOfAsset.md) |  | 

## Methods

### NewTokenTransferTransactionEventsInner

`func NewTokenTransferTransactionEventsInner(eventIndex int32, eventType string, txId string, contractLog SmartContractLogTransactionEventAllOfContractLog, stxLockEvent StxLockTransactionEventAllOfStxLockEvent, asset NonFungibleTokenAssetTransactionEventAllOfAsset, ) *TokenTransferTransactionEventsInner`

NewTokenTransferTransactionEventsInner instantiates a new TokenTransferTransactionEventsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenTransferTransactionEventsInnerWithDefaults

`func NewTokenTransferTransactionEventsInnerWithDefaults() *TokenTransferTransactionEventsInner`

NewTokenTransferTransactionEventsInnerWithDefaults instantiates a new TokenTransferTransactionEventsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventIndex

`func (o *TokenTransferTransactionEventsInner) GetEventIndex() int32`

GetEventIndex returns the EventIndex field if non-nil, zero value otherwise.

### GetEventIndexOk

`func (o *TokenTransferTransactionEventsInner) GetEventIndexOk() (*int32, bool)`

GetEventIndexOk returns a tuple with the EventIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIndex

`func (o *TokenTransferTransactionEventsInner) SetEventIndex(v int32)`

SetEventIndex sets EventIndex field to given value.


### GetEventType

`func (o *TokenTransferTransactionEventsInner) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *TokenTransferTransactionEventsInner) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *TokenTransferTransactionEventsInner) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetTxId

`func (o *TokenTransferTransactionEventsInner) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *TokenTransferTransactionEventsInner) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *TokenTransferTransactionEventsInner) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetContractLog

`func (o *TokenTransferTransactionEventsInner) GetContractLog() SmartContractLogTransactionEventAllOfContractLog`

GetContractLog returns the ContractLog field if non-nil, zero value otherwise.

### GetContractLogOk

`func (o *TokenTransferTransactionEventsInner) GetContractLogOk() (*SmartContractLogTransactionEventAllOfContractLog, bool)`

GetContractLogOk returns a tuple with the ContractLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractLog

`func (o *TokenTransferTransactionEventsInner) SetContractLog(v SmartContractLogTransactionEventAllOfContractLog)`

SetContractLog sets ContractLog field to given value.


### GetStxLockEvent

`func (o *TokenTransferTransactionEventsInner) GetStxLockEvent() StxLockTransactionEventAllOfStxLockEvent`

GetStxLockEvent returns the StxLockEvent field if non-nil, zero value otherwise.

### GetStxLockEventOk

`func (o *TokenTransferTransactionEventsInner) GetStxLockEventOk() (*StxLockTransactionEventAllOfStxLockEvent, bool)`

GetStxLockEventOk returns a tuple with the StxLockEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStxLockEvent

`func (o *TokenTransferTransactionEventsInner) SetStxLockEvent(v StxLockTransactionEventAllOfStxLockEvent)`

SetStxLockEvent sets StxLockEvent field to given value.


### GetAsset

`func (o *TokenTransferTransactionEventsInner) GetAsset() NonFungibleTokenAssetTransactionEventAllOfAsset`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *TokenTransferTransactionEventsInner) GetAssetOk() (*NonFungibleTokenAssetTransactionEventAllOfAsset, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *TokenTransferTransactionEventsInner) SetAsset(v NonFungibleTokenAssetTransactionEventAllOfAsset)`

SetAsset sets Asset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


