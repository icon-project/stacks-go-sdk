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

// checks if the TokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf{}

// TokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf struct for TokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf
type TokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf struct {
	TypeId string `json:"type_id"`
}

type _TokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf TokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf

// NewTokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf instantiates a new TokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf(typeId string) *TokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf {
	this := TokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf{}
	this.TypeId = typeId
	return &this
}

// NewTokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOfWithDefaults instantiates a new TokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOfWithDefaults() *TokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf {
	this := TokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf{}
	return &this
}

// GetTypeId returns the TypeId field value
func (o *TokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf) GetTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeId
}

// GetTypeIdOk returns a tuple with the TypeId field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf) GetTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeId, true
}

// SetTypeId sets field value
func (o *TokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf) SetTypeId(v string) {
	o.TypeId = v
}

func (o TokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type_id"] = o.TypeId
	return toSerialize, nil
}

func (o *TokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type_id",
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

	varTokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf := _TokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf)

	if err != nil {
		return err
	}

	*o = TokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf(varTokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf)

	return err
}

type NullableTokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf struct {
	value *TokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf
	isSet bool
}

func (v NullableTokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf) Get() *TokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf {
	return v.value
}

func (v *NullableTokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf) Set(val *TokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf(val *TokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf) *NullableTokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf {
	return &NullableTokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf{value: val, isSet: true}
}

func (v NullableTokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenTransferTransactionPostConditionsInnerAnyOfPrincipalAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


