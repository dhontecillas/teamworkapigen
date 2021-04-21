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

// InlineObject7 struct for InlineObject7
type InlineObject7 struct {
	ClockIn ClockinJsonClockIn `json:"clockIn"`
}

// NewInlineObject7 instantiates a new InlineObject7 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject7(clockIn ClockinJsonClockIn) *InlineObject7 {
	this := InlineObject7{}
	this.ClockIn = clockIn
	return &this
}

// NewInlineObject7WithDefaults instantiates a new InlineObject7 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject7WithDefaults() *InlineObject7 {
	this := InlineObject7{}
	return &this
}

// GetClockIn returns the ClockIn field value
func (o *InlineObject7) GetClockIn() ClockinJsonClockIn {
	if o == nil {
		var ret ClockinJsonClockIn
		return ret
	}

	return o.ClockIn
}

// GetClockInOk returns a tuple with the ClockIn field value
// and a boolean to check if the value has been set.
func (o *InlineObject7) GetClockInOk() (*ClockinJsonClockIn, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClockIn, true
}

// SetClockIn sets field value
func (o *InlineObject7) SetClockIn(v ClockinJsonClockIn) {
	o.ClockIn = v
}

func (o InlineObject7) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["clockIn"] = o.ClockIn
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject7 struct {
	value *InlineObject7
	isSet bool
}

func (v NullableInlineObject7) Get() *InlineObject7 {
	return v.value
}

func (v *NullableInlineObject7) Set(val *InlineObject7) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject7) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject7) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject7(val *InlineObject7) *NullableInlineObject7 {
	return &NullableInlineObject7{value: val, isSet: true}
}

func (v NullableInlineObject7) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject7) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


