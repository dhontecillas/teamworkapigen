# BudgetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Budget** | Pointer to [**BudgetProjectBudget**](BudgetProjectBudget.md) |  | [optional] 

## Methods

### NewBudgetRequest

`func NewBudgetRequest() *BudgetRequest`

NewBudgetRequest instantiates a new BudgetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetRequestWithDefaults

`func NewBudgetRequestWithDefaults() *BudgetRequest`

NewBudgetRequestWithDefaults instantiates a new BudgetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudget

`func (o *BudgetRequest) GetBudget() BudgetProjectBudget`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *BudgetRequest) GetBudgetOk() (*BudgetProjectBudget, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *BudgetRequest) SetBudget(v BudgetProjectBudget)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *BudgetRequest) HasBudget() bool`

HasBudget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


