# StxAssetTransactionEventAllOfAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetEventType** | [**StxAssetTransactionEventAllOfAssetAssetEventType**](StxAssetTransactionEventAllOfAssetAssetEventType.md) |  | 
**Sender** | **string** |  | 
**Recipient** | **string** |  | 
**Amount** | **string** |  | 
**Memo** | Pointer to **string** |  | [optional] 

## Methods

### NewStxAssetTransactionEventAllOfAsset

`func NewStxAssetTransactionEventAllOfAsset(assetEventType StxAssetTransactionEventAllOfAssetAssetEventType, sender string, recipient string, amount string, ) *StxAssetTransactionEventAllOfAsset`

NewStxAssetTransactionEventAllOfAsset instantiates a new StxAssetTransactionEventAllOfAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStxAssetTransactionEventAllOfAssetWithDefaults

`func NewStxAssetTransactionEventAllOfAssetWithDefaults() *StxAssetTransactionEventAllOfAsset`

NewStxAssetTransactionEventAllOfAssetWithDefaults instantiates a new StxAssetTransactionEventAllOfAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetEventType

`func (o *StxAssetTransactionEventAllOfAsset) GetAssetEventType() StxAssetTransactionEventAllOfAssetAssetEventType`

GetAssetEventType returns the AssetEventType field if non-nil, zero value otherwise.

### GetAssetEventTypeOk

`func (o *StxAssetTransactionEventAllOfAsset) GetAssetEventTypeOk() (*StxAssetTransactionEventAllOfAssetAssetEventType, bool)`

GetAssetEventTypeOk returns a tuple with the AssetEventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetEventType

`func (o *StxAssetTransactionEventAllOfAsset) SetAssetEventType(v StxAssetTransactionEventAllOfAssetAssetEventType)`

SetAssetEventType sets AssetEventType field to given value.


### GetSender

`func (o *StxAssetTransactionEventAllOfAsset) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *StxAssetTransactionEventAllOfAsset) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *StxAssetTransactionEventAllOfAsset) SetSender(v string)`

SetSender sets Sender field to given value.


### GetRecipient

`func (o *StxAssetTransactionEventAllOfAsset) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *StxAssetTransactionEventAllOfAsset) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *StxAssetTransactionEventAllOfAsset) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetAmount

`func (o *StxAssetTransactionEventAllOfAsset) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *StxAssetTransactionEventAllOfAsset) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *StxAssetTransactionEventAllOfAsset) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetMemo

`func (o *StxAssetTransactionEventAllOfAsset) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *StxAssetTransactionEventAllOfAsset) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *StxAssetTransactionEventAllOfAsset) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *StxAssetTransactionEventAllOfAsset) HasMemo() bool`

HasMemo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


