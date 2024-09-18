# \FaucetsAPI

All URIs are relative to *https://api.hiro.so*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBtcBalance**](FaucetsAPI.md#GetBtcBalance) | **Get** /extended/v1/faucets/btc/{address} | Get BTC balance for address
[**RunFaucetBtc**](FaucetsAPI.md#RunFaucetBtc) | **Post** /extended/v1/faucets/btc | Add testnet BTC tokens to address
[**RunFaucetStx**](FaucetsAPI.md#RunFaucetStx) | **Post** /extended/v1/faucets/stx | Get STX testnet tokens



## GetBtcBalance

> GetBtcBalance200Response GetBtcBalance(ctx, address).Execute()

Get BTC balance for address

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	address := "2N4M94S1ZPt8HfxydXzL2P7qyzgVq7MHWts" // string | A valid testnet BTC address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FaucetsAPI.GetBtcBalance(context.Background(), address).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FaucetsAPI.GetBtcBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBtcBalance`: GetBtcBalance200Response
	fmt.Fprintf(os.Stdout, "Response from `FaucetsAPI.GetBtcBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | A valid testnet BTC address | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBtcBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetBtcBalance200Response**](GetBtcBalance200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunFaucetBtc

> RunFaucetResponse RunFaucetBtc(ctx).Address(address).RunFaucetBtcRequest(runFaucetBtcRequest).Execute()

Add testnet BTC tokens to address



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	address := "2N4M94S1ZPt8HfxydXzL2P7qyzgVq7MHWts" // string | A valid testnet BTC address (optional)
	runFaucetBtcRequest := *openapiclient.NewRunFaucetBtcRequest() // RunFaucetBtcRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FaucetsAPI.RunFaucetBtc(context.Background()).Address(address).RunFaucetBtcRequest(runFaucetBtcRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FaucetsAPI.RunFaucetBtc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunFaucetBtc`: RunFaucetResponse
	fmt.Fprintf(os.Stdout, "Response from `FaucetsAPI.RunFaucetBtc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunFaucetBtcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **address** | **string** | A valid testnet BTC address | 
 **runFaucetBtcRequest** | [**RunFaucetBtcRequest**](RunFaucetBtcRequest.md) |  | 

### Return type

[**RunFaucetResponse**](RunFaucetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunFaucetStx

> RunFaucetResponse1 RunFaucetStx(ctx).Address(address).Stacking(stacking).RunFaucetStxRequest(runFaucetStxRequest).Execute()

Get STX testnet tokens



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	address := "ST3M7N9Q9HDRM7RVP1Q26P0EE69358PZZAZD7KMXQ" // string | A valid testnet STX address (optional)
	stacking := true // bool | Request the amount of STX tokens needed for individual address stacking (optional) (default to false)
	runFaucetStxRequest := *openapiclient.NewRunFaucetStxRequest() // RunFaucetStxRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FaucetsAPI.RunFaucetStx(context.Background()).Address(address).Stacking(stacking).RunFaucetStxRequest(runFaucetStxRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FaucetsAPI.RunFaucetStx``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunFaucetStx`: RunFaucetResponse1
	fmt.Fprintf(os.Stdout, "Response from `FaucetsAPI.RunFaucetStx`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunFaucetStxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **address** | **string** | A valid testnet STX address | 
 **stacking** | **bool** | Request the amount of STX tokens needed for individual address stacking | [default to false]
 **runFaucetStxRequest** | [**RunFaucetStxRequest**](RunFaucetStxRequest.md) |  | 

### Return type

[**RunFaucetResponse1**](RunFaucetResponse1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

