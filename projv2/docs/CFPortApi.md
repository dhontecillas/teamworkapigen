# \CFPortApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV2CompaniescompanyIdPeople**](CFPortApi.md#GETProjectsApiV2CompaniescompanyIdPeople) | **Get** /projects/api/v2/companies/{companyId}/people | Retrieves the details for all the people from the submitted company (excluding those you don’t have permission to see).
[**GETProjectsApiV2CompaniescompanyIdProjects**](CFPortApi.md#GETProjectsApiV2CompaniescompanyIdProjects) | **Get** /projects/api/v2/companies/{companyId}/projects | Returns the company projects
[**GETProjectsApiV2Defaultprojects**](CFPortApi.md#GETProjectsApiV2Defaultprojects) | **Get** /projects/api/v2/defaultprojects | Returns the project defaults
[**GETProjectsApiV2FilesobjectIdProject**](CFPortApi.md#GETProjectsApiV2FilesobjectIdProject) | **Get** /projects/api/v2/files/{objectId}/project | Get the project for the file
[**GETProjectsApiV2Fullsearch**](CFPortApi.md#GETProjectsApiV2Fullsearch) | **Get** /projects/api/v2/fullsearch | Perform a full search
[**GETProjectsApiV2LinksobjectIdProject**](CFPortApi.md#GETProjectsApiV2LinksobjectIdProject) | **Get** /projects/api/v2/links/{objectId}/project | Get the project for the link
[**GETProjectsApiV2Me**](CFPortApi.md#GETProjectsApiV2Me) | **Get** /projects/api/v2/me | Returns the details for the account you’re currently logged in with.
[**GETProjectsApiV2MePersonaltasks**](CFPortApi.md#GETProjectsApiV2MePersonaltasks) | **Get** /projects/api/v2/me/personaltasks | Get Tasks
[**GETProjectsApiV2MePersonaltaskstaskId**](CFPortApi.md#GETProjectsApiV2MePersonaltaskstaskId) | **Get** /projects/api/v2/me/personaltasks/{taskId} | Get Tasks
[**GETProjectsApiV2MessagesobjectIdProject**](CFPortApi.md#GETProjectsApiV2MessagesobjectIdProject) | **Get** /projects/api/v2/messages/{objectId}/project | Get the project for the message
[**GETProjectsApiV2MilestonesobjectIdProject**](CFPortApi.md#GETProjectsApiV2MilestonesobjectIdProject) | **Get** /projects/api/v2/milestones/{objectId}/project | Get the project for the milestone
[**GETProjectsApiV2NotebooksobjectIdProject**](CFPortApi.md#GETProjectsApiV2NotebooksobjectIdProject) | **Get** /projects/api/v2/notebooks/{objectId}/project | Get the project for the notebook
[**GETProjectsApiV2People**](CFPortApi.md#GETProjectsApiV2People) | **Get** /projects/api/v2/people | All people visible to the user will be returned, including the user themselves. By default 100 records are returned at a time.
[**GETProjectsApiV2PeopleuserId**](CFPortApi.md#GETProjectsApiV2PeopleuserId) | **Get** /projects/api/v2/people/{userId} | Retrieves the user details for the ID submitted.
[**GETProjectsApiV2PeopleuserIdProjects**](CFPortApi.md#GETProjectsApiV2PeopleuserIdProjects) | **Get** /projects/api/v2/people/{userId}/projects | APICall_GET_projects
[**GETProjectsApiV2ProjectcategoriesprojectCategoryIdTasks**](CFPortApi.md#GETProjectsApiV2ProjectcategoriesprojectCategoryIdTasks) | **Get** /projects/api/v2/projectcategories/{projectCategoryId}/tasks | Get Tasks
[**GETProjectsApiV2Projects**](CFPortApi.md#GETProjectsApiV2Projects) | **Get** /projects/api/v2/projects | APICall_GET_projects
[**GETProjectsApiV2ProjectsArchived**](CFPortApi.md#GETProjectsApiV2ProjectsArchived) | **Get** /projects/api/v2/projects/archived | Returns the project defaults
[**GETProjectsApiV2ProjectsChart**](CFPortApi.md#GETProjectsApiV2ProjectsChart) | **Get** /projects/api/v2/projects/chart | Returns the project chart
[**GETProjectsApiV2ProjectsChartoverview**](CFPortApi.md#GETProjectsApiV2ProjectsChartoverview) | **Get** /projects/api/v2/projects/chartoverview | Returns the project defaults (Objective: To deprecate getProjectsChart)
[**GETProjectsApiV2ProjectsEstimatedtimeTotal**](CFPortApi.md#GETProjectsApiV2ProjectsEstimatedtimeTotal) | **Get** /projects/api/v2/projects/estimatedtime/total | APICall_GET_projects
[**GETProjectsApiV2ProjectsLatest**](CFPortApi.md#GETProjectsApiV2ProjectsLatest) | **Get** /projects/api/v2/projects/latest | Get latest project
[**GETProjectsApiV2ProjectsMeLatest**](CFPortApi.md#GETProjectsApiV2ProjectsMeLatest) | **Get** /projects/api/v2/projects/me/latest | Returns the latest project
[**GETProjectsApiV2ProjectsNew**](CFPortApi.md#GETProjectsApiV2ProjectsNew) | **Get** /projects/api/v2/projects/new | Gets the milestone timeline for projects
[**GETProjectsApiV2ProjectsStarred**](CFPortApi.md#GETProjectsApiV2ProjectsStarred) | **Get** /projects/api/v2/projects/starred | Returns the starred projects
[**GETProjectsApiV2ProjectsTemplates**](CFPortApi.md#GETProjectsApiV2ProjectsTemplates) | **Get** /projects/api/v2/projects/templates | Returns the project defaults
[**GETProjectsApiV2ProjectsTimeTotal**](CFPortApi.md#GETProjectsApiV2ProjectsTimeTotal) | **Get** /projects/api/v2/projects/time/total | Get Projects Total Time
[**GETProjectsApiV2ProjectsprojectId**](CFPortApi.md#GETProjectsApiV2ProjectsprojectId) | **Get** /projects/api/v2/projects/{projectId} | Returns the project defaults
[**GETProjectsApiV2ProjectsprojectIdBillingTime**](CFPortApi.md#GETProjectsApiV2ProjectsprojectIdBillingTime) | **Get** /projects/api/v2/projects/{projectId}/billing/time | APICall_GET_projects_id_time_entries
[**GETProjectsApiV2ProjectsprojectIdEmailaddress**](CFPortApi.md#GETProjectsApiV2ProjectsprojectIdEmailaddress) | **Get** /projects/api/v2/projects/{projectId}/emailaddress | Returns the details for the account you’re currently logged in with.
[**GETProjectsApiV2ProjectsprojectIdPeople**](CFPortApi.md#GETProjectsApiV2ProjectsprojectIdPeople) | **Get** /projects/api/v2/projects/{projectId}/people | Retrieves all of the people in a given project.
[**GETProjectsApiV2ProjectsprojectIdPredecessors**](CFPortApi.md#GETProjectsApiV2ProjectsprojectIdPredecessors) | **Get** /projects/api/v2/projects/{projectId}/predecessors | Get Predecessor
[**GETProjectsApiV2ProjectsprojectIdSummary**](CFPortApi.md#GETProjectsApiV2ProjectsprojectIdSummary) | **Get** /projects/api/v2/projects/{projectId}/summary | APICall_GET_summary
[**GETProjectsApiV2ProjectsprojectIdTaskids**](CFPortApi.md#GETProjectsApiV2ProjectsprojectIdTaskids) | **Get** /projects/api/v2/projects/{projectId}/taskids | Get Tasks
[**GETProjectsApiV2ProjectsprojectIdTasks**](CFPortApi.md#GETProjectsApiV2ProjectsprojectIdTasks) | **Get** /projects/api/v2/projects/{projectId}/tasks | Get Tasks
[**GETProjectsApiV2ProjectsprojectIdTime**](CFPortApi.md#GETProjectsApiV2ProjectsprojectIdTime) | **Get** /projects/api/v2/projects/{projectId}/time | APICall_GET_projects_id_time_entries
[**GETProjectsApiV2ProjectsprojectIdTimeEntries**](CFPortApi.md#GETProjectsApiV2ProjectsprojectIdTimeEntries) | **Get** /projects/api/v2/projects/{projectId}/time_entries | APICall_GET_projects_id_time_entries
[**GETProjectsApiV2ProjectsprojectIdTimelogs**](CFPortApi.md#GETProjectsApiV2ProjectsprojectIdTimelogs) | **Get** /projects/api/v2/projects/{projectId}/timelogs | APICall_GET_projects_id_time_entries
[**GETProjectsApiV2ProjectsprojejctIdAssingLaterRoles**](CFPortApi.md#GETProjectsApiV2ProjectsprojejctIdAssingLaterRoles) | **Get** /projects/api/v2/projects/{projejctId}/assing_later_roles | Get assign later roles
[**GETProjectsApiV2Quicksearch**](CFPortApi.md#GETProjectsApiV2Quicksearch) | **Get** /projects/api/v2/quicksearch | APICall_POST_quicksearch
[**GETProjectsApiV2Summary**](CFPortApi.md#GETProjectsApiV2Summary) | **Get** /projects/api/v2/summary | APICall_GET_summary
[**GETProjectsApiV2TasklistsobjectIdProject**](CFPortApi.md#GETProjectsApiV2TasklistsobjectIdProject) | **Get** /projects/api/v2/tasklists/{objectId}/project | Get the project for the tasklist
[**GETProjectsApiV2TaskliststasklistIdTaskids**](CFPortApi.md#GETProjectsApiV2TaskliststasklistIdTaskids) | **Get** /projects/api/v2/tasklists/{tasklistId}/taskids | Get Tasks
[**GETProjectsApiV2TaskliststasklistIdTasks**](CFPortApi.md#GETProjectsApiV2TaskliststasklistIdTasks) | **Get** /projects/api/v2/tasklists/{tasklistId}/tasks | Get Tasks
[**GETProjectsApiV2Tasks**](CFPortApi.md#GETProjectsApiV2Tasks) | **Get** /projects/api/v2/tasks | Get Tasks
[**GETProjectsApiV2TasksdependentTaskIdPredecessors**](CFPortApi.md#GETProjectsApiV2TasksdependentTaskIdPredecessors) | **Get** /projects/api/v2/tasks/{dependentTaskId}/predecessors | Get Predecessor
[**GETProjectsApiV2TasksobjectIdProject**](CFPortApi.md#GETProjectsApiV2TasksobjectIdProject) | **Get** /projects/api/v2/tasks/{objectId}/project | Get the project for the task
[**GETProjectsApiV2TaskstaskId**](CFPortApi.md#GETProjectsApiV2TaskstaskId) | **Get** /projects/api/v2/tasks/{taskId} | Get Tasks
[**GETProjectsApiV2TaskstaskIdEstimatedtime**](CFPortApi.md#GETProjectsApiV2TaskstaskIdEstimatedtime) | **Get** /projects/api/v2/tasks/{taskId}/estimatedtime | Will return the total for all sub-tasks
[**GETProjectsApiV2TaskstaskIdSubtasks**](CFPortApi.md#GETProjectsApiV2TaskstaskIdSubtasks) | **Get** /projects/api/v2/tasks/{taskId}/subtasks | Get Tasks
[**GETProjectsApiV2TaskstaskIdTime**](CFPortApi.md#GETProjectsApiV2TaskstaskIdTime) | **Get** /projects/api/v2/tasks/{taskId}/time | APICall_GET_todo_items_id_time_entries
[**GETProjectsApiV2TaskstaskIdTimeEntries**](CFPortApi.md#GETProjectsApiV2TaskstaskIdTimeEntries) | **Get** /projects/api/v2/tasks/{taskId}/time_entries | APICall_GET_todo_items_id_time_entries
[**GETProjectsApiV2Time**](CFPortApi.md#GETProjectsApiV2Time) | **Get** /projects/api/v2/time | Get all timelogs
[**GETProjectsApiV2TimeEntries**](CFPortApi.md#GETProjectsApiV2TimeEntries) | **Get** /projects/api/v2/time_entries | Get all time entries
[**GETProjectsApiV2TimeEntriesMe**](CFPortApi.md#GETProjectsApiV2TimeEntriesMe) | **Get** /projects/api/v2/time_entries/me | APICall_GET_time_entries_currentUser
[**GETProjectsApiV2TimeEntriestimeLogId**](CFPortApi.md#GETProjectsApiV2TimeEntriestimeLogId) | **Get** /projects/api/v2/time_entries/{timeLogId} | Get time entries for a timelogId
[**GETProjectsApiV2TimeMe**](CFPortApi.md#GETProjectsApiV2TimeMe) | **Get** /projects/api/v2/time/me | APICall_GET_time_entries_currentUser
[**GETProjectsApiV2TimetimeLogId**](CFPortApi.md#GETProjectsApiV2TimetimeLogId) | **Get** /projects/api/v2/time/{timeLogId} | Get timelog for a timelogId
[**GETProjectsApiV2TodoItemstaskIdTime**](CFPortApi.md#GETProjectsApiV2TodoItemstaskIdTime) | **Get** /projects/api/v2/todo_items/{taskId}/time | APICall_GET_todo_items_id_time_entries
[**GETProjectsApiV2TodoItemstaskIdTimeEntries**](CFPortApi.md#GETProjectsApiV2TodoItemstaskIdTimeEntries) | **Get** /projects/api/v2/todo_items/{taskId}/time_entries | APICall_GET_todo_items_id_time_entries



## GETProjectsApiV2CompaniescompanyIdPeople

> PeopleResponse GETProjectsApiV2CompaniescompanyIdPeople(ctx, companyId).UpdatedAfterDateTime(updatedAfterDateTime).UpdatedAfterDate(updatedAfterDate).Type_(type_).SortOrder(sortOrder).Sort(sort).SearchTerm(searchTerm).ProjectType(projectType).InviteStatus(inviteStatus).FirstLetter(firstLetter).EmailAddress(emailAddress).DataSet(dataSet).ProjectId(projectId).ProjectCompanyId(projectCompanyId).PageSize(pageSize).Page(page).ExcludeProjectId(excludeProjectId).ShowDeleted(showDeleted).ReturnTeams(returnTeams).ReturnProjectIds(returnProjectIds).ReturnLetters(returnLetters).OwnerCompanyFirst(ownerCompanyFirst).Onlyids(onlyids).IncludeTags(includeTags).IncludeObservers(includeObservers).IncludeContacts(includeContacts).IncludeCompanyDetails(includeCompanyDetails).IncludeClockIn(includeClockIn).IncludeClients(includeClients).GroupByCompany(groupByCompany).GetProjectRoles(getProjectRoles).GetCounts(getCounts).Fullprofile(fullprofile).CountOnly(countOnly).UserIds(userIds).CompanyId2(companyId2).Execute()

Retrieves the details for all the people from the submitted company (excluding those you don’t have permission to see).

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
    companyId := int32(56) // int32 | 
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
    projectId := int32(56) // int32 |  (optional)
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
    userIds := []int32{int32(123)} // []int32 |  (optional)
    companyId2 := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2CompaniescompanyIdPeople(context.Background(), companyId).UpdatedAfterDateTime(updatedAfterDateTime).UpdatedAfterDate(updatedAfterDate).Type_(type_).SortOrder(sortOrder).Sort(sort).SearchTerm(searchTerm).ProjectType(projectType).InviteStatus(inviteStatus).FirstLetter(firstLetter).EmailAddress(emailAddress).DataSet(dataSet).ProjectId(projectId).ProjectCompanyId(projectCompanyId).PageSize(pageSize).Page(page).ExcludeProjectId(excludeProjectId).ShowDeleted(showDeleted).ReturnTeams(returnTeams).ReturnProjectIds(returnProjectIds).ReturnLetters(returnLetters).OwnerCompanyFirst(ownerCompanyFirst).Onlyids(onlyids).IncludeTags(includeTags).IncludeObservers(includeObservers).IncludeContacts(includeContacts).IncludeCompanyDetails(includeCompanyDetails).IncludeClockIn(includeClockIn).IncludeClients(includeClients).GroupByCompany(groupByCompany).GetProjectRoles(getProjectRoles).GetCounts(getCounts).Fullprofile(fullprofile).CountOnly(countOnly).UserIds(userIds).CompanyId2(companyId2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2CompaniescompanyIdPeople``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV2CompaniescompanyIdPeople`: PeopleResponse
    fmt.Fprintf(os.Stdout, "Response from `CFPortApi.GETProjectsApiV2CompaniescompanyIdPeople`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2CompaniescompanyIdPeopleRequest struct via the builder pattern


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
 **projectId** | **int32** |  | 
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
 **userIds** | **[]int32** |  | 
 **companyId2** | **[]int32** |  | 

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


## GETProjectsApiV2CompaniescompanyIdProjects

> GETProjectsApiV2CompaniescompanyIdProjects(ctx, companyId).UpdatedAfterDateTime(updatedAfterDateTime).UpdatedAfterDate(updatedAfterDate).SearchTerm(searchTerm).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).FirstLetter(firstLetter).CreatedAfterDateTime(createdAfterDateTime).CreatedAfterDate(createdAfterDate).Status(status).UserId(userId).PageSize(pageSize).Page(page).SearchCompany(searchCompany).SearchCategory(searchCategory).ReturnLetters(returnLetters).OnlyStarredProjects(onlyStarredProjects).OnlyArchivedProjects(onlyArchivedProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).IncludeUpdates(includeUpdates).IncludeTags(includeTags).IncludePeople(includePeople).IncludeArchivedProjects(includeArchivedProjects).HideDesc(hideDesc).GetTruePermissions(getTruePermissions).GetPermissions(getPermissions).GetDeleted(getDeleted).GetCategoryPath(getCategoryPath).GetActivePages(getActivePages).FormatMarkdown(formatMarkdown).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).FilterTagIds(filterTagIds).CompanyId2(companyId2).CategoryId(categoryId).CatId(catId).Execute()

Returns the company projects

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
    companyId := int32(56) // int32 | 
    updatedAfterDateTime := "updatedAfterDateTime_example" // string |  (optional)
    updatedAfterDate := "updatedAfterDate_example" // string |  (optional)
    searchTerm := "searchTerm_example" // string |  (optional)
    projectStatus := "projectStatus_example" // string |  (optional)
    orderMode := "orderMode_example" // string |  (optional)
    orderBy := "orderBy_example" // string |  (optional)
    firstLetter := "firstLetter_example" // string |  (optional)
    createdAfterDateTime := "createdAfterDateTime_example" // string |  (optional)
    createdAfterDate := "createdAfterDate_example" // string |  (optional)
    status := "status_example" // string | common project filters (optional)
    userId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    searchCompany := true // bool |  (optional)
    searchCategory := true // bool |  (optional)
    returnLetters := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    onlyArchivedProjects := true // bool |  (optional)
    matchAllTags := true // bool |  (optional)
    matchAllProjectTags := true // bool |  (optional)
    includeUpdates := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    includePeople := true // bool |  (optional)
    includeArchivedProjects := true // bool |  (optional)
    hideDesc := true // bool |  (optional)
    getTruePermissions := true // bool |  (optional)
    getPermissions := true // bool |  (optional)
    getDeleted := true // bool |  (optional)
    getCategoryPath := true // bool |  (optional)
    getActivePages := true // bool |  (optional)
    formatMarkdown := true // bool |  (optional)
    projectTagIds := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    projectHealths := []int32{int32(123)} // []int32 |  (optional)
    projectHealth := []int32{int32(123)} // []int32 |  (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 |  (optional)
    filterTagIds := []int32{int32(123)} // []int32 |  (optional)
    companyId2 := []int32{int32(123)} // []int32 |  (optional)
    categoryId := []int32{int32(123)} // []int32 |  (optional)
    catId := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2CompaniescompanyIdProjects(context.Background(), companyId).UpdatedAfterDateTime(updatedAfterDateTime).UpdatedAfterDate(updatedAfterDate).SearchTerm(searchTerm).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).FirstLetter(firstLetter).CreatedAfterDateTime(createdAfterDateTime).CreatedAfterDate(createdAfterDate).Status(status).UserId(userId).PageSize(pageSize).Page(page).SearchCompany(searchCompany).SearchCategory(searchCategory).ReturnLetters(returnLetters).OnlyStarredProjects(onlyStarredProjects).OnlyArchivedProjects(onlyArchivedProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).IncludeUpdates(includeUpdates).IncludeTags(includeTags).IncludePeople(includePeople).IncludeArchivedProjects(includeArchivedProjects).HideDesc(hideDesc).GetTruePermissions(getTruePermissions).GetPermissions(getPermissions).GetDeleted(getDeleted).GetCategoryPath(getCategoryPath).GetActivePages(getActivePages).FormatMarkdown(formatMarkdown).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).FilterTagIds(filterTagIds).CompanyId2(companyId2).CategoryId(categoryId).CatId(catId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2CompaniescompanyIdProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2CompaniescompanyIdProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedAfterDateTime** | **string** |  | 
 **updatedAfterDate** | **string** |  | 
 **searchTerm** | **string** |  | 
 **projectStatus** | **string** |  | 
 **orderMode** | **string** |  | 
 **orderBy** | **string** |  | 
 **firstLetter** | **string** |  | 
 **createdAfterDateTime** | **string** |  | 
 **createdAfterDate** | **string** |  | 
 **status** | **string** | common project filters | 
 **userId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **searchCompany** | **bool** |  | 
 **searchCategory** | **bool** |  | 
 **returnLetters** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **onlyArchivedProjects** | **bool** |  | 
 **matchAllTags** | **bool** |  | 
 **matchAllProjectTags** | **bool** |  | 
 **includeUpdates** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **includePeople** | **bool** |  | 
 **includeArchivedProjects** | **bool** |  | 
 **hideDesc** | **bool** |  | 
 **getTruePermissions** | **bool** |  | 
 **getPermissions** | **bool** |  | 
 **getDeleted** | **bool** |  | 
 **getCategoryPath** | **bool** |  | 
 **getActivePages** | **bool** |  | 
 **formatMarkdown** | **bool** |  | 
 **projectTagIds** | **[]int32** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **projectHealths** | **[]int32** |  | 
 **projectHealth** | **[]int32** |  | 
 **projectCompanyIds** | **[]int32** |  | 
 **projectCategoryIds** | **[]int32** |  | 
 **filterTagIds** | **[]int32** |  | 
 **companyId2** | **[]int32** |  | 
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


## GETProjectsApiV2Defaultprojects

> GETProjectsApiV2Defaultprojects(ctx).CompanyId(companyId).Execute()

Returns the project defaults

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
    companyId := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2Defaultprojects(context.Background()).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2Defaultprojects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2DefaultprojectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** |  | 

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


## GETProjectsApiV2FilesobjectIdProject

> GETProjectsApiV2FilesobjectIdProject(ctx, objectId2).UserId(userId).ObjectId(objectId).ShowDeleted(showDeleted).IncludeTags(includeTags).IncludePeople(includePeople).GetTruePermissions(getTruePermissions).GetPermissions(getPermissions).GetNotificationSettings(getNotificationSettings).GetEmoji(getEmoji).GetEmailAddress(getEmailAddress).GetDateInfo(getDateInfo).GetActivePages(getActivePages).FormatMarkdown(formatMarkdown).Execute()

Get the project for the file

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
    objectId2 := int32(56) // int32 | 
    userId := int32(56) // int32 |  (optional)
    objectId := int32(56) // int32 |  (optional)
    showDeleted := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    includePeople := true // bool |  (optional)
    getTruePermissions := true // bool |  (optional)
    getPermissions := true // bool |  (optional)
    getNotificationSettings := true // bool |  (optional)
    getEmoji := true // bool |  (optional)
    getEmailAddress := true // bool |  (optional)
    getDateInfo := true // bool |  (optional)
    getActivePages := true // bool |  (optional)
    formatMarkdown := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2FilesobjectIdProject(context.Background(), objectId2).UserId(userId).ObjectId(objectId).ShowDeleted(showDeleted).IncludeTags(includeTags).IncludePeople(includePeople).GetTruePermissions(getTruePermissions).GetPermissions(getPermissions).GetNotificationSettings(getNotificationSettings).GetEmoji(getEmoji).GetEmailAddress(getEmailAddress).GetDateInfo(getDateInfo).GetActivePages(getActivePages).FormatMarkdown(formatMarkdown).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2FilesobjectIdProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2FilesobjectIdProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **int32** |  | 
 **objectId** | **int32** |  | 
 **showDeleted** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **includePeople** | **bool** |  | 
 **getTruePermissions** | **bool** |  | 
 **getPermissions** | **bool** |  | 
 **getNotificationSettings** | **bool** |  | 
 **getEmoji** | **bool** |  | 
 **getEmailAddress** | **bool** |  | 
 **getDateInfo** | **bool** |  | 
 **getActivePages** | **bool** |  | 
 **formatMarkdown** | **bool** |  | 

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


## GETProjectsApiV2Fullsearch

> GETProjectsApiV2Fullsearch(ctx).SortOrder(sortOrder).SearchTerm(searchTerm).SearchFor(searchFor).ProjectId(projectId).PageSize(pageSize).Page(page).DayLimit(dayLimit).UseSmartSearch(useSmartSearch).SkipES(skipES).SeparateUsers(separateUsers).SearchWithTags(searchWithTags).SearchForIsList(searchForIsList).SearchArchivedMessages(searchArchivedMessages).MatchAllTags(matchAllTags).IncludeTags(includeTags).IncludeCompletedItems(includeCompletedItems).IncludeArchivedProjects(includeArchivedProjects).GetUsers(getUsers).EventsInUTC(eventsInUTC).Tags(tags).TagIds(tagIds).Execute()

Perform a full search

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
    sortOrder := "sortOrder_example" // string |  (optional)
    searchTerm := "searchTerm_example" // string |  (optional)
    searchFor := "searchFor_example" // string |  (optional)
    projectId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    dayLimit := int32(56) // int32 |  (optional)
    useSmartSearch := true // bool |  (optional)
    skipES := true // bool |  (optional)
    separateUsers := true // bool |  (optional)
    searchWithTags := true // bool |  (optional)
    searchForIsList := true // bool |  (optional)
    searchArchivedMessages := true // bool |  (optional)
    matchAllTags := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    includeCompletedItems := true // bool |  (optional)
    includeArchivedProjects := true // bool |  (optional)
    getUsers := true // bool |  (optional)
    eventsInUTC := true // bool |  (optional)
    tags := []string{"Inner_example"} // []string |  (optional)
    tagIds := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2Fullsearch(context.Background()).SortOrder(sortOrder).SearchTerm(searchTerm).SearchFor(searchFor).ProjectId(projectId).PageSize(pageSize).Page(page).DayLimit(dayLimit).UseSmartSearch(useSmartSearch).SkipES(skipES).SeparateUsers(separateUsers).SearchWithTags(searchWithTags).SearchForIsList(searchForIsList).SearchArchivedMessages(searchArchivedMessages).MatchAllTags(matchAllTags).IncludeTags(includeTags).IncludeCompletedItems(includeCompletedItems).IncludeArchivedProjects(includeArchivedProjects).GetUsers(getUsers).EventsInUTC(eventsInUTC).Tags(tags).TagIds(tagIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2Fullsearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2FullsearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortOrder** | **string** |  | 
 **searchTerm** | **string** |  | 
 **searchFor** | **string** |  | 
 **projectId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **dayLimit** | **int32** |  | 
 **useSmartSearch** | **bool** |  | 
 **skipES** | **bool** |  | 
 **separateUsers** | **bool** |  | 
 **searchWithTags** | **bool** |  | 
 **searchForIsList** | **bool** |  | 
 **searchArchivedMessages** | **bool** |  | 
 **matchAllTags** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **includeCompletedItems** | **bool** |  | 
 **includeArchivedProjects** | **bool** |  | 
 **getUsers** | **bool** |  | 
 **eventsInUTC** | **bool** |  | 
 **tags** | **[]string** |  | 
 **tagIds** | **[]int32** |  | 

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


## GETProjectsApiV2LinksobjectIdProject

> GETProjectsApiV2LinksobjectIdProject(ctx, objectId2).UserId(userId).ObjectId(objectId).ShowDeleted(showDeleted).IncludeTags(includeTags).IncludePeople(includePeople).GetTruePermissions(getTruePermissions).GetPermissions(getPermissions).GetNotificationSettings(getNotificationSettings).GetEmoji(getEmoji).GetEmailAddress(getEmailAddress).GetDateInfo(getDateInfo).GetActivePages(getActivePages).FormatMarkdown(formatMarkdown).Execute()

Get the project for the link

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
    objectId2 := int32(56) // int32 | 
    userId := int32(56) // int32 |  (optional)
    objectId := int32(56) // int32 |  (optional)
    showDeleted := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    includePeople := true // bool |  (optional)
    getTruePermissions := true // bool |  (optional)
    getPermissions := true // bool |  (optional)
    getNotificationSettings := true // bool |  (optional)
    getEmoji := true // bool |  (optional)
    getEmailAddress := true // bool |  (optional)
    getDateInfo := true // bool |  (optional)
    getActivePages := true // bool |  (optional)
    formatMarkdown := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2LinksobjectIdProject(context.Background(), objectId2).UserId(userId).ObjectId(objectId).ShowDeleted(showDeleted).IncludeTags(includeTags).IncludePeople(includePeople).GetTruePermissions(getTruePermissions).GetPermissions(getPermissions).GetNotificationSettings(getNotificationSettings).GetEmoji(getEmoji).GetEmailAddress(getEmailAddress).GetDateInfo(getDateInfo).GetActivePages(getActivePages).FormatMarkdown(formatMarkdown).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2LinksobjectIdProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2LinksobjectIdProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **int32** |  | 
 **objectId** | **int32** |  | 
 **showDeleted** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **includePeople** | **bool** |  | 
 **getTruePermissions** | **bool** |  | 
 **getPermissions** | **bool** |  | 
 **getNotificationSettings** | **bool** |  | 
 **getEmoji** | **bool** |  | 
 **getEmailAddress** | **bool** |  | 
 **getDateInfo** | **bool** |  | 
 **getActivePages** | **bool** |  | 
 **formatMarkdown** | **bool** |  | 

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


## GETProjectsApiV2Me

> PeoplePersonResponse GETProjectsApiV2Me(ctx).SharedFilter(sharedFilter).GetDefaultFilters(getDefaultFilters).UserID(userID).ShowDeleted(showDeleted).IncludeTags(includeTags).IncludeClockIn(includeClockIn).IncludeAuth(includeAuth).GetTeamsStats(getTeamsStats).GetTasksStats(getTasksStats).GetProjectsStats(getProjectsStats).GetPreferences(getPreferences).GetPermissions(getPermissions).GetMilestonesStats(getMilestonesStats).GetInitialPage(getInitialPage).GetAllStats(getAllStats).GetAccounts(getAccounts).Fullprofile(fullprofile).CleanPreferences(cleanPreferences).Execute()

Returns the details for the account you’re currently logged in with.

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
    sharedFilter := "sharedFilter_example" // string |  (optional)
    getDefaultFilters := "getDefaultFilters_example" // string |  (optional)
    userID := int32(56) // int32 |  (optional)
    showDeleted := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    includeClockIn := true // bool |  (optional)
    includeAuth := true // bool |  (optional)
    getTeamsStats := true // bool |  (optional)
    getTasksStats := true // bool |  (optional)
    getProjectsStats := true // bool |  (optional)
    getPreferences := true // bool |  (optional)
    getPermissions := true // bool |  (optional)
    getMilestonesStats := true // bool |  (optional)
    getInitialPage := true // bool |  (optional)
    getAllStats := true // bool | statistics (optional)
    getAccounts := true // bool |  (optional)
    fullprofile := true // bool |  (optional)
    cleanPreferences := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2Me(context.Background()).SharedFilter(sharedFilter).GetDefaultFilters(getDefaultFilters).UserID(userID).ShowDeleted(showDeleted).IncludeTags(includeTags).IncludeClockIn(includeClockIn).IncludeAuth(includeAuth).GetTeamsStats(getTeamsStats).GetTasksStats(getTasksStats).GetProjectsStats(getProjectsStats).GetPreferences(getPreferences).GetPermissions(getPermissions).GetMilestonesStats(getMilestonesStats).GetInitialPage(getInitialPage).GetAllStats(getAllStats).GetAccounts(getAccounts).Fullprofile(fullprofile).CleanPreferences(cleanPreferences).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2Me``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV2Me`: PeoplePersonResponse
    fmt.Fprintf(os.Stdout, "Response from `CFPortApi.GETProjectsApiV2Me`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2MeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sharedFilter** | **string** |  | 
 **getDefaultFilters** | **string** |  | 
 **userID** | **int32** |  | 
 **showDeleted** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **includeClockIn** | **bool** |  | 
 **includeAuth** | **bool** |  | 
 **getTeamsStats** | **bool** |  | 
 **getTasksStats** | **bool** |  | 
 **getProjectsStats** | **bool** |  | 
 **getPreferences** | **bool** |  | 
 **getPermissions** | **bool** |  | 
 **getMilestonesStats** | **bool** |  | 
 **getInitialPage** | **bool** |  | 
 **getAllStats** | **bool** | statistics | 
 **getAccounts** | **bool** |  | 
 **fullprofile** | **bool** |  | 
 **cleanPreferences** | **bool** |  | 

### Return type

[**PeoplePersonResponse**](PeoplePersonResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV2MePersonaltasks

> GETProjectsApiV2MePersonaltasks(ctx).UpdatedAfterDate(updatedAfterDate).Today(today).StartDate(startDate).SortOrder(sortOrder).Sort(sort).ReportFormat(reportFormat).ProjectStatus(projectStatus).Priority(priority).Include(include).Filter(filter).EndDate(endDate).DeletedAfterDate(deletedAfterDate).DateCode(dateCode).DataSet(dataSet).CreatedFilter(createdFilter).CreatedDateCode(createdDateCode).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TaskId(taskId).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).OffsetCount(offsetCount).IncludeTaskId(includeTaskId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyTasksWithUnreadComments(onlyTasksWithUnreadComments).OnlyTasksWithTickets(onlyTasksWithTickets).OnlyTasksWithEstimatedTime(onlyTasksWithEstimatedTime).OnlyStarredProjects(onlyStarredProjects).NestSubTasks(nestSubTasks).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTagIds(matchAllProjectTagIds).MatchAllExcludedTags(matchAllExcludedTags).IsReportDownload(isReportDownload).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksWithCards(includeTasksWithCards).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCustomFields(includeCustomFields).IncludeCrmDealIds(includeCrmDealIds).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompanyUserIds(includeCompanyUserIds).IncludeBlockedTasks(includeBlockedTasks).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).ExcludeAssigneeNotOnProjectTeams(excludeAssigneeNotOnProjectTeams).AllowTemplateTasks(allowTemplateTasks).TasklistId(tasklistId).TaskIdList(taskIdList).Tags(tags).TagIds(tagIds).SkipCrmDealIds(skipCrmDealIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ProjectCategoryId(projectCategoryId).FollowerIds(followerIds).FilterText(filterText).FilterBoardColumnIds(filterBoardColumnIds).ExcludeTagIds(excludeTagIds).CrmDealIds(crmDealIds).CreatorIds(creatorIds).CompanyId(companyId).AssigneeTeamIds(assigneeTeamIds).AssigneeCompanyIds(assigneeCompanyIds).Execute()

Get Tasks

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
    startDate := "startDate_example" // string |  (optional)
    sortOrder := "sortOrder_example" // string |  (optional)
    sort := "sort_example" // string |  (optional)
    reportFormat := "reportFormat_example" // string |  (optional)
    projectStatus := "projectStatus_example" // string |  (optional)
    priority := "priority_example" // string |  (optional)
    include := "include_example" // string |  (optional)
    filter := "filter_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    deletedAfterDate := "deletedAfterDate_example" // string |  (optional)
    dateCode := "dateCode_example" // string |  (optional)
    dataSet := "dataSet_example" // string |  (optional)
    createdFilter := "createdFilter_example" // string |  (optional)
    createdDateCode := "createdDateCode_example" // string |  (optional)
    createdBeforeDate := "createdBeforeDate_example" // string |  (optional)
    createdAfterDate := "createdAfterDate_example" // string |  (optional)
    completedBeforeDate := "completedBeforeDate_example" // string |  (optional)
    completedAfterDate := "completedAfterDate_example" // string |  (optional)
    callback := "callback_example" // string |  (optional)
    taskId := int32(56) // int32 |  (optional)
    parentTaskId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    offsetCount := int32(56) // int32 |  (optional)
    includeTaskId := int32(56) // int32 |  (optional)
    useStartDatesForTodaysTasks := true // bool |  (optional)
    useAllProjects := true // bool |  (optional)
    starredProjectsOnly := true // bool |  (optional)
    showDeleted := true // bool |  (optional)
    showCompletedLists := true // bool |  (optional)
    onlyUntaggedTasks := true // bool |  (optional)
    onlyTasksWithUnreadComments := true // bool |  (optional)
    onlyTasksWithTickets := true // bool |  (optional)
    onlyTasksWithEstimatedTime := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    nestSubTasks := true // bool |  (optional)
    matchAllTags := true // bool |  (optional)
    matchAllProjectTags := true // bool |  (optional)
    matchAllProjectTagIds := true // bool |  (optional)
    matchAllExcludedTags := true // bool |  (optional)
    isReportDownload := true // bool |  (optional)
    includeUntaggedTasks := true // bool |  (optional)
    includeToday := true // bool |  (optional)
    includeTeamUserIds := true // bool |  (optional)
    includeTasksWithoutDueDates := true // bool |  (optional)
    includeTasksWithCards := true // bool |  (optional)
    includeTasksFromDeletedLists := true // bool |  (optional)
    includeReminders := true // bool |  (optional)
    includeLoggedTime := true // bool |  (optional)
    includeCustomFields := true // bool |  (optional)
    includeCrmDealIds := true // bool |  (optional)
    includeCompletedTasks := true // bool |  (optional)
    includeCompletedSubtasks := true // bool |  (optional)
    includeCompanyUserIds := true // bool |  (optional)
    includeBlockedTasks := true // bool |  (optional)
    includeAssigneeTeams := true // bool |  (optional)
    includeAssigneeCompanies := true // bool |  (optional)
    includeArchivedProjects := true // bool |  (optional)
    ignoreStartDates := true // bool |  (optional)
    getSubTasks := true // bool |  (optional)
    getFiles := true // bool |  (optional)
    excludeAssigneeNotOnProjectTeams := true // bool |  (optional)
    allowTemplateTasks := true // bool |  (optional)
    tasklistId := []int32{int32(123)} // []int32 |  (optional)
    taskIdList := []int32{int32(123)} // []int32 |  (optional)
    tags := []string{"Inner_example"} // []string |  (optional)
    tagIds := []int32{int32(123)} // []int32 |  (optional)
    skipCrmDealIds := []int32{int32(123)} // []int32 |  (optional)
    responsiblePartyIds := []int32{int32(123)} // []int32 |  (optional)
    projectTagIds := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    projectIds := []int32{int32(123)} // []int32 |  (optional)
    projectId := []int32{int32(123)} // []int32 | project filters (optional)
    projectHealths := []int32{int32(123)} // []int32 |  (optional)
    projectHealth := []int32{int32(123)} // []int32 |  (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryId := []int32{int32(123)} // []int32 |  (optional)
    followerIds := []int32{int32(123)} // []int32 |  (optional)
    filterText := []string{"Inner_example"} // []string |  (optional)
    filterBoardColumnIds := []int32{int32(123)} // []int32 |  (optional)
    excludeTagIds := []int32{int32(123)} // []int32 |  (optional)
    crmDealIds := []int32{int32(123)} // []int32 |  (optional)
    creatorIds := []int32{int32(123)} // []int32 |  (optional)
    companyId := []int32{int32(123)} // []int32 |  (optional)
    assigneeTeamIds := []int32{int32(123)} // []int32 |  (optional)
    assigneeCompanyIds := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2MePersonaltasks(context.Background()).UpdatedAfterDate(updatedAfterDate).Today(today).StartDate(startDate).SortOrder(sortOrder).Sort(sort).ReportFormat(reportFormat).ProjectStatus(projectStatus).Priority(priority).Include(include).Filter(filter).EndDate(endDate).DeletedAfterDate(deletedAfterDate).DateCode(dateCode).DataSet(dataSet).CreatedFilter(createdFilter).CreatedDateCode(createdDateCode).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TaskId(taskId).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).OffsetCount(offsetCount).IncludeTaskId(includeTaskId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyTasksWithUnreadComments(onlyTasksWithUnreadComments).OnlyTasksWithTickets(onlyTasksWithTickets).OnlyTasksWithEstimatedTime(onlyTasksWithEstimatedTime).OnlyStarredProjects(onlyStarredProjects).NestSubTasks(nestSubTasks).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTagIds(matchAllProjectTagIds).MatchAllExcludedTags(matchAllExcludedTags).IsReportDownload(isReportDownload).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksWithCards(includeTasksWithCards).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCustomFields(includeCustomFields).IncludeCrmDealIds(includeCrmDealIds).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompanyUserIds(includeCompanyUserIds).IncludeBlockedTasks(includeBlockedTasks).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).ExcludeAssigneeNotOnProjectTeams(excludeAssigneeNotOnProjectTeams).AllowTemplateTasks(allowTemplateTasks).TasklistId(tasklistId).TaskIdList(taskIdList).Tags(tags).TagIds(tagIds).SkipCrmDealIds(skipCrmDealIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ProjectCategoryId(projectCategoryId).FollowerIds(followerIds).FilterText(filterText).FilterBoardColumnIds(filterBoardColumnIds).ExcludeTagIds(excludeTagIds).CrmDealIds(crmDealIds).CreatorIds(creatorIds).CompanyId(companyId).AssigneeTeamIds(assigneeTeamIds).AssigneeCompanyIds(assigneeCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2MePersonaltasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2MePersonaltasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfterDate** | **string** |  | 
 **today** | **string** |  | 
 **startDate** | **string** |  | 
 **sortOrder** | **string** |  | 
 **sort** | **string** |  | 
 **reportFormat** | **string** |  | 
 **projectStatus** | **string** |  | 
 **priority** | **string** |  | 
 **include** | **string** |  | 
 **filter** | **string** |  | 
 **endDate** | **string** |  | 
 **deletedAfterDate** | **string** |  | 
 **dateCode** | **string** |  | 
 **dataSet** | **string** |  | 
 **createdFilter** | **string** |  | 
 **createdDateCode** | **string** |  | 
 **createdBeforeDate** | **string** |  | 
 **createdAfterDate** | **string** |  | 
 **completedBeforeDate** | **string** |  | 
 **completedAfterDate** | **string** |  | 
 **callback** | **string** |  | 
 **taskId** | **int32** |  | 
 **parentTaskId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **offsetCount** | **int32** |  | 
 **includeTaskId** | **int32** |  | 
 **useStartDatesForTodaysTasks** | **bool** |  | 
 **useAllProjects** | **bool** |  | 
 **starredProjectsOnly** | **bool** |  | 
 **showDeleted** | **bool** |  | 
 **showCompletedLists** | **bool** |  | 
 **onlyUntaggedTasks** | **bool** |  | 
 **onlyTasksWithUnreadComments** | **bool** |  | 
 **onlyTasksWithTickets** | **bool** |  | 
 **onlyTasksWithEstimatedTime** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **nestSubTasks** | **bool** |  | 
 **matchAllTags** | **bool** |  | 
 **matchAllProjectTags** | **bool** |  | 
 **matchAllProjectTagIds** | **bool** |  | 
 **matchAllExcludedTags** | **bool** |  | 
 **isReportDownload** | **bool** |  | 
 **includeUntaggedTasks** | **bool** |  | 
 **includeToday** | **bool** |  | 
 **includeTeamUserIds** | **bool** |  | 
 **includeTasksWithoutDueDates** | **bool** |  | 
 **includeTasksWithCards** | **bool** |  | 
 **includeTasksFromDeletedLists** | **bool** |  | 
 **includeReminders** | **bool** |  | 
 **includeLoggedTime** | **bool** |  | 
 **includeCustomFields** | **bool** |  | 
 **includeCrmDealIds** | **bool** |  | 
 **includeCompletedTasks** | **bool** |  | 
 **includeCompletedSubtasks** | **bool** |  | 
 **includeCompanyUserIds** | **bool** |  | 
 **includeBlockedTasks** | **bool** |  | 
 **includeAssigneeTeams** | **bool** |  | 
 **includeAssigneeCompanies** | **bool** |  | 
 **includeArchivedProjects** | **bool** |  | 
 **ignoreStartDates** | **bool** |  | 
 **getSubTasks** | **bool** |  | 
 **getFiles** | **bool** |  | 
 **excludeAssigneeNotOnProjectTeams** | **bool** |  | 
 **allowTemplateTasks** | **bool** |  | 
 **tasklistId** | **[]int32** |  | 
 **taskIdList** | **[]int32** |  | 
 **tags** | **[]string** |  | 
 **tagIds** | **[]int32** |  | 
 **skipCrmDealIds** | **[]int32** |  | 
 **responsiblePartyIds** | **[]int32** |  | 
 **projectTagIds** | **[]int32** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **projectIds** | **[]int32** |  | 
 **projectId** | **[]int32** | project filters | 
 **projectHealths** | **[]int32** |  | 
 **projectHealth** | **[]int32** |  | 
 **projectCompanyIds** | **[]int32** |  | 
 **projectCategoryIds** | **[]int32** |  | 
 **projectCategoryId** | **[]int32** |  | 
 **followerIds** | **[]int32** |  | 
 **filterText** | **[]string** |  | 
 **filterBoardColumnIds** | **[]int32** |  | 
 **excludeTagIds** | **[]int32** |  | 
 **crmDealIds** | **[]int32** |  | 
 **creatorIds** | **[]int32** |  | 
 **companyId** | **[]int32** |  | 
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


## GETProjectsApiV2MePersonaltaskstaskId

> GETProjectsApiV2MePersonaltaskstaskId(ctx, taskId2).UpdatedAfterDate(updatedAfterDate).Today(today).StartDate(startDate).SortOrder(sortOrder).Sort(sort).ReportFormat(reportFormat).ProjectStatus(projectStatus).Priority(priority).Include(include).Filter(filter).EndDate(endDate).DeletedAfterDate(deletedAfterDate).DateCode(dateCode).DataSet(dataSet).CreatedFilter(createdFilter).CreatedDateCode(createdDateCode).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TaskId(taskId).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).OffsetCount(offsetCount).IncludeTaskId(includeTaskId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyTasksWithUnreadComments(onlyTasksWithUnreadComments).OnlyTasksWithTickets(onlyTasksWithTickets).OnlyTasksWithEstimatedTime(onlyTasksWithEstimatedTime).OnlyStarredProjects(onlyStarredProjects).NestSubTasks(nestSubTasks).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTagIds(matchAllProjectTagIds).MatchAllExcludedTags(matchAllExcludedTags).IsReportDownload(isReportDownload).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksWithCards(includeTasksWithCards).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCustomFields(includeCustomFields).IncludeCrmDealIds(includeCrmDealIds).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompanyUserIds(includeCompanyUserIds).IncludeBlockedTasks(includeBlockedTasks).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).ExcludeAssigneeNotOnProjectTeams(excludeAssigneeNotOnProjectTeams).AllowTemplateTasks(allowTemplateTasks).TasklistId(tasklistId).TaskIdList(taskIdList).Tags(tags).TagIds(tagIds).SkipCrmDealIds(skipCrmDealIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ProjectCategoryId(projectCategoryId).FollowerIds(followerIds).FilterText(filterText).FilterBoardColumnIds(filterBoardColumnIds).ExcludeTagIds(excludeTagIds).CrmDealIds(crmDealIds).CreatorIds(creatorIds).CompanyId(companyId).AssigneeTeamIds(assigneeTeamIds).AssigneeCompanyIds(assigneeCompanyIds).Execute()

Get Tasks

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
    updatedAfterDate := "updatedAfterDate_example" // string |  (optional)
    today := "today_example" // string |  (optional)
    startDate := "startDate_example" // string |  (optional)
    sortOrder := "sortOrder_example" // string |  (optional)
    sort := "sort_example" // string |  (optional)
    reportFormat := "reportFormat_example" // string |  (optional)
    projectStatus := "projectStatus_example" // string |  (optional)
    priority := "priority_example" // string |  (optional)
    include := "include_example" // string |  (optional)
    filter := "filter_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    deletedAfterDate := "deletedAfterDate_example" // string |  (optional)
    dateCode := "dateCode_example" // string |  (optional)
    dataSet := "dataSet_example" // string |  (optional)
    createdFilter := "createdFilter_example" // string |  (optional)
    createdDateCode := "createdDateCode_example" // string |  (optional)
    createdBeforeDate := "createdBeforeDate_example" // string |  (optional)
    createdAfterDate := "createdAfterDate_example" // string |  (optional)
    completedBeforeDate := "completedBeforeDate_example" // string |  (optional)
    completedAfterDate := "completedAfterDate_example" // string |  (optional)
    callback := "callback_example" // string |  (optional)
    taskId := int32(56) // int32 |  (optional)
    parentTaskId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    offsetCount := int32(56) // int32 |  (optional)
    includeTaskId := int32(56) // int32 |  (optional)
    useStartDatesForTodaysTasks := true // bool |  (optional)
    useAllProjects := true // bool |  (optional)
    starredProjectsOnly := true // bool |  (optional)
    showDeleted := true // bool |  (optional)
    showCompletedLists := true // bool |  (optional)
    onlyUntaggedTasks := true // bool |  (optional)
    onlyTasksWithUnreadComments := true // bool |  (optional)
    onlyTasksWithTickets := true // bool |  (optional)
    onlyTasksWithEstimatedTime := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    nestSubTasks := true // bool |  (optional)
    matchAllTags := true // bool |  (optional)
    matchAllProjectTags := true // bool |  (optional)
    matchAllProjectTagIds := true // bool |  (optional)
    matchAllExcludedTags := true // bool |  (optional)
    isReportDownload := true // bool |  (optional)
    includeUntaggedTasks := true // bool |  (optional)
    includeToday := true // bool |  (optional)
    includeTeamUserIds := true // bool |  (optional)
    includeTasksWithoutDueDates := true // bool |  (optional)
    includeTasksWithCards := true // bool |  (optional)
    includeTasksFromDeletedLists := true // bool |  (optional)
    includeReminders := true // bool |  (optional)
    includeLoggedTime := true // bool |  (optional)
    includeCustomFields := true // bool |  (optional)
    includeCrmDealIds := true // bool |  (optional)
    includeCompletedTasks := true // bool |  (optional)
    includeCompletedSubtasks := true // bool |  (optional)
    includeCompanyUserIds := true // bool |  (optional)
    includeBlockedTasks := true // bool |  (optional)
    includeAssigneeTeams := true // bool |  (optional)
    includeAssigneeCompanies := true // bool |  (optional)
    includeArchivedProjects := true // bool |  (optional)
    ignoreStartDates := true // bool |  (optional)
    getSubTasks := true // bool |  (optional)
    getFiles := true // bool |  (optional)
    excludeAssigneeNotOnProjectTeams := true // bool |  (optional)
    allowTemplateTasks := true // bool |  (optional)
    tasklistId := []int32{int32(123)} // []int32 |  (optional)
    taskIdList := []int32{int32(123)} // []int32 |  (optional)
    tags := []string{"Inner_example"} // []string |  (optional)
    tagIds := []int32{int32(123)} // []int32 |  (optional)
    skipCrmDealIds := []int32{int32(123)} // []int32 |  (optional)
    responsiblePartyIds := []int32{int32(123)} // []int32 |  (optional)
    projectTagIds := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    projectIds := []int32{int32(123)} // []int32 |  (optional)
    projectId := []int32{int32(123)} // []int32 | project filters (optional)
    projectHealths := []int32{int32(123)} // []int32 |  (optional)
    projectHealth := []int32{int32(123)} // []int32 |  (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryId := []int32{int32(123)} // []int32 |  (optional)
    followerIds := []int32{int32(123)} // []int32 |  (optional)
    filterText := []string{"Inner_example"} // []string |  (optional)
    filterBoardColumnIds := []int32{int32(123)} // []int32 |  (optional)
    excludeTagIds := []int32{int32(123)} // []int32 |  (optional)
    crmDealIds := []int32{int32(123)} // []int32 |  (optional)
    creatorIds := []int32{int32(123)} // []int32 |  (optional)
    companyId := []int32{int32(123)} // []int32 |  (optional)
    assigneeTeamIds := []int32{int32(123)} // []int32 |  (optional)
    assigneeCompanyIds := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2MePersonaltaskstaskId(context.Background(), taskId2).UpdatedAfterDate(updatedAfterDate).Today(today).StartDate(startDate).SortOrder(sortOrder).Sort(sort).ReportFormat(reportFormat).ProjectStatus(projectStatus).Priority(priority).Include(include).Filter(filter).EndDate(endDate).DeletedAfterDate(deletedAfterDate).DateCode(dateCode).DataSet(dataSet).CreatedFilter(createdFilter).CreatedDateCode(createdDateCode).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TaskId(taskId).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).OffsetCount(offsetCount).IncludeTaskId(includeTaskId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyTasksWithUnreadComments(onlyTasksWithUnreadComments).OnlyTasksWithTickets(onlyTasksWithTickets).OnlyTasksWithEstimatedTime(onlyTasksWithEstimatedTime).OnlyStarredProjects(onlyStarredProjects).NestSubTasks(nestSubTasks).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTagIds(matchAllProjectTagIds).MatchAllExcludedTags(matchAllExcludedTags).IsReportDownload(isReportDownload).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksWithCards(includeTasksWithCards).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCustomFields(includeCustomFields).IncludeCrmDealIds(includeCrmDealIds).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompanyUserIds(includeCompanyUserIds).IncludeBlockedTasks(includeBlockedTasks).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).ExcludeAssigneeNotOnProjectTeams(excludeAssigneeNotOnProjectTeams).AllowTemplateTasks(allowTemplateTasks).TasklistId(tasklistId).TaskIdList(taskIdList).Tags(tags).TagIds(tagIds).SkipCrmDealIds(skipCrmDealIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ProjectCategoryId(projectCategoryId).FollowerIds(followerIds).FilterText(filterText).FilterBoardColumnIds(filterBoardColumnIds).ExcludeTagIds(excludeTagIds).CrmDealIds(crmDealIds).CreatorIds(creatorIds).CompanyId(companyId).AssigneeTeamIds(assigneeTeamIds).AssigneeCompanyIds(assigneeCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2MePersonaltaskstaskId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2MePersonaltaskstaskIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedAfterDate** | **string** |  | 
 **today** | **string** |  | 
 **startDate** | **string** |  | 
 **sortOrder** | **string** |  | 
 **sort** | **string** |  | 
 **reportFormat** | **string** |  | 
 **projectStatus** | **string** |  | 
 **priority** | **string** |  | 
 **include** | **string** |  | 
 **filter** | **string** |  | 
 **endDate** | **string** |  | 
 **deletedAfterDate** | **string** |  | 
 **dateCode** | **string** |  | 
 **dataSet** | **string** |  | 
 **createdFilter** | **string** |  | 
 **createdDateCode** | **string** |  | 
 **createdBeforeDate** | **string** |  | 
 **createdAfterDate** | **string** |  | 
 **completedBeforeDate** | **string** |  | 
 **completedAfterDate** | **string** |  | 
 **callback** | **string** |  | 
 **taskId** | **int32** |  | 
 **parentTaskId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **offsetCount** | **int32** |  | 
 **includeTaskId** | **int32** |  | 
 **useStartDatesForTodaysTasks** | **bool** |  | 
 **useAllProjects** | **bool** |  | 
 **starredProjectsOnly** | **bool** |  | 
 **showDeleted** | **bool** |  | 
 **showCompletedLists** | **bool** |  | 
 **onlyUntaggedTasks** | **bool** |  | 
 **onlyTasksWithUnreadComments** | **bool** |  | 
 **onlyTasksWithTickets** | **bool** |  | 
 **onlyTasksWithEstimatedTime** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **nestSubTasks** | **bool** |  | 
 **matchAllTags** | **bool** |  | 
 **matchAllProjectTags** | **bool** |  | 
 **matchAllProjectTagIds** | **bool** |  | 
 **matchAllExcludedTags** | **bool** |  | 
 **isReportDownload** | **bool** |  | 
 **includeUntaggedTasks** | **bool** |  | 
 **includeToday** | **bool** |  | 
 **includeTeamUserIds** | **bool** |  | 
 **includeTasksWithoutDueDates** | **bool** |  | 
 **includeTasksWithCards** | **bool** |  | 
 **includeTasksFromDeletedLists** | **bool** |  | 
 **includeReminders** | **bool** |  | 
 **includeLoggedTime** | **bool** |  | 
 **includeCustomFields** | **bool** |  | 
 **includeCrmDealIds** | **bool** |  | 
 **includeCompletedTasks** | **bool** |  | 
 **includeCompletedSubtasks** | **bool** |  | 
 **includeCompanyUserIds** | **bool** |  | 
 **includeBlockedTasks** | **bool** |  | 
 **includeAssigneeTeams** | **bool** |  | 
 **includeAssigneeCompanies** | **bool** |  | 
 **includeArchivedProjects** | **bool** |  | 
 **ignoreStartDates** | **bool** |  | 
 **getSubTasks** | **bool** |  | 
 **getFiles** | **bool** |  | 
 **excludeAssigneeNotOnProjectTeams** | **bool** |  | 
 **allowTemplateTasks** | **bool** |  | 
 **tasklistId** | **[]int32** |  | 
 **taskIdList** | **[]int32** |  | 
 **tags** | **[]string** |  | 
 **tagIds** | **[]int32** |  | 
 **skipCrmDealIds** | **[]int32** |  | 
 **responsiblePartyIds** | **[]int32** |  | 
 **projectTagIds** | **[]int32** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **projectIds** | **[]int32** |  | 
 **projectId** | **[]int32** | project filters | 
 **projectHealths** | **[]int32** |  | 
 **projectHealth** | **[]int32** |  | 
 **projectCompanyIds** | **[]int32** |  | 
 **projectCategoryIds** | **[]int32** |  | 
 **projectCategoryId** | **[]int32** |  | 
 **followerIds** | **[]int32** |  | 
 **filterText** | **[]string** |  | 
 **filterBoardColumnIds** | **[]int32** |  | 
 **excludeTagIds** | **[]int32** |  | 
 **crmDealIds** | **[]int32** |  | 
 **creatorIds** | **[]int32** |  | 
 **companyId** | **[]int32** |  | 
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


## GETProjectsApiV2MessagesobjectIdProject

> GETProjectsApiV2MessagesobjectIdProject(ctx, objectId2).UserId(userId).ObjectId(objectId).ShowDeleted(showDeleted).IncludeTags(includeTags).IncludePeople(includePeople).GetTruePermissions(getTruePermissions).GetPermissions(getPermissions).GetNotificationSettings(getNotificationSettings).GetEmoji(getEmoji).GetEmailAddress(getEmailAddress).GetDateInfo(getDateInfo).GetActivePages(getActivePages).FormatMarkdown(formatMarkdown).Execute()

Get the project for the message

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
    objectId2 := int32(56) // int32 | 
    userId := int32(56) // int32 |  (optional)
    objectId := int32(56) // int32 |  (optional)
    showDeleted := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    includePeople := true // bool |  (optional)
    getTruePermissions := true // bool |  (optional)
    getPermissions := true // bool |  (optional)
    getNotificationSettings := true // bool |  (optional)
    getEmoji := true // bool |  (optional)
    getEmailAddress := true // bool |  (optional)
    getDateInfo := true // bool |  (optional)
    getActivePages := true // bool |  (optional)
    formatMarkdown := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2MessagesobjectIdProject(context.Background(), objectId2).UserId(userId).ObjectId(objectId).ShowDeleted(showDeleted).IncludeTags(includeTags).IncludePeople(includePeople).GetTruePermissions(getTruePermissions).GetPermissions(getPermissions).GetNotificationSettings(getNotificationSettings).GetEmoji(getEmoji).GetEmailAddress(getEmailAddress).GetDateInfo(getDateInfo).GetActivePages(getActivePages).FormatMarkdown(formatMarkdown).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2MessagesobjectIdProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2MessagesobjectIdProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **int32** |  | 
 **objectId** | **int32** |  | 
 **showDeleted** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **includePeople** | **bool** |  | 
 **getTruePermissions** | **bool** |  | 
 **getPermissions** | **bool** |  | 
 **getNotificationSettings** | **bool** |  | 
 **getEmoji** | **bool** |  | 
 **getEmailAddress** | **bool** |  | 
 **getDateInfo** | **bool** |  | 
 **getActivePages** | **bool** |  | 
 **formatMarkdown** | **bool** |  | 

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


## GETProjectsApiV2MilestonesobjectIdProject

> GETProjectsApiV2MilestonesobjectIdProject(ctx, objectId2).UserId(userId).ObjectId(objectId).ShowDeleted(showDeleted).IncludeTags(includeTags).IncludePeople(includePeople).GetTruePermissions(getTruePermissions).GetPermissions(getPermissions).GetNotificationSettings(getNotificationSettings).GetEmoji(getEmoji).GetEmailAddress(getEmailAddress).GetDateInfo(getDateInfo).GetActivePages(getActivePages).FormatMarkdown(formatMarkdown).Execute()

Get the project for the milestone

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
    objectId2 := int32(56) // int32 | 
    userId := int32(56) // int32 |  (optional)
    objectId := int32(56) // int32 |  (optional)
    showDeleted := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    includePeople := true // bool |  (optional)
    getTruePermissions := true // bool |  (optional)
    getPermissions := true // bool |  (optional)
    getNotificationSettings := true // bool |  (optional)
    getEmoji := true // bool |  (optional)
    getEmailAddress := true // bool |  (optional)
    getDateInfo := true // bool |  (optional)
    getActivePages := true // bool |  (optional)
    formatMarkdown := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2MilestonesobjectIdProject(context.Background(), objectId2).UserId(userId).ObjectId(objectId).ShowDeleted(showDeleted).IncludeTags(includeTags).IncludePeople(includePeople).GetTruePermissions(getTruePermissions).GetPermissions(getPermissions).GetNotificationSettings(getNotificationSettings).GetEmoji(getEmoji).GetEmailAddress(getEmailAddress).GetDateInfo(getDateInfo).GetActivePages(getActivePages).FormatMarkdown(formatMarkdown).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2MilestonesobjectIdProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2MilestonesobjectIdProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **int32** |  | 
 **objectId** | **int32** |  | 
 **showDeleted** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **includePeople** | **bool** |  | 
 **getTruePermissions** | **bool** |  | 
 **getPermissions** | **bool** |  | 
 **getNotificationSettings** | **bool** |  | 
 **getEmoji** | **bool** |  | 
 **getEmailAddress** | **bool** |  | 
 **getDateInfo** | **bool** |  | 
 **getActivePages** | **bool** |  | 
 **formatMarkdown** | **bool** |  | 

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


## GETProjectsApiV2NotebooksobjectIdProject

> GETProjectsApiV2NotebooksobjectIdProject(ctx, objectId2).UserId(userId).ObjectId(objectId).ShowDeleted(showDeleted).IncludeTags(includeTags).IncludePeople(includePeople).GetTruePermissions(getTruePermissions).GetPermissions(getPermissions).GetNotificationSettings(getNotificationSettings).GetEmoji(getEmoji).GetEmailAddress(getEmailAddress).GetDateInfo(getDateInfo).GetActivePages(getActivePages).FormatMarkdown(formatMarkdown).Execute()

Get the project for the notebook

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
    objectId2 := int32(56) // int32 | 
    userId := int32(56) // int32 |  (optional)
    objectId := int32(56) // int32 |  (optional)
    showDeleted := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    includePeople := true // bool |  (optional)
    getTruePermissions := true // bool |  (optional)
    getPermissions := true // bool |  (optional)
    getNotificationSettings := true // bool |  (optional)
    getEmoji := true // bool |  (optional)
    getEmailAddress := true // bool |  (optional)
    getDateInfo := true // bool |  (optional)
    getActivePages := true // bool |  (optional)
    formatMarkdown := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2NotebooksobjectIdProject(context.Background(), objectId2).UserId(userId).ObjectId(objectId).ShowDeleted(showDeleted).IncludeTags(includeTags).IncludePeople(includePeople).GetTruePermissions(getTruePermissions).GetPermissions(getPermissions).GetNotificationSettings(getNotificationSettings).GetEmoji(getEmoji).GetEmailAddress(getEmailAddress).GetDateInfo(getDateInfo).GetActivePages(getActivePages).FormatMarkdown(formatMarkdown).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2NotebooksobjectIdProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2NotebooksobjectIdProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **int32** |  | 
 **objectId** | **int32** |  | 
 **showDeleted** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **includePeople** | **bool** |  | 
 **getTruePermissions** | **bool** |  | 
 **getPermissions** | **bool** |  | 
 **getNotificationSettings** | **bool** |  | 
 **getEmoji** | **bool** |  | 
 **getEmailAddress** | **bool** |  | 
 **getDateInfo** | **bool** |  | 
 **getActivePages** | **bool** |  | 
 **formatMarkdown** | **bool** |  | 

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


## GETProjectsApiV2People

> PeopleResponse GETProjectsApiV2People(ctx).UpdatedAfterDateTime(updatedAfterDateTime).UpdatedAfterDate(updatedAfterDate).Type_(type_).SortOrder(sortOrder).Sort(sort).SearchTerm(searchTerm).ProjectType(projectType).InviteStatus(inviteStatus).FirstLetter(firstLetter).EmailAddress(emailAddress).DataSet(dataSet).ProjectId(projectId).ProjectCompanyId(projectCompanyId).PageSize(pageSize).Page(page).ExcludeProjectId(excludeProjectId).ShowDeleted(showDeleted).ReturnTeams(returnTeams).ReturnProjectIds(returnProjectIds).ReturnLetters(returnLetters).OwnerCompanyFirst(ownerCompanyFirst).Onlyids(onlyids).IncludeTags(includeTags).IncludeObservers(includeObservers).IncludeContacts(includeContacts).IncludeCompanyDetails(includeCompanyDetails).IncludeClockIn(includeClockIn).IncludeClients(includeClients).GroupByCompany(groupByCompany).GetProjectRoles(getProjectRoles).GetCounts(getCounts).Fullprofile(fullprofile).CountOnly(countOnly).UserIds(userIds).CompanyId(companyId).Execute()

All people visible to the user will be returned, including the user themselves. By default 100 records are returned at a time.

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
    type_ := "type__example" // string |  (optional)
    sortOrder := "sortOrder_example" // string |  (optional)
    sort := "sort_example" // string |  (optional)
    searchTerm := "searchTerm_example" // string |  (optional)
    projectType := "projectType_example" // string |  (optional)
    inviteStatus := "inviteStatus_example" // string |  (optional)
    firstLetter := "firstLetter_example" // string |  (optional)
    emailAddress := "emailAddress_example" // string |  (optional)
    dataSet := "dataSet_example" // string |  (optional)
    projectId := int32(56) // int32 |  (optional)
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
    userIds := []int32{int32(123)} // []int32 |  (optional)
    companyId := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2People(context.Background()).UpdatedAfterDateTime(updatedAfterDateTime).UpdatedAfterDate(updatedAfterDate).Type_(type_).SortOrder(sortOrder).Sort(sort).SearchTerm(searchTerm).ProjectType(projectType).InviteStatus(inviteStatus).FirstLetter(firstLetter).EmailAddress(emailAddress).DataSet(dataSet).ProjectId(projectId).ProjectCompanyId(projectCompanyId).PageSize(pageSize).Page(page).ExcludeProjectId(excludeProjectId).ShowDeleted(showDeleted).ReturnTeams(returnTeams).ReturnProjectIds(returnProjectIds).ReturnLetters(returnLetters).OwnerCompanyFirst(ownerCompanyFirst).Onlyids(onlyids).IncludeTags(includeTags).IncludeObservers(includeObservers).IncludeContacts(includeContacts).IncludeCompanyDetails(includeCompanyDetails).IncludeClockIn(includeClockIn).IncludeClients(includeClients).GroupByCompany(groupByCompany).GetProjectRoles(getProjectRoles).GetCounts(getCounts).Fullprofile(fullprofile).CountOnly(countOnly).UserIds(userIds).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2People``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV2People`: PeopleResponse
    fmt.Fprintf(os.Stdout, "Response from `CFPortApi.GETProjectsApiV2People`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2PeopleRequest struct via the builder pattern


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
 **projectId** | **int32** |  | 
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
 **userIds** | **[]int32** |  | 
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


## GETProjectsApiV2PeopleuserId

> PeoplePersonResponse GETProjectsApiV2PeopleuserId(ctx, userId).SharedFilter(sharedFilter).GetDefaultFilters(getDefaultFilters).UserID(userID).ShowDeleted(showDeleted).IncludeTags(includeTags).IncludeClockIn(includeClockIn).IncludeAuth(includeAuth).GetTeamsStats(getTeamsStats).GetTasksStats(getTasksStats).GetProjectsStats(getProjectsStats).GetPreferences(getPreferences).GetPermissions(getPermissions).GetMilestonesStats(getMilestonesStats).GetInitialPage(getInitialPage).GetAllStats(getAllStats).GetAccounts(getAccounts).Fullprofile(fullprofile).CleanPreferences(cleanPreferences).Execute()

Retrieves the user details for the ID submitted.

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
    sharedFilter := "sharedFilter_example" // string |  (optional)
    getDefaultFilters := "getDefaultFilters_example" // string |  (optional)
    userID := int32(56) // int32 |  (optional)
    showDeleted := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    includeClockIn := true // bool |  (optional)
    includeAuth := true // bool |  (optional)
    getTeamsStats := true // bool |  (optional)
    getTasksStats := true // bool |  (optional)
    getProjectsStats := true // bool |  (optional)
    getPreferences := true // bool |  (optional)
    getPermissions := true // bool |  (optional)
    getMilestonesStats := true // bool |  (optional)
    getInitialPage := true // bool |  (optional)
    getAllStats := true // bool | statistics (optional)
    getAccounts := true // bool |  (optional)
    fullprofile := true // bool |  (optional)
    cleanPreferences := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2PeopleuserId(context.Background(), userId).SharedFilter(sharedFilter).GetDefaultFilters(getDefaultFilters).UserID(userID).ShowDeleted(showDeleted).IncludeTags(includeTags).IncludeClockIn(includeClockIn).IncludeAuth(includeAuth).GetTeamsStats(getTeamsStats).GetTasksStats(getTasksStats).GetProjectsStats(getProjectsStats).GetPreferences(getPreferences).GetPermissions(getPermissions).GetMilestonesStats(getMilestonesStats).GetInitialPage(getInitialPage).GetAllStats(getAllStats).GetAccounts(getAccounts).Fullprofile(fullprofile).CleanPreferences(cleanPreferences).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2PeopleuserId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV2PeopleuserId`: PeoplePersonResponse
    fmt.Fprintf(os.Stdout, "Response from `CFPortApi.GETProjectsApiV2PeopleuserId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2PeopleuserIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sharedFilter** | **string** |  | 
 **getDefaultFilters** | **string** |  | 
 **userID** | **int32** |  | 
 **showDeleted** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **includeClockIn** | **bool** |  | 
 **includeAuth** | **bool** |  | 
 **getTeamsStats** | **bool** |  | 
 **getTasksStats** | **bool** |  | 
 **getProjectsStats** | **bool** |  | 
 **getPreferences** | **bool** |  | 
 **getPermissions** | **bool** |  | 
 **getMilestonesStats** | **bool** |  | 
 **getInitialPage** | **bool** |  | 
 **getAllStats** | **bool** | statistics | 
 **getAccounts** | **bool** |  | 
 **fullprofile** | **bool** |  | 
 **cleanPreferences** | **bool** |  | 

### Return type

[**PeoplePersonResponse**](PeoplePersonResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV2PeopleuserIdProjects

> GETProjectsApiV2PeopleuserIdProjects(ctx, userId2).WorkedOnAfterDateTime(workedOnAfterDateTime).UpdatedAfterDateTime(updatedAfterDateTime).UpdatedAfterDate(updatedAfterDate).Status(status).SearchTerm(searchTerm).ProjectType(projectType).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).Get(get).FirstLetter(firstLetter).DataSet(dataSet).CreatedAfterDateTime(createdAfterDateTime).CreatedAfterDate(createdAfterDate).CompletedBeforeDateTime(completedBeforeDateTime).CompletedAfterDateTime(completedAfterDateTime).UserId(userId).PageSize(pageSize).Page(page).SearchCompany(searchCompany).SearchCategory(searchCategory).SanitizeName(sanitizeName).ReturnLetters(returnLetters).OnlyStarredProjects(onlyStarredProjects).OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership).OnlyCustomFieldValues(onlyCustomFieldValues).OnlyArchivedProjects(onlyArchivedProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).IncludeUpdates(includeUpdates).IncludeTopPeople(includeTopPeople).IncludeTags(includeTags).IncludeProjectOwner(includeProjectOwner).IncludePeople(includePeople).IncludeCustomFields(includeCustomFields).IncludeCrmDealIds(includeCrmDealIds).IncludeArchivedProjects(includeArchivedProjects).HideObservedProjects(hideObservedProjects).HideDesc(hideDesc).GetTruePermissions(getTruePermissions).GetPermissions(getPermissions).GetNotificationSettings(getNotificationSettings).GetDeleted(getDeleted).GetCounts(getCounts).GetCategoryPath(getCategoryPath).GetAll(getAll).GetActivePages(getActivePages).FormatMarkdown(formatMarkdown).SkipIds(skipIds).SkipCrmDealIds(skipCrmDealIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).PortfolioColumnIds(portfolioColumnIds).PortfolioBoardIds(portfolioBoardIds).FilterTagIds(filterTagIds).FeaturesEnabledOnProject(featuresEnabledOnProject).CrmDealIds(crmDealIds).CompanyId(companyId).CategoryId(categoryId).CatId(catId).Execute()

APICall_GET_projects

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
    userId2 := int32(56) // int32 | 
    workedOnAfterDateTime := "workedOnAfterDateTime_example" // string |  (optional)
    updatedAfterDateTime := "updatedAfterDateTime_example" // string |  (optional)
    updatedAfterDate := "updatedAfterDate_example" // string |  (optional)
    status := "status_example" // string | common project filters (optional)
    searchTerm := "searchTerm_example" // string |  (optional)
    projectType := "projectType_example" // string |  (optional)
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
    searchCategory := true // bool |  (optional)
    sanitizeName := true // bool |  (optional)
    returnLetters := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    onlyProjectsWithExplicitMembership := true // bool |  (optional)
    onlyCustomFieldValues := true // bool |  (optional)
    onlyArchivedProjects := true // bool |  (optional)
    matchAllTags := true // bool |  (optional)
    matchAllProjectTags := true // bool |  (optional)
    includeUpdates := true // bool |  (optional)
    includeTopPeople := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    includeProjectOwner := true // bool |  (optional)
    includePeople := true // bool |  (optional)
    includeCustomFields := true // bool |  (optional)
    includeCrmDealIds := true // bool |  (optional)
    includeArchivedProjects := true // bool |  (optional)
    hideObservedProjects := true // bool |  (optional)
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
    skipCrmDealIds := []int32{int32(123)} // []int32 |  (optional)
    projectTagIds := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    projectHealths := []int32{int32(123)} // []int32 |  (optional)
    projectHealth := []int32{int32(123)} // []int32 |  (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 |  (optional)
    portfolioColumnIds := []int32{int32(123)} // []int32 |  (optional)
    portfolioBoardIds := []int32{int32(123)} // []int32 |  (optional)
    filterTagIds := []int32{int32(123)} // []int32 |  (optional)
    featuresEnabledOnProject := []string{"Inner_example"} // []string |  (optional)
    crmDealIds := []int32{int32(123)} // []int32 |  (optional)
    companyId := []int32{int32(123)} // []int32 |  (optional)
    categoryId := []int32{int32(123)} // []int32 |  (optional)
    catId := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2PeopleuserIdProjects(context.Background(), userId2).WorkedOnAfterDateTime(workedOnAfterDateTime).UpdatedAfterDateTime(updatedAfterDateTime).UpdatedAfterDate(updatedAfterDate).Status(status).SearchTerm(searchTerm).ProjectType(projectType).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).Get(get).FirstLetter(firstLetter).DataSet(dataSet).CreatedAfterDateTime(createdAfterDateTime).CreatedAfterDate(createdAfterDate).CompletedBeforeDateTime(completedBeforeDateTime).CompletedAfterDateTime(completedAfterDateTime).UserId(userId).PageSize(pageSize).Page(page).SearchCompany(searchCompany).SearchCategory(searchCategory).SanitizeName(sanitizeName).ReturnLetters(returnLetters).OnlyStarredProjects(onlyStarredProjects).OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership).OnlyCustomFieldValues(onlyCustomFieldValues).OnlyArchivedProjects(onlyArchivedProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).IncludeUpdates(includeUpdates).IncludeTopPeople(includeTopPeople).IncludeTags(includeTags).IncludeProjectOwner(includeProjectOwner).IncludePeople(includePeople).IncludeCustomFields(includeCustomFields).IncludeCrmDealIds(includeCrmDealIds).IncludeArchivedProjects(includeArchivedProjects).HideObservedProjects(hideObservedProjects).HideDesc(hideDesc).GetTruePermissions(getTruePermissions).GetPermissions(getPermissions).GetNotificationSettings(getNotificationSettings).GetDeleted(getDeleted).GetCounts(getCounts).GetCategoryPath(getCategoryPath).GetAll(getAll).GetActivePages(getActivePages).FormatMarkdown(formatMarkdown).SkipIds(skipIds).SkipCrmDealIds(skipCrmDealIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).PortfolioColumnIds(portfolioColumnIds).PortfolioBoardIds(portfolioBoardIds).FilterTagIds(filterTagIds).FeaturesEnabledOnProject(featuresEnabledOnProject).CrmDealIds(crmDealIds).CompanyId(companyId).CategoryId(categoryId).CatId(catId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2PeopleuserIdProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2PeopleuserIdProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workedOnAfterDateTime** | **string** |  | 
 **updatedAfterDateTime** | **string** |  | 
 **updatedAfterDate** | **string** |  | 
 **status** | **string** | common project filters | 
 **searchTerm** | **string** |  | 
 **projectType** | **string** |  | 
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
 **searchCategory** | **bool** |  | 
 **sanitizeName** | **bool** |  | 
 **returnLetters** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **onlyProjectsWithExplicitMembership** | **bool** |  | 
 **onlyCustomFieldValues** | **bool** |  | 
 **onlyArchivedProjects** | **bool** |  | 
 **matchAllTags** | **bool** |  | 
 **matchAllProjectTags** | **bool** |  | 
 **includeUpdates** | **bool** |  | 
 **includeTopPeople** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **includeProjectOwner** | **bool** |  | 
 **includePeople** | **bool** |  | 
 **includeCustomFields** | **bool** |  | 
 **includeCrmDealIds** | **bool** |  | 
 **includeArchivedProjects** | **bool** |  | 
 **hideObservedProjects** | **bool** |  | 
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
 **skipCrmDealIds** | **[]int32** |  | 
 **projectTagIds** | **[]int32** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **projectHealths** | **[]int32** |  | 
 **projectHealth** | **[]int32** |  | 
 **projectCompanyIds** | **[]int32** |  | 
 **projectCategoryIds** | **[]int32** |  | 
 **portfolioColumnIds** | **[]int32** |  | 
 **portfolioBoardIds** | **[]int32** |  | 
 **filterTagIds** | **[]int32** |  | 
 **featuresEnabledOnProject** | **[]string** |  | 
 **crmDealIds** | **[]int32** |  | 
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


## GETProjectsApiV2ProjectcategoriesprojectCategoryIdTasks

> GETProjectsApiV2ProjectcategoriesprojectCategoryIdTasks(ctx, projectCategoryId).UpdatedAfterDate(updatedAfterDate).Today(today).StartDate(startDate).SortOrder(sortOrder).Sort(sort).ReportFormat(reportFormat).ProjectStatus(projectStatus).Priority(priority).Include(include).Filter(filter).EndDate(endDate).DeletedAfterDate(deletedAfterDate).DateCode(dateCode).DataSet(dataSet).CreatedFilter(createdFilter).CreatedDateCode(createdDateCode).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TaskId(taskId).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).OffsetCount(offsetCount).IncludeTaskId(includeTaskId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyTasksWithUnreadComments(onlyTasksWithUnreadComments).OnlyTasksWithTickets(onlyTasksWithTickets).OnlyTasksWithEstimatedTime(onlyTasksWithEstimatedTime).OnlyStarredProjects(onlyStarredProjects).NestSubTasks(nestSubTasks).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTagIds(matchAllProjectTagIds).MatchAllExcludedTags(matchAllExcludedTags).IsReportDownload(isReportDownload).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksWithCards(includeTasksWithCards).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCustomFields(includeCustomFields).IncludeCrmDealIds(includeCrmDealIds).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompanyUserIds(includeCompanyUserIds).IncludeBlockedTasks(includeBlockedTasks).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).ExcludeAssigneeNotOnProjectTeams(excludeAssigneeNotOnProjectTeams).AllowTemplateTasks(allowTemplateTasks).TasklistId(tasklistId).TaskIdList(taskIdList).Tags(tags).TagIds(tagIds).SkipCrmDealIds(skipCrmDealIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ProjectCategoryId2(projectCategoryId2).FollowerIds(followerIds).FilterText(filterText).FilterBoardColumnIds(filterBoardColumnIds).ExcludeTagIds(excludeTagIds).CrmDealIds(crmDealIds).CreatorIds(creatorIds).CompanyId(companyId).AssigneeTeamIds(assigneeTeamIds).AssigneeCompanyIds(assigneeCompanyIds).Execute()

Get Tasks

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
    startDate := "startDate_example" // string |  (optional)
    sortOrder := "sortOrder_example" // string |  (optional)
    sort := "sort_example" // string |  (optional)
    reportFormat := "reportFormat_example" // string |  (optional)
    projectStatus := "projectStatus_example" // string |  (optional)
    priority := "priority_example" // string |  (optional)
    include := "include_example" // string |  (optional)
    filter := "filter_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    deletedAfterDate := "deletedAfterDate_example" // string |  (optional)
    dateCode := "dateCode_example" // string |  (optional)
    dataSet := "dataSet_example" // string |  (optional)
    createdFilter := "createdFilter_example" // string |  (optional)
    createdDateCode := "createdDateCode_example" // string |  (optional)
    createdBeforeDate := "createdBeforeDate_example" // string |  (optional)
    createdAfterDate := "createdAfterDate_example" // string |  (optional)
    completedBeforeDate := "completedBeforeDate_example" // string |  (optional)
    completedAfterDate := "completedAfterDate_example" // string |  (optional)
    callback := "callback_example" // string |  (optional)
    taskId := int32(56) // int32 |  (optional)
    parentTaskId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    offsetCount := int32(56) // int32 |  (optional)
    includeTaskId := int32(56) // int32 |  (optional)
    useStartDatesForTodaysTasks := true // bool |  (optional)
    useAllProjects := true // bool |  (optional)
    starredProjectsOnly := true // bool |  (optional)
    showDeleted := true // bool |  (optional)
    showCompletedLists := true // bool |  (optional)
    onlyUntaggedTasks := true // bool |  (optional)
    onlyTasksWithUnreadComments := true // bool |  (optional)
    onlyTasksWithTickets := true // bool |  (optional)
    onlyTasksWithEstimatedTime := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    nestSubTasks := true // bool |  (optional)
    matchAllTags := true // bool |  (optional)
    matchAllProjectTags := true // bool |  (optional)
    matchAllProjectTagIds := true // bool |  (optional)
    matchAllExcludedTags := true // bool |  (optional)
    isReportDownload := true // bool |  (optional)
    includeUntaggedTasks := true // bool |  (optional)
    includeToday := true // bool |  (optional)
    includeTeamUserIds := true // bool |  (optional)
    includeTasksWithoutDueDates := true // bool |  (optional)
    includeTasksWithCards := true // bool |  (optional)
    includeTasksFromDeletedLists := true // bool |  (optional)
    includeReminders := true // bool |  (optional)
    includeLoggedTime := true // bool |  (optional)
    includeCustomFields := true // bool |  (optional)
    includeCrmDealIds := true // bool |  (optional)
    includeCompletedTasks := true // bool |  (optional)
    includeCompletedSubtasks := true // bool |  (optional)
    includeCompanyUserIds := true // bool |  (optional)
    includeBlockedTasks := true // bool |  (optional)
    includeAssigneeTeams := true // bool |  (optional)
    includeAssigneeCompanies := true // bool |  (optional)
    includeArchivedProjects := true // bool |  (optional)
    ignoreStartDates := true // bool |  (optional)
    getSubTasks := true // bool |  (optional)
    getFiles := true // bool |  (optional)
    excludeAssigneeNotOnProjectTeams := true // bool |  (optional)
    allowTemplateTasks := true // bool |  (optional)
    tasklistId := []int32{int32(123)} // []int32 |  (optional)
    taskIdList := []int32{int32(123)} // []int32 |  (optional)
    tags := []string{"Inner_example"} // []string |  (optional)
    tagIds := []int32{int32(123)} // []int32 |  (optional)
    skipCrmDealIds := []int32{int32(123)} // []int32 |  (optional)
    responsiblePartyIds := []int32{int32(123)} // []int32 |  (optional)
    projectTagIds := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    projectIds := []int32{int32(123)} // []int32 |  (optional)
    projectId := []int32{int32(123)} // []int32 | project filters (optional)
    projectHealths := []int32{int32(123)} // []int32 |  (optional)
    projectHealth := []int32{int32(123)} // []int32 |  (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryId2 := []int32{int32(123)} // []int32 |  (optional)
    followerIds := []int32{int32(123)} // []int32 |  (optional)
    filterText := []string{"Inner_example"} // []string |  (optional)
    filterBoardColumnIds := []int32{int32(123)} // []int32 |  (optional)
    excludeTagIds := []int32{int32(123)} // []int32 |  (optional)
    crmDealIds := []int32{int32(123)} // []int32 |  (optional)
    creatorIds := []int32{int32(123)} // []int32 |  (optional)
    companyId := []int32{int32(123)} // []int32 |  (optional)
    assigneeTeamIds := []int32{int32(123)} // []int32 |  (optional)
    assigneeCompanyIds := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2ProjectcategoriesprojectCategoryIdTasks(context.Background(), projectCategoryId).UpdatedAfterDate(updatedAfterDate).Today(today).StartDate(startDate).SortOrder(sortOrder).Sort(sort).ReportFormat(reportFormat).ProjectStatus(projectStatus).Priority(priority).Include(include).Filter(filter).EndDate(endDate).DeletedAfterDate(deletedAfterDate).DateCode(dateCode).DataSet(dataSet).CreatedFilter(createdFilter).CreatedDateCode(createdDateCode).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TaskId(taskId).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).OffsetCount(offsetCount).IncludeTaskId(includeTaskId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyTasksWithUnreadComments(onlyTasksWithUnreadComments).OnlyTasksWithTickets(onlyTasksWithTickets).OnlyTasksWithEstimatedTime(onlyTasksWithEstimatedTime).OnlyStarredProjects(onlyStarredProjects).NestSubTasks(nestSubTasks).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTagIds(matchAllProjectTagIds).MatchAllExcludedTags(matchAllExcludedTags).IsReportDownload(isReportDownload).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksWithCards(includeTasksWithCards).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCustomFields(includeCustomFields).IncludeCrmDealIds(includeCrmDealIds).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompanyUserIds(includeCompanyUserIds).IncludeBlockedTasks(includeBlockedTasks).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).ExcludeAssigneeNotOnProjectTeams(excludeAssigneeNotOnProjectTeams).AllowTemplateTasks(allowTemplateTasks).TasklistId(tasklistId).TaskIdList(taskIdList).Tags(tags).TagIds(tagIds).SkipCrmDealIds(skipCrmDealIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ProjectCategoryId2(projectCategoryId2).FollowerIds(followerIds).FilterText(filterText).FilterBoardColumnIds(filterBoardColumnIds).ExcludeTagIds(excludeTagIds).CrmDealIds(crmDealIds).CreatorIds(creatorIds).CompanyId(companyId).AssigneeTeamIds(assigneeTeamIds).AssigneeCompanyIds(assigneeCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2ProjectcategoriesprojectCategoryIdTasks``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETProjectsApiV2ProjectcategoriesprojectCategoryIdTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedAfterDate** | **string** |  | 
 **today** | **string** |  | 
 **startDate** | **string** |  | 
 **sortOrder** | **string** |  | 
 **sort** | **string** |  | 
 **reportFormat** | **string** |  | 
 **projectStatus** | **string** |  | 
 **priority** | **string** |  | 
 **include** | **string** |  | 
 **filter** | **string** |  | 
 **endDate** | **string** |  | 
 **deletedAfterDate** | **string** |  | 
 **dateCode** | **string** |  | 
 **dataSet** | **string** |  | 
 **createdFilter** | **string** |  | 
 **createdDateCode** | **string** |  | 
 **createdBeforeDate** | **string** |  | 
 **createdAfterDate** | **string** |  | 
 **completedBeforeDate** | **string** |  | 
 **completedAfterDate** | **string** |  | 
 **callback** | **string** |  | 
 **taskId** | **int32** |  | 
 **parentTaskId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **offsetCount** | **int32** |  | 
 **includeTaskId** | **int32** |  | 
 **useStartDatesForTodaysTasks** | **bool** |  | 
 **useAllProjects** | **bool** |  | 
 **starredProjectsOnly** | **bool** |  | 
 **showDeleted** | **bool** |  | 
 **showCompletedLists** | **bool** |  | 
 **onlyUntaggedTasks** | **bool** |  | 
 **onlyTasksWithUnreadComments** | **bool** |  | 
 **onlyTasksWithTickets** | **bool** |  | 
 **onlyTasksWithEstimatedTime** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **nestSubTasks** | **bool** |  | 
 **matchAllTags** | **bool** |  | 
 **matchAllProjectTags** | **bool** |  | 
 **matchAllProjectTagIds** | **bool** |  | 
 **matchAllExcludedTags** | **bool** |  | 
 **isReportDownload** | **bool** |  | 
 **includeUntaggedTasks** | **bool** |  | 
 **includeToday** | **bool** |  | 
 **includeTeamUserIds** | **bool** |  | 
 **includeTasksWithoutDueDates** | **bool** |  | 
 **includeTasksWithCards** | **bool** |  | 
 **includeTasksFromDeletedLists** | **bool** |  | 
 **includeReminders** | **bool** |  | 
 **includeLoggedTime** | **bool** |  | 
 **includeCustomFields** | **bool** |  | 
 **includeCrmDealIds** | **bool** |  | 
 **includeCompletedTasks** | **bool** |  | 
 **includeCompletedSubtasks** | **bool** |  | 
 **includeCompanyUserIds** | **bool** |  | 
 **includeBlockedTasks** | **bool** |  | 
 **includeAssigneeTeams** | **bool** |  | 
 **includeAssigneeCompanies** | **bool** |  | 
 **includeArchivedProjects** | **bool** |  | 
 **ignoreStartDates** | **bool** |  | 
 **getSubTasks** | **bool** |  | 
 **getFiles** | **bool** |  | 
 **excludeAssigneeNotOnProjectTeams** | **bool** |  | 
 **allowTemplateTasks** | **bool** |  | 
 **tasklistId** | **[]int32** |  | 
 **taskIdList** | **[]int32** |  | 
 **tags** | **[]string** |  | 
 **tagIds** | **[]int32** |  | 
 **skipCrmDealIds** | **[]int32** |  | 
 **responsiblePartyIds** | **[]int32** |  | 
 **projectTagIds** | **[]int32** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **projectIds** | **[]int32** |  | 
 **projectId** | **[]int32** | project filters | 
 **projectHealths** | **[]int32** |  | 
 **projectHealth** | **[]int32** |  | 
 **projectCompanyIds** | **[]int32** |  | 
 **projectCategoryIds** | **[]int32** |  | 
 **projectCategoryId2** | **[]int32** |  | 
 **followerIds** | **[]int32** |  | 
 **filterText** | **[]string** |  | 
 **filterBoardColumnIds** | **[]int32** |  | 
 **excludeTagIds** | **[]int32** |  | 
 **crmDealIds** | **[]int32** |  | 
 **creatorIds** | **[]int32** |  | 
 **companyId** | **[]int32** |  | 
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


## GETProjectsApiV2Projects

> GETProjectsApiV2Projects(ctx).WorkedOnAfterDateTime(workedOnAfterDateTime).UpdatedAfterDateTime(updatedAfterDateTime).UpdatedAfterDate(updatedAfterDate).Status(status).SearchTerm(searchTerm).ProjectType(projectType).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).Get(get).FirstLetter(firstLetter).DataSet(dataSet).CreatedAfterDateTime(createdAfterDateTime).CreatedAfterDate(createdAfterDate).CompletedBeforeDateTime(completedBeforeDateTime).CompletedAfterDateTime(completedAfterDateTime).UserId(userId).PageSize(pageSize).Page(page).SearchCompany(searchCompany).SearchCategory(searchCategory).SanitizeName(sanitizeName).ReturnLetters(returnLetters).OnlyStarredProjects(onlyStarredProjects).OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership).OnlyCustomFieldValues(onlyCustomFieldValues).OnlyArchivedProjects(onlyArchivedProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).IncludeUpdates(includeUpdates).IncludeTopPeople(includeTopPeople).IncludeTags(includeTags).IncludeProjectOwner(includeProjectOwner).IncludePeople(includePeople).IncludeCustomFields(includeCustomFields).IncludeCrmDealIds(includeCrmDealIds).IncludeArchivedProjects(includeArchivedProjects).HideObservedProjects(hideObservedProjects).HideDesc(hideDesc).GetTruePermissions(getTruePermissions).GetPermissions(getPermissions).GetNotificationSettings(getNotificationSettings).GetDeleted(getDeleted).GetCounts(getCounts).GetCategoryPath(getCategoryPath).GetAll(getAll).GetActivePages(getActivePages).FormatMarkdown(formatMarkdown).SkipIds(skipIds).SkipCrmDealIds(skipCrmDealIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).PortfolioColumnIds(portfolioColumnIds).PortfolioBoardIds(portfolioBoardIds).FilterTagIds(filterTagIds).FeaturesEnabledOnProject(featuresEnabledOnProject).CrmDealIds(crmDealIds).CompanyId(companyId).CategoryId(categoryId).CatId(catId).Execute()

APICall_GET_projects

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
    workedOnAfterDateTime := "workedOnAfterDateTime_example" // string |  (optional)
    updatedAfterDateTime := "updatedAfterDateTime_example" // string |  (optional)
    updatedAfterDate := "updatedAfterDate_example" // string |  (optional)
    status := "status_example" // string | common project filters (optional)
    searchTerm := "searchTerm_example" // string |  (optional)
    projectType := "projectType_example" // string |  (optional)
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
    searchCategory := true // bool |  (optional)
    sanitizeName := true // bool |  (optional)
    returnLetters := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    onlyProjectsWithExplicitMembership := true // bool |  (optional)
    onlyCustomFieldValues := true // bool |  (optional)
    onlyArchivedProjects := true // bool |  (optional)
    matchAllTags := true // bool |  (optional)
    matchAllProjectTags := true // bool |  (optional)
    includeUpdates := true // bool |  (optional)
    includeTopPeople := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    includeProjectOwner := true // bool |  (optional)
    includePeople := true // bool |  (optional)
    includeCustomFields := true // bool |  (optional)
    includeCrmDealIds := true // bool |  (optional)
    includeArchivedProjects := true // bool |  (optional)
    hideObservedProjects := true // bool |  (optional)
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
    skipCrmDealIds := []int32{int32(123)} // []int32 |  (optional)
    projectTagIds := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    projectHealths := []int32{int32(123)} // []int32 |  (optional)
    projectHealth := []int32{int32(123)} // []int32 |  (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 |  (optional)
    portfolioColumnIds := []int32{int32(123)} // []int32 |  (optional)
    portfolioBoardIds := []int32{int32(123)} // []int32 |  (optional)
    filterTagIds := []int32{int32(123)} // []int32 |  (optional)
    featuresEnabledOnProject := []string{"Inner_example"} // []string |  (optional)
    crmDealIds := []int32{int32(123)} // []int32 |  (optional)
    companyId := []int32{int32(123)} // []int32 |  (optional)
    categoryId := []int32{int32(123)} // []int32 |  (optional)
    catId := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2Projects(context.Background()).WorkedOnAfterDateTime(workedOnAfterDateTime).UpdatedAfterDateTime(updatedAfterDateTime).UpdatedAfterDate(updatedAfterDate).Status(status).SearchTerm(searchTerm).ProjectType(projectType).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).Get(get).FirstLetter(firstLetter).DataSet(dataSet).CreatedAfterDateTime(createdAfterDateTime).CreatedAfterDate(createdAfterDate).CompletedBeforeDateTime(completedBeforeDateTime).CompletedAfterDateTime(completedAfterDateTime).UserId(userId).PageSize(pageSize).Page(page).SearchCompany(searchCompany).SearchCategory(searchCategory).SanitizeName(sanitizeName).ReturnLetters(returnLetters).OnlyStarredProjects(onlyStarredProjects).OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership).OnlyCustomFieldValues(onlyCustomFieldValues).OnlyArchivedProjects(onlyArchivedProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).IncludeUpdates(includeUpdates).IncludeTopPeople(includeTopPeople).IncludeTags(includeTags).IncludeProjectOwner(includeProjectOwner).IncludePeople(includePeople).IncludeCustomFields(includeCustomFields).IncludeCrmDealIds(includeCrmDealIds).IncludeArchivedProjects(includeArchivedProjects).HideObservedProjects(hideObservedProjects).HideDesc(hideDesc).GetTruePermissions(getTruePermissions).GetPermissions(getPermissions).GetNotificationSettings(getNotificationSettings).GetDeleted(getDeleted).GetCounts(getCounts).GetCategoryPath(getCategoryPath).GetAll(getAll).GetActivePages(getActivePages).FormatMarkdown(formatMarkdown).SkipIds(skipIds).SkipCrmDealIds(skipCrmDealIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).PortfolioColumnIds(portfolioColumnIds).PortfolioBoardIds(portfolioBoardIds).FilterTagIds(filterTagIds).FeaturesEnabledOnProject(featuresEnabledOnProject).CrmDealIds(crmDealIds).CompanyId(companyId).CategoryId(categoryId).CatId(catId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2Projects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2ProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workedOnAfterDateTime** | **string** |  | 
 **updatedAfterDateTime** | **string** |  | 
 **updatedAfterDate** | **string** |  | 
 **status** | **string** | common project filters | 
 **searchTerm** | **string** |  | 
 **projectType** | **string** |  | 
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
 **searchCategory** | **bool** |  | 
 **sanitizeName** | **bool** |  | 
 **returnLetters** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **onlyProjectsWithExplicitMembership** | **bool** |  | 
 **onlyCustomFieldValues** | **bool** |  | 
 **onlyArchivedProjects** | **bool** |  | 
 **matchAllTags** | **bool** |  | 
 **matchAllProjectTags** | **bool** |  | 
 **includeUpdates** | **bool** |  | 
 **includeTopPeople** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **includeProjectOwner** | **bool** |  | 
 **includePeople** | **bool** |  | 
 **includeCustomFields** | **bool** |  | 
 **includeCrmDealIds** | **bool** |  | 
 **includeArchivedProjects** | **bool** |  | 
 **hideObservedProjects** | **bool** |  | 
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
 **skipCrmDealIds** | **[]int32** |  | 
 **projectTagIds** | **[]int32** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **projectHealths** | **[]int32** |  | 
 **projectHealth** | **[]int32** |  | 
 **projectCompanyIds** | **[]int32** |  | 
 **projectCategoryIds** | **[]int32** |  | 
 **portfolioColumnIds** | **[]int32** |  | 
 **portfolioBoardIds** | **[]int32** |  | 
 **filterTagIds** | **[]int32** |  | 
 **featuresEnabledOnProject** | **[]string** |  | 
 **crmDealIds** | **[]int32** |  | 
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


## GETProjectsApiV2ProjectsArchived

> GETProjectsApiV2ProjectsArchived(ctx).PageSize(pageSize).Page(page).MaxProjects(maxProjects).IncludeUpdates(includeUpdates).GroupByCompany(groupByCompany).GetEmoji(getEmoji).Execute()

Returns the project defaults

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
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    maxProjects := int32(56) // int32 |  (optional)
    includeUpdates := true // bool |  (optional)
    groupByCompany := true // bool |  (optional)
    getEmoji := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2ProjectsArchived(context.Background()).PageSize(pageSize).Page(page).MaxProjects(maxProjects).IncludeUpdates(includeUpdates).GroupByCompany(groupByCompany).GetEmoji(getEmoji).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2ProjectsArchived``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2ProjectsArchivedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **maxProjects** | **int32** |  | 
 **includeUpdates** | **bool** |  | 
 **groupByCompany** | **bool** |  | 
 **getEmoji** | **bool** |  | 

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


## GETProjectsApiV2ProjectsChart

> GETProjectsApiV2ProjectsChart(ctx).UpdatedAfterDate(updatedAfterDate).StartDate(startDate).EndDate(endDate).CompanyId(companyId).UseMinDates(useMinDates).Starred(starred).OnlyStarredProjects(onlyStarredProjects).OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership).IncludeUpdates(includeUpdates).IncludeTags(includeTags).HideObservedProjects(hideObservedProjects).GetEmoji(getEmoji).CategoryId(categoryId).CatId(catId).Execute()

Returns the project chart

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
    startDate := "startDate_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    companyId := int32(56) // int32 |  (optional)
    useMinDates := true // bool |  (optional)
    starred := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    onlyProjectsWithExplicitMembership := true // bool |  (optional)
    includeUpdates := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    hideObservedProjects := true // bool |  (optional)
    getEmoji := true // bool |  (optional)
    categoryId := []int32{int32(123)} // []int32 |  (optional)
    catId := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2ProjectsChart(context.Background()).UpdatedAfterDate(updatedAfterDate).StartDate(startDate).EndDate(endDate).CompanyId(companyId).UseMinDates(useMinDates).Starred(starred).OnlyStarredProjects(onlyStarredProjects).OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership).IncludeUpdates(includeUpdates).IncludeTags(includeTags).HideObservedProjects(hideObservedProjects).GetEmoji(getEmoji).CategoryId(categoryId).CatId(catId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2ProjectsChart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2ProjectsChartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfterDate** | **string** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **companyId** | **int32** |  | 
 **useMinDates** | **bool** |  | 
 **starred** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **onlyProjectsWithExplicitMembership** | **bool** |  | 
 **includeUpdates** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **hideObservedProjects** | **bool** |  | 
 **getEmoji** | **bool** |  | 
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


## GETProjectsApiV2ProjectsChartoverview

> GETProjectsApiV2ProjectsChartoverview(ctx).UpdatedAfterDate(updatedAfterDate).StartDate(startDate).EndDate(endDate).ProjectCompanyId(projectCompanyId).UseMinDates(useMinDates).OnlyStarredProjects(onlyStarredProjects).IncludeUpdates(includeUpdates).IncludeTags(includeTags).GetEmoji(getEmoji).ProjectCategoryIds(projectCategoryIds).Execute()

Returns the project defaults (Objective: To deprecate getProjectsChart)

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
    startDate := "startDate_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    projectCompanyId := int32(56) // int32 |  (optional)
    useMinDates := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    includeUpdates := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    getEmoji := true // bool |  (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2ProjectsChartoverview(context.Background()).UpdatedAfterDate(updatedAfterDate).StartDate(startDate).EndDate(endDate).ProjectCompanyId(projectCompanyId).UseMinDates(useMinDates).OnlyStarredProjects(onlyStarredProjects).IncludeUpdates(includeUpdates).IncludeTags(includeTags).GetEmoji(getEmoji).ProjectCategoryIds(projectCategoryIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2ProjectsChartoverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2ProjectsChartoverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfterDate** | **string** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **projectCompanyId** | **int32** |  | 
 **useMinDates** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **includeUpdates** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **getEmoji** | **bool** |  | 
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


## GETProjectsApiV2ProjectsEstimatedtimeTotal

> GETProjectsApiV2ProjectsEstimatedtimeTotal(ctx).ToDate(toDate).ProjectStatus(projectStatus).FromDate(fromDate).UserId(userId).PageSize(pageSize).Page(page).Execute()

APICall_GET_projects

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
    toDate := "toDate_example" // string |  (optional)
    projectStatus := "projectStatus_example" // string |  (optional)
    fromDate := "fromDate_example" // string |  (optional)
    userId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2ProjectsEstimatedtimeTotal(context.Background()).ToDate(toDate).ProjectStatus(projectStatus).FromDate(fromDate).UserId(userId).PageSize(pageSize).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2ProjectsEstimatedtimeTotal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2ProjectsEstimatedtimeTotalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **toDate** | **string** |  | 
 **projectStatus** | **string** |  | 
 **fromDate** | **string** |  | 
 **userId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 

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


## GETProjectsApiV2ProjectsLatest

> GETProjectsApiV2ProjectsLatest(ctx).View(view).PageSize(pageSize).Page(page).MinProjects(minProjects).MaxProjects(maxProjects).IncludeUpdates(includeUpdates).IncludeStarredInLatest(includeStarredInLatest).GroupByCompany(groupByCompany).GetEmoji(getEmoji).Execute()

Get latest project

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
    view := "view_example" // string |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    minProjects := int32(56) // int32 |  (optional)
    maxProjects := int32(56) // int32 |  (optional)
    includeUpdates := true // bool |  (optional)
    includeStarredInLatest := true // bool |  (optional)
    groupByCompany := true // bool |  (optional)
    getEmoji := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2ProjectsLatest(context.Background()).View(view).PageSize(pageSize).Page(page).MinProjects(minProjects).MaxProjects(maxProjects).IncludeUpdates(includeUpdates).IncludeStarredInLatest(includeStarredInLatest).GroupByCompany(groupByCompany).GetEmoji(getEmoji).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2ProjectsLatest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2ProjectsLatestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **view** | **string** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **minProjects** | **int32** |  | 
 **maxProjects** | **int32** |  | 
 **includeUpdates** | **bool** |  | 
 **includeStarredInLatest** | **bool** |  | 
 **groupByCompany** | **bool** |  | 
 **getEmoji** | **bool** |  | 

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


## GETProjectsApiV2ProjectsMeLatest

> GETProjectsApiV2ProjectsMeLatest(ctx).MaxProjects(maxProjects).IncludeUpdates(includeUpdates).GroupByCompany(groupByCompany).GetEmoji(getEmoji).Execute()

Returns the latest project

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
    maxProjects := int32(56) // int32 |  (optional)
    includeUpdates := true // bool |  (optional)
    groupByCompany := true // bool |  (optional)
    getEmoji := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2ProjectsMeLatest(context.Background()).MaxProjects(maxProjects).IncludeUpdates(includeUpdates).GroupByCompany(groupByCompany).GetEmoji(getEmoji).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2ProjectsMeLatest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2ProjectsMeLatestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maxProjects** | **int32** |  | 
 **includeUpdates** | **bool** |  | 
 **groupByCompany** | **bool** |  | 
 **getEmoji** | **bool** |  | 

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


## GETProjectsApiV2ProjectsNew

> GETProjectsApiV2ProjectsNew(ctx).ProjectId(projectId).Execute()

Gets the milestone timeline for projects

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
    projectId := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2ProjectsNew(context.Background()).ProjectId(projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2ProjectsNew``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2ProjectsNewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **[]int32** |  | 

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


## GETProjectsApiV2ProjectsStarred

> GETProjectsApiV2ProjectsStarred(ctx).DataSet(dataSet).IncludeUpdates(includeUpdates).Execute()

Returns the starred projects

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
    dataSet := "dataSet_example" // string |  (optional)
    includeUpdates := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2ProjectsStarred(context.Background()).DataSet(dataSet).IncludeUpdates(includeUpdates).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2ProjectsStarred``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2ProjectsStarredRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSet** | **string** |  | 
 **includeUpdates** | **bool** |  | 

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


## GETProjectsApiV2ProjectsTemplates

> GETProjectsApiV2ProjectsTemplates(ctx).SearchTerm(searchTerm).OrderMode(orderMode).OrderBy(orderBy).UserId(userId).PageSize(pageSize).Page(page).MatchAllProjectTags(matchAllProjectTags).IncludeProjectOwner(includeProjectOwner).GetTruePermissions(getTruePermissions).GetPermissions(getPermissions).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).FilterTagIds(filterTagIds).Execute()

Returns the project defaults

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
    searchTerm := "searchTerm_example" // string |  (optional)
    orderMode := "orderMode_example" // string |  (optional)
    orderBy := "orderBy_example" // string |  (optional)
    userId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    matchAllProjectTags := true // bool |  (optional)
    includeProjectOwner := true // bool |  (optional)
    getTruePermissions := true // bool |  (optional)
    getPermissions := true // bool |  (optional)
    projectTagIds := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    filterTagIds := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2ProjectsTemplates(context.Background()).SearchTerm(searchTerm).OrderMode(orderMode).OrderBy(orderBy).UserId(userId).PageSize(pageSize).Page(page).MatchAllProjectTags(matchAllProjectTags).IncludeProjectOwner(includeProjectOwner).GetTruePermissions(getTruePermissions).GetPermissions(getPermissions).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).FilterTagIds(filterTagIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2ProjectsTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2ProjectsTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchTerm** | **string** |  | 
 **orderMode** | **string** |  | 
 **orderBy** | **string** |  | 
 **userId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **matchAllProjectTags** | **bool** |  | 
 **includeProjectOwner** | **bool** |  | 
 **getTruePermissions** | **bool** |  | 
 **getPermissions** | **bool** |  | 
 **projectTagIds** | **[]int32** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **filterTagIds** | **[]int32** |  | 

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


## GETProjectsApiV2ProjectsTimeTotal

> GETProjectsApiV2ProjectsTimeTotal(ctx).ToDate(toDate).ProjectStatus(projectStatus).FromDate(fromDate).UserId(userId).PageSize(pageSize).Page(page).Execute()

Get Projects Total Time

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
    toDate := "toDate_example" // string |  (optional)
    projectStatus := "projectStatus_example" // string |  (optional)
    fromDate := "fromDate_example" // string |  (optional)
    userId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2ProjectsTimeTotal(context.Background()).ToDate(toDate).ProjectStatus(projectStatus).FromDate(fromDate).UserId(userId).PageSize(pageSize).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2ProjectsTimeTotal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2ProjectsTimeTotalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **toDate** | **string** |  | 
 **projectStatus** | **string** |  | 
 **fromDate** | **string** |  | 
 **userId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 

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


## GETProjectsApiV2ProjectsprojectId

> GETProjectsApiV2ProjectsprojectId(ctx, projectId2).UserId(userId).ProjectId(projectId).ShowDeleted(showDeleted).OnlyCustomFieldValues(onlyCustomFieldValues).IncludeUpdates(includeUpdates).IncludeTags(includeTags).IncludeProjectOwner(includeProjectOwner).IncludePeople(includePeople).IncludeCustomFields(includeCustomFields).IncludeCrmDealIds(includeCrmDealIds).GetTruePermissions(getTruePermissions).GetPermissions(getPermissions).GetNotificationSettings(getNotificationSettings).GetEmoji(getEmoji).GetEmailAddress(getEmailAddress).GetDateInfo(getDateInfo).GetActivePages(getActivePages).FormatMarkdown(formatMarkdown).ProjectOwnerIds(projectOwnerIds).ProjectHealths(projectHealths).ProjectHealth(projectHealth).Execute()

Returns the project defaults

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
    projectId2 := int32(56) // int32 | 
    userId := int32(56) // int32 |  (optional)
    projectId := int32(56) // int32 |  (optional)
    showDeleted := true // bool |  (optional)
    onlyCustomFieldValues := true // bool |  (optional)
    includeUpdates := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    includeProjectOwner := true // bool |  (optional)
    includePeople := true // bool |  (optional)
    includeCustomFields := true // bool |  (optional)
    includeCrmDealIds := true // bool |  (optional)
    getTruePermissions := true // bool |  (optional)
    getPermissions := true // bool |  (optional)
    getNotificationSettings := true // bool |  (optional)
    getEmoji := true // bool |  (optional)
    getEmailAddress := true // bool |  (optional)
    getDateInfo := true // bool |  (optional)
    getActivePages := true // bool |  (optional)
    formatMarkdown := true // bool |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    projectHealths := []int32{int32(123)} // []int32 |  (optional)
    projectHealth := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2ProjectsprojectId(context.Background(), projectId2).UserId(userId).ProjectId(projectId).ShowDeleted(showDeleted).OnlyCustomFieldValues(onlyCustomFieldValues).IncludeUpdates(includeUpdates).IncludeTags(includeTags).IncludeProjectOwner(includeProjectOwner).IncludePeople(includePeople).IncludeCustomFields(includeCustomFields).IncludeCrmDealIds(includeCrmDealIds).GetTruePermissions(getTruePermissions).GetPermissions(getPermissions).GetNotificationSettings(getNotificationSettings).GetEmoji(getEmoji).GetEmailAddress(getEmailAddress).GetDateInfo(getDateInfo).GetActivePages(getActivePages).FormatMarkdown(formatMarkdown).ProjectOwnerIds(projectOwnerIds).ProjectHealths(projectHealths).ProjectHealth(projectHealth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2ProjectsprojectId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2ProjectsprojectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **int32** |  | 
 **projectId** | **int32** |  | 
 **showDeleted** | **bool** |  | 
 **onlyCustomFieldValues** | **bool** |  | 
 **includeUpdates** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **includeProjectOwner** | **bool** |  | 
 **includePeople** | **bool** |  | 
 **includeCustomFields** | **bool** |  | 
 **includeCrmDealIds** | **bool** |  | 
 **getTruePermissions** | **bool** |  | 
 **getPermissions** | **bool** |  | 
 **getNotificationSettings** | **bool** |  | 
 **getEmoji** | **bool** |  | 
 **getEmailAddress** | **bool** |  | 
 **getDateInfo** | **bool** |  | 
 **getActivePages** | **bool** |  | 
 **formatMarkdown** | **bool** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **projectHealths** | **[]int32** |  | 
 **projectHealth** | **[]int32** |  | 

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


## GETProjectsApiV2ProjectsprojectIdBillingTime

> GETProjectsApiV2ProjectsprojectIdBillingTime(ctx, projectId2).UpdatedAfterDate(updatedAfterDate).ToDate(toDate).SortOrder(sortOrder).SortBy(sortBy).Permission(permission).InvoicedType(invoicedType).FromDate(fromDate).BillableType(billableType).TicketId(ticketId).ProjectId(projectId).PageSize(pageSize).Page(page).InvoiceId(invoiceId).IncludeTimelogId(includeTimelogId).ShowDeleted(showDeleted).MatchAllTags(matchAllTags).IncludeTags(includeTags).GetTotals(getTotals).UserId(userId).TaskTagIds(taskTagIds).TagIds(tagIds).CompanyId(companyId).Execute()

APICall_GET_projects_id_time_entries

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
    projectId2 := int32(56) // int32 | 
    updatedAfterDate := "updatedAfterDate_example" // string |  (optional)
    toDate := "toDate_example" // string |  (optional)
    sortOrder := "sortOrder_example" // string |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    permission := "permission_example" // string |  (optional)
    invoicedType := "invoicedType_example" // string |  (optional)
    fromDate := "fromDate_example" // string |  (optional)
    billableType := "billableType_example" // string |  (optional)
    ticketId := int32(56) // int32 |  (optional)
    projectId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    invoiceId := int32(56) // int32 |  (optional)
    includeTimelogId := int32(56) // int32 |  (optional)
    showDeleted := true // bool |  (optional)
    matchAllTags := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    getTotals := true // bool |  (optional)
    userId := []int32{int32(123)} // []int32 |  (optional)
    taskTagIds := []int32{int32(123)} // []int32 |  (optional)
    tagIds := []int32{int32(123)} // []int32 |  (optional)
    companyId := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2ProjectsprojectIdBillingTime(context.Background(), projectId2).UpdatedAfterDate(updatedAfterDate).ToDate(toDate).SortOrder(sortOrder).SortBy(sortBy).Permission(permission).InvoicedType(invoicedType).FromDate(fromDate).BillableType(billableType).TicketId(ticketId).ProjectId(projectId).PageSize(pageSize).Page(page).InvoiceId(invoiceId).IncludeTimelogId(includeTimelogId).ShowDeleted(showDeleted).MatchAllTags(matchAllTags).IncludeTags(includeTags).GetTotals(getTotals).UserId(userId).TaskTagIds(taskTagIds).TagIds(tagIds).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2ProjectsprojectIdBillingTime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2ProjectsprojectIdBillingTimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedAfterDate** | **string** |  | 
 **toDate** | **string** |  | 
 **sortOrder** | **string** |  | 
 **sortBy** | **string** |  | 
 **permission** | **string** |  | 
 **invoicedType** | **string** |  | 
 **fromDate** | **string** |  | 
 **billableType** | **string** |  | 
 **ticketId** | **int32** |  | 
 **projectId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **invoiceId** | **int32** |  | 
 **includeTimelogId** | **int32** |  | 
 **showDeleted** | **bool** |  | 
 **matchAllTags** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **getTotals** | **bool** |  | 
 **userId** | **[]int32** |  | 
 **taskTagIds** | **[]int32** |  | 
 **tagIds** | **[]int32** |  | 
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


## GETProjectsApiV2ProjectsprojectIdEmailaddress

> GETProjectsApiV2ProjectsprojectIdEmailaddress(ctx, projectId2).SubCodeType(subCodeType).ProjectId(projectId).CategoryId(categoryId).Setup(setup).GetSubCode(getSubCode).Execute()

Returns the details for the account you’re currently logged in with.

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
    projectId2 := int32(56) // int32 | 
    subCodeType := int32(56) // int32 |  (optional)
    projectId := int32(56) // int32 |  (optional)
    categoryId := int32(56) // int32 |  (optional)
    setup := true // bool |  (optional)
    getSubCode := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2ProjectsprojectIdEmailaddress(context.Background(), projectId2).SubCodeType(subCodeType).ProjectId(projectId).CategoryId(categoryId).Setup(setup).GetSubCode(getSubCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2ProjectsprojectIdEmailaddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2ProjectsprojectIdEmailaddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subCodeType** | **int32** |  | 
 **projectId** | **int32** |  | 
 **categoryId** | **int32** |  | 
 **setup** | **bool** |  | 
 **getSubCode** | **bool** |  | 

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


## GETProjectsApiV2ProjectsprojectIdPeople

> PeopleResponse GETProjectsApiV2ProjectsprojectIdPeople(ctx, projectId).UpdatedAfterDateTime(updatedAfterDateTime).UpdatedAfterDate(updatedAfterDate).Type_(type_).SortOrder(sortOrder).Sort(sort).SearchTerm(searchTerm).ProjectType(projectType).InviteStatus(inviteStatus).FirstLetter(firstLetter).EmailAddress(emailAddress).DataSet(dataSet).ProjectId2(projectId2).ProjectCompanyId(projectCompanyId).PageSize(pageSize).Page(page).ExcludeProjectId(excludeProjectId).ShowDeleted(showDeleted).ReturnTeams(returnTeams).ReturnProjectIds(returnProjectIds).ReturnLetters(returnLetters).OwnerCompanyFirst(ownerCompanyFirst).Onlyids(onlyids).IncludeTags(includeTags).IncludeObservers(includeObservers).IncludeContacts(includeContacts).IncludeCompanyDetails(includeCompanyDetails).IncludeClockIn(includeClockIn).IncludeClients(includeClients).GroupByCompany(groupByCompany).GetProjectRoles(getProjectRoles).GetCounts(getCounts).Fullprofile(fullprofile).CountOnly(countOnly).UserIds(userIds).CompanyId(companyId).Execute()

Retrieves all of the people in a given project.

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
    userIds := []int32{int32(123)} // []int32 |  (optional)
    companyId := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2ProjectsprojectIdPeople(context.Background(), projectId).UpdatedAfterDateTime(updatedAfterDateTime).UpdatedAfterDate(updatedAfterDate).Type_(type_).SortOrder(sortOrder).Sort(sort).SearchTerm(searchTerm).ProjectType(projectType).InviteStatus(inviteStatus).FirstLetter(firstLetter).EmailAddress(emailAddress).DataSet(dataSet).ProjectId2(projectId2).ProjectCompanyId(projectCompanyId).PageSize(pageSize).Page(page).ExcludeProjectId(excludeProjectId).ShowDeleted(showDeleted).ReturnTeams(returnTeams).ReturnProjectIds(returnProjectIds).ReturnLetters(returnLetters).OwnerCompanyFirst(ownerCompanyFirst).Onlyids(onlyids).IncludeTags(includeTags).IncludeObservers(includeObservers).IncludeContacts(includeContacts).IncludeCompanyDetails(includeCompanyDetails).IncludeClockIn(includeClockIn).IncludeClients(includeClients).GroupByCompany(groupByCompany).GetProjectRoles(getProjectRoles).GetCounts(getCounts).Fullprofile(fullprofile).CountOnly(countOnly).UserIds(userIds).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2ProjectsprojectIdPeople``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV2ProjectsprojectIdPeople`: PeopleResponse
    fmt.Fprintf(os.Stdout, "Response from `CFPortApi.GETProjectsApiV2ProjectsprojectIdPeople`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2ProjectsprojectIdPeopleRequest struct via the builder pattern


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
 **userIds** | **[]int32** |  | 
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


## GETProjectsApiV2ProjectsprojectIdPredecessors

> GETProjectsApiV2ProjectsprojectIdPredecessors(ctx, projectId2).Include(include).ProjectId(projectId).PageSize(pageSize).Page(page).DependentTaskId(dependentTaskId).OnlyDependencies(onlyDependencies).OnlyBasicFields(onlyBasicFields).OnlyActionable(onlyActionable).IncludeDependencies(includeDependencies).GetRecursively(getRecursively).Execute()

Get Predecessor

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
    projectId2 := int32(56) // int32 | 
    include := "include_example" // string |  (optional)
    projectId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    dependentTaskId := int32(56) // int32 |  (optional)
    onlyDependencies := true // bool |  (optional)
    onlyBasicFields := true // bool |  (optional)
    onlyActionable := true // bool |  (optional)
    includeDependencies := true // bool |  (optional)
    getRecursively := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2ProjectsprojectIdPredecessors(context.Background(), projectId2).Include(include).ProjectId(projectId).PageSize(pageSize).Page(page).DependentTaskId(dependentTaskId).OnlyDependencies(onlyDependencies).OnlyBasicFields(onlyBasicFields).OnlyActionable(onlyActionable).IncludeDependencies(includeDependencies).GetRecursively(getRecursively).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2ProjectsprojectIdPredecessors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2ProjectsprojectIdPredecessorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **string** |  | 
 **projectId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **dependentTaskId** | **int32** |  | 
 **onlyDependencies** | **bool** |  | 
 **onlyBasicFields** | **bool** |  | 
 **onlyActionable** | **bool** |  | 
 **includeDependencies** | **bool** |  | 
 **getRecursively** | **bool** |  | 

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


## GETProjectsApiV2ProjectsprojectIdSummary

> GETProjectsApiV2ProjectsprojectIdSummary(ctx, projectId).FilterText(filterText).APIReturnType(aPIReturnType).PageSize(pageSize).Page(page).GetEmoji(getEmoji).Execute()

APICall_GET_summary

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
    filterText := "filterText_example" // string |  (optional)
    aPIReturnType := "aPIReturnType_example" // string |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    getEmoji := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2ProjectsprojectIdSummary(context.Background(), projectId).FilterText(filterText).APIReturnType(aPIReturnType).PageSize(pageSize).Page(page).GetEmoji(getEmoji).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2ProjectsprojectIdSummary``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETProjectsApiV2ProjectsprojectIdSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterText** | **string** |  | 
 **aPIReturnType** | **string** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **getEmoji** | **bool** |  | 

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


## GETProjectsApiV2ProjectsprojectIdTaskids

> GETProjectsApiV2ProjectsprojectIdTaskids(ctx, projectId).UpdatedAfterDate(updatedAfterDate).Today(today).StartDate(startDate).SortOrder(sortOrder).Sort(sort).ReportFormat(reportFormat).ProjectStatus(projectStatus).Priority(priority).Include(include).Filter(filter).EndDate(endDate).DeletedAfterDate(deletedAfterDate).DateCode(dateCode).DataSet(dataSet).CreatedFilter(createdFilter).CreatedDateCode(createdDateCode).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TaskId(taskId).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).OffsetCount(offsetCount).IncludeTaskId(includeTaskId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyTasksWithUnreadComments(onlyTasksWithUnreadComments).OnlyTasksWithTickets(onlyTasksWithTickets).OnlyTasksWithEstimatedTime(onlyTasksWithEstimatedTime).OnlyStarredProjects(onlyStarredProjects).NestSubTasks(nestSubTasks).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTagIds(matchAllProjectTagIds).MatchAllExcludedTags(matchAllExcludedTags).IsReportDownload(isReportDownload).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksWithCards(includeTasksWithCards).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCustomFields(includeCustomFields).IncludeCrmDealIds(includeCrmDealIds).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompanyUserIds(includeCompanyUserIds).IncludeBlockedTasks(includeBlockedTasks).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).ExcludeAssigneeNotOnProjectTeams(excludeAssigneeNotOnProjectTeams).AllowTemplateTasks(allowTemplateTasks).TasklistId(tasklistId).TaskIdList(taskIdList).Tags(tags).TagIds(tagIds).SkipCrmDealIds(skipCrmDealIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId2(projectId2).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ProjectCategoryId(projectCategoryId).FollowerIds(followerIds).FilterText(filterText).FilterBoardColumnIds(filterBoardColumnIds).ExcludeTagIds(excludeTagIds).CrmDealIds(crmDealIds).CreatorIds(creatorIds).CompanyId(companyId).AssigneeTeamIds(assigneeTeamIds).AssigneeCompanyIds(assigneeCompanyIds).Execute()

Get Tasks

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
    startDate := "startDate_example" // string |  (optional)
    sortOrder := "sortOrder_example" // string |  (optional)
    sort := "sort_example" // string |  (optional)
    reportFormat := "reportFormat_example" // string |  (optional)
    projectStatus := "projectStatus_example" // string |  (optional)
    priority := "priority_example" // string |  (optional)
    include := "include_example" // string |  (optional)
    filter := "filter_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    deletedAfterDate := "deletedAfterDate_example" // string |  (optional)
    dateCode := "dateCode_example" // string |  (optional)
    dataSet := "dataSet_example" // string |  (optional)
    createdFilter := "createdFilter_example" // string |  (optional)
    createdDateCode := "createdDateCode_example" // string |  (optional)
    createdBeforeDate := "createdBeforeDate_example" // string |  (optional)
    createdAfterDate := "createdAfterDate_example" // string |  (optional)
    completedBeforeDate := "completedBeforeDate_example" // string |  (optional)
    completedAfterDate := "completedAfterDate_example" // string |  (optional)
    callback := "callback_example" // string |  (optional)
    taskId := int32(56) // int32 |  (optional)
    parentTaskId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    offsetCount := int32(56) // int32 |  (optional)
    includeTaskId := int32(56) // int32 |  (optional)
    useStartDatesForTodaysTasks := true // bool |  (optional)
    useAllProjects := true // bool |  (optional)
    starredProjectsOnly := true // bool |  (optional)
    showDeleted := true // bool |  (optional)
    showCompletedLists := true // bool |  (optional)
    onlyUntaggedTasks := true // bool |  (optional)
    onlyTasksWithUnreadComments := true // bool |  (optional)
    onlyTasksWithTickets := true // bool |  (optional)
    onlyTasksWithEstimatedTime := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    nestSubTasks := true // bool |  (optional)
    matchAllTags := true // bool |  (optional)
    matchAllProjectTags := true // bool |  (optional)
    matchAllProjectTagIds := true // bool |  (optional)
    matchAllExcludedTags := true // bool |  (optional)
    isReportDownload := true // bool |  (optional)
    includeUntaggedTasks := true // bool |  (optional)
    includeToday := true // bool |  (optional)
    includeTeamUserIds := true // bool |  (optional)
    includeTasksWithoutDueDates := true // bool |  (optional)
    includeTasksWithCards := true // bool |  (optional)
    includeTasksFromDeletedLists := true // bool |  (optional)
    includeReminders := true // bool |  (optional)
    includeLoggedTime := true // bool |  (optional)
    includeCustomFields := true // bool |  (optional)
    includeCrmDealIds := true // bool |  (optional)
    includeCompletedTasks := true // bool |  (optional)
    includeCompletedSubtasks := true // bool |  (optional)
    includeCompanyUserIds := true // bool |  (optional)
    includeBlockedTasks := true // bool |  (optional)
    includeAssigneeTeams := true // bool |  (optional)
    includeAssigneeCompanies := true // bool |  (optional)
    includeArchivedProjects := true // bool |  (optional)
    ignoreStartDates := true // bool |  (optional)
    getSubTasks := true // bool |  (optional)
    getFiles := true // bool |  (optional)
    excludeAssigneeNotOnProjectTeams := true // bool |  (optional)
    allowTemplateTasks := true // bool |  (optional)
    tasklistId := []int32{int32(123)} // []int32 |  (optional)
    taskIdList := []int32{int32(123)} // []int32 |  (optional)
    tags := []string{"Inner_example"} // []string |  (optional)
    tagIds := []int32{int32(123)} // []int32 |  (optional)
    skipCrmDealIds := []int32{int32(123)} // []int32 |  (optional)
    responsiblePartyIds := []int32{int32(123)} // []int32 |  (optional)
    projectTagIds := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    projectIds := []int32{int32(123)} // []int32 |  (optional)
    projectId2 := []int32{int32(123)} // []int32 | project filters (optional)
    projectHealths := []int32{int32(123)} // []int32 |  (optional)
    projectHealth := []int32{int32(123)} // []int32 |  (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryId := []int32{int32(123)} // []int32 |  (optional)
    followerIds := []int32{int32(123)} // []int32 |  (optional)
    filterText := []string{"Inner_example"} // []string |  (optional)
    filterBoardColumnIds := []int32{int32(123)} // []int32 |  (optional)
    excludeTagIds := []int32{int32(123)} // []int32 |  (optional)
    crmDealIds := []int32{int32(123)} // []int32 |  (optional)
    creatorIds := []int32{int32(123)} // []int32 |  (optional)
    companyId := []int32{int32(123)} // []int32 |  (optional)
    assigneeTeamIds := []int32{int32(123)} // []int32 |  (optional)
    assigneeCompanyIds := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2ProjectsprojectIdTaskids(context.Background(), projectId).UpdatedAfterDate(updatedAfterDate).Today(today).StartDate(startDate).SortOrder(sortOrder).Sort(sort).ReportFormat(reportFormat).ProjectStatus(projectStatus).Priority(priority).Include(include).Filter(filter).EndDate(endDate).DeletedAfterDate(deletedAfterDate).DateCode(dateCode).DataSet(dataSet).CreatedFilter(createdFilter).CreatedDateCode(createdDateCode).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TaskId(taskId).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).OffsetCount(offsetCount).IncludeTaskId(includeTaskId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyTasksWithUnreadComments(onlyTasksWithUnreadComments).OnlyTasksWithTickets(onlyTasksWithTickets).OnlyTasksWithEstimatedTime(onlyTasksWithEstimatedTime).OnlyStarredProjects(onlyStarredProjects).NestSubTasks(nestSubTasks).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTagIds(matchAllProjectTagIds).MatchAllExcludedTags(matchAllExcludedTags).IsReportDownload(isReportDownload).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksWithCards(includeTasksWithCards).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCustomFields(includeCustomFields).IncludeCrmDealIds(includeCrmDealIds).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompanyUserIds(includeCompanyUserIds).IncludeBlockedTasks(includeBlockedTasks).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).ExcludeAssigneeNotOnProjectTeams(excludeAssigneeNotOnProjectTeams).AllowTemplateTasks(allowTemplateTasks).TasklistId(tasklistId).TaskIdList(taskIdList).Tags(tags).TagIds(tagIds).SkipCrmDealIds(skipCrmDealIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId2(projectId2).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ProjectCategoryId(projectCategoryId).FollowerIds(followerIds).FilterText(filterText).FilterBoardColumnIds(filterBoardColumnIds).ExcludeTagIds(excludeTagIds).CrmDealIds(crmDealIds).CreatorIds(creatorIds).CompanyId(companyId).AssigneeTeamIds(assigneeTeamIds).AssigneeCompanyIds(assigneeCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2ProjectsprojectIdTaskids``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETProjectsApiV2ProjectsprojectIdTaskidsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedAfterDate** | **string** |  | 
 **today** | **string** |  | 
 **startDate** | **string** |  | 
 **sortOrder** | **string** |  | 
 **sort** | **string** |  | 
 **reportFormat** | **string** |  | 
 **projectStatus** | **string** |  | 
 **priority** | **string** |  | 
 **include** | **string** |  | 
 **filter** | **string** |  | 
 **endDate** | **string** |  | 
 **deletedAfterDate** | **string** |  | 
 **dateCode** | **string** |  | 
 **dataSet** | **string** |  | 
 **createdFilter** | **string** |  | 
 **createdDateCode** | **string** |  | 
 **createdBeforeDate** | **string** |  | 
 **createdAfterDate** | **string** |  | 
 **completedBeforeDate** | **string** |  | 
 **completedAfterDate** | **string** |  | 
 **callback** | **string** |  | 
 **taskId** | **int32** |  | 
 **parentTaskId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **offsetCount** | **int32** |  | 
 **includeTaskId** | **int32** |  | 
 **useStartDatesForTodaysTasks** | **bool** |  | 
 **useAllProjects** | **bool** |  | 
 **starredProjectsOnly** | **bool** |  | 
 **showDeleted** | **bool** |  | 
 **showCompletedLists** | **bool** |  | 
 **onlyUntaggedTasks** | **bool** |  | 
 **onlyTasksWithUnreadComments** | **bool** |  | 
 **onlyTasksWithTickets** | **bool** |  | 
 **onlyTasksWithEstimatedTime** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **nestSubTasks** | **bool** |  | 
 **matchAllTags** | **bool** |  | 
 **matchAllProjectTags** | **bool** |  | 
 **matchAllProjectTagIds** | **bool** |  | 
 **matchAllExcludedTags** | **bool** |  | 
 **isReportDownload** | **bool** |  | 
 **includeUntaggedTasks** | **bool** |  | 
 **includeToday** | **bool** |  | 
 **includeTeamUserIds** | **bool** |  | 
 **includeTasksWithoutDueDates** | **bool** |  | 
 **includeTasksWithCards** | **bool** |  | 
 **includeTasksFromDeletedLists** | **bool** |  | 
 **includeReminders** | **bool** |  | 
 **includeLoggedTime** | **bool** |  | 
 **includeCustomFields** | **bool** |  | 
 **includeCrmDealIds** | **bool** |  | 
 **includeCompletedTasks** | **bool** |  | 
 **includeCompletedSubtasks** | **bool** |  | 
 **includeCompanyUserIds** | **bool** |  | 
 **includeBlockedTasks** | **bool** |  | 
 **includeAssigneeTeams** | **bool** |  | 
 **includeAssigneeCompanies** | **bool** |  | 
 **includeArchivedProjects** | **bool** |  | 
 **ignoreStartDates** | **bool** |  | 
 **getSubTasks** | **bool** |  | 
 **getFiles** | **bool** |  | 
 **excludeAssigneeNotOnProjectTeams** | **bool** |  | 
 **allowTemplateTasks** | **bool** |  | 
 **tasklistId** | **[]int32** |  | 
 **taskIdList** | **[]int32** |  | 
 **tags** | **[]string** |  | 
 **tagIds** | **[]int32** |  | 
 **skipCrmDealIds** | **[]int32** |  | 
 **responsiblePartyIds** | **[]int32** |  | 
 **projectTagIds** | **[]int32** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **projectIds** | **[]int32** |  | 
 **projectId2** | **[]int32** | project filters | 
 **projectHealths** | **[]int32** |  | 
 **projectHealth** | **[]int32** |  | 
 **projectCompanyIds** | **[]int32** |  | 
 **projectCategoryIds** | **[]int32** |  | 
 **projectCategoryId** | **[]int32** |  | 
 **followerIds** | **[]int32** |  | 
 **filterText** | **[]string** |  | 
 **filterBoardColumnIds** | **[]int32** |  | 
 **excludeTagIds** | **[]int32** |  | 
 **crmDealIds** | **[]int32** |  | 
 **creatorIds** | **[]int32** |  | 
 **companyId** | **[]int32** |  | 
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


## GETProjectsApiV2ProjectsprojectIdTasks

> GETProjectsApiV2ProjectsprojectIdTasks(ctx, projectId).UpdatedAfterDate(updatedAfterDate).Today(today).StartDate(startDate).SortOrder(sortOrder).Sort(sort).ReportFormat(reportFormat).ProjectStatus(projectStatus).Priority(priority).Include(include).Filter(filter).EndDate(endDate).DeletedAfterDate(deletedAfterDate).DateCode(dateCode).DataSet(dataSet).CreatedFilter(createdFilter).CreatedDateCode(createdDateCode).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TaskId(taskId).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).OffsetCount(offsetCount).IncludeTaskId(includeTaskId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyTasksWithUnreadComments(onlyTasksWithUnreadComments).OnlyTasksWithTickets(onlyTasksWithTickets).OnlyTasksWithEstimatedTime(onlyTasksWithEstimatedTime).OnlyStarredProjects(onlyStarredProjects).NestSubTasks(nestSubTasks).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTagIds(matchAllProjectTagIds).MatchAllExcludedTags(matchAllExcludedTags).IsReportDownload(isReportDownload).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksWithCards(includeTasksWithCards).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCustomFields(includeCustomFields).IncludeCrmDealIds(includeCrmDealIds).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompanyUserIds(includeCompanyUserIds).IncludeBlockedTasks(includeBlockedTasks).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).ExcludeAssigneeNotOnProjectTeams(excludeAssigneeNotOnProjectTeams).AllowTemplateTasks(allowTemplateTasks).TasklistId(tasklistId).TaskIdList(taskIdList).Tags(tags).TagIds(tagIds).SkipCrmDealIds(skipCrmDealIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId2(projectId2).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ProjectCategoryId(projectCategoryId).FollowerIds(followerIds).FilterText(filterText).FilterBoardColumnIds(filterBoardColumnIds).ExcludeTagIds(excludeTagIds).CrmDealIds(crmDealIds).CreatorIds(creatorIds).CompanyId(companyId).AssigneeTeamIds(assigneeTeamIds).AssigneeCompanyIds(assigneeCompanyIds).Execute()

Get Tasks

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
    startDate := "startDate_example" // string |  (optional)
    sortOrder := "sortOrder_example" // string |  (optional)
    sort := "sort_example" // string |  (optional)
    reportFormat := "reportFormat_example" // string |  (optional)
    projectStatus := "projectStatus_example" // string |  (optional)
    priority := "priority_example" // string |  (optional)
    include := "include_example" // string |  (optional)
    filter := "filter_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    deletedAfterDate := "deletedAfterDate_example" // string |  (optional)
    dateCode := "dateCode_example" // string |  (optional)
    dataSet := "dataSet_example" // string |  (optional)
    createdFilter := "createdFilter_example" // string |  (optional)
    createdDateCode := "createdDateCode_example" // string |  (optional)
    createdBeforeDate := "createdBeforeDate_example" // string |  (optional)
    createdAfterDate := "createdAfterDate_example" // string |  (optional)
    completedBeforeDate := "completedBeforeDate_example" // string |  (optional)
    completedAfterDate := "completedAfterDate_example" // string |  (optional)
    callback := "callback_example" // string |  (optional)
    taskId := int32(56) // int32 |  (optional)
    parentTaskId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    offsetCount := int32(56) // int32 |  (optional)
    includeTaskId := int32(56) // int32 |  (optional)
    useStartDatesForTodaysTasks := true // bool |  (optional)
    useAllProjects := true // bool |  (optional)
    starredProjectsOnly := true // bool |  (optional)
    showDeleted := true // bool |  (optional)
    showCompletedLists := true // bool |  (optional)
    onlyUntaggedTasks := true // bool |  (optional)
    onlyTasksWithUnreadComments := true // bool |  (optional)
    onlyTasksWithTickets := true // bool |  (optional)
    onlyTasksWithEstimatedTime := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    nestSubTasks := true // bool |  (optional)
    matchAllTags := true // bool |  (optional)
    matchAllProjectTags := true // bool |  (optional)
    matchAllProjectTagIds := true // bool |  (optional)
    matchAllExcludedTags := true // bool |  (optional)
    isReportDownload := true // bool |  (optional)
    includeUntaggedTasks := true // bool |  (optional)
    includeToday := true // bool |  (optional)
    includeTeamUserIds := true // bool |  (optional)
    includeTasksWithoutDueDates := true // bool |  (optional)
    includeTasksWithCards := true // bool |  (optional)
    includeTasksFromDeletedLists := true // bool |  (optional)
    includeReminders := true // bool |  (optional)
    includeLoggedTime := true // bool |  (optional)
    includeCustomFields := true // bool |  (optional)
    includeCrmDealIds := true // bool |  (optional)
    includeCompletedTasks := true // bool |  (optional)
    includeCompletedSubtasks := true // bool |  (optional)
    includeCompanyUserIds := true // bool |  (optional)
    includeBlockedTasks := true // bool |  (optional)
    includeAssigneeTeams := true // bool |  (optional)
    includeAssigneeCompanies := true // bool |  (optional)
    includeArchivedProjects := true // bool |  (optional)
    ignoreStartDates := true // bool |  (optional)
    getSubTasks := true // bool |  (optional)
    getFiles := true // bool |  (optional)
    excludeAssigneeNotOnProjectTeams := true // bool |  (optional)
    allowTemplateTasks := true // bool |  (optional)
    tasklistId := []int32{int32(123)} // []int32 |  (optional)
    taskIdList := []int32{int32(123)} // []int32 |  (optional)
    tags := []string{"Inner_example"} // []string |  (optional)
    tagIds := []int32{int32(123)} // []int32 |  (optional)
    skipCrmDealIds := []int32{int32(123)} // []int32 |  (optional)
    responsiblePartyIds := []int32{int32(123)} // []int32 |  (optional)
    projectTagIds := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    projectIds := []int32{int32(123)} // []int32 |  (optional)
    projectId2 := []int32{int32(123)} // []int32 | project filters (optional)
    projectHealths := []int32{int32(123)} // []int32 |  (optional)
    projectHealth := []int32{int32(123)} // []int32 |  (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryId := []int32{int32(123)} // []int32 |  (optional)
    followerIds := []int32{int32(123)} // []int32 |  (optional)
    filterText := []string{"Inner_example"} // []string |  (optional)
    filterBoardColumnIds := []int32{int32(123)} // []int32 |  (optional)
    excludeTagIds := []int32{int32(123)} // []int32 |  (optional)
    crmDealIds := []int32{int32(123)} // []int32 |  (optional)
    creatorIds := []int32{int32(123)} // []int32 |  (optional)
    companyId := []int32{int32(123)} // []int32 |  (optional)
    assigneeTeamIds := []int32{int32(123)} // []int32 |  (optional)
    assigneeCompanyIds := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2ProjectsprojectIdTasks(context.Background(), projectId).UpdatedAfterDate(updatedAfterDate).Today(today).StartDate(startDate).SortOrder(sortOrder).Sort(sort).ReportFormat(reportFormat).ProjectStatus(projectStatus).Priority(priority).Include(include).Filter(filter).EndDate(endDate).DeletedAfterDate(deletedAfterDate).DateCode(dateCode).DataSet(dataSet).CreatedFilter(createdFilter).CreatedDateCode(createdDateCode).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TaskId(taskId).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).OffsetCount(offsetCount).IncludeTaskId(includeTaskId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyTasksWithUnreadComments(onlyTasksWithUnreadComments).OnlyTasksWithTickets(onlyTasksWithTickets).OnlyTasksWithEstimatedTime(onlyTasksWithEstimatedTime).OnlyStarredProjects(onlyStarredProjects).NestSubTasks(nestSubTasks).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTagIds(matchAllProjectTagIds).MatchAllExcludedTags(matchAllExcludedTags).IsReportDownload(isReportDownload).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksWithCards(includeTasksWithCards).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCustomFields(includeCustomFields).IncludeCrmDealIds(includeCrmDealIds).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompanyUserIds(includeCompanyUserIds).IncludeBlockedTasks(includeBlockedTasks).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).ExcludeAssigneeNotOnProjectTeams(excludeAssigneeNotOnProjectTeams).AllowTemplateTasks(allowTemplateTasks).TasklistId(tasklistId).TaskIdList(taskIdList).Tags(tags).TagIds(tagIds).SkipCrmDealIds(skipCrmDealIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId2(projectId2).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ProjectCategoryId(projectCategoryId).FollowerIds(followerIds).FilterText(filterText).FilterBoardColumnIds(filterBoardColumnIds).ExcludeTagIds(excludeTagIds).CrmDealIds(crmDealIds).CreatorIds(creatorIds).CompanyId(companyId).AssigneeTeamIds(assigneeTeamIds).AssigneeCompanyIds(assigneeCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2ProjectsprojectIdTasks``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETProjectsApiV2ProjectsprojectIdTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedAfterDate** | **string** |  | 
 **today** | **string** |  | 
 **startDate** | **string** |  | 
 **sortOrder** | **string** |  | 
 **sort** | **string** |  | 
 **reportFormat** | **string** |  | 
 **projectStatus** | **string** |  | 
 **priority** | **string** |  | 
 **include** | **string** |  | 
 **filter** | **string** |  | 
 **endDate** | **string** |  | 
 **deletedAfterDate** | **string** |  | 
 **dateCode** | **string** |  | 
 **dataSet** | **string** |  | 
 **createdFilter** | **string** |  | 
 **createdDateCode** | **string** |  | 
 **createdBeforeDate** | **string** |  | 
 **createdAfterDate** | **string** |  | 
 **completedBeforeDate** | **string** |  | 
 **completedAfterDate** | **string** |  | 
 **callback** | **string** |  | 
 **taskId** | **int32** |  | 
 **parentTaskId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **offsetCount** | **int32** |  | 
 **includeTaskId** | **int32** |  | 
 **useStartDatesForTodaysTasks** | **bool** |  | 
 **useAllProjects** | **bool** |  | 
 **starredProjectsOnly** | **bool** |  | 
 **showDeleted** | **bool** |  | 
 **showCompletedLists** | **bool** |  | 
 **onlyUntaggedTasks** | **bool** |  | 
 **onlyTasksWithUnreadComments** | **bool** |  | 
 **onlyTasksWithTickets** | **bool** |  | 
 **onlyTasksWithEstimatedTime** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **nestSubTasks** | **bool** |  | 
 **matchAllTags** | **bool** |  | 
 **matchAllProjectTags** | **bool** |  | 
 **matchAllProjectTagIds** | **bool** |  | 
 **matchAllExcludedTags** | **bool** |  | 
 **isReportDownload** | **bool** |  | 
 **includeUntaggedTasks** | **bool** |  | 
 **includeToday** | **bool** |  | 
 **includeTeamUserIds** | **bool** |  | 
 **includeTasksWithoutDueDates** | **bool** |  | 
 **includeTasksWithCards** | **bool** |  | 
 **includeTasksFromDeletedLists** | **bool** |  | 
 **includeReminders** | **bool** |  | 
 **includeLoggedTime** | **bool** |  | 
 **includeCustomFields** | **bool** |  | 
 **includeCrmDealIds** | **bool** |  | 
 **includeCompletedTasks** | **bool** |  | 
 **includeCompletedSubtasks** | **bool** |  | 
 **includeCompanyUserIds** | **bool** |  | 
 **includeBlockedTasks** | **bool** |  | 
 **includeAssigneeTeams** | **bool** |  | 
 **includeAssigneeCompanies** | **bool** |  | 
 **includeArchivedProjects** | **bool** |  | 
 **ignoreStartDates** | **bool** |  | 
 **getSubTasks** | **bool** |  | 
 **getFiles** | **bool** |  | 
 **excludeAssigneeNotOnProjectTeams** | **bool** |  | 
 **allowTemplateTasks** | **bool** |  | 
 **tasklistId** | **[]int32** |  | 
 **taskIdList** | **[]int32** |  | 
 **tags** | **[]string** |  | 
 **tagIds** | **[]int32** |  | 
 **skipCrmDealIds** | **[]int32** |  | 
 **responsiblePartyIds** | **[]int32** |  | 
 **projectTagIds** | **[]int32** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **projectIds** | **[]int32** |  | 
 **projectId2** | **[]int32** | project filters | 
 **projectHealths** | **[]int32** |  | 
 **projectHealth** | **[]int32** |  | 
 **projectCompanyIds** | **[]int32** |  | 
 **projectCategoryIds** | **[]int32** |  | 
 **projectCategoryId** | **[]int32** |  | 
 **followerIds** | **[]int32** |  | 
 **filterText** | **[]string** |  | 
 **filterBoardColumnIds** | **[]int32** |  | 
 **excludeTagIds** | **[]int32** |  | 
 **crmDealIds** | **[]int32** |  | 
 **creatorIds** | **[]int32** |  | 
 **companyId** | **[]int32** |  | 
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


## GETProjectsApiV2ProjectsprojectIdTime

> GETProjectsApiV2ProjectsprojectIdTime(ctx, projectId2).UpdatedAfterDate(updatedAfterDate).ToDate(toDate).SortOrder(sortOrder).SortBy(sortBy).Permission(permission).InvoicedType(invoicedType).FromDate(fromDate).BillableType(billableType).TicketId(ticketId).ProjectId(projectId).PageSize(pageSize).Page(page).InvoiceId(invoiceId).IncludeTimelogId(includeTimelogId).ShowDeleted(showDeleted).MatchAllTags(matchAllTags).IncludeTags(includeTags).GetTotals(getTotals).UserId(userId).TaskTagIds(taskTagIds).TagIds(tagIds).CompanyId(companyId).Execute()

APICall_GET_projects_id_time_entries

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
    projectId2 := int32(56) // int32 | 
    updatedAfterDate := "updatedAfterDate_example" // string |  (optional)
    toDate := "toDate_example" // string |  (optional)
    sortOrder := "sortOrder_example" // string |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    permission := "permission_example" // string |  (optional)
    invoicedType := "invoicedType_example" // string |  (optional)
    fromDate := "fromDate_example" // string |  (optional)
    billableType := "billableType_example" // string |  (optional)
    ticketId := int32(56) // int32 |  (optional)
    projectId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    invoiceId := int32(56) // int32 |  (optional)
    includeTimelogId := int32(56) // int32 |  (optional)
    showDeleted := true // bool |  (optional)
    matchAllTags := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    getTotals := true // bool |  (optional)
    userId := []int32{int32(123)} // []int32 |  (optional)
    taskTagIds := []int32{int32(123)} // []int32 |  (optional)
    tagIds := []int32{int32(123)} // []int32 |  (optional)
    companyId := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2ProjectsprojectIdTime(context.Background(), projectId2).UpdatedAfterDate(updatedAfterDate).ToDate(toDate).SortOrder(sortOrder).SortBy(sortBy).Permission(permission).InvoicedType(invoicedType).FromDate(fromDate).BillableType(billableType).TicketId(ticketId).ProjectId(projectId).PageSize(pageSize).Page(page).InvoiceId(invoiceId).IncludeTimelogId(includeTimelogId).ShowDeleted(showDeleted).MatchAllTags(matchAllTags).IncludeTags(includeTags).GetTotals(getTotals).UserId(userId).TaskTagIds(taskTagIds).TagIds(tagIds).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2ProjectsprojectIdTime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2ProjectsprojectIdTimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedAfterDate** | **string** |  | 
 **toDate** | **string** |  | 
 **sortOrder** | **string** |  | 
 **sortBy** | **string** |  | 
 **permission** | **string** |  | 
 **invoicedType** | **string** |  | 
 **fromDate** | **string** |  | 
 **billableType** | **string** |  | 
 **ticketId** | **int32** |  | 
 **projectId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **invoiceId** | **int32** |  | 
 **includeTimelogId** | **int32** |  | 
 **showDeleted** | **bool** |  | 
 **matchAllTags** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **getTotals** | **bool** |  | 
 **userId** | **[]int32** |  | 
 **taskTagIds** | **[]int32** |  | 
 **tagIds** | **[]int32** |  | 
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


## GETProjectsApiV2ProjectsprojectIdTimeEntries

> GETProjectsApiV2ProjectsprojectIdTimeEntries(ctx, projectId2).UpdatedAfterDate(updatedAfterDate).ToDate(toDate).SortOrder(sortOrder).SortBy(sortBy).Permission(permission).InvoicedType(invoicedType).FromDate(fromDate).BillableType(billableType).TicketId(ticketId).ProjectId(projectId).PageSize(pageSize).Page(page).InvoiceId(invoiceId).IncludeTimelogId(includeTimelogId).ShowDeleted(showDeleted).MatchAllTags(matchAllTags).IncludeTags(includeTags).GetTotals(getTotals).UserId(userId).TaskTagIds(taskTagIds).TagIds(tagIds).CompanyId(companyId).Execute()

APICall_GET_projects_id_time_entries

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
    projectId2 := int32(56) // int32 | 
    updatedAfterDate := "updatedAfterDate_example" // string |  (optional)
    toDate := "toDate_example" // string |  (optional)
    sortOrder := "sortOrder_example" // string |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    permission := "permission_example" // string |  (optional)
    invoicedType := "invoicedType_example" // string |  (optional)
    fromDate := "fromDate_example" // string |  (optional)
    billableType := "billableType_example" // string |  (optional)
    ticketId := int32(56) // int32 |  (optional)
    projectId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    invoiceId := int32(56) // int32 |  (optional)
    includeTimelogId := int32(56) // int32 |  (optional)
    showDeleted := true // bool |  (optional)
    matchAllTags := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    getTotals := true // bool |  (optional)
    userId := []int32{int32(123)} // []int32 |  (optional)
    taskTagIds := []int32{int32(123)} // []int32 |  (optional)
    tagIds := []int32{int32(123)} // []int32 |  (optional)
    companyId := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2ProjectsprojectIdTimeEntries(context.Background(), projectId2).UpdatedAfterDate(updatedAfterDate).ToDate(toDate).SortOrder(sortOrder).SortBy(sortBy).Permission(permission).InvoicedType(invoicedType).FromDate(fromDate).BillableType(billableType).TicketId(ticketId).ProjectId(projectId).PageSize(pageSize).Page(page).InvoiceId(invoiceId).IncludeTimelogId(includeTimelogId).ShowDeleted(showDeleted).MatchAllTags(matchAllTags).IncludeTags(includeTags).GetTotals(getTotals).UserId(userId).TaskTagIds(taskTagIds).TagIds(tagIds).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2ProjectsprojectIdTimeEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2ProjectsprojectIdTimeEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedAfterDate** | **string** |  | 
 **toDate** | **string** |  | 
 **sortOrder** | **string** |  | 
 **sortBy** | **string** |  | 
 **permission** | **string** |  | 
 **invoicedType** | **string** |  | 
 **fromDate** | **string** |  | 
 **billableType** | **string** |  | 
 **ticketId** | **int32** |  | 
 **projectId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **invoiceId** | **int32** |  | 
 **includeTimelogId** | **int32** |  | 
 **showDeleted** | **bool** |  | 
 **matchAllTags** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **getTotals** | **bool** |  | 
 **userId** | **[]int32** |  | 
 **taskTagIds** | **[]int32** |  | 
 **tagIds** | **[]int32** |  | 
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


## GETProjectsApiV2ProjectsprojectIdTimelogs

> GETProjectsApiV2ProjectsprojectIdTimelogs(ctx, projectId2).UpdatedAfterDate(updatedAfterDate).ToDate(toDate).SortOrder(sortOrder).SortBy(sortBy).Permission(permission).InvoicedType(invoicedType).FromDate(fromDate).BillableType(billableType).TicketId(ticketId).ProjectId(projectId).PageSize(pageSize).Page(page).InvoiceId(invoiceId).IncludeTimelogId(includeTimelogId).ShowDeleted(showDeleted).MatchAllTags(matchAllTags).IncludeTags(includeTags).GetTotals(getTotals).UserId(userId).TaskTagIds(taskTagIds).TagIds(tagIds).CompanyId(companyId).Execute()

APICall_GET_projects_id_time_entries

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
    projectId2 := int32(56) // int32 | 
    updatedAfterDate := "updatedAfterDate_example" // string |  (optional)
    toDate := "toDate_example" // string |  (optional)
    sortOrder := "sortOrder_example" // string |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    permission := "permission_example" // string |  (optional)
    invoicedType := "invoicedType_example" // string |  (optional)
    fromDate := "fromDate_example" // string |  (optional)
    billableType := "billableType_example" // string |  (optional)
    ticketId := int32(56) // int32 |  (optional)
    projectId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    invoiceId := int32(56) // int32 |  (optional)
    includeTimelogId := int32(56) // int32 |  (optional)
    showDeleted := true // bool |  (optional)
    matchAllTags := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    getTotals := true // bool |  (optional)
    userId := []int32{int32(123)} // []int32 |  (optional)
    taskTagIds := []int32{int32(123)} // []int32 |  (optional)
    tagIds := []int32{int32(123)} // []int32 |  (optional)
    companyId := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2ProjectsprojectIdTimelogs(context.Background(), projectId2).UpdatedAfterDate(updatedAfterDate).ToDate(toDate).SortOrder(sortOrder).SortBy(sortBy).Permission(permission).InvoicedType(invoicedType).FromDate(fromDate).BillableType(billableType).TicketId(ticketId).ProjectId(projectId).PageSize(pageSize).Page(page).InvoiceId(invoiceId).IncludeTimelogId(includeTimelogId).ShowDeleted(showDeleted).MatchAllTags(matchAllTags).IncludeTags(includeTags).GetTotals(getTotals).UserId(userId).TaskTagIds(taskTagIds).TagIds(tagIds).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2ProjectsprojectIdTimelogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2ProjectsprojectIdTimelogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedAfterDate** | **string** |  | 
 **toDate** | **string** |  | 
 **sortOrder** | **string** |  | 
 **sortBy** | **string** |  | 
 **permission** | **string** |  | 
 **invoicedType** | **string** |  | 
 **fromDate** | **string** |  | 
 **billableType** | **string** |  | 
 **ticketId** | **int32** |  | 
 **projectId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **invoiceId** | **int32** |  | 
 **includeTimelogId** | **int32** |  | 
 **showDeleted** | **bool** |  | 
 **matchAllTags** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **getTotals** | **bool** |  | 
 **userId** | **[]int32** |  | 
 **taskTagIds** | **[]int32** |  | 
 **tagIds** | **[]int32** |  | 
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


## GETProjectsApiV2ProjectsprojejctIdAssingLaterRoles

> GETProjectsApiV2ProjectsprojejctIdAssingLaterRoles(ctx, projejctId).ProjectId(projectId).Execute()

Get assign later roles

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
    projejctId := int32(56) // int32 | 
    projectId := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2ProjectsprojejctIdAssingLaterRoles(context.Background(), projejctId).ProjectId(projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2ProjectsprojejctIdAssingLaterRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projejctId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2ProjectsprojejctIdAssingLaterRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectId** | **int32** |  | 

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


## GETProjectsApiV2Quicksearch

> GETProjectsApiV2Quicksearch(ctx).SearchTerm(searchTerm).Dateupdated(dateupdated).Timeout(timeout).ProjectId(projectId).PageSize(pageSize).Page(page).DayLimit(dayLimit).UseSmartSearch(useSmartSearch).SkipES(skipES).SearchWithTags(searchWithTags).SearchTeams(searchTeams).SearchTasks(searchTasks).SearchTasklists(searchTasklists).SearchProjects(searchProjects).SearchPeople(searchPeople).SearchNotebooks(searchNotebooks).SearchMilestones(searchMilestones).SearchMessages(searchMessages).SearchLinks(searchLinks).SearchFiles(searchFiles).SearchCompanies(searchCompanies).SearchComments(searchComments).SearchArchivedMessages(searchArchivedMessages).SearchAllProjects(searchAllProjects).ReturnExtras(returnExtras).MatchAllTags(matchAllTags).IncludeTags(includeTags).IncludeCompletedItems(includeCompletedItems).IncludeCompanyInSearch(includeCompanyInSearch).IncludeArchivedProjects(includeArchivedProjects).EventsInUTC(eventsInUTC).BasicSearch(basicSearch).Tags(tags).TagIds(tagIds).Execute()

APICall_POST_quicksearch

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
    searchTerm := "searchTerm_example" // string |  (optional)
    dateupdated := "dateupdated_example" // string |  (optional)
    timeout := int32(56) // int32 |  (optional)
    projectId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    dayLimit := int32(56) // int32 |  (optional)
    useSmartSearch := true // bool |  (optional)
    skipES := true // bool |  (optional)
    searchWithTags := true // bool |  (optional)
    searchTeams := true // bool |  (optional)
    searchTasks := true // bool |  (optional)
    searchTasklists := true // bool |  (optional)
    searchProjects := true // bool |  (optional)
    searchPeople := true // bool |  (optional)
    searchNotebooks := true // bool |  (optional)
    searchMilestones := true // bool |  (optional)
    searchMessages := true // bool |  (optional)
    searchLinks := true // bool |  (optional)
    searchFiles := true // bool |  (optional)
    searchCompanies := true // bool |  (optional)
    searchComments := true // bool |  (optional)
    searchArchivedMessages := true // bool |  (optional)
    searchAllProjects := true // bool |  (optional)
    returnExtras := true // bool |  (optional)
    matchAllTags := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    includeCompletedItems := true // bool |  (optional)
    includeCompanyInSearch := true // bool |  (optional)
    includeArchivedProjects := true // bool |  (optional)
    eventsInUTC := true // bool |  (optional)
    basicSearch := true // bool |  (optional)
    tags := []string{"Inner_example"} // []string |  (optional)
    tagIds := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2Quicksearch(context.Background()).SearchTerm(searchTerm).Dateupdated(dateupdated).Timeout(timeout).ProjectId(projectId).PageSize(pageSize).Page(page).DayLimit(dayLimit).UseSmartSearch(useSmartSearch).SkipES(skipES).SearchWithTags(searchWithTags).SearchTeams(searchTeams).SearchTasks(searchTasks).SearchTasklists(searchTasklists).SearchProjects(searchProjects).SearchPeople(searchPeople).SearchNotebooks(searchNotebooks).SearchMilestones(searchMilestones).SearchMessages(searchMessages).SearchLinks(searchLinks).SearchFiles(searchFiles).SearchCompanies(searchCompanies).SearchComments(searchComments).SearchArchivedMessages(searchArchivedMessages).SearchAllProjects(searchAllProjects).ReturnExtras(returnExtras).MatchAllTags(matchAllTags).IncludeTags(includeTags).IncludeCompletedItems(includeCompletedItems).IncludeCompanyInSearch(includeCompanyInSearch).IncludeArchivedProjects(includeArchivedProjects).EventsInUTC(eventsInUTC).BasicSearch(basicSearch).Tags(tags).TagIds(tagIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2Quicksearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2QuicksearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchTerm** | **string** |  | 
 **dateupdated** | **string** |  | 
 **timeout** | **int32** |  | 
 **projectId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **dayLimit** | **int32** |  | 
 **useSmartSearch** | **bool** |  | 
 **skipES** | **bool** |  | 
 **searchWithTags** | **bool** |  | 
 **searchTeams** | **bool** |  | 
 **searchTasks** | **bool** |  | 
 **searchTasklists** | **bool** |  | 
 **searchProjects** | **bool** |  | 
 **searchPeople** | **bool** |  | 
 **searchNotebooks** | **bool** |  | 
 **searchMilestones** | **bool** |  | 
 **searchMessages** | **bool** |  | 
 **searchLinks** | **bool** |  | 
 **searchFiles** | **bool** |  | 
 **searchCompanies** | **bool** |  | 
 **searchComments** | **bool** |  | 
 **searchArchivedMessages** | **bool** |  | 
 **searchAllProjects** | **bool** |  | 
 **returnExtras** | **bool** |  | 
 **matchAllTags** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **includeCompletedItems** | **bool** |  | 
 **includeCompanyInSearch** | **bool** |  | 
 **includeArchivedProjects** | **bool** |  | 
 **eventsInUTC** | **bool** |  | 
 **basicSearch** | **bool** |  | 
 **tags** | **[]string** |  | 
 **tagIds** | **[]int32** |  | 

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


## GETProjectsApiV2Summary

> GETProjectsApiV2Summary(ctx).TimeRangeStart(timeRangeStart).TimeRangeEnd(timeRangeEnd).Status(status).Set(set).ProjectStatus(projectStatus).Get(get).EventsRange(eventsRange).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UnreadMessagesMineOnly(unreadMessagesMineOnly).UnreadCommentsMineOnly(unreadCommentsMineOnly).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).IncludeArchivedProjects(includeArchivedProjects).UserId(userId).TimeLoggedByUserId(timeLoggedByUserId).TaskAssigneeUserId(taskAssigneeUserId).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).MilestoneAssigneeUserId(milestoneAssigneeUserId).FilterTagIds(filterTagIds).CompanyId(companyId).CategoryId(categoryId).AssigneeUserId(assigneeUserId).Execute()

APICall_GET_summary

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
    timeRangeStart := "timeRangeStart_example" // string |  (optional)
    timeRangeEnd := "timeRangeEnd_example" // string |  (optional)
    status := "status_example" // string |  (optional)
    set := "set_example" // string |  (optional)
    projectStatus := "projectStatus_example" // string |  (optional)
    get := "get_example" // string |  (optional)
    eventsRange := int32(56) // int32 |  (optional)
    useStartDatesForTodaysTasks := true // bool |  (optional)
    unreadMessagesMineOnly := true // bool |  (optional)
    unreadCommentsMineOnly := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    matchAllTags := true // bool |  (optional)
    matchAllProjectTags := true // bool |  (optional)
    includeArchivedProjects := true // bool |  (optional)
    userId := []int32{int32(123)} // []int32 |  (optional)
    timeLoggedByUserId := []int32{int32(123)} // []int32 |  (optional)
    taskAssigneeUserId := []int32{int32(123)} // []int32 |  (optional)
    projectTagIds := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    projectIds := []int32{int32(123)} // []int32 |  (optional)
    projectId := []int32{int32(123)} // []int32 | project filters (optional)
    projectHealths := []int32{int32(123)} // []int32 |  (optional)
    projectHealth := []int32{int32(123)} // []int32 |  (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 |  (optional)
    milestoneAssigneeUserId := []int32{int32(123)} // []int32 |  (optional)
    filterTagIds := []int32{int32(123)} // []int32 |  (optional)
    companyId := []int32{int32(123)} // []int32 |  (optional)
    categoryId := []int32{int32(123)} // []int32 |  (optional)
    assigneeUserId := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2Summary(context.Background()).TimeRangeStart(timeRangeStart).TimeRangeEnd(timeRangeEnd).Status(status).Set(set).ProjectStatus(projectStatus).Get(get).EventsRange(eventsRange).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UnreadMessagesMineOnly(unreadMessagesMineOnly).UnreadCommentsMineOnly(unreadCommentsMineOnly).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).IncludeArchivedProjects(includeArchivedProjects).UserId(userId).TimeLoggedByUserId(timeLoggedByUserId).TaskAssigneeUserId(taskAssigneeUserId).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).MilestoneAssigneeUserId(milestoneAssigneeUserId).FilterTagIds(filterTagIds).CompanyId(companyId).CategoryId(categoryId).AssigneeUserId(assigneeUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2Summary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2SummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timeRangeStart** | **string** |  | 
 **timeRangeEnd** | **string** |  | 
 **status** | **string** |  | 
 **set** | **string** |  | 
 **projectStatus** | **string** |  | 
 **get** | **string** |  | 
 **eventsRange** | **int32** |  | 
 **useStartDatesForTodaysTasks** | **bool** |  | 
 **unreadMessagesMineOnly** | **bool** |  | 
 **unreadCommentsMineOnly** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **matchAllTags** | **bool** |  | 
 **matchAllProjectTags** | **bool** |  | 
 **includeArchivedProjects** | **bool** |  | 
 **userId** | **[]int32** |  | 
 **timeLoggedByUserId** | **[]int32** |  | 
 **taskAssigneeUserId** | **[]int32** |  | 
 **projectTagIds** | **[]int32** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **projectIds** | **[]int32** |  | 
 **projectId** | **[]int32** | project filters | 
 **projectHealths** | **[]int32** |  | 
 **projectHealth** | **[]int32** |  | 
 **projectCompanyIds** | **[]int32** |  | 
 **projectCategoryIds** | **[]int32** |  | 
 **milestoneAssigneeUserId** | **[]int32** |  | 
 **filterTagIds** | **[]int32** |  | 
 **companyId** | **[]int32** |  | 
 **categoryId** | **[]int32** |  | 
 **assigneeUserId** | **[]int32** |  | 

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


## GETProjectsApiV2TasklistsobjectIdProject

> GETProjectsApiV2TasklistsobjectIdProject(ctx, objectId2).UserId(userId).ObjectId(objectId).ShowDeleted(showDeleted).IncludeTags(includeTags).IncludePeople(includePeople).GetTruePermissions(getTruePermissions).GetPermissions(getPermissions).GetNotificationSettings(getNotificationSettings).GetEmoji(getEmoji).GetEmailAddress(getEmailAddress).GetDateInfo(getDateInfo).GetActivePages(getActivePages).FormatMarkdown(formatMarkdown).Execute()

Get the project for the tasklist

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
    objectId2 := int32(56) // int32 | 
    userId := int32(56) // int32 |  (optional)
    objectId := int32(56) // int32 |  (optional)
    showDeleted := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    includePeople := true // bool |  (optional)
    getTruePermissions := true // bool |  (optional)
    getPermissions := true // bool |  (optional)
    getNotificationSettings := true // bool |  (optional)
    getEmoji := true // bool |  (optional)
    getEmailAddress := true // bool |  (optional)
    getDateInfo := true // bool |  (optional)
    getActivePages := true // bool |  (optional)
    formatMarkdown := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2TasklistsobjectIdProject(context.Background(), objectId2).UserId(userId).ObjectId(objectId).ShowDeleted(showDeleted).IncludeTags(includeTags).IncludePeople(includePeople).GetTruePermissions(getTruePermissions).GetPermissions(getPermissions).GetNotificationSettings(getNotificationSettings).GetEmoji(getEmoji).GetEmailAddress(getEmailAddress).GetDateInfo(getDateInfo).GetActivePages(getActivePages).FormatMarkdown(formatMarkdown).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2TasklistsobjectIdProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2TasklistsobjectIdProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **int32** |  | 
 **objectId** | **int32** |  | 
 **showDeleted** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **includePeople** | **bool** |  | 
 **getTruePermissions** | **bool** |  | 
 **getPermissions** | **bool** |  | 
 **getNotificationSettings** | **bool** |  | 
 **getEmoji** | **bool** |  | 
 **getEmailAddress** | **bool** |  | 
 **getDateInfo** | **bool** |  | 
 **getActivePages** | **bool** |  | 
 **formatMarkdown** | **bool** |  | 

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


## GETProjectsApiV2TaskliststasklistIdTaskids

> GETProjectsApiV2TaskliststasklistIdTaskids(ctx, tasklistId).UpdatedAfterDate(updatedAfterDate).Today(today).StartDate(startDate).SortOrder(sortOrder).Sort(sort).ReportFormat(reportFormat).ProjectStatus(projectStatus).Priority(priority).Include(include).Filter(filter).EndDate(endDate).DeletedAfterDate(deletedAfterDate).DateCode(dateCode).DataSet(dataSet).CreatedFilter(createdFilter).CreatedDateCode(createdDateCode).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TaskId(taskId).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).OffsetCount(offsetCount).IncludeTaskId(includeTaskId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyTasksWithUnreadComments(onlyTasksWithUnreadComments).OnlyTasksWithTickets(onlyTasksWithTickets).OnlyTasksWithEstimatedTime(onlyTasksWithEstimatedTime).OnlyStarredProjects(onlyStarredProjects).NestSubTasks(nestSubTasks).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTagIds(matchAllProjectTagIds).MatchAllExcludedTags(matchAllExcludedTags).IsReportDownload(isReportDownload).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksWithCards(includeTasksWithCards).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCustomFields(includeCustomFields).IncludeCrmDealIds(includeCrmDealIds).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompanyUserIds(includeCompanyUserIds).IncludeBlockedTasks(includeBlockedTasks).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).ExcludeAssigneeNotOnProjectTeams(excludeAssigneeNotOnProjectTeams).AllowTemplateTasks(allowTemplateTasks).TasklistId2(tasklistId2).TaskIdList(taskIdList).Tags(tags).TagIds(tagIds).SkipCrmDealIds(skipCrmDealIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ProjectCategoryId(projectCategoryId).FollowerIds(followerIds).FilterText(filterText).FilterBoardColumnIds(filterBoardColumnIds).ExcludeTagIds(excludeTagIds).CrmDealIds(crmDealIds).CreatorIds(creatorIds).CompanyId(companyId).AssigneeTeamIds(assigneeTeamIds).AssigneeCompanyIds(assigneeCompanyIds).Execute()

Get Tasks

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
    startDate := "startDate_example" // string |  (optional)
    sortOrder := "sortOrder_example" // string |  (optional)
    sort := "sort_example" // string |  (optional)
    reportFormat := "reportFormat_example" // string |  (optional)
    projectStatus := "projectStatus_example" // string |  (optional)
    priority := "priority_example" // string |  (optional)
    include := "include_example" // string |  (optional)
    filter := "filter_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    deletedAfterDate := "deletedAfterDate_example" // string |  (optional)
    dateCode := "dateCode_example" // string |  (optional)
    dataSet := "dataSet_example" // string |  (optional)
    createdFilter := "createdFilter_example" // string |  (optional)
    createdDateCode := "createdDateCode_example" // string |  (optional)
    createdBeforeDate := "createdBeforeDate_example" // string |  (optional)
    createdAfterDate := "createdAfterDate_example" // string |  (optional)
    completedBeforeDate := "completedBeforeDate_example" // string |  (optional)
    completedAfterDate := "completedAfterDate_example" // string |  (optional)
    callback := "callback_example" // string |  (optional)
    taskId := int32(56) // int32 |  (optional)
    parentTaskId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    offsetCount := int32(56) // int32 |  (optional)
    includeTaskId := int32(56) // int32 |  (optional)
    useStartDatesForTodaysTasks := true // bool |  (optional)
    useAllProjects := true // bool |  (optional)
    starredProjectsOnly := true // bool |  (optional)
    showDeleted := true // bool |  (optional)
    showCompletedLists := true // bool |  (optional)
    onlyUntaggedTasks := true // bool |  (optional)
    onlyTasksWithUnreadComments := true // bool |  (optional)
    onlyTasksWithTickets := true // bool |  (optional)
    onlyTasksWithEstimatedTime := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    nestSubTasks := true // bool |  (optional)
    matchAllTags := true // bool |  (optional)
    matchAllProjectTags := true // bool |  (optional)
    matchAllProjectTagIds := true // bool |  (optional)
    matchAllExcludedTags := true // bool |  (optional)
    isReportDownload := true // bool |  (optional)
    includeUntaggedTasks := true // bool |  (optional)
    includeToday := true // bool |  (optional)
    includeTeamUserIds := true // bool |  (optional)
    includeTasksWithoutDueDates := true // bool |  (optional)
    includeTasksWithCards := true // bool |  (optional)
    includeTasksFromDeletedLists := true // bool |  (optional)
    includeReminders := true // bool |  (optional)
    includeLoggedTime := true // bool |  (optional)
    includeCustomFields := true // bool |  (optional)
    includeCrmDealIds := true // bool |  (optional)
    includeCompletedTasks := true // bool |  (optional)
    includeCompletedSubtasks := true // bool |  (optional)
    includeCompanyUserIds := true // bool |  (optional)
    includeBlockedTasks := true // bool |  (optional)
    includeAssigneeTeams := true // bool |  (optional)
    includeAssigneeCompanies := true // bool |  (optional)
    includeArchivedProjects := true // bool |  (optional)
    ignoreStartDates := true // bool |  (optional)
    getSubTasks := true // bool |  (optional)
    getFiles := true // bool |  (optional)
    excludeAssigneeNotOnProjectTeams := true // bool |  (optional)
    allowTemplateTasks := true // bool |  (optional)
    tasklistId2 := []int32{int32(123)} // []int32 |  (optional)
    taskIdList := []int32{int32(123)} // []int32 |  (optional)
    tags := []string{"Inner_example"} // []string |  (optional)
    tagIds := []int32{int32(123)} // []int32 |  (optional)
    skipCrmDealIds := []int32{int32(123)} // []int32 |  (optional)
    responsiblePartyIds := []int32{int32(123)} // []int32 |  (optional)
    projectTagIds := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    projectIds := []int32{int32(123)} // []int32 |  (optional)
    projectId := []int32{int32(123)} // []int32 | project filters (optional)
    projectHealths := []int32{int32(123)} // []int32 |  (optional)
    projectHealth := []int32{int32(123)} // []int32 |  (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryId := []int32{int32(123)} // []int32 |  (optional)
    followerIds := []int32{int32(123)} // []int32 |  (optional)
    filterText := []string{"Inner_example"} // []string |  (optional)
    filterBoardColumnIds := []int32{int32(123)} // []int32 |  (optional)
    excludeTagIds := []int32{int32(123)} // []int32 |  (optional)
    crmDealIds := []int32{int32(123)} // []int32 |  (optional)
    creatorIds := []int32{int32(123)} // []int32 |  (optional)
    companyId := []int32{int32(123)} // []int32 |  (optional)
    assigneeTeamIds := []int32{int32(123)} // []int32 |  (optional)
    assigneeCompanyIds := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2TaskliststasklistIdTaskids(context.Background(), tasklistId).UpdatedAfterDate(updatedAfterDate).Today(today).StartDate(startDate).SortOrder(sortOrder).Sort(sort).ReportFormat(reportFormat).ProjectStatus(projectStatus).Priority(priority).Include(include).Filter(filter).EndDate(endDate).DeletedAfterDate(deletedAfterDate).DateCode(dateCode).DataSet(dataSet).CreatedFilter(createdFilter).CreatedDateCode(createdDateCode).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TaskId(taskId).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).OffsetCount(offsetCount).IncludeTaskId(includeTaskId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyTasksWithUnreadComments(onlyTasksWithUnreadComments).OnlyTasksWithTickets(onlyTasksWithTickets).OnlyTasksWithEstimatedTime(onlyTasksWithEstimatedTime).OnlyStarredProjects(onlyStarredProjects).NestSubTasks(nestSubTasks).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTagIds(matchAllProjectTagIds).MatchAllExcludedTags(matchAllExcludedTags).IsReportDownload(isReportDownload).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksWithCards(includeTasksWithCards).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCustomFields(includeCustomFields).IncludeCrmDealIds(includeCrmDealIds).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompanyUserIds(includeCompanyUserIds).IncludeBlockedTasks(includeBlockedTasks).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).ExcludeAssigneeNotOnProjectTeams(excludeAssigneeNotOnProjectTeams).AllowTemplateTasks(allowTemplateTasks).TasklistId2(tasklistId2).TaskIdList(taskIdList).Tags(tags).TagIds(tagIds).SkipCrmDealIds(skipCrmDealIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ProjectCategoryId(projectCategoryId).FollowerIds(followerIds).FilterText(filterText).FilterBoardColumnIds(filterBoardColumnIds).ExcludeTagIds(excludeTagIds).CrmDealIds(crmDealIds).CreatorIds(creatorIds).CompanyId(companyId).AssigneeTeamIds(assigneeTeamIds).AssigneeCompanyIds(assigneeCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2TaskliststasklistIdTaskids``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETProjectsApiV2TaskliststasklistIdTaskidsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedAfterDate** | **string** |  | 
 **today** | **string** |  | 
 **startDate** | **string** |  | 
 **sortOrder** | **string** |  | 
 **sort** | **string** |  | 
 **reportFormat** | **string** |  | 
 **projectStatus** | **string** |  | 
 **priority** | **string** |  | 
 **include** | **string** |  | 
 **filter** | **string** |  | 
 **endDate** | **string** |  | 
 **deletedAfterDate** | **string** |  | 
 **dateCode** | **string** |  | 
 **dataSet** | **string** |  | 
 **createdFilter** | **string** |  | 
 **createdDateCode** | **string** |  | 
 **createdBeforeDate** | **string** |  | 
 **createdAfterDate** | **string** |  | 
 **completedBeforeDate** | **string** |  | 
 **completedAfterDate** | **string** |  | 
 **callback** | **string** |  | 
 **taskId** | **int32** |  | 
 **parentTaskId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **offsetCount** | **int32** |  | 
 **includeTaskId** | **int32** |  | 
 **useStartDatesForTodaysTasks** | **bool** |  | 
 **useAllProjects** | **bool** |  | 
 **starredProjectsOnly** | **bool** |  | 
 **showDeleted** | **bool** |  | 
 **showCompletedLists** | **bool** |  | 
 **onlyUntaggedTasks** | **bool** |  | 
 **onlyTasksWithUnreadComments** | **bool** |  | 
 **onlyTasksWithTickets** | **bool** |  | 
 **onlyTasksWithEstimatedTime** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **nestSubTasks** | **bool** |  | 
 **matchAllTags** | **bool** |  | 
 **matchAllProjectTags** | **bool** |  | 
 **matchAllProjectTagIds** | **bool** |  | 
 **matchAllExcludedTags** | **bool** |  | 
 **isReportDownload** | **bool** |  | 
 **includeUntaggedTasks** | **bool** |  | 
 **includeToday** | **bool** |  | 
 **includeTeamUserIds** | **bool** |  | 
 **includeTasksWithoutDueDates** | **bool** |  | 
 **includeTasksWithCards** | **bool** |  | 
 **includeTasksFromDeletedLists** | **bool** |  | 
 **includeReminders** | **bool** |  | 
 **includeLoggedTime** | **bool** |  | 
 **includeCustomFields** | **bool** |  | 
 **includeCrmDealIds** | **bool** |  | 
 **includeCompletedTasks** | **bool** |  | 
 **includeCompletedSubtasks** | **bool** |  | 
 **includeCompanyUserIds** | **bool** |  | 
 **includeBlockedTasks** | **bool** |  | 
 **includeAssigneeTeams** | **bool** |  | 
 **includeAssigneeCompanies** | **bool** |  | 
 **includeArchivedProjects** | **bool** |  | 
 **ignoreStartDates** | **bool** |  | 
 **getSubTasks** | **bool** |  | 
 **getFiles** | **bool** |  | 
 **excludeAssigneeNotOnProjectTeams** | **bool** |  | 
 **allowTemplateTasks** | **bool** |  | 
 **tasklistId2** | **[]int32** |  | 
 **taskIdList** | **[]int32** |  | 
 **tags** | **[]string** |  | 
 **tagIds** | **[]int32** |  | 
 **skipCrmDealIds** | **[]int32** |  | 
 **responsiblePartyIds** | **[]int32** |  | 
 **projectTagIds** | **[]int32** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **projectIds** | **[]int32** |  | 
 **projectId** | **[]int32** | project filters | 
 **projectHealths** | **[]int32** |  | 
 **projectHealth** | **[]int32** |  | 
 **projectCompanyIds** | **[]int32** |  | 
 **projectCategoryIds** | **[]int32** |  | 
 **projectCategoryId** | **[]int32** |  | 
 **followerIds** | **[]int32** |  | 
 **filterText** | **[]string** |  | 
 **filterBoardColumnIds** | **[]int32** |  | 
 **excludeTagIds** | **[]int32** |  | 
 **crmDealIds** | **[]int32** |  | 
 **creatorIds** | **[]int32** |  | 
 **companyId** | **[]int32** |  | 
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


## GETProjectsApiV2TaskliststasklistIdTasks

> GETProjectsApiV2TaskliststasklistIdTasks(ctx, tasklistId).UpdatedAfterDate(updatedAfterDate).Today(today).StartDate(startDate).SortOrder(sortOrder).Sort(sort).ReportFormat(reportFormat).ProjectStatus(projectStatus).Priority(priority).Include(include).Filter(filter).EndDate(endDate).DeletedAfterDate(deletedAfterDate).DateCode(dateCode).DataSet(dataSet).CreatedFilter(createdFilter).CreatedDateCode(createdDateCode).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TaskId(taskId).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).OffsetCount(offsetCount).IncludeTaskId(includeTaskId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyTasksWithUnreadComments(onlyTasksWithUnreadComments).OnlyTasksWithTickets(onlyTasksWithTickets).OnlyTasksWithEstimatedTime(onlyTasksWithEstimatedTime).OnlyStarredProjects(onlyStarredProjects).NestSubTasks(nestSubTasks).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTagIds(matchAllProjectTagIds).MatchAllExcludedTags(matchAllExcludedTags).IsReportDownload(isReportDownload).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksWithCards(includeTasksWithCards).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCustomFields(includeCustomFields).IncludeCrmDealIds(includeCrmDealIds).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompanyUserIds(includeCompanyUserIds).IncludeBlockedTasks(includeBlockedTasks).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).ExcludeAssigneeNotOnProjectTeams(excludeAssigneeNotOnProjectTeams).AllowTemplateTasks(allowTemplateTasks).TasklistId2(tasklistId2).TaskIdList(taskIdList).Tags(tags).TagIds(tagIds).SkipCrmDealIds(skipCrmDealIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ProjectCategoryId(projectCategoryId).FollowerIds(followerIds).FilterText(filterText).FilterBoardColumnIds(filterBoardColumnIds).ExcludeTagIds(excludeTagIds).CrmDealIds(crmDealIds).CreatorIds(creatorIds).CompanyId(companyId).AssigneeTeamIds(assigneeTeamIds).AssigneeCompanyIds(assigneeCompanyIds).Execute()

Get Tasks

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
    startDate := "startDate_example" // string |  (optional)
    sortOrder := "sortOrder_example" // string |  (optional)
    sort := "sort_example" // string |  (optional)
    reportFormat := "reportFormat_example" // string |  (optional)
    projectStatus := "projectStatus_example" // string |  (optional)
    priority := "priority_example" // string |  (optional)
    include := "include_example" // string |  (optional)
    filter := "filter_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    deletedAfterDate := "deletedAfterDate_example" // string |  (optional)
    dateCode := "dateCode_example" // string |  (optional)
    dataSet := "dataSet_example" // string |  (optional)
    createdFilter := "createdFilter_example" // string |  (optional)
    createdDateCode := "createdDateCode_example" // string |  (optional)
    createdBeforeDate := "createdBeforeDate_example" // string |  (optional)
    createdAfterDate := "createdAfterDate_example" // string |  (optional)
    completedBeforeDate := "completedBeforeDate_example" // string |  (optional)
    completedAfterDate := "completedAfterDate_example" // string |  (optional)
    callback := "callback_example" // string |  (optional)
    taskId := int32(56) // int32 |  (optional)
    parentTaskId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    offsetCount := int32(56) // int32 |  (optional)
    includeTaskId := int32(56) // int32 |  (optional)
    useStartDatesForTodaysTasks := true // bool |  (optional)
    useAllProjects := true // bool |  (optional)
    starredProjectsOnly := true // bool |  (optional)
    showDeleted := true // bool |  (optional)
    showCompletedLists := true // bool |  (optional)
    onlyUntaggedTasks := true // bool |  (optional)
    onlyTasksWithUnreadComments := true // bool |  (optional)
    onlyTasksWithTickets := true // bool |  (optional)
    onlyTasksWithEstimatedTime := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    nestSubTasks := true // bool |  (optional)
    matchAllTags := true // bool |  (optional)
    matchAllProjectTags := true // bool |  (optional)
    matchAllProjectTagIds := true // bool |  (optional)
    matchAllExcludedTags := true // bool |  (optional)
    isReportDownload := true // bool |  (optional)
    includeUntaggedTasks := true // bool |  (optional)
    includeToday := true // bool |  (optional)
    includeTeamUserIds := true // bool |  (optional)
    includeTasksWithoutDueDates := true // bool |  (optional)
    includeTasksWithCards := true // bool |  (optional)
    includeTasksFromDeletedLists := true // bool |  (optional)
    includeReminders := true // bool |  (optional)
    includeLoggedTime := true // bool |  (optional)
    includeCustomFields := true // bool |  (optional)
    includeCrmDealIds := true // bool |  (optional)
    includeCompletedTasks := true // bool |  (optional)
    includeCompletedSubtasks := true // bool |  (optional)
    includeCompanyUserIds := true // bool |  (optional)
    includeBlockedTasks := true // bool |  (optional)
    includeAssigneeTeams := true // bool |  (optional)
    includeAssigneeCompanies := true // bool |  (optional)
    includeArchivedProjects := true // bool |  (optional)
    ignoreStartDates := true // bool |  (optional)
    getSubTasks := true // bool |  (optional)
    getFiles := true // bool |  (optional)
    excludeAssigneeNotOnProjectTeams := true // bool |  (optional)
    allowTemplateTasks := true // bool |  (optional)
    tasklistId2 := []int32{int32(123)} // []int32 |  (optional)
    taskIdList := []int32{int32(123)} // []int32 |  (optional)
    tags := []string{"Inner_example"} // []string |  (optional)
    tagIds := []int32{int32(123)} // []int32 |  (optional)
    skipCrmDealIds := []int32{int32(123)} // []int32 |  (optional)
    responsiblePartyIds := []int32{int32(123)} // []int32 |  (optional)
    projectTagIds := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    projectIds := []int32{int32(123)} // []int32 |  (optional)
    projectId := []int32{int32(123)} // []int32 | project filters (optional)
    projectHealths := []int32{int32(123)} // []int32 |  (optional)
    projectHealth := []int32{int32(123)} // []int32 |  (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryId := []int32{int32(123)} // []int32 |  (optional)
    followerIds := []int32{int32(123)} // []int32 |  (optional)
    filterText := []string{"Inner_example"} // []string |  (optional)
    filterBoardColumnIds := []int32{int32(123)} // []int32 |  (optional)
    excludeTagIds := []int32{int32(123)} // []int32 |  (optional)
    crmDealIds := []int32{int32(123)} // []int32 |  (optional)
    creatorIds := []int32{int32(123)} // []int32 |  (optional)
    companyId := []int32{int32(123)} // []int32 |  (optional)
    assigneeTeamIds := []int32{int32(123)} // []int32 |  (optional)
    assigneeCompanyIds := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2TaskliststasklistIdTasks(context.Background(), tasklistId).UpdatedAfterDate(updatedAfterDate).Today(today).StartDate(startDate).SortOrder(sortOrder).Sort(sort).ReportFormat(reportFormat).ProjectStatus(projectStatus).Priority(priority).Include(include).Filter(filter).EndDate(endDate).DeletedAfterDate(deletedAfterDate).DateCode(dateCode).DataSet(dataSet).CreatedFilter(createdFilter).CreatedDateCode(createdDateCode).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TaskId(taskId).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).OffsetCount(offsetCount).IncludeTaskId(includeTaskId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyTasksWithUnreadComments(onlyTasksWithUnreadComments).OnlyTasksWithTickets(onlyTasksWithTickets).OnlyTasksWithEstimatedTime(onlyTasksWithEstimatedTime).OnlyStarredProjects(onlyStarredProjects).NestSubTasks(nestSubTasks).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTagIds(matchAllProjectTagIds).MatchAllExcludedTags(matchAllExcludedTags).IsReportDownload(isReportDownload).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksWithCards(includeTasksWithCards).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCustomFields(includeCustomFields).IncludeCrmDealIds(includeCrmDealIds).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompanyUserIds(includeCompanyUserIds).IncludeBlockedTasks(includeBlockedTasks).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).ExcludeAssigneeNotOnProjectTeams(excludeAssigneeNotOnProjectTeams).AllowTemplateTasks(allowTemplateTasks).TasklistId2(tasklistId2).TaskIdList(taskIdList).Tags(tags).TagIds(tagIds).SkipCrmDealIds(skipCrmDealIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ProjectCategoryId(projectCategoryId).FollowerIds(followerIds).FilterText(filterText).FilterBoardColumnIds(filterBoardColumnIds).ExcludeTagIds(excludeTagIds).CrmDealIds(crmDealIds).CreatorIds(creatorIds).CompanyId(companyId).AssigneeTeamIds(assigneeTeamIds).AssigneeCompanyIds(assigneeCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2TaskliststasklistIdTasks``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETProjectsApiV2TaskliststasklistIdTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedAfterDate** | **string** |  | 
 **today** | **string** |  | 
 **startDate** | **string** |  | 
 **sortOrder** | **string** |  | 
 **sort** | **string** |  | 
 **reportFormat** | **string** |  | 
 **projectStatus** | **string** |  | 
 **priority** | **string** |  | 
 **include** | **string** |  | 
 **filter** | **string** |  | 
 **endDate** | **string** |  | 
 **deletedAfterDate** | **string** |  | 
 **dateCode** | **string** |  | 
 **dataSet** | **string** |  | 
 **createdFilter** | **string** |  | 
 **createdDateCode** | **string** |  | 
 **createdBeforeDate** | **string** |  | 
 **createdAfterDate** | **string** |  | 
 **completedBeforeDate** | **string** |  | 
 **completedAfterDate** | **string** |  | 
 **callback** | **string** |  | 
 **taskId** | **int32** |  | 
 **parentTaskId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **offsetCount** | **int32** |  | 
 **includeTaskId** | **int32** |  | 
 **useStartDatesForTodaysTasks** | **bool** |  | 
 **useAllProjects** | **bool** |  | 
 **starredProjectsOnly** | **bool** |  | 
 **showDeleted** | **bool** |  | 
 **showCompletedLists** | **bool** |  | 
 **onlyUntaggedTasks** | **bool** |  | 
 **onlyTasksWithUnreadComments** | **bool** |  | 
 **onlyTasksWithTickets** | **bool** |  | 
 **onlyTasksWithEstimatedTime** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **nestSubTasks** | **bool** |  | 
 **matchAllTags** | **bool** |  | 
 **matchAllProjectTags** | **bool** |  | 
 **matchAllProjectTagIds** | **bool** |  | 
 **matchAllExcludedTags** | **bool** |  | 
 **isReportDownload** | **bool** |  | 
 **includeUntaggedTasks** | **bool** |  | 
 **includeToday** | **bool** |  | 
 **includeTeamUserIds** | **bool** |  | 
 **includeTasksWithoutDueDates** | **bool** |  | 
 **includeTasksWithCards** | **bool** |  | 
 **includeTasksFromDeletedLists** | **bool** |  | 
 **includeReminders** | **bool** |  | 
 **includeLoggedTime** | **bool** |  | 
 **includeCustomFields** | **bool** |  | 
 **includeCrmDealIds** | **bool** |  | 
 **includeCompletedTasks** | **bool** |  | 
 **includeCompletedSubtasks** | **bool** |  | 
 **includeCompanyUserIds** | **bool** |  | 
 **includeBlockedTasks** | **bool** |  | 
 **includeAssigneeTeams** | **bool** |  | 
 **includeAssigneeCompanies** | **bool** |  | 
 **includeArchivedProjects** | **bool** |  | 
 **ignoreStartDates** | **bool** |  | 
 **getSubTasks** | **bool** |  | 
 **getFiles** | **bool** |  | 
 **excludeAssigneeNotOnProjectTeams** | **bool** |  | 
 **allowTemplateTasks** | **bool** |  | 
 **tasklistId2** | **[]int32** |  | 
 **taskIdList** | **[]int32** |  | 
 **tags** | **[]string** |  | 
 **tagIds** | **[]int32** |  | 
 **skipCrmDealIds** | **[]int32** |  | 
 **responsiblePartyIds** | **[]int32** |  | 
 **projectTagIds** | **[]int32** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **projectIds** | **[]int32** |  | 
 **projectId** | **[]int32** | project filters | 
 **projectHealths** | **[]int32** |  | 
 **projectHealth** | **[]int32** |  | 
 **projectCompanyIds** | **[]int32** |  | 
 **projectCategoryIds** | **[]int32** |  | 
 **projectCategoryId** | **[]int32** |  | 
 **followerIds** | **[]int32** |  | 
 **filterText** | **[]string** |  | 
 **filterBoardColumnIds** | **[]int32** |  | 
 **excludeTagIds** | **[]int32** |  | 
 **crmDealIds** | **[]int32** |  | 
 **creatorIds** | **[]int32** |  | 
 **companyId** | **[]int32** |  | 
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


## GETProjectsApiV2Tasks

> GETProjectsApiV2Tasks(ctx).UpdatedAfterDate(updatedAfterDate).Today(today).StartDate(startDate).SortOrder(sortOrder).Sort(sort).ReportFormat(reportFormat).ProjectStatus(projectStatus).Priority(priority).Include(include).Filter(filter).EndDate(endDate).DeletedAfterDate(deletedAfterDate).DateCode(dateCode).DataSet(dataSet).CreatedFilter(createdFilter).CreatedDateCode(createdDateCode).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TaskId(taskId).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).OffsetCount(offsetCount).IncludeTaskId(includeTaskId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyTasksWithUnreadComments(onlyTasksWithUnreadComments).OnlyTasksWithTickets(onlyTasksWithTickets).OnlyTasksWithEstimatedTime(onlyTasksWithEstimatedTime).OnlyStarredProjects(onlyStarredProjects).NestSubTasks(nestSubTasks).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTagIds(matchAllProjectTagIds).MatchAllExcludedTags(matchAllExcludedTags).IsReportDownload(isReportDownload).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksWithCards(includeTasksWithCards).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCustomFields(includeCustomFields).IncludeCrmDealIds(includeCrmDealIds).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompanyUserIds(includeCompanyUserIds).IncludeBlockedTasks(includeBlockedTasks).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).ExcludeAssigneeNotOnProjectTeams(excludeAssigneeNotOnProjectTeams).AllowTemplateTasks(allowTemplateTasks).TasklistId(tasklistId).TaskIdList(taskIdList).Tags(tags).TagIds(tagIds).SkipCrmDealIds(skipCrmDealIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ProjectCategoryId(projectCategoryId).FollowerIds(followerIds).FilterText(filterText).FilterBoardColumnIds(filterBoardColumnIds).ExcludeTagIds(excludeTagIds).CrmDealIds(crmDealIds).CreatorIds(creatorIds).CompanyId(companyId).AssigneeTeamIds(assigneeTeamIds).AssigneeCompanyIds(assigneeCompanyIds).Execute()

Get Tasks

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
    startDate := "startDate_example" // string |  (optional)
    sortOrder := "sortOrder_example" // string |  (optional)
    sort := "sort_example" // string |  (optional)
    reportFormat := "reportFormat_example" // string |  (optional)
    projectStatus := "projectStatus_example" // string |  (optional)
    priority := "priority_example" // string |  (optional)
    include := "include_example" // string |  (optional)
    filter := "filter_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    deletedAfterDate := "deletedAfterDate_example" // string |  (optional)
    dateCode := "dateCode_example" // string |  (optional)
    dataSet := "dataSet_example" // string |  (optional)
    createdFilter := "createdFilter_example" // string |  (optional)
    createdDateCode := "createdDateCode_example" // string |  (optional)
    createdBeforeDate := "createdBeforeDate_example" // string |  (optional)
    createdAfterDate := "createdAfterDate_example" // string |  (optional)
    completedBeforeDate := "completedBeforeDate_example" // string |  (optional)
    completedAfterDate := "completedAfterDate_example" // string |  (optional)
    callback := "callback_example" // string |  (optional)
    taskId := int32(56) // int32 |  (optional)
    parentTaskId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    offsetCount := int32(56) // int32 |  (optional)
    includeTaskId := int32(56) // int32 |  (optional)
    useStartDatesForTodaysTasks := true // bool |  (optional)
    useAllProjects := true // bool |  (optional)
    starredProjectsOnly := true // bool |  (optional)
    showDeleted := true // bool |  (optional)
    showCompletedLists := true // bool |  (optional)
    onlyUntaggedTasks := true // bool |  (optional)
    onlyTasksWithUnreadComments := true // bool |  (optional)
    onlyTasksWithTickets := true // bool |  (optional)
    onlyTasksWithEstimatedTime := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    nestSubTasks := true // bool |  (optional)
    matchAllTags := true // bool |  (optional)
    matchAllProjectTags := true // bool |  (optional)
    matchAllProjectTagIds := true // bool |  (optional)
    matchAllExcludedTags := true // bool |  (optional)
    isReportDownload := true // bool |  (optional)
    includeUntaggedTasks := true // bool |  (optional)
    includeToday := true // bool |  (optional)
    includeTeamUserIds := true // bool |  (optional)
    includeTasksWithoutDueDates := true // bool |  (optional)
    includeTasksWithCards := true // bool |  (optional)
    includeTasksFromDeletedLists := true // bool |  (optional)
    includeReminders := true // bool |  (optional)
    includeLoggedTime := true // bool |  (optional)
    includeCustomFields := true // bool |  (optional)
    includeCrmDealIds := true // bool |  (optional)
    includeCompletedTasks := true // bool |  (optional)
    includeCompletedSubtasks := true // bool |  (optional)
    includeCompanyUserIds := true // bool |  (optional)
    includeBlockedTasks := true // bool |  (optional)
    includeAssigneeTeams := true // bool |  (optional)
    includeAssigneeCompanies := true // bool |  (optional)
    includeArchivedProjects := true // bool |  (optional)
    ignoreStartDates := true // bool |  (optional)
    getSubTasks := true // bool |  (optional)
    getFiles := true // bool |  (optional)
    excludeAssigneeNotOnProjectTeams := true // bool |  (optional)
    allowTemplateTasks := true // bool |  (optional)
    tasklistId := []int32{int32(123)} // []int32 |  (optional)
    taskIdList := []int32{int32(123)} // []int32 |  (optional)
    tags := []string{"Inner_example"} // []string |  (optional)
    tagIds := []int32{int32(123)} // []int32 |  (optional)
    skipCrmDealIds := []int32{int32(123)} // []int32 |  (optional)
    responsiblePartyIds := []int32{int32(123)} // []int32 |  (optional)
    projectTagIds := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    projectIds := []int32{int32(123)} // []int32 |  (optional)
    projectId := []int32{int32(123)} // []int32 | project filters (optional)
    projectHealths := []int32{int32(123)} // []int32 |  (optional)
    projectHealth := []int32{int32(123)} // []int32 |  (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryId := []int32{int32(123)} // []int32 |  (optional)
    followerIds := []int32{int32(123)} // []int32 |  (optional)
    filterText := []string{"Inner_example"} // []string |  (optional)
    filterBoardColumnIds := []int32{int32(123)} // []int32 |  (optional)
    excludeTagIds := []int32{int32(123)} // []int32 |  (optional)
    crmDealIds := []int32{int32(123)} // []int32 |  (optional)
    creatorIds := []int32{int32(123)} // []int32 |  (optional)
    companyId := []int32{int32(123)} // []int32 |  (optional)
    assigneeTeamIds := []int32{int32(123)} // []int32 |  (optional)
    assigneeCompanyIds := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2Tasks(context.Background()).UpdatedAfterDate(updatedAfterDate).Today(today).StartDate(startDate).SortOrder(sortOrder).Sort(sort).ReportFormat(reportFormat).ProjectStatus(projectStatus).Priority(priority).Include(include).Filter(filter).EndDate(endDate).DeletedAfterDate(deletedAfterDate).DateCode(dateCode).DataSet(dataSet).CreatedFilter(createdFilter).CreatedDateCode(createdDateCode).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TaskId(taskId).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).OffsetCount(offsetCount).IncludeTaskId(includeTaskId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyTasksWithUnreadComments(onlyTasksWithUnreadComments).OnlyTasksWithTickets(onlyTasksWithTickets).OnlyTasksWithEstimatedTime(onlyTasksWithEstimatedTime).OnlyStarredProjects(onlyStarredProjects).NestSubTasks(nestSubTasks).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTagIds(matchAllProjectTagIds).MatchAllExcludedTags(matchAllExcludedTags).IsReportDownload(isReportDownload).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksWithCards(includeTasksWithCards).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCustomFields(includeCustomFields).IncludeCrmDealIds(includeCrmDealIds).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompanyUserIds(includeCompanyUserIds).IncludeBlockedTasks(includeBlockedTasks).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).ExcludeAssigneeNotOnProjectTeams(excludeAssigneeNotOnProjectTeams).AllowTemplateTasks(allowTemplateTasks).TasklistId(tasklistId).TaskIdList(taskIdList).Tags(tags).TagIds(tagIds).SkipCrmDealIds(skipCrmDealIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ProjectCategoryId(projectCategoryId).FollowerIds(followerIds).FilterText(filterText).FilterBoardColumnIds(filterBoardColumnIds).ExcludeTagIds(excludeTagIds).CrmDealIds(crmDealIds).CreatorIds(creatorIds).CompanyId(companyId).AssigneeTeamIds(assigneeTeamIds).AssigneeCompanyIds(assigneeCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2Tasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2TasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfterDate** | **string** |  | 
 **today** | **string** |  | 
 **startDate** | **string** |  | 
 **sortOrder** | **string** |  | 
 **sort** | **string** |  | 
 **reportFormat** | **string** |  | 
 **projectStatus** | **string** |  | 
 **priority** | **string** |  | 
 **include** | **string** |  | 
 **filter** | **string** |  | 
 **endDate** | **string** |  | 
 **deletedAfterDate** | **string** |  | 
 **dateCode** | **string** |  | 
 **dataSet** | **string** |  | 
 **createdFilter** | **string** |  | 
 **createdDateCode** | **string** |  | 
 **createdBeforeDate** | **string** |  | 
 **createdAfterDate** | **string** |  | 
 **completedBeforeDate** | **string** |  | 
 **completedAfterDate** | **string** |  | 
 **callback** | **string** |  | 
 **taskId** | **int32** |  | 
 **parentTaskId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **offsetCount** | **int32** |  | 
 **includeTaskId** | **int32** |  | 
 **useStartDatesForTodaysTasks** | **bool** |  | 
 **useAllProjects** | **bool** |  | 
 **starredProjectsOnly** | **bool** |  | 
 **showDeleted** | **bool** |  | 
 **showCompletedLists** | **bool** |  | 
 **onlyUntaggedTasks** | **bool** |  | 
 **onlyTasksWithUnreadComments** | **bool** |  | 
 **onlyTasksWithTickets** | **bool** |  | 
 **onlyTasksWithEstimatedTime** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **nestSubTasks** | **bool** |  | 
 **matchAllTags** | **bool** |  | 
 **matchAllProjectTags** | **bool** |  | 
 **matchAllProjectTagIds** | **bool** |  | 
 **matchAllExcludedTags** | **bool** |  | 
 **isReportDownload** | **bool** |  | 
 **includeUntaggedTasks** | **bool** |  | 
 **includeToday** | **bool** |  | 
 **includeTeamUserIds** | **bool** |  | 
 **includeTasksWithoutDueDates** | **bool** |  | 
 **includeTasksWithCards** | **bool** |  | 
 **includeTasksFromDeletedLists** | **bool** |  | 
 **includeReminders** | **bool** |  | 
 **includeLoggedTime** | **bool** |  | 
 **includeCustomFields** | **bool** |  | 
 **includeCrmDealIds** | **bool** |  | 
 **includeCompletedTasks** | **bool** |  | 
 **includeCompletedSubtasks** | **bool** |  | 
 **includeCompanyUserIds** | **bool** |  | 
 **includeBlockedTasks** | **bool** |  | 
 **includeAssigneeTeams** | **bool** |  | 
 **includeAssigneeCompanies** | **bool** |  | 
 **includeArchivedProjects** | **bool** |  | 
 **ignoreStartDates** | **bool** |  | 
 **getSubTasks** | **bool** |  | 
 **getFiles** | **bool** |  | 
 **excludeAssigneeNotOnProjectTeams** | **bool** |  | 
 **allowTemplateTasks** | **bool** |  | 
 **tasklistId** | **[]int32** |  | 
 **taskIdList** | **[]int32** |  | 
 **tags** | **[]string** |  | 
 **tagIds** | **[]int32** |  | 
 **skipCrmDealIds** | **[]int32** |  | 
 **responsiblePartyIds** | **[]int32** |  | 
 **projectTagIds** | **[]int32** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **projectIds** | **[]int32** |  | 
 **projectId** | **[]int32** | project filters | 
 **projectHealths** | **[]int32** |  | 
 **projectHealth** | **[]int32** |  | 
 **projectCompanyIds** | **[]int32** |  | 
 **projectCategoryIds** | **[]int32** |  | 
 **projectCategoryId** | **[]int32** |  | 
 **followerIds** | **[]int32** |  | 
 **filterText** | **[]string** |  | 
 **filterBoardColumnIds** | **[]int32** |  | 
 **excludeTagIds** | **[]int32** |  | 
 **crmDealIds** | **[]int32** |  | 
 **creatorIds** | **[]int32** |  | 
 **companyId** | **[]int32** |  | 
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


## GETProjectsApiV2TasksdependentTaskIdPredecessors

> GETProjectsApiV2TasksdependentTaskIdPredecessors(ctx, dependentTaskId2).Include(include).ProjectId(projectId).PageSize(pageSize).Page(page).DependentTaskId(dependentTaskId).OnlyDependencies(onlyDependencies).OnlyBasicFields(onlyBasicFields).OnlyActionable(onlyActionable).IncludeDependencies(includeDependencies).GetRecursively(getRecursively).Execute()

Get Predecessor

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
    dependentTaskId2 := int32(56) // int32 | 
    include := "include_example" // string |  (optional)
    projectId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    dependentTaskId := int32(56) // int32 |  (optional)
    onlyDependencies := true // bool |  (optional)
    onlyBasicFields := true // bool |  (optional)
    onlyActionable := true // bool |  (optional)
    includeDependencies := true // bool |  (optional)
    getRecursively := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2TasksdependentTaskIdPredecessors(context.Background(), dependentTaskId2).Include(include).ProjectId(projectId).PageSize(pageSize).Page(page).DependentTaskId(dependentTaskId).OnlyDependencies(onlyDependencies).OnlyBasicFields(onlyBasicFields).OnlyActionable(onlyActionable).IncludeDependencies(includeDependencies).GetRecursively(getRecursively).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2TasksdependentTaskIdPredecessors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dependentTaskId2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2TasksdependentTaskIdPredecessorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **string** |  | 
 **projectId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **dependentTaskId** | **int32** |  | 
 **onlyDependencies** | **bool** |  | 
 **onlyBasicFields** | **bool** |  | 
 **onlyActionable** | **bool** |  | 
 **includeDependencies** | **bool** |  | 
 **getRecursively** | **bool** |  | 

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


## GETProjectsApiV2TasksobjectIdProject

> GETProjectsApiV2TasksobjectIdProject(ctx, objectId2).UserId(userId).ObjectId(objectId).ShowDeleted(showDeleted).IncludeTags(includeTags).IncludePeople(includePeople).GetTruePermissions(getTruePermissions).GetPermissions(getPermissions).GetNotificationSettings(getNotificationSettings).GetEmoji(getEmoji).GetEmailAddress(getEmailAddress).GetDateInfo(getDateInfo).GetActivePages(getActivePages).FormatMarkdown(formatMarkdown).Execute()

Get the project for the task

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
    objectId2 := int32(56) // int32 | 
    userId := int32(56) // int32 |  (optional)
    objectId := int32(56) // int32 |  (optional)
    showDeleted := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    includePeople := true // bool |  (optional)
    getTruePermissions := true // bool |  (optional)
    getPermissions := true // bool |  (optional)
    getNotificationSettings := true // bool |  (optional)
    getEmoji := true // bool |  (optional)
    getEmailAddress := true // bool |  (optional)
    getDateInfo := true // bool |  (optional)
    getActivePages := true // bool |  (optional)
    formatMarkdown := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2TasksobjectIdProject(context.Background(), objectId2).UserId(userId).ObjectId(objectId).ShowDeleted(showDeleted).IncludeTags(includeTags).IncludePeople(includePeople).GetTruePermissions(getTruePermissions).GetPermissions(getPermissions).GetNotificationSettings(getNotificationSettings).GetEmoji(getEmoji).GetEmailAddress(getEmailAddress).GetDateInfo(getDateInfo).GetActivePages(getActivePages).FormatMarkdown(formatMarkdown).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2TasksobjectIdProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2TasksobjectIdProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **int32** |  | 
 **objectId** | **int32** |  | 
 **showDeleted** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **includePeople** | **bool** |  | 
 **getTruePermissions** | **bool** |  | 
 **getPermissions** | **bool** |  | 
 **getNotificationSettings** | **bool** |  | 
 **getEmoji** | **bool** |  | 
 **getEmailAddress** | **bool** |  | 
 **getDateInfo** | **bool** |  | 
 **getActivePages** | **bool** |  | 
 **formatMarkdown** | **bool** |  | 

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


## GETProjectsApiV2TaskstaskId

> GETProjectsApiV2TaskstaskId(ctx, taskId2).UpdatedAfterDate(updatedAfterDate).Today(today).StartDate(startDate).SortOrder(sortOrder).Sort(sort).ReportFormat(reportFormat).ProjectStatus(projectStatus).Priority(priority).Include(include).Filter(filter).EndDate(endDate).DeletedAfterDate(deletedAfterDate).DateCode(dateCode).DataSet(dataSet).CreatedFilter(createdFilter).CreatedDateCode(createdDateCode).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TaskId(taskId).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).OffsetCount(offsetCount).IncludeTaskId(includeTaskId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyTasksWithUnreadComments(onlyTasksWithUnreadComments).OnlyTasksWithTickets(onlyTasksWithTickets).OnlyTasksWithEstimatedTime(onlyTasksWithEstimatedTime).OnlyStarredProjects(onlyStarredProjects).NestSubTasks(nestSubTasks).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTagIds(matchAllProjectTagIds).MatchAllExcludedTags(matchAllExcludedTags).IsReportDownload(isReportDownload).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksWithCards(includeTasksWithCards).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCustomFields(includeCustomFields).IncludeCrmDealIds(includeCrmDealIds).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompanyUserIds(includeCompanyUserIds).IncludeBlockedTasks(includeBlockedTasks).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).ExcludeAssigneeNotOnProjectTeams(excludeAssigneeNotOnProjectTeams).AllowTemplateTasks(allowTemplateTasks).TasklistId(tasklistId).TaskIdList(taskIdList).Tags(tags).TagIds(tagIds).SkipCrmDealIds(skipCrmDealIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ProjectCategoryId(projectCategoryId).FollowerIds(followerIds).FilterText(filterText).FilterBoardColumnIds(filterBoardColumnIds).ExcludeTagIds(excludeTagIds).CrmDealIds(crmDealIds).CreatorIds(creatorIds).CompanyId(companyId).AssigneeTeamIds(assigneeTeamIds).AssigneeCompanyIds(assigneeCompanyIds).Execute()

Get Tasks

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
    updatedAfterDate := "updatedAfterDate_example" // string |  (optional)
    today := "today_example" // string |  (optional)
    startDate := "startDate_example" // string |  (optional)
    sortOrder := "sortOrder_example" // string |  (optional)
    sort := "sort_example" // string |  (optional)
    reportFormat := "reportFormat_example" // string |  (optional)
    projectStatus := "projectStatus_example" // string |  (optional)
    priority := "priority_example" // string |  (optional)
    include := "include_example" // string |  (optional)
    filter := "filter_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    deletedAfterDate := "deletedAfterDate_example" // string |  (optional)
    dateCode := "dateCode_example" // string |  (optional)
    dataSet := "dataSet_example" // string |  (optional)
    createdFilter := "createdFilter_example" // string |  (optional)
    createdDateCode := "createdDateCode_example" // string |  (optional)
    createdBeforeDate := "createdBeforeDate_example" // string |  (optional)
    createdAfterDate := "createdAfterDate_example" // string |  (optional)
    completedBeforeDate := "completedBeforeDate_example" // string |  (optional)
    completedAfterDate := "completedAfterDate_example" // string |  (optional)
    callback := "callback_example" // string |  (optional)
    taskId := int32(56) // int32 |  (optional)
    parentTaskId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    offsetCount := int32(56) // int32 |  (optional)
    includeTaskId := int32(56) // int32 |  (optional)
    useStartDatesForTodaysTasks := true // bool |  (optional)
    useAllProjects := true // bool |  (optional)
    starredProjectsOnly := true // bool |  (optional)
    showDeleted := true // bool |  (optional)
    showCompletedLists := true // bool |  (optional)
    onlyUntaggedTasks := true // bool |  (optional)
    onlyTasksWithUnreadComments := true // bool |  (optional)
    onlyTasksWithTickets := true // bool |  (optional)
    onlyTasksWithEstimatedTime := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    nestSubTasks := true // bool |  (optional)
    matchAllTags := true // bool |  (optional)
    matchAllProjectTags := true // bool |  (optional)
    matchAllProjectTagIds := true // bool |  (optional)
    matchAllExcludedTags := true // bool |  (optional)
    isReportDownload := true // bool |  (optional)
    includeUntaggedTasks := true // bool |  (optional)
    includeToday := true // bool |  (optional)
    includeTeamUserIds := true // bool |  (optional)
    includeTasksWithoutDueDates := true // bool |  (optional)
    includeTasksWithCards := true // bool |  (optional)
    includeTasksFromDeletedLists := true // bool |  (optional)
    includeReminders := true // bool |  (optional)
    includeLoggedTime := true // bool |  (optional)
    includeCustomFields := true // bool |  (optional)
    includeCrmDealIds := true // bool |  (optional)
    includeCompletedTasks := true // bool |  (optional)
    includeCompletedSubtasks := true // bool |  (optional)
    includeCompanyUserIds := true // bool |  (optional)
    includeBlockedTasks := true // bool |  (optional)
    includeAssigneeTeams := true // bool |  (optional)
    includeAssigneeCompanies := true // bool |  (optional)
    includeArchivedProjects := true // bool |  (optional)
    ignoreStartDates := true // bool |  (optional)
    getSubTasks := true // bool |  (optional)
    getFiles := true // bool |  (optional)
    excludeAssigneeNotOnProjectTeams := true // bool |  (optional)
    allowTemplateTasks := true // bool |  (optional)
    tasklistId := []int32{int32(123)} // []int32 |  (optional)
    taskIdList := []int32{int32(123)} // []int32 |  (optional)
    tags := []string{"Inner_example"} // []string |  (optional)
    tagIds := []int32{int32(123)} // []int32 |  (optional)
    skipCrmDealIds := []int32{int32(123)} // []int32 |  (optional)
    responsiblePartyIds := []int32{int32(123)} // []int32 |  (optional)
    projectTagIds := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    projectIds := []int32{int32(123)} // []int32 |  (optional)
    projectId := []int32{int32(123)} // []int32 | project filters (optional)
    projectHealths := []int32{int32(123)} // []int32 |  (optional)
    projectHealth := []int32{int32(123)} // []int32 |  (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryId := []int32{int32(123)} // []int32 |  (optional)
    followerIds := []int32{int32(123)} // []int32 |  (optional)
    filterText := []string{"Inner_example"} // []string |  (optional)
    filterBoardColumnIds := []int32{int32(123)} // []int32 |  (optional)
    excludeTagIds := []int32{int32(123)} // []int32 |  (optional)
    crmDealIds := []int32{int32(123)} // []int32 |  (optional)
    creatorIds := []int32{int32(123)} // []int32 |  (optional)
    companyId := []int32{int32(123)} // []int32 |  (optional)
    assigneeTeamIds := []int32{int32(123)} // []int32 |  (optional)
    assigneeCompanyIds := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2TaskstaskId(context.Background(), taskId2).UpdatedAfterDate(updatedAfterDate).Today(today).StartDate(startDate).SortOrder(sortOrder).Sort(sort).ReportFormat(reportFormat).ProjectStatus(projectStatus).Priority(priority).Include(include).Filter(filter).EndDate(endDate).DeletedAfterDate(deletedAfterDate).DateCode(dateCode).DataSet(dataSet).CreatedFilter(createdFilter).CreatedDateCode(createdDateCode).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TaskId(taskId).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).OffsetCount(offsetCount).IncludeTaskId(includeTaskId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyTasksWithUnreadComments(onlyTasksWithUnreadComments).OnlyTasksWithTickets(onlyTasksWithTickets).OnlyTasksWithEstimatedTime(onlyTasksWithEstimatedTime).OnlyStarredProjects(onlyStarredProjects).NestSubTasks(nestSubTasks).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTagIds(matchAllProjectTagIds).MatchAllExcludedTags(matchAllExcludedTags).IsReportDownload(isReportDownload).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksWithCards(includeTasksWithCards).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCustomFields(includeCustomFields).IncludeCrmDealIds(includeCrmDealIds).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompanyUserIds(includeCompanyUserIds).IncludeBlockedTasks(includeBlockedTasks).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).ExcludeAssigneeNotOnProjectTeams(excludeAssigneeNotOnProjectTeams).AllowTemplateTasks(allowTemplateTasks).TasklistId(tasklistId).TaskIdList(taskIdList).Tags(tags).TagIds(tagIds).SkipCrmDealIds(skipCrmDealIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ProjectCategoryId(projectCategoryId).FollowerIds(followerIds).FilterText(filterText).FilterBoardColumnIds(filterBoardColumnIds).ExcludeTagIds(excludeTagIds).CrmDealIds(crmDealIds).CreatorIds(creatorIds).CompanyId(companyId).AssigneeTeamIds(assigneeTeamIds).AssigneeCompanyIds(assigneeCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2TaskstaskId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2TaskstaskIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedAfterDate** | **string** |  | 
 **today** | **string** |  | 
 **startDate** | **string** |  | 
 **sortOrder** | **string** |  | 
 **sort** | **string** |  | 
 **reportFormat** | **string** |  | 
 **projectStatus** | **string** |  | 
 **priority** | **string** |  | 
 **include** | **string** |  | 
 **filter** | **string** |  | 
 **endDate** | **string** |  | 
 **deletedAfterDate** | **string** |  | 
 **dateCode** | **string** |  | 
 **dataSet** | **string** |  | 
 **createdFilter** | **string** |  | 
 **createdDateCode** | **string** |  | 
 **createdBeforeDate** | **string** |  | 
 **createdAfterDate** | **string** |  | 
 **completedBeforeDate** | **string** |  | 
 **completedAfterDate** | **string** |  | 
 **callback** | **string** |  | 
 **taskId** | **int32** |  | 
 **parentTaskId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **offsetCount** | **int32** |  | 
 **includeTaskId** | **int32** |  | 
 **useStartDatesForTodaysTasks** | **bool** |  | 
 **useAllProjects** | **bool** |  | 
 **starredProjectsOnly** | **bool** |  | 
 **showDeleted** | **bool** |  | 
 **showCompletedLists** | **bool** |  | 
 **onlyUntaggedTasks** | **bool** |  | 
 **onlyTasksWithUnreadComments** | **bool** |  | 
 **onlyTasksWithTickets** | **bool** |  | 
 **onlyTasksWithEstimatedTime** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **nestSubTasks** | **bool** |  | 
 **matchAllTags** | **bool** |  | 
 **matchAllProjectTags** | **bool** |  | 
 **matchAllProjectTagIds** | **bool** |  | 
 **matchAllExcludedTags** | **bool** |  | 
 **isReportDownload** | **bool** |  | 
 **includeUntaggedTasks** | **bool** |  | 
 **includeToday** | **bool** |  | 
 **includeTeamUserIds** | **bool** |  | 
 **includeTasksWithoutDueDates** | **bool** |  | 
 **includeTasksWithCards** | **bool** |  | 
 **includeTasksFromDeletedLists** | **bool** |  | 
 **includeReminders** | **bool** |  | 
 **includeLoggedTime** | **bool** |  | 
 **includeCustomFields** | **bool** |  | 
 **includeCrmDealIds** | **bool** |  | 
 **includeCompletedTasks** | **bool** |  | 
 **includeCompletedSubtasks** | **bool** |  | 
 **includeCompanyUserIds** | **bool** |  | 
 **includeBlockedTasks** | **bool** |  | 
 **includeAssigneeTeams** | **bool** |  | 
 **includeAssigneeCompanies** | **bool** |  | 
 **includeArchivedProjects** | **bool** |  | 
 **ignoreStartDates** | **bool** |  | 
 **getSubTasks** | **bool** |  | 
 **getFiles** | **bool** |  | 
 **excludeAssigneeNotOnProjectTeams** | **bool** |  | 
 **allowTemplateTasks** | **bool** |  | 
 **tasklistId** | **[]int32** |  | 
 **taskIdList** | **[]int32** |  | 
 **tags** | **[]string** |  | 
 **tagIds** | **[]int32** |  | 
 **skipCrmDealIds** | **[]int32** |  | 
 **responsiblePartyIds** | **[]int32** |  | 
 **projectTagIds** | **[]int32** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **projectIds** | **[]int32** |  | 
 **projectId** | **[]int32** | project filters | 
 **projectHealths** | **[]int32** |  | 
 **projectHealth** | **[]int32** |  | 
 **projectCompanyIds** | **[]int32** |  | 
 **projectCategoryIds** | **[]int32** |  | 
 **projectCategoryId** | **[]int32** |  | 
 **followerIds** | **[]int32** |  | 
 **filterText** | **[]string** |  | 
 **filterBoardColumnIds** | **[]int32** |  | 
 **excludeTagIds** | **[]int32** |  | 
 **crmDealIds** | **[]int32** |  | 
 **creatorIds** | **[]int32** |  | 
 **companyId** | **[]int32** |  | 
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


## GETProjectsApiV2TaskstaskIdEstimatedtime

> GETProjectsApiV2TaskstaskIdEstimatedtime(ctx, taskId).TaskID(taskID).Execute()

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
    taskId := int32(56) // int32 | 
    taskID := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2TaskstaskIdEstimatedtime(context.Background(), taskId).TaskID(taskID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2TaskstaskIdEstimatedtime``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETProjectsApiV2TaskstaskIdEstimatedtimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taskID** | **int32** |  | 

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


## GETProjectsApiV2TaskstaskIdSubtasks

> GETProjectsApiV2TaskstaskIdSubtasks(ctx, taskId2).UpdatedAfterDate(updatedAfterDate).Today(today).StartDate(startDate).SortOrder(sortOrder).Sort(sort).ReportFormat(reportFormat).ProjectStatus(projectStatus).Priority(priority).Include(include).Filter(filter).EndDate(endDate).DeletedAfterDate(deletedAfterDate).DateCode(dateCode).DataSet(dataSet).CreatedFilter(createdFilter).CreatedDateCode(createdDateCode).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TaskId(taskId).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).OffsetCount(offsetCount).IncludeTaskId(includeTaskId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyTasksWithUnreadComments(onlyTasksWithUnreadComments).OnlyTasksWithTickets(onlyTasksWithTickets).OnlyTasksWithEstimatedTime(onlyTasksWithEstimatedTime).OnlyStarredProjects(onlyStarredProjects).NestSubTasks(nestSubTasks).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTagIds(matchAllProjectTagIds).MatchAllExcludedTags(matchAllExcludedTags).IsReportDownload(isReportDownload).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksWithCards(includeTasksWithCards).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCustomFields(includeCustomFields).IncludeCrmDealIds(includeCrmDealIds).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompanyUserIds(includeCompanyUserIds).IncludeBlockedTasks(includeBlockedTasks).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).ExcludeAssigneeNotOnProjectTeams(excludeAssigneeNotOnProjectTeams).AllowTemplateTasks(allowTemplateTasks).TasklistId(tasklistId).TaskIdList(taskIdList).Tags(tags).TagIds(tagIds).SkipCrmDealIds(skipCrmDealIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ProjectCategoryId(projectCategoryId).FollowerIds(followerIds).FilterText(filterText).FilterBoardColumnIds(filterBoardColumnIds).ExcludeTagIds(excludeTagIds).CrmDealIds(crmDealIds).CreatorIds(creatorIds).CompanyId(companyId).AssigneeTeamIds(assigneeTeamIds).AssigneeCompanyIds(assigneeCompanyIds).Execute()

Get Tasks

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
    updatedAfterDate := "updatedAfterDate_example" // string |  (optional)
    today := "today_example" // string |  (optional)
    startDate := "startDate_example" // string |  (optional)
    sortOrder := "sortOrder_example" // string |  (optional)
    sort := "sort_example" // string |  (optional)
    reportFormat := "reportFormat_example" // string |  (optional)
    projectStatus := "projectStatus_example" // string |  (optional)
    priority := "priority_example" // string |  (optional)
    include := "include_example" // string |  (optional)
    filter := "filter_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    deletedAfterDate := "deletedAfterDate_example" // string |  (optional)
    dateCode := "dateCode_example" // string |  (optional)
    dataSet := "dataSet_example" // string |  (optional)
    createdFilter := "createdFilter_example" // string |  (optional)
    createdDateCode := "createdDateCode_example" // string |  (optional)
    createdBeforeDate := "createdBeforeDate_example" // string |  (optional)
    createdAfterDate := "createdAfterDate_example" // string |  (optional)
    completedBeforeDate := "completedBeforeDate_example" // string |  (optional)
    completedAfterDate := "completedAfterDate_example" // string |  (optional)
    callback := "callback_example" // string |  (optional)
    taskId := int32(56) // int32 |  (optional)
    parentTaskId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    offsetCount := int32(56) // int32 |  (optional)
    includeTaskId := int32(56) // int32 |  (optional)
    useStartDatesForTodaysTasks := true // bool |  (optional)
    useAllProjects := true // bool |  (optional)
    starredProjectsOnly := true // bool |  (optional)
    showDeleted := true // bool |  (optional)
    showCompletedLists := true // bool |  (optional)
    onlyUntaggedTasks := true // bool |  (optional)
    onlyTasksWithUnreadComments := true // bool |  (optional)
    onlyTasksWithTickets := true // bool |  (optional)
    onlyTasksWithEstimatedTime := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    nestSubTasks := true // bool |  (optional)
    matchAllTags := true // bool |  (optional)
    matchAllProjectTags := true // bool |  (optional)
    matchAllProjectTagIds := true // bool |  (optional)
    matchAllExcludedTags := true // bool |  (optional)
    isReportDownload := true // bool |  (optional)
    includeUntaggedTasks := true // bool |  (optional)
    includeToday := true // bool |  (optional)
    includeTeamUserIds := true // bool |  (optional)
    includeTasksWithoutDueDates := true // bool |  (optional)
    includeTasksWithCards := true // bool |  (optional)
    includeTasksFromDeletedLists := true // bool |  (optional)
    includeReminders := true // bool |  (optional)
    includeLoggedTime := true // bool |  (optional)
    includeCustomFields := true // bool |  (optional)
    includeCrmDealIds := true // bool |  (optional)
    includeCompletedTasks := true // bool |  (optional)
    includeCompletedSubtasks := true // bool |  (optional)
    includeCompanyUserIds := true // bool |  (optional)
    includeBlockedTasks := true // bool |  (optional)
    includeAssigneeTeams := true // bool |  (optional)
    includeAssigneeCompanies := true // bool |  (optional)
    includeArchivedProjects := true // bool |  (optional)
    ignoreStartDates := true // bool |  (optional)
    getSubTasks := true // bool |  (optional)
    getFiles := true // bool |  (optional)
    excludeAssigneeNotOnProjectTeams := true // bool |  (optional)
    allowTemplateTasks := true // bool |  (optional)
    tasklistId := []int32{int32(123)} // []int32 |  (optional)
    taskIdList := []int32{int32(123)} // []int32 |  (optional)
    tags := []string{"Inner_example"} // []string |  (optional)
    tagIds := []int32{int32(123)} // []int32 |  (optional)
    skipCrmDealIds := []int32{int32(123)} // []int32 |  (optional)
    responsiblePartyIds := []int32{int32(123)} // []int32 |  (optional)
    projectTagIds := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    projectIds := []int32{int32(123)} // []int32 |  (optional)
    projectId := []int32{int32(123)} // []int32 | project filters (optional)
    projectHealths := []int32{int32(123)} // []int32 |  (optional)
    projectHealth := []int32{int32(123)} // []int32 |  (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryId := []int32{int32(123)} // []int32 |  (optional)
    followerIds := []int32{int32(123)} // []int32 |  (optional)
    filterText := []string{"Inner_example"} // []string |  (optional)
    filterBoardColumnIds := []int32{int32(123)} // []int32 |  (optional)
    excludeTagIds := []int32{int32(123)} // []int32 |  (optional)
    crmDealIds := []int32{int32(123)} // []int32 |  (optional)
    creatorIds := []int32{int32(123)} // []int32 |  (optional)
    companyId := []int32{int32(123)} // []int32 |  (optional)
    assigneeTeamIds := []int32{int32(123)} // []int32 |  (optional)
    assigneeCompanyIds := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2TaskstaskIdSubtasks(context.Background(), taskId2).UpdatedAfterDate(updatedAfterDate).Today(today).StartDate(startDate).SortOrder(sortOrder).Sort(sort).ReportFormat(reportFormat).ProjectStatus(projectStatus).Priority(priority).Include(include).Filter(filter).EndDate(endDate).DeletedAfterDate(deletedAfterDate).DateCode(dateCode).DataSet(dataSet).CreatedFilter(createdFilter).CreatedDateCode(createdDateCode).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedAfterDate(completedAfterDate).Callback(callback).TaskId(taskId).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).OffsetCount(offsetCount).IncludeTaskId(includeTaskId).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).StarredProjectsOnly(starredProjectsOnly).ShowDeleted(showDeleted).ShowCompletedLists(showCompletedLists).OnlyUntaggedTasks(onlyUntaggedTasks).OnlyTasksWithUnreadComments(onlyTasksWithUnreadComments).OnlyTasksWithTickets(onlyTasksWithTickets).OnlyTasksWithEstimatedTime(onlyTasksWithEstimatedTime).OnlyStarredProjects(onlyStarredProjects).NestSubTasks(nestSubTasks).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTagIds(matchAllProjectTagIds).MatchAllExcludedTags(matchAllExcludedTags).IsReportDownload(isReportDownload).IncludeUntaggedTasks(includeUntaggedTasks).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksWithCards(includeTasksWithCards).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeReminders(includeReminders).IncludeLoggedTime(includeLoggedTime).IncludeCustomFields(includeCustomFields).IncludeCrmDealIds(includeCrmDealIds).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompanyUserIds(includeCompanyUserIds).IncludeBlockedTasks(includeBlockedTasks).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetFiles(getFiles).ExcludeAssigneeNotOnProjectTeams(excludeAssigneeNotOnProjectTeams).AllowTemplateTasks(allowTemplateTasks).TasklistId(tasklistId).TaskIdList(taskIdList).Tags(tags).TagIds(tagIds).SkipCrmDealIds(skipCrmDealIds).ResponsiblePartyIds(responsiblePartyIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ProjectCategoryId(projectCategoryId).FollowerIds(followerIds).FilterText(filterText).FilterBoardColumnIds(filterBoardColumnIds).ExcludeTagIds(excludeTagIds).CrmDealIds(crmDealIds).CreatorIds(creatorIds).CompanyId(companyId).AssigneeTeamIds(assigneeTeamIds).AssigneeCompanyIds(assigneeCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2TaskstaskIdSubtasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2TaskstaskIdSubtasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedAfterDate** | **string** |  | 
 **today** | **string** |  | 
 **startDate** | **string** |  | 
 **sortOrder** | **string** |  | 
 **sort** | **string** |  | 
 **reportFormat** | **string** |  | 
 **projectStatus** | **string** |  | 
 **priority** | **string** |  | 
 **include** | **string** |  | 
 **filter** | **string** |  | 
 **endDate** | **string** |  | 
 **deletedAfterDate** | **string** |  | 
 **dateCode** | **string** |  | 
 **dataSet** | **string** |  | 
 **createdFilter** | **string** |  | 
 **createdDateCode** | **string** |  | 
 **createdBeforeDate** | **string** |  | 
 **createdAfterDate** | **string** |  | 
 **completedBeforeDate** | **string** |  | 
 **completedAfterDate** | **string** |  | 
 **callback** | **string** |  | 
 **taskId** | **int32** |  | 
 **parentTaskId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **offsetCount** | **int32** |  | 
 **includeTaskId** | **int32** |  | 
 **useStartDatesForTodaysTasks** | **bool** |  | 
 **useAllProjects** | **bool** |  | 
 **starredProjectsOnly** | **bool** |  | 
 **showDeleted** | **bool** |  | 
 **showCompletedLists** | **bool** |  | 
 **onlyUntaggedTasks** | **bool** |  | 
 **onlyTasksWithUnreadComments** | **bool** |  | 
 **onlyTasksWithTickets** | **bool** |  | 
 **onlyTasksWithEstimatedTime** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **nestSubTasks** | **bool** |  | 
 **matchAllTags** | **bool** |  | 
 **matchAllProjectTags** | **bool** |  | 
 **matchAllProjectTagIds** | **bool** |  | 
 **matchAllExcludedTags** | **bool** |  | 
 **isReportDownload** | **bool** |  | 
 **includeUntaggedTasks** | **bool** |  | 
 **includeToday** | **bool** |  | 
 **includeTeamUserIds** | **bool** |  | 
 **includeTasksWithoutDueDates** | **bool** |  | 
 **includeTasksWithCards** | **bool** |  | 
 **includeTasksFromDeletedLists** | **bool** |  | 
 **includeReminders** | **bool** |  | 
 **includeLoggedTime** | **bool** |  | 
 **includeCustomFields** | **bool** |  | 
 **includeCrmDealIds** | **bool** |  | 
 **includeCompletedTasks** | **bool** |  | 
 **includeCompletedSubtasks** | **bool** |  | 
 **includeCompanyUserIds** | **bool** |  | 
 **includeBlockedTasks** | **bool** |  | 
 **includeAssigneeTeams** | **bool** |  | 
 **includeAssigneeCompanies** | **bool** |  | 
 **includeArchivedProjects** | **bool** |  | 
 **ignoreStartDates** | **bool** |  | 
 **getSubTasks** | **bool** |  | 
 **getFiles** | **bool** |  | 
 **excludeAssigneeNotOnProjectTeams** | **bool** |  | 
 **allowTemplateTasks** | **bool** |  | 
 **tasklistId** | **[]int32** |  | 
 **taskIdList** | **[]int32** |  | 
 **tags** | **[]string** |  | 
 **tagIds** | **[]int32** |  | 
 **skipCrmDealIds** | **[]int32** |  | 
 **responsiblePartyIds** | **[]int32** |  | 
 **projectTagIds** | **[]int32** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **projectIds** | **[]int32** |  | 
 **projectId** | **[]int32** | project filters | 
 **projectHealths** | **[]int32** |  | 
 **projectHealth** | **[]int32** |  | 
 **projectCompanyIds** | **[]int32** |  | 
 **projectCategoryIds** | **[]int32** |  | 
 **projectCategoryId** | **[]int32** |  | 
 **followerIds** | **[]int32** |  | 
 **filterText** | **[]string** |  | 
 **filterBoardColumnIds** | **[]int32** |  | 
 **excludeTagIds** | **[]int32** |  | 
 **crmDealIds** | **[]int32** |  | 
 **creatorIds** | **[]int32** |  | 
 **companyId** | **[]int32** |  | 
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


## GETProjectsApiV2TaskstaskIdTime

> GETProjectsApiV2TaskstaskIdTime(ctx, taskId2).SortOrder(sortOrder).TaskId(taskId).PageSize(pageSize).Page(page).IncludeTimelogId(includeTimelogId).IncludeTags(includeTags).IncludeSubTasks(includeSubTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).GetTotals(getTotals).Execute()

APICall_GET_todo_items_id_time_entries

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
    sortOrder := "sortOrder_example" // string |  (optional)
    taskId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    includeTimelogId := int32(56) // int32 |  (optional)
    includeTags := true // bool |  (optional)
    includeSubTasks := true // bool |  (optional)
    includeCompletedSubtasks := true // bool |  (optional)
    getTotals := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2TaskstaskIdTime(context.Background(), taskId2).SortOrder(sortOrder).TaskId(taskId).PageSize(pageSize).Page(page).IncludeTimelogId(includeTimelogId).IncludeTags(includeTags).IncludeSubTasks(includeSubTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).GetTotals(getTotals).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2TaskstaskIdTime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2TaskstaskIdTimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortOrder** | **string** |  | 
 **taskId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **includeTimelogId** | **int32** |  | 
 **includeTags** | **bool** |  | 
 **includeSubTasks** | **bool** |  | 
 **includeCompletedSubtasks** | **bool** |  | 
 **getTotals** | **bool** |  | 

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


## GETProjectsApiV2TaskstaskIdTimeEntries

> GETProjectsApiV2TaskstaskIdTimeEntries(ctx, taskId2).SortOrder(sortOrder).TaskId(taskId).PageSize(pageSize).Page(page).IncludeTimelogId(includeTimelogId).IncludeTags(includeTags).IncludeSubTasks(includeSubTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).GetTotals(getTotals).Execute()

APICall_GET_todo_items_id_time_entries

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
    sortOrder := "sortOrder_example" // string |  (optional)
    taskId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    includeTimelogId := int32(56) // int32 |  (optional)
    includeTags := true // bool |  (optional)
    includeSubTasks := true // bool |  (optional)
    includeCompletedSubtasks := true // bool |  (optional)
    getTotals := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2TaskstaskIdTimeEntries(context.Background(), taskId2).SortOrder(sortOrder).TaskId(taskId).PageSize(pageSize).Page(page).IncludeTimelogId(includeTimelogId).IncludeTags(includeTags).IncludeSubTasks(includeSubTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).GetTotals(getTotals).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2TaskstaskIdTimeEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2TaskstaskIdTimeEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortOrder** | **string** |  | 
 **taskId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **includeTimelogId** | **int32** |  | 
 **includeTags** | **bool** |  | 
 **includeSubTasks** | **bool** |  | 
 **includeCompletedSubtasks** | **bool** |  | 
 **getTotals** | **bool** |  | 

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


## GETProjectsApiV2Time

> GETProjectsApiV2Time(ctx).UpdatedAfterDate(updatedAfterDate).ToDate(toDate).SortOrder(sortOrder).SortBy(sortBy).ProjectStatus(projectStatus).InvoicedType(invoicedType).FromDate(fromDate).BillableType(billableType).TimeLogId(timeLogId).TicketId(ticketId).PageSize(pageSize).Page(page).IncludeTimelogId(includeTimelogId).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTagIds(matchAllProjectTagIds).IncludeTags(includeTags).IncludeMainTask(includeMainTask).IncludeArchivedProjects(includeArchivedProjects).GetTotals(getTotals).UserId(userId).TaskTagIds(taskTagIds).TagIds(tagIds).ProjectsFromCompanyId(projectsFromCompanyId).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ProjectCategoryId(projectCategoryId).CompanyId(companyId).Execute()

Get all timelogs

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
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2Time(context.Background()).UpdatedAfterDate(updatedAfterDate).ToDate(toDate).SortOrder(sortOrder).SortBy(sortBy).ProjectStatus(projectStatus).InvoicedType(invoicedType).FromDate(fromDate).BillableType(billableType).TimeLogId(timeLogId).TicketId(ticketId).PageSize(pageSize).Page(page).IncludeTimelogId(includeTimelogId).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTagIds(matchAllProjectTagIds).IncludeTags(includeTags).IncludeMainTask(includeMainTask).IncludeArchivedProjects(includeArchivedProjects).GetTotals(getTotals).UserId(userId).TaskTagIds(taskTagIds).TagIds(tagIds).ProjectsFromCompanyId(projectsFromCompanyId).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ProjectCategoryId(projectCategoryId).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2Time``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2TimeRequest struct via the builder pattern


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


## GETProjectsApiV2TimeEntries

> GETProjectsApiV2TimeEntries(ctx).UpdatedAfterDate(updatedAfterDate).ToDate(toDate).SortOrder(sortOrder).SortBy(sortBy).ProjectStatus(projectStatus).InvoicedType(invoicedType).FromDate(fromDate).BillableType(billableType).TimeLogId(timeLogId).TicketId(ticketId).PageSize(pageSize).Page(page).IncludeTimelogId(includeTimelogId).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTagIds(matchAllProjectTagIds).IncludeTags(includeTags).IncludeMainTask(includeMainTask).IncludeArchivedProjects(includeArchivedProjects).GetTotals(getTotals).UserId(userId).TaskTagIds(taskTagIds).TagIds(tagIds).ProjectsFromCompanyId(projectsFromCompanyId).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ProjectCategoryId(projectCategoryId).CompanyId(companyId).Execute()

Get all time entries

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
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2TimeEntries(context.Background()).UpdatedAfterDate(updatedAfterDate).ToDate(toDate).SortOrder(sortOrder).SortBy(sortBy).ProjectStatus(projectStatus).InvoicedType(invoicedType).FromDate(fromDate).BillableType(billableType).TimeLogId(timeLogId).TicketId(ticketId).PageSize(pageSize).Page(page).IncludeTimelogId(includeTimelogId).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTagIds(matchAllProjectTagIds).IncludeTags(includeTags).IncludeMainTask(includeMainTask).IncludeArchivedProjects(includeArchivedProjects).GetTotals(getTotals).UserId(userId).TaskTagIds(taskTagIds).TagIds(tagIds).ProjectsFromCompanyId(projectsFromCompanyId).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ProjectCategoryId(projectCategoryId).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2TimeEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2TimeEntriesRequest struct via the builder pattern


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


## GETProjectsApiV2TimeEntriesMe

> GETProjectsApiV2TimeEntriesMe(ctx).ToDate(toDate).SortOrder(sortOrder).FromDate(fromDate).TicketId(ticketId).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).IncludeTags(includeTags).GetTotals(getTotals).UserId(userId).Execute()

APICall_GET_time_entries_currentUser

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
    toDate := "toDate_example" // string |  (optional)
    sortOrder := "sortOrder_example" // string |  (optional)
    fromDate := "fromDate_example" // string |  (optional)
    ticketId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    showDeleted := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    getTotals := true // bool |  (optional)
    userId := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2TimeEntriesMe(context.Background()).ToDate(toDate).SortOrder(sortOrder).FromDate(fromDate).TicketId(ticketId).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).IncludeTags(includeTags).GetTotals(getTotals).UserId(userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2TimeEntriesMe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2TimeEntriesMeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **toDate** | **string** |  | 
 **sortOrder** | **string** |  | 
 **fromDate** | **string** |  | 
 **ticketId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **showDeleted** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **getTotals** | **bool** |  | 
 **userId** | **[]int32** |  | 

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


## GETProjectsApiV2TimeEntriestimeLogId

> GETProjectsApiV2TimeEntriestimeLogId(ctx, timeLogId2).UpdatedAfterDate(updatedAfterDate).ToDate(toDate).SortOrder(sortOrder).SortBy(sortBy).ProjectStatus(projectStatus).InvoicedType(invoicedType).FromDate(fromDate).BillableType(billableType).TimeLogId(timeLogId).TicketId(ticketId).PageSize(pageSize).Page(page).IncludeTimelogId(includeTimelogId).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTagIds(matchAllProjectTagIds).IncludeTags(includeTags).IncludeMainTask(includeMainTask).IncludeArchivedProjects(includeArchivedProjects).GetTotals(getTotals).UserId(userId).TaskTagIds(taskTagIds).TagIds(tagIds).ProjectsFromCompanyId(projectsFromCompanyId).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ProjectCategoryId(projectCategoryId).CompanyId(companyId).Execute()

Get time entries for a timelogId

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
    timeLogId2 := int32(56) // int32 | 
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
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2TimeEntriestimeLogId(context.Background(), timeLogId2).UpdatedAfterDate(updatedAfterDate).ToDate(toDate).SortOrder(sortOrder).SortBy(sortBy).ProjectStatus(projectStatus).InvoicedType(invoicedType).FromDate(fromDate).BillableType(billableType).TimeLogId(timeLogId).TicketId(ticketId).PageSize(pageSize).Page(page).IncludeTimelogId(includeTimelogId).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTagIds(matchAllProjectTagIds).IncludeTags(includeTags).IncludeMainTask(includeMainTask).IncludeArchivedProjects(includeArchivedProjects).GetTotals(getTotals).UserId(userId).TaskTagIds(taskTagIds).TagIds(tagIds).ProjectsFromCompanyId(projectsFromCompanyId).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ProjectCategoryId(projectCategoryId).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2TimeEntriestimeLogId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timeLogId2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2TimeEntriestimeLogIdRequest struct via the builder pattern


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


## GETProjectsApiV2TimeMe

> GETProjectsApiV2TimeMe(ctx).ToDate(toDate).SortOrder(sortOrder).FromDate(fromDate).TicketId(ticketId).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).IncludeTags(includeTags).GetTotals(getTotals).UserId(userId).Execute()

APICall_GET_time_entries_currentUser

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
    toDate := "toDate_example" // string |  (optional)
    sortOrder := "sortOrder_example" // string |  (optional)
    fromDate := "fromDate_example" // string |  (optional)
    ticketId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    showDeleted := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    getTotals := true // bool |  (optional)
    userId := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2TimeMe(context.Background()).ToDate(toDate).SortOrder(sortOrder).FromDate(fromDate).TicketId(ticketId).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).IncludeTags(includeTags).GetTotals(getTotals).UserId(userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2TimeMe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2TimeMeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **toDate** | **string** |  | 
 **sortOrder** | **string** |  | 
 **fromDate** | **string** |  | 
 **ticketId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **showDeleted** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **getTotals** | **bool** |  | 
 **userId** | **[]int32** |  | 

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


## GETProjectsApiV2TimetimeLogId

> GETProjectsApiV2TimetimeLogId(ctx, timeLogId2).UpdatedAfterDate(updatedAfterDate).ToDate(toDate).SortOrder(sortOrder).SortBy(sortBy).ProjectStatus(projectStatus).InvoicedType(invoicedType).FromDate(fromDate).BillableType(billableType).TimeLogId(timeLogId).TicketId(ticketId).PageSize(pageSize).Page(page).IncludeTimelogId(includeTimelogId).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTagIds(matchAllProjectTagIds).IncludeTags(includeTags).IncludeMainTask(includeMainTask).IncludeArchivedProjects(includeArchivedProjects).GetTotals(getTotals).UserId(userId).TaskTagIds(taskTagIds).TagIds(tagIds).ProjectsFromCompanyId(projectsFromCompanyId).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ProjectCategoryId(projectCategoryId).CompanyId(companyId).Execute()

Get timelog for a timelogId

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
    timeLogId2 := int32(56) // int32 | 
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
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2TimetimeLogId(context.Background(), timeLogId2).UpdatedAfterDate(updatedAfterDate).ToDate(toDate).SortOrder(sortOrder).SortBy(sortBy).ProjectStatus(projectStatus).InvoicedType(invoicedType).FromDate(fromDate).BillableType(billableType).TimeLogId(timeLogId).TicketId(ticketId).PageSize(pageSize).Page(page).IncludeTimelogId(includeTimelogId).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTagIds(matchAllProjectTagIds).IncludeTags(includeTags).IncludeMainTask(includeMainTask).IncludeArchivedProjects(includeArchivedProjects).GetTotals(getTotals).UserId(userId).TaskTagIds(taskTagIds).TagIds(tagIds).ProjectsFromCompanyId(projectsFromCompanyId).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ProjectCategoryId(projectCategoryId).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2TimetimeLogId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timeLogId2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2TimetimeLogIdRequest struct via the builder pattern


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


## GETProjectsApiV2TodoItemstaskIdTime

> GETProjectsApiV2TodoItemstaskIdTime(ctx, taskId2).SortOrder(sortOrder).TaskId(taskId).PageSize(pageSize).Page(page).IncludeTimelogId(includeTimelogId).IncludeTags(includeTags).IncludeSubTasks(includeSubTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).GetTotals(getTotals).Execute()

APICall_GET_todo_items_id_time_entries

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
    sortOrder := "sortOrder_example" // string |  (optional)
    taskId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    includeTimelogId := int32(56) // int32 |  (optional)
    includeTags := true // bool |  (optional)
    includeSubTasks := true // bool |  (optional)
    includeCompletedSubtasks := true // bool |  (optional)
    getTotals := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2TodoItemstaskIdTime(context.Background(), taskId2).SortOrder(sortOrder).TaskId(taskId).PageSize(pageSize).Page(page).IncludeTimelogId(includeTimelogId).IncludeTags(includeTags).IncludeSubTasks(includeSubTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).GetTotals(getTotals).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2TodoItemstaskIdTime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2TodoItemstaskIdTimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortOrder** | **string** |  | 
 **taskId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **includeTimelogId** | **int32** |  | 
 **includeTags** | **bool** |  | 
 **includeSubTasks** | **bool** |  | 
 **includeCompletedSubtasks** | **bool** |  | 
 **getTotals** | **bool** |  | 

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


## GETProjectsApiV2TodoItemstaskIdTimeEntries

> GETProjectsApiV2TodoItemstaskIdTimeEntries(ctx, taskId2).SortOrder(sortOrder).TaskId(taskId).PageSize(pageSize).Page(page).IncludeTimelogId(includeTimelogId).IncludeTags(includeTags).IncludeSubTasks(includeSubTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).GetTotals(getTotals).Execute()

APICall_GET_todo_items_id_time_entries

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
    sortOrder := "sortOrder_example" // string |  (optional)
    taskId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    includeTimelogId := int32(56) // int32 |  (optional)
    includeTags := true // bool |  (optional)
    includeSubTasks := true // bool |  (optional)
    includeCompletedSubtasks := true // bool |  (optional)
    getTotals := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CFPortApi.GETProjectsApiV2TodoItemstaskIdTimeEntries(context.Background(), taskId2).SortOrder(sortOrder).TaskId(taskId).PageSize(pageSize).Page(page).IncludeTimelogId(includeTimelogId).IncludeTags(includeTags).IncludeSubTasks(includeSubTasks).IncludeCompletedSubtasks(includeCompletedSubtasks).GetTotals(getTotals).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CFPortApi.GETProjectsApiV2TodoItemstaskIdTimeEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2TodoItemstaskIdTimeEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortOrder** | **string** |  | 
 **taskId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **includeTimelogId** | **int32** |  | 
 **includeTags** | **bool** |  | 
 **includeSubTasks** | **bool** |  | 
 **includeCompletedSubtasks** | **bool** |  | 
 **getTotals** | **bool** |  | 

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

