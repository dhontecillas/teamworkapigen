# TaskBulkUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomfieldTasks** | Pointer to [**[]TaskEditCustomFieldTask**](TaskEditCustomFieldTask.md) |  | [optional] 

## Methods

### NewTaskBulkUpdateRequest

`func NewTaskBulkUpdateRequest() *TaskBulkUpdateRequest`

NewTaskBulkUpdateRequest instantiates a new TaskBulkUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskBulkUpdateRequestWithDefaults

`func NewTaskBulkUpdateRequestWithDefaults() *TaskBulkUpdateRequest`

NewTaskBulkUpdateRequestWithDefaults instantiates a new TaskBulkUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomfieldTasks

`func (o *TaskBulkUpdateRequest) GetCustomfieldTasks() []TaskEditCustomFieldTask`

GetCustomfieldTasks returns the CustomfieldTasks field if non-nil, zero value otherwise.

### GetCustomfieldTasksOk

`func (o *TaskBulkUpdateRequest) GetCustomfieldTasksOk() (*[]TaskEditCustomFieldTask, bool)`

GetCustomfieldTasksOk returns a tuple with the CustomfieldTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomfieldTasks

`func (o *TaskBulkUpdateRequest) SetCustomfieldTasks(v []TaskEditCustomFieldTask)`

SetCustomfieldTasks sets CustomfieldTasks field to given value.

### HasCustomfieldTasks

`func (o *TaskBulkUpdateRequest) HasCustomfieldTasks() bool`

HasCustomfieldTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


