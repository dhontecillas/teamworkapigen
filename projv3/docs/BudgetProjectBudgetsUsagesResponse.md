# BudgetProjectBudgetsUsagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetUsages** | Pointer to [**[]BudgetProjectBudgetUsage**](BudgetProjectBudgetUsage.md) |  | [optional] 
**Meta** | Pointer to [**ViewMeta**](ViewMeta.md) |  | [optional] 

## Methods

### NewBudgetProjectBudgetsUsagesResponse

`func NewBudgetProjectBudgetsUsagesResponse() *BudgetProjectBudgetsUsagesResponse`

NewBudgetProjectBudgetsUsagesResponse instantiates a new BudgetProjectBudgetsUsagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetProjectBudgetsUsagesResponseWithDefaults

`func NewBudgetProjectBudgetsUsagesResponseWithDefaults() *BudgetProjectBudgetsUsagesResponse`

NewBudgetProjectBudgetsUsagesResponseWithDefaults instantiates a new BudgetProjectBudgetsUsagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgetUsages

`func (o *BudgetProjectBudgetsUsagesResponse) GetBudgetUsages() []BudgetProjectBudgetUsage`

GetBudgetUsages returns the BudgetUsages field if non-nil, zero value otherwise.

### GetBudgetUsagesOk

`func (o *BudgetProjectBudgetsUsagesResponse) GetBudgetUsagesOk() (*[]BudgetProjectBudgetUsage, bool)`

GetBudgetUsagesOk returns a tuple with the BudgetUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetUsages

`func (o *BudgetProjectBudgetsUsagesResponse) SetBudgetUsages(v []BudgetProjectBudgetUsage)`

SetBudgetUsages sets BudgetUsages field to given value.

### HasBudgetUsages

`func (o *BudgetProjectBudgetsUsagesResponse) HasBudgetUsages() bool`

HasBudgetUsages returns a boolean if a field has been set.

### GetMeta

`func (o *BudgetProjectBudgetsUsagesResponse) GetMeta() ViewMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BudgetProjectBudgetsUsagesResponse) GetMetaOk() (*ViewMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BudgetProjectBudgetsUsagesResponse) SetMeta(v ViewMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BudgetProjectBudgetsUsagesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


