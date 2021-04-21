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

// InlineResponse20060ProjectActivePages struct for InlineResponse20060ProjectActivePages
type InlineResponse20060ProjectActivePages struct {
	Billing *string `json:"billing,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Files *string `json:"files,omitempty"`
	Links *string `json:"links,omitempty"`
	Messages *string `json:"messages,omitempty"`
	Milestones *string `json:"milestones,omitempty"`
	Notebooks *string `json:"notebooks,omitempty"`
	RiskRegister *string `json:"riskRegister,omitempty"`
	Tasks *string `json:"tasks,omitempty"`
	Time *string `json:"time,omitempty"`
}

// NewInlineResponse20060ProjectActivePages instantiates a new InlineResponse20060ProjectActivePages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20060ProjectActivePages() *InlineResponse20060ProjectActivePages {
	this := InlineResponse20060ProjectActivePages{}
	return &this
}

// NewInlineResponse20060ProjectActivePagesWithDefaults instantiates a new InlineResponse20060ProjectActivePages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20060ProjectActivePagesWithDefaults() *InlineResponse20060ProjectActivePages {
	this := InlineResponse20060ProjectActivePages{}
	return &this
}

// GetBilling returns the Billing field value if set, zero value otherwise.
func (o *InlineResponse20060ProjectActivePages) GetBilling() string {
	if o == nil || o.Billing == nil {
		var ret string
		return ret
	}
	return *o.Billing
}

// GetBillingOk returns a tuple with the Billing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060ProjectActivePages) GetBillingOk() (*string, bool) {
	if o == nil || o.Billing == nil {
		return nil, false
	}
	return o.Billing, true
}

// HasBilling returns a boolean if a field has been set.
func (o *InlineResponse20060ProjectActivePages) HasBilling() bool {
	if o != nil && o.Billing != nil {
		return true
	}

	return false
}

// SetBilling gets a reference to the given string and assigns it to the Billing field.
func (o *InlineResponse20060ProjectActivePages) SetBilling(v string) {
	o.Billing = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *InlineResponse20060ProjectActivePages) GetComments() string {
	if o == nil || o.Comments == nil {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060ProjectActivePages) GetCommentsOk() (*string, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *InlineResponse20060ProjectActivePages) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *InlineResponse20060ProjectActivePages) SetComments(v string) {
	o.Comments = &v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *InlineResponse20060ProjectActivePages) GetFiles() string {
	if o == nil || o.Files == nil {
		var ret string
		return ret
	}
	return *o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060ProjectActivePages) GetFilesOk() (*string, bool) {
	if o == nil || o.Files == nil {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *InlineResponse20060ProjectActivePages) HasFiles() bool {
	if o != nil && o.Files != nil {
		return true
	}

	return false
}

// SetFiles gets a reference to the given string and assigns it to the Files field.
func (o *InlineResponse20060ProjectActivePages) SetFiles(v string) {
	o.Files = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *InlineResponse20060ProjectActivePages) GetLinks() string {
	if o == nil || o.Links == nil {
		var ret string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060ProjectActivePages) GetLinksOk() (*string, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *InlineResponse20060ProjectActivePages) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given string and assigns it to the Links field.
func (o *InlineResponse20060ProjectActivePages) SetLinks(v string) {
	o.Links = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *InlineResponse20060ProjectActivePages) GetMessages() string {
	if o == nil || o.Messages == nil {
		var ret string
		return ret
	}
	return *o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060ProjectActivePages) GetMessagesOk() (*string, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *InlineResponse20060ProjectActivePages) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given string and assigns it to the Messages field.
func (o *InlineResponse20060ProjectActivePages) SetMessages(v string) {
	o.Messages = &v
}

// GetMilestones returns the Milestones field value if set, zero value otherwise.
func (o *InlineResponse20060ProjectActivePages) GetMilestones() string {
	if o == nil || o.Milestones == nil {
		var ret string
		return ret
	}
	return *o.Milestones
}

// GetMilestonesOk returns a tuple with the Milestones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060ProjectActivePages) GetMilestonesOk() (*string, bool) {
	if o == nil || o.Milestones == nil {
		return nil, false
	}
	return o.Milestones, true
}

// HasMilestones returns a boolean if a field has been set.
func (o *InlineResponse20060ProjectActivePages) HasMilestones() bool {
	if o != nil && o.Milestones != nil {
		return true
	}

	return false
}

// SetMilestones gets a reference to the given string and assigns it to the Milestones field.
func (o *InlineResponse20060ProjectActivePages) SetMilestones(v string) {
	o.Milestones = &v
}

// GetNotebooks returns the Notebooks field value if set, zero value otherwise.
func (o *InlineResponse20060ProjectActivePages) GetNotebooks() string {
	if o == nil || o.Notebooks == nil {
		var ret string
		return ret
	}
	return *o.Notebooks
}

// GetNotebooksOk returns a tuple with the Notebooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060ProjectActivePages) GetNotebooksOk() (*string, bool) {
	if o == nil || o.Notebooks == nil {
		return nil, false
	}
	return o.Notebooks, true
}

// HasNotebooks returns a boolean if a field has been set.
func (o *InlineResponse20060ProjectActivePages) HasNotebooks() bool {
	if o != nil && o.Notebooks != nil {
		return true
	}

	return false
}

// SetNotebooks gets a reference to the given string and assigns it to the Notebooks field.
func (o *InlineResponse20060ProjectActivePages) SetNotebooks(v string) {
	o.Notebooks = &v
}

// GetRiskRegister returns the RiskRegister field value if set, zero value otherwise.
func (o *InlineResponse20060ProjectActivePages) GetRiskRegister() string {
	if o == nil || o.RiskRegister == nil {
		var ret string
		return ret
	}
	return *o.RiskRegister
}

// GetRiskRegisterOk returns a tuple with the RiskRegister field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060ProjectActivePages) GetRiskRegisterOk() (*string, bool) {
	if o == nil || o.RiskRegister == nil {
		return nil, false
	}
	return o.RiskRegister, true
}

// HasRiskRegister returns a boolean if a field has been set.
func (o *InlineResponse20060ProjectActivePages) HasRiskRegister() bool {
	if o != nil && o.RiskRegister != nil {
		return true
	}

	return false
}

// SetRiskRegister gets a reference to the given string and assigns it to the RiskRegister field.
func (o *InlineResponse20060ProjectActivePages) SetRiskRegister(v string) {
	o.RiskRegister = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *InlineResponse20060ProjectActivePages) GetTasks() string {
	if o == nil || o.Tasks == nil {
		var ret string
		return ret
	}
	return *o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060ProjectActivePages) GetTasksOk() (*string, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *InlineResponse20060ProjectActivePages) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given string and assigns it to the Tasks field.
func (o *InlineResponse20060ProjectActivePages) SetTasks(v string) {
	o.Tasks = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *InlineResponse20060ProjectActivePages) GetTime() string {
	if o == nil || o.Time == nil {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060ProjectActivePages) GetTimeOk() (*string, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *InlineResponse20060ProjectActivePages) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *InlineResponse20060ProjectActivePages) SetTime(v string) {
	o.Time = &v
}

func (o InlineResponse20060ProjectActivePages) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Billing != nil {
		toSerialize["billing"] = o.Billing
	}
	if o.Comments != nil {
		toSerialize["comments"] = o.Comments
	}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	if o.Milestones != nil {
		toSerialize["milestones"] = o.Milestones
	}
	if o.Notebooks != nil {
		toSerialize["notebooks"] = o.Notebooks
	}
	if o.RiskRegister != nil {
		toSerialize["riskRegister"] = o.RiskRegister
	}
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20060ProjectActivePages struct {
	value *InlineResponse20060ProjectActivePages
	isSet bool
}

func (v NullableInlineResponse20060ProjectActivePages) Get() *InlineResponse20060ProjectActivePages {
	return v.value
}

func (v *NullableInlineResponse20060ProjectActivePages) Set(val *InlineResponse20060ProjectActivePages) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20060ProjectActivePages) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20060ProjectActivePages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20060ProjectActivePages(val *InlineResponse20060ProjectActivePages) *NullableInlineResponse20060ProjectActivePages {
	return &NullableInlineResponse20060ProjectActivePages{value: val, isSet: true}
}

func (v NullableInlineResponse20060ProjectActivePages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20060ProjectActivePages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


