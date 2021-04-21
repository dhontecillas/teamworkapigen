# ViewProjectSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocatedTotalMinutes** | Pointer to **int32** |  | [optional] 
**BudgetIds** | Pointer to **[]int32** |  | [optional] 
**BudgetMinutes** | Pointer to **int32** |  | [optional] 
**Budgets** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**LoggedTotalMinutes** | Pointer to **int32** |  | [optional] 
**Project** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**Schedule** | Pointer to [**map[string]ViewProjectScheduleEntry**](ViewProjectScheduleEntry.md) |  | [optional] 
**UnavailableTotalMinutes** | Pointer to **int32** |  | [optional] 
**UserIds** | Pointer to **[]int32** |  | [optional] 
**Users** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 

## Methods

### NewViewProjectSchedule

`func NewViewProjectSchedule() *ViewProjectSchedule`

NewViewProjectSchedule instantiates a new ViewProjectSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewProjectScheduleWithDefaults

`func NewViewProjectScheduleWithDefaults() *ViewProjectSchedule`

NewViewProjectScheduleWithDefaults instantiates a new ViewProjectSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocatedTotalMinutes

`func (o *ViewProjectSchedule) GetAllocatedTotalMinutes() int32`

GetAllocatedTotalMinutes returns the AllocatedTotalMinutes field if non-nil, zero value otherwise.

### GetAllocatedTotalMinutesOk

`func (o *ViewProjectSchedule) GetAllocatedTotalMinutesOk() (*int32, bool)`

GetAllocatedTotalMinutesOk returns a tuple with the AllocatedTotalMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedTotalMinutes

`func (o *ViewProjectSchedule) SetAllocatedTotalMinutes(v int32)`

SetAllocatedTotalMinutes sets AllocatedTotalMinutes field to given value.

### HasAllocatedTotalMinutes

`func (o *ViewProjectSchedule) HasAllocatedTotalMinutes() bool`

HasAllocatedTotalMinutes returns a boolean if a field has been set.

### GetBudgetIds

`func (o *ViewProjectSchedule) GetBudgetIds() []int32`

GetBudgetIds returns the BudgetIds field if non-nil, zero value otherwise.

### GetBudgetIdsOk

`func (o *ViewProjectSchedule) GetBudgetIdsOk() (*[]int32, bool)`

GetBudgetIdsOk returns a tuple with the BudgetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetIds

`func (o *ViewProjectSchedule) SetBudgetIds(v []int32)`

SetBudgetIds sets BudgetIds field to given value.

### HasBudgetIds

`func (o *ViewProjectSchedule) HasBudgetIds() bool`

HasBudgetIds returns a boolean if a field has been set.

### GetBudgetMinutes

`func (o *ViewProjectSchedule) GetBudgetMinutes() int32`

GetBudgetMinutes returns the BudgetMinutes field if non-nil, zero value otherwise.

### GetBudgetMinutesOk

`func (o *ViewProjectSchedule) GetBudgetMinutesOk() (*int32, bool)`

GetBudgetMinutesOk returns a tuple with the BudgetMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetMinutes

`func (o *ViewProjectSchedule) SetBudgetMinutes(v int32)`

SetBudgetMinutes sets BudgetMinutes field to given value.

### HasBudgetMinutes

`func (o *ViewProjectSchedule) HasBudgetMinutes() bool`

HasBudgetMinutes returns a boolean if a field has been set.

### GetBudgets

`func (o *ViewProjectSchedule) GetBudgets() []ViewRelationship`

GetBudgets returns the Budgets field if non-nil, zero value otherwise.

### GetBudgetsOk

`func (o *ViewProjectSchedule) GetBudgetsOk() (*[]ViewRelationship, bool)`

GetBudgetsOk returns a tuple with the Budgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgets

`func (o *ViewProjectSchedule) SetBudgets(v []ViewRelationship)`

SetBudgets sets Budgets field to given value.

### HasBudgets

`func (o *ViewProjectSchedule) HasBudgets() bool`

HasBudgets returns a boolean if a field has been set.

### GetLoggedTotalMinutes

`func (o *ViewProjectSchedule) GetLoggedTotalMinutes() int32`

GetLoggedTotalMinutes returns the LoggedTotalMinutes field if non-nil, zero value otherwise.

### GetLoggedTotalMinutesOk

`func (o *ViewProjectSchedule) GetLoggedTotalMinutesOk() (*int32, bool)`

GetLoggedTotalMinutesOk returns a tuple with the LoggedTotalMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggedTotalMinutes

`func (o *ViewProjectSchedule) SetLoggedTotalMinutes(v int32)`

SetLoggedTotalMinutes sets LoggedTotalMinutes field to given value.

### HasLoggedTotalMinutes

`func (o *ViewProjectSchedule) HasLoggedTotalMinutes() bool`

HasLoggedTotalMinutes returns a boolean if a field has been set.

### GetProject

`func (o *ViewProjectSchedule) GetProject() ViewRelationship`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ViewProjectSchedule) GetProjectOk() (*ViewRelationship, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ViewProjectSchedule) SetProject(v ViewRelationship)`

SetProject sets Project field to given value.

### HasProject

`func (o *ViewProjectSchedule) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectId

`func (o *ViewProjectSchedule) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ViewProjectSchedule) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ViewProjectSchedule) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ViewProjectSchedule) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetSchedule

`func (o *ViewProjectSchedule) GetSchedule() map[string]ViewProjectScheduleEntry`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ViewProjectSchedule) GetScheduleOk() (*map[string]ViewProjectScheduleEntry, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ViewProjectSchedule) SetSchedule(v map[string]ViewProjectScheduleEntry)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ViewProjectSchedule) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetUnavailableTotalMinutes

`func (o *ViewProjectSchedule) GetUnavailableTotalMinutes() int32`

GetUnavailableTotalMinutes returns the UnavailableTotalMinutes field if non-nil, zero value otherwise.

### GetUnavailableTotalMinutesOk

`func (o *ViewProjectSchedule) GetUnavailableTotalMinutesOk() (*int32, bool)`

GetUnavailableTotalMinutesOk returns a tuple with the UnavailableTotalMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailableTotalMinutes

`func (o *ViewProjectSchedule) SetUnavailableTotalMinutes(v int32)`

SetUnavailableTotalMinutes sets UnavailableTotalMinutes field to given value.

### HasUnavailableTotalMinutes

`func (o *ViewProjectSchedule) HasUnavailableTotalMinutes() bool`

HasUnavailableTotalMinutes returns a boolean if a field has been set.

### GetUserIds

`func (o *ViewProjectSchedule) GetUserIds() []int32`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *ViewProjectSchedule) GetUserIdsOk() (*[]int32, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *ViewProjectSchedule) SetUserIds(v []int32)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *ViewProjectSchedule) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### GetUsers

`func (o *ViewProjectSchedule) GetUsers() []ViewRelationship`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ViewProjectSchedule) GetUsersOk() (*[]ViewRelationship, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ViewProjectSchedule) SetUsers(v []ViewRelationship)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *ViewProjectSchedule) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


