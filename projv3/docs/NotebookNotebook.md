# NotebookNotebook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | Pointer to **int32** |  | [optional] 
**Contents** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IsPrivate** | Pointer to **bool** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NewVersion** | Pointer to **bool** |  | [optional] 
**Notify** | Pointer to [**PayloadNotify**](PayloadNotify.md) |  | [optional] 
**NotifyCurrentUser** | Pointer to **bool** |  | [optional] 
**Privacy** | Pointer to [**PayloadUserGroups**](PayloadUserGroups.md) |  | [optional] 
**SecureContent** | Pointer to **bool** |  | [optional] 
**SendDiff** | Pointer to **bool** |  | [optional] 
**TagIds** | Pointer to **[]int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewNotebookNotebook

`func NewNotebookNotebook() *NotebookNotebook`

NewNotebookNotebook instantiates a new NotebookNotebook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotebookNotebookWithDefaults

`func NewNotebookNotebookWithDefaults() *NotebookNotebook`

NewNotebookNotebookWithDefaults instantiates a new NotebookNotebook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *NotebookNotebook) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *NotebookNotebook) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *NotebookNotebook) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *NotebookNotebook) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetContents

`func (o *NotebookNotebook) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *NotebookNotebook) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *NotebookNotebook) SetContents(v string)`

SetContents sets Contents field to given value.

### HasContents

`func (o *NotebookNotebook) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetDescription

`func (o *NotebookNotebook) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NotebookNotebook) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NotebookNotebook) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NotebookNotebook) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsPrivate

`func (o *NotebookNotebook) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *NotebookNotebook) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *NotebookNotebook) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *NotebookNotebook) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetLocked

`func (o *NotebookNotebook) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *NotebookNotebook) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *NotebookNotebook) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *NotebookNotebook) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetName

`func (o *NotebookNotebook) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotebookNotebook) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotebookNotebook) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NotebookNotebook) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNewVersion

`func (o *NotebookNotebook) GetNewVersion() bool`

GetNewVersion returns the NewVersion field if non-nil, zero value otherwise.

### GetNewVersionOk

`func (o *NotebookNotebook) GetNewVersionOk() (*bool, bool)`

GetNewVersionOk returns a tuple with the NewVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewVersion

`func (o *NotebookNotebook) SetNewVersion(v bool)`

SetNewVersion sets NewVersion field to given value.

### HasNewVersion

`func (o *NotebookNotebook) HasNewVersion() bool`

HasNewVersion returns a boolean if a field has been set.

### GetNotify

`func (o *NotebookNotebook) GetNotify() PayloadNotify`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *NotebookNotebook) GetNotifyOk() (*PayloadNotify, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *NotebookNotebook) SetNotify(v PayloadNotify)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *NotebookNotebook) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetNotifyCurrentUser

`func (o *NotebookNotebook) GetNotifyCurrentUser() bool`

GetNotifyCurrentUser returns the NotifyCurrentUser field if non-nil, zero value otherwise.

### GetNotifyCurrentUserOk

`func (o *NotebookNotebook) GetNotifyCurrentUserOk() (*bool, bool)`

GetNotifyCurrentUserOk returns a tuple with the NotifyCurrentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyCurrentUser

`func (o *NotebookNotebook) SetNotifyCurrentUser(v bool)`

SetNotifyCurrentUser sets NotifyCurrentUser field to given value.

### HasNotifyCurrentUser

`func (o *NotebookNotebook) HasNotifyCurrentUser() bool`

HasNotifyCurrentUser returns a boolean if a field has been set.

### GetPrivacy

`func (o *NotebookNotebook) GetPrivacy() PayloadUserGroups`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *NotebookNotebook) GetPrivacyOk() (*PayloadUserGroups, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *NotebookNotebook) SetPrivacy(v PayloadUserGroups)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *NotebookNotebook) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetSecureContent

`func (o *NotebookNotebook) GetSecureContent() bool`

GetSecureContent returns the SecureContent field if non-nil, zero value otherwise.

### GetSecureContentOk

`func (o *NotebookNotebook) GetSecureContentOk() (*bool, bool)`

GetSecureContentOk returns a tuple with the SecureContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureContent

`func (o *NotebookNotebook) SetSecureContent(v bool)`

SetSecureContent sets SecureContent field to given value.

### HasSecureContent

`func (o *NotebookNotebook) HasSecureContent() bool`

HasSecureContent returns a boolean if a field has been set.

### GetSendDiff

`func (o *NotebookNotebook) GetSendDiff() bool`

GetSendDiff returns the SendDiff field if non-nil, zero value otherwise.

### GetSendDiffOk

`func (o *NotebookNotebook) GetSendDiffOk() (*bool, bool)`

GetSendDiffOk returns a tuple with the SendDiff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDiff

`func (o *NotebookNotebook) SetSendDiff(v bool)`

SetSendDiff sets SendDiff field to given value.

### HasSendDiff

`func (o *NotebookNotebook) HasSendDiff() bool`

HasSendDiff returns a boolean if a field has been set.

### GetTagIds

`func (o *NotebookNotebook) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *NotebookNotebook) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *NotebookNotebook) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *NotebookNotebook) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetType

`func (o *NotebookNotebook) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotebookNotebook) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotebookNotebook) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NotebookNotebook) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


