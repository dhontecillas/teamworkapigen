# \TimezonesApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETTimezonesJson**](TimezonesApi.md#GETTimezonesJson) | **Get** /timezones.json | Get Timezones



## GETTimezonesJson

> InlineResponse200113 GETTimezonesJson(ctx).Execute()

Get Timezones



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
    resp, r, err := api_client.TimezonesApi.GETTimezonesJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimezonesApi.GETTimezonesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETTimezonesJson`: InlineResponse200113
    fmt.Fprintf(os.Stdout, "Response from `TimezonesApi.GETTimezonesJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETTimezonesJsonRequest struct via the builder pattern


### Return type

[**InlineResponse200113**](InlineResponse200113.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

