# \TasksV2Api

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV2ProjectsIdPredecessorsJson**](TasksV2Api.md#GETProjectsApiV2ProjectsIdPredecessorsJson) | **Get** /projects/api/v2/projects/{id}/predecessors.json | Retrieve all Task Predecessors/Dependencies on a given Project
[**GETProjectsApiV2ProjectsIdTaskidsJson**](TasksV2Api.md#GETProjectsApiV2ProjectsIdTaskidsJson) | **Get** /projects/api/v2/projects/{id}/taskids.json | Get all Tasks IDs on a given Project
[**GETProjectsApiV2TasklistsTasklistIdTaskidsJson**](TasksV2Api.md#GETProjectsApiV2TasklistsTasklistIdTaskidsJson) | **Get** /projects/api/v2/tasklists/{id}/taskids.json | Get all Task IDs on a given Task List



## GETProjectsApiV2ProjectsIdPredecessorsJson

> InlineResponse20055 GETProjectsApiV2ProjectsIdPredecessorsJson(ctx, id).IncludeDependencies(includeDependencies).OnlyDependencies(onlyDependencies).OnlyBasicFields(onlyBasicFields).OnlyActionable(onlyActionable).NextPageToken(nextPageToken).Execute()

Retrieve all Task Predecessors/Dependencies on a given Project



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
    id := int32(56) // int32 | ID of the Project to query
    includeDependencies := true // bool |  (optional) (default to false)
    onlyDependencies := true // bool |  (optional) (default to false)
    onlyBasicFields := true // bool | When true, means a reduced set of fields are returned. (optional) (default to false)
    onlyActionable := true // bool | When true, only return results for active projects, tasklists and tasks (optional) (default to false)
    nextPageToken := int32(56) // int32 | For pagination, this will be the token returned in the previous page results. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksV2Api.GETProjectsApiV2ProjectsIdPredecessorsJson(context.Background(), id).IncludeDependencies(includeDependencies).OnlyDependencies(onlyDependencies).OnlyBasicFields(onlyBasicFields).OnlyActionable(onlyActionable).NextPageToken(nextPageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksV2Api.GETProjectsApiV2ProjectsIdPredecessorsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV2ProjectsIdPredecessorsJson`: InlineResponse20055
    fmt.Fprintf(os.Stdout, "Response from `TasksV2Api.GETProjectsApiV2ProjectsIdPredecessorsJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the Project to query | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2ProjectsIdPredecessorsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeDependencies** | **bool** |  | [default to false]
 **onlyDependencies** | **bool** |  | [default to false]
 **onlyBasicFields** | **bool** | When true, means a reduced set of fields are returned. | [default to false]
 **onlyActionable** | **bool** | When true, only return results for active projects, tasklists and tasks | [default to false]
 **nextPageToken** | **int32** | For pagination, this will be the token returned in the previous page results. | 

### Return type

[**InlineResponse20055**](InlineResponse20055.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV2ProjectsIdTaskidsJson

> InlineResponse20056 GETProjectsApiV2ProjectsIdTaskidsJson(ctx, id).Filter(filter).FilterText(filterText).ResponsiblePartyIds(responsiblePartyIds).CreatorIds(creatorIds).IgnoreStartDates(ignoreStartDates).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).Include(include).IncludeToday(includeToday).IncludeBlockedTasks(includeBlockedTasks).UpdatedAfterDate(updatedAfterDate).CompletedAfterDate(completedAfterDate).CompletedBeforeDate(completedBeforeDate).Sort(sort).SortOrder(sortOrder).Page(page).PageSize(pageSize).OffsetCount(offsetCount).StartDate(startDate).EndDate(endDate).ShowDeleted(showDeleted).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).DeletedAfterDate(deletedAfterDate).ShowCompletedLists(showCompletedLists).IncludeCompletedSubtasks(includeCompletedSubtasks).GetSubTasks(getSubTasks).IncludeCompletedTasks(includeCompletedTasks).IncludeTaskId(includeTaskId).ProjectIds(projectIds).ProjectStatus(projectStatus).IncludeArchivedProjects(includeArchivedProjects).ProjectHealths(projectHealths).ProjectCategoryIds(projectCategoryIds).ProjectCompanyIds(projectCompanyIds).ProjectOwnerIds(projectOwnerIds).ProjectTagIds(projectTagIds).MatchAllProjectTags(matchAllProjectTags).OnlyStarredProjects(onlyStarredProjects).FollowerIds(followerIds).TagIds(tagIds).CreatedAfterDate(createdAfterDate).CreatedBeforeDate(createdBeforeDate).OnlyTasksWithTickets(onlyTasksWithTickets).OnlyTasksWithUnreadComments(onlyTasksWithUnreadComments).IncludeUntaggedTasks(includeUntaggedTasks).OnlyUntaggedTasks(onlyUntaggedTasks).Priority(priority).ExcludeTagIds(excludeTagIds).OnlyTasksWithEstimatedTime(onlyTasksWithEstimatedTime).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).AssigneeCompanyIds(assigneeCompanyIds).Execute()

Get all Tasks IDs on a given Project



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
    id := int32(56) // int32 | ID of Project to query
    filter := "filter_example" // string |  (optional) (default to "anytime")
    filterText := []string{"Inner_example"} // []string |  (optional)
    responsiblePartyIds := []int32{int32(123)} // []int32 |  (optional)
    creatorIds := []int32{int32(123)} // []int32 |  (optional)
    ignoreStartDates := true // bool |  (optional) (default to false)
    useStartDatesForTodaysTasks := true // bool |  (optional)
    include := "include_example" // string |  (optional)
    includeToday := true // bool |  (optional)
    includeBlockedTasks := true // bool |  (optional)
    updatedAfterDate := time.Now() // time.Time |  (optional)
    completedAfterDate := time.Now() // time.Time |  (optional)
    completedBeforeDate := time.Now() // time.Time |  (optional)
    sort := "sort_example" // string |  (optional) (default to "\"manual\"")
    sortOrder := "sortOrder_example" // string |  (optional) (default to "asc")
    page := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    offsetCount := int32(56) // int32 |  (optional)
    startDate := time.Now() // string |  (optional)
    endDate := time.Now() // string |  (optional)
    showDeleted := true // bool |  (optional)
    includeTasksFromDeletedLists := true // bool |  (optional)
    deletedAfterDate := time.Now() // time.Time |  (optional)
    showCompletedLists := true // bool |  (optional)
    includeCompletedSubtasks := true // bool |  (optional)
    getSubTasks := true // bool |  (optional) (default to true)
    includeCompletedTasks := true // bool |  (optional)
    includeTaskId := int32(56) // int32 |  (optional)
    projectIds := []int32{int32(123)} // []int32 |  (optional)
    projectStatus := "projectStatus_example" // string |  (optional) (default to "active")
    includeArchivedProjects := true // bool |  (optional) (default to false)
    projectHealths := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 |  (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    projectTagIds := []int32{int32(123)} // []int32 |  (optional)
    matchAllProjectTags := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    followerIds := []string{"Inner_example"} // []string |  (optional)
    tagIds := []int32{int32(123)} // []int32 |  (optional)
    createdAfterDate := time.Now() // time.Time |  (optional)
    createdBeforeDate := time.Now() // time.Time |  (optional)
    onlyTasksWithTickets := true // bool |  (optional)
    onlyTasksWithUnreadComments := true // bool |  (optional)
    includeUntaggedTasks := true // bool |  (optional)
    onlyUntaggedTasks := true // bool |  (optional)
    priority := "priority_example" // string |  (optional)
    excludeTagIds := []int32{int32(123)} // []int32 |  (optional)
    onlyTasksWithEstimatedTime := true // bool |  (optional)
    includeTasksWithoutDueDates := true // bool |  (optional)
    assigneeCompanyIds := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksV2Api.GETProjectsApiV2ProjectsIdTaskidsJson(context.Background(), id).Filter(filter).FilterText(filterText).ResponsiblePartyIds(responsiblePartyIds).CreatorIds(creatorIds).IgnoreStartDates(ignoreStartDates).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).Include(include).IncludeToday(includeToday).IncludeBlockedTasks(includeBlockedTasks).UpdatedAfterDate(updatedAfterDate).CompletedAfterDate(completedAfterDate).CompletedBeforeDate(completedBeforeDate).Sort(sort).SortOrder(sortOrder).Page(page).PageSize(pageSize).OffsetCount(offsetCount).StartDate(startDate).EndDate(endDate).ShowDeleted(showDeleted).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).DeletedAfterDate(deletedAfterDate).ShowCompletedLists(showCompletedLists).IncludeCompletedSubtasks(includeCompletedSubtasks).GetSubTasks(getSubTasks).IncludeCompletedTasks(includeCompletedTasks).IncludeTaskId(includeTaskId).ProjectIds(projectIds).ProjectStatus(projectStatus).IncludeArchivedProjects(includeArchivedProjects).ProjectHealths(projectHealths).ProjectCategoryIds(projectCategoryIds).ProjectCompanyIds(projectCompanyIds).ProjectOwnerIds(projectOwnerIds).ProjectTagIds(projectTagIds).MatchAllProjectTags(matchAllProjectTags).OnlyStarredProjects(onlyStarredProjects).FollowerIds(followerIds).TagIds(tagIds).CreatedAfterDate(createdAfterDate).CreatedBeforeDate(createdBeforeDate).OnlyTasksWithTickets(onlyTasksWithTickets).OnlyTasksWithUnreadComments(onlyTasksWithUnreadComments).IncludeUntaggedTasks(includeUntaggedTasks).OnlyUntaggedTasks(onlyUntaggedTasks).Priority(priority).ExcludeTagIds(excludeTagIds).OnlyTasksWithEstimatedTime(onlyTasksWithEstimatedTime).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).AssigneeCompanyIds(assigneeCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksV2Api.GETProjectsApiV2ProjectsIdTaskidsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV2ProjectsIdTaskidsJson`: InlineResponse20056
    fmt.Fprintf(os.Stdout, "Response from `TasksV2Api.GETProjectsApiV2ProjectsIdTaskidsJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of Project to query | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2ProjectsIdTaskidsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** |  | [default to &quot;anytime&quot;]
 **filterText** | **[]string** |  | 
 **responsiblePartyIds** | **[]int32** |  | 
 **creatorIds** | **[]int32** |  | 
 **ignoreStartDates** | **bool** |  | [default to false]
 **useStartDatesForTodaysTasks** | **bool** |  | 
 **include** | **string** |  | 
 **includeToday** | **bool** |  | 
 **includeBlockedTasks** | **bool** |  | 
 **updatedAfterDate** | **time.Time** |  | 
 **completedAfterDate** | **time.Time** |  | 
 **completedBeforeDate** | **time.Time** |  | 
 **sort** | **string** |  | [default to &quot;\&quot;manual\&quot;&quot;]
 **sortOrder** | **string** |  | [default to &quot;asc&quot;]
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **offsetCount** | **int32** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **showDeleted** | **bool** |  | 
 **includeTasksFromDeletedLists** | **bool** |  | 
 **deletedAfterDate** | **time.Time** |  | 
 **showCompletedLists** | **bool** |  | 
 **includeCompletedSubtasks** | **bool** |  | 
 **getSubTasks** | **bool** |  | [default to true]
 **includeCompletedTasks** | **bool** |  | 
 **includeTaskId** | **int32** |  | 
 **projectIds** | **[]int32** |  | 
 **projectStatus** | **string** |  | [default to &quot;active&quot;]
 **includeArchivedProjects** | **bool** |  | [default to false]
 **projectHealths** | **[]int32** |  | 
 **projectCategoryIds** | **[]int32** |  | 
 **projectCompanyIds** | **[]int32** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **projectTagIds** | **[]int32** |  | 
 **matchAllProjectTags** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **followerIds** | **[]string** |  | 
 **tagIds** | **[]int32** |  | 
 **createdAfterDate** | **time.Time** |  | 
 **createdBeforeDate** | **time.Time** |  | 
 **onlyTasksWithTickets** | **bool** |  | 
 **onlyTasksWithUnreadComments** | **bool** |  | 
 **includeUntaggedTasks** | **bool** |  | 
 **onlyUntaggedTasks** | **bool** |  | 
 **priority** | **string** |  | 
 **excludeTagIds** | **[]int32** |  | 
 **onlyTasksWithEstimatedTime** | **bool** |  | 
 **includeTasksWithoutDueDates** | **bool** |  | 
 **assigneeCompanyIds** | **[]int32** |  | 

### Return type

[**InlineResponse20056**](InlineResponse20056.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV2TasklistsTasklistIdTaskidsJson

> InlineResponse20056 GETProjectsApiV2TasklistsTasklistIdTaskidsJson(ctx, id).Filter(filter).FilterText(filterText).ResponsiblePartyIds(responsiblePartyIds).CreatorIds(creatorIds).IgnoreStartDates(ignoreStartDates).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).Include(include).IncludeToday(includeToday).IncludeBlockedTasks(includeBlockedTasks).UpdatedAfterDate(updatedAfterDate).CompletedAfterDate(completedAfterDate).CompletedBeforeDate(completedBeforeDate).Sort(sort).SortOrder(sortOrder).Page(page).PageSize(pageSize).OffsetCount(offsetCount).StartDate(startDate).EndDate(endDate).ShowDeleted(showDeleted).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).DeletedAfterDate(deletedAfterDate).ShowCompletedLists(showCompletedLists).IncludeCompletedSubtasks(includeCompletedSubtasks).GetSubTasks(getSubTasks).IncludeCompletedTasks(includeCompletedTasks).IncludeTaskId(includeTaskId).ProjectIds(projectIds).ProjectStatus(projectStatus).IncludeArchivedProjects(includeArchivedProjects).ProjectHealths(projectHealths).ProjectCategoryIds(projectCategoryIds).ProjectCompanyIds(projectCompanyIds).ProjectOwnerIds(projectOwnerIds).ProjectTagIds(projectTagIds).MatchAllProjectTags(matchAllProjectTags).OnlyStarredProjects(onlyStarredProjects).FollowerIds(followerIds).TagIds(tagIds).CreatedAfterDate(createdAfterDate).CreatedBeforeDate(createdBeforeDate).OnlyTasksWithTickets(onlyTasksWithTickets).OnlyTasksWithUnreadComments(onlyTasksWithUnreadComments).IncludeUntaggedTasks(includeUntaggedTasks).OnlyUntaggedTasks(onlyUntaggedTasks).Priority(priority).ExcludeTagIds(excludeTagIds).OnlyTasksWithEstimatedTime(onlyTasksWithEstimatedTime).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).AssigneeCompanyIds(assigneeCompanyIds).Execute()

Get all Task IDs on a given Task List



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
    id := int32(56) // int32 | ID of Project to query
    filter := "filter_example" // string |  (optional) (default to "anytime")
    filterText := []string{"Inner_example"} // []string |  (optional)
    responsiblePartyIds := []int32{int32(123)} // []int32 |  (optional)
    creatorIds := []int32{int32(123)} // []int32 |  (optional)
    ignoreStartDates := true // bool |  (optional) (default to false)
    useStartDatesForTodaysTasks := true // bool |  (optional)
    include := "include_example" // string |  (optional)
    includeToday := true // bool |  (optional)
    includeBlockedTasks := true // bool |  (optional)
    updatedAfterDate := time.Now() // time.Time |  (optional)
    completedAfterDate := time.Now() // time.Time |  (optional)
    completedBeforeDate := time.Now() // time.Time |  (optional)
    sort := "sort_example" // string |  (optional) (default to "\"manual\"")
    sortOrder := "sortOrder_example" // string |  (optional) (default to "asc")
    page := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    offsetCount := int32(56) // int32 |  (optional)
    startDate := time.Now() // string |  (optional)
    endDate := time.Now() // string |  (optional)
    showDeleted := true // bool |  (optional)
    includeTasksFromDeletedLists := true // bool |  (optional)
    deletedAfterDate := time.Now() // time.Time |  (optional)
    showCompletedLists := true // bool |  (optional)
    includeCompletedSubtasks := true // bool |  (optional)
    getSubTasks := true // bool |  (optional) (default to true)
    includeCompletedTasks := true // bool |  (optional)
    includeTaskId := int32(56) // int32 |  (optional)
    projectIds := []int32{int32(123)} // []int32 |  (optional)
    projectStatus := "projectStatus_example" // string |  (optional) (default to "active")
    includeArchivedProjects := true // bool |  (optional) (default to false)
    projectHealths := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 |  (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    projectTagIds := []int32{int32(123)} // []int32 |  (optional)
    matchAllProjectTags := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    followerIds := []string{"Inner_example"} // []string |  (optional)
    tagIds := []int32{int32(123)} // []int32 |  (optional)
    createdAfterDate := time.Now() // time.Time |  (optional)
    createdBeforeDate := time.Now() // time.Time |  (optional)
    onlyTasksWithTickets := true // bool |  (optional)
    onlyTasksWithUnreadComments := true // bool |  (optional)
    includeUntaggedTasks := true // bool |  (optional)
    onlyUntaggedTasks := true // bool |  (optional)
    priority := "priority_example" // string |  (optional)
    excludeTagIds := []int32{int32(123)} // []int32 |  (optional)
    onlyTasksWithEstimatedTime := true // bool |  (optional)
    includeTasksWithoutDueDates := true // bool |  (optional)
    assigneeCompanyIds := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksV2Api.GETProjectsApiV2TasklistsTasklistIdTaskidsJson(context.Background(), id).Filter(filter).FilterText(filterText).ResponsiblePartyIds(responsiblePartyIds).CreatorIds(creatorIds).IgnoreStartDates(ignoreStartDates).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).Include(include).IncludeToday(includeToday).IncludeBlockedTasks(includeBlockedTasks).UpdatedAfterDate(updatedAfterDate).CompletedAfterDate(completedAfterDate).CompletedBeforeDate(completedBeforeDate).Sort(sort).SortOrder(sortOrder).Page(page).PageSize(pageSize).OffsetCount(offsetCount).StartDate(startDate).EndDate(endDate).ShowDeleted(showDeleted).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).DeletedAfterDate(deletedAfterDate).ShowCompletedLists(showCompletedLists).IncludeCompletedSubtasks(includeCompletedSubtasks).GetSubTasks(getSubTasks).IncludeCompletedTasks(includeCompletedTasks).IncludeTaskId(includeTaskId).ProjectIds(projectIds).ProjectStatus(projectStatus).IncludeArchivedProjects(includeArchivedProjects).ProjectHealths(projectHealths).ProjectCategoryIds(projectCategoryIds).ProjectCompanyIds(projectCompanyIds).ProjectOwnerIds(projectOwnerIds).ProjectTagIds(projectTagIds).MatchAllProjectTags(matchAllProjectTags).OnlyStarredProjects(onlyStarredProjects).FollowerIds(followerIds).TagIds(tagIds).CreatedAfterDate(createdAfterDate).CreatedBeforeDate(createdBeforeDate).OnlyTasksWithTickets(onlyTasksWithTickets).OnlyTasksWithUnreadComments(onlyTasksWithUnreadComments).IncludeUntaggedTasks(includeUntaggedTasks).OnlyUntaggedTasks(onlyUntaggedTasks).Priority(priority).ExcludeTagIds(excludeTagIds).OnlyTasksWithEstimatedTime(onlyTasksWithEstimatedTime).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).AssigneeCompanyIds(assigneeCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksV2Api.GETProjectsApiV2TasklistsTasklistIdTaskidsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV2TasklistsTasklistIdTaskidsJson`: InlineResponse20056
    fmt.Fprintf(os.Stdout, "Response from `TasksV2Api.GETProjectsApiV2TasklistsTasklistIdTaskidsJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of Project to query | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2TasklistsTasklistIdTaskidsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** |  | [default to &quot;anytime&quot;]
 **filterText** | **[]string** |  | 
 **responsiblePartyIds** | **[]int32** |  | 
 **creatorIds** | **[]int32** |  | 
 **ignoreStartDates** | **bool** |  | [default to false]
 **useStartDatesForTodaysTasks** | **bool** |  | 
 **include** | **string** |  | 
 **includeToday** | **bool** |  | 
 **includeBlockedTasks** | **bool** |  | 
 **updatedAfterDate** | **time.Time** |  | 
 **completedAfterDate** | **time.Time** |  | 
 **completedBeforeDate** | **time.Time** |  | 
 **sort** | **string** |  | [default to &quot;\&quot;manual\&quot;&quot;]
 **sortOrder** | **string** |  | [default to &quot;asc&quot;]
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **offsetCount** | **int32** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **showDeleted** | **bool** |  | 
 **includeTasksFromDeletedLists** | **bool** |  | 
 **deletedAfterDate** | **time.Time** |  | 
 **showCompletedLists** | **bool** |  | 
 **includeCompletedSubtasks** | **bool** |  | 
 **getSubTasks** | **bool** |  | [default to true]
 **includeCompletedTasks** | **bool** |  | 
 **includeTaskId** | **int32** |  | 
 **projectIds** | **[]int32** |  | 
 **projectStatus** | **string** |  | [default to &quot;active&quot;]
 **includeArchivedProjects** | **bool** |  | [default to false]
 **projectHealths** | **[]int32** |  | 
 **projectCategoryIds** | **[]int32** |  | 
 **projectCompanyIds** | **[]int32** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **projectTagIds** | **[]int32** |  | 
 **matchAllProjectTags** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **followerIds** | **[]string** |  | 
 **tagIds** | **[]int32** |  | 
 **createdAfterDate** | **time.Time** |  | 
 **createdBeforeDate** | **time.Time** |  | 
 **onlyTasksWithTickets** | **bool** |  | 
 **onlyTasksWithUnreadComments** | **bool** |  | 
 **includeUntaggedTasks** | **bool** |  | 
 **onlyUntaggedTasks** | **bool** |  | 
 **priority** | **string** |  | 
 **excludeTagIds** | **[]int32** |  | 
 **onlyTasksWithEstimatedTime** | **bool** |  | 
 **includeTasksWithoutDueDates** | **bool** |  | 
 **assigneeCompanyIds** | **[]int32** |  | 

### Return type

[**InlineResponse20056**](InlineResponse20056.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

