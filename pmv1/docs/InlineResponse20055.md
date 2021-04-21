# InlineResponse20055

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**STATUS** | Pointer to **string** |  | [optional] 
**NextPageToken** | Pointer to **int32** | Optionally returned where more data is available. | [optional] 
**Tasks** | Pointer to **[]string** |  | [optional] 

## Methods

### NewInlineResponse20055

`func NewInlineResponse20055() *InlineResponse20055`

NewInlineResponse20055 instantiates a new InlineResponse20055 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20055WithDefaults

`func NewInlineResponse20055WithDefaults() *InlineResponse20055`

NewInlineResponse20055WithDefaults instantiates a new InlineResponse20055 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSTATUS

`func (o *InlineResponse20055) GetSTATUS() string`

GetSTATUS returns the STATUS field if non-nil, zero value otherwise.

### GetSTATUSOk

`func (o *InlineResponse20055) GetSTATUSOk() (*string, bool)`

GetSTATUSOk returns a tuple with the STATUS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS

`func (o *InlineResponse20055) SetSTATUS(v string)`

SetSTATUS sets STATUS field to given value.

### HasSTATUS

`func (o *InlineResponse20055) HasSTATUS() bool`

HasSTATUS returns a boolean if a field has been set.

### GetNextPageToken

`func (o *InlineResponse20055) GetNextPageToken() int32`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *InlineResponse20055) GetNextPageTokenOk() (*int32, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *InlineResponse20055) SetNextPageToken(v int32)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *InlineResponse20055) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetTasks

`func (o *InlineResponse20055) GetTasks() []string`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *InlineResponse20055) GetTasksOk() (*[]string, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *InlineResponse20055) SetTasks(v []string)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *InlineResponse20055) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


