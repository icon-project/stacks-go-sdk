# AddressTransactionEventAnyOfData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**StxAssetTransactionEventAllOfAssetAssetEventType**](StxAssetTransactionEventAllOfAssetAssetEventType.md) |  | 
**Amount** | **string** | Amount transferred in micro-STX as an integer string. | 
**Sender** | Pointer to **string** | Principal that sent STX. This is unspecified if the STX were minted. | [optional] 
**Recipient** | Pointer to **string** | Principal that received STX. This is unspecified if the STX were burned. | [optional] 

## Methods

### NewAddressTransactionEventAnyOfData

`func NewAddressTransactionEventAnyOfData(type_ StxAssetTransactionEventAllOfAssetAssetEventType, amount string, ) *AddressTransactionEventAnyOfData`

NewAddressTransactionEventAnyOfData instantiates a new AddressTransactionEventAnyOfData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressTransactionEventAnyOfDataWithDefaults

`func NewAddressTransactionEventAnyOfDataWithDefaults() *AddressTransactionEventAnyOfData`

NewAddressTransactionEventAnyOfDataWithDefaults instantiates a new AddressTransactionEventAnyOfData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AddressTransactionEventAnyOfData) GetType() StxAssetTransactionEventAllOfAssetAssetEventType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddressTransactionEventAnyOfData) GetTypeOk() (*StxAssetTransactionEventAllOfAssetAssetEventType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddressTransactionEventAnyOfData) SetType(v StxAssetTransactionEventAllOfAssetAssetEventType)`

SetType sets Type field to given value.


### GetAmount

`func (o *AddressTransactionEventAnyOfData) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AddressTransactionEventAnyOfData) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AddressTransactionEventAnyOfData) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetSender

`func (o *AddressTransactionEventAnyOfData) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *AddressTransactionEventAnyOfData) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *AddressTransactionEventAnyOfData) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *AddressTransactionEventAnyOfData) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetRecipient

`func (o *AddressTransactionEventAnyOfData) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *AddressTransactionEventAnyOfData) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *AddressTransactionEventAnyOfData) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *AddressTransactionEventAnyOfData) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


