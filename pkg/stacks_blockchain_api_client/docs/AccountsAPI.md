# \AccountsAPI

All URIs are relative to *https://api.hiro.so*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountAssets**](AccountsAPI.md#GetAccountAssets) | **Get** /extended/v1/address/{principal}/assets | Get account assets
[**GetAccountBalance**](AccountsAPI.md#GetAccountBalance) | **Get** /extended/v1/address/{principal}/balances | Get account balances
[**GetAccountInbound**](AccountsAPI.md#GetAccountInbound) | **Get** /extended/v1/address/{principal}/stx_inbound | Get inbound STX transfers
[**GetAccountNonces**](AccountsAPI.md#GetAccountNonces) | **Get** /extended/v1/address/{principal}/nonces | Get the latest nonce used by an account
[**GetAccountStxBalance**](AccountsAPI.md#GetAccountStxBalance) | **Get** /extended/v1/address/{principal}/stx | Get account STX balance
[**GetAccountTransactions**](AccountsAPI.md#GetAccountTransactions) | **Get** /extended/v1/address/{principal}/transactions | Get account transactions
[**GetAccountTransactionsWithTransfers**](AccountsAPI.md#GetAccountTransactionsWithTransfers) | **Get** /extended/v1/address/{principal}/transactions_with_transfers | Get account transactions including STX transfers for each transaction.
[**GetSingleTransactionWithTransfers**](AccountsAPI.md#GetSingleTransactionWithTransfers) | **Get** /extended/v1/address/{principal}/{tx_id}/with_transfers | Get account transaction information for specific transaction



## GetAccountAssets

> AddressAssetsListResponse GetAccountAssets(ctx, principal).Limit(limit).Offset(offset).Unanchored(unanchored).UntilBlock(untilBlock).Execute()

Get account assets



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
	principal := *openapiclient.NewGetFilteredEventsAddressParameter() // GetFilteredEventsAddressParameter | 
	limit := int32(56) // int32 | Results per page (optional) (default to 20)
	offset := int32(56) // int32 | Result offset (optional) (default to 0)
	unanchored := true // bool | Include data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)
	untilBlock := "60000" // string | Block hash or block height. Return data representing the state up until that point in time, rather than the current block. Note - Use either of the query parameters but not both at a time. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountAssets(context.Background(), principal).Limit(limit).Offset(offset).Unanchored(unanchored).UntilBlock(untilBlock).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountAssets`: AddressAssetsListResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountAssets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | [**GetFilteredEventsAddressParameter**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Results per page | [default to 20]
 **offset** | **int32** | Result offset | [default to 0]
 **unanchored** | **bool** | Include data from unanchored (i.e. unconfirmed) microblocks | [default to false]
 **untilBlock** | **string** | Block hash or block height. Return data representing the state up until that point in time, rather than the current block. Note - Use either of the query parameters but not both at a time. | 

### Return type

[**AddressAssetsListResponse**](AddressAssetsListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountBalance

> AddressBalanceResponse GetAccountBalance(ctx, principal).Unanchored(unanchored).UntilBlock(untilBlock).Execute()

Get account balances



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
	principal := *openapiclient.NewGetFilteredEventsAddressParameter() // GetFilteredEventsAddressParameter | 
	unanchored := true // bool | Include data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)
	untilBlock := "60000" // string | Block hash or block height. Return data representing the state up until that point in time, rather than the current block. Note - Use either of the query parameters but not both at a time. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountBalance(context.Background(), principal).Unanchored(unanchored).UntilBlock(untilBlock).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountBalance`: AddressBalanceResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | [**GetFilteredEventsAddressParameter**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unanchored** | **bool** | Include data from unanchored (i.e. unconfirmed) microblocks | [default to false]
 **untilBlock** | **string** | Block hash or block height. Return data representing the state up until that point in time, rather than the current block. Note - Use either of the query parameters but not both at a time. | 

### Return type

[**AddressBalanceResponse**](AddressBalanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountInbound

> AddressStxInboundListResponse GetAccountInbound(ctx, principal).Limit(limit).Offset(offset).Height(height).Unanchored(unanchored).UntilBlock(untilBlock).Execute()

Get inbound STX transfers



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
	principal := *openapiclient.NewGetFilteredEventsAddressParameter() // GetFilteredEventsAddressParameter | 
	limit := int32(56) // int32 | Results per page (optional) (default to 20)
	offset := int32(56) // int32 | Result offset (optional) (default to 0)
	height := int32(56) // int32 | Filter for transactions only at this given block height (optional)
	unanchored := true // bool | Include data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)
	untilBlock := "60000" // string | Block hash or block height. Return data representing the state up until that point in time, rather than the current block. Note - Use either of the query parameters but not both at a time. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountInbound(context.Background(), principal).Limit(limit).Offset(offset).Height(height).Unanchored(unanchored).UntilBlock(untilBlock).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountInbound``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountInbound`: AddressStxInboundListResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountInbound`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | [**GetFilteredEventsAddressParameter**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountInboundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Results per page | [default to 20]
 **offset** | **int32** | Result offset | [default to 0]
 **height** | **int32** | Filter for transactions only at this given block height | 
 **unanchored** | **bool** | Include data from unanchored (i.e. unconfirmed) microblocks | [default to false]
 **untilBlock** | **string** | Block hash or block height. Return data representing the state up until that point in time, rather than the current block. Note - Use either of the query parameters but not both at a time. | 

### Return type

[**AddressStxInboundListResponse**](AddressStxInboundListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountNonces

> AddressNonces GetAccountNonces(ctx, principal).BlockHeight(blockHeight).BlockHash(blockHash).Execute()

Get the latest nonce used by an account



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
	principal := *openapiclient.NewGetFilteredEventsAddressParameter() // GetFilteredEventsAddressParameter | 
	blockHeight := int32(66119) // int32 | Optionally get the nonce at a given block height. (optional)
	blockHash := "0x72d53f3cba39e149dcd42708e535bdae03d73e60d2fe853aaf61c0b392f521e9" // string | Optionally get the nonce at a given block hash. Note - Use either of the query parameters but not both at a time. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountNonces(context.Background(), principal).BlockHeight(blockHeight).BlockHash(blockHash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountNonces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountNonces`: AddressNonces
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountNonces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | [**GetFilteredEventsAddressParameter**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountNoncesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blockHeight** | **int32** | Optionally get the nonce at a given block height. | 
 **blockHash** | **string** | Optionally get the nonce at a given block hash. Note - Use either of the query parameters but not both at a time. | 

### Return type

[**AddressNonces**](AddressNonces.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountStxBalance

> AddressStxBalance GetAccountStxBalance(ctx, principal).Unanchored(unanchored).UntilBlock(untilBlock).Execute()

Get account STX balance



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
	principal := *openapiclient.NewGetFilteredEventsAddressParameter() // GetFilteredEventsAddressParameter | 
	unanchored := true // bool | Include data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)
	untilBlock := "60000" // string | Block hash or block height. Return data representing the state up until that point in time, rather than the current block. Note - Use either of the query parameters but not both at a time. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountStxBalance(context.Background(), principal).Unanchored(unanchored).UntilBlock(untilBlock).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountStxBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountStxBalance`: AddressStxBalance
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountStxBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | [**GetFilteredEventsAddressParameter**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountStxBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unanchored** | **bool** | Include data from unanchored (i.e. unconfirmed) microblocks | [default to false]
 **untilBlock** | **string** | Block hash or block height. Return data representing the state up until that point in time, rather than the current block. Note - Use either of the query parameters but not both at a time. | 

### Return type

[**AddressStxBalance**](AddressStxBalance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountTransactions

> AddressTransactionsListResponse GetAccountTransactions(ctx, principal).Limit(limit).Offset(offset).Height(height).Unanchored(unanchored).UntilBlock(untilBlock).Execute()

Get account transactions



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
	principal := *openapiclient.NewGetFilteredEventsAddressParameter() // GetFilteredEventsAddressParameter | 
	limit := int32(56) // int32 | Results per page (optional) (default to 20)
	offset := int32(56) // int32 | Result offset (optional) (default to 0)
	height := int32(56) // int32 | Filter for transactions only at this given block height (optional)
	unanchored := true // bool | Include data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)
	untilBlock := "60000" // string | Block hash or block height. Return data representing the state up until that point in time, rather than the current block. Note - Use either of the query parameters but not both at a time. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountTransactions(context.Background(), principal).Limit(limit).Offset(offset).Height(height).Unanchored(unanchored).UntilBlock(untilBlock).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountTransactions`: AddressTransactionsListResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | [**GetFilteredEventsAddressParameter**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Results per page | [default to 20]
 **offset** | **int32** | Result offset | [default to 0]
 **height** | **int32** | Filter for transactions only at this given block height | 
 **unanchored** | **bool** | Include data from unanchored (i.e. unconfirmed) microblocks | [default to false]
 **untilBlock** | **string** | Block hash or block height. Return data representing the state up until that point in time, rather than the current block. Note - Use either of the query parameters but not both at a time. | 

### Return type

[**AddressTransactionsListResponse**](AddressTransactionsListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountTransactionsWithTransfers

> AddressTransactionsWithTransfersListResponse GetAccountTransactionsWithTransfers(ctx, principal).Limit(limit).Offset(offset).Height(height).Unanchored(unanchored).UntilBlock(untilBlock).Execute()

Get account transactions including STX transfers for each transaction.



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
	principal := *openapiclient.NewGetFilteredEventsAddressParameter() // GetFilteredEventsAddressParameter | 
	limit := int32(56) // int32 | Results per page (optional) (default to 20)
	offset := int32(56) // int32 | Result offset (optional) (default to 0)
	height := int32(56) // int32 | Filter for transactions only at this given block height (optional)
	unanchored := true // bool | Include data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)
	untilBlock := "60000" // string | Block hash or block height. Return data representing the state up until that point in time, rather than the current block. Note - Use either of the query parameters but not both at a time. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountTransactionsWithTransfers(context.Background(), principal).Limit(limit).Offset(offset).Height(height).Unanchored(unanchored).UntilBlock(untilBlock).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountTransactionsWithTransfers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountTransactionsWithTransfers`: AddressTransactionsWithTransfersListResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountTransactionsWithTransfers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | [**GetFilteredEventsAddressParameter**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountTransactionsWithTransfersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Results per page | [default to 20]
 **offset** | **int32** | Result offset | [default to 0]
 **height** | **int32** | Filter for transactions only at this given block height | 
 **unanchored** | **bool** | Include data from unanchored (i.e. unconfirmed) microblocks | [default to false]
 **untilBlock** | **string** | Block hash or block height. Return data representing the state up until that point in time, rather than the current block. Note - Use either of the query parameters but not both at a time. | 

### Return type

[**AddressTransactionsWithTransfersListResponse**](AddressTransactionsWithTransfersListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSingleTransactionWithTransfers

> AddressTransactionWithTransfers GetSingleTransactionWithTransfers(ctx, principal, txId).Execute()

Get account transaction information for specific transaction



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
	principal := *openapiclient.NewGetFilteredEventsAddressParameter() // GetFilteredEventsAddressParameter | 
	txId := "0x34d79c7cfc2fe525438736733e501a4bf0308a5556e3e080d1e2c0858aad7448" // string | Transaction ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetSingleTransactionWithTransfers(context.Background(), principal, txId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetSingleTransactionWithTransfers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSingleTransactionWithTransfers`: AddressTransactionWithTransfers
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetSingleTransactionWithTransfers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | [**GetFilteredEventsAddressParameter**](.md) |  | 
**txId** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleTransactionWithTransfersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AddressTransactionWithTransfers**](AddressTransactionWithTransfers.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

