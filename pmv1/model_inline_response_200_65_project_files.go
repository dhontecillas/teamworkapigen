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

// InlineResponse20065ProjectFiles struct for InlineResponse20065ProjectFiles
type InlineResponse20065ProjectFiles struct {
	CategoryId *string `json:"category-id,omitempty"`
	CategoryName *string `json:"category-name,omitempty"`
	CommentsCount *string `json:"comments-count,omitempty"`
	Description *string `json:"description,omitempty"`
	ExtraData *string `json:"extraData,omitempty"`
	FileSource *string `json:"file-source,omitempty"`
	Id *string `json:"id,omitempty"`
	LastChangedOn *string `json:"last-changed-on,omitempty"`
	Name *string `json:"name,omitempty"`
	OriginalName *string `json:"originalName,omitempty"`
	Private *string `json:"private,omitempty"`
	ProjectId *string `json:"project-id,omitempty"`
	Size *string `json:"size,omitempty"`
	Status *string `json:"status,omitempty"`
	ThumbURL *string `json:"thumbURL,omitempty"`
	UploadedByUserFirstName *string `json:"uploaded-by-user-first-name,omitempty"`
	UploadedByUserLastName *string `json:"uploaded-by-user-last-name,omitempty"`
	UploadedByUserId *string `json:"uploaded-by-userId,omitempty"`
	UploadedDate *string `json:"uploaded-date,omitempty"`
	Version *string `json:"version,omitempty"`
	VersionId *string `json:"version-id,omitempty"`
	Versions *[]map[string]interface{} `json:"versions,omitempty"`
}

// NewInlineResponse20065ProjectFiles instantiates a new InlineResponse20065ProjectFiles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20065ProjectFiles() *InlineResponse20065ProjectFiles {
	this := InlineResponse20065ProjectFiles{}
	return &this
}

// NewInlineResponse20065ProjectFilesWithDefaults instantiates a new InlineResponse20065ProjectFiles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20065ProjectFilesWithDefaults() *InlineResponse20065ProjectFiles {
	this := InlineResponse20065ProjectFiles{}
	return &this
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *InlineResponse20065ProjectFiles) GetCategoryId() string {
	if o == nil || o.CategoryId == nil {
		var ret string
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065ProjectFiles) GetCategoryIdOk() (*string, bool) {
	if o == nil || o.CategoryId == nil {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *InlineResponse20065ProjectFiles) HasCategoryId() bool {
	if o != nil && o.CategoryId != nil {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given string and assigns it to the CategoryId field.
func (o *InlineResponse20065ProjectFiles) SetCategoryId(v string) {
	o.CategoryId = &v
}

// GetCategoryName returns the CategoryName field value if set, zero value otherwise.
func (o *InlineResponse20065ProjectFiles) GetCategoryName() string {
	if o == nil || o.CategoryName == nil {
		var ret string
		return ret
	}
	return *o.CategoryName
}

// GetCategoryNameOk returns a tuple with the CategoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065ProjectFiles) GetCategoryNameOk() (*string, bool) {
	if o == nil || o.CategoryName == nil {
		return nil, false
	}
	return o.CategoryName, true
}

// HasCategoryName returns a boolean if a field has been set.
func (o *InlineResponse20065ProjectFiles) HasCategoryName() bool {
	if o != nil && o.CategoryName != nil {
		return true
	}

	return false
}

// SetCategoryName gets a reference to the given string and assigns it to the CategoryName field.
func (o *InlineResponse20065ProjectFiles) SetCategoryName(v string) {
	o.CategoryName = &v
}

// GetCommentsCount returns the CommentsCount field value if set, zero value otherwise.
func (o *InlineResponse20065ProjectFiles) GetCommentsCount() string {
	if o == nil || o.CommentsCount == nil {
		var ret string
		return ret
	}
	return *o.CommentsCount
}

// GetCommentsCountOk returns a tuple with the CommentsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065ProjectFiles) GetCommentsCountOk() (*string, bool) {
	if o == nil || o.CommentsCount == nil {
		return nil, false
	}
	return o.CommentsCount, true
}

// HasCommentsCount returns a boolean if a field has been set.
func (o *InlineResponse20065ProjectFiles) HasCommentsCount() bool {
	if o != nil && o.CommentsCount != nil {
		return true
	}

	return false
}

// SetCommentsCount gets a reference to the given string and assigns it to the CommentsCount field.
func (o *InlineResponse20065ProjectFiles) SetCommentsCount(v string) {
	o.CommentsCount = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineResponse20065ProjectFiles) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065ProjectFiles) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineResponse20065ProjectFiles) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineResponse20065ProjectFiles) SetDescription(v string) {
	o.Description = &v
}

// GetExtraData returns the ExtraData field value if set, zero value otherwise.
func (o *InlineResponse20065ProjectFiles) GetExtraData() string {
	if o == nil || o.ExtraData == nil {
		var ret string
		return ret
	}
	return *o.ExtraData
}

// GetExtraDataOk returns a tuple with the ExtraData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065ProjectFiles) GetExtraDataOk() (*string, bool) {
	if o == nil || o.ExtraData == nil {
		return nil, false
	}
	return o.ExtraData, true
}

// HasExtraData returns a boolean if a field has been set.
func (o *InlineResponse20065ProjectFiles) HasExtraData() bool {
	if o != nil && o.ExtraData != nil {
		return true
	}

	return false
}

// SetExtraData gets a reference to the given string and assigns it to the ExtraData field.
func (o *InlineResponse20065ProjectFiles) SetExtraData(v string) {
	o.ExtraData = &v
}

// GetFileSource returns the FileSource field value if set, zero value otherwise.
func (o *InlineResponse20065ProjectFiles) GetFileSource() string {
	if o == nil || o.FileSource == nil {
		var ret string
		return ret
	}
	return *o.FileSource
}

// GetFileSourceOk returns a tuple with the FileSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065ProjectFiles) GetFileSourceOk() (*string, bool) {
	if o == nil || o.FileSource == nil {
		return nil, false
	}
	return o.FileSource, true
}

// HasFileSource returns a boolean if a field has been set.
func (o *InlineResponse20065ProjectFiles) HasFileSource() bool {
	if o != nil && o.FileSource != nil {
		return true
	}

	return false
}

// SetFileSource gets a reference to the given string and assigns it to the FileSource field.
func (o *InlineResponse20065ProjectFiles) SetFileSource(v string) {
	o.FileSource = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20065ProjectFiles) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065ProjectFiles) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20065ProjectFiles) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20065ProjectFiles) SetId(v string) {
	o.Id = &v
}

// GetLastChangedOn returns the LastChangedOn field value if set, zero value otherwise.
func (o *InlineResponse20065ProjectFiles) GetLastChangedOn() string {
	if o == nil || o.LastChangedOn == nil {
		var ret string
		return ret
	}
	return *o.LastChangedOn
}

// GetLastChangedOnOk returns a tuple with the LastChangedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065ProjectFiles) GetLastChangedOnOk() (*string, bool) {
	if o == nil || o.LastChangedOn == nil {
		return nil, false
	}
	return o.LastChangedOn, true
}

// HasLastChangedOn returns a boolean if a field has been set.
func (o *InlineResponse20065ProjectFiles) HasLastChangedOn() bool {
	if o != nil && o.LastChangedOn != nil {
		return true
	}

	return false
}

// SetLastChangedOn gets a reference to the given string and assigns it to the LastChangedOn field.
func (o *InlineResponse20065ProjectFiles) SetLastChangedOn(v string) {
	o.LastChangedOn = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20065ProjectFiles) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065ProjectFiles) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20065ProjectFiles) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20065ProjectFiles) SetName(v string) {
	o.Name = &v
}

// GetOriginalName returns the OriginalName field value if set, zero value otherwise.
func (o *InlineResponse20065ProjectFiles) GetOriginalName() string {
	if o == nil || o.OriginalName == nil {
		var ret string
		return ret
	}
	return *o.OriginalName
}

// GetOriginalNameOk returns a tuple with the OriginalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065ProjectFiles) GetOriginalNameOk() (*string, bool) {
	if o == nil || o.OriginalName == nil {
		return nil, false
	}
	return o.OriginalName, true
}

// HasOriginalName returns a boolean if a field has been set.
func (o *InlineResponse20065ProjectFiles) HasOriginalName() bool {
	if o != nil && o.OriginalName != nil {
		return true
	}

	return false
}

// SetOriginalName gets a reference to the given string and assigns it to the OriginalName field.
func (o *InlineResponse20065ProjectFiles) SetOriginalName(v string) {
	o.OriginalName = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *InlineResponse20065ProjectFiles) GetPrivate() string {
	if o == nil || o.Private == nil {
		var ret string
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065ProjectFiles) GetPrivateOk() (*string, bool) {
	if o == nil || o.Private == nil {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *InlineResponse20065ProjectFiles) HasPrivate() bool {
	if o != nil && o.Private != nil {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given string and assigns it to the Private field.
func (o *InlineResponse20065ProjectFiles) SetPrivate(v string) {
	o.Private = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *InlineResponse20065ProjectFiles) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065ProjectFiles) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *InlineResponse20065ProjectFiles) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *InlineResponse20065ProjectFiles) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *InlineResponse20065ProjectFiles) GetSize() string {
	if o == nil || o.Size == nil {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065ProjectFiles) GetSizeOk() (*string, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *InlineResponse20065ProjectFiles) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *InlineResponse20065ProjectFiles) SetSize(v string) {
	o.Size = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse20065ProjectFiles) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065ProjectFiles) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse20065ProjectFiles) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse20065ProjectFiles) SetStatus(v string) {
	o.Status = &v
}

// GetThumbURL returns the ThumbURL field value if set, zero value otherwise.
func (o *InlineResponse20065ProjectFiles) GetThumbURL() string {
	if o == nil || o.ThumbURL == nil {
		var ret string
		return ret
	}
	return *o.ThumbURL
}

// GetThumbURLOk returns a tuple with the ThumbURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065ProjectFiles) GetThumbURLOk() (*string, bool) {
	if o == nil || o.ThumbURL == nil {
		return nil, false
	}
	return o.ThumbURL, true
}

// HasThumbURL returns a boolean if a field has been set.
func (o *InlineResponse20065ProjectFiles) HasThumbURL() bool {
	if o != nil && o.ThumbURL != nil {
		return true
	}

	return false
}

// SetThumbURL gets a reference to the given string and assigns it to the ThumbURL field.
func (o *InlineResponse20065ProjectFiles) SetThumbURL(v string) {
	o.ThumbURL = &v
}

// GetUploadedByUserFirstName returns the UploadedByUserFirstName field value if set, zero value otherwise.
func (o *InlineResponse20065ProjectFiles) GetUploadedByUserFirstName() string {
	if o == nil || o.UploadedByUserFirstName == nil {
		var ret string
		return ret
	}
	return *o.UploadedByUserFirstName
}

// GetUploadedByUserFirstNameOk returns a tuple with the UploadedByUserFirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065ProjectFiles) GetUploadedByUserFirstNameOk() (*string, bool) {
	if o == nil || o.UploadedByUserFirstName == nil {
		return nil, false
	}
	return o.UploadedByUserFirstName, true
}

// HasUploadedByUserFirstName returns a boolean if a field has been set.
func (o *InlineResponse20065ProjectFiles) HasUploadedByUserFirstName() bool {
	if o != nil && o.UploadedByUserFirstName != nil {
		return true
	}

	return false
}

// SetUploadedByUserFirstName gets a reference to the given string and assigns it to the UploadedByUserFirstName field.
func (o *InlineResponse20065ProjectFiles) SetUploadedByUserFirstName(v string) {
	o.UploadedByUserFirstName = &v
}

// GetUploadedByUserLastName returns the UploadedByUserLastName field value if set, zero value otherwise.
func (o *InlineResponse20065ProjectFiles) GetUploadedByUserLastName() string {
	if o == nil || o.UploadedByUserLastName == nil {
		var ret string
		return ret
	}
	return *o.UploadedByUserLastName
}

// GetUploadedByUserLastNameOk returns a tuple with the UploadedByUserLastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065ProjectFiles) GetUploadedByUserLastNameOk() (*string, bool) {
	if o == nil || o.UploadedByUserLastName == nil {
		return nil, false
	}
	return o.UploadedByUserLastName, true
}

// HasUploadedByUserLastName returns a boolean if a field has been set.
func (o *InlineResponse20065ProjectFiles) HasUploadedByUserLastName() bool {
	if o != nil && o.UploadedByUserLastName != nil {
		return true
	}

	return false
}

// SetUploadedByUserLastName gets a reference to the given string and assigns it to the UploadedByUserLastName field.
func (o *InlineResponse20065ProjectFiles) SetUploadedByUserLastName(v string) {
	o.UploadedByUserLastName = &v
}

// GetUploadedByUserId returns the UploadedByUserId field value if set, zero value otherwise.
func (o *InlineResponse20065ProjectFiles) GetUploadedByUserId() string {
	if o == nil || o.UploadedByUserId == nil {
		var ret string
		return ret
	}
	return *o.UploadedByUserId
}

// GetUploadedByUserIdOk returns a tuple with the UploadedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065ProjectFiles) GetUploadedByUserIdOk() (*string, bool) {
	if o == nil || o.UploadedByUserId == nil {
		return nil, false
	}
	return o.UploadedByUserId, true
}

// HasUploadedByUserId returns a boolean if a field has been set.
func (o *InlineResponse20065ProjectFiles) HasUploadedByUserId() bool {
	if o != nil && o.UploadedByUserId != nil {
		return true
	}

	return false
}

// SetUploadedByUserId gets a reference to the given string and assigns it to the UploadedByUserId field.
func (o *InlineResponse20065ProjectFiles) SetUploadedByUserId(v string) {
	o.UploadedByUserId = &v
}

// GetUploadedDate returns the UploadedDate field value if set, zero value otherwise.
func (o *InlineResponse20065ProjectFiles) GetUploadedDate() string {
	if o == nil || o.UploadedDate == nil {
		var ret string
		return ret
	}
	return *o.UploadedDate
}

// GetUploadedDateOk returns a tuple with the UploadedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065ProjectFiles) GetUploadedDateOk() (*string, bool) {
	if o == nil || o.UploadedDate == nil {
		return nil, false
	}
	return o.UploadedDate, true
}

// HasUploadedDate returns a boolean if a field has been set.
func (o *InlineResponse20065ProjectFiles) HasUploadedDate() bool {
	if o != nil && o.UploadedDate != nil {
		return true
	}

	return false
}

// SetUploadedDate gets a reference to the given string and assigns it to the UploadedDate field.
func (o *InlineResponse20065ProjectFiles) SetUploadedDate(v string) {
	o.UploadedDate = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *InlineResponse20065ProjectFiles) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065ProjectFiles) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *InlineResponse20065ProjectFiles) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *InlineResponse20065ProjectFiles) SetVersion(v string) {
	o.Version = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *InlineResponse20065ProjectFiles) GetVersionId() string {
	if o == nil || o.VersionId == nil {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065ProjectFiles) GetVersionIdOk() (*string, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *InlineResponse20065ProjectFiles) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *InlineResponse20065ProjectFiles) SetVersionId(v string) {
	o.VersionId = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *InlineResponse20065ProjectFiles) GetVersions() []map[string]interface{} {
	if o == nil || o.Versions == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20065ProjectFiles) GetVersionsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Versions == nil {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *InlineResponse20065ProjectFiles) HasVersions() bool {
	if o != nil && o.Versions != nil {
		return true
	}

	return false
}

// SetVersions gets a reference to the given []map[string]interface{} and assigns it to the Versions field.
func (o *InlineResponse20065ProjectFiles) SetVersions(v []map[string]interface{}) {
	o.Versions = &v
}

func (o InlineResponse20065ProjectFiles) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CategoryId != nil {
		toSerialize["category-id"] = o.CategoryId
	}
	if o.CategoryName != nil {
		toSerialize["category-name"] = o.CategoryName
	}
	if o.CommentsCount != nil {
		toSerialize["comments-count"] = o.CommentsCount
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ExtraData != nil {
		toSerialize["extraData"] = o.ExtraData
	}
	if o.FileSource != nil {
		toSerialize["file-source"] = o.FileSource
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastChangedOn != nil {
		toSerialize["last-changed-on"] = o.LastChangedOn
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OriginalName != nil {
		toSerialize["originalName"] = o.OriginalName
	}
	if o.Private != nil {
		toSerialize["private"] = o.Private
	}
	if o.ProjectId != nil {
		toSerialize["project-id"] = o.ProjectId
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ThumbURL != nil {
		toSerialize["thumbURL"] = o.ThumbURL
	}
	if o.UploadedByUserFirstName != nil {
		toSerialize["uploaded-by-user-first-name"] = o.UploadedByUserFirstName
	}
	if o.UploadedByUserLastName != nil {
		toSerialize["uploaded-by-user-last-name"] = o.UploadedByUserLastName
	}
	if o.UploadedByUserId != nil {
		toSerialize["uploaded-by-userId"] = o.UploadedByUserId
	}
	if o.UploadedDate != nil {
		toSerialize["uploaded-date"] = o.UploadedDate
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.VersionId != nil {
		toSerialize["version-id"] = o.VersionId
	}
	if o.Versions != nil {
		toSerialize["versions"] = o.Versions
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20065ProjectFiles struct {
	value *InlineResponse20065ProjectFiles
	isSet bool
}

func (v NullableInlineResponse20065ProjectFiles) Get() *InlineResponse20065ProjectFiles {
	return v.value
}

func (v *NullableInlineResponse20065ProjectFiles) Set(val *InlineResponse20065ProjectFiles) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20065ProjectFiles) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20065ProjectFiles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20065ProjectFiles(val *InlineResponse20065ProjectFiles) *NullableInlineResponse20065ProjectFiles {
	return &NullableInlineResponse20065ProjectFiles{value: val, isSet: true}
}

func (v NullableInlineResponse20065ProjectFiles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20065ProjectFiles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

