# ProjectsTemplateJsonProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | Pointer to **int32** |  | [optional] 
**CompanyId** | Pointer to **int32** | Not set results to Owner Company by default | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**People** | Pointer to **string** |  | [optional] 
**ProjectOwnerId** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **string** |  | [optional] 
**TemplateDateTargetDefault** | **string** | \&quot;start\&quot; OR \&quot;end\&quot; | 
**UseBilling** | Pointer to **bool** |  | [optional] 
**UseComments** | Pointer to **bool** |  | [optional] 
**UseFiles** | Pointer to **bool** |  | [optional] 
**UseLinks** | Pointer to **bool** |  | [optional] 
**UseMessages** | Pointer to **bool** |  | [optional] 
**UseMilestones** | Pointer to **bool** |  | [optional] 
**UseNotebook** | Pointer to **bool** |  | [optional] 
**UseRiskregister** | Pointer to **bool** |  | [optional] 
**UseTasks** | Pointer to **bool** |  | [optional] 
**UseTime** | Pointer to **bool** |  | [optional] 

## Methods

### NewProjectsTemplateJsonProject

`func NewProjectsTemplateJsonProject(templateDateTargetDefault string, ) *ProjectsTemplateJsonProject`

NewProjectsTemplateJsonProject instantiates a new ProjectsTemplateJsonProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsTemplateJsonProjectWithDefaults

`func NewProjectsTemplateJsonProjectWithDefaults() *ProjectsTemplateJsonProject`

NewProjectsTemplateJsonProjectWithDefaults instantiates a new ProjectsTemplateJsonProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *ProjectsTemplateJsonProject) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *ProjectsTemplateJsonProject) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *ProjectsTemplateJsonProject) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *ProjectsTemplateJsonProject) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetCompanyId

`func (o *ProjectsTemplateJsonProject) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ProjectsTemplateJsonProject) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ProjectsTemplateJsonProject) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *ProjectsTemplateJsonProject) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetDescription

`func (o *ProjectsTemplateJsonProject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectsTemplateJsonProject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectsTemplateJsonProject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectsTemplateJsonProject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ProjectsTemplateJsonProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectsTemplateJsonProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectsTemplateJsonProject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectsTemplateJsonProject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPeople

`func (o *ProjectsTemplateJsonProject) GetPeople() string`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *ProjectsTemplateJsonProject) GetPeopleOk() (*string, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *ProjectsTemplateJsonProject) SetPeople(v string)`

SetPeople sets People field to given value.

### HasPeople

`func (o *ProjectsTemplateJsonProject) HasPeople() bool`

HasPeople returns a boolean if a field has been set.

### GetProjectOwnerId

`func (o *ProjectsTemplateJsonProject) GetProjectOwnerId() string`

GetProjectOwnerId returns the ProjectOwnerId field if non-nil, zero value otherwise.

### GetProjectOwnerIdOk

`func (o *ProjectsTemplateJsonProject) GetProjectOwnerIdOk() (*string, bool)`

GetProjectOwnerIdOk returns a tuple with the ProjectOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectOwnerId

`func (o *ProjectsTemplateJsonProject) SetProjectOwnerId(v string)`

SetProjectOwnerId sets ProjectOwnerId field to given value.

### HasProjectOwnerId

`func (o *ProjectsTemplateJsonProject) HasProjectOwnerId() bool`

HasProjectOwnerId returns a boolean if a field has been set.

### GetTags

`func (o *ProjectsTemplateJsonProject) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProjectsTemplateJsonProject) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProjectsTemplateJsonProject) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProjectsTemplateJsonProject) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTemplateDateTargetDefault

`func (o *ProjectsTemplateJsonProject) GetTemplateDateTargetDefault() string`

GetTemplateDateTargetDefault returns the TemplateDateTargetDefault field if non-nil, zero value otherwise.

### GetTemplateDateTargetDefaultOk

`func (o *ProjectsTemplateJsonProject) GetTemplateDateTargetDefaultOk() (*string, bool)`

GetTemplateDateTargetDefaultOk returns a tuple with the TemplateDateTargetDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateDateTargetDefault

`func (o *ProjectsTemplateJsonProject) SetTemplateDateTargetDefault(v string)`

SetTemplateDateTargetDefault sets TemplateDateTargetDefault field to given value.


### GetUseBilling

`func (o *ProjectsTemplateJsonProject) GetUseBilling() bool`

GetUseBilling returns the UseBilling field if non-nil, zero value otherwise.

### GetUseBillingOk

`func (o *ProjectsTemplateJsonProject) GetUseBillingOk() (*bool, bool)`

GetUseBillingOk returns a tuple with the UseBilling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBilling

`func (o *ProjectsTemplateJsonProject) SetUseBilling(v bool)`

SetUseBilling sets UseBilling field to given value.

### HasUseBilling

`func (o *ProjectsTemplateJsonProject) HasUseBilling() bool`

HasUseBilling returns a boolean if a field has been set.

### GetUseComments

`func (o *ProjectsTemplateJsonProject) GetUseComments() bool`

GetUseComments returns the UseComments field if non-nil, zero value otherwise.

### GetUseCommentsOk

`func (o *ProjectsTemplateJsonProject) GetUseCommentsOk() (*bool, bool)`

GetUseCommentsOk returns a tuple with the UseComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseComments

`func (o *ProjectsTemplateJsonProject) SetUseComments(v bool)`

SetUseComments sets UseComments field to given value.

### HasUseComments

`func (o *ProjectsTemplateJsonProject) HasUseComments() bool`

HasUseComments returns a boolean if a field has been set.

### GetUseFiles

`func (o *ProjectsTemplateJsonProject) GetUseFiles() bool`

GetUseFiles returns the UseFiles field if non-nil, zero value otherwise.

### GetUseFilesOk

`func (o *ProjectsTemplateJsonProject) GetUseFilesOk() (*bool, bool)`

GetUseFilesOk returns a tuple with the UseFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFiles

`func (o *ProjectsTemplateJsonProject) SetUseFiles(v bool)`

SetUseFiles sets UseFiles field to given value.

### HasUseFiles

`func (o *ProjectsTemplateJsonProject) HasUseFiles() bool`

HasUseFiles returns a boolean if a field has been set.

### GetUseLinks

`func (o *ProjectsTemplateJsonProject) GetUseLinks() bool`

GetUseLinks returns the UseLinks field if non-nil, zero value otherwise.

### GetUseLinksOk

`func (o *ProjectsTemplateJsonProject) GetUseLinksOk() (*bool, bool)`

GetUseLinksOk returns a tuple with the UseLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLinks

`func (o *ProjectsTemplateJsonProject) SetUseLinks(v bool)`

SetUseLinks sets UseLinks field to given value.

### HasUseLinks

`func (o *ProjectsTemplateJsonProject) HasUseLinks() bool`

HasUseLinks returns a boolean if a field has been set.

### GetUseMessages

`func (o *ProjectsTemplateJsonProject) GetUseMessages() bool`

GetUseMessages returns the UseMessages field if non-nil, zero value otherwise.

### GetUseMessagesOk

`func (o *ProjectsTemplateJsonProject) GetUseMessagesOk() (*bool, bool)`

GetUseMessagesOk returns a tuple with the UseMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMessages

`func (o *ProjectsTemplateJsonProject) SetUseMessages(v bool)`

SetUseMessages sets UseMessages field to given value.

### HasUseMessages

`func (o *ProjectsTemplateJsonProject) HasUseMessages() bool`

HasUseMessages returns a boolean if a field has been set.

### GetUseMilestones

`func (o *ProjectsTemplateJsonProject) GetUseMilestones() bool`

GetUseMilestones returns the UseMilestones field if non-nil, zero value otherwise.

### GetUseMilestonesOk

`func (o *ProjectsTemplateJsonProject) GetUseMilestonesOk() (*bool, bool)`

GetUseMilestonesOk returns a tuple with the UseMilestones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMilestones

`func (o *ProjectsTemplateJsonProject) SetUseMilestones(v bool)`

SetUseMilestones sets UseMilestones field to given value.

### HasUseMilestones

`func (o *ProjectsTemplateJsonProject) HasUseMilestones() bool`

HasUseMilestones returns a boolean if a field has been set.

### GetUseNotebook

`func (o *ProjectsTemplateJsonProject) GetUseNotebook() bool`

GetUseNotebook returns the UseNotebook field if non-nil, zero value otherwise.

### GetUseNotebookOk

`func (o *ProjectsTemplateJsonProject) GetUseNotebookOk() (*bool, bool)`

GetUseNotebookOk returns a tuple with the UseNotebook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNotebook

`func (o *ProjectsTemplateJsonProject) SetUseNotebook(v bool)`

SetUseNotebook sets UseNotebook field to given value.

### HasUseNotebook

`func (o *ProjectsTemplateJsonProject) HasUseNotebook() bool`

HasUseNotebook returns a boolean if a field has been set.

### GetUseRiskregister

`func (o *ProjectsTemplateJsonProject) GetUseRiskregister() bool`

GetUseRiskregister returns the UseRiskregister field if non-nil, zero value otherwise.

### GetUseRiskregisterOk

`func (o *ProjectsTemplateJsonProject) GetUseRiskregisterOk() (*bool, bool)`

GetUseRiskregisterOk returns a tuple with the UseRiskregister field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRiskregister

`func (o *ProjectsTemplateJsonProject) SetUseRiskregister(v bool)`

SetUseRiskregister sets UseRiskregister field to given value.

### HasUseRiskregister

`func (o *ProjectsTemplateJsonProject) HasUseRiskregister() bool`

HasUseRiskregister returns a boolean if a field has been set.

### GetUseTasks

`func (o *ProjectsTemplateJsonProject) GetUseTasks() bool`

GetUseTasks returns the UseTasks field if non-nil, zero value otherwise.

### GetUseTasksOk

`func (o *ProjectsTemplateJsonProject) GetUseTasksOk() (*bool, bool)`

GetUseTasksOk returns a tuple with the UseTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTasks

`func (o *ProjectsTemplateJsonProject) SetUseTasks(v bool)`

SetUseTasks sets UseTasks field to given value.

### HasUseTasks

`func (o *ProjectsTemplateJsonProject) HasUseTasks() bool`

HasUseTasks returns a boolean if a field has been set.

### GetUseTime

`func (o *ProjectsTemplateJsonProject) GetUseTime() bool`

GetUseTime returns the UseTime field if non-nil, zero value otherwise.

### GetUseTimeOk

`func (o *ProjectsTemplateJsonProject) GetUseTimeOk() (*bool, bool)`

GetUseTimeOk returns a tuple with the UseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTime

`func (o *ProjectsTemplateJsonProject) SetUseTime(v bool)`

SetUseTime sets UseTime field to given value.

### HasUseTime

`func (o *ProjectsTemplateJsonProject) HasUseTime() bool`

HasUseTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


