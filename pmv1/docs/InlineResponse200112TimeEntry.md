# InlineResponse200112TimeEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanEdit** | Pointer to **bool** |  | [optional] 
**CompanyId** | Pointer to **string** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 
**DateUserPerspective** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**HasStartTime** | Pointer to **string** |  | [optional] 
**Hours** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InvoiceNo** | Pointer to **string** |  | [optional] 
**Isbillable** | Pointer to **string** |  | [optional] 
**Isbilled** | Pointer to **string** |  | [optional] 
**Minutes** | Pointer to **string** |  | [optional] 
**ParentTaskId** | Pointer to **string** |  | [optional] 
**ParentTaskName** | Pointer to **string** |  | [optional] 
**PersonFirstName** | Pointer to **string** |  | [optional] 
**PersonId** | Pointer to **string** |  | [optional] 
**PersonLastName** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**ProjectName** | Pointer to **string** |  | [optional] 
**ProjectStatus** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]InlineResponse2002Column**](InlineResponse2002Column.md) |  | [optional] 
**TaskTags** | Pointer to **[]map[string]interface{}** |  | [optional] 
**TaskEstimatedTime** | Pointer to **string** |  | [optional] 
**TaskIsPrivate** | Pointer to **string** |  | [optional] 
**TaskIsSubTask** | Pointer to **string** |  | [optional] 
**TasklistId** | Pointer to **string** |  | [optional] 
**TicketId** | Pointer to **string** |  | [optional] 
**TodoItemId** | Pointer to **string** |  | [optional] 
**TodoItemName** | Pointer to **string** |  | [optional] 
**TodoListId** | Pointer to **string** |  | [optional] 
**TodoListName** | Pointer to **string** |  | [optional] 
**UpdatedDate** | Pointer to **string** |  | [optional] 

## Methods

### NewInlineResponse200112TimeEntry

`func NewInlineResponse200112TimeEntry() *InlineResponse200112TimeEntry`

NewInlineResponse200112TimeEntry instantiates a new InlineResponse200112TimeEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200112TimeEntryWithDefaults

`func NewInlineResponse200112TimeEntryWithDefaults() *InlineResponse200112TimeEntry`

NewInlineResponse200112TimeEntryWithDefaults instantiates a new InlineResponse200112TimeEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanEdit

`func (o *InlineResponse200112TimeEntry) GetCanEdit() bool`

GetCanEdit returns the CanEdit field if non-nil, zero value otherwise.

### GetCanEditOk

`func (o *InlineResponse200112TimeEntry) GetCanEditOk() (*bool, bool)`

GetCanEditOk returns a tuple with the CanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEdit

`func (o *InlineResponse200112TimeEntry) SetCanEdit(v bool)`

SetCanEdit sets CanEdit field to given value.

### HasCanEdit

`func (o *InlineResponse200112TimeEntry) HasCanEdit() bool`

HasCanEdit returns a boolean if a field has been set.

### GetCompanyId

`func (o *InlineResponse200112TimeEntry) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *InlineResponse200112TimeEntry) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *InlineResponse200112TimeEntry) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *InlineResponse200112TimeEntry) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCompanyName

`func (o *InlineResponse200112TimeEntry) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *InlineResponse200112TimeEntry) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *InlineResponse200112TimeEntry) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *InlineResponse200112TimeEntry) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse200112TimeEntry) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse200112TimeEntry) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse200112TimeEntry) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse200112TimeEntry) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDate

`func (o *InlineResponse200112TimeEntry) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *InlineResponse200112TimeEntry) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *InlineResponse200112TimeEntry) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *InlineResponse200112TimeEntry) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDateUserPerspective

`func (o *InlineResponse200112TimeEntry) GetDateUserPerspective() string`

GetDateUserPerspective returns the DateUserPerspective field if non-nil, zero value otherwise.

### GetDateUserPerspectiveOk

`func (o *InlineResponse200112TimeEntry) GetDateUserPerspectiveOk() (*string, bool)`

GetDateUserPerspectiveOk returns a tuple with the DateUserPerspective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUserPerspective

`func (o *InlineResponse200112TimeEntry) SetDateUserPerspective(v string)`

SetDateUserPerspective sets DateUserPerspective field to given value.

### HasDateUserPerspective

`func (o *InlineResponse200112TimeEntry) HasDateUserPerspective() bool`

HasDateUserPerspective returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse200112TimeEntry) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse200112TimeEntry) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse200112TimeEntry) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse200112TimeEntry) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHasStartTime

`func (o *InlineResponse200112TimeEntry) GetHasStartTime() string`

GetHasStartTime returns the HasStartTime field if non-nil, zero value otherwise.

### GetHasStartTimeOk

`func (o *InlineResponse200112TimeEntry) GetHasStartTimeOk() (*string, bool)`

GetHasStartTimeOk returns a tuple with the HasStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasStartTime

`func (o *InlineResponse200112TimeEntry) SetHasStartTime(v string)`

SetHasStartTime sets HasStartTime field to given value.

### HasHasStartTime

`func (o *InlineResponse200112TimeEntry) HasHasStartTime() bool`

HasHasStartTime returns a boolean if a field has been set.

### GetHours

`func (o *InlineResponse200112TimeEntry) GetHours() string`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *InlineResponse200112TimeEntry) GetHoursOk() (*string, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *InlineResponse200112TimeEntry) SetHours(v string)`

SetHours sets Hours field to given value.

### HasHours

`func (o *InlineResponse200112TimeEntry) HasHours() bool`

HasHours returns a boolean if a field has been set.

### GetId

`func (o *InlineResponse200112TimeEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200112TimeEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200112TimeEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200112TimeEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceNo

`func (o *InlineResponse200112TimeEntry) GetInvoiceNo() string`

GetInvoiceNo returns the InvoiceNo field if non-nil, zero value otherwise.

### GetInvoiceNoOk

`func (o *InlineResponse200112TimeEntry) GetInvoiceNoOk() (*string, bool)`

GetInvoiceNoOk returns a tuple with the InvoiceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNo

`func (o *InlineResponse200112TimeEntry) SetInvoiceNo(v string)`

SetInvoiceNo sets InvoiceNo field to given value.

### HasInvoiceNo

`func (o *InlineResponse200112TimeEntry) HasInvoiceNo() bool`

HasInvoiceNo returns a boolean if a field has been set.

### GetIsbillable

`func (o *InlineResponse200112TimeEntry) GetIsbillable() string`

GetIsbillable returns the Isbillable field if non-nil, zero value otherwise.

### GetIsbillableOk

`func (o *InlineResponse200112TimeEntry) GetIsbillableOk() (*string, bool)`

GetIsbillableOk returns a tuple with the Isbillable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsbillable

`func (o *InlineResponse200112TimeEntry) SetIsbillable(v string)`

SetIsbillable sets Isbillable field to given value.

### HasIsbillable

`func (o *InlineResponse200112TimeEntry) HasIsbillable() bool`

HasIsbillable returns a boolean if a field has been set.

### GetIsbilled

`func (o *InlineResponse200112TimeEntry) GetIsbilled() string`

GetIsbilled returns the Isbilled field if non-nil, zero value otherwise.

### GetIsbilledOk

`func (o *InlineResponse200112TimeEntry) GetIsbilledOk() (*string, bool)`

GetIsbilledOk returns a tuple with the Isbilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsbilled

`func (o *InlineResponse200112TimeEntry) SetIsbilled(v string)`

SetIsbilled sets Isbilled field to given value.

### HasIsbilled

`func (o *InlineResponse200112TimeEntry) HasIsbilled() bool`

HasIsbilled returns a boolean if a field has been set.

### GetMinutes

`func (o *InlineResponse200112TimeEntry) GetMinutes() string`

GetMinutes returns the Minutes field if non-nil, zero value otherwise.

### GetMinutesOk

`func (o *InlineResponse200112TimeEntry) GetMinutesOk() (*string, bool)`

GetMinutesOk returns a tuple with the Minutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutes

`func (o *InlineResponse200112TimeEntry) SetMinutes(v string)`

SetMinutes sets Minutes field to given value.

### HasMinutes

`func (o *InlineResponse200112TimeEntry) HasMinutes() bool`

HasMinutes returns a boolean if a field has been set.

### GetParentTaskId

`func (o *InlineResponse200112TimeEntry) GetParentTaskId() string`

GetParentTaskId returns the ParentTaskId field if non-nil, zero value otherwise.

### GetParentTaskIdOk

`func (o *InlineResponse200112TimeEntry) GetParentTaskIdOk() (*string, bool)`

GetParentTaskIdOk returns a tuple with the ParentTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTaskId

`func (o *InlineResponse200112TimeEntry) SetParentTaskId(v string)`

SetParentTaskId sets ParentTaskId field to given value.

### HasParentTaskId

`func (o *InlineResponse200112TimeEntry) HasParentTaskId() bool`

HasParentTaskId returns a boolean if a field has been set.

### GetParentTaskName

`func (o *InlineResponse200112TimeEntry) GetParentTaskName() string`

GetParentTaskName returns the ParentTaskName field if non-nil, zero value otherwise.

### GetParentTaskNameOk

`func (o *InlineResponse200112TimeEntry) GetParentTaskNameOk() (*string, bool)`

GetParentTaskNameOk returns a tuple with the ParentTaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTaskName

`func (o *InlineResponse200112TimeEntry) SetParentTaskName(v string)`

SetParentTaskName sets ParentTaskName field to given value.

### HasParentTaskName

`func (o *InlineResponse200112TimeEntry) HasParentTaskName() bool`

HasParentTaskName returns a boolean if a field has been set.

### GetPersonFirstName

`func (o *InlineResponse200112TimeEntry) GetPersonFirstName() string`

GetPersonFirstName returns the PersonFirstName field if non-nil, zero value otherwise.

### GetPersonFirstNameOk

`func (o *InlineResponse200112TimeEntry) GetPersonFirstNameOk() (*string, bool)`

GetPersonFirstNameOk returns a tuple with the PersonFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonFirstName

`func (o *InlineResponse200112TimeEntry) SetPersonFirstName(v string)`

SetPersonFirstName sets PersonFirstName field to given value.

### HasPersonFirstName

`func (o *InlineResponse200112TimeEntry) HasPersonFirstName() bool`

HasPersonFirstName returns a boolean if a field has been set.

### GetPersonId

`func (o *InlineResponse200112TimeEntry) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *InlineResponse200112TimeEntry) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *InlineResponse200112TimeEntry) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *InlineResponse200112TimeEntry) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### GetPersonLastName

`func (o *InlineResponse200112TimeEntry) GetPersonLastName() string`

GetPersonLastName returns the PersonLastName field if non-nil, zero value otherwise.

### GetPersonLastNameOk

`func (o *InlineResponse200112TimeEntry) GetPersonLastNameOk() (*string, bool)`

GetPersonLastNameOk returns a tuple with the PersonLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonLastName

`func (o *InlineResponse200112TimeEntry) SetPersonLastName(v string)`

SetPersonLastName sets PersonLastName field to given value.

### HasPersonLastName

`func (o *InlineResponse200112TimeEntry) HasPersonLastName() bool`

HasPersonLastName returns a boolean if a field has been set.

### GetProjectId

`func (o *InlineResponse200112TimeEntry) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *InlineResponse200112TimeEntry) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *InlineResponse200112TimeEntry) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *InlineResponse200112TimeEntry) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProjectName

`func (o *InlineResponse200112TimeEntry) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *InlineResponse200112TimeEntry) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *InlineResponse200112TimeEntry) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *InlineResponse200112TimeEntry) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetProjectStatus

`func (o *InlineResponse200112TimeEntry) GetProjectStatus() string`

GetProjectStatus returns the ProjectStatus field if non-nil, zero value otherwise.

### GetProjectStatusOk

`func (o *InlineResponse200112TimeEntry) GetProjectStatusOk() (*string, bool)`

GetProjectStatusOk returns a tuple with the ProjectStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectStatus

`func (o *InlineResponse200112TimeEntry) SetProjectStatus(v string)`

SetProjectStatus sets ProjectStatus field to given value.

### HasProjectStatus

`func (o *InlineResponse200112TimeEntry) HasProjectStatus() bool`

HasProjectStatus returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse200112TimeEntry) GetTags() []InlineResponse2002Column`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse200112TimeEntry) GetTagsOk() (*[]InlineResponse2002Column, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse200112TimeEntry) SetTags(v []InlineResponse2002Column)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse200112TimeEntry) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTaskTags

`func (o *InlineResponse200112TimeEntry) GetTaskTags() []map[string]interface{}`

GetTaskTags returns the TaskTags field if non-nil, zero value otherwise.

### GetTaskTagsOk

`func (o *InlineResponse200112TimeEntry) GetTaskTagsOk() (*[]map[string]interface{}, bool)`

GetTaskTagsOk returns a tuple with the TaskTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskTags

`func (o *InlineResponse200112TimeEntry) SetTaskTags(v []map[string]interface{})`

SetTaskTags sets TaskTags field to given value.

### HasTaskTags

`func (o *InlineResponse200112TimeEntry) HasTaskTags() bool`

HasTaskTags returns a boolean if a field has been set.

### GetTaskEstimatedTime

`func (o *InlineResponse200112TimeEntry) GetTaskEstimatedTime() string`

GetTaskEstimatedTime returns the TaskEstimatedTime field if non-nil, zero value otherwise.

### GetTaskEstimatedTimeOk

`func (o *InlineResponse200112TimeEntry) GetTaskEstimatedTimeOk() (*string, bool)`

GetTaskEstimatedTimeOk returns a tuple with the TaskEstimatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskEstimatedTime

`func (o *InlineResponse200112TimeEntry) SetTaskEstimatedTime(v string)`

SetTaskEstimatedTime sets TaskEstimatedTime field to given value.

### HasTaskEstimatedTime

`func (o *InlineResponse200112TimeEntry) HasTaskEstimatedTime() bool`

HasTaskEstimatedTime returns a boolean if a field has been set.

### GetTaskIsPrivate

`func (o *InlineResponse200112TimeEntry) GetTaskIsPrivate() string`

GetTaskIsPrivate returns the TaskIsPrivate field if non-nil, zero value otherwise.

### GetTaskIsPrivateOk

`func (o *InlineResponse200112TimeEntry) GetTaskIsPrivateOk() (*string, bool)`

GetTaskIsPrivateOk returns a tuple with the TaskIsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskIsPrivate

`func (o *InlineResponse200112TimeEntry) SetTaskIsPrivate(v string)`

SetTaskIsPrivate sets TaskIsPrivate field to given value.

### HasTaskIsPrivate

`func (o *InlineResponse200112TimeEntry) HasTaskIsPrivate() bool`

HasTaskIsPrivate returns a boolean if a field has been set.

### GetTaskIsSubTask

`func (o *InlineResponse200112TimeEntry) GetTaskIsSubTask() string`

GetTaskIsSubTask returns the TaskIsSubTask field if non-nil, zero value otherwise.

### GetTaskIsSubTaskOk

`func (o *InlineResponse200112TimeEntry) GetTaskIsSubTaskOk() (*string, bool)`

GetTaskIsSubTaskOk returns a tuple with the TaskIsSubTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskIsSubTask

`func (o *InlineResponse200112TimeEntry) SetTaskIsSubTask(v string)`

SetTaskIsSubTask sets TaskIsSubTask field to given value.

### HasTaskIsSubTask

`func (o *InlineResponse200112TimeEntry) HasTaskIsSubTask() bool`

HasTaskIsSubTask returns a boolean if a field has been set.

### GetTasklistId

`func (o *InlineResponse200112TimeEntry) GetTasklistId() string`

GetTasklistId returns the TasklistId field if non-nil, zero value otherwise.

### GetTasklistIdOk

`func (o *InlineResponse200112TimeEntry) GetTasklistIdOk() (*string, bool)`

GetTasklistIdOk returns a tuple with the TasklistId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasklistId

`func (o *InlineResponse200112TimeEntry) SetTasklistId(v string)`

SetTasklistId sets TasklistId field to given value.

### HasTasklistId

`func (o *InlineResponse200112TimeEntry) HasTasklistId() bool`

HasTasklistId returns a boolean if a field has been set.

### GetTicketId

`func (o *InlineResponse200112TimeEntry) GetTicketId() string`

GetTicketId returns the TicketId field if non-nil, zero value otherwise.

### GetTicketIdOk

`func (o *InlineResponse200112TimeEntry) GetTicketIdOk() (*string, bool)`

GetTicketIdOk returns a tuple with the TicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketId

`func (o *InlineResponse200112TimeEntry) SetTicketId(v string)`

SetTicketId sets TicketId field to given value.

### HasTicketId

`func (o *InlineResponse200112TimeEntry) HasTicketId() bool`

HasTicketId returns a boolean if a field has been set.

### GetTodoItemId

`func (o *InlineResponse200112TimeEntry) GetTodoItemId() string`

GetTodoItemId returns the TodoItemId field if non-nil, zero value otherwise.

### GetTodoItemIdOk

`func (o *InlineResponse200112TimeEntry) GetTodoItemIdOk() (*string, bool)`

GetTodoItemIdOk returns a tuple with the TodoItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodoItemId

`func (o *InlineResponse200112TimeEntry) SetTodoItemId(v string)`

SetTodoItemId sets TodoItemId field to given value.

### HasTodoItemId

`func (o *InlineResponse200112TimeEntry) HasTodoItemId() bool`

HasTodoItemId returns a boolean if a field has been set.

### GetTodoItemName

`func (o *InlineResponse200112TimeEntry) GetTodoItemName() string`

GetTodoItemName returns the TodoItemName field if non-nil, zero value otherwise.

### GetTodoItemNameOk

`func (o *InlineResponse200112TimeEntry) GetTodoItemNameOk() (*string, bool)`

GetTodoItemNameOk returns a tuple with the TodoItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodoItemName

`func (o *InlineResponse200112TimeEntry) SetTodoItemName(v string)`

SetTodoItemName sets TodoItemName field to given value.

### HasTodoItemName

`func (o *InlineResponse200112TimeEntry) HasTodoItemName() bool`

HasTodoItemName returns a boolean if a field has been set.

### GetTodoListId

`func (o *InlineResponse200112TimeEntry) GetTodoListId() string`

GetTodoListId returns the TodoListId field if non-nil, zero value otherwise.

### GetTodoListIdOk

`func (o *InlineResponse200112TimeEntry) GetTodoListIdOk() (*string, bool)`

GetTodoListIdOk returns a tuple with the TodoListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodoListId

`func (o *InlineResponse200112TimeEntry) SetTodoListId(v string)`

SetTodoListId sets TodoListId field to given value.

### HasTodoListId

`func (o *InlineResponse200112TimeEntry) HasTodoListId() bool`

HasTodoListId returns a boolean if a field has been set.

### GetTodoListName

`func (o *InlineResponse200112TimeEntry) GetTodoListName() string`

GetTodoListName returns the TodoListName field if non-nil, zero value otherwise.

### GetTodoListNameOk

`func (o *InlineResponse200112TimeEntry) GetTodoListNameOk() (*string, bool)`

GetTodoListNameOk returns a tuple with the TodoListName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodoListName

`func (o *InlineResponse200112TimeEntry) SetTodoListName(v string)`

SetTodoListName sets TodoListName field to given value.

### HasTodoListName

`func (o *InlineResponse200112TimeEntry) HasTodoListName() bool`

HasTodoListName returns a boolean if a field has been set.

### GetUpdatedDate

`func (o *InlineResponse200112TimeEntry) GetUpdatedDate() string`

GetUpdatedDate returns the UpdatedDate field if non-nil, zero value otherwise.

### GetUpdatedDateOk

`func (o *InlineResponse200112TimeEntry) GetUpdatedDateOk() (*string, bool)`

GetUpdatedDateOk returns a tuple with the UpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDate

`func (o *InlineResponse200112TimeEntry) SetUpdatedDate(v string)`

SetUpdatedDate sets UpdatedDate field to given value.

### HasUpdatedDate

`func (o *InlineResponse200112TimeEntry) HasUpdatedDate() bool`

HasUpdatedDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


