# SummaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | Pointer to [**SummaryColumnResponse**](SummaryColumnResponse.md) |  | [optional] 
**Events** | Pointer to [**SummaryEventResponse**](SummaryEventResponse.md) |  | [optional] 
**Health** | Pointer to [**SummaryHealthResponse**](SummaryHealthResponse.md) |  | [optional] 
**Milestones** | Pointer to [**SummaryMilestoneCountsResponse**](SummaryMilestoneCountsResponse.md) |  | [optional] 
**Risks** | Pointer to [**SummaryRiskResponse**](SummaryRiskResponse.md) |  | [optional] 
**Since** | Pointer to [**SummarySinceResponse**](SummarySinceResponse.md) |  | [optional] 
**Tasks** | Pointer to [**SummaryTaskResponse**](SummaryTaskResponse.md) |  | [optional] 
**Time** | Pointer to [**SummaryTimeCounterResponse**](SummaryTimeCounterResponse.md) |  | [optional] 
**Unread** | Pointer to [**SummaryUnreadResponse**](SummaryUnreadResponse.md) |  | [optional] 

## Methods

### NewSummaryResponse

`func NewSummaryResponse() *SummaryResponse`

NewSummaryResponse instantiates a new SummaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryResponseWithDefaults

`func NewSummaryResponseWithDefaults() *SummaryResponse`

NewSummaryResponseWithDefaults instantiates a new SummaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *SummaryResponse) GetColumns() SummaryColumnResponse`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *SummaryResponse) GetColumnsOk() (*SummaryColumnResponse, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *SummaryResponse) SetColumns(v SummaryColumnResponse)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *SummaryResponse) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetEvents

`func (o *SummaryResponse) GetEvents() SummaryEventResponse`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *SummaryResponse) GetEventsOk() (*SummaryEventResponse, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *SummaryResponse) SetEvents(v SummaryEventResponse)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *SummaryResponse) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetHealth

`func (o *SummaryResponse) GetHealth() SummaryHealthResponse`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *SummaryResponse) GetHealthOk() (*SummaryHealthResponse, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *SummaryResponse) SetHealth(v SummaryHealthResponse)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *SummaryResponse) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetMilestones

`func (o *SummaryResponse) GetMilestones() SummaryMilestoneCountsResponse`

GetMilestones returns the Milestones field if non-nil, zero value otherwise.

### GetMilestonesOk

`func (o *SummaryResponse) GetMilestonesOk() (*SummaryMilestoneCountsResponse, bool)`

GetMilestonesOk returns a tuple with the Milestones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestones

`func (o *SummaryResponse) SetMilestones(v SummaryMilestoneCountsResponse)`

SetMilestones sets Milestones field to given value.

### HasMilestones

`func (o *SummaryResponse) HasMilestones() bool`

HasMilestones returns a boolean if a field has been set.

### GetRisks

`func (o *SummaryResponse) GetRisks() SummaryRiskResponse`

GetRisks returns the Risks field if non-nil, zero value otherwise.

### GetRisksOk

`func (o *SummaryResponse) GetRisksOk() (*SummaryRiskResponse, bool)`

GetRisksOk returns a tuple with the Risks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisks

`func (o *SummaryResponse) SetRisks(v SummaryRiskResponse)`

SetRisks sets Risks field to given value.

### HasRisks

`func (o *SummaryResponse) HasRisks() bool`

HasRisks returns a boolean if a field has been set.

### GetSince

`func (o *SummaryResponse) GetSince() SummarySinceResponse`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *SummaryResponse) GetSinceOk() (*SummarySinceResponse, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *SummaryResponse) SetSince(v SummarySinceResponse)`

SetSince sets Since field to given value.

### HasSince

`func (o *SummaryResponse) HasSince() bool`

HasSince returns a boolean if a field has been set.

### GetTasks

`func (o *SummaryResponse) GetTasks() SummaryTaskResponse`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *SummaryResponse) GetTasksOk() (*SummaryTaskResponse, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *SummaryResponse) SetTasks(v SummaryTaskResponse)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *SummaryResponse) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetTime

`func (o *SummaryResponse) GetTime() SummaryTimeCounterResponse`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *SummaryResponse) GetTimeOk() (*SummaryTimeCounterResponse, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *SummaryResponse) SetTime(v SummaryTimeCounterResponse)`

SetTime sets Time field to given value.

### HasTime

`func (o *SummaryResponse) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetUnread

`func (o *SummaryResponse) GetUnread() SummaryUnreadResponse`

GetUnread returns the Unread field if non-nil, zero value otherwise.

### GetUnreadOk

`func (o *SummaryResponse) GetUnreadOk() (*SummaryUnreadResponse, bool)`

GetUnreadOk returns a tuple with the Unread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnread

`func (o *SummaryResponse) SetUnread(v SummaryUnreadResponse)`

SetUnread sets Unread field to given value.

### HasUnread

`func (o *SummaryResponse) HasUnread() bool`

HasUnread returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


