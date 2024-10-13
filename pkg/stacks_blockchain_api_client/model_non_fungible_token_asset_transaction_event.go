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

// checks if the NonFungibleTokenAssetTransactionEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NonFungibleTokenAssetTransactionEvent{}

// NonFungibleTokenAssetTransactionEvent struct for NonFungibleTokenAssetTransactionEvent
type NonFungibleTokenAssetTransactionEvent struct {
	EventIndex int32 `json:"event_index"`
	EventType string `json:"event_type"`
	TxId string `json:"tx_id"`
	Asset NonFungibleTokenAssetTransactionEventAllOfAsset `json:"asset"`
}

type _NonFungibleTokenAssetTransactionEvent NonFungibleTokenAssetTransactionEvent

// NewNonFungibleTokenAssetTransactionEvent instantiates a new NonFungibleTokenAssetTransactionEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonFungibleTokenAssetTransactionEvent(eventIndex int32, eventType string, txId string, asset NonFungibleTokenAssetTransactionEventAllOfAsset) *NonFungibleTokenAssetTransactionEvent {
	this := NonFungibleTokenAssetTransactionEvent{}
	this.EventIndex = eventIndex
	this.EventType = eventType
	this.TxId = txId
	this.Asset = asset
	return &this
}

// NewNonFungibleTokenAssetTransactionEventWithDefaults instantiates a new NonFungibleTokenAssetTransactionEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonFungibleTokenAssetTransactionEventWithDefaults() *NonFungibleTokenAssetTransactionEvent {
	this := NonFungibleTokenAssetTransactionEvent{}
	return &this
}

// GetEventIndex returns the EventIndex field value
func (o *NonFungibleTokenAssetTransactionEvent) GetEventIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventIndex
}

// GetEventIndexOk returns a tuple with the EventIndex field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenAssetTransactionEvent) GetEventIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventIndex, true
}

// SetEventIndex sets field value
func (o *NonFungibleTokenAssetTransactionEvent) SetEventIndex(v int32) {
	o.EventIndex = v
}

// GetEventType returns the EventType field value
func (o *NonFungibleTokenAssetTransactionEvent) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenAssetTransactionEvent) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *NonFungibleTokenAssetTransactionEvent) SetEventType(v string) {
	o.EventType = v
}

// GetTxId returns the TxId field value
func (o *NonFungibleTokenAssetTransactionEvent) GetTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenAssetTransactionEvent) GetTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *NonFungibleTokenAssetTransactionEvent) SetTxId(v string) {
	o.TxId = v
}

// GetAsset returns the Asset field value
func (o *NonFungibleTokenAssetTransactionEvent) GetAsset() NonFungibleTokenAssetTransactionEventAllOfAsset {
	if o == nil {
		var ret NonFungibleTokenAssetTransactionEventAllOfAsset
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenAssetTransactionEvent) GetAssetOk() (*NonFungibleTokenAssetTransactionEventAllOfAsset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *NonFungibleTokenAssetTransactionEvent) SetAsset(v NonFungibleTokenAssetTransactionEventAllOfAsset) {
	o.Asset = v
}

func (o NonFungibleTokenAssetTransactionEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonFungibleTokenAssetTransactionEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event_index"] = o.EventIndex
	toSerialize["event_type"] = o.EventType
	toSerialize["tx_id"] = o.TxId
	toSerialize["asset"] = o.Asset
	return toSerialize, nil
}

func (o *NonFungibleTokenAssetTransactionEvent) UnmarshalJSON(data []byte) (err error) {
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

	varNonFungibleTokenAssetTransactionEvent := _NonFungibleTokenAssetTransactionEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNonFungibleTokenAssetTransactionEvent)

	if err != nil {
		return err
	}

	*o = NonFungibleTokenAssetTransactionEvent(varNonFungibleTokenAssetTransactionEvent)

	return err
}

type NullableNonFungibleTokenAssetTransactionEvent struct {
	value *NonFungibleTokenAssetTransactionEvent
	isSet bool
}

func (v NullableNonFungibleTokenAssetTransactionEvent) Get() *NonFungibleTokenAssetTransactionEvent {
	return v.value
}

func (v *NullableNonFungibleTokenAssetTransactionEvent) Set(val *NonFungibleTokenAssetTransactionEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableNonFungibleTokenAssetTransactionEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableNonFungibleTokenAssetTransactionEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonFungibleTokenAssetTransactionEvent(val *NonFungibleTokenAssetTransactionEvent) *NullableNonFungibleTokenAssetTransactionEvent {
	return &NullableNonFungibleTokenAssetTransactionEvent{value: val, isSet: true}
}

func (v NullableNonFungibleTokenAssetTransactionEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonFungibleTokenAssetTransactionEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

