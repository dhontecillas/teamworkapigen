# ViewAppSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | Pointer to **string** | DefaultValue is the value returned if Value is empty. | [optional] 
**DisplayName** | Pointer to **string** | DisplayName is what the user should see when presented with the setting name. | [optional] 
**Id** | Pointer to **int32** | ID is the unique identifier for the setting. | [optional] 
**IsSecret** | Pointer to **bool** | IsSecret if true the DefaultValue and Value fields will be redacted. | [optional] 
**Name** | Pointer to **string** | Name is the code name for the setting - usually not displayed to a user. | [optional] 
**Type** | Pointer to **string** | Type. | [optional] 
**Value** | Pointer to **string** | Value is the current value for the setting - if empty DefaultValue is used. | [optional] 

## Methods

### NewViewAppSetting

`func NewViewAppSetting() *ViewAppSetting`

NewViewAppSetting instantiates a new ViewAppSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewAppSettingWithDefaults

`func NewViewAppSettingWithDefaults() *ViewAppSetting`

NewViewAppSettingWithDefaults instantiates a new ViewAppSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValue

`func (o *ViewAppSetting) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ViewAppSetting) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ViewAppSetting) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ViewAppSetting) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetDisplayName

`func (o *ViewAppSetting) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ViewAppSetting) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ViewAppSetting) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ViewAppSetting) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *ViewAppSetting) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewAppSetting) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewAppSetting) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewAppSetting) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsSecret

`func (o *ViewAppSetting) GetIsSecret() bool`

GetIsSecret returns the IsSecret field if non-nil, zero value otherwise.

### GetIsSecretOk

`func (o *ViewAppSetting) GetIsSecretOk() (*bool, bool)`

GetIsSecretOk returns a tuple with the IsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecret

`func (o *ViewAppSetting) SetIsSecret(v bool)`

SetIsSecret sets IsSecret field to given value.

### HasIsSecret

`func (o *ViewAppSetting) HasIsSecret() bool`

HasIsSecret returns a boolean if a field has been set.

### GetName

`func (o *ViewAppSetting) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ViewAppSetting) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ViewAppSetting) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ViewAppSetting) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ViewAppSetting) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ViewAppSetting) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ViewAppSetting) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ViewAppSetting) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *ViewAppSetting) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ViewAppSetting) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ViewAppSetting) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ViewAppSetting) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


