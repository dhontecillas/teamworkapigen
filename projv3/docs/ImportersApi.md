# \ImportersApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3ImporterStatsJson**](ImportersApi.md#GETProjectsApiV3ImporterStatsJson) | **Get** /projects/api/v3/importer/stats.json | Get stats about Importers



## GETProjectsApiV3ImporterStatsJson

> ImporterImportersResponse GETProjectsApiV3ImporterStatsJson(ctx).ResetImporters(resetImporters).Execute()

Get stats about Importers



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
    resetImporters := true // bool | reset importers (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImportersApi.GETProjectsApiV3ImporterStatsJson(context.Background()).ResetImporters(resetImporters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportersApi.GETProjectsApiV3ImporterStatsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ImporterStatsJson`: ImporterImportersResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportersApi.GETProjectsApiV3ImporterStatsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ImporterStatsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetImporters** | **bool** | reset importers | [default to false]

### Return type

[**ImporterImportersResponse**](ImporterImportersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

