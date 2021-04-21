# ViewUserSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocatedTotalMinutes** | Pointer to **int32** |  | [optional] 
**LoggedTotalMinutes** | Pointer to **int32** |  | [optional] 
**ProjectIds** | Pointer to **[]int32** |  | [optional] 
**Projects** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**Schedule** | Pointer to [**map[string]ViewScheduleEntry**](ViewScheduleEntry.md) |  | [optional] 
**UnavailableTotalMinutes** | Pointer to **int32** |  | [optional] 
**User** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 

## Methods

### NewViewUserSchedule

`func NewViewUserSchedule() *ViewUserSchedule`

NewViewUserSchedule instantiates a new ViewUserSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewUserScheduleWithDefaults

`func NewViewUserScheduleWithDefaults() *ViewUserSchedule`

NewViewUserScheduleWithDefaults instantiates a new ViewUserSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocatedTotalMinutes

`func (o *ViewUserSchedule) GetAllocatedTotalMinutes() int32`

GetAllocatedTotalMinutes returns the AllocatedTotalMinutes field if non-nil, zero value otherwise.

### GetAllocatedTotalMinutesOk

`func (o *ViewUserSchedule) GetAllocatedTotalMinutesOk() (*int32, bool)`

GetAllocatedTotalMinutesOk returns a tuple with the AllocatedTotalMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedTotalMinutes

`func (o *ViewUserSchedule) SetAllocatedTotalMinutes(v int32)`

SetAllocatedTotalMinutes sets AllocatedTotalMinutes field to given value.

### HasAllocatedTotalMinutes

`func (o *ViewUserSchedule) HasAllocatedTotalMinutes() bool`

HasAllocatedTotalMinutes returns a boolean if a field has been set.

### GetLoggedTotalMinutes

`func (o *ViewUserSchedule) GetLoggedTotalMinutes() int32`

GetLoggedTotalMinutes returns the LoggedTotalMinutes field if non-nil, zero value otherwise.

### GetLoggedTotalMinutesOk

`func (o *ViewUserSchedule) GetLoggedTotalMinutesOk() (*int32, bool)`

GetLoggedTotalMinutesOk returns a tuple with the LoggedTotalMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggedTotalMinutes

`func (o *ViewUserSchedule) SetLoggedTotalMinutes(v int32)`

SetLoggedTotalMinutes sets LoggedTotalMinutes field to given value.

### HasLoggedTotalMinutes

`func (o *ViewUserSchedule) HasLoggedTotalMinutes() bool`

HasLoggedTotalMinutes returns a boolean if a field has been set.

### GetProjectIds

`func (o *ViewUserSchedule) GetProjectIds() []int32`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *ViewUserSchedule) GetProjectIdsOk() (*[]int32, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *ViewUserSchedule) SetProjectIds(v []int32)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *ViewUserSchedule) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### GetProjects

`func (o *ViewUserSchedule) GetProjects() []ViewRelationship`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *ViewUserSchedule) GetProjectsOk() (*[]ViewRelationship, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *ViewUserSchedule) SetProjects(v []ViewRelationship)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *ViewUserSchedule) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetSchedule

`func (o *ViewUserSchedule) GetSchedule() map[string]ViewScheduleEntry`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ViewUserSchedule) GetScheduleOk() (*map[string]ViewScheduleEntry, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ViewUserSchedule) SetSchedule(v map[string]ViewScheduleEntry)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ViewUserSchedule) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetUnavailableTotalMinutes

`func (o *ViewUserSchedule) GetUnavailableTotalMinutes() int32`

GetUnavailableTotalMinutes returns the UnavailableTotalMinutes field if non-nil, zero value otherwise.

### GetUnavailableTotalMinutesOk

`func (o *ViewUserSchedule) GetUnavailableTotalMinutesOk() (*int32, bool)`

GetUnavailableTotalMinutesOk returns a tuple with the UnavailableTotalMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailableTotalMinutes

`func (o *ViewUserSchedule) SetUnavailableTotalMinutes(v int32)`

SetUnavailableTotalMinutes sets UnavailableTotalMinutes field to given value.

### HasUnavailableTotalMinutes

`func (o *ViewUserSchedule) HasUnavailableTotalMinutes() bool`

HasUnavailableTotalMinutes returns a boolean if a field has been set.

### GetUser

`func (o *ViewUserSchedule) GetUser() ViewRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ViewUserSchedule) GetUserOk() (*ViewRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ViewUserSchedule) SetUser(v ViewRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *ViewUserSchedule) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserId

`func (o *ViewUserSchedule) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ViewUserSchedule) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ViewUserSchedule) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ViewUserSchedule) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


