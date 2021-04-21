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

// InlineObject50 struct for InlineObject50
type InlineObject50 struct {
	Project ProjectsJsonProject `json:"project"`
}

// NewInlineObject50 instantiates a new InlineObject50 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject50(project ProjectsJsonProject) *InlineObject50 {
	this := InlineObject50{}
	this.Project = project
	return &this
}

// NewInlineObject50WithDefaults instantiates a new InlineObject50 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject50WithDefaults() *InlineObject50 {
	this := InlineObject50{}
	return &this
}

// GetProject returns the Project field value
func (o *InlineObject50) GetProject() ProjectsJsonProject {
	if o == nil {
		var ret ProjectsJsonProject
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *InlineObject50) GetProjectOk() (*ProjectsJsonProject, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *InlineObject50) SetProject(v ProjectsJsonProject) {
	o.Project = v
}

func (o InlineObject50) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["project"] = o.Project
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject50 struct {
	value *InlineObject50
	isSet bool
}

func (v NullableInlineObject50) Get() *InlineObject50 {
	return v.value
}

func (v *NullableInlineObject50) Set(val *InlineObject50) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject50) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject50) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject50(val *InlineObject50) *NullableInlineObject50 {
	return &NullableInlineObject50{value: val, isSet: true}
}

func (v NullableInlineObject50) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject50) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


