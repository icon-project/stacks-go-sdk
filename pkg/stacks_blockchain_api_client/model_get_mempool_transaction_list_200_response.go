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

// checks if the GetMempoolTransactionList200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMempoolTransactionList200Response{}

// GetMempoolTransactionList200Response List of mempool transactions
type GetMempoolTransactionList200Response struct {
	Limit int32 `json:"limit"`
	Offset int32 `json:"offset"`
	Total int32 `json:"total"`
	Results []GetMempoolTransactionList200ResponseResultsInner `json:"results"`
}

type _GetMempoolTransactionList200Response GetMempoolTransactionList200Response

// NewGetMempoolTransactionList200Response instantiates a new GetMempoolTransactionList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMempoolTransactionList200Response(limit int32, offset int32, total int32, results []GetMempoolTransactionList200ResponseResultsInner) *GetMempoolTransactionList200Response {
	this := GetMempoolTransactionList200Response{}
	this.Limit = limit
	this.Offset = offset
	this.Total = total
	this.Results = results
	return &this
}

// NewGetMempoolTransactionList200ResponseWithDefaults instantiates a new GetMempoolTransactionList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMempoolTransactionList200ResponseWithDefaults() *GetMempoolTransactionList200Response {
	this := GetMempoolTransactionList200Response{}
	return &this
}

// GetLimit returns the Limit field value
func (o *GetMempoolTransactionList200Response) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *GetMempoolTransactionList200Response) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *GetMempoolTransactionList200Response) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *GetMempoolTransactionList200Response) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *GetMempoolTransactionList200Response) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *GetMempoolTransactionList200Response) SetOffset(v int32) {
	o.Offset = v
}

// GetTotal returns the Total field value
func (o *GetMempoolTransactionList200Response) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *GetMempoolTransactionList200Response) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *GetMempoolTransactionList200Response) SetTotal(v int32) {
	o.Total = v
}

// GetResults returns the Results field value
func (o *GetMempoolTransactionList200Response) GetResults() []GetMempoolTransactionList200ResponseResultsInner {
	if o == nil {
		var ret []GetMempoolTransactionList200ResponseResultsInner
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *GetMempoolTransactionList200Response) GetResultsOk() ([]GetMempoolTransactionList200ResponseResultsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *GetMempoolTransactionList200Response) SetResults(v []GetMempoolTransactionList200ResponseResultsInner) {
	o.Results = v
}

func (o GetMempoolTransactionList200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMempoolTransactionList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["total"] = o.Total
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *GetMempoolTransactionList200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetMempoolTransactionList200Response := _GetMempoolTransactionList200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetMempoolTransactionList200Response)

	if err != nil {
		return err
	}

	*o = GetMempoolTransactionList200Response(varGetMempoolTransactionList200Response)

	return err
}

type NullableGetMempoolTransactionList200Response struct {
	value *GetMempoolTransactionList200Response
	isSet bool
}

func (v NullableGetMempoolTransactionList200Response) Get() *GetMempoolTransactionList200Response {
	return v.value
}

func (v *NullableGetMempoolTransactionList200Response) Set(val *GetMempoolTransactionList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMempoolTransactionList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMempoolTransactionList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMempoolTransactionList200Response(val *GetMempoolTransactionList200Response) *NullableGetMempoolTransactionList200Response {
	return &NullableGetMempoolTransactionList200Response{value: val, isSet: true}
}

func (v NullableGetMempoolTransactionList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMempoolTransactionList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


