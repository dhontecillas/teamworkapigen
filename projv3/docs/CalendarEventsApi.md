# \CalendarEventsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3CalendarEventsCsv**](CalendarEventsApi.md#GETProjectsApiV3CalendarEventsCsv) | **Get** /projects/api/v3/calendar/events.csv | Generate agenda report in CSV format
[**GETProjectsApiV3CalendarEventsHtml**](CalendarEventsApi.md#GETProjectsApiV3CalendarEventsHtml) | **Get** /projects/api/v3/calendar/events.html | Generate agenda report in HTML format
[**GETProjectsApiV3CalendarEventsJson**](CalendarEventsApi.md#GETProjectsApiV3CalendarEventsJson) | **Get** /projects/api/v3/calendar/events.json | Retrieve the calendar events
[**GETProjectsApiV3CalendarEventsPdf**](CalendarEventsApi.md#GETProjectsApiV3CalendarEventsPdf) | **Get** /projects/api/v3/calendar/events.pdf | Generate agenda report in PDF format
[**GETProjectsApiV3CalendarEventsXlsx**](CalendarEventsApi.md#GETProjectsApiV3CalendarEventsXlsx) | **Get** /projects/api/v3/calendar/events.xlsx | Generate agenda report in XLSX format



## GETProjectsApiV3CalendarEventsCsv

> GETProjectsApiV3CalendarEventsCsv(ctx).StartDate(startDate).ReportFormat(reportFormat).EndDate(endDate).WithTasks(withTasks).WithMilestones(withMilestones).WithEvents(withEvents).IsReportDownload(isReportDownload).IncludeTags(includeTags).AttendingOnly(attendingOnly).TypeIDs(typeIDs).TargetUserIDs(targetUserIDs).TargetProjectIDs(targetProjectIDs).Execute()

Generate agenda report in CSV format



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
    startDate := time.Now() // string | events that happen after this date (optional)
    reportFormat := "reportFormat_example" // string | define the format of the report (optional)
    endDate := time.Now() // string | events that happen before this date (optional)
    withTasks := true // bool | include tasks (optional) (default to false)
    withMilestones := true // bool | include milestones (optional) (default to false)
    withEvents := true // bool | include events (optional) (default to true)
    isReportDownload := true // bool | generate a report document (optional)
    includeTags := true // bool | should include tags (optional) (default to false)
    attendingOnly := true // bool | when filtering events with targetUserIDs, display only when attending the event (optional) (default to false)
    typeIDs := []int32{int32(123)} // []int32 | filter calendar events to show only the ones in typeIDs (optional)
    targetUserIDs := []int32{int32(123)} // []int32 | filter to show only events for the provided user (optional)
    targetProjectIDs := []int32{int32(123)} // []int32 | filter to show only events for the provided projects (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CalendarEventsApi.GETProjectsApiV3CalendarEventsCsv(context.Background()).StartDate(startDate).ReportFormat(reportFormat).EndDate(endDate).WithTasks(withTasks).WithMilestones(withMilestones).WithEvents(withEvents).IsReportDownload(isReportDownload).IncludeTags(includeTags).AttendingOnly(attendingOnly).TypeIDs(typeIDs).TargetUserIDs(targetUserIDs).TargetProjectIDs(targetProjectIDs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalendarEventsApi.GETProjectsApiV3CalendarEventsCsv``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3CalendarEventsCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | events that happen after this date | 
 **reportFormat** | **string** | define the format of the report | 
 **endDate** | **string** | events that happen before this date | 
 **withTasks** | **bool** | include tasks | [default to false]
 **withMilestones** | **bool** | include milestones | [default to false]
 **withEvents** | **bool** | include events | [default to true]
 **isReportDownload** | **bool** | generate a report document | 
 **includeTags** | **bool** | should include tags | [default to false]
 **attendingOnly** | **bool** | when filtering events with targetUserIDs, display only when attending the event | [default to false]
 **typeIDs** | **[]int32** | filter calendar events to show only the ones in typeIDs | 
 **targetUserIDs** | **[]int32** | filter to show only events for the provided user | 
 **targetProjectIDs** | **[]int32** | filter to show only events for the provided projects | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3CalendarEventsHtml

> GETProjectsApiV3CalendarEventsHtml(ctx).StartDate(startDate).ReportFormat(reportFormat).EndDate(endDate).WithTasks(withTasks).WithMilestones(withMilestones).WithEvents(withEvents).IsReportDownload(isReportDownload).IncludeTags(includeTags).AttendingOnly(attendingOnly).TypeIDs(typeIDs).TargetUserIDs(targetUserIDs).TargetProjectIDs(targetProjectIDs).Execute()

Generate agenda report in HTML format



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
    startDate := time.Now() // string | events that happen after this date (optional)
    reportFormat := "reportFormat_example" // string | define the format of the report (optional)
    endDate := time.Now() // string | events that happen before this date (optional)
    withTasks := true // bool | include tasks (optional) (default to false)
    withMilestones := true // bool | include milestones (optional) (default to false)
    withEvents := true // bool | include events (optional) (default to true)
    isReportDownload := true // bool | generate a report document (optional)
    includeTags := true // bool | should include tags (optional) (default to false)
    attendingOnly := true // bool | when filtering events with targetUserIDs, display only when attending the event (optional) (default to false)
    typeIDs := []int32{int32(123)} // []int32 | filter calendar events to show only the ones in typeIDs (optional)
    targetUserIDs := []int32{int32(123)} // []int32 | filter to show only events for the provided user (optional)
    targetProjectIDs := []int32{int32(123)} // []int32 | filter to show only events for the provided projects (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CalendarEventsApi.GETProjectsApiV3CalendarEventsHtml(context.Background()).StartDate(startDate).ReportFormat(reportFormat).EndDate(endDate).WithTasks(withTasks).WithMilestones(withMilestones).WithEvents(withEvents).IsReportDownload(isReportDownload).IncludeTags(includeTags).AttendingOnly(attendingOnly).TypeIDs(typeIDs).TargetUserIDs(targetUserIDs).TargetProjectIDs(targetProjectIDs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalendarEventsApi.GETProjectsApiV3CalendarEventsHtml``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3CalendarEventsHtmlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | events that happen after this date | 
 **reportFormat** | **string** | define the format of the report | 
 **endDate** | **string** | events that happen before this date | 
 **withTasks** | **bool** | include tasks | [default to false]
 **withMilestones** | **bool** | include milestones | [default to false]
 **withEvents** | **bool** | include events | [default to true]
 **isReportDownload** | **bool** | generate a report document | 
 **includeTags** | **bool** | should include tags | [default to false]
 **attendingOnly** | **bool** | when filtering events with targetUserIDs, display only when attending the event | [default to false]
 **typeIDs** | **[]int32** | filter calendar events to show only the ones in typeIDs | 
 **targetUserIDs** | **[]int32** | filter to show only events for the provided user | 
 **targetProjectIDs** | **[]int32** | filter to show only events for the provided projects | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3CalendarEventsJson

> CalendarEventsResponse GETProjectsApiV3CalendarEventsJson(ctx).StartDate(startDate).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).EndDate(endDate).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).OnlyStarredProjects(onlyStarredProjects).MatchAllProjectTags(matchAllProjectTags).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).FieldsCalendarEvents(fieldsCalendarEvents).Execute()

Retrieve the calendar events



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
    startDate := time.Now() // string | events that happen after this date (optional)
    projectStatuses := "projectStatuses_example" // string | filter by project statuses (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "date")
    endDate := time.Now() // string | events that happen before this date (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    matchAllProjectTags := true // bool | match all project tags (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    fieldsCalendarEvents := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CalendarEventsApi.GETProjectsApiV3CalendarEventsJson(context.Background()).StartDate(startDate).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).EndDate(endDate).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).OnlyStarredProjects(onlyStarredProjects).MatchAllProjectTags(matchAllProjectTags).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).FieldsCalendarEvents(fieldsCalendarEvents).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalendarEventsApi.GETProjectsApiV3CalendarEventsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3CalendarEventsJson`: CalendarEventsResponse
    fmt.Fprintf(os.Stdout, "Response from `CalendarEventsApi.GETProjectsApiV3CalendarEventsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3CalendarEventsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | events that happen after this date | 
 **projectStatuses** | **string** | filter by project statuses | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;date&quot;]
 **endDate** | **string** | events that happen before this date | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **matchAllProjectTags** | **bool** | match all project tags | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **fieldsCalendarEvents** | **[]string** |  | 

### Return type

[**CalendarEventsResponse**](CalendarEventsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3CalendarEventsPdf

> GETProjectsApiV3CalendarEventsPdf(ctx).StartDate(startDate).ReportFormat(reportFormat).EndDate(endDate).WithTasks(withTasks).WithMilestones(withMilestones).WithEvents(withEvents).IsReportDownload(isReportDownload).IncludeTags(includeTags).AttendingOnly(attendingOnly).TypeIDs(typeIDs).TargetUserIDs(targetUserIDs).TargetProjectIDs(targetProjectIDs).Execute()

Generate agenda report in PDF format



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
    startDate := time.Now() // string | events that happen after this date (optional)
    reportFormat := "reportFormat_example" // string | define the format of the report (optional)
    endDate := time.Now() // string | events that happen before this date (optional)
    withTasks := true // bool | include tasks (optional) (default to false)
    withMilestones := true // bool | include milestones (optional) (default to false)
    withEvents := true // bool | include events (optional) (default to true)
    isReportDownload := true // bool | generate a report document (optional)
    includeTags := true // bool | should include tags (optional) (default to false)
    attendingOnly := true // bool | when filtering events with targetUserIDs, display only when attending the event (optional) (default to false)
    typeIDs := []int32{int32(123)} // []int32 | filter calendar events to show only the ones in typeIDs (optional)
    targetUserIDs := []int32{int32(123)} // []int32 | filter to show only events for the provided user (optional)
    targetProjectIDs := []int32{int32(123)} // []int32 | filter to show only events for the provided projects (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CalendarEventsApi.GETProjectsApiV3CalendarEventsPdf(context.Background()).StartDate(startDate).ReportFormat(reportFormat).EndDate(endDate).WithTasks(withTasks).WithMilestones(withMilestones).WithEvents(withEvents).IsReportDownload(isReportDownload).IncludeTags(includeTags).AttendingOnly(attendingOnly).TypeIDs(typeIDs).TargetUserIDs(targetUserIDs).TargetProjectIDs(targetProjectIDs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalendarEventsApi.GETProjectsApiV3CalendarEventsPdf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3CalendarEventsPdfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | events that happen after this date | 
 **reportFormat** | **string** | define the format of the report | 
 **endDate** | **string** | events that happen before this date | 
 **withTasks** | **bool** | include tasks | [default to false]
 **withMilestones** | **bool** | include milestones | [default to false]
 **withEvents** | **bool** | include events | [default to true]
 **isReportDownload** | **bool** | generate a report document | 
 **includeTags** | **bool** | should include tags | [default to false]
 **attendingOnly** | **bool** | when filtering events with targetUserIDs, display only when attending the event | [default to false]
 **typeIDs** | **[]int32** | filter calendar events to show only the ones in typeIDs | 
 **targetUserIDs** | **[]int32** | filter to show only events for the provided user | 
 **targetProjectIDs** | **[]int32** | filter to show only events for the provided projects | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3CalendarEventsXlsx

> GETProjectsApiV3CalendarEventsXlsx(ctx).StartDate(startDate).ReportFormat(reportFormat).EndDate(endDate).WithTasks(withTasks).WithMilestones(withMilestones).WithEvents(withEvents).IsReportDownload(isReportDownload).IncludeTags(includeTags).AttendingOnly(attendingOnly).TypeIDs(typeIDs).TargetUserIDs(targetUserIDs).TargetProjectIDs(targetProjectIDs).Execute()

Generate agenda report in XLSX format



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
    startDate := time.Now() // string | events that happen after this date (optional)
    reportFormat := "reportFormat_example" // string | define the format of the report (optional)
    endDate := time.Now() // string | events that happen before this date (optional)
    withTasks := true // bool | include tasks (optional) (default to false)
    withMilestones := true // bool | include milestones (optional) (default to false)
    withEvents := true // bool | include events (optional) (default to true)
    isReportDownload := true // bool | generate a report document (optional)
    includeTags := true // bool | should include tags (optional) (default to false)
    attendingOnly := true // bool | when filtering events with targetUserIDs, display only when attending the event (optional) (default to false)
    typeIDs := []int32{int32(123)} // []int32 | filter calendar events to show only the ones in typeIDs (optional)
    targetUserIDs := []int32{int32(123)} // []int32 | filter to show only events for the provided user (optional)
    targetProjectIDs := []int32{int32(123)} // []int32 | filter to show only events for the provided projects (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CalendarEventsApi.GETProjectsApiV3CalendarEventsXlsx(context.Background()).StartDate(startDate).ReportFormat(reportFormat).EndDate(endDate).WithTasks(withTasks).WithMilestones(withMilestones).WithEvents(withEvents).IsReportDownload(isReportDownload).IncludeTags(includeTags).AttendingOnly(attendingOnly).TypeIDs(typeIDs).TargetUserIDs(targetUserIDs).TargetProjectIDs(targetProjectIDs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalendarEventsApi.GETProjectsApiV3CalendarEventsXlsx``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3CalendarEventsXlsxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | events that happen after this date | 
 **reportFormat** | **string** | define the format of the report | 
 **endDate** | **string** | events that happen before this date | 
 **withTasks** | **bool** | include tasks | [default to false]
 **withMilestones** | **bool** | include milestones | [default to false]
 **withEvents** | **bool** | include events | [default to true]
 **isReportDownload** | **bool** | generate a report document | 
 **includeTags** | **bool** | should include tags | [default to false]
 **attendingOnly** | **bool** | when filtering events with targetUserIDs, display only when attending the event | [default to false]
 **typeIDs** | **[]int32** | filter calendar events to show only the ones in typeIDs | 
 **targetUserIDs** | **[]int32** | filter to show only events for the provided user | 
 **targetProjectIDs** | **[]int32** | filter to show only events for the provided projects | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

