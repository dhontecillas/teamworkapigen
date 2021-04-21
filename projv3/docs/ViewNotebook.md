# ViewNotebook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**CategoryId** | Pointer to **int32** |  | [optional] 
**ChangeFollowers** | Pointer to [**ViewUserGroups**](ViewUserGroups.md) |  | [optional] 
**CommentFollowers** | Pointer to [**ViewUserGroups**](ViewUserGroups.md) |  | [optional] 
**CommentsCount** | Pointer to **int32** |  | [optional] 
**ContentHTML** | Pointer to **string** |  | [optional] 
**Contents** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **int32** |  | [optional] 
**CreatedByUserID** | Pointer to **int32** |  | [optional] 
**DateCreated** | Pointer to **string** |  | [optional] 
**DateDeleted** | Pointer to **string** |  | [optional] 
**DateUpdated** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**DeletedBy** | Pointer to **int32** |  | [optional] 
**DeletedByUserID** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**IsPrivate** | Pointer to **bool** |  | [optional] 
**LatestVersionNo** | Pointer to **int32** |  | [optional] 
**Lockdown** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**LockdownId** | Pointer to **int32** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**LockedAt** | Pointer to **string** |  | [optional] 
**LockedBy** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NotebookVersion** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**NotebookVersionCreatedDateTime** | Pointer to **string** |  | [optional] 
**NotebookVersionID** | Pointer to **int32** |  | [optional] 
**NotebookVersionUpdatedDateTime** | Pointer to **string** |  | [optional] 
**Privacy** | Pointer to [**ViewUserGroups**](ViewUserGroups.md) |  | [optional] 
**Project** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**ReadCommentsCount** | Pointer to **int32** |  | [optional] 
**SecureContent** | Pointer to **bool** |  | [optional] 
**TagIds** | Pointer to **[]int32** |  | [optional] 
**Tags** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **int32** |  | [optional] 
**UpdatedByUserID** | Pointer to **int32** |  | [optional] 
**UserFollowingChanges** | Pointer to **bool** |  | [optional] 
**UserFollowingComments** | Pointer to **bool** |  | [optional] 
**VersionCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewViewNotebook

`func NewViewNotebook() *ViewNotebook`

NewViewNotebook instantiates a new ViewNotebook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewNotebookWithDefaults

`func NewViewNotebookWithDefaults() *ViewNotebook`

NewViewNotebookWithDefaults instantiates a new ViewNotebook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *ViewNotebook) GetCategory() ViewRelationship`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ViewNotebook) GetCategoryOk() (*ViewRelationship, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ViewNotebook) SetCategory(v ViewRelationship)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ViewNotebook) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCategoryId

`func (o *ViewNotebook) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *ViewNotebook) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *ViewNotebook) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *ViewNotebook) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetChangeFollowers

`func (o *ViewNotebook) GetChangeFollowers() ViewUserGroups`

GetChangeFollowers returns the ChangeFollowers field if non-nil, zero value otherwise.

### GetChangeFollowersOk

`func (o *ViewNotebook) GetChangeFollowersOk() (*ViewUserGroups, bool)`

GetChangeFollowersOk returns a tuple with the ChangeFollowers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeFollowers

`func (o *ViewNotebook) SetChangeFollowers(v ViewUserGroups)`

SetChangeFollowers sets ChangeFollowers field to given value.

### HasChangeFollowers

`func (o *ViewNotebook) HasChangeFollowers() bool`

HasChangeFollowers returns a boolean if a field has been set.

### GetCommentFollowers

`func (o *ViewNotebook) GetCommentFollowers() ViewUserGroups`

GetCommentFollowers returns the CommentFollowers field if non-nil, zero value otherwise.

### GetCommentFollowersOk

`func (o *ViewNotebook) GetCommentFollowersOk() (*ViewUserGroups, bool)`

GetCommentFollowersOk returns a tuple with the CommentFollowers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentFollowers

`func (o *ViewNotebook) SetCommentFollowers(v ViewUserGroups)`

SetCommentFollowers sets CommentFollowers field to given value.

### HasCommentFollowers

`func (o *ViewNotebook) HasCommentFollowers() bool`

HasCommentFollowers returns a boolean if a field has been set.

### GetCommentsCount

`func (o *ViewNotebook) GetCommentsCount() int32`

GetCommentsCount returns the CommentsCount field if non-nil, zero value otherwise.

### GetCommentsCountOk

`func (o *ViewNotebook) GetCommentsCountOk() (*int32, bool)`

GetCommentsCountOk returns a tuple with the CommentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsCount

`func (o *ViewNotebook) SetCommentsCount(v int32)`

SetCommentsCount sets CommentsCount field to given value.

### HasCommentsCount

`func (o *ViewNotebook) HasCommentsCount() bool`

HasCommentsCount returns a boolean if a field has been set.

### GetContentHTML

`func (o *ViewNotebook) GetContentHTML() string`

GetContentHTML returns the ContentHTML field if non-nil, zero value otherwise.

### GetContentHTMLOk

`func (o *ViewNotebook) GetContentHTMLOk() (*string, bool)`

GetContentHTMLOk returns a tuple with the ContentHTML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentHTML

`func (o *ViewNotebook) SetContentHTML(v string)`

SetContentHTML sets ContentHTML field to given value.

### HasContentHTML

`func (o *ViewNotebook) HasContentHTML() bool`

HasContentHTML returns a boolean if a field has been set.

### GetContents

`func (o *ViewNotebook) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *ViewNotebook) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *ViewNotebook) SetContents(v string)`

SetContents sets Contents field to given value.

### HasContents

`func (o *ViewNotebook) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ViewNotebook) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ViewNotebook) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ViewNotebook) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ViewNotebook) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ViewNotebook) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ViewNotebook) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ViewNotebook) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ViewNotebook) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUserID

`func (o *ViewNotebook) GetCreatedByUserID() int32`

GetCreatedByUserID returns the CreatedByUserID field if non-nil, zero value otherwise.

### GetCreatedByUserIDOk

`func (o *ViewNotebook) GetCreatedByUserIDOk() (*int32, bool)`

GetCreatedByUserIDOk returns a tuple with the CreatedByUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserID

`func (o *ViewNotebook) SetCreatedByUserID(v int32)`

SetCreatedByUserID sets CreatedByUserID field to given value.

### HasCreatedByUserID

`func (o *ViewNotebook) HasCreatedByUserID() bool`

HasCreatedByUserID returns a boolean if a field has been set.

### GetDateCreated

`func (o *ViewNotebook) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ViewNotebook) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ViewNotebook) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ViewNotebook) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateDeleted

`func (o *ViewNotebook) GetDateDeleted() string`

GetDateDeleted returns the DateDeleted field if non-nil, zero value otherwise.

### GetDateDeletedOk

`func (o *ViewNotebook) GetDateDeletedOk() (*string, bool)`

GetDateDeletedOk returns a tuple with the DateDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeleted

`func (o *ViewNotebook) SetDateDeleted(v string)`

SetDateDeleted sets DateDeleted field to given value.

### HasDateDeleted

`func (o *ViewNotebook) HasDateDeleted() bool`

HasDateDeleted returns a boolean if a field has been set.

### GetDateUpdated

`func (o *ViewNotebook) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ViewNotebook) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ViewNotebook) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ViewNotebook) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### GetDeleted

`func (o *ViewNotebook) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *ViewNotebook) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *ViewNotebook) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *ViewNotebook) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ViewNotebook) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ViewNotebook) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ViewNotebook) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ViewNotebook) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeletedBy

`func (o *ViewNotebook) GetDeletedBy() int32`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *ViewNotebook) GetDeletedByOk() (*int32, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *ViewNotebook) SetDeletedBy(v int32)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *ViewNotebook) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetDeletedByUserID

`func (o *ViewNotebook) GetDeletedByUserID() int32`

GetDeletedByUserID returns the DeletedByUserID field if non-nil, zero value otherwise.

### GetDeletedByUserIDOk

`func (o *ViewNotebook) GetDeletedByUserIDOk() (*int32, bool)`

GetDeletedByUserIDOk returns a tuple with the DeletedByUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedByUserID

`func (o *ViewNotebook) SetDeletedByUserID(v int32)`

SetDeletedByUserID sets DeletedByUserID field to given value.

### HasDeletedByUserID

`func (o *ViewNotebook) HasDeletedByUserID() bool`

HasDeletedByUserID returns a boolean if a field has been set.

### GetDescription

`func (o *ViewNotebook) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ViewNotebook) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ViewNotebook) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ViewNotebook) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ViewNotebook) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewNotebook) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewNotebook) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewNotebook) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsPrivate

`func (o *ViewNotebook) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *ViewNotebook) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *ViewNotebook) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *ViewNotebook) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetLatestVersionNo

`func (o *ViewNotebook) GetLatestVersionNo() int32`

GetLatestVersionNo returns the LatestVersionNo field if non-nil, zero value otherwise.

### GetLatestVersionNoOk

`func (o *ViewNotebook) GetLatestVersionNoOk() (*int32, bool)`

GetLatestVersionNoOk returns a tuple with the LatestVersionNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersionNo

`func (o *ViewNotebook) SetLatestVersionNo(v int32)`

SetLatestVersionNo sets LatestVersionNo field to given value.

### HasLatestVersionNo

`func (o *ViewNotebook) HasLatestVersionNo() bool`

HasLatestVersionNo returns a boolean if a field has been set.

### GetLockdown

`func (o *ViewNotebook) GetLockdown() ViewRelationship`

GetLockdown returns the Lockdown field if non-nil, zero value otherwise.

### GetLockdownOk

`func (o *ViewNotebook) GetLockdownOk() (*ViewRelationship, bool)`

GetLockdownOk returns a tuple with the Lockdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockdown

`func (o *ViewNotebook) SetLockdown(v ViewRelationship)`

SetLockdown sets Lockdown field to given value.

### HasLockdown

`func (o *ViewNotebook) HasLockdown() bool`

HasLockdown returns a boolean if a field has been set.

### GetLockdownId

`func (o *ViewNotebook) GetLockdownId() int32`

GetLockdownId returns the LockdownId field if non-nil, zero value otherwise.

### GetLockdownIdOk

`func (o *ViewNotebook) GetLockdownIdOk() (*int32, bool)`

GetLockdownIdOk returns a tuple with the LockdownId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockdownId

`func (o *ViewNotebook) SetLockdownId(v int32)`

SetLockdownId sets LockdownId field to given value.

### HasLockdownId

`func (o *ViewNotebook) HasLockdownId() bool`

HasLockdownId returns a boolean if a field has been set.

### GetLocked

`func (o *ViewNotebook) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *ViewNotebook) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *ViewNotebook) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *ViewNotebook) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetLockedAt

`func (o *ViewNotebook) GetLockedAt() string`

GetLockedAt returns the LockedAt field if non-nil, zero value otherwise.

### GetLockedAtOk

`func (o *ViewNotebook) GetLockedAtOk() (*string, bool)`

GetLockedAtOk returns a tuple with the LockedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedAt

`func (o *ViewNotebook) SetLockedAt(v string)`

SetLockedAt sets LockedAt field to given value.

### HasLockedAt

`func (o *ViewNotebook) HasLockedAt() bool`

HasLockedAt returns a boolean if a field has been set.

### GetLockedBy

`func (o *ViewNotebook) GetLockedBy() int32`

GetLockedBy returns the LockedBy field if non-nil, zero value otherwise.

### GetLockedByOk

`func (o *ViewNotebook) GetLockedByOk() (*int32, bool)`

GetLockedByOk returns a tuple with the LockedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedBy

`func (o *ViewNotebook) SetLockedBy(v int32)`

SetLockedBy sets LockedBy field to given value.

### HasLockedBy

`func (o *ViewNotebook) HasLockedBy() bool`

HasLockedBy returns a boolean if a field has been set.

### GetName

`func (o *ViewNotebook) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ViewNotebook) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ViewNotebook) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ViewNotebook) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotebookVersion

`func (o *ViewNotebook) GetNotebookVersion() ViewRelationship`

GetNotebookVersion returns the NotebookVersion field if non-nil, zero value otherwise.

### GetNotebookVersionOk

`func (o *ViewNotebook) GetNotebookVersionOk() (*ViewRelationship, bool)`

GetNotebookVersionOk returns a tuple with the NotebookVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotebookVersion

`func (o *ViewNotebook) SetNotebookVersion(v ViewRelationship)`

SetNotebookVersion sets NotebookVersion field to given value.

### HasNotebookVersion

`func (o *ViewNotebook) HasNotebookVersion() bool`

HasNotebookVersion returns a boolean if a field has been set.

### GetNotebookVersionCreatedDateTime

`func (o *ViewNotebook) GetNotebookVersionCreatedDateTime() string`

GetNotebookVersionCreatedDateTime returns the NotebookVersionCreatedDateTime field if non-nil, zero value otherwise.

### GetNotebookVersionCreatedDateTimeOk

`func (o *ViewNotebook) GetNotebookVersionCreatedDateTimeOk() (*string, bool)`

GetNotebookVersionCreatedDateTimeOk returns a tuple with the NotebookVersionCreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotebookVersionCreatedDateTime

`func (o *ViewNotebook) SetNotebookVersionCreatedDateTime(v string)`

SetNotebookVersionCreatedDateTime sets NotebookVersionCreatedDateTime field to given value.

### HasNotebookVersionCreatedDateTime

`func (o *ViewNotebook) HasNotebookVersionCreatedDateTime() bool`

HasNotebookVersionCreatedDateTime returns a boolean if a field has been set.

### GetNotebookVersionID

`func (o *ViewNotebook) GetNotebookVersionID() int32`

GetNotebookVersionID returns the NotebookVersionID field if non-nil, zero value otherwise.

### GetNotebookVersionIDOk

`func (o *ViewNotebook) GetNotebookVersionIDOk() (*int32, bool)`

GetNotebookVersionIDOk returns a tuple with the NotebookVersionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotebookVersionID

`func (o *ViewNotebook) SetNotebookVersionID(v int32)`

SetNotebookVersionID sets NotebookVersionID field to given value.

### HasNotebookVersionID

`func (o *ViewNotebook) HasNotebookVersionID() bool`

HasNotebookVersionID returns a boolean if a field has been set.

### GetNotebookVersionUpdatedDateTime

`func (o *ViewNotebook) GetNotebookVersionUpdatedDateTime() string`

GetNotebookVersionUpdatedDateTime returns the NotebookVersionUpdatedDateTime field if non-nil, zero value otherwise.

### GetNotebookVersionUpdatedDateTimeOk

`func (o *ViewNotebook) GetNotebookVersionUpdatedDateTimeOk() (*string, bool)`

GetNotebookVersionUpdatedDateTimeOk returns a tuple with the NotebookVersionUpdatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotebookVersionUpdatedDateTime

`func (o *ViewNotebook) SetNotebookVersionUpdatedDateTime(v string)`

SetNotebookVersionUpdatedDateTime sets NotebookVersionUpdatedDateTime field to given value.

### HasNotebookVersionUpdatedDateTime

`func (o *ViewNotebook) HasNotebookVersionUpdatedDateTime() bool`

HasNotebookVersionUpdatedDateTime returns a boolean if a field has been set.

### GetPrivacy

`func (o *ViewNotebook) GetPrivacy() ViewUserGroups`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *ViewNotebook) GetPrivacyOk() (*ViewUserGroups, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *ViewNotebook) SetPrivacy(v ViewUserGroups)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *ViewNotebook) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetProject

`func (o *ViewNotebook) GetProject() ViewRelationship`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ViewNotebook) GetProjectOk() (*ViewRelationship, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ViewNotebook) SetProject(v ViewRelationship)`

SetProject sets Project field to given value.

### HasProject

`func (o *ViewNotebook) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectId

`func (o *ViewNotebook) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ViewNotebook) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ViewNotebook) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ViewNotebook) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetReadCommentsCount

`func (o *ViewNotebook) GetReadCommentsCount() int32`

GetReadCommentsCount returns the ReadCommentsCount field if non-nil, zero value otherwise.

### GetReadCommentsCountOk

`func (o *ViewNotebook) GetReadCommentsCountOk() (*int32, bool)`

GetReadCommentsCountOk returns a tuple with the ReadCommentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadCommentsCount

`func (o *ViewNotebook) SetReadCommentsCount(v int32)`

SetReadCommentsCount sets ReadCommentsCount field to given value.

### HasReadCommentsCount

`func (o *ViewNotebook) HasReadCommentsCount() bool`

HasReadCommentsCount returns a boolean if a field has been set.

### GetSecureContent

`func (o *ViewNotebook) GetSecureContent() bool`

GetSecureContent returns the SecureContent field if non-nil, zero value otherwise.

### GetSecureContentOk

`func (o *ViewNotebook) GetSecureContentOk() (*bool, bool)`

GetSecureContentOk returns a tuple with the SecureContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureContent

`func (o *ViewNotebook) SetSecureContent(v bool)`

SetSecureContent sets SecureContent field to given value.

### HasSecureContent

`func (o *ViewNotebook) HasSecureContent() bool`

HasSecureContent returns a boolean if a field has been set.

### GetTagIds

`func (o *ViewNotebook) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *ViewNotebook) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *ViewNotebook) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *ViewNotebook) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTags

`func (o *ViewNotebook) GetTags() []ViewRelationship`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ViewNotebook) GetTagsOk() (*[]ViewRelationship, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ViewNotebook) SetTags(v []ViewRelationship)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ViewNotebook) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *ViewNotebook) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ViewNotebook) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ViewNotebook) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ViewNotebook) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ViewNotebook) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ViewNotebook) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ViewNotebook) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ViewNotebook) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ViewNotebook) GetUpdatedBy() int32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ViewNotebook) GetUpdatedByOk() (*int32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ViewNotebook) SetUpdatedBy(v int32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ViewNotebook) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedByUserID

`func (o *ViewNotebook) GetUpdatedByUserID() int32`

GetUpdatedByUserID returns the UpdatedByUserID field if non-nil, zero value otherwise.

### GetUpdatedByUserIDOk

`func (o *ViewNotebook) GetUpdatedByUserIDOk() (*int32, bool)`

GetUpdatedByUserIDOk returns a tuple with the UpdatedByUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUserID

`func (o *ViewNotebook) SetUpdatedByUserID(v int32)`

SetUpdatedByUserID sets UpdatedByUserID field to given value.

### HasUpdatedByUserID

`func (o *ViewNotebook) HasUpdatedByUserID() bool`

HasUpdatedByUserID returns a boolean if a field has been set.

### GetUserFollowingChanges

`func (o *ViewNotebook) GetUserFollowingChanges() bool`

GetUserFollowingChanges returns the UserFollowingChanges field if non-nil, zero value otherwise.

### GetUserFollowingChangesOk

`func (o *ViewNotebook) GetUserFollowingChangesOk() (*bool, bool)`

GetUserFollowingChangesOk returns a tuple with the UserFollowingChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFollowingChanges

`func (o *ViewNotebook) SetUserFollowingChanges(v bool)`

SetUserFollowingChanges sets UserFollowingChanges field to given value.

### HasUserFollowingChanges

`func (o *ViewNotebook) HasUserFollowingChanges() bool`

HasUserFollowingChanges returns a boolean if a field has been set.

### GetUserFollowingComments

`func (o *ViewNotebook) GetUserFollowingComments() bool`

GetUserFollowingComments returns the UserFollowingComments field if non-nil, zero value otherwise.

### GetUserFollowingCommentsOk

`func (o *ViewNotebook) GetUserFollowingCommentsOk() (*bool, bool)`

GetUserFollowingCommentsOk returns a tuple with the UserFollowingComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFollowingComments

`func (o *ViewNotebook) SetUserFollowingComments(v bool)`

SetUserFollowingComments sets UserFollowingComments field to given value.

### HasUserFollowingComments

`func (o *ViewNotebook) HasUserFollowingComments() bool`

HasUserFollowingComments returns a boolean if a field has been set.

### GetVersionCount

`func (o *ViewNotebook) GetVersionCount() int32`

GetVersionCount returns the VersionCount field if non-nil, zero value otherwise.

### GetVersionCountOk

`func (o *ViewNotebook) GetVersionCountOk() (*int32, bool)`

GetVersionCountOk returns a tuple with the VersionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionCount

`func (o *ViewNotebook) SetVersionCount(v int32)`

SetVersionCount sets VersionCount field to given value.

### HasVersionCount

`func (o *ViewNotebook) HasVersionCount() bool`

HasVersionCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


