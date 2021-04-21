# \ProjectOwnerApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETYoursiteProjectsJson**](ProjectOwnerApi.md#GETYoursiteProjectsJson) | **Get** /yoursite/projects.json | Get Project Owner
[**PUTYoursiteProjectsProjIdJson**](ProjectOwnerApi.md#PUTYoursiteProjectsProjIdJson) | **Put** /yoursite/projects/{projId}.json | Setting a Project Owner



## GETYoursiteProjectsJson

> InlineResponse200121 GETYoursiteProjectsJson(ctx).IncludeProjectOwner(includeProjectOwner).Execute()

Get Project Owner



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
    includeProjectOwner := true // bool | =true to return owner

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectOwnerApi.GETYoursiteProjectsJson(context.Background()).IncludeProjectOwner(includeProjectOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectOwnerApi.GETYoursiteProjectsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETYoursiteProjectsJson`: InlineResponse200121
    fmt.Fprintf(os.Stdout, "Response from `ProjectOwnerApi.GETYoursiteProjectsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETYoursiteProjectsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeProjectOwner** | **bool** | &#x3D;true to return owner | 

### Return type

[**InlineResponse200121**](InlineResponse200121.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTYoursiteProjectsProjIdJson

> InlineResponse2001 PUTYoursiteProjectsProjIdJson(ctx, projId).Body(body).Execute()

Setting a Project Owner



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
    projId := "projId_example" // string | 
    body := *openapiclient.NewInlineObject109() // InlineObject109 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectOwnerApi.PUTYoursiteProjectsProjIdJson(context.Background(), projId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectOwnerApi.PUTYoursiteProjectsProjIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTYoursiteProjectsProjIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `ProjectOwnerApi.PUTYoursiteProjectsProjIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTYoursiteProjectsProjIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject109**](InlineObject109.md) |  | 

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

