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

// InlineResponse20020Category struct for InlineResponse20020Category
type InlineResponse20020Category struct {
	Count *string `json:"count,omitempty"`
	ElementsCount *string `json:"elements_count,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	ParentId *string `json:"parent-id,omitempty"`
	ProjectId *string `json:"project-id,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewInlineResponse20020Category instantiates a new InlineResponse20020Category object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20020Category() *InlineResponse20020Category {
	this := InlineResponse20020Category{}
	return &this
}

// NewInlineResponse20020CategoryWithDefaults instantiates a new InlineResponse20020Category object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20020CategoryWithDefaults() *InlineResponse20020Category {
	this := InlineResponse20020Category{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *InlineResponse20020Category) GetCount() string {
	if o == nil || o.Count == nil {
		var ret string
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20020Category) GetCountOk() (*string, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *InlineResponse20020Category) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given string and assigns it to the Count field.
func (o *InlineResponse20020Category) SetCount(v string) {
	o.Count = &v
}

// GetElementsCount returns the ElementsCount field value if set, zero value otherwise.
func (o *InlineResponse20020Category) GetElementsCount() string {
	if o == nil || o.ElementsCount == nil {
		var ret string
		return ret
	}
	return *o.ElementsCount
}

// GetElementsCountOk returns a tuple with the ElementsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20020Category) GetElementsCountOk() (*string, bool) {
	if o == nil || o.ElementsCount == nil {
		return nil, false
	}
	return o.ElementsCount, true
}

// HasElementsCount returns a boolean if a field has been set.
func (o *InlineResponse20020Category) HasElementsCount() bool {
	if o != nil && o.ElementsCount != nil {
		return true
	}

	return false
}

// SetElementsCount gets a reference to the given string and assigns it to the ElementsCount field.
func (o *InlineResponse20020Category) SetElementsCount(v string) {
	o.ElementsCount = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20020Category) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20020Category) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20020Category) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20020Category) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20020Category) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20020Category) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20020Category) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20020Category) SetName(v string) {
	o.Name = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *InlineResponse20020Category) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20020Category) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *InlineResponse20020Category) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *InlineResponse20020Category) SetParentId(v string) {
	o.ParentId = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *InlineResponse20020Category) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20020Category) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *InlineResponse20020Category) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *InlineResponse20020Category) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse20020Category) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20020Category) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse20020Category) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse20020Category) SetType(v string) {
	o.Type = &v
}

func (o InlineResponse20020Category) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.ElementsCount != nil {
		toSerialize["elements_count"] = o.ElementsCount
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ParentId != nil {
		toSerialize["parent-id"] = o.ParentId
	}
	if o.ProjectId != nil {
		toSerialize["project-id"] = o.ProjectId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20020Category struct {
	value *InlineResponse20020Category
	isSet bool
}

func (v NullableInlineResponse20020Category) Get() *InlineResponse20020Category {
	return v.value
}

func (v *NullableInlineResponse20020Category) Set(val *InlineResponse20020Category) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20020Category) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20020Category) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20020Category(val *InlineResponse20020Category) *NullableInlineResponse20020Category {
	return &NullableInlineResponse20020Category{value: val, isSet: true}
}

func (v NullableInlineResponse20020Category) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20020Category) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

