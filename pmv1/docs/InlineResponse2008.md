# InlineResponse2008

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**STATUS** | Pointer to **string** |  | [optional] 
**Reactions** | Pointer to [**InlineResponse2008Reactions**](InlineResponse2008Reactions.md) |  | [optional] 

## Methods

### NewInlineResponse2008

`func NewInlineResponse2008() *InlineResponse2008`

NewInlineResponse2008 instantiates a new InlineResponse2008 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2008WithDefaults

`func NewInlineResponse2008WithDefaults() *InlineResponse2008`

NewInlineResponse2008WithDefaults instantiates a new InlineResponse2008 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSTATUS

`func (o *InlineResponse2008) GetSTATUS() string`

GetSTATUS returns the STATUS field if non-nil, zero value otherwise.

### GetSTATUSOk

`func (o *InlineResponse2008) GetSTATUSOk() (*string, bool)`

GetSTATUSOk returns a tuple with the STATUS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS

`func (o *InlineResponse2008) SetSTATUS(v string)`

SetSTATUS sets STATUS field to given value.

### HasSTATUS

`func (o *InlineResponse2008) HasSTATUS() bool`

HasSTATUS returns a boolean if a field has been set.

### GetReactions

`func (o *InlineResponse2008) GetReactions() InlineResponse2008Reactions`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *InlineResponse2008) GetReactionsOk() (*InlineResponse2008Reactions, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *InlineResponse2008) SetReactions(v InlineResponse2008Reactions)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *InlineResponse2008) HasReactions() bool`

HasReactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


