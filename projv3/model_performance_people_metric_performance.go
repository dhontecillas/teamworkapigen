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

// PerformancePeopleMetricPerformance PeopleMetricPerformance contains all the information returned from a performance.
type PerformancePeopleMetricPerformance struct {
	Name *string `json:"name,omitempty"`
	Value *int32 `json:"value,omitempty"`
}

// NewPerformancePeopleMetricPerformance instantiates a new PerformancePeopleMetricPerformance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerformancePeopleMetricPerformance() *PerformancePeopleMetricPerformance {
	this := PerformancePeopleMetricPerformance{}
	return &this
}

// NewPerformancePeopleMetricPerformanceWithDefaults instantiates a new PerformancePeopleMetricPerformance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerformancePeopleMetricPerformanceWithDefaults() *PerformancePeopleMetricPerformance {
	this := PerformancePeopleMetricPerformance{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PerformancePeopleMetricPerformance) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformancePeopleMetricPerformance) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PerformancePeopleMetricPerformance) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PerformancePeopleMetricPerformance) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PerformancePeopleMetricPerformance) GetValue() int32 {
	if o == nil || o.Value == nil {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformancePeopleMetricPerformance) GetValueOk() (*int32, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PerformancePeopleMetricPerformance) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *PerformancePeopleMetricPerformance) SetValue(v int32) {
	o.Value = &v
}

func (o PerformancePeopleMetricPerformance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullablePerformancePeopleMetricPerformance struct {
	value *PerformancePeopleMetricPerformance
	isSet bool
}

func (v NullablePerformancePeopleMetricPerformance) Get() *PerformancePeopleMetricPerformance {
	return v.value
}

func (v *NullablePerformancePeopleMetricPerformance) Set(val *PerformancePeopleMetricPerformance) {
	v.value = val
	v.isSet = true
}

func (v NullablePerformancePeopleMetricPerformance) IsSet() bool {
	return v.isSet
}

func (v *NullablePerformancePeopleMetricPerformance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerformancePeopleMetricPerformance(val *PerformancePeopleMetricPerformance) *NullablePerformancePeopleMetricPerformance {
	return &NullablePerformancePeopleMetricPerformance{value: val, isSet: true}
}

func (v NullablePerformancePeopleMetricPerformance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerformancePeopleMetricPerformance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

