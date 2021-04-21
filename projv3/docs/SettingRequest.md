# SettingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Appsetting** | Pointer to [**SettingAppSetting**](SettingAppSetting.md) |  | [optional] 

## Methods

### NewSettingRequest

`func NewSettingRequest() *SettingRequest`

NewSettingRequest instantiates a new SettingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingRequestWithDefaults

`func NewSettingRequestWithDefaults() *SettingRequest`

NewSettingRequestWithDefaults instantiates a new SettingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppsetting

`func (o *SettingRequest) GetAppsetting() SettingAppSetting`

GetAppsetting returns the Appsetting field if non-nil, zero value otherwise.

### GetAppsettingOk

`func (o *SettingRequest) GetAppsettingOk() (*SettingAppSetting, bool)`

GetAppsettingOk returns a tuple with the Appsetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppsetting

`func (o *SettingRequest) SetAppsetting(v SettingAppSetting)`

SetAppsetting sets Appsetting field to given value.

### HasAppsetting

`func (o *SettingRequest) HasAppsetting() bool`

HasAppsetting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


