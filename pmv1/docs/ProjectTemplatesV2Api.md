# \ProjectTemplatesV2Api

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEProjectsTemplateTemplateIDJson**](ProjectTemplatesV2Api.md#DELETEProjectsTemplateTemplateIDJson) | **Delete** /projects/template/{TemplateID}.json | Delete a Project Template
[**GETProjectsApiV2ProjectsTemplateIdJson**](ProjectTemplatesV2Api.md#GETProjectsApiV2ProjectsTemplateIdJson) | **Get** /projects/api/v2/projects/{templateId}.json | Get a Project Template
[**GETProjectsApiV2ProjectsTemplatesJson**](ProjectTemplatesV2Api.md#GETProjectsApiV2ProjectsTemplatesJson) | **Get** /projects/api/v2/projects/templates.json | Get Project Templates
[**POSTProjectsTemplateJson**](ProjectTemplatesV2Api.md#POSTProjectsTemplateJson) | **Post** /projects/template.json | Create a Project Template
[**PUTProjectsTemplateTemplateIDJson**](ProjectTemplatesV2Api.md#PUTProjectsTemplateTemplateIDJson) | **Put** /projects/template/{TemplateID}.json | Update a specific Project Template



## DELETEProjectsTemplateTemplateIDJson

> InlineResponse2001 DELETEProjectsTemplateTemplateIDJson(ctx, templateID).Execute()

Delete a Project Template



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
    templateID := "templateID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectTemplatesV2Api.DELETEProjectsTemplateTemplateIDJson(context.Background(), templateID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplatesV2Api.DELETEProjectsTemplateTemplateIDJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DELETEProjectsTemplateTemplateIDJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `ProjectTemplatesV2Api.DELETEProjectsTemplateTemplateIDJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEProjectsTemplateTemplateIDJsonRequest struct via the builder pattern


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


## GETProjectsApiV2ProjectsTemplateIdJson

> InlineResponse20054 GETProjectsApiV2ProjectsTemplateIdJson(ctx, templateId, getpermissions).Execute()

Get a Project Template



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
    templateId := "templateId_example" // string | 
    getpermissions := true // bool | If this a required value in order to return the template. Values = true/false

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectTemplatesV2Api.GETProjectsApiV2ProjectsTemplateIdJson(context.Background(), templateId, getpermissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplatesV2Api.GETProjectsApiV2ProjectsTemplateIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV2ProjectsTemplateIdJson`: InlineResponse20054
    fmt.Fprintf(os.Stdout, "Response from `ProjectTemplatesV2Api.GETProjectsApiV2ProjectsTemplateIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** |  | 
**getpermissions** | **bool** | If this a required value in order to return the template. Values &#x3D; true/false | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2ProjectsTemplateIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20054**](InlineResponse20054.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV2ProjectsTemplatesJson

> InlineResponse20054 GETProjectsApiV2ProjectsTemplatesJson(ctx).SearchTerm(searchTerm).OrderBy(orderBy).ProjectOwnerIds(projectOwnerIds).ProjectTagIds(projectTagIds).Execute()

Get Project Templates



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
    searchTerm := "searchTerm_example" // string |  (optional)
    orderBy := "orderBy_example" // string | Options: createdOn OR name (optional)
    projectOwnerIds := "projectOwnerIds_example" // string | Comma separated list of ID's (optional)
    projectTagIds := "projectTagIds_example" // string | Comma separated list of ID's (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectTemplatesV2Api.GETProjectsApiV2ProjectsTemplatesJson(context.Background()).SearchTerm(searchTerm).OrderBy(orderBy).ProjectOwnerIds(projectOwnerIds).ProjectTagIds(projectTagIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplatesV2Api.GETProjectsApiV2ProjectsTemplatesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV2ProjectsTemplatesJson`: InlineResponse20054
    fmt.Fprintf(os.Stdout, "Response from `ProjectTemplatesV2Api.GETProjectsApiV2ProjectsTemplatesJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2ProjectsTemplatesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchTerm** | **string** |  | 
 **orderBy** | **string** | Options: createdOn OR name | 
 **projectOwnerIds** | **string** | Comma separated list of ID&#39;s | 
 **projectTagIds** | **string** | Comma separated list of ID&#39;s | 

### Return type

[**InlineResponse20054**](InlineResponse20054.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsTemplateJson

> InlineResponse2001 POSTProjectsTemplateJson(ctx).Body(body).Execute()

Create a Project Template



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
    body := *openapiclient.NewInlineObject51() // InlineObject51 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectTemplatesV2Api.POSTProjectsTemplateJson(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplatesV2Api.POSTProjectsTemplateJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsTemplateJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `ProjectTemplatesV2Api.POSTProjectsTemplateJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsTemplateJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InlineObject51**](InlineObject51.md) |  | 

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


## PUTProjectsTemplateTemplateIDJson

> InlineResponse2001 PUTProjectsTemplateTemplateIDJson(ctx, templateID).Body(body).Execute()

Update a specific Project Template



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
    templateID := "templateID_example" // string | 
    body := *openapiclient.NewInlineObject52(*openapiclient.NewProjectsTemplateTemplateIDJsonProject("Name_example", "TemplateDateTargetDefault_example")) // InlineObject52 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectTemplatesV2Api.PUTProjectsTemplateTemplateIDJson(context.Background(), templateID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplatesV2Api.PUTProjectsTemplateTemplateIDJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsTemplateTemplateIDJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `ProjectTemplatesV2Api.PUTProjectsTemplateTemplateIDJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsTemplateTemplateIDJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject52**](InlineObject52.md) |  | 

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

