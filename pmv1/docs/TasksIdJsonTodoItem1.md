# TasksIdJsonTodoItem1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | Pointer to **[]string** | Existing files can be attached to the task by specifying a comma-separated list of File IDs. | [optional] 
**ChangeFollowerIds** | Pointer to **string** | id&#39;s of people which will be followers. Change id means if notify&#x3D;true they will get notified if any changes occur on that task. | [optional] 
**ColumnId** | Pointer to **string** | Used to put a new Task directly in to a Column in the Boards view | [optional] 
**CommentFollowerIds** | Pointer to **string** | id&#39;s of people which will be followers if someone creates a comment. Comment id means if notify&#x3D;true they will get notified if any comments are created on that task. | [optional] 
**Content** | **string** | The name of the task you are adding. | 
**Description** | Pointer to **string** | Tasks can be assigned a description. | [optional] 
**DueDate** | Pointer to **string** | Tasks can be assigned a date for when they are due to be finished. The format should be YYYYMMDD. | [optional] 
**EstimatedMinutes** | Pointer to **string** | Set an estimated number of minutes for a task to be completed by setting this parameter. | [optional] 
**GrantAccessTo** | Pointer to **string** | Used in relation to private. If you set a Task as private (1), you can optionally add one or more User IDs to make the task visible to | [optional] 
**Notify** | Pointer to **bool** | This parameter can be set to true to notify people assigned to this task by email. | [optional] 
**ParentTaskId** | Pointer to **string** | Set this to the ID of a parent task if you wish to make your task a subtask. | [optional] 
**PendingFileAttachments** | Pointer to **string** | New files can be attached using this option (see the uploading files section for more info). This is a comma-separated list of PendingFileRef&#39;s. | [optional] 
**PositionAfterTask** | Pointer to **int32** | A task can be placed after another task by setting this parameter to a task ID. Additionally, some other options are available: -2 - Ignore -1 - Place task at the top of the list 0 - Place task at the bottom of the list | [optional] 
**Predecessors** | Pointer to [**[]TasksIdJsonTodoItem1Predecessors**](TasksIdJsonTodoItem1Predecessors.md) | An array of predecessor objects can be passed to specify which tasks need to have a specified state before this task can be marked as completed. Each predecessor object should contain two keys: id: The ID of the task that predecesses this one. type: The state the task needs to be, can be \&quot;complete\&quot; or \&quot;start\&quot;. | [optional] 
**Priority** | Pointer to **string** | Tasks can be assigned the following priorities: - not set - low - medium - high | [optional] 
**Private** | Pointer to **int32** | Set to 1 to make the task Private. Setting a 0 will set the Task back to normal | [optional] 
**Progress** | Pointer to **string** | Set the progress from 0 to 90 | [optional] 
**ResponsiblePartyId** | Pointer to **int32** | This can be used to assign the new task to a person or group of people. For example: -1 would assign the task to everyone 32 would assign the task to user 32. 32,55 would assign the task to users 32 and 55 etc. | [optional] 
**StartDate** | Pointer to **string** | Tasks can be assigned a date to start on. The format should be YYYYMMDD. | [optional] 
**Tags** | Pointer to **string** | A comma separated list of tags for the task | [optional] 

## Methods

### NewTasksIdJsonTodoItem1

`func NewTasksIdJsonTodoItem1(content string, ) *TasksIdJsonTodoItem1`

NewTasksIdJsonTodoItem1 instantiates a new TasksIdJsonTodoItem1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTasksIdJsonTodoItem1WithDefaults

`func NewTasksIdJsonTodoItem1WithDefaults() *TasksIdJsonTodoItem1`

NewTasksIdJsonTodoItem1WithDefaults instantiates a new TasksIdJsonTodoItem1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *TasksIdJsonTodoItem1) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TasksIdJsonTodoItem1) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TasksIdJsonTodoItem1) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *TasksIdJsonTodoItem1) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetChangeFollowerIds

`func (o *TasksIdJsonTodoItem1) GetChangeFollowerIds() string`

GetChangeFollowerIds returns the ChangeFollowerIds field if non-nil, zero value otherwise.

### GetChangeFollowerIdsOk

`func (o *TasksIdJsonTodoItem1) GetChangeFollowerIdsOk() (*string, bool)`

GetChangeFollowerIdsOk returns a tuple with the ChangeFollowerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeFollowerIds

`func (o *TasksIdJsonTodoItem1) SetChangeFollowerIds(v string)`

SetChangeFollowerIds sets ChangeFollowerIds field to given value.

### HasChangeFollowerIds

`func (o *TasksIdJsonTodoItem1) HasChangeFollowerIds() bool`

HasChangeFollowerIds returns a boolean if a field has been set.

### GetColumnId

`func (o *TasksIdJsonTodoItem1) GetColumnId() string`

GetColumnId returns the ColumnId field if non-nil, zero value otherwise.

### GetColumnIdOk

`func (o *TasksIdJsonTodoItem1) GetColumnIdOk() (*string, bool)`

GetColumnIdOk returns a tuple with the ColumnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnId

`func (o *TasksIdJsonTodoItem1) SetColumnId(v string)`

SetColumnId sets ColumnId field to given value.

### HasColumnId

`func (o *TasksIdJsonTodoItem1) HasColumnId() bool`

HasColumnId returns a boolean if a field has been set.

### GetCommentFollowerIds

`func (o *TasksIdJsonTodoItem1) GetCommentFollowerIds() string`

GetCommentFollowerIds returns the CommentFollowerIds field if non-nil, zero value otherwise.

### GetCommentFollowerIdsOk

`func (o *TasksIdJsonTodoItem1) GetCommentFollowerIdsOk() (*string, bool)`

GetCommentFollowerIdsOk returns a tuple with the CommentFollowerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentFollowerIds

`func (o *TasksIdJsonTodoItem1) SetCommentFollowerIds(v string)`

SetCommentFollowerIds sets CommentFollowerIds field to given value.

### HasCommentFollowerIds

`func (o *TasksIdJsonTodoItem1) HasCommentFollowerIds() bool`

HasCommentFollowerIds returns a boolean if a field has been set.

### GetContent

`func (o *TasksIdJsonTodoItem1) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *TasksIdJsonTodoItem1) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *TasksIdJsonTodoItem1) SetContent(v string)`

SetContent sets Content field to given value.


### GetDescription

`func (o *TasksIdJsonTodoItem1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TasksIdJsonTodoItem1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TasksIdJsonTodoItem1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TasksIdJsonTodoItem1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDueDate

`func (o *TasksIdJsonTodoItem1) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *TasksIdJsonTodoItem1) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *TasksIdJsonTodoItem1) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *TasksIdJsonTodoItem1) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetEstimatedMinutes

`func (o *TasksIdJsonTodoItem1) GetEstimatedMinutes() string`

GetEstimatedMinutes returns the EstimatedMinutes field if non-nil, zero value otherwise.

### GetEstimatedMinutesOk

`func (o *TasksIdJsonTodoItem1) GetEstimatedMinutesOk() (*string, bool)`

GetEstimatedMinutesOk returns a tuple with the EstimatedMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedMinutes

`func (o *TasksIdJsonTodoItem1) SetEstimatedMinutes(v string)`

SetEstimatedMinutes sets EstimatedMinutes field to given value.

### HasEstimatedMinutes

`func (o *TasksIdJsonTodoItem1) HasEstimatedMinutes() bool`

HasEstimatedMinutes returns a boolean if a field has been set.

### GetGrantAccessTo

`func (o *TasksIdJsonTodoItem1) GetGrantAccessTo() string`

GetGrantAccessTo returns the GrantAccessTo field if non-nil, zero value otherwise.

### GetGrantAccessToOk

`func (o *TasksIdJsonTodoItem1) GetGrantAccessToOk() (*string, bool)`

GetGrantAccessToOk returns a tuple with the GrantAccessTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantAccessTo

`func (o *TasksIdJsonTodoItem1) SetGrantAccessTo(v string)`

SetGrantAccessTo sets GrantAccessTo field to given value.

### HasGrantAccessTo

`func (o *TasksIdJsonTodoItem1) HasGrantAccessTo() bool`

HasGrantAccessTo returns a boolean if a field has been set.

### GetNotify

`func (o *TasksIdJsonTodoItem1) GetNotify() bool`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *TasksIdJsonTodoItem1) GetNotifyOk() (*bool, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *TasksIdJsonTodoItem1) SetNotify(v bool)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *TasksIdJsonTodoItem1) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetParentTaskId

`func (o *TasksIdJsonTodoItem1) GetParentTaskId() string`

GetParentTaskId returns the ParentTaskId field if non-nil, zero value otherwise.

### GetParentTaskIdOk

`func (o *TasksIdJsonTodoItem1) GetParentTaskIdOk() (*string, bool)`

GetParentTaskIdOk returns a tuple with the ParentTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTaskId

`func (o *TasksIdJsonTodoItem1) SetParentTaskId(v string)`

SetParentTaskId sets ParentTaskId field to given value.

### HasParentTaskId

`func (o *TasksIdJsonTodoItem1) HasParentTaskId() bool`

HasParentTaskId returns a boolean if a field has been set.

### GetPendingFileAttachments

`func (o *TasksIdJsonTodoItem1) GetPendingFileAttachments() string`

GetPendingFileAttachments returns the PendingFileAttachments field if non-nil, zero value otherwise.

### GetPendingFileAttachmentsOk

`func (o *TasksIdJsonTodoItem1) GetPendingFileAttachmentsOk() (*string, bool)`

GetPendingFileAttachmentsOk returns a tuple with the PendingFileAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingFileAttachments

`func (o *TasksIdJsonTodoItem1) SetPendingFileAttachments(v string)`

SetPendingFileAttachments sets PendingFileAttachments field to given value.

### HasPendingFileAttachments

`func (o *TasksIdJsonTodoItem1) HasPendingFileAttachments() bool`

HasPendingFileAttachments returns a boolean if a field has been set.

### GetPositionAfterTask

`func (o *TasksIdJsonTodoItem1) GetPositionAfterTask() int32`

GetPositionAfterTask returns the PositionAfterTask field if non-nil, zero value otherwise.

### GetPositionAfterTaskOk

`func (o *TasksIdJsonTodoItem1) GetPositionAfterTaskOk() (*int32, bool)`

GetPositionAfterTaskOk returns a tuple with the PositionAfterTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionAfterTask

`func (o *TasksIdJsonTodoItem1) SetPositionAfterTask(v int32)`

SetPositionAfterTask sets PositionAfterTask field to given value.

### HasPositionAfterTask

`func (o *TasksIdJsonTodoItem1) HasPositionAfterTask() bool`

HasPositionAfterTask returns a boolean if a field has been set.

### GetPredecessors

`func (o *TasksIdJsonTodoItem1) GetPredecessors() []TasksIdJsonTodoItem1Predecessors`

GetPredecessors returns the Predecessors field if non-nil, zero value otherwise.

### GetPredecessorsOk

`func (o *TasksIdJsonTodoItem1) GetPredecessorsOk() (*[]TasksIdJsonTodoItem1Predecessors, bool)`

GetPredecessorsOk returns a tuple with the Predecessors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredecessors

`func (o *TasksIdJsonTodoItem1) SetPredecessors(v []TasksIdJsonTodoItem1Predecessors)`

SetPredecessors sets Predecessors field to given value.

### HasPredecessors

`func (o *TasksIdJsonTodoItem1) HasPredecessors() bool`

HasPredecessors returns a boolean if a field has been set.

### GetPriority

`func (o *TasksIdJsonTodoItem1) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TasksIdJsonTodoItem1) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TasksIdJsonTodoItem1) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TasksIdJsonTodoItem1) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPrivate

`func (o *TasksIdJsonTodoItem1) GetPrivate() int32`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *TasksIdJsonTodoItem1) GetPrivateOk() (*int32, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *TasksIdJsonTodoItem1) SetPrivate(v int32)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *TasksIdJsonTodoItem1) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetProgress

`func (o *TasksIdJsonTodoItem1) GetProgress() string`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *TasksIdJsonTodoItem1) GetProgressOk() (*string, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *TasksIdJsonTodoItem1) SetProgress(v string)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *TasksIdJsonTodoItem1) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetResponsiblePartyId

`func (o *TasksIdJsonTodoItem1) GetResponsiblePartyId() int32`

GetResponsiblePartyId returns the ResponsiblePartyId field if non-nil, zero value otherwise.

### GetResponsiblePartyIdOk

`func (o *TasksIdJsonTodoItem1) GetResponsiblePartyIdOk() (*int32, bool)`

GetResponsiblePartyIdOk returns a tuple with the ResponsiblePartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsiblePartyId

`func (o *TasksIdJsonTodoItem1) SetResponsiblePartyId(v int32)`

SetResponsiblePartyId sets ResponsiblePartyId field to given value.

### HasResponsiblePartyId

`func (o *TasksIdJsonTodoItem1) HasResponsiblePartyId() bool`

HasResponsiblePartyId returns a boolean if a field has been set.

### GetStartDate

`func (o *TasksIdJsonTodoItem1) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TasksIdJsonTodoItem1) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TasksIdJsonTodoItem1) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *TasksIdJsonTodoItem1) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTags

`func (o *TasksIdJsonTodoItem1) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TasksIdJsonTodoItem1) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TasksIdJsonTodoItem1) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TasksIdJsonTodoItem1) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


