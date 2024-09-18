# StxLockTransactionEvent1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventIndex** | **int32** |  | 
**EventType** | **string** |  | 
**TxId** | **string** |  | 
**StxLockEvent** | [**StxLockTransactionEventAllOfStxLockEvent**](StxLockTransactionEventAllOfStxLockEvent.md) |  | 

## Methods

### NewStxLockTransactionEvent1

`func NewStxLockTransactionEvent1(eventIndex int32, eventType string, txId string, stxLockEvent StxLockTransactionEventAllOfStxLockEvent, ) *StxLockTransactionEvent1`

NewStxLockTransactionEvent1 instantiates a new StxLockTransactionEvent1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStxLockTransactionEvent1WithDefaults

`func NewStxLockTransactionEvent1WithDefaults() *StxLockTransactionEvent1`

NewStxLockTransactionEvent1WithDefaults instantiates a new StxLockTransactionEvent1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventIndex

`func (o *StxLockTransactionEvent1) GetEventIndex() int32`

GetEventIndex returns the EventIndex field if non-nil, zero value otherwise.

### GetEventIndexOk

`func (o *StxLockTransactionEvent1) GetEventIndexOk() (*int32, bool)`

GetEventIndexOk returns a tuple with the EventIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIndex

`func (o *StxLockTransactionEvent1) SetEventIndex(v int32)`

SetEventIndex sets EventIndex field to given value.


### GetEventType

`func (o *StxLockTransactionEvent1) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *StxLockTransactionEvent1) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *StxLockTransactionEvent1) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetTxId

`func (o *StxLockTransactionEvent1) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *StxLockTransactionEvent1) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *StxLockTransactionEvent1) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetStxLockEvent

`func (o *StxLockTransactionEvent1) GetStxLockEvent() StxLockTransactionEventAllOfStxLockEvent`

GetStxLockEvent returns the StxLockEvent field if non-nil, zero value otherwise.

### GetStxLockEventOk

`func (o *StxLockTransactionEvent1) GetStxLockEventOk() (*StxLockTransactionEventAllOfStxLockEvent, bool)`

GetStxLockEventOk returns a tuple with the StxLockEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStxLockEvent

`func (o *StxLockTransactionEvent1) SetStxLockEvent(v StxLockTransactionEventAllOfStxLockEvent)`

SetStxLockEvent sets StxLockEvent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


