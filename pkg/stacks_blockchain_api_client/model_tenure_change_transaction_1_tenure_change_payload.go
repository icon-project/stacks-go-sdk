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

// checks if the TenureChangeTransaction1TenureChangePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenureChangeTransaction1TenureChangePayload{}

// TenureChangeTransaction1TenureChangePayload struct for TenureChangeTransaction1TenureChangePayload
type TenureChangeTransaction1TenureChangePayload struct {
	// Consensus hash of this tenure. Corresponds to the sortition in which the miner of this block was chosen.
	TenureConsensusHash string `json:"tenure_consensus_hash"`
	// Consensus hash of the previous tenure. Corresponds to the sortition of the previous winning block-commit.
	PrevTenureConsensusHash string `json:"prev_tenure_consensus_hash"`
	// Current consensus hash on the underlying burnchain. Corresponds to the last-seen sortition.
	BurnViewConsensusHash string `json:"burn_view_consensus_hash"`
	// (Hex string) Stacks Block hash
	PreviousTenureEnd string `json:"previous_tenure_end"`
	// The number of blocks produced in the previous tenure.
	PreviousTenureBlocks int32 `json:"previous_tenure_blocks"`
	Cause TenureChangeTransaction1TenureChangePayloadCause `json:"cause"`
	// (Hex string) The ECDSA public key hash of the current tenure.
	PubkeyHash string `json:"pubkey_hash"`
}

type _TenureChangeTransaction1TenureChangePayload TenureChangeTransaction1TenureChangePayload

// NewTenureChangeTransaction1TenureChangePayload instantiates a new TenureChangeTransaction1TenureChangePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenureChangeTransaction1TenureChangePayload(tenureConsensusHash string, prevTenureConsensusHash string, burnViewConsensusHash string, previousTenureEnd string, previousTenureBlocks int32, cause TenureChangeTransaction1TenureChangePayloadCause, pubkeyHash string) *TenureChangeTransaction1TenureChangePayload {
	this := TenureChangeTransaction1TenureChangePayload{}
	this.TenureConsensusHash = tenureConsensusHash
	this.PrevTenureConsensusHash = prevTenureConsensusHash
	this.BurnViewConsensusHash = burnViewConsensusHash
	this.PreviousTenureEnd = previousTenureEnd
	this.PreviousTenureBlocks = previousTenureBlocks
	this.Cause = cause
	this.PubkeyHash = pubkeyHash
	return &this
}

// NewTenureChangeTransaction1TenureChangePayloadWithDefaults instantiates a new TenureChangeTransaction1TenureChangePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenureChangeTransaction1TenureChangePayloadWithDefaults() *TenureChangeTransaction1TenureChangePayload {
	this := TenureChangeTransaction1TenureChangePayload{}
	return &this
}

// GetTenureConsensusHash returns the TenureConsensusHash field value
func (o *TenureChangeTransaction1TenureChangePayload) GetTenureConsensusHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenureConsensusHash
}

// GetTenureConsensusHashOk returns a tuple with the TenureConsensusHash field value
// and a boolean to check if the value has been set.
func (o *TenureChangeTransaction1TenureChangePayload) GetTenureConsensusHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenureConsensusHash, true
}

// SetTenureConsensusHash sets field value
func (o *TenureChangeTransaction1TenureChangePayload) SetTenureConsensusHash(v string) {
	o.TenureConsensusHash = v
}

// GetPrevTenureConsensusHash returns the PrevTenureConsensusHash field value
func (o *TenureChangeTransaction1TenureChangePayload) GetPrevTenureConsensusHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrevTenureConsensusHash
}

// GetPrevTenureConsensusHashOk returns a tuple with the PrevTenureConsensusHash field value
// and a boolean to check if the value has been set.
func (o *TenureChangeTransaction1TenureChangePayload) GetPrevTenureConsensusHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrevTenureConsensusHash, true
}

// SetPrevTenureConsensusHash sets field value
func (o *TenureChangeTransaction1TenureChangePayload) SetPrevTenureConsensusHash(v string) {
	o.PrevTenureConsensusHash = v
}

// GetBurnViewConsensusHash returns the BurnViewConsensusHash field value
func (o *TenureChangeTransaction1TenureChangePayload) GetBurnViewConsensusHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BurnViewConsensusHash
}

// GetBurnViewConsensusHashOk returns a tuple with the BurnViewConsensusHash field value
// and a boolean to check if the value has been set.
func (o *TenureChangeTransaction1TenureChangePayload) GetBurnViewConsensusHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnViewConsensusHash, true
}

// SetBurnViewConsensusHash sets field value
func (o *TenureChangeTransaction1TenureChangePayload) SetBurnViewConsensusHash(v string) {
	o.BurnViewConsensusHash = v
}

// GetPreviousTenureEnd returns the PreviousTenureEnd field value
func (o *TenureChangeTransaction1TenureChangePayload) GetPreviousTenureEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreviousTenureEnd
}

// GetPreviousTenureEndOk returns a tuple with the PreviousTenureEnd field value
// and a boolean to check if the value has been set.
func (o *TenureChangeTransaction1TenureChangePayload) GetPreviousTenureEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviousTenureEnd, true
}

// SetPreviousTenureEnd sets field value
func (o *TenureChangeTransaction1TenureChangePayload) SetPreviousTenureEnd(v string) {
	o.PreviousTenureEnd = v
}

// GetPreviousTenureBlocks returns the PreviousTenureBlocks field value
func (o *TenureChangeTransaction1TenureChangePayload) GetPreviousTenureBlocks() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PreviousTenureBlocks
}

// GetPreviousTenureBlocksOk returns a tuple with the PreviousTenureBlocks field value
// and a boolean to check if the value has been set.
func (o *TenureChangeTransaction1TenureChangePayload) GetPreviousTenureBlocksOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviousTenureBlocks, true
}

// SetPreviousTenureBlocks sets field value
func (o *TenureChangeTransaction1TenureChangePayload) SetPreviousTenureBlocks(v int32) {
	o.PreviousTenureBlocks = v
}

// GetCause returns the Cause field value
func (o *TenureChangeTransaction1TenureChangePayload) GetCause() TenureChangeTransaction1TenureChangePayloadCause {
	if o == nil {
		var ret TenureChangeTransaction1TenureChangePayloadCause
		return ret
	}

	return o.Cause
}

// GetCauseOk returns a tuple with the Cause field value
// and a boolean to check if the value has been set.
func (o *TenureChangeTransaction1TenureChangePayload) GetCauseOk() (*TenureChangeTransaction1TenureChangePayloadCause, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cause, true
}

// SetCause sets field value
func (o *TenureChangeTransaction1TenureChangePayload) SetCause(v TenureChangeTransaction1TenureChangePayloadCause) {
	o.Cause = v
}

// GetPubkeyHash returns the PubkeyHash field value
func (o *TenureChangeTransaction1TenureChangePayload) GetPubkeyHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PubkeyHash
}

// GetPubkeyHashOk returns a tuple with the PubkeyHash field value
// and a boolean to check if the value has been set.
func (o *TenureChangeTransaction1TenureChangePayload) GetPubkeyHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PubkeyHash, true
}

// SetPubkeyHash sets field value
func (o *TenureChangeTransaction1TenureChangePayload) SetPubkeyHash(v string) {
	o.PubkeyHash = v
}

func (o TenureChangeTransaction1TenureChangePayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenureChangeTransaction1TenureChangePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenure_consensus_hash"] = o.TenureConsensusHash
	toSerialize["prev_tenure_consensus_hash"] = o.PrevTenureConsensusHash
	toSerialize["burn_view_consensus_hash"] = o.BurnViewConsensusHash
	toSerialize["previous_tenure_end"] = o.PreviousTenureEnd
	toSerialize["previous_tenure_blocks"] = o.PreviousTenureBlocks
	toSerialize["cause"] = o.Cause
	toSerialize["pubkey_hash"] = o.PubkeyHash
	return toSerialize, nil
}

func (o *TenureChangeTransaction1TenureChangePayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenure_consensus_hash",
		"prev_tenure_consensus_hash",
		"burn_view_consensus_hash",
		"previous_tenure_end",
		"previous_tenure_blocks",
		"cause",
		"pubkey_hash",
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

	varTenureChangeTransaction1TenureChangePayload := _TenureChangeTransaction1TenureChangePayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTenureChangeTransaction1TenureChangePayload)

	if err != nil {
		return err
	}

	*o = TenureChangeTransaction1TenureChangePayload(varTenureChangeTransaction1TenureChangePayload)

	return err
}

type NullableTenureChangeTransaction1TenureChangePayload struct {
	value *TenureChangeTransaction1TenureChangePayload
	isSet bool
}

func (v NullableTenureChangeTransaction1TenureChangePayload) Get() *TenureChangeTransaction1TenureChangePayload {
	return v.value
}

func (v *NullableTenureChangeTransaction1TenureChangePayload) Set(val *TenureChangeTransaction1TenureChangePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableTenureChangeTransaction1TenureChangePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableTenureChangeTransaction1TenureChangePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenureChangeTransaction1TenureChangePayload(val *TenureChangeTransaction1TenureChangePayload) *NullableTenureChangeTransaction1TenureChangePayload {
	return &NullableTenureChangeTransaction1TenureChangePayload{value: val, isSet: true}
}

func (v NullableTenureChangeTransaction1TenureChangePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenureChangeTransaction1TenureChangePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

