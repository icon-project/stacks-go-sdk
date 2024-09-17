# \BurnBlocksAPI

All URIs are relative to *https://api.hiro.so*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBlocksByBurnBlock**](BurnBlocksAPI.md#GetBlocksByBurnBlock) | **Get** /extended/v2/burn-blocks/{height_or_hash}/blocks | Get blocks by burn block
[**GetBurnBlock**](BurnBlocksAPI.md#GetBurnBlock) | **Get** /extended/v2/burn-blocks/{height_or_hash} | Get burn block
[**GetBurnBlocks**](BurnBlocksAPI.md#GetBurnBlocks) | **Get** /extended/v2/burn-blocks/ | Get burn blocks



## GetBlocksByBurnBlock

> GetBlocksByBurnBlock200Response GetBlocksByBurnBlock(ctx, heightOrHash).Limit(limit).Offset(offset).Execute()

Get blocks by burn block



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
	heightOrHash := *openapiclient.NewGetBurnBlockHeightOrHashParameter() // GetBurnBlockHeightOrHashParameter | 
	limit := int32(56) // int32 | Results per page (optional) (default to 20)
	offset := int32(56) // int32 | Result offset (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BurnBlocksAPI.GetBlocksByBurnBlock(context.Background(), heightOrHash).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BurnBlocksAPI.GetBlocksByBurnBlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlocksByBurnBlock`: GetBlocksByBurnBlock200Response
	fmt.Fprintf(os.Stdout, "Response from `BurnBlocksAPI.GetBlocksByBurnBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**heightOrHash** | [**GetBurnBlockHeightOrHashParameter**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlocksByBurnBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Results per page | [default to 20]
 **offset** | **int32** | Result offset | [default to 0]

### Return type

[**GetBlocksByBurnBlock200Response**](GetBlocksByBurnBlock200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBurnBlock

> GetBurnBlocks200ResponseResultsInner GetBurnBlock(ctx, heightOrHash).Execute()

Get burn block



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
	heightOrHash := *openapiclient.NewGetBurnBlockHeightOrHashParameter() // GetBurnBlockHeightOrHashParameter | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BurnBlocksAPI.GetBurnBlock(context.Background(), heightOrHash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BurnBlocksAPI.GetBurnBlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBurnBlock`: GetBurnBlocks200ResponseResultsInner
	fmt.Fprintf(os.Stdout, "Response from `BurnBlocksAPI.GetBurnBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**heightOrHash** | [**GetBurnBlockHeightOrHashParameter**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBurnBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetBurnBlocks200ResponseResultsInner**](GetBurnBlocks200ResponseResultsInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBurnBlocks

> GetBurnBlocks200Response GetBurnBlocks(ctx).Limit(limit).Offset(offset).Execute()

Get burn blocks



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
	limit := int32(56) // int32 | Results per page (optional) (default to 20)
	offset := int32(56) // int32 | Result offset (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BurnBlocksAPI.GetBurnBlocks(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BurnBlocksAPI.GetBurnBlocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBurnBlocks`: GetBurnBlocks200Response
	fmt.Fprintf(os.Stdout, "Response from `BurnBlocksAPI.GetBurnBlocks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBurnBlocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Results per page | [default to 20]
 **offset** | **int32** | Result offset | [default to 0]

### Return type

[**GetBurnBlocks200Response**](GetBurnBlocks200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

