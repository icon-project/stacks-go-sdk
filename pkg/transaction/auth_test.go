package transaction

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/icon-project/stacks-go-sdk/pkg/stacks"
	"github.com/stretchr/testify/assert"
)

func TestSingleSpendingConditionSerializationAndDeserialization(t *testing.T) {
	addressHashMode := stacks.AddressHashModeSerializeP2PKH
	nonce := uint64(0)
	fee := uint64(0)

	spendingCondition := createSingleSigSpendingCondition(addressHashMode, nonce, fee)
	emptySignature := emptyMessageSignature()

	serialized, err := spendingCondition.Serialize()
	assert.NoError(t, err, "Failed to serialize spending condition")

	deserialized := SpendingCondition{}
	_, err = deserialized.Deserialize(serialized)
	assert.NoError(t, err, "Failed to deserialize spending condition")

	assert.Equal(t, addressHashMode, deserialized.HashMode, "HashMode mismatch")
	assert.Equal(t, nonce, deserialized.Nonce, "Nonce mismatch")
	assert.Equal(t, fee, deserialized.Fee, "Fee mismatch")
	assert.True(t, bytes.Equal(deserialized.Signature[:], emptySignature[:]), "Signature mismatch")
}

func TestSingleSigP2PKHSpendingCondition(t *testing.T) {
	spCompressed := createSingleSigSpendingCondition(stacks.AddressHashModeSerializeP2PKH, 345, 456)
	spCompressed.KeyEncoding = stacks.PubKeyEncodingCompressed
	spCompressed.Signature = createMessageSignature("fe")
	spCompressed.Signer = [20]byte{0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11}

	serialized, err := spCompressed.Serialize()
	assert.NoError(t, err, "Failed to serialize compressed P2PKH spending condition")

	deserialized := SpendingCondition{}
	_, err = deserialized.Deserialize(serialized)
	assert.NoError(t, err, "Failed to deserialize compressed P2PKH spending condition")

	assert.Equal(t, spCompressed, deserialized, "Compressed P2PKH spending condition mismatch")

	spUncompressed := createSingleSigSpendingCondition(stacks.AddressHashModeSerializeP2PKH, 123, 456)
	spUncompressed.KeyEncoding = stacks.PubKeyEncodingUncompressed
	spUncompressed.Signature = createMessageSignature("ff")
	spUncompressed.Signer = [20]byte{0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11}

	serialized, err = spUncompressed.Serialize()
	assert.NoError(t, err, "Failed to serialize uncompressed P2PKH spending condition")

	deserialized = SpendingCondition{}
	_, err = deserialized.Deserialize(serialized)
	assert.NoError(t, err, "Failed to deserialize uncompressed P2PKH spending condition")

	assert.Equal(t, spUncompressed, deserialized, "Uncompressed P2PKH spending condition mismatch")
}

func TestSingleSigP2WPKHSpendingCondition(t *testing.T) {
	sp := createSingleSigSpendingCondition(stacks.AddressHashModeSerializeP2WPKH, 345, 567)
	sp.KeyEncoding = stacks.PubKeyEncodingCompressed
	sp.Signature = createMessageSignature("fe")
	sp.Signer = [20]byte{0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11}

	serialized, err := sp.Serialize()
	assert.NoError(t, err, "Failed to serialize P2WPKH spending condition")

	deserialized := SpendingCondition{}
	_, err = deserialized.Deserialize(serialized)
	assert.NoError(t, err, "Failed to deserialize P2WPKH spending condition")

	assert.Equal(t, sp, deserialized, "P2WPKH spending condition mismatch")
}

func TestInvalidSpendingConditions(t *testing.T) {
	// Test invalid hash mode
	invalidHashMode := []byte{
		0xff,                                                                                                                   // Invalid hash mode
		0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, // Signer
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0xc8, // Nonce
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x37, // Fee
		0x00, // Key encoding
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, // Signature
	}

	sp := SpendingCondition{}
	_, err := sp.Deserialize(invalidHashMode)
	assert.Error(t, err, "Expected error for invalid hash mode")

	// Test incompatible hash mode and key encoding (P2WPKH with uncompressed key)
	sp = createSingleSigSpendingCondition(stacks.AddressHashModeSerializeP2WPKH, 123, 567)
	sp.KeyEncoding = stacks.PubKeyEncodingUncompressed
	sp.Signature = createMessageSignature("ff")
	sp.Signer = [20]byte{0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11}

	serialized, _ := sp.Serialize()
	_, err = sp.Deserialize(serialized)
	assert.Error(t, err, "Expected error for incompatible hash mode and key encoding")
}

func createSingleSigSpendingCondition(hashMode stacks.AddressHashMode, nonce, fee uint64) SpendingCondition {
	return SpendingCondition{
		HashMode: hashMode,
		Signer:   [20]byte{},
		Nonce:    nonce,
		Fee:      fee,
	}
}

func createMessageSignature(hexString string) [65]byte {
	var signature [65]byte
	sigBytes, _ := hex.DecodeString(hexString)
	copy(signature[:], sigBytes)
	return signature
}

func emptyMessageSignature() [65]byte {
	return [65]byte{}
}
