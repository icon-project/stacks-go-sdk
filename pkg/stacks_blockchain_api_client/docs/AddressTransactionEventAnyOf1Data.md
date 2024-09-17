# AddressTransactionEventAnyOf1Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**StxAssetTransactionEventAllOfAssetAssetEventType**](StxAssetTransactionEventAllOfAssetAssetEventType.md) |  | 
**Amount** | **string** | Amount transferred as an integer string. This balance does not factor in possible SIP-010 decimals. | 
**AssetIdentifier** | **string** | Fungible Token asset identifier. | 
**Sender** | Pointer to **string** | Principal that sent the asset. | [optional] 
**Recipient** | Pointer to **string** | Principal that received the asset. | [optional] 

## Methods

### NewAddressTransactionEventAnyOf1Data

`func NewAddressTransactionEventAnyOf1Data(type_ StxAssetTransactionEventAllOfAssetAssetEventType, amount string, assetIdentifier string, ) *AddressTransactionEventAnyOf1Data`

NewAddressTransactionEventAnyOf1Data instantiates a new AddressTransactionEventAnyOf1Data object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressTransactionEventAnyOf1DataWithDefaults

`func NewAddressTransactionEventAnyOf1DataWithDefaults() *AddressTransactionEventAnyOf1Data`

NewAddressTransactionEventAnyOf1DataWithDefaults instantiates a new AddressTransactionEventAnyOf1Data object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AddressTransactionEventAnyOf1Data) GetType() StxAssetTransactionEventAllOfAssetAssetEventType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddressTransactionEventAnyOf1Data) GetTypeOk() (*StxAssetTransactionEventAllOfAssetAssetEventType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddressTransactionEventAnyOf1Data) SetType(v StxAssetTransactionEventAllOfAssetAssetEventType)`

SetType sets Type field to given value.


### GetAmount

`func (o *AddressTransactionEventAnyOf1Data) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AddressTransactionEventAnyOf1Data) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AddressTransactionEventAnyOf1Data) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetAssetIdentifier

`func (o *AddressTransactionEventAnyOf1Data) GetAssetIdentifier() string`

GetAssetIdentifier returns the AssetIdentifier field if non-nil, zero value otherwise.

### GetAssetIdentifierOk

`func (o *AddressTransactionEventAnyOf1Data) GetAssetIdentifierOk() (*string, bool)`

GetAssetIdentifierOk returns a tuple with the AssetIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetIdentifier

`func (o *AddressTransactionEventAnyOf1Data) SetAssetIdentifier(v string)`

SetAssetIdentifier sets AssetIdentifier field to given value.


### GetSender

`func (o *AddressTransactionEventAnyOf1Data) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *AddressTransactionEventAnyOf1Data) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *AddressTransactionEventAnyOf1Data) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *AddressTransactionEventAnyOf1Data) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetRecipient

`func (o *AddressTransactionEventAnyOf1Data) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *AddressTransactionEventAnyOf1Data) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *AddressTransactionEventAnyOf1Data) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *AddressTransactionEventAnyOf1Data) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


