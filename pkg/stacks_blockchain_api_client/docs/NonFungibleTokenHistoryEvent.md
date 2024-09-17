# NonFungibleTokenHistoryEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sender** | Pointer to **NullableString** |  | [optional] 
**Recipient** | Pointer to **string** |  | [optional] 
**EventIndex** | **int32** |  | 
**AssetEventType** | **string** |  | 
**TxId** | **string** |  | 
**Tx** | [**GetTransactionList200ResponseResultsInner**](GetTransactionList200ResponseResultsInner.md) |  | 

## Methods

### NewNonFungibleTokenHistoryEvent

`func NewNonFungibleTokenHistoryEvent(eventIndex int32, assetEventType string, txId string, tx GetTransactionList200ResponseResultsInner, ) *NonFungibleTokenHistoryEvent`

NewNonFungibleTokenHistoryEvent instantiates a new NonFungibleTokenHistoryEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonFungibleTokenHistoryEventWithDefaults

`func NewNonFungibleTokenHistoryEventWithDefaults() *NonFungibleTokenHistoryEvent`

NewNonFungibleTokenHistoryEventWithDefaults instantiates a new NonFungibleTokenHistoryEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSender

`func (o *NonFungibleTokenHistoryEvent) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *NonFungibleTokenHistoryEvent) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *NonFungibleTokenHistoryEvent) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *NonFungibleTokenHistoryEvent) HasSender() bool`

HasSender returns a boolean if a field has been set.

### SetSenderNil

`func (o *NonFungibleTokenHistoryEvent) SetSenderNil(b bool)`

 SetSenderNil sets the value for Sender to be an explicit nil

### UnsetSender
`func (o *NonFungibleTokenHistoryEvent) UnsetSender()`

UnsetSender ensures that no value is present for Sender, not even an explicit nil
### GetRecipient

`func (o *NonFungibleTokenHistoryEvent) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *NonFungibleTokenHistoryEvent) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *NonFungibleTokenHistoryEvent) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *NonFungibleTokenHistoryEvent) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetEventIndex

`func (o *NonFungibleTokenHistoryEvent) GetEventIndex() int32`

GetEventIndex returns the EventIndex field if non-nil, zero value otherwise.

### GetEventIndexOk

`func (o *NonFungibleTokenHistoryEvent) GetEventIndexOk() (*int32, bool)`

GetEventIndexOk returns a tuple with the EventIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIndex

`func (o *NonFungibleTokenHistoryEvent) SetEventIndex(v int32)`

SetEventIndex sets EventIndex field to given value.


### GetAssetEventType

`func (o *NonFungibleTokenHistoryEvent) GetAssetEventType() string`

GetAssetEventType returns the AssetEventType field if non-nil, zero value otherwise.

### GetAssetEventTypeOk

`func (o *NonFungibleTokenHistoryEvent) GetAssetEventTypeOk() (*string, bool)`

GetAssetEventTypeOk returns a tuple with the AssetEventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetEventType

`func (o *NonFungibleTokenHistoryEvent) SetAssetEventType(v string)`

SetAssetEventType sets AssetEventType field to given value.


### GetTxId

`func (o *NonFungibleTokenHistoryEvent) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *NonFungibleTokenHistoryEvent) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *NonFungibleTokenHistoryEvent) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetTx

`func (o *NonFungibleTokenHistoryEvent) GetTx() GetTransactionList200ResponseResultsInner`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *NonFungibleTokenHistoryEvent) GetTxOk() (*GetTransactionList200ResponseResultsInner, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *NonFungibleTokenHistoryEvent) SetTx(v GetTransactionList200ResponseResultsInner)`

SetTx sets Tx field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


