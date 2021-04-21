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

// InlineObject57 struct for InlineObject57
type InlineObject57 struct {
	Emailaddress *ProjectsIdEmailaddressJsonEmailaddress `json:"emailaddress,omitempty"`
}

// NewInlineObject57 instantiates a new InlineObject57 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject57() *InlineObject57 {
	this := InlineObject57{}
	return &this
}

// NewInlineObject57WithDefaults instantiates a new InlineObject57 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject57WithDefaults() *InlineObject57 {
	this := InlineObject57{}
	return &this
}

// GetEmailaddress returns the Emailaddress field value if set, zero value otherwise.
func (o *InlineObject57) GetEmailaddress() ProjectsIdEmailaddressJsonEmailaddress {
	if o == nil || o.Emailaddress == nil {
		var ret ProjectsIdEmailaddressJsonEmailaddress
		return ret
	}
	return *o.Emailaddress
}

// GetEmailaddressOk returns a tuple with the Emailaddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject57) GetEmailaddressOk() (*ProjectsIdEmailaddressJsonEmailaddress, bool) {
	if o == nil || o.Emailaddress == nil {
		return nil, false
	}
	return o.Emailaddress, true
}

// HasEmailaddress returns a boolean if a field has been set.
func (o *InlineObject57) HasEmailaddress() bool {
	if o != nil && o.Emailaddress != nil {
		return true
	}

	return false
}

// SetEmailaddress gets a reference to the given ProjectsIdEmailaddressJsonEmailaddress and assigns it to the Emailaddress field.
func (o *InlineObject57) SetEmailaddress(v ProjectsIdEmailaddressJsonEmailaddress) {
	o.Emailaddress = &v
}

func (o InlineObject57) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Emailaddress != nil {
		toSerialize["emailaddress"] = o.Emailaddress
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject57 struct {
	value *InlineObject57
	isSet bool
}

func (v NullableInlineObject57) Get() *InlineObject57 {
	return v.value
}

func (v *NullableInlineObject57) Set(val *InlineObject57) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject57) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject57) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject57(val *InlineObject57) *NullableInlineObject57 {
	return &NullableInlineObject57{value: val, isSet: true}
}

func (v NullableInlineObject57) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject57) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

