/*
Stacks Blockchain API

Welcome to the API reference overview for the [Stacks Blockchain API](https://docs.hiro.so/stacks-blockchain-api).        [Download Postman collection](https://hirosystems.github.io/stacks-blockchain-api/collection.json)

API version: v7.14.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stacks_blockchain_api_client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GetAverageBlockTimes200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAverageBlockTimes200Response{}

// GetAverageBlockTimes200Response struct for GetAverageBlockTimes200Response
type GetAverageBlockTimes200Response struct {
	// Average block times over the last hour (in seconds)
	Last1h float32 `json:"last_1h"`
	// Average block times over the last 24 hours (in seconds)
	Last24h float32 `json:"last_24h"`
	// Average block times over the last 7 days (in seconds)
	Last7d float32 `json:"last_7d"`
	// Average block times over the last 30 days (in seconds)
	Last30d float32 `json:"last_30d"`
}

type _GetAverageBlockTimes200Response GetAverageBlockTimes200Response

// NewGetAverageBlockTimes200Response instantiates a new GetAverageBlockTimes200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAverageBlockTimes200Response(last1h float32, last24h float32, last7d float32, last30d float32) *GetAverageBlockTimes200Response {
	this := GetAverageBlockTimes200Response{}
	this.Last1h = last1h
	this.Last24h = last24h
	this.Last7d = last7d
	this.Last30d = last30d
	return &this
}

// NewGetAverageBlockTimes200ResponseWithDefaults instantiates a new GetAverageBlockTimes200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAverageBlockTimes200ResponseWithDefaults() *GetAverageBlockTimes200Response {
	this := GetAverageBlockTimes200Response{}
	return &this
}

// GetLast1h returns the Last1h field value
func (o *GetAverageBlockTimes200Response) GetLast1h() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Last1h
}

// GetLast1hOk returns a tuple with the Last1h field value
// and a boolean to check if the value has been set.
func (o *GetAverageBlockTimes200Response) GetLast1hOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Last1h, true
}

// SetLast1h sets field value
func (o *GetAverageBlockTimes200Response) SetLast1h(v float32) {
	o.Last1h = v
}

// GetLast24h returns the Last24h field value
func (o *GetAverageBlockTimes200Response) GetLast24h() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Last24h
}

// GetLast24hOk returns a tuple with the Last24h field value
// and a boolean to check if the value has been set.
func (o *GetAverageBlockTimes200Response) GetLast24hOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Last24h, true
}

// SetLast24h sets field value
func (o *GetAverageBlockTimes200Response) SetLast24h(v float32) {
	o.Last24h = v
}

// GetLast7d returns the Last7d field value
func (o *GetAverageBlockTimes200Response) GetLast7d() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Last7d
}

// GetLast7dOk returns a tuple with the Last7d field value
// and a boolean to check if the value has been set.
func (o *GetAverageBlockTimes200Response) GetLast7dOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Last7d, true
}

// SetLast7d sets field value
func (o *GetAverageBlockTimes200Response) SetLast7d(v float32) {
	o.Last7d = v
}

// GetLast30d returns the Last30d field value
func (o *GetAverageBlockTimes200Response) GetLast30d() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Last30d
}

// GetLast30dOk returns a tuple with the Last30d field value
// and a boolean to check if the value has been set.
func (o *GetAverageBlockTimes200Response) GetLast30dOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Last30d, true
}

// SetLast30d sets field value
func (o *GetAverageBlockTimes200Response) SetLast30d(v float32) {
	o.Last30d = v
}

func (o GetAverageBlockTimes200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAverageBlockTimes200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["last_1h"] = o.Last1h
	toSerialize["last_24h"] = o.Last24h
	toSerialize["last_7d"] = o.Last7d
	toSerialize["last_30d"] = o.Last30d
	return toSerialize, nil
}

func (o *GetAverageBlockTimes200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"last_1h",
		"last_24h",
		"last_7d",
		"last_30d",
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

	varGetAverageBlockTimes200Response := _GetAverageBlockTimes200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetAverageBlockTimes200Response)

	if err != nil {
		return err
	}

	*o = GetAverageBlockTimes200Response(varGetAverageBlockTimes200Response)

	return err
}

type NullableGetAverageBlockTimes200Response struct {
	value *GetAverageBlockTimes200Response
	isSet bool
}

func (v NullableGetAverageBlockTimes200Response) Get() *GetAverageBlockTimes200Response {
	return v.value
}

func (v *NullableGetAverageBlockTimes200Response) Set(val *GetAverageBlockTimes200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAverageBlockTimes200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAverageBlockTimes200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAverageBlockTimes200Response(val *GetAverageBlockTimes200Response) *NullableGetAverageBlockTimes200Response {
	return &NullableGetAverageBlockTimes200Response{value: val, isSet: true}
}

func (v NullableGetAverageBlockTimes200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAverageBlockTimes200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

