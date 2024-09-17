# AddressSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | The id used to search this query. | 
**EntityType** | **string** |  | 
**Metadata** | Pointer to [**AddressStxBalance**](AddressStxBalance.md) |  | [optional] 

## Methods

### NewAddressSearchResult

`func NewAddressSearchResult(entityId string, entityType string, ) *AddressSearchResult`

NewAddressSearchResult instantiates a new AddressSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressSearchResultWithDefaults

`func NewAddressSearchResultWithDefaults() *AddressSearchResult`

NewAddressSearchResultWithDefaults instantiates a new AddressSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *AddressSearchResult) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AddressSearchResult) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AddressSearchResult) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetEntityType

`func (o *AddressSearchResult) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *AddressSearchResult) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *AddressSearchResult) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetMetadata

`func (o *AddressSearchResult) GetMetadata() AddressStxBalance`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AddressSearchResult) GetMetadataOk() (*AddressStxBalance, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AddressSearchResult) SetMetadata(v AddressStxBalance)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AddressSearchResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


