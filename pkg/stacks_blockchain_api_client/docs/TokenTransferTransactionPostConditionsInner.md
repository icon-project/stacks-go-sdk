# TokenTransferTransactionPostConditionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Principal** | [**TokenTransferTransactionPostConditionsInnerAnyOfPrincipal**](TokenTransferTransactionPostConditionsInnerAnyOfPrincipal.md) |  | 
**ConditionCode** | [**TokenTransferTransactionPostConditionsInnerAnyOf2ConditionCode**](TokenTransferTransactionPostConditionsInnerAnyOf2ConditionCode.md) |  | 
**Amount** | **string** |  | 
**Type** | **string** |  | 
**Asset** | [**TokenTransferTransactionPostConditionsInnerAnyOf1Asset**](TokenTransferTransactionPostConditionsInnerAnyOf1Asset.md) |  | 
**AssetValue** | [**TokenTransferTransactionPostConditionsInnerAnyOf2AssetValue**](TokenTransferTransactionPostConditionsInnerAnyOf2AssetValue.md) |  | 

## Methods

### NewTokenTransferTransactionPostConditionsInner

`func NewTokenTransferTransactionPostConditionsInner(principal TokenTransferTransactionPostConditionsInnerAnyOfPrincipal, conditionCode TokenTransferTransactionPostConditionsInnerAnyOf2ConditionCode, amount string, type_ string, asset TokenTransferTransactionPostConditionsInnerAnyOf1Asset, assetValue TokenTransferTransactionPostConditionsInnerAnyOf2AssetValue, ) *TokenTransferTransactionPostConditionsInner`

NewTokenTransferTransactionPostConditionsInner instantiates a new TokenTransferTransactionPostConditionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenTransferTransactionPostConditionsInnerWithDefaults

`func NewTokenTransferTransactionPostConditionsInnerWithDefaults() *TokenTransferTransactionPostConditionsInner`

NewTokenTransferTransactionPostConditionsInnerWithDefaults instantiates a new TokenTransferTransactionPostConditionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipal

`func (o *TokenTransferTransactionPostConditionsInner) GetPrincipal() TokenTransferTransactionPostConditionsInnerAnyOfPrincipal`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *TokenTransferTransactionPostConditionsInner) GetPrincipalOk() (*TokenTransferTransactionPostConditionsInnerAnyOfPrincipal, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *TokenTransferTransactionPostConditionsInner) SetPrincipal(v TokenTransferTransactionPostConditionsInnerAnyOfPrincipal)`

SetPrincipal sets Principal field to given value.


### GetConditionCode

`func (o *TokenTransferTransactionPostConditionsInner) GetConditionCode() TokenTransferTransactionPostConditionsInnerAnyOf2ConditionCode`

GetConditionCode returns the ConditionCode field if non-nil, zero value otherwise.

### GetConditionCodeOk

`func (o *TokenTransferTransactionPostConditionsInner) GetConditionCodeOk() (*TokenTransferTransactionPostConditionsInnerAnyOf2ConditionCode, bool)`

GetConditionCodeOk returns a tuple with the ConditionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionCode

`func (o *TokenTransferTransactionPostConditionsInner) SetConditionCode(v TokenTransferTransactionPostConditionsInnerAnyOf2ConditionCode)`

SetConditionCode sets ConditionCode field to given value.


### GetAmount

`func (o *TokenTransferTransactionPostConditionsInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TokenTransferTransactionPostConditionsInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TokenTransferTransactionPostConditionsInner) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetType

`func (o *TokenTransferTransactionPostConditionsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenTransferTransactionPostConditionsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenTransferTransactionPostConditionsInner) SetType(v string)`

SetType sets Type field to given value.


### GetAsset

`func (o *TokenTransferTransactionPostConditionsInner) GetAsset() TokenTransferTransactionPostConditionsInnerAnyOf1Asset`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *TokenTransferTransactionPostConditionsInner) GetAssetOk() (*TokenTransferTransactionPostConditionsInnerAnyOf1Asset, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *TokenTransferTransactionPostConditionsInner) SetAsset(v TokenTransferTransactionPostConditionsInnerAnyOf1Asset)`

SetAsset sets Asset field to given value.


### GetAssetValue

`func (o *TokenTransferTransactionPostConditionsInner) GetAssetValue() TokenTransferTransactionPostConditionsInnerAnyOf2AssetValue`

GetAssetValue returns the AssetValue field if non-nil, zero value otherwise.

### GetAssetValueOk

`func (o *TokenTransferTransactionPostConditionsInner) GetAssetValueOk() (*TokenTransferTransactionPostConditionsInnerAnyOf2AssetValue, bool)`

GetAssetValueOk returns a tuple with the AssetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetValue

`func (o *TokenTransferTransactionPostConditionsInner) SetAssetValue(v TokenTransferTransactionPostConditionsInnerAnyOf2AssetValue)`

SetAssetValue sets AssetValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


