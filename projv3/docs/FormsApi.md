# \FormsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEProjectsApiV3FormsformIdJson**](FormsApi.md#DELETEProjectsApiV3FormsformIdJson) | **Delete** /projects/api/v3/forms/{formId}.json | Delete an existing form.
[**GETProjectsApiV3FormsJson**](FormsApi.md#GETProjectsApiV3FormsJson) | **Get** /projects/api/v3/forms.json | Get all forms.
[**GETProjectsApiV3FormsformIdDraftJson**](FormsApi.md#GETProjectsApiV3FormsformIdDraftJson) | **Get** /projects/api/v3/forms/{formId}/draft.json | Get the draft version of a specific form.
[**GETProjectsApiV3FormsformIdJson**](FormsApi.md#GETProjectsApiV3FormsformIdJson) | **Get** /projects/api/v3/forms/{formId}.json | Get a specific form.
[**PATCHProjectsApiV3FormsformIdJson**](FormsApi.md#PATCHProjectsApiV3FormsformIdJson) | **Patch** /projects/api/v3/forms/{formId}.json | Update an existing form.
[**POSTProjectsApiV3FormsJson**](FormsApi.md#POSTProjectsApiV3FormsJson) | **Post** /projects/api/v3/forms.json | Create a new form.



## DELETEProjectsApiV3FormsformIdJson

> DELETEProjectsApiV3FormsformIdJson(ctx, formId).Execute()

Delete an existing form.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FormsApi.DELETEProjectsApiV3FormsformIdJson(context.Background(), formId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FormsApi.DELETEProjectsApiV3FormsformIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEProjectsApiV3FormsformIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3FormsJson

> FormFormsResponse GETProjectsApiV3FormsJson(ctx).Execute()

Get all forms.



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
    resp, r, err := api_client.FormsApi.GETProjectsApiV3FormsJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FormsApi.GETProjectsApiV3FormsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3FormsJson`: FormFormsResponse
    fmt.Fprintf(os.Stdout, "Response from `FormsApi.GETProjectsApiV3FormsJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3FormsJsonRequest struct via the builder pattern


### Return type

[**FormFormsResponse**](FormFormsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3FormsformIdDraftJson

> FormResponse GETProjectsApiV3FormsformIdDraftJson(ctx, formId).Execute()

Get the draft version of a specific form.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FormsApi.GETProjectsApiV3FormsformIdDraftJson(context.Background(), formId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FormsApi.GETProjectsApiV3FormsformIdDraftJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3FormsformIdDraftJson`: FormResponse
    fmt.Fprintf(os.Stdout, "Response from `FormsApi.GETProjectsApiV3FormsformIdDraftJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3FormsformIdDraftJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FormResponse**](FormResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3FormsformIdJson

> FormResponse GETProjectsApiV3FormsformIdJson(ctx, formId).Execute()

Get a specific form.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FormsApi.GETProjectsApiV3FormsformIdJson(context.Background(), formId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FormsApi.GETProjectsApiV3FormsformIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3FormsformIdJson`: FormResponse
    fmt.Fprintf(os.Stdout, "Response from `FormsApi.GETProjectsApiV3FormsformIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3FormsformIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FormResponse**](FormResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHProjectsApiV3FormsformIdJson

> FormResponse PATCHProjectsApiV3FormsformIdJson(ctx, formId).FormRequest(formRequest).Execute()

Update an existing form.

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
    formRequest := *openapiclient.NewFormRequest() // FormRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FormsApi.PATCHProjectsApiV3FormsformIdJson(context.Background(), formId).FormRequest(formRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FormsApi.PATCHProjectsApiV3FormsformIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHProjectsApiV3FormsformIdJson`: FormResponse
    fmt.Fprintf(os.Stdout, "Response from `FormsApi.PATCHProjectsApiV3FormsformIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHProjectsApiV3FormsformIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **formRequest** | [**FormRequest**](FormRequest.md) |  | 

### Return type

[**FormResponse**](FormResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsApiV3FormsJson

> FormResponse POSTProjectsApiV3FormsJson(ctx).FormRequest(formRequest).Execute()

Create a new form.

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
    formRequest := *openapiclient.NewFormRequest() // FormRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FormsApi.POSTProjectsApiV3FormsJson(context.Background()).FormRequest(formRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FormsApi.POSTProjectsApiV3FormsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsApiV3FormsJson`: FormResponse
    fmt.Fprintf(os.Stdout, "Response from `FormsApi.POSTProjectsApiV3FormsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3FormsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **formRequest** | [**FormRequest**](FormRequest.md) |  | 

### Return type

[**FormResponse**](FormResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

