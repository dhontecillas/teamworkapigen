# \ProjectsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3ProjectsCsv**](ProjectsApi.md#GETProjectsApiV3ProjectsCsv) | **Get** /projects/api/v3/projects.csv | Generate project report in CSV format
[**GETProjectsApiV3ProjectsHtml**](ProjectsApi.md#GETProjectsApiV3ProjectsHtml) | **Get** /projects/api/v3/projects.html | Generate project report in HTML format
[**GETProjectsApiV3ProjectsJson**](ProjectsApi.md#GETProjectsApiV3ProjectsJson) | **Get** /projects/api/v3/projects.json | Returns a list of projects
[**GETProjectsApiV3ProjectsMetricsActiveJson**](ProjectsApi.md#GETProjectsApiV3ProjectsMetricsActiveJson) | **Get** /projects/api/v3/projects/metrics/active.json | Get the total count of active projects
[**GETProjectsApiV3ProjectsMetricsBillableJson**](ProjectsApi.md#GETProjectsApiV3ProjectsMetricsBillableJson) | **Get** /projects/api/v3/projects/metrics/billable.json | Get the total billable time per project
[**GETProjectsApiV3ProjectsMetricsHealthsJson**](ProjectsApi.md#GETProjectsApiV3ProjectsMetricsHealthsJson) | **Get** /projects/api/v3/projects/metrics/healths.json | Get a health summary for all projects
[**GETProjectsApiV3ProjectsMetricsInvoiceJson**](ProjectsApi.md#GETProjectsApiV3ProjectsMetricsInvoiceJson) | **Get** /projects/api/v3/projects/metrics/invoice.json | Return open invoices across all projects
[**GETProjectsApiV3ProjectsMetricsOwnersJson**](ProjectsApi.md#GETProjectsApiV3ProjectsMetricsOwnersJson) | **Get** /projects/api/v3/projects/metrics/owners.json | Get number of owned and un-assigned projects
[**GETProjectsApiV3ProjectsMetricsUnbilledJson**](ProjectsApi.md#GETProjectsApiV3ProjectsMetricsUnbilledJson) | **Get** /projects/api/v3/projects/metrics/unbilled.json | Return un-billed expenses across all projects
[**GETProjectsApiV3ProjectsPdf**](ProjectsApi.md#GETProjectsApiV3ProjectsPdf) | **Get** /projects/api/v3/projects.pdf | Generate project report in PDF format
[**GETProjectsApiV3ProjectsTeamworkSamplesJson**](ProjectsApi.md#GETProjectsApiV3ProjectsTeamworkSamplesJson) | **Get** /projects/api/v3/projects/teamwork/samples.json | Returns a list of teamwork sample projects
[**GETProjectsApiV3ProjectsTemplatesJson**](ProjectsApi.md#GETProjectsApiV3ProjectsTemplatesJson) | **Get** /projects/api/v3/projects/templates.json | Returns a list of projects templates
[**GETProjectsApiV3ProjectsXlsx**](ProjectsApi.md#GETProjectsApiV3ProjectsXlsx) | **Get** /projects/api/v3/projects.xlsx | Generate project report in XLSX format
[**GETProjectsApiV3ProjectsprojectIdFeatureorderJson**](ProjectsApi.md#GETProjectsApiV3ProjectsprojectIdFeatureorderJson) | **Get** /projects/api/v3/projects/{projectId}/featureorder.json | Returns a project features order to display in tab
[**GETProjectsApiV3ProjectsprojectIdJson**](ProjectsApi.md#GETProjectsApiV3ProjectsprojectIdJson) | **Get** /projects/api/v3/projects/{projectId}.json | Returns a project
[**PUTProjectsApiV3ProjectsFeatureorderJson**](ProjectsApi.md#PUTProjectsApiV3ProjectsFeatureorderJson) | **Put** /projects/api/v3/projects/featureorder.json | Sets the default features order to display in tab
[**PUTProjectsApiV3ProjectsprojectIdFeatureorderJson**](ProjectsApi.md#PUTProjectsApiV3ProjectsprojectIdFeatureorderJson) | **Put** /projects/api/v3/projects/{projectId}/featureorder.json | Sets the the features order to display in tab



## GETProjectsApiV3ProjectsCsv

> GETProjectsApiV3ProjectsCsv(ctx).UpdatedAfter(updatedAfter).SearchTerm(searchTerm).ReportType(reportType).ReportFormat(reportFormat).ProjectType(projectType).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).MinLastActivityDate(minLastActivityDate).MaxLastActivityDate(maxLastActivityDate).UserId(userId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).OnlyStarredProjects(onlyStarredProjects).OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership).OnlyArchivedProjects(onlyArchivedProjects).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeProjectUserInfo(includeProjectUserInfo).IncludeCustomFields(includeCustomFields).IncludeCompletedStatus(includeCompletedStatus).IncludeArchivedProjects(includeArchivedProjects).HideObservedProjects(hideObservedProjects).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsProjectcategories(fieldsProjectcategories).FieldsProjectUpdates(fieldsProjectUpdates).FieldsProjectBudgets(fieldsProjectBudgets).FieldsPortfolioColumns(fieldsPortfolioColumns).FieldsPortfolioCards(fieldsPortfolioCards).FieldsPortfolioBoards(fieldsPortfolioBoards).FieldsCustomfields(fieldsCustomfields).FieldsCustomfieldProjects(fieldsCustomfieldProjects).FieldsCompanies(fieldsCompanies).Execute()

Generate project report in CSV format



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
    updatedAfter := time.Now() // time.Time | updated after (optional)
    searchTerm := "searchTerm_example" // string | filter by project name (optional)
    reportType := "reportType_example" // string | define the format of the report (optional) (default to "project")
    reportFormat := "reportFormat_example" // string | define the format of the report (optional) (default to "pdf")
    projectType := "projectType_example" // string | filter by project type (optional)
    projectStatuses := "projectStatuses_example" // string | filter by project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "name")
    minLastActivityDate := time.Now() // string | filter by min last activity date (optional)
    maxLastActivityDate := time.Now() // string | filter by max last activity date (optional)
    userId := int32(56) // int32 | filter by user id (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    pageSize := int32(56) // int32 | number of items in a page (not used when generating reports) (optional) (default to 50)
    page := int32(56) // int32 | page number (not used when generating reports) (optional) (default to 1)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    onlyProjectsWithExplicitMembership := true // bool | only show projects with explicit membership (optional) (default to false)
    onlyArchivedProjects := true // bool | return only archived projects (optional) (default to false)
    matchAllProjectTags := true // bool | match all project tags (optional)
    isReportDownload := true // bool | generate a report document (optional) (default to false)
    includeProjectUserInfo := true // bool | fetch user-specific data such as isStarred (optional) (default to false)
    includeCustomFields := true // bool | include custom fields (optional) (default to false)
    includeCompletedStatus := true // bool | optional to include completed projects when filtering by project statuses \"current,late\". (optional)
    includeArchivedProjects := true // bool | include archived projects (optional) (default to false)
    hideObservedProjects := true // bool | hide projects where the logged-in user is just an observer (optional) (default to false)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (not used when generating reports) (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTags := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsProjectcategories := []string{"Inner_example"} // []string |  (optional)
    fieldsProjectUpdates := []string{"Inner_example"} // []string |  (optional)
    fieldsProjectBudgets := []string{"Inner_example"} // []string |  (optional)
    fieldsPortfolioColumns := []string{"Inner_example"} // []string |  (optional)
    fieldsPortfolioCards := []string{"Inner_example"} // []string |  (optional)
    fieldsPortfolioBoards := []string{"Inner_example"} // []string |  (optional)
    fieldsCustomfields := []string{"Inner_example"} // []string |  (optional)
    fieldsCustomfieldProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.GETProjectsApiV3ProjectsCsv(context.Background()).UpdatedAfter(updatedAfter).SearchTerm(searchTerm).ReportType(reportType).ReportFormat(reportFormat).ProjectType(projectType).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).MinLastActivityDate(minLastActivityDate).MaxLastActivityDate(maxLastActivityDate).UserId(userId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).OnlyStarredProjects(onlyStarredProjects).OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership).OnlyArchivedProjects(onlyArchivedProjects).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeProjectUserInfo(includeProjectUserInfo).IncludeCustomFields(includeCustomFields).IncludeCompletedStatus(includeCompletedStatus).IncludeArchivedProjects(includeArchivedProjects).HideObservedProjects(hideObservedProjects).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsProjectcategories(fieldsProjectcategories).FieldsProjectUpdates(fieldsProjectUpdates).FieldsProjectBudgets(fieldsProjectBudgets).FieldsPortfolioColumns(fieldsPortfolioColumns).FieldsPortfolioCards(fieldsPortfolioCards).FieldsPortfolioBoards(fieldsPortfolioBoards).FieldsCustomfields(fieldsCustomfields).FieldsCustomfieldProjects(fieldsCustomfieldProjects).FieldsCompanies(fieldsCompanies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GETProjectsApiV3ProjectsCsv``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfter** | **time.Time** | updated after | 
 **searchTerm** | **string** | filter by project name | 
 **reportType** | **string** | define the format of the report | [default to &quot;project&quot;]
 **reportFormat** | **string** | define the format of the report | [default to &quot;pdf&quot;]
 **projectType** | **string** | filter by project type | 
 **projectStatuses** | **string** | filter by project status | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;name&quot;]
 **minLastActivityDate** | **string** | filter by min last activity date | 
 **maxLastActivityDate** | **string** | filter by max last activity date | 
 **userId** | **int32** | filter by user id | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **pageSize** | **int32** | number of items in a page (not used when generating reports) | [default to 50]
 **page** | **int32** | page number (not used when generating reports) | [default to 1]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **onlyProjectsWithExplicitMembership** | **bool** | only show projects with explicit membership | [default to false]
 **onlyArchivedProjects** | **bool** | return only archived projects | [default to false]
 **matchAllProjectTags** | **bool** | match all project tags | 
 **isReportDownload** | **bool** | generate a report document | [default to false]
 **includeProjectUserInfo** | **bool** | fetch user-specific data such as isStarred | [default to false]
 **includeCustomFields** | **bool** | include custom fields | [default to false]
 **includeCompletedStatus** | **bool** | optional to include completed projects when filtering by project statuses \&quot;current,late\&quot;. | 
 **includeArchivedProjects** | **bool** | include archived projects | [default to false]
 **hideObservedProjects** | **bool** | hide projects where the logged-in user is just an observer | [default to false]
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include (not used when generating reports) | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTags** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsProjectcategories** | **[]string** |  | 
 **fieldsProjectUpdates** | **[]string** |  | 
 **fieldsProjectBudgets** | **[]string** |  | 
 **fieldsPortfolioColumns** | **[]string** |  | 
 **fieldsPortfolioCards** | **[]string** |  | 
 **fieldsPortfolioBoards** | **[]string** |  | 
 **fieldsCustomfields** | **[]string** |  | 
 **fieldsCustomfieldProjects** | **[]string** |  | 
 **fieldsCompanies** | **[]string** |  | 

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


## GETProjectsApiV3ProjectsHtml

> GETProjectsApiV3ProjectsHtml(ctx).UpdatedAfter(updatedAfter).SearchTerm(searchTerm).ReportType(reportType).ReportFormat(reportFormat).ProjectType(projectType).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).MinLastActivityDate(minLastActivityDate).MaxLastActivityDate(maxLastActivityDate).UserId(userId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).OnlyStarredProjects(onlyStarredProjects).OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership).OnlyArchivedProjects(onlyArchivedProjects).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeProjectUserInfo(includeProjectUserInfo).IncludeCustomFields(includeCustomFields).IncludeCompletedStatus(includeCompletedStatus).IncludeArchivedProjects(includeArchivedProjects).HideObservedProjects(hideObservedProjects).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsProjectcategories(fieldsProjectcategories).FieldsProjectUpdates(fieldsProjectUpdates).FieldsProjectBudgets(fieldsProjectBudgets).FieldsPortfolioColumns(fieldsPortfolioColumns).FieldsPortfolioCards(fieldsPortfolioCards).FieldsPortfolioBoards(fieldsPortfolioBoards).FieldsCustomfields(fieldsCustomfields).FieldsCustomfieldProjects(fieldsCustomfieldProjects).FieldsCompanies(fieldsCompanies).Execute()

Generate project report in HTML format



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
    updatedAfter := time.Now() // time.Time | updated after (optional)
    searchTerm := "searchTerm_example" // string | filter by project name (optional)
    reportType := "reportType_example" // string | define the format of the report (optional) (default to "project")
    reportFormat := "reportFormat_example" // string | define the format of the report (optional) (default to "pdf")
    projectType := "projectType_example" // string | filter by project type (optional)
    projectStatuses := "projectStatuses_example" // string | filter by project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "name")
    minLastActivityDate := time.Now() // string | filter by min last activity date (optional)
    maxLastActivityDate := time.Now() // string | filter by max last activity date (optional)
    userId := int32(56) // int32 | filter by user id (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    pageSize := int32(56) // int32 | number of items in a page (not used when generating reports) (optional) (default to 50)
    page := int32(56) // int32 | page number (not used when generating reports) (optional) (default to 1)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    onlyProjectsWithExplicitMembership := true // bool | only show projects with explicit membership (optional) (default to false)
    onlyArchivedProjects := true // bool | return only archived projects (optional) (default to false)
    matchAllProjectTags := true // bool | match all project tags (optional)
    isReportDownload := true // bool | generate a report document (optional) (default to false)
    includeProjectUserInfo := true // bool | fetch user-specific data such as isStarred (optional) (default to false)
    includeCustomFields := true // bool | include custom fields (optional) (default to false)
    includeCompletedStatus := true // bool | optional to include completed projects when filtering by project statuses \"current,late\". (optional)
    includeArchivedProjects := true // bool | include archived projects (optional) (default to false)
    hideObservedProjects := true // bool | hide projects where the logged-in user is just an observer (optional) (default to false)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (not used when generating reports) (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTags := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsProjectcategories := []string{"Inner_example"} // []string |  (optional)
    fieldsProjectUpdates := []string{"Inner_example"} // []string |  (optional)
    fieldsProjectBudgets := []string{"Inner_example"} // []string |  (optional)
    fieldsPortfolioColumns := []string{"Inner_example"} // []string |  (optional)
    fieldsPortfolioCards := []string{"Inner_example"} // []string |  (optional)
    fieldsPortfolioBoards := []string{"Inner_example"} // []string |  (optional)
    fieldsCustomfields := []string{"Inner_example"} // []string |  (optional)
    fieldsCustomfieldProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.GETProjectsApiV3ProjectsHtml(context.Background()).UpdatedAfter(updatedAfter).SearchTerm(searchTerm).ReportType(reportType).ReportFormat(reportFormat).ProjectType(projectType).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).MinLastActivityDate(minLastActivityDate).MaxLastActivityDate(maxLastActivityDate).UserId(userId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).OnlyStarredProjects(onlyStarredProjects).OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership).OnlyArchivedProjects(onlyArchivedProjects).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeProjectUserInfo(includeProjectUserInfo).IncludeCustomFields(includeCustomFields).IncludeCompletedStatus(includeCompletedStatus).IncludeArchivedProjects(includeArchivedProjects).HideObservedProjects(hideObservedProjects).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsProjectcategories(fieldsProjectcategories).FieldsProjectUpdates(fieldsProjectUpdates).FieldsProjectBudgets(fieldsProjectBudgets).FieldsPortfolioColumns(fieldsPortfolioColumns).FieldsPortfolioCards(fieldsPortfolioCards).FieldsPortfolioBoards(fieldsPortfolioBoards).FieldsCustomfields(fieldsCustomfields).FieldsCustomfieldProjects(fieldsCustomfieldProjects).FieldsCompanies(fieldsCompanies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GETProjectsApiV3ProjectsHtml``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsHtmlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfter** | **time.Time** | updated after | 
 **searchTerm** | **string** | filter by project name | 
 **reportType** | **string** | define the format of the report | [default to &quot;project&quot;]
 **reportFormat** | **string** | define the format of the report | [default to &quot;pdf&quot;]
 **projectType** | **string** | filter by project type | 
 **projectStatuses** | **string** | filter by project status | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;name&quot;]
 **minLastActivityDate** | **string** | filter by min last activity date | 
 **maxLastActivityDate** | **string** | filter by max last activity date | 
 **userId** | **int32** | filter by user id | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **pageSize** | **int32** | number of items in a page (not used when generating reports) | [default to 50]
 **page** | **int32** | page number (not used when generating reports) | [default to 1]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **onlyProjectsWithExplicitMembership** | **bool** | only show projects with explicit membership | [default to false]
 **onlyArchivedProjects** | **bool** | return only archived projects | [default to false]
 **matchAllProjectTags** | **bool** | match all project tags | 
 **isReportDownload** | **bool** | generate a report document | [default to false]
 **includeProjectUserInfo** | **bool** | fetch user-specific data such as isStarred | [default to false]
 **includeCustomFields** | **bool** | include custom fields | [default to false]
 **includeCompletedStatus** | **bool** | optional to include completed projects when filtering by project statuses \&quot;current,late\&quot;. | 
 **includeArchivedProjects** | **bool** | include archived projects | [default to false]
 **hideObservedProjects** | **bool** | hide projects where the logged-in user is just an observer | [default to false]
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include (not used when generating reports) | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTags** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsProjectcategories** | **[]string** |  | 
 **fieldsProjectUpdates** | **[]string** |  | 
 **fieldsProjectBudgets** | **[]string** |  | 
 **fieldsPortfolioColumns** | **[]string** |  | 
 **fieldsPortfolioCards** | **[]string** |  | 
 **fieldsPortfolioBoards** | **[]string** |  | 
 **fieldsCustomfields** | **[]string** |  | 
 **fieldsCustomfieldProjects** | **[]string** |  | 
 **fieldsCompanies** | **[]string** |  | 

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


## GETProjectsApiV3ProjectsJson

> ProjectProjectsResponse GETProjectsApiV3ProjectsJson(ctx).UpdatedAfter(updatedAfter).SearchTerm(searchTerm).ReportType(reportType).ReportFormat(reportFormat).ProjectType(projectType).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).MinLastActivityDate(minLastActivityDate).MaxLastActivityDate(maxLastActivityDate).UserId(userId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).OnlyStarredProjects(onlyStarredProjects).OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership).OnlyArchivedProjects(onlyArchivedProjects).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeProjectUserInfo(includeProjectUserInfo).IncludeCustomFields(includeCustomFields).IncludeCompletedStatus(includeCompletedStatus).IncludeArchivedProjects(includeArchivedProjects).HideObservedProjects(hideObservedProjects).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsProjectcategories(fieldsProjectcategories).FieldsProjectUpdates(fieldsProjectUpdates).FieldsProjectBudgets(fieldsProjectBudgets).FieldsPortfolioColumns(fieldsPortfolioColumns).FieldsPortfolioCards(fieldsPortfolioCards).FieldsPortfolioBoards(fieldsPortfolioBoards).FieldsCustomfields(fieldsCustomfields).FieldsCustomfieldProjects(fieldsCustomfieldProjects).FieldsCompanies(fieldsCompanies).Execute()

Returns a list of projects



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
    updatedAfter := time.Now() // time.Time | updated after (optional)
    searchTerm := "searchTerm_example" // string | filter by project name (optional)
    reportType := "reportType_example" // string | define the format of the report (optional) (default to "project")
    reportFormat := "reportFormat_example" // string | define the format of the report (optional) (default to "pdf")
    projectType := "projectType_example" // string | filter by project type (optional)
    projectStatuses := "projectStatuses_example" // string | filter by project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "name")
    minLastActivityDate := time.Now() // string | filter by min last activity date (optional)
    maxLastActivityDate := time.Now() // string | filter by max last activity date (optional)
    userId := int32(56) // int32 | filter by user id (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    pageSize := int32(56) // int32 | number of items in a page (not used when generating reports) (optional) (default to 50)
    page := int32(56) // int32 | page number (not used when generating reports) (optional) (default to 1)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    onlyProjectsWithExplicitMembership := true // bool | only show projects with explicit membership (optional) (default to false)
    onlyArchivedProjects := true // bool | return only archived projects (optional) (default to false)
    matchAllProjectTags := true // bool | match all project tags (optional)
    isReportDownload := true // bool | generate a report document (optional) (default to false)
    includeProjectUserInfo := true // bool | fetch user-specific data such as isStarred (optional) (default to false)
    includeCustomFields := true // bool | include custom fields (optional) (default to false)
    includeCompletedStatus := true // bool | optional to include completed projects when filtering by project statuses \"current,late\". (optional)
    includeArchivedProjects := true // bool | include archived projects (optional) (default to false)
    hideObservedProjects := true // bool | hide projects where the logged-in user is just an observer (optional) (default to false)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (not used when generating reports) (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTags := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsProjectcategories := []string{"Inner_example"} // []string |  (optional)
    fieldsProjectUpdates := []string{"Inner_example"} // []string |  (optional)
    fieldsProjectBudgets := []string{"Inner_example"} // []string |  (optional)
    fieldsPortfolioColumns := []string{"Inner_example"} // []string |  (optional)
    fieldsPortfolioCards := []string{"Inner_example"} // []string |  (optional)
    fieldsPortfolioBoards := []string{"Inner_example"} // []string |  (optional)
    fieldsCustomfields := []string{"Inner_example"} // []string |  (optional)
    fieldsCustomfieldProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.GETProjectsApiV3ProjectsJson(context.Background()).UpdatedAfter(updatedAfter).SearchTerm(searchTerm).ReportType(reportType).ReportFormat(reportFormat).ProjectType(projectType).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).MinLastActivityDate(minLastActivityDate).MaxLastActivityDate(maxLastActivityDate).UserId(userId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).OnlyStarredProjects(onlyStarredProjects).OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership).OnlyArchivedProjects(onlyArchivedProjects).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeProjectUserInfo(includeProjectUserInfo).IncludeCustomFields(includeCustomFields).IncludeCompletedStatus(includeCompletedStatus).IncludeArchivedProjects(includeArchivedProjects).HideObservedProjects(hideObservedProjects).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsProjectcategories(fieldsProjectcategories).FieldsProjectUpdates(fieldsProjectUpdates).FieldsProjectBudgets(fieldsProjectBudgets).FieldsPortfolioColumns(fieldsPortfolioColumns).FieldsPortfolioCards(fieldsPortfolioCards).FieldsPortfolioBoards(fieldsPortfolioBoards).FieldsCustomfields(fieldsCustomfields).FieldsCustomfieldProjects(fieldsCustomfieldProjects).FieldsCompanies(fieldsCompanies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GETProjectsApiV3ProjectsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ProjectsJson`: ProjectProjectsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GETProjectsApiV3ProjectsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfter** | **time.Time** | updated after | 
 **searchTerm** | **string** | filter by project name | 
 **reportType** | **string** | define the format of the report | [default to &quot;project&quot;]
 **reportFormat** | **string** | define the format of the report | [default to &quot;pdf&quot;]
 **projectType** | **string** | filter by project type | 
 **projectStatuses** | **string** | filter by project status | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;name&quot;]
 **minLastActivityDate** | **string** | filter by min last activity date | 
 **maxLastActivityDate** | **string** | filter by max last activity date | 
 **userId** | **int32** | filter by user id | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **pageSize** | **int32** | number of items in a page (not used when generating reports) | [default to 50]
 **page** | **int32** | page number (not used when generating reports) | [default to 1]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **onlyProjectsWithExplicitMembership** | **bool** | only show projects with explicit membership | [default to false]
 **onlyArchivedProjects** | **bool** | return only archived projects | [default to false]
 **matchAllProjectTags** | **bool** | match all project tags | 
 **isReportDownload** | **bool** | generate a report document | [default to false]
 **includeProjectUserInfo** | **bool** | fetch user-specific data such as isStarred | [default to false]
 **includeCustomFields** | **bool** | include custom fields | [default to false]
 **includeCompletedStatus** | **bool** | optional to include completed projects when filtering by project statuses \&quot;current,late\&quot;. | 
 **includeArchivedProjects** | **bool** | include archived projects | [default to false]
 **hideObservedProjects** | **bool** | hide projects where the logged-in user is just an observer | [default to false]
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include (not used when generating reports) | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTags** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsProjectcategories** | **[]string** |  | 
 **fieldsProjectUpdates** | **[]string** |  | 
 **fieldsProjectBudgets** | **[]string** |  | 
 **fieldsPortfolioColumns** | **[]string** |  | 
 **fieldsPortfolioCards** | **[]string** |  | 
 **fieldsPortfolioBoards** | **[]string** |  | 
 **fieldsCustomfields** | **[]string** |  | 
 **fieldsCustomfieldProjects** | **[]string** |  | 
 **fieldsCompanies** | **[]string** |  | 

### Return type

[**ProjectProjectsResponse**](ProjectProjectsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3ProjectsMetricsActiveJson

> ActiveResponse GETProjectsApiV3ProjectsMetricsActiveJson(ctx).Execute()

Get the total count of active projects



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
    resp, r, err := api_client.ProjectsApi.GETProjectsApiV3ProjectsMetricsActiveJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GETProjectsApiV3ProjectsMetricsActiveJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ProjectsMetricsActiveJson`: ActiveResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GETProjectsApiV3ProjectsMetricsActiveJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsMetricsActiveJsonRequest struct via the builder pattern


### Return type

[**ActiveResponse**](ActiveResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3ProjectsMetricsBillableJson

> BillableProjectMetricBillablesResponse GETProjectsApiV3ProjectsMetricsBillableJson(ctx).StartDate(startDate).OrderMode(orderMode).EndDate(endDate).Execute()

Get the total billable time per project



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
    startDate := time.Now() // string |  (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "desc")
    endDate := time.Now() // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.GETProjectsApiV3ProjectsMetricsBillableJson(context.Background()).StartDate(startDate).OrderMode(orderMode).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GETProjectsApiV3ProjectsMetricsBillableJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ProjectsMetricsBillableJson`: BillableProjectMetricBillablesResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GETProjectsApiV3ProjectsMetricsBillableJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsMetricsBillableJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** |  | 
 **orderMode** | **string** | order mode | [default to &quot;desc&quot;]
 **endDate** | **string** |  | 

### Return type

[**BillableProjectMetricBillablesResponse**](BillableProjectMetricBillablesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3ProjectsMetricsHealthsJson

> HealthProjectMetricHealthsResponse GETProjectsApiV3ProjectsMetricsHealthsJson(ctx).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).ProjectHealths(projectHealths).OnlyStarredProjects(onlyStarredProjects).MatchAllProjectTags(matchAllProjectTags).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Execute()

Get a health summary for all projects



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
    projectHealths := int32(56) // int32 | project health  0: not set 1: bad 2: ok 3: good (optional)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    matchAllProjectTags := true // bool | match all project tags (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.GETProjectsApiV3ProjectsMetricsHealthsJson(context.Background()).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).ProjectHealths(projectHealths).OnlyStarredProjects(onlyStarredProjects).MatchAllProjectTags(matchAllProjectTags).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GETProjectsApiV3ProjectsMetricsHealthsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ProjectsMetricsHealthsJson`: HealthProjectMetricHealthsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GETProjectsApiV3ProjectsMetricsHealthsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsMetricsHealthsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectStatuses** | **string** | filter by project statuses | 
 **projectStatus** | **string** | filter by project status | 
 **projectHealths** | **int32** | project health  0: not set 1: bad 2: ok 3: good | 
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **matchAllProjectTags** | **bool** | match all project tags | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 

### Return type

[**HealthProjectMetricHealthsResponse**](HealthProjectMetricHealthsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3ProjectsMetricsInvoiceJson

> InvoiceProjectMetricInvoicesResponse GETProjectsApiV3ProjectsMetricsInvoiceJson(ctx).ProjectStatuses(projectStatuses).ProjectHealths(projectHealths).OnlyStarredProjects(onlyStarredProjects).MatchAllProjectTags(matchAllProjectTags).IncludeArchivedProjects(includeArchivedProjects).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Execute()

Return open invoices across all projects



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
    projectStatuses := "projectStatuses_example" // string | project status (optional)
    projectHealths := int32(56) // int32 | project health  0: not set 1: bad 2: ok 3: good (optional)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    matchAllProjectTags := true // bool | match all project tags (optional)
    includeArchivedProjects := true // bool | include archived projects (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.GETProjectsApiV3ProjectsMetricsInvoiceJson(context.Background()).ProjectStatuses(projectStatuses).ProjectHealths(projectHealths).OnlyStarredProjects(onlyStarredProjects).MatchAllProjectTags(matchAllProjectTags).IncludeArchivedProjects(includeArchivedProjects).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GETProjectsApiV3ProjectsMetricsInvoiceJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ProjectsMetricsInvoiceJson`: InvoiceProjectMetricInvoicesResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GETProjectsApiV3ProjectsMetricsInvoiceJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsMetricsInvoiceJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectStatuses** | **string** | project status | 
 **projectHealths** | **int32** | project health  0: not set 1: bad 2: ok 3: good | 
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **matchAllProjectTags** | **bool** | match all project tags | 
 **includeArchivedProjects** | **bool** | include archived projects | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 

### Return type

[**InvoiceProjectMetricInvoicesResponse**](InvoiceProjectMetricInvoicesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3ProjectsMetricsOwnersJson

> OwnerProjectMetricOwnersResponse GETProjectsApiV3ProjectsMetricsOwnersJson(ctx).ProjectStatuses(projectStatuses).OrderMode(orderMode).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).OnlyStarredProjects(onlyStarredProjects).MatchAllProjectTags(matchAllProjectTags).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Execute()

Get number of owned and un-assigned projects



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
    projectStatuses := "projectStatuses_example" // string | project status (optional)
    orderMode := "orderMode_example" // string | sort order (optional) (default to "desc")
    projectHealths := int32(56) // int32 | project health  0: not set 1: bad 2: ok 3: good (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    matchAllProjectTags := true // bool | match all project tags (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.GETProjectsApiV3ProjectsMetricsOwnersJson(context.Background()).ProjectStatuses(projectStatuses).OrderMode(orderMode).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).OnlyStarredProjects(onlyStarredProjects).MatchAllProjectTags(matchAllProjectTags).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GETProjectsApiV3ProjectsMetricsOwnersJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ProjectsMetricsOwnersJson`: OwnerProjectMetricOwnersResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GETProjectsApiV3ProjectsMetricsOwnersJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsMetricsOwnersJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectStatuses** | **string** | project status | 
 **orderMode** | **string** | sort order | [default to &quot;desc&quot;]
 **projectHealths** | **int32** | project health  0: not set 1: bad 2: ok 3: good | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **matchAllProjectTags** | **bool** | match all project tags | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 

### Return type

[**OwnerProjectMetricOwnersResponse**](OwnerProjectMetricOwnersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3ProjectsMetricsUnbilledJson

> UnbilledProjectMetricUnbilledResponse GETProjectsApiV3ProjectsMetricsUnbilledJson(ctx).ProjectStatuses(projectStatuses).ProjectHealths(projectHealths).OnlyStarredProjects(onlyStarredProjects).MatchAllProjectTags(matchAllProjectTags).IncludeArchivedProjects(includeArchivedProjects).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Execute()

Return un-billed expenses across all projects



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
    projectStatuses := "projectStatuses_example" // string | project status (optional)
    projectHealths := int32(56) // int32 | project health  0: not set 1: bad 2: ok 3: good (optional)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    matchAllProjectTags := true // bool | match all project tags (optional)
    includeArchivedProjects := true // bool | include archived projects (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.GETProjectsApiV3ProjectsMetricsUnbilledJson(context.Background()).ProjectStatuses(projectStatuses).ProjectHealths(projectHealths).OnlyStarredProjects(onlyStarredProjects).MatchAllProjectTags(matchAllProjectTags).IncludeArchivedProjects(includeArchivedProjects).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GETProjectsApiV3ProjectsMetricsUnbilledJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ProjectsMetricsUnbilledJson`: UnbilledProjectMetricUnbilledResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GETProjectsApiV3ProjectsMetricsUnbilledJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsMetricsUnbilledJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectStatuses** | **string** | project status | 
 **projectHealths** | **int32** | project health  0: not set 1: bad 2: ok 3: good | 
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **matchAllProjectTags** | **bool** | match all project tags | 
 **includeArchivedProjects** | **bool** | include archived projects | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 

### Return type

[**UnbilledProjectMetricUnbilledResponse**](UnbilledProjectMetricUnbilledResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3ProjectsPdf

> GETProjectsApiV3ProjectsPdf(ctx).UpdatedAfter(updatedAfter).SearchTerm(searchTerm).ReportType(reportType).ReportFormat(reportFormat).ProjectType(projectType).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).MinLastActivityDate(minLastActivityDate).MaxLastActivityDate(maxLastActivityDate).UserId(userId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).OnlyStarredProjects(onlyStarredProjects).OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership).OnlyArchivedProjects(onlyArchivedProjects).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeProjectUserInfo(includeProjectUserInfo).IncludeCustomFields(includeCustomFields).IncludeCompletedStatus(includeCompletedStatus).IncludeArchivedProjects(includeArchivedProjects).HideObservedProjects(hideObservedProjects).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsProjectcategories(fieldsProjectcategories).FieldsProjectUpdates(fieldsProjectUpdates).FieldsProjectBudgets(fieldsProjectBudgets).FieldsPortfolioColumns(fieldsPortfolioColumns).FieldsPortfolioCards(fieldsPortfolioCards).FieldsPortfolioBoards(fieldsPortfolioBoards).FieldsCustomfields(fieldsCustomfields).FieldsCustomfieldProjects(fieldsCustomfieldProjects).FieldsCompanies(fieldsCompanies).Execute()

Generate project report in PDF format



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
    updatedAfter := time.Now() // time.Time | updated after (optional)
    searchTerm := "searchTerm_example" // string | filter by project name (optional)
    reportType := "reportType_example" // string | define the format of the report (optional) (default to "project")
    reportFormat := "reportFormat_example" // string | define the format of the report (optional) (default to "pdf")
    projectType := "projectType_example" // string | filter by project type (optional)
    projectStatuses := "projectStatuses_example" // string | filter by project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "name")
    minLastActivityDate := time.Now() // string | filter by min last activity date (optional)
    maxLastActivityDate := time.Now() // string | filter by max last activity date (optional)
    userId := int32(56) // int32 | filter by user id (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    pageSize := int32(56) // int32 | number of items in a page (not used when generating reports) (optional) (default to 50)
    page := int32(56) // int32 | page number (not used when generating reports) (optional) (default to 1)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    onlyProjectsWithExplicitMembership := true // bool | only show projects with explicit membership (optional) (default to false)
    onlyArchivedProjects := true // bool | return only archived projects (optional) (default to false)
    matchAllProjectTags := true // bool | match all project tags (optional)
    isReportDownload := true // bool | generate a report document (optional) (default to false)
    includeProjectUserInfo := true // bool | fetch user-specific data such as isStarred (optional) (default to false)
    includeCustomFields := true // bool | include custom fields (optional) (default to false)
    includeCompletedStatus := true // bool | optional to include completed projects when filtering by project statuses \"current,late\". (optional)
    includeArchivedProjects := true // bool | include archived projects (optional) (default to false)
    hideObservedProjects := true // bool | hide projects where the logged-in user is just an observer (optional) (default to false)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (not used when generating reports) (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTags := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsProjectcategories := []string{"Inner_example"} // []string |  (optional)
    fieldsProjectUpdates := []string{"Inner_example"} // []string |  (optional)
    fieldsProjectBudgets := []string{"Inner_example"} // []string |  (optional)
    fieldsPortfolioColumns := []string{"Inner_example"} // []string |  (optional)
    fieldsPortfolioCards := []string{"Inner_example"} // []string |  (optional)
    fieldsPortfolioBoards := []string{"Inner_example"} // []string |  (optional)
    fieldsCustomfields := []string{"Inner_example"} // []string |  (optional)
    fieldsCustomfieldProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.GETProjectsApiV3ProjectsPdf(context.Background()).UpdatedAfter(updatedAfter).SearchTerm(searchTerm).ReportType(reportType).ReportFormat(reportFormat).ProjectType(projectType).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).MinLastActivityDate(minLastActivityDate).MaxLastActivityDate(maxLastActivityDate).UserId(userId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).OnlyStarredProjects(onlyStarredProjects).OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership).OnlyArchivedProjects(onlyArchivedProjects).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeProjectUserInfo(includeProjectUserInfo).IncludeCustomFields(includeCustomFields).IncludeCompletedStatus(includeCompletedStatus).IncludeArchivedProjects(includeArchivedProjects).HideObservedProjects(hideObservedProjects).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsProjectcategories(fieldsProjectcategories).FieldsProjectUpdates(fieldsProjectUpdates).FieldsProjectBudgets(fieldsProjectBudgets).FieldsPortfolioColumns(fieldsPortfolioColumns).FieldsPortfolioCards(fieldsPortfolioCards).FieldsPortfolioBoards(fieldsPortfolioBoards).FieldsCustomfields(fieldsCustomfields).FieldsCustomfieldProjects(fieldsCustomfieldProjects).FieldsCompanies(fieldsCompanies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GETProjectsApiV3ProjectsPdf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsPdfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfter** | **time.Time** | updated after | 
 **searchTerm** | **string** | filter by project name | 
 **reportType** | **string** | define the format of the report | [default to &quot;project&quot;]
 **reportFormat** | **string** | define the format of the report | [default to &quot;pdf&quot;]
 **projectType** | **string** | filter by project type | 
 **projectStatuses** | **string** | filter by project status | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;name&quot;]
 **minLastActivityDate** | **string** | filter by min last activity date | 
 **maxLastActivityDate** | **string** | filter by max last activity date | 
 **userId** | **int32** | filter by user id | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **pageSize** | **int32** | number of items in a page (not used when generating reports) | [default to 50]
 **page** | **int32** | page number (not used when generating reports) | [default to 1]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **onlyProjectsWithExplicitMembership** | **bool** | only show projects with explicit membership | [default to false]
 **onlyArchivedProjects** | **bool** | return only archived projects | [default to false]
 **matchAllProjectTags** | **bool** | match all project tags | 
 **isReportDownload** | **bool** | generate a report document | [default to false]
 **includeProjectUserInfo** | **bool** | fetch user-specific data such as isStarred | [default to false]
 **includeCustomFields** | **bool** | include custom fields | [default to false]
 **includeCompletedStatus** | **bool** | optional to include completed projects when filtering by project statuses \&quot;current,late\&quot;. | 
 **includeArchivedProjects** | **bool** | include archived projects | [default to false]
 **hideObservedProjects** | **bool** | hide projects where the logged-in user is just an observer | [default to false]
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include (not used when generating reports) | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTags** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsProjectcategories** | **[]string** |  | 
 **fieldsProjectUpdates** | **[]string** |  | 
 **fieldsProjectBudgets** | **[]string** |  | 
 **fieldsPortfolioColumns** | **[]string** |  | 
 **fieldsPortfolioCards** | **[]string** |  | 
 **fieldsPortfolioBoards** | **[]string** |  | 
 **fieldsCustomfields** | **[]string** |  | 
 **fieldsCustomfieldProjects** | **[]string** |  | 
 **fieldsCompanies** | **[]string** |  | 

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


## GETProjectsApiV3ProjectsTeamworkSamplesJson

> ProjectSampleProjectsResponse GETProjectsApiV3ProjectsTeamworkSamplesJson(ctx).SearchTerm(searchTerm).PageSize(pageSize).Page(page).Ids(ids).FieldsSampleprojects(fieldsSampleprojects).CategoryIds(categoryIds).Execute()

Returns a list of teamwork sample projects



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
    searchTerm := "searchTerm_example" // string | filter by project name (optional)
    pageSize := int32(56) // int32 | number of items in a page (not used when generating reports) (optional) (default to 50)
    page := int32(56) // int32 | page number (not used when generating reports) (optional) (default to 1)
    ids := []int32{int32(123)} // []int32 | filter by sample ids (optional)
    fieldsSampleprojects := []string{"Inner_example"} // []string |  (optional)
    categoryIds := []int32{int32(123)} // []int32 | filter by category ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.GETProjectsApiV3ProjectsTeamworkSamplesJson(context.Background()).SearchTerm(searchTerm).PageSize(pageSize).Page(page).Ids(ids).FieldsSampleprojects(fieldsSampleprojects).CategoryIds(categoryIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GETProjectsApiV3ProjectsTeamworkSamplesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ProjectsTeamworkSamplesJson`: ProjectSampleProjectsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GETProjectsApiV3ProjectsTeamworkSamplesJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsTeamworkSamplesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchTerm** | **string** | filter by project name | 
 **pageSize** | **int32** | number of items in a page (not used when generating reports) | [default to 50]
 **page** | **int32** | page number (not used when generating reports) | [default to 1]
 **ids** | **[]int32** | filter by sample ids | 
 **fieldsSampleprojects** | **[]string** |  | 
 **categoryIds** | **[]int32** | filter by category ids | 

### Return type

[**ProjectSampleProjectsResponse**](ProjectSampleProjectsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3ProjectsTemplatesJson

> ProjectProjectsResponse GETProjectsApiV3ProjectsTemplatesJson(ctx).UpdatedAfter(updatedAfter).SearchTerm(searchTerm).ReportType(reportType).ReportFormat(reportFormat).ProjectType(projectType).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).MinLastActivityDate(minLastActivityDate).MaxLastActivityDate(maxLastActivityDate).UserId(userId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).OnlyStarredProjects(onlyStarredProjects).OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership).OnlyArchivedProjects(onlyArchivedProjects).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeProjectUserInfo(includeProjectUserInfo).IncludeCustomFields(includeCustomFields).IncludeCompletedStatus(includeCompletedStatus).IncludeArchivedProjects(includeArchivedProjects).HideObservedProjects(hideObservedProjects).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsProjectcategories(fieldsProjectcategories).FieldsProjectUpdates(fieldsProjectUpdates).FieldsProjectBudgets(fieldsProjectBudgets).FieldsPortfolioColumns(fieldsPortfolioColumns).FieldsPortfolioCards(fieldsPortfolioCards).FieldsPortfolioBoards(fieldsPortfolioBoards).FieldsCustomfields(fieldsCustomfields).FieldsCustomfieldProjects(fieldsCustomfieldProjects).FieldsCompanies(fieldsCompanies).Execute()

Returns a list of projects templates



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
    updatedAfter := time.Now() // time.Time | updated after (optional)
    searchTerm := "searchTerm_example" // string | filter by project name (optional)
    reportType := "reportType_example" // string | define the format of the report (optional) (default to "project")
    reportFormat := "reportFormat_example" // string | define the format of the report (optional) (default to "pdf")
    projectType := "projectType_example" // string | filter by project type (optional)
    projectStatuses := "projectStatuses_example" // string | filter by project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "name")
    minLastActivityDate := time.Now() // string | filter by min last activity date (optional)
    maxLastActivityDate := time.Now() // string | filter by max last activity date (optional)
    userId := int32(56) // int32 | filter by user id (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    pageSize := int32(56) // int32 | number of items in a page (not used when generating reports) (optional) (default to 50)
    page := int32(56) // int32 | page number (not used when generating reports) (optional) (default to 1)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    onlyProjectsWithExplicitMembership := true // bool | only show projects with explicit membership (optional) (default to false)
    onlyArchivedProjects := true // bool | return only archived projects (optional) (default to false)
    matchAllProjectTags := true // bool | match all project tags (optional)
    isReportDownload := true // bool | generate a report document (optional) (default to false)
    includeProjectUserInfo := true // bool | fetch user-specific data such as isStarred (optional) (default to false)
    includeCustomFields := true // bool | include custom fields (optional) (default to false)
    includeCompletedStatus := true // bool | optional to include completed projects when filtering by project statuses \"current,late\". (optional)
    includeArchivedProjects := true // bool | include archived projects (optional) (default to false)
    hideObservedProjects := true // bool | hide projects where the logged-in user is just an observer (optional) (default to false)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (not used when generating reports) (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTags := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsProjectcategories := []string{"Inner_example"} // []string |  (optional)
    fieldsProjectUpdates := []string{"Inner_example"} // []string |  (optional)
    fieldsProjectBudgets := []string{"Inner_example"} // []string |  (optional)
    fieldsPortfolioColumns := []string{"Inner_example"} // []string |  (optional)
    fieldsPortfolioCards := []string{"Inner_example"} // []string |  (optional)
    fieldsPortfolioBoards := []string{"Inner_example"} // []string |  (optional)
    fieldsCustomfields := []string{"Inner_example"} // []string |  (optional)
    fieldsCustomfieldProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.GETProjectsApiV3ProjectsTemplatesJson(context.Background()).UpdatedAfter(updatedAfter).SearchTerm(searchTerm).ReportType(reportType).ReportFormat(reportFormat).ProjectType(projectType).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).MinLastActivityDate(minLastActivityDate).MaxLastActivityDate(maxLastActivityDate).UserId(userId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).OnlyStarredProjects(onlyStarredProjects).OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership).OnlyArchivedProjects(onlyArchivedProjects).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeProjectUserInfo(includeProjectUserInfo).IncludeCustomFields(includeCustomFields).IncludeCompletedStatus(includeCompletedStatus).IncludeArchivedProjects(includeArchivedProjects).HideObservedProjects(hideObservedProjects).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsProjectcategories(fieldsProjectcategories).FieldsProjectUpdates(fieldsProjectUpdates).FieldsProjectBudgets(fieldsProjectBudgets).FieldsPortfolioColumns(fieldsPortfolioColumns).FieldsPortfolioCards(fieldsPortfolioCards).FieldsPortfolioBoards(fieldsPortfolioBoards).FieldsCustomfields(fieldsCustomfields).FieldsCustomfieldProjects(fieldsCustomfieldProjects).FieldsCompanies(fieldsCompanies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GETProjectsApiV3ProjectsTemplatesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ProjectsTemplatesJson`: ProjectProjectsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GETProjectsApiV3ProjectsTemplatesJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsTemplatesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfter** | **time.Time** | updated after | 
 **searchTerm** | **string** | filter by project name | 
 **reportType** | **string** | define the format of the report | [default to &quot;project&quot;]
 **reportFormat** | **string** | define the format of the report | [default to &quot;pdf&quot;]
 **projectType** | **string** | filter by project type | 
 **projectStatuses** | **string** | filter by project status | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;name&quot;]
 **minLastActivityDate** | **string** | filter by min last activity date | 
 **maxLastActivityDate** | **string** | filter by max last activity date | 
 **userId** | **int32** | filter by user id | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **pageSize** | **int32** | number of items in a page (not used when generating reports) | [default to 50]
 **page** | **int32** | page number (not used when generating reports) | [default to 1]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **onlyProjectsWithExplicitMembership** | **bool** | only show projects with explicit membership | [default to false]
 **onlyArchivedProjects** | **bool** | return only archived projects | [default to false]
 **matchAllProjectTags** | **bool** | match all project tags | 
 **isReportDownload** | **bool** | generate a report document | [default to false]
 **includeProjectUserInfo** | **bool** | fetch user-specific data such as isStarred | [default to false]
 **includeCustomFields** | **bool** | include custom fields | [default to false]
 **includeCompletedStatus** | **bool** | optional to include completed projects when filtering by project statuses \&quot;current,late\&quot;. | 
 **includeArchivedProjects** | **bool** | include archived projects | [default to false]
 **hideObservedProjects** | **bool** | hide projects where the logged-in user is just an observer | [default to false]
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include (not used when generating reports) | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTags** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsProjectcategories** | **[]string** |  | 
 **fieldsProjectUpdates** | **[]string** |  | 
 **fieldsProjectBudgets** | **[]string** |  | 
 **fieldsPortfolioColumns** | **[]string** |  | 
 **fieldsPortfolioCards** | **[]string** |  | 
 **fieldsPortfolioBoards** | **[]string** |  | 
 **fieldsCustomfields** | **[]string** |  | 
 **fieldsCustomfieldProjects** | **[]string** |  | 
 **fieldsCompanies** | **[]string** |  | 

### Return type

[**ProjectProjectsResponse**](ProjectProjectsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3ProjectsXlsx

> GETProjectsApiV3ProjectsXlsx(ctx).UpdatedAfter(updatedAfter).SearchTerm(searchTerm).ReportType(reportType).ReportFormat(reportFormat).ProjectType(projectType).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).MinLastActivityDate(minLastActivityDate).MaxLastActivityDate(maxLastActivityDate).UserId(userId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).OnlyStarredProjects(onlyStarredProjects).OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership).OnlyArchivedProjects(onlyArchivedProjects).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeProjectUserInfo(includeProjectUserInfo).IncludeCustomFields(includeCustomFields).IncludeCompletedStatus(includeCompletedStatus).IncludeArchivedProjects(includeArchivedProjects).HideObservedProjects(hideObservedProjects).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsProjectcategories(fieldsProjectcategories).FieldsProjectUpdates(fieldsProjectUpdates).FieldsProjectBudgets(fieldsProjectBudgets).FieldsPortfolioColumns(fieldsPortfolioColumns).FieldsPortfolioCards(fieldsPortfolioCards).FieldsPortfolioBoards(fieldsPortfolioBoards).FieldsCustomfields(fieldsCustomfields).FieldsCustomfieldProjects(fieldsCustomfieldProjects).FieldsCompanies(fieldsCompanies).Execute()

Generate project report in XLSX format



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
    updatedAfter := time.Now() // time.Time | updated after (optional)
    searchTerm := "searchTerm_example" // string | filter by project name (optional)
    reportType := "reportType_example" // string | define the format of the report (optional) (default to "project")
    reportFormat := "reportFormat_example" // string | define the format of the report (optional) (default to "pdf")
    projectType := "projectType_example" // string | filter by project type (optional)
    projectStatuses := "projectStatuses_example" // string | filter by project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "name")
    minLastActivityDate := time.Now() // string | filter by min last activity date (optional)
    maxLastActivityDate := time.Now() // string | filter by max last activity date (optional)
    userId := int32(56) // int32 | filter by user id (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    pageSize := int32(56) // int32 | number of items in a page (not used when generating reports) (optional) (default to 50)
    page := int32(56) // int32 | page number (not used when generating reports) (optional) (default to 1)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    onlyProjectsWithExplicitMembership := true // bool | only show projects with explicit membership (optional) (default to false)
    onlyArchivedProjects := true // bool | return only archived projects (optional) (default to false)
    matchAllProjectTags := true // bool | match all project tags (optional)
    isReportDownload := true // bool | generate a report document (optional) (default to false)
    includeProjectUserInfo := true // bool | fetch user-specific data such as isStarred (optional) (default to false)
    includeCustomFields := true // bool | include custom fields (optional) (default to false)
    includeCompletedStatus := true // bool | optional to include completed projects when filtering by project statuses \"current,late\". (optional)
    includeArchivedProjects := true // bool | include archived projects (optional) (default to false)
    hideObservedProjects := true // bool | hide projects where the logged-in user is just an observer (optional) (default to false)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (not used when generating reports) (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTags := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsProjectcategories := []string{"Inner_example"} // []string |  (optional)
    fieldsProjectUpdates := []string{"Inner_example"} // []string |  (optional)
    fieldsProjectBudgets := []string{"Inner_example"} // []string |  (optional)
    fieldsPortfolioColumns := []string{"Inner_example"} // []string |  (optional)
    fieldsPortfolioCards := []string{"Inner_example"} // []string |  (optional)
    fieldsPortfolioBoards := []string{"Inner_example"} // []string |  (optional)
    fieldsCustomfields := []string{"Inner_example"} // []string |  (optional)
    fieldsCustomfieldProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.GETProjectsApiV3ProjectsXlsx(context.Background()).UpdatedAfter(updatedAfter).SearchTerm(searchTerm).ReportType(reportType).ReportFormat(reportFormat).ProjectType(projectType).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).MinLastActivityDate(minLastActivityDate).MaxLastActivityDate(maxLastActivityDate).UserId(userId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).OnlyStarredProjects(onlyStarredProjects).OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership).OnlyArchivedProjects(onlyArchivedProjects).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeProjectUserInfo(includeProjectUserInfo).IncludeCustomFields(includeCustomFields).IncludeCompletedStatus(includeCompletedStatus).IncludeArchivedProjects(includeArchivedProjects).HideObservedProjects(hideObservedProjects).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsProjectcategories(fieldsProjectcategories).FieldsProjectUpdates(fieldsProjectUpdates).FieldsProjectBudgets(fieldsProjectBudgets).FieldsPortfolioColumns(fieldsPortfolioColumns).FieldsPortfolioCards(fieldsPortfolioCards).FieldsPortfolioBoards(fieldsPortfolioBoards).FieldsCustomfields(fieldsCustomfields).FieldsCustomfieldProjects(fieldsCustomfieldProjects).FieldsCompanies(fieldsCompanies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GETProjectsApiV3ProjectsXlsx``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsXlsxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfter** | **time.Time** | updated after | 
 **searchTerm** | **string** | filter by project name | 
 **reportType** | **string** | define the format of the report | [default to &quot;project&quot;]
 **reportFormat** | **string** | define the format of the report | [default to &quot;pdf&quot;]
 **projectType** | **string** | filter by project type | 
 **projectStatuses** | **string** | filter by project status | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;name&quot;]
 **minLastActivityDate** | **string** | filter by min last activity date | 
 **maxLastActivityDate** | **string** | filter by max last activity date | 
 **userId** | **int32** | filter by user id | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **pageSize** | **int32** | number of items in a page (not used when generating reports) | [default to 50]
 **page** | **int32** | page number (not used when generating reports) | [default to 1]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **onlyProjectsWithExplicitMembership** | **bool** | only show projects with explicit membership | [default to false]
 **onlyArchivedProjects** | **bool** | return only archived projects | [default to false]
 **matchAllProjectTags** | **bool** | match all project tags | 
 **isReportDownload** | **bool** | generate a report document | [default to false]
 **includeProjectUserInfo** | **bool** | fetch user-specific data such as isStarred | [default to false]
 **includeCustomFields** | **bool** | include custom fields | [default to false]
 **includeCompletedStatus** | **bool** | optional to include completed projects when filtering by project statuses \&quot;current,late\&quot;. | 
 **includeArchivedProjects** | **bool** | include archived projects | [default to false]
 **hideObservedProjects** | **bool** | hide projects where the logged-in user is just an observer | [default to false]
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include (not used when generating reports) | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTags** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsProjectcategories** | **[]string** |  | 
 **fieldsProjectUpdates** | **[]string** |  | 
 **fieldsProjectBudgets** | **[]string** |  | 
 **fieldsPortfolioColumns** | **[]string** |  | 
 **fieldsPortfolioCards** | **[]string** |  | 
 **fieldsPortfolioBoards** | **[]string** |  | 
 **fieldsCustomfields** | **[]string** |  | 
 **fieldsCustomfieldProjects** | **[]string** |  | 
 **fieldsCompanies** | **[]string** |  | 

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


## GETProjectsApiV3ProjectsprojectIdFeatureorderJson

> ProjectFeatureOrderResponse GETProjectsApiV3ProjectsprojectIdFeatureorderJson(ctx, projectId).Execute()

Returns a project features order to display in tab

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
    projectId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.GETProjectsApiV3ProjectsprojectIdFeatureorderJson(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GETProjectsApiV3ProjectsprojectIdFeatureorderJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ProjectsprojectIdFeatureorderJson`: ProjectFeatureOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GETProjectsApiV3ProjectsprojectIdFeatureorderJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsprojectIdFeatureorderJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectFeatureOrderResponse**](ProjectFeatureOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3ProjectsprojectIdJson

> ProjectSingleResponse GETProjectsApiV3ProjectsprojectIdJson(ctx, projectId).UpdatedAfter(updatedAfter).SearchTerm(searchTerm).ReportType(reportType).ReportFormat(reportFormat).ProjectType(projectType).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).MinLastActivityDate(minLastActivityDate).MaxLastActivityDate(maxLastActivityDate).UserId(userId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).OnlyStarredProjects(onlyStarredProjects).OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership).OnlyArchivedProjects(onlyArchivedProjects).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeProjectUserInfo(includeProjectUserInfo).IncludeCustomFields(includeCustomFields).IncludeCompletedStatus(includeCompletedStatus).IncludeArchivedProjects(includeArchivedProjects).HideObservedProjects(hideObservedProjects).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsProjectcategories(fieldsProjectcategories).FieldsProjectUpdates(fieldsProjectUpdates).FieldsProjectBudgets(fieldsProjectBudgets).FieldsPortfolioColumns(fieldsPortfolioColumns).FieldsPortfolioCards(fieldsPortfolioCards).FieldsPortfolioBoards(fieldsPortfolioBoards).FieldsCustomfields(fieldsCustomfields).FieldsCustomfieldProjects(fieldsCustomfieldProjects).FieldsCompanies(fieldsCompanies).Execute()

Returns a project



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
    updatedAfter := time.Now() // time.Time | updated after (optional)
    searchTerm := "searchTerm_example" // string | filter by project name (optional)
    reportType := "reportType_example" // string | define the format of the report (optional) (default to "project")
    reportFormat := "reportFormat_example" // string | define the format of the report (optional) (default to "pdf")
    projectType := "projectType_example" // string | filter by project type (optional)
    projectStatuses := "projectStatuses_example" // string | filter by project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "name")
    minLastActivityDate := time.Now() // string | filter by min last activity date (optional)
    maxLastActivityDate := time.Now() // string | filter by max last activity date (optional)
    userId := int32(56) // int32 | filter by user id (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    pageSize := int32(56) // int32 | number of items in a page (not used when generating reports) (optional) (default to 50)
    page := int32(56) // int32 | page number (not used when generating reports) (optional) (default to 1)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    onlyProjectsWithExplicitMembership := true // bool | only show projects with explicit membership (optional) (default to false)
    onlyArchivedProjects := true // bool | return only archived projects (optional) (default to false)
    matchAllProjectTags := true // bool | match all project tags (optional)
    isReportDownload := true // bool | generate a report document (optional) (default to false)
    includeProjectUserInfo := true // bool | fetch user-specific data such as isStarred (optional) (default to false)
    includeCustomFields := true // bool | include custom fields (optional) (default to false)
    includeCompletedStatus := true // bool | optional to include completed projects when filtering by project statuses \"current,late\". (optional)
    includeArchivedProjects := true // bool | include archived projects (optional) (default to false)
    hideObservedProjects := true // bool | hide projects where the logged-in user is just an observer (optional) (default to false)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (not used when generating reports) (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTags := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsProjectcategories := []string{"Inner_example"} // []string |  (optional)
    fieldsProjectUpdates := []string{"Inner_example"} // []string |  (optional)
    fieldsProjectBudgets := []string{"Inner_example"} // []string |  (optional)
    fieldsPortfolioColumns := []string{"Inner_example"} // []string |  (optional)
    fieldsPortfolioCards := []string{"Inner_example"} // []string |  (optional)
    fieldsPortfolioBoards := []string{"Inner_example"} // []string |  (optional)
    fieldsCustomfields := []string{"Inner_example"} // []string |  (optional)
    fieldsCustomfieldProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.GETProjectsApiV3ProjectsprojectIdJson(context.Background(), projectId).UpdatedAfter(updatedAfter).SearchTerm(searchTerm).ReportType(reportType).ReportFormat(reportFormat).ProjectType(projectType).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).MinLastActivityDate(minLastActivityDate).MaxLastActivityDate(maxLastActivityDate).UserId(userId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).OnlyStarredProjects(onlyStarredProjects).OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership).OnlyArchivedProjects(onlyArchivedProjects).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeProjectUserInfo(includeProjectUserInfo).IncludeCustomFields(includeCustomFields).IncludeCompletedStatus(includeCompletedStatus).IncludeArchivedProjects(includeArchivedProjects).HideObservedProjects(hideObservedProjects).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsProjectcategories(fieldsProjectcategories).FieldsProjectUpdates(fieldsProjectUpdates).FieldsProjectBudgets(fieldsProjectBudgets).FieldsPortfolioColumns(fieldsPortfolioColumns).FieldsPortfolioCards(fieldsPortfolioCards).FieldsPortfolioBoards(fieldsPortfolioBoards).FieldsCustomfields(fieldsCustomfields).FieldsCustomfieldProjects(fieldsCustomfieldProjects).FieldsCompanies(fieldsCompanies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GETProjectsApiV3ProjectsprojectIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ProjectsprojectIdJson`: ProjectSingleResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GETProjectsApiV3ProjectsprojectIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsprojectIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedAfter** | **time.Time** | updated after | 
 **searchTerm** | **string** | filter by project name | 
 **reportType** | **string** | define the format of the report | [default to &quot;project&quot;]
 **reportFormat** | **string** | define the format of the report | [default to &quot;pdf&quot;]
 **projectType** | **string** | filter by project type | 
 **projectStatuses** | **string** | filter by project status | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;name&quot;]
 **minLastActivityDate** | **string** | filter by min last activity date | 
 **maxLastActivityDate** | **string** | filter by max last activity date | 
 **userId** | **int32** | filter by user id | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **pageSize** | **int32** | number of items in a page (not used when generating reports) | [default to 50]
 **page** | **int32** | page number (not used when generating reports) | [default to 1]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **onlyProjectsWithExplicitMembership** | **bool** | only show projects with explicit membership | [default to false]
 **onlyArchivedProjects** | **bool** | return only archived projects | [default to false]
 **matchAllProjectTags** | **bool** | match all project tags | 
 **isReportDownload** | **bool** | generate a report document | [default to false]
 **includeProjectUserInfo** | **bool** | fetch user-specific data such as isStarred | [default to false]
 **includeCustomFields** | **bool** | include custom fields | [default to false]
 **includeCompletedStatus** | **bool** | optional to include completed projects when filtering by project statuses \&quot;current,late\&quot;. | 
 **includeArchivedProjects** | **bool** | include archived projects | [default to false]
 **hideObservedProjects** | **bool** | hide projects where the logged-in user is just an observer | [default to false]
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include (not used when generating reports) | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTags** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsProjectcategories** | **[]string** |  | 
 **fieldsProjectUpdates** | **[]string** |  | 
 **fieldsProjectBudgets** | **[]string** |  | 
 **fieldsPortfolioColumns** | **[]string** |  | 
 **fieldsPortfolioCards** | **[]string** |  | 
 **fieldsPortfolioBoards** | **[]string** |  | 
 **fieldsCustomfields** | **[]string** |  | 
 **fieldsCustomfieldProjects** | **[]string** |  | 
 **fieldsCompanies** | **[]string** |  | 

### Return type

[**ProjectSingleResponse**](ProjectSingleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV3ProjectsFeatureorderJson

> ProjectFeatureOrderResponse PUTProjectsApiV3ProjectsFeatureorderJson(ctx).ProjectFeatureOrderDefaults(projectFeatureOrderDefaults).Execute()

Sets the default features order to display in tab



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
    projectFeatureOrderDefaults := *openapiclient.NewProjectFeatureOrderDefaults() // ProjectFeatureOrderDefaults | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.PUTProjectsApiV3ProjectsFeatureorderJson(context.Background()).ProjectFeatureOrderDefaults(projectFeatureOrderDefaults).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.PUTProjectsApiV3ProjectsFeatureorderJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsApiV3ProjectsFeatureorderJson`: ProjectFeatureOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.PUTProjectsApiV3ProjectsFeatureorderJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3ProjectsFeatureorderJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectFeatureOrderDefaults** | [**ProjectFeatureOrderDefaults**](ProjectFeatureOrderDefaults.md) |  | 

### Return type

[**ProjectFeatureOrderResponse**](ProjectFeatureOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV3ProjectsprojectIdFeatureorderJson

> ProjectFeatureOrderResponse PUTProjectsApiV3ProjectsprojectIdFeatureorderJson(ctx, projectId).ProjectFeatureOrder(projectFeatureOrder).Execute()

Sets the the features order to display in tab



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
    projectId := int32(56) // int32 | 
    projectFeatureOrder := *openapiclient.NewProjectFeatureOrder() // ProjectFeatureOrder | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.PUTProjectsApiV3ProjectsprojectIdFeatureorderJson(context.Background(), projectId).ProjectFeatureOrder(projectFeatureOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.PUTProjectsApiV3ProjectsprojectIdFeatureorderJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsApiV3ProjectsprojectIdFeatureorderJson`: ProjectFeatureOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.PUTProjectsApiV3ProjectsprojectIdFeatureorderJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3ProjectsprojectIdFeatureorderJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectFeatureOrder** | [**ProjectFeatureOrder**](ProjectFeatureOrder.md) |  | 

### Return type

[**ProjectFeatureOrderResponse**](ProjectFeatureOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

