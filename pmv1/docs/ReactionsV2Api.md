# \ReactionsV2Api

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETCommentsCommentIdReactionsJson**](ReactionsV2Api.md#GETCommentsCommentIdReactionsJson) | **Get** /comments/{commentId}/reactions.json | Get Reactions on a specific Comment
[**GETProjectsProjectIdUpdateJson**](ReactionsV2Api.md#GETProjectsProjectIdUpdateJson) | **Get** /projects/{projectId}/update.json | Get Reactions to a Project Update
[**GETV2MessagesMessageIdJson**](ReactionsV2Api.md#GETV2MessagesMessageIdJson) | **Get** /v/2/messages/{messageId}.json | Get Reactions to a Message
[**GETV2MessagesMessageIdRepliesJson**](ReactionsV2Api.md#GETV2MessagesMessageIdRepliesJson) | **Get** /v/2/messages/{messageId}/replies.json | Get Reactions on Message Replies
[**GETV2TasksIdCommentsJson**](ReactionsV2Api.md#GETV2TasksIdCommentsJson) | **Get** /v/2/tasks/{id}/comments.json | Get Comment Reactions on a task
[**PUTCommentsCommentIdReactJson**](ReactionsV2Api.md#PUTCommentsCommentIdReactJson) | **Put** /comments/{commentId}/react.json | React to a Comment
[**PUTMessagerepliesMessageIdReactJson**](ReactionsV2Api.md#PUTMessagerepliesMessageIdReactJson) | **Put** /messagereplies/{messageId}/react.json | React to a Message
[**PUTMessagerepliesMessagereplyIdReactJson**](ReactionsV2Api.md#PUTMessagerepliesMessagereplyIdReactJson) | **Put** /messagereplies/{messagereplyId}/react.json | React to a Message Reply
[**PUTProjectsUpdatesUpdateIdReactJson**](ReactionsV2Api.md#PUTProjectsUpdatesUpdateIdReactJson) | **Put** /projects/updates/{updateId}/react.json | React to a Project Update
[**PUTResourceIdUnreactJson**](ReactionsV2Api.md#PUTResourceIdUnreactJson) | **Put** /resource/{id}/unreact.json | Unreacting to a Comment/Message/ProjectUpdate



## GETCommentsCommentIdReactionsJson

> InlineResponse2009 GETCommentsCommentIdReactionsJson(ctx, commentId).ReactionType(reactionType).Execute()

Get Reactions on a specific Comment



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
    reactionType := "reactionType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReactionsV2Api.GETCommentsCommentIdReactionsJson(context.Background(), commentId).ReactionType(reactionType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsV2Api.GETCommentsCommentIdReactionsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCommentsCommentIdReactionsJson`: InlineResponse2009
    fmt.Fprintf(os.Stdout, "Response from `ReactionsV2Api.GETCommentsCommentIdReactionsJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCommentsCommentIdReactionsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reactionType** | **string** |  | 

### Return type

[**InlineResponse2009**](InlineResponse2009.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsProjectIdUpdateJson

> InlineResponse20080 GETProjectsProjectIdUpdateJson(ctx, projectId).GetReactions(getReactions).Body(body).Execute()

Get Reactions to a Project Update



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
    projectId := "projectId_example" // string | 
    getReactions := "getReactions_example" // string | 
    body := *openapiclient.NewInlineObject80() // InlineObject80 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReactionsV2Api.GETProjectsProjectIdUpdateJson(context.Background(), projectId).GetReactions(getReactions).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsV2Api.GETProjectsProjectIdUpdateJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsProjectIdUpdateJson`: InlineResponse20080
    fmt.Fprintf(os.Stdout, "Response from `ReactionsV2Api.GETProjectsProjectIdUpdateJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsProjectIdUpdateJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getReactions** | **string** |  | 
 **body** | [**InlineObject80**](InlineObject80.md) |  | 

### Return type

[**InlineResponse20080**](InlineResponse20080.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETV2MessagesMessageIdJson

> InlineResponse200115 GETV2MessagesMessageIdJson(ctx, messageId).GetReactions(getReactions).Execute()

Get Reactions to a Message



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
    messageId := "messageId_example" // string | 
    getReactions := true // bool | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReactionsV2Api.GETV2MessagesMessageIdJson(context.Background(), messageId).GetReactions(getReactions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsV2Api.GETV2MessagesMessageIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETV2MessagesMessageIdJson`: InlineResponse200115
    fmt.Fprintf(os.Stdout, "Response from `ReactionsV2Api.GETV2MessagesMessageIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETV2MessagesMessageIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getReactions** | **bool** |  | 

### Return type

[**InlineResponse200115**](InlineResponse200115.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETV2MessagesMessageIdRepliesJson

> InlineResponse200116 GETV2MessagesMessageIdRepliesJson(ctx, messageId).GetReactions(getReactions).Execute()

Get Reactions on Message Replies



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
    messageId := "messageId_example" // string | 
    getReactions := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReactionsV2Api.GETV2MessagesMessageIdRepliesJson(context.Background(), messageId).GetReactions(getReactions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsV2Api.GETV2MessagesMessageIdRepliesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETV2MessagesMessageIdRepliesJson`: InlineResponse200116
    fmt.Fprintf(os.Stdout, "Response from `ReactionsV2Api.GETV2MessagesMessageIdRepliesJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETV2MessagesMessageIdRepliesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getReactions** | **bool** |  | 

### Return type

[**InlineResponse200116**](InlineResponse200116.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETV2TasksIdCommentsJson

> map[string]interface{} GETV2TasksIdCommentsJson(ctx, id).GetReactions(getReactions).Page(page).PageSize(pageSize).GetLikes(getLikes).GetEmoji(getEmoji).OrderBy(orderBy).SortOrder(sortOrder).Execute()

Get Comment Reactions on a task



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
    getReactions := "getReactions_example" // string | 
    page := float32(8.14) // float32 |  (optional)
    pageSize := "pageSize_example" // string |  (optional)
    getLikes := "getLikes_example" // string |  (optional)
    getEmoji := "getEmoji_example" // string |  (optional)
    orderBy := "orderBy_example" // string | date (optional)
    sortOrder := "sortOrder_example" // string |  (optional) (default to "asc")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReactionsV2Api.GETV2TasksIdCommentsJson(context.Background(), id).GetReactions(getReactions).Page(page).PageSize(pageSize).GetLikes(getLikes).GetEmoji(getEmoji).OrderBy(orderBy).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsV2Api.GETV2TasksIdCommentsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETV2TasksIdCommentsJson`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ReactionsV2Api.GETV2TasksIdCommentsJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETV2TasksIdCommentsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getReactions** | **string** |  | 
 **page** | **float32** |  | 
 **pageSize** | **string** |  | 
 **getLikes** | **string** |  | 
 **getEmoji** | **string** |  | 
 **orderBy** | **string** | date | 
 **sortOrder** | **string** |  | [default to &quot;asc&quot;]

### Return type

**map[string]interface{}**

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTCommentsCommentIdReactJson

> InlineResponse2008 PUTCommentsCommentIdReactJson(ctx, commentId).Execute()

React to a Comment



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
    resp, r, err := api_client.ReactionsV2Api.PUTCommentsCommentIdReactJson(context.Background(), commentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsV2Api.PUTCommentsCommentIdReactJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTCommentsCommentIdReactJson`: InlineResponse2008
    fmt.Fprintf(os.Stdout, "Response from `ReactionsV2Api.PUTCommentsCommentIdReactJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTCommentsCommentIdReactJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2008**](InlineResponse2008.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTMessagerepliesMessageIdReactJson

> InlineResponse2008 PUTMessagerepliesMessageIdReactJson(ctx, messageId).Body(body).Execute()

React to a Message



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
    messageId := "messageId_example" // string | 
    body := *openapiclient.NewInlineObject28() // InlineObject28 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReactionsV2Api.PUTMessagerepliesMessageIdReactJson(context.Background(), messageId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsV2Api.PUTMessagerepliesMessageIdReactJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTMessagerepliesMessageIdReactJson`: InlineResponse2008
    fmt.Fprintf(os.Stdout, "Response from `ReactionsV2Api.PUTMessagerepliesMessageIdReactJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTMessagerepliesMessageIdReactJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject28**](InlineObject28.md) |  | 

### Return type

[**InlineResponse2008**](InlineResponse2008.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTMessagerepliesMessagereplyIdReactJson

> InlineResponse2008 PUTMessagerepliesMessagereplyIdReactJson(ctx, messagereplyId).Body(body).Execute()

React to a Message Reply



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
    messagereplyId := "messagereplyId_example" // string | 
    body := *openapiclient.NewInlineObject29() // InlineObject29 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReactionsV2Api.PUTMessagerepliesMessagereplyIdReactJson(context.Background(), messagereplyId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsV2Api.PUTMessagerepliesMessagereplyIdReactJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTMessagerepliesMessagereplyIdReactJson`: InlineResponse2008
    fmt.Fprintf(os.Stdout, "Response from `ReactionsV2Api.PUTMessagerepliesMessagereplyIdReactJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messagereplyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTMessagerepliesMessagereplyIdReactJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject29**](InlineObject29.md) |  | 

### Return type

[**InlineResponse2008**](InlineResponse2008.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsUpdatesUpdateIdReactJson

> InlineResponse2008 PUTProjectsUpdatesUpdateIdReactJson(ctx, updateId).Body(body).Execute()

React to a Project Update



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
    updateId := "updateId_example" // string | 
    body := *openapiclient.NewInlineObject54() // InlineObject54 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReactionsV2Api.PUTProjectsUpdatesUpdateIdReactJson(context.Background(), updateId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsV2Api.PUTProjectsUpdatesUpdateIdReactJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsUpdatesUpdateIdReactJson`: InlineResponse2008
    fmt.Fprintf(os.Stdout, "Response from `ReactionsV2Api.PUTProjectsUpdatesUpdateIdReactJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**updateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsUpdatesUpdateIdReactJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject54**](InlineObject54.md) |  | 

### Return type

[**InlineResponse2008**](InlineResponse2008.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTResourceIdUnreactJson

> InlineResponse20085 PUTResourceIdUnreactJson(ctx, id).Body(body).Execute()

Unreacting to a Comment/Message/ProjectUpdate



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
    body := *openapiclient.NewInlineObject82() // InlineObject82 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReactionsV2Api.PUTResourceIdUnreactJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsV2Api.PUTResourceIdUnreactJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTResourceIdUnreactJson`: InlineResponse20085
    fmt.Fprintf(os.Stdout, "Response from `ReactionsV2Api.PUTResourceIdUnreactJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTResourceIdUnreactJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject82**](InlineObject82.md) |  | 

### Return type

[**InlineResponse20085**](InlineResponse20085.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

