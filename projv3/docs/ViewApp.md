# ViewApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Appsettings** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) | Settings is an optional list of config settings for the App can be used to build up an options panel | [optional] 
**BaseUrl** | Pointer to **string** | BaseURL is the CDN location for the latest version of the App | [optional] 
**Code** | Pointer to **string** | Code is a text unique identifier slug for the App | [optional] 
**CreatedAt** | Pointer to **string** | CreatedAt is the date the app was created | [optional] 
**Description** | Pointer to **string** | Description is a detailed description of the App | [optional] 
**Id** | Pointer to **int32** | ID is the unique identifier for the App | [optional] 
**IsInstalled** | Pointer to **bool** | IsInstalled denotes if the app can be used by the installation | [optional] 
**IsUninstallable** | Pointer to **bool** | IsUninstallable denotes if the app can be removed, some apps are added by default and can&#39;t be removed. | [optional] 
**IsVerified** | Pointer to **bool** | IsVerified denotes if the App has been verified by Teamwork for usage | [optional] 
**LogoURL** | Pointer to **string** | LogoURL is the URL to the Apps logo if it has one | [optional] 
**Name** | Pointer to **string** | name is the human readable name | [optional] 
**Publisher** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**ShortDescription** | Pointer to **string** | ShortDescription is a snippet of information detailing the app | [optional] 
**UpdatedAt** | Pointer to **string** | CreatedAt is the date the app was last updated | [optional] 
**Url** | Pointer to **string** | URL is the location of the App feature page on the marketplace | [optional] 
**VerifiedAt** | Pointer to **string** | VerifiedAt, if provided, is the date the App was verified by Teamwork | [optional] 

## Methods

### NewViewApp

`func NewViewApp() *ViewApp`

NewViewApp instantiates a new ViewApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewAppWithDefaults

`func NewViewAppWithDefaults() *ViewApp`

NewViewAppWithDefaults instantiates a new ViewApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppsettings

`func (o *ViewApp) GetAppsettings() []ViewRelationship`

GetAppsettings returns the Appsettings field if non-nil, zero value otherwise.

### GetAppsettingsOk

`func (o *ViewApp) GetAppsettingsOk() (*[]ViewRelationship, bool)`

GetAppsettingsOk returns a tuple with the Appsettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppsettings

`func (o *ViewApp) SetAppsettings(v []ViewRelationship)`

SetAppsettings sets Appsettings field to given value.

### HasAppsettings

`func (o *ViewApp) HasAppsettings() bool`

HasAppsettings returns a boolean if a field has been set.

### GetBaseUrl

`func (o *ViewApp) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *ViewApp) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *ViewApp) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *ViewApp) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetCode

`func (o *ViewApp) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ViewApp) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ViewApp) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ViewApp) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ViewApp) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ViewApp) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ViewApp) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ViewApp) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *ViewApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ViewApp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ViewApp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ViewApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ViewApp) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewApp) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewApp) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewApp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsInstalled

`func (o *ViewApp) GetIsInstalled() bool`

GetIsInstalled returns the IsInstalled field if non-nil, zero value otherwise.

### GetIsInstalledOk

`func (o *ViewApp) GetIsInstalledOk() (*bool, bool)`

GetIsInstalledOk returns a tuple with the IsInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInstalled

`func (o *ViewApp) SetIsInstalled(v bool)`

SetIsInstalled sets IsInstalled field to given value.

### HasIsInstalled

`func (o *ViewApp) HasIsInstalled() bool`

HasIsInstalled returns a boolean if a field has been set.

### GetIsUninstallable

`func (o *ViewApp) GetIsUninstallable() bool`

GetIsUninstallable returns the IsUninstallable field if non-nil, zero value otherwise.

### GetIsUninstallableOk

`func (o *ViewApp) GetIsUninstallableOk() (*bool, bool)`

GetIsUninstallableOk returns a tuple with the IsUninstallable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUninstallable

`func (o *ViewApp) SetIsUninstallable(v bool)`

SetIsUninstallable sets IsUninstallable field to given value.

### HasIsUninstallable

`func (o *ViewApp) HasIsUninstallable() bool`

HasIsUninstallable returns a boolean if a field has been set.

### GetIsVerified

`func (o *ViewApp) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *ViewApp) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *ViewApp) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *ViewApp) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### GetLogoURL

`func (o *ViewApp) GetLogoURL() string`

GetLogoURL returns the LogoURL field if non-nil, zero value otherwise.

### GetLogoURLOk

`func (o *ViewApp) GetLogoURLOk() (*string, bool)`

GetLogoURLOk returns a tuple with the LogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoURL

`func (o *ViewApp) SetLogoURL(v string)`

SetLogoURL sets LogoURL field to given value.

### HasLogoURL

`func (o *ViewApp) HasLogoURL() bool`

HasLogoURL returns a boolean if a field has been set.

### GetName

`func (o *ViewApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ViewApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ViewApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ViewApp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublisher

`func (o *ViewApp) GetPublisher() ViewRelationship`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *ViewApp) GetPublisherOk() (*ViewRelationship, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *ViewApp) SetPublisher(v ViewRelationship)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *ViewApp) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### GetShortDescription

`func (o *ViewApp) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *ViewApp) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *ViewApp) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *ViewApp) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ViewApp) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ViewApp) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ViewApp) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ViewApp) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *ViewApp) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ViewApp) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ViewApp) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ViewApp) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVerifiedAt

`func (o *ViewApp) GetVerifiedAt() string`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *ViewApp) GetVerifiedAtOk() (*string, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *ViewApp) SetVerifiedAt(v string)`

SetVerifiedAt sets VerifiedAt field to given value.

### HasVerifiedAt

`func (o *ViewApp) HasVerifiedAt() bool`

HasVerifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


