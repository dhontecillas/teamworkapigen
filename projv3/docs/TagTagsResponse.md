# TagTagsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ViewMeta**](ViewMeta.md) |  | [optional] 
**Tags** | Pointer to [**[]ViewTag**](ViewTag.md) |  | [optional] 

## Methods

### NewTagTagsResponse

`func NewTagTagsResponse() *TagTagsResponse`

NewTagTagsResponse instantiates a new TagTagsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagTagsResponseWithDefaults

`func NewTagTagsResponseWithDefaults() *TagTagsResponse`

NewTagTagsResponseWithDefaults instantiates a new TagTagsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *TagTagsResponse) GetMeta() ViewMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TagTagsResponse) GetMetaOk() (*ViewMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TagTagsResponse) SetMeta(v ViewMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TagTagsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetTags

`func (o *TagTagsResponse) GetTags() []ViewTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TagTagsResponse) GetTagsOk() (*[]ViewTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TagTagsResponse) SetTags(v []ViewTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TagTagsResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


