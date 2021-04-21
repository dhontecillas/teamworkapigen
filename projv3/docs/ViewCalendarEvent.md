# ViewCalendarEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalPropertiesField** | Pointer to **string** |  | [optional] 
**AllDay** | Pointer to **bool** |  | [optional] 
**AttendeesCanEdit** | Pointer to **bool** |  | [optional] 
**AttendingUserIds** | Pointer to **[]int32** |  | [optional] 
**AttendingUsers** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **int32** |  | [optional] 
**CreatedByUserId** | Pointer to **int32** |  | [optional] 
**CreatedDateTime** | Pointer to **string** |  | [optional] 
**CurrentUserAssociationType** | Pointer to **string** |  | [optional] 
**DateDeleted** | Pointer to **string** |  | [optional] 
**DateLastUpdated** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**DeletedBy** | Pointer to **int32** |  | [optional] 
**DeletedByUserId** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**MonthlyRepeatType** | Pointer to **string** |  | [optional] 
**OwnedBy** | Pointer to **int32** |  | [optional] 
**OwnerUserId** | Pointer to **int32** |  | [optional] 
**PrivacyType** | Pointer to **string** |  | [optional] 
**Project** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**ProjectUsersCanEdit** | Pointer to **bool** |  | [optional] 
**Sequence** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**SequenceId** | Pointer to **int32** |  | [optional] 
**ShowAsBusy** | Pointer to **bool** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**TypeId** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **int32** |  | [optional] 
**UpdatedByUserId** | Pointer to **int32** |  | [optional] 

## Methods

### NewViewCalendarEvent

`func NewViewCalendarEvent() *ViewCalendarEvent`

NewViewCalendarEvent instantiates a new ViewCalendarEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewCalendarEventWithDefaults

`func NewViewCalendarEventWithDefaults() *ViewCalendarEvent`

NewViewCalendarEventWithDefaults instantiates a new ViewCalendarEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalPropertiesField

`func (o *ViewCalendarEvent) GetAdditionalPropertiesField() string`

GetAdditionalPropertiesField returns the AdditionalPropertiesField field if non-nil, zero value otherwise.

### GetAdditionalPropertiesFieldOk

`func (o *ViewCalendarEvent) GetAdditionalPropertiesFieldOk() (*string, bool)`

GetAdditionalPropertiesFieldOk returns a tuple with the AdditionalPropertiesField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPropertiesField

`func (o *ViewCalendarEvent) SetAdditionalPropertiesField(v string)`

SetAdditionalPropertiesField sets AdditionalPropertiesField field to given value.

### HasAdditionalPropertiesField

`func (o *ViewCalendarEvent) HasAdditionalPropertiesField() bool`

HasAdditionalPropertiesField returns a boolean if a field has been set.

### GetAllDay

`func (o *ViewCalendarEvent) GetAllDay() bool`

GetAllDay returns the AllDay field if non-nil, zero value otherwise.

### GetAllDayOk

`func (o *ViewCalendarEvent) GetAllDayOk() (*bool, bool)`

GetAllDayOk returns a tuple with the AllDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllDay

`func (o *ViewCalendarEvent) SetAllDay(v bool)`

SetAllDay sets AllDay field to given value.

### HasAllDay

`func (o *ViewCalendarEvent) HasAllDay() bool`

HasAllDay returns a boolean if a field has been set.

### GetAttendeesCanEdit

`func (o *ViewCalendarEvent) GetAttendeesCanEdit() bool`

GetAttendeesCanEdit returns the AttendeesCanEdit field if non-nil, zero value otherwise.

### GetAttendeesCanEditOk

`func (o *ViewCalendarEvent) GetAttendeesCanEditOk() (*bool, bool)`

GetAttendeesCanEditOk returns a tuple with the AttendeesCanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendeesCanEdit

`func (o *ViewCalendarEvent) SetAttendeesCanEdit(v bool)`

SetAttendeesCanEdit sets AttendeesCanEdit field to given value.

### HasAttendeesCanEdit

`func (o *ViewCalendarEvent) HasAttendeesCanEdit() bool`

HasAttendeesCanEdit returns a boolean if a field has been set.

### GetAttendingUserIds

`func (o *ViewCalendarEvent) GetAttendingUserIds() []int32`

GetAttendingUserIds returns the AttendingUserIds field if non-nil, zero value otherwise.

### GetAttendingUserIdsOk

`func (o *ViewCalendarEvent) GetAttendingUserIdsOk() (*[]int32, bool)`

GetAttendingUserIdsOk returns a tuple with the AttendingUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendingUserIds

`func (o *ViewCalendarEvent) SetAttendingUserIds(v []int32)`

SetAttendingUserIds sets AttendingUserIds field to given value.

### HasAttendingUserIds

`func (o *ViewCalendarEvent) HasAttendingUserIds() bool`

HasAttendingUserIds returns a boolean if a field has been set.

### GetAttendingUsers

`func (o *ViewCalendarEvent) GetAttendingUsers() []ViewRelationship`

GetAttendingUsers returns the AttendingUsers field if non-nil, zero value otherwise.

### GetAttendingUsersOk

`func (o *ViewCalendarEvent) GetAttendingUsersOk() (*[]ViewRelationship, bool)`

GetAttendingUsersOk returns a tuple with the AttendingUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendingUsers

`func (o *ViewCalendarEvent) SetAttendingUsers(v []ViewRelationship)`

SetAttendingUsers sets AttendingUsers field to given value.

### HasAttendingUsers

`func (o *ViewCalendarEvent) HasAttendingUsers() bool`

HasAttendingUsers returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ViewCalendarEvent) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ViewCalendarEvent) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ViewCalendarEvent) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ViewCalendarEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ViewCalendarEvent) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ViewCalendarEvent) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ViewCalendarEvent) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ViewCalendarEvent) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *ViewCalendarEvent) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *ViewCalendarEvent) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *ViewCalendarEvent) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *ViewCalendarEvent) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *ViewCalendarEvent) GetCreatedDateTime() string`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *ViewCalendarEvent) GetCreatedDateTimeOk() (*string, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *ViewCalendarEvent) SetCreatedDateTime(v string)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *ViewCalendarEvent) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetCurrentUserAssociationType

`func (o *ViewCalendarEvent) GetCurrentUserAssociationType() string`

GetCurrentUserAssociationType returns the CurrentUserAssociationType field if non-nil, zero value otherwise.

### GetCurrentUserAssociationTypeOk

`func (o *ViewCalendarEvent) GetCurrentUserAssociationTypeOk() (*string, bool)`

GetCurrentUserAssociationTypeOk returns a tuple with the CurrentUserAssociationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUserAssociationType

`func (o *ViewCalendarEvent) SetCurrentUserAssociationType(v string)`

SetCurrentUserAssociationType sets CurrentUserAssociationType field to given value.

### HasCurrentUserAssociationType

`func (o *ViewCalendarEvent) HasCurrentUserAssociationType() bool`

HasCurrentUserAssociationType returns a boolean if a field has been set.

### GetDateDeleted

`func (o *ViewCalendarEvent) GetDateDeleted() string`

GetDateDeleted returns the DateDeleted field if non-nil, zero value otherwise.

### GetDateDeletedOk

`func (o *ViewCalendarEvent) GetDateDeletedOk() (*string, bool)`

GetDateDeletedOk returns a tuple with the DateDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeleted

`func (o *ViewCalendarEvent) SetDateDeleted(v string)`

SetDateDeleted sets DateDeleted field to given value.

### HasDateDeleted

`func (o *ViewCalendarEvent) HasDateDeleted() bool`

HasDateDeleted returns a boolean if a field has been set.

### GetDateLastUpdated

`func (o *ViewCalendarEvent) GetDateLastUpdated() string`

GetDateLastUpdated returns the DateLastUpdated field if non-nil, zero value otherwise.

### GetDateLastUpdatedOk

`func (o *ViewCalendarEvent) GetDateLastUpdatedOk() (*string, bool)`

GetDateLastUpdatedOk returns a tuple with the DateLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateLastUpdated

`func (o *ViewCalendarEvent) SetDateLastUpdated(v string)`

SetDateLastUpdated sets DateLastUpdated field to given value.

### HasDateLastUpdated

`func (o *ViewCalendarEvent) HasDateLastUpdated() bool`

HasDateLastUpdated returns a boolean if a field has been set.

### GetDeleted

`func (o *ViewCalendarEvent) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *ViewCalendarEvent) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *ViewCalendarEvent) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *ViewCalendarEvent) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ViewCalendarEvent) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ViewCalendarEvent) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ViewCalendarEvent) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ViewCalendarEvent) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeletedBy

`func (o *ViewCalendarEvent) GetDeletedBy() int32`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *ViewCalendarEvent) GetDeletedByOk() (*int32, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *ViewCalendarEvent) SetDeletedBy(v int32)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *ViewCalendarEvent) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetDeletedByUserId

`func (o *ViewCalendarEvent) GetDeletedByUserId() int32`

GetDeletedByUserId returns the DeletedByUserId field if non-nil, zero value otherwise.

### GetDeletedByUserIdOk

`func (o *ViewCalendarEvent) GetDeletedByUserIdOk() (*int32, bool)`

GetDeletedByUserIdOk returns a tuple with the DeletedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedByUserId

`func (o *ViewCalendarEvent) SetDeletedByUserId(v int32)`

SetDeletedByUserId sets DeletedByUserId field to given value.

### HasDeletedByUserId

`func (o *ViewCalendarEvent) HasDeletedByUserId() bool`

HasDeletedByUserId returns a boolean if a field has been set.

### GetDescription

`func (o *ViewCalendarEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ViewCalendarEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ViewCalendarEvent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ViewCalendarEvent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndDate

`func (o *ViewCalendarEvent) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ViewCalendarEvent) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ViewCalendarEvent) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ViewCalendarEvent) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetId

`func (o *ViewCalendarEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewCalendarEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewCalendarEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewCalendarEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocation

`func (o *ViewCalendarEvent) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ViewCalendarEvent) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ViewCalendarEvent) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ViewCalendarEvent) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMonthlyRepeatType

`func (o *ViewCalendarEvent) GetMonthlyRepeatType() string`

GetMonthlyRepeatType returns the MonthlyRepeatType field if non-nil, zero value otherwise.

### GetMonthlyRepeatTypeOk

`func (o *ViewCalendarEvent) GetMonthlyRepeatTypeOk() (*string, bool)`

GetMonthlyRepeatTypeOk returns a tuple with the MonthlyRepeatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyRepeatType

`func (o *ViewCalendarEvent) SetMonthlyRepeatType(v string)`

SetMonthlyRepeatType sets MonthlyRepeatType field to given value.

### HasMonthlyRepeatType

`func (o *ViewCalendarEvent) HasMonthlyRepeatType() bool`

HasMonthlyRepeatType returns a boolean if a field has been set.

### GetOwnedBy

`func (o *ViewCalendarEvent) GetOwnedBy() int32`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *ViewCalendarEvent) GetOwnedByOk() (*int32, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *ViewCalendarEvent) SetOwnedBy(v int32)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *ViewCalendarEvent) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.

### GetOwnerUserId

`func (o *ViewCalendarEvent) GetOwnerUserId() int32`

GetOwnerUserId returns the OwnerUserId field if non-nil, zero value otherwise.

### GetOwnerUserIdOk

`func (o *ViewCalendarEvent) GetOwnerUserIdOk() (*int32, bool)`

GetOwnerUserIdOk returns a tuple with the OwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserId

`func (o *ViewCalendarEvent) SetOwnerUserId(v int32)`

SetOwnerUserId sets OwnerUserId field to given value.

### HasOwnerUserId

`func (o *ViewCalendarEvent) HasOwnerUserId() bool`

HasOwnerUserId returns a boolean if a field has been set.

### GetPrivacyType

`func (o *ViewCalendarEvent) GetPrivacyType() string`

GetPrivacyType returns the PrivacyType field if non-nil, zero value otherwise.

### GetPrivacyTypeOk

`func (o *ViewCalendarEvent) GetPrivacyTypeOk() (*string, bool)`

GetPrivacyTypeOk returns a tuple with the PrivacyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyType

`func (o *ViewCalendarEvent) SetPrivacyType(v string)`

SetPrivacyType sets PrivacyType field to given value.

### HasPrivacyType

`func (o *ViewCalendarEvent) HasPrivacyType() bool`

HasPrivacyType returns a boolean if a field has been set.

### GetProject

`func (o *ViewCalendarEvent) GetProject() ViewRelationship`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ViewCalendarEvent) GetProjectOk() (*ViewRelationship, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ViewCalendarEvent) SetProject(v ViewRelationship)`

SetProject sets Project field to given value.

### HasProject

`func (o *ViewCalendarEvent) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectId

`func (o *ViewCalendarEvent) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ViewCalendarEvent) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ViewCalendarEvent) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ViewCalendarEvent) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProjectUsersCanEdit

`func (o *ViewCalendarEvent) GetProjectUsersCanEdit() bool`

GetProjectUsersCanEdit returns the ProjectUsersCanEdit field if non-nil, zero value otherwise.

### GetProjectUsersCanEditOk

`func (o *ViewCalendarEvent) GetProjectUsersCanEditOk() (*bool, bool)`

GetProjectUsersCanEditOk returns a tuple with the ProjectUsersCanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectUsersCanEdit

`func (o *ViewCalendarEvent) SetProjectUsersCanEdit(v bool)`

SetProjectUsersCanEdit sets ProjectUsersCanEdit field to given value.

### HasProjectUsersCanEdit

`func (o *ViewCalendarEvent) HasProjectUsersCanEdit() bool`

HasProjectUsersCanEdit returns a boolean if a field has been set.

### GetSequence

`func (o *ViewCalendarEvent) GetSequence() ViewRelationship`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *ViewCalendarEvent) GetSequenceOk() (*ViewRelationship, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *ViewCalendarEvent) SetSequence(v ViewRelationship)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *ViewCalendarEvent) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSequenceId

`func (o *ViewCalendarEvent) GetSequenceId() int32`

GetSequenceId returns the SequenceId field if non-nil, zero value otherwise.

### GetSequenceIdOk

`func (o *ViewCalendarEvent) GetSequenceIdOk() (*int32, bool)`

GetSequenceIdOk returns a tuple with the SequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceId

`func (o *ViewCalendarEvent) SetSequenceId(v int32)`

SetSequenceId sets SequenceId field to given value.

### HasSequenceId

`func (o *ViewCalendarEvent) HasSequenceId() bool`

HasSequenceId returns a boolean if a field has been set.

### GetShowAsBusy

`func (o *ViewCalendarEvent) GetShowAsBusy() bool`

GetShowAsBusy returns the ShowAsBusy field if non-nil, zero value otherwise.

### GetShowAsBusyOk

`func (o *ViewCalendarEvent) GetShowAsBusyOk() (*bool, bool)`

GetShowAsBusyOk returns a tuple with the ShowAsBusy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAsBusy

`func (o *ViewCalendarEvent) SetShowAsBusy(v bool)`

SetShowAsBusy sets ShowAsBusy field to given value.

### HasShowAsBusy

`func (o *ViewCalendarEvent) HasShowAsBusy() bool`

HasShowAsBusy returns a boolean if a field has been set.

### GetStartDate

`func (o *ViewCalendarEvent) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ViewCalendarEvent) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ViewCalendarEvent) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ViewCalendarEvent) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTitle

`func (o *ViewCalendarEvent) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ViewCalendarEvent) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ViewCalendarEvent) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ViewCalendarEvent) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *ViewCalendarEvent) GetType() ViewRelationship`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ViewCalendarEvent) GetTypeOk() (*ViewRelationship, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ViewCalendarEvent) SetType(v ViewRelationship)`

SetType sets Type field to given value.

### HasType

`func (o *ViewCalendarEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypeId

`func (o *ViewCalendarEvent) GetTypeId() int32`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *ViewCalendarEvent) GetTypeIdOk() (*int32, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *ViewCalendarEvent) SetTypeId(v int32)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *ViewCalendarEvent) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ViewCalendarEvent) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ViewCalendarEvent) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ViewCalendarEvent) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ViewCalendarEvent) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ViewCalendarEvent) GetUpdatedBy() int32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ViewCalendarEvent) GetUpdatedByOk() (*int32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ViewCalendarEvent) SetUpdatedBy(v int32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ViewCalendarEvent) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedByUserId

`func (o *ViewCalendarEvent) GetUpdatedByUserId() int32`

GetUpdatedByUserId returns the UpdatedByUserId field if non-nil, zero value otherwise.

### GetUpdatedByUserIdOk

`func (o *ViewCalendarEvent) GetUpdatedByUserIdOk() (*int32, bool)`

GetUpdatedByUserIdOk returns a tuple with the UpdatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUserId

`func (o *ViewCalendarEvent) SetUpdatedByUserId(v int32)`

SetUpdatedByUserId sets UpdatedByUserId field to given value.

### HasUpdatedByUserId

`func (o *ViewCalendarEvent) HasUpdatedByUserId() bool`

HasUpdatedByUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


