# \TagsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEProjectsApiV3TagstagIdJson**](TagsApi.md#DELETEProjectsApiV3TagstagIdJson) | **Delete** /projects/api/v3/tags/{tagId}.json | Deletes a tag
[**GETProjectsApiV3TagsJson**](TagsApi.md#GETProjectsApiV3TagsJson) | **Get** /projects/api/v3/tags.json | Get all tags
[**GETProjectsApiV3TagstagIdJson**](TagsApi.md#GETProjectsApiV3TagstagIdJson) | **Get** /projects/api/v3/tags/{tagId}.json | Get details of an individual tag
[**PATCHProjectsApiV3TagstagIdJson**](TagsApi.md#PATCHProjectsApiV3TagstagIdJson) | **Patch** /projects/api/v3/tags/{tagId}.json | Update a single tag
[**POSTProjectsApiV3TagsBulkDeleteJson**](TagsApi.md#POSTProjectsApiV3TagsBulkDeleteJson) | **Post** /projects/api/v3/tags/bulk/delete.json | Delete many tags at once
[**POSTProjectsApiV3TagsJson**](TagsApi.md#POSTProjectsApiV3TagsJson) | **Post** /projects/api/v3/tags.json | Create a single tag



## DELETEProjectsApiV3TagstagIdJson

> DELETEProjectsApiV3TagstagIdJson(ctx, tagId).Execute()

Deletes a tag



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
    tagId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagsApi.DELETEProjectsApiV3TagstagIdJson(context.Background(), tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.DELETEProjectsApiV3TagstagIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEProjectsApiV3TagstagIdJsonRequest struct via the builder pattern


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


## GETProjectsApiV3TagsJson

> TagTagsResponse GETProjectsApiV3TagsJson(ctx).UpdatedAfter(updatedAfter).SearchTerm(searchTerm).OrderMode(orderMode).OrderBy(orderBy).ItemType(itemType).Filter(filter).PageSize(pageSize).Page(page).WithCounters(withCounters).SkipSpecial(skipSpecial).SearchRightOnly(searchRightOnly).SkipIds(skipIds).ProjectIds(projectIds).FieldsTags(fieldsTags).Execute()

Get all tags



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
    updatedAfter := time.Now() // time.Time | search for tags updated after the provided date (optional)
    searchTerm := "searchTerm_example" // string | filter by search term (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "name")
    itemType := "itemType_example" // string | filter by item type (optional) (default to "all")
    filter := "filter_example" // string | mode used when filtering the tags (optional) (default to "all")
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 500)
    page := int32(56) // int32 | page number (optional) (default to 1)
    withCounters := true // bool | include in the response the number of items that use the tag (optional) (default to false)
    skipSpecial := true // bool | do not include in the response special tags (optional) (default to true)
    searchRightOnly := true // bool | search term will be placed as a prefix to match the tag names (optional) (default to false)
    skipIds := []int32{int32(123)} // []int32 | skip from the result tags with the defined ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by projects (optional)
    fieldsTags := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagsApi.GETProjectsApiV3TagsJson(context.Background()).UpdatedAfter(updatedAfter).SearchTerm(searchTerm).OrderMode(orderMode).OrderBy(orderBy).ItemType(itemType).Filter(filter).PageSize(pageSize).Page(page).WithCounters(withCounters).SkipSpecial(skipSpecial).SearchRightOnly(searchRightOnly).SkipIds(skipIds).ProjectIds(projectIds).FieldsTags(fieldsTags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.GETProjectsApiV3TagsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3TagsJson`: TagTagsResponse
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.GETProjectsApiV3TagsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3TagsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfter** | **time.Time** | search for tags updated after the provided date | 
 **searchTerm** | **string** | filter by search term | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;name&quot;]
 **itemType** | **string** | filter by item type | [default to &quot;all&quot;]
 **filter** | **string** | mode used when filtering the tags | [default to &quot;all&quot;]
 **pageSize** | **int32** | number of items in a page | [default to 500]
 **page** | **int32** | page number | [default to 1]
 **withCounters** | **bool** | include in the response the number of items that use the tag | [default to false]
 **skipSpecial** | **bool** | do not include in the response special tags | [default to true]
 **searchRightOnly** | **bool** | search term will be placed as a prefix to match the tag names | [default to false]
 **skipIds** | **[]int32** | skip from the result tags with the defined ids | 
 **projectIds** | **[]int32** | filter by projects | 
 **fieldsTags** | **[]string** |  | 

### Return type

[**TagTagsResponse**](TagTagsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3TagstagIdJson

> TagResponse GETProjectsApiV3TagstagIdJson(ctx, tagId).Execute()

Get details of an individual tag



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
    tagId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagsApi.GETProjectsApiV3TagstagIdJson(context.Background(), tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.GETProjectsApiV3TagstagIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3TagstagIdJson`: TagResponse
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.GETProjectsApiV3TagstagIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3TagstagIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TagResponse**](TagResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHProjectsApiV3TagstagIdJson

> TagResponse PATCHProjectsApiV3TagstagIdJson(ctx, tagId).TagRequest(tagRequest).Execute()

Update a single tag



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
    tagId := int32(56) // int32 | 
    tagRequest := *openapiclient.NewTagRequest() // TagRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagsApi.PATCHProjectsApiV3TagstagIdJson(context.Background(), tagId).TagRequest(tagRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.PATCHProjectsApiV3TagstagIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHProjectsApiV3TagstagIdJson`: TagResponse
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.PATCHProjectsApiV3TagstagIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHProjectsApiV3TagstagIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagRequest** | [**TagRequest**](TagRequest.md) |  | 

### Return type

[**TagResponse**](TagResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsApiV3TagsBulkDeleteJson

> POSTProjectsApiV3TagsBulkDeleteJson(ctx).TagBulkDeleteRequest(tagBulkDeleteRequest).Execute()

Delete many tags at once



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
    tagBulkDeleteRequest := *openapiclient.NewTagBulkDeleteRequest() // TagBulkDeleteRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagsApi.POSTProjectsApiV3TagsBulkDeleteJson(context.Background()).TagBulkDeleteRequest(tagBulkDeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.POSTProjectsApiV3TagsBulkDeleteJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3TagsBulkDeleteJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagBulkDeleteRequest** | [**TagBulkDeleteRequest**](TagBulkDeleteRequest.md) |  | 

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


## POSTProjectsApiV3TagsJson

> TagResponse POSTProjectsApiV3TagsJson(ctx).TagRequest(tagRequest).Execute()

Create a single tag



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
    tagRequest := *openapiclient.NewTagRequest() // TagRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagsApi.POSTProjectsApiV3TagsJson(context.Background()).TagRequest(tagRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.POSTProjectsApiV3TagsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsApiV3TagsJson`: TagResponse
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.POSTProjectsApiV3TagsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3TagsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagRequest** | [**TagRequest**](TagRequest.md) |  | 

### Return type

[**TagResponse**](TagResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

