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

// InlineResponse20093 struct for InlineResponse20093
type InlineResponse20093 struct {
	STATUS *string `json:"STATUS,omitempty"`
	Tag *InlineResponse2002Column `json:"tag,omitempty"`
}

// NewInlineResponse20093 instantiates a new InlineResponse20093 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20093() *InlineResponse20093 {
	this := InlineResponse20093{}
	return &this
}

// NewInlineResponse20093WithDefaults instantiates a new InlineResponse20093 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20093WithDefaults() *InlineResponse20093 {
	this := InlineResponse20093{}
	return &this
}

// GetSTATUS returns the STATUS field value if set, zero value otherwise.
func (o *InlineResponse20093) GetSTATUS() string {
	if o == nil || o.STATUS == nil {
		var ret string
		return ret
	}
	return *o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20093) GetSTATUSOk() (*string, bool) {
	if o == nil || o.STATUS == nil {
		return nil, false
	}
	return o.STATUS, true
}

// HasSTATUS returns a boolean if a field has been set.
func (o *InlineResponse20093) HasSTATUS() bool {
	if o != nil && o.STATUS != nil {
		return true
	}

	return false
}

// SetSTATUS gets a reference to the given string and assigns it to the STATUS field.
func (o *InlineResponse20093) SetSTATUS(v string) {
	o.STATUS = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *InlineResponse20093) GetTag() InlineResponse2002Column {
	if o == nil || o.Tag == nil {
		var ret InlineResponse2002Column
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20093) GetTagOk() (*InlineResponse2002Column, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *InlineResponse20093) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given InlineResponse2002Column and assigns it to the Tag field.
func (o *InlineResponse20093) SetTag(v InlineResponse2002Column) {
	o.Tag = &v
}

func (o InlineResponse20093) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.STATUS != nil {
		toSerialize["STATUS"] = o.STATUS
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20093 struct {
	value *InlineResponse20093
	isSet bool
}

func (v NullableInlineResponse20093) Get() *InlineResponse20093 {
	return v.value
}

func (v *NullableInlineResponse20093) Set(val *InlineResponse20093) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20093) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20093) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20093(val *InlineResponse20093) *NullableInlineResponse20093 {
	return &NullableInlineResponse20093{value: val, isSet: true}
}

func (v NullableInlineResponse20093) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20093) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

