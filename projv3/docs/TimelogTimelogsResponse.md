# TimelogTimelogsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Included** | Pointer to [**TimelogTimelogsResponseIncluded**](TimelogTimelogsResponseIncluded.md) |  | [optional] 
**Meta** | Pointer to [**ViewMeta**](ViewMeta.md) |  | [optional] 
**Timelogs** | Pointer to [**[]ViewTimelog**](ViewTimelog.md) |  | [optional] 

## Methods

### NewTimelogTimelogsResponse

`func NewTimelogTimelogsResponse() *TimelogTimelogsResponse`

NewTimelogTimelogsResponse instantiates a new TimelogTimelogsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelogTimelogsResponseWithDefaults

`func NewTimelogTimelogsResponseWithDefaults() *TimelogTimelogsResponse`

NewTimelogTimelogsResponseWithDefaults instantiates a new TimelogTimelogsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncluded

`func (o *TimelogTimelogsResponse) GetIncluded() TimelogTimelogsResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *TimelogTimelogsResponse) GetIncludedOk() (*TimelogTimelogsResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *TimelogTimelogsResponse) SetIncluded(v TimelogTimelogsResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *TimelogTimelogsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *TimelogTimelogsResponse) GetMeta() ViewMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TimelogTimelogsResponse) GetMetaOk() (*ViewMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TimelogTimelogsResponse) SetMeta(v ViewMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TimelogTimelogsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetTimelogs

`func (o *TimelogTimelogsResponse) GetTimelogs() []ViewTimelog`

GetTimelogs returns the Timelogs field if non-nil, zero value otherwise.

### GetTimelogsOk

`func (o *TimelogTimelogsResponse) GetTimelogsOk() (*[]ViewTimelog, bool)`

GetTimelogsOk returns a tuple with the Timelogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelogs

`func (o *TimelogTimelogsResponse) SetTimelogs(v []ViewTimelog)`

SetTimelogs sets Timelogs field to given value.

### HasTimelogs

`func (o *TimelogTimelogsResponse) HasTimelogs() bool`

HasTimelogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


