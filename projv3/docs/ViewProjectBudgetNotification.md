# ViewProjectBudgetNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapacityThreshold** | Pointer to **float32** |  | [optional] 
**Companies** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**CompanyIds** | Pointer to **[]int32** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**NotificationMedium** | Pointer to **string** |  | [optional] 
**TeamIds** | Pointer to **[]int32** |  | [optional] 
**Teams** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**UserIds** | Pointer to **[]int32** |  | [optional] 
**Users** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 

## Methods

### NewViewProjectBudgetNotification

`func NewViewProjectBudgetNotification() *ViewProjectBudgetNotification`

NewViewProjectBudgetNotification instantiates a new ViewProjectBudgetNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewProjectBudgetNotificationWithDefaults

`func NewViewProjectBudgetNotificationWithDefaults() *ViewProjectBudgetNotification`

NewViewProjectBudgetNotificationWithDefaults instantiates a new ViewProjectBudgetNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacityThreshold

`func (o *ViewProjectBudgetNotification) GetCapacityThreshold() float32`

GetCapacityThreshold returns the CapacityThreshold field if non-nil, zero value otherwise.

### GetCapacityThresholdOk

`func (o *ViewProjectBudgetNotification) GetCapacityThresholdOk() (*float32, bool)`

GetCapacityThresholdOk returns a tuple with the CapacityThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityThreshold

`func (o *ViewProjectBudgetNotification) SetCapacityThreshold(v float32)`

SetCapacityThreshold sets CapacityThreshold field to given value.

### HasCapacityThreshold

`func (o *ViewProjectBudgetNotification) HasCapacityThreshold() bool`

HasCapacityThreshold returns a boolean if a field has been set.

### GetCompanies

`func (o *ViewProjectBudgetNotification) GetCompanies() []ViewRelationship`

GetCompanies returns the Companies field if non-nil, zero value otherwise.

### GetCompaniesOk

`func (o *ViewProjectBudgetNotification) GetCompaniesOk() (*[]ViewRelationship, bool)`

GetCompaniesOk returns a tuple with the Companies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanies

`func (o *ViewProjectBudgetNotification) SetCompanies(v []ViewRelationship)`

SetCompanies sets Companies field to given value.

### HasCompanies

`func (o *ViewProjectBudgetNotification) HasCompanies() bool`

HasCompanies returns a boolean if a field has been set.

### GetCompanyIds

`func (o *ViewProjectBudgetNotification) GetCompanyIds() []int32`

GetCompanyIds returns the CompanyIds field if non-nil, zero value otherwise.

### GetCompanyIdsOk

`func (o *ViewProjectBudgetNotification) GetCompanyIdsOk() (*[]int32, bool)`

GetCompanyIdsOk returns a tuple with the CompanyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIds

`func (o *ViewProjectBudgetNotification) SetCompanyIds(v []int32)`

SetCompanyIds sets CompanyIds field to given value.

### HasCompanyIds

`func (o *ViewProjectBudgetNotification) HasCompanyIds() bool`

HasCompanyIds returns a boolean if a field has been set.

### GetId

`func (o *ViewProjectBudgetNotification) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewProjectBudgetNotification) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewProjectBudgetNotification) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewProjectBudgetNotification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNotificationMedium

`func (o *ViewProjectBudgetNotification) GetNotificationMedium() string`

GetNotificationMedium returns the NotificationMedium field if non-nil, zero value otherwise.

### GetNotificationMediumOk

`func (o *ViewProjectBudgetNotification) GetNotificationMediumOk() (*string, bool)`

GetNotificationMediumOk returns a tuple with the NotificationMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationMedium

`func (o *ViewProjectBudgetNotification) SetNotificationMedium(v string)`

SetNotificationMedium sets NotificationMedium field to given value.

### HasNotificationMedium

`func (o *ViewProjectBudgetNotification) HasNotificationMedium() bool`

HasNotificationMedium returns a boolean if a field has been set.

### GetTeamIds

`func (o *ViewProjectBudgetNotification) GetTeamIds() []int32`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *ViewProjectBudgetNotification) GetTeamIdsOk() (*[]int32, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *ViewProjectBudgetNotification) SetTeamIds(v []int32)`

SetTeamIds sets TeamIds field to given value.

### HasTeamIds

`func (o *ViewProjectBudgetNotification) HasTeamIds() bool`

HasTeamIds returns a boolean if a field has been set.

### GetTeams

`func (o *ViewProjectBudgetNotification) GetTeams() []ViewRelationship`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *ViewProjectBudgetNotification) GetTeamsOk() (*[]ViewRelationship, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *ViewProjectBudgetNotification) SetTeams(v []ViewRelationship)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *ViewProjectBudgetNotification) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetUserIds

`func (o *ViewProjectBudgetNotification) GetUserIds() []int32`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *ViewProjectBudgetNotification) GetUserIdsOk() (*[]int32, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *ViewProjectBudgetNotification) SetUserIds(v []int32)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *ViewProjectBudgetNotification) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### GetUsers

`func (o *ViewProjectBudgetNotification) GetUsers() []ViewRelationship`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ViewProjectBudgetNotification) GetUsersOk() (*[]ViewRelationship, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ViewProjectBudgetNotification) SetUsers(v []ViewRelationship)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *ViewProjectBudgetNotification) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


