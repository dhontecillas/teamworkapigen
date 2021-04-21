# \MilestonesApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEMilestonesIdJson**](MilestonesApi.md#DELETEMilestonesIdJson) | **Delete** /milestones/{id}.json | Delete a Milestone
[**GETMilestonesIdJson**](MilestonesApi.md#GETMilestonesIdJson) | **Get** /milestones/{id}.json | Get a Single Milestone
[**GETMilestonesJson**](MilestonesApi.md#GETMilestonesJson) | **Get** /milestones.json | List All Milestones
[**GETProjectsIdMilestonesJson**](MilestonesApi.md#GETProjectsIdMilestonesJson) | **Get** /projects/{id}/milestones.json | List Milestones on a Project
[**POSTProjectsIdMilestonesJson**](MilestonesApi.md#POSTProjectsIdMilestonesJson) | **Post** /projects/{id}/milestones.json | Create a Single Milestone
[**PUTMilestonesIdCompleteJson**](MilestonesApi.md#PUTMilestonesIdCompleteJson) | **Put** /milestones/{id}/complete.json | Complete a Milestone
[**PUTMilestonesIdJson**](MilestonesApi.md#PUTMilestonesIdJson) | **Put** /milestones/{id}.json | Update a Single Milestone
[**PUTMilestonesIdUncompleteJson**](MilestonesApi.md#PUTMilestonesIdUncompleteJson) | **Put** /milestones/{id}/uncomplete.json | Un-complete a Milestone



## DELETEMilestonesIdJson

> InlineResponse2001 DELETEMilestonesIdJson(ctx, id).Execute()

Delete a Milestone



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
    resp, r, err := api_client.MilestonesApi.DELETEMilestonesIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MilestonesApi.DELETEMilestonesIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DELETEMilestonesIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `MilestonesApi.DELETEMilestonesIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEMilestonesIdJsonRequest struct via the builder pattern


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


## GETMilestonesIdJson

> InlineResponse20034 GETMilestonesIdJson(ctx, id).ShowTaskLists(showTaskLists).GetProgress(getProgress).Execute()

Get a Single Milestone



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
    showTaskLists := true // bool | If you pass &showTaskLists=true you can also pass &showTasks=true to include a list of tasks in each list   (optional)
    getProgress := true // bool | Use this parameter to return a field percentageComplete (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MilestonesApi.GETMilestonesIdJson(context.Background(), id).ShowTaskLists(showTaskLists).GetProgress(getProgress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MilestonesApi.GETMilestonesIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETMilestonesIdJson`: InlineResponse20034
    fmt.Fprintf(os.Stdout, "Response from `MilestonesApi.GETMilestonesIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETMilestonesIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **showTaskLists** | **bool** | If you pass &amp;showTaskLists&#x3D;true you can also pass &amp;showTasks&#x3D;true to include a list of tasks in each list   | 
 **getProgress** | **bool** | Use this parameter to return a field percentageComplete | 

### Return type

[**InlineResponse20034**](InlineResponse20034.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETMilestonesJson

> InlineResponse20033 GETMilestonesJson(ctx).PageSize(pageSize).Page(page).Find(find).GetProgress(getProgress).ProjectType(projectType).Execute()

List All Milestones



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
    pageSize := int32(56) // int32 | The amount of Milestones returned can be limited using this parameter. Normally used in conjunction with the page parameter. (optional)
    page := int32(56) // int32 | Optionally, you can set the page from which to start retrieving results. This is usually used in conjunction with the parameter pageSize. For example: ?page=2&pageSize=10 will retrieve results 10-20. (optional)
    find := "find_example" // string | Find certain milestones. Options: al|,complete,incomplete,late,upcoming   (optional)
    getProgress := true // bool | Use this parameter to return a field percentageComplete (optional)
    projectType := "projectType_example" // string | use ALL to bring back milestones related to archived projects. (optional) (default to "ACTIVE")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MilestonesApi.GETMilestonesJson(context.Background()).PageSize(pageSize).Page(page).Find(find).GetProgress(getProgress).ProjectType(projectType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MilestonesApi.GETMilestonesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETMilestonesJson`: InlineResponse20033
    fmt.Fprintf(os.Stdout, "Response from `MilestonesApi.GETMilestonesJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETMilestonesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The amount of Milestones returned can be limited using this parameter. Normally used in conjunction with the page parameter. | 
 **page** | **int32** | Optionally, you can set the page from which to start retrieving results. This is usually used in conjunction with the parameter pageSize. For example: ?page&#x3D;2&amp;pageSize&#x3D;10 will retrieve results 10-20. | 
 **find** | **string** | Find certain milestones. Options: al|,complete,incomplete,late,upcoming   | 
 **getProgress** | **bool** | Use this parameter to return a field percentageComplete | 
 **projectType** | **string** | use ALL to bring back milestones related to archived projects. | [default to &quot;ACTIVE&quot;]

### Return type

[**InlineResponse20033**](InlineResponse20033.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsIdMilestonesJson

> InlineResponse20033 GETProjectsIdMilestonesJson(ctx, id).PageSize(pageSize).Page(page).Find(find).GetProgress(getProgress).Execute()

List Milestones on a Project



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
    pageSize := int32(56) // int32 | The amount of Milestones returned can be limited using this parameter. Normally used in conjunction with the page parameter. (optional)
    page := int32(56) // int32 | Optionally, you can set the page from which to start retrieving results. This is usually used in conjunction with the parameter pageSize. For example: ?page=2&pageSize=10 will retrieve results 10-20. (optional)
    find := "find_example" // string | Find certain milestones. Options: al|,complete,incomplete,late,upcoming   (optional)
    getProgress := true // bool | Use this parameter to return a field percentageComplete (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MilestonesApi.GETProjectsIdMilestonesJson(context.Background(), id).PageSize(pageSize).Page(page).Find(find).GetProgress(getProgress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MilestonesApi.GETProjectsIdMilestonesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsIdMilestonesJson`: InlineResponse20033
    fmt.Fprintf(os.Stdout, "Response from `MilestonesApi.GETProjectsIdMilestonesJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsIdMilestonesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The amount of Milestones returned can be limited using this parameter. Normally used in conjunction with the page parameter. | 
 **page** | **int32** | Optionally, you can set the page from which to start retrieving results. This is usually used in conjunction with the parameter pageSize. For example: ?page&#x3D;2&amp;pageSize&#x3D;10 will retrieve results 10-20. | 
 **find** | **string** | Find certain milestones. Options: al|,complete,incomplete,late,upcoming   | 
 **getProgress** | **bool** | Use this parameter to return a field percentageComplete | 

### Return type

[**InlineResponse20033**](InlineResponse20033.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsIdMilestonesJson

> InlineResponse2013 POSTProjectsIdMilestonesJson(ctx, id).Body(body).Execute()

Create a Single Milestone



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
    body := *openapiclient.NewInlineObject64() // InlineObject64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MilestonesApi.POSTProjectsIdMilestonesJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MilestonesApi.POSTProjectsIdMilestonesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsIdMilestonesJson`: InlineResponse2013
    fmt.Fprintf(os.Stdout, "Response from `MilestonesApi.POSTProjectsIdMilestonesJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsIdMilestonesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject64**](InlineObject64.md) |  | 

### Return type

[**InlineResponse2013**](InlineResponse2013.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTMilestonesIdCompleteJson

> InlineResponse2001 PUTMilestonesIdCompleteJson(ctx, id).Execute()

Complete a Milestone



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
    resp, r, err := api_client.MilestonesApi.PUTMilestonesIdCompleteJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MilestonesApi.PUTMilestonesIdCompleteJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTMilestonesIdCompleteJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `MilestonesApi.PUTMilestonesIdCompleteJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTMilestonesIdCompleteJsonRequest struct via the builder pattern


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


## PUTMilestonesIdJson

> InlineResponse2001 PUTMilestonesIdJson(ctx, id).Body(body).Execute()

Update a Single Milestone



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
    body := *openapiclient.NewInlineObject31() // InlineObject31 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MilestonesApi.PUTMilestonesIdJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MilestonesApi.PUTMilestonesIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTMilestonesIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `MilestonesApi.PUTMilestonesIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTMilestonesIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject31**](InlineObject31.md) |  | 

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


## PUTMilestonesIdUncompleteJson

> InlineResponse2001 PUTMilestonesIdUncompleteJson(ctx, id).Execute()

Un-complete a Milestone



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
    resp, r, err := api_client.MilestonesApi.PUTMilestonesIdUncompleteJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MilestonesApi.PUTMilestonesIdUncompleteJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTMilestonesIdUncompleteJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `MilestonesApi.PUTMilestonesIdUncompleteJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTMilestonesIdUncompleteJsonRequest struct via the builder pattern


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

