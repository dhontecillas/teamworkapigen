# ViewCustomFieldProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **int32** |  | [optional] 
**Customfield** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**CustomfieldId** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Project** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**Value** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewViewCustomFieldProject

`func NewViewCustomFieldProject() *ViewCustomFieldProject`

NewViewCustomFieldProject instantiates a new ViewCustomFieldProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewCustomFieldProjectWithDefaults

`func NewViewCustomFieldProjectWithDefaults() *ViewCustomFieldProject`

NewViewCustomFieldProjectWithDefaults instantiates a new ViewCustomFieldProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ViewCustomFieldProject) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ViewCustomFieldProject) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ViewCustomFieldProject) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ViewCustomFieldProject) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ViewCustomFieldProject) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ViewCustomFieldProject) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ViewCustomFieldProject) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ViewCustomFieldProject) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCustomfield

`func (o *ViewCustomFieldProject) GetCustomfield() ViewRelationship`

GetCustomfield returns the Customfield field if non-nil, zero value otherwise.

### GetCustomfieldOk

`func (o *ViewCustomFieldProject) GetCustomfieldOk() (*ViewRelationship, bool)`

GetCustomfieldOk returns a tuple with the Customfield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomfield

`func (o *ViewCustomFieldProject) SetCustomfield(v ViewRelationship)`

SetCustomfield sets Customfield field to given value.

### HasCustomfield

`func (o *ViewCustomFieldProject) HasCustomfield() bool`

HasCustomfield returns a boolean if a field has been set.

### GetCustomfieldId

`func (o *ViewCustomFieldProject) GetCustomfieldId() int32`

GetCustomfieldId returns the CustomfieldId field if non-nil, zero value otherwise.

### GetCustomfieldIdOk

`func (o *ViewCustomFieldProject) GetCustomfieldIdOk() (*int32, bool)`

GetCustomfieldIdOk returns a tuple with the CustomfieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomfieldId

`func (o *ViewCustomFieldProject) SetCustomfieldId(v int32)`

SetCustomfieldId sets CustomfieldId field to given value.

### HasCustomfieldId

`func (o *ViewCustomFieldProject) HasCustomfieldId() bool`

HasCustomfieldId returns a boolean if a field has been set.

### GetId

`func (o *ViewCustomFieldProject) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewCustomFieldProject) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewCustomFieldProject) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewCustomFieldProject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProject

`func (o *ViewCustomFieldProject) GetProject() ViewRelationship`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ViewCustomFieldProject) GetProjectOk() (*ViewRelationship, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ViewCustomFieldProject) SetProject(v ViewRelationship)`

SetProject sets Project field to given value.

### HasProject

`func (o *ViewCustomFieldProject) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectId

`func (o *ViewCustomFieldProject) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ViewCustomFieldProject) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ViewCustomFieldProject) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ViewCustomFieldProject) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetValue

`func (o *ViewCustomFieldProject) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ViewCustomFieldProject) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ViewCustomFieldProject) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *ViewCustomFieldProject) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


