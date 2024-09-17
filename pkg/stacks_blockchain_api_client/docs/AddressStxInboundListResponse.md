# AddressStxInboundListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** |  | 
**Offset** | **int32** |  | 
**Total** | **int32** |  | 
**Results** | [**[]InboundStxTransfer**](InboundStxTransfer.md) |  | 

## Methods

### NewAddressStxInboundListResponse

`func NewAddressStxInboundListResponse(limit int32, offset int32, total int32, results []InboundStxTransfer, ) *AddressStxInboundListResponse`

NewAddressStxInboundListResponse instantiates a new AddressStxInboundListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressStxInboundListResponseWithDefaults

`func NewAddressStxInboundListResponseWithDefaults() *AddressStxInboundListResponse`

NewAddressStxInboundListResponseWithDefaults instantiates a new AddressStxInboundListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *AddressStxInboundListResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AddressStxInboundListResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AddressStxInboundListResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *AddressStxInboundListResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *AddressStxInboundListResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *AddressStxInboundListResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetTotal

`func (o *AddressStxInboundListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AddressStxInboundListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AddressStxInboundListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetResults

`func (o *AddressStxInboundListResponse) GetResults() []InboundStxTransfer`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AddressStxInboundListResponse) GetResultsOk() (*[]InboundStxTransfer, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AddressStxInboundListResponse) SetResults(v []InboundStxTransfer)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


