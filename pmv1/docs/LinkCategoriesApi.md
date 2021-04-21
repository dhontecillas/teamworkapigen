# \LinkCategoriesApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETELinkCategoriesIdJson**](LinkCategoriesApi.md#DELETELinkCategoriesIdJson) | **Delete** /linkCategories/{id}.json | Deleting a Link Category
[**GETLinkCategoriesIdJson**](LinkCategoriesApi.md#GETLinkCategoriesIdJson) | **Get** /linkCategories/{id}.json | Retrieve a Single Link Category
[**GETProjectsIdLinkCategoriesJson**](LinkCategoriesApi.md#GETProjectsIdLinkCategoriesJson) | **Get** /projects/{id}/linkCategories.json | Retrieving all of a Link Categories
[**POSTProjectsIdLinkCategoriesJson**](LinkCategoriesApi.md#POSTProjectsIdLinkCategoriesJson) | **Post** /projects/{id}/linkCategories.json | Creating Link Categories
[**PUTLinkCategoriesIdJson**](LinkCategoriesApi.md#PUTLinkCategoriesIdJson) | **Put** /linkCategories/{id}.json | Updating a Link Category



## DELETELinkCategoriesIdJson

> InlineResponse2001 DELETELinkCategoriesIdJson(ctx, id).Execute()

Deleting a Link Category



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
    resp, r, err := api_client.LinkCategoriesApi.DELETELinkCategoriesIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinkCategoriesApi.DELETELinkCategoriesIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DELETELinkCategoriesIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `LinkCategoriesApi.DELETELinkCategoriesIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETELinkCategoriesIdJsonRequest struct via the builder pattern


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


## GETLinkCategoriesIdJson

> InlineResponse20020 GETLinkCategoriesIdJson(ctx, id).Execute()

Retrieve a Single Link Category



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
    resp, r, err := api_client.LinkCategoriesApi.GETLinkCategoriesIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinkCategoriesApi.GETLinkCategoriesIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETLinkCategoriesIdJson`: InlineResponse20020
    fmt.Fprintf(os.Stdout, "Response from `LinkCategoriesApi.GETLinkCategoriesIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETLinkCategoriesIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20020**](InlineResponse20020.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsIdLinkCategoriesJson

> InlineResponse20064 GETProjectsIdLinkCategoriesJson(ctx, id).Execute()

Retrieving all of a Link Categories



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
    resp, r, err := api_client.LinkCategoriesApi.GETProjectsIdLinkCategoriesJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinkCategoriesApi.GETProjectsIdLinkCategoriesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsIdLinkCategoriesJson`: InlineResponse20064
    fmt.Fprintf(os.Stdout, "Response from `LinkCategoriesApi.GETProjectsIdLinkCategoriesJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsIdLinkCategoriesJsonRequest struct via the builder pattern


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


## POSTProjectsIdLinkCategoriesJson

> InlineResponse201 POSTProjectsIdLinkCategoriesJson(ctx, id).Body(body).Execute()

Creating Link Categories



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
    body := *openapiclient.NewInlineObject61() // InlineObject61 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinkCategoriesApi.POSTProjectsIdLinkCategoriesJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinkCategoriesApi.POSTProjectsIdLinkCategoriesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsIdLinkCategoriesJson`: InlineResponse201
    fmt.Fprintf(os.Stdout, "Response from `LinkCategoriesApi.POSTProjectsIdLinkCategoriesJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsIdLinkCategoriesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject61**](InlineObject61.md) |  | 

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


## PUTLinkCategoriesIdJson

> map[string]interface{} PUTLinkCategoriesIdJson(ctx, id).Body(body).Execute()

Updating a Link Category



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
    body := *openapiclient.NewInlineObject23(*openapiclient.NewLinkCategoriesIdJsonCategory("Name_example")) // InlineObject23 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinkCategoriesApi.PUTLinkCategoriesIdJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinkCategoriesApi.PUTLinkCategoriesIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTLinkCategoriesIdJson`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LinkCategoriesApi.PUTLinkCategoriesIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTLinkCategoriesIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject23**](InlineObject23.md) |  | 

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

