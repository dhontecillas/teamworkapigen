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

// OwnerProjectMetricOwner ProjectMetricOwner contains information about a specific owner.
type OwnerProjectMetricOwner struct {
	Name *string `json:"name,omitempty"`
	Value *int32 `json:"value,omitempty"`
}

// NewOwnerProjectMetricOwner instantiates a new OwnerProjectMetricOwner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOwnerProjectMetricOwner() *OwnerProjectMetricOwner {
	this := OwnerProjectMetricOwner{}
	return &this
}

// NewOwnerProjectMetricOwnerWithDefaults instantiates a new OwnerProjectMetricOwner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOwnerProjectMetricOwnerWithDefaults() *OwnerProjectMetricOwner {
	this := OwnerProjectMetricOwner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OwnerProjectMetricOwner) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerProjectMetricOwner) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OwnerProjectMetricOwner) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OwnerProjectMetricOwner) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *OwnerProjectMetricOwner) GetValue() int32 {
	if o == nil || o.Value == nil {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerProjectMetricOwner) GetValueOk() (*int32, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *OwnerProjectMetricOwner) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *OwnerProjectMetricOwner) SetValue(v int32) {
	o.Value = &v
}

func (o OwnerProjectMetricOwner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableOwnerProjectMetricOwner struct {
	value *OwnerProjectMetricOwner
	isSet bool
}

func (v NullableOwnerProjectMetricOwner) Get() *OwnerProjectMetricOwner {
	return v.value
}

func (v *NullableOwnerProjectMetricOwner) Set(val *OwnerProjectMetricOwner) {
	v.value = val
	v.isSet = true
}

func (v NullableOwnerProjectMetricOwner) IsSet() bool {
	return v.isSet
}

func (v *NullableOwnerProjectMetricOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOwnerProjectMetricOwner(val *OwnerProjectMetricOwner) *NullableOwnerProjectMetricOwner {
	return &NullableOwnerProjectMetricOwner{value: val, isSet: true}
}

func (v NullableOwnerProjectMetricOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOwnerProjectMetricOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


