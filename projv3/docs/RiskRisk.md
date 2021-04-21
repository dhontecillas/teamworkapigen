# RiskRisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanEdit** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **int32** |  | [optional] 
**CreatedByUserId** | Pointer to **int32** |  | [optional] 
**CreatedOn** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Impact** | Pointer to **string** |  | [optional] 
**ImpactCost** | Pointer to **bool** |  | [optional] 
**ImpactPerformance** | Pointer to **bool** |  | [optional] 
**ImpactSchedule** | Pointer to **bool** |  | [optional] 
**ImpactValue** | Pointer to **int32** |  | [optional] 
**LastChangedByUserId** | Pointer to **int32** |  | [optional] 
**LastChangedOn** | Pointer to **string** |  | [optional] 
**MitigationPlan** | Pointer to **string** |  | [optional] 
**Probability** | Pointer to **string** |  | [optional] 
**ProbabilityValue** | Pointer to **int32** |  | [optional] 
**Project** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**Result** | Pointer to **int32** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **int32** |  | [optional] 

## Methods

### NewRiskRisk

`func NewRiskRisk() *RiskRisk`

NewRiskRisk instantiates a new RiskRisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskRiskWithDefaults

`func NewRiskRiskWithDefaults() *RiskRisk`

NewRiskRiskWithDefaults instantiates a new RiskRisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanEdit

`func (o *RiskRisk) GetCanEdit() bool`

GetCanEdit returns the CanEdit field if non-nil, zero value otherwise.

### GetCanEditOk

`func (o *RiskRisk) GetCanEditOk() (*bool, bool)`

GetCanEditOk returns a tuple with the CanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEdit

`func (o *RiskRisk) SetCanEdit(v bool)`

SetCanEdit sets CanEdit field to given value.

### HasCanEdit

`func (o *RiskRisk) HasCanEdit() bool`

HasCanEdit returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RiskRisk) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RiskRisk) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RiskRisk) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RiskRisk) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *RiskRisk) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RiskRisk) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RiskRisk) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *RiskRisk) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *RiskRisk) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *RiskRisk) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *RiskRisk) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *RiskRisk) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetCreatedOn

`func (o *RiskRisk) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *RiskRisk) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *RiskRisk) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *RiskRisk) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetDeleted

`func (o *RiskRisk) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *RiskRisk) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *RiskRisk) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *RiskRisk) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetId

`func (o *RiskRisk) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskRisk) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskRisk) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RiskRisk) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImpact

`func (o *RiskRisk) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *RiskRisk) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *RiskRisk) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *RiskRisk) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetImpactCost

`func (o *RiskRisk) GetImpactCost() bool`

GetImpactCost returns the ImpactCost field if non-nil, zero value otherwise.

### GetImpactCostOk

`func (o *RiskRisk) GetImpactCostOk() (*bool, bool)`

GetImpactCostOk returns a tuple with the ImpactCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactCost

`func (o *RiskRisk) SetImpactCost(v bool)`

SetImpactCost sets ImpactCost field to given value.

### HasImpactCost

`func (o *RiskRisk) HasImpactCost() bool`

HasImpactCost returns a boolean if a field has been set.

### GetImpactPerformance

`func (o *RiskRisk) GetImpactPerformance() bool`

GetImpactPerformance returns the ImpactPerformance field if non-nil, zero value otherwise.

### GetImpactPerformanceOk

`func (o *RiskRisk) GetImpactPerformanceOk() (*bool, bool)`

GetImpactPerformanceOk returns a tuple with the ImpactPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactPerformance

`func (o *RiskRisk) SetImpactPerformance(v bool)`

SetImpactPerformance sets ImpactPerformance field to given value.

### HasImpactPerformance

`func (o *RiskRisk) HasImpactPerformance() bool`

HasImpactPerformance returns a boolean if a field has been set.

### GetImpactSchedule

`func (o *RiskRisk) GetImpactSchedule() bool`

GetImpactSchedule returns the ImpactSchedule field if non-nil, zero value otherwise.

### GetImpactScheduleOk

`func (o *RiskRisk) GetImpactScheduleOk() (*bool, bool)`

GetImpactScheduleOk returns a tuple with the ImpactSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactSchedule

`func (o *RiskRisk) SetImpactSchedule(v bool)`

SetImpactSchedule sets ImpactSchedule field to given value.

### HasImpactSchedule

`func (o *RiskRisk) HasImpactSchedule() bool`

HasImpactSchedule returns a boolean if a field has been set.

### GetImpactValue

`func (o *RiskRisk) GetImpactValue() int32`

GetImpactValue returns the ImpactValue field if non-nil, zero value otherwise.

### GetImpactValueOk

`func (o *RiskRisk) GetImpactValueOk() (*int32, bool)`

GetImpactValueOk returns a tuple with the ImpactValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactValue

`func (o *RiskRisk) SetImpactValue(v int32)`

SetImpactValue sets ImpactValue field to given value.

### HasImpactValue

`func (o *RiskRisk) HasImpactValue() bool`

HasImpactValue returns a boolean if a field has been set.

### GetLastChangedByUserId

`func (o *RiskRisk) GetLastChangedByUserId() int32`

GetLastChangedByUserId returns the LastChangedByUserId field if non-nil, zero value otherwise.

### GetLastChangedByUserIdOk

`func (o *RiskRisk) GetLastChangedByUserIdOk() (*int32, bool)`

GetLastChangedByUserIdOk returns a tuple with the LastChangedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangedByUserId

`func (o *RiskRisk) SetLastChangedByUserId(v int32)`

SetLastChangedByUserId sets LastChangedByUserId field to given value.

### HasLastChangedByUserId

`func (o *RiskRisk) HasLastChangedByUserId() bool`

HasLastChangedByUserId returns a boolean if a field has been set.

### GetLastChangedOn

`func (o *RiskRisk) GetLastChangedOn() string`

GetLastChangedOn returns the LastChangedOn field if non-nil, zero value otherwise.

### GetLastChangedOnOk

`func (o *RiskRisk) GetLastChangedOnOk() (*string, bool)`

GetLastChangedOnOk returns a tuple with the LastChangedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangedOn

`func (o *RiskRisk) SetLastChangedOn(v string)`

SetLastChangedOn sets LastChangedOn field to given value.

### HasLastChangedOn

`func (o *RiskRisk) HasLastChangedOn() bool`

HasLastChangedOn returns a boolean if a field has been set.

### GetMitigationPlan

`func (o *RiskRisk) GetMitigationPlan() string`

GetMitigationPlan returns the MitigationPlan field if non-nil, zero value otherwise.

### GetMitigationPlanOk

`func (o *RiskRisk) GetMitigationPlanOk() (*string, bool)`

GetMitigationPlanOk returns a tuple with the MitigationPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigationPlan

`func (o *RiskRisk) SetMitigationPlan(v string)`

SetMitigationPlan sets MitigationPlan field to given value.

### HasMitigationPlan

`func (o *RiskRisk) HasMitigationPlan() bool`

HasMitigationPlan returns a boolean if a field has been set.

### GetProbability

`func (o *RiskRisk) GetProbability() string`

GetProbability returns the Probability field if non-nil, zero value otherwise.

### GetProbabilityOk

`func (o *RiskRisk) GetProbabilityOk() (*string, bool)`

GetProbabilityOk returns a tuple with the Probability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbability

`func (o *RiskRisk) SetProbability(v string)`

SetProbability sets Probability field to given value.

### HasProbability

`func (o *RiskRisk) HasProbability() bool`

HasProbability returns a boolean if a field has been set.

### GetProbabilityValue

`func (o *RiskRisk) GetProbabilityValue() int32`

GetProbabilityValue returns the ProbabilityValue field if non-nil, zero value otherwise.

### GetProbabilityValueOk

`func (o *RiskRisk) GetProbabilityValueOk() (*int32, bool)`

GetProbabilityValueOk returns a tuple with the ProbabilityValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbabilityValue

`func (o *RiskRisk) SetProbabilityValue(v int32)`

SetProbabilityValue sets ProbabilityValue field to given value.

### HasProbabilityValue

`func (o *RiskRisk) HasProbabilityValue() bool`

HasProbabilityValue returns a boolean if a field has been set.

### GetProject

`func (o *RiskRisk) GetProject() ViewRelationship`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *RiskRisk) GetProjectOk() (*ViewRelationship, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *RiskRisk) SetProject(v ViewRelationship)`

SetProject sets Project field to given value.

### HasProject

`func (o *RiskRisk) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectId

`func (o *RiskRisk) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *RiskRisk) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *RiskRisk) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *RiskRisk) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetResult

`func (o *RiskRisk) GetResult() int32`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *RiskRisk) GetResultOk() (*int32, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *RiskRisk) SetResult(v int32)`

SetResult sets Result field to given value.

### HasResult

`func (o *RiskRisk) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetSource

`func (o *RiskRisk) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RiskRisk) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RiskRisk) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *RiskRisk) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *RiskRisk) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RiskRisk) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RiskRisk) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RiskRisk) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RiskRisk) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RiskRisk) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RiskRisk) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RiskRisk) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *RiskRisk) GetUpdatedBy() int32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *RiskRisk) GetUpdatedByOk() (*int32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *RiskRisk) SetUpdatedBy(v int32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *RiskRisk) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


