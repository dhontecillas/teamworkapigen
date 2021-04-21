# ScheduleProjectSchedulesResponseIncluded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allocations** | Pointer to [**map[string]ViewAllocation**](ViewAllocation.md) |  | [optional] 
**CalendarEvents** | Pointer to [**map[string]ViewCalendarEvent**](ViewCalendarEvent.md) |  | [optional] 
**Companies** | Pointer to [**map[string]ViewCompany**](ViewCompany.md) |  | [optional] 
**Milestones** | Pointer to [**map[string]ViewMilestone**](ViewMilestone.md) |  | [optional] 
**Projects** | Pointer to [**map[string]ViewProject**](ViewProject.md) |  | [optional] 
**Timelogs** | Pointer to [**map[string]ViewTimelog**](ViewTimelog.md) |  | [optional] 
**Users** | Pointer to [**map[string]ViewUser**](ViewUser.md) |  | [optional] 

## Methods

### NewScheduleProjectSchedulesResponseIncluded

`func NewScheduleProjectSchedulesResponseIncluded() *ScheduleProjectSchedulesResponseIncluded`

NewScheduleProjectSchedulesResponseIncluded instantiates a new ScheduleProjectSchedulesResponseIncluded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleProjectSchedulesResponseIncludedWithDefaults

`func NewScheduleProjectSchedulesResponseIncludedWithDefaults() *ScheduleProjectSchedulesResponseIncluded`

NewScheduleProjectSchedulesResponseIncludedWithDefaults instantiates a new ScheduleProjectSchedulesResponseIncluded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocations

`func (o *ScheduleProjectSchedulesResponseIncluded) GetAllocations() map[string]ViewAllocation`

GetAllocations returns the Allocations field if non-nil, zero value otherwise.

### GetAllocationsOk

`func (o *ScheduleProjectSchedulesResponseIncluded) GetAllocationsOk() (*map[string]ViewAllocation, bool)`

GetAllocationsOk returns a tuple with the Allocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocations

`func (o *ScheduleProjectSchedulesResponseIncluded) SetAllocations(v map[string]ViewAllocation)`

SetAllocations sets Allocations field to given value.

### HasAllocations

`func (o *ScheduleProjectSchedulesResponseIncluded) HasAllocations() bool`

HasAllocations returns a boolean if a field has been set.

### GetCalendarEvents

`func (o *ScheduleProjectSchedulesResponseIncluded) GetCalendarEvents() map[string]ViewCalendarEvent`

GetCalendarEvents returns the CalendarEvents field if non-nil, zero value otherwise.

### GetCalendarEventsOk

`func (o *ScheduleProjectSchedulesResponseIncluded) GetCalendarEventsOk() (*map[string]ViewCalendarEvent, bool)`

GetCalendarEventsOk returns a tuple with the CalendarEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarEvents

`func (o *ScheduleProjectSchedulesResponseIncluded) SetCalendarEvents(v map[string]ViewCalendarEvent)`

SetCalendarEvents sets CalendarEvents field to given value.

### HasCalendarEvents

`func (o *ScheduleProjectSchedulesResponseIncluded) HasCalendarEvents() bool`

HasCalendarEvents returns a boolean if a field has been set.

### GetCompanies

`func (o *ScheduleProjectSchedulesResponseIncluded) GetCompanies() map[string]ViewCompany`

GetCompanies returns the Companies field if non-nil, zero value otherwise.

### GetCompaniesOk

`func (o *ScheduleProjectSchedulesResponseIncluded) GetCompaniesOk() (*map[string]ViewCompany, bool)`

GetCompaniesOk returns a tuple with the Companies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanies

`func (o *ScheduleProjectSchedulesResponseIncluded) SetCompanies(v map[string]ViewCompany)`

SetCompanies sets Companies field to given value.

### HasCompanies

`func (o *ScheduleProjectSchedulesResponseIncluded) HasCompanies() bool`

HasCompanies returns a boolean if a field has been set.

### GetMilestones

`func (o *ScheduleProjectSchedulesResponseIncluded) GetMilestones() map[string]ViewMilestone`

GetMilestones returns the Milestones field if non-nil, zero value otherwise.

### GetMilestonesOk

`func (o *ScheduleProjectSchedulesResponseIncluded) GetMilestonesOk() (*map[string]ViewMilestone, bool)`

GetMilestonesOk returns a tuple with the Milestones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestones

`func (o *ScheduleProjectSchedulesResponseIncluded) SetMilestones(v map[string]ViewMilestone)`

SetMilestones sets Milestones field to given value.

### HasMilestones

`func (o *ScheduleProjectSchedulesResponseIncluded) HasMilestones() bool`

HasMilestones returns a boolean if a field has been set.

### GetProjects

`func (o *ScheduleProjectSchedulesResponseIncluded) GetProjects() map[string]ViewProject`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *ScheduleProjectSchedulesResponseIncluded) GetProjectsOk() (*map[string]ViewProject, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *ScheduleProjectSchedulesResponseIncluded) SetProjects(v map[string]ViewProject)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *ScheduleProjectSchedulesResponseIncluded) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetTimelogs

`func (o *ScheduleProjectSchedulesResponseIncluded) GetTimelogs() map[string]ViewTimelog`

GetTimelogs returns the Timelogs field if non-nil, zero value otherwise.

### GetTimelogsOk

`func (o *ScheduleProjectSchedulesResponseIncluded) GetTimelogsOk() (*map[string]ViewTimelog, bool)`

GetTimelogsOk returns a tuple with the Timelogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelogs

`func (o *ScheduleProjectSchedulesResponseIncluded) SetTimelogs(v map[string]ViewTimelog)`

SetTimelogs sets Timelogs field to given value.

### HasTimelogs

`func (o *ScheduleProjectSchedulesResponseIncluded) HasTimelogs() bool`

HasTimelogs returns a boolean if a field has been set.

### GetUsers

`func (o *ScheduleProjectSchedulesResponseIncluded) GetUsers() map[string]ViewUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ScheduleProjectSchedulesResponseIncluded) GetUsersOk() (*map[string]ViewUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ScheduleProjectSchedulesResponseIncluded) SetUsers(v map[string]ViewUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *ScheduleProjectSchedulesResponseIncluded) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


