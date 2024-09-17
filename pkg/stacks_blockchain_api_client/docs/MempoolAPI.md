# \MempoolAPI

All URIs are relative to *https://api.hiro.so*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMempoolFeePriorities**](MempoolAPI.md#GetMempoolFeePriorities) | **Get** /extended/v2/mempool/fees | Get mempool transaction fee priorities



## GetMempoolFeePriorities

> GetMempoolFeePriorities200Response GetMempoolFeePriorities(ctx).Execute()

Get mempool transaction fee priorities



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
	resp, r, err := apiClient.MempoolAPI.GetMempoolFeePriorities(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MempoolAPI.GetMempoolFeePriorities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMempoolFeePriorities`: GetMempoolFeePriorities200Response
	fmt.Fprintf(os.Stdout, "Response from `MempoolAPI.GetMempoolFeePriorities`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMempoolFeePrioritiesRequest struct via the builder pattern


### Return type

[**GetMempoolFeePriorities200Response**](GetMempoolFeePriorities200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

