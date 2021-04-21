/*
 * Teamwork Projects API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pmv1

import (
	"encoding/json"
)

// InlineResponse20044 struct for InlineResponse20044
type InlineResponse20044 struct {
	STATUS *string `json:"STATUS,omitempty"`
	Board *InlineResponse20043Board `json:"board,omitempty"`
	Columns *[]InlineResponse20044Columns `json:"columns,omitempty"`
}

// NewInlineResponse20044 instantiates a new InlineResponse20044 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20044() *InlineResponse20044 {
	this := InlineResponse20044{}
	return &this
}

// NewInlineResponse20044WithDefaults instantiates a new InlineResponse20044 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20044WithDefaults() *InlineResponse20044 {
	this := InlineResponse20044{}
	return &this
}

// GetSTATUS returns the STATUS field value if set, zero value otherwise.
func (o *InlineResponse20044) GetSTATUS() string {
	if o == nil || o.STATUS == nil {
		var ret string
		return ret
	}
	return *o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044) GetSTATUSOk() (*string, bool) {
	if o == nil || o.STATUS == nil {
		return nil, false
	}
	return o.STATUS, true
}

// HasSTATUS returns a boolean if a field has been set.
func (o *InlineResponse20044) HasSTATUS() bool {
	if o != nil && o.STATUS != nil {
		return true
	}

	return false
}

// SetSTATUS gets a reference to the given string and assigns it to the STATUS field.
func (o *InlineResponse20044) SetSTATUS(v string) {
	o.STATUS = &v
}

// GetBoard returns the Board field value if set, zero value otherwise.
func (o *InlineResponse20044) GetBoard() InlineResponse20043Board {
	if o == nil || o.Board == nil {
		var ret InlineResponse20043Board
		return ret
	}
	return *o.Board
}

// GetBoardOk returns a tuple with the Board field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044) GetBoardOk() (*InlineResponse20043Board, bool) {
	if o == nil || o.Board == nil {
		return nil, false
	}
	return o.Board, true
}

// HasBoard returns a boolean if a field has been set.
func (o *InlineResponse20044) HasBoard() bool {
	if o != nil && o.Board != nil {
		return true
	}

	return false
}

// SetBoard gets a reference to the given InlineResponse20043Board and assigns it to the Board field.
func (o *InlineResponse20044) SetBoard(v InlineResponse20043Board) {
	o.Board = &v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *InlineResponse20044) GetColumns() []InlineResponse20044Columns {
	if o == nil || o.Columns == nil {
		var ret []InlineResponse20044Columns
		return ret
	}
	return *o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044) GetColumnsOk() (*[]InlineResponse20044Columns, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *InlineResponse20044) HasColumns() bool {
	if o != nil && o.Columns != nil {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []InlineResponse20044Columns and assigns it to the Columns field.
func (o *InlineResponse20044) SetColumns(v []InlineResponse20044Columns) {
	o.Columns = &v
}

func (o InlineResponse20044) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.STATUS != nil {
		toSerialize["STATUS"] = o.STATUS
	}
	if o.Board != nil {
		toSerialize["board"] = o.Board
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20044 struct {
	value *InlineResponse20044
	isSet bool
}

func (v NullableInlineResponse20044) Get() *InlineResponse20044 {
	return v.value
}

func (v *NullableInlineResponse20044) Set(val *InlineResponse20044) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20044) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20044) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20044(val *InlineResponse20044) *NullableInlineResponse20044 {
	return &NullableInlineResponse20044{value: val, isSet: true}
}

func (v NullableInlineResponse20044) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20044) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


