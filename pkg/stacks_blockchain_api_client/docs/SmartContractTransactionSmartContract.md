# SmartContractTransactionSmartContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClarityVersion** | **NullableFloat32** | The Clarity version of the contract, only specified for versioned contract transactions, otherwise null | 
**ContractId** | **string** | Contract identifier formatted as &#x60;&lt;principaladdress&gt;.&lt;contract_name&gt;&#x60; | 
**SourceCode** | **string** | Clarity code of the smart contract being deployed | 

## Methods

### NewSmartContractTransactionSmartContract

`func NewSmartContractTransactionSmartContract(clarityVersion NullableFloat32, contractId string, sourceCode string, ) *SmartContractTransactionSmartContract`

NewSmartContractTransactionSmartContract instantiates a new SmartContractTransactionSmartContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartContractTransactionSmartContractWithDefaults

`func NewSmartContractTransactionSmartContractWithDefaults() *SmartContractTransactionSmartContract`

NewSmartContractTransactionSmartContractWithDefaults instantiates a new SmartContractTransactionSmartContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClarityVersion

`func (o *SmartContractTransactionSmartContract) GetClarityVersion() float32`

GetClarityVersion returns the ClarityVersion field if non-nil, zero value otherwise.

### GetClarityVersionOk

`func (o *SmartContractTransactionSmartContract) GetClarityVersionOk() (*float32, bool)`

GetClarityVersionOk returns a tuple with the ClarityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClarityVersion

`func (o *SmartContractTransactionSmartContract) SetClarityVersion(v float32)`

SetClarityVersion sets ClarityVersion field to given value.


### SetClarityVersionNil

`func (o *SmartContractTransactionSmartContract) SetClarityVersionNil(b bool)`

 SetClarityVersionNil sets the value for ClarityVersion to be an explicit nil

### UnsetClarityVersion
`func (o *SmartContractTransactionSmartContract) UnsetClarityVersion()`

UnsetClarityVersion ensures that no value is present for ClarityVersion, not even an explicit nil
### GetContractId

`func (o *SmartContractTransactionSmartContract) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *SmartContractTransactionSmartContract) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *SmartContractTransactionSmartContract) SetContractId(v string)`

SetContractId sets ContractId field to given value.


### GetSourceCode

`func (o *SmartContractTransactionSmartContract) GetSourceCode() string`

GetSourceCode returns the SourceCode field if non-nil, zero value otherwise.

### GetSourceCodeOk

`func (o *SmartContractTransactionSmartContract) GetSourceCodeOk() (*string, bool)`

GetSourceCodeOk returns a tuple with the SourceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCode

`func (o *SmartContractTransactionSmartContract) SetSourceCode(v string)`

SetSourceCode sets SourceCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


