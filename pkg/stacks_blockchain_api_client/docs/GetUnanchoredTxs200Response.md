# GetUnanchoredTxs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** |  | 
**Results** | [**[]GetTransactionList200ResponseResultsInner**](GetTransactionList200ResponseResultsInner.md) |  | 

## Methods

### NewGetUnanchoredTxs200Response

`func NewGetUnanchoredTxs200Response(total int32, results []GetTransactionList200ResponseResultsInner, ) *GetUnanchoredTxs200Response`

NewGetUnanchoredTxs200Response instantiates a new GetUnanchoredTxs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUnanchoredTxs200ResponseWithDefaults

`func NewGetUnanchoredTxs200ResponseWithDefaults() *GetUnanchoredTxs200Response`

NewGetUnanchoredTxs200ResponseWithDefaults instantiates a new GetUnanchoredTxs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetUnanchoredTxs200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetUnanchoredTxs200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetUnanchoredTxs200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetResults

`func (o *GetUnanchoredTxs200Response) GetResults() []GetTransactionList200ResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetUnanchoredTxs200Response) GetResultsOk() (*[]GetTransactionList200ResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetUnanchoredTxs200Response) SetResults(v []GetTransactionList200ResponseResultsInner)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


