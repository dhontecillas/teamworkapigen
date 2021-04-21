# ViewUserGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Companies** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**CompanyIds** | Pointer to **[]int32** |  | [optional] 
**TeamIds** | Pointer to **[]int32** |  | [optional] 
**Teams** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**UserIds** | Pointer to **[]int32** |  | [optional] 
**Users** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 

## Methods

### NewViewUserGroups

`func NewViewUserGroups() *ViewUserGroups`

NewViewUserGroups instantiates a new ViewUserGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewUserGroupsWithDefaults

`func NewViewUserGroupsWithDefaults() *ViewUserGroups`

NewViewUserGroupsWithDefaults instantiates a new ViewUserGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanies

`func (o *ViewUserGroups) GetCompanies() []ViewRelationship`

GetCompanies returns the Companies field if non-nil, zero value otherwise.

### GetCompaniesOk

`func (o *ViewUserGroups) GetCompaniesOk() (*[]ViewRelationship, bool)`

GetCompaniesOk returns a tuple with the Companies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanies

`func (o *ViewUserGroups) SetCompanies(v []ViewRelationship)`

SetCompanies sets Companies field to given value.

### HasCompanies

`func (o *ViewUserGroups) HasCompanies() bool`

HasCompanies returns a boolean if a field has been set.

### GetCompanyIds

`func (o *ViewUserGroups) GetCompanyIds() []int32`

GetCompanyIds returns the CompanyIds field if non-nil, zero value otherwise.

### GetCompanyIdsOk

`func (o *ViewUserGroups) GetCompanyIdsOk() (*[]int32, bool)`

GetCompanyIdsOk returns a tuple with the CompanyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIds

`func (o *ViewUserGroups) SetCompanyIds(v []int32)`

SetCompanyIds sets CompanyIds field to given value.

### HasCompanyIds

`func (o *ViewUserGroups) HasCompanyIds() bool`

HasCompanyIds returns a boolean if a field has been set.

### GetTeamIds

`func (o *ViewUserGroups) GetTeamIds() []int32`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *ViewUserGroups) GetTeamIdsOk() (*[]int32, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *ViewUserGroups) SetTeamIds(v []int32)`

SetTeamIds sets TeamIds field to given value.

### HasTeamIds

`func (o *ViewUserGroups) HasTeamIds() bool`

HasTeamIds returns a boolean if a field has been set.

### GetTeams

`func (o *ViewUserGroups) GetTeams() []ViewRelationship`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *ViewUserGroups) GetTeamsOk() (*[]ViewRelationship, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *ViewUserGroups) SetTeams(v []ViewRelationship)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *ViewUserGroups) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetUserIds

`func (o *ViewUserGroups) GetUserIds() []int32`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *ViewUserGroups) GetUserIdsOk() (*[]int32, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *ViewUserGroups) SetUserIds(v []int32)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *ViewUserGroups) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### GetUsers

`func (o *ViewUserGroups) GetUsers() []ViewRelationship`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ViewUserGroups) GetUsersOk() (*[]ViewRelationship, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ViewUserGroups) SetUsers(v []ViewRelationship)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *ViewUserGroups) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


