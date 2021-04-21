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

// InlineObject86 struct for InlineObject86
type InlineObject86 struct {
	Tag *EventtypesJsonEventtype `json:"tag,omitempty"`
}

// NewInlineObject86 instantiates a new InlineObject86 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject86() *InlineObject86 {
	this := InlineObject86{}
	return &this
}

// NewInlineObject86WithDefaults instantiates a new InlineObject86 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject86WithDefaults() *InlineObject86 {
	this := InlineObject86{}
	return &this
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *InlineObject86) GetTag() EventtypesJsonEventtype {
	if o == nil || o.Tag == nil {
		var ret EventtypesJsonEventtype
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject86) GetTagOk() (*EventtypesJsonEventtype, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *InlineObject86) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given EventtypesJsonEventtype and assigns it to the Tag field.
func (o *InlineObject86) SetTag(v EventtypesJsonEventtype) {
	o.Tag = &v
}

func (o InlineObject86) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject86 struct {
	value *InlineObject86
	isSet bool
}

func (v NullableInlineObject86) Get() *InlineObject86 {
	return v.value
}

func (v *NullableInlineObject86) Set(val *InlineObject86) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject86) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject86) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject86(val *InlineObject86) *NullableInlineObject86 {
	return &NullableInlineObject86{value: val, isSet: true}
}

func (v NullableInlineObject86) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject86) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

