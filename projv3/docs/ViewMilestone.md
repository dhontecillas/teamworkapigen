# ViewMilestone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanComplete** | Pointer to **bool** |  | [optional] 
**CanEdit** | Pointer to **bool** | permissions | [optional] 
**ChangeFollowerIds** | Pointer to **[]int32** |  | [optional] 
**ChangeFollowers** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**CommentFollowerIds** | Pointer to **[]int32** |  | [optional] 
**CommentFollowers** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**CommentsCount** | Pointer to **int32** |  | [optional] 
**Completed** | Pointer to **bool** |  | [optional] 
**CompletedBy** | Pointer to **int32** |  | [optional] 
**CompletedOn** | Pointer to **string** |  | [optional] 
**CompleterId** | Pointer to **int32** |  | [optional] 
**CreatedBy** | Pointer to **int32** |  | [optional] 
**CreatedOn** | Pointer to **string** |  | [optional] 
**CreatorUserId** | Pointer to **int32** |  | [optional] 
**Deadline** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**LastChangedOn** | Pointer to **string** |  | [optional] 
**LatestUpdates** | Pointer to [**[]ViewAudit**](ViewAudit.md) |  | [optional] 
**Lockdown** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**LockdownId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NumCommentsRead** | Pointer to **int32** |  | [optional] 
**OriginalDueDate** | Pointer to **string** |  | [optional] 
**PercentageComplete** | Pointer to **int32** |  | [optional] 
**PercentageTasksCompleted** | Pointer to **int32** |  | [optional] 
**Private** | Pointer to **bool** |  | [optional] 
**Project** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**Reminder** | Pointer to **bool** |  | [optional] 
**ResponsibleParties** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**ResponsiblePartyIds** | Pointer to **[]int32** | optional fields that are returned only when querying a milestone endpoint | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TagIds** | Pointer to **[]int32** |  | [optional] 
**Tags** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**TasklistIds** | Pointer to **[]int32** |  | [optional] 
**Tasklists** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**UpdatedBy** | Pointer to **int32** |  | [optional] 
**UserFollowingChanges** | Pointer to **bool** |  | [optional] 
**UserFollowingComments** | Pointer to **bool** |  | [optional] 

## Methods

### NewViewMilestone

`func NewViewMilestone() *ViewMilestone`

NewViewMilestone instantiates a new ViewMilestone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewMilestoneWithDefaults

`func NewViewMilestoneWithDefaults() *ViewMilestone`

NewViewMilestoneWithDefaults instantiates a new ViewMilestone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanComplete

`func (o *ViewMilestone) GetCanComplete() bool`

GetCanComplete returns the CanComplete field if non-nil, zero value otherwise.

### GetCanCompleteOk

`func (o *ViewMilestone) GetCanCompleteOk() (*bool, bool)`

GetCanCompleteOk returns a tuple with the CanComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanComplete

`func (o *ViewMilestone) SetCanComplete(v bool)`

SetCanComplete sets CanComplete field to given value.

### HasCanComplete

`func (o *ViewMilestone) HasCanComplete() bool`

HasCanComplete returns a boolean if a field has been set.

### GetCanEdit

`func (o *ViewMilestone) GetCanEdit() bool`

GetCanEdit returns the CanEdit field if non-nil, zero value otherwise.

### GetCanEditOk

`func (o *ViewMilestone) GetCanEditOk() (*bool, bool)`

GetCanEditOk returns a tuple with the CanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEdit

`func (o *ViewMilestone) SetCanEdit(v bool)`

SetCanEdit sets CanEdit field to given value.

### HasCanEdit

`func (o *ViewMilestone) HasCanEdit() bool`

HasCanEdit returns a boolean if a field has been set.

### GetChangeFollowerIds

`func (o *ViewMilestone) GetChangeFollowerIds() []int32`

GetChangeFollowerIds returns the ChangeFollowerIds field if non-nil, zero value otherwise.

### GetChangeFollowerIdsOk

`func (o *ViewMilestone) GetChangeFollowerIdsOk() (*[]int32, bool)`

GetChangeFollowerIdsOk returns a tuple with the ChangeFollowerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeFollowerIds

`func (o *ViewMilestone) SetChangeFollowerIds(v []int32)`

SetChangeFollowerIds sets ChangeFollowerIds field to given value.

### HasChangeFollowerIds

`func (o *ViewMilestone) HasChangeFollowerIds() bool`

HasChangeFollowerIds returns a boolean if a field has been set.

### GetChangeFollowers

`func (o *ViewMilestone) GetChangeFollowers() []ViewRelationship`

GetChangeFollowers returns the ChangeFollowers field if non-nil, zero value otherwise.

### GetChangeFollowersOk

`func (o *ViewMilestone) GetChangeFollowersOk() (*[]ViewRelationship, bool)`

GetChangeFollowersOk returns a tuple with the ChangeFollowers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeFollowers

`func (o *ViewMilestone) SetChangeFollowers(v []ViewRelationship)`

SetChangeFollowers sets ChangeFollowers field to given value.

### HasChangeFollowers

`func (o *ViewMilestone) HasChangeFollowers() bool`

HasChangeFollowers returns a boolean if a field has been set.

### GetCommentFollowerIds

`func (o *ViewMilestone) GetCommentFollowerIds() []int32`

GetCommentFollowerIds returns the CommentFollowerIds field if non-nil, zero value otherwise.

### GetCommentFollowerIdsOk

`func (o *ViewMilestone) GetCommentFollowerIdsOk() (*[]int32, bool)`

GetCommentFollowerIdsOk returns a tuple with the CommentFollowerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentFollowerIds

`func (o *ViewMilestone) SetCommentFollowerIds(v []int32)`

SetCommentFollowerIds sets CommentFollowerIds field to given value.

### HasCommentFollowerIds

`func (o *ViewMilestone) HasCommentFollowerIds() bool`

HasCommentFollowerIds returns a boolean if a field has been set.

### GetCommentFollowers

`func (o *ViewMilestone) GetCommentFollowers() []ViewRelationship`

GetCommentFollowers returns the CommentFollowers field if non-nil, zero value otherwise.

### GetCommentFollowersOk

`func (o *ViewMilestone) GetCommentFollowersOk() (*[]ViewRelationship, bool)`

GetCommentFollowersOk returns a tuple with the CommentFollowers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentFollowers

`func (o *ViewMilestone) SetCommentFollowers(v []ViewRelationship)`

SetCommentFollowers sets CommentFollowers field to given value.

### HasCommentFollowers

`func (o *ViewMilestone) HasCommentFollowers() bool`

HasCommentFollowers returns a boolean if a field has been set.

### GetCommentsCount

`func (o *ViewMilestone) GetCommentsCount() int32`

GetCommentsCount returns the CommentsCount field if non-nil, zero value otherwise.

### GetCommentsCountOk

`func (o *ViewMilestone) GetCommentsCountOk() (*int32, bool)`

GetCommentsCountOk returns a tuple with the CommentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsCount

`func (o *ViewMilestone) SetCommentsCount(v int32)`

SetCommentsCount sets CommentsCount field to given value.

### HasCommentsCount

`func (o *ViewMilestone) HasCommentsCount() bool`

HasCommentsCount returns a boolean if a field has been set.

### GetCompleted

`func (o *ViewMilestone) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *ViewMilestone) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *ViewMilestone) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *ViewMilestone) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetCompletedBy

`func (o *ViewMilestone) GetCompletedBy() int32`

GetCompletedBy returns the CompletedBy field if non-nil, zero value otherwise.

### GetCompletedByOk

`func (o *ViewMilestone) GetCompletedByOk() (*int32, bool)`

GetCompletedByOk returns a tuple with the CompletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedBy

`func (o *ViewMilestone) SetCompletedBy(v int32)`

SetCompletedBy sets CompletedBy field to given value.

### HasCompletedBy

`func (o *ViewMilestone) HasCompletedBy() bool`

HasCompletedBy returns a boolean if a field has been set.

### GetCompletedOn

`func (o *ViewMilestone) GetCompletedOn() string`

GetCompletedOn returns the CompletedOn field if non-nil, zero value otherwise.

### GetCompletedOnOk

`func (o *ViewMilestone) GetCompletedOnOk() (*string, bool)`

GetCompletedOnOk returns a tuple with the CompletedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedOn

`func (o *ViewMilestone) SetCompletedOn(v string)`

SetCompletedOn sets CompletedOn field to given value.

### HasCompletedOn

`func (o *ViewMilestone) HasCompletedOn() bool`

HasCompletedOn returns a boolean if a field has been set.

### GetCompleterId

`func (o *ViewMilestone) GetCompleterId() int32`

GetCompleterId returns the CompleterId field if non-nil, zero value otherwise.

### GetCompleterIdOk

`func (o *ViewMilestone) GetCompleterIdOk() (*int32, bool)`

GetCompleterIdOk returns a tuple with the CompleterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleterId

`func (o *ViewMilestone) SetCompleterId(v int32)`

SetCompleterId sets CompleterId field to given value.

### HasCompleterId

`func (o *ViewMilestone) HasCompleterId() bool`

HasCompleterId returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ViewMilestone) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ViewMilestone) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ViewMilestone) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ViewMilestone) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedOn

`func (o *ViewMilestone) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *ViewMilestone) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *ViewMilestone) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *ViewMilestone) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetCreatorUserId

`func (o *ViewMilestone) GetCreatorUserId() int32`

GetCreatorUserId returns the CreatorUserId field if non-nil, zero value otherwise.

### GetCreatorUserIdOk

`func (o *ViewMilestone) GetCreatorUserIdOk() (*int32, bool)`

GetCreatorUserIdOk returns a tuple with the CreatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorUserId

`func (o *ViewMilestone) SetCreatorUserId(v int32)`

SetCreatorUserId sets CreatorUserId field to given value.

### HasCreatorUserId

`func (o *ViewMilestone) HasCreatorUserId() bool`

HasCreatorUserId returns a boolean if a field has been set.

### GetDeadline

`func (o *ViewMilestone) GetDeadline() string`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *ViewMilestone) GetDeadlineOk() (*string, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *ViewMilestone) SetDeadline(v string)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *ViewMilestone) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### GetDescription

`func (o *ViewMilestone) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ViewMilestone) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ViewMilestone) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ViewMilestone) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ViewMilestone) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewMilestone) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewMilestone) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewMilestone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *ViewMilestone) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ViewMilestone) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ViewMilestone) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *ViewMilestone) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetLastChangedOn

`func (o *ViewMilestone) GetLastChangedOn() string`

GetLastChangedOn returns the LastChangedOn field if non-nil, zero value otherwise.

### GetLastChangedOnOk

`func (o *ViewMilestone) GetLastChangedOnOk() (*string, bool)`

GetLastChangedOnOk returns a tuple with the LastChangedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangedOn

`func (o *ViewMilestone) SetLastChangedOn(v string)`

SetLastChangedOn sets LastChangedOn field to given value.

### HasLastChangedOn

`func (o *ViewMilestone) HasLastChangedOn() bool`

HasLastChangedOn returns a boolean if a field has been set.

### GetLatestUpdates

`func (o *ViewMilestone) GetLatestUpdates() []ViewAudit`

GetLatestUpdates returns the LatestUpdates field if non-nil, zero value otherwise.

### GetLatestUpdatesOk

`func (o *ViewMilestone) GetLatestUpdatesOk() (*[]ViewAudit, bool)`

GetLatestUpdatesOk returns a tuple with the LatestUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestUpdates

`func (o *ViewMilestone) SetLatestUpdates(v []ViewAudit)`

SetLatestUpdates sets LatestUpdates field to given value.

### HasLatestUpdates

`func (o *ViewMilestone) HasLatestUpdates() bool`

HasLatestUpdates returns a boolean if a field has been set.

### GetLockdown

`func (o *ViewMilestone) GetLockdown() ViewRelationship`

GetLockdown returns the Lockdown field if non-nil, zero value otherwise.

### GetLockdownOk

`func (o *ViewMilestone) GetLockdownOk() (*ViewRelationship, bool)`

GetLockdownOk returns a tuple with the Lockdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockdown

`func (o *ViewMilestone) SetLockdown(v ViewRelationship)`

SetLockdown sets Lockdown field to given value.

### HasLockdown

`func (o *ViewMilestone) HasLockdown() bool`

HasLockdown returns a boolean if a field has been set.

### GetLockdownId

`func (o *ViewMilestone) GetLockdownId() int32`

GetLockdownId returns the LockdownId field if non-nil, zero value otherwise.

### GetLockdownIdOk

`func (o *ViewMilestone) GetLockdownIdOk() (*int32, bool)`

GetLockdownIdOk returns a tuple with the LockdownId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockdownId

`func (o *ViewMilestone) SetLockdownId(v int32)`

SetLockdownId sets LockdownId field to given value.

### HasLockdownId

`func (o *ViewMilestone) HasLockdownId() bool`

HasLockdownId returns a boolean if a field has been set.

### GetName

`func (o *ViewMilestone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ViewMilestone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ViewMilestone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ViewMilestone) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumCommentsRead

`func (o *ViewMilestone) GetNumCommentsRead() int32`

GetNumCommentsRead returns the NumCommentsRead field if non-nil, zero value otherwise.

### GetNumCommentsReadOk

`func (o *ViewMilestone) GetNumCommentsReadOk() (*int32, bool)`

GetNumCommentsReadOk returns a tuple with the NumCommentsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCommentsRead

`func (o *ViewMilestone) SetNumCommentsRead(v int32)`

SetNumCommentsRead sets NumCommentsRead field to given value.

### HasNumCommentsRead

`func (o *ViewMilestone) HasNumCommentsRead() bool`

HasNumCommentsRead returns a boolean if a field has been set.

### GetOriginalDueDate

`func (o *ViewMilestone) GetOriginalDueDate() string`

GetOriginalDueDate returns the OriginalDueDate field if non-nil, zero value otherwise.

### GetOriginalDueDateOk

`func (o *ViewMilestone) GetOriginalDueDateOk() (*string, bool)`

GetOriginalDueDateOk returns a tuple with the OriginalDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalDueDate

`func (o *ViewMilestone) SetOriginalDueDate(v string)`

SetOriginalDueDate sets OriginalDueDate field to given value.

### HasOriginalDueDate

`func (o *ViewMilestone) HasOriginalDueDate() bool`

HasOriginalDueDate returns a boolean if a field has been set.

### GetPercentageComplete

`func (o *ViewMilestone) GetPercentageComplete() int32`

GetPercentageComplete returns the PercentageComplete field if non-nil, zero value otherwise.

### GetPercentageCompleteOk

`func (o *ViewMilestone) GetPercentageCompleteOk() (*int32, bool)`

GetPercentageCompleteOk returns a tuple with the PercentageComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageComplete

`func (o *ViewMilestone) SetPercentageComplete(v int32)`

SetPercentageComplete sets PercentageComplete field to given value.

### HasPercentageComplete

`func (o *ViewMilestone) HasPercentageComplete() bool`

HasPercentageComplete returns a boolean if a field has been set.

### GetPercentageTasksCompleted

`func (o *ViewMilestone) GetPercentageTasksCompleted() int32`

GetPercentageTasksCompleted returns the PercentageTasksCompleted field if non-nil, zero value otherwise.

### GetPercentageTasksCompletedOk

`func (o *ViewMilestone) GetPercentageTasksCompletedOk() (*int32, bool)`

GetPercentageTasksCompletedOk returns a tuple with the PercentageTasksCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageTasksCompleted

`func (o *ViewMilestone) SetPercentageTasksCompleted(v int32)`

SetPercentageTasksCompleted sets PercentageTasksCompleted field to given value.

### HasPercentageTasksCompleted

`func (o *ViewMilestone) HasPercentageTasksCompleted() bool`

HasPercentageTasksCompleted returns a boolean if a field has been set.

### GetPrivate

`func (o *ViewMilestone) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *ViewMilestone) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *ViewMilestone) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *ViewMilestone) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetProject

`func (o *ViewMilestone) GetProject() ViewRelationship`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ViewMilestone) GetProjectOk() (*ViewRelationship, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ViewMilestone) SetProject(v ViewRelationship)`

SetProject sets Project field to given value.

### HasProject

`func (o *ViewMilestone) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectId

`func (o *ViewMilestone) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ViewMilestone) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ViewMilestone) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ViewMilestone) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetReminder

`func (o *ViewMilestone) GetReminder() bool`

GetReminder returns the Reminder field if non-nil, zero value otherwise.

### GetReminderOk

`func (o *ViewMilestone) GetReminderOk() (*bool, bool)`

GetReminderOk returns a tuple with the Reminder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminder

`func (o *ViewMilestone) SetReminder(v bool)`

SetReminder sets Reminder field to given value.

### HasReminder

`func (o *ViewMilestone) HasReminder() bool`

HasReminder returns a boolean if a field has been set.

### GetResponsibleParties

`func (o *ViewMilestone) GetResponsibleParties() []ViewRelationship`

GetResponsibleParties returns the ResponsibleParties field if non-nil, zero value otherwise.

### GetResponsiblePartiesOk

`func (o *ViewMilestone) GetResponsiblePartiesOk() (*[]ViewRelationship, bool)`

GetResponsiblePartiesOk returns a tuple with the ResponsibleParties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsibleParties

`func (o *ViewMilestone) SetResponsibleParties(v []ViewRelationship)`

SetResponsibleParties sets ResponsibleParties field to given value.

### HasResponsibleParties

`func (o *ViewMilestone) HasResponsibleParties() bool`

HasResponsibleParties returns a boolean if a field has been set.

### GetResponsiblePartyIds

`func (o *ViewMilestone) GetResponsiblePartyIds() []int32`

GetResponsiblePartyIds returns the ResponsiblePartyIds field if non-nil, zero value otherwise.

### GetResponsiblePartyIdsOk

`func (o *ViewMilestone) GetResponsiblePartyIdsOk() (*[]int32, bool)`

GetResponsiblePartyIdsOk returns a tuple with the ResponsiblePartyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsiblePartyIds

`func (o *ViewMilestone) SetResponsiblePartyIds(v []int32)`

SetResponsiblePartyIds sets ResponsiblePartyIds field to given value.

### HasResponsiblePartyIds

`func (o *ViewMilestone) HasResponsiblePartyIds() bool`

HasResponsiblePartyIds returns a boolean if a field has been set.

### GetStatus

`func (o *ViewMilestone) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ViewMilestone) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ViewMilestone) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ViewMilestone) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTagIds

`func (o *ViewMilestone) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *ViewMilestone) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *ViewMilestone) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *ViewMilestone) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTags

`func (o *ViewMilestone) GetTags() []ViewRelationship`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ViewMilestone) GetTagsOk() (*[]ViewRelationship, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ViewMilestone) SetTags(v []ViewRelationship)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ViewMilestone) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTasklistIds

`func (o *ViewMilestone) GetTasklistIds() []int32`

GetTasklistIds returns the TasklistIds field if non-nil, zero value otherwise.

### GetTasklistIdsOk

`func (o *ViewMilestone) GetTasklistIdsOk() (*[]int32, bool)`

GetTasklistIdsOk returns a tuple with the TasklistIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasklistIds

`func (o *ViewMilestone) SetTasklistIds(v []int32)`

SetTasklistIds sets TasklistIds field to given value.

### HasTasklistIds

`func (o *ViewMilestone) HasTasklistIds() bool`

HasTasklistIds returns a boolean if a field has been set.

### GetTasklists

`func (o *ViewMilestone) GetTasklists() []ViewRelationship`

GetTasklists returns the Tasklists field if non-nil, zero value otherwise.

### GetTasklistsOk

`func (o *ViewMilestone) GetTasklistsOk() (*[]ViewRelationship, bool)`

GetTasklistsOk returns a tuple with the Tasklists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasklists

`func (o *ViewMilestone) SetTasklists(v []ViewRelationship)`

SetTasklists sets Tasklists field to given value.

### HasTasklists

`func (o *ViewMilestone) HasTasklists() bool`

HasTasklists returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ViewMilestone) GetUpdatedBy() int32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ViewMilestone) GetUpdatedByOk() (*int32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ViewMilestone) SetUpdatedBy(v int32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ViewMilestone) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUserFollowingChanges

`func (o *ViewMilestone) GetUserFollowingChanges() bool`

GetUserFollowingChanges returns the UserFollowingChanges field if non-nil, zero value otherwise.

### GetUserFollowingChangesOk

`func (o *ViewMilestone) GetUserFollowingChangesOk() (*bool, bool)`

GetUserFollowingChangesOk returns a tuple with the UserFollowingChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFollowingChanges

`func (o *ViewMilestone) SetUserFollowingChanges(v bool)`

SetUserFollowingChanges sets UserFollowingChanges field to given value.

### HasUserFollowingChanges

`func (o *ViewMilestone) HasUserFollowingChanges() bool`

HasUserFollowingChanges returns a boolean if a field has been set.

### GetUserFollowingComments

`func (o *ViewMilestone) GetUserFollowingComments() bool`

GetUserFollowingComments returns the UserFollowingComments field if non-nil, zero value otherwise.

### GetUserFollowingCommentsOk

`func (o *ViewMilestone) GetUserFollowingCommentsOk() (*bool, bool)`

GetUserFollowingCommentsOk returns a tuple with the UserFollowingComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFollowingComments

`func (o *ViewMilestone) SetUserFollowingComments(v bool)`

SetUserFollowingComments sets UserFollowingComments field to given value.

### HasUserFollowingComments

`func (o *ViewMilestone) HasUserFollowingComments() bool`

HasUserFollowingComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


