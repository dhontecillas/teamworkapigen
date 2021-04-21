# \ExpensesApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEExpensesIdJson**](ExpensesApi.md#DELETEExpensesIdJson) | **Delete** /expenses/{id}.json | Delete a single Expense
[**GETExpensesIdJson**](ExpensesApi.md#GETExpensesIdJson) | **Get** /expenses/{id}.json | Get a single Expense
[**GETExpensesJson**](ExpensesApi.md#GETExpensesJson) | **Get** /expenses.json | Get all Expenses across your Projects
[**GETProjectsIdExpensesJson**](ExpensesApi.md#GETProjectsIdExpensesJson) | **Get** /projects/{id}/expenses.json | Get all Expenses on a single Project
[**POSTExpensesJson**](ExpensesApi.md#POSTExpensesJson) | **Post** /expenses.json | Create a new Expense in a Project
[**PUTExpensesIdJson**](ExpensesApi.md#PUTExpensesIdJson) | **Put** /expenses/{id}.json | Update a single Expense



## DELETEExpensesIdJson

> InlineResponse2001 DELETEExpensesIdJson(ctx, id).Execute()

Delete a single Expense



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
    resp, r, err := api_client.ExpensesApi.DELETEExpensesIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpensesApi.DELETEExpensesIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DELETEExpensesIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `ExpensesApi.DELETEExpensesIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEExpensesIdJsonRequest struct via the builder pattern


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


## GETExpensesIdJson

> InlineResponse20019 GETExpensesIdJson(ctx, id).Execute()

Get a single Expense



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
    resp, r, err := api_client.ExpensesApi.GETExpensesIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpensesApi.GETExpensesIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETExpensesIdJson`: InlineResponse20019
    fmt.Fprintf(os.Stdout, "Response from `ExpensesApi.GETExpensesIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETExpensesIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20019**](InlineResponse20019.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETExpensesJson

> InlineResponse20017 GETExpensesJson(ctx).Execute()

Get all Expenses across your Projects



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
    resp, r, err := api_client.ExpensesApi.GETExpensesJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpensesApi.GETExpensesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETExpensesJson`: InlineResponse20017
    fmt.Fprintf(os.Stdout, "Response from `ExpensesApi.GETExpensesJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETExpensesJsonRequest struct via the builder pattern


### Return type

[**InlineResponse20017**](InlineResponse20017.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsIdExpensesJson

> InlineResponse20017 GETProjectsIdExpensesJson(ctx, id).Execute()

Get all Expenses on a single Project



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
    resp, r, err := api_client.ExpensesApi.GETProjectsIdExpensesJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpensesApi.GETProjectsIdExpensesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsIdExpensesJson`: InlineResponse20017
    fmt.Fprintf(os.Stdout, "Response from `ExpensesApi.GETProjectsIdExpensesJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsIdExpensesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20017**](InlineResponse20017.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTExpensesJson

> InlineResponse20018 POSTExpensesJson(ctx).Body(body).Execute()

Create a new Expense in a Project



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
    body := *openapiclient.NewInlineObject14() // InlineObject14 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExpensesApi.POSTExpensesJson(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpensesApi.POSTExpensesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTExpensesJson`: InlineResponse20018
    fmt.Fprintf(os.Stdout, "Response from `ExpensesApi.POSTExpensesJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTExpensesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InlineObject14**](InlineObject14.md) |  | 

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


## PUTExpensesIdJson

> InlineResponse2001 PUTExpensesIdJson(ctx, id).Body(body).Execute()

Update a single Expense



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
    body := *openapiclient.NewInlineObject15() // InlineObject15 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExpensesApi.PUTExpensesIdJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExpensesApi.PUTExpensesIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTExpensesIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `ExpensesApi.PUTExpensesIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTExpensesIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject15**](InlineObject15.md) |  | 

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

