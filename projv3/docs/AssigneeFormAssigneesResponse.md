# AssigneeFormAssigneesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignees** | Pointer to [**[]ViewFormAssignee**](ViewFormAssignee.md) |  | [optional] 
**Included** | Pointer to [**AssigneeFormAssigneesResponseIncluded**](AssigneeFormAssigneesResponseIncluded.md) |  | [optional] 
**Meta** | Pointer to [**ViewMeta**](ViewMeta.md) |  | [optional] 

## Methods

### NewAssigneeFormAssigneesResponse

`func NewAssigneeFormAssigneesResponse() *AssigneeFormAssigneesResponse`

NewAssigneeFormAssigneesResponse instantiates a new AssigneeFormAssigneesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssigneeFormAssigneesResponseWithDefaults

`func NewAssigneeFormAssigneesResponseWithDefaults() *AssigneeFormAssigneesResponse`

NewAssigneeFormAssigneesResponseWithDefaults instantiates a new AssigneeFormAssigneesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignees

`func (o *AssigneeFormAssigneesResponse) GetAssignees() []ViewFormAssignee`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *AssigneeFormAssigneesResponse) GetAssigneesOk() (*[]ViewFormAssignee, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *AssigneeFormAssigneesResponse) SetAssignees(v []ViewFormAssignee)`

SetAssignees sets Assignees field to given value.

### HasAssignees

`func (o *AssigneeFormAssigneesResponse) HasAssignees() bool`

HasAssignees returns a boolean if a field has been set.

### GetIncluded

`func (o *AssigneeFormAssigneesResponse) GetIncluded() AssigneeFormAssigneesResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AssigneeFormAssigneesResponse) GetIncludedOk() (*AssigneeFormAssigneesResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AssigneeFormAssigneesResponse) SetIncluded(v AssigneeFormAssigneesResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AssigneeFormAssigneesResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *AssigneeFormAssigneesResponse) GetMeta() ViewMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AssigneeFormAssigneesResponse) GetMetaOk() (*ViewMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AssigneeFormAssigneesResponse) SetMeta(v ViewMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AssigneeFormAssigneesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


