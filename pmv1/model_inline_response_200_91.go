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

// InlineResponse20091 struct for InlineResponse20091
type InlineResponse20091 struct {
	UserStatuses *[]InlineResponse20091UserStatuses `json:"userStatuses,omitempty"`
}

// NewInlineResponse20091 instantiates a new InlineResponse20091 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20091() *InlineResponse20091 {
	this := InlineResponse20091{}
	return &this
}

// NewInlineResponse20091WithDefaults instantiates a new InlineResponse20091 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20091WithDefaults() *InlineResponse20091 {
	this := InlineResponse20091{}
	return &this
}

// GetUserStatuses returns the UserStatuses field value if set, zero value otherwise.
func (o *InlineResponse20091) GetUserStatuses() []InlineResponse20091UserStatuses {
	if o == nil || o.UserStatuses == nil {
		var ret []InlineResponse20091UserStatuses
		return ret
	}
	return *o.UserStatuses
}

// GetUserStatusesOk returns a tuple with the UserStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20091) GetUserStatusesOk() (*[]InlineResponse20091UserStatuses, bool) {
	if o == nil || o.UserStatuses == nil {
		return nil, false
	}
	return o.UserStatuses, true
}

// HasUserStatuses returns a boolean if a field has been set.
func (o *InlineResponse20091) HasUserStatuses() bool {
	if o != nil && o.UserStatuses != nil {
		return true
	}

	return false
}

// SetUserStatuses gets a reference to the given []InlineResponse20091UserStatuses and assigns it to the UserStatuses field.
func (o *InlineResponse20091) SetUserStatuses(v []InlineResponse20091UserStatuses) {
	o.UserStatuses = &v
}

func (o InlineResponse20091) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserStatuses != nil {
		toSerialize["userStatuses"] = o.UserStatuses
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20091 struct {
	value *InlineResponse20091
	isSet bool
}

func (v NullableInlineResponse20091) Get() *InlineResponse20091 {
	return v.value
}

func (v *NullableInlineResponse20091) Set(val *InlineResponse20091) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20091) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20091) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20091(val *InlineResponse20091) *NullableInlineResponse20091 {
	return &NullableInlineResponse20091{value: val, isSet: true}
}

func (v NullableInlineResponse20091) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20091) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

