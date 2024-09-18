# GetNftHistory200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** |  | 
**Offset** | **int32** |  | 
**Total** | **int32** |  | 
**Results** | [**[]NonFungibleTokenHistoryEvent**](NonFungibleTokenHistoryEvent.md) |  | 

## Methods

### NewGetNftHistory200Response

`func NewGetNftHistory200Response(limit int32, offset int32, total int32, results []NonFungibleTokenHistoryEvent, ) *GetNftHistory200Response`

NewGetNftHistory200Response instantiates a new GetNftHistory200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNftHistory200ResponseWithDefaults

`func NewGetNftHistory200ResponseWithDefaults() *GetNftHistory200Response`

NewGetNftHistory200ResponseWithDefaults instantiates a new GetNftHistory200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *GetNftHistory200Response) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetNftHistory200Response) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetNftHistory200Response) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *GetNftHistory200Response) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetNftHistory200Response) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetNftHistory200Response) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetTotal

`func (o *GetNftHistory200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetNftHistory200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetNftHistory200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetResults

`func (o *GetNftHistory200Response) GetResults() []NonFungibleTokenHistoryEvent`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetNftHistory200Response) GetResultsOk() (*[]NonFungibleTokenHistoryEvent, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetNftHistory200Response) SetResults(v []NonFungibleTokenHistoryEvent)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


