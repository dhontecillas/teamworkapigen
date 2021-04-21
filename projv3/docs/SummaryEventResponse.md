# SummaryEventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Today** | Pointer to **int32** |  | [optional] 
**Upcoming** | Pointer to **int32** |  | [optional] 

## Methods

### NewSummaryEventResponse

`func NewSummaryEventResponse() *SummaryEventResponse`

NewSummaryEventResponse instantiates a new SummaryEventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryEventResponseWithDefaults

`func NewSummaryEventResponseWithDefaults() *SummaryEventResponse`

NewSummaryEventResponseWithDefaults instantiates a new SummaryEventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToday

`func (o *SummaryEventResponse) GetToday() int32`

GetToday returns the Today field if non-nil, zero value otherwise.

### GetTodayOk

`func (o *SummaryEventResponse) GetTodayOk() (*int32, bool)`

GetTodayOk returns a tuple with the Today field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToday

`func (o *SummaryEventResponse) SetToday(v int32)`

SetToday sets Today field to given value.

### HasToday

`func (o *SummaryEventResponse) HasToday() bool`

HasToday returns a boolean if a field has been set.

### GetUpcoming

`func (o *SummaryEventResponse) GetUpcoming() int32`

GetUpcoming returns the Upcoming field if non-nil, zero value otherwise.

### GetUpcomingOk

`func (o *SummaryEventResponse) GetUpcomingOk() (*int32, bool)`

GetUpcomingOk returns a tuple with the Upcoming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpcoming

`func (o *SummaryEventResponse) SetUpcoming(v int32)`

SetUpcoming sets Upcoming field to given value.

### HasUpcoming

`func (o *SummaryEventResponse) HasUpcoming() bool`

HasUpcoming returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


