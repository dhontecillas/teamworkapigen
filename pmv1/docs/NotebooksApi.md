# \NotebooksApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETENotebooksIdJson**](NotebooksApi.md#DELETENotebooksIdJson) | **Delete** /notebooks/{id}.json | Delete a Single Notebook
[**GETNotebookCategoriesIdNotebooksJson**](NotebooksApi.md#GETNotebookCategoriesIdNotebooksJson) | **Get** /notebookCategories/{id}/notebooks.json | List Notebooks in a Specific Category
[**GETNotebooksIdJson**](NotebooksApi.md#GETNotebooksIdJson) | **Get** /notebooks/{id}.json | Get a Single Notebook
[**GETNotebooksJson**](NotebooksApi.md#GETNotebooksJson) | **Get** /notebooks.json | List All Notebooks
[**GETProjectsIdNotebooksJson**](NotebooksApi.md#GETProjectsIdNotebooksJson) | **Get** /projects/{id}/notebooks.json | List Notebooks on a Project
[**POSTProjectsIdNotebooksJson**](NotebooksApi.md#POSTProjectsIdNotebooksJson) | **Post** /projects/{id}/notebooks.json | Create a Single Notebook
[**PUTNotebooksIdCopyJson**](NotebooksApi.md#PUTNotebooksIdCopyJson) | **Put** /notebooks/{id}/copy.json | Copy a Notebook to another Project
[**PUTNotebooksIdJson**](NotebooksApi.md#PUTNotebooksIdJson) | **Put** /notebooks/{id}.json | Update a Single Notebook
[**PUTNotebooksIdLockJson**](NotebooksApi.md#PUTNotebooksIdLockJson) | **Put** /notebooks/{id}/lock.json | Lock a Single Notebook for Editing
[**PUTNotebooksIdMoveJson**](NotebooksApi.md#PUTNotebooksIdMoveJson) | **Put** /notebooks/{id}/move.json | Move a Notebook to another Project
[**PUTNotebooksIdUnlockJson**](NotebooksApi.md#PUTNotebooksIdUnlockJson) | **Put** /notebooks/{id}/unlock.json | Unlock a Single Notebook



## DELETENotebooksIdJson

> InlineResponse2001 DELETENotebooksIdJson(ctx, id).Execute()

Delete a Single Notebook



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
    resp, r, err := api_client.NotebooksApi.DELETENotebooksIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.DELETENotebooksIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DELETENotebooksIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `NotebooksApi.DELETENotebooksIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETENotebooksIdJsonRequest struct via the builder pattern


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


## GETNotebookCategoriesIdNotebooksJson

> InlineResponse20035 GETNotebookCategoriesIdNotebooksJson(ctx, id).IncludeContent(includeContent).Execute()

List Notebooks in a Specific Category



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
    includeContent := true // bool | You can pass includeContent=true to return the notebook HTML content with the notebook data (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotebooksApi.GETNotebookCategoriesIdNotebooksJson(context.Background(), id).IncludeContent(includeContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.GETNotebookCategoriesIdNotebooksJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETNotebookCategoriesIdNotebooksJson`: InlineResponse20035
    fmt.Fprintf(os.Stdout, "Response from `NotebooksApi.GETNotebookCategoriesIdNotebooksJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETNotebookCategoriesIdNotebooksJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeContent** | **bool** | You can pass includeContent&#x3D;true to return the notebook HTML content with the notebook data | 

### Return type

[**InlineResponse20035**](InlineResponse20035.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETNotebooksIdJson

> InlineResponse20036 GETNotebooksIdJson(ctx, id).Execute()

Get a Single Notebook



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
    resp, r, err := api_client.NotebooksApi.GETNotebooksIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.GETNotebooksIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETNotebooksIdJson`: InlineResponse20036
    fmt.Fprintf(os.Stdout, "Response from `NotebooksApi.GETNotebooksIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETNotebooksIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20036**](InlineResponse20036.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETNotebooksJson

> InlineResponse20035 GETNotebooksJson(ctx).IncludeContent(includeContent).Execute()

List All Notebooks



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
    includeContent := true // bool | You can pass includeContent=true to return the notebook HTML content with the notebook data (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotebooksApi.GETNotebooksJson(context.Background()).IncludeContent(includeContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.GETNotebooksJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETNotebooksJson`: InlineResponse20035
    fmt.Fprintf(os.Stdout, "Response from `NotebooksApi.GETNotebooksJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETNotebooksJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeContent** | **bool** | You can pass includeContent&#x3D;true to return the notebook HTML content with the notebook data | 

### Return type

[**InlineResponse20035**](InlineResponse20035.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsIdNotebooksJson

> InlineResponse20069 GETProjectsIdNotebooksJson(ctx, id).IncludeContent(includeContent).Execute()

List Notebooks on a Project



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
    includeContent := true // bool | You can pass includeContent=true to return the notebook HTML content with the notebook data (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotebooksApi.GETProjectsIdNotebooksJson(context.Background(), id).IncludeContent(includeContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.GETProjectsIdNotebooksJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsIdNotebooksJson`: InlineResponse20069
    fmt.Fprintf(os.Stdout, "Response from `NotebooksApi.GETProjectsIdNotebooksJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsIdNotebooksJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeContent** | **bool** | You can pass includeContent&#x3D;true to return the notebook HTML content with the notebook data | 

### Return type

[**InlineResponse20069**](InlineResponse20069.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsIdNotebooksJson

> InlineResponse20018 POSTProjectsIdNotebooksJson(ctx, id).Body(body).Execute()

Create a Single Notebook



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
    body := *openapiclient.NewInlineObject66() // InlineObject66 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotebooksApi.POSTProjectsIdNotebooksJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.POSTProjectsIdNotebooksJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsIdNotebooksJson`: InlineResponse20018
    fmt.Fprintf(os.Stdout, "Response from `NotebooksApi.POSTProjectsIdNotebooksJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsIdNotebooksJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject66**](InlineObject66.md) |  | 

### Return type

[**InlineResponse20018**](InlineResponse20018.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTNotebooksIdCopyJson

> InlineResponse201 PUTNotebooksIdCopyJson(ctx, id).Body(body).Execute()

Copy a Notebook to another Project



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
    body := *openapiclient.NewInlineObject34() // InlineObject34 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotebooksApi.PUTNotebooksIdCopyJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.PUTNotebooksIdCopyJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTNotebooksIdCopyJson`: InlineResponse201
    fmt.Fprintf(os.Stdout, "Response from `NotebooksApi.PUTNotebooksIdCopyJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTNotebooksIdCopyJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject34**](InlineObject34.md) |  | 

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


## PUTNotebooksIdJson

> InlineResponse2001 PUTNotebooksIdJson(ctx, id).Body(body).Execute()

Update a Single Notebook



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
    body := *openapiclient.NewInlineObject33() // InlineObject33 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotebooksApi.PUTNotebooksIdJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.PUTNotebooksIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTNotebooksIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `NotebooksApi.PUTNotebooksIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTNotebooksIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject33**](InlineObject33.md) |  | 

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


## PUTNotebooksIdLockJson

> InlineResponse2001 PUTNotebooksIdLockJson(ctx, id).Execute()

Lock a Single Notebook for Editing



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
    resp, r, err := api_client.NotebooksApi.PUTNotebooksIdLockJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.PUTNotebooksIdLockJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTNotebooksIdLockJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `NotebooksApi.PUTNotebooksIdLockJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTNotebooksIdLockJsonRequest struct via the builder pattern


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


## PUTNotebooksIdMoveJson

> InlineResponse2001 PUTNotebooksIdMoveJson(ctx, id).Body(body).Execute()

Move a Notebook to another Project



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
    body := *openapiclient.NewInlineObject35() // InlineObject35 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotebooksApi.PUTNotebooksIdMoveJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.PUTNotebooksIdMoveJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTNotebooksIdMoveJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `NotebooksApi.PUTNotebooksIdMoveJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTNotebooksIdMoveJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject35**](InlineObject35.md) |  | 

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


## PUTNotebooksIdUnlockJson

> InlineResponse2001 PUTNotebooksIdUnlockJson(ctx, id).Execute()

Unlock a Single Notebook



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
    resp, r, err := api_client.NotebooksApi.PUTNotebooksIdUnlockJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.PUTNotebooksIdUnlockJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTNotebooksIdUnlockJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `NotebooksApi.PUTNotebooksIdUnlockJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTNotebooksIdUnlockJsonRequest struct via the builder pattern


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

