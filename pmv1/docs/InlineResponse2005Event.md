# InlineResponse2005Event

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
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**InlineResponse2002Column**](InlineResponse2002Column.md) |  | [optional] 
**Where** | Pointer to **string** |  | [optional] 

## Methods

### NewInlineResponse2005Event

`func NewInlineResponse2005Event() *InlineResponse2005Event`

NewInlineResponse2005Event instantiates a new InlineResponse2005Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2005EventWithDefaults

`func NewInlineResponse2005EventWithDefaults() *InlineResponse2005Event`

NewInlineResponse2005EventWithDefaults instantiates a new InlineResponse2005Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllDay

`func (o *InlineResponse2005Event) GetAllDay() bool`

GetAllDay returns the AllDay field if non-nil, zero value otherwise.

### GetAllDayOk

`func (o *InlineResponse2005Event) GetAllDayOk() (*bool, bool)`

GetAllDayOk returns a tuple with the AllDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllDay

`func (o *InlineResponse2005Event) SetAllDay(v bool)`

SetAllDay sets AllDay field to given value.

### HasAllDay

`func (o *InlineResponse2005Event) HasAllDay() bool`

HasAllDay returns a boolean if a field has been set.

### GetAttendeesCanEdit

`func (o *InlineResponse2005Event) GetAttendeesCanEdit() bool`

GetAttendeesCanEdit returns the AttendeesCanEdit field if non-nil, zero value otherwise.

### GetAttendeesCanEditOk

`func (o *InlineResponse2005Event) GetAttendeesCanEditOk() (*bool, bool)`

GetAttendeesCanEditOk returns a tuple with the AttendeesCanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendeesCanEdit

`func (o *InlineResponse2005Event) SetAttendeesCanEdit(v bool)`

SetAttendeesCanEdit sets AttendeesCanEdit field to given value.

### HasAttendeesCanEdit

`func (o *InlineResponse2005Event) HasAttendeesCanEdit() bool`

HasAttendeesCanEdit returns a boolean if a field has been set.

### GetAttendingUserIds

`func (o *InlineResponse2005Event) GetAttendingUserIds() string`

GetAttendingUserIds returns the AttendingUserIds field if non-nil, zero value otherwise.

### GetAttendingUserIdsOk

`func (o *InlineResponse2005Event) GetAttendingUserIdsOk() (*string, bool)`

GetAttendingUserIdsOk returns a tuple with the AttendingUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendingUserIds

`func (o *InlineResponse2005Event) SetAttendingUserIds(v string)`

SetAttendingUserIds sets AttendingUserIds field to given value.

### HasAttendingUserIds

`func (o *InlineResponse2005Event) HasAttendingUserIds() bool`

HasAttendingUserIds returns a boolean if a field has been set.

### GetAttendingUserNames

`func (o *InlineResponse2005Event) GetAttendingUserNames() string`

GetAttendingUserNames returns the AttendingUserNames field if non-nil, zero value otherwise.

### GetAttendingUserNamesOk

`func (o *InlineResponse2005Event) GetAttendingUserNamesOk() (*string, bool)`

GetAttendingUserNamesOk returns a tuple with the AttendingUserNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendingUserNames

`func (o *InlineResponse2005Event) SetAttendingUserNames(v string)`

SetAttendingUserNames sets AttendingUserNames field to given value.

### HasAttendingUserNames

`func (o *InlineResponse2005Event) HasAttendingUserNames() bool`

HasAttendingUserNames returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse2005Event) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse2005Event) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse2005Event) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse2005Event) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnd

`func (o *InlineResponse2005Event) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *InlineResponse2005Event) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *InlineResponse2005Event) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *InlineResponse2005Event) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetId

`func (o *InlineResponse2005Event) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse2005Event) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse2005Event) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse2005Event) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastChangedOn

`func (o *InlineResponse2005Event) GetLastChangedOn() string`

GetLastChangedOn returns the LastChangedOn field if non-nil, zero value otherwise.

### GetLastChangedOnOk

`func (o *InlineResponse2005Event) GetLastChangedOnOk() (*string, bool)`

GetLastChangedOnOk returns a tuple with the LastChangedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangedOn

`func (o *InlineResponse2005Event) SetLastChangedOn(v string)`

SetLastChangedOn sets LastChangedOn field to given value.

### HasLastChangedOn

`func (o *InlineResponse2005Event) HasLastChangedOn() bool`

HasLastChangedOn returns a boolean if a field has been set.

### GetNotifyUserIds

`func (o *InlineResponse2005Event) GetNotifyUserIds() string`

GetNotifyUserIds returns the NotifyUserIds field if non-nil, zero value otherwise.

### GetNotifyUserIdsOk

`func (o *InlineResponse2005Event) GetNotifyUserIdsOk() (*string, bool)`

GetNotifyUserIdsOk returns a tuple with the NotifyUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUserIds

`func (o *InlineResponse2005Event) SetNotifyUserIds(v string)`

SetNotifyUserIds sets NotifyUserIds field to given value.

### HasNotifyUserIds

`func (o *InlineResponse2005Event) HasNotifyUserIds() bool`

HasNotifyUserIds returns a boolean if a field has been set.

### GetNotifyUserNames

`func (o *InlineResponse2005Event) GetNotifyUserNames() string`

GetNotifyUserNames returns the NotifyUserNames field if non-nil, zero value otherwise.

### GetNotifyUserNamesOk

`func (o *InlineResponse2005Event) GetNotifyUserNamesOk() (*string, bool)`

GetNotifyUserNamesOk returns a tuple with the NotifyUserNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUserNames

`func (o *InlineResponse2005Event) SetNotifyUserNames(v string)`

SetNotifyUserNames sets NotifyUserNames field to given value.

### HasNotifyUserNames

`func (o *InlineResponse2005Event) HasNotifyUserNames() bool`

HasNotifyUserNames returns a boolean if a field has been set.

### GetOwner

`func (o *InlineResponse2005Event) GetOwner() InlineResponse2003Owner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *InlineResponse2005Event) GetOwnerOk() (*InlineResponse2003Owner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *InlineResponse2005Event) SetOwner(v InlineResponse2003Owner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *InlineResponse2005Event) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPrivacy

`func (o *InlineResponse2005Event) GetPrivacy() InlineResponse2003Privacy`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *InlineResponse2005Event) GetPrivacyOk() (*InlineResponse2003Privacy, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *InlineResponse2005Event) SetPrivacy(v InlineResponse2003Privacy)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *InlineResponse2005Event) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetProjectUsersCanEdit

`func (o *InlineResponse2005Event) GetProjectUsersCanEdit() bool`

GetProjectUsersCanEdit returns the ProjectUsersCanEdit field if non-nil, zero value otherwise.

### GetProjectUsersCanEditOk

`func (o *InlineResponse2005Event) GetProjectUsersCanEditOk() (*bool, bool)`

GetProjectUsersCanEditOk returns a tuple with the ProjectUsersCanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectUsersCanEdit

`func (o *InlineResponse2005Event) SetProjectUsersCanEdit(v bool)`

SetProjectUsersCanEdit sets ProjectUsersCanEdit field to given value.

### HasProjectUsersCanEdit

`func (o *InlineResponse2005Event) HasProjectUsersCanEdit() bool`

HasProjectUsersCanEdit returns a boolean if a field has been set.

### GetReminders

`func (o *InlineResponse2005Event) GetReminders() []map[string]interface{}`

GetReminders returns the Reminders field if non-nil, zero value otherwise.

### GetRemindersOk

`func (o *InlineResponse2005Event) GetRemindersOk() (*[]map[string]interface{}, bool)`

GetRemindersOk returns a tuple with the Reminders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminders

`func (o *InlineResponse2005Event) SetReminders(v []map[string]interface{})`

SetReminders sets Reminders field to given value.

### HasReminders

`func (o *InlineResponse2005Event) HasReminders() bool`

HasReminders returns a boolean if a field has been set.

### GetRepeat

`func (o *InlineResponse2005Event) GetRepeat() map[string]interface{}`

GetRepeat returns the Repeat field if non-nil, zero value otherwise.

### GetRepeatOk

`func (o *InlineResponse2005Event) GetRepeatOk() (*map[string]interface{}, bool)`

GetRepeatOk returns a tuple with the Repeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeat

`func (o *InlineResponse2005Event) SetRepeat(v map[string]interface{})`

SetRepeat sets Repeat field to given value.

### HasRepeat

`func (o *InlineResponse2005Event) HasRepeat() bool`

HasRepeat returns a boolean if a field has been set.

### GetShowAsBusy

`func (o *InlineResponse2005Event) GetShowAsBusy() bool`

GetShowAsBusy returns the ShowAsBusy field if non-nil, zero value otherwise.

### GetShowAsBusyOk

`func (o *InlineResponse2005Event) GetShowAsBusyOk() (*bool, bool)`

GetShowAsBusyOk returns a tuple with the ShowAsBusy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAsBusy

`func (o *InlineResponse2005Event) SetShowAsBusy(v bool)`

SetShowAsBusy sets ShowAsBusy field to given value.

### HasShowAsBusy

`func (o *InlineResponse2005Event) HasShowAsBusy() bool`

HasShowAsBusy returns a boolean if a field has been set.

### GetStart

`func (o *InlineResponse2005Event) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *InlineResponse2005Event) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *InlineResponse2005Event) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *InlineResponse2005Event) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTitle

`func (o *InlineResponse2005Event) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InlineResponse2005Event) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InlineResponse2005Event) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InlineResponse2005Event) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse2005Event) GetType() InlineResponse2002Column`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse2005Event) GetTypeOk() (*InlineResponse2002Column, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse2005Event) SetType(v InlineResponse2002Column)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse2005Event) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWhere

`func (o *InlineResponse2005Event) GetWhere() string`

GetWhere returns the Where field if non-nil, zero value otherwise.

### GetWhereOk

`func (o *InlineResponse2005Event) GetWhereOk() (*string, bool)`

GetWhereOk returns a tuple with the Where field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhere

`func (o *InlineResponse2005Event) SetWhere(v string)`

SetWhere sets Where field to given value.

### HasWhere

`func (o *InlineResponse2005Event) HasWhere() bool`

HasWhere returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


