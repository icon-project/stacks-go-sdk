# TenureChangeTransactionTenureChangePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenureConsensusHash** | **string** | Consensus hash of this tenure. Corresponds to the sortition in which the miner of this block was chosen. | 
**PrevTenureConsensusHash** | **string** | Consensus hash of the previous tenure. Corresponds to the sortition of the previous winning block-commit. | 
**BurnViewConsensusHash** | **string** | Current consensus hash on the underlying burnchain. Corresponds to the last-seen sortition. | 
**PreviousTenureEnd** | **string** | (Hex string) Stacks Block hash | 
**PreviousTenureBlocks** | **int32** | The number of blocks produced in the previous tenure. | 
**Cause** | [**TenureChangeTransactionTenureChangePayloadCause**](TenureChangeTransactionTenureChangePayloadCause.md) |  | 
**PubkeyHash** | **string** | (Hex string) The ECDSA public key hash of the current tenure. | 

## Methods

### NewTenureChangeTransactionTenureChangePayload

`func NewTenureChangeTransactionTenureChangePayload(tenureConsensusHash string, prevTenureConsensusHash string, burnViewConsensusHash string, previousTenureEnd string, previousTenureBlocks int32, cause TenureChangeTransactionTenureChangePayloadCause, pubkeyHash string, ) *TenureChangeTransactionTenureChangePayload`

NewTenureChangeTransactionTenureChangePayload instantiates a new TenureChangeTransactionTenureChangePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenureChangeTransactionTenureChangePayloadWithDefaults

`func NewTenureChangeTransactionTenureChangePayloadWithDefaults() *TenureChangeTransactionTenureChangePayload`

NewTenureChangeTransactionTenureChangePayloadWithDefaults instantiates a new TenureChangeTransactionTenureChangePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenureConsensusHash

`func (o *TenureChangeTransactionTenureChangePayload) GetTenureConsensusHash() string`

GetTenureConsensusHash returns the TenureConsensusHash field if non-nil, zero value otherwise.

### GetTenureConsensusHashOk

`func (o *TenureChangeTransactionTenureChangePayload) GetTenureConsensusHashOk() (*string, bool)`

GetTenureConsensusHashOk returns a tuple with the TenureConsensusHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenureConsensusHash

`func (o *TenureChangeTransactionTenureChangePayload) SetTenureConsensusHash(v string)`

SetTenureConsensusHash sets TenureConsensusHash field to given value.


### GetPrevTenureConsensusHash

`func (o *TenureChangeTransactionTenureChangePayload) GetPrevTenureConsensusHash() string`

GetPrevTenureConsensusHash returns the PrevTenureConsensusHash field if non-nil, zero value otherwise.

### GetPrevTenureConsensusHashOk

`func (o *TenureChangeTransactionTenureChangePayload) GetPrevTenureConsensusHashOk() (*string, bool)`

GetPrevTenureConsensusHashOk returns a tuple with the PrevTenureConsensusHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevTenureConsensusHash

`func (o *TenureChangeTransactionTenureChangePayload) SetPrevTenureConsensusHash(v string)`

SetPrevTenureConsensusHash sets PrevTenureConsensusHash field to given value.


### GetBurnViewConsensusHash

`func (o *TenureChangeTransactionTenureChangePayload) GetBurnViewConsensusHash() string`

GetBurnViewConsensusHash returns the BurnViewConsensusHash field if non-nil, zero value otherwise.

### GetBurnViewConsensusHashOk

`func (o *TenureChangeTransactionTenureChangePayload) GetBurnViewConsensusHashOk() (*string, bool)`

GetBurnViewConsensusHashOk returns a tuple with the BurnViewConsensusHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnViewConsensusHash

`func (o *TenureChangeTransactionTenureChangePayload) SetBurnViewConsensusHash(v string)`

SetBurnViewConsensusHash sets BurnViewConsensusHash field to given value.


### GetPreviousTenureEnd

`func (o *TenureChangeTransactionTenureChangePayload) GetPreviousTenureEnd() string`

GetPreviousTenureEnd returns the PreviousTenureEnd field if non-nil, zero value otherwise.

### GetPreviousTenureEndOk

`func (o *TenureChangeTransactionTenureChangePayload) GetPreviousTenureEndOk() (*string, bool)`

GetPreviousTenureEndOk returns a tuple with the PreviousTenureEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousTenureEnd

`func (o *TenureChangeTransactionTenureChangePayload) SetPreviousTenureEnd(v string)`

SetPreviousTenureEnd sets PreviousTenureEnd field to given value.


### GetPreviousTenureBlocks

`func (o *TenureChangeTransactionTenureChangePayload) GetPreviousTenureBlocks() int32`

GetPreviousTenureBlocks returns the PreviousTenureBlocks field if non-nil, zero value otherwise.

### GetPreviousTenureBlocksOk

`func (o *TenureChangeTransactionTenureChangePayload) GetPreviousTenureBlocksOk() (*int32, bool)`

GetPreviousTenureBlocksOk returns a tuple with the PreviousTenureBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousTenureBlocks

`func (o *TenureChangeTransactionTenureChangePayload) SetPreviousTenureBlocks(v int32)`

SetPreviousTenureBlocks sets PreviousTenureBlocks field to given value.


### GetCause

`func (o *TenureChangeTransactionTenureChangePayload) GetCause() TenureChangeTransactionTenureChangePayloadCause`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *TenureChangeTransactionTenureChangePayload) GetCauseOk() (*TenureChangeTransactionTenureChangePayloadCause, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *TenureChangeTransactionTenureChangePayload) SetCause(v TenureChangeTransactionTenureChangePayloadCause)`

SetCause sets Cause field to given value.


### GetPubkeyHash

`func (o *TenureChangeTransactionTenureChangePayload) GetPubkeyHash() string`

GetPubkeyHash returns the PubkeyHash field if non-nil, zero value otherwise.

### GetPubkeyHashOk

`func (o *TenureChangeTransactionTenureChangePayload) GetPubkeyHashOk() (*string, bool)`

GetPubkeyHashOk returns a tuple with the PubkeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkeyHash

`func (o *TenureChangeTransactionTenureChangePayload) SetPubkeyHash(v string)`

SetPubkeyHash sets PubkeyHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


