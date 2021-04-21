# \ProjectUpdatesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3ProjectsUpdatesJson**](ProjectUpdatesApi.md#GETProjectsApiV3ProjectsUpdatesJson) | **Get** /projects/api/v3/projects/updates.json | Get all project updates
[**GETProjectsApiV3ProjectsprojectIdsUpdatesJson**](ProjectUpdatesApi.md#GETProjectsApiV3ProjectsprojectIdsUpdatesJson) | **Get** /projects/api/v3/projects/{projectIds}/updates.json | Get all updates from a specific project



## GETProjectsApiV3ProjectsUpdatesJson

> UpdateProjectUpdatesResponse GETProjectsApiV3ProjectsUpdatesJson(ctx).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).CreatedAfter(createdAfter).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).Reactions(reactions).OnlyStarredProjects(onlyStarredProjects).MatchAllProjectTags(matchAllProjectTags).IncludeArchivedProjects(includeArchivedProjects).Emoji(emoji).ActiveOnly(activeOnly).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsProjects(fieldsProjects).FieldsProjectUpdates(fieldsProjectUpdates).Execute()

Get all project updates



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
    projectStatuses := "projectStatuses_example" // string | list of project status (optional)
    projectStatus := "projectStatus_example" // string | filter by project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "date")
    createdAfter := time.Now() // time.Time | filter by creation date (optional)
    projectId := int32(56) // int32 | filter by project id (optional)
    projectHealths := int32(56) // int32 | filter by project health (optional)
    projectHealth := int32(56) // int32 | filter by project health (deprecated, use projectHealths) (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    showDeleted := true // bool | include deleted items (optional) (default to false)
    reactions := true // bool | add reactions to the response (optional) (default to false)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    matchAllProjectTags := true // bool | match all project tags (optional)
    includeArchivedProjects := true // bool | include archived projects (optional)
    emoji := true // bool | parse emojis to unicode (optional) (default to true)
    activeOnly := true // bool | filter by active (optional) (default to true)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsProjectUpdates := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectUpdatesApi.GETProjectsApiV3ProjectsUpdatesJson(context.Background()).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).CreatedAfter(createdAfter).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).Reactions(reactions).OnlyStarredProjects(onlyStarredProjects).MatchAllProjectTags(matchAllProjectTags).IncludeArchivedProjects(includeArchivedProjects).Emoji(emoji).ActiveOnly(activeOnly).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsProjects(fieldsProjects).FieldsProjectUpdates(fieldsProjectUpdates).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectUpdatesApi.GETProjectsApiV3ProjectsUpdatesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ProjectsUpdatesJson`: UpdateProjectUpdatesResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectUpdatesApi.GETProjectsApiV3ProjectsUpdatesJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsUpdatesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectStatuses** | **string** | list of project status | 
 **projectStatus** | **string** | filter by project status | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;date&quot;]
 **createdAfter** | **time.Time** | filter by creation date | 
 **projectId** | **int32** | filter by project id | 
 **projectHealths** | **int32** | filter by project health | 
 **projectHealth** | **int32** | filter by project health (deprecated, use projectHealths) | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **showDeleted** | **bool** | include deleted items | [default to false]
 **reactions** | **bool** | add reactions to the response | [default to false]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **matchAllProjectTags** | **bool** | match all project tags | 
 **includeArchivedProjects** | **bool** | include archived projects | 
 **emoji** | **bool** | parse emojis to unicode | [default to true]
 **activeOnly** | **bool** | filter by active | [default to true]
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsProjectUpdates** | **[]string** |  | 

### Return type

[**UpdateProjectUpdatesResponse**](UpdateProjectUpdatesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3ProjectsprojectIdsUpdatesJson

> UpdateProjectUpdatesResponse GETProjectsApiV3ProjectsprojectIdsUpdatesJson(ctx, projectIds).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).CreatedAfter(createdAfter).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).Reactions(reactions).OnlyStarredProjects(onlyStarredProjects).MatchAllProjectTags(matchAllProjectTags).IncludeArchivedProjects(includeArchivedProjects).Emoji(emoji).ActiveOnly(activeOnly).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsProjects(fieldsProjects).FieldsProjectUpdates(fieldsProjectUpdates).Execute()

Get all updates from a specific project



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
    projectIds := int32(56) // int32 | 
    projectStatuses := "projectStatuses_example" // string | list of project status (optional)
    projectStatus := "projectStatus_example" // string | filter by project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "date")
    createdAfter := time.Now() // time.Time | filter by creation date (optional)
    projectId := int32(56) // int32 | filter by project id (optional)
    projectHealths := int32(56) // int32 | filter by project health (optional)
    projectHealth := int32(56) // int32 | filter by project health (deprecated, use projectHealths) (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    showDeleted := true // bool | include deleted items (optional) (default to false)
    reactions := true // bool | add reactions to the response (optional) (default to false)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    matchAllProjectTags := true // bool | match all project tags (optional)
    includeArchivedProjects := true // bool | include archived projects (optional)
    emoji := true // bool | parse emojis to unicode (optional) (default to true)
    activeOnly := true // bool | filter by active (optional) (default to true)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsProjectUpdates := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectUpdatesApi.GETProjectsApiV3ProjectsprojectIdsUpdatesJson(context.Background(), projectIds).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).CreatedAfter(createdAfter).ProjectId(projectId).ProjectHealths(projectHealths).ProjectHealth(projectHealth).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).Reactions(reactions).OnlyStarredProjects(onlyStarredProjects).MatchAllProjectTags(matchAllProjectTags).IncludeArchivedProjects(includeArchivedProjects).Emoji(emoji).ActiveOnly(activeOnly).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsProjects(fieldsProjects).FieldsProjectUpdates(fieldsProjectUpdates).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectUpdatesApi.GETProjectsApiV3ProjectsprojectIdsUpdatesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ProjectsprojectIdsUpdatesJson`: UpdateProjectUpdatesResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectUpdatesApi.GETProjectsApiV3ProjectsprojectIdsUpdatesJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectIds** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectStatuses** | **string** | list of project status | 
 **projectStatus** | **string** | filter by project status | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;date&quot;]
 **createdAfter** | **time.Time** | filter by creation date | 
 **projectId** | **int32** | filter by project id | 
 **projectHealths** | **int32** | filter by project health | 
 **projectHealth** | **int32** | filter by project health (deprecated, use projectHealths) | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **showDeleted** | **bool** | include deleted items | [default to false]
 **reactions** | **bool** | add reactions to the response | [default to false]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **matchAllProjectTags** | **bool** | match all project tags | 
 **includeArchivedProjects** | **bool** | include archived projects | 
 **emoji** | **bool** | parse emojis to unicode | [default to true]
 **activeOnly** | **bool** | filter by active | [default to true]
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsProjectUpdates** | **[]string** |  | 

### Return type

[**UpdateProjectUpdatesResponse**](UpdateProjectUpdatesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

