# \FungibleTokensAPI

All URIs are relative to *https://api.hiro.so*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFtHolders**](FungibleTokensAPI.md#GetFtHolders) | **Get** /extended/v1/tokens/ft/{token}/holders | Fungible token holders



## GetFtHolders

> GetFtHolders200Response GetFtHolders(ctx, token).Limit(limit).Offset(offset).Execute()

Fungible token holders



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
	token := "stx" // string | fungible token identifier
	limit := int32(56) // int32 | max number of holders to fetch (optional) (default to 100)
	offset := int32(56) // int32 | index of first holder to fetch (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FungibleTokensAPI.GetFtHolders(context.Background(), token).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FungibleTokensAPI.GetFtHolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFtHolders`: GetFtHolders200Response
	fmt.Fprintf(os.Stdout, "Response from `FungibleTokensAPI.GetFtHolders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | fungible token identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFtHoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | max number of holders to fetch | [default to 100]
 **offset** | **int32** | index of first holder to fetch | [default to 0]

### Return type

[**GetFtHolders200Response**](GetFtHolders200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

