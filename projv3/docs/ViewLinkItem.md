# ViewLinkItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**CategoryId** | Pointer to **int32** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **int32** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**DeletedBy** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ForceNewWindow** | Pointer to **int32** |  | [optional] 
**Height** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**IsPrivate** | Pointer to **bool** |  | [optional] 
**Project** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **int32** |  | [optional] 
**Width** | Pointer to **int32** |  | [optional] 

## Methods

### NewViewLinkItem

`func NewViewLinkItem() *ViewLinkItem`

NewViewLinkItem instantiates a new ViewLinkItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewLinkItemWithDefaults

`func NewViewLinkItemWithDefaults() *ViewLinkItem`

NewViewLinkItemWithDefaults instantiates a new ViewLinkItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *ViewLinkItem) GetCategory() ViewRelationship`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ViewLinkItem) GetCategoryOk() (*ViewRelationship, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ViewLinkItem) SetCategory(v ViewRelationship)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ViewLinkItem) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCategoryId

`func (o *ViewLinkItem) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *ViewLinkItem) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *ViewLinkItem) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *ViewLinkItem) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetCode

`func (o *ViewLinkItem) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ViewLinkItem) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ViewLinkItem) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ViewLinkItem) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ViewLinkItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ViewLinkItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ViewLinkItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ViewLinkItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ViewLinkItem) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ViewLinkItem) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ViewLinkItem) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ViewLinkItem) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDeleted

`func (o *ViewLinkItem) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *ViewLinkItem) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *ViewLinkItem) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *ViewLinkItem) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ViewLinkItem) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ViewLinkItem) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ViewLinkItem) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ViewLinkItem) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeletedBy

`func (o *ViewLinkItem) GetDeletedBy() int32`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *ViewLinkItem) GetDeletedByOk() (*int32, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *ViewLinkItem) SetDeletedBy(v int32)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *ViewLinkItem) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetDescription

`func (o *ViewLinkItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ViewLinkItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ViewLinkItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ViewLinkItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetForceNewWindow

`func (o *ViewLinkItem) GetForceNewWindow() int32`

GetForceNewWindow returns the ForceNewWindow field if non-nil, zero value otherwise.

### GetForceNewWindowOk

`func (o *ViewLinkItem) GetForceNewWindowOk() (*int32, bool)`

GetForceNewWindowOk returns a tuple with the ForceNewWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceNewWindow

`func (o *ViewLinkItem) SetForceNewWindow(v int32)`

SetForceNewWindow sets ForceNewWindow field to given value.

### HasForceNewWindow

`func (o *ViewLinkItem) HasForceNewWindow() bool`

HasForceNewWindow returns a boolean if a field has been set.

### GetHeight

`func (o *ViewLinkItem) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ViewLinkItem) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ViewLinkItem) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ViewLinkItem) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *ViewLinkItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewLinkItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewLinkItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewLinkItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsPrivate

`func (o *ViewLinkItem) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *ViewLinkItem) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *ViewLinkItem) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *ViewLinkItem) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetProject

`func (o *ViewLinkItem) GetProject() ViewRelationship`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ViewLinkItem) GetProjectOk() (*ViewRelationship, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ViewLinkItem) SetProject(v ViewRelationship)`

SetProject sets Project field to given value.

### HasProject

`func (o *ViewLinkItem) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectId

`func (o *ViewLinkItem) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ViewLinkItem) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ViewLinkItem) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ViewLinkItem) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProvider

`func (o *ViewLinkItem) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ViewLinkItem) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ViewLinkItem) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ViewLinkItem) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetTitle

`func (o *ViewLinkItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ViewLinkItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ViewLinkItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ViewLinkItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ViewLinkItem) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ViewLinkItem) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ViewLinkItem) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ViewLinkItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ViewLinkItem) GetUpdatedBy() int32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ViewLinkItem) GetUpdatedByOk() (*int32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ViewLinkItem) SetUpdatedBy(v int32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ViewLinkItem) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetWidth

`func (o *ViewLinkItem) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ViewLinkItem) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ViewLinkItem) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ViewLinkItem) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


