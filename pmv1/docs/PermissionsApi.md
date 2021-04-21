# \PermissionsApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsProjectIdPeoplePersonIdJson**](PermissionsApi.md#GETProjectsProjectIdPeoplePersonIdJson) | **Get** /projects/{project_id}/people/{person_id}.json | Get a Users Permissions on a Project
[**POSTProjectsProjectIdPeoplePesonIdJson**](PermissionsApi.md#POSTProjectsProjectIdPeoplePesonIdJson) | **Post** /projects/{project_id}/people/{peson_id}.json | Add a New User to a Project
[**PUTProjectsIdPeopleIdJson**](PermissionsApi.md#PUTProjectsIdPeopleIdJson) | **Put** /projects/{id}/people/{id}.json | Update a Users Permissions on a Project
[**PUTProjectsIdPeopleJson**](PermissionsApi.md#PUTProjectsIdPeopleJson) | **Put** /projects/{id}/people.json | Add/Remove multiple People to/from a Project



## GETProjectsProjectIdPeoplePersonIdJson

> InlineResponse20083 GETProjectsProjectIdPeoplePersonIdJson(ctx, projectId, personId).Execute()

Get a Users Permissions on a Project



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
    personId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PermissionsApi.GETProjectsProjectIdPeoplePersonIdJson(context.Background(), projectId, personId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.GETProjectsProjectIdPeoplePersonIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsProjectIdPeoplePersonIdJson`: InlineResponse20083
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.GETProjectsProjectIdPeoplePersonIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**personId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsProjectIdPeoplePersonIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20083**](InlineResponse20083.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsProjectIdPeoplePesonIdJson

> InlineResponse2001 POSTProjectsProjectIdPeoplePesonIdJson(ctx, projectId, pesonId).Execute()

Add a New User to a Project



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
    pesonId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PermissionsApi.POSTProjectsProjectIdPeoplePesonIdJson(context.Background(), projectId, pesonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.POSTProjectsProjectIdPeoplePesonIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsProjectIdPeoplePesonIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.POSTProjectsProjectIdPeoplePesonIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**pesonId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsProjectIdPeoplePesonIdJsonRequest struct via the builder pattern


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


## PUTProjectsIdPeopleIdJson

> InlineResponse20071 PUTProjectsIdPeopleIdJson(ctx, id).Execute()

Update a Users Permissions on a Project



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
    resp, r, err := api_client.PermissionsApi.PUTProjectsIdPeopleIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.PUTProjectsIdPeopleIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsIdPeopleIdJson`: InlineResponse20071
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.PUTProjectsIdPeopleIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsIdPeopleIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20071**](InlineResponse20071.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsIdPeopleJson

> InlineResponse20070 PUTProjectsIdPeopleJson(ctx, id).Body(body).Execute()

Add/Remove multiple People to/from a Project



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
    body := *openapiclient.NewInlineObject67() // InlineObject67 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PermissionsApi.PUTProjectsIdPeopleJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.PUTProjectsIdPeopleJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsIdPeopleJson`: InlineResponse20070
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.PUTProjectsIdPeopleJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsIdPeopleJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject67**](InlineObject67.md) |  | 

### Return type

[**InlineResponse20070**](InlineResponse20070.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

