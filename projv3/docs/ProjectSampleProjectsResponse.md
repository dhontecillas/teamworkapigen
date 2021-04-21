# ProjectSampleProjectsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Included** | Pointer to [**ProjectSampleProjectsResponseIncluded**](ProjectSampleProjectsResponseIncluded.md) |  | [optional] 
**Meta** | Pointer to [**ViewMeta**](ViewMeta.md) |  | [optional] 
**Projects** | Pointer to [**[]ViewSampleProject**](ViewSampleProject.md) |  | [optional] 

## Methods

### NewProjectSampleProjectsResponse

`func NewProjectSampleProjectsResponse() *ProjectSampleProjectsResponse`

NewProjectSampleProjectsResponse instantiates a new ProjectSampleProjectsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectSampleProjectsResponseWithDefaults

`func NewProjectSampleProjectsResponseWithDefaults() *ProjectSampleProjectsResponse`

NewProjectSampleProjectsResponseWithDefaults instantiates a new ProjectSampleProjectsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncluded

`func (o *ProjectSampleProjectsResponse) GetIncluded() ProjectSampleProjectsResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *ProjectSampleProjectsResponse) GetIncludedOk() (*ProjectSampleProjectsResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *ProjectSampleProjectsResponse) SetIncluded(v ProjectSampleProjectsResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *ProjectSampleProjectsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *ProjectSampleProjectsResponse) GetMeta() ViewMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ProjectSampleProjectsResponse) GetMetaOk() (*ViewMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ProjectSampleProjectsResponse) SetMeta(v ViewMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ProjectSampleProjectsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetProjects

`func (o *ProjectSampleProjectsResponse) GetProjects() []ViewSampleProject`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *ProjectSampleProjectsResponse) GetProjectsOk() (*[]ViewSampleProject, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *ProjectSampleProjectsResponse) SetProjects(v []ViewSampleProject)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *ProjectSampleProjectsResponse) HasProjects() bool`

HasProjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


