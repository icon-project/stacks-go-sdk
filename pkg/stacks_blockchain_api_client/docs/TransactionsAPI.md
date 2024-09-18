# \TransactionsAPI

All URIs are relative to *https://api.hiro.so*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAddressMempoolTransactions**](TransactionsAPI.md#GetAddressMempoolTransactions) | **Get** /extended/v1/address/{principal}/mempool | Transactions for address
[**GetAddressTransactionEvents**](TransactionsAPI.md#GetAddressTransactionEvents) | **Get** /extended/v2/addresses/{address}/transactions/{tx_id}/events | Get events for an address transaction
[**GetAddressTransactions**](TransactionsAPI.md#GetAddressTransactions) | **Get** /extended/v2/addresses/{address}/transactions | Get address transactions
[**GetDroppedMempoolTransactionList**](TransactionsAPI.md#GetDroppedMempoolTransactionList) | **Get** /extended/v1/tx/mempool/dropped | Get dropped mempool transactions
[**GetFilteredEvents**](TransactionsAPI.md#GetFilteredEvents) | **Get** /extended/v1/tx/events | Transaction Events
[**GetMempoolTransactionList**](TransactionsAPI.md#GetMempoolTransactionList) | **Get** /extended/v1/tx/mempool | Get mempool transactions
[**GetMempoolTransactionStats**](TransactionsAPI.md#GetMempoolTransactionStats) | **Get** /extended/v1/tx/mempool/stats | Get statistics for mempool transactions
[**GetRawTransactionById**](TransactionsAPI.md#GetRawTransactionById) | **Get** /extended/v1/tx/{tx_id}/raw | Get raw transaction
[**GetTransactionById**](TransactionsAPI.md#GetTransactionById) | **Get** /extended/v1/tx/{tx_id} | Get transaction
[**GetTransactionList**](TransactionsAPI.md#GetTransactionList) | **Get** /extended/v1/tx/ | Get recent transactions
[**GetTransactionsByBlock**](TransactionsAPI.md#GetTransactionsByBlock) | **Get** /extended/v2/blocks/{height_or_hash}/transactions | Get transactions by block
[**GetTransactionsByBlockHash**](TransactionsAPI.md#GetTransactionsByBlockHash) | **Get** /extended/v1/tx/block/{block_hash} | Transactions by block hash
[**GetTransactionsByBlockHeight**](TransactionsAPI.md#GetTransactionsByBlockHeight) | **Get** /extended/v1/tx/block_height/{height} | Transactions by block height
[**GetTxListDetails**](TransactionsAPI.md#GetTxListDetails) | **Get** /extended/v1/tx/multiple | Get list of details for transactions



## GetAddressMempoolTransactions

> GetMempoolTransactionList200Response GetAddressMempoolTransactions(ctx, principal).Limit(limit).Offset(offset).Unanchored(unanchored).Execute()

Transactions for address



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetAddressMempoolTransactions(context.Background(), principal).Limit(limit).Offset(offset).Unanchored(unanchored).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetAddressMempoolTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddressMempoolTransactions`: GetMempoolTransactionList200Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetAddressMempoolTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | [**GetFilteredEventsAddressParameter**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressMempoolTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Results per page | [default to 20]
 **offset** | **int32** | Result offset | [default to 0]
 **unanchored** | **bool** | Include data from unanchored (i.e. unconfirmed) microblocks | [default to false]

### Return type

[**GetMempoolTransactionList200Response**](GetMempoolTransactionList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddressTransactionEvents

> GetAddressTransactionEvents200Response GetAddressTransactionEvents(ctx, address, txId).Limit(limit).Offset(offset).Execute()

Get events for an address transaction



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
	address := *openapiclient.NewGetFilteredEventsAddressParameter() // GetFilteredEventsAddressParameter | 
	txId := "0xf6bd5f4a7b26184a3466340b2e99fd003b4962c0e382a7e4b6a13df3dd7a91c6" // string | Transaction ID
	limit := int32(56) // int32 | Results per page (optional) (default to 20)
	offset := int32(56) // int32 | Result offset (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetAddressTransactionEvents(context.Background(), address, txId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetAddressTransactionEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddressTransactionEvents`: GetAddressTransactionEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetAddressTransactionEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | [**GetFilteredEventsAddressParameter**](.md) |  | 
**txId** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressTransactionEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Results per page | [default to 20]
 **offset** | **int32** | Result offset | [default to 0]

### Return type

[**GetAddressTransactionEvents200Response**](GetAddressTransactionEvents200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddressTransactions

> GetAddressTransactions200Response GetAddressTransactions(ctx, address).Limit(limit).Offset(offset).Execute()

Get address transactions



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
	address := *openapiclient.NewGetFilteredEventsAddressParameter() // GetFilteredEventsAddressParameter | 
	limit := int32(56) // int32 | Results per page (optional) (default to 20)
	offset := int32(56) // int32 | Result offset (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetAddressTransactions(context.Background(), address).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetAddressTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddressTransactions`: GetAddressTransactions200Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetAddressTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | [**GetFilteredEventsAddressParameter**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Results per page | [default to 20]
 **offset** | **int32** | Result offset | [default to 0]

### Return type

[**GetAddressTransactions200Response**](GetAddressTransactions200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDroppedMempoolTransactionList

> GetDroppedMempoolTransactionList200Response GetDroppedMempoolTransactionList(ctx).Offset(offset).Limit(limit).Execute()

Get dropped mempool transactions



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
	offset := int32(56) // int32 | Result offset (optional) (default to 0)
	limit := int32(56) // int32 | Results per page (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetDroppedMempoolTransactionList(context.Background()).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetDroppedMempoolTransactionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDroppedMempoolTransactionList`: GetDroppedMempoolTransactionList200Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetDroppedMempoolTransactionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDroppedMempoolTransactionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Result offset | [default to 0]
 **limit** | **int32** | Results per page | [default to 20]

### Return type

[**GetDroppedMempoolTransactionList200Response**](GetDroppedMempoolTransactionList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFilteredEvents

> ListOfEvents GetFilteredEvents(ctx).TxId(txId).Address(address).Type_(type_).Offset(offset).Limit(limit).Execute()

Transaction Events



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
	txId := "0xf6bd5f4a7b26184a3466340b2e99fd003b4962c0e382a7e4b6a13df3dd7a91c6" // string | Transaction ID (optional)
	address := *openapiclient.NewGetFilteredEventsAddressParameter() // GetFilteredEventsAddressParameter |  (optional)
	type_ := []openapiclient.GetFilteredEventsTypeParameterInner{*openapiclient.NewGetFilteredEventsTypeParameterInner()} // []GetFilteredEventsTypeParameterInner |  (optional)
	offset := int32(56) // int32 | Result offset (optional) (default to 0)
	limit := int32(56) // int32 | Results per page (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetFilteredEvents(context.Background()).TxId(txId).Address(address).Type_(type_).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetFilteredEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFilteredEvents`: ListOfEvents
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetFilteredEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFilteredEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **txId** | **string** | Transaction ID | 
 **address** | [**GetFilteredEventsAddressParameter**](GetFilteredEventsAddressParameter.md) |  | 
 **type_** | [**[]GetFilteredEventsTypeParameterInner**](GetFilteredEventsTypeParameterInner.md) |  | 
 **offset** | **int32** | Result offset | [default to 0]
 **limit** | **int32** | Results per page | [default to 20]

### Return type

[**ListOfEvents**](ListOfEvents.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMempoolTransactionList

> GetMempoolTransactionList200Response GetMempoolTransactionList(ctx).SenderAddress(senderAddress).RecipientAddress(recipientAddress).Address(address).OrderBy(orderBy).Order(order).Unanchored(unanchored).Offset(offset).Limit(limit).Execute()

Get mempool transactions



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
	senderAddress := "SP318Q55DEKHRXJK696033DQN5C54D9K2EE6DHRWP" // string | STX Address (optional)
	recipientAddress := "SP318Q55DEKHRXJK696033DQN5C54D9K2EE6DHRWP" // string | STX Address (optional)
	address := "SP318Q55DEKHRXJK696033DQN5C54D9K2EE6DHRWP" // string | STX Address (optional)
	orderBy := *openapiclient.NewOrderBy() // OrderBy | Option to sort results by transaction age, size, or fee rate. (optional)
	order := *openapiclient.NewOrder() // Order | Results order (optional)
	unanchored := true // bool | Include data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)
	offset := int32(56) // int32 | Result offset (optional) (default to 0)
	limit := int32(56) // int32 | Results per page (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetMempoolTransactionList(context.Background()).SenderAddress(senderAddress).RecipientAddress(recipientAddress).Address(address).OrderBy(orderBy).Order(order).Unanchored(unanchored).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetMempoolTransactionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMempoolTransactionList`: GetMempoolTransactionList200Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetMempoolTransactionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMempoolTransactionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **senderAddress** | **string** | STX Address | 
 **recipientAddress** | **string** | STX Address | 
 **address** | **string** | STX Address | 
 **orderBy** | [**OrderBy**](OrderBy.md) | Option to sort results by transaction age, size, or fee rate. | 
 **order** | [**Order**](Order.md) | Results order | 
 **unanchored** | **bool** | Include data from unanchored (i.e. unconfirmed) microblocks | [default to false]
 **offset** | **int32** | Result offset | [default to 0]
 **limit** | **int32** | Results per page | [default to 20]

### Return type

[**GetMempoolTransactionList200Response**](GetMempoolTransactionList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMempoolTransactionStats

> MempoolTransactionStatsResponse GetMempoolTransactionStats(ctx).Execute()

Get statistics for mempool transactions



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
	resp, r, err := apiClient.TransactionsAPI.GetMempoolTransactionStats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetMempoolTransactionStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMempoolTransactionStats`: MempoolTransactionStatsResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetMempoolTransactionStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMempoolTransactionStatsRequest struct via the builder pattern


### Return type

[**MempoolTransactionStatsResponse**](MempoolTransactionStatsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRawTransactionById

> GetRawTransactionResult GetRawTransactionById(ctx, txId).EventLimit(eventLimit).EventOffset(eventOffset).Execute()

Get raw transaction



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
	txId := "0xf6bd5f4a7b26184a3466340b2e99fd003b4962c0e382a7e4b6a13df3dd7a91c6" // string | Transaction ID
	eventLimit := int32(56) // int32 | Results per page (optional) (default to 20)
	eventOffset := int32(56) // int32 | Result offset (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetRawTransactionById(context.Background(), txId).EventLimit(eventLimit).EventOffset(eventOffset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetRawTransactionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRawTransactionById`: GetRawTransactionResult
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetRawTransactionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**txId** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRawTransactionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventLimit** | **int32** | Results per page | [default to 20]
 **eventOffset** | **int32** | Result offset | [default to 0]

### Return type

[**GetRawTransactionResult**](GetRawTransactionResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionById

> GetTransactionById200Response GetTransactionById(ctx, txId).EventLimit(eventLimit).EventOffset(eventOffset).Unanchored(unanchored).Execute()

Get transaction



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
	txId := "0xf6bd5f4a7b26184a3466340b2e99fd003b4962c0e382a7e4b6a13df3dd7a91c6" // string | Transaction ID
	eventLimit := int32(56) // int32 | Results per page (optional) (default to 20)
	eventOffset := int32(56) // int32 | Result offset (optional) (default to 0)
	unanchored := true // bool | Include data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetTransactionById(context.Background(), txId).EventLimit(eventLimit).EventOffset(eventOffset).Unanchored(unanchored).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransactionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionById`: GetTransactionById200Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransactionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**txId** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventLimit** | **int32** | Results per page | [default to 20]
 **eventOffset** | **int32** | Result offset | [default to 0]
 **unanchored** | **bool** | Include data from unanchored (i.e. unconfirmed) microblocks | [default to false]

### Return type

[**GetTransactionById200Response**](GetTransactionById200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionList

> GetTransactionList200Response GetTransactionList(ctx).Offset(offset).Limit(limit).Type_(type_).Unanchored(unanchored).Order(order).SortBy(sortBy).FromAddress(fromAddress).ToAddress(toAddress).StartTime(startTime).EndTime(endTime).ContractId(contractId).FunctionName(functionName).Nonce(nonce).Execute()

Get recent transactions



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
	offset := int32(56) // int32 | Result offset (optional) (default to 0)
	limit := int32(56) // int32 | Results per page (optional) (default to 20)
	type_ := []openapiclient.GetTransactionListTypeParameterInner{*openapiclient.NewGetTransactionListTypeParameterInner()} // []GetTransactionListTypeParameterInner |  (optional)
	unanchored := true // bool | Include data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)
	order := *openapiclient.NewGetTransactionListOrderParameter() // GetTransactionListOrderParameter |  (optional)
	sortBy := *openapiclient.NewGetTransactionListSortByParameter() // GetTransactionListSortByParameter | Option to sort results by block height, timestamp, or fee (optional) (default to block_height)
	fromAddress := "fromAddress_example" // string | Option to filter results by sender address (optional)
	toAddress := "toAddress_example" // string | Option to filter results by recipient address (optional)
	startTime := int32(1704067200) // int32 | Filter by transactions after this timestamp (unix timestamp in seconds) (optional)
	endTime := int32(1706745599) // int32 | Filter by transactions before this timestamp (unix timestamp in seconds) (optional)
	contractId := "SP000000000000000000002Q6VF78.pox-4" // string | Option to filter results by contract ID (optional)
	functionName := "delegate-stx" // string | Filter by contract call transactions involving this function name (optional)
	nonce := int32(123) // int32 | Filter by transactions with this nonce (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetTransactionList(context.Background()).Offset(offset).Limit(limit).Type_(type_).Unanchored(unanchored).Order(order).SortBy(sortBy).FromAddress(fromAddress).ToAddress(toAddress).StartTime(startTime).EndTime(endTime).ContractId(contractId).FunctionName(functionName).Nonce(nonce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransactionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionList`: GetTransactionList200Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransactionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Result offset | [default to 0]
 **limit** | **int32** | Results per page | [default to 20]
 **type_** | [**[]GetTransactionListTypeParameterInner**](GetTransactionListTypeParameterInner.md) |  | 
 **unanchored** | **bool** | Include data from unanchored (i.e. unconfirmed) microblocks | [default to false]
 **order** | [**GetTransactionListOrderParameter**](GetTransactionListOrderParameter.md) |  | 
 **sortBy** | [**GetTransactionListSortByParameter**](GetTransactionListSortByParameter.md) | Option to sort results by block height, timestamp, or fee | [default to block_height]
 **fromAddress** | **string** | Option to filter results by sender address | 
 **toAddress** | **string** | Option to filter results by recipient address | 
 **startTime** | **int32** | Filter by transactions after this timestamp (unix timestamp in seconds) | 
 **endTime** | **int32** | Filter by transactions before this timestamp (unix timestamp in seconds) | 
 **contractId** | **string** | Option to filter results by contract ID | 
 **functionName** | **string** | Filter by contract call transactions involving this function name | 
 **nonce** | **int32** | Filter by transactions with this nonce | 

### Return type

[**GetTransactionList200Response**](GetTransactionList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionsByBlock

> GetTransactionsByBlock200Response GetTransactionsByBlock(ctx, heightOrHash).Limit(limit).Offset(offset).Execute()

Get transactions by block



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
	limit := int32(56) // int32 | Results per page (optional) (default to 20)
	offset := int32(56) // int32 | Result offset (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetTransactionsByBlock(context.Background(), heightOrHash).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransactionsByBlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionsByBlock`: GetTransactionsByBlock200Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransactionsByBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**heightOrHash** | [**GetBlockHeightOrHashParameter**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsByBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Results per page | [default to 20]
 **offset** | **int32** | Result offset | [default to 0]

### Return type

[**GetTransactionsByBlock200Response**](GetTransactionsByBlock200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionsByBlockHash

> GetTransactionList200Response GetTransactionsByBlockHash(ctx, blockHash).Offset(offset).Limit(limit).Execute()

Transactions by block hash



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
	blockHash := "blockHash_example" // string | 
	offset := int32(56) // int32 | Result offset (optional) (default to 0)
	limit := int32(56) // int32 | Results per page (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetTransactionsByBlockHash(context.Background(), blockHash).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransactionsByBlockHash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionsByBlockHash`: GetTransactionList200Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransactionsByBlockHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockHash** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsByBlockHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | Result offset | [default to 0]
 **limit** | **int32** | Results per page | [default to 20]

### Return type

[**GetTransactionList200Response**](GetTransactionList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionsByBlockHeight

> GetTransactionList200Response GetTransactionsByBlockHeight(ctx, height).Offset(offset).Limit(limit).Execute()

Transactions by block height



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
	height := int32(777678) // int32 | Block height
	offset := int32(56) // int32 | Result offset (optional) (default to 0)
	limit := int32(56) // int32 | Results per page (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetTransactionsByBlockHeight(context.Background(), height).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransactionsByBlockHeight``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionsByBlockHeight`: GetTransactionList200Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransactionsByBlockHeight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**height** | **int32** | Block height | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsByBlockHeightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | Result offset | [default to 0]
 **limit** | **int32** | Results per page | [default to 20]

### Return type

[**GetTransactionList200Response**](GetTransactionList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTxListDetails

> map[string]GetTxListDetails200ResponseValue GetTxListDetails(ctx).TxId(txId).EventLimit(eventLimit).EventOffset(eventOffset).Unanchored(unanchored).Execute()

Get list of details for transactions



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
	txId := []string{"Inner_example"} // []string | 
	eventLimit := int32(56) // int32 | Results per page (optional) (default to 20)
	eventOffset := int32(56) // int32 | Result offset (optional) (default to 0)
	unanchored := true // bool | Include data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetTxListDetails(context.Background()).TxId(txId).EventLimit(eventLimit).EventOffset(eventOffset).Unanchored(unanchored).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTxListDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTxListDetails`: map[string]GetTxListDetails200ResponseValue
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTxListDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTxListDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **txId** | **[]string** |  | 
 **eventLimit** | **int32** | Results per page | [default to 20]
 **eventOffset** | **int32** | Result offset | [default to 0]
 **unanchored** | **bool** | Include data from unanchored (i.e. unconfirmed) microblocks | [default to false]

### Return type

[**map[string]GetTxListDetails200ResponseValue**](GetTxListDetails200ResponseValue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

