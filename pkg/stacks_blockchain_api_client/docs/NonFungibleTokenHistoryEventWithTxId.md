# NonFungibleTokenHistoryEventWithTxId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sender** | Pointer to **NullableString** |  | [optional] 
**Recipient** | Pointer to **string** |  | [optional] 
**EventIndex** | **int32** |  | 
**AssetEventType** | **string** |  | 
**TxId** | **string** |  | 

## Methods

### NewNonFungibleTokenHistoryEventWithTxId

`func NewNonFungibleTokenHistoryEventWithTxId(eventIndex int32, assetEventType string, txId string, ) *NonFungibleTokenHistoryEventWithTxId`

NewNonFungibleTokenHistoryEventWithTxId instantiates a new NonFungibleTokenHistoryEventWithTxId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonFungibleTokenHistoryEventWithTxIdWithDefaults

`func NewNonFungibleTokenHistoryEventWithTxIdWithDefaults() *NonFungibleTokenHistoryEventWithTxId`

NewNonFungibleTokenHistoryEventWithTxIdWithDefaults instantiates a new NonFungibleTokenHistoryEventWithTxId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSender

`func (o *NonFungibleTokenHistoryEventWithTxId) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *NonFungibleTokenHistoryEventWithTxId) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *NonFungibleTokenHistoryEventWithTxId) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *NonFungibleTokenHistoryEventWithTxId) HasSender() bool`

HasSender returns a boolean if a field has been set.

### SetSenderNil

`func (o *NonFungibleTokenHistoryEventWithTxId) SetSenderNil(b bool)`

 SetSenderNil sets the value for Sender to be an explicit nil

### UnsetSender
`func (o *NonFungibleTokenHistoryEventWithTxId) UnsetSender()`

UnsetSender ensures that no value is present for Sender, not even an explicit nil
### GetRecipient

`func (o *NonFungibleTokenHistoryEventWithTxId) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *NonFungibleTokenHistoryEventWithTxId) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *NonFungibleTokenHistoryEventWithTxId) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *NonFungibleTokenHistoryEventWithTxId) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetEventIndex

`func (o *NonFungibleTokenHistoryEventWithTxId) GetEventIndex() int32`

GetEventIndex returns the EventIndex field if non-nil, zero value otherwise.

### GetEventIndexOk

`func (o *NonFungibleTokenHistoryEventWithTxId) GetEventIndexOk() (*int32, bool)`

GetEventIndexOk returns a tuple with the EventIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIndex

`func (o *NonFungibleTokenHistoryEventWithTxId) SetEventIndex(v int32)`

SetEventIndex sets EventIndex field to given value.


### GetAssetEventType

`func (o *NonFungibleTokenHistoryEventWithTxId) GetAssetEventType() string`

GetAssetEventType returns the AssetEventType field if non-nil, zero value otherwise.

### GetAssetEventTypeOk

`func (o *NonFungibleTokenHistoryEventWithTxId) GetAssetEventTypeOk() (*string, bool)`

GetAssetEventTypeOk returns a tuple with the AssetEventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetEventType

`func (o *NonFungibleTokenHistoryEventWithTxId) SetAssetEventType(v string)`

SetAssetEventType sets AssetEventType field to given value.


### GetTxId

`func (o *NonFungibleTokenHistoryEventWithTxId) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *NonFungibleTokenHistoryEventWithTxId) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *NonFungibleTokenHistoryEventWithTxId) SetTxId(v string)`

SetTxId sets TxId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


