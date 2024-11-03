package transaction

import (
	"strings"
	"testing"

	"github.com/icon-project/stacks-go-sdk/pkg/clarity"
	"github.com/stretchr/testify/assert"
)

func TestTokenTransferPayloadSerializationDeserialization(t *testing.T) {
	recipient := "SP3FGQ8Z7JY9BWYZ5WM53E0M9NK7WHJF0691NZ159"
	amount := uint64(2500000)
	memo := "test memo"

	payload, err := NewTokenTransferPayload(recipient, amount, memo)
	assert.NoError(t, err)

	serialized, err := payload.Serialize()
	assert.NoError(t, err)

	deserialized := &TokenTransferPayload{}
	_, err = deserialized.Deserialize(serialized)
	assert.NoError(t, err)

	assert.Equal(t, payload.Recipient, deserialized.Recipient)
	assert.Equal(t, payload.Amount, deserialized.Amount)
	assert.Equal(t, payload.Memo, deserialized.Memo)
}

func TestTokenTransferPayloadToContractAddress(t *testing.T) {
	recipient := "SP3FGQ8Z7JY9BWYZ5WM53E0M9NK7WHJF0691NZ159.contract-name"
	amount := uint64(2500000)
	memo := "test memo"

	payload, err := NewTokenTransferPayload(recipient, amount, memo)
	assert.NoError(t, err)

	serialized, err := payload.Serialize()
	assert.NoError(t, err)

	deserialized := &TokenTransferPayload{}
	_, err = deserialized.Deserialize(serialized)
	assert.NoError(t, err)

	assert.Equal(t, payload.Recipient, deserialized.Recipient)
	assert.Equal(t, payload.Amount, deserialized.Amount)
	assert.Equal(t, payload.Memo, deserialized.Memo)
}

func TestContractCallPayloadSerializationDeserialization(t *testing.T) {
	contractAddress := "ST3FGQ8Z7JY9BWYZ5WM53E0M9NK7WHJF0691NZ159"
	contractName := "contract_name"
	functionName := "function_name"
	args := []clarity.ClarityValue{
		clarity.NewBool(true),
		clarity.NewBool(false),
	}

	payload, err := NewContractCallPayload(contractAddress, contractName, functionName, args)
	assert.NoError(t, err)

	serialized, err := payload.Serialize()
	assert.NoError(t, err)

	deserialized := &ContractCallPayload{}
	_, err = deserialized.Deserialize(serialized)
	assert.NoError(t, err)

	assert.Equal(t, payload.ContractAddress, deserialized.ContractAddress)
	assert.Equal(t, payload.ContractName, deserialized.ContractName)
	assert.Equal(t, payload.FunctionName, deserialized.FunctionName)
	assert.Equal(t, len(payload.FunctionArgs), len(deserialized.FunctionArgs))
	for i, arg := range payload.FunctionArgs {
		assert.Equal(t, arg.Type(), deserialized.FunctionArgs[i].Type())
	}
}

func TestSmartContractPayloadValidation(t *testing.T) {
	tests := []struct {
		name          string
		contractName  string
		codeBody      string
		expectedError string
	}{
		{
			name:          "Valid contract",
			contractName:  "my-contract",
			codeBody:      "(define-data-var counter int 0)\n(define-public (increment)\n\t(ok (var-set counter (+ (var-get counter) 1))))",
			expectedError: "",
		},
		{
			name:          "Empty contract name",
			contractName:  "",
			codeBody:      "(define-data-var counter int 0)",
			expectedError: "contract name must be between 1 and 128 characters",
		},
		{
			name:          "Contract name too long",
			contractName:  strings.Repeat("a", 129),
			codeBody:      "(define-data-var counter int 0)",
			expectedError: "contract name must be between 1 and 128 characters",
		},
		{
			name:          "Invalid contract name start",
			contractName:  "1invalid",
			codeBody:      "(define-data-var counter int 0)",
			expectedError: "contract name must start with a letter",
		},
		{
			name:          "Invalid contract name characters",
			contractName:  "invalid@name",
			codeBody:      "(define-data-var counter int 0)",
			expectedError: "contract name can only contain letters, numbers, hyphens, and underscores",
		},
		{
			name:          "Invalid code body characters",
			contractName:  "valid-name",
			codeBody:      "invalid\x00chars",
			expectedError: "invalid character in code body",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewSmartContractPayload(tt.contractName, tt.codeBody)
			if tt.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			}
		})
	}
}

func TestSmartContractPayloadSerializationAndDeserialization(t *testing.T) {
	contractName := "test-contract"
	codeBody := `(define-data-var counter int 0)
(define-public (increment)
    (ok (var-set counter (+ (var-get counter) 1))))
`
	payload, err := NewSmartContractPayload(contractName, codeBody)
	assert.NoError(t, err)

	serialized, err := payload.Serialize()
	assert.NoError(t, err)

	deserialized := &SmartContractPayload{}
	_, err = deserialized.Deserialize(serialized)
	assert.NoError(t, err)

	assert.Equal(t, payload.ContractName, deserialized.ContractName)
	assert.Equal(t, payload.CodeBody, deserialized.CodeBody)
}
