# InlineResponse20046

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**STATUS** | Pointer to **string** |  | [optional] 
**Cards** | Pointer to [**[]InlineResponse20046Cards**](InlineResponse20046Cards.md) |  | [optional] 
**Column** | Pointer to [**InlineResponse2002Column**](InlineResponse2002Column.md) |  | [optional] 

## Methods

### NewInlineResponse20046

`func NewInlineResponse20046() *InlineResponse20046`

NewInlineResponse20046 instantiates a new InlineResponse20046 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20046WithDefaults

`func NewInlineResponse20046WithDefaults() *InlineResponse20046`

NewInlineResponse20046WithDefaults instantiates a new InlineResponse20046 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSTATUS

`func (o *InlineResponse20046) GetSTATUS() string`

GetSTATUS returns the STATUS field if non-nil, zero value otherwise.

### GetSTATUSOk

`func (o *InlineResponse20046) GetSTATUSOk() (*string, bool)`

GetSTATUSOk returns a tuple with the STATUS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS

`func (o *InlineResponse20046) SetSTATUS(v string)`

SetSTATUS sets STATUS field to given value.

### HasSTATUS

`func (o *InlineResponse20046) HasSTATUS() bool`

HasSTATUS returns a boolean if a field has been set.

### GetCards

`func (o *InlineResponse20046) GetCards() []InlineResponse20046Cards`

GetCards returns the Cards field if non-nil, zero value otherwise.

### GetCardsOk

`func (o *InlineResponse20046) GetCardsOk() (*[]InlineResponse20046Cards, bool)`

GetCardsOk returns a tuple with the Cards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCards

`func (o *InlineResponse20046) SetCards(v []InlineResponse20046Cards)`

SetCards sets Cards field to given value.

### HasCards

`func (o *InlineResponse20046) HasCards() bool`

HasCards returns a boolean if a field has been set.

### GetColumn

`func (o *InlineResponse20046) GetColumn() InlineResponse2002Column`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *InlineResponse20046) GetColumnOk() (*InlineResponse2002Column, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *InlineResponse20046) SetColumn(v InlineResponse2002Column)`

SetColumn sets Column field to given value.

### HasColumn

`func (o *InlineResponse20046) HasColumn() bool`

HasColumn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


