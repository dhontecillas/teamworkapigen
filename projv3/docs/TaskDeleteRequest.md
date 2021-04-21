# TaskDeleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskIds** | Pointer to **[]int32** |  | [optional] 
**TaskOptions** | Pointer to [**TaskCommonOptions**](TaskCommonOptions.md) |  | [optional] 

## Methods

### NewTaskDeleteRequest

`func NewTaskDeleteRequest() *TaskDeleteRequest`

NewTaskDeleteRequest instantiates a new TaskDeleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskDeleteRequestWithDefaults

`func NewTaskDeleteRequestWithDefaults() *TaskDeleteRequest`

NewTaskDeleteRequestWithDefaults instantiates a new TaskDeleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskIds

`func (o *TaskDeleteRequest) GetTaskIds() []int32`

GetTaskIds returns the TaskIds field if non-nil, zero value otherwise.

### GetTaskIdsOk

`func (o *TaskDeleteRequest) GetTaskIdsOk() (*[]int32, bool)`

GetTaskIdsOk returns a tuple with the TaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskIds

`func (o *TaskDeleteRequest) SetTaskIds(v []int32)`

SetTaskIds sets TaskIds field to given value.

### HasTaskIds

`func (o *TaskDeleteRequest) HasTaskIds() bool`

HasTaskIds returns a boolean if a field has been set.

### GetTaskOptions

`func (o *TaskDeleteRequest) GetTaskOptions() TaskCommonOptions`

GetTaskOptions returns the TaskOptions field if non-nil, zero value otherwise.

### GetTaskOptionsOk

`func (o *TaskDeleteRequest) GetTaskOptionsOk() (*TaskCommonOptions, bool)`

GetTaskOptionsOk returns a tuple with the TaskOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskOptions

`func (o *TaskDeleteRequest) SetTaskOptions(v TaskCommonOptions)`

SetTaskOptions sets TaskOptions field to given value.

### HasTaskOptions

`func (o *TaskDeleteRequest) HasTaskOptions() bool`

HasTaskOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


