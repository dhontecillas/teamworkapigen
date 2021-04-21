# ActivityActivitiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activities** | Pointer to [**[]ActivityActivity**](ActivityActivity.md) |  | [optional] 
**Included** | Pointer to [**ActivityActivitiesResponseIncluded**](ActivityActivitiesResponseIncluded.md) |  | [optional] 
**Meta** | Pointer to [**ViewMeta**](ViewMeta.md) |  | [optional] 

## Methods

### NewActivityActivitiesResponse

`func NewActivityActivitiesResponse() *ActivityActivitiesResponse`

NewActivityActivitiesResponse instantiates a new ActivityActivitiesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityActivitiesResponseWithDefaults

`func NewActivityActivitiesResponseWithDefaults() *ActivityActivitiesResponse`

NewActivityActivitiesResponseWithDefaults instantiates a new ActivityActivitiesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivities

`func (o *ActivityActivitiesResponse) GetActivities() []ActivityActivity`

GetActivities returns the Activities field if non-nil, zero value otherwise.

### GetActivitiesOk

`func (o *ActivityActivitiesResponse) GetActivitiesOk() (*[]ActivityActivity, bool)`

GetActivitiesOk returns a tuple with the Activities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivities

`func (o *ActivityActivitiesResponse) SetActivities(v []ActivityActivity)`

SetActivities sets Activities field to given value.

### HasActivities

`func (o *ActivityActivitiesResponse) HasActivities() bool`

HasActivities returns a boolean if a field has been set.

### GetIncluded

`func (o *ActivityActivitiesResponse) GetIncluded() ActivityActivitiesResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *ActivityActivitiesResponse) GetIncludedOk() (*ActivityActivitiesResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *ActivityActivitiesResponse) SetIncluded(v ActivityActivitiesResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *ActivityActivitiesResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *ActivityActivitiesResponse) GetMeta() ViewMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ActivityActivitiesResponse) GetMetaOk() (*ViewMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ActivityActivitiesResponse) SetMeta(v ViewMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ActivityActivitiesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


