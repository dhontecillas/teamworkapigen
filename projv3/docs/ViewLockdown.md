# ViewLockdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantAccessTo** | Pointer to [**ViewUserGroups**](ViewUserGroups.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Item** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**ItemID** | Pointer to **int32** |  | [optional] 
**ItemType** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**UserID** | Pointer to **int32** |  | [optional] 

## Methods

### NewViewLockdown

`func NewViewLockdown() *ViewLockdown`

NewViewLockdown instantiates a new ViewLockdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewLockdownWithDefaults

`func NewViewLockdownWithDefaults() *ViewLockdown`

NewViewLockdownWithDefaults instantiates a new ViewLockdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantAccessTo

`func (o *ViewLockdown) GetGrantAccessTo() ViewUserGroups`

GetGrantAccessTo returns the GrantAccessTo field if non-nil, zero value otherwise.

### GetGrantAccessToOk

`func (o *ViewLockdown) GetGrantAccessToOk() (*ViewUserGroups, bool)`

GetGrantAccessToOk returns a tuple with the GrantAccessTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantAccessTo

`func (o *ViewLockdown) SetGrantAccessTo(v ViewUserGroups)`

SetGrantAccessTo sets GrantAccessTo field to given value.

### HasGrantAccessTo

`func (o *ViewLockdown) HasGrantAccessTo() bool`

HasGrantAccessTo returns a boolean if a field has been set.

### GetId

`func (o *ViewLockdown) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewLockdown) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewLockdown) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewLockdown) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItem

`func (o *ViewLockdown) GetItem() ViewRelationship`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *ViewLockdown) GetItemOk() (*ViewRelationship, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *ViewLockdown) SetItem(v ViewRelationship)`

SetItem sets Item field to given value.

### HasItem

`func (o *ViewLockdown) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetItemID

`func (o *ViewLockdown) GetItemID() int32`

GetItemID returns the ItemID field if non-nil, zero value otherwise.

### GetItemIDOk

`func (o *ViewLockdown) GetItemIDOk() (*int32, bool)`

GetItemIDOk returns a tuple with the ItemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemID

`func (o *ViewLockdown) SetItemID(v int32)`

SetItemID sets ItemID field to given value.

### HasItemID

`func (o *ViewLockdown) HasItemID() bool`

HasItemID returns a boolean if a field has been set.

### GetItemType

`func (o *ViewLockdown) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *ViewLockdown) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *ViewLockdown) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *ViewLockdown) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ViewLockdown) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ViewLockdown) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ViewLockdown) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ViewLockdown) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUser

`func (o *ViewLockdown) GetUser() ViewRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ViewLockdown) GetUserOk() (*ViewRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ViewLockdown) SetUser(v ViewRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *ViewLockdown) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserID

`func (o *ViewLockdown) GetUserID() int32`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *ViewLockdown) GetUserIDOk() (*int32, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *ViewLockdown) SetUserID(v int32)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *ViewLockdown) HasUserID() bool`

HasUserID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


