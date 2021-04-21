# \TasksApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETETasksIdJson**](TasksApi.md#DELETETasksIdJson) | **Delete** /tasks/{id}.json | Delete a Task
[**GETCompletedtasksJson**](TasksApi.md#GETCompletedtasksJson) | **Get** /completedtasks.json | Get completed Tasks
[**GETProjectsIdTasksJson**](TasksApi.md#GETProjectsIdTasksJson) | **Get** /projects/{id}/tasks.json | Get all Tasks on a given Project
[**GETTasklistsIdTasksJson**](TasksApi.md#GETTasklistsIdTasksJson) | **Get** /tasklists/{id}/tasks.json | Get all Tasks on a given Task List
[**GETTasksIdDependenciesJson**](TasksApi.md#GETTasksIdDependenciesJson) | **Get** /tasks/{id}/dependencies.json | Retrieve Task Dependencies
[**GETTasksIdFollowersJson**](TasksApi.md#GETTasksIdFollowersJson) | **Get** /tasks/{id}/followers.json | Get Task Followers
[**GETTasksIdJson**](TasksApi.md#GETTasksIdJson) | **Get** /tasks/{id}.json | Retrieve a Task
[**GETTasksIdPredecessorsJson**](TasksApi.md#GETTasksIdPredecessorsJson) | **Get** /tasks/{id}/predecessors.json | Get Task Predecessors
[**GETTasksJson**](TasksApi.md#GETTasksJson) | **Get** /tasks.json | Get all Tasks across all Projects
[**GETTasksParentTaskIdSubtasksJson**](TasksApi.md#GETTasksParentTaskIdSubtasksJson) | **Get** /tasks/{parentTaskId}/subtasks.json | Get Sub Tasks of a Task
[**GETTasksTaskIdRecurringJson**](TasksApi.md#GETTasksTaskIdRecurringJson) | **Get** /tasks/{taskId}/recurring.json | Get Recurring Tasks related to original Task.
[**POSTProjectsProjIdTasksJson**](TasksApi.md#POSTProjectsProjIdTasksJson) | **Post** /projects/{projId}/tasks.json | Create a Task on a Project
[**POSTProjectsProjidTasksQuickaddJson**](TasksApi.md#POSTProjectsProjidTasksQuickaddJson) | **Post** /projects/{projid}/tasks/quickadd.json | Quickly add multiple Tasks
[**POSTTasklistsIdTasksJson**](TasksApi.md#POSTTasklistsIdTasksJson) | **Post** /tasklists/{id}/tasks.json | Create a Task
[**POSTTasksIdJson**](TasksApi.md#POSTTasksIdJson) | **Post** /tasks/{id}.json | Create a Sub Task
[**POSTTasksIdQuickaddJson**](TasksApi.md#POSTTasksIdQuickaddJson) | **Post** /tasks/{id}/quickadd.json | Quickly add multiple Sub Tasks
[**PUTTasklistsIdTasksReorderJson**](TasksApi.md#PUTTasklistsIdTasksReorderJson) | **Put** /tasklists/{id}/tasks/reorder.json | Reorder the Tasks
[**PUTTasksIdCompleteJson**](TasksApi.md#PUTTasksIdCompleteJson) | **Put** /tasks/{id}/complete.json | Mark a Task complete
[**PUTTasksIdJson**](TasksApi.md#PUTTasksIdJson) | **Put** /tasks/{id}.json | Update a Task
[**PUTTasksIdUncompleteJson**](TasksApi.md#PUTTasksIdUncompleteJson) | **Put** /tasks/{id}/uncomplete.json | Mark a Task uncomplete
[**PUTTasksTaskIdCopyJson**](TasksApi.md#PUTTasksTaskIdCopyJson) | **Put** /tasks/{taskId}/copy.json | Copy a Task from one Project to Another
[**PUTTasksTaskIdJson**](TasksApi.md#PUTTasksTaskIdJson) | **Put** /tasks/{taskId}.json | Editing Task Predecessors
[**PUTTasksTaskIdMoveJson**](TasksApi.md#PUTTasksTaskIdMoveJson) | **Put** /tasks/{taskId}/move.json | Move a Task from one Project to Another



## DELETETasksIdJson

> map[string]interface{} DELETETasksIdJson(ctx, id).Body(body).Execute()

Delete a Task



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
    id := "id_example" // string | 
    body := *openapiclient.NewInlineObject96() // InlineObject96 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.DELETETasksIdJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.DELETETasksIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DELETETasksIdJson`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.DELETETasksIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETETasksIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject96**](InlineObject96.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETCompletedtasksJson

> InlineResponse20015 GETCompletedtasksJson(ctx).Page(page).PageSize(pageSize).Startdate(startdate).Enddate(enddate).IncludeArchivedProjects(includeArchivedProjects).Execute()

Get completed Tasks



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
    page := int32(56) // int32 | Optionally, you can set the page from which to start retrieving results. This is usually used in conjunction with the parameter pageSize. For example: ?page=2&pageSize=10 will retrieve results 10-20. (optional)
    pageSize := int32(56) // int32 | The amount of tasks returned can be limited using this parameter. Normally used in conjunction with the page parameter. (optional)
    startdate := "startdate_example" // string | Tasks within a range of dates can be returned by setting a startdate and enddate. The format should be YYYYMMDD. For example: ?startdate=20140512&enddate=20140513 will get all of the tasks from the 12th of May 2014 till the 13th of May 2014. (optional)
    enddate := "enddate_example" // string | Must be used in conjunction with startdate, see above. (optional)
    includeArchivedProjects := true // bool | Set to true or false to include archived Projects in the response (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.GETCompletedtasksJson(context.Background()).Page(page).PageSize(pageSize).Startdate(startdate).Enddate(enddate).IncludeArchivedProjects(includeArchivedProjects).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.GETCompletedtasksJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCompletedtasksJson`: InlineResponse20015
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.GETCompletedtasksJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETCompletedtasksJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Optionally, you can set the page from which to start retrieving results. This is usually used in conjunction with the parameter pageSize. For example: ?page&#x3D;2&amp;pageSize&#x3D;10 will retrieve results 10-20. | 
 **pageSize** | **int32** | The amount of tasks returned can be limited using this parameter. Normally used in conjunction with the page parameter. | 
 **startdate** | **string** | Tasks within a range of dates can be returned by setting a startdate and enddate. The format should be YYYYMMDD. For example: ?startdate&#x3D;20140512&amp;enddate&#x3D;20140513 will get all of the tasks from the 12th of May 2014 till the 13th of May 2014. | 
 **enddate** | **string** | Must be used in conjunction with startdate, see above. | 
 **includeArchivedProjects** | **bool** | Set to true or false to include archived Projects in the response | 

### Return type

[**InlineResponse20015**](InlineResponse20015.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsIdTasksJson

> InlineResponse20077 GETProjectsIdTasksJson(ctx, id).Filter(filter).Page(page).PageSize(pageSize).StartDate(startDate).EndDate(endDate).UpdatedAfterDate(updatedAfterDate).CompletedAfterDate(completedAfterDate).CompletedBeforeDate(completedBeforeDate).ShowDeleted(showDeleted).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).CreatorIds(creatorIds).Include(include).ResponsiblePartyIds(responsiblePartyIds).Sort(sort).GetSubTasks(getSubTasks).NestSubTasks(nestSubTasks).GetFiles(getFiles).IncludeToday(includeToday).IgnoreStartDates(ignoreStartDates).TagIds(tagIds).DateupdatedASC(dateupdatedASC).Execute()

Get all Tasks on a given Project



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
    id := "id_example" // string | 
    filter := "filter_example" // string | Tasks can be filtered by due dates using the following options: all anytime overdue today tomorrow thisweek within7 within14 within30 within365 nodate nostartdate completed Additionally, you can choose to include the start dates in this critera by setting ignore-start-dates to false. (optional)
    page := int32(56) // int32 | Optionally, you can set the page from which to start retrieving results. This is usually used in conjunction with the parameter pageSize. For example: ?page=2&pageSize=10 will retrieve results 10-20. (optional)
    pageSize := int32(56) // int32 | The amount of tasks returned can be limited using this parameter. Normally used in conjunction with the page parameter. (optional)
    startDate := "startDate_example" // string | Tasks within a range of dates can be returned by setting a startdate and enddate. The format should be YYYYMMDD. For example: ?startdate=20140512&enddate=20140513 will get all of the tasks from the 12th of May 2014 till the 13th of May 2014.  By default ignore-start-dates=false and that means we list:  - tasks started before or on specified startDate AND due before/on specified endDate - OR no start date but due date between specified dates - OR no duedate set but startdate has passed/is specified startDate - OR milestone due date set and its between specified dates (optional)
    endDate := "endDate_example" // string | Must be used in conjunction with startdate, see above. (optional)
    updatedAfterDate := "updatedAfterDate_example" // string | Will only return tasks that have been updated after a specified date. Timestamp must be in the following format: YYYYMMDDHHMMSS. For example: ?updatedAfterDate=20131123201022 will only return tasks updated after 23rd of November 2013 at 20:10:22. (UTC) (optional)
    completedAfterDate := "completedAfterDate_example" // string | Will only return tasks that have been completed after a specified date. Timestamp must be in the following format: YYYYMMDDHHMMSS For example: ?completedAfterDate=20141123201022 will only return tasks completed after 23rd of November 2014 at 20:10:22. (UTC) (optional)
    completedBeforeDate := "completedBeforeDate_example" // string | Will only return tasks that have been completed before a specified date. Timestamp must be in the following format: YYYYMMDDHHMMSS For example: ?completedBeforeDate=20150101235959 will only return tasks completed before 1st January 2015 at 23:59:59. (UTC) (optional)
    showDeleted := "showDeleted_example" // string | Tasks that have been deleted can be shown by setting this parameter to yes. (optional)
    includeCompletedTasks := true // bool | Tasks that have been marked as completed can be shown by setting this parameter to true. (optional)
    includeCompletedSubtasks := true // bool | Sub-tasks that have been marked as completed can be shown by setting this parameter to true if you have requested to include sub-tasks (optional)
    creatorIds := "creatorIds_example" // string | For requesting tasks made by a specific person or people. For example: 44 would return tasks made by user 44. 44,45 would return tasks made by users 44 and/or 45 etc. (optional)
    include := "include_example" // string | Extra tasks that can be included with the filter option. nodate nostartdate noduedate overdue (optional)
    responsiblePartyIds := "responsiblePartyIds_example" // string | Tasks can be filtered by the person/people a task is assigned to. For example: -1 would return all tasks with an assigned person. 0 would return all tasks with no assignment. 32 would return tasks assigned to user 32. 32,55 would return tasks assigned to users 32 and/or 55 etc. (optional)
    sort := "sort_example" // string | Tasks can be sorted by the following options: 'duedate', duedateDESC', 'startdate', 'dateadded', 'priority', 'project' ,'manual', 'duestartdate', 'duestartdatedesc', 'alldates', 'alldatesdesc', 'completedDateDESC', 'flattenedtasklist', 'company' or 'dateupdated'\".  (optional)
    getSubTasks := "getSubTasks_example" // string | Subtasks can be excluded from the results by adding this parameter with no as the value. (optional)
    nestSubTasks := "nestSubTasks_example" // string | no Subtasks can be nested within the parent task object by adding this paramter with yes as the value. (optional)
    getFiles := true // bool | Files attached to tasks can be returned within the task object by setting this parameter to true. (optional)
    includeToday := true // bool | When using the filter option with any of the following options; within7,within14,within30,within365. You can choose to exclude deadlines for today by passing this parameter as false.   (optional)
    ignoreStartDates := true // bool | When using the filter option, you can choose to include start dates matching the filtering critera by passing this parameter as true. By default, only due dates are checked against the filter. (optional)
    tagIds := "tagIds_example" // string | A comma separated list of tag ids to filter tasks on (optional)
    dateupdatedASC := "dateupdatedASC_example" // string | Tasks can be sorted by the following options: 'duedate', duedateDESC', 'startdate', 'dateadded', 'priority', 'project' ,'manual', 'duestartdate', 'duestartdatedesc', 'alldates', 'alldatesdesc', 'completedDateDESC', 'flattenedtasklist', 'company' or 'dateupdated'\".  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.GETProjectsIdTasksJson(context.Background(), id).Filter(filter).Page(page).PageSize(pageSize).StartDate(startDate).EndDate(endDate).UpdatedAfterDate(updatedAfterDate).CompletedAfterDate(completedAfterDate).CompletedBeforeDate(completedBeforeDate).ShowDeleted(showDeleted).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).CreatorIds(creatorIds).Include(include).ResponsiblePartyIds(responsiblePartyIds).Sort(sort).GetSubTasks(getSubTasks).NestSubTasks(nestSubTasks).GetFiles(getFiles).IncludeToday(includeToday).IgnoreStartDates(ignoreStartDates).TagIds(tagIds).DateupdatedASC(dateupdatedASC).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.GETProjectsIdTasksJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsIdTasksJson`: InlineResponse20077
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.GETProjectsIdTasksJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsIdTasksJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | Tasks can be filtered by due dates using the following options: all anytime overdue today tomorrow thisweek within7 within14 within30 within365 nodate nostartdate completed Additionally, you can choose to include the start dates in this critera by setting ignore-start-dates to false. | 
 **page** | **int32** | Optionally, you can set the page from which to start retrieving results. This is usually used in conjunction with the parameter pageSize. For example: ?page&#x3D;2&amp;pageSize&#x3D;10 will retrieve results 10-20. | 
 **pageSize** | **int32** | The amount of tasks returned can be limited using this parameter. Normally used in conjunction with the page parameter. | 
 **startDate** | **string** | Tasks within a range of dates can be returned by setting a startdate and enddate. The format should be YYYYMMDD. For example: ?startdate&#x3D;20140512&amp;enddate&#x3D;20140513 will get all of the tasks from the 12th of May 2014 till the 13th of May 2014.  By default ignore-start-dates&#x3D;false and that means we list:  - tasks started before or on specified startDate AND due before/on specified endDate - OR no start date but due date between specified dates - OR no duedate set but startdate has passed/is specified startDate - OR milestone due date set and its between specified dates | 
 **endDate** | **string** | Must be used in conjunction with startdate, see above. | 
 **updatedAfterDate** | **string** | Will only return tasks that have been updated after a specified date. Timestamp must be in the following format: YYYYMMDDHHMMSS. For example: ?updatedAfterDate&#x3D;20131123201022 will only return tasks updated after 23rd of November 2013 at 20:10:22. (UTC) | 
 **completedAfterDate** | **string** | Will only return tasks that have been completed after a specified date. Timestamp must be in the following format: YYYYMMDDHHMMSS For example: ?completedAfterDate&#x3D;20141123201022 will only return tasks completed after 23rd of November 2014 at 20:10:22. (UTC) | 
 **completedBeforeDate** | **string** | Will only return tasks that have been completed before a specified date. Timestamp must be in the following format: YYYYMMDDHHMMSS For example: ?completedBeforeDate&#x3D;20150101235959 will only return tasks completed before 1st January 2015 at 23:59:59. (UTC) | 
 **showDeleted** | **string** | Tasks that have been deleted can be shown by setting this parameter to yes. | 
 **includeCompletedTasks** | **bool** | Tasks that have been marked as completed can be shown by setting this parameter to true. | 
 **includeCompletedSubtasks** | **bool** | Sub-tasks that have been marked as completed can be shown by setting this parameter to true if you have requested to include sub-tasks | 
 **creatorIds** | **string** | For requesting tasks made by a specific person or people. For example: 44 would return tasks made by user 44. 44,45 would return tasks made by users 44 and/or 45 etc. | 
 **include** | **string** | Extra tasks that can be included with the filter option. nodate nostartdate noduedate overdue | 
 **responsiblePartyIds** | **string** | Tasks can be filtered by the person/people a task is assigned to. For example: -1 would return all tasks with an assigned person. 0 would return all tasks with no assignment. 32 would return tasks assigned to user 32. 32,55 would return tasks assigned to users 32 and/or 55 etc. | 
 **sort** | **string** | Tasks can be sorted by the following options: &#39;duedate&#39;, duedateDESC&#39;, &#39;startdate&#39;, &#39;dateadded&#39;, &#39;priority&#39;, &#39;project&#39; ,&#39;manual&#39;, &#39;duestartdate&#39;, &#39;duestartdatedesc&#39;, &#39;alldates&#39;, &#39;alldatesdesc&#39;, &#39;completedDateDESC&#39;, &#39;flattenedtasklist&#39;, &#39;company&#39; or &#39;dateupdated&#39;\&quot;.  | 
 **getSubTasks** | **string** | Subtasks can be excluded from the results by adding this parameter with no as the value. | 
 **nestSubTasks** | **string** | no Subtasks can be nested within the parent task object by adding this paramter with yes as the value. | 
 **getFiles** | **bool** | Files attached to tasks can be returned within the task object by setting this parameter to true. | 
 **includeToday** | **bool** | When using the filter option with any of the following options; within7,within14,within30,within365. You can choose to exclude deadlines for today by passing this parameter as false.   | 
 **ignoreStartDates** | **bool** | When using the filter option, you can choose to include start dates matching the filtering critera by passing this parameter as true. By default, only due dates are checked against the filter. | 
 **tagIds** | **string** | A comma separated list of tag ids to filter tasks on | 
 **dateupdatedASC** | **string** | Tasks can be sorted by the following options: &#39;duedate&#39;, duedateDESC&#39;, &#39;startdate&#39;, &#39;dateadded&#39;, &#39;priority&#39;, &#39;project&#39; ,&#39;manual&#39;, &#39;duestartdate&#39;, &#39;duestartdatedesc&#39;, &#39;alldates&#39;, &#39;alldatesdesc&#39;, &#39;completedDateDESC&#39;, &#39;flattenedtasklist&#39;, &#39;company&#39; or &#39;dateupdated&#39;\&quot;.  | 

### Return type

[**InlineResponse20077**](InlineResponse20077.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETTasklistsIdTasksJson

> InlineResponse20097 GETTasklistsIdTasksJson(ctx, id).Filter(filter).Page(page).PageSize(pageSize).StartDate(startDate).EndDate(endDate).UpdatedAfterDate(updatedAfterDate).CompletedAfterDate(completedAfterDate).CompletedBeforeDate(completedBeforeDate).ShowDeleted(showDeleted).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).CreatorIds(creatorIds).Include(include).ResponsiblePartyIds(responsiblePartyIds).Sort(sort).GetSubTasks(getSubTasks).NestSubTasks(nestSubTasks).GetFiles(getFiles).IncludeToday(includeToday).IgnoreStartDates(ignoreStartDates).TagIds(tagIds).DateupdatedASC(dateupdatedASC).Execute()

Get all Tasks on a given Task List



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
    id := "id_example" // string | 
    filter := "filter_example" // string | Tasks can be filtered by due dates using the following options: all anytime overdue today tomorrow thisweek within7 within14 within30 within365 nodate nostartdate completed Additionally, you can choose to include the start dates in this critera by setting ignore-start-dates to false. (optional)
    page := int32(56) // int32 | Optionally, you can set the page from which to start retrieving results. This is usually used in conjunction with the parameter pageSize. For example: ?page=2&pageSize=10 will retrieve results 10-20. (optional)
    pageSize := int32(56) // int32 | The amount of tasks returned can be limited using this parameter. Normally used in conjunction with the page parameter. (optional)
    startDate := "startDate_example" // string | Tasks within a range of dates can be returned by setting a startdate and enddate. The format should be YYYYMMDD. For example: ?startdate=20140512&enddate=20140513 will get all of the tasks from the 12th of May 2014 till the 13th of May 2014.  By default ignore-start-dates=false and that means we list:  - tasks started before or on specified startDate AND due before/on specified endDate - OR no start date but due date between specified dates - OR no duedate set but startdate has passed/is specified startDate - OR milestone due date set and its between specified dates (optional)
    endDate := "endDate_example" // string | Must be used in conjunction with startdate, see above. (optional)
    updatedAfterDate := "updatedAfterDate_example" // string | Will only return tasks that have been updated after a specified date. Timestamp must be in the following format: YYYYMMDDHHMMSS. For example: ?updatedAfterDate=20131123201022 will only return tasks updated after 23rd of November 2013 at 20:10:22. (UTC) (optional)
    completedAfterDate := "completedAfterDate_example" // string | Will only return tasks that have been completed after a specified date. Timestamp must be in the following format: YYYYMMDDHHMMSS For example: ?completedAfterDate=20141123201022 will only return tasks completed after 23rd of November 2014 at 20:10:22. (UTC) (optional)
    completedBeforeDate := "completedBeforeDate_example" // string | Will only return tasks that have been completed before a specified date. Timestamp must be in the following format: YYYYMMDDHHMMSS For example: ?completedBeforeDate=20150101235959 will only return tasks completed before 1st January 2015 at 23:59:59. (UTC) (optional)
    showDeleted := "showDeleted_example" // string | Tasks that have been deleted can be shown by setting this parameter to yes. (optional)
    includeCompletedTasks := true // bool | Tasks that have been marked as completed can be shown by setting this parameter to true. (optional)
    includeCompletedSubtasks := true // bool | Sub-tasks that have been marked as completed can be shown by setting this parameter to true if you have requested to include sub-tasks (optional)
    creatorIds := "creatorIds_example" // string | For requesting tasks made by a specific person or people. For example: 44 would return tasks made by user 44. 44,45 would return tasks made by users 44 and/or 45 etc. (optional)
    include := "include_example" // string | Extra tasks that can be included with the filter option. nodate nostartdate noduedate overdue (optional)
    responsiblePartyIds := "responsiblePartyIds_example" // string | Tasks can be filtered by the person/people a task is assigned to. For example: -1 would return all tasks with an assigned person. 0 would return all tasks with no assignment. 32 would return tasks assigned to user 32. 32,55 would return tasks assigned to users 32 and/or 55 etc. (optional)
    sort := "sort_example" // string | Tasks can be sorted by the following options: 'duedate', duedateDESC', 'startdate', 'dateadded', 'priority', 'project' ,'manual', 'duestartdate', 'duestartdatedesc', 'alldates', 'alldatesdesc', 'completedDateDESC', 'flattenedtasklist', 'company' or 'dateupdated'\".  (optional)
    getSubTasks := "getSubTasks_example" // string | Subtasks can be excluded from the results by adding this parameter with no as the value. (optional)
    nestSubTasks := "nestSubTasks_example" // string | no Subtasks can be nested within the parent task object by adding this paramter with yes as the value. (optional)
    getFiles := true // bool | Files attached to tasks can be returned within the task object by setting this parameter to true. (optional)
    includeToday := true // bool | When using the filter option with any of the following options; within7,within14,within30,within365. You can choose to exclude deadlines for today by passing this parameter as false.   (optional)
    ignoreStartDates := true // bool | When using the filter option, you can choose to include start dates matching the filtering critera by passing this parameter as true. By default, only due dates are checked against the filter. (optional)
    tagIds := "tagIds_example" // string | A comma separated list of tag ids to filter tasks on (optional)
    dateupdatedASC := "dateupdatedASC_example" // string | Tasks can be sorted by the following options: 'duedate', duedateDESC', 'startdate', 'dateadded', 'priority', 'project' ,'manual', 'duestartdate', 'duestartdatedesc', 'alldates', 'alldatesdesc', 'completedDateDESC', 'flattenedtasklist', 'company' or 'dateupdated'\".  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.GETTasklistsIdTasksJson(context.Background(), id).Filter(filter).Page(page).PageSize(pageSize).StartDate(startDate).EndDate(endDate).UpdatedAfterDate(updatedAfterDate).CompletedAfterDate(completedAfterDate).CompletedBeforeDate(completedBeforeDate).ShowDeleted(showDeleted).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).CreatorIds(creatorIds).Include(include).ResponsiblePartyIds(responsiblePartyIds).Sort(sort).GetSubTasks(getSubTasks).NestSubTasks(nestSubTasks).GetFiles(getFiles).IncludeToday(includeToday).IgnoreStartDates(ignoreStartDates).TagIds(tagIds).DateupdatedASC(dateupdatedASC).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.GETTasklistsIdTasksJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETTasklistsIdTasksJson`: InlineResponse20097
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.GETTasklistsIdTasksJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETTasklistsIdTasksJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | Tasks can be filtered by due dates using the following options: all anytime overdue today tomorrow thisweek within7 within14 within30 within365 nodate nostartdate completed Additionally, you can choose to include the start dates in this critera by setting ignore-start-dates to false. | 
 **page** | **int32** | Optionally, you can set the page from which to start retrieving results. This is usually used in conjunction with the parameter pageSize. For example: ?page&#x3D;2&amp;pageSize&#x3D;10 will retrieve results 10-20. | 
 **pageSize** | **int32** | The amount of tasks returned can be limited using this parameter. Normally used in conjunction with the page parameter. | 
 **startDate** | **string** | Tasks within a range of dates can be returned by setting a startdate and enddate. The format should be YYYYMMDD. For example: ?startdate&#x3D;20140512&amp;enddate&#x3D;20140513 will get all of the tasks from the 12th of May 2014 till the 13th of May 2014.  By default ignore-start-dates&#x3D;false and that means we list:  - tasks started before or on specified startDate AND due before/on specified endDate - OR no start date but due date between specified dates - OR no duedate set but startdate has passed/is specified startDate - OR milestone due date set and its between specified dates | 
 **endDate** | **string** | Must be used in conjunction with startdate, see above. | 
 **updatedAfterDate** | **string** | Will only return tasks that have been updated after a specified date. Timestamp must be in the following format: YYYYMMDDHHMMSS. For example: ?updatedAfterDate&#x3D;20131123201022 will only return tasks updated after 23rd of November 2013 at 20:10:22. (UTC) | 
 **completedAfterDate** | **string** | Will only return tasks that have been completed after a specified date. Timestamp must be in the following format: YYYYMMDDHHMMSS For example: ?completedAfterDate&#x3D;20141123201022 will only return tasks completed after 23rd of November 2014 at 20:10:22. (UTC) | 
 **completedBeforeDate** | **string** | Will only return tasks that have been completed before a specified date. Timestamp must be in the following format: YYYYMMDDHHMMSS For example: ?completedBeforeDate&#x3D;20150101235959 will only return tasks completed before 1st January 2015 at 23:59:59. (UTC) | 
 **showDeleted** | **string** | Tasks that have been deleted can be shown by setting this parameter to yes. | 
 **includeCompletedTasks** | **bool** | Tasks that have been marked as completed can be shown by setting this parameter to true. | 
 **includeCompletedSubtasks** | **bool** | Sub-tasks that have been marked as completed can be shown by setting this parameter to true if you have requested to include sub-tasks | 
 **creatorIds** | **string** | For requesting tasks made by a specific person or people. For example: 44 would return tasks made by user 44. 44,45 would return tasks made by users 44 and/or 45 etc. | 
 **include** | **string** | Extra tasks that can be included with the filter option. nodate nostartdate noduedate overdue | 
 **responsiblePartyIds** | **string** | Tasks can be filtered by the person/people a task is assigned to. For example: -1 would return all tasks with an assigned person. 0 would return all tasks with no assignment. 32 would return tasks assigned to user 32. 32,55 would return tasks assigned to users 32 and/or 55 etc. | 
 **sort** | **string** | Tasks can be sorted by the following options: &#39;duedate&#39;, duedateDESC&#39;, &#39;startdate&#39;, &#39;dateadded&#39;, &#39;priority&#39;, &#39;project&#39; ,&#39;manual&#39;, &#39;duestartdate&#39;, &#39;duestartdatedesc&#39;, &#39;alldates&#39;, &#39;alldatesdesc&#39;, &#39;completedDateDESC&#39;, &#39;flattenedtasklist&#39;, &#39;company&#39; or &#39;dateupdated&#39;\&quot;.  | 
 **getSubTasks** | **string** | Subtasks can be excluded from the results by adding this parameter with no as the value. | 
 **nestSubTasks** | **string** | no Subtasks can be nested within the parent task object by adding this paramter with yes as the value. | 
 **getFiles** | **bool** | Files attached to tasks can be returned within the task object by setting this parameter to true. | 
 **includeToday** | **bool** | When using the filter option with any of the following options; within7,within14,within30,within365. You can choose to exclude deadlines for today by passing this parameter as false.   | 
 **ignoreStartDates** | **bool** | When using the filter option, you can choose to include start dates matching the filtering critera by passing this parameter as true. By default, only due dates are checked against the filter. | 
 **tagIds** | **string** | A comma separated list of tag ids to filter tasks on | 
 **dateupdatedASC** | **string** | Tasks can be sorted by the following options: &#39;duedate&#39;, duedateDESC&#39;, &#39;startdate&#39;, &#39;dateadded&#39;, &#39;priority&#39;, &#39;project&#39; ,&#39;manual&#39;, &#39;duestartdate&#39;, &#39;duestartdatedesc&#39;, &#39;alldates&#39;, &#39;alldatesdesc&#39;, &#39;completedDateDESC&#39;, &#39;flattenedtasklist&#39;, &#39;company&#39; or &#39;dateupdated&#39;\&quot;.  | 

### Return type

[**InlineResponse20097**](InlineResponse20097.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETTasksIdDependenciesJson

> InlineResponse200102 GETTasksIdDependenciesJson(ctx, id).Execute()

Retrieve Task Dependencies



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.GETTasksIdDependenciesJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.GETTasksIdDependenciesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETTasksIdDependenciesJson`: InlineResponse200102
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.GETTasksIdDependenciesJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETTasksIdDependenciesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200102**](InlineResponse200102.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETTasksIdFollowersJson

> map[string]interface{} GETTasksIdFollowersJson(ctx, id).ReturnUserInfo(returnUserInfo).Execute()

Get Task Followers



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
    id := "id_example" // string | 
    returnUserInfo := true // bool | Optionally, you can pass a returnUserInfo=true paramter to expand on the user information. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.GETTasksIdFollowersJson(context.Background(), id).ReturnUserInfo(returnUserInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.GETTasksIdFollowersJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETTasksIdFollowersJson`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.GETTasksIdFollowersJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETTasksIdFollowersJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnUserInfo** | **bool** | Optionally, you can pass a returnUserInfo&#x3D;true paramter to expand on the user information. | 

### Return type

**map[string]interface{}**

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETTasksIdJson

> InlineResponse200100 GETTasksIdJson(ctx, id).GetFiles(getFiles).NestSubTasks(nestSubTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).Execute()

Retrieve a Task



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
    id := "id_example" // string | 
    getFiles := true // bool | Files attached to tasks can be returned within the task object by setting this parameter to true. (optional)
    nestSubTasks := true // bool | Sub tasks can be nested within each returned task object by setting this parameter to true. (optional)
    includeCompletedSubtasks := true // bool | Used in conjunction with nestSubTasks, this parameter can be used to include completed subtasks by setting it to true. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.GETTasksIdJson(context.Background(), id).GetFiles(getFiles).NestSubTasks(nestSubTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.GETTasksIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETTasksIdJson`: InlineResponse200100
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.GETTasksIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETTasksIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getFiles** | **bool** | Files attached to tasks can be returned within the task object by setting this parameter to true. | 
 **nestSubTasks** | **bool** | Sub tasks can be nested within each returned task object by setting this parameter to true. | 
 **includeCompletedSubtasks** | **bool** | Used in conjunction with nestSubTasks, this parameter can be used to include completed subtasks by setting it to true. | 

### Return type

[**InlineResponse200100**](InlineResponse200100.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETTasksIdPredecessorsJson

> InlineResponse200104 GETTasksIdPredecessorsJson(ctx, id).Execute()

Get Task Predecessors



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.GETTasksIdPredecessorsJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.GETTasksIdPredecessorsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETTasksIdPredecessorsJson`: InlineResponse200104
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.GETTasksIdPredecessorsJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETTasksIdPredecessorsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200104**](InlineResponse200104.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETTasksJson

> InlineResponse20099 GETTasksJson(ctx).Filter(filter).Page(page).PageSize(pageSize).StartDate(startDate).EndDate(endDate).UpdatedAfterDate(updatedAfterDate).CompletedAfterDate(completedAfterDate).CompletedBeforeDate(completedBeforeDate).ShowDeleted(showDeleted).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).CreatorIds(creatorIds).Include(include).ResponsiblePartyIds(responsiblePartyIds).Sort(sort).GetSubTasks(getSubTasks).NestSubTasks(nestSubTasks).GetFiles(getFiles).IncludeToday(includeToday).IgnoreStartDates(ignoreStartDates).TagIds(tagIds).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeArchivedProjects(includeArchivedProjects).DateupdatedASC(dateupdatedASC).Execute()

Get all Tasks across all Projects



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
    filter := "filter_example" // string | Tasks can be filtered by due dates using the following options: all anytime overdue today tomorrow thisweek within7 within14 within30 within365 nodate nostartdate completed Additionally, you can choose to include the start dates in this critera by setting ignore-start-dates to false. (optional)
    page := int32(56) // int32 | Optionally, you can set the page from which to start retrieving results. This is usually used in conjunction with the parameter pageSize. For example: ?page=2&pageSize=10 will retrieve results 10-20. (optional)
    pageSize := int32(56) // int32 | The amount of tasks returned can be limited using this parameter. Normally used in conjunction with the page parameter. (optional)
    startDate := "startDate_example" // string | Tasks within a range of dates can be returned by setting a startdate and enddate. The format should be YYYYMMDD. For example: ?startdate=20140512&enddate=20140513 will get all of the tasks from the 12th of May 2014 till the 13th of May 2014.   By default ignore-start-dates=false and that means we list:  - tasks started before or on specified startDate AND due before/on specified endDate - OR no start date but due date between specified dates - OR no duedate set but startdate has passed/is specified startDate - OR milestone due date set and its between specified dates (optional)
    endDate := "endDate_example" // string | Must be used in conjunction with startdate, see above. (optional)
    updatedAfterDate := "updatedAfterDate_example" // string | Will only return tasks that have been updated after a specified date. Timestamp must be in the following format: YYYYMMDDHHMMSS. For example: ?updatedAfterDate=20131123201022 will only return tasks updated after 23rd of November 2013 at 20:10:22. (UTC) (optional)
    completedAfterDate := "completedAfterDate_example" // string | Will only return tasks that have been completed after a specified date. Timestamp must be in the following format: YYYYMMDDHHMMSS For example: ?completedAfterDate=20141123201022 will only return tasks completed after 23rd of November 2014 at 20:10:22. (UTC) (optional)
    completedBeforeDate := "completedBeforeDate_example" // string | Will only return tasks that have been completed before a specified date. Timestamp must be in the following format: YYYYMMDDHHMMSS For example: ?completedBeforeDate=20150101235959 will only return tasks completed before 1st January 2015 at 23:59:59. (UTC) (optional)
    showDeleted := "showDeleted_example" // string | Tasks that have been deleted can be shown by setting this parameter to yes. (optional)
    includeCompletedTasks := true // bool | Tasks that have been marked as completed can be shown by setting this parameter to true. (optional)
    includeCompletedSubtasks := true // bool | Sub-tasks that have been marked as completed can be shown by setting this parameter to true if you have requested to include sub-tasks (optional)
    creatorIds := "creatorIds_example" // string | For requesting tasks made by a specific person or people. For example: 44 would return tasks made by user 44. 44,45 would return tasks made by users 44 and/or 45 etc. (optional)
    include := "include_example" // string | Extra tasks that can be included with the filter option. nodate nostartdate noduedate overdue (optional)
    responsiblePartyIds := "responsiblePartyIds_example" // string | Tasks can be filtered by the person/people a task is assigned to. For example: -1 would return all tasks with an assigned person. 0 would return all tasks with no assignment. 32 would return tasks assigned to user 32. 32,55 would return tasks assigned to users 32 and/or 55 etc. (optional)
    sort := "sort_example" // string | Tasks can be sorted by the following options: 'duedate', duedateDESC', 'startdate', 'dateadded', 'priority', 'project' ,'manual', 'duestartdate', 'duestartdatedesc', 'alldates', 'alldatesdesc', 'completedDateDESC', 'flattenedtasklist', 'company' or 'dateupdated'\".   (optional)
    getSubTasks := "getSubTasks_example" // string | Subtasks can be excluded from the results by adding this parameter with no as the value. (optional)
    nestSubTasks := "nestSubTasks_example" // string | Subtasks can be nested within the parent task object by adding this parameter with yes as the value. (optional)
    getFiles := true // bool | Files attached to tasks can be returned within the task object by setting this parameter to true. (optional)
    includeToday := true // bool | When using the filter option with any of the following options; within7,within14,within30,within365. You can choose to exclude deadlines for today by passing this parameter as false.   (optional)
    ignoreStartDates := true // bool | When using the filter option, you can choose to include start dates matching the filtering critera by passing this parameter as true. By default, only due dates are checked against the filter. (optional)
    tagIds := "tagIds_example" // string | A comma separated list of tag ids to filter tasks on (optional)
    includeTasksWithoutDueDates := true // bool | Allows you to get back tasks with no due date if set to false. (optional) (default to true)
    includeTasksFromDeletedLists := true // bool | Retrieve tasks from deleted task lists. (optional)
    includeArchivedProjects := true // bool |  (optional)
    dateupdatedASC := "dateupdatedASC_example" // string | Tasks can be sorted by the following options: 'duedate', duedateDESC', 'startdate', 'dateadded', 'priority', 'project' ,'manual', 'duestartdate', 'duestartdatedesc', 'alldates', 'alldatesdesc', 'completedDateDESC', 'flattenedtasklist', 'company' or 'dateupdated'\". (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.GETTasksJson(context.Background()).Filter(filter).Page(page).PageSize(pageSize).StartDate(startDate).EndDate(endDate).UpdatedAfterDate(updatedAfterDate).CompletedAfterDate(completedAfterDate).CompletedBeforeDate(completedBeforeDate).ShowDeleted(showDeleted).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).CreatorIds(creatorIds).Include(include).ResponsiblePartyIds(responsiblePartyIds).Sort(sort).GetSubTasks(getSubTasks).NestSubTasks(nestSubTasks).GetFiles(getFiles).IncludeToday(includeToday).IgnoreStartDates(ignoreStartDates).TagIds(tagIds).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeArchivedProjects(includeArchivedProjects).DateupdatedASC(dateupdatedASC).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.GETTasksJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETTasksJson`: InlineResponse20099
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.GETTasksJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETTasksJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Tasks can be filtered by due dates using the following options: all anytime overdue today tomorrow thisweek within7 within14 within30 within365 nodate nostartdate completed Additionally, you can choose to include the start dates in this critera by setting ignore-start-dates to false. | 
 **page** | **int32** | Optionally, you can set the page from which to start retrieving results. This is usually used in conjunction with the parameter pageSize. For example: ?page&#x3D;2&amp;pageSize&#x3D;10 will retrieve results 10-20. | 
 **pageSize** | **int32** | The amount of tasks returned can be limited using this parameter. Normally used in conjunction with the page parameter. | 
 **startDate** | **string** | Tasks within a range of dates can be returned by setting a startdate and enddate. The format should be YYYYMMDD. For example: ?startdate&#x3D;20140512&amp;enddate&#x3D;20140513 will get all of the tasks from the 12th of May 2014 till the 13th of May 2014.   By default ignore-start-dates&#x3D;false and that means we list:  - tasks started before or on specified startDate AND due before/on specified endDate - OR no start date but due date between specified dates - OR no duedate set but startdate has passed/is specified startDate - OR milestone due date set and its between specified dates | 
 **endDate** | **string** | Must be used in conjunction with startdate, see above. | 
 **updatedAfterDate** | **string** | Will only return tasks that have been updated after a specified date. Timestamp must be in the following format: YYYYMMDDHHMMSS. For example: ?updatedAfterDate&#x3D;20131123201022 will only return tasks updated after 23rd of November 2013 at 20:10:22. (UTC) | 
 **completedAfterDate** | **string** | Will only return tasks that have been completed after a specified date. Timestamp must be in the following format: YYYYMMDDHHMMSS For example: ?completedAfterDate&#x3D;20141123201022 will only return tasks completed after 23rd of November 2014 at 20:10:22. (UTC) | 
 **completedBeforeDate** | **string** | Will only return tasks that have been completed before a specified date. Timestamp must be in the following format: YYYYMMDDHHMMSS For example: ?completedBeforeDate&#x3D;20150101235959 will only return tasks completed before 1st January 2015 at 23:59:59. (UTC) | 
 **showDeleted** | **string** | Tasks that have been deleted can be shown by setting this parameter to yes. | 
 **includeCompletedTasks** | **bool** | Tasks that have been marked as completed can be shown by setting this parameter to true. | 
 **includeCompletedSubtasks** | **bool** | Sub-tasks that have been marked as completed can be shown by setting this parameter to true if you have requested to include sub-tasks | 
 **creatorIds** | **string** | For requesting tasks made by a specific person or people. For example: 44 would return tasks made by user 44. 44,45 would return tasks made by users 44 and/or 45 etc. | 
 **include** | **string** | Extra tasks that can be included with the filter option. nodate nostartdate noduedate overdue | 
 **responsiblePartyIds** | **string** | Tasks can be filtered by the person/people a task is assigned to. For example: -1 would return all tasks with an assigned person. 0 would return all tasks with no assignment. 32 would return tasks assigned to user 32. 32,55 would return tasks assigned to users 32 and/or 55 etc. | 
 **sort** | **string** | Tasks can be sorted by the following options: &#39;duedate&#39;, duedateDESC&#39;, &#39;startdate&#39;, &#39;dateadded&#39;, &#39;priority&#39;, &#39;project&#39; ,&#39;manual&#39;, &#39;duestartdate&#39;, &#39;duestartdatedesc&#39;, &#39;alldates&#39;, &#39;alldatesdesc&#39;, &#39;completedDateDESC&#39;, &#39;flattenedtasklist&#39;, &#39;company&#39; or &#39;dateupdated&#39;\&quot;.   | 
 **getSubTasks** | **string** | Subtasks can be excluded from the results by adding this parameter with no as the value. | 
 **nestSubTasks** | **string** | Subtasks can be nested within the parent task object by adding this parameter with yes as the value. | 
 **getFiles** | **bool** | Files attached to tasks can be returned within the task object by setting this parameter to true. | 
 **includeToday** | **bool** | When using the filter option with any of the following options; within7,within14,within30,within365. You can choose to exclude deadlines for today by passing this parameter as false.   | 
 **ignoreStartDates** | **bool** | When using the filter option, you can choose to include start dates matching the filtering critera by passing this parameter as true. By default, only due dates are checked against the filter. | 
 **tagIds** | **string** | A comma separated list of tag ids to filter tasks on | 
 **includeTasksWithoutDueDates** | **bool** | Allows you to get back tasks with no due date if set to false. | [default to true]
 **includeTasksFromDeletedLists** | **bool** | Retrieve tasks from deleted task lists. | 
 **includeArchivedProjects** | **bool** |  | 
 **dateupdatedASC** | **string** | Tasks can be sorted by the following options: &#39;duedate&#39;, duedateDESC&#39;, &#39;startdate&#39;, &#39;dateadded&#39;, &#39;priority&#39;, &#39;project&#39; ,&#39;manual&#39;, &#39;duestartdate&#39;, &#39;duestartdatedesc&#39;, &#39;alldates&#39;, &#39;alldatesdesc&#39;, &#39;completedDateDESC&#39;, &#39;flattenedtasklist&#39;, &#39;company&#39; or &#39;dateupdated&#39;\&quot;. | 

### Return type

[**InlineResponse20099**](InlineResponse20099.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETTasksParentTaskIdSubtasksJson

> InlineResponse200108 GETTasksParentTaskIdSubtasksJson(ctx, parentTaskId).Execute()

Get Sub Tasks of a Task



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
    parentTaskId := "parentTaskId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.GETTasksParentTaskIdSubtasksJson(context.Background(), parentTaskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.GETTasksParentTaskIdSubtasksJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETTasksParentTaskIdSubtasksJson`: InlineResponse200108
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.GETTasksParentTaskIdSubtasksJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentTaskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETTasksParentTaskIdSubtasksJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200108**](InlineResponse200108.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETTasksTaskIdRecurringJson

> InlineResponse200110 GETTasksTaskIdRecurringJson(ctx, taskId).OnlyFutureDates(onlyFutureDates).Execute()

Get Recurring Tasks related to original Task.



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
    taskId := "taskId_example" // string | 
    onlyFutureDates := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.GETTasksTaskIdRecurringJson(context.Background(), taskId).OnlyFutureDates(onlyFutureDates).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.GETTasksTaskIdRecurringJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETTasksTaskIdRecurringJson`: InlineResponse200110
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.GETTasksTaskIdRecurringJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETTasksTaskIdRecurringJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **onlyFutureDates** | **bool** |  | 

### Return type

[**InlineResponse200110**](InlineResponse200110.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsProjIdTasksJson

> InlineResponse2016 POSTProjectsProjIdTasksJson(ctx, projId).Body(body).Execute()

Create a Task on a Project



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
    projId := "projId_example" // string | 
    body := *openapiclient.NewInlineObject77() // InlineObject77 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.POSTProjectsProjIdTasksJson(context.Background(), projId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.POSTProjectsProjIdTasksJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsProjIdTasksJson`: InlineResponse2016
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.POSTProjectsProjIdTasksJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsProjIdTasksJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject77**](InlineObject77.md) |  | 

### Return type

[**InlineResponse2016**](InlineResponse2016.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsProjidTasksQuickaddJson

> InlineResponse2017 POSTProjectsProjidTasksQuickaddJson(ctx, projid).Body(body).Execute()

Quickly add multiple Tasks



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
    projid := "projid_example" // string | 
    body := *openapiclient.NewInlineObject81("Content_example") // InlineObject81 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.POSTProjectsProjidTasksQuickaddJson(context.Background(), projid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.POSTProjectsProjidTasksQuickaddJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsProjidTasksQuickaddJson`: InlineResponse2017
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.POSTProjectsProjidTasksQuickaddJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsProjidTasksQuickaddJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject81**](InlineObject81.md) |  | 

### Return type

[**InlineResponse2017**](InlineResponse2017.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTTasklistsIdTasksJson

> InlineResponse201 POSTTasklistsIdTasksJson(ctx, id).Body(body).Execute()

Create a Task



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
    id := "id_example" // string | 
    body := *openapiclient.NewInlineObject91(*openapiclient.NewTasklistsIdTasksJsonTodoItem("Content_example")) // InlineObject91 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.POSTTasklistsIdTasksJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.POSTTasklistsIdTasksJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTTasklistsIdTasksJson`: InlineResponse201
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.POSTTasklistsIdTasksJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTTasklistsIdTasksJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject91**](InlineObject91.md) |  | 

### Return type

[**InlineResponse201**](InlineResponse201.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTTasksIdJson

> InlineResponse20018 POSTTasksIdJson(ctx, id).Body(body).Execute()

Create a Sub Task



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
    id := "id_example" // string | 
    body := *openapiclient.NewInlineObject95() // InlineObject95 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.POSTTasksIdJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.POSTTasksIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTTasksIdJson`: InlineResponse20018
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.POSTTasksIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTTasksIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject95**](InlineObject95.md) |  | 

### Return type

[**InlineResponse20018**](InlineResponse20018.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTTasksIdQuickaddJson

> InlineResponse2017 POSTTasksIdQuickaddJson(ctx, id).Body(body).Execute()

Quickly add multiple Sub Tasks



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
    id := "id_example" // string | 
    body := *openapiclient.NewInlineObject99() // InlineObject99 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.POSTTasksIdQuickaddJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.POSTTasksIdQuickaddJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTTasksIdQuickaddJson`: InlineResponse2017
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.POSTTasksIdQuickaddJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTTasksIdQuickaddJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject99**](InlineObject99.md) |  | 

### Return type

[**InlineResponse2017**](InlineResponse2017.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTTasklistsIdTasksReorderJson

> InlineResponse2001 PUTTasklistsIdTasksReorderJson(ctx, id).Body(body).Execute()

Reorder the Tasks



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
    id := "id_example" // string | 
    body := *openapiclient.NewInlineObject92() // InlineObject92 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.PUTTasklistsIdTasksReorderJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.PUTTasklistsIdTasksReorderJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTTasklistsIdTasksReorderJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.PUTTasklistsIdTasksReorderJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTTasklistsIdTasksReorderJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject92**](InlineObject92.md) |  | 

### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTTasksIdCompleteJson

> InlineResponse2001 PUTTasksIdCompleteJson(ctx, id).Execute()

Mark a Task complete



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.PUTTasksIdCompleteJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.PUTTasksIdCompleteJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTTasksIdCompleteJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.PUTTasksIdCompleteJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTTasksIdCompleteJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTTasksIdJson

> PUTTasksIdJson(ctx, id).Body(body).Execute()

Update a Task



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
    id := "id_example" // string | 
    body := *openapiclient.NewInlineObject94() // InlineObject94 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.PUTTasksIdJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.PUTTasksIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTTasksIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject94**](InlineObject94.md) |  | 

### Return type

 (empty response body)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTTasksIdUncompleteJson

> InlineResponse2001 PUTTasksIdUncompleteJson(ctx, id).Execute()

Mark a Task uncomplete



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.PUTTasksIdUncompleteJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.PUTTasksIdUncompleteJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTTasksIdUncompleteJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.PUTTasksIdUncompleteJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTTasksIdUncompleteJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTTasksTaskIdCopyJson

> InlineResponse200109 PUTTasksTaskIdCopyJson(ctx, taskId).Body(body).Execute()

Copy a Task from one Project to Another



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
    taskId := "taskId_example" // string | 
    body := *openapiclient.NewInlineObject103() // InlineObject103 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.PUTTasksTaskIdCopyJson(context.Background(), taskId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.PUTTasksTaskIdCopyJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTTasksTaskIdCopyJson`: InlineResponse200109
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.PUTTasksTaskIdCopyJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTTasksTaskIdCopyJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject103**](InlineObject103.md) |  | 

### Return type

[**InlineResponse200109**](InlineResponse200109.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTTasksTaskIdJson

> InlineResponse2001 PUTTasksTaskIdJson(ctx, taskId).Body(body).Execute()

Editing Task Predecessors



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
    taskId := "taskId_example" // string | 
    body := *openapiclient.NewInlineObject102() // InlineObject102 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.PUTTasksTaskIdJson(context.Background(), taskId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.PUTTasksTaskIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTTasksTaskIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.PUTTasksTaskIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTTasksTaskIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject102**](InlineObject102.md) |  | 

### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTTasksTaskIdMoveJson

> InlineResponse200109 PUTTasksTaskIdMoveJson(ctx, taskId).Body(body).Execute()

Move a Task from one Project to Another



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
    taskId := "taskId_example" // string | 
    body := *openapiclient.NewInlineObject104() // InlineObject104 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.PUTTasksTaskIdMoveJson(context.Background(), taskId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.PUTTasksTaskIdMoveJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTTasksTaskIdMoveJson`: InlineResponse200109
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.PUTTasksTaskIdMoveJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTTasksTaskIdMoveJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject104**](InlineObject104.md) |  | 

### Return type

[**InlineResponse200109**](InlineResponse200109.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

