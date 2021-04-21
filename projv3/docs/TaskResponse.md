# TaskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Included** | Pointer to [**TaskResponseIncluded**](TaskResponseIncluded.md) |  | [optional] 
**Task** | Pointer to [**ViewTask**](ViewTask.md) |  | [optional] 

## Methods

### NewTaskResponse

`func NewTaskResponse() *TaskResponse`

NewTaskResponse instantiates a new TaskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskResponseWithDefaults

`func NewTaskResponseWithDefaults() *TaskResponse`

NewTaskResponseWithDefaults instantiates a new TaskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncluded

`func (o *TaskResponse) GetIncluded() TaskResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *TaskResponse) GetIncludedOk() (*TaskResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *TaskResponse) SetIncluded(v TaskResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *TaskResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetTask

`func (o *TaskResponse) GetTask() ViewTask`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *TaskResponse) GetTaskOk() (*ViewTask, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *TaskResponse) SetTask(v ViewTask)`

SetTask sets Task field to given value.

### HasTask

`func (o *TaskResponse) HasTask() bool`

HasTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


