# \NotebookCategoriesApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETENotebookCategoriesIdJson**](NotebookCategoriesApi.md#DELETENotebookCategoriesIdJson) | **Delete** /notebookCategories/{id}.json | Deleting a Notebook Category
[**GETNotebookCategoriesIdJson**](NotebookCategoriesApi.md#GETNotebookCategoriesIdJson) | **Get** /notebookCategories/{id}.json | Retrieve a Single Notebook Category
[**GETProjectsIdNotebookCategoriesJson**](NotebookCategoriesApi.md#GETProjectsIdNotebookCategoriesJson) | **Get** /projects/{id}/notebookCategories.json | Retrieving all of a Notebook Categories
[**POSTProjectsIdNotebookCategoriesJson**](NotebookCategoriesApi.md#POSTProjectsIdNotebookCategoriesJson) | **Post** /projects/{id}/notebookCategories.json | Creating Notebook Categories
[**PUTNotebookCategoriesIdJson**](NotebookCategoriesApi.md#PUTNotebookCategoriesIdJson) | **Put** /notebookCategories/{id}.json | Updating a Notebook Category



## DELETENotebookCategoriesIdJson

> InlineResponse2001 DELETENotebookCategoriesIdJson(ctx, id).Execute()

Deleting a Notebook Category



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
    resp, r, err := api_client.NotebookCategoriesApi.DELETENotebookCategoriesIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebookCategoriesApi.DELETENotebookCategoriesIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DELETENotebookCategoriesIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `NotebookCategoriesApi.DELETENotebookCategoriesIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETENotebookCategoriesIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETNotebookCategoriesIdJson

> InlineResponse20020 GETNotebookCategoriesIdJson(ctx, id).Execute()

Retrieve a Single Notebook Category



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
    resp, r, err := api_client.NotebookCategoriesApi.GETNotebookCategoriesIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebookCategoriesApi.GETNotebookCategoriesIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETNotebookCategoriesIdJson`: InlineResponse20020
    fmt.Fprintf(os.Stdout, "Response from `NotebookCategoriesApi.GETNotebookCategoriesIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETNotebookCategoriesIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20020**](InlineResponse20020.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsIdNotebookCategoriesJson

> InlineResponse20064 GETProjectsIdNotebookCategoriesJson(ctx, id).Execute()

Retrieving all of a Notebook Categories



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
    resp, r, err := api_client.NotebookCategoriesApi.GETProjectsIdNotebookCategoriesJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebookCategoriesApi.GETProjectsIdNotebookCategoriesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsIdNotebookCategoriesJson`: InlineResponse20064
    fmt.Fprintf(os.Stdout, "Response from `NotebookCategoriesApi.GETProjectsIdNotebookCategoriesJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsIdNotebookCategoriesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20064**](InlineResponse20064.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsIdNotebookCategoriesJson

> InlineResponse201 POSTProjectsIdNotebookCategoriesJson(ctx, id).Body(body).Execute()

Creating Notebook Categories



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
    body := *openapiclient.NewInlineObject65() // InlineObject65 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotebookCategoriesApi.POSTProjectsIdNotebookCategoriesJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebookCategoriesApi.POSTProjectsIdNotebookCategoriesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsIdNotebookCategoriesJson`: InlineResponse201
    fmt.Fprintf(os.Stdout, "Response from `NotebookCategoriesApi.POSTProjectsIdNotebookCategoriesJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsIdNotebookCategoriesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject65**](InlineObject65.md) |  | 

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


## PUTNotebookCategoriesIdJson

> InlineResponse2001 PUTNotebookCategoriesIdJson(ctx, id).Body(body).Execute()

Updating a Notebook Category



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
    body := *openapiclient.NewInlineObject32() // InlineObject32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotebookCategoriesApi.PUTNotebookCategoriesIdJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebookCategoriesApi.PUTNotebookCategoriesIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTNotebookCategoriesIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `NotebookCategoriesApi.PUTNotebookCategoriesIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTNotebookCategoriesIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject32**](InlineObject32.md) |  | 

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

