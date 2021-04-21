# InlineResponse20054Projects

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivePages** | Pointer to [**InlineResponse20054ActivePages**](InlineResponse20054ActivePages.md) |  | [optional] 
**ActiveUserIsProjectAdmin** | Pointer to **bool** |  | [optional] 
**Announcement** | Pointer to **string** |  | [optional] 
**AnnouncementHTML** | Pointer to **string** |  | [optional] 
**Category** | Pointer to [**InlineResponse20054Category**](InlineResponse20054Category.md) |  | [optional] 
**Company** | Pointer to [**InlineResponse20054Company**](InlineResponse20054Company.md) |  | [optional] 
**CreatedOn** | Pointer to **string** |  | [optional] 
**DefaultPrivacy** | Pointer to **string** |  | [optional] 
**Defaults** | Pointer to [**InlineResponse20053Defaults**](InlineResponse20053Defaults.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DirectFileUploadsEnabled** | Pointer to **bool** |  | [optional] 
**FilesAutoNewVersion** | Pointer to **bool** |  | [optional] 
**HarvestTimersEnabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Integrations** | Pointer to [**InlineResponse20054Integrations**](InlineResponse20054Integrations.md) |  | [optional] 
**IsProjectAdmin** | Pointer to **bool** |  | [optional] 
**LastChangedOn** | Pointer to **string** |  | [optional] 
**LastWorkedOn** | Pointer to **string** |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**LogoFromCompany** | Pointer to **bool** |  | [optional] 
**MinMaxAvailableDates** | Pointer to [**InlineResponse20054MinMaxAvailableDates**](InlineResponse20054MinMaxAvailableDates.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NotifyEveryone** | Pointer to **bool** |  | [optional] 
**OverviewStartPage** | Pointer to **string** |  | [optional] 
**PrivacyEnabled** | Pointer to **bool** |  | [optional] 
**ProjectOwnerId** | Pointer to **int32** |  | [optional] 
**ReplyByEmailEnabled** | Pointer to **bool** |  | [optional] 
**ShowAnnouncement** | Pointer to **bool** |  | [optional] 
**SkipWeekends** | Pointer to **bool** |  | [optional] 
**Starred** | Pointer to **bool** |  | [optional] 
**StartPage** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SubStatus** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]InlineResponse20054Tags**](InlineResponse20054Tags.md) |  | [optional] 
**TasksStartPage** | Pointer to **string** |  | [optional] 
**TemplateDateTargetDefault** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewInlineResponse20054Projects

`func NewInlineResponse20054Projects() *InlineResponse20054Projects`

NewInlineResponse20054Projects instantiates a new InlineResponse20054Projects object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20054ProjectsWithDefaults

`func NewInlineResponse20054ProjectsWithDefaults() *InlineResponse20054Projects`

NewInlineResponse20054ProjectsWithDefaults instantiates a new InlineResponse20054Projects object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivePages

`func (o *InlineResponse20054Projects) GetActivePages() InlineResponse20054ActivePages`

GetActivePages returns the ActivePages field if non-nil, zero value otherwise.

### GetActivePagesOk

`func (o *InlineResponse20054Projects) GetActivePagesOk() (*InlineResponse20054ActivePages, bool)`

GetActivePagesOk returns a tuple with the ActivePages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePages

`func (o *InlineResponse20054Projects) SetActivePages(v InlineResponse20054ActivePages)`

SetActivePages sets ActivePages field to given value.

### HasActivePages

`func (o *InlineResponse20054Projects) HasActivePages() bool`

HasActivePages returns a boolean if a field has been set.

### GetActiveUserIsProjectAdmin

`func (o *InlineResponse20054Projects) GetActiveUserIsProjectAdmin() bool`

GetActiveUserIsProjectAdmin returns the ActiveUserIsProjectAdmin field if non-nil, zero value otherwise.

### GetActiveUserIsProjectAdminOk

`func (o *InlineResponse20054Projects) GetActiveUserIsProjectAdminOk() (*bool, bool)`

GetActiveUserIsProjectAdminOk returns a tuple with the ActiveUserIsProjectAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveUserIsProjectAdmin

`func (o *InlineResponse20054Projects) SetActiveUserIsProjectAdmin(v bool)`

SetActiveUserIsProjectAdmin sets ActiveUserIsProjectAdmin field to given value.

### HasActiveUserIsProjectAdmin

`func (o *InlineResponse20054Projects) HasActiveUserIsProjectAdmin() bool`

HasActiveUserIsProjectAdmin returns a boolean if a field has been set.

### GetAnnouncement

`func (o *InlineResponse20054Projects) GetAnnouncement() string`

GetAnnouncement returns the Announcement field if non-nil, zero value otherwise.

### GetAnnouncementOk

`func (o *InlineResponse20054Projects) GetAnnouncementOk() (*string, bool)`

GetAnnouncementOk returns a tuple with the Announcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncement

`func (o *InlineResponse20054Projects) SetAnnouncement(v string)`

SetAnnouncement sets Announcement field to given value.

### HasAnnouncement

`func (o *InlineResponse20054Projects) HasAnnouncement() bool`

HasAnnouncement returns a boolean if a field has been set.

### GetAnnouncementHTML

`func (o *InlineResponse20054Projects) GetAnnouncementHTML() string`

GetAnnouncementHTML returns the AnnouncementHTML field if non-nil, zero value otherwise.

### GetAnnouncementHTMLOk

`func (o *InlineResponse20054Projects) GetAnnouncementHTMLOk() (*string, bool)`

GetAnnouncementHTMLOk returns a tuple with the AnnouncementHTML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementHTML

`func (o *InlineResponse20054Projects) SetAnnouncementHTML(v string)`

SetAnnouncementHTML sets AnnouncementHTML field to given value.

### HasAnnouncementHTML

`func (o *InlineResponse20054Projects) HasAnnouncementHTML() bool`

HasAnnouncementHTML returns a boolean if a field has been set.

### GetCategory

`func (o *InlineResponse20054Projects) GetCategory() InlineResponse20054Category`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InlineResponse20054Projects) GetCategoryOk() (*InlineResponse20054Category, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InlineResponse20054Projects) SetCategory(v InlineResponse20054Category)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *InlineResponse20054Projects) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCompany

`func (o *InlineResponse20054Projects) GetCompany() InlineResponse20054Company`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *InlineResponse20054Projects) GetCompanyOk() (*InlineResponse20054Company, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *InlineResponse20054Projects) SetCompany(v InlineResponse20054Company)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *InlineResponse20054Projects) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCreatedOn

`func (o *InlineResponse20054Projects) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *InlineResponse20054Projects) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *InlineResponse20054Projects) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *InlineResponse20054Projects) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetDefaultPrivacy

`func (o *InlineResponse20054Projects) GetDefaultPrivacy() string`

GetDefaultPrivacy returns the DefaultPrivacy field if non-nil, zero value otherwise.

### GetDefaultPrivacyOk

`func (o *InlineResponse20054Projects) GetDefaultPrivacyOk() (*string, bool)`

GetDefaultPrivacyOk returns a tuple with the DefaultPrivacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPrivacy

`func (o *InlineResponse20054Projects) SetDefaultPrivacy(v string)`

SetDefaultPrivacy sets DefaultPrivacy field to given value.

### HasDefaultPrivacy

`func (o *InlineResponse20054Projects) HasDefaultPrivacy() bool`

HasDefaultPrivacy returns a boolean if a field has been set.

### GetDefaults

`func (o *InlineResponse20054Projects) GetDefaults() InlineResponse20053Defaults`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *InlineResponse20054Projects) GetDefaultsOk() (*InlineResponse20053Defaults, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *InlineResponse20054Projects) SetDefaults(v InlineResponse20053Defaults)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *InlineResponse20054Projects) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse20054Projects) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse20054Projects) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse20054Projects) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse20054Projects) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDirectFileUploadsEnabled

`func (o *InlineResponse20054Projects) GetDirectFileUploadsEnabled() bool`

GetDirectFileUploadsEnabled returns the DirectFileUploadsEnabled field if non-nil, zero value otherwise.

### GetDirectFileUploadsEnabledOk

`func (o *InlineResponse20054Projects) GetDirectFileUploadsEnabledOk() (*bool, bool)`

GetDirectFileUploadsEnabledOk returns a tuple with the DirectFileUploadsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectFileUploadsEnabled

`func (o *InlineResponse20054Projects) SetDirectFileUploadsEnabled(v bool)`

SetDirectFileUploadsEnabled sets DirectFileUploadsEnabled field to given value.

### HasDirectFileUploadsEnabled

`func (o *InlineResponse20054Projects) HasDirectFileUploadsEnabled() bool`

HasDirectFileUploadsEnabled returns a boolean if a field has been set.

### GetFilesAutoNewVersion

`func (o *InlineResponse20054Projects) GetFilesAutoNewVersion() bool`

GetFilesAutoNewVersion returns the FilesAutoNewVersion field if non-nil, zero value otherwise.

### GetFilesAutoNewVersionOk

`func (o *InlineResponse20054Projects) GetFilesAutoNewVersionOk() (*bool, bool)`

GetFilesAutoNewVersionOk returns a tuple with the FilesAutoNewVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesAutoNewVersion

`func (o *InlineResponse20054Projects) SetFilesAutoNewVersion(v bool)`

SetFilesAutoNewVersion sets FilesAutoNewVersion field to given value.

### HasFilesAutoNewVersion

`func (o *InlineResponse20054Projects) HasFilesAutoNewVersion() bool`

HasFilesAutoNewVersion returns a boolean if a field has been set.

### GetHarvestTimersEnabled

`func (o *InlineResponse20054Projects) GetHarvestTimersEnabled() bool`

GetHarvestTimersEnabled returns the HarvestTimersEnabled field if non-nil, zero value otherwise.

### GetHarvestTimersEnabledOk

`func (o *InlineResponse20054Projects) GetHarvestTimersEnabledOk() (*bool, bool)`

GetHarvestTimersEnabledOk returns a tuple with the HarvestTimersEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHarvestTimersEnabled

`func (o *InlineResponse20054Projects) SetHarvestTimersEnabled(v bool)`

SetHarvestTimersEnabled sets HarvestTimersEnabled field to given value.

### HasHarvestTimersEnabled

`func (o *InlineResponse20054Projects) HasHarvestTimersEnabled() bool`

HasHarvestTimersEnabled returns a boolean if a field has been set.

### GetId

`func (o *InlineResponse20054Projects) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20054Projects) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20054Projects) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20054Projects) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIntegrations

`func (o *InlineResponse20054Projects) GetIntegrations() InlineResponse20054Integrations`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *InlineResponse20054Projects) GetIntegrationsOk() (*InlineResponse20054Integrations, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *InlineResponse20054Projects) SetIntegrations(v InlineResponse20054Integrations)`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *InlineResponse20054Projects) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.

### GetIsProjectAdmin

`func (o *InlineResponse20054Projects) GetIsProjectAdmin() bool`

GetIsProjectAdmin returns the IsProjectAdmin field if non-nil, zero value otherwise.

### GetIsProjectAdminOk

`func (o *InlineResponse20054Projects) GetIsProjectAdminOk() (*bool, bool)`

GetIsProjectAdminOk returns a tuple with the IsProjectAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProjectAdmin

`func (o *InlineResponse20054Projects) SetIsProjectAdmin(v bool)`

SetIsProjectAdmin sets IsProjectAdmin field to given value.

### HasIsProjectAdmin

`func (o *InlineResponse20054Projects) HasIsProjectAdmin() bool`

HasIsProjectAdmin returns a boolean if a field has been set.

### GetLastChangedOn

`func (o *InlineResponse20054Projects) GetLastChangedOn() string`

GetLastChangedOn returns the LastChangedOn field if non-nil, zero value otherwise.

### GetLastChangedOnOk

`func (o *InlineResponse20054Projects) GetLastChangedOnOk() (*string, bool)`

GetLastChangedOnOk returns a tuple with the LastChangedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangedOn

`func (o *InlineResponse20054Projects) SetLastChangedOn(v string)`

SetLastChangedOn sets LastChangedOn field to given value.

### HasLastChangedOn

`func (o *InlineResponse20054Projects) HasLastChangedOn() bool`

HasLastChangedOn returns a boolean if a field has been set.

### GetLastWorkedOn

`func (o *InlineResponse20054Projects) GetLastWorkedOn() string`

GetLastWorkedOn returns the LastWorkedOn field if non-nil, zero value otherwise.

### GetLastWorkedOnOk

`func (o *InlineResponse20054Projects) GetLastWorkedOnOk() (*string, bool)`

GetLastWorkedOnOk returns a tuple with the LastWorkedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWorkedOn

`func (o *InlineResponse20054Projects) SetLastWorkedOn(v string)`

SetLastWorkedOn sets LastWorkedOn field to given value.

### HasLastWorkedOn

`func (o *InlineResponse20054Projects) HasLastWorkedOn() bool`

HasLastWorkedOn returns a boolean if a field has been set.

### GetLogo

`func (o *InlineResponse20054Projects) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *InlineResponse20054Projects) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *InlineResponse20054Projects) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *InlineResponse20054Projects) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetLogoFromCompany

`func (o *InlineResponse20054Projects) GetLogoFromCompany() bool`

GetLogoFromCompany returns the LogoFromCompany field if non-nil, zero value otherwise.

### GetLogoFromCompanyOk

`func (o *InlineResponse20054Projects) GetLogoFromCompanyOk() (*bool, bool)`

GetLogoFromCompanyOk returns a tuple with the LogoFromCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoFromCompany

`func (o *InlineResponse20054Projects) SetLogoFromCompany(v bool)`

SetLogoFromCompany sets LogoFromCompany field to given value.

### HasLogoFromCompany

`func (o *InlineResponse20054Projects) HasLogoFromCompany() bool`

HasLogoFromCompany returns a boolean if a field has been set.

### GetMinMaxAvailableDates

`func (o *InlineResponse20054Projects) GetMinMaxAvailableDates() InlineResponse20054MinMaxAvailableDates`

GetMinMaxAvailableDates returns the MinMaxAvailableDates field if non-nil, zero value otherwise.

### GetMinMaxAvailableDatesOk

`func (o *InlineResponse20054Projects) GetMinMaxAvailableDatesOk() (*InlineResponse20054MinMaxAvailableDates, bool)`

GetMinMaxAvailableDatesOk returns a tuple with the MinMaxAvailableDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMaxAvailableDates

`func (o *InlineResponse20054Projects) SetMinMaxAvailableDates(v InlineResponse20054MinMaxAvailableDates)`

SetMinMaxAvailableDates sets MinMaxAvailableDates field to given value.

### HasMinMaxAvailableDates

`func (o *InlineResponse20054Projects) HasMinMaxAvailableDates() bool`

HasMinMaxAvailableDates returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20054Projects) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20054Projects) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20054Projects) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20054Projects) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotifyEveryone

`func (o *InlineResponse20054Projects) GetNotifyEveryone() bool`

GetNotifyEveryone returns the NotifyEveryone field if non-nil, zero value otherwise.

### GetNotifyEveryoneOk

`func (o *InlineResponse20054Projects) GetNotifyEveryoneOk() (*bool, bool)`

GetNotifyEveryoneOk returns a tuple with the NotifyEveryone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyEveryone

`func (o *InlineResponse20054Projects) SetNotifyEveryone(v bool)`

SetNotifyEveryone sets NotifyEveryone field to given value.

### HasNotifyEveryone

`func (o *InlineResponse20054Projects) HasNotifyEveryone() bool`

HasNotifyEveryone returns a boolean if a field has been set.

### GetOverviewStartPage

`func (o *InlineResponse20054Projects) GetOverviewStartPage() string`

GetOverviewStartPage returns the OverviewStartPage field if non-nil, zero value otherwise.

### GetOverviewStartPageOk

`func (o *InlineResponse20054Projects) GetOverviewStartPageOk() (*string, bool)`

GetOverviewStartPageOk returns a tuple with the OverviewStartPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverviewStartPage

`func (o *InlineResponse20054Projects) SetOverviewStartPage(v string)`

SetOverviewStartPage sets OverviewStartPage field to given value.

### HasOverviewStartPage

`func (o *InlineResponse20054Projects) HasOverviewStartPage() bool`

HasOverviewStartPage returns a boolean if a field has been set.

### GetPrivacyEnabled

`func (o *InlineResponse20054Projects) GetPrivacyEnabled() bool`

GetPrivacyEnabled returns the PrivacyEnabled field if non-nil, zero value otherwise.

### GetPrivacyEnabledOk

`func (o *InlineResponse20054Projects) GetPrivacyEnabledOk() (*bool, bool)`

GetPrivacyEnabledOk returns a tuple with the PrivacyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyEnabled

`func (o *InlineResponse20054Projects) SetPrivacyEnabled(v bool)`

SetPrivacyEnabled sets PrivacyEnabled field to given value.

### HasPrivacyEnabled

`func (o *InlineResponse20054Projects) HasPrivacyEnabled() bool`

HasPrivacyEnabled returns a boolean if a field has been set.

### GetProjectOwnerId

`func (o *InlineResponse20054Projects) GetProjectOwnerId() int32`

GetProjectOwnerId returns the ProjectOwnerId field if non-nil, zero value otherwise.

### GetProjectOwnerIdOk

`func (o *InlineResponse20054Projects) GetProjectOwnerIdOk() (*int32, bool)`

GetProjectOwnerIdOk returns a tuple with the ProjectOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectOwnerId

`func (o *InlineResponse20054Projects) SetProjectOwnerId(v int32)`

SetProjectOwnerId sets ProjectOwnerId field to given value.

### HasProjectOwnerId

`func (o *InlineResponse20054Projects) HasProjectOwnerId() bool`

HasProjectOwnerId returns a boolean if a field has been set.

### GetReplyByEmailEnabled

`func (o *InlineResponse20054Projects) GetReplyByEmailEnabled() bool`

GetReplyByEmailEnabled returns the ReplyByEmailEnabled field if non-nil, zero value otherwise.

### GetReplyByEmailEnabledOk

`func (o *InlineResponse20054Projects) GetReplyByEmailEnabledOk() (*bool, bool)`

GetReplyByEmailEnabledOk returns a tuple with the ReplyByEmailEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyByEmailEnabled

`func (o *InlineResponse20054Projects) SetReplyByEmailEnabled(v bool)`

SetReplyByEmailEnabled sets ReplyByEmailEnabled field to given value.

### HasReplyByEmailEnabled

`func (o *InlineResponse20054Projects) HasReplyByEmailEnabled() bool`

HasReplyByEmailEnabled returns a boolean if a field has been set.

### GetShowAnnouncement

`func (o *InlineResponse20054Projects) GetShowAnnouncement() bool`

GetShowAnnouncement returns the ShowAnnouncement field if non-nil, zero value otherwise.

### GetShowAnnouncementOk

`func (o *InlineResponse20054Projects) GetShowAnnouncementOk() (*bool, bool)`

GetShowAnnouncementOk returns a tuple with the ShowAnnouncement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAnnouncement

`func (o *InlineResponse20054Projects) SetShowAnnouncement(v bool)`

SetShowAnnouncement sets ShowAnnouncement field to given value.

### HasShowAnnouncement

`func (o *InlineResponse20054Projects) HasShowAnnouncement() bool`

HasShowAnnouncement returns a boolean if a field has been set.

### GetSkipWeekends

`func (o *InlineResponse20054Projects) GetSkipWeekends() bool`

GetSkipWeekends returns the SkipWeekends field if non-nil, zero value otherwise.

### GetSkipWeekendsOk

`func (o *InlineResponse20054Projects) GetSkipWeekendsOk() (*bool, bool)`

GetSkipWeekendsOk returns a tuple with the SkipWeekends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipWeekends

`func (o *InlineResponse20054Projects) SetSkipWeekends(v bool)`

SetSkipWeekends sets SkipWeekends field to given value.

### HasSkipWeekends

`func (o *InlineResponse20054Projects) HasSkipWeekends() bool`

HasSkipWeekends returns a boolean if a field has been set.

### GetStarred

`func (o *InlineResponse20054Projects) GetStarred() bool`

GetStarred returns the Starred field if non-nil, zero value otherwise.

### GetStarredOk

`func (o *InlineResponse20054Projects) GetStarredOk() (*bool, bool)`

GetStarredOk returns a tuple with the Starred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarred

`func (o *InlineResponse20054Projects) SetStarred(v bool)`

SetStarred sets Starred field to given value.

### HasStarred

`func (o *InlineResponse20054Projects) HasStarred() bool`

HasStarred returns a boolean if a field has been set.

### GetStartPage

`func (o *InlineResponse20054Projects) GetStartPage() string`

GetStartPage returns the StartPage field if non-nil, zero value otherwise.

### GetStartPageOk

`func (o *InlineResponse20054Projects) GetStartPageOk() (*string, bool)`

GetStartPageOk returns a tuple with the StartPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPage

`func (o *InlineResponse20054Projects) SetStartPage(v string)`

SetStartPage sets StartPage field to given value.

### HasStartPage

`func (o *InlineResponse20054Projects) HasStartPage() bool`

HasStartPage returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20054Projects) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20054Projects) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20054Projects) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20054Projects) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubStatus

`func (o *InlineResponse20054Projects) GetSubStatus() string`

GetSubStatus returns the SubStatus field if non-nil, zero value otherwise.

### GetSubStatusOk

`func (o *InlineResponse20054Projects) GetSubStatusOk() (*string, bool)`

GetSubStatusOk returns a tuple with the SubStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStatus

`func (o *InlineResponse20054Projects) SetSubStatus(v string)`

SetSubStatus sets SubStatus field to given value.

### HasSubStatus

`func (o *InlineResponse20054Projects) HasSubStatus() bool`

HasSubStatus returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse20054Projects) GetTags() []InlineResponse20054Tags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse20054Projects) GetTagsOk() (*[]InlineResponse20054Tags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse20054Projects) SetTags(v []InlineResponse20054Tags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse20054Projects) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTasksStartPage

`func (o *InlineResponse20054Projects) GetTasksStartPage() string`

GetTasksStartPage returns the TasksStartPage field if non-nil, zero value otherwise.

### GetTasksStartPageOk

`func (o *InlineResponse20054Projects) GetTasksStartPageOk() (*string, bool)`

GetTasksStartPageOk returns a tuple with the TasksStartPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksStartPage

`func (o *InlineResponse20054Projects) SetTasksStartPage(v string)`

SetTasksStartPage sets TasksStartPage field to given value.

### HasTasksStartPage

`func (o *InlineResponse20054Projects) HasTasksStartPage() bool`

HasTasksStartPage returns a boolean if a field has been set.

### GetTemplateDateTargetDefault

`func (o *InlineResponse20054Projects) GetTemplateDateTargetDefault() string`

GetTemplateDateTargetDefault returns the TemplateDateTargetDefault field if non-nil, zero value otherwise.

### GetTemplateDateTargetDefaultOk

`func (o *InlineResponse20054Projects) GetTemplateDateTargetDefaultOk() (*string, bool)`

GetTemplateDateTargetDefaultOk returns a tuple with the TemplateDateTargetDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateDateTargetDefault

`func (o *InlineResponse20054Projects) SetTemplateDateTargetDefault(v string)`

SetTemplateDateTargetDefault sets TemplateDateTargetDefault field to given value.

### HasTemplateDateTargetDefault

`func (o *InlineResponse20054Projects) HasTemplateDateTargetDefault() bool`

HasTemplateDateTargetDefault returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse20054Projects) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse20054Projects) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse20054Projects) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse20054Projects) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


