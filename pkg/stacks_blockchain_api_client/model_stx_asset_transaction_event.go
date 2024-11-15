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

// checks if the StxAssetTransactionEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StxAssetTransactionEvent{}

// StxAssetTransactionEvent Only present in `smart_contract` and `contract_call` tx types.
type StxAssetTransactionEvent struct {
	EventIndex int32 `json:"event_index"`
	EventType string `json:"event_type"`
	TxId string `json:"tx_id"`
	Asset StxAssetTransactionEventAllOfAsset `json:"asset"`
}

type _StxAssetTransactionEvent StxAssetTransactionEvent

// NewStxAssetTransactionEvent instantiates a new StxAssetTransactionEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStxAssetTransactionEvent(eventIndex int32, eventType string, txId string, asset StxAssetTransactionEventAllOfAsset) *StxAssetTransactionEvent {
	this := StxAssetTransactionEvent{}
	this.EventIndex = eventIndex
	this.EventType = eventType
	this.TxId = txId
	this.Asset = asset
	return &this
}

// NewStxAssetTransactionEventWithDefaults instantiates a new StxAssetTransactionEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStxAssetTransactionEventWithDefaults() *StxAssetTransactionEvent {
	this := StxAssetTransactionEvent{}
	return &this
}

// GetEventIndex returns the EventIndex field value
func (o *StxAssetTransactionEvent) GetEventIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventIndex
}

// GetEventIndexOk returns a tuple with the EventIndex field value
// and a boolean to check if the value has been set.
func (o *StxAssetTransactionEvent) GetEventIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventIndex, true
}

// SetEventIndex sets field value
func (o *StxAssetTransactionEvent) SetEventIndex(v int32) {
	o.EventIndex = v
}

// GetEventType returns the EventType field value
func (o *StxAssetTransactionEvent) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *StxAssetTransactionEvent) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *StxAssetTransactionEvent) SetEventType(v string) {
	o.EventType = v
}

// GetTxId returns the TxId field value
func (o *StxAssetTransactionEvent) GetTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *StxAssetTransactionEvent) GetTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *StxAssetTransactionEvent) SetTxId(v string) {
	o.TxId = v
}

// GetAsset returns the Asset field value
func (o *StxAssetTransactionEvent) GetAsset() StxAssetTransactionEventAllOfAsset {
	if o == nil {
		var ret StxAssetTransactionEventAllOfAsset
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *StxAssetTransactionEvent) GetAssetOk() (*StxAssetTransactionEventAllOfAsset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *StxAssetTransactionEvent) SetAsset(v StxAssetTransactionEventAllOfAsset) {
	o.Asset = v
}

func (o StxAssetTransactionEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StxAssetTransactionEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event_index"] = o.EventIndex
	toSerialize["event_type"] = o.EventType
	toSerialize["tx_id"] = o.TxId
	toSerialize["asset"] = o.Asset
	return toSerialize, nil
}

func (o *StxAssetTransactionEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"event_index",
		"event_type",
		"tx_id",
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

	varStxAssetTransactionEvent := _StxAssetTransactionEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStxAssetTransactionEvent)

	if err != nil {
		return err
	}

	*o = StxAssetTransactionEvent(varStxAssetTransactionEvent)

	return err
}

type NullableStxAssetTransactionEvent struct {
	value *StxAssetTransactionEvent
	isSet bool
}

func (v NullableStxAssetTransactionEvent) Get() *StxAssetTransactionEvent {
	return v.value
}

func (v *NullableStxAssetTransactionEvent) Set(val *StxAssetTransactionEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableStxAssetTransactionEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableStxAssetTransactionEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStxAssetTransactionEvent(val *StxAssetTransactionEvent) *NullableStxAssetTransactionEvent {
	return &NullableStxAssetTransactionEvent{value: val, isSet: true}
}

func (v NullableStxAssetTransactionEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStxAssetTransactionEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


