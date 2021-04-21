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

// InlineObject107 struct for InlineObject107
type InlineObject107 struct {
	STATUS *string `json:"STATUS,omitempty"`
	Webhook *WebhooksJsonWebhook `json:"webhook,omitempty"`
}

// NewInlineObject107 instantiates a new InlineObject107 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject107() *InlineObject107 {
	this := InlineObject107{}
	return &this
}

// NewInlineObject107WithDefaults instantiates a new InlineObject107 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject107WithDefaults() *InlineObject107 {
	this := InlineObject107{}
	return &this
}

// GetSTATUS returns the STATUS field value if set, zero value otherwise.
func (o *InlineObject107) GetSTATUS() string {
	if o == nil || o.STATUS == nil {
		var ret string
		return ret
	}
	return *o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject107) GetSTATUSOk() (*string, bool) {
	if o == nil || o.STATUS == nil {
		return nil, false
	}
	return o.STATUS, true
}

// HasSTATUS returns a boolean if a field has been set.
func (o *InlineObject107) HasSTATUS() bool {
	if o != nil && o.STATUS != nil {
		return true
	}

	return false
}

// SetSTATUS gets a reference to the given string and assigns it to the STATUS field.
func (o *InlineObject107) SetSTATUS(v string) {
	o.STATUS = &v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise.
func (o *InlineObject107) GetWebhook() WebhooksJsonWebhook {
	if o == nil || o.Webhook == nil {
		var ret WebhooksJsonWebhook
		return ret
	}
	return *o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject107) GetWebhookOk() (*WebhooksJsonWebhook, bool) {
	if o == nil || o.Webhook == nil {
		return nil, false
	}
	return o.Webhook, true
}

// HasWebhook returns a boolean if a field has been set.
func (o *InlineObject107) HasWebhook() bool {
	if o != nil && o.Webhook != nil {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given WebhooksJsonWebhook and assigns it to the Webhook field.
func (o *InlineObject107) SetWebhook(v WebhooksJsonWebhook) {
	o.Webhook = &v
}

func (o InlineObject107) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.STATUS != nil {
		toSerialize["STATUS"] = o.STATUS
	}
	if o.Webhook != nil {
		toSerialize["webhook"] = o.Webhook
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject107 struct {
	value *InlineObject107
	isSet bool
}

func (v NullableInlineObject107) Get() *InlineObject107 {
	return v.value
}

func (v *NullableInlineObject107) Set(val *InlineObject107) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject107) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject107) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject107(val *InlineObject107) *NullableInlineObject107 {
	return &NullableInlineObject107{value: val, isSet: true}
}

func (v NullableInlineObject107) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject107) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

