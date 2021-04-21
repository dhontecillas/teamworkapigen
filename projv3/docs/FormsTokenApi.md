# \FormsTokenApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PATCHProjectsApiV3FormsformIdTokenJson**](FormsTokenApi.md#PATCHProjectsApiV3FormsformIdTokenJson) | **Patch** /projects/api/v3/forms/{formId}/token.json | Update an existing token.
[**PUTProjectsApiV3FormsformIdTokenRefreshJson**](FormsTokenApi.md#PUTProjectsApiV3FormsformIdTokenRefreshJson) | **Put** /projects/api/v3/forms/{formId}/token/refresh.json | Refresh the value of a token



## PATCHProjectsApiV3FormsformIdTokenJson

> TokenResponse PATCHProjectsApiV3FormsformIdTokenJson(ctx, formId).TokenRequest(tokenRequest).Execute()

Update an existing token.



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
    formId := int32(56) // int32 | 
    tokenRequest := *openapiclient.NewTokenRequest() // TokenRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FormsTokenApi.PATCHProjectsApiV3FormsformIdTokenJson(context.Background(), formId).TokenRequest(tokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FormsTokenApi.PATCHProjectsApiV3FormsformIdTokenJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHProjectsApiV3FormsformIdTokenJson`: TokenResponse
    fmt.Fprintf(os.Stdout, "Response from `FormsTokenApi.PATCHProjectsApiV3FormsformIdTokenJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHProjectsApiV3FormsformIdTokenJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenRequest** | [**TokenRequest**](TokenRequest.md) |  | 

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV3FormsformIdTokenRefreshJson

> TokenResponse PUTProjectsApiV3FormsformIdTokenRefreshJson(ctx, formId).TokenRequest(tokenRequest).Execute()

Refresh the value of a token



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
    formId := int32(56) // int32 | 
    tokenRequest := *openapiclient.NewTokenRequest() // TokenRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FormsTokenApi.PUTProjectsApiV3FormsformIdTokenRefreshJson(context.Background(), formId).TokenRequest(tokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FormsTokenApi.PUTProjectsApiV3FormsformIdTokenRefreshJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsApiV3FormsformIdTokenRefreshJson`: TokenResponse
    fmt.Fprintf(os.Stdout, "Response from `FormsTokenApi.PUTProjectsApiV3FormsformIdTokenRefreshJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3FormsformIdTokenRefreshJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenRequest** | [**TokenRequest**](TokenRequest.md) |  | 

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

