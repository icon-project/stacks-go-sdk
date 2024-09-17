# TokenTransferTransaction1EventsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventIndex** | **int32** |  | 
**EventType** | **string** |  | 
**TxId** | **string** |  | 
**ContractLog** | [**SmartContractLogTransactionEventAllOfContractLog**](SmartContractLogTransactionEventAllOfContractLog.md) |  | 
**StxLockEvent** | [**StxLockTransactionEventAllOfStxLockEvent**](StxLockTransactionEventAllOfStxLockEvent.md) |  | 
**Asset** | [**NonFungibleTokenAssetTransactionEvent1AllOfAsset**](NonFungibleTokenAssetTransactionEvent1AllOfAsset.md) |  | 

## Methods

### NewTokenTransferTransaction1EventsInner

`func NewTokenTransferTransaction1EventsInner(eventIndex int32, eventType string, txId string, contractLog SmartContractLogTransactionEventAllOfContractLog, stxLockEvent StxLockTransactionEventAllOfStxLockEvent, asset NonFungibleTokenAssetTransactionEvent1AllOfAsset, ) *TokenTransferTransaction1EventsInner`

NewTokenTransferTransaction1EventsInner instantiates a new TokenTransferTransaction1EventsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenTransferTransaction1EventsInnerWithDefaults

`func NewTokenTransferTransaction1EventsInnerWithDefaults() *TokenTransferTransaction1EventsInner`

NewTokenTransferTransaction1EventsInnerWithDefaults instantiates a new TokenTransferTransaction1EventsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventIndex

`func (o *TokenTransferTransaction1EventsInner) GetEventIndex() int32`

GetEventIndex returns the EventIndex field if non-nil, zero value otherwise.

### GetEventIndexOk

`func (o *TokenTransferTransaction1EventsInner) GetEventIndexOk() (*int32, bool)`

GetEventIndexOk returns a tuple with the EventIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIndex

`func (o *TokenTransferTransaction1EventsInner) SetEventIndex(v int32)`

SetEventIndex sets EventIndex field to given value.


### GetEventType

`func (o *TokenTransferTransaction1EventsInner) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *TokenTransferTransaction1EventsInner) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *TokenTransferTransaction1EventsInner) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetTxId

`func (o *TokenTransferTransaction1EventsInner) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *TokenTransferTransaction1EventsInner) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *TokenTransferTransaction1EventsInner) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetContractLog

`func (o *TokenTransferTransaction1EventsInner) GetContractLog() SmartContractLogTransactionEventAllOfContractLog`

GetContractLog returns the ContractLog field if non-nil, zero value otherwise.

### GetContractLogOk

`func (o *TokenTransferTransaction1EventsInner) GetContractLogOk() (*SmartContractLogTransactionEventAllOfContractLog, bool)`

GetContractLogOk returns a tuple with the ContractLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractLog

`func (o *TokenTransferTransaction1EventsInner) SetContractLog(v SmartContractLogTransactionEventAllOfContractLog)`

SetContractLog sets ContractLog field to given value.


### GetStxLockEvent

`func (o *TokenTransferTransaction1EventsInner) GetStxLockEvent() StxLockTransactionEventAllOfStxLockEvent`

GetStxLockEvent returns the StxLockEvent field if non-nil, zero value otherwise.

### GetStxLockEventOk

`func (o *TokenTransferTransaction1EventsInner) GetStxLockEventOk() (*StxLockTransactionEventAllOfStxLockEvent, bool)`

GetStxLockEventOk returns a tuple with the StxLockEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStxLockEvent

`func (o *TokenTransferTransaction1EventsInner) SetStxLockEvent(v StxLockTransactionEventAllOfStxLockEvent)`

SetStxLockEvent sets StxLockEvent field to given value.


### GetAsset

`func (o *TokenTransferTransaction1EventsInner) GetAsset() NonFungibleTokenAssetTransactionEvent1AllOfAsset`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *TokenTransferTransaction1EventsInner) GetAssetOk() (*NonFungibleTokenAssetTransactionEvent1AllOfAsset, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *TokenTransferTransaction1EventsInner) SetAsset(v NonFungibleTokenAssetTransactionEvent1AllOfAsset)`

SetAsset sets Asset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


