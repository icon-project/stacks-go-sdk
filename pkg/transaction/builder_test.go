package transaction

import (
	"bytes"
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/icon-project/stacks-go-sdk/pkg/clarity"
	"github.com/icon-project/stacks-go-sdk/pkg/crypto"
	"github.com/icon-project/stacks-go-sdk/pkg/stacks"
	"github.com/stretchr/testify/assert"
)

func TestMakeSTXTokenTransfer(t *testing.T) {
	recipient := "ST3YJD5Y1WTMC8R09ZKR3HJF562R3NM8HHXW2S2R9"
	amount := big.NewInt(1000000) // 1 STX
	memo := "Test transfer"
	network := stacks.NewStacksTestnet()
	senderAddress := "ST15C893XJFJ6FSKM020P9JQDB5T7X6MQTXMBPAVH"
	senderKeyHex := "c1d5bb638aa70862621667f9997711fce692cad782694103f8d9561f62e9f19701"
	senderKey, _ := hex.DecodeString(senderKeyHex)

	tx, err := MakeSTXTokenTransfer(recipient, *amount, memo, *network, senderAddress, senderKey, nil, nil)
	if err != nil {
		t.Fatalf("makeSTXTokenTransfer failed: %v", err)
	}

	expectedRecipient, err := clarity.StringToPrincipal(recipient)
	if err != nil {
		t.Fatalf("failed to cast recipient string to clarity principal: %v", err)
	}

	recipientPrincipal, _ := tx.Payload.Recipient.(*clarity.StandardPrincipal)
	expectedPrincipal, _ := expectedRecipient.(*clarity.StandardPrincipal)

	if !bytes.Equal(recipientPrincipal.Hash160[:], expectedPrincipal.Hash160[:]) {
		t.Errorf("Expected recipient %s, got %s", expectedRecipient, tx.Payload.Recipient)
	}

	if tx.Payload.Amount != amount.Uint64() {
		t.Errorf("Expected amount %d, got %d", amount.Uint64(), tx.Payload.Amount)
	}

	if tx.Payload.Memo != memo {
		t.Errorf("Expected memo %s, got %s", memo, tx.Payload.Memo)
	}

	if tx.Auth.OriginAuth.Fee == 0 {
		t.Error("Fee was not estimated")
	}

	if tx.Auth.OriginAuth.Nonce == 0 {
		t.Error("Nonce was not fetched")
	}

	if len(tx.Auth.OriginAuth.Signature) == 0 {
		t.Error("Transaction was not signed")
	}

	t.Logf("Estimated fee: %d", tx.Auth.OriginAuth.Fee)
	t.Logf("Fetched nonce: %d", tx.Auth.OriginAuth.Nonce)

	specifiedFee := big.NewInt(180)
	specifiedNonce := big.NewInt(3)
	tx2, err := MakeSTXTokenTransfer(recipient, *amount, memo, *network, senderAddress, senderKey, specifiedFee, specifiedNonce)
	if err != nil {
		t.Fatalf("makeSTXTokenTransfer with specified fee and nonce failed: %v", err)
	}

	if tx2.Auth.OriginAuth.Fee != specifiedFee.Uint64() {
		t.Errorf("Expected fee %d, got %d", specifiedFee.Uint64(), tx2.Auth.OriginAuth.Fee)
	}

	if tx2.Auth.OriginAuth.Nonce != specifiedNonce.Uint64() {
		t.Errorf("Expected nonce %d, got %d", specifiedNonce.Uint64(), tx2.Auth.OriginAuth.Nonce)
	}
}

func TestMakeContractDeploy(t *testing.T) {
	contractName := "test-contract"
	codeBody := `(define-data-var counter int 0)
(define-public (increment)
    (ok (var-set counter (+ (var-get counter) 1))))
`
	network := stacks.NewStacksTestnet()
	senderAddress := "ST15C893XJFJ6FSKM020P9JQDB5T7X6MQTXMBPAVH"
	senderKeyHex := "c1d5bb638aa70862621667f9997711fce692cad782694103f8d9561f62e9f19701"
	senderKey, _ := hex.DecodeString(senderKeyHex)

	tests := []struct {
		name         string
		contractName string
		codeBody     string
		fee          *big.Int
		nonce        *big.Int
		expectedErr  bool
	}{
		{
			name:         "Valid contract with auto fee and nonce",
			contractName: contractName,
			codeBody:     codeBody,
			fee:          nil,
			nonce:        nil,
			expectedErr:  false,
		},
		{
			name:         "Valid contract with specified fee and nonce",
			contractName: contractName,
			codeBody:     codeBody,
			fee:          big.NewInt(1000),
			nonce:        big.NewInt(1),
			expectedErr:  false,
		},
		{
			name:         "Empty contract name",
			contractName: "",
			codeBody:     codeBody,
			fee:          nil,
			nonce:        nil,
			expectedErr:  true,
		},
		{
			name:         "Empty code body",
			contractName: contractName,
			codeBody:     "",
			fee:          nil,
			nonce:        nil,
			expectedErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, err := MakeContractDeploy(
				tt.contractName,
				tt.codeBody,
				*network,
				senderAddress,
				senderKey,
				tt.fee,
				tt.nonce,
			)

			if tt.expectedErr {
				assert.Error(t, err)
				assert.Nil(t, tx)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, tx)

			assert.Equal(t, tt.contractName, tx.Payload.ContractName)
			assert.Equal(t, tt.codeBody, tx.Payload.CodeBody)

			if tt.fee != nil {
				assert.Equal(t, tt.fee.Uint64(), tx.Auth.OriginAuth.Fee)
			} else {
				assert.Greater(t, tx.Auth.OriginAuth.Fee, uint64(0), "Fee should be estimated")
			}

			if tt.nonce != nil {
				assert.Equal(t, tt.nonce.Uint64(), tx.Auth.OriginAuth.Nonce)
			} else {
				assert.Greater(t, tx.Auth.OriginAuth.Nonce, uint64(0), "Nonce should be fetched")
			}

			senderPublicKey := crypto.GetPublicKeyFromPrivate(senderKey)
			isValid, err := VerifyTransaction(tx, senderPublicKey)
			assert.NoError(t, err)
			assert.True(t, isValid)

			serialized, err := tx.Serialize()
			assert.NoError(t, err)

			deserialized, err := DeserializeTransaction(serialized)
			assert.NoError(t, err)

			contractTx, ok := deserialized.(*SmartContractTransaction)
			assert.True(t, ok)
			assert.Equal(t, tx.Payload.ContractName, contractTx.Payload.ContractName)
			assert.Equal(t, tx.Payload.CodeBody, contractTx.Payload.CodeBody)
		})
	}
}
