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


// NonFungibleTokenHistoryEvent Describes an event from the history of a Non-Fungible Token
type NonFungibleTokenHistoryEvent struct {
	NonFungibleTokenHistoryEventWithTxId *NonFungibleTokenHistoryEventWithTxId
	NonFungibleTokenHistoryEventWithTxMetadata *NonFungibleTokenHistoryEventWithTxMetadata
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NonFungibleTokenHistoryEvent) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NonFungibleTokenHistoryEventWithTxId
	err = json.Unmarshal(data, &dst.NonFungibleTokenHistoryEventWithTxId);
	if err == nil {
		jsonNonFungibleTokenHistoryEventWithTxId, _ := json.Marshal(dst.NonFungibleTokenHistoryEventWithTxId)
		if string(jsonNonFungibleTokenHistoryEventWithTxId) == "{}" { // empty struct
			dst.NonFungibleTokenHistoryEventWithTxId = nil
		} else {
			return nil // data stored in dst.NonFungibleTokenHistoryEventWithTxId, return on the first match
		}
	} else {
		dst.NonFungibleTokenHistoryEventWithTxId = nil
	}

	// try to unmarshal JSON data into NonFungibleTokenHistoryEventWithTxMetadata
	err = json.Unmarshal(data, &dst.NonFungibleTokenHistoryEventWithTxMetadata);
	if err == nil {
		jsonNonFungibleTokenHistoryEventWithTxMetadata, _ := json.Marshal(dst.NonFungibleTokenHistoryEventWithTxMetadata)
		if string(jsonNonFungibleTokenHistoryEventWithTxMetadata) == "{}" { // empty struct
			dst.NonFungibleTokenHistoryEventWithTxMetadata = nil
		} else {
			return nil // data stored in dst.NonFungibleTokenHistoryEventWithTxMetadata, return on the first match
		}
	} else {
		dst.NonFungibleTokenHistoryEventWithTxMetadata = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(NonFungibleTokenHistoryEvent)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NonFungibleTokenHistoryEvent) MarshalJSON() ([]byte, error) {
	if src.NonFungibleTokenHistoryEventWithTxId != nil {
		return json.Marshal(&src.NonFungibleTokenHistoryEventWithTxId)
	}

	if src.NonFungibleTokenHistoryEventWithTxMetadata != nil {
		return json.Marshal(&src.NonFungibleTokenHistoryEventWithTxMetadata)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableNonFungibleTokenHistoryEvent struct {
	value *NonFungibleTokenHistoryEvent
	isSet bool
}

func (v NullableNonFungibleTokenHistoryEvent) Get() *NonFungibleTokenHistoryEvent {
	return v.value
}

func (v *NullableNonFungibleTokenHistoryEvent) Set(val *NonFungibleTokenHistoryEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableNonFungibleTokenHistoryEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableNonFungibleTokenHistoryEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonFungibleTokenHistoryEvent(val *NonFungibleTokenHistoryEvent) *NullableNonFungibleTokenHistoryEvent {
	return &NullableNonFungibleTokenHistoryEvent{value: val, isSet: true}
}

func (v NullableNonFungibleTokenHistoryEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonFungibleTokenHistoryEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


