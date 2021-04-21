# CommentCommentsResponseIncluded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Companies** | Pointer to [**map[string]ViewCompany**](ViewCompany.md) |  | [optional] 
**Files** | Pointer to [**map[string]ViewProjectFile**](ViewProjectFile.md) |  | [optional] 
**Fileversions** | Pointer to [**map[string]ViewFileversion**](ViewFileversion.md) |  | [optional] 
**Links** | Pointer to [**map[string]ViewLinkItem**](ViewLinkItem.md) |  | [optional] 
**Milestones** | Pointer to [**map[string]ViewMilestone**](ViewMilestone.md) |  | [optional] 
**Notebooks** | Pointer to [**map[string]ViewNotebook**](ViewNotebook.md) |  | [optional] 
**Projects** | Pointer to [**map[string]ViewProject**](ViewProject.md) |  | [optional] 
**Tasks** | Pointer to [**map[string]ViewTask**](ViewTask.md) |  | [optional] 
**Users** | Pointer to [**map[string]ViewUser**](ViewUser.md) |  | [optional] 

## Methods

### NewCommentCommentsResponseIncluded

`func NewCommentCommentsResponseIncluded() *CommentCommentsResponseIncluded`

NewCommentCommentsResponseIncluded instantiates a new CommentCommentsResponseIncluded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentCommentsResponseIncludedWithDefaults

`func NewCommentCommentsResponseIncludedWithDefaults() *CommentCommentsResponseIncluded`

NewCommentCommentsResponseIncludedWithDefaults instantiates a new CommentCommentsResponseIncluded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanies

`func (o *CommentCommentsResponseIncluded) GetCompanies() map[string]ViewCompany`

GetCompanies returns the Companies field if non-nil, zero value otherwise.

### GetCompaniesOk

`func (o *CommentCommentsResponseIncluded) GetCompaniesOk() (*map[string]ViewCompany, bool)`

GetCompaniesOk returns a tuple with the Companies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanies

`func (o *CommentCommentsResponseIncluded) SetCompanies(v map[string]ViewCompany)`

SetCompanies sets Companies field to given value.

### HasCompanies

`func (o *CommentCommentsResponseIncluded) HasCompanies() bool`

HasCompanies returns a boolean if a field has been set.

### GetFiles

`func (o *CommentCommentsResponseIncluded) GetFiles() map[string]ViewProjectFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *CommentCommentsResponseIncluded) GetFilesOk() (*map[string]ViewProjectFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *CommentCommentsResponseIncluded) SetFiles(v map[string]ViewProjectFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *CommentCommentsResponseIncluded) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetFileversions

`func (o *CommentCommentsResponseIncluded) GetFileversions() map[string]ViewFileversion`

GetFileversions returns the Fileversions field if non-nil, zero value otherwise.

### GetFileversionsOk

`func (o *CommentCommentsResponseIncluded) GetFileversionsOk() (*map[string]ViewFileversion, bool)`

GetFileversionsOk returns a tuple with the Fileversions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileversions

`func (o *CommentCommentsResponseIncluded) SetFileversions(v map[string]ViewFileversion)`

SetFileversions sets Fileversions field to given value.

### HasFileversions

`func (o *CommentCommentsResponseIncluded) HasFileversions() bool`

HasFileversions returns a boolean if a field has been set.

### GetLinks

`func (o *CommentCommentsResponseIncluded) GetLinks() map[string]ViewLinkItem`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CommentCommentsResponseIncluded) GetLinksOk() (*map[string]ViewLinkItem, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CommentCommentsResponseIncluded) SetLinks(v map[string]ViewLinkItem)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CommentCommentsResponseIncluded) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMilestones

`func (o *CommentCommentsResponseIncluded) GetMilestones() map[string]ViewMilestone`

GetMilestones returns the Milestones field if non-nil, zero value otherwise.

### GetMilestonesOk

`func (o *CommentCommentsResponseIncluded) GetMilestonesOk() (*map[string]ViewMilestone, bool)`

GetMilestonesOk returns a tuple with the Milestones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestones

`func (o *CommentCommentsResponseIncluded) SetMilestones(v map[string]ViewMilestone)`

SetMilestones sets Milestones field to given value.

### HasMilestones

`func (o *CommentCommentsResponseIncluded) HasMilestones() bool`

HasMilestones returns a boolean if a field has been set.

### GetNotebooks

`func (o *CommentCommentsResponseIncluded) GetNotebooks() map[string]ViewNotebook`

GetNotebooks returns the Notebooks field if non-nil, zero value otherwise.

### GetNotebooksOk

`func (o *CommentCommentsResponseIncluded) GetNotebooksOk() (*map[string]ViewNotebook, bool)`

GetNotebooksOk returns a tuple with the Notebooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotebooks

`func (o *CommentCommentsResponseIncluded) SetNotebooks(v map[string]ViewNotebook)`

SetNotebooks sets Notebooks field to given value.

### HasNotebooks

`func (o *CommentCommentsResponseIncluded) HasNotebooks() bool`

HasNotebooks returns a boolean if a field has been set.

### GetProjects

`func (o *CommentCommentsResponseIncluded) GetProjects() map[string]ViewProject`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *CommentCommentsResponseIncluded) GetProjectsOk() (*map[string]ViewProject, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *CommentCommentsResponseIncluded) SetProjects(v map[string]ViewProject)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *CommentCommentsResponseIncluded) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetTasks

`func (o *CommentCommentsResponseIncluded) GetTasks() map[string]ViewTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *CommentCommentsResponseIncluded) GetTasksOk() (*map[string]ViewTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *CommentCommentsResponseIncluded) SetTasks(v map[string]ViewTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *CommentCommentsResponseIncluded) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetUsers

`func (o *CommentCommentsResponseIncluded) GetUsers() map[string]ViewUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *CommentCommentsResponseIncluded) GetUsersOk() (*map[string]ViewUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *CommentCommentsResponseIncluded) SetUsers(v map[string]ViewUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *CommentCommentsResponseIncluded) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


