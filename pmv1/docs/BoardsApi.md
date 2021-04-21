# \BoardsApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEBoardsColumnsCardsIdJson**](BoardsApi.md#DELETEBoardsColumnsCardsIdJson) | **Delete** /boards/columns/cards/{id}.json | Remove a Card
[**DELETEBoardsColumnsIdJson**](BoardsApi.md#DELETEBoardsColumnsIdJson) | **Delete** /boards/columns/{id}.json | Delete a Column
[**GETBoardsColumnsIdCardsJson**](BoardsApi.md#GETBoardsColumnsIdCardsJson) | **Get** /boards/columns/{id}/cards.json | List Cards in a Column
[**GETProjectsIdBoardsColumnsJson**](BoardsApi.md#GETProjectsIdBoardsColumnsJson) | **Get** /projects/{id}/boards/columns.json | List Columns
[**POSTBoardsColumnsIdCardsJson**](BoardsApi.md#POSTBoardsColumnsIdCardsJson) | **Post** /boards/columns/{id}/cards.json | Add a Task from the Backlog to a Column
[**POSTProjectsIdBoardsColumnsJson**](BoardsApi.md#POSTProjectsIdBoardsColumnsJson) | **Post** /projects/{id}/boards/columns.json | Create a new Column
[**PUTBoardsColumnsCardsIdJson**](BoardsApi.md#PUTBoardsColumnsCardsIdJson) | **Put** /boards/columns/cards/{id}.json | Edit a Card
[**PUTBoardsColumnsCardsIdMoveJson**](BoardsApi.md#PUTBoardsColumnsCardsIdMoveJson) | **Put** /boards/columns/cards/{id}/move.json | Move a Card



## DELETEBoardsColumnsCardsIdJson

> InlineResponse2001 DELETEBoardsColumnsCardsIdJson(ctx, id).Execute()

Remove a Card



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
    resp, r, err := api_client.BoardsApi.DELETEBoardsColumnsCardsIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BoardsApi.DELETEBoardsColumnsCardsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DELETEBoardsColumnsCardsIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `BoardsApi.DELETEBoardsColumnsCardsIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEBoardsColumnsCardsIdJsonRequest struct via the builder pattern


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


## DELETEBoardsColumnsIdJson

> InlineResponse2001 DELETEBoardsColumnsIdJson(ctx, id).Execute()

Delete a Column



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
    resp, r, err := api_client.BoardsApi.DELETEBoardsColumnsIdJson(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BoardsApi.DELETEBoardsColumnsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DELETEBoardsColumnsIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `BoardsApi.DELETEBoardsColumnsIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEBoardsColumnsIdJsonRequest struct via the builder pattern


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


## GETBoardsColumnsIdCardsJson

> InlineResponse2002 GETBoardsColumnsIdCardsJson(ctx, id).Page(page).PageSize(pageSize).ShowDeleted(showDeleted).DeletedAfterDate(deletedAfterDate).UpdatedAfterDate(updatedAfterDate).SearchTerm(searchTerm).Execute()

List Cards in a Column



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
    page := "page_example" // string | The page number to get (optional)
    pageSize := int32(56) // int32 | The number of records to return per page (optional)
    showDeleted := true // bool | Show deleted Columns in the API response or not (optional)
    deletedAfterDate := "deletedAfterDate_example" // string | deletedAfterDate (optional)
    updatedAfterDate := "updatedAfterDate_example" // string | Show Columns updated after a given date (optional)
    searchTerm := "searchTerm_example" // string | An optional term to filter down to Cards containing this text in the Cards name (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BoardsApi.GETBoardsColumnsIdCardsJson(context.Background(), id).Page(page).PageSize(pageSize).ShowDeleted(showDeleted).DeletedAfterDate(deletedAfterDate).UpdatedAfterDate(updatedAfterDate).SearchTerm(searchTerm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BoardsApi.GETBoardsColumnsIdCardsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETBoardsColumnsIdCardsJson`: InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `BoardsApi.GETBoardsColumnsIdCardsJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETBoardsColumnsIdCardsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | The page number to get | 
 **pageSize** | **int32** | The number of records to return per page | 
 **showDeleted** | **bool** | Show deleted Columns in the API response or not | 
 **deletedAfterDate** | **string** | deletedAfterDate | 
 **updatedAfterDate** | **string** | Show Columns updated after a given date | 
 **searchTerm** | **string** | An optional term to filter down to Cards containing this text in the Cards name | 

### Return type

[**InlineResponse2002**](InlineResponse2002.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsIdBoardsColumnsJson

> InlineResponse20061 GETProjectsIdBoardsColumnsJson(ctx, id).Page(page).PageSize(pageSize).ShowDeleted(showDeleted).DeletedAfterDate(deletedAfterDate).UpdatedAfterDate(updatedAfterDate).Execute()

List Columns



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
    page := int32(56) // int32 | The page number to get (optional)
    pageSize := int32(56) // int32 | The number of records to return per page (optional)
    showDeleted := true // bool | Show deleted Columns in the API response or not (optional)
    deletedAfterDate := "deletedAfterDate_example" // string | Show Columns deleted after a given date (optional)
    updatedAfterDate := "updatedAfterDate_example" // string | Show Columns updated after a given date (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BoardsApi.GETProjectsIdBoardsColumnsJson(context.Background(), id).Page(page).PageSize(pageSize).ShowDeleted(showDeleted).DeletedAfterDate(deletedAfterDate).UpdatedAfterDate(updatedAfterDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BoardsApi.GETProjectsIdBoardsColumnsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsIdBoardsColumnsJson`: InlineResponse20061
    fmt.Fprintf(os.Stdout, "Response from `BoardsApi.GETProjectsIdBoardsColumnsJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsIdBoardsColumnsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page number to get | 
 **pageSize** | **int32** | The number of records to return per page | 
 **showDeleted** | **bool** | Show deleted Columns in the API response or not | 
 **deletedAfterDate** | **string** | Show Columns deleted after a given date | 
 **updatedAfterDate** | **string** | Show Columns updated after a given date | 

### Return type

[**InlineResponse20061**](InlineResponse20061.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTBoardsColumnsIdCardsJson

> POSTBoardsColumnsIdCardsJson(ctx, id).Body(body).Execute()

Add a Task from the Backlog to a Column



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
    body := *openapiclient.NewInlineObject3() // InlineObject3 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BoardsApi.POSTBoardsColumnsIdCardsJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BoardsApi.POSTBoardsColumnsIdCardsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTBoardsColumnsIdCardsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject3**](InlineObject3.md) |  | 

### Return type

 (empty response body)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsIdBoardsColumnsJson

> InlineResponse201 POSTProjectsIdBoardsColumnsJson(ctx, id).Body(body).Execute()

Create a new Column



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
    body := *openapiclient.NewInlineObject56() // InlineObject56 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BoardsApi.POSTProjectsIdBoardsColumnsJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BoardsApi.POSTProjectsIdBoardsColumnsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsIdBoardsColumnsJson`: InlineResponse201
    fmt.Fprintf(os.Stdout, "Response from `BoardsApi.POSTProjectsIdBoardsColumnsJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsIdBoardsColumnsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject56**](InlineObject56.md) |  | 

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


## PUTBoardsColumnsCardsIdJson

> InlineResponse2001 PUTBoardsColumnsCardsIdJson(ctx, id).Body(body).Execute()

Edit a Card



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
    body := *openapiclient.NewInlineObject1() // InlineObject1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BoardsApi.PUTBoardsColumnsCardsIdJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BoardsApi.PUTBoardsColumnsCardsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTBoardsColumnsCardsIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `BoardsApi.PUTBoardsColumnsCardsIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTBoardsColumnsCardsIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject1**](InlineObject1.md) |  | 

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


## PUTBoardsColumnsCardsIdMoveJson

> map[string]interface{} PUTBoardsColumnsCardsIdMoveJson(ctx, id).Body(body).Execute()

Move a Card



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
    body := *openapiclient.NewInlineObject2() // InlineObject2 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BoardsApi.PUTBoardsColumnsCardsIdMoveJson(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BoardsApi.PUTBoardsColumnsCardsIdMoveJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTBoardsColumnsCardsIdMoveJson`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BoardsApi.PUTBoardsColumnsCardsIdMoveJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTBoardsColumnsCardsIdMoveJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject2**](InlineObject2.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

