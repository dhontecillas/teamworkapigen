# \MessagesApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEPostsIdJson**](MessagesApi.md#DELETEPostsIdJson) | **Delete** /posts/{id}.json | Delete Message
[**GETPostsIdJson**](MessagesApi.md#GETPostsIdJson) | **Get** /posts/{id}.json | Retrieve a Single Message
[**GETPostsJson**](MessagesApi.md#GETPostsJson) | **Get** /posts.json | Retrieve Latest Messages
[**GETProjectsIdPostsArchiveJson**](MessagesApi.md#GETProjectsIdPostsArchiveJson) | **Get** /projects/{id}/posts/archive.json | Get Archived Messages
[**GETProjectsIdPostsJson**](MessagesApi.md#GETProjectsIdPostsJson) | **Get** /projects/{id}/posts.json | Retrieve Latest Messages by Project
[**GETProjectsProjectIdCatCategoryIdPostsArchiveJson**](MessagesApi.md#GETProjectsProjectIdCatCategoryIdPostsArchiveJson) | **Get** /projects/{project_id}/cat/{category_id}/posts/archive.json | Get Archived Messages by Category
[**GETProjectsProjectIdCatCategoryIdPostsJson**](MessagesApi.md#GETProjectsProjectIdCatCategoryIdPostsJson) | **Get** /projects/{project_id}/cat/{category_id}/posts.json | Retrieve Messages by Category
[**POSTProjectsIdPostsJson**](MessagesApi.md#POSTProjectsIdPostsJson) | **Post** /projects/{id}/posts.json | Create a Message
[**PUTMessagesIdArchiveJson**](MessagesApi.md#PUTMessagesIdArchiveJson) | **Put** /messages/{id}/archive.json | Archive a Message
[**PUTMessagesIdMarkreadJson**](MessagesApi.md#PUTMessagesIdMarkreadJson) | **Put** /messages/{id}/markread.json | Mark Message As Read
[**PUTMessagesIdUnarchiveJson**](MessagesApi.md#PUTMessagesIdUnarchiveJson) | **Put** /messages/{id}/unarchive.json | Un-archive a Message
[**PUTPostsIdJson**](MessagesApi.md#PUTPostsIdJson) | **Put** /posts/{id}.json | Update Message



## DELETEPostsIdJson

> InlineResponse2001 DELETEPostsIdJson(ctx, id).Execute()

Delete Message



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
    resp, r, err := api_client.MessagesApi.DELETEPostsIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagesApi.DELETEPostsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DELETEPostsIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `MessagesApi.DELETEPostsIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEPostsIdJsonRequest struct via the builder pattern


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


## GETPostsIdJson

> InlineResponse20049 GETPostsIdJson(ctx, id).Execute()

Retrieve a Single Message



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
    resp, r, err := api_client.MessagesApi.GETPostsIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagesApi.GETPostsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPostsIdJson`: InlineResponse20049
    fmt.Fprintf(os.Stdout, "Response from `MessagesApi.GETPostsIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPostsIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20049**](InlineResponse20049.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPostsJson

> InlineResponse20048 GETPostsJson(ctx).Execute()

Retrieve Latest Messages



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MessagesApi.GETPostsJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagesApi.GETPostsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPostsJson`: InlineResponse20048
    fmt.Fprintf(os.Stdout, "Response from `MessagesApi.GETPostsJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETPostsJsonRequest struct via the builder pattern


### Return type

[**InlineResponse20048**](InlineResponse20048.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsIdPostsArchiveJson

> InlineResponse20048 GETProjectsIdPostsArchiveJson(ctx, id).Execute()

Get Archived Messages



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
    resp, r, err := api_client.MessagesApi.GETProjectsIdPostsArchiveJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagesApi.GETProjectsIdPostsArchiveJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsIdPostsArchiveJson`: InlineResponse20048
    fmt.Fprintf(os.Stdout, "Response from `MessagesApi.GETProjectsIdPostsArchiveJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsIdPostsArchiveJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20048**](InlineResponse20048.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsIdPostsJson

> InlineResponse20048 GETProjectsIdPostsJson(ctx, id).Execute()

Retrieve Latest Messages by Project



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
    resp, r, err := api_client.MessagesApi.GETProjectsIdPostsJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagesApi.GETProjectsIdPostsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsIdPostsJson`: InlineResponse20048
    fmt.Fprintf(os.Stdout, "Response from `MessagesApi.GETProjectsIdPostsJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsIdPostsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20048**](InlineResponse20048.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsProjectIdCatCategoryIdPostsArchiveJson

> InlineResponse20082 GETProjectsProjectIdCatCategoryIdPostsArchiveJson(ctx, projectId, categoryId).Execute()

Get Archived Messages by Category



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
    categoryId := "categoryId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MessagesApi.GETProjectsProjectIdCatCategoryIdPostsArchiveJson(context.Background(), projectId, categoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagesApi.GETProjectsProjectIdCatCategoryIdPostsArchiveJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsProjectIdCatCategoryIdPostsArchiveJson`: InlineResponse20082
    fmt.Fprintf(os.Stdout, "Response from `MessagesApi.GETProjectsProjectIdCatCategoryIdPostsArchiveJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**categoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsProjectIdCatCategoryIdPostsArchiveJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20082**](InlineResponse20082.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsProjectIdCatCategoryIdPostsJson

> InlineResponse20082 GETProjectsProjectIdCatCategoryIdPostsJson(ctx, projectId, categoryId).Execute()

Retrieve Messages by Category



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
    categoryId := "categoryId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MessagesApi.GETProjectsProjectIdCatCategoryIdPostsJson(context.Background(), projectId, categoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagesApi.GETProjectsProjectIdCatCategoryIdPostsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsProjectIdCatCategoryIdPostsJson`: InlineResponse20082
    fmt.Fprintf(os.Stdout, "Response from `MessagesApi.GETProjectsProjectIdCatCategoryIdPostsJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**categoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsProjectIdCatCategoryIdPostsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20082**](InlineResponse20082.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsIdPostsJson

> InlineResponse201 POSTProjectsIdPostsJson(ctx, id).Body(body).Execute()

Create a Message



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
    body := *openapiclient.NewInlineObject69(*openapiclient.NewProjectsIdPostsJsonPost("Body_example", "Title_example")) // InlineObject69 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MessagesApi.POSTProjectsIdPostsJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagesApi.POSTProjectsIdPostsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsIdPostsJson`: InlineResponse201
    fmt.Fprintf(os.Stdout, "Response from `MessagesApi.POSTProjectsIdPostsJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsIdPostsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject69**](InlineObject69.md) |  | 

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


## PUTMessagesIdArchiveJson

> InlineResponse2001 PUTMessagesIdArchiveJson(ctx, id).Execute()

Archive a Message



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
    resp, r, err := api_client.MessagesApi.PUTMessagesIdArchiveJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagesApi.PUTMessagesIdArchiveJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTMessagesIdArchiveJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `MessagesApi.PUTMessagesIdArchiveJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTMessagesIdArchiveJsonRequest struct via the builder pattern


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


## PUTMessagesIdMarkreadJson

> InlineResponse2001 PUTMessagesIdMarkreadJson(ctx, id).Execute()

Mark Message As Read



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
    resp, r, err := api_client.MessagesApi.PUTMessagesIdMarkreadJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagesApi.PUTMessagesIdMarkreadJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTMessagesIdMarkreadJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `MessagesApi.PUTMessagesIdMarkreadJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTMessagesIdMarkreadJsonRequest struct via the builder pattern


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


## PUTMessagesIdUnarchiveJson

> InlineResponse2001 PUTMessagesIdUnarchiveJson(ctx, id).Execute()

Un-archive a Message



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
    resp, r, err := api_client.MessagesApi.PUTMessagesIdUnarchiveJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagesApi.PUTMessagesIdUnarchiveJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTMessagesIdUnarchiveJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `MessagesApi.PUTMessagesIdUnarchiveJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTMessagesIdUnarchiveJsonRequest struct via the builder pattern


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


## PUTPostsIdJson

> InlineResponse2001 PUTPostsIdJson(ctx, id).Body(body).Execute()

Update Message



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
    body := *openapiclient.NewInlineObject47(*openapiclient.NewPostsIdJsonPost("Title_example")) // InlineObject47 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MessagesApi.PUTPostsIdJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagesApi.PUTPostsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTPostsIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `MessagesApi.PUTPostsIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTPostsIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject47**](InlineObject47.md) |  | 

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

