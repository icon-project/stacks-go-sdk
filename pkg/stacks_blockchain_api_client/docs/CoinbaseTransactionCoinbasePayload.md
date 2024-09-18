# CoinbaseTransactionCoinbasePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **string** | Hex encoded 32-byte scratch space for block leader&#39;s use | 
**AltRecipient** | Pointer to **NullableString** | A principal that will receive the miner rewards for this coinbase transaction. Can be either a standard principal or contract principal. Only specified for &#x60;coinbase-to-alt-recipient&#x60; transaction types, otherwise null. | [optional] 
**VrfProof** | Pointer to **NullableString** | Hex encoded 80-byte VRF proof | [optional] 

## Methods

### NewCoinbaseTransactionCoinbasePayload

`func NewCoinbaseTransactionCoinbasePayload(data string, ) *CoinbaseTransactionCoinbasePayload`

NewCoinbaseTransactionCoinbasePayload instantiates a new CoinbaseTransactionCoinbasePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoinbaseTransactionCoinbasePayloadWithDefaults

`func NewCoinbaseTransactionCoinbasePayloadWithDefaults() *CoinbaseTransactionCoinbasePayload`

NewCoinbaseTransactionCoinbasePayloadWithDefaults instantiates a new CoinbaseTransactionCoinbasePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CoinbaseTransactionCoinbasePayload) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CoinbaseTransactionCoinbasePayload) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CoinbaseTransactionCoinbasePayload) SetData(v string)`

SetData sets Data field to given value.


### GetAltRecipient

`func (o *CoinbaseTransactionCoinbasePayload) GetAltRecipient() string`

GetAltRecipient returns the AltRecipient field if non-nil, zero value otherwise.

### GetAltRecipientOk

`func (o *CoinbaseTransactionCoinbasePayload) GetAltRecipientOk() (*string, bool)`

GetAltRecipientOk returns a tuple with the AltRecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltRecipient

`func (o *CoinbaseTransactionCoinbasePayload) SetAltRecipient(v string)`

SetAltRecipient sets AltRecipient field to given value.

### HasAltRecipient

`func (o *CoinbaseTransactionCoinbasePayload) HasAltRecipient() bool`

HasAltRecipient returns a boolean if a field has been set.

### SetAltRecipientNil

`func (o *CoinbaseTransactionCoinbasePayload) SetAltRecipientNil(b bool)`

 SetAltRecipientNil sets the value for AltRecipient to be an explicit nil

### UnsetAltRecipient
`func (o *CoinbaseTransactionCoinbasePayload) UnsetAltRecipient()`

UnsetAltRecipient ensures that no value is present for AltRecipient, not even an explicit nil
### GetVrfProof

`func (o *CoinbaseTransactionCoinbasePayload) GetVrfProof() string`

GetVrfProof returns the VrfProof field if non-nil, zero value otherwise.

### GetVrfProofOk

`func (o *CoinbaseTransactionCoinbasePayload) GetVrfProofOk() (*string, bool)`

GetVrfProofOk returns a tuple with the VrfProof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfProof

`func (o *CoinbaseTransactionCoinbasePayload) SetVrfProof(v string)`

SetVrfProof sets VrfProof field to given value.

### HasVrfProof

`func (o *CoinbaseTransactionCoinbasePayload) HasVrfProof() bool`

HasVrfProof returns a boolean if a field has been set.

### SetVrfProofNil

`func (o *CoinbaseTransactionCoinbasePayload) SetVrfProofNil(b bool)`

 SetVrfProofNil sets the value for VrfProof to be an explicit nil

### UnsetVrfProof
`func (o *CoinbaseTransactionCoinbasePayload) UnsetVrfProof()`

UnsetVrfProof ensures that no value is present for VrfProof, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


