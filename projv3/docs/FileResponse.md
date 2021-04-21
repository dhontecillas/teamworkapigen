# FileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | Pointer to [**ViewProjectFile**](ViewProjectFile.md) |  | [optional] 
**Included** | Pointer to [**FileProjectFilesResponseIncluded**](FileProjectFilesResponseIncluded.md) |  | [optional] 

## Methods

### NewFileResponse

`func NewFileResponse() *FileResponse`

NewFileResponse instantiates a new FileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileResponseWithDefaults

`func NewFileResponseWithDefaults() *FileResponse`

NewFileResponseWithDefaults instantiates a new FileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *FileResponse) GetFile() ViewProjectFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *FileResponse) GetFileOk() (*ViewProjectFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *FileResponse) SetFile(v ViewProjectFile)`

SetFile sets File field to given value.

### HasFile

`func (o *FileResponse) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetIncluded

`func (o *FileResponse) GetIncluded() FileProjectFilesResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *FileResponse) GetIncludedOk() (*FileProjectFilesResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *FileResponse) SetIncluded(v FileProjectFilesResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *FileResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


