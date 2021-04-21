# InlineResponse200115

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | Pointer to **[]map[string]interface{}** |  | [optional] 
**AttachmentsCount** | Pointer to **string** |  | [optional] 
**Author** | Pointer to [**InlineResponse200115Author**](InlineResponse200115Author.md) |  | [optional] 
**Body** | Pointer to **string** |  | [optional] 
**CanDelete** | Pointer to **bool** |  | [optional] 
**CanEditContents** | Pointer to **bool** |  | [optional] 
**Category** | Pointer to [**InlineResponse2002Column**](InlineResponse2002Column.md) |  | [optional] 
**CommentsCount** | Pointer to **string** |  | [optional] 
**Company** | Pointer to [**InlineResponse200115Company**](InlineResponse200115Company.md) |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**ContributingUsers** | Pointer to [**[]InlineResponse200115ContributingUsers**](InlineResponse200115ContributingUsers.md) |  | [optional] 
**FollowerIds** | Pointer to **string** |  | [optional] 
**HasEditExpired** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsOriginal** | Pointer to **string** |  | [optional] 
**IsRead** | Pointer to **string** |  | [optional] 
**LastChangedOn** | Pointer to **string** |  | [optional] 
**LastCommentDate** | Pointer to **string** |  | [optional] 
**MessageStatus** | Pointer to **string** |  | [optional] 
**NumNotified** | Pointer to **string** |  | [optional] 
**PostId** | Pointer to **string** |  | [optional] 
**PostStatus** | Pointer to **string** |  | [optional] 
**PostedOn** | Pointer to **string** |  | [optional] 
**Private** | Pointer to **string** |  | [optional] 
**Project** | Pointer to [**InlineResponse2002People12345Company**](InlineResponse2002People12345Company.md) |  | [optional] 
**Reactions** | Pointer to [**InlineResponse2008Reactions**](InlineResponse2008Reactions.md) |  | [optional] 
**Tags** | Pointer to **[]map[string]interface{}** |  | [optional] 
**TextBody** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UserDisplayPostedDate** | Pointer to **string** |  | [optional] 
**UserDisplayPostedTime** | Pointer to **string** |  | [optional] 
**UserFollowing** | Pointer to **bool** |  | [optional] 

## Methods

### NewInlineResponse200115

`func NewInlineResponse200115() *InlineResponse200115`

NewInlineResponse200115 instantiates a new InlineResponse200115 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200115WithDefaults

`func NewInlineResponse200115WithDefaults() *InlineResponse200115`

NewInlineResponse200115WithDefaults instantiates a new InlineResponse200115 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *InlineResponse200115) GetAttachments() []map[string]interface{}`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *InlineResponse200115) GetAttachmentsOk() (*[]map[string]interface{}, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *InlineResponse200115) SetAttachments(v []map[string]interface{})`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *InlineResponse200115) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetAttachmentsCount

`func (o *InlineResponse200115) GetAttachmentsCount() string`

GetAttachmentsCount returns the AttachmentsCount field if non-nil, zero value otherwise.

### GetAttachmentsCountOk

`func (o *InlineResponse200115) GetAttachmentsCountOk() (*string, bool)`

GetAttachmentsCountOk returns a tuple with the AttachmentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentsCount

`func (o *InlineResponse200115) SetAttachmentsCount(v string)`

SetAttachmentsCount sets AttachmentsCount field to given value.

### HasAttachmentsCount

`func (o *InlineResponse200115) HasAttachmentsCount() bool`

HasAttachmentsCount returns a boolean if a field has been set.

### GetAuthor

`func (o *InlineResponse200115) GetAuthor() InlineResponse200115Author`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *InlineResponse200115) GetAuthorOk() (*InlineResponse200115Author, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *InlineResponse200115) SetAuthor(v InlineResponse200115Author)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *InlineResponse200115) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetBody

`func (o *InlineResponse200115) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *InlineResponse200115) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *InlineResponse200115) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *InlineResponse200115) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCanDelete

`func (o *InlineResponse200115) GetCanDelete() bool`

GetCanDelete returns the CanDelete field if non-nil, zero value otherwise.

### GetCanDeleteOk

`func (o *InlineResponse200115) GetCanDeleteOk() (*bool, bool)`

GetCanDeleteOk returns a tuple with the CanDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDelete

`func (o *InlineResponse200115) SetCanDelete(v bool)`

SetCanDelete sets CanDelete field to given value.

### HasCanDelete

`func (o *InlineResponse200115) HasCanDelete() bool`

HasCanDelete returns a boolean if a field has been set.

### GetCanEditContents

`func (o *InlineResponse200115) GetCanEditContents() bool`

GetCanEditContents returns the CanEditContents field if non-nil, zero value otherwise.

### GetCanEditContentsOk

`func (o *InlineResponse200115) GetCanEditContentsOk() (*bool, bool)`

GetCanEditContentsOk returns a tuple with the CanEditContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEditContents

`func (o *InlineResponse200115) SetCanEditContents(v bool)`

SetCanEditContents sets CanEditContents field to given value.

### HasCanEditContents

`func (o *InlineResponse200115) HasCanEditContents() bool`

HasCanEditContents returns a boolean if a field has been set.

### GetCategory

`func (o *InlineResponse200115) GetCategory() InlineResponse2002Column`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InlineResponse200115) GetCategoryOk() (*InlineResponse2002Column, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InlineResponse200115) SetCategory(v InlineResponse2002Column)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *InlineResponse200115) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCommentsCount

`func (o *InlineResponse200115) GetCommentsCount() string`

GetCommentsCount returns the CommentsCount field if non-nil, zero value otherwise.

### GetCommentsCountOk

`func (o *InlineResponse200115) GetCommentsCountOk() (*string, bool)`

GetCommentsCountOk returns a tuple with the CommentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsCount

`func (o *InlineResponse200115) SetCommentsCount(v string)`

SetCommentsCount sets CommentsCount field to given value.

### HasCommentsCount

`func (o *InlineResponse200115) HasCommentsCount() bool`

HasCommentsCount returns a boolean if a field has been set.

### GetCompany

`func (o *InlineResponse200115) GetCompany() InlineResponse200115Company`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *InlineResponse200115) GetCompanyOk() (*InlineResponse200115Company, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *InlineResponse200115) SetCompany(v InlineResponse200115Company)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *InlineResponse200115) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetContentType

`func (o *InlineResponse200115) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *InlineResponse200115) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *InlineResponse200115) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *InlineResponse200115) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetContributingUsers

`func (o *InlineResponse200115) GetContributingUsers() []InlineResponse200115ContributingUsers`

GetContributingUsers returns the ContributingUsers field if non-nil, zero value otherwise.

### GetContributingUsersOk

`func (o *InlineResponse200115) GetContributingUsersOk() (*[]InlineResponse200115ContributingUsers, bool)`

GetContributingUsersOk returns a tuple with the ContributingUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributingUsers

`func (o *InlineResponse200115) SetContributingUsers(v []InlineResponse200115ContributingUsers)`

SetContributingUsers sets ContributingUsers field to given value.

### HasContributingUsers

`func (o *InlineResponse200115) HasContributingUsers() bool`

HasContributingUsers returns a boolean if a field has been set.

### GetFollowerIds

`func (o *InlineResponse200115) GetFollowerIds() string`

GetFollowerIds returns the FollowerIds field if non-nil, zero value otherwise.

### GetFollowerIdsOk

`func (o *InlineResponse200115) GetFollowerIdsOk() (*string, bool)`

GetFollowerIdsOk returns a tuple with the FollowerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerIds

`func (o *InlineResponse200115) SetFollowerIds(v string)`

SetFollowerIds sets FollowerIds field to given value.

### HasFollowerIds

`func (o *InlineResponse200115) HasFollowerIds() bool`

HasFollowerIds returns a boolean if a field has been set.

### GetHasEditExpired

`func (o *InlineResponse200115) GetHasEditExpired() bool`

GetHasEditExpired returns the HasEditExpired field if non-nil, zero value otherwise.

### GetHasEditExpiredOk

`func (o *InlineResponse200115) GetHasEditExpiredOk() (*bool, bool)`

GetHasEditExpiredOk returns a tuple with the HasEditExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEditExpired

`func (o *InlineResponse200115) SetHasEditExpired(v bool)`

SetHasEditExpired sets HasEditExpired field to given value.

### HasHasEditExpired

`func (o *InlineResponse200115) HasHasEditExpired() bool`

HasHasEditExpired returns a boolean if a field has been set.

### GetId

`func (o *InlineResponse200115) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200115) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200115) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200115) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsOriginal

`func (o *InlineResponse200115) GetIsOriginal() string`

GetIsOriginal returns the IsOriginal field if non-nil, zero value otherwise.

### GetIsOriginalOk

`func (o *InlineResponse200115) GetIsOriginalOk() (*string, bool)`

GetIsOriginalOk returns a tuple with the IsOriginal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOriginal

`func (o *InlineResponse200115) SetIsOriginal(v string)`

SetIsOriginal sets IsOriginal field to given value.

### HasIsOriginal

`func (o *InlineResponse200115) HasIsOriginal() bool`

HasIsOriginal returns a boolean if a field has been set.

### GetIsRead

`func (o *InlineResponse200115) GetIsRead() string`

GetIsRead returns the IsRead field if non-nil, zero value otherwise.

### GetIsReadOk

`func (o *InlineResponse200115) GetIsReadOk() (*string, bool)`

GetIsReadOk returns a tuple with the IsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRead

`func (o *InlineResponse200115) SetIsRead(v string)`

SetIsRead sets IsRead field to given value.

### HasIsRead

`func (o *InlineResponse200115) HasIsRead() bool`

HasIsRead returns a boolean if a field has been set.

### GetLastChangedOn

`func (o *InlineResponse200115) GetLastChangedOn() string`

GetLastChangedOn returns the LastChangedOn field if non-nil, zero value otherwise.

### GetLastChangedOnOk

`func (o *InlineResponse200115) GetLastChangedOnOk() (*string, bool)`

GetLastChangedOnOk returns a tuple with the LastChangedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangedOn

`func (o *InlineResponse200115) SetLastChangedOn(v string)`

SetLastChangedOn sets LastChangedOn field to given value.

### HasLastChangedOn

`func (o *InlineResponse200115) HasLastChangedOn() bool`

HasLastChangedOn returns a boolean if a field has been set.

### GetLastCommentDate

`func (o *InlineResponse200115) GetLastCommentDate() string`

GetLastCommentDate returns the LastCommentDate field if non-nil, zero value otherwise.

### GetLastCommentDateOk

`func (o *InlineResponse200115) GetLastCommentDateOk() (*string, bool)`

GetLastCommentDateOk returns a tuple with the LastCommentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCommentDate

`func (o *InlineResponse200115) SetLastCommentDate(v string)`

SetLastCommentDate sets LastCommentDate field to given value.

### HasLastCommentDate

`func (o *InlineResponse200115) HasLastCommentDate() bool`

HasLastCommentDate returns a boolean if a field has been set.

### GetMessageStatus

`func (o *InlineResponse200115) GetMessageStatus() string`

GetMessageStatus returns the MessageStatus field if non-nil, zero value otherwise.

### GetMessageStatusOk

`func (o *InlineResponse200115) GetMessageStatusOk() (*string, bool)`

GetMessageStatusOk returns a tuple with the MessageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageStatus

`func (o *InlineResponse200115) SetMessageStatus(v string)`

SetMessageStatus sets MessageStatus field to given value.

### HasMessageStatus

`func (o *InlineResponse200115) HasMessageStatus() bool`

HasMessageStatus returns a boolean if a field has been set.

### GetNumNotified

`func (o *InlineResponse200115) GetNumNotified() string`

GetNumNotified returns the NumNotified field if non-nil, zero value otherwise.

### GetNumNotifiedOk

`func (o *InlineResponse200115) GetNumNotifiedOk() (*string, bool)`

GetNumNotifiedOk returns a tuple with the NumNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumNotified

`func (o *InlineResponse200115) SetNumNotified(v string)`

SetNumNotified sets NumNotified field to given value.

### HasNumNotified

`func (o *InlineResponse200115) HasNumNotified() bool`

HasNumNotified returns a boolean if a field has been set.

### GetPostId

`func (o *InlineResponse200115) GetPostId() string`

GetPostId returns the PostId field if non-nil, zero value otherwise.

### GetPostIdOk

`func (o *InlineResponse200115) GetPostIdOk() (*string, bool)`

GetPostIdOk returns a tuple with the PostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostId

`func (o *InlineResponse200115) SetPostId(v string)`

SetPostId sets PostId field to given value.

### HasPostId

`func (o *InlineResponse200115) HasPostId() bool`

HasPostId returns a boolean if a field has been set.

### GetPostStatus

`func (o *InlineResponse200115) GetPostStatus() string`

GetPostStatus returns the PostStatus field if non-nil, zero value otherwise.

### GetPostStatusOk

`func (o *InlineResponse200115) GetPostStatusOk() (*string, bool)`

GetPostStatusOk returns a tuple with the PostStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostStatus

`func (o *InlineResponse200115) SetPostStatus(v string)`

SetPostStatus sets PostStatus field to given value.

### HasPostStatus

`func (o *InlineResponse200115) HasPostStatus() bool`

HasPostStatus returns a boolean if a field has been set.

### GetPostedOn

`func (o *InlineResponse200115) GetPostedOn() string`

GetPostedOn returns the PostedOn field if non-nil, zero value otherwise.

### GetPostedOnOk

`func (o *InlineResponse200115) GetPostedOnOk() (*string, bool)`

GetPostedOnOk returns a tuple with the PostedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedOn

`func (o *InlineResponse200115) SetPostedOn(v string)`

SetPostedOn sets PostedOn field to given value.

### HasPostedOn

`func (o *InlineResponse200115) HasPostedOn() bool`

HasPostedOn returns a boolean if a field has been set.

### GetPrivate

`func (o *InlineResponse200115) GetPrivate() string`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *InlineResponse200115) GetPrivateOk() (*string, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *InlineResponse200115) SetPrivate(v string)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *InlineResponse200115) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetProject

`func (o *InlineResponse200115) GetProject() InlineResponse2002People12345Company`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *InlineResponse200115) GetProjectOk() (*InlineResponse2002People12345Company, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *InlineResponse200115) SetProject(v InlineResponse2002People12345Company)`

SetProject sets Project field to given value.

### HasProject

`func (o *InlineResponse200115) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetReactions

`func (o *InlineResponse200115) GetReactions() InlineResponse2008Reactions`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *InlineResponse200115) GetReactionsOk() (*InlineResponse2008Reactions, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *InlineResponse200115) SetReactions(v InlineResponse2008Reactions)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *InlineResponse200115) HasReactions() bool`

HasReactions returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse200115) GetTags() []map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse200115) GetTagsOk() (*[]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse200115) SetTags(v []map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse200115) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTextBody

`func (o *InlineResponse200115) GetTextBody() string`

GetTextBody returns the TextBody field if non-nil, zero value otherwise.

### GetTextBodyOk

`func (o *InlineResponse200115) GetTextBodyOk() (*string, bool)`

GetTextBodyOk returns a tuple with the TextBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextBody

`func (o *InlineResponse200115) SetTextBody(v string)`

SetTextBody sets TextBody field to given value.

### HasTextBody

`func (o *InlineResponse200115) HasTextBody() bool`

HasTextBody returns a boolean if a field has been set.

### GetTitle

`func (o *InlineResponse200115) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InlineResponse200115) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InlineResponse200115) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InlineResponse200115) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUserDisplayPostedDate

`func (o *InlineResponse200115) GetUserDisplayPostedDate() string`

GetUserDisplayPostedDate returns the UserDisplayPostedDate field if non-nil, zero value otherwise.

### GetUserDisplayPostedDateOk

`func (o *InlineResponse200115) GetUserDisplayPostedDateOk() (*string, bool)`

GetUserDisplayPostedDateOk returns a tuple with the UserDisplayPostedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDisplayPostedDate

`func (o *InlineResponse200115) SetUserDisplayPostedDate(v string)`

SetUserDisplayPostedDate sets UserDisplayPostedDate field to given value.

### HasUserDisplayPostedDate

`func (o *InlineResponse200115) HasUserDisplayPostedDate() bool`

HasUserDisplayPostedDate returns a boolean if a field has been set.

### GetUserDisplayPostedTime

`func (o *InlineResponse200115) GetUserDisplayPostedTime() string`

GetUserDisplayPostedTime returns the UserDisplayPostedTime field if non-nil, zero value otherwise.

### GetUserDisplayPostedTimeOk

`func (o *InlineResponse200115) GetUserDisplayPostedTimeOk() (*string, bool)`

GetUserDisplayPostedTimeOk returns a tuple with the UserDisplayPostedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDisplayPostedTime

`func (o *InlineResponse200115) SetUserDisplayPostedTime(v string)`

SetUserDisplayPostedTime sets UserDisplayPostedTime field to given value.

### HasUserDisplayPostedTime

`func (o *InlineResponse200115) HasUserDisplayPostedTime() bool`

HasUserDisplayPostedTime returns a boolean if a field has been set.

### GetUserFollowing

`func (o *InlineResponse200115) GetUserFollowing() bool`

GetUserFollowing returns the UserFollowing field if non-nil, zero value otherwise.

### GetUserFollowingOk

`func (o *InlineResponse200115) GetUserFollowingOk() (*bool, bool)`

GetUserFollowingOk returns a tuple with the UserFollowing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFollowing

`func (o *InlineResponse200115) SetUserFollowing(v bool)`

SetUserFollowing sets UserFollowing field to given value.

### HasUserFollowing

`func (o *InlineResponse200115) HasUserFollowing() bool`

HasUserFollowing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


