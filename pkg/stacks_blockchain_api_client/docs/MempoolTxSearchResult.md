# MempoolTxSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | The id used to search this query. | 
**EntityType** | **string** |  | 
**TxData** | [**MempoolTxSearchResultTxData**](MempoolTxSearchResultTxData.md) |  | 
**Metadata** | Pointer to [**GetMempoolTransactionList200ResponseResultsInner**](GetMempoolTransactionList200ResponseResultsInner.md) |  | [optional] 

## Methods

### NewMempoolTxSearchResult

`func NewMempoolTxSearchResult(entityId string, entityType string, txData MempoolTxSearchResultTxData, ) *MempoolTxSearchResult`

NewMempoolTxSearchResult instantiates a new MempoolTxSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMempoolTxSearchResultWithDefaults

`func NewMempoolTxSearchResultWithDefaults() *MempoolTxSearchResult`

NewMempoolTxSearchResultWithDefaults instantiates a new MempoolTxSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *MempoolTxSearchResult) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *MempoolTxSearchResult) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *MempoolTxSearchResult) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetEntityType

`func (o *MempoolTxSearchResult) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *MempoolTxSearchResult) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *MempoolTxSearchResult) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetTxData

`func (o *MempoolTxSearchResult) GetTxData() MempoolTxSearchResultTxData`

GetTxData returns the TxData field if non-nil, zero value otherwise.

### GetTxDataOk

`func (o *MempoolTxSearchResult) GetTxDataOk() (*MempoolTxSearchResultTxData, bool)`

GetTxDataOk returns a tuple with the TxData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxData

`func (o *MempoolTxSearchResult) SetTxData(v MempoolTxSearchResultTxData)`

SetTxData sets TxData field to given value.


### GetMetadata

`func (o *MempoolTxSearchResult) GetMetadata() GetMempoolTransactionList200ResponseResultsInner`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MempoolTxSearchResult) GetMetadataOk() (*GetMempoolTransactionList200ResponseResultsInner, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MempoolTxSearchResult) SetMetadata(v GetMempoolTransactionList200ResponseResultsInner)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MempoolTxSearchResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


