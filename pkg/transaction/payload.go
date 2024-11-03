package transaction

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"unicode"

	"github.com/icon-project/stacks-go-sdk/internal/utils"
	"github.com/icon-project/stacks-go-sdk/pkg/clarity"
	"github.com/icon-project/stacks-go-sdk/pkg/stacks"
)

type Payload interface {
	Serialize() ([]byte, error)
	Deserialize([]byte) (int, error)
}

type TokenTransferPayload struct {
	Recipient clarity.ClarityValue // Can be either StandardPrincipal or ContractPrincipal
	Amount    uint64
	Memo      string
}

type SmartContractPayload struct {
	ContractName string
	CodeBody     string
}

type ContractCallPayload struct {
	ContractAddress clarity.ClarityValue // Can be either StandardPrincipal or ContractPrincipal
	ContractName    string
	FunctionName    string
	FunctionArgs    []clarity.ClarityValue
}

func (t *TokenTransferTransaction) GetPayload() Payload {
	return &t.Payload
}

func (t *SmartContractTransaction) GetPayload() Payload {
	return &t.Payload
}

func (t *ContractCallTransaction) GetPayload() Payload {
	return &t.Payload
}

func NewTokenTransferPayload(recipient string, amount uint64, memo string) (*TokenTransferPayload, error) {
	principalCV, err := clarity.StringToPrincipal(recipient)
	if err != nil {
		return nil, err
	}

	return &TokenTransferPayload{
		Recipient: principalCV,
		Amount:    amount,
		Memo:      memo,
	}, nil
}

func NewSmartContractPayload(contractName string, codeBody string) (*SmartContractPayload, error) {
	if err := validateContractName(contractName); err != nil {
		return nil, err
	}

	if err := validateCodeBody(codeBody); err != nil {
		return nil, err
	}

	return &SmartContractPayload{
		ContractName: contractName,
		CodeBody:     codeBody,
	}, nil
}

func NewContractCallPayload(contractAddress string, contractName string, functionName string, functionArgs []clarity.ClarityValue) (*ContractCallPayload, error) {
	principalCV, err := clarity.StringToPrincipal(contractAddress)
	if err != nil {
		return nil, err
	}

	return &ContractCallPayload{
		ContractAddress: principalCV,
		ContractName:    contractName,
		FunctionName:    functionName,
		FunctionArgs:    functionArgs,
	}, nil
}

func (p *TokenTransferPayload) Serialize() ([]byte, error) {
	buf := make([]byte, 0, 128)

	buf = append(buf, byte(stacks.PayloadTypeTokenTransfer))

	recipientBytes, err := p.Recipient.Serialize()
	if err != nil {
		return nil, err
	}
	buf = append(buf, recipientBytes...)

	amountBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(amountBytes, p.Amount)
	buf = append(buf, amountBytes...)

	memoBytes := make([]byte, stacks.MemoMaxLengthBytes)
	copy(memoBytes, p.Memo)
	buf = append(buf, memoBytes...)

	return buf, nil
}

func (p *TokenTransferPayload) Deserialize(data []byte) (int, error) {
	if len(data) < 1 || stacks.PayloadType(data[0]) != stacks.PayloadTypeTokenTransfer {
		return 0, errors.New("invalid token transfer payload")
	}

	offset := 1

	recipient, n, err := clarity.DeserializePrincipal(data[offset:])
	if err != nil {
		return 0, err
	}
	p.Recipient = recipient
	offset += n

	if len(data[offset:]) < 8 {
		return 0, errors.New("insufficient data for amount")
	}
	p.Amount = binary.BigEndian.Uint64(data[offset : offset+8])
	offset += 8

	if len(data[offset:]) < stacks.MemoMaxLengthBytes {
		return 0, errors.New("insufficient data for memo")
	}
	p.Memo = string(bytes.TrimRight(data[offset:offset+stacks.MemoMaxLengthBytes], "\x00"))
	offset += stacks.MemoMaxLengthBytes

	return offset, nil
}

func (p *SmartContractPayload) Serialize() ([]byte, error) {
	buf := make([]byte, 0, len(p.ContractName)+len(p.CodeBody)+6) // 1 + 1 + 4 bytes for headers

	// Payload type
	buf = append(buf, byte(stacks.PayloadTypeSmartContract))

	// Contract name
	if len(p.ContractName) > stacks.MaxStringLengthBytes {
		return nil, fmt.Errorf("contract name exceeds maximum length of %d", stacks.MaxStringLengthBytes)
	}
	buf = append(buf, byte(len(p.ContractName)))
	buf = append(buf, []byte(p.ContractName)...)

	// Code body with 4-byte length prefix
	codeBodyLen := make([]byte, 4)
	binary.BigEndian.PutUint32(codeBodyLen, uint32(len(p.CodeBody)))
	buf = append(buf, codeBodyLen...)
	buf = append(buf, []byte(p.CodeBody)...)

	return buf, nil
}

func (p *SmartContractPayload) Deserialize(data []byte) (int, error) {
	if len(data) < 2 || stacks.PayloadType(data[0]) != stacks.PayloadTypeSmartContract {
		return 0, errors.New("invalid smart contract payload")
	}

	offset := 1

	// Contract name
	nameLen := int(data[offset])
	offset++
	if nameLen > stacks.MaxStringLengthBytes {
		return 0, fmt.Errorf("contract name length %d exceeds maximum %d", nameLen, stacks.MaxStringLengthBytes)
	}
	if len(data[offset:]) < nameLen {
		return 0, errors.New("insufficient data for contract name")
	}
	p.ContractName = string(data[offset : offset+nameLen])
	offset += nameLen

	// Code body
	if len(data[offset:]) < 4 {
		return 0, errors.New("insufficient data for code body length")
	}
	codeBodyLen := binary.BigEndian.Uint32(data[offset : offset+4])
	offset += 4

	if len(data[offset:]) < int(codeBodyLen) {
		return 0, errors.New("insufficient data for code body")
	}
	p.CodeBody = string(data[offset : offset+int(codeBodyLen)])
	offset += int(codeBodyLen)

	return offset, nil
}

func (p *ContractCallPayload) Serialize() ([]byte, error) {
	buf := make([]byte, 0, 128)

	buf = append(buf, byte(stacks.PayloadTypeContractCall))

	contractAddressBytes, err := p.ContractAddress.Serialize()
	if err != nil {
		return nil, err
	}
	buf = append(buf, contractAddressBytes[1:]...) // skip clarity type byte

	contractNameBytes, err := utils.SerializeString(p.ContractName, stacks.MaxStringLengthBytes)
	if err != nil {
		return nil, err
	}
	buf = append(buf, contractNameBytes...)

	functionNameBytes, err := utils.SerializeString(p.FunctionName, stacks.MaxStringLengthBytes)
	if err != nil {
		return nil, err
	}
	buf = append(buf, functionNameBytes...)

	argCountBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(argCountBytes, uint32(len(p.FunctionArgs)))
	buf = append(buf, argCountBytes...)

	for _, arg := range p.FunctionArgs {
		argBytes, err := arg.Serialize()
		if err != nil {
			return nil, err
		}
		buf = append(buf, argBytes...)
	}

	return buf, nil
}

func (p *ContractCallPayload) Deserialize(data []byte) (int, error) {
	if len(data) < 1 || stacks.PayloadType(data[0]) != stacks.PayloadTypeContractCall {
		return 0, errors.New("invalid contract call payload")
	}

	offset := 1

	// add back in clarity type byte. note this is ClarityTypeStandardPrincipal but should be ClarityTypeContractPrincipal
	dataWithClarityType := append([]byte{byte(clarity.ClarityTypeStandardPrincipal)}, data[offset:]...)

	contractPrincipal, n, err := clarity.DeserializePrincipal(dataWithClarityType)
	if err != nil {
		return 0, err
	}
	p.ContractAddress = contractPrincipal
	offset += n - 1

	contractName, n, err := utils.DeserializeString(data[offset:], stacks.MaxStringLengthBytes)
	if err != nil {
		return 0, err
	}
	p.ContractName = contractName
	offset += n

	functionName, n, err := utils.DeserializeString(dataWithClarityType[offset:], stacks.MaxStringLengthBytes)
	if err != nil {
		return 0, err
	}
	p.FunctionName = functionName
	offset += n

	if len(data[offset:]) < 4 {
		return 0, errors.New("insufficient data for function args count")
	}
	argCount := binary.BigEndian.Uint32(data[offset : offset+4])
	offset += 4

	p.FunctionArgs = make([]clarity.ClarityValue, argCount)
	for i := uint32(0); i < argCount; i++ {
		arg, err := clarity.DeserializeClarityValue(data[offset:])
		if err != nil {
			return 0, err
		}
		p.FunctionArgs[i] = arg
		serialized, _ := arg.Serialize()
		offset += len(serialized)
	}

	return offset, nil
}

func validateContractName(name string) error {
	if len(name) == 0 || len(name) > stacks.MaxStringLengthBytes {
		return fmt.Errorf("contract name must be between 1 and %d characters", stacks.MaxStringLengthBytes)
	}

	// First character must be a letter
	if !unicode.IsLetter(rune(name[0])) {
		return errors.New("contract name must start with a letter")
	}

	// Subsequent characters must be letters, numbers, hyphens, or underscores
	for _, r := range name[1:] {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != '-' && r != '_' {
			return errors.New("contract name can only contain letters, numbers, hyphens, and underscores")
		}
	}

	return nil
}

func validateCodeBody(code string) error {
	for i, r := range code {
		if r != '\n' && r != '\t' && (r < 0x20 || r > 0x7e) {
			return fmt.Errorf("invalid character in code body at position %d: %q", i, r)
		}
	}
	return nil
}
