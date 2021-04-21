# FileUsageResponseFilesUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overuse** | Pointer to **int32** |  | [optional] 
**OveruseAsText** | Pointer to **string** |  | [optional] 
**Projects** | Pointer to [**[]FileProjectUsage**](FileProjectUsage.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**TotalAllowed** | Pointer to **int32** |  | [optional] 
**TotalAllowedAsText** | Pointer to **string** |  | [optional] 
**TotalAsText** | Pointer to **string** |  | [optional] 
**TotalRemaining** | Pointer to **int32** |  | [optional] 
**TotalRemainingAsText** | Pointer to **string** |  | [optional] 

## Methods

### NewFileUsageResponseFilesUsage

`func NewFileUsageResponseFilesUsage() *FileUsageResponseFilesUsage`

NewFileUsageResponseFilesUsage instantiates a new FileUsageResponseFilesUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileUsageResponseFilesUsageWithDefaults

`func NewFileUsageResponseFilesUsageWithDefaults() *FileUsageResponseFilesUsage`

NewFileUsageResponseFilesUsageWithDefaults instantiates a new FileUsageResponseFilesUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOveruse

`func (o *FileUsageResponseFilesUsage) GetOveruse() int32`

GetOveruse returns the Overuse field if non-nil, zero value otherwise.

### GetOveruseOk

`func (o *FileUsageResponseFilesUsage) GetOveruseOk() (*int32, bool)`

GetOveruseOk returns a tuple with the Overuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOveruse

`func (o *FileUsageResponseFilesUsage) SetOveruse(v int32)`

SetOveruse sets Overuse field to given value.

### HasOveruse

`func (o *FileUsageResponseFilesUsage) HasOveruse() bool`

HasOveruse returns a boolean if a field has been set.

### GetOveruseAsText

`func (o *FileUsageResponseFilesUsage) GetOveruseAsText() string`

GetOveruseAsText returns the OveruseAsText field if non-nil, zero value otherwise.

### GetOveruseAsTextOk

`func (o *FileUsageResponseFilesUsage) GetOveruseAsTextOk() (*string, bool)`

GetOveruseAsTextOk returns a tuple with the OveruseAsText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOveruseAsText

`func (o *FileUsageResponseFilesUsage) SetOveruseAsText(v string)`

SetOveruseAsText sets OveruseAsText field to given value.

### HasOveruseAsText

`func (o *FileUsageResponseFilesUsage) HasOveruseAsText() bool`

HasOveruseAsText returns a boolean if a field has been set.

### GetProjects

`func (o *FileUsageResponseFilesUsage) GetProjects() []FileProjectUsage`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *FileUsageResponseFilesUsage) GetProjectsOk() (*[]FileProjectUsage, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *FileUsageResponseFilesUsage) SetProjects(v []FileProjectUsage)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *FileUsageResponseFilesUsage) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetTotal

`func (o *FileUsageResponseFilesUsage) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *FileUsageResponseFilesUsage) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *FileUsageResponseFilesUsage) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *FileUsageResponseFilesUsage) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTotalAllowed

`func (o *FileUsageResponseFilesUsage) GetTotalAllowed() int32`

GetTotalAllowed returns the TotalAllowed field if non-nil, zero value otherwise.

### GetTotalAllowedOk

`func (o *FileUsageResponseFilesUsage) GetTotalAllowedOk() (*int32, bool)`

GetTotalAllowedOk returns a tuple with the TotalAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAllowed

`func (o *FileUsageResponseFilesUsage) SetTotalAllowed(v int32)`

SetTotalAllowed sets TotalAllowed field to given value.

### HasTotalAllowed

`func (o *FileUsageResponseFilesUsage) HasTotalAllowed() bool`

HasTotalAllowed returns a boolean if a field has been set.

### GetTotalAllowedAsText

`func (o *FileUsageResponseFilesUsage) GetTotalAllowedAsText() string`

GetTotalAllowedAsText returns the TotalAllowedAsText field if non-nil, zero value otherwise.

### GetTotalAllowedAsTextOk

`func (o *FileUsageResponseFilesUsage) GetTotalAllowedAsTextOk() (*string, bool)`

GetTotalAllowedAsTextOk returns a tuple with the TotalAllowedAsText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAllowedAsText

`func (o *FileUsageResponseFilesUsage) SetTotalAllowedAsText(v string)`

SetTotalAllowedAsText sets TotalAllowedAsText field to given value.

### HasTotalAllowedAsText

`func (o *FileUsageResponseFilesUsage) HasTotalAllowedAsText() bool`

HasTotalAllowedAsText returns a boolean if a field has been set.

### GetTotalAsText

`func (o *FileUsageResponseFilesUsage) GetTotalAsText() string`

GetTotalAsText returns the TotalAsText field if non-nil, zero value otherwise.

### GetTotalAsTextOk

`func (o *FileUsageResponseFilesUsage) GetTotalAsTextOk() (*string, bool)`

GetTotalAsTextOk returns a tuple with the TotalAsText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAsText

`func (o *FileUsageResponseFilesUsage) SetTotalAsText(v string)`

SetTotalAsText sets TotalAsText field to given value.

### HasTotalAsText

`func (o *FileUsageResponseFilesUsage) HasTotalAsText() bool`

HasTotalAsText returns a boolean if a field has been set.

### GetTotalRemaining

`func (o *FileUsageResponseFilesUsage) GetTotalRemaining() int32`

GetTotalRemaining returns the TotalRemaining field if non-nil, zero value otherwise.

### GetTotalRemainingOk

`func (o *FileUsageResponseFilesUsage) GetTotalRemainingOk() (*int32, bool)`

GetTotalRemainingOk returns a tuple with the TotalRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRemaining

`func (o *FileUsageResponseFilesUsage) SetTotalRemaining(v int32)`

SetTotalRemaining sets TotalRemaining field to given value.

### HasTotalRemaining

`func (o *FileUsageResponseFilesUsage) HasTotalRemaining() bool`

HasTotalRemaining returns a boolean if a field has been set.

### GetTotalRemainingAsText

`func (o *FileUsageResponseFilesUsage) GetTotalRemainingAsText() string`

GetTotalRemainingAsText returns the TotalRemainingAsText field if non-nil, zero value otherwise.

### GetTotalRemainingAsTextOk

`func (o *FileUsageResponseFilesUsage) GetTotalRemainingAsTextOk() (*string, bool)`

GetTotalRemainingAsTextOk returns a tuple with the TotalRemainingAsText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRemainingAsText

`func (o *FileUsageResponseFilesUsage) SetTotalRemainingAsText(v string)`

SetTotalRemainingAsText sets TotalRemainingAsText field to given value.

### HasTotalRemainingAsText

`func (o *FileUsageResponseFilesUsage) HasTotalRemainingAsText() bool`

HasTotalRemainingAsText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


