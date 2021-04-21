# \DashboardsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3DashboardsJson**](DashboardsApi.md#GETProjectsApiV3DashboardsJson) | **Get** /projects/api/v3/dashboards.json | Get all dashboards



## GETProjectsApiV3DashboardsJson

> DashboardUserDashboardsResponse GETProjectsApiV3DashboardsJson(ctx).UpdatedAfter(updatedAfter).OrderMode(orderMode).OrderBy(orderBy).UserId(userId).ProjectId(projectId).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).Emoji(emoji).Include(include).FieldsUsers(fieldsUsers).FieldsProjects(fieldsProjects).FieldsDashboards(fieldsDashboards).FieldsDashboardSettings(fieldsDashboardSettings).FieldsDashboardPanels(fieldsDashboardPanels).FieldsDashboardPanelSettings(fieldsDashboardPanelSettings).DashboardIds(dashboardIds).Execute()

Get all dashboards



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    updatedAfter := time.Now() // string | filter by updated after date (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "isDefault")
    userId := int32(56) // int32 | filter by user id (optional)
    projectId := int32(56) // int32 | filter by project id (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    showDeleted := true // bool | include deleted items (optional) (default to false)
    emoji := true // bool | parse emoji alias to unicode (optional) (default to true)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsDashboards := []string{"Inner_example"} // []string |  (optional)
    fieldsDashboardSettings := []string{"Inner_example"} // []string |  (optional)
    fieldsDashboardPanels := []string{"Inner_example"} // []string |  (optional)
    fieldsDashboardPanelSettings := []string{"Inner_example"} // []string |  (optional)
    dashboardIds := []int32{int32(123)} // []int32 | filter by user dashboard ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardsApi.GETProjectsApiV3DashboardsJson(context.Background()).UpdatedAfter(updatedAfter).OrderMode(orderMode).OrderBy(orderBy).UserId(userId).ProjectId(projectId).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).Emoji(emoji).Include(include).FieldsUsers(fieldsUsers).FieldsProjects(fieldsProjects).FieldsDashboards(fieldsDashboards).FieldsDashboardSettings(fieldsDashboardSettings).FieldsDashboardPanels(fieldsDashboardPanels).FieldsDashboardPanelSettings(fieldsDashboardPanelSettings).DashboardIds(dashboardIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.GETProjectsApiV3DashboardsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3DashboardsJson`: DashboardUserDashboardsResponse
    fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.GETProjectsApiV3DashboardsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3DashboardsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfter** | **string** | filter by updated after date | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;isDefault&quot;]
 **userId** | **int32** | filter by user id | 
 **projectId** | **int32** | filter by project id | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **showDeleted** | **bool** | include deleted items | [default to false]
 **emoji** | **bool** | parse emoji alias to unicode | [default to true]
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsDashboards** | **[]string** |  | 
 **fieldsDashboardSettings** | **[]string** |  | 
 **fieldsDashboardPanels** | **[]string** |  | 
 **fieldsDashboardPanelSettings** | **[]string** |  | 
 **dashboardIds** | **[]int32** | filter by user dashboard ids | 

### Return type

[**DashboardUserDashboardsResponse**](DashboardUserDashboardsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

