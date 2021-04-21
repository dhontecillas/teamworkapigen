# FormFormsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Forms** | Pointer to [**[]ViewForm**](ViewForm.md) |  | [optional] 
**Included** | Pointer to [**CalendarEventsResponseIncluded**](CalendarEventsResponseIncluded.md) |  | [optional] 
**Meta** | Pointer to [**ViewMeta**](ViewMeta.md) |  | [optional] 

## Methods

### NewFormFormsResponse

`func NewFormFormsResponse() *FormFormsResponse`

NewFormFormsResponse instantiates a new FormFormsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFormsResponseWithDefaults

`func NewFormFormsResponseWithDefaults() *FormFormsResponse`

NewFormFormsResponseWithDefaults instantiates a new FormFormsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForms

`func (o *FormFormsResponse) GetForms() []ViewForm`

GetForms returns the Forms field if non-nil, zero value otherwise.

### GetFormsOk

`func (o *FormFormsResponse) GetFormsOk() (*[]ViewForm, bool)`

GetFormsOk returns a tuple with the Forms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForms

`func (o *FormFormsResponse) SetForms(v []ViewForm)`

SetForms sets Forms field to given value.

### HasForms

`func (o *FormFormsResponse) HasForms() bool`

HasForms returns a boolean if a field has been set.

### GetIncluded

`func (o *FormFormsResponse) GetIncluded() CalendarEventsResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *FormFormsResponse) GetIncludedOk() (*CalendarEventsResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *FormFormsResponse) SetIncluded(v CalendarEventsResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *FormFormsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *FormFormsResponse) GetMeta() ViewMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FormFormsResponse) GetMetaOk() (*ViewMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FormFormsResponse) SetMeta(v ViewMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FormFormsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


