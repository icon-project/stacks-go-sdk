package transaction

import (
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyPostConditionsSerialization(t *testing.T) {
	emptyPostConditions := []PostCondition{}
	serialized := SerializePostConditions(emptyPostConditions)

	// Check that the serialized data is 4 bytes long (uint32 for length)
	assert.Equal(t, 4, len(serialized), "Serialized empty post conditions should be 4 bytes long")

	// Check that the serialized length is 0
	length := binary.BigEndian.Uint32(serialized)
	assert.Equal(t, uint32(0), length, "Serialized length of empty post conditions should be 0")
}

func TestEmptyPostConditionsDeserialization(t *testing.T) {
	serialized := make([]byte, 4)
	binary.BigEndian.PutUint32(serialized, 0)

	deserialized, bytesRead, err := DeserializePostConditions(serialized)

	assert.NoError(t, err, "Deserialization of empty post conditions should not produce an error")

	assert.Equal(t, 4, bytesRead, "Deserialization should have read 4 bytes")

	assert.Empty(t, deserialized, "Deserialized post conditions should be empty")
}

func TestEmptyPostConditionsRoundTrip(t *testing.T) {
	originalPostConditions := []PostCondition{}

	serialized := SerializePostConditions(originalPostConditions)

	deserialized, bytesRead, err := DeserializePostConditions(serialized)

	assert.NoError(t, err, "Round trip should not produce an error")
	assert.Equal(t, 4, bytesRead, "Round trip should read 4 bytes")
	assert.Empty(t, deserialized, "Round trip should result in empty post conditions")
	assert.Equal(t, len(originalPostConditions), len(deserialized), "Original and deserialized lengths should match")
}
