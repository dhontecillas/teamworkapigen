# \NotebookCommentsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3NotebooksnotebookIdCommentsJson**](NotebookCommentsApi.md#GETProjectsApiV3NotebooksnotebookIdCommentsJson) | **Get** /projects/api/v3/notebooks/{notebookId}/comments.json | Get a list of comments for a notebook



## GETProjectsApiV3NotebooksnotebookIdCommentsJson

> CommentCommentsResponse GETProjectsApiV3NotebooksnotebookIdCommentsJson(ctx, notebookId).Execute()

Get a list of comments for a notebook

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
    notebookId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotebookCommentsApi.GETProjectsApiV3NotebooksnotebookIdCommentsJson(context.Background(), notebookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebookCommentsApi.GETProjectsApiV3NotebooksnotebookIdCommentsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3NotebooksnotebookIdCommentsJson`: CommentCommentsResponse
    fmt.Fprintf(os.Stdout, "Response from `NotebookCommentsApi.GETProjectsApiV3NotebooksnotebookIdCommentsJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notebookId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3NotebooksnotebookIdCommentsJsonRequest struct via the builder pattern


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

