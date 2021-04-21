# FileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | Pointer to [**FileProjectFile**](FileProjectFile.md) |  | [optional] 

## Methods

### NewFileRequest

`func NewFileRequest() *FileRequest`

NewFileRequest instantiates a new FileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileRequestWithDefaults

`func NewFileRequestWithDefaults() *FileRequest`

NewFileRequestWithDefaults instantiates a new FileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *FileRequest) GetFile() FileProjectFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *FileRequest) GetFileOk() (*FileProjectFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *FileRequest) SetFile(v FileProjectFile)`

SetFile sets File field to given value.

### HasFile

`func (o *FileRequest) HasFile() bool`

HasFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


