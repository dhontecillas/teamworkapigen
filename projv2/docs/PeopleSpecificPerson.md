# PeopleSpecificPerson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**APIEnabled** | Pointer to **bool** |  | [optional] 
**Accounts** | Pointer to [**[]UsersbaseAccount**](UsersbaseAccount.md) |  | [optional] 
**Address** | Pointer to [**PeopleAddress**](PeopleAddress.md) |  | [optional] 
**Administrator** | Pointer to **bool** |  | [optional] 
**Auth** | Pointer to [**PeopleAuthData**](PeopleAuthData.md) |  | [optional] 
**AvatarUrl** | Pointer to **string** |  | [optional] 
**CanManageProjectTemplates** | Pointer to **bool** |  | [optional] 
**CanViewProjectTemplates** | Pointer to **bool** |  | [optional] 
**Company** | Pointer to [**PeopleCompany**](PeopleCompany.md) |  | [optional] 
**CompanyId** | Pointer to **int32** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CurrentFeatureAnnouncement** | Pointer to [**AccountAnnouncementResponse**](AccountAnnouncementResponse.md) |  | [optional] 
**DefaultFilters** | Pointer to [**map[string]TwdatausersFilterResponse**](TwdatausersFilterResponse.md) |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**DocumentEditorInstalled** | Pointer to **bool** |  | [optional] 
**EmailAddress** | Pointer to **string** |  | [optional] 
**EmailAlt1** | Pointer to **string** |  | [optional] 
**EmailAlt2** | Pointer to **string** |  | [optional] 
**EmailAlt3** | Pointer to **string** |  | [optional] 
**FeatureUpdatesCount** | Pointer to **int32** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**HasAccessToNewProjects** | Pointer to **bool** |  | [optional] 
**HasDeskAccount** | Pointer to **bool** |  | [optional] 
**HasFeatureUpdates** | Pointer to **bool** |  | [optional] 
**HubspotEnabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**ImHandle** | Pointer to **string** |  | [optional] 
**ImService** | Pointer to **string** |  | [optional] 
**Impersonating** | Pointer to [**PeopleImpersonating**](PeopleImpersonating.md) |  | [optional] 
**InOwnerCompany** | Pointer to **bool** |  | [optional] 
**InitialPage** | Pointer to **string** |  | [optional] 
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
**Milestones** | Pointer to **map[string]interface{}** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Notifications** | Pointer to [**PeopleNotifications**](PeopleNotifications.md) |  | [optional] 
**NumActiveProjects** | Pointer to **int32** |  | [optional] 
**OpenId** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to [**PeoplePermissions**](PeoplePermissions.md) |  | [optional] 
**PhoneNumberFax** | Pointer to **string** |  | [optional] 
**PhoneNumberHome** | Pointer to **string** |  | [optional] 
**PhoneNumberMobile** | Pointer to **string** |  | [optional] 
**PhoneNumberMobileParts** | Pointer to [**PeoplePhoneParts**](PeoplePhoneParts.md) |  | [optional] 
**PhoneNumberOffice** | Pointer to **string** |  | [optional] 
**PhoneNumberOfficeExt** | Pointer to **string** |  | [optional] 
**Pid** | Pointer to **string** |  | [optional] 
**Preferences** | Pointer to **map[string]interface{}** |  | [optional] 
**PrivateNotes** | Pointer to **string** |  | [optional] 
**PrivateNotesText** | Pointer to **string** |  | [optional] 
**Profile** | Pointer to **string** |  | [optional] 
**ProfileText** | Pointer to **string** |  | [optional] 
**Projects** | Pointer to **map[string]interface{}** |  | [optional] 
**SharedUserFilter** | Pointer to [**map[string]TwdatausersFilterResponse**](TwdatausersFilterResponse.md) |  | [optional] 
**SiteOwner** | Pointer to **bool** |  | [optional] 
**Social** | Pointer to [**PeopleSocial**](PeopleSocial.md) |  | [optional] 
**Tags** | Pointer to [**[]TagsTag**](TagsTag.md) |  | [optional] 
**Tasks** | Pointer to **map[string]interface{}** |  | [optional] 
**Teams** | Pointer to **[]string** |  | [optional] 
**TeamsCount** | Pointer to **int32** |  | [optional] 
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

### NewPeopleSpecificPerson

`func NewPeopleSpecificPerson() *PeopleSpecificPerson`

NewPeopleSpecificPerson instantiates a new PeopleSpecificPerson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeopleSpecificPersonWithDefaults

`func NewPeopleSpecificPersonWithDefaults() *PeopleSpecificPerson`

NewPeopleSpecificPersonWithDefaults instantiates a new PeopleSpecificPerson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAPIEnabled

`func (o *PeopleSpecificPerson) GetAPIEnabled() bool`

GetAPIEnabled returns the APIEnabled field if non-nil, zero value otherwise.

### GetAPIEnabledOk

`func (o *PeopleSpecificPerson) GetAPIEnabledOk() (*bool, bool)`

GetAPIEnabledOk returns a tuple with the APIEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPIEnabled

`func (o *PeopleSpecificPerson) SetAPIEnabled(v bool)`

SetAPIEnabled sets APIEnabled field to given value.

### HasAPIEnabled

`func (o *PeopleSpecificPerson) HasAPIEnabled() bool`

HasAPIEnabled returns a boolean if a field has been set.

### GetAccounts

`func (o *PeopleSpecificPerson) GetAccounts() []UsersbaseAccount`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *PeopleSpecificPerson) GetAccountsOk() (*[]UsersbaseAccount, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *PeopleSpecificPerson) SetAccounts(v []UsersbaseAccount)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *PeopleSpecificPerson) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetAddress

`func (o *PeopleSpecificPerson) GetAddress() PeopleAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PeopleSpecificPerson) GetAddressOk() (*PeopleAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PeopleSpecificPerson) SetAddress(v PeopleAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PeopleSpecificPerson) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAdministrator

`func (o *PeopleSpecificPerson) GetAdministrator() bool`

GetAdministrator returns the Administrator field if non-nil, zero value otherwise.

### GetAdministratorOk

`func (o *PeopleSpecificPerson) GetAdministratorOk() (*bool, bool)`

GetAdministratorOk returns a tuple with the Administrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrator

`func (o *PeopleSpecificPerson) SetAdministrator(v bool)`

SetAdministrator sets Administrator field to given value.

### HasAdministrator

`func (o *PeopleSpecificPerson) HasAdministrator() bool`

HasAdministrator returns a boolean if a field has been set.

### GetAuth

`func (o *PeopleSpecificPerson) GetAuth() PeopleAuthData`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *PeopleSpecificPerson) GetAuthOk() (*PeopleAuthData, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *PeopleSpecificPerson) SetAuth(v PeopleAuthData)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *PeopleSpecificPerson) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *PeopleSpecificPerson) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *PeopleSpecificPerson) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *PeopleSpecificPerson) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *PeopleSpecificPerson) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetCanManageProjectTemplates

`func (o *PeopleSpecificPerson) GetCanManageProjectTemplates() bool`

GetCanManageProjectTemplates returns the CanManageProjectTemplates field if non-nil, zero value otherwise.

### GetCanManageProjectTemplatesOk

`func (o *PeopleSpecificPerson) GetCanManageProjectTemplatesOk() (*bool, bool)`

GetCanManageProjectTemplatesOk returns a tuple with the CanManageProjectTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageProjectTemplates

`func (o *PeopleSpecificPerson) SetCanManageProjectTemplates(v bool)`

SetCanManageProjectTemplates sets CanManageProjectTemplates field to given value.

### HasCanManageProjectTemplates

`func (o *PeopleSpecificPerson) HasCanManageProjectTemplates() bool`

HasCanManageProjectTemplates returns a boolean if a field has been set.

### GetCanViewProjectTemplates

`func (o *PeopleSpecificPerson) GetCanViewProjectTemplates() bool`

GetCanViewProjectTemplates returns the CanViewProjectTemplates field if non-nil, zero value otherwise.

### GetCanViewProjectTemplatesOk

`func (o *PeopleSpecificPerson) GetCanViewProjectTemplatesOk() (*bool, bool)`

GetCanViewProjectTemplatesOk returns a tuple with the CanViewProjectTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewProjectTemplates

`func (o *PeopleSpecificPerson) SetCanViewProjectTemplates(v bool)`

SetCanViewProjectTemplates sets CanViewProjectTemplates field to given value.

### HasCanViewProjectTemplates

`func (o *PeopleSpecificPerson) HasCanViewProjectTemplates() bool`

HasCanViewProjectTemplates returns a boolean if a field has been set.

### GetCompany

`func (o *PeopleSpecificPerson) GetCompany() PeopleCompany`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *PeopleSpecificPerson) GetCompanyOk() (*PeopleCompany, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *PeopleSpecificPerson) SetCompany(v PeopleCompany)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *PeopleSpecificPerson) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCompanyId

`func (o *PeopleSpecificPerson) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *PeopleSpecificPerson) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *PeopleSpecificPerson) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *PeopleSpecificPerson) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCompanyName

`func (o *PeopleSpecificPerson) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *PeopleSpecificPerson) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *PeopleSpecificPerson) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *PeopleSpecificPerson) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PeopleSpecificPerson) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PeopleSpecificPerson) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PeopleSpecificPerson) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PeopleSpecificPerson) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrentFeatureAnnouncement

`func (o *PeopleSpecificPerson) GetCurrentFeatureAnnouncement() AccountAnnouncementResponse`

GetCurrentFeatureAnnouncement returns the CurrentFeatureAnnouncement field if non-nil, zero value otherwise.

### GetCurrentFeatureAnnouncementOk

`func (o *PeopleSpecificPerson) GetCurrentFeatureAnnouncementOk() (*AccountAnnouncementResponse, bool)`

GetCurrentFeatureAnnouncementOk returns a tuple with the CurrentFeatureAnnouncement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentFeatureAnnouncement

`func (o *PeopleSpecificPerson) SetCurrentFeatureAnnouncement(v AccountAnnouncementResponse)`

SetCurrentFeatureAnnouncement sets CurrentFeatureAnnouncement field to given value.

### HasCurrentFeatureAnnouncement

`func (o *PeopleSpecificPerson) HasCurrentFeatureAnnouncement() bool`

HasCurrentFeatureAnnouncement returns a boolean if a field has been set.

### GetDefaultFilters

`func (o *PeopleSpecificPerson) GetDefaultFilters() map[string]TwdatausersFilterResponse`

GetDefaultFilters returns the DefaultFilters field if non-nil, zero value otherwise.

### GetDefaultFiltersOk

`func (o *PeopleSpecificPerson) GetDefaultFiltersOk() (*map[string]TwdatausersFilterResponse, bool)`

GetDefaultFiltersOk returns a tuple with the DefaultFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFilters

`func (o *PeopleSpecificPerson) SetDefaultFilters(v map[string]TwdatausersFilterResponse)`

SetDefaultFilters sets DefaultFilters field to given value.

### HasDefaultFilters

`func (o *PeopleSpecificPerson) HasDefaultFilters() bool`

HasDefaultFilters returns a boolean if a field has been set.

### GetDeleted

`func (o *PeopleSpecificPerson) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *PeopleSpecificPerson) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *PeopleSpecificPerson) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *PeopleSpecificPerson) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDocumentEditorInstalled

`func (o *PeopleSpecificPerson) GetDocumentEditorInstalled() bool`

GetDocumentEditorInstalled returns the DocumentEditorInstalled field if non-nil, zero value otherwise.

### GetDocumentEditorInstalledOk

`func (o *PeopleSpecificPerson) GetDocumentEditorInstalledOk() (*bool, bool)`

GetDocumentEditorInstalledOk returns a tuple with the DocumentEditorInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentEditorInstalled

`func (o *PeopleSpecificPerson) SetDocumentEditorInstalled(v bool)`

SetDocumentEditorInstalled sets DocumentEditorInstalled field to given value.

### HasDocumentEditorInstalled

`func (o *PeopleSpecificPerson) HasDocumentEditorInstalled() bool`

HasDocumentEditorInstalled returns a boolean if a field has been set.

### GetEmailAddress

`func (o *PeopleSpecificPerson) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *PeopleSpecificPerson) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *PeopleSpecificPerson) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *PeopleSpecificPerson) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetEmailAlt1

`func (o *PeopleSpecificPerson) GetEmailAlt1() string`

GetEmailAlt1 returns the EmailAlt1 field if non-nil, zero value otherwise.

### GetEmailAlt1Ok

`func (o *PeopleSpecificPerson) GetEmailAlt1Ok() (*string, bool)`

GetEmailAlt1Ok returns a tuple with the EmailAlt1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAlt1

`func (o *PeopleSpecificPerson) SetEmailAlt1(v string)`

SetEmailAlt1 sets EmailAlt1 field to given value.

### HasEmailAlt1

`func (o *PeopleSpecificPerson) HasEmailAlt1() bool`

HasEmailAlt1 returns a boolean if a field has been set.

### GetEmailAlt2

`func (o *PeopleSpecificPerson) GetEmailAlt2() string`

GetEmailAlt2 returns the EmailAlt2 field if non-nil, zero value otherwise.

### GetEmailAlt2Ok

`func (o *PeopleSpecificPerson) GetEmailAlt2Ok() (*string, bool)`

GetEmailAlt2Ok returns a tuple with the EmailAlt2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAlt2

`func (o *PeopleSpecificPerson) SetEmailAlt2(v string)`

SetEmailAlt2 sets EmailAlt2 field to given value.

### HasEmailAlt2

`func (o *PeopleSpecificPerson) HasEmailAlt2() bool`

HasEmailAlt2 returns a boolean if a field has been set.

### GetEmailAlt3

`func (o *PeopleSpecificPerson) GetEmailAlt3() string`

GetEmailAlt3 returns the EmailAlt3 field if non-nil, zero value otherwise.

### GetEmailAlt3Ok

`func (o *PeopleSpecificPerson) GetEmailAlt3Ok() (*string, bool)`

GetEmailAlt3Ok returns a tuple with the EmailAlt3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAlt3

`func (o *PeopleSpecificPerson) SetEmailAlt3(v string)`

SetEmailAlt3 sets EmailAlt3 field to given value.

### HasEmailAlt3

`func (o *PeopleSpecificPerson) HasEmailAlt3() bool`

HasEmailAlt3 returns a boolean if a field has been set.

### GetFeatureUpdatesCount

`func (o *PeopleSpecificPerson) GetFeatureUpdatesCount() int32`

GetFeatureUpdatesCount returns the FeatureUpdatesCount field if non-nil, zero value otherwise.

### GetFeatureUpdatesCountOk

`func (o *PeopleSpecificPerson) GetFeatureUpdatesCountOk() (*int32, bool)`

GetFeatureUpdatesCountOk returns a tuple with the FeatureUpdatesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureUpdatesCount

`func (o *PeopleSpecificPerson) SetFeatureUpdatesCount(v int32)`

SetFeatureUpdatesCount sets FeatureUpdatesCount field to given value.

### HasFeatureUpdatesCount

`func (o *PeopleSpecificPerson) HasFeatureUpdatesCount() bool`

HasFeatureUpdatesCount returns a boolean if a field has been set.

### GetFirstName

`func (o *PeopleSpecificPerson) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PeopleSpecificPerson) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PeopleSpecificPerson) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PeopleSpecificPerson) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetHasAccessToNewProjects

`func (o *PeopleSpecificPerson) GetHasAccessToNewProjects() bool`

GetHasAccessToNewProjects returns the HasAccessToNewProjects field if non-nil, zero value otherwise.

### GetHasAccessToNewProjectsOk

`func (o *PeopleSpecificPerson) GetHasAccessToNewProjectsOk() (*bool, bool)`

GetHasAccessToNewProjectsOk returns a tuple with the HasAccessToNewProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccessToNewProjects

`func (o *PeopleSpecificPerson) SetHasAccessToNewProjects(v bool)`

SetHasAccessToNewProjects sets HasAccessToNewProjects field to given value.

### HasHasAccessToNewProjects

`func (o *PeopleSpecificPerson) HasHasAccessToNewProjects() bool`

HasHasAccessToNewProjects returns a boolean if a field has been set.

### GetHasDeskAccount

`func (o *PeopleSpecificPerson) GetHasDeskAccount() bool`

GetHasDeskAccount returns the HasDeskAccount field if non-nil, zero value otherwise.

### GetHasDeskAccountOk

`func (o *PeopleSpecificPerson) GetHasDeskAccountOk() (*bool, bool)`

GetHasDeskAccountOk returns a tuple with the HasDeskAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDeskAccount

`func (o *PeopleSpecificPerson) SetHasDeskAccount(v bool)`

SetHasDeskAccount sets HasDeskAccount field to given value.

### HasHasDeskAccount

`func (o *PeopleSpecificPerson) HasHasDeskAccount() bool`

HasHasDeskAccount returns a boolean if a field has been set.

### GetHasFeatureUpdates

`func (o *PeopleSpecificPerson) GetHasFeatureUpdates() bool`

GetHasFeatureUpdates returns the HasFeatureUpdates field if non-nil, zero value otherwise.

### GetHasFeatureUpdatesOk

`func (o *PeopleSpecificPerson) GetHasFeatureUpdatesOk() (*bool, bool)`

GetHasFeatureUpdatesOk returns a tuple with the HasFeatureUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFeatureUpdates

`func (o *PeopleSpecificPerson) SetHasFeatureUpdates(v bool)`

SetHasFeatureUpdates sets HasFeatureUpdates field to given value.

### HasHasFeatureUpdates

`func (o *PeopleSpecificPerson) HasHasFeatureUpdates() bool`

HasHasFeatureUpdates returns a boolean if a field has been set.

### GetHubspotEnabled

`func (o *PeopleSpecificPerson) GetHubspotEnabled() bool`

GetHubspotEnabled returns the HubspotEnabled field if non-nil, zero value otherwise.

### GetHubspotEnabledOk

`func (o *PeopleSpecificPerson) GetHubspotEnabledOk() (*bool, bool)`

GetHubspotEnabledOk returns a tuple with the HubspotEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubspotEnabled

`func (o *PeopleSpecificPerson) SetHubspotEnabled(v bool)`

SetHubspotEnabled sets HubspotEnabled field to given value.

### HasHubspotEnabled

`func (o *PeopleSpecificPerson) HasHubspotEnabled() bool`

HasHubspotEnabled returns a boolean if a field has been set.

### GetId

`func (o *PeopleSpecificPerson) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PeopleSpecificPerson) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PeopleSpecificPerson) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PeopleSpecificPerson) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImHandle

`func (o *PeopleSpecificPerson) GetImHandle() string`

GetImHandle returns the ImHandle field if non-nil, zero value otherwise.

### GetImHandleOk

`func (o *PeopleSpecificPerson) GetImHandleOk() (*string, bool)`

GetImHandleOk returns a tuple with the ImHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImHandle

`func (o *PeopleSpecificPerson) SetImHandle(v string)`

SetImHandle sets ImHandle field to given value.

### HasImHandle

`func (o *PeopleSpecificPerson) HasImHandle() bool`

HasImHandle returns a boolean if a field has been set.

### GetImService

`func (o *PeopleSpecificPerson) GetImService() string`

GetImService returns the ImService field if non-nil, zero value otherwise.

### GetImServiceOk

`func (o *PeopleSpecificPerson) GetImServiceOk() (*string, bool)`

GetImServiceOk returns a tuple with the ImService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImService

`func (o *PeopleSpecificPerson) SetImService(v string)`

SetImService sets ImService field to given value.

### HasImService

`func (o *PeopleSpecificPerson) HasImService() bool`

HasImService returns a boolean if a field has been set.

### GetImpersonating

`func (o *PeopleSpecificPerson) GetImpersonating() PeopleImpersonating`

GetImpersonating returns the Impersonating field if non-nil, zero value otherwise.

### GetImpersonatingOk

`func (o *PeopleSpecificPerson) GetImpersonatingOk() (*PeopleImpersonating, bool)`

GetImpersonatingOk returns a tuple with the Impersonating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpersonating

`func (o *PeopleSpecificPerson) SetImpersonating(v PeopleImpersonating)`

SetImpersonating sets Impersonating field to given value.

### HasImpersonating

`func (o *PeopleSpecificPerson) HasImpersonating() bool`

HasImpersonating returns a boolean if a field has been set.

### GetInOwnerCompany

`func (o *PeopleSpecificPerson) GetInOwnerCompany() bool`

GetInOwnerCompany returns the InOwnerCompany field if non-nil, zero value otherwise.

### GetInOwnerCompanyOk

`func (o *PeopleSpecificPerson) GetInOwnerCompanyOk() (*bool, bool)`

GetInOwnerCompanyOk returns a tuple with the InOwnerCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInOwnerCompany

`func (o *PeopleSpecificPerson) SetInOwnerCompany(v bool)`

SetInOwnerCompany sets InOwnerCompany field to given value.

### HasInOwnerCompany

`func (o *PeopleSpecificPerson) HasInOwnerCompany() bool`

HasInOwnerCompany returns a boolean if a field has been set.

### GetInitialPage

`func (o *PeopleSpecificPerson) GetInitialPage() string`

GetInitialPage returns the InitialPage field if non-nil, zero value otherwise.

### GetInitialPageOk

`func (o *PeopleSpecificPerson) GetInitialPageOk() (*string, bool)`

GetInitialPageOk returns a tuple with the InitialPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialPage

`func (o *PeopleSpecificPerson) SetInitialPage(v string)`

SetInitialPage sets InitialPage field to given value.

### HasInitialPage

`func (o *PeopleSpecificPerson) HasInitialPage() bool`

HasInitialPage returns a boolean if a field has been set.

### GetIsClientUser

`func (o *PeopleSpecificPerson) GetIsClientUser() bool`

GetIsClientUser returns the IsClientUser field if non-nil, zero value otherwise.

### GetIsClientUserOk

`func (o *PeopleSpecificPerson) GetIsClientUserOk() (*bool, bool)`

GetIsClientUserOk returns a tuple with the IsClientUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClientUser

`func (o *PeopleSpecificPerson) SetIsClientUser(v bool)`

SetIsClientUser sets IsClientUser field to given value.

### HasIsClientUser

`func (o *PeopleSpecificPerson) HasIsClientUser() bool`

HasIsClientUser returns a boolean if a field has been set.

### GetIsClockedIn

`func (o *PeopleSpecificPerson) GetIsClockedIn() bool`

GetIsClockedIn returns the IsClockedIn field if non-nil, zero value otherwise.

### GetIsClockedInOk

`func (o *PeopleSpecificPerson) GetIsClockedInOk() (*bool, bool)`

GetIsClockedInOk returns a tuple with the IsClockedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClockedIn

`func (o *PeopleSpecificPerson) SetIsClockedIn(v bool)`

SetIsClockedIn sets IsClockedIn field to given value.

### HasIsClockedIn

`func (o *PeopleSpecificPerson) HasIsClockedIn() bool`

HasIsClockedIn returns a boolean if a field has been set.

### GetLastActive

`func (o *PeopleSpecificPerson) GetLastActive() string`

GetLastActive returns the LastActive field if non-nil, zero value otherwise.

### GetLastActiveOk

`func (o *PeopleSpecificPerson) GetLastActiveOk() (*string, bool)`

GetLastActiveOk returns a tuple with the LastActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActive

`func (o *PeopleSpecificPerson) SetLastActive(v string)`

SetLastActive sets LastActive field to given value.

### HasLastActive

`func (o *PeopleSpecificPerson) HasLastActive() bool`

HasLastActive returns a boolean if a field has been set.

### GetLastChangedOn

`func (o *PeopleSpecificPerson) GetLastChangedOn() string`

GetLastChangedOn returns the LastChangedOn field if non-nil, zero value otherwise.

### GetLastChangedOnOk

`func (o *PeopleSpecificPerson) GetLastChangedOnOk() (*string, bool)`

GetLastChangedOnOk returns a tuple with the LastChangedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangedOn

`func (o *PeopleSpecificPerson) SetLastChangedOn(v string)`

SetLastChangedOn sets LastChangedOn field to given value.

### HasLastChangedOn

`func (o *PeopleSpecificPerson) HasLastChangedOn() bool`

HasLastChangedOn returns a boolean if a field has been set.

### GetLastLogin

`func (o *PeopleSpecificPerson) GetLastLogin() string`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *PeopleSpecificPerson) GetLastLoginOk() (*string, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *PeopleSpecificPerson) SetLastLogin(v string)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *PeopleSpecificPerson) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetLastName

`func (o *PeopleSpecificPerson) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PeopleSpecificPerson) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PeopleSpecificPerson) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PeopleSpecificPerson) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLengthOfDay

`func (o *PeopleSpecificPerson) GetLengthOfDay() float32`

GetLengthOfDay returns the LengthOfDay field if non-nil, zero value otherwise.

### GetLengthOfDayOk

`func (o *PeopleSpecificPerson) GetLengthOfDayOk() (*float32, bool)`

GetLengthOfDayOk returns a tuple with the LengthOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthOfDay

`func (o *PeopleSpecificPerson) SetLengthOfDay(v float32)`

SetLengthOfDay sets LengthOfDay field to given value.

### HasLengthOfDay

`func (o *PeopleSpecificPerson) HasLengthOfDay() bool`

HasLengthOfDay returns a boolean if a field has been set.

### GetLocalization

`func (o *PeopleSpecificPerson) GetLocalization() PeopleLocalization`

GetLocalization returns the Localization field if non-nil, zero value otherwise.

### GetLocalizationOk

`func (o *PeopleSpecificPerson) GetLocalizationOk() (*PeopleLocalization, bool)`

GetLocalizationOk returns a tuple with the Localization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalization

`func (o *PeopleSpecificPerson) SetLocalization(v PeopleLocalization)`

SetLocalization sets Localization field to given value.

### HasLocalization

`func (o *PeopleSpecificPerson) HasLocalization() bool`

HasLocalization returns a boolean if a field has been set.

### GetLoginCount

`func (o *PeopleSpecificPerson) GetLoginCount() int32`

GetLoginCount returns the LoginCount field if non-nil, zero value otherwise.

### GetLoginCountOk

`func (o *PeopleSpecificPerson) GetLoginCountOk() (*int32, bool)`

GetLoginCountOk returns a tuple with the LoginCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCount

`func (o *PeopleSpecificPerson) SetLoginCount(v int32)`

SetLoginCount sets LoginCount field to given value.

### HasLoginCount

`func (o *PeopleSpecificPerson) HasLoginCount() bool`

HasLoginCount returns a boolean if a field has been set.

### GetMentionName

`func (o *PeopleSpecificPerson) GetMentionName() string`

GetMentionName returns the MentionName field if non-nil, zero value otherwise.

### GetMentionNameOk

`func (o *PeopleSpecificPerson) GetMentionNameOk() (*string, bool)`

GetMentionNameOk returns a tuple with the MentionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionName

`func (o *PeopleSpecificPerson) SetMentionName(v string)`

SetMentionName sets MentionName field to given value.

### HasMentionName

`func (o *PeopleSpecificPerson) HasMentionName() bool`

HasMentionName returns a boolean if a field has been set.

### GetMilestones

`func (o *PeopleSpecificPerson) GetMilestones() map[string]interface{}`

GetMilestones returns the Milestones field if non-nil, zero value otherwise.

### GetMilestonesOk

`func (o *PeopleSpecificPerson) GetMilestonesOk() (*map[string]interface{}, bool)`

GetMilestonesOk returns a tuple with the Milestones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestones

`func (o *PeopleSpecificPerson) SetMilestones(v map[string]interface{})`

SetMilestones sets Milestones field to given value.

### HasMilestones

`func (o *PeopleSpecificPerson) HasMilestones() bool`

HasMilestones returns a boolean if a field has been set.

### GetNotes

`func (o *PeopleSpecificPerson) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *PeopleSpecificPerson) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *PeopleSpecificPerson) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *PeopleSpecificPerson) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetNotifications

`func (o *PeopleSpecificPerson) GetNotifications() PeopleNotifications`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *PeopleSpecificPerson) GetNotificationsOk() (*PeopleNotifications, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *PeopleSpecificPerson) SetNotifications(v PeopleNotifications)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *PeopleSpecificPerson) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetNumActiveProjects

`func (o *PeopleSpecificPerson) GetNumActiveProjects() int32`

GetNumActiveProjects returns the NumActiveProjects field if non-nil, zero value otherwise.

### GetNumActiveProjectsOk

`func (o *PeopleSpecificPerson) GetNumActiveProjectsOk() (*int32, bool)`

GetNumActiveProjectsOk returns a tuple with the NumActiveProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumActiveProjects

`func (o *PeopleSpecificPerson) SetNumActiveProjects(v int32)`

SetNumActiveProjects sets NumActiveProjects field to given value.

### HasNumActiveProjects

`func (o *PeopleSpecificPerson) HasNumActiveProjects() bool`

HasNumActiveProjects returns a boolean if a field has been set.

### GetOpenId

`func (o *PeopleSpecificPerson) GetOpenId() string`

GetOpenId returns the OpenId field if non-nil, zero value otherwise.

### GetOpenIdOk

`func (o *PeopleSpecificPerson) GetOpenIdOk() (*string, bool)`

GetOpenIdOk returns a tuple with the OpenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenId

`func (o *PeopleSpecificPerson) SetOpenId(v string)`

SetOpenId sets OpenId field to given value.

### HasOpenId

`func (o *PeopleSpecificPerson) HasOpenId() bool`

HasOpenId returns a boolean if a field has been set.

### GetPermissions

`func (o *PeopleSpecificPerson) GetPermissions() PeoplePermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PeopleSpecificPerson) GetPermissionsOk() (*PeoplePermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PeopleSpecificPerson) SetPermissions(v PeoplePermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *PeopleSpecificPerson) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetPhoneNumberFax

`func (o *PeopleSpecificPerson) GetPhoneNumberFax() string`

GetPhoneNumberFax returns the PhoneNumberFax field if non-nil, zero value otherwise.

### GetPhoneNumberFaxOk

`func (o *PeopleSpecificPerson) GetPhoneNumberFaxOk() (*string, bool)`

GetPhoneNumberFaxOk returns a tuple with the PhoneNumberFax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberFax

`func (o *PeopleSpecificPerson) SetPhoneNumberFax(v string)`

SetPhoneNumberFax sets PhoneNumberFax field to given value.

### HasPhoneNumberFax

`func (o *PeopleSpecificPerson) HasPhoneNumberFax() bool`

HasPhoneNumberFax returns a boolean if a field has been set.

### GetPhoneNumberHome

`func (o *PeopleSpecificPerson) GetPhoneNumberHome() string`

GetPhoneNumberHome returns the PhoneNumberHome field if non-nil, zero value otherwise.

### GetPhoneNumberHomeOk

`func (o *PeopleSpecificPerson) GetPhoneNumberHomeOk() (*string, bool)`

GetPhoneNumberHomeOk returns a tuple with the PhoneNumberHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberHome

`func (o *PeopleSpecificPerson) SetPhoneNumberHome(v string)`

SetPhoneNumberHome sets PhoneNumberHome field to given value.

### HasPhoneNumberHome

`func (o *PeopleSpecificPerson) HasPhoneNumberHome() bool`

HasPhoneNumberHome returns a boolean if a field has been set.

### GetPhoneNumberMobile

`func (o *PeopleSpecificPerson) GetPhoneNumberMobile() string`

GetPhoneNumberMobile returns the PhoneNumberMobile field if non-nil, zero value otherwise.

### GetPhoneNumberMobileOk

`func (o *PeopleSpecificPerson) GetPhoneNumberMobileOk() (*string, bool)`

GetPhoneNumberMobileOk returns a tuple with the PhoneNumberMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberMobile

`func (o *PeopleSpecificPerson) SetPhoneNumberMobile(v string)`

SetPhoneNumberMobile sets PhoneNumberMobile field to given value.

### HasPhoneNumberMobile

`func (o *PeopleSpecificPerson) HasPhoneNumberMobile() bool`

HasPhoneNumberMobile returns a boolean if a field has been set.

### GetPhoneNumberMobileParts

`func (o *PeopleSpecificPerson) GetPhoneNumberMobileParts() PeoplePhoneParts`

GetPhoneNumberMobileParts returns the PhoneNumberMobileParts field if non-nil, zero value otherwise.

### GetPhoneNumberMobilePartsOk

`func (o *PeopleSpecificPerson) GetPhoneNumberMobilePartsOk() (*PeoplePhoneParts, bool)`

GetPhoneNumberMobilePartsOk returns a tuple with the PhoneNumberMobileParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberMobileParts

`func (o *PeopleSpecificPerson) SetPhoneNumberMobileParts(v PeoplePhoneParts)`

SetPhoneNumberMobileParts sets PhoneNumberMobileParts field to given value.

### HasPhoneNumberMobileParts

`func (o *PeopleSpecificPerson) HasPhoneNumberMobileParts() bool`

HasPhoneNumberMobileParts returns a boolean if a field has been set.

### GetPhoneNumberOffice

`func (o *PeopleSpecificPerson) GetPhoneNumberOffice() string`

GetPhoneNumberOffice returns the PhoneNumberOffice field if non-nil, zero value otherwise.

### GetPhoneNumberOfficeOk

`func (o *PeopleSpecificPerson) GetPhoneNumberOfficeOk() (*string, bool)`

GetPhoneNumberOfficeOk returns a tuple with the PhoneNumberOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberOffice

`func (o *PeopleSpecificPerson) SetPhoneNumberOffice(v string)`

SetPhoneNumberOffice sets PhoneNumberOffice field to given value.

### HasPhoneNumberOffice

`func (o *PeopleSpecificPerson) HasPhoneNumberOffice() bool`

HasPhoneNumberOffice returns a boolean if a field has been set.

### GetPhoneNumberOfficeExt

`func (o *PeopleSpecificPerson) GetPhoneNumberOfficeExt() string`

GetPhoneNumberOfficeExt returns the PhoneNumberOfficeExt field if non-nil, zero value otherwise.

### GetPhoneNumberOfficeExtOk

`func (o *PeopleSpecificPerson) GetPhoneNumberOfficeExtOk() (*string, bool)`

GetPhoneNumberOfficeExtOk returns a tuple with the PhoneNumberOfficeExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberOfficeExt

`func (o *PeopleSpecificPerson) SetPhoneNumberOfficeExt(v string)`

SetPhoneNumberOfficeExt sets PhoneNumberOfficeExt field to given value.

### HasPhoneNumberOfficeExt

`func (o *PeopleSpecificPerson) HasPhoneNumberOfficeExt() bool`

HasPhoneNumberOfficeExt returns a boolean if a field has been set.

### GetPid

`func (o *PeopleSpecificPerson) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *PeopleSpecificPerson) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *PeopleSpecificPerson) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *PeopleSpecificPerson) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetPreferences

`func (o *PeopleSpecificPerson) GetPreferences() map[string]interface{}`

GetPreferences returns the Preferences field if non-nil, zero value otherwise.

### GetPreferencesOk

`func (o *PeopleSpecificPerson) GetPreferencesOk() (*map[string]interface{}, bool)`

GetPreferencesOk returns a tuple with the Preferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferences

`func (o *PeopleSpecificPerson) SetPreferences(v map[string]interface{})`

SetPreferences sets Preferences field to given value.

### HasPreferences

`func (o *PeopleSpecificPerson) HasPreferences() bool`

HasPreferences returns a boolean if a field has been set.

### GetPrivateNotes

`func (o *PeopleSpecificPerson) GetPrivateNotes() string`

GetPrivateNotes returns the PrivateNotes field if non-nil, zero value otherwise.

### GetPrivateNotesOk

`func (o *PeopleSpecificPerson) GetPrivateNotesOk() (*string, bool)`

GetPrivateNotesOk returns a tuple with the PrivateNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNotes

`func (o *PeopleSpecificPerson) SetPrivateNotes(v string)`

SetPrivateNotes sets PrivateNotes field to given value.

### HasPrivateNotes

`func (o *PeopleSpecificPerson) HasPrivateNotes() bool`

HasPrivateNotes returns a boolean if a field has been set.

### GetPrivateNotesText

`func (o *PeopleSpecificPerson) GetPrivateNotesText() string`

GetPrivateNotesText returns the PrivateNotesText field if non-nil, zero value otherwise.

### GetPrivateNotesTextOk

`func (o *PeopleSpecificPerson) GetPrivateNotesTextOk() (*string, bool)`

GetPrivateNotesTextOk returns a tuple with the PrivateNotesText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNotesText

`func (o *PeopleSpecificPerson) SetPrivateNotesText(v string)`

SetPrivateNotesText sets PrivateNotesText field to given value.

### HasPrivateNotesText

`func (o *PeopleSpecificPerson) HasPrivateNotesText() bool`

HasPrivateNotesText returns a boolean if a field has been set.

### GetProfile

`func (o *PeopleSpecificPerson) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *PeopleSpecificPerson) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *PeopleSpecificPerson) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *PeopleSpecificPerson) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetProfileText

`func (o *PeopleSpecificPerson) GetProfileText() string`

GetProfileText returns the ProfileText field if non-nil, zero value otherwise.

### GetProfileTextOk

`func (o *PeopleSpecificPerson) GetProfileTextOk() (*string, bool)`

GetProfileTextOk returns a tuple with the ProfileText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileText

`func (o *PeopleSpecificPerson) SetProfileText(v string)`

SetProfileText sets ProfileText field to given value.

### HasProfileText

`func (o *PeopleSpecificPerson) HasProfileText() bool`

HasProfileText returns a boolean if a field has been set.

### GetProjects

`func (o *PeopleSpecificPerson) GetProjects() map[string]interface{}`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *PeopleSpecificPerson) GetProjectsOk() (*map[string]interface{}, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *PeopleSpecificPerson) SetProjects(v map[string]interface{})`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *PeopleSpecificPerson) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetSharedUserFilter

`func (o *PeopleSpecificPerson) GetSharedUserFilter() map[string]TwdatausersFilterResponse`

GetSharedUserFilter returns the SharedUserFilter field if non-nil, zero value otherwise.

### GetSharedUserFilterOk

`func (o *PeopleSpecificPerson) GetSharedUserFilterOk() (*map[string]TwdatausersFilterResponse, bool)`

GetSharedUserFilterOk returns a tuple with the SharedUserFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedUserFilter

`func (o *PeopleSpecificPerson) SetSharedUserFilter(v map[string]TwdatausersFilterResponse)`

SetSharedUserFilter sets SharedUserFilter field to given value.

### HasSharedUserFilter

`func (o *PeopleSpecificPerson) HasSharedUserFilter() bool`

HasSharedUserFilter returns a boolean if a field has been set.

### GetSiteOwner

`func (o *PeopleSpecificPerson) GetSiteOwner() bool`

GetSiteOwner returns the SiteOwner field if non-nil, zero value otherwise.

### GetSiteOwnerOk

`func (o *PeopleSpecificPerson) GetSiteOwnerOk() (*bool, bool)`

GetSiteOwnerOk returns a tuple with the SiteOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteOwner

`func (o *PeopleSpecificPerson) SetSiteOwner(v bool)`

SetSiteOwner sets SiteOwner field to given value.

### HasSiteOwner

`func (o *PeopleSpecificPerson) HasSiteOwner() bool`

HasSiteOwner returns a boolean if a field has been set.

### GetSocial

`func (o *PeopleSpecificPerson) GetSocial() PeopleSocial`

GetSocial returns the Social field if non-nil, zero value otherwise.

### GetSocialOk

`func (o *PeopleSpecificPerson) GetSocialOk() (*PeopleSocial, bool)`

GetSocialOk returns a tuple with the Social field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocial

`func (o *PeopleSpecificPerson) SetSocial(v PeopleSocial)`

SetSocial sets Social field to given value.

### HasSocial

`func (o *PeopleSpecificPerson) HasSocial() bool`

HasSocial returns a boolean if a field has been set.

### GetTags

`func (o *PeopleSpecificPerson) GetTags() []TagsTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PeopleSpecificPerson) GetTagsOk() (*[]TagsTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PeopleSpecificPerson) SetTags(v []TagsTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PeopleSpecificPerson) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTasks

`func (o *PeopleSpecificPerson) GetTasks() map[string]interface{}`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *PeopleSpecificPerson) GetTasksOk() (*map[string]interface{}, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *PeopleSpecificPerson) SetTasks(v map[string]interface{})`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *PeopleSpecificPerson) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetTeams

`func (o *PeopleSpecificPerson) GetTeams() []string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *PeopleSpecificPerson) GetTeamsOk() (*[]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *PeopleSpecificPerson) SetTeams(v []string)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *PeopleSpecificPerson) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetTeamsCount

`func (o *PeopleSpecificPerson) GetTeamsCount() int32`

GetTeamsCount returns the TeamsCount field if non-nil, zero value otherwise.

### GetTeamsCountOk

`func (o *PeopleSpecificPerson) GetTeamsCountOk() (*int32, bool)`

GetTeamsCountOk returns a tuple with the TeamsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsCount

`func (o *PeopleSpecificPerson) SetTeamsCount(v int32)`

SetTeamsCount sets TeamsCount field to given value.

### HasTeamsCount

`func (o *PeopleSpecificPerson) HasTeamsCount() bool`

HasTeamsCount returns a boolean if a field has been set.

### GetTextFormat

`func (o *PeopleSpecificPerson) GetTextFormat() string`

GetTextFormat returns the TextFormat field if non-nil, zero value otherwise.

### GetTextFormatOk

`func (o *PeopleSpecificPerson) GetTextFormatOk() (*string, bool)`

GetTextFormatOk returns a tuple with the TextFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextFormat

`func (o *PeopleSpecificPerson) SetTextFormat(v string)`

SetTextFormat sets TextFormat field to given value.

### HasTextFormat

`func (o *PeopleSpecificPerson) HasTextFormat() bool`

HasTextFormat returns a boolean if a field has been set.

### GetTitle

`func (o *PeopleSpecificPerson) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PeopleSpecificPerson) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PeopleSpecificPerson) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PeopleSpecificPerson) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTwitter

`func (o *PeopleSpecificPerson) GetTwitter() string`

GetTwitter returns the Twitter field if non-nil, zero value otherwise.

### GetTwitterOk

`func (o *PeopleSpecificPerson) GetTwitterOk() (*string, bool)`

GetTwitterOk returns a tuple with the Twitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitter

`func (o *PeopleSpecificPerson) SetTwitter(v string)`

SetTwitter sets Twitter field to given value.

### HasTwitter

`func (o *PeopleSpecificPerson) HasTwitter() bool`

HasTwitter returns a boolean if a field has been set.

### GetTwoFactorAuthEnabled

`func (o *PeopleSpecificPerson) GetTwoFactorAuthEnabled() bool`

GetTwoFactorAuthEnabled returns the TwoFactorAuthEnabled field if non-nil, zero value otherwise.

### GetTwoFactorAuthEnabledOk

`func (o *PeopleSpecificPerson) GetTwoFactorAuthEnabledOk() (*bool, bool)`

GetTwoFactorAuthEnabledOk returns a tuple with the TwoFactorAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorAuthEnabled

`func (o *PeopleSpecificPerson) SetTwoFactorAuthEnabled(v bool)`

SetTwoFactorAuthEnabled sets TwoFactorAuthEnabled field to given value.

### HasTwoFactorAuthEnabled

`func (o *PeopleSpecificPerson) HasTwoFactorAuthEnabled() bool`

HasTwoFactorAuthEnabled returns a boolean if a field has been set.

### GetUseShorthandDurations

`func (o *PeopleSpecificPerson) GetUseShorthandDurations() bool`

GetUseShorthandDurations returns the UseShorthandDurations field if non-nil, zero value otherwise.

### GetUseShorthandDurationsOk

`func (o *PeopleSpecificPerson) GetUseShorthandDurationsOk() (*bool, bool)`

GetUseShorthandDurationsOk returns a tuple with the UseShorthandDurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseShorthandDurations

`func (o *PeopleSpecificPerson) SetUseShorthandDurations(v bool)`

SetUseShorthandDurations sets UseShorthandDurations field to given value.

### HasUseShorthandDurations

`func (o *PeopleSpecificPerson) HasUseShorthandDurations() bool`

HasUseShorthandDurations returns a boolean if a field has been set.

### GetUserInvited

`func (o *PeopleSpecificPerson) GetUserInvited() int32`

GetUserInvited returns the UserInvited field if non-nil, zero value otherwise.

### GetUserInvitedOk

`func (o *PeopleSpecificPerson) GetUserInvitedOk() (*int32, bool)`

GetUserInvitedOk returns a tuple with the UserInvited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInvited

`func (o *PeopleSpecificPerson) SetUserInvited(v int32)`

SetUserInvited sets UserInvited field to given value.

### HasUserInvited

`func (o *PeopleSpecificPerson) HasUserInvited() bool`

HasUserInvited returns a boolean if a field has been set.

### GetUserInvitedDate

`func (o *PeopleSpecificPerson) GetUserInvitedDate() string`

GetUserInvitedDate returns the UserInvitedDate field if non-nil, zero value otherwise.

### GetUserInvitedDateOk

`func (o *PeopleSpecificPerson) GetUserInvitedDateOk() (*string, bool)`

GetUserInvitedDateOk returns a tuple with the UserInvitedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInvitedDate

`func (o *PeopleSpecificPerson) SetUserInvitedDate(v string)`

SetUserInvitedDate sets UserInvitedDate field to given value.

### HasUserInvitedDate

`func (o *PeopleSpecificPerson) HasUserInvitedDate() bool`

HasUserInvitedDate returns a boolean if a field has been set.

### GetUserInvitedStatus

`func (o *PeopleSpecificPerson) GetUserInvitedStatus() string`

GetUserInvitedStatus returns the UserInvitedStatus field if non-nil, zero value otherwise.

### GetUserInvitedStatusOk

`func (o *PeopleSpecificPerson) GetUserInvitedStatusOk() (*string, bool)`

GetUserInvitedStatusOk returns a tuple with the UserInvitedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInvitedStatus

`func (o *PeopleSpecificPerson) SetUserInvitedStatus(v string)`

SetUserInvitedStatus sets UserInvitedStatus field to given value.

### HasUserInvitedStatus

`func (o *PeopleSpecificPerson) HasUserInvitedStatus() bool`

HasUserInvitedStatus returns a boolean if a field has been set.

### GetUserName

`func (o *PeopleSpecificPerson) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *PeopleSpecificPerson) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *PeopleSpecificPerson) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *PeopleSpecificPerson) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetUserType

`func (o *PeopleSpecificPerson) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *PeopleSpecificPerson) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *PeopleSpecificPerson) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *PeopleSpecificPerson) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetUserUUID

`func (o *PeopleSpecificPerson) GetUserUUID() string`

GetUserUUID returns the UserUUID field if non-nil, zero value otherwise.

### GetUserUUIDOk

`func (o *PeopleSpecificPerson) GetUserUUIDOk() (*string, bool)`

GetUserUUIDOk returns a tuple with the UserUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUUID

`func (o *PeopleSpecificPerson) SetUserUUID(v string)`

SetUserUUID sets UserUUID field to given value.

### HasUserUUID

`func (o *PeopleSpecificPerson) HasUserUUID() bool`

HasUserUUID returns a boolean if a field has been set.

### GetWorkingHours

`func (o *PeopleSpecificPerson) GetWorkingHours() PeopleWorkingHours`

GetWorkingHours returns the WorkingHours field if non-nil, zero value otherwise.

### GetWorkingHoursOk

`func (o *PeopleSpecificPerson) GetWorkingHoursOk() (*PeopleWorkingHours, bool)`

GetWorkingHoursOk returns a tuple with the WorkingHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingHours

`func (o *PeopleSpecificPerson) SetWorkingHours(v PeopleWorkingHours)`

SetWorkingHours sets WorkingHours field to given value.

### HasWorkingHours

`func (o *PeopleSpecificPerson) HasWorkingHours() bool`

HasWorkingHours returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


