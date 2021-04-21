# ActivityActivitiesResponseIncluded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Companies** | Pointer to [**map[string]ViewCompany**](ViewCompany.md) |  | [optional] 
**Projects** | Pointer to [**map[string]ViewProject**](ViewProject.md) |  | [optional] 
**Users** | Pointer to [**map[string]ViewUser**](ViewUser.md) |  | [optional] 

## Methods

### NewActivityActivitiesResponseIncluded

`func NewActivityActivitiesResponseIncluded() *ActivityActivitiesResponseIncluded`

NewActivityActivitiesResponseIncluded instantiates a new ActivityActivitiesResponseIncluded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityActivitiesResponseIncludedWithDefaults

`func NewActivityActivitiesResponseIncludedWithDefaults() *ActivityActivitiesResponseIncluded`

NewActivityActivitiesResponseIncludedWithDefaults instantiates a new ActivityActivitiesResponseIncluded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanies

`func (o *ActivityActivitiesResponseIncluded) GetCompanies() map[string]ViewCompany`

GetCompanies returns the Companies field if non-nil, zero value otherwise.

### GetCompaniesOk

`func (o *ActivityActivitiesResponseIncluded) GetCompaniesOk() (*map[string]ViewCompany, bool)`

GetCompaniesOk returns a tuple with the Companies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanies

`func (o *ActivityActivitiesResponseIncluded) SetCompanies(v map[string]ViewCompany)`

SetCompanies sets Companies field to given value.

### HasCompanies

`func (o *ActivityActivitiesResponseIncluded) HasCompanies() bool`

HasCompanies returns a boolean if a field has been set.

### GetProjects

`func (o *ActivityActivitiesResponseIncluded) GetProjects() map[string]ViewProject`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *ActivityActivitiesResponseIncluded) GetProjectsOk() (*map[string]ViewProject, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *ActivityActivitiesResponseIncluded) SetProjects(v map[string]ViewProject)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *ActivityActivitiesResponseIncluded) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetUsers

`func (o *ActivityActivitiesResponseIncluded) GetUsers() map[string]ViewUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ActivityActivitiesResponseIncluded) GetUsersOk() (*map[string]ViewUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ActivityActivitiesResponseIncluded) SetUsers(v map[string]ViewUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *ActivityActivitiesResponseIncluded) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


