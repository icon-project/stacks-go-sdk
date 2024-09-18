# TokenTransferTransactionTokenTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecipientAddress** | **string** |  | 
**Amount** | **string** | Transfer amount as Integer string (64-bit unsigned integer) | 
**Memo** | **string** | Hex encoded arbitrary message, up to 34 bytes length (should try decoding to an ASCII string) | 

## Methods

### NewTokenTransferTransactionTokenTransfer

`func NewTokenTransferTransactionTokenTransfer(recipientAddress string, amount string, memo string, ) *TokenTransferTransactionTokenTransfer`

NewTokenTransferTransactionTokenTransfer instantiates a new TokenTransferTransactionTokenTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenTransferTransactionTokenTransferWithDefaults

`func NewTokenTransferTransactionTokenTransferWithDefaults() *TokenTransferTransactionTokenTransfer`

NewTokenTransferTransactionTokenTransferWithDefaults instantiates a new TokenTransferTransactionTokenTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipientAddress

`func (o *TokenTransferTransactionTokenTransfer) GetRecipientAddress() string`

GetRecipientAddress returns the RecipientAddress field if non-nil, zero value otherwise.

### GetRecipientAddressOk

`func (o *TokenTransferTransactionTokenTransfer) GetRecipientAddressOk() (*string, bool)`

GetRecipientAddressOk returns a tuple with the RecipientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientAddress

`func (o *TokenTransferTransactionTokenTransfer) SetRecipientAddress(v string)`

SetRecipientAddress sets RecipientAddress field to given value.


### GetAmount

`func (o *TokenTransferTransactionTokenTransfer) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TokenTransferTransactionTokenTransfer) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TokenTransferTransactionTokenTransfer) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetMemo

`func (o *TokenTransferTransactionTokenTransfer) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *TokenTransferTransactionTokenTransfer) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *TokenTransferTransactionTokenTransfer) SetMemo(v string)`

SetMemo sets Memo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


