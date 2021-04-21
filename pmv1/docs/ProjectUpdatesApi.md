# \ProjectUpdatesApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEProjectsUpdatesUpdateIdJson**](ProjectUpdatesApi.md#DELETEProjectsUpdatesUpdateIdJson) | **Delete** /projects/updates/{updateId}.json | Delete a Project Update
[**GETProjectsProjIdUpdateJson**](ProjectUpdatesApi.md#GETProjectsProjIdUpdateJson) | **Get** /projects/{projId}/update.json | Get most recent Update
[**GETProjectsProjIdUpdatesJson**](ProjectUpdatesApi.md#GETProjectsProjIdUpdatesJson) | **Get** /projects/{projId}/updates.json | Get Project Updates
[**POSTProjectsProjIdUpdateJson**](ProjectUpdatesApi.md#POSTProjectsProjIdUpdateJson) | **Post** /projects/{projId}/update.json | Create a Project Update
[**PUTProjectsProjIdUpdateRequestJson**](ProjectUpdatesApi.md#PUTProjectsProjIdUpdateRequestJson) | **Put** /projects/{projId}/update/request.json | Requesting a Project Update from a Project Owner
[**PUTProjectsUpdatesUpdateIdJson**](ProjectUpdatesApi.md#PUTProjectsUpdatesUpdateIdJson) | **Put** /projects/updates/{updateId}.json | Modify a Project Update



## DELETEProjectsUpdatesUpdateIdJson

> InlineResponse20059 DELETEProjectsUpdatesUpdateIdJson(ctx, updateId).Execute()

Delete a Project Update



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectUpdatesApi.DELETEProjectsUpdatesUpdateIdJson(context.Background(), updateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectUpdatesApi.DELETEProjectsUpdatesUpdateIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DELETEProjectsUpdatesUpdateIdJson`: InlineResponse20059
    fmt.Fprintf(os.Stdout, "Response from `ProjectUpdatesApi.DELETEProjectsUpdatesUpdateIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**updateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEProjectsUpdatesUpdateIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20059**](InlineResponse20059.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsProjIdUpdateJson

> InlineResponse20080 GETProjectsProjIdUpdateJson(ctx, projId).Execute()

Get most recent Update



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
    projId := "projId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectUpdatesApi.GETProjectsProjIdUpdateJson(context.Background(), projId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectUpdatesApi.GETProjectsProjIdUpdateJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsProjIdUpdateJson`: InlineResponse20080
    fmt.Fprintf(os.Stdout, "Response from `ProjectUpdatesApi.GETProjectsProjIdUpdateJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsProjIdUpdateJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20080**](InlineResponse20080.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsProjIdUpdatesJson

> InlineResponse20080 GETProjectsProjIdUpdatesJson(ctx, projId).Execute()

Get Project Updates



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
    projId := "projId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectUpdatesApi.GETProjectsProjIdUpdatesJson(context.Background(), projId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectUpdatesApi.GETProjectsProjIdUpdatesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsProjIdUpdatesJson`: InlineResponse20080
    fmt.Fprintf(os.Stdout, "Response from `ProjectUpdatesApi.GETProjectsProjIdUpdatesJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsProjIdUpdatesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20080**](InlineResponse20080.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsProjIdUpdateJson

> InlineResponse20059 POSTProjectsProjIdUpdateJson(ctx, projId).Body(body).Execute()

Create a Project Update



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
    projId := "projId_example" // string | 
    body := *openapiclient.NewInlineObject78() // InlineObject78 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectUpdatesApi.POSTProjectsProjIdUpdateJson(context.Background(), projId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectUpdatesApi.POSTProjectsProjIdUpdateJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsProjIdUpdateJson`: InlineResponse20059
    fmt.Fprintf(os.Stdout, "Response from `ProjectUpdatesApi.POSTProjectsProjIdUpdateJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsProjIdUpdateJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject78**](InlineObject78.md) |  | 

### Return type

[**InlineResponse20059**](InlineResponse20059.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsProjIdUpdateRequestJson

> InlineResponse2001 PUTProjectsProjIdUpdateRequestJson(ctx, projId).Body(body).Execute()

Requesting a Project Update from a Project Owner



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
    projId := "projId_example" // string | 
    body := *openapiclient.NewInlineObject79() // InlineObject79 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectUpdatesApi.PUTProjectsProjIdUpdateRequestJson(context.Background(), projId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectUpdatesApi.PUTProjectsProjIdUpdateRequestJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsProjIdUpdateRequestJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `ProjectUpdatesApi.PUTProjectsProjIdUpdateRequestJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsProjIdUpdateRequestJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject79**](InlineObject79.md) |  | 

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


## PUTProjectsUpdatesUpdateIdJson

> InlineResponse20059 PUTProjectsUpdatesUpdateIdJson(ctx, updateId).Body(body).Execute()

Modify a Project Update



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
    body := *openapiclient.NewInlineObject53() // InlineObject53 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectUpdatesApi.PUTProjectsUpdatesUpdateIdJson(context.Background(), updateId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectUpdatesApi.PUTProjectsUpdatesUpdateIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsUpdatesUpdateIdJson`: InlineResponse20059
    fmt.Fprintf(os.Stdout, "Response from `ProjectUpdatesApi.PUTProjectsUpdatesUpdateIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**updateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsUpdatesUpdateIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject53**](InlineObject53.md) |  | 

### Return type

[**InlineResponse20059**](InlineResponse20059.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

