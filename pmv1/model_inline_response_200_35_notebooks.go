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

// InlineResponse20035Notebooks struct for InlineResponse20035Notebooks
type InlineResponse20035Notebooks struct {
	CategoryId *string `json:"category-id,omitempty"`
	CategoryName *string `json:"category-name,omitempty"`
	CommentsCount *string `json:"comments-count,omitempty"`
	CreatedByUserId *string `json:"created-by-userId,omitempty"`
	CreatedByUserfirstname *string `json:"created-by-userfirstname,omitempty"`
	CreatedByUserlastname *string `json:"created-by-userlastname,omitempty"`
	CreatedDate *string `json:"created-date,omitempty"`
	Description *string `json:"description,omitempty"`
	Id *string `json:"id,omitempty"`
	LockdownId *string `json:"lockdown-id,omitempty"`
	Locked *string `json:"locked,omitempty"`
	Name *string `json:"name,omitempty"`
	NrComments *string `json:"nr-comments,omitempty"`
	Private *string `json:"private,omitempty"`
	ProjectId *string `json:"project-id,omitempty"`
	ReadCommentsCount *string `json:"read-comments-count,omitempty"`
	UpdatedByUserId *string `json:"updated-by-userId,omitempty"`
	UpdatedByUserfirstname *string `json:"updated-by-userfirstname,omitempty"`
	UpdatedByUserlastname *string `json:"updated-by-userlastname,omitempty"`
	UpdatedDate *string `json:"updated-date,omitempty"`
	UserDisplayUpdatedDate *string `json:"user-display-updated-date,omitempty"`
	UserDisplayUpdatedTime *string `json:"user-display-updated-time,omitempty"`
	Version *string `json:"version,omitempty"`
	VersionId *string `json:"version-id,omitempty"`
}

// NewInlineResponse20035Notebooks instantiates a new InlineResponse20035Notebooks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20035Notebooks() *InlineResponse20035Notebooks {
	this := InlineResponse20035Notebooks{}
	return &this
}

// NewInlineResponse20035NotebooksWithDefaults instantiates a new InlineResponse20035Notebooks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20035NotebooksWithDefaults() *InlineResponse20035Notebooks {
	this := InlineResponse20035Notebooks{}
	return &this
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *InlineResponse20035Notebooks) GetCategoryId() string {
	if o == nil || o.CategoryId == nil {
		var ret string
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035Notebooks) GetCategoryIdOk() (*string, bool) {
	if o == nil || o.CategoryId == nil {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *InlineResponse20035Notebooks) HasCategoryId() bool {
	if o != nil && o.CategoryId != nil {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given string and assigns it to the CategoryId field.
func (o *InlineResponse20035Notebooks) SetCategoryId(v string) {
	o.CategoryId = &v
}

// GetCategoryName returns the CategoryName field value if set, zero value otherwise.
func (o *InlineResponse20035Notebooks) GetCategoryName() string {
	if o == nil || o.CategoryName == nil {
		var ret string
		return ret
	}
	return *o.CategoryName
}

// GetCategoryNameOk returns a tuple with the CategoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035Notebooks) GetCategoryNameOk() (*string, bool) {
	if o == nil || o.CategoryName == nil {
		return nil, false
	}
	return o.CategoryName, true
}

// HasCategoryName returns a boolean if a field has been set.
func (o *InlineResponse20035Notebooks) HasCategoryName() bool {
	if o != nil && o.CategoryName != nil {
		return true
	}

	return false
}

// SetCategoryName gets a reference to the given string and assigns it to the CategoryName field.
func (o *InlineResponse20035Notebooks) SetCategoryName(v string) {
	o.CategoryName = &v
}

// GetCommentsCount returns the CommentsCount field value if set, zero value otherwise.
func (o *InlineResponse20035Notebooks) GetCommentsCount() string {
	if o == nil || o.CommentsCount == nil {
		var ret string
		return ret
	}
	return *o.CommentsCount
}

// GetCommentsCountOk returns a tuple with the CommentsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035Notebooks) GetCommentsCountOk() (*string, bool) {
	if o == nil || o.CommentsCount == nil {
		return nil, false
	}
	return o.CommentsCount, true
}

// HasCommentsCount returns a boolean if a field has been set.
func (o *InlineResponse20035Notebooks) HasCommentsCount() bool {
	if o != nil && o.CommentsCount != nil {
		return true
	}

	return false
}

// SetCommentsCount gets a reference to the given string and assigns it to the CommentsCount field.
func (o *InlineResponse20035Notebooks) SetCommentsCount(v string) {
	o.CommentsCount = &v
}

// GetCreatedByUserId returns the CreatedByUserId field value if set, zero value otherwise.
func (o *InlineResponse20035Notebooks) GetCreatedByUserId() string {
	if o == nil || o.CreatedByUserId == nil {
		var ret string
		return ret
	}
	return *o.CreatedByUserId
}

// GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035Notebooks) GetCreatedByUserIdOk() (*string, bool) {
	if o == nil || o.CreatedByUserId == nil {
		return nil, false
	}
	return o.CreatedByUserId, true
}

// HasCreatedByUserId returns a boolean if a field has been set.
func (o *InlineResponse20035Notebooks) HasCreatedByUserId() bool {
	if o != nil && o.CreatedByUserId != nil {
		return true
	}

	return false
}

// SetCreatedByUserId gets a reference to the given string and assigns it to the CreatedByUserId field.
func (o *InlineResponse20035Notebooks) SetCreatedByUserId(v string) {
	o.CreatedByUserId = &v
}

// GetCreatedByUserfirstname returns the CreatedByUserfirstname field value if set, zero value otherwise.
func (o *InlineResponse20035Notebooks) GetCreatedByUserfirstname() string {
	if o == nil || o.CreatedByUserfirstname == nil {
		var ret string
		return ret
	}
	return *o.CreatedByUserfirstname
}

// GetCreatedByUserfirstnameOk returns a tuple with the CreatedByUserfirstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035Notebooks) GetCreatedByUserfirstnameOk() (*string, bool) {
	if o == nil || o.CreatedByUserfirstname == nil {
		return nil, false
	}
	return o.CreatedByUserfirstname, true
}

// HasCreatedByUserfirstname returns a boolean if a field has been set.
func (o *InlineResponse20035Notebooks) HasCreatedByUserfirstname() bool {
	if o != nil && o.CreatedByUserfirstname != nil {
		return true
	}

	return false
}

// SetCreatedByUserfirstname gets a reference to the given string and assigns it to the CreatedByUserfirstname field.
func (o *InlineResponse20035Notebooks) SetCreatedByUserfirstname(v string) {
	o.CreatedByUserfirstname = &v
}

// GetCreatedByUserlastname returns the CreatedByUserlastname field value if set, zero value otherwise.
func (o *InlineResponse20035Notebooks) GetCreatedByUserlastname() string {
	if o == nil || o.CreatedByUserlastname == nil {
		var ret string
		return ret
	}
	return *o.CreatedByUserlastname
}

// GetCreatedByUserlastnameOk returns a tuple with the CreatedByUserlastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035Notebooks) GetCreatedByUserlastnameOk() (*string, bool) {
	if o == nil || o.CreatedByUserlastname == nil {
		return nil, false
	}
	return o.CreatedByUserlastname, true
}

// HasCreatedByUserlastname returns a boolean if a field has been set.
func (o *InlineResponse20035Notebooks) HasCreatedByUserlastname() bool {
	if o != nil && o.CreatedByUserlastname != nil {
		return true
	}

	return false
}

// SetCreatedByUserlastname gets a reference to the given string and assigns it to the CreatedByUserlastname field.
func (o *InlineResponse20035Notebooks) SetCreatedByUserlastname(v string) {
	o.CreatedByUserlastname = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *InlineResponse20035Notebooks) GetCreatedDate() string {
	if o == nil || o.CreatedDate == nil {
		var ret string
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035Notebooks) GetCreatedDateOk() (*string, bool) {
	if o == nil || o.CreatedDate == nil {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *InlineResponse20035Notebooks) HasCreatedDate() bool {
	if o != nil && o.CreatedDate != nil {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given string and assigns it to the CreatedDate field.
func (o *InlineResponse20035Notebooks) SetCreatedDate(v string) {
	o.CreatedDate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineResponse20035Notebooks) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035Notebooks) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineResponse20035Notebooks) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineResponse20035Notebooks) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20035Notebooks) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035Notebooks) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20035Notebooks) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20035Notebooks) SetId(v string) {
	o.Id = &v
}

// GetLockdownId returns the LockdownId field value if set, zero value otherwise.
func (o *InlineResponse20035Notebooks) GetLockdownId() string {
	if o == nil || o.LockdownId == nil {
		var ret string
		return ret
	}
	return *o.LockdownId
}

// GetLockdownIdOk returns a tuple with the LockdownId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035Notebooks) GetLockdownIdOk() (*string, bool) {
	if o == nil || o.LockdownId == nil {
		return nil, false
	}
	return o.LockdownId, true
}

// HasLockdownId returns a boolean if a field has been set.
func (o *InlineResponse20035Notebooks) HasLockdownId() bool {
	if o != nil && o.LockdownId != nil {
		return true
	}

	return false
}

// SetLockdownId gets a reference to the given string and assigns it to the LockdownId field.
func (o *InlineResponse20035Notebooks) SetLockdownId(v string) {
	o.LockdownId = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *InlineResponse20035Notebooks) GetLocked() string {
	if o == nil || o.Locked == nil {
		var ret string
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035Notebooks) GetLockedOk() (*string, bool) {
	if o == nil || o.Locked == nil {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *InlineResponse20035Notebooks) HasLocked() bool {
	if o != nil && o.Locked != nil {
		return true
	}

	return false
}

// SetLocked gets a reference to the given string and assigns it to the Locked field.
func (o *InlineResponse20035Notebooks) SetLocked(v string) {
	o.Locked = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20035Notebooks) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035Notebooks) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20035Notebooks) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20035Notebooks) SetName(v string) {
	o.Name = &v
}

// GetNrComments returns the NrComments field value if set, zero value otherwise.
func (o *InlineResponse20035Notebooks) GetNrComments() string {
	if o == nil || o.NrComments == nil {
		var ret string
		return ret
	}
	return *o.NrComments
}

// GetNrCommentsOk returns a tuple with the NrComments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035Notebooks) GetNrCommentsOk() (*string, bool) {
	if o == nil || o.NrComments == nil {
		return nil, false
	}
	return o.NrComments, true
}

// HasNrComments returns a boolean if a field has been set.
func (o *InlineResponse20035Notebooks) HasNrComments() bool {
	if o != nil && o.NrComments != nil {
		return true
	}

	return false
}

// SetNrComments gets a reference to the given string and assigns it to the NrComments field.
func (o *InlineResponse20035Notebooks) SetNrComments(v string) {
	o.NrComments = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *InlineResponse20035Notebooks) GetPrivate() string {
	if o == nil || o.Private == nil {
		var ret string
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035Notebooks) GetPrivateOk() (*string, bool) {
	if o == nil || o.Private == nil {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *InlineResponse20035Notebooks) HasPrivate() bool {
	if o != nil && o.Private != nil {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given string and assigns it to the Private field.
func (o *InlineResponse20035Notebooks) SetPrivate(v string) {
	o.Private = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *InlineResponse20035Notebooks) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035Notebooks) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *InlineResponse20035Notebooks) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *InlineResponse20035Notebooks) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetReadCommentsCount returns the ReadCommentsCount field value if set, zero value otherwise.
func (o *InlineResponse20035Notebooks) GetReadCommentsCount() string {
	if o == nil || o.ReadCommentsCount == nil {
		var ret string
		return ret
	}
	return *o.ReadCommentsCount
}

// GetReadCommentsCountOk returns a tuple with the ReadCommentsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035Notebooks) GetReadCommentsCountOk() (*string, bool) {
	if o == nil || o.ReadCommentsCount == nil {
		return nil, false
	}
	return o.ReadCommentsCount, true
}

// HasReadCommentsCount returns a boolean if a field has been set.
func (o *InlineResponse20035Notebooks) HasReadCommentsCount() bool {
	if o != nil && o.ReadCommentsCount != nil {
		return true
	}

	return false
}

// SetReadCommentsCount gets a reference to the given string and assigns it to the ReadCommentsCount field.
func (o *InlineResponse20035Notebooks) SetReadCommentsCount(v string) {
	o.ReadCommentsCount = &v
}

// GetUpdatedByUserId returns the UpdatedByUserId field value if set, zero value otherwise.
func (o *InlineResponse20035Notebooks) GetUpdatedByUserId() string {
	if o == nil || o.UpdatedByUserId == nil {
		var ret string
		return ret
	}
	return *o.UpdatedByUserId
}

// GetUpdatedByUserIdOk returns a tuple with the UpdatedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035Notebooks) GetUpdatedByUserIdOk() (*string, bool) {
	if o == nil || o.UpdatedByUserId == nil {
		return nil, false
	}
	return o.UpdatedByUserId, true
}

// HasUpdatedByUserId returns a boolean if a field has been set.
func (o *InlineResponse20035Notebooks) HasUpdatedByUserId() bool {
	if o != nil && o.UpdatedByUserId != nil {
		return true
	}

	return false
}

// SetUpdatedByUserId gets a reference to the given string and assigns it to the UpdatedByUserId field.
func (o *InlineResponse20035Notebooks) SetUpdatedByUserId(v string) {
	o.UpdatedByUserId = &v
}

// GetUpdatedByUserfirstname returns the UpdatedByUserfirstname field value if set, zero value otherwise.
func (o *InlineResponse20035Notebooks) GetUpdatedByUserfirstname() string {
	if o == nil || o.UpdatedByUserfirstname == nil {
		var ret string
		return ret
	}
	return *o.UpdatedByUserfirstname
}

// GetUpdatedByUserfirstnameOk returns a tuple with the UpdatedByUserfirstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035Notebooks) GetUpdatedByUserfirstnameOk() (*string, bool) {
	if o == nil || o.UpdatedByUserfirstname == nil {
		return nil, false
	}
	return o.UpdatedByUserfirstname, true
}

// HasUpdatedByUserfirstname returns a boolean if a field has been set.
func (o *InlineResponse20035Notebooks) HasUpdatedByUserfirstname() bool {
	if o != nil && o.UpdatedByUserfirstname != nil {
		return true
	}

	return false
}

// SetUpdatedByUserfirstname gets a reference to the given string and assigns it to the UpdatedByUserfirstname field.
func (o *InlineResponse20035Notebooks) SetUpdatedByUserfirstname(v string) {
	o.UpdatedByUserfirstname = &v
}

// GetUpdatedByUserlastname returns the UpdatedByUserlastname field value if set, zero value otherwise.
func (o *InlineResponse20035Notebooks) GetUpdatedByUserlastname() string {
	if o == nil || o.UpdatedByUserlastname == nil {
		var ret string
		return ret
	}
	return *o.UpdatedByUserlastname
}

// GetUpdatedByUserlastnameOk returns a tuple with the UpdatedByUserlastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035Notebooks) GetUpdatedByUserlastnameOk() (*string, bool) {
	if o == nil || o.UpdatedByUserlastname == nil {
		return nil, false
	}
	return o.UpdatedByUserlastname, true
}

// HasUpdatedByUserlastname returns a boolean if a field has been set.
func (o *InlineResponse20035Notebooks) HasUpdatedByUserlastname() bool {
	if o != nil && o.UpdatedByUserlastname != nil {
		return true
	}

	return false
}

// SetUpdatedByUserlastname gets a reference to the given string and assigns it to the UpdatedByUserlastname field.
func (o *InlineResponse20035Notebooks) SetUpdatedByUserlastname(v string) {
	o.UpdatedByUserlastname = &v
}

// GetUpdatedDate returns the UpdatedDate field value if set, zero value otherwise.
func (o *InlineResponse20035Notebooks) GetUpdatedDate() string {
	if o == nil || o.UpdatedDate == nil {
		var ret string
		return ret
	}
	return *o.UpdatedDate
}

// GetUpdatedDateOk returns a tuple with the UpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035Notebooks) GetUpdatedDateOk() (*string, bool) {
	if o == nil || o.UpdatedDate == nil {
		return nil, false
	}
	return o.UpdatedDate, true
}

// HasUpdatedDate returns a boolean if a field has been set.
func (o *InlineResponse20035Notebooks) HasUpdatedDate() bool {
	if o != nil && o.UpdatedDate != nil {
		return true
	}

	return false
}

// SetUpdatedDate gets a reference to the given string and assigns it to the UpdatedDate field.
func (o *InlineResponse20035Notebooks) SetUpdatedDate(v string) {
	o.UpdatedDate = &v
}

// GetUserDisplayUpdatedDate returns the UserDisplayUpdatedDate field value if set, zero value otherwise.
func (o *InlineResponse20035Notebooks) GetUserDisplayUpdatedDate() string {
	if o == nil || o.UserDisplayUpdatedDate == nil {
		var ret string
		return ret
	}
	return *o.UserDisplayUpdatedDate
}

// GetUserDisplayUpdatedDateOk returns a tuple with the UserDisplayUpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035Notebooks) GetUserDisplayUpdatedDateOk() (*string, bool) {
	if o == nil || o.UserDisplayUpdatedDate == nil {
		return nil, false
	}
	return o.UserDisplayUpdatedDate, true
}

// HasUserDisplayUpdatedDate returns a boolean if a field has been set.
func (o *InlineResponse20035Notebooks) HasUserDisplayUpdatedDate() bool {
	if o != nil && o.UserDisplayUpdatedDate != nil {
		return true
	}

	return false
}

// SetUserDisplayUpdatedDate gets a reference to the given string and assigns it to the UserDisplayUpdatedDate field.
func (o *InlineResponse20035Notebooks) SetUserDisplayUpdatedDate(v string) {
	o.UserDisplayUpdatedDate = &v
}

// GetUserDisplayUpdatedTime returns the UserDisplayUpdatedTime field value if set, zero value otherwise.
func (o *InlineResponse20035Notebooks) GetUserDisplayUpdatedTime() string {
	if o == nil || o.UserDisplayUpdatedTime == nil {
		var ret string
		return ret
	}
	return *o.UserDisplayUpdatedTime
}

// GetUserDisplayUpdatedTimeOk returns a tuple with the UserDisplayUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035Notebooks) GetUserDisplayUpdatedTimeOk() (*string, bool) {
	if o == nil || o.UserDisplayUpdatedTime == nil {
		return nil, false
	}
	return o.UserDisplayUpdatedTime, true
}

// HasUserDisplayUpdatedTime returns a boolean if a field has been set.
func (o *InlineResponse20035Notebooks) HasUserDisplayUpdatedTime() bool {
	if o != nil && o.UserDisplayUpdatedTime != nil {
		return true
	}

	return false
}

// SetUserDisplayUpdatedTime gets a reference to the given string and assigns it to the UserDisplayUpdatedTime field.
func (o *InlineResponse20035Notebooks) SetUserDisplayUpdatedTime(v string) {
	o.UserDisplayUpdatedTime = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *InlineResponse20035Notebooks) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035Notebooks) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *InlineResponse20035Notebooks) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *InlineResponse20035Notebooks) SetVersion(v string) {
	o.Version = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *InlineResponse20035Notebooks) GetVersionId() string {
	if o == nil || o.VersionId == nil {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035Notebooks) GetVersionIdOk() (*string, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *InlineResponse20035Notebooks) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *InlineResponse20035Notebooks) SetVersionId(v string) {
	o.VersionId = &v
}

func (o InlineResponse20035Notebooks) MarshalJSON() ([]byte, error) {
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
	if o.CreatedByUserId != nil {
		toSerialize["created-by-userId"] = o.CreatedByUserId
	}
	if o.CreatedByUserfirstname != nil {
		toSerialize["created-by-userfirstname"] = o.CreatedByUserfirstname
	}
	if o.CreatedByUserlastname != nil {
		toSerialize["created-by-userlastname"] = o.CreatedByUserlastname
	}
	if o.CreatedDate != nil {
		toSerialize["created-date"] = o.CreatedDate
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LockdownId != nil {
		toSerialize["lockdown-id"] = o.LockdownId
	}
	if o.Locked != nil {
		toSerialize["locked"] = o.Locked
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NrComments != nil {
		toSerialize["nr-comments"] = o.NrComments
	}
	if o.Private != nil {
		toSerialize["private"] = o.Private
	}
	if o.ProjectId != nil {
		toSerialize["project-id"] = o.ProjectId
	}
	if o.ReadCommentsCount != nil {
		toSerialize["read-comments-count"] = o.ReadCommentsCount
	}
	if o.UpdatedByUserId != nil {
		toSerialize["updated-by-userId"] = o.UpdatedByUserId
	}
	if o.UpdatedByUserfirstname != nil {
		toSerialize["updated-by-userfirstname"] = o.UpdatedByUserfirstname
	}
	if o.UpdatedByUserlastname != nil {
		toSerialize["updated-by-userlastname"] = o.UpdatedByUserlastname
	}
	if o.UpdatedDate != nil {
		toSerialize["updated-date"] = o.UpdatedDate
	}
	if o.UserDisplayUpdatedDate != nil {
		toSerialize["user-display-updated-date"] = o.UserDisplayUpdatedDate
	}
	if o.UserDisplayUpdatedTime != nil {
		toSerialize["user-display-updated-time"] = o.UserDisplayUpdatedTime
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.VersionId != nil {
		toSerialize["version-id"] = o.VersionId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20035Notebooks struct {
	value *InlineResponse20035Notebooks
	isSet bool
}

func (v NullableInlineResponse20035Notebooks) Get() *InlineResponse20035Notebooks {
	return v.value
}

func (v *NullableInlineResponse20035Notebooks) Set(val *InlineResponse20035Notebooks) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20035Notebooks) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20035Notebooks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20035Notebooks(val *InlineResponse20035Notebooks) *NullableInlineResponse20035Notebooks {
	return &NullableInlineResponse20035Notebooks{value: val, isSet: true}
}

func (v NullableInlineResponse20035Notebooks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20035Notebooks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


