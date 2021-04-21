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

// ViewBoardColumnSettings BoardColumnSettings contains all the settings for a column.
type ViewBoardColumnSettings struct {
	Assignee *bool `json:"assignee,omitempty"`
	Avatar *bool `json:"avatar,omitempty"`
	Comments *bool `json:"comments,omitempty"`
	Dependencies *bool `json:"dependencies,omitempty"`
	EndAt *bool `json:"endAt,omitempty"`
	Estimatedtime *bool `json:"estimatedtime,omitempty"`
	Files *bool `json:"files,omitempty"`
	Followers *bool `json:"followers,omitempty"`
	Name *bool `json:"name,omitempty"`
	Priority *bool `json:"priority,omitempty"`
	Private *bool `json:"private,omitempty"`
	Progress *bool `json:"progress,omitempty"`
	Recurring *bool `json:"recurring,omitempty"`
	Reminders *bool `json:"reminders,omitempty"`
	StartAt *bool `json:"startAt,omitempty"`
	Subtasklabel *bool `json:"subtasklabel,omitempty"`
	Subtasktext *bool `json:"subtasktext,omitempty"`
	Tags *bool `json:"tags,omitempty"`
	Tasklist *bool `json:"tasklist,omitempty"`
	Tickets *bool `json:"tickets,omitempty"`
	Time *bool `json:"time,omitempty"`
}

// NewViewBoardColumnSettings instantiates a new ViewBoardColumnSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewBoardColumnSettings() *ViewBoardColumnSettings {
	this := ViewBoardColumnSettings{}
	return &this
}

// NewViewBoardColumnSettingsWithDefaults instantiates a new ViewBoardColumnSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewBoardColumnSettingsWithDefaults() *ViewBoardColumnSettings {
	this := ViewBoardColumnSettings{}
	return &this
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *ViewBoardColumnSettings) GetAssignee() bool {
	if o == nil || o.Assignee == nil {
		var ret bool
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewBoardColumnSettings) GetAssigneeOk() (*bool, bool) {
	if o == nil || o.Assignee == nil {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *ViewBoardColumnSettings) HasAssignee() bool {
	if o != nil && o.Assignee != nil {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given bool and assigns it to the Assignee field.
func (o *ViewBoardColumnSettings) SetAssignee(v bool) {
	o.Assignee = &v
}

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *ViewBoardColumnSettings) GetAvatar() bool {
	if o == nil || o.Avatar == nil {
		var ret bool
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewBoardColumnSettings) GetAvatarOk() (*bool, bool) {
	if o == nil || o.Avatar == nil {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *ViewBoardColumnSettings) HasAvatar() bool {
	if o != nil && o.Avatar != nil {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given bool and assigns it to the Avatar field.
func (o *ViewBoardColumnSettings) SetAvatar(v bool) {
	o.Avatar = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *ViewBoardColumnSettings) GetComments() bool {
	if o == nil || o.Comments == nil {
		var ret bool
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewBoardColumnSettings) GetCommentsOk() (*bool, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *ViewBoardColumnSettings) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given bool and assigns it to the Comments field.
func (o *ViewBoardColumnSettings) SetComments(v bool) {
	o.Comments = &v
}

// GetDependencies returns the Dependencies field value if set, zero value otherwise.
func (o *ViewBoardColumnSettings) GetDependencies() bool {
	if o == nil || o.Dependencies == nil {
		var ret bool
		return ret
	}
	return *o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewBoardColumnSettings) GetDependenciesOk() (*bool, bool) {
	if o == nil || o.Dependencies == nil {
		return nil, false
	}
	return o.Dependencies, true
}

// HasDependencies returns a boolean if a field has been set.
func (o *ViewBoardColumnSettings) HasDependencies() bool {
	if o != nil && o.Dependencies != nil {
		return true
	}

	return false
}

// SetDependencies gets a reference to the given bool and assigns it to the Dependencies field.
func (o *ViewBoardColumnSettings) SetDependencies(v bool) {
	o.Dependencies = &v
}

// GetEndAt returns the EndAt field value if set, zero value otherwise.
func (o *ViewBoardColumnSettings) GetEndAt() bool {
	if o == nil || o.EndAt == nil {
		var ret bool
		return ret
	}
	return *o.EndAt
}

// GetEndAtOk returns a tuple with the EndAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewBoardColumnSettings) GetEndAtOk() (*bool, bool) {
	if o == nil || o.EndAt == nil {
		return nil, false
	}
	return o.EndAt, true
}

// HasEndAt returns a boolean if a field has been set.
func (o *ViewBoardColumnSettings) HasEndAt() bool {
	if o != nil && o.EndAt != nil {
		return true
	}

	return false
}

// SetEndAt gets a reference to the given bool and assigns it to the EndAt field.
func (o *ViewBoardColumnSettings) SetEndAt(v bool) {
	o.EndAt = &v
}

// GetEstimatedtime returns the Estimatedtime field value if set, zero value otherwise.
func (o *ViewBoardColumnSettings) GetEstimatedtime() bool {
	if o == nil || o.Estimatedtime == nil {
		var ret bool
		return ret
	}
	return *o.Estimatedtime
}

// GetEstimatedtimeOk returns a tuple with the Estimatedtime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewBoardColumnSettings) GetEstimatedtimeOk() (*bool, bool) {
	if o == nil || o.Estimatedtime == nil {
		return nil, false
	}
	return o.Estimatedtime, true
}

// HasEstimatedtime returns a boolean if a field has been set.
func (o *ViewBoardColumnSettings) HasEstimatedtime() bool {
	if o != nil && o.Estimatedtime != nil {
		return true
	}

	return false
}

// SetEstimatedtime gets a reference to the given bool and assigns it to the Estimatedtime field.
func (o *ViewBoardColumnSettings) SetEstimatedtime(v bool) {
	o.Estimatedtime = &v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *ViewBoardColumnSettings) GetFiles() bool {
	if o == nil || o.Files == nil {
		var ret bool
		return ret
	}
	return *o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewBoardColumnSettings) GetFilesOk() (*bool, bool) {
	if o == nil || o.Files == nil {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *ViewBoardColumnSettings) HasFiles() bool {
	if o != nil && o.Files != nil {
		return true
	}

	return false
}

// SetFiles gets a reference to the given bool and assigns it to the Files field.
func (o *ViewBoardColumnSettings) SetFiles(v bool) {
	o.Files = &v
}

// GetFollowers returns the Followers field value if set, zero value otherwise.
func (o *ViewBoardColumnSettings) GetFollowers() bool {
	if o == nil || o.Followers == nil {
		var ret bool
		return ret
	}
	return *o.Followers
}

// GetFollowersOk returns a tuple with the Followers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewBoardColumnSettings) GetFollowersOk() (*bool, bool) {
	if o == nil || o.Followers == nil {
		return nil, false
	}
	return o.Followers, true
}

// HasFollowers returns a boolean if a field has been set.
func (o *ViewBoardColumnSettings) HasFollowers() bool {
	if o != nil && o.Followers != nil {
		return true
	}

	return false
}

// SetFollowers gets a reference to the given bool and assigns it to the Followers field.
func (o *ViewBoardColumnSettings) SetFollowers(v bool) {
	o.Followers = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ViewBoardColumnSettings) GetName() bool {
	if o == nil || o.Name == nil {
		var ret bool
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewBoardColumnSettings) GetNameOk() (*bool, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ViewBoardColumnSettings) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given bool and assigns it to the Name field.
func (o *ViewBoardColumnSettings) SetName(v bool) {
	o.Name = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *ViewBoardColumnSettings) GetPriority() bool {
	if o == nil || o.Priority == nil {
		var ret bool
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewBoardColumnSettings) GetPriorityOk() (*bool, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *ViewBoardColumnSettings) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given bool and assigns it to the Priority field.
func (o *ViewBoardColumnSettings) SetPriority(v bool) {
	o.Priority = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *ViewBoardColumnSettings) GetPrivate() bool {
	if o == nil || o.Private == nil {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewBoardColumnSettings) GetPrivateOk() (*bool, bool) {
	if o == nil || o.Private == nil {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *ViewBoardColumnSettings) HasPrivate() bool {
	if o != nil && o.Private != nil {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *ViewBoardColumnSettings) SetPrivate(v bool) {
	o.Private = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *ViewBoardColumnSettings) GetProgress() bool {
	if o == nil || o.Progress == nil {
		var ret bool
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewBoardColumnSettings) GetProgressOk() (*bool, bool) {
	if o == nil || o.Progress == nil {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *ViewBoardColumnSettings) HasProgress() bool {
	if o != nil && o.Progress != nil {
		return true
	}

	return false
}

// SetProgress gets a reference to the given bool and assigns it to the Progress field.
func (o *ViewBoardColumnSettings) SetProgress(v bool) {
	o.Progress = &v
}

// GetRecurring returns the Recurring field value if set, zero value otherwise.
func (o *ViewBoardColumnSettings) GetRecurring() bool {
	if o == nil || o.Recurring == nil {
		var ret bool
		return ret
	}
	return *o.Recurring
}

// GetRecurringOk returns a tuple with the Recurring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewBoardColumnSettings) GetRecurringOk() (*bool, bool) {
	if o == nil || o.Recurring == nil {
		return nil, false
	}
	return o.Recurring, true
}

// HasRecurring returns a boolean if a field has been set.
func (o *ViewBoardColumnSettings) HasRecurring() bool {
	if o != nil && o.Recurring != nil {
		return true
	}

	return false
}

// SetRecurring gets a reference to the given bool and assigns it to the Recurring field.
func (o *ViewBoardColumnSettings) SetRecurring(v bool) {
	o.Recurring = &v
}

// GetReminders returns the Reminders field value if set, zero value otherwise.
func (o *ViewBoardColumnSettings) GetReminders() bool {
	if o == nil || o.Reminders == nil {
		var ret bool
		return ret
	}
	return *o.Reminders
}

// GetRemindersOk returns a tuple with the Reminders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewBoardColumnSettings) GetRemindersOk() (*bool, bool) {
	if o == nil || o.Reminders == nil {
		return nil, false
	}
	return o.Reminders, true
}

// HasReminders returns a boolean if a field has been set.
func (o *ViewBoardColumnSettings) HasReminders() bool {
	if o != nil && o.Reminders != nil {
		return true
	}

	return false
}

// SetReminders gets a reference to the given bool and assigns it to the Reminders field.
func (o *ViewBoardColumnSettings) SetReminders(v bool) {
	o.Reminders = &v
}

// GetStartAt returns the StartAt field value if set, zero value otherwise.
func (o *ViewBoardColumnSettings) GetStartAt() bool {
	if o == nil || o.StartAt == nil {
		var ret bool
		return ret
	}
	return *o.StartAt
}

// GetStartAtOk returns a tuple with the StartAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewBoardColumnSettings) GetStartAtOk() (*bool, bool) {
	if o == nil || o.StartAt == nil {
		return nil, false
	}
	return o.StartAt, true
}

// HasStartAt returns a boolean if a field has been set.
func (o *ViewBoardColumnSettings) HasStartAt() bool {
	if o != nil && o.StartAt != nil {
		return true
	}

	return false
}

// SetStartAt gets a reference to the given bool and assigns it to the StartAt field.
func (o *ViewBoardColumnSettings) SetStartAt(v bool) {
	o.StartAt = &v
}

// GetSubtasklabel returns the Subtasklabel field value if set, zero value otherwise.
func (o *ViewBoardColumnSettings) GetSubtasklabel() bool {
	if o == nil || o.Subtasklabel == nil {
		var ret bool
		return ret
	}
	return *o.Subtasklabel
}

// GetSubtasklabelOk returns a tuple with the Subtasklabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewBoardColumnSettings) GetSubtasklabelOk() (*bool, bool) {
	if o == nil || o.Subtasklabel == nil {
		return nil, false
	}
	return o.Subtasklabel, true
}

// HasSubtasklabel returns a boolean if a field has been set.
func (o *ViewBoardColumnSettings) HasSubtasklabel() bool {
	if o != nil && o.Subtasklabel != nil {
		return true
	}

	return false
}

// SetSubtasklabel gets a reference to the given bool and assigns it to the Subtasklabel field.
func (o *ViewBoardColumnSettings) SetSubtasklabel(v bool) {
	o.Subtasklabel = &v
}

// GetSubtasktext returns the Subtasktext field value if set, zero value otherwise.
func (o *ViewBoardColumnSettings) GetSubtasktext() bool {
	if o == nil || o.Subtasktext == nil {
		var ret bool
		return ret
	}
	return *o.Subtasktext
}

// GetSubtasktextOk returns a tuple with the Subtasktext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewBoardColumnSettings) GetSubtasktextOk() (*bool, bool) {
	if o == nil || o.Subtasktext == nil {
		return nil, false
	}
	return o.Subtasktext, true
}

// HasSubtasktext returns a boolean if a field has been set.
func (o *ViewBoardColumnSettings) HasSubtasktext() bool {
	if o != nil && o.Subtasktext != nil {
		return true
	}

	return false
}

// SetSubtasktext gets a reference to the given bool and assigns it to the Subtasktext field.
func (o *ViewBoardColumnSettings) SetSubtasktext(v bool) {
	o.Subtasktext = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ViewBoardColumnSettings) GetTags() bool {
	if o == nil || o.Tags == nil {
		var ret bool
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewBoardColumnSettings) GetTagsOk() (*bool, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ViewBoardColumnSettings) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given bool and assigns it to the Tags field.
func (o *ViewBoardColumnSettings) SetTags(v bool) {
	o.Tags = &v
}

// GetTasklist returns the Tasklist field value if set, zero value otherwise.
func (o *ViewBoardColumnSettings) GetTasklist() bool {
	if o == nil || o.Tasklist == nil {
		var ret bool
		return ret
	}
	return *o.Tasklist
}

// GetTasklistOk returns a tuple with the Tasklist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewBoardColumnSettings) GetTasklistOk() (*bool, bool) {
	if o == nil || o.Tasklist == nil {
		return nil, false
	}
	return o.Tasklist, true
}

// HasTasklist returns a boolean if a field has been set.
func (o *ViewBoardColumnSettings) HasTasklist() bool {
	if o != nil && o.Tasklist != nil {
		return true
	}

	return false
}

// SetTasklist gets a reference to the given bool and assigns it to the Tasklist field.
func (o *ViewBoardColumnSettings) SetTasklist(v bool) {
	o.Tasklist = &v
}

// GetTickets returns the Tickets field value if set, zero value otherwise.
func (o *ViewBoardColumnSettings) GetTickets() bool {
	if o == nil || o.Tickets == nil {
		var ret bool
		return ret
	}
	return *o.Tickets
}

// GetTicketsOk returns a tuple with the Tickets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewBoardColumnSettings) GetTicketsOk() (*bool, bool) {
	if o == nil || o.Tickets == nil {
		return nil, false
	}
	return o.Tickets, true
}

// HasTickets returns a boolean if a field has been set.
func (o *ViewBoardColumnSettings) HasTickets() bool {
	if o != nil && o.Tickets != nil {
		return true
	}

	return false
}

// SetTickets gets a reference to the given bool and assigns it to the Tickets field.
func (o *ViewBoardColumnSettings) SetTickets(v bool) {
	o.Tickets = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *ViewBoardColumnSettings) GetTime() bool {
	if o == nil || o.Time == nil {
		var ret bool
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewBoardColumnSettings) GetTimeOk() (*bool, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *ViewBoardColumnSettings) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given bool and assigns it to the Time field.
func (o *ViewBoardColumnSettings) SetTime(v bool) {
	o.Time = &v
}

func (o ViewBoardColumnSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Assignee != nil {
		toSerialize["assignee"] = o.Assignee
	}
	if o.Avatar != nil {
		toSerialize["avatar"] = o.Avatar
	}
	if o.Comments != nil {
		toSerialize["comments"] = o.Comments
	}
	if o.Dependencies != nil {
		toSerialize["dependencies"] = o.Dependencies
	}
	if o.EndAt != nil {
		toSerialize["endAt"] = o.EndAt
	}
	if o.Estimatedtime != nil {
		toSerialize["estimatedtime"] = o.Estimatedtime
	}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	if o.Followers != nil {
		toSerialize["followers"] = o.Followers
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.Private != nil {
		toSerialize["private"] = o.Private
	}
	if o.Progress != nil {
		toSerialize["progress"] = o.Progress
	}
	if o.Recurring != nil {
		toSerialize["recurring"] = o.Recurring
	}
	if o.Reminders != nil {
		toSerialize["reminders"] = o.Reminders
	}
	if o.StartAt != nil {
		toSerialize["startAt"] = o.StartAt
	}
	if o.Subtasklabel != nil {
		toSerialize["subtasklabel"] = o.Subtasklabel
	}
	if o.Subtasktext != nil {
		toSerialize["subtasktext"] = o.Subtasktext
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Tasklist != nil {
		toSerialize["tasklist"] = o.Tasklist
	}
	if o.Tickets != nil {
		toSerialize["tickets"] = o.Tickets
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	return json.Marshal(toSerialize)
}

type NullableViewBoardColumnSettings struct {
	value *ViewBoardColumnSettings
	isSet bool
}

func (v NullableViewBoardColumnSettings) Get() *ViewBoardColumnSettings {
	return v.value
}

func (v *NullableViewBoardColumnSettings) Set(val *ViewBoardColumnSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableViewBoardColumnSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableViewBoardColumnSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewBoardColumnSettings(val *ViewBoardColumnSettings) *NullableViewBoardColumnSettings {
	return &NullableViewBoardColumnSettings{value: val, isSet: true}
}

func (v NullableViewBoardColumnSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewBoardColumnSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

