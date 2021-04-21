# InlineResponse2005

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**STATUS** | Pointer to **string** |  | [optional] 
**Event** | Pointer to [**InlineResponse2005Event**](InlineResponse2005Event.md) |  | [optional] 

## Methods

### NewInlineResponse2005

`func NewInlineResponse2005() *InlineResponse2005`

NewInlineResponse2005 instantiates a new InlineResponse2005 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2005WithDefaults

`func NewInlineResponse2005WithDefaults() *InlineResponse2005`

NewInlineResponse2005WithDefaults instantiates a new InlineResponse2005 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSTATUS

`func (o *InlineResponse2005) GetSTATUS() string`

GetSTATUS returns the STATUS field if non-nil, zero value otherwise.

### GetSTATUSOk

`func (o *InlineResponse2005) GetSTATUSOk() (*string, bool)`

GetSTATUSOk returns a tuple with the STATUS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS

`func (o *InlineResponse2005) SetSTATUS(v string)`

SetSTATUS sets STATUS field to given value.

### HasSTATUS

`func (o *InlineResponse2005) HasSTATUS() bool`

HasSTATUS returns a boolean if a field has been set.

### GetEvent

`func (o *InlineResponse2005) GetEvent() InlineResponse2005Event`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *InlineResponse2005) GetEventOk() (*InlineResponse2005Event, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *InlineResponse2005) SetEvent(v InlineResponse2005Event)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *InlineResponse2005) HasEvent() bool`

HasEvent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


