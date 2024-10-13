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

// checks if the MempoolTransactionStatsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MempoolTransactionStatsResponse{}

// MempoolTransactionStatsResponse GET request that returns stats on mempool transactions
type MempoolTransactionStatsResponse struct {
	// Number of tranasction in the mempool, broken down by transaction type.
	TxTypeCounts map[string]int32 `json:"tx_type_counts"`
	// The simple mean (average) transaction fee, broken down by transaction type. Note that this does not factor in actual execution costs. The average fee is not a reliable metric for calculating a fee for a new transaction.
	TxSimpleFeeAverages map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue `json:"tx_simple_fee_averages"`
	// The average time (in blocks) that transactions have lived in the mempool. The start block height is simply the current chain-tip of when the attached Stacks node receives the transaction. This timing can be different across Stacks nodes / API instances due to propagation timing differences in the p2p network.
	TxAges map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue `json:"tx_ages"`
	// The average byte size of transactions in the mempool, broken down by transaction type.
	TxByteSizes map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue `json:"tx_byte_sizes"`
}

type _MempoolTransactionStatsResponse MempoolTransactionStatsResponse

// NewMempoolTransactionStatsResponse instantiates a new MempoolTransactionStatsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMempoolTransactionStatsResponse(txTypeCounts map[string]int32, txSimpleFeeAverages map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue, txAges map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue, txByteSizes map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue) *MempoolTransactionStatsResponse {
	this := MempoolTransactionStatsResponse{}
	this.TxTypeCounts = txTypeCounts
	this.TxSimpleFeeAverages = txSimpleFeeAverages
	this.TxAges = txAges
	this.TxByteSizes = txByteSizes
	return &this
}

// NewMempoolTransactionStatsResponseWithDefaults instantiates a new MempoolTransactionStatsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMempoolTransactionStatsResponseWithDefaults() *MempoolTransactionStatsResponse {
	this := MempoolTransactionStatsResponse{}
	return &this
}

// GetTxTypeCounts returns the TxTypeCounts field value
func (o *MempoolTransactionStatsResponse) GetTxTypeCounts() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.TxTypeCounts
}

// GetTxTypeCountsOk returns a tuple with the TxTypeCounts field value
// and a boolean to check if the value has been set.
func (o *MempoolTransactionStatsResponse) GetTxTypeCountsOk() (*map[string]int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxTypeCounts, true
}

// SetTxTypeCounts sets field value
func (o *MempoolTransactionStatsResponse) SetTxTypeCounts(v map[string]int32) {
	o.TxTypeCounts = v
}

// GetTxSimpleFeeAverages returns the TxSimpleFeeAverages field value
func (o *MempoolTransactionStatsResponse) GetTxSimpleFeeAverages() map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue {
	if o == nil {
		var ret map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue
		return ret
	}

	return o.TxSimpleFeeAverages
}

// GetTxSimpleFeeAveragesOk returns a tuple with the TxSimpleFeeAverages field value
// and a boolean to check if the value has been set.
func (o *MempoolTransactionStatsResponse) GetTxSimpleFeeAveragesOk() (*map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxSimpleFeeAverages, true
}

// SetTxSimpleFeeAverages sets field value
func (o *MempoolTransactionStatsResponse) SetTxSimpleFeeAverages(v map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue) {
	o.TxSimpleFeeAverages = v
}

// GetTxAges returns the TxAges field value
func (o *MempoolTransactionStatsResponse) GetTxAges() map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue {
	if o == nil {
		var ret map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue
		return ret
	}

	return o.TxAges
}

// GetTxAgesOk returns a tuple with the TxAges field value
// and a boolean to check if the value has been set.
func (o *MempoolTransactionStatsResponse) GetTxAgesOk() (*map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxAges, true
}

// SetTxAges sets field value
func (o *MempoolTransactionStatsResponse) SetTxAges(v map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue) {
	o.TxAges = v
}

// GetTxByteSizes returns the TxByteSizes field value
func (o *MempoolTransactionStatsResponse) GetTxByteSizes() map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue {
	if o == nil {
		var ret map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue
		return ret
	}

	return o.TxByteSizes
}

// GetTxByteSizesOk returns a tuple with the TxByteSizes field value
// and a boolean to check if the value has been set.
func (o *MempoolTransactionStatsResponse) GetTxByteSizesOk() (*map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxByteSizes, true
}

// SetTxByteSizes sets field value
func (o *MempoolTransactionStatsResponse) SetTxByteSizes(v map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue) {
	o.TxByteSizes = v
}

func (o MempoolTransactionStatsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MempoolTransactionStatsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tx_type_counts"] = o.TxTypeCounts
	toSerialize["tx_simple_fee_averages"] = o.TxSimpleFeeAverages
	toSerialize["tx_ages"] = o.TxAges
	toSerialize["tx_byte_sizes"] = o.TxByteSizes
	return toSerialize, nil
}

func (o *MempoolTransactionStatsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tx_type_counts",
		"tx_simple_fee_averages",
		"tx_ages",
		"tx_byte_sizes",
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

	varMempoolTransactionStatsResponse := _MempoolTransactionStatsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMempoolTransactionStatsResponse)

	if err != nil {
		return err
	}

	*o = MempoolTransactionStatsResponse(varMempoolTransactionStatsResponse)

	return err
}

type NullableMempoolTransactionStatsResponse struct {
	value *MempoolTransactionStatsResponse
	isSet bool
}

func (v NullableMempoolTransactionStatsResponse) Get() *MempoolTransactionStatsResponse {
	return v.value
}

func (v *NullableMempoolTransactionStatsResponse) Set(val *MempoolTransactionStatsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMempoolTransactionStatsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMempoolTransactionStatsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMempoolTransactionStatsResponse(val *MempoolTransactionStatsResponse) *NullableMempoolTransactionStatsResponse {
	return &NullableMempoolTransactionStatsResponse{value: val, isSet: true}
}

func (v NullableMempoolTransactionStatsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMempoolTransactionStatsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

