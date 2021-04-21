# TaskOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EveryoneMustDo** | Pointer to **bool** |  | [optional] 
**FireWebhook** | Pointer to **bool** |  | [optional] 
**IsTemplate** | Pointer to **bool** |  | [optional] 
**LogActivity** | Pointer to **bool** |  | [optional] 
**Notify** | Pointer to **bool** |  | [optional] 
**ParseInlineTags** | Pointer to **bool** |  | [optional] 
**PositionAfterTaskId** | Pointer to **int32** |  | [optional] 
**ShiftProjectDates** | Pointer to **bool** |  | [optional] 
**UseDefaults** | Pointer to **bool** |  | [optional] 
**UseNotifyViaTWIM** | Pointer to **bool** |  | [optional] 

## Methods

### NewTaskOptions

`func NewTaskOptions() *TaskOptions`

NewTaskOptions instantiates a new TaskOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskOptionsWithDefaults

`func NewTaskOptionsWithDefaults() *TaskOptions`

NewTaskOptionsWithDefaults instantiates a new TaskOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEveryoneMustDo

`func (o *TaskOptions) GetEveryoneMustDo() bool`

GetEveryoneMustDo returns the EveryoneMustDo field if non-nil, zero value otherwise.

### GetEveryoneMustDoOk

`func (o *TaskOptions) GetEveryoneMustDoOk() (*bool, bool)`

GetEveryoneMustDoOk returns a tuple with the EveryoneMustDo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEveryoneMustDo

`func (o *TaskOptions) SetEveryoneMustDo(v bool)`

SetEveryoneMustDo sets EveryoneMustDo field to given value.

### HasEveryoneMustDo

`func (o *TaskOptions) HasEveryoneMustDo() bool`

HasEveryoneMustDo returns a boolean if a field has been set.

### GetFireWebhook

`func (o *TaskOptions) GetFireWebhook() bool`

GetFireWebhook returns the FireWebhook field if non-nil, zero value otherwise.

### GetFireWebhookOk

`func (o *TaskOptions) GetFireWebhookOk() (*bool, bool)`

GetFireWebhookOk returns a tuple with the FireWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFireWebhook

`func (o *TaskOptions) SetFireWebhook(v bool)`

SetFireWebhook sets FireWebhook field to given value.

### HasFireWebhook

`func (o *TaskOptions) HasFireWebhook() bool`

HasFireWebhook returns a boolean if a field has been set.

### GetIsTemplate

`func (o *TaskOptions) GetIsTemplate() bool`

GetIsTemplate returns the IsTemplate field if non-nil, zero value otherwise.

### GetIsTemplateOk

`func (o *TaskOptions) GetIsTemplateOk() (*bool, bool)`

GetIsTemplateOk returns a tuple with the IsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemplate

`func (o *TaskOptions) SetIsTemplate(v bool)`

SetIsTemplate sets IsTemplate field to given value.

### HasIsTemplate

`func (o *TaskOptions) HasIsTemplate() bool`

HasIsTemplate returns a boolean if a field has been set.

### GetLogActivity

`func (o *TaskOptions) GetLogActivity() bool`

GetLogActivity returns the LogActivity field if non-nil, zero value otherwise.

### GetLogActivityOk

`func (o *TaskOptions) GetLogActivityOk() (*bool, bool)`

GetLogActivityOk returns a tuple with the LogActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogActivity

`func (o *TaskOptions) SetLogActivity(v bool)`

SetLogActivity sets LogActivity field to given value.

### HasLogActivity

`func (o *TaskOptions) HasLogActivity() bool`

HasLogActivity returns a boolean if a field has been set.

### GetNotify

`func (o *TaskOptions) GetNotify() bool`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *TaskOptions) GetNotifyOk() (*bool, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *TaskOptions) SetNotify(v bool)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *TaskOptions) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetParseInlineTags

`func (o *TaskOptions) GetParseInlineTags() bool`

GetParseInlineTags returns the ParseInlineTags field if non-nil, zero value otherwise.

### GetParseInlineTagsOk

`func (o *TaskOptions) GetParseInlineTagsOk() (*bool, bool)`

GetParseInlineTagsOk returns a tuple with the ParseInlineTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseInlineTags

`func (o *TaskOptions) SetParseInlineTags(v bool)`

SetParseInlineTags sets ParseInlineTags field to given value.

### HasParseInlineTags

`func (o *TaskOptions) HasParseInlineTags() bool`

HasParseInlineTags returns a boolean if a field has been set.

### GetPositionAfterTaskId

`func (o *TaskOptions) GetPositionAfterTaskId() int32`

GetPositionAfterTaskId returns the PositionAfterTaskId field if non-nil, zero value otherwise.

### GetPositionAfterTaskIdOk

`func (o *TaskOptions) GetPositionAfterTaskIdOk() (*int32, bool)`

GetPositionAfterTaskIdOk returns a tuple with the PositionAfterTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionAfterTaskId

`func (o *TaskOptions) SetPositionAfterTaskId(v int32)`

SetPositionAfterTaskId sets PositionAfterTaskId field to given value.

### HasPositionAfterTaskId

`func (o *TaskOptions) HasPositionAfterTaskId() bool`

HasPositionAfterTaskId returns a boolean if a field has been set.

### GetShiftProjectDates

`func (o *TaskOptions) GetShiftProjectDates() bool`

GetShiftProjectDates returns the ShiftProjectDates field if non-nil, zero value otherwise.

### GetShiftProjectDatesOk

`func (o *TaskOptions) GetShiftProjectDatesOk() (*bool, bool)`

GetShiftProjectDatesOk returns a tuple with the ShiftProjectDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiftProjectDates

`func (o *TaskOptions) SetShiftProjectDates(v bool)`

SetShiftProjectDates sets ShiftProjectDates field to given value.

### HasShiftProjectDates

`func (o *TaskOptions) HasShiftProjectDates() bool`

HasShiftProjectDates returns a boolean if a field has been set.

### GetUseDefaults

`func (o *TaskOptions) GetUseDefaults() bool`

GetUseDefaults returns the UseDefaults field if non-nil, zero value otherwise.

### GetUseDefaultsOk

`func (o *TaskOptions) GetUseDefaultsOk() (*bool, bool)`

GetUseDefaultsOk returns a tuple with the UseDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDefaults

`func (o *TaskOptions) SetUseDefaults(v bool)`

SetUseDefaults sets UseDefaults field to given value.

### HasUseDefaults

`func (o *TaskOptions) HasUseDefaults() bool`

HasUseDefaults returns a boolean if a field has been set.

### GetUseNotifyViaTWIM

`func (o *TaskOptions) GetUseNotifyViaTWIM() bool`

GetUseNotifyViaTWIM returns the UseNotifyViaTWIM field if non-nil, zero value otherwise.

### GetUseNotifyViaTWIMOk

`func (o *TaskOptions) GetUseNotifyViaTWIMOk() (*bool, bool)`

GetUseNotifyViaTWIMOk returns a tuple with the UseNotifyViaTWIM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNotifyViaTWIM

`func (o *TaskOptions) SetUseNotifyViaTWIM(v bool)`

SetUseNotifyViaTWIM sets UseNotifyViaTWIM field to given value.

### HasUseNotifyViaTWIM

`func (o *TaskOptions) HasUseNotifyViaTWIM() bool`

HasUseNotifyViaTWIM returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


