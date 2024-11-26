package transaction

import (
	"encoding/binary"
	"errors"

	"github.com/icon-project/stacks-go-sdk/pkg/clarity"
	"github.com/icon-project/stacks-go-sdk/pkg/stacks"
)

type PostCondition interface {
	Serialize() ([]byte, error)
	Deserialize([]byte) (int, error)
	GetType() stacks.PostConditionType
}

type AssetInfo struct {
	Contract  *clarity.ContractPrincipal
	AssetName string
}

func (a *AssetInfo) Serialize() ([]byte, error) {
	contractBytes, err := a.Contract.Serialize()
	if err != nil {
		return nil, err
	}
	contractBytes = contractBytes[1:]
	buf := append(contractBytes, byte(len(a.AssetName)))
	return append(buf, []byte(a.AssetName)...), nil
}

func (a *AssetInfo) Deserialize(data []byte) (int, error) {
	if len(data) < 22 {
		return 0, errors.New("invalid asset info")
	}
	contractNameLength := int(data[21])
	if len(data) < 22+contractNameLength {
		return 0, errors.New("invalid contract principal name length")
	}
	contractName := string(data[22 : 22+contractNameLength])
	contract, err := clarity.NewContractPrincipal(data[0], [20]byte(data[1:21]), contractName)
	if err != nil {
		return 0, err
	}
	a.Contract = contract
	offset := 22 + contractNameLength
	assetNameLength := int(data[offset])
	offset++
	if len(data) < offset+assetNameLength {
		return 0, errors.New("invalid asset name length")
	}
	a.AssetName = string(data[offset : offset+assetNameLength])
	return offset + assetNameLength, nil
}

type PostConditionPrincipal interface {
	Serialize() ([]byte, error)
	Deserialize([]byte) (int, error)
	GetID() stacks.PostConditionPrincipalType
}

type PostConditionOriginPrincipal struct {
}

func (pc *PostConditionOriginPrincipal) GetID() stacks.PostConditionPrincipalType {
	return stacks.PostConditionPrincipalTypeOrigin
}

func (pc *PostConditionOriginPrincipal) Serialize() ([]byte, error) {
	return []byte{byte(stacks.PostConditionPrincipalTypeOrigin)}, nil
}

func (pc *PostConditionOriginPrincipal) Deserialize(data []byte) (int, error) {
	if len(data) < 1 || stacks.PostConditionPrincipalType(data[0]) != stacks.PostConditionPrincipalTypeOrigin {
		return 0, errors.New("invalid origin post condition principal")
	}
	return 1, nil
}

type PostConditionStandardPrincipal struct {
	Address *clarity.StandardPrincipal
}

func (pc *PostConditionStandardPrincipal) GetID() stacks.PostConditionPrincipalType {
	return stacks.PostConditionPrincipalTypeStandard
}

func (pc *PostConditionStandardPrincipal) Serialize() ([]byte, error) {
	addressBytes := make([]byte, 21)
	addressBytes[0] = byte(pc.Address.Version)
	copy(addressBytes[1:], pc.Address.Hash160[:])
	return append([]byte{byte(stacks.PostConditionPrincipalTypeStandard)}, addressBytes...), nil
}

func (pc *PostConditionStandardPrincipal) Deserialize(data []byte) (int, error) {
	if len(data) < 22 || stacks.PostConditionPrincipalType(data[0]) != stacks.PostConditionPrincipalTypeStandard {
		return 0, errors.New("invalid standard post condition principal")
	}
	address := clarity.NewStandardPrincipal(data[1], [20]byte(data[2:22]))
	pc.Address = address
	return 22, nil
}

type PostConditionContractPrincipal struct {
	Contract *clarity.ContractPrincipal
}

func (pc *PostConditionContractPrincipal) GetID() stacks.PostConditionPrincipalType {
	return stacks.PostConditionPrincipalTypeContract
}

func (pc *PostConditionContractPrincipal) Serialize() ([]byte, error) {
	contractBytes, err := pc.Contract.Serialize()
	if err != nil {
		return nil, err
	}
	return append([]byte{byte(stacks.PostConditionPrincipalTypeContract)}, contractBytes[1:]...), nil
}

func (pc *PostConditionContractPrincipal) Deserialize(data []byte) (int, error) {
	if len(data) < 23 || stacks.PostConditionPrincipalType(data[0]) != stacks.PostConditionPrincipalTypeContract {
		return 0, errors.New("invalid contract post condition principal")
	}
	contractNameLength := int(data[22])
	if len(data) < 23+contractNameLength {
		return 0, errors.New("invalid contract principal name length")
	}
	contractName := string(data[23 : 23+contractNameLength])
	contract, err := clarity.NewContractPrincipal(data[1], [20]byte(data[2:22]), contractName)
	if err != nil {
		return 0, err
	}
	pc.Contract = contract
	return 23 + contractNameLength, nil
}

func DeserializePostConditionPrincipal(data []byte) (PostConditionPrincipal, int, error) {
	if len(data) < 1 {
		return nil, 0, errors.New("insufficient data for post condition principal")
	}
	principalType := stacks.PostConditionPrincipalType(data[0])
	switch principalType {
	case stacks.PostConditionPrincipalTypeOrigin:
		principal := &PostConditionOriginPrincipal{}
		l, err := principal.Deserialize(data)
		return principal, l, err
	case stacks.PostConditionPrincipalTypeStandard:
		principal := &PostConditionStandardPrincipal{}
		l, err := principal.Deserialize(data)
		if err != nil {
			return nil, 0, err
		}
		return principal, l, nil
	case stacks.PostConditionPrincipalTypeContract:
		principal := &PostConditionContractPrincipal{}
		l, err := principal.Deserialize(data)
		if err != nil {
			return nil, 0, err
		}
		return principal, l, nil
	default:
		return nil, 0, errors.New("invalid post condition principal type")
	}
}

type PostConditionSTX struct {
	Principal     PostConditionPrincipal
	ConditionCode stacks.FungibleConditionCode
	Amount        uint64
}

func (pc *PostConditionSTX) GetType() stacks.PostConditionType {
	return stacks.PostConditionTypeSTX
}

func (pc *PostConditionSTX) Serialize() ([]byte, error) {
	buf := []byte{byte(stacks.PostConditionTypeSTX)}
	principalBytes, err := pc.Principal.Serialize()
	if err != nil {
		return nil, err
	}
	buf = append(buf, principalBytes...)
	buf = append(buf, byte(pc.ConditionCode))
	amountBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(amountBytes, pc.Amount)
	buf = append(buf, amountBytes...)
	return buf, nil
}

func (pc *PostConditionSTX) Deserialize(data []byte) (int, error) {
	if len(data) < 1 || stacks.PostConditionType(data[0]) != stacks.PostConditionTypeSTX {
		return 0, errors.New("invalid STX post condition")
	}
	offset := 1
	principal, l, err := DeserializePostConditionPrincipal(data[offset:])
	if err != nil {
		return 0, err
	}
	pc.Principal = principal
	offset += l
	pc.ConditionCode = stacks.FungibleConditionCode(data[offset])
	offset++
	pc.Amount = binary.BigEndian.Uint64(data[offset : offset+8])
	return offset + 8, nil
}

type PostConditionFungible struct {
	Principal     PostConditionPrincipal
	ConditionCode stacks.FungibleConditionCode
	Amount        uint64
	Asset         *AssetInfo
}

func (pc *PostConditionFungible) GetType() stacks.PostConditionType {
	return stacks.PostConditionTypeFungible
}

func (pc *PostConditionFungible) Serialize() ([]byte, error) {
	buf := []byte{byte(stacks.PostConditionTypeFungible)}
	principalBytes, err := pc.Principal.Serialize()
	if err != nil {
		return nil, err
	}
	buf = append(buf, principalBytes...)
	assetBytes, err := pc.Asset.Serialize()
	if err != nil {
		return nil, err
	}
	buf = append(buf, assetBytes...)
	buf = append(buf, byte(pc.ConditionCode))
	amountBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(amountBytes, pc.Amount)
	buf = append(buf, amountBytes...)
	return buf, nil
}

func (pc *PostConditionFungible) Deserialize(data []byte) (int, error) {
	if len(data) < 1 || stacks.PostConditionType(data[0]) != stacks.PostConditionTypeFungible {
		return 0, errors.New("invalid fungible post condition")
	}
	offset := 1
	principal, l, err := DeserializePostConditionPrincipal(data[offset:])
	if err != nil {
		return 0, err
	}
	pc.Principal = principal
	offset += l
	asset := &AssetInfo{}
	l, err = asset.Deserialize(data[offset:])
	if err != nil {
		return 0, err
	}
	pc.Asset = asset
	offset += l
	pc.ConditionCode = stacks.FungibleConditionCode(data[offset])
	offset++
	pc.Amount = binary.BigEndian.Uint64(data[offset : offset+8])
	return offset + 8, nil
}

type PostConditionNonFungible struct {
	Principal     PostConditionPrincipal
	ConditionCode stacks.NonFungibleConditionCode
	Asset         *AssetInfo
	AssetValue    clarity.ClarityValue
}

func (pc *PostConditionNonFungible) GetType() stacks.PostConditionType {
	return stacks.PostConditionTypeNonFungible
}

func (pc *PostConditionNonFungible) Serialize() ([]byte, error) {
	buf := []byte{byte(stacks.PostConditionTypeNonFungible)}
	principalBytes, err := pc.Principal.Serialize()
	if err != nil {
		return nil, err
	}
	buf = append(buf, principalBytes...)
	assetBytes, err := pc.Asset.Serialize()
	if err != nil {
		return nil, err
	}
	buf = append(buf, assetBytes...)
	assetValueBytes, err := pc.AssetValue.Serialize()
	if err != nil {
		return nil, err
	}
	buf = append(buf, assetValueBytes...)
	buf = append(buf, byte(pc.ConditionCode))
	return buf, nil
}

func (pc *PostConditionNonFungible) Deserialize(data []byte) (int, error) {
	if len(data) < 1 || stacks.PostConditionType(data[0]) != stacks.PostConditionTypeNonFungible {
		return 0, errors.New("invalid non fungible post condition")
	}
	offset := 1
	principal, l, err := DeserializePostConditionPrincipal(data[offset:])
	if err != nil {
		return 0, err
	}
	pc.Principal = principal
	offset += l
	asset := &AssetInfo{}
	l, err = asset.Deserialize(data[offset:])
	if err != nil {
		return 0, err
	}
	pc.Asset = asset
	offset += l
	assetValue, err := clarity.DeserializeClarityValue(data[offset:])
	if err != nil {
		return 0, err
	}
	assetValueBytes, err := assetValue.Serialize()
	if err != nil {
		return 0, err
	}
	pc.AssetValue = assetValue
	offset += len(assetValueBytes)
	pc.ConditionCode = stacks.NonFungibleConditionCode(data[offset])
	return offset + 1, nil
}

func DeserializePostCondition(data []byte) (PostCondition, int, error) {
	if len(data) < 1 {
		return nil, 0, errors.New("insufficient data for post condition")
	}
	postConditionType := stacks.PostConditionType(data[0])
	switch postConditionType {
	case stacks.PostConditionTypeSTX:
		postCondition := &PostConditionSTX{}
		l, err := postCondition.Deserialize(data)
		if err != nil {
			return nil, 0, err
		}
		return postCondition, l, nil
	case stacks.PostConditionTypeFungible:
		postCondition := &PostConditionFungible{}
		l, err := postCondition.Deserialize(data)
		if err != nil {
			return nil, 0, err
		}
		return postCondition, l, nil
	case stacks.PostConditionTypeNonFungible:
		postCondition := &PostConditionNonFungible{}
		l, err := postCondition.Deserialize(data)
		if err != nil {
			return nil, 0, err
		}
		return postCondition, l, nil
	default:
		return nil, 0, errors.New("invalid post condition type")
	}
}

func SerializePostConditions(postConditions []PostCondition) ([]byte, error) {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(len(postConditions)))
	for _, postCondition := range postConditions {
		postConditionBytes, err := postCondition.Serialize()
		if err != nil {
			return nil, err
		}
		buf = append(buf, postConditionBytes...)
	}
	return buf, nil
}

func DeserializePostConditions(data []byte) ([]PostCondition, int, error) {
	if len(data) < 4 {
		return nil, 0, errors.New("insufficient data for post conditions length")
	}
	length := binary.BigEndian.Uint32(data[:4])
	postConditions := make([]PostCondition, length)
	offset := 4
	for i := range postConditions {
		postCondition, l, err := DeserializePostCondition(data[offset:])
		if err != nil {
			return nil, 0, err
		}
		postConditions[i] = postCondition
		offset += l
	}
	return postConditions, offset, nil
}
