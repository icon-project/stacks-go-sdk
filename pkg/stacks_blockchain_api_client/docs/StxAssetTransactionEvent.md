# StxAssetTransactionEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventIndex** | **int32** |  | 
**EventType** | **string** |  | 
**TxId** | **string** |  | 
**Asset** | [**StxAssetTransactionEventAllOfAsset**](StxAssetTransactionEventAllOfAsset.md) |  | 

## Methods

### NewStxAssetTransactionEvent

`func NewStxAssetTransactionEvent(eventIndex int32, eventType string, txId string, asset StxAssetTransactionEventAllOfAsset, ) *StxAssetTransactionEvent`

NewStxAssetTransactionEvent instantiates a new StxAssetTransactionEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStxAssetTransactionEventWithDefaults

`func NewStxAssetTransactionEventWithDefaults() *StxAssetTransactionEvent`

NewStxAssetTransactionEventWithDefaults instantiates a new StxAssetTransactionEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventIndex

`func (o *StxAssetTransactionEvent) GetEventIndex() int32`

GetEventIndex returns the EventIndex field if non-nil, zero value otherwise.

### GetEventIndexOk

`func (o *StxAssetTransactionEvent) GetEventIndexOk() (*int32, bool)`

GetEventIndexOk returns a tuple with the EventIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIndex

`func (o *StxAssetTransactionEvent) SetEventIndex(v int32)`

SetEventIndex sets EventIndex field to given value.


### GetEventType

`func (o *StxAssetTransactionEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *StxAssetTransactionEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *StxAssetTransactionEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetTxId

`func (o *StxAssetTransactionEvent) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *StxAssetTransactionEvent) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *StxAssetTransactionEvent) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetAsset

`func (o *StxAssetTransactionEvent) GetAsset() StxAssetTransactionEventAllOfAsset`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *StxAssetTransactionEvent) GetAssetOk() (*StxAssetTransactionEventAllOfAsset, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *StxAssetTransactionEvent) SetAsset(v StxAssetTransactionEventAllOfAsset)`

SetAsset sets Asset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


