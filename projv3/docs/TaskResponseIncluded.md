# TaskResponseIncluded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cards** | Pointer to [**map[string]ViewTaskCard**](ViewTaskCard.md) |  | [optional] 
**Columns** | Pointer to [**map[string]ViewBoardColumn**](ViewBoardColumn.md) |  | [optional] 
**Comments** | Pointer to [**map[string]ViewComment**](ViewComment.md) |  | [optional] 
**Companies** | Pointer to [**map[string]ViewCompany**](ViewCompany.md) |  | [optional] 
**Milestones** | Pointer to [**map[string]ViewMilestone**](ViewMilestone.md) |  | [optional] 
**Projects** | Pointer to [**map[string]ViewProject**](ViewProject.md) |  | [optional] 
**Tags** | Pointer to [**map[string]ViewTag**](ViewTag.md) |  | [optional] 
**Tasklists** | Pointer to [**map[string]ViewTasklist**](ViewTasklist.md) |  | [optional] 
**Teams** | Pointer to [**map[string]ViewTeam**](ViewTeam.md) |  | [optional] 
**TimeTotals** | Pointer to [**map[string]ViewTaskTimeTotals**](ViewTaskTimeTotals.md) |  | [optional] 
**Users** | Pointer to [**map[string]ViewUser**](ViewUser.md) |  | [optional] 

## Methods

### NewTaskResponseIncluded

`func NewTaskResponseIncluded() *TaskResponseIncluded`

NewTaskResponseIncluded instantiates a new TaskResponseIncluded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskResponseIncludedWithDefaults

`func NewTaskResponseIncludedWithDefaults() *TaskResponseIncluded`

NewTaskResponseIncludedWithDefaults instantiates a new TaskResponseIncluded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCards

`func (o *TaskResponseIncluded) GetCards() map[string]ViewTaskCard`

GetCards returns the Cards field if non-nil, zero value otherwise.

### GetCardsOk

`func (o *TaskResponseIncluded) GetCardsOk() (*map[string]ViewTaskCard, bool)`

GetCardsOk returns a tuple with the Cards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCards

`func (o *TaskResponseIncluded) SetCards(v map[string]ViewTaskCard)`

SetCards sets Cards field to given value.

### HasCards

`func (o *TaskResponseIncluded) HasCards() bool`

HasCards returns a boolean if a field has been set.

### GetColumns

`func (o *TaskResponseIncluded) GetColumns() map[string]ViewBoardColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *TaskResponseIncluded) GetColumnsOk() (*map[string]ViewBoardColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *TaskResponseIncluded) SetColumns(v map[string]ViewBoardColumn)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *TaskResponseIncluded) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetComments

`func (o *TaskResponseIncluded) GetComments() map[string]ViewComment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *TaskResponseIncluded) GetCommentsOk() (*map[string]ViewComment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *TaskResponseIncluded) SetComments(v map[string]ViewComment)`

SetComments sets Comments field to given value.

### HasComments

`func (o *TaskResponseIncluded) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCompanies

`func (o *TaskResponseIncluded) GetCompanies() map[string]ViewCompany`

GetCompanies returns the Companies field if non-nil, zero value otherwise.

### GetCompaniesOk

`func (o *TaskResponseIncluded) GetCompaniesOk() (*map[string]ViewCompany, bool)`

GetCompaniesOk returns a tuple with the Companies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanies

`func (o *TaskResponseIncluded) SetCompanies(v map[string]ViewCompany)`

SetCompanies sets Companies field to given value.

### HasCompanies

`func (o *TaskResponseIncluded) HasCompanies() bool`

HasCompanies returns a boolean if a field has been set.

### GetMilestones

`func (o *TaskResponseIncluded) GetMilestones() map[string]ViewMilestone`

GetMilestones returns the Milestones field if non-nil, zero value otherwise.

### GetMilestonesOk

`func (o *TaskResponseIncluded) GetMilestonesOk() (*map[string]ViewMilestone, bool)`

GetMilestonesOk returns a tuple with the Milestones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestones

`func (o *TaskResponseIncluded) SetMilestones(v map[string]ViewMilestone)`

SetMilestones sets Milestones field to given value.

### HasMilestones

`func (o *TaskResponseIncluded) HasMilestones() bool`

HasMilestones returns a boolean if a field has been set.

### GetProjects

`func (o *TaskResponseIncluded) GetProjects() map[string]ViewProject`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *TaskResponseIncluded) GetProjectsOk() (*map[string]ViewProject, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *TaskResponseIncluded) SetProjects(v map[string]ViewProject)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *TaskResponseIncluded) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetTags

`func (o *TaskResponseIncluded) GetTags() map[string]ViewTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TaskResponseIncluded) GetTagsOk() (*map[string]ViewTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TaskResponseIncluded) SetTags(v map[string]ViewTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TaskResponseIncluded) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTasklists

`func (o *TaskResponseIncluded) GetTasklists() map[string]ViewTasklist`

GetTasklists returns the Tasklists field if non-nil, zero value otherwise.

### GetTasklistsOk

`func (o *TaskResponseIncluded) GetTasklistsOk() (*map[string]ViewTasklist, bool)`

GetTasklistsOk returns a tuple with the Tasklists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasklists

`func (o *TaskResponseIncluded) SetTasklists(v map[string]ViewTasklist)`

SetTasklists sets Tasklists field to given value.

### HasTasklists

`func (o *TaskResponseIncluded) HasTasklists() bool`

HasTasklists returns a boolean if a field has been set.

### GetTeams

`func (o *TaskResponseIncluded) GetTeams() map[string]ViewTeam`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *TaskResponseIncluded) GetTeamsOk() (*map[string]ViewTeam, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *TaskResponseIncluded) SetTeams(v map[string]ViewTeam)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *TaskResponseIncluded) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetTimeTotals

`func (o *TaskResponseIncluded) GetTimeTotals() map[string]ViewTaskTimeTotals`

GetTimeTotals returns the TimeTotals field if non-nil, zero value otherwise.

### GetTimeTotalsOk

`func (o *TaskResponseIncluded) GetTimeTotalsOk() (*map[string]ViewTaskTimeTotals, bool)`

GetTimeTotalsOk returns a tuple with the TimeTotals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeTotals

`func (o *TaskResponseIncluded) SetTimeTotals(v map[string]ViewTaskTimeTotals)`

SetTimeTotals sets TimeTotals field to given value.

### HasTimeTotals

`func (o *TaskResponseIncluded) HasTimeTotals() bool`

HasTimeTotals returns a boolean if a field has been set.

### GetUsers

`func (o *TaskResponseIncluded) GetUsers() map[string]ViewUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *TaskResponseIncluded) GetUsersOk() (*map[string]ViewUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *TaskResponseIncluded) SetUsers(v map[string]ViewUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *TaskResponseIncluded) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


