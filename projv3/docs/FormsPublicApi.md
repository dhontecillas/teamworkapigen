# \FormsPublicApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3FormsPublictokenJson**](FormsPublicApi.md#GETProjectsApiV3FormsPublictokenJson) | **Get** /projects/api/v3/forms/public/{token}.json | Get a forms via its token.
[**POSTProjectsApiV3FormsPublictokenSubmitJson**](FormsPublicApi.md#POSTProjectsApiV3FormsPublictokenSubmitJson) | **Post** /projects/api/v3/forms/public/{token}/submit.json | Submit a form response



## GETProjectsApiV3FormsPublictokenJson

> FormPublicResponse GETProjectsApiV3FormsPublictokenJson(ctx, token).OrderMode(orderMode).HostObjectType(hostObjectType).ContentState(contentState).UserId(userId).PageSize(pageSize).Page(page).HostObjectId(hostObjectId).ProjectIds(projectIds).Include(include).FieldsUsers(fieldsUsers).FieldsProjects(fieldsProjects).FieldsForms(fieldsForms).Execute()

Get a forms via its token.



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
    token := int32(56) // int32 | 
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    hostObjectType := "hostObjectType_example" // string | query by hostObject type (optional)
    contentState := "contentState_example" // string | query by form state (optional)
    userId := int32(56) // int32 | filter by user id (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    hostObjectId := int32(56) // int32 | filter by host id (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsForms := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FormsPublicApi.GETProjectsApiV3FormsPublictokenJson(context.Background(), token).OrderMode(orderMode).HostObjectType(hostObjectType).ContentState(contentState).UserId(userId).PageSize(pageSize).Page(page).HostObjectId(hostObjectId).ProjectIds(projectIds).Include(include).FieldsUsers(fieldsUsers).FieldsProjects(fieldsProjects).FieldsForms(fieldsForms).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FormsPublicApi.GETProjectsApiV3FormsPublictokenJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3FormsPublictokenJson`: FormPublicResponse
    fmt.Fprintf(os.Stdout, "Response from `FormsPublicApi.GETProjectsApiV3FormsPublictokenJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3FormsPublictokenJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **hostObjectType** | **string** | query by hostObject type | 
 **contentState** | **string** | query by form state | 
 **userId** | **int32** | filter by user id | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **hostObjectId** | **int32** | filter by host id | 
 **projectIds** | **[]int32** | filter by project ids | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsForms** | **[]string** |  | 

### Return type

[**FormPublicResponse**](FormPublicResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsApiV3FormsPublictokenSubmitJson

> FormFormsResponse POSTProjectsApiV3FormsPublictokenSubmitJson(ctx, token).FormSubmissionRequest(formSubmissionRequest).Execute()

Submit a form response



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
    token := int32(56) // int32 | 
    formSubmissionRequest := *openapiclient.NewFormSubmissionRequest() // FormSubmissionRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FormsPublicApi.POSTProjectsApiV3FormsPublictokenSubmitJson(context.Background(), token).FormSubmissionRequest(formSubmissionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FormsPublicApi.POSTProjectsApiV3FormsPublictokenSubmitJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsApiV3FormsPublictokenSubmitJson`: FormFormsResponse
    fmt.Fprintf(os.Stdout, "Response from `FormsPublicApi.POSTProjectsApiV3FormsPublictokenSubmitJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3FormsPublictokenSubmitJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **formSubmissionRequest** | [**FormSubmissionRequest**](FormSubmissionRequest.md) |  | 

### Return type

[**FormFormsResponse**](FormFormsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

