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

// InlineResponse200116 struct for InlineResponse200116
type InlineResponse200116 struct {
	STATUS *string `json:"STATUS,omitempty"`
	MessageReplies *[]InlineResponse200116MessageReplies `json:"messageReplies,omitempty"`
}

// NewInlineResponse200116 instantiates a new InlineResponse200116 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200116() *InlineResponse200116 {
	this := InlineResponse200116{}
	return &this
}

// NewInlineResponse200116WithDefaults instantiates a new InlineResponse200116 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200116WithDefaults() *InlineResponse200116 {
	this := InlineResponse200116{}
	return &this
}

// GetSTATUS returns the STATUS field value if set, zero value otherwise.
func (o *InlineResponse200116) GetSTATUS() string {
	if o == nil || o.STATUS == nil {
		var ret string
		return ret
	}
	return *o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200116) GetSTATUSOk() (*string, bool) {
	if o == nil || o.STATUS == nil {
		return nil, false
	}
	return o.STATUS, true
}

// HasSTATUS returns a boolean if a field has been set.
func (o *InlineResponse200116) HasSTATUS() bool {
	if o != nil && o.STATUS != nil {
		return true
	}

	return false
}

// SetSTATUS gets a reference to the given string and assigns it to the STATUS field.
func (o *InlineResponse200116) SetSTATUS(v string) {
	o.STATUS = &v
}

// GetMessageReplies returns the MessageReplies field value if set, zero value otherwise.
func (o *InlineResponse200116) GetMessageReplies() []InlineResponse200116MessageReplies {
	if o == nil || o.MessageReplies == nil {
		var ret []InlineResponse200116MessageReplies
		return ret
	}
	return *o.MessageReplies
}

// GetMessageRepliesOk returns a tuple with the MessageReplies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200116) GetMessageRepliesOk() (*[]InlineResponse200116MessageReplies, bool) {
	if o == nil || o.MessageReplies == nil {
		return nil, false
	}
	return o.MessageReplies, true
}

// HasMessageReplies returns a boolean if a field has been set.
func (o *InlineResponse200116) HasMessageReplies() bool {
	if o != nil && o.MessageReplies != nil {
		return true
	}

	return false
}

// SetMessageReplies gets a reference to the given []InlineResponse200116MessageReplies and assigns it to the MessageReplies field.
func (o *InlineResponse200116) SetMessageReplies(v []InlineResponse200116MessageReplies) {
	o.MessageReplies = &v
}

func (o InlineResponse200116) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.STATUS != nil {
		toSerialize["STATUS"] = o.STATUS
	}
	if o.MessageReplies != nil {
		toSerialize["messageReplies"] = o.MessageReplies
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200116 struct {
	value *InlineResponse200116
	isSet bool
}

func (v NullableInlineResponse200116) Get() *InlineResponse200116 {
	return v.value
}

func (v *NullableInlineResponse200116) Set(val *InlineResponse200116) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200116) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200116) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200116(val *InlineResponse200116) *NullableInlineResponse200116 {
	return &NullableInlineResponse200116{value: val, isSet: true}
}

func (v NullableInlineResponse200116) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200116) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

