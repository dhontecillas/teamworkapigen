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

// InlineResponse20012 struct for InlineResponse20012
type InlineResponse20012 struct {
	STATUS *string `json:"STATUS,omitempty"`
	Company *InlineResponse20011Companies `json:"company,omitempty"`
}

// NewInlineResponse20012 instantiates a new InlineResponse20012 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20012() *InlineResponse20012 {
	this := InlineResponse20012{}
	return &this
}

// NewInlineResponse20012WithDefaults instantiates a new InlineResponse20012 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20012WithDefaults() *InlineResponse20012 {
	this := InlineResponse20012{}
	return &this
}

// GetSTATUS returns the STATUS field value if set, zero value otherwise.
func (o *InlineResponse20012) GetSTATUS() string {
	if o == nil || o.STATUS == nil {
		var ret string
		return ret
	}
	return *o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20012) GetSTATUSOk() (*string, bool) {
	if o == nil || o.STATUS == nil {
		return nil, false
	}
	return o.STATUS, true
}

// HasSTATUS returns a boolean if a field has been set.
func (o *InlineResponse20012) HasSTATUS() bool {
	if o != nil && o.STATUS != nil {
		return true
	}

	return false
}

// SetSTATUS gets a reference to the given string and assigns it to the STATUS field.
func (o *InlineResponse20012) SetSTATUS(v string) {
	o.STATUS = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *InlineResponse20012) GetCompany() InlineResponse20011Companies {
	if o == nil || o.Company == nil {
		var ret InlineResponse20011Companies
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20012) GetCompanyOk() (*InlineResponse20011Companies, bool) {
	if o == nil || o.Company == nil {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *InlineResponse20012) HasCompany() bool {
	if o != nil && o.Company != nil {
		return true
	}

	return false
}

// SetCompany gets a reference to the given InlineResponse20011Companies and assigns it to the Company field.
func (o *InlineResponse20012) SetCompany(v InlineResponse20011Companies) {
	o.Company = &v
}

func (o InlineResponse20012) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.STATUS != nil {
		toSerialize["STATUS"] = o.STATUS
	}
	if o.Company != nil {
		toSerialize["company"] = o.Company
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20012 struct {
	value *InlineResponse20012
	isSet bool
}

func (v NullableInlineResponse20012) Get() *InlineResponse20012 {
	return v.value
}

func (v *NullableInlineResponse20012) Set(val *InlineResponse20012) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20012) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20012) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20012(val *InlineResponse20012) *NullableInlineResponse20012 {
	return &NullableInlineResponse20012{value: val, isSet: true}
}

func (v NullableInlineResponse20012) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20012) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

