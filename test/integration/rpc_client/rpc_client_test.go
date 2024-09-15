package main

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/icon-project/stacks-go-sdk/pkg/rpc_client"
)

const (
	stacksNodeURL = "https://stacks-node-api.testnet.stacks.co"
)

func TestStacksRPCIntegration(t *testing.T) {
	// Create a new client with responses
	client, err := rpc_client.NewClientWithResponses(stacksNodeURL)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Test GetCoreApiInfo
	t.Run("GetCoreApiInfo", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		resp, err := client.GetCoreApiInfoWithResponse(ctx)
		if err != nil {
			t.Fatalf("GetCoreApiInfo failed: %v", err)
		}

		if resp.StatusCode() != 200 {
			t.Errorf("Expected status code 200, got %d", resp.StatusCode())
		}

		if resp.JSON200 == nil {
			t.Fatalf("Expected non-nil JSON200 response")
		}

		fmt.Printf("Core API Info: %+v\n", resp.JSON200)
	})

	// Test GetAccountInfo
	t.Run("GetAccountInfo", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Use a known Stacks address for testing
		testAddress := "SP2J6ZY48GV1EZ5V2V5RB9MP66SW86PYKKNRV9EJ7"

		resp, err := client.GetAccountInfoWithResponse(ctx, testAddress, nil)
		if err != nil {
			t.Fatalf("GetAccountInfo failed: %v", err)
		}

		if resp.StatusCode() != 200 {
			t.Errorf("Expected status code 200, got %d", resp.StatusCode())
		}

		if resp.JSON200 == nil {
			t.Fatalf("Expected non-nil JSON200 response")
		}

		fmt.Printf("Account Info: %+v\n", resp.JSON200)
	})

	// Test GetPoxInfo
	t.Run("GetPoxInfo", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		resp, err := client.GetPoxInfoWithResponse(ctx, nil)
		if err != nil {
			t.Fatalf("GetPoxInfo failed: %v", err)
		}

		if resp.StatusCode() != 200 {
			t.Errorf("Expected status code 200, got %d", resp.StatusCode())
		}

		if resp.JSON200 == nil {
			t.Fatalf("Expected non-nil JSON200 response")
		}

		fmt.Printf("PoX Info: %+v\n", resp.JSON200)
	})

	// Add more test cases for other endpoints as needed
}

func main() {
	log.Println("Running Stacks RPC Integration Tests...")
	testing.Main(func(pat, str string) (bool, error) { return true, nil },
		[]testing.InternalTest{{Name: "TestStacksRPCIntegration", F: TestStacksRPCIntegration}},
		nil,
		nil)
}
