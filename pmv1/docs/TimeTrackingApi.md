# \TimeTrackingApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETETimeEntriesIdJson**](TimeTrackingApi.md#DELETETimeEntriesIdJson) | **Delete** /time_entries/{id}.json | Delete a Time-Entry
[**GETMeTimersJson**](TimeTrackingApi.md#GETMeTimersJson) | **Get** /me/timers.json | Get all your Running Timers
[**GETPeopleIdLoggedtimeJson**](TimeTrackingApi.md#GETPeopleIdLoggedtimeJson) | **Get** /people/{id}/loggedtime.json | Get Logged Time by Person
[**GETPeopleIdTimersJson**](TimeTrackingApi.md#GETPeopleIdTimersJson) | **Get** /people/{id}/timers.json | Get all Running Timers for a specific Person
[**GETProjectsEstimatedtimeTotalJson**](TimeTrackingApi.md#GETProjectsEstimatedtimeTotalJson) | **Get** /projects/estimatedtime/total.json. | Estimated Time Totals on Projects
[**GETProjectsIdTimeEntriesJson**](TimeTrackingApi.md#GETProjectsIdTimeEntriesJson) | **Get** /projects/{id}/time_entries.json | Retrieve all Time Entries for a Project
[**GETProjectsIdTimeTotalJson**](TimeTrackingApi.md#GETProjectsIdTimeTotalJson) | **Get** /projects/{id}/time/total.json | Time Totals on a Project
[**GETProjectsTimeTotalJson**](TimeTrackingApi.md#GETProjectsTimeTotalJson) | **Get** /projects/time/total.json | Time Totals across Projects
[**GETTasklistsIdTimeTotalJson**](TimeTrackingApi.md#GETTasklistsIdTimeTotalJson) | **Get** /tasklists/{id}/time/total.json | Total Time on a Tasklist
[**GETTasksIdTimeEntriesJson**](TimeTrackingApi.md#GETTasksIdTimeEntriesJson) | **Get** /tasks/{id}/time_entries.json | Retrieve all Task times
[**GETTasksIdTimeTotalJson**](TimeTrackingApi.md#GETTasksIdTimeTotalJson) | **Get** /tasks/{id}/time/total.json | Total Time on a Task
[**GETTimeEntriesIdJson**](TimeTrackingApi.md#GETTimeEntriesIdJson) | **Get** /time_entries/{id}.json | Retrieve single Time-Entry
[**GETTimeEntriesJson**](TimeTrackingApi.md#GETTimeEntriesJson) | **Get** /time_entries.json | Retrieve all Time Entries across all Projects
[**GETTimeTotalJson**](TimeTrackingApi.md#GETTimeTotalJson) | **Get** /time/total.json | Get Total Time across an Account
[**GETTimersJson**](TimeTrackingApi.md#GETTimersJson) | **Get** /timers.json | Get all Running Timers
[**POSTProjectsIdTimeEntriesJson**](TimeTrackingApi.md#POSTProjectsIdTimeEntriesJson) | **Post** /projects/{id}/time_entries.json | Create a Time-Entry
[**POSTTasksIdTimeEntriesJson**](TimeTrackingApi.md#POSTTasksIdTimeEntriesJson) | **Post** /tasks/{id}/time_entries.json | Create a Time-Entry (for a Task)
[**PUTTasksIdEstimatedtimeJson**](TimeTrackingApi.md#PUTTasksIdEstimatedtimeJson) | **Put** /tasks/{id}/estimatedtime.json | Add a Time estimate to a Task
[**PUTTimeEntriesIdJson**](TimeTrackingApi.md#PUTTimeEntriesIdJson) | **Put** /time_entries/{id}.json | Update a Time Entry



## DELETETimeEntriesIdJson

> InlineResponse2001 DELETETimeEntriesIdJson(ctx, id).Execute()

Delete a Time-Entry



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
    resp, r, err := api_client.TimeTrackingApi.DELETETimeEntriesIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.DELETETimeEntriesIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DELETETimeEntriesIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.DELETETimeEntriesIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETETimeEntriesIdJsonRequest struct via the builder pattern


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


## GETMeTimersJson

> InlineResponse20030 GETMeTimersJson(ctx).ProjectId(projectId).TaskId(taskId).RunningTimersOnly(runningTimersOnly).Execute()

Get all your Running Timers



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
    projectId := "projectId_example" // string | timers on a given project (optional)
    taskId := "taskId_example" // string | Timers on a task (optional)
    runningTimersOnly := true // bool | all active timers or only timers which are currently running (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.GETMeTimersJson(context.Background()).ProjectId(projectId).TaskId(taskId).RunningTimersOnly(runningTimersOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.GETMeTimersJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETMeTimersJson`: InlineResponse20030
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.GETMeTimersJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETMeTimersJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string** | timers on a given project | 
 **taskId** | **string** | Timers on a task | 
 **runningTimersOnly** | **bool** | all active timers or only timers which are currently running | 

### Return type

[**InlineResponse20030**](InlineResponse20030.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPeopleIdLoggedtimeJson

> InlineResponse20040 GETPeopleIdLoggedtimeJson(ctx, id).M(m).Y(y).Page(page).PageSize(pageSize).Execute()

Get Logged Time by Person



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
    m := "m_example" // string | Month of the year. Needs to be used with 'y' param in order to work. (optional)
    y := "y_example" // string | year (optional) (default to "YYYY")
    page := "page_example" // string |  (optional)
    pageSize := "pageSize_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.GETPeopleIdLoggedtimeJson(context.Background(), id).M(m).Y(y).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.GETPeopleIdLoggedtimeJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPeopleIdLoggedtimeJson`: InlineResponse20040
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.GETPeopleIdLoggedtimeJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPeopleIdLoggedtimeJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **m** | **string** | Month of the year. Needs to be used with &#39;y&#39; param in order to work. | 
 **y** | **string** | year | [default to &quot;YYYY&quot;]
 **page** | **string** |  | 
 **pageSize** | **string** |  | 

### Return type

[**InlineResponse20040**](InlineResponse20040.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPeopleIdTimersJson

> InlineResponse20030 GETPeopleIdTimersJson(ctx, id).ProjectId(projectId).TaskId(taskId).RunningTimersOnly(runningTimersOnly).Execute()

Get all Running Timers for a specific Person



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
    projectId := "projectId_example" // string | timers on a given project (optional)
    taskId := "taskId_example" // string | Timers on a task (optional)
    runningTimersOnly := true // bool | all active timers or only timers which are currently running (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.GETPeopleIdTimersJson(context.Background(), id).ProjectId(projectId).TaskId(taskId).RunningTimersOnly(runningTimersOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.GETPeopleIdTimersJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPeopleIdTimersJson`: InlineResponse20030
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.GETPeopleIdTimersJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPeopleIdTimersJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectId** | **string** | timers on a given project | 
 **taskId** | **string** | Timers on a task | 
 **runningTimersOnly** | **bool** | all active timers or only timers which are currently running | 

### Return type

[**InlineResponse20030**](InlineResponse20030.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsEstimatedtimeTotalJson

> InlineResponse20057 GETProjectsEstimatedtimeTotalJson(ctx).Execute()

Estimated Time Totals on Projects



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
    resp, r, err := api_client.TimeTrackingApi.GETProjectsEstimatedtimeTotalJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.GETProjectsEstimatedtimeTotalJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsEstimatedtimeTotalJson`: InlineResponse20057
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.GETProjectsEstimatedtimeTotalJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsEstimatedtimeTotalJsonRequest struct via the builder pattern


### Return type

[**InlineResponse20057**](InlineResponse20057.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsIdTimeEntriesJson

> InlineResponse20079 GETProjectsIdTimeEntriesJson(ctx, id).Page(page).Fromdate(fromdate).Fromtime(fromtime).Todate(todate).Totime(totime).Sortby(sortby).Sortorder(sortorder).UserId(userId).BillableType(billableType).InvoicedType(invoicedType).ProjectType(projectType).ShowDeleted(showDeleted).TagIds(tagIds).UpdatedAfterDate(updatedAfterDate).PageSize(pageSize).TaskTagIds(taskTagIds).Execute()

Retrieve all Time Entries for a Project



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
    page := int32(56) // int32 | The page to start retrieving entries from ( e.g: page=1 gives records 1 - 500, page=2 gives records 501-1001 etc) (optional)
    fromdate := "fromdate_example" // string | string (YYYYMMDD) - The start date to retrieve from (optional)
    fromtime := "fromtime_example" // string | string (HH:MM) - The start time only if fromdate is passed (optional)
    todate := "todate_example" // string | string (YYYYMMDD) - The end date to retrieve to (optional)
    totime := "totime_example" // string | string (HH:MM) - The end time only if todate is passed (optional)
    sortby := "sortby_example" // string | string - One of the following: date, user, task, tasklist, project, company, dateupdated (Default: date) (optional)
    sortorder := "sortorder_example" // string | ASC or DESC - The order to sort the returned data (optional)
    userId := int32(56) // int32 | Return time logs for a specific user only (optional)
    billableType := "billableType_example" // string | string (billable or nonbillable).Filter the Time Entries to those that are Billable or Not Billable. (optional)
    invoicedType := "invoicedType_example" // string | string (invoiced or noninvoiced) - filter the time entries to those that have been Invoiced or not. (optional)
    projectType := "projectType_example" // string | string (all, active, archived) - Filter the time entries to those in Active projects, Archived projects or All projects. (optional)
    showDeleted := "showDeleted_example" // string | string (true or false) - Filter time entries to include deleted time sheet entries or not (optional)
    tagIds := "tagIds_example" // string | numeric - Include one or more Tag IDs here to return only the time entries with those tags attached (eg &tagIds=23,445,454) (optional)
    updatedAfterDate := "updatedAfterDate_example" // string | string YYYYMMDDHHMMSS (optional)
    pageSize := "pageSize_example" // string | pageSize=500 (optional)
    taskTagIds := "taskTagIds_example" // string | filter on task tags (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.GETProjectsIdTimeEntriesJson(context.Background(), id).Page(page).Fromdate(fromdate).Fromtime(fromtime).Todate(todate).Totime(totime).Sortby(sortby).Sortorder(sortorder).UserId(userId).BillableType(billableType).InvoicedType(invoicedType).ProjectType(projectType).ShowDeleted(showDeleted).TagIds(tagIds).UpdatedAfterDate(updatedAfterDate).PageSize(pageSize).TaskTagIds(taskTagIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.GETProjectsIdTimeEntriesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsIdTimeEntriesJson`: InlineResponse20079
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.GETProjectsIdTimeEntriesJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsIdTimeEntriesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page to start retrieving entries from ( e.g: page&#x3D;1 gives records 1 - 500, page&#x3D;2 gives records 501-1001 etc) | 
 **fromdate** | **string** | string (YYYYMMDD) - The start date to retrieve from | 
 **fromtime** | **string** | string (HH:MM) - The start time only if fromdate is passed | 
 **todate** | **string** | string (YYYYMMDD) - The end date to retrieve to | 
 **totime** | **string** | string (HH:MM) - The end time only if todate is passed | 
 **sortby** | **string** | string - One of the following: date, user, task, tasklist, project, company, dateupdated (Default: date) | 
 **sortorder** | **string** | ASC or DESC - The order to sort the returned data | 
 **userId** | **int32** | Return time logs for a specific user only | 
 **billableType** | **string** | string (billable or nonbillable).Filter the Time Entries to those that are Billable or Not Billable. | 
 **invoicedType** | **string** | string (invoiced or noninvoiced) - filter the time entries to those that have been Invoiced or not. | 
 **projectType** | **string** | string (all, active, archived) - Filter the time entries to those in Active projects, Archived projects or All projects. | 
 **showDeleted** | **string** | string (true or false) - Filter time entries to include deleted time sheet entries or not | 
 **tagIds** | **string** | numeric - Include one or more Tag IDs here to return only the time entries with those tags attached (eg &amp;tagIds&#x3D;23,445,454) | 
 **updatedAfterDate** | **string** | string YYYYMMDDHHMMSS | 
 **pageSize** | **string** | pageSize&#x3D;500 | 
 **taskTagIds** | **string** | filter on task tags | 

### Return type

[**InlineResponse20079**](InlineResponse20079.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsIdTimeTotalJson

> InlineResponse20078 GETProjectsIdTimeTotalJson(ctx, id).FromDate(fromDate).FromTime(fromTime).ToDate(toDate).ToTime(toTime).UserId(userId).ProjectType(projectType).Page(page).PageSize(pageSize).Execute()

Time Totals on a Project



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
    fromDate := "fromDate_example" // string | string (YYYYMMDD) - The start date to retrieve from   (optional)
    fromTime := "fromTime_example" // string | string (HH:MM) - The start time only if fromdate is passed (optional)
    toDate := "toDate_example" // string | string (YYYYMMDD) - The end date to retrieve to (optional)
    toTime := "toTime_example" // string | string (HH:MM) - The end time only if todate is passed (optional)
    userId := int32(56) // int32 | Return time logs for a specific user only (optional)
    projectType := "projectType_example" // string | string (all, active, archived) - Filter the time entries to those in Active projects, Archived projects or All projects. (optional)
    page := int32(56) // int32 | The page number to show (optional)
    pageSize := int32(56) // int32 | The number of results per page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.GETProjectsIdTimeTotalJson(context.Background(), id).FromDate(fromDate).FromTime(fromTime).ToDate(toDate).ToTime(toTime).UserId(userId).ProjectType(projectType).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.GETProjectsIdTimeTotalJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsIdTimeTotalJson`: InlineResponse20078
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.GETProjectsIdTimeTotalJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsIdTimeTotalJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromDate** | **string** | string (YYYYMMDD) - The start date to retrieve from   | 
 **fromTime** | **string** | string (HH:MM) - The start time only if fromdate is passed | 
 **toDate** | **string** | string (YYYYMMDD) - The end date to retrieve to | 
 **toTime** | **string** | string (HH:MM) - The end time only if todate is passed | 
 **userId** | **int32** | Return time logs for a specific user only | 
 **projectType** | **string** | string (all, active, archived) - Filter the time entries to those in Active projects, Archived projects or All projects. | 
 **page** | **int32** | The page number to show | 
 **pageSize** | **int32** | The number of results per page | 

### Return type

[**InlineResponse20078**](InlineResponse20078.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsTimeTotalJson

> InlineResponse20058 GETProjectsTimeTotalJson(ctx).Execute()

Time Totals across Projects



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
    resp, r, err := api_client.TimeTrackingApi.GETProjectsTimeTotalJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.GETProjectsTimeTotalJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsTimeTotalJson`: InlineResponse20058
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.GETProjectsTimeTotalJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsTimeTotalJsonRequest struct via the builder pattern


### Return type

[**InlineResponse20058**](InlineResponse20058.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETTasklistsIdTimeTotalJson

> InlineResponse20098 GETTasklistsIdTimeTotalJson(ctx, id).FromDate(fromDate).FromTime(fromTime).ToDate(toDate).ToTime(toTime).UserId(userId).ProjectType(projectType).Execute()

Total Time on a Tasklist



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
    fromDate := "fromDate_example" // string | string (YYYYMMDD) - The start date to retrieve from   (optional)
    fromTime := "fromTime_example" // string | string (HH:MM) - The start time only if fromdate is passed (optional)
    toDate := "toDate_example" // string | string (YYYYMMDD) - The end date to retrieve to (optional)
    toTime := "toTime_example" // string | string (HH:MM) - The end time only if todate is passed (optional)
    userId := int32(56) // int32 | Return time logs for a specific user only (optional)
    projectType := "projectType_example" // string | string (all, active, archived) - Filter the time entries to those in Active projects, Archived projects or All projects. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.GETTasklistsIdTimeTotalJson(context.Background(), id).FromDate(fromDate).FromTime(fromTime).ToDate(toDate).ToTime(toTime).UserId(userId).ProjectType(projectType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.GETTasklistsIdTimeTotalJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETTasklistsIdTimeTotalJson`: InlineResponse20098
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.GETTasklistsIdTimeTotalJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETTasklistsIdTimeTotalJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromDate** | **string** | string (YYYYMMDD) - The start date to retrieve from   | 
 **fromTime** | **string** | string (HH:MM) - The start time only if fromdate is passed | 
 **toDate** | **string** | string (YYYYMMDD) - The end date to retrieve to | 
 **toTime** | **string** | string (HH:MM) - The end time only if todate is passed | 
 **userId** | **int32** | Return time logs for a specific user only | 
 **projectType** | **string** | string (all, active, archived) - Filter the time entries to those in Active projects, Archived projects or All projects. | 

### Return type

[**InlineResponse20098**](InlineResponse20098.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETTasksIdTimeEntriesJson

> InlineResponse200107 GETTasksIdTimeEntriesJson(ctx, id).Execute()

Retrieve all Task times



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
    resp, r, err := api_client.TimeTrackingApi.GETTasksIdTimeEntriesJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.GETTasksIdTimeEntriesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETTasksIdTimeEntriesJson`: InlineResponse200107
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.GETTasksIdTimeEntriesJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETTasksIdTimeEntriesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200107**](InlineResponse200107.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETTasksIdTimeTotalJson

> InlineResponse200106 GETTasksIdTimeTotalJson(ctx, id).FromDate(fromDate).FromTime(fromTime).ToDate(toDate).ToTime(toTime).UserId(userId).ProjectType(projectType).Execute()

Total Time on a Task



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
    fromDate := "fromDate_example" // string | string (YYYYMMDD) - The start date to retrieve from   (optional)
    fromTime := "fromTime_example" // string | string (HH:MM) - The start time only if fromdate is passed (optional)
    toDate := "toDate_example" // string | string (YYYYMMDD) - The end date to retrieve to (optional)
    toTime := "toTime_example" // string | string (HH:MM) - The end time only if todate is passed (optional)
    userId := int32(56) // int32 | Return time logs for a specific user only (optional)
    projectType := "projectType_example" // string | string (all, active, archived) - Filter the time entries to those in Active projects, Archived projects or All projects. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.GETTasksIdTimeTotalJson(context.Background(), id).FromDate(fromDate).FromTime(fromTime).ToDate(toDate).ToTime(toTime).UserId(userId).ProjectType(projectType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.GETTasksIdTimeTotalJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETTasksIdTimeTotalJson`: InlineResponse200106
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.GETTasksIdTimeTotalJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETTasksIdTimeTotalJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromDate** | **string** | string (YYYYMMDD) - The start date to retrieve from   | 
 **fromTime** | **string** | string (HH:MM) - The start time only if fromdate is passed | 
 **toDate** | **string** | string (YYYYMMDD) - The end date to retrieve to | 
 **toTime** | **string** | string (HH:MM) - The end time only if todate is passed | 
 **userId** | **int32** | Return time logs for a specific user only | 
 **projectType** | **string** | string (all, active, archived) - Filter the time entries to those in Active projects, Archived projects or All projects. | 

### Return type

[**InlineResponse200106**](InlineResponse200106.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETTimeEntriesIdJson

> InlineResponse200112 GETTimeEntriesIdJson(ctx, id).Execute()

Retrieve single Time-Entry



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
    resp, r, err := api_client.TimeTrackingApi.GETTimeEntriesIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.GETTimeEntriesIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETTimeEntriesIdJson`: InlineResponse200112
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.GETTimeEntriesIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETTimeEntriesIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200112**](InlineResponse200112.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETTimeEntriesJson

> InlineResponse200107 GETTimeEntriesJson(ctx).Page(page).Fromdate(fromdate).Fromtime(fromtime).Todate(todate).Totime(totime).Sortby(sortby).Sortorder(sortorder).UserId(userId).BillableType(billableType).InvoicedType(invoicedType).ProjectType(projectType).ShowDeleted(showDeleted).TagIds(tagIds).UpdatedAfterDate(updatedAfterDate).ProjectId(projectId).PageSize(pageSize).TaskTagIds(taskTagIds).Execute()

Retrieve all Time Entries across all Projects



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
    page := int32(56) // int32 | The page to start retrieving entries from ( e.g: page=1 gives records 1 - 500, page=2 gives records 501-1001 etc) (optional)
    fromdate := "fromdate_example" // string | string (YYYYMMDD) - The start date to retrieve from (optional)
    fromtime := "fromtime_example" // string | string (HH:MM) - The start time only if fromdate is passed (optional)
    todate := "todate_example" // string | string (YYYYMMDD) - The end date to retrieve to (optional)
    totime := "totime_example" // string | string (HH:MM) - The end time only if todate is passed (optional)
    sortby := "sortby_example" // string | string - One of the following: date, user, task, tasklist, project, company, dateupdated (Default: date) (optional)
    sortorder := "sortorder_example" // string | ASC or DESC - The order to sort the returned data (optional)
    userId := int32(56) // int32 | Return time logs for a specific user only (optional)
    billableType := "billableType_example" // string | string (billable or nonbillable).Filter the Time Entries to those that are Billable or Not Billable. (optional)
    invoicedType := "invoicedType_example" // string | string (invoiced or noninvoiced) - filter the time entries to those that have been Invoiced or not. (optional)
    projectType := "projectType_example" // string | string (all, active, archived) - Filter the time entries to those in Active projects, Archived projects or All projects. (optional)
    showDeleted := "showDeleted_example" // string | string (true or false) - Filter time entries to include deleted time sheet entries or not (optional)
    tagIds := "tagIds_example" // string | numeric - Include one or more Tag IDs here to return only the time entries with those tags attached (eg &tagIds=23,445,454) (optional)
    updatedAfterDate := "updatedAfterDate_example" // string | string YYYYMMDDHHMMSS (optional)
    projectId := "projectId_example" // string |  (optional)
    pageSize := "pageSize_example" // string | eg. pageSize=500 (optional)
    taskTagIds := "taskTagIds_example" // string | filter on task tags (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.GETTimeEntriesJson(context.Background()).Page(page).Fromdate(fromdate).Fromtime(fromtime).Todate(todate).Totime(totime).Sortby(sortby).Sortorder(sortorder).UserId(userId).BillableType(billableType).InvoicedType(invoicedType).ProjectType(projectType).ShowDeleted(showDeleted).TagIds(tagIds).UpdatedAfterDate(updatedAfterDate).ProjectId(projectId).PageSize(pageSize).TaskTagIds(taskTagIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.GETTimeEntriesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETTimeEntriesJson`: InlineResponse200107
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.GETTimeEntriesJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETTimeEntriesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page to start retrieving entries from ( e.g: page&#x3D;1 gives records 1 - 500, page&#x3D;2 gives records 501-1001 etc) | 
 **fromdate** | **string** | string (YYYYMMDD) - The start date to retrieve from | 
 **fromtime** | **string** | string (HH:MM) - The start time only if fromdate is passed | 
 **todate** | **string** | string (YYYYMMDD) - The end date to retrieve to | 
 **totime** | **string** | string (HH:MM) - The end time only if todate is passed | 
 **sortby** | **string** | string - One of the following: date, user, task, tasklist, project, company, dateupdated (Default: date) | 
 **sortorder** | **string** | ASC or DESC - The order to sort the returned data | 
 **userId** | **int32** | Return time logs for a specific user only | 
 **billableType** | **string** | string (billable or nonbillable).Filter the Time Entries to those that are Billable or Not Billable. | 
 **invoicedType** | **string** | string (invoiced or noninvoiced) - filter the time entries to those that have been Invoiced or not. | 
 **projectType** | **string** | string (all, active, archived) - Filter the time entries to those in Active projects, Archived projects or All projects. | 
 **showDeleted** | **string** | string (true or false) - Filter time entries to include deleted time sheet entries or not | 
 **tagIds** | **string** | numeric - Include one or more Tag IDs here to return only the time entries with those tags attached (eg &amp;tagIds&#x3D;23,445,454) | 
 **updatedAfterDate** | **string** | string YYYYMMDDHHMMSS | 
 **projectId** | **string** |  | 
 **pageSize** | **string** | eg. pageSize&#x3D;500 | 
 **taskTagIds** | **string** | filter on task tags | 

### Return type

[**InlineResponse200107**](InlineResponse200107.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETTimeTotalJson

> InlineResponse200111 GETTimeTotalJson(ctx).FromDate(fromDate).FromTime(fromTime).ToDate(toDate).ToTime(toTime).UserId(userId).ProjectType(projectType).Execute()

Get Total Time across an Account



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
    fromDate := "fromDate_example" // string | string (YYYYMMDD) - The start date to retrieve from   (optional)
    fromTime := "fromTime_example" // string | string (HH:MM) - The start time only if fromdate is passed (optional)
    toDate := "toDate_example" // string | string (YYYYMMDD) - The end date to retrieve to (optional)
    toTime := "toTime_example" // string | string (HH:MM) - The end time only if todate is passed (optional)
    userId := int32(56) // int32 | Return time logs for a specific user only (optional)
    projectType := "projectType_example" // string | string (all, active, archived) - Filter the time entries to those in Active projects, Archived projects or All projects. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.GETTimeTotalJson(context.Background()).FromDate(fromDate).FromTime(fromTime).ToDate(toDate).ToTime(toTime).UserId(userId).ProjectType(projectType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.GETTimeTotalJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETTimeTotalJson`: InlineResponse200111
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.GETTimeTotalJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETTimeTotalJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromDate** | **string** | string (YYYYMMDD) - The start date to retrieve from   | 
 **fromTime** | **string** | string (HH:MM) - The start time only if fromdate is passed | 
 **toDate** | **string** | string (YYYYMMDD) - The end date to retrieve to | 
 **toTime** | **string** | string (HH:MM) - The end time only if todate is passed | 
 **userId** | **int32** | Return time logs for a specific user only | 
 **projectType** | **string** | string (all, active, archived) - Filter the time entries to those in Active projects, Archived projects or All projects. | 

### Return type

[**InlineResponse200111**](InlineResponse200111.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETTimersJson

> InlineResponse20030 GETTimersJson(ctx).ProjectId(projectId).TaskId(taskId).RunningTimersOnly(runningTimersOnly).Execute()

Get all Running Timers



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
    projectId := "projectId_example" // string | timers on a given project (optional)
    taskId := "taskId_example" // string | Timers on a task (optional)
    runningTimersOnly := true // bool | all active timers or only timers which are currently running (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.GETTimersJson(context.Background()).ProjectId(projectId).TaskId(taskId).RunningTimersOnly(runningTimersOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.GETTimersJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETTimersJson`: InlineResponse20030
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.GETTimersJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETTimersJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string** | timers on a given project | 
 **taskId** | **string** | Timers on a task | 
 **runningTimersOnly** | **bool** | all active timers or only timers which are currently running | 

### Return type

[**InlineResponse20030**](InlineResponse20030.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsIdTimeEntriesJson

> InlineResponse2015 POSTProjectsIdTimeEntriesJson(ctx, id).Body(body).Execute()

Create a Time-Entry



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
    body := *openapiclient.NewInlineObject75() // InlineObject75 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.POSTProjectsIdTimeEntriesJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.POSTProjectsIdTimeEntriesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsIdTimeEntriesJson`: InlineResponse2015
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.POSTProjectsIdTimeEntriesJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsIdTimeEntriesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject75**](InlineObject75.md) |  | 

### Return type

[**InlineResponse2015**](InlineResponse2015.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTTasksIdTimeEntriesJson

> InlineResponse20018 POSTTasksIdTimeEntriesJson(ctx, id).Body(body).Execute()

Create a Time-Entry (for a Task)



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
    body := *openapiclient.NewInlineObject101() // InlineObject101 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.POSTTasksIdTimeEntriesJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.POSTTasksIdTimeEntriesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTTasksIdTimeEntriesJson`: InlineResponse20018
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.POSTTasksIdTimeEntriesJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTTasksIdTimeEntriesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject101**](InlineObject101.md) |  | 

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


## PUTTasksIdEstimatedtimeJson

> InlineResponse2001 PUTTasksIdEstimatedtimeJson(ctx, id).Body(body).Execute()

Add a Time estimate to a Task



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
    body := *openapiclient.NewInlineObject97() // InlineObject97 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.PUTTasksIdEstimatedtimeJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.PUTTasksIdEstimatedtimeJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTTasksIdEstimatedtimeJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.PUTTasksIdEstimatedtimeJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTTasksIdEstimatedtimeJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject97**](InlineObject97.md) |  | 

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


## PUTTimeEntriesIdJson

> InlineResponse2001 PUTTimeEntriesIdJson(ctx, id).Body(body).Execute()

Update a Time Entry



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
    body := *openapiclient.NewInlineObject105() // InlineObject105 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.PUTTimeEntriesIdJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.PUTTimeEntriesIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTTimeEntriesIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.PUTTimeEntriesIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTTimeEntriesIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject105**](InlineObject105.md) |  | 

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

