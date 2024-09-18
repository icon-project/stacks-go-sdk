# RunFaucetResponse1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | Indicates if the faucet call was successful | 
**TxId** | **string** | The transaction ID for the faucet call | 
**TxRaw** | **string** | Raw transaction in hex string representation | 

## Methods

### NewRunFaucetResponse1

`func NewRunFaucetResponse1(success bool, txId string, txRaw string, ) *RunFaucetResponse1`

NewRunFaucetResponse1 instantiates a new RunFaucetResponse1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunFaucetResponse1WithDefaults

`func NewRunFaucetResponse1WithDefaults() *RunFaucetResponse1`

NewRunFaucetResponse1WithDefaults instantiates a new RunFaucetResponse1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *RunFaucetResponse1) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *RunFaucetResponse1) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *RunFaucetResponse1) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetTxId

`func (o *RunFaucetResponse1) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *RunFaucetResponse1) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *RunFaucetResponse1) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetTxRaw

`func (o *RunFaucetResponse1) GetTxRaw() string`

GetTxRaw returns the TxRaw field if non-nil, zero value otherwise.

### GetTxRawOk

`func (o *RunFaucetResponse1) GetTxRawOk() (*string, bool)`

GetTxRawOk returns a tuple with the TxRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRaw

`func (o *RunFaucetResponse1) SetTxRaw(v string)`

SetTxRaw sets TxRaw field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


