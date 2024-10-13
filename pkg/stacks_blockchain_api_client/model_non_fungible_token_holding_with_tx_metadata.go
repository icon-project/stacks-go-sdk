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

// checks if the NonFungibleTokenHoldingWithTxMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NonFungibleTokenHoldingWithTxMetadata{}

// NonFungibleTokenHoldingWithTxMetadata Ownership of a Non-Fungible Token with transaction metadata
type NonFungibleTokenHoldingWithTxMetadata struct {
	AssetIdentifier string `json:"asset_identifier"`
	Value NonFungibleTokenHoldingWithTxIdValue `json:"value"`
	BlockHeight int32 `json:"block_height"`
	Tx GetTransactionList200ResponseResultsInner `json:"tx"`
}

type _NonFungibleTokenHoldingWithTxMetadata NonFungibleTokenHoldingWithTxMetadata

// NewNonFungibleTokenHoldingWithTxMetadata instantiates a new NonFungibleTokenHoldingWithTxMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonFungibleTokenHoldingWithTxMetadata(assetIdentifier string, value NonFungibleTokenHoldingWithTxIdValue, blockHeight int32, tx GetTransactionList200ResponseResultsInner) *NonFungibleTokenHoldingWithTxMetadata {
	this := NonFungibleTokenHoldingWithTxMetadata{}
	this.AssetIdentifier = assetIdentifier
	this.Value = value
	this.BlockHeight = blockHeight
	this.Tx = tx
	return &this
}

// NewNonFungibleTokenHoldingWithTxMetadataWithDefaults instantiates a new NonFungibleTokenHoldingWithTxMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonFungibleTokenHoldingWithTxMetadataWithDefaults() *NonFungibleTokenHoldingWithTxMetadata {
	this := NonFungibleTokenHoldingWithTxMetadata{}
	return &this
}

// GetAssetIdentifier returns the AssetIdentifier field value
func (o *NonFungibleTokenHoldingWithTxMetadata) GetAssetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetIdentifier
}

// GetAssetIdentifierOk returns a tuple with the AssetIdentifier field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenHoldingWithTxMetadata) GetAssetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetIdentifier, true
}

// SetAssetIdentifier sets field value
func (o *NonFungibleTokenHoldingWithTxMetadata) SetAssetIdentifier(v string) {
	o.AssetIdentifier = v
}

// GetValue returns the Value field value
func (o *NonFungibleTokenHoldingWithTxMetadata) GetValue() NonFungibleTokenHoldingWithTxIdValue {
	if o == nil {
		var ret NonFungibleTokenHoldingWithTxIdValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenHoldingWithTxMetadata) GetValueOk() (*NonFungibleTokenHoldingWithTxIdValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *NonFungibleTokenHoldingWithTxMetadata) SetValue(v NonFungibleTokenHoldingWithTxIdValue) {
	o.Value = v
}

// GetBlockHeight returns the BlockHeight field value
func (o *NonFungibleTokenHoldingWithTxMetadata) GetBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BlockHeight
}

// GetBlockHeightOk returns a tuple with the BlockHeight field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenHoldingWithTxMetadata) GetBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHeight, true
}

// SetBlockHeight sets field value
func (o *NonFungibleTokenHoldingWithTxMetadata) SetBlockHeight(v int32) {
	o.BlockHeight = v
}

// GetTx returns the Tx field value
func (o *NonFungibleTokenHoldingWithTxMetadata) GetTx() GetTransactionList200ResponseResultsInner {
	if o == nil {
		var ret GetTransactionList200ResponseResultsInner
		return ret
	}

	return o.Tx
}

// GetTxOk returns a tuple with the Tx field value
// and a boolean to check if the value has been set.
func (o *NonFungibleTokenHoldingWithTxMetadata) GetTxOk() (*GetTransactionList200ResponseResultsInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tx, true
}

// SetTx sets field value
func (o *NonFungibleTokenHoldingWithTxMetadata) SetTx(v GetTransactionList200ResponseResultsInner) {
	o.Tx = v
}

func (o NonFungibleTokenHoldingWithTxMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonFungibleTokenHoldingWithTxMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asset_identifier"] = o.AssetIdentifier
	toSerialize["value"] = o.Value
	toSerialize["block_height"] = o.BlockHeight
	toSerialize["tx"] = o.Tx
	return toSerialize, nil
}

func (o *NonFungibleTokenHoldingWithTxMetadata) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"asset_identifier",
		"value",
		"block_height",
		"tx",
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

	varNonFungibleTokenHoldingWithTxMetadata := _NonFungibleTokenHoldingWithTxMetadata{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNonFungibleTokenHoldingWithTxMetadata)

	if err != nil {
		return err
	}

	*o = NonFungibleTokenHoldingWithTxMetadata(varNonFungibleTokenHoldingWithTxMetadata)

	return err
}

type NullableNonFungibleTokenHoldingWithTxMetadata struct {
	value *NonFungibleTokenHoldingWithTxMetadata
	isSet bool
}

func (v NullableNonFungibleTokenHoldingWithTxMetadata) Get() *NonFungibleTokenHoldingWithTxMetadata {
	return v.value
}

func (v *NullableNonFungibleTokenHoldingWithTxMetadata) Set(val *NonFungibleTokenHoldingWithTxMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableNonFungibleTokenHoldingWithTxMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableNonFungibleTokenHoldingWithTxMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonFungibleTokenHoldingWithTxMetadata(val *NonFungibleTokenHoldingWithTxMetadata) *NullableNonFungibleTokenHoldingWithTxMetadata {
	return &NullableNonFungibleTokenHoldingWithTxMetadata{value: val, isSet: true}
}

func (v NullableNonFungibleTokenHoldingWithTxMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonFungibleTokenHoldingWithTxMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

