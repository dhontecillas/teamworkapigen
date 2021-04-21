# \MessageRepliesApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEMessageRepliesIdJson**](MessageRepliesApi.md#DELETEMessageRepliesIdJson) | **Delete** /messageReplies/{id}.json | Delete Message Reply
[**GETMessageRepliesIdJson**](MessageRepliesApi.md#GETMessageRepliesIdJson) | **Get** /messageReplies/{id}.json | Retrieve a Single Message Reply
[**GETMessagesIdRepliesJson**](MessageRepliesApi.md#GETMessagesIdRepliesJson) | **Get** /messages/{id}/replies.json | Retrieve Replies to a Message
[**POSTMessagesIdMessageRepliesJson**](MessageRepliesApi.md#POSTMessagesIdMessageRepliesJson) | **Post** /messages/{id}/messageReplies.json | Create a Message Reply
[**PUTMessageRepliesIdJson**](MessageRepliesApi.md#PUTMessageRepliesIdJson) | **Put** /messageReplies/{id}.json | Update Message Reply
[**PUTMessageRepliesIdMarkreadJson**](MessageRepliesApi.md#PUTMessageRepliesIdMarkreadJson) | **Put** /messageReplies/{id}/markread.json | Mark Message Reply as Read



## DELETEMessageRepliesIdJson

> InlineResponse2001 DELETEMessageRepliesIdJson(ctx, id).Execute()

Delete Message Reply



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MessageRepliesApi.DELETEMessageRepliesIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageRepliesApi.DELETEMessageRepliesIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DELETEMessageRepliesIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `MessageRepliesApi.DELETEMessageRepliesIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEMessageRepliesIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETMessageRepliesIdJson

> InlineResponse20032 GETMessageRepliesIdJson(ctx, id).Execute()

Retrieve a Single Message Reply



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MessageRepliesApi.GETMessageRepliesIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageRepliesApi.GETMessageRepliesIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETMessageRepliesIdJson`: InlineResponse20032
    fmt.Fprintf(os.Stdout, "Response from `MessageRepliesApi.GETMessageRepliesIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETMessageRepliesIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20032**](InlineResponse20032.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETMessagesIdRepliesJson

> InlineResponse20032 GETMessagesIdRepliesJson(ctx, id).Page(page).PageSize(pageSize).Execute()

Retrieve Replies to a Message



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
    id := "id_example" // string | 
    page := int32(56) // int32 | The page you want (optional)
    pageSize := int32(56) // int32 | The size of each page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MessageRepliesApi.GETMessagesIdRepliesJson(context.Background(), id).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageRepliesApi.GETMessagesIdRepliesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETMessagesIdRepliesJson`: InlineResponse20032
    fmt.Fprintf(os.Stdout, "Response from `MessageRepliesApi.GETMessagesIdRepliesJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETMessagesIdRepliesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page you want | 
 **pageSize** | **int32** | The size of each page | 

### Return type

[**InlineResponse20032**](InlineResponse20032.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTMessagesIdMessageRepliesJson

> InlineResponse201 POSTMessagesIdMessageRepliesJson(ctx, id).Body(body).Execute()

Create a Message Reply



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
    id := "id_example" // string | 
    body := *openapiclient.NewInlineObject30() // InlineObject30 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MessageRepliesApi.POSTMessagesIdMessageRepliesJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageRepliesApi.POSTMessagesIdMessageRepliesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTMessagesIdMessageRepliesJson`: InlineResponse201
    fmt.Fprintf(os.Stdout, "Response from `MessageRepliesApi.POSTMessagesIdMessageRepliesJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTMessagesIdMessageRepliesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject30**](InlineObject30.md) |  | 

### Return type

[**InlineResponse201**](InlineResponse201.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTMessageRepliesIdJson

> InlineResponse2001 PUTMessageRepliesIdJson(ctx, id).Body(body).Execute()

Update Message Reply



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
    id := "id_example" // string | 
    body := *openapiclient.NewInlineObject27() // InlineObject27 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MessageRepliesApi.PUTMessageRepliesIdJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageRepliesApi.PUTMessageRepliesIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTMessageRepliesIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `MessageRepliesApi.PUTMessageRepliesIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTMessageRepliesIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject27**](InlineObject27.md) |  | 

### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTMessageRepliesIdMarkreadJson

> InlineResponse2001 PUTMessageRepliesIdMarkreadJson(ctx, id).Execute()

Mark Message Reply as Read



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MessageRepliesApi.PUTMessageRepliesIdMarkreadJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageRepliesApi.PUTMessageRepliesIdMarkreadJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTMessageRepliesIdMarkreadJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `MessageRepliesApi.PUTMessageRepliesIdMarkreadJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTMessageRepliesIdMarkreadJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

