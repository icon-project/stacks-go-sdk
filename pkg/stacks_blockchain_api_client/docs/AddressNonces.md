# AddressNonces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastMempoolTxNonce** | **NullableInt32** | The latest nonce found within mempool transactions sent by this address. Will be null if there are no current mempool transactions for this address. | 
**LastExecutedTxNonce** | **NullableInt32** | The latest nonce found within transactions sent by this address, including unanchored microblock transactions. Will be null if there are no current transactions for this address. | 
**PossibleNextNonce** | **int32** | The likely nonce required for creating the next transaction, based on the last nonces seen by the API. This can be incorrect if the API&#39;s mempool or transactions aren&#39;t fully synchronized, even by a small amount, or if a previous transaction is still propagating through the Stacks blockchain network when this endpoint is called. | 
**DetectedMissingNonces** | **[]int32** | Nonces that appear to be missing and likely causing a mempool transaction to be stuck. | 
**DetectedMempoolNonces** | **[]int32** | Nonces currently in mempool for this address. | 

## Methods

### NewAddressNonces

`func NewAddressNonces(lastMempoolTxNonce NullableInt32, lastExecutedTxNonce NullableInt32, possibleNextNonce int32, detectedMissingNonces []int32, detectedMempoolNonces []int32, ) *AddressNonces`

NewAddressNonces instantiates a new AddressNonces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressNoncesWithDefaults

`func NewAddressNoncesWithDefaults() *AddressNonces`

NewAddressNoncesWithDefaults instantiates a new AddressNonces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastMempoolTxNonce

`func (o *AddressNonces) GetLastMempoolTxNonce() int32`

GetLastMempoolTxNonce returns the LastMempoolTxNonce field if non-nil, zero value otherwise.

### GetLastMempoolTxNonceOk

`func (o *AddressNonces) GetLastMempoolTxNonceOk() (*int32, bool)`

GetLastMempoolTxNonceOk returns a tuple with the LastMempoolTxNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMempoolTxNonce

`func (o *AddressNonces) SetLastMempoolTxNonce(v int32)`

SetLastMempoolTxNonce sets LastMempoolTxNonce field to given value.


### SetLastMempoolTxNonceNil

`func (o *AddressNonces) SetLastMempoolTxNonceNil(b bool)`

 SetLastMempoolTxNonceNil sets the value for LastMempoolTxNonce to be an explicit nil

### UnsetLastMempoolTxNonce
`func (o *AddressNonces) UnsetLastMempoolTxNonce()`

UnsetLastMempoolTxNonce ensures that no value is present for LastMempoolTxNonce, not even an explicit nil
### GetLastExecutedTxNonce

`func (o *AddressNonces) GetLastExecutedTxNonce() int32`

GetLastExecutedTxNonce returns the LastExecutedTxNonce field if non-nil, zero value otherwise.

### GetLastExecutedTxNonceOk

`func (o *AddressNonces) GetLastExecutedTxNonceOk() (*int32, bool)`

GetLastExecutedTxNonceOk returns a tuple with the LastExecutedTxNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutedTxNonce

`func (o *AddressNonces) SetLastExecutedTxNonce(v int32)`

SetLastExecutedTxNonce sets LastExecutedTxNonce field to given value.


### SetLastExecutedTxNonceNil

`func (o *AddressNonces) SetLastExecutedTxNonceNil(b bool)`

 SetLastExecutedTxNonceNil sets the value for LastExecutedTxNonce to be an explicit nil

### UnsetLastExecutedTxNonce
`func (o *AddressNonces) UnsetLastExecutedTxNonce()`

UnsetLastExecutedTxNonce ensures that no value is present for LastExecutedTxNonce, not even an explicit nil
### GetPossibleNextNonce

`func (o *AddressNonces) GetPossibleNextNonce() int32`

GetPossibleNextNonce returns the PossibleNextNonce field if non-nil, zero value otherwise.

### GetPossibleNextNonceOk

`func (o *AddressNonces) GetPossibleNextNonceOk() (*int32, bool)`

GetPossibleNextNonceOk returns a tuple with the PossibleNextNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossibleNextNonce

`func (o *AddressNonces) SetPossibleNextNonce(v int32)`

SetPossibleNextNonce sets PossibleNextNonce field to given value.


### GetDetectedMissingNonces

`func (o *AddressNonces) GetDetectedMissingNonces() []int32`

GetDetectedMissingNonces returns the DetectedMissingNonces field if non-nil, zero value otherwise.

### GetDetectedMissingNoncesOk

`func (o *AddressNonces) GetDetectedMissingNoncesOk() (*[]int32, bool)`

GetDetectedMissingNoncesOk returns a tuple with the DetectedMissingNonces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedMissingNonces

`func (o *AddressNonces) SetDetectedMissingNonces(v []int32)`

SetDetectedMissingNonces sets DetectedMissingNonces field to given value.


### GetDetectedMempoolNonces

`func (o *AddressNonces) GetDetectedMempoolNonces() []int32`

GetDetectedMempoolNonces returns the DetectedMempoolNonces field if non-nil, zero value otherwise.

### GetDetectedMempoolNoncesOk

`func (o *AddressNonces) GetDetectedMempoolNoncesOk() (*[]int32, bool)`

GetDetectedMempoolNoncesOk returns a tuple with the DetectedMempoolNonces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedMempoolNonces

`func (o *AddressNonces) SetDetectedMempoolNonces(v []int32)`

SetDetectedMempoolNonces sets DetectedMempoolNonces field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


