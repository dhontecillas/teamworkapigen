# \PeopleApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3PeopleJson**](PeopleApi.md#GETProjectsApiV3PeopleJson) | **Get** /projects/api/v3/people.json | Returns a list of people
[**GETProjectsApiV3PeopleMetricsPerformanceJson**](PeopleApi.md#GETProjectsApiV3PeopleMetricsPerformanceJson) | **Get** /projects/api/v3/people/metrics/performance.json | Performance of users completing the most tasks
[**GETProjectsApiV3PeopleUtilizationJson**](PeopleApi.md#GETProjectsApiV3PeopleUtilizationJson) | **Get** /projects/api/v3/people/utilization.json | Return the user utilization.
[**GETProjectsApiV3PeopleuserIdAvailabilityJson**](PeopleApi.md#GETProjectsApiV3PeopleuserIdAvailabilityJson) | **Get** /projects/api/v3/people/{userId}/availability.json | Return the user availability.



## GETProjectsApiV3PeopleJson

> PeopleMultiResponse GETProjectsApiV3PeopleJson(ctx).OrderMode(orderMode).OrderBy(orderBy).PageSize(pageSize).Page(page).Include(include).Ids(ids).FieldsCompanies(fieldsCompanies).Execute()

Returns a list of people



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
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "name")
    pageSize := int32(56) // int32 | number of items in a page (not used when generating reports) (optional) (default to 50)
    page := int32(56) // int32 | page number (not used when generating reports) (optional) (default to 1)
    include := []string{"Inner_example"} // []string | include (not used when generating reports) (optional)
    ids := []int32{int32(123)} // []int32 | filter by user ids (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeopleApi.GETProjectsApiV3PeopleJson(context.Background()).OrderMode(orderMode).OrderBy(orderBy).PageSize(pageSize).Page(page).Include(include).Ids(ids).FieldsCompanies(fieldsCompanies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.GETProjectsApiV3PeopleJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3PeopleJson`: PeopleMultiResponse
    fmt.Fprintf(os.Stdout, "Response from `PeopleApi.GETProjectsApiV3PeopleJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3PeopleJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;name&quot;]
 **pageSize** | **int32** | number of items in a page (not used when generating reports) | [default to 50]
 **page** | **int32** | page number (not used when generating reports) | [default to 1]
 **include** | **[]string** | include (not used when generating reports) | 
 **ids** | **[]int32** | filter by user ids | 
 **fieldsCompanies** | **[]string** |  | 

### Return type

[**PeopleMultiResponse**](PeopleMultiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3PeopleMetricsPerformanceJson

> PerformancePeopleMetricPerformancesResponse GETProjectsApiV3PeopleMetricsPerformanceJson(ctx).StartDate(startDate).OrderMode(orderMode).EndDate(endDate).Execute()

Performance of users completing the most tasks



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
    startDate := time.Now() // string |  (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "desc")
    endDate := time.Now() // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeopleApi.GETProjectsApiV3PeopleMetricsPerformanceJson(context.Background()).StartDate(startDate).OrderMode(orderMode).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.GETProjectsApiV3PeopleMetricsPerformanceJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3PeopleMetricsPerformanceJson`: PerformancePeopleMetricPerformancesResponse
    fmt.Fprintf(os.Stdout, "Response from `PeopleApi.GETProjectsApiV3PeopleMetricsPerformanceJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3PeopleMetricsPerformanceJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** |  | 
 **orderMode** | **string** | order mode | [default to &quot;desc&quot;]
 **endDate** | **string** |  | 

### Return type

[**PerformancePeopleMetricPerformancesResponse**](PerformancePeopleMetricPerformancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3PeopleUtilizationJson

> UtilizationResponse GETProjectsApiV3PeopleUtilizationJson(ctx).StartDate(startDate).SortOrder(sortOrder).Sort(sort).SearchTerm(searchTerm).GroupBy(groupBy).EndDate(endDate).PageSize(pageSize).Page(page).IncludeCollaborators(includeCollaborators).IncludeClients(includeClients).UserIds(userIds).TeamIds(teamIds).ProjectIds(projectIds).Include(include).FieldsUtilizations(fieldsUtilizations).FieldsUsers(fieldsUsers).CompanyIds(companyIds).Execute()

Return the user utilization.



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
    sortOrder := "sortOrder_example" // string | order mode (optional) (default to "asc")
    sort := "sort_example" // string | sort by (optional) (default to "name")
    searchTerm := "searchTerm_example" // string | filter by user first or last name (optional)
    groupBy := "groupBy_example" // string | group by (optional)
    endDate := time.Now() // string | filter by end date (optional)
    pageSize := int32(56) // int32 | number of items in a page default: 50 (optional)
    page := int32(56) // int32 | page number default: 1 (optional)
    includeCollaborators := true // bool | include collaborators (optional) (default to true)
    includeClients := true // bool | include client users (optional) (default to true)
    userIds := []int32{int32(123)} // []int32 | filter by userIds (optional)
    teamIds := []int32{int32(123)} // []int32 | filter by team ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUtilizations := []string{"Inner_example"} // []string |  (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    companyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeopleApi.GETProjectsApiV3PeopleUtilizationJson(context.Background()).StartDate(startDate).SortOrder(sortOrder).Sort(sort).SearchTerm(searchTerm).GroupBy(groupBy).EndDate(endDate).PageSize(pageSize).Page(page).IncludeCollaborators(includeCollaborators).IncludeClients(includeClients).UserIds(userIds).TeamIds(teamIds).ProjectIds(projectIds).Include(include).FieldsUtilizations(fieldsUtilizations).FieldsUsers(fieldsUsers).CompanyIds(companyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.GETProjectsApiV3PeopleUtilizationJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3PeopleUtilizationJson`: UtilizationResponse
    fmt.Fprintf(os.Stdout, "Response from `PeopleApi.GETProjectsApiV3PeopleUtilizationJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3PeopleUtilizationJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | filter by start date | 
 **sortOrder** | **string** | order mode | [default to &quot;asc&quot;]
 **sort** | **string** | sort by | [default to &quot;name&quot;]
 **searchTerm** | **string** | filter by user first or last name | 
 **groupBy** | **string** | group by | 
 **endDate** | **string** | filter by end date | 
 **pageSize** | **int32** | number of items in a page default: 50 | 
 **page** | **int32** | page number default: 1 | 
 **includeCollaborators** | **bool** | include collaborators | [default to true]
 **includeClients** | **bool** | include client users | [default to true]
 **userIds** | **[]int32** | filter by userIds | 
 **teamIds** | **[]int32** | filter by team ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **include** | **[]string** | include | 
 **fieldsUtilizations** | **[]string** |  | 
 **fieldsUsers** | **[]string** |  | 
 **companyIds** | **[]int32** | filter by company ids | 

### Return type

[**UtilizationResponse**](UtilizationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3PeopleuserIdAvailabilityJson

> AvailabilityResponse GETProjectsApiV3PeopleuserIdAvailabilityJson(ctx, userId2).StartDate(startDate).EndDate(endDate).UserId(userId).Include(include).FieldsUsers(fieldsUsers).FieldsProjects(fieldsProjects).FieldsAvailability(fieldsAvailability).Execute()

Return the user availability.



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
    userId2 := int32(56) // int32 | 
    startDate := time.Now() // string | filter by start date (optional)
    endDate := time.Now() // string | filter by end date (optional)
    userId := int32(56) // int32 | filter by user (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsAvailability := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeopleApi.GETProjectsApiV3PeopleuserIdAvailabilityJson(context.Background(), userId2).StartDate(startDate).EndDate(endDate).UserId(userId).Include(include).FieldsUsers(fieldsUsers).FieldsProjects(fieldsProjects).FieldsAvailability(fieldsAvailability).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.GETProjectsApiV3PeopleuserIdAvailabilityJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3PeopleuserIdAvailabilityJson`: AvailabilityResponse
    fmt.Fprintf(os.Stdout, "Response from `PeopleApi.GETProjectsApiV3PeopleuserIdAvailabilityJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId2** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3PeopleuserIdAvailabilityJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | filter by start date | 
 **endDate** | **string** | filter by end date | 
 **userId** | **int32** | filter by user | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsAvailability** | **[]string** |  | 

### Return type

[**AvailabilityResponse**](AvailabilityResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

