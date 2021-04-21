# CalendareventsIdJsonEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllDay** | Pointer to **string** |  | [optional] 
**AttendeesCanEdit** | Pointer to **string** |  | [optional] 
**AttendingUserIds** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**End** | Pointer to **string** |  | [optional] 
**Notify** | Pointer to **string** |  | [optional] 
**NotifyUserIds** | Pointer to **string** |  | [optional] 
**Privacy** | Pointer to [**InlineResponse2003Privacy**](InlineResponse2003Privacy.md) |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**ProjectUsersCanEdit** | Pointer to **string** |  | [optional] 
**Reminders** | Pointer to [**[]CalendareventsIdJsonEventReminders**](CalendareventsIdJsonEventReminders.md) |  | [optional] 
**ShowAsBusy** | Pointer to **string** |  | [optional] 
**Start** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**CalendareventsIdJsonEventType**](CalendareventsIdJsonEventType.md) |  | [optional] 
**Utc** | Pointer to **bool** |  | [optional] 
**Where** | Pointer to **string** |  | [optional] 

## Methods

### NewCalendareventsIdJsonEvent

`func NewCalendareventsIdJsonEvent() *CalendareventsIdJsonEvent`

NewCalendareventsIdJsonEvent instantiates a new CalendareventsIdJsonEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalendareventsIdJsonEventWithDefaults

`func NewCalendareventsIdJsonEventWithDefaults() *CalendareventsIdJsonEvent`

NewCalendareventsIdJsonEventWithDefaults instantiates a new CalendareventsIdJsonEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllDay

`func (o *CalendareventsIdJsonEvent) GetAllDay() string`

GetAllDay returns the AllDay field if non-nil, zero value otherwise.

### GetAllDayOk

`func (o *CalendareventsIdJsonEvent) GetAllDayOk() (*string, bool)`

GetAllDayOk returns a tuple with the AllDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllDay

`func (o *CalendareventsIdJsonEvent) SetAllDay(v string)`

SetAllDay sets AllDay field to given value.

### HasAllDay

`func (o *CalendareventsIdJsonEvent) HasAllDay() bool`

HasAllDay returns a boolean if a field has been set.

### GetAttendeesCanEdit

`func (o *CalendareventsIdJsonEvent) GetAttendeesCanEdit() string`

GetAttendeesCanEdit returns the AttendeesCanEdit field if non-nil, zero value otherwise.

### GetAttendeesCanEditOk

`func (o *CalendareventsIdJsonEvent) GetAttendeesCanEditOk() (*string, bool)`

GetAttendeesCanEditOk returns a tuple with the AttendeesCanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendeesCanEdit

`func (o *CalendareventsIdJsonEvent) SetAttendeesCanEdit(v string)`

SetAttendeesCanEdit sets AttendeesCanEdit field to given value.

### HasAttendeesCanEdit

`func (o *CalendareventsIdJsonEvent) HasAttendeesCanEdit() bool`

HasAttendeesCanEdit returns a boolean if a field has been set.

### GetAttendingUserIds

`func (o *CalendareventsIdJsonEvent) GetAttendingUserIds() string`

GetAttendingUserIds returns the AttendingUserIds field if non-nil, zero value otherwise.

### GetAttendingUserIdsOk

`func (o *CalendareventsIdJsonEvent) GetAttendingUserIdsOk() (*string, bool)`

GetAttendingUserIdsOk returns a tuple with the AttendingUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendingUserIds

`func (o *CalendareventsIdJsonEvent) SetAttendingUserIds(v string)`

SetAttendingUserIds sets AttendingUserIds field to given value.

### HasAttendingUserIds

`func (o *CalendareventsIdJsonEvent) HasAttendingUserIds() bool`

HasAttendingUserIds returns a boolean if a field has been set.

### GetDescription

`func (o *CalendareventsIdJsonEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CalendareventsIdJsonEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CalendareventsIdJsonEvent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CalendareventsIdJsonEvent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnd

`func (o *CalendareventsIdJsonEvent) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *CalendareventsIdJsonEvent) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *CalendareventsIdJsonEvent) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *CalendareventsIdJsonEvent) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetNotify

`func (o *CalendareventsIdJsonEvent) GetNotify() string`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *CalendareventsIdJsonEvent) GetNotifyOk() (*string, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *CalendareventsIdJsonEvent) SetNotify(v string)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *CalendareventsIdJsonEvent) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetNotifyUserIds

`func (o *CalendareventsIdJsonEvent) GetNotifyUserIds() string`

GetNotifyUserIds returns the NotifyUserIds field if non-nil, zero value otherwise.

### GetNotifyUserIdsOk

`func (o *CalendareventsIdJsonEvent) GetNotifyUserIdsOk() (*string, bool)`

GetNotifyUserIdsOk returns a tuple with the NotifyUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUserIds

`func (o *CalendareventsIdJsonEvent) SetNotifyUserIds(v string)`

SetNotifyUserIds sets NotifyUserIds field to given value.

### HasNotifyUserIds

`func (o *CalendareventsIdJsonEvent) HasNotifyUserIds() bool`

HasNotifyUserIds returns a boolean if a field has been set.

### GetPrivacy

`func (o *CalendareventsIdJsonEvent) GetPrivacy() InlineResponse2003Privacy`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *CalendareventsIdJsonEvent) GetPrivacyOk() (*InlineResponse2003Privacy, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *CalendareventsIdJsonEvent) SetPrivacy(v InlineResponse2003Privacy)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *CalendareventsIdJsonEvent) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetProjectId

`func (o *CalendareventsIdJsonEvent) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CalendareventsIdJsonEvent) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CalendareventsIdJsonEvent) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CalendareventsIdJsonEvent) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProjectUsersCanEdit

`func (o *CalendareventsIdJsonEvent) GetProjectUsersCanEdit() string`

GetProjectUsersCanEdit returns the ProjectUsersCanEdit field if non-nil, zero value otherwise.

### GetProjectUsersCanEditOk

`func (o *CalendareventsIdJsonEvent) GetProjectUsersCanEditOk() (*string, bool)`

GetProjectUsersCanEditOk returns a tuple with the ProjectUsersCanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectUsersCanEdit

`func (o *CalendareventsIdJsonEvent) SetProjectUsersCanEdit(v string)`

SetProjectUsersCanEdit sets ProjectUsersCanEdit field to given value.

### HasProjectUsersCanEdit

`func (o *CalendareventsIdJsonEvent) HasProjectUsersCanEdit() bool`

HasProjectUsersCanEdit returns a boolean if a field has been set.

### GetReminders

`func (o *CalendareventsIdJsonEvent) GetReminders() []CalendareventsIdJsonEventReminders`

GetReminders returns the Reminders field if non-nil, zero value otherwise.

### GetRemindersOk

`func (o *CalendareventsIdJsonEvent) GetRemindersOk() (*[]CalendareventsIdJsonEventReminders, bool)`

GetRemindersOk returns a tuple with the Reminders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminders

`func (o *CalendareventsIdJsonEvent) SetReminders(v []CalendareventsIdJsonEventReminders)`

SetReminders sets Reminders field to given value.

### HasReminders

`func (o *CalendareventsIdJsonEvent) HasReminders() bool`

HasReminders returns a boolean if a field has been set.

### GetShowAsBusy

`func (o *CalendareventsIdJsonEvent) GetShowAsBusy() string`

GetShowAsBusy returns the ShowAsBusy field if non-nil, zero value otherwise.

### GetShowAsBusyOk

`func (o *CalendareventsIdJsonEvent) GetShowAsBusyOk() (*string, bool)`

GetShowAsBusyOk returns a tuple with the ShowAsBusy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAsBusy

`func (o *CalendareventsIdJsonEvent) SetShowAsBusy(v string)`

SetShowAsBusy sets ShowAsBusy field to given value.

### HasShowAsBusy

`func (o *CalendareventsIdJsonEvent) HasShowAsBusy() bool`

HasShowAsBusy returns a boolean if a field has been set.

### GetStart

`func (o *CalendareventsIdJsonEvent) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *CalendareventsIdJsonEvent) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *CalendareventsIdJsonEvent) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *CalendareventsIdJsonEvent) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTitle

`func (o *CalendareventsIdJsonEvent) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CalendareventsIdJsonEvent) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CalendareventsIdJsonEvent) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CalendareventsIdJsonEvent) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *CalendareventsIdJsonEvent) GetType() CalendareventsIdJsonEventType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CalendareventsIdJsonEvent) GetTypeOk() (*CalendareventsIdJsonEventType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CalendareventsIdJsonEvent) SetType(v CalendareventsIdJsonEventType)`

SetType sets Type field to given value.

### HasType

`func (o *CalendareventsIdJsonEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUtc

`func (o *CalendareventsIdJsonEvent) GetUtc() bool`

GetUtc returns the Utc field if non-nil, zero value otherwise.

### GetUtcOk

`func (o *CalendareventsIdJsonEvent) GetUtcOk() (*bool, bool)`

GetUtcOk returns a tuple with the Utc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtc

`func (o *CalendareventsIdJsonEvent) SetUtc(v bool)`

SetUtc sets Utc field to given value.

### HasUtc

`func (o *CalendareventsIdJsonEvent) HasUtc() bool`

HasUtc returns a boolean if a field has been set.

### GetWhere

`func (o *CalendareventsIdJsonEvent) GetWhere() string`

GetWhere returns the Where field if non-nil, zero value otherwise.

### GetWhereOk

`func (o *CalendareventsIdJsonEvent) GetWhereOk() (*string, bool)`

GetWhereOk returns a tuple with the Where field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhere

`func (o *CalendareventsIdJsonEvent) SetWhere(v string)`

SetWhere sets Where field to given value.

### HasWhere

`func (o *CalendareventsIdJsonEvent) HasWhere() bool`

HasWhere returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


