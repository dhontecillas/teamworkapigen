# NotebookVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Included** | Pointer to [**NotebookVersionResponseIncluded**](NotebookVersionResponseIncluded.md) |  | [optional] 
**Version** | Pointer to [**ViewNotebookVersion**](ViewNotebookVersion.md) |  | [optional] 

## Methods

### NewNotebookVersionResponse

`func NewNotebookVersionResponse() *NotebookVersionResponse`

NewNotebookVersionResponse instantiates a new NotebookVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotebookVersionResponseWithDefaults

`func NewNotebookVersionResponseWithDefaults() *NotebookVersionResponse`

NewNotebookVersionResponseWithDefaults instantiates a new NotebookVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncluded

`func (o *NotebookVersionResponse) GetIncluded() NotebookVersionResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *NotebookVersionResponse) GetIncludedOk() (*NotebookVersionResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *NotebookVersionResponse) SetIncluded(v NotebookVersionResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *NotebookVersionResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetVersion

`func (o *NotebookVersionResponse) GetVersion() ViewNotebookVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NotebookVersionResponse) GetVersionOk() (*ViewNotebookVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NotebookVersionResponse) SetVersion(v ViewNotebookVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NotebookVersionResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


