# \FilesApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEFilesIdJson**](FilesApi.md#DELETEFilesIdJson) | **Delete** /files/{id}.json | Delete a File from a Project
[**GETFilesIdJson**](FilesApi.md#GETFilesIdJson) | **Get** /files/{id}.json | Get a Single File
[**GETFilesIdSharedlinkJson**](FilesApi.md#GETFilesIdSharedlinkJson) | **Get** /files/{id}/sharedlink.json | Get a short URL for sharing a File
[**GETProjectsIdFilesJson**](FilesApi.md#GETProjectsIdFilesJson) | **Get** /projects/{id}/files.json | List Files on a Project
[**GETTasksIdFilesJson**](FilesApi.md#GETTasksIdFilesJson) | **Get** /tasks/{id}/files.json | List Files on a Task
[**POSTFilesIdJson**](FilesApi.md#POSTFilesIdJson) | **Post** /files/{id}.json | Add a new File Version to a File
[**POSTProjectsIdFilesJson**](FilesApi.md#POSTProjectsIdFilesJson) | **Post** /projects/{id}/files.json | Add a File to a Project
[**POSTTasksIdFilesJson**](FilesApi.md#POSTTasksIdFilesJson) | **Post** /tasks/{id}/files.json | Add a File to a Task
[**PUTFilesIdCopyJson**](FilesApi.md#PUTFilesIdCopyJson) | **Put** /files/{id}/copy.json | Copy a File to another Project
[**PUTFilesIdMoveJson**](FilesApi.md#PUTFilesIdMoveJson) | **Put** /files/{id}/move.json | Move a file to another Project



## DELETEFilesIdJson

> InlineResponse2001 DELETEFilesIdJson(ctx, id).Execute()

Delete a File from a Project



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
    resp, r, err := api_client.FilesApi.DELETEFilesIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.DELETEFilesIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DELETEFilesIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.DELETEFilesIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEFilesIdJsonRequest struct via the builder pattern


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


## GETFilesIdJson

> map[string]interface{} GETFilesIdJson(ctx, id).Body(body).Execute()

Get a Single File



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
    body := *openapiclient.NewInlineObject17() // InlineObject17 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FilesApi.GETFilesIdJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.GETFilesIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETFilesIdJson`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.GETFilesIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETFilesIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject17**](InlineObject17.md) |  | 

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


## GETFilesIdSharedlinkJson

> InlineResponse20021 GETFilesIdSharedlinkJson(ctx, id).Version(version).Execute()

Get a short URL for sharing a File



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
    version := "version_example" // string | If you have multiple versions of a file, you can request a short link to other versions by appending 'version'. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FilesApi.GETFilesIdSharedlinkJson(context.Background(), id).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.GETFilesIdSharedlinkJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETFilesIdSharedlinkJson`: InlineResponse20021
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.GETFilesIdSharedlinkJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETFilesIdSharedlinkJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | If you have multiple versions of a file, you can request a short link to other versions by appending &#39;version&#39;. | 

### Return type

[**InlineResponse20021**](InlineResponse20021.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsIdFilesJson

> InlineResponse20065 GETProjectsIdFilesJson(ctx, id).Execute()

List Files on a Project



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
    resp, r, err := api_client.FilesApi.GETProjectsIdFilesJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.GETProjectsIdFilesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsIdFilesJson`: InlineResponse20065
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.GETProjectsIdFilesJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsIdFilesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20065**](InlineResponse20065.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETTasksIdFilesJson

> InlineResponse200103 GETTasksIdFilesJson(ctx, id).Execute()

List Files on a Task



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
    resp, r, err := api_client.FilesApi.GETTasksIdFilesJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.GETTasksIdFilesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETTasksIdFilesJson`: InlineResponse200103
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.GETTasksIdFilesJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETTasksIdFilesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200103**](InlineResponse200103.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTFilesIdJson

> InlineResponse201 POSTFilesIdJson(ctx, id).Body(body).Execute()

Add a new File Version to a File



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
    body := *openapiclient.NewInlineObject18() // InlineObject18 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FilesApi.POSTFilesIdJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.POSTFilesIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTFilesIdJson`: InlineResponse201
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.POSTFilesIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTFilesIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject18**](InlineObject18.md) |  | 

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


## POSTProjectsIdFilesJson

> InlineResponse201 POSTProjectsIdFilesJson(ctx, id).Body(body).Execute()

Add a File to a Project



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
    body := *openapiclient.NewInlineObject59() // InlineObject59 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FilesApi.POSTProjectsIdFilesJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.POSTProjectsIdFilesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsIdFilesJson`: InlineResponse201
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.POSTProjectsIdFilesJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsIdFilesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject59**](InlineObject59.md) |  | 

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


## POSTTasksIdFilesJson

> InlineResponse201 POSTTasksIdFilesJson(ctx, id).Body(body).Execute()

Add a File to a Task



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
    body := *openapiclient.NewInlineObject98() // InlineObject98 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FilesApi.POSTTasksIdFilesJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.POSTTasksIdFilesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTTasksIdFilesJson`: InlineResponse201
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.POSTTasksIdFilesJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTTasksIdFilesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject98**](InlineObject98.md) |  | 

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


## PUTFilesIdCopyJson

> InlineResponse2001 PUTFilesIdCopyJson(ctx, id).Body(body).Execute()

Copy a File to another Project



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
    body := *openapiclient.NewInlineObject19() // InlineObject19 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FilesApi.PUTFilesIdCopyJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.PUTFilesIdCopyJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTFilesIdCopyJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.PUTFilesIdCopyJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTFilesIdCopyJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject19**](InlineObject19.md) |  | 

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


## PUTFilesIdMoveJson

> PUTFilesIdMoveJson(ctx, id).Body(body).Execute()

Move a file to another Project



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
    body := *openapiclient.NewInlineObject20() // InlineObject20 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FilesApi.PUTFilesIdMoveJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.PUTFilesIdMoveJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTFilesIdMoveJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject20**](InlineObject20.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

