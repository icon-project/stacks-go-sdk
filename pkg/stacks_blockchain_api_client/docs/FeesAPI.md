# \FeesAPI

All URIs are relative to *https://api.hiro.so*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchFeeRate**](FeesAPI.md#FetchFeeRate) | **Post** /extended/v1/fee_rate/ | Fetch fee rate



## FetchFeeRate

> FeeRate FetchFeeRate(ctx).FeeRateRequest(feeRateRequest).Execute()

Fetch fee rate



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
	feeRateRequest := *openapiclient.NewFeeRateRequest("Transaction_example") // FeeRateRequest | Request to fetch fee for a transaction

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeesAPI.FetchFeeRate(context.Background()).FeeRateRequest(feeRateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeesAPI.FetchFeeRate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchFeeRate`: FeeRate
	fmt.Fprintf(os.Stdout, "Response from `FeesAPI.FetchFeeRate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchFeeRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **feeRateRequest** | [**FeeRateRequest**](FeeRateRequest.md) | Request to fetch fee for a transaction | 

### Return type

[**FeeRate**](FeeRate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

