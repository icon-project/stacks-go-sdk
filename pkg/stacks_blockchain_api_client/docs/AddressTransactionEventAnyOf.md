# AddressTransactionEventAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**EventIndex** | **int32** |  | 
**Data** | [**AddressTransactionEventAnyOfData**](AddressTransactionEventAnyOfData.md) |  | 

## Methods

### NewAddressTransactionEventAnyOf

`func NewAddressTransactionEventAnyOf(type_ string, eventIndex int32, data AddressTransactionEventAnyOfData, ) *AddressTransactionEventAnyOf`

NewAddressTransactionEventAnyOf instantiates a new AddressTransactionEventAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressTransactionEventAnyOfWithDefaults

`func NewAddressTransactionEventAnyOfWithDefaults() *AddressTransactionEventAnyOf`

NewAddressTransactionEventAnyOfWithDefaults instantiates a new AddressTransactionEventAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AddressTransactionEventAnyOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddressTransactionEventAnyOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddressTransactionEventAnyOf) SetType(v string)`

SetType sets Type field to given value.


### GetEventIndex

`func (o *AddressTransactionEventAnyOf) GetEventIndex() int32`

GetEventIndex returns the EventIndex field if non-nil, zero value otherwise.

### GetEventIndexOk

`func (o *AddressTransactionEventAnyOf) GetEventIndexOk() (*int32, bool)`

GetEventIndexOk returns a tuple with the EventIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIndex

`func (o *AddressTransactionEventAnyOf) SetEventIndex(v int32)`

SetEventIndex sets EventIndex field to given value.


### GetData

`func (o *AddressTransactionEventAnyOf) GetData() AddressTransactionEventAnyOfData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AddressTransactionEventAnyOf) GetDataOk() (*AddressTransactionEventAnyOfData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AddressTransactionEventAnyOf) SetData(v AddressTransactionEventAnyOfData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


