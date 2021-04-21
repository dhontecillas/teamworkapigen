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

// InlineResponse200124 struct for InlineResponse200124
type InlineResponse200124 struct {
	STATUS *string `json:"STATUS,omitempty"`
	Deliveries *[]InlineResponse200124Deliveries `json:"deliveries,omitempty"`
}

// NewInlineResponse200124 instantiates a new InlineResponse200124 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200124() *InlineResponse200124 {
	this := InlineResponse200124{}
	return &this
}

// NewInlineResponse200124WithDefaults instantiates a new InlineResponse200124 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200124WithDefaults() *InlineResponse200124 {
	this := InlineResponse200124{}
	return &this
}

// GetSTATUS returns the STATUS field value if set, zero value otherwise.
func (o *InlineResponse200124) GetSTATUS() string {
	if o == nil || o.STATUS == nil {
		var ret string
		return ret
	}
	return *o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200124) GetSTATUSOk() (*string, bool) {
	if o == nil || o.STATUS == nil {
		return nil, false
	}
	return o.STATUS, true
}

// HasSTATUS returns a boolean if a field has been set.
func (o *InlineResponse200124) HasSTATUS() bool {
	if o != nil && o.STATUS != nil {
		return true
	}

	return false
}

// SetSTATUS gets a reference to the given string and assigns it to the STATUS field.
func (o *InlineResponse200124) SetSTATUS(v string) {
	o.STATUS = &v
}

// GetDeliveries returns the Deliveries field value if set, zero value otherwise.
func (o *InlineResponse200124) GetDeliveries() []InlineResponse200124Deliveries {
	if o == nil || o.Deliveries == nil {
		var ret []InlineResponse200124Deliveries
		return ret
	}
	return *o.Deliveries
}

// GetDeliveriesOk returns a tuple with the Deliveries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200124) GetDeliveriesOk() (*[]InlineResponse200124Deliveries, bool) {
	if o == nil || o.Deliveries == nil {
		return nil, false
	}
	return o.Deliveries, true
}

// HasDeliveries returns a boolean if a field has been set.
func (o *InlineResponse200124) HasDeliveries() bool {
	if o != nil && o.Deliveries != nil {
		return true
	}

	return false
}

// SetDeliveries gets a reference to the given []InlineResponse200124Deliveries and assigns it to the Deliveries field.
func (o *InlineResponse200124) SetDeliveries(v []InlineResponse200124Deliveries) {
	o.Deliveries = &v
}

func (o InlineResponse200124) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.STATUS != nil {
		toSerialize["STATUS"] = o.STATUS
	}
	if o.Deliveries != nil {
		toSerialize["deliveries"] = o.Deliveries
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200124 struct {
	value *InlineResponse200124
	isSet bool
}

func (v NullableInlineResponse200124) Get() *InlineResponse200124 {
	return v.value
}

func (v *NullableInlineResponse200124) Set(val *InlineResponse200124) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200124) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200124) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200124(val *InlineResponse200124) *NullableInlineResponse200124 {
	return &NullableInlineResponse200124{value: val, isSet: true}
}

func (v NullableInlineResponse200124) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200124) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

