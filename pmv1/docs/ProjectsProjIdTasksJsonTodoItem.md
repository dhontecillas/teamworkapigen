# ProjectsProjIdTasksJsonTodoItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | Pointer to **string** |  | [optional] 
**AttachmentsCategoryIds** | Pointer to **string** |  | [optional] 
**ColumnId** | Pointer to **int32** |  | [optional] 
**Completed** | Pointer to **bool** |  | [optional] 
**Content** | **string** |  | 
**CreatorId** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EstimatedMinutes** | Pointer to **int32** |  | [optional] 
**EveryoneMustDo** | Pointer to **bool** |  | [optional] 
**GrantAccessTo** | Pointer to **string** |  | [optional] 
**ParentTaskId** | Pointer to **int32** |  | [optional] 
**PendingFileAttachments** | Pointer to **[]map[string]interface{}** |  | [optional] 
**PendingFileAttachmentsCategoryIds** | Pointer to **string** |  | [optional] 
**Predecessors** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**Private** | Pointer to **bool** |  | [optional] 
**Progress** | Pointer to **int32** |  | [optional] 
**RemoveOtherFiles** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **string** |  | [optional] 
**UpdateFiles** | Pointer to **bool** |  | [optional] 
**UseDefaults** | Pointer to **bool** |  | [optional] 

## Methods

### NewProjectsProjIdTasksJsonTodoItem

`func NewProjectsProjIdTasksJsonTodoItem(content string, ) *ProjectsProjIdTasksJsonTodoItem`

NewProjectsProjIdTasksJsonTodoItem instantiates a new ProjectsProjIdTasksJsonTodoItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsProjIdTasksJsonTodoItemWithDefaults

`func NewProjectsProjIdTasksJsonTodoItemWithDefaults() *ProjectsProjIdTasksJsonTodoItem`

NewProjectsProjIdTasksJsonTodoItemWithDefaults instantiates a new ProjectsProjIdTasksJsonTodoItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *ProjectsProjIdTasksJsonTodoItem) GetAttachments() string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ProjectsProjIdTasksJsonTodoItem) GetAttachmentsOk() (*string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ProjectsProjIdTasksJsonTodoItem) SetAttachments(v string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ProjectsProjIdTasksJsonTodoItem) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetAttachmentsCategoryIds

`func (o *ProjectsProjIdTasksJsonTodoItem) GetAttachmentsCategoryIds() string`

GetAttachmentsCategoryIds returns the AttachmentsCategoryIds field if non-nil, zero value otherwise.

### GetAttachmentsCategoryIdsOk

`func (o *ProjectsProjIdTasksJsonTodoItem) GetAttachmentsCategoryIdsOk() (*string, bool)`

GetAttachmentsCategoryIdsOk returns a tuple with the AttachmentsCategoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentsCategoryIds

`func (o *ProjectsProjIdTasksJsonTodoItem) SetAttachmentsCategoryIds(v string)`

SetAttachmentsCategoryIds sets AttachmentsCategoryIds field to given value.

### HasAttachmentsCategoryIds

`func (o *ProjectsProjIdTasksJsonTodoItem) HasAttachmentsCategoryIds() bool`

HasAttachmentsCategoryIds returns a boolean if a field has been set.

### GetColumnId

`func (o *ProjectsProjIdTasksJsonTodoItem) GetColumnId() int32`

GetColumnId returns the ColumnId field if non-nil, zero value otherwise.

### GetColumnIdOk

`func (o *ProjectsProjIdTasksJsonTodoItem) GetColumnIdOk() (*int32, bool)`

GetColumnIdOk returns a tuple with the ColumnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnId

`func (o *ProjectsProjIdTasksJsonTodoItem) SetColumnId(v int32)`

SetColumnId sets ColumnId field to given value.

### HasColumnId

`func (o *ProjectsProjIdTasksJsonTodoItem) HasColumnId() bool`

HasColumnId returns a boolean if a field has been set.

### GetCompleted

`func (o *ProjectsProjIdTasksJsonTodoItem) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *ProjectsProjIdTasksJsonTodoItem) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *ProjectsProjIdTasksJsonTodoItem) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *ProjectsProjIdTasksJsonTodoItem) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetContent

`func (o *ProjectsProjIdTasksJsonTodoItem) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ProjectsProjIdTasksJsonTodoItem) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ProjectsProjIdTasksJsonTodoItem) SetContent(v string)`

SetContent sets Content field to given value.


### GetCreatorId

`func (o *ProjectsProjIdTasksJsonTodoItem) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *ProjectsProjIdTasksJsonTodoItem) GetCreatorIdOk() (*int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *ProjectsProjIdTasksJsonTodoItem) SetCreatorId(v int32)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *ProjectsProjIdTasksJsonTodoItem) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetDescription

`func (o *ProjectsProjIdTasksJsonTodoItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectsProjIdTasksJsonTodoItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectsProjIdTasksJsonTodoItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectsProjIdTasksJsonTodoItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEstimatedMinutes

`func (o *ProjectsProjIdTasksJsonTodoItem) GetEstimatedMinutes() int32`

GetEstimatedMinutes returns the EstimatedMinutes field if non-nil, zero value otherwise.

### GetEstimatedMinutesOk

`func (o *ProjectsProjIdTasksJsonTodoItem) GetEstimatedMinutesOk() (*int32, bool)`

GetEstimatedMinutesOk returns a tuple with the EstimatedMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedMinutes

`func (o *ProjectsProjIdTasksJsonTodoItem) SetEstimatedMinutes(v int32)`

SetEstimatedMinutes sets EstimatedMinutes field to given value.

### HasEstimatedMinutes

`func (o *ProjectsProjIdTasksJsonTodoItem) HasEstimatedMinutes() bool`

HasEstimatedMinutes returns a boolean if a field has been set.

### GetEveryoneMustDo

`func (o *ProjectsProjIdTasksJsonTodoItem) GetEveryoneMustDo() bool`

GetEveryoneMustDo returns the EveryoneMustDo field if non-nil, zero value otherwise.

### GetEveryoneMustDoOk

`func (o *ProjectsProjIdTasksJsonTodoItem) GetEveryoneMustDoOk() (*bool, bool)`

GetEveryoneMustDoOk returns a tuple with the EveryoneMustDo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEveryoneMustDo

`func (o *ProjectsProjIdTasksJsonTodoItem) SetEveryoneMustDo(v bool)`

SetEveryoneMustDo sets EveryoneMustDo field to given value.

### HasEveryoneMustDo

`func (o *ProjectsProjIdTasksJsonTodoItem) HasEveryoneMustDo() bool`

HasEveryoneMustDo returns a boolean if a field has been set.

### GetGrantAccessTo

`func (o *ProjectsProjIdTasksJsonTodoItem) GetGrantAccessTo() string`

GetGrantAccessTo returns the GrantAccessTo field if non-nil, zero value otherwise.

### GetGrantAccessToOk

`func (o *ProjectsProjIdTasksJsonTodoItem) GetGrantAccessToOk() (*string, bool)`

GetGrantAccessToOk returns a tuple with the GrantAccessTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantAccessTo

`func (o *ProjectsProjIdTasksJsonTodoItem) SetGrantAccessTo(v string)`

SetGrantAccessTo sets GrantAccessTo field to given value.

### HasGrantAccessTo

`func (o *ProjectsProjIdTasksJsonTodoItem) HasGrantAccessTo() bool`

HasGrantAccessTo returns a boolean if a field has been set.

### GetParentTaskId

`func (o *ProjectsProjIdTasksJsonTodoItem) GetParentTaskId() int32`

GetParentTaskId returns the ParentTaskId field if non-nil, zero value otherwise.

### GetParentTaskIdOk

`func (o *ProjectsProjIdTasksJsonTodoItem) GetParentTaskIdOk() (*int32, bool)`

GetParentTaskIdOk returns a tuple with the ParentTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTaskId

`func (o *ProjectsProjIdTasksJsonTodoItem) SetParentTaskId(v int32)`

SetParentTaskId sets ParentTaskId field to given value.

### HasParentTaskId

`func (o *ProjectsProjIdTasksJsonTodoItem) HasParentTaskId() bool`

HasParentTaskId returns a boolean if a field has been set.

### GetPendingFileAttachments

`func (o *ProjectsProjIdTasksJsonTodoItem) GetPendingFileAttachments() []map[string]interface{}`

GetPendingFileAttachments returns the PendingFileAttachments field if non-nil, zero value otherwise.

### GetPendingFileAttachmentsOk

`func (o *ProjectsProjIdTasksJsonTodoItem) GetPendingFileAttachmentsOk() (*[]map[string]interface{}, bool)`

GetPendingFileAttachmentsOk returns a tuple with the PendingFileAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingFileAttachments

`func (o *ProjectsProjIdTasksJsonTodoItem) SetPendingFileAttachments(v []map[string]interface{})`

SetPendingFileAttachments sets PendingFileAttachments field to given value.

### HasPendingFileAttachments

`func (o *ProjectsProjIdTasksJsonTodoItem) HasPendingFileAttachments() bool`

HasPendingFileAttachments returns a boolean if a field has been set.

### GetPendingFileAttachmentsCategoryIds

`func (o *ProjectsProjIdTasksJsonTodoItem) GetPendingFileAttachmentsCategoryIds() string`

GetPendingFileAttachmentsCategoryIds returns the PendingFileAttachmentsCategoryIds field if non-nil, zero value otherwise.

### GetPendingFileAttachmentsCategoryIdsOk

`func (o *ProjectsProjIdTasksJsonTodoItem) GetPendingFileAttachmentsCategoryIdsOk() (*string, bool)`

GetPendingFileAttachmentsCategoryIdsOk returns a tuple with the PendingFileAttachmentsCategoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingFileAttachmentsCategoryIds

`func (o *ProjectsProjIdTasksJsonTodoItem) SetPendingFileAttachmentsCategoryIds(v string)`

SetPendingFileAttachmentsCategoryIds sets PendingFileAttachmentsCategoryIds field to given value.

### HasPendingFileAttachmentsCategoryIds

`func (o *ProjectsProjIdTasksJsonTodoItem) HasPendingFileAttachmentsCategoryIds() bool`

HasPendingFileAttachmentsCategoryIds returns a boolean if a field has been set.

### GetPredecessors

`func (o *ProjectsProjIdTasksJsonTodoItem) GetPredecessors() []map[string]interface{}`

GetPredecessors returns the Predecessors field if non-nil, zero value otherwise.

### GetPredecessorsOk

`func (o *ProjectsProjIdTasksJsonTodoItem) GetPredecessorsOk() (*[]map[string]interface{}, bool)`

GetPredecessorsOk returns a tuple with the Predecessors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredecessors

`func (o *ProjectsProjIdTasksJsonTodoItem) SetPredecessors(v []map[string]interface{})`

SetPredecessors sets Predecessors field to given value.

### HasPredecessors

`func (o *ProjectsProjIdTasksJsonTodoItem) HasPredecessors() bool`

HasPredecessors returns a boolean if a field has been set.

### GetPriority

`func (o *ProjectsProjIdTasksJsonTodoItem) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ProjectsProjIdTasksJsonTodoItem) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ProjectsProjIdTasksJsonTodoItem) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ProjectsProjIdTasksJsonTodoItem) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPrivate

`func (o *ProjectsProjIdTasksJsonTodoItem) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *ProjectsProjIdTasksJsonTodoItem) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *ProjectsProjIdTasksJsonTodoItem) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *ProjectsProjIdTasksJsonTodoItem) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetProgress

`func (o *ProjectsProjIdTasksJsonTodoItem) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ProjectsProjIdTasksJsonTodoItem) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ProjectsProjIdTasksJsonTodoItem) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *ProjectsProjIdTasksJsonTodoItem) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetRemoveOtherFiles

`func (o *ProjectsProjIdTasksJsonTodoItem) GetRemoveOtherFiles() bool`

GetRemoveOtherFiles returns the RemoveOtherFiles field if non-nil, zero value otherwise.

### GetRemoveOtherFilesOk

`func (o *ProjectsProjIdTasksJsonTodoItem) GetRemoveOtherFilesOk() (*bool, bool)`

GetRemoveOtherFilesOk returns a tuple with the RemoveOtherFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveOtherFiles

`func (o *ProjectsProjIdTasksJsonTodoItem) SetRemoveOtherFiles(v bool)`

SetRemoveOtherFiles sets RemoveOtherFiles field to given value.

### HasRemoveOtherFiles

`func (o *ProjectsProjIdTasksJsonTodoItem) HasRemoveOtherFiles() bool`

HasRemoveOtherFiles returns a boolean if a field has been set.

### GetTags

`func (o *ProjectsProjIdTasksJsonTodoItem) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProjectsProjIdTasksJsonTodoItem) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProjectsProjIdTasksJsonTodoItem) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProjectsProjIdTasksJsonTodoItem) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdateFiles

`func (o *ProjectsProjIdTasksJsonTodoItem) GetUpdateFiles() bool`

GetUpdateFiles returns the UpdateFiles field if non-nil, zero value otherwise.

### GetUpdateFilesOk

`func (o *ProjectsProjIdTasksJsonTodoItem) GetUpdateFilesOk() (*bool, bool)`

GetUpdateFilesOk returns a tuple with the UpdateFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateFiles

`func (o *ProjectsProjIdTasksJsonTodoItem) SetUpdateFiles(v bool)`

SetUpdateFiles sets UpdateFiles field to given value.

### HasUpdateFiles

`func (o *ProjectsProjIdTasksJsonTodoItem) HasUpdateFiles() bool`

HasUpdateFiles returns a boolean if a field has been set.

### GetUseDefaults

`func (o *ProjectsProjIdTasksJsonTodoItem) GetUseDefaults() bool`

GetUseDefaults returns the UseDefaults field if non-nil, zero value otherwise.

### GetUseDefaultsOk

`func (o *ProjectsProjIdTasksJsonTodoItem) GetUseDefaultsOk() (*bool, bool)`

GetUseDefaultsOk returns a tuple with the UseDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDefaults

`func (o *ProjectsProjIdTasksJsonTodoItem) SetUseDefaults(v bool)`

SetUseDefaults sets UseDefaults field to given value.

### HasUseDefaults

`func (o *ProjectsProjIdTasksJsonTodoItem) HasUseDefaults() bool`

HasUseDefaults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


