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

// InlineResponse200124Deliveries struct for InlineResponse200124Deliveries
type InlineResponse200124Deliveries struct {
	DateCalled *string `json:"dateCalled,omitempty"`
	Event *string `json:"event,omitempty"`
	Id *string `json:"id,omitempty"`
	Request *InlineResponse200124Request `json:"request,omitempty"`
	Response *InlineResponse200124Response `json:"response,omitempty"`
	Url *string `json:"url,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewInlineResponse200124Deliveries instantiates a new InlineResponse200124Deliveries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200124Deliveries() *InlineResponse200124Deliveries {
	this := InlineResponse200124Deliveries{}
	return &this
}

// NewInlineResponse200124DeliveriesWithDefaults instantiates a new InlineResponse200124Deliveries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200124DeliveriesWithDefaults() *InlineResponse200124Deliveries {
	this := InlineResponse200124Deliveries{}
	return &this
}

// GetDateCalled returns the DateCalled field value if set, zero value otherwise.
func (o *InlineResponse200124Deliveries) GetDateCalled() string {
	if o == nil || o.DateCalled == nil {
		var ret string
		return ret
	}
	return *o.DateCalled
}

// GetDateCalledOk returns a tuple with the DateCalled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200124Deliveries) GetDateCalledOk() (*string, bool) {
	if o == nil || o.DateCalled == nil {
		return nil, false
	}
	return o.DateCalled, true
}

// HasDateCalled returns a boolean if a field has been set.
func (o *InlineResponse200124Deliveries) HasDateCalled() bool {
	if o != nil && o.DateCalled != nil {
		return true
	}

	return false
}

// SetDateCalled gets a reference to the given string and assigns it to the DateCalled field.
func (o *InlineResponse200124Deliveries) SetDateCalled(v string) {
	o.DateCalled = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *InlineResponse200124Deliveries) GetEvent() string {
	if o == nil || o.Event == nil {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200124Deliveries) GetEventOk() (*string, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *InlineResponse200124Deliveries) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *InlineResponse200124Deliveries) SetEvent(v string) {
	o.Event = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200124Deliveries) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200124Deliveries) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200124Deliveries) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200124Deliveries) SetId(v string) {
	o.Id = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *InlineResponse200124Deliveries) GetRequest() InlineResponse200124Request {
	if o == nil || o.Request == nil {
		var ret InlineResponse200124Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200124Deliveries) GetRequestOk() (*InlineResponse200124Request, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *InlineResponse200124Deliveries) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given InlineResponse200124Request and assigns it to the Request field.
func (o *InlineResponse200124Deliveries) SetRequest(v InlineResponse200124Request) {
	o.Request = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *InlineResponse200124Deliveries) GetResponse() InlineResponse200124Response {
	if o == nil || o.Response == nil {
		var ret InlineResponse200124Response
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200124Deliveries) GetResponseOk() (*InlineResponse200124Response, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *InlineResponse200124Deliveries) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given InlineResponse200124Response and assigns it to the Response field.
func (o *InlineResponse200124Deliveries) SetResponse(v InlineResponse200124Response) {
	o.Response = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InlineResponse200124Deliveries) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200124Deliveries) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InlineResponse200124Deliveries) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *InlineResponse200124Deliveries) SetUrl(v string) {
	o.Url = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *InlineResponse200124Deliveries) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200124Deliveries) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *InlineResponse200124Deliveries) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *InlineResponse200124Deliveries) SetUuid(v string) {
	o.Uuid = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *InlineResponse200124Deliveries) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200124Deliveries) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *InlineResponse200124Deliveries) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *InlineResponse200124Deliveries) SetVersion(v string) {
	o.Version = &v
}

func (o InlineResponse200124Deliveries) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DateCalled != nil {
		toSerialize["dateCalled"] = o.DateCalled
	}
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200124Deliveries struct {
	value *InlineResponse200124Deliveries
	isSet bool
}

func (v NullableInlineResponse200124Deliveries) Get() *InlineResponse200124Deliveries {
	return v.value
}

func (v *NullableInlineResponse200124Deliveries) Set(val *InlineResponse200124Deliveries) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200124Deliveries) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200124Deliveries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200124Deliveries(val *InlineResponse200124Deliveries) *NullableInlineResponse200124Deliveries {
	return &NullableInlineResponse200124Deliveries{value: val, isSet: true}
}

func (v NullableInlineResponse200124Deliveries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200124Deliveries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


