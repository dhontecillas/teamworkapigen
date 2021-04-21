# InlineResponse200121Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivePages** | Pointer to [**InlineResponse20060ProjectActivePages**](InlineResponse20060ProjectActivePages.md) |  | [optional] 
**Announcement** | Pointer to **string** |  | [optional] 
**AnnouncementHTML** | Pointer to **string** |  | [optional] 
**BoardData** | Pointer to [**InlineResponse200121ProjectBoardData**](InlineResponse200121ProjectBoardData.md) |  | [optional] 
**Category** | Pointer to [**InlineResponse2002Column**](InlineResponse2002Column.md) |  | [optional] 
**Company** | Pointer to [**InlineResponse2002People12345Company**](InlineResponse2002People12345Company.md) |  | [optional] 
**CreatedOn** | Pointer to **string** |  | [optional] 
**DefaultPrivacy** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DirectFileUploadsEnabled** | Pointer to **bool** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**FilesAutoNewVersion** | Pointer to **bool** |  | [optional] 
**HarvestTimersEnabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Integrations** | Pointer to [**InlineResponse20060ProjectIntegrations**](InlineResponse20060ProjectIntegrations.md) |  | [optional] 
**LastChangedOn** | Pointer to **string** |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**LogoFromCompany** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Notifyeveryone** | Pointer to **bool** |  | [optional] 
**OverviewStartPage** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**InlineResponse20045CardOwner**](InlineResponse20045CardOwner.md) |  | [optional] 
**PrivacyEnabled** | Pointer to **bool** |  | [optional] 
**ReplyByEmailEnabled** | Pointer to **bool** |  | [optional] 
**ShowAnnouncement** | Pointer to **bool** |  | [optional] 
**SkipWeekends** | Pointer to **bool** |  | [optional] 
**Starred** | Pointer to **bool** |  | [optional] 
**StartPage** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SubStatus** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]map[string]interface{}** |  | [optional] 
**TasksStartPage** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewInlineResponse200121Project

`func NewInlineResponse200121Project() *InlineResponse200121Project`

NewInlineResponse200121Project instantiates a new InlineResponse200121Project object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200121ProjectWithDefaults

`func NewInlineResponse200121ProjectWithDefaults() *InlineResponse200121Project`

NewInlineResponse200121ProjectWithDefaults instantiates a new InlineResponse200121Project object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivePages

`func (o *InlineResponse200121Project) GetActivePages() InlineResponse20060ProjectActivePages`

GetActivePages returns the ActivePages field if non-nil, zero value otherwise.

### GetActivePagesOk

`func (o *InlineResponse200121Project) GetActivePagesOk() (*InlineResponse20060ProjectActivePages, bool)`

GetActivePagesOk returns a tuple with the ActivePages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePages

`func (o *InlineResponse200121Project) SetActivePages(v InlineResponse20060ProjectActivePages)`

SetActivePages sets ActivePages field to given value.

### HasActivePages

`func (o *InlineResponse200121Project) HasActivePages() bool`

HasActivePages returns a boolean if a field has been set.

### GetAnnouncement

`func (o *InlineResponse200121Project) GetAnnouncement() string`

GetAnnouncement returns the Announcement field if non-nil, zero value otherwise.

### GetAnnouncementOk

`func (o *InlineResponse200121Project) GetAnnouncementOk() (*string, bool)`

GetAnnouncementOk returns a tuple with the Announcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncement

`func (o *InlineResponse200121Project) SetAnnouncement(v string)`

SetAnnouncement sets Announcement field to given value.

### HasAnnouncement

`func (o *InlineResponse200121Project) HasAnnouncement() bool`

HasAnnouncement returns a boolean if a field has been set.

### GetAnnouncementHTML

`func (o *InlineResponse200121Project) GetAnnouncementHTML() string`

GetAnnouncementHTML returns the AnnouncementHTML field if non-nil, zero value otherwise.

### GetAnnouncementHTMLOk

`func (o *InlineResponse200121Project) GetAnnouncementHTMLOk() (*string, bool)`

GetAnnouncementHTMLOk returns a tuple with the AnnouncementHTML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementHTML

`func (o *InlineResponse200121Project) SetAnnouncementHTML(v string)`

SetAnnouncementHTML sets AnnouncementHTML field to given value.

### HasAnnouncementHTML

`func (o *InlineResponse200121Project) HasAnnouncementHTML() bool`

HasAnnouncementHTML returns a boolean if a field has been set.

### GetBoardData

`func (o *InlineResponse200121Project) GetBoardData() InlineResponse200121ProjectBoardData`

GetBoardData returns the BoardData field if non-nil, zero value otherwise.

### GetBoardDataOk

`func (o *InlineResponse200121Project) GetBoardDataOk() (*InlineResponse200121ProjectBoardData, bool)`

GetBoardDataOk returns a tuple with the BoardData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardData

`func (o *InlineResponse200121Project) SetBoardData(v InlineResponse200121ProjectBoardData)`

SetBoardData sets BoardData field to given value.

### HasBoardData

`func (o *InlineResponse200121Project) HasBoardData() bool`

HasBoardData returns a boolean if a field has been set.

### GetCategory

`func (o *InlineResponse200121Project) GetCategory() InlineResponse2002Column`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InlineResponse200121Project) GetCategoryOk() (*InlineResponse2002Column, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InlineResponse200121Project) SetCategory(v InlineResponse2002Column)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *InlineResponse200121Project) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCompany

`func (o *InlineResponse200121Project) GetCompany() InlineResponse2002People12345Company`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *InlineResponse200121Project) GetCompanyOk() (*InlineResponse2002People12345Company, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *InlineResponse200121Project) SetCompany(v InlineResponse2002People12345Company)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *InlineResponse200121Project) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCreatedOn

`func (o *InlineResponse200121Project) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *InlineResponse200121Project) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *InlineResponse200121Project) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *InlineResponse200121Project) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetDefaultPrivacy

`func (o *InlineResponse200121Project) GetDefaultPrivacy() string`

GetDefaultPrivacy returns the DefaultPrivacy field if non-nil, zero value otherwise.

### GetDefaultPrivacyOk

`func (o *InlineResponse200121Project) GetDefaultPrivacyOk() (*string, bool)`

GetDefaultPrivacyOk returns a tuple with the DefaultPrivacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPrivacy

`func (o *InlineResponse200121Project) SetDefaultPrivacy(v string)`

SetDefaultPrivacy sets DefaultPrivacy field to given value.

### HasDefaultPrivacy

`func (o *InlineResponse200121Project) HasDefaultPrivacy() bool`

HasDefaultPrivacy returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse200121Project) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse200121Project) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse200121Project) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse200121Project) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDirectFileUploadsEnabled

`func (o *InlineResponse200121Project) GetDirectFileUploadsEnabled() bool`

GetDirectFileUploadsEnabled returns the DirectFileUploadsEnabled field if non-nil, zero value otherwise.

### GetDirectFileUploadsEnabledOk

`func (o *InlineResponse200121Project) GetDirectFileUploadsEnabledOk() (*bool, bool)`

GetDirectFileUploadsEnabledOk returns a tuple with the DirectFileUploadsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectFileUploadsEnabled

`func (o *InlineResponse200121Project) SetDirectFileUploadsEnabled(v bool)`

SetDirectFileUploadsEnabled sets DirectFileUploadsEnabled field to given value.

### HasDirectFileUploadsEnabled

`func (o *InlineResponse200121Project) HasDirectFileUploadsEnabled() bool`

HasDirectFileUploadsEnabled returns a boolean if a field has been set.

### GetEndDate

`func (o *InlineResponse200121Project) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *InlineResponse200121Project) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *InlineResponse200121Project) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *InlineResponse200121Project) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetFilesAutoNewVersion

`func (o *InlineResponse200121Project) GetFilesAutoNewVersion() bool`

GetFilesAutoNewVersion returns the FilesAutoNewVersion field if non-nil, zero value otherwise.

### GetFilesAutoNewVersionOk

`func (o *InlineResponse200121Project) GetFilesAutoNewVersionOk() (*bool, bool)`

GetFilesAutoNewVersionOk returns a tuple with the FilesAutoNewVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesAutoNewVersion

`func (o *InlineResponse200121Project) SetFilesAutoNewVersion(v bool)`

SetFilesAutoNewVersion sets FilesAutoNewVersion field to given value.

### HasFilesAutoNewVersion

`func (o *InlineResponse200121Project) HasFilesAutoNewVersion() bool`

HasFilesAutoNewVersion returns a boolean if a field has been set.

### GetHarvestTimersEnabled

`func (o *InlineResponse200121Project) GetHarvestTimersEnabled() bool`

GetHarvestTimersEnabled returns the HarvestTimersEnabled field if non-nil, zero value otherwise.

### GetHarvestTimersEnabledOk

`func (o *InlineResponse200121Project) GetHarvestTimersEnabledOk() (*bool, bool)`

GetHarvestTimersEnabledOk returns a tuple with the HarvestTimersEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHarvestTimersEnabled

`func (o *InlineResponse200121Project) SetHarvestTimersEnabled(v bool)`

SetHarvestTimersEnabled sets HarvestTimersEnabled field to given value.

### HasHarvestTimersEnabled

`func (o *InlineResponse200121Project) HasHarvestTimersEnabled() bool`

HasHarvestTimersEnabled returns a boolean if a field has been set.

### GetId

`func (o *InlineResponse200121Project) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200121Project) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200121Project) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200121Project) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIntegrations

`func (o *InlineResponse200121Project) GetIntegrations() InlineResponse20060ProjectIntegrations`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *InlineResponse200121Project) GetIntegrationsOk() (*InlineResponse20060ProjectIntegrations, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *InlineResponse200121Project) SetIntegrations(v InlineResponse20060ProjectIntegrations)`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *InlineResponse200121Project) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.

### GetLastChangedOn

`func (o *InlineResponse200121Project) GetLastChangedOn() string`

GetLastChangedOn returns the LastChangedOn field if non-nil, zero value otherwise.

### GetLastChangedOnOk

`func (o *InlineResponse200121Project) GetLastChangedOnOk() (*string, bool)`

GetLastChangedOnOk returns a tuple with the LastChangedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangedOn

`func (o *InlineResponse200121Project) SetLastChangedOn(v string)`

SetLastChangedOn sets LastChangedOn field to given value.

### HasLastChangedOn

`func (o *InlineResponse200121Project) HasLastChangedOn() bool`

HasLastChangedOn returns a boolean if a field has been set.

### GetLogo

`func (o *InlineResponse200121Project) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *InlineResponse200121Project) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *InlineResponse200121Project) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *InlineResponse200121Project) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetLogoFromCompany

`func (o *InlineResponse200121Project) GetLogoFromCompany() bool`

GetLogoFromCompany returns the LogoFromCompany field if non-nil, zero value otherwise.

### GetLogoFromCompanyOk

`func (o *InlineResponse200121Project) GetLogoFromCompanyOk() (*bool, bool)`

GetLogoFromCompanyOk returns a tuple with the LogoFromCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoFromCompany

`func (o *InlineResponse200121Project) SetLogoFromCompany(v bool)`

SetLogoFromCompany sets LogoFromCompany field to given value.

### HasLogoFromCompany

`func (o *InlineResponse200121Project) HasLogoFromCompany() bool`

HasLogoFromCompany returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200121Project) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200121Project) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200121Project) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200121Project) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotifyeveryone

`func (o *InlineResponse200121Project) GetNotifyeveryone() bool`

GetNotifyeveryone returns the Notifyeveryone field if non-nil, zero value otherwise.

### GetNotifyeveryoneOk

`func (o *InlineResponse200121Project) GetNotifyeveryoneOk() (*bool, bool)`

GetNotifyeveryoneOk returns a tuple with the Notifyeveryone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyeveryone

`func (o *InlineResponse200121Project) SetNotifyeveryone(v bool)`

SetNotifyeveryone sets Notifyeveryone field to given value.

### HasNotifyeveryone

`func (o *InlineResponse200121Project) HasNotifyeveryone() bool`

HasNotifyeveryone returns a boolean if a field has been set.

### GetOverviewStartPage

`func (o *InlineResponse200121Project) GetOverviewStartPage() string`

GetOverviewStartPage returns the OverviewStartPage field if non-nil, zero value otherwise.

### GetOverviewStartPageOk

`func (o *InlineResponse200121Project) GetOverviewStartPageOk() (*string, bool)`

GetOverviewStartPageOk returns a tuple with the OverviewStartPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverviewStartPage

`func (o *InlineResponse200121Project) SetOverviewStartPage(v string)`

SetOverviewStartPage sets OverviewStartPage field to given value.

### HasOverviewStartPage

`func (o *InlineResponse200121Project) HasOverviewStartPage() bool`

HasOverviewStartPage returns a boolean if a field has been set.

### GetOwner

`func (o *InlineResponse200121Project) GetOwner() InlineResponse20045CardOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *InlineResponse200121Project) GetOwnerOk() (*InlineResponse20045CardOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *InlineResponse200121Project) SetOwner(v InlineResponse20045CardOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *InlineResponse200121Project) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPrivacyEnabled

`func (o *InlineResponse200121Project) GetPrivacyEnabled() bool`

GetPrivacyEnabled returns the PrivacyEnabled field if non-nil, zero value otherwise.

### GetPrivacyEnabledOk

`func (o *InlineResponse200121Project) GetPrivacyEnabledOk() (*bool, bool)`

GetPrivacyEnabledOk returns a tuple with the PrivacyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyEnabled

`func (o *InlineResponse200121Project) SetPrivacyEnabled(v bool)`

SetPrivacyEnabled sets PrivacyEnabled field to given value.

### HasPrivacyEnabled

`func (o *InlineResponse200121Project) HasPrivacyEnabled() bool`

HasPrivacyEnabled returns a boolean if a field has been set.

### GetReplyByEmailEnabled

`func (o *InlineResponse200121Project) GetReplyByEmailEnabled() bool`

GetReplyByEmailEnabled returns the ReplyByEmailEnabled field if non-nil, zero value otherwise.

### GetReplyByEmailEnabledOk

`func (o *InlineResponse200121Project) GetReplyByEmailEnabledOk() (*bool, bool)`

GetReplyByEmailEnabledOk returns a tuple with the ReplyByEmailEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyByEmailEnabled

`func (o *InlineResponse200121Project) SetReplyByEmailEnabled(v bool)`

SetReplyByEmailEnabled sets ReplyByEmailEnabled field to given value.

### HasReplyByEmailEnabled

`func (o *InlineResponse200121Project) HasReplyByEmailEnabled() bool`

HasReplyByEmailEnabled returns a boolean if a field has been set.

### GetShowAnnouncement

`func (o *InlineResponse200121Project) GetShowAnnouncement() bool`

GetShowAnnouncement returns the ShowAnnouncement field if non-nil, zero value otherwise.

### GetShowAnnouncementOk

`func (o *InlineResponse200121Project) GetShowAnnouncementOk() (*bool, bool)`

GetShowAnnouncementOk returns a tuple with the ShowAnnouncement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAnnouncement

`func (o *InlineResponse200121Project) SetShowAnnouncement(v bool)`

SetShowAnnouncement sets ShowAnnouncement field to given value.

### HasShowAnnouncement

`func (o *InlineResponse200121Project) HasShowAnnouncement() bool`

HasShowAnnouncement returns a boolean if a field has been set.

### GetSkipWeekends

`func (o *InlineResponse200121Project) GetSkipWeekends() bool`

GetSkipWeekends returns the SkipWeekends field if non-nil, zero value otherwise.

### GetSkipWeekendsOk

`func (o *InlineResponse200121Project) GetSkipWeekendsOk() (*bool, bool)`

GetSkipWeekendsOk returns a tuple with the SkipWeekends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipWeekends

`func (o *InlineResponse200121Project) SetSkipWeekends(v bool)`

SetSkipWeekends sets SkipWeekends field to given value.

### HasSkipWeekends

`func (o *InlineResponse200121Project) HasSkipWeekends() bool`

HasSkipWeekends returns a boolean if a field has been set.

### GetStarred

`func (o *InlineResponse200121Project) GetStarred() bool`

GetStarred returns the Starred field if non-nil, zero value otherwise.

### GetStarredOk

`func (o *InlineResponse200121Project) GetStarredOk() (*bool, bool)`

GetStarredOk returns a tuple with the Starred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarred

`func (o *InlineResponse200121Project) SetStarred(v bool)`

SetStarred sets Starred field to given value.

### HasStarred

`func (o *InlineResponse200121Project) HasStarred() bool`

HasStarred returns a boolean if a field has been set.

### GetStartPage

`func (o *InlineResponse200121Project) GetStartPage() string`

GetStartPage returns the StartPage field if non-nil, zero value otherwise.

### GetStartPageOk

`func (o *InlineResponse200121Project) GetStartPageOk() (*string, bool)`

GetStartPageOk returns a tuple with the StartPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPage

`func (o *InlineResponse200121Project) SetStartPage(v string)`

SetStartPage sets StartPage field to given value.

### HasStartPage

`func (o *InlineResponse200121Project) HasStartPage() bool`

HasStartPage returns a boolean if a field has been set.

### GetStartDate

`func (o *InlineResponse200121Project) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InlineResponse200121Project) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InlineResponse200121Project) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InlineResponse200121Project) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200121Project) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200121Project) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200121Project) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200121Project) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubStatus

`func (o *InlineResponse200121Project) GetSubStatus() string`

GetSubStatus returns the SubStatus field if non-nil, zero value otherwise.

### GetSubStatusOk

`func (o *InlineResponse200121Project) GetSubStatusOk() (*string, bool)`

GetSubStatusOk returns a tuple with the SubStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStatus

`func (o *InlineResponse200121Project) SetSubStatus(v string)`

SetSubStatus sets SubStatus field to given value.

### HasSubStatus

`func (o *InlineResponse200121Project) HasSubStatus() bool`

HasSubStatus returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse200121Project) GetTags() []map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse200121Project) GetTagsOk() (*[]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse200121Project) SetTags(v []map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse200121Project) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTasksStartPage

`func (o *InlineResponse200121Project) GetTasksStartPage() string`

GetTasksStartPage returns the TasksStartPage field if non-nil, zero value otherwise.

### GetTasksStartPageOk

`func (o *InlineResponse200121Project) GetTasksStartPageOk() (*string, bool)`

GetTasksStartPageOk returns a tuple with the TasksStartPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksStartPage

`func (o *InlineResponse200121Project) SetTasksStartPage(v string)`

SetTasksStartPage sets TasksStartPage field to given value.

### HasTasksStartPage

`func (o *InlineResponse200121Project) HasTasksStartPage() bool`

HasTasksStartPage returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200121Project) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200121Project) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200121Project) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200121Project) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


