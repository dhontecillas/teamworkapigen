# PlannerWorkloadPlannersResponseIncluded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CalendarEvents** | Pointer to [**map[string]ViewCalendarEvent**](ViewCalendarEvent.md) |  | [optional] 
**Companies** | Pointer to [**map[string]ViewCompany**](ViewCompany.md) |  | [optional] 
**Milestones** | Pointer to [**map[string]ViewMilestone**](ViewMilestone.md) |  | [optional] 
**Tasklists** | Pointer to [**map[string]ViewTasklist**](ViewTasklist.md) |  | [optional] 
**Tasks** | Pointer to [**map[string]ViewTask**](ViewTask.md) |  | [optional] 
**Timelogs** | Pointer to [**map[string]ViewTimelog**](ViewTimelog.md) |  | [optional] 
**Users** | Pointer to [**map[string]ViewUser**](ViewUser.md) |  | [optional] 
**WorkingHourEntries** | Pointer to [**map[string]ViewWorkingHourEntry**](ViewWorkingHourEntry.md) |  | [optional] 
**WorkingHours** | Pointer to [**map[string]ViewWorkingHour**](ViewWorkingHour.md) |  | [optional] 

## Methods

### NewPlannerWorkloadPlannersResponseIncluded

`func NewPlannerWorkloadPlannersResponseIncluded() *PlannerWorkloadPlannersResponseIncluded`

NewPlannerWorkloadPlannersResponseIncluded instantiates a new PlannerWorkloadPlannersResponseIncluded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlannerWorkloadPlannersResponseIncludedWithDefaults

`func NewPlannerWorkloadPlannersResponseIncludedWithDefaults() *PlannerWorkloadPlannersResponseIncluded`

NewPlannerWorkloadPlannersResponseIncludedWithDefaults instantiates a new PlannerWorkloadPlannersResponseIncluded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalendarEvents

`func (o *PlannerWorkloadPlannersResponseIncluded) GetCalendarEvents() map[string]ViewCalendarEvent`

GetCalendarEvents returns the CalendarEvents field if non-nil, zero value otherwise.

### GetCalendarEventsOk

`func (o *PlannerWorkloadPlannersResponseIncluded) GetCalendarEventsOk() (*map[string]ViewCalendarEvent, bool)`

GetCalendarEventsOk returns a tuple with the CalendarEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarEvents

`func (o *PlannerWorkloadPlannersResponseIncluded) SetCalendarEvents(v map[string]ViewCalendarEvent)`

SetCalendarEvents sets CalendarEvents field to given value.

### HasCalendarEvents

`func (o *PlannerWorkloadPlannersResponseIncluded) HasCalendarEvents() bool`

HasCalendarEvents returns a boolean if a field has been set.

### GetCompanies

`func (o *PlannerWorkloadPlannersResponseIncluded) GetCompanies() map[string]ViewCompany`

GetCompanies returns the Companies field if non-nil, zero value otherwise.

### GetCompaniesOk

`func (o *PlannerWorkloadPlannersResponseIncluded) GetCompaniesOk() (*map[string]ViewCompany, bool)`

GetCompaniesOk returns a tuple with the Companies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanies

`func (o *PlannerWorkloadPlannersResponseIncluded) SetCompanies(v map[string]ViewCompany)`

SetCompanies sets Companies field to given value.

### HasCompanies

`func (o *PlannerWorkloadPlannersResponseIncluded) HasCompanies() bool`

HasCompanies returns a boolean if a field has been set.

### GetMilestones

`func (o *PlannerWorkloadPlannersResponseIncluded) GetMilestones() map[string]ViewMilestone`

GetMilestones returns the Milestones field if non-nil, zero value otherwise.

### GetMilestonesOk

`func (o *PlannerWorkloadPlannersResponseIncluded) GetMilestonesOk() (*map[string]ViewMilestone, bool)`

GetMilestonesOk returns a tuple with the Milestones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestones

`func (o *PlannerWorkloadPlannersResponseIncluded) SetMilestones(v map[string]ViewMilestone)`

SetMilestones sets Milestones field to given value.

### HasMilestones

`func (o *PlannerWorkloadPlannersResponseIncluded) HasMilestones() bool`

HasMilestones returns a boolean if a field has been set.

### GetTasklists

`func (o *PlannerWorkloadPlannersResponseIncluded) GetTasklists() map[string]ViewTasklist`

GetTasklists returns the Tasklists field if non-nil, zero value otherwise.

### GetTasklistsOk

`func (o *PlannerWorkloadPlannersResponseIncluded) GetTasklistsOk() (*map[string]ViewTasklist, bool)`

GetTasklistsOk returns a tuple with the Tasklists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasklists

`func (o *PlannerWorkloadPlannersResponseIncluded) SetTasklists(v map[string]ViewTasklist)`

SetTasklists sets Tasklists field to given value.

### HasTasklists

`func (o *PlannerWorkloadPlannersResponseIncluded) HasTasklists() bool`

HasTasklists returns a boolean if a field has been set.

### GetTasks

`func (o *PlannerWorkloadPlannersResponseIncluded) GetTasks() map[string]ViewTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *PlannerWorkloadPlannersResponseIncluded) GetTasksOk() (*map[string]ViewTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *PlannerWorkloadPlannersResponseIncluded) SetTasks(v map[string]ViewTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *PlannerWorkloadPlannersResponseIncluded) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetTimelogs

`func (o *PlannerWorkloadPlannersResponseIncluded) GetTimelogs() map[string]ViewTimelog`

GetTimelogs returns the Timelogs field if non-nil, zero value otherwise.

### GetTimelogsOk

`func (o *PlannerWorkloadPlannersResponseIncluded) GetTimelogsOk() (*map[string]ViewTimelog, bool)`

GetTimelogsOk returns a tuple with the Timelogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelogs

`func (o *PlannerWorkloadPlannersResponseIncluded) SetTimelogs(v map[string]ViewTimelog)`

SetTimelogs sets Timelogs field to given value.

### HasTimelogs

`func (o *PlannerWorkloadPlannersResponseIncluded) HasTimelogs() bool`

HasTimelogs returns a boolean if a field has been set.

### GetUsers

`func (o *PlannerWorkloadPlannersResponseIncluded) GetUsers() map[string]ViewUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *PlannerWorkloadPlannersResponseIncluded) GetUsersOk() (*map[string]ViewUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *PlannerWorkloadPlannersResponseIncluded) SetUsers(v map[string]ViewUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *PlannerWorkloadPlannersResponseIncluded) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetWorkingHourEntries

`func (o *PlannerWorkloadPlannersResponseIncluded) GetWorkingHourEntries() map[string]ViewWorkingHourEntry`

GetWorkingHourEntries returns the WorkingHourEntries field if non-nil, zero value otherwise.

### GetWorkingHourEntriesOk

`func (o *PlannerWorkloadPlannersResponseIncluded) GetWorkingHourEntriesOk() (*map[string]ViewWorkingHourEntry, bool)`

GetWorkingHourEntriesOk returns a tuple with the WorkingHourEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingHourEntries

`func (o *PlannerWorkloadPlannersResponseIncluded) SetWorkingHourEntries(v map[string]ViewWorkingHourEntry)`

SetWorkingHourEntries sets WorkingHourEntries field to given value.

### HasWorkingHourEntries

`func (o *PlannerWorkloadPlannersResponseIncluded) HasWorkingHourEntries() bool`

HasWorkingHourEntries returns a boolean if a field has been set.

### GetWorkingHours

`func (o *PlannerWorkloadPlannersResponseIncluded) GetWorkingHours() map[string]ViewWorkingHour`

GetWorkingHours returns the WorkingHours field if non-nil, zero value otherwise.

### GetWorkingHoursOk

`func (o *PlannerWorkloadPlannersResponseIncluded) GetWorkingHoursOk() (*map[string]ViewWorkingHour, bool)`

GetWorkingHoursOk returns a tuple with the WorkingHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingHours

`func (o *PlannerWorkloadPlannersResponseIncluded) SetWorkingHours(v map[string]ViewWorkingHour)`

SetWorkingHours sets WorkingHours field to given value.

### HasWorkingHours

`func (o *PlannerWorkloadPlannersResponseIncluded) HasWorkingHours() bool`

HasWorkingHours returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


