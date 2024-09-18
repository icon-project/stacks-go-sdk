# AddressTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tx** | [**GetTransactionList200ResponseResultsInner**](GetTransactionList200ResponseResultsInner.md) |  | 
**StxSent** | **string** | Total sent from the given address, including the tx fee, in micro-STX as an integer string. | 
**StxReceived** | **string** | Total received by the given address in micro-STX as an integer string. | 
**Events** | [**AddressTransactionEvents**](AddressTransactionEvents.md) |  | 

## Methods

### NewAddressTransaction

`func NewAddressTransaction(tx GetTransactionList200ResponseResultsInner, stxSent string, stxReceived string, events AddressTransactionEvents, ) *AddressTransaction`

NewAddressTransaction instantiates a new AddressTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressTransactionWithDefaults

`func NewAddressTransactionWithDefaults() *AddressTransaction`

NewAddressTransactionWithDefaults instantiates a new AddressTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTx

`func (o *AddressTransaction) GetTx() GetTransactionList200ResponseResultsInner`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *AddressTransaction) GetTxOk() (*GetTransactionList200ResponseResultsInner, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *AddressTransaction) SetTx(v GetTransactionList200ResponseResultsInner)`

SetTx sets Tx field to given value.


### GetStxSent

`func (o *AddressTransaction) GetStxSent() string`

GetStxSent returns the StxSent field if non-nil, zero value otherwise.

### GetStxSentOk

`func (o *AddressTransaction) GetStxSentOk() (*string, bool)`

GetStxSentOk returns a tuple with the StxSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStxSent

`func (o *AddressTransaction) SetStxSent(v string)`

SetStxSent sets StxSent field to given value.


### GetStxReceived

`func (o *AddressTransaction) GetStxReceived() string`

GetStxReceived returns the StxReceived field if non-nil, zero value otherwise.

### GetStxReceivedOk

`func (o *AddressTransaction) GetStxReceivedOk() (*string, bool)`

GetStxReceivedOk returns a tuple with the StxReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStxReceived

`func (o *AddressTransaction) SetStxReceived(v string)`

SetStxReceived sets StxReceived field to given value.


### GetEvents

`func (o *AddressTransaction) GetEvents() AddressTransactionEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *AddressTransaction) GetEventsOk() (*AddressTransactionEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *AddressTransaction) SetEvents(v AddressTransactionEvents)`

SetEvents sets Events field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


