# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV2PeopleuserIdAvatar**](DefaultApi.md#GETProjectsApiV2PeopleuserIdAvatar) | **Get** /projects/api/v2/people/{userId}/avatar | Get user avatar
[**GETProjectsApiV2ProjectsprojectIdPeopleuserIds**](DefaultApi.md#GETProjectsApiV2ProjectsprojectIdPeopleuserIds) | **Get** /projects/api/v2/projects/{projectId}/people/{userIds} | Get people project
[**GETProjectsApiV2ProjectsprojectIdTasksGantt**](DefaultApi.md#GETProjectsApiV2ProjectsprojectIdTasksGantt) | **Get** /projects/api/v2/projects/{projectId}/tasks/gantt | Will return the total for all sub-tasks



## GETProjectsApiV2PeopleuserIdAvatar

> GETProjectsApiV2PeopleuserIdAvatar(ctx, userId).Execute()

Get user avatar

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
    userId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GETProjectsApiV2PeopleuserIdAvatar(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GETProjectsApiV2PeopleuserIdAvatar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2PeopleuserIdAvatarRequest struct via the builder pattern


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


## GETProjectsApiV2ProjectsprojectIdPeopleuserIds

> PeopleResponse GETProjectsApiV2ProjectsprojectIdPeopleuserIds(ctx, userIds, projectId).UpdatedAfterDateTime(updatedAfterDateTime).UpdatedAfterDate(updatedAfterDate).Type_(type_).SortOrder(sortOrder).Sort(sort).SearchTerm(searchTerm).ProjectType(projectType).InviteStatus(inviteStatus).FirstLetter(firstLetter).EmailAddress(emailAddress).DataSet(dataSet).ProjectId2(projectId2).ProjectCompanyId(projectCompanyId).PageSize(pageSize).Page(page).ExcludeProjectId(excludeProjectId).ShowDeleted(showDeleted).ReturnTeams(returnTeams).ReturnProjectIds(returnProjectIds).ReturnLetters(returnLetters).OwnerCompanyFirst(ownerCompanyFirst).Onlyids(onlyids).IncludeTags(includeTags).IncludeObservers(includeObservers).IncludeContacts(includeContacts).IncludeCompanyDetails(includeCompanyDetails).IncludeClockIn(includeClockIn).IncludeClients(includeClients).GroupByCompany(groupByCompany).GetProjectRoles(getProjectRoles).GetCounts(getCounts).Fullprofile(fullprofile).CountOnly(countOnly).UserIds2(userIds2).CompanyId(companyId).Execute()

Get people project

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
    userIds := int32(56) // int32 | 
    projectId := int32(56) // int32 | 
    updatedAfterDateTime := "updatedAfterDateTime_example" // string |  (optional)
    updatedAfterDate := "updatedAfterDate_example" // string |  (optional)
    type_ := "type__example" // string |  (optional)
    sortOrder := "sortOrder_example" // string |  (optional)
    sort := "sort_example" // string |  (optional)
    searchTerm := "searchTerm_example" // string |  (optional)
    projectType := "projectType_example" // string |  (optional)
    inviteStatus := "inviteStatus_example" // string |  (optional)
    firstLetter := "firstLetter_example" // string |  (optional)
    emailAddress := "emailAddress_example" // string |  (optional)
    dataSet := "dataSet_example" // string |  (optional)
    projectId2 := int32(56) // int32 |  (optional)
    projectCompanyId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    excludeProjectId := int32(56) // int32 |  (optional)
    showDeleted := true // bool |  (optional)
    returnTeams := true // bool |  (optional)
    returnProjectIds := true // bool |  (optional)
    returnLetters := true // bool |  (optional)
    ownerCompanyFirst := true // bool |  (optional)
    onlyids := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    includeObservers := true // bool |  (optional)
    includeContacts := true // bool |  (optional)
    includeCompanyDetails := true // bool |  (optional)
    includeClockIn := true // bool |  (optional)
    includeClients := true // bool |  (optional)
    groupByCompany := true // bool |  (optional)
    getProjectRoles := true // bool |  (optional)
    getCounts := true // bool |  (optional)
    fullprofile := true // bool |  (optional)
    countOnly := true // bool |  (optional)
    userIds2 := []int32{int32(123)} // []int32 |  (optional)
    companyId := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GETProjectsApiV2ProjectsprojectIdPeopleuserIds(context.Background(), userIds, projectId).UpdatedAfterDateTime(updatedAfterDateTime).UpdatedAfterDate(updatedAfterDate).Type_(type_).SortOrder(sortOrder).Sort(sort).SearchTerm(searchTerm).ProjectType(projectType).InviteStatus(inviteStatus).FirstLetter(firstLetter).EmailAddress(emailAddress).DataSet(dataSet).ProjectId2(projectId2).ProjectCompanyId(projectCompanyId).PageSize(pageSize).Page(page).ExcludeProjectId(excludeProjectId).ShowDeleted(showDeleted).ReturnTeams(returnTeams).ReturnProjectIds(returnProjectIds).ReturnLetters(returnLetters).OwnerCompanyFirst(ownerCompanyFirst).Onlyids(onlyids).IncludeTags(includeTags).IncludeObservers(includeObservers).IncludeContacts(includeContacts).IncludeCompanyDetails(includeCompanyDetails).IncludeClockIn(includeClockIn).IncludeClients(includeClients).GroupByCompany(groupByCompany).GetProjectRoles(getProjectRoles).GetCounts(getCounts).Fullprofile(fullprofile).CountOnly(countOnly).UserIds2(userIds2).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GETProjectsApiV2ProjectsprojectIdPeopleuserIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV2ProjectsprojectIdPeopleuserIds`: PeopleResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GETProjectsApiV2ProjectsprojectIdPeopleuserIds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userIds** | **int32** |  | 
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2ProjectsprojectIdPeopleuserIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updatedAfterDateTime** | **string** |  | 
 **updatedAfterDate** | **string** |  | 
 **type_** | **string** |  | 
 **sortOrder** | **string** |  | 
 **sort** | **string** |  | 
 **searchTerm** | **string** |  | 
 **projectType** | **string** |  | 
 **inviteStatus** | **string** |  | 
 **firstLetter** | **string** |  | 
 **emailAddress** | **string** |  | 
 **dataSet** | **string** |  | 
 **projectId2** | **int32** |  | 
 **projectCompanyId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **excludeProjectId** | **int32** |  | 
 **showDeleted** | **bool** |  | 
 **returnTeams** | **bool** |  | 
 **returnProjectIds** | **bool** |  | 
 **returnLetters** | **bool** |  | 
 **ownerCompanyFirst** | **bool** |  | 
 **onlyids** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **includeObservers** | **bool** |  | 
 **includeContacts** | **bool** |  | 
 **includeCompanyDetails** | **bool** |  | 
 **includeClockIn** | **bool** |  | 
 **includeClients** | **bool** |  | 
 **groupByCompany** | **bool** |  | 
 **getProjectRoles** | **bool** |  | 
 **getCounts** | **bool** |  | 
 **fullprofile** | **bool** |  | 
 **countOnly** | **bool** |  | 
 **userIds2** | **[]int32** |  | 
 **companyId** | **[]int32** |  | 

### Return type

[**PeopleResponse**](PeopleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV2ProjectsprojectIdTasksGantt

> GETProjectsApiV2ProjectsprojectIdTasksGantt(ctx, projectId).UpdatedAfterDate(updatedAfterDate).Today(today).TagIds(tagIds).StartDate(startDate).Sort(sort).ResponsiblePartyIds(responsiblePartyIds).ProjectStatus(projectStatus).Include(include).FollowedByUserIds(followedByUserIds).Filter(filter).EndDate(endDate).DataSet(dataSet).CreatorIds(creatorIds).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TasklistId(tasklistId).TaskId(taskId).ProjectCategoryId(projectCategoryId).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).IncludeTaskId(includeTaskId).CompanyId(companyId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyStarredProjects(onlyStarredProjects).OnlyArchivedProjects(onlyArchivedProjects).NestSubTasks(nestSubTasks).MatchAllProjectTags(matchAllProjectTags).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompletedPredecessors(includeCompletedPredecessors).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).CountOnly(countOnly).AllowTemplateTasks(allowTemplateTasks).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId2(projectId2).ProjectHealths(projectHealths).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Execute()

Will return the total for all sub-tasks

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
    resp, r, err := api_client.DefaultApi.GETProjectsApiV2ProjectsprojectIdTasksGantt(context.Background(), projectId).UpdatedAfterDate(updatedAfterDate).Today(today).TagIds(tagIds).StartDate(startDate).Sort(sort).ResponsiblePartyIds(responsiblePartyIds).ProjectStatus(projectStatus).Include(include).FollowedByUserIds(followedByUserIds).Filter(filter).EndDate(endDate).DataSet(dataSet).CreatorIds(creatorIds).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TasklistId(tasklistId).TaskId(taskId).ProjectCategoryId(projectCategoryId).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).IncludeTaskId(includeTaskId).CompanyId(companyId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyStarredProjects(onlyStarredProjects).OnlyArchivedProjects(onlyArchivedProjects).NestSubTasks(nestSubTasks).MatchAllProjectTags(matchAllProjectTags).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompletedPredecessors(includeCompletedPredecessors).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).CountOnly(countOnly).AllowTemplateTasks(allowTemplateTasks).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId2(projectId2).ProjectHealths(projectHealths).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GETProjectsApiV2ProjectsprojectIdTasksGantt``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETProjectsApiV2ProjectsprojectIdTasksGanttRequest struct via the builder pattern


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

