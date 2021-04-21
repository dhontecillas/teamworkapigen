# FormSubmissionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Form** | Pointer to [**FormSubmission**](FormSubmission.md) |  | [optional] 

## Methods

### NewFormSubmissionRequest

`func NewFormSubmissionRequest() *FormSubmissionRequest`

NewFormSubmissionRequest instantiates a new FormSubmissionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormSubmissionRequestWithDefaults

`func NewFormSubmissionRequestWithDefaults() *FormSubmissionRequest`

NewFormSubmissionRequestWithDefaults instantiates a new FormSubmissionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForm

`func (o *FormSubmissionRequest) GetForm() FormSubmission`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *FormSubmissionRequest) GetFormOk() (*FormSubmission, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *FormSubmissionRequest) SetForm(v FormSubmission)`

SetForm sets Form field to given value.

### HasForm

`func (o *FormSubmissionRequest) HasForm() bool`

HasForm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


