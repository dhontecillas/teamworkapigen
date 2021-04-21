# \TasksApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEProjectsApiV3TaskstaskIdJson**](TasksApi.md#DELETEProjectsApiV3TaskstaskIdJson) | **Delete** /projects/api/v3/tasks/{taskId}.json | Delete an existing task.
[**GETProjectsApiV3TasksJson**](TasksApi.md#GETProjectsApiV3TasksJson) | **Get** /projects/api/v3/tasks.json | Get all tasks.
[**GETProjectsApiV3TasksMetricsCompleteJson**](TasksApi.md#GETProjectsApiV3TasksMetricsCompleteJson) | **Get** /projects/api/v3/tasks/metrics/complete.json | Total count of completed tasks
[**GETProjectsApiV3TasksMetricsLateJson**](TasksApi.md#GETProjectsApiV3TasksMetricsLateJson) | **Get** /projects/api/v3/tasks/metrics/late.json | Get total count of late tasks
[**GETProjectsApiV3TaskstaskIdJson**](TasksApi.md#GETProjectsApiV3TaskstaskIdJson) | **Get** /projects/api/v3/tasks/{taskId}.json | Get a specific task.
[**POSTProjectsApiV3TaskliststasklistIdTasksJson**](TasksApi.md#POSTProjectsApiV3TaskliststasklistIdTasksJson) | **Post** /projects/api/v3/tasklists/{tasklistId}/tasks.json | Creates a task.
[**POSTProjectsApiV3TasksBulkDeleteJson**](TasksApi.md#POSTProjectsApiV3TasksBulkDeleteJson) | **Post** /projects/api/v3/tasks/bulk/delete.json | Delete many tasks at once.



## DELETEProjectsApiV3TaskstaskIdJson

> DELETEProjectsApiV3TaskstaskIdJson(ctx, taskId).Execute()

Delete an existing task.



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
    resp, r, err := api_client.TasksApi.DELETEProjectsApiV3TaskstaskIdJson(context.Background(), taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.DELETEProjectsApiV3TaskstaskIdJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDELETEProjectsApiV3TaskstaskIdJsonRequest struct via the builder pattern


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


## GETProjectsApiV3TasksJson

> TaskTasksResponse GETProjectsApiV3TasksJson(ctx).UpdatedAfter(updatedAfter).Today(today).TaskIncludedSet(taskIncludedSet).TaskFilter(taskFilter).Status(status).StartDate(startDate).SearchTerm(searchTerm).ProjectStatuses(projectStatuses).Priority(priority).OrderMode(orderMode).OrderBy(orderBy).EndDate(endDate).DeletedAfter(deletedAfter).CreatedFilter(createdFilter).CreatedDateCode(createdDateCode).CreatedBefore(createdBefore).CreatedAfter(createdAfter).CompletedBefore(completedBefore).CompletedAfter(completedAfter).ProjectHealths(projectHealths).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).IncludeTaskId(includeTaskId).CompletedByUserId(completedByUserId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).SearchCompaniesTeams(searchCompaniesTeams).SearchAssignees(searchAssignees).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyTasksWithUnreadComments(onlyTasksWithUnreadComments).OnlyTasksWithTickets(onlyTasksWithTickets).OnlyStarredProjects(onlyStarredProjects).OnlyAdminProjects(onlyAdminProjects).NestSubTasks(nestSubTasks).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllExcludedTags(matchAllExcludedTags).IncludeUpdate(includeUpdate).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksWithCards(includeTasksWithCards).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeRelatedTasks(includeRelatedTasks).IncludePrivateItems(includePrivateItems).IncludeOverdueTasks(includeOverdueTasks).IncludeOriginalDueDate(includeOriginalDueDate).IncludeLoggedTime(includeLoggedTime).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCRMDealIds(includeCRMDealIds).IncludeBlocked(includeBlocked).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeAllComments(includeAllComments).GroupByTasklist(groupByTasklist).GetSubTasks(getSubTasks).GetFiles(getFiles).ExcludeAssigneeNotOnProjectTeams(excludeAssigneeNotOnProjectTeams).OnlyTasksWithEstimatedTime(onlyTasksWithEstimatedTime).TasklistIds(tasklistIds).Tags(tags).TagIds(tagIds).SkipCRMDealIds(skipCRMDealIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).Ids(ids).FollowedByUserIds(followedByUserIds).FilterBoardColumnIds(filterBoardColumnIds).FieldsUsers(fieldsUsers).FieldsTeams(fieldsTeams).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsMilestones(fieldsMilestones).FieldsCompanies(fieldsCompanies).FieldsComments(fieldsComments).FieldsColumns(fieldsColumns).FieldsCards(fieldsCards).ExcludeTagIds(excludeTagIds).CustomFields(customFields).CrmDealIds(crmDealIds).CreatedByUserIds(createdByUserIds).AssigneeTeamIds(assigneeTeamIds).AssigneeCompanyIds(assigneeCompanyIds).Execute()

Get all tasks.



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
    updatedAfter := time.Now() // string | filter by updated after date (optional)
    today := time.Now() // time.Time | filter by today (optional)
    taskIncludedSet := "taskIncludedSet_example" // string | filter by task included set (optional)
    taskFilter := "taskFilter_example" // string | filter by a taskFilter (optional)
    status := "status_example" // string | filter by list of task status (optional)
    startDate := time.Now() // string | filter on start date (optional)
    searchTerm := "searchTerm_example" // string | filter by search term (optional)
    projectStatuses := "projectStatuses_example" // string | filter by project status (optional)
    priority := "priority_example" // string | filter by task priority (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "duedate")
    endDate := time.Now() // string | filter on end date (optional)
    deletedAfter := time.Now() // time.Time | filter on deleted after date (optional)
    createdFilter := "createdFilter_example" // string | filter by created filter (optional)
    createdDateCode := "createdDateCode_example" // string | filter by created date code (optional)
    createdBefore := time.Now() // string | filter by created before date (optional)
    createdAfter := time.Now() // time.Time | filter by created after date (optional)
    completedBefore := time.Now() // time.Time | filter by completed before date (optional)
    completedAfter := time.Now() // time.Time | filter by completed after date (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    parentTaskId := int32(56) // int32 | filter by parent task ids (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    includeTaskId := int32(56) // int32 | include task id (optional)
    completedByUserId := int32(56) // int32 | filter by completer user id (optional)
    useStartDatesForTodaysTasks := true // bool | use start dates for todays tasks (optional)
    useAllProjects := true // bool | filter on all projects (optional)
    showDeleted := true // bool | include deleted items (optional) (default to false)
    showCompletedLists := true // bool | include tasks from completed lists (optional)
    searchCompaniesTeams := true // bool | include companies and teams in the search term (optional)
    searchAssignees := true // bool | include assignees in the search (optional)
    onlyUntaggedTasks := true // bool | only untagged tasks (optional)
    onlyTasksWithUnreadComments := true // bool | filter by only tasks with unread comments (optional)
    onlyTasksWithTickets := true // bool | filter by only tasks with tickets (optional)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    onlyAdminProjects := true // bool | only include tasks from projects where the user is strictly a project admin. site admins have visibility to all projects. (optional)
    nestSubTasks := true // bool | nest sub tasks (optional)
    matchAllTags := true // bool | match all tags (optional)
    matchAllProjectTags := true // bool | match all project tags (optional)
    matchAllExcludedTags := true // bool | match all exclude tags (optional)
    includeUpdate := true // bool | include tasks latest update action (optional)
    includeUntaggedTasks := true // bool | include untagged tasks (optional)
    includeToday := true // bool | filter by include today (optional)
    includeTasksWithoutDueDates := true // bool | include tasks without due dates (optional)
    includeTasksWithCards := true // bool | include tasks with cards (optional)
    includeTasksFromDeletedLists := true // bool | include tasks from deleted lists (optional)
    includeReminders := true // bool | include reminders (optional)
    includeRelatedTasks := true // bool | include ids of completed and active subtasks, dependencies, predecessors (optional)
    includePrivateItems := true // bool | include private items (optional)
    includeOverdueTasks := true // bool | include overdue tasks (optional)
    includeOriginalDueDate := true // bool | include original due date of a task (optional)
    includeLoggedTime := true // bool | include logged time (optional)
    includeCompletedTasks := true // bool | include completed tasks (optional)
    includeCompletedSubtasks := true // bool | include completed sub tasks (optional)
    includeCRMDealIds := true // bool | filter by include crm deal ids (optional)
    includeBlocked := true // bool | filter by include blocked (optional)
    includeAssigneeTeams := true // bool | include teams related to the responsible user ids (optional)
    includeAssigneeCompanies := true // bool | include companies related to the responsible user ids (optional)
    includeAllComments := true // bool | include all comments (optional)
    groupByTasklist := true // bool | group by tasklist (optional)
    getSubTasks := true // bool | get sub tasks (optional)
    getFiles := true // bool | get files (optional)
    excludeAssigneeNotOnProjectTeams := true // bool | exclude assignee not on project teams (optional)
    onlyTasksWithEstimatedTime := true // bool | only return tasks with estimated time (optional)
    tasklistIds := []int32{int32(123)} // []int32 | filter by tasklist ids (optional)
    tags := []string{"Inner_example"} // []string | filter by tag values (optional)
    tagIds := []int32{int32(123)} // []int32 | filter by tag ids (optional)
    skipCRMDealIds := []int32{int32(123)} // []int32 | skip crm deal ids (optional)
    responsiblePartyIds := []int32{int32(123)} // []int32 | filter by responsible party ids (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    ids := []int32{int32(123)} // []int32 | filter by task ids (optional)
    followedByUserIds := []int32{int32(123)} // []int32 | filter by followed by user ids (optional)
    filterBoardColumnIds := []int32{int32(123)} // []int32 | filter by board column ids (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTeams := []string{"Inner_example"} // []string |  (optional)
    fieldsTasks := []string{"Inner_example"} // []string |  (optional)
    fieldsTasklists := []string{"Inner_example"} // []string |  (optional)
    fieldsTags := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsMilestones := []string{"Inner_example"} // []string |  (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)
    fieldsComments := []string{"Inner_example"} // []string |  (optional)
    fieldsColumns := []string{"Inner_example"} // []string |  (optional)
    fieldsCards := []string{"Inner_example"} // []string |  (optional)
    excludeTagIds := []int32{int32(123)} // []int32 | filter by excluded tag ids (optional)
    customFields := []string{"Inner_example"} // []string | filter by custom fields (optional)
    crmDealIds := []int32{int32(123)} // []int32 | filter by crm deal ids (optional)
    createdByUserIds := []int32{int32(123)} // []int32 | filter by creator user ids (optional)
    assigneeTeamIds := []int32{int32(123)} // []int32 | filter by assignee team ids (optional)
    assigneeCompanyIds := []int32{int32(123)} // []int32 | filter by assignee company ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.GETProjectsApiV3TasksJson(context.Background()).UpdatedAfter(updatedAfter).Today(today).TaskIncludedSet(taskIncludedSet).TaskFilter(taskFilter).Status(status).StartDate(startDate).SearchTerm(searchTerm).ProjectStatuses(projectStatuses).Priority(priority).OrderMode(orderMode).OrderBy(orderBy).EndDate(endDate).DeletedAfter(deletedAfter).CreatedFilter(createdFilter).CreatedDateCode(createdDateCode).CreatedBefore(createdBefore).CreatedAfter(createdAfter).CompletedBefore(completedBefore).CompletedAfter(completedAfter).ProjectHealths(projectHealths).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).IncludeTaskId(includeTaskId).CompletedByUserId(completedByUserId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).SearchCompaniesTeams(searchCompaniesTeams).SearchAssignees(searchAssignees).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyTasksWithUnreadComments(onlyTasksWithUnreadComments).OnlyTasksWithTickets(onlyTasksWithTickets).OnlyStarredProjects(onlyStarredProjects).OnlyAdminProjects(onlyAdminProjects).NestSubTasks(nestSubTasks).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllExcludedTags(matchAllExcludedTags).IncludeUpdate(includeUpdate).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksWithCards(includeTasksWithCards).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeRelatedTasks(includeRelatedTasks).IncludePrivateItems(includePrivateItems).IncludeOverdueTasks(includeOverdueTasks).IncludeOriginalDueDate(includeOriginalDueDate).IncludeLoggedTime(includeLoggedTime).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCRMDealIds(includeCRMDealIds).IncludeBlocked(includeBlocked).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeAllComments(includeAllComments).GroupByTasklist(groupByTasklist).GetSubTasks(getSubTasks).GetFiles(getFiles).ExcludeAssigneeNotOnProjectTeams(excludeAssigneeNotOnProjectTeams).OnlyTasksWithEstimatedTime(onlyTasksWithEstimatedTime).TasklistIds(tasklistIds).Tags(tags).TagIds(tagIds).SkipCRMDealIds(skipCRMDealIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).Ids(ids).FollowedByUserIds(followedByUserIds).FilterBoardColumnIds(filterBoardColumnIds).FieldsUsers(fieldsUsers).FieldsTeams(fieldsTeams).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsMilestones(fieldsMilestones).FieldsCompanies(fieldsCompanies).FieldsComments(fieldsComments).FieldsColumns(fieldsColumns).FieldsCards(fieldsCards).ExcludeTagIds(excludeTagIds).CustomFields(customFields).CrmDealIds(crmDealIds).CreatedByUserIds(createdByUserIds).AssigneeTeamIds(assigneeTeamIds).AssigneeCompanyIds(assigneeCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.GETProjectsApiV3TasksJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3TasksJson`: TaskTasksResponse
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.GETProjectsApiV3TasksJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3TasksJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfter** | **string** | filter by updated after date | 
 **today** | **time.Time** | filter by today | 
 **taskIncludedSet** | **string** | filter by task included set | 
 **taskFilter** | **string** | filter by a taskFilter | 
 **status** | **string** | filter by list of task status | 
 **startDate** | **string** | filter on start date | 
 **searchTerm** | **string** | filter by search term | 
 **projectStatuses** | **string** | filter by project status | 
 **priority** | **string** | filter by task priority | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;duedate&quot;]
 **endDate** | **string** | filter on end date | 
 **deletedAfter** | **time.Time** | filter on deleted after date | 
 **createdFilter** | **string** | filter by created filter | 
 **createdDateCode** | **string** | filter by created date code | 
 **createdBefore** | **string** | filter by created before date | 
 **createdAfter** | **time.Time** | filter by created after date | 
 **completedBefore** | **time.Time** | filter by completed before date | 
 **completedAfter** | **time.Time** | filter by completed after date | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **parentTaskId** | **int32** | filter by parent task ids | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **includeTaskId** | **int32** | include task id | 
 **completedByUserId** | **int32** | filter by completer user id | 
 **useStartDatesForTodaysTasks** | **bool** | use start dates for todays tasks | 
 **useAllProjects** | **bool** | filter on all projects | 
 **showDeleted** | **bool** | include deleted items | [default to false]
 **showCompletedLists** | **bool** | include tasks from completed lists | 
 **searchCompaniesTeams** | **bool** | include companies and teams in the search term | 
 **searchAssignees** | **bool** | include assignees in the search | 
 **onlyUntaggedTasks** | **bool** | only untagged tasks | 
 **onlyTasksWithUnreadComments** | **bool** | filter by only tasks with unread comments | 
 **onlyTasksWithTickets** | **bool** | filter by only tasks with tickets | 
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **onlyAdminProjects** | **bool** | only include tasks from projects where the user is strictly a project admin. site admins have visibility to all projects. | 
 **nestSubTasks** | **bool** | nest sub tasks | 
 **matchAllTags** | **bool** | match all tags | 
 **matchAllProjectTags** | **bool** | match all project tags | 
 **matchAllExcludedTags** | **bool** | match all exclude tags | 
 **includeUpdate** | **bool** | include tasks latest update action | 
 **includeUntaggedTasks** | **bool** | include untagged tasks | 
 **includeToday** | **bool** | filter by include today | 
 **includeTasksWithoutDueDates** | **bool** | include tasks without due dates | 
 **includeTasksWithCards** | **bool** | include tasks with cards | 
 **includeTasksFromDeletedLists** | **bool** | include tasks from deleted lists | 
 **includeReminders** | **bool** | include reminders | 
 **includeRelatedTasks** | **bool** | include ids of completed and active subtasks, dependencies, predecessors | 
 **includePrivateItems** | **bool** | include private items | 
 **includeOverdueTasks** | **bool** | include overdue tasks | 
 **includeOriginalDueDate** | **bool** | include original due date of a task | 
 **includeLoggedTime** | **bool** | include logged time | 
 **includeCompletedTasks** | **bool** | include completed tasks | 
 **includeCompletedSubtasks** | **bool** | include completed sub tasks | 
 **includeCRMDealIds** | **bool** | filter by include crm deal ids | 
 **includeBlocked** | **bool** | filter by include blocked | 
 **includeAssigneeTeams** | **bool** | include teams related to the responsible user ids | 
 **includeAssigneeCompanies** | **bool** | include companies related to the responsible user ids | 
 **includeAllComments** | **bool** | include all comments | 
 **groupByTasklist** | **bool** | group by tasklist | 
 **getSubTasks** | **bool** | get sub tasks | 
 **getFiles** | **bool** | get files | 
 **excludeAssigneeNotOnProjectTeams** | **bool** | exclude assignee not on project teams | 
 **onlyTasksWithEstimatedTime** | **bool** | only return tasks with estimated time | 
 **tasklistIds** | **[]int32** | filter by tasklist ids | 
 **tags** | **[]string** | filter by tag values | 
 **tagIds** | **[]int32** | filter by tag ids | 
 **skipCRMDealIds** | **[]int32** | skip crm deal ids | 
 **responsiblePartyIds** | **[]int32** | filter by responsible party ids | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include | 
 **ids** | **[]int32** | filter by task ids | 
 **followedByUserIds** | **[]int32** | filter by followed by user ids | 
 **filterBoardColumnIds** | **[]int32** | filter by board column ids | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTeams** | **[]string** |  | 
 **fieldsTasks** | **[]string** |  | 
 **fieldsTasklists** | **[]string** |  | 
 **fieldsTags** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsMilestones** | **[]string** |  | 
 **fieldsCompanies** | **[]string** |  | 
 **fieldsComments** | **[]string** |  | 
 **fieldsColumns** | **[]string** |  | 
 **fieldsCards** | **[]string** |  | 
 **excludeTagIds** | **[]int32** | filter by excluded tag ids | 
 **customFields** | **[]string** | filter by custom fields | 
 **crmDealIds** | **[]int32** | filter by crm deal ids | 
 **createdByUserIds** | **[]int32** | filter by creator user ids | 
 **assigneeTeamIds** | **[]int32** | filter by assignee team ids | 
 **assigneeCompanyIds** | **[]int32** | filter by assignee company ids | 

### Return type

[**TaskTasksResponse**](TaskTasksResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3TasksMetricsCompleteJson

> CompleteResponse GETProjectsApiV3TasksMetricsCompleteJson(ctx).Execute()

Total count of completed tasks



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
    resp, r, err := api_client.TasksApi.GETProjectsApiV3TasksMetricsCompleteJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.GETProjectsApiV3TasksMetricsCompleteJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3TasksMetricsCompleteJson`: CompleteResponse
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.GETProjectsApiV3TasksMetricsCompleteJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3TasksMetricsCompleteJsonRequest struct via the builder pattern


### Return type

[**CompleteResponse**](CompleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3TasksMetricsLateJson

> LateResponse GETProjectsApiV3TasksMetricsLateJson(ctx).Execute()

Get total count of late tasks



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
    resp, r, err := api_client.TasksApi.GETProjectsApiV3TasksMetricsLateJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.GETProjectsApiV3TasksMetricsLateJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3TasksMetricsLateJson`: LateResponse
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.GETProjectsApiV3TasksMetricsLateJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3TasksMetricsLateJsonRequest struct via the builder pattern


### Return type

[**LateResponse**](LateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3TaskstaskIdJson

> TaskResponse GETProjectsApiV3TaskstaskIdJson(ctx, taskId).Execute()

Get a specific task.



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
    resp, r, err := api_client.TasksApi.GETProjectsApiV3TaskstaskIdJson(context.Background(), taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.GETProjectsApiV3TaskstaskIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3TaskstaskIdJson`: TaskResponse
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.GETProjectsApiV3TaskstaskIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3TaskstaskIdJsonRequest struct via the builder pattern


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


## POSTProjectsApiV3TaskliststasklistIdTasksJson

> TaskResponse POSTProjectsApiV3TaskliststasklistIdTasksJson(ctx, tasklistId).TaskRequest(taskRequest).Execute()

Creates a task.



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
    taskRequest := *openapiclient.NewTaskRequest() // TaskRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.POSTProjectsApiV3TaskliststasklistIdTasksJson(context.Background(), tasklistId).TaskRequest(taskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.POSTProjectsApiV3TaskliststasklistIdTasksJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsApiV3TaskliststasklistIdTasksJson`: TaskResponse
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.POSTProjectsApiV3TaskliststasklistIdTasksJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tasklistId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3TaskliststasklistIdTasksJsonRequest struct via the builder pattern


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


## POSTProjectsApiV3TasksBulkDeleteJson

> POSTProjectsApiV3TasksBulkDeleteJson(ctx).TaskDeleteRequest(taskDeleteRequest).Execute()

Delete many tasks at once.



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
    taskDeleteRequest := *openapiclient.NewTaskDeleteRequest() // TaskDeleteRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.POSTProjectsApiV3TasksBulkDeleteJson(context.Background()).TaskDeleteRequest(taskDeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.POSTProjectsApiV3TasksBulkDeleteJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3TasksBulkDeleteJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskDeleteRequest** | [**TaskDeleteRequest**](TaskDeleteRequest.md) |  | 

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

