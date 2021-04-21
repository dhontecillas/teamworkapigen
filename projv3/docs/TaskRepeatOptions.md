# TaskRepeatOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **int32** |  | [optional] 
**EndsAt** | Pointer to **map[string]interface{}** | Date unmarshals represents a Unified API Spec date format. | [optional] 
**Frequency** | Pointer to **string** |  | [optional] 
**MonthlyRepeatType** | Pointer to **string** |  | [optional] 
**SelectedDays** | Pointer to **string** |  | [optional] 

## Methods

### NewTaskRepeatOptions

`func NewTaskRepeatOptions() *TaskRepeatOptions`

NewTaskRepeatOptions instantiates a new TaskRepeatOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskRepeatOptionsWithDefaults

`func NewTaskRepeatOptionsWithDefaults() *TaskRepeatOptions`

NewTaskRepeatOptionsWithDefaults instantiates a new TaskRepeatOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *TaskRepeatOptions) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TaskRepeatOptions) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TaskRepeatOptions) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *TaskRepeatOptions) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEndsAt

`func (o *TaskRepeatOptions) GetEndsAt() map[string]interface{}`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *TaskRepeatOptions) GetEndsAtOk() (*map[string]interface{}, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *TaskRepeatOptions) SetEndsAt(v map[string]interface{})`

SetEndsAt sets EndsAt field to given value.

### HasEndsAt

`func (o *TaskRepeatOptions) HasEndsAt() bool`

HasEndsAt returns a boolean if a field has been set.

### GetFrequency

`func (o *TaskRepeatOptions) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *TaskRepeatOptions) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *TaskRepeatOptions) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *TaskRepeatOptions) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetMonthlyRepeatType

`func (o *TaskRepeatOptions) GetMonthlyRepeatType() string`

GetMonthlyRepeatType returns the MonthlyRepeatType field if non-nil, zero value otherwise.

### GetMonthlyRepeatTypeOk

`func (o *TaskRepeatOptions) GetMonthlyRepeatTypeOk() (*string, bool)`

GetMonthlyRepeatTypeOk returns a tuple with the MonthlyRepeatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyRepeatType

`func (o *TaskRepeatOptions) SetMonthlyRepeatType(v string)`

SetMonthlyRepeatType sets MonthlyRepeatType field to given value.

### HasMonthlyRepeatType

`func (o *TaskRepeatOptions) HasMonthlyRepeatType() bool`

HasMonthlyRepeatType returns a boolean if a field has been set.

### GetSelectedDays

`func (o *TaskRepeatOptions) GetSelectedDays() string`

GetSelectedDays returns the SelectedDays field if non-nil, zero value otherwise.

### GetSelectedDaysOk

`func (o *TaskRepeatOptions) GetSelectedDaysOk() (*string, bool)`

GetSelectedDaysOk returns a tuple with the SelectedDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedDays

`func (o *TaskRepeatOptions) SetSelectedDays(v string)`

SetSelectedDays sets SelectedDays field to given value.

### HasSelectedDays

`func (o *TaskRepeatOptions) HasSelectedDays() bool`

HasSelectedDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


