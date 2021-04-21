# BudgetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Budget** | Pointer to [**ViewProjectBudget**](ViewProjectBudget.md) |  | [optional] 
**RepeatStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewBudgetResponse

`func NewBudgetResponse() *BudgetResponse`

NewBudgetResponse instantiates a new BudgetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetResponseWithDefaults

`func NewBudgetResponseWithDefaults() *BudgetResponse`

NewBudgetResponseWithDefaults instantiates a new BudgetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudget

`func (o *BudgetResponse) GetBudget() ViewProjectBudget`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *BudgetResponse) GetBudgetOk() (*ViewProjectBudget, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *BudgetResponse) SetBudget(v ViewProjectBudget)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *BudgetResponse) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetRepeatStatus

`func (o *BudgetResponse) GetRepeatStatus() string`

GetRepeatStatus returns the RepeatStatus field if non-nil, zero value otherwise.

### GetRepeatStatusOk

`func (o *BudgetResponse) GetRepeatStatusOk() (*string, bool)`

GetRepeatStatusOk returns a tuple with the RepeatStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatStatus

`func (o *BudgetResponse) SetRepeatStatus(v string)`

SetRepeatStatus sets RepeatStatus field to given value.

### HasRepeatStatus

`func (o *BudgetResponse) HasRepeatStatus() bool`

HasRepeatStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


