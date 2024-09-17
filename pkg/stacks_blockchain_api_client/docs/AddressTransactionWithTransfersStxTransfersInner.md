# AddressTransactionWithTransfersStxTransfersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Amount transferred in micro-STX as an integer string. | 
**Sender** | Pointer to **string** | Principal that sent STX. This is unspecified if the STX were minted. | [optional] 
**Recipient** | Pointer to **string** | Principal that received STX. This is unspecified if the STX were burned. | [optional] 

## Methods

### NewAddressTransactionWithTransfersStxTransfersInner

`func NewAddressTransactionWithTransfersStxTransfersInner(amount string, ) *AddressTransactionWithTransfersStxTransfersInner`

NewAddressTransactionWithTransfersStxTransfersInner instantiates a new AddressTransactionWithTransfersStxTransfersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressTransactionWithTransfersStxTransfersInnerWithDefaults

`func NewAddressTransactionWithTransfersStxTransfersInnerWithDefaults() *AddressTransactionWithTransfersStxTransfersInner`

NewAddressTransactionWithTransfersStxTransfersInnerWithDefaults instantiates a new AddressTransactionWithTransfersStxTransfersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *AddressTransactionWithTransfersStxTransfersInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AddressTransactionWithTransfersStxTransfersInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AddressTransactionWithTransfersStxTransfersInner) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetSender

`func (o *AddressTransactionWithTransfersStxTransfersInner) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *AddressTransactionWithTransfersStxTransfersInner) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *AddressTransactionWithTransfersStxTransfersInner) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *AddressTransactionWithTransfersStxTransfersInner) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetRecipient

`func (o *AddressTransactionWithTransfersStxTransfersInner) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *AddressTransactionWithTransfersStxTransfersInner) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *AddressTransactionWithTransfersStxTransfersInner) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *AddressTransactionWithTransfersStxTransfersInner) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


