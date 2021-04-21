# FileProjectFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Notify** | Pointer to [**PayloadNotify**](PayloadNotify.md) |  | [optional] 
**NotifyCurrentUser** | Pointer to **bool** |  | [optional] 
**Privacy** | Pointer to [**PayloadUserGroups**](PayloadUserGroups.md) |  | [optional] 
**Private** | Pointer to **bool** |  | [optional] 
**TagIds** | Pointer to [**PayloadNullableInt64Slice**](PayloadNullableInt64Slice.md) |  | [optional] 
**VersionId** | Pointer to **int32** |  | [optional] 

## Methods

### NewFileProjectFile

`func NewFileProjectFile() *FileProjectFile`

NewFileProjectFile instantiates a new FileProjectFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileProjectFileWithDefaults

`func NewFileProjectFileWithDefaults() *FileProjectFile`

NewFileProjectFileWithDefaults instantiates a new FileProjectFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *FileProjectFile) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *FileProjectFile) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *FileProjectFile) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *FileProjectFile) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetDescription

`func (o *FileProjectFile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FileProjectFile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FileProjectFile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FileProjectFile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *FileProjectFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileProjectFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileProjectFile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FileProjectFile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotify

`func (o *FileProjectFile) GetNotify() PayloadNotify`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *FileProjectFile) GetNotifyOk() (*PayloadNotify, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *FileProjectFile) SetNotify(v PayloadNotify)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *FileProjectFile) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetNotifyCurrentUser

`func (o *FileProjectFile) GetNotifyCurrentUser() bool`

GetNotifyCurrentUser returns the NotifyCurrentUser field if non-nil, zero value otherwise.

### GetNotifyCurrentUserOk

`func (o *FileProjectFile) GetNotifyCurrentUserOk() (*bool, bool)`

GetNotifyCurrentUserOk returns a tuple with the NotifyCurrentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyCurrentUser

`func (o *FileProjectFile) SetNotifyCurrentUser(v bool)`

SetNotifyCurrentUser sets NotifyCurrentUser field to given value.

### HasNotifyCurrentUser

`func (o *FileProjectFile) HasNotifyCurrentUser() bool`

HasNotifyCurrentUser returns a boolean if a field has been set.

### GetPrivacy

`func (o *FileProjectFile) GetPrivacy() PayloadUserGroups`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *FileProjectFile) GetPrivacyOk() (*PayloadUserGroups, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *FileProjectFile) SetPrivacy(v PayloadUserGroups)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *FileProjectFile) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetPrivate

`func (o *FileProjectFile) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *FileProjectFile) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *FileProjectFile) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *FileProjectFile) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetTagIds

`func (o *FileProjectFile) GetTagIds() PayloadNullableInt64Slice`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *FileProjectFile) GetTagIdsOk() (*PayloadNullableInt64Slice, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *FileProjectFile) SetTagIds(v PayloadNullableInt64Slice)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *FileProjectFile) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetVersionId

`func (o *FileProjectFile) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *FileProjectFile) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *FileProjectFile) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *FileProjectFile) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


