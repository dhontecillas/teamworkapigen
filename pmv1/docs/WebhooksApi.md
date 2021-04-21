# \WebhooksApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEWebhooksIdJson**](WebhooksApi.md#DELETEWebhooksIdJson) | **Delete** /webhooks/{id}.json | Delete a specific Webhook on your Account
[**GETWebhookidDeliveriesJson**](WebhooksApi.md#GETWebhookidDeliveriesJson) | **Get** /{webhookid}/deliveries.json | Get deliveries from specific Webhook
[**GETWebhooksEventsJson**](WebhooksApi.md#GETWebhooksEventsJson) | **Get** /webhooks/events.json | Get a list of all Webhook Events that can be used
[**GETWebhooksIdJson**](WebhooksApi.md#GETWebhooksIdJson) | **Get** /webhooks/{id}.json | Get details of a specific Webhook set on your account
[**GETWebhooksJson**](WebhooksApi.md#GETWebhooksJson) | **Get** /webhooks.json | Get all Webhooks set on your Account
[**POSTWebhookWebhookidRedeliverJson**](WebhooksApi.md#POSTWebhookWebhookidRedeliverJson) | **Post** /webhook/{webhookid}/redeliver.json | Resend a Specific Webhook
[**POSTWebhooksJson**](WebhooksApi.md#POSTWebhooksJson) | **Post** /webhooks.json | Create a new Webhook on your Account
[**PUTWebhooksDisableJson**](WebhooksApi.md#PUTWebhooksDisableJson) | **Put** /webhooks/disable.json | Disable Webhooks on your Teamwork.com Projects Account
[**PUTWebhooksEnableJson**](WebhooksApi.md#PUTWebhooksEnableJson) | **Put** /webhooks/enable.json | Enable Webhooks on your Teamwork.com Projects Account
[**PUTWebhooksIdJson**](WebhooksApi.md#PUTWebhooksIdJson) | **Put** /webhooks/{id}.json | Update a specific Webhook set on your Account
[**PUTWebhooksIdPauseJson**](WebhooksApi.md#PUTWebhooksIdPauseJson) | **Put** /webhooks/{id}/pause.json | Pause a specific Webhook set on your Account
[**PUTWebhooksIdResumeJson**](WebhooksApi.md#PUTWebhooksIdResumeJson) | **Put** /webhooks/{id}/resume.json | Resume a specific Webhook set on your Account



## DELETEWebhooksIdJson

> InlineResponse2001 DELETEWebhooksIdJson(ctx, id).Execute()

Delete a specific Webhook on your Account



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
    resp, r, err := api_client.WebhooksApi.DELETEWebhooksIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.DELETEWebhooksIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DELETEWebhooksIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.DELETEWebhooksIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEWebhooksIdJsonRequest struct via the builder pattern


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


## GETWebhookidDeliveriesJson

> InlineResponse200124 GETWebhookidDeliveriesJson(ctx, webhookid).Execute()

Get deliveries from specific Webhook



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
    webhookid := "webhookid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhooksApi.GETWebhookidDeliveriesJson(context.Background(), webhookid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.GETWebhookidDeliveriesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETWebhookidDeliveriesJson`: InlineResponse200124
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.GETWebhookidDeliveriesJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETWebhookidDeliveriesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200124**](InlineResponse200124.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETWebhooksEventsJson

> InlineResponse200118 GETWebhooksEventsJson(ctx).Execute()

Get a list of all Webhook Events that can be used



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
    resp, r, err := api_client.WebhooksApi.GETWebhooksEventsJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.GETWebhooksEventsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETWebhooksEventsJson`: InlineResponse200118
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.GETWebhooksEventsJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETWebhooksEventsJsonRequest struct via the builder pattern


### Return type

[**InlineResponse200118**](InlineResponse200118.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETWebhooksIdJson

> InlineResponse200119 GETWebhooksIdJson(ctx, id).Execute()

Get details of a specific Webhook set on your account



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
    resp, r, err := api_client.WebhooksApi.GETWebhooksIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.GETWebhooksIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETWebhooksIdJson`: InlineResponse200119
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.GETWebhooksIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETWebhooksIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200119**](InlineResponse200119.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETWebhooksJson

> InlineResponse200117 GETWebhooksJson(ctx).Execute()

Get all Webhooks set on your Account



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
    resp, r, err := api_client.WebhooksApi.GETWebhooksJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.GETWebhooksJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETWebhooksJson`: InlineResponse200117
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.GETWebhooksJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETWebhooksJsonRequest struct via the builder pattern


### Return type

[**InlineResponse200117**](InlineResponse200117.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTWebhookWebhookidRedeliverJson

> InlineResponse2001 POSTWebhookWebhookidRedeliverJson(ctx, webhookid).Body(body).Execute()

Resend a Specific Webhook



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
    webhookid := "webhookid_example" // string | 
    body := *openapiclient.NewInlineObject106() // InlineObject106 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhooksApi.POSTWebhookWebhookidRedeliverJson(context.Background(), webhookid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.POSTWebhookWebhookidRedeliverJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTWebhookWebhookidRedeliverJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.POSTWebhookWebhookidRedeliverJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTWebhookWebhookidRedeliverJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject106**](InlineObject106.md) |  | 

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


## POSTWebhooksJson

> InlineResponse20018 POSTWebhooksJson(ctx).Body(body).Execute()

Create a new Webhook on your Account



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
    body := *openapiclient.NewInlineObject107() // InlineObject107 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhooksApi.POSTWebhooksJson(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.POSTWebhooksJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTWebhooksJson`: InlineResponse20018
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.POSTWebhooksJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTWebhooksJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InlineObject107**](InlineObject107.md) |  | 

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


## PUTWebhooksDisableJson

> InlineResponse2001 PUTWebhooksDisableJson(ctx).Execute()

Disable Webhooks on your Teamwork.com Projects Account



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
    resp, r, err := api_client.WebhooksApi.PUTWebhooksDisableJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.PUTWebhooksDisableJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTWebhooksDisableJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.PUTWebhooksDisableJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPUTWebhooksDisableJsonRequest struct via the builder pattern


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


## PUTWebhooksEnableJson

> InlineResponse2001 PUTWebhooksEnableJson(ctx).Execute()

Enable Webhooks on your Teamwork.com Projects Account



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
    resp, r, err := api_client.WebhooksApi.PUTWebhooksEnableJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.PUTWebhooksEnableJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTWebhooksEnableJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.PUTWebhooksEnableJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPUTWebhooksEnableJsonRequest struct via the builder pattern


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


## PUTWebhooksIdJson

> InlineResponse2001 PUTWebhooksIdJson(ctx, id).Body(body).Execute()

Update a specific Webhook set on your Account



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
    body := *openapiclient.NewInlineObject108() // InlineObject108 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhooksApi.PUTWebhooksIdJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.PUTWebhooksIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTWebhooksIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.PUTWebhooksIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTWebhooksIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject108**](InlineObject108.md) |  | 

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


## PUTWebhooksIdPauseJson

> InlineResponse2001 PUTWebhooksIdPauseJson(ctx, id).Execute()

Pause a specific Webhook set on your Account



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
    resp, r, err := api_client.WebhooksApi.PUTWebhooksIdPauseJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.PUTWebhooksIdPauseJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTWebhooksIdPauseJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.PUTWebhooksIdPauseJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTWebhooksIdPauseJsonRequest struct via the builder pattern


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


## PUTWebhooksIdResumeJson

> InlineResponse2001 PUTWebhooksIdResumeJson(ctx, id).Execute()

Resume a specific Webhook set on your Account



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
    resp, r, err := api_client.WebhooksApi.PUTWebhooksIdResumeJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.PUTWebhooksIdResumeJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTWebhooksIdResumeJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.PUTWebhooksIdResumeJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTWebhooksIdResumeJsonRequest struct via the builder pattern


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

