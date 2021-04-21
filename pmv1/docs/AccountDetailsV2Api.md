# \AccountDetailsV2Api

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV2MeJson**](AccountDetailsV2Api.md#GETProjectsApiV2MeJson) | **Get** /projects/api/v2/me.json | Get account details



## GETProjectsApiV2MeJson

> map[string]interface{} GETProjectsApiV2MeJson(ctx).GetPermissions(getPermissions).Execute()

Get account details



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
    getPermissions := true // bool | To get permissions on your account. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountDetailsV2Api.GETProjectsApiV2MeJson(context.Background()).GetPermissions(getPermissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountDetailsV2Api.GETProjectsApiV2MeJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV2MeJson`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountDetailsV2Api.GETProjectsApiV2MeJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2MeJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPermissions** | **bool** | To get permissions on your account. | [default to false]

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

