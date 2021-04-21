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

// InlineResponse20095Tasklists struct for InlineResponse20095Tasklists
type InlineResponse20095Tasklists struct {
	Complete *bool `json:"complete,omitempty"`
	Description *string `json:"description,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Position *string `json:"position,omitempty"`
	Private *bool `json:"private,omitempty"`
	UncompletedCount *string `json:"uncompleted-count,omitempty"`
}

// NewInlineResponse20095Tasklists instantiates a new InlineResponse20095Tasklists object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20095Tasklists() *InlineResponse20095Tasklists {
	this := InlineResponse20095Tasklists{}
	return &this
}

// NewInlineResponse20095TasklistsWithDefaults instantiates a new InlineResponse20095Tasklists object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20095TasklistsWithDefaults() *InlineResponse20095Tasklists {
	this := InlineResponse20095Tasklists{}
	return &this
}

// GetComplete returns the Complete field value if set, zero value otherwise.
func (o *InlineResponse20095Tasklists) GetComplete() bool {
	if o == nil || o.Complete == nil {
		var ret bool
		return ret
	}
	return *o.Complete
}

// GetCompleteOk returns a tuple with the Complete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095Tasklists) GetCompleteOk() (*bool, bool) {
	if o == nil || o.Complete == nil {
		return nil, false
	}
	return o.Complete, true
}

// HasComplete returns a boolean if a field has been set.
func (o *InlineResponse20095Tasklists) HasComplete() bool {
	if o != nil && o.Complete != nil {
		return true
	}

	return false
}

// SetComplete gets a reference to the given bool and assigns it to the Complete field.
func (o *InlineResponse20095Tasklists) SetComplete(v bool) {
	o.Complete = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineResponse20095Tasklists) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095Tasklists) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineResponse20095Tasklists) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineResponse20095Tasklists) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20095Tasklists) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095Tasklists) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20095Tasklists) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20095Tasklists) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20095Tasklists) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095Tasklists) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20095Tasklists) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20095Tasklists) SetName(v string) {
	o.Name = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *InlineResponse20095Tasklists) GetPosition() string {
	if o == nil || o.Position == nil {
		var ret string
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095Tasklists) GetPositionOk() (*string, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *InlineResponse20095Tasklists) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given string and assigns it to the Position field.
func (o *InlineResponse20095Tasklists) SetPosition(v string) {
	o.Position = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *InlineResponse20095Tasklists) GetPrivate() bool {
	if o == nil || o.Private == nil {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095Tasklists) GetPrivateOk() (*bool, bool) {
	if o == nil || o.Private == nil {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *InlineResponse20095Tasklists) HasPrivate() bool {
	if o != nil && o.Private != nil {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *InlineResponse20095Tasklists) SetPrivate(v bool) {
	o.Private = &v
}

// GetUncompletedCount returns the UncompletedCount field value if set, zero value otherwise.
func (o *InlineResponse20095Tasklists) GetUncompletedCount() string {
	if o == nil || o.UncompletedCount == nil {
		var ret string
		return ret
	}
	return *o.UncompletedCount
}

// GetUncompletedCountOk returns a tuple with the UncompletedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095Tasklists) GetUncompletedCountOk() (*string, bool) {
	if o == nil || o.UncompletedCount == nil {
		return nil, false
	}
	return o.UncompletedCount, true
}

// HasUncompletedCount returns a boolean if a field has been set.
func (o *InlineResponse20095Tasklists) HasUncompletedCount() bool {
	if o != nil && o.UncompletedCount != nil {
		return true
	}

	return false
}

// SetUncompletedCount gets a reference to the given string and assigns it to the UncompletedCount field.
func (o *InlineResponse20095Tasklists) SetUncompletedCount(v string) {
	o.UncompletedCount = &v
}

func (o InlineResponse20095Tasklists) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Complete != nil {
		toSerialize["complete"] = o.Complete
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}
	if o.Private != nil {
		toSerialize["private"] = o.Private
	}
	if o.UncompletedCount != nil {
		toSerialize["uncompleted-count"] = o.UncompletedCount
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20095Tasklists struct {
	value *InlineResponse20095Tasklists
	isSet bool
}

func (v NullableInlineResponse20095Tasklists) Get() *InlineResponse20095Tasklists {
	return v.value
}

func (v *NullableInlineResponse20095Tasklists) Set(val *InlineResponse20095Tasklists) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20095Tasklists) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20095Tasklists) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20095Tasklists(val *InlineResponse20095Tasklists) *NullableInlineResponse20095Tasklists {
	return &NullableInlineResponse20095Tasklists{value: val, isSet: true}
}

func (v NullableInlineResponse20095Tasklists) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20095Tasklists) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

