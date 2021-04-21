# BudgetProjectBudgetSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to **int32** |  | [optional] 
**CapacityUsed** | Pointer to **int32** |  | [optional] 
**EndDateTime** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Project** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**StartDateTime** | Pointer to **string** |  | [optional] 
**TimelogType** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewBudgetProjectBudgetSummary

`func NewBudgetProjectBudgetSummary() *BudgetProjectBudgetSummary`

NewBudgetProjectBudgetSummary instantiates a new BudgetProjectBudgetSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetProjectBudgetSummaryWithDefaults

`func NewBudgetProjectBudgetSummaryWithDefaults() *BudgetProjectBudgetSummary`

NewBudgetProjectBudgetSummaryWithDefaults instantiates a new BudgetProjectBudgetSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *BudgetProjectBudgetSummary) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *BudgetProjectBudgetSummary) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *BudgetProjectBudgetSummary) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *BudgetProjectBudgetSummary) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetCapacityUsed

`func (o *BudgetProjectBudgetSummary) GetCapacityUsed() int32`

GetCapacityUsed returns the CapacityUsed field if non-nil, zero value otherwise.

### GetCapacityUsedOk

`func (o *BudgetProjectBudgetSummary) GetCapacityUsedOk() (*int32, bool)`

GetCapacityUsedOk returns a tuple with the CapacityUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityUsed

`func (o *BudgetProjectBudgetSummary) SetCapacityUsed(v int32)`

SetCapacityUsed sets CapacityUsed field to given value.

### HasCapacityUsed

`func (o *BudgetProjectBudgetSummary) HasCapacityUsed() bool`

HasCapacityUsed returns a boolean if a field has been set.

### GetEndDateTime

`func (o *BudgetProjectBudgetSummary) GetEndDateTime() string`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *BudgetProjectBudgetSummary) GetEndDateTimeOk() (*string, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *BudgetProjectBudgetSummary) SetEndDateTime(v string)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *BudgetProjectBudgetSummary) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### GetId

`func (o *BudgetProjectBudgetSummary) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BudgetProjectBudgetSummary) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BudgetProjectBudgetSummary) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BudgetProjectBudgetSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProject

`func (o *BudgetProjectBudgetSummary) GetProject() ViewRelationship`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *BudgetProjectBudgetSummary) GetProjectOk() (*ViewRelationship, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *BudgetProjectBudgetSummary) SetProject(v ViewRelationship)`

SetProject sets Project field to given value.

### HasProject

`func (o *BudgetProjectBudgetSummary) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectId

`func (o *BudgetProjectBudgetSummary) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BudgetProjectBudgetSummary) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BudgetProjectBudgetSummary) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BudgetProjectBudgetSummary) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetStartDateTime

`func (o *BudgetProjectBudgetSummary) GetStartDateTime() string`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *BudgetProjectBudgetSummary) GetStartDateTimeOk() (*string, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *BudgetProjectBudgetSummary) SetStartDateTime(v string)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *BudgetProjectBudgetSummary) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetTimelogType

`func (o *BudgetProjectBudgetSummary) GetTimelogType() string`

GetTimelogType returns the TimelogType field if non-nil, zero value otherwise.

### GetTimelogTypeOk

`func (o *BudgetProjectBudgetSummary) GetTimelogTypeOk() (*string, bool)`

GetTimelogTypeOk returns a tuple with the TimelogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelogType

`func (o *BudgetProjectBudgetSummary) SetTimelogType(v string)`

SetTimelogType sets TimelogType field to given value.

### HasTimelogType

`func (o *BudgetProjectBudgetSummary) HasTimelogType() bool`

HasTimelogType returns a boolean if a field has been set.

### GetType

`func (o *BudgetProjectBudgetSummary) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BudgetProjectBudgetSummary) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BudgetProjectBudgetSummary) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BudgetProjectBudgetSummary) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


