# UtilizationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Included** | Pointer to [**NotebookVersionResponseIncluded**](NotebookVersionResponseIncluded.md) |  | [optional] 
**Meta** | Pointer to [**ViewMeta**](ViewMeta.md) |  | [optional] 
**Utilization** | Pointer to [**[]ViewUserUtilization**](ViewUserUtilization.md) |  | [optional] 

## Methods

### NewUtilizationResponse

`func NewUtilizationResponse() *UtilizationResponse`

NewUtilizationResponse instantiates a new UtilizationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilizationResponseWithDefaults

`func NewUtilizationResponseWithDefaults() *UtilizationResponse`

NewUtilizationResponseWithDefaults instantiates a new UtilizationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncluded

`func (o *UtilizationResponse) GetIncluded() NotebookVersionResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *UtilizationResponse) GetIncludedOk() (*NotebookVersionResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *UtilizationResponse) SetIncluded(v NotebookVersionResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *UtilizationResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *UtilizationResponse) GetMeta() ViewMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *UtilizationResponse) GetMetaOk() (*ViewMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *UtilizationResponse) SetMeta(v ViewMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *UtilizationResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUtilization

`func (o *UtilizationResponse) GetUtilization() []ViewUserUtilization`

GetUtilization returns the Utilization field if non-nil, zero value otherwise.

### GetUtilizationOk

`func (o *UtilizationResponse) GetUtilizationOk() (*[]ViewUserUtilization, bool)`

GetUtilizationOk returns a tuple with the Utilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilization

`func (o *UtilizationResponse) SetUtilization(v []ViewUserUtilization)`

SetUtilization sets Utilization field to given value.

### HasUtilization

`func (o *UtilizationResponse) HasUtilization() bool`

HasUtilization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


