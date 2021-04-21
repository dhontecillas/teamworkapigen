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

// LateTaskMetricLate TaskMetricLate contains count information about late tasks.
type LateTaskMetricLate struct {
	Value *int32 `json:"value,omitempty"`
}

// NewLateTaskMetricLate instantiates a new LateTaskMetricLate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLateTaskMetricLate() *LateTaskMetricLate {
	this := LateTaskMetricLate{}
	return &this
}

// NewLateTaskMetricLateWithDefaults instantiates a new LateTaskMetricLate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLateTaskMetricLateWithDefaults() *LateTaskMetricLate {
	this := LateTaskMetricLate{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *LateTaskMetricLate) GetValue() int32 {
	if o == nil || o.Value == nil {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LateTaskMetricLate) GetValueOk() (*int32, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *LateTaskMetricLate) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *LateTaskMetricLate) SetValue(v int32) {
	o.Value = &v
}

func (o LateTaskMetricLate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableLateTaskMetricLate struct {
	value *LateTaskMetricLate
	isSet bool
}

func (v NullableLateTaskMetricLate) Get() *LateTaskMetricLate {
	return v.value
}

func (v *NullableLateTaskMetricLate) Set(val *LateTaskMetricLate) {
	v.value = val
	v.isSet = true
}

func (v NullableLateTaskMetricLate) IsSet() bool {
	return v.isSet
}

func (v *NullableLateTaskMetricLate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLateTaskMetricLate(val *LateTaskMetricLate) *NullableLateTaskMetricLate {
	return &NullableLateTaskMetricLate{value: val, isSet: true}
}

func (v NullableLateTaskMetricLate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLateTaskMetricLate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

