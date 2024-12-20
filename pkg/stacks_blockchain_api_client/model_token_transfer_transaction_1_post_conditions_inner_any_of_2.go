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

// checks if the TokenTransferTransaction1PostConditionsInnerAnyOf2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenTransferTransaction1PostConditionsInnerAnyOf2{}

// TokenTransferTransaction1PostConditionsInnerAnyOf2 struct for TokenTransferTransaction1PostConditionsInnerAnyOf2
type TokenTransferTransaction1PostConditionsInnerAnyOf2 struct {
	Principal TokenTransferTransaction1PostConditionsInnerAnyOfPrincipal `json:"principal"`
	ConditionCode TokenTransferTransaction1PostConditionMode `json:"condition_code"`
	Type string `json:"type"`
	AssetValue TokenTransferTransactionPostConditionsInnerAnyOf2AssetValue `json:"asset_value"`
	Asset TokenTransferTransactionPostConditionsInnerAnyOf1Asset `json:"asset"`
}

type _TokenTransferTransaction1PostConditionsInnerAnyOf2 TokenTransferTransaction1PostConditionsInnerAnyOf2

// NewTokenTransferTransaction1PostConditionsInnerAnyOf2 instantiates a new TokenTransferTransaction1PostConditionsInnerAnyOf2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenTransferTransaction1PostConditionsInnerAnyOf2(principal TokenTransferTransaction1PostConditionsInnerAnyOfPrincipal, conditionCode TokenTransferTransaction1PostConditionMode, type_ string, assetValue TokenTransferTransactionPostConditionsInnerAnyOf2AssetValue, asset TokenTransferTransactionPostConditionsInnerAnyOf1Asset) *TokenTransferTransaction1PostConditionsInnerAnyOf2 {
	this := TokenTransferTransaction1PostConditionsInnerAnyOf2{}
	this.Principal = principal
	this.ConditionCode = conditionCode
	this.Type = type_
	this.AssetValue = assetValue
	this.Asset = asset
	return &this
}

// NewTokenTransferTransaction1PostConditionsInnerAnyOf2WithDefaults instantiates a new TokenTransferTransaction1PostConditionsInnerAnyOf2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenTransferTransaction1PostConditionsInnerAnyOf2WithDefaults() *TokenTransferTransaction1PostConditionsInnerAnyOf2 {
	this := TokenTransferTransaction1PostConditionsInnerAnyOf2{}
	return &this
}

// GetPrincipal returns the Principal field value
func (o *TokenTransferTransaction1PostConditionsInnerAnyOf2) GetPrincipal() TokenTransferTransaction1PostConditionsInnerAnyOfPrincipal {
	if o == nil {
		var ret TokenTransferTransaction1PostConditionsInnerAnyOfPrincipal
		return ret
	}

	return o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction1PostConditionsInnerAnyOf2) GetPrincipalOk() (*TokenTransferTransaction1PostConditionsInnerAnyOfPrincipal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Principal, true
}

// SetPrincipal sets field value
func (o *TokenTransferTransaction1PostConditionsInnerAnyOf2) SetPrincipal(v TokenTransferTransaction1PostConditionsInnerAnyOfPrincipal) {
	o.Principal = v
}

// GetConditionCode returns the ConditionCode field value
func (o *TokenTransferTransaction1PostConditionsInnerAnyOf2) GetConditionCode() TokenTransferTransaction1PostConditionMode {
	if o == nil {
		var ret TokenTransferTransaction1PostConditionMode
		return ret
	}

	return o.ConditionCode
}

// GetConditionCodeOk returns a tuple with the ConditionCode field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction1PostConditionsInnerAnyOf2) GetConditionCodeOk() (*TokenTransferTransaction1PostConditionMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConditionCode, true
}

// SetConditionCode sets field value
func (o *TokenTransferTransaction1PostConditionsInnerAnyOf2) SetConditionCode(v TokenTransferTransaction1PostConditionMode) {
	o.ConditionCode = v
}

// GetType returns the Type field value
func (o *TokenTransferTransaction1PostConditionsInnerAnyOf2) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction1PostConditionsInnerAnyOf2) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TokenTransferTransaction1PostConditionsInnerAnyOf2) SetType(v string) {
	o.Type = v
}

// GetAssetValue returns the AssetValue field value
func (o *TokenTransferTransaction1PostConditionsInnerAnyOf2) GetAssetValue() TokenTransferTransactionPostConditionsInnerAnyOf2AssetValue {
	if o == nil {
		var ret TokenTransferTransactionPostConditionsInnerAnyOf2AssetValue
		return ret
	}

	return o.AssetValue
}

// GetAssetValueOk returns a tuple with the AssetValue field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction1PostConditionsInnerAnyOf2) GetAssetValueOk() (*TokenTransferTransactionPostConditionsInnerAnyOf2AssetValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetValue, true
}

// SetAssetValue sets field value
func (o *TokenTransferTransaction1PostConditionsInnerAnyOf2) SetAssetValue(v TokenTransferTransactionPostConditionsInnerAnyOf2AssetValue) {
	o.AssetValue = v
}

// GetAsset returns the Asset field value
func (o *TokenTransferTransaction1PostConditionsInnerAnyOf2) GetAsset() TokenTransferTransactionPostConditionsInnerAnyOf1Asset {
	if o == nil {
		var ret TokenTransferTransactionPostConditionsInnerAnyOf1Asset
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *TokenTransferTransaction1PostConditionsInnerAnyOf2) GetAssetOk() (*TokenTransferTransactionPostConditionsInnerAnyOf1Asset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *TokenTransferTransaction1PostConditionsInnerAnyOf2) SetAsset(v TokenTransferTransactionPostConditionsInnerAnyOf1Asset) {
	o.Asset = v
}

func (o TokenTransferTransaction1PostConditionsInnerAnyOf2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenTransferTransaction1PostConditionsInnerAnyOf2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["principal"] = o.Principal
	toSerialize["condition_code"] = o.ConditionCode
	toSerialize["type"] = o.Type
	toSerialize["asset_value"] = o.AssetValue
	toSerialize["asset"] = o.Asset
	return toSerialize, nil
}

func (o *TokenTransferTransaction1PostConditionsInnerAnyOf2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"principal",
		"condition_code",
		"type",
		"asset_value",
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

	varTokenTransferTransaction1PostConditionsInnerAnyOf2 := _TokenTransferTransaction1PostConditionsInnerAnyOf2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTokenTransferTransaction1PostConditionsInnerAnyOf2)

	if err != nil {
		return err
	}

	*o = TokenTransferTransaction1PostConditionsInnerAnyOf2(varTokenTransferTransaction1PostConditionsInnerAnyOf2)

	return err
}

type NullableTokenTransferTransaction1PostConditionsInnerAnyOf2 struct {
	value *TokenTransferTransaction1PostConditionsInnerAnyOf2
	isSet bool
}

func (v NullableTokenTransferTransaction1PostConditionsInnerAnyOf2) Get() *TokenTransferTransaction1PostConditionsInnerAnyOf2 {
	return v.value
}

func (v *NullableTokenTransferTransaction1PostConditionsInnerAnyOf2) Set(val *TokenTransferTransaction1PostConditionsInnerAnyOf2) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenTransferTransaction1PostConditionsInnerAnyOf2) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenTransferTransaction1PostConditionsInnerAnyOf2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenTransferTransaction1PostConditionsInnerAnyOf2(val *TokenTransferTransaction1PostConditionsInnerAnyOf2) *NullableTokenTransferTransaction1PostConditionsInnerAnyOf2 {
	return &NullableTokenTransferTransaction1PostConditionsInnerAnyOf2{value: val, isSet: true}
}

func (v NullableTokenTransferTransaction1PostConditionsInnerAnyOf2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenTransferTransaction1PostConditionsInnerAnyOf2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


