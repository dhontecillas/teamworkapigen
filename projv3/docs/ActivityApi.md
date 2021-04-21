# \ActivityApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3LatestactivityJson**](ActivityApi.md#GETProjectsApiV3LatestactivityJson) | **Get** /projects/api/v3/latestactivity.json | Latest activity (all projects)
[**GETProjectsApiV3ProjectsprojectIdLatestactivity**](ActivityApi.md#GETProjectsApiV3ProjectsprojectIdLatestactivity) | **Get** /projects/api/v3/projects/{projectId}/latestactivity | List latest activity for a specific project.



## GETProjectsApiV3LatestactivityJson

> ActivityActivitiesResponse GETProjectsApiV3LatestactivityJson(ctx).UpdatedAfter(updatedAfter).StartDate(startDate).Sort(sort).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).EndDate(endDate).ActivityTypes(activityTypes).ProjectId(projectId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).MaxId(maxId).CatchupFromId(catchupFromId).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).MatchAllProjectTags(matchAllProjectTags).IncludeArchivedProjects(includeArchivedProjects).GroupCommentActivityType(groupCommentActivityType).UserIds(userIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsProjects(fieldsProjects).FieldsCompanies(fieldsCompanies).FieldsActivities(fieldsActivities).Execute()

Latest activity (all projects)



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
    startDate := time.Now() // time.Time | filter by start datetime (optional)
    sort := "sort_example" // string | sort by (optional)
    projectStatuses := "projectStatuses_example" // string | list of project status (optional)
    projectStatus := "projectStatus_example" // string | filter by project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional)
    endDate := time.Now() // time.Time | filter by end datetime (optional)
    activityTypes := "activityTypes_example" // string | filter by activity types (optional)
    projectId := int32(56) // int32 | filter by project id (optional)
    projectHealths := int32(56) // int32 | list of project health (optional)
    pageSize := int32(56) // int32 | number of items in a page default: 50 (optional)
    page := int32(56) // int32 | page number default: 1 (optional)
    maxId := int32(56) // int32 | activity id offset (optional)
    catchupFromId := int32(56) // int32 | limit results from previous id (optional)
    showDeleted := true // bool | include deleted items default: false (optional)
    onlyStarredProjects := true // bool | filter by starred projects only default: false (optional)
    matchAllProjectTags := true // bool | enforce all tag ids must be matched (optional)
    includeArchivedProjects := true // bool | include archived projects (optional)
    groupCommentActivityType := true // bool | group all activities on comments (optional)
    userIds := []int32{int32(123)} // []int32 | filter by user ids (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)
    fieldsActivities := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActivityApi.GETProjectsApiV3LatestactivityJson(context.Background()).UpdatedAfter(updatedAfter).StartDate(startDate).Sort(sort).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).EndDate(endDate).ActivityTypes(activityTypes).ProjectId(projectId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).MaxId(maxId).CatchupFromId(catchupFromId).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).MatchAllProjectTags(matchAllProjectTags).IncludeArchivedProjects(includeArchivedProjects).GroupCommentActivityType(groupCommentActivityType).UserIds(userIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsProjects(fieldsProjects).FieldsCompanies(fieldsCompanies).FieldsActivities(fieldsActivities).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.GETProjectsApiV3LatestactivityJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3LatestactivityJson`: ActivityActivitiesResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.GETProjectsApiV3LatestactivityJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3LatestactivityJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfter** | **time.Time** | filter by updated after date | 
 **startDate** | **time.Time** | filter by start datetime | 
 **sort** | **string** | sort by | 
 **projectStatuses** | **string** | list of project status | 
 **projectStatus** | **string** | filter by project status | 
 **orderMode** | **string** | order mode | 
 **endDate** | **time.Time** | filter by end datetime | 
 **activityTypes** | **string** | filter by activity types | 
 **projectId** | **int32** | filter by project id | 
 **projectHealths** | **int32** | list of project health | 
 **pageSize** | **int32** | number of items in a page default: 50 | 
 **page** | **int32** | page number default: 1 | 
 **maxId** | **int32** | activity id offset | 
 **catchupFromId** | **int32** | limit results from previous id | 
 **showDeleted** | **bool** | include deleted items default: false | 
 **onlyStarredProjects** | **bool** | filter by starred projects only default: false | 
 **matchAllProjectTags** | **bool** | enforce all tag ids must be matched | 
 **includeArchivedProjects** | **bool** | include archived projects | 
 **groupCommentActivityType** | **bool** | group all activities on comments | 
 **userIds** | **[]int32** | filter by user ids | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsCompanies** | **[]string** |  | 
 **fieldsActivities** | **[]string** |  | 

### Return type

[**ActivityActivitiesResponse**](ActivityActivitiesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3ProjectsprojectIdLatestactivity

> ActivityActivitiesResponse GETProjectsApiV3ProjectsprojectIdLatestactivity(ctx, projectId2).UpdatedAfter(updatedAfter).StartDate(startDate).Sort(sort).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).EndDate(endDate).ActivityTypes(activityTypes).ProjectId(projectId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).MaxId(maxId).CatchupFromId(catchupFromId).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).MatchAllProjectTags(matchAllProjectTags).IncludeArchivedProjects(includeArchivedProjects).GroupCommentActivityType(groupCommentActivityType).UserIds(userIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsProjects(fieldsProjects).FieldsCompanies(fieldsCompanies).FieldsActivities(fieldsActivities).Execute()

List latest activity for a specific project.



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
    projectId2 := int32(56) // int32 | 
    updatedAfter := time.Now() // time.Time | filter by updated after date (optional)
    startDate := time.Now() // time.Time | filter by start datetime (optional)
    sort := "sort_example" // string | sort by (optional)
    projectStatuses := "projectStatuses_example" // string | list of project status (optional)
    projectStatus := "projectStatus_example" // string | filter by project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional)
    endDate := time.Now() // time.Time | filter by end datetime (optional)
    activityTypes := "activityTypes_example" // string | filter by activity types (optional)
    projectId := int32(56) // int32 | filter by project id (optional)
    projectHealths := int32(56) // int32 | list of project health (optional)
    pageSize := int32(56) // int32 | number of items in a page default: 50 (optional)
    page := int32(56) // int32 | page number default: 1 (optional)
    maxId := int32(56) // int32 | activity id offset (optional)
    catchupFromId := int32(56) // int32 | limit results from previous id (optional)
    showDeleted := true // bool | include deleted items default: false (optional)
    onlyStarredProjects := true // bool | filter by starred projects only default: false (optional)
    matchAllProjectTags := true // bool | enforce all tag ids must be matched (optional)
    includeArchivedProjects := true // bool | include archived projects (optional)
    groupCommentActivityType := true // bool | group all activities on comments (optional)
    userIds := []int32{int32(123)} // []int32 | filter by user id (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)
    fieldsActivities := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActivityApi.GETProjectsApiV3ProjectsprojectIdLatestactivity(context.Background(), projectId2).UpdatedAfter(updatedAfter).StartDate(startDate).Sort(sort).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).EndDate(endDate).ActivityTypes(activityTypes).ProjectId(projectId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).MaxId(maxId).CatchupFromId(catchupFromId).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).MatchAllProjectTags(matchAllProjectTags).IncludeArchivedProjects(includeArchivedProjects).GroupCommentActivityType(groupCommentActivityType).UserIds(userIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsProjects(fieldsProjects).FieldsCompanies(fieldsCompanies).FieldsActivities(fieldsActivities).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.GETProjectsApiV3ProjectsprojectIdLatestactivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ProjectsprojectIdLatestactivity`: ActivityActivitiesResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.GETProjectsApiV3ProjectsprojectIdLatestactivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsprojectIdLatestactivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedAfter** | **time.Time** | filter by updated after date | 
 **startDate** | **time.Time** | filter by start datetime | 
 **sort** | **string** | sort by | 
 **projectStatuses** | **string** | list of project status | 
 **projectStatus** | **string** | filter by project status | 
 **orderMode** | **string** | order mode | 
 **endDate** | **time.Time** | filter by end datetime | 
 **activityTypes** | **string** | filter by activity types | 
 **projectId** | **int32** | filter by project id | 
 **projectHealths** | **int32** | list of project health | 
 **pageSize** | **int32** | number of items in a page default: 50 | 
 **page** | **int32** | page number default: 1 | 
 **maxId** | **int32** | activity id offset | 
 **catchupFromId** | **int32** | limit results from previous id | 
 **showDeleted** | **bool** | include deleted items default: false | 
 **onlyStarredProjects** | **bool** | filter by starred projects only default: false | 
 **matchAllProjectTags** | **bool** | enforce all tag ids must be matched | 
 **includeArchivedProjects** | **bool** | include archived projects | 
 **groupCommentActivityType** | **bool** | group all activities on comments | 
 **userIds** | **[]int32** | filter by user id | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsCompanies** | **[]string** |  | 
 **fieldsActivities** | **[]string** |  | 

### Return type

[**ActivityActivitiesResponse**](ActivityActivitiesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

