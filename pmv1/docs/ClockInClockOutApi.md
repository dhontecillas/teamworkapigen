# \ClockInClockOutApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETPeoplePersonIdClockinsJson**](ClockInClockOutApi.md#GETPeoplePersonIdClockinsJson) | **Get** /people/{personId}/clockins.json | Get all clock Ins
[**POSTClockinJson**](ClockInClockOutApi.md#POSTClockinJson) | **Post** /clockin.json | Quick log for another user
[**POSTMeClockinJson**](ClockInClockOutApi.md#POSTMeClockinJson) | **Post** /me/clockin.json | Clock me in
[**POSTMeClockoutJson**](ClockInClockOutApi.md#POSTMeClockoutJson) | **Post** /me/clockout.json | Clock me out
[**PUTClockinClockinIdJson**](ClockInClockOutApi.md#PUTClockinClockinIdJson) | **Put** /clockin/{clockinId}.json | Edit a clock in/out entry



## GETPeoplePersonIdClockinsJson

> InlineResponse20041 GETPeoplePersonIdClockinsJson(ctx, personId).Page(page).PageSize(pageSize).Execute()

Get all clock Ins



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
    personId := "personId_example" // string | 
    page := float32(8.14) // float32 |  (optional)
    pageSize := float32(8.14) // float32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClockInClockOutApi.GETPeoplePersonIdClockinsJson(context.Background(), personId).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClockInClockOutApi.GETPeoplePersonIdClockinsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPeoplePersonIdClockinsJson`: InlineResponse20041
    fmt.Fprintf(os.Stdout, "Response from `ClockInClockOutApi.GETPeoplePersonIdClockinsJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPeoplePersonIdClockinsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** |  | 
 **pageSize** | **float32** |  | 

### Return type

[**InlineResponse20041**](InlineResponse20041.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTClockinJson

> InlineResponse2001 POSTClockinJson(ctx).Body(body).Execute()

Quick log for another user



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
    body := *openapiclient.NewInlineObject7(*openapiclient.NewClockinJsonClockIn(int32(123))) // InlineObject7 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClockInClockOutApi.POSTClockinJson(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClockInClockOutApi.POSTClockinJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTClockinJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `ClockInClockOutApi.POSTClockinJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTClockinJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InlineObject7**](InlineObject7.md) |  | 

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


## POSTMeClockinJson

> InlineResponse2001 POSTMeClockinJson(ctx).Execute()

Clock me in



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
    resp, r, err := api_client.ClockInClockOutApi.POSTMeClockinJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClockInClockOutApi.POSTMeClockinJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTMeClockinJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `ClockInClockOutApi.POSTMeClockinJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTMeClockinJsonRequest struct via the builder pattern


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


## POSTMeClockoutJson

> InlineResponse2001 POSTMeClockoutJson(ctx).Execute()

Clock me out



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
    resp, r, err := api_client.ClockInClockOutApi.POSTMeClockoutJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClockInClockOutApi.POSTMeClockoutJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTMeClockoutJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `ClockInClockOutApi.POSTMeClockoutJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTMeClockoutJsonRequest struct via the builder pattern


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


## PUTClockinClockinIdJson

> InlineResponse2001 PUTClockinClockinIdJson(ctx, clockinId).Body(body).Execute()

Edit a clock in/out entry

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
    clockinId := "clockinId_example" // string | 
    body := *openapiclient.NewInlineObject8() // InlineObject8 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClockInClockOutApi.PUTClockinClockinIdJson(context.Background(), clockinId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClockInClockOutApi.PUTClockinClockinIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTClockinClockinIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `ClockInClockOutApi.PUTClockinClockinIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clockinId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTClockinClockinIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject8**](InlineObject8.md) |  | 

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

