# TaskCustomFieldTasksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomfieldTasks** | Pointer to [**[]ViewCustomFieldTask**](ViewCustomFieldTask.md) |  | [optional] 
**Included** | Pointer to [**TaskCustomFieldTasksResponseIncluded**](TaskCustomFieldTasksResponseIncluded.md) |  | [optional] 
**Meta** | Pointer to [**ViewMeta**](ViewMeta.md) |  | [optional] 

## Methods

### NewTaskCustomFieldTasksResponse

`func NewTaskCustomFieldTasksResponse() *TaskCustomFieldTasksResponse`

NewTaskCustomFieldTasksResponse instantiates a new TaskCustomFieldTasksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskCustomFieldTasksResponseWithDefaults

`func NewTaskCustomFieldTasksResponseWithDefaults() *TaskCustomFieldTasksResponse`

NewTaskCustomFieldTasksResponseWithDefaults instantiates a new TaskCustomFieldTasksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomfieldTasks

`func (o *TaskCustomFieldTasksResponse) GetCustomfieldTasks() []ViewCustomFieldTask`

GetCustomfieldTasks returns the CustomfieldTasks field if non-nil, zero value otherwise.

### GetCustomfieldTasksOk

`func (o *TaskCustomFieldTasksResponse) GetCustomfieldTasksOk() (*[]ViewCustomFieldTask, bool)`

GetCustomfieldTasksOk returns a tuple with the CustomfieldTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomfieldTasks

`func (o *TaskCustomFieldTasksResponse) SetCustomfieldTasks(v []ViewCustomFieldTask)`

SetCustomfieldTasks sets CustomfieldTasks field to given value.

### HasCustomfieldTasks

`func (o *TaskCustomFieldTasksResponse) HasCustomfieldTasks() bool`

HasCustomfieldTasks returns a boolean if a field has been set.

### GetIncluded

`func (o *TaskCustomFieldTasksResponse) GetIncluded() TaskCustomFieldTasksResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *TaskCustomFieldTasksResponse) GetIncludedOk() (*TaskCustomFieldTasksResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *TaskCustomFieldTasksResponse) SetIncluded(v TaskCustomFieldTasksResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *TaskCustomFieldTasksResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *TaskCustomFieldTasksResponse) GetMeta() ViewMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TaskCustomFieldTasksResponse) GetMetaOk() (*ViewMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TaskCustomFieldTasksResponse) SetMeta(v ViewMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TaskCustomFieldTasksResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


