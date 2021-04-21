# InlineResponse200101

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audit** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Columns** | Pointer to **map[string]interface{}** |  | [optional] 
**Companies** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedByUserId** | Pointer to **int32** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**People** | Pointer to [**InlineResponse200101People**](InlineResponse200101People.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 
**Tasklists** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewInlineResponse200101

`func NewInlineResponse200101() *InlineResponse200101`

NewInlineResponse200101 instantiates a new InlineResponse200101 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200101WithDefaults

`func NewInlineResponse200101WithDefaults() *InlineResponse200101`

NewInlineResponse200101WithDefaults instantiates a new InlineResponse200101 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *InlineResponse200101) GetAudit() []map[string]interface{}`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *InlineResponse200101) GetAuditOk() (*[]map[string]interface{}, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *InlineResponse200101) SetAudit(v []map[string]interface{})`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *InlineResponse200101) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetColumns

`func (o *InlineResponse200101) GetColumns() map[string]interface{}`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *InlineResponse200101) GetColumnsOk() (*map[string]interface{}, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *InlineResponse200101) SetColumns(v map[string]interface{})`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *InlineResponse200101) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetCompanies

`func (o *InlineResponse200101) GetCompanies() map[string]interface{}`

GetCompanies returns the Companies field if non-nil, zero value otherwise.

### GetCompaniesOk

`func (o *InlineResponse200101) GetCompaniesOk() (*map[string]interface{}, bool)`

GetCompaniesOk returns a tuple with the Companies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanies

`func (o *InlineResponse200101) SetCompanies(v map[string]interface{})`

SetCompanies sets Companies field to given value.

### HasCompanies

`func (o *InlineResponse200101) HasCompanies() bool`

HasCompanies returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *InlineResponse200101) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *InlineResponse200101) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *InlineResponse200101) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *InlineResponse200101) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetCreatedTime

`func (o *InlineResponse200101) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *InlineResponse200101) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *InlineResponse200101) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *InlineResponse200101) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetPeople

`func (o *InlineResponse200101) GetPeople() InlineResponse200101People`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *InlineResponse200101) GetPeopleOk() (*InlineResponse200101People, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *InlineResponse200101) SetPeople(v InlineResponse200101People)`

SetPeople sets People field to given value.

### HasPeople

`func (o *InlineResponse200101) HasPeople() bool`

HasPeople returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200101) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200101) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200101) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200101) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse200101) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse200101) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse200101) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse200101) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTasklists

`func (o *InlineResponse200101) GetTasklists() map[string]interface{}`

GetTasklists returns the Tasklists field if non-nil, zero value otherwise.

### GetTasklistsOk

`func (o *InlineResponse200101) GetTasklistsOk() (*map[string]interface{}, bool)`

GetTasklistsOk returns a tuple with the Tasklists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasklists

`func (o *InlineResponse200101) SetTasklists(v map[string]interface{})`

SetTasklists sets Tasklists field to given value.

### HasTasklists

`func (o *InlineResponse200101) HasTasklists() bool`

HasTasklists returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


