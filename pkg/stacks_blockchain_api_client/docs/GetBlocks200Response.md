# GetBlocks200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** |  | 
**Offset** | **int32** |  | 
**Total** | **int32** |  | 
**NextCursor** | **NullableString** | Next page cursor | 
**PrevCursor** | **NullableString** | Previous page cursor | 
**Cursor** | **NullableString** | Current page cursor | 
**Results** | [**[]GetBlocks200ResponseResultsInner**](GetBlocks200ResponseResultsInner.md) |  | 

## Methods

### NewGetBlocks200Response

`func NewGetBlocks200Response(limit int32, offset int32, total int32, nextCursor NullableString, prevCursor NullableString, cursor NullableString, results []GetBlocks200ResponseResultsInner, ) *GetBlocks200Response`

NewGetBlocks200Response instantiates a new GetBlocks200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBlocks200ResponseWithDefaults

`func NewGetBlocks200ResponseWithDefaults() *GetBlocks200Response`

NewGetBlocks200ResponseWithDefaults instantiates a new GetBlocks200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *GetBlocks200Response) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetBlocks200Response) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetBlocks200Response) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *GetBlocks200Response) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetBlocks200Response) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetBlocks200Response) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetTotal

`func (o *GetBlocks200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetBlocks200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetBlocks200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetNextCursor

`func (o *GetBlocks200Response) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *GetBlocks200Response) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *GetBlocks200Response) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *GetBlocks200Response) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *GetBlocks200Response) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetPrevCursor

`func (o *GetBlocks200Response) GetPrevCursor() string`

GetPrevCursor returns the PrevCursor field if non-nil, zero value otherwise.

### GetPrevCursorOk

`func (o *GetBlocks200Response) GetPrevCursorOk() (*string, bool)`

GetPrevCursorOk returns a tuple with the PrevCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevCursor

`func (o *GetBlocks200Response) SetPrevCursor(v string)`

SetPrevCursor sets PrevCursor field to given value.


### SetPrevCursorNil

`func (o *GetBlocks200Response) SetPrevCursorNil(b bool)`

 SetPrevCursorNil sets the value for PrevCursor to be an explicit nil

### UnsetPrevCursor
`func (o *GetBlocks200Response) UnsetPrevCursor()`

UnsetPrevCursor ensures that no value is present for PrevCursor, not even an explicit nil
### GetCursor

`func (o *GetBlocks200Response) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *GetBlocks200Response) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *GetBlocks200Response) SetCursor(v string)`

SetCursor sets Cursor field to given value.


### SetCursorNil

`func (o *GetBlocks200Response) SetCursorNil(b bool)`

 SetCursorNil sets the value for Cursor to be an explicit nil

### UnsetCursor
`func (o *GetBlocks200Response) UnsetCursor()`

UnsetCursor ensures that no value is present for Cursor, not even an explicit nil
### GetResults

`func (o *GetBlocks200Response) GetResults() []GetBlocks200ResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetBlocks200Response) GetResultsOk() (*[]GetBlocks200ResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetBlocks200Response) SetResults(v []GetBlocks200ResponseResultsInner)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


