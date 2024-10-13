# GetDroppedMempoolTransactionList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** |  | 
**Offset** | **int32** |  | 
**Total** | **int32** |  | 
**Results** | [**[]GetMempoolTransactionList200ResponseResultsInner**](GetMempoolTransactionList200ResponseResultsInner.md) |  | 

## Methods

### NewGetDroppedMempoolTransactionList200Response

`func NewGetDroppedMempoolTransactionList200Response(limit int32, offset int32, total int32, results []GetMempoolTransactionList200ResponseResultsInner, ) *GetDroppedMempoolTransactionList200Response`

NewGetDroppedMempoolTransactionList200Response instantiates a new GetDroppedMempoolTransactionList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDroppedMempoolTransactionList200ResponseWithDefaults

`func NewGetDroppedMempoolTransactionList200ResponseWithDefaults() *GetDroppedMempoolTransactionList200Response`

NewGetDroppedMempoolTransactionList200ResponseWithDefaults instantiates a new GetDroppedMempoolTransactionList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *GetDroppedMempoolTransactionList200Response) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetDroppedMempoolTransactionList200Response) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetDroppedMempoolTransactionList200Response) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *GetDroppedMempoolTransactionList200Response) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetDroppedMempoolTransactionList200Response) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetDroppedMempoolTransactionList200Response) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetTotal

`func (o *GetDroppedMempoolTransactionList200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetDroppedMempoolTransactionList200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetDroppedMempoolTransactionList200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetResults

`func (o *GetDroppedMempoolTransactionList200Response) GetResults() []GetMempoolTransactionList200ResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetDroppedMempoolTransactionList200Response) GetResultsOk() (*[]GetMempoolTransactionList200ResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetDroppedMempoolTransactionList200Response) SetResults(v []GetMempoolTransactionList200ResponseResultsInner)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

