# TaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachmentOptions** | Pointer to [**TaskRequestAttachmentOptions**](TaskRequestAttachmentOptions.md) |  | [optional] 
**Attachments** | Pointer to [**TaskRequestAttachments**](TaskRequestAttachments.md) |  | [optional] 
**Card** | Pointer to [**TaskCard**](TaskCard.md) |  | [optional] 
**Predecessors** | Pointer to [**[]TaskPredecessor**](TaskPredecessor.md) |  | [optional] 
**Tags** | Pointer to [**[]TagTag**](TagTag.md) |  | [optional] 
**Task** | Pointer to [**TaskTask**](TaskTask.md) |  | [optional] 
**TaskOptions** | Pointer to [**TaskOptions**](TaskOptions.md) |  | [optional] 

## Methods

### NewTaskRequest

`func NewTaskRequest() *TaskRequest`

NewTaskRequest instantiates a new TaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskRequestWithDefaults

`func NewTaskRequestWithDefaults() *TaskRequest`

NewTaskRequestWithDefaults instantiates a new TaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachmentOptions

`func (o *TaskRequest) GetAttachmentOptions() TaskRequestAttachmentOptions`

GetAttachmentOptions returns the AttachmentOptions field if non-nil, zero value otherwise.

### GetAttachmentOptionsOk

`func (o *TaskRequest) GetAttachmentOptionsOk() (*TaskRequestAttachmentOptions, bool)`

GetAttachmentOptionsOk returns a tuple with the AttachmentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentOptions

`func (o *TaskRequest) SetAttachmentOptions(v TaskRequestAttachmentOptions)`

SetAttachmentOptions sets AttachmentOptions field to given value.

### HasAttachmentOptions

`func (o *TaskRequest) HasAttachmentOptions() bool`

HasAttachmentOptions returns a boolean if a field has been set.

### GetAttachments

`func (o *TaskRequest) GetAttachments() TaskRequestAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TaskRequest) GetAttachmentsOk() (*TaskRequestAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TaskRequest) SetAttachments(v TaskRequestAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *TaskRequest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetCard

`func (o *TaskRequest) GetCard() TaskCard`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *TaskRequest) GetCardOk() (*TaskCard, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *TaskRequest) SetCard(v TaskCard)`

SetCard sets Card field to given value.

### HasCard

`func (o *TaskRequest) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetPredecessors

`func (o *TaskRequest) GetPredecessors() []TaskPredecessor`

GetPredecessors returns the Predecessors field if non-nil, zero value otherwise.

### GetPredecessorsOk

`func (o *TaskRequest) GetPredecessorsOk() (*[]TaskPredecessor, bool)`

GetPredecessorsOk returns a tuple with the Predecessors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredecessors

`func (o *TaskRequest) SetPredecessors(v []TaskPredecessor)`

SetPredecessors sets Predecessors field to given value.

### HasPredecessors

`func (o *TaskRequest) HasPredecessors() bool`

HasPredecessors returns a boolean if a field has been set.

### GetTags

`func (o *TaskRequest) GetTags() []TagTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TaskRequest) GetTagsOk() (*[]TagTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TaskRequest) SetTags(v []TagTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TaskRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTask

`func (o *TaskRequest) GetTask() TaskTask`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *TaskRequest) GetTaskOk() (*TaskTask, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *TaskRequest) SetTask(v TaskTask)`

SetTask sets Task field to given value.

### HasTask

`func (o *TaskRequest) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetTaskOptions

`func (o *TaskRequest) GetTaskOptions() TaskOptions`

GetTaskOptions returns the TaskOptions field if non-nil, zero value otherwise.

### GetTaskOptionsOk

`func (o *TaskRequest) GetTaskOptionsOk() (*TaskOptions, bool)`

GetTaskOptionsOk returns a tuple with the TaskOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskOptions

`func (o *TaskRequest) SetTaskOptions(v TaskOptions)`

SetTaskOptions sets TaskOptions field to given value.

### HasTaskOptions

`func (o *TaskRequest) HasTaskOptions() bool`

HasTaskOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


