# NonFungibleTokenAssetTransactionEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventIndex** | **int32** |  | 
**EventType** | **string** |  | 
**TxId** | **string** |  | 
**Asset** | [**NonFungibleTokenAssetTransactionEventAllOfAsset**](NonFungibleTokenAssetTransactionEventAllOfAsset.md) |  | 

## Methods

### NewNonFungibleTokenAssetTransactionEvent

`func NewNonFungibleTokenAssetTransactionEvent(eventIndex int32, eventType string, txId string, asset NonFungibleTokenAssetTransactionEventAllOfAsset, ) *NonFungibleTokenAssetTransactionEvent`

NewNonFungibleTokenAssetTransactionEvent instantiates a new NonFungibleTokenAssetTransactionEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonFungibleTokenAssetTransactionEventWithDefaults

`func NewNonFungibleTokenAssetTransactionEventWithDefaults() *NonFungibleTokenAssetTransactionEvent`

NewNonFungibleTokenAssetTransactionEventWithDefaults instantiates a new NonFungibleTokenAssetTransactionEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventIndex

`func (o *NonFungibleTokenAssetTransactionEvent) GetEventIndex() int32`

GetEventIndex returns the EventIndex field if non-nil, zero value otherwise.

### GetEventIndexOk

`func (o *NonFungibleTokenAssetTransactionEvent) GetEventIndexOk() (*int32, bool)`

GetEventIndexOk returns a tuple with the EventIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIndex

`func (o *NonFungibleTokenAssetTransactionEvent) SetEventIndex(v int32)`

SetEventIndex sets EventIndex field to given value.


### GetEventType

`func (o *NonFungibleTokenAssetTransactionEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *NonFungibleTokenAssetTransactionEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *NonFungibleTokenAssetTransactionEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetTxId

`func (o *NonFungibleTokenAssetTransactionEvent) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *NonFungibleTokenAssetTransactionEvent) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *NonFungibleTokenAssetTransactionEvent) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetAsset

`func (o *NonFungibleTokenAssetTransactionEvent) GetAsset() NonFungibleTokenAssetTransactionEventAllOfAsset`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *NonFungibleTokenAssetTransactionEvent) GetAssetOk() (*NonFungibleTokenAssetTransactionEventAllOfAsset, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *NonFungibleTokenAssetTransactionEvent) SetAsset(v NonFungibleTokenAssetTransactionEventAllOfAsset)`

SetAsset sets Asset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


