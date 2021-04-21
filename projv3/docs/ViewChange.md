# ViewChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**LastUpdatedDate** | Pointer to **string** |  | [optional] 
**Parent** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**ParentId** | Pointer to **int32** |  | [optional] 

## Methods

### NewViewChange

`func NewViewChange() *ViewChange`

NewViewChange instantiates a new ViewChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewChangeWithDefaults

`func NewViewChangeWithDefaults() *ViewChange`

NewViewChangeWithDefaults instantiates a new ViewChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ViewChange) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewChange) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewChange) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewChange) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedDate

`func (o *ViewChange) GetLastUpdatedDate() string`

GetLastUpdatedDate returns the LastUpdatedDate field if non-nil, zero value otherwise.

### GetLastUpdatedDateOk

`func (o *ViewChange) GetLastUpdatedDateOk() (*string, bool)`

GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDate

`func (o *ViewChange) SetLastUpdatedDate(v string)`

SetLastUpdatedDate sets LastUpdatedDate field to given value.

### HasLastUpdatedDate

`func (o *ViewChange) HasLastUpdatedDate() bool`

HasLastUpdatedDate returns a boolean if a field has been set.

### GetParent

`func (o *ViewChange) GetParent() ViewRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ViewChange) GetParentOk() (*ViewRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ViewChange) SetParent(v ViewRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ViewChange) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetParentId

`func (o *ViewChange) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ViewChange) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ViewChange) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ViewChange) HasParentId() bool`

HasParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


