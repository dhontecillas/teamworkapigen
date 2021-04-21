# \TimeTrackingApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEProjectsApiV3MeTimerstimerIdJson**](TimeTrackingApi.md#DELETEProjectsApiV3MeTimerstimerIdJson) | **Delete** /projects/api/v3/me/timers/{timerId}.json | Delete a timer by ID
[**GETProjectsApiV3AllocationsallocationIdTimeJson**](TimeTrackingApi.md#GETProjectsApiV3AllocationsallocationIdTimeJson) | **Get** /projects/api/v3/allocations/{allocationId}/time.json | Get time entries for a specific allocation
[**GETProjectsApiV3MeTimersJson**](TimeTrackingApi.md#GETProjectsApiV3MeTimersJson) | **Get** /projects/api/v3/me/timers.json | Get all your running timers
[**GETProjectsApiV3ProjectsprojectIdTimeJson**](TimeTrackingApi.md#GETProjectsApiV3ProjectsprojectIdTimeJson) | **Get** /projects/api/v3/projects/{projectId}/time.json | Get time entries for a specific project
[**GETProjectsApiV3ProjectsprojectIdTimeTotalJson**](TimeTrackingApi.md#GETProjectsApiV3ProjectsprojectIdTimeTotalJson) | **Get** /projects/api/v3/projects/{projectId}/time/total.json | Get timelog totals.
[**GETProjectsApiV3TimeJson**](TimeTrackingApi.md#GETProjectsApiV3TimeJson) | **Get** /projects/api/v3/time.json | Get all time entries
[**GETProjectsApiV3TimeTotalJson**](TimeTrackingApi.md#GETProjectsApiV3TimeTotalJson) | **Get** /projects/api/v3/time/total.json | Get timelog totals.
[**GETProjectsApiV3TimersJson**](TimeTrackingApi.md#GETProjectsApiV3TimersJson) | **Get** /projects/api/v3/timers.json | Get all running timers
[**GETProjectsApiV3TimerstimerIdJson**](TimeTrackingApi.md#GETProjectsApiV3TimerstimerIdJson) | **Get** /projects/api/v3/timers/{timerId}.json | Get a specific timer
[**POSTProjectsApiV3MeTimersJson**](TimeTrackingApi.md#POSTProjectsApiV3MeTimersJson) | **Post** /projects/api/v3/me/timers.json | Create a new timer
[**PUTProjectsApiV3MeTimerstimerIdCompleteJson**](TimeTrackingApi.md#PUTProjectsApiV3MeTimerstimerIdCompleteJson) | **Put** /projects/api/v3/me/timers/{timerId}/complete.json | Complete a timer by ID
[**PUTProjectsApiV3MeTimerstimerIdJson**](TimeTrackingApi.md#PUTProjectsApiV3MeTimerstimerIdJson) | **Put** /projects/api/v3/me/timers/{timerId}.json | Edits a timer
[**PUTProjectsApiV3MeTimerstimerIdPauseJson**](TimeTrackingApi.md#PUTProjectsApiV3MeTimerstimerIdPauseJson) | **Put** /projects/api/v3/me/timers/{timerId}/pause.json | Pause a timer by ID
[**PUTProjectsApiV3MeTimerstimerIdResumeJson**](TimeTrackingApi.md#PUTProjectsApiV3MeTimerstimerIdResumeJson) | **Put** /projects/api/v3/me/timers/{timerId}/resume.json | Resume a timer by ID
[**PUTProjectsApiV3MeTimerstimerIdUndeleteJson**](TimeTrackingApi.md#PUTProjectsApiV3MeTimerstimerIdUndeleteJson) | **Put** /projects/api/v3/me/timers/{timerId}/undelete.json | Undelete a timer by ID



## DELETEProjectsApiV3MeTimerstimerIdJson

> DELETEProjectsApiV3MeTimerstimerIdJson(ctx, timerId).TimerDeleteRequest(timerDeleteRequest).Execute()

Delete a timer by ID



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
    timerId := int32(56) // int32 | 
    timerDeleteRequest := *openapiclient.NewTimerDeleteRequest() // TimerDeleteRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.DELETEProjectsApiV3MeTimerstimerIdJson(context.Background(), timerId).TimerDeleteRequest(timerDeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.DELETEProjectsApiV3MeTimerstimerIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEProjectsApiV3MeTimerstimerIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timerDeleteRequest** | [**TimerDeleteRequest**](TimerDeleteRequest.md) |  | 

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


## GETProjectsApiV3AllocationsallocationIdTimeJson

> TimelogTimelogsResponse GETProjectsApiV3AllocationsallocationIdTimeJson(ctx, allocationId).UpdatedAfter(updatedAfter).StartDate(startDate).SelectedColumns(selectedColumns).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).InvoicedType(invoicedType).EndDate(endDate).BillableType(billableType).TicketId(ticketId).TasklistId(tasklistId).ProjectId(projectId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).AllocationId2(allocationId2).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).IsReportDownload(isReportDownload).IncludeArchivedProjects(includeArchivedProjects).TaskTagIds(taskTagIds).TagIds(tagIds).ProjectsFromCompanyId(projectsFromCompanyId).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsTimelogs(fieldsTimelogs).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).CompanyIds(companyIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()

Get time entries for a specific allocation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    allocationId := int32(56) // int32 | 
    updatedAfter := time.Now() // time.Time | filter by updated after date (optional)
    startDate := time.Now() // string | filter by a starting date (optional)
    selectedColumns := "selectedColumns_example" // string | customise the report by selecting columns (optional)
    reportFormat := "reportFormat_example" // string | define the format of the report (optional)
    projectStatuses := "projectStatuses_example" // string | filter by project statuses (optional)
    projectStatus := "projectStatus_example" // string | filter by project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | sort order (optional) (default to "date")
    invoicedType := "invoicedType_example" // string | filter by invoiced type (optional) (default to "all")
    endDate := time.Now() // string | filter by an ending date (optional)
    billableType := "billableType_example" // string | filter by billable type (optional) (default to "all")
    ticketId := int32(56) // int32 | filter by ticket id (optional)
    tasklistId := int32(56) // int32 | filter by tasklist id (optional)
    projectId := int32(56) // int32 | filter by project id (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    allocationId2 := int32(56) // int32 | filter by allocation id (optional)
    showDeleted := true // bool | include deleted items (optional) (default to false)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    matchAllTags := true // bool | match all tags (optional)
    isReportDownload := true // bool | generate a report document (optional)
    includeArchivedProjects := true // bool | include archived projects (optional) (default to false)
    taskTagIds := []int32{int32(123)} // []int32 | filter by task tag ids (optional)
    tagIds := []int32{int32(123)} // []int32 | filter by tag ids (optional)
    projectsFromCompanyId := []int32{int32(123)} // []int32 | filter by project company ids (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTimelogs := []string{"Inner_example"} // []string |  (optional)
    fieldsTasks := []string{"Inner_example"} // []string |  (optional)
    fieldsTasklists := []string{"Inner_example"} // []string |  (optional)
    fieldsTags := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    companyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    assignedToUserIds := []int32{int32(123)} // []int32 | filter by assigned user ids (optional)
    assignedToTeamIds := []int32{int32(123)} // []int32 | filter by assigned team ids (optional)
    assignedToCompanyIds := []int32{int32(123)} // []int32 | filter by assigned company ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.GETProjectsApiV3AllocationsallocationIdTimeJson(context.Background(), allocationId).UpdatedAfter(updatedAfter).StartDate(startDate).SelectedColumns(selectedColumns).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).InvoicedType(invoicedType).EndDate(endDate).BillableType(billableType).TicketId(ticketId).TasklistId(tasklistId).ProjectId(projectId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).AllocationId2(allocationId2).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).IsReportDownload(isReportDownload).IncludeArchivedProjects(includeArchivedProjects).TaskTagIds(taskTagIds).TagIds(tagIds).ProjectsFromCompanyId(projectsFromCompanyId).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsTimelogs(fieldsTimelogs).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).CompanyIds(companyIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.GETProjectsApiV3AllocationsallocationIdTimeJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3AllocationsallocationIdTimeJson`: TimelogTimelogsResponse
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.GETProjectsApiV3AllocationsallocationIdTimeJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allocationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3AllocationsallocationIdTimeJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedAfter** | **time.Time** | filter by updated after date | 
 **startDate** | **string** | filter by a starting date | 
 **selectedColumns** | **string** | customise the report by selecting columns | 
 **reportFormat** | **string** | define the format of the report | 
 **projectStatuses** | **string** | filter by project statuses | 
 **projectStatus** | **string** | filter by project status | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | sort order | [default to &quot;date&quot;]
 **invoicedType** | **string** | filter by invoiced type | [default to &quot;all&quot;]
 **endDate** | **string** | filter by an ending date | 
 **billableType** | **string** | filter by billable type | [default to &quot;all&quot;]
 **ticketId** | **int32** | filter by ticket id | 
 **tasklistId** | **int32** | filter by tasklist id | 
 **projectId** | **int32** | filter by project id | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **allocationId2** | **int32** | filter by allocation id | 
 **showDeleted** | **bool** | include deleted items | [default to false]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **matchAllTags** | **bool** | match all tags | 
 **isReportDownload** | **bool** | generate a report document | 
 **includeArchivedProjects** | **bool** | include archived projects | [default to false]
 **taskTagIds** | **[]int32** | filter by task tag ids | 
 **tagIds** | **[]int32** | filter by tag ids | 
 **projectsFromCompanyId** | **[]int32** | filter by project company ids | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTimelogs** | **[]string** |  | 
 **fieldsTasks** | **[]string** |  | 
 **fieldsTasklists** | **[]string** |  | 
 **fieldsTags** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **companyIds** | **[]int32** | filter by company ids | 
 **assignedToUserIds** | **[]int32** | filter by assigned user ids | 
 **assignedToTeamIds** | **[]int32** | filter by assigned team ids | 
 **assignedToCompanyIds** | **[]int32** | filter by assigned company ids | 

### Return type

[**TimelogTimelogsResponse**](TimelogTimelogsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3MeTimersJson

> TimerTimersResponse GETProjectsApiV3MeTimersJson(ctx).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).UserId(userId).TaskId(taskId).ProjectId(projectId).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).RunningTimersOnly(runningTimersOnly).Include(include).FieldsUsers(fieldsUsers).FieldsTimers(fieldsTimers).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsProjects(fieldsProjects).Execute()

Get all your running timers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    updatedAfterDate := time.Now() // time.Time | filter by updated after date (deprecated, use updatedAfter) (optional)
    updatedAfter := time.Now() // time.Time | filter by updated after date (optional)
    userId := int32(56) // int32 | filter by user id (optional)
    taskId := int32(56) // int32 | filter by task id (optional)
    projectId := int32(56) // int32 | filter by project id (optional)
    pageSize := int32(56) // int32 | number of items in a page default: 50 (optional)
    page := int32(56) // int32 | page number default: 1 (optional)
    showDeleted := true // bool | include deleted items default: false (optional)
    runningTimersOnly := true // bool | enforce running timers only default: false (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTimers := []string{"Inner_example"} // []string |  (optional)
    fieldsTasks := []string{"Inner_example"} // []string |  (optional)
    fieldsTasklists := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.GETProjectsApiV3MeTimersJson(context.Background()).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).UserId(userId).TaskId(taskId).ProjectId(projectId).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).RunningTimersOnly(runningTimersOnly).Include(include).FieldsUsers(fieldsUsers).FieldsTimers(fieldsTimers).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsProjects(fieldsProjects).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.GETProjectsApiV3MeTimersJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3MeTimersJson`: TimerTimersResponse
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.GETProjectsApiV3MeTimersJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3MeTimersJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfterDate** | **time.Time** | filter by updated after date (deprecated, use updatedAfter) | 
 **updatedAfter** | **time.Time** | filter by updated after date | 
 **userId** | **int32** | filter by user id | 
 **taskId** | **int32** | filter by task id | 
 **projectId** | **int32** | filter by project id | 
 **pageSize** | **int32** | number of items in a page default: 50 | 
 **page** | **int32** | page number default: 1 | 
 **showDeleted** | **bool** | include deleted items default: false | 
 **runningTimersOnly** | **bool** | enforce running timers only default: false | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTimers** | **[]string** |  | 
 **fieldsTasks** | **[]string** |  | 
 **fieldsTasklists** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 

### Return type

[**TimerTimersResponse**](TimerTimersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3ProjectsprojectIdTimeJson

> TimelogTimelogsResponse GETProjectsApiV3ProjectsprojectIdTimeJson(ctx, projectId).UpdatedAfter(updatedAfter).StartDate(startDate).SelectedColumns(selectedColumns).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).InvoicedType(invoicedType).EndDate(endDate).BillableType(billableType).TicketId(ticketId).TasklistId(tasklistId).ProjectId2(projectId2).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).AllocationId(allocationId).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).IsReportDownload(isReportDownload).IncludeArchivedProjects(includeArchivedProjects).TaskTagIds(taskTagIds).TagIds(tagIds).ProjectsFromCompanyId(projectsFromCompanyId).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsTimelogs(fieldsTimelogs).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).CompanyIds(companyIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()

Get time entries for a specific project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    projectId := int32(56) // int32 | 
    updatedAfter := time.Now() // time.Time | filter by updated after date (optional)
    startDate := time.Now() // string | filter by a starting date (optional)
    selectedColumns := "selectedColumns_example" // string | customise the report by selecting columns (optional)
    reportFormat := "reportFormat_example" // string | define the format of the report (optional)
    projectStatuses := "projectStatuses_example" // string | filter by project statuses (optional)
    projectStatus := "projectStatus_example" // string | filter by project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | sort order (optional) (default to "date")
    invoicedType := "invoicedType_example" // string | filter by invoiced type (optional) (default to "all")
    endDate := time.Now() // string | filter by an ending date (optional)
    billableType := "billableType_example" // string | filter by billable type (optional) (default to "all")
    ticketId := int32(56) // int32 | filter by ticket id (optional)
    tasklistId := int32(56) // int32 | filter by tasklist id (optional)
    projectId2 := int32(56) // int32 | filter by project id (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    allocationId := int32(56) // int32 | filter by allocation id (optional)
    showDeleted := true // bool | include deleted items (optional) (default to false)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    matchAllTags := true // bool | match all tags (optional)
    isReportDownload := true // bool | generate a report document (optional)
    includeArchivedProjects := true // bool | include archived projects (optional) (default to false)
    taskTagIds := []int32{int32(123)} // []int32 | filter by task tag ids (optional)
    tagIds := []int32{int32(123)} // []int32 | filter by tag ids (optional)
    projectsFromCompanyId := []int32{int32(123)} // []int32 | filter by project company ids (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTimelogs := []string{"Inner_example"} // []string |  (optional)
    fieldsTasks := []string{"Inner_example"} // []string |  (optional)
    fieldsTasklists := []string{"Inner_example"} // []string |  (optional)
    fieldsTags := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    companyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    assignedToUserIds := []int32{int32(123)} // []int32 | filter by assigned user ids (optional)
    assignedToTeamIds := []int32{int32(123)} // []int32 | filter by assigned team ids (optional)
    assignedToCompanyIds := []int32{int32(123)} // []int32 | filter by assigned company ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.GETProjectsApiV3ProjectsprojectIdTimeJson(context.Background(), projectId).UpdatedAfter(updatedAfter).StartDate(startDate).SelectedColumns(selectedColumns).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).InvoicedType(invoicedType).EndDate(endDate).BillableType(billableType).TicketId(ticketId).TasklistId(tasklistId).ProjectId2(projectId2).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).AllocationId(allocationId).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).IsReportDownload(isReportDownload).IncludeArchivedProjects(includeArchivedProjects).TaskTagIds(taskTagIds).TagIds(tagIds).ProjectsFromCompanyId(projectsFromCompanyId).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsTimelogs(fieldsTimelogs).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).CompanyIds(companyIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.GETProjectsApiV3ProjectsprojectIdTimeJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ProjectsprojectIdTimeJson`: TimelogTimelogsResponse
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.GETProjectsApiV3ProjectsprojectIdTimeJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsprojectIdTimeJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedAfter** | **time.Time** | filter by updated after date | 
 **startDate** | **string** | filter by a starting date | 
 **selectedColumns** | **string** | customise the report by selecting columns | 
 **reportFormat** | **string** | define the format of the report | 
 **projectStatuses** | **string** | filter by project statuses | 
 **projectStatus** | **string** | filter by project status | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | sort order | [default to &quot;date&quot;]
 **invoicedType** | **string** | filter by invoiced type | [default to &quot;all&quot;]
 **endDate** | **string** | filter by an ending date | 
 **billableType** | **string** | filter by billable type | [default to &quot;all&quot;]
 **ticketId** | **int32** | filter by ticket id | 
 **tasklistId** | **int32** | filter by tasklist id | 
 **projectId2** | **int32** | filter by project id | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **allocationId** | **int32** | filter by allocation id | 
 **showDeleted** | **bool** | include deleted items | [default to false]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **matchAllTags** | **bool** | match all tags | 
 **isReportDownload** | **bool** | generate a report document | 
 **includeArchivedProjects** | **bool** | include archived projects | [default to false]
 **taskTagIds** | **[]int32** | filter by task tag ids | 
 **tagIds** | **[]int32** | filter by tag ids | 
 **projectsFromCompanyId** | **[]int32** | filter by project company ids | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTimelogs** | **[]string** |  | 
 **fieldsTasks** | **[]string** |  | 
 **fieldsTasklists** | **[]string** |  | 
 **fieldsTags** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **companyIds** | **[]int32** | filter by company ids | 
 **assignedToUserIds** | **[]int32** | filter by assigned user ids | 
 **assignedToTeamIds** | **[]int32** | filter by assigned team ids | 
 **assignedToCompanyIds** | **[]int32** | filter by assigned company ids | 

### Return type

[**TimelogTimelogsResponse**](TimelogTimelogsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3ProjectsprojectIdTimeTotalJson

> TimelogTotalsResponse GETProjectsApiV3ProjectsprojectIdTimeTotalJson(ctx, projectId2).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).UserId(userId).TasklistId(tasklistId).TaskId(taskId).ProjectId(projectId).ProjectHealths(projectHealths).DeskTicketId(deskTicketId).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).UserIds(userIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).CompanyIds(companyIds).Execute()

Get timelog totals.

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
    projectId2 := int32(56) // int32 | 
    projectStatuses := "projectStatuses_example" // string | filter by project statuses (optional)
    projectStatus := "projectStatus_example" // string | filter by project status (optional)
    userId := int32(56) // int32 | filter by user id (optional)
    tasklistId := int32(56) // int32 | filter by tasklist id (optional)
    taskId := int32(56) // int32 | filter by task id (optional)
    projectId := int32(56) // int32 | filter by project id (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    deskTicketId := int32(56) // int32 | filter by desk ticket id (optional)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    matchAllTags := true // bool | match all tags (optional)
    userIds := []int32{int32(123)} // []int32 | filter by user ids (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by project company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    companyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.GETProjectsApiV3ProjectsprojectIdTimeTotalJson(context.Background(), projectId2).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).UserId(userId).TasklistId(tasklistId).TaskId(taskId).ProjectId(projectId).ProjectHealths(projectHealths).DeskTicketId(deskTicketId).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).UserIds(userIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).CompanyIds(companyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.GETProjectsApiV3ProjectsprojectIdTimeTotalJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ProjectsprojectIdTimeTotalJson`: TimelogTotalsResponse
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.GETProjectsApiV3ProjectsprojectIdTimeTotalJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsprojectIdTimeTotalJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectStatuses** | **string** | filter by project statuses | 
 **projectStatus** | **string** | filter by project status | 
 **userId** | **int32** | filter by user id | 
 **tasklistId** | **int32** | filter by tasklist id | 
 **taskId** | **int32** | filter by task id | 
 **projectId** | **int32** | filter by project id | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **deskTicketId** | **int32** | filter by desk ticket id | 
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **matchAllTags** | **bool** | match all tags | 
 **userIds** | **[]int32** | filter by user ids | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by project company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **companyIds** | **[]int32** | filter by company ids | 

### Return type

[**TimelogTotalsResponse**](TimelogTotalsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3TimeJson

> TimelogTimelogsResponse GETProjectsApiV3TimeJson(ctx).UpdatedAfter(updatedAfter).StartDate(startDate).SelectedColumns(selectedColumns).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).InvoicedType(invoicedType).EndDate(endDate).BillableType(billableType).TicketId(ticketId).TasklistId(tasklistId).ProjectId(projectId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).AllocationId(allocationId).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).IsReportDownload(isReportDownload).IncludeArchivedProjects(includeArchivedProjects).TaskTagIds(taskTagIds).TagIds(tagIds).ProjectsFromCompanyId(projectsFromCompanyId).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsTimelogs(fieldsTimelogs).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).CompanyIds(companyIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()

Get all time entries



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    updatedAfter := time.Now() // time.Time | filter by updated after date (optional)
    startDate := time.Now() // string | filter by a starting date (optional)
    selectedColumns := "selectedColumns_example" // string | customise the report by selecting columns (optional)
    reportFormat := "reportFormat_example" // string | define the format of the report (optional)
    projectStatuses := "projectStatuses_example" // string | filter by project statuses (optional)
    projectStatus := "projectStatus_example" // string | filter by project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | sort order (optional) (default to "date")
    invoicedType := "invoicedType_example" // string | filter by invoiced type (optional) (default to "all")
    endDate := time.Now() // string | filter by an ending date (optional)
    billableType := "billableType_example" // string | filter by billable type (optional) (default to "all")
    ticketId := int32(56) // int32 | filter by ticket id (optional)
    tasklistId := int32(56) // int32 | filter by tasklist id (optional)
    projectId := int32(56) // int32 | filter by project id (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    allocationId := int32(56) // int32 | filter by allocation id (optional)
    showDeleted := true // bool | include deleted items (optional) (default to false)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    matchAllTags := true // bool | match all tags (optional)
    isReportDownload := true // bool | generate a report document (optional)
    includeArchivedProjects := true // bool | include archived projects (optional) (default to false)
    taskTagIds := []int32{int32(123)} // []int32 | filter by task tag ids (optional)
    tagIds := []int32{int32(123)} // []int32 | filter by tag ids (optional)
    projectsFromCompanyId := []int32{int32(123)} // []int32 | filter by project company ids (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTimelogs := []string{"Inner_example"} // []string |  (optional)
    fieldsTasks := []string{"Inner_example"} // []string |  (optional)
    fieldsTasklists := []string{"Inner_example"} // []string |  (optional)
    fieldsTags := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    companyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    assignedToUserIds := []int32{int32(123)} // []int32 | filter by assigned user ids (optional)
    assignedToTeamIds := []int32{int32(123)} // []int32 | filter by assigned team ids (optional)
    assignedToCompanyIds := []int32{int32(123)} // []int32 | filter by assigned company ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.GETProjectsApiV3TimeJson(context.Background()).UpdatedAfter(updatedAfter).StartDate(startDate).SelectedColumns(selectedColumns).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).InvoicedType(invoicedType).EndDate(endDate).BillableType(billableType).TicketId(ticketId).TasklistId(tasklistId).ProjectId(projectId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).AllocationId(allocationId).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).IsReportDownload(isReportDownload).IncludeArchivedProjects(includeArchivedProjects).TaskTagIds(taskTagIds).TagIds(tagIds).ProjectsFromCompanyId(projectsFromCompanyId).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsTimelogs(fieldsTimelogs).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).CompanyIds(companyIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.GETProjectsApiV3TimeJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3TimeJson`: TimelogTimelogsResponse
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.GETProjectsApiV3TimeJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3TimeJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfter** | **time.Time** | filter by updated after date | 
 **startDate** | **string** | filter by a starting date | 
 **selectedColumns** | **string** | customise the report by selecting columns | 
 **reportFormat** | **string** | define the format of the report | 
 **projectStatuses** | **string** | filter by project statuses | 
 **projectStatus** | **string** | filter by project status | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | sort order | [default to &quot;date&quot;]
 **invoicedType** | **string** | filter by invoiced type | [default to &quot;all&quot;]
 **endDate** | **string** | filter by an ending date | 
 **billableType** | **string** | filter by billable type | [default to &quot;all&quot;]
 **ticketId** | **int32** | filter by ticket id | 
 **tasklistId** | **int32** | filter by tasklist id | 
 **projectId** | **int32** | filter by project id | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **allocationId** | **int32** | filter by allocation id | 
 **showDeleted** | **bool** | include deleted items | [default to false]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **matchAllTags** | **bool** | match all tags | 
 **isReportDownload** | **bool** | generate a report document | 
 **includeArchivedProjects** | **bool** | include archived projects | [default to false]
 **taskTagIds** | **[]int32** | filter by task tag ids | 
 **tagIds** | **[]int32** | filter by tag ids | 
 **projectsFromCompanyId** | **[]int32** | filter by project company ids | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTimelogs** | **[]string** |  | 
 **fieldsTasks** | **[]string** |  | 
 **fieldsTasklists** | **[]string** |  | 
 **fieldsTags** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **companyIds** | **[]int32** | filter by company ids | 
 **assignedToUserIds** | **[]int32** | filter by assigned user ids | 
 **assignedToTeamIds** | **[]int32** | filter by assigned team ids | 
 **assignedToCompanyIds** | **[]int32** | filter by assigned company ids | 

### Return type

[**TimelogTimelogsResponse**](TimelogTimelogsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3TimeTotalJson

> TimelogTotalsResponse GETProjectsApiV3TimeTotalJson(ctx).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).UserId(userId).TasklistId(tasklistId).TaskId(taskId).ProjectId(projectId).ProjectHealths(projectHealths).DeskTicketId(deskTicketId).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).UserIds(userIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).CompanyIds(companyIds).Execute()

Get timelog totals.

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
    projectStatuses := "projectStatuses_example" // string | filter by project statuses (optional)
    projectStatus := "projectStatus_example" // string | filter by project status (optional)
    userId := int32(56) // int32 | filter by user id (optional)
    tasklistId := int32(56) // int32 | filter by tasklist id (optional)
    taskId := int32(56) // int32 | filter by task id (optional)
    projectId := int32(56) // int32 | filter by project id (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    deskTicketId := int32(56) // int32 | filter by desk ticket id (optional)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    matchAllTags := true // bool | match all tags (optional)
    userIds := []int32{int32(123)} // []int32 | filter by user ids (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by project company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    companyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.GETProjectsApiV3TimeTotalJson(context.Background()).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).UserId(userId).TasklistId(tasklistId).TaskId(taskId).ProjectId(projectId).ProjectHealths(projectHealths).DeskTicketId(deskTicketId).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).UserIds(userIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).CompanyIds(companyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.GETProjectsApiV3TimeTotalJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3TimeTotalJson`: TimelogTotalsResponse
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.GETProjectsApiV3TimeTotalJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3TimeTotalJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectStatuses** | **string** | filter by project statuses | 
 **projectStatus** | **string** | filter by project status | 
 **userId** | **int32** | filter by user id | 
 **tasklistId** | **int32** | filter by tasklist id | 
 **taskId** | **int32** | filter by task id | 
 **projectId** | **int32** | filter by project id | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **deskTicketId** | **int32** | filter by desk ticket id | 
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **matchAllTags** | **bool** | match all tags | 
 **userIds** | **[]int32** | filter by user ids | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by project company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **companyIds** | **[]int32** | filter by company ids | 

### Return type

[**TimelogTotalsResponse**](TimelogTotalsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3TimersJson

> TimerTimersResponse GETProjectsApiV3TimersJson(ctx).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).UserId(userId).TaskId(taskId).ProjectId(projectId).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).RunningTimersOnly(runningTimersOnly).Include(include).FieldsUsers(fieldsUsers).FieldsTimers(fieldsTimers).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsProjects(fieldsProjects).Execute()

Get all running timers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    updatedAfterDate := time.Now() // time.Time | filter by updated after date (deprecated, use updatedAfter) (optional)
    updatedAfter := time.Now() // time.Time | filter by updated after date (optional)
    userId := int32(56) // int32 | filter by user id (optional)
    taskId := int32(56) // int32 | filter by task id (optional)
    projectId := int32(56) // int32 | filter by project id (optional)
    pageSize := int32(56) // int32 | number of items in a page default: 50 (optional)
    page := int32(56) // int32 | page number default: 1 (optional)
    showDeleted := true // bool | include deleted items default: false (optional)
    runningTimersOnly := true // bool | enforce running timers only default: false (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTimers := []string{"Inner_example"} // []string |  (optional)
    fieldsTasks := []string{"Inner_example"} // []string |  (optional)
    fieldsTasklists := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.GETProjectsApiV3TimersJson(context.Background()).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).UserId(userId).TaskId(taskId).ProjectId(projectId).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).RunningTimersOnly(runningTimersOnly).Include(include).FieldsUsers(fieldsUsers).FieldsTimers(fieldsTimers).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsProjects(fieldsProjects).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.GETProjectsApiV3TimersJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3TimersJson`: TimerTimersResponse
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.GETProjectsApiV3TimersJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3TimersJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfterDate** | **time.Time** | filter by updated after date (deprecated, use updatedAfter) | 
 **updatedAfter** | **time.Time** | filter by updated after date | 
 **userId** | **int32** | filter by user id | 
 **taskId** | **int32** | filter by task id | 
 **projectId** | **int32** | filter by project id | 
 **pageSize** | **int32** | number of items in a page default: 50 | 
 **page** | **int32** | page number default: 1 | 
 **showDeleted** | **bool** | include deleted items default: false | 
 **runningTimersOnly** | **bool** | enforce running timers only default: false | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTimers** | **[]string** |  | 
 **fieldsTasks** | **[]string** |  | 
 **fieldsTasklists** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 

### Return type

[**TimerTimersResponse**](TimerTimersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3TimerstimerIdJson

> TimerResponse GETProjectsApiV3TimerstimerIdJson(ctx, timerId).UserId(userId).ShowDeleted(showDeleted).Include(include).FieldsUsers(fieldsUsers).FieldsTimers(fieldsTimers).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsProjects(fieldsProjects).Execute()

Get a specific timer



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
    timerId := int32(56) // int32 | 
    userId := int32(56) // int32 | filter by user id (optional)
    showDeleted := true // bool | filter by task id (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTimers := []string{"Inner_example"} // []string |  (optional)
    fieldsTasks := []string{"Inner_example"} // []string |  (optional)
    fieldsTasklists := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.GETProjectsApiV3TimerstimerIdJson(context.Background(), timerId).UserId(userId).ShowDeleted(showDeleted).Include(include).FieldsUsers(fieldsUsers).FieldsTimers(fieldsTimers).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsProjects(fieldsProjects).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.GETProjectsApiV3TimerstimerIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3TimerstimerIdJson`: TimerResponse
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.GETProjectsApiV3TimerstimerIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3TimerstimerIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **int32** | filter by user id | 
 **showDeleted** | **bool** | filter by task id | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTimers** | **[]string** |  | 
 **fieldsTasks** | **[]string** |  | 
 **fieldsTasklists** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 

### Return type

[**TimerResponse**](TimerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsApiV3MeTimersJson

> TimerResponse POSTProjectsApiV3MeTimersJson(ctx).TimerRequest(timerRequest).Execute()

Create a new timer



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
    timerRequest := *openapiclient.NewTimerRequest() // TimerRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.POSTProjectsApiV3MeTimersJson(context.Background()).TimerRequest(timerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.POSTProjectsApiV3MeTimersJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsApiV3MeTimersJson`: TimerResponse
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.POSTProjectsApiV3MeTimersJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3MeTimersJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timerRequest** | [**TimerRequest**](TimerRequest.md) |  | 

### Return type

[**TimerResponse**](TimerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV3MeTimerstimerIdCompleteJson

> TimerResponse PUTProjectsApiV3MeTimerstimerIdCompleteJson(ctx, timerId).Execute()

Complete a timer by ID



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
    timerId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.PUTProjectsApiV3MeTimerstimerIdCompleteJson(context.Background(), timerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.PUTProjectsApiV3MeTimerstimerIdCompleteJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsApiV3MeTimerstimerIdCompleteJson`: TimerResponse
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.PUTProjectsApiV3MeTimerstimerIdCompleteJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3MeTimerstimerIdCompleteJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TimerResponse**](TimerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV3MeTimerstimerIdJson

> TimerResponse PUTProjectsApiV3MeTimerstimerIdJson(ctx, timerId).TimerRequest(timerRequest).Execute()

Edits a timer



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
    timerId := int32(56) // int32 | 
    timerRequest := *openapiclient.NewTimerRequest() // TimerRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.PUTProjectsApiV3MeTimerstimerIdJson(context.Background(), timerId).TimerRequest(timerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.PUTProjectsApiV3MeTimerstimerIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsApiV3MeTimerstimerIdJson`: TimerResponse
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.PUTProjectsApiV3MeTimerstimerIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3MeTimerstimerIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timerRequest** | [**TimerRequest**](TimerRequest.md) |  | 

### Return type

[**TimerResponse**](TimerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV3MeTimerstimerIdPauseJson

> TimerResponse PUTProjectsApiV3MeTimerstimerIdPauseJson(ctx, timerId).Execute()

Pause a timer by ID



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
    timerId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.PUTProjectsApiV3MeTimerstimerIdPauseJson(context.Background(), timerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.PUTProjectsApiV3MeTimerstimerIdPauseJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsApiV3MeTimerstimerIdPauseJson`: TimerResponse
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.PUTProjectsApiV3MeTimerstimerIdPauseJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3MeTimerstimerIdPauseJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TimerResponse**](TimerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV3MeTimerstimerIdResumeJson

> TimerResponse PUTProjectsApiV3MeTimerstimerIdResumeJson(ctx, timerId).Execute()

Resume a timer by ID



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
    timerId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.PUTProjectsApiV3MeTimerstimerIdResumeJson(context.Background(), timerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.PUTProjectsApiV3MeTimerstimerIdResumeJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsApiV3MeTimerstimerIdResumeJson`: TimerResponse
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.PUTProjectsApiV3MeTimerstimerIdResumeJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3MeTimerstimerIdResumeJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TimerResponse**](TimerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV3MeTimerstimerIdUndeleteJson

> PUTProjectsApiV3MeTimerstimerIdUndeleteJson(ctx, timerId).Execute()

Undelete a timer by ID



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
    timerId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.PUTProjectsApiV3MeTimerstimerIdUndeleteJson(context.Background(), timerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.PUTProjectsApiV3MeTimerstimerIdUndeleteJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3MeTimerstimerIdUndeleteJsonRequest struct via the builder pattern


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

