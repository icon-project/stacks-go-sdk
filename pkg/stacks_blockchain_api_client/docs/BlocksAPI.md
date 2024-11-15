# \BlocksAPI

All URIs are relative to *https://api.hiro.so*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAverageBlockTimes**](BlocksAPI.md#GetAverageBlockTimes) | **Get** /extended/v2/blocks/average-times | Get average block times
[**GetBlock**](BlocksAPI.md#GetBlock) | **Get** /extended/v2/blocks/{height_or_hash} | Get block
[**GetBlockByBurnBlockHash**](BlocksAPI.md#GetBlockByBurnBlockHash) | **Get** /extended/v1/block/by_burn_block_hash/{burn_block_hash} | Get block by burnchain block hash
[**GetBlockByBurnBlockHeight**](BlocksAPI.md#GetBlockByBurnBlockHeight) | **Get** /extended/v1/block/by_burn_block_height/{burn_block_height} | Get block by burnchain height
[**GetBlockByHash**](BlocksAPI.md#GetBlockByHash) | **Get** /extended/v1/block/{hash} | Get block by hash
[**GetBlockByHeight**](BlocksAPI.md#GetBlockByHeight) | **Get** /extended/v1/block/by_height/{height} | Get block by height
[**GetBlockList**](BlocksAPI.md#GetBlockList) | **Get** /extended/v1/block/ | Get recent blocks
[**GetBlocks**](BlocksAPI.md#GetBlocks) | **Get** /extended/v2/blocks/ | Get blocks
[**GetSignerSignaturesForBlock**](BlocksAPI.md#GetSignerSignaturesForBlock) | **Get** /extended/v2/blocks/{height_or_hash}/signer-signatures | Get signer signatures for block



## GetAverageBlockTimes

> GetAverageBlockTimes200Response GetAverageBlockTimes(ctx).Execute()

Get average block times



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
	resp, r, err := apiClient.BlocksAPI.GetAverageBlockTimes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocksAPI.GetAverageBlockTimes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAverageBlockTimes`: GetAverageBlockTimes200Response
	fmt.Fprintf(os.Stdout, "Response from `BlocksAPI.GetAverageBlockTimes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAverageBlockTimesRequest struct via the builder pattern


### Return type

[**GetAverageBlockTimes200Response**](GetAverageBlockTimes200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlock

> GetBlocks200ResponseResultsInner GetBlock(ctx, heightOrHash).Execute()

Get block



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
	heightOrHash := *openapiclient.NewGetBlockHeightOrHashParameter() // GetBlockHeightOrHashParameter | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlocksAPI.GetBlock(context.Background(), heightOrHash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocksAPI.GetBlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlock`: GetBlocks200ResponseResultsInner
	fmt.Fprintf(os.Stdout, "Response from `BlocksAPI.GetBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**heightOrHash** | [**GetBlockHeightOrHashParameter**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetBlocks200ResponseResultsInner**](GetBlocks200ResponseResultsInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockByBurnBlockHash

> Block GetBlockByBurnBlockHash(ctx, burnBlockHash).Execute()

Get block by burnchain block hash



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
	burnBlockHash := "0x00000000000000000002bba732926cf68b6eda3e2cdbc2a85af79f10efeeeb10" // string | Hash of the burnchain block

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlocksAPI.GetBlockByBurnBlockHash(context.Background(), burnBlockHash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocksAPI.GetBlockByBurnBlockHash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockByBurnBlockHash`: Block
	fmt.Fprintf(os.Stdout, "Response from `BlocksAPI.GetBlockByBurnBlockHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**burnBlockHash** | **string** | Hash of the burnchain block | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockByBurnBlockHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Block**](Block.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockByBurnBlockHeight

> Block GetBlockByBurnBlockHeight(ctx, burnBlockHeight).Execute()

Get block by burnchain height



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
	burnBlockHeight := int32(744603) // int32 | Height of the burn chain block

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlocksAPI.GetBlockByBurnBlockHeight(context.Background(), burnBlockHeight).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocksAPI.GetBlockByBurnBlockHeight``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockByBurnBlockHeight`: Block
	fmt.Fprintf(os.Stdout, "Response from `BlocksAPI.GetBlockByBurnBlockHeight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**burnBlockHeight** | **int32** | Height of the burn chain block | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockByBurnBlockHeightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Block**](Block.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockByHash

> Block GetBlockByHash(ctx, hash).Execute()

Get block by hash



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
	hash := "0x4839a8b01cfb39ffcc0d07d3db31e848d5adf5279d529ed5062300b9f353ff79" // string | Hash of the block

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlocksAPI.GetBlockByHash(context.Background(), hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocksAPI.GetBlockByHash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockByHash`: Block
	fmt.Fprintf(os.Stdout, "Response from `BlocksAPI.GetBlockByHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | Hash of the block | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockByHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Block**](Block.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockByHeight

> Block GetBlockByHeight(ctx, height).Execute()

Get block by height



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
	height := int32(10000) // int32 | Height of the block

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlocksAPI.GetBlockByHeight(context.Background(), height).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocksAPI.GetBlockByHeight``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockByHeight`: Block
	fmt.Fprintf(os.Stdout, "Response from `BlocksAPI.GetBlockByHeight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**height** | **int32** | Height of the block | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockByHeightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Block**](Block.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockList

> BlockListResponse GetBlockList(ctx).Limit(limit).Offset(offset).Execute()

Get recent blocks



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
	limit := int32(56) // int32 | max number of blocks to fetch (optional) (default to 20)
	offset := int32(56) // int32 | Result offset (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlocksAPI.GetBlockList(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocksAPI.GetBlockList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockList`: BlockListResponse
	fmt.Fprintf(os.Stdout, "Response from `BlocksAPI.GetBlockList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | max number of blocks to fetch | [default to 20]
 **offset** | **int32** | Result offset | [default to 0]

### Return type

[**BlockListResponse**](BlockListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlocks

> GetBlocks200Response GetBlocks(ctx).Limit(limit).Offset(offset).Cursor(cursor).Execute()

Get blocks



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
	cursor := "cursor_example" // string | Cursor for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlocksAPI.GetBlocks(context.Background()).Limit(limit).Offset(offset).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocksAPI.GetBlocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlocks`: GetBlocks200Response
	fmt.Fprintf(os.Stdout, "Response from `BlocksAPI.GetBlocks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Results per page | [default to 20]
 **offset** | **int32** | Result offset | [default to 0]
 **cursor** | **string** | Cursor for pagination | 

### Return type

[**GetBlocks200Response**](GetBlocks200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSignerSignaturesForBlock

> GetSignerSignaturesForBlock200Response GetSignerSignaturesForBlock(ctx, heightOrHash).Limit(limit).Offset(offset).Execute()

Get signer signatures for block



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
	heightOrHash := *openapiclient.NewGetBlockHeightOrHashParameter() // GetBlockHeightOrHashParameter | 
	limit := int32(56) // int32 | Results per page (optional) (default to 500)
	offset := int32(56) // int32 | Result offset (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlocksAPI.GetSignerSignaturesForBlock(context.Background(), heightOrHash).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocksAPI.GetSignerSignaturesForBlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSignerSignaturesForBlock`: GetSignerSignaturesForBlock200Response
	fmt.Fprintf(os.Stdout, "Response from `BlocksAPI.GetSignerSignaturesForBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**heightOrHash** | [**GetBlockHeightOrHashParameter**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSignerSignaturesForBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Results per page | [default to 500]
 **offset** | **int32** | Result offset | [default to 0]

### Return type

[**GetSignerSignaturesForBlock200Response**](GetSignerSignaturesForBlock200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

