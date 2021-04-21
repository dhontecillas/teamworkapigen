# ViewUsercost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | Pointer to **int32** |  | [optional] 
**User** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 

## Methods

### NewViewUsercost

`func NewViewUsercost() *ViewUsercost`

NewViewUsercost instantiates a new ViewUsercost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewUsercostWithDefaults

`func NewViewUsercostWithDefaults() *ViewUsercost`

NewViewUsercostWithDefaults instantiates a new ViewUsercost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *ViewUsercost) GetCost() int32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ViewUsercost) GetCostOk() (*int32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ViewUsercost) SetCost(v int32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ViewUsercost) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetUser

`func (o *ViewUsercost) GetUser() ViewRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ViewUsercost) GetUserOk() (*ViewRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ViewUsercost) SetUser(v ViewRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *ViewUsercost) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserId

`func (o *ViewUsercost) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ViewUsercost) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ViewUsercost) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ViewUsercost) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


