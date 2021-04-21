# \AllocationsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEProjectsApiV3AllocationsallocationIdJson**](AllocationsApi.md#DELETEProjectsApiV3AllocationsallocationIdJson) | **Delete** /projects/api/v3/allocations/{allocationId}.json | Delete an existing allocation.
[**GETProjectsApiV3AllocationsJson**](AllocationsApi.md#GETProjectsApiV3AllocationsJson) | **Get** /projects/api/v3/allocations.json | Get all allocations.
[**GETProjectsApiV3AllocationsallocationIdJson**](AllocationsApi.md#GETProjectsApiV3AllocationsallocationIdJson) | **Get** /projects/api/v3/allocations/{allocationId}.json | Get a specific allocation.
[**PATCHProjectsApiV3AllocationsallocationIdJson**](AllocationsApi.md#PATCHProjectsApiV3AllocationsallocationIdJson) | **Patch** /projects/api/v3/allocations/{allocationId}.json | Update an existing allocation.
[**POSTProjectsApiV3AllocationsBulkDeleteJson**](AllocationsApi.md#POSTProjectsApiV3AllocationsBulkDeleteJson) | **Post** /projects/api/v3/allocations/bulk/delete.json | Delete many allocations at once.
[**POSTProjectsApiV3AllocationsJson**](AllocationsApi.md#POSTProjectsApiV3AllocationsJson) | **Post** /projects/api/v3/allocations.json | Create a new allocation.
[**PUTProjectsApiV3AllocationsallocationIdJson**](AllocationsApi.md#PUTProjectsApiV3AllocationsallocationIdJson) | **Put** /projects/api/v3/allocations/{allocationId}.json | Update an existing allocation.
[**PUTProjectsApiV3AllocationsallocationIdRestoreJson**](AllocationsApi.md#PUTProjectsApiV3AllocationsallocationIdRestoreJson) | **Put** /projects/api/v3/allocations/{allocationId}/restore.json | Restore a soft deleted allocation.
[**PUTProjectsApiV3AllocationsallocationIdSplitJson**](AllocationsApi.md#PUTProjectsApiV3AllocationsallocationIdSplitJson) | **Put** /projects/api/v3/allocations/{allocationId}/split.json | Split an allocation in two



## DELETEProjectsApiV3AllocationsallocationIdJson

> DELETEProjectsApiV3AllocationsallocationIdJson(ctx, allocationId).AllocationDeleteRequest(allocationDeleteRequest).Execute()

Delete an existing allocation.



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
    allocationId := int32(56) // int32 | 
    allocationDeleteRequest := *openapiclient.NewAllocationDeleteRequest() // AllocationDeleteRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllocationsApi.DELETEProjectsApiV3AllocationsallocationIdJson(context.Background(), allocationId).AllocationDeleteRequest(allocationDeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllocationsApi.DELETEProjectsApiV3AllocationsallocationIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allocationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEProjectsApiV3AllocationsallocationIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **allocationDeleteRequest** | [**AllocationDeleteRequest**](AllocationDeleteRequest.md) |  | 

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


## GETProjectsApiV3AllocationsJson

> AllocationAllocationsResponse GETProjectsApiV3AllocationsJson(ctx).StartDate(startDate).SearchTerm(searchTerm).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).EndDate(endDate).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).OnlyStarredProjects(onlyStarredProjects).OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership).MatchAllProjectTags(matchAllProjectTags).HideObservedProjects(hideObservedProjects).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsProjects(fieldsProjects).FieldsCompanies(fieldsCompanies).FieldsAllocations(fieldsAllocations).AssignedUserIds(assignedUserIds).Execute()

Get all allocations.



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
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)
    fieldsAllocations := []string{"Inner_example"} // []string |  (optional)
    assignedUserIds := []int32{int32(123)} // []int32 | filter by assigned user ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllocationsApi.GETProjectsApiV3AllocationsJson(context.Background()).StartDate(startDate).SearchTerm(searchTerm).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).EndDate(endDate).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).OnlyStarredProjects(onlyStarredProjects).OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership).MatchAllProjectTags(matchAllProjectTags).HideObservedProjects(hideObservedProjects).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsProjects(fieldsProjects).FieldsCompanies(fieldsCompanies).FieldsAllocations(fieldsAllocations).AssignedUserIds(assignedUserIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllocationsApi.GETProjectsApiV3AllocationsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3AllocationsJson`: AllocationAllocationsResponse
    fmt.Fprintf(os.Stdout, "Response from `AllocationsApi.GETProjectsApiV3AllocationsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3AllocationsJsonRequest struct via the builder pattern


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
 **fieldsProjects** | **[]string** |  | 
 **fieldsCompanies** | **[]string** |  | 
 **fieldsAllocations** | **[]string** |  | 
 **assignedUserIds** | **[]int32** | filter by assigned user ids | 

### Return type

[**AllocationAllocationsResponse**](AllocationAllocationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3AllocationsallocationIdJson

> AllocationResponse GETProjectsApiV3AllocationsallocationIdJson(ctx, allocationId).Execute()

Get a specific allocation.



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
    allocationId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllocationsApi.GETProjectsApiV3AllocationsallocationIdJson(context.Background(), allocationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllocationsApi.GETProjectsApiV3AllocationsallocationIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3AllocationsallocationIdJson`: AllocationResponse
    fmt.Fprintf(os.Stdout, "Response from `AllocationsApi.GETProjectsApiV3AllocationsallocationIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allocationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3AllocationsallocationIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AllocationResponse**](AllocationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHProjectsApiV3AllocationsallocationIdJson

> AllocationResponse PATCHProjectsApiV3AllocationsallocationIdJson(ctx, allocationId).AllocationRequest(allocationRequest).Execute()

Update an existing allocation.



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
    allocationId := int32(56) // int32 | 
    allocationRequest := *openapiclient.NewAllocationRequest() // AllocationRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllocationsApi.PATCHProjectsApiV3AllocationsallocationIdJson(context.Background(), allocationId).AllocationRequest(allocationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllocationsApi.PATCHProjectsApiV3AllocationsallocationIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHProjectsApiV3AllocationsallocationIdJson`: AllocationResponse
    fmt.Fprintf(os.Stdout, "Response from `AllocationsApi.PATCHProjectsApiV3AllocationsallocationIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allocationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHProjectsApiV3AllocationsallocationIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **allocationRequest** | [**AllocationRequest**](AllocationRequest.md) |  | 

### Return type

[**AllocationResponse**](AllocationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsApiV3AllocationsBulkDeleteJson

> POSTProjectsApiV3AllocationsBulkDeleteJson(ctx).AllocationBulkDeleteRequest(allocationBulkDeleteRequest).Execute()

Delete many allocations at once.



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
    allocationBulkDeleteRequest := *openapiclient.NewAllocationBulkDeleteRequest() // AllocationBulkDeleteRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllocationsApi.POSTProjectsApiV3AllocationsBulkDeleteJson(context.Background()).AllocationBulkDeleteRequest(allocationBulkDeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllocationsApi.POSTProjectsApiV3AllocationsBulkDeleteJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3AllocationsBulkDeleteJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allocationBulkDeleteRequest** | [**AllocationBulkDeleteRequest**](AllocationBulkDeleteRequest.md) |  | 

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


## POSTProjectsApiV3AllocationsJson

> AllocationResponse POSTProjectsApiV3AllocationsJson(ctx).AllocationRequest(allocationRequest).Execute()

Create a new allocation.



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
    allocationRequest := *openapiclient.NewAllocationRequest() // AllocationRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllocationsApi.POSTProjectsApiV3AllocationsJson(context.Background()).AllocationRequest(allocationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllocationsApi.POSTProjectsApiV3AllocationsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsApiV3AllocationsJson`: AllocationResponse
    fmt.Fprintf(os.Stdout, "Response from `AllocationsApi.POSTProjectsApiV3AllocationsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3AllocationsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allocationRequest** | [**AllocationRequest**](AllocationRequest.md) |  | 

### Return type

[**AllocationResponse**](AllocationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV3AllocationsallocationIdJson

> AllocationResponse PUTProjectsApiV3AllocationsallocationIdJson(ctx, allocationId).AllocationRequest(allocationRequest).Execute()

Update an existing allocation.



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
    allocationId := int32(56) // int32 | 
    allocationRequest := *openapiclient.NewAllocationRequest() // AllocationRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllocationsApi.PUTProjectsApiV3AllocationsallocationIdJson(context.Background(), allocationId).AllocationRequest(allocationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllocationsApi.PUTProjectsApiV3AllocationsallocationIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsApiV3AllocationsallocationIdJson`: AllocationResponse
    fmt.Fprintf(os.Stdout, "Response from `AllocationsApi.PUTProjectsApiV3AllocationsallocationIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allocationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3AllocationsallocationIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **allocationRequest** | [**AllocationRequest**](AllocationRequest.md) |  | 

### Return type

[**AllocationResponse**](AllocationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV3AllocationsallocationIdRestoreJson

> AllocationResponse PUTProjectsApiV3AllocationsallocationIdRestoreJson(ctx, allocationId).Execute()

Restore a soft deleted allocation.



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
    allocationId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllocationsApi.PUTProjectsApiV3AllocationsallocationIdRestoreJson(context.Background(), allocationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllocationsApi.PUTProjectsApiV3AllocationsallocationIdRestoreJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsApiV3AllocationsallocationIdRestoreJson`: AllocationResponse
    fmt.Fprintf(os.Stdout, "Response from `AllocationsApi.PUTProjectsApiV3AllocationsallocationIdRestoreJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allocationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3AllocationsallocationIdRestoreJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AllocationResponse**](AllocationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV3AllocationsallocationIdSplitJson

> AllocationAllocationsResponse PUTProjectsApiV3AllocationsallocationIdSplitJson(ctx, allocationId).AllocationSplitRequest(allocationSplitRequest).Execute()

Split an allocation in two



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
    allocationId := int32(56) // int32 | 
    allocationSplitRequest := *openapiclient.NewAllocationSplitRequest() // AllocationSplitRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllocationsApi.PUTProjectsApiV3AllocationsallocationIdSplitJson(context.Background(), allocationId).AllocationSplitRequest(allocationSplitRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllocationsApi.PUTProjectsApiV3AllocationsallocationIdSplitJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsApiV3AllocationsallocationIdSplitJson`: AllocationAllocationsResponse
    fmt.Fprintf(os.Stdout, "Response from `AllocationsApi.PUTProjectsApiV3AllocationsallocationIdSplitJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allocationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3AllocationsallocationIdSplitJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **allocationSplitRequest** | [**AllocationSplitRequest**](AllocationSplitRequest.md) |  | 

### Return type

[**AllocationAllocationsResponse**](AllocationAllocationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

