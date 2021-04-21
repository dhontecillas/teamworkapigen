# \ActivityApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEActivityIdJson**](ActivityApi.md#DELETEActivityIdJson) | **Delete** /activity/{id}.json | Delete an Activity Entry
[**GETLatestActivityJson**](ActivityApi.md#GETLatestActivityJson) | **Get** /latestActivity.json | Latest Activity across all Projects
[**GETProjectsIdLatestActivityJson**](ActivityApi.md#GETProjectsIdLatestActivityJson) | **Get** /projects/{id}/latestActivity.json | List Latest Activity for a Specific Project
[**GETTasksIdAuditJson**](ActivityApi.md#GETTasksIdAuditJson) | **Get** /tasks/{id}/audit.json | Get a Task&#39;s Audit History (Premium and Enterprise plans only)
[**GETYoursiteTasksTaskIdActivityJson**](ActivityApi.md#GETYoursiteTasksTaskIdActivityJson) | **Get** /yoursite/tasks/{taskId}/activity.json | Get Task Activity



## DELETEActivityIdJson

> InlineResponse2001 DELETEActivityIdJson(ctx, id).Execute()

Delete an Activity Entry



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
    resp, r, err := api_client.ActivityApi.DELETEActivityIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.DELETEActivityIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DELETEActivityIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.DELETEActivityIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEActivityIdJsonRequest struct via the builder pattern


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


## GETLatestActivityJson

> InlineResponse20024 GETLatestActivityJson(ctx).Page(page).PageSize(pageSize).IncludeArchivedProjects(includeArchivedProjects).Execute()

Latest Activity across all Projects



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
    page := "page_example" // string | page=1&pageSize=10 (optional)
    pageSize := "pageSize_example" // string | page=1&pageSize=10 (optional)
    includeArchivedProjects := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActivityApi.GETLatestActivityJson(context.Background()).Page(page).PageSize(pageSize).IncludeArchivedProjects(includeArchivedProjects).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.GETLatestActivityJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETLatestActivityJson`: InlineResponse20024
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.GETLatestActivityJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETLatestActivityJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | page&#x3D;1&amp;pageSize&#x3D;10 | 
 **pageSize** | **string** | page&#x3D;1&amp;pageSize&#x3D;10 | 
 **includeArchivedProjects** | **bool** |  | 

### Return type

[**InlineResponse20024**](InlineResponse20024.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsIdLatestActivityJson

> InlineResponse20066 GETProjectsIdLatestActivityJson(ctx, id).MaxItems(maxItems).OnlyStarred(onlyStarred).Page(page).PageSize(pageSize).Execute()

List Latest Activity for a Specific Project



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
    maxItems := "maxItems_example" // string | default=60, Max=200 (optional)
    onlyStarred := true // bool |  (optional)
    page := "page_example" // string | page=1&pageSize=10 (optional)
    pageSize := "pageSize_example" // string | page=1&pageSize=10 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActivityApi.GETProjectsIdLatestActivityJson(context.Background(), id).MaxItems(maxItems).OnlyStarred(onlyStarred).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.GETProjectsIdLatestActivityJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsIdLatestActivityJson`: InlineResponse20066
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.GETProjectsIdLatestActivityJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsIdLatestActivityJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxItems** | **string** | default&#x3D;60, Max&#x3D;200 | 
 **onlyStarred** | **bool** |  | 
 **page** | **string** | page&#x3D;1&amp;pageSize&#x3D;10 | 
 **pageSize** | **string** | page&#x3D;1&amp;pageSize&#x3D;10 | 

### Return type

[**InlineResponse20066**](InlineResponse20066.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETTasksIdAuditJson

> InlineResponse200101 GETTasksIdAuditJson(ctx, id).Execute()

Get a Task's Audit History (Premium and Enterprise plans only)



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
    resp, r, err := api_client.ActivityApi.GETTasksIdAuditJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.GETTasksIdAuditJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETTasksIdAuditJson`: InlineResponse200101
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.GETTasksIdAuditJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETTasksIdAuditJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200101**](InlineResponse200101.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETYoursiteTasksTaskIdActivityJson

> InlineResponse200122 GETYoursiteTasksTaskIdActivityJson(ctx, taskId).Execute()

Get Task Activity



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
    taskId := "taskId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActivityApi.GETYoursiteTasksTaskIdActivityJson(context.Background(), taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.GETYoursiteTasksTaskIdActivityJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETYoursiteTasksTaskIdActivityJson`: InlineResponse200122
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.GETYoursiteTasksTaskIdActivityJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETYoursiteTasksTaskIdActivityJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200122**](InlineResponse200122.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

