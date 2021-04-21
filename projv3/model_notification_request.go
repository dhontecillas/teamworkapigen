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

// NotificationRequest Request contains information of a notification to be created or updated.
type NotificationRequest struct {
	Notification *NotificationProjectBudgetNotification `json:"notification,omitempty"`
}

// NewNotificationRequest instantiates a new NotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationRequest() *NotificationRequest {
	this := NotificationRequest{}
	return &this
}

// NewNotificationRequestWithDefaults instantiates a new NotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationRequestWithDefaults() *NotificationRequest {
	this := NotificationRequest{}
	return &this
}

// GetNotification returns the Notification field value if set, zero value otherwise.
func (o *NotificationRequest) GetNotification() NotificationProjectBudgetNotification {
	if o == nil || o.Notification == nil {
		var ret NotificationProjectBudgetNotification
		return ret
	}
	return *o.Notification
}

// GetNotificationOk returns a tuple with the Notification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRequest) GetNotificationOk() (*NotificationProjectBudgetNotification, bool) {
	if o == nil || o.Notification == nil {
		return nil, false
	}
	return o.Notification, true
}

// HasNotification returns a boolean if a field has been set.
func (o *NotificationRequest) HasNotification() bool {
	if o != nil && o.Notification != nil {
		return true
	}

	return false
}

// SetNotification gets a reference to the given NotificationProjectBudgetNotification and assigns it to the Notification field.
func (o *NotificationRequest) SetNotification(v NotificationProjectBudgetNotification) {
	o.Notification = &v
}

func (o NotificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Notification != nil {
		toSerialize["notification"] = o.Notification
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationRequest struct {
	value *NotificationRequest
	isSet bool
}

func (v NullableNotificationRequest) Get() *NotificationRequest {
	return v.value
}

func (v *NullableNotificationRequest) Set(val *NotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationRequest(val *NotificationRequest) *NullableNotificationRequest {
	return &NullableNotificationRequest{value: val, isSet: true}
}

func (v NullableNotificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

