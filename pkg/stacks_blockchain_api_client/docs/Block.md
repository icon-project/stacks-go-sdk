# Block

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Canonical** | **bool** | Set to &#x60;true&#x60; if block corresponds to the canonical chain tip | 
**Height** | **int32** | Height of the block | 
**Hash** | **string** | Hash representing the block | 
**BlockTime** | **float32** | Unix timestamp (in seconds) indicating when this block was mined. | 
**BlockTimeIso** | **string** | An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) indicating when this block was mined. | 
**IndexBlockHash** | **string** | The only hash that can uniquely identify an anchored block or an unconfirmed state trie | 
**ParentBlockHash** | **string** | Hash of the parent block | 
**BurnBlockTime** | **float32** | Unix timestamp (in seconds) indicating when this block was mined. | 
**BurnBlockTimeIso** | **string** | An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) indicating when this block was mined. | 
**BurnBlockHash** | **string** | Hash of the anchor chain block | 
**BurnBlockHeight** | **int32** | Height of the anchor chain block | 
**MinerTxid** | **string** | Anchor chain transaction ID | 
**ExecutionCostReadCount** | **int32** | Execution cost read count. | 
**ExecutionCostReadLength** | **int32** | Execution cost read length. | 
**ExecutionCostRuntime** | **int32** | Execution cost runtime. | 
**ExecutionCostWriteCount** | **int32** | Execution cost write count. | 
**ExecutionCostWriteLength** | **int32** | Execution cost write length. | 
**Txs** | **[]string** | List of transactions included in the block | 
**ParentMicroblockHash** | **string** | The hash of the last streamed block that precedes this block to which this block is to be appended. Not every anchored block will have a parent microblock stream. An anchored block that does not have a parent microblock stream has the parent microblock hash set to an empty string, and the parent microblock sequence number set to -1. | 
**ParentMicroblockSequence** | **int32** | The hash of the last streamed block that precedes this block to which this block is to be appended. Not every anchored block will have a parent microblock stream. An anchored block that does not have a parent microblock stream has the parent microblock hash set to an empty string, and the parent microblock sequence number set to -1. | 
**MicroblocksAccepted** | **[]string** | List of microblocks that were accepted in this anchor block. Not every anchored block will have a accepted all (or any) of the previously streamed microblocks. Microblocks that were orphaned are not included in this list. | 
**MicroblocksStreamed** | **[]string** | List of microblocks that were streamed/produced by this anchor block&#39;s miner. This list only includes microblocks that were accepted in the following anchor block. Microblocks that were orphaned are not included in this list. | 
**MicroblockTxCount** | **map[string]int32** | List of txs counts in each accepted microblock | 

## Methods

### NewBlock

`func NewBlock(canonical bool, height int32, hash string, blockTime float32, blockTimeIso string, indexBlockHash string, parentBlockHash string, burnBlockTime float32, burnBlockTimeIso string, burnBlockHash string, burnBlockHeight int32, minerTxid string, executionCostReadCount int32, executionCostReadLength int32, executionCostRuntime int32, executionCostWriteCount int32, executionCostWriteLength int32, txs []string, parentMicroblockHash string, parentMicroblockSequence int32, microblocksAccepted []string, microblocksStreamed []string, microblockTxCount map[string]int32, ) *Block`

NewBlock instantiates a new Block object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockWithDefaults

`func NewBlockWithDefaults() *Block`

NewBlockWithDefaults instantiates a new Block object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanonical

`func (o *Block) GetCanonical() bool`

GetCanonical returns the Canonical field if non-nil, zero value otherwise.

### GetCanonicalOk

`func (o *Block) GetCanonicalOk() (*bool, bool)`

GetCanonicalOk returns a tuple with the Canonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonical

`func (o *Block) SetCanonical(v bool)`

SetCanonical sets Canonical field to given value.


### GetHeight

`func (o *Block) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *Block) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *Block) SetHeight(v int32)`

SetHeight sets Height field to given value.


### GetHash

`func (o *Block) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Block) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Block) SetHash(v string)`

SetHash sets Hash field to given value.


### GetBlockTime

`func (o *Block) GetBlockTime() float32`

GetBlockTime returns the BlockTime field if non-nil, zero value otherwise.

### GetBlockTimeOk

`func (o *Block) GetBlockTimeOk() (*float32, bool)`

GetBlockTimeOk returns a tuple with the BlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTime

`func (o *Block) SetBlockTime(v float32)`

SetBlockTime sets BlockTime field to given value.


### GetBlockTimeIso

`func (o *Block) GetBlockTimeIso() string`

GetBlockTimeIso returns the BlockTimeIso field if non-nil, zero value otherwise.

### GetBlockTimeIsoOk

`func (o *Block) GetBlockTimeIsoOk() (*string, bool)`

GetBlockTimeIsoOk returns a tuple with the BlockTimeIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTimeIso

`func (o *Block) SetBlockTimeIso(v string)`

SetBlockTimeIso sets BlockTimeIso field to given value.


### GetIndexBlockHash

`func (o *Block) GetIndexBlockHash() string`

GetIndexBlockHash returns the IndexBlockHash field if non-nil, zero value otherwise.

### GetIndexBlockHashOk

`func (o *Block) GetIndexBlockHashOk() (*string, bool)`

GetIndexBlockHashOk returns a tuple with the IndexBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexBlockHash

`func (o *Block) SetIndexBlockHash(v string)`

SetIndexBlockHash sets IndexBlockHash field to given value.


### GetParentBlockHash

`func (o *Block) GetParentBlockHash() string`

GetParentBlockHash returns the ParentBlockHash field if non-nil, zero value otherwise.

### GetParentBlockHashOk

`func (o *Block) GetParentBlockHashOk() (*string, bool)`

GetParentBlockHashOk returns a tuple with the ParentBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBlockHash

`func (o *Block) SetParentBlockHash(v string)`

SetParentBlockHash sets ParentBlockHash field to given value.


### GetBurnBlockTime

`func (o *Block) GetBurnBlockTime() float32`

GetBurnBlockTime returns the BurnBlockTime field if non-nil, zero value otherwise.

### GetBurnBlockTimeOk

`func (o *Block) GetBurnBlockTimeOk() (*float32, bool)`

GetBurnBlockTimeOk returns a tuple with the BurnBlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockTime

`func (o *Block) SetBurnBlockTime(v float32)`

SetBurnBlockTime sets BurnBlockTime field to given value.


### GetBurnBlockTimeIso

`func (o *Block) GetBurnBlockTimeIso() string`

GetBurnBlockTimeIso returns the BurnBlockTimeIso field if non-nil, zero value otherwise.

### GetBurnBlockTimeIsoOk

`func (o *Block) GetBurnBlockTimeIsoOk() (*string, bool)`

GetBurnBlockTimeIsoOk returns a tuple with the BurnBlockTimeIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockTimeIso

`func (o *Block) SetBurnBlockTimeIso(v string)`

SetBurnBlockTimeIso sets BurnBlockTimeIso field to given value.


### GetBurnBlockHash

`func (o *Block) GetBurnBlockHash() string`

GetBurnBlockHash returns the BurnBlockHash field if non-nil, zero value otherwise.

### GetBurnBlockHashOk

`func (o *Block) GetBurnBlockHashOk() (*string, bool)`

GetBurnBlockHashOk returns a tuple with the BurnBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockHash

`func (o *Block) SetBurnBlockHash(v string)`

SetBurnBlockHash sets BurnBlockHash field to given value.


### GetBurnBlockHeight

`func (o *Block) GetBurnBlockHeight() int32`

GetBurnBlockHeight returns the BurnBlockHeight field if non-nil, zero value otherwise.

### GetBurnBlockHeightOk

`func (o *Block) GetBurnBlockHeightOk() (*int32, bool)`

GetBurnBlockHeightOk returns a tuple with the BurnBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockHeight

`func (o *Block) SetBurnBlockHeight(v int32)`

SetBurnBlockHeight sets BurnBlockHeight field to given value.


### GetMinerTxid

`func (o *Block) GetMinerTxid() string`

GetMinerTxid returns the MinerTxid field if non-nil, zero value otherwise.

### GetMinerTxidOk

`func (o *Block) GetMinerTxidOk() (*string, bool)`

GetMinerTxidOk returns a tuple with the MinerTxid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinerTxid

`func (o *Block) SetMinerTxid(v string)`

SetMinerTxid sets MinerTxid field to given value.


### GetExecutionCostReadCount

`func (o *Block) GetExecutionCostReadCount() int32`

GetExecutionCostReadCount returns the ExecutionCostReadCount field if non-nil, zero value otherwise.

### GetExecutionCostReadCountOk

`func (o *Block) GetExecutionCostReadCountOk() (*int32, bool)`

GetExecutionCostReadCountOk returns a tuple with the ExecutionCostReadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCostReadCount

`func (o *Block) SetExecutionCostReadCount(v int32)`

SetExecutionCostReadCount sets ExecutionCostReadCount field to given value.


### GetExecutionCostReadLength

`func (o *Block) GetExecutionCostReadLength() int32`

GetExecutionCostReadLength returns the ExecutionCostReadLength field if non-nil, zero value otherwise.

### GetExecutionCostReadLengthOk

`func (o *Block) GetExecutionCostReadLengthOk() (*int32, bool)`

GetExecutionCostReadLengthOk returns a tuple with the ExecutionCostReadLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCostReadLength

`func (o *Block) SetExecutionCostReadLength(v int32)`

SetExecutionCostReadLength sets ExecutionCostReadLength field to given value.


### GetExecutionCostRuntime

`func (o *Block) GetExecutionCostRuntime() int32`

GetExecutionCostRuntime returns the ExecutionCostRuntime field if non-nil, zero value otherwise.

### GetExecutionCostRuntimeOk

`func (o *Block) GetExecutionCostRuntimeOk() (*int32, bool)`

GetExecutionCostRuntimeOk returns a tuple with the ExecutionCostRuntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCostRuntime

`func (o *Block) SetExecutionCostRuntime(v int32)`

SetExecutionCostRuntime sets ExecutionCostRuntime field to given value.


### GetExecutionCostWriteCount

`func (o *Block) GetExecutionCostWriteCount() int32`

GetExecutionCostWriteCount returns the ExecutionCostWriteCount field if non-nil, zero value otherwise.

### GetExecutionCostWriteCountOk

`func (o *Block) GetExecutionCostWriteCountOk() (*int32, bool)`

GetExecutionCostWriteCountOk returns a tuple with the ExecutionCostWriteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCostWriteCount

`func (o *Block) SetExecutionCostWriteCount(v int32)`

SetExecutionCostWriteCount sets ExecutionCostWriteCount field to given value.


### GetExecutionCostWriteLength

`func (o *Block) GetExecutionCostWriteLength() int32`

GetExecutionCostWriteLength returns the ExecutionCostWriteLength field if non-nil, zero value otherwise.

### GetExecutionCostWriteLengthOk

`func (o *Block) GetExecutionCostWriteLengthOk() (*int32, bool)`

GetExecutionCostWriteLengthOk returns a tuple with the ExecutionCostWriteLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCostWriteLength

`func (o *Block) SetExecutionCostWriteLength(v int32)`

SetExecutionCostWriteLength sets ExecutionCostWriteLength field to given value.


### GetTxs

`func (o *Block) GetTxs() []string`

GetTxs returns the Txs field if non-nil, zero value otherwise.

### GetTxsOk

`func (o *Block) GetTxsOk() (*[]string, bool)`

GetTxsOk returns a tuple with the Txs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxs

`func (o *Block) SetTxs(v []string)`

SetTxs sets Txs field to given value.


### GetParentMicroblockHash

`func (o *Block) GetParentMicroblockHash() string`

GetParentMicroblockHash returns the ParentMicroblockHash field if non-nil, zero value otherwise.

### GetParentMicroblockHashOk

`func (o *Block) GetParentMicroblockHashOk() (*string, bool)`

GetParentMicroblockHashOk returns a tuple with the ParentMicroblockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentMicroblockHash

`func (o *Block) SetParentMicroblockHash(v string)`

SetParentMicroblockHash sets ParentMicroblockHash field to given value.


### GetParentMicroblockSequence

`func (o *Block) GetParentMicroblockSequence() int32`

GetParentMicroblockSequence returns the ParentMicroblockSequence field if non-nil, zero value otherwise.

### GetParentMicroblockSequenceOk

`func (o *Block) GetParentMicroblockSequenceOk() (*int32, bool)`

GetParentMicroblockSequenceOk returns a tuple with the ParentMicroblockSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentMicroblockSequence

`func (o *Block) SetParentMicroblockSequence(v int32)`

SetParentMicroblockSequence sets ParentMicroblockSequence field to given value.


### GetMicroblocksAccepted

`func (o *Block) GetMicroblocksAccepted() []string`

GetMicroblocksAccepted returns the MicroblocksAccepted field if non-nil, zero value otherwise.

### GetMicroblocksAcceptedOk

`func (o *Block) GetMicroblocksAcceptedOk() (*[]string, bool)`

GetMicroblocksAcceptedOk returns a tuple with the MicroblocksAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroblocksAccepted

`func (o *Block) SetMicroblocksAccepted(v []string)`

SetMicroblocksAccepted sets MicroblocksAccepted field to given value.


### GetMicroblocksStreamed

`func (o *Block) GetMicroblocksStreamed() []string`

GetMicroblocksStreamed returns the MicroblocksStreamed field if non-nil, zero value otherwise.

### GetMicroblocksStreamedOk

`func (o *Block) GetMicroblocksStreamedOk() (*[]string, bool)`

GetMicroblocksStreamedOk returns a tuple with the MicroblocksStreamed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroblocksStreamed

`func (o *Block) SetMicroblocksStreamed(v []string)`

SetMicroblocksStreamed sets MicroblocksStreamed field to given value.


### GetMicroblockTxCount

`func (o *Block) GetMicroblockTxCount() map[string]int32`

GetMicroblockTxCount returns the MicroblockTxCount field if non-nil, zero value otherwise.

### GetMicroblockTxCountOk

`func (o *Block) GetMicroblockTxCountOk() (*map[string]int32, bool)`

GetMicroblockTxCountOk returns a tuple with the MicroblockTxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroblockTxCount

`func (o *Block) SetMicroblockTxCount(v map[string]int32)`

SetMicroblockTxCount sets MicroblockTxCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


