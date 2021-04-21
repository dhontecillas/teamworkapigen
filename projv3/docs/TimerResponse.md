# TimerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Included** | Pointer to [**TimerResponseIncluded**](TimerResponseIncluded.md) |  | [optional] 
**Timer** | Pointer to [**ViewTimer**](ViewTimer.md) |  | [optional] 

## Methods

### NewTimerResponse

`func NewTimerResponse() *TimerResponse`

NewTimerResponse instantiates a new TimerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimerResponseWithDefaults

`func NewTimerResponseWithDefaults() *TimerResponse`

NewTimerResponseWithDefaults instantiates a new TimerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncluded

`func (o *TimerResponse) GetIncluded() TimerResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *TimerResponse) GetIncludedOk() (*TimerResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *TimerResponse) SetIncluded(v TimerResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *TimerResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetTimer

`func (o *TimerResponse) GetTimer() ViewTimer`

GetTimer returns the Timer field if non-nil, zero value otherwise.

### GetTimerOk

`func (o *TimerResponse) GetTimerOk() (*ViewTimer, bool)`

GetTimerOk returns a tuple with the Timer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimer

`func (o *TimerResponse) SetTimer(v ViewTimer)`

SetTimer sets Timer field to given value.

### HasTimer

`func (o *TimerResponse) HasTimer() bool`

HasTimer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


