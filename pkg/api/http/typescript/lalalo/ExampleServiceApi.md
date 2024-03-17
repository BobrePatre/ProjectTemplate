# .ExampleServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**exampleServiceExample**](ExampleServiceApi.md#exampleServiceExample) | **POST** /grpc/api/example | Create an example


# **exampleServiceExample**
> ApilalaExampleResponse exampleServiceExample(body)


### Example


```typescript
import {  } from '';
import * as fs from 'fs';

const configuration = .createConfiguration();
const apiInstance = new .ExampleServiceApi(configuration);

let body:.ExampleServiceApiExampleServiceExampleRequest = {
  // ApilalaExampleRequest
  body: {
    title: "title_example",
    description: "description_example",
    body: "body_example",
  },
};

apiInstance.exampleServiceExample(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **ApilalaExampleRequest**|  |


### Return type

**ApilalaExampleResponse**

### Authorization

[BearerAuth](README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints) [[Back to Model list]](README.md#documentation-for-models) [[Back to README]](README.md)


