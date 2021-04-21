# BudgetCapacityUsedAtDate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** | RFC3339 formated datetime, when the day starts int UTC time | [optional] 
**Usage** | Pointer to **float32** |  | [optional] 

## Methods

### NewBudgetCapacityUsedAtDate

`func NewBudgetCapacityUsedAtDate() *BudgetCapacityUsedAtDate`

NewBudgetCapacityUsedAtDate instantiates a new BudgetCapacityUsedAtDate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetCapacityUsedAtDateWithDefaults

`func NewBudgetCapacityUsedAtDateWithDefaults() *BudgetCapacityUsedAtDate`

NewBudgetCapacityUsedAtDateWithDefaults instantiates a new BudgetCapacityUsedAtDate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *BudgetCapacityUsedAtDate) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *BudgetCapacityUsedAtDate) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *BudgetCapacityUsedAtDate) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *BudgetCapacityUsedAtDate) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetUsage

`func (o *BudgetCapacityUsedAtDate) GetUsage() float32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *BudgetCapacityUsedAtDate) GetUsageOk() (*float32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *BudgetCapacityUsedAtDate) SetUsage(v float32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *BudgetCapacityUsedAtDate) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


