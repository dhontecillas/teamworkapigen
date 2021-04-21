# PeopleNotifications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DailyReportDaysFilter** | Pointer to **int32** |  | [optional] 
**DailyReportEventsType** | Pointer to **string** |  | [optional] 
**DailyReportIncludeStartDate** | Pointer to **bool** |  | [optional] 
**DailyReportSortBy** | Pointer to **string** |  | [optional] 
**NotifyOnAddedAsFollower** | Pointer to **bool** |  | [optional] 
**NotifyOnStatusUpdate** | Pointer to **bool** |  | [optional] 
**NotifyOnTaskComplete** | Pointer to **bool** |  | [optional] 
**ReceiveDailyReports** | Pointer to **bool** |  | [optional] 
**ReceiveDailyReportsAtTime** | Pointer to **int32** |  | [optional] 
**ReceiveDailyReportsAtWeekend** | Pointer to **bool** |  | [optional] 
**ReceiveDailyReportsIfEmpty** | Pointer to **bool** |  | [optional] 
**ReceiveMyNotificationsOnly** | Pointer to **bool** |  | [optional] 
**ReceiveNotifyWarnings** | Pointer to **bool** |  | [optional] 
**SoundAlertsEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewPeopleNotifications

`func NewPeopleNotifications() *PeopleNotifications`

NewPeopleNotifications instantiates a new PeopleNotifications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeopleNotificationsWithDefaults

`func NewPeopleNotificationsWithDefaults() *PeopleNotifications`

NewPeopleNotificationsWithDefaults instantiates a new PeopleNotifications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDailyReportDaysFilter

`func (o *PeopleNotifications) GetDailyReportDaysFilter() int32`

GetDailyReportDaysFilter returns the DailyReportDaysFilter field if non-nil, zero value otherwise.

### GetDailyReportDaysFilterOk

`func (o *PeopleNotifications) GetDailyReportDaysFilterOk() (*int32, bool)`

GetDailyReportDaysFilterOk returns a tuple with the DailyReportDaysFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyReportDaysFilter

`func (o *PeopleNotifications) SetDailyReportDaysFilter(v int32)`

SetDailyReportDaysFilter sets DailyReportDaysFilter field to given value.

### HasDailyReportDaysFilter

`func (o *PeopleNotifications) HasDailyReportDaysFilter() bool`

HasDailyReportDaysFilter returns a boolean if a field has been set.

### GetDailyReportEventsType

`func (o *PeopleNotifications) GetDailyReportEventsType() string`

GetDailyReportEventsType returns the DailyReportEventsType field if non-nil, zero value otherwise.

### GetDailyReportEventsTypeOk

`func (o *PeopleNotifications) GetDailyReportEventsTypeOk() (*string, bool)`

GetDailyReportEventsTypeOk returns a tuple with the DailyReportEventsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyReportEventsType

`func (o *PeopleNotifications) SetDailyReportEventsType(v string)`

SetDailyReportEventsType sets DailyReportEventsType field to given value.

### HasDailyReportEventsType

`func (o *PeopleNotifications) HasDailyReportEventsType() bool`

HasDailyReportEventsType returns a boolean if a field has been set.

### GetDailyReportIncludeStartDate

`func (o *PeopleNotifications) GetDailyReportIncludeStartDate() bool`

GetDailyReportIncludeStartDate returns the DailyReportIncludeStartDate field if non-nil, zero value otherwise.

### GetDailyReportIncludeStartDateOk

`func (o *PeopleNotifications) GetDailyReportIncludeStartDateOk() (*bool, bool)`

GetDailyReportIncludeStartDateOk returns a tuple with the DailyReportIncludeStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyReportIncludeStartDate

`func (o *PeopleNotifications) SetDailyReportIncludeStartDate(v bool)`

SetDailyReportIncludeStartDate sets DailyReportIncludeStartDate field to given value.

### HasDailyReportIncludeStartDate

`func (o *PeopleNotifications) HasDailyReportIncludeStartDate() bool`

HasDailyReportIncludeStartDate returns a boolean if a field has been set.

### GetDailyReportSortBy

`func (o *PeopleNotifications) GetDailyReportSortBy() string`

GetDailyReportSortBy returns the DailyReportSortBy field if non-nil, zero value otherwise.

### GetDailyReportSortByOk

`func (o *PeopleNotifications) GetDailyReportSortByOk() (*string, bool)`

GetDailyReportSortByOk returns a tuple with the DailyReportSortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyReportSortBy

`func (o *PeopleNotifications) SetDailyReportSortBy(v string)`

SetDailyReportSortBy sets DailyReportSortBy field to given value.

### HasDailyReportSortBy

`func (o *PeopleNotifications) HasDailyReportSortBy() bool`

HasDailyReportSortBy returns a boolean if a field has been set.

### GetNotifyOnAddedAsFollower

`func (o *PeopleNotifications) GetNotifyOnAddedAsFollower() bool`

GetNotifyOnAddedAsFollower returns the NotifyOnAddedAsFollower field if non-nil, zero value otherwise.

### GetNotifyOnAddedAsFollowerOk

`func (o *PeopleNotifications) GetNotifyOnAddedAsFollowerOk() (*bool, bool)`

GetNotifyOnAddedAsFollowerOk returns a tuple with the NotifyOnAddedAsFollower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyOnAddedAsFollower

`func (o *PeopleNotifications) SetNotifyOnAddedAsFollower(v bool)`

SetNotifyOnAddedAsFollower sets NotifyOnAddedAsFollower field to given value.

### HasNotifyOnAddedAsFollower

`func (o *PeopleNotifications) HasNotifyOnAddedAsFollower() bool`

HasNotifyOnAddedAsFollower returns a boolean if a field has been set.

### GetNotifyOnStatusUpdate

`func (o *PeopleNotifications) GetNotifyOnStatusUpdate() bool`

GetNotifyOnStatusUpdate returns the NotifyOnStatusUpdate field if non-nil, zero value otherwise.

### GetNotifyOnStatusUpdateOk

`func (o *PeopleNotifications) GetNotifyOnStatusUpdateOk() (*bool, bool)`

GetNotifyOnStatusUpdateOk returns a tuple with the NotifyOnStatusUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyOnStatusUpdate

`func (o *PeopleNotifications) SetNotifyOnStatusUpdate(v bool)`

SetNotifyOnStatusUpdate sets NotifyOnStatusUpdate field to given value.

### HasNotifyOnStatusUpdate

`func (o *PeopleNotifications) HasNotifyOnStatusUpdate() bool`

HasNotifyOnStatusUpdate returns a boolean if a field has been set.

### GetNotifyOnTaskComplete

`func (o *PeopleNotifications) GetNotifyOnTaskComplete() bool`

GetNotifyOnTaskComplete returns the NotifyOnTaskComplete field if non-nil, zero value otherwise.

### GetNotifyOnTaskCompleteOk

`func (o *PeopleNotifications) GetNotifyOnTaskCompleteOk() (*bool, bool)`

GetNotifyOnTaskCompleteOk returns a tuple with the NotifyOnTaskComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyOnTaskComplete

`func (o *PeopleNotifications) SetNotifyOnTaskComplete(v bool)`

SetNotifyOnTaskComplete sets NotifyOnTaskComplete field to given value.

### HasNotifyOnTaskComplete

`func (o *PeopleNotifications) HasNotifyOnTaskComplete() bool`

HasNotifyOnTaskComplete returns a boolean if a field has been set.

### GetReceiveDailyReports

`func (o *PeopleNotifications) GetReceiveDailyReports() bool`

GetReceiveDailyReports returns the ReceiveDailyReports field if non-nil, zero value otherwise.

### GetReceiveDailyReportsOk

`func (o *PeopleNotifications) GetReceiveDailyReportsOk() (*bool, bool)`

GetReceiveDailyReportsOk returns a tuple with the ReceiveDailyReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveDailyReports

`func (o *PeopleNotifications) SetReceiveDailyReports(v bool)`

SetReceiveDailyReports sets ReceiveDailyReports field to given value.

### HasReceiveDailyReports

`func (o *PeopleNotifications) HasReceiveDailyReports() bool`

HasReceiveDailyReports returns a boolean if a field has been set.

### GetReceiveDailyReportsAtTime

`func (o *PeopleNotifications) GetReceiveDailyReportsAtTime() int32`

GetReceiveDailyReportsAtTime returns the ReceiveDailyReportsAtTime field if non-nil, zero value otherwise.

### GetReceiveDailyReportsAtTimeOk

`func (o *PeopleNotifications) GetReceiveDailyReportsAtTimeOk() (*int32, bool)`

GetReceiveDailyReportsAtTimeOk returns a tuple with the ReceiveDailyReportsAtTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveDailyReportsAtTime

`func (o *PeopleNotifications) SetReceiveDailyReportsAtTime(v int32)`

SetReceiveDailyReportsAtTime sets ReceiveDailyReportsAtTime field to given value.

### HasReceiveDailyReportsAtTime

`func (o *PeopleNotifications) HasReceiveDailyReportsAtTime() bool`

HasReceiveDailyReportsAtTime returns a boolean if a field has been set.

### GetReceiveDailyReportsAtWeekend

`func (o *PeopleNotifications) GetReceiveDailyReportsAtWeekend() bool`

GetReceiveDailyReportsAtWeekend returns the ReceiveDailyReportsAtWeekend field if non-nil, zero value otherwise.

### GetReceiveDailyReportsAtWeekendOk

`func (o *PeopleNotifications) GetReceiveDailyReportsAtWeekendOk() (*bool, bool)`

GetReceiveDailyReportsAtWeekendOk returns a tuple with the ReceiveDailyReportsAtWeekend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveDailyReportsAtWeekend

`func (o *PeopleNotifications) SetReceiveDailyReportsAtWeekend(v bool)`

SetReceiveDailyReportsAtWeekend sets ReceiveDailyReportsAtWeekend field to given value.

### HasReceiveDailyReportsAtWeekend

`func (o *PeopleNotifications) HasReceiveDailyReportsAtWeekend() bool`

HasReceiveDailyReportsAtWeekend returns a boolean if a field has been set.

### GetReceiveDailyReportsIfEmpty

`func (o *PeopleNotifications) GetReceiveDailyReportsIfEmpty() bool`

GetReceiveDailyReportsIfEmpty returns the ReceiveDailyReportsIfEmpty field if non-nil, zero value otherwise.

### GetReceiveDailyReportsIfEmptyOk

`func (o *PeopleNotifications) GetReceiveDailyReportsIfEmptyOk() (*bool, bool)`

GetReceiveDailyReportsIfEmptyOk returns a tuple with the ReceiveDailyReportsIfEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveDailyReportsIfEmpty

`func (o *PeopleNotifications) SetReceiveDailyReportsIfEmpty(v bool)`

SetReceiveDailyReportsIfEmpty sets ReceiveDailyReportsIfEmpty field to given value.

### HasReceiveDailyReportsIfEmpty

`func (o *PeopleNotifications) HasReceiveDailyReportsIfEmpty() bool`

HasReceiveDailyReportsIfEmpty returns a boolean if a field has been set.

### GetReceiveMyNotificationsOnly

`func (o *PeopleNotifications) GetReceiveMyNotificationsOnly() bool`

GetReceiveMyNotificationsOnly returns the ReceiveMyNotificationsOnly field if non-nil, zero value otherwise.

### GetReceiveMyNotificationsOnlyOk

`func (o *PeopleNotifications) GetReceiveMyNotificationsOnlyOk() (*bool, bool)`

GetReceiveMyNotificationsOnlyOk returns a tuple with the ReceiveMyNotificationsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveMyNotificationsOnly

`func (o *PeopleNotifications) SetReceiveMyNotificationsOnly(v bool)`

SetReceiveMyNotificationsOnly sets ReceiveMyNotificationsOnly field to given value.

### HasReceiveMyNotificationsOnly

`func (o *PeopleNotifications) HasReceiveMyNotificationsOnly() bool`

HasReceiveMyNotificationsOnly returns a boolean if a field has been set.

### GetReceiveNotifyWarnings

`func (o *PeopleNotifications) GetReceiveNotifyWarnings() bool`

GetReceiveNotifyWarnings returns the ReceiveNotifyWarnings field if non-nil, zero value otherwise.

### GetReceiveNotifyWarningsOk

`func (o *PeopleNotifications) GetReceiveNotifyWarningsOk() (*bool, bool)`

GetReceiveNotifyWarningsOk returns a tuple with the ReceiveNotifyWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveNotifyWarnings

`func (o *PeopleNotifications) SetReceiveNotifyWarnings(v bool)`

SetReceiveNotifyWarnings sets ReceiveNotifyWarnings field to given value.

### HasReceiveNotifyWarnings

`func (o *PeopleNotifications) HasReceiveNotifyWarnings() bool`

HasReceiveNotifyWarnings returns a boolean if a field has been set.

### GetSoundAlertsEnabled

`func (o *PeopleNotifications) GetSoundAlertsEnabled() bool`

GetSoundAlertsEnabled returns the SoundAlertsEnabled field if non-nil, zero value otherwise.

### GetSoundAlertsEnabledOk

`func (o *PeopleNotifications) GetSoundAlertsEnabledOk() (*bool, bool)`

GetSoundAlertsEnabledOk returns a tuple with the SoundAlertsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoundAlertsEnabled

`func (o *PeopleNotifications) SetSoundAlertsEnabled(v bool)`

SetSoundAlertsEnabled sets SoundAlertsEnabled field to given value.

### HasSoundAlertsEnabled

`func (o *PeopleNotifications) HasSoundAlertsEnabled() bool`

HasSoundAlertsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


