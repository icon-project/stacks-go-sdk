# GetBurnchainRewardList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** |  | 
**Offset** | **int32** |  | 
**Results** | [**[]BurnchainReward**](BurnchainReward.md) |  | 

## Methods

### NewGetBurnchainRewardList200Response

`func NewGetBurnchainRewardList200Response(limit int32, offset int32, results []BurnchainReward, ) *GetBurnchainRewardList200Response`

NewGetBurnchainRewardList200Response instantiates a new GetBurnchainRewardList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBurnchainRewardList200ResponseWithDefaults

`func NewGetBurnchainRewardList200ResponseWithDefaults() *GetBurnchainRewardList200Response`

NewGetBurnchainRewardList200ResponseWithDefaults instantiates a new GetBurnchainRewardList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *GetBurnchainRewardList200Response) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetBurnchainRewardList200Response) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetBurnchainRewardList200Response) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *GetBurnchainRewardList200Response) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetBurnchainRewardList200Response) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetBurnchainRewardList200Response) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetResults

`func (o *GetBurnchainRewardList200Response) GetResults() []BurnchainReward`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetBurnchainRewardList200Response) GetResultsOk() (*[]BurnchainReward, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetBurnchainRewardList200Response) SetResults(v []BurnchainReward)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


