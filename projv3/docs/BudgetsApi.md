# \BudgetsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEProjectsApiV3BudgetNotificationsnotificationIdJson**](BudgetsApi.md#DELETEProjectsApiV3BudgetNotificationsnotificationIdJson) | **Delete** /projects/api/v3/budget/notifications/{notificationId}.json | Delete an existing notification.
[**DELETEProjectsApiV3BudgetsidJson**](BudgetsApi.md#DELETEProjectsApiV3BudgetsidJson) | **Delete** /projects/api/v3/budgets/{id}.json | Delete an existing budget.
[**GETProjectsApiV3BudgetsJson**](BudgetsApi.md#GETProjectsApiV3BudgetsJson) | **Get** /projects/api/v3/budgets.json | Get all budgets.
[**GETProjectsApiV3BudgetsidJson**](BudgetsApi.md#GETProjectsApiV3BudgetsidJson) | **Get** /projects/api/v3/budgets/{id}.json | Get a specific budget.
[**GETProjectsApiV3BudgetsidUsagesJson**](BudgetsApi.md#GETProjectsApiV3BudgetsidUsagesJson) | **Get** /projects/api/v3/budgets/{id}/usages.json | Get daily budget capacity used.
[**GETProjectsApiV3ProjectsBudgetsUsagesJson**](BudgetsApi.md#GETProjectsApiV3ProjectsBudgetsUsagesJson) | **Get** /projects/api/v3/projects/budgets/usages.json | Get the budgets usages for a set of projects
[**PATCHProjectsApiV3BudgetNotificationsnotificationIdJson**](BudgetsApi.md#PATCHProjectsApiV3BudgetNotificationsnotificationIdJson) | **Patch** /projects/api/v3/budget/notifications/{notificationId}.json | Update an existing notification.
[**PATCHProjectsApiV3BudgetsidJson**](BudgetsApi.md#PATCHProjectsApiV3BudgetsidJson) | **Patch** /projects/api/v3/budgets/{id}.json | Update an existing budget.
[**POSTProjectsApiV3BudgetsBulkDeleteJson**](BudgetsApi.md#POSTProjectsApiV3BudgetsBulkDeleteJson) | **Post** /projects/api/v3/budgets/bulk/delete.json | Delete many budgets at once.
[**POSTProjectsApiV3BudgetsJson**](BudgetsApi.md#POSTProjectsApiV3BudgetsJson) | **Post** /projects/api/v3/budgets.json | Create a new budget.



## DELETEProjectsApiV3BudgetNotificationsnotificationIdJson

> DELETEProjectsApiV3BudgetNotificationsnotificationIdJson(ctx, notificationId).Execute()

Delete an existing notification.

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
    notificationId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BudgetsApi.DELETEProjectsApiV3BudgetNotificationsnotificationIdJson(context.Background(), notificationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsApi.DELETEProjectsApiV3BudgetNotificationsnotificationIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEProjectsApiV3BudgetNotificationsnotificationIdJsonRequest struct via the builder pattern


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


## DELETEProjectsApiV3BudgetsidJson

> DELETEProjectsApiV3BudgetsidJson(ctx, id).Execute()

Delete an existing budget.



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
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BudgetsApi.DELETEProjectsApiV3BudgetsidJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsApi.DELETEProjectsApiV3BudgetsidJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEProjectsApiV3BudgetsidJsonRequest struct via the builder pattern


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


## GETProjectsApiV3BudgetsJson

> BudgetProjectBudgetsResponse GETProjectsApiV3BudgetsJson(ctx).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).Type_(type_).Status(status).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).ShowArchived(showArchived).OnlyStarredProjects(onlyStarredProjects).OnlyArchived(onlyArchived).MatchAllProjectTags(matchAllProjectTags).IncludeArchivedProjects(includeArchivedProjects).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).Ids(ids).FieldsUsers(fieldsUsers).FieldsTeams(fieldsTeams).FieldsProjects(fieldsProjects).FieldsProjectBudgetNotifications(fieldsProjectBudgetNotifications).FieldsCompanies(fieldsCompanies).Execute()

Get all budgets.



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
    type_ := "type__example" // string | filter by budget type (optional)
    status := "status_example" // string | filter by budget status (optional)
    projectStatuses := "projectStatuses_example" // string | project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "dateCreated")
    projectHealths := int32(56) // int32 | project healths [] 0: not set 1: bad 2: ok 3: good (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    showDeleted := true // bool | include deleted items (optional) (default to false)
    showArchived := true // bool | include archived items. usually in conjunction with one provided projectId for a project's budget history. (optional) (default to false)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    onlyArchived := true // bool | only archived items (optional)
    matchAllProjectTags := true // bool | match all project tags (optional)
    includeArchivedProjects := true // bool | include archived projects (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    ids := []int32{int32(123)} // []int32 | filter by budget ids (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTeams := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsProjectBudgetNotifications := []string{"Inner_example"} // []string |  (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BudgetsApi.GETProjectsApiV3BudgetsJson(context.Background()).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).Type_(type_).Status(status).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).ShowArchived(showArchived).OnlyStarredProjects(onlyStarredProjects).OnlyArchived(onlyArchived).MatchAllProjectTags(matchAllProjectTags).IncludeArchivedProjects(includeArchivedProjects).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).Ids(ids).FieldsUsers(fieldsUsers).FieldsTeams(fieldsTeams).FieldsProjects(fieldsProjects).FieldsProjectBudgetNotifications(fieldsProjectBudgetNotifications).FieldsCompanies(fieldsCompanies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsApi.GETProjectsApiV3BudgetsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3BudgetsJson`: BudgetProjectBudgetsResponse
    fmt.Fprintf(os.Stdout, "Response from `BudgetsApi.GETProjectsApiV3BudgetsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3BudgetsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfterDate** | **time.Time** | filter by updated after date (deprecated, use updatedAfter) | 
 **updatedAfter** | **time.Time** | filter by updated after date | 
 **type_** | **string** | filter by budget type | 
 **status** | **string** | filter by budget status | 
 **projectStatuses** | **string** | project status | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;dateCreated&quot;]
 **projectHealths** | **int32** | project healths [] 0: not set 1: bad 2: ok 3: good | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **showDeleted** | **bool** | include deleted items | [default to false]
 **showArchived** | **bool** | include archived items. usually in conjunction with one provided projectId for a project&#39;s budget history. | [default to false]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **onlyArchived** | **bool** | only archived items | 
 **matchAllProjectTags** | **bool** | match all project tags | 
 **includeArchivedProjects** | **bool** | include archived projects | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include | 
 **ids** | **[]int32** | filter by budget ids | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTeams** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsProjectBudgetNotifications** | **[]string** |  | 
 **fieldsCompanies** | **[]string** |  | 

### Return type

[**BudgetProjectBudgetsResponse**](BudgetProjectBudgetsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3BudgetsidJson

> BudgetResponse GETProjectsApiV3BudgetsidJson(ctx, id).Execute()

Get a specific budget.



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
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BudgetsApi.GETProjectsApiV3BudgetsidJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsApi.GETProjectsApiV3BudgetsidJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3BudgetsidJson`: BudgetResponse
    fmt.Fprintf(os.Stdout, "Response from `BudgetsApi.GETProjectsApiV3BudgetsidJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3BudgetsidJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BudgetResponse**](BudgetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3BudgetsidUsagesJson

> BudgetProjectBudgetUsageResponse GETProjectsApiV3BudgetsidUsagesJson(ctx, id).Execute()

Get daily budget capacity used.



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
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BudgetsApi.GETProjectsApiV3BudgetsidUsagesJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsApi.GETProjectsApiV3BudgetsidUsagesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3BudgetsidUsagesJson`: BudgetProjectBudgetUsageResponse
    fmt.Fprintf(os.Stdout, "Response from `BudgetsApi.GETProjectsApiV3BudgetsidUsagesJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3BudgetsidUsagesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BudgetProjectBudgetUsageResponse**](BudgetProjectBudgetUsageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3ProjectsBudgetsUsagesJson

> BudgetProjectBudgetsUsagesResponse GETProjectsApiV3ProjectsBudgetsUsagesJson(ctx).Execute()

Get the budgets usages for a set of projects



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
    resp, r, err := api_client.BudgetsApi.GETProjectsApiV3ProjectsBudgetsUsagesJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsApi.GETProjectsApiV3ProjectsBudgetsUsagesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ProjectsBudgetsUsagesJson`: BudgetProjectBudgetsUsagesResponse
    fmt.Fprintf(os.Stdout, "Response from `BudgetsApi.GETProjectsApiV3ProjectsBudgetsUsagesJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsBudgetsUsagesJsonRequest struct via the builder pattern


### Return type

[**BudgetProjectBudgetsUsagesResponse**](BudgetProjectBudgetsUsagesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHProjectsApiV3BudgetNotificationsnotificationIdJson

> NotificationResponse PATCHProjectsApiV3BudgetNotificationsnotificationIdJson(ctx, notificationId).NotificationRequest(notificationRequest).Execute()

Update an existing notification.

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
    notificationId := int32(56) // int32 | 
    notificationRequest := *openapiclient.NewNotificationRequest() // NotificationRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BudgetsApi.PATCHProjectsApiV3BudgetNotificationsnotificationIdJson(context.Background(), notificationId).NotificationRequest(notificationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsApi.PATCHProjectsApiV3BudgetNotificationsnotificationIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHProjectsApiV3BudgetNotificationsnotificationIdJson`: NotificationResponse
    fmt.Fprintf(os.Stdout, "Response from `BudgetsApi.PATCHProjectsApiV3BudgetNotificationsnotificationIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHProjectsApiV3BudgetNotificationsnotificationIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationRequest** | [**NotificationRequest**](NotificationRequest.md) |  | 

### Return type

[**NotificationResponse**](NotificationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHProjectsApiV3BudgetsidJson

> BudgetResponse PATCHProjectsApiV3BudgetsidJson(ctx, id).BudgetRequest(budgetRequest).Execute()

Update an existing budget.

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
    id := int32(56) // int32 | 
    budgetRequest := *openapiclient.NewBudgetRequest() // BudgetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BudgetsApi.PATCHProjectsApiV3BudgetsidJson(context.Background(), id).BudgetRequest(budgetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsApi.PATCHProjectsApiV3BudgetsidJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHProjectsApiV3BudgetsidJson`: BudgetResponse
    fmt.Fprintf(os.Stdout, "Response from `BudgetsApi.PATCHProjectsApiV3BudgetsidJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHProjectsApiV3BudgetsidJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **budgetRequest** | [**BudgetRequest**](BudgetRequest.md) |  | 

### Return type

[**BudgetResponse**](BudgetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsApiV3BudgetsBulkDeleteJson

> POSTProjectsApiV3BudgetsBulkDeleteJson(ctx).BudgetBulkDeleteRequest(budgetBulkDeleteRequest).Execute()

Delete many budgets at once.



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
    budgetBulkDeleteRequest := *openapiclient.NewBudgetBulkDeleteRequest() // BudgetBulkDeleteRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BudgetsApi.POSTProjectsApiV3BudgetsBulkDeleteJson(context.Background()).BudgetBulkDeleteRequest(budgetBulkDeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsApi.POSTProjectsApiV3BudgetsBulkDeleteJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3BudgetsBulkDeleteJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **budgetBulkDeleteRequest** | [**BudgetBulkDeleteRequest**](BudgetBulkDeleteRequest.md) |  | 

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


## POSTProjectsApiV3BudgetsJson

> BudgetResponse POSTProjectsApiV3BudgetsJson(ctx).BudgetRequest(budgetRequest).Execute()

Create a new budget.



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
    budgetRequest := *openapiclient.NewBudgetRequest() // BudgetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BudgetsApi.POSTProjectsApiV3BudgetsJson(context.Background()).BudgetRequest(budgetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsApi.POSTProjectsApiV3BudgetsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsApiV3BudgetsJson`: BudgetResponse
    fmt.Fprintf(os.Stdout, "Response from `BudgetsApi.POSTProjectsApiV3BudgetsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3BudgetsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **budgetRequest** | [**BudgetRequest**](BudgetRequest.md) |  | 

### Return type

[**BudgetResponse**](BudgetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

