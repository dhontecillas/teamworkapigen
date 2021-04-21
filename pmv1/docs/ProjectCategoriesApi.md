# \ProjectCategoriesApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEProjectCategoriesIdJson**](ProjectCategoriesApi.md#DELETEProjectCategoriesIdJson) | **Delete** /projectCategories/{id}.json | Deleting a Project Category
[**GETProjectCategoriesCategoryIdTasksJson**](ProjectCategoriesApi.md#GETProjectCategoriesCategoryIdTasksJson) | **Get** /projectCategories/{categoryId}/tasks.json | Retrieving Tasks associated with the Category
[**GETProjectCategoriesIdJson**](ProjectCategoriesApi.md#GETProjectCategoriesIdJson) | **Get** /projectCategories/{id}.json | Retrieve a Single Project Category
[**GETProjectCategoriesJson**](ProjectCategoriesApi.md#GETProjectCategoriesJson) | **Get** /projectCategories.json | Retrieving all of a Project Categories
[**POSTProjectCategoriesJson**](ProjectCategoriesApi.md#POSTProjectCategoriesJson) | **Post** /projectCategories.json | Creating Project Categories
[**PUTProjectCategoriesIdJson**](ProjectCategoriesApi.md#PUTProjectCategoriesIdJson) | **Put** /projectCategories/{id}.json | Updating a Project Category



## DELETEProjectCategoriesIdJson

> InlineResponse2001 DELETEProjectCategoriesIdJson(ctx, id).Execute()

Deleting a Project Category



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
    resp, r, err := api_client.ProjectCategoriesApi.DELETEProjectCategoriesIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectCategoriesApi.DELETEProjectCategoriesIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DELETEProjectCategoriesIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `ProjectCategoriesApi.DELETEProjectCategoriesIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEProjectCategoriesIdJsonRequest struct via the builder pattern


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


## GETProjectCategoriesCategoryIdTasksJson

> InlineResponse20051 GETProjectCategoriesCategoryIdTasksJson(ctx, categoryId).Execute()

Retrieving Tasks associated with the Category



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
    categoryId := "categoryId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectCategoriesApi.GETProjectCategoriesCategoryIdTasksJson(context.Background(), categoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectCategoriesApi.GETProjectCategoriesCategoryIdTasksJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectCategoriesCategoryIdTasksJson`: InlineResponse20051
    fmt.Fprintf(os.Stdout, "Response from `ProjectCategoriesApi.GETProjectCategoriesCategoryIdTasksJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectCategoriesCategoryIdTasksJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20051**](InlineResponse20051.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectCategoriesIdJson

> InlineResponse20052 GETProjectCategoriesIdJson(ctx, id).Execute()

Retrieve a Single Project Category



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
    resp, r, err := api_client.ProjectCategoriesApi.GETProjectCategoriesIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectCategoriesApi.GETProjectCategoriesIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectCategoriesIdJson`: InlineResponse20052
    fmt.Fprintf(os.Stdout, "Response from `ProjectCategoriesApi.GETProjectCategoriesIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectCategoriesIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20052**](InlineResponse20052.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectCategoriesJson

> InlineResponse20050 GETProjectCategoriesJson(ctx).Execute()

Retrieving all of a Project Categories



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
    resp, r, err := api_client.ProjectCategoriesApi.GETProjectCategoriesJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectCategoriesApi.GETProjectCategoriesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectCategoriesJson`: InlineResponse20050
    fmt.Fprintf(os.Stdout, "Response from `ProjectCategoriesApi.GETProjectCategoriesJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectCategoriesJsonRequest struct via the builder pattern


### Return type

[**InlineResponse20050**](InlineResponse20050.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectCategoriesJson

> InlineResponse2012 POSTProjectCategoriesJson(ctx).Body(body).Execute()

Creating Project Categories



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
    body := *openapiclient.NewInlineObject48() // InlineObject48 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectCategoriesApi.POSTProjectCategoriesJson(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectCategoriesApi.POSTProjectCategoriesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectCategoriesJson`: InlineResponse2012
    fmt.Fprintf(os.Stdout, "Response from `ProjectCategoriesApi.POSTProjectCategoriesJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectCategoriesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InlineObject48**](InlineObject48.md) |  | 

### Return type

[**InlineResponse2012**](InlineResponse2012.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectCategoriesIdJson

> InlineResponse2001 PUTProjectCategoriesIdJson(ctx, id).Body(body).Execute()

Updating a Project Category



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
    body := *openapiclient.NewInlineObject49() // InlineObject49 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectCategoriesApi.PUTProjectCategoriesIdJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectCategoriesApi.PUTProjectCategoriesIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectCategoriesIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `ProjectCategoriesApi.PUTProjectCategoriesIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectCategoriesIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject49**](InlineObject49.md) |  | 

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

