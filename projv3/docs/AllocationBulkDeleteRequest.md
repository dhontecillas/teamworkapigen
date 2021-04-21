# AllocationBulkDeleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocationIds** | Pointer to **[]int32** |  | [optional] 
**HardDelete** | Pointer to **bool** |  | [optional] 

## Methods

### NewAllocationBulkDeleteRequest

`func NewAllocationBulkDeleteRequest() *AllocationBulkDeleteRequest`

NewAllocationBulkDeleteRequest instantiates a new AllocationBulkDeleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllocationBulkDeleteRequestWithDefaults

`func NewAllocationBulkDeleteRequestWithDefaults() *AllocationBulkDeleteRequest`

NewAllocationBulkDeleteRequestWithDefaults instantiates a new AllocationBulkDeleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocationIds

`func (o *AllocationBulkDeleteRequest) GetAllocationIds() []int32`

GetAllocationIds returns the AllocationIds field if non-nil, zero value otherwise.

### GetAllocationIdsOk

`func (o *AllocationBulkDeleteRequest) GetAllocationIdsOk() (*[]int32, bool)`

GetAllocationIdsOk returns a tuple with the AllocationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationIds

`func (o *AllocationBulkDeleteRequest) SetAllocationIds(v []int32)`

SetAllocationIds sets AllocationIds field to given value.

### HasAllocationIds

`func (o *AllocationBulkDeleteRequest) HasAllocationIds() bool`

HasAllocationIds returns a boolean if a field has been set.

### GetHardDelete

`func (o *AllocationBulkDeleteRequest) GetHardDelete() bool`

GetHardDelete returns the HardDelete field if non-nil, zero value otherwise.

### GetHardDeleteOk

`func (o *AllocationBulkDeleteRequest) GetHardDeleteOk() (*bool, bool)`

GetHardDeleteOk returns a tuple with the HardDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardDelete

`func (o *AllocationBulkDeleteRequest) SetHardDelete(v bool)`

SetHardDelete sets HardDelete field to given value.

### HasHardDelete

`func (o *AllocationBulkDeleteRequest) HasHardDelete() bool`

HasHardDelete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


