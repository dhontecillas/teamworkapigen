# \FilesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEProjectsApiV3FilesfileIdJson**](FilesApi.md#DELETEProjectsApiV3FilesfileIdJson) | **Delete** /projects/api/v3/files/{fileId}.json | Delete an existing file and it&#39;s versions.
[**GETProjectsApiV3FilesAvailableJson**](FilesApi.md#GETProjectsApiV3FilesAvailableJson) | **Get** /projects/api/v3/files/available.json | Retrieves available space on account
[**GETProjectsApiV3FilesChangesJson**](FilesApi.md#GETProjectsApiV3FilesChangesJson) | **Get** /projects/api/v3/files/changes.json | Get all recently changed files.
[**GETProjectsApiV3FilesJson**](FilesApi.md#GETProjectsApiV3FilesJson) | **Get** /projects/api/v3/files.json | Get a list of files.
[**GETProjectsApiV3FilesUsageJson**](FilesApi.md#GETProjectsApiV3FilesUsageJson) | **Get** /projects/api/v3/files/usage.json | Retrieve file usage on account
[**GETProjectsApiV3FilesfileIdJson**](FilesApi.md#GETProjectsApiV3FilesfileIdJson) | **Get** /projects/api/v3/files/{fileId}.json | Get a specific file.
[**GETProjectsApiV3FileversionidJson**](FilesApi.md#GETProjectsApiV3FileversionidJson) | **Get** /projects/api/v3/fileversion/{id}.json | Get a specific fileversion.
[**GETProjectsApiV3ProjectsprojectIdFilesUsageJson**](FilesApi.md#GETProjectsApiV3ProjectsprojectIdFilesUsageJson) | **Get** /projects/api/v3/projects/{projectId}/files/usage.json | Retrieve file usage on a project
[**PATCHProjectsApiV3FilesfileIdJson**](FilesApi.md#PATCHProjectsApiV3FilesfileIdJson) | **Patch** /projects/api/v3/files/{fileId}.json | Update an existing file
[**POSTProjectsApiV3FilesArchiveJson**](FilesApi.md#POSTProjectsApiV3FilesArchiveJson) | **Post** /projects/api/v3/files/archive.json | Returns an URL for multiple files
[**POSTProjectsApiV3FilesprojectFileIdJson**](FilesApi.md#POSTProjectsApiV3FilesprojectFileIdJson) | **Post** /projects/api/v3/files/{projectFileId}.json | Create a new fileversion for the project file id.



## DELETEProjectsApiV3FilesfileIdJson

> DELETEProjectsApiV3FilesfileIdJson(ctx, fileId).Execute()

Delete an existing file and it's versions.



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
    fileId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FilesApi.DELETEProjectsApiV3FilesfileIdJson(context.Background(), fileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.DELETEProjectsApiV3FilesfileIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEProjectsApiV3FilesfileIdJsonRequest struct via the builder pattern


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


## GETProjectsApiV3FilesAvailableJson

> FileAvailableResponse GETProjectsApiV3FilesAvailableJson(ctx).Execute()

Retrieves available space on account



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
    resp, r, err := api_client.FilesApi.GETProjectsApiV3FilesAvailableJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.GETProjectsApiV3FilesAvailableJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3FilesAvailableJson`: FileAvailableResponse
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.GETProjectsApiV3FilesAvailableJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3FilesAvailableJsonRequest struct via the builder pattern


### Return type

[**FileAvailableResponse**](FileAvailableResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3FilesChangesJson

> ChangeChangesResponse GETProjectsApiV3FilesChangesJson(ctx).UpdatedAfterDateTime(updatedAfterDateTime).UpdatedAfter(updatedAfter).PageSize(pageSize).IncludeParentId(includeParentId).Execute()

Get all recently changed files.



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
    updatedAfterDateTime := time.Now() // time.Time | filter by risks updated after specified date (deprecated, use updatedAfter) (optional)
    updatedAfter := time.Now() // time.Time | filter by risks updated after specified date (optional)
    pageSize := int32(56) // int32 | max number of items to retrieve (optional) (default to 1000)
    includeParentId := true // bool | return parent ID (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FilesApi.GETProjectsApiV3FilesChangesJson(context.Background()).UpdatedAfterDateTime(updatedAfterDateTime).UpdatedAfter(updatedAfter).PageSize(pageSize).IncludeParentId(includeParentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.GETProjectsApiV3FilesChangesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3FilesChangesJson`: ChangeChangesResponse
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.GETProjectsApiV3FilesChangesJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3FilesChangesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfterDateTime** | **time.Time** | filter by risks updated after specified date (deprecated, use updatedAfter) | 
 **updatedAfter** | **time.Time** | filter by risks updated after specified date | 
 **pageSize** | **int32** | max number of items to retrieve | [default to 1000]
 **includeParentId** | **bool** | return parent ID | [default to false]

### Return type

[**ChangeChangesResponse**](ChangeChangesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3FilesJson

> FileProjectFilesResponse GETProjectsApiV3FilesJson(ctx).UploadedStartDate(uploadedStartDate).UploadedEndDate(uploadedEndDate).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).Status(status).SearchTerm(searchTerm).ProjectType(projectType).OrderMode(orderMode).OrderBy(orderBy).VersionId(versionId).Version(version).TaskId(taskId).ProjectId(projectId).PageSize(pageSize).Page(page).Id(id).CategoryId(categoryId).SkipInternalFiles(skipInternalFiles).SkipExternalFiles(skipExternalFiles).ShowDeleted(showDeleted).SearchAllFields(searchAllFields).MatchAllTags(matchAllTags).LockedOnly(lockedOnly).GetVersions(getVersions).GetVersionReactions(getVersionReactions).GetRelatedItems(getRelatedItems).GetReactions(getReactions).GetLikes(getLikes).GetComments(getComments).FromDocumentEditor(fromDocumentEditor).VersionIds(versionIds).UserIds(userIds).TagIds(tagIds).Include(include).FieldsUsers(fieldsUsers).FieldsTasks(fieldsTasks).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsMessages(fieldsMessages).FieldsFileCategories(fieldsFileCategories).FieldsComments(fieldsComments).Execute()

Get a list of files.



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
    uploadedStartDate := time.Now() // string | filter by uploaded start date (optional)
    uploadedEndDate := time.Now() // string | filter by uploaded end date (optional)
    updatedAfterDate := time.Now() // time.Time | filter by updated after date (deprecated, use updatedAfter) (optional)
    updatedAfter := time.Now() // time.Time | filter by updated after date (optional)
    status := "status_example" // string | filter by 'all', 'active', 'inactive' (optional)
    searchTerm := "searchTerm_example" // string | filter by file name (optional)
    projectType := "projectType_example" // string | filter by project type. Leave empty to retrieve all types eg. &projectType= (optional) (default to "normal")
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "name")
    versionId := int32(56) // int32 | filter by version id, returns the latest file (optional)
    version := int32(56) // int32 | select the version of the file to include when filtering by id. (optional)
    taskId := int32(56) // int32 | filter by task ids. (optional)
    projectId := int32(56) // int32 | filter by project id (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    id := int32(56) // int32 | filter by project file id (optional)
    categoryId := int32(56) // int32 | filter by category id (optional)
    skipInternalFiles := true // bool | filter by omitting internal files (optional)
    skipExternalFiles := true // bool | filter by omitting external files (optional)
    showDeleted := true // bool | include deleted files. (optional)
    searchAllFields := true // bool | filter by display name, file extension, file category, original file name and the name of the latest uploader. (optional)
    matchAllTags := true // bool | filter by matching all tag ids (optional)
    lockedOnly := true // bool | filter by files that are locked (optional)
    getVersions := true // bool | include the file versions of the file (optional)
    getVersionReactions := true // bool | include the reactions of each file version (optional)
    getRelatedItems := true // bool | get related tasks, messages, comments for a file. default true when filtering by id. (optional)
    getReactions := true // bool | include reactions of the file (optional)
    getLikes := true // bool | include the number of likes (optional)
    getComments := true // bool | include the number of comments and comments read. default true when filtering by id (optional)
    fromDocumentEditor := true // bool | document editor files (optional)
    versionIds := []int32{int32(123)} // []int32 | filter by version ids, returns the latest files (optional)
    userIds := []int32{int32(123)} // []int32 | filter by user ids (optional)
    tagIds := []int32{int32(123)} // []int32 | filter by files tag ids. (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTasks := []string{"Inner_example"} // []string |  (optional)
    fieldsTags := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsMessages := []string{"Inner_example"} // []string |  (optional)
    fieldsFileCategories := []string{"Inner_example"} // []string |  (optional)
    fieldsComments := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FilesApi.GETProjectsApiV3FilesJson(context.Background()).UploadedStartDate(uploadedStartDate).UploadedEndDate(uploadedEndDate).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).Status(status).SearchTerm(searchTerm).ProjectType(projectType).OrderMode(orderMode).OrderBy(orderBy).VersionId(versionId).Version(version).TaskId(taskId).ProjectId(projectId).PageSize(pageSize).Page(page).Id(id).CategoryId(categoryId).SkipInternalFiles(skipInternalFiles).SkipExternalFiles(skipExternalFiles).ShowDeleted(showDeleted).SearchAllFields(searchAllFields).MatchAllTags(matchAllTags).LockedOnly(lockedOnly).GetVersions(getVersions).GetVersionReactions(getVersionReactions).GetRelatedItems(getRelatedItems).GetReactions(getReactions).GetLikes(getLikes).GetComments(getComments).FromDocumentEditor(fromDocumentEditor).VersionIds(versionIds).UserIds(userIds).TagIds(tagIds).Include(include).FieldsUsers(fieldsUsers).FieldsTasks(fieldsTasks).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsMessages(fieldsMessages).FieldsFileCategories(fieldsFileCategories).FieldsComments(fieldsComments).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.GETProjectsApiV3FilesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3FilesJson`: FileProjectFilesResponse
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.GETProjectsApiV3FilesJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3FilesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uploadedStartDate** | **string** | filter by uploaded start date | 
 **uploadedEndDate** | **string** | filter by uploaded end date | 
 **updatedAfterDate** | **time.Time** | filter by updated after date (deprecated, use updatedAfter) | 
 **updatedAfter** | **time.Time** | filter by updated after date | 
 **status** | **string** | filter by &#39;all&#39;, &#39;active&#39;, &#39;inactive&#39; | 
 **searchTerm** | **string** | filter by file name | 
 **projectType** | **string** | filter by project type. Leave empty to retrieve all types eg. &amp;projectType&#x3D; | [default to &quot;normal&quot;]
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;name&quot;]
 **versionId** | **int32** | filter by version id, returns the latest file | 
 **version** | **int32** | select the version of the file to include when filtering by id. | 
 **taskId** | **int32** | filter by task ids. | 
 **projectId** | **int32** | filter by project id | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **id** | **int32** | filter by project file id | 
 **categoryId** | **int32** | filter by category id | 
 **skipInternalFiles** | **bool** | filter by omitting internal files | 
 **skipExternalFiles** | **bool** | filter by omitting external files | 
 **showDeleted** | **bool** | include deleted files. | 
 **searchAllFields** | **bool** | filter by display name, file extension, file category, original file name and the name of the latest uploader. | 
 **matchAllTags** | **bool** | filter by matching all tag ids | 
 **lockedOnly** | **bool** | filter by files that are locked | 
 **getVersions** | **bool** | include the file versions of the file | 
 **getVersionReactions** | **bool** | include the reactions of each file version | 
 **getRelatedItems** | **bool** | get related tasks, messages, comments for a file. default true when filtering by id. | 
 **getReactions** | **bool** | include reactions of the file | 
 **getLikes** | **bool** | include the number of likes | 
 **getComments** | **bool** | include the number of comments and comments read. default true when filtering by id | 
 **fromDocumentEditor** | **bool** | document editor files | 
 **versionIds** | **[]int32** | filter by version ids, returns the latest files | 
 **userIds** | **[]int32** | filter by user ids | 
 **tagIds** | **[]int32** | filter by files tag ids. | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTasks** | **[]string** |  | 
 **fieldsTags** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsMessages** | **[]string** |  | 
 **fieldsFileCategories** | **[]string** |  | 
 **fieldsComments** | **[]string** |  | 

### Return type

[**FileProjectFilesResponse**](FileProjectFilesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3FilesUsageJson

> FileUsageResponse GETProjectsApiV3FilesUsageJson(ctx).ProjectId(projectId).Include(include).FieldsProjects(fieldsProjects).Execute()

Retrieve file usage on account



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
    projectId := int32(56) // int32 | filter by a specific project (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FilesApi.GETProjectsApiV3FilesUsageJson(context.Background()).ProjectId(projectId).Include(include).FieldsProjects(fieldsProjects).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.GETProjectsApiV3FilesUsageJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3FilesUsageJson`: FileUsageResponse
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.GETProjectsApiV3FilesUsageJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3FilesUsageJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int32** | filter by a specific project | 
 **include** | **[]string** | include | 
 **fieldsProjects** | **[]string** |  | 

### Return type

[**FileUsageResponse**](FileUsageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3FilesfileIdJson

> FileResponse GETProjectsApiV3FilesfileIdJson(ctx, fileId).UploadedStartDate(uploadedStartDate).UploadedEndDate(uploadedEndDate).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).Status(status).SearchTerm(searchTerm).ProjectType(projectType).OrderMode(orderMode).OrderBy(orderBy).VersionId(versionId).Version(version).TaskId(taskId).ProjectId(projectId).PageSize(pageSize).Page(page).Id(id).CategoryId(categoryId).SkipInternalFiles(skipInternalFiles).SkipExternalFiles(skipExternalFiles).ShowDeleted(showDeleted).SearchAllFields(searchAllFields).MatchAllTags(matchAllTags).LockedOnly(lockedOnly).GetVersions(getVersions).GetVersionReactions(getVersionReactions).GetRelatedItems(getRelatedItems).GetReactions(getReactions).GetLikes(getLikes).GetComments(getComments).FromDocumentEditor(fromDocumentEditor).VersionIds(versionIds).UserIds(userIds).TagIds(tagIds).Include(include).FieldsUsers(fieldsUsers).FieldsTasks(fieldsTasks).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsMessages(fieldsMessages).FieldsFileCategories(fieldsFileCategories).FieldsComments(fieldsComments).Execute()

Get a specific file.



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
    fileId := int32(56) // int32 | 
    uploadedStartDate := time.Now() // string | filter by uploaded start date (optional)
    uploadedEndDate := time.Now() // string | filter by uploaded end date (optional)
    updatedAfterDate := time.Now() // time.Time | filter by updated after date (deprecated, use updatedAfter) (optional)
    updatedAfter := time.Now() // time.Time | filter by updated after date (optional)
    status := "status_example" // string | filter by 'all', 'active', 'inactive' (optional)
    searchTerm := "searchTerm_example" // string | filter by file name (optional)
    projectType := "projectType_example" // string | filter by project type. Leave empty to retrieve all types eg. &projectType= (optional) (default to "normal")
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "name")
    versionId := int32(56) // int32 | filter by version id, returns the latest file (optional)
    version := int32(56) // int32 | select the version of the file to include when filtering by id. (optional)
    taskId := int32(56) // int32 | filter by task ids. (optional)
    projectId := int32(56) // int32 | filter by project id (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    id := int32(56) // int32 | filter by project file id (optional)
    categoryId := int32(56) // int32 | filter by category id (optional)
    skipInternalFiles := true // bool | filter by omitting internal files (optional)
    skipExternalFiles := true // bool | filter by omitting external files (optional)
    showDeleted := true // bool | include deleted files. (optional)
    searchAllFields := true // bool | filter by display name, file extension, file category, original file name and the name of the latest uploader. (optional)
    matchAllTags := true // bool | filter by matching all tag ids (optional)
    lockedOnly := true // bool | filter by files that are locked (optional)
    getVersions := true // bool | include the file versions of the file (optional)
    getVersionReactions := true // bool | include the reactions of each file version (optional)
    getRelatedItems := true // bool | get related tasks, messages, comments for a file. default true when filtering by id. (optional)
    getReactions := true // bool | include reactions of the file (optional)
    getLikes := true // bool | include the number of likes (optional)
    getComments := true // bool | include the number of comments and comments read. default true when filtering by id (optional)
    fromDocumentEditor := true // bool | document editor files (optional)
    versionIds := []int32{int32(123)} // []int32 | filter by version ids, returns the latest files (optional)
    userIds := []int32{int32(123)} // []int32 | filter by user ids (optional)
    tagIds := []int32{int32(123)} // []int32 | filter by files tag ids. (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTasks := []string{"Inner_example"} // []string |  (optional)
    fieldsTags := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsMessages := []string{"Inner_example"} // []string |  (optional)
    fieldsFileCategories := []string{"Inner_example"} // []string |  (optional)
    fieldsComments := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FilesApi.GETProjectsApiV3FilesfileIdJson(context.Background(), fileId).UploadedStartDate(uploadedStartDate).UploadedEndDate(uploadedEndDate).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).Status(status).SearchTerm(searchTerm).ProjectType(projectType).OrderMode(orderMode).OrderBy(orderBy).VersionId(versionId).Version(version).TaskId(taskId).ProjectId(projectId).PageSize(pageSize).Page(page).Id(id).CategoryId(categoryId).SkipInternalFiles(skipInternalFiles).SkipExternalFiles(skipExternalFiles).ShowDeleted(showDeleted).SearchAllFields(searchAllFields).MatchAllTags(matchAllTags).LockedOnly(lockedOnly).GetVersions(getVersions).GetVersionReactions(getVersionReactions).GetRelatedItems(getRelatedItems).GetReactions(getReactions).GetLikes(getLikes).GetComments(getComments).FromDocumentEditor(fromDocumentEditor).VersionIds(versionIds).UserIds(userIds).TagIds(tagIds).Include(include).FieldsUsers(fieldsUsers).FieldsTasks(fieldsTasks).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsMessages(fieldsMessages).FieldsFileCategories(fieldsFileCategories).FieldsComments(fieldsComments).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.GETProjectsApiV3FilesfileIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3FilesfileIdJson`: FileResponse
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.GETProjectsApiV3FilesfileIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3FilesfileIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uploadedStartDate** | **string** | filter by uploaded start date | 
 **uploadedEndDate** | **string** | filter by uploaded end date | 
 **updatedAfterDate** | **time.Time** | filter by updated after date (deprecated, use updatedAfter) | 
 **updatedAfter** | **time.Time** | filter by updated after date | 
 **status** | **string** | filter by &#39;all&#39;, &#39;active&#39;, &#39;inactive&#39; | 
 **searchTerm** | **string** | filter by file name | 
 **projectType** | **string** | filter by project type. Leave empty to retrieve all types eg. &amp;projectType&#x3D; | [default to &quot;normal&quot;]
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;name&quot;]
 **versionId** | **int32** | filter by version id, returns the latest file | 
 **version** | **int32** | select the version of the file to include when filtering by id. | 
 **taskId** | **int32** | filter by task ids. | 
 **projectId** | **int32** | filter by project id | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **id** | **int32** | filter by project file id | 
 **categoryId** | **int32** | filter by category id | 
 **skipInternalFiles** | **bool** | filter by omitting internal files | 
 **skipExternalFiles** | **bool** | filter by omitting external files | 
 **showDeleted** | **bool** | include deleted files. | 
 **searchAllFields** | **bool** | filter by display name, file extension, file category, original file name and the name of the latest uploader. | 
 **matchAllTags** | **bool** | filter by matching all tag ids | 
 **lockedOnly** | **bool** | filter by files that are locked | 
 **getVersions** | **bool** | include the file versions of the file | 
 **getVersionReactions** | **bool** | include the reactions of each file version | 
 **getRelatedItems** | **bool** | get related tasks, messages, comments for a file. default true when filtering by id. | 
 **getReactions** | **bool** | include reactions of the file | 
 **getLikes** | **bool** | include the number of likes | 
 **getComments** | **bool** | include the number of comments and comments read. default true when filtering by id | 
 **fromDocumentEditor** | **bool** | document editor files | 
 **versionIds** | **[]int32** | filter by version ids, returns the latest files | 
 **userIds** | **[]int32** | filter by user ids | 
 **tagIds** | **[]int32** | filter by files tag ids. | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTasks** | **[]string** |  | 
 **fieldsTags** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsMessages** | **[]string** |  | 
 **fieldsFileCategories** | **[]string** |  | 
 **fieldsComments** | **[]string** |  | 

### Return type

[**FileResponse**](FileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3FileversionidJson

> FileversionResponse GETProjectsApiV3FileversionidJson(ctx, id2).Id(id).GetReactions(getReactions).Include(include).FieldsUsers(fieldsUsers).FieldsProjects(fieldsProjects).FieldsFiles(fieldsFiles).Execute()

Get a specific fileversion.



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
    id2 := int32(56) // int32 | 
    id := int32(56) // int32 | filter by project file id (optional)
    getReactions := true // bool | include reactions of the file version (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsFiles := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FilesApi.GETProjectsApiV3FileversionidJson(context.Background(), id2).Id(id).GetReactions(getReactions).Include(include).FieldsUsers(fieldsUsers).FieldsProjects(fieldsProjects).FieldsFiles(fieldsFiles).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.GETProjectsApiV3FileversionidJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3FileversionidJson`: FileversionResponse
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.GETProjectsApiV3FileversionidJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3FileversionidJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **int32** | filter by project file id | 
 **getReactions** | **bool** | include reactions of the file version | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsFiles** | **[]string** |  | 

### Return type

[**FileversionResponse**](FileversionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3ProjectsprojectIdFilesUsageJson

> FileUsageResponse GETProjectsApiV3ProjectsprojectIdFilesUsageJson(ctx, projectId2).ProjectId(projectId).Include(include).FieldsProjects(fieldsProjects).Execute()

Retrieve file usage on a project



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
    projectId := int32(56) // int32 | filter by a specific project (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FilesApi.GETProjectsApiV3ProjectsprojectIdFilesUsageJson(context.Background(), projectId2).ProjectId(projectId).Include(include).FieldsProjects(fieldsProjects).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.GETProjectsApiV3ProjectsprojectIdFilesUsageJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ProjectsprojectIdFilesUsageJson`: FileUsageResponse
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.GETProjectsApiV3ProjectsprojectIdFilesUsageJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsprojectIdFilesUsageJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectId** | **int32** | filter by a specific project | 
 **include** | **[]string** | include | 
 **fieldsProjects** | **[]string** |  | 

### Return type

[**FileUsageResponse**](FileUsageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHProjectsApiV3FilesfileIdJson

> FileResponse PATCHProjectsApiV3FilesfileIdJson(ctx, fileId).FileRequest(fileRequest).Execute()

Update an existing file



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
    fileId := int32(56) // int32 | 
    fileRequest := *openapiclient.NewFileRequest() // FileRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FilesApi.PATCHProjectsApiV3FilesfileIdJson(context.Background(), fileId).FileRequest(fileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.PATCHProjectsApiV3FilesfileIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHProjectsApiV3FilesfileIdJson`: FileResponse
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.PATCHProjectsApiV3FilesfileIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHProjectsApiV3FilesfileIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileRequest** | [**FileRequest**](FileRequest.md) |  | 

### Return type

[**FileResponse**](FileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsApiV3FilesArchiveJson

> FileArchiveResponse POSTProjectsApiV3FilesArchiveJson(ctx).FileArchiveRequest(fileArchiveRequest).Execute()

Returns an URL for multiple files



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
    fileArchiveRequest := *openapiclient.NewFileArchiveRequest() // FileArchiveRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FilesApi.POSTProjectsApiV3FilesArchiveJson(context.Background()).FileArchiveRequest(fileArchiveRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.POSTProjectsApiV3FilesArchiveJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsApiV3FilesArchiveJson`: FileArchiveResponse
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.POSTProjectsApiV3FilesArchiveJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3FilesArchiveJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileArchiveRequest** | [**FileArchiveRequest**](FileArchiveRequest.md) |  | 

### Return type

[**FileArchiveResponse**](FileArchiveResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsApiV3FilesprojectFileIdJson

> FileversionResponse POSTProjectsApiV3FilesprojectFileIdJson(ctx, projectFileId).FileversionRequest(fileversionRequest).Execute()

Create a new fileversion for the project file id.



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
    projectFileId := int32(56) // int32 | 
    fileversionRequest := *openapiclient.NewFileversionRequest() // FileversionRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FilesApi.POSTProjectsApiV3FilesprojectFileIdJson(context.Background(), projectFileId).FileversionRequest(fileversionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.POSTProjectsApiV3FilesprojectFileIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsApiV3FilesprojectFileIdJson`: FileversionResponse
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.POSTProjectsApiV3FilesprojectFileIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectFileId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3FilesprojectFileIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileversionRequest** | [**FileversionRequest**](FileversionRequest.md) |  | 

### Return type

[**FileversionResponse**](FileversionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

