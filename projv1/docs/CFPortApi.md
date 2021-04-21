# \CFPortApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV1Completedtasks**](CFPortApi.md#GETProjectsApiV1Completedtasks) | **Get** /projects/api/v1/completedtasks | Will return the completed tasks
[**GETProjectsApiV1PendingfilesPresignedurl**](CFPortApi.md#GETProjectsApiV1PendingfilesPresignedurl) | **Get** /projects/api/v1/pendingfiles/presignedurl | Get Presigned Url for Pending files
[**GETProjectsApiV1ProjectcategoriesprojectCategoryIdTasks**](CFPortApi.md#GETProjectsApiV1ProjectcategoriesprojectCategoryIdTasks) | **Get** /projects/api/v1/projectcategories/{projectCategoryId}/tasks | Will return the tasks
[**GETProjectsApiV1ProjectprojectIdCompletedtasks**](CFPortApi.md#GETProjectsApiV1ProjectprojectIdCompletedtasks) | **Get** /projects/api/v1/project/{projectId}/completedtasks | Will return the project completed tasks
[**GETProjectsApiV1Projects**](CFPortApi.md#GETProjectsApiV1Projects) | **Get** /projects/api/v1/projects | Get Projects (APICall_GET_projects)
[**GETProjectsApiV1ProjectsprojectIdTasklists**](CFPortApi.md#GETProjectsApiV1ProjectsprojectIdTasklists) | **Get** /projects/api/v1/projects/{projectId}/tasklists | Get tasklists
[**GETProjectsApiV1ProjectsprojectIdTasks**](CFPortApi.md#GETProjectsApiV1ProjectsprojectIdTasks) | **Get** /projects/api/v1/projects/{projectId}/tasks | Will return the tasks
[**GETProjectsApiV1Tasklists**](CFPortApi.md#GETProjectsApiV1Tasklists) | **Get** /projects/api/v1/tasklists | Get tasklists
[**GETProjectsApiV1TaskliststasklistId**](CFPortApi.md#GETProjectsApiV1TaskliststasklistId) | **Get** /projects/api/v1/tasklists/{tasklistId} | Get tasklists
[**GETProjectsApiV1TaskliststasklistIdTasks**](CFPortApi.md#GETProjectsApiV1TaskliststasklistIdTasks) | **Get** /projects/api/v1/tasklists/{tasklistId}/tasks | Will return the tasks
[**GETProjectsApiV1Tasks**](CFPortApi.md#GETProjectsApiV1Tasks) | **Get** /projects/api/v1/tasks | Will return the tasks
[**GETProjectsApiV1TaskstaskId**](CFPortApi.md#GETProjectsApiV1TaskstaskId) | **Get** /projects/api/v1/tasks/{taskId} | Will return the tasks
[**GETProjectsApiV1TaskstaskIdDependents**](CFPortApi.md#GETProjectsApiV1TaskstaskIdDependents) | **Get** /projects/api/v1/tasks/{taskId}/dependents | Will return the tasks dependents tasks
[**GETProjectsApiV1TaskstaskIdPredecessors**](CFPortApi.md#GETProjectsApiV1TaskstaskIdPredecessors) | **Get** /projects/api/v1/tasks/{taskId}/predecessors | Will return the tasks predecessor tasks
[**GETProjectsApiV1TaskstaskIdSubtasks**](CFPortApi.md#GETProjectsApiV1TaskstaskIdSubtasks) | **Get** /projects/api/v1/tasks/{taskId}/subtasks | Will return the tasks
[**GETProjectsApiV1Time**](CFPortApi.md#GETProjectsApiV1Time) | **Get** /projects/api/v1/time | APICall_GET_time_entries
[**GETProjectsApiV1TimeEntries**](CFPortApi.md#GETProjectsApiV1TimeEntries) | **Get** /projects/api/v1/time_entries | APICall_GET_time_entries
[**POSTProjectsApiV1Pendingfiles**](CFPortApi.md#POSTProjectsApiV1Pendingfiles) | **Post** /projects/api/v1/pendingfiles | Posts a pending file



## GETProjectsApiV1Completedtasks

> GETProjectsApiV1Completedtasks(ctx).UpdatedAfterDate(updatedAfterDate).Type_(type_).StartDate(startDate).SortOrder(sortOrder).SortBy(sortBy).SearchTerm(searchTerm).ReportFormat(reportFormat).ProjectStatus(projectStatus).Priority(priority).FilterBy(filterBy).EndDate(endDate).DateType(dateType).UserId(userId).PageSize(pageSize).Page(page).StarredProjectsOnly(starredProjectsOnly).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyTasksWithTickets(onlyTasksWithTickets).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllExcludedTags(matchAllExcludedTags).IsReportDownload(isReportDownload).IncludeTags(includeTags).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).TagIds(tagIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).FollowerIds(followerIds).ExcludeTagIds(excludeTagIds).CreatorIds(creatorIds).CompletedByIds(completedByIds).AssigneeTeamIds(assigneeTeamIds).AssigneeCompanyIds(assigneeCompanyIds).Execute()

Will return the completed tasks

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
    updatedAfterDate := "updatedAfterDate_example" // string |  (optional)
    type_ := "type__example" // string |  (optional)
    startDate := "startDate_example" // string |  (optional)
    sortOrder := "sortOrder_example" // string |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    searchTerm := "searchTerm_example" // string |  (optional)
    reportFormat := "reportFormat_example" // string |  (optional)
    projectStatus := "projectStatus_example" // string |  (optional)
    priority := "priority_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    dateType := "dateType_example" // string |  (optional)
    userId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    starredProjectsOnly := true // bool |  (optional)
    onlyUntaggedTasks := true // bool |  (optional)
    onlyTasksWithTickets := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    matchAllTags := true // bool |  (optional)
    matchAllProjectTags := true // bool |  (optional)
    matchAllExcludedTags := true // bool |  (optional)
    isReportDownload := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    includeAssigneeTeams := true // bool |  (optional)
    includeAssigneeCompanies := true // bool |  (optional)
    includeArchivedProjects := true // bool |  (optional)
    tagIds := []int32{int32(123)} // []int32 |  (optional)
    responsiblePartyIds := []int32{int32(123)} // []int32 |  (optional)
    projectTagIds := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    projectIds := []int32{int32(123)} // []int32 |  (optional)
    projectId := []int32{int32(123)} // []int32 | project filters (optional)
    projectHealths := []int32{int32(123)} // []int32 |  (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 |  (optional)
    followerIds := []int32{int32(123)} // []int32 |  (optional)
    excludeTagIds := []int32{int32(123)} // []int32 |  (optional)
    creatorIds := []int32{int32(123)} // []int32 |  (optional)
    completedByIds := []int32{int32(123)} // []int32 |  (optional)
    assigneeTeamIds := []int32{int32(123)} // []int32 |  (optional)
    assigneeCompanyIds := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV1Completedtasks(context.Background()).UpdatedAfterDate(updatedAfterDate).Type_(type_).StartDate(startDate).SortOrder(sortOrder).SortBy(sortBy).SearchTerm(searchTerm).ReportFormat(reportFormat).ProjectStatus(projectStatus).Priority(priority).FilterBy(filterBy).EndDate(endDate).DateType(dateType).UserId(userId).PageSize(pageSize).Page(page).StarredProjectsOnly(starredProjectsOnly).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyTasksWithTickets(onlyTasksWithTickets).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllExcludedTags(matchAllExcludedTags).IsReportDownload(isReportDownload).IncludeTags(includeTags).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).TagIds(tagIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).FollowerIds(followerIds).ExcludeTagIds(excludeTagIds).CreatorIds(creatorIds).CompletedByIds(completedByIds).AssigneeTeamIds(assigneeTeamIds).AssigneeCompanyIds(assigneeCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV1Completedtasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV1CompletedtasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfterDate** | **string** |  | 
 **type_** | **string** |  | 
 **startDate** | **string** |  | 
 **sortOrder** | **string** |  | 
 **sortBy** | **string** |  | 
 **searchTerm** | **string** |  | 
 **reportFormat** | **string** |  | 
 **projectStatus** | **string** |  | 
 **priority** | **string** |  | 
 **filterBy** | **string** |  | 
 **endDate** | **string** |  | 
 **dateType** | **string** |  | 
 **userId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **starredProjectsOnly** | **bool** |  | 
 **onlyUntaggedTasks** | **bool** |  | 
 **onlyTasksWithTickets** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **matchAllTags** | **bool** |  | 
 **matchAllProjectTags** | **bool** |  | 
 **matchAllExcludedTags** | **bool** |  | 
 **isReportDownload** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **includeAssigneeTeams** | **bool** |  | 
 **includeAssigneeCompanies** | **bool** |  | 
 **includeArchivedProjects** | **bool** |  | 
 **tagIds** | **[]int32** |  | 
 **responsiblePartyIds** | **[]int32** |  | 
 **projectTagIds** | **[]int32** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **projectIds** | **[]int32** |  | 
 **projectId** | **[]int32** | project filters | 
 **projectHealths** | **[]int32** |  | 
 **projectCompanyIds** | **[]int32** |  | 
 **projectCategoryIds** | **[]int32** |  | 
 **followerIds** | **[]int32** |  | 
 **excludeTagIds** | **[]int32** |  | 
 **creatorIds** | **[]int32** |  | 
 **completedByIds** | **[]int32** |  | 
 **assigneeTeamIds** | **[]int32** |  | 
 **assigneeCompanyIds** | **[]int32** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV1PendingfilesPresignedurl

> GETProjectsApiV1PendingfilesPresignedurl(ctx).FileName(fileName).FileSize(fileSize).Execute()

Get Presigned Url for Pending files

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
    fileName := "fileName_example" // string |  (optional)
    fileSize := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV1PendingfilesPresignedurl(context.Background()).FileName(fileName).FileSize(fileSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV1PendingfilesPresignedurl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV1PendingfilesPresignedurlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileName** | **string** |  | 
 **fileSize** | **int32** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV1ProjectcategoriesprojectCategoryIdTasks

> GETProjectsApiV1ProjectcategoriesprojectCategoryIdTasks(ctx, projectCategoryId).UpdatedAfterDate(updatedAfterDate).Today(today).TagIds(tagIds).StartDate(startDate).Sort(sort).ResponsiblePartyIds(responsiblePartyIds).ProjectStatus(projectStatus).Include(include).FollowedByUserIds(followedByUserIds).Filter(filter).EndDate(endDate).DataSet(dataSet).CreatorIds(creatorIds).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TasklistId(tasklistId).TaskId(taskId).ProjectCategoryId2(projectCategoryId2).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).IncludeTaskId(includeTaskId).CompanyId(companyId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyStarredProjects(onlyStarredProjects).OnlyArchivedProjects(onlyArchivedProjects).NestSubTasks(nestSubTasks).MatchAllProjectTags(matchAllProjectTags).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompletedPredecessors(includeCompletedPredecessors).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).CountOnly(countOnly).AllowTemplateTasks(allowTemplateTasks).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Execute()

Will return the tasks

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
    projectCategoryId := int32(56) // int32 | 
    updatedAfterDate := "updatedAfterDate_example" // string |  (optional)
    today := "today_example" // string |  (optional)
    tagIds := "tagIds_example" // string |  (optional)
    startDate := "startDate_example" // string |  (optional)
    sort := "sort_example" // string |  (optional)
    responsiblePartyIds := "responsiblePartyIds_example" // string |  (optional)
    projectStatus := "projectStatus_example" // string |  (optional)
    include := "include_example" // string |  (optional)
    followedByUserIds := "followedByUserIds_example" // string |  (optional)
    filter := "filter_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    dataSet := "dataSet_example" // string |  (optional)
    creatorIds := "creatorIds_example" // string |  (optional)
    createdBeforeDate := "createdBeforeDate_example" // string |  (optional)
    createdAfterDate := "createdAfterDate_example" // string |  (optional)
    completedBeforeDate := "completedBeforeDate_example" // string |  (optional)
    completedAfterDate := "completedAfterDate_example" // string |  (optional)
    callback := "callback_example" // string |  (optional)
    tasklistId := int32(56) // int32 |  (optional)
    taskId := int32(56) // int32 |  (optional)
    projectCategoryId2 := int32(56) // int32 |  (optional)
    parentTaskId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    includeTaskId := int32(56) // int32 |  (optional)
    companyId := int32(56) // int32 |  (optional)
    useStartDatesForTodaysTasks := true // bool |  (optional)
    useAllProjects := true // bool |  (optional)
    starredProjectsOnly := true // bool |  (optional)
    showDeleted := true // bool |  (optional)
    showCompletedLists := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    onlyArchivedProjects := true // bool |  (optional)
    nestSubTasks := true // bool |  (optional)
    matchAllProjectTags := true // bool |  (optional)
    includeUntaggedTasks := true // bool |  (optional)
    includeToday := true // bool |  (optional)
    includeTasksWithoutDueDates := true // bool |  (optional)
    includeTasksFromDeletedLists := true // bool |  (optional)
    includeReminders := true // bool |  (optional)
    includeLoggedTime := true // bool |  (optional)
    includeCompletedTasks := true // bool |  (optional)
    includeCompletedSubtasks := true // bool |  (optional)
    includeCompletedPredecessors := true // bool |  (optional)
    includeArchivedProjects := true // bool |  (optional)
    ignoreStartDates := true // bool |  (optional)
    getSubTasks := true // bool |  (optional)
    getFiles := true // bool |  (optional)
    countOnly := true // bool |  (optional)
    allowTemplateTasks := true // bool |  (optional)
    projectTagIds := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    projectIds := []int32{int32(123)} // []int32 |  (optional)
    projectId := []int32{int32(123)} // []int32 | project *filters (optional)
    projectHealths := []int32{int32(123)} // []int32 |  (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV1ProjectcategoriesprojectCategoryIdTasks(context.Background(), projectCategoryId).UpdatedAfterDate(updatedAfterDate).Today(today).TagIds(tagIds).StartDate(startDate).Sort(sort).ResponsiblePartyIds(responsiblePartyIds).ProjectStatus(projectStatus).Include(include).FollowedByUserIds(followedByUserIds).Filter(filter).EndDate(endDate).DataSet(dataSet).CreatorIds(creatorIds).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TasklistId(tasklistId).TaskId(taskId).ProjectCategoryId2(projectCategoryId2).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).IncludeTaskId(includeTaskId).CompanyId(companyId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyStarredProjects(onlyStarredProjects).OnlyArchivedProjects(onlyArchivedProjects).NestSubTasks(nestSubTasks).MatchAllProjectTags(matchAllProjectTags).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompletedPredecessors(includeCompletedPredecessors).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).CountOnly(countOnly).AllowTemplateTasks(allowTemplateTasks).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV1ProjectcategoriesprojectCategoryIdTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectCategoryId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV1ProjectcategoriesprojectCategoryIdTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedAfterDate** | **string** |  | 
 **today** | **string** |  | 
 **tagIds** | **string** |  | 
 **startDate** | **string** |  | 
 **sort** | **string** |  | 
 **responsiblePartyIds** | **string** |  | 
 **projectStatus** | **string** |  | 
 **include** | **string** |  | 
 **followedByUserIds** | **string** |  | 
 **filter** | **string** |  | 
 **endDate** | **string** |  | 
 **dataSet** | **string** |  | 
 **creatorIds** | **string** |  | 
 **createdBeforeDate** | **string** |  | 
 **createdAfterDate** | **string** |  | 
 **completedBeforeDate** | **string** |  | 
 **completedAfterDate** | **string** |  | 
 **callback** | **string** |  | 
 **tasklistId** | **int32** |  | 
 **taskId** | **int32** |  | 
 **projectCategoryId2** | **int32** |  | 
 **parentTaskId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **includeTaskId** | **int32** |  | 
 **companyId** | **int32** |  | 
 **useStartDatesForTodaysTasks** | **bool** |  | 
 **useAllProjects** | **bool** |  | 
 **starredProjectsOnly** | **bool** |  | 
 **showDeleted** | **bool** |  | 
 **showCompletedLists** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **onlyArchivedProjects** | **bool** |  | 
 **nestSubTasks** | **bool** |  | 
 **matchAllProjectTags** | **bool** |  | 
 **includeUntaggedTasks** | **bool** |  | 
 **includeToday** | **bool** |  | 
 **includeTasksWithoutDueDates** | **bool** |  | 
 **includeTasksFromDeletedLists** | **bool** |  | 
 **includeReminders** | **bool** |  | 
 **includeLoggedTime** | **bool** |  | 
 **includeCompletedTasks** | **bool** |  | 
 **includeCompletedSubtasks** | **bool** |  | 
 **includeCompletedPredecessors** | **bool** |  | 
 **includeArchivedProjects** | **bool** |  | 
 **ignoreStartDates** | **bool** |  | 
 **getSubTasks** | **bool** |  | 
 **getFiles** | **bool** |  | 
 **countOnly** | **bool** |  | 
 **allowTemplateTasks** | **bool** |  | 
 **projectTagIds** | **[]int32** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **projectIds** | **[]int32** |  | 
 **projectId** | **[]int32** | project *filters | 
 **projectHealths** | **[]int32** |  | 
 **projectCompanyIds** | **[]int32** |  | 
 **projectCategoryIds** | **[]int32** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV1ProjectprojectIdCompletedtasks

> GETProjectsApiV1ProjectprojectIdCompletedtasks(ctx, projectId).UpdatedAfterDate(updatedAfterDate).Type_(type_).StartDate(startDate).SortOrder(sortOrder).SortBy(sortBy).SearchTerm(searchTerm).ReportFormat(reportFormat).ProjectStatus(projectStatus).Priority(priority).FilterBy(filterBy).EndDate(endDate).DateType(dateType).UserId(userId).PageSize(pageSize).Page(page).StarredProjectsOnly(starredProjectsOnly).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyTasksWithTickets(onlyTasksWithTickets).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllExcludedTags(matchAllExcludedTags).IsReportDownload(isReportDownload).IncludeTags(includeTags).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).TagIds(tagIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId2(projectId2).ProjectHealths(projectHealths).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).FollowerIds(followerIds).ExcludeTagIds(excludeTagIds).CreatorIds(creatorIds).CompletedByIds(completedByIds).AssigneeTeamIds(assigneeTeamIds).AssigneeCompanyIds(assigneeCompanyIds).Execute()

Will return the project completed tasks

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
    updatedAfterDate := "updatedAfterDate_example" // string |  (optional)
    type_ := "type__example" // string |  (optional)
    startDate := "startDate_example" // string |  (optional)
    sortOrder := "sortOrder_example" // string |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    searchTerm := "searchTerm_example" // string |  (optional)
    reportFormat := "reportFormat_example" // string |  (optional)
    projectStatus := "projectStatus_example" // string |  (optional)
    priority := "priority_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    dateType := "dateType_example" // string |  (optional)
    userId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    starredProjectsOnly := true // bool |  (optional)
    onlyUntaggedTasks := true // bool |  (optional)
    onlyTasksWithTickets := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    matchAllTags := true // bool |  (optional)
    matchAllProjectTags := true // bool |  (optional)
    matchAllExcludedTags := true // bool |  (optional)
    isReportDownload := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    includeAssigneeTeams := true // bool |  (optional)
    includeAssigneeCompanies := true // bool |  (optional)
    includeArchivedProjects := true // bool |  (optional)
    tagIds := []int32{int32(123)} // []int32 |  (optional)
    responsiblePartyIds := []int32{int32(123)} // []int32 |  (optional)
    projectTagIds := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    projectIds := []int32{int32(123)} // []int32 |  (optional)
    projectId2 := []int32{int32(123)} // []int32 | project filters (optional)
    projectHealths := []int32{int32(123)} // []int32 |  (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 |  (optional)
    followerIds := []int32{int32(123)} // []int32 |  (optional)
    excludeTagIds := []int32{int32(123)} // []int32 |  (optional)
    creatorIds := []int32{int32(123)} // []int32 |  (optional)
    completedByIds := []int32{int32(123)} // []int32 |  (optional)
    assigneeTeamIds := []int32{int32(123)} // []int32 |  (optional)
    assigneeCompanyIds := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV1ProjectprojectIdCompletedtasks(context.Background(), projectId).UpdatedAfterDate(updatedAfterDate).Type_(type_).StartDate(startDate).SortOrder(sortOrder).SortBy(sortBy).SearchTerm(searchTerm).ReportFormat(reportFormat).ProjectStatus(projectStatus).Priority(priority).FilterBy(filterBy).EndDate(endDate).DateType(dateType).UserId(userId).PageSize(pageSize).Page(page).StarredProjectsOnly(starredProjectsOnly).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyTasksWithTickets(onlyTasksWithTickets).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllExcludedTags(matchAllExcludedTags).IsReportDownload(isReportDownload).IncludeTags(includeTags).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).TagIds(tagIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId2(projectId2).ProjectHealths(projectHealths).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).FollowerIds(followerIds).ExcludeTagIds(excludeTagIds).CreatorIds(creatorIds).CompletedByIds(completedByIds).AssigneeTeamIds(assigneeTeamIds).AssigneeCompanyIds(assigneeCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV1ProjectprojectIdCompletedtasks``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETProjectsApiV1ProjectprojectIdCompletedtasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedAfterDate** | **string** |  | 
 **type_** | **string** |  | 
 **startDate** | **string** |  | 
 **sortOrder** | **string** |  | 
 **sortBy** | **string** |  | 
 **searchTerm** | **string** |  | 
 **reportFormat** | **string** |  | 
 **projectStatus** | **string** |  | 
 **priority** | **string** |  | 
 **filterBy** | **string** |  | 
 **endDate** | **string** |  | 
 **dateType** | **string** |  | 
 **userId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **starredProjectsOnly** | **bool** |  | 
 **onlyUntaggedTasks** | **bool** |  | 
 **onlyTasksWithTickets** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **matchAllTags** | **bool** |  | 
 **matchAllProjectTags** | **bool** |  | 
 **matchAllExcludedTags** | **bool** |  | 
 **isReportDownload** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **includeAssigneeTeams** | **bool** |  | 
 **includeAssigneeCompanies** | **bool** |  | 
 **includeArchivedProjects** | **bool** |  | 
 **tagIds** | **[]int32** |  | 
 **responsiblePartyIds** | **[]int32** |  | 
 **projectTagIds** | **[]int32** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **projectIds** | **[]int32** |  | 
 **projectId2** | **[]int32** | project filters | 
 **projectHealths** | **[]int32** |  | 
 **projectCompanyIds** | **[]int32** |  | 
 **projectCategoryIds** | **[]int32** |  | 
 **followerIds** | **[]int32** |  | 
 **excludeTagIds** | **[]int32** |  | 
 **creatorIds** | **[]int32** |  | 
 **completedByIds** | **[]int32** |  | 
 **assigneeTeamIds** | **[]int32** |  | 
 **assigneeCompanyIds** | **[]int32** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV1Projects

> GETProjectsApiV1Projects(ctx).UpdatedAfterDateTime(updatedAfterDateTime).UpdatedAfterDate(updatedAfterDate).Status(status).SearchTerm(searchTerm).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).Get(get).FirstLetter(firstLetter).DataSet(dataSet).CreatedAfterDateTime(createdAfterDateTime).CreatedAfterDate(createdAfterDate).CompletedBeforeDateTime(completedBeforeDateTime).CompletedAfterDateTime(completedAfterDateTime).UserId(userId).PageSize(pageSize).Page(page).SearchCompany(searchCompany).ReturnLetters(returnLetters).OnlyStarredProjects(onlyStarredProjects).OnlyArchivedProjects(onlyArchivedProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).IncludeUpdates(includeUpdates).IncludeTopPeople(includeTopPeople).IncludeTags(includeTags).IncludeProjectOwner(includeProjectOwner).IncludePeople(includePeople).IncludeArchivedProjects(includeArchivedProjects).HideDesc(hideDesc).GetTruePermissions(getTruePermissions).GetPermissions(getPermissions).GetNotificationSettings(getNotificationSettings).GetDeleted(getDeleted).GetCounts(getCounts).GetCategoryPath(getCategoryPath).GetAll(getAll).GetActivePages(getActivePages).FormatMarkdown(formatMarkdown).SkipIds(skipIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).FilterTagIds(filterTagIds).FeaturesEnabledOnProject(featuresEnabledOnProject).CompanyId(companyId).CategoryId(categoryId).CatId(catId).Execute()

Get Projects (APICall_GET_projects)

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
    updatedAfterDateTime := "updatedAfterDateTime_example" // string |  (optional)
    updatedAfterDate := "updatedAfterDate_example" // string |  (optional)
    status := "status_example" // string | project filters (optional)
    searchTerm := "searchTerm_example" // string |  (optional)
    projectStatus := "projectStatus_example" // string |  (optional)
    orderMode := "orderMode_example" // string |  (optional)
    orderBy := "orderBy_example" // string |  (optional)
    get := "get_example" // string |  (optional)
    firstLetter := "firstLetter_example" // string |  (optional)
    dataSet := "dataSet_example" // string |  (optional)
    createdAfterDateTime := "createdAfterDateTime_example" // string |  (optional)
    createdAfterDate := "createdAfterDate_example" // string |  (optional)
    completedBeforeDateTime := "completedBeforeDateTime_example" // string |  (optional)
    completedAfterDateTime := "completedAfterDateTime_example" // string |  (optional)
    userId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    searchCompany := true // bool |  (optional)
    returnLetters := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    onlyArchivedProjects := true // bool |  (optional)
    matchAllTags := true // bool |  (optional)
    matchAllProjectTags := true // bool |  (optional)
    includeUpdates := true // bool |  (optional)
    includeTopPeople := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    includeProjectOwner := true // bool |  (optional)
    includePeople := true // bool |  (optional)
    includeArchivedProjects := true // bool |  (optional)
    hideDesc := true // bool |  (optional)
    getTruePermissions := true // bool |  (optional)
    getPermissions := true // bool |  (optional)
    getNotificationSettings := true // bool |  (optional)
    getDeleted := true // bool |  (optional)
    getCounts := true // bool |  (optional)
    getCategoryPath := true // bool |  (optional)
    getAll := true // bool |  (optional)
    getActivePages := true // bool |  (optional)
    formatMarkdown := true // bool |  (optional)
    skipIds := []int32{int32(123)} // []int32 |  (optional)
    projectTagIds := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    projectHealths := []int32{int32(123)} // []int32 |  (optional)
    projectHealth := []int32{int32(123)} // []int32 |  (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 |  (optional)
    filterTagIds := []int32{int32(123)} // []int32 |  (optional)
    featuresEnabledOnProject := []string{"Inner_example"} // []string |  (optional)
    companyId := []int32{int32(123)} // []int32 |  (optional)
    categoryId := []int32{int32(123)} // []int32 |  (optional)
    catId := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV1Projects(context.Background()).UpdatedAfterDateTime(updatedAfterDateTime).UpdatedAfterDate(updatedAfterDate).Status(status).SearchTerm(searchTerm).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).Get(get).FirstLetter(firstLetter).DataSet(dataSet).CreatedAfterDateTime(createdAfterDateTime).CreatedAfterDate(createdAfterDate).CompletedBeforeDateTime(completedBeforeDateTime).CompletedAfterDateTime(completedAfterDateTime).UserId(userId).PageSize(pageSize).Page(page).SearchCompany(searchCompany).ReturnLetters(returnLetters).OnlyStarredProjects(onlyStarredProjects).OnlyArchivedProjects(onlyArchivedProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).IncludeUpdates(includeUpdates).IncludeTopPeople(includeTopPeople).IncludeTags(includeTags).IncludeProjectOwner(includeProjectOwner).IncludePeople(includePeople).IncludeArchivedProjects(includeArchivedProjects).HideDesc(hideDesc).GetTruePermissions(getTruePermissions).GetPermissions(getPermissions).GetNotificationSettings(getNotificationSettings).GetDeleted(getDeleted).GetCounts(getCounts).GetCategoryPath(getCategoryPath).GetAll(getAll).GetActivePages(getActivePages).FormatMarkdown(formatMarkdown).SkipIds(skipIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).FilterTagIds(filterTagIds).FeaturesEnabledOnProject(featuresEnabledOnProject).CompanyId(companyId).CategoryId(categoryId).CatId(catId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV1Projects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV1ProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfterDateTime** | **string** |  | 
 **updatedAfterDate** | **string** |  | 
 **status** | **string** | project filters | 
 **searchTerm** | **string** |  | 
 **projectStatus** | **string** |  | 
 **orderMode** | **string** |  | 
 **orderBy** | **string** |  | 
 **get** | **string** |  | 
 **firstLetter** | **string** |  | 
 **dataSet** | **string** |  | 
 **createdAfterDateTime** | **string** |  | 
 **createdAfterDate** | **string** |  | 
 **completedBeforeDateTime** | **string** |  | 
 **completedAfterDateTime** | **string** |  | 
 **userId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **searchCompany** | **bool** |  | 
 **returnLetters** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **onlyArchivedProjects** | **bool** |  | 
 **matchAllTags** | **bool** |  | 
 **matchAllProjectTags** | **bool** |  | 
 **includeUpdates** | **bool** |  | 
 **includeTopPeople** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **includeProjectOwner** | **bool** |  | 
 **includePeople** | **bool** |  | 
 **includeArchivedProjects** | **bool** |  | 
 **hideDesc** | **bool** |  | 
 **getTruePermissions** | **bool** |  | 
 **getPermissions** | **bool** |  | 
 **getNotificationSettings** | **bool** |  | 
 **getDeleted** | **bool** |  | 
 **getCounts** | **bool** |  | 
 **getCategoryPath** | **bool** |  | 
 **getAll** | **bool** |  | 
 **getActivePages** | **bool** |  | 
 **formatMarkdown** | **bool** |  | 
 **skipIds** | **[]int32** |  | 
 **projectTagIds** | **[]int32** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **projectHealths** | **[]int32** |  | 
 **projectHealth** | **[]int32** |  | 
 **projectCompanyIds** | **[]int32** |  | 
 **projectCategoryIds** | **[]int32** |  | 
 **filterTagIds** | **[]int32** |  | 
 **featuresEnabledOnProject** | **[]string** |  | 
 **companyId** | **[]int32** |  | 
 **categoryId** | **[]int32** |  | 
 **catId** | **[]int32** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV1ProjectsprojectIdTasklists

> GETProjectsApiV1ProjectsprojectIdTasklists(ctx, projectId).UpdatedAfterDate(updatedAfterDate).Status(status).Status2(status2).SortBy(sortBy).SearchTerm(searchTerm).ProjectStatus(projectStatus).FilterText(filterText).Filter(filter).DataSet(dataSet).UserId(userId).TasklistId(tasklistId).ResponsiblePartyId(responsiblePartyId).PageSize(pageSize).Page(page).OnlyUnattachedExceptMilestoneId(onlyUnattachedExceptMilestoneId).ShowPrivate(showPrivate).ShowMilestones(showMilestones).ShowDeleted(showDeleted).ShowCompleted(showCompleted).OnlyStarredProjects(onlyStarredProjects).OnlyArchivedProjects(onlyArchivedProjects).MatchAllProjectTags(matchAllProjectTags).IncludeTags(includeTags).IncludeCustomFields(includeCustomFields).IncludeArchivedProjects(includeArchivedProjects).GetTasks(getTasks).GetOverdueCount(getOverdueCount).GetNewTaskDefaults(getNewTaskDefaults).GetEmptyLists(getEmptyLists).GetDLMs(getDLMs).GetCompletedCount(getCompletedCount).GetCategoryPath(getCategoryPath).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId2(projectId2).ProjectHealths(projectHealths).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Execute()

Get tasklists

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
    updatedAfterDate := "updatedAfterDate_example" // string |  (optional)
    status := "status_example" // string |  (optional)
    status2 := "status_example" // string |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    searchTerm := "searchTerm_example" // string |  (optional)
    projectStatus := "projectStatus_example" // string |  (optional)
    filterText := "filterText_example" // string |  (optional)
    filter := "filter_example" // string |  (optional)
    dataSet := "dataSet_example" // string |  (optional)
    userId := int32(56) // int32 |  (optional)
    tasklistId := int32(56) // int32 |  (optional)
    responsiblePartyId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    onlyUnattachedExceptMilestoneId := int32(56) // int32 |  (optional)
    showPrivate := true // bool |  (optional)
    showMilestones := true // bool |  (optional)
    showDeleted := true // bool |  (optional)
    showCompleted := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    onlyArchivedProjects := true // bool |  (optional)
    matchAllProjectTags := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    includeCustomFields := true // bool |  (optional)
    includeArchivedProjects := true // bool |  (optional)
    getTasks := true // bool |  (optional)
    getOverdueCount := true // bool |  (optional)
    getNewTaskDefaults := true // bool |  (optional)
    getEmptyLists := true // bool |  (optional)
    getDLMs := true // bool |  (optional)
    getCompletedCount := true // bool |  (optional)
    getCategoryPath := true // bool |  (optional)
    projectTagIds := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    projectIds := []int32{int32(123)} // []int32 |  (optional)
    projectId2 := []int32{int32(123)} // []int32 | project filters (optional)
    projectHealths := []int32{int32(123)} // []int32 |  (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV1ProjectsprojectIdTasklists(context.Background(), projectId).UpdatedAfterDate(updatedAfterDate).Status(status).Status2(status2).SortBy(sortBy).SearchTerm(searchTerm).ProjectStatus(projectStatus).FilterText(filterText).Filter(filter).DataSet(dataSet).UserId(userId).TasklistId(tasklistId).ResponsiblePartyId(responsiblePartyId).PageSize(pageSize).Page(page).OnlyUnattachedExceptMilestoneId(onlyUnattachedExceptMilestoneId).ShowPrivate(showPrivate).ShowMilestones(showMilestones).ShowDeleted(showDeleted).ShowCompleted(showCompleted).OnlyStarredProjects(onlyStarredProjects).OnlyArchivedProjects(onlyArchivedProjects).MatchAllProjectTags(matchAllProjectTags).IncludeTags(includeTags).IncludeCustomFields(includeCustomFields).IncludeArchivedProjects(includeArchivedProjects).GetTasks(getTasks).GetOverdueCount(getOverdueCount).GetNewTaskDefaults(getNewTaskDefaults).GetEmptyLists(getEmptyLists).GetDLMs(getDLMs).GetCompletedCount(getCompletedCount).GetCategoryPath(getCategoryPath).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId2(projectId2).ProjectHealths(projectHealths).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV1ProjectsprojectIdTasklists``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETProjectsApiV1ProjectsprojectIdTasklistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedAfterDate** | **string** |  | 
 **status** | **string** |  | 
 **status2** | **string** |  | 
 **sortBy** | **string** |  | 
 **searchTerm** | **string** |  | 
 **projectStatus** | **string** |  | 
 **filterText** | **string** |  | 
 **filter** | **string** |  | 
 **dataSet** | **string** |  | 
 **userId** | **int32** |  | 
 **tasklistId** | **int32** |  | 
 **responsiblePartyId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **onlyUnattachedExceptMilestoneId** | **int32** |  | 
 **showPrivate** | **bool** |  | 
 **showMilestones** | **bool** |  | 
 **showDeleted** | **bool** |  | 
 **showCompleted** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **onlyArchivedProjects** | **bool** |  | 
 **matchAllProjectTags** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **includeCustomFields** | **bool** |  | 
 **includeArchivedProjects** | **bool** |  | 
 **getTasks** | **bool** |  | 
 **getOverdueCount** | **bool** |  | 
 **getNewTaskDefaults** | **bool** |  | 
 **getEmptyLists** | **bool** |  | 
 **getDLMs** | **bool** |  | 
 **getCompletedCount** | **bool** |  | 
 **getCategoryPath** | **bool** |  | 
 **projectTagIds** | **[]int32** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **projectIds** | **[]int32** |  | 
 **projectId2** | **[]int32** | project filters | 
 **projectHealths** | **[]int32** |  | 
 **projectCompanyIds** | **[]int32** |  | 
 **projectCategoryIds** | **[]int32** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV1ProjectsprojectIdTasks

> GETProjectsApiV1ProjectsprojectIdTasks(ctx, projectId).UpdatedAfterDate(updatedAfterDate).Today(today).TagIds(tagIds).StartDate(startDate).Sort(sort).ResponsiblePartyIds(responsiblePartyIds).ProjectStatus(projectStatus).Include(include).FollowedByUserIds(followedByUserIds).Filter(filter).EndDate(endDate).DataSet(dataSet).CreatorIds(creatorIds).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TasklistId(tasklistId).TaskId(taskId).ProjectCategoryId(projectCategoryId).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).IncludeTaskId(includeTaskId).CompanyId(companyId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyStarredProjects(onlyStarredProjects).OnlyArchivedProjects(onlyArchivedProjects).NestSubTasks(nestSubTasks).MatchAllProjectTags(matchAllProjectTags).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompletedPredecessors(includeCompletedPredecessors).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).CountOnly(countOnly).AllowTemplateTasks(allowTemplateTasks).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId2(projectId2).ProjectHealths(projectHealths).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Execute()

Will return the tasks

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
    updatedAfterDate := "updatedAfterDate_example" // string |  (optional)
    today := "today_example" // string |  (optional)
    tagIds := "tagIds_example" // string |  (optional)
    startDate := "startDate_example" // string |  (optional)
    sort := "sort_example" // string |  (optional)
    responsiblePartyIds := "responsiblePartyIds_example" // string |  (optional)
    projectStatus := "projectStatus_example" // string |  (optional)
    include := "include_example" // string |  (optional)
    followedByUserIds := "followedByUserIds_example" // string |  (optional)
    filter := "filter_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    dataSet := "dataSet_example" // string |  (optional)
    creatorIds := "creatorIds_example" // string |  (optional)
    createdBeforeDate := "createdBeforeDate_example" // string |  (optional)
    createdAfterDate := "createdAfterDate_example" // string |  (optional)
    completedBeforeDate := "completedBeforeDate_example" // string |  (optional)
    completedAfterDate := "completedAfterDate_example" // string |  (optional)
    callback := "callback_example" // string |  (optional)
    tasklistId := int32(56) // int32 |  (optional)
    taskId := int32(56) // int32 |  (optional)
    projectCategoryId := int32(56) // int32 |  (optional)
    parentTaskId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    includeTaskId := int32(56) // int32 |  (optional)
    companyId := int32(56) // int32 |  (optional)
    useStartDatesForTodaysTasks := true // bool |  (optional)
    useAllProjects := true // bool |  (optional)
    starredProjectsOnly := true // bool |  (optional)
    showDeleted := true // bool |  (optional)
    showCompletedLists := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    onlyArchivedProjects := true // bool |  (optional)
    nestSubTasks := true // bool |  (optional)
    matchAllProjectTags := true // bool |  (optional)
    includeUntaggedTasks := true // bool |  (optional)
    includeToday := true // bool |  (optional)
    includeTasksWithoutDueDates := true // bool |  (optional)
    includeTasksFromDeletedLists := true // bool |  (optional)
    includeReminders := true // bool |  (optional)
    includeLoggedTime := true // bool |  (optional)
    includeCompletedTasks := true // bool |  (optional)
    includeCompletedSubtasks := true // bool |  (optional)
    includeCompletedPredecessors := true // bool |  (optional)
    includeArchivedProjects := true // bool |  (optional)
    ignoreStartDates := true // bool |  (optional)
    getSubTasks := true // bool |  (optional)
    getFiles := true // bool |  (optional)
    countOnly := true // bool |  (optional)
    allowTemplateTasks := true // bool |  (optional)
    projectTagIds := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    projectIds := []int32{int32(123)} // []int32 |  (optional)
    projectId2 := []int32{int32(123)} // []int32 | project *filters (optional)
    projectHealths := []int32{int32(123)} // []int32 |  (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV1ProjectsprojectIdTasks(context.Background(), projectId).UpdatedAfterDate(updatedAfterDate).Today(today).TagIds(tagIds).StartDate(startDate).Sort(sort).ResponsiblePartyIds(responsiblePartyIds).ProjectStatus(projectStatus).Include(include).FollowedByUserIds(followedByUserIds).Filter(filter).EndDate(endDate).DataSet(dataSet).CreatorIds(creatorIds).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TasklistId(tasklistId).TaskId(taskId).ProjectCategoryId(projectCategoryId).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).IncludeTaskId(includeTaskId).CompanyId(companyId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyStarredProjects(onlyStarredProjects).OnlyArchivedProjects(onlyArchivedProjects).NestSubTasks(nestSubTasks).MatchAllProjectTags(matchAllProjectTags).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompletedPredecessors(includeCompletedPredecessors).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).CountOnly(countOnly).AllowTemplateTasks(allowTemplateTasks).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId2(projectId2).ProjectHealths(projectHealths).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV1ProjectsprojectIdTasks``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETProjectsApiV1ProjectsprojectIdTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedAfterDate** | **string** |  | 
 **today** | **string** |  | 
 **tagIds** | **string** |  | 
 **startDate** | **string** |  | 
 **sort** | **string** |  | 
 **responsiblePartyIds** | **string** |  | 
 **projectStatus** | **string** |  | 
 **include** | **string** |  | 
 **followedByUserIds** | **string** |  | 
 **filter** | **string** |  | 
 **endDate** | **string** |  | 
 **dataSet** | **string** |  | 
 **creatorIds** | **string** |  | 
 **createdBeforeDate** | **string** |  | 
 **createdAfterDate** | **string** |  | 
 **completedBeforeDate** | **string** |  | 
 **completedAfterDate** | **string** |  | 
 **callback** | **string** |  | 
 **tasklistId** | **int32** |  | 
 **taskId** | **int32** |  | 
 **projectCategoryId** | **int32** |  | 
 **parentTaskId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **includeTaskId** | **int32** |  | 
 **companyId** | **int32** |  | 
 **useStartDatesForTodaysTasks** | **bool** |  | 
 **useAllProjects** | **bool** |  | 
 **starredProjectsOnly** | **bool** |  | 
 **showDeleted** | **bool** |  | 
 **showCompletedLists** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **onlyArchivedProjects** | **bool** |  | 
 **nestSubTasks** | **bool** |  | 
 **matchAllProjectTags** | **bool** |  | 
 **includeUntaggedTasks** | **bool** |  | 
 **includeToday** | **bool** |  | 
 **includeTasksWithoutDueDates** | **bool** |  | 
 **includeTasksFromDeletedLists** | **bool** |  | 
 **includeReminders** | **bool** |  | 
 **includeLoggedTime** | **bool** |  | 
 **includeCompletedTasks** | **bool** |  | 
 **includeCompletedSubtasks** | **bool** |  | 
 **includeCompletedPredecessors** | **bool** |  | 
 **includeArchivedProjects** | **bool** |  | 
 **ignoreStartDates** | **bool** |  | 
 **getSubTasks** | **bool** |  | 
 **getFiles** | **bool** |  | 
 **countOnly** | **bool** |  | 
 **allowTemplateTasks** | **bool** |  | 
 **projectTagIds** | **[]int32** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **projectIds** | **[]int32** |  | 
 **projectId2** | **[]int32** | project *filters | 
 **projectHealths** | **[]int32** |  | 
 **projectCompanyIds** | **[]int32** |  | 
 **projectCategoryIds** | **[]int32** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV1Tasklists

> GETProjectsApiV1Tasklists(ctx).UpdatedAfterDate(updatedAfterDate).Status(status).Status2(status2).SortBy(sortBy).SearchTerm(searchTerm).ProjectStatus(projectStatus).FilterText(filterText).Filter(filter).DataSet(dataSet).UserId(userId).TasklistId(tasklistId).ResponsiblePartyId(responsiblePartyId).PageSize(pageSize).Page(page).OnlyUnattachedExceptMilestoneId(onlyUnattachedExceptMilestoneId).ShowPrivate(showPrivate).ShowMilestones(showMilestones).ShowDeleted(showDeleted).ShowCompleted(showCompleted).OnlyStarredProjects(onlyStarredProjects).OnlyArchivedProjects(onlyArchivedProjects).MatchAllProjectTags(matchAllProjectTags).IncludeTags(includeTags).IncludeCustomFields(includeCustomFields).IncludeArchivedProjects(includeArchivedProjects).GetTasks(getTasks).GetOverdueCount(getOverdueCount).GetNewTaskDefaults(getNewTaskDefaults).GetEmptyLists(getEmptyLists).GetDLMs(getDLMs).GetCompletedCount(getCompletedCount).GetCategoryPath(getCategoryPath).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Execute()

Get tasklists

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
    updatedAfterDate := "updatedAfterDate_example" // string |  (optional)
    status := "status_example" // string |  (optional)
    status2 := "status_example" // string |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    searchTerm := "searchTerm_example" // string |  (optional)
    projectStatus := "projectStatus_example" // string |  (optional)
    filterText := "filterText_example" // string |  (optional)
    filter := "filter_example" // string |  (optional)
    dataSet := "dataSet_example" // string |  (optional)
    userId := int32(56) // int32 |  (optional)
    tasklistId := int32(56) // int32 |  (optional)
    responsiblePartyId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    onlyUnattachedExceptMilestoneId := int32(56) // int32 |  (optional)
    showPrivate := true // bool |  (optional)
    showMilestones := true // bool |  (optional)
    showDeleted := true // bool |  (optional)
    showCompleted := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    onlyArchivedProjects := true // bool |  (optional)
    matchAllProjectTags := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    includeCustomFields := true // bool |  (optional)
    includeArchivedProjects := true // bool |  (optional)
    getTasks := true // bool |  (optional)
    getOverdueCount := true // bool |  (optional)
    getNewTaskDefaults := true // bool |  (optional)
    getEmptyLists := true // bool |  (optional)
    getDLMs := true // bool |  (optional)
    getCompletedCount := true // bool |  (optional)
    getCategoryPath := true // bool |  (optional)
    projectTagIds := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    projectIds := []int32{int32(123)} // []int32 |  (optional)
    projectId := []int32{int32(123)} // []int32 | project filters (optional)
    projectHealths := []int32{int32(123)} // []int32 |  (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV1Tasklists(context.Background()).UpdatedAfterDate(updatedAfterDate).Status(status).Status2(status2).SortBy(sortBy).SearchTerm(searchTerm).ProjectStatus(projectStatus).FilterText(filterText).Filter(filter).DataSet(dataSet).UserId(userId).TasklistId(tasklistId).ResponsiblePartyId(responsiblePartyId).PageSize(pageSize).Page(page).OnlyUnattachedExceptMilestoneId(onlyUnattachedExceptMilestoneId).ShowPrivate(showPrivate).ShowMilestones(showMilestones).ShowDeleted(showDeleted).ShowCompleted(showCompleted).OnlyStarredProjects(onlyStarredProjects).OnlyArchivedProjects(onlyArchivedProjects).MatchAllProjectTags(matchAllProjectTags).IncludeTags(includeTags).IncludeCustomFields(includeCustomFields).IncludeArchivedProjects(includeArchivedProjects).GetTasks(getTasks).GetOverdueCount(getOverdueCount).GetNewTaskDefaults(getNewTaskDefaults).GetEmptyLists(getEmptyLists).GetDLMs(getDLMs).GetCompletedCount(getCompletedCount).GetCategoryPath(getCategoryPath).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV1Tasklists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV1TasklistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfterDate** | **string** |  | 
 **status** | **string** |  | 
 **status2** | **string** |  | 
 **sortBy** | **string** |  | 
 **searchTerm** | **string** |  | 
 **projectStatus** | **string** |  | 
 **filterText** | **string** |  | 
 **filter** | **string** |  | 
 **dataSet** | **string** |  | 
 **userId** | **int32** |  | 
 **tasklistId** | **int32** |  | 
 **responsiblePartyId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **onlyUnattachedExceptMilestoneId** | **int32** |  | 
 **showPrivate** | **bool** |  | 
 **showMilestones** | **bool** |  | 
 **showDeleted** | **bool** |  | 
 **showCompleted** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **onlyArchivedProjects** | **bool** |  | 
 **matchAllProjectTags** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **includeCustomFields** | **bool** |  | 
 **includeArchivedProjects** | **bool** |  | 
 **getTasks** | **bool** |  | 
 **getOverdueCount** | **bool** |  | 
 **getNewTaskDefaults** | **bool** |  | 
 **getEmptyLists** | **bool** |  | 
 **getDLMs** | **bool** |  | 
 **getCompletedCount** | **bool** |  | 
 **getCategoryPath** | **bool** |  | 
 **projectTagIds** | **[]int32** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **projectIds** | **[]int32** |  | 
 **projectId** | **[]int32** | project filters | 
 **projectHealths** | **[]int32** |  | 
 **projectCompanyIds** | **[]int32** |  | 
 **projectCategoryIds** | **[]int32** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV1TaskliststasklistId

> GETProjectsApiV1TaskliststasklistId(ctx, tasklistId2).UpdatedAfterDate(updatedAfterDate).Status(status).Status2(status2).SortBy(sortBy).SearchTerm(searchTerm).ProjectStatus(projectStatus).FilterText(filterText).Filter(filter).DataSet(dataSet).UserId(userId).TasklistId(tasklistId).ResponsiblePartyId(responsiblePartyId).PageSize(pageSize).Page(page).OnlyUnattachedExceptMilestoneId(onlyUnattachedExceptMilestoneId).ShowPrivate(showPrivate).ShowMilestones(showMilestones).ShowDeleted(showDeleted).ShowCompleted(showCompleted).OnlyStarredProjects(onlyStarredProjects).OnlyArchivedProjects(onlyArchivedProjects).MatchAllProjectTags(matchAllProjectTags).IncludeTags(includeTags).IncludeCustomFields(includeCustomFields).IncludeArchivedProjects(includeArchivedProjects).GetTasks(getTasks).GetOverdueCount(getOverdueCount).GetNewTaskDefaults(getNewTaskDefaults).GetEmptyLists(getEmptyLists).GetDLMs(getDLMs).GetCompletedCount(getCompletedCount).GetCategoryPath(getCategoryPath).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Execute()

Get tasklists

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
    tasklistId2 := int32(56) // int32 | 
    updatedAfterDate := "updatedAfterDate_example" // string |  (optional)
    status := "status_example" // string |  (optional)
    status2 := "status_example" // string |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    searchTerm := "searchTerm_example" // string |  (optional)
    projectStatus := "projectStatus_example" // string |  (optional)
    filterText := "filterText_example" // string |  (optional)
    filter := "filter_example" // string |  (optional)
    dataSet := "dataSet_example" // string |  (optional)
    userId := int32(56) // int32 |  (optional)
    tasklistId := int32(56) // int32 |  (optional)
    responsiblePartyId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    onlyUnattachedExceptMilestoneId := int32(56) // int32 |  (optional)
    showPrivate := true // bool |  (optional)
    showMilestones := true // bool |  (optional)
    showDeleted := true // bool |  (optional)
    showCompleted := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    onlyArchivedProjects := true // bool |  (optional)
    matchAllProjectTags := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    includeCustomFields := true // bool |  (optional)
    includeArchivedProjects := true // bool |  (optional)
    getTasks := true // bool |  (optional)
    getOverdueCount := true // bool |  (optional)
    getNewTaskDefaults := true // bool |  (optional)
    getEmptyLists := true // bool |  (optional)
    getDLMs := true // bool |  (optional)
    getCompletedCount := true // bool |  (optional)
    getCategoryPath := true // bool |  (optional)
    projectTagIds := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    projectIds := []int32{int32(123)} // []int32 |  (optional)
    projectId := []int32{int32(123)} // []int32 | project filters (optional)
    projectHealths := []int32{int32(123)} // []int32 |  (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV1TaskliststasklistId(context.Background(), tasklistId2).UpdatedAfterDate(updatedAfterDate).Status(status).Status2(status2).SortBy(sortBy).SearchTerm(searchTerm).ProjectStatus(projectStatus).FilterText(filterText).Filter(filter).DataSet(dataSet).UserId(userId).TasklistId(tasklistId).ResponsiblePartyId(responsiblePartyId).PageSize(pageSize).Page(page).OnlyUnattachedExceptMilestoneId(onlyUnattachedExceptMilestoneId).ShowPrivate(showPrivate).ShowMilestones(showMilestones).ShowDeleted(showDeleted).ShowCompleted(showCompleted).OnlyStarredProjects(onlyStarredProjects).OnlyArchivedProjects(onlyArchivedProjects).MatchAllProjectTags(matchAllProjectTags).IncludeTags(includeTags).IncludeCustomFields(includeCustomFields).IncludeArchivedProjects(includeArchivedProjects).GetTasks(getTasks).GetOverdueCount(getOverdueCount).GetNewTaskDefaults(getNewTaskDefaults).GetEmptyLists(getEmptyLists).GetDLMs(getDLMs).GetCompletedCount(getCompletedCount).GetCategoryPath(getCategoryPath).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV1TaskliststasklistId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tasklistId2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV1TaskliststasklistIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedAfterDate** | **string** |  | 
 **status** | **string** |  | 
 **status2** | **string** |  | 
 **sortBy** | **string** |  | 
 **searchTerm** | **string** |  | 
 **projectStatus** | **string** |  | 
 **filterText** | **string** |  | 
 **filter** | **string** |  | 
 **dataSet** | **string** |  | 
 **userId** | **int32** |  | 
 **tasklistId** | **int32** |  | 
 **responsiblePartyId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **onlyUnattachedExceptMilestoneId** | **int32** |  | 
 **showPrivate** | **bool** |  | 
 **showMilestones** | **bool** |  | 
 **showDeleted** | **bool** |  | 
 **showCompleted** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **onlyArchivedProjects** | **bool** |  | 
 **matchAllProjectTags** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **includeCustomFields** | **bool** |  | 
 **includeArchivedProjects** | **bool** |  | 
 **getTasks** | **bool** |  | 
 **getOverdueCount** | **bool** |  | 
 **getNewTaskDefaults** | **bool** |  | 
 **getEmptyLists** | **bool** |  | 
 **getDLMs** | **bool** |  | 
 **getCompletedCount** | **bool** |  | 
 **getCategoryPath** | **bool** |  | 
 **projectTagIds** | **[]int32** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **projectIds** | **[]int32** |  | 
 **projectId** | **[]int32** | project filters | 
 **projectHealths** | **[]int32** |  | 
 **projectCompanyIds** | **[]int32** |  | 
 **projectCategoryIds** | **[]int32** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV1TaskliststasklistIdTasks

> GETProjectsApiV1TaskliststasklistIdTasks(ctx, tasklistId).UpdatedAfterDate(updatedAfterDate).Today(today).TagIds(tagIds).StartDate(startDate).Sort(sort).ResponsiblePartyIds(responsiblePartyIds).ProjectStatus(projectStatus).Include(include).FollowedByUserIds(followedByUserIds).Filter(filter).EndDate(endDate).DataSet(dataSet).CreatorIds(creatorIds).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TasklistId2(tasklistId2).TaskId(taskId).ProjectCategoryId(projectCategoryId).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).IncludeTaskId(includeTaskId).CompanyId(companyId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyStarredProjects(onlyStarredProjects).OnlyArchivedProjects(onlyArchivedProjects).NestSubTasks(nestSubTasks).MatchAllProjectTags(matchAllProjectTags).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompletedPredecessors(includeCompletedPredecessors).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).CountOnly(countOnly).AllowTemplateTasks(allowTemplateTasks).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Execute()

Will return the tasks

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
    tasklistId := int32(56) // int32 | 
    updatedAfterDate := "updatedAfterDate_example" // string |  (optional)
    today := "today_example" // string |  (optional)
    tagIds := "tagIds_example" // string |  (optional)
    startDate := "startDate_example" // string |  (optional)
    sort := "sort_example" // string |  (optional)
    responsiblePartyIds := "responsiblePartyIds_example" // string |  (optional)
    projectStatus := "projectStatus_example" // string |  (optional)
    include := "include_example" // string |  (optional)
    followedByUserIds := "followedByUserIds_example" // string |  (optional)
    filter := "filter_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    dataSet := "dataSet_example" // string |  (optional)
    creatorIds := "creatorIds_example" // string |  (optional)
    createdBeforeDate := "createdBeforeDate_example" // string |  (optional)
    createdAfterDate := "createdAfterDate_example" // string |  (optional)
    completedBeforeDate := "completedBeforeDate_example" // string |  (optional)
    completedAfterDate := "completedAfterDate_example" // string |  (optional)
    callback := "callback_example" // string |  (optional)
    tasklistId2 := int32(56) // int32 |  (optional)
    taskId := int32(56) // int32 |  (optional)
    projectCategoryId := int32(56) // int32 |  (optional)
    parentTaskId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    includeTaskId := int32(56) // int32 |  (optional)
    companyId := int32(56) // int32 |  (optional)
    useStartDatesForTodaysTasks := true // bool |  (optional)
    useAllProjects := true // bool |  (optional)
    starredProjectsOnly := true // bool |  (optional)
    showDeleted := true // bool |  (optional)
    showCompletedLists := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    onlyArchivedProjects := true // bool |  (optional)
    nestSubTasks := true // bool |  (optional)
    matchAllProjectTags := true // bool |  (optional)
    includeUntaggedTasks := true // bool |  (optional)
    includeToday := true // bool |  (optional)
    includeTasksWithoutDueDates := true // bool |  (optional)
    includeTasksFromDeletedLists := true // bool |  (optional)
    includeReminders := true // bool |  (optional)
    includeLoggedTime := true // bool |  (optional)
    includeCompletedTasks := true // bool |  (optional)
    includeCompletedSubtasks := true // bool |  (optional)
    includeCompletedPredecessors := true // bool |  (optional)
    includeArchivedProjects := true // bool |  (optional)
    ignoreStartDates := true // bool |  (optional)
    getSubTasks := true // bool |  (optional)
    getFiles := true // bool |  (optional)
    countOnly := true // bool |  (optional)
    allowTemplateTasks := true // bool |  (optional)
    projectTagIds := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    projectIds := []int32{int32(123)} // []int32 |  (optional)
    projectId := []int32{int32(123)} // []int32 | project *filters (optional)
    projectHealths := []int32{int32(123)} // []int32 |  (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV1TaskliststasklistIdTasks(context.Background(), tasklistId).UpdatedAfterDate(updatedAfterDate).Today(today).TagIds(tagIds).StartDate(startDate).Sort(sort).ResponsiblePartyIds(responsiblePartyIds).ProjectStatus(projectStatus).Include(include).FollowedByUserIds(followedByUserIds).Filter(filter).EndDate(endDate).DataSet(dataSet).CreatorIds(creatorIds).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TasklistId2(tasklistId2).TaskId(taskId).ProjectCategoryId(projectCategoryId).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).IncludeTaskId(includeTaskId).CompanyId(companyId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyStarredProjects(onlyStarredProjects).OnlyArchivedProjects(onlyArchivedProjects).NestSubTasks(nestSubTasks).MatchAllProjectTags(matchAllProjectTags).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompletedPredecessors(includeCompletedPredecessors).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).CountOnly(countOnly).AllowTemplateTasks(allowTemplateTasks).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV1TaskliststasklistIdTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tasklistId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV1TaskliststasklistIdTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedAfterDate** | **string** |  | 
 **today** | **string** |  | 
 **tagIds** | **string** |  | 
 **startDate** | **string** |  | 
 **sort** | **string** |  | 
 **responsiblePartyIds** | **string** |  | 
 **projectStatus** | **string** |  | 
 **include** | **string** |  | 
 **followedByUserIds** | **string** |  | 
 **filter** | **string** |  | 
 **endDate** | **string** |  | 
 **dataSet** | **string** |  | 
 **creatorIds** | **string** |  | 
 **createdBeforeDate** | **string** |  | 
 **createdAfterDate** | **string** |  | 
 **completedBeforeDate** | **string** |  | 
 **completedAfterDate** | **string** |  | 
 **callback** | **string** |  | 
 **tasklistId2** | **int32** |  | 
 **taskId** | **int32** |  | 
 **projectCategoryId** | **int32** |  | 
 **parentTaskId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **includeTaskId** | **int32** |  | 
 **companyId** | **int32** |  | 
 **useStartDatesForTodaysTasks** | **bool** |  | 
 **useAllProjects** | **bool** |  | 
 **starredProjectsOnly** | **bool** |  | 
 **showDeleted** | **bool** |  | 
 **showCompletedLists** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **onlyArchivedProjects** | **bool** |  | 
 **nestSubTasks** | **bool** |  | 
 **matchAllProjectTags** | **bool** |  | 
 **includeUntaggedTasks** | **bool** |  | 
 **includeToday** | **bool** |  | 
 **includeTasksWithoutDueDates** | **bool** |  | 
 **includeTasksFromDeletedLists** | **bool** |  | 
 **includeReminders** | **bool** |  | 
 **includeLoggedTime** | **bool** |  | 
 **includeCompletedTasks** | **bool** |  | 
 **includeCompletedSubtasks** | **bool** |  | 
 **includeCompletedPredecessors** | **bool** |  | 
 **includeArchivedProjects** | **bool** |  | 
 **ignoreStartDates** | **bool** |  | 
 **getSubTasks** | **bool** |  | 
 **getFiles** | **bool** |  | 
 **countOnly** | **bool** |  | 
 **allowTemplateTasks** | **bool** |  | 
 **projectTagIds** | **[]int32** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **projectIds** | **[]int32** |  | 
 **projectId** | **[]int32** | project *filters | 
 **projectHealths** | **[]int32** |  | 
 **projectCompanyIds** | **[]int32** |  | 
 **projectCategoryIds** | **[]int32** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV1Tasks

> GETProjectsApiV1Tasks(ctx).UpdatedAfterDate(updatedAfterDate).Today(today).TagIds(tagIds).StartDate(startDate).Sort(sort).ResponsiblePartyIds(responsiblePartyIds).ProjectStatus(projectStatus).Include(include).FollowedByUserIds(followedByUserIds).Filter(filter).EndDate(endDate).DataSet(dataSet).CreatorIds(creatorIds).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TasklistId(tasklistId).TaskId(taskId).ProjectCategoryId(projectCategoryId).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).IncludeTaskId(includeTaskId).CompanyId(companyId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyStarredProjects(onlyStarredProjects).OnlyArchivedProjects(onlyArchivedProjects).NestSubTasks(nestSubTasks).MatchAllProjectTags(matchAllProjectTags).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompletedPredecessors(includeCompletedPredecessors).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).CountOnly(countOnly).AllowTemplateTasks(allowTemplateTasks).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Execute()

Will return the tasks

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
    updatedAfterDate := "updatedAfterDate_example" // string |  (optional)
    today := "today_example" // string |  (optional)
    tagIds := "tagIds_example" // string |  (optional)
    startDate := "startDate_example" // string |  (optional)
    sort := "sort_example" // string |  (optional)
    responsiblePartyIds := "responsiblePartyIds_example" // string |  (optional)
    projectStatus := "projectStatus_example" // string |  (optional)
    include := "include_example" // string |  (optional)
    followedByUserIds := "followedByUserIds_example" // string |  (optional)
    filter := "filter_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    dataSet := "dataSet_example" // string |  (optional)
    creatorIds := "creatorIds_example" // string |  (optional)
    createdBeforeDate := "createdBeforeDate_example" // string |  (optional)
    createdAfterDate := "createdAfterDate_example" // string |  (optional)
    completedBeforeDate := "completedBeforeDate_example" // string |  (optional)
    completedAfterDate := "completedAfterDate_example" // string |  (optional)
    callback := "callback_example" // string |  (optional)
    tasklistId := int32(56) // int32 |  (optional)
    taskId := int32(56) // int32 |  (optional)
    projectCategoryId := int32(56) // int32 |  (optional)
    parentTaskId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    includeTaskId := int32(56) // int32 |  (optional)
    companyId := int32(56) // int32 |  (optional)
    useStartDatesForTodaysTasks := true // bool |  (optional)
    useAllProjects := true // bool |  (optional)
    starredProjectsOnly := true // bool |  (optional)
    showDeleted := true // bool |  (optional)
    showCompletedLists := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    onlyArchivedProjects := true // bool |  (optional)
    nestSubTasks := true // bool |  (optional)
    matchAllProjectTags := true // bool |  (optional)
    includeUntaggedTasks := true // bool |  (optional)
    includeToday := true // bool |  (optional)
    includeTasksWithoutDueDates := true // bool |  (optional)
    includeTasksFromDeletedLists := true // bool |  (optional)
    includeReminders := true // bool |  (optional)
    includeLoggedTime := true // bool |  (optional)
    includeCompletedTasks := true // bool |  (optional)
    includeCompletedSubtasks := true // bool |  (optional)
    includeCompletedPredecessors := true // bool |  (optional)
    includeArchivedProjects := true // bool |  (optional)
    ignoreStartDates := true // bool |  (optional)
    getSubTasks := true // bool |  (optional)
    getFiles := true // bool |  (optional)
    countOnly := true // bool |  (optional)
    allowTemplateTasks := true // bool |  (optional)
    projectTagIds := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    projectIds := []int32{int32(123)} // []int32 |  (optional)
    projectId := []int32{int32(123)} // []int32 | project *filters (optional)
    projectHealths := []int32{int32(123)} // []int32 |  (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV1Tasks(context.Background()).UpdatedAfterDate(updatedAfterDate).Today(today).TagIds(tagIds).StartDate(startDate).Sort(sort).ResponsiblePartyIds(responsiblePartyIds).ProjectStatus(projectStatus).Include(include).FollowedByUserIds(followedByUserIds).Filter(filter).EndDate(endDate).DataSet(dataSet).CreatorIds(creatorIds).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TasklistId(tasklistId).TaskId(taskId).ProjectCategoryId(projectCategoryId).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).IncludeTaskId(includeTaskId).CompanyId(companyId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyStarredProjects(onlyStarredProjects).OnlyArchivedProjects(onlyArchivedProjects).NestSubTasks(nestSubTasks).MatchAllProjectTags(matchAllProjectTags).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompletedPredecessors(includeCompletedPredecessors).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).CountOnly(countOnly).AllowTemplateTasks(allowTemplateTasks).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV1Tasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV1TasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfterDate** | **string** |  | 
 **today** | **string** |  | 
 **tagIds** | **string** |  | 
 **startDate** | **string** |  | 
 **sort** | **string** |  | 
 **responsiblePartyIds** | **string** |  | 
 **projectStatus** | **string** |  | 
 **include** | **string** |  | 
 **followedByUserIds** | **string** |  | 
 **filter** | **string** |  | 
 **endDate** | **string** |  | 
 **dataSet** | **string** |  | 
 **creatorIds** | **string** |  | 
 **createdBeforeDate** | **string** |  | 
 **createdAfterDate** | **string** |  | 
 **completedBeforeDate** | **string** |  | 
 **completedAfterDate** | **string** |  | 
 **callback** | **string** |  | 
 **tasklistId** | **int32** |  | 
 **taskId** | **int32** |  | 
 **projectCategoryId** | **int32** |  | 
 **parentTaskId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **includeTaskId** | **int32** |  | 
 **companyId** | **int32** |  | 
 **useStartDatesForTodaysTasks** | **bool** |  | 
 **useAllProjects** | **bool** |  | 
 **starredProjectsOnly** | **bool** |  | 
 **showDeleted** | **bool** |  | 
 **showCompletedLists** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **onlyArchivedProjects** | **bool** |  | 
 **nestSubTasks** | **bool** |  | 
 **matchAllProjectTags** | **bool** |  | 
 **includeUntaggedTasks** | **bool** |  | 
 **includeToday** | **bool** |  | 
 **includeTasksWithoutDueDates** | **bool** |  | 
 **includeTasksFromDeletedLists** | **bool** |  | 
 **includeReminders** | **bool** |  | 
 **includeLoggedTime** | **bool** |  | 
 **includeCompletedTasks** | **bool** |  | 
 **includeCompletedSubtasks** | **bool** |  | 
 **includeCompletedPredecessors** | **bool** |  | 
 **includeArchivedProjects** | **bool** |  | 
 **ignoreStartDates** | **bool** |  | 
 **getSubTasks** | **bool** |  | 
 **getFiles** | **bool** |  | 
 **countOnly** | **bool** |  | 
 **allowTemplateTasks** | **bool** |  | 
 **projectTagIds** | **[]int32** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **projectIds** | **[]int32** |  | 
 **projectId** | **[]int32** | project *filters | 
 **projectHealths** | **[]int32** |  | 
 **projectCompanyIds** | **[]int32** |  | 
 **projectCategoryIds** | **[]int32** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV1TaskstaskId

> GETProjectsApiV1TaskstaskId(ctx, taskId).UpdatedAfterDate(updatedAfterDate).Today(today).TagIds(tagIds).StartDate(startDate).Sort(sort).ResponsiblePartyIds(responsiblePartyIds).ProjectStatus(projectStatus).Include(include).FollowedByUserIds(followedByUserIds).Filter(filter).EndDate(endDate).DataSet(dataSet).CreatorIds(creatorIds).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TasklistId(tasklistId).TaskId2(taskId2).ProjectCategoryId(projectCategoryId).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).IncludeTaskId(includeTaskId).CompanyId(companyId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyStarredProjects(onlyStarredProjects).OnlyArchivedProjects(onlyArchivedProjects).NestSubTasks(nestSubTasks).MatchAllProjectTags(matchAllProjectTags).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompletedPredecessors(includeCompletedPredecessors).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).CountOnly(countOnly).AllowTemplateTasks(allowTemplateTasks).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Execute()

Will return the tasks

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
    updatedAfterDate := "updatedAfterDate_example" // string |  (optional)
    today := "today_example" // string |  (optional)
    tagIds := "tagIds_example" // string |  (optional)
    startDate := "startDate_example" // string |  (optional)
    sort := "sort_example" // string |  (optional)
    responsiblePartyIds := "responsiblePartyIds_example" // string |  (optional)
    projectStatus := "projectStatus_example" // string |  (optional)
    include := "include_example" // string |  (optional)
    followedByUserIds := "followedByUserIds_example" // string |  (optional)
    filter := "filter_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    dataSet := "dataSet_example" // string |  (optional)
    creatorIds := "creatorIds_example" // string |  (optional)
    createdBeforeDate := "createdBeforeDate_example" // string |  (optional)
    createdAfterDate := "createdAfterDate_example" // string |  (optional)
    completedBeforeDate := "completedBeforeDate_example" // string |  (optional)
    completedAfterDate := "completedAfterDate_example" // string |  (optional)
    callback := "callback_example" // string |  (optional)
    tasklistId := int32(56) // int32 |  (optional)
    taskId2 := int32(56) // int32 |  (optional)
    projectCategoryId := int32(56) // int32 |  (optional)
    parentTaskId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    includeTaskId := int32(56) // int32 |  (optional)
    companyId := int32(56) // int32 |  (optional)
    useStartDatesForTodaysTasks := true // bool |  (optional)
    useAllProjects := true // bool |  (optional)
    starredProjectsOnly := true // bool |  (optional)
    showDeleted := true // bool |  (optional)
    showCompletedLists := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    onlyArchivedProjects := true // bool |  (optional)
    nestSubTasks := true // bool |  (optional)
    matchAllProjectTags := true // bool |  (optional)
    includeUntaggedTasks := true // bool |  (optional)
    includeToday := true // bool |  (optional)
    includeTasksWithoutDueDates := true // bool |  (optional)
    includeTasksFromDeletedLists := true // bool |  (optional)
    includeReminders := true // bool |  (optional)
    includeLoggedTime := true // bool |  (optional)
    includeCompletedTasks := true // bool |  (optional)
    includeCompletedSubtasks := true // bool |  (optional)
    includeCompletedPredecessors := true // bool |  (optional)
    includeArchivedProjects := true // bool |  (optional)
    ignoreStartDates := true // bool |  (optional)
    getSubTasks := true // bool |  (optional)
    getFiles := true // bool |  (optional)
    countOnly := true // bool |  (optional)
    allowTemplateTasks := true // bool |  (optional)
    projectTagIds := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    projectIds := []int32{int32(123)} // []int32 |  (optional)
    projectId := []int32{int32(123)} // []int32 | project *filters (optional)
    projectHealths := []int32{int32(123)} // []int32 |  (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV1TaskstaskId(context.Background(), taskId).UpdatedAfterDate(updatedAfterDate).Today(today).TagIds(tagIds).StartDate(startDate).Sort(sort).ResponsiblePartyIds(responsiblePartyIds).ProjectStatus(projectStatus).Include(include).FollowedByUserIds(followedByUserIds).Filter(filter).EndDate(endDate).DataSet(dataSet).CreatorIds(creatorIds).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TasklistId(tasklistId).TaskId2(taskId2).ProjectCategoryId(projectCategoryId).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).IncludeTaskId(includeTaskId).CompanyId(companyId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyStarredProjects(onlyStarredProjects).OnlyArchivedProjects(onlyArchivedProjects).NestSubTasks(nestSubTasks).MatchAllProjectTags(matchAllProjectTags).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompletedPredecessors(includeCompletedPredecessors).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).CountOnly(countOnly).AllowTemplateTasks(allowTemplateTasks).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV1TaskstaskId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETProjectsApiV1TaskstaskIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedAfterDate** | **string** |  | 
 **today** | **string** |  | 
 **tagIds** | **string** |  | 
 **startDate** | **string** |  | 
 **sort** | **string** |  | 
 **responsiblePartyIds** | **string** |  | 
 **projectStatus** | **string** |  | 
 **include** | **string** |  | 
 **followedByUserIds** | **string** |  | 
 **filter** | **string** |  | 
 **endDate** | **string** |  | 
 **dataSet** | **string** |  | 
 **creatorIds** | **string** |  | 
 **createdBeforeDate** | **string** |  | 
 **createdAfterDate** | **string** |  | 
 **completedBeforeDate** | **string** |  | 
 **completedAfterDate** | **string** |  | 
 **callback** | **string** |  | 
 **tasklistId** | **int32** |  | 
 **taskId2** | **int32** |  | 
 **projectCategoryId** | **int32** |  | 
 **parentTaskId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **includeTaskId** | **int32** |  | 
 **companyId** | **int32** |  | 
 **useStartDatesForTodaysTasks** | **bool** |  | 
 **useAllProjects** | **bool** |  | 
 **starredProjectsOnly** | **bool** |  | 
 **showDeleted** | **bool** |  | 
 **showCompletedLists** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **onlyArchivedProjects** | **bool** |  | 
 **nestSubTasks** | **bool** |  | 
 **matchAllProjectTags** | **bool** |  | 
 **includeUntaggedTasks** | **bool** |  | 
 **includeToday** | **bool** |  | 
 **includeTasksWithoutDueDates** | **bool** |  | 
 **includeTasksFromDeletedLists** | **bool** |  | 
 **includeReminders** | **bool** |  | 
 **includeLoggedTime** | **bool** |  | 
 **includeCompletedTasks** | **bool** |  | 
 **includeCompletedSubtasks** | **bool** |  | 
 **includeCompletedPredecessors** | **bool** |  | 
 **includeArchivedProjects** | **bool** |  | 
 **ignoreStartDates** | **bool** |  | 
 **getSubTasks** | **bool** |  | 
 **getFiles** | **bool** |  | 
 **countOnly** | **bool** |  | 
 **allowTemplateTasks** | **bool** |  | 
 **projectTagIds** | **[]int32** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **projectIds** | **[]int32** |  | 
 **projectId** | **[]int32** | project *filters | 
 **projectHealths** | **[]int32** |  | 
 **projectCompanyIds** | **[]int32** |  | 
 **projectCategoryIds** | **[]int32** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV1TaskstaskIdDependents

> GETProjectsApiV1TaskstaskIdDependents(ctx, taskId).Execute()

Will return the tasks dependents tasks

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV1TaskstaskIdDependents(context.Background(), taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV1TaskstaskIdDependents``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETProjectsApiV1TaskstaskIdDependentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV1TaskstaskIdPredecessors

> GETProjectsApiV1TaskstaskIdPredecessors(ctx, taskId).Execute()

Will return the tasks predecessor tasks

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV1TaskstaskIdPredecessors(context.Background(), taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV1TaskstaskIdPredecessors``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETProjectsApiV1TaskstaskIdPredecessorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV1TaskstaskIdSubtasks

> GETProjectsApiV1TaskstaskIdSubtasks(ctx, taskId).UpdatedAfterDate(updatedAfterDate).Today(today).TagIds(tagIds).StartDate(startDate).Sort(sort).ResponsiblePartyIds(responsiblePartyIds).ProjectStatus(projectStatus).Include(include).FollowedByUserIds(followedByUserIds).Filter(filter).EndDate(endDate).DataSet(dataSet).CreatorIds(creatorIds).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TasklistId(tasklistId).TaskId2(taskId2).ProjectCategoryId(projectCategoryId).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).IncludeTaskId(includeTaskId).CompanyId(companyId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyStarredProjects(onlyStarredProjects).OnlyArchivedProjects(onlyArchivedProjects).NestSubTasks(nestSubTasks).MatchAllProjectTags(matchAllProjectTags).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompletedPredecessors(includeCompletedPredecessors).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).CountOnly(countOnly).AllowTemplateTasks(allowTemplateTasks).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Execute()

Will return the tasks

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
    updatedAfterDate := "updatedAfterDate_example" // string |  (optional)
    today := "today_example" // string |  (optional)
    tagIds := "tagIds_example" // string |  (optional)
    startDate := "startDate_example" // string |  (optional)
    sort := "sort_example" // string |  (optional)
    responsiblePartyIds := "responsiblePartyIds_example" // string |  (optional)
    projectStatus := "projectStatus_example" // string |  (optional)
    include := "include_example" // string |  (optional)
    followedByUserIds := "followedByUserIds_example" // string |  (optional)
    filter := "filter_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    dataSet := "dataSet_example" // string |  (optional)
    creatorIds := "creatorIds_example" // string |  (optional)
    createdBeforeDate := "createdBeforeDate_example" // string |  (optional)
    createdAfterDate := "createdAfterDate_example" // string |  (optional)
    completedBeforeDate := "completedBeforeDate_example" // string |  (optional)
    completedAfterDate := "completedAfterDate_example" // string |  (optional)
    callback := "callback_example" // string |  (optional)
    tasklistId := int32(56) // int32 |  (optional)
    taskId2 := int32(56) // int32 |  (optional)
    projectCategoryId := int32(56) // int32 |  (optional)
    parentTaskId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    includeTaskId := int32(56) // int32 |  (optional)
    companyId := int32(56) // int32 |  (optional)
    useStartDatesForTodaysTasks := true // bool |  (optional)
    useAllProjects := true // bool |  (optional)
    starredProjectsOnly := true // bool |  (optional)
    showDeleted := true // bool |  (optional)
    showCompletedLists := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    onlyArchivedProjects := true // bool |  (optional)
    nestSubTasks := true // bool |  (optional)
    matchAllProjectTags := true // bool |  (optional)
    includeUntaggedTasks := true // bool |  (optional)
    includeToday := true // bool |  (optional)
    includeTasksWithoutDueDates := true // bool |  (optional)
    includeTasksFromDeletedLists := true // bool |  (optional)
    includeReminders := true // bool |  (optional)
    includeLoggedTime := true // bool |  (optional)
    includeCompletedTasks := true // bool |  (optional)
    includeCompletedSubtasks := true // bool |  (optional)
    includeCompletedPredecessors := true // bool |  (optional)
    includeArchivedProjects := true // bool |  (optional)
    ignoreStartDates := true // bool |  (optional)
    getSubTasks := true // bool |  (optional)
    getFiles := true // bool |  (optional)
    countOnly := true // bool |  (optional)
    allowTemplateTasks := true // bool |  (optional)
    projectTagIds := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    projectIds := []int32{int32(123)} // []int32 |  (optional)
    projectId := []int32{int32(123)} // []int32 | project *filters (optional)
    projectHealths := []int32{int32(123)} // []int32 |  (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV1TaskstaskIdSubtasks(context.Background(), taskId).UpdatedAfterDate(updatedAfterDate).Today(today).TagIds(tagIds).StartDate(startDate).Sort(sort).ResponsiblePartyIds(responsiblePartyIds).ProjectStatus(projectStatus).Include(include).FollowedByUserIds(followedByUserIds).Filter(filter).EndDate(endDate).DataSet(dataSet).CreatorIds(creatorIds).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TasklistId(tasklistId).TaskId2(taskId2).ProjectCategoryId(projectCategoryId).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).IncludeTaskId(includeTaskId).CompanyId(companyId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyStarredProjects(onlyStarredProjects).OnlyArchivedProjects(onlyArchivedProjects).NestSubTasks(nestSubTasks).MatchAllProjectTags(matchAllProjectTags).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompletedPredecessors(includeCompletedPredecessors).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).CountOnly(countOnly).AllowTemplateTasks(allowTemplateTasks).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV1TaskstaskIdSubtasks``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETProjectsApiV1TaskstaskIdSubtasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedAfterDate** | **string** |  | 
 **today** | **string** |  | 
 **tagIds** | **string** |  | 
 **startDate** | **string** |  | 
 **sort** | **string** |  | 
 **responsiblePartyIds** | **string** |  | 
 **projectStatus** | **string** |  | 
 **include** | **string** |  | 
 **followedByUserIds** | **string** |  | 
 **filter** | **string** |  | 
 **endDate** | **string** |  | 
 **dataSet** | **string** |  | 
 **creatorIds** | **string** |  | 
 **createdBeforeDate** | **string** |  | 
 **createdAfterDate** | **string** |  | 
 **completedBeforeDate** | **string** |  | 
 **completedAfterDate** | **string** |  | 
 **callback** | **string** |  | 
 **tasklistId** | **int32** |  | 
 **taskId2** | **int32** |  | 
 **projectCategoryId** | **int32** |  | 
 **parentTaskId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **includeTaskId** | **int32** |  | 
 **companyId** | **int32** |  | 
 **useStartDatesForTodaysTasks** | **bool** |  | 
 **useAllProjects** | **bool** |  | 
 **starredProjectsOnly** | **bool** |  | 
 **showDeleted** | **bool** |  | 
 **showCompletedLists** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **onlyArchivedProjects** | **bool** |  | 
 **nestSubTasks** | **bool** |  | 
 **matchAllProjectTags** | **bool** |  | 
 **includeUntaggedTasks** | **bool** |  | 
 **includeToday** | **bool** |  | 
 **includeTasksWithoutDueDates** | **bool** |  | 
 **includeTasksFromDeletedLists** | **bool** |  | 
 **includeReminders** | **bool** |  | 
 **includeLoggedTime** | **bool** |  | 
 **includeCompletedTasks** | **bool** |  | 
 **includeCompletedSubtasks** | **bool** |  | 
 **includeCompletedPredecessors** | **bool** |  | 
 **includeArchivedProjects** | **bool** |  | 
 **ignoreStartDates** | **bool** |  | 
 **getSubTasks** | **bool** |  | 
 **getFiles** | **bool** |  | 
 **countOnly** | **bool** |  | 
 **allowTemplateTasks** | **bool** |  | 
 **projectTagIds** | **[]int32** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **projectIds** | **[]int32** |  | 
 **projectId** | **[]int32** | project *filters | 
 **projectHealths** | **[]int32** |  | 
 **projectCompanyIds** | **[]int32** |  | 
 **projectCategoryIds** | **[]int32** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV1Time

> GETProjectsApiV1Time(ctx).UpdatedAfterDate(updatedAfterDate).ToDate(toDate).SortOrder(sortOrder).SortBy(sortBy).ProjectStatus(projectStatus).InvoicedType(invoicedType).FromDate(fromDate).BillableType(billableType).TimeLogId(timeLogId).TicketId(ticketId).PageSize(pageSize).Page(page).IncludeTimelogId(includeTimelogId).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTagIds(matchAllProjectTagIds).IncludeTags(includeTags).IncludeMainTask(includeMainTask).IncludeArchivedProjects(includeArchivedProjects).GetTotals(getTotals).UserId(userId).TaskTagIds(taskTagIds).TagIds(tagIds).ProjectsFromCompanyId(projectsFromCompanyId).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ProjectCategoryId(projectCategoryId).CompanyId(companyId).Execute()

APICall_GET_time_entries

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
    updatedAfterDate := "updatedAfterDate_example" // string |  (optional)
    toDate := "toDate_example" // string |  (optional)
    sortOrder := "sortOrder_example" // string |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    projectStatus := "projectStatus_example" // string |  (optional)
    invoicedType := "invoicedType_example" // string |  (optional)
    fromDate := "fromDate_example" // string |  (optional)
    billableType := "billableType_example" // string |  (optional)
    timeLogId := int32(56) // int32 |  (optional)
    ticketId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    includeTimelogId := int32(56) // int32 |  (optional)
    showDeleted := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    matchAllTags := true // bool |  (optional)
    matchAllProjectTags := true // bool |  (optional)
    matchAllProjectTagIds := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    includeMainTask := true // bool |  (optional)
    includeArchivedProjects := true // bool |  (optional)
    getTotals := true // bool |  (optional)
    userId := []int32{int32(123)} // []int32 |  (optional)
    taskTagIds := []int32{int32(123)} // []int32 |  (optional)
    tagIds := []int32{int32(123)} // []int32 |  (optional)
    projectsFromCompanyId := []int32{int32(123)} // []int32 |  (optional)
    projectTagIds := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    projectIds := []int32{int32(123)} // []int32 |  (optional)
    projectId := []int32{int32(123)} // []int32 | project filters (optional)
    projectHealths := []int32{int32(123)} // []int32 |  (optional)
    projectHealth := []int32{int32(123)} // []int32 |  (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryId := []int32{int32(123)} // []int32 |  (optional)
    companyId := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV1Time(context.Background()).UpdatedAfterDate(updatedAfterDate).ToDate(toDate).SortOrder(sortOrder).SortBy(sortBy).ProjectStatus(projectStatus).InvoicedType(invoicedType).FromDate(fromDate).BillableType(billableType).TimeLogId(timeLogId).TicketId(ticketId).PageSize(pageSize).Page(page).IncludeTimelogId(includeTimelogId).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTagIds(matchAllProjectTagIds).IncludeTags(includeTags).IncludeMainTask(includeMainTask).IncludeArchivedProjects(includeArchivedProjects).GetTotals(getTotals).UserId(userId).TaskTagIds(taskTagIds).TagIds(tagIds).ProjectsFromCompanyId(projectsFromCompanyId).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ProjectCategoryId(projectCategoryId).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV1Time``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV1TimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfterDate** | **string** |  | 
 **toDate** | **string** |  | 
 **sortOrder** | **string** |  | 
 **sortBy** | **string** |  | 
 **projectStatus** | **string** |  | 
 **invoicedType** | **string** |  | 
 **fromDate** | **string** |  | 
 **billableType** | **string** |  | 
 **timeLogId** | **int32** |  | 
 **ticketId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **includeTimelogId** | **int32** |  | 
 **showDeleted** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **matchAllTags** | **bool** |  | 
 **matchAllProjectTags** | **bool** |  | 
 **matchAllProjectTagIds** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **includeMainTask** | **bool** |  | 
 **includeArchivedProjects** | **bool** |  | 
 **getTotals** | **bool** |  | 
 **userId** | **[]int32** |  | 
 **taskTagIds** | **[]int32** |  | 
 **tagIds** | **[]int32** |  | 
 **projectsFromCompanyId** | **[]int32** |  | 
 **projectTagIds** | **[]int32** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **projectIds** | **[]int32** |  | 
 **projectId** | **[]int32** | project filters | 
 **projectHealths** | **[]int32** |  | 
 **projectHealth** | **[]int32** |  | 
 **projectCompanyIds** | **[]int32** |  | 
 **projectCategoryIds** | **[]int32** |  | 
 **projectCategoryId** | **[]int32** |  | 
 **companyId** | **[]int32** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV1TimeEntries

> GETProjectsApiV1TimeEntries(ctx).UpdatedAfterDate(updatedAfterDate).ToDate(toDate).SortOrder(sortOrder).SortBy(sortBy).ProjectStatus(projectStatus).InvoicedType(invoicedType).FromDate(fromDate).BillableType(billableType).TimeLogId(timeLogId).TicketId(ticketId).PageSize(pageSize).Page(page).IncludeTimelogId(includeTimelogId).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTagIds(matchAllProjectTagIds).IncludeTags(includeTags).IncludeMainTask(includeMainTask).IncludeArchivedProjects(includeArchivedProjects).GetTotals(getTotals).UserId(userId).TaskTagIds(taskTagIds).TagIds(tagIds).ProjectsFromCompanyId(projectsFromCompanyId).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ProjectCategoryId(projectCategoryId).CompanyId(companyId).Execute()

APICall_GET_time_entries

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
    updatedAfterDate := "updatedAfterDate_example" // string |  (optional)
    toDate := "toDate_example" // string |  (optional)
    sortOrder := "sortOrder_example" // string |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    projectStatus := "projectStatus_example" // string |  (optional)
    invoicedType := "invoicedType_example" // string |  (optional)
    fromDate := "fromDate_example" // string |  (optional)
    billableType := "billableType_example" // string |  (optional)
    timeLogId := int32(56) // int32 |  (optional)
    ticketId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    includeTimelogId := int32(56) // int32 |  (optional)
    showDeleted := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    matchAllTags := true // bool |  (optional)
    matchAllProjectTags := true // bool |  (optional)
    matchAllProjectTagIds := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    includeMainTask := true // bool |  (optional)
    includeArchivedProjects := true // bool |  (optional)
    getTotals := true // bool |  (optional)
    userId := []int32{int32(123)} // []int32 |  (optional)
    taskTagIds := []int32{int32(123)} // []int32 |  (optional)
    tagIds := []int32{int32(123)} // []int32 |  (optional)
    projectsFromCompanyId := []int32{int32(123)} // []int32 |  (optional)
    projectTagIds := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    projectIds := []int32{int32(123)} // []int32 |  (optional)
    projectId := []int32{int32(123)} // []int32 | project filters (optional)
    projectHealths := []int32{int32(123)} // []int32 |  (optional)
    projectHealth := []int32{int32(123)} // []int32 |  (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryId := []int32{int32(123)} // []int32 |  (optional)
    companyId := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV1TimeEntries(context.Background()).UpdatedAfterDate(updatedAfterDate).ToDate(toDate).SortOrder(sortOrder).SortBy(sortBy).ProjectStatus(projectStatus).InvoicedType(invoicedType).FromDate(fromDate).BillableType(billableType).TimeLogId(timeLogId).TicketId(ticketId).PageSize(pageSize).Page(page).IncludeTimelogId(includeTimelogId).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTagIds(matchAllProjectTagIds).IncludeTags(includeTags).IncludeMainTask(includeMainTask).IncludeArchivedProjects(includeArchivedProjects).GetTotals(getTotals).UserId(userId).TaskTagIds(taskTagIds).TagIds(tagIds).ProjectsFromCompanyId(projectsFromCompanyId).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ProjectCategoryId(projectCategoryId).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV1TimeEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV1TimeEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfterDate** | **string** |  | 
 **toDate** | **string** |  | 
 **sortOrder** | **string** |  | 
 **sortBy** | **string** |  | 
 **projectStatus** | **string** |  | 
 **invoicedType** | **string** |  | 
 **fromDate** | **string** |  | 
 **billableType** | **string** |  | 
 **timeLogId** | **int32** |  | 
 **ticketId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **includeTimelogId** | **int32** |  | 
 **showDeleted** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **matchAllTags** | **bool** |  | 
 **matchAllProjectTags** | **bool** |  | 
 **matchAllProjectTagIds** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **includeMainTask** | **bool** |  | 
 **includeArchivedProjects** | **bool** |  | 
 **getTotals** | **bool** |  | 
 **userId** | **[]int32** |  | 
 **taskTagIds** | **[]int32** |  | 
 **tagIds** | **[]int32** |  | 
 **projectsFromCompanyId** | **[]int32** |  | 
 **projectTagIds** | **[]int32** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **projectIds** | **[]int32** |  | 
 **projectId** | **[]int32** | project filters | 
 **projectHealths** | **[]int32** |  | 
 **projectHealth** | **[]int32** |  | 
 **projectCompanyIds** | **[]int32** |  | 
 **projectCategoryIds** | **[]int32** |  | 
 **projectCategoryId** | **[]int32** |  | 
 **companyId** | **[]int32** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsApiV1Pendingfiles

> POSTProjectsApiV1Pendingfiles(ctx).Type_(type_).ReturnFormat(returnFormat).ImageType(imageType).ReturnLocation(returnLocation).File(file).Execute()

Posts a pending file

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
    type_ := "type__example" // string |  (optional)
    returnFormat := "returnFormat_example" // string |  (optional)
    imageType := "imageType_example" // string |  (optional)
    returnLocation := true // bool |  (optional)
    file := "file_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.POSTProjectsApiV1Pendingfiles(context.Background()).Type_(type_).ReturnFormat(returnFormat).ImageType(imageType).ReturnLocation(returnLocation).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.POSTProjectsApiV1Pendingfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV1PendingfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** |  | 
 **returnFormat** | **string** |  | 
 **imageType** | **string** |  | 
 **returnLocation** | **bool** |  | 
 **file** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

