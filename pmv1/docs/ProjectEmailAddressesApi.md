# \ProjectEmailAddressesApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsIdEmailaddressJson**](ProjectEmailAddressesApi.md#GETProjectsIdEmailaddressJson) | **Get** /projects/{id}/emailaddress.json | Get Project Email Address
[**PUTProjectsIdEmailaddressJson**](ProjectEmailAddressesApi.md#PUTProjectsIdEmailaddressJson) | **Put** /projects/{id}/emailaddress.json | Update Project Email Address



## GETProjectsIdEmailaddressJson

> InlineResponse20063 GETProjectsIdEmailaddressJson(ctx, id).Execute()

Get Project Email Address



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
    resp, r, err := api_client.ProjectEmailAddressesApi.GETProjectsIdEmailaddressJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectEmailAddressesApi.GETProjectsIdEmailaddressJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsIdEmailaddressJson`: InlineResponse20063
    fmt.Fprintf(os.Stdout, "Response from `ProjectEmailAddressesApi.GETProjectsIdEmailaddressJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsIdEmailaddressJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20063**](InlineResponse20063.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsIdEmailaddressJson

> InlineResponse20063 PUTProjectsIdEmailaddressJson(ctx, id).Body(body).Execute()

Update Project Email Address



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
    body := *openapiclient.NewInlineObject57() // InlineObject57 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectEmailAddressesApi.PUTProjectsIdEmailaddressJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectEmailAddressesApi.PUTProjectsIdEmailaddressJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsIdEmailaddressJson`: InlineResponse20063
    fmt.Fprintf(os.Stdout, "Response from `ProjectEmailAddressesApi.PUTProjectsIdEmailaddressJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsIdEmailaddressJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject57**](InlineObject57.md) |  | 

### Return type

[**InlineResponse20063**](InlineResponse20063.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

