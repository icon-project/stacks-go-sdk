package crypto

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"

	// nolint:staticcheck // SA1019: RIPEMD-160 is required for compatibility
	"golang.org/x/crypto/ripemd160"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
	"github.com/icon-project/stacks-go-sdk/pkg/stacks"
)

type MessageSignature struct {
	Type StacksMessageType
	Data string
}

type StacksMessageType int

const (
	Address StacksMessageType = iota
	MessageSignatureType
)

type StacksPrivateKey struct {
	Data       []byte
	Compressed bool
}

func SignWithKey(privateKey []byte, messageHash string) (MessageSignature, error) {
	privKey, _ := btcec.PrivKeyFromBytes(privateKey[:32])
	messageHashBytes, err := hex.DecodeString(messageHash)
	if err != nil {
		return MessageSignature{}, err
	}
	signature := ecdsa.SignCompact(privKey, messageHashBytes, true)

	recoveryID := signature[0] - 27 - 4 // Remove the base Bitcoin offset (27) and account for account for compressed key flag (4). For more information, see https://github.com/btcsuite/btcd/blob/master/btcec/ecdsa/signature.go#L233
	vrsSignature := fmt.Sprintf("%02x%s", recoveryID, hex.EncodeToString(signature[1:]))
	return CreateMessageSignature(vrsSignature)
}

func CreateMessageSignature(signature string) (MessageSignature, error) {
	signatureBytes, err := hex.DecodeString(signature)
	if err != nil {
		return MessageSignature{}, err
	}
	if len(signatureBytes) != stacks.RecoverableECDSASigLengthBytes {
		return MessageSignature{}, errors.New("invalid signature")
	}
	return MessageSignature{
		Type: MessageSignatureType,
		Data: signature,
	}, nil
}

func GetPublicKeyFromPrivate(privateKey []byte) []byte {
	_, pubKey := btcec.PrivKeyFromBytes(privateKey)
	return pubKey.SerializeCompressed()
}

func VerifySignature(messageHash string, signature MessageSignature, publicKey []byte) (bool, error) {
	messageHashBytes, err := hex.DecodeString(messageHash)
	if err != nil {
		return false, errors.New("invalid message hash")
	}

	signatureBytes, err := hex.DecodeString(signature.Data)
	if err != nil {
		return false, errors.New("invalid signature")
	}

	// The signature is in [R || S] format, where R and S are 32 bytes each
	if len(signatureBytes) != 65 {
		return false, errors.New("invalid signature length")
	}

	// Parse the public key
	pubKey, err := btcec.ParsePubKey(publicKey)
	if err != nil {
		return false, errors.New("failed to parse public key")
	}

	// Create a new ECDSA signature from R and S components
	r := new(btcec.ModNScalar)
	r.SetByteSlice(signatureBytes[1:33])
	s := new(btcec.ModNScalar)
	s.SetByteSlice(signatureBytes[33:])
	ecdsaSignature := ecdsa.NewSignature(r, s)

	// Verify the signature
	return ecdsaSignature.Verify(messageHashBytes, pubKey), nil
}

func Hash160(b []byte) []byte {
	h := sha256.Sum256(b)
	ripemd160Hasher := ripemd160.New()
	ripemd160Hasher.Write(h[:])
	return ripemd160Hasher.Sum(nil)
}

func CalculateSighash(serializedTx []byte) []byte {
	h := sha512.New512_256()
	h.Write(serializedTx)
	return h.Sum(nil)
}

func CalculatePresignSighash(sighash []byte, authType stacks.AuthType, fee uint64, nonce uint64) []byte {
	data := make([]byte, 0, len(sighash)+1+8+8)
	data = append(data, sighash...)
	data = append(data, byte(authType))
	feeBytes := make([]byte, 8)
	nonceBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(feeBytes, fee)
	binary.BigEndian.PutUint64(nonceBytes, nonce)
	data = append(data, feeBytes...)
	data = append(data, nonceBytes...)

	h := sha512.New512_256()
	h.Write(data)
	return h.Sum(nil)
}
