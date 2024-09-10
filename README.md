[![codecov](https://codecov.io/gh/icon-project/stacks-go-sdk/graph/badge.svg?token=CvLaGibygJ)](https://codecov.io/gh/icon-project/stacks-go-sdk)

# Unofficial Stacks Blockchain SDK for Go

Send tokens and call Clarity smart contracts on the Stacks blockchain with Golang.

## Features

- Transaction signing for standard authorization and single-signature spending conditions
- Supports Token Transfer transactions (Type-0) and Contract Call transactions (Type-2)
- Clarity value encoding and decoding
- Stacks address handling and conversion

## Installation

To use this SDK in your Go project, run:

```bash
go get github.com/balancednetwork/stacks-go-sdk
```

## Usage
Here are some basic usage examples:

### Creating and Broadcasting a Token Transfer Transaction
```go
import (
    "github.com/balancednetwork/stacks-go-sdk/stacks"
    "math/big"
)

func main() {
    network := stacks.NewStacksMainnet()

    recipient := "SP2J6ZY48GV1EZ5V2V5RB9MP66SW86PYKKNRV9EJ7"
    amount := big.NewInt(1000000) // amount in microSTX
    memo := "Transfer memo"
    senderAddress := "SP1P72Z3704VMT3DMHPP2CB8TGQWGDBHD3RPR9GZS"
    senderKey := []byte{...}  // sender's private key

    // Create and sign the transaction
    tx, err := stacks.MakeSTXTokenTransfer(
        recipient,
        *amount,
        memo,
        *network,
        senderAddress,
        senderKey,
        nil,  // Let the function estimate the fee
        nil,  // Let the function fetch the nonce
    )
    if err != nil {
        // Handle error
    }

    // Broadcast the transaction
    txID, err := stacks.BroadcastTransaction(tx, network)
    if err != nil {
        // Handle error
    }

    println("Transaction broadcast successfully. Transaction ID:", txID)
}
```

### Creating and Broadcasting a Token Transfer Transaction
```go
import (
    "github.com/balancednetwork/stacks-go-sdk/stacks"
    "github.com/balancednetwork/stacks-go-sdk/clarity"
)

func main() {
    network := stacks.NewStacksMainnet()

    contractAddress := "SP466FNC0P7JWTNM2R9T199QRZN1MYEDTAR0KP27"
    contractName := "contract-name"
    functionName := "function-name"
    senderAddress := "SP1P72Z3704VMT3DMHPP2CB8TGQWGDBHD3RPR9GZS"
    senderKey := []byte{...}  // sender's private key

    // Prepare function arguments
    arg1, _ := clarity.NewInt(123)
    arg2 := clarity.NewStringType("example")
    functionArgs := []clarity.ClarityValue{arg1, arg2}

    // Create and sign the transaction
    tx, err := stacks.MakeContractCall(
        contractAddress,
        contractName,
        functionName,
        functionArgs,
        *network,
        senderAddress,
        senderKey,
        nil,  // Let the function estimate the fee
        nil,  // Let the function fetch the nonce
    )
    if err != nil {
        // Handle error
    }

    // Broadcast the transaction
    txID, err := stacks.BroadcastTransaction(tx, network)
    if err != nil {
        // Handle error
    }

    println("Contract call transaction broadcast successfully. Transaction ID:", txID)
}
```

### Working with Clarity Values
```golang
import (
    "github.com/balancednetwork/stacks-go-sdk/clarity"
)

func main() {
    // Create a Clarity integer
    intValue, err := clarity.NewInt(12345)
    if err != nil {
        // Handle error
    }

    // Serialize the Clarity value
    serialized, err := intValue.Serialize()
    if err != nil {
        // Handle error
    }

    // Deserialize a Clarity value
    deserialized, err := clarity.DeserializeClarityValue(serialized)
    if err != nil {
        // Handle error
    }
}
```

## Contributing
Contributions are welcome! Please feel free to submit a Pull Request.

## Disclaimer
This is an unofficial SDK and is not affiliated with or endorsed by the Stacks Foundation or Hiro PBC. Use at your own risk.

## License
This project is licensed under the MIT License.
