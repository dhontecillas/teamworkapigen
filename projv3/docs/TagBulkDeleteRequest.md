# TagBulkDeleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagIds** | Pointer to **[]int32** |  | [optional] 
**TagNames** | Pointer to **string** |  | [optional] 

## Methods

### NewTagBulkDeleteRequest

`func NewTagBulkDeleteRequest() *TagBulkDeleteRequest`

NewTagBulkDeleteRequest instantiates a new TagBulkDeleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagBulkDeleteRequestWithDefaults

`func NewTagBulkDeleteRequestWithDefaults() *TagBulkDeleteRequest`

NewTagBulkDeleteRequestWithDefaults instantiates a new TagBulkDeleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagIds

`func (o *TagBulkDeleteRequest) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *TagBulkDeleteRequest) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *TagBulkDeleteRequest) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *TagBulkDeleteRequest) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTagNames

`func (o *TagBulkDeleteRequest) GetTagNames() string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *TagBulkDeleteRequest) GetTagNamesOk() (*string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *TagBulkDeleteRequest) SetTagNames(v string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *TagBulkDeleteRequest) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


