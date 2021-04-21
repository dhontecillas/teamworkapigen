# FileUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilesUsage** | Pointer to [**FileUsageResponseFilesUsage**](FileUsageResponseFilesUsage.md) |  | [optional] 
**Included** | Pointer to [**FileUsageResponseIncluded**](FileUsageResponseIncluded.md) |  | [optional] 

## Methods

### NewFileUsageResponse

`func NewFileUsageResponse() *FileUsageResponse`

NewFileUsageResponse instantiates a new FileUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileUsageResponseWithDefaults

`func NewFileUsageResponseWithDefaults() *FileUsageResponse`

NewFileUsageResponseWithDefaults instantiates a new FileUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilesUsage

`func (o *FileUsageResponse) GetFilesUsage() FileUsageResponseFilesUsage`

GetFilesUsage returns the FilesUsage field if non-nil, zero value otherwise.

### GetFilesUsageOk

`func (o *FileUsageResponse) GetFilesUsageOk() (*FileUsageResponseFilesUsage, bool)`

GetFilesUsageOk returns a tuple with the FilesUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesUsage

`func (o *FileUsageResponse) SetFilesUsage(v FileUsageResponseFilesUsage)`

SetFilesUsage sets FilesUsage field to given value.

### HasFilesUsage

`func (o *FileUsageResponse) HasFilesUsage() bool`

HasFilesUsage returns a boolean if a field has been set.

### GetIncluded

`func (o *FileUsageResponse) GetIncluded() FileUsageResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *FileUsageResponse) GetIncludedOk() (*FileUsageResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *FileUsageResponse) SetIncluded(v FileUsageResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *FileUsageResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


