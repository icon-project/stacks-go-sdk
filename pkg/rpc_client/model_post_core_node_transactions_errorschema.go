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

// checks if the PostCoreNodeTransactionsErrorschema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCoreNodeTransactionsErrorschema{}

// PostCoreNodeTransactionsErrorschema GET request that returns transactions
type PostCoreNodeTransactionsErrorschema struct {
	// The error
	Error string `json:"error"`
	// The reason for the error
	Reason string `json:"reason"`
	// More details about the reason
	ReasonData map[string]interface{} `json:"reason_data"`
	// The relevant transaction id
	Txid string `json:"txid"`
}

type _PostCoreNodeTransactionsErrorschema PostCoreNodeTransactionsErrorschema

// NewPostCoreNodeTransactionsErrorschema instantiates a new PostCoreNodeTransactionsErrorschema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCoreNodeTransactionsErrorschema(error_ string, reason string, reasonData map[string]interface{}, txid string) *PostCoreNodeTransactionsErrorschema {
	this := PostCoreNodeTransactionsErrorschema{}
	this.Error = error_
	this.Reason = reason
	this.ReasonData = reasonData
	this.Txid = txid
	return &this
}

// NewPostCoreNodeTransactionsErrorschemaWithDefaults instantiates a new PostCoreNodeTransactionsErrorschema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCoreNodeTransactionsErrorschemaWithDefaults() *PostCoreNodeTransactionsErrorschema {
	this := PostCoreNodeTransactionsErrorschema{}
	return &this
}

// GetError returns the Error field value
func (o *PostCoreNodeTransactionsErrorschema) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *PostCoreNodeTransactionsErrorschema) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *PostCoreNodeTransactionsErrorschema) SetError(v string) {
	o.Error = v
}

// GetReason returns the Reason field value
func (o *PostCoreNodeTransactionsErrorschema) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *PostCoreNodeTransactionsErrorschema) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *PostCoreNodeTransactionsErrorschema) SetReason(v string) {
	o.Reason = v
}

// GetReasonData returns the ReasonData field value
func (o *PostCoreNodeTransactionsErrorschema) GetReasonData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ReasonData
}

// GetReasonDataOk returns a tuple with the ReasonData field value
// and a boolean to check if the value has been set.
func (o *PostCoreNodeTransactionsErrorschema) GetReasonDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.ReasonData, true
}

// SetReasonData sets field value
func (o *PostCoreNodeTransactionsErrorschema) SetReasonData(v map[string]interface{}) {
	o.ReasonData = v
}

// GetTxid returns the Txid field value
func (o *PostCoreNodeTransactionsErrorschema) GetTxid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Txid
}

// GetTxidOk returns a tuple with the Txid field value
// and a boolean to check if the value has been set.
func (o *PostCoreNodeTransactionsErrorschema) GetTxidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Txid, true
}

// SetTxid sets field value
func (o *PostCoreNodeTransactionsErrorschema) SetTxid(v string) {
	o.Txid = v
}

func (o PostCoreNodeTransactionsErrorschema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCoreNodeTransactionsErrorschema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	toSerialize["reason"] = o.Reason
	toSerialize["reason_data"] = o.ReasonData
	toSerialize["txid"] = o.Txid
	return toSerialize, nil
}

func (o *PostCoreNodeTransactionsErrorschema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"error",
		"reason",
		"reason_data",
		"txid",
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

	varPostCoreNodeTransactionsErrorschema := _PostCoreNodeTransactionsErrorschema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostCoreNodeTransactionsErrorschema)

	if err != nil {
		return err
	}

	*o = PostCoreNodeTransactionsErrorschema(varPostCoreNodeTransactionsErrorschema)

	return err
}

type NullablePostCoreNodeTransactionsErrorschema struct {
	value *PostCoreNodeTransactionsErrorschema
	isSet bool
}

func (v NullablePostCoreNodeTransactionsErrorschema) Get() *PostCoreNodeTransactionsErrorschema {
	return v.value
}

func (v *NullablePostCoreNodeTransactionsErrorschema) Set(val *PostCoreNodeTransactionsErrorschema) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCoreNodeTransactionsErrorschema) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCoreNodeTransactionsErrorschema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCoreNodeTransactionsErrorschema(val *PostCoreNodeTransactionsErrorschema) *NullablePostCoreNodeTransactionsErrorschema {
	return &NullablePostCoreNodeTransactionsErrorschema{value: val, isSet: true}
}

func (v NullablePostCoreNodeTransactionsErrorschema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCoreNodeTransactionsErrorschema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


