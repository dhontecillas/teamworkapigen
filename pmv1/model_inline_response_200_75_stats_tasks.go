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

// InlineResponse20075StatsTasks struct for InlineResponse20075StatsTasks
type InlineResponse20075StatsTasks struct {
	Active *string `json:"active,omitempty"`
	Complete *string `json:"complete,omitempty"`
	Late *string `json:"late,omitempty"`
	Nodate *string `json:"nodate,omitempty"`
	Today *string `json:"today,omitempty"`
	Upcoming *string `json:"upcoming,omitempty"`
}

// NewInlineResponse20075StatsTasks instantiates a new InlineResponse20075StatsTasks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20075StatsTasks() *InlineResponse20075StatsTasks {
	this := InlineResponse20075StatsTasks{}
	return &this
}

// NewInlineResponse20075StatsTasksWithDefaults instantiates a new InlineResponse20075StatsTasks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20075StatsTasksWithDefaults() *InlineResponse20075StatsTasks {
	this := InlineResponse20075StatsTasks{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *InlineResponse20075StatsTasks) GetActive() string {
	if o == nil || o.Active == nil {
		var ret string
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20075StatsTasks) GetActiveOk() (*string, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *InlineResponse20075StatsTasks) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given string and assigns it to the Active field.
func (o *InlineResponse20075StatsTasks) SetActive(v string) {
	o.Active = &v
}

// GetComplete returns the Complete field value if set, zero value otherwise.
func (o *InlineResponse20075StatsTasks) GetComplete() string {
	if o == nil || o.Complete == nil {
		var ret string
		return ret
	}
	return *o.Complete
}

// GetCompleteOk returns a tuple with the Complete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20075StatsTasks) GetCompleteOk() (*string, bool) {
	if o == nil || o.Complete == nil {
		return nil, false
	}
	return o.Complete, true
}

// HasComplete returns a boolean if a field has been set.
func (o *InlineResponse20075StatsTasks) HasComplete() bool {
	if o != nil && o.Complete != nil {
		return true
	}

	return false
}

// SetComplete gets a reference to the given string and assigns it to the Complete field.
func (o *InlineResponse20075StatsTasks) SetComplete(v string) {
	o.Complete = &v
}

// GetLate returns the Late field value if set, zero value otherwise.
func (o *InlineResponse20075StatsTasks) GetLate() string {
	if o == nil || o.Late == nil {
		var ret string
		return ret
	}
	return *o.Late
}

// GetLateOk returns a tuple with the Late field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20075StatsTasks) GetLateOk() (*string, bool) {
	if o == nil || o.Late == nil {
		return nil, false
	}
	return o.Late, true
}

// HasLate returns a boolean if a field has been set.
func (o *InlineResponse20075StatsTasks) HasLate() bool {
	if o != nil && o.Late != nil {
		return true
	}

	return false
}

// SetLate gets a reference to the given string and assigns it to the Late field.
func (o *InlineResponse20075StatsTasks) SetLate(v string) {
	o.Late = &v
}

// GetNodate returns the Nodate field value if set, zero value otherwise.
func (o *InlineResponse20075StatsTasks) GetNodate() string {
	if o == nil || o.Nodate == nil {
		var ret string
		return ret
	}
	return *o.Nodate
}

// GetNodateOk returns a tuple with the Nodate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20075StatsTasks) GetNodateOk() (*string, bool) {
	if o == nil || o.Nodate == nil {
		return nil, false
	}
	return o.Nodate, true
}

// HasNodate returns a boolean if a field has been set.
func (o *InlineResponse20075StatsTasks) HasNodate() bool {
	if o != nil && o.Nodate != nil {
		return true
	}

	return false
}

// SetNodate gets a reference to the given string and assigns it to the Nodate field.
func (o *InlineResponse20075StatsTasks) SetNodate(v string) {
	o.Nodate = &v
}

// GetToday returns the Today field value if set, zero value otherwise.
func (o *InlineResponse20075StatsTasks) GetToday() string {
	if o == nil || o.Today == nil {
		var ret string
		return ret
	}
	return *o.Today
}

// GetTodayOk returns a tuple with the Today field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20075StatsTasks) GetTodayOk() (*string, bool) {
	if o == nil || o.Today == nil {
		return nil, false
	}
	return o.Today, true
}

// HasToday returns a boolean if a field has been set.
func (o *InlineResponse20075StatsTasks) HasToday() bool {
	if o != nil && o.Today != nil {
		return true
	}

	return false
}

// SetToday gets a reference to the given string and assigns it to the Today field.
func (o *InlineResponse20075StatsTasks) SetToday(v string) {
	o.Today = &v
}

// GetUpcoming returns the Upcoming field value if set, zero value otherwise.
func (o *InlineResponse20075StatsTasks) GetUpcoming() string {
	if o == nil || o.Upcoming == nil {
		var ret string
		return ret
	}
	return *o.Upcoming
}

// GetUpcomingOk returns a tuple with the Upcoming field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20075StatsTasks) GetUpcomingOk() (*string, bool) {
	if o == nil || o.Upcoming == nil {
		return nil, false
	}
	return o.Upcoming, true
}

// HasUpcoming returns a boolean if a field has been set.
func (o *InlineResponse20075StatsTasks) HasUpcoming() bool {
	if o != nil && o.Upcoming != nil {
		return true
	}

	return false
}

// SetUpcoming gets a reference to the given string and assigns it to the Upcoming field.
func (o *InlineResponse20075StatsTasks) SetUpcoming(v string) {
	o.Upcoming = &v
}

func (o InlineResponse20075StatsTasks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Complete != nil {
		toSerialize["complete"] = o.Complete
	}
	if o.Late != nil {
		toSerialize["late"] = o.Late
	}
	if o.Nodate != nil {
		toSerialize["nodate"] = o.Nodate
	}
	if o.Today != nil {
		toSerialize["today"] = o.Today
	}
	if o.Upcoming != nil {
		toSerialize["upcoming"] = o.Upcoming
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20075StatsTasks struct {
	value *InlineResponse20075StatsTasks
	isSet bool
}

func (v NullableInlineResponse20075StatsTasks) Get() *InlineResponse20075StatsTasks {
	return v.value
}

func (v *NullableInlineResponse20075StatsTasks) Set(val *InlineResponse20075StatsTasks) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20075StatsTasks) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20075StatsTasks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20075StatsTasks(val *InlineResponse20075StatsTasks) *NullableInlineResponse20075StatsTasks {
	return &NullableInlineResponse20075StatsTasks{value: val, isSet: true}
}

func (v NullableInlineResponse20075StatsTasks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20075StatsTasks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


