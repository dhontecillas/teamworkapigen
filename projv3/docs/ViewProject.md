# ViewProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivePages** | Pointer to [**ViewProjectActivePages**](ViewProjectActivePages.md) |  | [optional] 
**Category** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**CategoryId** | Pointer to **int32** |  | [optional] 
**Company** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**CompanyId** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **int32** |  | [optional] 
**CustomFieldValueIds** | Pointer to **[]int32** |  | [optional] 
**CustomFieldValues** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**DefaultPrivacy** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DirectFileUploadsEnabled** | Pointer to **bool** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**FinancialBudget** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**FinancialBudgetId** | Pointer to **int32** |  | [optional] 
**HarvestTimersEnabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Integrations** | Pointer to [**ViewProjectIntegrations**](ViewProjectIntegrations.md) |  | [optional] 
**IsOnBoardingProject** | Pointer to **bool** |  | [optional] 
**IsSampleProject** | Pointer to **bool** |  | [optional] 
**IsStarred** | Pointer to **bool** |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**MinMaxAvailableDates** | Pointer to [**ViewProjectMinMaxAvailableDates**](ViewProjectMinMaxAvailableDates.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NotifyCommentIncludeCreator** | Pointer to **bool** |  | [optional] 
**NotifyEveryone** | Pointer to **bool** |  | [optional] 
**OverviewStartPage** | Pointer to **string** |  | [optional] 
**OwnedBy** | Pointer to **int32** |  | [optional] 
**OwnerId** | Pointer to **int32** |  | [optional] 
**PrivacyEnabled** | Pointer to **bool** |  | [optional] 
**ProjectOwnerId** | Pointer to **int32** |  | [optional] 
**ReplyByEmailEnabled** | Pointer to **bool** |  | [optional] 
**ShowAnnouncement** | Pointer to **bool** |  | [optional] 
**SkipWeekends** | Pointer to **bool** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**StartPage** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SubStatus** | Pointer to **string** |  | [optional] 
**TagIds** | Pointer to **[]int32** |  | [optional] 
**Tags** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**TasksStartPage** | Pointer to **string** |  | [optional] 
**TimeBudget** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**TimeBudgetId** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdateId** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **int32** |  | [optional] 

## Methods

### NewViewProject

`func NewViewProject() *ViewProject`

NewViewProject instantiates a new ViewProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewProjectWithDefaults

`func NewViewProjectWithDefaults() *ViewProject`

NewViewProjectWithDefaults instantiates a new ViewProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivePages

`func (o *ViewProject) GetActivePages() ViewProjectActivePages`

GetActivePages returns the ActivePages field if non-nil, zero value otherwise.

### GetActivePagesOk

`func (o *ViewProject) GetActivePagesOk() (*ViewProjectActivePages, bool)`

GetActivePagesOk returns a tuple with the ActivePages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePages

`func (o *ViewProject) SetActivePages(v ViewProjectActivePages)`

SetActivePages sets ActivePages field to given value.

### HasActivePages

`func (o *ViewProject) HasActivePages() bool`

HasActivePages returns a boolean if a field has been set.

### GetCategory

`func (o *ViewProject) GetCategory() ViewRelationship`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ViewProject) GetCategoryOk() (*ViewRelationship, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ViewProject) SetCategory(v ViewRelationship)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ViewProject) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCategoryId

`func (o *ViewProject) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *ViewProject) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *ViewProject) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *ViewProject) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetCompany

`func (o *ViewProject) GetCompany() ViewRelationship`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ViewProject) GetCompanyOk() (*ViewRelationship, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ViewProject) SetCompany(v ViewRelationship)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ViewProject) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCompanyId

`func (o *ViewProject) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ViewProject) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ViewProject) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *ViewProject) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ViewProject) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ViewProject) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ViewProject) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ViewProject) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ViewProject) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ViewProject) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ViewProject) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ViewProject) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCustomFieldValueIds

`func (o *ViewProject) GetCustomFieldValueIds() []int32`

GetCustomFieldValueIds returns the CustomFieldValueIds field if non-nil, zero value otherwise.

### GetCustomFieldValueIdsOk

`func (o *ViewProject) GetCustomFieldValueIdsOk() (*[]int32, bool)`

GetCustomFieldValueIdsOk returns a tuple with the CustomFieldValueIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldValueIds

`func (o *ViewProject) SetCustomFieldValueIds(v []int32)`

SetCustomFieldValueIds sets CustomFieldValueIds field to given value.

### HasCustomFieldValueIds

`func (o *ViewProject) HasCustomFieldValueIds() bool`

HasCustomFieldValueIds returns a boolean if a field has been set.

### GetCustomFieldValues

`func (o *ViewProject) GetCustomFieldValues() []ViewRelationship`

GetCustomFieldValues returns the CustomFieldValues field if non-nil, zero value otherwise.

### GetCustomFieldValuesOk

`func (o *ViewProject) GetCustomFieldValuesOk() (*[]ViewRelationship, bool)`

GetCustomFieldValuesOk returns a tuple with the CustomFieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldValues

`func (o *ViewProject) SetCustomFieldValues(v []ViewRelationship)`

SetCustomFieldValues sets CustomFieldValues field to given value.

### HasCustomFieldValues

`func (o *ViewProject) HasCustomFieldValues() bool`

HasCustomFieldValues returns a boolean if a field has been set.

### GetDefaultPrivacy

`func (o *ViewProject) GetDefaultPrivacy() string`

GetDefaultPrivacy returns the DefaultPrivacy field if non-nil, zero value otherwise.

### GetDefaultPrivacyOk

`func (o *ViewProject) GetDefaultPrivacyOk() (*string, bool)`

GetDefaultPrivacyOk returns a tuple with the DefaultPrivacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPrivacy

`func (o *ViewProject) SetDefaultPrivacy(v string)`

SetDefaultPrivacy sets DefaultPrivacy field to given value.

### HasDefaultPrivacy

`func (o *ViewProject) HasDefaultPrivacy() bool`

HasDefaultPrivacy returns a boolean if a field has been set.

### GetDescription

`func (o *ViewProject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ViewProject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ViewProject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ViewProject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDirectFileUploadsEnabled

`func (o *ViewProject) GetDirectFileUploadsEnabled() bool`

GetDirectFileUploadsEnabled returns the DirectFileUploadsEnabled field if non-nil, zero value otherwise.

### GetDirectFileUploadsEnabledOk

`func (o *ViewProject) GetDirectFileUploadsEnabledOk() (*bool, bool)`

GetDirectFileUploadsEnabledOk returns a tuple with the DirectFileUploadsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectFileUploadsEnabled

`func (o *ViewProject) SetDirectFileUploadsEnabled(v bool)`

SetDirectFileUploadsEnabled sets DirectFileUploadsEnabled field to given value.

### HasDirectFileUploadsEnabled

`func (o *ViewProject) HasDirectFileUploadsEnabled() bool`

HasDirectFileUploadsEnabled returns a boolean if a field has been set.

### GetEndDate

`func (o *ViewProject) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ViewProject) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ViewProject) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ViewProject) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetFinancialBudget

`func (o *ViewProject) GetFinancialBudget() ViewRelationship`

GetFinancialBudget returns the FinancialBudget field if non-nil, zero value otherwise.

### GetFinancialBudgetOk

`func (o *ViewProject) GetFinancialBudgetOk() (*ViewRelationship, bool)`

GetFinancialBudgetOk returns a tuple with the FinancialBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialBudget

`func (o *ViewProject) SetFinancialBudget(v ViewRelationship)`

SetFinancialBudget sets FinancialBudget field to given value.

### HasFinancialBudget

`func (o *ViewProject) HasFinancialBudget() bool`

HasFinancialBudget returns a boolean if a field has been set.

### GetFinancialBudgetId

`func (o *ViewProject) GetFinancialBudgetId() int32`

GetFinancialBudgetId returns the FinancialBudgetId field if non-nil, zero value otherwise.

### GetFinancialBudgetIdOk

`func (o *ViewProject) GetFinancialBudgetIdOk() (*int32, bool)`

GetFinancialBudgetIdOk returns a tuple with the FinancialBudgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialBudgetId

`func (o *ViewProject) SetFinancialBudgetId(v int32)`

SetFinancialBudgetId sets FinancialBudgetId field to given value.

### HasFinancialBudgetId

`func (o *ViewProject) HasFinancialBudgetId() bool`

HasFinancialBudgetId returns a boolean if a field has been set.

### GetHarvestTimersEnabled

`func (o *ViewProject) GetHarvestTimersEnabled() bool`

GetHarvestTimersEnabled returns the HarvestTimersEnabled field if non-nil, zero value otherwise.

### GetHarvestTimersEnabledOk

`func (o *ViewProject) GetHarvestTimersEnabledOk() (*bool, bool)`

GetHarvestTimersEnabledOk returns a tuple with the HarvestTimersEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHarvestTimersEnabled

`func (o *ViewProject) SetHarvestTimersEnabled(v bool)`

SetHarvestTimersEnabled sets HarvestTimersEnabled field to given value.

### HasHarvestTimersEnabled

`func (o *ViewProject) HasHarvestTimersEnabled() bool`

HasHarvestTimersEnabled returns a boolean if a field has been set.

### GetId

`func (o *ViewProject) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewProject) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewProject) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewProject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIntegrations

`func (o *ViewProject) GetIntegrations() ViewProjectIntegrations`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *ViewProject) GetIntegrationsOk() (*ViewProjectIntegrations, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *ViewProject) SetIntegrations(v ViewProjectIntegrations)`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *ViewProject) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.

### GetIsOnBoardingProject

`func (o *ViewProject) GetIsOnBoardingProject() bool`

GetIsOnBoardingProject returns the IsOnBoardingProject field if non-nil, zero value otherwise.

### GetIsOnBoardingProjectOk

`func (o *ViewProject) GetIsOnBoardingProjectOk() (*bool, bool)`

GetIsOnBoardingProjectOk returns a tuple with the IsOnBoardingProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnBoardingProject

`func (o *ViewProject) SetIsOnBoardingProject(v bool)`

SetIsOnBoardingProject sets IsOnBoardingProject field to given value.

### HasIsOnBoardingProject

`func (o *ViewProject) HasIsOnBoardingProject() bool`

HasIsOnBoardingProject returns a boolean if a field has been set.

### GetIsSampleProject

`func (o *ViewProject) GetIsSampleProject() bool`

GetIsSampleProject returns the IsSampleProject field if non-nil, zero value otherwise.

### GetIsSampleProjectOk

`func (o *ViewProject) GetIsSampleProjectOk() (*bool, bool)`

GetIsSampleProjectOk returns a tuple with the IsSampleProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSampleProject

`func (o *ViewProject) SetIsSampleProject(v bool)`

SetIsSampleProject sets IsSampleProject field to given value.

### HasIsSampleProject

`func (o *ViewProject) HasIsSampleProject() bool`

HasIsSampleProject returns a boolean if a field has been set.

### GetIsStarred

`func (o *ViewProject) GetIsStarred() bool`

GetIsStarred returns the IsStarred field if non-nil, zero value otherwise.

### GetIsStarredOk

`func (o *ViewProject) GetIsStarredOk() (*bool, bool)`

GetIsStarredOk returns a tuple with the IsStarred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStarred

`func (o *ViewProject) SetIsStarred(v bool)`

SetIsStarred sets IsStarred field to given value.

### HasIsStarred

`func (o *ViewProject) HasIsStarred() bool`

HasIsStarred returns a boolean if a field has been set.

### GetLogo

`func (o *ViewProject) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *ViewProject) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *ViewProject) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *ViewProject) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetMinMaxAvailableDates

`func (o *ViewProject) GetMinMaxAvailableDates() ViewProjectMinMaxAvailableDates`

GetMinMaxAvailableDates returns the MinMaxAvailableDates field if non-nil, zero value otherwise.

### GetMinMaxAvailableDatesOk

`func (o *ViewProject) GetMinMaxAvailableDatesOk() (*ViewProjectMinMaxAvailableDates, bool)`

GetMinMaxAvailableDatesOk returns a tuple with the MinMaxAvailableDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMaxAvailableDates

`func (o *ViewProject) SetMinMaxAvailableDates(v ViewProjectMinMaxAvailableDates)`

SetMinMaxAvailableDates sets MinMaxAvailableDates field to given value.

### HasMinMaxAvailableDates

`func (o *ViewProject) HasMinMaxAvailableDates() bool`

HasMinMaxAvailableDates returns a boolean if a field has been set.

### GetName

`func (o *ViewProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ViewProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ViewProject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ViewProject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotifyCommentIncludeCreator

`func (o *ViewProject) GetNotifyCommentIncludeCreator() bool`

GetNotifyCommentIncludeCreator returns the NotifyCommentIncludeCreator field if non-nil, zero value otherwise.

### GetNotifyCommentIncludeCreatorOk

`func (o *ViewProject) GetNotifyCommentIncludeCreatorOk() (*bool, bool)`

GetNotifyCommentIncludeCreatorOk returns a tuple with the NotifyCommentIncludeCreator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyCommentIncludeCreator

`func (o *ViewProject) SetNotifyCommentIncludeCreator(v bool)`

SetNotifyCommentIncludeCreator sets NotifyCommentIncludeCreator field to given value.

### HasNotifyCommentIncludeCreator

`func (o *ViewProject) HasNotifyCommentIncludeCreator() bool`

HasNotifyCommentIncludeCreator returns a boolean if a field has been set.

### GetNotifyEveryone

`func (o *ViewProject) GetNotifyEveryone() bool`

GetNotifyEveryone returns the NotifyEveryone field if non-nil, zero value otherwise.

### GetNotifyEveryoneOk

`func (o *ViewProject) GetNotifyEveryoneOk() (*bool, bool)`

GetNotifyEveryoneOk returns a tuple with the NotifyEveryone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyEveryone

`func (o *ViewProject) SetNotifyEveryone(v bool)`

SetNotifyEveryone sets NotifyEveryone field to given value.

### HasNotifyEveryone

`func (o *ViewProject) HasNotifyEveryone() bool`

HasNotifyEveryone returns a boolean if a field has been set.

### GetOverviewStartPage

`func (o *ViewProject) GetOverviewStartPage() string`

GetOverviewStartPage returns the OverviewStartPage field if non-nil, zero value otherwise.

### GetOverviewStartPageOk

`func (o *ViewProject) GetOverviewStartPageOk() (*string, bool)`

GetOverviewStartPageOk returns a tuple with the OverviewStartPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverviewStartPage

`func (o *ViewProject) SetOverviewStartPage(v string)`

SetOverviewStartPage sets OverviewStartPage field to given value.

### HasOverviewStartPage

`func (o *ViewProject) HasOverviewStartPage() bool`

HasOverviewStartPage returns a boolean if a field has been set.

### GetOwnedBy

`func (o *ViewProject) GetOwnedBy() int32`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *ViewProject) GetOwnedByOk() (*int32, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *ViewProject) SetOwnedBy(v int32)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *ViewProject) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.

### GetOwnerId

`func (o *ViewProject) GetOwnerId() int32`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ViewProject) GetOwnerIdOk() (*int32, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ViewProject) SetOwnerId(v int32)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ViewProject) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetPrivacyEnabled

`func (o *ViewProject) GetPrivacyEnabled() bool`

GetPrivacyEnabled returns the PrivacyEnabled field if non-nil, zero value otherwise.

### GetPrivacyEnabledOk

`func (o *ViewProject) GetPrivacyEnabledOk() (*bool, bool)`

GetPrivacyEnabledOk returns a tuple with the PrivacyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyEnabled

`func (o *ViewProject) SetPrivacyEnabled(v bool)`

SetPrivacyEnabled sets PrivacyEnabled field to given value.

### HasPrivacyEnabled

`func (o *ViewProject) HasPrivacyEnabled() bool`

HasPrivacyEnabled returns a boolean if a field has been set.

### GetProjectOwnerId

`func (o *ViewProject) GetProjectOwnerId() int32`

GetProjectOwnerId returns the ProjectOwnerId field if non-nil, zero value otherwise.

### GetProjectOwnerIdOk

`func (o *ViewProject) GetProjectOwnerIdOk() (*int32, bool)`

GetProjectOwnerIdOk returns a tuple with the ProjectOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectOwnerId

`func (o *ViewProject) SetProjectOwnerId(v int32)`

SetProjectOwnerId sets ProjectOwnerId field to given value.

### HasProjectOwnerId

`func (o *ViewProject) HasProjectOwnerId() bool`

HasProjectOwnerId returns a boolean if a field has been set.

### GetReplyByEmailEnabled

`func (o *ViewProject) GetReplyByEmailEnabled() bool`

GetReplyByEmailEnabled returns the ReplyByEmailEnabled field if non-nil, zero value otherwise.

### GetReplyByEmailEnabledOk

`func (o *ViewProject) GetReplyByEmailEnabledOk() (*bool, bool)`

GetReplyByEmailEnabledOk returns a tuple with the ReplyByEmailEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyByEmailEnabled

`func (o *ViewProject) SetReplyByEmailEnabled(v bool)`

SetReplyByEmailEnabled sets ReplyByEmailEnabled field to given value.

### HasReplyByEmailEnabled

`func (o *ViewProject) HasReplyByEmailEnabled() bool`

HasReplyByEmailEnabled returns a boolean if a field has been set.

### GetShowAnnouncement

`func (o *ViewProject) GetShowAnnouncement() bool`

GetShowAnnouncement returns the ShowAnnouncement field if non-nil, zero value otherwise.

### GetShowAnnouncementOk

`func (o *ViewProject) GetShowAnnouncementOk() (*bool, bool)`

GetShowAnnouncementOk returns a tuple with the ShowAnnouncement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAnnouncement

`func (o *ViewProject) SetShowAnnouncement(v bool)`

SetShowAnnouncement sets ShowAnnouncement field to given value.

### HasShowAnnouncement

`func (o *ViewProject) HasShowAnnouncement() bool`

HasShowAnnouncement returns a boolean if a field has been set.

### GetSkipWeekends

`func (o *ViewProject) GetSkipWeekends() bool`

GetSkipWeekends returns the SkipWeekends field if non-nil, zero value otherwise.

### GetSkipWeekendsOk

`func (o *ViewProject) GetSkipWeekendsOk() (*bool, bool)`

GetSkipWeekendsOk returns a tuple with the SkipWeekends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipWeekends

`func (o *ViewProject) SetSkipWeekends(v bool)`

SetSkipWeekends sets SkipWeekends field to given value.

### HasSkipWeekends

`func (o *ViewProject) HasSkipWeekends() bool`

HasSkipWeekends returns a boolean if a field has been set.

### GetStartDate

`func (o *ViewProject) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ViewProject) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ViewProject) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ViewProject) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStartPage

`func (o *ViewProject) GetStartPage() string`

GetStartPage returns the StartPage field if non-nil, zero value otherwise.

### GetStartPageOk

`func (o *ViewProject) GetStartPageOk() (*string, bool)`

GetStartPageOk returns a tuple with the StartPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPage

`func (o *ViewProject) SetStartPage(v string)`

SetStartPage sets StartPage field to given value.

### HasStartPage

`func (o *ViewProject) HasStartPage() bool`

HasStartPage returns a boolean if a field has been set.

### GetStatus

`func (o *ViewProject) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ViewProject) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ViewProject) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ViewProject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubStatus

`func (o *ViewProject) GetSubStatus() string`

GetSubStatus returns the SubStatus field if non-nil, zero value otherwise.

### GetSubStatusOk

`func (o *ViewProject) GetSubStatusOk() (*string, bool)`

GetSubStatusOk returns a tuple with the SubStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStatus

`func (o *ViewProject) SetSubStatus(v string)`

SetSubStatus sets SubStatus field to given value.

### HasSubStatus

`func (o *ViewProject) HasSubStatus() bool`

HasSubStatus returns a boolean if a field has been set.

### GetTagIds

`func (o *ViewProject) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *ViewProject) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *ViewProject) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *ViewProject) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTags

`func (o *ViewProject) GetTags() []ViewRelationship`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ViewProject) GetTagsOk() (*[]ViewRelationship, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ViewProject) SetTags(v []ViewRelationship)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ViewProject) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTasksStartPage

`func (o *ViewProject) GetTasksStartPage() string`

GetTasksStartPage returns the TasksStartPage field if non-nil, zero value otherwise.

### GetTasksStartPageOk

`func (o *ViewProject) GetTasksStartPageOk() (*string, bool)`

GetTasksStartPageOk returns a tuple with the TasksStartPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksStartPage

`func (o *ViewProject) SetTasksStartPage(v string)`

SetTasksStartPage sets TasksStartPage field to given value.

### HasTasksStartPage

`func (o *ViewProject) HasTasksStartPage() bool`

HasTasksStartPage returns a boolean if a field has been set.

### GetTimeBudget

`func (o *ViewProject) GetTimeBudget() ViewRelationship`

GetTimeBudget returns the TimeBudget field if non-nil, zero value otherwise.

### GetTimeBudgetOk

`func (o *ViewProject) GetTimeBudgetOk() (*ViewRelationship, bool)`

GetTimeBudgetOk returns a tuple with the TimeBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBudget

`func (o *ViewProject) SetTimeBudget(v ViewRelationship)`

SetTimeBudget sets TimeBudget field to given value.

### HasTimeBudget

`func (o *ViewProject) HasTimeBudget() bool`

HasTimeBudget returns a boolean if a field has been set.

### GetTimeBudgetId

`func (o *ViewProject) GetTimeBudgetId() int32`

GetTimeBudgetId returns the TimeBudgetId field if non-nil, zero value otherwise.

### GetTimeBudgetIdOk

`func (o *ViewProject) GetTimeBudgetIdOk() (*int32, bool)`

GetTimeBudgetIdOk returns a tuple with the TimeBudgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBudgetId

`func (o *ViewProject) SetTimeBudgetId(v int32)`

SetTimeBudgetId sets TimeBudgetId field to given value.

### HasTimeBudgetId

`func (o *ViewProject) HasTimeBudgetId() bool`

HasTimeBudgetId returns a boolean if a field has been set.

### GetType

`func (o *ViewProject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ViewProject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ViewProject) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ViewProject) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdateId

`func (o *ViewProject) GetUpdateId() int32`

GetUpdateId returns the UpdateId field if non-nil, zero value otherwise.

### GetUpdateIdOk

`func (o *ViewProject) GetUpdateIdOk() (*int32, bool)`

GetUpdateIdOk returns a tuple with the UpdateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateId

`func (o *ViewProject) SetUpdateId(v int32)`

SetUpdateId sets UpdateId field to given value.

### HasUpdateId

`func (o *ViewProject) HasUpdateId() bool`

HasUpdateId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ViewProject) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ViewProject) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ViewProject) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ViewProject) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ViewProject) GetUpdatedBy() int32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ViewProject) GetUpdatedByOk() (*int32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ViewProject) SetUpdatedBy(v int32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ViewProject) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


