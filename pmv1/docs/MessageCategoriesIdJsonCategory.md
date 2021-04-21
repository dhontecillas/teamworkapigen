# MessageCategoriesIdJsonCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ParentId** | Pointer to **string** | If you pass in a parent id for a category, you will be creating a sub category. If you don&#39;t include this parameter, you will be not be creating a sub category. | [optional] 

## Methods

### NewMessageCategoriesIdJsonCategory

`func NewMessageCategoriesIdJsonCategory(name string, ) *MessageCategoriesIdJsonCategory`

NewMessageCategoriesIdJsonCategory instantiates a new MessageCategoriesIdJsonCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageCategoriesIdJsonCategoryWithDefaults

`func NewMessageCategoriesIdJsonCategoryWithDefaults() *MessageCategoriesIdJsonCategory`

NewMessageCategoriesIdJsonCategoryWithDefaults instantiates a new MessageCategoriesIdJsonCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MessageCategoriesIdJsonCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MessageCategoriesIdJsonCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MessageCategoriesIdJsonCategory) SetName(v string)`

SetName sets Name field to given value.


### GetParentId

`func (o *MessageCategoriesIdJsonCategory) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *MessageCategoriesIdJsonCategory) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *MessageCategoriesIdJsonCategory) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *MessageCategoriesIdJsonCategory) HasParentId() bool`

HasParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


