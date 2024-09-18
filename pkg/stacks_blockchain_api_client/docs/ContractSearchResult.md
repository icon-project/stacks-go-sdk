# ContractSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | The id used to search this query. | 
**EntityType** | **string** |  | 
**TxData** | Pointer to [**ContractSearchResultTxData**](ContractSearchResultTxData.md) |  | [optional] 
**Metadata** | Pointer to [**GetTransactionById200Response**](GetTransactionById200Response.md) |  | [optional] 

## Methods

### NewContractSearchResult

`func NewContractSearchResult(entityId string, entityType string, ) *ContractSearchResult`

NewContractSearchResult instantiates a new ContractSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractSearchResultWithDefaults

`func NewContractSearchResultWithDefaults() *ContractSearchResult`

NewContractSearchResultWithDefaults instantiates a new ContractSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *ContractSearchResult) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *ContractSearchResult) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *ContractSearchResult) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetEntityType

`func (o *ContractSearchResult) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *ContractSearchResult) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *ContractSearchResult) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetTxData

`func (o *ContractSearchResult) GetTxData() ContractSearchResultTxData`

GetTxData returns the TxData field if non-nil, zero value otherwise.

### GetTxDataOk

`func (o *ContractSearchResult) GetTxDataOk() (*ContractSearchResultTxData, bool)`

GetTxDataOk returns a tuple with the TxData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxData

`func (o *ContractSearchResult) SetTxData(v ContractSearchResultTxData)`

SetTxData sets TxData field to given value.

### HasTxData

`func (o *ContractSearchResult) HasTxData() bool`

HasTxData returns a boolean if a field has been set.

### GetMetadata

`func (o *ContractSearchResult) GetMetadata() GetTransactionById200Response`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ContractSearchResult) GetMetadataOk() (*GetTransactionById200Response, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ContractSearchResult) SetMetadata(v GetTransactionById200Response)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ContractSearchResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


