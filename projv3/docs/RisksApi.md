# \RisksApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3ProjectsprojectIdRisks**](RisksApi.md#GETProjectsApiV3ProjectsprojectIdRisks) | **Get** /projects/api/v3/projects/{projectId}/risks | Get risks for a specific project
[**GETProjectsApiV3RisksJson**](RisksApi.md#GETProjectsApiV3RisksJson) | **Get** /projects/api/v3/risks.json | Get all risks



## GETProjectsApiV3ProjectsprojectIdRisks

> RiskRisksResponse GETProjectsApiV3ProjectsprojectIdRisks(ctx, projectId).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).Statuses(statuses).SortOrder(sortOrder).Sort(sort).SearchTerm(searchTerm).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).ProjectHealths(projectHealths).ProbabilityTo(probabilityTo).ProbabilityFrom(probabilityFrom).PageSize(pageSize).Page(page).ImpactTo(impactTo).ImpactFrom(impactFrom).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeArchivedProjects(includeArchivedProjects).ImpactSchedule(impactSchedule).ImpactPerformance(impactPerformance).ImpactCost(impactCost).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsRisks(fieldsRisks).FieldsProjects(fieldsProjects).FieldsCompanies(fieldsCompanies).Execute()

Get risks for a specific project



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
    updatedAfterDate := time.Now() // time.Time | filter by risks updated after specified date (deprecated, use updatedAfter) (optional)
    updatedAfter := time.Now() // time.Time | filter by risks updated after specified date (optional)
    statuses := "statuses_example" // string | filter by risk statuses (optional)
    sortOrder := "sortOrder_example" // string | sort order (deprecated, use orderMode) (optional)
    sort := "sort_example" // string | sort by (deprecated, use orderBy) (optional)
    searchTerm := "searchTerm_example" // string | search term (optional)
    reportFormat := "reportFormat_example" // string | define the format of the report default: pdf (optional)
    projectStatuses := "projectStatuses_example" // string | list of project statuses (optional)
    orderMode := "orderMode_example" // string | sort order (optional)
    orderBy := "orderBy_example" // string | sort by (optional)
    projectHealths := int32(56) // int32 | list of project health (optional)
    probabilityTo := int32(56) // int32 | filter by probability (optional)
    probabilityFrom := int32(56) // int32 | filter by probability (optional)
    pageSize := int32(56) // int32 | number of items in a page default: 50 (optional)
    page := int32(56) // int32 | page number default: 1 (optional)
    impactTo := int32(56) // int32 | filter by impact (optional)
    impactFrom := int32(56) // int32 | filter by impact (optional)
    showDeleted := true // bool | include deleted items default: false (optional)
    onlyStarredProjects := true // bool | filter by starred projects only default: false (optional)
    matchAllProjectTags := true // bool | enforce all tag ids must be matched (optional)
    isReportDownload := true // bool | generate a report document (optional)
    includeArchivedProjects := true // bool | include archived projects (optional)
    impactSchedule := true // bool | filter by risks that impact schedule (optional)
    impactPerformance := true // bool | filter by risks that impact performance (optional)
    impactCost := true // bool | filter by risks that impact cost (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsRisks := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RisksApi.GETProjectsApiV3ProjectsprojectIdRisks(context.Background(), projectId).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).Statuses(statuses).SortOrder(sortOrder).Sort(sort).SearchTerm(searchTerm).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).ProjectHealths(projectHealths).ProbabilityTo(probabilityTo).ProbabilityFrom(probabilityFrom).PageSize(pageSize).Page(page).ImpactTo(impactTo).ImpactFrom(impactFrom).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeArchivedProjects(includeArchivedProjects).ImpactSchedule(impactSchedule).ImpactPerformance(impactPerformance).ImpactCost(impactCost).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsRisks(fieldsRisks).FieldsProjects(fieldsProjects).FieldsCompanies(fieldsCompanies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RisksApi.GETProjectsApiV3ProjectsprojectIdRisks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ProjectsprojectIdRisks`: RiskRisksResponse
    fmt.Fprintf(os.Stdout, "Response from `RisksApi.GETProjectsApiV3ProjectsprojectIdRisks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsprojectIdRisksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedAfterDate** | **time.Time** | filter by risks updated after specified date (deprecated, use updatedAfter) | 
 **updatedAfter** | **time.Time** | filter by risks updated after specified date | 
 **statuses** | **string** | filter by risk statuses | 
 **sortOrder** | **string** | sort order (deprecated, use orderMode) | 
 **sort** | **string** | sort by (deprecated, use orderBy) | 
 **searchTerm** | **string** | search term | 
 **reportFormat** | **string** | define the format of the report default: pdf | 
 **projectStatuses** | **string** | list of project statuses | 
 **orderMode** | **string** | sort order | 
 **orderBy** | **string** | sort by | 
 **projectHealths** | **int32** | list of project health | 
 **probabilityTo** | **int32** | filter by probability | 
 **probabilityFrom** | **int32** | filter by probability | 
 **pageSize** | **int32** | number of items in a page default: 50 | 
 **page** | **int32** | page number default: 1 | 
 **impactTo** | **int32** | filter by impact | 
 **impactFrom** | **int32** | filter by impact | 
 **showDeleted** | **bool** | include deleted items default: false | 
 **onlyStarredProjects** | **bool** | filter by starred projects only default: false | 
 **matchAllProjectTags** | **bool** | enforce all tag ids must be matched | 
 **isReportDownload** | **bool** | generate a report document | 
 **includeArchivedProjects** | **bool** | include archived projects | 
 **impactSchedule** | **bool** | filter by risks that impact schedule | 
 **impactPerformance** | **bool** | filter by risks that impact performance | 
 **impactCost** | **bool** | filter by risks that impact cost | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsRisks** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsCompanies** | **[]string** |  | 

### Return type

[**RiskRisksResponse**](RiskRisksResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3RisksJson

> RiskRisksResponse GETProjectsApiV3RisksJson(ctx).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).Statuses(statuses).SortOrder(sortOrder).Sort(sort).SearchTerm(searchTerm).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).ProjectHealths(projectHealths).ProbabilityTo(probabilityTo).ProbabilityFrom(probabilityFrom).PageSize(pageSize).Page(page).ImpactTo(impactTo).ImpactFrom(impactFrom).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeArchivedProjects(includeArchivedProjects).ImpactSchedule(impactSchedule).ImpactPerformance(impactPerformance).ImpactCost(impactCost).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsRisks(fieldsRisks).FieldsProjects(fieldsProjects).FieldsCompanies(fieldsCompanies).Execute()

Get all risks



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
    updatedAfterDate := time.Now() // time.Time | filter by risks updated after specified date (deprecated, use updatedAfter) (optional)
    updatedAfter := time.Now() // time.Time | filter by risks updated after specified date (optional)
    statuses := "statuses_example" // string | filter by risk statuses (optional)
    sortOrder := "sortOrder_example" // string | sort order (deprecated, use orderMode) (optional)
    sort := "sort_example" // string | sort by (deprecated, use orderBy) (optional)
    searchTerm := "searchTerm_example" // string | search term (optional)
    reportFormat := "reportFormat_example" // string | define the format of the report default: pdf (optional)
    projectStatuses := "projectStatuses_example" // string | list of project statuses (optional)
    orderMode := "orderMode_example" // string | sort order (optional)
    orderBy := "orderBy_example" // string | sort by (optional)
    projectHealths := int32(56) // int32 | list of project health (optional)
    probabilityTo := int32(56) // int32 | filter by probability (optional)
    probabilityFrom := int32(56) // int32 | filter by probability (optional)
    pageSize := int32(56) // int32 | number of items in a page default: 50 (optional)
    page := int32(56) // int32 | page number default: 1 (optional)
    impactTo := int32(56) // int32 | filter by impact (optional)
    impactFrom := int32(56) // int32 | filter by impact (optional)
    showDeleted := true // bool | include deleted items default: false (optional)
    onlyStarredProjects := true // bool | filter by starred projects only default: false (optional)
    matchAllProjectTags := true // bool | enforce all tag ids must be matched (optional)
    isReportDownload := true // bool | generate a report document (optional)
    includeArchivedProjects := true // bool | include archived projects (optional)
    impactSchedule := true // bool | filter by risks that impact schedule (optional)
    impactPerformance := true // bool | filter by risks that impact performance (optional)
    impactCost := true // bool | filter by risks that impact cost (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsRisks := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RisksApi.GETProjectsApiV3RisksJson(context.Background()).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).Statuses(statuses).SortOrder(sortOrder).Sort(sort).SearchTerm(searchTerm).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).ProjectHealths(projectHealths).ProbabilityTo(probabilityTo).ProbabilityFrom(probabilityFrom).PageSize(pageSize).Page(page).ImpactTo(impactTo).ImpactFrom(impactFrom).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeArchivedProjects(includeArchivedProjects).ImpactSchedule(impactSchedule).ImpactPerformance(impactPerformance).ImpactCost(impactCost).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsRisks(fieldsRisks).FieldsProjects(fieldsProjects).FieldsCompanies(fieldsCompanies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RisksApi.GETProjectsApiV3RisksJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3RisksJson`: RiskRisksResponse
    fmt.Fprintf(os.Stdout, "Response from `RisksApi.GETProjectsApiV3RisksJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3RisksJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfterDate** | **time.Time** | filter by risks updated after specified date (deprecated, use updatedAfter) | 
 **updatedAfter** | **time.Time** | filter by risks updated after specified date | 
 **statuses** | **string** | filter by risk statuses | 
 **sortOrder** | **string** | sort order (deprecated, use orderMode) | 
 **sort** | **string** | sort by (deprecated, use orderBy) | 
 **searchTerm** | **string** | search term | 
 **reportFormat** | **string** | define the format of the report default: pdf | 
 **projectStatuses** | **string** | list of project statuses | 
 **orderMode** | **string** | sort order | 
 **orderBy** | **string** | sort by | 
 **projectHealths** | **int32** | list of project health | 
 **probabilityTo** | **int32** | filter by probability | 
 **probabilityFrom** | **int32** | filter by probability | 
 **pageSize** | **int32** | number of items in a page default: 50 | 
 **page** | **int32** | page number default: 1 | 
 **impactTo** | **int32** | filter by impact | 
 **impactFrom** | **int32** | filter by impact | 
 **showDeleted** | **bool** | include deleted items default: false | 
 **onlyStarredProjects** | **bool** | filter by starred projects only default: false | 
 **matchAllProjectTags** | **bool** | enforce all tag ids must be matched | 
 **isReportDownload** | **bool** | generate a report document | 
 **includeArchivedProjects** | **bool** | include archived projects | 
 **impactSchedule** | **bool** | filter by risks that impact schedule | 
 **impactPerformance** | **bool** | filter by risks that impact performance | 
 **impactCost** | **bool** | filter by risks that impact cost | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsRisks** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsCompanies** | **[]string** |  | 

### Return type

[**RiskRisksResponse**](RiskRisksResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

