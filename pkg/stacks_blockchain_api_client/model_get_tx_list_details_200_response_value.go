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


// GetTxListDetails200ResponseValue struct for GetTxListDetails200ResponseValue
type GetTxListDetails200ResponseValue struct {
	TransactionFound *TransactionFound
	TransactionNotFound *TransactionNotFound
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *GetTxListDetails200ResponseValue) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into TransactionFound
	err = json.Unmarshal(data, &dst.TransactionFound);
	if err == nil {
		jsonTransactionFound, _ := json.Marshal(dst.TransactionFound)
		if string(jsonTransactionFound) == "{}" { // empty struct
			dst.TransactionFound = nil
		} else {
			return nil // data stored in dst.TransactionFound, return on the first match
		}
	} else {
		dst.TransactionFound = nil
	}

	// try to unmarshal JSON data into TransactionNotFound
	err = json.Unmarshal(data, &dst.TransactionNotFound);
	if err == nil {
		jsonTransactionNotFound, _ := json.Marshal(dst.TransactionNotFound)
		if string(jsonTransactionNotFound) == "{}" { // empty struct
			dst.TransactionNotFound = nil
		} else {
			return nil // data stored in dst.TransactionNotFound, return on the first match
		}
	} else {
		dst.TransactionNotFound = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(GetTxListDetails200ResponseValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *GetTxListDetails200ResponseValue) MarshalJSON() ([]byte, error) {
	if src.TransactionFound != nil {
		return json.Marshal(&src.TransactionFound)
	}

	if src.TransactionNotFound != nil {
		return json.Marshal(&src.TransactionNotFound)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableGetTxListDetails200ResponseValue struct {
	value *GetTxListDetails200ResponseValue
	isSet bool
}

func (v NullableGetTxListDetails200ResponseValue) Get() *GetTxListDetails200ResponseValue {
	return v.value
}

func (v *NullableGetTxListDetails200ResponseValue) Set(val *GetTxListDetails200ResponseValue) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTxListDetails200ResponseValue) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTxListDetails200ResponseValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTxListDetails200ResponseValue(val *GetTxListDetails200ResponseValue) *NullableGetTxListDetails200ResponseValue {
	return &NullableGetTxListDetails200ResponseValue{value: val, isSet: true}
}

func (v NullableGetTxListDetails200ResponseValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTxListDetails200ResponseValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


