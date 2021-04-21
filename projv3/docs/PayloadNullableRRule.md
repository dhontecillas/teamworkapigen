# PayloadNullableRRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to **map[string]interface{}** | A Location maps time instants to the zone in use at that time. Typically, the Location represents the collection of time offsets in use in a geographical area. For many Locations the time offset varies depending on whether daylight savings time is in use at the time instant. | [optional] 
**Null** | Pointer to **bool** |  | [optional] 
**Set** | Pointer to **bool** |  | [optional] 

## Methods

### NewPayloadNullableRRule

`func NewPayloadNullableRRule() *PayloadNullableRRule`

NewPayloadNullableRRule instantiates a new PayloadNullableRRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayloadNullableRRuleWithDefaults

`func NewPayloadNullableRRuleWithDefaults() *PayloadNullableRRule`

NewPayloadNullableRRuleWithDefaults instantiates a new PayloadNullableRRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *PayloadNullableRRule) GetLocation() map[string]interface{}`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PayloadNullableRRule) GetLocationOk() (*map[string]interface{}, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PayloadNullableRRule) SetLocation(v map[string]interface{})`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PayloadNullableRRule) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetNull

`func (o *PayloadNullableRRule) GetNull() bool`

GetNull returns the Null field if non-nil, zero value otherwise.

### GetNullOk

`func (o *PayloadNullableRRule) GetNullOk() (*bool, bool)`

GetNullOk returns a tuple with the Null field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNull

`func (o *PayloadNullableRRule) SetNull(v bool)`

SetNull sets Null field to given value.

### HasNull

`func (o *PayloadNullableRRule) HasNull() bool`

HasNull returns a boolean if a field has been set.

### GetSet

`func (o *PayloadNullableRRule) GetSet() bool`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *PayloadNullableRRule) GetSetOk() (*bool, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *PayloadNullableRRule) SetSet(v bool)`

SetSet sets Set field to given value.

### HasSet

`func (o *PayloadNullableRRule) HasSet() bool`

HasSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


