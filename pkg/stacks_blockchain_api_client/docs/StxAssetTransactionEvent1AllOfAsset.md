# StxAssetTransactionEvent1AllOfAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetEventType** | [**StxAssetTransactionEvent1AllOfAssetAssetEventType**](StxAssetTransactionEvent1AllOfAssetAssetEventType.md) |  | 
**Sender** | **string** |  | 
**Recipient** | **string** |  | 
**Amount** | **string** |  | 
**Memo** | Pointer to **string** |  | [optional] 

## Methods

### NewStxAssetTransactionEvent1AllOfAsset

`func NewStxAssetTransactionEvent1AllOfAsset(assetEventType StxAssetTransactionEvent1AllOfAssetAssetEventType, sender string, recipient string, amount string, ) *StxAssetTransactionEvent1AllOfAsset`

NewStxAssetTransactionEvent1AllOfAsset instantiates a new StxAssetTransactionEvent1AllOfAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStxAssetTransactionEvent1AllOfAssetWithDefaults

`func NewStxAssetTransactionEvent1AllOfAssetWithDefaults() *StxAssetTransactionEvent1AllOfAsset`

NewStxAssetTransactionEvent1AllOfAssetWithDefaults instantiates a new StxAssetTransactionEvent1AllOfAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetEventType

`func (o *StxAssetTransactionEvent1AllOfAsset) GetAssetEventType() StxAssetTransactionEvent1AllOfAssetAssetEventType`

GetAssetEventType returns the AssetEventType field if non-nil, zero value otherwise.

### GetAssetEventTypeOk

`func (o *StxAssetTransactionEvent1AllOfAsset) GetAssetEventTypeOk() (*StxAssetTransactionEvent1AllOfAssetAssetEventType, bool)`

GetAssetEventTypeOk returns a tuple with the AssetEventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetEventType

`func (o *StxAssetTransactionEvent1AllOfAsset) SetAssetEventType(v StxAssetTransactionEvent1AllOfAssetAssetEventType)`

SetAssetEventType sets AssetEventType field to given value.


### GetSender

`func (o *StxAssetTransactionEvent1AllOfAsset) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *StxAssetTransactionEvent1AllOfAsset) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *StxAssetTransactionEvent1AllOfAsset) SetSender(v string)`

SetSender sets Sender field to given value.


### GetRecipient

`func (o *StxAssetTransactionEvent1AllOfAsset) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *StxAssetTransactionEvent1AllOfAsset) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *StxAssetTransactionEvent1AllOfAsset) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetAmount

`func (o *StxAssetTransactionEvent1AllOfAsset) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *StxAssetTransactionEvent1AllOfAsset) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *StxAssetTransactionEvent1AllOfAsset) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetMemo

`func (o *StxAssetTransactionEvent1AllOfAsset) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *StxAssetTransactionEvent1AllOfAsset) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *StxAssetTransactionEvent1AllOfAsset) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *StxAssetTransactionEvent1AllOfAsset) HasMemo() bool`

HasMemo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


