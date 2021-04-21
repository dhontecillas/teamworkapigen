# PlannerWorkloadPlanner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacities** | Pointer to [**map[string]PlannerWorkloadPlannerCapacity**](PlannerWorkloadPlannerCapacity.md) |  | [optional] 
**User** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 

## Methods

### NewPlannerWorkloadPlanner

`func NewPlannerWorkloadPlanner() *PlannerWorkloadPlanner`

NewPlannerWorkloadPlanner instantiates a new PlannerWorkloadPlanner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlannerWorkloadPlannerWithDefaults

`func NewPlannerWorkloadPlannerWithDefaults() *PlannerWorkloadPlanner`

NewPlannerWorkloadPlannerWithDefaults instantiates a new PlannerWorkloadPlanner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacities

`func (o *PlannerWorkloadPlanner) GetCapacities() map[string]PlannerWorkloadPlannerCapacity`

GetCapacities returns the Capacities field if non-nil, zero value otherwise.

### GetCapacitiesOk

`func (o *PlannerWorkloadPlanner) GetCapacitiesOk() (*map[string]PlannerWorkloadPlannerCapacity, bool)`

GetCapacitiesOk returns a tuple with the Capacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacities

`func (o *PlannerWorkloadPlanner) SetCapacities(v map[string]PlannerWorkloadPlannerCapacity)`

SetCapacities sets Capacities field to given value.

### HasCapacities

`func (o *PlannerWorkloadPlanner) HasCapacities() bool`

HasCapacities returns a boolean if a field has been set.

### GetUser

`func (o *PlannerWorkloadPlanner) GetUser() ViewRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PlannerWorkloadPlanner) GetUserOk() (*ViewRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PlannerWorkloadPlanner) SetUser(v ViewRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *PlannerWorkloadPlanner) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserId

`func (o *PlannerWorkloadPlanner) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PlannerWorkloadPlanner) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PlannerWorkloadPlanner) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PlannerWorkloadPlanner) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


