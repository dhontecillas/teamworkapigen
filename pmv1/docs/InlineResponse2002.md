# InlineResponse2002

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**STATUS** | Pointer to **string** |  | [optional] 
**Cards** | Pointer to [**[]InlineResponse2002Cards**](InlineResponse2002Cards.md) |  | [optional] 
**Column** | Pointer to [**InlineResponse2002Column**](InlineResponse2002Column.md) |  | [optional] 
**People** | Pointer to [**InlineResponse2002People**](InlineResponse2002People.md) |  | [optional] 

## Methods

### NewInlineResponse2002

`func NewInlineResponse2002() *InlineResponse2002`

NewInlineResponse2002 instantiates a new InlineResponse2002 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2002WithDefaults

`func NewInlineResponse2002WithDefaults() *InlineResponse2002`

NewInlineResponse2002WithDefaults instantiates a new InlineResponse2002 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSTATUS

`func (o *InlineResponse2002) GetSTATUS() string`

GetSTATUS returns the STATUS field if non-nil, zero value otherwise.

### GetSTATUSOk

`func (o *InlineResponse2002) GetSTATUSOk() (*string, bool)`

GetSTATUSOk returns a tuple with the STATUS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS

`func (o *InlineResponse2002) SetSTATUS(v string)`

SetSTATUS sets STATUS field to given value.

### HasSTATUS

`func (o *InlineResponse2002) HasSTATUS() bool`

HasSTATUS returns a boolean if a field has been set.

### GetCards

`func (o *InlineResponse2002) GetCards() []InlineResponse2002Cards`

GetCards returns the Cards field if non-nil, zero value otherwise.

### GetCardsOk

`func (o *InlineResponse2002) GetCardsOk() (*[]InlineResponse2002Cards, bool)`

GetCardsOk returns a tuple with the Cards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCards

`func (o *InlineResponse2002) SetCards(v []InlineResponse2002Cards)`

SetCards sets Cards field to given value.

### HasCards

`func (o *InlineResponse2002) HasCards() bool`

HasCards returns a boolean if a field has been set.

### GetColumn

`func (o *InlineResponse2002) GetColumn() InlineResponse2002Column`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *InlineResponse2002) GetColumnOk() (*InlineResponse2002Column, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *InlineResponse2002) SetColumn(v InlineResponse2002Column)`

SetColumn sets Column field to given value.

### HasColumn

`func (o *InlineResponse2002) HasColumn() bool`

HasColumn returns a boolean if a field has been set.

### GetPeople

`func (o *InlineResponse2002) GetPeople() InlineResponse2002People`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *InlineResponse2002) GetPeopleOk() (*InlineResponse2002People, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *InlineResponse2002) SetPeople(v InlineResponse2002People)`

SetPeople sets People field to given value.

### HasPeople

`func (o *InlineResponse2002) HasPeople() bool`

HasPeople returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


