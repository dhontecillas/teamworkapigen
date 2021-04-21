# InlineResponse200112

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**STATUS** | Pointer to **string** |  | [optional] 
**TimeEntry** | Pointer to [**InlineResponse200112TimeEntry**](InlineResponse200112TimeEntry.md) |  | [optional] 

## Methods

### NewInlineResponse200112

`func NewInlineResponse200112() *InlineResponse200112`

NewInlineResponse200112 instantiates a new InlineResponse200112 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200112WithDefaults

`func NewInlineResponse200112WithDefaults() *InlineResponse200112`

NewInlineResponse200112WithDefaults instantiates a new InlineResponse200112 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSTATUS

`func (o *InlineResponse200112) GetSTATUS() string`

GetSTATUS returns the STATUS field if non-nil, zero value otherwise.

### GetSTATUSOk

`func (o *InlineResponse200112) GetSTATUSOk() (*string, bool)`

GetSTATUSOk returns a tuple with the STATUS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS

`func (o *InlineResponse200112) SetSTATUS(v string)`

SetSTATUS sets STATUS field to given value.

### HasSTATUS

`func (o *InlineResponse200112) HasSTATUS() bool`

HasSTATUS returns a boolean if a field has been set.

### GetTimeEntry

`func (o *InlineResponse200112) GetTimeEntry() InlineResponse200112TimeEntry`

GetTimeEntry returns the TimeEntry field if non-nil, zero value otherwise.

### GetTimeEntryOk

`func (o *InlineResponse200112) GetTimeEntryOk() (*InlineResponse200112TimeEntry, bool)`

GetTimeEntryOk returns a tuple with the TimeEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEntry

`func (o *InlineResponse200112) SetTimeEntry(v InlineResponse200112TimeEntry)`

SetTimeEntry sets TimeEntry field to given value.

### HasTimeEntry

`func (o *InlineResponse200112) HasTimeEntry() bool`

HasTimeEntry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


