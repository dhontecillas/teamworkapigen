# \LikesApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETResourceResourceIdLikesJson**](LikesApi.md#GETResourceResourceIdLikesJson) | **Get** /{resource}/{resource_id}/likes.json | Get Likes on an Resource
[**PUTResourceResourceIdLikeJson**](LikesApi.md#PUTResourceResourceIdLikeJson) | **Put** /{resource}/{resource_id}/like.json | Like a Resource Item
[**PUTResourceResourceIdUnlikeJson**](LikesApi.md#PUTResourceResourceIdUnlikeJson) | **Put** /{resource}/{resource_id}/unlike.json | UnLike a Resource Item



## GETResourceResourceIdLikesJson

> InlineResponse200123 GETResourceResourceIdLikesJson(ctx, resource, resourceId).Execute()

Get Likes on an Resource



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
    resource := "resource_example" // string | Resources such as: Message Replies, FileVersions, Comments
    resourceId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LikesApi.GETResourceResourceIdLikesJson(context.Background(), resource, resourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LikesApi.GETResourceResourceIdLikesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETResourceResourceIdLikesJson`: InlineResponse200123
    fmt.Fprintf(os.Stdout, "Response from `LikesApi.GETResourceResourceIdLikesJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resource** | **string** | Resources such as: Message Replies, FileVersions, Comments | 
**resourceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETResourceResourceIdLikesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200123**](InlineResponse200123.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTResourceResourceIdLikeJson

> InlineResponse2001 PUTResourceResourceIdLikeJson(ctx, resource, resourceId).Execute()

Like a Resource Item



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
    resource := "resource_example" // string | Resources such as: Message Replies, FileVersions, Comments
    resourceId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LikesApi.PUTResourceResourceIdLikeJson(context.Background(), resource, resourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LikesApi.PUTResourceResourceIdLikeJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTResourceResourceIdLikeJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `LikesApi.PUTResourceResourceIdLikeJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resource** | **string** | Resources such as: Message Replies, FileVersions, Comments | 
**resourceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTResourceResourceIdLikeJsonRequest struct via the builder pattern


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


## PUTResourceResourceIdUnlikeJson

> InlineResponse2001 PUTResourceResourceIdUnlikeJson(ctx, resource, resourceId).Execute()

UnLike a Resource Item



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
    resource := "resource_example" // string | Resources such as: Message Replies, FileVersions, Comments
    resourceId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LikesApi.PUTResourceResourceIdUnlikeJson(context.Background(), resource, resourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LikesApi.PUTResourceResourceIdUnlikeJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTResourceResourceIdUnlikeJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `LikesApi.PUTResourceResourceIdUnlikeJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resource** | **string** | Resources such as: Message Replies, FileVersions, Comments | 
**resourceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTResourceResourceIdUnlikeJsonRequest struct via the builder pattern


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

