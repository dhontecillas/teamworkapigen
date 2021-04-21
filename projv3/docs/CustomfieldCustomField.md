# CustomfieldCustomField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Entity** | Pointer to **string** | using a verbose name to avoid conflict | [optional] 
**IsPrivate** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Options** | Pointer to **map[string]interface{}** |  | [optional] 
**Privacy** | Pointer to [**PayloadUserGroups**](PayloadUserGroups.md) |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Visibilities** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomfieldCustomField

`func NewCustomfieldCustomField() *CustomfieldCustomField`

NewCustomfieldCustomField instantiates a new CustomfieldCustomField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomfieldCustomFieldWithDefaults

`func NewCustomfieldCustomFieldWithDefaults() *CustomfieldCustomField`

NewCustomfieldCustomFieldWithDefaults instantiates a new CustomfieldCustomField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CustomfieldCustomField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomfieldCustomField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomfieldCustomField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomfieldCustomField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntity

`func (o *CustomfieldCustomField) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *CustomfieldCustomField) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *CustomfieldCustomField) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *CustomfieldCustomField) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetIsPrivate

`func (o *CustomfieldCustomField) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *CustomfieldCustomField) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *CustomfieldCustomField) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *CustomfieldCustomField) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetName

`func (o *CustomfieldCustomField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomfieldCustomField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomfieldCustomField) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomfieldCustomField) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *CustomfieldCustomField) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CustomfieldCustomField) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CustomfieldCustomField) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CustomfieldCustomField) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPrivacy

`func (o *CustomfieldCustomField) GetPrivacy() PayloadUserGroups`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *CustomfieldCustomField) GetPrivacyOk() (*PayloadUserGroups, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *CustomfieldCustomField) SetPrivacy(v PayloadUserGroups)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *CustomfieldCustomField) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetProjectId

`func (o *CustomfieldCustomField) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CustomfieldCustomField) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CustomfieldCustomField) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CustomfieldCustomField) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetRequired

`func (o *CustomfieldCustomField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *CustomfieldCustomField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *CustomfieldCustomField) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *CustomfieldCustomField) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetType

`func (o *CustomfieldCustomField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomfieldCustomField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomfieldCustomField) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomfieldCustomField) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVisibilities

`func (o *CustomfieldCustomField) GetVisibilities() string`

GetVisibilities returns the Visibilities field if non-nil, zero value otherwise.

### GetVisibilitiesOk

`func (o *CustomfieldCustomField) GetVisibilitiesOk() (*string, bool)`

GetVisibilitiesOk returns a tuple with the Visibilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibilities

`func (o *CustomfieldCustomField) SetVisibilities(v string)`

SetVisibilities sets Visibilities field to given value.

### HasVisibilities

`func (o *CustomfieldCustomField) HasVisibilities() bool`

HasVisibilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


