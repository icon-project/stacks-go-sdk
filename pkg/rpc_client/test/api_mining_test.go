/*
Stacks 2.0+ RPC API

Testing MiningAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package rpc_client

import (
	"context"
	"testing"

	"github.com/icon-project/stacks-go-sdk/pkg/rpc_client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_rpc_client_MiningAPIService(t *testing.T) {

	configuration := rpc_client.NewConfiguration()
	apiClient := rpc_client.NewAPIClient(configuration)

	t.Run("Test MiningAPIService GetStackerSet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var cycleNumber int32

		httpRes, err := apiClient.MiningAPI.GetStackerSet(context.Background(), cycleNumber).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MiningAPIService PostBlockProposal", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.MiningAPI.PostBlockProposal(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
