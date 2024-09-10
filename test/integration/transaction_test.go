package transaction_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/icon-project/stacks-go-sdk/pkg/clarity"
	"github.com/icon-project/stacks-go-sdk/pkg/crypto"
	"github.com/icon-project/stacks-go-sdk/pkg/stacks"
	"github.com/icon-project/stacks-go-sdk/pkg/transaction"
	"github.com/stretchr/testify/require"
)

func TestBroadcastSTXTokenTransferTransaction(t *testing.T) {
	mnemonic := "vapor unhappy gather snap project ball gain puzzle comic error avocado bounce letter anxiety wheel provide canyon promote sniff improve figure daughter mansion baby"
	privateKey, err := crypto.DeriveStxPrivateKey(mnemonic, 0)
	if err != nil {
		t.Fatalf("Failed to derive private key: %v", err)
	}

	senderAddress := "ST15C893XJFJ6FSKM020P9JQDB5T7X6MQTXMBPAVH"
	senderPublicKey := crypto.GetPublicKeyFromPrivate(privateKey)
	var signerArray [20]byte
	copy(signerArray[:], crypto.Hash160(senderPublicKey))

	recipient := "ST3YJD5Y1WTMC8R09ZKR3HJF562R3NM8HHXW2S2R9"
	amount := big.NewInt(1000000) // 1 STX
	memo := "Test transfer"
	network := stacks.NewStacksTestnet()
	tx, err := transaction.MakeSTXTokenTransfer(recipient, *amount, memo, *network, senderAddress, privateKey, nil, nil)
	if err != nil {
		t.Fatalf("Failed to create transaction: %v", err)
	}

	err = transaction.SignTransaction(tx, privateKey)
	if err != nil {
		t.Fatalf("Failed to sign transaction: %v", err)
	}

	txID, err := transaction.BroadcastTransaction(tx, network)
	if err != nil {
		t.Fatalf("Failed to broadcast transaction: %v", err)
	}

	fmt.Printf("Transaction broadcasted successfully. TxID: %s\n", txID)
}

func TestBroadcastContractCallTransaction(t *testing.T) {
	mnemonic := "vapor unhappy gather snap project ball gain puzzle comic error avocado bounce letter anxiety wheel provide canyon promote sniff improve figure daughter mansion baby"
	privateKey, err := crypto.DeriveStxPrivateKey(mnemonic, 0)
	if err != nil {
		t.Fatalf("Failed to derive private key: %v", err)
	}

	senderAddress := "ST15C893XJFJ6FSKM020P9JQDB5T7X6MQTXMBPAVH"
	senderPublicKey := crypto.GetPublicKeyFromPrivate(privateKey)
	var signerArray [20]byte
	copy(signerArray[:], crypto.Hash160(senderPublicKey))

	contractAddress := "ST15C893XJFJ6FSKM020P9JQDB5T7X6MQTXMBPAVH"
	contractName := "contract_name"
	functionName := "address-string-to-principal"

	strArg, err := clarity.NewStringASCII("test")
	require.NoError(t, err, "Failed to create string argument")

	functionArgs := []clarity.ClarityValue{
		strArg,
	}

	network := stacks.NewStacksTestnet()
	tx, err := transaction.MakeContractCall(contractAddress, contractName, functionName, functionArgs, *network, senderAddress, privateKey, nil, nil)
	if err != nil {
		t.Fatalf("Failed to create transaction: %v", err)
	}

	err = transaction.SignTransaction(tx, privateKey)
	if err != nil {
		t.Fatalf("Failed to sign transaction: %v", err)
	}

	txID, err := transaction.BroadcastTransaction(tx, network)
	if err != nil {
		t.Fatalf("Failed to broadcast transaction: %v", err)
	}

	fmt.Printf("Transaction broadcasted successfully. TxID: %s\n", txID)
}
