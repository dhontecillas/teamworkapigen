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

// FormForm Form contains information of a form to be created or updated.
type FormForm struct {
	ConfirmationMessage *string `json:"confirmationMessage,omitempty"`
	Content *FormContent `json:"content,omitempty"`
	HostObject *FormHostObject `json:"hostObject,omitempty"`
	PromptAdditionalSubmissions *bool `json:"promptAdditionalSubmissions,omitempty"`
	TaskTitleFieldId *string `json:"taskTitleFieldId,omitempty"`
}

// NewFormForm instantiates a new FormForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormForm() *FormForm {
	this := FormForm{}
	return &this
}

// NewFormFormWithDefaults instantiates a new FormForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormFormWithDefaults() *FormForm {
	this := FormForm{}
	return &this
}

// GetConfirmationMessage returns the ConfirmationMessage field value if set, zero value otherwise.
func (o *FormForm) GetConfirmationMessage() string {
	if o == nil || o.ConfirmationMessage == nil {
		var ret string
		return ret
	}
	return *o.ConfirmationMessage
}

// GetConfirmationMessageOk returns a tuple with the ConfirmationMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormForm) GetConfirmationMessageOk() (*string, bool) {
	if o == nil || o.ConfirmationMessage == nil {
		return nil, false
	}
	return o.ConfirmationMessage, true
}

// HasConfirmationMessage returns a boolean if a field has been set.
func (o *FormForm) HasConfirmationMessage() bool {
	if o != nil && o.ConfirmationMessage != nil {
		return true
	}

	return false
}

// SetConfirmationMessage gets a reference to the given string and assigns it to the ConfirmationMessage field.
func (o *FormForm) SetConfirmationMessage(v string) {
	o.ConfirmationMessage = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *FormForm) GetContent() FormContent {
	if o == nil || o.Content == nil {
		var ret FormContent
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormForm) GetContentOk() (*FormContent, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *FormForm) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given FormContent and assigns it to the Content field.
func (o *FormForm) SetContent(v FormContent) {
	o.Content = &v
}

// GetHostObject returns the HostObject field value if set, zero value otherwise.
func (o *FormForm) GetHostObject() FormHostObject {
	if o == nil || o.HostObject == nil {
		var ret FormHostObject
		return ret
	}
	return *o.HostObject
}

// GetHostObjectOk returns a tuple with the HostObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormForm) GetHostObjectOk() (*FormHostObject, bool) {
	if o == nil || o.HostObject == nil {
		return nil, false
	}
	return o.HostObject, true
}

// HasHostObject returns a boolean if a field has been set.
func (o *FormForm) HasHostObject() bool {
	if o != nil && o.HostObject != nil {
		return true
	}

	return false
}

// SetHostObject gets a reference to the given FormHostObject and assigns it to the HostObject field.
func (o *FormForm) SetHostObject(v FormHostObject) {
	o.HostObject = &v
}

// GetPromptAdditionalSubmissions returns the PromptAdditionalSubmissions field value if set, zero value otherwise.
func (o *FormForm) GetPromptAdditionalSubmissions() bool {
	if o == nil || o.PromptAdditionalSubmissions == nil {
		var ret bool
		return ret
	}
	return *o.PromptAdditionalSubmissions
}

// GetPromptAdditionalSubmissionsOk returns a tuple with the PromptAdditionalSubmissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormForm) GetPromptAdditionalSubmissionsOk() (*bool, bool) {
	if o == nil || o.PromptAdditionalSubmissions == nil {
		return nil, false
	}
	return o.PromptAdditionalSubmissions, true
}

// HasPromptAdditionalSubmissions returns a boolean if a field has been set.
func (o *FormForm) HasPromptAdditionalSubmissions() bool {
	if o != nil && o.PromptAdditionalSubmissions != nil {
		return true
	}

	return false
}

// SetPromptAdditionalSubmissions gets a reference to the given bool and assigns it to the PromptAdditionalSubmissions field.
func (o *FormForm) SetPromptAdditionalSubmissions(v bool) {
	o.PromptAdditionalSubmissions = &v
}

// GetTaskTitleFieldId returns the TaskTitleFieldId field value if set, zero value otherwise.
func (o *FormForm) GetTaskTitleFieldId() string {
	if o == nil || o.TaskTitleFieldId == nil {
		var ret string
		return ret
	}
	return *o.TaskTitleFieldId
}

// GetTaskTitleFieldIdOk returns a tuple with the TaskTitleFieldId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormForm) GetTaskTitleFieldIdOk() (*string, bool) {
	if o == nil || o.TaskTitleFieldId == nil {
		return nil, false
	}
	return o.TaskTitleFieldId, true
}

// HasTaskTitleFieldId returns a boolean if a field has been set.
func (o *FormForm) HasTaskTitleFieldId() bool {
	if o != nil && o.TaskTitleFieldId != nil {
		return true
	}

	return false
}

// SetTaskTitleFieldId gets a reference to the given string and assigns it to the TaskTitleFieldId field.
func (o *FormForm) SetTaskTitleFieldId(v string) {
	o.TaskTitleFieldId = &v
}

func (o FormForm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConfirmationMessage != nil {
		toSerialize["confirmationMessage"] = o.ConfirmationMessage
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.HostObject != nil {
		toSerialize["hostObject"] = o.HostObject
	}
	if o.PromptAdditionalSubmissions != nil {
		toSerialize["promptAdditionalSubmissions"] = o.PromptAdditionalSubmissions
	}
	if o.TaskTitleFieldId != nil {
		toSerialize["taskTitleFieldId"] = o.TaskTitleFieldId
	}
	return json.Marshal(toSerialize)
}

type NullableFormForm struct {
	value *FormForm
	isSet bool
}

func (v NullableFormForm) Get() *FormForm {
	return v.value
}

func (v *NullableFormForm) Set(val *FormForm) {
	v.value = val
	v.isSet = true
}

func (v NullableFormForm) IsSet() bool {
	return v.isSet
}

func (v *NullableFormForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormForm(val *FormForm) *NullableFormForm {
	return &NullableFormForm{value: val, isSet: true}
}

func (v NullableFormForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


