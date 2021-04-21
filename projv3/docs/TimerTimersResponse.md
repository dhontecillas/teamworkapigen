# TimerTimersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Included** | Pointer to [**TimerResponseIncluded**](TimerResponseIncluded.md) |  | [optional] 
**Meta** | Pointer to [**ViewMeta**](ViewMeta.md) |  | [optional] 
**Timers** | Pointer to [**[]ViewTimer**](ViewTimer.md) |  | [optional] 

## Methods

### NewTimerTimersResponse

`func NewTimerTimersResponse() *TimerTimersResponse`

NewTimerTimersResponse instantiates a new TimerTimersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimerTimersResponseWithDefaults

`func NewTimerTimersResponseWithDefaults() *TimerTimersResponse`

NewTimerTimersResponseWithDefaults instantiates a new TimerTimersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncluded

`func (o *TimerTimersResponse) GetIncluded() TimerResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *TimerTimersResponse) GetIncludedOk() (*TimerResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *TimerTimersResponse) SetIncluded(v TimerResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *TimerTimersResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *TimerTimersResponse) GetMeta() ViewMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TimerTimersResponse) GetMetaOk() (*ViewMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TimerTimersResponse) SetMeta(v ViewMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TimerTimersResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetTimers

`func (o *TimerTimersResponse) GetTimers() []ViewTimer`

GetTimers returns the Timers field if non-nil, zero value otherwise.

### GetTimersOk

`func (o *TimerTimersResponse) GetTimersOk() (*[]ViewTimer, bool)`

GetTimersOk returns a tuple with the Timers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimers

`func (o *TimerTimersResponse) SetTimers(v []ViewTimer)`

SetTimers sets Timers field to given value.

### HasTimers

`func (o *TimerTimersResponse) HasTimers() bool`

HasTimers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


