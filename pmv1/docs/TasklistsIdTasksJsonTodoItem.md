# TasklistsIdTasksJsonTodoItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | Pointer to **string** |  | [optional] 
**AttachmentsCategoryIds** | Pointer to **string** |  | [optional] 
**ChangeFollowerIds** | Pointer to **int32** |  | [optional] 
**ColumnId** | Pointer to **int32** |  | [optional] 
**CommentFollowerIds** | Pointer to **int32** |  | [optional] 
**Completed** | Pointer to **bool** |  | [optional] 
**Content** | **string** |  | 
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
**TasklistId** | Pointer to **int32** |  | [optional] 
**UpdateFiles** | Pointer to **bool** |  | [optional] 
**UseDefaults** | Pointer to **bool** |  | [optional] 

## Methods

### NewTasklistsIdTasksJsonTodoItem

`func NewTasklistsIdTasksJsonTodoItem(content string, ) *TasklistsIdTasksJsonTodoItem`

NewTasklistsIdTasksJsonTodoItem instantiates a new TasklistsIdTasksJsonTodoItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTasklistsIdTasksJsonTodoItemWithDefaults

`func NewTasklistsIdTasksJsonTodoItemWithDefaults() *TasklistsIdTasksJsonTodoItem`

NewTasklistsIdTasksJsonTodoItemWithDefaults instantiates a new TasklistsIdTasksJsonTodoItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *TasklistsIdTasksJsonTodoItem) GetAttachments() string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TasklistsIdTasksJsonTodoItem) GetAttachmentsOk() (*string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TasklistsIdTasksJsonTodoItem) SetAttachments(v string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *TasklistsIdTasksJsonTodoItem) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetAttachmentsCategoryIds

`func (o *TasklistsIdTasksJsonTodoItem) GetAttachmentsCategoryIds() string`

GetAttachmentsCategoryIds returns the AttachmentsCategoryIds field if non-nil, zero value otherwise.

### GetAttachmentsCategoryIdsOk

`func (o *TasklistsIdTasksJsonTodoItem) GetAttachmentsCategoryIdsOk() (*string, bool)`

GetAttachmentsCategoryIdsOk returns a tuple with the AttachmentsCategoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentsCategoryIds

`func (o *TasklistsIdTasksJsonTodoItem) SetAttachmentsCategoryIds(v string)`

SetAttachmentsCategoryIds sets AttachmentsCategoryIds field to given value.

### HasAttachmentsCategoryIds

`func (o *TasklistsIdTasksJsonTodoItem) HasAttachmentsCategoryIds() bool`

HasAttachmentsCategoryIds returns a boolean if a field has been set.

### GetChangeFollowerIds

`func (o *TasklistsIdTasksJsonTodoItem) GetChangeFollowerIds() int32`

GetChangeFollowerIds returns the ChangeFollowerIds field if non-nil, zero value otherwise.

### GetChangeFollowerIdsOk

`func (o *TasklistsIdTasksJsonTodoItem) GetChangeFollowerIdsOk() (*int32, bool)`

GetChangeFollowerIdsOk returns a tuple with the ChangeFollowerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeFollowerIds

`func (o *TasklistsIdTasksJsonTodoItem) SetChangeFollowerIds(v int32)`

SetChangeFollowerIds sets ChangeFollowerIds field to given value.

### HasChangeFollowerIds

`func (o *TasklistsIdTasksJsonTodoItem) HasChangeFollowerIds() bool`

HasChangeFollowerIds returns a boolean if a field has been set.

### GetColumnId

`func (o *TasklistsIdTasksJsonTodoItem) GetColumnId() int32`

GetColumnId returns the ColumnId field if non-nil, zero value otherwise.

### GetColumnIdOk

`func (o *TasklistsIdTasksJsonTodoItem) GetColumnIdOk() (*int32, bool)`

GetColumnIdOk returns a tuple with the ColumnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnId

`func (o *TasklistsIdTasksJsonTodoItem) SetColumnId(v int32)`

SetColumnId sets ColumnId field to given value.

### HasColumnId

`func (o *TasklistsIdTasksJsonTodoItem) HasColumnId() bool`

HasColumnId returns a boolean if a field has been set.

### GetCommentFollowerIds

`func (o *TasklistsIdTasksJsonTodoItem) GetCommentFollowerIds() int32`

GetCommentFollowerIds returns the CommentFollowerIds field if non-nil, zero value otherwise.

### GetCommentFollowerIdsOk

`func (o *TasklistsIdTasksJsonTodoItem) GetCommentFollowerIdsOk() (*int32, bool)`

GetCommentFollowerIdsOk returns a tuple with the CommentFollowerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentFollowerIds

`func (o *TasklistsIdTasksJsonTodoItem) SetCommentFollowerIds(v int32)`

SetCommentFollowerIds sets CommentFollowerIds field to given value.

### HasCommentFollowerIds

`func (o *TasklistsIdTasksJsonTodoItem) HasCommentFollowerIds() bool`

HasCommentFollowerIds returns a boolean if a field has been set.

### GetCompleted

`func (o *TasklistsIdTasksJsonTodoItem) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *TasklistsIdTasksJsonTodoItem) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *TasklistsIdTasksJsonTodoItem) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *TasklistsIdTasksJsonTodoItem) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetContent

`func (o *TasklistsIdTasksJsonTodoItem) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *TasklistsIdTasksJsonTodoItem) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *TasklistsIdTasksJsonTodoItem) SetContent(v string)`

SetContent sets Content field to given value.


### GetCreatorId

`func (o *TasklistsIdTasksJsonTodoItem) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *TasklistsIdTasksJsonTodoItem) GetCreatorIdOk() (*int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *TasklistsIdTasksJsonTodoItem) SetCreatorId(v int32)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *TasklistsIdTasksJsonTodoItem) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetDescription

`func (o *TasklistsIdTasksJsonTodoItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TasklistsIdTasksJsonTodoItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TasklistsIdTasksJsonTodoItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TasklistsIdTasksJsonTodoItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDueDate

`func (o *TasklistsIdTasksJsonTodoItem) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *TasklistsIdTasksJsonTodoItem) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *TasklistsIdTasksJsonTodoItem) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *TasklistsIdTasksJsonTodoItem) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetEstimatedMinutes

`func (o *TasklistsIdTasksJsonTodoItem) GetEstimatedMinutes() int32`

GetEstimatedMinutes returns the EstimatedMinutes field if non-nil, zero value otherwise.

### GetEstimatedMinutesOk

`func (o *TasklistsIdTasksJsonTodoItem) GetEstimatedMinutesOk() (*int32, bool)`

GetEstimatedMinutesOk returns a tuple with the EstimatedMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedMinutes

`func (o *TasklistsIdTasksJsonTodoItem) SetEstimatedMinutes(v int32)`

SetEstimatedMinutes sets EstimatedMinutes field to given value.

### HasEstimatedMinutes

`func (o *TasklistsIdTasksJsonTodoItem) HasEstimatedMinutes() bool`

HasEstimatedMinutes returns a boolean if a field has been set.

### GetEveryoneMustDo

`func (o *TasklistsIdTasksJsonTodoItem) GetEveryoneMustDo() bool`

GetEveryoneMustDo returns the EveryoneMustDo field if non-nil, zero value otherwise.

### GetEveryoneMustDoOk

`func (o *TasklistsIdTasksJsonTodoItem) GetEveryoneMustDoOk() (*bool, bool)`

GetEveryoneMustDoOk returns a tuple with the EveryoneMustDo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEveryoneMustDo

`func (o *TasklistsIdTasksJsonTodoItem) SetEveryoneMustDo(v bool)`

SetEveryoneMustDo sets EveryoneMustDo field to given value.

### HasEveryoneMustDo

`func (o *TasklistsIdTasksJsonTodoItem) HasEveryoneMustDo() bool`

HasEveryoneMustDo returns a boolean if a field has been set.

### GetGrantAccessTo

`func (o *TasklistsIdTasksJsonTodoItem) GetGrantAccessTo() string`

GetGrantAccessTo returns the GrantAccessTo field if non-nil, zero value otherwise.

### GetGrantAccessToOk

`func (o *TasklistsIdTasksJsonTodoItem) GetGrantAccessToOk() (*string, bool)`

GetGrantAccessToOk returns a tuple with the GrantAccessTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantAccessTo

`func (o *TasklistsIdTasksJsonTodoItem) SetGrantAccessTo(v string)`

SetGrantAccessTo sets GrantAccessTo field to given value.

### HasGrantAccessTo

`func (o *TasklistsIdTasksJsonTodoItem) HasGrantAccessTo() bool`

HasGrantAccessTo returns a boolean if a field has been set.

### GetNotify

`func (o *TasklistsIdTasksJsonTodoItem) GetNotify() bool`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *TasklistsIdTasksJsonTodoItem) GetNotifyOk() (*bool, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *TasklistsIdTasksJsonTodoItem) SetNotify(v bool)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *TasklistsIdTasksJsonTodoItem) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetParentTaskId

`func (o *TasklistsIdTasksJsonTodoItem) GetParentTaskId() int32`

GetParentTaskId returns the ParentTaskId field if non-nil, zero value otherwise.

### GetParentTaskIdOk

`func (o *TasklistsIdTasksJsonTodoItem) GetParentTaskIdOk() (*int32, bool)`

GetParentTaskIdOk returns a tuple with the ParentTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTaskId

`func (o *TasklistsIdTasksJsonTodoItem) SetParentTaskId(v int32)`

SetParentTaskId sets ParentTaskId field to given value.

### HasParentTaskId

`func (o *TasklistsIdTasksJsonTodoItem) HasParentTaskId() bool`

HasParentTaskId returns a boolean if a field has been set.

### GetPendingFileAttachments

`func (o *TasklistsIdTasksJsonTodoItem) GetPendingFileAttachments() []map[string]interface{}`

GetPendingFileAttachments returns the PendingFileAttachments field if non-nil, zero value otherwise.

### GetPendingFileAttachmentsOk

`func (o *TasklistsIdTasksJsonTodoItem) GetPendingFileAttachmentsOk() (*[]map[string]interface{}, bool)`

GetPendingFileAttachmentsOk returns a tuple with the PendingFileAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingFileAttachments

`func (o *TasklistsIdTasksJsonTodoItem) SetPendingFileAttachments(v []map[string]interface{})`

SetPendingFileAttachments sets PendingFileAttachments field to given value.

### HasPendingFileAttachments

`func (o *TasklistsIdTasksJsonTodoItem) HasPendingFileAttachments() bool`

HasPendingFileAttachments returns a boolean if a field has been set.

### GetPendingFileAttachmentsCategoryIds

`func (o *TasklistsIdTasksJsonTodoItem) GetPendingFileAttachmentsCategoryIds() string`

GetPendingFileAttachmentsCategoryIds returns the PendingFileAttachmentsCategoryIds field if non-nil, zero value otherwise.

### GetPendingFileAttachmentsCategoryIdsOk

`func (o *TasklistsIdTasksJsonTodoItem) GetPendingFileAttachmentsCategoryIdsOk() (*string, bool)`

GetPendingFileAttachmentsCategoryIdsOk returns a tuple with the PendingFileAttachmentsCategoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingFileAttachmentsCategoryIds

`func (o *TasklistsIdTasksJsonTodoItem) SetPendingFileAttachmentsCategoryIds(v string)`

SetPendingFileAttachmentsCategoryIds sets PendingFileAttachmentsCategoryIds field to given value.

### HasPendingFileAttachmentsCategoryIds

`func (o *TasklistsIdTasksJsonTodoItem) HasPendingFileAttachmentsCategoryIds() bool`

HasPendingFileAttachmentsCategoryIds returns a boolean if a field has been set.

### GetPredecessors

`func (o *TasklistsIdTasksJsonTodoItem) GetPredecessors() []map[string]interface{}`

GetPredecessors returns the Predecessors field if non-nil, zero value otherwise.

### GetPredecessorsOk

`func (o *TasklistsIdTasksJsonTodoItem) GetPredecessorsOk() (*[]map[string]interface{}, bool)`

GetPredecessorsOk returns a tuple with the Predecessors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredecessors

`func (o *TasklistsIdTasksJsonTodoItem) SetPredecessors(v []map[string]interface{})`

SetPredecessors sets Predecessors field to given value.

### HasPredecessors

`func (o *TasklistsIdTasksJsonTodoItem) HasPredecessors() bool`

HasPredecessors returns a boolean if a field has been set.

### GetPriority

`func (o *TasklistsIdTasksJsonTodoItem) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TasklistsIdTasksJsonTodoItem) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TasklistsIdTasksJsonTodoItem) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TasklistsIdTasksJsonTodoItem) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPrivate

`func (o *TasklistsIdTasksJsonTodoItem) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *TasklistsIdTasksJsonTodoItem) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *TasklistsIdTasksJsonTodoItem) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *TasklistsIdTasksJsonTodoItem) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetProgress

`func (o *TasklistsIdTasksJsonTodoItem) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *TasklistsIdTasksJsonTodoItem) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *TasklistsIdTasksJsonTodoItem) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *TasklistsIdTasksJsonTodoItem) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetRemoveOtherFiles

`func (o *TasklistsIdTasksJsonTodoItem) GetRemoveOtherFiles() bool`

GetRemoveOtherFiles returns the RemoveOtherFiles field if non-nil, zero value otherwise.

### GetRemoveOtherFilesOk

`func (o *TasklistsIdTasksJsonTodoItem) GetRemoveOtherFilesOk() (*bool, bool)`

GetRemoveOtherFilesOk returns a tuple with the RemoveOtherFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveOtherFiles

`func (o *TasklistsIdTasksJsonTodoItem) SetRemoveOtherFiles(v bool)`

SetRemoveOtherFiles sets RemoveOtherFiles field to given value.

### HasRemoveOtherFiles

`func (o *TasklistsIdTasksJsonTodoItem) HasRemoveOtherFiles() bool`

HasRemoveOtherFiles returns a boolean if a field has been set.

### GetRepeatOptions

`func (o *TasklistsIdTasksJsonTodoItem) GetRepeatOptions() TasklistsIdTasksJsonTodoItemRepeatOptions`

GetRepeatOptions returns the RepeatOptions field if non-nil, zero value otherwise.

### GetRepeatOptionsOk

`func (o *TasklistsIdTasksJsonTodoItem) GetRepeatOptionsOk() (*TasklistsIdTasksJsonTodoItemRepeatOptions, bool)`

GetRepeatOptionsOk returns a tuple with the RepeatOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatOptions

`func (o *TasklistsIdTasksJsonTodoItem) SetRepeatOptions(v TasklistsIdTasksJsonTodoItemRepeatOptions)`

SetRepeatOptions sets RepeatOptions field to given value.

### HasRepeatOptions

`func (o *TasklistsIdTasksJsonTodoItem) HasRepeatOptions() bool`

HasRepeatOptions returns a boolean if a field has been set.

### GetResponsiblePartyId

`func (o *TasklistsIdTasksJsonTodoItem) GetResponsiblePartyId() string`

GetResponsiblePartyId returns the ResponsiblePartyId field if non-nil, zero value otherwise.

### GetResponsiblePartyIdOk

`func (o *TasklistsIdTasksJsonTodoItem) GetResponsiblePartyIdOk() (*string, bool)`

GetResponsiblePartyIdOk returns a tuple with the ResponsiblePartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsiblePartyId

`func (o *TasklistsIdTasksJsonTodoItem) SetResponsiblePartyId(v string)`

SetResponsiblePartyId sets ResponsiblePartyId field to given value.

### HasResponsiblePartyId

`func (o *TasklistsIdTasksJsonTodoItem) HasResponsiblePartyId() bool`

HasResponsiblePartyId returns a boolean if a field has been set.

### GetStartDate

`func (o *TasklistsIdTasksJsonTodoItem) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TasklistsIdTasksJsonTodoItem) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TasklistsIdTasksJsonTodoItem) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *TasklistsIdTasksJsonTodoItem) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTags

`func (o *TasklistsIdTasksJsonTodoItem) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TasklistsIdTasksJsonTodoItem) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TasklistsIdTasksJsonTodoItem) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TasklistsIdTasksJsonTodoItem) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTasklistId

`func (o *TasklistsIdTasksJsonTodoItem) GetTasklistId() int32`

GetTasklistId returns the TasklistId field if non-nil, zero value otherwise.

### GetTasklistIdOk

`func (o *TasklistsIdTasksJsonTodoItem) GetTasklistIdOk() (*int32, bool)`

GetTasklistIdOk returns a tuple with the TasklistId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasklistId

`func (o *TasklistsIdTasksJsonTodoItem) SetTasklistId(v int32)`

SetTasklistId sets TasklistId field to given value.

### HasTasklistId

`func (o *TasklistsIdTasksJsonTodoItem) HasTasklistId() bool`

HasTasklistId returns a boolean if a field has been set.

### GetUpdateFiles

`func (o *TasklistsIdTasksJsonTodoItem) GetUpdateFiles() bool`

GetUpdateFiles returns the UpdateFiles field if non-nil, zero value otherwise.

### GetUpdateFilesOk

`func (o *TasklistsIdTasksJsonTodoItem) GetUpdateFilesOk() (*bool, bool)`

GetUpdateFilesOk returns a tuple with the UpdateFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateFiles

`func (o *TasklistsIdTasksJsonTodoItem) SetUpdateFiles(v bool)`

SetUpdateFiles sets UpdateFiles field to given value.

### HasUpdateFiles

`func (o *TasklistsIdTasksJsonTodoItem) HasUpdateFiles() bool`

HasUpdateFiles returns a boolean if a field has been set.

### GetUseDefaults

`func (o *TasklistsIdTasksJsonTodoItem) GetUseDefaults() bool`

GetUseDefaults returns the UseDefaults field if non-nil, zero value otherwise.

### GetUseDefaultsOk

`func (o *TasklistsIdTasksJsonTodoItem) GetUseDefaultsOk() (*bool, bool)`

GetUseDefaultsOk returns a tuple with the UseDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDefaults

`func (o *TasklistsIdTasksJsonTodoItem) SetUseDefaults(v bool)`

SetUseDefaults sets UseDefaults field to given value.

### HasUseDefaults

`func (o *TasklistsIdTasksJsonTodoItem) HasUseDefaults() bool`

HasUseDefaults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


