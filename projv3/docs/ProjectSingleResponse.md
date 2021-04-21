# ProjectSingleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Included** | Pointer to [**ProjectProjectsResponseIncluded**](ProjectProjectsResponseIncluded.md) |  | [optional] 
**Project** | Pointer to [**ViewProject**](ViewProject.md) |  | [optional] 

## Methods

### NewProjectSingleResponse

`func NewProjectSingleResponse() *ProjectSingleResponse`

NewProjectSingleResponse instantiates a new ProjectSingleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectSingleResponseWithDefaults

`func NewProjectSingleResponseWithDefaults() *ProjectSingleResponse`

NewProjectSingleResponseWithDefaults instantiates a new ProjectSingleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncluded

`func (o *ProjectSingleResponse) GetIncluded() ProjectProjectsResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *ProjectSingleResponse) GetIncludedOk() (*ProjectProjectsResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *ProjectSingleResponse) SetIncluded(v ProjectProjectsResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *ProjectSingleResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetProject

`func (o *ProjectSingleResponse) GetProject() ViewProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ProjectSingleResponse) GetProjectOk() (*ViewProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ProjectSingleResponse) SetProject(v ViewProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *ProjectSingleResponse) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


