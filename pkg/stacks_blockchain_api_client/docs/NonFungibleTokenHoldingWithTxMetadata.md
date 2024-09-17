# NonFungibleTokenHoldingWithTxMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetIdentifier** | **string** |  | 
**Value** | [**NonFungibleTokenHoldingWithTxIdValue**](NonFungibleTokenHoldingWithTxIdValue.md) |  | 
**BlockHeight** | **int32** |  | 
**Tx** | [**GetTransactionList200ResponseResultsInner**](GetTransactionList200ResponseResultsInner.md) |  | 

## Methods

### NewNonFungibleTokenHoldingWithTxMetadata

`func NewNonFungibleTokenHoldingWithTxMetadata(assetIdentifier string, value NonFungibleTokenHoldingWithTxIdValue, blockHeight int32, tx GetTransactionList200ResponseResultsInner, ) *NonFungibleTokenHoldingWithTxMetadata`

NewNonFungibleTokenHoldingWithTxMetadata instantiates a new NonFungibleTokenHoldingWithTxMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonFungibleTokenHoldingWithTxMetadataWithDefaults

`func NewNonFungibleTokenHoldingWithTxMetadataWithDefaults() *NonFungibleTokenHoldingWithTxMetadata`

NewNonFungibleTokenHoldingWithTxMetadataWithDefaults instantiates a new NonFungibleTokenHoldingWithTxMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetIdentifier

`func (o *NonFungibleTokenHoldingWithTxMetadata) GetAssetIdentifier() string`

GetAssetIdentifier returns the AssetIdentifier field if non-nil, zero value otherwise.

### GetAssetIdentifierOk

`func (o *NonFungibleTokenHoldingWithTxMetadata) GetAssetIdentifierOk() (*string, bool)`

GetAssetIdentifierOk returns a tuple with the AssetIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetIdentifier

`func (o *NonFungibleTokenHoldingWithTxMetadata) SetAssetIdentifier(v string)`

SetAssetIdentifier sets AssetIdentifier field to given value.


### GetValue

`func (o *NonFungibleTokenHoldingWithTxMetadata) GetValue() NonFungibleTokenHoldingWithTxIdValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NonFungibleTokenHoldingWithTxMetadata) GetValueOk() (*NonFungibleTokenHoldingWithTxIdValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NonFungibleTokenHoldingWithTxMetadata) SetValue(v NonFungibleTokenHoldingWithTxIdValue)`

SetValue sets Value field to given value.


### GetBlockHeight

`func (o *NonFungibleTokenHoldingWithTxMetadata) GetBlockHeight() int32`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *NonFungibleTokenHoldingWithTxMetadata) GetBlockHeightOk() (*int32, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *NonFungibleTokenHoldingWithTxMetadata) SetBlockHeight(v int32)`

SetBlockHeight sets BlockHeight field to given value.


### GetTx

`func (o *NonFungibleTokenHoldingWithTxMetadata) GetTx() GetTransactionList200ResponseResultsInner`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *NonFungibleTokenHoldingWithTxMetadata) GetTxOk() (*GetTransactionList200ResponseResultsInner, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *NonFungibleTokenHoldingWithTxMetadata) SetTx(v GetTransactionList200ResponseResultsInner)`

SetTx sets Tx field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


