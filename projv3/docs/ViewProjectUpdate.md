# ViewProjectUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **int32** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**DeletedBy** | Pointer to **int32** |  | [optional] 
**Health** | Pointer to **int32** |  | [optional] 
**HealthLabel** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**LikeFromUserIDs** | Pointer to **[]int32** |  | [optional] 
**LikeFromUsers** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**Project** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**Reactions** | Pointer to [**ViewReactionsForObject**](ViewReactionsForObject.md) |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewViewProjectUpdate

`func NewViewProjectUpdate() *ViewProjectUpdate`

NewViewProjectUpdate instantiates a new ViewProjectUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewProjectUpdateWithDefaults

`func NewViewProjectUpdateWithDefaults() *ViewProjectUpdate`

NewViewProjectUpdateWithDefaults instantiates a new ViewProjectUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *ViewProjectUpdate) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ViewProjectUpdate) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ViewProjectUpdate) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *ViewProjectUpdate) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ViewProjectUpdate) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ViewProjectUpdate) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ViewProjectUpdate) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ViewProjectUpdate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ViewProjectUpdate) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ViewProjectUpdate) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ViewProjectUpdate) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ViewProjectUpdate) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDeleted

`func (o *ViewProjectUpdate) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *ViewProjectUpdate) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *ViewProjectUpdate) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *ViewProjectUpdate) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ViewProjectUpdate) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ViewProjectUpdate) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ViewProjectUpdate) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ViewProjectUpdate) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeletedBy

`func (o *ViewProjectUpdate) GetDeletedBy() int32`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *ViewProjectUpdate) GetDeletedByOk() (*int32, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *ViewProjectUpdate) SetDeletedBy(v int32)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *ViewProjectUpdate) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetHealth

`func (o *ViewProjectUpdate) GetHealth() int32`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ViewProjectUpdate) GetHealthOk() (*int32, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ViewProjectUpdate) SetHealth(v int32)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *ViewProjectUpdate) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHealthLabel

`func (o *ViewProjectUpdate) GetHealthLabel() string`

GetHealthLabel returns the HealthLabel field if non-nil, zero value otherwise.

### GetHealthLabelOk

`func (o *ViewProjectUpdate) GetHealthLabelOk() (*string, bool)`

GetHealthLabelOk returns a tuple with the HealthLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthLabel

`func (o *ViewProjectUpdate) SetHealthLabel(v string)`

SetHealthLabel sets HealthLabel field to given value.

### HasHealthLabel

`func (o *ViewProjectUpdate) HasHealthLabel() bool`

HasHealthLabel returns a boolean if a field has been set.

### GetId

`func (o *ViewProjectUpdate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewProjectUpdate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewProjectUpdate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewProjectUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLikeFromUserIDs

`func (o *ViewProjectUpdate) GetLikeFromUserIDs() []int32`

GetLikeFromUserIDs returns the LikeFromUserIDs field if non-nil, zero value otherwise.

### GetLikeFromUserIDsOk

`func (o *ViewProjectUpdate) GetLikeFromUserIDsOk() (*[]int32, bool)`

GetLikeFromUserIDsOk returns a tuple with the LikeFromUserIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikeFromUserIDs

`func (o *ViewProjectUpdate) SetLikeFromUserIDs(v []int32)`

SetLikeFromUserIDs sets LikeFromUserIDs field to given value.

### HasLikeFromUserIDs

`func (o *ViewProjectUpdate) HasLikeFromUserIDs() bool`

HasLikeFromUserIDs returns a boolean if a field has been set.

### GetLikeFromUsers

`func (o *ViewProjectUpdate) GetLikeFromUsers() []ViewRelationship`

GetLikeFromUsers returns the LikeFromUsers field if non-nil, zero value otherwise.

### GetLikeFromUsersOk

`func (o *ViewProjectUpdate) GetLikeFromUsersOk() (*[]ViewRelationship, bool)`

GetLikeFromUsersOk returns a tuple with the LikeFromUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikeFromUsers

`func (o *ViewProjectUpdate) SetLikeFromUsers(v []ViewRelationship)`

SetLikeFromUsers sets LikeFromUsers field to given value.

### HasLikeFromUsers

`func (o *ViewProjectUpdate) HasLikeFromUsers() bool`

HasLikeFromUsers returns a boolean if a field has been set.

### GetProject

`func (o *ViewProjectUpdate) GetProject() ViewRelationship`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ViewProjectUpdate) GetProjectOk() (*ViewRelationship, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ViewProjectUpdate) SetProject(v ViewRelationship)`

SetProject sets Project field to given value.

### HasProject

`func (o *ViewProjectUpdate) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectId

`func (o *ViewProjectUpdate) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ViewProjectUpdate) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ViewProjectUpdate) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ViewProjectUpdate) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetReactions

`func (o *ViewProjectUpdate) GetReactions() ViewReactionsForObject`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *ViewProjectUpdate) GetReactionsOk() (*ViewReactionsForObject, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *ViewProjectUpdate) SetReactions(v ViewReactionsForObject)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *ViewProjectUpdate) HasReactions() bool`

HasReactions returns a boolean if a field has been set.

### GetText

`func (o *ViewProjectUpdate) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ViewProjectUpdate) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ViewProjectUpdate) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ViewProjectUpdate) HasText() bool`

HasText returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ViewProjectUpdate) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ViewProjectUpdate) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ViewProjectUpdate) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ViewProjectUpdate) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


