# TasksIdJsonTodoItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowUpdate** | Pointer to **bool** | This parameter can be used to allow the completed date of a Task to be updated on a closed task. A completer-id is optional if you want to change the person that completed the task. | [optional] 
**AttachmentsCategoryIds** | Pointer to **string** |  | [optional] 
**ChangeFollowerIds** | Pointer to **int32** | To set one or more users as Followers of a specified Task. To remove followers, send in a blank string. | [optional] 
**ColumnId** | Pointer to **int32** |  | [optional] 
**CommentFollowerIds** | Pointer to **int32** | To set one or more users as Followers of a specified Task. To remove followers, send in blank string. | [optional] 
**Completed** | Pointer to **bool** |  | [optional] 
**CompletedOn** | Pointer to **string** | Used with &#39;allow-update&#39;. Pass in the date of completion of the closed task.  | [optional] 
**CompleterId** | Pointer to **string** | If you are updating a task completion date that is closed,  you can optionally pass in the ID of the user you want to have completed the task.  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**CreatorId** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DueDate** | Pointer to **string** |  | [optional] 
**EstimatedMinutes** | Pointer to **int32** |  | [optional] 
**EveryoneMustDo** | Pointer to **bool** |  | [optional] 
**GrantAccessTo** | Pointer to **string** |  | [optional] 
**Notify** | Pointer to **bool** |  | [optional] 
**ParentTaskId** | Pointer to **int32** |  | [optional] 
**PendingFileAttachments** | Pointer to **[]map[string]interface{}** |  | [optional] 
**PendingFileAttachmentsCategoryIds** | Pointer to **string** |  | [optional] 
**Predecessors** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**Private** | Pointer to **bool** |  | [optional] 
**Progress** | Pointer to **int32** |  | [optional] 
**RemoveOtherFiles** | Pointer to **bool** |  | [optional] 
**RepeatOptions** | Pointer to [**TasklistsIdTasksJsonTodoItemRepeatOptions**](TasklistsIdTasksJsonTodoItemRepeatOptions.md) |  | [optional] 
**ResponsiblePartyId** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **string** |  | [optional] 
**UpdateFiles** | Pointer to **bool** |  | [optional] 
**UseDefaults** | Pointer to **bool** |  | [optional] 

## Methods

### NewTasksIdJsonTodoItem

`func NewTasksIdJsonTodoItem() *TasksIdJsonTodoItem`

NewTasksIdJsonTodoItem instantiates a new TasksIdJsonTodoItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTasksIdJsonTodoItemWithDefaults

`func NewTasksIdJsonTodoItemWithDefaults() *TasksIdJsonTodoItem`

NewTasksIdJsonTodoItemWithDefaults instantiates a new TasksIdJsonTodoItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowUpdate

`func (o *TasksIdJsonTodoItem) GetAllowUpdate() bool`

GetAllowUpdate returns the AllowUpdate field if non-nil, zero value otherwise.

### GetAllowUpdateOk

`func (o *TasksIdJsonTodoItem) GetAllowUpdateOk() (*bool, bool)`

GetAllowUpdateOk returns a tuple with the AllowUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUpdate

`func (o *TasksIdJsonTodoItem) SetAllowUpdate(v bool)`

SetAllowUpdate sets AllowUpdate field to given value.

### HasAllowUpdate

`func (o *TasksIdJsonTodoItem) HasAllowUpdate() bool`

HasAllowUpdate returns a boolean if a field has been set.

### GetAttachmentsCategoryIds

`func (o *TasksIdJsonTodoItem) GetAttachmentsCategoryIds() string`

GetAttachmentsCategoryIds returns the AttachmentsCategoryIds field if non-nil, zero value otherwise.

### GetAttachmentsCategoryIdsOk

`func (o *TasksIdJsonTodoItem) GetAttachmentsCategoryIdsOk() (*string, bool)`

GetAttachmentsCategoryIdsOk returns a tuple with the AttachmentsCategoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentsCategoryIds

`func (o *TasksIdJsonTodoItem) SetAttachmentsCategoryIds(v string)`

SetAttachmentsCategoryIds sets AttachmentsCategoryIds field to given value.

### HasAttachmentsCategoryIds

`func (o *TasksIdJsonTodoItem) HasAttachmentsCategoryIds() bool`

HasAttachmentsCategoryIds returns a boolean if a field has been set.

### GetChangeFollowerIds

`func (o *TasksIdJsonTodoItem) GetChangeFollowerIds() int32`

GetChangeFollowerIds returns the ChangeFollowerIds field if non-nil, zero value otherwise.

### GetChangeFollowerIdsOk

`func (o *TasksIdJsonTodoItem) GetChangeFollowerIdsOk() (*int32, bool)`

GetChangeFollowerIdsOk returns a tuple with the ChangeFollowerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeFollowerIds

`func (o *TasksIdJsonTodoItem) SetChangeFollowerIds(v int32)`

SetChangeFollowerIds sets ChangeFollowerIds field to given value.

### HasChangeFollowerIds

`func (o *TasksIdJsonTodoItem) HasChangeFollowerIds() bool`

HasChangeFollowerIds returns a boolean if a field has been set.

### GetColumnId

`func (o *TasksIdJsonTodoItem) GetColumnId() int32`

GetColumnId returns the ColumnId field if non-nil, zero value otherwise.

### GetColumnIdOk

`func (o *TasksIdJsonTodoItem) GetColumnIdOk() (*int32, bool)`

GetColumnIdOk returns a tuple with the ColumnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnId

`func (o *TasksIdJsonTodoItem) SetColumnId(v int32)`

SetColumnId sets ColumnId field to given value.

### HasColumnId

`func (o *TasksIdJsonTodoItem) HasColumnId() bool`

HasColumnId returns a boolean if a field has been set.

### GetCommentFollowerIds

`func (o *TasksIdJsonTodoItem) GetCommentFollowerIds() int32`

GetCommentFollowerIds returns the CommentFollowerIds field if non-nil, zero value otherwise.

### GetCommentFollowerIdsOk

`func (o *TasksIdJsonTodoItem) GetCommentFollowerIdsOk() (*int32, bool)`

GetCommentFollowerIdsOk returns a tuple with the CommentFollowerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentFollowerIds

`func (o *TasksIdJsonTodoItem) SetCommentFollowerIds(v int32)`

SetCommentFollowerIds sets CommentFollowerIds field to given value.

### HasCommentFollowerIds

`func (o *TasksIdJsonTodoItem) HasCommentFollowerIds() bool`

HasCommentFollowerIds returns a boolean if a field has been set.

### GetCompleted

`func (o *TasksIdJsonTodoItem) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *TasksIdJsonTodoItem) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *TasksIdJsonTodoItem) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *TasksIdJsonTodoItem) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetCompletedOn

`func (o *TasksIdJsonTodoItem) GetCompletedOn() string`

GetCompletedOn returns the CompletedOn field if non-nil, zero value otherwise.

### GetCompletedOnOk

`func (o *TasksIdJsonTodoItem) GetCompletedOnOk() (*string, bool)`

GetCompletedOnOk returns a tuple with the CompletedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedOn

`func (o *TasksIdJsonTodoItem) SetCompletedOn(v string)`

SetCompletedOn sets CompletedOn field to given value.

### HasCompletedOn

`func (o *TasksIdJsonTodoItem) HasCompletedOn() bool`

HasCompletedOn returns a boolean if a field has been set.

### GetCompleterId

`func (o *TasksIdJsonTodoItem) GetCompleterId() string`

GetCompleterId returns the CompleterId field if non-nil, zero value otherwise.

### GetCompleterIdOk

`func (o *TasksIdJsonTodoItem) GetCompleterIdOk() (*string, bool)`

GetCompleterIdOk returns a tuple with the CompleterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleterId

`func (o *TasksIdJsonTodoItem) SetCompleterId(v string)`

SetCompleterId sets CompleterId field to given value.

### HasCompleterId

`func (o *TasksIdJsonTodoItem) HasCompleterId() bool`

HasCompleterId returns a boolean if a field has been set.

### GetContent

`func (o *TasksIdJsonTodoItem) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *TasksIdJsonTodoItem) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *TasksIdJsonTodoItem) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *TasksIdJsonTodoItem) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatorId

`func (o *TasksIdJsonTodoItem) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *TasksIdJsonTodoItem) GetCreatorIdOk() (*int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *TasksIdJsonTodoItem) SetCreatorId(v int32)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *TasksIdJsonTodoItem) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetDescription

`func (o *TasksIdJsonTodoItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TasksIdJsonTodoItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TasksIdJsonTodoItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TasksIdJsonTodoItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDueDate

`func (o *TasksIdJsonTodoItem) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *TasksIdJsonTodoItem) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *TasksIdJsonTodoItem) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *TasksIdJsonTodoItem) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetEstimatedMinutes

`func (o *TasksIdJsonTodoItem) GetEstimatedMinutes() int32`

GetEstimatedMinutes returns the EstimatedMinutes field if non-nil, zero value otherwise.

### GetEstimatedMinutesOk

`func (o *TasksIdJsonTodoItem) GetEstimatedMinutesOk() (*int32, bool)`

GetEstimatedMinutesOk returns a tuple with the EstimatedMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedMinutes

`func (o *TasksIdJsonTodoItem) SetEstimatedMinutes(v int32)`

SetEstimatedMinutes sets EstimatedMinutes field to given value.

### HasEstimatedMinutes

`func (o *TasksIdJsonTodoItem) HasEstimatedMinutes() bool`

HasEstimatedMinutes returns a boolean if a field has been set.

### GetEveryoneMustDo

`func (o *TasksIdJsonTodoItem) GetEveryoneMustDo() bool`

GetEveryoneMustDo returns the EveryoneMustDo field if non-nil, zero value otherwise.

### GetEveryoneMustDoOk

`func (o *TasksIdJsonTodoItem) GetEveryoneMustDoOk() (*bool, bool)`

GetEveryoneMustDoOk returns a tuple with the EveryoneMustDo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEveryoneMustDo

`func (o *TasksIdJsonTodoItem) SetEveryoneMustDo(v bool)`

SetEveryoneMustDo sets EveryoneMustDo field to given value.

### HasEveryoneMustDo

`func (o *TasksIdJsonTodoItem) HasEveryoneMustDo() bool`

HasEveryoneMustDo returns a boolean if a field has been set.

### GetGrantAccessTo

`func (o *TasksIdJsonTodoItem) GetGrantAccessTo() string`

GetGrantAccessTo returns the GrantAccessTo field if non-nil, zero value otherwise.

### GetGrantAccessToOk

`func (o *TasksIdJsonTodoItem) GetGrantAccessToOk() (*string, bool)`

GetGrantAccessToOk returns a tuple with the GrantAccessTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantAccessTo

`func (o *TasksIdJsonTodoItem) SetGrantAccessTo(v string)`

SetGrantAccessTo sets GrantAccessTo field to given value.

### HasGrantAccessTo

`func (o *TasksIdJsonTodoItem) HasGrantAccessTo() bool`

HasGrantAccessTo returns a boolean if a field has been set.

### GetNotify

`func (o *TasksIdJsonTodoItem) GetNotify() bool`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *TasksIdJsonTodoItem) GetNotifyOk() (*bool, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *TasksIdJsonTodoItem) SetNotify(v bool)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *TasksIdJsonTodoItem) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetParentTaskId

`func (o *TasksIdJsonTodoItem) GetParentTaskId() int32`

GetParentTaskId returns the ParentTaskId field if non-nil, zero value otherwise.

### GetParentTaskIdOk

`func (o *TasksIdJsonTodoItem) GetParentTaskIdOk() (*int32, bool)`

GetParentTaskIdOk returns a tuple with the ParentTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTaskId

`func (o *TasksIdJsonTodoItem) SetParentTaskId(v int32)`

SetParentTaskId sets ParentTaskId field to given value.

### HasParentTaskId

`func (o *TasksIdJsonTodoItem) HasParentTaskId() bool`

HasParentTaskId returns a boolean if a field has been set.

### GetPendingFileAttachments

`func (o *TasksIdJsonTodoItem) GetPendingFileAttachments() []map[string]interface{}`

GetPendingFileAttachments returns the PendingFileAttachments field if non-nil, zero value otherwise.

### GetPendingFileAttachmentsOk

`func (o *TasksIdJsonTodoItem) GetPendingFileAttachmentsOk() (*[]map[string]interface{}, bool)`

GetPendingFileAttachmentsOk returns a tuple with the PendingFileAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingFileAttachments

`func (o *TasksIdJsonTodoItem) SetPendingFileAttachments(v []map[string]interface{})`

SetPendingFileAttachments sets PendingFileAttachments field to given value.

### HasPendingFileAttachments

`func (o *TasksIdJsonTodoItem) HasPendingFileAttachments() bool`

HasPendingFileAttachments returns a boolean if a field has been set.

### GetPendingFileAttachmentsCategoryIds

`func (o *TasksIdJsonTodoItem) GetPendingFileAttachmentsCategoryIds() string`

GetPendingFileAttachmentsCategoryIds returns the PendingFileAttachmentsCategoryIds field if non-nil, zero value otherwise.

### GetPendingFileAttachmentsCategoryIdsOk

`func (o *TasksIdJsonTodoItem) GetPendingFileAttachmentsCategoryIdsOk() (*string, bool)`

GetPendingFileAttachmentsCategoryIdsOk returns a tuple with the PendingFileAttachmentsCategoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingFileAttachmentsCategoryIds

`func (o *TasksIdJsonTodoItem) SetPendingFileAttachmentsCategoryIds(v string)`

SetPendingFileAttachmentsCategoryIds sets PendingFileAttachmentsCategoryIds field to given value.

### HasPendingFileAttachmentsCategoryIds

`func (o *TasksIdJsonTodoItem) HasPendingFileAttachmentsCategoryIds() bool`

HasPendingFileAttachmentsCategoryIds returns a boolean if a field has been set.

### GetPredecessors

`func (o *TasksIdJsonTodoItem) GetPredecessors() []map[string]interface{}`

GetPredecessors returns the Predecessors field if non-nil, zero value otherwise.

### GetPredecessorsOk

`func (o *TasksIdJsonTodoItem) GetPredecessorsOk() (*[]map[string]interface{}, bool)`

GetPredecessorsOk returns a tuple with the Predecessors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredecessors

`func (o *TasksIdJsonTodoItem) SetPredecessors(v []map[string]interface{})`

SetPredecessors sets Predecessors field to given value.

### HasPredecessors

`func (o *TasksIdJsonTodoItem) HasPredecessors() bool`

HasPredecessors returns a boolean if a field has been set.

### GetPriority

`func (o *TasksIdJsonTodoItem) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TasksIdJsonTodoItem) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TasksIdJsonTodoItem) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TasksIdJsonTodoItem) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPrivate

`func (o *TasksIdJsonTodoItem) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *TasksIdJsonTodoItem) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *TasksIdJsonTodoItem) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *TasksIdJsonTodoItem) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetProgress

`func (o *TasksIdJsonTodoItem) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *TasksIdJsonTodoItem) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *TasksIdJsonTodoItem) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *TasksIdJsonTodoItem) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetRemoveOtherFiles

`func (o *TasksIdJsonTodoItem) GetRemoveOtherFiles() bool`

GetRemoveOtherFiles returns the RemoveOtherFiles field if non-nil, zero value otherwise.

### GetRemoveOtherFilesOk

`func (o *TasksIdJsonTodoItem) GetRemoveOtherFilesOk() (*bool, bool)`

GetRemoveOtherFilesOk returns a tuple with the RemoveOtherFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveOtherFiles

`func (o *TasksIdJsonTodoItem) SetRemoveOtherFiles(v bool)`

SetRemoveOtherFiles sets RemoveOtherFiles field to given value.

### HasRemoveOtherFiles

`func (o *TasksIdJsonTodoItem) HasRemoveOtherFiles() bool`

HasRemoveOtherFiles returns a boolean if a field has been set.

### GetRepeatOptions

`func (o *TasksIdJsonTodoItem) GetRepeatOptions() TasklistsIdTasksJsonTodoItemRepeatOptions`

GetRepeatOptions returns the RepeatOptions field if non-nil, zero value otherwise.

### GetRepeatOptionsOk

`func (o *TasksIdJsonTodoItem) GetRepeatOptionsOk() (*TasklistsIdTasksJsonTodoItemRepeatOptions, bool)`

GetRepeatOptionsOk returns a tuple with the RepeatOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatOptions

`func (o *TasksIdJsonTodoItem) SetRepeatOptions(v TasklistsIdTasksJsonTodoItemRepeatOptions)`

SetRepeatOptions sets RepeatOptions field to given value.

### HasRepeatOptions

`func (o *TasksIdJsonTodoItem) HasRepeatOptions() bool`

HasRepeatOptions returns a boolean if a field has been set.

### GetResponsiblePartyId

`func (o *TasksIdJsonTodoItem) GetResponsiblePartyId() string`

GetResponsiblePartyId returns the ResponsiblePartyId field if non-nil, zero value otherwise.

### GetResponsiblePartyIdOk

`func (o *TasksIdJsonTodoItem) GetResponsiblePartyIdOk() (*string, bool)`

GetResponsiblePartyIdOk returns a tuple with the ResponsiblePartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsiblePartyId

`func (o *TasksIdJsonTodoItem) SetResponsiblePartyId(v string)`

SetResponsiblePartyId sets ResponsiblePartyId field to given value.

### HasResponsiblePartyId

`func (o *TasksIdJsonTodoItem) HasResponsiblePartyId() bool`

HasResponsiblePartyId returns a boolean if a field has been set.

### GetStartDate

`func (o *TasksIdJsonTodoItem) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TasksIdJsonTodoItem) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TasksIdJsonTodoItem) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *TasksIdJsonTodoItem) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTags

`func (o *TasksIdJsonTodoItem) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TasksIdJsonTodoItem) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TasksIdJsonTodoItem) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TasksIdJsonTodoItem) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdateFiles

`func (o *TasksIdJsonTodoItem) GetUpdateFiles() bool`

GetUpdateFiles returns the UpdateFiles field if non-nil, zero value otherwise.

### GetUpdateFilesOk

`func (o *TasksIdJsonTodoItem) GetUpdateFilesOk() (*bool, bool)`

GetUpdateFilesOk returns a tuple with the UpdateFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateFiles

`func (o *TasksIdJsonTodoItem) SetUpdateFiles(v bool)`

SetUpdateFiles sets UpdateFiles field to given value.

### HasUpdateFiles

`func (o *TasksIdJsonTodoItem) HasUpdateFiles() bool`

HasUpdateFiles returns a boolean if a field has been set.

### GetUseDefaults

`func (o *TasksIdJsonTodoItem) GetUseDefaults() bool`

GetUseDefaults returns the UseDefaults field if non-nil, zero value otherwise.

### GetUseDefaultsOk

`func (o *TasksIdJsonTodoItem) GetUseDefaultsOk() (*bool, bool)`

GetUseDefaultsOk returns a tuple with the UseDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDefaults

`func (o *TasksIdJsonTodoItem) SetUseDefaults(v bool)`

SetUseDefaults sets UseDefaults field to given value.

### HasUseDefaults

`func (o *TasksIdJsonTodoItem) HasUseDefaults() bool`

HasUseDefaults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


