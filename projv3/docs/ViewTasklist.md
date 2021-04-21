# ViewTasklist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Milestone** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**MilestoneId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NewTaskDefaults** | Pointer to [**ViewNewTaskDefaults**](ViewNewTaskDefaults.md) |  | [optional] 
**Project** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**TaskIds** | Pointer to **[]int32** |  | [optional] 
**Tasks** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 

## Methods

### NewViewTasklist

`func NewViewTasklist() *ViewTasklist`

NewViewTasklist instantiates a new ViewTasklist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewTasklistWithDefaults

`func NewViewTasklistWithDefaults() *ViewTasklist`

NewViewTasklistWithDefaults instantiates a new ViewTasklist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ViewTasklist) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewTasklist) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewTasklist) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewTasklist) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMilestone

`func (o *ViewTasklist) GetMilestone() ViewRelationship`

GetMilestone returns the Milestone field if non-nil, zero value otherwise.

### GetMilestoneOk

`func (o *ViewTasklist) GetMilestoneOk() (*ViewRelationship, bool)`

GetMilestoneOk returns a tuple with the Milestone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestone

`func (o *ViewTasklist) SetMilestone(v ViewRelationship)`

SetMilestone sets Milestone field to given value.

### HasMilestone

`func (o *ViewTasklist) HasMilestone() bool`

HasMilestone returns a boolean if a field has been set.

### GetMilestoneId

`func (o *ViewTasklist) GetMilestoneId() int32`

GetMilestoneId returns the MilestoneId field if non-nil, zero value otherwise.

### GetMilestoneIdOk

`func (o *ViewTasklist) GetMilestoneIdOk() (*int32, bool)`

GetMilestoneIdOk returns a tuple with the MilestoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestoneId

`func (o *ViewTasklist) SetMilestoneId(v int32)`

SetMilestoneId sets MilestoneId field to given value.

### HasMilestoneId

`func (o *ViewTasklist) HasMilestoneId() bool`

HasMilestoneId returns a boolean if a field has been set.

### GetName

`func (o *ViewTasklist) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ViewTasklist) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ViewTasklist) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ViewTasklist) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNewTaskDefaults

`func (o *ViewTasklist) GetNewTaskDefaults() ViewNewTaskDefaults`

GetNewTaskDefaults returns the NewTaskDefaults field if non-nil, zero value otherwise.

### GetNewTaskDefaultsOk

`func (o *ViewTasklist) GetNewTaskDefaultsOk() (*ViewNewTaskDefaults, bool)`

GetNewTaskDefaultsOk returns a tuple with the NewTaskDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTaskDefaults

`func (o *ViewTasklist) SetNewTaskDefaults(v ViewNewTaskDefaults)`

SetNewTaskDefaults sets NewTaskDefaults field to given value.

### HasNewTaskDefaults

`func (o *ViewTasklist) HasNewTaskDefaults() bool`

HasNewTaskDefaults returns a boolean if a field has been set.

### GetProject

`func (o *ViewTasklist) GetProject() ViewRelationship`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ViewTasklist) GetProjectOk() (*ViewRelationship, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ViewTasklist) SetProject(v ViewRelationship)`

SetProject sets Project field to given value.

### HasProject

`func (o *ViewTasklist) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectId

`func (o *ViewTasklist) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ViewTasklist) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ViewTasklist) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ViewTasklist) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetTaskIds

`func (o *ViewTasklist) GetTaskIds() []int32`

GetTaskIds returns the TaskIds field if non-nil, zero value otherwise.

### GetTaskIdsOk

`func (o *ViewTasklist) GetTaskIdsOk() (*[]int32, bool)`

GetTaskIdsOk returns a tuple with the TaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskIds

`func (o *ViewTasklist) SetTaskIds(v []int32)`

SetTaskIds sets TaskIds field to given value.

### HasTaskIds

`func (o *ViewTasklist) HasTaskIds() bool`

HasTaskIds returns a boolean if a field has been set.

### GetTasks

`func (o *ViewTasklist) GetTasks() []ViewRelationship`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ViewTasklist) GetTasksOk() (*[]ViewRelationship, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ViewTasklist) SetTasks(v []ViewRelationship)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *ViewTasklist) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


