# FormResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Form** | Pointer to [**ViewForm**](ViewForm.md) |  | [optional] 
**Included** | Pointer to [**CalendarEventsResponseIncluded**](CalendarEventsResponseIncluded.md) |  | [optional] 

## Methods

### NewFormResponse

`func NewFormResponse() *FormResponse`

NewFormResponse instantiates a new FormResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormResponseWithDefaults

`func NewFormResponseWithDefaults() *FormResponse`

NewFormResponseWithDefaults instantiates a new FormResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForm

`func (o *FormResponse) GetForm() ViewForm`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *FormResponse) GetFormOk() (*ViewForm, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *FormResponse) SetForm(v ViewForm)`

SetForm sets Form field to given value.

### HasForm

`func (o *FormResponse) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetIncluded

`func (o *FormResponse) GetIncluded() CalendarEventsResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *FormResponse) GetIncludedOk() (*CalendarEventsResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *FormResponse) SetIncluded(v CalendarEventsResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *FormResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


