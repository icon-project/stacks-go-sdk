package stacks

import (
	"fmt"
	"net/http"
)

type StacksNetwork struct {
	CoreAPIURL string
	Version    TransactionVersion
	ChainID    ChainID
}

func NewStacksMainnet() *StacksNetwork {
	return &StacksNetwork{
		CoreAPIURL: "https://api.mainnet.hiro.so",
		Version:    TransactionVersionMainnet,
		ChainID:    ChainIDMainnet,
	}
}

func NewStacksTestnet() *StacksNetwork {
	return &StacksNetwork{
		CoreAPIURL: "https://api.testnet.hiro.so",
		Version:    TransactionVersionTestnet,
		ChainID:    ChainIDTestnet,
	}
}

func (n *StacksNetwork) GetAccountAPIURL(address string) string {
	return fmt.Sprintf("%s/v2/accounts/%s?proof=0", n.CoreAPIURL, address)
}

func (n *StacksNetwork) GetBroadcastAPIURL() string {
	return fmt.Sprintf("%s/v2/transactions", n.CoreAPIURL)
}

func (n *StacksNetwork) GetTransferFeeEstimateAPIURL() string {
	return fmt.Sprintf("%s/v2/fees/transfer", n.CoreAPIURL)
}

func (n *StacksNetwork) GetTransactionFeeEstimateAPIURL() string {
	return fmt.Sprintf("%s/v2/fees/transaction", n.CoreAPIURL)
}

func (n *StacksNetwork) GetNonceAPIURL(address string) string {
	return fmt.Sprintf("%s/extended/v1/address/%s/nonces", n.CoreAPIURL, address)
}

func (n *StacksNetwork) FetchFn(url string) (*http.Response, error) {
	return http.Get(url)
}
