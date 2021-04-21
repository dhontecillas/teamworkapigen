# ProjectsProjIdUpdateJsonUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Health** | Pointer to **string** | You can add health of the project when updating. It will attach a colour status to the project. Options include: 0 - no health status  1 - need attention - Red 2 - ready to release - yellow 3 - in progress - green | [optional] [default to ""]
**Text** | Pointer to **string** |  | [optional] 

## Methods

### NewProjectsProjIdUpdateJsonUpdate

`func NewProjectsProjIdUpdateJsonUpdate() *ProjectsProjIdUpdateJsonUpdate`

NewProjectsProjIdUpdateJsonUpdate instantiates a new ProjectsProjIdUpdateJsonUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsProjIdUpdateJsonUpdateWithDefaults

`func NewProjectsProjIdUpdateJsonUpdateWithDefaults() *ProjectsProjIdUpdateJsonUpdate`

NewProjectsProjIdUpdateJsonUpdateWithDefaults instantiates a new ProjectsProjIdUpdateJsonUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealth

`func (o *ProjectsProjIdUpdateJsonUpdate) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ProjectsProjIdUpdateJsonUpdate) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ProjectsProjIdUpdateJsonUpdate) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *ProjectsProjIdUpdateJsonUpdate) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetText

`func (o *ProjectsProjIdUpdateJsonUpdate) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ProjectsProjIdUpdateJsonUpdate) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ProjectsProjIdUpdateJsonUpdate) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ProjectsProjIdUpdateJsonUpdate) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


