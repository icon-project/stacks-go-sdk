# RunFaucetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | Indicates if the faucet call was successful | 
**Txid** | **string** | The transaction ID for the faucet call | 
**RawTx** | **string** | Raw transaction in hex string representation | 

## Methods

### NewRunFaucetResponse

`func NewRunFaucetResponse(success bool, txid string, rawTx string, ) *RunFaucetResponse`

NewRunFaucetResponse instantiates a new RunFaucetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunFaucetResponseWithDefaults

`func NewRunFaucetResponseWithDefaults() *RunFaucetResponse`

NewRunFaucetResponseWithDefaults instantiates a new RunFaucetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *RunFaucetResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *RunFaucetResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *RunFaucetResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetTxid

`func (o *RunFaucetResponse) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *RunFaucetResponse) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *RunFaucetResponse) SetTxid(v string)`

SetTxid sets Txid field to given value.


### GetRawTx

`func (o *RunFaucetResponse) GetRawTx() string`

GetRawTx returns the RawTx field if non-nil, zero value otherwise.

### GetRawTxOk

`func (o *RunFaucetResponse) GetRawTxOk() (*string, bool)`

GetRawTxOk returns a tuple with the RawTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawTx

`func (o *RunFaucetResponse) SetRawTx(v string)`

SetRawTx sets RawTx field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


