# \CommentsApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETECommentsIdJson**](CommentsApi.md#DELETECommentsIdJson) | **Delete** /comments/{id}.json | Deleting a Comment
[**GETCommentsCommentIdJson**](CommentsApi.md#GETCommentsCommentIdJson) | **Get** /comments/{comment_id}.json | Retrieving a Specific Comment
[**GETCommentsJson**](CommentsApi.md#GETCommentsJson) | **Get** /comments.json | Retrieving Comments across all types
[**GETResourceResourceIdCommentsJson**](CommentsApi.md#GETResourceResourceIdCommentsJson) | **Get** /resource/{resource_id}/comments.json | Retrieving Recent Comments
[**POSTResourceResourceIdCommentsJson**](CommentsApi.md#POSTResourceResourceIdCommentsJson) | **Post** /resource/{resource_id}/comments.json | Creating a Comment
[**PUTCommentsIdJson**](CommentsApi.md#PUTCommentsIdJson) | **Put** /comments/{id}.json | Updating a Comment
[**PUTCommentsIdMarkreadJson**](CommentsApi.md#PUTCommentsIdMarkreadJson) | **Put** /comments/{id}/markread.json | Mark a Comment as Read



## DELETECommentsIdJson

> InlineResponse2001 DELETECommentsIdJson(ctx, id).Execute()

Deleting a Comment



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
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommentsApi.DELETECommentsIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommentsApi.DELETECommentsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DELETECommentsIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `CommentsApi.DELETECommentsIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETECommentsIdJsonRequest struct via the builder pattern


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


## GETCommentsCommentIdJson

> InlineResponse20010 GETCommentsCommentIdJson(ctx, commentId).Execute()

Retrieving a Specific Comment



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
    commentId := "commentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommentsApi.GETCommentsCommentIdJson(context.Background(), commentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommentsApi.GETCommentsCommentIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCommentsCommentIdJson`: InlineResponse20010
    fmt.Fprintf(os.Stdout, "Response from `CommentsApi.GETCommentsCommentIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCommentsCommentIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20010**](InlineResponse20010.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETCommentsJson

> InlineResponse2007 GETCommentsJson(ctx).Page(page).PageSize(pageSize).Execute()

Retrieving Comments across all types



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
    page := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommentsApi.GETCommentsJson(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommentsApi.GETCommentsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCommentsJson`: InlineResponse2007
    fmt.Fprintf(os.Stdout, "Response from `CommentsApi.GETCommentsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETCommentsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 

### Return type

[**InlineResponse2007**](InlineResponse2007.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETResourceResourceIdCommentsJson

> InlineResponse2007 GETResourceResourceIdCommentsJson(ctx, resourceId).Execute()

Retrieving Recent Comments



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
    resourceId := "resourceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommentsApi.GETResourceResourceIdCommentsJson(context.Background(), resourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommentsApi.GETResourceResourceIdCommentsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETResourceResourceIdCommentsJson`: InlineResponse2007
    fmt.Fprintf(os.Stdout, "Response from `CommentsApi.GETResourceResourceIdCommentsJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETResourceResourceIdCommentsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2007**](InlineResponse2007.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTResourceResourceIdCommentsJson

> InlineResponse201 POSTResourceResourceIdCommentsJson(ctx, resourceId).Body(body).Execute()

Creating a Comment



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
    resourceId := "resourceId_example" // string | 
    body := *openapiclient.NewInlineObject83() // InlineObject83 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommentsApi.POSTResourceResourceIdCommentsJson(context.Background(), resourceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommentsApi.POSTResourceResourceIdCommentsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTResourceResourceIdCommentsJson`: InlineResponse201
    fmt.Fprintf(os.Stdout, "Response from `CommentsApi.POSTResourceResourceIdCommentsJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTResourceResourceIdCommentsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject83**](InlineObject83.md) |  | 

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


## PUTCommentsIdJson

> InlineResponse2001 PUTCommentsIdJson(ctx, id).Body(body).Execute()

Updating a Comment



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
    id := int32(56) // int32 | 
    body := *openapiclient.NewInlineObject9(*openapiclient.NewCommentsIdJsonComment("Body_example")) // InlineObject9 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommentsApi.PUTCommentsIdJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommentsApi.PUTCommentsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTCommentsIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `CommentsApi.PUTCommentsIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTCommentsIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject9**](InlineObject9.md) |  | 

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


## PUTCommentsIdMarkreadJson

> InlineResponse2001 PUTCommentsIdMarkreadJson(ctx, id).Execute()

Mark a Comment as Read



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
    resp, r, err := api_client.CommentsApi.PUTCommentsIdMarkreadJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommentsApi.PUTCommentsIdMarkreadJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTCommentsIdMarkreadJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `CommentsApi.PUTCommentsIdMarkreadJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTCommentsIdMarkreadJsonRequest struct via the builder pattern


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

