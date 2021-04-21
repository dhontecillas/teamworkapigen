# InlineResponse20027Person

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**InlineResponse20013Address**](InlineResponse20013Address.md) |  | [optional] 
**Administrator** | Pointer to **bool** |  | [optional] 
**AvatarUrl** | Pointer to **string** |  | [optional] 
**CompanyId** | Pointer to **string** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**EmailAddress** | Pointer to **string** |  | [optional] 
**EmailAlt1** | Pointer to **string** |  | [optional] 
**EmailAlt2** | Pointer to **string** |  | [optional] 
**EmailAlt3** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**HasAccessToNewProjects** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ImHandle** | Pointer to **string** |  | [optional] 
**ImService** | Pointer to **string** |  | [optional] 
**InOwnerCompany** | Pointer to **string** |  | [optional] 
**LastChangedOn** | Pointer to **string** |  | [optional] 
**LastLogin** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to [**InlineResponse20027PersonPermissions**](InlineResponse20027PersonPermissions.md) |  | [optional] 
**PhoneNumberFax** | Pointer to **string** |  | [optional] 
**PhoneNumberHome** | Pointer to **string** |  | [optional] 
**PhoneNumberMobile** | Pointer to **string** |  | [optional] 
**PhoneNumberOffice** | Pointer to **string** |  | [optional] 
**PhoneNumberOfficeExt** | Pointer to **string** |  | [optional] 
**Pid** | Pointer to **string** |  | [optional] 
**PrivateNotes** | Pointer to **string** |  | [optional] 
**Profile** | Pointer to **string** |  | [optional] 
**SiteOwner** | Pointer to **bool** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Twitter** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**UserType** | Pointer to **string** |  | [optional] 
**UserUUID** | Pointer to **string** |  | [optional] 

## Methods

### NewInlineResponse20027Person

`func NewInlineResponse20027Person() *InlineResponse20027Person`

NewInlineResponse20027Person instantiates a new InlineResponse20027Person object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20027PersonWithDefaults

`func NewInlineResponse20027PersonWithDefaults() *InlineResponse20027Person`

NewInlineResponse20027PersonWithDefaults instantiates a new InlineResponse20027Person object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *InlineResponse20027Person) GetAddress() InlineResponse20013Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InlineResponse20027Person) GetAddressOk() (*InlineResponse20013Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InlineResponse20027Person) SetAddress(v InlineResponse20013Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *InlineResponse20027Person) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAdministrator

`func (o *InlineResponse20027Person) GetAdministrator() bool`

GetAdministrator returns the Administrator field if non-nil, zero value otherwise.

### GetAdministratorOk

`func (o *InlineResponse20027Person) GetAdministratorOk() (*bool, bool)`

GetAdministratorOk returns a tuple with the Administrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrator

`func (o *InlineResponse20027Person) SetAdministrator(v bool)`

SetAdministrator sets Administrator field to given value.

### HasAdministrator

`func (o *InlineResponse20027Person) HasAdministrator() bool`

HasAdministrator returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *InlineResponse20027Person) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *InlineResponse20027Person) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *InlineResponse20027Person) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *InlineResponse20027Person) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetCompanyId

`func (o *InlineResponse20027Person) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *InlineResponse20027Person) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *InlineResponse20027Person) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *InlineResponse20027Person) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCompanyName

`func (o *InlineResponse20027Person) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *InlineResponse20027Person) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *InlineResponse20027Person) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *InlineResponse20027Person) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse20027Person) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse20027Person) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse20027Person) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse20027Person) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeleted

`func (o *InlineResponse20027Person) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *InlineResponse20027Person) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *InlineResponse20027Person) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *InlineResponse20027Person) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetEmailAddress

`func (o *InlineResponse20027Person) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *InlineResponse20027Person) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *InlineResponse20027Person) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *InlineResponse20027Person) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetEmailAlt1

`func (o *InlineResponse20027Person) GetEmailAlt1() string`

GetEmailAlt1 returns the EmailAlt1 field if non-nil, zero value otherwise.

### GetEmailAlt1Ok

`func (o *InlineResponse20027Person) GetEmailAlt1Ok() (*string, bool)`

GetEmailAlt1Ok returns a tuple with the EmailAlt1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAlt1

`func (o *InlineResponse20027Person) SetEmailAlt1(v string)`

SetEmailAlt1 sets EmailAlt1 field to given value.

### HasEmailAlt1

`func (o *InlineResponse20027Person) HasEmailAlt1() bool`

HasEmailAlt1 returns a boolean if a field has been set.

### GetEmailAlt2

`func (o *InlineResponse20027Person) GetEmailAlt2() string`

GetEmailAlt2 returns the EmailAlt2 field if non-nil, zero value otherwise.

### GetEmailAlt2Ok

`func (o *InlineResponse20027Person) GetEmailAlt2Ok() (*string, bool)`

GetEmailAlt2Ok returns a tuple with the EmailAlt2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAlt2

`func (o *InlineResponse20027Person) SetEmailAlt2(v string)`

SetEmailAlt2 sets EmailAlt2 field to given value.

### HasEmailAlt2

`func (o *InlineResponse20027Person) HasEmailAlt2() bool`

HasEmailAlt2 returns a boolean if a field has been set.

### GetEmailAlt3

`func (o *InlineResponse20027Person) GetEmailAlt3() string`

GetEmailAlt3 returns the EmailAlt3 field if non-nil, zero value otherwise.

### GetEmailAlt3Ok

`func (o *InlineResponse20027Person) GetEmailAlt3Ok() (*string, bool)`

GetEmailAlt3Ok returns a tuple with the EmailAlt3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAlt3

`func (o *InlineResponse20027Person) SetEmailAlt3(v string)`

SetEmailAlt3 sets EmailAlt3 field to given value.

### HasEmailAlt3

`func (o *InlineResponse20027Person) HasEmailAlt3() bool`

HasEmailAlt3 returns a boolean if a field has been set.

### GetFirstName

`func (o *InlineResponse20027Person) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *InlineResponse20027Person) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *InlineResponse20027Person) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *InlineResponse20027Person) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetHasAccessToNewProjects

`func (o *InlineResponse20027Person) GetHasAccessToNewProjects() bool`

GetHasAccessToNewProjects returns the HasAccessToNewProjects field if non-nil, zero value otherwise.

### GetHasAccessToNewProjectsOk

`func (o *InlineResponse20027Person) GetHasAccessToNewProjectsOk() (*bool, bool)`

GetHasAccessToNewProjectsOk returns a tuple with the HasAccessToNewProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccessToNewProjects

`func (o *InlineResponse20027Person) SetHasAccessToNewProjects(v bool)`

SetHasAccessToNewProjects sets HasAccessToNewProjects field to given value.

### HasHasAccessToNewProjects

`func (o *InlineResponse20027Person) HasHasAccessToNewProjects() bool`

HasHasAccessToNewProjects returns a boolean if a field has been set.

### GetId

`func (o *InlineResponse20027Person) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20027Person) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20027Person) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20027Person) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImHandle

`func (o *InlineResponse20027Person) GetImHandle() string`

GetImHandle returns the ImHandle field if non-nil, zero value otherwise.

### GetImHandleOk

`func (o *InlineResponse20027Person) GetImHandleOk() (*string, bool)`

GetImHandleOk returns a tuple with the ImHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImHandle

`func (o *InlineResponse20027Person) SetImHandle(v string)`

SetImHandle sets ImHandle field to given value.

### HasImHandle

`func (o *InlineResponse20027Person) HasImHandle() bool`

HasImHandle returns a boolean if a field has been set.

### GetImService

`func (o *InlineResponse20027Person) GetImService() string`

GetImService returns the ImService field if non-nil, zero value otherwise.

### GetImServiceOk

`func (o *InlineResponse20027Person) GetImServiceOk() (*string, bool)`

GetImServiceOk returns a tuple with the ImService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImService

`func (o *InlineResponse20027Person) SetImService(v string)`

SetImService sets ImService field to given value.

### HasImService

`func (o *InlineResponse20027Person) HasImService() bool`

HasImService returns a boolean if a field has been set.

### GetInOwnerCompany

`func (o *InlineResponse20027Person) GetInOwnerCompany() string`

GetInOwnerCompany returns the InOwnerCompany field if non-nil, zero value otherwise.

### GetInOwnerCompanyOk

`func (o *InlineResponse20027Person) GetInOwnerCompanyOk() (*string, bool)`

GetInOwnerCompanyOk returns a tuple with the InOwnerCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInOwnerCompany

`func (o *InlineResponse20027Person) SetInOwnerCompany(v string)`

SetInOwnerCompany sets InOwnerCompany field to given value.

### HasInOwnerCompany

`func (o *InlineResponse20027Person) HasInOwnerCompany() bool`

HasInOwnerCompany returns a boolean if a field has been set.

### GetLastChangedOn

`func (o *InlineResponse20027Person) GetLastChangedOn() string`

GetLastChangedOn returns the LastChangedOn field if non-nil, zero value otherwise.

### GetLastChangedOnOk

`func (o *InlineResponse20027Person) GetLastChangedOnOk() (*string, bool)`

GetLastChangedOnOk returns a tuple with the LastChangedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangedOn

`func (o *InlineResponse20027Person) SetLastChangedOn(v string)`

SetLastChangedOn sets LastChangedOn field to given value.

### HasLastChangedOn

`func (o *InlineResponse20027Person) HasLastChangedOn() bool`

HasLastChangedOn returns a boolean if a field has been set.

### GetLastLogin

`func (o *InlineResponse20027Person) GetLastLogin() string`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *InlineResponse20027Person) GetLastLoginOk() (*string, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *InlineResponse20027Person) SetLastLogin(v string)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *InlineResponse20027Person) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetLastName

`func (o *InlineResponse20027Person) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *InlineResponse20027Person) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *InlineResponse20027Person) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *InlineResponse20027Person) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetPermissions

`func (o *InlineResponse20027Person) GetPermissions() InlineResponse20027PersonPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *InlineResponse20027Person) GetPermissionsOk() (*InlineResponse20027PersonPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *InlineResponse20027Person) SetPermissions(v InlineResponse20027PersonPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *InlineResponse20027Person) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetPhoneNumberFax

`func (o *InlineResponse20027Person) GetPhoneNumberFax() string`

GetPhoneNumberFax returns the PhoneNumberFax field if non-nil, zero value otherwise.

### GetPhoneNumberFaxOk

`func (o *InlineResponse20027Person) GetPhoneNumberFaxOk() (*string, bool)`

GetPhoneNumberFaxOk returns a tuple with the PhoneNumberFax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberFax

`func (o *InlineResponse20027Person) SetPhoneNumberFax(v string)`

SetPhoneNumberFax sets PhoneNumberFax field to given value.

### HasPhoneNumberFax

`func (o *InlineResponse20027Person) HasPhoneNumberFax() bool`

HasPhoneNumberFax returns a boolean if a field has been set.

### GetPhoneNumberHome

`func (o *InlineResponse20027Person) GetPhoneNumberHome() string`

GetPhoneNumberHome returns the PhoneNumberHome field if non-nil, zero value otherwise.

### GetPhoneNumberHomeOk

`func (o *InlineResponse20027Person) GetPhoneNumberHomeOk() (*string, bool)`

GetPhoneNumberHomeOk returns a tuple with the PhoneNumberHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberHome

`func (o *InlineResponse20027Person) SetPhoneNumberHome(v string)`

SetPhoneNumberHome sets PhoneNumberHome field to given value.

### HasPhoneNumberHome

`func (o *InlineResponse20027Person) HasPhoneNumberHome() bool`

HasPhoneNumberHome returns a boolean if a field has been set.

### GetPhoneNumberMobile

`func (o *InlineResponse20027Person) GetPhoneNumberMobile() string`

GetPhoneNumberMobile returns the PhoneNumberMobile field if non-nil, zero value otherwise.

### GetPhoneNumberMobileOk

`func (o *InlineResponse20027Person) GetPhoneNumberMobileOk() (*string, bool)`

GetPhoneNumberMobileOk returns a tuple with the PhoneNumberMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberMobile

`func (o *InlineResponse20027Person) SetPhoneNumberMobile(v string)`

SetPhoneNumberMobile sets PhoneNumberMobile field to given value.

### HasPhoneNumberMobile

`func (o *InlineResponse20027Person) HasPhoneNumberMobile() bool`

HasPhoneNumberMobile returns a boolean if a field has been set.

### GetPhoneNumberOffice

`func (o *InlineResponse20027Person) GetPhoneNumberOffice() string`

GetPhoneNumberOffice returns the PhoneNumberOffice field if non-nil, zero value otherwise.

### GetPhoneNumberOfficeOk

`func (o *InlineResponse20027Person) GetPhoneNumberOfficeOk() (*string, bool)`

GetPhoneNumberOfficeOk returns a tuple with the PhoneNumberOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberOffice

`func (o *InlineResponse20027Person) SetPhoneNumberOffice(v string)`

SetPhoneNumberOffice sets PhoneNumberOffice field to given value.

### HasPhoneNumberOffice

`func (o *InlineResponse20027Person) HasPhoneNumberOffice() bool`

HasPhoneNumberOffice returns a boolean if a field has been set.

### GetPhoneNumberOfficeExt

`func (o *InlineResponse20027Person) GetPhoneNumberOfficeExt() string`

GetPhoneNumberOfficeExt returns the PhoneNumberOfficeExt field if non-nil, zero value otherwise.

### GetPhoneNumberOfficeExtOk

`func (o *InlineResponse20027Person) GetPhoneNumberOfficeExtOk() (*string, bool)`

GetPhoneNumberOfficeExtOk returns a tuple with the PhoneNumberOfficeExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberOfficeExt

`func (o *InlineResponse20027Person) SetPhoneNumberOfficeExt(v string)`

SetPhoneNumberOfficeExt sets PhoneNumberOfficeExt field to given value.

### HasPhoneNumberOfficeExt

`func (o *InlineResponse20027Person) HasPhoneNumberOfficeExt() bool`

HasPhoneNumberOfficeExt returns a boolean if a field has been set.

### GetPid

`func (o *InlineResponse20027Person) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *InlineResponse20027Person) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *InlineResponse20027Person) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *InlineResponse20027Person) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetPrivateNotes

`func (o *InlineResponse20027Person) GetPrivateNotes() string`

GetPrivateNotes returns the PrivateNotes field if non-nil, zero value otherwise.

### GetPrivateNotesOk

`func (o *InlineResponse20027Person) GetPrivateNotesOk() (*string, bool)`

GetPrivateNotesOk returns a tuple with the PrivateNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNotes

`func (o *InlineResponse20027Person) SetPrivateNotes(v string)`

SetPrivateNotes sets PrivateNotes field to given value.

### HasPrivateNotes

`func (o *InlineResponse20027Person) HasPrivateNotes() bool`

HasPrivateNotes returns a boolean if a field has been set.

### GetProfile

`func (o *InlineResponse20027Person) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *InlineResponse20027Person) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *InlineResponse20027Person) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *InlineResponse20027Person) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetSiteOwner

`func (o *InlineResponse20027Person) GetSiteOwner() bool`

GetSiteOwner returns the SiteOwner field if non-nil, zero value otherwise.

### GetSiteOwnerOk

`func (o *InlineResponse20027Person) GetSiteOwnerOk() (*bool, bool)`

GetSiteOwnerOk returns a tuple with the SiteOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteOwner

`func (o *InlineResponse20027Person) SetSiteOwner(v bool)`

SetSiteOwner sets SiteOwner field to given value.

### HasSiteOwner

`func (o *InlineResponse20027Person) HasSiteOwner() bool`

HasSiteOwner returns a boolean if a field has been set.

### GetTitle

`func (o *InlineResponse20027Person) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InlineResponse20027Person) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InlineResponse20027Person) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InlineResponse20027Person) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTwitter

`func (o *InlineResponse20027Person) GetTwitter() string`

GetTwitter returns the Twitter field if non-nil, zero value otherwise.

### GetTwitterOk

`func (o *InlineResponse20027Person) GetTwitterOk() (*string, bool)`

GetTwitterOk returns a tuple with the Twitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitter

`func (o *InlineResponse20027Person) SetTwitter(v string)`

SetTwitter sets Twitter field to given value.

### HasTwitter

`func (o *InlineResponse20027Person) HasTwitter() bool`

HasTwitter returns a boolean if a field has been set.

### GetUserName

`func (o *InlineResponse20027Person) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *InlineResponse20027Person) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *InlineResponse20027Person) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *InlineResponse20027Person) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetUserType

`func (o *InlineResponse20027Person) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *InlineResponse20027Person) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *InlineResponse20027Person) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *InlineResponse20027Person) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetUserUUID

`func (o *InlineResponse20027Person) GetUserUUID() string`

GetUserUUID returns the UserUUID field if non-nil, zero value otherwise.

### GetUserUUIDOk

`func (o *InlineResponse20027Person) GetUserUUIDOk() (*string, bool)`

GetUserUUIDOk returns a tuple with the UserUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUUID

`func (o *InlineResponse20027Person) SetUserUUID(v string)`

SetUserUUID sets UserUUID field to given value.

### HasUserUUID

`func (o *InlineResponse20027Person) HasUserUUID() bool`

HasUserUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


