# \ProjectsApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEProjectsIdJson**](ProjectsApi.md#DELETEProjectsIdJson) | **Delete** /projects/{id}.json | Delete Project
[**GETCompaniesIdProjectsJson**](ProjectsApi.md#GETCompaniesIdProjectsJson) | **Get** /companies/{id}/projects.json | Retrieve Projects assigned to a specific Company
[**GETProjectsIdJson**](ProjectsApi.md#GETProjectsIdJson) | **Get** /projects/{id}.json | Retrieve a Single Project
[**GETProjectsIdRatesJson**](ProjectsApi.md#GETProjectsIdRatesJson) | **Get** /projects/{id}/rates.json | Get Project Rates
[**GETProjectsIdStatsJson**](ProjectsApi.md#GETProjectsIdStatsJson) | **Get** /projects/{id}/stats.json | Get Project Stats
[**GETProjectsJson**](ProjectsApi.md#GETProjectsJson) | **Get** /projects.json | Retrieve All Projects
[**GETProjectsStarredJson**](ProjectsApi.md#GETProjectsStarredJson) | **Get** /projects/starred.json | Retrieve your Starred Projects
[**POSTProjectsIdRatesJson**](ProjectsApi.md#POSTProjectsIdRatesJson) | **Post** /projects/{id}/rates.json | Set Project Rates
[**POSTProjectsJson**](ProjectsApi.md#POSTProjectsJson) | **Post** /projects.json | Create Project
[**PUTProjectsIdJson**](ProjectsApi.md#PUTProjectsIdJson) | **Put** /projects/{id}.json | Update Project
[**PUTProjectsIdStarJson**](ProjectsApi.md#PUTProjectsIdStarJson) | **Put** /projects/{id}/star.json | Star a Project
[**PUTProjectsIdUnstarJson**](ProjectsApi.md#PUTProjectsIdUnstarJson) | **Put** /projects/{id}/unstar.json | Un-star a Project



## DELETEProjectsIdJson

> InlineResponse2001 DELETEProjectsIdJson(ctx, id).Execute()

Delete Project



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.DELETEProjectsIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.DELETEProjectsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DELETEProjectsIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.DELETEProjectsIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEProjectsIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETCompaniesIdProjectsJson

> InlineResponse20014 GETCompaniesIdProjectsJson(ctx, id).IncludePeople(includePeople).Execute()

Retrieve Projects assigned to a specific Company



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
    includePeople := true // bool | You can pass includePeople=true as an optional parameter to get an array of ID's of all people associated with each project included in the response under the field people (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.GETCompaniesIdProjectsJson(context.Background(), id).IncludePeople(includePeople).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GETCompaniesIdProjectsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCompaniesIdProjectsJson`: InlineResponse20014
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GETCompaniesIdProjectsJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCompaniesIdProjectsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includePeople** | **bool** | You can pass includePeople&#x3D;true as an optional parameter to get an array of ID&#39;s of all people associated with each project included in the response under the field people | 

### Return type

[**InlineResponse20014**](InlineResponse20014.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsIdJson

> InlineResponse20060 GETProjectsIdJson(ctx, id).IncludePeople(includePeople).ProjectOwnerIds(projectOwnerIds).ProjectHealth(projectHealth).IncludeProjectOwner(includeProjectOwner).UserId(userId).Execute()

Retrieve a Single Project



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
    includePeople := true // bool | You can pass includePeople=true as an optional parameter to get an array of ID's of all people associated with each project included in the response under the field people (optional)
    projectOwnerIds := "projectOwnerIds_example" // string |  (optional)
    projectHealth := "projectHealth_example" // string |  (optional)
    includeProjectOwner := true // bool |  (optional) (default to false)
    userId := "userId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.GETProjectsIdJson(context.Background(), id).IncludePeople(includePeople).ProjectOwnerIds(projectOwnerIds).ProjectHealth(projectHealth).IncludeProjectOwner(includeProjectOwner).UserId(userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GETProjectsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsIdJson`: InlineResponse20060
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GETProjectsIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includePeople** | **bool** | You can pass includePeople&#x3D;true as an optional parameter to get an array of ID&#39;s of all people associated with each project included in the response under the field people | 
 **projectOwnerIds** | **string** |  | 
 **projectHealth** | **string** |  | 
 **includeProjectOwner** | **bool** |  | [default to false]
 **userId** | **string** |  | 

### Return type

[**InlineResponse20060**](InlineResponse20060.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsIdRatesJson

> InlineResponse20072 GETProjectsIdRatesJson(ctx, id).Page(page).PageSize(pageSize).Execute()

Get Project Rates



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
    page := int32(56) // int32 | The page number to show (optional)
    pageSize := int32(56) // int32 | The number of results to show per page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.GETProjectsIdRatesJson(context.Background(), id).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GETProjectsIdRatesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsIdRatesJson`: InlineResponse20072
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GETProjectsIdRatesJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsIdRatesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page number to show | 
 **pageSize** | **int32** | The number of results to show per page | 

### Return type

[**InlineResponse20072**](InlineResponse20072.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsIdStatsJson

> InlineResponse20075 GETProjectsIdStatsJson(ctx, id).ResponsiblePartyIds(responsiblePartyIds).Onlymyprojects(onlymyprojects).ForStarredProjects(forStarredProjects).OnlyMyEvents(onlyMyEvents).GetPermissions(getPermissions).EventsInNext(eventsInNext).Execute()

Get Project Stats



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
    responsiblePartyIds := "responsiblePartyIds_example" // string |  (optional)
    onlymyprojects := true // bool |  (optional)
    forStarredProjects := true // bool |  (optional)
    onlyMyEvents := true // bool |  (optional)
    getPermissions := true // bool |  (optional)
    eventsInNext := float32(8.14) // float32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.GETProjectsIdStatsJson(context.Background(), id).ResponsiblePartyIds(responsiblePartyIds).Onlymyprojects(onlymyprojects).ForStarredProjects(forStarredProjects).OnlyMyEvents(onlyMyEvents).GetPermissions(getPermissions).EventsInNext(eventsInNext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GETProjectsIdStatsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsIdStatsJson`: InlineResponse20075
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GETProjectsIdStatsJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsIdStatsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **responsiblePartyIds** | **string** |  | 
 **onlymyprojects** | **bool** |  | 
 **forStarredProjects** | **bool** |  | 
 **onlyMyEvents** | **bool** |  | 
 **getPermissions** | **bool** |  | 
 **eventsInNext** | **float32** |  | 

### Return type

[**InlineResponse20075**](InlineResponse20075.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsJson

> InlineResponse20053 GETProjectsJson(ctx).Status(status).UpdatedAfterDate(updatedAfterDate).Orderby(orderby).CreatedAfterDate(createdAfterDate).CreatedAfterTime(createdAfterTime).CatId(catId).IncludePeople(includePeople).IncludeProjectOwner(includeProjectOwner).Page(page).PageSize(pageSize).OrderMode(orderMode).OnlyStarredProjects(onlyStarredProjects).CompanyId(companyId).ProjectOwnerIds(projectOwnerIds).SearchTerm(searchTerm).GetDeleted(getDeleted).IncludeTags(includeTags).UserId(userId).UpdatedAfterDateTime(updatedAfterDateTime).Execute()

Retrieve All Projects



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
    status := "status_example" // string | You can pass a status option to retrieve projects filtered by status (ALL, ACTIVE, ARCHIVED, CURRENT, LATE, COMPLETED) (optional) (default to "active")
    updatedAfterDate := "updatedAfterDate_example" // string | Send back only those projects recently updated after a certain date. Eg. 20100603 - datetime (optional)
    orderby := "orderby_example" // string | Order projects by a parameter, eg companyName, name, lastActivityDate (optional) (default to "name")
    createdAfterDate := "createdAfterDate_example" // string |  (optional)
    createdAfterTime := "createdAfterTime_example" // string | Return projects created after a certain time. (optional)
    catId := int32(56) // int32 | Return projects in a certain category, using cat Id (optional)
    includePeople := true // bool | You can pass includePeople=true as an optional parameter to get an array of ID's of all people associated with each project included in the response under the field people (optional)
    includeProjectOwner := true // bool | The project owner can be returned by adding this parameter to the project's endpoint.  (optional) (default to false)
    page := "page_example" // string |  (optional)
    pageSize := float32(8.14) // float32 | Max page size is 500. (optional) (default to 500.0)
    orderMode := "orderMode_example" // string |  (optional) (default to "asc")
    onlyStarredProjects := true // bool |  (optional)
    companyId := "companyId_example" // string |  (optional)
    projectOwnerIds := "projectOwnerIds_example" // string |  (optional)
    searchTerm := "searchTerm_example" // string |  (optional)
    getDeleted := true // bool |  (optional) (default to false)
    includeTags := true // bool |  (optional)
    userId := "userId_example" // string |  (optional)
    updatedAfterDateTime := "updatedAfterDateTime_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.GETProjectsJson(context.Background()).Status(status).UpdatedAfterDate(updatedAfterDate).Orderby(orderby).CreatedAfterDate(createdAfterDate).CreatedAfterTime(createdAfterTime).CatId(catId).IncludePeople(includePeople).IncludeProjectOwner(includeProjectOwner).Page(page).PageSize(pageSize).OrderMode(orderMode).OnlyStarredProjects(onlyStarredProjects).CompanyId(companyId).ProjectOwnerIds(projectOwnerIds).SearchTerm(searchTerm).GetDeleted(getDeleted).IncludeTags(includeTags).UserId(userId).UpdatedAfterDateTime(updatedAfterDateTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GETProjectsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsJson`: InlineResponse20053
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GETProjectsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | You can pass a status option to retrieve projects filtered by status (ALL, ACTIVE, ARCHIVED, CURRENT, LATE, COMPLETED) | [default to &quot;active&quot;]
 **updatedAfterDate** | **string** | Send back only those projects recently updated after a certain date. Eg. 20100603 - datetime | 
 **orderby** | **string** | Order projects by a parameter, eg companyName, name, lastActivityDate | [default to &quot;name&quot;]
 **createdAfterDate** | **string** |  | 
 **createdAfterTime** | **string** | Return projects created after a certain time. | 
 **catId** | **int32** | Return projects in a certain category, using cat Id | 
 **includePeople** | **bool** | You can pass includePeople&#x3D;true as an optional parameter to get an array of ID&#39;s of all people associated with each project included in the response under the field people | 
 **includeProjectOwner** | **bool** | The project owner can be returned by adding this parameter to the project&#39;s endpoint.  | [default to false]
 **page** | **string** |  | 
 **pageSize** | **float32** | Max page size is 500. | [default to 500.0]
 **orderMode** | **string** |  | [default to &quot;asc&quot;]
 **onlyStarredProjects** | **bool** |  | 
 **companyId** | **string** |  | 
 **projectOwnerIds** | **string** |  | 
 **searchTerm** | **string** |  | 
 **getDeleted** | **bool** |  | [default to false]
 **includeTags** | **bool** |  | 
 **userId** | **string** |  | 
 **updatedAfterDateTime** | **string** |  | 

### Return type

[**InlineResponse20053**](InlineResponse20053.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsStarredJson

> InlineResponse20014 GETProjectsStarredJson(ctx).Execute()

Retrieve your Starred Projects



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
    resp, r, err := api_client.ProjectsApi.GETProjectsStarredJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GETProjectsStarredJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsStarredJson`: InlineResponse20014
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GETProjectsStarredJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsStarredJsonRequest struct via the builder pattern


### Return type

[**InlineResponse20014**](InlineResponse20014.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsIdRatesJson

> InlineResponse2001 POSTProjectsIdRatesJson(ctx, id).Body(body).Execute()

Set Project Rates



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
    body := *openapiclient.NewInlineObject70() // InlineObject70 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.POSTProjectsIdRatesJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.POSTProjectsIdRatesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsIdRatesJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.POSTProjectsIdRatesJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsIdRatesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject70**](InlineObject70.md) |  | 

### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsJson

> InlineResponse201 POSTProjectsJson(ctx).Body(body).Execute()

Create Project



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
    body := *openapiclient.NewInlineObject50(*openapiclient.NewProjectsJsonProject("Name_example")) // InlineObject50 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.POSTProjectsJson(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.POSTProjectsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsJson`: InlineResponse201
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.POSTProjectsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InlineObject50**](InlineObject50.md) |  | 

### Return type

[**InlineResponse201**](InlineResponse201.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsIdJson

> InlineResponse2001 PUTProjectsIdJson(ctx, id).Body(body).Execute()

Update Project



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
    body := *openapiclient.NewInlineObject55() // InlineObject55 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.PUTProjectsIdJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.PUTProjectsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.PUTProjectsIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject55**](InlineObject55.md) |  | 

### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsIdStarJson

> InlineResponse2001 PUTProjectsIdStarJson(ctx, id).Execute()

Star a Project



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.PUTProjectsIdStarJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.PUTProjectsIdStarJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsIdStarJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.PUTProjectsIdStarJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsIdStarJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsIdUnstarJson

> map[string]interface{} PUTProjectsIdUnstarJson(ctx, id).Body(body).Execute()

Un-star a Project



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
    body := *openapiclient.NewInlineObject76() // InlineObject76 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.PUTProjectsIdUnstarJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.PUTProjectsIdUnstarJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsIdUnstarJson`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.PUTProjectsIdUnstarJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsIdUnstarJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject76**](InlineObject76.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

