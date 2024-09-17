# NonFungibleTokenMint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recipient** | Pointer to **string** |  | [optional] 
**EventIndex** | **int32** |  | 
**Value** | [**NonFungibleTokenHoldingWithTxIdValue**](NonFungibleTokenHoldingWithTxIdValue.md) |  | 
**TxId** | **string** |  | 
**Tx** | [**GetTransactionList200ResponseResultsInner**](GetTransactionList200ResponseResultsInner.md) |  | 

## Methods

### NewNonFungibleTokenMint

`func NewNonFungibleTokenMint(eventIndex int32, value NonFungibleTokenHoldingWithTxIdValue, txId string, tx GetTransactionList200ResponseResultsInner, ) *NonFungibleTokenMint`

NewNonFungibleTokenMint instantiates a new NonFungibleTokenMint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonFungibleTokenMintWithDefaults

`func NewNonFungibleTokenMintWithDefaults() *NonFungibleTokenMint`

NewNonFungibleTokenMintWithDefaults instantiates a new NonFungibleTokenMint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipient

`func (o *NonFungibleTokenMint) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *NonFungibleTokenMint) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *NonFungibleTokenMint) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *NonFungibleTokenMint) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetEventIndex

`func (o *NonFungibleTokenMint) GetEventIndex() int32`

GetEventIndex returns the EventIndex field if non-nil, zero value otherwise.

### GetEventIndexOk

`func (o *NonFungibleTokenMint) GetEventIndexOk() (*int32, bool)`

GetEventIndexOk returns a tuple with the EventIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIndex

`func (o *NonFungibleTokenMint) SetEventIndex(v int32)`

SetEventIndex sets EventIndex field to given value.


### GetValue

`func (o *NonFungibleTokenMint) GetValue() NonFungibleTokenHoldingWithTxIdValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NonFungibleTokenMint) GetValueOk() (*NonFungibleTokenHoldingWithTxIdValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NonFungibleTokenMint) SetValue(v NonFungibleTokenHoldingWithTxIdValue)`

SetValue sets Value field to given value.


### GetTxId

`func (o *NonFungibleTokenMint) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *NonFungibleTokenMint) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *NonFungibleTokenMint) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetTx

`func (o *NonFungibleTokenMint) GetTx() GetTransactionList200ResponseResultsInner`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *NonFungibleTokenMint) GetTxOk() (*GetTransactionList200ResponseResultsInner, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *NonFungibleTokenMint) SetTx(v GetTransactionList200ResponseResultsInner)`

SetTx sets Tx field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


