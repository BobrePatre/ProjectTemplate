# openapi.api.ExampleServiceApi

## Load the API package
```dart
import 'package:openapi/api.dart';
```

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**exampleServiceExample**](ExampleServiceApi.md#exampleserviceexample) | **POST** /grpc/api/example | Create an example


# **exampleServiceExample**
> ApiexampleExampleResponse exampleServiceExample(body)

Create an example

### Example
```dart
import 'package:openapi/api.dart';
// TODO Configure OAuth2 access token for authorization: BearerAuth
//defaultApiClient.getAuthentication<OAuth>('BearerAuth').accessToken = 'YOUR_ACCESS_TOKEN';

final api = Openapi().getExampleServiceApi();
final ApiexampleExampleRequest body = ; // ApiexampleExampleRequest | 

try {
    final response = api.exampleServiceExample(body);
    print(response);
} catch on DioException (e) {
    print('Exception when calling ExampleServiceApi->exampleServiceExample: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiexampleExampleRequest**](ApiexampleExampleRequest.md)|  | 

### Return type

[**ApiexampleExampleResponse**](ApiexampleExampleResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

