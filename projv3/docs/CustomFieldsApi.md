# \CustomFieldsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEProjectsApiV3CustomfieldscustomFieldIdJson**](CustomFieldsApi.md#DELETEProjectsApiV3CustomfieldscustomFieldIdJson) | **Delete** /projects/api/v3/customfields/{customFieldId}.json | Delete an existing custom field
[**DELETEProjectsApiV3ProjectsprojectIdCustomfieldscustomFieldIdJson**](CustomFieldsApi.md#DELETEProjectsApiV3ProjectsprojectIdCustomfieldscustomFieldIdJson) | **Delete** /projects/api/v3/projects/{projectId}/customfields/{customFieldId}.json | Delete an existing project custom field value.
[**DELETEProjectsApiV3TaskstaskIdCustomfieldscustomFieldIdJson**](CustomFieldsApi.md#DELETEProjectsApiV3TaskstaskIdCustomfieldscustomFieldIdJson) | **Delete** /projects/api/v3/tasks/{taskId}/customfields/{customFieldId}.json | Delete an existing task custom field value.
[**GETProjectsApiV3CustomfieldsJson**](CustomFieldsApi.md#GETProjectsApiV3CustomfieldsJson) | **Get** /projects/api/v3/customfields.json | Get all custom fields
[**GETProjectsApiV3CustomfieldscustomFieldIdJson**](CustomFieldsApi.md#GETProjectsApiV3CustomfieldscustomFieldIdJson) | **Get** /projects/api/v3/customfields/{customFieldId}.json | Get a custom field by id.
[**GETProjectsApiV3ProjectsprojectIdCustomfieldsJson**](CustomFieldsApi.md#GETProjectsApiV3ProjectsprojectIdCustomfieldsJson) | **Get** /projects/api/v3/projects/{projectId}/customfields.json | Project custom field values.
[**GETProjectsApiV3ProjectsprojectIdCustomfieldscustomFieldIdJson**](CustomFieldsApi.md#GETProjectsApiV3ProjectsprojectIdCustomfieldscustomFieldIdJson) | **Get** /projects/api/v3/projects/{projectId}/customfields/{customFieldId}.json | Project custom field value.
[**GETProjectsApiV3TaskstaskIdCustomfieldsJson**](CustomFieldsApi.md#GETProjectsApiV3TaskstaskIdCustomfieldsJson) | **Get** /projects/api/v3/tasks/{taskId}/customfields.json | Task custom field values.
[**GETProjectsApiV3TaskstaskIdCustomfieldscustomFieldIdJson**](CustomFieldsApi.md#GETProjectsApiV3TaskstaskIdCustomfieldscustomFieldIdJson) | **Get** /projects/api/v3/tasks/{taskId}/customfields/{customFieldId}.json | Task custom field value.
[**PATCHProjectsApiV3CustomfieldscustomFieldIdJson**](CustomFieldsApi.md#PATCHProjectsApiV3CustomfieldscustomFieldIdJson) | **Patch** /projects/api/v3/customfields/{customFieldId}.json | Update an existing custom field.
[**PATCHProjectsApiV3ProjectsprojectIdCustomfieldscustomFieldIdJson**](CustomFieldsApi.md#PATCHProjectsApiV3ProjectsprojectIdCustomfieldscustomFieldIdJson) | **Patch** /projects/api/v3/projects/{projectId}/customfields/{customFieldId}.json | Update an existing project custom field value.
[**PATCHProjectsApiV3TaskstaskIdCustomfieldscustomFieldIdJson**](CustomFieldsApi.md#PATCHProjectsApiV3TaskstaskIdCustomfieldscustomFieldIdJson) | **Patch** /projects/api/v3/tasks/{taskId}/customfields/{customFieldId}.json | Update an existing task custom field value.
[**POSTProjectsApiV3CustomfieldsBulkDeleteJson**](CustomFieldsApi.md#POSTProjectsApiV3CustomfieldsBulkDeleteJson) | **Post** /projects/api/v3/customfields/bulk/delete.json | Delete many custom fields at once
[**POSTProjectsApiV3CustomfieldsJson**](CustomFieldsApi.md#POSTProjectsApiV3CustomfieldsJson) | **Post** /projects/api/v3/customfields.json | Create a new custom field
[**POSTProjectsApiV3ProjectsprojectIdCustomfieldsBulkDeleteJson**](CustomFieldsApi.md#POSTProjectsApiV3ProjectsprojectIdCustomfieldsBulkDeleteJson) | **Post** /projects/api/v3/projects/{projectId}/customfields/bulk/delete.json | Delete many project custom fields values at once.
[**POSTProjectsApiV3ProjectsprojectIdCustomfieldsBulkUpdateJson**](CustomFieldsApi.md#POSTProjectsApiV3ProjectsprojectIdCustomfieldsBulkUpdateJson) | **Post** /projects/api/v3/projects/{projectId}/customfields/bulk/update.json | Update many project custom field values at once.
[**POSTProjectsApiV3ProjectsprojectIdCustomfieldsJson**](CustomFieldsApi.md#POSTProjectsApiV3ProjectsprojectIdCustomfieldsJson) | **Post** /projects/api/v3/projects/{projectId}/customfields.json | Add project custom field value.
[**POSTProjectsApiV3TaskstaskIdCustomfieldsBulkDeleteJson**](CustomFieldsApi.md#POSTProjectsApiV3TaskstaskIdCustomfieldsBulkDeleteJson) | **Post** /projects/api/v3/tasks/{taskId}/customfields/bulk/delete.json | Delete many task custom fields values at once.
[**POSTProjectsApiV3TaskstaskIdCustomfieldsBulkUpdateJson**](CustomFieldsApi.md#POSTProjectsApiV3TaskstaskIdCustomfieldsBulkUpdateJson) | **Post** /projects/api/v3/tasks/{taskId}/customfields/bulk/update.json | Update many task custom field values at once.
[**POSTProjectsApiV3TaskstaskIdCustomfieldsJson**](CustomFieldsApi.md#POSTProjectsApiV3TaskstaskIdCustomfieldsJson) | **Post** /projects/api/v3/tasks/{taskId}/customfields.json | Add task custom field value.
[**PUTProjectsApiV3CustomfieldscustomFieldIdJson**](CustomFieldsApi.md#PUTProjectsApiV3CustomfieldscustomFieldIdJson) | **Put** /projects/api/v3/customfields/{customFieldId}.json | Update an existing custom field.



## DELETEProjectsApiV3CustomfieldscustomFieldIdJson

> DELETEProjectsApiV3CustomfieldscustomFieldIdJson(ctx, customFieldId).Execute()

Delete an existing custom field



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
    customFieldId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.DELETEProjectsApiV3CustomfieldscustomFieldIdJson(context.Background(), customFieldId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.DELETEProjectsApiV3CustomfieldscustomFieldIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customFieldId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEProjectsApiV3CustomfieldscustomFieldIdJsonRequest struct via the builder pattern


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


## DELETEProjectsApiV3ProjectsprojectIdCustomfieldscustomFieldIdJson

> DELETEProjectsApiV3ProjectsprojectIdCustomfieldscustomFieldIdJson(ctx, projectId, customFieldId).Execute()

Delete an existing project custom field value.



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
    customFieldId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.DELETEProjectsApiV3ProjectsprojectIdCustomfieldscustomFieldIdJson(context.Background(), projectId, customFieldId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.DELETEProjectsApiV3ProjectsprojectIdCustomfieldscustomFieldIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**customFieldId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEProjectsApiV3ProjectsprojectIdCustomfieldscustomFieldIdJsonRequest struct via the builder pattern


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


## DELETEProjectsApiV3TaskstaskIdCustomfieldscustomFieldIdJson

> DELETEProjectsApiV3TaskstaskIdCustomfieldscustomFieldIdJson(ctx, taskId, customFieldId).Execute()

Delete an existing task custom field value.



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
    taskId := int32(56) // int32 | 
    customFieldId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.DELETEProjectsApiV3TaskstaskIdCustomfieldscustomFieldIdJson(context.Background(), taskId, customFieldId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.DELETEProjectsApiV3TaskstaskIdCustomfieldscustomFieldIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **int32** |  | 
**customFieldId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEProjectsApiV3TaskstaskIdCustomfieldscustomFieldIdJsonRequest struct via the builder pattern


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


## GETProjectsApiV3CustomfieldsJson

> CustomfieldCustomFieldsResponse GETProjectsApiV3CustomfieldsJson(ctx).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).Entities(entities).ProjectId(projectId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).OnlySiteLevel(onlySiteLevel).MatchAllProjectTags(matchAllProjectTags).IncludeSiteLevel(includeSiteLevel).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsProjects(fieldsProjects).Execute()

Get all custom fields



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
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "name")
    entities := "entities_example" // string | filter by entities (optional)
    projectId := int32(56) // int32 | filter by project id (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    showDeleted := true // bool | include deleted items (optional) (default to false)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    onlySiteLevel := true // bool | only return site-level custom fields (optional) (default to false)
    matchAllProjectTags := true // bool | match all project tags (optional)
    includeSiteLevel := true // bool | include custom fields for site-level in the response (optional) (default to true)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.GETProjectsApiV3CustomfieldsJson(context.Background()).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).Entities(entities).ProjectId(projectId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).OnlySiteLevel(onlySiteLevel).MatchAllProjectTags(matchAllProjectTags).IncludeSiteLevel(includeSiteLevel).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsProjects(fieldsProjects).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.GETProjectsApiV3CustomfieldsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3CustomfieldsJson`: CustomfieldCustomFieldsResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFieldsApi.GETProjectsApiV3CustomfieldsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3CustomfieldsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectStatuses** | **string** | filter by project statuses | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;name&quot;]
 **entities** | **string** | filter by entities | 
 **projectId** | **int32** | filter by project id | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **showDeleted** | **bool** | include deleted items | [default to false]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **onlySiteLevel** | **bool** | only return site-level custom fields | [default to false]
 **matchAllProjectTags** | **bool** | match all project tags | 
 **includeSiteLevel** | **bool** | include custom fields for site-level in the response | [default to true]
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 

### Return type

[**CustomfieldCustomFieldsResponse**](CustomfieldCustomFieldsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3CustomfieldscustomFieldIdJson

> CustomfieldResponse GETProjectsApiV3CustomfieldscustomFieldIdJson(ctx, customFieldId).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).Entities(entities).ProjectId(projectId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).OnlySiteLevel(onlySiteLevel).MatchAllProjectTags(matchAllProjectTags).IncludeSiteLevel(includeSiteLevel).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsProjects(fieldsProjects).Execute()

Get a custom field by id.



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
    customFieldId := int32(56) // int32 | 
    projectStatuses := "projectStatuses_example" // string | filter by project statuses (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "name")
    entities := "entities_example" // string | filter by entities (optional)
    projectId := int32(56) // int32 | filter by project id (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    showDeleted := true // bool | include deleted items (optional) (default to false)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    onlySiteLevel := true // bool | only return site-level custom fields (optional) (default to false)
    matchAllProjectTags := true // bool | match all project tags (optional)
    includeSiteLevel := true // bool | include custom fields for site-level in the response (optional) (default to true)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.GETProjectsApiV3CustomfieldscustomFieldIdJson(context.Background(), customFieldId).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).Entities(entities).ProjectId(projectId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).OnlySiteLevel(onlySiteLevel).MatchAllProjectTags(matchAllProjectTags).IncludeSiteLevel(includeSiteLevel).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsProjects(fieldsProjects).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.GETProjectsApiV3CustomfieldscustomFieldIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3CustomfieldscustomFieldIdJson`: CustomfieldResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFieldsApi.GETProjectsApiV3CustomfieldscustomFieldIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customFieldId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3CustomfieldscustomFieldIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectStatuses** | **string** | filter by project statuses | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;name&quot;]
 **entities** | **string** | filter by entities | 
 **projectId** | **int32** | filter by project id | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **showDeleted** | **bool** | include deleted items | [default to false]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **onlySiteLevel** | **bool** | only return site-level custom fields | [default to false]
 **matchAllProjectTags** | **bool** | match all project tags | 
 **includeSiteLevel** | **bool** | include custom fields for site-level in the response | [default to true]
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 

### Return type

[**CustomfieldResponse**](CustomfieldResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3ProjectsprojectIdCustomfieldsJson

> ProjectCustomFieldProjectsResponse GETProjectsApiV3ProjectsprojectIdCustomfieldsJson(ctx, projectId).UpdatedAfter(updatedAfter).SearchTerm(searchTerm).ReportType(reportType).ReportFormat(reportFormat).ProjectType(projectType).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).MinLastActivityDate(minLastActivityDate).MaxLastActivityDate(maxLastActivityDate).UserId(userId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).OnlyStarredProjects(onlyStarredProjects).OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership).OnlyArchivedProjects(onlyArchivedProjects).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeProjectUserInfo(includeProjectUserInfo).IncludeCustomFields(includeCustomFields).IncludeCompletedStatus(includeCompletedStatus).IncludeArchivedProjects(includeArchivedProjects).HideObservedProjects(hideObservedProjects).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsProjectcategories(fieldsProjectcategories).FieldsProjectUpdates(fieldsProjectUpdates).FieldsProjectBudgets(fieldsProjectBudgets).FieldsPortfolioColumns(fieldsPortfolioColumns).FieldsPortfolioCards(fieldsPortfolioCards).FieldsPortfolioBoards(fieldsPortfolioBoards).FieldsCustomfields(fieldsCustomfields).FieldsCustomfieldProjects(fieldsCustomfieldProjects).FieldsCompanies(fieldsCompanies).Execute()

Project custom field values.



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
    resp, r, err := api_client.CustomFieldsApi.GETProjectsApiV3ProjectsprojectIdCustomfieldsJson(context.Background(), projectId).UpdatedAfter(updatedAfter).SearchTerm(searchTerm).ReportType(reportType).ReportFormat(reportFormat).ProjectType(projectType).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).MinLastActivityDate(minLastActivityDate).MaxLastActivityDate(maxLastActivityDate).UserId(userId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).OnlyStarredProjects(onlyStarredProjects).OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership).OnlyArchivedProjects(onlyArchivedProjects).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeProjectUserInfo(includeProjectUserInfo).IncludeCustomFields(includeCustomFields).IncludeCompletedStatus(includeCompletedStatus).IncludeArchivedProjects(includeArchivedProjects).HideObservedProjects(hideObservedProjects).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsProjectcategories(fieldsProjectcategories).FieldsProjectUpdates(fieldsProjectUpdates).FieldsProjectBudgets(fieldsProjectBudgets).FieldsPortfolioColumns(fieldsPortfolioColumns).FieldsPortfolioCards(fieldsPortfolioCards).FieldsPortfolioBoards(fieldsPortfolioBoards).FieldsCustomfields(fieldsCustomfields).FieldsCustomfieldProjects(fieldsCustomfieldProjects).FieldsCompanies(fieldsCompanies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.GETProjectsApiV3ProjectsprojectIdCustomfieldsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ProjectsprojectIdCustomfieldsJson`: ProjectCustomFieldProjectsResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFieldsApi.GETProjectsApiV3ProjectsprojectIdCustomfieldsJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsprojectIdCustomfieldsJsonRequest struct via the builder pattern


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

[**ProjectCustomFieldProjectsResponse**](ProjectCustomFieldProjectsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3ProjectsprojectIdCustomfieldscustomFieldIdJson

> ProjectResponse GETProjectsApiV3ProjectsprojectIdCustomfieldscustomFieldIdJson(ctx, projectId, customFieldId).Execute()

Project custom field value.



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
    customFieldId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.GETProjectsApiV3ProjectsprojectIdCustomfieldscustomFieldIdJson(context.Background(), projectId, customFieldId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.GETProjectsApiV3ProjectsprojectIdCustomfieldscustomFieldIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ProjectsprojectIdCustomfieldscustomFieldIdJson`: ProjectResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFieldsApi.GETProjectsApiV3ProjectsprojectIdCustomfieldscustomFieldIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**customFieldId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsprojectIdCustomfieldscustomFieldIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProjectResponse**](ProjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3TaskstaskIdCustomfieldsJson

> TaskCustomFieldTasksResponse GETProjectsApiV3TaskstaskIdCustomfieldsJson(ctx, taskId2).TaskId(taskId).PageSize(pageSize).Page(page).Include(include).FieldsTasks(fieldsTasks).FieldsCustomfields(fieldsCustomfields).FieldsCustomfieldTasks(fieldsCustomfieldTasks).CustomFieldIds(customFieldIds).Execute()

Task custom field values.



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
    taskId2 := int32(56) // int32 | 
    taskId := int32(56) // int32 | filter by task id (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsTasks := []string{"Inner_example"} // []string |  (optional)
    fieldsCustomfields := []string{"Inner_example"} // []string |  (optional)
    fieldsCustomfieldTasks := []string{"Inner_example"} // []string |  (optional)
    customFieldIds := []int32{int32(123)} // []int32 | filter by custom field ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.GETProjectsApiV3TaskstaskIdCustomfieldsJson(context.Background(), taskId2).TaskId(taskId).PageSize(pageSize).Page(page).Include(include).FieldsTasks(fieldsTasks).FieldsCustomfields(fieldsCustomfields).FieldsCustomfieldTasks(fieldsCustomfieldTasks).CustomFieldIds(customFieldIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.GETProjectsApiV3TaskstaskIdCustomfieldsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3TaskstaskIdCustomfieldsJson`: TaskCustomFieldTasksResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFieldsApi.GETProjectsApiV3TaskstaskIdCustomfieldsJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3TaskstaskIdCustomfieldsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taskId** | **int32** | filter by task id | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **include** | **[]string** | include | 
 **fieldsTasks** | **[]string** |  | 
 **fieldsCustomfields** | **[]string** |  | 
 **fieldsCustomfieldTasks** | **[]string** |  | 
 **customFieldIds** | **[]int32** | filter by custom field ids | 

### Return type

[**TaskCustomFieldTasksResponse**](TaskCustomFieldTasksResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3TaskstaskIdCustomfieldscustomFieldIdJson

> TaskResponse GETProjectsApiV3TaskstaskIdCustomfieldscustomFieldIdJson(ctx, taskId, customFieldId).Execute()

Task custom field value.



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
    taskId := int32(56) // int32 | 
    customFieldId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.GETProjectsApiV3TaskstaskIdCustomfieldscustomFieldIdJson(context.Background(), taskId, customFieldId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.GETProjectsApiV3TaskstaskIdCustomfieldscustomFieldIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3TaskstaskIdCustomfieldscustomFieldIdJson`: TaskResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFieldsApi.GETProjectsApiV3TaskstaskIdCustomfieldscustomFieldIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **int32** |  | 
**customFieldId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3TaskstaskIdCustomfieldscustomFieldIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TaskResponse**](TaskResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHProjectsApiV3CustomfieldscustomFieldIdJson

> CustomfieldResponse PATCHProjectsApiV3CustomfieldscustomFieldIdJson(ctx, customFieldId).CustomfieldRequest(customfieldRequest).Execute()

Update an existing custom field.



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
    customFieldId := int32(56) // int32 | 
    customfieldRequest := *openapiclient.NewCustomfieldRequest() // CustomfieldRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.PATCHProjectsApiV3CustomfieldscustomFieldIdJson(context.Background(), customFieldId).CustomfieldRequest(customfieldRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.PATCHProjectsApiV3CustomfieldscustomFieldIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHProjectsApiV3CustomfieldscustomFieldIdJson`: CustomfieldResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFieldsApi.PATCHProjectsApiV3CustomfieldscustomFieldIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customFieldId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHProjectsApiV3CustomfieldscustomFieldIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customfieldRequest** | [**CustomfieldRequest**](CustomfieldRequest.md) |  | 

### Return type

[**CustomfieldResponse**](CustomfieldResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHProjectsApiV3ProjectsprojectIdCustomfieldscustomFieldIdJson

> ProjectResponse PATCHProjectsApiV3ProjectsprojectIdCustomfieldscustomFieldIdJson(ctx, projectId, customFieldId).ProjectRequest(projectRequest).Execute()

Update an existing project custom field value.



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
    customFieldId := int32(56) // int32 | 
    projectRequest := *openapiclient.NewProjectRequest() // ProjectRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.PATCHProjectsApiV3ProjectsprojectIdCustomfieldscustomFieldIdJson(context.Background(), projectId, customFieldId).ProjectRequest(projectRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.PATCHProjectsApiV3ProjectsprojectIdCustomfieldscustomFieldIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHProjectsApiV3ProjectsprojectIdCustomfieldscustomFieldIdJson`: ProjectResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFieldsApi.PATCHProjectsApiV3ProjectsprojectIdCustomfieldscustomFieldIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**customFieldId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHProjectsApiV3ProjectsprojectIdCustomfieldscustomFieldIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **projectRequest** | [**ProjectRequest**](ProjectRequest.md) |  | 

### Return type

[**ProjectResponse**](ProjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHProjectsApiV3TaskstaskIdCustomfieldscustomFieldIdJson

> TaskResponse PATCHProjectsApiV3TaskstaskIdCustomfieldscustomFieldIdJson(ctx, taskId, customFieldId).TaskRequest(taskRequest).Execute()

Update an existing task custom field value.



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
    taskId := int32(56) // int32 | 
    customFieldId := int32(56) // int32 | 
    taskRequest := *openapiclient.NewTaskRequest() // TaskRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.PATCHProjectsApiV3TaskstaskIdCustomfieldscustomFieldIdJson(context.Background(), taskId, customFieldId).TaskRequest(taskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.PATCHProjectsApiV3TaskstaskIdCustomfieldscustomFieldIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHProjectsApiV3TaskstaskIdCustomfieldscustomFieldIdJson`: TaskResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFieldsApi.PATCHProjectsApiV3TaskstaskIdCustomfieldscustomFieldIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **int32** |  | 
**customFieldId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHProjectsApiV3TaskstaskIdCustomfieldscustomFieldIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **taskRequest** | [**TaskRequest**](TaskRequest.md) |  | 

### Return type

[**TaskResponse**](TaskResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsApiV3CustomfieldsBulkDeleteJson

> POSTProjectsApiV3CustomfieldsBulkDeleteJson(ctx).CustomfieldBulkDeleteRequest(customfieldBulkDeleteRequest).Execute()

Delete many custom fields at once



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
    customfieldBulkDeleteRequest := *openapiclient.NewCustomfieldBulkDeleteRequest() // CustomfieldBulkDeleteRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.POSTProjectsApiV3CustomfieldsBulkDeleteJson(context.Background()).CustomfieldBulkDeleteRequest(customfieldBulkDeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.POSTProjectsApiV3CustomfieldsBulkDeleteJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3CustomfieldsBulkDeleteJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customfieldBulkDeleteRequest** | [**CustomfieldBulkDeleteRequest**](CustomfieldBulkDeleteRequest.md) |  | 

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


## POSTProjectsApiV3CustomfieldsJson

> CustomfieldResponse POSTProjectsApiV3CustomfieldsJson(ctx).CustomfieldRequest(customfieldRequest).Execute()

Create a new custom field



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
    customfieldRequest := *openapiclient.NewCustomfieldRequest() // CustomfieldRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.POSTProjectsApiV3CustomfieldsJson(context.Background()).CustomfieldRequest(customfieldRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.POSTProjectsApiV3CustomfieldsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsApiV3CustomfieldsJson`: CustomfieldResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFieldsApi.POSTProjectsApiV3CustomfieldsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3CustomfieldsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customfieldRequest** | [**CustomfieldRequest**](CustomfieldRequest.md) |  | 

### Return type

[**CustomfieldResponse**](CustomfieldResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsApiV3ProjectsprojectIdCustomfieldsBulkDeleteJson

> POSTProjectsApiV3ProjectsprojectIdCustomfieldsBulkDeleteJson(ctx, projectId).ProjectBulkDeleteRequest(projectBulkDeleteRequest).Execute()

Delete many project custom fields values at once.



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
    projectBulkDeleteRequest := *openapiclient.NewProjectBulkDeleteRequest() // ProjectBulkDeleteRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.POSTProjectsApiV3ProjectsprojectIdCustomfieldsBulkDeleteJson(context.Background(), projectId).ProjectBulkDeleteRequest(projectBulkDeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.POSTProjectsApiV3ProjectsprojectIdCustomfieldsBulkDeleteJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3ProjectsprojectIdCustomfieldsBulkDeleteJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectBulkDeleteRequest** | [**ProjectBulkDeleteRequest**](ProjectBulkDeleteRequest.md) |  | 

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


## POSTProjectsApiV3ProjectsprojectIdCustomfieldsBulkUpdateJson

> ProjectCustomFieldProjectsResponse POSTProjectsApiV3ProjectsprojectIdCustomfieldsBulkUpdateJson(ctx, projectId).ProjectBulkUpdateRequest(projectBulkUpdateRequest).Execute()

Update many project custom field values at once.



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
    projectBulkUpdateRequest := *openapiclient.NewProjectBulkUpdateRequest() // ProjectBulkUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.POSTProjectsApiV3ProjectsprojectIdCustomfieldsBulkUpdateJson(context.Background(), projectId).ProjectBulkUpdateRequest(projectBulkUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.POSTProjectsApiV3ProjectsprojectIdCustomfieldsBulkUpdateJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsApiV3ProjectsprojectIdCustomfieldsBulkUpdateJson`: ProjectCustomFieldProjectsResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFieldsApi.POSTProjectsApiV3ProjectsprojectIdCustomfieldsBulkUpdateJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3ProjectsprojectIdCustomfieldsBulkUpdateJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectBulkUpdateRequest** | [**ProjectBulkUpdateRequest**](ProjectBulkUpdateRequest.md) |  | 

### Return type

[**ProjectCustomFieldProjectsResponse**](ProjectCustomFieldProjectsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsApiV3ProjectsprojectIdCustomfieldsJson

> ProjectResponse POSTProjectsApiV3ProjectsprojectIdCustomfieldsJson(ctx, projectId).ProjectRequest(projectRequest).Execute()

Add project custom field value.



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
    projectRequest := *openapiclient.NewProjectRequest() // ProjectRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.POSTProjectsApiV3ProjectsprojectIdCustomfieldsJson(context.Background(), projectId).ProjectRequest(projectRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.POSTProjectsApiV3ProjectsprojectIdCustomfieldsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsApiV3ProjectsprojectIdCustomfieldsJson`: ProjectResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFieldsApi.POSTProjectsApiV3ProjectsprojectIdCustomfieldsJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3ProjectsprojectIdCustomfieldsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectRequest** | [**ProjectRequest**](ProjectRequest.md) |  | 

### Return type

[**ProjectResponse**](ProjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsApiV3TaskstaskIdCustomfieldsBulkDeleteJson

> POSTProjectsApiV3TaskstaskIdCustomfieldsBulkDeleteJson(ctx, taskId).TaskBulkDeleteRequest(taskBulkDeleteRequest).Execute()

Delete many task custom fields values at once.



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
    taskId := int32(56) // int32 | 
    taskBulkDeleteRequest := *openapiclient.NewTaskBulkDeleteRequest() // TaskBulkDeleteRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.POSTProjectsApiV3TaskstaskIdCustomfieldsBulkDeleteJson(context.Background(), taskId).TaskBulkDeleteRequest(taskBulkDeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.POSTProjectsApiV3TaskstaskIdCustomfieldsBulkDeleteJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3TaskstaskIdCustomfieldsBulkDeleteJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taskBulkDeleteRequest** | [**TaskBulkDeleteRequest**](TaskBulkDeleteRequest.md) |  | 

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


## POSTProjectsApiV3TaskstaskIdCustomfieldsBulkUpdateJson

> TaskCustomFieldTasksResponse POSTProjectsApiV3TaskstaskIdCustomfieldsBulkUpdateJson(ctx, taskId).TaskBulkUpdateRequest(taskBulkUpdateRequest).Execute()

Update many task custom field values at once.



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
    taskId := int32(56) // int32 | 
    taskBulkUpdateRequest := *openapiclient.NewTaskBulkUpdateRequest() // TaskBulkUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.POSTProjectsApiV3TaskstaskIdCustomfieldsBulkUpdateJson(context.Background(), taskId).TaskBulkUpdateRequest(taskBulkUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.POSTProjectsApiV3TaskstaskIdCustomfieldsBulkUpdateJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsApiV3TaskstaskIdCustomfieldsBulkUpdateJson`: TaskCustomFieldTasksResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFieldsApi.POSTProjectsApiV3TaskstaskIdCustomfieldsBulkUpdateJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3TaskstaskIdCustomfieldsBulkUpdateJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taskBulkUpdateRequest** | [**TaskBulkUpdateRequest**](TaskBulkUpdateRequest.md) |  | 

### Return type

[**TaskCustomFieldTasksResponse**](TaskCustomFieldTasksResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsApiV3TaskstaskIdCustomfieldsJson

> TaskResponse POSTProjectsApiV3TaskstaskIdCustomfieldsJson(ctx, taskId).TaskRequest(taskRequest).Execute()

Add task custom field value.



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
    taskId := int32(56) // int32 | 
    taskRequest := *openapiclient.NewTaskRequest() // TaskRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.POSTProjectsApiV3TaskstaskIdCustomfieldsJson(context.Background(), taskId).TaskRequest(taskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.POSTProjectsApiV3TaskstaskIdCustomfieldsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsApiV3TaskstaskIdCustomfieldsJson`: TaskResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFieldsApi.POSTProjectsApiV3TaskstaskIdCustomfieldsJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3TaskstaskIdCustomfieldsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taskRequest** | [**TaskRequest**](TaskRequest.md) |  | 

### Return type

[**TaskResponse**](TaskResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV3CustomfieldscustomFieldIdJson

> CustomfieldResponse PUTProjectsApiV3CustomfieldscustomFieldIdJson(ctx, customFieldId).CustomfieldRequest(customfieldRequest).Execute()

Update an existing custom field.



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
    customFieldId := int32(56) // int32 | 
    customfieldRequest := *openapiclient.NewCustomfieldRequest() // CustomfieldRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.PUTProjectsApiV3CustomfieldscustomFieldIdJson(context.Background(), customFieldId).CustomfieldRequest(customfieldRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.PUTProjectsApiV3CustomfieldscustomFieldIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsApiV3CustomfieldscustomFieldIdJson`: CustomfieldResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFieldsApi.PUTProjectsApiV3CustomfieldscustomFieldIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customFieldId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3CustomfieldscustomFieldIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customfieldRequest** | [**CustomfieldRequest**](CustomfieldRequest.md) |  | 

### Return type

[**CustomfieldResponse**](CustomfieldResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

