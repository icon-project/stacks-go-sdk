# ContractSearchResultTxData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Canonical** | Pointer to **bool** |  | [optional] 
**BlockHash** | Pointer to **string** |  | [optional] 
**BurnBlockTime** | Pointer to **int32** |  | [optional] 
**BlockHeight** | Pointer to **int32** |  | [optional] 
**TxType** | **string** |  | 
**TxId** | **string** |  | 

## Methods

### NewContractSearchResultTxData

`func NewContractSearchResultTxData(txType string, txId string, ) *ContractSearchResultTxData`

NewContractSearchResultTxData instantiates a new ContractSearchResultTxData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractSearchResultTxDataWithDefaults

`func NewContractSearchResultTxDataWithDefaults() *ContractSearchResultTxData`

NewContractSearchResultTxDataWithDefaults instantiates a new ContractSearchResultTxData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanonical

`func (o *ContractSearchResultTxData) GetCanonical() bool`

GetCanonical returns the Canonical field if non-nil, zero value otherwise.

### GetCanonicalOk

`func (o *ContractSearchResultTxData) GetCanonicalOk() (*bool, bool)`

GetCanonicalOk returns a tuple with the Canonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonical

`func (o *ContractSearchResultTxData) SetCanonical(v bool)`

SetCanonical sets Canonical field to given value.

### HasCanonical

`func (o *ContractSearchResultTxData) HasCanonical() bool`

HasCanonical returns a boolean if a field has been set.

### GetBlockHash

`func (o *ContractSearchResultTxData) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *ContractSearchResultTxData) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *ContractSearchResultTxData) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.

### HasBlockHash

`func (o *ContractSearchResultTxData) HasBlockHash() bool`

HasBlockHash returns a boolean if a field has been set.

### GetBurnBlockTime

`func (o *ContractSearchResultTxData) GetBurnBlockTime() int32`

GetBurnBlockTime returns the BurnBlockTime field if non-nil, zero value otherwise.

### GetBurnBlockTimeOk

`func (o *ContractSearchResultTxData) GetBurnBlockTimeOk() (*int32, bool)`

GetBurnBlockTimeOk returns a tuple with the BurnBlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockTime

`func (o *ContractSearchResultTxData) SetBurnBlockTime(v int32)`

SetBurnBlockTime sets BurnBlockTime field to given value.

### HasBurnBlockTime

`func (o *ContractSearchResultTxData) HasBurnBlockTime() bool`

HasBurnBlockTime returns a boolean if a field has been set.

### GetBlockHeight

`func (o *ContractSearchResultTxData) GetBlockHeight() int32`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *ContractSearchResultTxData) GetBlockHeightOk() (*int32, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *ContractSearchResultTxData) SetBlockHeight(v int32)`

SetBlockHeight sets BlockHeight field to given value.

### HasBlockHeight

`func (o *ContractSearchResultTxData) HasBlockHeight() bool`

HasBlockHeight returns a boolean if a field has been set.

### GetTxType

`func (o *ContractSearchResultTxData) GetTxType() string`

GetTxType returns the TxType field if non-nil, zero value otherwise.

### GetTxTypeOk

`func (o *ContractSearchResultTxData) GetTxTypeOk() (*string, bool)`

GetTxTypeOk returns a tuple with the TxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxType

`func (o *ContractSearchResultTxData) SetTxType(v string)`

SetTxType sets TxType field to given value.


### GetTxId

`func (o *ContractSearchResultTxData) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *ContractSearchResultTxData) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *ContractSearchResultTxData) SetTxId(v string)`

SetTxId sets TxId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


