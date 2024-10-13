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

// checks if the StxLockTransactionEvent1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StxLockTransactionEvent1{}

// StxLockTransactionEvent1 Only present in `smart_contract` and `contract_call` tx types.
type StxLockTransactionEvent1 struct {
	EventIndex int32 `json:"event_index"`
	EventType string `json:"event_type"`
	TxId string `json:"tx_id"`
	StxLockEvent StxLockTransactionEventAllOfStxLockEvent `json:"stx_lock_event"`
}

type _StxLockTransactionEvent1 StxLockTransactionEvent1

// NewStxLockTransactionEvent1 instantiates a new StxLockTransactionEvent1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStxLockTransactionEvent1(eventIndex int32, eventType string, txId string, stxLockEvent StxLockTransactionEventAllOfStxLockEvent) *StxLockTransactionEvent1 {
	this := StxLockTransactionEvent1{}
	this.EventIndex = eventIndex
	this.EventType = eventType
	this.TxId = txId
	this.StxLockEvent = stxLockEvent
	return &this
}

// NewStxLockTransactionEvent1WithDefaults instantiates a new StxLockTransactionEvent1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStxLockTransactionEvent1WithDefaults() *StxLockTransactionEvent1 {
	this := StxLockTransactionEvent1{}
	return &this
}

// GetEventIndex returns the EventIndex field value
func (o *StxLockTransactionEvent1) GetEventIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventIndex
}

// GetEventIndexOk returns a tuple with the EventIndex field value
// and a boolean to check if the value has been set.
func (o *StxLockTransactionEvent1) GetEventIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventIndex, true
}

// SetEventIndex sets field value
func (o *StxLockTransactionEvent1) SetEventIndex(v int32) {
	o.EventIndex = v
}

// GetEventType returns the EventType field value
func (o *StxLockTransactionEvent1) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *StxLockTransactionEvent1) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *StxLockTransactionEvent1) SetEventType(v string) {
	o.EventType = v
}

// GetTxId returns the TxId field value
func (o *StxLockTransactionEvent1) GetTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *StxLockTransactionEvent1) GetTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *StxLockTransactionEvent1) SetTxId(v string) {
	o.TxId = v
}

// GetStxLockEvent returns the StxLockEvent field value
func (o *StxLockTransactionEvent1) GetStxLockEvent() StxLockTransactionEventAllOfStxLockEvent {
	if o == nil {
		var ret StxLockTransactionEventAllOfStxLockEvent
		return ret
	}

	return o.StxLockEvent
}

// GetStxLockEventOk returns a tuple with the StxLockEvent field value
// and a boolean to check if the value has been set.
func (o *StxLockTransactionEvent1) GetStxLockEventOk() (*StxLockTransactionEventAllOfStxLockEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StxLockEvent, true
}

// SetStxLockEvent sets field value
func (o *StxLockTransactionEvent1) SetStxLockEvent(v StxLockTransactionEventAllOfStxLockEvent) {
	o.StxLockEvent = v
}

func (o StxLockTransactionEvent1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StxLockTransactionEvent1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event_index"] = o.EventIndex
	toSerialize["event_type"] = o.EventType
	toSerialize["tx_id"] = o.TxId
	toSerialize["stx_lock_event"] = o.StxLockEvent
	return toSerialize, nil
}

func (o *StxLockTransactionEvent1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"event_index",
		"event_type",
		"tx_id",
		"stx_lock_event",
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

	varStxLockTransactionEvent1 := _StxLockTransactionEvent1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStxLockTransactionEvent1)

	if err != nil {
		return err
	}

	*o = StxLockTransactionEvent1(varStxLockTransactionEvent1)

	return err
}

type NullableStxLockTransactionEvent1 struct {
	value *StxLockTransactionEvent1
	isSet bool
}

func (v NullableStxLockTransactionEvent1) Get() *StxLockTransactionEvent1 {
	return v.value
}

func (v *NullableStxLockTransactionEvent1) Set(val *StxLockTransactionEvent1) {
	v.value = val
	v.isSet = true
}

func (v NullableStxLockTransactionEvent1) IsSet() bool {
	return v.isSet
}

func (v *NullableStxLockTransactionEvent1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStxLockTransactionEvent1(val *StxLockTransactionEvent1) *NullableStxLockTransactionEvent1 {
	return &NullableStxLockTransactionEvent1{value: val, isSet: true}
}

func (v NullableStxLockTransactionEvent1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStxLockTransactionEvent1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

