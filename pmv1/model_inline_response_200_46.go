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

// InlineResponse20046 struct for InlineResponse20046
type InlineResponse20046 struct {
	STATUS *string `json:"STATUS,omitempty"`
	Cards *[]InlineResponse20046Cards `json:"cards,omitempty"`
	Column *InlineResponse2002Column `json:"column,omitempty"`
}

// NewInlineResponse20046 instantiates a new InlineResponse20046 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20046() *InlineResponse20046 {
	this := InlineResponse20046{}
	return &this
}

// NewInlineResponse20046WithDefaults instantiates a new InlineResponse20046 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20046WithDefaults() *InlineResponse20046 {
	this := InlineResponse20046{}
	return &this
}

// GetSTATUS returns the STATUS field value if set, zero value otherwise.
func (o *InlineResponse20046) GetSTATUS() string {
	if o == nil || o.STATUS == nil {
		var ret string
		return ret
	}
	return *o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20046) GetSTATUSOk() (*string, bool) {
	if o == nil || o.STATUS == nil {
		return nil, false
	}
	return o.STATUS, true
}

// HasSTATUS returns a boolean if a field has been set.
func (o *InlineResponse20046) HasSTATUS() bool {
	if o != nil && o.STATUS != nil {
		return true
	}

	return false
}

// SetSTATUS gets a reference to the given string and assigns it to the STATUS field.
func (o *InlineResponse20046) SetSTATUS(v string) {
	o.STATUS = &v
}

// GetCards returns the Cards field value if set, zero value otherwise.
func (o *InlineResponse20046) GetCards() []InlineResponse20046Cards {
	if o == nil || o.Cards == nil {
		var ret []InlineResponse20046Cards
		return ret
	}
	return *o.Cards
}

// GetCardsOk returns a tuple with the Cards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20046) GetCardsOk() (*[]InlineResponse20046Cards, bool) {
	if o == nil || o.Cards == nil {
		return nil, false
	}
	return o.Cards, true
}

// HasCards returns a boolean if a field has been set.
func (o *InlineResponse20046) HasCards() bool {
	if o != nil && o.Cards != nil {
		return true
	}

	return false
}

// SetCards gets a reference to the given []InlineResponse20046Cards and assigns it to the Cards field.
func (o *InlineResponse20046) SetCards(v []InlineResponse20046Cards) {
	o.Cards = &v
}

// GetColumn returns the Column field value if set, zero value otherwise.
func (o *InlineResponse20046) GetColumn() InlineResponse2002Column {
	if o == nil || o.Column == nil {
		var ret InlineResponse2002Column
		return ret
	}
	return *o.Column
}

// GetColumnOk returns a tuple with the Column field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20046) GetColumnOk() (*InlineResponse2002Column, bool) {
	if o == nil || o.Column == nil {
		return nil, false
	}
	return o.Column, true
}

// HasColumn returns a boolean if a field has been set.
func (o *InlineResponse20046) HasColumn() bool {
	if o != nil && o.Column != nil {
		return true
	}

	return false
}

// SetColumn gets a reference to the given InlineResponse2002Column and assigns it to the Column field.
func (o *InlineResponse20046) SetColumn(v InlineResponse2002Column) {
	o.Column = &v
}

func (o InlineResponse20046) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.STATUS != nil {
		toSerialize["STATUS"] = o.STATUS
	}
	if o.Cards != nil {
		toSerialize["cards"] = o.Cards
	}
	if o.Column != nil {
		toSerialize["column"] = o.Column
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20046 struct {
	value *InlineResponse20046
	isSet bool
}

func (v NullableInlineResponse20046) Get() *InlineResponse20046 {
	return v.value
}

func (v *NullableInlineResponse20046) Set(val *InlineResponse20046) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20046) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20046) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20046(val *InlineResponse20046) *NullableInlineResponse20046 {
	return &NullableInlineResponse20046{value: val, isSet: true}
}

func (v NullableInlineResponse20046) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20046) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


