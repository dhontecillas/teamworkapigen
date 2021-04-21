# ViewTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssigneeCompanies** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**AssigneeCompanyIds** | Pointer to **[]int32** |  | [optional] 
**AssigneeTeamIds** | Pointer to **[]int32** |  | [optional] 
**AssigneeTeams** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**AssigneeUserIds** | Pointer to **[]int32** |  | [optional] 
**AssigneeUsers** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**CompletedBy** | Pointer to **int32** |  | [optional] 
**CompletedOn** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **int32** |  | [optional] 
**CreatedByUserId** | Pointer to **int32** |  | [optional] 
**CrmDealIds** | Pointer to **[]int32** |  | [optional] 
**DateUpdated** | Pointer to **string** |  | [optional] 
**DependencyIds** | Pointer to **[]int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayOrder** | Pointer to **int32** |  | [optional] 
**DueDate** | Pointer to **string** |  | [optional] 
**EstimateMinutes** | Pointer to **int32** |  | [optional] 
**HasDeskTickets** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**IsPrivate** | Pointer to **int32** |  | [optional] 
**LatestUpdates** | Pointer to [**[]ViewAudit**](ViewAudit.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OriginalDueDate** | Pointer to **map[string]interface{}** | Date represents a Unified API Spec date format. | [optional] 
**ParentTask** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**ParentTaskId** | Pointer to **int32** |  | [optional] 
**PredecessorIds** | Pointer to **[]int32** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**Progress** | Pointer to **int32** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SubTaskIds** | Pointer to **[]int32** |  | [optional] 
**TagIds** | Pointer to **[]int32** |  | [optional] 
**Tasklist** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**TasklistId** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **int32** |  | [optional] 
**UserPermissions** | Pointer to [**ViewTaskPermissions**](ViewTaskPermissions.md) |  | [optional] 

## Methods

### NewViewTask

`func NewViewTask() *ViewTask`

NewViewTask instantiates a new ViewTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewTaskWithDefaults

`func NewViewTaskWithDefaults() *ViewTask`

NewViewTaskWithDefaults instantiates a new ViewTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssigneeCompanies

`func (o *ViewTask) GetAssigneeCompanies() []ViewRelationship`

GetAssigneeCompanies returns the AssigneeCompanies field if non-nil, zero value otherwise.

### GetAssigneeCompaniesOk

`func (o *ViewTask) GetAssigneeCompaniesOk() (*[]ViewRelationship, bool)`

GetAssigneeCompaniesOk returns a tuple with the AssigneeCompanies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeCompanies

`func (o *ViewTask) SetAssigneeCompanies(v []ViewRelationship)`

SetAssigneeCompanies sets AssigneeCompanies field to given value.

### HasAssigneeCompanies

`func (o *ViewTask) HasAssigneeCompanies() bool`

HasAssigneeCompanies returns a boolean if a field has been set.

### GetAssigneeCompanyIds

`func (o *ViewTask) GetAssigneeCompanyIds() []int32`

GetAssigneeCompanyIds returns the AssigneeCompanyIds field if non-nil, zero value otherwise.

### GetAssigneeCompanyIdsOk

`func (o *ViewTask) GetAssigneeCompanyIdsOk() (*[]int32, bool)`

GetAssigneeCompanyIdsOk returns a tuple with the AssigneeCompanyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeCompanyIds

`func (o *ViewTask) SetAssigneeCompanyIds(v []int32)`

SetAssigneeCompanyIds sets AssigneeCompanyIds field to given value.

### HasAssigneeCompanyIds

`func (o *ViewTask) HasAssigneeCompanyIds() bool`

HasAssigneeCompanyIds returns a boolean if a field has been set.

### GetAssigneeTeamIds

`func (o *ViewTask) GetAssigneeTeamIds() []int32`

GetAssigneeTeamIds returns the AssigneeTeamIds field if non-nil, zero value otherwise.

### GetAssigneeTeamIdsOk

`func (o *ViewTask) GetAssigneeTeamIdsOk() (*[]int32, bool)`

GetAssigneeTeamIdsOk returns a tuple with the AssigneeTeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeTeamIds

`func (o *ViewTask) SetAssigneeTeamIds(v []int32)`

SetAssigneeTeamIds sets AssigneeTeamIds field to given value.

### HasAssigneeTeamIds

`func (o *ViewTask) HasAssigneeTeamIds() bool`

HasAssigneeTeamIds returns a boolean if a field has been set.

### GetAssigneeTeams

`func (o *ViewTask) GetAssigneeTeams() []ViewRelationship`

GetAssigneeTeams returns the AssigneeTeams field if non-nil, zero value otherwise.

### GetAssigneeTeamsOk

`func (o *ViewTask) GetAssigneeTeamsOk() (*[]ViewRelationship, bool)`

GetAssigneeTeamsOk returns a tuple with the AssigneeTeams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeTeams

`func (o *ViewTask) SetAssigneeTeams(v []ViewRelationship)`

SetAssigneeTeams sets AssigneeTeams field to given value.

### HasAssigneeTeams

`func (o *ViewTask) HasAssigneeTeams() bool`

HasAssigneeTeams returns a boolean if a field has been set.

### GetAssigneeUserIds

`func (o *ViewTask) GetAssigneeUserIds() []int32`

GetAssigneeUserIds returns the AssigneeUserIds field if non-nil, zero value otherwise.

### GetAssigneeUserIdsOk

`func (o *ViewTask) GetAssigneeUserIdsOk() (*[]int32, bool)`

GetAssigneeUserIdsOk returns a tuple with the AssigneeUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeUserIds

`func (o *ViewTask) SetAssigneeUserIds(v []int32)`

SetAssigneeUserIds sets AssigneeUserIds field to given value.

### HasAssigneeUserIds

`func (o *ViewTask) HasAssigneeUserIds() bool`

HasAssigneeUserIds returns a boolean if a field has been set.

### GetAssigneeUsers

`func (o *ViewTask) GetAssigneeUsers() []ViewRelationship`

GetAssigneeUsers returns the AssigneeUsers field if non-nil, zero value otherwise.

### GetAssigneeUsersOk

`func (o *ViewTask) GetAssigneeUsersOk() (*[]ViewRelationship, bool)`

GetAssigneeUsersOk returns a tuple with the AssigneeUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeUsers

`func (o *ViewTask) SetAssigneeUsers(v []ViewRelationship)`

SetAssigneeUsers sets AssigneeUsers field to given value.

### HasAssigneeUsers

`func (o *ViewTask) HasAssigneeUsers() bool`

HasAssigneeUsers returns a boolean if a field has been set.

### GetCompletedBy

`func (o *ViewTask) GetCompletedBy() int32`

GetCompletedBy returns the CompletedBy field if non-nil, zero value otherwise.

### GetCompletedByOk

`func (o *ViewTask) GetCompletedByOk() (*int32, bool)`

GetCompletedByOk returns a tuple with the CompletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedBy

`func (o *ViewTask) SetCompletedBy(v int32)`

SetCompletedBy sets CompletedBy field to given value.

### HasCompletedBy

`func (o *ViewTask) HasCompletedBy() bool`

HasCompletedBy returns a boolean if a field has been set.

### GetCompletedOn

`func (o *ViewTask) GetCompletedOn() string`

GetCompletedOn returns the CompletedOn field if non-nil, zero value otherwise.

### GetCompletedOnOk

`func (o *ViewTask) GetCompletedOnOk() (*string, bool)`

GetCompletedOnOk returns a tuple with the CompletedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedOn

`func (o *ViewTask) SetCompletedOn(v string)`

SetCompletedOn sets CompletedOn field to given value.

### HasCompletedOn

`func (o *ViewTask) HasCompletedOn() bool`

HasCompletedOn returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ViewTask) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ViewTask) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ViewTask) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ViewTask) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *ViewTask) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *ViewTask) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *ViewTask) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *ViewTask) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetCrmDealIds

`func (o *ViewTask) GetCrmDealIds() []int32`

GetCrmDealIds returns the CrmDealIds field if non-nil, zero value otherwise.

### GetCrmDealIdsOk

`func (o *ViewTask) GetCrmDealIdsOk() (*[]int32, bool)`

GetCrmDealIdsOk returns a tuple with the CrmDealIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrmDealIds

`func (o *ViewTask) SetCrmDealIds(v []int32)`

SetCrmDealIds sets CrmDealIds field to given value.

### HasCrmDealIds

`func (o *ViewTask) HasCrmDealIds() bool`

HasCrmDealIds returns a boolean if a field has been set.

### GetDateUpdated

`func (o *ViewTask) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ViewTask) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ViewTask) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ViewTask) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### GetDependencyIds

`func (o *ViewTask) GetDependencyIds() []int32`

GetDependencyIds returns the DependencyIds field if non-nil, zero value otherwise.

### GetDependencyIdsOk

`func (o *ViewTask) GetDependencyIdsOk() (*[]int32, bool)`

GetDependencyIdsOk returns a tuple with the DependencyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyIds

`func (o *ViewTask) SetDependencyIds(v []int32)`

SetDependencyIds sets DependencyIds field to given value.

### HasDependencyIds

`func (o *ViewTask) HasDependencyIds() bool`

HasDependencyIds returns a boolean if a field has been set.

### GetDescription

`func (o *ViewTask) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ViewTask) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ViewTask) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ViewTask) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *ViewTask) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *ViewTask) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *ViewTask) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *ViewTask) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetDueDate

`func (o *ViewTask) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *ViewTask) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *ViewTask) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *ViewTask) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetEstimateMinutes

`func (o *ViewTask) GetEstimateMinutes() int32`

GetEstimateMinutes returns the EstimateMinutes field if non-nil, zero value otherwise.

### GetEstimateMinutesOk

`func (o *ViewTask) GetEstimateMinutesOk() (*int32, bool)`

GetEstimateMinutesOk returns a tuple with the EstimateMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimateMinutes

`func (o *ViewTask) SetEstimateMinutes(v int32)`

SetEstimateMinutes sets EstimateMinutes field to given value.

### HasEstimateMinutes

`func (o *ViewTask) HasEstimateMinutes() bool`

HasEstimateMinutes returns a boolean if a field has been set.

### GetHasDeskTickets

`func (o *ViewTask) GetHasDeskTickets() bool`

GetHasDeskTickets returns the HasDeskTickets field if non-nil, zero value otherwise.

### GetHasDeskTicketsOk

`func (o *ViewTask) GetHasDeskTicketsOk() (*bool, bool)`

GetHasDeskTicketsOk returns a tuple with the HasDeskTickets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDeskTickets

`func (o *ViewTask) SetHasDeskTickets(v bool)`

SetHasDeskTickets sets HasDeskTickets field to given value.

### HasHasDeskTickets

`func (o *ViewTask) HasHasDeskTickets() bool`

HasHasDeskTickets returns a boolean if a field has been set.

### GetId

`func (o *ViewTask) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewTask) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewTask) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsPrivate

`func (o *ViewTask) GetIsPrivate() int32`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *ViewTask) GetIsPrivateOk() (*int32, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *ViewTask) SetIsPrivate(v int32)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *ViewTask) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetLatestUpdates

`func (o *ViewTask) GetLatestUpdates() []ViewAudit`

GetLatestUpdates returns the LatestUpdates field if non-nil, zero value otherwise.

### GetLatestUpdatesOk

`func (o *ViewTask) GetLatestUpdatesOk() (*[]ViewAudit, bool)`

GetLatestUpdatesOk returns a tuple with the LatestUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestUpdates

`func (o *ViewTask) SetLatestUpdates(v []ViewAudit)`

SetLatestUpdates sets LatestUpdates field to given value.

### HasLatestUpdates

`func (o *ViewTask) HasLatestUpdates() bool`

HasLatestUpdates returns a boolean if a field has been set.

### GetName

`func (o *ViewTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ViewTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ViewTask) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ViewTask) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOriginalDueDate

`func (o *ViewTask) GetOriginalDueDate() map[string]interface{}`

GetOriginalDueDate returns the OriginalDueDate field if non-nil, zero value otherwise.

### GetOriginalDueDateOk

`func (o *ViewTask) GetOriginalDueDateOk() (*map[string]interface{}, bool)`

GetOriginalDueDateOk returns a tuple with the OriginalDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalDueDate

`func (o *ViewTask) SetOriginalDueDate(v map[string]interface{})`

SetOriginalDueDate sets OriginalDueDate field to given value.

### HasOriginalDueDate

`func (o *ViewTask) HasOriginalDueDate() bool`

HasOriginalDueDate returns a boolean if a field has been set.

### GetParentTask

`func (o *ViewTask) GetParentTask() ViewRelationship`

GetParentTask returns the ParentTask field if non-nil, zero value otherwise.

### GetParentTaskOk

`func (o *ViewTask) GetParentTaskOk() (*ViewRelationship, bool)`

GetParentTaskOk returns a tuple with the ParentTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTask

`func (o *ViewTask) SetParentTask(v ViewRelationship)`

SetParentTask sets ParentTask field to given value.

### HasParentTask

`func (o *ViewTask) HasParentTask() bool`

HasParentTask returns a boolean if a field has been set.

### GetParentTaskId

`func (o *ViewTask) GetParentTaskId() int32`

GetParentTaskId returns the ParentTaskId field if non-nil, zero value otherwise.

### GetParentTaskIdOk

`func (o *ViewTask) GetParentTaskIdOk() (*int32, bool)`

GetParentTaskIdOk returns a tuple with the ParentTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTaskId

`func (o *ViewTask) SetParentTaskId(v int32)`

SetParentTaskId sets ParentTaskId field to given value.

### HasParentTaskId

`func (o *ViewTask) HasParentTaskId() bool`

HasParentTaskId returns a boolean if a field has been set.

### GetPredecessorIds

`func (o *ViewTask) GetPredecessorIds() []int32`

GetPredecessorIds returns the PredecessorIds field if non-nil, zero value otherwise.

### GetPredecessorIdsOk

`func (o *ViewTask) GetPredecessorIdsOk() (*[]int32, bool)`

GetPredecessorIdsOk returns a tuple with the PredecessorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredecessorIds

`func (o *ViewTask) SetPredecessorIds(v []int32)`

SetPredecessorIds sets PredecessorIds field to given value.

### HasPredecessorIds

`func (o *ViewTask) HasPredecessorIds() bool`

HasPredecessorIds returns a boolean if a field has been set.

### GetPriority

`func (o *ViewTask) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ViewTask) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ViewTask) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ViewTask) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProgress

`func (o *ViewTask) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ViewTask) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ViewTask) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *ViewTask) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetStartDate

`func (o *ViewTask) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ViewTask) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ViewTask) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ViewTask) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *ViewTask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ViewTask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ViewTask) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ViewTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubTaskIds

`func (o *ViewTask) GetSubTaskIds() []int32`

GetSubTaskIds returns the SubTaskIds field if non-nil, zero value otherwise.

### GetSubTaskIdsOk

`func (o *ViewTask) GetSubTaskIdsOk() (*[]int32, bool)`

GetSubTaskIdsOk returns a tuple with the SubTaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTaskIds

`func (o *ViewTask) SetSubTaskIds(v []int32)`

SetSubTaskIds sets SubTaskIds field to given value.

### HasSubTaskIds

`func (o *ViewTask) HasSubTaskIds() bool`

HasSubTaskIds returns a boolean if a field has been set.

### GetTagIds

`func (o *ViewTask) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *ViewTask) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *ViewTask) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *ViewTask) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTasklist

`func (o *ViewTask) GetTasklist() ViewRelationship`

GetTasklist returns the Tasklist field if non-nil, zero value otherwise.

### GetTasklistOk

`func (o *ViewTask) GetTasklistOk() (*ViewRelationship, bool)`

GetTasklistOk returns a tuple with the Tasklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasklist

`func (o *ViewTask) SetTasklist(v ViewRelationship)`

SetTasklist sets Tasklist field to given value.

### HasTasklist

`func (o *ViewTask) HasTasklist() bool`

HasTasklist returns a boolean if a field has been set.

### GetTasklistId

`func (o *ViewTask) GetTasklistId() int32`

GetTasklistId returns the TasklistId field if non-nil, zero value otherwise.

### GetTasklistIdOk

`func (o *ViewTask) GetTasklistIdOk() (*int32, bool)`

GetTasklistIdOk returns a tuple with the TasklistId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasklistId

`func (o *ViewTask) SetTasklistId(v int32)`

SetTasklistId sets TasklistId field to given value.

### HasTasklistId

`func (o *ViewTask) HasTasklistId() bool`

HasTasklistId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ViewTask) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ViewTask) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ViewTask) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ViewTask) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ViewTask) GetUpdatedBy() int32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ViewTask) GetUpdatedByOk() (*int32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ViewTask) SetUpdatedBy(v int32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ViewTask) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUserPermissions

`func (o *ViewTask) GetUserPermissions() ViewTaskPermissions`

GetUserPermissions returns the UserPermissions field if non-nil, zero value otherwise.

### GetUserPermissionsOk

`func (o *ViewTask) GetUserPermissionsOk() (*ViewTaskPermissions, bool)`

GetUserPermissionsOk returns a tuple with the UserPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPermissions

`func (o *ViewTask) SetUserPermissions(v ViewTaskPermissions)`

SetUserPermissions sets UserPermissions field to given value.

### HasUserPermissions

`func (o *ViewTask) HasUserPermissions() bool`

HasUserPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


