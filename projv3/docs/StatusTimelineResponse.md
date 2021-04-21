# StatusTimelineResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Included** | Pointer to [**StatusTimelineResponseIncluded**](StatusTimelineResponseIncluded.md) |  | [optional] 
**Meta** | Pointer to [**ViewMeta**](ViewMeta.md) |  | [optional] 
**Statuses** | Pointer to [**[]ViewStatus**](ViewStatus.md) |  | [optional] 

## Methods

### NewStatusTimelineResponse

`func NewStatusTimelineResponse() *StatusTimelineResponse`

NewStatusTimelineResponse instantiates a new StatusTimelineResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusTimelineResponseWithDefaults

`func NewStatusTimelineResponseWithDefaults() *StatusTimelineResponse`

NewStatusTimelineResponseWithDefaults instantiates a new StatusTimelineResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncluded

`func (o *StatusTimelineResponse) GetIncluded() StatusTimelineResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *StatusTimelineResponse) GetIncludedOk() (*StatusTimelineResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *StatusTimelineResponse) SetIncluded(v StatusTimelineResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *StatusTimelineResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *StatusTimelineResponse) GetMeta() ViewMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *StatusTimelineResponse) GetMetaOk() (*ViewMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *StatusTimelineResponse) SetMeta(v ViewMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *StatusTimelineResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetStatuses

`func (o *StatusTimelineResponse) GetStatuses() []ViewStatus`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *StatusTimelineResponse) GetStatusesOk() (*[]ViewStatus, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *StatusTimelineResponse) SetStatuses(v []ViewStatus)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *StatusTimelineResponse) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


