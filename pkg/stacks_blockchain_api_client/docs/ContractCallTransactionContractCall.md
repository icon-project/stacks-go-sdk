# ContractCallTransactionContractCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractId** | **string** | Contract identifier formatted as &#x60;&lt;principaladdress&gt;.&lt;contract_name&gt;&#x60; | 
**FunctionName** | **string** | Name of the Clarity function to be invoked | 
**FunctionSignature** | **string** | Function definition, including function name and type as well as parameter names and types | 
**FunctionArgs** | Pointer to [**[]ContractCallTransactionContractCallFunctionArgsInner**](ContractCallTransactionContractCallFunctionArgsInner.md) |  | [optional] 

## Methods

### NewContractCallTransactionContractCall

`func NewContractCallTransactionContractCall(contractId string, functionName string, functionSignature string, ) *ContractCallTransactionContractCall`

NewContractCallTransactionContractCall instantiates a new ContractCallTransactionContractCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractCallTransactionContractCallWithDefaults

`func NewContractCallTransactionContractCallWithDefaults() *ContractCallTransactionContractCall`

NewContractCallTransactionContractCallWithDefaults instantiates a new ContractCallTransactionContractCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractId

`func (o *ContractCallTransactionContractCall) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *ContractCallTransactionContractCall) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *ContractCallTransactionContractCall) SetContractId(v string)`

SetContractId sets ContractId field to given value.


### GetFunctionName

`func (o *ContractCallTransactionContractCall) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *ContractCallTransactionContractCall) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *ContractCallTransactionContractCall) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetFunctionSignature

`func (o *ContractCallTransactionContractCall) GetFunctionSignature() string`

GetFunctionSignature returns the FunctionSignature field if non-nil, zero value otherwise.

### GetFunctionSignatureOk

`func (o *ContractCallTransactionContractCall) GetFunctionSignatureOk() (*string, bool)`

GetFunctionSignatureOk returns a tuple with the FunctionSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionSignature

`func (o *ContractCallTransactionContractCall) SetFunctionSignature(v string)`

SetFunctionSignature sets FunctionSignature field to given value.


### GetFunctionArgs

`func (o *ContractCallTransactionContractCall) GetFunctionArgs() []ContractCallTransactionContractCallFunctionArgsInner`

GetFunctionArgs returns the FunctionArgs field if non-nil, zero value otherwise.

### GetFunctionArgsOk

`func (o *ContractCallTransactionContractCall) GetFunctionArgsOk() (*[]ContractCallTransactionContractCallFunctionArgsInner, bool)`

GetFunctionArgsOk returns a tuple with the FunctionArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionArgs

`func (o *ContractCallTransactionContractCall) SetFunctionArgs(v []ContractCallTransactionContractCallFunctionArgsInner)`

SetFunctionArgs sets FunctionArgs field to given value.

### HasFunctionArgs

`func (o *ContractCallTransactionContractCall) HasFunctionArgs() bool`

HasFunctionArgs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


