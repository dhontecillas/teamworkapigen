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

// SummaryColumnCardResponse ColumnCardResponse contains counters from column's cards.
type SummaryColumnCardResponse struct {
	Active *int32 `json:"active,omitempty"`
	Archived *int32 `json:"archived,omitempty"`
	Completed *int32 `json:"completed,omitempty"`
	Count *int32 `json:"count,omitempty"`
}

// NewSummaryColumnCardResponse instantiates a new SummaryColumnCardResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSummaryColumnCardResponse() *SummaryColumnCardResponse {
	this := SummaryColumnCardResponse{}
	return &this
}

// NewSummaryColumnCardResponseWithDefaults instantiates a new SummaryColumnCardResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSummaryColumnCardResponseWithDefaults() *SummaryColumnCardResponse {
	this := SummaryColumnCardResponse{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *SummaryColumnCardResponse) GetActive() int32 {
	if o == nil || o.Active == nil {
		var ret int32
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryColumnCardResponse) GetActiveOk() (*int32, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *SummaryColumnCardResponse) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given int32 and assigns it to the Active field.
func (o *SummaryColumnCardResponse) SetActive(v int32) {
	o.Active = &v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *SummaryColumnCardResponse) GetArchived() int32 {
	if o == nil || o.Archived == nil {
		var ret int32
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryColumnCardResponse) GetArchivedOk() (*int32, bool) {
	if o == nil || o.Archived == nil {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *SummaryColumnCardResponse) HasArchived() bool {
	if o != nil && o.Archived != nil {
		return true
	}

	return false
}

// SetArchived gets a reference to the given int32 and assigns it to the Archived field.
func (o *SummaryColumnCardResponse) SetArchived(v int32) {
	o.Archived = &v
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *SummaryColumnCardResponse) GetCompleted() int32 {
	if o == nil || o.Completed == nil {
		var ret int32
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryColumnCardResponse) GetCompletedOk() (*int32, bool) {
	if o == nil || o.Completed == nil {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *SummaryColumnCardResponse) HasCompleted() bool {
	if o != nil && o.Completed != nil {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given int32 and assigns it to the Completed field.
func (o *SummaryColumnCardResponse) SetCompleted(v int32) {
	o.Completed = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *SummaryColumnCardResponse) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryColumnCardResponse) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *SummaryColumnCardResponse) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *SummaryColumnCardResponse) SetCount(v int32) {
	o.Count = &v
}

func (o SummaryColumnCardResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Archived != nil {
		toSerialize["archived"] = o.Archived
	}
	if o.Completed != nil {
		toSerialize["completed"] = o.Completed
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableSummaryColumnCardResponse struct {
	value *SummaryColumnCardResponse
	isSet bool
}

func (v NullableSummaryColumnCardResponse) Get() *SummaryColumnCardResponse {
	return v.value
}

func (v *NullableSummaryColumnCardResponse) Set(val *SummaryColumnCardResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSummaryColumnCardResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSummaryColumnCardResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSummaryColumnCardResponse(val *SummaryColumnCardResponse) *NullableSummaryColumnCardResponse {
	return &NullableSummaryColumnCardResponse{value: val, isSet: true}
}

func (v NullableSummaryColumnCardResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSummaryColumnCardResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


