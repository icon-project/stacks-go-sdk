# TokenTransferTransaction1PostConditionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Principal** | [**TokenTransferTransaction1PostConditionsInnerAnyOfPrincipal**](TokenTransferTransaction1PostConditionsInnerAnyOfPrincipal.md) |  | 
**ConditionCode** | [**TokenTransferTransaction1PostConditionMode**](TokenTransferTransaction1PostConditionMode.md) |  | 
**Amount** | **string** |  | 
**Type** | **string** |  | 
**Asset** | [**TokenTransferTransactionPostConditionsInnerAnyOf1Asset**](TokenTransferTransactionPostConditionsInnerAnyOf1Asset.md) |  | 
**AssetValue** | [**TokenTransferTransactionPostConditionsInnerAnyOf2AssetValue**](TokenTransferTransactionPostConditionsInnerAnyOf2AssetValue.md) |  | 

## Methods

### NewTokenTransferTransaction1PostConditionsInner

`func NewTokenTransferTransaction1PostConditionsInner(principal TokenTransferTransaction1PostConditionsInnerAnyOfPrincipal, conditionCode TokenTransferTransaction1PostConditionMode, amount string, type_ string, asset TokenTransferTransactionPostConditionsInnerAnyOf1Asset, assetValue TokenTransferTransactionPostConditionsInnerAnyOf2AssetValue, ) *TokenTransferTransaction1PostConditionsInner`

NewTokenTransferTransaction1PostConditionsInner instantiates a new TokenTransferTransaction1PostConditionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenTransferTransaction1PostConditionsInnerWithDefaults

`func NewTokenTransferTransaction1PostConditionsInnerWithDefaults() *TokenTransferTransaction1PostConditionsInner`

NewTokenTransferTransaction1PostConditionsInnerWithDefaults instantiates a new TokenTransferTransaction1PostConditionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipal

`func (o *TokenTransferTransaction1PostConditionsInner) GetPrincipal() TokenTransferTransaction1PostConditionsInnerAnyOfPrincipal`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *TokenTransferTransaction1PostConditionsInner) GetPrincipalOk() (*TokenTransferTransaction1PostConditionsInnerAnyOfPrincipal, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *TokenTransferTransaction1PostConditionsInner) SetPrincipal(v TokenTransferTransaction1PostConditionsInnerAnyOfPrincipal)`

SetPrincipal sets Principal field to given value.


### GetConditionCode

`func (o *TokenTransferTransaction1PostConditionsInner) GetConditionCode() TokenTransferTransaction1PostConditionMode`

GetConditionCode returns the ConditionCode field if non-nil, zero value otherwise.

### GetConditionCodeOk

`func (o *TokenTransferTransaction1PostConditionsInner) GetConditionCodeOk() (*TokenTransferTransaction1PostConditionMode, bool)`

GetConditionCodeOk returns a tuple with the ConditionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionCode

`func (o *TokenTransferTransaction1PostConditionsInner) SetConditionCode(v TokenTransferTransaction1PostConditionMode)`

SetConditionCode sets ConditionCode field to given value.


### GetAmount

`func (o *TokenTransferTransaction1PostConditionsInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TokenTransferTransaction1PostConditionsInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TokenTransferTransaction1PostConditionsInner) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetType

`func (o *TokenTransferTransaction1PostConditionsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenTransferTransaction1PostConditionsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenTransferTransaction1PostConditionsInner) SetType(v string)`

SetType sets Type field to given value.


### GetAsset

`func (o *TokenTransferTransaction1PostConditionsInner) GetAsset() TokenTransferTransactionPostConditionsInnerAnyOf1Asset`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *TokenTransferTransaction1PostConditionsInner) GetAssetOk() (*TokenTransferTransactionPostConditionsInnerAnyOf1Asset, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *TokenTransferTransaction1PostConditionsInner) SetAsset(v TokenTransferTransactionPostConditionsInnerAnyOf1Asset)`

SetAsset sets Asset field to given value.


### GetAssetValue

`func (o *TokenTransferTransaction1PostConditionsInner) GetAssetValue() TokenTransferTransactionPostConditionsInnerAnyOf2AssetValue`

GetAssetValue returns the AssetValue field if non-nil, zero value otherwise.

### GetAssetValueOk

`func (o *TokenTransferTransaction1PostConditionsInner) GetAssetValueOk() (*TokenTransferTransactionPostConditionsInnerAnyOf2AssetValue, bool)`

GetAssetValueOk returns a tuple with the AssetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetValue

`func (o *TokenTransferTransaction1PostConditionsInner) SetAssetValue(v TokenTransferTransactionPostConditionsInnerAnyOf2AssetValue)`

SetAssetValue sets AssetValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


