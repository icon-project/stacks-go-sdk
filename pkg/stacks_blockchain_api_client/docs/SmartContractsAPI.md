# \SmartContractsAPI

All URIs are relative to *https://api.hiro.so*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContractById**](SmartContractsAPI.md#GetContractById) | **Get** /extended/v1/contract/{contract_id} | Get contract info
[**GetContractEventsById**](SmartContractsAPI.md#GetContractEventsById) | **Get** /extended/v1/contract/{contract_id}/events | Get contract events
[**GetContractsByTrait**](SmartContractsAPI.md#GetContractsByTrait) | **Get** /extended/v1/contract/by_trait | Get contracts by trait
[**GetSmartContractsStatus**](SmartContractsAPI.md#GetSmartContractsStatus) | **Get** /extended/v2/smart-contracts/status | Get smart contracts status



## GetContractById

> SmartContract GetContractById(ctx, contractId).Execute()

Get contract info



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
	contractId := "SP6P4EJF0VG8V0RB3TQQKJBHDQKEF6NVRD1KZE3C.satoshibles" // string | Contract identifier formatted as `<contract_address>.<contract_name>`

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmartContractsAPI.GetContractById(context.Background(), contractId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartContractsAPI.GetContractById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContractById`: SmartContract
	fmt.Fprintf(os.Stdout, "Response from `SmartContractsAPI.GetContractById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | Contract identifier formatted as &#x60;&lt;contract_address&gt;.&lt;contract_name&gt;&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SmartContract**](SmartContract.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContractEventsById

> GetContractEventsById200Response GetContractEventsById(ctx, contractId).Limit(limit).Offset(offset).Execute()

Get contract events



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
	contractId := "SP6P4EJF0VG8V0RB3TQQKJBHDQKEF6NVRD1KZE3C.satoshibles" // string | Contract identifier formatted as `<contract_address>.<contract_name>`
	limit := int32(56) // int32 | max number of events to fetch (optional) (default to 20)
	offset := int32(56) // int32 | Result offset (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmartContractsAPI.GetContractEventsById(context.Background(), contractId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartContractsAPI.GetContractEventsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContractEventsById`: GetContractEventsById200Response
	fmt.Fprintf(os.Stdout, "Response from `SmartContractsAPI.GetContractEventsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | Contract identifier formatted as &#x60;&lt;contract_address&gt;.&lt;contract_name&gt;&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractEventsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | max number of events to fetch | [default to 20]
 **offset** | **int32** | Result offset | [default to 0]

### Return type

[**GetContractEventsById200Response**](GetContractEventsById200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContractsByTrait

> ContractListResponse GetContractsByTrait(ctx).TraitAbi(traitAbi).Limit(limit).Offset(offset).Execute()

Get contracts by trait



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
	traitAbi := "traitAbi_example" // string | JSON abi of the trait.
	limit := int32(56) // int32 | max number of contracts fetch (optional) (default to 20)
	offset := int32(56) // int32 | index of first contract event to fetch (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmartContractsAPI.GetContractsByTrait(context.Background()).TraitAbi(traitAbi).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartContractsAPI.GetContractsByTrait``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContractsByTrait`: ContractListResponse
	fmt.Fprintf(os.Stdout, "Response from `SmartContractsAPI.GetContractsByTrait`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContractsByTraitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **traitAbi** | **string** | JSON abi of the trait. | 
 **limit** | **int32** | max number of contracts fetch | [default to 20]
 **offset** | **int32** | index of first contract event to fetch | [default to 0]

### Return type

[**ContractListResponse**](ContractListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSmartContractsStatus

> map[string]GetSmartContractsStatus200ResponseValue GetSmartContractsStatus(ctx).ContractId(contractId).Execute()

Get smart contracts status



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
	contractId := *openapiclient.NewGetSmartContractsStatusContractIdParameter() // GetSmartContractsStatusContractIdParameter | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmartContractsAPI.GetSmartContractsStatus(context.Background()).ContractId(contractId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartContractsAPI.GetSmartContractsStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSmartContractsStatus`: map[string]GetSmartContractsStatus200ResponseValue
	fmt.Fprintf(os.Stdout, "Response from `SmartContractsAPI.GetSmartContractsStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSmartContractsStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contractId** | [**GetSmartContractsStatusContractIdParameter**](GetSmartContractsStatusContractIdParameter.md) |  | 

### Return type

[**map[string]GetSmartContractsStatus200ResponseValue**](GetSmartContractsStatus200ResponseValue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

