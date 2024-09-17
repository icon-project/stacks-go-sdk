# AddressTokenOfferingLocked

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalLocked** | **string** | Micro-STX amount still locked at current block height. | 
**TotalUnlocked** | **string** | Micro-STX amount unlocked at current block height. | 
**UnlockSchedule** | [**[]AddressUnlockSchedule**](AddressUnlockSchedule.md) |  | 

## Methods

### NewAddressTokenOfferingLocked

`func NewAddressTokenOfferingLocked(totalLocked string, totalUnlocked string, unlockSchedule []AddressUnlockSchedule, ) *AddressTokenOfferingLocked`

NewAddressTokenOfferingLocked instantiates a new AddressTokenOfferingLocked object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressTokenOfferingLockedWithDefaults

`func NewAddressTokenOfferingLockedWithDefaults() *AddressTokenOfferingLocked`

NewAddressTokenOfferingLockedWithDefaults instantiates a new AddressTokenOfferingLocked object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalLocked

`func (o *AddressTokenOfferingLocked) GetTotalLocked() string`

GetTotalLocked returns the TotalLocked field if non-nil, zero value otherwise.

### GetTotalLockedOk

`func (o *AddressTokenOfferingLocked) GetTotalLockedOk() (*string, bool)`

GetTotalLockedOk returns a tuple with the TotalLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLocked

`func (o *AddressTokenOfferingLocked) SetTotalLocked(v string)`

SetTotalLocked sets TotalLocked field to given value.


### GetTotalUnlocked

`func (o *AddressTokenOfferingLocked) GetTotalUnlocked() string`

GetTotalUnlocked returns the TotalUnlocked field if non-nil, zero value otherwise.

### GetTotalUnlockedOk

`func (o *AddressTokenOfferingLocked) GetTotalUnlockedOk() (*string, bool)`

GetTotalUnlockedOk returns a tuple with the TotalUnlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUnlocked

`func (o *AddressTokenOfferingLocked) SetTotalUnlocked(v string)`

SetTotalUnlocked sets TotalUnlocked field to given value.


### GetUnlockSchedule

`func (o *AddressTokenOfferingLocked) GetUnlockSchedule() []AddressUnlockSchedule`

GetUnlockSchedule returns the UnlockSchedule field if non-nil, zero value otherwise.

### GetUnlockScheduleOk

`func (o *AddressTokenOfferingLocked) GetUnlockScheduleOk() (*[]AddressUnlockSchedule, bool)`

GetUnlockScheduleOk returns a tuple with the UnlockSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockSchedule

`func (o *AddressTokenOfferingLocked) SetUnlockSchedule(v []AddressUnlockSchedule)`

SetUnlockSchedule sets UnlockSchedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


