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

// InlineResponse200115ContributingUsers struct for InlineResponse200115ContributingUsers
type InlineResponse200115ContributingUsers struct {
	AvatarURL *string `json:"avatarURL,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	FullName *string `json:"fullName,omitempty"`
	Id *string `json:"id,omitempty"`
	LastName *string `json:"lastName,omitempty"`
}

// NewInlineResponse200115ContributingUsers instantiates a new InlineResponse200115ContributingUsers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200115ContributingUsers() *InlineResponse200115ContributingUsers {
	this := InlineResponse200115ContributingUsers{}
	return &this
}

// NewInlineResponse200115ContributingUsersWithDefaults instantiates a new InlineResponse200115ContributingUsers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200115ContributingUsersWithDefaults() *InlineResponse200115ContributingUsers {
	this := InlineResponse200115ContributingUsers{}
	return &this
}

// GetAvatarURL returns the AvatarURL field value if set, zero value otherwise.
func (o *InlineResponse200115ContributingUsers) GetAvatarURL() string {
	if o == nil || o.AvatarURL == nil {
		var ret string
		return ret
	}
	return *o.AvatarURL
}

// GetAvatarURLOk returns a tuple with the AvatarURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200115ContributingUsers) GetAvatarURLOk() (*string, bool) {
	if o == nil || o.AvatarURL == nil {
		return nil, false
	}
	return o.AvatarURL, true
}

// HasAvatarURL returns a boolean if a field has been set.
func (o *InlineResponse200115ContributingUsers) HasAvatarURL() bool {
	if o != nil && o.AvatarURL != nil {
		return true
	}

	return false
}

// SetAvatarURL gets a reference to the given string and assigns it to the AvatarURL field.
func (o *InlineResponse200115ContributingUsers) SetAvatarURL(v string) {
	o.AvatarURL = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *InlineResponse200115ContributingUsers) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200115ContributingUsers) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *InlineResponse200115ContributingUsers) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *InlineResponse200115ContributingUsers) SetFirstName(v string) {
	o.FirstName = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *InlineResponse200115ContributingUsers) GetFullName() string {
	if o == nil || o.FullName == nil {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200115ContributingUsers) GetFullNameOk() (*string, bool) {
	if o == nil || o.FullName == nil {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *InlineResponse200115ContributingUsers) HasFullName() bool {
	if o != nil && o.FullName != nil {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *InlineResponse200115ContributingUsers) SetFullName(v string) {
	o.FullName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200115ContributingUsers) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200115ContributingUsers) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200115ContributingUsers) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200115ContributingUsers) SetId(v string) {
	o.Id = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *InlineResponse200115ContributingUsers) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200115ContributingUsers) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *InlineResponse200115ContributingUsers) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *InlineResponse200115ContributingUsers) SetLastName(v string) {
	o.LastName = &v
}

func (o InlineResponse200115ContributingUsers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AvatarURL != nil {
		toSerialize["avatarURL"] = o.AvatarURL
	}
	if o.FirstName != nil {
		toSerialize["firstName"] = o.FirstName
	}
	if o.FullName != nil {
		toSerialize["fullName"] = o.FullName
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastName != nil {
		toSerialize["lastName"] = o.LastName
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200115ContributingUsers struct {
	value *InlineResponse200115ContributingUsers
	isSet bool
}

func (v NullableInlineResponse200115ContributingUsers) Get() *InlineResponse200115ContributingUsers {
	return v.value
}

func (v *NullableInlineResponse200115ContributingUsers) Set(val *InlineResponse200115ContributingUsers) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200115ContributingUsers) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200115ContributingUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200115ContributingUsers(val *InlineResponse200115ContributingUsers) *NullableInlineResponse200115ContributingUsers {
	return &NullableInlineResponse200115ContributingUsers{value: val, isSet: true}
}

func (v NullableInlineResponse200115ContributingUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200115ContributingUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

