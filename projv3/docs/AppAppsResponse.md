# AppAppsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to [**[]ViewApp**](ViewApp.md) |  | [optional] 
**Included** | Pointer to [**AppAppsResponseIncluded**](AppAppsResponseIncluded.md) |  | [optional] 
**Meta** | Pointer to [**ViewMeta**](ViewMeta.md) |  | [optional] 

## Methods

### NewAppAppsResponse

`func NewAppAppsResponse() *AppAppsResponse`

NewAppAppsResponse instantiates a new AppAppsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppAppsResponseWithDefaults

`func NewAppAppsResponseWithDefaults() *AppAppsResponse`

NewAppAppsResponseWithDefaults instantiates a new AppAppsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *AppAppsResponse) GetApps() []ViewApp`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *AppAppsResponse) GetAppsOk() (*[]ViewApp, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *AppAppsResponse) SetApps(v []ViewApp)`

SetApps sets Apps field to given value.

### HasApps

`func (o *AppAppsResponse) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetIncluded

`func (o *AppAppsResponse) GetIncluded() AppAppsResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AppAppsResponse) GetIncludedOk() (*AppAppsResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AppAppsResponse) SetIncluded(v AppAppsResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AppAppsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *AppAppsResponse) GetMeta() ViewMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AppAppsResponse) GetMetaOk() (*ViewMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AppAppsResponse) SetMeta(v ViewMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AppAppsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


