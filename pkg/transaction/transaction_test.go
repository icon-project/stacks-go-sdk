package transaction

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/icon-project/stacks-go-sdk/pkg/c32"
	"github.com/icon-project/stacks-go-sdk/pkg/clarity"
	"github.com/icon-project/stacks-go-sdk/pkg/crypto"
	"github.com/icon-project/stacks-go-sdk/pkg/stacks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainnetSTXTokenTransferTransactionSerializationAndDeserialization(t *testing.T) {
	// https://explorer.stxer.xyz/txid/0x07de75c128a47589327d177330faf88a300d4ee4b56de4831fae678b0fa15e1b
	txHex := "00000000010400c2cd5f81d93ca8d87aa2885e3b614c88f5f27b14000000000000000b0000000000020f080000ce732dfe7c17c11b7a7c48ee8be30b9a5c1d567d9dca228b9b565592e7dbdf303c3720c733ff8f12ab8fa4102339261b568e9ee4dd0a80643f6ad1cd79ee7022030200000000000516421a579acf61a681d9a6ca9e74938dd491e35b4b000000002e6e98c06b72616b656e00000000000000000000000000000000000000000000000000000000"
	txBytes, err := hex.DecodeString(txHex)
	assert.NoError(t, err)

	tx, err := DeserializeTransaction(txBytes)
	assert.NoError(t, err)

	assert.IsType(t, &TokenTransferTransaction{}, tx)

	payload, ok := tx.GetPayload().(*TokenTransferPayload)
	assert.True(t, ok)

	version, senderSigner, err := c32.DecodeC32Address("SP31CTQW1V4YAHP3TMA45WEV19J4FBWKV2JVDYDPT")
	assert.Equal(t, uint8(0x16), version)
	assert.NoError(t, err)
	assert.Equal(t, tx.GetAuth().OriginAuth.Signer, senderSigner)

	assert.Equal(t, uint64(779000000), payload.Amount)
	assert.Equal(t, "kraken", payload.Memo)
	recipient, err := clarity.StringToPrincipal("SP111MNWTSXGTD0ESMV59WX4KHQA93RTV9F82EK0K")
	assert.NoError(t, err)
	assert.Equal(t, recipient, payload.Recipient)

	txBytes, err = tx.Serialize()
	assert.NoError(t, err)
	assert.Equal(t, txHex, hex.EncodeToString(txBytes))
}

func TestSTXTokenTransferTransactionSerializationAndDeserialization(t *testing.T) {
	transactionVersion := stacks.TransactionVersionTestnet
	chainID := stacks.ChainIDTestnet
	anchorMode := stacks.AnchorModeOnChainOnly
	postConditionMode := stacks.PostConditionModeDeny

	recipientAddress := "SP3FGQ8Z7JY9BWYZ5WM53E0M9NK7WHJF0691NZ159"
	amount := uint64(2500000)
	memo := "memo (not included)"

	payload, err := NewTokenTransferPayload(recipientAddress, amount, memo)
	assert.NoError(t, err)

	addressHashMode := stacks.AddressHashModeSerializeP2PKH
	nonce := uint64(0)
	fee := uint64(0)

	pubKey := "03ef788b3830c00abe8f64f62dc32fc863bc0b2cafeb073b6c8e1c7657d9c2c3ab"
	pubKeyBytes, err := hex.DecodeString(pubKey)
	assert.NoError(t, err)

	secretKey := "edf9aee84d9b7abc145504dde6726c64f369d37ee34ded868fabd876c26570bc01"
	secretKeyBytes, err := hex.DecodeString(secretKey)
	assert.NoError(t, err)

	spendingCondition := SpendingCondition{
		HashMode:    addressHashMode,
		Signer:      [20]byte{},
		Nonce:       nonce,
		Fee:         fee,
		KeyEncoding: stacks.PubKeyEncodingCompressed,
		Signature:   [65]byte{},
	}

	auth := TransactionAuth{
		AuthType:   stacks.AuthTypeStandard,
		OriginAuth: spendingCondition,
	}

	transaction := &TokenTransferTransaction{
		BaseTransaction: BaseTransaction{
			Version:           transactionVersion,
			ChainID:           chainID,
			Auth:              auth,
			AnchorMode:        anchorMode,
			PostConditionMode: postConditionMode,
			PostConditions:    []PostCondition{},
		},
		Payload: *payload,
	}

	err = SignTransaction(transaction, secretKeyBytes)
	assert.NoError(t, err)

	isValid, err := VerifyTransaction(transaction, pubKeyBytes)
	assert.NoError(t, err)

	assert.True(t, isValid)

	serialized, err := transaction.Serialize()
	assert.NoError(t, err)

	deserialized, err := DeserializeTransaction(serialized)
	assert.NoError(t, err)

	tokenTx, ok := deserialized.(*TokenTransferTransaction)
	assert.True(t, ok, "Deserialized transaction is not a TokenTransferTransaction")

	assert.Equal(t, transactionVersion, tokenTx.Version)
	assert.Equal(t, chainID, tokenTx.ChainID)
	assert.Equal(t, stacks.AuthTypeStandard, tokenTx.Auth.AuthType)
	assert.Equal(t, addressHashMode, tokenTx.Auth.OriginAuth.HashMode)
	assert.Equal(t, nonce, tokenTx.Auth.OriginAuth.Nonce)
	assert.Equal(t, fee, tokenTx.Auth.OriginAuth.Fee)
	assert.Equal(t, anchorMode, tokenTx.AnchorMode)
	assert.Equal(t, postConditionMode, tokenTx.PostConditionMode)
	assert.Empty(t, tokenTx.PostConditions)

	recipientPrincipal, _ := clarity.StringToPrincipal(recipientAddress)

	assert.Equal(t, recipientPrincipal, tokenTx.Payload.Recipient)
	assert.Equal(t, amount, tokenTx.Payload.Amount)
	assert.Equal(t, memo, tokenTx.Payload.Memo)

	isValid, err = VerifyTransaction(tokenTx, pubKeyBytes)
	assert.NoError(t, err)
	assert.True(t, isValid)

	reserializedBytes, err := deserialized.Serialize()
	assert.NoError(t, err)
	assert.Equal(t, serialized, reserializedBytes)
}

func TestNewTokenTransferTransaction(t *testing.T) {
	tests := []struct {
		name              string
		recipient         string
		amount            uint64
		memo              string
		version           stacks.TransactionVersion
		chainID           stacks.ChainID
		signer            [20]byte
		nonce             uint64
		fee               uint64
		postConditionMode stacks.PostConditionMode
		postConditions    []PostCondition
	}{
		{
			name:              "Valid transaction",
			recipient:         "ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM",
			amount:            1000000,
			memo:              "Test transfer",
			version:           stacks.TransactionVersionMainnet,
			chainID:           stacks.ChainIDMainnet,
			signer:            [20]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			nonce:             1,
			fee:               1000,
			postConditionMode: stacks.PostConditionModeAllow,
			postConditions:    []PostCondition{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, err := NewTokenTransferTransaction(
				tt.recipient,
				tt.amount,
				tt.memo,
				tt.version,
				tt.chainID,
				tt.signer,
				tt.nonce,
				tt.fee,
				tt.postConditionMode,
				tt.postConditions,
			)

			require.NoError(t, err)
			require.NotNil(t, tx)

			// Check initial values
			assertTokenTransferTransactionFields(t, tx, tt)

			// Serialize the transaction
			serialized, err := tx.Serialize()
			require.NoError(t, err)
			require.NotEmpty(t, serialized)

			// Deserialize the transaction
			deserialized, err := DeserializeTransaction(serialized)
			require.NoError(t, err)
			require.NotNil(t, deserialized)

			// Check that the deserialized transaction is of the correct type
			deserializedTx, ok := deserialized.(*TokenTransferTransaction)
			require.True(t, ok, "Deserialized transaction is not a TokenTransferTransaction")

			// Check that all fields are preserved after serialization and deserialization
			assertTokenTransferTransactionFields(t, deserializedTx, tt)
		})
	}
}

func TestContractCallTransactionSerializationAndDeserialization(t *testing.T) {
	transactionVersion := stacks.TransactionVersionTestnet
	chainID := stacks.ChainIDTestnet
	anchorMode := stacks.AnchorModeAny
	postConditionMode := stacks.PostConditionModeDeny

	contractAddress := "ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM"
	contractName := "test-contract"
	functionName := "test-function"

	intArg, err := clarity.NewInt(1)
	require.NoError(t, err, "Failed to create int argument")

	strArg, err := clarity.NewStringASCII("test")
	require.NoError(t, err, "Failed to create string argument")

	functionArgs := []clarity.ClarityValue{
		intArg,
		clarity.NewBool(true),
		strArg,
	}

	addressHashMode := stacks.AddressHashModeSerializeP2PKH
	nonce := uint64(10)
	fee := uint64(456)

	pubKey := "0332fc778e5beb5f944c75b2b63c21dd12c40bdcdf99ba0663168ae0b2be880aef"
	pubKeyBytes, err := hex.DecodeString(pubKey)
	assert.NoError(t, err)

	secretKey := "c1d5bb638aa70862621667f9997711fce692cad782694103f8d9561f62e9f19701"
	secretKeyBytes, err := hex.DecodeString(secretKey)
	assert.NoError(t, err)

	senderPublicKey := crypto.GetPublicKeyFromPrivate(secretKeyBytes)
	var signer [20]byte
	copy(signer[:], crypto.Hash160(senderPublicKey))

	transaction, err := NewContractCallTransaction(contractAddress, contractName, functionName, functionArgs, transactionVersion, chainID, signer, nonce, fee, postConditionMode, []PostCondition{})
	assert.NoError(t, err)

	err = SignTransaction(transaction, secretKeyBytes)
	assert.NoError(t, err)

	isValid, err := VerifyTransaction(transaction, pubKeyBytes)
	assert.NoError(t, err)
	assert.True(t, isValid)

	serialized, err := transaction.Serialize()
	assert.NoError(t, err)

	deserialized, err := DeserializeTransaction(serialized)
	assert.NoError(t, err)

	contractTx, ok := deserialized.(*ContractCallTransaction)
	assert.True(t, ok, "Deserialized transaction is not a ContractCallTransaction")

	assert.Equal(t, transactionVersion, contractTx.Version)
	assert.Equal(t, chainID, contractTx.ChainID)
	assert.Equal(t, stacks.AuthTypeStandard, contractTx.Auth.AuthType)
	assert.Equal(t, addressHashMode, contractTx.Auth.OriginAuth.HashMode)
	assert.Equal(t, nonce, contractTx.Auth.OriginAuth.Nonce)
	assert.Equal(t, fee, contractTx.Auth.OriginAuth.Fee)
	assert.Equal(t, anchorMode, contractTx.AnchorMode)
	assert.Equal(t, postConditionMode, contractTx.PostConditionMode)
	assert.Empty(t, contractTx.PostConditions)

	contractPrincipal, _ := clarity.StringToPrincipal(contractAddress)

	assert.Equal(t, contractPrincipal, contractTx.Payload.ContractAddress)
	assert.Equal(t, contractName, contractTx.Payload.ContractName)
	assert.Equal(t, functionName, contractTx.Payload.FunctionName)
	assert.Equal(t, len(functionArgs), len(contractTx.Payload.FunctionArgs))

	for i, arg := range functionArgs {
		assert.Equal(t, arg.Type(), contractTx.Payload.FunctionArgs[i].Type())
		switch typedArg := arg.(type) {
		case *clarity.Int:
			deserializedArg, ok := contractTx.Payload.FunctionArgs[i].(*clarity.Int)
			assert.True(t, ok)
			assert.Equal(t, 0, typedArg.Value.Cmp(deserializedArg.Value))
		case clarity.Bool:
			deserializedArg, ok := contractTx.Payload.FunctionArgs[i].(clarity.Bool)
			assert.True(t, ok)
			assert.Equal(t, typedArg, deserializedArg)
		case *clarity.StringASCII:
			deserializedArg, ok := contractTx.Payload.FunctionArgs[i].(*clarity.StringASCII)
			assert.True(t, ok)
			assert.Equal(t, typedArg.Data, deserializedArg.Data)
		}
	}

	isValid, err = VerifyTransaction(contractTx, pubKeyBytes)
	assert.NoError(t, err)
	assert.True(t, isValid)

	reserializedBytes, err := deserialized.Serialize()
	assert.NoError(t, err)
	assert.Equal(t, serialized, reserializedBytes)
}

func TestMainnetContractCallTransactionSerializationAndDeserialization(t *testing.T) {
	// https://explorer.stxer.xyz/txid/0xdc11af04b0d4f490d9925f4e3278304d74a2316ffd1ce94c30ffc30c723e4414
	txHex := "00000000010400465b9141e262fd03fc16d84781d8189d253fa5fd0000000000000b2c0000000000002710000019487a42bd01bef957ff4949bdaa57b3caa13e85524e981d9b20ffc2759633ab148c56cc72ef5ba9569b0aec4df194dddcfcd957053410ed37e207bc7c993fca030200000003010216465b9141e262fd03fc16d84781d8189d253fa5fd16e357773d94c58f5605b1bd9738a37e1dd310b2b40c6d656d652d73747863697479044d454d45010000008720176c3b01031681d1b585c012192166f5afa30d1245701e2073a916756e6976322d706f6f6c2d76315f305f302d3030333316e357773d94c58f5605b1bd9738a37e1dd310b2b40c6d656d652d73747863697479044d454d4503000000000000000000031681d1b585c012192166f5afa30d1245701e2073a916756e6976322d706f6f6c2d76315f305f302d3030333303000000000019782702167c5f674a8fd08efa61dd9b11121e046dd2c8927311706174682d6170706c795f76315f325f30056170706c79000000180b000000010c0000000601610d00000001760162061681d1b585c012192166f5afa30d1245701e2073a916756e6976322d706f6f6c2d76315f305f302d3030333301630100000000000000000000000001406f6101640616e357773d94c58f5605b1bd9738a37e1dd310b2b40c6d656d652d73747863697479016506167c5f674a8fd08efa61dd9b11121e046dd2c8927304777374780166040100000000000000000000008720176c3b0a0616e357773d94c58f5605b1bd9738a37e1dd310b2b40c6d656d652d737478636974790a06167c5f674a8fd08efa61dd9b11121e046dd2c8927304777374780909090a06167c5f674a8fd08efa61dd9b11121e046dd2c8927312756e6976322d73686172652d6665652d746f0a061681d1b585c012192166f5afa30d1245701e2073a916756e6976322d706f6f6c2d76315f305f302d303033330909090a061681d1b585c012192166f5afa30d1245701e2073a916756e6976322d666565732d76315f305f302d303033330909090909090909090909"
	txBytes, err := hex.DecodeString(txHex)
	assert.NoError(t, err)

	tx, err := DeserializeTransaction(txBytes)
	assert.NoError(t, err)

	assert.IsType(t, &ContractCallTransaction{}, tx)

	payload, ok := tx.GetPayload().(*ContractCallPayload)
	assert.True(t, ok)

	version, senderSigner, err := c32.DecodeC32Address("SP135Q4A1W9HFT0ZW2VC4F0ER32EJAFX5ZME05JYW")
	assert.Equal(t, uint8(0x16), version)
	assert.NoError(t, err)
	assert.Equal(t, tx.GetAuth().OriginAuth.Signer, senderSigner)

	assert.Equal(t, "path-apply_v1_2_0", payload.ContractName)
	assert.Equal(t, "apply", payload.FunctionName)

	txBytes, err = tx.Serialize()
	assert.NoError(t, err)
	assert.Equal(t, txHex, hex.EncodeToString(txBytes))
}

func TestSmartContractTransactionInterfaces(t *testing.T) {
	contractName := "test-contract"
	codeBody := `(define-data-var counter int 0)`
	transactionVersion := stacks.TransactionVersionTestnet
	chainID := stacks.ChainIDTestnet
	postConditionMode := stacks.PostConditionModeDeny

	var signer [20]byte
	copy(signer[:], bytes.Repeat([]byte{1}, 20))

	tx, err := NewSmartContractTransaction(
		contractName,
		codeBody,
		stacks.ClarityVersionUnspecified,
		transactionVersion,
		chainID,
		signer,
		0,
		0,
		postConditionMode,
		[]PostCondition{},
	)
	assert.NoError(t, err)

	// Test GetAuth
	auth := tx.GetAuth()
	assert.NotNil(t, auth)
	assert.Equal(t, stacks.AuthTypeStandard, auth.AuthType)
	assert.Equal(t, signer, auth.OriginAuth.Signer)

	// Test GetPayload
	payload := tx.GetPayload()
	assert.NotNil(t, payload)
	contractPayload, ok := payload.(*SmartContractPayload)
	assert.True(t, ok)
	assert.Equal(t, contractName, contractPayload.ContractName)
	assert.Equal(t, codeBody, contractPayload.CodeBody)

	// Verify transaction implements StacksTransaction interface
	var _ StacksTransaction = (*SmartContractTransaction)(nil)
}

func TestSmartContractTransactionSerializationAndDeserialization(t *testing.T) {
	contractName := "test-contract"
	codeBody := `(define-data-var counter int 0)
(define-public (increment)
    (ok (var-set counter (+ (var-get counter) 1))))
`
	transactionVersion := stacks.TransactionVersionTestnet
	chainID := stacks.ChainIDTestnet
	anchorMode := stacks.AnchorModeAny
	postConditionMode := stacks.PostConditionModeDeny

	addressHashMode := stacks.AddressHashModeSerializeP2PKH
	nonce := uint64(1)
	fee := uint64(1000)

	pubKey := "0332fc778e5beb5f944c75b2b63c21dd12c40bdcdf99ba0663168ae0b2be880aef"
	pubKeyBytes, err := hex.DecodeString(pubKey)
	assert.NoError(t, err)

	secretKey := "c1d5bb638aa70862621667f9997711fce692cad782694103f8d9561f62e9f19701"
	secretKeyBytes, err := hex.DecodeString(secretKey)
	assert.NoError(t, err)

	senderPublicKey := crypto.GetPublicKeyFromPrivate(secretKeyBytes)
	var signer [20]byte
	copy(signer[:], crypto.Hash160(senderPublicKey))

	transaction, err := NewSmartContractTransaction(
		contractName,
		codeBody,
		stacks.ClarityVersionUnspecified,
		transactionVersion,
		chainID,
		signer,
		nonce,
		fee,
		postConditionMode,
		[]PostCondition{},
	)
	assert.NoError(t, err)

	err = SignTransaction(transaction, secretKeyBytes)
	assert.NoError(t, err)

	isValid, err := VerifyTransaction(transaction, pubKeyBytes)
	assert.NoError(t, err)
	assert.True(t, isValid)

	serialized, err := transaction.Serialize()
	assert.NoError(t, err)

	deserialized, err := DeserializeTransaction(serialized)
	assert.NoError(t, err)

	contractTx, ok := deserialized.(*SmartContractTransaction)
	assert.True(t, ok, "Deserialized transaction is not a SmartContractTransaction")

	// Verify fields match
	assert.Equal(t, transactionVersion, contractTx.Version)
	assert.Equal(t, chainID, contractTx.ChainID)
	assert.Equal(t, stacks.AuthTypeStandard, contractTx.Auth.AuthType)
	assert.Equal(t, addressHashMode, contractTx.Auth.OriginAuth.HashMode)
	assert.Equal(t, nonce, contractTx.Auth.OriginAuth.Nonce)
	assert.Equal(t, fee, contractTx.Auth.OriginAuth.Fee)
	assert.Equal(t, anchorMode, contractTx.AnchorMode)
	assert.Equal(t, postConditionMode, contractTx.PostConditionMode)
	assert.Empty(t, contractTx.PostConditions)

	assert.Equal(t, contractName, contractTx.Payload.ContractName)
	assert.Equal(t, codeBody, contractTx.Payload.CodeBody)

	// Verify signature is still valid
	isValid, err = VerifyTransaction(contractTx, pubKeyBytes)
	assert.NoError(t, err)
	assert.True(t, isValid)

	// Verify re-serialization matches
	reserializedBytes, err := deserialized.Serialize()
	assert.NoError(t, err)
	assert.Equal(t, serialized, reserializedBytes)
}

func TestMainnetVersionedSmartContractTransactionSerializationAndDeserialization(t *testing.T) {
	// https://explorer.stxer.xyz/txid/0xdbc688121e3d3d4632710b43b63d0114b02236bfe46973037e130f31923ec3e5
	txHex := "00000000010400040330949aaa99d6c1aefa3058d758358b02fe37000000000000000200000000000b71b00001f9f0db80562fa27374f8059d5138f5be09a250cf978997efa7e58f9cc55ff4fa578a89c1a910bffa792bcf60a9cfd4405db48e4b414323d8b4d37d06a3fe3a20030200000001020216040330949aaa99d6c1aefa3058d758358b02fe3716000000000000000000000000000000000000000003626e73056e616d65730c00000002046e616d65020000000c6573636f72746c6164696573096e616d657370616365020000000362746310060214626e732d313638343237303935343838372d7631000012b50a20203b3b2076657273696f6e3a20310a20203b3b206e616d653a206573636f72746c61646965730a20203b3b206e616d6573706163653a206274630a0a2020287573652d747261697420636f6d6d697373696f6e2d74726169742027535033443650563241434250454b594a54434d483748454e30324b503837515350384b5445483333352e636f6d6d697373696f6e2d74726169742e636f6d6d697373696f6e290a20200a202028646566696e652d636f6e7374616e74204445504c4f5945525f434f4e54524143545f5052494e434950414c202861732d636f6e74726163742074782d73656e64657229290a202028646566696e652d636f6e7374616e7420434f4d4d2d4144445220275350374e445836595241483643393957435a4a574b4632535952314751524637583638393451534a290a20200a202028646566696e652d636f6e7374616e74204552522d414c52454144592d4c49535445442028657272207534303129290a202028646566696e652d636f6e7374616e74204552522d57524f4e472d434f4d4d495353494f4e2028657272207534303229290a202028646566696e652d636f6e7374616e74204552522d4e4f542d415554484f52495a45442028657272207534303329290a202028646566696e652d636f6e7374616e74204552522d4e4f542d464f554e442028657272207534303429290a202028646566696e652d636f6e7374616e74204552522d57524f4e472d50524943452028657272207534303529290a202028646566696e652d636f6e7374616e74204552522d5452414e534645522d4641494c45442028657272207535303029290a20200a202028646566696e652d646174612d7661722063757272656e742d6e616d657370616365202862756666203230292030783030290a202028646566696e652d646174612d7661722063757272656e742d6e616d65202862756666203438292030783030290a20200a202028646566696e652d6d6170206c697374696e6773207b206e616d6573706163653a202862756666203230292c206e616d653a20286275666620343829207d207b2070726963653a2075696e742c206c69737465723a207072696e636970616c2c20636f6d6d697373696f6e3a207072696e636970616c207d290a20200a202028646566696e652d726561642d6f6e6c79202869732d61646d696e290a202020202869732d65712074782d73656e64657220434f4d4d2d41444452290a2020290a20200a202028646566696e652d726561642d6f6e6c7920286765742d6c697374696e67290a20202020286d61702d6765743f206c697374696e6773207b206e616d6573706163653a20287661722d6765742063757272656e742d6e616d657370616365292c206e616d653a20287661722d6765742063757272656e742d6e616d6529207d290a2020290a20200a202028646566696e652d726561642d6f6e6c7920286765742d63757272656e742d6e616d65290a202020207b206e616d6573706163653a20287661722d6765742063757272656e742d6e616d657370616365292c206e616d653a20287661722d6765742063757272656e742d6e616d6529207d0a2020290a20200a202028646566696e652d7072697661746520286c6973742d6e616d6520286e616d657370616365202862756666203230292920286e616d652028627566662034382929202870726963652075696e74292028636f6d6d697373696f6e203c636f6d6d697373696f6e2d74726169743e29290a20202020202028626567696e0a20202020202020202020286173736572747321202869732d6e6f6e6520286765742d6c697374696e672929204552522d414c52454144592d4c4953544544290a2020202020202020202028747279212028746f2d626f6f6c2d726573706f6e73652028636f6e74726163742d63616c6c3f0a20202020202020202020202020202753503030303030303030303030303030303030303030325136564637382e626e730a20202020202020202020202020206e616d652d7472616e73666572200a20202020202020202020202020206e616d6573706163650a20202020202020202020202020206e616d650a20202020202020202020202020204445504c4f5945525f434f4e54524143545f5052494e434950414c0a20202020202020202020202020206e6f6e650a202020202020202020202929290a20202020202020202020287661722d7365742063757272656e742d6e616d657370616365206e616d657370616365290a20202020202020202020287661722d7365742063757272656e742d6e616d65206e616d65290a20202020202020202020286f6b20286d61702d736574206c697374696e6773207b6e616d653a206e616d652c206e616d6573706163653a206e616d6573706163657d200a20202020202020202020202020202020202020202020202020202020202020207b70726963653a2070726963652c206c69737465723a2074782d73656e6465722c20636f6d6d697373696f6e3a2028636f6e74726163742d6f6620636f6d6d697373696f6e297d29290a202020202020290a2020290a20200a202028646566696e652d7075626c696320286368616e67652d707269636520286e616d657370616365202862756666203230292920286e616d65202862756666203438292920286e65772d70726963652075696e74292028636f6d6d697373696f6e203c636f6d6d697373696f6e2d74726169743e29290a20202020286c657420280a2020202020202020286c697374696e672028756e777261702120286d61702d6765743f206c697374696e6773207b6e616d6573706163653a206e616d6573706163652c206e616d653a206e616d657d29204552522d4e4f542d464f554e4429290a20202020202020202870726963652028676574207072696365206c697374696e6729290a2020202020202020286c69737465722028676574206c6973746572206c697374696e672929290a202020202020286173736572747321202869732d65712074782d73656e646572206c697374657229204552522d4e4f542d415554484f52495a4544290a202020202020286f6b20286d61702d736574206c697374696e6773207b206e616d6573706163653a206e616d6573706163652c206e616d653a206e616d65207d207b2070726963653a206e65772d70726963652c206c69737465723a206c69737465722c20636f6d6d697373696f6e3a2028636f6e74726163742d6f6620636f6d6d697373696f6e29207d29290a20202020290a2020290a20200a202028646566696e652d7075626c69632028756e6c6973742d6e616d6520286e616d657370616365202862756666203230292920286e616d652028627566662034382929290a20202020286c657420280a2020202020202020286c697374696e672028756e777261702120286d61702d6765743f206c697374696e6773207b6e616d6573706163653a206e616d6573706163652c206e616d653a206e616d657d29204552522d4e4f542d464f554e4429290a20202020202020202870726963652028676574207072696365206c697374696e6729290a2020202020202020286c69737465722028676574206c6973746572206c697374696e672929290a20202020202028617373657274732120286f72202869732d65712074782d73656e646572206c697374657229202869732d61646d696e2929204552522d4e4f542d415554484f52495a4544290a202020202020286d61702d64656c657465206c697374696e6773207b6e616d6573706163653a206e616d6573706163652c206e616d653a206e616d657d290a2020202020202861732d636f6e74726163740a2020202020202020202028746f2d626f6f6c2d726573706f6e73652028636f6e74726163742d63616c6c3f0a20202020202020202020202020202753503030303030303030303030303030303030303030325136564637382e626e73200a20202020202020202020202020206e616d652d7472616e73666572200a20202020202020202020202020206e616d6573706163650a20202020202020202020202020206e616d650a20202020202020202020202020206c69737465720a20202020202020202020202020206e6f6e65200a2020202020202020202029290a202020202020290a20202020290a2020290a20200a202028646566696e652d7075626c6963202870757263686173652d6e616d6520286e616d657370616365202862756666203230292920286e616d652028627566662034382929202865787065637465642d70726963652075696e74292028636f6d6d697373696f6e203c636f6d6d697373696f6e2d74726169743e292028726563697069656e7420286f7074696f6e616c207072696e636970616c2929290a202020202020286c657420280a2020202020202020286e65772d6f776e657220286966202869732d736f6d6520726563697069656e74292028756e777261702d70616e696320726563697069656e74292074782d73656e64657229290a2020202020202020286c697374696e672028756e777261702120286d61702d6765743f206c697374696e6773207b6e616d6573706163653a206e616d6573706163652c206e616d653a206e616d657d29204552522d4e4f542d464f554e4429290a20202020202020202870726963652028676574207072696365206c697374696e6729290a2020202020202020286c69737465722028676574206c6973746572206c697374696e6729290a2020202020202020286c6973742d636f6d6d697373696f6e202867657420636f6d6d697373696f6e206c697374696e6729290a202020202020290a2020202020202020286173736572747321202869732d65712028636f6e74726163742d6f6620636f6d6d697373696f6e29206c6973742d636f6d6d697373696f6e29204552522d57524f4e472d434f4d4d495353494f4e290a2020202020202020286173736572747321202869732d65712070726963652065787065637465642d707269636529204552522d57524f4e472d5052494345290a202020202020202028747279212028636f6e74726163742d63616c6c3f20636f6d6d697373696f6e2070617920753020707269636529290a2020202020202020287472792120287374782d7472616e736665723f2070726963652074782d73656e646572206c697374657229290a2020202020202020286d61702d64656c657465206c697374696e6773207b6e616d6573706163653a206e616d6573706163652c206e616d653a206e616d657d290a202020202020202028746f2d626f6f6c2d726573706f6e7365202861732d636f6e74726163740a20202020202020202020202028636f6e74726163742d63616c6c3f0a202020202020202020202020202020202753503030303030303030303030303030303030303030325136564637382e626e73200a202020202020202020202020202020206e616d652d7472616e73666572200a202020202020202020202020202020206e616d6573706163650a202020202020202020202020202020206e616d650a202020202020202020202020202020206e65772d6f776e65720a202020202020202020202020202020206e6f6e65200a202020202020202020202020290a202020202020202029290a202020202020290a2020290a20200a202028646566696e652d7075626c6963202877697468647261772d7374782028616d6f756e742075696e7429290a202020202020286c657420280a2020202020202020286c697374696e672028756e777261702120286765742d6c697374696e6729204552522d4e4f542d464f554e4429290a2020202020202020286c69737465722028676574206c6973746572206c697374696e6729290a2020202020202920200a202020202020202028617373657274732120286f72202869732d65712074782d73656e646572206c697374657229202869732d61646d696e2929204552522d4e4f542d415554484f52495a4544290a20202020202020202874727921202861732d636f6e747261637420287374782d7472616e736665723f20616d6f756e742074782d73656e646572206c69737465722929290a2020202020202020286f6b20616d6f756e74290a202020202020290a2020290a0a202028646566696e652d707269766174652028746f2d626f6f6c2d726573706f6e7365202876616c75652028726573706f6e736520626f6f6c20696e742929290a202020202020286d617463682076616c75650a202020202020202020202020207375636365737320286f6b2073756363657373290a202020202020202020202020206572726f7220286572722028746f2d75696e74206572726f72292929290a20200a2020286c6973742d6e616d6520307836323734363320307836353733363336663732373436633631363436393635373320753139333232373035333135202753504e575a3556325450574751475644523654374236525134584d475a34505854454530565130532e67616d6d612d636f6d6d697373696f6e2d332d35290a2020"
	txBytes, err := hex.DecodeString(txHex)
	assert.NoError(t, err)

	tx, err := DeserializeTransaction(txBytes)
	assert.NoError(t, err)

	assert.IsType(t, &SmartContractTransaction{}, tx)

	version, senderSigner, err := c32.DecodeC32Address("SP206C4MKAN9KNP1NVX30P6QB0TRP0QY6X479H82")
	assert.Equal(t, uint8(0x16), version)
	assert.NoError(t, err)
	assert.Equal(t, tx.GetAuth().OriginAuth.Signer, senderSigner)
	sc, ok := tx.GetPayload().(*SmartContractPayload)
	assert.True(t, ok)
	assert.Equal(t, stacks.ClarityVersion2, sc.ClarityVersion)
	assert.Equal(t, "bns-1684270954887-v1", sc.ContractName)

	txBytes, err = tx.Serialize()
	assert.NoError(t, err)
	assert.Equal(t, txHex, hex.EncodeToString(txBytes))
}

func assertTokenTransferTransactionFields(t *testing.T, tx *TokenTransferTransaction, expected struct {
	name              string
	recipient         string
	amount            uint64
	memo              string
	version           stacks.TransactionVersion
	chainID           stacks.ChainID
	signer            [20]byte
	nonce             uint64
	fee               uint64
	postConditionMode stacks.PostConditionMode
	postConditions    []PostCondition
},
) {
	assert.Equal(t, expected.version, tx.Version)
	assert.Equal(t, expected.chainID, tx.ChainID)
	assert.Equal(t, expected.signer, tx.Auth.OriginAuth.Signer)
	assert.Equal(t, expected.nonce, tx.Auth.OriginAuth.Nonce)
	assert.Equal(t, expected.fee, tx.Auth.OriginAuth.Fee)
	assert.Equal(t, expected.postConditionMode, tx.PostConditionMode)
	assert.Equal(t, expected.amount, tx.Payload.Amount)
	assert.Equal(t, expected.memo, tx.Payload.Memo)

	recipient := tx.Payload.Recipient
	expectedRecipient, _ := clarity.StringToPrincipal(expected.recipient)
	assert.Equal(t, expectedRecipient, recipient)
}
