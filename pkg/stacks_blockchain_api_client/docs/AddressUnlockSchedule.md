# AddressUnlockSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Micro-STX amount locked at this block height. | 
**BlockHeight** | **int32** |  | 

## Methods

### NewAddressUnlockSchedule

`func NewAddressUnlockSchedule(amount string, blockHeight int32, ) *AddressUnlockSchedule`

NewAddressUnlockSchedule instantiates a new AddressUnlockSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressUnlockScheduleWithDefaults

`func NewAddressUnlockScheduleWithDefaults() *AddressUnlockSchedule`

NewAddressUnlockScheduleWithDefaults instantiates a new AddressUnlockSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *AddressUnlockSchedule) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AddressUnlockSchedule) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AddressUnlockSchedule) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetBlockHeight

`func (o *AddressUnlockSchedule) GetBlockHeight() int32`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *AddressUnlockSchedule) GetBlockHeightOk() (*int32, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *AddressUnlockSchedule) SetBlockHeight(v int32)`

SetBlockHeight sets BlockHeight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


