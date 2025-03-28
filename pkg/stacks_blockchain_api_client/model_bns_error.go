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

// checks if the BnsError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BnsError{}

// BnsError Error
type BnsError struct {
	Error string `json:"error"`
}

type _BnsError BnsError

// NewBnsError instantiates a new BnsError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBnsError(error_ string) *BnsError {
	this := BnsError{}
	this.Error = error_
	return &this
}

// NewBnsErrorWithDefaults instantiates a new BnsError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBnsErrorWithDefaults() *BnsError {
	this := BnsError{}
	return &this
}

// GetError returns the Error field value
func (o *BnsError) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *BnsError) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *BnsError) SetError(v string) {
	o.Error = v
}

func (o BnsError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BnsError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

func (o *BnsError) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"error",
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

	varBnsError := _BnsError{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBnsError)

	if err != nil {
		return err
	}

	*o = BnsError(varBnsError)

	return err
}

type NullableBnsError struct {
	value *BnsError
	isSet bool
}

func (v NullableBnsError) Get() *BnsError {
	return v.value
}

func (v *NullableBnsError) Set(val *BnsError) {
	v.value = val
	v.isSet = true
}

func (v NullableBnsError) IsSet() bool {
	return v.isSet
}

func (v *NullableBnsError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBnsError(val *BnsError) *NullableBnsError {
	return &NullableBnsError{value: val, isSet: true}
}

func (v NullableBnsError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBnsError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


