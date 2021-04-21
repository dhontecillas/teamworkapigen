# ViewProjectScheduleEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocatedMinutes** | Pointer to **int32** |  | [optional] 
**AllocationIds** | Pointer to **[]int32** |  | [optional] 
**Allocations** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**EventIds** | Pointer to **[]int32** |  | [optional] 
**Events** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**LoggedMinutes** | Pointer to **int32** |  | [optional] 
**MilestoneIds** | Pointer to **[]int32** |  | [optional] 
**Milestones** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**TimelogIds** | Pointer to **[]int32** |  | [optional] 
**Timelogs** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**UnavailableMinutes** | Pointer to **int32** |  | [optional] 

## Methods

### NewViewProjectScheduleEntry

`func NewViewProjectScheduleEntry() *ViewProjectScheduleEntry`

NewViewProjectScheduleEntry instantiates a new ViewProjectScheduleEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewProjectScheduleEntryWithDefaults

`func NewViewProjectScheduleEntryWithDefaults() *ViewProjectScheduleEntry`

NewViewProjectScheduleEntryWithDefaults instantiates a new ViewProjectScheduleEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocatedMinutes

`func (o *ViewProjectScheduleEntry) GetAllocatedMinutes() int32`

GetAllocatedMinutes returns the AllocatedMinutes field if non-nil, zero value otherwise.

### GetAllocatedMinutesOk

`func (o *ViewProjectScheduleEntry) GetAllocatedMinutesOk() (*int32, bool)`

GetAllocatedMinutesOk returns a tuple with the AllocatedMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedMinutes

`func (o *ViewProjectScheduleEntry) SetAllocatedMinutes(v int32)`

SetAllocatedMinutes sets AllocatedMinutes field to given value.

### HasAllocatedMinutes

`func (o *ViewProjectScheduleEntry) HasAllocatedMinutes() bool`

HasAllocatedMinutes returns a boolean if a field has been set.

### GetAllocationIds

`func (o *ViewProjectScheduleEntry) GetAllocationIds() []int32`

GetAllocationIds returns the AllocationIds field if non-nil, zero value otherwise.

### GetAllocationIdsOk

`func (o *ViewProjectScheduleEntry) GetAllocationIdsOk() (*[]int32, bool)`

GetAllocationIdsOk returns a tuple with the AllocationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationIds

`func (o *ViewProjectScheduleEntry) SetAllocationIds(v []int32)`

SetAllocationIds sets AllocationIds field to given value.

### HasAllocationIds

`func (o *ViewProjectScheduleEntry) HasAllocationIds() bool`

HasAllocationIds returns a boolean if a field has been set.

### GetAllocations

`func (o *ViewProjectScheduleEntry) GetAllocations() []ViewRelationship`

GetAllocations returns the Allocations field if non-nil, zero value otherwise.

### GetAllocationsOk

`func (o *ViewProjectScheduleEntry) GetAllocationsOk() (*[]ViewRelationship, bool)`

GetAllocationsOk returns a tuple with the Allocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocations

`func (o *ViewProjectScheduleEntry) SetAllocations(v []ViewRelationship)`

SetAllocations sets Allocations field to given value.

### HasAllocations

`func (o *ViewProjectScheduleEntry) HasAllocations() bool`

HasAllocations returns a boolean if a field has been set.

### GetEventIds

`func (o *ViewProjectScheduleEntry) GetEventIds() []int32`

GetEventIds returns the EventIds field if non-nil, zero value otherwise.

### GetEventIdsOk

`func (o *ViewProjectScheduleEntry) GetEventIdsOk() (*[]int32, bool)`

GetEventIdsOk returns a tuple with the EventIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIds

`func (o *ViewProjectScheduleEntry) SetEventIds(v []int32)`

SetEventIds sets EventIds field to given value.

### HasEventIds

`func (o *ViewProjectScheduleEntry) HasEventIds() bool`

HasEventIds returns a boolean if a field has been set.

### GetEvents

`func (o *ViewProjectScheduleEntry) GetEvents() []ViewRelationship`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ViewProjectScheduleEntry) GetEventsOk() (*[]ViewRelationship, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ViewProjectScheduleEntry) SetEvents(v []ViewRelationship)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ViewProjectScheduleEntry) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetLoggedMinutes

`func (o *ViewProjectScheduleEntry) GetLoggedMinutes() int32`

GetLoggedMinutes returns the LoggedMinutes field if non-nil, zero value otherwise.

### GetLoggedMinutesOk

`func (o *ViewProjectScheduleEntry) GetLoggedMinutesOk() (*int32, bool)`

GetLoggedMinutesOk returns a tuple with the LoggedMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggedMinutes

`func (o *ViewProjectScheduleEntry) SetLoggedMinutes(v int32)`

SetLoggedMinutes sets LoggedMinutes field to given value.

### HasLoggedMinutes

`func (o *ViewProjectScheduleEntry) HasLoggedMinutes() bool`

HasLoggedMinutes returns a boolean if a field has been set.

### GetMilestoneIds

`func (o *ViewProjectScheduleEntry) GetMilestoneIds() []int32`

GetMilestoneIds returns the MilestoneIds field if non-nil, zero value otherwise.

### GetMilestoneIdsOk

`func (o *ViewProjectScheduleEntry) GetMilestoneIdsOk() (*[]int32, bool)`

GetMilestoneIdsOk returns a tuple with the MilestoneIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestoneIds

`func (o *ViewProjectScheduleEntry) SetMilestoneIds(v []int32)`

SetMilestoneIds sets MilestoneIds field to given value.

### HasMilestoneIds

`func (o *ViewProjectScheduleEntry) HasMilestoneIds() bool`

HasMilestoneIds returns a boolean if a field has been set.

### GetMilestones

`func (o *ViewProjectScheduleEntry) GetMilestones() []ViewRelationship`

GetMilestones returns the Milestones field if non-nil, zero value otherwise.

### GetMilestonesOk

`func (o *ViewProjectScheduleEntry) GetMilestonesOk() (*[]ViewRelationship, bool)`

GetMilestonesOk returns a tuple with the Milestones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestones

`func (o *ViewProjectScheduleEntry) SetMilestones(v []ViewRelationship)`

SetMilestones sets Milestones field to given value.

### HasMilestones

`func (o *ViewProjectScheduleEntry) HasMilestones() bool`

HasMilestones returns a boolean if a field has been set.

### GetTimelogIds

`func (o *ViewProjectScheduleEntry) GetTimelogIds() []int32`

GetTimelogIds returns the TimelogIds field if non-nil, zero value otherwise.

### GetTimelogIdsOk

`func (o *ViewProjectScheduleEntry) GetTimelogIdsOk() (*[]int32, bool)`

GetTimelogIdsOk returns a tuple with the TimelogIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelogIds

`func (o *ViewProjectScheduleEntry) SetTimelogIds(v []int32)`

SetTimelogIds sets TimelogIds field to given value.

### HasTimelogIds

`func (o *ViewProjectScheduleEntry) HasTimelogIds() bool`

HasTimelogIds returns a boolean if a field has been set.

### GetTimelogs

`func (o *ViewProjectScheduleEntry) GetTimelogs() []ViewRelationship`

GetTimelogs returns the Timelogs field if non-nil, zero value otherwise.

### GetTimelogsOk

`func (o *ViewProjectScheduleEntry) GetTimelogsOk() (*[]ViewRelationship, bool)`

GetTimelogsOk returns a tuple with the Timelogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelogs

`func (o *ViewProjectScheduleEntry) SetTimelogs(v []ViewRelationship)`

SetTimelogs sets Timelogs field to given value.

### HasTimelogs

`func (o *ViewProjectScheduleEntry) HasTimelogs() bool`

HasTimelogs returns a boolean if a field has been set.

### GetUnavailableMinutes

`func (o *ViewProjectScheduleEntry) GetUnavailableMinutes() int32`

GetUnavailableMinutes returns the UnavailableMinutes field if non-nil, zero value otherwise.

### GetUnavailableMinutesOk

`func (o *ViewProjectScheduleEntry) GetUnavailableMinutesOk() (*int32, bool)`

GetUnavailableMinutesOk returns a tuple with the UnavailableMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailableMinutes

`func (o *ViewProjectScheduleEntry) SetUnavailableMinutes(v int32)`

SetUnavailableMinutes sets UnavailableMinutes field to given value.

### HasUnavailableMinutes

`func (o *ViewProjectScheduleEntry) HasUnavailableMinutes() bool`

HasUnavailableMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


