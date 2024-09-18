# \NamesAPI

All URIs are relative to *https://api.hiro.so*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchSubdomainsListForName**](NamesAPI.md#FetchSubdomainsListForName) | **Get** /v1/names/{name}/subdomains | Get Name Subdomains
[**FetchZoneFile**](NamesAPI.md#FetchZoneFile) | **Get** /v1/names/{name}/zonefile | Get Zone File
[**GetAllNames**](NamesAPI.md#GetAllNames) | **Get** /v1/names/ | Get All Names
[**GetAllNamespaces**](NamesAPI.md#GetAllNamespaces) | **Get** /v1/namespaces/ | Get All Namespaces
[**GetHistoricalZoneFile**](NamesAPI.md#GetHistoricalZoneFile) | **Get** /v1/names/{name}/zonefile/{zoneFileHash} | Get Historical Zone File
[**GetNameInfo**](NamesAPI.md#GetNameInfo) | **Get** /v1/names/{name} | Get Name Details
[**GetNamePrice**](NamesAPI.md#GetNamePrice) | **Get** /v2/prices/names/{name} | Get Name Price
[**GetNamesOwnedByAddress**](NamesAPI.md#GetNamesOwnedByAddress) | **Get** /v1/addresses/{blockchain}/{address} | Get Names Owned by Address
[**GetNamespaceNames**](NamesAPI.md#GetNamespaceNames) | **Get** /v1/namespaces/{tld}/names | Get Namespace Names
[**GetNamespacePrice**](NamesAPI.md#GetNamespacePrice) | **Get** /v2/prices/namespaces/{tld} | Get Namespace Price



## FetchSubdomainsListForName

> []string FetchSubdomainsListForName(ctx, name).Unanchored(unanchored).Execute()

Get Name Subdomains



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
	name := "id.blockstack" // string | fully-qualified name
	unanchored := true // bool | Include data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamesAPI.FetchSubdomainsListForName(context.Background(), name).Unanchored(unanchored).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamesAPI.FetchSubdomainsListForName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchSubdomainsListForName`: []string
	fmt.Fprintf(os.Stdout, "Response from `NamesAPI.FetchSubdomainsListForName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | fully-qualified name | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchSubdomainsListForNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unanchored** | **bool** | Include data from unanchored (i.e. unconfirmed) microblocks | [default to false]

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchZoneFile

> BnsFetchFileZoneResponse FetchZoneFile(ctx, name).Unanchored(unanchored).Execute()

Get Zone File



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
	name := "bar.test" // string | fully-qualified name
	unanchored := true // bool | Include data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamesAPI.FetchZoneFile(context.Background(), name).Unanchored(unanchored).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamesAPI.FetchZoneFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchZoneFile`: BnsFetchFileZoneResponse
	fmt.Fprintf(os.Stdout, "Response from `NamesAPI.FetchZoneFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | fully-qualified name | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchZoneFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unanchored** | **bool** | Include data from unanchored (i.e. unconfirmed) microblocks | [default to false]

### Return type

[**BnsFetchFileZoneResponse**](BnsFetchFileZoneResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllNames

> []string GetAllNames(ctx).Unanchored(unanchored).Page(page).Execute()

Get All Names



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
	unanchored := true // bool | Include data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)
	page := int32(56) // int32 | names are defaulted to page 1 with 100 results. You can query specific page results by using the 'page' query parameter. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamesAPI.GetAllNames(context.Background()).Unanchored(unanchored).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamesAPI.GetAllNames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllNames`: []string
	fmt.Fprintf(os.Stdout, "Response from `NamesAPI.GetAllNames`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unanchored** | **bool** | Include data from unanchored (i.e. unconfirmed) microblocks | [default to false]
 **page** | **int32** | names are defaulted to page 1 with 100 results. You can query specific page results by using the &#39;page&#39; query parameter. | [default to 0]

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllNamespaces

> GetAllNamespaces200Response GetAllNamespaces(ctx).Unanchored(unanchored).Execute()

Get All Namespaces



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
	unanchored := true // bool | Include data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamesAPI.GetAllNamespaces(context.Background()).Unanchored(unanchored).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamesAPI.GetAllNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllNamespaces`: GetAllNamespaces200Response
	fmt.Fprintf(os.Stdout, "Response from `NamesAPI.GetAllNamespaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unanchored** | **bool** | Include data from unanchored (i.e. unconfirmed) microblocks | [default to false]

### Return type

[**GetAllNamespaces200Response**](GetAllNamespaces200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoricalZoneFile

> BnsFetchHistoricalZoneFileResponse GetHistoricalZoneFile(ctx, name, zoneFileHash).Unanchored(unanchored).Execute()

Get Historical Zone File



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
	name := "muneeb.id" // string | fully-qualified name
	zoneFileHash := "b100a68235244b012854a95f9114695679002af9" // string | zone file hash
	unanchored := true // bool | Include data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamesAPI.GetHistoricalZoneFile(context.Background(), name, zoneFileHash).Unanchored(unanchored).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamesAPI.GetHistoricalZoneFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricalZoneFile`: BnsFetchHistoricalZoneFileResponse
	fmt.Fprintf(os.Stdout, "Response from `NamesAPI.GetHistoricalZoneFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | fully-qualified name | 
**zoneFileHash** | **string** | zone file hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricalZoneFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **unanchored** | **bool** | Include data from unanchored (i.e. unconfirmed) microblocks | [default to false]

### Return type

[**BnsFetchHistoricalZoneFileResponse**](BnsFetchHistoricalZoneFileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNameInfo

> BnsGetNameInfoResponse GetNameInfo(ctx, name).Unanchored(unanchored).Execute()

Get Name Details



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
	name := "muneeb.id" // string | fully-qualified name
	unanchored := true // bool | Include data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamesAPI.GetNameInfo(context.Background(), name).Unanchored(unanchored).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamesAPI.GetNameInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNameInfo`: BnsGetNameInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `NamesAPI.GetNameInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | fully-qualified name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNameInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unanchored** | **bool** | Include data from unanchored (i.e. unconfirmed) microblocks | [default to false]

### Return type

[**BnsGetNameInfoResponse**](BnsGetNameInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamePrice

> BnsGetNamePriceResponse GetNamePrice(ctx, name).Execute()

Get Name Price



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
	name := "muneeb.id" // string | the name to query price information for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamesAPI.GetNamePrice(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamesAPI.GetNamePrice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamePrice`: BnsGetNamePriceResponse
	fmt.Fprintf(os.Stdout, "Response from `NamesAPI.GetNamePrice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | the name to query price information for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamePriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BnsGetNamePriceResponse**](BnsGetNamePriceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamesOwnedByAddress

> BnsNamesOwnByAddressResponse GetNamesOwnedByAddress(ctx, blockchain, address).Unanchored(unanchored).Execute()

Get Names Owned by Address



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
	blockchain := "stacks" // string | the layer-1 blockchain for the address
	address := "SP2J6ZY48GV1EZ5V2V5RB9MP66SW86PYKKNRV9EJ7" // string | the address to lookup
	unanchored := true // bool | Include data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamesAPI.GetNamesOwnedByAddress(context.Background(), blockchain, address).Unanchored(unanchored).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamesAPI.GetNamesOwnedByAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamesOwnedByAddress`: BnsNamesOwnByAddressResponse
	fmt.Fprintf(os.Stdout, "Response from `NamesAPI.GetNamesOwnedByAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockchain** | **string** | the layer-1 blockchain for the address | 
**address** | **string** | the address to lookup | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamesOwnedByAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **unanchored** | **bool** | Include data from unanchored (i.e. unconfirmed) microblocks | [default to false]

### Return type

[**BnsNamesOwnByAddressResponse**](BnsNamesOwnByAddressResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespaceNames

> []string GetNamespaceNames(ctx, tld).Page(page).Unanchored(unanchored).Execute()

Get Namespace Names



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
	tld := "id" // string | the namespace to fetch names from.
	page := float32(22) // float32 | namespace values are defaulted to page 1 with 100 results. You can query specific page results by using the 'page' query parameter. (optional)
	unanchored := true // bool | Include data from unanchored (i.e. unconfirmed) microblocks (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamesAPI.GetNamespaceNames(context.Background(), tld).Page(page).Unanchored(unanchored).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamesAPI.GetNamespaceNames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamespaceNames`: []string
	fmt.Fprintf(os.Stdout, "Response from `NamesAPI.GetNamespaceNames`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tld** | **string** | the namespace to fetch names from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespaceNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | namespace values are defaulted to page 1 with 100 results. You can query specific page results by using the &#39;page&#39; query parameter. | 
 **unanchored** | **bool** | Include data from unanchored (i.e. unconfirmed) microblocks | [default to false]

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespacePrice

> BnsGetNamespacePriceResponse GetNamespacePrice(ctx, tld).Execute()

Get Namespace Price



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
	tld := "id" // string | the namespace to fetch price for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamesAPI.GetNamespacePrice(context.Background(), tld).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamesAPI.GetNamespacePrice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamespacePrice`: BnsGetNamespacePriceResponse
	fmt.Fprintf(os.Stdout, "Response from `NamesAPI.GetNamespacePrice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tld** | **string** | the namespace to fetch price for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespacePriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BnsGetNamespacePriceResponse**](BnsGetNamespacePriceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

