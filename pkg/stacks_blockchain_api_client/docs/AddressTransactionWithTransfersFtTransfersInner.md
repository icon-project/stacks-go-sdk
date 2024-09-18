# AddressTransactionWithTransfersFtTransfersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Amount transferred as an integer string. This balance does not factor in possible SIP-010 decimals. | 
**AssetIdentifier** | **string** | Fungible Token asset identifier. | 
**Sender** | Pointer to **string** | Principal that sent the asset. | [optional] 
**Recipient** | Pointer to **string** | Principal that received the asset. | [optional] 

## Methods

### NewAddressTransactionWithTransfersFtTransfersInner

`func NewAddressTransactionWithTransfersFtTransfersInner(amount string, assetIdentifier string, ) *AddressTransactionWithTransfersFtTransfersInner`

NewAddressTransactionWithTransfersFtTransfersInner instantiates a new AddressTransactionWithTransfersFtTransfersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressTransactionWithTransfersFtTransfersInnerWithDefaults

`func NewAddressTransactionWithTransfersFtTransfersInnerWithDefaults() *AddressTransactionWithTransfersFtTransfersInner`

NewAddressTransactionWithTransfersFtTransfersInnerWithDefaults instantiates a new AddressTransactionWithTransfersFtTransfersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *AddressTransactionWithTransfersFtTransfersInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AddressTransactionWithTransfersFtTransfersInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AddressTransactionWithTransfersFtTransfersInner) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetAssetIdentifier

`func (o *AddressTransactionWithTransfersFtTransfersInner) GetAssetIdentifier() string`

GetAssetIdentifier returns the AssetIdentifier field if non-nil, zero value otherwise.

### GetAssetIdentifierOk

`func (o *AddressTransactionWithTransfersFtTransfersInner) GetAssetIdentifierOk() (*string, bool)`

GetAssetIdentifierOk returns a tuple with the AssetIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetIdentifier

`func (o *AddressTransactionWithTransfersFtTransfersInner) SetAssetIdentifier(v string)`

SetAssetIdentifier sets AssetIdentifier field to given value.


### GetSender

`func (o *AddressTransactionWithTransfersFtTransfersInner) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *AddressTransactionWithTransfersFtTransfersInner) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *AddressTransactionWithTransfersFtTransfersInner) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *AddressTransactionWithTransfersFtTransfersInner) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetRecipient

`func (o *AddressTransactionWithTransfersFtTransfersInner) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *AddressTransactionWithTransfersFtTransfersInner) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *AddressTransactionWithTransfersFtTransfersInner) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *AddressTransactionWithTransfersFtTransfersInner) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


