# TransactionFound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Found** | **bool** |  | 
**Result** | [**TransactionFoundResult**](TransactionFoundResult.md) |  | 

## Methods

### NewTransactionFound

`func NewTransactionFound(found bool, result TransactionFoundResult, ) *TransactionFound`

NewTransactionFound instantiates a new TransactionFound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionFoundWithDefaults

`func NewTransactionFoundWithDefaults() *TransactionFound`

NewTransactionFoundWithDefaults instantiates a new TransactionFound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFound

`func (o *TransactionFound) GetFound() bool`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *TransactionFound) GetFoundOk() (*bool, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *TransactionFound) SetFound(v bool)`

SetFound sets Found field to given value.


### GetResult

`func (o *TransactionFound) GetResult() TransactionFoundResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TransactionFound) GetResultOk() (*TransactionFoundResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TransactionFound) SetResult(v TransactionFoundResult)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


