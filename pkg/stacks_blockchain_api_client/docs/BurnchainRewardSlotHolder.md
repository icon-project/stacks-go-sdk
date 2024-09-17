# BurnchainRewardSlotHolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Canonical** | **bool** | Set to &#x60;true&#x60; if block corresponds to the canonical burchchain tip | 
**BurnBlockHash** | **string** | The hash representing the burnchain block | 
**BurnBlockHeight** | **int32** | Height of the burnchain block | 
**Address** | **string** | The recipient address that validly received PoX commitments, in the format native to the burnchain (e.g. B58 encoded for Bitcoin) | 
**SlotIndex** | **int32** | The index position of the reward entry, useful for ordering when there&#39;s more than one slot per burnchain block | 

## Methods

### NewBurnchainRewardSlotHolder

`func NewBurnchainRewardSlotHolder(canonical bool, burnBlockHash string, burnBlockHeight int32, address string, slotIndex int32, ) *BurnchainRewardSlotHolder`

NewBurnchainRewardSlotHolder instantiates a new BurnchainRewardSlotHolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBurnchainRewardSlotHolderWithDefaults

`func NewBurnchainRewardSlotHolderWithDefaults() *BurnchainRewardSlotHolder`

NewBurnchainRewardSlotHolderWithDefaults instantiates a new BurnchainRewardSlotHolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanonical

`func (o *BurnchainRewardSlotHolder) GetCanonical() bool`

GetCanonical returns the Canonical field if non-nil, zero value otherwise.

### GetCanonicalOk

`func (o *BurnchainRewardSlotHolder) GetCanonicalOk() (*bool, bool)`

GetCanonicalOk returns a tuple with the Canonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonical

`func (o *BurnchainRewardSlotHolder) SetCanonical(v bool)`

SetCanonical sets Canonical field to given value.


### GetBurnBlockHash

`func (o *BurnchainRewardSlotHolder) GetBurnBlockHash() string`

GetBurnBlockHash returns the BurnBlockHash field if non-nil, zero value otherwise.

### GetBurnBlockHashOk

`func (o *BurnchainRewardSlotHolder) GetBurnBlockHashOk() (*string, bool)`

GetBurnBlockHashOk returns a tuple with the BurnBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockHash

`func (o *BurnchainRewardSlotHolder) SetBurnBlockHash(v string)`

SetBurnBlockHash sets BurnBlockHash field to given value.


### GetBurnBlockHeight

`func (o *BurnchainRewardSlotHolder) GetBurnBlockHeight() int32`

GetBurnBlockHeight returns the BurnBlockHeight field if non-nil, zero value otherwise.

### GetBurnBlockHeightOk

`func (o *BurnchainRewardSlotHolder) GetBurnBlockHeightOk() (*int32, bool)`

GetBurnBlockHeightOk returns a tuple with the BurnBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockHeight

`func (o *BurnchainRewardSlotHolder) SetBurnBlockHeight(v int32)`

SetBurnBlockHeight sets BurnBlockHeight field to given value.


### GetAddress

`func (o *BurnchainRewardSlotHolder) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BurnchainRewardSlotHolder) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BurnchainRewardSlotHolder) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetSlotIndex

`func (o *BurnchainRewardSlotHolder) GetSlotIndex() int32`

GetSlotIndex returns the SlotIndex field if non-nil, zero value otherwise.

### GetSlotIndexOk

`func (o *BurnchainRewardSlotHolder) GetSlotIndexOk() (*int32, bool)`

GetSlotIndexOk returns a tuple with the SlotIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotIndex

`func (o *BurnchainRewardSlotHolder) SetSlotIndex(v int32)`

SetSlotIndex sets SlotIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


