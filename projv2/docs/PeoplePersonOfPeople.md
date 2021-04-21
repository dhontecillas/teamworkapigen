# PeoplePersonOfPeople

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**APIEnabled** | Pointer to **bool** |  | [optional] 
**Address** | Pointer to [**PeopleAddress**](PeopleAddress.md) |  | [optional] 
**AddressCity** | Pointer to **string** |  | [optional] 
**AddressCountry** | Pointer to **string** |  | [optional] 
**AddressLine1** | Pointer to **string** |  | [optional] 
**AddressLine2** | Pointer to **string** |  | [optional] 
**AddressState** | Pointer to **string** |  | [optional] 
**AddressZip** | Pointer to **string** |  | [optional] 
**Administrator** | Pointer to **bool** |  | [optional] 
**AvatarUrl** | Pointer to **string** |  | [optional] 
**CanManageProjectTemplates** | Pointer to **bool** |  | [optional] 
**CanViewProjectTemplates** | Pointer to **bool** |  | [optional] 
**Company** | Pointer to [**PeopleCompany**](PeopleCompany.md) |  | [optional] 
**CompanyId** | Pointer to **int32** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**EmailAddress** | Pointer to **string** |  | [optional] 
**EmailAlt1** | Pointer to **string** |  | [optional] 
**EmailAlt2** | Pointer to **string** |  | [optional] 
**EmailAlt3** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**HasAccessToNewProjects** | Pointer to **bool** |  | [optional] 
**HubspotEnabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**ImHandle** | Pointer to **string** |  | [optional] 
**ImService** | Pointer to **string** |  | [optional] 
**InOwnerCompany** | Pointer to **bool** |  | [optional] 
**IsClientUser** | Pointer to **bool** |  | [optional] 
**IsClockedIn** | Pointer to **bool** |  | [optional] 
**LastActive** | Pointer to **string** |  | [optional] 
**LastChangedOn** | Pointer to **string** |  | [optional] 
**LastLogin** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**LengthOfDay** | Pointer to **float32** |  | [optional] 
**Localization** | Pointer to [**PeopleLocalization**](PeopleLocalization.md) |  | [optional] 
**LoginCount** | Pointer to **int32** |  | [optional] 
**MentionName** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Notifications** | Pointer to [**PeopleNotifications**](PeopleNotifications.md) |  | [optional] 
**OpenId** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to [**PeoplePermissions**](PeoplePermissions.md) |  | [optional] 
**PhoneNumberFax** | Pointer to **string** |  | [optional] 
**PhoneNumberHome** | Pointer to **string** |  | [optional] 
**PhoneNumberMobile** | Pointer to **string** |  | [optional] 
**PhoneNumberMobileParts** | Pointer to [**PeoplePhoneParts**](PeoplePhoneParts.md) |  | [optional] 
**PhoneNumberOffice** | Pointer to **string** |  | [optional] 
**PhoneNumberOfficeExt** | Pointer to **string** |  | [optional] 
**Pid** | Pointer to **string** |  | [optional] 
**PrivateNotes** | Pointer to **string** |  | [optional] 
**PrivateNotesText** | Pointer to **string** |  | [optional] 
**Profile** | Pointer to **string** |  | [optional] 
**ProfileText** | Pointer to **string** |  | [optional] 
**ProjectRoles** | Pointer to **string** |  | [optional] 
**Projects** | Pointer to **[]int32** |  | [optional] 
**SiteOwner** | Pointer to **bool** |  | [optional] 
**Social** | Pointer to [**PeopleSocial**](PeopleSocial.md) |  | [optional] 
**Tags** | Pointer to [**[]TagsTag**](TagsTag.md) |  | [optional] 
**Teams** | Pointer to **[]string** |  | [optional] 
**TextFormat** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Twitter** | Pointer to **string** |  | [optional] 
**TwoFactorAuthEnabled** | Pointer to **bool** |  | [optional] 
**UseShorthandDurations** | Pointer to **bool** |  | [optional] 
**UserInvited** | Pointer to **int32** |  | [optional] 
**UserInvitedDate** | Pointer to **string** |  | [optional] 
**UserInvitedStatus** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**UserType** | Pointer to **string** |  | [optional] 
**UserUUID** | Pointer to **string** |  | [optional] 
**WorkingHours** | Pointer to [**PeopleWorkingHours**](PeopleWorkingHours.md) |  | [optional] 

## Methods

### NewPeoplePersonOfPeople

`func NewPeoplePersonOfPeople() *PeoplePersonOfPeople`

NewPeoplePersonOfPeople instantiates a new PeoplePersonOfPeople object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeoplePersonOfPeopleWithDefaults

`func NewPeoplePersonOfPeopleWithDefaults() *PeoplePersonOfPeople`

NewPeoplePersonOfPeopleWithDefaults instantiates a new PeoplePersonOfPeople object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAPIEnabled

`func (o *PeoplePersonOfPeople) GetAPIEnabled() bool`

GetAPIEnabled returns the APIEnabled field if non-nil, zero value otherwise.

### GetAPIEnabledOk

`func (o *PeoplePersonOfPeople) GetAPIEnabledOk() (*bool, bool)`

GetAPIEnabledOk returns a tuple with the APIEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPIEnabled

`func (o *PeoplePersonOfPeople) SetAPIEnabled(v bool)`

SetAPIEnabled sets APIEnabled field to given value.

### HasAPIEnabled

`func (o *PeoplePersonOfPeople) HasAPIEnabled() bool`

HasAPIEnabled returns a boolean if a field has been set.

### GetAddress

`func (o *PeoplePersonOfPeople) GetAddress() PeopleAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PeoplePersonOfPeople) GetAddressOk() (*PeopleAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PeoplePersonOfPeople) SetAddress(v PeopleAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PeoplePersonOfPeople) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddressCity

`func (o *PeoplePersonOfPeople) GetAddressCity() string`

GetAddressCity returns the AddressCity field if non-nil, zero value otherwise.

### GetAddressCityOk

`func (o *PeoplePersonOfPeople) GetAddressCityOk() (*string, bool)`

GetAddressCityOk returns a tuple with the AddressCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCity

`func (o *PeoplePersonOfPeople) SetAddressCity(v string)`

SetAddressCity sets AddressCity field to given value.

### HasAddressCity

`func (o *PeoplePersonOfPeople) HasAddressCity() bool`

HasAddressCity returns a boolean if a field has been set.

### GetAddressCountry

`func (o *PeoplePersonOfPeople) GetAddressCountry() string`

GetAddressCountry returns the AddressCountry field if non-nil, zero value otherwise.

### GetAddressCountryOk

`func (o *PeoplePersonOfPeople) GetAddressCountryOk() (*string, bool)`

GetAddressCountryOk returns a tuple with the AddressCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCountry

`func (o *PeoplePersonOfPeople) SetAddressCountry(v string)`

SetAddressCountry sets AddressCountry field to given value.

### HasAddressCountry

`func (o *PeoplePersonOfPeople) HasAddressCountry() bool`

HasAddressCountry returns a boolean if a field has been set.

### GetAddressLine1

`func (o *PeoplePersonOfPeople) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *PeoplePersonOfPeople) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *PeoplePersonOfPeople) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *PeoplePersonOfPeople) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *PeoplePersonOfPeople) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *PeoplePersonOfPeople) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *PeoplePersonOfPeople) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *PeoplePersonOfPeople) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAddressState

`func (o *PeoplePersonOfPeople) GetAddressState() string`

GetAddressState returns the AddressState field if non-nil, zero value otherwise.

### GetAddressStateOk

`func (o *PeoplePersonOfPeople) GetAddressStateOk() (*string, bool)`

GetAddressStateOk returns a tuple with the AddressState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressState

`func (o *PeoplePersonOfPeople) SetAddressState(v string)`

SetAddressState sets AddressState field to given value.

### HasAddressState

`func (o *PeoplePersonOfPeople) HasAddressState() bool`

HasAddressState returns a boolean if a field has been set.

### GetAddressZip

`func (o *PeoplePersonOfPeople) GetAddressZip() string`

GetAddressZip returns the AddressZip field if non-nil, zero value otherwise.

### GetAddressZipOk

`func (o *PeoplePersonOfPeople) GetAddressZipOk() (*string, bool)`

GetAddressZipOk returns a tuple with the AddressZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressZip

`func (o *PeoplePersonOfPeople) SetAddressZip(v string)`

SetAddressZip sets AddressZip field to given value.

### HasAddressZip

`func (o *PeoplePersonOfPeople) HasAddressZip() bool`

HasAddressZip returns a boolean if a field has been set.

### GetAdministrator

`func (o *PeoplePersonOfPeople) GetAdministrator() bool`

GetAdministrator returns the Administrator field if non-nil, zero value otherwise.

### GetAdministratorOk

`func (o *PeoplePersonOfPeople) GetAdministratorOk() (*bool, bool)`

GetAdministratorOk returns a tuple with the Administrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrator

`func (o *PeoplePersonOfPeople) SetAdministrator(v bool)`

SetAdministrator sets Administrator field to given value.

### HasAdministrator

`func (o *PeoplePersonOfPeople) HasAdministrator() bool`

HasAdministrator returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *PeoplePersonOfPeople) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *PeoplePersonOfPeople) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *PeoplePersonOfPeople) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *PeoplePersonOfPeople) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetCanManageProjectTemplates

`func (o *PeoplePersonOfPeople) GetCanManageProjectTemplates() bool`

GetCanManageProjectTemplates returns the CanManageProjectTemplates field if non-nil, zero value otherwise.

### GetCanManageProjectTemplatesOk

`func (o *PeoplePersonOfPeople) GetCanManageProjectTemplatesOk() (*bool, bool)`

GetCanManageProjectTemplatesOk returns a tuple with the CanManageProjectTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageProjectTemplates

`func (o *PeoplePersonOfPeople) SetCanManageProjectTemplates(v bool)`

SetCanManageProjectTemplates sets CanManageProjectTemplates field to given value.

### HasCanManageProjectTemplates

`func (o *PeoplePersonOfPeople) HasCanManageProjectTemplates() bool`

HasCanManageProjectTemplates returns a boolean if a field has been set.

### GetCanViewProjectTemplates

`func (o *PeoplePersonOfPeople) GetCanViewProjectTemplates() bool`

GetCanViewProjectTemplates returns the CanViewProjectTemplates field if non-nil, zero value otherwise.

### GetCanViewProjectTemplatesOk

`func (o *PeoplePersonOfPeople) GetCanViewProjectTemplatesOk() (*bool, bool)`

GetCanViewProjectTemplatesOk returns a tuple with the CanViewProjectTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewProjectTemplates

`func (o *PeoplePersonOfPeople) SetCanViewProjectTemplates(v bool)`

SetCanViewProjectTemplates sets CanViewProjectTemplates field to given value.

### HasCanViewProjectTemplates

`func (o *PeoplePersonOfPeople) HasCanViewProjectTemplates() bool`

HasCanViewProjectTemplates returns a boolean if a field has been set.

### GetCompany

`func (o *PeoplePersonOfPeople) GetCompany() PeopleCompany`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *PeoplePersonOfPeople) GetCompanyOk() (*PeopleCompany, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *PeoplePersonOfPeople) SetCompany(v PeopleCompany)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *PeoplePersonOfPeople) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCompanyId

`func (o *PeoplePersonOfPeople) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *PeoplePersonOfPeople) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *PeoplePersonOfPeople) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *PeoplePersonOfPeople) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCompanyName

`func (o *PeoplePersonOfPeople) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *PeoplePersonOfPeople) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *PeoplePersonOfPeople) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *PeoplePersonOfPeople) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PeoplePersonOfPeople) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PeoplePersonOfPeople) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PeoplePersonOfPeople) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PeoplePersonOfPeople) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeleted

`func (o *PeoplePersonOfPeople) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *PeoplePersonOfPeople) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *PeoplePersonOfPeople) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *PeoplePersonOfPeople) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetEmailAddress

`func (o *PeoplePersonOfPeople) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *PeoplePersonOfPeople) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *PeoplePersonOfPeople) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *PeoplePersonOfPeople) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetEmailAlt1

`func (o *PeoplePersonOfPeople) GetEmailAlt1() string`

GetEmailAlt1 returns the EmailAlt1 field if non-nil, zero value otherwise.

### GetEmailAlt1Ok

`func (o *PeoplePersonOfPeople) GetEmailAlt1Ok() (*string, bool)`

GetEmailAlt1Ok returns a tuple with the EmailAlt1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAlt1

`func (o *PeoplePersonOfPeople) SetEmailAlt1(v string)`

SetEmailAlt1 sets EmailAlt1 field to given value.

### HasEmailAlt1

`func (o *PeoplePersonOfPeople) HasEmailAlt1() bool`

HasEmailAlt1 returns a boolean if a field has been set.

### GetEmailAlt2

`func (o *PeoplePersonOfPeople) GetEmailAlt2() string`

GetEmailAlt2 returns the EmailAlt2 field if non-nil, zero value otherwise.

### GetEmailAlt2Ok

`func (o *PeoplePersonOfPeople) GetEmailAlt2Ok() (*string, bool)`

GetEmailAlt2Ok returns a tuple with the EmailAlt2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAlt2

`func (o *PeoplePersonOfPeople) SetEmailAlt2(v string)`

SetEmailAlt2 sets EmailAlt2 field to given value.

### HasEmailAlt2

`func (o *PeoplePersonOfPeople) HasEmailAlt2() bool`

HasEmailAlt2 returns a boolean if a field has been set.

### GetEmailAlt3

`func (o *PeoplePersonOfPeople) GetEmailAlt3() string`

GetEmailAlt3 returns the EmailAlt3 field if non-nil, zero value otherwise.

### GetEmailAlt3Ok

`func (o *PeoplePersonOfPeople) GetEmailAlt3Ok() (*string, bool)`

GetEmailAlt3Ok returns a tuple with the EmailAlt3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAlt3

`func (o *PeoplePersonOfPeople) SetEmailAlt3(v string)`

SetEmailAlt3 sets EmailAlt3 field to given value.

### HasEmailAlt3

`func (o *PeoplePersonOfPeople) HasEmailAlt3() bool`

HasEmailAlt3 returns a boolean if a field has been set.

### GetFirstName

`func (o *PeoplePersonOfPeople) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PeoplePersonOfPeople) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PeoplePersonOfPeople) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PeoplePersonOfPeople) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetHasAccessToNewProjects

`func (o *PeoplePersonOfPeople) GetHasAccessToNewProjects() bool`

GetHasAccessToNewProjects returns the HasAccessToNewProjects field if non-nil, zero value otherwise.

### GetHasAccessToNewProjectsOk

`func (o *PeoplePersonOfPeople) GetHasAccessToNewProjectsOk() (*bool, bool)`

GetHasAccessToNewProjectsOk returns a tuple with the HasAccessToNewProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccessToNewProjects

`func (o *PeoplePersonOfPeople) SetHasAccessToNewProjects(v bool)`

SetHasAccessToNewProjects sets HasAccessToNewProjects field to given value.

### HasHasAccessToNewProjects

`func (o *PeoplePersonOfPeople) HasHasAccessToNewProjects() bool`

HasHasAccessToNewProjects returns a boolean if a field has been set.

### GetHubspotEnabled

`func (o *PeoplePersonOfPeople) GetHubspotEnabled() bool`

GetHubspotEnabled returns the HubspotEnabled field if non-nil, zero value otherwise.

### GetHubspotEnabledOk

`func (o *PeoplePersonOfPeople) GetHubspotEnabledOk() (*bool, bool)`

GetHubspotEnabledOk returns a tuple with the HubspotEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubspotEnabled

`func (o *PeoplePersonOfPeople) SetHubspotEnabled(v bool)`

SetHubspotEnabled sets HubspotEnabled field to given value.

### HasHubspotEnabled

`func (o *PeoplePersonOfPeople) HasHubspotEnabled() bool`

HasHubspotEnabled returns a boolean if a field has been set.

### GetId

`func (o *PeoplePersonOfPeople) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PeoplePersonOfPeople) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PeoplePersonOfPeople) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PeoplePersonOfPeople) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImHandle

`func (o *PeoplePersonOfPeople) GetImHandle() string`

GetImHandle returns the ImHandle field if non-nil, zero value otherwise.

### GetImHandleOk

`func (o *PeoplePersonOfPeople) GetImHandleOk() (*string, bool)`

GetImHandleOk returns a tuple with the ImHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImHandle

`func (o *PeoplePersonOfPeople) SetImHandle(v string)`

SetImHandle sets ImHandle field to given value.

### HasImHandle

`func (o *PeoplePersonOfPeople) HasImHandle() bool`

HasImHandle returns a boolean if a field has been set.

### GetImService

`func (o *PeoplePersonOfPeople) GetImService() string`

GetImService returns the ImService field if non-nil, zero value otherwise.

### GetImServiceOk

`func (o *PeoplePersonOfPeople) GetImServiceOk() (*string, bool)`

GetImServiceOk returns a tuple with the ImService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImService

`func (o *PeoplePersonOfPeople) SetImService(v string)`

SetImService sets ImService field to given value.

### HasImService

`func (o *PeoplePersonOfPeople) HasImService() bool`

HasImService returns a boolean if a field has been set.

### GetInOwnerCompany

`func (o *PeoplePersonOfPeople) GetInOwnerCompany() bool`

GetInOwnerCompany returns the InOwnerCompany field if non-nil, zero value otherwise.

### GetInOwnerCompanyOk

`func (o *PeoplePersonOfPeople) GetInOwnerCompanyOk() (*bool, bool)`

GetInOwnerCompanyOk returns a tuple with the InOwnerCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInOwnerCompany

`func (o *PeoplePersonOfPeople) SetInOwnerCompany(v bool)`

SetInOwnerCompany sets InOwnerCompany field to given value.

### HasInOwnerCompany

`func (o *PeoplePersonOfPeople) HasInOwnerCompany() bool`

HasInOwnerCompany returns a boolean if a field has been set.

### GetIsClientUser

`func (o *PeoplePersonOfPeople) GetIsClientUser() bool`

GetIsClientUser returns the IsClientUser field if non-nil, zero value otherwise.

### GetIsClientUserOk

`func (o *PeoplePersonOfPeople) GetIsClientUserOk() (*bool, bool)`

GetIsClientUserOk returns a tuple with the IsClientUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClientUser

`func (o *PeoplePersonOfPeople) SetIsClientUser(v bool)`

SetIsClientUser sets IsClientUser field to given value.

### HasIsClientUser

`func (o *PeoplePersonOfPeople) HasIsClientUser() bool`

HasIsClientUser returns a boolean if a field has been set.

### GetIsClockedIn

`func (o *PeoplePersonOfPeople) GetIsClockedIn() bool`

GetIsClockedIn returns the IsClockedIn field if non-nil, zero value otherwise.

### GetIsClockedInOk

`func (o *PeoplePersonOfPeople) GetIsClockedInOk() (*bool, bool)`

GetIsClockedInOk returns a tuple with the IsClockedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClockedIn

`func (o *PeoplePersonOfPeople) SetIsClockedIn(v bool)`

SetIsClockedIn sets IsClockedIn field to given value.

### HasIsClockedIn

`func (o *PeoplePersonOfPeople) HasIsClockedIn() bool`

HasIsClockedIn returns a boolean if a field has been set.

### GetLastActive

`func (o *PeoplePersonOfPeople) GetLastActive() string`

GetLastActive returns the LastActive field if non-nil, zero value otherwise.

### GetLastActiveOk

`func (o *PeoplePersonOfPeople) GetLastActiveOk() (*string, bool)`

GetLastActiveOk returns a tuple with the LastActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActive

`func (o *PeoplePersonOfPeople) SetLastActive(v string)`

SetLastActive sets LastActive field to given value.

### HasLastActive

`func (o *PeoplePersonOfPeople) HasLastActive() bool`

HasLastActive returns a boolean if a field has been set.

### GetLastChangedOn

`func (o *PeoplePersonOfPeople) GetLastChangedOn() string`

GetLastChangedOn returns the LastChangedOn field if non-nil, zero value otherwise.

### GetLastChangedOnOk

`func (o *PeoplePersonOfPeople) GetLastChangedOnOk() (*string, bool)`

GetLastChangedOnOk returns a tuple with the LastChangedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangedOn

`func (o *PeoplePersonOfPeople) SetLastChangedOn(v string)`

SetLastChangedOn sets LastChangedOn field to given value.

### HasLastChangedOn

`func (o *PeoplePersonOfPeople) HasLastChangedOn() bool`

HasLastChangedOn returns a boolean if a field has been set.

### GetLastLogin

`func (o *PeoplePersonOfPeople) GetLastLogin() string`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *PeoplePersonOfPeople) GetLastLoginOk() (*string, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *PeoplePersonOfPeople) SetLastLogin(v string)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *PeoplePersonOfPeople) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetLastName

`func (o *PeoplePersonOfPeople) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PeoplePersonOfPeople) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PeoplePersonOfPeople) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PeoplePersonOfPeople) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLengthOfDay

`func (o *PeoplePersonOfPeople) GetLengthOfDay() float32`

GetLengthOfDay returns the LengthOfDay field if non-nil, zero value otherwise.

### GetLengthOfDayOk

`func (o *PeoplePersonOfPeople) GetLengthOfDayOk() (*float32, bool)`

GetLengthOfDayOk returns a tuple with the LengthOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthOfDay

`func (o *PeoplePersonOfPeople) SetLengthOfDay(v float32)`

SetLengthOfDay sets LengthOfDay field to given value.

### HasLengthOfDay

`func (o *PeoplePersonOfPeople) HasLengthOfDay() bool`

HasLengthOfDay returns a boolean if a field has been set.

### GetLocalization

`func (o *PeoplePersonOfPeople) GetLocalization() PeopleLocalization`

GetLocalization returns the Localization field if non-nil, zero value otherwise.

### GetLocalizationOk

`func (o *PeoplePersonOfPeople) GetLocalizationOk() (*PeopleLocalization, bool)`

GetLocalizationOk returns a tuple with the Localization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalization

`func (o *PeoplePersonOfPeople) SetLocalization(v PeopleLocalization)`

SetLocalization sets Localization field to given value.

### HasLocalization

`func (o *PeoplePersonOfPeople) HasLocalization() bool`

HasLocalization returns a boolean if a field has been set.

### GetLoginCount

`func (o *PeoplePersonOfPeople) GetLoginCount() int32`

GetLoginCount returns the LoginCount field if non-nil, zero value otherwise.

### GetLoginCountOk

`func (o *PeoplePersonOfPeople) GetLoginCountOk() (*int32, bool)`

GetLoginCountOk returns a tuple with the LoginCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCount

`func (o *PeoplePersonOfPeople) SetLoginCount(v int32)`

SetLoginCount sets LoginCount field to given value.

### HasLoginCount

`func (o *PeoplePersonOfPeople) HasLoginCount() bool`

HasLoginCount returns a boolean if a field has been set.

### GetMentionName

`func (o *PeoplePersonOfPeople) GetMentionName() string`

GetMentionName returns the MentionName field if non-nil, zero value otherwise.

### GetMentionNameOk

`func (o *PeoplePersonOfPeople) GetMentionNameOk() (*string, bool)`

GetMentionNameOk returns a tuple with the MentionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionName

`func (o *PeoplePersonOfPeople) SetMentionName(v string)`

SetMentionName sets MentionName field to given value.

### HasMentionName

`func (o *PeoplePersonOfPeople) HasMentionName() bool`

HasMentionName returns a boolean if a field has been set.

### GetNotes

`func (o *PeoplePersonOfPeople) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *PeoplePersonOfPeople) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *PeoplePersonOfPeople) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *PeoplePersonOfPeople) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetNotifications

`func (o *PeoplePersonOfPeople) GetNotifications() PeopleNotifications`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *PeoplePersonOfPeople) GetNotificationsOk() (*PeopleNotifications, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *PeoplePersonOfPeople) SetNotifications(v PeopleNotifications)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *PeoplePersonOfPeople) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetOpenId

`func (o *PeoplePersonOfPeople) GetOpenId() string`

GetOpenId returns the OpenId field if non-nil, zero value otherwise.

### GetOpenIdOk

`func (o *PeoplePersonOfPeople) GetOpenIdOk() (*string, bool)`

GetOpenIdOk returns a tuple with the OpenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenId

`func (o *PeoplePersonOfPeople) SetOpenId(v string)`

SetOpenId sets OpenId field to given value.

### HasOpenId

`func (o *PeoplePersonOfPeople) HasOpenId() bool`

HasOpenId returns a boolean if a field has been set.

### GetPermissions

`func (o *PeoplePersonOfPeople) GetPermissions() PeoplePermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PeoplePersonOfPeople) GetPermissionsOk() (*PeoplePermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PeoplePersonOfPeople) SetPermissions(v PeoplePermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *PeoplePersonOfPeople) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetPhoneNumberFax

`func (o *PeoplePersonOfPeople) GetPhoneNumberFax() string`

GetPhoneNumberFax returns the PhoneNumberFax field if non-nil, zero value otherwise.

### GetPhoneNumberFaxOk

`func (o *PeoplePersonOfPeople) GetPhoneNumberFaxOk() (*string, bool)`

GetPhoneNumberFaxOk returns a tuple with the PhoneNumberFax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberFax

`func (o *PeoplePersonOfPeople) SetPhoneNumberFax(v string)`

SetPhoneNumberFax sets PhoneNumberFax field to given value.

### HasPhoneNumberFax

`func (o *PeoplePersonOfPeople) HasPhoneNumberFax() bool`

HasPhoneNumberFax returns a boolean if a field has been set.

### GetPhoneNumberHome

`func (o *PeoplePersonOfPeople) GetPhoneNumberHome() string`

GetPhoneNumberHome returns the PhoneNumberHome field if non-nil, zero value otherwise.

### GetPhoneNumberHomeOk

`func (o *PeoplePersonOfPeople) GetPhoneNumberHomeOk() (*string, bool)`

GetPhoneNumberHomeOk returns a tuple with the PhoneNumberHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberHome

`func (o *PeoplePersonOfPeople) SetPhoneNumberHome(v string)`

SetPhoneNumberHome sets PhoneNumberHome field to given value.

### HasPhoneNumberHome

`func (o *PeoplePersonOfPeople) HasPhoneNumberHome() bool`

HasPhoneNumberHome returns a boolean if a field has been set.

### GetPhoneNumberMobile

`func (o *PeoplePersonOfPeople) GetPhoneNumberMobile() string`

GetPhoneNumberMobile returns the PhoneNumberMobile field if non-nil, zero value otherwise.

### GetPhoneNumberMobileOk

`func (o *PeoplePersonOfPeople) GetPhoneNumberMobileOk() (*string, bool)`

GetPhoneNumberMobileOk returns a tuple with the PhoneNumberMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberMobile

`func (o *PeoplePersonOfPeople) SetPhoneNumberMobile(v string)`

SetPhoneNumberMobile sets PhoneNumberMobile field to given value.

### HasPhoneNumberMobile

`func (o *PeoplePersonOfPeople) HasPhoneNumberMobile() bool`

HasPhoneNumberMobile returns a boolean if a field has been set.

### GetPhoneNumberMobileParts

`func (o *PeoplePersonOfPeople) GetPhoneNumberMobileParts() PeoplePhoneParts`

GetPhoneNumberMobileParts returns the PhoneNumberMobileParts field if non-nil, zero value otherwise.

### GetPhoneNumberMobilePartsOk

`func (o *PeoplePersonOfPeople) GetPhoneNumberMobilePartsOk() (*PeoplePhoneParts, bool)`

GetPhoneNumberMobilePartsOk returns a tuple with the PhoneNumberMobileParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberMobileParts

`func (o *PeoplePersonOfPeople) SetPhoneNumberMobileParts(v PeoplePhoneParts)`

SetPhoneNumberMobileParts sets PhoneNumberMobileParts field to given value.

### HasPhoneNumberMobileParts

`func (o *PeoplePersonOfPeople) HasPhoneNumberMobileParts() bool`

HasPhoneNumberMobileParts returns a boolean if a field has been set.

### GetPhoneNumberOffice

`func (o *PeoplePersonOfPeople) GetPhoneNumberOffice() string`

GetPhoneNumberOffice returns the PhoneNumberOffice field if non-nil, zero value otherwise.

### GetPhoneNumberOfficeOk

`func (o *PeoplePersonOfPeople) GetPhoneNumberOfficeOk() (*string, bool)`

GetPhoneNumberOfficeOk returns a tuple with the PhoneNumberOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberOffice

`func (o *PeoplePersonOfPeople) SetPhoneNumberOffice(v string)`

SetPhoneNumberOffice sets PhoneNumberOffice field to given value.

### HasPhoneNumberOffice

`func (o *PeoplePersonOfPeople) HasPhoneNumberOffice() bool`

HasPhoneNumberOffice returns a boolean if a field has been set.

### GetPhoneNumberOfficeExt

`func (o *PeoplePersonOfPeople) GetPhoneNumberOfficeExt() string`

GetPhoneNumberOfficeExt returns the PhoneNumberOfficeExt field if non-nil, zero value otherwise.

### GetPhoneNumberOfficeExtOk

`func (o *PeoplePersonOfPeople) GetPhoneNumberOfficeExtOk() (*string, bool)`

GetPhoneNumberOfficeExtOk returns a tuple with the PhoneNumberOfficeExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberOfficeExt

`func (o *PeoplePersonOfPeople) SetPhoneNumberOfficeExt(v string)`

SetPhoneNumberOfficeExt sets PhoneNumberOfficeExt field to given value.

### HasPhoneNumberOfficeExt

`func (o *PeoplePersonOfPeople) HasPhoneNumberOfficeExt() bool`

HasPhoneNumberOfficeExt returns a boolean if a field has been set.

### GetPid

`func (o *PeoplePersonOfPeople) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *PeoplePersonOfPeople) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *PeoplePersonOfPeople) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *PeoplePersonOfPeople) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetPrivateNotes

`func (o *PeoplePersonOfPeople) GetPrivateNotes() string`

GetPrivateNotes returns the PrivateNotes field if non-nil, zero value otherwise.

### GetPrivateNotesOk

`func (o *PeoplePersonOfPeople) GetPrivateNotesOk() (*string, bool)`

GetPrivateNotesOk returns a tuple with the PrivateNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNotes

`func (o *PeoplePersonOfPeople) SetPrivateNotes(v string)`

SetPrivateNotes sets PrivateNotes field to given value.

### HasPrivateNotes

`func (o *PeoplePersonOfPeople) HasPrivateNotes() bool`

HasPrivateNotes returns a boolean if a field has been set.

### GetPrivateNotesText

`func (o *PeoplePersonOfPeople) GetPrivateNotesText() string`

GetPrivateNotesText returns the PrivateNotesText field if non-nil, zero value otherwise.

### GetPrivateNotesTextOk

`func (o *PeoplePersonOfPeople) GetPrivateNotesTextOk() (*string, bool)`

GetPrivateNotesTextOk returns a tuple with the PrivateNotesText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNotesText

`func (o *PeoplePersonOfPeople) SetPrivateNotesText(v string)`

SetPrivateNotesText sets PrivateNotesText field to given value.

### HasPrivateNotesText

`func (o *PeoplePersonOfPeople) HasPrivateNotesText() bool`

HasPrivateNotesText returns a boolean if a field has been set.

### GetProfile

`func (o *PeoplePersonOfPeople) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *PeoplePersonOfPeople) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *PeoplePersonOfPeople) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *PeoplePersonOfPeople) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetProfileText

`func (o *PeoplePersonOfPeople) GetProfileText() string`

GetProfileText returns the ProfileText field if non-nil, zero value otherwise.

### GetProfileTextOk

`func (o *PeoplePersonOfPeople) GetProfileTextOk() (*string, bool)`

GetProfileTextOk returns a tuple with the ProfileText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileText

`func (o *PeoplePersonOfPeople) SetProfileText(v string)`

SetProfileText sets ProfileText field to given value.

### HasProfileText

`func (o *PeoplePersonOfPeople) HasProfileText() bool`

HasProfileText returns a boolean if a field has been set.

### GetProjectRoles

`func (o *PeoplePersonOfPeople) GetProjectRoles() string`

GetProjectRoles returns the ProjectRoles field if non-nil, zero value otherwise.

### GetProjectRolesOk

`func (o *PeoplePersonOfPeople) GetProjectRolesOk() (*string, bool)`

GetProjectRolesOk returns a tuple with the ProjectRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectRoles

`func (o *PeoplePersonOfPeople) SetProjectRoles(v string)`

SetProjectRoles sets ProjectRoles field to given value.

### HasProjectRoles

`func (o *PeoplePersonOfPeople) HasProjectRoles() bool`

HasProjectRoles returns a boolean if a field has been set.

### GetProjects

`func (o *PeoplePersonOfPeople) GetProjects() []int32`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *PeoplePersonOfPeople) GetProjectsOk() (*[]int32, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *PeoplePersonOfPeople) SetProjects(v []int32)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *PeoplePersonOfPeople) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetSiteOwner

`func (o *PeoplePersonOfPeople) GetSiteOwner() bool`

GetSiteOwner returns the SiteOwner field if non-nil, zero value otherwise.

### GetSiteOwnerOk

`func (o *PeoplePersonOfPeople) GetSiteOwnerOk() (*bool, bool)`

GetSiteOwnerOk returns a tuple with the SiteOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteOwner

`func (o *PeoplePersonOfPeople) SetSiteOwner(v bool)`

SetSiteOwner sets SiteOwner field to given value.

### HasSiteOwner

`func (o *PeoplePersonOfPeople) HasSiteOwner() bool`

HasSiteOwner returns a boolean if a field has been set.

### GetSocial

`func (o *PeoplePersonOfPeople) GetSocial() PeopleSocial`

GetSocial returns the Social field if non-nil, zero value otherwise.

### GetSocialOk

`func (o *PeoplePersonOfPeople) GetSocialOk() (*PeopleSocial, bool)`

GetSocialOk returns a tuple with the Social field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocial

`func (o *PeoplePersonOfPeople) SetSocial(v PeopleSocial)`

SetSocial sets Social field to given value.

### HasSocial

`func (o *PeoplePersonOfPeople) HasSocial() bool`

HasSocial returns a boolean if a field has been set.

### GetTags

`func (o *PeoplePersonOfPeople) GetTags() []TagsTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PeoplePersonOfPeople) GetTagsOk() (*[]TagsTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PeoplePersonOfPeople) SetTags(v []TagsTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PeoplePersonOfPeople) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTeams

`func (o *PeoplePersonOfPeople) GetTeams() []string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *PeoplePersonOfPeople) GetTeamsOk() (*[]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *PeoplePersonOfPeople) SetTeams(v []string)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *PeoplePersonOfPeople) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetTextFormat

`func (o *PeoplePersonOfPeople) GetTextFormat() string`

GetTextFormat returns the TextFormat field if non-nil, zero value otherwise.

### GetTextFormatOk

`func (o *PeoplePersonOfPeople) GetTextFormatOk() (*string, bool)`

GetTextFormatOk returns a tuple with the TextFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextFormat

`func (o *PeoplePersonOfPeople) SetTextFormat(v string)`

SetTextFormat sets TextFormat field to given value.

### HasTextFormat

`func (o *PeoplePersonOfPeople) HasTextFormat() bool`

HasTextFormat returns a boolean if a field has been set.

### GetTitle

`func (o *PeoplePersonOfPeople) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PeoplePersonOfPeople) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PeoplePersonOfPeople) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PeoplePersonOfPeople) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTwitter

`func (o *PeoplePersonOfPeople) GetTwitter() string`

GetTwitter returns the Twitter field if non-nil, zero value otherwise.

### GetTwitterOk

`func (o *PeoplePersonOfPeople) GetTwitterOk() (*string, bool)`

GetTwitterOk returns a tuple with the Twitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitter

`func (o *PeoplePersonOfPeople) SetTwitter(v string)`

SetTwitter sets Twitter field to given value.

### HasTwitter

`func (o *PeoplePersonOfPeople) HasTwitter() bool`

HasTwitter returns a boolean if a field has been set.

### GetTwoFactorAuthEnabled

`func (o *PeoplePersonOfPeople) GetTwoFactorAuthEnabled() bool`

GetTwoFactorAuthEnabled returns the TwoFactorAuthEnabled field if non-nil, zero value otherwise.

### GetTwoFactorAuthEnabledOk

`func (o *PeoplePersonOfPeople) GetTwoFactorAuthEnabledOk() (*bool, bool)`

GetTwoFactorAuthEnabledOk returns a tuple with the TwoFactorAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorAuthEnabled

`func (o *PeoplePersonOfPeople) SetTwoFactorAuthEnabled(v bool)`

SetTwoFactorAuthEnabled sets TwoFactorAuthEnabled field to given value.

### HasTwoFactorAuthEnabled

`func (o *PeoplePersonOfPeople) HasTwoFactorAuthEnabled() bool`

HasTwoFactorAuthEnabled returns a boolean if a field has been set.

### GetUseShorthandDurations

`func (o *PeoplePersonOfPeople) GetUseShorthandDurations() bool`

GetUseShorthandDurations returns the UseShorthandDurations field if non-nil, zero value otherwise.

### GetUseShorthandDurationsOk

`func (o *PeoplePersonOfPeople) GetUseShorthandDurationsOk() (*bool, bool)`

GetUseShorthandDurationsOk returns a tuple with the UseShorthandDurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseShorthandDurations

`func (o *PeoplePersonOfPeople) SetUseShorthandDurations(v bool)`

SetUseShorthandDurations sets UseShorthandDurations field to given value.

### HasUseShorthandDurations

`func (o *PeoplePersonOfPeople) HasUseShorthandDurations() bool`

HasUseShorthandDurations returns a boolean if a field has been set.

### GetUserInvited

`func (o *PeoplePersonOfPeople) GetUserInvited() int32`

GetUserInvited returns the UserInvited field if non-nil, zero value otherwise.

### GetUserInvitedOk

`func (o *PeoplePersonOfPeople) GetUserInvitedOk() (*int32, bool)`

GetUserInvitedOk returns a tuple with the UserInvited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInvited

`func (o *PeoplePersonOfPeople) SetUserInvited(v int32)`

SetUserInvited sets UserInvited field to given value.

### HasUserInvited

`func (o *PeoplePersonOfPeople) HasUserInvited() bool`

HasUserInvited returns a boolean if a field has been set.

### GetUserInvitedDate

`func (o *PeoplePersonOfPeople) GetUserInvitedDate() string`

GetUserInvitedDate returns the UserInvitedDate field if non-nil, zero value otherwise.

### GetUserInvitedDateOk

`func (o *PeoplePersonOfPeople) GetUserInvitedDateOk() (*string, bool)`

GetUserInvitedDateOk returns a tuple with the UserInvitedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInvitedDate

`func (o *PeoplePersonOfPeople) SetUserInvitedDate(v string)`

SetUserInvitedDate sets UserInvitedDate field to given value.

### HasUserInvitedDate

`func (o *PeoplePersonOfPeople) HasUserInvitedDate() bool`

HasUserInvitedDate returns a boolean if a field has been set.

### GetUserInvitedStatus

`func (o *PeoplePersonOfPeople) GetUserInvitedStatus() string`

GetUserInvitedStatus returns the UserInvitedStatus field if non-nil, zero value otherwise.

### GetUserInvitedStatusOk

`func (o *PeoplePersonOfPeople) GetUserInvitedStatusOk() (*string, bool)`

GetUserInvitedStatusOk returns a tuple with the UserInvitedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInvitedStatus

`func (o *PeoplePersonOfPeople) SetUserInvitedStatus(v string)`

SetUserInvitedStatus sets UserInvitedStatus field to given value.

### HasUserInvitedStatus

`func (o *PeoplePersonOfPeople) HasUserInvitedStatus() bool`

HasUserInvitedStatus returns a boolean if a field has been set.

### GetUserName

`func (o *PeoplePersonOfPeople) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *PeoplePersonOfPeople) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *PeoplePersonOfPeople) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *PeoplePersonOfPeople) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetUserType

`func (o *PeoplePersonOfPeople) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *PeoplePersonOfPeople) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *PeoplePersonOfPeople) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *PeoplePersonOfPeople) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetUserUUID

`func (o *PeoplePersonOfPeople) GetUserUUID() string`

GetUserUUID returns the UserUUID field if non-nil, zero value otherwise.

### GetUserUUIDOk

`func (o *PeoplePersonOfPeople) GetUserUUIDOk() (*string, bool)`

GetUserUUIDOk returns a tuple with the UserUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUUID

`func (o *PeoplePersonOfPeople) SetUserUUID(v string)`

SetUserUUID sets UserUUID field to given value.

### HasUserUUID

`func (o *PeoplePersonOfPeople) HasUserUUID() bool`

HasUserUUID returns a boolean if a field has been set.

### GetWorkingHours

`func (o *PeoplePersonOfPeople) GetWorkingHours() PeopleWorkingHours`

GetWorkingHours returns the WorkingHours field if non-nil, zero value otherwise.

### GetWorkingHoursOk

`func (o *PeoplePersonOfPeople) GetWorkingHoursOk() (*PeopleWorkingHours, bool)`

GetWorkingHoursOk returns a tuple with the WorkingHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingHours

`func (o *PeoplePersonOfPeople) SetWorkingHours(v PeopleWorkingHours)`

SetWorkingHours sets WorkingHours field to given value.

### HasWorkingHours

`func (o *PeoplePersonOfPeople) HasWorkingHours() bool`

HasWorkingHours returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


