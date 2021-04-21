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

// InlineResponse20039PersonPermissions struct for InlineResponse20039PersonPermissions
type InlineResponse20039PersonPermissions struct {
	CanAccessTemplates *bool `json:"can-access-templates,omitempty"`
	CanAddProjects *bool `json:"can-add-projects,omitempty"`
	CanManagePeople *bool `json:"can-manage-people,omitempty"`
}

// NewInlineResponse20039PersonPermissions instantiates a new InlineResponse20039PersonPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20039PersonPermissions() *InlineResponse20039PersonPermissions {
	this := InlineResponse20039PersonPermissions{}
	return &this
}

// NewInlineResponse20039PersonPermissionsWithDefaults instantiates a new InlineResponse20039PersonPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20039PersonPermissionsWithDefaults() *InlineResponse20039PersonPermissions {
	this := InlineResponse20039PersonPermissions{}
	return &this
}

// GetCanAccessTemplates returns the CanAccessTemplates field value if set, zero value otherwise.
func (o *InlineResponse20039PersonPermissions) GetCanAccessTemplates() bool {
	if o == nil || o.CanAccessTemplates == nil {
		var ret bool
		return ret
	}
	return *o.CanAccessTemplates
}

// GetCanAccessTemplatesOk returns a tuple with the CanAccessTemplates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20039PersonPermissions) GetCanAccessTemplatesOk() (*bool, bool) {
	if o == nil || o.CanAccessTemplates == nil {
		return nil, false
	}
	return o.CanAccessTemplates, true
}

// HasCanAccessTemplates returns a boolean if a field has been set.
func (o *InlineResponse20039PersonPermissions) HasCanAccessTemplates() bool {
	if o != nil && o.CanAccessTemplates != nil {
		return true
	}

	return false
}

// SetCanAccessTemplates gets a reference to the given bool and assigns it to the CanAccessTemplates field.
func (o *InlineResponse20039PersonPermissions) SetCanAccessTemplates(v bool) {
	o.CanAccessTemplates = &v
}

// GetCanAddProjects returns the CanAddProjects field value if set, zero value otherwise.
func (o *InlineResponse20039PersonPermissions) GetCanAddProjects() bool {
	if o == nil || o.CanAddProjects == nil {
		var ret bool
		return ret
	}
	return *o.CanAddProjects
}

// GetCanAddProjectsOk returns a tuple with the CanAddProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20039PersonPermissions) GetCanAddProjectsOk() (*bool, bool) {
	if o == nil || o.CanAddProjects == nil {
		return nil, false
	}
	return o.CanAddProjects, true
}

// HasCanAddProjects returns a boolean if a field has been set.
func (o *InlineResponse20039PersonPermissions) HasCanAddProjects() bool {
	if o != nil && o.CanAddProjects != nil {
		return true
	}

	return false
}

// SetCanAddProjects gets a reference to the given bool and assigns it to the CanAddProjects field.
func (o *InlineResponse20039PersonPermissions) SetCanAddProjects(v bool) {
	o.CanAddProjects = &v
}

// GetCanManagePeople returns the CanManagePeople field value if set, zero value otherwise.
func (o *InlineResponse20039PersonPermissions) GetCanManagePeople() bool {
	if o == nil || o.CanManagePeople == nil {
		var ret bool
		return ret
	}
	return *o.CanManagePeople
}

// GetCanManagePeopleOk returns a tuple with the CanManagePeople field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20039PersonPermissions) GetCanManagePeopleOk() (*bool, bool) {
	if o == nil || o.CanManagePeople == nil {
		return nil, false
	}
	return o.CanManagePeople, true
}

// HasCanManagePeople returns a boolean if a field has been set.
func (o *InlineResponse20039PersonPermissions) HasCanManagePeople() bool {
	if o != nil && o.CanManagePeople != nil {
		return true
	}

	return false
}

// SetCanManagePeople gets a reference to the given bool and assigns it to the CanManagePeople field.
func (o *InlineResponse20039PersonPermissions) SetCanManagePeople(v bool) {
	o.CanManagePeople = &v
}

func (o InlineResponse20039PersonPermissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CanAccessTemplates != nil {
		toSerialize["can-access-templates"] = o.CanAccessTemplates
	}
	if o.CanAddProjects != nil {
		toSerialize["can-add-projects"] = o.CanAddProjects
	}
	if o.CanManagePeople != nil {
		toSerialize["can-manage-people"] = o.CanManagePeople
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20039PersonPermissions struct {
	value *InlineResponse20039PersonPermissions
	isSet bool
}

func (v NullableInlineResponse20039PersonPermissions) Get() *InlineResponse20039PersonPermissions {
	return v.value
}

func (v *NullableInlineResponse20039PersonPermissions) Set(val *InlineResponse20039PersonPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20039PersonPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20039PersonPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20039PersonPermissions(val *InlineResponse20039PersonPermissions) *NullableInlineResponse20039PersonPermissions {
	return &NullableInlineResponse20039PersonPermissions{value: val, isSet: true}
}

func (v NullableInlineResponse20039PersonPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20039PersonPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

