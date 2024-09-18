# SmartContractLogTransactionEvent1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventIndex** | **int32** |  | 
**EventType** | **string** |  | 
**TxId** | **string** |  | 
**ContractLog** | [**SmartContractLogTransactionEventAllOfContractLog**](SmartContractLogTransactionEventAllOfContractLog.md) |  | 

## Methods

### NewSmartContractLogTransactionEvent1

`func NewSmartContractLogTransactionEvent1(eventIndex int32, eventType string, txId string, contractLog SmartContractLogTransactionEventAllOfContractLog, ) *SmartContractLogTransactionEvent1`

NewSmartContractLogTransactionEvent1 instantiates a new SmartContractLogTransactionEvent1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartContractLogTransactionEvent1WithDefaults

`func NewSmartContractLogTransactionEvent1WithDefaults() *SmartContractLogTransactionEvent1`

NewSmartContractLogTransactionEvent1WithDefaults instantiates a new SmartContractLogTransactionEvent1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventIndex

`func (o *SmartContractLogTransactionEvent1) GetEventIndex() int32`

GetEventIndex returns the EventIndex field if non-nil, zero value otherwise.

### GetEventIndexOk

`func (o *SmartContractLogTransactionEvent1) GetEventIndexOk() (*int32, bool)`

GetEventIndexOk returns a tuple with the EventIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIndex

`func (o *SmartContractLogTransactionEvent1) SetEventIndex(v int32)`

SetEventIndex sets EventIndex field to given value.


### GetEventType

`func (o *SmartContractLogTransactionEvent1) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *SmartContractLogTransactionEvent1) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *SmartContractLogTransactionEvent1) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetTxId

`func (o *SmartContractLogTransactionEvent1) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *SmartContractLogTransactionEvent1) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *SmartContractLogTransactionEvent1) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetContractLog

`func (o *SmartContractLogTransactionEvent1) GetContractLog() SmartContractLogTransactionEventAllOfContractLog`

GetContractLog returns the ContractLog field if non-nil, zero value otherwise.

### GetContractLogOk

`func (o *SmartContractLogTransactionEvent1) GetContractLogOk() (*SmartContractLogTransactionEventAllOfContractLog, bool)`

GetContractLogOk returns a tuple with the ContractLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractLog

`func (o *SmartContractLogTransactionEvent1) SetContractLog(v SmartContractLogTransactionEventAllOfContractLog)`

SetContractLog sets ContractLog field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


