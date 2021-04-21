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

// HealthProjectMetricHealthsResponse ProjectMetricHealthsResponse contains information about a group of healths. Following this format to satisfy the Numerics integration.
type HealthProjectMetricHealthsResponse struct {
	Data *[]HealthProjectMetricHealth `json:"data,omitempty"`
}

// NewHealthProjectMetricHealthsResponse instantiates a new HealthProjectMetricHealthsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthProjectMetricHealthsResponse() *HealthProjectMetricHealthsResponse {
	this := HealthProjectMetricHealthsResponse{}
	return &this
}

// NewHealthProjectMetricHealthsResponseWithDefaults instantiates a new HealthProjectMetricHealthsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthProjectMetricHealthsResponseWithDefaults() *HealthProjectMetricHealthsResponse {
	this := HealthProjectMetricHealthsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *HealthProjectMetricHealthsResponse) GetData() []HealthProjectMetricHealth {
	if o == nil || o.Data == nil {
		var ret []HealthProjectMetricHealth
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthProjectMetricHealthsResponse) GetDataOk() (*[]HealthProjectMetricHealth, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *HealthProjectMetricHealthsResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []HealthProjectMetricHealth and assigns it to the Data field.
func (o *HealthProjectMetricHealthsResponse) SetData(v []HealthProjectMetricHealth) {
	o.Data = &v
}

func (o HealthProjectMetricHealthsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableHealthProjectMetricHealthsResponse struct {
	value *HealthProjectMetricHealthsResponse
	isSet bool
}

func (v NullableHealthProjectMetricHealthsResponse) Get() *HealthProjectMetricHealthsResponse {
	return v.value
}

func (v *NullableHealthProjectMetricHealthsResponse) Set(val *HealthProjectMetricHealthsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthProjectMetricHealthsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthProjectMetricHealthsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthProjectMetricHealthsResponse(val *HealthProjectMetricHealthsResponse) *NullableHealthProjectMetricHealthsResponse {
	return &NullableHealthProjectMetricHealthsResponse{value: val, isSet: true}
}

func (v NullableHealthProjectMetricHealthsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthProjectMetricHealthsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

