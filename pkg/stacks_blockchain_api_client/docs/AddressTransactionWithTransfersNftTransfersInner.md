# AddressTransactionWithTransfersNftTransfersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | [**AddressTransactionWithTransfersNftTransfersInnerValue**](AddressTransactionWithTransfersNftTransfersInnerValue.md) |  | 
**AssetIdentifier** | **string** | Non Fungible Token asset identifier. | 
**Sender** | Pointer to **string** | Principal that sent the asset. | [optional] 
**Recipient** | Pointer to **string** | Principal that received the asset. | [optional] 

## Methods

### NewAddressTransactionWithTransfersNftTransfersInner

`func NewAddressTransactionWithTransfersNftTransfersInner(value AddressTransactionWithTransfersNftTransfersInnerValue, assetIdentifier string, ) *AddressTransactionWithTransfersNftTransfersInner`

NewAddressTransactionWithTransfersNftTransfersInner instantiates a new AddressTransactionWithTransfersNftTransfersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressTransactionWithTransfersNftTransfersInnerWithDefaults

`func NewAddressTransactionWithTransfersNftTransfersInnerWithDefaults() *AddressTransactionWithTransfersNftTransfersInner`

NewAddressTransactionWithTransfersNftTransfersInnerWithDefaults instantiates a new AddressTransactionWithTransfersNftTransfersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *AddressTransactionWithTransfersNftTransfersInner) GetValue() AddressTransactionWithTransfersNftTransfersInnerValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AddressTransactionWithTransfersNftTransfersInner) GetValueOk() (*AddressTransactionWithTransfersNftTransfersInnerValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AddressTransactionWithTransfersNftTransfersInner) SetValue(v AddressTransactionWithTransfersNftTransfersInnerValue)`

SetValue sets Value field to given value.


### GetAssetIdentifier

`func (o *AddressTransactionWithTransfersNftTransfersInner) GetAssetIdentifier() string`

GetAssetIdentifier returns the AssetIdentifier field if non-nil, zero value otherwise.

### GetAssetIdentifierOk

`func (o *AddressTransactionWithTransfersNftTransfersInner) GetAssetIdentifierOk() (*string, bool)`

GetAssetIdentifierOk returns a tuple with the AssetIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetIdentifier

`func (o *AddressTransactionWithTransfersNftTransfersInner) SetAssetIdentifier(v string)`

SetAssetIdentifier sets AssetIdentifier field to given value.


### GetSender

`func (o *AddressTransactionWithTransfersNftTransfersInner) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *AddressTransactionWithTransfersNftTransfersInner) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *AddressTransactionWithTransfersNftTransfersInner) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *AddressTransactionWithTransfersNftTransfersInner) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetRecipient

`func (o *AddressTransactionWithTransfersNftTransfersInner) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *AddressTransactionWithTransfersNftTransfersInner) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *AddressTransactionWithTransfersNftTransfersInner) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *AddressTransactionWithTransfersNftTransfersInner) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


