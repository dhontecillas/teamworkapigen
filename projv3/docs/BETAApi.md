# \BETAApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3ColumnscolumnIdJson**](BETAApi.md#GETProjectsApiV3ColumnscolumnIdJson) | **Get** /projects/api/v3/columns/{columnId}.json | Get a specific column.
[**GETProjectsApiV3CostsUsersuserIdJson**](BETAApi.md#GETProjectsApiV3CostsUsersuserIdJson) | **Get** /projects/api/v3/costs/users/{userId}.json | Get a specific usercost.
[**GETProjectsApiV3RatesInstallationJson**](BETAApi.md#GETProjectsApiV3RatesInstallationJson) | **Get** /projects/api/v3/rates/installation.json | Get an installation default rate.
[**GETProjectsApiV3RatesInstallationUsersJson**](BETAApi.md#GETProjectsApiV3RatesInstallationUsersJson) | **Get** /projects/api/v3/rates/installation/users.json | Get a specific rates.
[**GETProjectsApiV3RatesProjectsprojectIdJson**](BETAApi.md#GETProjectsApiV3RatesProjectsprojectIdJson) | **Get** /projects/api/v3/rates/projects/{projectId}.json | Get a projects default rate.
[**GETProjectsApiV3RatesUseruserIdJson**](BETAApi.md#GETProjectsApiV3RatesUseruserIdJson) | **Get** /projects/api/v3/rates/user/{userId}.json | Get a user default rate.
[**PUTProjectsApiV3CostsUsersuserIdJson**](BETAApi.md#PUTProjectsApiV3CostsUsersuserIdJson) | **Put** /projects/api/v3/costs/users/{userId}.json | set a usercost.
[**PUTProjectsApiV3RatesInstallationJson**](BETAApi.md#PUTProjectsApiV3RatesInstallationJson) | **Put** /projects/api/v3/rates/installation.json | set an installation default rate.
[**PUTProjectsApiV3RatesProjectsprojectIdJson**](BETAApi.md#PUTProjectsApiV3RatesProjectsprojectIdJson) | **Put** /projects/api/v3/rates/projects/{projectId}.json | set a project default rate.
[**PUTProjectsApiV3RatesUsersuserIdJson**](BETAApi.md#PUTProjectsApiV3RatesUsersuserIdJson) | **Put** /projects/api/v3/rates/users/{userId}.json | set a user default rate.



## GETProjectsApiV3ColumnscolumnIdJson

> ColumnResponse GETProjectsApiV3ColumnscolumnIdJson(ctx, columnId).Execute()

Get a specific column.



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
    columnId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BETAApi.GETProjectsApiV3ColumnscolumnIdJson(context.Background(), columnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BETAApi.GETProjectsApiV3ColumnscolumnIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ColumnscolumnIdJson`: ColumnResponse
    fmt.Fprintf(os.Stdout, "Response from `BETAApi.GETProjectsApiV3ColumnscolumnIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**columnId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ColumnscolumnIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ColumnResponse**](ColumnResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3CostsUsersuserIdJson

> UsercostResponse GETProjectsApiV3CostsUsersuserIdJson(ctx, userId).Execute()

Get a specific usercost.



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
    userId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BETAApi.GETProjectsApiV3CostsUsersuserIdJson(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BETAApi.GETProjectsApiV3CostsUsersuserIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3CostsUsersuserIdJson`: UsercostResponse
    fmt.Fprintf(os.Stdout, "Response from `BETAApi.GETProjectsApiV3CostsUsersuserIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3CostsUsersuserIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UsercostResponse**](UsercostResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3RatesInstallationJson

> RatesRateResponse GETProjectsApiV3RatesInstallationJson(ctx).Execute()

Get an installation default rate.



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
    resp, r, err := api_client.BETAApi.GETProjectsApiV3RatesInstallationJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BETAApi.GETProjectsApiV3RatesInstallationJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3RatesInstallationJson`: RatesRateResponse
    fmt.Fprintf(os.Stdout, "Response from `BETAApi.GETProjectsApiV3RatesInstallationJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3RatesInstallationJsonRequest struct via the builder pattern


### Return type

[**RatesRateResponse**](RatesRateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3RatesInstallationUsersJson

> RatesInstallationUserRateResponse GETProjectsApiV3RatesInstallationUsersJson(ctx).PageSize(pageSize).Page(page).Execute()

Get a specific rates.



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
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BETAApi.GETProjectsApiV3RatesInstallationUsersJson(context.Background()).PageSize(pageSize).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BETAApi.GETProjectsApiV3RatesInstallationUsersJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3RatesInstallationUsersJson`: RatesInstallationUserRateResponse
    fmt.Fprintf(os.Stdout, "Response from `BETAApi.GETProjectsApiV3RatesInstallationUsersJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3RatesInstallationUsersJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]

### Return type

[**RatesInstallationUserRateResponse**](RatesInstallationUserRateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3RatesProjectsprojectIdJson

> RatesRateResponse GETProjectsApiV3RatesProjectsprojectIdJson(ctx, projectId).Execute()

Get a projects default rate.



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
    projectId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BETAApi.GETProjectsApiV3RatesProjectsprojectIdJson(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BETAApi.GETProjectsApiV3RatesProjectsprojectIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3RatesProjectsprojectIdJson`: RatesRateResponse
    fmt.Fprintf(os.Stdout, "Response from `BETAApi.GETProjectsApiV3RatesProjectsprojectIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3RatesProjectsprojectIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RatesRateResponse**](RatesRateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3RatesUseruserIdJson

> RatesRateResponse GETProjectsApiV3RatesUseruserIdJson(ctx, userId).Execute()

Get a user default rate.



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
    userId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BETAApi.GETProjectsApiV3RatesUseruserIdJson(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BETAApi.GETProjectsApiV3RatesUseruserIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3RatesUseruserIdJson`: RatesRateResponse
    fmt.Fprintf(os.Stdout, "Response from `BETAApi.GETProjectsApiV3RatesUseruserIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3RatesUseruserIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RatesRateResponse**](RatesRateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV3CostsUsersuserIdJson

> UsercostResponse PUTProjectsApiV3CostsUsersuserIdJson(ctx, userId).UsercostRequest(usercostRequest).Execute()

set a usercost.



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
    userId := int32(56) // int32 | 
    usercostRequest := *openapiclient.NewUsercostRequest() // UsercostRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BETAApi.PUTProjectsApiV3CostsUsersuserIdJson(context.Background(), userId).UsercostRequest(usercostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BETAApi.PUTProjectsApiV3CostsUsersuserIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsApiV3CostsUsersuserIdJson`: UsercostResponse
    fmt.Fprintf(os.Stdout, "Response from `BETAApi.PUTProjectsApiV3CostsUsersuserIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3CostsUsersuserIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **usercostRequest** | [**UsercostRequest**](UsercostRequest.md) |  | 

### Return type

[**UsercostResponse**](UsercostResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV3RatesInstallationJson

> RatesRateResponse PUTProjectsApiV3RatesInstallationJson(ctx).RatesRequest(ratesRequest).Execute()

set an installation default rate.



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
    ratesRequest := *openapiclient.NewRatesRequest() // RatesRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BETAApi.PUTProjectsApiV3RatesInstallationJson(context.Background()).RatesRequest(ratesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BETAApi.PUTProjectsApiV3RatesInstallationJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsApiV3RatesInstallationJson`: RatesRateResponse
    fmt.Fprintf(os.Stdout, "Response from `BETAApi.PUTProjectsApiV3RatesInstallationJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3RatesInstallationJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ratesRequest** | [**RatesRequest**](RatesRequest.md) |  | 

### Return type

[**RatesRateResponse**](RatesRateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV3RatesProjectsprojectIdJson

> RatesRateResponse PUTProjectsApiV3RatesProjectsprojectIdJson(ctx, projectId).RatesRequest(ratesRequest).Execute()

set a project default rate.



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
    projectId := int32(56) // int32 | 
    ratesRequest := *openapiclient.NewRatesRequest() // RatesRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BETAApi.PUTProjectsApiV3RatesProjectsprojectIdJson(context.Background(), projectId).RatesRequest(ratesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BETAApi.PUTProjectsApiV3RatesProjectsprojectIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsApiV3RatesProjectsprojectIdJson`: RatesRateResponse
    fmt.Fprintf(os.Stdout, "Response from `BETAApi.PUTProjectsApiV3RatesProjectsprojectIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3RatesProjectsprojectIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ratesRequest** | [**RatesRequest**](RatesRequest.md) |  | 

### Return type

[**RatesRateResponse**](RatesRateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV3RatesUsersuserIdJson

> RatesRateResponse PUTProjectsApiV3RatesUsersuserIdJson(ctx, userId).RatesRequest(ratesRequest).Execute()

set a user default rate.



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
    userId := int32(56) // int32 | 
    ratesRequest := *openapiclient.NewRatesRequest() // RatesRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BETAApi.PUTProjectsApiV3RatesUsersuserIdJson(context.Background(), userId).RatesRequest(ratesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BETAApi.PUTProjectsApiV3RatesUsersuserIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsApiV3RatesUsersuserIdJson`: RatesRateResponse
    fmt.Fprintf(os.Stdout, "Response from `BETAApi.PUTProjectsApiV3RatesUsersuserIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3RatesUsersuserIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ratesRequest** | [**RatesRequest**](RatesRequest.md) |  | 

### Return type

[**RatesRateResponse**](RatesRateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

