# TimerResponseIncluded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Projects** | Pointer to [**map[string]ViewProject**](ViewProject.md) |  | [optional] 
**Tasklists** | Pointer to [**map[string]ViewTasklist**](ViewTasklist.md) |  | [optional] 
**Tasks** | Pointer to [**map[string]ViewTask**](ViewTask.md) |  | [optional] 
**Users** | Pointer to [**map[string]ViewUser**](ViewUser.md) |  | [optional] 

## Methods

### NewTimerResponseIncluded

`func NewTimerResponseIncluded() *TimerResponseIncluded`

NewTimerResponseIncluded instantiates a new TimerResponseIncluded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimerResponseIncludedWithDefaults

`func NewTimerResponseIncludedWithDefaults() *TimerResponseIncluded`

NewTimerResponseIncludedWithDefaults instantiates a new TimerResponseIncluded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjects

`func (o *TimerResponseIncluded) GetProjects() map[string]ViewProject`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *TimerResponseIncluded) GetProjectsOk() (*map[string]ViewProject, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *TimerResponseIncluded) SetProjects(v map[string]ViewProject)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *TimerResponseIncluded) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetTasklists

`func (o *TimerResponseIncluded) GetTasklists() map[string]ViewTasklist`

GetTasklists returns the Tasklists field if non-nil, zero value otherwise.

### GetTasklistsOk

`func (o *TimerResponseIncluded) GetTasklistsOk() (*map[string]ViewTasklist, bool)`

GetTasklistsOk returns a tuple with the Tasklists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasklists

`func (o *TimerResponseIncluded) SetTasklists(v map[string]ViewTasklist)`

SetTasklists sets Tasklists field to given value.

### HasTasklists

`func (o *TimerResponseIncluded) HasTasklists() bool`

HasTasklists returns a boolean if a field has been set.

### GetTasks

`func (o *TimerResponseIncluded) GetTasks() map[string]ViewTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *TimerResponseIncluded) GetTasksOk() (*map[string]ViewTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *TimerResponseIncluded) SetTasks(v map[string]ViewTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *TimerResponseIncluded) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetUsers

`func (o *TimerResponseIncluded) GetUsers() map[string]ViewUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *TimerResponseIncluded) GetUsersOk() (*map[string]ViewUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *TimerResponseIncluded) SetUsers(v map[string]ViewUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *TimerResponseIncluded) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


