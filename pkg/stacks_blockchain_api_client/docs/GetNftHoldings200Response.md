# GetNftHoldings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** |  | 
**Offset** | **int32** |  | 
**Total** | **int32** |  | 
**Results** | [**[]NonFungibleTokenHoldingsList**](NonFungibleTokenHoldingsList.md) |  | 

## Methods

### NewGetNftHoldings200Response

`func NewGetNftHoldings200Response(limit int32, offset int32, total int32, results []NonFungibleTokenHoldingsList, ) *GetNftHoldings200Response`

NewGetNftHoldings200Response instantiates a new GetNftHoldings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNftHoldings200ResponseWithDefaults

`func NewGetNftHoldings200ResponseWithDefaults() *GetNftHoldings200Response`

NewGetNftHoldings200ResponseWithDefaults instantiates a new GetNftHoldings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *GetNftHoldings200Response) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetNftHoldings200Response) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetNftHoldings200Response) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *GetNftHoldings200Response) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetNftHoldings200Response) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetNftHoldings200Response) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetTotal

`func (o *GetNftHoldings200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetNftHoldings200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetNftHoldings200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetResults

`func (o *GetNftHoldings200Response) GetResults() []NonFungibleTokenHoldingsList`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetNftHoldings200Response) GetResultsOk() (*[]NonFungibleTokenHoldingsList, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetNftHoldings200Response) SetResults(v []NonFungibleTokenHoldingsList)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


