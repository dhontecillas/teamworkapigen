# TwdatausersFilterResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateCreated** | Pointer to **string** |  | [optional] 
**DateDeleted** | Pointer to **string** |  | [optional] 
**DateUpdated** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayOrder** | Pointer to **int32** |  | [optional] 
**Fulldata** | Pointer to **map[string]interface{}** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**IncludesSort** | Pointer to **bool** |  | [optional] 
**IsProjectSpecific** | Pointer to **bool** |  | [optional] 
**IsTemporary** | Pointer to **bool** |  | [optional] 
**Parameters** | Pointer to **map[string]string** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**Section** | Pointer to **string** |  | [optional] 
**ShareLink** | Pointer to **string** |  | [optional] 
**Shared** | Pointer to **bool** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 

## Methods

### NewTwdatausersFilterResponse

`func NewTwdatausersFilterResponse() *TwdatausersFilterResponse`

NewTwdatausersFilterResponse instantiates a new TwdatausersFilterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTwdatausersFilterResponseWithDefaults

`func NewTwdatausersFilterResponseWithDefaults() *TwdatausersFilterResponse`

NewTwdatausersFilterResponseWithDefaults instantiates a new TwdatausersFilterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateCreated

`func (o *TwdatausersFilterResponse) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *TwdatausersFilterResponse) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *TwdatausersFilterResponse) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *TwdatausersFilterResponse) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateDeleted

`func (o *TwdatausersFilterResponse) GetDateDeleted() string`

GetDateDeleted returns the DateDeleted field if non-nil, zero value otherwise.

### GetDateDeletedOk

`func (o *TwdatausersFilterResponse) GetDateDeletedOk() (*string, bool)`

GetDateDeletedOk returns a tuple with the DateDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeleted

`func (o *TwdatausersFilterResponse) SetDateDeleted(v string)`

SetDateDeleted sets DateDeleted field to given value.

### HasDateDeleted

`func (o *TwdatausersFilterResponse) HasDateDeleted() bool`

HasDateDeleted returns a boolean if a field has been set.

### GetDateUpdated

`func (o *TwdatausersFilterResponse) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *TwdatausersFilterResponse) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *TwdatausersFilterResponse) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *TwdatausersFilterResponse) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### GetDeleted

`func (o *TwdatausersFilterResponse) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *TwdatausersFilterResponse) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *TwdatausersFilterResponse) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *TwdatausersFilterResponse) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDescription

`func (o *TwdatausersFilterResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TwdatausersFilterResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TwdatausersFilterResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TwdatausersFilterResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *TwdatausersFilterResponse) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *TwdatausersFilterResponse) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *TwdatausersFilterResponse) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *TwdatausersFilterResponse) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetFulldata

`func (o *TwdatausersFilterResponse) GetFulldata() map[string]interface{}`

GetFulldata returns the Fulldata field if non-nil, zero value otherwise.

### GetFulldataOk

`func (o *TwdatausersFilterResponse) GetFulldataOk() (*map[string]interface{}, bool)`

GetFulldataOk returns a tuple with the Fulldata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulldata

`func (o *TwdatausersFilterResponse) SetFulldata(v map[string]interface{})`

SetFulldata sets Fulldata field to given value.

### HasFulldata

`func (o *TwdatausersFilterResponse) HasFulldata() bool`

HasFulldata returns a boolean if a field has been set.

### GetId

`func (o *TwdatausersFilterResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TwdatausersFilterResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TwdatausersFilterResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TwdatausersFilterResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncludesSort

`func (o *TwdatausersFilterResponse) GetIncludesSort() bool`

GetIncludesSort returns the IncludesSort field if non-nil, zero value otherwise.

### GetIncludesSortOk

`func (o *TwdatausersFilterResponse) GetIncludesSortOk() (*bool, bool)`

GetIncludesSortOk returns a tuple with the IncludesSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludesSort

`func (o *TwdatausersFilterResponse) SetIncludesSort(v bool)`

SetIncludesSort sets IncludesSort field to given value.

### HasIncludesSort

`func (o *TwdatausersFilterResponse) HasIncludesSort() bool`

HasIncludesSort returns a boolean if a field has been set.

### GetIsProjectSpecific

`func (o *TwdatausersFilterResponse) GetIsProjectSpecific() bool`

GetIsProjectSpecific returns the IsProjectSpecific field if non-nil, zero value otherwise.

### GetIsProjectSpecificOk

`func (o *TwdatausersFilterResponse) GetIsProjectSpecificOk() (*bool, bool)`

GetIsProjectSpecificOk returns a tuple with the IsProjectSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProjectSpecific

`func (o *TwdatausersFilterResponse) SetIsProjectSpecific(v bool)`

SetIsProjectSpecific sets IsProjectSpecific field to given value.

### HasIsProjectSpecific

`func (o *TwdatausersFilterResponse) HasIsProjectSpecific() bool`

HasIsProjectSpecific returns a boolean if a field has been set.

### GetIsTemporary

`func (o *TwdatausersFilterResponse) GetIsTemporary() bool`

GetIsTemporary returns the IsTemporary field if non-nil, zero value otherwise.

### GetIsTemporaryOk

`func (o *TwdatausersFilterResponse) GetIsTemporaryOk() (*bool, bool)`

GetIsTemporaryOk returns a tuple with the IsTemporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemporary

`func (o *TwdatausersFilterResponse) SetIsTemporary(v bool)`

SetIsTemporary sets IsTemporary field to given value.

### HasIsTemporary

`func (o *TwdatausersFilterResponse) HasIsTemporary() bool`

HasIsTemporary returns a boolean if a field has been set.

### GetParameters

`func (o *TwdatausersFilterResponse) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *TwdatausersFilterResponse) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *TwdatausersFilterResponse) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *TwdatausersFilterResponse) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetProjectId

`func (o *TwdatausersFilterResponse) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TwdatausersFilterResponse) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TwdatausersFilterResponse) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *TwdatausersFilterResponse) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetSection

`func (o *TwdatausersFilterResponse) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *TwdatausersFilterResponse) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *TwdatausersFilterResponse) SetSection(v string)`

SetSection sets Section field to given value.

### HasSection

`func (o *TwdatausersFilterResponse) HasSection() bool`

HasSection returns a boolean if a field has been set.

### GetShareLink

`func (o *TwdatausersFilterResponse) GetShareLink() string`

GetShareLink returns the ShareLink field if non-nil, zero value otherwise.

### GetShareLinkOk

`func (o *TwdatausersFilterResponse) GetShareLinkOk() (*string, bool)`

GetShareLinkOk returns a tuple with the ShareLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareLink

`func (o *TwdatausersFilterResponse) SetShareLink(v string)`

SetShareLink sets ShareLink field to given value.

### HasShareLink

`func (o *TwdatausersFilterResponse) HasShareLink() bool`

HasShareLink returns a boolean if a field has been set.

### GetShared

`func (o *TwdatausersFilterResponse) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *TwdatausersFilterResponse) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *TwdatausersFilterResponse) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *TwdatausersFilterResponse) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetTitle

`func (o *TwdatausersFilterResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TwdatausersFilterResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TwdatausersFilterResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TwdatausersFilterResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUserId

`func (o *TwdatausersFilterResponse) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TwdatausersFilterResponse) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TwdatausersFilterResponse) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *TwdatausersFilterResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


