# DashboardUserDashboardsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dashboards** | Pointer to [**[]ViewUserDashboard**](ViewUserDashboard.md) |  | [optional] 
**Included** | Pointer to [**DashboardUserDashboardsResponseIncluded**](DashboardUserDashboardsResponseIncluded.md) |  | [optional] 
**Meta** | Pointer to [**ViewMeta**](ViewMeta.md) |  | [optional] 

## Methods

### NewDashboardUserDashboardsResponse

`func NewDashboardUserDashboardsResponse() *DashboardUserDashboardsResponse`

NewDashboardUserDashboardsResponse instantiates a new DashboardUserDashboardsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardUserDashboardsResponseWithDefaults

`func NewDashboardUserDashboardsResponseWithDefaults() *DashboardUserDashboardsResponse`

NewDashboardUserDashboardsResponseWithDefaults instantiates a new DashboardUserDashboardsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDashboards

`func (o *DashboardUserDashboardsResponse) GetDashboards() []ViewUserDashboard`

GetDashboards returns the Dashboards field if non-nil, zero value otherwise.

### GetDashboardsOk

`func (o *DashboardUserDashboardsResponse) GetDashboardsOk() (*[]ViewUserDashboard, bool)`

GetDashboardsOk returns a tuple with the Dashboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboards

`func (o *DashboardUserDashboardsResponse) SetDashboards(v []ViewUserDashboard)`

SetDashboards sets Dashboards field to given value.

### HasDashboards

`func (o *DashboardUserDashboardsResponse) HasDashboards() bool`

HasDashboards returns a boolean if a field has been set.

### GetIncluded

`func (o *DashboardUserDashboardsResponse) GetIncluded() DashboardUserDashboardsResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *DashboardUserDashboardsResponse) GetIncludedOk() (*DashboardUserDashboardsResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *DashboardUserDashboardsResponse) SetIncluded(v DashboardUserDashboardsResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *DashboardUserDashboardsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *DashboardUserDashboardsResponse) GetMeta() ViewMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DashboardUserDashboardsResponse) GetMetaOk() (*ViewMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DashboardUserDashboardsResponse) SetMeta(v ViewMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *DashboardUserDashboardsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


