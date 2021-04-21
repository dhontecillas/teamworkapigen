# InlineResponse20099TodoItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DLM** | Pointer to **int32** |  | [optional] 
**AttachmentsCount** | Pointer to **int32** |  | [optional] 
**BoardColumn** | Pointer to [**InlineResponse20054Category**](InlineResponse20054Category.md) |  | [optional] 
**CanComplete** | Pointer to **bool** |  | [optional] 
**CanEdit** | Pointer to **bool** |  | [optional] 
**CanLogTime** | Pointer to **bool** |  | [optional] 
**CommentsCount** | Pointer to **int32** |  | [optional] 
**CompanyId** | Pointer to **int32** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**Completed** | Pointer to **bool** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**CreatedOn** | Pointer to **string** |  | [optional] 
**CreatorAvatarUrl** | Pointer to **string** |  | [optional] 
**CreatorFirstname** | Pointer to **string** |  | [optional] 
**CreatorId** | Pointer to **int32** |  | [optional] 
**CreatorLastname** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DueDate** | Pointer to **string** | If a task due date is defined, it will be returned in this field matching due date base. If there is no task due date set, but a milestone is assoicated, it will return the milestone due date.  | [optional] 
**DueDateBase** | Pointer to **string** | If the task has a defined due date, this will be set. Otherwise, it will be empty. | [optional] 
**EstimatedMinutes** | Pointer to **int32** |  | [optional] 
**HarvestEnabled** | Pointer to **bool** |  | [optional] 
**HasDependencies** | Pointer to **int32** |  | [optional] 
**HasReminders** | Pointer to **bool** |  | [optional] 
**HasUnreadComments** | Pointer to **bool** |  | [optional] 
**HasTickets** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**LastChangedOn** | Pointer to **string** |  | [optional] 
**LockdownId** | Pointer to **string** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**ParentTaskId** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**Predecessors** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**Private** | Pointer to **int32** |  | [optional] 
**Progress** | Pointer to **int32** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**ProjectName** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TasklistIsTemplate** | Pointer to **bool** |  | [optional] 
**TasklistLockdownId** | Pointer to **string** |  | [optional] 
**TasklistPrivate** | Pointer to **bool** |  | [optional] 
**TimeIsLogged** | Pointer to **string** |  | [optional] 
**TodoListId** | Pointer to **int32** |  | [optional] 
**TodoListName** | Pointer to **string** |  | [optional] 
**UpdaterFirstname** | Pointer to **string** |  | [optional] 
**UpdaterId** | Pointer to **int32** |  | [optional] 
**UpdaterLastname** | Pointer to **string** |  | [optional] 
**UserFollowingChanges** | Pointer to **bool** |  | [optional] 
**UserFollowingComments** | Pointer to **bool** |  | [optional] 
**ViewEstimatedTime** | Pointer to **bool** |  | [optional] 

## Methods

### NewInlineResponse20099TodoItems

`func NewInlineResponse20099TodoItems() *InlineResponse20099TodoItems`

NewInlineResponse20099TodoItems instantiates a new InlineResponse20099TodoItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20099TodoItemsWithDefaults

`func NewInlineResponse20099TodoItemsWithDefaults() *InlineResponse20099TodoItems`

NewInlineResponse20099TodoItemsWithDefaults instantiates a new InlineResponse20099TodoItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDLM

`func (o *InlineResponse20099TodoItems) GetDLM() int32`

GetDLM returns the DLM field if non-nil, zero value otherwise.

### GetDLMOk

`func (o *InlineResponse20099TodoItems) GetDLMOk() (*int32, bool)`

GetDLMOk returns a tuple with the DLM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDLM

`func (o *InlineResponse20099TodoItems) SetDLM(v int32)`

SetDLM sets DLM field to given value.

### HasDLM

`func (o *InlineResponse20099TodoItems) HasDLM() bool`

HasDLM returns a boolean if a field has been set.

### GetAttachmentsCount

`func (o *InlineResponse20099TodoItems) GetAttachmentsCount() int32`

GetAttachmentsCount returns the AttachmentsCount field if non-nil, zero value otherwise.

### GetAttachmentsCountOk

`func (o *InlineResponse20099TodoItems) GetAttachmentsCountOk() (*int32, bool)`

GetAttachmentsCountOk returns a tuple with the AttachmentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentsCount

`func (o *InlineResponse20099TodoItems) SetAttachmentsCount(v int32)`

SetAttachmentsCount sets AttachmentsCount field to given value.

### HasAttachmentsCount

`func (o *InlineResponse20099TodoItems) HasAttachmentsCount() bool`

HasAttachmentsCount returns a boolean if a field has been set.

### GetBoardColumn

`func (o *InlineResponse20099TodoItems) GetBoardColumn() InlineResponse20054Category`

GetBoardColumn returns the BoardColumn field if non-nil, zero value otherwise.

### GetBoardColumnOk

`func (o *InlineResponse20099TodoItems) GetBoardColumnOk() (*InlineResponse20054Category, bool)`

GetBoardColumnOk returns a tuple with the BoardColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardColumn

`func (o *InlineResponse20099TodoItems) SetBoardColumn(v InlineResponse20054Category)`

SetBoardColumn sets BoardColumn field to given value.

### HasBoardColumn

`func (o *InlineResponse20099TodoItems) HasBoardColumn() bool`

HasBoardColumn returns a boolean if a field has been set.

### GetCanComplete

`func (o *InlineResponse20099TodoItems) GetCanComplete() bool`

GetCanComplete returns the CanComplete field if non-nil, zero value otherwise.

### GetCanCompleteOk

`func (o *InlineResponse20099TodoItems) GetCanCompleteOk() (*bool, bool)`

GetCanCompleteOk returns a tuple with the CanComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanComplete

`func (o *InlineResponse20099TodoItems) SetCanComplete(v bool)`

SetCanComplete sets CanComplete field to given value.

### HasCanComplete

`func (o *InlineResponse20099TodoItems) HasCanComplete() bool`

HasCanComplete returns a boolean if a field has been set.

### GetCanEdit

`func (o *InlineResponse20099TodoItems) GetCanEdit() bool`

GetCanEdit returns the CanEdit field if non-nil, zero value otherwise.

### GetCanEditOk

`func (o *InlineResponse20099TodoItems) GetCanEditOk() (*bool, bool)`

GetCanEditOk returns a tuple with the CanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEdit

`func (o *InlineResponse20099TodoItems) SetCanEdit(v bool)`

SetCanEdit sets CanEdit field to given value.

### HasCanEdit

`func (o *InlineResponse20099TodoItems) HasCanEdit() bool`

HasCanEdit returns a boolean if a field has been set.

### GetCanLogTime

`func (o *InlineResponse20099TodoItems) GetCanLogTime() bool`

GetCanLogTime returns the CanLogTime field if non-nil, zero value otherwise.

### GetCanLogTimeOk

`func (o *InlineResponse20099TodoItems) GetCanLogTimeOk() (*bool, bool)`

GetCanLogTimeOk returns a tuple with the CanLogTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanLogTime

`func (o *InlineResponse20099TodoItems) SetCanLogTime(v bool)`

SetCanLogTime sets CanLogTime field to given value.

### HasCanLogTime

`func (o *InlineResponse20099TodoItems) HasCanLogTime() bool`

HasCanLogTime returns a boolean if a field has been set.

### GetCommentsCount

`func (o *InlineResponse20099TodoItems) GetCommentsCount() int32`

GetCommentsCount returns the CommentsCount field if non-nil, zero value otherwise.

### GetCommentsCountOk

`func (o *InlineResponse20099TodoItems) GetCommentsCountOk() (*int32, bool)`

GetCommentsCountOk returns a tuple with the CommentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsCount

`func (o *InlineResponse20099TodoItems) SetCommentsCount(v int32)`

SetCommentsCount sets CommentsCount field to given value.

### HasCommentsCount

`func (o *InlineResponse20099TodoItems) HasCommentsCount() bool`

HasCommentsCount returns a boolean if a field has been set.

### GetCompanyId

`func (o *InlineResponse20099TodoItems) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *InlineResponse20099TodoItems) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *InlineResponse20099TodoItems) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *InlineResponse20099TodoItems) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCompanyName

`func (o *InlineResponse20099TodoItems) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *InlineResponse20099TodoItems) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *InlineResponse20099TodoItems) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *InlineResponse20099TodoItems) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCompleted

`func (o *InlineResponse20099TodoItems) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *InlineResponse20099TodoItems) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *InlineResponse20099TodoItems) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *InlineResponse20099TodoItems) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetContent

`func (o *InlineResponse20099TodoItems) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *InlineResponse20099TodoItems) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *InlineResponse20099TodoItems) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *InlineResponse20099TodoItems) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedOn

`func (o *InlineResponse20099TodoItems) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *InlineResponse20099TodoItems) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *InlineResponse20099TodoItems) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *InlineResponse20099TodoItems) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetCreatorAvatarUrl

`func (o *InlineResponse20099TodoItems) GetCreatorAvatarUrl() string`

GetCreatorAvatarUrl returns the CreatorAvatarUrl field if non-nil, zero value otherwise.

### GetCreatorAvatarUrlOk

`func (o *InlineResponse20099TodoItems) GetCreatorAvatarUrlOk() (*string, bool)`

GetCreatorAvatarUrlOk returns a tuple with the CreatorAvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorAvatarUrl

`func (o *InlineResponse20099TodoItems) SetCreatorAvatarUrl(v string)`

SetCreatorAvatarUrl sets CreatorAvatarUrl field to given value.

### HasCreatorAvatarUrl

`func (o *InlineResponse20099TodoItems) HasCreatorAvatarUrl() bool`

HasCreatorAvatarUrl returns a boolean if a field has been set.

### GetCreatorFirstname

`func (o *InlineResponse20099TodoItems) GetCreatorFirstname() string`

GetCreatorFirstname returns the CreatorFirstname field if non-nil, zero value otherwise.

### GetCreatorFirstnameOk

`func (o *InlineResponse20099TodoItems) GetCreatorFirstnameOk() (*string, bool)`

GetCreatorFirstnameOk returns a tuple with the CreatorFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorFirstname

`func (o *InlineResponse20099TodoItems) SetCreatorFirstname(v string)`

SetCreatorFirstname sets CreatorFirstname field to given value.

### HasCreatorFirstname

`func (o *InlineResponse20099TodoItems) HasCreatorFirstname() bool`

HasCreatorFirstname returns a boolean if a field has been set.

### GetCreatorId

`func (o *InlineResponse20099TodoItems) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *InlineResponse20099TodoItems) GetCreatorIdOk() (*int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *InlineResponse20099TodoItems) SetCreatorId(v int32)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *InlineResponse20099TodoItems) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetCreatorLastname

`func (o *InlineResponse20099TodoItems) GetCreatorLastname() string`

GetCreatorLastname returns the CreatorLastname field if non-nil, zero value otherwise.

### GetCreatorLastnameOk

`func (o *InlineResponse20099TodoItems) GetCreatorLastnameOk() (*string, bool)`

GetCreatorLastnameOk returns a tuple with the CreatorLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorLastname

`func (o *InlineResponse20099TodoItems) SetCreatorLastname(v string)`

SetCreatorLastname sets CreatorLastname field to given value.

### HasCreatorLastname

`func (o *InlineResponse20099TodoItems) HasCreatorLastname() bool`

HasCreatorLastname returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse20099TodoItems) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse20099TodoItems) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse20099TodoItems) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse20099TodoItems) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDueDate

`func (o *InlineResponse20099TodoItems) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *InlineResponse20099TodoItems) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *InlineResponse20099TodoItems) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *InlineResponse20099TodoItems) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetDueDateBase

`func (o *InlineResponse20099TodoItems) GetDueDateBase() string`

GetDueDateBase returns the DueDateBase field if non-nil, zero value otherwise.

### GetDueDateBaseOk

`func (o *InlineResponse20099TodoItems) GetDueDateBaseOk() (*string, bool)`

GetDueDateBaseOk returns a tuple with the DueDateBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDateBase

`func (o *InlineResponse20099TodoItems) SetDueDateBase(v string)`

SetDueDateBase sets DueDateBase field to given value.

### HasDueDateBase

`func (o *InlineResponse20099TodoItems) HasDueDateBase() bool`

HasDueDateBase returns a boolean if a field has been set.

### GetEstimatedMinutes

`func (o *InlineResponse20099TodoItems) GetEstimatedMinutes() int32`

GetEstimatedMinutes returns the EstimatedMinutes field if non-nil, zero value otherwise.

### GetEstimatedMinutesOk

`func (o *InlineResponse20099TodoItems) GetEstimatedMinutesOk() (*int32, bool)`

GetEstimatedMinutesOk returns a tuple with the EstimatedMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedMinutes

`func (o *InlineResponse20099TodoItems) SetEstimatedMinutes(v int32)`

SetEstimatedMinutes sets EstimatedMinutes field to given value.

### HasEstimatedMinutes

`func (o *InlineResponse20099TodoItems) HasEstimatedMinutes() bool`

HasEstimatedMinutes returns a boolean if a field has been set.

### GetHarvestEnabled

`func (o *InlineResponse20099TodoItems) GetHarvestEnabled() bool`

GetHarvestEnabled returns the HarvestEnabled field if non-nil, zero value otherwise.

### GetHarvestEnabledOk

`func (o *InlineResponse20099TodoItems) GetHarvestEnabledOk() (*bool, bool)`

GetHarvestEnabledOk returns a tuple with the HarvestEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHarvestEnabled

`func (o *InlineResponse20099TodoItems) SetHarvestEnabled(v bool)`

SetHarvestEnabled sets HarvestEnabled field to given value.

### HasHarvestEnabled

`func (o *InlineResponse20099TodoItems) HasHarvestEnabled() bool`

HasHarvestEnabled returns a boolean if a field has been set.

### GetHasDependencies

`func (o *InlineResponse20099TodoItems) GetHasDependencies() int32`

GetHasDependencies returns the HasDependencies field if non-nil, zero value otherwise.

### GetHasDependenciesOk

`func (o *InlineResponse20099TodoItems) GetHasDependenciesOk() (*int32, bool)`

GetHasDependenciesOk returns a tuple with the HasDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDependencies

`func (o *InlineResponse20099TodoItems) SetHasDependencies(v int32)`

SetHasDependencies sets HasDependencies field to given value.

### HasHasDependencies

`func (o *InlineResponse20099TodoItems) HasHasDependencies() bool`

HasHasDependencies returns a boolean if a field has been set.

### GetHasReminders

`func (o *InlineResponse20099TodoItems) GetHasReminders() bool`

GetHasReminders returns the HasReminders field if non-nil, zero value otherwise.

### GetHasRemindersOk

`func (o *InlineResponse20099TodoItems) GetHasRemindersOk() (*bool, bool)`

GetHasRemindersOk returns a tuple with the HasReminders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasReminders

`func (o *InlineResponse20099TodoItems) SetHasReminders(v bool)`

SetHasReminders sets HasReminders field to given value.

### HasHasReminders

`func (o *InlineResponse20099TodoItems) HasHasReminders() bool`

HasHasReminders returns a boolean if a field has been set.

### GetHasUnreadComments

`func (o *InlineResponse20099TodoItems) GetHasUnreadComments() bool`

GetHasUnreadComments returns the HasUnreadComments field if non-nil, zero value otherwise.

### GetHasUnreadCommentsOk

`func (o *InlineResponse20099TodoItems) GetHasUnreadCommentsOk() (*bool, bool)`

GetHasUnreadCommentsOk returns a tuple with the HasUnreadComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUnreadComments

`func (o *InlineResponse20099TodoItems) SetHasUnreadComments(v bool)`

SetHasUnreadComments sets HasUnreadComments field to given value.

### HasHasUnreadComments

`func (o *InlineResponse20099TodoItems) HasHasUnreadComments() bool`

HasHasUnreadComments returns a boolean if a field has been set.

### GetHasTickets

`func (o *InlineResponse20099TodoItems) GetHasTickets() bool`

GetHasTickets returns the HasTickets field if non-nil, zero value otherwise.

### GetHasTicketsOk

`func (o *InlineResponse20099TodoItems) GetHasTicketsOk() (*bool, bool)`

GetHasTicketsOk returns a tuple with the HasTickets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTickets

`func (o *InlineResponse20099TodoItems) SetHasTickets(v bool)`

SetHasTickets sets HasTickets field to given value.

### HasHasTickets

`func (o *InlineResponse20099TodoItems) HasHasTickets() bool`

HasHasTickets returns a boolean if a field has been set.

### GetId

`func (o *InlineResponse20099TodoItems) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20099TodoItems) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20099TodoItems) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20099TodoItems) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastChangedOn

`func (o *InlineResponse20099TodoItems) GetLastChangedOn() string`

GetLastChangedOn returns the LastChangedOn field if non-nil, zero value otherwise.

### GetLastChangedOnOk

`func (o *InlineResponse20099TodoItems) GetLastChangedOnOk() (*string, bool)`

GetLastChangedOnOk returns a tuple with the LastChangedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangedOn

`func (o *InlineResponse20099TodoItems) SetLastChangedOn(v string)`

SetLastChangedOn sets LastChangedOn field to given value.

### HasLastChangedOn

`func (o *InlineResponse20099TodoItems) HasLastChangedOn() bool`

HasLastChangedOn returns a boolean if a field has been set.

### GetLockdownId

`func (o *InlineResponse20099TodoItems) GetLockdownId() string`

GetLockdownId returns the LockdownId field if non-nil, zero value otherwise.

### GetLockdownIdOk

`func (o *InlineResponse20099TodoItems) GetLockdownIdOk() (*string, bool)`

GetLockdownIdOk returns a tuple with the LockdownId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockdownId

`func (o *InlineResponse20099TodoItems) SetLockdownId(v string)`

SetLockdownId sets LockdownId field to given value.

### HasLockdownId

`func (o *InlineResponse20099TodoItems) HasLockdownId() bool`

HasLockdownId returns a boolean if a field has been set.

### GetOrder

`func (o *InlineResponse20099TodoItems) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *InlineResponse20099TodoItems) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *InlineResponse20099TodoItems) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *InlineResponse20099TodoItems) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetParentTaskId

`func (o *InlineResponse20099TodoItems) GetParentTaskId() string`

GetParentTaskId returns the ParentTaskId field if non-nil, zero value otherwise.

### GetParentTaskIdOk

`func (o *InlineResponse20099TodoItems) GetParentTaskIdOk() (*string, bool)`

GetParentTaskIdOk returns a tuple with the ParentTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTaskId

`func (o *InlineResponse20099TodoItems) SetParentTaskId(v string)`

SetParentTaskId sets ParentTaskId field to given value.

### HasParentTaskId

`func (o *InlineResponse20099TodoItems) HasParentTaskId() bool`

HasParentTaskId returns a boolean if a field has been set.

### GetPosition

`func (o *InlineResponse20099TodoItems) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *InlineResponse20099TodoItems) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *InlineResponse20099TodoItems) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *InlineResponse20099TodoItems) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPredecessors

`func (o *InlineResponse20099TodoItems) GetPredecessors() []map[string]interface{}`

GetPredecessors returns the Predecessors field if non-nil, zero value otherwise.

### GetPredecessorsOk

`func (o *InlineResponse20099TodoItems) GetPredecessorsOk() (*[]map[string]interface{}, bool)`

GetPredecessorsOk returns a tuple with the Predecessors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredecessors

`func (o *InlineResponse20099TodoItems) SetPredecessors(v []map[string]interface{})`

SetPredecessors sets Predecessors field to given value.

### HasPredecessors

`func (o *InlineResponse20099TodoItems) HasPredecessors() bool`

HasPredecessors returns a boolean if a field has been set.

### GetPriority

`func (o *InlineResponse20099TodoItems) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *InlineResponse20099TodoItems) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *InlineResponse20099TodoItems) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *InlineResponse20099TodoItems) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPrivate

`func (o *InlineResponse20099TodoItems) GetPrivate() int32`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *InlineResponse20099TodoItems) GetPrivateOk() (*int32, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *InlineResponse20099TodoItems) SetPrivate(v int32)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *InlineResponse20099TodoItems) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetProgress

`func (o *InlineResponse20099TodoItems) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *InlineResponse20099TodoItems) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *InlineResponse20099TodoItems) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *InlineResponse20099TodoItems) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetProjectId

`func (o *InlineResponse20099TodoItems) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *InlineResponse20099TodoItems) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *InlineResponse20099TodoItems) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *InlineResponse20099TodoItems) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProjectName

`func (o *InlineResponse20099TodoItems) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *InlineResponse20099TodoItems) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *InlineResponse20099TodoItems) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *InlineResponse20099TodoItems) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetStartDate

`func (o *InlineResponse20099TodoItems) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InlineResponse20099TodoItems) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InlineResponse20099TodoItems) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InlineResponse20099TodoItems) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20099TodoItems) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20099TodoItems) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20099TodoItems) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20099TodoItems) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTasklistIsTemplate

`func (o *InlineResponse20099TodoItems) GetTasklistIsTemplate() bool`

GetTasklistIsTemplate returns the TasklistIsTemplate field if non-nil, zero value otherwise.

### GetTasklistIsTemplateOk

`func (o *InlineResponse20099TodoItems) GetTasklistIsTemplateOk() (*bool, bool)`

GetTasklistIsTemplateOk returns a tuple with the TasklistIsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasklistIsTemplate

`func (o *InlineResponse20099TodoItems) SetTasklistIsTemplate(v bool)`

SetTasklistIsTemplate sets TasklistIsTemplate field to given value.

### HasTasklistIsTemplate

`func (o *InlineResponse20099TodoItems) HasTasklistIsTemplate() bool`

HasTasklistIsTemplate returns a boolean if a field has been set.

### GetTasklistLockdownId

`func (o *InlineResponse20099TodoItems) GetTasklistLockdownId() string`

GetTasklistLockdownId returns the TasklistLockdownId field if non-nil, zero value otherwise.

### GetTasklistLockdownIdOk

`func (o *InlineResponse20099TodoItems) GetTasklistLockdownIdOk() (*string, bool)`

GetTasklistLockdownIdOk returns a tuple with the TasklistLockdownId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasklistLockdownId

`func (o *InlineResponse20099TodoItems) SetTasklistLockdownId(v string)`

SetTasklistLockdownId sets TasklistLockdownId field to given value.

### HasTasklistLockdownId

`func (o *InlineResponse20099TodoItems) HasTasklistLockdownId() bool`

HasTasklistLockdownId returns a boolean if a field has been set.

### GetTasklistPrivate

`func (o *InlineResponse20099TodoItems) GetTasklistPrivate() bool`

GetTasklistPrivate returns the TasklistPrivate field if non-nil, zero value otherwise.

### GetTasklistPrivateOk

`func (o *InlineResponse20099TodoItems) GetTasklistPrivateOk() (*bool, bool)`

GetTasklistPrivateOk returns a tuple with the TasklistPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasklistPrivate

`func (o *InlineResponse20099TodoItems) SetTasklistPrivate(v bool)`

SetTasklistPrivate sets TasklistPrivate field to given value.

### HasTasklistPrivate

`func (o *InlineResponse20099TodoItems) HasTasklistPrivate() bool`

HasTasklistPrivate returns a boolean if a field has been set.

### GetTimeIsLogged

`func (o *InlineResponse20099TodoItems) GetTimeIsLogged() string`

GetTimeIsLogged returns the TimeIsLogged field if non-nil, zero value otherwise.

### GetTimeIsLoggedOk

`func (o *InlineResponse20099TodoItems) GetTimeIsLoggedOk() (*string, bool)`

GetTimeIsLoggedOk returns a tuple with the TimeIsLogged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeIsLogged

`func (o *InlineResponse20099TodoItems) SetTimeIsLogged(v string)`

SetTimeIsLogged sets TimeIsLogged field to given value.

### HasTimeIsLogged

`func (o *InlineResponse20099TodoItems) HasTimeIsLogged() bool`

HasTimeIsLogged returns a boolean if a field has been set.

### GetTodoListId

`func (o *InlineResponse20099TodoItems) GetTodoListId() int32`

GetTodoListId returns the TodoListId field if non-nil, zero value otherwise.

### GetTodoListIdOk

`func (o *InlineResponse20099TodoItems) GetTodoListIdOk() (*int32, bool)`

GetTodoListIdOk returns a tuple with the TodoListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodoListId

`func (o *InlineResponse20099TodoItems) SetTodoListId(v int32)`

SetTodoListId sets TodoListId field to given value.

### HasTodoListId

`func (o *InlineResponse20099TodoItems) HasTodoListId() bool`

HasTodoListId returns a boolean if a field has been set.

### GetTodoListName

`func (o *InlineResponse20099TodoItems) GetTodoListName() string`

GetTodoListName returns the TodoListName field if non-nil, zero value otherwise.

### GetTodoListNameOk

`func (o *InlineResponse20099TodoItems) GetTodoListNameOk() (*string, bool)`

GetTodoListNameOk returns a tuple with the TodoListName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodoListName

`func (o *InlineResponse20099TodoItems) SetTodoListName(v string)`

SetTodoListName sets TodoListName field to given value.

### HasTodoListName

`func (o *InlineResponse20099TodoItems) HasTodoListName() bool`

HasTodoListName returns a boolean if a field has been set.

### GetUpdaterFirstname

`func (o *InlineResponse20099TodoItems) GetUpdaterFirstname() string`

GetUpdaterFirstname returns the UpdaterFirstname field if non-nil, zero value otherwise.

### GetUpdaterFirstnameOk

`func (o *InlineResponse20099TodoItems) GetUpdaterFirstnameOk() (*string, bool)`

GetUpdaterFirstnameOk returns a tuple with the UpdaterFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdaterFirstname

`func (o *InlineResponse20099TodoItems) SetUpdaterFirstname(v string)`

SetUpdaterFirstname sets UpdaterFirstname field to given value.

### HasUpdaterFirstname

`func (o *InlineResponse20099TodoItems) HasUpdaterFirstname() bool`

HasUpdaterFirstname returns a boolean if a field has been set.

### GetUpdaterId

`func (o *InlineResponse20099TodoItems) GetUpdaterId() int32`

GetUpdaterId returns the UpdaterId field if non-nil, zero value otherwise.

### GetUpdaterIdOk

`func (o *InlineResponse20099TodoItems) GetUpdaterIdOk() (*int32, bool)`

GetUpdaterIdOk returns a tuple with the UpdaterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdaterId

`func (o *InlineResponse20099TodoItems) SetUpdaterId(v int32)`

SetUpdaterId sets UpdaterId field to given value.

### HasUpdaterId

`func (o *InlineResponse20099TodoItems) HasUpdaterId() bool`

HasUpdaterId returns a boolean if a field has been set.

### GetUpdaterLastname

`func (o *InlineResponse20099TodoItems) GetUpdaterLastname() string`

GetUpdaterLastname returns the UpdaterLastname field if non-nil, zero value otherwise.

### GetUpdaterLastnameOk

`func (o *InlineResponse20099TodoItems) GetUpdaterLastnameOk() (*string, bool)`

GetUpdaterLastnameOk returns a tuple with the UpdaterLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdaterLastname

`func (o *InlineResponse20099TodoItems) SetUpdaterLastname(v string)`

SetUpdaterLastname sets UpdaterLastname field to given value.

### HasUpdaterLastname

`func (o *InlineResponse20099TodoItems) HasUpdaterLastname() bool`

HasUpdaterLastname returns a boolean if a field has been set.

### GetUserFollowingChanges

`func (o *InlineResponse20099TodoItems) GetUserFollowingChanges() bool`

GetUserFollowingChanges returns the UserFollowingChanges field if non-nil, zero value otherwise.

### GetUserFollowingChangesOk

`func (o *InlineResponse20099TodoItems) GetUserFollowingChangesOk() (*bool, bool)`

GetUserFollowingChangesOk returns a tuple with the UserFollowingChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFollowingChanges

`func (o *InlineResponse20099TodoItems) SetUserFollowingChanges(v bool)`

SetUserFollowingChanges sets UserFollowingChanges field to given value.

### HasUserFollowingChanges

`func (o *InlineResponse20099TodoItems) HasUserFollowingChanges() bool`

HasUserFollowingChanges returns a boolean if a field has been set.

### GetUserFollowingComments

`func (o *InlineResponse20099TodoItems) GetUserFollowingComments() bool`

GetUserFollowingComments returns the UserFollowingComments field if non-nil, zero value otherwise.

### GetUserFollowingCommentsOk

`func (o *InlineResponse20099TodoItems) GetUserFollowingCommentsOk() (*bool, bool)`

GetUserFollowingCommentsOk returns a tuple with the UserFollowingComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFollowingComments

`func (o *InlineResponse20099TodoItems) SetUserFollowingComments(v bool)`

SetUserFollowingComments sets UserFollowingComments field to given value.

### HasUserFollowingComments

`func (o *InlineResponse20099TodoItems) HasUserFollowingComments() bool`

HasUserFollowingComments returns a boolean if a field has been set.

### GetViewEstimatedTime

`func (o *InlineResponse20099TodoItems) GetViewEstimatedTime() bool`

GetViewEstimatedTime returns the ViewEstimatedTime field if non-nil, zero value otherwise.

### GetViewEstimatedTimeOk

`func (o *InlineResponse20099TodoItems) GetViewEstimatedTimeOk() (*bool, bool)`

GetViewEstimatedTimeOk returns a tuple with the ViewEstimatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewEstimatedTime

`func (o *InlineResponse20099TodoItems) SetViewEstimatedTime(v bool)`

SetViewEstimatedTime sets ViewEstimatedTime field to given value.

### HasViewEstimatedTime

`func (o *InlineResponse20099TodoItems) HasViewEstimatedTime() bool`

HasViewEstimatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


