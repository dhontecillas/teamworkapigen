# InlineResponse2003Events

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllDay** | Pointer to **bool** |  | [optional] 
**AttendeesCanEdit** | Pointer to **bool** |  | [optional] 
**AttendingUserIds** | Pointer to **string** |  | [optional] 
**AttendingUserNames** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**End** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastChangedOn** | Pointer to **string** |  | [optional] 
**NotifyUserIds** | Pointer to **string** |  | [optional] 
**NotifyUserNames** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**InlineResponse2003Owner**](InlineResponse2003Owner.md) |  | [optional] 
**Privacy** | Pointer to [**InlineResponse2003Privacy**](InlineResponse2003Privacy.md) |  | [optional] 
**ProjectUsersCanEdit** | Pointer to **bool** |  | [optional] 
**Reminders** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Repeat** | Pointer to **map[string]interface{}** |  | [optional] 
**ShowAsBusy** | Pointer to **bool** |  | [optional] 
**Start** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**InlineResponse2002Column**](InlineResponse2002Column.md) |  | [optional] 
**Where** | Pointer to **string** |  | [optional] 

## Methods

### NewInlineResponse2003Events

`func NewInlineResponse2003Events() *InlineResponse2003Events`

NewInlineResponse2003Events instantiates a new InlineResponse2003Events object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2003EventsWithDefaults

`func NewInlineResponse2003EventsWithDefaults() *InlineResponse2003Events`

NewInlineResponse2003EventsWithDefaults instantiates a new InlineResponse2003Events object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllDay

`func (o *InlineResponse2003Events) GetAllDay() bool`

GetAllDay returns the AllDay field if non-nil, zero value otherwise.

### GetAllDayOk

`func (o *InlineResponse2003Events) GetAllDayOk() (*bool, bool)`

GetAllDayOk returns a tuple with the AllDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllDay

`func (o *InlineResponse2003Events) SetAllDay(v bool)`

SetAllDay sets AllDay field to given value.

### HasAllDay

`func (o *InlineResponse2003Events) HasAllDay() bool`

HasAllDay returns a boolean if a field has been set.

### GetAttendeesCanEdit

`func (o *InlineResponse2003Events) GetAttendeesCanEdit() bool`

GetAttendeesCanEdit returns the AttendeesCanEdit field if non-nil, zero value otherwise.

### GetAttendeesCanEditOk

`func (o *InlineResponse2003Events) GetAttendeesCanEditOk() (*bool, bool)`

GetAttendeesCanEditOk returns a tuple with the AttendeesCanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendeesCanEdit

`func (o *InlineResponse2003Events) SetAttendeesCanEdit(v bool)`

SetAttendeesCanEdit sets AttendeesCanEdit field to given value.

### HasAttendeesCanEdit

`func (o *InlineResponse2003Events) HasAttendeesCanEdit() bool`

HasAttendeesCanEdit returns a boolean if a field has been set.

### GetAttendingUserIds

`func (o *InlineResponse2003Events) GetAttendingUserIds() string`

GetAttendingUserIds returns the AttendingUserIds field if non-nil, zero value otherwise.

### GetAttendingUserIdsOk

`func (o *InlineResponse2003Events) GetAttendingUserIdsOk() (*string, bool)`

GetAttendingUserIdsOk returns a tuple with the AttendingUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendingUserIds

`func (o *InlineResponse2003Events) SetAttendingUserIds(v string)`

SetAttendingUserIds sets AttendingUserIds field to given value.

### HasAttendingUserIds

`func (o *InlineResponse2003Events) HasAttendingUserIds() bool`

HasAttendingUserIds returns a boolean if a field has been set.

### GetAttendingUserNames

`func (o *InlineResponse2003Events) GetAttendingUserNames() string`

GetAttendingUserNames returns the AttendingUserNames field if non-nil, zero value otherwise.

### GetAttendingUserNamesOk

`func (o *InlineResponse2003Events) GetAttendingUserNamesOk() (*string, bool)`

GetAttendingUserNamesOk returns a tuple with the AttendingUserNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendingUserNames

`func (o *InlineResponse2003Events) SetAttendingUserNames(v string)`

SetAttendingUserNames sets AttendingUserNames field to given value.

### HasAttendingUserNames

`func (o *InlineResponse2003Events) HasAttendingUserNames() bool`

HasAttendingUserNames returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse2003Events) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse2003Events) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse2003Events) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse2003Events) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnd

`func (o *InlineResponse2003Events) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *InlineResponse2003Events) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *InlineResponse2003Events) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *InlineResponse2003Events) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetId

`func (o *InlineResponse2003Events) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse2003Events) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse2003Events) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse2003Events) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastChangedOn

`func (o *InlineResponse2003Events) GetLastChangedOn() string`

GetLastChangedOn returns the LastChangedOn field if non-nil, zero value otherwise.

### GetLastChangedOnOk

`func (o *InlineResponse2003Events) GetLastChangedOnOk() (*string, bool)`

GetLastChangedOnOk returns a tuple with the LastChangedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangedOn

`func (o *InlineResponse2003Events) SetLastChangedOn(v string)`

SetLastChangedOn sets LastChangedOn field to given value.

### HasLastChangedOn

`func (o *InlineResponse2003Events) HasLastChangedOn() bool`

HasLastChangedOn returns a boolean if a field has been set.

### GetNotifyUserIds

`func (o *InlineResponse2003Events) GetNotifyUserIds() string`

GetNotifyUserIds returns the NotifyUserIds field if non-nil, zero value otherwise.

### GetNotifyUserIdsOk

`func (o *InlineResponse2003Events) GetNotifyUserIdsOk() (*string, bool)`

GetNotifyUserIdsOk returns a tuple with the NotifyUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUserIds

`func (o *InlineResponse2003Events) SetNotifyUserIds(v string)`

SetNotifyUserIds sets NotifyUserIds field to given value.

### HasNotifyUserIds

`func (o *InlineResponse2003Events) HasNotifyUserIds() bool`

HasNotifyUserIds returns a boolean if a field has been set.

### GetNotifyUserNames

`func (o *InlineResponse2003Events) GetNotifyUserNames() string`

GetNotifyUserNames returns the NotifyUserNames field if non-nil, zero value otherwise.

### GetNotifyUserNamesOk

`func (o *InlineResponse2003Events) GetNotifyUserNamesOk() (*string, bool)`

GetNotifyUserNamesOk returns a tuple with the NotifyUserNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUserNames

`func (o *InlineResponse2003Events) SetNotifyUserNames(v string)`

SetNotifyUserNames sets NotifyUserNames field to given value.

### HasNotifyUserNames

`func (o *InlineResponse2003Events) HasNotifyUserNames() bool`

HasNotifyUserNames returns a boolean if a field has been set.

### GetOwner

`func (o *InlineResponse2003Events) GetOwner() InlineResponse2003Owner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *InlineResponse2003Events) GetOwnerOk() (*InlineResponse2003Owner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *InlineResponse2003Events) SetOwner(v InlineResponse2003Owner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *InlineResponse2003Events) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPrivacy

`func (o *InlineResponse2003Events) GetPrivacy() InlineResponse2003Privacy`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *InlineResponse2003Events) GetPrivacyOk() (*InlineResponse2003Privacy, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *InlineResponse2003Events) SetPrivacy(v InlineResponse2003Privacy)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *InlineResponse2003Events) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetProjectUsersCanEdit

`func (o *InlineResponse2003Events) GetProjectUsersCanEdit() bool`

GetProjectUsersCanEdit returns the ProjectUsersCanEdit field if non-nil, zero value otherwise.

### GetProjectUsersCanEditOk

`func (o *InlineResponse2003Events) GetProjectUsersCanEditOk() (*bool, bool)`

GetProjectUsersCanEditOk returns a tuple with the ProjectUsersCanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectUsersCanEdit

`func (o *InlineResponse2003Events) SetProjectUsersCanEdit(v bool)`

SetProjectUsersCanEdit sets ProjectUsersCanEdit field to given value.

### HasProjectUsersCanEdit

`func (o *InlineResponse2003Events) HasProjectUsersCanEdit() bool`

HasProjectUsersCanEdit returns a boolean if a field has been set.

### GetReminders

`func (o *InlineResponse2003Events) GetReminders() []map[string]interface{}`

GetReminders returns the Reminders field if non-nil, zero value otherwise.

### GetRemindersOk

`func (o *InlineResponse2003Events) GetRemindersOk() (*[]map[string]interface{}, bool)`

GetRemindersOk returns a tuple with the Reminders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminders

`func (o *InlineResponse2003Events) SetReminders(v []map[string]interface{})`

SetReminders sets Reminders field to given value.

### HasReminders

`func (o *InlineResponse2003Events) HasReminders() bool`

HasReminders returns a boolean if a field has been set.

### GetRepeat

`func (o *InlineResponse2003Events) GetRepeat() map[string]interface{}`

GetRepeat returns the Repeat field if non-nil, zero value otherwise.

### GetRepeatOk

`func (o *InlineResponse2003Events) GetRepeatOk() (*map[string]interface{}, bool)`

GetRepeatOk returns a tuple with the Repeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeat

`func (o *InlineResponse2003Events) SetRepeat(v map[string]interface{})`

SetRepeat sets Repeat field to given value.

### HasRepeat

`func (o *InlineResponse2003Events) HasRepeat() bool`

HasRepeat returns a boolean if a field has been set.

### GetShowAsBusy

`func (o *InlineResponse2003Events) GetShowAsBusy() bool`

GetShowAsBusy returns the ShowAsBusy field if non-nil, zero value otherwise.

### GetShowAsBusyOk

`func (o *InlineResponse2003Events) GetShowAsBusyOk() (*bool, bool)`

GetShowAsBusyOk returns a tuple with the ShowAsBusy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAsBusy

`func (o *InlineResponse2003Events) SetShowAsBusy(v bool)`

SetShowAsBusy sets ShowAsBusy field to given value.

### HasShowAsBusy

`func (o *InlineResponse2003Events) HasShowAsBusy() bool`

HasShowAsBusy returns a boolean if a field has been set.

### GetStart

`func (o *InlineResponse2003Events) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *InlineResponse2003Events) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *InlineResponse2003Events) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *InlineResponse2003Events) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse2003Events) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse2003Events) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse2003Events) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse2003Events) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *InlineResponse2003Events) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InlineResponse2003Events) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InlineResponse2003Events) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InlineResponse2003Events) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse2003Events) GetType() InlineResponse2002Column`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse2003Events) GetTypeOk() (*InlineResponse2002Column, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse2003Events) SetType(v InlineResponse2002Column)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse2003Events) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWhere

`func (o *InlineResponse2003Events) GetWhere() string`

GetWhere returns the Where field if non-nil, zero value otherwise.

### GetWhereOk

`func (o *InlineResponse2003Events) GetWhereOk() (*string, bool)`

GetWhereOk returns a tuple with the Where field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhere

`func (o *InlineResponse2003Events) SetWhere(v string)`

SetWhere sets Where field to given value.

### HasWhere

`func (o *InlineResponse2003Events) HasWhere() bool`

HasWhere returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


