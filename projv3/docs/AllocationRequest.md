# AllocationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allocation** | Pointer to [**AllocationAllocation**](AllocationAllocation.md) |  | [optional] 

## Methods

### NewAllocationRequest

`func NewAllocationRequest() *AllocationRequest`

NewAllocationRequest instantiates a new AllocationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllocationRequestWithDefaults

`func NewAllocationRequestWithDefaults() *AllocationRequest`

NewAllocationRequestWithDefaults instantiates a new AllocationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocation

`func (o *AllocationRequest) GetAllocation() AllocationAllocation`

GetAllocation returns the Allocation field if non-nil, zero value otherwise.

### GetAllocationOk

`func (o *AllocationRequest) GetAllocationOk() (*AllocationAllocation, bool)`

GetAllocationOk returns a tuple with the Allocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocation

`func (o *AllocationRequest) SetAllocation(v AllocationAllocation)`

SetAllocation sets Allocation field to given value.

### HasAllocation

`func (o *AllocationRequest) HasAllocation() bool`

HasAllocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


