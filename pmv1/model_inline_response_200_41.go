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

// InlineResponse20041 struct for InlineResponse20041
type InlineResponse20041 struct {
	STATUS *string `json:"STATUS,omitempty"`
	ClockIns *[]InlineResponse20041ClockIns `json:"clockIns,omitempty"`
}

// NewInlineResponse20041 instantiates a new InlineResponse20041 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20041() *InlineResponse20041 {
	this := InlineResponse20041{}
	return &this
}

// NewInlineResponse20041WithDefaults instantiates a new InlineResponse20041 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20041WithDefaults() *InlineResponse20041 {
	this := InlineResponse20041{}
	return &this
}

// GetSTATUS returns the STATUS field value if set, zero value otherwise.
func (o *InlineResponse20041) GetSTATUS() string {
	if o == nil || o.STATUS == nil {
		var ret string
		return ret
	}
	return *o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20041) GetSTATUSOk() (*string, bool) {
	if o == nil || o.STATUS == nil {
		return nil, false
	}
	return o.STATUS, true
}

// HasSTATUS returns a boolean if a field has been set.
func (o *InlineResponse20041) HasSTATUS() bool {
	if o != nil && o.STATUS != nil {
		return true
	}

	return false
}

// SetSTATUS gets a reference to the given string and assigns it to the STATUS field.
func (o *InlineResponse20041) SetSTATUS(v string) {
	o.STATUS = &v
}

// GetClockIns returns the ClockIns field value if set, zero value otherwise.
func (o *InlineResponse20041) GetClockIns() []InlineResponse20041ClockIns {
	if o == nil || o.ClockIns == nil {
		var ret []InlineResponse20041ClockIns
		return ret
	}
	return *o.ClockIns
}

// GetClockInsOk returns a tuple with the ClockIns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20041) GetClockInsOk() (*[]InlineResponse20041ClockIns, bool) {
	if o == nil || o.ClockIns == nil {
		return nil, false
	}
	return o.ClockIns, true
}

// HasClockIns returns a boolean if a field has been set.
func (o *InlineResponse20041) HasClockIns() bool {
	if o != nil && o.ClockIns != nil {
		return true
	}

	return false
}

// SetClockIns gets a reference to the given []InlineResponse20041ClockIns and assigns it to the ClockIns field.
func (o *InlineResponse20041) SetClockIns(v []InlineResponse20041ClockIns) {
	o.ClockIns = &v
}

func (o InlineResponse20041) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.STATUS != nil {
		toSerialize["STATUS"] = o.STATUS
	}
	if o.ClockIns != nil {
		toSerialize["clockIns"] = o.ClockIns
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20041 struct {
	value *InlineResponse20041
	isSet bool
}

func (v NullableInlineResponse20041) Get() *InlineResponse20041 {
	return v.value
}

func (v *NullableInlineResponse20041) Set(val *InlineResponse20041) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20041) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20041) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20041(val *InlineResponse20041) *NullableInlineResponse20041 {
	return &NullableInlineResponse20041{value: val, isSet: true}
}

func (v NullableInlineResponse20041) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20041) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

