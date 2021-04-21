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

// InlineResponse2011 struct for InlineResponse2011
type InlineResponse2011 struct {
	PendingFile *InlineResponse2011PendingFile `json:"pendingFile,omitempty"`
}

// NewInlineResponse2011 instantiates a new InlineResponse2011 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2011() *InlineResponse2011 {
	this := InlineResponse2011{}
	return &this
}

// NewInlineResponse2011WithDefaults instantiates a new InlineResponse2011 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2011WithDefaults() *InlineResponse2011 {
	this := InlineResponse2011{}
	return &this
}

// GetPendingFile returns the PendingFile field value if set, zero value otherwise.
func (o *InlineResponse2011) GetPendingFile() InlineResponse2011PendingFile {
	if o == nil || o.PendingFile == nil {
		var ret InlineResponse2011PendingFile
		return ret
	}
	return *o.PendingFile
}

// GetPendingFileOk returns a tuple with the PendingFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2011) GetPendingFileOk() (*InlineResponse2011PendingFile, bool) {
	if o == nil || o.PendingFile == nil {
		return nil, false
	}
	return o.PendingFile, true
}

// HasPendingFile returns a boolean if a field has been set.
func (o *InlineResponse2011) HasPendingFile() bool {
	if o != nil && o.PendingFile != nil {
		return true
	}

	return false
}

// SetPendingFile gets a reference to the given InlineResponse2011PendingFile and assigns it to the PendingFile field.
func (o *InlineResponse2011) SetPendingFile(v InlineResponse2011PendingFile) {
	o.PendingFile = &v
}

func (o InlineResponse2011) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PendingFile != nil {
		toSerialize["pendingFile"] = o.PendingFile
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2011 struct {
	value *InlineResponse2011
	isSet bool
}

func (v NullableInlineResponse2011) Get() *InlineResponse2011 {
	return v.value
}

func (v *NullableInlineResponse2011) Set(val *InlineResponse2011) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2011) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2011) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2011(val *InlineResponse2011) *NullableInlineResponse2011 {
	return &NullableInlineResponse2011{value: val, isSet: true}
}

func (v NullableInlineResponse2011) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2011) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


