# BurnchainRewardsTotal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RewardRecipient** | **string** | The recipient address that received the burnchain rewards, in the format native to the burnchain (e.g. B58 encoded for Bitcoin) | 
**RewardAmount** | **string** | The total amount of burnchain tokens rewarded to the recipient, in the smallest unit (e.g. satoshis for Bitcoin) | 

## Methods

### NewBurnchainRewardsTotal

`func NewBurnchainRewardsTotal(rewardRecipient string, rewardAmount string, ) *BurnchainRewardsTotal`

NewBurnchainRewardsTotal instantiates a new BurnchainRewardsTotal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBurnchainRewardsTotalWithDefaults

`func NewBurnchainRewardsTotalWithDefaults() *BurnchainRewardsTotal`

NewBurnchainRewardsTotalWithDefaults instantiates a new BurnchainRewardsTotal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRewardRecipient

`func (o *BurnchainRewardsTotal) GetRewardRecipient() string`

GetRewardRecipient returns the RewardRecipient field if non-nil, zero value otherwise.

### GetRewardRecipientOk

`func (o *BurnchainRewardsTotal) GetRewardRecipientOk() (*string, bool)`

GetRewardRecipientOk returns a tuple with the RewardRecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardRecipient

`func (o *BurnchainRewardsTotal) SetRewardRecipient(v string)`

SetRewardRecipient sets RewardRecipient field to given value.


### GetRewardAmount

`func (o *BurnchainRewardsTotal) GetRewardAmount() string`

GetRewardAmount returns the RewardAmount field if non-nil, zero value otherwise.

### GetRewardAmountOk

`func (o *BurnchainRewardsTotal) GetRewardAmountOk() (*string, bool)`

GetRewardAmountOk returns a tuple with the RewardAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardAmount

`func (o *BurnchainRewardsTotal) SetRewardAmount(v string)`

SetRewardAmount sets RewardAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


