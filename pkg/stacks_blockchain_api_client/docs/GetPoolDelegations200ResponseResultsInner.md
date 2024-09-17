# GetPoolDelegations200ResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stacker** | **string** | The principal of the pool member that issued the delegation | 
**PoxAddr** | Pointer to **string** | The pox-addr value specified by the stacker in the delegation operation | [optional] 
**AmountUstx** | **string** | The amount of uSTX delegated by the stacker | 
**BurnBlockUnlockHeight** | Pointer to **int32** | The optional burnchain block unlock height that the stacker may have specified | [optional] 
**BlockHeight** | **int32** | The block height at which the stacker delegation transaction was mined at | 
**TxId** | **string** | The tx_id of the stacker delegation operation | 

## Methods

### NewGetPoolDelegations200ResponseResultsInner

`func NewGetPoolDelegations200ResponseResultsInner(stacker string, amountUstx string, blockHeight int32, txId string, ) *GetPoolDelegations200ResponseResultsInner`

NewGetPoolDelegations200ResponseResultsInner instantiates a new GetPoolDelegations200ResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPoolDelegations200ResponseResultsInnerWithDefaults

`func NewGetPoolDelegations200ResponseResultsInnerWithDefaults() *GetPoolDelegations200ResponseResultsInner`

NewGetPoolDelegations200ResponseResultsInnerWithDefaults instantiates a new GetPoolDelegations200ResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStacker

`func (o *GetPoolDelegations200ResponseResultsInner) GetStacker() string`

GetStacker returns the Stacker field if non-nil, zero value otherwise.

### GetStackerOk

`func (o *GetPoolDelegations200ResponseResultsInner) GetStackerOk() (*string, bool)`

GetStackerOk returns a tuple with the Stacker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacker

`func (o *GetPoolDelegations200ResponseResultsInner) SetStacker(v string)`

SetStacker sets Stacker field to given value.


### GetPoxAddr

`func (o *GetPoolDelegations200ResponseResultsInner) GetPoxAddr() string`

GetPoxAddr returns the PoxAddr field if non-nil, zero value otherwise.

### GetPoxAddrOk

`func (o *GetPoolDelegations200ResponseResultsInner) GetPoxAddrOk() (*string, bool)`

GetPoxAddrOk returns a tuple with the PoxAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoxAddr

`func (o *GetPoolDelegations200ResponseResultsInner) SetPoxAddr(v string)`

SetPoxAddr sets PoxAddr field to given value.

### HasPoxAddr

`func (o *GetPoolDelegations200ResponseResultsInner) HasPoxAddr() bool`

HasPoxAddr returns a boolean if a field has been set.

### GetAmountUstx

`func (o *GetPoolDelegations200ResponseResultsInner) GetAmountUstx() string`

GetAmountUstx returns the AmountUstx field if non-nil, zero value otherwise.

### GetAmountUstxOk

`func (o *GetPoolDelegations200ResponseResultsInner) GetAmountUstxOk() (*string, bool)`

GetAmountUstxOk returns a tuple with the AmountUstx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountUstx

`func (o *GetPoolDelegations200ResponseResultsInner) SetAmountUstx(v string)`

SetAmountUstx sets AmountUstx field to given value.


### GetBurnBlockUnlockHeight

`func (o *GetPoolDelegations200ResponseResultsInner) GetBurnBlockUnlockHeight() int32`

GetBurnBlockUnlockHeight returns the BurnBlockUnlockHeight field if non-nil, zero value otherwise.

### GetBurnBlockUnlockHeightOk

`func (o *GetPoolDelegations200ResponseResultsInner) GetBurnBlockUnlockHeightOk() (*int32, bool)`

GetBurnBlockUnlockHeightOk returns a tuple with the BurnBlockUnlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockUnlockHeight

`func (o *GetPoolDelegations200ResponseResultsInner) SetBurnBlockUnlockHeight(v int32)`

SetBurnBlockUnlockHeight sets BurnBlockUnlockHeight field to given value.

### HasBurnBlockUnlockHeight

`func (o *GetPoolDelegations200ResponseResultsInner) HasBurnBlockUnlockHeight() bool`

HasBurnBlockUnlockHeight returns a boolean if a field has been set.

### GetBlockHeight

`func (o *GetPoolDelegations200ResponseResultsInner) GetBlockHeight() int32`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *GetPoolDelegations200ResponseResultsInner) GetBlockHeightOk() (*int32, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *GetPoolDelegations200ResponseResultsInner) SetBlockHeight(v int32)`

SetBlockHeight sets BlockHeight field to given value.


### GetTxId

`func (o *GetPoolDelegations200ResponseResultsInner) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *GetPoolDelegations200ResponseResultsInner) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *GetPoolDelegations200ResponseResultsInner) SetTxId(v string)`

SetTxId sets TxId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


