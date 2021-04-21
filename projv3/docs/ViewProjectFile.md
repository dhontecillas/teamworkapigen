# ViewProjectFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**CategoryId** | Pointer to **int32** |  | [optional] 
**ChangeFollowers** | Pointer to [**ViewUserGroups**](ViewUserGroups.md) |  | [optional] 
**CommentFollowers** | Pointer to [**ViewUserGroups**](ViewUserGroups.md) |  | [optional] 
**CommentsCount** | Pointer to **int32** |  | [optional] 
**CommentsCountRead** | Pointer to **int32** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**DeletedBy** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**DownloadURL** | Pointer to **string** |  | [optional] 
**FileSource** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**IsPrivate** | Pointer to **int32** | IsPrivate can have the values 0 for public, 1 for private and 2 for inherited. | [optional] 
**LatestFileVersionNo** | Pointer to **int32** |  | [optional] 
**Lockdown** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**LockdownId** | Pointer to **int32** |  | [optional] 
**LockedBy** | Pointer to **int32** |  | [optional] 
**LockedByUserId** | Pointer to **int32** |  | [optional] 
**LockedDate** | Pointer to **string** |  | [optional] 
**NumLikes** | Pointer to **int32** |  | [optional] 
**OriginalName** | Pointer to **string** |  | [optional] 
**PreviewURL** | Pointer to **string** |  | [optional] 
**Project** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**Reactions** | Pointer to [**ViewReactionsForObject**](ViewReactionsForObject.md) |  | [optional] 
**RelatedItems** | Pointer to [**ViewRelatedItems**](ViewRelatedItems.md) |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TagIds** | Pointer to **[]int32** |  | [optional] 
**Tags** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**ThumbURL** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UploadedBy** | Pointer to **int32** |  | [optional] 
**UploadedByUserID** | Pointer to **int32** |  | [optional] 
**UploadedDate** | Pointer to **string** |  | [optional] 
**Version** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**VersionId** | Pointer to **int32** |  | [optional] 
**Versions** | Pointer to [**[]ViewFileversion**](ViewFileversion.md) |  | [optional] 

## Methods

### NewViewProjectFile

`func NewViewProjectFile() *ViewProjectFile`

NewViewProjectFile instantiates a new ViewProjectFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewProjectFileWithDefaults

`func NewViewProjectFileWithDefaults() *ViewProjectFile`

NewViewProjectFileWithDefaults instantiates a new ViewProjectFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *ViewProjectFile) GetCategory() ViewRelationship`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ViewProjectFile) GetCategoryOk() (*ViewRelationship, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ViewProjectFile) SetCategory(v ViewRelationship)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ViewProjectFile) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCategoryId

`func (o *ViewProjectFile) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *ViewProjectFile) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *ViewProjectFile) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *ViewProjectFile) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetChangeFollowers

`func (o *ViewProjectFile) GetChangeFollowers() ViewUserGroups`

GetChangeFollowers returns the ChangeFollowers field if non-nil, zero value otherwise.

### GetChangeFollowersOk

`func (o *ViewProjectFile) GetChangeFollowersOk() (*ViewUserGroups, bool)`

GetChangeFollowersOk returns a tuple with the ChangeFollowers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeFollowers

`func (o *ViewProjectFile) SetChangeFollowers(v ViewUserGroups)`

SetChangeFollowers sets ChangeFollowers field to given value.

### HasChangeFollowers

`func (o *ViewProjectFile) HasChangeFollowers() bool`

HasChangeFollowers returns a boolean if a field has been set.

### GetCommentFollowers

`func (o *ViewProjectFile) GetCommentFollowers() ViewUserGroups`

GetCommentFollowers returns the CommentFollowers field if non-nil, zero value otherwise.

### GetCommentFollowersOk

`func (o *ViewProjectFile) GetCommentFollowersOk() (*ViewUserGroups, bool)`

GetCommentFollowersOk returns a tuple with the CommentFollowers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentFollowers

`func (o *ViewProjectFile) SetCommentFollowers(v ViewUserGroups)`

SetCommentFollowers sets CommentFollowers field to given value.

### HasCommentFollowers

`func (o *ViewProjectFile) HasCommentFollowers() bool`

HasCommentFollowers returns a boolean if a field has been set.

### GetCommentsCount

`func (o *ViewProjectFile) GetCommentsCount() int32`

GetCommentsCount returns the CommentsCount field if non-nil, zero value otherwise.

### GetCommentsCountOk

`func (o *ViewProjectFile) GetCommentsCountOk() (*int32, bool)`

GetCommentsCountOk returns a tuple with the CommentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsCount

`func (o *ViewProjectFile) SetCommentsCount(v int32)`

SetCommentsCount sets CommentsCount field to given value.

### HasCommentsCount

`func (o *ViewProjectFile) HasCommentsCount() bool`

HasCommentsCount returns a boolean if a field has been set.

### GetCommentsCountRead

`func (o *ViewProjectFile) GetCommentsCountRead() int32`

GetCommentsCountRead returns the CommentsCountRead field if non-nil, zero value otherwise.

### GetCommentsCountReadOk

`func (o *ViewProjectFile) GetCommentsCountReadOk() (*int32, bool)`

GetCommentsCountReadOk returns a tuple with the CommentsCountRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsCountRead

`func (o *ViewProjectFile) SetCommentsCountRead(v int32)`

SetCommentsCountRead sets CommentsCountRead field to given value.

### HasCommentsCountRead

`func (o *ViewProjectFile) HasCommentsCountRead() bool`

HasCommentsCountRead returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ViewProjectFile) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ViewProjectFile) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ViewProjectFile) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ViewProjectFile) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeletedBy

`func (o *ViewProjectFile) GetDeletedBy() int32`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *ViewProjectFile) GetDeletedByOk() (*int32, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *ViewProjectFile) SetDeletedBy(v int32)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *ViewProjectFile) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetDescription

`func (o *ViewProjectFile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ViewProjectFile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ViewProjectFile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ViewProjectFile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *ViewProjectFile) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ViewProjectFile) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ViewProjectFile) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ViewProjectFile) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDownloadURL

`func (o *ViewProjectFile) GetDownloadURL() string`

GetDownloadURL returns the DownloadURL field if non-nil, zero value otherwise.

### GetDownloadURLOk

`func (o *ViewProjectFile) GetDownloadURLOk() (*string, bool)`

GetDownloadURLOk returns a tuple with the DownloadURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadURL

`func (o *ViewProjectFile) SetDownloadURL(v string)`

SetDownloadURL sets DownloadURL field to given value.

### HasDownloadURL

`func (o *ViewProjectFile) HasDownloadURL() bool`

HasDownloadURL returns a boolean if a field has been set.

### GetFileSource

`func (o *ViewProjectFile) GetFileSource() string`

GetFileSource returns the FileSource field if non-nil, zero value otherwise.

### GetFileSourceOk

`func (o *ViewProjectFile) GetFileSourceOk() (*string, bool)`

GetFileSourceOk returns a tuple with the FileSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSource

`func (o *ViewProjectFile) SetFileSource(v string)`

SetFileSource sets FileSource field to given value.

### HasFileSource

`func (o *ViewProjectFile) HasFileSource() bool`

HasFileSource returns a boolean if a field has been set.

### GetId

`func (o *ViewProjectFile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewProjectFile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewProjectFile) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewProjectFile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsLocked

`func (o *ViewProjectFile) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *ViewProjectFile) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *ViewProjectFile) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *ViewProjectFile) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetIsPrivate

`func (o *ViewProjectFile) GetIsPrivate() int32`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *ViewProjectFile) GetIsPrivateOk() (*int32, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *ViewProjectFile) SetIsPrivate(v int32)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *ViewProjectFile) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetLatestFileVersionNo

`func (o *ViewProjectFile) GetLatestFileVersionNo() int32`

GetLatestFileVersionNo returns the LatestFileVersionNo field if non-nil, zero value otherwise.

### GetLatestFileVersionNoOk

`func (o *ViewProjectFile) GetLatestFileVersionNoOk() (*int32, bool)`

GetLatestFileVersionNoOk returns a tuple with the LatestFileVersionNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestFileVersionNo

`func (o *ViewProjectFile) SetLatestFileVersionNo(v int32)`

SetLatestFileVersionNo sets LatestFileVersionNo field to given value.

### HasLatestFileVersionNo

`func (o *ViewProjectFile) HasLatestFileVersionNo() bool`

HasLatestFileVersionNo returns a boolean if a field has been set.

### GetLockdown

`func (o *ViewProjectFile) GetLockdown() ViewRelationship`

GetLockdown returns the Lockdown field if non-nil, zero value otherwise.

### GetLockdownOk

`func (o *ViewProjectFile) GetLockdownOk() (*ViewRelationship, bool)`

GetLockdownOk returns a tuple with the Lockdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockdown

`func (o *ViewProjectFile) SetLockdown(v ViewRelationship)`

SetLockdown sets Lockdown field to given value.

### HasLockdown

`func (o *ViewProjectFile) HasLockdown() bool`

HasLockdown returns a boolean if a field has been set.

### GetLockdownId

`func (o *ViewProjectFile) GetLockdownId() int32`

GetLockdownId returns the LockdownId field if non-nil, zero value otherwise.

### GetLockdownIdOk

`func (o *ViewProjectFile) GetLockdownIdOk() (*int32, bool)`

GetLockdownIdOk returns a tuple with the LockdownId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockdownId

`func (o *ViewProjectFile) SetLockdownId(v int32)`

SetLockdownId sets LockdownId field to given value.

### HasLockdownId

`func (o *ViewProjectFile) HasLockdownId() bool`

HasLockdownId returns a boolean if a field has been set.

### GetLockedBy

`func (o *ViewProjectFile) GetLockedBy() int32`

GetLockedBy returns the LockedBy field if non-nil, zero value otherwise.

### GetLockedByOk

`func (o *ViewProjectFile) GetLockedByOk() (*int32, bool)`

GetLockedByOk returns a tuple with the LockedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedBy

`func (o *ViewProjectFile) SetLockedBy(v int32)`

SetLockedBy sets LockedBy field to given value.

### HasLockedBy

`func (o *ViewProjectFile) HasLockedBy() bool`

HasLockedBy returns a boolean if a field has been set.

### GetLockedByUserId

`func (o *ViewProjectFile) GetLockedByUserId() int32`

GetLockedByUserId returns the LockedByUserId field if non-nil, zero value otherwise.

### GetLockedByUserIdOk

`func (o *ViewProjectFile) GetLockedByUserIdOk() (*int32, bool)`

GetLockedByUserIdOk returns a tuple with the LockedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedByUserId

`func (o *ViewProjectFile) SetLockedByUserId(v int32)`

SetLockedByUserId sets LockedByUserId field to given value.

### HasLockedByUserId

`func (o *ViewProjectFile) HasLockedByUserId() bool`

HasLockedByUserId returns a boolean if a field has been set.

### GetLockedDate

`func (o *ViewProjectFile) GetLockedDate() string`

GetLockedDate returns the LockedDate field if non-nil, zero value otherwise.

### GetLockedDateOk

`func (o *ViewProjectFile) GetLockedDateOk() (*string, bool)`

GetLockedDateOk returns a tuple with the LockedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedDate

`func (o *ViewProjectFile) SetLockedDate(v string)`

SetLockedDate sets LockedDate field to given value.

### HasLockedDate

`func (o *ViewProjectFile) HasLockedDate() bool`

HasLockedDate returns a boolean if a field has been set.

### GetNumLikes

`func (o *ViewProjectFile) GetNumLikes() int32`

GetNumLikes returns the NumLikes field if non-nil, zero value otherwise.

### GetNumLikesOk

`func (o *ViewProjectFile) GetNumLikesOk() (*int32, bool)`

GetNumLikesOk returns a tuple with the NumLikes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLikes

`func (o *ViewProjectFile) SetNumLikes(v int32)`

SetNumLikes sets NumLikes field to given value.

### HasNumLikes

`func (o *ViewProjectFile) HasNumLikes() bool`

HasNumLikes returns a boolean if a field has been set.

### GetOriginalName

`func (o *ViewProjectFile) GetOriginalName() string`

GetOriginalName returns the OriginalName field if non-nil, zero value otherwise.

### GetOriginalNameOk

`func (o *ViewProjectFile) GetOriginalNameOk() (*string, bool)`

GetOriginalNameOk returns a tuple with the OriginalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalName

`func (o *ViewProjectFile) SetOriginalName(v string)`

SetOriginalName sets OriginalName field to given value.

### HasOriginalName

`func (o *ViewProjectFile) HasOriginalName() bool`

HasOriginalName returns a boolean if a field has been set.

### GetPreviewURL

`func (o *ViewProjectFile) GetPreviewURL() string`

GetPreviewURL returns the PreviewURL field if non-nil, zero value otherwise.

### GetPreviewURLOk

`func (o *ViewProjectFile) GetPreviewURLOk() (*string, bool)`

GetPreviewURLOk returns a tuple with the PreviewURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewURL

`func (o *ViewProjectFile) SetPreviewURL(v string)`

SetPreviewURL sets PreviewURL field to given value.

### HasPreviewURL

`func (o *ViewProjectFile) HasPreviewURL() bool`

HasPreviewURL returns a boolean if a field has been set.

### GetProject

`func (o *ViewProjectFile) GetProject() ViewRelationship`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ViewProjectFile) GetProjectOk() (*ViewRelationship, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ViewProjectFile) SetProject(v ViewRelationship)`

SetProject sets Project field to given value.

### HasProject

`func (o *ViewProjectFile) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectId

`func (o *ViewProjectFile) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ViewProjectFile) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ViewProjectFile) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ViewProjectFile) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetReactions

`func (o *ViewProjectFile) GetReactions() ViewReactionsForObject`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *ViewProjectFile) GetReactionsOk() (*ViewReactionsForObject, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *ViewProjectFile) SetReactions(v ViewReactionsForObject)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *ViewProjectFile) HasReactions() bool`

HasReactions returns a boolean if a field has been set.

### GetRelatedItems

`func (o *ViewProjectFile) GetRelatedItems() ViewRelatedItems`

GetRelatedItems returns the RelatedItems field if non-nil, zero value otherwise.

### GetRelatedItemsOk

`func (o *ViewProjectFile) GetRelatedItemsOk() (*ViewRelatedItems, bool)`

GetRelatedItemsOk returns a tuple with the RelatedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedItems

`func (o *ViewProjectFile) SetRelatedItems(v ViewRelatedItems)`

SetRelatedItems sets RelatedItems field to given value.

### HasRelatedItems

`func (o *ViewProjectFile) HasRelatedItems() bool`

HasRelatedItems returns a boolean if a field has been set.

### GetSize

`func (o *ViewProjectFile) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ViewProjectFile) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ViewProjectFile) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ViewProjectFile) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStatus

`func (o *ViewProjectFile) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ViewProjectFile) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ViewProjectFile) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ViewProjectFile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTagIds

`func (o *ViewProjectFile) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *ViewProjectFile) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *ViewProjectFile) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *ViewProjectFile) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTags

`func (o *ViewProjectFile) GetTags() []ViewRelationship`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ViewProjectFile) GetTagsOk() (*[]ViewRelationship, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ViewProjectFile) SetTags(v []ViewRelationship)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ViewProjectFile) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThumbURL

`func (o *ViewProjectFile) GetThumbURL() string`

GetThumbURL returns the ThumbURL field if non-nil, zero value otherwise.

### GetThumbURLOk

`func (o *ViewProjectFile) GetThumbURLOk() (*string, bool)`

GetThumbURLOk returns a tuple with the ThumbURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbURL

`func (o *ViewProjectFile) SetThumbURL(v string)`

SetThumbURL sets ThumbURL field to given value.

### HasThumbURL

`func (o *ViewProjectFile) HasThumbURL() bool`

HasThumbURL returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ViewProjectFile) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ViewProjectFile) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ViewProjectFile) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ViewProjectFile) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUploadedBy

`func (o *ViewProjectFile) GetUploadedBy() int32`

GetUploadedBy returns the UploadedBy field if non-nil, zero value otherwise.

### GetUploadedByOk

`func (o *ViewProjectFile) GetUploadedByOk() (*int32, bool)`

GetUploadedByOk returns a tuple with the UploadedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedBy

`func (o *ViewProjectFile) SetUploadedBy(v int32)`

SetUploadedBy sets UploadedBy field to given value.

### HasUploadedBy

`func (o *ViewProjectFile) HasUploadedBy() bool`

HasUploadedBy returns a boolean if a field has been set.

### GetUploadedByUserID

`func (o *ViewProjectFile) GetUploadedByUserID() int32`

GetUploadedByUserID returns the UploadedByUserID field if non-nil, zero value otherwise.

### GetUploadedByUserIDOk

`func (o *ViewProjectFile) GetUploadedByUserIDOk() (*int32, bool)`

GetUploadedByUserIDOk returns a tuple with the UploadedByUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedByUserID

`func (o *ViewProjectFile) SetUploadedByUserID(v int32)`

SetUploadedByUserID sets UploadedByUserID field to given value.

### HasUploadedByUserID

`func (o *ViewProjectFile) HasUploadedByUserID() bool`

HasUploadedByUserID returns a boolean if a field has been set.

### GetUploadedDate

`func (o *ViewProjectFile) GetUploadedDate() string`

GetUploadedDate returns the UploadedDate field if non-nil, zero value otherwise.

### GetUploadedDateOk

`func (o *ViewProjectFile) GetUploadedDateOk() (*string, bool)`

GetUploadedDateOk returns a tuple with the UploadedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedDate

`func (o *ViewProjectFile) SetUploadedDate(v string)`

SetUploadedDate sets UploadedDate field to given value.

### HasUploadedDate

`func (o *ViewProjectFile) HasUploadedDate() bool`

HasUploadedDate returns a boolean if a field has been set.

### GetVersion

`func (o *ViewProjectFile) GetVersion() ViewRelationship`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ViewProjectFile) GetVersionOk() (*ViewRelationship, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ViewProjectFile) SetVersion(v ViewRelationship)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ViewProjectFile) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionId

`func (o *ViewProjectFile) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *ViewProjectFile) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *ViewProjectFile) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *ViewProjectFile) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetVersions

`func (o *ViewProjectFile) GetVersions() []ViewFileversion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ViewProjectFile) GetVersionsOk() (*[]ViewFileversion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ViewProjectFile) SetVersions(v []ViewFileversion)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *ViewProjectFile) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


