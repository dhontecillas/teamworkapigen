# FileversionResponseIncluded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to [**map[string]ViewProjectFile**](ViewProjectFile.md) |  | [optional] 
**Projects** | Pointer to [**map[string]ViewProject**](ViewProject.md) |  | [optional] 
**Users** | Pointer to [**map[string]ViewUser**](ViewUser.md) |  | [optional] 

## Methods

### NewFileversionResponseIncluded

`func NewFileversionResponseIncluded() *FileversionResponseIncluded`

NewFileversionResponseIncluded instantiates a new FileversionResponseIncluded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileversionResponseIncludedWithDefaults

`func NewFileversionResponseIncludedWithDefaults() *FileversionResponseIncluded`

NewFileversionResponseIncludedWithDefaults instantiates a new FileversionResponseIncluded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *FileversionResponseIncluded) GetFiles() map[string]ViewProjectFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *FileversionResponseIncluded) GetFilesOk() (*map[string]ViewProjectFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *FileversionResponseIncluded) SetFiles(v map[string]ViewProjectFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *FileversionResponseIncluded) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetProjects

`func (o *FileversionResponseIncluded) GetProjects() map[string]ViewProject`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *FileversionResponseIncluded) GetProjectsOk() (*map[string]ViewProject, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *FileversionResponseIncluded) SetProjects(v map[string]ViewProject)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *FileversionResponseIncluded) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetUsers

`func (o *FileversionResponseIncluded) GetUsers() map[string]ViewUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *FileversionResponseIncluded) GetUsersOk() (*map[string]ViewUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *FileversionResponseIncluded) SetUsers(v map[string]ViewUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *FileversionResponseIncluded) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


