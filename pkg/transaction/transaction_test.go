package transaction

import (
	"encoding/hex"
	"testing"

	"github.com/icon-project/stacks-go-sdk/pkg/clarity"
	"github.com/icon-project/stacks-go-sdk/pkg/crypto"
	"github.com/icon-project/stacks-go-sdk/pkg/stacks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSTXTokenTransferTransactionSerializationAndDeserialization(t *testing.T) {
	transactionVersion := stacks.TransactionVersionTestnet
	chainID := stacks.ChainIDTestnet
	anchorMode := stacks.AnchorModeOnChainOnly
	postConditionMode := stacks.PostConditionModeDeny

	recipientAddress := "SP3FGQ8Z7JY9BWYZ5WM53E0M9NK7WHJF0691NZ159"
	amount := uint64(2500000)
	memo := "memo (not included)"

	payload, err := NewTokenTransferPayload(recipientAddress, amount, memo)
	assert.NoError(t, err)

	addressHashMode := stacks.AddressHashModeSerializeP2PKH
	nonce := uint64(0)
	fee := uint64(0)

	pubKey := "03ef788b3830c00abe8f64f62dc32fc863bc0b2cafeb073b6c8e1c7657d9c2c3ab"
	pubKeyBytes, err := hex.DecodeString(pubKey)
	assert.NoError(t, err)

	secretKey := "edf9aee84d9b7abc145504dde6726c64f369d37ee34ded868fabd876c26570bc01"
	secretKeyBytes, err := hex.DecodeString(secretKey)
	assert.NoError(t, err)

	spendingCondition := SpendingCondition{
		HashMode:    addressHashMode,
		Signer:      [20]byte{},
		Nonce:       nonce,
		Fee:         fee,
		KeyEncoding: stacks.PubKeyEncodingCompressed,
		Signature:   [65]byte{},
	}

	auth := TransactionAuth{
		AuthType:   stacks.AuthTypeStandard,
		OriginAuth: spendingCondition,
	}

	transaction := &TokenTransferTransaction{
		BaseTransaction: BaseTransaction{
			Version:           transactionVersion,
			ChainID:           chainID,
			Auth:              auth,
			AnchorMode:        anchorMode,
			PostConditionMode: postConditionMode,
			PostConditions:    []PostCondition{},
		},
		Payload: *payload,
	}

	err = SignTransaction(transaction, secretKeyBytes)
	assert.NoError(t, err)

	isValid, err := VerifyTransaction(transaction, pubKeyBytes)
	assert.NoError(t, err)

	assert.True(t, isValid)

	serialized, err := transaction.Serialize()
	assert.NoError(t, err)

	deserialized, err := DeserializeTransaction(serialized)
	assert.NoError(t, err)

	tokenTx, ok := deserialized.(*TokenTransferTransaction)
	assert.True(t, ok, "Deserialized transaction is not a TokenTransferTransaction")

	assert.Equal(t, transactionVersion, tokenTx.Version)
	assert.Equal(t, chainID, tokenTx.ChainID)
	assert.Equal(t, stacks.AuthTypeStandard, tokenTx.Auth.AuthType)
	assert.Equal(t, addressHashMode, tokenTx.Auth.OriginAuth.HashMode)
	assert.Equal(t, nonce, tokenTx.Auth.OriginAuth.Nonce)
	assert.Equal(t, fee, tokenTx.Auth.OriginAuth.Fee)
	assert.Equal(t, anchorMode, tokenTx.AnchorMode)
	assert.Equal(t, postConditionMode, tokenTx.PostConditionMode)
	assert.Empty(t, tokenTx.PostConditions)

	recipientPrincipal, _ := clarity.StringToPrincipal(recipientAddress)

	assert.Equal(t, recipientPrincipal, tokenTx.Payload.Recipient)
	assert.Equal(t, amount, tokenTx.Payload.Amount)
	assert.Equal(t, memo, tokenTx.Payload.Memo)

	isValid, err = VerifyTransaction(tokenTx, pubKeyBytes)
	assert.NoError(t, err)
	assert.True(t, isValid)

	reserializedBytes, err := deserialized.Serialize()
	assert.NoError(t, err)
	assert.Equal(t, serialized, reserializedBytes)
}

func TestNewTokenTransferTransaction(t *testing.T) {
	tests := []struct {
		name              string
		recipient         string
		amount            uint64
		memo              string
		version           stacks.TransactionVersion
		chainID           stacks.ChainID
		signer            [20]byte
		nonce             uint64
		fee               uint64
		anchorMode        stacks.AnchorMode
		postConditionMode stacks.PostConditionMode
	}{
		{
			name:              "Valid transaction",
			recipient:         "ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM",
			amount:            1000000,
			memo:              "Test transfer",
			version:           stacks.TransactionVersionMainnet,
			chainID:           stacks.ChainIDMainnet,
			signer:            [20]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			nonce:             1,
			fee:               1000,
			anchorMode:        stacks.AnchorModeOnChainOnly,
			postConditionMode: stacks.PostConditionModeAllow,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, err := NewTokenTransferTransaction(
				tt.recipient,
				tt.amount,
				tt.memo,
				tt.version,
				tt.chainID,
				tt.signer,
				tt.nonce,
				tt.fee,
				tt.anchorMode,
				tt.postConditionMode,
			)

			require.NoError(t, err)
			require.NotNil(t, tx)

			// Check initial values
			assertTokenTransferTransactionFields(t, tx, tt)

			// Serialize the transaction
			serialized, err := tx.Serialize()
			require.NoError(t, err)
			require.NotEmpty(t, serialized)

			// Deserialize the transaction
			deserialized, err := DeserializeTransaction(serialized)
			require.NoError(t, err)
			require.NotNil(t, deserialized)

			// Check that the deserialized transaction is of the correct type
			deserializedTx, ok := deserialized.(*TokenTransferTransaction)
			require.True(t, ok, "Deserialized transaction is not a TokenTransferTransaction")

			// Check that all fields are preserved after serialization and deserialization
			assertTokenTransferTransactionFields(t, deserializedTx, tt)
		})
	}
}

func TestContractCallTransactionSerializationAndDeserialization(t *testing.T) {
	transactionVersion := stacks.TransactionVersionTestnet
	chainID := stacks.ChainIDTestnet
	anchorMode := stacks.AnchorModeOnChainOnly
	postConditionMode := stacks.PostConditionModeDeny

	contractAddress := "ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM"
	contractName := "test-contract"
	functionName := "test-function"

	intArg, err := clarity.NewInt(1)
	require.NoError(t, err, "Failed to create int argument")

	strArg, err := clarity.NewStringASCII("test")
	require.NoError(t, err, "Failed to create string argument")

	functionArgs := []clarity.ClarityValue{
		intArg,
		clarity.NewBool(true),
		strArg,
	}

	addressHashMode := stacks.AddressHashModeSerializeP2PKH
	nonce := uint64(10)
	fee := uint64(456)

	pubKey := "0332fc778e5beb5f944c75b2b63c21dd12c40bdcdf99ba0663168ae0b2be880aef"
	pubKeyBytes, err := hex.DecodeString(pubKey)
	assert.NoError(t, err)

	secretKey := "c1d5bb638aa70862621667f9997711fce692cad782694103f8d9561f62e9f19701"
	secretKeyBytes, err := hex.DecodeString(secretKey)
	assert.NoError(t, err)

	senderPublicKey := crypto.GetPublicKeyFromPrivate(secretKeyBytes)
	var signer [20]byte
	copy(signer[:], crypto.Hash160(senderPublicKey))

	transaction, err := NewContractCallTransaction(contractAddress, contractName, functionName, functionArgs, transactionVersion, chainID, signer, nonce, fee, anchorMode, postConditionMode)
	assert.NoError(t, err)

	err = SignTransaction(transaction, secretKeyBytes)
	assert.NoError(t, err)

	isValid, err := VerifyTransaction(transaction, pubKeyBytes)
	assert.NoError(t, err)
	assert.True(t, isValid)

	serialized, err := transaction.Serialize()
	assert.NoError(t, err)

	deserialized, err := DeserializeTransaction(serialized)
	assert.NoError(t, err)

	contractTx, ok := deserialized.(*ContractCallTransaction)
	assert.True(t, ok, "Deserialized transaction is not a ContractCallTransaction")

	assert.Equal(t, transactionVersion, contractTx.Version)
	assert.Equal(t, chainID, contractTx.ChainID)
	assert.Equal(t, stacks.AuthTypeStandard, contractTx.Auth.AuthType)
	assert.Equal(t, addressHashMode, contractTx.Auth.OriginAuth.HashMode)
	assert.Equal(t, nonce, contractTx.Auth.OriginAuth.Nonce)
	assert.Equal(t, fee, contractTx.Auth.OriginAuth.Fee)
	assert.Equal(t, anchorMode, contractTx.AnchorMode)
	assert.Equal(t, postConditionMode, contractTx.PostConditionMode)
	assert.Empty(t, contractTx.PostConditions)

	contractPrincipal, _ := clarity.StringToPrincipal(contractAddress)

	assert.Equal(t, contractPrincipal, contractTx.Payload.ContractAddress)
	assert.Equal(t, contractName, contractTx.Payload.ContractName)
	assert.Equal(t, functionName, contractTx.Payload.FunctionName)
	assert.Equal(t, len(functionArgs), len(contractTx.Payload.FunctionArgs))

	for i, arg := range functionArgs {
		assert.Equal(t, arg.Type(), contractTx.Payload.FunctionArgs[i].Type())
		switch typedArg := arg.(type) {
		case *clarity.Int:
			deserializedArg, ok := contractTx.Payload.FunctionArgs[i].(*clarity.Int)
			assert.True(t, ok)
			assert.Equal(t, 0, typedArg.Value.Cmp(deserializedArg.Value))
		case clarity.Bool:
			deserializedArg, ok := contractTx.Payload.FunctionArgs[i].(clarity.Bool)
			assert.True(t, ok)
			assert.Equal(t, typedArg, deserializedArg)
		case *clarity.StringASCII:
			deserializedArg, ok := contractTx.Payload.FunctionArgs[i].(*clarity.StringASCII)
			assert.True(t, ok)
			assert.Equal(t, typedArg.Data, deserializedArg.Data)
		}
	}

	isValid, err = VerifyTransaction(contractTx, pubKeyBytes)
	assert.NoError(t, err)
	assert.True(t, isValid)

	reserializedBytes, err := deserialized.Serialize()
	assert.NoError(t, err)
	assert.Equal(t, serialized, reserializedBytes)
}

func assertTokenTransferTransactionFields(t *testing.T, tx *TokenTransferTransaction, expected struct {
	name              string
	recipient         string
	amount            uint64
	memo              string
	version           stacks.TransactionVersion
	chainID           stacks.ChainID
	signer            [20]byte
	nonce             uint64
	fee               uint64
	anchorMode        stacks.AnchorMode
	postConditionMode stacks.PostConditionMode
}) {
	assert.Equal(t, expected.version, tx.Version)
	assert.Equal(t, expected.chainID, tx.ChainID)
	assert.Equal(t, expected.signer, tx.Auth.OriginAuth.Signer)
	assert.Equal(t, expected.nonce, tx.Auth.OriginAuth.Nonce)
	assert.Equal(t, expected.fee, tx.Auth.OriginAuth.Fee)
	assert.Equal(t, expected.anchorMode, tx.AnchorMode)
	assert.Equal(t, expected.postConditionMode, tx.PostConditionMode)
	assert.Equal(t, expected.amount, tx.Payload.Amount)
	assert.Equal(t, expected.memo, tx.Payload.Memo)

	recipient := tx.Payload.Recipient
	expectedRecipient, _ := clarity.StringToPrincipal(expected.recipient)
	assert.Equal(t, expectedRecipient, recipient)
}
