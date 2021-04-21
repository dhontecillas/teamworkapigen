# InlineResponse20045Card

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | Pointer to **bool** |  | [optional] 
**BoardId** | Pointer to **string** |  | [optional] 
**CanEdit** | Pointer to **bool** |  | [optional] 
**Category** | Pointer to [**InlineResponse2002Column**](InlineResponse2002Column.md) |  | [optional] 
**ColumnId** | Pointer to **string** |  | [optional] 
**Company** | Pointer to [**InlineResponse2002People12345Company**](InlineResponse2002People12345Company.md) |  | [optional] 
**DateCreated** | Pointer to **string** |  | [optional] 
**DateUpdated** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**DeletedDate** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayOrder** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**InlineResponse20045CardOwner**](InlineResponse20045CardOwner.md) |  | [optional] 
**Progress** | Pointer to **map[string]interface{}** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**Starred** | Pointer to **bool** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Update** | Pointer to [**InlineResponse20045CardUpdate**](InlineResponse20045CardUpdate.md) |  | [optional] 

## Methods

### NewInlineResponse20045Card

`func NewInlineResponse20045Card() *InlineResponse20045Card`

NewInlineResponse20045Card instantiates a new InlineResponse20045Card object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20045CardWithDefaults

`func NewInlineResponse20045CardWithDefaults() *InlineResponse20045Card`

NewInlineResponse20045CardWithDefaults instantiates a new InlineResponse20045Card object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *InlineResponse20045Card) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *InlineResponse20045Card) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *InlineResponse20045Card) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *InlineResponse20045Card) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetBoardId

`func (o *InlineResponse20045Card) GetBoardId() string`

GetBoardId returns the BoardId field if non-nil, zero value otherwise.

### GetBoardIdOk

`func (o *InlineResponse20045Card) GetBoardIdOk() (*string, bool)`

GetBoardIdOk returns a tuple with the BoardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardId

`func (o *InlineResponse20045Card) SetBoardId(v string)`

SetBoardId sets BoardId field to given value.

### HasBoardId

`func (o *InlineResponse20045Card) HasBoardId() bool`

HasBoardId returns a boolean if a field has been set.

### GetCanEdit

`func (o *InlineResponse20045Card) GetCanEdit() bool`

GetCanEdit returns the CanEdit field if non-nil, zero value otherwise.

### GetCanEditOk

`func (o *InlineResponse20045Card) GetCanEditOk() (*bool, bool)`

GetCanEditOk returns a tuple with the CanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEdit

`func (o *InlineResponse20045Card) SetCanEdit(v bool)`

SetCanEdit sets CanEdit field to given value.

### HasCanEdit

`func (o *InlineResponse20045Card) HasCanEdit() bool`

HasCanEdit returns a boolean if a field has been set.

### GetCategory

`func (o *InlineResponse20045Card) GetCategory() InlineResponse2002Column`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InlineResponse20045Card) GetCategoryOk() (*InlineResponse2002Column, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InlineResponse20045Card) SetCategory(v InlineResponse2002Column)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *InlineResponse20045Card) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetColumnId

`func (o *InlineResponse20045Card) GetColumnId() string`

GetColumnId returns the ColumnId field if non-nil, zero value otherwise.

### GetColumnIdOk

`func (o *InlineResponse20045Card) GetColumnIdOk() (*string, bool)`

GetColumnIdOk returns a tuple with the ColumnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnId

`func (o *InlineResponse20045Card) SetColumnId(v string)`

SetColumnId sets ColumnId field to given value.

### HasColumnId

`func (o *InlineResponse20045Card) HasColumnId() bool`

HasColumnId returns a boolean if a field has been set.

### GetCompany

`func (o *InlineResponse20045Card) GetCompany() InlineResponse2002People12345Company`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *InlineResponse20045Card) GetCompanyOk() (*InlineResponse2002People12345Company, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *InlineResponse20045Card) SetCompany(v InlineResponse2002People12345Company)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *InlineResponse20045Card) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetDateCreated

`func (o *InlineResponse20045Card) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *InlineResponse20045Card) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *InlineResponse20045Card) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *InlineResponse20045Card) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateUpdated

`func (o *InlineResponse20045Card) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *InlineResponse20045Card) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *InlineResponse20045Card) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *InlineResponse20045Card) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### GetDeleted

`func (o *InlineResponse20045Card) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *InlineResponse20045Card) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *InlineResponse20045Card) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *InlineResponse20045Card) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDeletedDate

`func (o *InlineResponse20045Card) GetDeletedDate() string`

GetDeletedDate returns the DeletedDate field if non-nil, zero value otherwise.

### GetDeletedDateOk

`func (o *InlineResponse20045Card) GetDeletedDateOk() (*string, bool)`

GetDeletedDateOk returns a tuple with the DeletedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDate

`func (o *InlineResponse20045Card) SetDeletedDate(v string)`

SetDeletedDate sets DeletedDate field to given value.

### HasDeletedDate

`func (o *InlineResponse20045Card) HasDeletedDate() bool`

HasDeletedDate returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse20045Card) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse20045Card) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse20045Card) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse20045Card) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *InlineResponse20045Card) GetDisplayOrder() string`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *InlineResponse20045Card) GetDisplayOrderOk() (*string, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *InlineResponse20045Card) SetDisplayOrder(v string)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *InlineResponse20045Card) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetEndDate

`func (o *InlineResponse20045Card) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *InlineResponse20045Card) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *InlineResponse20045Card) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *InlineResponse20045Card) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetId

`func (o *InlineResponse20045Card) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20045Card) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20045Card) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20045Card) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20045Card) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20045Card) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20045Card) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20045Card) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *InlineResponse20045Card) GetOwner() InlineResponse20045CardOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *InlineResponse20045Card) GetOwnerOk() (*InlineResponse20045CardOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *InlineResponse20045Card) SetOwner(v InlineResponse20045CardOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *InlineResponse20045Card) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetProgress

`func (o *InlineResponse20045Card) GetProgress() map[string]interface{}`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *InlineResponse20045Card) GetProgressOk() (*map[string]interface{}, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *InlineResponse20045Card) SetProgress(v map[string]interface{})`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *InlineResponse20045Card) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetProjectId

`func (o *InlineResponse20045Card) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *InlineResponse20045Card) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *InlineResponse20045Card) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *InlineResponse20045Card) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetStarred

`func (o *InlineResponse20045Card) GetStarred() bool`

GetStarred returns the Starred field if non-nil, zero value otherwise.

### GetStarredOk

`func (o *InlineResponse20045Card) GetStarredOk() (*bool, bool)`

GetStarredOk returns a tuple with the Starred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarred

`func (o *InlineResponse20045Card) SetStarred(v bool)`

SetStarred sets Starred field to given value.

### HasStarred

`func (o *InlineResponse20045Card) HasStarred() bool`

HasStarred returns a boolean if a field has been set.

### GetStartDate

`func (o *InlineResponse20045Card) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InlineResponse20045Card) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InlineResponse20045Card) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InlineResponse20045Card) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20045Card) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20045Card) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20045Card) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20045Card) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse20045Card) GetTags() []map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse20045Card) GetTagsOk() (*[]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse20045Card) SetTags(v []map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse20045Card) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdate

`func (o *InlineResponse20045Card) GetUpdate() InlineResponse20045CardUpdate`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *InlineResponse20045Card) GetUpdateOk() (*InlineResponse20045CardUpdate, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *InlineResponse20045Card) SetUpdate(v InlineResponse20045CardUpdate)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *InlineResponse20045Card) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


