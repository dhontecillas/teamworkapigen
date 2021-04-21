# \AppsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3AppsJson**](AppsApi.md#GETProjectsApiV3AppsJson) | **Get** /projects/api/v3/apps.json | Get all apps.
[**GETProjectsApiV3AppsappIdJson**](AppsApi.md#GETProjectsApiV3AppsappIdJson) | **Get** /projects/api/v3/apps/{appId}.json | Get a specific app.
[**POSTProjectsApiV3AppsappIdInstallJson**](AppsApi.md#POSTProjectsApiV3AppsappIdInstallJson) | **Post** /projects/api/v3/apps/{appId}/install.json | Install an app onto an installation
[**POSTProjectsApiV3AppsappIdUninstallJson**](AppsApi.md#POSTProjectsApiV3AppsappIdUninstallJson) | **Post** /projects/api/v3/apps/{appId}/uninstall.json | Uninstall an app from an installation



## GETProjectsApiV3AppsJson

> AppAppsResponse GETProjectsApiV3AppsJson(ctx).Product(product).PageSize(pageSize).Page(page).Include(include).Execute()

Get all apps.

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
    product := "product_example" // string | request apps for a given product (optional) (default to "projects")
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    include := []string{"Inner_example"} // []string | include (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.GETProjectsApiV3AppsJson(context.Background()).Product(product).PageSize(pageSize).Page(page).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.GETProjectsApiV3AppsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3AppsJson`: AppAppsResponse
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.GETProjectsApiV3AppsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3AppsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **product** | **string** | request apps for a given product | [default to &quot;projects&quot;]
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **include** | **[]string** | include | 

### Return type

[**AppAppsResponse**](AppAppsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3AppsappIdJson

> AppResponse GETProjectsApiV3AppsappIdJson(ctx, appId).Execute()

Get a specific app.

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
    appId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.GETProjectsApiV3AppsappIdJson(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.GETProjectsApiV3AppsappIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3AppsappIdJson`: AppResponse
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.GETProjectsApiV3AppsappIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3AppsappIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppResponse**](AppResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsApiV3AppsappIdInstallJson

> AppResponse POSTProjectsApiV3AppsappIdInstallJson(ctx, appId).Execute()

Install an app onto an installation

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
    appId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.POSTProjectsApiV3AppsappIdInstallJson(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.POSTProjectsApiV3AppsappIdInstallJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsApiV3AppsappIdInstallJson`: AppResponse
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.POSTProjectsApiV3AppsappIdInstallJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3AppsappIdInstallJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppResponse**](AppResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsApiV3AppsappIdUninstallJson

> POSTProjectsApiV3AppsappIdUninstallJson(ctx, appId).Execute()

Uninstall an app from an installation

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
    appId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.POSTProjectsApiV3AppsappIdUninstallJson(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.POSTProjectsApiV3AppsappIdUninstallJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3AppsappIdUninstallJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

