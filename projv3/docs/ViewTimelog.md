# ViewTimelog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Billable** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **string** |  | [optional] 
**DateDeleted** | Pointer to **string** |  | [optional] 
**DateEdited** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**DeletedBy** | Pointer to **int32** |  | [optional] 
**DeletedByUserId** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DeskTicketId** | Pointer to **int32** |  | [optional] 
**EditedByUserId** | Pointer to **int32** |  | [optional] 
**HasStartTime** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**InvoiceNumber** | Pointer to **string** |  | [optional] 
**LoggedBy** | Pointer to **int32** |  | [optional] 
**LoggedByUserId** | Pointer to **int32** |  | [optional] 
**Minutes** | Pointer to **int32** |  | [optional] 
**Project** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**ProjectBillingInvoice** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**ProjectBillingInvoiceId** | Pointer to **int32** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**TagIds** | Pointer to **[]int32** |  | [optional] 
**Tags** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**Task** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**TaskId** | Pointer to **int32** |  | [optional] 
**TaskIdPreMove** | Pointer to **int32** |  | [optional] 
**TaskPreMove** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**TimeLogged** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **int32** |  | [optional] 
**User** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 

## Methods

### NewViewTimelog

`func NewViewTimelog() *ViewTimelog`

NewViewTimelog instantiates a new ViewTimelog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewTimelogWithDefaults

`func NewViewTimelogWithDefaults() *ViewTimelog`

NewViewTimelogWithDefaults instantiates a new ViewTimelog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillable

`func (o *ViewTimelog) GetBillable() bool`

GetBillable returns the Billable field if non-nil, zero value otherwise.

### GetBillableOk

`func (o *ViewTimelog) GetBillableOk() (*bool, bool)`

GetBillableOk returns a tuple with the Billable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillable

`func (o *ViewTimelog) SetBillable(v bool)`

SetBillable sets Billable field to given value.

### HasBillable

`func (o *ViewTimelog) HasBillable() bool`

HasBillable returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ViewTimelog) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ViewTimelog) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ViewTimelog) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ViewTimelog) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDateCreated

`func (o *ViewTimelog) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ViewTimelog) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ViewTimelog) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ViewTimelog) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateDeleted

`func (o *ViewTimelog) GetDateDeleted() string`

GetDateDeleted returns the DateDeleted field if non-nil, zero value otherwise.

### GetDateDeletedOk

`func (o *ViewTimelog) GetDateDeletedOk() (*string, bool)`

GetDateDeletedOk returns a tuple with the DateDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeleted

`func (o *ViewTimelog) SetDateDeleted(v string)`

SetDateDeleted sets DateDeleted field to given value.

### HasDateDeleted

`func (o *ViewTimelog) HasDateDeleted() bool`

HasDateDeleted returns a boolean if a field has been set.

### GetDateEdited

`func (o *ViewTimelog) GetDateEdited() string`

GetDateEdited returns the DateEdited field if non-nil, zero value otherwise.

### GetDateEditedOk

`func (o *ViewTimelog) GetDateEditedOk() (*string, bool)`

GetDateEditedOk returns a tuple with the DateEdited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEdited

`func (o *ViewTimelog) SetDateEdited(v string)`

SetDateEdited sets DateEdited field to given value.

### HasDateEdited

`func (o *ViewTimelog) HasDateEdited() bool`

HasDateEdited returns a boolean if a field has been set.

### GetDeleted

`func (o *ViewTimelog) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *ViewTimelog) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *ViewTimelog) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *ViewTimelog) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ViewTimelog) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ViewTimelog) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ViewTimelog) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ViewTimelog) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeletedBy

`func (o *ViewTimelog) GetDeletedBy() int32`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *ViewTimelog) GetDeletedByOk() (*int32, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *ViewTimelog) SetDeletedBy(v int32)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *ViewTimelog) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetDeletedByUserId

`func (o *ViewTimelog) GetDeletedByUserId() int32`

GetDeletedByUserId returns the DeletedByUserId field if non-nil, zero value otherwise.

### GetDeletedByUserIdOk

`func (o *ViewTimelog) GetDeletedByUserIdOk() (*int32, bool)`

GetDeletedByUserIdOk returns a tuple with the DeletedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedByUserId

`func (o *ViewTimelog) SetDeletedByUserId(v int32)`

SetDeletedByUserId sets DeletedByUserId field to given value.

### HasDeletedByUserId

`func (o *ViewTimelog) HasDeletedByUserId() bool`

HasDeletedByUserId returns a boolean if a field has been set.

### GetDescription

`func (o *ViewTimelog) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ViewTimelog) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ViewTimelog) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ViewTimelog) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeskTicketId

`func (o *ViewTimelog) GetDeskTicketId() int32`

GetDeskTicketId returns the DeskTicketId field if non-nil, zero value otherwise.

### GetDeskTicketIdOk

`func (o *ViewTimelog) GetDeskTicketIdOk() (*int32, bool)`

GetDeskTicketIdOk returns a tuple with the DeskTicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeskTicketId

`func (o *ViewTimelog) SetDeskTicketId(v int32)`

SetDeskTicketId sets DeskTicketId field to given value.

### HasDeskTicketId

`func (o *ViewTimelog) HasDeskTicketId() bool`

HasDeskTicketId returns a boolean if a field has been set.

### GetEditedByUserId

`func (o *ViewTimelog) GetEditedByUserId() int32`

GetEditedByUserId returns the EditedByUserId field if non-nil, zero value otherwise.

### GetEditedByUserIdOk

`func (o *ViewTimelog) GetEditedByUserIdOk() (*int32, bool)`

GetEditedByUserIdOk returns a tuple with the EditedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedByUserId

`func (o *ViewTimelog) SetEditedByUserId(v int32)`

SetEditedByUserId sets EditedByUserId field to given value.

### HasEditedByUserId

`func (o *ViewTimelog) HasEditedByUserId() bool`

HasEditedByUserId returns a boolean if a field has been set.

### GetHasStartTime

`func (o *ViewTimelog) GetHasStartTime() bool`

GetHasStartTime returns the HasStartTime field if non-nil, zero value otherwise.

### GetHasStartTimeOk

`func (o *ViewTimelog) GetHasStartTimeOk() (*bool, bool)`

GetHasStartTimeOk returns a tuple with the HasStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasStartTime

`func (o *ViewTimelog) SetHasStartTime(v bool)`

SetHasStartTime sets HasStartTime field to given value.

### HasHasStartTime

`func (o *ViewTimelog) HasHasStartTime() bool`

HasHasStartTime returns a boolean if a field has been set.

### GetId

`func (o *ViewTimelog) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewTimelog) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewTimelog) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewTimelog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceNumber

`func (o *ViewTimelog) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *ViewTimelog) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *ViewTimelog) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *ViewTimelog) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetLoggedBy

`func (o *ViewTimelog) GetLoggedBy() int32`

GetLoggedBy returns the LoggedBy field if non-nil, zero value otherwise.

### GetLoggedByOk

`func (o *ViewTimelog) GetLoggedByOk() (*int32, bool)`

GetLoggedByOk returns a tuple with the LoggedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggedBy

`func (o *ViewTimelog) SetLoggedBy(v int32)`

SetLoggedBy sets LoggedBy field to given value.

### HasLoggedBy

`func (o *ViewTimelog) HasLoggedBy() bool`

HasLoggedBy returns a boolean if a field has been set.

### GetLoggedByUserId

`func (o *ViewTimelog) GetLoggedByUserId() int32`

GetLoggedByUserId returns the LoggedByUserId field if non-nil, zero value otherwise.

### GetLoggedByUserIdOk

`func (o *ViewTimelog) GetLoggedByUserIdOk() (*int32, bool)`

GetLoggedByUserIdOk returns a tuple with the LoggedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggedByUserId

`func (o *ViewTimelog) SetLoggedByUserId(v int32)`

SetLoggedByUserId sets LoggedByUserId field to given value.

### HasLoggedByUserId

`func (o *ViewTimelog) HasLoggedByUserId() bool`

HasLoggedByUserId returns a boolean if a field has been set.

### GetMinutes

`func (o *ViewTimelog) GetMinutes() int32`

GetMinutes returns the Minutes field if non-nil, zero value otherwise.

### GetMinutesOk

`func (o *ViewTimelog) GetMinutesOk() (*int32, bool)`

GetMinutesOk returns a tuple with the Minutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutes

`func (o *ViewTimelog) SetMinutes(v int32)`

SetMinutes sets Minutes field to given value.

### HasMinutes

`func (o *ViewTimelog) HasMinutes() bool`

HasMinutes returns a boolean if a field has been set.

### GetProject

`func (o *ViewTimelog) GetProject() ViewRelationship`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ViewTimelog) GetProjectOk() (*ViewRelationship, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ViewTimelog) SetProject(v ViewRelationship)`

SetProject sets Project field to given value.

### HasProject

`func (o *ViewTimelog) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectBillingInvoice

`func (o *ViewTimelog) GetProjectBillingInvoice() ViewRelationship`

GetProjectBillingInvoice returns the ProjectBillingInvoice field if non-nil, zero value otherwise.

### GetProjectBillingInvoiceOk

`func (o *ViewTimelog) GetProjectBillingInvoiceOk() (*ViewRelationship, bool)`

GetProjectBillingInvoiceOk returns a tuple with the ProjectBillingInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectBillingInvoice

`func (o *ViewTimelog) SetProjectBillingInvoice(v ViewRelationship)`

SetProjectBillingInvoice sets ProjectBillingInvoice field to given value.

### HasProjectBillingInvoice

`func (o *ViewTimelog) HasProjectBillingInvoice() bool`

HasProjectBillingInvoice returns a boolean if a field has been set.

### GetProjectBillingInvoiceId

`func (o *ViewTimelog) GetProjectBillingInvoiceId() int32`

GetProjectBillingInvoiceId returns the ProjectBillingInvoiceId field if non-nil, zero value otherwise.

### GetProjectBillingInvoiceIdOk

`func (o *ViewTimelog) GetProjectBillingInvoiceIdOk() (*int32, bool)`

GetProjectBillingInvoiceIdOk returns a tuple with the ProjectBillingInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectBillingInvoiceId

`func (o *ViewTimelog) SetProjectBillingInvoiceId(v int32)`

SetProjectBillingInvoiceId sets ProjectBillingInvoiceId field to given value.

### HasProjectBillingInvoiceId

`func (o *ViewTimelog) HasProjectBillingInvoiceId() bool`

HasProjectBillingInvoiceId returns a boolean if a field has been set.

### GetProjectId

`func (o *ViewTimelog) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ViewTimelog) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ViewTimelog) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ViewTimelog) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetTagIds

`func (o *ViewTimelog) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *ViewTimelog) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *ViewTimelog) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *ViewTimelog) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTags

`func (o *ViewTimelog) GetTags() []ViewRelationship`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ViewTimelog) GetTagsOk() (*[]ViewRelationship, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ViewTimelog) SetTags(v []ViewRelationship)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ViewTimelog) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTask

`func (o *ViewTimelog) GetTask() ViewRelationship`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *ViewTimelog) GetTaskOk() (*ViewRelationship, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *ViewTimelog) SetTask(v ViewRelationship)`

SetTask sets Task field to given value.

### HasTask

`func (o *ViewTimelog) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetTaskId

`func (o *ViewTimelog) GetTaskId() int32`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *ViewTimelog) GetTaskIdOk() (*int32, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *ViewTimelog) SetTaskId(v int32)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *ViewTimelog) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetTaskIdPreMove

`func (o *ViewTimelog) GetTaskIdPreMove() int32`

GetTaskIdPreMove returns the TaskIdPreMove field if non-nil, zero value otherwise.

### GetTaskIdPreMoveOk

`func (o *ViewTimelog) GetTaskIdPreMoveOk() (*int32, bool)`

GetTaskIdPreMoveOk returns a tuple with the TaskIdPreMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskIdPreMove

`func (o *ViewTimelog) SetTaskIdPreMove(v int32)`

SetTaskIdPreMove sets TaskIdPreMove field to given value.

### HasTaskIdPreMove

`func (o *ViewTimelog) HasTaskIdPreMove() bool`

HasTaskIdPreMove returns a boolean if a field has been set.

### GetTaskPreMove

`func (o *ViewTimelog) GetTaskPreMove() ViewRelationship`

GetTaskPreMove returns the TaskPreMove field if non-nil, zero value otherwise.

### GetTaskPreMoveOk

`func (o *ViewTimelog) GetTaskPreMoveOk() (*ViewRelationship, bool)`

GetTaskPreMoveOk returns a tuple with the TaskPreMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskPreMove

`func (o *ViewTimelog) SetTaskPreMove(v ViewRelationship)`

SetTaskPreMove sets TaskPreMove field to given value.

### HasTaskPreMove

`func (o *ViewTimelog) HasTaskPreMove() bool`

HasTaskPreMove returns a boolean if a field has been set.

### GetTimeLogged

`func (o *ViewTimelog) GetTimeLogged() string`

GetTimeLogged returns the TimeLogged field if non-nil, zero value otherwise.

### GetTimeLoggedOk

`func (o *ViewTimelog) GetTimeLoggedOk() (*string, bool)`

GetTimeLoggedOk returns a tuple with the TimeLogged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLogged

`func (o *ViewTimelog) SetTimeLogged(v string)`

SetTimeLogged sets TimeLogged field to given value.

### HasTimeLogged

`func (o *ViewTimelog) HasTimeLogged() bool`

HasTimeLogged returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ViewTimelog) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ViewTimelog) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ViewTimelog) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ViewTimelog) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ViewTimelog) GetUpdatedBy() int32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ViewTimelog) GetUpdatedByOk() (*int32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ViewTimelog) SetUpdatedBy(v int32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ViewTimelog) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUser

`func (o *ViewTimelog) GetUser() ViewRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ViewTimelog) GetUserOk() (*ViewRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ViewTimelog) SetUser(v ViewRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *ViewTimelog) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserId

`func (o *ViewTimelog) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ViewTimelog) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ViewTimelog) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ViewTimelog) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


