# ViewError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code is an application-specific error code. | [optional] 
**Detail** | Pointer to **string** | Detail is a human-readable explanation specific to this occurrence of the problem. | [optional] 
**Id** | Pointer to **string** | ID is a reference for this exact instance of the error. | [optional] 
**Meta** | Pointer to **map[string]interface{}** | Meta contains tags that are useful to detect specific issues. | [optional] 
**Title** | Pointer to **string** | Title is a short, human-readable summary of the problem. | [optional] 

## Methods

### NewViewError

`func NewViewError() *ViewError`

NewViewError instantiates a new ViewError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewErrorWithDefaults

`func NewViewErrorWithDefaults() *ViewError`

NewViewErrorWithDefaults instantiates a new ViewError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ViewError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ViewError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ViewError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ViewError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetail

`func (o *ViewError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ViewError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ViewError) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ViewError) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetId

`func (o *ViewError) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewError) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewError) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ViewError) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMeta

`func (o *ViewError) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ViewError) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ViewError) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ViewError) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetTitle

`func (o *ViewError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ViewError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ViewError) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ViewError) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


