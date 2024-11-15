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

// checks if the NetworkBlockTimesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkBlockTimesResponse{}

// NetworkBlockTimesResponse GET request that returns network target block times
type NetworkBlockTimesResponse struct {
	Mainnet NetworkBlockTimesResponseMainnet `json:"mainnet"`
	Testnet NetworkBlockTimesResponseMainnet `json:"testnet"`
}

type _NetworkBlockTimesResponse NetworkBlockTimesResponse

// NewNetworkBlockTimesResponse instantiates a new NetworkBlockTimesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkBlockTimesResponse(mainnet NetworkBlockTimesResponseMainnet, testnet NetworkBlockTimesResponseMainnet) *NetworkBlockTimesResponse {
	this := NetworkBlockTimesResponse{}
	this.Mainnet = mainnet
	this.Testnet = testnet
	return &this
}

// NewNetworkBlockTimesResponseWithDefaults instantiates a new NetworkBlockTimesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkBlockTimesResponseWithDefaults() *NetworkBlockTimesResponse {
	this := NetworkBlockTimesResponse{}
	return &this
}

// GetMainnet returns the Mainnet field value
func (o *NetworkBlockTimesResponse) GetMainnet() NetworkBlockTimesResponseMainnet {
	if o == nil {
		var ret NetworkBlockTimesResponseMainnet
		return ret
	}

	return o.Mainnet
}

// GetMainnetOk returns a tuple with the Mainnet field value
// and a boolean to check if the value has been set.
func (o *NetworkBlockTimesResponse) GetMainnetOk() (*NetworkBlockTimesResponseMainnet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mainnet, true
}

// SetMainnet sets field value
func (o *NetworkBlockTimesResponse) SetMainnet(v NetworkBlockTimesResponseMainnet) {
	o.Mainnet = v
}

// GetTestnet returns the Testnet field value
func (o *NetworkBlockTimesResponse) GetTestnet() NetworkBlockTimesResponseMainnet {
	if o == nil {
		var ret NetworkBlockTimesResponseMainnet
		return ret
	}

	return o.Testnet
}

// GetTestnetOk returns a tuple with the Testnet field value
// and a boolean to check if the value has been set.
func (o *NetworkBlockTimesResponse) GetTestnetOk() (*NetworkBlockTimesResponseMainnet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Testnet, true
}

// SetTestnet sets field value
func (o *NetworkBlockTimesResponse) SetTestnet(v NetworkBlockTimesResponseMainnet) {
	o.Testnet = v
}

func (o NetworkBlockTimesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkBlockTimesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mainnet"] = o.Mainnet
	toSerialize["testnet"] = o.Testnet
	return toSerialize, nil
}

func (o *NetworkBlockTimesResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mainnet",
		"testnet",
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

	varNetworkBlockTimesResponse := _NetworkBlockTimesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNetworkBlockTimesResponse)

	if err != nil {
		return err
	}

	*o = NetworkBlockTimesResponse(varNetworkBlockTimesResponse)

	return err
}

type NullableNetworkBlockTimesResponse struct {
	value *NetworkBlockTimesResponse
	isSet bool
}

func (v NullableNetworkBlockTimesResponse) Get() *NetworkBlockTimesResponse {
	return v.value
}

func (v *NullableNetworkBlockTimesResponse) Set(val *NetworkBlockTimesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkBlockTimesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkBlockTimesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkBlockTimesResponse(val *NetworkBlockTimesResponse) *NullableNetworkBlockTimesResponse {
	return &NullableNetworkBlockTimesResponse{value: val, isSet: true}
}

func (v NullableNetworkBlockTimesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkBlockTimesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


