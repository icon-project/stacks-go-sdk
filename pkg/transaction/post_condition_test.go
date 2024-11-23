package transaction

import (
	"encoding/binary"
	"encoding/hex"
	"testing"

	"github.com/icon-project/stacks-go-sdk/pkg/c32"
	"github.com/icon-project/stacks-go-sdk/pkg/clarity"
	"github.com/icon-project/stacks-go-sdk/pkg/stacks"
	"github.com/stretchr/testify/assert"
)

func TestEmptyPostConditionsSerialization(t *testing.T) {
	emptyPostConditions := []PostCondition{}
	serialized, err := SerializePostConditions(emptyPostConditions)
	assert.NoError(t, err, "Serialization of empty post conditions should not produce an error")
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

	serialized, err := SerializePostConditions(originalPostConditions)
	assert.NoError(t, err, "Serialization of empty post conditions should not produce an error")

	deserialized, bytesRead, err := DeserializePostConditions(serialized)

	assert.NoError(t, err, "Round trip should not produce an error")
	assert.Equal(t, 4, bytesRead, "Round trip should read 4 bytes")
	assert.Empty(t, deserialized, "Round trip should result in empty post conditions")
	assert.Equal(t, len(originalPostConditions), len(deserialized), "Original and deserialized lengths should match")
}

func TestPostConditionSTXSerializationAndDeserialization(t *testing.T) {
	pcHex := "00031681d1b585c012192166f5afa30d1245701e2073a916756e6976322d706f6f6c2d76315f305f302d30303333030000000000197827"
	pcBytes, err := hex.DecodeString(pcHex)
	assert.NoError(t, err)
	pc, bytesRead, err := DeserializePostCondition(pcBytes)
	assert.NoError(t, err)
	assert.Equal(t, len(pcBytes), bytesRead)
	assert.IsType(t, &PostConditionSTX{}, pc)
	pcSTX, ok := pc.(*PostConditionSTX)
	assert.True(t, ok)
	assert.Equal(t, uint64(0x197827), pcSTX.Amount)
	assert.Equal(t, stacks.PostConditionPrincipalTypeContract, pcSTX.Principal.GetID())
	principal, ok := pcSTX.Principal.(*PostConditionContractPrincipal)
	assert.True(t, ok)
	version, hash, err := c32.DecodeC32Address("SP20X3DC5R091J8B6YPQT638J8NR1W83KN6TN5BJY")
	assert.NoError(t, err)
	assert.Equal(t, version, principal.Contract.Version)
	assert.Equal(t, hash, principal.Contract.Hash160)
	assert.Equal(t, "univ2-pool-v1_0_0-0033", principal.Contract.ContractName)
	assert.Equal(t, stacks.FungibleConditionCodeSentGe, pcSTX.ConditionCode)

	serialized, err := pcSTX.Serialize()
	assert.NoError(t, err)
	assert.Equal(t, pcHex, hex.EncodeToString(serialized))
}

func TestPostConditionFungibleSerializationAndDeserialization(t *testing.T) {
	pcHex := "010216465b9141e262fd03fc16d84781d8189d253fa5fd16e357773d94c58f5605b1bd9738a37e1dd310b2b40c6d656d652d73747863697479044d454d45010000008720176c3b"
	pcBytes, err := hex.DecodeString(pcHex)
	assert.NoError(t, err)
	pc, bytesRead, err := DeserializePostCondition(pcBytes)
	assert.NoError(t, err)
	assert.Equal(t, len(pcBytes), bytesRead)
	assert.IsType(t, &PostConditionFungible{}, pc)
	pcFungible, ok := pc.(*PostConditionFungible)
	assert.True(t, ok)
	assert.Equal(t, "MEME", pcFungible.Asset.AssetName)
	assert.Equal(t, uint64(0x8720176c3b), pcFungible.Amount)
	assert.Equal(t, stacks.FungibleConditionCodeSentEq, pcFungible.ConditionCode)

	serialized, err := pcFungible.Serialize()
	assert.NoError(t, err)
	assert.Equal(t, pcHex, hex.EncodeToString(serialized))
}

func TestPostConditionNonFungibleSerializationAndDeserialization(t *testing.T) {
	pcHex := "020216040330949aaa99d6c1aefa3058d758358b02fe3716000000000000000000000000000000000000000003626e73056e616d65730c00000002046e616d65020000000c6573636f72746c6164696573096e616d657370616365020000000362746310"
	pcBytes, err := hex.DecodeString(pcHex)
	assert.NoError(t, err)
	pc, bytesRead, err := DeserializePostCondition(pcBytes)
	assert.NoError(t, err)
	assert.Equal(t, len(pcBytes), bytesRead)
	assert.IsType(t, &PostConditionNonFungible{}, pc)
	pcNonFungible, ok := pc.(*PostConditionNonFungible)
	assert.True(t, ok)
	assert.Equal(t, "names", pcNonFungible.Asset.AssetName)

	assetValue, ok := pcNonFungible.AssetValue.(*clarity.Tuple)
	assert.True(t, ok)
	assetValueName, ok := assetValue.Data["name"]
	assert.True(t, ok)
	assetValueNamespace, ok := assetValue.Data["namespace"]
	assert.True(t, ok)
	assetValueNameBuffer, ok := assetValueName.(*clarity.Buffer)
	assert.True(t, ok)
	assetValueNamespaceBuffer, ok := assetValueNamespace.(*clarity.Buffer)
	assert.True(t, ok)
	assert.Equal(t, "escortladies", string(assetValueNameBuffer.Data))
	assert.Equal(t, "btc", string(assetValueNamespaceBuffer.Data))
	assert.Equal(t, stacks.NonFungibleConditionCodeSent, pcNonFungible.ConditionCode)

	serialized, err := pcNonFungible.Serialize()
	assert.NoError(t, err)
	assert.Equal(t, pcHex, hex.EncodeToString(serialized))
}
