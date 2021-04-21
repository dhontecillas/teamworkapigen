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

// ViewComment Comment contains all the information returned from a comment.
type ViewComment struct {
	Id *int32 `json:"id,omitempty"`
	Object *ViewRelationship `json:"object,omitempty"`
	ObjectId *int32 `json:"objectId,omitempty"`
	ObjectType *string `json:"objectType,omitempty"`
	Title *string `json:"title,omitempty"`
}

// NewViewComment instantiates a new ViewComment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewComment() *ViewComment {
	this := ViewComment{}
	return &this
}

// NewViewCommentWithDefaults instantiates a new ViewComment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewCommentWithDefaults() *ViewComment {
	this := ViewComment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ViewComment) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewComment) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ViewComment) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ViewComment) SetId(v int32) {
	o.Id = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *ViewComment) GetObject() ViewRelationship {
	if o == nil || o.Object == nil {
		var ret ViewRelationship
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewComment) GetObjectOk() (*ViewRelationship, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *ViewComment) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given ViewRelationship and assigns it to the Object field.
func (o *ViewComment) SetObject(v ViewRelationship) {
	o.Object = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *ViewComment) GetObjectId() int32 {
	if o == nil || o.ObjectId == nil {
		var ret int32
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewComment) GetObjectIdOk() (*int32, bool) {
	if o == nil || o.ObjectId == nil {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *ViewComment) HasObjectId() bool {
	if o != nil && o.ObjectId != nil {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given int32 and assigns it to the ObjectId field.
func (o *ViewComment) SetObjectId(v int32) {
	o.ObjectId = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *ViewComment) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewComment) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *ViewComment) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *ViewComment) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ViewComment) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewComment) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ViewComment) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ViewComment) SetTitle(v string) {
	o.Title = &v
}

func (o ViewComment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.ObjectId != nil {
		toSerialize["objectId"] = o.ObjectId
	}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullableViewComment struct {
	value *ViewComment
	isSet bool
}

func (v NullableViewComment) Get() *ViewComment {
	return v.value
}

func (v *NullableViewComment) Set(val *ViewComment) {
	v.value = val
	v.isSet = true
}

func (v NullableViewComment) IsSet() bool {
	return v.isSet
}

func (v *NullableViewComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewComment(val *ViewComment) *NullableViewComment {
	return &NullableViewComment{value: val, isSet: true}
}

func (v NullableViewComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


