# BlockSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | The id used to search this query. | 
**EntityType** | **string** |  | 
**BlockData** | [**BlockSearchResultBlockData**](BlockSearchResultBlockData.md) |  | 
**Metadata** | Pointer to [**Block**](Block.md) |  | [optional] 

## Methods

### NewBlockSearchResult

`func NewBlockSearchResult(entityId string, entityType string, blockData BlockSearchResultBlockData, ) *BlockSearchResult`

NewBlockSearchResult instantiates a new BlockSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockSearchResultWithDefaults

`func NewBlockSearchResultWithDefaults() *BlockSearchResult`

NewBlockSearchResultWithDefaults instantiates a new BlockSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *BlockSearchResult) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *BlockSearchResult) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *BlockSearchResult) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetEntityType

`func (o *BlockSearchResult) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *BlockSearchResult) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *BlockSearchResult) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetBlockData

`func (o *BlockSearchResult) GetBlockData() BlockSearchResultBlockData`

GetBlockData returns the BlockData field if non-nil, zero value otherwise.

### GetBlockDataOk

`func (o *BlockSearchResult) GetBlockDataOk() (*BlockSearchResultBlockData, bool)`

GetBlockDataOk returns a tuple with the BlockData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockData

`func (o *BlockSearchResult) SetBlockData(v BlockSearchResultBlockData)`

SetBlockData sets BlockData field to given value.


### GetMetadata

`func (o *BlockSearchResult) GetMetadata() Block`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BlockSearchResult) GetMetadataOk() (*Block, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BlockSearchResult) SetMetadata(v Block)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BlockSearchResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


