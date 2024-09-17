# GetContractEventsById200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** |  | 
**Offset** | **int32** |  | 
**Results** | [**[]TokenTransferTransactionEventsInner**](TokenTransferTransactionEventsInner.md) |  | 

## Methods

### NewGetContractEventsById200Response

`func NewGetContractEventsById200Response(limit int32, offset int32, results []TokenTransferTransactionEventsInner, ) *GetContractEventsById200Response`

NewGetContractEventsById200Response instantiates a new GetContractEventsById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContractEventsById200ResponseWithDefaults

`func NewGetContractEventsById200ResponseWithDefaults() *GetContractEventsById200Response`

NewGetContractEventsById200ResponseWithDefaults instantiates a new GetContractEventsById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *GetContractEventsById200Response) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetContractEventsById200Response) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetContractEventsById200Response) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *GetContractEventsById200Response) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetContractEventsById200Response) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetContractEventsById200Response) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetResults

`func (o *GetContractEventsById200Response) GetResults() []TokenTransferTransactionEventsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetContractEventsById200Response) GetResultsOk() (*[]TokenTransferTransactionEventsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetContractEventsById200Response) SetResults(v []TokenTransferTransactionEventsInner)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


