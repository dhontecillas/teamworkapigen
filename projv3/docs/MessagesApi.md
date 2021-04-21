# \MessagesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PATCHProjectsApiV3MessagesmessageIdJson**](MessagesApi.md#PATCHProjectsApiV3MessagesmessageIdJson) | **Patch** /projects/api/v3/messages/{messageId}.json | Edit a message.



## PATCHProjectsApiV3MessagesmessageIdJson

> PATCHProjectsApiV3MessagesmessageIdJson(ctx, messageId).MessageRequest(messageRequest).Execute()

Edit a message.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    messageId := int32(56) // int32 | 
    messageRequest := *openapiclient.NewMessageRequest() // MessageRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MessagesApi.PATCHProjectsApiV3MessagesmessageIdJson(context.Background(), messageId).MessageRequest(messageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagesApi.PATCHProjectsApiV3MessagesmessageIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHProjectsApiV3MessagesmessageIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **messageRequest** | [**MessageRequest**](MessageRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

