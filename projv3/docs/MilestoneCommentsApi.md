# \MilestoneCommentsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3MilestonesmilestoneIdCommentsJson**](MilestoneCommentsApi.md#GETProjectsApiV3MilestonesmilestoneIdCommentsJson) | **Get** /projects/api/v3/milestones/{milestoneId}/comments.json | Get a list of comments for a milestone



## GETProjectsApiV3MilestonesmilestoneIdCommentsJson

> CommentCommentsResponse GETProjectsApiV3MilestonesmilestoneIdCommentsJson(ctx, milestoneId).Execute()

Get a list of comments for a milestone

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
    milestoneId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MilestoneCommentsApi.GETProjectsApiV3MilestonesmilestoneIdCommentsJson(context.Background(), milestoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MilestoneCommentsApi.GETProjectsApiV3MilestonesmilestoneIdCommentsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3MilestonesmilestoneIdCommentsJson`: CommentCommentsResponse
    fmt.Fprintf(os.Stdout, "Response from `MilestoneCommentsApi.GETProjectsApiV3MilestonesmilestoneIdCommentsJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**milestoneId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3MilestonesmilestoneIdCommentsJsonRequest struct via the builder pattern


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

