# SearchById200ResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | The id used to search this query. | 
**EntityType** | **string** |  | 
**Metadata** | Pointer to [**GetTransactionList200ResponseResultsInner**](GetTransactionList200ResponseResultsInner.md) |  | [optional] 
**BlockData** | [**BlockSearchResultBlockData**](BlockSearchResultBlockData.md) |  | 
**TxData** | [**TxSearchResultTxData**](TxSearchResultTxData.md) |  | 

## Methods

### NewSearchById200ResponseResult

`func NewSearchById200ResponseResult(entityId string, entityType string, blockData BlockSearchResultBlockData, txData TxSearchResultTxData, ) *SearchById200ResponseResult`

NewSearchById200ResponseResult instantiates a new SearchById200ResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchById200ResponseResultWithDefaults

`func NewSearchById200ResponseResultWithDefaults() *SearchById200ResponseResult`

NewSearchById200ResponseResultWithDefaults instantiates a new SearchById200ResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *SearchById200ResponseResult) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *SearchById200ResponseResult) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *SearchById200ResponseResult) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetEntityType

`func (o *SearchById200ResponseResult) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *SearchById200ResponseResult) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *SearchById200ResponseResult) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetMetadata

`func (o *SearchById200ResponseResult) GetMetadata() GetTransactionList200ResponseResultsInner`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SearchById200ResponseResult) GetMetadataOk() (*GetTransactionList200ResponseResultsInner, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SearchById200ResponseResult) SetMetadata(v GetTransactionList200ResponseResultsInner)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SearchById200ResponseResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetBlockData

`func (o *SearchById200ResponseResult) GetBlockData() BlockSearchResultBlockData`

GetBlockData returns the BlockData field if non-nil, zero value otherwise.

### GetBlockDataOk

`func (o *SearchById200ResponseResult) GetBlockDataOk() (*BlockSearchResultBlockData, bool)`

GetBlockDataOk returns a tuple with the BlockData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockData

`func (o *SearchById200ResponseResult) SetBlockData(v BlockSearchResultBlockData)`

SetBlockData sets BlockData field to given value.


### GetTxData

`func (o *SearchById200ResponseResult) GetTxData() TxSearchResultTxData`

GetTxData returns the TxData field if non-nil, zero value otherwise.

### GetTxDataOk

`func (o *SearchById200ResponseResult) GetTxDataOk() (*TxSearchResultTxData, bool)`

GetTxDataOk returns a tuple with the TxData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxData

`func (o *SearchById200ResponseResult) SetTxData(v TxSearchResultTxData)`

SetTxData sets TxData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


