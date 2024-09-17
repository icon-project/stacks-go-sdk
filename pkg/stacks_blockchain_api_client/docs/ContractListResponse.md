# ContractListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** |  | 
**Offset** | **int32** |  | 
**Results** | [**[]SmartContract**](SmartContract.md) |  | 

## Methods

### NewContractListResponse

`func NewContractListResponse(limit int32, offset int32, results []SmartContract, ) *ContractListResponse`

NewContractListResponse instantiates a new ContractListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractListResponseWithDefaults

`func NewContractListResponseWithDefaults() *ContractListResponse`

NewContractListResponseWithDefaults instantiates a new ContractListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *ContractListResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ContractListResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ContractListResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *ContractListResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ContractListResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ContractListResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetResults

`func (o *ContractListResponse) GetResults() []SmartContract`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ContractListResponse) GetResultsOk() (*[]SmartContract, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ContractListResponse) SetResults(v []SmartContract)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


