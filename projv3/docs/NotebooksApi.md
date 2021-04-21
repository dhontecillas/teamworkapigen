# \NotebooksApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEProjectsApiV3NotebooksnotebookIdJson**](NotebooksApi.md#DELETEProjectsApiV3NotebooksnotebookIdJson) | **Delete** /projects/api/v3/notebooks/{notebookId}.json | Delete an existing notebook.
[**DELETEProjectsApiV3NotebooksnotebookIdVersionsJson**](NotebooksApi.md#DELETEProjectsApiV3NotebooksnotebookIdVersionsJson) | **Delete** /projects/api/v3/notebooks/{notebookId}/versions.json | Delete notebook versions.
[**GETProjectsApiV3NotebooksJson**](NotebooksApi.md#GETProjectsApiV3NotebooksJson) | **Get** /projects/api/v3/notebooks.json | Get a list of notebooks.
[**GETProjectsApiV3NotebooksnotebookIdCompareJson**](NotebooksApi.md#GETProjectsApiV3NotebooksnotebookIdCompareJson) | **Get** /projects/api/v3/notebooks/{notebookId}/compare.json | Compare 2 notebook versions
[**GETProjectsApiV3NotebooksnotebookIdJson**](NotebooksApi.md#GETProjectsApiV3NotebooksnotebookIdJson) | **Get** /projects/api/v3/notebooks/{notebookId}.json | Get a specific notebook.
[**GETProjectsApiV3NotebooksnotebookIdVersionsJson**](NotebooksApi.md#GETProjectsApiV3NotebooksnotebookIdVersionsJson) | **Get** /projects/api/v3/notebooks/{notebookId}/versions.json | Get notebook versions
[**GETProjectsApiV3NotebooksnotebookIdVersionsversionIdJson**](NotebooksApi.md#GETProjectsApiV3NotebooksnotebookIdVersionsversionIdJson) | **Get** /projects/api/v3/notebooks/{notebookId}/versions/{versionId}.json | Get a notebook version
[**PATCHProjectsApiV3NotebooksnotebookIdJson**](NotebooksApi.md#PATCHProjectsApiV3NotebooksnotebookIdJson) | **Patch** /projects/api/v3/notebooks/{notebookId}.json | Edits a notebook
[**POSTProjectsApiV3ProjectsprojectIdNotebooksJson**](NotebooksApi.md#POSTProjectsApiV3ProjectsprojectIdNotebooksJson) | **Post** /projects/api/v3/projects/{projectId}/notebooks.json | Create a new notebook.
[**PUTProjectsApiV3NotebooksnotebookIdLockJson**](NotebooksApi.md#PUTProjectsApiV3NotebooksnotebookIdLockJson) | **Put** /projects/api/v3/notebooks/{notebookId}/lock.json | Lock a notebook
[**PUTProjectsApiV3NotebooksnotebookIdUnlockJson**](NotebooksApi.md#PUTProjectsApiV3NotebooksnotebookIdUnlockJson) | **Put** /projects/api/v3/notebooks/{notebookId}/unlock.json | Unlock a notebook



## DELETEProjectsApiV3NotebooksnotebookIdJson

> DELETEProjectsApiV3NotebooksnotebookIdJson(ctx, notebookId).Execute()

Delete an existing notebook.

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
    notebookId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotebooksApi.DELETEProjectsApiV3NotebooksnotebookIdJson(context.Background(), notebookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.DELETEProjectsApiV3NotebooksnotebookIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notebookId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEProjectsApiV3NotebooksnotebookIdJsonRequest struct via the builder pattern


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


## DELETEProjectsApiV3NotebooksnotebookIdVersionsJson

> DELETEProjectsApiV3NotebooksnotebookIdVersionsJson(ctx, notebookId).Execute()

Delete notebook versions.



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
    notebookId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotebooksApi.DELETEProjectsApiV3NotebooksnotebookIdVersionsJson(context.Background(), notebookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.DELETEProjectsApiV3NotebooksnotebookIdVersionsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notebookId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEProjectsApiV3NotebooksnotebookIdVersionsJsonRequest struct via the builder pattern


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


## GETProjectsApiV3NotebooksJson

> NotebookNotebooksResponse GETProjectsApiV3NotebooksJson(ctx).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).SearchTerm(searchTerm).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).CreatedAfterDate(createdAfterDate).CreatedAfter(createdAfter).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).SecureOnly(secureOnly).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).LockedOnly(lockedOnly).GetEmoji(getEmoji).TagIds(tagIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).Ids(ids).FieldsUsers(fieldsUsers).FieldsTeams(fieldsTeams).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsNotebooks(fieldsNotebooks).FieldsNotebookCategories(fieldsNotebookCategories).FieldsCompanies(fieldsCompanies).CreatorIds(creatorIds).CategoryIds(categoryIds).Execute()

Get a list of notebooks.

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
    searchTerm := "searchTerm_example" // string | filter by notebook name and description (optional)
    projectStatuses := "projectStatuses_example" // string | filter by project statuses (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "dateUpdated")
    createdAfterDate := time.Now() // string | filter by created after date (deprecated, use createdAfter) (optional)
    createdAfter := time.Now() // string | filter by created after date (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    showDeleted := true // bool | include deleted items (optional) (default to false)
    secureOnly := true // bool | filter by showing only secure notebooks (optional) (default to false)
    matchAllTags := true // bool | match all notebook tags (optional)
    matchAllProjectTags := true // bool | match all project tags (optional)
    lockedOnly := true // bool | filter by showing only locked notebooks (optional) (default to false)
    getEmoji := true // bool | parse Emoji's (optional)
    tagIds := []int32{int32(123)} // []int32 | filter by notebook tag ids (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    ids := []int32{int32(123)} // []int32 | filter by ids (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTeams := []string{"Inner_example"} // []string |  (optional)
    fieldsTags := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsNotebooks := []string{"Inner_example"} // []string |  (optional)
    fieldsNotebookCategories := []string{"Inner_example"} // []string |  (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)
    creatorIds := []int32{int32(123)} // []int32 | filter by creators' ids (optional)
    categoryIds := []int32{int32(123)} // []int32 | filter by notebook category ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotebooksApi.GETProjectsApiV3NotebooksJson(context.Background()).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).SearchTerm(searchTerm).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).CreatedAfterDate(createdAfterDate).CreatedAfter(createdAfter).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).SecureOnly(secureOnly).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).LockedOnly(lockedOnly).GetEmoji(getEmoji).TagIds(tagIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).Ids(ids).FieldsUsers(fieldsUsers).FieldsTeams(fieldsTeams).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsNotebooks(fieldsNotebooks).FieldsNotebookCategories(fieldsNotebookCategories).FieldsCompanies(fieldsCompanies).CreatorIds(creatorIds).CategoryIds(categoryIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.GETProjectsApiV3NotebooksJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3NotebooksJson`: NotebookNotebooksResponse
    fmt.Fprintf(os.Stdout, "Response from `NotebooksApi.GETProjectsApiV3NotebooksJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3NotebooksJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfterDate** | **time.Time** | filter by updated after date (deprecated, use updatedAfter) | 
 **updatedAfter** | **time.Time** | filter by updated after date | 
 **searchTerm** | **string** | filter by notebook name and description | 
 **projectStatuses** | **string** | filter by project statuses | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;dateUpdated&quot;]
 **createdAfterDate** | **string** | filter by created after date (deprecated, use createdAfter) | 
 **createdAfter** | **string** | filter by created after date | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **showDeleted** | **bool** | include deleted items | [default to false]
 **secureOnly** | **bool** | filter by showing only secure notebooks | [default to false]
 **matchAllTags** | **bool** | match all notebook tags | 
 **matchAllProjectTags** | **bool** | match all project tags | 
 **lockedOnly** | **bool** | filter by showing only locked notebooks | [default to false]
 **getEmoji** | **bool** | parse Emoji&#39;s | 
 **tagIds** | **[]int32** | filter by notebook tag ids | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include | 
 **ids** | **[]int32** | filter by ids | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTeams** | **[]string** |  | 
 **fieldsTags** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsNotebooks** | **[]string** |  | 
 **fieldsNotebookCategories** | **[]string** |  | 
 **fieldsCompanies** | **[]string** |  | 
 **creatorIds** | **[]int32** | filter by creators&#39; ids | 
 **categoryIds** | **[]int32** | filter by notebook category ids | 

### Return type

[**NotebookNotebooksResponse**](NotebookNotebooksResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3NotebooksnotebookIdCompareJson

> NotebookVersionResponse GETProjectsApiV3NotebooksnotebookIdCompareJson(ctx, notebookId).Include(include).FieldsUsers(fieldsUsers).Execute()

Compare 2 notebook versions



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
    notebookId := int32(56) // int32 | 
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotebooksApi.GETProjectsApiV3NotebooksnotebookIdCompareJson(context.Background(), notebookId).Include(include).FieldsUsers(fieldsUsers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.GETProjectsApiV3NotebooksnotebookIdCompareJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3NotebooksnotebookIdCompareJson`: NotebookVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `NotebooksApi.GETProjectsApiV3NotebooksnotebookIdCompareJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notebookId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3NotebooksnotebookIdCompareJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 

### Return type

[**NotebookVersionResponse**](NotebookVersionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3NotebooksnotebookIdJson

> NotebookResponse GETProjectsApiV3NotebooksnotebookIdJson(ctx, notebookId).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).ShowDeleted(showDeleted).GetEmoji(getEmoji).Include(include).FieldsUsers(fieldsUsers).FieldsTeams(fieldsTeams).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsNotebooks(fieldsNotebooks).FieldsNotebookCategories(fieldsNotebookCategories).FieldsCompanies(fieldsCompanies).Execute()

Get a specific notebook.

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
    notebookId := int32(56) // int32 | 
    updatedAfterDate := time.Now() // time.Time | filter by updated after date (deprecated, use updatedAfter) (optional)
    updatedAfter := time.Now() // time.Time | filter by updated after date (optional)
    showDeleted := true // bool | show deleted notebooks (optional)
    getEmoji := true // bool | parse Emoji's (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTeams := []string{"Inner_example"} // []string |  (optional)
    fieldsTags := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsNotebooks := []string{"Inner_example"} // []string |  (optional)
    fieldsNotebookCategories := []string{"Inner_example"} // []string |  (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotebooksApi.GETProjectsApiV3NotebooksnotebookIdJson(context.Background(), notebookId).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).ShowDeleted(showDeleted).GetEmoji(getEmoji).Include(include).FieldsUsers(fieldsUsers).FieldsTeams(fieldsTeams).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsNotebooks(fieldsNotebooks).FieldsNotebookCategories(fieldsNotebookCategories).FieldsCompanies(fieldsCompanies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.GETProjectsApiV3NotebooksnotebookIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3NotebooksnotebookIdJson`: NotebookResponse
    fmt.Fprintf(os.Stdout, "Response from `NotebooksApi.GETProjectsApiV3NotebooksnotebookIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notebookId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3NotebooksnotebookIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedAfterDate** | **time.Time** | filter by updated after date (deprecated, use updatedAfter) | 
 **updatedAfter** | **time.Time** | filter by updated after date | 
 **showDeleted** | **bool** | show deleted notebooks | 
 **getEmoji** | **bool** | parse Emoji&#39;s | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTeams** | **[]string** |  | 
 **fieldsTags** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsNotebooks** | **[]string** |  | 
 **fieldsNotebookCategories** | **[]string** |  | 
 **fieldsCompanies** | **[]string** |  | 

### Return type

[**NotebookResponse**](NotebookResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3NotebooksnotebookIdVersionsJson

> NotebookVersionsResponse GETProjectsApiV3NotebooksnotebookIdVersionsJson(ctx, notebookId).Include(include).FieldsUsers(fieldsUsers).Execute()

Get notebook versions



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
    notebookId := int32(56) // int32 | 
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotebooksApi.GETProjectsApiV3NotebooksnotebookIdVersionsJson(context.Background(), notebookId).Include(include).FieldsUsers(fieldsUsers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.GETProjectsApiV3NotebooksnotebookIdVersionsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3NotebooksnotebookIdVersionsJson`: NotebookVersionsResponse
    fmt.Fprintf(os.Stdout, "Response from `NotebooksApi.GETProjectsApiV3NotebooksnotebookIdVersionsJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notebookId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3NotebooksnotebookIdVersionsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 

### Return type

[**NotebookVersionsResponse**](NotebookVersionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3NotebooksnotebookIdVersionsversionIdJson

> NotebookVersionResponse GETProjectsApiV3NotebooksnotebookIdVersionsversionIdJson(ctx, versionId, notebookId).Include(include).FieldsUsers(fieldsUsers).Execute()

Get a notebook version



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
    versionId := int32(56) // int32 | 
    notebookId := int32(56) // int32 | 
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotebooksApi.GETProjectsApiV3NotebooksnotebookIdVersionsversionIdJson(context.Background(), versionId, notebookId).Include(include).FieldsUsers(fieldsUsers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.GETProjectsApiV3NotebooksnotebookIdVersionsversionIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3NotebooksnotebookIdVersionsversionIdJson`: NotebookVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `NotebooksApi.GETProjectsApiV3NotebooksnotebookIdVersionsversionIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**versionId** | **int32** |  | 
**notebookId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3NotebooksnotebookIdVersionsversionIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 

### Return type

[**NotebookVersionResponse**](NotebookVersionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHProjectsApiV3NotebooksnotebookIdJson

> NotebookResponse PATCHProjectsApiV3NotebooksnotebookIdJson(ctx, notebookId).NotebookRequest(notebookRequest).GetEmoji(getEmoji).Include(include).FieldsUsers(fieldsUsers).FieldsTeams(fieldsTeams).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsNotebooks(fieldsNotebooks).FieldsNotebookCategories(fieldsNotebookCategories).FieldsCompanies(fieldsCompanies).Execute()

Edits a notebook



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
    notebookId := int32(56) // int32 | 
    notebookRequest := *openapiclient.NewNotebookRequest() // NotebookRequest | 
    getEmoji := true // bool | parse Emoji's (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTeams := []string{"Inner_example"} // []string |  (optional)
    fieldsTags := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsNotebooks := []string{"Inner_example"} // []string |  (optional)
    fieldsNotebookCategories := []string{"Inner_example"} // []string |  (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotebooksApi.PATCHProjectsApiV3NotebooksnotebookIdJson(context.Background(), notebookId).NotebookRequest(notebookRequest).GetEmoji(getEmoji).Include(include).FieldsUsers(fieldsUsers).FieldsTeams(fieldsTeams).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsNotebooks(fieldsNotebooks).FieldsNotebookCategories(fieldsNotebookCategories).FieldsCompanies(fieldsCompanies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.PATCHProjectsApiV3NotebooksnotebookIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHProjectsApiV3NotebooksnotebookIdJson`: NotebookResponse
    fmt.Fprintf(os.Stdout, "Response from `NotebooksApi.PATCHProjectsApiV3NotebooksnotebookIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notebookId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHProjectsApiV3NotebooksnotebookIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notebookRequest** | [**NotebookRequest**](NotebookRequest.md) |  | 
 **getEmoji** | **bool** | parse Emoji&#39;s | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTeams** | **[]string** |  | 
 **fieldsTags** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsNotebooks** | **[]string** |  | 
 **fieldsNotebookCategories** | **[]string** |  | 
 **fieldsCompanies** | **[]string** |  | 

### Return type

[**NotebookResponse**](NotebookResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsApiV3ProjectsprojectIdNotebooksJson

> NotebookResponse POSTProjectsApiV3ProjectsprojectIdNotebooksJson(ctx, projectId).NotebookRequest(notebookRequest).GetEmoji(getEmoji).Include(include).FieldsUsers(fieldsUsers).FieldsTeams(fieldsTeams).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsNotebooks(fieldsNotebooks).FieldsNotebookCategories(fieldsNotebookCategories).FieldsCompanies(fieldsCompanies).Execute()

Create a new notebook.

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
    notebookRequest := *openapiclient.NewNotebookRequest() // NotebookRequest | 
    getEmoji := true // bool | parse Emoji's (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTeams := []string{"Inner_example"} // []string |  (optional)
    fieldsTags := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsNotebooks := []string{"Inner_example"} // []string |  (optional)
    fieldsNotebookCategories := []string{"Inner_example"} // []string |  (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotebooksApi.POSTProjectsApiV3ProjectsprojectIdNotebooksJson(context.Background(), projectId).NotebookRequest(notebookRequest).GetEmoji(getEmoji).Include(include).FieldsUsers(fieldsUsers).FieldsTeams(fieldsTeams).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsNotebooks(fieldsNotebooks).FieldsNotebookCategories(fieldsNotebookCategories).FieldsCompanies(fieldsCompanies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.POSTProjectsApiV3ProjectsprojectIdNotebooksJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsApiV3ProjectsprojectIdNotebooksJson`: NotebookResponse
    fmt.Fprintf(os.Stdout, "Response from `NotebooksApi.POSTProjectsApiV3ProjectsprojectIdNotebooksJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3ProjectsprojectIdNotebooksJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notebookRequest** | [**NotebookRequest**](NotebookRequest.md) |  | 
 **getEmoji** | **bool** | parse Emoji&#39;s | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTeams** | **[]string** |  | 
 **fieldsTags** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsNotebooks** | **[]string** |  | 
 **fieldsNotebookCategories** | **[]string** |  | 
 **fieldsCompanies** | **[]string** |  | 

### Return type

[**NotebookResponse**](NotebookResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV3NotebooksnotebookIdLockJson

> PUTProjectsApiV3NotebooksnotebookIdLockJson(ctx, notebookId).Execute()

Lock a notebook



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
    notebookId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotebooksApi.PUTProjectsApiV3NotebooksnotebookIdLockJson(context.Background(), notebookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.PUTProjectsApiV3NotebooksnotebookIdLockJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notebookId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3NotebooksnotebookIdLockJsonRequest struct via the builder pattern


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


## PUTProjectsApiV3NotebooksnotebookIdUnlockJson

> PUTProjectsApiV3NotebooksnotebookIdUnlockJson(ctx, notebookId).Execute()

Unlock a notebook



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
    notebookId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotebooksApi.PUTProjectsApiV3NotebooksnotebookIdUnlockJson(context.Background(), notebookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.PUTProjectsApiV3NotebooksnotebookIdUnlockJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notebookId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3NotebooksnotebookIdUnlockJsonRequest struct via the builder pattern


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

