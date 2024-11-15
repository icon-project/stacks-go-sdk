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

// checks if the GetPoxCycles200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPoxCycles200Response{}

// GetPoxCycles200Response struct for GetPoxCycles200Response
type GetPoxCycles200Response struct {
	Limit int32 `json:"limit"`
	Offset int32 `json:"offset"`
	Total int32 `json:"total"`
	Results []GetPoxCycles200ResponseResultsInner `json:"results"`
}

type _GetPoxCycles200Response GetPoxCycles200Response

// NewGetPoxCycles200Response instantiates a new GetPoxCycles200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPoxCycles200Response(limit int32, offset int32, total int32, results []GetPoxCycles200ResponseResultsInner) *GetPoxCycles200Response {
	this := GetPoxCycles200Response{}
	this.Limit = limit
	this.Offset = offset
	this.Total = total
	this.Results = results
	return &this
}

// NewGetPoxCycles200ResponseWithDefaults instantiates a new GetPoxCycles200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPoxCycles200ResponseWithDefaults() *GetPoxCycles200Response {
	this := GetPoxCycles200Response{}
	return &this
}

// GetLimit returns the Limit field value
func (o *GetPoxCycles200Response) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *GetPoxCycles200Response) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *GetPoxCycles200Response) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *GetPoxCycles200Response) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *GetPoxCycles200Response) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *GetPoxCycles200Response) SetOffset(v int32) {
	o.Offset = v
}

// GetTotal returns the Total field value
func (o *GetPoxCycles200Response) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *GetPoxCycles200Response) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *GetPoxCycles200Response) SetTotal(v int32) {
	o.Total = v
}

// GetResults returns the Results field value
func (o *GetPoxCycles200Response) GetResults() []GetPoxCycles200ResponseResultsInner {
	if o == nil {
		var ret []GetPoxCycles200ResponseResultsInner
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *GetPoxCycles200Response) GetResultsOk() ([]GetPoxCycles200ResponseResultsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *GetPoxCycles200Response) SetResults(v []GetPoxCycles200ResponseResultsInner) {
	o.Results = v
}

func (o GetPoxCycles200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPoxCycles200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["total"] = o.Total
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *GetPoxCycles200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"limit",
		"offset",
		"total",
		"results",
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

	varGetPoxCycles200Response := _GetPoxCycles200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetPoxCycles200Response)

	if err != nil {
		return err
	}

	*o = GetPoxCycles200Response(varGetPoxCycles200Response)

	return err
}

type NullableGetPoxCycles200Response struct {
	value *GetPoxCycles200Response
	isSet bool
}

func (v NullableGetPoxCycles200Response) Get() *GetPoxCycles200Response {
	return v.value
}

func (v *NullableGetPoxCycles200Response) Set(val *GetPoxCycles200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPoxCycles200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPoxCycles200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPoxCycles200Response(val *GetPoxCycles200Response) *NullableGetPoxCycles200Response {
	return &NullableGetPoxCycles200Response{value: val, isSet: true}
}

func (v NullableGetPoxCycles200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPoxCycles200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


