# AppResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to [**ViewApp**](ViewApp.md) |  | [optional] 
**Included** | Pointer to [**AppAppsResponseIncluded**](AppAppsResponseIncluded.md) |  | [optional] 

## Methods

### NewAppResponse

`func NewAppResponse() *AppResponse`

NewAppResponse instantiates a new AppResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppResponseWithDefaults

`func NewAppResponseWithDefaults() *AppResponse`

NewAppResponseWithDefaults instantiates a new AppResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *AppResponse) GetApp() ViewApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *AppResponse) GetAppOk() (*ViewApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *AppResponse) SetApp(v ViewApp)`

SetApp sets App field to given value.

### HasApp

`func (o *AppResponse) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetIncluded

`func (o *AppResponse) GetIncluded() AppAppsResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AppResponse) GetIncludedOk() (*AppAppsResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AppResponse) SetIncluded(v AppAppsResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AppResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


