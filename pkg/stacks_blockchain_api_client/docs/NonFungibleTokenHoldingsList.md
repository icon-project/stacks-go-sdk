# NonFungibleTokenHoldingsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetIdentifier** | **string** |  | 
**Value** | [**NonFungibleTokenHoldingWithTxIdValue**](NonFungibleTokenHoldingWithTxIdValue.md) |  | 
**BlockHeight** | **int32** |  | 
**TxId** | **string** |  | 
**Tx** | [**GetTransactionList200ResponseResultsInner**](GetTransactionList200ResponseResultsInner.md) |  | 

## Methods

### NewNonFungibleTokenHoldingsList

`func NewNonFungibleTokenHoldingsList(assetIdentifier string, value NonFungibleTokenHoldingWithTxIdValue, blockHeight int32, txId string, tx GetTransactionList200ResponseResultsInner, ) *NonFungibleTokenHoldingsList`

NewNonFungibleTokenHoldingsList instantiates a new NonFungibleTokenHoldingsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonFungibleTokenHoldingsListWithDefaults

`func NewNonFungibleTokenHoldingsListWithDefaults() *NonFungibleTokenHoldingsList`

NewNonFungibleTokenHoldingsListWithDefaults instantiates a new NonFungibleTokenHoldingsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetIdentifier

`func (o *NonFungibleTokenHoldingsList) GetAssetIdentifier() string`

GetAssetIdentifier returns the AssetIdentifier field if non-nil, zero value otherwise.

### GetAssetIdentifierOk

`func (o *NonFungibleTokenHoldingsList) GetAssetIdentifierOk() (*string, bool)`

GetAssetIdentifierOk returns a tuple with the AssetIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetIdentifier

`func (o *NonFungibleTokenHoldingsList) SetAssetIdentifier(v string)`

SetAssetIdentifier sets AssetIdentifier field to given value.


### GetValue

`func (o *NonFungibleTokenHoldingsList) GetValue() NonFungibleTokenHoldingWithTxIdValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NonFungibleTokenHoldingsList) GetValueOk() (*NonFungibleTokenHoldingWithTxIdValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NonFungibleTokenHoldingsList) SetValue(v NonFungibleTokenHoldingWithTxIdValue)`

SetValue sets Value field to given value.


### GetBlockHeight

`func (o *NonFungibleTokenHoldingsList) GetBlockHeight() int32`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *NonFungibleTokenHoldingsList) GetBlockHeightOk() (*int32, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *NonFungibleTokenHoldingsList) SetBlockHeight(v int32)`

SetBlockHeight sets BlockHeight field to given value.


### GetTxId

`func (o *NonFungibleTokenHoldingsList) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *NonFungibleTokenHoldingsList) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *NonFungibleTokenHoldingsList) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetTx

`func (o *NonFungibleTokenHoldingsList) GetTx() GetTransactionList200ResponseResultsInner`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *NonFungibleTokenHoldingsList) GetTxOk() (*GetTransactionList200ResponseResultsInner, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *NonFungibleTokenHoldingsList) SetTx(v GetTransactionList200ResponseResultsInner)`

SetTx sets Tx field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


