# ViewForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfirmationMessage** | Pointer to **string** |  | [optional] 
**Content** | Pointer to [**ViewContent**](ViewContent.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**DeletedBy** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**HostObject** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Installation** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**IsShared** | Pointer to **bool** |  | [optional] 
**PromptAdditionalSubmissions** | Pointer to **bool** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**TaskTitleFieldId** | Pointer to **string** |  | [optional] 
**Token** | Pointer to [**ViewToken**](ViewToken.md) |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 

## Methods

### NewViewForm

`func NewViewForm() *ViewForm`

NewViewForm instantiates a new ViewForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewFormWithDefaults

`func NewViewFormWithDefaults() *ViewForm`

NewViewFormWithDefaults instantiates a new ViewForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfirmationMessage

`func (o *ViewForm) GetConfirmationMessage() string`

GetConfirmationMessage returns the ConfirmationMessage field if non-nil, zero value otherwise.

### GetConfirmationMessageOk

`func (o *ViewForm) GetConfirmationMessageOk() (*string, bool)`

GetConfirmationMessageOk returns a tuple with the ConfirmationMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationMessage

`func (o *ViewForm) SetConfirmationMessage(v string)`

SetConfirmationMessage sets ConfirmationMessage field to given value.

### HasConfirmationMessage

`func (o *ViewForm) HasConfirmationMessage() bool`

HasConfirmationMessage returns a boolean if a field has been set.

### GetContent

`func (o *ViewForm) GetContent() ViewContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ViewForm) GetContentOk() (*ViewContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ViewForm) SetContent(v ViewContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *ViewForm) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ViewForm) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ViewForm) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ViewForm) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ViewForm) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ViewForm) GetCreatedBy() ViewRelationship`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ViewForm) GetCreatedByOk() (*ViewRelationship, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ViewForm) SetCreatedBy(v ViewRelationship)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ViewForm) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ViewForm) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ViewForm) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ViewForm) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ViewForm) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeletedBy

`func (o *ViewForm) GetDeletedBy() ViewRelationship`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *ViewForm) GetDeletedByOk() (*ViewRelationship, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *ViewForm) SetDeletedBy(v ViewRelationship)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *ViewForm) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetHostObject

`func (o *ViewForm) GetHostObject() ViewRelationship`

GetHostObject returns the HostObject field if non-nil, zero value otherwise.

### GetHostObjectOk

`func (o *ViewForm) GetHostObjectOk() (*ViewRelationship, bool)`

GetHostObjectOk returns a tuple with the HostObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostObject

`func (o *ViewForm) SetHostObject(v ViewRelationship)`

SetHostObject sets HostObject field to given value.

### HasHostObject

`func (o *ViewForm) HasHostObject() bool`

HasHostObject returns a boolean if a field has been set.

### GetId

`func (o *ViewForm) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewForm) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewForm) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewForm) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstallation

`func (o *ViewForm) GetInstallation() ViewRelationship`

GetInstallation returns the Installation field if non-nil, zero value otherwise.

### GetInstallationOk

`func (o *ViewForm) GetInstallationOk() (*ViewRelationship, bool)`

GetInstallationOk returns a tuple with the Installation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallation

`func (o *ViewForm) SetInstallation(v ViewRelationship)`

SetInstallation sets Installation field to given value.

### HasInstallation

`func (o *ViewForm) HasInstallation() bool`

HasInstallation returns a boolean if a field has been set.

### GetIsShared

`func (o *ViewForm) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *ViewForm) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *ViewForm) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *ViewForm) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### GetPromptAdditionalSubmissions

`func (o *ViewForm) GetPromptAdditionalSubmissions() bool`

GetPromptAdditionalSubmissions returns the PromptAdditionalSubmissions field if non-nil, zero value otherwise.

### GetPromptAdditionalSubmissionsOk

`func (o *ViewForm) GetPromptAdditionalSubmissionsOk() (*bool, bool)`

GetPromptAdditionalSubmissionsOk returns a tuple with the PromptAdditionalSubmissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptAdditionalSubmissions

`func (o *ViewForm) SetPromptAdditionalSubmissions(v bool)`

SetPromptAdditionalSubmissions sets PromptAdditionalSubmissions field to given value.

### HasPromptAdditionalSubmissions

`func (o *ViewForm) HasPromptAdditionalSubmissions() bool`

HasPromptAdditionalSubmissions returns a boolean if a field has been set.

### GetState

`func (o *ViewForm) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ViewForm) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ViewForm) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ViewForm) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTaskTitleFieldId

`func (o *ViewForm) GetTaskTitleFieldId() string`

GetTaskTitleFieldId returns the TaskTitleFieldId field if non-nil, zero value otherwise.

### GetTaskTitleFieldIdOk

`func (o *ViewForm) GetTaskTitleFieldIdOk() (*string, bool)`

GetTaskTitleFieldIdOk returns a tuple with the TaskTitleFieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskTitleFieldId

`func (o *ViewForm) SetTaskTitleFieldId(v string)`

SetTaskTitleFieldId sets TaskTitleFieldId field to given value.

### HasTaskTitleFieldId

`func (o *ViewForm) HasTaskTitleFieldId() bool`

HasTaskTitleFieldId returns a boolean if a field has been set.

### GetToken

`func (o *ViewForm) GetToken() ViewToken`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ViewForm) GetTokenOk() (*ViewToken, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ViewForm) SetToken(v ViewToken)`

SetToken sets Token field to given value.

### HasToken

`func (o *ViewForm) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ViewForm) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ViewForm) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ViewForm) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ViewForm) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ViewForm) GetUpdatedBy() ViewRelationship`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ViewForm) GetUpdatedByOk() (*ViewRelationship, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ViewForm) SetUpdatedBy(v ViewRelationship)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ViewForm) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


