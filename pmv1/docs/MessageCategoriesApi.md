# \MessageCategoriesApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEMessageCategoriesIdJson**](MessageCategoriesApi.md#DELETEMessageCategoriesIdJson) | **Delete** /messageCategories/{id}.json | Deleting a Message Category
[**GETMessageCategoriesIdJson**](MessageCategoriesApi.md#GETMessageCategoriesIdJson) | **Get** /messageCategories/{id}.json | Retrieve a Single Message Category
[**GETProjectsIdMessageCategoriesJson**](MessageCategoriesApi.md#GETProjectsIdMessageCategoriesJson) | **Get** /projects/{id}/messageCategories.json | Retrieving all of a Message Categories
[**POSTProjectsIdMessageCategoriesJson**](MessageCategoriesApi.md#POSTProjectsIdMessageCategoriesJson) | **Post** /projects/{id}/messageCategories.json | Creating Message Categories
[**PUTMessageCategoriesIdJson**](MessageCategoriesApi.md#PUTMessageCategoriesIdJson) | **Put** /messageCategories/{id}.json | Updating a Category



## DELETEMessageCategoriesIdJson

> InlineResponse2001 DELETEMessageCategoriesIdJson(ctx, id).Execute()

Deleting a Message Category



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
    resp, r, err := api_client.MessageCategoriesApi.DELETEMessageCategoriesIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageCategoriesApi.DELETEMessageCategoriesIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DELETEMessageCategoriesIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `MessageCategoriesApi.DELETEMessageCategoriesIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEMessageCategoriesIdJsonRequest struct via the builder pattern


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


## GETMessageCategoriesIdJson

> InlineResponse20031 GETMessageCategoriesIdJson(ctx, id).Execute()

Retrieve a Single Message Category



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
    resp, r, err := api_client.MessageCategoriesApi.GETMessageCategoriesIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageCategoriesApi.GETMessageCategoriesIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETMessageCategoriesIdJson`: InlineResponse20031
    fmt.Fprintf(os.Stdout, "Response from `MessageCategoriesApi.GETMessageCategoriesIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETMessageCategoriesIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20031**](InlineResponse20031.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsIdMessageCategoriesJson

> InlineResponse20068 GETProjectsIdMessageCategoriesJson(ctx, id).Execute()

Retrieving all of a Message Categories



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
    resp, r, err := api_client.MessageCategoriesApi.GETProjectsIdMessageCategoriesJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageCategoriesApi.GETProjectsIdMessageCategoriesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsIdMessageCategoriesJson`: InlineResponse20068
    fmt.Fprintf(os.Stdout, "Response from `MessageCategoriesApi.GETProjectsIdMessageCategoriesJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsIdMessageCategoriesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20068**](InlineResponse20068.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsIdMessageCategoriesJson

> InlineResponse201 POSTProjectsIdMessageCategoriesJson(ctx, id).Body(body).Execute()

Creating Message Categories



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
    body := *openapiclient.NewInlineObject63(*openapiclient.NewProjectsIdMessageCategoriesJsonCategory("Name_example")) // InlineObject63 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MessageCategoriesApi.POSTProjectsIdMessageCategoriesJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageCategoriesApi.POSTProjectsIdMessageCategoriesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsIdMessageCategoriesJson`: InlineResponse201
    fmt.Fprintf(os.Stdout, "Response from `MessageCategoriesApi.POSTProjectsIdMessageCategoriesJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsIdMessageCategoriesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject63**](InlineObject63.md) |  | 

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


## PUTMessageCategoriesIdJson

> map[string]interface{} PUTMessageCategoriesIdJson(ctx, id).Body(body).Execute()

Updating a Category



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
    body := *openapiclient.NewInlineObject26(*openapiclient.NewMessageCategoriesIdJsonCategory("Name_example")) // InlineObject26 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MessageCategoriesApi.PUTMessageCategoriesIdJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageCategoriesApi.PUTMessageCategoriesIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTMessageCategoriesIdJson`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MessageCategoriesApi.PUTMessageCategoriesIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTMessageCategoriesIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject26**](InlineObject26.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

