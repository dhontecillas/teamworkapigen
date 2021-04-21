# \SummaryApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3ProjectsprojectIdSummaryJson**](SummaryApi.md#GETProjectsApiV3ProjectsprojectIdSummaryJson) | **Get** /projects/api/v3/projects/{projectId}/summary.json | Get project summary dashboard
[**GETProjectsApiV3SummaryJson**](SummaryApi.md#GETProjectsApiV3SummaryJson) | **Get** /projects/api/v3/summary.json | Get installation summary dashboard



## GETProjectsApiV3ProjectsprojectIdSummaryJson

> SummaryProjectResponse GETProjectsApiV3ProjectsprojectIdSummaryJson(ctx, projectId).Until(until).TimeRangeStart(timeRangeStart).TimeRangeEnd(timeRangeEnd).TimeRange(timeRange).Status(status).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).ProjectHealths(projectHealths).ProjectHealth(projectHealth).Health(health).EventsDaysAhead(eventsDaysAhead).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UnreadMessagesMineOnly(unreadMessagesMineOnly).UnreadCommentsMineOnly(unreadCommentsMineOnly).OnlyStarredProjects(onlyStarredProjects).OnlyProjectEvents(onlyProjectEvents).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).IncludeAssigneeTeams(includeAssigneeTeams).IncludeArchivedProjects(includeArchivedProjects).EventsAttendingOnly(eventsAttendingOnly).ApplyTaskAssigneeUsersToSince(applyTaskAssigneeUsersToSince).ApplySinceOnUnread(applySinceOnUnread).UserIds(userIds).TimeLoggedByUserIds(timeLoggedByUserIds).TaskAssigneeUserIds(taskAssigneeUserIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).MilestoneAssigneeUserIds(milestoneAssigneeUserIds).FilterTagIds(filterTagIds).FieldsUnread(fieldsUnread).FieldsTime(fieldsTime).FieldsTasks(fieldsTasks).FieldsSince(fieldsSince).FieldsRisks(fieldsRisks).FieldsMilestones(fieldsMilestones).FieldsHealth(fieldsHealth).FieldsEvents(fieldsEvents).FieldsColumns(fieldsColumns).CompanyIds(companyIds).AssigneeUserIds(assigneeUserIds).Execute()

Get project summary dashboard



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
    until := "until_example" // string | used to limit the since counters to a specific period (optional)
    timeRangeStart := time.Now() // time.Time | filter by start datetime (optional)
    timeRangeEnd := time.Now() // time.Time | filter by end datetime (optional)
    timeRange := "timeRange_example" // string | filter by time range. It will be ignored if timeRangeStart and timeRangeEnd were informed. (optional)
    status := "status_example" // string | project status (deprecated, use projectStatuses) (optional)
    projectStatuses := "projectStatuses_example" // string | project statuses (optional)
    projectStatus := "projectStatus_example" // string | project status (optional)
    projectHealths := int32(56) // int32 | project healths  0: not set 1: bad 2: ok 3: good (optional)
    projectHealth := int32(56) // int32 | project health (deprecated, use projectHealths)  0: not set 1: bad 2: ok 3: good (optional)
    health := int32(56) // int32 | project healths (deprecated, use projectHealths)  0: not set 1: bad 2: ok 3: good (optional)
    eventsDaysAhead := int32(56) // int32 | number of days remaining to event's start (optional) (default to 7)
    useStartDatesForTodaysTasks := true // bool | enforce today is considered as start date for today's tasks (optional)
    unreadMessagesMineOnly := true // bool | filter by my unread messages only (optional)
    unreadCommentsMineOnly := true // bool | filter by my unread comments only (optional)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    onlyProjectEvents := true // bool | filter only project events (optional)
    matchAllTags := true // bool | enforce all tag ids must be matched (deprecated, use matchAllProjectTags) (optional)
    matchAllProjectTags := true // bool | enforce all tag ids must be matched (optional)
    includeAssigneeTeams := true // bool | include teams related to the taskAssigneeUserIds (optional)
    includeArchivedProjects := true // bool | include archived projects (optional)
    eventsAttendingOnly := true // bool | count only events that the users are attending (optional)
    applyTaskAssigneeUsersToSince := true // bool | the since section will only use taskAssigneeUserIds if this flag is true (keeps backward compatibility) (optional)
    applySinceOnUnread := true // bool | when enabled unread counters will respect the time range period. (optional) (default to false)
    userIds := []int32{int32(123)} // []int32 | filter by user ids (optional)
    timeLoggedByUserIds := []int32{int32(123)} // []int32 | filter by user ids who time logged (optional)
    taskAssigneeUserIds := []int32{int32(123)} // []int32 | filter by user ids with tasks assigned (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    milestoneAssigneeUserIds := []int32{int32(123)} // []int32 | filter by user ids with milestones assigned (optional)
    filterTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (deprecated, use projectTagIds) (optional)
    fieldsUnread := []string{"Inner_example"} // []string |  (optional)
    fieldsTime := []string{"Inner_example"} // []string |  (optional)
    fieldsTasks := []string{"Inner_example"} // []string | sparse fields (optional)
    fieldsSince := []string{"Inner_example"} // []string |  (optional)
    fieldsRisks := []string{"Inner_example"} // []string |  (optional)
    fieldsMilestones := []string{"Inner_example"} // []string |  (optional)
    fieldsHealth := []string{"Inner_example"} // []string |  0: not set 1: bad 2: ok 3: good (optional)
    fieldsEvents := []string{"Inner_example"} // []string |  (optional)
    fieldsColumns := []string{"Inner_example"} // []string |  (optional)
    companyIds := []int32{int32(123)} // []int32 | filter by company ids (deprecated, use projectCompanyIds) (optional)
    assigneeUserIds := []int32{int32(123)} // []int32 | filter by user ids with tasks or milestones assigned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SummaryApi.GETProjectsApiV3ProjectsprojectIdSummaryJson(context.Background(), projectId).Until(until).TimeRangeStart(timeRangeStart).TimeRangeEnd(timeRangeEnd).TimeRange(timeRange).Status(status).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).ProjectHealths(projectHealths).ProjectHealth(projectHealth).Health(health).EventsDaysAhead(eventsDaysAhead).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UnreadMessagesMineOnly(unreadMessagesMineOnly).UnreadCommentsMineOnly(unreadCommentsMineOnly).OnlyStarredProjects(onlyStarredProjects).OnlyProjectEvents(onlyProjectEvents).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).IncludeAssigneeTeams(includeAssigneeTeams).IncludeArchivedProjects(includeArchivedProjects).EventsAttendingOnly(eventsAttendingOnly).ApplyTaskAssigneeUsersToSince(applyTaskAssigneeUsersToSince).ApplySinceOnUnread(applySinceOnUnread).UserIds(userIds).TimeLoggedByUserIds(timeLoggedByUserIds).TaskAssigneeUserIds(taskAssigneeUserIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).MilestoneAssigneeUserIds(milestoneAssigneeUserIds).FilterTagIds(filterTagIds).FieldsUnread(fieldsUnread).FieldsTime(fieldsTime).FieldsTasks(fieldsTasks).FieldsSince(fieldsSince).FieldsRisks(fieldsRisks).FieldsMilestones(fieldsMilestones).FieldsHealth(fieldsHealth).FieldsEvents(fieldsEvents).FieldsColumns(fieldsColumns).CompanyIds(companyIds).AssigneeUserIds(assigneeUserIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SummaryApi.GETProjectsApiV3ProjectsprojectIdSummaryJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ProjectsprojectIdSummaryJson`: SummaryProjectResponse
    fmt.Fprintf(os.Stdout, "Response from `SummaryApi.GETProjectsApiV3ProjectsprojectIdSummaryJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsprojectIdSummaryJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **until** | **string** | used to limit the since counters to a specific period | 
 **timeRangeStart** | **time.Time** | filter by start datetime | 
 **timeRangeEnd** | **time.Time** | filter by end datetime | 
 **timeRange** | **string** | filter by time range. It will be ignored if timeRangeStart and timeRangeEnd were informed. | 
 **status** | **string** | project status (deprecated, use projectStatuses) | 
 **projectStatuses** | **string** | project statuses | 
 **projectStatus** | **string** | project status | 
 **projectHealths** | **int32** | project healths  0: not set 1: bad 2: ok 3: good | 
 **projectHealth** | **int32** | project health (deprecated, use projectHealths)  0: not set 1: bad 2: ok 3: good | 
 **health** | **int32** | project healths (deprecated, use projectHealths)  0: not set 1: bad 2: ok 3: good | 
 **eventsDaysAhead** | **int32** | number of days remaining to event&#39;s start | [default to 7]
 **useStartDatesForTodaysTasks** | **bool** | enforce today is considered as start date for today&#39;s tasks | 
 **unreadMessagesMineOnly** | **bool** | filter by my unread messages only | 
 **unreadCommentsMineOnly** | **bool** | filter by my unread comments only | 
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **onlyProjectEvents** | **bool** | filter only project events | 
 **matchAllTags** | **bool** | enforce all tag ids must be matched (deprecated, use matchAllProjectTags) | 
 **matchAllProjectTags** | **bool** | enforce all tag ids must be matched | 
 **includeAssigneeTeams** | **bool** | include teams related to the taskAssigneeUserIds | 
 **includeArchivedProjects** | **bool** | include archived projects | 
 **eventsAttendingOnly** | **bool** | count only events that the users are attending | 
 **applyTaskAssigneeUsersToSince** | **bool** | the since section will only use taskAssigneeUserIds if this flag is true (keeps backward compatibility) | 
 **applySinceOnUnread** | **bool** | when enabled unread counters will respect the time range period. | [default to false]
 **userIds** | **[]int32** | filter by user ids | 
 **timeLoggedByUserIds** | **[]int32** | filter by user ids who time logged | 
 **taskAssigneeUserIds** | **[]int32** | filter by user ids with tasks assigned | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **milestoneAssigneeUserIds** | **[]int32** | filter by user ids with milestones assigned | 
 **filterTagIds** | **[]int32** | filter by project tag ids (deprecated, use projectTagIds) | 
 **fieldsUnread** | **[]string** |  | 
 **fieldsTime** | **[]string** |  | 
 **fieldsTasks** | **[]string** | sparse fields | 
 **fieldsSince** | **[]string** |  | 
 **fieldsRisks** | **[]string** |  | 
 **fieldsMilestones** | **[]string** |  | 
 **fieldsHealth** | **[]string** |  0: not set 1: bad 2: ok 3: good | 
 **fieldsEvents** | **[]string** |  | 
 **fieldsColumns** | **[]string** |  | 
 **companyIds** | **[]int32** | filter by company ids (deprecated, use projectCompanyIds) | 
 **assigneeUserIds** | **[]int32** | filter by user ids with tasks or milestones assigned | 

### Return type

[**SummaryProjectResponse**](SummaryProjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3SummaryJson

> SummaryResponse GETProjectsApiV3SummaryJson(ctx).Until(until).TimeRangeStart(timeRangeStart).TimeRangeEnd(timeRangeEnd).TimeRange(timeRange).Status(status).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).ProjectHealths(projectHealths).ProjectHealth(projectHealth).EventsDaysAhead(eventsDaysAhead).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UnreadMessagesMineOnly(unreadMessagesMineOnly).UnreadCommentsMineOnly(unreadCommentsMineOnly).OnlyStarredProjects(onlyStarredProjects).OnlyProjectEvents(onlyProjectEvents).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).IncludeAssigneeTeams(includeAssigneeTeams).IncludeArchivedProjects(includeArchivedProjects).EventsAttendingOnly(eventsAttendingOnly).ApplyTaskAssigneeUsersToSince(applyTaskAssigneeUsersToSince).ApplySinceOnUnread(applySinceOnUnread).UserIds(userIds).TimeLoggedByUserIds(timeLoggedByUserIds).TaskAssigneeUserIds(taskAssigneeUserIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).MilestoneAssigneeUserIds(milestoneAssigneeUserIds).FilterTagIds(filterTagIds).FieldsUnread(fieldsUnread).FieldsTime(fieldsTime).FieldsTasks(fieldsTasks).FieldsSince(fieldsSince).FieldsRisks(fieldsRisks).FieldsMilestones(fieldsMilestones).FieldsHealth(fieldsHealth).FieldsEvents(fieldsEvents).FieldsColumns(fieldsColumns).CompanyIds(companyIds).AssigneeUserIds(assigneeUserIds).Execute()

Get installation summary dashboard



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
    until := "until_example" // string | used to limit the since counters to a specific period (optional)
    timeRangeStart := time.Now() // time.Time | filter by start datetime (optional)
    timeRangeEnd := time.Now() // time.Time | filter by end datetime (optional)
    timeRange := "timeRange_example" // string | filter by time range. It will be ignored if timeRangeStart and timeRangeEnd are provided. (optional)
    status := "status_example" // string | project status (deprecated, use projectStatuses) (optional)
    projectStatuses := "projectStatuses_example" // string | project statuses (optional)
    projectStatus := "projectStatus_example" // string | project status (optional)
    projectHealths := int32(56) // int32 | project healths  0: not set 1: bad 2: ok 3: good (optional)
    projectHealth := int32(56) // int32 | project healths (deprecated, use projectHealths)  0: not set 1: bad 2: ok 3: good (optional)
    eventsDaysAhead := int32(56) // int32 | number of days remaining to event's start (optional) (default to 7)
    useStartDatesForTodaysTasks := true // bool | enforce today is considered as start date for today's tasks (optional)
    unreadMessagesMineOnly := true // bool | filter by my unread messages only (optional)
    unreadCommentsMineOnly := true // bool | filter by my unread comments only (optional)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    onlyProjectEvents := true // bool | filter only project events (optional)
    matchAllTags := true // bool | enforce all tag ids must be matched (deprecated, use matchAllProjectTags) (optional)
    matchAllProjectTags := true // bool | enforce all tag ids must be matched (optional)
    includeAssigneeTeams := true // bool | include teams related to the taskAssigneeUserIds (optional)
    includeArchivedProjects := true // bool | include archived projects (optional)
    eventsAttendingOnly := true // bool | count only events that the users are attending (optional)
    applyTaskAssigneeUsersToSince := true // bool | the since section will only use taskAssigneeUserIds if this flag is true (keeps backward compatibility) (optional)
    applySinceOnUnread := true // bool | when enabled unread counters will respect the time range period. (optional) (default to false)
    userIds := []int32{int32(123)} // []int32 | filter by user ids (optional)
    timeLoggedByUserIds := []int32{int32(123)} // []int32 | filter by user ids who time logged (optional)
    taskAssigneeUserIds := []int32{int32(123)} // []int32 | filter by user ids with tasks assigned (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    milestoneAssigneeUserIds := []int32{int32(123)} // []int32 | filter by user ids with milestones assigned (optional)
    filterTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (deprecated, use projectTagIds) (optional)
    fieldsUnread := []string{"Inner_example"} // []string |  (optional)
    fieldsTime := []string{"Inner_example"} // []string |  (optional)
    fieldsTasks := []string{"Inner_example"} // []string | sparse fields (optional)
    fieldsSince := []string{"Inner_example"} // []string |  (optional)
    fieldsRisks := []string{"Inner_example"} // []string |  (optional)
    fieldsMilestones := []string{"Inner_example"} // []string |  (optional)
    fieldsHealth := []string{"Inner_example"} // []string |  0: not set 1: bad 2: ok 3: good (optional)
    fieldsEvents := []string{"Inner_example"} // []string |  (optional)
    fieldsColumns := []string{"Inner_example"} // []string |  (optional)
    companyIds := []int32{int32(123)} // []int32 | filter by company ids (deprecated, use projectCompanyIds) (optional)
    assigneeUserIds := []int32{int32(123)} // []int32 | filter by user ids with tasks or milestones assigned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SummaryApi.GETProjectsApiV3SummaryJson(context.Background()).Until(until).TimeRangeStart(timeRangeStart).TimeRangeEnd(timeRangeEnd).TimeRange(timeRange).Status(status).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).ProjectHealths(projectHealths).ProjectHealth(projectHealth).EventsDaysAhead(eventsDaysAhead).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UnreadMessagesMineOnly(unreadMessagesMineOnly).UnreadCommentsMineOnly(unreadCommentsMineOnly).OnlyStarredProjects(onlyStarredProjects).OnlyProjectEvents(onlyProjectEvents).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).IncludeAssigneeTeams(includeAssigneeTeams).IncludeArchivedProjects(includeArchivedProjects).EventsAttendingOnly(eventsAttendingOnly).ApplyTaskAssigneeUsersToSince(applyTaskAssigneeUsersToSince).ApplySinceOnUnread(applySinceOnUnread).UserIds(userIds).TimeLoggedByUserIds(timeLoggedByUserIds).TaskAssigneeUserIds(taskAssigneeUserIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).MilestoneAssigneeUserIds(milestoneAssigneeUserIds).FilterTagIds(filterTagIds).FieldsUnread(fieldsUnread).FieldsTime(fieldsTime).FieldsTasks(fieldsTasks).FieldsSince(fieldsSince).FieldsRisks(fieldsRisks).FieldsMilestones(fieldsMilestones).FieldsHealth(fieldsHealth).FieldsEvents(fieldsEvents).FieldsColumns(fieldsColumns).CompanyIds(companyIds).AssigneeUserIds(assigneeUserIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SummaryApi.GETProjectsApiV3SummaryJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3SummaryJson`: SummaryResponse
    fmt.Fprintf(os.Stdout, "Response from `SummaryApi.GETProjectsApiV3SummaryJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3SummaryJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **until** | **string** | used to limit the since counters to a specific period | 
 **timeRangeStart** | **time.Time** | filter by start datetime | 
 **timeRangeEnd** | **time.Time** | filter by end datetime | 
 **timeRange** | **string** | filter by time range. It will be ignored if timeRangeStart and timeRangeEnd are provided. | 
 **status** | **string** | project status (deprecated, use projectStatuses) | 
 **projectStatuses** | **string** | project statuses | 
 **projectStatus** | **string** | project status | 
 **projectHealths** | **int32** | project healths  0: not set 1: bad 2: ok 3: good | 
 **projectHealth** | **int32** | project healths (deprecated, use projectHealths)  0: not set 1: bad 2: ok 3: good | 
 **eventsDaysAhead** | **int32** | number of days remaining to event&#39;s start | [default to 7]
 **useStartDatesForTodaysTasks** | **bool** | enforce today is considered as start date for today&#39;s tasks | 
 **unreadMessagesMineOnly** | **bool** | filter by my unread messages only | 
 **unreadCommentsMineOnly** | **bool** | filter by my unread comments only | 
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **onlyProjectEvents** | **bool** | filter only project events | 
 **matchAllTags** | **bool** | enforce all tag ids must be matched (deprecated, use matchAllProjectTags) | 
 **matchAllProjectTags** | **bool** | enforce all tag ids must be matched | 
 **includeAssigneeTeams** | **bool** | include teams related to the taskAssigneeUserIds | 
 **includeArchivedProjects** | **bool** | include archived projects | 
 **eventsAttendingOnly** | **bool** | count only events that the users are attending | 
 **applyTaskAssigneeUsersToSince** | **bool** | the since section will only use taskAssigneeUserIds if this flag is true (keeps backward compatibility) | 
 **applySinceOnUnread** | **bool** | when enabled unread counters will respect the time range period. | [default to false]
 **userIds** | **[]int32** | filter by user ids | 
 **timeLoggedByUserIds** | **[]int32** | filter by user ids who time logged | 
 **taskAssigneeUserIds** | **[]int32** | filter by user ids with tasks assigned | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **milestoneAssigneeUserIds** | **[]int32** | filter by user ids with milestones assigned | 
 **filterTagIds** | **[]int32** | filter by project tag ids (deprecated, use projectTagIds) | 
 **fieldsUnread** | **[]string** |  | 
 **fieldsTime** | **[]string** |  | 
 **fieldsTasks** | **[]string** | sparse fields | 
 **fieldsSince** | **[]string** |  | 
 **fieldsRisks** | **[]string** |  | 
 **fieldsMilestones** | **[]string** |  | 
 **fieldsHealth** | **[]string** |  0: not set 1: bad 2: ok 3: good | 
 **fieldsEvents** | **[]string** |  | 
 **fieldsColumns** | **[]string** |  | 
 **companyIds** | **[]int32** | filter by company ids (deprecated, use projectCompanyIds) | 
 **assigneeUserIds** | **[]int32** | filter by user ids with tasks or milestones assigned | 

### Return type

[**SummaryResponse**](SummaryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

