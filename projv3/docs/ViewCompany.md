# ViewCompany

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to **int32** |  | [optional] 
**AddressOne** | Pointer to **string** |  | [optional] 
**AddressTwo** | Pointer to **string** |  | [optional] 
**CanSeePrivate** | Pointer to **bool** |  | [optional] 
**Cid** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**ClientUsers** | Pointer to **int32** |  | [optional] 
**CompanyNameUrl** | Pointer to **string** |  | [optional] 
**Contacts** | Pointer to **int32** |  | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**EmailOne** | Pointer to **string** |  | [optional] 
**EmailThree** | Pointer to **string** |  | [optional] 
**EmailTwo** | Pointer to **string** |  | [optional] 
**Fax** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Industry** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**IndustryId** | Pointer to **int32** |  | [optional] 
**IsOwner** | Pointer to **bool** |  | [optional] 
**LogoUrl** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**PrivateNotesText** | Pointer to **string** |  | [optional] 
**ProfileText** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Website** | Pointer to **string** |  | [optional] 
**Zip** | Pointer to **string** |  | [optional] 

## Methods

### NewViewCompany

`func NewViewCompany() *ViewCompany`

NewViewCompany instantiates a new ViewCompany object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewCompanyWithDefaults

`func NewViewCompanyWithDefaults() *ViewCompany`

NewViewCompanyWithDefaults instantiates a new ViewCompany object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *ViewCompany) GetAccounts() int32`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *ViewCompany) GetAccountsOk() (*int32, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *ViewCompany) SetAccounts(v int32)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *ViewCompany) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetAddressOne

`func (o *ViewCompany) GetAddressOne() string`

GetAddressOne returns the AddressOne field if non-nil, zero value otherwise.

### GetAddressOneOk

`func (o *ViewCompany) GetAddressOneOk() (*string, bool)`

GetAddressOneOk returns a tuple with the AddressOne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressOne

`func (o *ViewCompany) SetAddressOne(v string)`

SetAddressOne sets AddressOne field to given value.

### HasAddressOne

`func (o *ViewCompany) HasAddressOne() bool`

HasAddressOne returns a boolean if a field has been set.

### GetAddressTwo

`func (o *ViewCompany) GetAddressTwo() string`

GetAddressTwo returns the AddressTwo field if non-nil, zero value otherwise.

### GetAddressTwoOk

`func (o *ViewCompany) GetAddressTwoOk() (*string, bool)`

GetAddressTwoOk returns a tuple with the AddressTwo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressTwo

`func (o *ViewCompany) SetAddressTwo(v string)`

SetAddressTwo sets AddressTwo field to given value.

### HasAddressTwo

`func (o *ViewCompany) HasAddressTwo() bool`

HasAddressTwo returns a boolean if a field has been set.

### GetCanSeePrivate

`func (o *ViewCompany) GetCanSeePrivate() bool`

GetCanSeePrivate returns the CanSeePrivate field if non-nil, zero value otherwise.

### GetCanSeePrivateOk

`func (o *ViewCompany) GetCanSeePrivateOk() (*bool, bool)`

GetCanSeePrivateOk returns a tuple with the CanSeePrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSeePrivate

`func (o *ViewCompany) SetCanSeePrivate(v bool)`

SetCanSeePrivate sets CanSeePrivate field to given value.

### HasCanSeePrivate

`func (o *ViewCompany) HasCanSeePrivate() bool`

HasCanSeePrivate returns a boolean if a field has been set.

### GetCid

`func (o *ViewCompany) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *ViewCompany) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *ViewCompany) SetCid(v string)`

SetCid sets Cid field to given value.

### HasCid

`func (o *ViewCompany) HasCid() bool`

HasCid returns a boolean if a field has been set.

### GetCity

`func (o *ViewCompany) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ViewCompany) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ViewCompany) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *ViewCompany) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetClientUsers

`func (o *ViewCompany) GetClientUsers() int32`

GetClientUsers returns the ClientUsers field if non-nil, zero value otherwise.

### GetClientUsersOk

`func (o *ViewCompany) GetClientUsersOk() (*int32, bool)`

GetClientUsersOk returns a tuple with the ClientUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUsers

`func (o *ViewCompany) SetClientUsers(v int32)`

SetClientUsers sets ClientUsers field to given value.

### HasClientUsers

`func (o *ViewCompany) HasClientUsers() bool`

HasClientUsers returns a boolean if a field has been set.

### GetCompanyNameUrl

`func (o *ViewCompany) GetCompanyNameUrl() string`

GetCompanyNameUrl returns the CompanyNameUrl field if non-nil, zero value otherwise.

### GetCompanyNameUrlOk

`func (o *ViewCompany) GetCompanyNameUrlOk() (*string, bool)`

GetCompanyNameUrlOk returns a tuple with the CompanyNameUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyNameUrl

`func (o *ViewCompany) SetCompanyNameUrl(v string)`

SetCompanyNameUrl sets CompanyNameUrl field to given value.

### HasCompanyNameUrl

`func (o *ViewCompany) HasCompanyNameUrl() bool`

HasCompanyNameUrl returns a boolean if a field has been set.

### GetContacts

`func (o *ViewCompany) GetContacts() int32`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *ViewCompany) GetContactsOk() (*int32, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *ViewCompany) SetContacts(v int32)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *ViewCompany) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetCountryCode

`func (o *ViewCompany) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *ViewCompany) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *ViewCompany) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *ViewCompany) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetEmailOne

`func (o *ViewCompany) GetEmailOne() string`

GetEmailOne returns the EmailOne field if non-nil, zero value otherwise.

### GetEmailOneOk

`func (o *ViewCompany) GetEmailOneOk() (*string, bool)`

GetEmailOneOk returns a tuple with the EmailOne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOne

`func (o *ViewCompany) SetEmailOne(v string)`

SetEmailOne sets EmailOne field to given value.

### HasEmailOne

`func (o *ViewCompany) HasEmailOne() bool`

HasEmailOne returns a boolean if a field has been set.

### GetEmailThree

`func (o *ViewCompany) GetEmailThree() string`

GetEmailThree returns the EmailThree field if non-nil, zero value otherwise.

### GetEmailThreeOk

`func (o *ViewCompany) GetEmailThreeOk() (*string, bool)`

GetEmailThreeOk returns a tuple with the EmailThree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailThree

`func (o *ViewCompany) SetEmailThree(v string)`

SetEmailThree sets EmailThree field to given value.

### HasEmailThree

`func (o *ViewCompany) HasEmailThree() bool`

HasEmailThree returns a boolean if a field has been set.

### GetEmailTwo

`func (o *ViewCompany) GetEmailTwo() string`

GetEmailTwo returns the EmailTwo field if non-nil, zero value otherwise.

### GetEmailTwoOk

`func (o *ViewCompany) GetEmailTwoOk() (*string, bool)`

GetEmailTwoOk returns a tuple with the EmailTwo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTwo

`func (o *ViewCompany) SetEmailTwo(v string)`

SetEmailTwo sets EmailTwo field to given value.

### HasEmailTwo

`func (o *ViewCompany) HasEmailTwo() bool`

HasEmailTwo returns a boolean if a field has been set.

### GetFax

`func (o *ViewCompany) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *ViewCompany) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *ViewCompany) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *ViewCompany) HasFax() bool`

HasFax returns a boolean if a field has been set.

### GetId

`func (o *ViewCompany) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewCompany) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewCompany) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewCompany) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndustry

`func (o *ViewCompany) GetIndustry() ViewRelationship`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *ViewCompany) GetIndustryOk() (*ViewRelationship, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *ViewCompany) SetIndustry(v ViewRelationship)`

SetIndustry sets Industry field to given value.

### HasIndustry

`func (o *ViewCompany) HasIndustry() bool`

HasIndustry returns a boolean if a field has been set.

### GetIndustryId

`func (o *ViewCompany) GetIndustryId() int32`

GetIndustryId returns the IndustryId field if non-nil, zero value otherwise.

### GetIndustryIdOk

`func (o *ViewCompany) GetIndustryIdOk() (*int32, bool)`

GetIndustryIdOk returns a tuple with the IndustryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryId

`func (o *ViewCompany) SetIndustryId(v int32)`

SetIndustryId sets IndustryId field to given value.

### HasIndustryId

`func (o *ViewCompany) HasIndustryId() bool`

HasIndustryId returns a boolean if a field has been set.

### GetIsOwner

`func (o *ViewCompany) GetIsOwner() bool`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *ViewCompany) GetIsOwnerOk() (*bool, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *ViewCompany) SetIsOwner(v bool)`

SetIsOwner sets IsOwner field to given value.

### HasIsOwner

`func (o *ViewCompany) HasIsOwner() bool`

HasIsOwner returns a boolean if a field has been set.

### GetLogoUrl

`func (o *ViewCompany) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *ViewCompany) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *ViewCompany) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *ViewCompany) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetName

`func (o *ViewCompany) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ViewCompany) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ViewCompany) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ViewCompany) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhone

`func (o *ViewCompany) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ViewCompany) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ViewCompany) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ViewCompany) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPrivateNotesText

`func (o *ViewCompany) GetPrivateNotesText() string`

GetPrivateNotesText returns the PrivateNotesText field if non-nil, zero value otherwise.

### GetPrivateNotesTextOk

`func (o *ViewCompany) GetPrivateNotesTextOk() (*string, bool)`

GetPrivateNotesTextOk returns a tuple with the PrivateNotesText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNotesText

`func (o *ViewCompany) SetPrivateNotesText(v string)`

SetPrivateNotesText sets PrivateNotesText field to given value.

### HasPrivateNotesText

`func (o *ViewCompany) HasPrivateNotesText() bool`

HasPrivateNotesText returns a boolean if a field has been set.

### GetProfileText

`func (o *ViewCompany) GetProfileText() string`

GetProfileText returns the ProfileText field if non-nil, zero value otherwise.

### GetProfileTextOk

`func (o *ViewCompany) GetProfileTextOk() (*string, bool)`

GetProfileTextOk returns a tuple with the ProfileText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileText

`func (o *ViewCompany) SetProfileText(v string)`

SetProfileText sets ProfileText field to given value.

### HasProfileText

`func (o *ViewCompany) HasProfileText() bool`

HasProfileText returns a boolean if a field has been set.

### GetState

`func (o *ViewCompany) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ViewCompany) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ViewCompany) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ViewCompany) HasState() bool`

HasState returns a boolean if a field has been set.

### GetWebsite

`func (o *ViewCompany) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *ViewCompany) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *ViewCompany) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *ViewCompany) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetZip

`func (o *ViewCompany) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *ViewCompany) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *ViewCompany) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *ViewCompany) HasZip() bool`

HasZip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


