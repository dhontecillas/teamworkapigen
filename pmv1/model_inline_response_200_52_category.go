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

// InlineResponse20052Category struct for InlineResponse20052Category
type InlineResponse20052Category struct {
	Color *string `json:"color,omitempty"`
	Count *string `json:"count,omitempty"`
	ElementsCount *string `json:"elements_count,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	ParentId *string `json:"parent-id,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewInlineResponse20052Category instantiates a new InlineResponse20052Category object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20052Category() *InlineResponse20052Category {
	this := InlineResponse20052Category{}
	return &this
}

// NewInlineResponse20052CategoryWithDefaults instantiates a new InlineResponse20052Category object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20052CategoryWithDefaults() *InlineResponse20052Category {
	this := InlineResponse20052Category{}
	return &this
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *InlineResponse20052Category) GetColor() string {
	if o == nil || o.Color == nil {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20052Category) GetColorOk() (*string, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *InlineResponse20052Category) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *InlineResponse20052Category) SetColor(v string) {
	o.Color = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *InlineResponse20052Category) GetCount() string {
	if o == nil || o.Count == nil {
		var ret string
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20052Category) GetCountOk() (*string, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *InlineResponse20052Category) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given string and assigns it to the Count field.
func (o *InlineResponse20052Category) SetCount(v string) {
	o.Count = &v
}

// GetElementsCount returns the ElementsCount field value if set, zero value otherwise.
func (o *InlineResponse20052Category) GetElementsCount() string {
	if o == nil || o.ElementsCount == nil {
		var ret string
		return ret
	}
	return *o.ElementsCount
}

// GetElementsCountOk returns a tuple with the ElementsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20052Category) GetElementsCountOk() (*string, bool) {
	if o == nil || o.ElementsCount == nil {
		return nil, false
	}
	return o.ElementsCount, true
}

// HasElementsCount returns a boolean if a field has been set.
func (o *InlineResponse20052Category) HasElementsCount() bool {
	if o != nil && o.ElementsCount != nil {
		return true
	}

	return false
}

// SetElementsCount gets a reference to the given string and assigns it to the ElementsCount field.
func (o *InlineResponse20052Category) SetElementsCount(v string) {
	o.ElementsCount = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20052Category) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20052Category) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20052Category) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20052Category) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20052Category) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20052Category) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20052Category) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20052Category) SetName(v string) {
	o.Name = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *InlineResponse20052Category) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20052Category) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *InlineResponse20052Category) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *InlineResponse20052Category) SetParentId(v string) {
	o.ParentId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse20052Category) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20052Category) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse20052Category) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse20052Category) SetType(v string) {
	o.Type = &v
}

func (o InlineResponse20052Category) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Color != nil {
		toSerialize["color"] = o.Color
	}
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
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20052Category struct {
	value *InlineResponse20052Category
	isSet bool
}

func (v NullableInlineResponse20052Category) Get() *InlineResponse20052Category {
	return v.value
}

func (v *NullableInlineResponse20052Category) Set(val *InlineResponse20052Category) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20052Category) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20052Category) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20052Category(val *InlineResponse20052Category) *NullableInlineResponse20052Category {
	return &NullableInlineResponse20052Category{value: val, isSet: true}
}

func (v NullableInlineResponse20052Category) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20052Category) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

