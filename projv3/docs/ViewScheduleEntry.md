# ViewScheduleEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocatedMinutes** | Pointer to **int32** |  | [optional] 
**AllocationIds** | Pointer to **[]int32** |  | [optional] 
**Allocations** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**EventIds** | Pointer to **[]int32** |  | [optional] 
**Events** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**LoggedMinutes** | Pointer to **int32** |  | [optional] 
**TimelogIds** | Pointer to **[]int32** |  | [optional] 
**Timelogs** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**UnavailableMinutes** | Pointer to **int32** |  | [optional] 

## Methods

### NewViewScheduleEntry

`func NewViewScheduleEntry() *ViewScheduleEntry`

NewViewScheduleEntry instantiates a new ViewScheduleEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewScheduleEntryWithDefaults

`func NewViewScheduleEntryWithDefaults() *ViewScheduleEntry`

NewViewScheduleEntryWithDefaults instantiates a new ViewScheduleEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocatedMinutes

`func (o *ViewScheduleEntry) GetAllocatedMinutes() int32`

GetAllocatedMinutes returns the AllocatedMinutes field if non-nil, zero value otherwise.

### GetAllocatedMinutesOk

`func (o *ViewScheduleEntry) GetAllocatedMinutesOk() (*int32, bool)`

GetAllocatedMinutesOk returns a tuple with the AllocatedMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedMinutes

`func (o *ViewScheduleEntry) SetAllocatedMinutes(v int32)`

SetAllocatedMinutes sets AllocatedMinutes field to given value.

### HasAllocatedMinutes

`func (o *ViewScheduleEntry) HasAllocatedMinutes() bool`

HasAllocatedMinutes returns a boolean if a field has been set.

### GetAllocationIds

`func (o *ViewScheduleEntry) GetAllocationIds() []int32`

GetAllocationIds returns the AllocationIds field if non-nil, zero value otherwise.

### GetAllocationIdsOk

`func (o *ViewScheduleEntry) GetAllocationIdsOk() (*[]int32, bool)`

GetAllocationIdsOk returns a tuple with the AllocationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationIds

`func (o *ViewScheduleEntry) SetAllocationIds(v []int32)`

SetAllocationIds sets AllocationIds field to given value.

### HasAllocationIds

`func (o *ViewScheduleEntry) HasAllocationIds() bool`

HasAllocationIds returns a boolean if a field has been set.

### GetAllocations

`func (o *ViewScheduleEntry) GetAllocations() []ViewRelationship`

GetAllocations returns the Allocations field if non-nil, zero value otherwise.

### GetAllocationsOk

`func (o *ViewScheduleEntry) GetAllocationsOk() (*[]ViewRelationship, bool)`

GetAllocationsOk returns a tuple with the Allocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocations

`func (o *ViewScheduleEntry) SetAllocations(v []ViewRelationship)`

SetAllocations sets Allocations field to given value.

### HasAllocations

`func (o *ViewScheduleEntry) HasAllocations() bool`

HasAllocations returns a boolean if a field has been set.

### GetEventIds

`func (o *ViewScheduleEntry) GetEventIds() []int32`

GetEventIds returns the EventIds field if non-nil, zero value otherwise.

### GetEventIdsOk

`func (o *ViewScheduleEntry) GetEventIdsOk() (*[]int32, bool)`

GetEventIdsOk returns a tuple with the EventIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIds

`func (o *ViewScheduleEntry) SetEventIds(v []int32)`

SetEventIds sets EventIds field to given value.

### HasEventIds

`func (o *ViewScheduleEntry) HasEventIds() bool`

HasEventIds returns a boolean if a field has been set.

### GetEvents

`func (o *ViewScheduleEntry) GetEvents() []ViewRelationship`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ViewScheduleEntry) GetEventsOk() (*[]ViewRelationship, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ViewScheduleEntry) SetEvents(v []ViewRelationship)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ViewScheduleEntry) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetLoggedMinutes

`func (o *ViewScheduleEntry) GetLoggedMinutes() int32`

GetLoggedMinutes returns the LoggedMinutes field if non-nil, zero value otherwise.

### GetLoggedMinutesOk

`func (o *ViewScheduleEntry) GetLoggedMinutesOk() (*int32, bool)`

GetLoggedMinutesOk returns a tuple with the LoggedMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggedMinutes

`func (o *ViewScheduleEntry) SetLoggedMinutes(v int32)`

SetLoggedMinutes sets LoggedMinutes field to given value.

### HasLoggedMinutes

`func (o *ViewScheduleEntry) HasLoggedMinutes() bool`

HasLoggedMinutes returns a boolean if a field has been set.

### GetTimelogIds

`func (o *ViewScheduleEntry) GetTimelogIds() []int32`

GetTimelogIds returns the TimelogIds field if non-nil, zero value otherwise.

### GetTimelogIdsOk

`func (o *ViewScheduleEntry) GetTimelogIdsOk() (*[]int32, bool)`

GetTimelogIdsOk returns a tuple with the TimelogIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelogIds

`func (o *ViewScheduleEntry) SetTimelogIds(v []int32)`

SetTimelogIds sets TimelogIds field to given value.

### HasTimelogIds

`func (o *ViewScheduleEntry) HasTimelogIds() bool`

HasTimelogIds returns a boolean if a field has been set.

### GetTimelogs

`func (o *ViewScheduleEntry) GetTimelogs() []ViewRelationship`

GetTimelogs returns the Timelogs field if non-nil, zero value otherwise.

### GetTimelogsOk

`func (o *ViewScheduleEntry) GetTimelogsOk() (*[]ViewRelationship, bool)`

GetTimelogsOk returns a tuple with the Timelogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelogs

`func (o *ViewScheduleEntry) SetTimelogs(v []ViewRelationship)`

SetTimelogs sets Timelogs field to given value.

### HasTimelogs

`func (o *ViewScheduleEntry) HasTimelogs() bool`

HasTimelogs returns a boolean if a field has been set.

### GetUnavailableMinutes

`func (o *ViewScheduleEntry) GetUnavailableMinutes() int32`

GetUnavailableMinutes returns the UnavailableMinutes field if non-nil, zero value otherwise.

### GetUnavailableMinutesOk

`func (o *ViewScheduleEntry) GetUnavailableMinutesOk() (*int32, bool)`

GetUnavailableMinutesOk returns a tuple with the UnavailableMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailableMinutes

`func (o *ViewScheduleEntry) SetUnavailableMinutes(v int32)`

SetUnavailableMinutes sets UnavailableMinutes field to given value.

### HasUnavailableMinutes

`func (o *ViewScheduleEntry) HasUnavailableMinutes() bool`

HasUnavailableMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


