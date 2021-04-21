# TaskReminder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsRelative** | Pointer to **bool** |  | [optional] 
**Note** | Pointer to **string** |  | [optional] 
**PeopleAssigned** | Pointer to **bool** |  | [optional] 
**RelativeNumberDays** | Pointer to **int32** |  | [optional] 
**RemindAt** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 
**UsingOffsetDueDate** | Pointer to **bool** |  | [optional] 

## Methods

### NewTaskReminder

`func NewTaskReminder() *TaskReminder`

NewTaskReminder instantiates a new TaskReminder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskReminderWithDefaults

`func NewTaskReminderWithDefaults() *TaskReminder`

NewTaskReminderWithDefaults instantiates a new TaskReminder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsRelative

`func (o *TaskReminder) GetIsRelative() bool`

GetIsRelative returns the IsRelative field if non-nil, zero value otherwise.

### GetIsRelativeOk

`func (o *TaskReminder) GetIsRelativeOk() (*bool, bool)`

GetIsRelativeOk returns a tuple with the IsRelative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRelative

`func (o *TaskReminder) SetIsRelative(v bool)`

SetIsRelative sets IsRelative field to given value.

### HasIsRelative

`func (o *TaskReminder) HasIsRelative() bool`

HasIsRelative returns a boolean if a field has been set.

### GetNote

`func (o *TaskReminder) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *TaskReminder) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *TaskReminder) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *TaskReminder) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetPeopleAssigned

`func (o *TaskReminder) GetPeopleAssigned() bool`

GetPeopleAssigned returns the PeopleAssigned field if non-nil, zero value otherwise.

### GetPeopleAssignedOk

`func (o *TaskReminder) GetPeopleAssignedOk() (*bool, bool)`

GetPeopleAssignedOk returns a tuple with the PeopleAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeopleAssigned

`func (o *TaskReminder) SetPeopleAssigned(v bool)`

SetPeopleAssigned sets PeopleAssigned field to given value.

### HasPeopleAssigned

`func (o *TaskReminder) HasPeopleAssigned() bool`

HasPeopleAssigned returns a boolean if a field has been set.

### GetRelativeNumberDays

`func (o *TaskReminder) GetRelativeNumberDays() int32`

GetRelativeNumberDays returns the RelativeNumberDays field if non-nil, zero value otherwise.

### GetRelativeNumberDaysOk

`func (o *TaskReminder) GetRelativeNumberDaysOk() (*int32, bool)`

GetRelativeNumberDaysOk returns a tuple with the RelativeNumberDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeNumberDays

`func (o *TaskReminder) SetRelativeNumberDays(v int32)`

SetRelativeNumberDays sets RelativeNumberDays field to given value.

### HasRelativeNumberDays

`func (o *TaskReminder) HasRelativeNumberDays() bool`

HasRelativeNumberDays returns a boolean if a field has been set.

### GetRemindAt

`func (o *TaskReminder) GetRemindAt() string`

GetRemindAt returns the RemindAt field if non-nil, zero value otherwise.

### GetRemindAtOk

`func (o *TaskReminder) GetRemindAtOk() (*string, bool)`

GetRemindAtOk returns a tuple with the RemindAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemindAt

`func (o *TaskReminder) SetRemindAt(v string)`

SetRemindAt sets RemindAt field to given value.

### HasRemindAt

`func (o *TaskReminder) HasRemindAt() bool`

HasRemindAt returns a boolean if a field has been set.

### GetType

`func (o *TaskReminder) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaskReminder) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaskReminder) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TaskReminder) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserId

`func (o *TaskReminder) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TaskReminder) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TaskReminder) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *TaskReminder) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUsingOffsetDueDate

`func (o *TaskReminder) GetUsingOffsetDueDate() bool`

GetUsingOffsetDueDate returns the UsingOffsetDueDate field if non-nil, zero value otherwise.

### GetUsingOffsetDueDateOk

`func (o *TaskReminder) GetUsingOffsetDueDateOk() (*bool, bool)`

GetUsingOffsetDueDateOk returns a tuple with the UsingOffsetDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsingOffsetDueDate

`func (o *TaskReminder) SetUsingOffsetDueDate(v bool)`

SetUsingOffsetDueDate sets UsingOffsetDueDate field to given value.

### HasUsingOffsetDueDate

`func (o *TaskReminder) HasUsingOffsetDueDate() bool`

HasUsingOffsetDueDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


