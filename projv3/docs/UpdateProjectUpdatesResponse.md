# UpdateProjectUpdatesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Included** | Pointer to [**CalendarEventsResponseIncluded**](CalendarEventsResponseIncluded.md) |  | [optional] 
**Meta** | Pointer to [**ViewMeta**](ViewMeta.md) |  | [optional] 
**ProjectUpdates** | Pointer to [**[]ViewProjectUpdate**](ViewProjectUpdate.md) |  | [optional] 

## Methods

### NewUpdateProjectUpdatesResponse

`func NewUpdateProjectUpdatesResponse() *UpdateProjectUpdatesResponse`

NewUpdateProjectUpdatesResponse instantiates a new UpdateProjectUpdatesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProjectUpdatesResponseWithDefaults

`func NewUpdateProjectUpdatesResponseWithDefaults() *UpdateProjectUpdatesResponse`

NewUpdateProjectUpdatesResponseWithDefaults instantiates a new UpdateProjectUpdatesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncluded

`func (o *UpdateProjectUpdatesResponse) GetIncluded() CalendarEventsResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *UpdateProjectUpdatesResponse) GetIncludedOk() (*CalendarEventsResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *UpdateProjectUpdatesResponse) SetIncluded(v CalendarEventsResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *UpdateProjectUpdatesResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *UpdateProjectUpdatesResponse) GetMeta() ViewMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *UpdateProjectUpdatesResponse) GetMetaOk() (*ViewMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *UpdateProjectUpdatesResponse) SetMeta(v ViewMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *UpdateProjectUpdatesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetProjectUpdates

`func (o *UpdateProjectUpdatesResponse) GetProjectUpdates() []ViewProjectUpdate`

GetProjectUpdates returns the ProjectUpdates field if non-nil, zero value otherwise.

### GetProjectUpdatesOk

`func (o *UpdateProjectUpdatesResponse) GetProjectUpdatesOk() (*[]ViewProjectUpdate, bool)`

GetProjectUpdatesOk returns a tuple with the ProjectUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectUpdates

`func (o *UpdateProjectUpdatesResponse) SetProjectUpdates(v []ViewProjectUpdate)`

SetProjectUpdates sets ProjectUpdates field to given value.

### HasProjectUpdates

`func (o *UpdateProjectUpdatesResponse) HasProjectUpdates() bool`

HasProjectUpdates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


