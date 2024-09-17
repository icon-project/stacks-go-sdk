# ListOfEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** |  | 
**Offset** | **int32** |  | 
**Events** | [**[]TokenTransferTransactionEventsInner**](TokenTransferTransactionEventsInner.md) |  | 

## Methods

### NewListOfEvents

`func NewListOfEvents(limit int32, offset int32, events []TokenTransferTransactionEventsInner, ) *ListOfEvents`

NewListOfEvents instantiates a new ListOfEvents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOfEventsWithDefaults

`func NewListOfEventsWithDefaults() *ListOfEvents`

NewListOfEventsWithDefaults instantiates a new ListOfEvents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *ListOfEvents) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListOfEvents) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListOfEvents) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *ListOfEvents) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListOfEvents) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListOfEvents) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetEvents

`func (o *ListOfEvents) GetEvents() []TokenTransferTransactionEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ListOfEvents) GetEventsOk() (*[]TokenTransferTransactionEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ListOfEvents) SetEvents(v []TokenTransferTransactionEventsInner)`

SetEvents sets Events field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


