# InlineResponse20053Projects

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Announcement** | Pointer to **string** |  | [optional] 
**AnnouncementHTML** | Pointer to **string** |  | [optional] 
**BoardData** | Pointer to **map[string]interface{}** |  | [optional] 
**Category** | Pointer to [**InlineResponse2002Column**](InlineResponse2002Column.md) |  | [optional] 
**Company** | Pointer to [**InlineResponse20014Company**](InlineResponse20014Company.md) |  | [optional] 
**CreatedOn** | Pointer to **string** |  | [optional] 
**DefaultPrivacy** | Pointer to **string** |  | [optional] 
**Defaults** | Pointer to [**InlineResponse20053Defaults**](InlineResponse20053Defaults.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**FilesAutoNewVersion** | Pointer to **bool** |  | [optional] 
**HarvestTimersEnabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Integrations** | Pointer to [**InlineResponse20053Integrations**](InlineResponse20053Integrations.md) |  | [optional] 
**IsProjectAdmin** | Pointer to **bool** |  | [optional] 
**LastChangedOn** | Pointer to **string** |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Notifyeveryone** | Pointer to **bool** |  | [optional] 
**OverviewStartPage** | Pointer to **string** |  | [optional] 
**PrivacyEnabled** | Pointer to **bool** |  | [optional] 
**ReplyByEmailEnabled** | Pointer to **bool** |  | [optional] 
**ShowAnnouncement** | Pointer to **bool** |  | [optional] 
**Starred** | Pointer to **bool** |  | [optional] 
**StartPage** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SubStatus** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]map[string]interface{}** |  | [optional] 
**TasksStartPage** | Pointer to **string** |  | [optional] 

## Methods

### NewInlineResponse20053Projects

`func NewInlineResponse20053Projects() *InlineResponse20053Projects`

NewInlineResponse20053Projects instantiates a new InlineResponse20053Projects object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20053ProjectsWithDefaults

`func NewInlineResponse20053ProjectsWithDefaults() *InlineResponse20053Projects`

NewInlineResponse20053ProjectsWithDefaults instantiates a new InlineResponse20053Projects object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnouncement

`func (o *InlineResponse20053Projects) GetAnnouncement() string`

GetAnnouncement returns the Announcement field if non-nil, zero value otherwise.

### GetAnnouncementOk

`func (o *InlineResponse20053Projects) GetAnnouncementOk() (*string, bool)`

GetAnnouncementOk returns a tuple with the Announcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncement

`func (o *InlineResponse20053Projects) SetAnnouncement(v string)`

SetAnnouncement sets Announcement field to given value.

### HasAnnouncement

`func (o *InlineResponse20053Projects) HasAnnouncement() bool`

HasAnnouncement returns a boolean if a field has been set.

### GetAnnouncementHTML

`func (o *InlineResponse20053Projects) GetAnnouncementHTML() string`

GetAnnouncementHTML returns the AnnouncementHTML field if non-nil, zero value otherwise.

### GetAnnouncementHTMLOk

`func (o *InlineResponse20053Projects) GetAnnouncementHTMLOk() (*string, bool)`

GetAnnouncementHTMLOk returns a tuple with the AnnouncementHTML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementHTML

`func (o *InlineResponse20053Projects) SetAnnouncementHTML(v string)`

SetAnnouncementHTML sets AnnouncementHTML field to given value.

### HasAnnouncementHTML

`func (o *InlineResponse20053Projects) HasAnnouncementHTML() bool`

HasAnnouncementHTML returns a boolean if a field has been set.

### GetBoardData

`func (o *InlineResponse20053Projects) GetBoardData() map[string]interface{}`

GetBoardData returns the BoardData field if non-nil, zero value otherwise.

### GetBoardDataOk

`func (o *InlineResponse20053Projects) GetBoardDataOk() (*map[string]interface{}, bool)`

GetBoardDataOk returns a tuple with the BoardData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardData

`func (o *InlineResponse20053Projects) SetBoardData(v map[string]interface{})`

SetBoardData sets BoardData field to given value.

### HasBoardData

`func (o *InlineResponse20053Projects) HasBoardData() bool`

HasBoardData returns a boolean if a field has been set.

### GetCategory

`func (o *InlineResponse20053Projects) GetCategory() InlineResponse2002Column`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InlineResponse20053Projects) GetCategoryOk() (*InlineResponse2002Column, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InlineResponse20053Projects) SetCategory(v InlineResponse2002Column)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *InlineResponse20053Projects) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCompany

`func (o *InlineResponse20053Projects) GetCompany() InlineResponse20014Company`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *InlineResponse20053Projects) GetCompanyOk() (*InlineResponse20014Company, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *InlineResponse20053Projects) SetCompany(v InlineResponse20014Company)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *InlineResponse20053Projects) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCreatedOn

`func (o *InlineResponse20053Projects) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *InlineResponse20053Projects) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *InlineResponse20053Projects) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *InlineResponse20053Projects) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetDefaultPrivacy

`func (o *InlineResponse20053Projects) GetDefaultPrivacy() string`

GetDefaultPrivacy returns the DefaultPrivacy field if non-nil, zero value otherwise.

### GetDefaultPrivacyOk

`func (o *InlineResponse20053Projects) GetDefaultPrivacyOk() (*string, bool)`

GetDefaultPrivacyOk returns a tuple with the DefaultPrivacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPrivacy

`func (o *InlineResponse20053Projects) SetDefaultPrivacy(v string)`

SetDefaultPrivacy sets DefaultPrivacy field to given value.

### HasDefaultPrivacy

`func (o *InlineResponse20053Projects) HasDefaultPrivacy() bool`

HasDefaultPrivacy returns a boolean if a field has been set.

### GetDefaults

`func (o *InlineResponse20053Projects) GetDefaults() InlineResponse20053Defaults`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *InlineResponse20053Projects) GetDefaultsOk() (*InlineResponse20053Defaults, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *InlineResponse20053Projects) SetDefaults(v InlineResponse20053Defaults)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *InlineResponse20053Projects) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse20053Projects) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse20053Projects) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse20053Projects) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse20053Projects) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndDate

`func (o *InlineResponse20053Projects) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *InlineResponse20053Projects) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *InlineResponse20053Projects) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *InlineResponse20053Projects) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetFilesAutoNewVersion

`func (o *InlineResponse20053Projects) GetFilesAutoNewVersion() bool`

GetFilesAutoNewVersion returns the FilesAutoNewVersion field if non-nil, zero value otherwise.

### GetFilesAutoNewVersionOk

`func (o *InlineResponse20053Projects) GetFilesAutoNewVersionOk() (*bool, bool)`

GetFilesAutoNewVersionOk returns a tuple with the FilesAutoNewVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesAutoNewVersion

`func (o *InlineResponse20053Projects) SetFilesAutoNewVersion(v bool)`

SetFilesAutoNewVersion sets FilesAutoNewVersion field to given value.

### HasFilesAutoNewVersion

`func (o *InlineResponse20053Projects) HasFilesAutoNewVersion() bool`

HasFilesAutoNewVersion returns a boolean if a field has been set.

### GetHarvestTimersEnabled

`func (o *InlineResponse20053Projects) GetHarvestTimersEnabled() bool`

GetHarvestTimersEnabled returns the HarvestTimersEnabled field if non-nil, zero value otherwise.

### GetHarvestTimersEnabledOk

`func (o *InlineResponse20053Projects) GetHarvestTimersEnabledOk() (*bool, bool)`

GetHarvestTimersEnabledOk returns a tuple with the HarvestTimersEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHarvestTimersEnabled

`func (o *InlineResponse20053Projects) SetHarvestTimersEnabled(v bool)`

SetHarvestTimersEnabled sets HarvestTimersEnabled field to given value.

### HasHarvestTimersEnabled

`func (o *InlineResponse20053Projects) HasHarvestTimersEnabled() bool`

HasHarvestTimersEnabled returns a boolean if a field has been set.

### GetId

`func (o *InlineResponse20053Projects) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20053Projects) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20053Projects) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20053Projects) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIntegrations

`func (o *InlineResponse20053Projects) GetIntegrations() InlineResponse20053Integrations`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *InlineResponse20053Projects) GetIntegrationsOk() (*InlineResponse20053Integrations, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *InlineResponse20053Projects) SetIntegrations(v InlineResponse20053Integrations)`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *InlineResponse20053Projects) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.

### GetIsProjectAdmin

`func (o *InlineResponse20053Projects) GetIsProjectAdmin() bool`

GetIsProjectAdmin returns the IsProjectAdmin field if non-nil, zero value otherwise.

### GetIsProjectAdminOk

`func (o *InlineResponse20053Projects) GetIsProjectAdminOk() (*bool, bool)`

GetIsProjectAdminOk returns a tuple with the IsProjectAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProjectAdmin

`func (o *InlineResponse20053Projects) SetIsProjectAdmin(v bool)`

SetIsProjectAdmin sets IsProjectAdmin field to given value.

### HasIsProjectAdmin

`func (o *InlineResponse20053Projects) HasIsProjectAdmin() bool`

HasIsProjectAdmin returns a boolean if a field has been set.

### GetLastChangedOn

`func (o *InlineResponse20053Projects) GetLastChangedOn() string`

GetLastChangedOn returns the LastChangedOn field if non-nil, zero value otherwise.

### GetLastChangedOnOk

`func (o *InlineResponse20053Projects) GetLastChangedOnOk() (*string, bool)`

GetLastChangedOnOk returns a tuple with the LastChangedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangedOn

`func (o *InlineResponse20053Projects) SetLastChangedOn(v string)`

SetLastChangedOn sets LastChangedOn field to given value.

### HasLastChangedOn

`func (o *InlineResponse20053Projects) HasLastChangedOn() bool`

HasLastChangedOn returns a boolean if a field has been set.

### GetLogo

`func (o *InlineResponse20053Projects) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *InlineResponse20053Projects) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *InlineResponse20053Projects) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *InlineResponse20053Projects) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20053Projects) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20053Projects) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20053Projects) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20053Projects) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotifyeveryone

`func (o *InlineResponse20053Projects) GetNotifyeveryone() bool`

GetNotifyeveryone returns the Notifyeveryone field if non-nil, zero value otherwise.

### GetNotifyeveryoneOk

`func (o *InlineResponse20053Projects) GetNotifyeveryoneOk() (*bool, bool)`

GetNotifyeveryoneOk returns a tuple with the Notifyeveryone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyeveryone

`func (o *InlineResponse20053Projects) SetNotifyeveryone(v bool)`

SetNotifyeveryone sets Notifyeveryone field to given value.

### HasNotifyeveryone

`func (o *InlineResponse20053Projects) HasNotifyeveryone() bool`

HasNotifyeveryone returns a boolean if a field has been set.

### GetOverviewStartPage

`func (o *InlineResponse20053Projects) GetOverviewStartPage() string`

GetOverviewStartPage returns the OverviewStartPage field if non-nil, zero value otherwise.

### GetOverviewStartPageOk

`func (o *InlineResponse20053Projects) GetOverviewStartPageOk() (*string, bool)`

GetOverviewStartPageOk returns a tuple with the OverviewStartPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverviewStartPage

`func (o *InlineResponse20053Projects) SetOverviewStartPage(v string)`

SetOverviewStartPage sets OverviewStartPage field to given value.

### HasOverviewStartPage

`func (o *InlineResponse20053Projects) HasOverviewStartPage() bool`

HasOverviewStartPage returns a boolean if a field has been set.

### GetPrivacyEnabled

`func (o *InlineResponse20053Projects) GetPrivacyEnabled() bool`

GetPrivacyEnabled returns the PrivacyEnabled field if non-nil, zero value otherwise.

### GetPrivacyEnabledOk

`func (o *InlineResponse20053Projects) GetPrivacyEnabledOk() (*bool, bool)`

GetPrivacyEnabledOk returns a tuple with the PrivacyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyEnabled

`func (o *InlineResponse20053Projects) SetPrivacyEnabled(v bool)`

SetPrivacyEnabled sets PrivacyEnabled field to given value.

### HasPrivacyEnabled

`func (o *InlineResponse20053Projects) HasPrivacyEnabled() bool`

HasPrivacyEnabled returns a boolean if a field has been set.

### GetReplyByEmailEnabled

`func (o *InlineResponse20053Projects) GetReplyByEmailEnabled() bool`

GetReplyByEmailEnabled returns the ReplyByEmailEnabled field if non-nil, zero value otherwise.

### GetReplyByEmailEnabledOk

`func (o *InlineResponse20053Projects) GetReplyByEmailEnabledOk() (*bool, bool)`

GetReplyByEmailEnabledOk returns a tuple with the ReplyByEmailEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyByEmailEnabled

`func (o *InlineResponse20053Projects) SetReplyByEmailEnabled(v bool)`

SetReplyByEmailEnabled sets ReplyByEmailEnabled field to given value.

### HasReplyByEmailEnabled

`func (o *InlineResponse20053Projects) HasReplyByEmailEnabled() bool`

HasReplyByEmailEnabled returns a boolean if a field has been set.

### GetShowAnnouncement

`func (o *InlineResponse20053Projects) GetShowAnnouncement() bool`

GetShowAnnouncement returns the ShowAnnouncement field if non-nil, zero value otherwise.

### GetShowAnnouncementOk

`func (o *InlineResponse20053Projects) GetShowAnnouncementOk() (*bool, bool)`

GetShowAnnouncementOk returns a tuple with the ShowAnnouncement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAnnouncement

`func (o *InlineResponse20053Projects) SetShowAnnouncement(v bool)`

SetShowAnnouncement sets ShowAnnouncement field to given value.

### HasShowAnnouncement

`func (o *InlineResponse20053Projects) HasShowAnnouncement() bool`

HasShowAnnouncement returns a boolean if a field has been set.

### GetStarred

`func (o *InlineResponse20053Projects) GetStarred() bool`

GetStarred returns the Starred field if non-nil, zero value otherwise.

### GetStarredOk

`func (o *InlineResponse20053Projects) GetStarredOk() (*bool, bool)`

GetStarredOk returns a tuple with the Starred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarred

`func (o *InlineResponse20053Projects) SetStarred(v bool)`

SetStarred sets Starred field to given value.

### HasStarred

`func (o *InlineResponse20053Projects) HasStarred() bool`

HasStarred returns a boolean if a field has been set.

### GetStartPage

`func (o *InlineResponse20053Projects) GetStartPage() string`

GetStartPage returns the StartPage field if non-nil, zero value otherwise.

### GetStartPageOk

`func (o *InlineResponse20053Projects) GetStartPageOk() (*string, bool)`

GetStartPageOk returns a tuple with the StartPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPage

`func (o *InlineResponse20053Projects) SetStartPage(v string)`

SetStartPage sets StartPage field to given value.

### HasStartPage

`func (o *InlineResponse20053Projects) HasStartPage() bool`

HasStartPage returns a boolean if a field has been set.

### GetStartDate

`func (o *InlineResponse20053Projects) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InlineResponse20053Projects) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InlineResponse20053Projects) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InlineResponse20053Projects) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20053Projects) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20053Projects) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20053Projects) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20053Projects) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubStatus

`func (o *InlineResponse20053Projects) GetSubStatus() string`

GetSubStatus returns the SubStatus field if non-nil, zero value otherwise.

### GetSubStatusOk

`func (o *InlineResponse20053Projects) GetSubStatusOk() (*string, bool)`

GetSubStatusOk returns a tuple with the SubStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStatus

`func (o *InlineResponse20053Projects) SetSubStatus(v string)`

SetSubStatus sets SubStatus field to given value.

### HasSubStatus

`func (o *InlineResponse20053Projects) HasSubStatus() bool`

HasSubStatus returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse20053Projects) GetTags() []map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse20053Projects) GetTagsOk() (*[]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse20053Projects) SetTags(v []map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse20053Projects) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTasksStartPage

`func (o *InlineResponse20053Projects) GetTasksStartPage() string`

GetTasksStartPage returns the TasksStartPage field if non-nil, zero value otherwise.

### GetTasksStartPageOk

`func (o *InlineResponse20053Projects) GetTasksStartPageOk() (*string, bool)`

GetTasksStartPageOk returns a tuple with the TasksStartPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksStartPage

`func (o *InlineResponse20053Projects) SetTasksStartPage(v string)`

SetTasksStartPage sets TasksStartPage field to given value.

### HasTasksStartPage

`func (o *InlineResponse20053Projects) HasTasksStartPage() bool`

HasTasksStartPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


