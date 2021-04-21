# ViewUserAvailability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dates** | Pointer to [**map[string]ViewUserAvailabilityDate**](ViewUserAvailabilityDate.md) |  | [optional] 
**User** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 

## Methods

### NewViewUserAvailability

`func NewViewUserAvailability() *ViewUserAvailability`

NewViewUserAvailability instantiates a new ViewUserAvailability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewUserAvailabilityWithDefaults

`func NewViewUserAvailabilityWithDefaults() *ViewUserAvailability`

NewViewUserAvailabilityWithDefaults instantiates a new ViewUserAvailability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDates

`func (o *ViewUserAvailability) GetDates() map[string]ViewUserAvailabilityDate`

GetDates returns the Dates field if non-nil, zero value otherwise.

### GetDatesOk

`func (o *ViewUserAvailability) GetDatesOk() (*map[string]ViewUserAvailabilityDate, bool)`

GetDatesOk returns a tuple with the Dates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDates

`func (o *ViewUserAvailability) SetDates(v map[string]ViewUserAvailabilityDate)`

SetDates sets Dates field to given value.

### HasDates

`func (o *ViewUserAvailability) HasDates() bool`

HasDates returns a boolean if a field has been set.

### GetUser

`func (o *ViewUserAvailability) GetUser() ViewRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ViewUserAvailability) GetUserOk() (*ViewRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ViewUserAvailability) SetUser(v ViewRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *ViewUserAvailability) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserId

`func (o *ViewUserAvailability) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ViewUserAvailability) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ViewUserAvailability) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ViewUserAvailability) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


