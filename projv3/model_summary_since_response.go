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

// SummarySinceResponse SinceResponse contains counters for since.
type SummarySinceResponse struct {
	DateTime *string `json:"dateTime,omitempty"`
	Events *int32 `json:"events,omitempty"`
	TasksComplete *int32 `json:"tasksComplete,omitempty"`
	TasksCreated *int32 `json:"tasksCreated,omitempty"`
}

// NewSummarySinceResponse instantiates a new SummarySinceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSummarySinceResponse() *SummarySinceResponse {
	this := SummarySinceResponse{}
	return &this
}

// NewSummarySinceResponseWithDefaults instantiates a new SummarySinceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSummarySinceResponseWithDefaults() *SummarySinceResponse {
	this := SummarySinceResponse{}
	return &this
}

// GetDateTime returns the DateTime field value if set, zero value otherwise.
func (o *SummarySinceResponse) GetDateTime() string {
	if o == nil || o.DateTime == nil {
		var ret string
		return ret
	}
	return *o.DateTime
}

// GetDateTimeOk returns a tuple with the DateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummarySinceResponse) GetDateTimeOk() (*string, bool) {
	if o == nil || o.DateTime == nil {
		return nil, false
	}
	return o.DateTime, true
}

// HasDateTime returns a boolean if a field has been set.
func (o *SummarySinceResponse) HasDateTime() bool {
	if o != nil && o.DateTime != nil {
		return true
	}

	return false
}

// SetDateTime gets a reference to the given string and assigns it to the DateTime field.
func (o *SummarySinceResponse) SetDateTime(v string) {
	o.DateTime = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *SummarySinceResponse) GetEvents() int32 {
	if o == nil || o.Events == nil {
		var ret int32
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummarySinceResponse) GetEventsOk() (*int32, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *SummarySinceResponse) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given int32 and assigns it to the Events field.
func (o *SummarySinceResponse) SetEvents(v int32) {
	o.Events = &v
}

// GetTasksComplete returns the TasksComplete field value if set, zero value otherwise.
func (o *SummarySinceResponse) GetTasksComplete() int32 {
	if o == nil || o.TasksComplete == nil {
		var ret int32
		return ret
	}
	return *o.TasksComplete
}

// GetTasksCompleteOk returns a tuple with the TasksComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummarySinceResponse) GetTasksCompleteOk() (*int32, bool) {
	if o == nil || o.TasksComplete == nil {
		return nil, false
	}
	return o.TasksComplete, true
}

// HasTasksComplete returns a boolean if a field has been set.
func (o *SummarySinceResponse) HasTasksComplete() bool {
	if o != nil && o.TasksComplete != nil {
		return true
	}

	return false
}

// SetTasksComplete gets a reference to the given int32 and assigns it to the TasksComplete field.
func (o *SummarySinceResponse) SetTasksComplete(v int32) {
	o.TasksComplete = &v
}

// GetTasksCreated returns the TasksCreated field value if set, zero value otherwise.
func (o *SummarySinceResponse) GetTasksCreated() int32 {
	if o == nil || o.TasksCreated == nil {
		var ret int32
		return ret
	}
	return *o.TasksCreated
}

// GetTasksCreatedOk returns a tuple with the TasksCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummarySinceResponse) GetTasksCreatedOk() (*int32, bool) {
	if o == nil || o.TasksCreated == nil {
		return nil, false
	}
	return o.TasksCreated, true
}

// HasTasksCreated returns a boolean if a field has been set.
func (o *SummarySinceResponse) HasTasksCreated() bool {
	if o != nil && o.TasksCreated != nil {
		return true
	}

	return false
}

// SetTasksCreated gets a reference to the given int32 and assigns it to the TasksCreated field.
func (o *SummarySinceResponse) SetTasksCreated(v int32) {
	o.TasksCreated = &v
}

func (o SummarySinceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DateTime != nil {
		toSerialize["dateTime"] = o.DateTime
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	if o.TasksComplete != nil {
		toSerialize["tasksComplete"] = o.TasksComplete
	}
	if o.TasksCreated != nil {
		toSerialize["tasksCreated"] = o.TasksCreated
	}
	return json.Marshal(toSerialize)
}

type NullableSummarySinceResponse struct {
	value *SummarySinceResponse
	isSet bool
}

func (v NullableSummarySinceResponse) Get() *SummarySinceResponse {
	return v.value
}

func (v *NullableSummarySinceResponse) Set(val *SummarySinceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSummarySinceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSummarySinceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSummarySinceResponse(val *SummarySinceResponse) *NullableSummarySinceResponse {
	return &NullableSummarySinceResponse{value: val, isSet: true}
}

func (v NullableSummarySinceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSummarySinceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


