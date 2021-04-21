# \LockdownsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3LockdownslockdownIdJson**](LockdownsApi.md#GETProjectsApiV3LockdownslockdownIdJson) | **Get** /projects/api/v3/lockdowns/{lockdownId}.json | Get a specific lockdown.



## GETProjectsApiV3LockdownslockdownIdJson

> LockdownResponse GETProjectsApiV3LockdownslockdownIdJson(ctx, lockdownId).Id(id).IncludeItems(includeItems).Include(include).FieldsUsers(fieldsUsers).FieldsTeams(fieldsTeams).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsNotebooks(fieldsNotebooks).FieldsMilestones(fieldsMilestones).FieldsMessages(fieldsMessages).FieldsLockdowns(fieldsLockdowns).FieldsLinks(fieldsLinks).FieldsFiles(fieldsFiles).FieldsCompanies(fieldsCompanies).FieldsComments(fieldsComments).Execute()

Get a specific lockdown.



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
    lockdownId := int32(56) // int32 | 
    id := int32(56) // int32 | filter by lockdown ID (optional)
    includeItems := true // bool | include items (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTeams := []string{"Inner_example"} // []string |  (optional)
    fieldsTasks := []string{"Inner_example"} // []string |  (optional)
    fieldsTasklists := []string{"Inner_example"} // []string |  (optional)
    fieldsNotebooks := []string{"Inner_example"} // []string |  (optional)
    fieldsMilestones := []string{"Inner_example"} // []string |  (optional)
    fieldsMessages := []string{"Inner_example"} // []string |  (optional)
    fieldsLockdowns := []string{"Inner_example"} // []string |  (optional)
    fieldsLinks := []string{"Inner_example"} // []string |  (optional)
    fieldsFiles := []string{"Inner_example"} // []string |  (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)
    fieldsComments := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LockdownsApi.GETProjectsApiV3LockdownslockdownIdJson(context.Background(), lockdownId).Id(id).IncludeItems(includeItems).Include(include).FieldsUsers(fieldsUsers).FieldsTeams(fieldsTeams).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsNotebooks(fieldsNotebooks).FieldsMilestones(fieldsMilestones).FieldsMessages(fieldsMessages).FieldsLockdowns(fieldsLockdowns).FieldsLinks(fieldsLinks).FieldsFiles(fieldsFiles).FieldsCompanies(fieldsCompanies).FieldsComments(fieldsComments).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LockdownsApi.GETProjectsApiV3LockdownslockdownIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3LockdownslockdownIdJson`: LockdownResponse
    fmt.Fprintf(os.Stdout, "Response from `LockdownsApi.GETProjectsApiV3LockdownslockdownIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lockdownId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3LockdownslockdownIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **int32** | filter by lockdown ID | 
 **includeItems** | **bool** | include items | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTeams** | **[]string** |  | 
 **fieldsTasks** | **[]string** |  | 
 **fieldsTasklists** | **[]string** |  | 
 **fieldsNotebooks** | **[]string** |  | 
 **fieldsMilestones** | **[]string** |  | 
 **fieldsMessages** | **[]string** |  | 
 **fieldsLockdowns** | **[]string** |  | 
 **fieldsLinks** | **[]string** |  | 
 **fieldsFiles** | **[]string** |  | 
 **fieldsCompanies** | **[]string** |  | 
 **fieldsComments** | **[]string** |  | 

### Return type

[**LockdownResponse**](LockdownResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

