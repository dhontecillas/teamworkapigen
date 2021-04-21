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

// InlineResponse20016CurrencyCodes struct for InlineResponse20016CurrencyCodes
type InlineResponse20016CurrencyCodes struct {
	Code *string `json:"code,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewInlineResponse20016CurrencyCodes instantiates a new InlineResponse20016CurrencyCodes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20016CurrencyCodes() *InlineResponse20016CurrencyCodes {
	this := InlineResponse20016CurrencyCodes{}
	return &this
}

// NewInlineResponse20016CurrencyCodesWithDefaults instantiates a new InlineResponse20016CurrencyCodes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20016CurrencyCodesWithDefaults() *InlineResponse20016CurrencyCodes {
	this := InlineResponse20016CurrencyCodes{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *InlineResponse20016CurrencyCodes) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20016CurrencyCodes) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *InlineResponse20016CurrencyCodes) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *InlineResponse20016CurrencyCodes) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20016CurrencyCodes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20016CurrencyCodes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20016CurrencyCodes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20016CurrencyCodes) SetName(v string) {
	o.Name = &v
}

func (o InlineResponse20016CurrencyCodes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20016CurrencyCodes struct {
	value *InlineResponse20016CurrencyCodes
	isSet bool
}

func (v NullableInlineResponse20016CurrencyCodes) Get() *InlineResponse20016CurrencyCodes {
	return v.value
}

func (v *NullableInlineResponse20016CurrencyCodes) Set(val *InlineResponse20016CurrencyCodes) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20016CurrencyCodes) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20016CurrencyCodes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20016CurrencyCodes(val *InlineResponse20016CurrencyCodes) *NullableInlineResponse20016CurrencyCodes {
	return &NullableInlineResponse20016CurrencyCodes{value: val, isSet: true}
}

func (v NullableInlineResponse20016CurrencyCodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20016CurrencyCodes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

