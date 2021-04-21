# PayloadNotify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to [**PayloadUserGroups**](PayloadUserGroups.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewPayloadNotify

`func NewPayloadNotify() *PayloadNotify`

NewPayloadNotify instantiates a new PayloadNotify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayloadNotifyWithDefaults

`func NewPayloadNotifyWithDefaults() *PayloadNotify`

NewPayloadNotifyWithDefaults instantiates a new PayloadNotify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *PayloadNotify) GetIds() PayloadUserGroups`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *PayloadNotify) GetIdsOk() (*PayloadUserGroups, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *PayloadNotify) SetIds(v PayloadUserGroups)`

SetIds sets Ids field to given value.

### HasIds

`func (o *PayloadNotify) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetType

`func (o *PayloadNotify) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PayloadNotify) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PayloadNotify) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PayloadNotify) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


