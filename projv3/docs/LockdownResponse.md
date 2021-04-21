# LockdownResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Included** | Pointer to [**LockdownResponseIncluded**](LockdownResponseIncluded.md) |  | [optional] 
**Lockdown** | Pointer to [**ViewLockdown**](ViewLockdown.md) |  | [optional] 

## Methods

### NewLockdownResponse

`func NewLockdownResponse() *LockdownResponse`

NewLockdownResponse instantiates a new LockdownResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLockdownResponseWithDefaults

`func NewLockdownResponseWithDefaults() *LockdownResponse`

NewLockdownResponseWithDefaults instantiates a new LockdownResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncluded

`func (o *LockdownResponse) GetIncluded() LockdownResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *LockdownResponse) GetIncludedOk() (*LockdownResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *LockdownResponse) SetIncluded(v LockdownResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *LockdownResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLockdown

`func (o *LockdownResponse) GetLockdown() ViewLockdown`

GetLockdown returns the Lockdown field if non-nil, zero value otherwise.

### GetLockdownOk

`func (o *LockdownResponse) GetLockdownOk() (*ViewLockdown, bool)`

GetLockdownOk returns a tuple with the Lockdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockdown

`func (o *LockdownResponse) SetLockdown(v ViewLockdown)`

SetLockdown sets Lockdown field to given value.

### HasLockdown

`func (o *LockdownResponse) HasLockdown() bool`

HasLockdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


