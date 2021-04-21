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

// InvoicesIdLineitemsJsonLineitemsAdd struct for InvoicesIdLineitemsJsonLineitemsAdd
type InvoicesIdLineitemsJsonLineitemsAdd struct {
	Timelogs *string `json:"timelogs,omitempty"`
}

// NewInvoicesIdLineitemsJsonLineitemsAdd instantiates a new InvoicesIdLineitemsJsonLineitemsAdd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoicesIdLineitemsJsonLineitemsAdd() *InvoicesIdLineitemsJsonLineitemsAdd {
	this := InvoicesIdLineitemsJsonLineitemsAdd{}
	return &this
}

// NewInvoicesIdLineitemsJsonLineitemsAddWithDefaults instantiates a new InvoicesIdLineitemsJsonLineitemsAdd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoicesIdLineitemsJsonLineitemsAddWithDefaults() *InvoicesIdLineitemsJsonLineitemsAdd {
	this := InvoicesIdLineitemsJsonLineitemsAdd{}
	return &this
}

// GetTimelogs returns the Timelogs field value if set, zero value otherwise.
func (o *InvoicesIdLineitemsJsonLineitemsAdd) GetTimelogs() string {
	if o == nil || o.Timelogs == nil {
		var ret string
		return ret
	}
	return *o.Timelogs
}

// GetTimelogsOk returns a tuple with the Timelogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoicesIdLineitemsJsonLineitemsAdd) GetTimelogsOk() (*string, bool) {
	if o == nil || o.Timelogs == nil {
		return nil, false
	}
	return o.Timelogs, true
}

// HasTimelogs returns a boolean if a field has been set.
func (o *InvoicesIdLineitemsJsonLineitemsAdd) HasTimelogs() bool {
	if o != nil && o.Timelogs != nil {
		return true
	}

	return false
}

// SetTimelogs gets a reference to the given string and assigns it to the Timelogs field.
func (o *InvoicesIdLineitemsJsonLineitemsAdd) SetTimelogs(v string) {
	o.Timelogs = &v
}

func (o InvoicesIdLineitemsJsonLineitemsAdd) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Timelogs != nil {
		toSerialize["timelogs"] = o.Timelogs
	}
	return json.Marshal(toSerialize)
}

type NullableInvoicesIdLineitemsJsonLineitemsAdd struct {
	value *InvoicesIdLineitemsJsonLineitemsAdd
	isSet bool
}

func (v NullableInvoicesIdLineitemsJsonLineitemsAdd) Get() *InvoicesIdLineitemsJsonLineitemsAdd {
	return v.value
}

func (v *NullableInvoicesIdLineitemsJsonLineitemsAdd) Set(val *InvoicesIdLineitemsJsonLineitemsAdd) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoicesIdLineitemsJsonLineitemsAdd) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoicesIdLineitemsJsonLineitemsAdd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoicesIdLineitemsJsonLineitemsAdd(val *InvoicesIdLineitemsJsonLineitemsAdd) *NullableInvoicesIdLineitemsJsonLineitemsAdd {
	return &NullableInvoicesIdLineitemsJsonLineitemsAdd{value: val, isSet: true}
}

func (v NullableInvoicesIdLineitemsJsonLineitemsAdd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoicesIdLineitemsJsonLineitemsAdd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

