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

// InlineResponse200103 struct for InlineResponse200103
type InlineResponse200103 struct {
	Project *InlineResponse200103Project `json:"project,omitempty"`
}

// NewInlineResponse200103 instantiates a new InlineResponse200103 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200103() *InlineResponse200103 {
	this := InlineResponse200103{}
	return &this
}

// NewInlineResponse200103WithDefaults instantiates a new InlineResponse200103 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200103WithDefaults() *InlineResponse200103 {
	this := InlineResponse200103{}
	return &this
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *InlineResponse200103) GetProject() InlineResponse200103Project {
	if o == nil || o.Project == nil {
		var ret InlineResponse200103Project
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200103) GetProjectOk() (*InlineResponse200103Project, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *InlineResponse200103) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given InlineResponse200103Project and assigns it to the Project field.
func (o *InlineResponse200103) SetProject(v InlineResponse200103Project) {
	o.Project = &v
}

func (o InlineResponse200103) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200103 struct {
	value *InlineResponse200103
	isSet bool
}

func (v NullableInlineResponse200103) Get() *InlineResponse200103 {
	return v.value
}

func (v *NullableInlineResponse200103) Set(val *InlineResponse200103) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200103) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200103) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200103(val *InlineResponse200103) *NullableInlineResponse200103 {
	return &NullableInlineResponse200103{value: val, isSet: true}
}

func (v NullableInlineResponse200103) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200103) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


