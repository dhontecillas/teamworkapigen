# \FileCommentsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3FilesfileIdCommentsJson**](FileCommentsApi.md#GETProjectsApiV3FilesfileIdCommentsJson) | **Get** /projects/api/v3/files/{fileId}/comments.json | Get a list of comments for a file



## GETProjectsApiV3FilesfileIdCommentsJson

> CommentCommentsResponse GETProjectsApiV3FilesfileIdCommentsJson(ctx, fileId).Execute()

Get a list of comments for a file

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
    fileId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FileCommentsApi.GETProjectsApiV3FilesfileIdCommentsJson(context.Background(), fileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileCommentsApi.GETProjectsApiV3FilesfileIdCommentsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3FilesfileIdCommentsJson`: CommentCommentsResponse
    fmt.Fprintf(os.Stdout, "Response from `FileCommentsApi.GETProjectsApiV3FilesfileIdCommentsJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3FilesfileIdCommentsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommentCommentsResponse**](CommentCommentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

