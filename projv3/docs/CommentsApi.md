# \CommentsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3CommentsJson**](CommentsApi.md#GETProjectsApiV3CommentsJson) | **Get** /projects/api/v3/comments.json | Get a list of comments



## GETProjectsApiV3CommentsJson

> CommentCommentsResponse GETProjectsApiV3CommentsJson(ctx).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).SearchTerm(searchTerm).PublishedStartDate(publishedStartDate).PublishedEndDate(publishedEndDate).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).ObjectTypes(objectTypes).ObjectType(objectType).CommentStatus(commentStatus).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).StrictHTML(strictHTML).OnlyStarredProjects(onlyStarredProjects).MatchAllProjectTags(matchAllProjectTags).IncludeArchivedProjects(includeArchivedProjects).UserIds(userIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ObjectIds(objectIds).NotifiedUserIds(notifiedUserIds).Include(include).FieldsUsers(fieldsUsers).FieldsTasks(fieldsTasks).FieldsProjects(fieldsProjects).FieldsNotebooks(fieldsNotebooks).FieldsMilestones(fieldsMilestones).FieldsLinks(fieldsLinks).FieldsFileversions(fieldsFileversions).FieldsFiles(fieldsFiles).FieldsCompanies(fieldsCompanies).Execute()

Get a list of comments

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
    updatedAfterDate := time.Now() // time.Time | filter by risks updated after specified date (deprecated, use updatedAfter) (optional)
    updatedAfter := time.Now() // time.Time | filter by risks updated after specified date (optional)
    searchTerm := "searchTerm_example" // string | filter by comment content (like v1 filterText param) (optional)
    publishedStartDate := time.Now() // string | filter by published after date (optional)
    publishedEndDate := time.Now() // string | filter by published before date (optional)
    projectStatuses := "projectStatuses_example" // string | filter by project statuses (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "data")
    objectTypes := "objectTypes_example" // string | filter by object types (optional)
    objectType := "objectType_example" // string | object type where the comment is attached to (deprecated, use objectTypes) (optional)
    commentStatus := "commentStatus_example" // string | filter by commment status (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    strictHTML := true // bool | use strict html filtering for content (optional) (default to false)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    matchAllProjectTags := true // bool | match all project tags (optional)
    includeArchivedProjects := true // bool | include archived projects (optional)
    userIds := []int32{int32(123)} // []int32 | filter by user ids (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    objectIds := []int32{int32(123)} // []int32 | filter by object IDs (optional)
    notifiedUserIds := []int32{int32(123)} // []int32 | filter by users who got notified for the comments (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTasks := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsNotebooks := []string{"Inner_example"} // []string |  (optional)
    fieldsMilestones := []string{"Inner_example"} // []string |  (optional)
    fieldsLinks := []string{"Inner_example"} // []string |  (optional)
    fieldsFileversions := []string{"Inner_example"} // []string |  (optional)
    fieldsFiles := []string{"Inner_example"} // []string |  (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommentsApi.GETProjectsApiV3CommentsJson(context.Background()).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).SearchTerm(searchTerm).PublishedStartDate(publishedStartDate).PublishedEndDate(publishedEndDate).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).ObjectTypes(objectTypes).ObjectType(objectType).CommentStatus(commentStatus).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).StrictHTML(strictHTML).OnlyStarredProjects(onlyStarredProjects).MatchAllProjectTags(matchAllProjectTags).IncludeArchivedProjects(includeArchivedProjects).UserIds(userIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).ObjectIds(objectIds).NotifiedUserIds(notifiedUserIds).Include(include).FieldsUsers(fieldsUsers).FieldsTasks(fieldsTasks).FieldsProjects(fieldsProjects).FieldsNotebooks(fieldsNotebooks).FieldsMilestones(fieldsMilestones).FieldsLinks(fieldsLinks).FieldsFileversions(fieldsFileversions).FieldsFiles(fieldsFiles).FieldsCompanies(fieldsCompanies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommentsApi.GETProjectsApiV3CommentsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3CommentsJson`: CommentCommentsResponse
    fmt.Fprintf(os.Stdout, "Response from `CommentsApi.GETProjectsApiV3CommentsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3CommentsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfterDate** | **time.Time** | filter by risks updated after specified date (deprecated, use updatedAfter) | 
 **updatedAfter** | **time.Time** | filter by risks updated after specified date | 
 **searchTerm** | **string** | filter by comment content (like v1 filterText param) | 
 **publishedStartDate** | **string** | filter by published after date | 
 **publishedEndDate** | **string** | filter by published before date | 
 **projectStatuses** | **string** | filter by project statuses | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;data&quot;]
 **objectTypes** | **string** | filter by object types | 
 **objectType** | **string** | object type where the comment is attached to (deprecated, use objectTypes) | 
 **commentStatus** | **string** | filter by commment status | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **strictHTML** | **bool** | use strict html filtering for content | [default to false]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **matchAllProjectTags** | **bool** | match all project tags | 
 **includeArchivedProjects** | **bool** | include archived projects | 
 **userIds** | **[]int32** | filter by user ids | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **objectIds** | **[]int32** | filter by object IDs | 
 **notifiedUserIds** | **[]int32** | filter by users who got notified for the comments | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTasks** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsNotebooks** | **[]string** |  | 
 **fieldsMilestones** | **[]string** |  | 
 **fieldsLinks** | **[]string** |  | 
 **fieldsFileversions** | **[]string** |  | 
 **fieldsFiles** | **[]string** |  | 
 **fieldsCompanies** | **[]string** |  | 

### Return type

[**CommentCommentsResponse**](CommentCommentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

