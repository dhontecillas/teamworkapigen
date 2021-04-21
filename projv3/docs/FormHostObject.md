# FormHostObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Meta** | Pointer to [**FormHostObjectMeta**](FormHostObjectMeta.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewFormHostObject

`func NewFormHostObject() *FormHostObject`

NewFormHostObject instantiates a new FormHostObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormHostObjectWithDefaults

`func NewFormHostObjectWithDefaults() *FormHostObject`

NewFormHostObjectWithDefaults instantiates a new FormHostObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FormHostObject) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormHostObject) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormHostObject) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FormHostObject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMeta

`func (o *FormHostObject) GetMeta() FormHostObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FormHostObject) GetMetaOk() (*FormHostObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FormHostObject) SetMeta(v FormHostObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FormHostObject) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetType

`func (o *FormHostObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormHostObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormHostObject) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FormHostObject) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


