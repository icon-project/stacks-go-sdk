# BurnchainReward

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Canonical** | **bool** | Set to &#x60;true&#x60; if block corresponds to the canonical burchchain tip | 
**BurnBlockHash** | **string** | The hash representing the burnchain block | 
**BurnBlockHeight** | **int32** | Height of the burnchain block | 
**BurnAmount** | **string** | The total amount of burnchain tokens burned for this burnchain block, in the smallest unit (e.g. satoshis for Bitcoin) | 
**RewardRecipient** | **string** | The recipient address that received the burnchain rewards, in the format native to the burnchain (e.g. B58 encoded for Bitcoin) | 
**RewardAmount** | **string** | The amount of burnchain tokens rewarded to the recipient, in the smallest unit (e.g. satoshis for Bitcoin) | 
**RewardIndex** | **int32** | The index position of the reward entry, useful for ordering when there&#39;s more than one recipient per burnchain block | 

## Methods

### NewBurnchainReward

`func NewBurnchainReward(canonical bool, burnBlockHash string, burnBlockHeight int32, burnAmount string, rewardRecipient string, rewardAmount string, rewardIndex int32, ) *BurnchainReward`

NewBurnchainReward instantiates a new BurnchainReward object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBurnchainRewardWithDefaults

`func NewBurnchainRewardWithDefaults() *BurnchainReward`

NewBurnchainRewardWithDefaults instantiates a new BurnchainReward object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanonical

`func (o *BurnchainReward) GetCanonical() bool`

GetCanonical returns the Canonical field if non-nil, zero value otherwise.

### GetCanonicalOk

`func (o *BurnchainReward) GetCanonicalOk() (*bool, bool)`

GetCanonicalOk returns a tuple with the Canonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonical

`func (o *BurnchainReward) SetCanonical(v bool)`

SetCanonical sets Canonical field to given value.


### GetBurnBlockHash

`func (o *BurnchainReward) GetBurnBlockHash() string`

GetBurnBlockHash returns the BurnBlockHash field if non-nil, zero value otherwise.

### GetBurnBlockHashOk

`func (o *BurnchainReward) GetBurnBlockHashOk() (*string, bool)`

GetBurnBlockHashOk returns a tuple with the BurnBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockHash

`func (o *BurnchainReward) SetBurnBlockHash(v string)`

SetBurnBlockHash sets BurnBlockHash field to given value.


### GetBurnBlockHeight

`func (o *BurnchainReward) GetBurnBlockHeight() int32`

GetBurnBlockHeight returns the BurnBlockHeight field if non-nil, zero value otherwise.

### GetBurnBlockHeightOk

`func (o *BurnchainReward) GetBurnBlockHeightOk() (*int32, bool)`

GetBurnBlockHeightOk returns a tuple with the BurnBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBlockHeight

`func (o *BurnchainReward) SetBurnBlockHeight(v int32)`

SetBurnBlockHeight sets BurnBlockHeight field to given value.


### GetBurnAmount

`func (o *BurnchainReward) GetBurnAmount() string`

GetBurnAmount returns the BurnAmount field if non-nil, zero value otherwise.

### GetBurnAmountOk

`func (o *BurnchainReward) GetBurnAmountOk() (*string, bool)`

GetBurnAmountOk returns a tuple with the BurnAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnAmount

`func (o *BurnchainReward) SetBurnAmount(v string)`

SetBurnAmount sets BurnAmount field to given value.


### GetRewardRecipient

`func (o *BurnchainReward) GetRewardRecipient() string`

GetRewardRecipient returns the RewardRecipient field if non-nil, zero value otherwise.

### GetRewardRecipientOk

`func (o *BurnchainReward) GetRewardRecipientOk() (*string, bool)`

GetRewardRecipientOk returns a tuple with the RewardRecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardRecipient

`func (o *BurnchainReward) SetRewardRecipient(v string)`

SetRewardRecipient sets RewardRecipient field to given value.


### GetRewardAmount

`func (o *BurnchainReward) GetRewardAmount() string`

GetRewardAmount returns the RewardAmount field if non-nil, zero value otherwise.

### GetRewardAmountOk

`func (o *BurnchainReward) GetRewardAmountOk() (*string, bool)`

GetRewardAmountOk returns a tuple with the RewardAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardAmount

`func (o *BurnchainReward) SetRewardAmount(v string)`

SetRewardAmount sets RewardAmount field to given value.


### GetRewardIndex

`func (o *BurnchainReward) GetRewardIndex() int32`

GetRewardIndex returns the RewardIndex field if non-nil, zero value otherwise.

### GetRewardIndexOk

`func (o *BurnchainReward) GetRewardIndexOk() (*int32, bool)`

GetRewardIndexOk returns a tuple with the RewardIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardIndex

`func (o *BurnchainReward) SetRewardIndex(v int32)`

SetRewardIndex sets RewardIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


