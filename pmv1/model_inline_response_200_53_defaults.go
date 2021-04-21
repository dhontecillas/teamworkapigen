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

// InlineResponse20053Defaults struct for InlineResponse20053Defaults
type InlineResponse20053Defaults struct {
	Privacy *string `json:"privacy,omitempty"`
}

// NewInlineResponse20053Defaults instantiates a new InlineResponse20053Defaults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20053Defaults() *InlineResponse20053Defaults {
	this := InlineResponse20053Defaults{}
	return &this
}

// NewInlineResponse20053DefaultsWithDefaults instantiates a new InlineResponse20053Defaults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20053DefaultsWithDefaults() *InlineResponse20053Defaults {
	this := InlineResponse20053Defaults{}
	return &this
}

// GetPrivacy returns the Privacy field value if set, zero value otherwise.
func (o *InlineResponse20053Defaults) GetPrivacy() string {
	if o == nil || o.Privacy == nil {
		var ret string
		return ret
	}
	return *o.Privacy
}

// GetPrivacyOk returns a tuple with the Privacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Defaults) GetPrivacyOk() (*string, bool) {
	if o == nil || o.Privacy == nil {
		return nil, false
	}
	return o.Privacy, true
}

// HasPrivacy returns a boolean if a field has been set.
func (o *InlineResponse20053Defaults) HasPrivacy() bool {
	if o != nil && o.Privacy != nil {
		return true
	}

	return false
}

// SetPrivacy gets a reference to the given string and assigns it to the Privacy field.
func (o *InlineResponse20053Defaults) SetPrivacy(v string) {
	o.Privacy = &v
}

func (o InlineResponse20053Defaults) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Privacy != nil {
		toSerialize["privacy"] = o.Privacy
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20053Defaults struct {
	value *InlineResponse20053Defaults
	isSet bool
}

func (v NullableInlineResponse20053Defaults) Get() *InlineResponse20053Defaults {
	return v.value
}

func (v *NullableInlineResponse20053Defaults) Set(val *InlineResponse20053Defaults) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20053Defaults) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20053Defaults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20053Defaults(val *InlineResponse20053Defaults) *NullableInlineResponse20053Defaults {
	return &NullableInlineResponse20053Defaults{value: val, isSet: true}
}

func (v NullableInlineResponse20053Defaults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20053Defaults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

