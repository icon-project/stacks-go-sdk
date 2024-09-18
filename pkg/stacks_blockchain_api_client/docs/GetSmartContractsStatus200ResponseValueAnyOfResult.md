# GetSmartContractsStatus200ResponseValueAnyOfResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Smart contract deployment transaction status | 
**TxId** | **string** | Deployment transaction ID | 
**ContractId** | **string** | Smart contract ID | 
**BlockHeight** | Pointer to **int32** | Height of the transaction confirmation block | [optional] 

## Methods

### NewGetSmartContractsStatus200ResponseValueAnyOfResult

`func NewGetSmartContractsStatus200ResponseValueAnyOfResult(status string, txId string, contractId string, ) *GetSmartContractsStatus200ResponseValueAnyOfResult`

NewGetSmartContractsStatus200ResponseValueAnyOfResult instantiates a new GetSmartContractsStatus200ResponseValueAnyOfResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSmartContractsStatus200ResponseValueAnyOfResultWithDefaults

`func NewGetSmartContractsStatus200ResponseValueAnyOfResultWithDefaults() *GetSmartContractsStatus200ResponseValueAnyOfResult`

NewGetSmartContractsStatus200ResponseValueAnyOfResultWithDefaults instantiates a new GetSmartContractsStatus200ResponseValueAnyOfResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetSmartContractsStatus200ResponseValueAnyOfResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSmartContractsStatus200ResponseValueAnyOfResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSmartContractsStatus200ResponseValueAnyOfResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTxId

`func (o *GetSmartContractsStatus200ResponseValueAnyOfResult) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *GetSmartContractsStatus200ResponseValueAnyOfResult) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *GetSmartContractsStatus200ResponseValueAnyOfResult) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetContractId

`func (o *GetSmartContractsStatus200ResponseValueAnyOfResult) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *GetSmartContractsStatus200ResponseValueAnyOfResult) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *GetSmartContractsStatus200ResponseValueAnyOfResult) SetContractId(v string)`

SetContractId sets ContractId field to given value.


### GetBlockHeight

`func (o *GetSmartContractsStatus200ResponseValueAnyOfResult) GetBlockHeight() int32`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *GetSmartContractsStatus200ResponseValueAnyOfResult) GetBlockHeightOk() (*int32, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *GetSmartContractsStatus200ResponseValueAnyOfResult) SetBlockHeight(v int32)`

SetBlockHeight sets BlockHeight field to given value.

### HasBlockHeight

`func (o *GetSmartContractsStatus200ResponseValueAnyOfResult) HasBlockHeight() bool`

HasBlockHeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


