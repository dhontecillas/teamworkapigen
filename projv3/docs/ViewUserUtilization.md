# ViewUserUtilization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ViewUserUtilizationData**](ViewUserUtilizationData.md) |  | [optional] 
**User** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 

## Methods

### NewViewUserUtilization

`func NewViewUserUtilization() *ViewUserUtilization`

NewViewUserUtilization instantiates a new ViewUserUtilization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewUserUtilizationWithDefaults

`func NewViewUserUtilizationWithDefaults() *ViewUserUtilization`

NewViewUserUtilizationWithDefaults instantiates a new ViewUserUtilization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ViewUserUtilization) GetData() []ViewUserUtilizationData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ViewUserUtilization) GetDataOk() (*[]ViewUserUtilizationData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ViewUserUtilization) SetData(v []ViewUserUtilizationData)`

SetData sets Data field to given value.

### HasData

`func (o *ViewUserUtilization) HasData() bool`

HasData returns a boolean if a field has been set.

### GetUser

`func (o *ViewUserUtilization) GetUser() ViewRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ViewUserUtilization) GetUserOk() (*ViewRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ViewUserUtilization) SetUser(v ViewRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *ViewUserUtilization) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserId

`func (o *ViewUserUtilization) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ViewUserUtilization) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ViewUserUtilization) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ViewUserUtilization) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


