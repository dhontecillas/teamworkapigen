# \FileUploadingApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**POSTPendingfilesJson**](FileUploadingApi.md#POSTPendingfilesJson) | **Post** /pendingfiles.json | Upload a File (Classic)
[**PUTProjectsApiV1PendingfilesPresignedurlJson**](FileUploadingApi.md#PUTProjectsApiV1PendingfilesPresignedurlJson) | **Put** /projects/api/v1/pendingfiles/presignedurl.json | File Upload (Preferred)



## POSTPendingfilesJson

> InlineResponse2011 POSTPendingfilesJson(ctx).Body(body).Execute()

Upload a File (Classic)



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
    body := *openapiclient.NewInlineObject36(map[string]interface{}(123)) // InlineObject36 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FileUploadingApi.POSTPendingfilesJson(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileUploadingApi.POSTPendingfilesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTPendingfilesJson`: InlineResponse2011
    fmt.Fprintf(os.Stdout, "Response from `FileUploadingApi.POSTPendingfilesJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTPendingfilesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InlineObject36**](InlineObject36.md) |  | 

### Return type

[**InlineResponse2011**](InlineResponse2011.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV1PendingfilesPresignedurlJson

> map[string]interface{} PUTProjectsApiV1PendingfilesPresignedurlJson(ctx).Execute()

File Upload (Preferred)



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
    resp, r, err := api_client.FileUploadingApi.PUTProjectsApiV1PendingfilesPresignedurlJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileUploadingApi.PUTProjectsApiV1PendingfilesPresignedurlJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsApiV1PendingfilesPresignedurlJson`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `FileUploadingApi.PUTProjectsApiV1PendingfilesPresignedurlJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV1PendingfilesPresignedurlJsonRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

