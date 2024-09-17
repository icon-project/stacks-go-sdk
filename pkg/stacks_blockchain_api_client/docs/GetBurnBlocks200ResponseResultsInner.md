# GetBurnBlocks200ResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BurnBlockTime** | **int32** | Unix timestamp (in seconds) indicating when this block was mined. | 
**BurnBlockTimeIso** | **string** | An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) indicating when this block was mined. | 
**BurnBlockHash** | **string** | Hash of the anchor chain block | 
**BurnBlockHeight** | **int32** | Height of the anchor chain block | 
**StacksBlocks** | **[]string** | Hashes of the Stacks blocks included in the burn block | 
**AvgBlockTime** | **int32** | Average time between blocks in seconds. Returns 0 if there is only one block in the burn block. | 
**TotalTxCount** | **int32** | Total number of transactions in the Stacks blocks associated with this burn block | 

## Methods

### NewGetBurnBlocks200ResponseResultsInner

`func NewGetBurnBlocks200ResponseResultsInner(burnBlockTime int32, burnBlockTimeIso string, burnBlockHash string, burnBlockHeight int32, stacksBlocks []string, avgBlockTime int32, totalTxCount int32, ) *GetBurnBlocks200ResponseResultsInner`

NewGetBurnBlocks200ResponseResultsInner instantiates a new GetBurnBlocks200ResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBurnBlocks200ResponseResultsInnerWithDefaults

`func NewGetBurnBlocks200ResponseResultsInnerWithDefaults() *GetBurnBlocks200ResponseResultsInner`

NewGetBurnBlocks200ResponseResultsInnerWithDefaults instantiates a new GetBurnBlocks200ResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBurnBlockTime

`func (o *GetBurnBlocks200ResponseResultsInner) GetBurnBlockTime() int32`

GetBurnBlockTime returns the BurnBlockTime field if non-nil, zero value otherwise.

### GetBurnBlockTimeOk

`func (o *GetBurnBlocks200ResponseResultsInner) GetBurnBlockTimeOk() (*int32, bool)`

GetBurnBlockTimeOk returns a tuple with the BurnBlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockTime

`func (o *GetBurnBlocks200ResponseResultsInner) SetBurnBlockTime(v int32)`

SetBurnBlockTime sets BurnBlockTime field to given value.


### GetBurnBlockTimeIso

`func (o *GetBurnBlocks200ResponseResultsInner) GetBurnBlockTimeIso() string`

GetBurnBlockTimeIso returns the BurnBlockTimeIso field if non-nil, zero value otherwise.

### GetBurnBlockTimeIsoOk

`func (o *GetBurnBlocks200ResponseResultsInner) GetBurnBlockTimeIsoOk() (*string, bool)`

GetBurnBlockTimeIsoOk returns a tuple with the BurnBlockTimeIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockTimeIso

`func (o *GetBurnBlocks200ResponseResultsInner) SetBurnBlockTimeIso(v string)`

SetBurnBlockTimeIso sets BurnBlockTimeIso field to given value.


### GetBurnBlockHash

`func (o *GetBurnBlocks200ResponseResultsInner) GetBurnBlockHash() string`

GetBurnBlockHash returns the BurnBlockHash field if non-nil, zero value otherwise.

### GetBurnBlockHashOk

`func (o *GetBurnBlocks200ResponseResultsInner) GetBurnBlockHashOk() (*string, bool)`

GetBurnBlockHashOk returns a tuple with the BurnBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockHash

`func (o *GetBurnBlocks200ResponseResultsInner) SetBurnBlockHash(v string)`

SetBurnBlockHash sets BurnBlockHash field to given value.


### GetBurnBlockHeight

`func (o *GetBurnBlocks200ResponseResultsInner) GetBurnBlockHeight() int32`

GetBurnBlockHeight returns the BurnBlockHeight field if non-nil, zero value otherwise.

### GetBurnBlockHeightOk

`func (o *GetBurnBlocks200ResponseResultsInner) GetBurnBlockHeightOk() (*int32, bool)`

GetBurnBlockHeightOk returns a tuple with the BurnBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockHeight

`func (o *GetBurnBlocks200ResponseResultsInner) SetBurnBlockHeight(v int32)`

SetBurnBlockHeight sets BurnBlockHeight field to given value.


### GetStacksBlocks

`func (o *GetBurnBlocks200ResponseResultsInner) GetStacksBlocks() []string`

GetStacksBlocks returns the StacksBlocks field if non-nil, zero value otherwise.

### GetStacksBlocksOk

`func (o *GetBurnBlocks200ResponseResultsInner) GetStacksBlocksOk() (*[]string, bool)`

GetStacksBlocksOk returns a tuple with the StacksBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacksBlocks

`func (o *GetBurnBlocks200ResponseResultsInner) SetStacksBlocks(v []string)`

SetStacksBlocks sets StacksBlocks field to given value.


### GetAvgBlockTime

`func (o *GetBurnBlocks200ResponseResultsInner) GetAvgBlockTime() int32`

GetAvgBlockTime returns the AvgBlockTime field if non-nil, zero value otherwise.

### GetAvgBlockTimeOk

`func (o *GetBurnBlocks200ResponseResultsInner) GetAvgBlockTimeOk() (*int32, bool)`

GetAvgBlockTimeOk returns a tuple with the AvgBlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgBlockTime

`func (o *GetBurnBlocks200ResponseResultsInner) SetAvgBlockTime(v int32)`

SetAvgBlockTime sets AvgBlockTime field to given value.


### GetTotalTxCount

`func (o *GetBurnBlocks200ResponseResultsInner) GetTotalTxCount() int32`

GetTotalTxCount returns the TotalTxCount field if non-nil, zero value otherwise.

### GetTotalTxCountOk

`func (o *GetBurnBlocks200ResponseResultsInner) GetTotalTxCountOk() (*int32, bool)`

GetTotalTxCountOk returns a tuple with the TotalTxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTxCount

`func (o *GetBurnBlocks200ResponseResultsInner) SetTotalTxCount(v int32)`

SetTotalTxCount sets TotalTxCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


