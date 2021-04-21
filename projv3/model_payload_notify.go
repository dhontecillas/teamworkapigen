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

// PayloadNotify Notify defines the access lists.
type PayloadNotify struct {
	Ids *PayloadUserGroups `json:"ids,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewPayloadNotify instantiates a new PayloadNotify object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayloadNotify() *PayloadNotify {
	this := PayloadNotify{}
	return &this
}

// NewPayloadNotifyWithDefaults instantiates a new PayloadNotify object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayloadNotifyWithDefaults() *PayloadNotify {
	this := PayloadNotify{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *PayloadNotify) GetIds() PayloadUserGroups {
	if o == nil || o.Ids == nil {
		var ret PayloadUserGroups
		return ret
	}
	return *o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayloadNotify) GetIdsOk() (*PayloadUserGroups, bool) {
	if o == nil || o.Ids == nil {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *PayloadNotify) HasIds() bool {
	if o != nil && o.Ids != nil {
		return true
	}

	return false
}

// SetIds gets a reference to the given PayloadUserGroups and assigns it to the Ids field.
func (o *PayloadNotify) SetIds(v PayloadUserGroups) {
	o.Ids = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PayloadNotify) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayloadNotify) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PayloadNotify) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PayloadNotify) SetType(v string) {
	o.Type = &v
}

func (o PayloadNotify) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullablePayloadNotify struct {
	value *PayloadNotify
	isSet bool
}

func (v NullablePayloadNotify) Get() *PayloadNotify {
	return v.value
}

func (v *NullablePayloadNotify) Set(val *PayloadNotify) {
	v.value = val
	v.isSet = true
}

func (v NullablePayloadNotify) IsSet() bool {
	return v.isSet
}

func (v *NullablePayloadNotify) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayloadNotify(val *PayloadNotify) *NullablePayloadNotify {
	return &NullablePayloadNotify{value: val, isSet: true}
}

func (v NullablePayloadNotify) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayloadNotify) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

