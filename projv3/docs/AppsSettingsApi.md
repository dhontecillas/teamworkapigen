# \AppsSettingsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PUTProjectsApiV3AppsappIdSettingssettingIdJson**](AppsSettingsApi.md#PUTProjectsApiV3AppsappIdSettingssettingIdJson) | **Put** /projects/api/v3/apps/{appId}/settings/{settingId}.json | Update an existing setting.



## PUTProjectsApiV3AppsappIdSettingssettingIdJson

> SettingResponse PUTProjectsApiV3AppsappIdSettingssettingIdJson(ctx, settingId, appId).SettingRequest(settingRequest).Execute()

Update an existing setting.

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
    settingId := int32(56) // int32 | 
    appId := int32(56) // int32 | 
    settingRequest := *openapiclient.NewSettingRequest() // SettingRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsSettingsApi.PUTProjectsApiV3AppsappIdSettingssettingIdJson(context.Background(), settingId, appId).SettingRequest(settingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsSettingsApi.PUTProjectsApiV3AppsappIdSettingssettingIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsApiV3AppsappIdSettingssettingIdJson`: SettingResponse
    fmt.Fprintf(os.Stdout, "Response from `AppsSettingsApi.PUTProjectsApiV3AppsappIdSettingssettingIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settingId** | **int32** |  | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3AppsappIdSettingssettingIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **settingRequest** | [**SettingRequest**](SettingRequest.md) |  | 

### Return type

[**SettingResponse**](SettingResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

