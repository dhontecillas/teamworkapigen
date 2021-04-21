# PeopleJsonPerson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**PeopleJsonPersonAddress**](PeopleJsonPersonAddress.md) |  | [optional] 
**Administrator** | Pointer to **bool** |  | [optional] 
**AllowEmailNotifications** | Pointer to **bool** |  | [optional] 
**AutoGiveProjectAccess** | Pointer to **bool** |  | [optional] 
**AvatarPendingFileRef** | Pointer to **string** |  | [optional] 
**CalendarStartsOnSunday** | Pointer to **string** |  | [optional] 
**CanAccessPortfolio** | Pointer to **bool** |  | [optional] 
**CanAccessTemplates** | Pointer to **bool** |  | [optional] 
**CanAddProjects** | Pointer to **bool** |  | [optional] 
**CanManagePeople** | Pointer to **bool** |  | [optional] 
**CanManagePortfolio** | Pointer to **bool** |  | [optional] 
**ChangeForEveryone** | Pointer to **bool** |  | [optional] 
**CompanyId** | Pointer to **string** |  | [optional] 
**DailyReportDaysFilter** | Pointer to **int32** |  | [optional] 
**DailyReportSort** | Pointer to **string** |  | [optional] 
**DateFormatId** | Pointer to **int32** |  | [optional] 
**EmailAddress** | Pointer to **string** |  | [optional] 
**EmailAlt1** | Pointer to **string** |  | [optional] 
**EmailAlt2** | Pointer to **string** |  | [optional] 
**EmailAlt3** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**GetUserDetails** | Pointer to **bool** |  | [optional] 
**ImHandle** | Pointer to **string** |  | [optional] 
**ImService** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**LengthOfDay** | Pointer to **int32** |  | [optional] 
**NotifyOnAddedAsFollower** | Pointer to **bool** |  | [optional] 
**NotifyOnStatusUpdate** | Pointer to **bool** |  | [optional] 
**NotifyOnTaskComplete** | Pointer to **bool** |  | [optional] 
**PhoneNumberFax** | Pointer to **string** |  | [optional] 
**PhoneNumberHome** | Pointer to **string** |  | [optional] 
**PhoneNumberMobile** | Pointer to **string** |  | [optional] 
**PhoneNumberMobileCountrycode** | Pointer to **string** |  | [optional] 
**PhoneNumberMobilePrefix** | Pointer to **string** |  | [optional] 
**PhoneNumberOffice** | Pointer to **string** |  | [optional] 
**PhoneNumberOfficeExt** | Pointer to **string** |  | [optional] 
**PrivateNotes** | Pointer to **string** |  | [optional] 
**Profile** | Pointer to **string** |  | [optional] 
**ReceiveDailyReports** | Pointer to **bool** |  | [optional] 
**ReceiveDailyReportsAtTime** | Pointer to **string** |  | [optional] 
**ReceiveDailyReportsAtWeekend** | Pointer to **bool** |  | [optional] 
**ReceiveDailyReportsIfEmpty** | Pointer to **bool** |  | [optional] 
**RemoveAvatar** | Pointer to **bool** |  | [optional] 
**SendInvite** | Pointer to **bool** |  | [optional] 
**SoundAlertsEnabled** | Pointer to **bool** |  | [optional] 
**TextFormat** | Pointer to **string** |  | [optional] 
**TimeFormatId** | Pointer to **int32** |  | [optional] 
**TimezoneId** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UseShorthandDurations** | Pointer to **bool** |  | [optional] 
**UserType** | Pointer to **string** |  | [optional] 
**UserFacebook** | Pointer to **string** |  | [optional] 
**UserGoogleplus** | Pointer to **string** |  | [optional] 
**UserLinkedin** | Pointer to **string** |  | [optional] 
**UserReceiveMyNotificationsOnly** | Pointer to **bool** |  | [optional] 
**UserReceiveNotifyWarnings** | Pointer to **bool** |  | [optional] 
**UserTwitterName** | Pointer to **string** |  | [optional] 
**UserWebsite** | Pointer to **string** |  | [optional] 

## Methods

### NewPeopleJsonPerson

`func NewPeopleJsonPerson() *PeopleJsonPerson`

NewPeopleJsonPerson instantiates a new PeopleJsonPerson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeopleJsonPersonWithDefaults

`func NewPeopleJsonPersonWithDefaults() *PeopleJsonPerson`

NewPeopleJsonPersonWithDefaults instantiates a new PeopleJsonPerson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *PeopleJsonPerson) GetAddress() PeopleJsonPersonAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PeopleJsonPerson) GetAddressOk() (*PeopleJsonPersonAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PeopleJsonPerson) SetAddress(v PeopleJsonPersonAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PeopleJsonPerson) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAdministrator

`func (o *PeopleJsonPerson) GetAdministrator() bool`

GetAdministrator returns the Administrator field if non-nil, zero value otherwise.

### GetAdministratorOk

`func (o *PeopleJsonPerson) GetAdministratorOk() (*bool, bool)`

GetAdministratorOk returns a tuple with the Administrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrator

`func (o *PeopleJsonPerson) SetAdministrator(v bool)`

SetAdministrator sets Administrator field to given value.

### HasAdministrator

`func (o *PeopleJsonPerson) HasAdministrator() bool`

HasAdministrator returns a boolean if a field has been set.

### GetAllowEmailNotifications

`func (o *PeopleJsonPerson) GetAllowEmailNotifications() bool`

GetAllowEmailNotifications returns the AllowEmailNotifications field if non-nil, zero value otherwise.

### GetAllowEmailNotificationsOk

`func (o *PeopleJsonPerson) GetAllowEmailNotificationsOk() (*bool, bool)`

GetAllowEmailNotificationsOk returns a tuple with the AllowEmailNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEmailNotifications

`func (o *PeopleJsonPerson) SetAllowEmailNotifications(v bool)`

SetAllowEmailNotifications sets AllowEmailNotifications field to given value.

### HasAllowEmailNotifications

`func (o *PeopleJsonPerson) HasAllowEmailNotifications() bool`

HasAllowEmailNotifications returns a boolean if a field has been set.

### GetAutoGiveProjectAccess

`func (o *PeopleJsonPerson) GetAutoGiveProjectAccess() bool`

GetAutoGiveProjectAccess returns the AutoGiveProjectAccess field if non-nil, zero value otherwise.

### GetAutoGiveProjectAccessOk

`func (o *PeopleJsonPerson) GetAutoGiveProjectAccessOk() (*bool, bool)`

GetAutoGiveProjectAccessOk returns a tuple with the AutoGiveProjectAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoGiveProjectAccess

`func (o *PeopleJsonPerson) SetAutoGiveProjectAccess(v bool)`

SetAutoGiveProjectAccess sets AutoGiveProjectAccess field to given value.

### HasAutoGiveProjectAccess

`func (o *PeopleJsonPerson) HasAutoGiveProjectAccess() bool`

HasAutoGiveProjectAccess returns a boolean if a field has been set.

### GetAvatarPendingFileRef

`func (o *PeopleJsonPerson) GetAvatarPendingFileRef() string`

GetAvatarPendingFileRef returns the AvatarPendingFileRef field if non-nil, zero value otherwise.

### GetAvatarPendingFileRefOk

`func (o *PeopleJsonPerson) GetAvatarPendingFileRefOk() (*string, bool)`

GetAvatarPendingFileRefOk returns a tuple with the AvatarPendingFileRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarPendingFileRef

`func (o *PeopleJsonPerson) SetAvatarPendingFileRef(v string)`

SetAvatarPendingFileRef sets AvatarPendingFileRef field to given value.

### HasAvatarPendingFileRef

`func (o *PeopleJsonPerson) HasAvatarPendingFileRef() bool`

HasAvatarPendingFileRef returns a boolean if a field has been set.

### GetCalendarStartsOnSunday

`func (o *PeopleJsonPerson) GetCalendarStartsOnSunday() string`

GetCalendarStartsOnSunday returns the CalendarStartsOnSunday field if non-nil, zero value otherwise.

### GetCalendarStartsOnSundayOk

`func (o *PeopleJsonPerson) GetCalendarStartsOnSundayOk() (*string, bool)`

GetCalendarStartsOnSundayOk returns a tuple with the CalendarStartsOnSunday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarStartsOnSunday

`func (o *PeopleJsonPerson) SetCalendarStartsOnSunday(v string)`

SetCalendarStartsOnSunday sets CalendarStartsOnSunday field to given value.

### HasCalendarStartsOnSunday

`func (o *PeopleJsonPerson) HasCalendarStartsOnSunday() bool`

HasCalendarStartsOnSunday returns a boolean if a field has been set.

### GetCanAccessPortfolio

`func (o *PeopleJsonPerson) GetCanAccessPortfolio() bool`

GetCanAccessPortfolio returns the CanAccessPortfolio field if non-nil, zero value otherwise.

### GetCanAccessPortfolioOk

`func (o *PeopleJsonPerson) GetCanAccessPortfolioOk() (*bool, bool)`

GetCanAccessPortfolioOk returns a tuple with the CanAccessPortfolio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAccessPortfolio

`func (o *PeopleJsonPerson) SetCanAccessPortfolio(v bool)`

SetCanAccessPortfolio sets CanAccessPortfolio field to given value.

### HasCanAccessPortfolio

`func (o *PeopleJsonPerson) HasCanAccessPortfolio() bool`

HasCanAccessPortfolio returns a boolean if a field has been set.

### GetCanAccessTemplates

`func (o *PeopleJsonPerson) GetCanAccessTemplates() bool`

GetCanAccessTemplates returns the CanAccessTemplates field if non-nil, zero value otherwise.

### GetCanAccessTemplatesOk

`func (o *PeopleJsonPerson) GetCanAccessTemplatesOk() (*bool, bool)`

GetCanAccessTemplatesOk returns a tuple with the CanAccessTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAccessTemplates

`func (o *PeopleJsonPerson) SetCanAccessTemplates(v bool)`

SetCanAccessTemplates sets CanAccessTemplates field to given value.

### HasCanAccessTemplates

`func (o *PeopleJsonPerson) HasCanAccessTemplates() bool`

HasCanAccessTemplates returns a boolean if a field has been set.

### GetCanAddProjects

`func (o *PeopleJsonPerson) GetCanAddProjects() bool`

GetCanAddProjects returns the CanAddProjects field if non-nil, zero value otherwise.

### GetCanAddProjectsOk

`func (o *PeopleJsonPerson) GetCanAddProjectsOk() (*bool, bool)`

GetCanAddProjectsOk returns a tuple with the CanAddProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAddProjects

`func (o *PeopleJsonPerson) SetCanAddProjects(v bool)`

SetCanAddProjects sets CanAddProjects field to given value.

### HasCanAddProjects

`func (o *PeopleJsonPerson) HasCanAddProjects() bool`

HasCanAddProjects returns a boolean if a field has been set.

### GetCanManagePeople

`func (o *PeopleJsonPerson) GetCanManagePeople() bool`

GetCanManagePeople returns the CanManagePeople field if non-nil, zero value otherwise.

### GetCanManagePeopleOk

`func (o *PeopleJsonPerson) GetCanManagePeopleOk() (*bool, bool)`

GetCanManagePeopleOk returns a tuple with the CanManagePeople field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManagePeople

`func (o *PeopleJsonPerson) SetCanManagePeople(v bool)`

SetCanManagePeople sets CanManagePeople field to given value.

### HasCanManagePeople

`func (o *PeopleJsonPerson) HasCanManagePeople() bool`

HasCanManagePeople returns a boolean if a field has been set.

### GetCanManagePortfolio

`func (o *PeopleJsonPerson) GetCanManagePortfolio() bool`

GetCanManagePortfolio returns the CanManagePortfolio field if non-nil, zero value otherwise.

### GetCanManagePortfolioOk

`func (o *PeopleJsonPerson) GetCanManagePortfolioOk() (*bool, bool)`

GetCanManagePortfolioOk returns a tuple with the CanManagePortfolio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManagePortfolio

`func (o *PeopleJsonPerson) SetCanManagePortfolio(v bool)`

SetCanManagePortfolio sets CanManagePortfolio field to given value.

### HasCanManagePortfolio

`func (o *PeopleJsonPerson) HasCanManagePortfolio() bool`

HasCanManagePortfolio returns a boolean if a field has been set.

### GetChangeForEveryone

`func (o *PeopleJsonPerson) GetChangeForEveryone() bool`

GetChangeForEveryone returns the ChangeForEveryone field if non-nil, zero value otherwise.

### GetChangeForEveryoneOk

`func (o *PeopleJsonPerson) GetChangeForEveryoneOk() (*bool, bool)`

GetChangeForEveryoneOk returns a tuple with the ChangeForEveryone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeForEveryone

`func (o *PeopleJsonPerson) SetChangeForEveryone(v bool)`

SetChangeForEveryone sets ChangeForEveryone field to given value.

### HasChangeForEveryone

`func (o *PeopleJsonPerson) HasChangeForEveryone() bool`

HasChangeForEveryone returns a boolean if a field has been set.

### GetCompanyId

`func (o *PeopleJsonPerson) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *PeopleJsonPerson) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *PeopleJsonPerson) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *PeopleJsonPerson) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetDailyReportDaysFilter

`func (o *PeopleJsonPerson) GetDailyReportDaysFilter() int32`

GetDailyReportDaysFilter returns the DailyReportDaysFilter field if non-nil, zero value otherwise.

### GetDailyReportDaysFilterOk

`func (o *PeopleJsonPerson) GetDailyReportDaysFilterOk() (*int32, bool)`

GetDailyReportDaysFilterOk returns a tuple with the DailyReportDaysFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyReportDaysFilter

`func (o *PeopleJsonPerson) SetDailyReportDaysFilter(v int32)`

SetDailyReportDaysFilter sets DailyReportDaysFilter field to given value.

### HasDailyReportDaysFilter

`func (o *PeopleJsonPerson) HasDailyReportDaysFilter() bool`

HasDailyReportDaysFilter returns a boolean if a field has been set.

### GetDailyReportSort

`func (o *PeopleJsonPerson) GetDailyReportSort() string`

GetDailyReportSort returns the DailyReportSort field if non-nil, zero value otherwise.

### GetDailyReportSortOk

`func (o *PeopleJsonPerson) GetDailyReportSortOk() (*string, bool)`

GetDailyReportSortOk returns a tuple with the DailyReportSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyReportSort

`func (o *PeopleJsonPerson) SetDailyReportSort(v string)`

SetDailyReportSort sets DailyReportSort field to given value.

### HasDailyReportSort

`func (o *PeopleJsonPerson) HasDailyReportSort() bool`

HasDailyReportSort returns a boolean if a field has been set.

### GetDateFormatId

`func (o *PeopleJsonPerson) GetDateFormatId() int32`

GetDateFormatId returns the DateFormatId field if non-nil, zero value otherwise.

### GetDateFormatIdOk

`func (o *PeopleJsonPerson) GetDateFormatIdOk() (*int32, bool)`

GetDateFormatIdOk returns a tuple with the DateFormatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFormatId

`func (o *PeopleJsonPerson) SetDateFormatId(v int32)`

SetDateFormatId sets DateFormatId field to given value.

### HasDateFormatId

`func (o *PeopleJsonPerson) HasDateFormatId() bool`

HasDateFormatId returns a boolean if a field has been set.

### GetEmailAddress

`func (o *PeopleJsonPerson) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *PeopleJsonPerson) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *PeopleJsonPerson) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *PeopleJsonPerson) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetEmailAlt1

`func (o *PeopleJsonPerson) GetEmailAlt1() string`

GetEmailAlt1 returns the EmailAlt1 field if non-nil, zero value otherwise.

### GetEmailAlt1Ok

`func (o *PeopleJsonPerson) GetEmailAlt1Ok() (*string, bool)`

GetEmailAlt1Ok returns a tuple with the EmailAlt1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAlt1

`func (o *PeopleJsonPerson) SetEmailAlt1(v string)`

SetEmailAlt1 sets EmailAlt1 field to given value.

### HasEmailAlt1

`func (o *PeopleJsonPerson) HasEmailAlt1() bool`

HasEmailAlt1 returns a boolean if a field has been set.

### GetEmailAlt2

`func (o *PeopleJsonPerson) GetEmailAlt2() string`

GetEmailAlt2 returns the EmailAlt2 field if non-nil, zero value otherwise.

### GetEmailAlt2Ok

`func (o *PeopleJsonPerson) GetEmailAlt2Ok() (*string, bool)`

GetEmailAlt2Ok returns a tuple with the EmailAlt2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAlt2

`func (o *PeopleJsonPerson) SetEmailAlt2(v string)`

SetEmailAlt2 sets EmailAlt2 field to given value.

### HasEmailAlt2

`func (o *PeopleJsonPerson) HasEmailAlt2() bool`

HasEmailAlt2 returns a boolean if a field has been set.

### GetEmailAlt3

`func (o *PeopleJsonPerson) GetEmailAlt3() string`

GetEmailAlt3 returns the EmailAlt3 field if non-nil, zero value otherwise.

### GetEmailAlt3Ok

`func (o *PeopleJsonPerson) GetEmailAlt3Ok() (*string, bool)`

GetEmailAlt3Ok returns a tuple with the EmailAlt3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAlt3

`func (o *PeopleJsonPerson) SetEmailAlt3(v string)`

SetEmailAlt3 sets EmailAlt3 field to given value.

### HasEmailAlt3

`func (o *PeopleJsonPerson) HasEmailAlt3() bool`

HasEmailAlt3 returns a boolean if a field has been set.

### GetFirstName

`func (o *PeopleJsonPerson) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PeopleJsonPerson) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PeopleJsonPerson) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PeopleJsonPerson) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetGetUserDetails

`func (o *PeopleJsonPerson) GetGetUserDetails() bool`

GetGetUserDetails returns the GetUserDetails field if non-nil, zero value otherwise.

### GetGetUserDetailsOk

`func (o *PeopleJsonPerson) GetGetUserDetailsOk() (*bool, bool)`

GetGetUserDetailsOk returns a tuple with the GetUserDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetUserDetails

`func (o *PeopleJsonPerson) SetGetUserDetails(v bool)`

SetGetUserDetails sets GetUserDetails field to given value.

### HasGetUserDetails

`func (o *PeopleJsonPerson) HasGetUserDetails() bool`

HasGetUserDetails returns a boolean if a field has been set.

### GetImHandle

`func (o *PeopleJsonPerson) GetImHandle() string`

GetImHandle returns the ImHandle field if non-nil, zero value otherwise.

### GetImHandleOk

`func (o *PeopleJsonPerson) GetImHandleOk() (*string, bool)`

GetImHandleOk returns a tuple with the ImHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImHandle

`func (o *PeopleJsonPerson) SetImHandle(v string)`

SetImHandle sets ImHandle field to given value.

### HasImHandle

`func (o *PeopleJsonPerson) HasImHandle() bool`

HasImHandle returns a boolean if a field has been set.

### GetImService

`func (o *PeopleJsonPerson) GetImService() string`

GetImService returns the ImService field if non-nil, zero value otherwise.

### GetImServiceOk

`func (o *PeopleJsonPerson) GetImServiceOk() (*string, bool)`

GetImServiceOk returns a tuple with the ImService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImService

`func (o *PeopleJsonPerson) SetImService(v string)`

SetImService sets ImService field to given value.

### HasImService

`func (o *PeopleJsonPerson) HasImService() bool`

HasImService returns a boolean if a field has been set.

### GetLanguage

`func (o *PeopleJsonPerson) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *PeopleJsonPerson) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *PeopleJsonPerson) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *PeopleJsonPerson) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLastName

`func (o *PeopleJsonPerson) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PeopleJsonPerson) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PeopleJsonPerson) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PeopleJsonPerson) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLengthOfDay

`func (o *PeopleJsonPerson) GetLengthOfDay() int32`

GetLengthOfDay returns the LengthOfDay field if non-nil, zero value otherwise.

### GetLengthOfDayOk

`func (o *PeopleJsonPerson) GetLengthOfDayOk() (*int32, bool)`

GetLengthOfDayOk returns a tuple with the LengthOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthOfDay

`func (o *PeopleJsonPerson) SetLengthOfDay(v int32)`

SetLengthOfDay sets LengthOfDay field to given value.

### HasLengthOfDay

`func (o *PeopleJsonPerson) HasLengthOfDay() bool`

HasLengthOfDay returns a boolean if a field has been set.

### GetNotifyOnAddedAsFollower

`func (o *PeopleJsonPerson) GetNotifyOnAddedAsFollower() bool`

GetNotifyOnAddedAsFollower returns the NotifyOnAddedAsFollower field if non-nil, zero value otherwise.

### GetNotifyOnAddedAsFollowerOk

`func (o *PeopleJsonPerson) GetNotifyOnAddedAsFollowerOk() (*bool, bool)`

GetNotifyOnAddedAsFollowerOk returns a tuple with the NotifyOnAddedAsFollower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyOnAddedAsFollower

`func (o *PeopleJsonPerson) SetNotifyOnAddedAsFollower(v bool)`

SetNotifyOnAddedAsFollower sets NotifyOnAddedAsFollower field to given value.

### HasNotifyOnAddedAsFollower

`func (o *PeopleJsonPerson) HasNotifyOnAddedAsFollower() bool`

HasNotifyOnAddedAsFollower returns a boolean if a field has been set.

### GetNotifyOnStatusUpdate

`func (o *PeopleJsonPerson) GetNotifyOnStatusUpdate() bool`

GetNotifyOnStatusUpdate returns the NotifyOnStatusUpdate field if non-nil, zero value otherwise.

### GetNotifyOnStatusUpdateOk

`func (o *PeopleJsonPerson) GetNotifyOnStatusUpdateOk() (*bool, bool)`

GetNotifyOnStatusUpdateOk returns a tuple with the NotifyOnStatusUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyOnStatusUpdate

`func (o *PeopleJsonPerson) SetNotifyOnStatusUpdate(v bool)`

SetNotifyOnStatusUpdate sets NotifyOnStatusUpdate field to given value.

### HasNotifyOnStatusUpdate

`func (o *PeopleJsonPerson) HasNotifyOnStatusUpdate() bool`

HasNotifyOnStatusUpdate returns a boolean if a field has been set.

### GetNotifyOnTaskComplete

`func (o *PeopleJsonPerson) GetNotifyOnTaskComplete() bool`

GetNotifyOnTaskComplete returns the NotifyOnTaskComplete field if non-nil, zero value otherwise.

### GetNotifyOnTaskCompleteOk

`func (o *PeopleJsonPerson) GetNotifyOnTaskCompleteOk() (*bool, bool)`

GetNotifyOnTaskCompleteOk returns a tuple with the NotifyOnTaskComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyOnTaskComplete

`func (o *PeopleJsonPerson) SetNotifyOnTaskComplete(v bool)`

SetNotifyOnTaskComplete sets NotifyOnTaskComplete field to given value.

### HasNotifyOnTaskComplete

`func (o *PeopleJsonPerson) HasNotifyOnTaskComplete() bool`

HasNotifyOnTaskComplete returns a boolean if a field has been set.

### GetPhoneNumberFax

`func (o *PeopleJsonPerson) GetPhoneNumberFax() string`

GetPhoneNumberFax returns the PhoneNumberFax field if non-nil, zero value otherwise.

### GetPhoneNumberFaxOk

`func (o *PeopleJsonPerson) GetPhoneNumberFaxOk() (*string, bool)`

GetPhoneNumberFaxOk returns a tuple with the PhoneNumberFax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberFax

`func (o *PeopleJsonPerson) SetPhoneNumberFax(v string)`

SetPhoneNumberFax sets PhoneNumberFax field to given value.

### HasPhoneNumberFax

`func (o *PeopleJsonPerson) HasPhoneNumberFax() bool`

HasPhoneNumberFax returns a boolean if a field has been set.

### GetPhoneNumberHome

`func (o *PeopleJsonPerson) GetPhoneNumberHome() string`

GetPhoneNumberHome returns the PhoneNumberHome field if non-nil, zero value otherwise.

### GetPhoneNumberHomeOk

`func (o *PeopleJsonPerson) GetPhoneNumberHomeOk() (*string, bool)`

GetPhoneNumberHomeOk returns a tuple with the PhoneNumberHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberHome

`func (o *PeopleJsonPerson) SetPhoneNumberHome(v string)`

SetPhoneNumberHome sets PhoneNumberHome field to given value.

### HasPhoneNumberHome

`func (o *PeopleJsonPerson) HasPhoneNumberHome() bool`

HasPhoneNumberHome returns a boolean if a field has been set.

### GetPhoneNumberMobile

`func (o *PeopleJsonPerson) GetPhoneNumberMobile() string`

GetPhoneNumberMobile returns the PhoneNumberMobile field if non-nil, zero value otherwise.

### GetPhoneNumberMobileOk

`func (o *PeopleJsonPerson) GetPhoneNumberMobileOk() (*string, bool)`

GetPhoneNumberMobileOk returns a tuple with the PhoneNumberMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberMobile

`func (o *PeopleJsonPerson) SetPhoneNumberMobile(v string)`

SetPhoneNumberMobile sets PhoneNumberMobile field to given value.

### HasPhoneNumberMobile

`func (o *PeopleJsonPerson) HasPhoneNumberMobile() bool`

HasPhoneNumberMobile returns a boolean if a field has been set.

### GetPhoneNumberMobileCountrycode

`func (o *PeopleJsonPerson) GetPhoneNumberMobileCountrycode() string`

GetPhoneNumberMobileCountrycode returns the PhoneNumberMobileCountrycode field if non-nil, zero value otherwise.

### GetPhoneNumberMobileCountrycodeOk

`func (o *PeopleJsonPerson) GetPhoneNumberMobileCountrycodeOk() (*string, bool)`

GetPhoneNumberMobileCountrycodeOk returns a tuple with the PhoneNumberMobileCountrycode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberMobileCountrycode

`func (o *PeopleJsonPerson) SetPhoneNumberMobileCountrycode(v string)`

SetPhoneNumberMobileCountrycode sets PhoneNumberMobileCountrycode field to given value.

### HasPhoneNumberMobileCountrycode

`func (o *PeopleJsonPerson) HasPhoneNumberMobileCountrycode() bool`

HasPhoneNumberMobileCountrycode returns a boolean if a field has been set.

### GetPhoneNumberMobilePrefix

`func (o *PeopleJsonPerson) GetPhoneNumberMobilePrefix() string`

GetPhoneNumberMobilePrefix returns the PhoneNumberMobilePrefix field if non-nil, zero value otherwise.

### GetPhoneNumberMobilePrefixOk

`func (o *PeopleJsonPerson) GetPhoneNumberMobilePrefixOk() (*string, bool)`

GetPhoneNumberMobilePrefixOk returns a tuple with the PhoneNumberMobilePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberMobilePrefix

`func (o *PeopleJsonPerson) SetPhoneNumberMobilePrefix(v string)`

SetPhoneNumberMobilePrefix sets PhoneNumberMobilePrefix field to given value.

### HasPhoneNumberMobilePrefix

`func (o *PeopleJsonPerson) HasPhoneNumberMobilePrefix() bool`

HasPhoneNumberMobilePrefix returns a boolean if a field has been set.

### GetPhoneNumberOffice

`func (o *PeopleJsonPerson) GetPhoneNumberOffice() string`

GetPhoneNumberOffice returns the PhoneNumberOffice field if non-nil, zero value otherwise.

### GetPhoneNumberOfficeOk

`func (o *PeopleJsonPerson) GetPhoneNumberOfficeOk() (*string, bool)`

GetPhoneNumberOfficeOk returns a tuple with the PhoneNumberOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberOffice

`func (o *PeopleJsonPerson) SetPhoneNumberOffice(v string)`

SetPhoneNumberOffice sets PhoneNumberOffice field to given value.

### HasPhoneNumberOffice

`func (o *PeopleJsonPerson) HasPhoneNumberOffice() bool`

HasPhoneNumberOffice returns a boolean if a field has been set.

### GetPhoneNumberOfficeExt

`func (o *PeopleJsonPerson) GetPhoneNumberOfficeExt() string`

GetPhoneNumberOfficeExt returns the PhoneNumberOfficeExt field if non-nil, zero value otherwise.

### GetPhoneNumberOfficeExtOk

`func (o *PeopleJsonPerson) GetPhoneNumberOfficeExtOk() (*string, bool)`

GetPhoneNumberOfficeExtOk returns a tuple with the PhoneNumberOfficeExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberOfficeExt

`func (o *PeopleJsonPerson) SetPhoneNumberOfficeExt(v string)`

SetPhoneNumberOfficeExt sets PhoneNumberOfficeExt field to given value.

### HasPhoneNumberOfficeExt

`func (o *PeopleJsonPerson) HasPhoneNumberOfficeExt() bool`

HasPhoneNumberOfficeExt returns a boolean if a field has been set.

### GetPrivateNotes

`func (o *PeopleJsonPerson) GetPrivateNotes() string`

GetPrivateNotes returns the PrivateNotes field if non-nil, zero value otherwise.

### GetPrivateNotesOk

`func (o *PeopleJsonPerson) GetPrivateNotesOk() (*string, bool)`

GetPrivateNotesOk returns a tuple with the PrivateNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNotes

`func (o *PeopleJsonPerson) SetPrivateNotes(v string)`

SetPrivateNotes sets PrivateNotes field to given value.

### HasPrivateNotes

`func (o *PeopleJsonPerson) HasPrivateNotes() bool`

HasPrivateNotes returns a boolean if a field has been set.

### GetProfile

`func (o *PeopleJsonPerson) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *PeopleJsonPerson) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *PeopleJsonPerson) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *PeopleJsonPerson) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetReceiveDailyReports

`func (o *PeopleJsonPerson) GetReceiveDailyReports() bool`

GetReceiveDailyReports returns the ReceiveDailyReports field if non-nil, zero value otherwise.

### GetReceiveDailyReportsOk

`func (o *PeopleJsonPerson) GetReceiveDailyReportsOk() (*bool, bool)`

GetReceiveDailyReportsOk returns a tuple with the ReceiveDailyReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveDailyReports

`func (o *PeopleJsonPerson) SetReceiveDailyReports(v bool)`

SetReceiveDailyReports sets ReceiveDailyReports field to given value.

### HasReceiveDailyReports

`func (o *PeopleJsonPerson) HasReceiveDailyReports() bool`

HasReceiveDailyReports returns a boolean if a field has been set.

### GetReceiveDailyReportsAtTime

`func (o *PeopleJsonPerson) GetReceiveDailyReportsAtTime() string`

GetReceiveDailyReportsAtTime returns the ReceiveDailyReportsAtTime field if non-nil, zero value otherwise.

### GetReceiveDailyReportsAtTimeOk

`func (o *PeopleJsonPerson) GetReceiveDailyReportsAtTimeOk() (*string, bool)`

GetReceiveDailyReportsAtTimeOk returns a tuple with the ReceiveDailyReportsAtTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveDailyReportsAtTime

`func (o *PeopleJsonPerson) SetReceiveDailyReportsAtTime(v string)`

SetReceiveDailyReportsAtTime sets ReceiveDailyReportsAtTime field to given value.

### HasReceiveDailyReportsAtTime

`func (o *PeopleJsonPerson) HasReceiveDailyReportsAtTime() bool`

HasReceiveDailyReportsAtTime returns a boolean if a field has been set.

### GetReceiveDailyReportsAtWeekend

`func (o *PeopleJsonPerson) GetReceiveDailyReportsAtWeekend() bool`

GetReceiveDailyReportsAtWeekend returns the ReceiveDailyReportsAtWeekend field if non-nil, zero value otherwise.

### GetReceiveDailyReportsAtWeekendOk

`func (o *PeopleJsonPerson) GetReceiveDailyReportsAtWeekendOk() (*bool, bool)`

GetReceiveDailyReportsAtWeekendOk returns a tuple with the ReceiveDailyReportsAtWeekend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveDailyReportsAtWeekend

`func (o *PeopleJsonPerson) SetReceiveDailyReportsAtWeekend(v bool)`

SetReceiveDailyReportsAtWeekend sets ReceiveDailyReportsAtWeekend field to given value.

### HasReceiveDailyReportsAtWeekend

`func (o *PeopleJsonPerson) HasReceiveDailyReportsAtWeekend() bool`

HasReceiveDailyReportsAtWeekend returns a boolean if a field has been set.

### GetReceiveDailyReportsIfEmpty

`func (o *PeopleJsonPerson) GetReceiveDailyReportsIfEmpty() bool`

GetReceiveDailyReportsIfEmpty returns the ReceiveDailyReportsIfEmpty field if non-nil, zero value otherwise.

### GetReceiveDailyReportsIfEmptyOk

`func (o *PeopleJsonPerson) GetReceiveDailyReportsIfEmptyOk() (*bool, bool)`

GetReceiveDailyReportsIfEmptyOk returns a tuple with the ReceiveDailyReportsIfEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveDailyReportsIfEmpty

`func (o *PeopleJsonPerson) SetReceiveDailyReportsIfEmpty(v bool)`

SetReceiveDailyReportsIfEmpty sets ReceiveDailyReportsIfEmpty field to given value.

### HasReceiveDailyReportsIfEmpty

`func (o *PeopleJsonPerson) HasReceiveDailyReportsIfEmpty() bool`

HasReceiveDailyReportsIfEmpty returns a boolean if a field has been set.

### GetRemoveAvatar

`func (o *PeopleJsonPerson) GetRemoveAvatar() bool`

GetRemoveAvatar returns the RemoveAvatar field if non-nil, zero value otherwise.

### GetRemoveAvatarOk

`func (o *PeopleJsonPerson) GetRemoveAvatarOk() (*bool, bool)`

GetRemoveAvatarOk returns a tuple with the RemoveAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAvatar

`func (o *PeopleJsonPerson) SetRemoveAvatar(v bool)`

SetRemoveAvatar sets RemoveAvatar field to given value.

### HasRemoveAvatar

`func (o *PeopleJsonPerson) HasRemoveAvatar() bool`

HasRemoveAvatar returns a boolean if a field has been set.

### GetSendInvite

`func (o *PeopleJsonPerson) GetSendInvite() bool`

GetSendInvite returns the SendInvite field if non-nil, zero value otherwise.

### GetSendInviteOk

`func (o *PeopleJsonPerson) GetSendInviteOk() (*bool, bool)`

GetSendInviteOk returns a tuple with the SendInvite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendInvite

`func (o *PeopleJsonPerson) SetSendInvite(v bool)`

SetSendInvite sets SendInvite field to given value.

### HasSendInvite

`func (o *PeopleJsonPerson) HasSendInvite() bool`

HasSendInvite returns a boolean if a field has been set.

### GetSoundAlertsEnabled

`func (o *PeopleJsonPerson) GetSoundAlertsEnabled() bool`

GetSoundAlertsEnabled returns the SoundAlertsEnabled field if non-nil, zero value otherwise.

### GetSoundAlertsEnabledOk

`func (o *PeopleJsonPerson) GetSoundAlertsEnabledOk() (*bool, bool)`

GetSoundAlertsEnabledOk returns a tuple with the SoundAlertsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoundAlertsEnabled

`func (o *PeopleJsonPerson) SetSoundAlertsEnabled(v bool)`

SetSoundAlertsEnabled sets SoundAlertsEnabled field to given value.

### HasSoundAlertsEnabled

`func (o *PeopleJsonPerson) HasSoundAlertsEnabled() bool`

HasSoundAlertsEnabled returns a boolean if a field has been set.

### GetTextFormat

`func (o *PeopleJsonPerson) GetTextFormat() string`

GetTextFormat returns the TextFormat field if non-nil, zero value otherwise.

### GetTextFormatOk

`func (o *PeopleJsonPerson) GetTextFormatOk() (*string, bool)`

GetTextFormatOk returns a tuple with the TextFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextFormat

`func (o *PeopleJsonPerson) SetTextFormat(v string)`

SetTextFormat sets TextFormat field to given value.

### HasTextFormat

`func (o *PeopleJsonPerson) HasTextFormat() bool`

HasTextFormat returns a boolean if a field has been set.

### GetTimeFormatId

`func (o *PeopleJsonPerson) GetTimeFormatId() int32`

GetTimeFormatId returns the TimeFormatId field if non-nil, zero value otherwise.

### GetTimeFormatIdOk

`func (o *PeopleJsonPerson) GetTimeFormatIdOk() (*int32, bool)`

GetTimeFormatIdOk returns a tuple with the TimeFormatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFormatId

`func (o *PeopleJsonPerson) SetTimeFormatId(v int32)`

SetTimeFormatId sets TimeFormatId field to given value.

### HasTimeFormatId

`func (o *PeopleJsonPerson) HasTimeFormatId() bool`

HasTimeFormatId returns a boolean if a field has been set.

### GetTimezoneId

`func (o *PeopleJsonPerson) GetTimezoneId() int32`

GetTimezoneId returns the TimezoneId field if non-nil, zero value otherwise.

### GetTimezoneIdOk

`func (o *PeopleJsonPerson) GetTimezoneIdOk() (*int32, bool)`

GetTimezoneIdOk returns a tuple with the TimezoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneId

`func (o *PeopleJsonPerson) SetTimezoneId(v int32)`

SetTimezoneId sets TimezoneId field to given value.

### HasTimezoneId

`func (o *PeopleJsonPerson) HasTimezoneId() bool`

HasTimezoneId returns a boolean if a field has been set.

### GetTitle

`func (o *PeopleJsonPerson) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PeopleJsonPerson) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PeopleJsonPerson) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PeopleJsonPerson) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUseShorthandDurations

`func (o *PeopleJsonPerson) GetUseShorthandDurations() bool`

GetUseShorthandDurations returns the UseShorthandDurations field if non-nil, zero value otherwise.

### GetUseShorthandDurationsOk

`func (o *PeopleJsonPerson) GetUseShorthandDurationsOk() (*bool, bool)`

GetUseShorthandDurationsOk returns a tuple with the UseShorthandDurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseShorthandDurations

`func (o *PeopleJsonPerson) SetUseShorthandDurations(v bool)`

SetUseShorthandDurations sets UseShorthandDurations field to given value.

### HasUseShorthandDurations

`func (o *PeopleJsonPerson) HasUseShorthandDurations() bool`

HasUseShorthandDurations returns a boolean if a field has been set.

### GetUserType

`func (o *PeopleJsonPerson) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *PeopleJsonPerson) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *PeopleJsonPerson) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *PeopleJsonPerson) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetUserFacebook

`func (o *PeopleJsonPerson) GetUserFacebook() string`

GetUserFacebook returns the UserFacebook field if non-nil, zero value otherwise.

### GetUserFacebookOk

`func (o *PeopleJsonPerson) GetUserFacebookOk() (*string, bool)`

GetUserFacebookOk returns a tuple with the UserFacebook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFacebook

`func (o *PeopleJsonPerson) SetUserFacebook(v string)`

SetUserFacebook sets UserFacebook field to given value.

### HasUserFacebook

`func (o *PeopleJsonPerson) HasUserFacebook() bool`

HasUserFacebook returns a boolean if a field has been set.

### GetUserGoogleplus

`func (o *PeopleJsonPerson) GetUserGoogleplus() string`

GetUserGoogleplus returns the UserGoogleplus field if non-nil, zero value otherwise.

### GetUserGoogleplusOk

`func (o *PeopleJsonPerson) GetUserGoogleplusOk() (*string, bool)`

GetUserGoogleplusOk returns a tuple with the UserGoogleplus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGoogleplus

`func (o *PeopleJsonPerson) SetUserGoogleplus(v string)`

SetUserGoogleplus sets UserGoogleplus field to given value.

### HasUserGoogleplus

`func (o *PeopleJsonPerson) HasUserGoogleplus() bool`

HasUserGoogleplus returns a boolean if a field has been set.

### GetUserLinkedin

`func (o *PeopleJsonPerson) GetUserLinkedin() string`

GetUserLinkedin returns the UserLinkedin field if non-nil, zero value otherwise.

### GetUserLinkedinOk

`func (o *PeopleJsonPerson) GetUserLinkedinOk() (*string, bool)`

GetUserLinkedinOk returns a tuple with the UserLinkedin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLinkedin

`func (o *PeopleJsonPerson) SetUserLinkedin(v string)`

SetUserLinkedin sets UserLinkedin field to given value.

### HasUserLinkedin

`func (o *PeopleJsonPerson) HasUserLinkedin() bool`

HasUserLinkedin returns a boolean if a field has been set.

### GetUserReceiveMyNotificationsOnly

`func (o *PeopleJsonPerson) GetUserReceiveMyNotificationsOnly() bool`

GetUserReceiveMyNotificationsOnly returns the UserReceiveMyNotificationsOnly field if non-nil, zero value otherwise.

### GetUserReceiveMyNotificationsOnlyOk

`func (o *PeopleJsonPerson) GetUserReceiveMyNotificationsOnlyOk() (*bool, bool)`

GetUserReceiveMyNotificationsOnlyOk returns a tuple with the UserReceiveMyNotificationsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserReceiveMyNotificationsOnly

`func (o *PeopleJsonPerson) SetUserReceiveMyNotificationsOnly(v bool)`

SetUserReceiveMyNotificationsOnly sets UserReceiveMyNotificationsOnly field to given value.

### HasUserReceiveMyNotificationsOnly

`func (o *PeopleJsonPerson) HasUserReceiveMyNotificationsOnly() bool`

HasUserReceiveMyNotificationsOnly returns a boolean if a field has been set.

### GetUserReceiveNotifyWarnings

`func (o *PeopleJsonPerson) GetUserReceiveNotifyWarnings() bool`

GetUserReceiveNotifyWarnings returns the UserReceiveNotifyWarnings field if non-nil, zero value otherwise.

### GetUserReceiveNotifyWarningsOk

`func (o *PeopleJsonPerson) GetUserReceiveNotifyWarningsOk() (*bool, bool)`

GetUserReceiveNotifyWarningsOk returns a tuple with the UserReceiveNotifyWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserReceiveNotifyWarnings

`func (o *PeopleJsonPerson) SetUserReceiveNotifyWarnings(v bool)`

SetUserReceiveNotifyWarnings sets UserReceiveNotifyWarnings field to given value.

### HasUserReceiveNotifyWarnings

`func (o *PeopleJsonPerson) HasUserReceiveNotifyWarnings() bool`

HasUserReceiveNotifyWarnings returns a boolean if a field has been set.

### GetUserTwitterName

`func (o *PeopleJsonPerson) GetUserTwitterName() string`

GetUserTwitterName returns the UserTwitterName field if non-nil, zero value otherwise.

### GetUserTwitterNameOk

`func (o *PeopleJsonPerson) GetUserTwitterNameOk() (*string, bool)`

GetUserTwitterNameOk returns a tuple with the UserTwitterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTwitterName

`func (o *PeopleJsonPerson) SetUserTwitterName(v string)`

SetUserTwitterName sets UserTwitterName field to given value.

### HasUserTwitterName

`func (o *PeopleJsonPerson) HasUserTwitterName() bool`

HasUserTwitterName returns a boolean if a field has been set.

### GetUserWebsite

`func (o *PeopleJsonPerson) GetUserWebsite() string`

GetUserWebsite returns the UserWebsite field if non-nil, zero value otherwise.

### GetUserWebsiteOk

`func (o *PeopleJsonPerson) GetUserWebsiteOk() (*string, bool)`

GetUserWebsiteOk returns a tuple with the UserWebsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserWebsite

`func (o *PeopleJsonPerson) SetUserWebsite(v string)`

SetUserWebsite sets UserWebsite field to given value.

### HasUserWebsite

`func (o *PeopleJsonPerson) HasUserWebsite() bool`

HasUserWebsite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


