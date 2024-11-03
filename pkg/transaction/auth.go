package transaction

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/icon-project/stacks-go-sdk/internal/utils"
	"github.com/icon-project/stacks-go-sdk/pkg/crypto"
	"github.com/icon-project/stacks-go-sdk/pkg/stacks"
)

type TransactionAuth struct {
	AuthType    stacks.AuthType
	OriginAuth  SpendingCondition
	SponsorAuth *SpendingCondition
}

type SpendingCondition struct {
	HashMode    stacks.AddressHashMode
	Signer      [20]byte
	Nonce       uint64
	Fee         uint64
	KeyEncoding stacks.PubKeyEncoding
	Signature   [stacks.RecoverableECDSASigLengthBytes]byte
}

func (t *TokenTransferTransaction) GetAuth() *TransactionAuth {
	return &t.Auth
}

func (t *SmartContractTransaction) GetAuth() *TransactionAuth {
	return &t.Auth
}

func (t *ContractCallTransaction) GetAuth() *TransactionAuth {
	return &t.Auth
}

func (t *TokenTransferTransaction) Sign(privateKey []byte) error {
	return SignTransaction(t, privateKey)
}

func (t *TokenTransferTransaction) Verify(publicKey []byte) (bool, error) {
	return VerifyTransaction(t, publicKey)
}

func (t *ContractCallTransaction) Sign(privateKey []byte) error {
	return SignTransaction(t, privateKey)
}

func (t *ContractCallTransaction) Verify(publicKey []byte) (bool, error) {
	return VerifyTransaction(t, publicKey)
}

func (auth *TransactionAuth) Serialize() ([]byte, error) {
	buf := make([]byte, 0, 256)

	buf = append(buf, byte(auth.AuthType))

	originAuthBytes, err := auth.OriginAuth.Serialize()
	if err != nil {
		return nil, err
	}
	buf = append(buf, originAuthBytes...)

	if auth.AuthType == stacks.AuthTypeSponsored {
		if auth.SponsorAuth == nil {
			return nil, errors.New("sponsor auth is required for sponsored transactions")
		}
		sponsorAuthBytes, err := auth.SponsorAuth.Serialize()
		if err != nil {
			return nil, err
		}
		buf = append(buf, sponsorAuthBytes...)
	}

	return buf, nil
}

func (auth *TransactionAuth) Deserialize(data []byte) (int, error) {
	if len(data) < 1 {
		return 0, errors.New("invalid auth data length")
	}

	auth.AuthType = stacks.AuthType(data[0])
	offset := 1

	originAuthLen, err := auth.OriginAuth.Deserialize(data[offset:])
	if err != nil {
		return 0, err
	}
	offset += originAuthLen

	if auth.AuthType == stacks.AuthTypeSponsored {
		auth.SponsorAuth = &SpendingCondition{}
		sponsorAuthLen, err := auth.SponsorAuth.Deserialize(data[offset:])
		if err != nil {
			return 0, err
		}
		offset += sponsorAuthLen
	}

	return offset, nil
}

func (sc *SpendingCondition) Serialize() ([]byte, error) {
	buf := make([]byte, 0, 103) // 1 + 20 + 8 + 8 + 1 + 65

	buf = append(buf, byte(sc.HashMode))
	buf = append(buf, sc.Signer[:]...)

	nonceBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(nonceBytes, sc.Nonce)
	buf = append(buf, nonceBytes...)

	feeBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(feeBytes, sc.Fee)
	buf = append(buf, feeBytes...)

	buf = append(buf, byte(sc.KeyEncoding))
	buf = append(buf, sc.Signature[:]...)

	return buf, nil
}

func (sc *SpendingCondition) Deserialize(data []byte) (int, error) {
	if len(data) < 103 {
		return 0, errors.New("invalid spending condition data length")
	}

	hashMode := stacks.AddressHashMode(data[0])
	if !utils.IsValidAddressHashMode(hashMode) {
		return 0, fmt.Errorf("invalid address hash mode: %d", hashMode)
	}
	sc.HashMode = hashMode

	copy(sc.Signer[:], data[1:21])
	sc.Nonce = binary.BigEndian.Uint64(data[21:29])
	sc.Fee = binary.BigEndian.Uint64(data[29:37])

	keyEncoding := stacks.PubKeyEncoding(data[37])
	if !utils.IsValidPubKeyEncoding(keyEncoding) {
		return 0, fmt.Errorf("invalid public key encoding: %d", keyEncoding)
	}
	sc.KeyEncoding = keyEncoding

	if !utils.IsCompatibleHashModeAndKeyEncoding(sc.HashMode, sc.KeyEncoding) {
		return 0, fmt.Errorf("incompatible hash mode (%d) and key encoding (%d)", sc.HashMode, sc.KeyEncoding)
	}

	copy(sc.Signature[:], data[38:103])

	return 103, nil
}

func SignTransaction(tx StacksTransaction, privateKey []byte) error {
	txCopy := tx.Clone()

	// 1. Set the fee and nonce to 0, and set the signature bytes to 0
	authCopy := txCopy.GetAuth()

	authCopy.OriginAuth.Signature = [65]byte{}
	authCopy.OriginAuth.Nonce = 0
	authCopy.OriginAuth.Fee = 0

	// 2. Serialize the transaction
	serializedTx, err := txCopy.Serialize()
	if err != nil {
		return err
	}

	// 3. Calculate the initial sighash
	sighash := crypto.CalculateSighash(serializedTx)

	// 4. Calculate the presign-sighash
	originalAuth := tx.GetAuth()
	presignSighash := crypto.CalculatePresignSighash(sighash, originalAuth.AuthType, originalAuth.OriginAuth.Fee, originalAuth.OriginAuth.Nonce)

	// 5. Sign the presign-sighash
	signature, err := crypto.SignWithKey(privateKey, hex.EncodeToString(presignSighash))
	if err != nil {
		return err
	}

	// 6. Set the signature in the spending condition
	signatureBytes, _ := hex.DecodeString(signature.Data)
	copy(originalAuth.OriginAuth.Signature[:], signatureBytes)

	return nil
}

func VerifyTransaction(tx StacksTransaction, publicKey []byte) (bool, error) {
	txCopy := tx.Clone()

	// 1. Extract the signature
	authCopy := txCopy.GetAuth()
	signature := authCopy.OriginAuth.Signature

	// 2. Clear the signature in the spending condition
	authCopy.OriginAuth.Signature = [stacks.RecoverableECDSASigLengthBytes]byte{}
	authCopy.OriginAuth.Nonce = 0
	authCopy.OriginAuth.Fee = 0

	// 3. Serialize the transaction
	serializedTx, err := txCopy.Serialize()
	if err != nil {
		return false, err
	}

	// 4. Calculate the initial sighash
	sighash := crypto.CalculateSighash(serializedTx)

	// 5. Calculate the presign-sighash
	originalAuth := tx.GetAuth()
	presignSighash := crypto.CalculatePresignSighash(sighash, originalAuth.AuthType, originalAuth.OriginAuth.Fee, originalAuth.OriginAuth.Nonce)

	// 6. Verify the signature
	messageSignature := crypto.MessageSignature{
		Type: crypto.MessageSignatureType,
		Data: hex.EncodeToString(signature[:]),
	}
	return crypto.VerifySignature(hex.EncodeToString(presignSighash), messageSignature, publicKey)
}
