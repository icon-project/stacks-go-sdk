package c32

import (
	"encoding/base32"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/icon-project/stacks-go-sdk/pkg/stacks"
)

var crockfordAlphabet = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"

func C32Encode(input []byte) string {
	encoder := base32.NewEncoding(crockfordAlphabet).WithPadding(base32.NoPadding)
	return encoder.EncodeToString(input)
}

func C32Decode(input string) ([]byte, error) {
	input = strings.ToUpper(strings.ReplaceAll(input, "-", ""))
	input = strings.ReplaceAll(input, "O", "0")
	input = strings.ReplaceAll(input, "I", "1")
	input = strings.ReplaceAll(input, "L", "1")

	bi := big.NewInt(0)
	for _, char := range input {
		bi.Mul(bi, big.NewInt(32))
		index := strings.IndexRune(crockfordAlphabet, char)
		if index == -1 {
			return nil, fmt.Errorf("invalid character: %c", char)
		}
		bi.Add(bi, big.NewInt(int64(index)))
	}

	bytes := bi.Bytes()

	for len(bytes) > 0 && bytes[0] == 0 {
		bytes = bytes[1:]
	}

	return bytes, nil
}

func DecodeC32Address(address string) (version byte, hash160 [20]byte, err error) {
	if len(address) < 5 || address[0] != 'S' {
		return 0, [20]byte{}, fmt.Errorf("invalid C32 address: must start with 'S' and be at least 5 characters long")
	}

	versionChar := address[1]
	version = byte(strings.IndexRune(crockfordAlphabet, rune(versionChar)))

	decoded, err := C32Decode(address[2:])
	if err != nil {
		return 0, [20]byte{}, fmt.Errorf("failed to decode address: %v", err)
	}

	if len(decoded) != 24 { // 20 bytes hash160 + 4 bytes checksum
		return 0, [20]byte{}, fmt.Errorf("invalid decoded length: expected 24, got %d", len(decoded))
	}

	copy(hash160[:], decoded[:20])

	return version, hash160, nil
}

func SerializeAddress(address string) ([]byte, error) {
	if len(address) != 1+20*2 { // 'S' + version + 40 hex chars
		return nil, fmt.Errorf("invalid address length: %d", len(address))
	}

	var version byte
	switch address[0] {
	case 'S':
		version = byte(stacks.AddressVersionMainnetSingleSig)
	case 'T':
		version = byte(stacks.AddressVersionTestnetSingleSig)
	default:
		return nil, fmt.Errorf("invalid address version: %c", address[0])
	}

	hashBytes, err := C32Decode(address[1:])
	if err != nil {
		return nil, fmt.Errorf("invalid address hash: %v", err)
	}

	result := make([]byte, 1+len(hashBytes))
	result[0] = version
	copy(result[1:], hashBytes)

	return result, nil
}

func DeserializeAddress(data []byte) (string, int, error) {
	if len(data) < 1+stacks.AddressHashLength {
		return "", 0, errors.New("insufficient data for address")
	}

	version := stacks.AddressVersion(data[0])
	var prefix string
	switch version {
	case stacks.AddressVersionMainnetSingleSig:
		prefix = "S"
	case stacks.AddressVersionTestnetSingleSig:
		prefix = "T"
	default:
		return "", 0, fmt.Errorf("invalid address version: %d", version)
	}

	c32hash := C32Encode(data[1 : 1+stacks.AddressHashLength+5])
	address := fmt.Sprintf("%s%s", prefix, c32hash)

	return address, 1 + stacks.AddressHashLength + 5, nil
}
