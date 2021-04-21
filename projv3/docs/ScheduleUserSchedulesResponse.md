# ScheduleUserSchedulesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Included** | Pointer to [**ScheduleUserSchedulesResponseIncluded**](ScheduleUserSchedulesResponseIncluded.md) |  | [optional] 
**Meta** | Pointer to [**ViewMeta**](ViewMeta.md) |  | [optional] 
**UserSchedules** | Pointer to [**[]ViewUserSchedule**](ViewUserSchedule.md) |  | [optional] 

## Methods

### NewScheduleUserSchedulesResponse

`func NewScheduleUserSchedulesResponse() *ScheduleUserSchedulesResponse`

NewScheduleUserSchedulesResponse instantiates a new ScheduleUserSchedulesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleUserSchedulesResponseWithDefaults

`func NewScheduleUserSchedulesResponseWithDefaults() *ScheduleUserSchedulesResponse`

NewScheduleUserSchedulesResponseWithDefaults instantiates a new ScheduleUserSchedulesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncluded

`func (o *ScheduleUserSchedulesResponse) GetIncluded() ScheduleUserSchedulesResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *ScheduleUserSchedulesResponse) GetIncludedOk() (*ScheduleUserSchedulesResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *ScheduleUserSchedulesResponse) SetIncluded(v ScheduleUserSchedulesResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *ScheduleUserSchedulesResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *ScheduleUserSchedulesResponse) GetMeta() ViewMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ScheduleUserSchedulesResponse) GetMetaOk() (*ViewMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ScheduleUserSchedulesResponse) SetMeta(v ViewMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ScheduleUserSchedulesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUserSchedules

`func (o *ScheduleUserSchedulesResponse) GetUserSchedules() []ViewUserSchedule`

GetUserSchedules returns the UserSchedules field if non-nil, zero value otherwise.

### GetUserSchedulesOk

`func (o *ScheduleUserSchedulesResponse) GetUserSchedulesOk() (*[]ViewUserSchedule, bool)`

GetUserSchedulesOk returns a tuple with the UserSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSchedules

`func (o *ScheduleUserSchedulesResponse) SetUserSchedules(v []ViewUserSchedule)`

SetUserSchedules sets UserSchedules field to given value.

### HasUserSchedules

`func (o *ScheduleUserSchedulesResponse) HasUserSchedules() bool`

HasUserSchedules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


