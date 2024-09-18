# TxSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | The id used to search this query. | 
**EntityType** | **string** |  | 
**TxData** | [**TxSearchResultTxData**](TxSearchResultTxData.md) |  | 
**Metadata** | Pointer to [**GetTransactionList200ResponseResultsInner**](GetTransactionList200ResponseResultsInner.md) |  | [optional] 

## Methods

### NewTxSearchResult

`func NewTxSearchResult(entityId string, entityType string, txData TxSearchResultTxData, ) *TxSearchResult`

NewTxSearchResult instantiates a new TxSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTxSearchResultWithDefaults

`func NewTxSearchResultWithDefaults() *TxSearchResult`

NewTxSearchResultWithDefaults instantiates a new TxSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *TxSearchResult) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *TxSearchResult) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *TxSearchResult) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetEntityType

`func (o *TxSearchResult) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *TxSearchResult) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *TxSearchResult) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetTxData

`func (o *TxSearchResult) GetTxData() TxSearchResultTxData`

GetTxData returns the TxData field if non-nil, zero value otherwise.

### GetTxDataOk

`func (o *TxSearchResult) GetTxDataOk() (*TxSearchResultTxData, bool)`

GetTxDataOk returns a tuple with the TxData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxData

`func (o *TxSearchResult) SetTxData(v TxSearchResultTxData)`

SetTxData sets TxData field to given value.


### GetMetadata

`func (o *TxSearchResult) GetMetadata() GetTransactionList200ResponseResultsInner`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TxSearchResult) GetMetadataOk() (*GetTransactionList200ResponseResultsInner, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TxSearchResult) SetMetadata(v GetTransactionList200ResponseResultsInner)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TxSearchResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


