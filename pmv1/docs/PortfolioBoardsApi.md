# \PortfolioBoardsApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEPortfolioBoardsBoardIdJson**](PortfolioBoardsApi.md#DELETEPortfolioBoardsBoardIdJson) | **Delete** /portfolio/boards/{boardId}.json | Delete a board
[**DELETEPortfolioCardsCardIdJson**](PortfolioBoardsApi.md#DELETEPortfolioCardsCardIdJson) | **Delete** /portfolio/cards/{cardId}.json | Delete a Card in a Portfolio Column
[**DELETEPortfolioColumnsColumnIdJson**](PortfolioBoardsApi.md#DELETEPortfolioColumnsColumnIdJson) | **Delete** /portfolio/columns/{columnId}.json | Delete a Portfolio Column
[**GETPortfolioBoardsBoardIdColumnsJson**](PortfolioBoardsApi.md#GETPortfolioBoardsBoardIdColumnsJson) | **Get** /portfolio/boards/{boardId}/columns.json | Columns inside a Portfolio Board
[**GETPortfolioBoardsBoardIdJson**](PortfolioBoardsApi.md#GETPortfolioBoardsBoardIdJson) | **Get** /portfolio/boards/{boardId}.json | Get a specific Board
[**GETPortfolioBoardsJson**](PortfolioBoardsApi.md#GETPortfolioBoardsJson) | **Get** /portfolio/boards.json | Boards in Portfolio View
[**GETPortfolioCardsCardIdJson**](PortfolioBoardsApi.md#GETPortfolioCardsCardIdJson) | **Get** /portfolio/cards/{cardId}.json | Get a specific Card inside a Column
[**GETPortfolioColumnsColumnIdCardsJson**](PortfolioBoardsApi.md#GETPortfolioColumnsColumnIdCardsJson) | **Get** /portfolio/columns/{columnId}/cards.json | Get Cards inside a Portfolio Column
[**POSTPortfolioBoardsBoardIdColumnsJson**](PortfolioBoardsApi.md#POSTPortfolioBoardsBoardIdColumnsJson) | **Post** /portfolio/boards/{boardId}/columns.json | Add a column to the given Board
[**POSTPortfolioBoardsJson**](PortfolioBoardsApi.md#POSTPortfolioBoardsJson) | **Post** /portfolio/boards.json | Create a new Board
[**POSTPortfolioColumnsColumnIdCardsJson**](PortfolioBoardsApi.md#POSTPortfolioColumnsColumnIdCardsJson) | **Post** /portfolio/columns/{columnId}/cards.json | Adding a Project to a Column from Backlog
[**PUTPortfolioBoardsBoardIdJson**](PortfolioBoardsApi.md#PUTPortfolioBoardsBoardIdJson) | **Put** /portfolio/boards/{boardId}.json | Edit a Board
[**PUTPortfolioCardsCardIdMoveJson**](PortfolioBoardsApi.md#PUTPortfolioCardsCardIdMoveJson) | **Put** /portfolio/cards/{cardId}/move.json | Move Projects on a Portfolio Board
[**PUTPortfolioColumnsColumnIdJson**](PortfolioBoardsApi.md#PUTPortfolioColumnsColumnIdJson) | **Put** /portfolio/columns/{columnId}.json | Edit a Portfolio Column



## DELETEPortfolioBoardsBoardIdJson

> InlineResponse2001 DELETEPortfolioBoardsBoardIdJson(ctx, boardId).Execute()

Delete a board



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
    boardId := "boardId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortfolioBoardsApi.DELETEPortfolioBoardsBoardIdJson(context.Background(), boardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortfolioBoardsApi.DELETEPortfolioBoardsBoardIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DELETEPortfolioBoardsBoardIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `PortfolioBoardsApi.DELETEPortfolioBoardsBoardIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEPortfolioBoardsBoardIdJsonRequest struct via the builder pattern


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


## DELETEPortfolioCardsCardIdJson

> InlineResponse2001 DELETEPortfolioCardsCardIdJson(ctx, cardId).Execute()

Delete a Card in a Portfolio Column



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
    cardId := "cardId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortfolioBoardsApi.DELETEPortfolioCardsCardIdJson(context.Background(), cardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortfolioBoardsApi.DELETEPortfolioCardsCardIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DELETEPortfolioCardsCardIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `PortfolioBoardsApi.DELETEPortfolioCardsCardIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEPortfolioCardsCardIdJsonRequest struct via the builder pattern


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


## DELETEPortfolioColumnsColumnIdJson

> InlineResponse2001 DELETEPortfolioColumnsColumnIdJson(ctx, columnId).Execute()

Delete a Portfolio Column



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
    columnId := "columnId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortfolioBoardsApi.DELETEPortfolioColumnsColumnIdJson(context.Background(), columnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortfolioBoardsApi.DELETEPortfolioColumnsColumnIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DELETEPortfolioColumnsColumnIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `PortfolioBoardsApi.DELETEPortfolioColumnsColumnIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**columnId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEPortfolioColumnsColumnIdJsonRequest struct via the builder pattern


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


## GETPortfolioBoardsBoardIdColumnsJson

> InlineResponse20044 GETPortfolioBoardsBoardIdColumnsJson(ctx, boardId).Execute()

Columns inside a Portfolio Board



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
    boardId := "boardId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortfolioBoardsApi.GETPortfolioBoardsBoardIdColumnsJson(context.Background(), boardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortfolioBoardsApi.GETPortfolioBoardsBoardIdColumnsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPortfolioBoardsBoardIdColumnsJson`: InlineResponse20044
    fmt.Fprintf(os.Stdout, "Response from `PortfolioBoardsApi.GETPortfolioBoardsBoardIdColumnsJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPortfolioBoardsBoardIdColumnsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20044**](InlineResponse20044.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPortfolioBoardsBoardIdJson

> map[string]interface{} GETPortfolioBoardsBoardIdJson(ctx, boardId).Body(body).Execute()

Get a specific Board



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
    boardId := "boardId_example" // string | 
    body := *openapiclient.NewInlineObject42() // InlineObject42 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortfolioBoardsApi.GETPortfolioBoardsBoardIdJson(context.Background(), boardId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortfolioBoardsApi.GETPortfolioBoardsBoardIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPortfolioBoardsBoardIdJson`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PortfolioBoardsApi.GETPortfolioBoardsBoardIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPortfolioBoardsBoardIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject42**](InlineObject42.md) |  | 

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


## GETPortfolioBoardsJson

> InlineResponse20042 GETPortfolioBoardsJson(ctx).Execute()

Boards in Portfolio View



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
    resp, r, err := api_client.PortfolioBoardsApi.GETPortfolioBoardsJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortfolioBoardsApi.GETPortfolioBoardsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPortfolioBoardsJson`: InlineResponse20042
    fmt.Fprintf(os.Stdout, "Response from `PortfolioBoardsApi.GETPortfolioBoardsJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETPortfolioBoardsJsonRequest struct via the builder pattern


### Return type

[**InlineResponse20042**](InlineResponse20042.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPortfolioCardsCardIdJson

> InlineResponse20045 GETPortfolioCardsCardIdJson(ctx, cardId).Execute()

Get a specific Card inside a Column



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
    cardId := "cardId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortfolioBoardsApi.GETPortfolioCardsCardIdJson(context.Background(), cardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortfolioBoardsApi.GETPortfolioCardsCardIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPortfolioCardsCardIdJson`: InlineResponse20045
    fmt.Fprintf(os.Stdout, "Response from `PortfolioBoardsApi.GETPortfolioCardsCardIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPortfolioCardsCardIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20045**](InlineResponse20045.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPortfolioColumnsColumnIdCardsJson

> InlineResponse20046 GETPortfolioColumnsColumnIdCardsJson(ctx, columnId).Execute()

Get Cards inside a Portfolio Column



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
    columnId := "columnId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortfolioBoardsApi.GETPortfolioColumnsColumnIdCardsJson(context.Background(), columnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortfolioBoardsApi.GETPortfolioColumnsColumnIdCardsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPortfolioColumnsColumnIdCardsJson`: InlineResponse20046
    fmt.Fprintf(os.Stdout, "Response from `PortfolioBoardsApi.GETPortfolioColumnsColumnIdCardsJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**columnId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPortfolioColumnsColumnIdCardsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20046**](InlineResponse20046.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTPortfolioBoardsBoardIdColumnsJson

> InlineResponse201 POSTPortfolioBoardsBoardIdColumnsJson(ctx, boardId).Execute()

Add a column to the given Board



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
    boardId := "boardId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortfolioBoardsApi.POSTPortfolioBoardsBoardIdColumnsJson(context.Background(), boardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortfolioBoardsApi.POSTPortfolioBoardsBoardIdColumnsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTPortfolioBoardsBoardIdColumnsJson`: InlineResponse201
    fmt.Fprintf(os.Stdout, "Response from `PortfolioBoardsApi.POSTPortfolioBoardsBoardIdColumnsJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTPortfolioBoardsBoardIdColumnsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse201**](InlineResponse201.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTPortfolioBoardsJson

> InlineResponse20043 POSTPortfolioBoardsJson(ctx).Body(body).Execute()

Create a new Board



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
    body := *openapiclient.NewInlineObject41() // InlineObject41 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortfolioBoardsApi.POSTPortfolioBoardsJson(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortfolioBoardsApi.POSTPortfolioBoardsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTPortfolioBoardsJson`: InlineResponse20043
    fmt.Fprintf(os.Stdout, "Response from `PortfolioBoardsApi.POSTPortfolioBoardsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTPortfolioBoardsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InlineObject41**](InlineObject41.md) |  | 

### Return type

[**InlineResponse20043**](InlineResponse20043.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTPortfolioColumnsColumnIdCardsJson

> InlineResponse20047 POSTPortfolioColumnsColumnIdCardsJson(ctx, columnId).Body(body).Execute()

Adding a Project to a Column from Backlog



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
    columnId := "columnId_example" // string | 
    body := *openapiclient.NewInlineObject46() // InlineObject46 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortfolioBoardsApi.POSTPortfolioColumnsColumnIdCardsJson(context.Background(), columnId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortfolioBoardsApi.POSTPortfolioColumnsColumnIdCardsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTPortfolioColumnsColumnIdCardsJson`: InlineResponse20047
    fmt.Fprintf(os.Stdout, "Response from `PortfolioBoardsApi.POSTPortfolioColumnsColumnIdCardsJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**columnId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTPortfolioColumnsColumnIdCardsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject46**](InlineObject46.md) |  | 

### Return type

[**InlineResponse20047**](InlineResponse20047.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTPortfolioBoardsBoardIdJson

> InlineResponse2001 PUTPortfolioBoardsBoardIdJson(ctx, boardId).Body(body).Execute()

Edit a Board



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
    boardId := "boardId_example" // string | 
    body := *openapiclient.NewInlineObject43() // InlineObject43 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortfolioBoardsApi.PUTPortfolioBoardsBoardIdJson(context.Background(), boardId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortfolioBoardsApi.PUTPortfolioBoardsBoardIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTPortfolioBoardsBoardIdJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `PortfolioBoardsApi.PUTPortfolioBoardsBoardIdJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTPortfolioBoardsBoardIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject43**](InlineObject43.md) |  | 

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


## PUTPortfolioCardsCardIdMoveJson

> InlineResponse2001 PUTPortfolioCardsCardIdMoveJson(ctx, cardId).Body(body).Execute()

Move Projects on a Portfolio Board



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
    cardId := "cardId_example" // string | 
    body := *openapiclient.NewInlineObject44() // InlineObject44 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortfolioBoardsApi.PUTPortfolioCardsCardIdMoveJson(context.Background(), cardId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortfolioBoardsApi.PUTPortfolioCardsCardIdMoveJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTPortfolioCardsCardIdMoveJson`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `PortfolioBoardsApi.PUTPortfolioCardsCardIdMoveJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTPortfolioCardsCardIdMoveJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject44**](InlineObject44.md) |  | 

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


## PUTPortfolioColumnsColumnIdJson

> PUTPortfolioColumnsColumnIdJson(ctx, columnId).Body(body).Execute()

Edit a Portfolio Column



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
    columnId := "columnId_example" // string | 
    body := *openapiclient.NewInlineObject45() // InlineObject45 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortfolioBoardsApi.PUTPortfolioColumnsColumnIdJson(context.Background(), columnId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortfolioBoardsApi.PUTPortfolioColumnsColumnIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**columnId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPUTPortfolioColumnsColumnIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InlineObject45**](InlineObject45.md) |  | 

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

