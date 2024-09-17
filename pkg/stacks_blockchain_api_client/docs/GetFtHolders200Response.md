# GetFtHolders200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalSupply** | **string** | The total supply of the token (the sum of all balances) | 
**Limit** | **int32** |  | 
**Offset** | **int32** |  | 
**Total** | **int32** |  | 
**Results** | [**[]FtHolderEntry**](FtHolderEntry.md) |  | 

## Methods

### NewGetFtHolders200Response

`func NewGetFtHolders200Response(totalSupply string, limit int32, offset int32, total int32, results []FtHolderEntry, ) *GetFtHolders200Response`

NewGetFtHolders200Response instantiates a new GetFtHolders200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFtHolders200ResponseWithDefaults

`func NewGetFtHolders200ResponseWithDefaults() *GetFtHolders200Response`

NewGetFtHolders200ResponseWithDefaults instantiates a new GetFtHolders200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalSupply

`func (o *GetFtHolders200Response) GetTotalSupply() string`

GetTotalSupply returns the TotalSupply field if non-nil, zero value otherwise.

### GetTotalSupplyOk

`func (o *GetFtHolders200Response) GetTotalSupplyOk() (*string, bool)`

GetTotalSupplyOk returns a tuple with the TotalSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSupply

`func (o *GetFtHolders200Response) SetTotalSupply(v string)`

SetTotalSupply sets TotalSupply field to given value.


### GetLimit

`func (o *GetFtHolders200Response) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetFtHolders200Response) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetFtHolders200Response) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *GetFtHolders200Response) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetFtHolders200Response) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetFtHolders200Response) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetTotal

`func (o *GetFtHolders200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetFtHolders200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetFtHolders200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetResults

`func (o *GetFtHolders200Response) GetResults() []FtHolderEntry`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetFtHolders200Response) GetResultsOk() (*[]FtHolderEntry, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetFtHolders200Response) SetResults(v []FtHolderEntry)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


