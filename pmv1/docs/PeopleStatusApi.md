# \PeopleStatusApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEMeStatusIdJson**](PeopleStatusApi.md#DELETEMeStatusIdJson) | **Delete** /me/status/{id}.json | Delete my Status
[**DELETEPeoplePersonIdStatusStatusIdJson**](PeopleStatusApi.md#DELETEPeoplePersonIdStatusStatusIdJson) | **Delete** /people/{person_id}/status/{status_id}.json | Delete User Status
[**DELETEPeopleStatusIdJson**](PeopleStatusApi.md#DELETEPeopleStatusIdJson) | **Delete** /people/status/{id}.json | Delete a Status
[**GETMeStatusJson**](PeopleStatusApi.md#GETMeStatusJson) | **Get** /me/status.json | Retrieve my Status
[**GETPeopleIdStatusJson**](PeopleStatusApi.md#GETPeopleIdStatusJson) | **Get** /people/{id}/status.json | Retrieve a Persons Status
[**GETStatusesJson**](PeopleStatusApi.md#GETStatusesJson) | **Get** /statuses.json | Retrieve all Statuses
[**POSTMeStatusJson**](PeopleStatusApi.md#POSTMeStatusJson) | **Post** /me/status.json | Create my Status
[**PUTMeStatusIdJson**](PeopleStatusApi.md#PUTMeStatusIdJson) | **Put** /me/status/id.json | Update my Status
[**PUTPeoplePersonIdStatusStatusIdJson**](PeopleStatusApi.md#PUTPeoplePersonIdStatusStatusIdJson) | **Put** /people/{person_id}/status/{status_id}.json | Update User Status
[**PUTPeopleStatusIdJson**](PeopleStatusApi.md#PUTPeopleStatusIdJson) | **Put** /people/status/{id}.json | Update Status



## DELETEMeStatusIdJson

> InlineResponse2001 DELETEMeStatusIdJson(ctx, id).Execute()

Delete my Status



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
    resp, r, err := api_client.PeopleStatusApi.DELETEMeStatusIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleStatusApi.DELETEMeStatusIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DELETEMeStatusIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `PeopleStatusApi.DELETEMeStatusIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEMeStatusIdJsonRequest struct via the builder pattern


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


## DELETEPeoplePersonIdStatusStatusIdJson

> InlineResponse2001 DELETEPeoplePersonIdStatusStatusIdJson(ctx, personId, statusId).Execute()

Delete User Status



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
    statusId := "statusId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeopleStatusApi.DELETEPeoplePersonIdStatusStatusIdJson(context.Background(), personId, statusId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleStatusApi.DELETEPeoplePersonIdStatusStatusIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DELETEPeoplePersonIdStatusStatusIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `PeopleStatusApi.DELETEPeoplePersonIdStatusStatusIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**statusId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEPeoplePersonIdStatusStatusIdJsonRequest struct via the builder pattern


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


## DELETEPeopleStatusIdJson

> InlineResponse2001 DELETEPeopleStatusIdJson(ctx, id).Execute()

Delete a Status



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
    resp, r, err := api_client.PeopleStatusApi.DELETEPeopleStatusIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleStatusApi.DELETEPeopleStatusIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DELETEPeopleStatusIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `PeopleStatusApi.DELETEPeopleStatusIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEPeopleStatusIdJsonRequest struct via the builder pattern


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


## GETMeStatusJson

> InlineResponse20028 GETMeStatusJson(ctx).Execute()

Retrieve my Status



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
    resp, r, err := api_client.PeopleStatusApi.GETMeStatusJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleStatusApi.GETMeStatusJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETMeStatusJson`: InlineResponse20028
    fmt.Fprintf(os.Stdout, "Response from `PeopleStatusApi.GETMeStatusJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETMeStatusJsonRequest struct via the builder pattern


### Return type

[**InlineResponse20028**](InlineResponse20028.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPeopleIdStatusJson

> InlineResponse20028 GETPeopleIdStatusJson(ctx, id).Execute()

Retrieve a Persons Status



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
    resp, r, err := api_client.PeopleStatusApi.GETPeopleIdStatusJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleStatusApi.GETPeopleIdStatusJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPeopleIdStatusJson`: InlineResponse20028
    fmt.Fprintf(os.Stdout, "Response from `PeopleStatusApi.GETPeopleIdStatusJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPeopleIdStatusJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20028**](InlineResponse20028.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETStatusesJson

> InlineResponse20091 GETStatusesJson(ctx).IncludeClockin(includeClockin).Execute()

Retrieve all Statuses



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
    includeClockin := true // bool | Will return 0 or 1 based on if the user is clocked in or not.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeopleStatusApi.GETStatusesJson(context.Background()).IncludeClockin(includeClockin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleStatusApi.GETStatusesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETStatusesJson`: InlineResponse20091
    fmt.Fprintf(os.Stdout, "Response from `PeopleStatusApi.GETStatusesJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETStatusesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeClockin** | **bool** | Will return 0 or 1 based on if the user is clocked in or not.  | 

### Return type

[**InlineResponse20091**](InlineResponse20091.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTMeStatusJson

> POSTMeStatusJson(ctx).Body(body).Execute()

Create my Status



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
    body := *openapiclient.NewInlineObject25() // InlineObject25 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeopleStatusApi.POSTMeStatusJson(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleStatusApi.POSTMeStatusJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTMeStatusJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InlineObject25**](InlineObject25.md) |  | 

### Return type

 (empty response body)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTMeStatusIdJson

> InlineResponse20029 PUTMeStatusIdJson(ctx).Execute()

Update my Status



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
    resp, r, err := api_client.PeopleStatusApi.PUTMeStatusIdJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleStatusApi.PUTMeStatusIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTMeStatusIdJson`: InlineResponse20029
    fmt.Fprintf(os.Stdout, "Response from `PeopleStatusApi.PUTMeStatusIdJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPUTMeStatusIdJsonRequest struct via the builder pattern


### Return type

[**InlineResponse20029**](InlineResponse20029.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTPeoplePersonIdStatusStatusIdJson

> InlineResponse2001 PUTPeoplePersonIdStatusStatusIdJson(ctx, personId, statusId).Body(body).Execute()

Update User Status



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
    statusId := "statusId_example" // string | 
    body := *openapiclient.NewInlineObject40() // InlineObject40 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeopleStatusApi.PUTPeoplePersonIdStatusStatusIdJson(context.Background(), personId, statusId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleStatusApi.PUTPeoplePersonIdStatusStatusIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTPeoplePersonIdStatusStatusIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `PeopleStatusApi.PUTPeoplePersonIdStatusStatusIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**statusId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTPeoplePersonIdStatusStatusIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**InlineObject40**](InlineObject40.md) |  | 

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


## PUTPeopleStatusIdJson

> InlineResponse2001 PUTPeopleStatusIdJson(ctx, id).Body(body).Execute()

Update Status



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
    body := *openapiclient.NewInlineObject38() // InlineObject38 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeopleStatusApi.PUTPeopleStatusIdJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleStatusApi.PUTPeopleStatusIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTPeopleStatusIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `PeopleStatusApi.PUTPeopleStatusIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTPeopleStatusIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject38**](InlineObject38.md) |  | 

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

