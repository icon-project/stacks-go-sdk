package transaction

import (
	"encoding/binary"
	"errors"

	"github.com/icon-project/stacks-go-sdk/pkg/stacks"
)

type PostCondition interface {
	Serialize() ([]byte, error)
	Deserialize([]byte) (int, error)
	GetType() stacks.PostConditionType
}

// Assume empty post condition
func SerializePostConditions(postConditions []PostCondition) []byte {
	lenBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lenBytes, uint32(len(postConditions)))
	return lenBytes
}

// Assume empty post condition
func DeserializePostConditions(data []byte) ([]PostCondition, int, error) {
	if len(data) < 4 {
		return nil, 0, errors.New("insufficient data for post conditions length")
	}
	length := binary.BigEndian.Uint32(data[:4])
	return make([]PostCondition, length), 4, nil
}
