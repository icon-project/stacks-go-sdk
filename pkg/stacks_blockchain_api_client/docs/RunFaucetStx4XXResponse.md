# RunFaucetStx4XXResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | Indicates if the faucet call was successful | 
**Error** | **string** | Error message | 
**Help** | Pointer to **string** |  | [optional] 

## Methods

### NewRunFaucetStx4XXResponse

`func NewRunFaucetStx4XXResponse(success bool, error_ string, ) *RunFaucetStx4XXResponse`

NewRunFaucetStx4XXResponse instantiates a new RunFaucetStx4XXResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunFaucetStx4XXResponseWithDefaults

`func NewRunFaucetStx4XXResponseWithDefaults() *RunFaucetStx4XXResponse`

NewRunFaucetStx4XXResponseWithDefaults instantiates a new RunFaucetStx4XXResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *RunFaucetStx4XXResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *RunFaucetStx4XXResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *RunFaucetStx4XXResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetError

`func (o *RunFaucetStx4XXResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RunFaucetStx4XXResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RunFaucetStx4XXResponse) SetError(v string)`

SetError sets Error field to given value.


### GetHelp

`func (o *RunFaucetStx4XXResponse) GetHelp() string`

GetHelp returns the Help field if non-nil, zero value otherwise.

### GetHelpOk

`func (o *RunFaucetStx4XXResponse) GetHelpOk() (*string, bool)`

GetHelpOk returns a tuple with the Help field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelp

`func (o *RunFaucetStx4XXResponse) SetHelp(v string)`

SetHelp sets Help field to given value.

### HasHelp

`func (o *RunFaucetStx4XXResponse) HasHelp() bool`

HasHelp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


