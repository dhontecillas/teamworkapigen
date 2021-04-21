# PlannerWorkloadPlannersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Included** | Pointer to [**PlannerWorkloadPlannersResponseIncluded**](PlannerWorkloadPlannersResponseIncluded.md) |  | [optional] 
**Meta** | Pointer to [**ViewMeta**](ViewMeta.md) |  | [optional] 
**Planners** | Pointer to [**[]PlannerWorkloadPlanner**](PlannerWorkloadPlanner.md) |  | [optional] 

## Methods

### NewPlannerWorkloadPlannersResponse

`func NewPlannerWorkloadPlannersResponse() *PlannerWorkloadPlannersResponse`

NewPlannerWorkloadPlannersResponse instantiates a new PlannerWorkloadPlannersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlannerWorkloadPlannersResponseWithDefaults

`func NewPlannerWorkloadPlannersResponseWithDefaults() *PlannerWorkloadPlannersResponse`

NewPlannerWorkloadPlannersResponseWithDefaults instantiates a new PlannerWorkloadPlannersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncluded

`func (o *PlannerWorkloadPlannersResponse) GetIncluded() PlannerWorkloadPlannersResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *PlannerWorkloadPlannersResponse) GetIncludedOk() (*PlannerWorkloadPlannersResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *PlannerWorkloadPlannersResponse) SetIncluded(v PlannerWorkloadPlannersResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *PlannerWorkloadPlannersResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *PlannerWorkloadPlannersResponse) GetMeta() ViewMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PlannerWorkloadPlannersResponse) GetMetaOk() (*ViewMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PlannerWorkloadPlannersResponse) SetMeta(v ViewMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PlannerWorkloadPlannersResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetPlanners

`func (o *PlannerWorkloadPlannersResponse) GetPlanners() []PlannerWorkloadPlanner`

GetPlanners returns the Planners field if non-nil, zero value otherwise.

### GetPlannersOk

`func (o *PlannerWorkloadPlannersResponse) GetPlannersOk() (*[]PlannerWorkloadPlanner, bool)`

GetPlannersOk returns a tuple with the Planners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanners

`func (o *PlannerWorkloadPlannersResponse) SetPlanners(v []PlannerWorkloadPlanner)`

SetPlanners sets Planners field to given value.

### HasPlanners

`func (o *PlannerWorkloadPlannersResponse) HasPlanners() bool`

HasPlanners returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


