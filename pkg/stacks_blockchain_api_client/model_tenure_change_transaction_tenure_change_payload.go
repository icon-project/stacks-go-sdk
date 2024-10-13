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

// checks if the TenureChangeTransactionTenureChangePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenureChangeTransactionTenureChangePayload{}

// TenureChangeTransactionTenureChangePayload struct for TenureChangeTransactionTenureChangePayload
type TenureChangeTransactionTenureChangePayload struct {
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
	Cause TenureChangeTransactionTenureChangePayloadCause `json:"cause"`
	// (Hex string) The ECDSA public key hash of the current tenure.
	PubkeyHash string `json:"pubkey_hash"`
}

type _TenureChangeTransactionTenureChangePayload TenureChangeTransactionTenureChangePayload

// NewTenureChangeTransactionTenureChangePayload instantiates a new TenureChangeTransactionTenureChangePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenureChangeTransactionTenureChangePayload(tenureConsensusHash string, prevTenureConsensusHash string, burnViewConsensusHash string, previousTenureEnd string, previousTenureBlocks int32, cause TenureChangeTransactionTenureChangePayloadCause, pubkeyHash string) *TenureChangeTransactionTenureChangePayload {
	this := TenureChangeTransactionTenureChangePayload{}
	this.TenureConsensusHash = tenureConsensusHash
	this.PrevTenureConsensusHash = prevTenureConsensusHash
	this.BurnViewConsensusHash = burnViewConsensusHash
	this.PreviousTenureEnd = previousTenureEnd
	this.PreviousTenureBlocks = previousTenureBlocks
	this.Cause = cause
	this.PubkeyHash = pubkeyHash
	return &this
}

// NewTenureChangeTransactionTenureChangePayloadWithDefaults instantiates a new TenureChangeTransactionTenureChangePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenureChangeTransactionTenureChangePayloadWithDefaults() *TenureChangeTransactionTenureChangePayload {
	this := TenureChangeTransactionTenureChangePayload{}
	return &this
}

// GetTenureConsensusHash returns the TenureConsensusHash field value
func (o *TenureChangeTransactionTenureChangePayload) GetTenureConsensusHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenureConsensusHash
}

// GetTenureConsensusHashOk returns a tuple with the TenureConsensusHash field value
// and a boolean to check if the value has been set.
func (o *TenureChangeTransactionTenureChangePayload) GetTenureConsensusHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenureConsensusHash, true
}

// SetTenureConsensusHash sets field value
func (o *TenureChangeTransactionTenureChangePayload) SetTenureConsensusHash(v string) {
	o.TenureConsensusHash = v
}

// GetPrevTenureConsensusHash returns the PrevTenureConsensusHash field value
func (o *TenureChangeTransactionTenureChangePayload) GetPrevTenureConsensusHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrevTenureConsensusHash
}

// GetPrevTenureConsensusHashOk returns a tuple with the PrevTenureConsensusHash field value
// and a boolean to check if the value has been set.
func (o *TenureChangeTransactionTenureChangePayload) GetPrevTenureConsensusHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrevTenureConsensusHash, true
}

// SetPrevTenureConsensusHash sets field value
func (o *TenureChangeTransactionTenureChangePayload) SetPrevTenureConsensusHash(v string) {
	o.PrevTenureConsensusHash = v
}

// GetBurnViewConsensusHash returns the BurnViewConsensusHash field value
func (o *TenureChangeTransactionTenureChangePayload) GetBurnViewConsensusHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BurnViewConsensusHash
}

// GetBurnViewConsensusHashOk returns a tuple with the BurnViewConsensusHash field value
// and a boolean to check if the value has been set.
func (o *TenureChangeTransactionTenureChangePayload) GetBurnViewConsensusHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnViewConsensusHash, true
}

// SetBurnViewConsensusHash sets field value
func (o *TenureChangeTransactionTenureChangePayload) SetBurnViewConsensusHash(v string) {
	o.BurnViewConsensusHash = v
}

// GetPreviousTenureEnd returns the PreviousTenureEnd field value
func (o *TenureChangeTransactionTenureChangePayload) GetPreviousTenureEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreviousTenureEnd
}

// GetPreviousTenureEndOk returns a tuple with the PreviousTenureEnd field value
// and a boolean to check if the value has been set.
func (o *TenureChangeTransactionTenureChangePayload) GetPreviousTenureEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviousTenureEnd, true
}

// SetPreviousTenureEnd sets field value
func (o *TenureChangeTransactionTenureChangePayload) SetPreviousTenureEnd(v string) {
	o.PreviousTenureEnd = v
}

// GetPreviousTenureBlocks returns the PreviousTenureBlocks field value
func (o *TenureChangeTransactionTenureChangePayload) GetPreviousTenureBlocks() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PreviousTenureBlocks
}

// GetPreviousTenureBlocksOk returns a tuple with the PreviousTenureBlocks field value
// and a boolean to check if the value has been set.
func (o *TenureChangeTransactionTenureChangePayload) GetPreviousTenureBlocksOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviousTenureBlocks, true
}

// SetPreviousTenureBlocks sets field value
func (o *TenureChangeTransactionTenureChangePayload) SetPreviousTenureBlocks(v int32) {
	o.PreviousTenureBlocks = v
}

// GetCause returns the Cause field value
func (o *TenureChangeTransactionTenureChangePayload) GetCause() TenureChangeTransactionTenureChangePayloadCause {
	if o == nil {
		var ret TenureChangeTransactionTenureChangePayloadCause
		return ret
	}

	return o.Cause
}

// GetCauseOk returns a tuple with the Cause field value
// and a boolean to check if the value has been set.
func (o *TenureChangeTransactionTenureChangePayload) GetCauseOk() (*TenureChangeTransactionTenureChangePayloadCause, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cause, true
}

// SetCause sets field value
func (o *TenureChangeTransactionTenureChangePayload) SetCause(v TenureChangeTransactionTenureChangePayloadCause) {
	o.Cause = v
}

// GetPubkeyHash returns the PubkeyHash field value
func (o *TenureChangeTransactionTenureChangePayload) GetPubkeyHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PubkeyHash
}

// GetPubkeyHashOk returns a tuple with the PubkeyHash field value
// and a boolean to check if the value has been set.
func (o *TenureChangeTransactionTenureChangePayload) GetPubkeyHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PubkeyHash, true
}

// SetPubkeyHash sets field value
func (o *TenureChangeTransactionTenureChangePayload) SetPubkeyHash(v string) {
	o.PubkeyHash = v
}

func (o TenureChangeTransactionTenureChangePayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenureChangeTransactionTenureChangePayload) ToMap() (map[string]interface{}, error) {
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

func (o *TenureChangeTransactionTenureChangePayload) UnmarshalJSON(data []byte) (err error) {
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

	varTenureChangeTransactionTenureChangePayload := _TenureChangeTransactionTenureChangePayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTenureChangeTransactionTenureChangePayload)

	if err != nil {
		return err
	}

	*o = TenureChangeTransactionTenureChangePayload(varTenureChangeTransactionTenureChangePayload)

	return err
}

type NullableTenureChangeTransactionTenureChangePayload struct {
	value *TenureChangeTransactionTenureChangePayload
	isSet bool
}

func (v NullableTenureChangeTransactionTenureChangePayload) Get() *TenureChangeTransactionTenureChangePayload {
	return v.value
}

func (v *NullableTenureChangeTransactionTenureChangePayload) Set(val *TenureChangeTransactionTenureChangePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableTenureChangeTransactionTenureChangePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableTenureChangeTransactionTenureChangePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenureChangeTransactionTenureChangePayload(val *TenureChangeTransactionTenureChangePayload) *NullableTenureChangeTransactionTenureChangePayload {
	return &NullableTenureChangeTransactionTenureChangePayload{value: val, isSet: true}
}

func (v NullableTenureChangeTransactionTenureChangePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenureChangeTransactionTenureChangePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

