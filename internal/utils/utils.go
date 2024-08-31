package utils

import (
	"errors"

	"github.com/icon-project/stacks-go-sdk/pkg/stacks"
)

func SerializeString(s string, maxLength int) ([]byte, error) {
	if len(s) > maxLength {
		return nil, errors.New("string exceeds maximum length")
	}
	buf := make([]byte, 1+len(s))
	buf[0] = byte(len(s))
	copy(buf[1:], s)
	return buf, nil
}

func DeserializeString(data []byte, maxLength int) (string, int, error) {
	if len(data) < 1 {
		return "", 0, errors.New("insufficient data for string length")
	}
	length := int(data[0])
	if length > maxLength || len(data) < 1+length {
		return "", 0, errors.New("invalid string length")
	}
	return string(data[1 : 1+length]), 1 + length, nil
}

func IsValidAddressHashMode(mode stacks.AddressHashMode) bool {
	return mode == stacks.AddressHashModeSerializeP2PKH ||
		mode == stacks.AddressHashModeSerializeP2WPKH
}

func IsValidPubKeyEncoding(encoding stacks.PubKeyEncoding) bool {
	return encoding == stacks.PubKeyEncodingCompressed || encoding == stacks.PubKeyEncodingUncompressed
}

func IsCompatibleHashModeAndKeyEncoding(hashMode stacks.AddressHashMode, keyEncoding stacks.PubKeyEncoding) bool {
	// P2WPKH and P2WSH only support compressed public keys
	if (hashMode == stacks.AddressHashModeSerializeP2WPKH) &&
		keyEncoding != stacks.PubKeyEncodingCompressed {
		return false
	}
	return true
}
