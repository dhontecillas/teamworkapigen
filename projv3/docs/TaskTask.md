# TaskTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignees** | Pointer to [**PayloadUserGroups**](PayloadUserGroups.md) |  | [optional] 
**AttachmentIds** | Pointer to **[]int32** |  | [optional] 
**ChangeFollowers** | Pointer to [**PayloadUserGroups**](PayloadUserGroups.md) |  | [optional] 
**CommentFollowers** | Pointer to [**PayloadUserGroups**](PayloadUserGroups.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **int32** |  | [optional] 
**CrmDealIds** | Pointer to **[]int32** |  | [optional] 
**CustomFields** | Pointer to [**[]TaskCustomFieldTask**](TaskCustomFieldTask.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DueAt** | Pointer to **map[string]interface{}** | Date unmarshals represents a Unified API Spec date format. | [optional] 
**EstimatedMinutes** | Pointer to **int32** |  | [optional] 
**GrantAccessTo** | Pointer to [**PayloadUserGroups**](PayloadUserGroups.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ParentTaskId** | Pointer to **int32** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**Private** | Pointer to **bool** |  | [optional] 
**Progress** | Pointer to **int32** |  | [optional] 
**Reminders** | Pointer to [**[]TaskReminder**](TaskReminder.md) |  | [optional] 
**RepeatOptions** | Pointer to [**TaskRepeatOptions**](TaskRepeatOptions.md) |  | [optional] 
**StartAt** | Pointer to **map[string]interface{}** | Date unmarshals represents a Unified API Spec date format. | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TagIds** | Pointer to **[]int32** |  | [optional] 
**TasklistId** | Pointer to **int32** |  | [optional] 
**TemplateRoleName** | Pointer to **string** |  | [optional] 
**TicketId** | Pointer to **int32** |  | [optional] 

## Methods

### NewTaskTask

`func NewTaskTask() *TaskTask`

NewTaskTask instantiates a new TaskTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskTaskWithDefaults

`func NewTaskTaskWithDefaults() *TaskTask`

NewTaskTaskWithDefaults instantiates a new TaskTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignees

`func (o *TaskTask) GetAssignees() PayloadUserGroups`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *TaskTask) GetAssigneesOk() (*PayloadUserGroups, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *TaskTask) SetAssignees(v PayloadUserGroups)`

SetAssignees sets Assignees field to given value.

### HasAssignees

`func (o *TaskTask) HasAssignees() bool`

HasAssignees returns a boolean if a field has been set.

### GetAttachmentIds

`func (o *TaskTask) GetAttachmentIds() []int32`

GetAttachmentIds returns the AttachmentIds field if non-nil, zero value otherwise.

### GetAttachmentIdsOk

`func (o *TaskTask) GetAttachmentIdsOk() (*[]int32, bool)`

GetAttachmentIdsOk returns a tuple with the AttachmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentIds

`func (o *TaskTask) SetAttachmentIds(v []int32)`

SetAttachmentIds sets AttachmentIds field to given value.

### HasAttachmentIds

`func (o *TaskTask) HasAttachmentIds() bool`

HasAttachmentIds returns a boolean if a field has been set.

### GetChangeFollowers

`func (o *TaskTask) GetChangeFollowers() PayloadUserGroups`

GetChangeFollowers returns the ChangeFollowers field if non-nil, zero value otherwise.

### GetChangeFollowersOk

`func (o *TaskTask) GetChangeFollowersOk() (*PayloadUserGroups, bool)`

GetChangeFollowersOk returns a tuple with the ChangeFollowers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeFollowers

`func (o *TaskTask) SetChangeFollowers(v PayloadUserGroups)`

SetChangeFollowers sets ChangeFollowers field to given value.

### HasChangeFollowers

`func (o *TaskTask) HasChangeFollowers() bool`

HasChangeFollowers returns a boolean if a field has been set.

### GetCommentFollowers

`func (o *TaskTask) GetCommentFollowers() PayloadUserGroups`

GetCommentFollowers returns the CommentFollowers field if non-nil, zero value otherwise.

### GetCommentFollowersOk

`func (o *TaskTask) GetCommentFollowersOk() (*PayloadUserGroups, bool)`

GetCommentFollowersOk returns a tuple with the CommentFollowers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentFollowers

`func (o *TaskTask) SetCommentFollowers(v PayloadUserGroups)`

SetCommentFollowers sets CommentFollowers field to given value.

### HasCommentFollowers

`func (o *TaskTask) HasCommentFollowers() bool`

HasCommentFollowers returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TaskTask) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskTask) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskTask) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TaskTask) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *TaskTask) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *TaskTask) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *TaskTask) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *TaskTask) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCrmDealIds

`func (o *TaskTask) GetCrmDealIds() []int32`

GetCrmDealIds returns the CrmDealIds field if non-nil, zero value otherwise.

### GetCrmDealIdsOk

`func (o *TaskTask) GetCrmDealIdsOk() (*[]int32, bool)`

GetCrmDealIdsOk returns a tuple with the CrmDealIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrmDealIds

`func (o *TaskTask) SetCrmDealIds(v []int32)`

SetCrmDealIds sets CrmDealIds field to given value.

### HasCrmDealIds

`func (o *TaskTask) HasCrmDealIds() bool`

HasCrmDealIds returns a boolean if a field has been set.

### GetCustomFields

`func (o *TaskTask) GetCustomFields() []TaskCustomFieldTask`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *TaskTask) GetCustomFieldsOk() (*[]TaskCustomFieldTask, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *TaskTask) SetCustomFields(v []TaskCustomFieldTask)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *TaskTask) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetDescription

`func (o *TaskTask) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaskTask) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaskTask) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TaskTask) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDueAt

`func (o *TaskTask) GetDueAt() map[string]interface{}`

GetDueAt returns the DueAt field if non-nil, zero value otherwise.

### GetDueAtOk

`func (o *TaskTask) GetDueAtOk() (*map[string]interface{}, bool)`

GetDueAtOk returns a tuple with the DueAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueAt

`func (o *TaskTask) SetDueAt(v map[string]interface{})`

SetDueAt sets DueAt field to given value.

### HasDueAt

`func (o *TaskTask) HasDueAt() bool`

HasDueAt returns a boolean if a field has been set.

### GetEstimatedMinutes

`func (o *TaskTask) GetEstimatedMinutes() int32`

GetEstimatedMinutes returns the EstimatedMinutes field if non-nil, zero value otherwise.

### GetEstimatedMinutesOk

`func (o *TaskTask) GetEstimatedMinutesOk() (*int32, bool)`

GetEstimatedMinutesOk returns a tuple with the EstimatedMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedMinutes

`func (o *TaskTask) SetEstimatedMinutes(v int32)`

SetEstimatedMinutes sets EstimatedMinutes field to given value.

### HasEstimatedMinutes

`func (o *TaskTask) HasEstimatedMinutes() bool`

HasEstimatedMinutes returns a boolean if a field has been set.

### GetGrantAccessTo

`func (o *TaskTask) GetGrantAccessTo() PayloadUserGroups`

GetGrantAccessTo returns the GrantAccessTo field if non-nil, zero value otherwise.

### GetGrantAccessToOk

`func (o *TaskTask) GetGrantAccessToOk() (*PayloadUserGroups, bool)`

GetGrantAccessToOk returns a tuple with the GrantAccessTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantAccessTo

`func (o *TaskTask) SetGrantAccessTo(v PayloadUserGroups)`

SetGrantAccessTo sets GrantAccessTo field to given value.

### HasGrantAccessTo

`func (o *TaskTask) HasGrantAccessTo() bool`

HasGrantAccessTo returns a boolean if a field has been set.

### GetName

`func (o *TaskTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskTask) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaskTask) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentTaskId

`func (o *TaskTask) GetParentTaskId() int32`

GetParentTaskId returns the ParentTaskId field if non-nil, zero value otherwise.

### GetParentTaskIdOk

`func (o *TaskTask) GetParentTaskIdOk() (*int32, bool)`

GetParentTaskIdOk returns a tuple with the ParentTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTaskId

`func (o *TaskTask) SetParentTaskId(v int32)`

SetParentTaskId sets ParentTaskId field to given value.

### HasParentTaskId

`func (o *TaskTask) HasParentTaskId() bool`

HasParentTaskId returns a boolean if a field has been set.

### GetPriority

`func (o *TaskTask) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TaskTask) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TaskTask) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TaskTask) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPrivate

`func (o *TaskTask) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *TaskTask) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *TaskTask) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *TaskTask) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetProgress

`func (o *TaskTask) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *TaskTask) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *TaskTask) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *TaskTask) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetReminders

`func (o *TaskTask) GetReminders() []TaskReminder`

GetReminders returns the Reminders field if non-nil, zero value otherwise.

### GetRemindersOk

`func (o *TaskTask) GetRemindersOk() (*[]TaskReminder, bool)`

GetRemindersOk returns a tuple with the Reminders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminders

`func (o *TaskTask) SetReminders(v []TaskReminder)`

SetReminders sets Reminders field to given value.

### HasReminders

`func (o *TaskTask) HasReminders() bool`

HasReminders returns a boolean if a field has been set.

### GetRepeatOptions

`func (o *TaskTask) GetRepeatOptions() TaskRepeatOptions`

GetRepeatOptions returns the RepeatOptions field if non-nil, zero value otherwise.

### GetRepeatOptionsOk

`func (o *TaskTask) GetRepeatOptionsOk() (*TaskRepeatOptions, bool)`

GetRepeatOptionsOk returns a tuple with the RepeatOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatOptions

`func (o *TaskTask) SetRepeatOptions(v TaskRepeatOptions)`

SetRepeatOptions sets RepeatOptions field to given value.

### HasRepeatOptions

`func (o *TaskTask) HasRepeatOptions() bool`

HasRepeatOptions returns a boolean if a field has been set.

### GetStartAt

`func (o *TaskTask) GetStartAt() map[string]interface{}`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *TaskTask) GetStartAtOk() (*map[string]interface{}, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *TaskTask) SetStartAt(v map[string]interface{})`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *TaskTask) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### GetStatus

`func (o *TaskTask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskTask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskTask) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaskTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTagIds

`func (o *TaskTask) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *TaskTask) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *TaskTask) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *TaskTask) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTasklistId

`func (o *TaskTask) GetTasklistId() int32`

GetTasklistId returns the TasklistId field if non-nil, zero value otherwise.

### GetTasklistIdOk

`func (o *TaskTask) GetTasklistIdOk() (*int32, bool)`

GetTasklistIdOk returns a tuple with the TasklistId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasklistId

`func (o *TaskTask) SetTasklistId(v int32)`

SetTasklistId sets TasklistId field to given value.

### HasTasklistId

`func (o *TaskTask) HasTasklistId() bool`

HasTasklistId returns a boolean if a field has been set.

### GetTemplateRoleName

`func (o *TaskTask) GetTemplateRoleName() string`

GetTemplateRoleName returns the TemplateRoleName field if non-nil, zero value otherwise.

### GetTemplateRoleNameOk

`func (o *TaskTask) GetTemplateRoleNameOk() (*string, bool)`

GetTemplateRoleNameOk returns a tuple with the TemplateRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateRoleName

`func (o *TaskTask) SetTemplateRoleName(v string)`

SetTemplateRoleName sets TemplateRoleName field to given value.

### HasTemplateRoleName

`func (o *TaskTask) HasTemplateRoleName() bool`

HasTemplateRoleName returns a boolean if a field has been set.

### GetTicketId

`func (o *TaskTask) GetTicketId() int32`

GetTicketId returns the TicketId field if non-nil, zero value otherwise.

### GetTicketIdOk

`func (o *TaskTask) GetTicketIdOk() (*int32, bool)`

GetTicketIdOk returns a tuple with the TicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketId

`func (o *TaskTask) SetTicketId(v int32)`

SetTicketId sets TicketId field to given value.

### HasTicketId

`func (o *TaskTask) HasTicketId() bool`

HasTicketId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


