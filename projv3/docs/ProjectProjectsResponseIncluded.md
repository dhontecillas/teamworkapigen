# ProjectProjectsResponseIncluded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Companies** | Pointer to [**map[string]ViewCompany**](ViewCompany.md) |  | [optional] 
**CustomfieldProjects** | Pointer to [**map[string]ViewCustomFieldProject**](ViewCustomFieldProject.md) |  | [optional] 
**Customfields** | Pointer to [**map[string]ViewCustomField**](ViewCustomField.md) |  | [optional] 
**PortfolioBoards** | Pointer to [**map[string]ViewPortfolioBoard**](ViewPortfolioBoard.md) |  | [optional] 
**PortfolioCards** | Pointer to [**map[string]ViewPortfolioCard**](ViewPortfolioCard.md) |  | [optional] 
**PortfolioColumns** | Pointer to [**map[string]ViewPortfolioColumn**](ViewPortfolioColumn.md) |  | [optional] 
**ProjectBudgets** | Pointer to [**map[string]ViewProjectBudget**](ViewProjectBudget.md) |  | [optional] 
**ProjectCategories** | Pointer to [**map[string]ViewProjectCategory**](ViewProjectCategory.md) |  | [optional] 
**ProjectTaskStats** | Pointer to [**map[string]ViewProjectTaskStats**](ViewProjectTaskStats.md) |  | [optional] 
**ProjectUpdates** | Pointer to [**map[string]ViewProjectUpdate**](ViewProjectUpdate.md) |  | [optional] 
**Tags** | Pointer to [**map[string]ViewTag**](ViewTag.md) |  | [optional] 
**Users** | Pointer to [**map[string]ViewUser**](ViewUser.md) |  | [optional] 

## Methods

### NewProjectProjectsResponseIncluded

`func NewProjectProjectsResponseIncluded() *ProjectProjectsResponseIncluded`

NewProjectProjectsResponseIncluded instantiates a new ProjectProjectsResponseIncluded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectProjectsResponseIncludedWithDefaults

`func NewProjectProjectsResponseIncludedWithDefaults() *ProjectProjectsResponseIncluded`

NewProjectProjectsResponseIncludedWithDefaults instantiates a new ProjectProjectsResponseIncluded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanies

`func (o *ProjectProjectsResponseIncluded) GetCompanies() map[string]ViewCompany`

GetCompanies returns the Companies field if non-nil, zero value otherwise.

### GetCompaniesOk

`func (o *ProjectProjectsResponseIncluded) GetCompaniesOk() (*map[string]ViewCompany, bool)`

GetCompaniesOk returns a tuple with the Companies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanies

`func (o *ProjectProjectsResponseIncluded) SetCompanies(v map[string]ViewCompany)`

SetCompanies sets Companies field to given value.

### HasCompanies

`func (o *ProjectProjectsResponseIncluded) HasCompanies() bool`

HasCompanies returns a boolean if a field has been set.

### GetCustomfieldProjects

`func (o *ProjectProjectsResponseIncluded) GetCustomfieldProjects() map[string]ViewCustomFieldProject`

GetCustomfieldProjects returns the CustomfieldProjects field if non-nil, zero value otherwise.

### GetCustomfieldProjectsOk

`func (o *ProjectProjectsResponseIncluded) GetCustomfieldProjectsOk() (*map[string]ViewCustomFieldProject, bool)`

GetCustomfieldProjectsOk returns a tuple with the CustomfieldProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomfieldProjects

`func (o *ProjectProjectsResponseIncluded) SetCustomfieldProjects(v map[string]ViewCustomFieldProject)`

SetCustomfieldProjects sets CustomfieldProjects field to given value.

### HasCustomfieldProjects

`func (o *ProjectProjectsResponseIncluded) HasCustomfieldProjects() bool`

HasCustomfieldProjects returns a boolean if a field has been set.

### GetCustomfields

`func (o *ProjectProjectsResponseIncluded) GetCustomfields() map[string]ViewCustomField`

GetCustomfields returns the Customfields field if non-nil, zero value otherwise.

### GetCustomfieldsOk

`func (o *ProjectProjectsResponseIncluded) GetCustomfieldsOk() (*map[string]ViewCustomField, bool)`

GetCustomfieldsOk returns a tuple with the Customfields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomfields

`func (o *ProjectProjectsResponseIncluded) SetCustomfields(v map[string]ViewCustomField)`

SetCustomfields sets Customfields field to given value.

### HasCustomfields

`func (o *ProjectProjectsResponseIncluded) HasCustomfields() bool`

HasCustomfields returns a boolean if a field has been set.

### GetPortfolioBoards

`func (o *ProjectProjectsResponseIncluded) GetPortfolioBoards() map[string]ViewPortfolioBoard`

GetPortfolioBoards returns the PortfolioBoards field if non-nil, zero value otherwise.

### GetPortfolioBoardsOk

`func (o *ProjectProjectsResponseIncluded) GetPortfolioBoardsOk() (*map[string]ViewPortfolioBoard, bool)`

GetPortfolioBoardsOk returns a tuple with the PortfolioBoards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortfolioBoards

`func (o *ProjectProjectsResponseIncluded) SetPortfolioBoards(v map[string]ViewPortfolioBoard)`

SetPortfolioBoards sets PortfolioBoards field to given value.

### HasPortfolioBoards

`func (o *ProjectProjectsResponseIncluded) HasPortfolioBoards() bool`

HasPortfolioBoards returns a boolean if a field has been set.

### GetPortfolioCards

`func (o *ProjectProjectsResponseIncluded) GetPortfolioCards() map[string]ViewPortfolioCard`

GetPortfolioCards returns the PortfolioCards field if non-nil, zero value otherwise.

### GetPortfolioCardsOk

`func (o *ProjectProjectsResponseIncluded) GetPortfolioCardsOk() (*map[string]ViewPortfolioCard, bool)`

GetPortfolioCardsOk returns a tuple with the PortfolioCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortfolioCards

`func (o *ProjectProjectsResponseIncluded) SetPortfolioCards(v map[string]ViewPortfolioCard)`

SetPortfolioCards sets PortfolioCards field to given value.

### HasPortfolioCards

`func (o *ProjectProjectsResponseIncluded) HasPortfolioCards() bool`

HasPortfolioCards returns a boolean if a field has been set.

### GetPortfolioColumns

`func (o *ProjectProjectsResponseIncluded) GetPortfolioColumns() map[string]ViewPortfolioColumn`

GetPortfolioColumns returns the PortfolioColumns field if non-nil, zero value otherwise.

### GetPortfolioColumnsOk

`func (o *ProjectProjectsResponseIncluded) GetPortfolioColumnsOk() (*map[string]ViewPortfolioColumn, bool)`

GetPortfolioColumnsOk returns a tuple with the PortfolioColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortfolioColumns

`func (o *ProjectProjectsResponseIncluded) SetPortfolioColumns(v map[string]ViewPortfolioColumn)`

SetPortfolioColumns sets PortfolioColumns field to given value.

### HasPortfolioColumns

`func (o *ProjectProjectsResponseIncluded) HasPortfolioColumns() bool`

HasPortfolioColumns returns a boolean if a field has been set.

### GetProjectBudgets

`func (o *ProjectProjectsResponseIncluded) GetProjectBudgets() map[string]ViewProjectBudget`

GetProjectBudgets returns the ProjectBudgets field if non-nil, zero value otherwise.

### GetProjectBudgetsOk

`func (o *ProjectProjectsResponseIncluded) GetProjectBudgetsOk() (*map[string]ViewProjectBudget, bool)`

GetProjectBudgetsOk returns a tuple with the ProjectBudgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectBudgets

`func (o *ProjectProjectsResponseIncluded) SetProjectBudgets(v map[string]ViewProjectBudget)`

SetProjectBudgets sets ProjectBudgets field to given value.

### HasProjectBudgets

`func (o *ProjectProjectsResponseIncluded) HasProjectBudgets() bool`

HasProjectBudgets returns a boolean if a field has been set.

### GetProjectCategories

`func (o *ProjectProjectsResponseIncluded) GetProjectCategories() map[string]ViewProjectCategory`

GetProjectCategories returns the ProjectCategories field if non-nil, zero value otherwise.

### GetProjectCategoriesOk

`func (o *ProjectProjectsResponseIncluded) GetProjectCategoriesOk() (*map[string]ViewProjectCategory, bool)`

GetProjectCategoriesOk returns a tuple with the ProjectCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCategories

`func (o *ProjectProjectsResponseIncluded) SetProjectCategories(v map[string]ViewProjectCategory)`

SetProjectCategories sets ProjectCategories field to given value.

### HasProjectCategories

`func (o *ProjectProjectsResponseIncluded) HasProjectCategories() bool`

HasProjectCategories returns a boolean if a field has been set.

### GetProjectTaskStats

`func (o *ProjectProjectsResponseIncluded) GetProjectTaskStats() map[string]ViewProjectTaskStats`

GetProjectTaskStats returns the ProjectTaskStats field if non-nil, zero value otherwise.

### GetProjectTaskStatsOk

`func (o *ProjectProjectsResponseIncluded) GetProjectTaskStatsOk() (*map[string]ViewProjectTaskStats, bool)`

GetProjectTaskStatsOk returns a tuple with the ProjectTaskStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTaskStats

`func (o *ProjectProjectsResponseIncluded) SetProjectTaskStats(v map[string]ViewProjectTaskStats)`

SetProjectTaskStats sets ProjectTaskStats field to given value.

### HasProjectTaskStats

`func (o *ProjectProjectsResponseIncluded) HasProjectTaskStats() bool`

HasProjectTaskStats returns a boolean if a field has been set.

### GetProjectUpdates

`func (o *ProjectProjectsResponseIncluded) GetProjectUpdates() map[string]ViewProjectUpdate`

GetProjectUpdates returns the ProjectUpdates field if non-nil, zero value otherwise.

### GetProjectUpdatesOk

`func (o *ProjectProjectsResponseIncluded) GetProjectUpdatesOk() (*map[string]ViewProjectUpdate, bool)`

GetProjectUpdatesOk returns a tuple with the ProjectUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectUpdates

`func (o *ProjectProjectsResponseIncluded) SetProjectUpdates(v map[string]ViewProjectUpdate)`

SetProjectUpdates sets ProjectUpdates field to given value.

### HasProjectUpdates

`func (o *ProjectProjectsResponseIncluded) HasProjectUpdates() bool`

HasProjectUpdates returns a boolean if a field has been set.

### GetTags

`func (o *ProjectProjectsResponseIncluded) GetTags() map[string]ViewTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProjectProjectsResponseIncluded) GetTagsOk() (*map[string]ViewTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProjectProjectsResponseIncluded) SetTags(v map[string]ViewTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProjectProjectsResponseIncluded) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUsers

`func (o *ProjectProjectsResponseIncluded) GetUsers() map[string]ViewUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ProjectProjectsResponseIncluded) GetUsersOk() (*map[string]ViewUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ProjectProjectsResponseIncluded) SetUsers(v map[string]ViewUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *ProjectProjectsResponseIncluded) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


