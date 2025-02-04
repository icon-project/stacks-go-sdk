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

// checks if the GetPoxCycleSignerStackers200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPoxCycleSignerStackers200Response{}

// GetPoxCycleSignerStackers200Response struct for GetPoxCycleSignerStackers200Response
type GetPoxCycleSignerStackers200Response struct {
	Limit int32 `json:"limit"`
	Offset int32 `json:"offset"`
	Total int32 `json:"total"`
	Results []GetPoxCycleSignerStackers200ResponseResultsInner `json:"results"`
}

type _GetPoxCycleSignerStackers200Response GetPoxCycleSignerStackers200Response

// NewGetPoxCycleSignerStackers200Response instantiates a new GetPoxCycleSignerStackers200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPoxCycleSignerStackers200Response(limit int32, offset int32, total int32, results []GetPoxCycleSignerStackers200ResponseResultsInner) *GetPoxCycleSignerStackers200Response {
	this := GetPoxCycleSignerStackers200Response{}
	this.Limit = limit
	this.Offset = offset
	this.Total = total
	this.Results = results
	return &this
}

// NewGetPoxCycleSignerStackers200ResponseWithDefaults instantiates a new GetPoxCycleSignerStackers200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPoxCycleSignerStackers200ResponseWithDefaults() *GetPoxCycleSignerStackers200Response {
	this := GetPoxCycleSignerStackers200Response{}
	return &this
}

// GetLimit returns the Limit field value
func (o *GetPoxCycleSignerStackers200Response) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *GetPoxCycleSignerStackers200Response) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *GetPoxCycleSignerStackers200Response) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *GetPoxCycleSignerStackers200Response) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *GetPoxCycleSignerStackers200Response) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *GetPoxCycleSignerStackers200Response) SetOffset(v int32) {
	o.Offset = v
}

// GetTotal returns the Total field value
func (o *GetPoxCycleSignerStackers200Response) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *GetPoxCycleSignerStackers200Response) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *GetPoxCycleSignerStackers200Response) SetTotal(v int32) {
	o.Total = v
}

// GetResults returns the Results field value
func (o *GetPoxCycleSignerStackers200Response) GetResults() []GetPoxCycleSignerStackers200ResponseResultsInner {
	if o == nil {
		var ret []GetPoxCycleSignerStackers200ResponseResultsInner
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *GetPoxCycleSignerStackers200Response) GetResultsOk() ([]GetPoxCycleSignerStackers200ResponseResultsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *GetPoxCycleSignerStackers200Response) SetResults(v []GetPoxCycleSignerStackers200ResponseResultsInner) {
	o.Results = v
}

func (o GetPoxCycleSignerStackers200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPoxCycleSignerStackers200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["total"] = o.Total
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *GetPoxCycleSignerStackers200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetPoxCycleSignerStackers200Response := _GetPoxCycleSignerStackers200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetPoxCycleSignerStackers200Response)

	if err != nil {
		return err
	}

	*o = GetPoxCycleSignerStackers200Response(varGetPoxCycleSignerStackers200Response)

	return err
}

type NullableGetPoxCycleSignerStackers200Response struct {
	value *GetPoxCycleSignerStackers200Response
	isSet bool
}

func (v NullableGetPoxCycleSignerStackers200Response) Get() *GetPoxCycleSignerStackers200Response {
	return v.value
}

func (v *NullableGetPoxCycleSignerStackers200Response) Set(val *GetPoxCycleSignerStackers200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPoxCycleSignerStackers200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPoxCycleSignerStackers200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPoxCycleSignerStackers200Response(val *GetPoxCycleSignerStackers200Response) *NullableGetPoxCycleSignerStackers200Response {
	return &NullableGetPoxCycleSignerStackers200Response{value: val, isSet: true}
}

func (v NullableGetPoxCycleSignerStackers200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPoxCycleSignerStackers200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


