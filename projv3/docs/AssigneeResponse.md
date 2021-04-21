# AssigneeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignees** | Pointer to [**[]ViewFormAssignee**](ViewFormAssignee.md) |  | [optional] 
**Errors** | Pointer to [**[]ErrorsBulkError**](ErrorsBulkError.md) |  | [optional] 
**Included** | Pointer to [**AssigneeFormAssigneesResponseIncluded**](AssigneeFormAssigneesResponseIncluded.md) |  | [optional] 

## Methods

### NewAssigneeResponse

`func NewAssigneeResponse() *AssigneeResponse`

NewAssigneeResponse instantiates a new AssigneeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssigneeResponseWithDefaults

`func NewAssigneeResponseWithDefaults() *AssigneeResponse`

NewAssigneeResponseWithDefaults instantiates a new AssigneeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignees

`func (o *AssigneeResponse) GetAssignees() []ViewFormAssignee`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *AssigneeResponse) GetAssigneesOk() (*[]ViewFormAssignee, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *AssigneeResponse) SetAssignees(v []ViewFormAssignee)`

SetAssignees sets Assignees field to given value.

### HasAssignees

`func (o *AssigneeResponse) HasAssignees() bool`

HasAssignees returns a boolean if a field has been set.

### GetErrors

`func (o *AssigneeResponse) GetErrors() []ErrorsBulkError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *AssigneeResponse) GetErrorsOk() (*[]ErrorsBulkError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *AssigneeResponse) SetErrors(v []ErrorsBulkError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *AssigneeResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetIncluded

`func (o *AssigneeResponse) GetIncluded() AssigneeFormAssigneesResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AssigneeResponse) GetIncludedOk() (*AssigneeFormAssigneesResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AssigneeResponse) SetIncluded(v AssigneeFormAssigneesResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AssigneeResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


