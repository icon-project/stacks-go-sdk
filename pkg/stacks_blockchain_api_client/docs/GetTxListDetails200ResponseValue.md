# GetTxListDetails200ResponseValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Found** | **bool** |  | 
**Result** | [**TransactionNotFoundResult**](TransactionNotFoundResult.md) |  | 

## Methods

### NewGetTxListDetails200ResponseValue

`func NewGetTxListDetails200ResponseValue(found bool, result TransactionNotFoundResult, ) *GetTxListDetails200ResponseValue`

NewGetTxListDetails200ResponseValue instantiates a new GetTxListDetails200ResponseValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTxListDetails200ResponseValueWithDefaults

`func NewGetTxListDetails200ResponseValueWithDefaults() *GetTxListDetails200ResponseValue`

NewGetTxListDetails200ResponseValueWithDefaults instantiates a new GetTxListDetails200ResponseValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFound

`func (o *GetTxListDetails200ResponseValue) GetFound() bool`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *GetTxListDetails200ResponseValue) GetFoundOk() (*bool, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *GetTxListDetails200ResponseValue) SetFound(v bool)`

SetFound sets Found field to given value.


### GetResult

`func (o *GetTxListDetails200ResponseValue) GetResult() TransactionNotFoundResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetTxListDetails200ResponseValue) GetResultOk() (*TransactionNotFoundResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetTxListDetails200ResponseValue) SetResult(v TransactionNotFoundResult)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


