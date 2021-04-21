# \CalendarEventApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETECalendareventsIdJson**](CalendarEventApi.md#DELETECalendareventsIdJson) | **Delete** /calendarevents/{id}.json | Delete Event
[**DELETEEventtypesIdJson**](CalendarEventApi.md#DELETEEventtypesIdJson) | **Delete** /eventtypes/{id}.json | Delete Event Type
[**GETCalendareventsIdJson**](CalendarEventApi.md#GETCalendareventsIdJson) | **Get** /calendarevents/{id}.json | Get an Event
[**GETCalendareventsJson**](CalendarEventApi.md#GETCalendareventsJson) | **Get** /calendarevents.json | Get Events
[**GETCalendareventtypesJson**](CalendarEventApi.md#GETCalendareventtypesJson) | **Get** /calendareventtypes.json | Get Event Types
[**POSTCalendareventsJson**](CalendarEventApi.md#POSTCalendareventsJson) | **Post** /calendarevents.json | Create a Recurring Event
[**POSTEventtypesJson**](CalendarEventApi.md#POSTEventtypesJson) | **Post** /eventtypes.json | Create an Event Type
[**PUTCalendareventsIdJson**](CalendarEventApi.md#PUTCalendareventsIdJson) | **Put** /calendarevents/{id}.json | Edit an Event
[**PUTEventtypesIdJson**](CalendarEventApi.md#PUTEventtypesIdJson) | **Put** /eventtypes/{id}.json | Edit an Event Type



## DELETECalendareventsIdJson

> InlineResponse2001 DELETECalendareventsIdJson(ctx, id).Body(body).Execute()

Delete Event



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
    id := "id_example" // string | 
    body := *openapiclient.NewInlineObject6() // InlineObject6 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CalendarEventApi.DELETECalendareventsIdJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalendarEventApi.DELETECalendareventsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DELETECalendareventsIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `CalendarEventApi.DELETECalendareventsIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETECalendareventsIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject6**](InlineObject6.md) |  | 

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


## DELETEEventtypesIdJson

> InlineResponse2001 DELETEEventtypesIdJson(ctx, id).Execute()

Delete Event Type



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CalendarEventApi.DELETEEventtypesIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalendarEventApi.DELETEEventtypesIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DELETEEventtypesIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `CalendarEventApi.DELETEEventtypesIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEEventtypesIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETCalendareventsIdJson

> InlineResponse2005 GETCalendareventsIdJson(ctx, id).Execute()

Get an Event



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CalendarEventApi.GETCalendareventsIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalendarEventApi.GETCalendareventsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCalendareventsIdJson`: InlineResponse2005
    fmt.Fprintf(os.Stdout, "Response from `CalendarEventApi.GETCalendareventsIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCalendareventsIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2005**](InlineResponse2005.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETCalendareventsJson

> InlineResponse2003 GETCalendareventsJson(ctx).Startdate(startdate).EndDate(endDate).ShowDeleted(showDeleted).UpdatedAfterDate(updatedAfterDate).EventTypeId(eventTypeId).Page(page).UserId(userId).AttendingOnly(attendingOnly).Execute()

Get Events



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
    startdate := "startdate_example" // string | YYYYMMDD
    endDate := "endDate_example" // string | YYYYMMDD (default to "tomorrow's date")
    showDeleted := true // bool | Whether to include deleted events (optional)
    updatedAfterDate := "updatedAfterDate_example" // string | Only return events updated after a certain datetime (YYYYMMDDHHMMSS) (optional)
    eventTypeId := int32(56) // int32 | Only return events with a given eventTypeId (Numeric only - default 0) (optional)
    page := int32(56) // int32 |  (optional)
    userId := float32(8.14) // float32 | If you want to bring back records which this user is attending or has created. (optional)
    attendingOnly := true // bool | If set to true, along with passing in the userId, it will return events that user is attending only. Not ones they have also created.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CalendarEventApi.GETCalendareventsJson(context.Background()).Startdate(startdate).EndDate(endDate).ShowDeleted(showDeleted).UpdatedAfterDate(updatedAfterDate).EventTypeId(eventTypeId).Page(page).UserId(userId).AttendingOnly(attendingOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalendarEventApi.GETCalendareventsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCalendareventsJson`: InlineResponse2003
    fmt.Fprintf(os.Stdout, "Response from `CalendarEventApi.GETCalendareventsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETCalendareventsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startdate** | **string** | YYYYMMDD | 
 **endDate** | **string** | YYYYMMDD | [default to &quot;tomorrow&#39;s date&quot;]
 **showDeleted** | **bool** | Whether to include deleted events | 
 **updatedAfterDate** | **string** | Only return events updated after a certain datetime (YYYYMMDDHHMMSS) | 
 **eventTypeId** | **int32** | Only return events with a given eventTypeId (Numeric only - default 0) | 
 **page** | **int32** |  | 
 **userId** | **float32** | If you want to bring back records which this user is attending or has created. | 
 **attendingOnly** | **bool** | If set to true, along with passing in the userId, it will return events that user is attending only. Not ones they have also created.  | 

### Return type

[**InlineResponse2003**](InlineResponse2003.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETCalendareventtypesJson

> InlineResponse2006 GETCalendareventtypesJson(ctx).Execute()

Get Event Types



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
    resp, r, err := api_client.CalendarEventApi.GETCalendareventtypesJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalendarEventApi.GETCalendareventtypesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCalendareventtypesJson`: InlineResponse2006
    fmt.Fprintf(os.Stdout, "Response from `CalendarEventApi.GETCalendareventtypesJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETCalendareventtypesJsonRequest struct via the builder pattern


### Return type

[**InlineResponse2006**](InlineResponse2006.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTCalendareventsJson

> InlineResponse201 POSTCalendareventsJson(ctx).Body(body).Execute()

Create a Recurring Event



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
    body := *openapiclient.NewInlineObject4() // InlineObject4 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CalendarEventApi.POSTCalendareventsJson(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalendarEventApi.POSTCalendareventsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTCalendareventsJson`: InlineResponse201
    fmt.Fprintf(os.Stdout, "Response from `CalendarEventApi.POSTCalendareventsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTCalendareventsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InlineObject4**](InlineObject4.md) |  | 

### Return type

[**InlineResponse201**](InlineResponse201.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTEventtypesJson

> InlineResponse201 POSTEventtypesJson(ctx).Body(body).Execute()

Create an Event Type



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
    body := *openapiclient.NewInlineObject12() // InlineObject12 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CalendarEventApi.POSTEventtypesJson(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalendarEventApi.POSTEventtypesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTEventtypesJson`: InlineResponse201
    fmt.Fprintf(os.Stdout, "Response from `CalendarEventApi.POSTEventtypesJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTEventtypesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InlineObject12**](InlineObject12.md) |  | 

### Return type

[**InlineResponse201**](InlineResponse201.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTCalendareventsIdJson

> InlineResponse2001 PUTCalendareventsIdJson(ctx, id).Body(body).Execute()

Edit an Event



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
    id := "id_example" // string | 
    body := *openapiclient.NewInlineObject5() // InlineObject5 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CalendarEventApi.PUTCalendareventsIdJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalendarEventApi.PUTCalendareventsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTCalendareventsIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `CalendarEventApi.PUTCalendareventsIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTCalendareventsIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject5**](InlineObject5.md) |  | 

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


## PUTEventtypesIdJson

> InlineResponse2001 PUTEventtypesIdJson(ctx, id).Body(body).Execute()

Edit an Event Type



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
    id := "id_example" // string | 
    body := *openapiclient.NewInlineObject13() // InlineObject13 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CalendarEventApi.PUTEventtypesIdJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalendarEventApi.PUTEventtypesIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTEventtypesIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `CalendarEventApi.PUTEventtypesIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTEventtypesIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject13**](InlineObject13.md) |  | 

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

