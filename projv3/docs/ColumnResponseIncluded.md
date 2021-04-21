# ColumnResponseIncluded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tasklists** | Pointer to [**map[string]ViewTasklist**](ViewTasklist.md) |  | [optional] 

## Methods

### NewColumnResponseIncluded

`func NewColumnResponseIncluded() *ColumnResponseIncluded`

NewColumnResponseIncluded instantiates a new ColumnResponseIncluded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColumnResponseIncludedWithDefaults

`func NewColumnResponseIncludedWithDefaults() *ColumnResponseIncluded`

NewColumnResponseIncludedWithDefaults instantiates a new ColumnResponseIncluded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTasklists

`func (o *ColumnResponseIncluded) GetTasklists() map[string]ViewTasklist`

GetTasklists returns the Tasklists field if non-nil, zero value otherwise.

### GetTasklistsOk

`func (o *ColumnResponseIncluded) GetTasklistsOk() (*map[string]ViewTasklist, bool)`

GetTasklistsOk returns a tuple with the Tasklists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasklists

`func (o *ColumnResponseIncluded) SetTasklists(v map[string]ViewTasklist)`

SetTasklists sets Tasklists field to given value.

### HasTasklists

`func (o *ColumnResponseIncluded) HasTasklists() bool`

HasTasklists returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


