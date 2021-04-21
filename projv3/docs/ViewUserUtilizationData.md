# ViewUserUtilizationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocatedMinutes** | Pointer to **int32** |  | [optional] 
**AvailableMinutes** | Pointer to **int32** |  | [optional] 
**BillableMinutes** | Pointer to **int32** |  | [optional] 
**EndDate** | Pointer to **map[string]interface{}** | Date represents a Unified API Spec date format. | [optional] 
**EstimatedMinutes** | Pointer to **int32** |  | [optional] 
**LoggedMinutes** | Pointer to **int32** |  | [optional] 
**StartDate** | Pointer to **map[string]interface{}** | Date represents a Unified API Spec date format. | [optional] 
**UnavailableMinutes** | Pointer to **int32** |  | [optional] 
**UnbillableMinutes** | Pointer to **int32** |  | [optional] 

## Methods

### NewViewUserUtilizationData

`func NewViewUserUtilizationData() *ViewUserUtilizationData`

NewViewUserUtilizationData instantiates a new ViewUserUtilizationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewUserUtilizationDataWithDefaults

`func NewViewUserUtilizationDataWithDefaults() *ViewUserUtilizationData`

NewViewUserUtilizationDataWithDefaults instantiates a new ViewUserUtilizationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocatedMinutes

`func (o *ViewUserUtilizationData) GetAllocatedMinutes() int32`

GetAllocatedMinutes returns the AllocatedMinutes field if non-nil, zero value otherwise.

### GetAllocatedMinutesOk

`func (o *ViewUserUtilizationData) GetAllocatedMinutesOk() (*int32, bool)`

GetAllocatedMinutesOk returns a tuple with the AllocatedMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedMinutes

`func (o *ViewUserUtilizationData) SetAllocatedMinutes(v int32)`

SetAllocatedMinutes sets AllocatedMinutes field to given value.

### HasAllocatedMinutes

`func (o *ViewUserUtilizationData) HasAllocatedMinutes() bool`

HasAllocatedMinutes returns a boolean if a field has been set.

### GetAvailableMinutes

`func (o *ViewUserUtilizationData) GetAvailableMinutes() int32`

GetAvailableMinutes returns the AvailableMinutes field if non-nil, zero value otherwise.

### GetAvailableMinutesOk

`func (o *ViewUserUtilizationData) GetAvailableMinutesOk() (*int32, bool)`

GetAvailableMinutesOk returns a tuple with the AvailableMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMinutes

`func (o *ViewUserUtilizationData) SetAvailableMinutes(v int32)`

SetAvailableMinutes sets AvailableMinutes field to given value.

### HasAvailableMinutes

`func (o *ViewUserUtilizationData) HasAvailableMinutes() bool`

HasAvailableMinutes returns a boolean if a field has been set.

### GetBillableMinutes

`func (o *ViewUserUtilizationData) GetBillableMinutes() int32`

GetBillableMinutes returns the BillableMinutes field if non-nil, zero value otherwise.

### GetBillableMinutesOk

`func (o *ViewUserUtilizationData) GetBillableMinutesOk() (*int32, bool)`

GetBillableMinutesOk returns a tuple with the BillableMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableMinutes

`func (o *ViewUserUtilizationData) SetBillableMinutes(v int32)`

SetBillableMinutes sets BillableMinutes field to given value.

### HasBillableMinutes

`func (o *ViewUserUtilizationData) HasBillableMinutes() bool`

HasBillableMinutes returns a boolean if a field has been set.

### GetEndDate

`func (o *ViewUserUtilizationData) GetEndDate() map[string]interface{}`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ViewUserUtilizationData) GetEndDateOk() (*map[string]interface{}, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ViewUserUtilizationData) SetEndDate(v map[string]interface{})`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ViewUserUtilizationData) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetEstimatedMinutes

`func (o *ViewUserUtilizationData) GetEstimatedMinutes() int32`

GetEstimatedMinutes returns the EstimatedMinutes field if non-nil, zero value otherwise.

### GetEstimatedMinutesOk

`func (o *ViewUserUtilizationData) GetEstimatedMinutesOk() (*int32, bool)`

GetEstimatedMinutesOk returns a tuple with the EstimatedMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedMinutes

`func (o *ViewUserUtilizationData) SetEstimatedMinutes(v int32)`

SetEstimatedMinutes sets EstimatedMinutes field to given value.

### HasEstimatedMinutes

`func (o *ViewUserUtilizationData) HasEstimatedMinutes() bool`

HasEstimatedMinutes returns a boolean if a field has been set.

### GetLoggedMinutes

`func (o *ViewUserUtilizationData) GetLoggedMinutes() int32`

GetLoggedMinutes returns the LoggedMinutes field if non-nil, zero value otherwise.

### GetLoggedMinutesOk

`func (o *ViewUserUtilizationData) GetLoggedMinutesOk() (*int32, bool)`

GetLoggedMinutesOk returns a tuple with the LoggedMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggedMinutes

`func (o *ViewUserUtilizationData) SetLoggedMinutes(v int32)`

SetLoggedMinutes sets LoggedMinutes field to given value.

### HasLoggedMinutes

`func (o *ViewUserUtilizationData) HasLoggedMinutes() bool`

HasLoggedMinutes returns a boolean if a field has been set.

### GetStartDate

`func (o *ViewUserUtilizationData) GetStartDate() map[string]interface{}`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ViewUserUtilizationData) GetStartDateOk() (*map[string]interface{}, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ViewUserUtilizationData) SetStartDate(v map[string]interface{})`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ViewUserUtilizationData) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetUnavailableMinutes

`func (o *ViewUserUtilizationData) GetUnavailableMinutes() int32`

GetUnavailableMinutes returns the UnavailableMinutes field if non-nil, zero value otherwise.

### GetUnavailableMinutesOk

`func (o *ViewUserUtilizationData) GetUnavailableMinutesOk() (*int32, bool)`

GetUnavailableMinutesOk returns a tuple with the UnavailableMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailableMinutes

`func (o *ViewUserUtilizationData) SetUnavailableMinutes(v int32)`

SetUnavailableMinutes sets UnavailableMinutes field to given value.

### HasUnavailableMinutes

`func (o *ViewUserUtilizationData) HasUnavailableMinutes() bool`

HasUnavailableMinutes returns a boolean if a field has been set.

### GetUnbillableMinutes

`func (o *ViewUserUtilizationData) GetUnbillableMinutes() int32`

GetUnbillableMinutes returns the UnbillableMinutes field if non-nil, zero value otherwise.

### GetUnbillableMinutesOk

`func (o *ViewUserUtilizationData) GetUnbillableMinutesOk() (*int32, bool)`

GetUnbillableMinutesOk returns a tuple with the UnbillableMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnbillableMinutes

`func (o *ViewUserUtilizationData) SetUnbillableMinutes(v int32)`

SetUnbillableMinutes sets UnbillableMinutes field to given value.

### HasUnbillableMinutes

`func (o *ViewUserUtilizationData) HasUnbillableMinutes() bool`

HasUnbillableMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


