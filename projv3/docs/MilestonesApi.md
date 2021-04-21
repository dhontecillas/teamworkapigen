# \MilestonesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3MilestonesCsv**](MilestonesApi.md#GETProjectsApiV3MilestonesCsv) | **Get** /projects/api/v3/milestones.csv | Generate milestone report in CSV format
[**GETProjectsApiV3MilestonesHtml**](MilestonesApi.md#GETProjectsApiV3MilestonesHtml) | **Get** /projects/api/v3/milestones.html | Generate milestone report in HTML format
[**GETProjectsApiV3MilestonesJson**](MilestonesApi.md#GETProjectsApiV3MilestonesJson) | **Get** /projects/api/v3/milestones.json | Get all milestones
[**GETProjectsApiV3MilestonesMetricsDeadlinesJson**](MilestonesApi.md#GETProjectsApiV3MilestonesMetricsDeadlinesJson) | **Get** /projects/api/v3/milestones/metrics/deadlines.json | Get milestones by due date in a time range
[**GETProjectsApiV3MilestonesPdf**](MilestonesApi.md#GETProjectsApiV3MilestonesPdf) | **Get** /projects/api/v3/milestones.pdf | Generate milestone report in PDF format
[**GETProjectsApiV3MilestonesXlsx**](MilestonesApi.md#GETProjectsApiV3MilestonesXlsx) | **Get** /projects/api/v3/milestones.xlsx | Generate milestone report in XLSX format
[**GETProjectsApiV3MilestonesmilestonIdJson**](MilestonesApi.md#GETProjectsApiV3MilestonesmilestonIdJson) | **Get** /projects/api/v3/milestones/{milestonId}.json | Get a milestone by id.
[**GETProjectsApiV3ProjectsprojectIdMilestonesJson**](MilestonesApi.md#GETProjectsApiV3ProjectsprojectIdMilestonesJson) | **Get** /projects/api/v3/projects/{projectId}/milestones.json | Get milestones in a project



## GETProjectsApiV3MilestonesCsv

> GETProjectsApiV3MilestonesCsv(ctx).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).Status(status).SearchTerm(searchTerm).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).DueBeforeDate(dueBeforeDate).DueBefore(dueBefore).DueAfterDate(dueAfterDate).DueAfter(dueAfter).ProjectHealths(projectHealths).ProjectHealth(projectHealth).PageSize(pageSize).Page(page).ShowPrivate(showPrivate).ShowPercentageCompleted(showPercentageCompleted).ShowDeleted(showDeleted).ShowAttachedTasklists(showAttachedTasklists).SearchCompaniesTeams(searchCompaniesTeams).SearchAssignees(searchAssignees).Progress(progress).OnlyStarredProjects(onlyStarredProjects).OnlyMyProjects(onlyMyProjects).OnlyAdminProjects(onlyAdminProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeUpdate(includeUpdate).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTags(includeTags).IncludeProgress(includeProgress).IncludeCompleted(includeCompleted).IncludeCompanyUserIds(includeCompanyUserIds).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).Emoji(emoji).CountUnreadComments(countUnreadComments).Calendar(calendar).TagIds(tagIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).Ids(ids).FollowerIds(followerIds).FieldsUsers(fieldsUsers).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsMilestones(fieldsMilestones).FieldsCompanies(fieldsCompanies).CreatorIds(creatorIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()

Generate milestone report in CSV format



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
    updatedAfterDate := time.Now() // time.Time | filter by updated after date (deprecated, use updatedAfter) (optional)
    updatedAfter := time.Now() // time.Time | filter by updated after date (optional)
    status := "status_example" // string | filter by list of milestone status (optional)
    searchTerm := "searchTerm_example" // string | filter by milestone name and description (optional)
    reportFormat := "reportFormat_example" // string | define the format of the report (optional)
    projectStatuses := "projectStatuses_example" // string | filter by list of project status (optional)
    projectStatus := "projectStatus_example" // string | filter by project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "date")
    dueBeforeDate := time.Now() // string | filter by due before date (deprecated, use dueBefore) (optional)
    dueBefore := time.Now() // string | filter by due before date (optional)
    dueAfterDate := time.Now() // string | filter by due after date (deprecated, use dueAfter) (optional)
    dueAfter := time.Now() // string | filter by due after date (optional)
    projectHealths := int32(56) // int32 | filter by list of project health (optional)
    projectHealth := int32(56) // int32 | filter by list of project health (deprecated, use projectHealths) (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    showPrivate := true // bool | show private milestones (optional) (default to true)
    showPercentageCompleted := true // bool | show percentages completed (optional)
    showDeleted := true // bool | include deleted items (optional) (default to false)
    showAttachedTasklists := true // bool | show attached tasklists (optional)
    searchCompaniesTeams := true // bool | include companies and teams in the search term (optional)
    searchAssignees := true // bool | include assignees in the search (optional)
    progress := true // bool | include progress of each milestone (optional) (default to false)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    onlyMyProjects := true // bool | filter by projects that the logged-in user is related (optional) (default to false)
    onlyAdminProjects := true // bool | only include milestones from projects where the user is strictly a project admin. site admins have visibility to all projects. (optional)
    matchAllTags := true // bool | match all milestone tags (optional)
    matchAllProjectTags := true // bool | match all project tags (optional)
    isReportDownload := true // bool | generate a report document (optional)
    includeUpdate := true // bool | include milestones latest update action (optional)
    includeToday := true // bool | include today when milestone status is upcoming (optional) (default to true)
    includeTeamUserIds := true // bool | include team users to the responsible party ids (optional) (default to true)
    includeTags := true // bool | include tags in the reports (optional) (default to true)
    includeProgress := true // bool | include percentage of tasks completed for all linked tasklists (optional)
    includeCompleted := true // bool | optional to include completed milestones (optional)
    includeCompanyUserIds := true // bool | include company users to the responsible party ids (optional) (default to true)
    includeAssigneeTeams := true // bool | include teams related to the responsible user ids (optional)
    includeAssigneeCompanies := true // bool | include companies related to the responsible user ids (optional)
    includeArchivedProjects := true // bool | include archived projects (optional) (default to false)
    emoji := true // bool | convert emoji alias to unicode (optional) (default to true)
    countUnreadComments := true // bool | count unread comments (optional) (default to false)
    calendar := true // bool | enables the calendar style layout (optional) (default to false)
    tagIds := []int32{int32(123)} // []int32 | filter by milestone tag ids (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    ids := []int32{int32(123)} // []int32 | filter by ids (optional)
    followerIds := []int32{int32(123)} // []int32 | filter by followers' ids (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTasks := []string{"Inner_example"} // []string |  (optional)
    fieldsTasklists := []string{"Inner_example"} // []string |  (optional)
    fieldsTags := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsMilestones := []string{"Inner_example"} // []string |  (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)
    creatorIds := []int32{int32(123)} // []int32 | filter by creators' ids (optional)
    assignedToUserIds := []int32{int32(123)} // []int32 | filter by assigned user ids (optional)
    assignedToTeamIds := []int32{int32(123)} // []int32 | filter by assigned team ids (optional)
    assignedToCompanyIds := []int32{int32(123)} // []int32 | filter by assigned company ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MilestonesApi.GETProjectsApiV3MilestonesCsv(context.Background()).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).Status(status).SearchTerm(searchTerm).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).DueBeforeDate(dueBeforeDate).DueBefore(dueBefore).DueAfterDate(dueAfterDate).DueAfter(dueAfter).ProjectHealths(projectHealths).ProjectHealth(projectHealth).PageSize(pageSize).Page(page).ShowPrivate(showPrivate).ShowPercentageCompleted(showPercentageCompleted).ShowDeleted(showDeleted).ShowAttachedTasklists(showAttachedTasklists).SearchCompaniesTeams(searchCompaniesTeams).SearchAssignees(searchAssignees).Progress(progress).OnlyStarredProjects(onlyStarredProjects).OnlyMyProjects(onlyMyProjects).OnlyAdminProjects(onlyAdminProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeUpdate(includeUpdate).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTags(includeTags).IncludeProgress(includeProgress).IncludeCompleted(includeCompleted).IncludeCompanyUserIds(includeCompanyUserIds).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).Emoji(emoji).CountUnreadComments(countUnreadComments).Calendar(calendar).TagIds(tagIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).Ids(ids).FollowerIds(followerIds).FieldsUsers(fieldsUsers).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsMilestones(fieldsMilestones).FieldsCompanies(fieldsCompanies).CreatorIds(creatorIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MilestonesApi.GETProjectsApiV3MilestonesCsv``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3MilestonesCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfterDate** | **time.Time** | filter by updated after date (deprecated, use updatedAfter) | 
 **updatedAfter** | **time.Time** | filter by updated after date | 
 **status** | **string** | filter by list of milestone status | 
 **searchTerm** | **string** | filter by milestone name and description | 
 **reportFormat** | **string** | define the format of the report | 
 **projectStatuses** | **string** | filter by list of project status | 
 **projectStatus** | **string** | filter by project status | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;date&quot;]
 **dueBeforeDate** | **string** | filter by due before date (deprecated, use dueBefore) | 
 **dueBefore** | **string** | filter by due before date | 
 **dueAfterDate** | **string** | filter by due after date (deprecated, use dueAfter) | 
 **dueAfter** | **string** | filter by due after date | 
 **projectHealths** | **int32** | filter by list of project health | 
 **projectHealth** | **int32** | filter by list of project health (deprecated, use projectHealths) | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **showPrivate** | **bool** | show private milestones | [default to true]
 **showPercentageCompleted** | **bool** | show percentages completed | 
 **showDeleted** | **bool** | include deleted items | [default to false]
 **showAttachedTasklists** | **bool** | show attached tasklists | 
 **searchCompaniesTeams** | **bool** | include companies and teams in the search term | 
 **searchAssignees** | **bool** | include assignees in the search | 
 **progress** | **bool** | include progress of each milestone | [default to false]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **onlyMyProjects** | **bool** | filter by projects that the logged-in user is related | [default to false]
 **onlyAdminProjects** | **bool** | only include milestones from projects where the user is strictly a project admin. site admins have visibility to all projects. | 
 **matchAllTags** | **bool** | match all milestone tags | 
 **matchAllProjectTags** | **bool** | match all project tags | 
 **isReportDownload** | **bool** | generate a report document | 
 **includeUpdate** | **bool** | include milestones latest update action | 
 **includeToday** | **bool** | include today when milestone status is upcoming | [default to true]
 **includeTeamUserIds** | **bool** | include team users to the responsible party ids | [default to true]
 **includeTags** | **bool** | include tags in the reports | [default to true]
 **includeProgress** | **bool** | include percentage of tasks completed for all linked tasklists | 
 **includeCompleted** | **bool** | optional to include completed milestones | 
 **includeCompanyUserIds** | **bool** | include company users to the responsible party ids | [default to true]
 **includeAssigneeTeams** | **bool** | include teams related to the responsible user ids | 
 **includeAssigneeCompanies** | **bool** | include companies related to the responsible user ids | 
 **includeArchivedProjects** | **bool** | include archived projects | [default to false]
 **emoji** | **bool** | convert emoji alias to unicode | [default to true]
 **countUnreadComments** | **bool** | count unread comments | [default to false]
 **calendar** | **bool** | enables the calendar style layout | [default to false]
 **tagIds** | **[]int32** | filter by milestone tag ids | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include | 
 **ids** | **[]int32** | filter by ids | 
 **followerIds** | **[]int32** | filter by followers&#39; ids | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTasks** | **[]string** |  | 
 **fieldsTasklists** | **[]string** |  | 
 **fieldsTags** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsMilestones** | **[]string** |  | 
 **fieldsCompanies** | **[]string** |  | 
 **creatorIds** | **[]int32** | filter by creators&#39; ids | 
 **assignedToUserIds** | **[]int32** | filter by assigned user ids | 
 **assignedToTeamIds** | **[]int32** | filter by assigned team ids | 
 **assignedToCompanyIds** | **[]int32** | filter by assigned company ids | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3MilestonesHtml

> GETProjectsApiV3MilestonesHtml(ctx).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).Status(status).SearchTerm(searchTerm).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).DueBeforeDate(dueBeforeDate).DueBefore(dueBefore).DueAfterDate(dueAfterDate).DueAfter(dueAfter).ProjectHealths(projectHealths).ProjectHealth(projectHealth).PageSize(pageSize).Page(page).ShowPrivate(showPrivate).ShowPercentageCompleted(showPercentageCompleted).ShowDeleted(showDeleted).ShowAttachedTasklists(showAttachedTasklists).SearchCompaniesTeams(searchCompaniesTeams).SearchAssignees(searchAssignees).Progress(progress).OnlyStarredProjects(onlyStarredProjects).OnlyMyProjects(onlyMyProjects).OnlyAdminProjects(onlyAdminProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeUpdate(includeUpdate).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTags(includeTags).IncludeProgress(includeProgress).IncludeCompleted(includeCompleted).IncludeCompanyUserIds(includeCompanyUserIds).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).Emoji(emoji).CountUnreadComments(countUnreadComments).Calendar(calendar).TagIds(tagIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).Ids(ids).FollowerIds(followerIds).FieldsUsers(fieldsUsers).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsMilestones(fieldsMilestones).FieldsCompanies(fieldsCompanies).CreatorIds(creatorIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()

Generate milestone report in HTML format



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
    updatedAfterDate := time.Now() // time.Time | filter by updated after date (deprecated, use updatedAfter) (optional)
    updatedAfter := time.Now() // time.Time | filter by updated after date (optional)
    status := "status_example" // string | filter by list of milestone status (optional)
    searchTerm := "searchTerm_example" // string | filter by milestone name and description (optional)
    reportFormat := "reportFormat_example" // string | define the format of the report (optional)
    projectStatuses := "projectStatuses_example" // string | filter by list of project status (optional)
    projectStatus := "projectStatus_example" // string | filter by project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "date")
    dueBeforeDate := time.Now() // string | filter by due before date (deprecated, use dueBefore) (optional)
    dueBefore := time.Now() // string | filter by due before date (optional)
    dueAfterDate := time.Now() // string | filter by due after date (deprecated, use dueAfter) (optional)
    dueAfter := time.Now() // string | filter by due after date (optional)
    projectHealths := int32(56) // int32 | filter by list of project health (optional)
    projectHealth := int32(56) // int32 | filter by list of project health (deprecated, use projectHealths) (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    showPrivate := true // bool | show private milestones (optional) (default to true)
    showPercentageCompleted := true // bool | show percentages completed (optional)
    showDeleted := true // bool | include deleted items (optional) (default to false)
    showAttachedTasklists := true // bool | show attached tasklists (optional)
    searchCompaniesTeams := true // bool | include companies and teams in the search term (optional)
    searchAssignees := true // bool | include assignees in the search (optional)
    progress := true // bool | include progress of each milestone (optional) (default to false)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    onlyMyProjects := true // bool | filter by projects that the logged-in user is related (optional) (default to false)
    onlyAdminProjects := true // bool | only include milestones from projects where the user is strictly a project admin. site admins have visibility to all projects. (optional)
    matchAllTags := true // bool | match all milestone tags (optional)
    matchAllProjectTags := true // bool | match all project tags (optional)
    isReportDownload := true // bool | generate a report document (optional)
    includeUpdate := true // bool | include milestones latest update action (optional)
    includeToday := true // bool | include today when milestone status is upcoming (optional) (default to true)
    includeTeamUserIds := true // bool | include team users to the responsible party ids (optional) (default to true)
    includeTags := true // bool | include tags in the reports (optional) (default to true)
    includeProgress := true // bool | include percentage of tasks completed for all linked tasklists (optional)
    includeCompleted := true // bool | optional to include completed milestones (optional)
    includeCompanyUserIds := true // bool | include company users to the responsible party ids (optional) (default to true)
    includeAssigneeTeams := true // bool | include teams related to the responsible user ids (optional)
    includeAssigneeCompanies := true // bool | include companies related to the responsible user ids (optional)
    includeArchivedProjects := true // bool | include archived projects (optional) (default to false)
    emoji := true // bool | convert emoji alias to unicode (optional) (default to true)
    countUnreadComments := true // bool | count unread comments (optional) (default to false)
    calendar := true // bool | enables the calendar style layout (optional) (default to false)
    tagIds := []int32{int32(123)} // []int32 | filter by milestone tag ids (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    ids := []int32{int32(123)} // []int32 | filter by ids (optional)
    followerIds := []int32{int32(123)} // []int32 | filter by followers' ids (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTasks := []string{"Inner_example"} // []string |  (optional)
    fieldsTasklists := []string{"Inner_example"} // []string |  (optional)
    fieldsTags := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsMilestones := []string{"Inner_example"} // []string |  (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)
    creatorIds := []int32{int32(123)} // []int32 | filter by creators' ids (optional)
    assignedToUserIds := []int32{int32(123)} // []int32 | filter by assigned user ids (optional)
    assignedToTeamIds := []int32{int32(123)} // []int32 | filter by assigned team ids (optional)
    assignedToCompanyIds := []int32{int32(123)} // []int32 | filter by assigned company ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MilestonesApi.GETProjectsApiV3MilestonesHtml(context.Background()).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).Status(status).SearchTerm(searchTerm).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).DueBeforeDate(dueBeforeDate).DueBefore(dueBefore).DueAfterDate(dueAfterDate).DueAfter(dueAfter).ProjectHealths(projectHealths).ProjectHealth(projectHealth).PageSize(pageSize).Page(page).ShowPrivate(showPrivate).ShowPercentageCompleted(showPercentageCompleted).ShowDeleted(showDeleted).ShowAttachedTasklists(showAttachedTasklists).SearchCompaniesTeams(searchCompaniesTeams).SearchAssignees(searchAssignees).Progress(progress).OnlyStarredProjects(onlyStarredProjects).OnlyMyProjects(onlyMyProjects).OnlyAdminProjects(onlyAdminProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeUpdate(includeUpdate).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTags(includeTags).IncludeProgress(includeProgress).IncludeCompleted(includeCompleted).IncludeCompanyUserIds(includeCompanyUserIds).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).Emoji(emoji).CountUnreadComments(countUnreadComments).Calendar(calendar).TagIds(tagIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).Ids(ids).FollowerIds(followerIds).FieldsUsers(fieldsUsers).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsMilestones(fieldsMilestones).FieldsCompanies(fieldsCompanies).CreatorIds(creatorIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MilestonesApi.GETProjectsApiV3MilestonesHtml``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3MilestonesHtmlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfterDate** | **time.Time** | filter by updated after date (deprecated, use updatedAfter) | 
 **updatedAfter** | **time.Time** | filter by updated after date | 
 **status** | **string** | filter by list of milestone status | 
 **searchTerm** | **string** | filter by milestone name and description | 
 **reportFormat** | **string** | define the format of the report | 
 **projectStatuses** | **string** | filter by list of project status | 
 **projectStatus** | **string** | filter by project status | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;date&quot;]
 **dueBeforeDate** | **string** | filter by due before date (deprecated, use dueBefore) | 
 **dueBefore** | **string** | filter by due before date | 
 **dueAfterDate** | **string** | filter by due after date (deprecated, use dueAfter) | 
 **dueAfter** | **string** | filter by due after date | 
 **projectHealths** | **int32** | filter by list of project health | 
 **projectHealth** | **int32** | filter by list of project health (deprecated, use projectHealths) | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **showPrivate** | **bool** | show private milestones | [default to true]
 **showPercentageCompleted** | **bool** | show percentages completed | 
 **showDeleted** | **bool** | include deleted items | [default to false]
 **showAttachedTasklists** | **bool** | show attached tasklists | 
 **searchCompaniesTeams** | **bool** | include companies and teams in the search term | 
 **searchAssignees** | **bool** | include assignees in the search | 
 **progress** | **bool** | include progress of each milestone | [default to false]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **onlyMyProjects** | **bool** | filter by projects that the logged-in user is related | [default to false]
 **onlyAdminProjects** | **bool** | only include milestones from projects where the user is strictly a project admin. site admins have visibility to all projects. | 
 **matchAllTags** | **bool** | match all milestone tags | 
 **matchAllProjectTags** | **bool** | match all project tags | 
 **isReportDownload** | **bool** | generate a report document | 
 **includeUpdate** | **bool** | include milestones latest update action | 
 **includeToday** | **bool** | include today when milestone status is upcoming | [default to true]
 **includeTeamUserIds** | **bool** | include team users to the responsible party ids | [default to true]
 **includeTags** | **bool** | include tags in the reports | [default to true]
 **includeProgress** | **bool** | include percentage of tasks completed for all linked tasklists | 
 **includeCompleted** | **bool** | optional to include completed milestones | 
 **includeCompanyUserIds** | **bool** | include company users to the responsible party ids | [default to true]
 **includeAssigneeTeams** | **bool** | include teams related to the responsible user ids | 
 **includeAssigneeCompanies** | **bool** | include companies related to the responsible user ids | 
 **includeArchivedProjects** | **bool** | include archived projects | [default to false]
 **emoji** | **bool** | convert emoji alias to unicode | [default to true]
 **countUnreadComments** | **bool** | count unread comments | [default to false]
 **calendar** | **bool** | enables the calendar style layout | [default to false]
 **tagIds** | **[]int32** | filter by milestone tag ids | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include | 
 **ids** | **[]int32** | filter by ids | 
 **followerIds** | **[]int32** | filter by followers&#39; ids | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTasks** | **[]string** |  | 
 **fieldsTasklists** | **[]string** |  | 
 **fieldsTags** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsMilestones** | **[]string** |  | 
 **fieldsCompanies** | **[]string** |  | 
 **creatorIds** | **[]int32** | filter by creators&#39; ids | 
 **assignedToUserIds** | **[]int32** | filter by assigned user ids | 
 **assignedToTeamIds** | **[]int32** | filter by assigned team ids | 
 **assignedToCompanyIds** | **[]int32** | filter by assigned company ids | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3MilestonesJson

> MilestoneMilestonesResponse GETProjectsApiV3MilestonesJson(ctx).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).Status(status).SearchTerm(searchTerm).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).DueBeforeDate(dueBeforeDate).DueBefore(dueBefore).DueAfterDate(dueAfterDate).DueAfter(dueAfter).ProjectHealths(projectHealths).ProjectHealth(projectHealth).PageSize(pageSize).Page(page).ShowPrivate(showPrivate).ShowPercentageCompleted(showPercentageCompleted).ShowDeleted(showDeleted).ShowAttachedTasklists(showAttachedTasklists).SearchCompaniesTeams(searchCompaniesTeams).SearchAssignees(searchAssignees).Progress(progress).OnlyStarredProjects(onlyStarredProjects).OnlyMyProjects(onlyMyProjects).OnlyAdminProjects(onlyAdminProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeUpdate(includeUpdate).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTags(includeTags).IncludeProgress(includeProgress).IncludeCompleted(includeCompleted).IncludeCompanyUserIds(includeCompanyUserIds).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).Emoji(emoji).CountUnreadComments(countUnreadComments).Calendar(calendar).TagIds(tagIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).Ids(ids).FollowerIds(followerIds).FieldsUsers(fieldsUsers).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsMilestones(fieldsMilestones).FieldsCompanies(fieldsCompanies).CreatorIds(creatorIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()

Get all milestones



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
    updatedAfterDate := time.Now() // time.Time | filter by updated after date (deprecated, use updatedAfter) (optional)
    updatedAfter := time.Now() // time.Time | filter by updated after date (optional)
    status := "status_example" // string | filter by list of milestone status (optional)
    searchTerm := "searchTerm_example" // string | filter by milestone name and description (optional)
    reportFormat := "reportFormat_example" // string | define the format of the report (optional)
    projectStatuses := "projectStatuses_example" // string | filter by list of project status (optional)
    projectStatus := "projectStatus_example" // string | filter by project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "date")
    dueBeforeDate := time.Now() // string | filter by due before date (deprecated, use dueBefore) (optional)
    dueBefore := time.Now() // string | filter by due before date (optional)
    dueAfterDate := time.Now() // string | filter by due after date (deprecated, use dueAfter) (optional)
    dueAfter := time.Now() // string | filter by due after date (optional)
    projectHealths := int32(56) // int32 | filter by list of project health (optional)
    projectHealth := int32(56) // int32 | filter by list of project health (deprecated, use projectHealths) (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    showPrivate := true // bool | show private milestones (optional) (default to true)
    showPercentageCompleted := true // bool | show percentages completed (optional)
    showDeleted := true // bool | include deleted items (optional) (default to false)
    showAttachedTasklists := true // bool | show attached tasklists (optional)
    searchCompaniesTeams := true // bool | include companies and teams in the search term (optional)
    searchAssignees := true // bool | include assignees in the search (optional)
    progress := true // bool | include progress of each milestone (optional) (default to false)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    onlyMyProjects := true // bool | filter by projects that the logged-in user is related (optional) (default to false)
    onlyAdminProjects := true // bool | only include milestones from projects where the user is strictly a project admin. site admins have visibility to all projects. (optional)
    matchAllTags := true // bool | match all milestone tags (optional)
    matchAllProjectTags := true // bool | match all project tags (optional)
    isReportDownload := true // bool | generate a report document (optional)
    includeUpdate := true // bool | include milestones latest update action (optional)
    includeToday := true // bool | include today when milestone status is upcoming (optional) (default to true)
    includeTeamUserIds := true // bool | include team users to the responsible party ids (optional) (default to true)
    includeTags := true // bool | include tags in the reports (optional) (default to true)
    includeProgress := true // bool | include percentage of tasks completed for all linked tasklists (optional)
    includeCompleted := true // bool | optional to include completed milestones (optional)
    includeCompanyUserIds := true // bool | include company users to the responsible party ids (optional) (default to true)
    includeAssigneeTeams := true // bool | include teams related to the responsible user ids (optional)
    includeAssigneeCompanies := true // bool | include companies related to the responsible user ids (optional)
    includeArchivedProjects := true // bool | include archived projects (optional) (default to false)
    emoji := true // bool | convert emoji alias to unicode (optional) (default to true)
    countUnreadComments := true // bool | count unread comments (optional) (default to false)
    calendar := true // bool | enables the calendar style layout (optional) (default to false)
    tagIds := []int32{int32(123)} // []int32 | filter by milestone tag ids (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    ids := []int32{int32(123)} // []int32 | filter by ids (optional)
    followerIds := []int32{int32(123)} // []int32 | filter by followers' ids (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTasks := []string{"Inner_example"} // []string |  (optional)
    fieldsTasklists := []string{"Inner_example"} // []string |  (optional)
    fieldsTags := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsMilestones := []string{"Inner_example"} // []string |  (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)
    creatorIds := []int32{int32(123)} // []int32 | filter by creators' ids (optional)
    assignedToUserIds := []int32{int32(123)} // []int32 | filter by assigned user ids (optional)
    assignedToTeamIds := []int32{int32(123)} // []int32 | filter by assigned team ids (optional)
    assignedToCompanyIds := []int32{int32(123)} // []int32 | filter by assigned company ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MilestonesApi.GETProjectsApiV3MilestonesJson(context.Background()).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).Status(status).SearchTerm(searchTerm).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).DueBeforeDate(dueBeforeDate).DueBefore(dueBefore).DueAfterDate(dueAfterDate).DueAfter(dueAfter).ProjectHealths(projectHealths).ProjectHealth(projectHealth).PageSize(pageSize).Page(page).ShowPrivate(showPrivate).ShowPercentageCompleted(showPercentageCompleted).ShowDeleted(showDeleted).ShowAttachedTasklists(showAttachedTasklists).SearchCompaniesTeams(searchCompaniesTeams).SearchAssignees(searchAssignees).Progress(progress).OnlyStarredProjects(onlyStarredProjects).OnlyMyProjects(onlyMyProjects).OnlyAdminProjects(onlyAdminProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeUpdate(includeUpdate).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTags(includeTags).IncludeProgress(includeProgress).IncludeCompleted(includeCompleted).IncludeCompanyUserIds(includeCompanyUserIds).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).Emoji(emoji).CountUnreadComments(countUnreadComments).Calendar(calendar).TagIds(tagIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).Ids(ids).FollowerIds(followerIds).FieldsUsers(fieldsUsers).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsMilestones(fieldsMilestones).FieldsCompanies(fieldsCompanies).CreatorIds(creatorIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MilestonesApi.GETProjectsApiV3MilestonesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3MilestonesJson`: MilestoneMilestonesResponse
    fmt.Fprintf(os.Stdout, "Response from `MilestonesApi.GETProjectsApiV3MilestonesJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3MilestonesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfterDate** | **time.Time** | filter by updated after date (deprecated, use updatedAfter) | 
 **updatedAfter** | **time.Time** | filter by updated after date | 
 **status** | **string** | filter by list of milestone status | 
 **searchTerm** | **string** | filter by milestone name and description | 
 **reportFormat** | **string** | define the format of the report | 
 **projectStatuses** | **string** | filter by list of project status | 
 **projectStatus** | **string** | filter by project status | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;date&quot;]
 **dueBeforeDate** | **string** | filter by due before date (deprecated, use dueBefore) | 
 **dueBefore** | **string** | filter by due before date | 
 **dueAfterDate** | **string** | filter by due after date (deprecated, use dueAfter) | 
 **dueAfter** | **string** | filter by due after date | 
 **projectHealths** | **int32** | filter by list of project health | 
 **projectHealth** | **int32** | filter by list of project health (deprecated, use projectHealths) | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **showPrivate** | **bool** | show private milestones | [default to true]
 **showPercentageCompleted** | **bool** | show percentages completed | 
 **showDeleted** | **bool** | include deleted items | [default to false]
 **showAttachedTasklists** | **bool** | show attached tasklists | 
 **searchCompaniesTeams** | **bool** | include companies and teams in the search term | 
 **searchAssignees** | **bool** | include assignees in the search | 
 **progress** | **bool** | include progress of each milestone | [default to false]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **onlyMyProjects** | **bool** | filter by projects that the logged-in user is related | [default to false]
 **onlyAdminProjects** | **bool** | only include milestones from projects where the user is strictly a project admin. site admins have visibility to all projects. | 
 **matchAllTags** | **bool** | match all milestone tags | 
 **matchAllProjectTags** | **bool** | match all project tags | 
 **isReportDownload** | **bool** | generate a report document | 
 **includeUpdate** | **bool** | include milestones latest update action | 
 **includeToday** | **bool** | include today when milestone status is upcoming | [default to true]
 **includeTeamUserIds** | **bool** | include team users to the responsible party ids | [default to true]
 **includeTags** | **bool** | include tags in the reports | [default to true]
 **includeProgress** | **bool** | include percentage of tasks completed for all linked tasklists | 
 **includeCompleted** | **bool** | optional to include completed milestones | 
 **includeCompanyUserIds** | **bool** | include company users to the responsible party ids | [default to true]
 **includeAssigneeTeams** | **bool** | include teams related to the responsible user ids | 
 **includeAssigneeCompanies** | **bool** | include companies related to the responsible user ids | 
 **includeArchivedProjects** | **bool** | include archived projects | [default to false]
 **emoji** | **bool** | convert emoji alias to unicode | [default to true]
 **countUnreadComments** | **bool** | count unread comments | [default to false]
 **calendar** | **bool** | enables the calendar style layout | [default to false]
 **tagIds** | **[]int32** | filter by milestone tag ids | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include | 
 **ids** | **[]int32** | filter by ids | 
 **followerIds** | **[]int32** | filter by followers&#39; ids | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTasks** | **[]string** |  | 
 **fieldsTasklists** | **[]string** |  | 
 **fieldsTags** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsMilestones** | **[]string** |  | 
 **fieldsCompanies** | **[]string** |  | 
 **creatorIds** | **[]int32** | filter by creators&#39; ids | 
 **assignedToUserIds** | **[]int32** | filter by assigned user ids | 
 **assignedToTeamIds** | **[]int32** | filter by assigned team ids | 
 **assignedToCompanyIds** | **[]int32** | filter by assigned company ids | 

### Return type

[**MilestoneMilestonesResponse**](MilestoneMilestonesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3MilestonesMetricsDeadlinesJson

> DeadlineMilestoneMetricDeadlinesResponse GETProjectsApiV3MilestonesMetricsDeadlinesJson(ctx).StartDate(startDate).EndDate(endDate).Execute()

Get milestones by due date in a time range



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
    startDate := time.Now() // string |  (optional)
    endDate := time.Now() // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MilestonesApi.GETProjectsApiV3MilestonesMetricsDeadlinesJson(context.Background()).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MilestonesApi.GETProjectsApiV3MilestonesMetricsDeadlinesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3MilestonesMetricsDeadlinesJson`: DeadlineMilestoneMetricDeadlinesResponse
    fmt.Fprintf(os.Stdout, "Response from `MilestonesApi.GETProjectsApiV3MilestonesMetricsDeadlinesJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3MilestonesMetricsDeadlinesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 

### Return type

[**DeadlineMilestoneMetricDeadlinesResponse**](DeadlineMilestoneMetricDeadlinesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3MilestonesPdf

> GETProjectsApiV3MilestonesPdf(ctx).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).Status(status).SearchTerm(searchTerm).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).DueBeforeDate(dueBeforeDate).DueBefore(dueBefore).DueAfterDate(dueAfterDate).DueAfter(dueAfter).ProjectHealths(projectHealths).ProjectHealth(projectHealth).PageSize(pageSize).Page(page).ShowPrivate(showPrivate).ShowPercentageCompleted(showPercentageCompleted).ShowDeleted(showDeleted).ShowAttachedTasklists(showAttachedTasklists).SearchCompaniesTeams(searchCompaniesTeams).SearchAssignees(searchAssignees).Progress(progress).OnlyStarredProjects(onlyStarredProjects).OnlyMyProjects(onlyMyProjects).OnlyAdminProjects(onlyAdminProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeUpdate(includeUpdate).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTags(includeTags).IncludeProgress(includeProgress).IncludeCompleted(includeCompleted).IncludeCompanyUserIds(includeCompanyUserIds).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).Emoji(emoji).CountUnreadComments(countUnreadComments).Calendar(calendar).TagIds(tagIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).Ids(ids).FollowerIds(followerIds).FieldsUsers(fieldsUsers).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsMilestones(fieldsMilestones).FieldsCompanies(fieldsCompanies).CreatorIds(creatorIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()

Generate milestone report in PDF format



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
    updatedAfterDate := time.Now() // time.Time | filter by updated after date (deprecated, use updatedAfter) (optional)
    updatedAfter := time.Now() // time.Time | filter by updated after date (optional)
    status := "status_example" // string | filter by list of milestone status (optional)
    searchTerm := "searchTerm_example" // string | filter by milestone name and description (optional)
    reportFormat := "reportFormat_example" // string | define the format of the report (optional)
    projectStatuses := "projectStatuses_example" // string | filter by list of project status (optional)
    projectStatus := "projectStatus_example" // string | filter by project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "date")
    dueBeforeDate := time.Now() // string | filter by due before date (deprecated, use dueBefore) (optional)
    dueBefore := time.Now() // string | filter by due before date (optional)
    dueAfterDate := time.Now() // string | filter by due after date (deprecated, use dueAfter) (optional)
    dueAfter := time.Now() // string | filter by due after date (optional)
    projectHealths := int32(56) // int32 | filter by list of project health (optional)
    projectHealth := int32(56) // int32 | filter by list of project health (deprecated, use projectHealths) (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    showPrivate := true // bool | show private milestones (optional) (default to true)
    showPercentageCompleted := true // bool | show percentages completed (optional)
    showDeleted := true // bool | include deleted items (optional) (default to false)
    showAttachedTasklists := true // bool | show attached tasklists (optional)
    searchCompaniesTeams := true // bool | include companies and teams in the search term (optional)
    searchAssignees := true // bool | include assignees in the search (optional)
    progress := true // bool | include progress of each milestone (optional) (default to false)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    onlyMyProjects := true // bool | filter by projects that the logged-in user is related (optional) (default to false)
    onlyAdminProjects := true // bool | only include milestones from projects where the user is strictly a project admin. site admins have visibility to all projects. (optional)
    matchAllTags := true // bool | match all milestone tags (optional)
    matchAllProjectTags := true // bool | match all project tags (optional)
    isReportDownload := true // bool | generate a report document (optional)
    includeUpdate := true // bool | include milestones latest update action (optional)
    includeToday := true // bool | include today when milestone status is upcoming (optional) (default to true)
    includeTeamUserIds := true // bool | include team users to the responsible party ids (optional) (default to true)
    includeTags := true // bool | include tags in the reports (optional) (default to true)
    includeProgress := true // bool | include percentage of tasks completed for all linked tasklists (optional)
    includeCompleted := true // bool | optional to include completed milestones (optional)
    includeCompanyUserIds := true // bool | include company users to the responsible party ids (optional) (default to true)
    includeAssigneeTeams := true // bool | include teams related to the responsible user ids (optional)
    includeAssigneeCompanies := true // bool | include companies related to the responsible user ids (optional)
    includeArchivedProjects := true // bool | include archived projects (optional) (default to false)
    emoji := true // bool | convert emoji alias to unicode (optional) (default to true)
    countUnreadComments := true // bool | count unread comments (optional) (default to false)
    calendar := true // bool | enables the calendar style layout (optional) (default to false)
    tagIds := []int32{int32(123)} // []int32 | filter by milestone tag ids (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    ids := []int32{int32(123)} // []int32 | filter by ids (optional)
    followerIds := []int32{int32(123)} // []int32 | filter by followers' ids (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTasks := []string{"Inner_example"} // []string |  (optional)
    fieldsTasklists := []string{"Inner_example"} // []string |  (optional)
    fieldsTags := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsMilestones := []string{"Inner_example"} // []string |  (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)
    creatorIds := []int32{int32(123)} // []int32 | filter by creators' ids (optional)
    assignedToUserIds := []int32{int32(123)} // []int32 | filter by assigned user ids (optional)
    assignedToTeamIds := []int32{int32(123)} // []int32 | filter by assigned team ids (optional)
    assignedToCompanyIds := []int32{int32(123)} // []int32 | filter by assigned company ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MilestonesApi.GETProjectsApiV3MilestonesPdf(context.Background()).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).Status(status).SearchTerm(searchTerm).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).DueBeforeDate(dueBeforeDate).DueBefore(dueBefore).DueAfterDate(dueAfterDate).DueAfter(dueAfter).ProjectHealths(projectHealths).ProjectHealth(projectHealth).PageSize(pageSize).Page(page).ShowPrivate(showPrivate).ShowPercentageCompleted(showPercentageCompleted).ShowDeleted(showDeleted).ShowAttachedTasklists(showAttachedTasklists).SearchCompaniesTeams(searchCompaniesTeams).SearchAssignees(searchAssignees).Progress(progress).OnlyStarredProjects(onlyStarredProjects).OnlyMyProjects(onlyMyProjects).OnlyAdminProjects(onlyAdminProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeUpdate(includeUpdate).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTags(includeTags).IncludeProgress(includeProgress).IncludeCompleted(includeCompleted).IncludeCompanyUserIds(includeCompanyUserIds).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).Emoji(emoji).CountUnreadComments(countUnreadComments).Calendar(calendar).TagIds(tagIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).Ids(ids).FollowerIds(followerIds).FieldsUsers(fieldsUsers).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsMilestones(fieldsMilestones).FieldsCompanies(fieldsCompanies).CreatorIds(creatorIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MilestonesApi.GETProjectsApiV3MilestonesPdf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3MilestonesPdfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfterDate** | **time.Time** | filter by updated after date (deprecated, use updatedAfter) | 
 **updatedAfter** | **time.Time** | filter by updated after date | 
 **status** | **string** | filter by list of milestone status | 
 **searchTerm** | **string** | filter by milestone name and description | 
 **reportFormat** | **string** | define the format of the report | 
 **projectStatuses** | **string** | filter by list of project status | 
 **projectStatus** | **string** | filter by project status | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;date&quot;]
 **dueBeforeDate** | **string** | filter by due before date (deprecated, use dueBefore) | 
 **dueBefore** | **string** | filter by due before date | 
 **dueAfterDate** | **string** | filter by due after date (deprecated, use dueAfter) | 
 **dueAfter** | **string** | filter by due after date | 
 **projectHealths** | **int32** | filter by list of project health | 
 **projectHealth** | **int32** | filter by list of project health (deprecated, use projectHealths) | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **showPrivate** | **bool** | show private milestones | [default to true]
 **showPercentageCompleted** | **bool** | show percentages completed | 
 **showDeleted** | **bool** | include deleted items | [default to false]
 **showAttachedTasklists** | **bool** | show attached tasklists | 
 **searchCompaniesTeams** | **bool** | include companies and teams in the search term | 
 **searchAssignees** | **bool** | include assignees in the search | 
 **progress** | **bool** | include progress of each milestone | [default to false]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **onlyMyProjects** | **bool** | filter by projects that the logged-in user is related | [default to false]
 **onlyAdminProjects** | **bool** | only include milestones from projects where the user is strictly a project admin. site admins have visibility to all projects. | 
 **matchAllTags** | **bool** | match all milestone tags | 
 **matchAllProjectTags** | **bool** | match all project tags | 
 **isReportDownload** | **bool** | generate a report document | 
 **includeUpdate** | **bool** | include milestones latest update action | 
 **includeToday** | **bool** | include today when milestone status is upcoming | [default to true]
 **includeTeamUserIds** | **bool** | include team users to the responsible party ids | [default to true]
 **includeTags** | **bool** | include tags in the reports | [default to true]
 **includeProgress** | **bool** | include percentage of tasks completed for all linked tasklists | 
 **includeCompleted** | **bool** | optional to include completed milestones | 
 **includeCompanyUserIds** | **bool** | include company users to the responsible party ids | [default to true]
 **includeAssigneeTeams** | **bool** | include teams related to the responsible user ids | 
 **includeAssigneeCompanies** | **bool** | include companies related to the responsible user ids | 
 **includeArchivedProjects** | **bool** | include archived projects | [default to false]
 **emoji** | **bool** | convert emoji alias to unicode | [default to true]
 **countUnreadComments** | **bool** | count unread comments | [default to false]
 **calendar** | **bool** | enables the calendar style layout | [default to false]
 **tagIds** | **[]int32** | filter by milestone tag ids | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include | 
 **ids** | **[]int32** | filter by ids | 
 **followerIds** | **[]int32** | filter by followers&#39; ids | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTasks** | **[]string** |  | 
 **fieldsTasklists** | **[]string** |  | 
 **fieldsTags** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsMilestones** | **[]string** |  | 
 **fieldsCompanies** | **[]string** |  | 
 **creatorIds** | **[]int32** | filter by creators&#39; ids | 
 **assignedToUserIds** | **[]int32** | filter by assigned user ids | 
 **assignedToTeamIds** | **[]int32** | filter by assigned team ids | 
 **assignedToCompanyIds** | **[]int32** | filter by assigned company ids | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3MilestonesXlsx

> GETProjectsApiV3MilestonesXlsx(ctx).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).Status(status).SearchTerm(searchTerm).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).DueBeforeDate(dueBeforeDate).DueBefore(dueBefore).DueAfterDate(dueAfterDate).DueAfter(dueAfter).ProjectHealths(projectHealths).ProjectHealth(projectHealth).PageSize(pageSize).Page(page).ShowPrivate(showPrivate).ShowPercentageCompleted(showPercentageCompleted).ShowDeleted(showDeleted).ShowAttachedTasklists(showAttachedTasklists).SearchCompaniesTeams(searchCompaniesTeams).SearchAssignees(searchAssignees).Progress(progress).OnlyStarredProjects(onlyStarredProjects).OnlyMyProjects(onlyMyProjects).OnlyAdminProjects(onlyAdminProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeUpdate(includeUpdate).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTags(includeTags).IncludeProgress(includeProgress).IncludeCompleted(includeCompleted).IncludeCompanyUserIds(includeCompanyUserIds).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).Emoji(emoji).CountUnreadComments(countUnreadComments).Calendar(calendar).TagIds(tagIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).Ids(ids).FollowerIds(followerIds).FieldsUsers(fieldsUsers).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsMilestones(fieldsMilestones).FieldsCompanies(fieldsCompanies).CreatorIds(creatorIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()

Generate milestone report in XLSX format



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
    updatedAfterDate := time.Now() // time.Time | filter by updated after date (deprecated, use updatedAfter) (optional)
    updatedAfter := time.Now() // time.Time | filter by updated after date (optional)
    status := "status_example" // string | filter by list of milestone status (optional)
    searchTerm := "searchTerm_example" // string | filter by milestone name and description (optional)
    reportFormat := "reportFormat_example" // string | define the format of the report (optional)
    projectStatuses := "projectStatuses_example" // string | filter by list of project status (optional)
    projectStatus := "projectStatus_example" // string | filter by project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "date")
    dueBeforeDate := time.Now() // string | filter by due before date (deprecated, use dueBefore) (optional)
    dueBefore := time.Now() // string | filter by due before date (optional)
    dueAfterDate := time.Now() // string | filter by due after date (deprecated, use dueAfter) (optional)
    dueAfter := time.Now() // string | filter by due after date (optional)
    projectHealths := int32(56) // int32 | filter by list of project health (optional)
    projectHealth := int32(56) // int32 | filter by list of project health (deprecated, use projectHealths) (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    showPrivate := true // bool | show private milestones (optional) (default to true)
    showPercentageCompleted := true // bool | show percentages completed (optional)
    showDeleted := true // bool | include deleted items (optional) (default to false)
    showAttachedTasklists := true // bool | show attached tasklists (optional)
    searchCompaniesTeams := true // bool | include companies and teams in the search term (optional)
    searchAssignees := true // bool | include assignees in the search (optional)
    progress := true // bool | include progress of each milestone (optional) (default to false)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    onlyMyProjects := true // bool | filter by projects that the logged-in user is related (optional) (default to false)
    onlyAdminProjects := true // bool | only include milestones from projects where the user is strictly a project admin. site admins have visibility to all projects. (optional)
    matchAllTags := true // bool | match all milestone tags (optional)
    matchAllProjectTags := true // bool | match all project tags (optional)
    isReportDownload := true // bool | generate a report document (optional)
    includeUpdate := true // bool | include milestones latest update action (optional)
    includeToday := true // bool | include today when milestone status is upcoming (optional) (default to true)
    includeTeamUserIds := true // bool | include team users to the responsible party ids (optional) (default to true)
    includeTags := true // bool | include tags in the reports (optional) (default to true)
    includeProgress := true // bool | include percentage of tasks completed for all linked tasklists (optional)
    includeCompleted := true // bool | optional to include completed milestones (optional)
    includeCompanyUserIds := true // bool | include company users to the responsible party ids (optional) (default to true)
    includeAssigneeTeams := true // bool | include teams related to the responsible user ids (optional)
    includeAssigneeCompanies := true // bool | include companies related to the responsible user ids (optional)
    includeArchivedProjects := true // bool | include archived projects (optional) (default to false)
    emoji := true // bool | convert emoji alias to unicode (optional) (default to true)
    countUnreadComments := true // bool | count unread comments (optional) (default to false)
    calendar := true // bool | enables the calendar style layout (optional) (default to false)
    tagIds := []int32{int32(123)} // []int32 | filter by milestone tag ids (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    ids := []int32{int32(123)} // []int32 | filter by ids (optional)
    followerIds := []int32{int32(123)} // []int32 | filter by followers' ids (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTasks := []string{"Inner_example"} // []string |  (optional)
    fieldsTasklists := []string{"Inner_example"} // []string |  (optional)
    fieldsTags := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsMilestones := []string{"Inner_example"} // []string |  (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)
    creatorIds := []int32{int32(123)} // []int32 | filter by creators' ids (optional)
    assignedToUserIds := []int32{int32(123)} // []int32 | filter by assigned user ids (optional)
    assignedToTeamIds := []int32{int32(123)} // []int32 | filter by assigned team ids (optional)
    assignedToCompanyIds := []int32{int32(123)} // []int32 | filter by assigned company ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MilestonesApi.GETProjectsApiV3MilestonesXlsx(context.Background()).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).Status(status).SearchTerm(searchTerm).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).DueBeforeDate(dueBeforeDate).DueBefore(dueBefore).DueAfterDate(dueAfterDate).DueAfter(dueAfter).ProjectHealths(projectHealths).ProjectHealth(projectHealth).PageSize(pageSize).Page(page).ShowPrivate(showPrivate).ShowPercentageCompleted(showPercentageCompleted).ShowDeleted(showDeleted).ShowAttachedTasklists(showAttachedTasklists).SearchCompaniesTeams(searchCompaniesTeams).SearchAssignees(searchAssignees).Progress(progress).OnlyStarredProjects(onlyStarredProjects).OnlyMyProjects(onlyMyProjects).OnlyAdminProjects(onlyAdminProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeUpdate(includeUpdate).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTags(includeTags).IncludeProgress(includeProgress).IncludeCompleted(includeCompleted).IncludeCompanyUserIds(includeCompanyUserIds).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).Emoji(emoji).CountUnreadComments(countUnreadComments).Calendar(calendar).TagIds(tagIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).Ids(ids).FollowerIds(followerIds).FieldsUsers(fieldsUsers).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsMilestones(fieldsMilestones).FieldsCompanies(fieldsCompanies).CreatorIds(creatorIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MilestonesApi.GETProjectsApiV3MilestonesXlsx``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3MilestonesXlsxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfterDate** | **time.Time** | filter by updated after date (deprecated, use updatedAfter) | 
 **updatedAfter** | **time.Time** | filter by updated after date | 
 **status** | **string** | filter by list of milestone status | 
 **searchTerm** | **string** | filter by milestone name and description | 
 **reportFormat** | **string** | define the format of the report | 
 **projectStatuses** | **string** | filter by list of project status | 
 **projectStatus** | **string** | filter by project status | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;date&quot;]
 **dueBeforeDate** | **string** | filter by due before date (deprecated, use dueBefore) | 
 **dueBefore** | **string** | filter by due before date | 
 **dueAfterDate** | **string** | filter by due after date (deprecated, use dueAfter) | 
 **dueAfter** | **string** | filter by due after date | 
 **projectHealths** | **int32** | filter by list of project health | 
 **projectHealth** | **int32** | filter by list of project health (deprecated, use projectHealths) | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **showPrivate** | **bool** | show private milestones | [default to true]
 **showPercentageCompleted** | **bool** | show percentages completed | 
 **showDeleted** | **bool** | include deleted items | [default to false]
 **showAttachedTasklists** | **bool** | show attached tasklists | 
 **searchCompaniesTeams** | **bool** | include companies and teams in the search term | 
 **searchAssignees** | **bool** | include assignees in the search | 
 **progress** | **bool** | include progress of each milestone | [default to false]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **onlyMyProjects** | **bool** | filter by projects that the logged-in user is related | [default to false]
 **onlyAdminProjects** | **bool** | only include milestones from projects where the user is strictly a project admin. site admins have visibility to all projects. | 
 **matchAllTags** | **bool** | match all milestone tags | 
 **matchAllProjectTags** | **bool** | match all project tags | 
 **isReportDownload** | **bool** | generate a report document | 
 **includeUpdate** | **bool** | include milestones latest update action | 
 **includeToday** | **bool** | include today when milestone status is upcoming | [default to true]
 **includeTeamUserIds** | **bool** | include team users to the responsible party ids | [default to true]
 **includeTags** | **bool** | include tags in the reports | [default to true]
 **includeProgress** | **bool** | include percentage of tasks completed for all linked tasklists | 
 **includeCompleted** | **bool** | optional to include completed milestones | 
 **includeCompanyUserIds** | **bool** | include company users to the responsible party ids | [default to true]
 **includeAssigneeTeams** | **bool** | include teams related to the responsible user ids | 
 **includeAssigneeCompanies** | **bool** | include companies related to the responsible user ids | 
 **includeArchivedProjects** | **bool** | include archived projects | [default to false]
 **emoji** | **bool** | convert emoji alias to unicode | [default to true]
 **countUnreadComments** | **bool** | count unread comments | [default to false]
 **calendar** | **bool** | enables the calendar style layout | [default to false]
 **tagIds** | **[]int32** | filter by milestone tag ids | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include | 
 **ids** | **[]int32** | filter by ids | 
 **followerIds** | **[]int32** | filter by followers&#39; ids | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTasks** | **[]string** |  | 
 **fieldsTasklists** | **[]string** |  | 
 **fieldsTags** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsMilestones** | **[]string** |  | 
 **fieldsCompanies** | **[]string** |  | 
 **creatorIds** | **[]int32** | filter by creators&#39; ids | 
 **assignedToUserIds** | **[]int32** | filter by assigned user ids | 
 **assignedToTeamIds** | **[]int32** | filter by assigned team ids | 
 **assignedToCompanyIds** | **[]int32** | filter by assigned company ids | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3MilestonesmilestonIdJson

> MilestoneResponse GETProjectsApiV3MilestonesmilestonIdJson(ctx, milestonId).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).Status(status).SearchTerm(searchTerm).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).DueBeforeDate(dueBeforeDate).DueBefore(dueBefore).DueAfterDate(dueAfterDate).DueAfter(dueAfter).ProjectHealths(projectHealths).ProjectHealth(projectHealth).PageSize(pageSize).Page(page).ShowPrivate(showPrivate).ShowPercentageCompleted(showPercentageCompleted).ShowDeleted(showDeleted).ShowAttachedTasklists(showAttachedTasklists).SearchCompaniesTeams(searchCompaniesTeams).SearchAssignees(searchAssignees).Progress(progress).OnlyStarredProjects(onlyStarredProjects).OnlyMyProjects(onlyMyProjects).OnlyAdminProjects(onlyAdminProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeUpdate(includeUpdate).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTags(includeTags).IncludeProgress(includeProgress).IncludeCompleted(includeCompleted).IncludeCompanyUserIds(includeCompanyUserIds).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).Emoji(emoji).CountUnreadComments(countUnreadComments).Calendar(calendar).TagIds(tagIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).Ids(ids).FollowerIds(followerIds).FieldsUsers(fieldsUsers).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsMilestones(fieldsMilestones).FieldsCompanies(fieldsCompanies).CreatorIds(creatorIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()

Get a milestone by id.



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
    milestonId := int32(56) // int32 | 
    updatedAfterDate := time.Now() // time.Time | filter by updated after date (deprecated, use updatedAfter) (optional)
    updatedAfter := time.Now() // time.Time | filter by updated after date (optional)
    status := "status_example" // string | filter by list of milestone status (optional)
    searchTerm := "searchTerm_example" // string | filter by milestone name and description (optional)
    reportFormat := "reportFormat_example" // string | define the format of the report (optional)
    projectStatuses := "projectStatuses_example" // string | filter by list of project status (optional)
    projectStatus := "projectStatus_example" // string | filter by project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "date")
    dueBeforeDate := time.Now() // string | filter by due before date (deprecated, use dueBefore) (optional)
    dueBefore := time.Now() // string | filter by due before date (optional)
    dueAfterDate := time.Now() // string | filter by due after date (deprecated, use dueAfter) (optional)
    dueAfter := time.Now() // string | filter by due after date (optional)
    projectHealths := int32(56) // int32 | filter by list of project health (optional)
    projectHealth := int32(56) // int32 | filter by list of project health (deprecated, use projectHealths) (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    showPrivate := true // bool | show private milestones (optional) (default to true)
    showPercentageCompleted := true // bool | show percentages completed (optional)
    showDeleted := true // bool | include deleted items (optional) (default to false)
    showAttachedTasklists := true // bool | show attached tasklists (optional)
    searchCompaniesTeams := true // bool | include companies and teams in the search term (optional)
    searchAssignees := true // bool | include assignees in the search (optional)
    progress := true // bool | include progress of each milestone (optional) (default to false)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    onlyMyProjects := true // bool | filter by projects that the logged-in user is related (optional) (default to false)
    onlyAdminProjects := true // bool | only include milestones from projects where the user is strictly a project admin. site admins have visibility to all projects. (optional)
    matchAllTags := true // bool | match all milestone tags (optional)
    matchAllProjectTags := true // bool | match all project tags (optional)
    isReportDownload := true // bool | generate a report document (optional)
    includeUpdate := true // bool | include milestones latest update action (optional)
    includeToday := true // bool | include today when milestone status is upcoming (optional) (default to true)
    includeTeamUserIds := true // bool | include team users to the responsible party ids (optional) (default to true)
    includeTags := true // bool | include tags in the reports (optional) (default to true)
    includeProgress := true // bool | include percentage of tasks completed for all linked tasklists (optional)
    includeCompleted := true // bool | optional to include completed milestones (optional)
    includeCompanyUserIds := true // bool | include company users to the responsible party ids (optional) (default to true)
    includeAssigneeTeams := true // bool | include teams related to the responsible user ids (optional)
    includeAssigneeCompanies := true // bool | include companies related to the responsible user ids (optional)
    includeArchivedProjects := true // bool | include archived projects (optional) (default to false)
    emoji := true // bool | convert emoji alias to unicode (optional) (default to true)
    countUnreadComments := true // bool | count unread comments (optional) (default to false)
    calendar := true // bool | enables the calendar style layout (optional) (default to false)
    tagIds := []int32{int32(123)} // []int32 | filter by milestone tag ids (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    ids := []int32{int32(123)} // []int32 | filter by ids (optional)
    followerIds := []int32{int32(123)} // []int32 | filter by followers' ids (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTasks := []string{"Inner_example"} // []string |  (optional)
    fieldsTasklists := []string{"Inner_example"} // []string |  (optional)
    fieldsTags := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsMilestones := []string{"Inner_example"} // []string |  (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)
    creatorIds := []int32{int32(123)} // []int32 | filter by creators' ids (optional)
    assignedToUserIds := []int32{int32(123)} // []int32 | filter by assigned user ids (optional)
    assignedToTeamIds := []int32{int32(123)} // []int32 | filter by assigned team ids (optional)
    assignedToCompanyIds := []int32{int32(123)} // []int32 | filter by assigned company ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MilestonesApi.GETProjectsApiV3MilestonesmilestonIdJson(context.Background(), milestonId).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).Status(status).SearchTerm(searchTerm).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).DueBeforeDate(dueBeforeDate).DueBefore(dueBefore).DueAfterDate(dueAfterDate).DueAfter(dueAfter).ProjectHealths(projectHealths).ProjectHealth(projectHealth).PageSize(pageSize).Page(page).ShowPrivate(showPrivate).ShowPercentageCompleted(showPercentageCompleted).ShowDeleted(showDeleted).ShowAttachedTasklists(showAttachedTasklists).SearchCompaniesTeams(searchCompaniesTeams).SearchAssignees(searchAssignees).Progress(progress).OnlyStarredProjects(onlyStarredProjects).OnlyMyProjects(onlyMyProjects).OnlyAdminProjects(onlyAdminProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeUpdate(includeUpdate).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTags(includeTags).IncludeProgress(includeProgress).IncludeCompleted(includeCompleted).IncludeCompanyUserIds(includeCompanyUserIds).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).Emoji(emoji).CountUnreadComments(countUnreadComments).Calendar(calendar).TagIds(tagIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).Ids(ids).FollowerIds(followerIds).FieldsUsers(fieldsUsers).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsMilestones(fieldsMilestones).FieldsCompanies(fieldsCompanies).CreatorIds(creatorIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MilestonesApi.GETProjectsApiV3MilestonesmilestonIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3MilestonesmilestonIdJson`: MilestoneResponse
    fmt.Fprintf(os.Stdout, "Response from `MilestonesApi.GETProjectsApiV3MilestonesmilestonIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**milestonId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3MilestonesmilestonIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedAfterDate** | **time.Time** | filter by updated after date (deprecated, use updatedAfter) | 
 **updatedAfter** | **time.Time** | filter by updated after date | 
 **status** | **string** | filter by list of milestone status | 
 **searchTerm** | **string** | filter by milestone name and description | 
 **reportFormat** | **string** | define the format of the report | 
 **projectStatuses** | **string** | filter by list of project status | 
 **projectStatus** | **string** | filter by project status | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;date&quot;]
 **dueBeforeDate** | **string** | filter by due before date (deprecated, use dueBefore) | 
 **dueBefore** | **string** | filter by due before date | 
 **dueAfterDate** | **string** | filter by due after date (deprecated, use dueAfter) | 
 **dueAfter** | **string** | filter by due after date | 
 **projectHealths** | **int32** | filter by list of project health | 
 **projectHealth** | **int32** | filter by list of project health (deprecated, use projectHealths) | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **showPrivate** | **bool** | show private milestones | [default to true]
 **showPercentageCompleted** | **bool** | show percentages completed | 
 **showDeleted** | **bool** | include deleted items | [default to false]
 **showAttachedTasklists** | **bool** | show attached tasklists | 
 **searchCompaniesTeams** | **bool** | include companies and teams in the search term | 
 **searchAssignees** | **bool** | include assignees in the search | 
 **progress** | **bool** | include progress of each milestone | [default to false]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **onlyMyProjects** | **bool** | filter by projects that the logged-in user is related | [default to false]
 **onlyAdminProjects** | **bool** | only include milestones from projects where the user is strictly a project admin. site admins have visibility to all projects. | 
 **matchAllTags** | **bool** | match all milestone tags | 
 **matchAllProjectTags** | **bool** | match all project tags | 
 **isReportDownload** | **bool** | generate a report document | 
 **includeUpdate** | **bool** | include milestones latest update action | 
 **includeToday** | **bool** | include today when milestone status is upcoming | [default to true]
 **includeTeamUserIds** | **bool** | include team users to the responsible party ids | [default to true]
 **includeTags** | **bool** | include tags in the reports | [default to true]
 **includeProgress** | **bool** | include percentage of tasks completed for all linked tasklists | 
 **includeCompleted** | **bool** | optional to include completed milestones | 
 **includeCompanyUserIds** | **bool** | include company users to the responsible party ids | [default to true]
 **includeAssigneeTeams** | **bool** | include teams related to the responsible user ids | 
 **includeAssigneeCompanies** | **bool** | include companies related to the responsible user ids | 
 **includeArchivedProjects** | **bool** | include archived projects | [default to false]
 **emoji** | **bool** | convert emoji alias to unicode | [default to true]
 **countUnreadComments** | **bool** | count unread comments | [default to false]
 **calendar** | **bool** | enables the calendar style layout | [default to false]
 **tagIds** | **[]int32** | filter by milestone tag ids | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include | 
 **ids** | **[]int32** | filter by ids | 
 **followerIds** | **[]int32** | filter by followers&#39; ids | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTasks** | **[]string** |  | 
 **fieldsTasklists** | **[]string** |  | 
 **fieldsTags** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsMilestones** | **[]string** |  | 
 **fieldsCompanies** | **[]string** |  | 
 **creatorIds** | **[]int32** | filter by creators&#39; ids | 
 **assignedToUserIds** | **[]int32** | filter by assigned user ids | 
 **assignedToTeamIds** | **[]int32** | filter by assigned team ids | 
 **assignedToCompanyIds** | **[]int32** | filter by assigned company ids | 

### Return type

[**MilestoneResponse**](MilestoneResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3ProjectsprojectIdMilestonesJson

> MilestoneMilestonesResponse GETProjectsApiV3ProjectsprojectIdMilestonesJson(ctx, projectId).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).Status(status).SearchTerm(searchTerm).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).DueBeforeDate(dueBeforeDate).DueBefore(dueBefore).DueAfterDate(dueAfterDate).DueAfter(dueAfter).ProjectHealths(projectHealths).ProjectHealth(projectHealth).PageSize(pageSize).Page(page).ShowPrivate(showPrivate).ShowPercentageCompleted(showPercentageCompleted).ShowDeleted(showDeleted).ShowAttachedTasklists(showAttachedTasklists).SearchCompaniesTeams(searchCompaniesTeams).SearchAssignees(searchAssignees).Progress(progress).OnlyStarredProjects(onlyStarredProjects).OnlyMyProjects(onlyMyProjects).OnlyAdminProjects(onlyAdminProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeUpdate(includeUpdate).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTags(includeTags).IncludeProgress(includeProgress).IncludeCompleted(includeCompleted).IncludeCompanyUserIds(includeCompanyUserIds).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).Emoji(emoji).CountUnreadComments(countUnreadComments).Calendar(calendar).TagIds(tagIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).Ids(ids).FollowerIds(followerIds).FieldsUsers(fieldsUsers).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsMilestones(fieldsMilestones).FieldsCompanies(fieldsCompanies).CreatorIds(creatorIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()

Get milestones in a project



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
    updatedAfterDate := time.Now() // time.Time | filter by updated after date (deprecated, use updatedAfter) (optional)
    updatedAfter := time.Now() // time.Time | filter by updated after date (optional)
    status := "status_example" // string | filter by list of milestone status (optional)
    searchTerm := "searchTerm_example" // string | filter by milestone name and description (optional)
    reportFormat := "reportFormat_example" // string | define the format of the report (optional)
    projectStatuses := "projectStatuses_example" // string | filter by list of project status (optional)
    projectStatus := "projectStatus_example" // string | filter by project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "date")
    dueBeforeDate := time.Now() // string | filter by due before date (deprecated, use dueBefore) (optional)
    dueBefore := time.Now() // string | filter by due before date (optional)
    dueAfterDate := time.Now() // string | filter by due after date (deprecated, use dueAfter) (optional)
    dueAfter := time.Now() // string | filter by due after date (optional)
    projectHealths := int32(56) // int32 | filter by list of project health (optional)
    projectHealth := int32(56) // int32 | filter by list of project health (deprecated, use projectHealths) (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    showPrivate := true // bool | show private milestones (optional) (default to true)
    showPercentageCompleted := true // bool | show percentages completed (optional)
    showDeleted := true // bool | include deleted items (optional) (default to false)
    showAttachedTasklists := true // bool | show attached tasklists (optional)
    searchCompaniesTeams := true // bool | include companies and teams in the search term (optional)
    searchAssignees := true // bool | include assignees in the search (optional)
    progress := true // bool | include progress of each milestone (optional) (default to false)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    onlyMyProjects := true // bool | filter by projects that the logged-in user is related (optional) (default to false)
    onlyAdminProjects := true // bool | only include milestones from projects where the user is strictly a project admin. site admins have visibility to all projects. (optional)
    matchAllTags := true // bool | match all milestone tags (optional)
    matchAllProjectTags := true // bool | match all project tags (optional)
    isReportDownload := true // bool | generate a report document (optional)
    includeUpdate := true // bool | include milestones latest update action (optional)
    includeToday := true // bool | include today when milestone status is upcoming (optional) (default to true)
    includeTeamUserIds := true // bool | include team users to the responsible party ids (optional) (default to true)
    includeTags := true // bool | include tags in the reports (optional) (default to true)
    includeProgress := true // bool | include percentage of tasks completed for all linked tasklists (optional)
    includeCompleted := true // bool | optional to include completed milestones (optional)
    includeCompanyUserIds := true // bool | include company users to the responsible party ids (optional) (default to true)
    includeAssigneeTeams := true // bool | include teams related to the responsible user ids (optional)
    includeAssigneeCompanies := true // bool | include companies related to the responsible user ids (optional)
    includeArchivedProjects := true // bool | include archived projects (optional) (default to false)
    emoji := true // bool | convert emoji alias to unicode (optional) (default to true)
    countUnreadComments := true // bool | count unread comments (optional) (default to false)
    calendar := true // bool | enables the calendar style layout (optional) (default to false)
    tagIds := []int32{int32(123)} // []int32 | filter by milestone tag ids (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    ids := []int32{int32(123)} // []int32 | filter by ids (optional)
    followerIds := []int32{int32(123)} // []int32 | filter by followers' ids (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTasks := []string{"Inner_example"} // []string |  (optional)
    fieldsTasklists := []string{"Inner_example"} // []string |  (optional)
    fieldsTags := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsMilestones := []string{"Inner_example"} // []string |  (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)
    creatorIds := []int32{int32(123)} // []int32 | filter by creators' ids (optional)
    assignedToUserIds := []int32{int32(123)} // []int32 | filter by assigned user ids (optional)
    assignedToTeamIds := []int32{int32(123)} // []int32 | filter by assigned team ids (optional)
    assignedToCompanyIds := []int32{int32(123)} // []int32 | filter by assigned company ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MilestonesApi.GETProjectsApiV3ProjectsprojectIdMilestonesJson(context.Background(), projectId).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).Status(status).SearchTerm(searchTerm).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).DueBeforeDate(dueBeforeDate).DueBefore(dueBefore).DueAfterDate(dueAfterDate).DueAfter(dueAfter).ProjectHealths(projectHealths).ProjectHealth(projectHealth).PageSize(pageSize).Page(page).ShowPrivate(showPrivate).ShowPercentageCompleted(showPercentageCompleted).ShowDeleted(showDeleted).ShowAttachedTasklists(showAttachedTasklists).SearchCompaniesTeams(searchCompaniesTeams).SearchAssignees(searchAssignees).Progress(progress).OnlyStarredProjects(onlyStarredProjects).OnlyMyProjects(onlyMyProjects).OnlyAdminProjects(onlyAdminProjects).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeUpdate(includeUpdate).IncludeToday(includeToday).IncludeTeamUserIds(includeTeamUserIds).IncludeTags(includeTags).IncludeProgress(includeProgress).IncludeCompleted(includeCompleted).IncludeCompanyUserIds(includeCompanyUserIds).IncludeAssigneeTeams(includeAssigneeTeams).IncludeAssigneeCompanies(includeAssigneeCompanies).IncludeArchivedProjects(includeArchivedProjects).Emoji(emoji).CountUnreadComments(countUnreadComments).Calendar(calendar).TagIds(tagIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).Ids(ids).FollowerIds(followerIds).FieldsUsers(fieldsUsers).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsMilestones(fieldsMilestones).FieldsCompanies(fieldsCompanies).CreatorIds(creatorIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MilestonesApi.GETProjectsApiV3ProjectsprojectIdMilestonesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ProjectsprojectIdMilestonesJson`: MilestoneMilestonesResponse
    fmt.Fprintf(os.Stdout, "Response from `MilestonesApi.GETProjectsApiV3ProjectsprojectIdMilestonesJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsprojectIdMilestonesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedAfterDate** | **time.Time** | filter by updated after date (deprecated, use updatedAfter) | 
 **updatedAfter** | **time.Time** | filter by updated after date | 
 **status** | **string** | filter by list of milestone status | 
 **searchTerm** | **string** | filter by milestone name and description | 
 **reportFormat** | **string** | define the format of the report | 
 **projectStatuses** | **string** | filter by list of project status | 
 **projectStatus** | **string** | filter by project status | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;date&quot;]
 **dueBeforeDate** | **string** | filter by due before date (deprecated, use dueBefore) | 
 **dueBefore** | **string** | filter by due before date | 
 **dueAfterDate** | **string** | filter by due after date (deprecated, use dueAfter) | 
 **dueAfter** | **string** | filter by due after date | 
 **projectHealths** | **int32** | filter by list of project health | 
 **projectHealth** | **int32** | filter by list of project health (deprecated, use projectHealths) | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **showPrivate** | **bool** | show private milestones | [default to true]
 **showPercentageCompleted** | **bool** | show percentages completed | 
 **showDeleted** | **bool** | include deleted items | [default to false]
 **showAttachedTasklists** | **bool** | show attached tasklists | 
 **searchCompaniesTeams** | **bool** | include companies and teams in the search term | 
 **searchAssignees** | **bool** | include assignees in the search | 
 **progress** | **bool** | include progress of each milestone | [default to false]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **onlyMyProjects** | **bool** | filter by projects that the logged-in user is related | [default to false]
 **onlyAdminProjects** | **bool** | only include milestones from projects where the user is strictly a project admin. site admins have visibility to all projects. | 
 **matchAllTags** | **bool** | match all milestone tags | 
 **matchAllProjectTags** | **bool** | match all project tags | 
 **isReportDownload** | **bool** | generate a report document | 
 **includeUpdate** | **bool** | include milestones latest update action | 
 **includeToday** | **bool** | include today when milestone status is upcoming | [default to true]
 **includeTeamUserIds** | **bool** | include team users to the responsible party ids | [default to true]
 **includeTags** | **bool** | include tags in the reports | [default to true]
 **includeProgress** | **bool** | include percentage of tasks completed for all linked tasklists | 
 **includeCompleted** | **bool** | optional to include completed milestones | 
 **includeCompanyUserIds** | **bool** | include company users to the responsible party ids | [default to true]
 **includeAssigneeTeams** | **bool** | include teams related to the responsible user ids | 
 **includeAssigneeCompanies** | **bool** | include companies related to the responsible user ids | 
 **includeArchivedProjects** | **bool** | include archived projects | [default to false]
 **emoji** | **bool** | convert emoji alias to unicode | [default to true]
 **countUnreadComments** | **bool** | count unread comments | [default to false]
 **calendar** | **bool** | enables the calendar style layout | [default to false]
 **tagIds** | **[]int32** | filter by milestone tag ids | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include | 
 **ids** | **[]int32** | filter by ids | 
 **followerIds** | **[]int32** | filter by followers&#39; ids | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTasks** | **[]string** |  | 
 **fieldsTasklists** | **[]string** |  | 
 **fieldsTags** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsMilestones** | **[]string** |  | 
 **fieldsCompanies** | **[]string** |  | 
 **creatorIds** | **[]int32** | filter by creators&#39; ids | 
 **assignedToUserIds** | **[]int32** | filter by assigned user ids | 
 **assignedToTeamIds** | **[]int32** | filter by assigned team ids | 
 **assignedToCompanyIds** | **[]int32** | filter by assigned company ids | 

### Return type

[**MilestoneMilestonesResponse**](MilestoneMilestonesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

