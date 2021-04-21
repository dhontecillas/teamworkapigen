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

// InlineResponse20025 struct for InlineResponse20025
type InlineResponse20025 struct {
	Projects *[]InlineResponse20025Projects `json:"projects,omitempty"`
}

// NewInlineResponse20025 instantiates a new InlineResponse20025 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20025() *InlineResponse20025 {
	this := InlineResponse20025{}
	return &this
}

// NewInlineResponse20025WithDefaults instantiates a new InlineResponse20025 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20025WithDefaults() *InlineResponse20025 {
	this := InlineResponse20025{}
	return &this
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *InlineResponse20025) GetProjects() []InlineResponse20025Projects {
	if o == nil || o.Projects == nil {
		var ret []InlineResponse20025Projects
		return ret
	}
	return *o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20025) GetProjectsOk() (*[]InlineResponse20025Projects, bool) {
	if o == nil || o.Projects == nil {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *InlineResponse20025) HasProjects() bool {
	if o != nil && o.Projects != nil {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []InlineResponse20025Projects and assigns it to the Projects field.
func (o *InlineResponse20025) SetProjects(v []InlineResponse20025Projects) {
	o.Projects = &v
}

func (o InlineResponse20025) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20025 struct {
	value *InlineResponse20025
	isSet bool
}

func (v NullableInlineResponse20025) Get() *InlineResponse20025 {
	return v.value
}

func (v *NullableInlineResponse20025) Set(val *InlineResponse20025) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20025) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20025) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20025(val *InlineResponse20025) *NullableInlineResponse20025 {
	return &NullableInlineResponse20025{value: val, isSet: true}
}

func (v NullableInlineResponse20025) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20025) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

