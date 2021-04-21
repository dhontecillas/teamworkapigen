# InlineResponse200116MessageReplies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | Pointer to **[]map[string]interface{}** |  | [optional] 
**AttachmentsCount** | Pointer to **string** |  | [optional] 
**Author** | Pointer to [**InlineResponse200115Author**](InlineResponse200115Author.md) |  | [optional] 
**Body** | Pointer to **string** |  | [optional] 
**CanDelete** | Pointer to **bool** |  | [optional] 
**CanEditContents** | Pointer to **bool** |  | [optional] 
**Company** | Pointer to [**InlineResponse200115Company**](InlineResponse200115Company.md) |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**Datetime** | Pointer to **string** |  | [optional] 
**HasEditExpired** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsOriginal** | Pointer to **string** |  | [optional] 
**IsRead** | Pointer to **string** |  | [optional] 
**LastChangedOn** | Pointer to **string** |  | [optional] 
**LockdownId** | Pointer to **string** |  | [optional] 
**MessageId** | Pointer to **string** |  | [optional] 
**MessageStatus** | Pointer to **string** |  | [optional] 
**NumNotified** | Pointer to **string** |  | [optional] 
**PostStatus** | Pointer to **string** |  | [optional] 
**PostedOn** | Pointer to **string** |  | [optional] 
**Private** | Pointer to **bool** |  | [optional] 
**Project** | Pointer to [**InlineResponse2002People12345Company**](InlineResponse2002People12345Company.md) |  | [optional] 
**Reactions** | Pointer to [**InlineResponse2008Reactions**](InlineResponse2008Reactions.md) |  | [optional] 
**ReplyNo** | Pointer to **string** |  | [optional] 
**TextBody** | Pointer to **string** |  | [optional] 
**UserDisplayPostedDate** | Pointer to **string** |  | [optional] 
**UserDisplayPostedTime** | Pointer to **string** |  | [optional] 

## Methods

### NewInlineResponse200116MessageReplies

`func NewInlineResponse200116MessageReplies() *InlineResponse200116MessageReplies`

NewInlineResponse200116MessageReplies instantiates a new InlineResponse200116MessageReplies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200116MessageRepliesWithDefaults

`func NewInlineResponse200116MessageRepliesWithDefaults() *InlineResponse200116MessageReplies`

NewInlineResponse200116MessageRepliesWithDefaults instantiates a new InlineResponse200116MessageReplies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *InlineResponse200116MessageReplies) GetAttachments() []map[string]interface{}`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *InlineResponse200116MessageReplies) GetAttachmentsOk() (*[]map[string]interface{}, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *InlineResponse200116MessageReplies) SetAttachments(v []map[string]interface{})`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *InlineResponse200116MessageReplies) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetAttachmentsCount

`func (o *InlineResponse200116MessageReplies) GetAttachmentsCount() string`

GetAttachmentsCount returns the AttachmentsCount field if non-nil, zero value otherwise.

### GetAttachmentsCountOk

`func (o *InlineResponse200116MessageReplies) GetAttachmentsCountOk() (*string, bool)`

GetAttachmentsCountOk returns a tuple with the AttachmentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentsCount

`func (o *InlineResponse200116MessageReplies) SetAttachmentsCount(v string)`

SetAttachmentsCount sets AttachmentsCount field to given value.

### HasAttachmentsCount

`func (o *InlineResponse200116MessageReplies) HasAttachmentsCount() bool`

HasAttachmentsCount returns a boolean if a field has been set.

### GetAuthor

`func (o *InlineResponse200116MessageReplies) GetAuthor() InlineResponse200115Author`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *InlineResponse200116MessageReplies) GetAuthorOk() (*InlineResponse200115Author, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *InlineResponse200116MessageReplies) SetAuthor(v InlineResponse200115Author)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *InlineResponse200116MessageReplies) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetBody

`func (o *InlineResponse200116MessageReplies) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *InlineResponse200116MessageReplies) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *InlineResponse200116MessageReplies) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *InlineResponse200116MessageReplies) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCanDelete

`func (o *InlineResponse200116MessageReplies) GetCanDelete() bool`

GetCanDelete returns the CanDelete field if non-nil, zero value otherwise.

### GetCanDeleteOk

`func (o *InlineResponse200116MessageReplies) GetCanDeleteOk() (*bool, bool)`

GetCanDeleteOk returns a tuple with the CanDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDelete

`func (o *InlineResponse200116MessageReplies) SetCanDelete(v bool)`

SetCanDelete sets CanDelete field to given value.

### HasCanDelete

`func (o *InlineResponse200116MessageReplies) HasCanDelete() bool`

HasCanDelete returns a boolean if a field has been set.

### GetCanEditContents

`func (o *InlineResponse200116MessageReplies) GetCanEditContents() bool`

GetCanEditContents returns the CanEditContents field if non-nil, zero value otherwise.

### GetCanEditContentsOk

`func (o *InlineResponse200116MessageReplies) GetCanEditContentsOk() (*bool, bool)`

GetCanEditContentsOk returns a tuple with the CanEditContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEditContents

`func (o *InlineResponse200116MessageReplies) SetCanEditContents(v bool)`

SetCanEditContents sets CanEditContents field to given value.

### HasCanEditContents

`func (o *InlineResponse200116MessageReplies) HasCanEditContents() bool`

HasCanEditContents returns a boolean if a field has been set.

### GetCompany

`func (o *InlineResponse200116MessageReplies) GetCompany() InlineResponse200115Company`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *InlineResponse200116MessageReplies) GetCompanyOk() (*InlineResponse200115Company, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *InlineResponse200116MessageReplies) SetCompany(v InlineResponse200115Company)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *InlineResponse200116MessageReplies) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetContentType

`func (o *InlineResponse200116MessageReplies) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *InlineResponse200116MessageReplies) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *InlineResponse200116MessageReplies) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *InlineResponse200116MessageReplies) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetDatetime

`func (o *InlineResponse200116MessageReplies) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *InlineResponse200116MessageReplies) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *InlineResponse200116MessageReplies) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *InlineResponse200116MessageReplies) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetHasEditExpired

`func (o *InlineResponse200116MessageReplies) GetHasEditExpired() bool`

GetHasEditExpired returns the HasEditExpired field if non-nil, zero value otherwise.

### GetHasEditExpiredOk

`func (o *InlineResponse200116MessageReplies) GetHasEditExpiredOk() (*bool, bool)`

GetHasEditExpiredOk returns a tuple with the HasEditExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEditExpired

`func (o *InlineResponse200116MessageReplies) SetHasEditExpired(v bool)`

SetHasEditExpired sets HasEditExpired field to given value.

### HasHasEditExpired

`func (o *InlineResponse200116MessageReplies) HasHasEditExpired() bool`

HasHasEditExpired returns a boolean if a field has been set.

### GetId

`func (o *InlineResponse200116MessageReplies) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200116MessageReplies) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200116MessageReplies) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200116MessageReplies) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsOriginal

`func (o *InlineResponse200116MessageReplies) GetIsOriginal() string`

GetIsOriginal returns the IsOriginal field if non-nil, zero value otherwise.

### GetIsOriginalOk

`func (o *InlineResponse200116MessageReplies) GetIsOriginalOk() (*string, bool)`

GetIsOriginalOk returns a tuple with the IsOriginal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOriginal

`func (o *InlineResponse200116MessageReplies) SetIsOriginal(v string)`

SetIsOriginal sets IsOriginal field to given value.

### HasIsOriginal

`func (o *InlineResponse200116MessageReplies) HasIsOriginal() bool`

HasIsOriginal returns a boolean if a field has been set.

### GetIsRead

`func (o *InlineResponse200116MessageReplies) GetIsRead() string`

GetIsRead returns the IsRead field if non-nil, zero value otherwise.

### GetIsReadOk

`func (o *InlineResponse200116MessageReplies) GetIsReadOk() (*string, bool)`

GetIsReadOk returns a tuple with the IsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRead

`func (o *InlineResponse200116MessageReplies) SetIsRead(v string)`

SetIsRead sets IsRead field to given value.

### HasIsRead

`func (o *InlineResponse200116MessageReplies) HasIsRead() bool`

HasIsRead returns a boolean if a field has been set.

### GetLastChangedOn

`func (o *InlineResponse200116MessageReplies) GetLastChangedOn() string`

GetLastChangedOn returns the LastChangedOn field if non-nil, zero value otherwise.

### GetLastChangedOnOk

`func (o *InlineResponse200116MessageReplies) GetLastChangedOnOk() (*string, bool)`

GetLastChangedOnOk returns a tuple with the LastChangedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangedOn

`func (o *InlineResponse200116MessageReplies) SetLastChangedOn(v string)`

SetLastChangedOn sets LastChangedOn field to given value.

### HasLastChangedOn

`func (o *InlineResponse200116MessageReplies) HasLastChangedOn() bool`

HasLastChangedOn returns a boolean if a field has been set.

### GetLockdownId

`func (o *InlineResponse200116MessageReplies) GetLockdownId() string`

GetLockdownId returns the LockdownId field if non-nil, zero value otherwise.

### GetLockdownIdOk

`func (o *InlineResponse200116MessageReplies) GetLockdownIdOk() (*string, bool)`

GetLockdownIdOk returns a tuple with the LockdownId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockdownId

`func (o *InlineResponse200116MessageReplies) SetLockdownId(v string)`

SetLockdownId sets LockdownId field to given value.

### HasLockdownId

`func (o *InlineResponse200116MessageReplies) HasLockdownId() bool`

HasLockdownId returns a boolean if a field has been set.

### GetMessageId

`func (o *InlineResponse200116MessageReplies) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *InlineResponse200116MessageReplies) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *InlineResponse200116MessageReplies) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *InlineResponse200116MessageReplies) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMessageStatus

`func (o *InlineResponse200116MessageReplies) GetMessageStatus() string`

GetMessageStatus returns the MessageStatus field if non-nil, zero value otherwise.

### GetMessageStatusOk

`func (o *InlineResponse200116MessageReplies) GetMessageStatusOk() (*string, bool)`

GetMessageStatusOk returns a tuple with the MessageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageStatus

`func (o *InlineResponse200116MessageReplies) SetMessageStatus(v string)`

SetMessageStatus sets MessageStatus field to given value.

### HasMessageStatus

`func (o *InlineResponse200116MessageReplies) HasMessageStatus() bool`

HasMessageStatus returns a boolean if a field has been set.

### GetNumNotified

`func (o *InlineResponse200116MessageReplies) GetNumNotified() string`

GetNumNotified returns the NumNotified field if non-nil, zero value otherwise.

### GetNumNotifiedOk

`func (o *InlineResponse200116MessageReplies) GetNumNotifiedOk() (*string, bool)`

GetNumNotifiedOk returns a tuple with the NumNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumNotified

`func (o *InlineResponse200116MessageReplies) SetNumNotified(v string)`

SetNumNotified sets NumNotified field to given value.

### HasNumNotified

`func (o *InlineResponse200116MessageReplies) HasNumNotified() bool`

HasNumNotified returns a boolean if a field has been set.

### GetPostStatus

`func (o *InlineResponse200116MessageReplies) GetPostStatus() string`

GetPostStatus returns the PostStatus field if non-nil, zero value otherwise.

### GetPostStatusOk

`func (o *InlineResponse200116MessageReplies) GetPostStatusOk() (*string, bool)`

GetPostStatusOk returns a tuple with the PostStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostStatus

`func (o *InlineResponse200116MessageReplies) SetPostStatus(v string)`

SetPostStatus sets PostStatus field to given value.

### HasPostStatus

`func (o *InlineResponse200116MessageReplies) HasPostStatus() bool`

HasPostStatus returns a boolean if a field has been set.

### GetPostedOn

`func (o *InlineResponse200116MessageReplies) GetPostedOn() string`

GetPostedOn returns the PostedOn field if non-nil, zero value otherwise.

### GetPostedOnOk

`func (o *InlineResponse200116MessageReplies) GetPostedOnOk() (*string, bool)`

GetPostedOnOk returns a tuple with the PostedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedOn

`func (o *InlineResponse200116MessageReplies) SetPostedOn(v string)`

SetPostedOn sets PostedOn field to given value.

### HasPostedOn

`func (o *InlineResponse200116MessageReplies) HasPostedOn() bool`

HasPostedOn returns a boolean if a field has been set.

### GetPrivate

`func (o *InlineResponse200116MessageReplies) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *InlineResponse200116MessageReplies) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *InlineResponse200116MessageReplies) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *InlineResponse200116MessageReplies) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetProject

`func (o *InlineResponse200116MessageReplies) GetProject() InlineResponse2002People12345Company`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *InlineResponse200116MessageReplies) GetProjectOk() (*InlineResponse2002People12345Company, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *InlineResponse200116MessageReplies) SetProject(v InlineResponse2002People12345Company)`

SetProject sets Project field to given value.

### HasProject

`func (o *InlineResponse200116MessageReplies) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetReactions

`func (o *InlineResponse200116MessageReplies) GetReactions() InlineResponse2008Reactions`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *InlineResponse200116MessageReplies) GetReactionsOk() (*InlineResponse2008Reactions, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *InlineResponse200116MessageReplies) SetReactions(v InlineResponse2008Reactions)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *InlineResponse200116MessageReplies) HasReactions() bool`

HasReactions returns a boolean if a field has been set.

### GetReplyNo

`func (o *InlineResponse200116MessageReplies) GetReplyNo() string`

GetReplyNo returns the ReplyNo field if non-nil, zero value otherwise.

### GetReplyNoOk

`func (o *InlineResponse200116MessageReplies) GetReplyNoOk() (*string, bool)`

GetReplyNoOk returns a tuple with the ReplyNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyNo

`func (o *InlineResponse200116MessageReplies) SetReplyNo(v string)`

SetReplyNo sets ReplyNo field to given value.

### HasReplyNo

`func (o *InlineResponse200116MessageReplies) HasReplyNo() bool`

HasReplyNo returns a boolean if a field has been set.

### GetTextBody

`func (o *InlineResponse200116MessageReplies) GetTextBody() string`

GetTextBody returns the TextBody field if non-nil, zero value otherwise.

### GetTextBodyOk

`func (o *InlineResponse200116MessageReplies) GetTextBodyOk() (*string, bool)`

GetTextBodyOk returns a tuple with the TextBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextBody

`func (o *InlineResponse200116MessageReplies) SetTextBody(v string)`

SetTextBody sets TextBody field to given value.

### HasTextBody

`func (o *InlineResponse200116MessageReplies) HasTextBody() bool`

HasTextBody returns a boolean if a field has been set.

### GetUserDisplayPostedDate

`func (o *InlineResponse200116MessageReplies) GetUserDisplayPostedDate() string`

GetUserDisplayPostedDate returns the UserDisplayPostedDate field if non-nil, zero value otherwise.

### GetUserDisplayPostedDateOk

`func (o *InlineResponse200116MessageReplies) GetUserDisplayPostedDateOk() (*string, bool)`

GetUserDisplayPostedDateOk returns a tuple with the UserDisplayPostedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDisplayPostedDate

`func (o *InlineResponse200116MessageReplies) SetUserDisplayPostedDate(v string)`

SetUserDisplayPostedDate sets UserDisplayPostedDate field to given value.

### HasUserDisplayPostedDate

`func (o *InlineResponse200116MessageReplies) HasUserDisplayPostedDate() bool`

HasUserDisplayPostedDate returns a boolean if a field has been set.

### GetUserDisplayPostedTime

`func (o *InlineResponse200116MessageReplies) GetUserDisplayPostedTime() string`

GetUserDisplayPostedTime returns the UserDisplayPostedTime field if non-nil, zero value otherwise.

### GetUserDisplayPostedTimeOk

`func (o *InlineResponse200116MessageReplies) GetUserDisplayPostedTimeOk() (*string, bool)`

GetUserDisplayPostedTimeOk returns a tuple with the UserDisplayPostedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDisplayPostedTime

`func (o *InlineResponse200116MessageReplies) SetUserDisplayPostedTime(v string)`

SetUserDisplayPostedTime sets UserDisplayPostedTime field to given value.

### HasUserDisplayPostedTime

`func (o *InlineResponse200116MessageReplies) HasUserDisplayPostedTime() bool`

HasUserDisplayPostedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


