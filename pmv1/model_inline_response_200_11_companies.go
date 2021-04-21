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

// InlineResponse20011Companies struct for InlineResponse20011Companies
type InlineResponse20011Companies struct {
	Accounts *string `json:"accounts,omitempty"`
	AddressOne *string `json:"address_one,omitempty"`
	AddressTwo *string `json:"address_two,omitempty"`
	CanSeePrivate *bool `json:"can_see_private,omitempty"`
	Cid *string `json:"cid,omitempty"`
	City *string `json:"city,omitempty"`
	CompanyNameUrl *string `json:"company_name_url,omitempty"`
	Contacts *string `json:"contacts,omitempty"`
	Country *string `json:"country,omitempty"`
	Countrycode *string `json:"countrycode,omitempty"`
	EmailOne *string `json:"email_one,omitempty"`
	EmailThree *string `json:"email_three,omitempty"`
	EmailTwo *string `json:"email_two,omitempty"`
	Fax *string `json:"fax,omitempty"`
	Id *string `json:"id,omitempty"`
	Industry *string `json:"industry,omitempty"`
	Isowner *string `json:"isowner,omitempty"`
	LogoURL *string `json:"logo-URL,omitempty"`
	Name *string `json:"name,omitempty"`
	Phone *string `json:"phone,omitempty"`
	State *string `json:"state,omitempty"`
	Website *string `json:"website,omitempty"`
	Zip *string `json:"zip,omitempty"`
}

// NewInlineResponse20011Companies instantiates a new InlineResponse20011Companies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20011Companies() *InlineResponse20011Companies {
	this := InlineResponse20011Companies{}
	return &this
}

// NewInlineResponse20011CompaniesWithDefaults instantiates a new InlineResponse20011Companies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20011CompaniesWithDefaults() *InlineResponse20011Companies {
	this := InlineResponse20011Companies{}
	return &this
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *InlineResponse20011Companies) GetAccounts() string {
	if o == nil || o.Accounts == nil {
		var ret string
		return ret
	}
	return *o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011Companies) GetAccountsOk() (*string, bool) {
	if o == nil || o.Accounts == nil {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *InlineResponse20011Companies) HasAccounts() bool {
	if o != nil && o.Accounts != nil {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given string and assigns it to the Accounts field.
func (o *InlineResponse20011Companies) SetAccounts(v string) {
	o.Accounts = &v
}

// GetAddressOne returns the AddressOne field value if set, zero value otherwise.
func (o *InlineResponse20011Companies) GetAddressOne() string {
	if o == nil || o.AddressOne == nil {
		var ret string
		return ret
	}
	return *o.AddressOne
}

// GetAddressOneOk returns a tuple with the AddressOne field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011Companies) GetAddressOneOk() (*string, bool) {
	if o == nil || o.AddressOne == nil {
		return nil, false
	}
	return o.AddressOne, true
}

// HasAddressOne returns a boolean if a field has been set.
func (o *InlineResponse20011Companies) HasAddressOne() bool {
	if o != nil && o.AddressOne != nil {
		return true
	}

	return false
}

// SetAddressOne gets a reference to the given string and assigns it to the AddressOne field.
func (o *InlineResponse20011Companies) SetAddressOne(v string) {
	o.AddressOne = &v
}

// GetAddressTwo returns the AddressTwo field value if set, zero value otherwise.
func (o *InlineResponse20011Companies) GetAddressTwo() string {
	if o == nil || o.AddressTwo == nil {
		var ret string
		return ret
	}
	return *o.AddressTwo
}

// GetAddressTwoOk returns a tuple with the AddressTwo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011Companies) GetAddressTwoOk() (*string, bool) {
	if o == nil || o.AddressTwo == nil {
		return nil, false
	}
	return o.AddressTwo, true
}

// HasAddressTwo returns a boolean if a field has been set.
func (o *InlineResponse20011Companies) HasAddressTwo() bool {
	if o != nil && o.AddressTwo != nil {
		return true
	}

	return false
}

// SetAddressTwo gets a reference to the given string and assigns it to the AddressTwo field.
func (o *InlineResponse20011Companies) SetAddressTwo(v string) {
	o.AddressTwo = &v
}

// GetCanSeePrivate returns the CanSeePrivate field value if set, zero value otherwise.
func (o *InlineResponse20011Companies) GetCanSeePrivate() bool {
	if o == nil || o.CanSeePrivate == nil {
		var ret bool
		return ret
	}
	return *o.CanSeePrivate
}

// GetCanSeePrivateOk returns a tuple with the CanSeePrivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011Companies) GetCanSeePrivateOk() (*bool, bool) {
	if o == nil || o.CanSeePrivate == nil {
		return nil, false
	}
	return o.CanSeePrivate, true
}

// HasCanSeePrivate returns a boolean if a field has been set.
func (o *InlineResponse20011Companies) HasCanSeePrivate() bool {
	if o != nil && o.CanSeePrivate != nil {
		return true
	}

	return false
}

// SetCanSeePrivate gets a reference to the given bool and assigns it to the CanSeePrivate field.
func (o *InlineResponse20011Companies) SetCanSeePrivate(v bool) {
	o.CanSeePrivate = &v
}

// GetCid returns the Cid field value if set, zero value otherwise.
func (o *InlineResponse20011Companies) GetCid() string {
	if o == nil || o.Cid == nil {
		var ret string
		return ret
	}
	return *o.Cid
}

// GetCidOk returns a tuple with the Cid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011Companies) GetCidOk() (*string, bool) {
	if o == nil || o.Cid == nil {
		return nil, false
	}
	return o.Cid, true
}

// HasCid returns a boolean if a field has been set.
func (o *InlineResponse20011Companies) HasCid() bool {
	if o != nil && o.Cid != nil {
		return true
	}

	return false
}

// SetCid gets a reference to the given string and assigns it to the Cid field.
func (o *InlineResponse20011Companies) SetCid(v string) {
	o.Cid = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *InlineResponse20011Companies) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011Companies) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *InlineResponse20011Companies) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *InlineResponse20011Companies) SetCity(v string) {
	o.City = &v
}

// GetCompanyNameUrl returns the CompanyNameUrl field value if set, zero value otherwise.
func (o *InlineResponse20011Companies) GetCompanyNameUrl() string {
	if o == nil || o.CompanyNameUrl == nil {
		var ret string
		return ret
	}
	return *o.CompanyNameUrl
}

// GetCompanyNameUrlOk returns a tuple with the CompanyNameUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011Companies) GetCompanyNameUrlOk() (*string, bool) {
	if o == nil || o.CompanyNameUrl == nil {
		return nil, false
	}
	return o.CompanyNameUrl, true
}

// HasCompanyNameUrl returns a boolean if a field has been set.
func (o *InlineResponse20011Companies) HasCompanyNameUrl() bool {
	if o != nil && o.CompanyNameUrl != nil {
		return true
	}

	return false
}

// SetCompanyNameUrl gets a reference to the given string and assigns it to the CompanyNameUrl field.
func (o *InlineResponse20011Companies) SetCompanyNameUrl(v string) {
	o.CompanyNameUrl = &v
}

// GetContacts returns the Contacts field value if set, zero value otherwise.
func (o *InlineResponse20011Companies) GetContacts() string {
	if o == nil || o.Contacts == nil {
		var ret string
		return ret
	}
	return *o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011Companies) GetContactsOk() (*string, bool) {
	if o == nil || o.Contacts == nil {
		return nil, false
	}
	return o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *InlineResponse20011Companies) HasContacts() bool {
	if o != nil && o.Contacts != nil {
		return true
	}

	return false
}

// SetContacts gets a reference to the given string and assigns it to the Contacts field.
func (o *InlineResponse20011Companies) SetContacts(v string) {
	o.Contacts = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *InlineResponse20011Companies) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011Companies) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *InlineResponse20011Companies) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *InlineResponse20011Companies) SetCountry(v string) {
	o.Country = &v
}

// GetCountrycode returns the Countrycode field value if set, zero value otherwise.
func (o *InlineResponse20011Companies) GetCountrycode() string {
	if o == nil || o.Countrycode == nil {
		var ret string
		return ret
	}
	return *o.Countrycode
}

// GetCountrycodeOk returns a tuple with the Countrycode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011Companies) GetCountrycodeOk() (*string, bool) {
	if o == nil || o.Countrycode == nil {
		return nil, false
	}
	return o.Countrycode, true
}

// HasCountrycode returns a boolean if a field has been set.
func (o *InlineResponse20011Companies) HasCountrycode() bool {
	if o != nil && o.Countrycode != nil {
		return true
	}

	return false
}

// SetCountrycode gets a reference to the given string and assigns it to the Countrycode field.
func (o *InlineResponse20011Companies) SetCountrycode(v string) {
	o.Countrycode = &v
}

// GetEmailOne returns the EmailOne field value if set, zero value otherwise.
func (o *InlineResponse20011Companies) GetEmailOne() string {
	if o == nil || o.EmailOne == nil {
		var ret string
		return ret
	}
	return *o.EmailOne
}

// GetEmailOneOk returns a tuple with the EmailOne field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011Companies) GetEmailOneOk() (*string, bool) {
	if o == nil || o.EmailOne == nil {
		return nil, false
	}
	return o.EmailOne, true
}

// HasEmailOne returns a boolean if a field has been set.
func (o *InlineResponse20011Companies) HasEmailOne() bool {
	if o != nil && o.EmailOne != nil {
		return true
	}

	return false
}

// SetEmailOne gets a reference to the given string and assigns it to the EmailOne field.
func (o *InlineResponse20011Companies) SetEmailOne(v string) {
	o.EmailOne = &v
}

// GetEmailThree returns the EmailThree field value if set, zero value otherwise.
func (o *InlineResponse20011Companies) GetEmailThree() string {
	if o == nil || o.EmailThree == nil {
		var ret string
		return ret
	}
	return *o.EmailThree
}

// GetEmailThreeOk returns a tuple with the EmailThree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011Companies) GetEmailThreeOk() (*string, bool) {
	if o == nil || o.EmailThree == nil {
		return nil, false
	}
	return o.EmailThree, true
}

// HasEmailThree returns a boolean if a field has been set.
func (o *InlineResponse20011Companies) HasEmailThree() bool {
	if o != nil && o.EmailThree != nil {
		return true
	}

	return false
}

// SetEmailThree gets a reference to the given string and assigns it to the EmailThree field.
func (o *InlineResponse20011Companies) SetEmailThree(v string) {
	o.EmailThree = &v
}

// GetEmailTwo returns the EmailTwo field value if set, zero value otherwise.
func (o *InlineResponse20011Companies) GetEmailTwo() string {
	if o == nil || o.EmailTwo == nil {
		var ret string
		return ret
	}
	return *o.EmailTwo
}

// GetEmailTwoOk returns a tuple with the EmailTwo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011Companies) GetEmailTwoOk() (*string, bool) {
	if o == nil || o.EmailTwo == nil {
		return nil, false
	}
	return o.EmailTwo, true
}

// HasEmailTwo returns a boolean if a field has been set.
func (o *InlineResponse20011Companies) HasEmailTwo() bool {
	if o != nil && o.EmailTwo != nil {
		return true
	}

	return false
}

// SetEmailTwo gets a reference to the given string and assigns it to the EmailTwo field.
func (o *InlineResponse20011Companies) SetEmailTwo(v string) {
	o.EmailTwo = &v
}

// GetFax returns the Fax field value if set, zero value otherwise.
func (o *InlineResponse20011Companies) GetFax() string {
	if o == nil || o.Fax == nil {
		var ret string
		return ret
	}
	return *o.Fax
}

// GetFaxOk returns a tuple with the Fax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011Companies) GetFaxOk() (*string, bool) {
	if o == nil || o.Fax == nil {
		return nil, false
	}
	return o.Fax, true
}

// HasFax returns a boolean if a field has been set.
func (o *InlineResponse20011Companies) HasFax() bool {
	if o != nil && o.Fax != nil {
		return true
	}

	return false
}

// SetFax gets a reference to the given string and assigns it to the Fax field.
func (o *InlineResponse20011Companies) SetFax(v string) {
	o.Fax = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20011Companies) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011Companies) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20011Companies) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20011Companies) SetId(v string) {
	o.Id = &v
}

// GetIndustry returns the Industry field value if set, zero value otherwise.
func (o *InlineResponse20011Companies) GetIndustry() string {
	if o == nil || o.Industry == nil {
		var ret string
		return ret
	}
	return *o.Industry
}

// GetIndustryOk returns a tuple with the Industry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011Companies) GetIndustryOk() (*string, bool) {
	if o == nil || o.Industry == nil {
		return nil, false
	}
	return o.Industry, true
}

// HasIndustry returns a boolean if a field has been set.
func (o *InlineResponse20011Companies) HasIndustry() bool {
	if o != nil && o.Industry != nil {
		return true
	}

	return false
}

// SetIndustry gets a reference to the given string and assigns it to the Industry field.
func (o *InlineResponse20011Companies) SetIndustry(v string) {
	o.Industry = &v
}

// GetIsowner returns the Isowner field value if set, zero value otherwise.
func (o *InlineResponse20011Companies) GetIsowner() string {
	if o == nil || o.Isowner == nil {
		var ret string
		return ret
	}
	return *o.Isowner
}

// GetIsownerOk returns a tuple with the Isowner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011Companies) GetIsownerOk() (*string, bool) {
	if o == nil || o.Isowner == nil {
		return nil, false
	}
	return o.Isowner, true
}

// HasIsowner returns a boolean if a field has been set.
func (o *InlineResponse20011Companies) HasIsowner() bool {
	if o != nil && o.Isowner != nil {
		return true
	}

	return false
}

// SetIsowner gets a reference to the given string and assigns it to the Isowner field.
func (o *InlineResponse20011Companies) SetIsowner(v string) {
	o.Isowner = &v
}

// GetLogoURL returns the LogoURL field value if set, zero value otherwise.
func (o *InlineResponse20011Companies) GetLogoURL() string {
	if o == nil || o.LogoURL == nil {
		var ret string
		return ret
	}
	return *o.LogoURL
}

// GetLogoURLOk returns a tuple with the LogoURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011Companies) GetLogoURLOk() (*string, bool) {
	if o == nil || o.LogoURL == nil {
		return nil, false
	}
	return o.LogoURL, true
}

// HasLogoURL returns a boolean if a field has been set.
func (o *InlineResponse20011Companies) HasLogoURL() bool {
	if o != nil && o.LogoURL != nil {
		return true
	}

	return false
}

// SetLogoURL gets a reference to the given string and assigns it to the LogoURL field.
func (o *InlineResponse20011Companies) SetLogoURL(v string) {
	o.LogoURL = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20011Companies) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011Companies) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20011Companies) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20011Companies) SetName(v string) {
	o.Name = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *InlineResponse20011Companies) GetPhone() string {
	if o == nil || o.Phone == nil {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011Companies) GetPhoneOk() (*string, bool) {
	if o == nil || o.Phone == nil {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *InlineResponse20011Companies) HasPhone() bool {
	if o != nil && o.Phone != nil {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *InlineResponse20011Companies) SetPhone(v string) {
	o.Phone = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *InlineResponse20011Companies) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011Companies) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *InlineResponse20011Companies) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *InlineResponse20011Companies) SetState(v string) {
	o.State = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *InlineResponse20011Companies) GetWebsite() string {
	if o == nil || o.Website == nil {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011Companies) GetWebsiteOk() (*string, bool) {
	if o == nil || o.Website == nil {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *InlineResponse20011Companies) HasWebsite() bool {
	if o != nil && o.Website != nil {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *InlineResponse20011Companies) SetWebsite(v string) {
	o.Website = &v
}

// GetZip returns the Zip field value if set, zero value otherwise.
func (o *InlineResponse20011Companies) GetZip() string {
	if o == nil || o.Zip == nil {
		var ret string
		return ret
	}
	return *o.Zip
}

// GetZipOk returns a tuple with the Zip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011Companies) GetZipOk() (*string, bool) {
	if o == nil || o.Zip == nil {
		return nil, false
	}
	return o.Zip, true
}

// HasZip returns a boolean if a field has been set.
func (o *InlineResponse20011Companies) HasZip() bool {
	if o != nil && o.Zip != nil {
		return true
	}

	return false
}

// SetZip gets a reference to the given string and assigns it to the Zip field.
func (o *InlineResponse20011Companies) SetZip(v string) {
	o.Zip = &v
}

func (o InlineResponse20011Companies) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Accounts != nil {
		toSerialize["accounts"] = o.Accounts
	}
	if o.AddressOne != nil {
		toSerialize["address_one"] = o.AddressOne
	}
	if o.AddressTwo != nil {
		toSerialize["address_two"] = o.AddressTwo
	}
	if o.CanSeePrivate != nil {
		toSerialize["can_see_private"] = o.CanSeePrivate
	}
	if o.Cid != nil {
		toSerialize["cid"] = o.Cid
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.CompanyNameUrl != nil {
		toSerialize["company_name_url"] = o.CompanyNameUrl
	}
	if o.Contacts != nil {
		toSerialize["contacts"] = o.Contacts
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.Countrycode != nil {
		toSerialize["countrycode"] = o.Countrycode
	}
	if o.EmailOne != nil {
		toSerialize["email_one"] = o.EmailOne
	}
	if o.EmailThree != nil {
		toSerialize["email_three"] = o.EmailThree
	}
	if o.EmailTwo != nil {
		toSerialize["email_two"] = o.EmailTwo
	}
	if o.Fax != nil {
		toSerialize["fax"] = o.Fax
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Industry != nil {
		toSerialize["industry"] = o.Industry
	}
	if o.Isowner != nil {
		toSerialize["isowner"] = o.Isowner
	}
	if o.LogoURL != nil {
		toSerialize["logo-URL"] = o.LogoURL
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Phone != nil {
		toSerialize["phone"] = o.Phone
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Website != nil {
		toSerialize["website"] = o.Website
	}
	if o.Zip != nil {
		toSerialize["zip"] = o.Zip
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20011Companies struct {
	value *InlineResponse20011Companies
	isSet bool
}

func (v NullableInlineResponse20011Companies) Get() *InlineResponse20011Companies {
	return v.value
}

func (v *NullableInlineResponse20011Companies) Set(val *InlineResponse20011Companies) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20011Companies) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20011Companies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20011Companies(val *InlineResponse20011Companies) *NullableInlineResponse20011Companies {
	return &NullableInlineResponse20011Companies{value: val, isSet: true}
}

func (v NullableInlineResponse20011Companies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20011Companies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

