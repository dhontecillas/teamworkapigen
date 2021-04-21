# FileversionFileversion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FileData** | Pointer to **string** |  | [optional] 
**FileRef** | Pointer to **string** |  | [optional] 
**FileSize** | Pointer to **int32** |  | [optional] 
**FileSource** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Notify** | Pointer to [**PayloadNotify**](PayloadNotify.md) |  | [optional] 
**NotifyCurrentUser** | Pointer to **bool** |  | [optional] 
**PendingFileRef** | Pointer to **string** |  | [optional] 
**Privacy** | Pointer to [**PayloadUserGroups**](PayloadUserGroups.md) |  | [optional] 
**Private** | Pointer to **bool** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 

## Methods

### NewFileversionFileversion

`func NewFileversionFileversion() *FileversionFileversion`

NewFileversionFileversion instantiates a new FileversionFileversion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileversionFileversionWithDefaults

`func NewFileversionFileversionWithDefaults() *FileversionFileversion`

NewFileversionFileversionWithDefaults instantiates a new FileversionFileversion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *FileversionFileversion) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *FileversionFileversion) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *FileversionFileversion) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *FileversionFileversion) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetDescription

`func (o *FileversionFileversion) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FileversionFileversion) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FileversionFileversion) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FileversionFileversion) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFileData

`func (o *FileversionFileversion) GetFileData() string`

GetFileData returns the FileData field if non-nil, zero value otherwise.

### GetFileDataOk

`func (o *FileversionFileversion) GetFileDataOk() (*string, bool)`

GetFileDataOk returns a tuple with the FileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileData

`func (o *FileversionFileversion) SetFileData(v string)`

SetFileData sets FileData field to given value.

### HasFileData

`func (o *FileversionFileversion) HasFileData() bool`

HasFileData returns a boolean if a field has been set.

### GetFileRef

`func (o *FileversionFileversion) GetFileRef() string`

GetFileRef returns the FileRef field if non-nil, zero value otherwise.

### GetFileRefOk

`func (o *FileversionFileversion) GetFileRefOk() (*string, bool)`

GetFileRefOk returns a tuple with the FileRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileRef

`func (o *FileversionFileversion) SetFileRef(v string)`

SetFileRef sets FileRef field to given value.

### HasFileRef

`func (o *FileversionFileversion) HasFileRef() bool`

HasFileRef returns a boolean if a field has been set.

### GetFileSize

`func (o *FileversionFileversion) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *FileversionFileversion) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *FileversionFileversion) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *FileversionFileversion) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetFileSource

`func (o *FileversionFileversion) GetFileSource() string`

GetFileSource returns the FileSource field if non-nil, zero value otherwise.

### GetFileSourceOk

`func (o *FileversionFileversion) GetFileSourceOk() (*string, bool)`

GetFileSourceOk returns a tuple with the FileSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSource

`func (o *FileversionFileversion) SetFileSource(v string)`

SetFileSource sets FileSource field to given value.

### HasFileSource

`func (o *FileversionFileversion) HasFileSource() bool`

HasFileSource returns a boolean if a field has been set.

### GetId

`func (o *FileversionFileversion) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileversionFileversion) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileversionFileversion) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FileversionFileversion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FileversionFileversion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileversionFileversion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileversionFileversion) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FileversionFileversion) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotify

`func (o *FileversionFileversion) GetNotify() PayloadNotify`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *FileversionFileversion) GetNotifyOk() (*PayloadNotify, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *FileversionFileversion) SetNotify(v PayloadNotify)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *FileversionFileversion) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetNotifyCurrentUser

`func (o *FileversionFileversion) GetNotifyCurrentUser() bool`

GetNotifyCurrentUser returns the NotifyCurrentUser field if non-nil, zero value otherwise.

### GetNotifyCurrentUserOk

`func (o *FileversionFileversion) GetNotifyCurrentUserOk() (*bool, bool)`

GetNotifyCurrentUserOk returns a tuple with the NotifyCurrentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyCurrentUser

`func (o *FileversionFileversion) SetNotifyCurrentUser(v bool)`

SetNotifyCurrentUser sets NotifyCurrentUser field to given value.

### HasNotifyCurrentUser

`func (o *FileversionFileversion) HasNotifyCurrentUser() bool`

HasNotifyCurrentUser returns a boolean if a field has been set.

### GetPendingFileRef

`func (o *FileversionFileversion) GetPendingFileRef() string`

GetPendingFileRef returns the PendingFileRef field if non-nil, zero value otherwise.

### GetPendingFileRefOk

`func (o *FileversionFileversion) GetPendingFileRefOk() (*string, bool)`

GetPendingFileRefOk returns a tuple with the PendingFileRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingFileRef

`func (o *FileversionFileversion) SetPendingFileRef(v string)`

SetPendingFileRef sets PendingFileRef field to given value.

### HasPendingFileRef

`func (o *FileversionFileversion) HasPendingFileRef() bool`

HasPendingFileRef returns a boolean if a field has been set.

### GetPrivacy

`func (o *FileversionFileversion) GetPrivacy() PayloadUserGroups`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *FileversionFileversion) GetPrivacyOk() (*PayloadUserGroups, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *FileversionFileversion) SetPrivacy(v PayloadUserGroups)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *FileversionFileversion) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetPrivate

`func (o *FileversionFileversion) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *FileversionFileversion) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *FileversionFileversion) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *FileversionFileversion) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetProjectId

`func (o *FileversionFileversion) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *FileversionFileversion) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *FileversionFileversion) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *FileversionFileversion) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


