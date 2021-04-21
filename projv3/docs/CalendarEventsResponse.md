# CalendarEventsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CalendarEvents** | Pointer to [**[]ViewCalendarEvent**](ViewCalendarEvent.md) |  | [optional] 
**Included** | Pointer to [**CalendarEventsResponseIncluded**](CalendarEventsResponseIncluded.md) |  | [optional] 
**Meta** | Pointer to [**ViewMeta**](ViewMeta.md) |  | [optional] 

## Methods

### NewCalendarEventsResponse

`func NewCalendarEventsResponse() *CalendarEventsResponse`

NewCalendarEventsResponse instantiates a new CalendarEventsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalendarEventsResponseWithDefaults

`func NewCalendarEventsResponseWithDefaults() *CalendarEventsResponse`

NewCalendarEventsResponseWithDefaults instantiates a new CalendarEventsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalendarEvents

`func (o *CalendarEventsResponse) GetCalendarEvents() []ViewCalendarEvent`

GetCalendarEvents returns the CalendarEvents field if non-nil, zero value otherwise.

### GetCalendarEventsOk

`func (o *CalendarEventsResponse) GetCalendarEventsOk() (*[]ViewCalendarEvent, bool)`

GetCalendarEventsOk returns a tuple with the CalendarEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarEvents

`func (o *CalendarEventsResponse) SetCalendarEvents(v []ViewCalendarEvent)`

SetCalendarEvents sets CalendarEvents field to given value.

### HasCalendarEvents

`func (o *CalendarEventsResponse) HasCalendarEvents() bool`

HasCalendarEvents returns a boolean if a field has been set.

### GetIncluded

`func (o *CalendarEventsResponse) GetIncluded() CalendarEventsResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *CalendarEventsResponse) GetIncludedOk() (*CalendarEventsResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *CalendarEventsResponse) SetIncluded(v CalendarEventsResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *CalendarEventsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *CalendarEventsResponse) GetMeta() ViewMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CalendarEventsResponse) GetMetaOk() (*ViewMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CalendarEventsResponse) SetMeta(v ViewMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CalendarEventsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


