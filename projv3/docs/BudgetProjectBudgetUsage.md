# BudgetProjectBudgetUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetSummary** | Pointer to [**BudgetProjectBudgetSummary**](BudgetProjectBudgetSummary.md) |  | [optional] 
**UsageTimeline** | Pointer to [**[]BudgetCapacityUsedAtDate**](BudgetCapacityUsedAtDate.md) |  | [optional] 

## Methods

### NewBudgetProjectBudgetUsage

`func NewBudgetProjectBudgetUsage() *BudgetProjectBudgetUsage`

NewBudgetProjectBudgetUsage instantiates a new BudgetProjectBudgetUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetProjectBudgetUsageWithDefaults

`func NewBudgetProjectBudgetUsageWithDefaults() *BudgetProjectBudgetUsage`

NewBudgetProjectBudgetUsageWithDefaults instantiates a new BudgetProjectBudgetUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgetSummary

`func (o *BudgetProjectBudgetUsage) GetBudgetSummary() BudgetProjectBudgetSummary`

GetBudgetSummary returns the BudgetSummary field if non-nil, zero value otherwise.

### GetBudgetSummaryOk

`func (o *BudgetProjectBudgetUsage) GetBudgetSummaryOk() (*BudgetProjectBudgetSummary, bool)`

GetBudgetSummaryOk returns a tuple with the BudgetSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetSummary

`func (o *BudgetProjectBudgetUsage) SetBudgetSummary(v BudgetProjectBudgetSummary)`

SetBudgetSummary sets BudgetSummary field to given value.

### HasBudgetSummary

`func (o *BudgetProjectBudgetUsage) HasBudgetSummary() bool`

HasBudgetSummary returns a boolean if a field has been set.

### GetUsageTimeline

`func (o *BudgetProjectBudgetUsage) GetUsageTimeline() []BudgetCapacityUsedAtDate`

GetUsageTimeline returns the UsageTimeline field if non-nil, zero value otherwise.

### GetUsageTimelineOk

`func (o *BudgetProjectBudgetUsage) GetUsageTimelineOk() (*[]BudgetCapacityUsedAtDate, bool)`

GetUsageTimelineOk returns a tuple with the UsageTimeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageTimeline

`func (o *BudgetProjectBudgetUsage) SetUsageTimeline(v []BudgetCapacityUsedAtDate)`

SetUsageTimeline sets UsageTimeline field to given value.

### HasUsageTimeline

`func (o *BudgetProjectBudgetUsage) HasUsageTimeline() bool`

HasUsageTimeline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


