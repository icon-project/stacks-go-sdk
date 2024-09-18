# AddressTransactionEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**EventIndex** | **int32** |  | 
**Data** | [**AddressTransactionEventAnyOf2Data**](AddressTransactionEventAnyOf2Data.md) |  | 

## Methods

### NewAddressTransactionEvent

`func NewAddressTransactionEvent(type_ string, eventIndex int32, data AddressTransactionEventAnyOf2Data, ) *AddressTransactionEvent`

NewAddressTransactionEvent instantiates a new AddressTransactionEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressTransactionEventWithDefaults

`func NewAddressTransactionEventWithDefaults() *AddressTransactionEvent`

NewAddressTransactionEventWithDefaults instantiates a new AddressTransactionEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AddressTransactionEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddressTransactionEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddressTransactionEvent) SetType(v string)`

SetType sets Type field to given value.


### GetEventIndex

`func (o *AddressTransactionEvent) GetEventIndex() int32`

GetEventIndex returns the EventIndex field if non-nil, zero value otherwise.

### GetEventIndexOk

`func (o *AddressTransactionEvent) GetEventIndexOk() (*int32, bool)`

GetEventIndexOk returns a tuple with the EventIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIndex

`func (o *AddressTransactionEvent) SetEventIndex(v int32)`

SetEventIndex sets EventIndex field to given value.


### GetData

`func (o *AddressTransactionEvent) GetData() AddressTransactionEventAnyOf2Data`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AddressTransactionEvent) GetDataOk() (*AddressTransactionEventAnyOf2Data, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AddressTransactionEvent) SetData(v AddressTransactionEventAnyOf2Data)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


