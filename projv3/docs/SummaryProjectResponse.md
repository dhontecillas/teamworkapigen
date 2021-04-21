# SummaryProjectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | Pointer to [**SummaryColumnResponse**](SummaryColumnResponse.md) |  | [optional] 
**Events** | Pointer to [**SummaryEventResponse**](SummaryEventResponse.md) |  | [optional] 
**Health** | Pointer to [**SummaryHealthResponse**](SummaryHealthResponse.md) |  | [optional] 
**Milestones** | Pointer to [**SummaryMilestoneResponse**](SummaryMilestoneResponse.md) |  | [optional] 
**Risks** | Pointer to [**SummaryRiskResponse**](SummaryRiskResponse.md) |  | [optional] 
**Since** | Pointer to [**SummarySinceResponse**](SummarySinceResponse.md) |  | [optional] 
**Tasks** | Pointer to [**SummaryProjectTasksResponse**](SummaryProjectTasksResponse.md) |  | [optional] 
**Time** | Pointer to [**SummaryTimeResponse**](SummaryTimeResponse.md) |  | [optional] 
**Unread** | Pointer to [**SummaryUnreadResponse**](SummaryUnreadResponse.md) |  | [optional] 

## Methods

### NewSummaryProjectResponse

`func NewSummaryProjectResponse() *SummaryProjectResponse`

NewSummaryProjectResponse instantiates a new SummaryProjectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryProjectResponseWithDefaults

`func NewSummaryProjectResponseWithDefaults() *SummaryProjectResponse`

NewSummaryProjectResponseWithDefaults instantiates a new SummaryProjectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *SummaryProjectResponse) GetColumns() SummaryColumnResponse`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *SummaryProjectResponse) GetColumnsOk() (*SummaryColumnResponse, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *SummaryProjectResponse) SetColumns(v SummaryColumnResponse)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *SummaryProjectResponse) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetEvents

`func (o *SummaryProjectResponse) GetEvents() SummaryEventResponse`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *SummaryProjectResponse) GetEventsOk() (*SummaryEventResponse, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *SummaryProjectResponse) SetEvents(v SummaryEventResponse)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *SummaryProjectResponse) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetHealth

`func (o *SummaryProjectResponse) GetHealth() SummaryHealthResponse`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *SummaryProjectResponse) GetHealthOk() (*SummaryHealthResponse, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *SummaryProjectResponse) SetHealth(v SummaryHealthResponse)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *SummaryProjectResponse) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetMilestones

`func (o *SummaryProjectResponse) GetMilestones() SummaryMilestoneResponse`

GetMilestones returns the Milestones field if non-nil, zero value otherwise.

### GetMilestonesOk

`func (o *SummaryProjectResponse) GetMilestonesOk() (*SummaryMilestoneResponse, bool)`

GetMilestonesOk returns a tuple with the Milestones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestones

`func (o *SummaryProjectResponse) SetMilestones(v SummaryMilestoneResponse)`

SetMilestones sets Milestones field to given value.

### HasMilestones

`func (o *SummaryProjectResponse) HasMilestones() bool`

HasMilestones returns a boolean if a field has been set.

### GetRisks

`func (o *SummaryProjectResponse) GetRisks() SummaryRiskResponse`

GetRisks returns the Risks field if non-nil, zero value otherwise.

### GetRisksOk

`func (o *SummaryProjectResponse) GetRisksOk() (*SummaryRiskResponse, bool)`

GetRisksOk returns a tuple with the Risks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisks

`func (o *SummaryProjectResponse) SetRisks(v SummaryRiskResponse)`

SetRisks sets Risks field to given value.

### HasRisks

`func (o *SummaryProjectResponse) HasRisks() bool`

HasRisks returns a boolean if a field has been set.

### GetSince

`func (o *SummaryProjectResponse) GetSince() SummarySinceResponse`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *SummaryProjectResponse) GetSinceOk() (*SummarySinceResponse, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *SummaryProjectResponse) SetSince(v SummarySinceResponse)`

SetSince sets Since field to given value.

### HasSince

`func (o *SummaryProjectResponse) HasSince() bool`

HasSince returns a boolean if a field has been set.

### GetTasks

`func (o *SummaryProjectResponse) GetTasks() SummaryProjectTasksResponse`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *SummaryProjectResponse) GetTasksOk() (*SummaryProjectTasksResponse, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *SummaryProjectResponse) SetTasks(v SummaryProjectTasksResponse)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *SummaryProjectResponse) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetTime

`func (o *SummaryProjectResponse) GetTime() SummaryTimeResponse`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *SummaryProjectResponse) GetTimeOk() (*SummaryTimeResponse, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *SummaryProjectResponse) SetTime(v SummaryTimeResponse)`

SetTime sets Time field to given value.

### HasTime

`func (o *SummaryProjectResponse) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetUnread

`func (o *SummaryProjectResponse) GetUnread() SummaryUnreadResponse`

GetUnread returns the Unread field if non-nil, zero value otherwise.

### GetUnreadOk

`func (o *SummaryProjectResponse) GetUnreadOk() (*SummaryUnreadResponse, bool)`

GetUnreadOk returns a tuple with the Unread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnread

`func (o *SummaryProjectResponse) SetUnread(v SummaryUnreadResponse)`

SetUnread sets Unread field to given value.

### HasUnread

`func (o *SummaryProjectResponse) HasUnread() bool`

HasUnread returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


