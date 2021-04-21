# NotebookNotebooksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Included** | Pointer to [**NotebookNotebooksResponseIncluded**](NotebookNotebooksResponseIncluded.md) |  | [optional] 
**Meta** | Pointer to [**ViewMeta**](ViewMeta.md) |  | [optional] 
**Notebooks** | Pointer to [**[]ViewNotebook**](ViewNotebook.md) |  | [optional] 

## Methods

### NewNotebookNotebooksResponse

`func NewNotebookNotebooksResponse() *NotebookNotebooksResponse`

NewNotebookNotebooksResponse instantiates a new NotebookNotebooksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotebookNotebooksResponseWithDefaults

`func NewNotebookNotebooksResponseWithDefaults() *NotebookNotebooksResponse`

NewNotebookNotebooksResponseWithDefaults instantiates a new NotebookNotebooksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncluded

`func (o *NotebookNotebooksResponse) GetIncluded() NotebookNotebooksResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *NotebookNotebooksResponse) GetIncludedOk() (*NotebookNotebooksResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *NotebookNotebooksResponse) SetIncluded(v NotebookNotebooksResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *NotebookNotebooksResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *NotebookNotebooksResponse) GetMeta() ViewMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *NotebookNotebooksResponse) GetMetaOk() (*ViewMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *NotebookNotebooksResponse) SetMeta(v ViewMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *NotebookNotebooksResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetNotebooks

`func (o *NotebookNotebooksResponse) GetNotebooks() []ViewNotebook`

GetNotebooks returns the Notebooks field if non-nil, zero value otherwise.

### GetNotebooksOk

`func (o *NotebookNotebooksResponse) GetNotebooksOk() (*[]ViewNotebook, bool)`

GetNotebooksOk returns a tuple with the Notebooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotebooks

`func (o *NotebookNotebooksResponse) SetNotebooks(v []ViewNotebook)`

SetNotebooks sets Notebooks field to given value.

### HasNotebooks

`func (o *NotebookNotebooksResponse) HasNotebooks() bool`

HasNotebooks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


