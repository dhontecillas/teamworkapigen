# ViewBoardColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DefaultTasklist** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**DisplayOrder** | Pointer to **int32** |  | [optional] 
**HasTriggers** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Project** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**Settings** | Pointer to [**ViewBoardColumnSettings**](ViewBoardColumnSettings.md) |  | [optional] 
**Sort** | Pointer to **string** |  | [optional] 
**SortOrder** | Pointer to **string** |  | [optional] 
**Stats** | Pointer to [**ViewColumnStats**](ViewColumnStats.md) |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewViewBoardColumn

`func NewViewBoardColumn() *ViewBoardColumn`

NewViewBoardColumn instantiates a new ViewBoardColumn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewBoardColumnWithDefaults

`func NewViewBoardColumnWithDefaults() *ViewBoardColumn`

NewViewBoardColumnWithDefaults instantiates a new ViewBoardColumn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *ViewBoardColumn) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ViewBoardColumn) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ViewBoardColumn) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *ViewBoardColumn) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ViewBoardColumn) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ViewBoardColumn) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ViewBoardColumn) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ViewBoardColumn) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefaultTasklist

`func (o *ViewBoardColumn) GetDefaultTasklist() ViewRelationship`

GetDefaultTasklist returns the DefaultTasklist field if non-nil, zero value otherwise.

### GetDefaultTasklistOk

`func (o *ViewBoardColumn) GetDefaultTasklistOk() (*ViewRelationship, bool)`

GetDefaultTasklistOk returns a tuple with the DefaultTasklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTasklist

`func (o *ViewBoardColumn) SetDefaultTasklist(v ViewRelationship)`

SetDefaultTasklist sets DefaultTasklist field to given value.

### HasDefaultTasklist

`func (o *ViewBoardColumn) HasDefaultTasklist() bool`

HasDefaultTasklist returns a boolean if a field has been set.

### GetDeleted

`func (o *ViewBoardColumn) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *ViewBoardColumn) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *ViewBoardColumn) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *ViewBoardColumn) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ViewBoardColumn) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ViewBoardColumn) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ViewBoardColumn) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ViewBoardColumn) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *ViewBoardColumn) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *ViewBoardColumn) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *ViewBoardColumn) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *ViewBoardColumn) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetHasTriggers

`func (o *ViewBoardColumn) GetHasTriggers() bool`

GetHasTriggers returns the HasTriggers field if non-nil, zero value otherwise.

### GetHasTriggersOk

`func (o *ViewBoardColumn) GetHasTriggersOk() (*bool, bool)`

GetHasTriggersOk returns a tuple with the HasTriggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTriggers

`func (o *ViewBoardColumn) SetHasTriggers(v bool)`

SetHasTriggers sets HasTriggers field to given value.

### HasHasTriggers

`func (o *ViewBoardColumn) HasHasTriggers() bool`

HasHasTriggers returns a boolean if a field has been set.

### GetId

`func (o *ViewBoardColumn) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewBoardColumn) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewBoardColumn) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewBoardColumn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ViewBoardColumn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ViewBoardColumn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ViewBoardColumn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ViewBoardColumn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProject

`func (o *ViewBoardColumn) GetProject() ViewRelationship`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ViewBoardColumn) GetProjectOk() (*ViewRelationship, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ViewBoardColumn) SetProject(v ViewRelationship)`

SetProject sets Project field to given value.

### HasProject

`func (o *ViewBoardColumn) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetSettings

`func (o *ViewBoardColumn) GetSettings() ViewBoardColumnSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ViewBoardColumn) GetSettingsOk() (*ViewBoardColumnSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ViewBoardColumn) SetSettings(v ViewBoardColumnSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *ViewBoardColumn) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetSort

`func (o *ViewBoardColumn) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *ViewBoardColumn) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *ViewBoardColumn) SetSort(v string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *ViewBoardColumn) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetSortOrder

`func (o *ViewBoardColumn) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ViewBoardColumn) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ViewBoardColumn) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *ViewBoardColumn) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetStats

`func (o *ViewBoardColumn) GetStats() ViewColumnStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ViewBoardColumn) GetStatsOk() (*ViewColumnStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ViewBoardColumn) SetStats(v ViewColumnStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ViewBoardColumn) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ViewBoardColumn) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ViewBoardColumn) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ViewBoardColumn) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ViewBoardColumn) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


