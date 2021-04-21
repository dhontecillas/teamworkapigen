# \TrashcanApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETTrashcanProjectsIdJson**](TrashcanApi.md#GETTrashcanProjectsIdJson) | **Get** /trashcan/projects/{id}.json | Get a Project Resources in the Trashcan
[**PUTTrashcanResourceIdRestoreJson**](TrashcanApi.md#PUTTrashcanResourceIdRestoreJson) | **Put** /trashcan/{resource}/{id}/restore.json | Restore Resource Items from the Trashcan



## GETTrashcanProjectsIdJson

> InlineResponse200114 GETTrashcanProjectsIdJson(ctx, id).Execute()

Get a Project Resources in the Trashcan



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
    resp, r, err := api_client.TrashcanApi.GETTrashcanProjectsIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrashcanApi.GETTrashcanProjectsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETTrashcanProjectsIdJson`: InlineResponse200114
    fmt.Fprintf(os.Stdout, "Response from `TrashcanApi.GETTrashcanProjectsIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETTrashcanProjectsIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200114**](InlineResponse200114.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTTrashcanResourceIdRestoreJson

> InlineResponse2001 PUTTrashcanResourceIdRestoreJson(ctx, resource, id).Execute()

Restore Resource Items from the Trashcan



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
    resource := "resource_example" // string | 
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TrashcanApi.PUTTrashcanResourceIdRestoreJson(context.Background(), resource, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrashcanApi.PUTTrashcanResourceIdRestoreJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTTrashcanResourceIdRestoreJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `TrashcanApi.PUTTrashcanResourceIdRestoreJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resource** | **string** |  | 
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTTrashcanResourceIdRestoreJsonRequest struct via the builder pattern


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

