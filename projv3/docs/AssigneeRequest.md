# AssigneeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignees** | Pointer to [**[]AssigneeFormAssignee**](AssigneeFormAssignee.md) |  | [optional] 

## Methods

### NewAssigneeRequest

`func NewAssigneeRequest() *AssigneeRequest`

NewAssigneeRequest instantiates a new AssigneeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssigneeRequestWithDefaults

`func NewAssigneeRequestWithDefaults() *AssigneeRequest`

NewAssigneeRequestWithDefaults instantiates a new AssigneeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignees

`func (o *AssigneeRequest) GetAssignees() []AssigneeFormAssignee`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *AssigneeRequest) GetAssigneesOk() (*[]AssigneeFormAssignee, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *AssigneeRequest) SetAssignees(v []AssigneeFormAssignee)`

SetAssignees sets Assignees field to given value.

### HasAssignees

`func (o *AssigneeRequest) HasAssignees() bool`

HasAssignees returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


