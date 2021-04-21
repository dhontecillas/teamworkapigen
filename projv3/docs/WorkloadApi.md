# \WorkloadApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3WorkloadCsv**](WorkloadApi.md#GETProjectsApiV3WorkloadCsv) | **Get** /projects/api/v3/workload.csv | Generate workload report in CSV format
[**GETProjectsApiV3WorkloadHtml**](WorkloadApi.md#GETProjectsApiV3WorkloadHtml) | **Get** /projects/api/v3/workload.html | Generate workload report in HTML format
[**GETProjectsApiV3WorkloadPdf**](WorkloadApi.md#GETProjectsApiV3WorkloadPdf) | **Get** /projects/api/v3/workload.pdf | Generate workload report in PDF format
[**GETProjectsApiV3WorkloadPlannersJson**](WorkloadApi.md#GETProjectsApiV3WorkloadPlannersJson) | **Get** /projects/api/v3/workload/planners.json | Retrieve user workload planner
[**GETProjectsApiV3WorkloadXlsx**](WorkloadApi.md#GETProjectsApiV3WorkloadXlsx) | **Get** /projects/api/v3/workload.xlsx | Generate workload report in XLSX format.



## GETProjectsApiV3WorkloadCsv

> GETProjectsApiV3WorkloadCsv(ctx).StartDate(startDate).SearchTerm(searchTerm).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).EndDate(endDate).ProjectHealths(projectHealths).ProrataEstimatedTime(prorataEstimatedTime).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllExcludedTags(matchAllExcludedTags).IsReportDownload(isReportDownload).IncludeTasksWithoutDates(includeTasksWithoutDates).DistributeEstimatedTimeToAssignees(distributeEstimatedTimeToAssignees).TagIds(tagIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ExcludeTagIds(excludeTagIds).Execute()

Generate workload report in CSV format



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
    startDate := time.Now() // string | filter by start date (optional)
    searchTerm := "searchTerm_example" // string | filter by search term (optional)
    reportFormat := "reportFormat_example" // string | define the format of the report (optional)
    projectStatuses := "projectStatuses_example" // string | filter by project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "user")
    endDate := time.Now() // string | filter by end date (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    prorataEstimatedTime := true // bool | enable prorata estimated time (optional)
    onlyUntaggedTasks := true // bool | filter for only untagged tasks (optional)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    matchAllTags := true // bool | math all task tags (optional)
    matchAllProjectTags := true // bool | match all project tags (optional)
    matchAllExcludedTags := true // bool | match all excluded task tags (optional)
    isReportDownload := true // bool | generate a report document (optional)
    includeTasksWithoutDates := true // bool | include tasks without dates (optional)
    distributeEstimatedTimeToAssignees := true // bool | distribute estimated time to assignees (optional)
    tagIds := []int32{int32(123)} // []int32 | filter by task tag ids (optional)
    responsiblePartyIds := []int32{int32(123)} // []int32 | filter by responsible party IDs (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    excludeTagIds := []int32{int32(123)} // []int32 | exclude some task tag ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadApi.GETProjectsApiV3WorkloadCsv(context.Background()).StartDate(startDate).SearchTerm(searchTerm).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).EndDate(endDate).ProjectHealths(projectHealths).ProrataEstimatedTime(prorataEstimatedTime).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllExcludedTags(matchAllExcludedTags).IsReportDownload(isReportDownload).IncludeTasksWithoutDates(includeTasksWithoutDates).DistributeEstimatedTimeToAssignees(distributeEstimatedTimeToAssignees).TagIds(tagIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ExcludeTagIds(excludeTagIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadApi.GETProjectsApiV3WorkloadCsv``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3WorkloadCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | filter by start date | 
 **searchTerm** | **string** | filter by search term | 
 **reportFormat** | **string** | define the format of the report | 
 **projectStatuses** | **string** | filter by project status | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;user&quot;]
 **endDate** | **string** | filter by end date | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **prorataEstimatedTime** | **bool** | enable prorata estimated time | 
 **onlyUntaggedTasks** | **bool** | filter for only untagged tasks | 
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **matchAllTags** | **bool** | math all task tags | 
 **matchAllProjectTags** | **bool** | match all project tags | 
 **matchAllExcludedTags** | **bool** | match all excluded task tags | 
 **isReportDownload** | **bool** | generate a report document | 
 **includeTasksWithoutDates** | **bool** | include tasks without dates | 
 **distributeEstimatedTimeToAssignees** | **bool** | distribute estimated time to assignees | 
 **tagIds** | **[]int32** | filter by task tag ids | 
 **responsiblePartyIds** | **[]int32** | filter by responsible party IDs | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **excludeTagIds** | **[]int32** | exclude some task tag ids | 

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


## GETProjectsApiV3WorkloadHtml

> GETProjectsApiV3WorkloadHtml(ctx).StartDate(startDate).SearchTerm(searchTerm).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).EndDate(endDate).ProjectHealths(projectHealths).ProrataEstimatedTime(prorataEstimatedTime).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllExcludedTags(matchAllExcludedTags).IsReportDownload(isReportDownload).IncludeTasksWithoutDates(includeTasksWithoutDates).DistributeEstimatedTimeToAssignees(distributeEstimatedTimeToAssignees).TagIds(tagIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ExcludeTagIds(excludeTagIds).Execute()

Generate workload report in HTML format



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
    startDate := time.Now() // string | filter by start date (optional)
    searchTerm := "searchTerm_example" // string | filter by search term (optional)
    reportFormat := "reportFormat_example" // string | define the format of the report (optional)
    projectStatuses := "projectStatuses_example" // string | filter by project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "user")
    endDate := time.Now() // string | filter by end date (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    prorataEstimatedTime := true // bool | enable prorata estimated time (optional)
    onlyUntaggedTasks := true // bool | filter for only untagged tasks (optional)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    matchAllTags := true // bool | math all task tags (optional)
    matchAllProjectTags := true // bool | match all project tags (optional)
    matchAllExcludedTags := true // bool | match all excluded task tags (optional)
    isReportDownload := true // bool | generate a report document (optional)
    includeTasksWithoutDates := true // bool | include tasks without dates (optional)
    distributeEstimatedTimeToAssignees := true // bool | distribute estimated time to assignees (optional)
    tagIds := []int32{int32(123)} // []int32 | filter by task tag ids (optional)
    responsiblePartyIds := []int32{int32(123)} // []int32 | filter by responsible party IDs (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    excludeTagIds := []int32{int32(123)} // []int32 | exclude some task tag ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadApi.GETProjectsApiV3WorkloadHtml(context.Background()).StartDate(startDate).SearchTerm(searchTerm).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).EndDate(endDate).ProjectHealths(projectHealths).ProrataEstimatedTime(prorataEstimatedTime).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllExcludedTags(matchAllExcludedTags).IsReportDownload(isReportDownload).IncludeTasksWithoutDates(includeTasksWithoutDates).DistributeEstimatedTimeToAssignees(distributeEstimatedTimeToAssignees).TagIds(tagIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ExcludeTagIds(excludeTagIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadApi.GETProjectsApiV3WorkloadHtml``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3WorkloadHtmlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | filter by start date | 
 **searchTerm** | **string** | filter by search term | 
 **reportFormat** | **string** | define the format of the report | 
 **projectStatuses** | **string** | filter by project status | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;user&quot;]
 **endDate** | **string** | filter by end date | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **prorataEstimatedTime** | **bool** | enable prorata estimated time | 
 **onlyUntaggedTasks** | **bool** | filter for only untagged tasks | 
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **matchAllTags** | **bool** | math all task tags | 
 **matchAllProjectTags** | **bool** | match all project tags | 
 **matchAllExcludedTags** | **bool** | match all excluded task tags | 
 **isReportDownload** | **bool** | generate a report document | 
 **includeTasksWithoutDates** | **bool** | include tasks without dates | 
 **distributeEstimatedTimeToAssignees** | **bool** | distribute estimated time to assignees | 
 **tagIds** | **[]int32** | filter by task tag ids | 
 **responsiblePartyIds** | **[]int32** | filter by responsible party IDs | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **excludeTagIds** | **[]int32** | exclude some task tag ids | 

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


## GETProjectsApiV3WorkloadPdf

> GETProjectsApiV3WorkloadPdf(ctx).StartDate(startDate).SearchTerm(searchTerm).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).EndDate(endDate).ProjectHealths(projectHealths).ProrataEstimatedTime(prorataEstimatedTime).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllExcludedTags(matchAllExcludedTags).IsReportDownload(isReportDownload).IncludeTasksWithoutDates(includeTasksWithoutDates).DistributeEstimatedTimeToAssignees(distributeEstimatedTimeToAssignees).TagIds(tagIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ExcludeTagIds(excludeTagIds).Execute()

Generate workload report in PDF format



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
    startDate := time.Now() // string | filter by start date (optional)
    searchTerm := "searchTerm_example" // string | filter by search term (optional)
    reportFormat := "reportFormat_example" // string | define the format of the report (optional)
    projectStatuses := "projectStatuses_example" // string | filter by project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "user")
    endDate := time.Now() // string | filter by end date (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    prorataEstimatedTime := true // bool | enable prorata estimated time (optional)
    onlyUntaggedTasks := true // bool | filter for only untagged tasks (optional)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    matchAllTags := true // bool | math all task tags (optional)
    matchAllProjectTags := true // bool | match all project tags (optional)
    matchAllExcludedTags := true // bool | match all excluded task tags (optional)
    isReportDownload := true // bool | generate a report document (optional)
    includeTasksWithoutDates := true // bool | include tasks without dates (optional)
    distributeEstimatedTimeToAssignees := true // bool | distribute estimated time to assignees (optional)
    tagIds := []int32{int32(123)} // []int32 | filter by task tag ids (optional)
    responsiblePartyIds := []int32{int32(123)} // []int32 | filter by responsible party IDs (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    excludeTagIds := []int32{int32(123)} // []int32 | exclude some task tag ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadApi.GETProjectsApiV3WorkloadPdf(context.Background()).StartDate(startDate).SearchTerm(searchTerm).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).EndDate(endDate).ProjectHealths(projectHealths).ProrataEstimatedTime(prorataEstimatedTime).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllExcludedTags(matchAllExcludedTags).IsReportDownload(isReportDownload).IncludeTasksWithoutDates(includeTasksWithoutDates).DistributeEstimatedTimeToAssignees(distributeEstimatedTimeToAssignees).TagIds(tagIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ExcludeTagIds(excludeTagIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadApi.GETProjectsApiV3WorkloadPdf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3WorkloadPdfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | filter by start date | 
 **searchTerm** | **string** | filter by search term | 
 **reportFormat** | **string** | define the format of the report | 
 **projectStatuses** | **string** | filter by project status | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;user&quot;]
 **endDate** | **string** | filter by end date | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **prorataEstimatedTime** | **bool** | enable prorata estimated time | 
 **onlyUntaggedTasks** | **bool** | filter for only untagged tasks | 
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **matchAllTags** | **bool** | math all task tags | 
 **matchAllProjectTags** | **bool** | match all project tags | 
 **matchAllExcludedTags** | **bool** | match all excluded task tags | 
 **isReportDownload** | **bool** | generate a report document | 
 **includeTasksWithoutDates** | **bool** | include tasks without dates | 
 **distributeEstimatedTimeToAssignees** | **bool** | distribute estimated time to assignees | 
 **tagIds** | **[]int32** | filter by task tag ids | 
 **responsiblePartyIds** | **[]int32** | filter by responsible party IDs | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **excludeTagIds** | **[]int32** | exclude some task tag ids | 

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


## GETProjectsApiV3WorkloadPlannersJson

> PlannerWorkloadPlannersResponse GETProjectsApiV3WorkloadPlannersJson(ctx).StartDate(startDate).SearchUserName(searchUserName).SearchTaskName(searchTaskName).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).EndDate(endDate).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).SubtractLoggedTimeFromEstimates(subtractLoggedTimeFromEstimates).Prorating(prorating).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyStarredProjects(onlyStarredProjects).MatchAllProjectTags(matchAllProjectTags).IncludeTeamUserIds(includeTeamUserIds).IncludeTasksWithoutDates(includeTasksWithoutDates).IncludeCompletedTasks(includeCompletedTasks).IncludeCompanyUserIds(includeCompanyUserIds).IncludeCalendarEvents(includeCalendarEvents).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).Distribute(distribute).TeamIds(teamIds).TagIds(tagIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).MemberOfProjectIds(memberOfProjectIds).Include(include).FieldsWorkingHours(fieldsWorkingHours).FieldsWorkingHourEntries(fieldsWorkingHourEntries).FieldsUsers(fieldsUsers).FieldsTimelogs(fieldsTimelogs).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsMilestones(fieldsMilestones).FieldsCompanies(fieldsCompanies).FieldsCalendarEvents(fieldsCalendarEvents).ExcludeTagIds(excludeTagIds).CompanyIds(companyIds).AssigneeTeamIds(assigneeTeamIds).AssigneeCompanyIds(assigneeCompanyIds).Execute()

Retrieve user workload planner



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
    startDate := time.Now() // string | define the start date of the planner (optional)
    searchUserName := "searchUserName_example" // string | filter by user name (optional)
    searchTaskName := "searchTaskName_example" // string | filter by task name (optional)
    projectStatuses := "projectStatuses_example" // string | list of project status (optional)
    projectStatus := "projectStatus_example" // string | filter by project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "name")
    endDate := time.Now() // string | define the end date of the planner (optional)
    projectHealths := int32(56) // int32 | list of project health (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    subtractLoggedTimeFromEstimates := true // bool | subtract logged time from task estimated time (optional)
    prorating := true // bool | include tasks that have the start and due dates outside the window range (optional) (default to false)
    onlyUntaggedTasks := true // bool | filter ony untagged tasks (optional)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    matchAllProjectTags := true // bool | enforce all tag ids must be matched (optional)
    includeTeamUserIds := true // bool | include team users to the responsible party ids (optional) (default to true)
    includeTasksWithoutDates := true // bool | include tasks without start or due date (optional)
    includeCompletedTasks := true // bool | include completed tasks (optional)
    includeCompanyUserIds := true // bool | include company users to the responsible party ids (optional) (default to true)
    includeCalendarEvents := true // bool | include calendar events on capacity calculation (optional)
    includeAssigneeTeams := true // bool | include teams related to the responsible party ids (optional)
    includeAssigneeCompanies := true // bool | include companies related to the responsible party ids (optional)
    includeArchivedProjects := true // bool | include archived projects (optional)
    distribute := true // bool | distribute the estimated time for a task between all the assignees (optional)
    teamIds := []int32{int32(123)} // []int32 | filter by member of team ids (optional)
    tagIds := []int32{int32(123)} // []int32 | filter by task tag ids (optional)
    responsiblePartyIds := []int32{int32(123)} // []int32 | filter by responsible party ids (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    memberOfProjectIds := []int32{int32(123)} // []int32 | filter by member of project ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsWorkingHours := []string{"Inner_example"} // []string |  (optional)
    fieldsWorkingHourEntries := []string{"Inner_example"} // []string |  (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTimelogs := []string{"Inner_example"} // []string |  (optional)
    fieldsTasks := []string{"Inner_example"} // []string |  (optional)
    fieldsTasklists := []string{"Inner_example"} // []string |  (optional)
    fieldsMilestones := []string{"Inner_example"} // []string |  (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)
    fieldsCalendarEvents := []string{"Inner_example"} // []string |  (optional)
    excludeTagIds := []int32{int32(123)} // []int32 | filter by removing task tag ids (optional)
    companyIds := []int32{int32(123)} // []int32 | filter by user company id (optional)
    assigneeTeamIds := []int32{int32(123)} // []int32 | filter by assignee team ids (optional)
    assigneeCompanyIds := []int32{int32(123)} // []int32 | filter by assignee company ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadApi.GETProjectsApiV3WorkloadPlannersJson(context.Background()).StartDate(startDate).SearchUserName(searchUserName).SearchTaskName(searchTaskName).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).EndDate(endDate).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).SubtractLoggedTimeFromEstimates(subtractLoggedTimeFromEstimates).Prorating(prorating).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyStarredProjects(onlyStarredProjects).MatchAllProjectTags(matchAllProjectTags).IncludeTeamUserIds(includeTeamUserIds).IncludeTasksWithoutDates(includeTasksWithoutDates).IncludeCompletedTasks(includeCompletedTasks).IncludeCompanyUserIds(includeCompanyUserIds).IncludeCalendarEvents(includeCalendarEvents).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).Distribute(distribute).TeamIds(teamIds).TagIds(tagIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).MemberOfProjectIds(memberOfProjectIds).Include(include).FieldsWorkingHours(fieldsWorkingHours).FieldsWorkingHourEntries(fieldsWorkingHourEntries).FieldsUsers(fieldsUsers).FieldsTimelogs(fieldsTimelogs).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsMilestones(fieldsMilestones).FieldsCompanies(fieldsCompanies).FieldsCalendarEvents(fieldsCalendarEvents).ExcludeTagIds(excludeTagIds).CompanyIds(companyIds).AssigneeTeamIds(assigneeTeamIds).AssigneeCompanyIds(assigneeCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadApi.GETProjectsApiV3WorkloadPlannersJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3WorkloadPlannersJson`: PlannerWorkloadPlannersResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkloadApi.GETProjectsApiV3WorkloadPlannersJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3WorkloadPlannersJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | define the start date of the planner | 
 **searchUserName** | **string** | filter by user name | 
 **searchTaskName** | **string** | filter by task name | 
 **projectStatuses** | **string** | list of project status | 
 **projectStatus** | **string** | filter by project status | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;name&quot;]
 **endDate** | **string** | define the end date of the planner | 
 **projectHealths** | **int32** | list of project health | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **subtractLoggedTimeFromEstimates** | **bool** | subtract logged time from task estimated time | 
 **prorating** | **bool** | include tasks that have the start and due dates outside the window range | [default to false]
 **onlyUntaggedTasks** | **bool** | filter ony untagged tasks | 
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **matchAllProjectTags** | **bool** | enforce all tag ids must be matched | 
 **includeTeamUserIds** | **bool** | include team users to the responsible party ids | [default to true]
 **includeTasksWithoutDates** | **bool** | include tasks without start or due date | 
 **includeCompletedTasks** | **bool** | include completed tasks | 
 **includeCompanyUserIds** | **bool** | include company users to the responsible party ids | [default to true]
 **includeCalendarEvents** | **bool** | include calendar events on capacity calculation | 
 **includeAssigneeTeams** | **bool** | include teams related to the responsible party ids | 
 **includeAssigneeCompanies** | **bool** | include companies related to the responsible party ids | 
 **includeArchivedProjects** | **bool** | include archived projects | 
 **distribute** | **bool** | distribute the estimated time for a task between all the assignees | 
 **teamIds** | **[]int32** | filter by member of team ids | 
 **tagIds** | **[]int32** | filter by task tag ids | 
 **responsiblePartyIds** | **[]int32** | filter by responsible party ids | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **memberOfProjectIds** | **[]int32** | filter by member of project ids | 
 **include** | **[]string** | include | 
 **fieldsWorkingHours** | **[]string** |  | 
 **fieldsWorkingHourEntries** | **[]string** |  | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTimelogs** | **[]string** |  | 
 **fieldsTasks** | **[]string** |  | 
 **fieldsTasklists** | **[]string** |  | 
 **fieldsMilestones** | **[]string** |  | 
 **fieldsCompanies** | **[]string** |  | 
 **fieldsCalendarEvents** | **[]string** |  | 
 **excludeTagIds** | **[]int32** | filter by removing task tag ids | 
 **companyIds** | **[]int32** | filter by user company id | 
 **assigneeTeamIds** | **[]int32** | filter by assignee team ids | 
 **assigneeCompanyIds** | **[]int32** | filter by assignee company ids | 

### Return type

[**PlannerWorkloadPlannersResponse**](PlannerWorkloadPlannersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3WorkloadXlsx

> GETProjectsApiV3WorkloadXlsx(ctx).StartDate(startDate).SearchTerm(searchTerm).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).EndDate(endDate).ProjectHealths(projectHealths).ProrataEstimatedTime(prorataEstimatedTime).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllExcludedTags(matchAllExcludedTags).IsReportDownload(isReportDownload).IncludeTasksWithoutDates(includeTasksWithoutDates).DistributeEstimatedTimeToAssignees(distributeEstimatedTimeToAssignees).TagIds(tagIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ExcludeTagIds(excludeTagIds).Execute()

Generate workload report in XLSX format.



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
    startDate := time.Now() // string | filter by start date (optional)
    searchTerm := "searchTerm_example" // string | filter by search term (optional)
    reportFormat := "reportFormat_example" // string | define the format of the report (optional)
    projectStatuses := "projectStatuses_example" // string | filter by project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "user")
    endDate := time.Now() // string | filter by end date (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    prorataEstimatedTime := true // bool | enable prorata estimated time (optional)
    onlyUntaggedTasks := true // bool | filter for only untagged tasks (optional)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    matchAllTags := true // bool | math all task tags (optional)
    matchAllProjectTags := true // bool | match all project tags (optional)
    matchAllExcludedTags := true // bool | match all excluded task tags (optional)
    isReportDownload := true // bool | generate a report document (optional)
    includeTasksWithoutDates := true // bool | include tasks without dates (optional)
    distributeEstimatedTimeToAssignees := true // bool | distribute estimated time to assignees (optional)
    tagIds := []int32{int32(123)} // []int32 | filter by task tag ids (optional)
    responsiblePartyIds := []int32{int32(123)} // []int32 | filter by responsible party IDs (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    excludeTagIds := []int32{int32(123)} // []int32 | exclude some task tag ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadApi.GETProjectsApiV3WorkloadXlsx(context.Background()).StartDate(startDate).SearchTerm(searchTerm).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).EndDate(endDate).ProjectHealths(projectHealths).ProrataEstimatedTime(prorataEstimatedTime).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllExcludedTags(matchAllExcludedTags).IsReportDownload(isReportDownload).IncludeTasksWithoutDates(includeTasksWithoutDates).DistributeEstimatedTimeToAssignees(distributeEstimatedTimeToAssignees).TagIds(tagIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ExcludeTagIds(excludeTagIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadApi.GETProjectsApiV3WorkloadXlsx``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3WorkloadXlsxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | filter by start date | 
 **searchTerm** | **string** | filter by search term | 
 **reportFormat** | **string** | define the format of the report | 
 **projectStatuses** | **string** | filter by project status | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;user&quot;]
 **endDate** | **string** | filter by end date | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **prorataEstimatedTime** | **bool** | enable prorata estimated time | 
 **onlyUntaggedTasks** | **bool** | filter for only untagged tasks | 
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **matchAllTags** | **bool** | math all task tags | 
 **matchAllProjectTags** | **bool** | match all project tags | 
 **matchAllExcludedTags** | **bool** | match all excluded task tags | 
 **isReportDownload** | **bool** | generate a report document | 
 **includeTasksWithoutDates** | **bool** | include tasks without dates | 
 **distributeEstimatedTimeToAssignees** | **bool** | distribute estimated time to assignees | 
 **tagIds** | **[]int32** | filter by task tag ids | 
 **responsiblePartyIds** | **[]int32** | filter by responsible party IDs | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **excludeTagIds** | **[]int32** | exclude some task tag ids | 

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

