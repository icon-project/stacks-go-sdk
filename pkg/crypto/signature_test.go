package crypto

import (
	"encoding/hex"
	"errors"
	"testing"

	"github.com/icon-project/stacks-go-sdk/pkg/stacks"
)

func privateKeyToBytes(privateKey interface{}) ([]byte, error) {
	var privateKeyBuffer []byte
	switch v := privateKey.(type) {
	case string:
		var err error
		privateKeyBuffer, err = hex.DecodeString(v)
		if err != nil {
			return nil, err
		}
	case []byte:
		privateKeyBuffer = v
	default:
		return nil, errors.New("privateKey must be a string or []byte")
	}

	if len(privateKeyBuffer) != 32 && len(privateKeyBuffer) != 33 {
		return nil, errors.New("improperly formatted private-key. Private-key byte length should be 32 or 33")
	}

	if len(privateKeyBuffer) == 33 && privateKeyBuffer[32] != 1 {
		return nil, errors.New("improperly formatted private-key. 33 bytes indicate compressed key, but the last byte must be == 01")
	}

	return privateKeyBuffer, nil
}

func createStacksPrivateKey(key interface{}) (StacksPrivateKey, error) {
	data, err := privateKeyToBytes(key)
	if err != nil {
		return StacksPrivateKey{}, err
	}
	compressed := len(data) == stacks.CompressedPubkeyLengthBytes
	return StacksPrivateKey{Data: data, Compressed: compressed}, nil
}

func TestSignWithKey(t *testing.T) {
	privateKey, err := createStacksPrivateKey("bcf62fdd286f9b30b2c289cce3189dbf3b502dcd955b2dc4f67d18d77f3e73c7")
	if err != nil {
		t.Fatalf("Failed to create private key: %v", err)
	}

	publicKey := GetPublicKeyFromPrivate(privateKey.Data)

	hash := CalculateSighash([]byte("Hello World"))
	messageHash := hex.EncodeToString(hash[:])

	signature, err := SignWithKey(privateKey.Data, messageHash)
	if err != nil {
		t.Fatalf("Failed to sign: %v", err)
	}

	isValid, err := VerifySignature(messageHash, signature, publicKey)
	if err != nil {
		t.Fatalf("Error verifying signature: %v", err)
	}

	if !isValid {
		t.Fatalf("Signature verification failed: expected valid signature")
	}

	incorrectMessageHash := "0000000000000000000000000000000000000000000000000000000000000000"
	isValid, err = VerifySignature(incorrectMessageHash, signature, publicKey)
	if err != nil {
		t.Fatalf("Signature verification failed: %v", err)
	}

	if isValid {
		t.Errorf("Signature verification failed: expected invalid signature for incorrect message hash")
	}

	incorrectPublicKey := make([]byte, len(publicKey))
	copy(incorrectPublicKey, publicKey)
	incorrectPublicKey[0] ^= 0xFF // Flip bits in the first byte

	isValid, _ = VerifySignature(messageHash, signature, incorrectPublicKey)

	if isValid {
		t.Errorf("Signature verification failed: expected invalid signature for incorrect public key")
	}
}

func TestCalculatePresignSighash(t *testing.T) {
	curSigHash := "2be5719a864803128bc90aba66a2c42a28e07431fb3c7362f43dadf3c6be6cb5"
	authType := stacks.AuthTypeStandard
	fee := uint64(1)
	nonce := uint64(1)

	expectedHash := "9047149584c3e1af556484afc14dd599351c04cf5caca37c6e4d438490cead7b"

	curSigHashBytes, _ := hex.DecodeString(curSigHash)
	result := CalculatePresignSighash(curSigHashBytes, authType, fee, nonce)

	if hex.EncodeToString(result) != expectedHash {
		t.Errorf("calculatePresignSighash mismatch. Got %s, want %s", hex.EncodeToString(result), expectedHash)
	}
}

func TestCalculatePresignSighash2(t *testing.T) {
	curSigHash := "9af2794c87ae019025001231dabef4417ca2a5ca1ba2285ef4fa917195ad35dc"
	authType := stacks.AuthTypeStandard
	fee := uint64(180)
	nonce := uint64(1)

	expectedHash := "5e37d024804806b0c7657731541b4a2c8c5dc69f66e48223d8883634ddaa980c"

	curSigHashBytes, _ := hex.DecodeString(curSigHash)
	result := CalculatePresignSighash(curSigHashBytes, authType, fee, nonce)

	if hex.EncodeToString(result) != expectedHash {
		t.Errorf("calculatePresignSighash mismatch. Got %s, want %s", hex.EncodeToString(result), expectedHash)
	}
}

func TestSignWithKey2(t *testing.T) {
	privateKey := "c1d5bb638aa70862621667f9997711fce692cad782694103f8d9561f62e9f19701"
	messageHash := "9047149584c3e1af556484afc14dd599351c04cf5caca37c6e4d438490cead7b"

	expectedSignature := "001ca86b18448768a0eacc41ed749eceeb2c3073e424e9fc36b8e86481d92ad2244ac572e30782158b5a1d954a0f50ed3f6f02d1ff9c9dd6efe6d14c52878c52da"

	privateKeyBytes, _ := hex.DecodeString(privateKey)
	signature, err := SignWithKey(privateKeyBytes, messageHash)
	if err != nil {
		t.Fatalf("SignWithKey failed: %v", err)
	}

	if signature.Data != expectedSignature {
		t.Fatalf("SignWithKey signature mismatch. Got %s, want %s", signature.Data, expectedSignature)
	}
}

func TestCalculateStacksAddress(t *testing.T) {
	tests := []struct {
		name            string
		publicKeyHex    string
		network         stacks.ChainID
		expectedAddress string
	}{
		{
			name:            "Testnet Address",
			publicKeyHex:    "0332fc778e5beb5f944c75b2b63c21dd12c40bdcdf99ba0663168ae0b2be880aef",
			network:         stacks.ChainIDTestnet,
			expectedAddress: "ST15C893XJFJ6FSKM020P9JQDB5T7X6MQTXMBPAVH",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			publicKeyBytes, err := hex.DecodeString(tt.publicKeyHex)
			if err != nil {
				t.Fatalf("Failed to decode public key hex: %v", err)
			}

			address, err := CalculateStacksAddress(publicKeyBytes, tt.network)
			if err != nil {
				t.Fatalf("CalculateStacksAddress failed: %v", err)
			}

			if address != tt.expectedAddress {
				t.Errorf("CalculateStacksAddress mismatch.\nGot:  %s\nWant: %s", address, tt.expectedAddress)
			}
		})
	}
}

func TestGetAddressFromPrivateKey(t *testing.T) {
	tests := []struct {
		name            string
		privateKeyHex   string
		network         stacks.ChainID
		expectedAddress string
	}{
		{
			name:            "Mainnet Address from Private Key",
			privateKeyHex:   "c1d5bb638aa70862621667f9997711fce692cad782694103f8d9561f62e9f19701",
			network:         stacks.ChainIDTestnet,
			expectedAddress: "ST15C893XJFJ6FSKM020P9JQDB5T7X6MQTXMBPAVH",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			privateKeyBytes, err := hex.DecodeString(tt.privateKeyHex)
			if err != nil {
				t.Fatalf("Failed to decode private key hex: %v", err)
			}

			address, err := GetAddressFromPrivateKey(privateKeyBytes, tt.network)
			if err != nil {
				t.Fatalf("GetAddressFromPrivateKey failed: %v", err)
			}

			if address != tt.expectedAddress {
				t.Errorf("GetAddressFromPrivateKey mismatch.\nGot:  %s\nWant: %s", address, tt.expectedAddress)
			}
		})
	}
}
