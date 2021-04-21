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

// InlineResponse20027PersonPermissions struct for InlineResponse20027PersonPermissions
type InlineResponse20027PersonPermissions struct {
	CanAddProjects *bool `json:"can-add-projects,omitempty"`
	CanManagePeople *bool `json:"can-manage-people,omitempty"`
}

// NewInlineResponse20027PersonPermissions instantiates a new InlineResponse20027PersonPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20027PersonPermissions() *InlineResponse20027PersonPermissions {
	this := InlineResponse20027PersonPermissions{}
	return &this
}

// NewInlineResponse20027PersonPermissionsWithDefaults instantiates a new InlineResponse20027PersonPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20027PersonPermissionsWithDefaults() *InlineResponse20027PersonPermissions {
	this := InlineResponse20027PersonPermissions{}
	return &this
}

// GetCanAddProjects returns the CanAddProjects field value if set, zero value otherwise.
func (o *InlineResponse20027PersonPermissions) GetCanAddProjects() bool {
	if o == nil || o.CanAddProjects == nil {
		var ret bool
		return ret
	}
	return *o.CanAddProjects
}

// GetCanAddProjectsOk returns a tuple with the CanAddProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027PersonPermissions) GetCanAddProjectsOk() (*bool, bool) {
	if o == nil || o.CanAddProjects == nil {
		return nil, false
	}
	return o.CanAddProjects, true
}

// HasCanAddProjects returns a boolean if a field has been set.
func (o *InlineResponse20027PersonPermissions) HasCanAddProjects() bool {
	if o != nil && o.CanAddProjects != nil {
		return true
	}

	return false
}

// SetCanAddProjects gets a reference to the given bool and assigns it to the CanAddProjects field.
func (o *InlineResponse20027PersonPermissions) SetCanAddProjects(v bool) {
	o.CanAddProjects = &v
}

// GetCanManagePeople returns the CanManagePeople field value if set, zero value otherwise.
func (o *InlineResponse20027PersonPermissions) GetCanManagePeople() bool {
	if o == nil || o.CanManagePeople == nil {
		var ret bool
		return ret
	}
	return *o.CanManagePeople
}

// GetCanManagePeopleOk returns a tuple with the CanManagePeople field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027PersonPermissions) GetCanManagePeopleOk() (*bool, bool) {
	if o == nil || o.CanManagePeople == nil {
		return nil, false
	}
	return o.CanManagePeople, true
}

// HasCanManagePeople returns a boolean if a field has been set.
func (o *InlineResponse20027PersonPermissions) HasCanManagePeople() bool {
	if o != nil && o.CanManagePeople != nil {
		return true
	}

	return false
}

// SetCanManagePeople gets a reference to the given bool and assigns it to the CanManagePeople field.
func (o *InlineResponse20027PersonPermissions) SetCanManagePeople(v bool) {
	o.CanManagePeople = &v
}

func (o InlineResponse20027PersonPermissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CanAddProjects != nil {
		toSerialize["can-add-projects"] = o.CanAddProjects
	}
	if o.CanManagePeople != nil {
		toSerialize["can-manage-people"] = o.CanManagePeople
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20027PersonPermissions struct {
	value *InlineResponse20027PersonPermissions
	isSet bool
}

func (v NullableInlineResponse20027PersonPermissions) Get() *InlineResponse20027PersonPermissions {
	return v.value
}

func (v *NullableInlineResponse20027PersonPermissions) Set(val *InlineResponse20027PersonPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20027PersonPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20027PersonPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20027PersonPermissions(val *InlineResponse20027PersonPermissions) *NullableInlineResponse20027PersonPermissions {
	return &NullableInlineResponse20027PersonPermissions{value: val, isSet: true}
}

func (v NullableInlineResponse20027PersonPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20027PersonPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

