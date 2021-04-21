# ViewCustomFieldTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **int32** |  | [optional] 
**Customfield** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**CustomfieldId** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Task** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**TaskId** | Pointer to **int32** |  | [optional] 
**Value** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewViewCustomFieldTask

`func NewViewCustomFieldTask() *ViewCustomFieldTask`

NewViewCustomFieldTask instantiates a new ViewCustomFieldTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewCustomFieldTaskWithDefaults

`func NewViewCustomFieldTaskWithDefaults() *ViewCustomFieldTask`

NewViewCustomFieldTaskWithDefaults instantiates a new ViewCustomFieldTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ViewCustomFieldTask) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ViewCustomFieldTask) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ViewCustomFieldTask) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ViewCustomFieldTask) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ViewCustomFieldTask) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ViewCustomFieldTask) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ViewCustomFieldTask) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ViewCustomFieldTask) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCustomfield

`func (o *ViewCustomFieldTask) GetCustomfield() ViewRelationship`

GetCustomfield returns the Customfield field if non-nil, zero value otherwise.

### GetCustomfieldOk

`func (o *ViewCustomFieldTask) GetCustomfieldOk() (*ViewRelationship, bool)`

GetCustomfieldOk returns a tuple with the Customfield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomfield

`func (o *ViewCustomFieldTask) SetCustomfield(v ViewRelationship)`

SetCustomfield sets Customfield field to given value.

### HasCustomfield

`func (o *ViewCustomFieldTask) HasCustomfield() bool`

HasCustomfield returns a boolean if a field has been set.

### GetCustomfieldId

`func (o *ViewCustomFieldTask) GetCustomfieldId() int32`

GetCustomfieldId returns the CustomfieldId field if non-nil, zero value otherwise.

### GetCustomfieldIdOk

`func (o *ViewCustomFieldTask) GetCustomfieldIdOk() (*int32, bool)`

GetCustomfieldIdOk returns a tuple with the CustomfieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomfieldId

`func (o *ViewCustomFieldTask) SetCustomfieldId(v int32)`

SetCustomfieldId sets CustomfieldId field to given value.

### HasCustomfieldId

`func (o *ViewCustomFieldTask) HasCustomfieldId() bool`

HasCustomfieldId returns a boolean if a field has been set.

### GetId

`func (o *ViewCustomFieldTask) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewCustomFieldTask) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewCustomFieldTask) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewCustomFieldTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTask

`func (o *ViewCustomFieldTask) GetTask() ViewRelationship`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *ViewCustomFieldTask) GetTaskOk() (*ViewRelationship, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *ViewCustomFieldTask) SetTask(v ViewRelationship)`

SetTask sets Task field to given value.

### HasTask

`func (o *ViewCustomFieldTask) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetTaskId

`func (o *ViewCustomFieldTask) GetTaskId() int32`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *ViewCustomFieldTask) GetTaskIdOk() (*int32, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *ViewCustomFieldTask) SetTaskId(v int32)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *ViewCustomFieldTask) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetValue

`func (o *ViewCustomFieldTask) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ViewCustomFieldTask) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ViewCustomFieldTask) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *ViewCustomFieldTask) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


