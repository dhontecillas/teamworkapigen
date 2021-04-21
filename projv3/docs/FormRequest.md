# FormRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Form** | Pointer to [**FormForm**](FormForm.md) |  | [optional] 

## Methods

### NewFormRequest

`func NewFormRequest() *FormRequest`

NewFormRequest instantiates a new FormRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormRequestWithDefaults

`func NewFormRequestWithDefaults() *FormRequest`

NewFormRequestWithDefaults instantiates a new FormRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForm

`func (o *FormRequest) GetForm() FormForm`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *FormRequest) GetFormOk() (*FormForm, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *FormRequest) SetForm(v FormForm)`

SetForm sets Form field to given value.

### HasForm

`func (o *FormRequest) HasForm() bool`

HasForm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


