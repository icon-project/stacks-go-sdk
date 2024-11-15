/*
Stacks Blockchain API

Welcome to the API reference overview for the [Stacks Blockchain API](https://docs.hiro.so/stacks-blockchain-api).        [Download Postman collection](https://hirosystems.github.io/stacks-blockchain-api/collection.json)

API version: v8.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stacks_blockchain_api_client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RunFaucetResponse1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunFaucetResponse1{}

// RunFaucetResponse1 POST request that initiates a transfer of tokens to a specified testnet address
type RunFaucetResponse1 struct {
	// Indicates if the faucet call was successful
	Success bool `json:"success"`
	// The transaction ID for the faucet call
	TxId string `json:"txId"`
	// Raw transaction in hex string representation
	TxRaw string `json:"txRaw"`
}

type _RunFaucetResponse1 RunFaucetResponse1

// NewRunFaucetResponse1 instantiates a new RunFaucetResponse1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunFaucetResponse1(success bool, txId string, txRaw string) *RunFaucetResponse1 {
	this := RunFaucetResponse1{}
	this.Success = success
	this.TxId = txId
	this.TxRaw = txRaw
	return &this
}

// NewRunFaucetResponse1WithDefaults instantiates a new RunFaucetResponse1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunFaucetResponse1WithDefaults() *RunFaucetResponse1 {
	this := RunFaucetResponse1{}
	return &this
}

// GetSuccess returns the Success field value
func (o *RunFaucetResponse1) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *RunFaucetResponse1) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *RunFaucetResponse1) SetSuccess(v bool) {
	o.Success = v
}

// GetTxId returns the TxId field value
func (o *RunFaucetResponse1) GetTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *RunFaucetResponse1) GetTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *RunFaucetResponse1) SetTxId(v string) {
	o.TxId = v
}

// GetTxRaw returns the TxRaw field value
func (o *RunFaucetResponse1) GetTxRaw() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxRaw
}

// GetTxRawOk returns a tuple with the TxRaw field value
// and a boolean to check if the value has been set.
func (o *RunFaucetResponse1) GetTxRawOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxRaw, true
}

// SetTxRaw sets field value
func (o *RunFaucetResponse1) SetTxRaw(v string) {
	o.TxRaw = v
}

func (o RunFaucetResponse1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunFaucetResponse1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	toSerialize["txId"] = o.TxId
	toSerialize["txRaw"] = o.TxRaw
	return toSerialize, nil
}

func (o *RunFaucetResponse1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"success",
		"txId",
		"txRaw",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRunFaucetResponse1 := _RunFaucetResponse1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRunFaucetResponse1)

	if err != nil {
		return err
	}

	*o = RunFaucetResponse1(varRunFaucetResponse1)

	return err
}

type NullableRunFaucetResponse1 struct {
	value *RunFaucetResponse1
	isSet bool
}

func (v NullableRunFaucetResponse1) Get() *RunFaucetResponse1 {
	return v.value
}

func (v *NullableRunFaucetResponse1) Set(val *RunFaucetResponse1) {
	v.value = val
	v.isSet = true
}

func (v NullableRunFaucetResponse1) IsSet() bool {
	return v.isSet
}

func (v *NullableRunFaucetResponse1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunFaucetResponse1(val *RunFaucetResponse1) *NullableRunFaucetResponse1 {
	return &NullableRunFaucetResponse1{value: val, isSet: true}
}

func (v NullableRunFaucetResponse1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunFaucetResponse1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


