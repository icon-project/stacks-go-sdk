# ApiStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerVersion** | **string** | the server version that is currently running | 
**Status** | **string** | the current server status | 
**PoxV1UnlockHeight** | Pointer to **NullableInt32** |  | [optional] 
**PoxV2UnlockHeight** | Pointer to **NullableInt32** |  | [optional] 
**PoxV3UnlockHeight** | Pointer to **NullableInt32** |  | [optional] 
**ChainTip** | Pointer to [**NullableApiStatusResponseChainTip**](ApiStatusResponseChainTip.md) |  | [optional] 

## Methods

### NewApiStatusResponse

`func NewApiStatusResponse(serverVersion string, status string, ) *ApiStatusResponse`

NewApiStatusResponse instantiates a new ApiStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStatusResponseWithDefaults

`func NewApiStatusResponseWithDefaults() *ApiStatusResponse`

NewApiStatusResponseWithDefaults instantiates a new ApiStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerVersion

`func (o *ApiStatusResponse) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *ApiStatusResponse) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *ApiStatusResponse) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.


### GetStatus

`func (o *ApiStatusResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiStatusResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiStatusResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPoxV1UnlockHeight

`func (o *ApiStatusResponse) GetPoxV1UnlockHeight() int32`

GetPoxV1UnlockHeight returns the PoxV1UnlockHeight field if non-nil, zero value otherwise.

### GetPoxV1UnlockHeightOk

`func (o *ApiStatusResponse) GetPoxV1UnlockHeightOk() (*int32, bool)`

GetPoxV1UnlockHeightOk returns a tuple with the PoxV1UnlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoxV1UnlockHeight

`func (o *ApiStatusResponse) SetPoxV1UnlockHeight(v int32)`

SetPoxV1UnlockHeight sets PoxV1UnlockHeight field to given value.

### HasPoxV1UnlockHeight

`func (o *ApiStatusResponse) HasPoxV1UnlockHeight() bool`

HasPoxV1UnlockHeight returns a boolean if a field has been set.

### SetPoxV1UnlockHeightNil

`func (o *ApiStatusResponse) SetPoxV1UnlockHeightNil(b bool)`

 SetPoxV1UnlockHeightNil sets the value for PoxV1UnlockHeight to be an explicit nil

### UnsetPoxV1UnlockHeight
`func (o *ApiStatusResponse) UnsetPoxV1UnlockHeight()`

UnsetPoxV1UnlockHeight ensures that no value is present for PoxV1UnlockHeight, not even an explicit nil
### GetPoxV2UnlockHeight

`func (o *ApiStatusResponse) GetPoxV2UnlockHeight() int32`

GetPoxV2UnlockHeight returns the PoxV2UnlockHeight field if non-nil, zero value otherwise.

### GetPoxV2UnlockHeightOk

`func (o *ApiStatusResponse) GetPoxV2UnlockHeightOk() (*int32, bool)`

GetPoxV2UnlockHeightOk returns a tuple with the PoxV2UnlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoxV2UnlockHeight

`func (o *ApiStatusResponse) SetPoxV2UnlockHeight(v int32)`

SetPoxV2UnlockHeight sets PoxV2UnlockHeight field to given value.

### HasPoxV2UnlockHeight

`func (o *ApiStatusResponse) HasPoxV2UnlockHeight() bool`

HasPoxV2UnlockHeight returns a boolean if a field has been set.

### SetPoxV2UnlockHeightNil

`func (o *ApiStatusResponse) SetPoxV2UnlockHeightNil(b bool)`

 SetPoxV2UnlockHeightNil sets the value for PoxV2UnlockHeight to be an explicit nil

### UnsetPoxV2UnlockHeight
`func (o *ApiStatusResponse) UnsetPoxV2UnlockHeight()`

UnsetPoxV2UnlockHeight ensures that no value is present for PoxV2UnlockHeight, not even an explicit nil
### GetPoxV3UnlockHeight

`func (o *ApiStatusResponse) GetPoxV3UnlockHeight() int32`

GetPoxV3UnlockHeight returns the PoxV3UnlockHeight field if non-nil, zero value otherwise.

### GetPoxV3UnlockHeightOk

`func (o *ApiStatusResponse) GetPoxV3UnlockHeightOk() (*int32, bool)`

GetPoxV3UnlockHeightOk returns a tuple with the PoxV3UnlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoxV3UnlockHeight

`func (o *ApiStatusResponse) SetPoxV3UnlockHeight(v int32)`

SetPoxV3UnlockHeight sets PoxV3UnlockHeight field to given value.

### HasPoxV3UnlockHeight

`func (o *ApiStatusResponse) HasPoxV3UnlockHeight() bool`

HasPoxV3UnlockHeight returns a boolean if a field has been set.

### SetPoxV3UnlockHeightNil

`func (o *ApiStatusResponse) SetPoxV3UnlockHeightNil(b bool)`

 SetPoxV3UnlockHeightNil sets the value for PoxV3UnlockHeight to be an explicit nil

### UnsetPoxV3UnlockHeight
`func (o *ApiStatusResponse) UnsetPoxV3UnlockHeight()`

UnsetPoxV3UnlockHeight ensures that no value is present for PoxV3UnlockHeight, not even an explicit nil
### GetChainTip

`func (o *ApiStatusResponse) GetChainTip() ApiStatusResponseChainTip`

GetChainTip returns the ChainTip field if non-nil, zero value otherwise.

### GetChainTipOk

`func (o *ApiStatusResponse) GetChainTipOk() (*ApiStatusResponseChainTip, bool)`

GetChainTipOk returns a tuple with the ChainTip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainTip

`func (o *ApiStatusResponse) SetChainTip(v ApiStatusResponseChainTip)`

SetChainTip sets ChainTip field to given value.

### HasChainTip

`func (o *ApiStatusResponse) HasChainTip() bool`

HasChainTip returns a boolean if a field has been set.

### SetChainTipNil

`func (o *ApiStatusResponse) SetChainTipNil(b bool)`

 SetChainTipNil sets the value for ChainTip to be an explicit nil

### UnsetChainTip
`func (o *ApiStatusResponse) UnsetChainTip()`

UnsetChainTip ensures that no value is present for ChainTip, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


