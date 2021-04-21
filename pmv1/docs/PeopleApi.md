# \PeopleApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEPeopleIdJson**](PeopleApi.md#DELETEPeopleIdJson) | **Delete** /people/{id}.json | Delete User
[**GETCalendareventsEventIdAvailablePeopleJson**](PeopleApi.md#GETCalendareventsEventIdAvailablePeopleJson) | **Get** /calendarevents/{eventId}/availablePeople.json | Get available People for a Calendar Event
[**GETCompaniesIdPeopleJson**](PeopleApi.md#GETCompaniesIdPeopleJson) | **Get** /companies/{id}/people.json | Get People (within a Company)
[**GETMeJson**](PeopleApi.md#GETMeJson) | **Get** /me.json | Get Current User Details
[**GETPeopleAPIKeysJson**](PeopleApi.md#GETPeopleAPIKeysJson) | **Get** /people/APIKeys.json | Retrieve all API Keys for all People on account
[**GETPeopleDeletedJson**](PeopleApi.md#GETPeopleDeletedJson) | **Get** /people/deleted.json | Get all deleted People
[**GETPeopleIdJson**](PeopleApi.md#GETPeopleIdJson) | **Get** /people/{id}.json | Retrieve a Specific Person
[**GETPeopleJson**](PeopleApi.md#GETPeopleJson) | **Get** /people.json | Get All People
[**GETProjectsProjectIdFilesAvailablepeopleJson**](PeopleApi.md#GETProjectsProjectIdFilesAvailablepeopleJson) | **Get** /projects/{projectId}/files/availablepeople.json | Get available People to notify when adding a File
[**GETProjectsProjectIdLinksAvailablepeopleJson**](PeopleApi.md#GETProjectsProjectIdLinksAvailablepeopleJson) | **Get** /projects/{projectId}/links/availablepeople.json | Get available People to notify when adding a Link
[**GETProjectsProjectIdMessagesAvailablepeopleJson**](PeopleApi.md#GETProjectsProjectIdMessagesAvailablepeopleJson) | **Get** /projects/{projectId}/messages/availablepeople.json | Get available People for a Message
[**GETProjectsProjectIdMilestonesAvailablepeopleJson**](PeopleApi.md#GETProjectsProjectIdMilestonesAvailablepeopleJson) | **Get** /projects/{projectId}/milestones/availablepeople.json | Get available People for a Milestone
[**GETProjectsProjectIdNotebooksAvailablepeopleJson**](PeopleApi.md#GETProjectsProjectIdNotebooksAvailablepeopleJson) | **Get** /projects/{projectId}/notebooks/availablepeople.json | Get available People for following a Notebook
[**GETProjectsProjectIdTasksAvailablepeopleJson**](PeopleApi.md#GETProjectsProjectIdTasksAvailablepeopleJson) | **Get** /projects/{projectId}/tasks/availablepeople.json | Get available People for a Task
[**GETProjectsProjectidPeopleJson**](PeopleApi.md#GETProjectsProjectidPeopleJson) | **Get** /projects/{projectid}/people.json | Get all People (within a Project)
[**GETStatsJson**](PeopleApi.md#GETStatsJson) | **Get** /stats.json | Current User Summary Stats
[**POSTPeopleJson**](PeopleApi.md#POSTPeopleJson) | **Post** /people.json | Creates a new User Account
[**POSTProjectsIdPeopleJson**](PeopleApi.md#POSTProjectsIdPeopleJson) | **Post** /projects/{id}/people.json | Add/Remove People to existing Project
[**PUTPeopleIdJson**](PeopleApi.md#PUTPeopleIdJson) | **Put** /people/{id}.json | Editing a User



## DELETEPeopleIdJson

> InlineResponse2001 DELETEPeopleIdJson(ctx, id).Execute()

Delete User



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
    id := "id_example" // string | Id of the person.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeopleApi.DELETEPeopleIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.DELETEPeopleIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DELETEPeopleIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `PeopleApi.DELETEPeopleIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the person. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEPeopleIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETCalendareventsEventIdAvailablePeopleJson

> InlineResponse2004 GETCalendareventsEventIdAvailablePeopleJson(ctx, eventId).Privacy(privacy).Execute()

Get available People for a Calendar Event



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
    eventId := "eventId_example" // string | 
    privacy := "privacy_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeopleApi.GETCalendareventsEventIdAvailablePeopleJson(context.Background(), eventId).Privacy(privacy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.GETCalendareventsEventIdAvailablePeopleJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCalendareventsEventIdAvailablePeopleJson`: InlineResponse2004
    fmt.Fprintf(os.Stdout, "Response from `PeopleApi.GETCalendareventsEventIdAvailablePeopleJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCalendareventsEventIdAvailablePeopleJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **privacy** | **string** |  | 

### Return type

[**InlineResponse2004**](InlineResponse2004.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETCompaniesIdPeopleJson

> InlineResponse20013 GETCompaniesIdPeopleJson(ctx, id).Execute()

Get People (within a Company)



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
    id := "id_example" // string | Id of the company

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeopleApi.GETCompaniesIdPeopleJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.GETCompaniesIdPeopleJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCompaniesIdPeopleJson`: InlineResponse20013
    fmt.Fprintf(os.Stdout, "Response from `PeopleApi.GETCompaniesIdPeopleJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the company | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCompaniesIdPeopleJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20013**](InlineResponse20013.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETMeJson

> InlineResponse20027 GETMeJson(ctx).GetPreferences(getPreferences).FullProfile(fullProfile).GetDefaultFilters(getDefaultFilters).SharedFilter(sharedFilter).GetAccounts(getAccounts).IncludeAuth(includeAuth).Execute()

Get Current User Details



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
    getPreferences := true // bool | Get user preferences in teamwork projects, eg. ordering of comments. (optional)
    fullProfile := true // bool | Get additional user details. (optional)
    getDefaultFilters := true // bool |  (optional)
    sharedFilter := true // bool |  (optional)
    getAccounts := true // bool |  (optional)
    includeAuth := true // bool | Includes API key, please be careful sharing this data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeopleApi.GETMeJson(context.Background()).GetPreferences(getPreferences).FullProfile(fullProfile).GetDefaultFilters(getDefaultFilters).SharedFilter(sharedFilter).GetAccounts(getAccounts).IncludeAuth(includeAuth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.GETMeJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETMeJson`: InlineResponse20027
    fmt.Fprintf(os.Stdout, "Response from `PeopleApi.GETMeJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETMeJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPreferences** | **bool** | Get user preferences in teamwork projects, eg. ordering of comments. | 
 **fullProfile** | **bool** | Get additional user details. | 
 **getDefaultFilters** | **bool** |  | 
 **sharedFilter** | **bool** |  | 
 **getAccounts** | **bool** |  | 
 **includeAuth** | **bool** | Includes API key, please be careful sharing this data. | 

### Return type

[**InlineResponse20027**](InlineResponse20027.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPeopleAPIKeysJson

> InlineResponse20037 GETPeopleAPIKeysJson(ctx).Execute()

Retrieve all API Keys for all People on account



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
    resp, r, err := api_client.PeopleApi.GETPeopleAPIKeysJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.GETPeopleAPIKeysJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPeopleAPIKeysJson`: InlineResponse20037
    fmt.Fprintf(os.Stdout, "Response from `PeopleApi.GETPeopleAPIKeysJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETPeopleAPIKeysJsonRequest struct via the builder pattern


### Return type

[**InlineResponse20037**](InlineResponse20037.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPeopleDeletedJson

> InlineResponse20038 GETPeopleDeletedJson(ctx).GroupByCompany(groupByCompany).Execute()

Get all deleted People



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
    groupByCompany := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeopleApi.GETPeopleDeletedJson(context.Background()).GroupByCompany(groupByCompany).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.GETPeopleDeletedJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPeopleDeletedJson`: InlineResponse20038
    fmt.Fprintf(os.Stdout, "Response from `PeopleApi.GETPeopleDeletedJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETPeopleDeletedJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupByCompany** | **bool** |  | 

### Return type

[**InlineResponse20038**](InlineResponse20038.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPeopleIdJson

> InlineResponse20039 GETPeopleIdJson(ctx, id).Execute()

Retrieve a Specific Person



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
    id := "id_example" // string | Id of the person.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeopleApi.GETPeopleIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.GETPeopleIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPeopleIdJson`: InlineResponse20039
    fmt.Fprintf(os.Stdout, "Response from `PeopleApi.GETPeopleIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the person. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPeopleIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20039**](InlineResponse20039.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPeopleJson

> InlineResponse20013 GETPeopleJson(ctx).ProjectId(projectId).Page(page).PageSize(pageSize).GroupByCompany(groupByCompany).Sort(sort).SortOrder(sortOrder).Fullprofile(fullprofile).Execute()

Get All People



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
    projectId := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    groupByCompany := true // bool |  (optional)
    sort := "sort_example" // string |  (optional) (default to "first name")
    sortOrder := "sortOrder_example" // string | options are: lastactive, lastName, title, dateAdded. (optional) (default to "asc")
    fullprofile := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeopleApi.GETPeopleJson(context.Background()).ProjectId(projectId).Page(page).PageSize(pageSize).GroupByCompany(groupByCompany).Sort(sort).SortOrder(sortOrder).Fullprofile(fullprofile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.GETPeopleJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPeopleJson`: InlineResponse20013
    fmt.Fprintf(os.Stdout, "Response from `PeopleApi.GETPeopleJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETPeopleJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int32** |  | 
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **groupByCompany** | **bool** |  | 
 **sort** | **string** |  | [default to &quot;first name&quot;]
 **sortOrder** | **string** | options are: lastactive, lastName, title, dateAdded. | [default to &quot;asc&quot;]
 **fullprofile** | **bool** |  | 

### Return type

[**InlineResponse20013**](InlineResponse20013.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsProjectIdFilesAvailablepeopleJson

> InlineResponse20081 GETProjectsProjectIdFilesAvailablepeopleJson(ctx, projectId).Execute()

Get available People to notify when adding a File

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
    projectId := "projectId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeopleApi.GETProjectsProjectIdFilesAvailablepeopleJson(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.GETProjectsProjectIdFilesAvailablepeopleJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsProjectIdFilesAvailablepeopleJson`: InlineResponse20081
    fmt.Fprintf(os.Stdout, "Response from `PeopleApi.GETProjectsProjectIdFilesAvailablepeopleJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsProjectIdFilesAvailablepeopleJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20081**](InlineResponse20081.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsProjectIdLinksAvailablepeopleJson

> InlineResponse20081 GETProjectsProjectIdLinksAvailablepeopleJson(ctx, projectId).Execute()

Get available People to notify when adding a Link

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
    projectId := "projectId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeopleApi.GETProjectsProjectIdLinksAvailablepeopleJson(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.GETProjectsProjectIdLinksAvailablepeopleJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsProjectIdLinksAvailablepeopleJson`: InlineResponse20081
    fmt.Fprintf(os.Stdout, "Response from `PeopleApi.GETProjectsProjectIdLinksAvailablepeopleJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsProjectIdLinksAvailablepeopleJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20081**](InlineResponse20081.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsProjectIdMessagesAvailablepeopleJson

> InlineResponse20081 GETProjectsProjectIdMessagesAvailablepeopleJson(ctx, projectId).Execute()

Get available People for a Message

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
    projectId := "projectId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeopleApi.GETProjectsProjectIdMessagesAvailablepeopleJson(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.GETProjectsProjectIdMessagesAvailablepeopleJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsProjectIdMessagesAvailablepeopleJson`: InlineResponse20081
    fmt.Fprintf(os.Stdout, "Response from `PeopleApi.GETProjectsProjectIdMessagesAvailablepeopleJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsProjectIdMessagesAvailablepeopleJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20081**](InlineResponse20081.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsProjectIdMilestonesAvailablepeopleJson

> InlineResponse20081 GETProjectsProjectIdMilestonesAvailablepeopleJson(ctx, projectId).Execute()

Get available People for a Milestone

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
    projectId := "projectId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeopleApi.GETProjectsProjectIdMilestonesAvailablepeopleJson(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.GETProjectsProjectIdMilestonesAvailablepeopleJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsProjectIdMilestonesAvailablepeopleJson`: InlineResponse20081
    fmt.Fprintf(os.Stdout, "Response from `PeopleApi.GETProjectsProjectIdMilestonesAvailablepeopleJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsProjectIdMilestonesAvailablepeopleJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20081**](InlineResponse20081.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsProjectIdNotebooksAvailablepeopleJson

> InlineResponse20081 GETProjectsProjectIdNotebooksAvailablepeopleJson(ctx, projectId).Execute()

Get available People for following a Notebook

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
    projectId := "projectId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeopleApi.GETProjectsProjectIdNotebooksAvailablepeopleJson(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.GETProjectsProjectIdNotebooksAvailablepeopleJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsProjectIdNotebooksAvailablepeopleJson`: InlineResponse20081
    fmt.Fprintf(os.Stdout, "Response from `PeopleApi.GETProjectsProjectIdNotebooksAvailablepeopleJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsProjectIdNotebooksAvailablepeopleJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20081**](InlineResponse20081.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsProjectIdTasksAvailablepeopleJson

> InlineResponse20081 GETProjectsProjectIdTasksAvailablepeopleJson(ctx, projectId).Execute()

Get available People for a Task

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
    projectId := "projectId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeopleApi.GETProjectsProjectIdTasksAvailablepeopleJson(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.GETProjectsProjectIdTasksAvailablepeopleJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsProjectIdTasksAvailablepeopleJson`: InlineResponse20081
    fmt.Fprintf(os.Stdout, "Response from `PeopleApi.GETProjectsProjectIdTasksAvailablepeopleJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsProjectIdTasksAvailablepeopleJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20081**](InlineResponse20081.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsProjectidPeopleJson

> InlineResponse20084 GETProjectsProjectidPeopleJson(ctx, projectid).Execute()

Get all People (within a Project)



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
    projectid := "projectid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeopleApi.GETProjectsProjectidPeopleJson(context.Background(), projectid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.GETProjectsProjectidPeopleJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsProjectidPeopleJson`: InlineResponse20084
    fmt.Fprintf(os.Stdout, "Response from `PeopleApi.GETProjectsProjectidPeopleJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsProjectidPeopleJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20084**](InlineResponse20084.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETStatsJson

> InlineResponse20090 GETStatsJson(ctx).GetPermissions(getPermissions).OnlyMyProjects(onlyMyProjects).Execute()

Current User Summary Stats



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
    getPermissions := true // bool | Return Permissions information in the response (optional)
    onlyMyProjects := true // bool | Summary of the users Projects only or all Projects (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeopleApi.GETStatsJson(context.Background()).GetPermissions(getPermissions).OnlyMyProjects(onlyMyProjects).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.GETStatsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETStatsJson`: InlineResponse20090
    fmt.Fprintf(os.Stdout, "Response from `PeopleApi.GETStatsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETStatsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPermissions** | **bool** | Return Permissions information in the response | 
 **onlyMyProjects** | **bool** | Summary of the users Projects only or all Projects | 

### Return type

[**InlineResponse20090**](InlineResponse20090.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTPeopleJson

> InlineResponse20018 POSTPeopleJson(ctx).Body(body).Execute()

Creates a new User Account



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
    body := *openapiclient.NewInlineObject37() // InlineObject37 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeopleApi.POSTPeopleJson(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.POSTPeopleJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTPeopleJson`: InlineResponse20018
    fmt.Fprintf(os.Stdout, "Response from `PeopleApi.POSTPeopleJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTPeopleJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InlineObject37**](InlineObject37.md) |  | 

### Return type

[**InlineResponse20018**](InlineResponse20018.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsIdPeopleJson

> InlineResponse2001 POSTProjectsIdPeopleJson(ctx, id).Body(body).Execute()

Add/Remove People to existing Project



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
    id := "id_example" // string | 
    body := *openapiclient.NewInlineObject68() // InlineObject68 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeopleApi.POSTProjectsIdPeopleJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.POSTProjectsIdPeopleJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsIdPeopleJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `PeopleApi.POSTProjectsIdPeopleJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsIdPeopleJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject68**](InlineObject68.md) |  | 

### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTPeopleIdJson

> InlineResponse20018 PUTPeopleIdJson(ctx, id).Body(body).Execute()

Editing a User



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
    id := "id_example" // string | Id of the person.
    body := *openapiclient.NewInlineObject39() // InlineObject39 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeopleApi.PUTPeopleIdJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.PUTPeopleIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTPeopleIdJson`: InlineResponse20018
    fmt.Fprintf(os.Stdout, "Response from `PeopleApi.PUTPeopleIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the person. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTPeopleIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject39**](InlineObject39.md) |  | 

### Return type

[**InlineResponse20018**](InlineResponse20018.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

