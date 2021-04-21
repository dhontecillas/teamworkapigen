# ViewUserDashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DashboardPanelIds** | Pointer to **[]int32** |  | [optional] 
**DashboardPanels** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**DashboardSettingIds** | Pointer to **[]int32** |  | [optional] 
**DashboardSettings** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayOrder** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**Project** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 

## Methods

### NewViewUserDashboard

`func NewViewUserDashboard() *ViewUserDashboard`

NewViewUserDashboard instantiates a new ViewUserDashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewUserDashboardWithDefaults

`func NewViewUserDashboardWithDefaults() *ViewUserDashboard`

NewViewUserDashboardWithDefaults instantiates a new ViewUserDashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *ViewUserDashboard) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ViewUserDashboard) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ViewUserDashboard) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *ViewUserDashboard) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ViewUserDashboard) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ViewUserDashboard) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ViewUserDashboard) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ViewUserDashboard) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDashboardPanelIds

`func (o *ViewUserDashboard) GetDashboardPanelIds() []int32`

GetDashboardPanelIds returns the DashboardPanelIds field if non-nil, zero value otherwise.

### GetDashboardPanelIdsOk

`func (o *ViewUserDashboard) GetDashboardPanelIdsOk() (*[]int32, bool)`

GetDashboardPanelIdsOk returns a tuple with the DashboardPanelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardPanelIds

`func (o *ViewUserDashboard) SetDashboardPanelIds(v []int32)`

SetDashboardPanelIds sets DashboardPanelIds field to given value.

### HasDashboardPanelIds

`func (o *ViewUserDashboard) HasDashboardPanelIds() bool`

HasDashboardPanelIds returns a boolean if a field has been set.

### GetDashboardPanels

`func (o *ViewUserDashboard) GetDashboardPanels() []ViewRelationship`

GetDashboardPanels returns the DashboardPanels field if non-nil, zero value otherwise.

### GetDashboardPanelsOk

`func (o *ViewUserDashboard) GetDashboardPanelsOk() (*[]ViewRelationship, bool)`

GetDashboardPanelsOk returns a tuple with the DashboardPanels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardPanels

`func (o *ViewUserDashboard) SetDashboardPanels(v []ViewRelationship)`

SetDashboardPanels sets DashboardPanels field to given value.

### HasDashboardPanels

`func (o *ViewUserDashboard) HasDashboardPanels() bool`

HasDashboardPanels returns a boolean if a field has been set.

### GetDashboardSettingIds

`func (o *ViewUserDashboard) GetDashboardSettingIds() []int32`

GetDashboardSettingIds returns the DashboardSettingIds field if non-nil, zero value otherwise.

### GetDashboardSettingIdsOk

`func (o *ViewUserDashboard) GetDashboardSettingIdsOk() (*[]int32, bool)`

GetDashboardSettingIdsOk returns a tuple with the DashboardSettingIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardSettingIds

`func (o *ViewUserDashboard) SetDashboardSettingIds(v []int32)`

SetDashboardSettingIds sets DashboardSettingIds field to given value.

### HasDashboardSettingIds

`func (o *ViewUserDashboard) HasDashboardSettingIds() bool`

HasDashboardSettingIds returns a boolean if a field has been set.

### GetDashboardSettings

`func (o *ViewUserDashboard) GetDashboardSettings() []ViewRelationship`

GetDashboardSettings returns the DashboardSettings field if non-nil, zero value otherwise.

### GetDashboardSettingsOk

`func (o *ViewUserDashboard) GetDashboardSettingsOk() (*[]ViewRelationship, bool)`

GetDashboardSettingsOk returns a tuple with the DashboardSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardSettings

`func (o *ViewUserDashboard) SetDashboardSettings(v []ViewRelationship)`

SetDashboardSettings sets DashboardSettings field to given value.

### HasDashboardSettings

`func (o *ViewUserDashboard) HasDashboardSettings() bool`

HasDashboardSettings returns a boolean if a field has been set.

### GetDeleted

`func (o *ViewUserDashboard) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *ViewUserDashboard) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *ViewUserDashboard) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *ViewUserDashboard) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ViewUserDashboard) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ViewUserDashboard) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ViewUserDashboard) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ViewUserDashboard) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDescription

`func (o *ViewUserDashboard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ViewUserDashboard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ViewUserDashboard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ViewUserDashboard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *ViewUserDashboard) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *ViewUserDashboard) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *ViewUserDashboard) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *ViewUserDashboard) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetId

`func (o *ViewUserDashboard) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewUserDashboard) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewUserDashboard) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewUserDashboard) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *ViewUserDashboard) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ViewUserDashboard) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ViewUserDashboard) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *ViewUserDashboard) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetProject

`func (o *ViewUserDashboard) GetProject() ViewRelationship`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ViewUserDashboard) GetProjectOk() (*ViewRelationship, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ViewUserDashboard) SetProject(v ViewRelationship)`

SetProject sets Project field to given value.

### HasProject

`func (o *ViewUserDashboard) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectId

`func (o *ViewUserDashboard) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ViewUserDashboard) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ViewUserDashboard) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ViewUserDashboard) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetTitle

`func (o *ViewUserDashboard) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ViewUserDashboard) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ViewUserDashboard) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ViewUserDashboard) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ViewUserDashboard) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ViewUserDashboard) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ViewUserDashboard) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ViewUserDashboard) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUser

`func (o *ViewUserDashboard) GetUser() ViewRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ViewUserDashboard) GetUserOk() (*ViewRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ViewUserDashboard) SetUser(v ViewRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *ViewUserDashboard) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserId

`func (o *ViewUserDashboard) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ViewUserDashboard) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ViewUserDashboard) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ViewUserDashboard) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


