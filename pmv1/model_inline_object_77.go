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

// InlineObject77 struct for InlineObject77
type InlineObject77 struct {
	TodoItem *ProjectsProjIdTasksJsonTodoItem `json:"todo-item,omitempty"`
}

// NewInlineObject77 instantiates a new InlineObject77 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject77() *InlineObject77 {
	this := InlineObject77{}
	return &this
}

// NewInlineObject77WithDefaults instantiates a new InlineObject77 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject77WithDefaults() *InlineObject77 {
	this := InlineObject77{}
	return &this
}

// GetTodoItem returns the TodoItem field value if set, zero value otherwise.
func (o *InlineObject77) GetTodoItem() ProjectsProjIdTasksJsonTodoItem {
	if o == nil || o.TodoItem == nil {
		var ret ProjectsProjIdTasksJsonTodoItem
		return ret
	}
	return *o.TodoItem
}

// GetTodoItemOk returns a tuple with the TodoItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject77) GetTodoItemOk() (*ProjectsProjIdTasksJsonTodoItem, bool) {
	if o == nil || o.TodoItem == nil {
		return nil, false
	}
	return o.TodoItem, true
}

// HasTodoItem returns a boolean if a field has been set.
func (o *InlineObject77) HasTodoItem() bool {
	if o != nil && o.TodoItem != nil {
		return true
	}

	return false
}

// SetTodoItem gets a reference to the given ProjectsProjIdTasksJsonTodoItem and assigns it to the TodoItem field.
func (o *InlineObject77) SetTodoItem(v ProjectsProjIdTasksJsonTodoItem) {
	o.TodoItem = &v
}

func (o InlineObject77) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TodoItem != nil {
		toSerialize["todo-item"] = o.TodoItem
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject77 struct {
	value *InlineObject77
	isSet bool
}

func (v NullableInlineObject77) Get() *InlineObject77 {
	return v.value
}

func (v *NullableInlineObject77) Set(val *InlineObject77) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject77) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject77) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject77(val *InlineObject77) *NullableInlineObject77 {
	return &NullableInlineObject77{value: val, isSet: true}
}

func (v NullableInlineObject77) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject77) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


