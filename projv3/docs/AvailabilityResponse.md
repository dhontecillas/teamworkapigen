# AvailabilityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Availability** | Pointer to [**ViewUserAvailability**](ViewUserAvailability.md) |  | [optional] 
**Included** | Pointer to [**AvailabilityResponseIncluded**](AvailabilityResponseIncluded.md) |  | [optional] 

## Methods

### NewAvailabilityResponse

`func NewAvailabilityResponse() *AvailabilityResponse`

NewAvailabilityResponse instantiates a new AvailabilityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailabilityResponseWithDefaults

`func NewAvailabilityResponseWithDefaults() *AvailabilityResponse`

NewAvailabilityResponseWithDefaults instantiates a new AvailabilityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailability

`func (o *AvailabilityResponse) GetAvailability() ViewUserAvailability`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *AvailabilityResponse) GetAvailabilityOk() (*ViewUserAvailability, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *AvailabilityResponse) SetAvailability(v ViewUserAvailability)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *AvailabilityResponse) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetIncluded

`func (o *AvailabilityResponse) GetIncluded() AvailabilityResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AvailabilityResponse) GetIncludedOk() (*AvailabilityResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AvailabilityResponse) SetIncluded(v AvailabilityResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AvailabilityResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


