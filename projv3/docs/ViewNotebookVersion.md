# ViewNotebookVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentHTML** | Pointer to **string** |  | [optional] 
**Contents** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **int32** |  | [optional] 
**CreatedByUserID** | Pointer to **int32** |  | [optional] 
**DateCreated** | Pointer to **string** |  | [optional] 
**DateUpdated** | Pointer to **string** |  | [optional] 
**Notebook** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**NotebookId** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**VersionId** | Pointer to **int32** |  | [optional] 
**VersionNumber** | Pointer to **int32** |  | [optional] 

## Methods

### NewViewNotebookVersion

`func NewViewNotebookVersion() *ViewNotebookVersion`

NewViewNotebookVersion instantiates a new ViewNotebookVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewNotebookVersionWithDefaults

`func NewViewNotebookVersionWithDefaults() *ViewNotebookVersion`

NewViewNotebookVersionWithDefaults instantiates a new ViewNotebookVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentHTML

`func (o *ViewNotebookVersion) GetContentHTML() string`

GetContentHTML returns the ContentHTML field if non-nil, zero value otherwise.

### GetContentHTMLOk

`func (o *ViewNotebookVersion) GetContentHTMLOk() (*string, bool)`

GetContentHTMLOk returns a tuple with the ContentHTML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentHTML

`func (o *ViewNotebookVersion) SetContentHTML(v string)`

SetContentHTML sets ContentHTML field to given value.

### HasContentHTML

`func (o *ViewNotebookVersion) HasContentHTML() bool`

HasContentHTML returns a boolean if a field has been set.

### GetContents

`func (o *ViewNotebookVersion) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *ViewNotebookVersion) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *ViewNotebookVersion) SetContents(v string)`

SetContents sets Contents field to given value.

### HasContents

`func (o *ViewNotebookVersion) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ViewNotebookVersion) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ViewNotebookVersion) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ViewNotebookVersion) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ViewNotebookVersion) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ViewNotebookVersion) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ViewNotebookVersion) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ViewNotebookVersion) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ViewNotebookVersion) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUserID

`func (o *ViewNotebookVersion) GetCreatedByUserID() int32`

GetCreatedByUserID returns the CreatedByUserID field if non-nil, zero value otherwise.

### GetCreatedByUserIDOk

`func (o *ViewNotebookVersion) GetCreatedByUserIDOk() (*int32, bool)`

GetCreatedByUserIDOk returns a tuple with the CreatedByUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserID

`func (o *ViewNotebookVersion) SetCreatedByUserID(v int32)`

SetCreatedByUserID sets CreatedByUserID field to given value.

### HasCreatedByUserID

`func (o *ViewNotebookVersion) HasCreatedByUserID() bool`

HasCreatedByUserID returns a boolean if a field has been set.

### GetDateCreated

`func (o *ViewNotebookVersion) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ViewNotebookVersion) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ViewNotebookVersion) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ViewNotebookVersion) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateUpdated

`func (o *ViewNotebookVersion) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ViewNotebookVersion) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ViewNotebookVersion) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ViewNotebookVersion) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### GetNotebook

`func (o *ViewNotebookVersion) GetNotebook() ViewRelationship`

GetNotebook returns the Notebook field if non-nil, zero value otherwise.

### GetNotebookOk

`func (o *ViewNotebookVersion) GetNotebookOk() (*ViewRelationship, bool)`

GetNotebookOk returns a tuple with the Notebook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotebook

`func (o *ViewNotebookVersion) SetNotebook(v ViewRelationship)`

SetNotebook sets Notebook field to given value.

### HasNotebook

`func (o *ViewNotebookVersion) HasNotebook() bool`

HasNotebook returns a boolean if a field has been set.

### GetNotebookId

`func (o *ViewNotebookVersion) GetNotebookId() int32`

GetNotebookId returns the NotebookId field if non-nil, zero value otherwise.

### GetNotebookIdOk

`func (o *ViewNotebookVersion) GetNotebookIdOk() (*int32, bool)`

GetNotebookIdOk returns a tuple with the NotebookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotebookId

`func (o *ViewNotebookVersion) SetNotebookId(v int32)`

SetNotebookId sets NotebookId field to given value.

### HasNotebookId

`func (o *ViewNotebookVersion) HasNotebookId() bool`

HasNotebookId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ViewNotebookVersion) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ViewNotebookVersion) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ViewNotebookVersion) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ViewNotebookVersion) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersionId

`func (o *ViewNotebookVersion) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *ViewNotebookVersion) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *ViewNotebookVersion) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *ViewNotebookVersion) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetVersionNumber

`func (o *ViewNotebookVersion) GetVersionNumber() int32`

GetVersionNumber returns the VersionNumber field if non-nil, zero value otherwise.

### GetVersionNumberOk

`func (o *ViewNotebookVersion) GetVersionNumberOk() (*int32, bool)`

GetVersionNumberOk returns a tuple with the VersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionNumber

`func (o *ViewNotebookVersion) SetVersionNumber(v int32)`

SetVersionNumber sets VersionNumber field to given value.

### HasVersionNumber

`func (o *ViewNotebookVersion) HasVersionNumber() bool`

HasVersionNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


