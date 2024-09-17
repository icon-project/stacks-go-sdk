# FeeRateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transaction** | **string** | A serialized transaction | 

## Methods

### NewFeeRateRequest

`func NewFeeRateRequest(transaction string, ) *FeeRateRequest`

NewFeeRateRequest instantiates a new FeeRateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeRateRequestWithDefaults

`func NewFeeRateRequestWithDefaults() *FeeRateRequest`

NewFeeRateRequestWithDefaults instantiates a new FeeRateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransaction

`func (o *FeeRateRequest) GetTransaction() string`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *FeeRateRequest) GetTransactionOk() (*string, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *FeeRateRequest) SetTransaction(v string)`

SetTransaction sets Transaction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


