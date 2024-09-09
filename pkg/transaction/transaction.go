package transaction

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/icon-project/stacks-go-sdk/pkg/clarity"
	"github.com/icon-project/stacks-go-sdk/pkg/stacks"
)

type StacksTransaction interface {
	Serialize() ([]byte, error)
	Deserialize([]byte) error
	GetAuth() *TransactionAuth
	GetPayload() Payload
	Clone() StacksTransaction
}

type BaseTransaction struct {
	Version           stacks.TransactionVersion
	ChainID           stacks.ChainID
	Auth              TransactionAuth
	AnchorMode        stacks.AnchorMode
	PostConditionMode stacks.PostConditionMode
	PostConditions    []PostCondition
}

type TokenTransferTransaction struct {
	BaseTransaction
	Payload TokenTransferPayload
}

type ContractCallTransaction struct {
	BaseTransaction
	Payload ContractCallPayload
}

func (t *TokenTransferTransaction) Clone() StacksTransaction {
	clone := *t
	return &clone
}

func (t *ContractCallTransaction) Clone() StacksTransaction {
	clone := *t
	return &clone
}

func NewTokenTransferTransaction(
	recipient string,
	amount uint64,
	memo string,
	version stacks.TransactionVersion,
	chainID stacks.ChainID,
	signer [20]byte,
	nonce uint64,
	fee uint64,
	anchorMode stacks.AnchorMode,
	postConditionMode stacks.PostConditionMode,
) (*TokenTransferTransaction, error) {
	payload, err := NewTokenTransferPayload(recipient, amount, memo)
	if err != nil {
		return nil, err
	}
	return &TokenTransferTransaction{
		BaseTransaction: BaseTransaction{
			Version: version,
			ChainID: chainID,
			Auth: TransactionAuth{
				AuthType: stacks.AuthTypeStandard,
				OriginAuth: SpendingCondition{
					HashMode:    stacks.AddressHashModeSerializeP2PKH,
					Signer:      signer,
					Nonce:       nonce,
					Fee:         fee,
					KeyEncoding: stacks.PubKeyEncodingCompressed,
					Signature:   [65]byte{}, // This will be filled when signing the transaction
				},
			},
			AnchorMode:        anchorMode,
			PostConditionMode: postConditionMode,
			PostConditions:    []PostCondition{}, // Empty post condition
		},
		Payload: *payload,
	}, nil
}

func NewContractCallTransaction(
	contractAddress string,
	contractName string,
	functionName string,
	functionArgs []clarity.ClarityValue,
	version stacks.TransactionVersion,
	chainID stacks.ChainID,
	signer [20]byte,
	nonce uint64,
	fee uint64,
	anchorMode stacks.AnchorMode,
	postConditionMode stacks.PostConditionMode) (*ContractCallTransaction, error) {
	payload, err := NewContractCallPayload(contractAddress, contractName, functionName, functionArgs)
	if err != nil {
		return nil, err
	}

	return &ContractCallTransaction{
		BaseTransaction: BaseTransaction{
			Version: version,
			ChainID: chainID,
			Auth: TransactionAuth{
				AuthType: stacks.AuthTypeStandard,
				OriginAuth: SpendingCondition{
					HashMode:    stacks.AddressHashModeSerializeP2PKH,
					Signer:      signer,
					Nonce:       nonce,
					Fee:         fee,
					KeyEncoding: stacks.PubKeyEncodingCompressed,
					Signature:   [65]byte{}, // This will be filled when signing the transaction
				},
			},
			AnchorMode:        anchorMode,
			PostConditionMode: postConditionMode,
			PostConditions:    []PostCondition{}, // Empty post condition
		},
		Payload: *payload,
	}, nil
}

func (t *TokenTransferTransaction) Serialize() ([]byte, error) {
	buf := make([]byte, 0, 128)

	buf = append(buf, byte(t.Version))
	chainIDBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(chainIDBytes, uint32(t.ChainID))
	buf = append(buf, chainIDBytes...)

	authBytes, err := t.Auth.Serialize()
	if err != nil {
		return nil, err
	}
	buf = append(buf, authBytes...)

	buf = append(buf, byte(t.AnchorMode))
	buf = append(buf, byte(t.PostConditionMode))

	postConditionsBytes := SerializePostConditions(t.PostConditions)
	buf = append(buf, postConditionsBytes...)

	payloadBytes, err := t.Payload.Serialize()
	if err != nil {
		return nil, err
	}
	buf = append(buf, payloadBytes...)

	return buf, nil
}

func (t *TokenTransferTransaction) Deserialize(data []byte) error {
	if len(data) < 7 {
		return errors.New("insufficient data for transaction")
	}

	t.Version = stacks.TransactionVersion(data[0])
	t.ChainID = stacks.ChainID(binary.BigEndian.Uint32(data[1:5]))
	authType := stacks.AuthType(data[5])

	offset := 5 // Start after version and chain ID

	auth := TransactionAuth{AuthType: authType}
	authLen, err := auth.Deserialize(data[offset:])
	if err != nil {
		return fmt.Errorf("failed to deserialize auth: %w", err)
	}
	offset += authLen
	t.Auth = auth

	if len(data) < offset+1 {
		return errors.New("insufficient data for anchor mode")
	}
	t.AnchorMode = stacks.AnchorMode(data[offset])
	offset++

	if len(data) < offset+1 {
		return errors.New("insufficient data for post condition mode")
	}
	t.PostConditionMode = stacks.PostConditionMode(data[offset])
	offset++

	postConditions, postConditionsLen, err := DeserializePostConditions(data[offset:])
	if err != nil {
		return fmt.Errorf("failed to deserialize post conditions: %w", err)
	}
	t.PostConditions = postConditions
	offset += postConditionsLen

	if len(data) < offset+1 {
		return errors.New("insufficient data for payload type")
	}
	if stacks.PayloadType(data[offset]) != stacks.PayloadTypeTokenTransfer {
		return errors.New("invalid payload type for token transfer transaction")
	}
	offset++

	_, err = t.Payload.Deserialize(data[offset:])
	if err != nil {
		return fmt.Errorf("failed to deserialize token transfer payload: %w", err)
	}

	return nil
}

func (t *ContractCallTransaction) Serialize() ([]byte, error) {
	buf := make([]byte, 0, 128)

	buf = append(buf, byte(t.Version))

	chainIDBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(chainIDBytes, uint32(t.ChainID))
	buf = append(buf, chainIDBytes...)

	authBytes, err := t.Auth.Serialize()
	if err != nil {
		return nil, fmt.Errorf("failed to serialize auth: %w", err)
	}
	buf = append(buf, authBytes...)

	buf = append(buf, byte(t.AnchorMode))

	buf = append(buf, byte(t.PostConditionMode))

	postConditionsBytes := SerializePostConditions(t.PostConditions)
	buf = append(buf, postConditionsBytes...)

	payloadBytes, err := t.Payload.Serialize()
	if err != nil {
		return nil, fmt.Errorf("failed to serialize contract call payload: %w", err)
	}
	buf = append(buf, payloadBytes...)

	return buf, nil
}

func (t *ContractCallTransaction) Deserialize(data []byte) error {
	if len(data) < 7 {
		return errors.New("insufficient data for transaction")
	}

	t.Version = stacks.TransactionVersion(data[0])
	t.ChainID = stacks.ChainID(binary.BigEndian.Uint32(data[1:5]))
	authType := stacks.AuthType(data[5])

	offset := 5 // Start after version and chain ID

	auth := TransactionAuth{AuthType: authType}
	authLen, err := auth.Deserialize(data[offset:])
	if err != nil {
		return fmt.Errorf("failed to deserialize auth: %w", err)
	}
	offset += authLen
	t.Auth = auth

	if len(data) < offset+1 {
		return errors.New("insufficient data for anchor mode")
	}
	t.AnchorMode = stacks.AnchorMode(data[offset])
	offset++

	if len(data) < offset+1 {
		return errors.New("insufficient data for post condition mode")
	}
	t.PostConditionMode = stacks.PostConditionMode(data[offset])
	offset++

	postConditions, postConditionsLen, err := DeserializePostConditions(data[offset:])
	if err != nil {
		return fmt.Errorf("failed to deserialize post conditions: %w", err)
	}
	t.PostConditions = postConditions
	offset += postConditionsLen

	if len(data) < offset+1 {
		return errors.New("insufficient data for payload type")
	}
	if stacks.PayloadType(data[offset]) != stacks.PayloadTypeContractCall {
		return errors.New("invalid payload type for contract call transaction")
	}
	offset++

	_, err = t.Payload.Deserialize(data[offset:])
	if err != nil {
		return fmt.Errorf("failed to deserialize contract call payload: %w", err)
	}

	return nil
}

func DeserializeTransaction(data []byte) (StacksTransaction, error) {
	if len(data) < 7 { // Minimum length for version, chain ID, and auth type
		return nil, errors.New("insufficient data for transaction")
	}

	version := stacks.TransactionVersion(data[0])
	chainID := stacks.ChainID(binary.BigEndian.Uint32(data[1:5]))

	offset := 5 // Start after version and chain ID

	baseTx := BaseTransaction{
		Version: version,
		ChainID: chainID,
	}

	authLen, err := baseTx.Auth.Deserialize(data[offset:])
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize auth: %w", err)
	}
	offset += authLen

	if len(data) < offset+1 {
		return nil, errors.New("insufficient data for anchor mode")
	}
	baseTx.AnchorMode = stacks.AnchorMode(data[offset])
	offset++

	if len(data) < offset+1 {
		return nil, errors.New("insufficient data for post condition mode")
	}
	baseTx.PostConditionMode = stacks.PostConditionMode(data[offset])
	offset++

	postConditions, postConditionsLen, err := DeserializePostConditions(data[offset:])
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize post conditions: %w", err)
	}
	baseTx.PostConditions = postConditions
	offset += postConditionsLen

	if len(data) < offset+1 {
		return nil, errors.New("insufficient data for payload type")
	}
	payloadType := stacks.PayloadType(data[offset])

	var tx StacksTransaction

	switch payloadType {
	case stacks.PayloadTypeTokenTransfer:
		tokenTx := &TokenTransferTransaction{BaseTransaction: baseTx}
		_, err := tokenTx.Payload.Deserialize(data[offset:])
		if err != nil {
			return nil, fmt.Errorf("failed to deserialize token transfer payload: %w", err)
		}
		tx = tokenTx
	case stacks.PayloadTypeContractCall:
		contractTx := &ContractCallTransaction{BaseTransaction: baseTx}
		_, err := contractTx.Payload.Deserialize(data[offset:])
		if err != nil {
			return nil, fmt.Errorf("failed to deserialize contract call payload: %w", err)
		}
		tx = contractTx
	default:
		return nil, fmt.Errorf("unsupported payload type: %d", payloadType)
	}

	return tx, nil
}
