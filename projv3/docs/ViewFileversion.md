# ViewFileversion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommentsCount** | Pointer to **int32** |  | [optional] 
**CommentsCountRead** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**File** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**FileId** | Pointer to **int32** |  | [optional] 
**FileVersionId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OriginalName** | Pointer to **string** |  | [optional] 
**Project** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**Reactions** | Pointer to [**ViewReactionsForObject**](ViewReactionsForObject.md) |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UploadedBy** | Pointer to **int32** |  | [optional] 
**VersionNo** | Pointer to **int32** |  | [optional] 

## Methods

### NewViewFileversion

`func NewViewFileversion() *ViewFileversion`

NewViewFileversion instantiates a new ViewFileversion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewFileversionWithDefaults

`func NewViewFileversionWithDefaults() *ViewFileversion`

NewViewFileversionWithDefaults instantiates a new ViewFileversion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommentsCount

`func (o *ViewFileversion) GetCommentsCount() int32`

GetCommentsCount returns the CommentsCount field if non-nil, zero value otherwise.

### GetCommentsCountOk

`func (o *ViewFileversion) GetCommentsCountOk() (*int32, bool)`

GetCommentsCountOk returns a tuple with the CommentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsCount

`func (o *ViewFileversion) SetCommentsCount(v int32)`

SetCommentsCount sets CommentsCount field to given value.

### HasCommentsCount

`func (o *ViewFileversion) HasCommentsCount() bool`

HasCommentsCount returns a boolean if a field has been set.

### GetCommentsCountRead

`func (o *ViewFileversion) GetCommentsCountRead() int32`

GetCommentsCountRead returns the CommentsCountRead field if non-nil, zero value otherwise.

### GetCommentsCountReadOk

`func (o *ViewFileversion) GetCommentsCountReadOk() (*int32, bool)`

GetCommentsCountReadOk returns a tuple with the CommentsCountRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsCountRead

`func (o *ViewFileversion) SetCommentsCountRead(v int32)`

SetCommentsCountRead sets CommentsCountRead field to given value.

### HasCommentsCountRead

`func (o *ViewFileversion) HasCommentsCountRead() bool`

HasCommentsCountRead returns a boolean if a field has been set.

### GetDescription

`func (o *ViewFileversion) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ViewFileversion) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ViewFileversion) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ViewFileversion) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *ViewFileversion) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ViewFileversion) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ViewFileversion) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ViewFileversion) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFile

`func (o *ViewFileversion) GetFile() ViewRelationship`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ViewFileversion) GetFileOk() (*ViewRelationship, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ViewFileversion) SetFile(v ViewRelationship)`

SetFile sets File field to given value.

### HasFile

`func (o *ViewFileversion) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetFileId

`func (o *ViewFileversion) GetFileId() int32`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *ViewFileversion) GetFileIdOk() (*int32, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *ViewFileversion) SetFileId(v int32)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *ViewFileversion) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetFileVersionId

`func (o *ViewFileversion) GetFileVersionId() int32`

GetFileVersionId returns the FileVersionId field if non-nil, zero value otherwise.

### GetFileVersionIdOk

`func (o *ViewFileversion) GetFileVersionIdOk() (*int32, bool)`

GetFileVersionIdOk returns a tuple with the FileVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileVersionId

`func (o *ViewFileversion) SetFileVersionId(v int32)`

SetFileVersionId sets FileVersionId field to given value.

### HasFileVersionId

`func (o *ViewFileversion) HasFileVersionId() bool`

HasFileVersionId returns a boolean if a field has been set.

### GetName

`func (o *ViewFileversion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ViewFileversion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ViewFileversion) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ViewFileversion) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOriginalName

`func (o *ViewFileversion) GetOriginalName() string`

GetOriginalName returns the OriginalName field if non-nil, zero value otherwise.

### GetOriginalNameOk

`func (o *ViewFileversion) GetOriginalNameOk() (*string, bool)`

GetOriginalNameOk returns a tuple with the OriginalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalName

`func (o *ViewFileversion) SetOriginalName(v string)`

SetOriginalName sets OriginalName field to given value.

### HasOriginalName

`func (o *ViewFileversion) HasOriginalName() bool`

HasOriginalName returns a boolean if a field has been set.

### GetProject

`func (o *ViewFileversion) GetProject() ViewRelationship`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ViewFileversion) GetProjectOk() (*ViewRelationship, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ViewFileversion) SetProject(v ViewRelationship)`

SetProject sets Project field to given value.

### HasProject

`func (o *ViewFileversion) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectId

`func (o *ViewFileversion) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ViewFileversion) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ViewFileversion) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ViewFileversion) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetReactions

`func (o *ViewFileversion) GetReactions() ViewReactionsForObject`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *ViewFileversion) GetReactionsOk() (*ViewReactionsForObject, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *ViewFileversion) SetReactions(v ViewReactionsForObject)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *ViewFileversion) HasReactions() bool`

HasReactions returns a boolean if a field has been set.

### GetSize

`func (o *ViewFileversion) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ViewFileversion) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ViewFileversion) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ViewFileversion) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStatus

`func (o *ViewFileversion) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ViewFileversion) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ViewFileversion) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ViewFileversion) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUploadedBy

`func (o *ViewFileversion) GetUploadedBy() int32`

GetUploadedBy returns the UploadedBy field if non-nil, zero value otherwise.

### GetUploadedByOk

`func (o *ViewFileversion) GetUploadedByOk() (*int32, bool)`

GetUploadedByOk returns a tuple with the UploadedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedBy

`func (o *ViewFileversion) SetUploadedBy(v int32)`

SetUploadedBy sets UploadedBy field to given value.

### HasUploadedBy

`func (o *ViewFileversion) HasUploadedBy() bool`

HasUploadedBy returns a boolean if a field has been set.

### GetVersionNo

`func (o *ViewFileversion) GetVersionNo() int32`

GetVersionNo returns the VersionNo field if non-nil, zero value otherwise.

### GetVersionNoOk

`func (o *ViewFileversion) GetVersionNoOk() (*int32, bool)`

GetVersionNoOk returns a tuple with the VersionNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionNo

`func (o *ViewFileversion) SetVersionNo(v int32)`

SetVersionNo sets VersionNo field to given value.

### HasVersionNo

`func (o *ViewFileversion) HasVersionNo() bool`

HasVersionNo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


