/*
Stacks Blockchain API

Welcome to the API reference overview for the [Stacks Blockchain API](https://docs.hiro.so/stacks-blockchain-api).        [Download Postman collection](https://hirosystems.github.io/stacks-blockchain-api/collection.json)

API version: v8.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stacks_blockchain_api_client

import (
	"encoding/json"
	"fmt"
)


// GetFilteredEventsAddressParameter struct for GetFilteredEventsAddressParameter
type GetFilteredEventsAddressParameter struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *GetFilteredEventsAddressParameter) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into String
	err = json.Unmarshal(data, &dst.String);
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(GetFilteredEventsAddressParameter)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *GetFilteredEventsAddressParameter) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableGetFilteredEventsAddressParameter struct {
	value *GetFilteredEventsAddressParameter
	isSet bool
}

func (v NullableGetFilteredEventsAddressParameter) Get() *GetFilteredEventsAddressParameter {
	return v.value
}

func (v *NullableGetFilteredEventsAddressParameter) Set(val *GetFilteredEventsAddressParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFilteredEventsAddressParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFilteredEventsAddressParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFilteredEventsAddressParameter(val *GetFilteredEventsAddressParameter) *NullableGetFilteredEventsAddressParameter {
	return &NullableGetFilteredEventsAddressParameter{value: val, isSet: true}
}

func (v NullableGetFilteredEventsAddressParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFilteredEventsAddressParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


