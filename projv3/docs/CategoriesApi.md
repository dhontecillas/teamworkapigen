# \CategoriesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3ProjectsTeamworkCategoriesJson**](CategoriesApi.md#GETProjectsApiV3ProjectsTeamworkCategoriesJson) | **Get** /projects/api/v3/projects/teamwork/categories.json | Returns a list of teamwork project categories



## GETProjectsApiV3ProjectsTeamworkCategoriesJson

> CategoryCategoriesResponse GETProjectsApiV3ProjectsTeamworkCategoriesJson(ctx).Execute()

Returns a list of teamwork project categories

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
    resp, r, err := api_client.CategoriesApi.GETProjectsApiV3ProjectsTeamworkCategoriesJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.GETProjectsApiV3ProjectsTeamworkCategoriesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ProjectsTeamworkCategoriesJson`: CategoryCategoriesResponse
    fmt.Fprintf(os.Stdout, "Response from `CategoriesApi.GETProjectsApiV3ProjectsTeamworkCategoriesJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsTeamworkCategoriesJsonRequest struct via the builder pattern


### Return type

[**CategoryCategoriesResponse**](CategoryCategoriesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

