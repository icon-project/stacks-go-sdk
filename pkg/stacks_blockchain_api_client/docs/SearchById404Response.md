# SearchById404Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Found** | **bool** |  | 
**Result** | [**SearchById404ResponseResult**](SearchById404ResponseResult.md) |  | 
**Error** | **string** |  | 

## Methods

### NewSearchById404Response

`func NewSearchById404Response(found bool, result SearchById404ResponseResult, error_ string, ) *SearchById404Response`

NewSearchById404Response instantiates a new SearchById404Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchById404ResponseWithDefaults

`func NewSearchById404ResponseWithDefaults() *SearchById404Response`

NewSearchById404ResponseWithDefaults instantiates a new SearchById404Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFound

`func (o *SearchById404Response) GetFound() bool`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *SearchById404Response) GetFoundOk() (*bool, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *SearchById404Response) SetFound(v bool)`

SetFound sets Found field to given value.


### GetResult

`func (o *SearchById404Response) GetResult() SearchById404ResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SearchById404Response) GetResultOk() (*SearchById404ResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SearchById404Response) SetResult(v SearchById404ResponseResult)`

SetResult sets Result field to given value.


### GetError

`func (o *SearchById404Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SearchById404Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SearchById404Response) SetError(v string)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


