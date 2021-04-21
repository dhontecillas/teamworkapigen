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

// InlineObject104 struct for InlineObject104
type InlineObject104 struct {
	ProjectId *int32 `json:"projectId,omitempty"`
	TaskId *int32 `json:"taskId,omitempty"`
	TasklistId *int32 `json:"tasklistId,omitempty"`
}

// NewInlineObject104 instantiates a new InlineObject104 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject104() *InlineObject104 {
	this := InlineObject104{}
	return &this
}

// NewInlineObject104WithDefaults instantiates a new InlineObject104 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject104WithDefaults() *InlineObject104 {
	this := InlineObject104{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *InlineObject104) GetProjectId() int32 {
	if o == nil || o.ProjectId == nil {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject104) GetProjectIdOk() (*int32, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *InlineObject104) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *InlineObject104) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *InlineObject104) GetTaskId() int32 {
	if o == nil || o.TaskId == nil {
		var ret int32
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject104) GetTaskIdOk() (*int32, bool) {
	if o == nil || o.TaskId == nil {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *InlineObject104) HasTaskId() bool {
	if o != nil && o.TaskId != nil {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given int32 and assigns it to the TaskId field.
func (o *InlineObject104) SetTaskId(v int32) {
	o.TaskId = &v
}

// GetTasklistId returns the TasklistId field value if set, zero value otherwise.
func (o *InlineObject104) GetTasklistId() int32 {
	if o == nil || o.TasklistId == nil {
		var ret int32
		return ret
	}
	return *o.TasklistId
}

// GetTasklistIdOk returns a tuple with the TasklistId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject104) GetTasklistIdOk() (*int32, bool) {
	if o == nil || o.TasklistId == nil {
		return nil, false
	}
	return o.TasklistId, true
}

// HasTasklistId returns a boolean if a field has been set.
func (o *InlineObject104) HasTasklistId() bool {
	if o != nil && o.TasklistId != nil {
		return true
	}

	return false
}

// SetTasklistId gets a reference to the given int32 and assigns it to the TasklistId field.
func (o *InlineObject104) SetTasklistId(v int32) {
	o.TasklistId = &v
}

func (o InlineObject104) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.TaskId != nil {
		toSerialize["taskId"] = o.TaskId
	}
	if o.TasklistId != nil {
		toSerialize["tasklistId"] = o.TasklistId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject104 struct {
	value *InlineObject104
	isSet bool
}

func (v NullableInlineObject104) Get() *InlineObject104 {
	return v.value
}

func (v *NullableInlineObject104) Set(val *InlineObject104) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject104) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject104) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject104(val *InlineObject104) *NullableInlineObject104 {
	return &NullableInlineObject104{value: val, isSet: true}
}

func (v NullableInlineObject104) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject104) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

