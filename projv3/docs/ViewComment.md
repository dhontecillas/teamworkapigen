# ViewComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Object** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**ObjectId** | Pointer to **int32** |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewViewComment

`func NewViewComment() *ViewComment`

NewViewComment instantiates a new ViewComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewCommentWithDefaults

`func NewViewCommentWithDefaults() *ViewComment`

NewViewCommentWithDefaults instantiates a new ViewComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ViewComment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewComment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewComment) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewComment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *ViewComment) GetObject() ViewRelationship`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ViewComment) GetObjectOk() (*ViewRelationship, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ViewComment) SetObject(v ViewRelationship)`

SetObject sets Object field to given value.

### HasObject

`func (o *ViewComment) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetObjectId

`func (o *ViewComment) GetObjectId() int32`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *ViewComment) GetObjectIdOk() (*int32, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *ViewComment) SetObjectId(v int32)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *ViewComment) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetObjectType

`func (o *ViewComment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ViewComment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ViewComment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *ViewComment) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetTitle

`func (o *ViewComment) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ViewComment) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ViewComment) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ViewComment) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


