# ViewAllocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedUser** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**AssignedUserID** | Pointer to **int32** |  | [optional] 
**AvailableDuration** | Pointer to **int32** | in minutes | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **int32** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**DeletedBy** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DistributeType** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **int32** | in minutes | [optional] 
**EndedAt** | Pointer to **map[string]interface{}** | Date represents a Unified API Spec date format. | [optional] 
**HoursPerDay** | Pointer to **float32** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Installation** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**InstallationId** | Pointer to **int32** |  | [optional] 
**Project** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**RecurringRule** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **map[string]interface{}** | Date represents a Unified API Spec date format. | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **int32** |  | [optional] 

## Methods

### NewViewAllocation

`func NewViewAllocation() *ViewAllocation`

NewViewAllocation instantiates a new ViewAllocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewAllocationWithDefaults

`func NewViewAllocationWithDefaults() *ViewAllocation`

NewViewAllocationWithDefaults instantiates a new ViewAllocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedUser

`func (o *ViewAllocation) GetAssignedUser() ViewRelationship`

GetAssignedUser returns the AssignedUser field if non-nil, zero value otherwise.

### GetAssignedUserOk

`func (o *ViewAllocation) GetAssignedUserOk() (*ViewRelationship, bool)`

GetAssignedUserOk returns a tuple with the AssignedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUser

`func (o *ViewAllocation) SetAssignedUser(v ViewRelationship)`

SetAssignedUser sets AssignedUser field to given value.

### HasAssignedUser

`func (o *ViewAllocation) HasAssignedUser() bool`

HasAssignedUser returns a boolean if a field has been set.

### GetAssignedUserID

`func (o *ViewAllocation) GetAssignedUserID() int32`

GetAssignedUserID returns the AssignedUserID field if non-nil, zero value otherwise.

### GetAssignedUserIDOk

`func (o *ViewAllocation) GetAssignedUserIDOk() (*int32, bool)`

GetAssignedUserIDOk returns a tuple with the AssignedUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUserID

`func (o *ViewAllocation) SetAssignedUserID(v int32)`

SetAssignedUserID sets AssignedUserID field to given value.

### HasAssignedUserID

`func (o *ViewAllocation) HasAssignedUserID() bool`

HasAssignedUserID returns a boolean if a field has been set.

### GetAvailableDuration

`func (o *ViewAllocation) GetAvailableDuration() int32`

GetAvailableDuration returns the AvailableDuration field if non-nil, zero value otherwise.

### GetAvailableDurationOk

`func (o *ViewAllocation) GetAvailableDurationOk() (*int32, bool)`

GetAvailableDurationOk returns a tuple with the AvailableDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableDuration

`func (o *ViewAllocation) SetAvailableDuration(v int32)`

SetAvailableDuration sets AvailableDuration field to given value.

### HasAvailableDuration

`func (o *ViewAllocation) HasAvailableDuration() bool`

HasAvailableDuration returns a boolean if a field has been set.

### GetColor

`func (o *ViewAllocation) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ViewAllocation) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ViewAllocation) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *ViewAllocation) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ViewAllocation) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ViewAllocation) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ViewAllocation) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ViewAllocation) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ViewAllocation) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ViewAllocation) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ViewAllocation) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ViewAllocation) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ViewAllocation) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ViewAllocation) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ViewAllocation) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ViewAllocation) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeletedBy

`func (o *ViewAllocation) GetDeletedBy() int32`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *ViewAllocation) GetDeletedByOk() (*int32, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *ViewAllocation) SetDeletedBy(v int32)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *ViewAllocation) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetDescription

`func (o *ViewAllocation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ViewAllocation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ViewAllocation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ViewAllocation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDistributeType

`func (o *ViewAllocation) GetDistributeType() string`

GetDistributeType returns the DistributeType field if non-nil, zero value otherwise.

### GetDistributeTypeOk

`func (o *ViewAllocation) GetDistributeTypeOk() (*string, bool)`

GetDistributeTypeOk returns a tuple with the DistributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributeType

`func (o *ViewAllocation) SetDistributeType(v string)`

SetDistributeType sets DistributeType field to given value.

### HasDistributeType

`func (o *ViewAllocation) HasDistributeType() bool`

HasDistributeType returns a boolean if a field has been set.

### GetDuration

`func (o *ViewAllocation) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ViewAllocation) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ViewAllocation) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ViewAllocation) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEndedAt

`func (o *ViewAllocation) GetEndedAt() map[string]interface{}`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *ViewAllocation) GetEndedAtOk() (*map[string]interface{}, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *ViewAllocation) SetEndedAt(v map[string]interface{})`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *ViewAllocation) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetHoursPerDay

`func (o *ViewAllocation) GetHoursPerDay() float32`

GetHoursPerDay returns the HoursPerDay field if non-nil, zero value otherwise.

### GetHoursPerDayOk

`func (o *ViewAllocation) GetHoursPerDayOk() (*float32, bool)`

GetHoursPerDayOk returns a tuple with the HoursPerDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursPerDay

`func (o *ViewAllocation) SetHoursPerDay(v float32)`

SetHoursPerDay sets HoursPerDay field to given value.

### HasHoursPerDay

`func (o *ViewAllocation) HasHoursPerDay() bool`

HasHoursPerDay returns a boolean if a field has been set.

### GetId

`func (o *ViewAllocation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewAllocation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewAllocation) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewAllocation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstallation

`func (o *ViewAllocation) GetInstallation() ViewRelationship`

GetInstallation returns the Installation field if non-nil, zero value otherwise.

### GetInstallationOk

`func (o *ViewAllocation) GetInstallationOk() (*ViewRelationship, bool)`

GetInstallationOk returns a tuple with the Installation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallation

`func (o *ViewAllocation) SetInstallation(v ViewRelationship)`

SetInstallation sets Installation field to given value.

### HasInstallation

`func (o *ViewAllocation) HasInstallation() bool`

HasInstallation returns a boolean if a field has been set.

### GetInstallationId

`func (o *ViewAllocation) GetInstallationId() int32`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *ViewAllocation) GetInstallationIdOk() (*int32, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *ViewAllocation) SetInstallationId(v int32)`

SetInstallationId sets InstallationId field to given value.

### HasInstallationId

`func (o *ViewAllocation) HasInstallationId() bool`

HasInstallationId returns a boolean if a field has been set.

### GetProject

`func (o *ViewAllocation) GetProject() ViewRelationship`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ViewAllocation) GetProjectOk() (*ViewRelationship, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ViewAllocation) SetProject(v ViewRelationship)`

SetProject sets Project field to given value.

### HasProject

`func (o *ViewAllocation) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectId

`func (o *ViewAllocation) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ViewAllocation) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ViewAllocation) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ViewAllocation) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetRecurringRule

`func (o *ViewAllocation) GetRecurringRule() string`

GetRecurringRule returns the RecurringRule field if non-nil, zero value otherwise.

### GetRecurringRuleOk

`func (o *ViewAllocation) GetRecurringRuleOk() (*string, bool)`

GetRecurringRuleOk returns a tuple with the RecurringRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringRule

`func (o *ViewAllocation) SetRecurringRule(v string)`

SetRecurringRule sets RecurringRule field to given value.

### HasRecurringRule

`func (o *ViewAllocation) HasRecurringRule() bool`

HasRecurringRule returns a boolean if a field has been set.

### GetStartedAt

`func (o *ViewAllocation) GetStartedAt() map[string]interface{}`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ViewAllocation) GetStartedAtOk() (*map[string]interface{}, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ViewAllocation) SetStartedAt(v map[string]interface{})`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ViewAllocation) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *ViewAllocation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ViewAllocation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ViewAllocation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ViewAllocation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *ViewAllocation) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ViewAllocation) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ViewAllocation) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ViewAllocation) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ViewAllocation) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ViewAllocation) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ViewAllocation) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ViewAllocation) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ViewAllocation) GetUpdatedBy() int32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ViewAllocation) GetUpdatedByOk() (*int32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ViewAllocation) SetUpdatedBy(v int32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ViewAllocation) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


