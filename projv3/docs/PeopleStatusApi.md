# \PeopleStatusApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3StatusesJson**](PeopleStatusApi.md#GETProjectsApiV3StatusesJson) | **Get** /projects/api/v3/statuses.json | Get all statuses
[**GETProjectsApiV3StatusesTimelineJson**](PeopleStatusApi.md#GETProjectsApiV3StatusesTimelineJson) | **Get** /projects/api/v3/statuses/timeline.json | Get the people statuses timeline.
[**GETProjectsApiV3TeamsteamIdStatusesTimelineJson**](PeopleStatusApi.md#GETProjectsApiV3TeamsteamIdStatusesTimelineJson) | **Get** /projects/api/v3/teams/{teamId}/statuses/timeline.json | Get statuses timeline for a specific team



## GETProjectsApiV3StatusesJson

> StatusTimelineResponse GETProjectsApiV3StatusesJson(ctx).UpdatedAfter(updatedAfter).SearchTerm(searchTerm).TeamId(teamId).PageSize(pageSize).Page(page).ShowEveryone(showEveryone).ShowDeleted(showDeleted).IncludeSubteams(includeSubteams).Include(include).FieldsUsers(fieldsUsers).FieldsStatuses(fieldsStatuses).Execute()

Get all statuses



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
    updatedAfter := time.Now() // time.Time | return only statuses updated after a specific date (optional)
    searchTerm := "searchTerm_example" // string | filter by the user name or status message (optional)
    teamId := int32(56) // int32 | filter by team (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 100)
    page := int32(56) // int32 | page number (optional) (default to 1)
    showEveryone := true // bool | show also users without status (optional)
    showDeleted := true // bool | showDeleted statuses (optional) (default to false)
    includeSubteams := true // bool | also include statuses from subteams (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsStatuses := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeopleStatusApi.GETProjectsApiV3StatusesJson(context.Background()).UpdatedAfter(updatedAfter).SearchTerm(searchTerm).TeamId(teamId).PageSize(pageSize).Page(page).ShowEveryone(showEveryone).ShowDeleted(showDeleted).IncludeSubteams(includeSubteams).Include(include).FieldsUsers(fieldsUsers).FieldsStatuses(fieldsStatuses).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleStatusApi.GETProjectsApiV3StatusesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3StatusesJson`: StatusTimelineResponse
    fmt.Fprintf(os.Stdout, "Response from `PeopleStatusApi.GETProjectsApiV3StatusesJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3StatusesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfter** | **time.Time** | return only statuses updated after a specific date | 
 **searchTerm** | **string** | filter by the user name or status message | 
 **teamId** | **int32** | filter by team | 
 **pageSize** | **int32** | number of items in a page | [default to 100]
 **page** | **int32** | page number | [default to 1]
 **showEveryone** | **bool** | show also users without status | 
 **showDeleted** | **bool** | showDeleted statuses | [default to false]
 **includeSubteams** | **bool** | also include statuses from subteams | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsStatuses** | **[]string** |  | 

### Return type

[**StatusTimelineResponse**](StatusTimelineResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3StatusesTimelineJson

> StatusTimelineResponse GETProjectsApiV3StatusesTimelineJson(ctx).SearchTerm(searchTerm).TeamId(teamId).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).IncludeSubteams(includeSubteams).Include(include).FieldsUsers(fieldsUsers).Execute()

Get the people statuses timeline.



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
    searchTerm := "searchTerm_example" // string | filter by the user name or status message (optional)
    teamId := int32(56) // int32 | filter by team (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    showDeleted := true // bool | showDeleted statuses (optional) (default to false)
    includeSubteams := true // bool | also include statuses from subteams (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeopleStatusApi.GETProjectsApiV3StatusesTimelineJson(context.Background()).SearchTerm(searchTerm).TeamId(teamId).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).IncludeSubteams(includeSubteams).Include(include).FieldsUsers(fieldsUsers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleStatusApi.GETProjectsApiV3StatusesTimelineJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3StatusesTimelineJson`: StatusTimelineResponse
    fmt.Fprintf(os.Stdout, "Response from `PeopleStatusApi.GETProjectsApiV3StatusesTimelineJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3StatusesTimelineJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchTerm** | **string** | filter by the user name or status message | 
 **teamId** | **int32** | filter by team | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **showDeleted** | **bool** | showDeleted statuses | [default to false]
 **includeSubteams** | **bool** | also include statuses from subteams | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 

### Return type

[**StatusTimelineResponse**](StatusTimelineResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3TeamsteamIdStatusesTimelineJson

> StatusTimelineResponse GETProjectsApiV3TeamsteamIdStatusesTimelineJson(ctx, teamId2).SearchTerm(searchTerm).TeamId(teamId).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).IncludeSubteams(includeSubteams).Include(include).FieldsUsers(fieldsUsers).Execute()

Get statuses timeline for a specific team



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
    teamId2 := int32(56) // int32 | 
    searchTerm := "searchTerm_example" // string | filter by the user name or status message (optional)
    teamId := int32(56) // int32 | filter by team (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    showDeleted := true // bool | showDeleted statuses (optional) (default to false)
    includeSubteams := true // bool | also include statuses from subteams (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeopleStatusApi.GETProjectsApiV3TeamsteamIdStatusesTimelineJson(context.Background(), teamId2).SearchTerm(searchTerm).TeamId(teamId).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).IncludeSubteams(includeSubteams).Include(include).FieldsUsers(fieldsUsers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleStatusApi.GETProjectsApiV3TeamsteamIdStatusesTimelineJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3TeamsteamIdStatusesTimelineJson`: StatusTimelineResponse
    fmt.Fprintf(os.Stdout, "Response from `PeopleStatusApi.GETProjectsApiV3TeamsteamIdStatusesTimelineJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3TeamsteamIdStatusesTimelineJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchTerm** | **string** | filter by the user name or status message | 
 **teamId** | **int32** | filter by team | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **showDeleted** | **bool** | showDeleted statuses | [default to false]
 **includeSubteams** | **bool** | also include statuses from subteams | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 

### Return type

[**StatusTimelineResponse**](StatusTimelineResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

