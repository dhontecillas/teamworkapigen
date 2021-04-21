# ViewPublicForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Banner** | Pointer to [**ViewBanner**](ViewBanner.md) |  | [optional] 
**ConfirmationMessage** | Pointer to **string** |  | [optional] 
**Content** | Pointer to [**ViewContent**](ViewContent.md) |  | [optional] 
**PromptAdditionalSubmissions** | Pointer to **bool** |  | [optional] 
**Token** | Pointer to [**ViewToken**](ViewToken.md) |  | [optional] 

## Methods

### NewViewPublicForm

`func NewViewPublicForm() *ViewPublicForm`

NewViewPublicForm instantiates a new ViewPublicForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewPublicFormWithDefaults

`func NewViewPublicFormWithDefaults() *ViewPublicForm`

NewViewPublicFormWithDefaults instantiates a new ViewPublicForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBanner

`func (o *ViewPublicForm) GetBanner() ViewBanner`

GetBanner returns the Banner field if non-nil, zero value otherwise.

### GetBannerOk

`func (o *ViewPublicForm) GetBannerOk() (*ViewBanner, bool)`

GetBannerOk returns a tuple with the Banner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanner

`func (o *ViewPublicForm) SetBanner(v ViewBanner)`

SetBanner sets Banner field to given value.

### HasBanner

`func (o *ViewPublicForm) HasBanner() bool`

HasBanner returns a boolean if a field has been set.

### GetConfirmationMessage

`func (o *ViewPublicForm) GetConfirmationMessage() string`

GetConfirmationMessage returns the ConfirmationMessage field if non-nil, zero value otherwise.

### GetConfirmationMessageOk

`func (o *ViewPublicForm) GetConfirmationMessageOk() (*string, bool)`

GetConfirmationMessageOk returns a tuple with the ConfirmationMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationMessage

`func (o *ViewPublicForm) SetConfirmationMessage(v string)`

SetConfirmationMessage sets ConfirmationMessage field to given value.

### HasConfirmationMessage

`func (o *ViewPublicForm) HasConfirmationMessage() bool`

HasConfirmationMessage returns a boolean if a field has been set.

### GetContent

`func (o *ViewPublicForm) GetContent() ViewContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ViewPublicForm) GetContentOk() (*ViewContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ViewPublicForm) SetContent(v ViewContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *ViewPublicForm) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetPromptAdditionalSubmissions

`func (o *ViewPublicForm) GetPromptAdditionalSubmissions() bool`

GetPromptAdditionalSubmissions returns the PromptAdditionalSubmissions field if non-nil, zero value otherwise.

### GetPromptAdditionalSubmissionsOk

`func (o *ViewPublicForm) GetPromptAdditionalSubmissionsOk() (*bool, bool)`

GetPromptAdditionalSubmissionsOk returns a tuple with the PromptAdditionalSubmissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptAdditionalSubmissions

`func (o *ViewPublicForm) SetPromptAdditionalSubmissions(v bool)`

SetPromptAdditionalSubmissions sets PromptAdditionalSubmissions field to given value.

### HasPromptAdditionalSubmissions

`func (o *ViewPublicForm) HasPromptAdditionalSubmissions() bool`

HasPromptAdditionalSubmissions returns a boolean if a field has been set.

### GetToken

`func (o *ViewPublicForm) GetToken() ViewToken`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ViewPublicForm) GetTokenOk() (*ViewToken, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ViewPublicForm) SetToken(v ViewToken)`

SetToken sets Token field to given value.

### HasToken

`func (o *ViewPublicForm) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


