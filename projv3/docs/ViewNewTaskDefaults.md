# ViewNewTaskDefaults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DueDate** | Pointer to **string** |  | [optional] 
**EstimatedMinutes** | Pointer to **int32** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**ResponisblePartyIds** | Pointer to [**ViewUserGroups**](ViewUserGroups.md) |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewViewNewTaskDefaults

`func NewViewNewTaskDefaults() *ViewNewTaskDefaults`

NewViewNewTaskDefaults instantiates a new ViewNewTaskDefaults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewNewTaskDefaultsWithDefaults

`func NewViewNewTaskDefaultsWithDefaults() *ViewNewTaskDefaults`

NewViewNewTaskDefaultsWithDefaults instantiates a new ViewNewTaskDefaults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDueDate

`func (o *ViewNewTaskDefaults) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *ViewNewTaskDefaults) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *ViewNewTaskDefaults) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *ViewNewTaskDefaults) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetEstimatedMinutes

`func (o *ViewNewTaskDefaults) GetEstimatedMinutes() int32`

GetEstimatedMinutes returns the EstimatedMinutes field if non-nil, zero value otherwise.

### GetEstimatedMinutesOk

`func (o *ViewNewTaskDefaults) GetEstimatedMinutesOk() (*int32, bool)`

GetEstimatedMinutesOk returns a tuple with the EstimatedMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedMinutes

`func (o *ViewNewTaskDefaults) SetEstimatedMinutes(v int32)`

SetEstimatedMinutes sets EstimatedMinutes field to given value.

### HasEstimatedMinutes

`func (o *ViewNewTaskDefaults) HasEstimatedMinutes() bool`

HasEstimatedMinutes returns a boolean if a field has been set.

### GetPriority

`func (o *ViewNewTaskDefaults) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ViewNewTaskDefaults) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ViewNewTaskDefaults) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ViewNewTaskDefaults) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetResponisblePartyIds

`func (o *ViewNewTaskDefaults) GetResponisblePartyIds() ViewUserGroups`

GetResponisblePartyIds returns the ResponisblePartyIds field if non-nil, zero value otherwise.

### GetResponisblePartyIdsOk

`func (o *ViewNewTaskDefaults) GetResponisblePartyIdsOk() (*ViewUserGroups, bool)`

GetResponisblePartyIdsOk returns a tuple with the ResponisblePartyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponisblePartyIds

`func (o *ViewNewTaskDefaults) SetResponisblePartyIds(v ViewUserGroups)`

SetResponisblePartyIds sets ResponisblePartyIds field to given value.

### HasResponisblePartyIds

`func (o *ViewNewTaskDefaults) HasResponisblePartyIds() bool`

HasResponisblePartyIds returns a boolean if a field has been set.

### GetStartDate

`func (o *ViewNewTaskDefaults) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ViewNewTaskDefaults) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ViewNewTaskDefaults) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ViewNewTaskDefaults) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTags

`func (o *ViewNewTaskDefaults) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ViewNewTaskDefaults) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ViewNewTaskDefaults) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ViewNewTaskDefaults) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


