# NotebookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Included** | Pointer to [**NotebookNotebooksResponseIncluded**](NotebookNotebooksResponseIncluded.md) |  | [optional] 
**Notebook** | Pointer to [**ViewNotebook**](ViewNotebook.md) |  | [optional] 

## Methods

### NewNotebookResponse

`func NewNotebookResponse() *NotebookResponse`

NewNotebookResponse instantiates a new NotebookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotebookResponseWithDefaults

`func NewNotebookResponseWithDefaults() *NotebookResponse`

NewNotebookResponseWithDefaults instantiates a new NotebookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncluded

`func (o *NotebookResponse) GetIncluded() NotebookNotebooksResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *NotebookResponse) GetIncludedOk() (*NotebookNotebooksResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *NotebookResponse) SetIncluded(v NotebookNotebooksResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *NotebookResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetNotebook

`func (o *NotebookResponse) GetNotebook() ViewNotebook`

GetNotebook returns the Notebook field if non-nil, zero value otherwise.

### GetNotebookOk

`func (o *NotebookResponse) GetNotebookOk() (*ViewNotebook, bool)`

GetNotebookOk returns a tuple with the Notebook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotebook

`func (o *NotebookResponse) SetNotebook(v ViewNotebook)`

SetNotebook sets Notebook field to given value.

### HasNotebook

`func (o *NotebookResponse) HasNotebook() bool`

HasNotebook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


