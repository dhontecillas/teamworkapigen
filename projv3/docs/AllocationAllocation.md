# AllocationAllocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedUserID** | Pointer to **int32** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DistributeType** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **int32** | in minutes | [optional] 
**EndedAt** | Pointer to **map[string]interface{}** | Date unmarshals represents a Unified API Spec date format. | [optional] 
**HoursPerDay** | Pointer to **float32** |  | [optional] 
**IgnoreCollisions** | Pointer to **bool** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**RecurringRule** | Pointer to [**PayloadNullableRRule**](PayloadNullableRRule.md) |  | [optional] 
**StartedAt** | Pointer to **map[string]interface{}** | Date unmarshals represents a Unified API Spec date format. | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewAllocationAllocation

`func NewAllocationAllocation() *AllocationAllocation`

NewAllocationAllocation instantiates a new AllocationAllocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllocationAllocationWithDefaults

`func NewAllocationAllocationWithDefaults() *AllocationAllocation`

NewAllocationAllocationWithDefaults instantiates a new AllocationAllocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedUserID

`func (o *AllocationAllocation) GetAssignedUserID() int32`

GetAssignedUserID returns the AssignedUserID field if non-nil, zero value otherwise.

### GetAssignedUserIDOk

`func (o *AllocationAllocation) GetAssignedUserIDOk() (*int32, bool)`

GetAssignedUserIDOk returns a tuple with the AssignedUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUserID

`func (o *AllocationAllocation) SetAssignedUserID(v int32)`

SetAssignedUserID sets AssignedUserID field to given value.

### HasAssignedUserID

`func (o *AllocationAllocation) HasAssignedUserID() bool`

HasAssignedUserID returns a boolean if a field has been set.

### GetColor

`func (o *AllocationAllocation) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *AllocationAllocation) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *AllocationAllocation) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *AllocationAllocation) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetDescription

`func (o *AllocationAllocation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AllocationAllocation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AllocationAllocation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AllocationAllocation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDistributeType

`func (o *AllocationAllocation) GetDistributeType() string`

GetDistributeType returns the DistributeType field if non-nil, zero value otherwise.

### GetDistributeTypeOk

`func (o *AllocationAllocation) GetDistributeTypeOk() (*string, bool)`

GetDistributeTypeOk returns a tuple with the DistributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributeType

`func (o *AllocationAllocation) SetDistributeType(v string)`

SetDistributeType sets DistributeType field to given value.

### HasDistributeType

`func (o *AllocationAllocation) HasDistributeType() bool`

HasDistributeType returns a boolean if a field has been set.

### GetDuration

`func (o *AllocationAllocation) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AllocationAllocation) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AllocationAllocation) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AllocationAllocation) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEndedAt

`func (o *AllocationAllocation) GetEndedAt() map[string]interface{}`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *AllocationAllocation) GetEndedAtOk() (*map[string]interface{}, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *AllocationAllocation) SetEndedAt(v map[string]interface{})`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *AllocationAllocation) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetHoursPerDay

`func (o *AllocationAllocation) GetHoursPerDay() float32`

GetHoursPerDay returns the HoursPerDay field if non-nil, zero value otherwise.

### GetHoursPerDayOk

`func (o *AllocationAllocation) GetHoursPerDayOk() (*float32, bool)`

GetHoursPerDayOk returns a tuple with the HoursPerDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursPerDay

`func (o *AllocationAllocation) SetHoursPerDay(v float32)`

SetHoursPerDay sets HoursPerDay field to given value.

### HasHoursPerDay

`func (o *AllocationAllocation) HasHoursPerDay() bool`

HasHoursPerDay returns a boolean if a field has been set.

### GetIgnoreCollisions

`func (o *AllocationAllocation) GetIgnoreCollisions() bool`

GetIgnoreCollisions returns the IgnoreCollisions field if non-nil, zero value otherwise.

### GetIgnoreCollisionsOk

`func (o *AllocationAllocation) GetIgnoreCollisionsOk() (*bool, bool)`

GetIgnoreCollisionsOk returns a tuple with the IgnoreCollisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreCollisions

`func (o *AllocationAllocation) SetIgnoreCollisions(v bool)`

SetIgnoreCollisions sets IgnoreCollisions field to given value.

### HasIgnoreCollisions

`func (o *AllocationAllocation) HasIgnoreCollisions() bool`

HasIgnoreCollisions returns a boolean if a field has been set.

### GetProjectId

`func (o *AllocationAllocation) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AllocationAllocation) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AllocationAllocation) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *AllocationAllocation) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetRecurringRule

`func (o *AllocationAllocation) GetRecurringRule() PayloadNullableRRule`

GetRecurringRule returns the RecurringRule field if non-nil, zero value otherwise.

### GetRecurringRuleOk

`func (o *AllocationAllocation) GetRecurringRuleOk() (*PayloadNullableRRule, bool)`

GetRecurringRuleOk returns a tuple with the RecurringRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringRule

`func (o *AllocationAllocation) SetRecurringRule(v PayloadNullableRRule)`

SetRecurringRule sets RecurringRule field to given value.

### HasRecurringRule

`func (o *AllocationAllocation) HasRecurringRule() bool`

HasRecurringRule returns a boolean if a field has been set.

### GetStartedAt

`func (o *AllocationAllocation) GetStartedAt() map[string]interface{}`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *AllocationAllocation) GetStartedAtOk() (*map[string]interface{}, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *AllocationAllocation) SetStartedAt(v map[string]interface{})`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *AllocationAllocation) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetTitle

`func (o *AllocationAllocation) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AllocationAllocation) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AllocationAllocation) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AllocationAllocation) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


