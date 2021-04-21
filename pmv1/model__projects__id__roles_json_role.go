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

// ProjectsIdRolesJsonRole struct for ProjectsIdRolesJsonRole
type ProjectsIdRolesJsonRole struct {
	Description *string `json:"description,omitempty"`
	Name *string `json:"name,omitempty"`
	Users *string `json:"users,omitempty"`
}

// NewProjectsIdRolesJsonRole instantiates a new ProjectsIdRolesJsonRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectsIdRolesJsonRole() *ProjectsIdRolesJsonRole {
	this := ProjectsIdRolesJsonRole{}
	return &this
}

// NewProjectsIdRolesJsonRoleWithDefaults instantiates a new ProjectsIdRolesJsonRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectsIdRolesJsonRoleWithDefaults() *ProjectsIdRolesJsonRole {
	this := ProjectsIdRolesJsonRole{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProjectsIdRolesJsonRole) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdRolesJsonRole) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProjectsIdRolesJsonRole) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProjectsIdRolesJsonRole) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProjectsIdRolesJsonRole) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdRolesJsonRole) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProjectsIdRolesJsonRole) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProjectsIdRolesJsonRole) SetName(v string) {
	o.Name = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *ProjectsIdRolesJsonRole) GetUsers() string {
	if o == nil || o.Users == nil {
		var ret string
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdRolesJsonRole) GetUsersOk() (*string, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *ProjectsIdRolesJsonRole) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given string and assigns it to the Users field.
func (o *ProjectsIdRolesJsonRole) SetUsers(v string) {
	o.Users = &v
}

func (o ProjectsIdRolesJsonRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	return json.Marshal(toSerialize)
}

type NullableProjectsIdRolesJsonRole struct {
	value *ProjectsIdRolesJsonRole
	isSet bool
}

func (v NullableProjectsIdRolesJsonRole) Get() *ProjectsIdRolesJsonRole {
	return v.value
}

func (v *NullableProjectsIdRolesJsonRole) Set(val *ProjectsIdRolesJsonRole) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectsIdRolesJsonRole) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectsIdRolesJsonRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectsIdRolesJsonRole(val *ProjectsIdRolesJsonRole) *NullableProjectsIdRolesJsonRole {
	return &NullableProjectsIdRolesJsonRole{value: val, isSet: true}
}

func (v NullableProjectsIdRolesJsonRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectsIdRolesJsonRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


