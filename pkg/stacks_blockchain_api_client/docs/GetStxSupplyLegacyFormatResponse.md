# GetStxSupplyLegacyFormatResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnlockedPercent** | **string** | String quoted decimal number of the percentage of STX that have unlocked | 
**TotalStacks** | **string** | String quoted decimal number of the total possible number of STX | 
**TotalStacksFormatted** | **string** | Same as &#x60;totalStacks&#x60; but formatted with comma thousands separators | 
**UnlockedSupply** | **string** | String quoted decimal number of the STX that have been mined or unlocked | 
**UnlockedSupplyFormatted** | **string** | Same as &#x60;unlockedSupply&#x60; but formatted with comma thousands separators | 
**BlockHeight** | **string** | The block height at which this information was queried | 

## Methods

### NewGetStxSupplyLegacyFormatResponse

`func NewGetStxSupplyLegacyFormatResponse(unlockedPercent string, totalStacks string, totalStacksFormatted string, unlockedSupply string, unlockedSupplyFormatted string, blockHeight string, ) *GetStxSupplyLegacyFormatResponse`

NewGetStxSupplyLegacyFormatResponse instantiates a new GetStxSupplyLegacyFormatResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStxSupplyLegacyFormatResponseWithDefaults

`func NewGetStxSupplyLegacyFormatResponseWithDefaults() *GetStxSupplyLegacyFormatResponse`

NewGetStxSupplyLegacyFormatResponseWithDefaults instantiates a new GetStxSupplyLegacyFormatResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnlockedPercent

`func (o *GetStxSupplyLegacyFormatResponse) GetUnlockedPercent() string`

GetUnlockedPercent returns the UnlockedPercent field if non-nil, zero value otherwise.

### GetUnlockedPercentOk

`func (o *GetStxSupplyLegacyFormatResponse) GetUnlockedPercentOk() (*string, bool)`

GetUnlockedPercentOk returns a tuple with the UnlockedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockedPercent

`func (o *GetStxSupplyLegacyFormatResponse) SetUnlockedPercent(v string)`

SetUnlockedPercent sets UnlockedPercent field to given value.


### GetTotalStacks

`func (o *GetStxSupplyLegacyFormatResponse) GetTotalStacks() string`

GetTotalStacks returns the TotalStacks field if non-nil, zero value otherwise.

### GetTotalStacksOk

`func (o *GetStxSupplyLegacyFormatResponse) GetTotalStacksOk() (*string, bool)`

GetTotalStacksOk returns a tuple with the TotalStacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStacks

`func (o *GetStxSupplyLegacyFormatResponse) SetTotalStacks(v string)`

SetTotalStacks sets TotalStacks field to given value.


### GetTotalStacksFormatted

`func (o *GetStxSupplyLegacyFormatResponse) GetTotalStacksFormatted() string`

GetTotalStacksFormatted returns the TotalStacksFormatted field if non-nil, zero value otherwise.

### GetTotalStacksFormattedOk

`func (o *GetStxSupplyLegacyFormatResponse) GetTotalStacksFormattedOk() (*string, bool)`

GetTotalStacksFormattedOk returns a tuple with the TotalStacksFormatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStacksFormatted

`func (o *GetStxSupplyLegacyFormatResponse) SetTotalStacksFormatted(v string)`

SetTotalStacksFormatted sets TotalStacksFormatted field to given value.


### GetUnlockedSupply

`func (o *GetStxSupplyLegacyFormatResponse) GetUnlockedSupply() string`

GetUnlockedSupply returns the UnlockedSupply field if non-nil, zero value otherwise.

### GetUnlockedSupplyOk

`func (o *GetStxSupplyLegacyFormatResponse) GetUnlockedSupplyOk() (*string, bool)`

GetUnlockedSupplyOk returns a tuple with the UnlockedSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockedSupply

`func (o *GetStxSupplyLegacyFormatResponse) SetUnlockedSupply(v string)`

SetUnlockedSupply sets UnlockedSupply field to given value.


### GetUnlockedSupplyFormatted

`func (o *GetStxSupplyLegacyFormatResponse) GetUnlockedSupplyFormatted() string`

GetUnlockedSupplyFormatted returns the UnlockedSupplyFormatted field if non-nil, zero value otherwise.

### GetUnlockedSupplyFormattedOk

`func (o *GetStxSupplyLegacyFormatResponse) GetUnlockedSupplyFormattedOk() (*string, bool)`

GetUnlockedSupplyFormattedOk returns a tuple with the UnlockedSupplyFormatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockedSupplyFormatted

`func (o *GetStxSupplyLegacyFormatResponse) SetUnlockedSupplyFormatted(v string)`

SetUnlockedSupplyFormatted sets UnlockedSupplyFormatted field to given value.


### GetBlockHeight

`func (o *GetStxSupplyLegacyFormatResponse) GetBlockHeight() string`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *GetStxSupplyLegacyFormatResponse) GetBlockHeightOk() (*string, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *GetStxSupplyLegacyFormatResponse) SetBlockHeight(v string)`

SetBlockHeight sets BlockHeight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


