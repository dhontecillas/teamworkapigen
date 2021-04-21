# InlineResponse20043

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**STATUS** | Pointer to **string** |  | [optional] 
**Board** | Pointer to [**InlineResponse20043Board**](InlineResponse20043Board.md) |  | [optional] 
**Columns** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewInlineResponse20043

`func NewInlineResponse20043() *InlineResponse20043`

NewInlineResponse20043 instantiates a new InlineResponse20043 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20043WithDefaults

`func NewInlineResponse20043WithDefaults() *InlineResponse20043`

NewInlineResponse20043WithDefaults instantiates a new InlineResponse20043 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSTATUS

`func (o *InlineResponse20043) GetSTATUS() string`

GetSTATUS returns the STATUS field if non-nil, zero value otherwise.

### GetSTATUSOk

`func (o *InlineResponse20043) GetSTATUSOk() (*string, bool)`

GetSTATUSOk returns a tuple with the STATUS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS

`func (o *InlineResponse20043) SetSTATUS(v string)`

SetSTATUS sets STATUS field to given value.

### HasSTATUS

`func (o *InlineResponse20043) HasSTATUS() bool`

HasSTATUS returns a boolean if a field has been set.

### GetBoard

`func (o *InlineResponse20043) GetBoard() InlineResponse20043Board`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *InlineResponse20043) GetBoardOk() (*InlineResponse20043Board, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *InlineResponse20043) SetBoard(v InlineResponse20043Board)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *InlineResponse20043) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetColumns

`func (o *InlineResponse20043) GetColumns() []map[string]interface{}`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *InlineResponse20043) GetColumnsOk() (*[]map[string]interface{}, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *InlineResponse20043) SetColumns(v []map[string]interface{})`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *InlineResponse20043) HasColumns() bool`

HasColumns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


