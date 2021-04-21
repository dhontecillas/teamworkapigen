# TaskRequestAttachments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to [**[]TaskFile**](TaskFile.md) |  | [optional] 
**PendingFiles** | Pointer to [**[]TaskPendingFile**](TaskPendingFile.md) |  | [optional] 

## Methods

### NewTaskRequestAttachments

`func NewTaskRequestAttachments() *TaskRequestAttachments`

NewTaskRequestAttachments instantiates a new TaskRequestAttachments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskRequestAttachmentsWithDefaults

`func NewTaskRequestAttachmentsWithDefaults() *TaskRequestAttachments`

NewTaskRequestAttachmentsWithDefaults instantiates a new TaskRequestAttachments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *TaskRequestAttachments) GetFiles() []TaskFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *TaskRequestAttachments) GetFilesOk() (*[]TaskFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *TaskRequestAttachments) SetFiles(v []TaskFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *TaskRequestAttachments) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetPendingFiles

`func (o *TaskRequestAttachments) GetPendingFiles() []TaskPendingFile`

GetPendingFiles returns the PendingFiles field if non-nil, zero value otherwise.

### GetPendingFilesOk

`func (o *TaskRequestAttachments) GetPendingFilesOk() (*[]TaskPendingFile, bool)`

GetPendingFilesOk returns a tuple with the PendingFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingFiles

`func (o *TaskRequestAttachments) SetPendingFiles(v []TaskPendingFile)`

SetPendingFiles sets PendingFiles field to given value.

### HasPendingFiles

`func (o *TaskRequestAttachments) HasPendingFiles() bool`

HasPendingFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


