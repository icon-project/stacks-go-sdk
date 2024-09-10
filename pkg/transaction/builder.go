package transaction

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/icon-project/stacks-go-sdk/pkg/clarity"
	"github.com/icon-project/stacks-go-sdk/pkg/crypto"
	"github.com/icon-project/stacks-go-sdk/pkg/stacks"
)

type CustomError struct {
	Message string
	Err     error
}

func (e *CustomError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func makeRequest(url string, method string, payload interface{}) ([]byte, error) {
	client := resty.New()
	var resp *resty.Response
	var err error

	switch method {
	case "GET":
		resp, err = client.R().Get(url)
	case "POST":
		resp, err = client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(payload).
			Post(url)
	default:
		return nil, &CustomError{Message: "Unsupported HTTP method"}
	}

	if err != nil {
		return nil, &CustomError{Message: "Error making request", Err: err}
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, &CustomError{Message: fmt.Sprintf("Request failed with status code: %d, body: %s", resp.StatusCode(), string(resp.Body()))}
	}

	return resp.Body(), nil
}

func getNonce(address string, network stacks.StacksNetwork) (*big.Int, error) {
	url := network.GetAccountAPIURL(address)
	body, err := makeRequest(url, "GET", nil)
	if err != nil {
		return nil, &CustomError{Message: "Error fetching nonce", Err: err}
	}

	var result struct {
		Nonce uint64 `json:"nonce"`
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, &CustomError{Message: "Error parsing JSON response", Err: err}
	}

	return big.NewInt(int64(result.Nonce)), nil
}

func estimateTransactionFee(tx StacksTransaction, network stacks.StacksNetwork) (*big.Int, error) {
	url := network.GetTransactionFeeEstimateAPIURL()
	serializedTxPayload, err := tx.GetPayload().Serialize()
	if err != nil {
		return nil, &CustomError{Message: "Error serializing transaction payload", Err: err}
	}

	serializedTx, err := tx.Serialize()
	if err != nil {
		return nil, fmt.Errorf("error serializing entire transaction: %w", err)
	}

	payload := map[string]interface{}{
		"transaction_payload": hex.EncodeToString(serializedTxPayload),
		"estimated_len":       len(serializedTx),
	}

	body, err := makeRequest(url, "POST", payload)
	if err != nil {
		if strings.Contains(err.Error(), "NoEstimateAvailable") {
			return estimateTransferUnsafe(tx, network)
		}
		return nil, &CustomError{Message: "Error estimating fee", Err: err}
	}

	var result struct {
		Estimations []struct {
			FeeRate float64 `json:"fee_rate"`
			Fee     int64   `json:"fee"`
		} `json:"estimations"`
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, &CustomError{Message: "Error parsing JSON response", Err: err}
	}

	if len(result.Estimations) < 2 {
		return nil, &CustomError{Message: "Insufficient estimations in response"}
	}

	fee := big.NewInt(result.Estimations[1].Fee)

	return fee, nil
}

func estimateTransferUnsafe(tx StacksTransaction, network stacks.StacksNetwork) (*big.Int, error) {
	url := network.GetTransferFeeEstimateAPIURL()
	body, err := makeRequest(url, "GET", nil)
	if err != nil {
		return nil, &CustomError{Message: "Error fetching unsafe fee estimate", Err: err}
	}

	feeRate, err := strconv.ParseInt(string(body), 10, 64)
	if err != nil {
		return nil, &CustomError{Message: "Error parsing unsafe fee estimate", Err: err}
	}

	serializedTx, err := tx.Serialize()
	if err != nil {
		return nil, &CustomError{Message: "Error serializing transaction", Err: err}
	}

	txBytes := big.NewInt(int64(len(serializedTx)))
	fee := big.NewInt(feeRate)
	fee.Mul(fee, txBytes)

	return fee, nil
}

func createAndSignTransaction(tx StacksTransaction, network stacks.StacksNetwork, senderAddress string, senderKey []byte, fee *big.Int, nonce *big.Int) error {
	var err error
	if fee == nil {
		fee, err = estimateTransactionFee(tx, network)
		if err != nil {
			return &CustomError{Message: "Failed to estimate fee", Err: err}
		}
	}

	if nonce == nil {
		nonce, err = getNonce(senderAddress, network)
		if err != nil {
			return &CustomError{Message: "Failed to fetch nonce", Err: err}
		}
	}

	auth := tx.GetAuth()
	if auth == nil {
		return &CustomError{Message: "Transaction authentication is nil"}
	}

	auth.OriginAuth.Fee = fee.Uint64()
	auth.OriginAuth.Nonce = nonce.Uint64()

	err = SignTransaction(tx, senderKey)
	if err != nil {
		return &CustomError{Message: "Failed to sign transaction", Err: err}
	}

	return nil
}

func deriveSigner(senderKey []byte) [20]byte {
	senderPublicKey := crypto.GetPublicKeyFromPrivate(senderKey)
	var signer [20]byte
	copy(signer[:], crypto.Hash160(senderPublicKey))
	return signer
}

func MakeSTXTokenTransfer(
	recipient string,
	amount big.Int,
	memo string,
	network stacks.StacksNetwork,
	senderAddress string,
	senderKey []byte,
	fee *big.Int,
	nonce *big.Int,
) (*TokenTransferTransaction, error) {
	if recipient == "" || len(senderKey) == 0 {
		return nil, &CustomError{Message: "Invalid parameters: recipient or senderKey are empty"}
	}

	signer := deriveSigner(senderKey)

	tx, err := NewTokenTransferTransaction(recipient, amount.Uint64(), memo, network.Version, network.ChainID, signer, 0, 0, stacks.AnchorModeOnChainOnly, stacks.PostConditionModeDeny)
	if err != nil {
		return nil, &CustomError{Message: "Failed to create transaction", Err: err}
	}

	err = createAndSignTransaction(tx, network, senderAddress, senderKey, fee, nonce)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func MakeContractCall(
	contractAddress string,
	contractName string,
	functionName string,
	functionArgs []clarity.ClarityValue,
	network stacks.StacksNetwork,
	senderAddress string,
	senderKey []byte,
	fee *big.Int,
	nonce *big.Int,
) (*ContractCallTransaction, error) {
	if contractAddress == "" || contractName == "" || functionName == "" || len(senderKey) == 0 {
		return nil, &CustomError{Message: "Invalid parameters: contractAddress, contractName, functionName, or senderKey are empty"}
	}

	signer := deriveSigner(senderKey)

	tx, err := NewContractCallTransaction(contractAddress, contractName, functionName, functionArgs, network.Version, network.ChainID, signer, 0, 0, stacks.AnchorModeOnChainOnly, stacks.PostConditionModeDeny)
	if err != nil {
		return nil, &CustomError{Message: "Failed to create transaction", Err: err}
	}

	err = createAndSignTransaction(tx, network, senderAddress, senderKey, fee, nonce)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func BroadcastTransaction(tx StacksTransaction, network *stacks.StacksNetwork) (string, error) {
	serializedTx, err := tx.Serialize()
	if err != nil {
		return "", &CustomError{Message: "Failed to serialize transaction", Err: err}
	}

	url := network.GetBroadcastAPIURL()
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/octet-stream").
		SetBody(serializedTx).
		Post(url)
	if err != nil {
		return "", &CustomError{Message: "Failed to send transaction", Err: err}
	}

	text := string(resp.Body())

	if resp.StatusCode() != http.StatusOK {
		return "", &CustomError{Message: fmt.Sprintf("Transaction submission failed with status code: %d, body: %s", resp.StatusCode(), text)}
	}

	txId := strings.Trim(text, "\"")

	if !isValidTransactionID(txId) {
		return "", &CustomError{Message: fmt.Sprintf("Received invalid transaction ID: %s", txId)}
	}

	return txId, nil
}

func isValidTransactionID(txID string) bool {
	if len(txID) != 64 {
		return false
	}

	_, err := hex.DecodeString(txID)
	return err == nil
}
