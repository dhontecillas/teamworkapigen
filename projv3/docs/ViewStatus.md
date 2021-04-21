# ViewStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateLastUpdated** | Pointer to **string** |  | [optional] 
**DateTime** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Emoji** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 

## Methods

### NewViewStatus

`func NewViewStatus() *ViewStatus`

NewViewStatus instantiates a new ViewStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewStatusWithDefaults

`func NewViewStatusWithDefaults() *ViewStatus`

NewViewStatusWithDefaults instantiates a new ViewStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateLastUpdated

`func (o *ViewStatus) GetDateLastUpdated() string`

GetDateLastUpdated returns the DateLastUpdated field if non-nil, zero value otherwise.

### GetDateLastUpdatedOk

`func (o *ViewStatus) GetDateLastUpdatedOk() (*string, bool)`

GetDateLastUpdatedOk returns a tuple with the DateLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateLastUpdated

`func (o *ViewStatus) SetDateLastUpdated(v string)`

SetDateLastUpdated sets DateLastUpdated field to given value.

### HasDateLastUpdated

`func (o *ViewStatus) HasDateLastUpdated() bool`

HasDateLastUpdated returns a boolean if a field has been set.

### GetDateTime

`func (o *ViewStatus) GetDateTime() string`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *ViewStatus) GetDateTimeOk() (*string, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *ViewStatus) SetDateTime(v string)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *ViewStatus) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetDeleted

`func (o *ViewStatus) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *ViewStatus) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *ViewStatus) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *ViewStatus) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetEmoji

`func (o *ViewStatus) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *ViewStatus) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *ViewStatus) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *ViewStatus) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### GetId

`func (o *ViewStatus) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewStatus) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewStatus) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpAddress

`func (o *ViewStatus) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ViewStatus) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ViewStatus) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ViewStatus) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetText

`func (o *ViewStatus) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ViewStatus) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ViewStatus) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ViewStatus) HasText() bool`

HasText returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ViewStatus) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ViewStatus) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ViewStatus) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ViewStatus) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUser

`func (o *ViewStatus) GetUser() ViewRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ViewStatus) GetUserOk() (*ViewRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ViewStatus) SetUser(v ViewRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *ViewStatus) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserId

`func (o *ViewStatus) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ViewStatus) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ViewStatus) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ViewStatus) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


