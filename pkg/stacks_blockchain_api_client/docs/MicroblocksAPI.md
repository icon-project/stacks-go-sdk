# \MicroblocksAPI

All URIs are relative to *https://api.hiro.so*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMicroblockByHash**](MicroblocksAPI.md#GetMicroblockByHash) | **Get** /extended/v1/microblock/{hash} | Get microblock
[**GetMicroblockList**](MicroblocksAPI.md#GetMicroblockList) | **Get** /extended/v1/microblock/ | Get recent microblocks
[**GetUnanchoredTxs**](MicroblocksAPI.md#GetUnanchoredTxs) | **Get** /extended/v1/microblock/unanchored/txs | Get the list of current transactions that belong to unanchored microblocks



## GetMicroblockByHash

> Microblock GetMicroblockByHash(ctx, hash).Execute()

Get microblock



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
	hash := "0x3bfcdf84b3012adb544cf0f6df4835f93418c2269a3881885e27b3d58eb82d47" // string | Hash of the microblock

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MicroblocksAPI.GetMicroblockByHash(context.Background(), hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MicroblocksAPI.GetMicroblockByHash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMicroblockByHash`: Microblock
	fmt.Fprintf(os.Stdout, "Response from `MicroblocksAPI.GetMicroblockByHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | Hash of the microblock | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMicroblockByHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Microblock**](Microblock.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMicroblockList

> MicroblockListResponse GetMicroblockList(ctx).Limit(limit).Offset(offset).Execute()

Get recent microblocks



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
	limit := int32(56) // int32 | Max number of microblocks to fetch (optional) (default to 20)
	offset := int32(56) // int32 | Result offset (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MicroblocksAPI.GetMicroblockList(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MicroblocksAPI.GetMicroblockList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMicroblockList`: MicroblockListResponse
	fmt.Fprintf(os.Stdout, "Response from `MicroblocksAPI.GetMicroblockList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMicroblockListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of microblocks to fetch | [default to 20]
 **offset** | **int32** | Result offset | [default to 0]

### Return type

[**MicroblockListResponse**](MicroblockListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUnanchoredTxs

> GetUnanchoredTxs200Response GetUnanchoredTxs(ctx).Execute()

Get the list of current transactions that belong to unanchored microblocks



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MicroblocksAPI.GetUnanchoredTxs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MicroblocksAPI.GetUnanchoredTxs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUnanchoredTxs`: GetUnanchoredTxs200Response
	fmt.Fprintf(os.Stdout, "Response from `MicroblocksAPI.GetUnanchoredTxs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUnanchoredTxsRequest struct via the builder pattern


### Return type

[**GetUnanchoredTxs200Response**](GetUnanchoredTxs200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

