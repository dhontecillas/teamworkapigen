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

// InlineResponse20020 struct for InlineResponse20020
type InlineResponse20020 struct {
	Category *InlineResponse20020Category `json:"category,omitempty"`
}

// NewInlineResponse20020 instantiates a new InlineResponse20020 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20020() *InlineResponse20020 {
	this := InlineResponse20020{}
	return &this
}

// NewInlineResponse20020WithDefaults instantiates a new InlineResponse20020 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20020WithDefaults() *InlineResponse20020 {
	this := InlineResponse20020{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *InlineResponse20020) GetCategory() InlineResponse20020Category {
	if o == nil || o.Category == nil {
		var ret InlineResponse20020Category
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20020) GetCategoryOk() (*InlineResponse20020Category, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *InlineResponse20020) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given InlineResponse20020Category and assigns it to the Category field.
func (o *InlineResponse20020) SetCategory(v InlineResponse20020Category) {
	o.Category = &v
}

func (o InlineResponse20020) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20020 struct {
	value *InlineResponse20020
	isSet bool
}

func (v NullableInlineResponse20020) Get() *InlineResponse20020 {
	return v.value
}

func (v *NullableInlineResponse20020) Set(val *InlineResponse20020) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20020) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20020) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20020(val *InlineResponse20020) *NullableInlineResponse20020 {
	return &NullableInlineResponse20020{value: val, isSet: true}
}

func (v NullableInlineResponse20020) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20020) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


