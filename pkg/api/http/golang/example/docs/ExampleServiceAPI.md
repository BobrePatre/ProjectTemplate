# \ExampleServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExampleServiceExample**](ExampleServiceAPI.md#ExampleServiceExample) | **Post** /grpc/api/example | Create an example



## ExampleServiceExample

> ApiexampleExampleResponse ExampleServiceExample(ctx).Body(body).Execute()

Create an example

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
	body := *openapiclient.NewApiexampleExampleRequest() // ApiexampleExampleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExampleServiceAPI.ExampleServiceExample(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExampleServiceAPI.ExampleServiceExample``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExampleServiceExample`: ApiexampleExampleResponse
	fmt.Fprintf(os.Stdout, "Response from `ExampleServiceAPI.ExampleServiceExample`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExampleServiceExampleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiexampleExampleRequest**](ApiexampleExampleRequest.md) |  | 

### Return type

[**ApiexampleExampleResponse**](ApiexampleExampleResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

