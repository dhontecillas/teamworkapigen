# StatusTimelineResponseIncluded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserEvents** | Pointer to [**map[string]ViewUserEvents**](ViewUserEvents.md) |  | [optional] 
**Users** | Pointer to [**map[string]ViewUser**](ViewUser.md) |  | [optional] 

## Methods

### NewStatusTimelineResponseIncluded

`func NewStatusTimelineResponseIncluded() *StatusTimelineResponseIncluded`

NewStatusTimelineResponseIncluded instantiates a new StatusTimelineResponseIncluded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusTimelineResponseIncludedWithDefaults

`func NewStatusTimelineResponseIncludedWithDefaults() *StatusTimelineResponseIncluded`

NewStatusTimelineResponseIncludedWithDefaults instantiates a new StatusTimelineResponseIncluded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserEvents

`func (o *StatusTimelineResponseIncluded) GetUserEvents() map[string]ViewUserEvents`

GetUserEvents returns the UserEvents field if non-nil, zero value otherwise.

### GetUserEventsOk

`func (o *StatusTimelineResponseIncluded) GetUserEventsOk() (*map[string]ViewUserEvents, bool)`

GetUserEventsOk returns a tuple with the UserEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEvents

`func (o *StatusTimelineResponseIncluded) SetUserEvents(v map[string]ViewUserEvents)`

SetUserEvents sets UserEvents field to given value.

### HasUserEvents

`func (o *StatusTimelineResponseIncluded) HasUserEvents() bool`

HasUserEvents returns a boolean if a field has been set.

### GetUsers

`func (o *StatusTimelineResponseIncluded) GetUsers() map[string]ViewUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *StatusTimelineResponseIncluded) GetUsersOk() (*map[string]ViewUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *StatusTimelineResponseIncluded) SetUsers(v map[string]ViewUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *StatusTimelineResponseIncluded) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


