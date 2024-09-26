package c32

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/icon-project/stacks-go-sdk/pkg/stacks"
)

var crockfordAlphabet = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"

func C32Encode(data []byte) string {
	// Convert bytes to big.Int
	bi := new(big.Int).SetBytes(data)

	// Convert big.Int to base32 string
	var encoded strings.Builder
	for bi.Cmp(big.NewInt(0)) > 0 {
		mod := new(big.Int)
		bi.DivMod(bi, big.NewInt(32), mod)
		encoded.WriteByte(crockfordAlphabet[mod.Int64()])
	}

	// Reverse the string
	encodedStr := reverseString(encoded.String())

	// Handle leading zeros
	leadingZeros := 0
	for _, b := range data {
		if b == 0 {
			leadingZeros++
		} else {
			break
		}
	}
	for i := 0; i < leadingZeros; i++ {
		encodedStr = "0" + encodedStr
	}

	return encodedStr
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

func DecodeWithChecksum(c32addr string) (byte, []byte, error) {
	if len(c32addr) < 1 {
		return 0, nil, errors.New("address too short")
	}
	if c32addr[0] != 'S' {
		return 0, nil, errors.New("address must start with 'S'")
	}

	c32str := c32addr[1:]

	data, err := C32Decode(c32str)
	if err != nil {
		return 0, nil, err
	}

	if len(data) != 1+stacks.AddressHashLength+4 {
		return 0, nil, fmt.Errorf("invalid decoded length: expected %d, got %d", 1+stacks.AddressHashLength+4, len(data))
	}

	version := data[0]
	payload := data[1 : 1+stacks.AddressHashLength]
	checksum := data[1+stacks.AddressHashLength:]

	// Recompute checksum
	versionedData := append([]byte{version}, payload...)
	computedChecksum := sha256.Sum256(versionedData)
	computedChecksum = sha256.Sum256(computedChecksum[:])
	computedChecksum = sha256.Sum256(computedChecksum[:])
	computedChecksumBytes := computedChecksum[:4]

	// Compare checksums
	for i := 0; i < 4; i++ {
		if checksum[i] != computedChecksumBytes[i] {
			return 0, nil, errors.New("checksum mismatch")
		}
	}

	return version, payload, nil
}

func EncodeWithChecksum(version byte, data []byte) (string, error) {
	if len(data) != stacks.AddressHashLength {
		return "", errors.New("data must be 20 bytes for P2PKH")
	}

	// Version byte + data
	versionedData := append([]byte{version}, data...)

	// Compute checksum: double SHA256, first 4 bytes
	checksum := sha256.Sum256(versionedData)
	checksum = sha256.Sum256(checksum[:])
	checksumBytes := checksum[:4]

	// Append checksum
	fullData := append(data, checksumBytes...)

	// Encode to c32
	c32str := C32Encode(fullData)

	// Add prefix 'S'
	return "S" + string(crockfordAlphabet[version]) + c32str, nil
}

func SerializeAddress(version stacks.AddressVersion, hash160 []byte) (string, error) {
	return EncodeWithChecksum(byte(version), hash160)
}

func DeserializeAddress(address string) (stacks.AddressVersion, []byte, error) {
	version, payload, err := DecodeWithChecksum(address)
	if err != nil {
		return 0, nil, err
	}

	var addrVersion stacks.AddressVersion
	switch version {
	case byte(stacks.AddressVersionMainnetSingleSig):
		addrVersion = stacks.AddressVersionMainnetSingleSig
	case byte(stacks.AddressVersionTestnetSingleSig):
		addrVersion = stacks.AddressVersionTestnetSingleSig
	default:
		return 0, nil, fmt.Errorf("unknown address version: %d", version)
	}

	return addrVersion, payload, nil
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
