# NonFungibleTokenMintWithTxId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recipient** | Pointer to **string** |  | [optional] 
**EventIndex** | **int32** |  | 
**Value** | [**NonFungibleTokenHoldingWithTxIdValue**](NonFungibleTokenHoldingWithTxIdValue.md) |  | 
**TxId** | **string** |  | 

## Methods

### NewNonFungibleTokenMintWithTxId

`func NewNonFungibleTokenMintWithTxId(eventIndex int32, value NonFungibleTokenHoldingWithTxIdValue, txId string, ) *NonFungibleTokenMintWithTxId`

NewNonFungibleTokenMintWithTxId instantiates a new NonFungibleTokenMintWithTxId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonFungibleTokenMintWithTxIdWithDefaults

`func NewNonFungibleTokenMintWithTxIdWithDefaults() *NonFungibleTokenMintWithTxId`

NewNonFungibleTokenMintWithTxIdWithDefaults instantiates a new NonFungibleTokenMintWithTxId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipient

`func (o *NonFungibleTokenMintWithTxId) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *NonFungibleTokenMintWithTxId) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *NonFungibleTokenMintWithTxId) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *NonFungibleTokenMintWithTxId) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetEventIndex

`func (o *NonFungibleTokenMintWithTxId) GetEventIndex() int32`

GetEventIndex returns the EventIndex field if non-nil, zero value otherwise.

### GetEventIndexOk

`func (o *NonFungibleTokenMintWithTxId) GetEventIndexOk() (*int32, bool)`

GetEventIndexOk returns a tuple with the EventIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIndex

`func (o *NonFungibleTokenMintWithTxId) SetEventIndex(v int32)`

SetEventIndex sets EventIndex field to given value.


### GetValue

`func (o *NonFungibleTokenMintWithTxId) GetValue() NonFungibleTokenHoldingWithTxIdValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NonFungibleTokenMintWithTxId) GetValueOk() (*NonFungibleTokenHoldingWithTxIdValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NonFungibleTokenMintWithTxId) SetValue(v NonFungibleTokenHoldingWithTxIdValue)`

SetValue sets Value field to given value.


### GetTxId

`func (o *NonFungibleTokenMintWithTxId) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *NonFungibleTokenMintWithTxId) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *NonFungibleTokenMintWithTxId) SetTxId(v string)`

SetTxId sets TxId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


