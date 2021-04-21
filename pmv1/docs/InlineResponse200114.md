# InlineResponse200114

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**STATUS** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Files** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Links** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Messages** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Milestones** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Notebooks** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Tasklists** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Tasks** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Timelogs** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewInlineResponse200114

`func NewInlineResponse200114() *InlineResponse200114`

NewInlineResponse200114 instantiates a new InlineResponse200114 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200114WithDefaults

`func NewInlineResponse200114WithDefaults() *InlineResponse200114`

NewInlineResponse200114WithDefaults instantiates a new InlineResponse200114 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSTATUS

`func (o *InlineResponse200114) GetSTATUS() string`

GetSTATUS returns the STATUS field if non-nil, zero value otherwise.

### GetSTATUSOk

`func (o *InlineResponse200114) GetSTATUSOk() (*string, bool)`

GetSTATUSOk returns a tuple with the STATUS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS

`func (o *InlineResponse200114) SetSTATUS(v string)`

SetSTATUS sets STATUS field to given value.

### HasSTATUS

`func (o *InlineResponse200114) HasSTATUS() bool`

HasSTATUS returns a boolean if a field has been set.

### GetComments

`func (o *InlineResponse200114) GetComments() []map[string]interface{}`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *InlineResponse200114) GetCommentsOk() (*[]map[string]interface{}, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *InlineResponse200114) SetComments(v []map[string]interface{})`

SetComments sets Comments field to given value.

### HasComments

`func (o *InlineResponse200114) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetFiles

`func (o *InlineResponse200114) GetFiles() []map[string]interface{}`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *InlineResponse200114) GetFilesOk() (*[]map[string]interface{}, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *InlineResponse200114) SetFiles(v []map[string]interface{})`

SetFiles sets Files field to given value.

### HasFiles

`func (o *InlineResponse200114) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetLinks

`func (o *InlineResponse200114) GetLinks() []map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *InlineResponse200114) GetLinksOk() (*[]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *InlineResponse200114) SetLinks(v []map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *InlineResponse200114) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMessages

`func (o *InlineResponse200114) GetMessages() []map[string]interface{}`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *InlineResponse200114) GetMessagesOk() (*[]map[string]interface{}, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *InlineResponse200114) SetMessages(v []map[string]interface{})`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *InlineResponse200114) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetMilestones

`func (o *InlineResponse200114) GetMilestones() []map[string]interface{}`

GetMilestones returns the Milestones field if non-nil, zero value otherwise.

### GetMilestonesOk

`func (o *InlineResponse200114) GetMilestonesOk() (*[]map[string]interface{}, bool)`

GetMilestonesOk returns a tuple with the Milestones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestones

`func (o *InlineResponse200114) SetMilestones(v []map[string]interface{})`

SetMilestones sets Milestones field to given value.

### HasMilestones

`func (o *InlineResponse200114) HasMilestones() bool`

HasMilestones returns a boolean if a field has been set.

### GetNotebooks

`func (o *InlineResponse200114) GetNotebooks() []map[string]interface{}`

GetNotebooks returns the Notebooks field if non-nil, zero value otherwise.

### GetNotebooksOk

`func (o *InlineResponse200114) GetNotebooksOk() (*[]map[string]interface{}, bool)`

GetNotebooksOk returns a tuple with the Notebooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotebooks

`func (o *InlineResponse200114) SetNotebooks(v []map[string]interface{})`

SetNotebooks sets Notebooks field to given value.

### HasNotebooks

`func (o *InlineResponse200114) HasNotebooks() bool`

HasNotebooks returns a boolean if a field has been set.

### GetTasklists

`func (o *InlineResponse200114) GetTasklists() []map[string]interface{}`

GetTasklists returns the Tasklists field if non-nil, zero value otherwise.

### GetTasklistsOk

`func (o *InlineResponse200114) GetTasklistsOk() (*[]map[string]interface{}, bool)`

GetTasklistsOk returns a tuple with the Tasklists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasklists

`func (o *InlineResponse200114) SetTasklists(v []map[string]interface{})`

SetTasklists sets Tasklists field to given value.

### HasTasklists

`func (o *InlineResponse200114) HasTasklists() bool`

HasTasklists returns a boolean if a field has been set.

### GetTasks

`func (o *InlineResponse200114) GetTasks() []map[string]interface{}`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *InlineResponse200114) GetTasksOk() (*[]map[string]interface{}, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *InlineResponse200114) SetTasks(v []map[string]interface{})`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *InlineResponse200114) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetTimelogs

`func (o *InlineResponse200114) GetTimelogs() []map[string]interface{}`

GetTimelogs returns the Timelogs field if non-nil, zero value otherwise.

### GetTimelogsOk

`func (o *InlineResponse200114) GetTimelogsOk() (*[]map[string]interface{}, bool)`

GetTimelogsOk returns a tuple with the Timelogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelogs

`func (o *InlineResponse200114) SetTimelogs(v []map[string]interface{})`

SetTimelogs sets Timelogs field to given value.

### HasTimelogs

`func (o *InlineResponse200114) HasTimelogs() bool`

HasTimelogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


