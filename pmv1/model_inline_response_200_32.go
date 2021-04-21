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

// InlineResponse20032 struct for InlineResponse20032
type InlineResponse20032 struct {
	STATUS *string `json:"STATUS,omitempty"`
	MessageReplies *[]InlineResponse20032MessageReplies `json:"messageReplies,omitempty"`
}

// NewInlineResponse20032 instantiates a new InlineResponse20032 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20032() *InlineResponse20032 {
	this := InlineResponse20032{}
	return &this
}

// NewInlineResponse20032WithDefaults instantiates a new InlineResponse20032 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20032WithDefaults() *InlineResponse20032 {
	this := InlineResponse20032{}
	return &this
}

// GetSTATUS returns the STATUS field value if set, zero value otherwise.
func (o *InlineResponse20032) GetSTATUS() string {
	if o == nil || o.STATUS == nil {
		var ret string
		return ret
	}
	return *o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20032) GetSTATUSOk() (*string, bool) {
	if o == nil || o.STATUS == nil {
		return nil, false
	}
	return o.STATUS, true
}

// HasSTATUS returns a boolean if a field has been set.
func (o *InlineResponse20032) HasSTATUS() bool {
	if o != nil && o.STATUS != nil {
		return true
	}

	return false
}

// SetSTATUS gets a reference to the given string and assigns it to the STATUS field.
func (o *InlineResponse20032) SetSTATUS(v string) {
	o.STATUS = &v
}

// GetMessageReplies returns the MessageReplies field value if set, zero value otherwise.
func (o *InlineResponse20032) GetMessageReplies() []InlineResponse20032MessageReplies {
	if o == nil || o.MessageReplies == nil {
		var ret []InlineResponse20032MessageReplies
		return ret
	}
	return *o.MessageReplies
}

// GetMessageRepliesOk returns a tuple with the MessageReplies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20032) GetMessageRepliesOk() (*[]InlineResponse20032MessageReplies, bool) {
	if o == nil || o.MessageReplies == nil {
		return nil, false
	}
	return o.MessageReplies, true
}

// HasMessageReplies returns a boolean if a field has been set.
func (o *InlineResponse20032) HasMessageReplies() bool {
	if o != nil && o.MessageReplies != nil {
		return true
	}

	return false
}

// SetMessageReplies gets a reference to the given []InlineResponse20032MessageReplies and assigns it to the MessageReplies field.
func (o *InlineResponse20032) SetMessageReplies(v []InlineResponse20032MessageReplies) {
	o.MessageReplies = &v
}

func (o InlineResponse20032) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.STATUS != nil {
		toSerialize["STATUS"] = o.STATUS
	}
	if o.MessageReplies != nil {
		toSerialize["messageReplies"] = o.MessageReplies
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20032 struct {
	value *InlineResponse20032
	isSet bool
}

func (v NullableInlineResponse20032) Get() *InlineResponse20032 {
	return v.value
}

func (v *NullableInlineResponse20032) Set(val *InlineResponse20032) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20032) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20032) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20032(val *InlineResponse20032) *NullableInlineResponse20032 {
	return &NullableInlineResponse20032{value: val, isSet: true}
}

func (v NullableInlineResponse20032) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20032) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


