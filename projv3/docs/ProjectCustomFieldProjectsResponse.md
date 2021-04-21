# ProjectCustomFieldProjectsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomfieldProjects** | Pointer to [**[]ViewCustomFieldProject**](ViewCustomFieldProject.md) |  | [optional] 
**Included** | Pointer to [**ProjectCustomFieldProjectsResponseIncluded**](ProjectCustomFieldProjectsResponseIncluded.md) |  | [optional] 
**Meta** | Pointer to [**ViewMeta**](ViewMeta.md) |  | [optional] 

## Methods

### NewProjectCustomFieldProjectsResponse

`func NewProjectCustomFieldProjectsResponse() *ProjectCustomFieldProjectsResponse`

NewProjectCustomFieldProjectsResponse instantiates a new ProjectCustomFieldProjectsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectCustomFieldProjectsResponseWithDefaults

`func NewProjectCustomFieldProjectsResponseWithDefaults() *ProjectCustomFieldProjectsResponse`

NewProjectCustomFieldProjectsResponseWithDefaults instantiates a new ProjectCustomFieldProjectsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomfieldProjects

`func (o *ProjectCustomFieldProjectsResponse) GetCustomfieldProjects() []ViewCustomFieldProject`

GetCustomfieldProjects returns the CustomfieldProjects field if non-nil, zero value otherwise.

### GetCustomfieldProjectsOk

`func (o *ProjectCustomFieldProjectsResponse) GetCustomfieldProjectsOk() (*[]ViewCustomFieldProject, bool)`

GetCustomfieldProjectsOk returns a tuple with the CustomfieldProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomfieldProjects

`func (o *ProjectCustomFieldProjectsResponse) SetCustomfieldProjects(v []ViewCustomFieldProject)`

SetCustomfieldProjects sets CustomfieldProjects field to given value.

### HasCustomfieldProjects

`func (o *ProjectCustomFieldProjectsResponse) HasCustomfieldProjects() bool`

HasCustomfieldProjects returns a boolean if a field has been set.

### GetIncluded

`func (o *ProjectCustomFieldProjectsResponse) GetIncluded() ProjectCustomFieldProjectsResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *ProjectCustomFieldProjectsResponse) GetIncludedOk() (*ProjectCustomFieldProjectsResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *ProjectCustomFieldProjectsResponse) SetIncluded(v ProjectCustomFieldProjectsResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *ProjectCustomFieldProjectsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *ProjectCustomFieldProjectsResponse) GetMeta() ViewMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ProjectCustomFieldProjectsResponse) GetMetaOk() (*ViewMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ProjectCustomFieldProjectsResponse) SetMeta(v ViewMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ProjectCustomFieldProjectsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


