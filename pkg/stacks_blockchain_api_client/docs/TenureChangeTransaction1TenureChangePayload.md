# TenureChangeTransaction1TenureChangePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenureConsensusHash** | **string** | Consensus hash of this tenure. Corresponds to the sortition in which the miner of this block was chosen. | 
**PrevTenureConsensusHash** | **string** | Consensus hash of the previous tenure. Corresponds to the sortition of the previous winning block-commit. | 
**BurnViewConsensusHash** | **string** | Current consensus hash on the underlying burnchain. Corresponds to the last-seen sortition. | 
**PreviousTenureEnd** | **string** | (Hex string) Stacks Block hash | 
**PreviousTenureBlocks** | **int32** | The number of blocks produced in the previous tenure. | 
**Cause** | [**TenureChangeTransaction1TenureChangePayloadCause**](TenureChangeTransaction1TenureChangePayloadCause.md) |  | 
**PubkeyHash** | **string** | (Hex string) The ECDSA public key hash of the current tenure. | 

## Methods

### NewTenureChangeTransaction1TenureChangePayload

`func NewTenureChangeTransaction1TenureChangePayload(tenureConsensusHash string, prevTenureConsensusHash string, burnViewConsensusHash string, previousTenureEnd string, previousTenureBlocks int32, cause TenureChangeTransaction1TenureChangePayloadCause, pubkeyHash string, ) *TenureChangeTransaction1TenureChangePayload`

NewTenureChangeTransaction1TenureChangePayload instantiates a new TenureChangeTransaction1TenureChangePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenureChangeTransaction1TenureChangePayloadWithDefaults

`func NewTenureChangeTransaction1TenureChangePayloadWithDefaults() *TenureChangeTransaction1TenureChangePayload`

NewTenureChangeTransaction1TenureChangePayloadWithDefaults instantiates a new TenureChangeTransaction1TenureChangePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenureConsensusHash

`func (o *TenureChangeTransaction1TenureChangePayload) GetTenureConsensusHash() string`

GetTenureConsensusHash returns the TenureConsensusHash field if non-nil, zero value otherwise.

### GetTenureConsensusHashOk

`func (o *TenureChangeTransaction1TenureChangePayload) GetTenureConsensusHashOk() (*string, bool)`

GetTenureConsensusHashOk returns a tuple with the TenureConsensusHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenureConsensusHash

`func (o *TenureChangeTransaction1TenureChangePayload) SetTenureConsensusHash(v string)`

SetTenureConsensusHash sets TenureConsensusHash field to given value.


### GetPrevTenureConsensusHash

`func (o *TenureChangeTransaction1TenureChangePayload) GetPrevTenureConsensusHash() string`

GetPrevTenureConsensusHash returns the PrevTenureConsensusHash field if non-nil, zero value otherwise.

### GetPrevTenureConsensusHashOk

`func (o *TenureChangeTransaction1TenureChangePayload) GetPrevTenureConsensusHashOk() (*string, bool)`

GetPrevTenureConsensusHashOk returns a tuple with the PrevTenureConsensusHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevTenureConsensusHash

`func (o *TenureChangeTransaction1TenureChangePayload) SetPrevTenureConsensusHash(v string)`

SetPrevTenureConsensusHash sets PrevTenureConsensusHash field to given value.


### GetBurnViewConsensusHash

`func (o *TenureChangeTransaction1TenureChangePayload) GetBurnViewConsensusHash() string`

GetBurnViewConsensusHash returns the BurnViewConsensusHash field if non-nil, zero value otherwise.

### GetBurnViewConsensusHashOk

`func (o *TenureChangeTransaction1TenureChangePayload) GetBurnViewConsensusHashOk() (*string, bool)`

GetBurnViewConsensusHashOk returns a tuple with the BurnViewConsensusHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnViewConsensusHash

`func (o *TenureChangeTransaction1TenureChangePayload) SetBurnViewConsensusHash(v string)`

SetBurnViewConsensusHash sets BurnViewConsensusHash field to given value.


### GetPreviousTenureEnd

`func (o *TenureChangeTransaction1TenureChangePayload) GetPreviousTenureEnd() string`

GetPreviousTenureEnd returns the PreviousTenureEnd field if non-nil, zero value otherwise.

### GetPreviousTenureEndOk

`func (o *TenureChangeTransaction1TenureChangePayload) GetPreviousTenureEndOk() (*string, bool)`

GetPreviousTenureEndOk returns a tuple with the PreviousTenureEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousTenureEnd

`func (o *TenureChangeTransaction1TenureChangePayload) SetPreviousTenureEnd(v string)`

SetPreviousTenureEnd sets PreviousTenureEnd field to given value.


### GetPreviousTenureBlocks

`func (o *TenureChangeTransaction1TenureChangePayload) GetPreviousTenureBlocks() int32`

GetPreviousTenureBlocks returns the PreviousTenureBlocks field if non-nil, zero value otherwise.

### GetPreviousTenureBlocksOk

`func (o *TenureChangeTransaction1TenureChangePayload) GetPreviousTenureBlocksOk() (*int32, bool)`

GetPreviousTenureBlocksOk returns a tuple with the PreviousTenureBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousTenureBlocks

`func (o *TenureChangeTransaction1TenureChangePayload) SetPreviousTenureBlocks(v int32)`

SetPreviousTenureBlocks sets PreviousTenureBlocks field to given value.


### GetCause

`func (o *TenureChangeTransaction1TenureChangePayload) GetCause() TenureChangeTransaction1TenureChangePayloadCause`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *TenureChangeTransaction1TenureChangePayload) GetCauseOk() (*TenureChangeTransaction1TenureChangePayloadCause, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *TenureChangeTransaction1TenureChangePayload) SetCause(v TenureChangeTransaction1TenureChangePayloadCause)`

SetCause sets Cause field to given value.


### GetPubkeyHash

`func (o *TenureChangeTransaction1TenureChangePayload) GetPubkeyHash() string`

GetPubkeyHash returns the PubkeyHash field if non-nil, zero value otherwise.

### GetPubkeyHashOk

`func (o *TenureChangeTransaction1TenureChangePayload) GetPubkeyHashOk() (*string, bool)`

GetPubkeyHashOk returns a tuple with the PubkeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkeyHash

`func (o *TenureChangeTransaction1TenureChangePayload) SetPubkeyHash(v string)`

SetPubkeyHash sets PubkeyHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


