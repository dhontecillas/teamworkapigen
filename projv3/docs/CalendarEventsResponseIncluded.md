# CalendarEventsResponseIncluded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Projects** | Pointer to [**map[string]ViewProject**](ViewProject.md) |  | [optional] 
**Users** | Pointer to [**map[string]ViewUser**](ViewUser.md) |  | [optional] 

## Methods

### NewCalendarEventsResponseIncluded

`func NewCalendarEventsResponseIncluded() *CalendarEventsResponseIncluded`

NewCalendarEventsResponseIncluded instantiates a new CalendarEventsResponseIncluded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalendarEventsResponseIncludedWithDefaults

`func NewCalendarEventsResponseIncludedWithDefaults() *CalendarEventsResponseIncluded`

NewCalendarEventsResponseIncludedWithDefaults instantiates a new CalendarEventsResponseIncluded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjects

`func (o *CalendarEventsResponseIncluded) GetProjects() map[string]ViewProject`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *CalendarEventsResponseIncluded) GetProjectsOk() (*map[string]ViewProject, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *CalendarEventsResponseIncluded) SetProjects(v map[string]ViewProject)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *CalendarEventsResponseIncluded) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetUsers

`func (o *CalendarEventsResponseIncluded) GetUsers() map[string]ViewUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *CalendarEventsResponseIncluded) GetUsersOk() (*map[string]ViewUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *CalendarEventsResponseIncluded) SetUsers(v map[string]ViewUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *CalendarEventsResponseIncluded) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


