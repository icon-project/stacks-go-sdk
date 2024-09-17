# BlockSearchResultBlockData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Canonical** | **bool** |  | 
**Hash** | **string** |  | 
**ParentBlockHash** | **string** |  | 
**BurnBlockTime** | **int32** |  | 
**Height** | **int32** |  | 

## Methods

### NewBlockSearchResultBlockData

`func NewBlockSearchResultBlockData(canonical bool, hash string, parentBlockHash string, burnBlockTime int32, height int32, ) *BlockSearchResultBlockData`

NewBlockSearchResultBlockData instantiates a new BlockSearchResultBlockData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockSearchResultBlockDataWithDefaults

`func NewBlockSearchResultBlockDataWithDefaults() *BlockSearchResultBlockData`

NewBlockSearchResultBlockDataWithDefaults instantiates a new BlockSearchResultBlockData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanonical

`func (o *BlockSearchResultBlockData) GetCanonical() bool`

GetCanonical returns the Canonical field if non-nil, zero value otherwise.

### GetCanonicalOk

`func (o *BlockSearchResultBlockData) GetCanonicalOk() (*bool, bool)`

GetCanonicalOk returns a tuple with the Canonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonical

`func (o *BlockSearchResultBlockData) SetCanonical(v bool)`

SetCanonical sets Canonical field to given value.


### GetHash

`func (o *BlockSearchResultBlockData) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *BlockSearchResultBlockData) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *BlockSearchResultBlockData) SetHash(v string)`

SetHash sets Hash field to given value.


### GetParentBlockHash

`func (o *BlockSearchResultBlockData) GetParentBlockHash() string`

GetParentBlockHash returns the ParentBlockHash field if non-nil, zero value otherwise.

### GetParentBlockHashOk

`func (o *BlockSearchResultBlockData) GetParentBlockHashOk() (*string, bool)`

GetParentBlockHashOk returns a tuple with the ParentBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBlockHash

`func (o *BlockSearchResultBlockData) SetParentBlockHash(v string)`

SetParentBlockHash sets ParentBlockHash field to given value.


### GetBurnBlockTime

`func (o *BlockSearchResultBlockData) GetBurnBlockTime() int32`

GetBurnBlockTime returns the BurnBlockTime field if non-nil, zero value otherwise.

### GetBurnBlockTimeOk

`func (o *BlockSearchResultBlockData) GetBurnBlockTimeOk() (*int32, bool)`

GetBurnBlockTimeOk returns a tuple with the BurnBlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockTime

`func (o *BlockSearchResultBlockData) SetBurnBlockTime(v int32)`

SetBurnBlockTime sets BurnBlockTime field to given value.


### GetHeight

`func (o *BlockSearchResultBlockData) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *BlockSearchResultBlockData) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *BlockSearchResultBlockData) SetHeight(v int32)`

SetHeight sets Height field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


