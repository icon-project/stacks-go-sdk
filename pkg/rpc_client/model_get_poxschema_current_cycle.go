/*
Stacks 2.0+ RPC API

This is the documentation for the `stacks-node` RPC interface. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rpc_client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GetPoxschemaCurrentCycle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPoxschemaCurrentCycle{}

// GetPoxschemaCurrentCycle struct for GetPoxschemaCurrentCycle
type GetPoxschemaCurrentCycle struct {
	// The reward cycle number
	Id int32 `json:"id"`
	// The threshold amount for obtaining a slot in this reward cycle.
	MinThresholdUstx int32 `json:"min_threshold_ustx"`
	// The total amount of stacked microstacks in this reward cycle.
	StackedUstx int32 `json:"stacked_ustx"`
	// Whether or not PoX is active during this reward cycle.
	IsPoxActive bool `json:"is_pox_active"`
}

type _GetPoxschemaCurrentCycle GetPoxschemaCurrentCycle

// NewGetPoxschemaCurrentCycle instantiates a new GetPoxschemaCurrentCycle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPoxschemaCurrentCycle(id int32, minThresholdUstx int32, stackedUstx int32, isPoxActive bool) *GetPoxschemaCurrentCycle {
	this := GetPoxschemaCurrentCycle{}
	this.Id = id
	this.MinThresholdUstx = minThresholdUstx
	this.StackedUstx = stackedUstx
	this.IsPoxActive = isPoxActive
	return &this
}

// NewGetPoxschemaCurrentCycleWithDefaults instantiates a new GetPoxschemaCurrentCycle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPoxschemaCurrentCycleWithDefaults() *GetPoxschemaCurrentCycle {
	this := GetPoxschemaCurrentCycle{}
	return &this
}

// GetId returns the Id field value
func (o *GetPoxschemaCurrentCycle) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetPoxschemaCurrentCycle) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetPoxschemaCurrentCycle) SetId(v int32) {
	o.Id = v
}

// GetMinThresholdUstx returns the MinThresholdUstx field value
func (o *GetPoxschemaCurrentCycle) GetMinThresholdUstx() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinThresholdUstx
}

// GetMinThresholdUstxOk returns a tuple with the MinThresholdUstx field value
// and a boolean to check if the value has been set.
func (o *GetPoxschemaCurrentCycle) GetMinThresholdUstxOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinThresholdUstx, true
}

// SetMinThresholdUstx sets field value
func (o *GetPoxschemaCurrentCycle) SetMinThresholdUstx(v int32) {
	o.MinThresholdUstx = v
}

// GetStackedUstx returns the StackedUstx field value
func (o *GetPoxschemaCurrentCycle) GetStackedUstx() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StackedUstx
}

// GetStackedUstxOk returns a tuple with the StackedUstx field value
// and a boolean to check if the value has been set.
func (o *GetPoxschemaCurrentCycle) GetStackedUstxOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StackedUstx, true
}

// SetStackedUstx sets field value
func (o *GetPoxschemaCurrentCycle) SetStackedUstx(v int32) {
	o.StackedUstx = v
}

// GetIsPoxActive returns the IsPoxActive field value
func (o *GetPoxschemaCurrentCycle) GetIsPoxActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPoxActive
}

// GetIsPoxActiveOk returns a tuple with the IsPoxActive field value
// and a boolean to check if the value has been set.
func (o *GetPoxschemaCurrentCycle) GetIsPoxActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPoxActive, true
}

// SetIsPoxActive sets field value
func (o *GetPoxschemaCurrentCycle) SetIsPoxActive(v bool) {
	o.IsPoxActive = v
}

func (o GetPoxschemaCurrentCycle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPoxschemaCurrentCycle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["min_threshold_ustx"] = o.MinThresholdUstx
	toSerialize["stacked_ustx"] = o.StackedUstx
	toSerialize["is_pox_active"] = o.IsPoxActive
	return toSerialize, nil
}

func (o *GetPoxschemaCurrentCycle) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"min_threshold_ustx",
		"stacked_ustx",
		"is_pox_active",
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

	varGetPoxschemaCurrentCycle := _GetPoxschemaCurrentCycle{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetPoxschemaCurrentCycle)

	if err != nil {
		return err
	}

	*o = GetPoxschemaCurrentCycle(varGetPoxschemaCurrentCycle)

	return err
}

type NullableGetPoxschemaCurrentCycle struct {
	value *GetPoxschemaCurrentCycle
	isSet bool
}

func (v NullableGetPoxschemaCurrentCycle) Get() *GetPoxschemaCurrentCycle {
	return v.value
}

func (v *NullableGetPoxschemaCurrentCycle) Set(val *GetPoxschemaCurrentCycle) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPoxschemaCurrentCycle) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPoxschemaCurrentCycle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPoxschemaCurrentCycle(val *GetPoxschemaCurrentCycle) *NullableGetPoxschemaCurrentCycle {
	return &NullableGetPoxschemaCurrentCycle{value: val, isSet: true}
}

func (v NullableGetPoxschemaCurrentCycle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPoxschemaCurrentCycle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

