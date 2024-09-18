# NonFungibleTokenAssetTransactionEventAllOfAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetEventType** | [**StxAssetTransactionEventAllOfAssetAssetEventType**](StxAssetTransactionEventAllOfAssetAssetEventType.md) |  | 
**AssetId** | **string** |  | 
**Sender** | **string** |  | 
**Recipient** | **string** |  | 
**Value** | [**TokenTransferTransactionPostConditionsInnerAnyOf2AssetValue**](TokenTransferTransactionPostConditionsInnerAnyOf2AssetValue.md) |  | 

## Methods

### NewNonFungibleTokenAssetTransactionEventAllOfAsset

`func NewNonFungibleTokenAssetTransactionEventAllOfAsset(assetEventType StxAssetTransactionEventAllOfAssetAssetEventType, assetId string, sender string, recipient string, value TokenTransferTransactionPostConditionsInnerAnyOf2AssetValue, ) *NonFungibleTokenAssetTransactionEventAllOfAsset`

NewNonFungibleTokenAssetTransactionEventAllOfAsset instantiates a new NonFungibleTokenAssetTransactionEventAllOfAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonFungibleTokenAssetTransactionEventAllOfAssetWithDefaults

`func NewNonFungibleTokenAssetTransactionEventAllOfAssetWithDefaults() *NonFungibleTokenAssetTransactionEventAllOfAsset`

NewNonFungibleTokenAssetTransactionEventAllOfAssetWithDefaults instantiates a new NonFungibleTokenAssetTransactionEventAllOfAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetEventType

`func (o *NonFungibleTokenAssetTransactionEventAllOfAsset) GetAssetEventType() StxAssetTransactionEventAllOfAssetAssetEventType`

GetAssetEventType returns the AssetEventType field if non-nil, zero value otherwise.

### GetAssetEventTypeOk

`func (o *NonFungibleTokenAssetTransactionEventAllOfAsset) GetAssetEventTypeOk() (*StxAssetTransactionEventAllOfAssetAssetEventType, bool)`

GetAssetEventTypeOk returns a tuple with the AssetEventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetEventType

`func (o *NonFungibleTokenAssetTransactionEventAllOfAsset) SetAssetEventType(v StxAssetTransactionEventAllOfAssetAssetEventType)`

SetAssetEventType sets AssetEventType field to given value.


### GetAssetId

`func (o *NonFungibleTokenAssetTransactionEventAllOfAsset) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *NonFungibleTokenAssetTransactionEventAllOfAsset) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *NonFungibleTokenAssetTransactionEventAllOfAsset) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetSender

`func (o *NonFungibleTokenAssetTransactionEventAllOfAsset) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *NonFungibleTokenAssetTransactionEventAllOfAsset) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *NonFungibleTokenAssetTransactionEventAllOfAsset) SetSender(v string)`

SetSender sets Sender field to given value.


### GetRecipient

`func (o *NonFungibleTokenAssetTransactionEventAllOfAsset) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *NonFungibleTokenAssetTransactionEventAllOfAsset) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *NonFungibleTokenAssetTransactionEventAllOfAsset) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetValue

`func (o *NonFungibleTokenAssetTransactionEventAllOfAsset) GetValue() TokenTransferTransactionPostConditionsInnerAnyOf2AssetValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NonFungibleTokenAssetTransactionEventAllOfAsset) GetValueOk() (*TokenTransferTransactionPostConditionsInnerAnyOf2AssetValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NonFungibleTokenAssetTransactionEventAllOfAsset) SetValue(v TokenTransferTransactionPostConditionsInnerAnyOf2AssetValue)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


