# \TaskListsApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETETasklistsIdJson**](TaskListsApi.md#DELETETasklistsIdJson) | **Delete** /tasklists/{id}.json | Delete a Task List
[**GETProjectsIdTasklistsJson**](TaskListsApi.md#GETProjectsIdTasklistsJson) | **Get** /projects/{id}/tasklists.json | Get all Task Lists for a Project
[**GETTasklistsIdJson**](TaskListsApi.md#GETTasklistsIdJson) | **Get** /tasklists/{id}.json | Retrieve Single Task List
[**GETTasklistsJson**](TaskListsApi.md#GETTasklistsJson) | **Get** /tasklists.json | Get all Task Lists
[**GETTasklistsTemplatesJson**](TaskListsApi.md#GETTasklistsTemplatesJson) | **Get** /tasklists/templates.json | Template Task Lists: get all template task lists
[**POSTProjectsIdTasklistsJson**](TaskListsApi.md#POSTProjectsIdTasklistsJson) | **Post** /projects/{id}/tasklists.json | Create Task List
[**PUTProjectsIdTasklistsReorderJson**](TaskListsApi.md#PUTProjectsIdTasklistsReorderJson) | **Put** /projects/{id}/tasklists/reorder.json | Reorder Lists
[**PUTTasklistIdCopyJson**](TaskListsApi.md#PUTTasklistIdCopyJson) | **Put** /tasklist/{id}/copy.json | Copy a Task List, or Copy a Task List to another Project
[**PUTTasklistIdMoveJson**](TaskListsApi.md#PUTTasklistIdMoveJson) | **Put** /tasklist/{id}/move.json | Move a Task List to another Project
[**PUTTasklistsIdJson**](TaskListsApi.md#PUTTasklistsIdJson) | **Put** /tasklists/{id}.json | Update Task List



## DELETETasklistsIdJson

> InlineResponse2001 DELETETasklistsIdJson(ctx, id).Execute()

Delete a Task List



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
    resp, r, err := api_client.TaskListsApi.DELETETasklistsIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskListsApi.DELETETasklistsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DELETETasklistsIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `TaskListsApi.DELETETasklistsIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETETasklistsIdJsonRequest struct via the builder pattern


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


## GETProjectsIdTasklistsJson

> InlineResponse20076 GETProjectsIdTasklistsJson(ctx, id).ResponsiblePartyId(responsiblePartyId).GetOverdueCount(getOverdueCount).Status(status).ShowMilestones(showMilestones).GetCompletedCount(getCompletedCount).Filter(filter).Execute()

Get all Task Lists for a Project



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
    responsiblePartyId := "responsiblePartyId_example" // string |  (optional)
    getOverdueCount := true // bool | Passing getOverdueCount will return the number of overdue tasks overdue-count for each task list. (optional)
    status := "status_example" // string |  string Status: You can use the status option to restrict the Task Lists returned - valid values are all, active, and completed. The default is \"active\" (optional)
    showMilestones := true // bool | Passing showMilestones=1 will add Milestone information in to the response, if a Milestone is attached to the Task List   (optional)
    getCompletedCount := true // bool |  (optional)
    filter := "filter_example" // string | anytime, today, tomorrow, overdue, yesterday (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaskListsApi.GETProjectsIdTasklistsJson(context.Background(), id).ResponsiblePartyId(responsiblePartyId).GetOverdueCount(getOverdueCount).Status(status).ShowMilestones(showMilestones).GetCompletedCount(getCompletedCount).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskListsApi.GETProjectsIdTasklistsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsIdTasklistsJson`: InlineResponse20076
    fmt.Fprintf(os.Stdout, "Response from `TaskListsApi.GETProjectsIdTasklistsJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsIdTasklistsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **responsiblePartyId** | **string** |  | 
 **getOverdueCount** | **bool** | Passing getOverdueCount will return the number of overdue tasks overdue-count for each task list. | 
 **status** | **string** |  string Status: You can use the status option to restrict the Task Lists returned - valid values are all, active, and completed. The default is \&quot;active\&quot; | 
 **showMilestones** | **bool** | Passing showMilestones&#x3D;1 will add Milestone information in to the response, if a Milestone is attached to the Task List   | 
 **getCompletedCount** | **bool** |  | 
 **filter** | **string** | anytime, today, tomorrow, overdue, yesterday | 

### Return type

[**InlineResponse20076**](InlineResponse20076.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETTasklistsIdJson

> InlineResponse20096 GETTasklistsIdJson(ctx, id).Execute()

Retrieve Single Task List



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
    resp, r, err := api_client.TaskListsApi.GETTasklistsIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskListsApi.GETTasklistsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETTasklistsIdJson`: InlineResponse20096
    fmt.Fprintf(os.Stdout, "Response from `TaskListsApi.GETTasklistsIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETTasklistsIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20096**](InlineResponse20096.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETTasklistsJson

> InlineResponse20094 GETTasklistsJson(ctx).Status(status).Page(page).PageSize(pageSize).ShowDeleted(showDeleted).IncludeArchivedProjects(includeArchivedProjects).Execute()

Get all Task Lists



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
    status := "status_example" // string | string Status: You can use the status option to restrict the Task Lists returned - valid values are all, active, and completed. The default is \"active\"   (optional)
    page := "page_example" // string |  (optional)
    pageSize := "pageSize_example" // string |  (optional)
    showDeleted := true // bool |  (optional)
    includeArchivedProjects := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaskListsApi.GETTasklistsJson(context.Background()).Status(status).Page(page).PageSize(pageSize).ShowDeleted(showDeleted).IncludeArchivedProjects(includeArchivedProjects).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskListsApi.GETTasklistsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETTasklistsJson`: InlineResponse20094
    fmt.Fprintf(os.Stdout, "Response from `TaskListsApi.GETTasklistsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETTasklistsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | string Status: You can use the status option to restrict the Task Lists returned - valid values are all, active, and completed. The default is \&quot;active\&quot;   | 
 **page** | **string** |  | 
 **pageSize** | **string** |  | 
 **showDeleted** | **bool** |  | 
 **includeArchivedProjects** | **bool** |  | 

### Return type

[**InlineResponse20094**](InlineResponse20094.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETTasklistsTemplatesJson

> InlineResponse20095 GETTasklistsTemplatesJson(ctx).Execute()

Template Task Lists: get all template task lists



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
    resp, r, err := api_client.TaskListsApi.GETTasklistsTemplatesJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskListsApi.GETTasklistsTemplatesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETTasklistsTemplatesJson`: InlineResponse20095
    fmt.Fprintf(os.Stdout, "Response from `TaskListsApi.GETTasklistsTemplatesJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETTasklistsTemplatesJsonRequest struct via the builder pattern


### Return type

[**InlineResponse20095**](InlineResponse20095.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsIdTasklistsJson

> InlineResponse2014 POSTProjectsIdTasklistsJson(ctx, id).Body(body).Execute()

Create Task List



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
    body := *openapiclient.NewInlineObject73() // InlineObject73 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaskListsApi.POSTProjectsIdTasklistsJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskListsApi.POSTProjectsIdTasklistsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsIdTasklistsJson`: InlineResponse2014
    fmt.Fprintf(os.Stdout, "Response from `TaskListsApi.POSTProjectsIdTasklistsJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsIdTasklistsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject73**](InlineObject73.md) |  | 

### Return type

[**InlineResponse2014**](InlineResponse2014.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsIdTasklistsReorderJson

> InlineResponse2001 PUTProjectsIdTasklistsReorderJson(ctx, id).Body(body).Execute()

Reorder Lists



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
    body := *openapiclient.NewInlineObject74() // InlineObject74 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaskListsApi.PUTProjectsIdTasklistsReorderJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskListsApi.PUTProjectsIdTasklistsReorderJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsIdTasklistsReorderJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `TaskListsApi.PUTProjectsIdTasklistsReorderJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsIdTasklistsReorderJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject74**](InlineObject74.md) |  | 

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


## PUTTasklistIdCopyJson

> InlineResponse2001 PUTTasklistIdCopyJson(ctx, id).Body(body).Execute()

Copy a Task List, or Copy a Task List to another Project



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
    body := *openapiclient.NewInlineObject88() // InlineObject88 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaskListsApi.PUTTasklistIdCopyJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskListsApi.PUTTasklistIdCopyJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTTasklistIdCopyJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `TaskListsApi.PUTTasklistIdCopyJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTTasklistIdCopyJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject88**](InlineObject88.md) |  | 

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


## PUTTasklistIdMoveJson

> InlineResponse2001 PUTTasklistIdMoveJson(ctx, id).Body(body).Execute()

Move a Task List to another Project



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
    body := *openapiclient.NewInlineObject89("ProjectId_example") // InlineObject89 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaskListsApi.PUTTasklistIdMoveJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskListsApi.PUTTasklistIdMoveJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTTasklistIdMoveJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `TaskListsApi.PUTTasklistIdMoveJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTTasklistIdMoveJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject89**](InlineObject89.md) |  | 

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


## PUTTasklistsIdJson

> InlineResponse2001 PUTTasklistsIdJson(ctx, id).Body(body).Execute()

Update Task List



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
    body := *openapiclient.NewInlineObject90() // InlineObject90 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaskListsApi.PUTTasklistsIdJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskListsApi.PUTTasklistsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTTasklistsIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `TaskListsApi.PUTTasklistsIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTTasklistsIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject90**](InlineObject90.md) |  | 

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

