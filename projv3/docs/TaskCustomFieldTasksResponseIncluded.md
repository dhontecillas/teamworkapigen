# TaskCustomFieldTasksResponseIncluded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customfields** | Pointer to [**map[string]ViewCustomField**](ViewCustomField.md) |  | [optional] 
**Tasks** | Pointer to [**map[string]ViewTask**](ViewTask.md) |  | [optional] 

## Methods

### NewTaskCustomFieldTasksResponseIncluded

`func NewTaskCustomFieldTasksResponseIncluded() *TaskCustomFieldTasksResponseIncluded`

NewTaskCustomFieldTasksResponseIncluded instantiates a new TaskCustomFieldTasksResponseIncluded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskCustomFieldTasksResponseIncludedWithDefaults

`func NewTaskCustomFieldTasksResponseIncludedWithDefaults() *TaskCustomFieldTasksResponseIncluded`

NewTaskCustomFieldTasksResponseIncludedWithDefaults instantiates a new TaskCustomFieldTasksResponseIncluded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomfields

`func (o *TaskCustomFieldTasksResponseIncluded) GetCustomfields() map[string]ViewCustomField`

GetCustomfields returns the Customfields field if non-nil, zero value otherwise.

### GetCustomfieldsOk

`func (o *TaskCustomFieldTasksResponseIncluded) GetCustomfieldsOk() (*map[string]ViewCustomField, bool)`

GetCustomfieldsOk returns a tuple with the Customfields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomfields

`func (o *TaskCustomFieldTasksResponseIncluded) SetCustomfields(v map[string]ViewCustomField)`

SetCustomfields sets Customfields field to given value.

### HasCustomfields

`func (o *TaskCustomFieldTasksResponseIncluded) HasCustomfields() bool`

HasCustomfields returns a boolean if a field has been set.

### GetTasks

`func (o *TaskCustomFieldTasksResponseIncluded) GetTasks() map[string]ViewTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *TaskCustomFieldTasksResponseIncluded) GetTasksOk() (*map[string]ViewTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *TaskCustomFieldTasksResponseIncluded) SetTasks(v map[string]ViewTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *TaskCustomFieldTasksResponseIncluded) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


