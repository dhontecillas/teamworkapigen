# \InvoicesApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEInvoicesIdJson**](InvoicesApi.md#DELETEInvoicesIdJson) | **Delete** /invoices/{id}.json | Delete a Specific Invoice
[**GETCurrencycodesJson**](InvoicesApi.md#GETCurrencycodesJson) | **Get** /currencycodes.json | Get a list of valid currency Codes
[**GETInvoicesInvoiceIdJson**](InvoicesApi.md#GETInvoicesInvoiceIdJson) | **Get** /invoices/{invoiceId}.json | Get invoice Line Items
[**GETInvoicesJson**](InvoicesApi.md#GETInvoicesJson) | **Get** /invoices.json | Get all Invoices across your Projects
[**GETProjectsIdInvoicesJson**](InvoicesApi.md#GETProjectsIdInvoicesJson) | **Get** /projects/{id}/invoices.json | Get all invoices on a single Project
[**POSTProjectsIdInvoicesJson**](InvoicesApi.md#POSTProjectsIdInvoicesJson) | **Post** /projects/{id}/invoices.json | Create a New Invoice in a Project
[**PUTInvoicesIdCompleteJson**](InvoicesApi.md#PUTInvoicesIdCompleteJson) | **Put** /invoices/{id}/complete.json | Mark a specific Invoice as Complete
[**PUTInvoicesIdJson**](InvoicesApi.md#PUTInvoicesIdJson) | **Put** /invoices/{id}.json | Update a Specific Invoice
[**PUTInvoicesIdLineitemsJson**](InvoicesApi.md#PUTInvoicesIdLineitemsJson) | **Put** /invoices/{id}/lineitems.json | Add a Time Entry to an Invoice
[**PUTInvoicesIdUncompleteJson**](InvoicesApi.md#PUTInvoicesIdUncompleteJson) | **Put** /invoices/{id}/uncomplete.json | Mark a specific Invoice as not Complete



## DELETEInvoicesIdJson

> InlineResponse2001 DELETEInvoicesIdJson(ctx, id).Execute()

Delete a Specific Invoice



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
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvoicesApi.DELETEInvoicesIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.DELETEInvoicesIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DELETEInvoicesIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.DELETEInvoicesIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEInvoicesIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETCurrencycodesJson

> InlineResponse20016 GETCurrencycodesJson(ctx).Execute()

Get a list of valid currency Codes



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
    resp, r, err := api_client.InvoicesApi.GETCurrencycodesJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.GETCurrencycodesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCurrencycodesJson`: InlineResponse20016
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.GETCurrencycodesJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETCurrencycodesJsonRequest struct via the builder pattern


### Return type

[**InlineResponse20016**](InlineResponse20016.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETInvoicesInvoiceIdJson

> InlineResponse20023 GETInvoicesInvoiceIdJson(ctx, invoiceId).Execute()

Get invoice Line Items



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
    invoiceId := "invoiceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvoicesApi.GETInvoicesInvoiceIdJson(context.Background(), invoiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.GETInvoicesInvoiceIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETInvoicesInvoiceIdJson`: InlineResponse20023
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.GETInvoicesInvoiceIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETInvoicesInvoiceIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20023**](InlineResponse20023.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETInvoicesJson

> InlineResponse20022 GETInvoicesJson(ctx).Page(page).Type_(type_).Execute()

Get all Invoices across your Projects



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
    page := int32(56) // int32 | The page of results to return Inspect the response headers for x-page, x-pages, x-records (optional)
    type_ := "type__example" // string | Which types of invoices to return - options are all, completed or active (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvoicesApi.GETInvoicesJson(context.Background()).Page(page).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.GETInvoicesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETInvoicesJson`: InlineResponse20022
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.GETInvoicesJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETInvoicesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page of results to return Inspect the response headers for x-page, x-pages, x-records | 
 **type_** | **string** | Which types of invoices to return - options are all, completed or active | 

### Return type

[**InlineResponse20022**](InlineResponse20022.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsIdInvoicesJson

> InlineResponse20022 GETProjectsIdInvoicesJson(ctx, id).Type_(type_).Page(page).Execute()

Get all invoices on a single Project



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
    id := int32(56) // int32 | 
    type_ := "type__example" // string | Which types of invoices to return. Options are: all,completed or active (optional)
    page := int32(56) // int32 | The page of results to return Inspect the response headers for x-page, x-pages, x-records (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvoicesApi.GETProjectsIdInvoicesJson(context.Background(), id).Type_(type_).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.GETProjectsIdInvoicesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsIdInvoicesJson`: InlineResponse20022
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.GETProjectsIdInvoicesJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsIdInvoicesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Which types of invoices to return. Options are: all,completed or active | 
 **page** | **int32** | The page of results to return Inspect the response headers for x-page, x-pages, x-records | 

### Return type

[**InlineResponse20022**](InlineResponse20022.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsIdInvoicesJson

> InlineResponse20018 POSTProjectsIdInvoicesJson(ctx, id).Body(body).Execute()

Create a New Invoice in a Project



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
    id := int32(56) // int32 | 
    body := *openapiclient.NewInlineObject60(*openapiclient.NewProjectsIdInvoicesJsonInvoice("DisplayDate_example", "Number_example")) // InlineObject60 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvoicesApi.POSTProjectsIdInvoicesJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.POSTProjectsIdInvoicesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsIdInvoicesJson`: InlineResponse20018
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.POSTProjectsIdInvoicesJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsIdInvoicesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject60**](InlineObject60.md) |  | 

### Return type

[**InlineResponse20018**](InlineResponse20018.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTInvoicesIdCompleteJson

> InlineResponse2001 PUTInvoicesIdCompleteJson(ctx, id).Execute()

Mark a specific Invoice as Complete



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
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvoicesApi.PUTInvoicesIdCompleteJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.PUTInvoicesIdCompleteJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTInvoicesIdCompleteJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.PUTInvoicesIdCompleteJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTInvoicesIdCompleteJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTInvoicesIdJson

> InlineResponse2001 PUTInvoicesIdJson(ctx, id).Body(body).Execute()

Update a Specific Invoice



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
    id := int32(56) // int32 | 
    body := *openapiclient.NewInlineObject21(*openapiclient.NewInvoicesIdJsonInvoice()) // InlineObject21 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvoicesApi.PUTInvoicesIdJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.PUTInvoicesIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTInvoicesIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.PUTInvoicesIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTInvoicesIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject21**](InlineObject21.md) |  | 

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


## PUTInvoicesIdLineitemsJson

> InlineResponse2001 PUTInvoicesIdLineitemsJson(ctx, id).Body(body).Execute()

Add a Time Entry to an Invoice



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
    id := int32(56) // int32 | 
    body := *openapiclient.NewInlineObject22() // InlineObject22 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvoicesApi.PUTInvoicesIdLineitemsJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.PUTInvoicesIdLineitemsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTInvoicesIdLineitemsJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.PUTInvoicesIdLineitemsJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTInvoicesIdLineitemsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject22**](InlineObject22.md) |  | 

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


## PUTInvoicesIdUncompleteJson

> InlineResponse2001 PUTInvoicesIdUncompleteJson(ctx, id).Execute()

Mark a specific Invoice as not Complete



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
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvoicesApi.PUTInvoicesIdUncompleteJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.PUTInvoicesIdUncompleteJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTInvoicesIdUncompleteJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.PUTInvoicesIdUncompleteJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTInvoicesIdUncompleteJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

