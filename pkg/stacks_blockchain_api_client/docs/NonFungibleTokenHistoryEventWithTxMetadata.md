# NonFungibleTokenHistoryEventWithTxMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sender** | Pointer to **NullableString** |  | [optional] 
**Recipient** | Pointer to **string** |  | [optional] 
**EventIndex** | **int32** |  | 
**AssetEventType** | **string** |  | 
**Tx** | [**GetTransactionList200ResponseResultsInner**](GetTransactionList200ResponseResultsInner.md) |  | 

## Methods

### NewNonFungibleTokenHistoryEventWithTxMetadata

`func NewNonFungibleTokenHistoryEventWithTxMetadata(eventIndex int32, assetEventType string, tx GetTransactionList200ResponseResultsInner, ) *NonFungibleTokenHistoryEventWithTxMetadata`

NewNonFungibleTokenHistoryEventWithTxMetadata instantiates a new NonFungibleTokenHistoryEventWithTxMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonFungibleTokenHistoryEventWithTxMetadataWithDefaults

`func NewNonFungibleTokenHistoryEventWithTxMetadataWithDefaults() *NonFungibleTokenHistoryEventWithTxMetadata`

NewNonFungibleTokenHistoryEventWithTxMetadataWithDefaults instantiates a new NonFungibleTokenHistoryEventWithTxMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSender

`func (o *NonFungibleTokenHistoryEventWithTxMetadata) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *NonFungibleTokenHistoryEventWithTxMetadata) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *NonFungibleTokenHistoryEventWithTxMetadata) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *NonFungibleTokenHistoryEventWithTxMetadata) HasSender() bool`

HasSender returns a boolean if a field has been set.

### SetSenderNil

`func (o *NonFungibleTokenHistoryEventWithTxMetadata) SetSenderNil(b bool)`

 SetSenderNil sets the value for Sender to be an explicit nil

### UnsetSender
`func (o *NonFungibleTokenHistoryEventWithTxMetadata) UnsetSender()`

UnsetSender ensures that no value is present for Sender, not even an explicit nil
### GetRecipient

`func (o *NonFungibleTokenHistoryEventWithTxMetadata) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *NonFungibleTokenHistoryEventWithTxMetadata) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *NonFungibleTokenHistoryEventWithTxMetadata) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *NonFungibleTokenHistoryEventWithTxMetadata) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetEventIndex

`func (o *NonFungibleTokenHistoryEventWithTxMetadata) GetEventIndex() int32`

GetEventIndex returns the EventIndex field if non-nil, zero value otherwise.

### GetEventIndexOk

`func (o *NonFungibleTokenHistoryEventWithTxMetadata) GetEventIndexOk() (*int32, bool)`

GetEventIndexOk returns a tuple with the EventIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIndex

`func (o *NonFungibleTokenHistoryEventWithTxMetadata) SetEventIndex(v int32)`

SetEventIndex sets EventIndex field to given value.


### GetAssetEventType

`func (o *NonFungibleTokenHistoryEventWithTxMetadata) GetAssetEventType() string`

GetAssetEventType returns the AssetEventType field if non-nil, zero value otherwise.

### GetAssetEventTypeOk

`func (o *NonFungibleTokenHistoryEventWithTxMetadata) GetAssetEventTypeOk() (*string, bool)`

GetAssetEventTypeOk returns a tuple with the AssetEventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetEventType

`func (o *NonFungibleTokenHistoryEventWithTxMetadata) SetAssetEventType(v string)`

SetAssetEventType sets AssetEventType field to given value.


### GetTx

`func (o *NonFungibleTokenHistoryEventWithTxMetadata) GetTx() GetTransactionList200ResponseResultsInner`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *NonFungibleTokenHistoryEventWithTxMetadata) GetTxOk() (*GetTransactionList200ResponseResultsInner, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *NonFungibleTokenHistoryEventWithTxMetadata) SetTx(v GetTransactionList200ResponseResultsInner)`

SetTx sets Tx field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


