# AllocationSplitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **map[string]interface{}** | Date unmarshals represents a Unified API Spec date format. | [optional] 
**Duration** | Pointer to **int32** |  | [optional] 

## Methods

### NewAllocationSplitRequest

`func NewAllocationSplitRequest() *AllocationSplitRequest`

NewAllocationSplitRequest instantiates a new AllocationSplitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllocationSplitRequestWithDefaults

`func NewAllocationSplitRequestWithDefaults() *AllocationSplitRequest`

NewAllocationSplitRequestWithDefaults instantiates a new AllocationSplitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *AllocationSplitRequest) GetAt() map[string]interface{}`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *AllocationSplitRequest) GetAtOk() (*map[string]interface{}, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *AllocationSplitRequest) SetAt(v map[string]interface{})`

SetAt sets At field to given value.

### HasAt

`func (o *AllocationSplitRequest) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetDuration

`func (o *AllocationSplitRequest) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AllocationSplitRequest) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AllocationSplitRequest) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AllocationSplitRequest) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


