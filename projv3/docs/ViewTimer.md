# ViewTimer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Billable** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Intervals** | Pointer to [**[]ViewTimerInterval**](ViewTimerInterval.md) |  | [optional] 
**LastStartedAt** | Pointer to **string** |  | [optional] 
**Project** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**Running** | Pointer to **bool** |  | [optional] 
**ServerTime** | Pointer to **string** |  | [optional] 
**Task** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**TaskId** | Pointer to **int32** |  | [optional] 
**TimeLogId** | Pointer to **int32** |  | [optional] 
**Timelog** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**TimerLastIntervalEnd** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 

## Methods

### NewViewTimer

`func NewViewTimer() *ViewTimer`

NewViewTimer instantiates a new ViewTimer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewTimerWithDefaults

`func NewViewTimerWithDefaults() *ViewTimer`

NewViewTimerWithDefaults instantiates a new ViewTimer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillable

`func (o *ViewTimer) GetBillable() bool`

GetBillable returns the Billable field if non-nil, zero value otherwise.

### GetBillableOk

`func (o *ViewTimer) GetBillableOk() (*bool, bool)`

GetBillableOk returns a tuple with the Billable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillable

`func (o *ViewTimer) SetBillable(v bool)`

SetBillable sets Billable field to given value.

### HasBillable

`func (o *ViewTimer) HasBillable() bool`

HasBillable returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ViewTimer) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ViewTimer) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ViewTimer) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ViewTimer) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeleted

`func (o *ViewTimer) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *ViewTimer) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *ViewTimer) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *ViewTimer) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ViewTimer) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ViewTimer) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ViewTimer) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ViewTimer) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDescription

`func (o *ViewTimer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ViewTimer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ViewTimer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ViewTimer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDuration

`func (o *ViewTimer) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ViewTimer) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ViewTimer) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ViewTimer) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetId

`func (o *ViewTimer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewTimer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewTimer) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewTimer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIntervals

`func (o *ViewTimer) GetIntervals() []ViewTimerInterval`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *ViewTimer) GetIntervalsOk() (*[]ViewTimerInterval, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *ViewTimer) SetIntervals(v []ViewTimerInterval)`

SetIntervals sets Intervals field to given value.

### HasIntervals

`func (o *ViewTimer) HasIntervals() bool`

HasIntervals returns a boolean if a field has been set.

### GetLastStartedAt

`func (o *ViewTimer) GetLastStartedAt() string`

GetLastStartedAt returns the LastStartedAt field if non-nil, zero value otherwise.

### GetLastStartedAtOk

`func (o *ViewTimer) GetLastStartedAtOk() (*string, bool)`

GetLastStartedAtOk returns a tuple with the LastStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStartedAt

`func (o *ViewTimer) SetLastStartedAt(v string)`

SetLastStartedAt sets LastStartedAt field to given value.

### HasLastStartedAt

`func (o *ViewTimer) HasLastStartedAt() bool`

HasLastStartedAt returns a boolean if a field has been set.

### GetProject

`func (o *ViewTimer) GetProject() ViewRelationship`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ViewTimer) GetProjectOk() (*ViewRelationship, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ViewTimer) SetProject(v ViewRelationship)`

SetProject sets Project field to given value.

### HasProject

`func (o *ViewTimer) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectId

`func (o *ViewTimer) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ViewTimer) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ViewTimer) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ViewTimer) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetRunning

`func (o *ViewTimer) GetRunning() bool`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *ViewTimer) GetRunningOk() (*bool, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *ViewTimer) SetRunning(v bool)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *ViewTimer) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetServerTime

`func (o *ViewTimer) GetServerTime() string`

GetServerTime returns the ServerTime field if non-nil, zero value otherwise.

### GetServerTimeOk

`func (o *ViewTimer) GetServerTimeOk() (*string, bool)`

GetServerTimeOk returns a tuple with the ServerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTime

`func (o *ViewTimer) SetServerTime(v string)`

SetServerTime sets ServerTime field to given value.

### HasServerTime

`func (o *ViewTimer) HasServerTime() bool`

HasServerTime returns a boolean if a field has been set.

### GetTask

`func (o *ViewTimer) GetTask() ViewRelationship`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *ViewTimer) GetTaskOk() (*ViewRelationship, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *ViewTimer) SetTask(v ViewRelationship)`

SetTask sets Task field to given value.

### HasTask

`func (o *ViewTimer) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetTaskId

`func (o *ViewTimer) GetTaskId() int32`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *ViewTimer) GetTaskIdOk() (*int32, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *ViewTimer) SetTaskId(v int32)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *ViewTimer) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetTimeLogId

`func (o *ViewTimer) GetTimeLogId() int32`

GetTimeLogId returns the TimeLogId field if non-nil, zero value otherwise.

### GetTimeLogIdOk

`func (o *ViewTimer) GetTimeLogIdOk() (*int32, bool)`

GetTimeLogIdOk returns a tuple with the TimeLogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLogId

`func (o *ViewTimer) SetTimeLogId(v int32)`

SetTimeLogId sets TimeLogId field to given value.

### HasTimeLogId

`func (o *ViewTimer) HasTimeLogId() bool`

HasTimeLogId returns a boolean if a field has been set.

### GetTimelog

`func (o *ViewTimer) GetTimelog() ViewRelationship`

GetTimelog returns the Timelog field if non-nil, zero value otherwise.

### GetTimelogOk

`func (o *ViewTimer) GetTimelogOk() (*ViewRelationship, bool)`

GetTimelogOk returns a tuple with the Timelog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelog

`func (o *ViewTimer) SetTimelog(v ViewRelationship)`

SetTimelog sets Timelog field to given value.

### HasTimelog

`func (o *ViewTimer) HasTimelog() bool`

HasTimelog returns a boolean if a field has been set.

### GetTimerLastIntervalEnd

`func (o *ViewTimer) GetTimerLastIntervalEnd() string`

GetTimerLastIntervalEnd returns the TimerLastIntervalEnd field if non-nil, zero value otherwise.

### GetTimerLastIntervalEndOk

`func (o *ViewTimer) GetTimerLastIntervalEndOk() (*string, bool)`

GetTimerLastIntervalEndOk returns a tuple with the TimerLastIntervalEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimerLastIntervalEnd

`func (o *ViewTimer) SetTimerLastIntervalEnd(v string)`

SetTimerLastIntervalEnd sets TimerLastIntervalEnd field to given value.

### HasTimerLastIntervalEnd

`func (o *ViewTimer) HasTimerLastIntervalEnd() bool`

HasTimerLastIntervalEnd returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ViewTimer) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ViewTimer) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ViewTimer) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ViewTimer) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUser

`func (o *ViewTimer) GetUser() ViewRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ViewTimer) GetUserOk() (*ViewRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ViewTimer) SetUser(v ViewRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *ViewTimer) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserId

`func (o *ViewTimer) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ViewTimer) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ViewTimer) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ViewTimer) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


