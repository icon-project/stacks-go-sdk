# Microblock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Canonical** | **bool** | Set to &#x60;true&#x60; if the microblock corresponds to the canonical chain tip. | 
**MicroblockCanonical** | **bool** | Set to &#x60;true&#x60; if the microblock was not orphaned in a following anchor block. Defaults to &#x60;true&#x60; if the following anchor block has not yet been created. | 
**MicroblockHash** | **string** | The SHA512/256 hash of this microblock. | 
**MicroblockSequence** | **int32** | A hint to describe how to order a set of microblocks. Starts at 0. | 
**MicroblockParentHash** | **string** | The SHA512/256 hash of the previous signed microblock in this stream. | 
**BlockHeight** | **int32** | The anchor block height that confirmed this microblock. | 
**ParentBlockHeight** | **int32** | The height of the anchor block that preceded this microblock. | 
**ParentBlockHash** | **string** | The hash of the anchor block that preceded this microblock. | 
**ParentBurnBlockHash** | **string** | The hash of the Bitcoin block that preceded this microblock. | 
**ParentBurnBlockTime** | **int32** | The block timestamp of the Bitcoin block that preceded this microblock. | 
**ParentBurnBlockTimeIso** | **string** | The ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) formatted block time of the bitcoin block that preceded this microblock. | 
**ParentBurnBlockHeight** | **int32** | The height of the Bitcoin block that preceded this microblock. | 
**BlockHash** | **NullableString** | The hash of the anchor block that confirmed this microblock. This wil be empty for unanchored microblocks | 
**Txs** | **[]string** | List of transactions included in the microblock | 

## Methods

### NewMicroblock

`func NewMicroblock(canonical bool, microblockCanonical bool, microblockHash string, microblockSequence int32, microblockParentHash string, blockHeight int32, parentBlockHeight int32, parentBlockHash string, parentBurnBlockHash string, parentBurnBlockTime int32, parentBurnBlockTimeIso string, parentBurnBlockHeight int32, blockHash NullableString, txs []string, ) *Microblock`

NewMicroblock instantiates a new Microblock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicroblockWithDefaults

`func NewMicroblockWithDefaults() *Microblock`

NewMicroblockWithDefaults instantiates a new Microblock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanonical

`func (o *Microblock) GetCanonical() bool`

GetCanonical returns the Canonical field if non-nil, zero value otherwise.

### GetCanonicalOk

`func (o *Microblock) GetCanonicalOk() (*bool, bool)`

GetCanonicalOk returns a tuple with the Canonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonical

`func (o *Microblock) SetCanonical(v bool)`

SetCanonical sets Canonical field to given value.


### GetMicroblockCanonical

`func (o *Microblock) GetMicroblockCanonical() bool`

GetMicroblockCanonical returns the MicroblockCanonical field if non-nil, zero value otherwise.

### GetMicroblockCanonicalOk

`func (o *Microblock) GetMicroblockCanonicalOk() (*bool, bool)`

GetMicroblockCanonicalOk returns a tuple with the MicroblockCanonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroblockCanonical

`func (o *Microblock) SetMicroblockCanonical(v bool)`

SetMicroblockCanonical sets MicroblockCanonical field to given value.


### GetMicroblockHash

`func (o *Microblock) GetMicroblockHash() string`

GetMicroblockHash returns the MicroblockHash field if non-nil, zero value otherwise.

### GetMicroblockHashOk

`func (o *Microblock) GetMicroblockHashOk() (*string, bool)`

GetMicroblockHashOk returns a tuple with the MicroblockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroblockHash

`func (o *Microblock) SetMicroblockHash(v string)`

SetMicroblockHash sets MicroblockHash field to given value.


### GetMicroblockSequence

`func (o *Microblock) GetMicroblockSequence() int32`

GetMicroblockSequence returns the MicroblockSequence field if non-nil, zero value otherwise.

### GetMicroblockSequenceOk

`func (o *Microblock) GetMicroblockSequenceOk() (*int32, bool)`

GetMicroblockSequenceOk returns a tuple with the MicroblockSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroblockSequence

`func (o *Microblock) SetMicroblockSequence(v int32)`

SetMicroblockSequence sets MicroblockSequence field to given value.


### GetMicroblockParentHash

`func (o *Microblock) GetMicroblockParentHash() string`

GetMicroblockParentHash returns the MicroblockParentHash field if non-nil, zero value otherwise.

### GetMicroblockParentHashOk

`func (o *Microblock) GetMicroblockParentHashOk() (*string, bool)`

GetMicroblockParentHashOk returns a tuple with the MicroblockParentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroblockParentHash

`func (o *Microblock) SetMicroblockParentHash(v string)`

SetMicroblockParentHash sets MicroblockParentHash field to given value.


### GetBlockHeight

`func (o *Microblock) GetBlockHeight() int32`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *Microblock) GetBlockHeightOk() (*int32, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *Microblock) SetBlockHeight(v int32)`

SetBlockHeight sets BlockHeight field to given value.


### GetParentBlockHeight

`func (o *Microblock) GetParentBlockHeight() int32`

GetParentBlockHeight returns the ParentBlockHeight field if non-nil, zero value otherwise.

### GetParentBlockHeightOk

`func (o *Microblock) GetParentBlockHeightOk() (*int32, bool)`

GetParentBlockHeightOk returns a tuple with the ParentBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBlockHeight

`func (o *Microblock) SetParentBlockHeight(v int32)`

SetParentBlockHeight sets ParentBlockHeight field to given value.


### GetParentBlockHash

`func (o *Microblock) GetParentBlockHash() string`

GetParentBlockHash returns the ParentBlockHash field if non-nil, zero value otherwise.

### GetParentBlockHashOk

`func (o *Microblock) GetParentBlockHashOk() (*string, bool)`

GetParentBlockHashOk returns a tuple with the ParentBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBlockHash

`func (o *Microblock) SetParentBlockHash(v string)`

SetParentBlockHash sets ParentBlockHash field to given value.


### GetParentBurnBlockHash

`func (o *Microblock) GetParentBurnBlockHash() string`

GetParentBurnBlockHash returns the ParentBurnBlockHash field if non-nil, zero value otherwise.

### GetParentBurnBlockHashOk

`func (o *Microblock) GetParentBurnBlockHashOk() (*string, bool)`

GetParentBurnBlockHashOk returns a tuple with the ParentBurnBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBurnBlockHash

`func (o *Microblock) SetParentBurnBlockHash(v string)`

SetParentBurnBlockHash sets ParentBurnBlockHash field to given value.


### GetParentBurnBlockTime

`func (o *Microblock) GetParentBurnBlockTime() int32`

GetParentBurnBlockTime returns the ParentBurnBlockTime field if non-nil, zero value otherwise.

### GetParentBurnBlockTimeOk

`func (o *Microblock) GetParentBurnBlockTimeOk() (*int32, bool)`

GetParentBurnBlockTimeOk returns a tuple with the ParentBurnBlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBurnBlockTime

`func (o *Microblock) SetParentBurnBlockTime(v int32)`

SetParentBurnBlockTime sets ParentBurnBlockTime field to given value.


### GetParentBurnBlockTimeIso

`func (o *Microblock) GetParentBurnBlockTimeIso() string`

GetParentBurnBlockTimeIso returns the ParentBurnBlockTimeIso field if non-nil, zero value otherwise.

### GetParentBurnBlockTimeIsoOk

`func (o *Microblock) GetParentBurnBlockTimeIsoOk() (*string, bool)`

GetParentBurnBlockTimeIsoOk returns a tuple with the ParentBurnBlockTimeIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBurnBlockTimeIso

`func (o *Microblock) SetParentBurnBlockTimeIso(v string)`

SetParentBurnBlockTimeIso sets ParentBurnBlockTimeIso field to given value.


### GetParentBurnBlockHeight

`func (o *Microblock) GetParentBurnBlockHeight() int32`

GetParentBurnBlockHeight returns the ParentBurnBlockHeight field if non-nil, zero value otherwise.

### GetParentBurnBlockHeightOk

`func (o *Microblock) GetParentBurnBlockHeightOk() (*int32, bool)`

GetParentBurnBlockHeightOk returns a tuple with the ParentBurnBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBurnBlockHeight

`func (o *Microblock) SetParentBurnBlockHeight(v int32)`

SetParentBurnBlockHeight sets ParentBurnBlockHeight field to given value.


### GetBlockHash

`func (o *Microblock) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *Microblock) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *Microblock) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.


### SetBlockHashNil

`func (o *Microblock) SetBlockHashNil(b bool)`

 SetBlockHashNil sets the value for BlockHash to be an explicit nil

### UnsetBlockHash
`func (o *Microblock) UnsetBlockHash()`

UnsetBlockHash ensures that no value is present for BlockHash, not even an explicit nil
### GetTxs

`func (o *Microblock) GetTxs() []string`

GetTxs returns the Txs field if non-nil, zero value otherwise.

### GetTxsOk

`func (o *Microblock) GetTxsOk() (*[]string, bool)`

GetTxsOk returns a tuple with the Txs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxs

`func (o *Microblock) SetTxs(v []string)`

SetTxs sets Txs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


