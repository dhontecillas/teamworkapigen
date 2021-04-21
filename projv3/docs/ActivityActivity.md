# ActivityActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityType** | Pointer to **string** |  | [optional] 
**Company** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**CompanyId** | Pointer to **int32** |  | [optional] 
**DateTime** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DueDate** | Pointer to **string** |  | [optional] 
**ExtraDescription** | Pointer to **string** |  | [optional] 
**ForUser** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**ForUserId** | Pointer to **int32** |  | [optional] 
**ForUserName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**IsPrivate** | Pointer to **int32** |  | [optional] 
**Item** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**ItemId** | Pointer to **int32** |  | [optional] 
**ItemLink** | Pointer to **string** |  | [optional] 
**LatestActivityType** | Pointer to **string** |  | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**Lockdown** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**LockdownId** | Pointer to **int32** |  | [optional] 
**Project** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**PublicInfo** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 

## Methods

### NewActivityActivity

`func NewActivityActivity() *ActivityActivity`

NewActivityActivity instantiates a new ActivityActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityActivityWithDefaults

`func NewActivityActivityWithDefaults() *ActivityActivity`

NewActivityActivityWithDefaults instantiates a new ActivityActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityType

`func (o *ActivityActivity) GetActivityType() string`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *ActivityActivity) GetActivityTypeOk() (*string, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *ActivityActivity) SetActivityType(v string)`

SetActivityType sets ActivityType field to given value.

### HasActivityType

`func (o *ActivityActivity) HasActivityType() bool`

HasActivityType returns a boolean if a field has been set.

### GetCompany

`func (o *ActivityActivity) GetCompany() ViewRelationship`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ActivityActivity) GetCompanyOk() (*ViewRelationship, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ActivityActivity) SetCompany(v ViewRelationship)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ActivityActivity) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCompanyId

`func (o *ActivityActivity) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ActivityActivity) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ActivityActivity) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *ActivityActivity) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetDateTime

`func (o *ActivityActivity) GetDateTime() string`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *ActivityActivity) GetDateTimeOk() (*string, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *ActivityActivity) SetDateTime(v string)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *ActivityActivity) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetDescription

`func (o *ActivityActivity) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ActivityActivity) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ActivityActivity) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ActivityActivity) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDueDate

`func (o *ActivityActivity) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *ActivityActivity) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *ActivityActivity) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *ActivityActivity) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetExtraDescription

`func (o *ActivityActivity) GetExtraDescription() string`

GetExtraDescription returns the ExtraDescription field if non-nil, zero value otherwise.

### GetExtraDescriptionOk

`func (o *ActivityActivity) GetExtraDescriptionOk() (*string, bool)`

GetExtraDescriptionOk returns a tuple with the ExtraDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraDescription

`func (o *ActivityActivity) SetExtraDescription(v string)`

SetExtraDescription sets ExtraDescription field to given value.

### HasExtraDescription

`func (o *ActivityActivity) HasExtraDescription() bool`

HasExtraDescription returns a boolean if a field has been set.

### GetForUser

`func (o *ActivityActivity) GetForUser() ViewRelationship`

GetForUser returns the ForUser field if non-nil, zero value otherwise.

### GetForUserOk

`func (o *ActivityActivity) GetForUserOk() (*ViewRelationship, bool)`

GetForUserOk returns a tuple with the ForUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForUser

`func (o *ActivityActivity) SetForUser(v ViewRelationship)`

SetForUser sets ForUser field to given value.

### HasForUser

`func (o *ActivityActivity) HasForUser() bool`

HasForUser returns a boolean if a field has been set.

### GetForUserId

`func (o *ActivityActivity) GetForUserId() int32`

GetForUserId returns the ForUserId field if non-nil, zero value otherwise.

### GetForUserIdOk

`func (o *ActivityActivity) GetForUserIdOk() (*int32, bool)`

GetForUserIdOk returns a tuple with the ForUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForUserId

`func (o *ActivityActivity) SetForUserId(v int32)`

SetForUserId sets ForUserId field to given value.

### HasForUserId

`func (o *ActivityActivity) HasForUserId() bool`

HasForUserId returns a boolean if a field has been set.

### GetForUserName

`func (o *ActivityActivity) GetForUserName() string`

GetForUserName returns the ForUserName field if non-nil, zero value otherwise.

### GetForUserNameOk

`func (o *ActivityActivity) GetForUserNameOk() (*string, bool)`

GetForUserNameOk returns a tuple with the ForUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForUserName

`func (o *ActivityActivity) SetForUserName(v string)`

SetForUserName sets ForUserName field to given value.

### HasForUserName

`func (o *ActivityActivity) HasForUserName() bool`

HasForUserName returns a boolean if a field has been set.

### GetId

`func (o *ActivityActivity) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivityActivity) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivityActivity) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ActivityActivity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsPrivate

`func (o *ActivityActivity) GetIsPrivate() int32`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *ActivityActivity) GetIsPrivateOk() (*int32, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *ActivityActivity) SetIsPrivate(v int32)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *ActivityActivity) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetItem

`func (o *ActivityActivity) GetItem() ViewRelationship`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *ActivityActivity) GetItemOk() (*ViewRelationship, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *ActivityActivity) SetItem(v ViewRelationship)`

SetItem sets Item field to given value.

### HasItem

`func (o *ActivityActivity) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetItemId

`func (o *ActivityActivity) GetItemId() int32`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ActivityActivity) GetItemIdOk() (*int32, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ActivityActivity) SetItemId(v int32)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ActivityActivity) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetItemLink

`func (o *ActivityActivity) GetItemLink() string`

GetItemLink returns the ItemLink field if non-nil, zero value otherwise.

### GetItemLinkOk

`func (o *ActivityActivity) GetItemLinkOk() (*string, bool)`

GetItemLinkOk returns a tuple with the ItemLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLink

`func (o *ActivityActivity) SetItemLink(v string)`

SetItemLink sets ItemLink field to given value.

### HasItemLink

`func (o *ActivityActivity) HasItemLink() bool`

HasItemLink returns a boolean if a field has been set.

### GetLatestActivityType

`func (o *ActivityActivity) GetLatestActivityType() string`

GetLatestActivityType returns the LatestActivityType field if non-nil, zero value otherwise.

### GetLatestActivityTypeOk

`func (o *ActivityActivity) GetLatestActivityTypeOk() (*string, bool)`

GetLatestActivityTypeOk returns a tuple with the LatestActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestActivityType

`func (o *ActivityActivity) SetLatestActivityType(v string)`

SetLatestActivityType sets LatestActivityType field to given value.

### HasLatestActivityType

`func (o *ActivityActivity) HasLatestActivityType() bool`

HasLatestActivityType returns a boolean if a field has been set.

### GetLink

`func (o *ActivityActivity) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *ActivityActivity) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *ActivityActivity) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *ActivityActivity) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetLockdown

`func (o *ActivityActivity) GetLockdown() ViewRelationship`

GetLockdown returns the Lockdown field if non-nil, zero value otherwise.

### GetLockdownOk

`func (o *ActivityActivity) GetLockdownOk() (*ViewRelationship, bool)`

GetLockdownOk returns a tuple with the Lockdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockdown

`func (o *ActivityActivity) SetLockdown(v ViewRelationship)`

SetLockdown sets Lockdown field to given value.

### HasLockdown

`func (o *ActivityActivity) HasLockdown() bool`

HasLockdown returns a boolean if a field has been set.

### GetLockdownId

`func (o *ActivityActivity) GetLockdownId() int32`

GetLockdownId returns the LockdownId field if non-nil, zero value otherwise.

### GetLockdownIdOk

`func (o *ActivityActivity) GetLockdownIdOk() (*int32, bool)`

GetLockdownIdOk returns a tuple with the LockdownId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockdownId

`func (o *ActivityActivity) SetLockdownId(v int32)`

SetLockdownId sets LockdownId field to given value.

### HasLockdownId

`func (o *ActivityActivity) HasLockdownId() bool`

HasLockdownId returns a boolean if a field has been set.

### GetProject

`func (o *ActivityActivity) GetProject() ViewRelationship`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ActivityActivity) GetProjectOk() (*ViewRelationship, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ActivityActivity) SetProject(v ViewRelationship)`

SetProject sets Project field to given value.

### HasProject

`func (o *ActivityActivity) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectId

`func (o *ActivityActivity) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ActivityActivity) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ActivityActivity) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ActivityActivity) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetPublicInfo

`func (o *ActivityActivity) GetPublicInfo() string`

GetPublicInfo returns the PublicInfo field if non-nil, zero value otherwise.

### GetPublicInfoOk

`func (o *ActivityActivity) GetPublicInfoOk() (*string, bool)`

GetPublicInfoOk returns a tuple with the PublicInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicInfo

`func (o *ActivityActivity) SetPublicInfo(v string)`

SetPublicInfo sets PublicInfo field to given value.

### HasPublicInfo

`func (o *ActivityActivity) HasPublicInfo() bool`

HasPublicInfo returns a boolean if a field has been set.

### GetType

`func (o *ActivityActivity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActivityActivity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActivityActivity) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ActivityActivity) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUser

`func (o *ActivityActivity) GetUser() ViewRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ActivityActivity) GetUserOk() (*ViewRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ActivityActivity) SetUser(v ViewRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *ActivityActivity) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserId

`func (o *ActivityActivity) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ActivityActivity) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ActivityActivity) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ActivityActivity) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


