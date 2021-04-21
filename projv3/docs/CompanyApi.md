# \CompanyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3CompaniescompanyIdJson**](CompanyApi.md#GETProjectsApiV3CompaniescompanyIdJson) | **Get** /projects/api/v3/companies/{companyId}.json | Get a specific company.



## GETProjectsApiV3CompaniescompanyIdJson

> CompanyResponse GETProjectsApiV3CompaniescompanyIdJson(ctx, companyId).Execute()

Get a specific company.



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
    companyId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompanyApi.GETProjectsApiV3CompaniescompanyIdJson(context.Background(), companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompanyApi.GETProjectsApiV3CompaniescompanyIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3CompaniescompanyIdJson`: CompanyResponse
    fmt.Fprintf(os.Stdout, "Response from `CompanyApi.GETProjectsApiV3CompaniescompanyIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3CompaniescompanyIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CompanyResponse**](CompanyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

