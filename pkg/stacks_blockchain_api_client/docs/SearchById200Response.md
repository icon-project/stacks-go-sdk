# SearchById200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Found** | **bool** |  | 
**Result** | [**SearchById200ResponseResult**](SearchById200ResponseResult.md) |  | 

## Methods

### NewSearchById200Response

`func NewSearchById200Response(found bool, result SearchById200ResponseResult, ) *SearchById200Response`

NewSearchById200Response instantiates a new SearchById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchById200ResponseWithDefaults

`func NewSearchById200ResponseWithDefaults() *SearchById200Response`

NewSearchById200ResponseWithDefaults instantiates a new SearchById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFound

`func (o *SearchById200Response) GetFound() bool`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *SearchById200Response) GetFoundOk() (*bool, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *SearchById200Response) SetFound(v bool)`

SetFound sets Found field to given value.


### GetResult

`func (o *SearchById200Response) GetResult() SearchById200ResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SearchById200Response) GetResultOk() (*SearchById200ResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SearchById200Response) SetResult(v SearchById200ResponseResult)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


