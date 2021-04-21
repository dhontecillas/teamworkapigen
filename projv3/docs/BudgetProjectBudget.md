# BudgetProjectBudget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to **int32** |  | [optional] 
**CurrencyCode** | Pointer to **string** |  | [optional] 
**DefaultRate** | Pointer to **float32** |  | [optional] 
**EndDateTime** | Pointer to **string** |  | [optional] 
**ExpenseType** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**IsRepeating** | Pointer to **bool** |  | [optional] 
**Notifications** | Pointer to [**[]NotificationProjectBudgetNotification**](NotificationProjectBudgetNotification.md) |  | [optional] 
**OriginatorBudgetID** | Pointer to **int32** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**RepeatPeriod** | Pointer to **int32** |  | [optional] 
**RepeatUnit** | Pointer to **string** |  | [optional] 
**RepeatsRemaining** | Pointer to **int32** |  | [optional] 
**SequenceNumber** | Pointer to **int32** |  | [optional] 
**StartDateTime** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TimelogType** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewBudgetProjectBudget

`func NewBudgetProjectBudget() *BudgetProjectBudget`

NewBudgetProjectBudget instantiates a new BudgetProjectBudget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetProjectBudgetWithDefaults

`func NewBudgetProjectBudgetWithDefaults() *BudgetProjectBudget`

NewBudgetProjectBudgetWithDefaults instantiates a new BudgetProjectBudget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *BudgetProjectBudget) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *BudgetProjectBudget) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *BudgetProjectBudget) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *BudgetProjectBudget) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *BudgetProjectBudget) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *BudgetProjectBudget) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *BudgetProjectBudget) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *BudgetProjectBudget) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetDefaultRate

`func (o *BudgetProjectBudget) GetDefaultRate() float32`

GetDefaultRate returns the DefaultRate field if non-nil, zero value otherwise.

### GetDefaultRateOk

`func (o *BudgetProjectBudget) GetDefaultRateOk() (*float32, bool)`

GetDefaultRateOk returns a tuple with the DefaultRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRate

`func (o *BudgetProjectBudget) SetDefaultRate(v float32)`

SetDefaultRate sets DefaultRate field to given value.

### HasDefaultRate

`func (o *BudgetProjectBudget) HasDefaultRate() bool`

HasDefaultRate returns a boolean if a field has been set.

### GetEndDateTime

`func (o *BudgetProjectBudget) GetEndDateTime() string`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *BudgetProjectBudget) GetEndDateTimeOk() (*string, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *BudgetProjectBudget) SetEndDateTime(v string)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *BudgetProjectBudget) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### GetExpenseType

`func (o *BudgetProjectBudget) GetExpenseType() string`

GetExpenseType returns the ExpenseType field if non-nil, zero value otherwise.

### GetExpenseTypeOk

`func (o *BudgetProjectBudget) GetExpenseTypeOk() (*string, bool)`

GetExpenseTypeOk returns a tuple with the ExpenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseType

`func (o *BudgetProjectBudget) SetExpenseType(v string)`

SetExpenseType sets ExpenseType field to given value.

### HasExpenseType

`func (o *BudgetProjectBudget) HasExpenseType() bool`

HasExpenseType returns a boolean if a field has been set.

### GetId

`func (o *BudgetProjectBudget) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BudgetProjectBudget) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BudgetProjectBudget) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BudgetProjectBudget) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsRepeating

`func (o *BudgetProjectBudget) GetIsRepeating() bool`

GetIsRepeating returns the IsRepeating field if non-nil, zero value otherwise.

### GetIsRepeatingOk

`func (o *BudgetProjectBudget) GetIsRepeatingOk() (*bool, bool)`

GetIsRepeatingOk returns a tuple with the IsRepeating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRepeating

`func (o *BudgetProjectBudget) SetIsRepeating(v bool)`

SetIsRepeating sets IsRepeating field to given value.

### HasIsRepeating

`func (o *BudgetProjectBudget) HasIsRepeating() bool`

HasIsRepeating returns a boolean if a field has been set.

### GetNotifications

`func (o *BudgetProjectBudget) GetNotifications() []NotificationProjectBudgetNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *BudgetProjectBudget) GetNotificationsOk() (*[]NotificationProjectBudgetNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *BudgetProjectBudget) SetNotifications(v []NotificationProjectBudgetNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *BudgetProjectBudget) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetOriginatorBudgetID

`func (o *BudgetProjectBudget) GetOriginatorBudgetID() int32`

GetOriginatorBudgetID returns the OriginatorBudgetID field if non-nil, zero value otherwise.

### GetOriginatorBudgetIDOk

`func (o *BudgetProjectBudget) GetOriginatorBudgetIDOk() (*int32, bool)`

GetOriginatorBudgetIDOk returns a tuple with the OriginatorBudgetID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatorBudgetID

`func (o *BudgetProjectBudget) SetOriginatorBudgetID(v int32)`

SetOriginatorBudgetID sets OriginatorBudgetID field to given value.

### HasOriginatorBudgetID

`func (o *BudgetProjectBudget) HasOriginatorBudgetID() bool`

HasOriginatorBudgetID returns a boolean if a field has been set.

### GetProjectId

`func (o *BudgetProjectBudget) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BudgetProjectBudget) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BudgetProjectBudget) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BudgetProjectBudget) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetRepeatPeriod

`func (o *BudgetProjectBudget) GetRepeatPeriod() int32`

GetRepeatPeriod returns the RepeatPeriod field if non-nil, zero value otherwise.

### GetRepeatPeriodOk

`func (o *BudgetProjectBudget) GetRepeatPeriodOk() (*int32, bool)`

GetRepeatPeriodOk returns a tuple with the RepeatPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatPeriod

`func (o *BudgetProjectBudget) SetRepeatPeriod(v int32)`

SetRepeatPeriod sets RepeatPeriod field to given value.

### HasRepeatPeriod

`func (o *BudgetProjectBudget) HasRepeatPeriod() bool`

HasRepeatPeriod returns a boolean if a field has been set.

### GetRepeatUnit

`func (o *BudgetProjectBudget) GetRepeatUnit() string`

GetRepeatUnit returns the RepeatUnit field if non-nil, zero value otherwise.

### GetRepeatUnitOk

`func (o *BudgetProjectBudget) GetRepeatUnitOk() (*string, bool)`

GetRepeatUnitOk returns a tuple with the RepeatUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatUnit

`func (o *BudgetProjectBudget) SetRepeatUnit(v string)`

SetRepeatUnit sets RepeatUnit field to given value.

### HasRepeatUnit

`func (o *BudgetProjectBudget) HasRepeatUnit() bool`

HasRepeatUnit returns a boolean if a field has been set.

### GetRepeatsRemaining

`func (o *BudgetProjectBudget) GetRepeatsRemaining() int32`

GetRepeatsRemaining returns the RepeatsRemaining field if non-nil, zero value otherwise.

### GetRepeatsRemainingOk

`func (o *BudgetProjectBudget) GetRepeatsRemainingOk() (*int32, bool)`

GetRepeatsRemainingOk returns a tuple with the RepeatsRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatsRemaining

`func (o *BudgetProjectBudget) SetRepeatsRemaining(v int32)`

SetRepeatsRemaining sets RepeatsRemaining field to given value.

### HasRepeatsRemaining

`func (o *BudgetProjectBudget) HasRepeatsRemaining() bool`

HasRepeatsRemaining returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *BudgetProjectBudget) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *BudgetProjectBudget) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *BudgetProjectBudget) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *BudgetProjectBudget) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetStartDateTime

`func (o *BudgetProjectBudget) GetStartDateTime() string`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *BudgetProjectBudget) GetStartDateTimeOk() (*string, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *BudgetProjectBudget) SetStartDateTime(v string)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *BudgetProjectBudget) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetStatus

`func (o *BudgetProjectBudget) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BudgetProjectBudget) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BudgetProjectBudget) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BudgetProjectBudget) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimelogType

`func (o *BudgetProjectBudget) GetTimelogType() string`

GetTimelogType returns the TimelogType field if non-nil, zero value otherwise.

### GetTimelogTypeOk

`func (o *BudgetProjectBudget) GetTimelogTypeOk() (*string, bool)`

GetTimelogTypeOk returns a tuple with the TimelogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelogType

`func (o *BudgetProjectBudget) SetTimelogType(v string)`

SetTimelogType sets TimelogType field to given value.

### HasTimelogType

`func (o *BudgetProjectBudget) HasTimelogType() bool`

HasTimelogType returns a boolean if a field has been set.

### GetType

`func (o *BudgetProjectBudget) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BudgetProjectBudget) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BudgetProjectBudget) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BudgetProjectBudget) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


