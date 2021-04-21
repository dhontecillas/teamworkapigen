# ViewTaskCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | Pointer to **bool** |  | [optional] 
**ArchivedAt** | Pointer to **string** |  | [optional] 
**ArchivedBy** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**CreateBy** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DeleteBy** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**DisplayOrder** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Visible** | Pointer to **bool** |  | [optional] 

## Methods

### NewViewTaskCard

`func NewViewTaskCard() *ViewTaskCard`

NewViewTaskCard instantiates a new ViewTaskCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewTaskCardWithDefaults

`func NewViewTaskCardWithDefaults() *ViewTaskCard`

NewViewTaskCardWithDefaults instantiates a new ViewTaskCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *ViewTaskCard) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ViewTaskCard) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ViewTaskCard) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *ViewTaskCard) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetArchivedAt

`func (o *ViewTaskCard) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *ViewTaskCard) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *ViewTaskCard) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *ViewTaskCard) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetArchivedBy

`func (o *ViewTaskCard) GetArchivedBy() ViewRelationship`

GetArchivedBy returns the ArchivedBy field if non-nil, zero value otherwise.

### GetArchivedByOk

`func (o *ViewTaskCard) GetArchivedByOk() (*ViewRelationship, bool)`

GetArchivedByOk returns a tuple with the ArchivedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedBy

`func (o *ViewTaskCard) SetArchivedBy(v ViewRelationship)`

SetArchivedBy sets ArchivedBy field to given value.

### HasArchivedBy

`func (o *ViewTaskCard) HasArchivedBy() bool`

HasArchivedBy returns a boolean if a field has been set.

### GetCreateBy

`func (o *ViewTaskCard) GetCreateBy() ViewRelationship`

GetCreateBy returns the CreateBy field if non-nil, zero value otherwise.

### GetCreateByOk

`func (o *ViewTaskCard) GetCreateByOk() (*ViewRelationship, bool)`

GetCreateByOk returns a tuple with the CreateBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBy

`func (o *ViewTaskCard) SetCreateBy(v ViewRelationship)`

SetCreateBy sets CreateBy field to given value.

### HasCreateBy

`func (o *ViewTaskCard) HasCreateBy() bool`

HasCreateBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ViewTaskCard) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ViewTaskCard) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ViewTaskCard) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ViewTaskCard) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeleteBy

`func (o *ViewTaskCard) GetDeleteBy() ViewRelationship`

GetDeleteBy returns the DeleteBy field if non-nil, zero value otherwise.

### GetDeleteByOk

`func (o *ViewTaskCard) GetDeleteByOk() (*ViewRelationship, bool)`

GetDeleteByOk returns a tuple with the DeleteBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteBy

`func (o *ViewTaskCard) SetDeleteBy(v ViewRelationship)`

SetDeleteBy sets DeleteBy field to given value.

### HasDeleteBy

`func (o *ViewTaskCard) HasDeleteBy() bool`

HasDeleteBy returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ViewTaskCard) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ViewTaskCard) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ViewTaskCard) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ViewTaskCard) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *ViewTaskCard) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *ViewTaskCard) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *ViewTaskCard) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *ViewTaskCard) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetId

`func (o *ViewTaskCard) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewTaskCard) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewTaskCard) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewTaskCard) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *ViewTaskCard) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ViewTaskCard) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ViewTaskCard) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ViewTaskCard) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ViewTaskCard) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ViewTaskCard) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ViewTaskCard) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ViewTaskCard) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVisible

`func (o *ViewTaskCard) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *ViewTaskCard) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *ViewTaskCard) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *ViewTaskCard) HasVisible() bool`

HasVisible returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


