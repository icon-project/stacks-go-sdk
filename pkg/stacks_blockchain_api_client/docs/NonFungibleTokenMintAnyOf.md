# NonFungibleTokenMintAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recipient** | Pointer to **string** |  | [optional] 
**EventIndex** | **int32** |  | 
**Value** | [**NonFungibleTokenHoldingWithTxIdValue**](NonFungibleTokenHoldingWithTxIdValue.md) |  | 
**Tx** | [**GetTransactionList200ResponseResultsInner**](GetTransactionList200ResponseResultsInner.md) |  | 

## Methods

### NewNonFungibleTokenMintAnyOf

`func NewNonFungibleTokenMintAnyOf(eventIndex int32, value NonFungibleTokenHoldingWithTxIdValue, tx GetTransactionList200ResponseResultsInner, ) *NonFungibleTokenMintAnyOf`

NewNonFungibleTokenMintAnyOf instantiates a new NonFungibleTokenMintAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonFungibleTokenMintAnyOfWithDefaults

`func NewNonFungibleTokenMintAnyOfWithDefaults() *NonFungibleTokenMintAnyOf`

NewNonFungibleTokenMintAnyOfWithDefaults instantiates a new NonFungibleTokenMintAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipient

`func (o *NonFungibleTokenMintAnyOf) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *NonFungibleTokenMintAnyOf) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *NonFungibleTokenMintAnyOf) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *NonFungibleTokenMintAnyOf) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetEventIndex

`func (o *NonFungibleTokenMintAnyOf) GetEventIndex() int32`

GetEventIndex returns the EventIndex field if non-nil, zero value otherwise.

### GetEventIndexOk

`func (o *NonFungibleTokenMintAnyOf) GetEventIndexOk() (*int32, bool)`

GetEventIndexOk returns a tuple with the EventIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIndex

`func (o *NonFungibleTokenMintAnyOf) SetEventIndex(v int32)`

SetEventIndex sets EventIndex field to given value.


### GetValue

`func (o *NonFungibleTokenMintAnyOf) GetValue() NonFungibleTokenHoldingWithTxIdValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NonFungibleTokenMintAnyOf) GetValueOk() (*NonFungibleTokenHoldingWithTxIdValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NonFungibleTokenMintAnyOf) SetValue(v NonFungibleTokenHoldingWithTxIdValue)`

SetValue sets Value field to given value.


### GetTx

`func (o *NonFungibleTokenMintAnyOf) GetTx() GetTransactionList200ResponseResultsInner`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *NonFungibleTokenMintAnyOf) GetTxOk() (*GetTransactionList200ResponseResultsInner, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *NonFungibleTokenMintAnyOf) SetTx(v GetTransactionList200ResponseResultsInner)`

SetTx sets Tx field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


