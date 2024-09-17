# GetBurnBlocks200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** |  | 
**Offset** | **int32** |  | 
**Total** | **int32** |  | 
**Results** | [**[]GetBurnBlocks200ResponseResultsInner**](GetBurnBlocks200ResponseResultsInner.md) |  | 

## Methods

### NewGetBurnBlocks200Response

`func NewGetBurnBlocks200Response(limit int32, offset int32, total int32, results []GetBurnBlocks200ResponseResultsInner, ) *GetBurnBlocks200Response`

NewGetBurnBlocks200Response instantiates a new GetBurnBlocks200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBurnBlocks200ResponseWithDefaults

`func NewGetBurnBlocks200ResponseWithDefaults() *GetBurnBlocks200Response`

NewGetBurnBlocks200ResponseWithDefaults instantiates a new GetBurnBlocks200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *GetBurnBlocks200Response) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetBurnBlocks200Response) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetBurnBlocks200Response) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *GetBurnBlocks200Response) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetBurnBlocks200Response) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetBurnBlocks200Response) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetTotal

`func (o *GetBurnBlocks200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetBurnBlocks200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetBurnBlocks200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetResults

`func (o *GetBurnBlocks200Response) GetResults() []GetBurnBlocks200ResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetBurnBlocks200Response) GetResultsOk() (*[]GetBurnBlocks200ResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetBurnBlocks200Response) SetResults(v []GetBurnBlocks200ResponseResultsInner)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


