# TimerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timer** | Pointer to [**TimerTimer**](TimerTimer.md) |  | [optional] 

## Methods

### NewTimerRequest

`func NewTimerRequest() *TimerRequest`

NewTimerRequest instantiates a new TimerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimerRequestWithDefaults

`func NewTimerRequestWithDefaults() *TimerRequest`

NewTimerRequestWithDefaults instantiates a new TimerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimer

`func (o *TimerRequest) GetTimer() TimerTimer`

GetTimer returns the Timer field if non-nil, zero value otherwise.

### GetTimerOk

`func (o *TimerRequest) GetTimerOk() (*TimerTimer, bool)`

GetTimerOk returns a tuple with the Timer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimer

`func (o *TimerRequest) SetTimer(v TimerTimer)`

SetTimer sets Timer field to given value.

### HasTimer

`func (o *TimerRequest) HasTimer() bool`

HasTimer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


