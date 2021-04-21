# ViewCustomField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **int32** |  | [optional] 
**CreatedByUserId** | Pointer to **int32** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**DeletedBy** | Pointer to **int32** |  | [optional] 
**DeletedByUserId** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Entity** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**IsPrivate** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Options** | Pointer to **map[string]interface{}** |  | [optional] 
**Project** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **int32** |  | [optional] 
**UpdatedByUserId** | Pointer to **int32** |  | [optional] 
**Visibilities** | Pointer to **string** |  | [optional] 

## Methods

### NewViewCustomField

`func NewViewCustomField() *ViewCustomField`

NewViewCustomField instantiates a new ViewCustomField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewCustomFieldWithDefaults

`func NewViewCustomFieldWithDefaults() *ViewCustomField`

NewViewCustomFieldWithDefaults instantiates a new ViewCustomField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ViewCustomField) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ViewCustomField) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ViewCustomField) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ViewCustomField) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ViewCustomField) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ViewCustomField) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ViewCustomField) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ViewCustomField) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *ViewCustomField) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *ViewCustomField) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *ViewCustomField) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *ViewCustomField) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetDeleted

`func (o *ViewCustomField) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *ViewCustomField) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *ViewCustomField) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *ViewCustomField) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ViewCustomField) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ViewCustomField) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ViewCustomField) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ViewCustomField) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeletedBy

`func (o *ViewCustomField) GetDeletedBy() int32`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *ViewCustomField) GetDeletedByOk() (*int32, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *ViewCustomField) SetDeletedBy(v int32)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *ViewCustomField) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetDeletedByUserId

`func (o *ViewCustomField) GetDeletedByUserId() int32`

GetDeletedByUserId returns the DeletedByUserId field if non-nil, zero value otherwise.

### GetDeletedByUserIdOk

`func (o *ViewCustomField) GetDeletedByUserIdOk() (*int32, bool)`

GetDeletedByUserIdOk returns a tuple with the DeletedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedByUserId

`func (o *ViewCustomField) SetDeletedByUserId(v int32)`

SetDeletedByUserId sets DeletedByUserId field to given value.

### HasDeletedByUserId

`func (o *ViewCustomField) HasDeletedByUserId() bool`

HasDeletedByUserId returns a boolean if a field has been set.

### GetDescription

`func (o *ViewCustomField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ViewCustomField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ViewCustomField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ViewCustomField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntity

`func (o *ViewCustomField) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *ViewCustomField) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *ViewCustomField) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *ViewCustomField) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetId

`func (o *ViewCustomField) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewCustomField) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewCustomField) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewCustomField) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsPrivate

`func (o *ViewCustomField) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *ViewCustomField) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *ViewCustomField) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *ViewCustomField) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetName

`func (o *ViewCustomField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ViewCustomField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ViewCustomField) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ViewCustomField) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *ViewCustomField) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ViewCustomField) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ViewCustomField) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ViewCustomField) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetProject

`func (o *ViewCustomField) GetProject() ViewRelationship`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ViewCustomField) GetProjectOk() (*ViewRelationship, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ViewCustomField) SetProject(v ViewRelationship)`

SetProject sets Project field to given value.

### HasProject

`func (o *ViewCustomField) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectId

`func (o *ViewCustomField) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ViewCustomField) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ViewCustomField) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ViewCustomField) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetRequired

`func (o *ViewCustomField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ViewCustomField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ViewCustomField) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ViewCustomField) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetType

`func (o *ViewCustomField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ViewCustomField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ViewCustomField) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ViewCustomField) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ViewCustomField) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ViewCustomField) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ViewCustomField) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ViewCustomField) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ViewCustomField) GetUpdatedBy() int32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ViewCustomField) GetUpdatedByOk() (*int32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ViewCustomField) SetUpdatedBy(v int32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ViewCustomField) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedByUserId

`func (o *ViewCustomField) GetUpdatedByUserId() int32`

GetUpdatedByUserId returns the UpdatedByUserId field if non-nil, zero value otherwise.

### GetUpdatedByUserIdOk

`func (o *ViewCustomField) GetUpdatedByUserIdOk() (*int32, bool)`

GetUpdatedByUserIdOk returns a tuple with the UpdatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUserId

`func (o *ViewCustomField) SetUpdatedByUserId(v int32)`

SetUpdatedByUserId sets UpdatedByUserId field to given value.

### HasUpdatedByUserId

`func (o *ViewCustomField) HasUpdatedByUserId() bool`

HasUpdatedByUserId returns a boolean if a field has been set.

### GetVisibilities

`func (o *ViewCustomField) GetVisibilities() string`

GetVisibilities returns the Visibilities field if non-nil, zero value otherwise.

### GetVisibilitiesOk

`func (o *ViewCustomField) GetVisibilitiesOk() (*string, bool)`

GetVisibilitiesOk returns a tuple with the Visibilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibilities

`func (o *ViewCustomField) SetVisibilities(v string)`

SetVisibilities sets Visibilities field to given value.

### HasVisibilities

`func (o *ViewCustomField) HasVisibilities() bool`

HasVisibilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


