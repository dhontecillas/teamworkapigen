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

// PostsIdJsonPost struct for PostsIdJsonPost
type PostsIdJsonPost struct {
	Attachments *string `json:"attachments,omitempty"`
	Body *string `json:"body,omitempty"`
	CategoryId *string `json:"category-id,omitempty"`
	Notify *[]string `json:"notify,omitempty"`
	PendingFileAttachments *string `json:"pendingFileAttachments,omitempty"`
	Private *string `json:"private,omitempty"`
	Title string `json:"title"`
}

// NewPostsIdJsonPost instantiates a new PostsIdJsonPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostsIdJsonPost(title string) *PostsIdJsonPost {
	this := PostsIdJsonPost{}
	this.Title = title
	return &this
}

// NewPostsIdJsonPostWithDefaults instantiates a new PostsIdJsonPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostsIdJsonPostWithDefaults() *PostsIdJsonPost {
	this := PostsIdJsonPost{}
	return &this
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *PostsIdJsonPost) GetAttachments() string {
	if o == nil || o.Attachments == nil {
		var ret string
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostsIdJsonPost) GetAttachmentsOk() (*string, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *PostsIdJsonPost) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given string and assigns it to the Attachments field.
func (o *PostsIdJsonPost) SetAttachments(v string) {
	o.Attachments = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *PostsIdJsonPost) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostsIdJsonPost) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *PostsIdJsonPost) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *PostsIdJsonPost) SetBody(v string) {
	o.Body = &v
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *PostsIdJsonPost) GetCategoryId() string {
	if o == nil || o.CategoryId == nil {
		var ret string
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostsIdJsonPost) GetCategoryIdOk() (*string, bool) {
	if o == nil || o.CategoryId == nil {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *PostsIdJsonPost) HasCategoryId() bool {
	if o != nil && o.CategoryId != nil {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given string and assigns it to the CategoryId field.
func (o *PostsIdJsonPost) SetCategoryId(v string) {
	o.CategoryId = &v
}

// GetNotify returns the Notify field value if set, zero value otherwise.
func (o *PostsIdJsonPost) GetNotify() []string {
	if o == nil || o.Notify == nil {
		var ret []string
		return ret
	}
	return *o.Notify
}

// GetNotifyOk returns a tuple with the Notify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostsIdJsonPost) GetNotifyOk() (*[]string, bool) {
	if o == nil || o.Notify == nil {
		return nil, false
	}
	return o.Notify, true
}

// HasNotify returns a boolean if a field has been set.
func (o *PostsIdJsonPost) HasNotify() bool {
	if o != nil && o.Notify != nil {
		return true
	}

	return false
}

// SetNotify gets a reference to the given []string and assigns it to the Notify field.
func (o *PostsIdJsonPost) SetNotify(v []string) {
	o.Notify = &v
}

// GetPendingFileAttachments returns the PendingFileAttachments field value if set, zero value otherwise.
func (o *PostsIdJsonPost) GetPendingFileAttachments() string {
	if o == nil || o.PendingFileAttachments == nil {
		var ret string
		return ret
	}
	return *o.PendingFileAttachments
}

// GetPendingFileAttachmentsOk returns a tuple with the PendingFileAttachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostsIdJsonPost) GetPendingFileAttachmentsOk() (*string, bool) {
	if o == nil || o.PendingFileAttachments == nil {
		return nil, false
	}
	return o.PendingFileAttachments, true
}

// HasPendingFileAttachments returns a boolean if a field has been set.
func (o *PostsIdJsonPost) HasPendingFileAttachments() bool {
	if o != nil && o.PendingFileAttachments != nil {
		return true
	}

	return false
}

// SetPendingFileAttachments gets a reference to the given string and assigns it to the PendingFileAttachments field.
func (o *PostsIdJsonPost) SetPendingFileAttachments(v string) {
	o.PendingFileAttachments = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *PostsIdJsonPost) GetPrivate() string {
	if o == nil || o.Private == nil {
		var ret string
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostsIdJsonPost) GetPrivateOk() (*string, bool) {
	if o == nil || o.Private == nil {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *PostsIdJsonPost) HasPrivate() bool {
	if o != nil && o.Private != nil {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given string and assigns it to the Private field.
func (o *PostsIdJsonPost) SetPrivate(v string) {
	o.Private = &v
}

// GetTitle returns the Title field value
func (o *PostsIdJsonPost) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *PostsIdJsonPost) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *PostsIdJsonPost) SetTitle(v string) {
	o.Title = v
}

func (o PostsIdJsonPost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.CategoryId != nil {
		toSerialize["category-id"] = o.CategoryId
	}
	if o.Notify != nil {
		toSerialize["notify"] = o.Notify
	}
	if o.PendingFileAttachments != nil {
		toSerialize["pendingFileAttachments"] = o.PendingFileAttachments
	}
	if o.Private != nil {
		toSerialize["private"] = o.Private
	}
	if true {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullablePostsIdJsonPost struct {
	value *PostsIdJsonPost
	isSet bool
}

func (v NullablePostsIdJsonPost) Get() *PostsIdJsonPost {
	return v.value
}

func (v *NullablePostsIdJsonPost) Set(val *PostsIdJsonPost) {
	v.value = val
	v.isSet = true
}

func (v NullablePostsIdJsonPost) IsSet() bool {
	return v.isSet
}

func (v *NullablePostsIdJsonPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostsIdJsonPost(val *PostsIdJsonPost) *NullablePostsIdJsonPost {
	return &NullablePostsIdJsonPost{value: val, isSet: true}
}

func (v NullablePostsIdJsonPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostsIdJsonPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


