/*
Stacks Blockchain API

Testing AccountsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package stacks_blockchain_api_client

import (
	"context"
	"testing"

	"github.com/icon-project/stacks-go-sdk/pkg/stacks_blockchain_api_client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_stacks_blockchain_api_client_AccountsAPIService(t *testing.T) {

	configuration := stacks_blockchain_api_client.NewConfiguration()
	configuration.Servers = stacks_blockchain_api_client.ServerConfigurations{configuration.Servers[1]}
	apiClient := stacks_blockchain_api_client.NewAPIClient(configuration)
	testAddress := "ST15C893XJFJ6FSKM020P9JQDB5T7X6MQTXMBPAVH"
	principal := stacks_blockchain_api_client.GetFilteredEventsAddressParameter{
		String: &testAddress,
	}

	t.Run("Test AccountsAPIService GetAccountAssets", func(t *testing.T) {
		resp, httpRes, err := apiClient.AccountsAPI.GetAccountAssets(context.Background(), principal).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountsAPIService GetAccountBalance", func(t *testing.T) {
		resp, httpRes, err := apiClient.AccountsAPI.GetAccountBalance(context.Background(), principal).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountsAPIService GetAccountInbound", func(t *testing.T) {
		resp, httpRes, err := apiClient.AccountsAPI.GetAccountInbound(context.Background(), principal).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountsAPIService GetAccountNonces", func(t *testing.T) {
		resp, httpRes, err := apiClient.AccountsAPI.GetAccountNonces(context.Background(), principal).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountsAPIService GetAccountStxBalance", func(t *testing.T) {
		resp, httpRes, err := apiClient.AccountsAPI.GetAccountStxBalance(context.Background(), principal).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountsAPIService GetAccountTransactions", func(t *testing.T) {
		resp, httpRes, err := apiClient.AccountsAPI.GetAccountTransactions(context.Background(), principal).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountsAPIService GetAccountTransactionsWithTransfers", func(t *testing.T) {
		resp, httpRes, err := apiClient.AccountsAPI.GetAccountTransactionsWithTransfers(context.Background(), principal).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountsAPIService GetSingleTransactionWithTransfers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var principal stacks_blockchain_api_client.GetFilteredEventsAddressParameter
		var txId string

		resp, httpRes, err := apiClient.AccountsAPI.GetSingleTransactionWithTransfers(context.Background(), principal, txId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}