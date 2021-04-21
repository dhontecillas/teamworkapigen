# LockdownResponseIncluded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comments** | Pointer to [**map[string]ViewComment**](ViewComment.md) |  | [optional] 
**Companies** | Pointer to [**map[string]ViewCompany**](ViewCompany.md) |  | [optional] 
**Files** | Pointer to [**map[string]ViewProjectFile**](ViewProjectFile.md) |  | [optional] 
**Links** | Pointer to [**map[string]ViewLinkItem**](ViewLinkItem.md) |  | [optional] 
**Messages** | Pointer to [**map[string]ViewProjectMessage**](ViewProjectMessage.md) |  | [optional] 
**Milestones** | Pointer to [**map[string]ViewMilestone**](ViewMilestone.md) |  | [optional] 
**Notebooks** | Pointer to [**map[string]ViewNotebook**](ViewNotebook.md) |  | [optional] 
**Tasklists** | Pointer to [**map[string]ViewTasklist**](ViewTasklist.md) |  | [optional] 
**Tasks** | Pointer to [**map[string]ViewTask**](ViewTask.md) |  | [optional] 
**Teams** | Pointer to [**map[string]ViewTeam**](ViewTeam.md) |  | [optional] 
**Users** | Pointer to [**map[string]ViewUser**](ViewUser.md) |  | [optional] 

## Methods

### NewLockdownResponseIncluded

`func NewLockdownResponseIncluded() *LockdownResponseIncluded`

NewLockdownResponseIncluded instantiates a new LockdownResponseIncluded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLockdownResponseIncludedWithDefaults

`func NewLockdownResponseIncludedWithDefaults() *LockdownResponseIncluded`

NewLockdownResponseIncludedWithDefaults instantiates a new LockdownResponseIncluded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComments

`func (o *LockdownResponseIncluded) GetComments() map[string]ViewComment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *LockdownResponseIncluded) GetCommentsOk() (*map[string]ViewComment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *LockdownResponseIncluded) SetComments(v map[string]ViewComment)`

SetComments sets Comments field to given value.

### HasComments

`func (o *LockdownResponseIncluded) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCompanies

`func (o *LockdownResponseIncluded) GetCompanies() map[string]ViewCompany`

GetCompanies returns the Companies field if non-nil, zero value otherwise.

### GetCompaniesOk

`func (o *LockdownResponseIncluded) GetCompaniesOk() (*map[string]ViewCompany, bool)`

GetCompaniesOk returns a tuple with the Companies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanies

`func (o *LockdownResponseIncluded) SetCompanies(v map[string]ViewCompany)`

SetCompanies sets Companies field to given value.

### HasCompanies

`func (o *LockdownResponseIncluded) HasCompanies() bool`

HasCompanies returns a boolean if a field has been set.

### GetFiles

`func (o *LockdownResponseIncluded) GetFiles() map[string]ViewProjectFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *LockdownResponseIncluded) GetFilesOk() (*map[string]ViewProjectFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *LockdownResponseIncluded) SetFiles(v map[string]ViewProjectFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *LockdownResponseIncluded) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetLinks

`func (o *LockdownResponseIncluded) GetLinks() map[string]ViewLinkItem`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LockdownResponseIncluded) GetLinksOk() (*map[string]ViewLinkItem, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LockdownResponseIncluded) SetLinks(v map[string]ViewLinkItem)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *LockdownResponseIncluded) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMessages

`func (o *LockdownResponseIncluded) GetMessages() map[string]ViewProjectMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *LockdownResponseIncluded) GetMessagesOk() (*map[string]ViewProjectMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *LockdownResponseIncluded) SetMessages(v map[string]ViewProjectMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *LockdownResponseIncluded) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetMilestones

`func (o *LockdownResponseIncluded) GetMilestones() map[string]ViewMilestone`

GetMilestones returns the Milestones field if non-nil, zero value otherwise.

### GetMilestonesOk

`func (o *LockdownResponseIncluded) GetMilestonesOk() (*map[string]ViewMilestone, bool)`

GetMilestonesOk returns a tuple with the Milestones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestones

`func (o *LockdownResponseIncluded) SetMilestones(v map[string]ViewMilestone)`

SetMilestones sets Milestones field to given value.

### HasMilestones

`func (o *LockdownResponseIncluded) HasMilestones() bool`

HasMilestones returns a boolean if a field has been set.

### GetNotebooks

`func (o *LockdownResponseIncluded) GetNotebooks() map[string]ViewNotebook`

GetNotebooks returns the Notebooks field if non-nil, zero value otherwise.

### GetNotebooksOk

`func (o *LockdownResponseIncluded) GetNotebooksOk() (*map[string]ViewNotebook, bool)`

GetNotebooksOk returns a tuple with the Notebooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotebooks

`func (o *LockdownResponseIncluded) SetNotebooks(v map[string]ViewNotebook)`

SetNotebooks sets Notebooks field to given value.

### HasNotebooks

`func (o *LockdownResponseIncluded) HasNotebooks() bool`

HasNotebooks returns a boolean if a field has been set.

### GetTasklists

`func (o *LockdownResponseIncluded) GetTasklists() map[string]ViewTasklist`

GetTasklists returns the Tasklists field if non-nil, zero value otherwise.

### GetTasklistsOk

`func (o *LockdownResponseIncluded) GetTasklistsOk() (*map[string]ViewTasklist, bool)`

GetTasklistsOk returns a tuple with the Tasklists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasklists

`func (o *LockdownResponseIncluded) SetTasklists(v map[string]ViewTasklist)`

SetTasklists sets Tasklists field to given value.

### HasTasklists

`func (o *LockdownResponseIncluded) HasTasklists() bool`

HasTasklists returns a boolean if a field has been set.

### GetTasks

`func (o *LockdownResponseIncluded) GetTasks() map[string]ViewTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *LockdownResponseIncluded) GetTasksOk() (*map[string]ViewTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *LockdownResponseIncluded) SetTasks(v map[string]ViewTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *LockdownResponseIncluded) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetTeams

`func (o *LockdownResponseIncluded) GetTeams() map[string]ViewTeam`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *LockdownResponseIncluded) GetTeamsOk() (*map[string]ViewTeam, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *LockdownResponseIncluded) SetTeams(v map[string]ViewTeam)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *LockdownResponseIncluded) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetUsers

`func (o *LockdownResponseIncluded) GetUsers() map[string]ViewUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *LockdownResponseIncluded) GetUsersOk() (*map[string]ViewUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *LockdownResponseIncluded) SetUsers(v map[string]ViewUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *LockdownResponseIncluded) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


