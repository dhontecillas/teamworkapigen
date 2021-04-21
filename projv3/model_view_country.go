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

// ViewCountry Country represents all the information returned from a country.
type ViewCountry struct {
	Code *string `json:"code,omitempty"`
	Country *string `json:"country,omitempty"`
	Eu *bool `json:"eu,omitempty"`
	PhoneCode *string `json:"phoneCode,omitempty"`
	VatName *string `json:"vatName,omitempty"`
	VatPercent *int32 `json:"vatPercent,omitempty"`
}

// NewViewCountry instantiates a new ViewCountry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewCountry() *ViewCountry {
	this := ViewCountry{}
	return &this
}

// NewViewCountryWithDefaults instantiates a new ViewCountry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewCountryWithDefaults() *ViewCountry {
	this := ViewCountry{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ViewCountry) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewCountry) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ViewCountry) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ViewCountry) SetCode(v string) {
	o.Code = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *ViewCountry) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewCountry) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *ViewCountry) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *ViewCountry) SetCountry(v string) {
	o.Country = &v
}

// GetEu returns the Eu field value if set, zero value otherwise.
func (o *ViewCountry) GetEu() bool {
	if o == nil || o.Eu == nil {
		var ret bool
		return ret
	}
	return *o.Eu
}

// GetEuOk returns a tuple with the Eu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewCountry) GetEuOk() (*bool, bool) {
	if o == nil || o.Eu == nil {
		return nil, false
	}
	return o.Eu, true
}

// HasEu returns a boolean if a field has been set.
func (o *ViewCountry) HasEu() bool {
	if o != nil && o.Eu != nil {
		return true
	}

	return false
}

// SetEu gets a reference to the given bool and assigns it to the Eu field.
func (o *ViewCountry) SetEu(v bool) {
	o.Eu = &v
}

// GetPhoneCode returns the PhoneCode field value if set, zero value otherwise.
func (o *ViewCountry) GetPhoneCode() string {
	if o == nil || o.PhoneCode == nil {
		var ret string
		return ret
	}
	return *o.PhoneCode
}

// GetPhoneCodeOk returns a tuple with the PhoneCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewCountry) GetPhoneCodeOk() (*string, bool) {
	if o == nil || o.PhoneCode == nil {
		return nil, false
	}
	return o.PhoneCode, true
}

// HasPhoneCode returns a boolean if a field has been set.
func (o *ViewCountry) HasPhoneCode() bool {
	if o != nil && o.PhoneCode != nil {
		return true
	}

	return false
}

// SetPhoneCode gets a reference to the given string and assigns it to the PhoneCode field.
func (o *ViewCountry) SetPhoneCode(v string) {
	o.PhoneCode = &v
}

// GetVatName returns the VatName field value if set, zero value otherwise.
func (o *ViewCountry) GetVatName() string {
	if o == nil || o.VatName == nil {
		var ret string
		return ret
	}
	return *o.VatName
}

// GetVatNameOk returns a tuple with the VatName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewCountry) GetVatNameOk() (*string, bool) {
	if o == nil || o.VatName == nil {
		return nil, false
	}
	return o.VatName, true
}

// HasVatName returns a boolean if a field has been set.
func (o *ViewCountry) HasVatName() bool {
	if o != nil && o.VatName != nil {
		return true
	}

	return false
}

// SetVatName gets a reference to the given string and assigns it to the VatName field.
func (o *ViewCountry) SetVatName(v string) {
	o.VatName = &v
}

// GetVatPercent returns the VatPercent field value if set, zero value otherwise.
func (o *ViewCountry) GetVatPercent() int32 {
	if o == nil || o.VatPercent == nil {
		var ret int32
		return ret
	}
	return *o.VatPercent
}

// GetVatPercentOk returns a tuple with the VatPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewCountry) GetVatPercentOk() (*int32, bool) {
	if o == nil || o.VatPercent == nil {
		return nil, false
	}
	return o.VatPercent, true
}

// HasVatPercent returns a boolean if a field has been set.
func (o *ViewCountry) HasVatPercent() bool {
	if o != nil && o.VatPercent != nil {
		return true
	}

	return false
}

// SetVatPercent gets a reference to the given int32 and assigns it to the VatPercent field.
func (o *ViewCountry) SetVatPercent(v int32) {
	o.VatPercent = &v
}

func (o ViewCountry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.Eu != nil {
		toSerialize["eu"] = o.Eu
	}
	if o.PhoneCode != nil {
		toSerialize["phoneCode"] = o.PhoneCode
	}
	if o.VatName != nil {
		toSerialize["vatName"] = o.VatName
	}
	if o.VatPercent != nil {
		toSerialize["vatPercent"] = o.VatPercent
	}
	return json.Marshal(toSerialize)
}

type NullableViewCountry struct {
	value *ViewCountry
	isSet bool
}

func (v NullableViewCountry) Get() *ViewCountry {
	return v.value
}

func (v *NullableViewCountry) Set(val *ViewCountry) {
	v.value = val
	v.isSet = true
}

func (v NullableViewCountry) IsSet() bool {
	return v.isSet
}

func (v *NullableViewCountry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewCountry(val *ViewCountry) *NullableViewCountry {
	return &NullableViewCountry{value: val, isSet: true}
}

func (v NullableViewCountry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewCountry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

