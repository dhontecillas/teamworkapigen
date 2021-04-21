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

// ViewContent Content contains all the information returned from a form token.
type ViewContent struct {
	Banner *ViewBanner `json:"banner,omitempty"`
	Definition *string `json:"definition,omitempty"`
	Description *string `json:"description,omitempty"`
	Name *string `json:"name,omitempty"`
	State *string `json:"state,omitempty"`
	Version *int32 `json:"version,omitempty"`
}

// NewViewContent instantiates a new ViewContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewContent() *ViewContent {
	this := ViewContent{}
	return &this
}

// NewViewContentWithDefaults instantiates a new ViewContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewContentWithDefaults() *ViewContent {
	this := ViewContent{}
	return &this
}

// GetBanner returns the Banner field value if set, zero value otherwise.
func (o *ViewContent) GetBanner() ViewBanner {
	if o == nil || o.Banner == nil {
		var ret ViewBanner
		return ret
	}
	return *o.Banner
}

// GetBannerOk returns a tuple with the Banner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewContent) GetBannerOk() (*ViewBanner, bool) {
	if o == nil || o.Banner == nil {
		return nil, false
	}
	return o.Banner, true
}

// HasBanner returns a boolean if a field has been set.
func (o *ViewContent) HasBanner() bool {
	if o != nil && o.Banner != nil {
		return true
	}

	return false
}

// SetBanner gets a reference to the given ViewBanner and assigns it to the Banner field.
func (o *ViewContent) SetBanner(v ViewBanner) {
	o.Banner = &v
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *ViewContent) GetDefinition() string {
	if o == nil || o.Definition == nil {
		var ret string
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewContent) GetDefinitionOk() (*string, bool) {
	if o == nil || o.Definition == nil {
		return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *ViewContent) HasDefinition() bool {
	if o != nil && o.Definition != nil {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given string and assigns it to the Definition field.
func (o *ViewContent) SetDefinition(v string) {
	o.Definition = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ViewContent) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewContent) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ViewContent) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ViewContent) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ViewContent) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewContent) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ViewContent) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ViewContent) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ViewContent) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewContent) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ViewContent) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ViewContent) SetState(v string) {
	o.State = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ViewContent) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewContent) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ViewContent) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *ViewContent) SetVersion(v int32) {
	o.Version = &v
}

func (o ViewContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Banner != nil {
		toSerialize["banner"] = o.Banner
	}
	if o.Definition != nil {
		toSerialize["definition"] = o.Definition
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableViewContent struct {
	value *ViewContent
	isSet bool
}

func (v NullableViewContent) Get() *ViewContent {
	return v.value
}

func (v *NullableViewContent) Set(val *ViewContent) {
	v.value = val
	v.isSet = true
}

func (v NullableViewContent) IsSet() bool {
	return v.isSet
}

func (v *NullableViewContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewContent(val *ViewContent) *NullableViewContent {
	return &NullableViewContent{value: val, isSet: true}
}

func (v NullableViewContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

