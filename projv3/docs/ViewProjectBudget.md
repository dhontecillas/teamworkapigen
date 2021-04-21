# ViewProjectBudget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to **int32** |  | [optional] 
**CapacityUsed** | Pointer to **int32** |  | [optional] 
**CompletedAt** | Pointer to **string** |  | [optional] 
**CompletedBy** | Pointer to **int32** |  | [optional] 
**CompletedByUserId** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **int32** |  | [optional] 
**CreatedByUserId** | Pointer to **int32** |  | [optional] 
**CurrencyCode** | Pointer to **string** |  | [optional] 
**DateCompleted** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **string** |  | [optional] 
**DateDeleted** | Pointer to **string** |  | [optional] 
**DateUpdated** | Pointer to **string** |  | [optional] 
**DefaultRate** | Pointer to **float32** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**DeletedBy** | Pointer to **int32** |  | [optional] 
**DeletedByUserId** | Pointer to **int32** |  | [optional] 
**EndDateTime** | Pointer to **string** |  | [optional] 
**ExpenseType** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**IsRepeating** | Pointer to **bool** |  | [optional] 
**NotificationIds** | Pointer to **[]int32** |  | [optional] 
**Notifications** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**OriginatorBudget** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**OriginatorBudgetId** | Pointer to **int32** |  | [optional] 
**Project** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**RepeatPeriod** | Pointer to **int32** |  | [optional] 
**RepeatUnit** | Pointer to **string** |  | [optional] 
**RepeatsRemaining** | Pointer to **int32** |  | [optional] 
**SequenceNumber** | Pointer to **int32** |  | [optional] 
**StartDateTime** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TimelogType** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **int32** |  | [optional] 
**UpdatedByUserId** | Pointer to **int32** |  | [optional] 

## Methods

### NewViewProjectBudget

`func NewViewProjectBudget() *ViewProjectBudget`

NewViewProjectBudget instantiates a new ViewProjectBudget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewProjectBudgetWithDefaults

`func NewViewProjectBudgetWithDefaults() *ViewProjectBudget`

NewViewProjectBudgetWithDefaults instantiates a new ViewProjectBudget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *ViewProjectBudget) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *ViewProjectBudget) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *ViewProjectBudget) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *ViewProjectBudget) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetCapacityUsed

`func (o *ViewProjectBudget) GetCapacityUsed() int32`

GetCapacityUsed returns the CapacityUsed field if non-nil, zero value otherwise.

### GetCapacityUsedOk

`func (o *ViewProjectBudget) GetCapacityUsedOk() (*int32, bool)`

GetCapacityUsedOk returns a tuple with the CapacityUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityUsed

`func (o *ViewProjectBudget) SetCapacityUsed(v int32)`

SetCapacityUsed sets CapacityUsed field to given value.

### HasCapacityUsed

`func (o *ViewProjectBudget) HasCapacityUsed() bool`

HasCapacityUsed returns a boolean if a field has been set.

### GetCompletedAt

`func (o *ViewProjectBudget) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *ViewProjectBudget) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *ViewProjectBudget) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *ViewProjectBudget) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetCompletedBy

`func (o *ViewProjectBudget) GetCompletedBy() int32`

GetCompletedBy returns the CompletedBy field if non-nil, zero value otherwise.

### GetCompletedByOk

`func (o *ViewProjectBudget) GetCompletedByOk() (*int32, bool)`

GetCompletedByOk returns a tuple with the CompletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedBy

`func (o *ViewProjectBudget) SetCompletedBy(v int32)`

SetCompletedBy sets CompletedBy field to given value.

### HasCompletedBy

`func (o *ViewProjectBudget) HasCompletedBy() bool`

HasCompletedBy returns a boolean if a field has been set.

### GetCompletedByUserId

`func (o *ViewProjectBudget) GetCompletedByUserId() int32`

GetCompletedByUserId returns the CompletedByUserId field if non-nil, zero value otherwise.

### GetCompletedByUserIdOk

`func (o *ViewProjectBudget) GetCompletedByUserIdOk() (*int32, bool)`

GetCompletedByUserIdOk returns a tuple with the CompletedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedByUserId

`func (o *ViewProjectBudget) SetCompletedByUserId(v int32)`

SetCompletedByUserId sets CompletedByUserId field to given value.

### HasCompletedByUserId

`func (o *ViewProjectBudget) HasCompletedByUserId() bool`

HasCompletedByUserId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ViewProjectBudget) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ViewProjectBudget) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ViewProjectBudget) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ViewProjectBudget) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ViewProjectBudget) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ViewProjectBudget) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ViewProjectBudget) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ViewProjectBudget) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *ViewProjectBudget) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *ViewProjectBudget) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *ViewProjectBudget) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *ViewProjectBudget) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *ViewProjectBudget) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *ViewProjectBudget) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *ViewProjectBudget) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *ViewProjectBudget) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetDateCompleted

`func (o *ViewProjectBudget) GetDateCompleted() string`

GetDateCompleted returns the DateCompleted field if non-nil, zero value otherwise.

### GetDateCompletedOk

`func (o *ViewProjectBudget) GetDateCompletedOk() (*string, bool)`

GetDateCompletedOk returns a tuple with the DateCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCompleted

`func (o *ViewProjectBudget) SetDateCompleted(v string)`

SetDateCompleted sets DateCompleted field to given value.

### HasDateCompleted

`func (o *ViewProjectBudget) HasDateCompleted() bool`

HasDateCompleted returns a boolean if a field has been set.

### GetDateCreated

`func (o *ViewProjectBudget) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ViewProjectBudget) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ViewProjectBudget) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ViewProjectBudget) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateDeleted

`func (o *ViewProjectBudget) GetDateDeleted() string`

GetDateDeleted returns the DateDeleted field if non-nil, zero value otherwise.

### GetDateDeletedOk

`func (o *ViewProjectBudget) GetDateDeletedOk() (*string, bool)`

GetDateDeletedOk returns a tuple with the DateDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeleted

`func (o *ViewProjectBudget) SetDateDeleted(v string)`

SetDateDeleted sets DateDeleted field to given value.

### HasDateDeleted

`func (o *ViewProjectBudget) HasDateDeleted() bool`

HasDateDeleted returns a boolean if a field has been set.

### GetDateUpdated

`func (o *ViewProjectBudget) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ViewProjectBudget) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ViewProjectBudget) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ViewProjectBudget) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### GetDefaultRate

`func (o *ViewProjectBudget) GetDefaultRate() float32`

GetDefaultRate returns the DefaultRate field if non-nil, zero value otherwise.

### GetDefaultRateOk

`func (o *ViewProjectBudget) GetDefaultRateOk() (*float32, bool)`

GetDefaultRateOk returns a tuple with the DefaultRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRate

`func (o *ViewProjectBudget) SetDefaultRate(v float32)`

SetDefaultRate sets DefaultRate field to given value.

### HasDefaultRate

`func (o *ViewProjectBudget) HasDefaultRate() bool`

HasDefaultRate returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ViewProjectBudget) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ViewProjectBudget) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ViewProjectBudget) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ViewProjectBudget) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeletedBy

`func (o *ViewProjectBudget) GetDeletedBy() int32`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *ViewProjectBudget) GetDeletedByOk() (*int32, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *ViewProjectBudget) SetDeletedBy(v int32)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *ViewProjectBudget) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetDeletedByUserId

`func (o *ViewProjectBudget) GetDeletedByUserId() int32`

GetDeletedByUserId returns the DeletedByUserId field if non-nil, zero value otherwise.

### GetDeletedByUserIdOk

`func (o *ViewProjectBudget) GetDeletedByUserIdOk() (*int32, bool)`

GetDeletedByUserIdOk returns a tuple with the DeletedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedByUserId

`func (o *ViewProjectBudget) SetDeletedByUserId(v int32)`

SetDeletedByUserId sets DeletedByUserId field to given value.

### HasDeletedByUserId

`func (o *ViewProjectBudget) HasDeletedByUserId() bool`

HasDeletedByUserId returns a boolean if a field has been set.

### GetEndDateTime

`func (o *ViewProjectBudget) GetEndDateTime() string`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *ViewProjectBudget) GetEndDateTimeOk() (*string, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *ViewProjectBudget) SetEndDateTime(v string)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *ViewProjectBudget) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### GetExpenseType

`func (o *ViewProjectBudget) GetExpenseType() string`

GetExpenseType returns the ExpenseType field if non-nil, zero value otherwise.

### GetExpenseTypeOk

`func (o *ViewProjectBudget) GetExpenseTypeOk() (*string, bool)`

GetExpenseTypeOk returns a tuple with the ExpenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseType

`func (o *ViewProjectBudget) SetExpenseType(v string)`

SetExpenseType sets ExpenseType field to given value.

### HasExpenseType

`func (o *ViewProjectBudget) HasExpenseType() bool`

HasExpenseType returns a boolean if a field has been set.

### GetId

`func (o *ViewProjectBudget) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewProjectBudget) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewProjectBudget) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewProjectBudget) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsRepeating

`func (o *ViewProjectBudget) GetIsRepeating() bool`

GetIsRepeating returns the IsRepeating field if non-nil, zero value otherwise.

### GetIsRepeatingOk

`func (o *ViewProjectBudget) GetIsRepeatingOk() (*bool, bool)`

GetIsRepeatingOk returns a tuple with the IsRepeating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRepeating

`func (o *ViewProjectBudget) SetIsRepeating(v bool)`

SetIsRepeating sets IsRepeating field to given value.

### HasIsRepeating

`func (o *ViewProjectBudget) HasIsRepeating() bool`

HasIsRepeating returns a boolean if a field has been set.

### GetNotificationIds

`func (o *ViewProjectBudget) GetNotificationIds() []int32`

GetNotificationIds returns the NotificationIds field if non-nil, zero value otherwise.

### GetNotificationIdsOk

`func (o *ViewProjectBudget) GetNotificationIdsOk() (*[]int32, bool)`

GetNotificationIdsOk returns a tuple with the NotificationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationIds

`func (o *ViewProjectBudget) SetNotificationIds(v []int32)`

SetNotificationIds sets NotificationIds field to given value.

### HasNotificationIds

`func (o *ViewProjectBudget) HasNotificationIds() bool`

HasNotificationIds returns a boolean if a field has been set.

### GetNotifications

`func (o *ViewProjectBudget) GetNotifications() []ViewRelationship`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *ViewProjectBudget) GetNotificationsOk() (*[]ViewRelationship, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *ViewProjectBudget) SetNotifications(v []ViewRelationship)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *ViewProjectBudget) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetOriginatorBudget

`func (o *ViewProjectBudget) GetOriginatorBudget() ViewRelationship`

GetOriginatorBudget returns the OriginatorBudget field if non-nil, zero value otherwise.

### GetOriginatorBudgetOk

`func (o *ViewProjectBudget) GetOriginatorBudgetOk() (*ViewRelationship, bool)`

GetOriginatorBudgetOk returns a tuple with the OriginatorBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatorBudget

`func (o *ViewProjectBudget) SetOriginatorBudget(v ViewRelationship)`

SetOriginatorBudget sets OriginatorBudget field to given value.

### HasOriginatorBudget

`func (o *ViewProjectBudget) HasOriginatorBudget() bool`

HasOriginatorBudget returns a boolean if a field has been set.

### GetOriginatorBudgetId

`func (o *ViewProjectBudget) GetOriginatorBudgetId() int32`

GetOriginatorBudgetId returns the OriginatorBudgetId field if non-nil, zero value otherwise.

### GetOriginatorBudgetIdOk

`func (o *ViewProjectBudget) GetOriginatorBudgetIdOk() (*int32, bool)`

GetOriginatorBudgetIdOk returns a tuple with the OriginatorBudgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatorBudgetId

`func (o *ViewProjectBudget) SetOriginatorBudgetId(v int32)`

SetOriginatorBudgetId sets OriginatorBudgetId field to given value.

### HasOriginatorBudgetId

`func (o *ViewProjectBudget) HasOriginatorBudgetId() bool`

HasOriginatorBudgetId returns a boolean if a field has been set.

### GetProject

`func (o *ViewProjectBudget) GetProject() ViewRelationship`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ViewProjectBudget) GetProjectOk() (*ViewRelationship, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ViewProjectBudget) SetProject(v ViewRelationship)`

SetProject sets Project field to given value.

### HasProject

`func (o *ViewProjectBudget) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectId

`func (o *ViewProjectBudget) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ViewProjectBudget) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ViewProjectBudget) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ViewProjectBudget) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetRepeatPeriod

`func (o *ViewProjectBudget) GetRepeatPeriod() int32`

GetRepeatPeriod returns the RepeatPeriod field if non-nil, zero value otherwise.

### GetRepeatPeriodOk

`func (o *ViewProjectBudget) GetRepeatPeriodOk() (*int32, bool)`

GetRepeatPeriodOk returns a tuple with the RepeatPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatPeriod

`func (o *ViewProjectBudget) SetRepeatPeriod(v int32)`

SetRepeatPeriod sets RepeatPeriod field to given value.

### HasRepeatPeriod

`func (o *ViewProjectBudget) HasRepeatPeriod() bool`

HasRepeatPeriod returns a boolean if a field has been set.

### GetRepeatUnit

`func (o *ViewProjectBudget) GetRepeatUnit() string`

GetRepeatUnit returns the RepeatUnit field if non-nil, zero value otherwise.

### GetRepeatUnitOk

`func (o *ViewProjectBudget) GetRepeatUnitOk() (*string, bool)`

GetRepeatUnitOk returns a tuple with the RepeatUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatUnit

`func (o *ViewProjectBudget) SetRepeatUnit(v string)`

SetRepeatUnit sets RepeatUnit field to given value.

### HasRepeatUnit

`func (o *ViewProjectBudget) HasRepeatUnit() bool`

HasRepeatUnit returns a boolean if a field has been set.

### GetRepeatsRemaining

`func (o *ViewProjectBudget) GetRepeatsRemaining() int32`

GetRepeatsRemaining returns the RepeatsRemaining field if non-nil, zero value otherwise.

### GetRepeatsRemainingOk

`func (o *ViewProjectBudget) GetRepeatsRemainingOk() (*int32, bool)`

GetRepeatsRemainingOk returns a tuple with the RepeatsRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatsRemaining

`func (o *ViewProjectBudget) SetRepeatsRemaining(v int32)`

SetRepeatsRemaining sets RepeatsRemaining field to given value.

### HasRepeatsRemaining

`func (o *ViewProjectBudget) HasRepeatsRemaining() bool`

HasRepeatsRemaining returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *ViewProjectBudget) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *ViewProjectBudget) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *ViewProjectBudget) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *ViewProjectBudget) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetStartDateTime

`func (o *ViewProjectBudget) GetStartDateTime() string`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *ViewProjectBudget) GetStartDateTimeOk() (*string, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *ViewProjectBudget) SetStartDateTime(v string)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *ViewProjectBudget) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetStatus

`func (o *ViewProjectBudget) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ViewProjectBudget) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ViewProjectBudget) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ViewProjectBudget) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimelogType

`func (o *ViewProjectBudget) GetTimelogType() string`

GetTimelogType returns the TimelogType field if non-nil, zero value otherwise.

### GetTimelogTypeOk

`func (o *ViewProjectBudget) GetTimelogTypeOk() (*string, bool)`

GetTimelogTypeOk returns a tuple with the TimelogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelogType

`func (o *ViewProjectBudget) SetTimelogType(v string)`

SetTimelogType sets TimelogType field to given value.

### HasTimelogType

`func (o *ViewProjectBudget) HasTimelogType() bool`

HasTimelogType returns a boolean if a field has been set.

### GetType

`func (o *ViewProjectBudget) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ViewProjectBudget) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ViewProjectBudget) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ViewProjectBudget) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ViewProjectBudget) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ViewProjectBudget) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ViewProjectBudget) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ViewProjectBudget) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ViewProjectBudget) GetUpdatedBy() int32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ViewProjectBudget) GetUpdatedByOk() (*int32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ViewProjectBudget) SetUpdatedBy(v int32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ViewProjectBudget) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedByUserId

`func (o *ViewProjectBudget) GetUpdatedByUserId() int32`

GetUpdatedByUserId returns the UpdatedByUserId field if non-nil, zero value otherwise.

### GetUpdatedByUserIdOk

`func (o *ViewProjectBudget) GetUpdatedByUserIdOk() (*int32, bool)`

GetUpdatedByUserIdOk returns a tuple with the UpdatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUserId

`func (o *ViewProjectBudget) SetUpdatedByUserId(v int32)`

SetUpdatedByUserId sets UpdatedByUserId field to given value.

### HasUpdatedByUserId

`func (o *ViewProjectBudget) HasUpdatedByUserId() bool`

HasUpdatedByUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


