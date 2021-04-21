# FileProjectFilesResponseIncluded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comments** | Pointer to [**map[string]ViewComment**](ViewComment.md) |  | [optional] 
**FileCategories** | Pointer to [**map[string]ViewProjectFileCategory**](ViewProjectFileCategory.md) |  | [optional] 
**Messages** | Pointer to [**map[string]ViewProjectMessage**](ViewProjectMessage.md) |  | [optional] 
**Permissions** | Pointer to [**map[string]ViewFilePermissions**](ViewFilePermissions.md) |  | [optional] 
**Projects** | Pointer to [**map[string]ViewProject**](ViewProject.md) |  | [optional] 
**Tags** | Pointer to [**map[string]ViewTag**](ViewTag.md) |  | [optional] 
**Tasks** | Pointer to [**map[string]ViewTask**](ViewTask.md) |  | [optional] 
**Users** | Pointer to [**map[string]ViewUser**](ViewUser.md) |  | [optional] 

## Methods

### NewFileProjectFilesResponseIncluded

`func NewFileProjectFilesResponseIncluded() *FileProjectFilesResponseIncluded`

NewFileProjectFilesResponseIncluded instantiates a new FileProjectFilesResponseIncluded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileProjectFilesResponseIncludedWithDefaults

`func NewFileProjectFilesResponseIncludedWithDefaults() *FileProjectFilesResponseIncluded`

NewFileProjectFilesResponseIncludedWithDefaults instantiates a new FileProjectFilesResponseIncluded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComments

`func (o *FileProjectFilesResponseIncluded) GetComments() map[string]ViewComment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *FileProjectFilesResponseIncluded) GetCommentsOk() (*map[string]ViewComment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *FileProjectFilesResponseIncluded) SetComments(v map[string]ViewComment)`

SetComments sets Comments field to given value.

### HasComments

`func (o *FileProjectFilesResponseIncluded) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetFileCategories

`func (o *FileProjectFilesResponseIncluded) GetFileCategories() map[string]ViewProjectFileCategory`

GetFileCategories returns the FileCategories field if non-nil, zero value otherwise.

### GetFileCategoriesOk

`func (o *FileProjectFilesResponseIncluded) GetFileCategoriesOk() (*map[string]ViewProjectFileCategory, bool)`

GetFileCategoriesOk returns a tuple with the FileCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCategories

`func (o *FileProjectFilesResponseIncluded) SetFileCategories(v map[string]ViewProjectFileCategory)`

SetFileCategories sets FileCategories field to given value.

### HasFileCategories

`func (o *FileProjectFilesResponseIncluded) HasFileCategories() bool`

HasFileCategories returns a boolean if a field has been set.

### GetMessages

`func (o *FileProjectFilesResponseIncluded) GetMessages() map[string]ViewProjectMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *FileProjectFilesResponseIncluded) GetMessagesOk() (*map[string]ViewProjectMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *FileProjectFilesResponseIncluded) SetMessages(v map[string]ViewProjectMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *FileProjectFilesResponseIncluded) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetPermissions

`func (o *FileProjectFilesResponseIncluded) GetPermissions() map[string]ViewFilePermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *FileProjectFilesResponseIncluded) GetPermissionsOk() (*map[string]ViewFilePermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *FileProjectFilesResponseIncluded) SetPermissions(v map[string]ViewFilePermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *FileProjectFilesResponseIncluded) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetProjects

`func (o *FileProjectFilesResponseIncluded) GetProjects() map[string]ViewProject`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *FileProjectFilesResponseIncluded) GetProjectsOk() (*map[string]ViewProject, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *FileProjectFilesResponseIncluded) SetProjects(v map[string]ViewProject)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *FileProjectFilesResponseIncluded) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetTags

`func (o *FileProjectFilesResponseIncluded) GetTags() map[string]ViewTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FileProjectFilesResponseIncluded) GetTagsOk() (*map[string]ViewTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FileProjectFilesResponseIncluded) SetTags(v map[string]ViewTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FileProjectFilesResponseIncluded) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTasks

`func (o *FileProjectFilesResponseIncluded) GetTasks() map[string]ViewTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *FileProjectFilesResponseIncluded) GetTasksOk() (*map[string]ViewTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *FileProjectFilesResponseIncluded) SetTasks(v map[string]ViewTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *FileProjectFilesResponseIncluded) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetUsers

`func (o *FileProjectFilesResponseIncluded) GetUsers() map[string]ViewUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *FileProjectFilesResponseIncluded) GetUsersOk() (*map[string]ViewUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *FileProjectFilesResponseIncluded) SetUsers(v map[string]ViewUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *FileProjectFilesResponseIncluded) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


