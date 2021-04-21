# InlineResponse2002Cards

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | Pointer to **bool** |  | [optional] 
**ArchivedDate** | Pointer to **string** |  | [optional] 
**AssignedPeople** | Pointer to **[]string** |  | [optional] 
**CanComplete** | Pointer to **bool** |  | [optional] 
**CanEdit** | Pointer to **bool** |  | [optional] 
**ColumnId** | Pointer to **string** |  | [optional] 
**Completed** | Pointer to **bool** |  | [optional] 
**DateCompleted** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **string** |  | [optional] 
**DateUpdated** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**DeletedDate** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayOrder** | Pointer to **string** |  | [optional] 
**DueDate** | Pointer to **string** |  | [optional] 
**EstimatedMinutes** | Pointer to **string** |  | [optional] 
**HasLoggedTime** | Pointer to **bool** |  | [optional] 
**HasTickets** | Pointer to **bool** |  | [optional] 
**HasUnreadComments** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsRecurring** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NumActiveSubTasks** | Pointer to **string** |  | [optional] 
**NumAttachments** | Pointer to **string** |  | [optional] 
**NumComments** | Pointer to **string** |  | [optional] 
**NumDependencies** | Pointer to **string** |  | [optional] 
**NumPredecessors** | Pointer to **string** |  | [optional] 
**ParentTaskId** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**Private** | Pointer to **string** |  | [optional] 
**Progress** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]map[string]interface{}** |  | [optional] 
**TaskId** | Pointer to **string** |  | [optional] 
**TaskListId** | Pointer to **string** |  | [optional] 
**TaskListName** | Pointer to **string** |  | [optional] 
**TaskStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewInlineResponse2002Cards

`func NewInlineResponse2002Cards() *InlineResponse2002Cards`

NewInlineResponse2002Cards instantiates a new InlineResponse2002Cards object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2002CardsWithDefaults

`func NewInlineResponse2002CardsWithDefaults() *InlineResponse2002Cards`

NewInlineResponse2002CardsWithDefaults instantiates a new InlineResponse2002Cards object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *InlineResponse2002Cards) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *InlineResponse2002Cards) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *InlineResponse2002Cards) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *InlineResponse2002Cards) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetArchivedDate

`func (o *InlineResponse2002Cards) GetArchivedDate() string`

GetArchivedDate returns the ArchivedDate field if non-nil, zero value otherwise.

### GetArchivedDateOk

`func (o *InlineResponse2002Cards) GetArchivedDateOk() (*string, bool)`

GetArchivedDateOk returns a tuple with the ArchivedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedDate

`func (o *InlineResponse2002Cards) SetArchivedDate(v string)`

SetArchivedDate sets ArchivedDate field to given value.

### HasArchivedDate

`func (o *InlineResponse2002Cards) HasArchivedDate() bool`

HasArchivedDate returns a boolean if a field has been set.

### GetAssignedPeople

`func (o *InlineResponse2002Cards) GetAssignedPeople() []string`

GetAssignedPeople returns the AssignedPeople field if non-nil, zero value otherwise.

### GetAssignedPeopleOk

`func (o *InlineResponse2002Cards) GetAssignedPeopleOk() (*[]string, bool)`

GetAssignedPeopleOk returns a tuple with the AssignedPeople field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedPeople

`func (o *InlineResponse2002Cards) SetAssignedPeople(v []string)`

SetAssignedPeople sets AssignedPeople field to given value.

### HasAssignedPeople

`func (o *InlineResponse2002Cards) HasAssignedPeople() bool`

HasAssignedPeople returns a boolean if a field has been set.

### GetCanComplete

`func (o *InlineResponse2002Cards) GetCanComplete() bool`

GetCanComplete returns the CanComplete field if non-nil, zero value otherwise.

### GetCanCompleteOk

`func (o *InlineResponse2002Cards) GetCanCompleteOk() (*bool, bool)`

GetCanCompleteOk returns a tuple with the CanComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanComplete

`func (o *InlineResponse2002Cards) SetCanComplete(v bool)`

SetCanComplete sets CanComplete field to given value.

### HasCanComplete

`func (o *InlineResponse2002Cards) HasCanComplete() bool`

HasCanComplete returns a boolean if a field has been set.

### GetCanEdit

`func (o *InlineResponse2002Cards) GetCanEdit() bool`

GetCanEdit returns the CanEdit field if non-nil, zero value otherwise.

### GetCanEditOk

`func (o *InlineResponse2002Cards) GetCanEditOk() (*bool, bool)`

GetCanEditOk returns a tuple with the CanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEdit

`func (o *InlineResponse2002Cards) SetCanEdit(v bool)`

SetCanEdit sets CanEdit field to given value.

### HasCanEdit

`func (o *InlineResponse2002Cards) HasCanEdit() bool`

HasCanEdit returns a boolean if a field has been set.

### GetColumnId

`func (o *InlineResponse2002Cards) GetColumnId() string`

GetColumnId returns the ColumnId field if non-nil, zero value otherwise.

### GetColumnIdOk

`func (o *InlineResponse2002Cards) GetColumnIdOk() (*string, bool)`

GetColumnIdOk returns a tuple with the ColumnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnId

`func (o *InlineResponse2002Cards) SetColumnId(v string)`

SetColumnId sets ColumnId field to given value.

### HasColumnId

`func (o *InlineResponse2002Cards) HasColumnId() bool`

HasColumnId returns a boolean if a field has been set.

### GetCompleted

`func (o *InlineResponse2002Cards) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *InlineResponse2002Cards) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *InlineResponse2002Cards) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *InlineResponse2002Cards) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetDateCompleted

`func (o *InlineResponse2002Cards) GetDateCompleted() string`

GetDateCompleted returns the DateCompleted field if non-nil, zero value otherwise.

### GetDateCompletedOk

`func (o *InlineResponse2002Cards) GetDateCompletedOk() (*string, bool)`

GetDateCompletedOk returns a tuple with the DateCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCompleted

`func (o *InlineResponse2002Cards) SetDateCompleted(v string)`

SetDateCompleted sets DateCompleted field to given value.

### HasDateCompleted

`func (o *InlineResponse2002Cards) HasDateCompleted() bool`

HasDateCompleted returns a boolean if a field has been set.

### GetDateCreated

`func (o *InlineResponse2002Cards) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *InlineResponse2002Cards) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *InlineResponse2002Cards) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *InlineResponse2002Cards) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateUpdated

`func (o *InlineResponse2002Cards) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *InlineResponse2002Cards) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *InlineResponse2002Cards) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *InlineResponse2002Cards) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### GetDeleted

`func (o *InlineResponse2002Cards) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *InlineResponse2002Cards) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *InlineResponse2002Cards) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *InlineResponse2002Cards) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDeletedDate

`func (o *InlineResponse2002Cards) GetDeletedDate() string`

GetDeletedDate returns the DeletedDate field if non-nil, zero value otherwise.

### GetDeletedDateOk

`func (o *InlineResponse2002Cards) GetDeletedDateOk() (*string, bool)`

GetDeletedDateOk returns a tuple with the DeletedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDate

`func (o *InlineResponse2002Cards) SetDeletedDate(v string)`

SetDeletedDate sets DeletedDate field to given value.

### HasDeletedDate

`func (o *InlineResponse2002Cards) HasDeletedDate() bool`

HasDeletedDate returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse2002Cards) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse2002Cards) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse2002Cards) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse2002Cards) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *InlineResponse2002Cards) GetDisplayOrder() string`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *InlineResponse2002Cards) GetDisplayOrderOk() (*string, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *InlineResponse2002Cards) SetDisplayOrder(v string)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *InlineResponse2002Cards) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetDueDate

`func (o *InlineResponse2002Cards) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *InlineResponse2002Cards) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *InlineResponse2002Cards) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *InlineResponse2002Cards) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetEstimatedMinutes

`func (o *InlineResponse2002Cards) GetEstimatedMinutes() string`

GetEstimatedMinutes returns the EstimatedMinutes field if non-nil, zero value otherwise.

### GetEstimatedMinutesOk

`func (o *InlineResponse2002Cards) GetEstimatedMinutesOk() (*string, bool)`

GetEstimatedMinutesOk returns a tuple with the EstimatedMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedMinutes

`func (o *InlineResponse2002Cards) SetEstimatedMinutes(v string)`

SetEstimatedMinutes sets EstimatedMinutes field to given value.

### HasEstimatedMinutes

`func (o *InlineResponse2002Cards) HasEstimatedMinutes() bool`

HasEstimatedMinutes returns a boolean if a field has been set.

### GetHasLoggedTime

`func (o *InlineResponse2002Cards) GetHasLoggedTime() bool`

GetHasLoggedTime returns the HasLoggedTime field if non-nil, zero value otherwise.

### GetHasLoggedTimeOk

`func (o *InlineResponse2002Cards) GetHasLoggedTimeOk() (*bool, bool)`

GetHasLoggedTimeOk returns a tuple with the HasLoggedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLoggedTime

`func (o *InlineResponse2002Cards) SetHasLoggedTime(v bool)`

SetHasLoggedTime sets HasLoggedTime field to given value.

### HasHasLoggedTime

`func (o *InlineResponse2002Cards) HasHasLoggedTime() bool`

HasHasLoggedTime returns a boolean if a field has been set.

### GetHasTickets

`func (o *InlineResponse2002Cards) GetHasTickets() bool`

GetHasTickets returns the HasTickets field if non-nil, zero value otherwise.

### GetHasTicketsOk

`func (o *InlineResponse2002Cards) GetHasTicketsOk() (*bool, bool)`

GetHasTicketsOk returns a tuple with the HasTickets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTickets

`func (o *InlineResponse2002Cards) SetHasTickets(v bool)`

SetHasTickets sets HasTickets field to given value.

### HasHasTickets

`func (o *InlineResponse2002Cards) HasHasTickets() bool`

HasHasTickets returns a boolean if a field has been set.

### GetHasUnreadComments

`func (o *InlineResponse2002Cards) GetHasUnreadComments() bool`

GetHasUnreadComments returns the HasUnreadComments field if non-nil, zero value otherwise.

### GetHasUnreadCommentsOk

`func (o *InlineResponse2002Cards) GetHasUnreadCommentsOk() (*bool, bool)`

GetHasUnreadCommentsOk returns a tuple with the HasUnreadComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUnreadComments

`func (o *InlineResponse2002Cards) SetHasUnreadComments(v bool)`

SetHasUnreadComments sets HasUnreadComments field to given value.

### HasHasUnreadComments

`func (o *InlineResponse2002Cards) HasHasUnreadComments() bool`

HasHasUnreadComments returns a boolean if a field has been set.

### GetId

`func (o *InlineResponse2002Cards) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse2002Cards) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse2002Cards) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse2002Cards) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsRecurring

`func (o *InlineResponse2002Cards) GetIsRecurring() bool`

GetIsRecurring returns the IsRecurring field if non-nil, zero value otherwise.

### GetIsRecurringOk

`func (o *InlineResponse2002Cards) GetIsRecurringOk() (*bool, bool)`

GetIsRecurringOk returns a tuple with the IsRecurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecurring

`func (o *InlineResponse2002Cards) SetIsRecurring(v bool)`

SetIsRecurring sets IsRecurring field to given value.

### HasIsRecurring

`func (o *InlineResponse2002Cards) HasIsRecurring() bool`

HasIsRecurring returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse2002Cards) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse2002Cards) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse2002Cards) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse2002Cards) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumActiveSubTasks

`func (o *InlineResponse2002Cards) GetNumActiveSubTasks() string`

GetNumActiveSubTasks returns the NumActiveSubTasks field if non-nil, zero value otherwise.

### GetNumActiveSubTasksOk

`func (o *InlineResponse2002Cards) GetNumActiveSubTasksOk() (*string, bool)`

GetNumActiveSubTasksOk returns a tuple with the NumActiveSubTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumActiveSubTasks

`func (o *InlineResponse2002Cards) SetNumActiveSubTasks(v string)`

SetNumActiveSubTasks sets NumActiveSubTasks field to given value.

### HasNumActiveSubTasks

`func (o *InlineResponse2002Cards) HasNumActiveSubTasks() bool`

HasNumActiveSubTasks returns a boolean if a field has been set.

### GetNumAttachments

`func (o *InlineResponse2002Cards) GetNumAttachments() string`

GetNumAttachments returns the NumAttachments field if non-nil, zero value otherwise.

### GetNumAttachmentsOk

`func (o *InlineResponse2002Cards) GetNumAttachmentsOk() (*string, bool)`

GetNumAttachmentsOk returns a tuple with the NumAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAttachments

`func (o *InlineResponse2002Cards) SetNumAttachments(v string)`

SetNumAttachments sets NumAttachments field to given value.

### HasNumAttachments

`func (o *InlineResponse2002Cards) HasNumAttachments() bool`

HasNumAttachments returns a boolean if a field has been set.

### GetNumComments

`func (o *InlineResponse2002Cards) GetNumComments() string`

GetNumComments returns the NumComments field if non-nil, zero value otherwise.

### GetNumCommentsOk

`func (o *InlineResponse2002Cards) GetNumCommentsOk() (*string, bool)`

GetNumCommentsOk returns a tuple with the NumComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumComments

`func (o *InlineResponse2002Cards) SetNumComments(v string)`

SetNumComments sets NumComments field to given value.

### HasNumComments

`func (o *InlineResponse2002Cards) HasNumComments() bool`

HasNumComments returns a boolean if a field has been set.

### GetNumDependencies

`func (o *InlineResponse2002Cards) GetNumDependencies() string`

GetNumDependencies returns the NumDependencies field if non-nil, zero value otherwise.

### GetNumDependenciesOk

`func (o *InlineResponse2002Cards) GetNumDependenciesOk() (*string, bool)`

GetNumDependenciesOk returns a tuple with the NumDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDependencies

`func (o *InlineResponse2002Cards) SetNumDependencies(v string)`

SetNumDependencies sets NumDependencies field to given value.

### HasNumDependencies

`func (o *InlineResponse2002Cards) HasNumDependencies() bool`

HasNumDependencies returns a boolean if a field has been set.

### GetNumPredecessors

`func (o *InlineResponse2002Cards) GetNumPredecessors() string`

GetNumPredecessors returns the NumPredecessors field if non-nil, zero value otherwise.

### GetNumPredecessorsOk

`func (o *InlineResponse2002Cards) GetNumPredecessorsOk() (*string, bool)`

GetNumPredecessorsOk returns a tuple with the NumPredecessors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPredecessors

`func (o *InlineResponse2002Cards) SetNumPredecessors(v string)`

SetNumPredecessors sets NumPredecessors field to given value.

### HasNumPredecessors

`func (o *InlineResponse2002Cards) HasNumPredecessors() bool`

HasNumPredecessors returns a boolean if a field has been set.

### GetParentTaskId

`func (o *InlineResponse2002Cards) GetParentTaskId() string`

GetParentTaskId returns the ParentTaskId field if non-nil, zero value otherwise.

### GetParentTaskIdOk

`func (o *InlineResponse2002Cards) GetParentTaskIdOk() (*string, bool)`

GetParentTaskIdOk returns a tuple with the ParentTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTaskId

`func (o *InlineResponse2002Cards) SetParentTaskId(v string)`

SetParentTaskId sets ParentTaskId field to given value.

### HasParentTaskId

`func (o *InlineResponse2002Cards) HasParentTaskId() bool`

HasParentTaskId returns a boolean if a field has been set.

### GetPriority

`func (o *InlineResponse2002Cards) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *InlineResponse2002Cards) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *InlineResponse2002Cards) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *InlineResponse2002Cards) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPrivate

`func (o *InlineResponse2002Cards) GetPrivate() string`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *InlineResponse2002Cards) GetPrivateOk() (*string, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *InlineResponse2002Cards) SetPrivate(v string)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *InlineResponse2002Cards) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetProgress

`func (o *InlineResponse2002Cards) GetProgress() string`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *InlineResponse2002Cards) GetProgressOk() (*string, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *InlineResponse2002Cards) SetProgress(v string)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *InlineResponse2002Cards) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetProjectId

`func (o *InlineResponse2002Cards) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *InlineResponse2002Cards) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *InlineResponse2002Cards) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *InlineResponse2002Cards) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetStartDate

`func (o *InlineResponse2002Cards) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InlineResponse2002Cards) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InlineResponse2002Cards) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InlineResponse2002Cards) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse2002Cards) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse2002Cards) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse2002Cards) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse2002Cards) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse2002Cards) GetTags() []map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse2002Cards) GetTagsOk() (*[]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse2002Cards) SetTags(v []map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse2002Cards) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTaskId

`func (o *InlineResponse2002Cards) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *InlineResponse2002Cards) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *InlineResponse2002Cards) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *InlineResponse2002Cards) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetTaskListId

`func (o *InlineResponse2002Cards) GetTaskListId() string`

GetTaskListId returns the TaskListId field if non-nil, zero value otherwise.

### GetTaskListIdOk

`func (o *InlineResponse2002Cards) GetTaskListIdOk() (*string, bool)`

GetTaskListIdOk returns a tuple with the TaskListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskListId

`func (o *InlineResponse2002Cards) SetTaskListId(v string)`

SetTaskListId sets TaskListId field to given value.

### HasTaskListId

`func (o *InlineResponse2002Cards) HasTaskListId() bool`

HasTaskListId returns a boolean if a field has been set.

### GetTaskListName

`func (o *InlineResponse2002Cards) GetTaskListName() string`

GetTaskListName returns the TaskListName field if non-nil, zero value otherwise.

### GetTaskListNameOk

`func (o *InlineResponse2002Cards) GetTaskListNameOk() (*string, bool)`

GetTaskListNameOk returns a tuple with the TaskListName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskListName

`func (o *InlineResponse2002Cards) SetTaskListName(v string)`

SetTaskListName sets TaskListName field to given value.

### HasTaskListName

`func (o *InlineResponse2002Cards) HasTaskListName() bool`

HasTaskListName returns a boolean if a field has been set.

### GetTaskStatus

`func (o *InlineResponse2002Cards) GetTaskStatus() string`

GetTaskStatus returns the TaskStatus field if non-nil, zero value otherwise.

### GetTaskStatusOk

`func (o *InlineResponse2002Cards) GetTaskStatusOk() (*string, bool)`

GetTaskStatusOk returns a tuple with the TaskStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskStatus

`func (o *InlineResponse2002Cards) SetTaskStatus(v string)`

SetTaskStatus sets TaskStatus field to given value.

### HasTaskStatus

`func (o *InlineResponse2002Cards) HasTaskStatus() bool`

HasTaskStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


