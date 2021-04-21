/*
 * Teamwork.com Projects API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package projv3

import (
	"encoding/json"
)

// ProjectRequest Request contains information of a customfield to be created or updated.
type ProjectRequest struct {
	CustomfieldProject *ProjectCustomFieldProject `json:"customfieldProject,omitempty"`
}

// NewProjectRequest instantiates a new ProjectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectRequest() *ProjectRequest {
	this := ProjectRequest{}
	return &this
}

// NewProjectRequestWithDefaults instantiates a new ProjectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectRequestWithDefaults() *ProjectRequest {
	this := ProjectRequest{}
	return &this
}

// GetCustomfieldProject returns the CustomfieldProject field value if set, zero value otherwise.
func (o *ProjectRequest) GetCustomfieldProject() ProjectCustomFieldProject {
	if o == nil || o.CustomfieldProject == nil {
		var ret ProjectCustomFieldProject
		return ret
	}
	return *o.CustomfieldProject
}

// GetCustomfieldProjectOk returns a tuple with the CustomfieldProject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRequest) GetCustomfieldProjectOk() (*ProjectCustomFieldProject, bool) {
	if o == nil || o.CustomfieldProject == nil {
		return nil, false
	}
	return o.CustomfieldProject, true
}

// HasCustomfieldProject returns a boolean if a field has been set.
func (o *ProjectRequest) HasCustomfieldProject() bool {
	if o != nil && o.CustomfieldProject != nil {
		return true
	}

	return false
}

// SetCustomfieldProject gets a reference to the given ProjectCustomFieldProject and assigns it to the CustomfieldProject field.
func (o *ProjectRequest) SetCustomfieldProject(v ProjectCustomFieldProject) {
	o.CustomfieldProject = &v
}

func (o ProjectRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomfieldProject != nil {
		toSerialize["customfieldProject"] = o.CustomfieldProject
	}
	return json.Marshal(toSerialize)
}

type NullableProjectRequest struct {
	value *ProjectRequest
	isSet bool
}

func (v NullableProjectRequest) Get() *ProjectRequest {
	return v.value
}

func (v *NullableProjectRequest) Set(val *ProjectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectRequest(val *ProjectRequest) *NullableProjectRequest {
	return &NullableProjectRequest{value: val, isSet: true}
}

func (v NullableProjectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


