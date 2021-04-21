# NotebookVersionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Included** | Pointer to [**NotebookVersionResponseIncluded**](NotebookVersionResponseIncluded.md) |  | [optional] 
**Versions** | Pointer to [**[]ViewNotebookVersion**](ViewNotebookVersion.md) |  | [optional] 

## Methods

### NewNotebookVersionsResponse

`func NewNotebookVersionsResponse() *NotebookVersionsResponse`

NewNotebookVersionsResponse instantiates a new NotebookVersionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotebookVersionsResponseWithDefaults

`func NewNotebookVersionsResponseWithDefaults() *NotebookVersionsResponse`

NewNotebookVersionsResponseWithDefaults instantiates a new NotebookVersionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncluded

`func (o *NotebookVersionsResponse) GetIncluded() NotebookVersionResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *NotebookVersionsResponse) GetIncludedOk() (*NotebookVersionResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *NotebookVersionsResponse) SetIncluded(v NotebookVersionResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *NotebookVersionsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetVersions

`func (o *NotebookVersionsResponse) GetVersions() []ViewNotebookVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *NotebookVersionsResponse) GetVersionsOk() (*[]ViewNotebookVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *NotebookVersionsResponse) SetVersions(v []ViewNotebookVersion)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *NotebookVersionsResponse) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


