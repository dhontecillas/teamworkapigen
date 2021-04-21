# TaskTasksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Included** | Pointer to [**TaskResponseIncluded**](TaskResponseIncluded.md) |  | [optional] 
**Meta** | Pointer to [**ViewMeta**](ViewMeta.md) |  | [optional] 
**Tasks** | Pointer to [**[]ViewTask**](ViewTask.md) |  | [optional] 

## Methods

### NewTaskTasksResponse

`func NewTaskTasksResponse() *TaskTasksResponse`

NewTaskTasksResponse instantiates a new TaskTasksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskTasksResponseWithDefaults

`func NewTaskTasksResponseWithDefaults() *TaskTasksResponse`

NewTaskTasksResponseWithDefaults instantiates a new TaskTasksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncluded

`func (o *TaskTasksResponse) GetIncluded() TaskResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *TaskTasksResponse) GetIncludedOk() (*TaskResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *TaskTasksResponse) SetIncluded(v TaskResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *TaskTasksResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *TaskTasksResponse) GetMeta() ViewMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TaskTasksResponse) GetMetaOk() (*ViewMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TaskTasksResponse) SetMeta(v ViewMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TaskTasksResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetTasks

`func (o *TaskTasksResponse) GetTasks() []ViewTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *TaskTasksResponse) GetTasksOk() (*[]ViewTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *TaskTasksResponse) SetTasks(v []ViewTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *TaskTasksResponse) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


