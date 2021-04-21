# ScheduleProjectSchedulesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Included** | Pointer to [**ScheduleProjectSchedulesResponseIncluded**](ScheduleProjectSchedulesResponseIncluded.md) |  | [optional] 
**Meta** | Pointer to [**ViewMeta**](ViewMeta.md) |  | [optional] 
**ProjectSchedules** | Pointer to [**[]ViewProjectSchedule**](ViewProjectSchedule.md) |  | [optional] 

## Methods

### NewScheduleProjectSchedulesResponse

`func NewScheduleProjectSchedulesResponse() *ScheduleProjectSchedulesResponse`

NewScheduleProjectSchedulesResponse instantiates a new ScheduleProjectSchedulesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleProjectSchedulesResponseWithDefaults

`func NewScheduleProjectSchedulesResponseWithDefaults() *ScheduleProjectSchedulesResponse`

NewScheduleProjectSchedulesResponseWithDefaults instantiates a new ScheduleProjectSchedulesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncluded

`func (o *ScheduleProjectSchedulesResponse) GetIncluded() ScheduleProjectSchedulesResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *ScheduleProjectSchedulesResponse) GetIncludedOk() (*ScheduleProjectSchedulesResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *ScheduleProjectSchedulesResponse) SetIncluded(v ScheduleProjectSchedulesResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *ScheduleProjectSchedulesResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *ScheduleProjectSchedulesResponse) GetMeta() ViewMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ScheduleProjectSchedulesResponse) GetMetaOk() (*ViewMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ScheduleProjectSchedulesResponse) SetMeta(v ViewMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ScheduleProjectSchedulesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetProjectSchedules

`func (o *ScheduleProjectSchedulesResponse) GetProjectSchedules() []ViewProjectSchedule`

GetProjectSchedules returns the ProjectSchedules field if non-nil, zero value otherwise.

### GetProjectSchedulesOk

`func (o *ScheduleProjectSchedulesResponse) GetProjectSchedulesOk() (*[]ViewProjectSchedule, bool)`

GetProjectSchedulesOk returns a tuple with the ProjectSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectSchedules

`func (o *ScheduleProjectSchedulesResponse) SetProjectSchedules(v []ViewProjectSchedule)`

SetProjectSchedules sets ProjectSchedules field to given value.

### HasProjectSchedules

`func (o *ScheduleProjectSchedulesResponse) HasProjectSchedules() bool`

HasProjectSchedules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


