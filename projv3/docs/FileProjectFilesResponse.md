# FileProjectFilesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to [**[]ViewProjectFile**](ViewProjectFile.md) |  | [optional] 
**Included** | Pointer to [**FileProjectFilesResponseIncluded**](FileProjectFilesResponseIncluded.md) |  | [optional] 
**Meta** | Pointer to [**ViewMeta**](ViewMeta.md) |  | [optional] 

## Methods

### NewFileProjectFilesResponse

`func NewFileProjectFilesResponse() *FileProjectFilesResponse`

NewFileProjectFilesResponse instantiates a new FileProjectFilesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileProjectFilesResponseWithDefaults

`func NewFileProjectFilesResponseWithDefaults() *FileProjectFilesResponse`

NewFileProjectFilesResponseWithDefaults instantiates a new FileProjectFilesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *FileProjectFilesResponse) GetFiles() []ViewProjectFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *FileProjectFilesResponse) GetFilesOk() (*[]ViewProjectFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *FileProjectFilesResponse) SetFiles(v []ViewProjectFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *FileProjectFilesResponse) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetIncluded

`func (o *FileProjectFilesResponse) GetIncluded() FileProjectFilesResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *FileProjectFilesResponse) GetIncludedOk() (*FileProjectFilesResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *FileProjectFilesResponse) SetIncluded(v FileProjectFilesResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *FileProjectFilesResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *FileProjectFilesResponse) GetMeta() ViewMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FileProjectFilesResponse) GetMetaOk() (*ViewMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FileProjectFilesResponse) SetMeta(v ViewMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FileProjectFilesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


