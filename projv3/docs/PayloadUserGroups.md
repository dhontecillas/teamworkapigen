# PayloadUserGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyIds** | Pointer to [**PayloadNullableInt64Slice**](PayloadNullableInt64Slice.md) |  | [optional] 
**TeamIds** | Pointer to [**PayloadNullableInt64Slice**](PayloadNullableInt64Slice.md) |  | [optional] 
**UserIds** | Pointer to [**PayloadNullableInt64Slice**](PayloadNullableInt64Slice.md) |  | [optional] 

## Methods

### NewPayloadUserGroups

`func NewPayloadUserGroups() *PayloadUserGroups`

NewPayloadUserGroups instantiates a new PayloadUserGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayloadUserGroupsWithDefaults

`func NewPayloadUserGroupsWithDefaults() *PayloadUserGroups`

NewPayloadUserGroupsWithDefaults instantiates a new PayloadUserGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyIds

`func (o *PayloadUserGroups) GetCompanyIds() PayloadNullableInt64Slice`

GetCompanyIds returns the CompanyIds field if non-nil, zero value otherwise.

### GetCompanyIdsOk

`func (o *PayloadUserGroups) GetCompanyIdsOk() (*PayloadNullableInt64Slice, bool)`

GetCompanyIdsOk returns a tuple with the CompanyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIds

`func (o *PayloadUserGroups) SetCompanyIds(v PayloadNullableInt64Slice)`

SetCompanyIds sets CompanyIds field to given value.

### HasCompanyIds

`func (o *PayloadUserGroups) HasCompanyIds() bool`

HasCompanyIds returns a boolean if a field has been set.

### GetTeamIds

`func (o *PayloadUserGroups) GetTeamIds() PayloadNullableInt64Slice`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *PayloadUserGroups) GetTeamIdsOk() (*PayloadNullableInt64Slice, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *PayloadUserGroups) SetTeamIds(v PayloadNullableInt64Slice)`

SetTeamIds sets TeamIds field to given value.

### HasTeamIds

`func (o *PayloadUserGroups) HasTeamIds() bool`

HasTeamIds returns a boolean if a field has been set.

### GetUserIds

`func (o *PayloadUserGroups) GetUserIds() PayloadNullableInt64Slice`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *PayloadUserGroups) GetUserIdsOk() (*PayloadNullableInt64Slice, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *PayloadUserGroups) SetUserIds(v PayloadNullableInt64Slice)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *PayloadUserGroups) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


