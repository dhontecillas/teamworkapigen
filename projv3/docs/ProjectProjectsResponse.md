# ProjectProjectsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Included** | Pointer to [**ProjectProjectsResponseIncluded**](ProjectProjectsResponseIncluded.md) |  | [optional] 
**Meta** | Pointer to [**ViewMeta**](ViewMeta.md) |  | [optional] 
**Projects** | Pointer to [**[]ViewProject**](ViewProject.md) |  | [optional] 

## Methods

### NewProjectProjectsResponse

`func NewProjectProjectsResponse() *ProjectProjectsResponse`

NewProjectProjectsResponse instantiates a new ProjectProjectsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectProjectsResponseWithDefaults

`func NewProjectProjectsResponseWithDefaults() *ProjectProjectsResponse`

NewProjectProjectsResponseWithDefaults instantiates a new ProjectProjectsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncluded

`func (o *ProjectProjectsResponse) GetIncluded() ProjectProjectsResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *ProjectProjectsResponse) GetIncludedOk() (*ProjectProjectsResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *ProjectProjectsResponse) SetIncluded(v ProjectProjectsResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *ProjectProjectsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *ProjectProjectsResponse) GetMeta() ViewMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ProjectProjectsResponse) GetMetaOk() (*ViewMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ProjectProjectsResponse) SetMeta(v ViewMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ProjectProjectsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetProjects

`func (o *ProjectProjectsResponse) GetProjects() []ViewProject`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *ProjectProjectsResponse) GetProjectsOk() (*[]ViewProject, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *ProjectProjectsResponse) SetProjects(v []ViewProject)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *ProjectProjectsResponse) HasProjects() bool`

HasProjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


