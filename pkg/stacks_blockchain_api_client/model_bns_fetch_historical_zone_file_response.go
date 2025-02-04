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

// checks if the BnsFetchHistoricalZoneFileResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BnsFetchHistoricalZoneFileResponse{}

// BnsFetchHistoricalZoneFileResponse Fetches the historical zonefile specified by the username and zone hash.
type BnsFetchHistoricalZoneFileResponse struct {
	Zonefile string `json:"zonefile"`
}

type _BnsFetchHistoricalZoneFileResponse BnsFetchHistoricalZoneFileResponse

// NewBnsFetchHistoricalZoneFileResponse instantiates a new BnsFetchHistoricalZoneFileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBnsFetchHistoricalZoneFileResponse(zonefile string) *BnsFetchHistoricalZoneFileResponse {
	this := BnsFetchHistoricalZoneFileResponse{}
	this.Zonefile = zonefile
	return &this
}

// NewBnsFetchHistoricalZoneFileResponseWithDefaults instantiates a new BnsFetchHistoricalZoneFileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBnsFetchHistoricalZoneFileResponseWithDefaults() *BnsFetchHistoricalZoneFileResponse {
	this := BnsFetchHistoricalZoneFileResponse{}
	return &this
}

// GetZonefile returns the Zonefile field value
func (o *BnsFetchHistoricalZoneFileResponse) GetZonefile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Zonefile
}

// GetZonefileOk returns a tuple with the Zonefile field value
// and a boolean to check if the value has been set.
func (o *BnsFetchHistoricalZoneFileResponse) GetZonefileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Zonefile, true
}

// SetZonefile sets field value
func (o *BnsFetchHistoricalZoneFileResponse) SetZonefile(v string) {
	o.Zonefile = v
}

func (o BnsFetchHistoricalZoneFileResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BnsFetchHistoricalZoneFileResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["zonefile"] = o.Zonefile
	return toSerialize, nil
}

func (o *BnsFetchHistoricalZoneFileResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"zonefile",
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

	varBnsFetchHistoricalZoneFileResponse := _BnsFetchHistoricalZoneFileResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBnsFetchHistoricalZoneFileResponse)

	if err != nil {
		return err
	}

	*o = BnsFetchHistoricalZoneFileResponse(varBnsFetchHistoricalZoneFileResponse)

	return err
}

type NullableBnsFetchHistoricalZoneFileResponse struct {
	value *BnsFetchHistoricalZoneFileResponse
	isSet bool
}

func (v NullableBnsFetchHistoricalZoneFileResponse) Get() *BnsFetchHistoricalZoneFileResponse {
	return v.value
}

func (v *NullableBnsFetchHistoricalZoneFileResponse) Set(val *BnsFetchHistoricalZoneFileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBnsFetchHistoricalZoneFileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBnsFetchHistoricalZoneFileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBnsFetchHistoricalZoneFileResponse(val *BnsFetchHistoricalZoneFileResponse) *NullableBnsFetchHistoricalZoneFileResponse {
	return &NullableBnsFetchHistoricalZoneFileResponse{value: val, isSet: true}
}

func (v NullableBnsFetchHistoricalZoneFileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBnsFetchHistoricalZoneFileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


