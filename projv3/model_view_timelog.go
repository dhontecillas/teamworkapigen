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

// ViewTimelog Timelog contains all the information returned from a timelog.
type ViewTimelog struct {
	Billable *bool `json:"billable,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	DateCreated *string `json:"dateCreated,omitempty"`
	DateDeleted *string `json:"dateDeleted,omitempty"`
	DateEdited *string `json:"dateEdited,omitempty"`
	Deleted *bool `json:"deleted,omitempty"`
	DeletedAt *string `json:"deletedAt,omitempty"`
	DeletedBy *int32 `json:"deletedBy,omitempty"`
	DeletedByUserId *int32 `json:"deletedByUserId,omitempty"`
	Description *string `json:"description,omitempty"`
	DeskTicketId *int32 `json:"deskTicketId,omitempty"`
	EditedByUserId *int32 `json:"editedByUserId,omitempty"`
	HasStartTime *bool `json:"hasStartTime,omitempty"`
	Id *int32 `json:"id,omitempty"`
	InvoiceNumber *string `json:"invoiceNumber,omitempty"`
	LoggedBy *int32 `json:"loggedBy,omitempty"`
	LoggedByUserId *int32 `json:"loggedByUserId,omitempty"`
	Minutes *int32 `json:"minutes,omitempty"`
	Project *ViewRelationship `json:"project,omitempty"`
	ProjectBillingInvoice *ViewRelationship `json:"projectBillingInvoice,omitempty"`
	ProjectBillingInvoiceId *int32 `json:"projectBillingInvoiceId,omitempty"`
	ProjectId *int32 `json:"projectId,omitempty"`
	TagIds *[]int32 `json:"tagIds,omitempty"`
	Tags *[]ViewRelationship `json:"tags,omitempty"`
	Task *ViewRelationship `json:"task,omitempty"`
	TaskId *int32 `json:"taskId,omitempty"`
	TaskIdPreMove *int32 `json:"taskIdPreMove,omitempty"`
	TaskPreMove *ViewRelationship `json:"taskPreMove,omitempty"`
	TimeLogged *string `json:"timeLogged,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	UpdatedBy *int32 `json:"updatedBy,omitempty"`
	User *ViewRelationship `json:"user,omitempty"`
	UserId *int32 `json:"userId,omitempty"`
}

// NewViewTimelog instantiates a new ViewTimelog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewTimelog() *ViewTimelog {
	this := ViewTimelog{}
	return &this
}

// NewViewTimelogWithDefaults instantiates a new ViewTimelog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewTimelogWithDefaults() *ViewTimelog {
	this := ViewTimelog{}
	return &this
}

// GetBillable returns the Billable field value if set, zero value otherwise.
func (o *ViewTimelog) GetBillable() bool {
	if o == nil || o.Billable == nil {
		var ret bool
		return ret
	}
	return *o.Billable
}

// GetBillableOk returns a tuple with the Billable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetBillableOk() (*bool, bool) {
	if o == nil || o.Billable == nil {
		return nil, false
	}
	return o.Billable, true
}

// HasBillable returns a boolean if a field has been set.
func (o *ViewTimelog) HasBillable() bool {
	if o != nil && o.Billable != nil {
		return true
	}

	return false
}

// SetBillable gets a reference to the given bool and assigns it to the Billable field.
func (o *ViewTimelog) SetBillable(v bool) {
	o.Billable = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ViewTimelog) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ViewTimelog) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ViewTimelog) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *ViewTimelog) GetDateCreated() string {
	if o == nil || o.DateCreated == nil {
		var ret string
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetDateCreatedOk() (*string, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *ViewTimelog) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given string and assigns it to the DateCreated field.
func (o *ViewTimelog) SetDateCreated(v string) {
	o.DateCreated = &v
}

// GetDateDeleted returns the DateDeleted field value if set, zero value otherwise.
func (o *ViewTimelog) GetDateDeleted() string {
	if o == nil || o.DateDeleted == nil {
		var ret string
		return ret
	}
	return *o.DateDeleted
}

// GetDateDeletedOk returns a tuple with the DateDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetDateDeletedOk() (*string, bool) {
	if o == nil || o.DateDeleted == nil {
		return nil, false
	}
	return o.DateDeleted, true
}

// HasDateDeleted returns a boolean if a field has been set.
func (o *ViewTimelog) HasDateDeleted() bool {
	if o != nil && o.DateDeleted != nil {
		return true
	}

	return false
}

// SetDateDeleted gets a reference to the given string and assigns it to the DateDeleted field.
func (o *ViewTimelog) SetDateDeleted(v string) {
	o.DateDeleted = &v
}

// GetDateEdited returns the DateEdited field value if set, zero value otherwise.
func (o *ViewTimelog) GetDateEdited() string {
	if o == nil || o.DateEdited == nil {
		var ret string
		return ret
	}
	return *o.DateEdited
}

// GetDateEditedOk returns a tuple with the DateEdited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetDateEditedOk() (*string, bool) {
	if o == nil || o.DateEdited == nil {
		return nil, false
	}
	return o.DateEdited, true
}

// HasDateEdited returns a boolean if a field has been set.
func (o *ViewTimelog) HasDateEdited() bool {
	if o != nil && o.DateEdited != nil {
		return true
	}

	return false
}

// SetDateEdited gets a reference to the given string and assigns it to the DateEdited field.
func (o *ViewTimelog) SetDateEdited(v string) {
	o.DateEdited = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *ViewTimelog) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *ViewTimelog) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *ViewTimelog) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *ViewTimelog) GetDeletedAt() string {
	if o == nil || o.DeletedAt == nil {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetDeletedAtOk() (*string, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ViewTimelog) HasDeletedAt() bool {
	if o != nil && o.DeletedAt != nil {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *ViewTimelog) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetDeletedBy returns the DeletedBy field value if set, zero value otherwise.
func (o *ViewTimelog) GetDeletedBy() int32 {
	if o == nil || o.DeletedBy == nil {
		var ret int32
		return ret
	}
	return *o.DeletedBy
}

// GetDeletedByOk returns a tuple with the DeletedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetDeletedByOk() (*int32, bool) {
	if o == nil || o.DeletedBy == nil {
		return nil, false
	}
	return o.DeletedBy, true
}

// HasDeletedBy returns a boolean if a field has been set.
func (o *ViewTimelog) HasDeletedBy() bool {
	if o != nil && o.DeletedBy != nil {
		return true
	}

	return false
}

// SetDeletedBy gets a reference to the given int32 and assigns it to the DeletedBy field.
func (o *ViewTimelog) SetDeletedBy(v int32) {
	o.DeletedBy = &v
}

// GetDeletedByUserId returns the DeletedByUserId field value if set, zero value otherwise.
func (o *ViewTimelog) GetDeletedByUserId() int32 {
	if o == nil || o.DeletedByUserId == nil {
		var ret int32
		return ret
	}
	return *o.DeletedByUserId
}

// GetDeletedByUserIdOk returns a tuple with the DeletedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetDeletedByUserIdOk() (*int32, bool) {
	if o == nil || o.DeletedByUserId == nil {
		return nil, false
	}
	return o.DeletedByUserId, true
}

// HasDeletedByUserId returns a boolean if a field has been set.
func (o *ViewTimelog) HasDeletedByUserId() bool {
	if o != nil && o.DeletedByUserId != nil {
		return true
	}

	return false
}

// SetDeletedByUserId gets a reference to the given int32 and assigns it to the DeletedByUserId field.
func (o *ViewTimelog) SetDeletedByUserId(v int32) {
	o.DeletedByUserId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ViewTimelog) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ViewTimelog) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ViewTimelog) SetDescription(v string) {
	o.Description = &v
}

// GetDeskTicketId returns the DeskTicketId field value if set, zero value otherwise.
func (o *ViewTimelog) GetDeskTicketId() int32 {
	if o == nil || o.DeskTicketId == nil {
		var ret int32
		return ret
	}
	return *o.DeskTicketId
}

// GetDeskTicketIdOk returns a tuple with the DeskTicketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetDeskTicketIdOk() (*int32, bool) {
	if o == nil || o.DeskTicketId == nil {
		return nil, false
	}
	return o.DeskTicketId, true
}

// HasDeskTicketId returns a boolean if a field has been set.
func (o *ViewTimelog) HasDeskTicketId() bool {
	if o != nil && o.DeskTicketId != nil {
		return true
	}

	return false
}

// SetDeskTicketId gets a reference to the given int32 and assigns it to the DeskTicketId field.
func (o *ViewTimelog) SetDeskTicketId(v int32) {
	o.DeskTicketId = &v
}

// GetEditedByUserId returns the EditedByUserId field value if set, zero value otherwise.
func (o *ViewTimelog) GetEditedByUserId() int32 {
	if o == nil || o.EditedByUserId == nil {
		var ret int32
		return ret
	}
	return *o.EditedByUserId
}

// GetEditedByUserIdOk returns a tuple with the EditedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetEditedByUserIdOk() (*int32, bool) {
	if o == nil || o.EditedByUserId == nil {
		return nil, false
	}
	return o.EditedByUserId, true
}

// HasEditedByUserId returns a boolean if a field has been set.
func (o *ViewTimelog) HasEditedByUserId() bool {
	if o != nil && o.EditedByUserId != nil {
		return true
	}

	return false
}

// SetEditedByUserId gets a reference to the given int32 and assigns it to the EditedByUserId field.
func (o *ViewTimelog) SetEditedByUserId(v int32) {
	o.EditedByUserId = &v
}

// GetHasStartTime returns the HasStartTime field value if set, zero value otherwise.
func (o *ViewTimelog) GetHasStartTime() bool {
	if o == nil || o.HasStartTime == nil {
		var ret bool
		return ret
	}
	return *o.HasStartTime
}

// GetHasStartTimeOk returns a tuple with the HasStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetHasStartTimeOk() (*bool, bool) {
	if o == nil || o.HasStartTime == nil {
		return nil, false
	}
	return o.HasStartTime, true
}

// HasHasStartTime returns a boolean if a field has been set.
func (o *ViewTimelog) HasHasStartTime() bool {
	if o != nil && o.HasStartTime != nil {
		return true
	}

	return false
}

// SetHasStartTime gets a reference to the given bool and assigns it to the HasStartTime field.
func (o *ViewTimelog) SetHasStartTime(v bool) {
	o.HasStartTime = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ViewTimelog) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ViewTimelog) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ViewTimelog) SetId(v int32) {
	o.Id = &v
}

// GetInvoiceNumber returns the InvoiceNumber field value if set, zero value otherwise.
func (o *ViewTimelog) GetInvoiceNumber() string {
	if o == nil || o.InvoiceNumber == nil {
		var ret string
		return ret
	}
	return *o.InvoiceNumber
}

// GetInvoiceNumberOk returns a tuple with the InvoiceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetInvoiceNumberOk() (*string, bool) {
	if o == nil || o.InvoiceNumber == nil {
		return nil, false
	}
	return o.InvoiceNumber, true
}

// HasInvoiceNumber returns a boolean if a field has been set.
func (o *ViewTimelog) HasInvoiceNumber() bool {
	if o != nil && o.InvoiceNumber != nil {
		return true
	}

	return false
}

// SetInvoiceNumber gets a reference to the given string and assigns it to the InvoiceNumber field.
func (o *ViewTimelog) SetInvoiceNumber(v string) {
	o.InvoiceNumber = &v
}

// GetLoggedBy returns the LoggedBy field value if set, zero value otherwise.
func (o *ViewTimelog) GetLoggedBy() int32 {
	if o == nil || o.LoggedBy == nil {
		var ret int32
		return ret
	}
	return *o.LoggedBy
}

// GetLoggedByOk returns a tuple with the LoggedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetLoggedByOk() (*int32, bool) {
	if o == nil || o.LoggedBy == nil {
		return nil, false
	}
	return o.LoggedBy, true
}

// HasLoggedBy returns a boolean if a field has been set.
func (o *ViewTimelog) HasLoggedBy() bool {
	if o != nil && o.LoggedBy != nil {
		return true
	}

	return false
}

// SetLoggedBy gets a reference to the given int32 and assigns it to the LoggedBy field.
func (o *ViewTimelog) SetLoggedBy(v int32) {
	o.LoggedBy = &v
}

// GetLoggedByUserId returns the LoggedByUserId field value if set, zero value otherwise.
func (o *ViewTimelog) GetLoggedByUserId() int32 {
	if o == nil || o.LoggedByUserId == nil {
		var ret int32
		return ret
	}
	return *o.LoggedByUserId
}

// GetLoggedByUserIdOk returns a tuple with the LoggedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetLoggedByUserIdOk() (*int32, bool) {
	if o == nil || o.LoggedByUserId == nil {
		return nil, false
	}
	return o.LoggedByUserId, true
}

// HasLoggedByUserId returns a boolean if a field has been set.
func (o *ViewTimelog) HasLoggedByUserId() bool {
	if o != nil && o.LoggedByUserId != nil {
		return true
	}

	return false
}

// SetLoggedByUserId gets a reference to the given int32 and assigns it to the LoggedByUserId field.
func (o *ViewTimelog) SetLoggedByUserId(v int32) {
	o.LoggedByUserId = &v
}

// GetMinutes returns the Minutes field value if set, zero value otherwise.
func (o *ViewTimelog) GetMinutes() int32 {
	if o == nil || o.Minutes == nil {
		var ret int32
		return ret
	}
	return *o.Minutes
}

// GetMinutesOk returns a tuple with the Minutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetMinutesOk() (*int32, bool) {
	if o == nil || o.Minutes == nil {
		return nil, false
	}
	return o.Minutes, true
}

// HasMinutes returns a boolean if a field has been set.
func (o *ViewTimelog) HasMinutes() bool {
	if o != nil && o.Minutes != nil {
		return true
	}

	return false
}

// SetMinutes gets a reference to the given int32 and assigns it to the Minutes field.
func (o *ViewTimelog) SetMinutes(v int32) {
	o.Minutes = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *ViewTimelog) GetProject() ViewRelationship {
	if o == nil || o.Project == nil {
		var ret ViewRelationship
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetProjectOk() (*ViewRelationship, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *ViewTimelog) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given ViewRelationship and assigns it to the Project field.
func (o *ViewTimelog) SetProject(v ViewRelationship) {
	o.Project = &v
}

// GetProjectBillingInvoice returns the ProjectBillingInvoice field value if set, zero value otherwise.
func (o *ViewTimelog) GetProjectBillingInvoice() ViewRelationship {
	if o == nil || o.ProjectBillingInvoice == nil {
		var ret ViewRelationship
		return ret
	}
	return *o.ProjectBillingInvoice
}

// GetProjectBillingInvoiceOk returns a tuple with the ProjectBillingInvoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetProjectBillingInvoiceOk() (*ViewRelationship, bool) {
	if o == nil || o.ProjectBillingInvoice == nil {
		return nil, false
	}
	return o.ProjectBillingInvoice, true
}

// HasProjectBillingInvoice returns a boolean if a field has been set.
func (o *ViewTimelog) HasProjectBillingInvoice() bool {
	if o != nil && o.ProjectBillingInvoice != nil {
		return true
	}

	return false
}

// SetProjectBillingInvoice gets a reference to the given ViewRelationship and assigns it to the ProjectBillingInvoice field.
func (o *ViewTimelog) SetProjectBillingInvoice(v ViewRelationship) {
	o.ProjectBillingInvoice = &v
}

// GetProjectBillingInvoiceId returns the ProjectBillingInvoiceId field value if set, zero value otherwise.
func (o *ViewTimelog) GetProjectBillingInvoiceId() int32 {
	if o == nil || o.ProjectBillingInvoiceId == nil {
		var ret int32
		return ret
	}
	return *o.ProjectBillingInvoiceId
}

// GetProjectBillingInvoiceIdOk returns a tuple with the ProjectBillingInvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetProjectBillingInvoiceIdOk() (*int32, bool) {
	if o == nil || o.ProjectBillingInvoiceId == nil {
		return nil, false
	}
	return o.ProjectBillingInvoiceId, true
}

// HasProjectBillingInvoiceId returns a boolean if a field has been set.
func (o *ViewTimelog) HasProjectBillingInvoiceId() bool {
	if o != nil && o.ProjectBillingInvoiceId != nil {
		return true
	}

	return false
}

// SetProjectBillingInvoiceId gets a reference to the given int32 and assigns it to the ProjectBillingInvoiceId field.
func (o *ViewTimelog) SetProjectBillingInvoiceId(v int32) {
	o.ProjectBillingInvoiceId = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *ViewTimelog) GetProjectId() int32 {
	if o == nil || o.ProjectId == nil {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetProjectIdOk() (*int32, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ViewTimelog) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *ViewTimelog) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetTagIds returns the TagIds field value if set, zero value otherwise.
func (o *ViewTimelog) GetTagIds() []int32 {
	if o == nil || o.TagIds == nil {
		var ret []int32
		return ret
	}
	return *o.TagIds
}

// GetTagIdsOk returns a tuple with the TagIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetTagIdsOk() (*[]int32, bool) {
	if o == nil || o.TagIds == nil {
		return nil, false
	}
	return o.TagIds, true
}

// HasTagIds returns a boolean if a field has been set.
func (o *ViewTimelog) HasTagIds() bool {
	if o != nil && o.TagIds != nil {
		return true
	}

	return false
}

// SetTagIds gets a reference to the given []int32 and assigns it to the TagIds field.
func (o *ViewTimelog) SetTagIds(v []int32) {
	o.TagIds = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ViewTimelog) GetTags() []ViewRelationship {
	if o == nil || o.Tags == nil {
		var ret []ViewRelationship
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetTagsOk() (*[]ViewRelationship, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ViewTimelog) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ViewRelationship and assigns it to the Tags field.
func (o *ViewTimelog) SetTags(v []ViewRelationship) {
	o.Tags = &v
}

// GetTask returns the Task field value if set, zero value otherwise.
func (o *ViewTimelog) GetTask() ViewRelationship {
	if o == nil || o.Task == nil {
		var ret ViewRelationship
		return ret
	}
	return *o.Task
}

// GetTaskOk returns a tuple with the Task field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetTaskOk() (*ViewRelationship, bool) {
	if o == nil || o.Task == nil {
		return nil, false
	}
	return o.Task, true
}

// HasTask returns a boolean if a field has been set.
func (o *ViewTimelog) HasTask() bool {
	if o != nil && o.Task != nil {
		return true
	}

	return false
}

// SetTask gets a reference to the given ViewRelationship and assigns it to the Task field.
func (o *ViewTimelog) SetTask(v ViewRelationship) {
	o.Task = &v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *ViewTimelog) GetTaskId() int32 {
	if o == nil || o.TaskId == nil {
		var ret int32
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetTaskIdOk() (*int32, bool) {
	if o == nil || o.TaskId == nil {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *ViewTimelog) HasTaskId() bool {
	if o != nil && o.TaskId != nil {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given int32 and assigns it to the TaskId field.
func (o *ViewTimelog) SetTaskId(v int32) {
	o.TaskId = &v
}

// GetTaskIdPreMove returns the TaskIdPreMove field value if set, zero value otherwise.
func (o *ViewTimelog) GetTaskIdPreMove() int32 {
	if o == nil || o.TaskIdPreMove == nil {
		var ret int32
		return ret
	}
	return *o.TaskIdPreMove
}

// GetTaskIdPreMoveOk returns a tuple with the TaskIdPreMove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetTaskIdPreMoveOk() (*int32, bool) {
	if o == nil || o.TaskIdPreMove == nil {
		return nil, false
	}
	return o.TaskIdPreMove, true
}

// HasTaskIdPreMove returns a boolean if a field has been set.
func (o *ViewTimelog) HasTaskIdPreMove() bool {
	if o != nil && o.TaskIdPreMove != nil {
		return true
	}

	return false
}

// SetTaskIdPreMove gets a reference to the given int32 and assigns it to the TaskIdPreMove field.
func (o *ViewTimelog) SetTaskIdPreMove(v int32) {
	o.TaskIdPreMove = &v
}

// GetTaskPreMove returns the TaskPreMove field value if set, zero value otherwise.
func (o *ViewTimelog) GetTaskPreMove() ViewRelationship {
	if o == nil || o.TaskPreMove == nil {
		var ret ViewRelationship
		return ret
	}
	return *o.TaskPreMove
}

// GetTaskPreMoveOk returns a tuple with the TaskPreMove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetTaskPreMoveOk() (*ViewRelationship, bool) {
	if o == nil || o.TaskPreMove == nil {
		return nil, false
	}
	return o.TaskPreMove, true
}

// HasTaskPreMove returns a boolean if a field has been set.
func (o *ViewTimelog) HasTaskPreMove() bool {
	if o != nil && o.TaskPreMove != nil {
		return true
	}

	return false
}

// SetTaskPreMove gets a reference to the given ViewRelationship and assigns it to the TaskPreMove field.
func (o *ViewTimelog) SetTaskPreMove(v ViewRelationship) {
	o.TaskPreMove = &v
}

// GetTimeLogged returns the TimeLogged field value if set, zero value otherwise.
func (o *ViewTimelog) GetTimeLogged() string {
	if o == nil || o.TimeLogged == nil {
		var ret string
		return ret
	}
	return *o.TimeLogged
}

// GetTimeLoggedOk returns a tuple with the TimeLogged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetTimeLoggedOk() (*string, bool) {
	if o == nil || o.TimeLogged == nil {
		return nil, false
	}
	return o.TimeLogged, true
}

// HasTimeLogged returns a boolean if a field has been set.
func (o *ViewTimelog) HasTimeLogged() bool {
	if o != nil && o.TimeLogged != nil {
		return true
	}

	return false
}

// SetTimeLogged gets a reference to the given string and assigns it to the TimeLogged field.
func (o *ViewTimelog) SetTimeLogged(v string) {
	o.TimeLogged = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ViewTimelog) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ViewTimelog) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ViewTimelog) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *ViewTimelog) GetUpdatedBy() int32 {
	if o == nil || o.UpdatedBy == nil {
		var ret int32
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetUpdatedByOk() (*int32, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *ViewTimelog) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given int32 and assigns it to the UpdatedBy field.
func (o *ViewTimelog) SetUpdatedBy(v int32) {
	o.UpdatedBy = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ViewTimelog) GetUser() ViewRelationship {
	if o == nil || o.User == nil {
		var ret ViewRelationship
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetUserOk() (*ViewRelationship, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ViewTimelog) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given ViewRelationship and assigns it to the User field.
func (o *ViewTimelog) SetUser(v ViewRelationship) {
	o.User = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ViewTimelog) GetUserId() int32 {
	if o == nil || o.UserId == nil {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewTimelog) GetUserIdOk() (*int32, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ViewTimelog) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *ViewTimelog) SetUserId(v int32) {
	o.UserId = &v
}

func (o ViewTimelog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Billable != nil {
		toSerialize["billable"] = o.Billable
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.DateCreated != nil {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if o.DateDeleted != nil {
		toSerialize["dateDeleted"] = o.DateDeleted
	}
	if o.DateEdited != nil {
		toSerialize["dateEdited"] = o.DateEdited
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	if o.DeletedAt != nil {
		toSerialize["deletedAt"] = o.DeletedAt
	}
	if o.DeletedBy != nil {
		toSerialize["deletedBy"] = o.DeletedBy
	}
	if o.DeletedByUserId != nil {
		toSerialize["deletedByUserId"] = o.DeletedByUserId
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DeskTicketId != nil {
		toSerialize["deskTicketId"] = o.DeskTicketId
	}
	if o.EditedByUserId != nil {
		toSerialize["editedByUserId"] = o.EditedByUserId
	}
	if o.HasStartTime != nil {
		toSerialize["hasStartTime"] = o.HasStartTime
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.InvoiceNumber != nil {
		toSerialize["invoiceNumber"] = o.InvoiceNumber
	}
	if o.LoggedBy != nil {
		toSerialize["loggedBy"] = o.LoggedBy
	}
	if o.LoggedByUserId != nil {
		toSerialize["loggedByUserId"] = o.LoggedByUserId
	}
	if o.Minutes != nil {
		toSerialize["minutes"] = o.Minutes
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.ProjectBillingInvoice != nil {
		toSerialize["projectBillingInvoice"] = o.ProjectBillingInvoice
	}
	if o.ProjectBillingInvoiceId != nil {
		toSerialize["projectBillingInvoiceId"] = o.ProjectBillingInvoiceId
	}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.TagIds != nil {
		toSerialize["tagIds"] = o.TagIds
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Task != nil {
		toSerialize["task"] = o.Task
	}
	if o.TaskId != nil {
		toSerialize["taskId"] = o.TaskId
	}
	if o.TaskIdPreMove != nil {
		toSerialize["taskIdPreMove"] = o.TaskIdPreMove
	}
	if o.TaskPreMove != nil {
		toSerialize["taskPreMove"] = o.TaskPreMove
	}
	if o.TimeLogged != nil {
		toSerialize["timeLogged"] = o.TimeLogged
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.UpdatedBy != nil {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullableViewTimelog struct {
	value *ViewTimelog
	isSet bool
}

func (v NullableViewTimelog) Get() *ViewTimelog {
	return v.value
}

func (v *NullableViewTimelog) Set(val *ViewTimelog) {
	v.value = val
	v.isSet = true
}

func (v NullableViewTimelog) IsSet() bool {
	return v.isSet
}

func (v *NullableViewTimelog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewTimelog(val *ViewTimelog) *NullableViewTimelog {
	return &NullableViewTimelog{value: val, isSet: true}
}

func (v NullableViewTimelog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewTimelog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


