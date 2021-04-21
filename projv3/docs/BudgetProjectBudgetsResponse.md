# BudgetProjectBudgetsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Budgets** | Pointer to [**[]ViewProjectBudget**](ViewProjectBudget.md) |  | [optional] 
**Included** | Pointer to [**BudgetProjectBudgetsResponseIncluded**](BudgetProjectBudgetsResponseIncluded.md) |  | [optional] 
**Meta** | Pointer to [**ViewMeta**](ViewMeta.md) |  | [optional] 

## Methods

### NewBudgetProjectBudgetsResponse

`func NewBudgetProjectBudgetsResponse() *BudgetProjectBudgetsResponse`

NewBudgetProjectBudgetsResponse instantiates a new BudgetProjectBudgetsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetProjectBudgetsResponseWithDefaults

`func NewBudgetProjectBudgetsResponseWithDefaults() *BudgetProjectBudgetsResponse`

NewBudgetProjectBudgetsResponseWithDefaults instantiates a new BudgetProjectBudgetsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgets

`func (o *BudgetProjectBudgetsResponse) GetBudgets() []ViewProjectBudget`

GetBudgets returns the Budgets field if non-nil, zero value otherwise.

### GetBudgetsOk

`func (o *BudgetProjectBudgetsResponse) GetBudgetsOk() (*[]ViewProjectBudget, bool)`

GetBudgetsOk returns a tuple with the Budgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgets

`func (o *BudgetProjectBudgetsResponse) SetBudgets(v []ViewProjectBudget)`

SetBudgets sets Budgets field to given value.

### HasBudgets

`func (o *BudgetProjectBudgetsResponse) HasBudgets() bool`

HasBudgets returns a boolean if a field has been set.

### GetIncluded

`func (o *BudgetProjectBudgetsResponse) GetIncluded() BudgetProjectBudgetsResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *BudgetProjectBudgetsResponse) GetIncludedOk() (*BudgetProjectBudgetsResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *BudgetProjectBudgetsResponse) SetIncluded(v BudgetProjectBudgetsResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *BudgetProjectBudgetsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *BudgetProjectBudgetsResponse) GetMeta() ViewMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BudgetProjectBudgetsResponse) GetMetaOk() (*ViewMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BudgetProjectBudgetsResponse) SetMeta(v ViewMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BudgetProjectBudgetsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


