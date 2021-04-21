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

// BudgetProjectBudgetSummary ProjectBudgetSummary has a summary with some of the view.ProjectBudget fields
type BudgetProjectBudgetSummary struct {
	Capacity *int32 `json:"capacity,omitempty"`
	CapacityUsed *int32 `json:"capacityUsed,omitempty"`
	EndDateTime *string `json:"endDateTime,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Project *ViewRelationship `json:"project,omitempty"`
	ProjectId *int32 `json:"projectId,omitempty"`
	StartDateTime *string `json:"startDateTime,omitempty"`
	TimelogType *string `json:"timelogType,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewBudgetProjectBudgetSummary instantiates a new BudgetProjectBudgetSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetProjectBudgetSummary() *BudgetProjectBudgetSummary {
	this := BudgetProjectBudgetSummary{}
	return &this
}

// NewBudgetProjectBudgetSummaryWithDefaults instantiates a new BudgetProjectBudgetSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetProjectBudgetSummaryWithDefaults() *BudgetProjectBudgetSummary {
	this := BudgetProjectBudgetSummary{}
	return &this
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *BudgetProjectBudgetSummary) GetCapacity() int32 {
	if o == nil || o.Capacity == nil {
		var ret int32
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetProjectBudgetSummary) GetCapacityOk() (*int32, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *BudgetProjectBudgetSummary) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given int32 and assigns it to the Capacity field.
func (o *BudgetProjectBudgetSummary) SetCapacity(v int32) {
	o.Capacity = &v
}

// GetCapacityUsed returns the CapacityUsed field value if set, zero value otherwise.
func (o *BudgetProjectBudgetSummary) GetCapacityUsed() int32 {
	if o == nil || o.CapacityUsed == nil {
		var ret int32
		return ret
	}
	return *o.CapacityUsed
}

// GetCapacityUsedOk returns a tuple with the CapacityUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetProjectBudgetSummary) GetCapacityUsedOk() (*int32, bool) {
	if o == nil || o.CapacityUsed == nil {
		return nil, false
	}
	return o.CapacityUsed, true
}

// HasCapacityUsed returns a boolean if a field has been set.
func (o *BudgetProjectBudgetSummary) HasCapacityUsed() bool {
	if o != nil && o.CapacityUsed != nil {
		return true
	}

	return false
}

// SetCapacityUsed gets a reference to the given int32 and assigns it to the CapacityUsed field.
func (o *BudgetProjectBudgetSummary) SetCapacityUsed(v int32) {
	o.CapacityUsed = &v
}

// GetEndDateTime returns the EndDateTime field value if set, zero value otherwise.
func (o *BudgetProjectBudgetSummary) GetEndDateTime() string {
	if o == nil || o.EndDateTime == nil {
		var ret string
		return ret
	}
	return *o.EndDateTime
}

// GetEndDateTimeOk returns a tuple with the EndDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetProjectBudgetSummary) GetEndDateTimeOk() (*string, bool) {
	if o == nil || o.EndDateTime == nil {
		return nil, false
	}
	return o.EndDateTime, true
}

// HasEndDateTime returns a boolean if a field has been set.
func (o *BudgetProjectBudgetSummary) HasEndDateTime() bool {
	if o != nil && o.EndDateTime != nil {
		return true
	}

	return false
}

// SetEndDateTime gets a reference to the given string and assigns it to the EndDateTime field.
func (o *BudgetProjectBudgetSummary) SetEndDateTime(v string) {
	o.EndDateTime = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BudgetProjectBudgetSummary) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetProjectBudgetSummary) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BudgetProjectBudgetSummary) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *BudgetProjectBudgetSummary) SetId(v int32) {
	o.Id = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *BudgetProjectBudgetSummary) GetProject() ViewRelationship {
	if o == nil || o.Project == nil {
		var ret ViewRelationship
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetProjectBudgetSummary) GetProjectOk() (*ViewRelationship, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *BudgetProjectBudgetSummary) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given ViewRelationship and assigns it to the Project field.
func (o *BudgetProjectBudgetSummary) SetProject(v ViewRelationship) {
	o.Project = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *BudgetProjectBudgetSummary) GetProjectId() int32 {
	if o == nil || o.ProjectId == nil {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetProjectBudgetSummary) GetProjectIdOk() (*int32, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *BudgetProjectBudgetSummary) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *BudgetProjectBudgetSummary) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetStartDateTime returns the StartDateTime field value if set, zero value otherwise.
func (o *BudgetProjectBudgetSummary) GetStartDateTime() string {
	if o == nil || o.StartDateTime == nil {
		var ret string
		return ret
	}
	return *o.StartDateTime
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetProjectBudgetSummary) GetStartDateTimeOk() (*string, bool) {
	if o == nil || o.StartDateTime == nil {
		return nil, false
	}
	return o.StartDateTime, true
}

// HasStartDateTime returns a boolean if a field has been set.
func (o *BudgetProjectBudgetSummary) HasStartDateTime() bool {
	if o != nil && o.StartDateTime != nil {
		return true
	}

	return false
}

// SetStartDateTime gets a reference to the given string and assigns it to the StartDateTime field.
func (o *BudgetProjectBudgetSummary) SetStartDateTime(v string) {
	o.StartDateTime = &v
}

// GetTimelogType returns the TimelogType field value if set, zero value otherwise.
func (o *BudgetProjectBudgetSummary) GetTimelogType() string {
	if o == nil || o.TimelogType == nil {
		var ret string
		return ret
	}
	return *o.TimelogType
}

// GetTimelogTypeOk returns a tuple with the TimelogType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetProjectBudgetSummary) GetTimelogTypeOk() (*string, bool) {
	if o == nil || o.TimelogType == nil {
		return nil, false
	}
	return o.TimelogType, true
}

// HasTimelogType returns a boolean if a field has been set.
func (o *BudgetProjectBudgetSummary) HasTimelogType() bool {
	if o != nil && o.TimelogType != nil {
		return true
	}

	return false
}

// SetTimelogType gets a reference to the given string and assigns it to the TimelogType field.
func (o *BudgetProjectBudgetSummary) SetTimelogType(v string) {
	o.TimelogType = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BudgetProjectBudgetSummary) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetProjectBudgetSummary) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BudgetProjectBudgetSummary) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BudgetProjectBudgetSummary) SetType(v string) {
	o.Type = &v
}

func (o BudgetProjectBudgetSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Capacity != nil {
		toSerialize["capacity"] = o.Capacity
	}
	if o.CapacityUsed != nil {
		toSerialize["capacityUsed"] = o.CapacityUsed
	}
	if o.EndDateTime != nil {
		toSerialize["endDateTime"] = o.EndDateTime
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.StartDateTime != nil {
		toSerialize["startDateTime"] = o.StartDateTime
	}
	if o.TimelogType != nil {
		toSerialize["timelogType"] = o.TimelogType
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableBudgetProjectBudgetSummary struct {
	value *BudgetProjectBudgetSummary
	isSet bool
}

func (v NullableBudgetProjectBudgetSummary) Get() *BudgetProjectBudgetSummary {
	return v.value
}

func (v *NullableBudgetProjectBudgetSummary) Set(val *BudgetProjectBudgetSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetProjectBudgetSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetProjectBudgetSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetProjectBudgetSummary(val *BudgetProjectBudgetSummary) *NullableBudgetProjectBudgetSummary {
	return &NullableBudgetProjectBudgetSummary{value: val, isSet: true}
}

func (v NullableBudgetProjectBudgetSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBudgetProjectBudgetSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

