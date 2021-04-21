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

// InlineResponse20079 struct for InlineResponse20079
type InlineResponse20079 struct {
	STATUS *string `json:"STATUS,omitempty"`
	TimeEntries *[]InlineResponse20079TimeEntries `json:"time-entries,omitempty"`
}

// NewInlineResponse20079 instantiates a new InlineResponse20079 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20079() *InlineResponse20079 {
	this := InlineResponse20079{}
	return &this
}

// NewInlineResponse20079WithDefaults instantiates a new InlineResponse20079 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20079WithDefaults() *InlineResponse20079 {
	this := InlineResponse20079{}
	return &this
}

// GetSTATUS returns the STATUS field value if set, zero value otherwise.
func (o *InlineResponse20079) GetSTATUS() string {
	if o == nil || o.STATUS == nil {
		var ret string
		return ret
	}
	return *o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetSTATUSOk() (*string, bool) {
	if o == nil || o.STATUS == nil {
		return nil, false
	}
	return o.STATUS, true
}

// HasSTATUS returns a boolean if a field has been set.
func (o *InlineResponse20079) HasSTATUS() bool {
	if o != nil && o.STATUS != nil {
		return true
	}

	return false
}

// SetSTATUS gets a reference to the given string and assigns it to the STATUS field.
func (o *InlineResponse20079) SetSTATUS(v string) {
	o.STATUS = &v
}

// GetTimeEntries returns the TimeEntries field value if set, zero value otherwise.
func (o *InlineResponse20079) GetTimeEntries() []InlineResponse20079TimeEntries {
	if o == nil || o.TimeEntries == nil {
		var ret []InlineResponse20079TimeEntries
		return ret
	}
	return *o.TimeEntries
}

// GetTimeEntriesOk returns a tuple with the TimeEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetTimeEntriesOk() (*[]InlineResponse20079TimeEntries, bool) {
	if o == nil || o.TimeEntries == nil {
		return nil, false
	}
	return o.TimeEntries, true
}

// HasTimeEntries returns a boolean if a field has been set.
func (o *InlineResponse20079) HasTimeEntries() bool {
	if o != nil && o.TimeEntries != nil {
		return true
	}

	return false
}

// SetTimeEntries gets a reference to the given []InlineResponse20079TimeEntries and assigns it to the TimeEntries field.
func (o *InlineResponse20079) SetTimeEntries(v []InlineResponse20079TimeEntries) {
	o.TimeEntries = &v
}

func (o InlineResponse20079) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.STATUS != nil {
		toSerialize["STATUS"] = o.STATUS
	}
	if o.TimeEntries != nil {
		toSerialize["time-entries"] = o.TimeEntries
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20079 struct {
	value *InlineResponse20079
	isSet bool
}

func (v NullableInlineResponse20079) Get() *InlineResponse20079 {
	return v.value
}

func (v *NullableInlineResponse20079) Set(val *InlineResponse20079) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20079) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20079) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20079(val *InlineResponse20079) *NullableInlineResponse20079 {
	return &NullableInlineResponse20079{value: val, isSet: true}
}

func (v NullableInlineResponse20079) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20079) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

