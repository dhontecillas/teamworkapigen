# \SearchApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETSearchJson**](SearchApi.md#GETSearchJson) | **Get** /search.json | Search



## GETSearchJson

> InlineResponse20089 GETSearchJson(ctx).SearchFor(searchFor).SearchTerm(searchTerm).ProjectId(projectId).SortOrder(sortOrder).IncludeArchivedProjects(includeArchivedProjects).IncludeCompletedItems(includeCompletedItems).PageSize(pageSize).Execute()

Search



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
    searchFor := "searchFor_example" // string | 
    searchTerm := "searchTerm_example" // string | 
    projectId := "projectId_example" // string | The ProjectID you would like to search within (optional)
    sortOrder := "sortOrder_example" // string | 'name' or 'dateupdated' to decide the order of the search results (optional)
    includeArchivedProjects := "includeArchivedProjects_example" // string | True or False to include archived items or not. (optional)
    includeCompletedItems := "includeCompletedItems_example" // string | True or False to include Completed items or not (optional)
    pageSize := "pageSize_example" // string | Define number of results to show, for example 20 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SearchApi.GETSearchJson(context.Background()).SearchFor(searchFor).SearchTerm(searchTerm).ProjectId(projectId).SortOrder(sortOrder).IncludeArchivedProjects(includeArchivedProjects).IncludeCompletedItems(includeCompletedItems).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.GETSearchJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETSearchJson`: InlineResponse20089
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.GETSearchJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETSearchJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchFor** | **string** |  | 
 **searchTerm** | **string** |  | 
 **projectId** | **string** | The ProjectID you would like to search within | 
 **sortOrder** | **string** | &#39;name&#39; or &#39;dateupdated&#39; to decide the order of the search results | 
 **includeArchivedProjects** | **string** | True or False to include archived items or not. | 
 **includeCompletedItems** | **string** | True or False to include Completed items or not | 
 **pageSize** | **string** | Define number of results to show, for example 20 | 

### Return type

[**InlineResponse20089**](InlineResponse20089.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

