# ViewProjectMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Post** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**PostId** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewViewProjectMessage

`func NewViewProjectMessage() *ViewProjectMessage`

NewViewProjectMessage instantiates a new ViewProjectMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewProjectMessageWithDefaults

`func NewViewProjectMessageWithDefaults() *ViewProjectMessage`

NewViewProjectMessageWithDefaults instantiates a new ViewProjectMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ViewProjectMessage) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewProjectMessage) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewProjectMessage) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewProjectMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPost

`func (o *ViewProjectMessage) GetPost() ViewRelationship`

GetPost returns the Post field if non-nil, zero value otherwise.

### GetPostOk

`func (o *ViewProjectMessage) GetPostOk() (*ViewRelationship, bool)`

GetPostOk returns a tuple with the Post field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPost

`func (o *ViewProjectMessage) SetPost(v ViewRelationship)`

SetPost sets Post field to given value.

### HasPost

`func (o *ViewProjectMessage) HasPost() bool`

HasPost returns a boolean if a field has been set.

### GetPostId

`func (o *ViewProjectMessage) GetPostId() int32`

GetPostId returns the PostId field if non-nil, zero value otherwise.

### GetPostIdOk

`func (o *ViewProjectMessage) GetPostIdOk() (*int32, bool)`

GetPostIdOk returns a tuple with the PostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostId

`func (o *ViewProjectMessage) SetPostId(v int32)`

SetPostId sets PostId field to given value.

### HasPostId

`func (o *ViewProjectMessage) HasPostId() bool`

HasPostId returns a boolean if a field has been set.

### GetTitle

`func (o *ViewProjectMessage) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ViewProjectMessage) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ViewProjectMessage) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ViewProjectMessage) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


