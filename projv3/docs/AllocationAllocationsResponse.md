# AllocationAllocationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allocations** | Pointer to [**[]ViewAllocation**](ViewAllocation.md) |  | [optional] 
**Included** | Pointer to [**ActivityActivitiesResponseIncluded**](ActivityActivitiesResponseIncluded.md) |  | [optional] 
**Meta** | Pointer to [**ViewMeta**](ViewMeta.md) |  | [optional] 

## Methods

### NewAllocationAllocationsResponse

`func NewAllocationAllocationsResponse() *AllocationAllocationsResponse`

NewAllocationAllocationsResponse instantiates a new AllocationAllocationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllocationAllocationsResponseWithDefaults

`func NewAllocationAllocationsResponseWithDefaults() *AllocationAllocationsResponse`

NewAllocationAllocationsResponseWithDefaults instantiates a new AllocationAllocationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocations

`func (o *AllocationAllocationsResponse) GetAllocations() []ViewAllocation`

GetAllocations returns the Allocations field if non-nil, zero value otherwise.

### GetAllocationsOk

`func (o *AllocationAllocationsResponse) GetAllocationsOk() (*[]ViewAllocation, bool)`

GetAllocationsOk returns a tuple with the Allocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocations

`func (o *AllocationAllocationsResponse) SetAllocations(v []ViewAllocation)`

SetAllocations sets Allocations field to given value.

### HasAllocations

`func (o *AllocationAllocationsResponse) HasAllocations() bool`

HasAllocations returns a boolean if a field has been set.

### GetIncluded

`func (o *AllocationAllocationsResponse) GetIncluded() ActivityActivitiesResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AllocationAllocationsResponse) GetIncludedOk() (*ActivityActivitiesResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AllocationAllocationsResponse) SetIncluded(v ActivityActivitiesResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AllocationAllocationsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *AllocationAllocationsResponse) GetMeta() ViewMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AllocationAllocationsResponse) GetMetaOk() (*ViewMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AllocationAllocationsResponse) SetMeta(v ViewMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AllocationAllocationsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


