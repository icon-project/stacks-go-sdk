# \StackingAPI

All URIs are relative to *https://api.hiro.so*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExtendedV1PoxEventsGet**](StackingAPI.md#ExtendedV1PoxEventsGet) | **Get** /extended/v1/{pox}/events | Get latest PoX events
[**ExtendedV1PoxStackerPrincipalGet**](StackingAPI.md#ExtendedV1PoxStackerPrincipalGet) | **Get** /extended/v1/{pox}/stacker/{principal} | Get events for a stacking address
[**ExtendedV1PoxTxTxIdGet**](StackingAPI.md#ExtendedV1PoxTxTxIdGet) | **Get** /extended/v1/{pox}/tx/{tx_id} | Get PoX events for a transaction
[**GetPoolDelegations**](StackingAPI.md#GetPoolDelegations) | **Get** /extended/v1/{pox}/{pool_principal}/delegations | Stacking pool members



## ExtendedV1PoxEventsGet

> ExtendedV1PoxEventsGet(ctx, pox).Limit(limit).Offset(offset).Execute()

Get latest PoX events

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
	pox := *openapiclient.NewExtendedV1PoxEventsGetPoxParameter() // ExtendedV1PoxEventsGetPoxParameter | 
	limit := int32(56) // int32 | Results per page (optional) (default to 96)
	offset := int32(56) // int32 | Result offset (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StackingAPI.ExtendedV1PoxEventsGet(context.Background(), pox).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackingAPI.ExtendedV1PoxEventsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pox** | [**ExtendedV1PoxEventsGetPoxParameter**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtendedV1PoxEventsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Results per page | [default to 96]
 **offset** | **int32** | Result offset | [default to 0]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExtendedV1PoxStackerPrincipalGet

> ExtendedV1PoxStackerPrincipalGet(ctx, pox, principal).Execute()

Get events for a stacking address

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
	pox := *openapiclient.NewExtendedV1PoxEventsGetPoxParameter() // ExtendedV1PoxEventsGetPoxParameter | 
	principal := *openapiclient.NewGetFilteredEventsAddressParameter() // GetFilteredEventsAddressParameter | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StackingAPI.ExtendedV1PoxStackerPrincipalGet(context.Background(), pox, principal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackingAPI.ExtendedV1PoxStackerPrincipalGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pox** | [**ExtendedV1PoxEventsGetPoxParameter**](.md) |  | 
**principal** | [**GetFilteredEventsAddressParameter**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtendedV1PoxStackerPrincipalGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExtendedV1PoxTxTxIdGet

> ExtendedV1PoxTxTxIdGet(ctx, pox, txId).Execute()

Get PoX events for a transaction

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
	pox := *openapiclient.NewExtendedV1PoxEventsGetPoxParameter() // ExtendedV1PoxEventsGetPoxParameter | 
	txId := "txId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StackingAPI.ExtendedV1PoxTxTxIdGet(context.Background(), pox, txId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackingAPI.ExtendedV1PoxTxTxIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pox** | [**ExtendedV1PoxEventsGetPoxParameter**](.md) |  | 
**txId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtendedV1PoxTxTxIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoolDelegations

> GetPoolDelegations200Response GetPoolDelegations(ctx, pox, poolPrincipal).Limit(limit).Offset(offset).AfterBlock(afterBlock).Height(height).Unanchored(unanchored).Execute()

Stacking pool members



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
	pox := *openapiclient.NewExtendedV1PoxEventsGetPoxParameter() // ExtendedV1PoxEventsGetPoxParameter | 
	poolPrincipal := "SPSCWDV3RKV5ZRN1FQD84YE1NQFEDJ9R1F4DYQ11" // string | Address principal of the stacking pool delegator
	limit := int32(56) // int32 | Results per page (optional) (default to 100)
	offset := int32(56) // int32 | Result offset (optional) (default to 0)
	afterBlock := int32(56) // int32 | If specified, only delegation events after the given block will be included (optional)
	height := int32(56) // int32 |  (optional)
	unanchored := true // bool | Include data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackingAPI.GetPoolDelegations(context.Background(), pox, poolPrincipal).Limit(limit).Offset(offset).AfterBlock(afterBlock).Height(height).Unanchored(unanchored).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackingAPI.GetPoolDelegations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPoolDelegations`: GetPoolDelegations200Response
	fmt.Fprintf(os.Stdout, "Response from `StackingAPI.GetPoolDelegations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pox** | [**ExtendedV1PoxEventsGetPoxParameter**](.md) |  | 
**poolPrincipal** | **string** | Address principal of the stacking pool delegator | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoolDelegationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Results per page | [default to 100]
 **offset** | **int32** | Result offset | [default to 0]
 **afterBlock** | **int32** | If specified, only delegation events after the given block will be included | 
 **height** | **int32** |  | 
 **unanchored** | **bool** | Include data from unanchored (i.e. unconfirmed) microblocks | [default to false]

### Return type

[**GetPoolDelegations200Response**](GetPoolDelegations200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

