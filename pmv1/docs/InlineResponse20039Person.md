# InlineResponse20039Person

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
**DocumentEditorInstalled** | Pointer to **bool** |  | [optional] 
**EmailAddress** | Pointer to **string** |  | [optional] 
**EmailAlt1** | Pointer to **string** |  | [optional] 
**EmailAlt2** | Pointer to **string** |  | [optional] 
**EmailAlt3** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**HasAccessToNewProjects** | Pointer to **bool** |  | [optional] 
**HasDeskAccount** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ImHandle** | Pointer to **string** |  | [optional] 
**ImService** | Pointer to **string** |  | [optional] 
**InOwnerCompany** | Pointer to **string** |  | [optional] 
**Integrations** | Pointer to [**InlineResponse20039PersonIntegrations**](InlineResponse20039PersonIntegrations.md) |  | [optional] 
**LastChangedOn** | Pointer to **string** |  | [optional] 
**LastLogin** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**LengthOfDay** | Pointer to **string** |  | [optional] 
**LoginCount** | Pointer to **string** |  | [optional] 
**MarketoId** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**OpenId** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to [**InlineResponse20039PersonPermissions**](InlineResponse20039PersonPermissions.md) |  | [optional] 
**PhoneNumberFax** | Pointer to **string** |  | [optional] 
**PhoneNumberHome** | Pointer to **string** |  | [optional] 
**PhoneNumberMobile** | Pointer to **string** |  | [optional] 
**PhoneNumberMobileParts** | Pointer to [**InlineResponse20013PhoneNumberMobileParts**](InlineResponse20013PhoneNumberMobileParts.md) |  | [optional] 
**PhoneNumberOffice** | Pointer to **string** |  | [optional] 
**PhoneNumberOfficeExt** | Pointer to **string** |  | [optional] 
**Pid** | Pointer to **string** |  | [optional] 
**PrivateNotes** | Pointer to **string** |  | [optional] 
**PrivateNotesText** | Pointer to **string** |  | [optional] 
**Profile** | Pointer to **string** |  | [optional] 
**ProfileText** | Pointer to **string** |  | [optional] 
**SiteOwner** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Twitter** | Pointer to **string** |  | [optional] 
**TwoFactorAuthEnabled** | Pointer to **bool** |  | [optional] 
**UserInvited** | Pointer to **string** |  | [optional] 
**UserInvitedDate** | Pointer to **string** |  | [optional] 
**UserInvitedStatus** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**UserType** | Pointer to **string** |  | [optional] 
**UserUUID** | Pointer to **string** |  | [optional] 

## Methods

### NewInlineResponse20039Person

`func NewInlineResponse20039Person() *InlineResponse20039Person`

NewInlineResponse20039Person instantiates a new InlineResponse20039Person object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20039PersonWithDefaults

`func NewInlineResponse20039PersonWithDefaults() *InlineResponse20039Person`

NewInlineResponse20039PersonWithDefaults instantiates a new InlineResponse20039Person object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *InlineResponse20039Person) GetAddress() InlineResponse20013Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InlineResponse20039Person) GetAddressOk() (*InlineResponse20013Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InlineResponse20039Person) SetAddress(v InlineResponse20013Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *InlineResponse20039Person) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAdministrator

`func (o *InlineResponse20039Person) GetAdministrator() bool`

GetAdministrator returns the Administrator field if non-nil, zero value otherwise.

### GetAdministratorOk

`func (o *InlineResponse20039Person) GetAdministratorOk() (*bool, bool)`

GetAdministratorOk returns a tuple with the Administrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrator

`func (o *InlineResponse20039Person) SetAdministrator(v bool)`

SetAdministrator sets Administrator field to given value.

### HasAdministrator

`func (o *InlineResponse20039Person) HasAdministrator() bool`

HasAdministrator returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *InlineResponse20039Person) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *InlineResponse20039Person) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *InlineResponse20039Person) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *InlineResponse20039Person) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetCompanyId

`func (o *InlineResponse20039Person) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *InlineResponse20039Person) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *InlineResponse20039Person) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *InlineResponse20039Person) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCompanyName

`func (o *InlineResponse20039Person) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *InlineResponse20039Person) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *InlineResponse20039Person) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *InlineResponse20039Person) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse20039Person) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse20039Person) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse20039Person) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse20039Person) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeleted

`func (o *InlineResponse20039Person) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *InlineResponse20039Person) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *InlineResponse20039Person) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *InlineResponse20039Person) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDocumentEditorInstalled

`func (o *InlineResponse20039Person) GetDocumentEditorInstalled() bool`

GetDocumentEditorInstalled returns the DocumentEditorInstalled field if non-nil, zero value otherwise.

### GetDocumentEditorInstalledOk

`func (o *InlineResponse20039Person) GetDocumentEditorInstalledOk() (*bool, bool)`

GetDocumentEditorInstalledOk returns a tuple with the DocumentEditorInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentEditorInstalled

`func (o *InlineResponse20039Person) SetDocumentEditorInstalled(v bool)`

SetDocumentEditorInstalled sets DocumentEditorInstalled field to given value.

### HasDocumentEditorInstalled

`func (o *InlineResponse20039Person) HasDocumentEditorInstalled() bool`

HasDocumentEditorInstalled returns a boolean if a field has been set.

### GetEmailAddress

`func (o *InlineResponse20039Person) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *InlineResponse20039Person) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *InlineResponse20039Person) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *InlineResponse20039Person) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetEmailAlt1

`func (o *InlineResponse20039Person) GetEmailAlt1() string`

GetEmailAlt1 returns the EmailAlt1 field if non-nil, zero value otherwise.

### GetEmailAlt1Ok

`func (o *InlineResponse20039Person) GetEmailAlt1Ok() (*string, bool)`

GetEmailAlt1Ok returns a tuple with the EmailAlt1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAlt1

`func (o *InlineResponse20039Person) SetEmailAlt1(v string)`

SetEmailAlt1 sets EmailAlt1 field to given value.

### HasEmailAlt1

`func (o *InlineResponse20039Person) HasEmailAlt1() bool`

HasEmailAlt1 returns a boolean if a field has been set.

### GetEmailAlt2

`func (o *InlineResponse20039Person) GetEmailAlt2() string`

GetEmailAlt2 returns the EmailAlt2 field if non-nil, zero value otherwise.

### GetEmailAlt2Ok

`func (o *InlineResponse20039Person) GetEmailAlt2Ok() (*string, bool)`

GetEmailAlt2Ok returns a tuple with the EmailAlt2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAlt2

`func (o *InlineResponse20039Person) SetEmailAlt2(v string)`

SetEmailAlt2 sets EmailAlt2 field to given value.

### HasEmailAlt2

`func (o *InlineResponse20039Person) HasEmailAlt2() bool`

HasEmailAlt2 returns a boolean if a field has been set.

### GetEmailAlt3

`func (o *InlineResponse20039Person) GetEmailAlt3() string`

GetEmailAlt3 returns the EmailAlt3 field if non-nil, zero value otherwise.

### GetEmailAlt3Ok

`func (o *InlineResponse20039Person) GetEmailAlt3Ok() (*string, bool)`

GetEmailAlt3Ok returns a tuple with the EmailAlt3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAlt3

`func (o *InlineResponse20039Person) SetEmailAlt3(v string)`

SetEmailAlt3 sets EmailAlt3 field to given value.

### HasEmailAlt3

`func (o *InlineResponse20039Person) HasEmailAlt3() bool`

HasEmailAlt3 returns a boolean if a field has been set.

### GetFirstName

`func (o *InlineResponse20039Person) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *InlineResponse20039Person) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *InlineResponse20039Person) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *InlineResponse20039Person) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetHasAccessToNewProjects

`func (o *InlineResponse20039Person) GetHasAccessToNewProjects() bool`

GetHasAccessToNewProjects returns the HasAccessToNewProjects field if non-nil, zero value otherwise.

### GetHasAccessToNewProjectsOk

`func (o *InlineResponse20039Person) GetHasAccessToNewProjectsOk() (*bool, bool)`

GetHasAccessToNewProjectsOk returns a tuple with the HasAccessToNewProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccessToNewProjects

`func (o *InlineResponse20039Person) SetHasAccessToNewProjects(v bool)`

SetHasAccessToNewProjects sets HasAccessToNewProjects field to given value.

### HasHasAccessToNewProjects

`func (o *InlineResponse20039Person) HasHasAccessToNewProjects() bool`

HasHasAccessToNewProjects returns a boolean if a field has been set.

### GetHasDeskAccount

`func (o *InlineResponse20039Person) GetHasDeskAccount() bool`

GetHasDeskAccount returns the HasDeskAccount field if non-nil, zero value otherwise.

### GetHasDeskAccountOk

`func (o *InlineResponse20039Person) GetHasDeskAccountOk() (*bool, bool)`

GetHasDeskAccountOk returns a tuple with the HasDeskAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDeskAccount

`func (o *InlineResponse20039Person) SetHasDeskAccount(v bool)`

SetHasDeskAccount sets HasDeskAccount field to given value.

### HasHasDeskAccount

`func (o *InlineResponse20039Person) HasHasDeskAccount() bool`

HasHasDeskAccount returns a boolean if a field has been set.

### GetId

`func (o *InlineResponse20039Person) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20039Person) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20039Person) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20039Person) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImHandle

`func (o *InlineResponse20039Person) GetImHandle() string`

GetImHandle returns the ImHandle field if non-nil, zero value otherwise.

### GetImHandleOk

`func (o *InlineResponse20039Person) GetImHandleOk() (*string, bool)`

GetImHandleOk returns a tuple with the ImHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImHandle

`func (o *InlineResponse20039Person) SetImHandle(v string)`

SetImHandle sets ImHandle field to given value.

### HasImHandle

`func (o *InlineResponse20039Person) HasImHandle() bool`

HasImHandle returns a boolean if a field has been set.

### GetImService

`func (o *InlineResponse20039Person) GetImService() string`

GetImService returns the ImService field if non-nil, zero value otherwise.

### GetImServiceOk

`func (o *InlineResponse20039Person) GetImServiceOk() (*string, bool)`

GetImServiceOk returns a tuple with the ImService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImService

`func (o *InlineResponse20039Person) SetImService(v string)`

SetImService sets ImService field to given value.

### HasImService

`func (o *InlineResponse20039Person) HasImService() bool`

HasImService returns a boolean if a field has been set.

### GetInOwnerCompany

`func (o *InlineResponse20039Person) GetInOwnerCompany() string`

GetInOwnerCompany returns the InOwnerCompany field if non-nil, zero value otherwise.

### GetInOwnerCompanyOk

`func (o *InlineResponse20039Person) GetInOwnerCompanyOk() (*string, bool)`

GetInOwnerCompanyOk returns a tuple with the InOwnerCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInOwnerCompany

`func (o *InlineResponse20039Person) SetInOwnerCompany(v string)`

SetInOwnerCompany sets InOwnerCompany field to given value.

### HasInOwnerCompany

`func (o *InlineResponse20039Person) HasInOwnerCompany() bool`

HasInOwnerCompany returns a boolean if a field has been set.

### GetIntegrations

`func (o *InlineResponse20039Person) GetIntegrations() InlineResponse20039PersonIntegrations`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *InlineResponse20039Person) GetIntegrationsOk() (*InlineResponse20039PersonIntegrations, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *InlineResponse20039Person) SetIntegrations(v InlineResponse20039PersonIntegrations)`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *InlineResponse20039Person) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.

### GetLastChangedOn

`func (o *InlineResponse20039Person) GetLastChangedOn() string`

GetLastChangedOn returns the LastChangedOn field if non-nil, zero value otherwise.

### GetLastChangedOnOk

`func (o *InlineResponse20039Person) GetLastChangedOnOk() (*string, bool)`

GetLastChangedOnOk returns a tuple with the LastChangedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangedOn

`func (o *InlineResponse20039Person) SetLastChangedOn(v string)`

SetLastChangedOn sets LastChangedOn field to given value.

### HasLastChangedOn

`func (o *InlineResponse20039Person) HasLastChangedOn() bool`

HasLastChangedOn returns a boolean if a field has been set.

### GetLastLogin

`func (o *InlineResponse20039Person) GetLastLogin() string`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *InlineResponse20039Person) GetLastLoginOk() (*string, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *InlineResponse20039Person) SetLastLogin(v string)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *InlineResponse20039Person) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetLastName

`func (o *InlineResponse20039Person) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *InlineResponse20039Person) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *InlineResponse20039Person) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *InlineResponse20039Person) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLengthOfDay

`func (o *InlineResponse20039Person) GetLengthOfDay() string`

GetLengthOfDay returns the LengthOfDay field if non-nil, zero value otherwise.

### GetLengthOfDayOk

`func (o *InlineResponse20039Person) GetLengthOfDayOk() (*string, bool)`

GetLengthOfDayOk returns a tuple with the LengthOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthOfDay

`func (o *InlineResponse20039Person) SetLengthOfDay(v string)`

SetLengthOfDay sets LengthOfDay field to given value.

### HasLengthOfDay

`func (o *InlineResponse20039Person) HasLengthOfDay() bool`

HasLengthOfDay returns a boolean if a field has been set.

### GetLoginCount

`func (o *InlineResponse20039Person) GetLoginCount() string`

GetLoginCount returns the LoginCount field if non-nil, zero value otherwise.

### GetLoginCountOk

`func (o *InlineResponse20039Person) GetLoginCountOk() (*string, bool)`

GetLoginCountOk returns a tuple with the LoginCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCount

`func (o *InlineResponse20039Person) SetLoginCount(v string)`

SetLoginCount sets LoginCount field to given value.

### HasLoginCount

`func (o *InlineResponse20039Person) HasLoginCount() bool`

HasLoginCount returns a boolean if a field has been set.

### GetMarketoId

`func (o *InlineResponse20039Person) GetMarketoId() string`

GetMarketoId returns the MarketoId field if non-nil, zero value otherwise.

### GetMarketoIdOk

`func (o *InlineResponse20039Person) GetMarketoIdOk() (*string, bool)`

GetMarketoIdOk returns a tuple with the MarketoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketoId

`func (o *InlineResponse20039Person) SetMarketoId(v string)`

SetMarketoId sets MarketoId field to given value.

### HasMarketoId

`func (o *InlineResponse20039Person) HasMarketoId() bool`

HasMarketoId returns a boolean if a field has been set.

### GetNotes

`func (o *InlineResponse20039Person) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InlineResponse20039Person) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InlineResponse20039Person) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InlineResponse20039Person) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetOpenId

`func (o *InlineResponse20039Person) GetOpenId() string`

GetOpenId returns the OpenId field if non-nil, zero value otherwise.

### GetOpenIdOk

`func (o *InlineResponse20039Person) GetOpenIdOk() (*string, bool)`

GetOpenIdOk returns a tuple with the OpenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenId

`func (o *InlineResponse20039Person) SetOpenId(v string)`

SetOpenId sets OpenId field to given value.

### HasOpenId

`func (o *InlineResponse20039Person) HasOpenId() bool`

HasOpenId returns a boolean if a field has been set.

### GetPermissions

`func (o *InlineResponse20039Person) GetPermissions() InlineResponse20039PersonPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *InlineResponse20039Person) GetPermissionsOk() (*InlineResponse20039PersonPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *InlineResponse20039Person) SetPermissions(v InlineResponse20039PersonPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *InlineResponse20039Person) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetPhoneNumberFax

`func (o *InlineResponse20039Person) GetPhoneNumberFax() string`

GetPhoneNumberFax returns the PhoneNumberFax field if non-nil, zero value otherwise.

### GetPhoneNumberFaxOk

`func (o *InlineResponse20039Person) GetPhoneNumberFaxOk() (*string, bool)`

GetPhoneNumberFaxOk returns a tuple with the PhoneNumberFax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberFax

`func (o *InlineResponse20039Person) SetPhoneNumberFax(v string)`

SetPhoneNumberFax sets PhoneNumberFax field to given value.

### HasPhoneNumberFax

`func (o *InlineResponse20039Person) HasPhoneNumberFax() bool`

HasPhoneNumberFax returns a boolean if a field has been set.

### GetPhoneNumberHome

`func (o *InlineResponse20039Person) GetPhoneNumberHome() string`

GetPhoneNumberHome returns the PhoneNumberHome field if non-nil, zero value otherwise.

### GetPhoneNumberHomeOk

`func (o *InlineResponse20039Person) GetPhoneNumberHomeOk() (*string, bool)`

GetPhoneNumberHomeOk returns a tuple with the PhoneNumberHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberHome

`func (o *InlineResponse20039Person) SetPhoneNumberHome(v string)`

SetPhoneNumberHome sets PhoneNumberHome field to given value.

### HasPhoneNumberHome

`func (o *InlineResponse20039Person) HasPhoneNumberHome() bool`

HasPhoneNumberHome returns a boolean if a field has been set.

### GetPhoneNumberMobile

`func (o *InlineResponse20039Person) GetPhoneNumberMobile() string`

GetPhoneNumberMobile returns the PhoneNumberMobile field if non-nil, zero value otherwise.

### GetPhoneNumberMobileOk

`func (o *InlineResponse20039Person) GetPhoneNumberMobileOk() (*string, bool)`

GetPhoneNumberMobileOk returns a tuple with the PhoneNumberMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberMobile

`func (o *InlineResponse20039Person) SetPhoneNumberMobile(v string)`

SetPhoneNumberMobile sets PhoneNumberMobile field to given value.

### HasPhoneNumberMobile

`func (o *InlineResponse20039Person) HasPhoneNumberMobile() bool`

HasPhoneNumberMobile returns a boolean if a field has been set.

### GetPhoneNumberMobileParts

`func (o *InlineResponse20039Person) GetPhoneNumberMobileParts() InlineResponse20013PhoneNumberMobileParts`

GetPhoneNumberMobileParts returns the PhoneNumberMobileParts field if non-nil, zero value otherwise.

### GetPhoneNumberMobilePartsOk

`func (o *InlineResponse20039Person) GetPhoneNumberMobilePartsOk() (*InlineResponse20013PhoneNumberMobileParts, bool)`

GetPhoneNumberMobilePartsOk returns a tuple with the PhoneNumberMobileParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberMobileParts

`func (o *InlineResponse20039Person) SetPhoneNumberMobileParts(v InlineResponse20013PhoneNumberMobileParts)`

SetPhoneNumberMobileParts sets PhoneNumberMobileParts field to given value.

### HasPhoneNumberMobileParts

`func (o *InlineResponse20039Person) HasPhoneNumberMobileParts() bool`

HasPhoneNumberMobileParts returns a boolean if a field has been set.

### GetPhoneNumberOffice

`func (o *InlineResponse20039Person) GetPhoneNumberOffice() string`

GetPhoneNumberOffice returns the PhoneNumberOffice field if non-nil, zero value otherwise.

### GetPhoneNumberOfficeOk

`func (o *InlineResponse20039Person) GetPhoneNumberOfficeOk() (*string, bool)`

GetPhoneNumberOfficeOk returns a tuple with the PhoneNumberOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberOffice

`func (o *InlineResponse20039Person) SetPhoneNumberOffice(v string)`

SetPhoneNumberOffice sets PhoneNumberOffice field to given value.

### HasPhoneNumberOffice

`func (o *InlineResponse20039Person) HasPhoneNumberOffice() bool`

HasPhoneNumberOffice returns a boolean if a field has been set.

### GetPhoneNumberOfficeExt

`func (o *InlineResponse20039Person) GetPhoneNumberOfficeExt() string`

GetPhoneNumberOfficeExt returns the PhoneNumberOfficeExt field if non-nil, zero value otherwise.

### GetPhoneNumberOfficeExtOk

`func (o *InlineResponse20039Person) GetPhoneNumberOfficeExtOk() (*string, bool)`

GetPhoneNumberOfficeExtOk returns a tuple with the PhoneNumberOfficeExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberOfficeExt

`func (o *InlineResponse20039Person) SetPhoneNumberOfficeExt(v string)`

SetPhoneNumberOfficeExt sets PhoneNumberOfficeExt field to given value.

### HasPhoneNumberOfficeExt

`func (o *InlineResponse20039Person) HasPhoneNumberOfficeExt() bool`

HasPhoneNumberOfficeExt returns a boolean if a field has been set.

### GetPid

`func (o *InlineResponse20039Person) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *InlineResponse20039Person) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *InlineResponse20039Person) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *InlineResponse20039Person) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetPrivateNotes

`func (o *InlineResponse20039Person) GetPrivateNotes() string`

GetPrivateNotes returns the PrivateNotes field if non-nil, zero value otherwise.

### GetPrivateNotesOk

`func (o *InlineResponse20039Person) GetPrivateNotesOk() (*string, bool)`

GetPrivateNotesOk returns a tuple with the PrivateNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNotes

`func (o *InlineResponse20039Person) SetPrivateNotes(v string)`

SetPrivateNotes sets PrivateNotes field to given value.

### HasPrivateNotes

`func (o *InlineResponse20039Person) HasPrivateNotes() bool`

HasPrivateNotes returns a boolean if a field has been set.

### GetPrivateNotesText

`func (o *InlineResponse20039Person) GetPrivateNotesText() string`

GetPrivateNotesText returns the PrivateNotesText field if non-nil, zero value otherwise.

### GetPrivateNotesTextOk

`func (o *InlineResponse20039Person) GetPrivateNotesTextOk() (*string, bool)`

GetPrivateNotesTextOk returns a tuple with the PrivateNotesText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNotesText

`func (o *InlineResponse20039Person) SetPrivateNotesText(v string)`

SetPrivateNotesText sets PrivateNotesText field to given value.

### HasPrivateNotesText

`func (o *InlineResponse20039Person) HasPrivateNotesText() bool`

HasPrivateNotesText returns a boolean if a field has been set.

### GetProfile

`func (o *InlineResponse20039Person) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *InlineResponse20039Person) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *InlineResponse20039Person) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *InlineResponse20039Person) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetProfileText

`func (o *InlineResponse20039Person) GetProfileText() string`

GetProfileText returns the ProfileText field if non-nil, zero value otherwise.

### GetProfileTextOk

`func (o *InlineResponse20039Person) GetProfileTextOk() (*string, bool)`

GetProfileTextOk returns a tuple with the ProfileText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileText

`func (o *InlineResponse20039Person) SetProfileText(v string)`

SetProfileText sets ProfileText field to given value.

### HasProfileText

`func (o *InlineResponse20039Person) HasProfileText() bool`

HasProfileText returns a boolean if a field has been set.

### GetSiteOwner

`func (o *InlineResponse20039Person) GetSiteOwner() bool`

GetSiteOwner returns the SiteOwner field if non-nil, zero value otherwise.

### GetSiteOwnerOk

`func (o *InlineResponse20039Person) GetSiteOwnerOk() (*bool, bool)`

GetSiteOwnerOk returns a tuple with the SiteOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteOwner

`func (o *InlineResponse20039Person) SetSiteOwner(v bool)`

SetSiteOwner sets SiteOwner field to given value.

### HasSiteOwner

`func (o *InlineResponse20039Person) HasSiteOwner() bool`

HasSiteOwner returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse20039Person) GetTags() []map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse20039Person) GetTagsOk() (*[]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse20039Person) SetTags(v []map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse20039Person) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTitle

`func (o *InlineResponse20039Person) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InlineResponse20039Person) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InlineResponse20039Person) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InlineResponse20039Person) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTwitter

`func (o *InlineResponse20039Person) GetTwitter() string`

GetTwitter returns the Twitter field if non-nil, zero value otherwise.

### GetTwitterOk

`func (o *InlineResponse20039Person) GetTwitterOk() (*string, bool)`

GetTwitterOk returns a tuple with the Twitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitter

`func (o *InlineResponse20039Person) SetTwitter(v string)`

SetTwitter sets Twitter field to given value.

### HasTwitter

`func (o *InlineResponse20039Person) HasTwitter() bool`

HasTwitter returns a boolean if a field has been set.

### GetTwoFactorAuthEnabled

`func (o *InlineResponse20039Person) GetTwoFactorAuthEnabled() bool`

GetTwoFactorAuthEnabled returns the TwoFactorAuthEnabled field if non-nil, zero value otherwise.

### GetTwoFactorAuthEnabledOk

`func (o *InlineResponse20039Person) GetTwoFactorAuthEnabledOk() (*bool, bool)`

GetTwoFactorAuthEnabledOk returns a tuple with the TwoFactorAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorAuthEnabled

`func (o *InlineResponse20039Person) SetTwoFactorAuthEnabled(v bool)`

SetTwoFactorAuthEnabled sets TwoFactorAuthEnabled field to given value.

### HasTwoFactorAuthEnabled

`func (o *InlineResponse20039Person) HasTwoFactorAuthEnabled() bool`

HasTwoFactorAuthEnabled returns a boolean if a field has been set.

### GetUserInvited

`func (o *InlineResponse20039Person) GetUserInvited() string`

GetUserInvited returns the UserInvited field if non-nil, zero value otherwise.

### GetUserInvitedOk

`func (o *InlineResponse20039Person) GetUserInvitedOk() (*string, bool)`

GetUserInvitedOk returns a tuple with the UserInvited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInvited

`func (o *InlineResponse20039Person) SetUserInvited(v string)`

SetUserInvited sets UserInvited field to given value.

### HasUserInvited

`func (o *InlineResponse20039Person) HasUserInvited() bool`

HasUserInvited returns a boolean if a field has been set.

### GetUserInvitedDate

`func (o *InlineResponse20039Person) GetUserInvitedDate() string`

GetUserInvitedDate returns the UserInvitedDate field if non-nil, zero value otherwise.

### GetUserInvitedDateOk

`func (o *InlineResponse20039Person) GetUserInvitedDateOk() (*string, bool)`

GetUserInvitedDateOk returns a tuple with the UserInvitedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInvitedDate

`func (o *InlineResponse20039Person) SetUserInvitedDate(v string)`

SetUserInvitedDate sets UserInvitedDate field to given value.

### HasUserInvitedDate

`func (o *InlineResponse20039Person) HasUserInvitedDate() bool`

HasUserInvitedDate returns a boolean if a field has been set.

### GetUserInvitedStatus

`func (o *InlineResponse20039Person) GetUserInvitedStatus() string`

GetUserInvitedStatus returns the UserInvitedStatus field if non-nil, zero value otherwise.

### GetUserInvitedStatusOk

`func (o *InlineResponse20039Person) GetUserInvitedStatusOk() (*string, bool)`

GetUserInvitedStatusOk returns a tuple with the UserInvitedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInvitedStatus

`func (o *InlineResponse20039Person) SetUserInvitedStatus(v string)`

SetUserInvitedStatus sets UserInvitedStatus field to given value.

### HasUserInvitedStatus

`func (o *InlineResponse20039Person) HasUserInvitedStatus() bool`

HasUserInvitedStatus returns a boolean if a field has been set.

### GetUserName

`func (o *InlineResponse20039Person) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *InlineResponse20039Person) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *InlineResponse20039Person) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *InlineResponse20039Person) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetUserType

`func (o *InlineResponse20039Person) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *InlineResponse20039Person) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *InlineResponse20039Person) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *InlineResponse20039Person) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetUserUUID

`func (o *InlineResponse20039Person) GetUserUUID() string`

GetUserUUID returns the UserUUID field if non-nil, zero value otherwise.

### GetUserUUIDOk

`func (o *InlineResponse20039Person) GetUserUUIDOk() (*string, bool)`

GetUserUUIDOk returns a tuple with the UserUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUUID

`func (o *InlineResponse20039Person) SetUserUUID(v string)`

SetUserUUID sets UserUUID field to given value.

### HasUserUUID

`func (o *InlineResponse20039Person) HasUserUUID() bool`

HasUserUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


