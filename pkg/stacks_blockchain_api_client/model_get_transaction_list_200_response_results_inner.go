/*
Stacks Blockchain API

Welcome to the API reference overview for the [Stacks Blockchain API](https://docs.hiro.so/stacks-blockchain-api).        [Download Postman collection](https://hirosystems.github.io/stacks-blockchain-api/collection.json)

API version: v8.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stacks_blockchain_api_client

import (
	"encoding/json"
	"fmt"
)


// GetTransactionList200ResponseResultsInner struct for GetTransactionList200ResponseResultsInner
type GetTransactionList200ResponseResultsInner struct {
	CoinbaseTransaction *CoinbaseTransaction
	ContractCallTransaction *ContractCallTransaction
	PoisonMicroblockTransaction *PoisonMicroblockTransaction
	SmartContractTransaction *SmartContractTransaction
	TenureChangeTransaction *TenureChangeTransaction
	TokenTransferTransaction *TokenTransferTransaction
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *GetTransactionList200ResponseResultsInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into CoinbaseTransaction
	err = json.Unmarshal(data, &dst.CoinbaseTransaction);
	if err == nil {
		jsonCoinbaseTransaction, _ := json.Marshal(dst.CoinbaseTransaction)
		if string(jsonCoinbaseTransaction) == "{}" { // empty struct
			dst.CoinbaseTransaction = nil
		} else {
			return nil // data stored in dst.CoinbaseTransaction, return on the first match
		}
	} else {
		dst.CoinbaseTransaction = nil
	}

	// try to unmarshal JSON data into ContractCallTransaction
	err = json.Unmarshal(data, &dst.ContractCallTransaction);
	if err == nil {
		jsonContractCallTransaction, _ := json.Marshal(dst.ContractCallTransaction)
		if string(jsonContractCallTransaction) == "{}" { // empty struct
			dst.ContractCallTransaction = nil
		} else {
			return nil // data stored in dst.ContractCallTransaction, return on the first match
		}
	} else {
		dst.ContractCallTransaction = nil
	}

	// try to unmarshal JSON data into PoisonMicroblockTransaction
	err = json.Unmarshal(data, &dst.PoisonMicroblockTransaction);
	if err == nil {
		jsonPoisonMicroblockTransaction, _ := json.Marshal(dst.PoisonMicroblockTransaction)
		if string(jsonPoisonMicroblockTransaction) == "{}" { // empty struct
			dst.PoisonMicroblockTransaction = nil
		} else {
			return nil // data stored in dst.PoisonMicroblockTransaction, return on the first match
		}
	} else {
		dst.PoisonMicroblockTransaction = nil
	}

	// try to unmarshal JSON data into SmartContractTransaction
	err = json.Unmarshal(data, &dst.SmartContractTransaction);
	if err == nil {
		jsonSmartContractTransaction, _ := json.Marshal(dst.SmartContractTransaction)
		if string(jsonSmartContractTransaction) == "{}" { // empty struct
			dst.SmartContractTransaction = nil
		} else {
			return nil // data stored in dst.SmartContractTransaction, return on the first match
		}
	} else {
		dst.SmartContractTransaction = nil
	}

	// try to unmarshal JSON data into TenureChangeTransaction
	err = json.Unmarshal(data, &dst.TenureChangeTransaction);
	if err == nil {
		jsonTenureChangeTransaction, _ := json.Marshal(dst.TenureChangeTransaction)
		if string(jsonTenureChangeTransaction) == "{}" { // empty struct
			dst.TenureChangeTransaction = nil
		} else {
			return nil // data stored in dst.TenureChangeTransaction, return on the first match
		}
	} else {
		dst.TenureChangeTransaction = nil
	}

	// try to unmarshal JSON data into TokenTransferTransaction
	err = json.Unmarshal(data, &dst.TokenTransferTransaction);
	if err == nil {
		jsonTokenTransferTransaction, _ := json.Marshal(dst.TokenTransferTransaction)
		if string(jsonTokenTransferTransaction) == "{}" { // empty struct
			dst.TokenTransferTransaction = nil
		} else {
			return nil // data stored in dst.TokenTransferTransaction, return on the first match
		}
	} else {
		dst.TokenTransferTransaction = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(GetTransactionList200ResponseResultsInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *GetTransactionList200ResponseResultsInner) MarshalJSON() ([]byte, error) {
	if src.CoinbaseTransaction != nil {
		return json.Marshal(&src.CoinbaseTransaction)
	}

	if src.ContractCallTransaction != nil {
		return json.Marshal(&src.ContractCallTransaction)
	}

	if src.PoisonMicroblockTransaction != nil {
		return json.Marshal(&src.PoisonMicroblockTransaction)
	}

	if src.SmartContractTransaction != nil {
		return json.Marshal(&src.SmartContractTransaction)
	}

	if src.TenureChangeTransaction != nil {
		return json.Marshal(&src.TenureChangeTransaction)
	}

	if src.TokenTransferTransaction != nil {
		return json.Marshal(&src.TokenTransferTransaction)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableGetTransactionList200ResponseResultsInner struct {
	value *GetTransactionList200ResponseResultsInner
	isSet bool
}

func (v NullableGetTransactionList200ResponseResultsInner) Get() *GetTransactionList200ResponseResultsInner {
	return v.value
}

func (v *NullableGetTransactionList200ResponseResultsInner) Set(val *GetTransactionList200ResponseResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionList200ResponseResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionList200ResponseResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionList200ResponseResultsInner(val *GetTransactionList200ResponseResultsInner) *NullableGetTransactionList200ResponseResultsInner {
	return &NullableGetTransactionList200ResponseResultsInner{value: val, isSet: true}
}

func (v NullableGetTransactionList200ResponseResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionList200ResponseResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


