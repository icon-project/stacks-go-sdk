# \ProofOfTransferAPI

All URIs are relative to *https://api.hiro.so*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPoxCycle**](ProofOfTransferAPI.md#GetPoxCycle) | **Get** /extended/v2/pox/cycles/{cycle_number} | Get PoX cycle
[**GetPoxCycleSigner**](ProofOfTransferAPI.md#GetPoxCycleSigner) | **Get** /extended/v2/pox/cycles/{cycle_number}/signers/{signer_key} | Get signer in PoX cycle
[**GetPoxCycleSignerStackers**](ProofOfTransferAPI.md#GetPoxCycleSignerStackers) | **Get** /extended/v2/pox/cycles/{cycle_number}/signers/{signer_key}/stackers | Get stackers for signer in PoX cycle
[**GetPoxCycleSigners**](ProofOfTransferAPI.md#GetPoxCycleSigners) | **Get** /extended/v2/pox/cycles/{cycle_number}/signers | Get signers in PoX cycle
[**GetPoxCycles**](ProofOfTransferAPI.md#GetPoxCycles) | **Get** /extended/v2/pox/cycles | Get PoX cycles



## GetPoxCycle

> GetPoxCycles200ResponseResultsInner GetPoxCycle(ctx, cycleNumber).Execute()

Get PoX cycle



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
	cycleNumber := int32(56) // int32 | PoX cycle number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProofOfTransferAPI.GetPoxCycle(context.Background(), cycleNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProofOfTransferAPI.GetPoxCycle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPoxCycle`: GetPoxCycles200ResponseResultsInner
	fmt.Fprintf(os.Stdout, "Response from `ProofOfTransferAPI.GetPoxCycle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cycleNumber** | **int32** | PoX cycle number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoxCycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPoxCycles200ResponseResultsInner**](GetPoxCycles200ResponseResultsInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoxCycleSigner

> GetPoxCycleSigners200ResponseResultsInner GetPoxCycleSigner(ctx, cycleNumber, signerKey).Execute()

Get signer in PoX cycle



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
	cycleNumber := int32(56) // int32 | PoX cycle number
	signerKey := "0x038e3c4529395611be9abf6fa3b6987e81d402385e3d605a073f42f407565a4a3d" // string | Signer key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProofOfTransferAPI.GetPoxCycleSigner(context.Background(), cycleNumber, signerKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProofOfTransferAPI.GetPoxCycleSigner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPoxCycleSigner`: GetPoxCycleSigners200ResponseResultsInner
	fmt.Fprintf(os.Stdout, "Response from `ProofOfTransferAPI.GetPoxCycleSigner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cycleNumber** | **int32** | PoX cycle number | 
**signerKey** | **string** | Signer key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoxCycleSignerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetPoxCycleSigners200ResponseResultsInner**](GetPoxCycleSigners200ResponseResultsInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoxCycleSignerStackers

> GetPoxCycleSignerStackers200Response GetPoxCycleSignerStackers(ctx, cycleNumber, signerKey).Limit(limit).Offset(offset).Execute()

Get stackers for signer in PoX cycle



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
	cycleNumber := int32(56) // int32 | PoX cycle number
	signerKey := "0x038e3c4529395611be9abf6fa3b6987e81d402385e3d605a073f42f407565a4a3d" // string | Signer key
	limit := int32(56) // int32 | Results per page (optional) (default to 100)
	offset := int32(56) // int32 | Result offset (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProofOfTransferAPI.GetPoxCycleSignerStackers(context.Background(), cycleNumber, signerKey).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProofOfTransferAPI.GetPoxCycleSignerStackers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPoxCycleSignerStackers`: GetPoxCycleSignerStackers200Response
	fmt.Fprintf(os.Stdout, "Response from `ProofOfTransferAPI.GetPoxCycleSignerStackers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cycleNumber** | **int32** | PoX cycle number | 
**signerKey** | **string** | Signer key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoxCycleSignerStackersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Results per page | [default to 100]
 **offset** | **int32** | Result offset | [default to 0]

### Return type

[**GetPoxCycleSignerStackers200Response**](GetPoxCycleSignerStackers200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoxCycleSigners

> GetPoxCycleSigners200Response GetPoxCycleSigners(ctx, cycleNumber).Limit(limit).Offset(offset).Execute()

Get signers in PoX cycle



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
	cycleNumber := int32(56) // int32 | PoX cycle number
	limit := int32(56) // int32 | Results per page (optional) (default to 100)
	offset := int32(56) // int32 | Result offset (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProofOfTransferAPI.GetPoxCycleSigners(context.Background(), cycleNumber).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProofOfTransferAPI.GetPoxCycleSigners``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPoxCycleSigners`: GetPoxCycleSigners200Response
	fmt.Fprintf(os.Stdout, "Response from `ProofOfTransferAPI.GetPoxCycleSigners`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cycleNumber** | **int32** | PoX cycle number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoxCycleSignersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Results per page | [default to 100]
 **offset** | **int32** | Result offset | [default to 0]

### Return type

[**GetPoxCycleSigners200Response**](GetPoxCycleSigners200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoxCycles

> GetPoxCycles200Response GetPoxCycles(ctx).Limit(limit).Offset(offset).Execute()

Get PoX cycles



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
	resp, r, err := apiClient.ProofOfTransferAPI.GetPoxCycles(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProofOfTransferAPI.GetPoxCycles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPoxCycles`: GetPoxCycles200Response
	fmt.Fprintf(os.Stdout, "Response from `ProofOfTransferAPI.GetPoxCycles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPoxCyclesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Results per page | [default to 20]
 **offset** | **int32** | Result offset | [default to 0]

### Return type

[**GetPoxCycles200Response**](GetPoxCycles200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

