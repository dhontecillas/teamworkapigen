# \ResourceSchedulingApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3SchedulePeopleJson**](ResourceSchedulingApi.md#GETProjectsApiV3SchedulePeopleJson) | **Get** /projects/api/v3/schedule/people.json | Return the summary for users&#39; allocations.
[**GETProjectsApiV3ScheduleProjectsJson**](ResourceSchedulingApi.md#GETProjectsApiV3ScheduleProjectsJson) | **Get** /projects/api/v3/schedule/projects.json | Return the summary for projects&#39; allocations.



## GETProjectsApiV3SchedulePeopleJson

> ScheduleUserSchedulesResponse GETProjectsApiV3SchedulePeopleJson(ctx).StartDate(startDate).SearchTerm(searchTerm).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).EndDate(endDate).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).OnlyStarredProjects(onlyStarredProjects).OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership).MatchAllProjectTags(matchAllProjectTags).HideObservedProjects(hideObservedProjects).TeamIds(teamIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsUserSchedules(fieldsUserSchedules).FieldsTimelogs(fieldsTimelogs).FieldsProjects(fieldsProjects).FieldsProjectSchedules(fieldsProjectSchedules).FieldsMilestones(fieldsMilestones).FieldsCompanies(fieldsCompanies).FieldsCalendarEvents(fieldsCalendarEvents).FieldsAllocations(fieldsAllocations).CompanyIds(companyIds).AssignedUserIds(assignedUserIds).Execute()

Return the summary for users' allocations.



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
    startDate := time.Now() // string | filter by start date (optional)
    searchTerm := "searchTerm_example" // string | filter by search term (optional)
    projectStatus := "projectStatus_example" // string | filter by project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "project")
    endDate := time.Now() // string | filter by end date (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    onlyProjectsWithExplicitMembership := true // bool | only show projects with explicit membership (optional) (default to false)
    matchAllProjectTags := true // bool | match all project tags (optional)
    hideObservedProjects := true // bool | hide projects where the logged-in user is just an observer (optional) (default to false)
    teamIds := []int32{int32(123)} // []int32 | filter by teams (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsUserSchedules := []string{"Inner_example"} // []string |  (optional)
    fieldsTimelogs := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsProjectSchedules := []string{"Inner_example"} // []string |  (optional)
    fieldsMilestones := []string{"Inner_example"} // []string |  (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)
    fieldsCalendarEvents := []string{"Inner_example"} // []string |  (optional)
    fieldsAllocations := []string{"Inner_example"} // []string |  (optional)
    companyIds := []int32{int32(123)} // []int32 | filter by companies (optional)
    assignedUserIds := []int32{int32(123)} // []int32 | filter by assigned user ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourceSchedulingApi.GETProjectsApiV3SchedulePeopleJson(context.Background()).StartDate(startDate).SearchTerm(searchTerm).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).EndDate(endDate).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).OnlyStarredProjects(onlyStarredProjects).OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership).MatchAllProjectTags(matchAllProjectTags).HideObservedProjects(hideObservedProjects).TeamIds(teamIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsUserSchedules(fieldsUserSchedules).FieldsTimelogs(fieldsTimelogs).FieldsProjects(fieldsProjects).FieldsProjectSchedules(fieldsProjectSchedules).FieldsMilestones(fieldsMilestones).FieldsCompanies(fieldsCompanies).FieldsCalendarEvents(fieldsCalendarEvents).FieldsAllocations(fieldsAllocations).CompanyIds(companyIds).AssignedUserIds(assignedUserIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceSchedulingApi.GETProjectsApiV3SchedulePeopleJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3SchedulePeopleJson`: ScheduleUserSchedulesResponse
    fmt.Fprintf(os.Stdout, "Response from `ResourceSchedulingApi.GETProjectsApiV3SchedulePeopleJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3SchedulePeopleJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | filter by start date | 
 **searchTerm** | **string** | filter by search term | 
 **projectStatus** | **string** | filter by project status | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;project&quot;]
 **endDate** | **string** | filter by end date | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **onlyProjectsWithExplicitMembership** | **bool** | only show projects with explicit membership | [default to false]
 **matchAllProjectTags** | **bool** | match all project tags | 
 **hideObservedProjects** | **bool** | hide projects where the logged-in user is just an observer | [default to false]
 **teamIds** | **[]int32** | filter by teams | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsUserSchedules** | **[]string** |  | 
 **fieldsTimelogs** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsProjectSchedules** | **[]string** |  | 
 **fieldsMilestones** | **[]string** |  | 
 **fieldsCompanies** | **[]string** |  | 
 **fieldsCalendarEvents** | **[]string** |  | 
 **fieldsAllocations** | **[]string** |  | 
 **companyIds** | **[]int32** | filter by companies | 
 **assignedUserIds** | **[]int32** | filter by assigned user ids | 

### Return type

[**ScheduleUserSchedulesResponse**](ScheduleUserSchedulesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3ScheduleProjectsJson

> ScheduleProjectSchedulesResponse GETProjectsApiV3ScheduleProjectsJson(ctx).StartDate(startDate).SearchTerm(searchTerm).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).EndDate(endDate).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).OnlyStarredProjects(onlyStarredProjects).OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership).MatchAllProjectTags(matchAllProjectTags).HideObservedProjects(hideObservedProjects).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsUserSchedules(fieldsUserSchedules).FieldsTimelogs(fieldsTimelogs).FieldsProjects(fieldsProjects).FieldsProjectSchedules(fieldsProjectSchedules).FieldsMilestones(fieldsMilestones).FieldsCompanies(fieldsCompanies).FieldsCalendarEvents(fieldsCalendarEvents).FieldsAllocations(fieldsAllocations).AssignedUserIds(assignedUserIds).Execute()

Return the summary for projects' allocations.



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
    startDate := time.Now() // string | filter by start date (optional)
    searchTerm := "searchTerm_example" // string | filter by search term (optional)
    projectStatus := "projectStatus_example" // string | filter by project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "project")
    endDate := time.Now() // string | filter by end date (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    onlyProjectsWithExplicitMembership := true // bool | only show projects with explicit membership (optional) (default to false)
    matchAllProjectTags := true // bool | match all project tags (optional)
    hideObservedProjects := true // bool | hide projects where the logged-in user is just an observer (optional) (default to false)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsUserSchedules := []string{"Inner_example"} // []string |  (optional)
    fieldsTimelogs := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsProjectSchedules := []string{"Inner_example"} // []string |  (optional)
    fieldsMilestones := []string{"Inner_example"} // []string |  (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)
    fieldsCalendarEvents := []string{"Inner_example"} // []string |  (optional)
    fieldsAllocations := []string{"Inner_example"} // []string |  (optional)
    assignedUserIds := []int32{int32(123)} // []int32 | filter by assigned user ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourceSchedulingApi.GETProjectsApiV3ScheduleProjectsJson(context.Background()).StartDate(startDate).SearchTerm(searchTerm).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).EndDate(endDate).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).OnlyStarredProjects(onlyStarredProjects).OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership).MatchAllProjectTags(matchAllProjectTags).HideObservedProjects(hideObservedProjects).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsUserSchedules(fieldsUserSchedules).FieldsTimelogs(fieldsTimelogs).FieldsProjects(fieldsProjects).FieldsProjectSchedules(fieldsProjectSchedules).FieldsMilestones(fieldsMilestones).FieldsCompanies(fieldsCompanies).FieldsCalendarEvents(fieldsCalendarEvents).FieldsAllocations(fieldsAllocations).AssignedUserIds(assignedUserIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceSchedulingApi.GETProjectsApiV3ScheduleProjectsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ScheduleProjectsJson`: ScheduleProjectSchedulesResponse
    fmt.Fprintf(os.Stdout, "Response from `ResourceSchedulingApi.GETProjectsApiV3ScheduleProjectsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ScheduleProjectsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | filter by start date | 
 **searchTerm** | **string** | filter by search term | 
 **projectStatus** | **string** | filter by project status | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;project&quot;]
 **endDate** | **string** | filter by end date | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **onlyProjectsWithExplicitMembership** | **bool** | only show projects with explicit membership | [default to false]
 **matchAllProjectTags** | **bool** | match all project tags | 
 **hideObservedProjects** | **bool** | hide projects where the logged-in user is just an observer | [default to false]
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsUserSchedules** | **[]string** |  | 
 **fieldsTimelogs** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsProjectSchedules** | **[]string** |  | 
 **fieldsMilestones** | **[]string** |  | 
 **fieldsCompanies** | **[]string** |  | 
 **fieldsCalendarEvents** | **[]string** |  | 
 **fieldsAllocations** | **[]string** |  | 
 **assignedUserIds** | **[]int32** | filter by assigned user ids | 

### Return type

[**ScheduleProjectSchedulesResponse**](ScheduleProjectSchedulesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

