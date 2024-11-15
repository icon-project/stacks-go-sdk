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


// TransactionFoundResultAnyOf1 struct for TransactionFoundResultAnyOf1
type TransactionFoundResultAnyOf1 struct {
	CoinbaseMempoolTransaction *CoinbaseMempoolTransaction
	ContractCallMempoolTransaction *ContractCallMempoolTransaction
	PoisonMicroblockMempoolTransaction *PoisonMicroblockMempoolTransaction
	SmartContractMempoolTransaction *SmartContractMempoolTransaction
	TenureChangeMempoolTransaction *TenureChangeMempoolTransaction
	TokenTransferMempoolTransaction *TokenTransferMempoolTransaction
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TransactionFoundResultAnyOf1) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into CoinbaseMempoolTransaction
	err = json.Unmarshal(data, &dst.CoinbaseMempoolTransaction);
	if err == nil {
		jsonCoinbaseMempoolTransaction, _ := json.Marshal(dst.CoinbaseMempoolTransaction)
		if string(jsonCoinbaseMempoolTransaction) == "{}" { // empty struct
			dst.CoinbaseMempoolTransaction = nil
		} else {
			return nil // data stored in dst.CoinbaseMempoolTransaction, return on the first match
		}
	} else {
		dst.CoinbaseMempoolTransaction = nil
	}

	// try to unmarshal JSON data into ContractCallMempoolTransaction
	err = json.Unmarshal(data, &dst.ContractCallMempoolTransaction);
	if err == nil {
		jsonContractCallMempoolTransaction, _ := json.Marshal(dst.ContractCallMempoolTransaction)
		if string(jsonContractCallMempoolTransaction) == "{}" { // empty struct
			dst.ContractCallMempoolTransaction = nil
		} else {
			return nil // data stored in dst.ContractCallMempoolTransaction, return on the first match
		}
	} else {
		dst.ContractCallMempoolTransaction = nil
	}

	// try to unmarshal JSON data into PoisonMicroblockMempoolTransaction
	err = json.Unmarshal(data, &dst.PoisonMicroblockMempoolTransaction);
	if err == nil {
		jsonPoisonMicroblockMempoolTransaction, _ := json.Marshal(dst.PoisonMicroblockMempoolTransaction)
		if string(jsonPoisonMicroblockMempoolTransaction) == "{}" { // empty struct
			dst.PoisonMicroblockMempoolTransaction = nil
		} else {
			return nil // data stored in dst.PoisonMicroblockMempoolTransaction, return on the first match
		}
	} else {
		dst.PoisonMicroblockMempoolTransaction = nil
	}

	// try to unmarshal JSON data into SmartContractMempoolTransaction
	err = json.Unmarshal(data, &dst.SmartContractMempoolTransaction);
	if err == nil {
		jsonSmartContractMempoolTransaction, _ := json.Marshal(dst.SmartContractMempoolTransaction)
		if string(jsonSmartContractMempoolTransaction) == "{}" { // empty struct
			dst.SmartContractMempoolTransaction = nil
		} else {
			return nil // data stored in dst.SmartContractMempoolTransaction, return on the first match
		}
	} else {
		dst.SmartContractMempoolTransaction = nil
	}

	// try to unmarshal JSON data into TenureChangeMempoolTransaction
	err = json.Unmarshal(data, &dst.TenureChangeMempoolTransaction);
	if err == nil {
		jsonTenureChangeMempoolTransaction, _ := json.Marshal(dst.TenureChangeMempoolTransaction)
		if string(jsonTenureChangeMempoolTransaction) == "{}" { // empty struct
			dst.TenureChangeMempoolTransaction = nil
		} else {
			return nil // data stored in dst.TenureChangeMempoolTransaction, return on the first match
		}
	} else {
		dst.TenureChangeMempoolTransaction = nil
	}

	// try to unmarshal JSON data into TokenTransferMempoolTransaction
	err = json.Unmarshal(data, &dst.TokenTransferMempoolTransaction);
	if err == nil {
		jsonTokenTransferMempoolTransaction, _ := json.Marshal(dst.TokenTransferMempoolTransaction)
		if string(jsonTokenTransferMempoolTransaction) == "{}" { // empty struct
			dst.TokenTransferMempoolTransaction = nil
		} else {
			return nil // data stored in dst.TokenTransferMempoolTransaction, return on the first match
		}
	} else {
		dst.TokenTransferMempoolTransaction = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(TransactionFoundResultAnyOf1)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TransactionFoundResultAnyOf1) MarshalJSON() ([]byte, error) {
	if src.CoinbaseMempoolTransaction != nil {
		return json.Marshal(&src.CoinbaseMempoolTransaction)
	}

	if src.ContractCallMempoolTransaction != nil {
		return json.Marshal(&src.ContractCallMempoolTransaction)
	}

	if src.PoisonMicroblockMempoolTransaction != nil {
		return json.Marshal(&src.PoisonMicroblockMempoolTransaction)
	}

	if src.SmartContractMempoolTransaction != nil {
		return json.Marshal(&src.SmartContractMempoolTransaction)
	}

	if src.TenureChangeMempoolTransaction != nil {
		return json.Marshal(&src.TenureChangeMempoolTransaction)
	}

	if src.TokenTransferMempoolTransaction != nil {
		return json.Marshal(&src.TokenTransferMempoolTransaction)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableTransactionFoundResultAnyOf1 struct {
	value *TransactionFoundResultAnyOf1
	isSet bool
}

func (v NullableTransactionFoundResultAnyOf1) Get() *TransactionFoundResultAnyOf1 {
	return v.value
}

func (v *NullableTransactionFoundResultAnyOf1) Set(val *TransactionFoundResultAnyOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionFoundResultAnyOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionFoundResultAnyOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionFoundResultAnyOf1(val *TransactionFoundResultAnyOf1) *NullableTransactionFoundResultAnyOf1 {
	return &NullableTransactionFoundResultAnyOf1{value: val, isSet: true}
}

func (v NullableTransactionFoundResultAnyOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionFoundResultAnyOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


