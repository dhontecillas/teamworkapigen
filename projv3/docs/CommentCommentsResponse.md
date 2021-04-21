# CommentCommentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comments** | Pointer to [**[]CommentFullComment**](CommentFullComment.md) |  | [optional] 
**Included** | Pointer to [**CommentCommentsResponseIncluded**](CommentCommentsResponseIncluded.md) |  | [optional] 
**Meta** | Pointer to [**ViewMeta**](ViewMeta.md) |  | [optional] 

## Methods

### NewCommentCommentsResponse

`func NewCommentCommentsResponse() *CommentCommentsResponse`

NewCommentCommentsResponse instantiates a new CommentCommentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentCommentsResponseWithDefaults

`func NewCommentCommentsResponseWithDefaults() *CommentCommentsResponse`

NewCommentCommentsResponseWithDefaults instantiates a new CommentCommentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComments

`func (o *CommentCommentsResponse) GetComments() []CommentFullComment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *CommentCommentsResponse) GetCommentsOk() (*[]CommentFullComment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *CommentCommentsResponse) SetComments(v []CommentFullComment)`

SetComments sets Comments field to given value.

### HasComments

`func (o *CommentCommentsResponse) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetIncluded

`func (o *CommentCommentsResponse) GetIncluded() CommentCommentsResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *CommentCommentsResponse) GetIncludedOk() (*CommentCommentsResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *CommentCommentsResponse) SetIncluded(v CommentCommentsResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *CommentCommentsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *CommentCommentsResponse) GetMeta() ViewMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CommentCommentsResponse) GetMetaOk() (*ViewMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CommentCommentsResponse) SetMeta(v ViewMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CommentCommentsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


