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

// checks if the TokenTransferTransactionPostConditionsInnerAnyOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenTransferTransactionPostConditionsInnerAnyOf1{}

// TokenTransferTransactionPostConditionsInnerAnyOf1 struct for TokenTransferTransactionPostConditionsInnerAnyOf1
type TokenTransferTransactionPostConditionsInnerAnyOf1 struct {
	Principal TokenTransferTransactionPostConditionsInnerAnyOfPrincipal `json:"principal"`
	ConditionCode TokenTransferTransactionPostConditionsInnerAnyOfConditionCode `json:"condition_code"`
	Amount string `json:"amount"`
	Type string `json:"type"`
	Asset TokenTransferTransactionPostConditionsInnerAnyOf1Asset `json:"asset"`
}

type _TokenTransferTransactionPostConditionsInnerAnyOf1 TokenTransferTransactionPostConditionsInnerAnyOf1

// NewTokenTransferTransactionPostConditionsInnerAnyOf1 instantiates a new TokenTransferTransactionPostConditionsInnerAnyOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenTransferTransactionPostConditionsInnerAnyOf1(principal TokenTransferTransactionPostConditionsInnerAnyOfPrincipal, conditionCode TokenTransferTransactionPostConditionsInnerAnyOfConditionCode, amount string, type_ string, asset TokenTransferTransactionPostConditionsInnerAnyOf1Asset) *TokenTransferTransactionPostConditionsInnerAnyOf1 {
	this := TokenTransferTransactionPostConditionsInnerAnyOf1{}
	this.Principal = principal
	this.ConditionCode = conditionCode
	this.Amount = amount
	this.Type = type_
	this.Asset = asset
	return &this
}

// NewTokenTransferTransactionPostConditionsInnerAnyOf1WithDefaults instantiates a new TokenTransferTransactionPostConditionsInnerAnyOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenTransferTransactionPostConditionsInnerAnyOf1WithDefaults() *TokenTransferTransactionPostConditionsInnerAnyOf1 {
	this := TokenTransferTransactionPostConditionsInnerAnyOf1{}
	return &this
}

// GetPrincipal returns the Principal field value
func (o *TokenTransferTransactionPostConditionsInnerAnyOf1) GetPrincipal() TokenTransferTransactionPostConditionsInnerAnyOfPrincipal {
	if o == nil {
		var ret TokenTransferTransactionPostConditionsInnerAnyOfPrincipal
		return ret
	}

	return o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransactionPostConditionsInnerAnyOf1) GetPrincipalOk() (*TokenTransferTransactionPostConditionsInnerAnyOfPrincipal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Principal, true
}

// SetPrincipal sets field value
func (o *TokenTransferTransactionPostConditionsInnerAnyOf1) SetPrincipal(v TokenTransferTransactionPostConditionsInnerAnyOfPrincipal) {
	o.Principal = v
}

// GetConditionCode returns the ConditionCode field value
func (o *TokenTransferTransactionPostConditionsInnerAnyOf1) GetConditionCode() TokenTransferTransactionPostConditionsInnerAnyOfConditionCode {
	if o == nil {
		var ret TokenTransferTransactionPostConditionsInnerAnyOfConditionCode
		return ret
	}

	return o.ConditionCode
}

// GetConditionCodeOk returns a tuple with the ConditionCode field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransactionPostConditionsInnerAnyOf1) GetConditionCodeOk() (*TokenTransferTransactionPostConditionsInnerAnyOfConditionCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConditionCode, true
}

// SetConditionCode sets field value
func (o *TokenTransferTransactionPostConditionsInnerAnyOf1) SetConditionCode(v TokenTransferTransactionPostConditionsInnerAnyOfConditionCode) {
	o.ConditionCode = v
}

// GetAmount returns the Amount field value
func (o *TokenTransferTransactionPostConditionsInnerAnyOf1) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransactionPostConditionsInnerAnyOf1) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TokenTransferTransactionPostConditionsInnerAnyOf1) SetAmount(v string) {
	o.Amount = v
}

// GetType returns the Type field value
func (o *TokenTransferTransactionPostConditionsInnerAnyOf1) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransactionPostConditionsInnerAnyOf1) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TokenTransferTransactionPostConditionsInnerAnyOf1) SetType(v string) {
	o.Type = v
}

// GetAsset returns the Asset field value
func (o *TokenTransferTransactionPostConditionsInnerAnyOf1) GetAsset() TokenTransferTransactionPostConditionsInnerAnyOf1Asset {
	if o == nil {
		var ret TokenTransferTransactionPostConditionsInnerAnyOf1Asset
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransactionPostConditionsInnerAnyOf1) GetAssetOk() (*TokenTransferTransactionPostConditionsInnerAnyOf1Asset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *TokenTransferTransactionPostConditionsInnerAnyOf1) SetAsset(v TokenTransferTransactionPostConditionsInnerAnyOf1Asset) {
	o.Asset = v
}

func (o TokenTransferTransactionPostConditionsInnerAnyOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenTransferTransactionPostConditionsInnerAnyOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["principal"] = o.Principal
	toSerialize["condition_code"] = o.ConditionCode
	toSerialize["amount"] = o.Amount
	toSerialize["type"] = o.Type
	toSerialize["asset"] = o.Asset
	return toSerialize, nil
}

func (o *TokenTransferTransactionPostConditionsInnerAnyOf1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"principal",
		"condition_code",
		"amount",
		"type",
		"asset",
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

	varTokenTransferTransactionPostConditionsInnerAnyOf1 := _TokenTransferTransactionPostConditionsInnerAnyOf1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTokenTransferTransactionPostConditionsInnerAnyOf1)

	if err != nil {
		return err
	}

	*o = TokenTransferTransactionPostConditionsInnerAnyOf1(varTokenTransferTransactionPostConditionsInnerAnyOf1)

	return err
}

type NullableTokenTransferTransactionPostConditionsInnerAnyOf1 struct {
	value *TokenTransferTransactionPostConditionsInnerAnyOf1
	isSet bool
}

func (v NullableTokenTransferTransactionPostConditionsInnerAnyOf1) Get() *TokenTransferTransactionPostConditionsInnerAnyOf1 {
	return v.value
}

func (v *NullableTokenTransferTransactionPostConditionsInnerAnyOf1) Set(val *TokenTransferTransactionPostConditionsInnerAnyOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenTransferTransactionPostConditionsInnerAnyOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenTransferTransactionPostConditionsInnerAnyOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenTransferTransactionPostConditionsInnerAnyOf1(val *TokenTransferTransactionPostConditionsInnerAnyOf1) *NullableTokenTransferTransactionPostConditionsInnerAnyOf1 {
	return &NullableTokenTransferTransactionPostConditionsInnerAnyOf1{value: val, isSet: true}
}

func (v NullableTokenTransferTransactionPostConditionsInnerAnyOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenTransferTransactionPostConditionsInnerAnyOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


