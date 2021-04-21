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

// ArgCustomField CustomField stores the id, value and operator of the custom field.
type ArgCustomField struct {
	ID *int32 `json:"ID,omitempty"`
	Operator *string `json:"Operator,omitempty"`
	Value *string `json:"Value,omitempty"`
}

// NewArgCustomField instantiates a new ArgCustomField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArgCustomField() *ArgCustomField {
	this := ArgCustomField{}
	return &this
}

// NewArgCustomFieldWithDefaults instantiates a new ArgCustomField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArgCustomFieldWithDefaults() *ArgCustomField {
	this := ArgCustomField{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *ArgCustomField) GetID() int32 {
	if o == nil || o.ID == nil {
		var ret int32
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArgCustomField) GetIDOk() (*int32, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *ArgCustomField) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given int32 and assigns it to the ID field.
func (o *ArgCustomField) SetID(v int32) {
	o.ID = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *ArgCustomField) GetOperator() string {
	if o == nil || o.Operator == nil {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArgCustomField) GetOperatorOk() (*string, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *ArgCustomField) HasOperator() bool {
	if o != nil && o.Operator != nil {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *ArgCustomField) SetOperator(v string) {
	o.Operator = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ArgCustomField) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArgCustomField) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ArgCustomField) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ArgCustomField) SetValue(v string) {
	o.Value = &v
}

func (o ArgCustomField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ID != nil {
		toSerialize["ID"] = o.ID
	}
	if o.Operator != nil {
		toSerialize["Operator"] = o.Operator
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableArgCustomField struct {
	value *ArgCustomField
	isSet bool
}

func (v NullableArgCustomField) Get() *ArgCustomField {
	return v.value
}

func (v *NullableArgCustomField) Set(val *ArgCustomField) {
	v.value = val
	v.isSet = true
}

func (v NullableArgCustomField) IsSet() bool {
	return v.isSet
}

func (v *NullableArgCustomField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArgCustomField(val *ArgCustomField) *NullableArgCustomField {
	return &NullableArgCustomField{value: val, isSet: true}
}

func (v NullableArgCustomField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArgCustomField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


