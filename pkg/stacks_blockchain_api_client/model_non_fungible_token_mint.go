/*
Stacks Blockchain API

Welcome to the API reference overview for the [Stacks Blockchain API](https://docs.hiro.so/stacks-blockchain-api).        [Download Postman collection](https://hirosystems.github.io/stacks-blockchain-api/collection.json)

API version: v7.14.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stacks_blockchain_api_client

import (
	"encoding/json"
	"fmt"
)


// NonFungibleTokenMint Describes the minting of a Non-Fungible Token
type NonFungibleTokenMint struct {
	NonFungibleTokenMintAnyOf *NonFungibleTokenMintAnyOf
	NonFungibleTokenMintWithTxId *NonFungibleTokenMintWithTxId
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NonFungibleTokenMint) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NonFungibleTokenMintAnyOf
	err = json.Unmarshal(data, &dst.NonFungibleTokenMintAnyOf);
	if err == nil {
		jsonNonFungibleTokenMintAnyOf, _ := json.Marshal(dst.NonFungibleTokenMintAnyOf)
		if string(jsonNonFungibleTokenMintAnyOf) == "{}" { // empty struct
			dst.NonFungibleTokenMintAnyOf = nil
		} else {
			return nil // data stored in dst.NonFungibleTokenMintAnyOf, return on the first match
		}
	} else {
		dst.NonFungibleTokenMintAnyOf = nil
	}

	// try to unmarshal JSON data into NonFungibleTokenMintWithTxId
	err = json.Unmarshal(data, &dst.NonFungibleTokenMintWithTxId);
	if err == nil {
		jsonNonFungibleTokenMintWithTxId, _ := json.Marshal(dst.NonFungibleTokenMintWithTxId)
		if string(jsonNonFungibleTokenMintWithTxId) == "{}" { // empty struct
			dst.NonFungibleTokenMintWithTxId = nil
		} else {
			return nil // data stored in dst.NonFungibleTokenMintWithTxId, return on the first match
		}
	} else {
		dst.NonFungibleTokenMintWithTxId = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(NonFungibleTokenMint)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NonFungibleTokenMint) MarshalJSON() ([]byte, error) {
	if src.NonFungibleTokenMintAnyOf != nil {
		return json.Marshal(&src.NonFungibleTokenMintAnyOf)
	}

	if src.NonFungibleTokenMintWithTxId != nil {
		return json.Marshal(&src.NonFungibleTokenMintWithTxId)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableNonFungibleTokenMint struct {
	value *NonFungibleTokenMint
	isSet bool
}

func (v NullableNonFungibleTokenMint) Get() *NonFungibleTokenMint {
	return v.value
}

func (v *NullableNonFungibleTokenMint) Set(val *NonFungibleTokenMint) {
	v.value = val
	v.isSet = true
}

func (v NullableNonFungibleTokenMint) IsSet() bool {
	return v.isSet
}

func (v *NullableNonFungibleTokenMint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonFungibleTokenMint(val *NonFungibleTokenMint) *NullableNonFungibleTokenMint {
	return &NullableNonFungibleTokenMint{value: val, isSet: true}
}

func (v NullableNonFungibleTokenMint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonFungibleTokenMint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

