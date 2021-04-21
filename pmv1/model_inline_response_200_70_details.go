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

// InlineResponse20070Details struct for InlineResponse20070Details
type InlineResponse20070Details struct {
	Added *string `json:"added,omitempty"`
	Failed *string `json:"failed,omitempty"`
	Removed *string `json:"removed,omitempty"`
}

// NewInlineResponse20070Details instantiates a new InlineResponse20070Details object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20070Details() *InlineResponse20070Details {
	this := InlineResponse20070Details{}
	return &this
}

// NewInlineResponse20070DetailsWithDefaults instantiates a new InlineResponse20070Details object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20070DetailsWithDefaults() *InlineResponse20070Details {
	this := InlineResponse20070Details{}
	return &this
}

// GetAdded returns the Added field value if set, zero value otherwise.
func (o *InlineResponse20070Details) GetAdded() string {
	if o == nil || o.Added == nil {
		var ret string
		return ret
	}
	return *o.Added
}

// GetAddedOk returns a tuple with the Added field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20070Details) GetAddedOk() (*string, bool) {
	if o == nil || o.Added == nil {
		return nil, false
	}
	return o.Added, true
}

// HasAdded returns a boolean if a field has been set.
func (o *InlineResponse20070Details) HasAdded() bool {
	if o != nil && o.Added != nil {
		return true
	}

	return false
}

// SetAdded gets a reference to the given string and assigns it to the Added field.
func (o *InlineResponse20070Details) SetAdded(v string) {
	o.Added = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *InlineResponse20070Details) GetFailed() string {
	if o == nil || o.Failed == nil {
		var ret string
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20070Details) GetFailedOk() (*string, bool) {
	if o == nil || o.Failed == nil {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *InlineResponse20070Details) HasFailed() bool {
	if o != nil && o.Failed != nil {
		return true
	}

	return false
}

// SetFailed gets a reference to the given string and assigns it to the Failed field.
func (o *InlineResponse20070Details) SetFailed(v string) {
	o.Failed = &v
}

// GetRemoved returns the Removed field value if set, zero value otherwise.
func (o *InlineResponse20070Details) GetRemoved() string {
	if o == nil || o.Removed == nil {
		var ret string
		return ret
	}
	return *o.Removed
}

// GetRemovedOk returns a tuple with the Removed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20070Details) GetRemovedOk() (*string, bool) {
	if o == nil || o.Removed == nil {
		return nil, false
	}
	return o.Removed, true
}

// HasRemoved returns a boolean if a field has been set.
func (o *InlineResponse20070Details) HasRemoved() bool {
	if o != nil && o.Removed != nil {
		return true
	}

	return false
}

// SetRemoved gets a reference to the given string and assigns it to the Removed field.
func (o *InlineResponse20070Details) SetRemoved(v string) {
	o.Removed = &v
}

func (o InlineResponse20070Details) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Added != nil {
		toSerialize["added"] = o.Added
	}
	if o.Failed != nil {
		toSerialize["failed"] = o.Failed
	}
	if o.Removed != nil {
		toSerialize["removed"] = o.Removed
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20070Details struct {
	value *InlineResponse20070Details
	isSet bool
}

func (v NullableInlineResponse20070Details) Get() *InlineResponse20070Details {
	return v.value
}

func (v *NullableInlineResponse20070Details) Set(val *InlineResponse20070Details) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20070Details) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20070Details) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20070Details(val *InlineResponse20070Details) *NullableInlineResponse20070Details {
	return &NullableInlineResponse20070Details{value: val, isSet: true}
}

func (v NullableInlineResponse20070Details) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20070Details) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

