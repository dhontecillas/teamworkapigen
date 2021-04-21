# ViewWorkingHour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **string** |  | [optional] 
**DateUpdated** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Entries** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**EntryIds** | Pointer to **[]int32** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Object** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**ObjectId** | Pointer to **int32** |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewViewWorkingHour

`func NewViewWorkingHour() *ViewWorkingHour`

NewViewWorkingHour instantiates a new ViewWorkingHour object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewWorkingHourWithDefaults

`func NewViewWorkingHourWithDefaults() *ViewWorkingHour`

NewViewWorkingHourWithDefaults instantiates a new ViewWorkingHour object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ViewWorkingHour) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ViewWorkingHour) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ViewWorkingHour) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ViewWorkingHour) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDateCreated

`func (o *ViewWorkingHour) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ViewWorkingHour) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ViewWorkingHour) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ViewWorkingHour) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateUpdated

`func (o *ViewWorkingHour) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ViewWorkingHour) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ViewWorkingHour) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ViewWorkingHour) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### GetDescription

`func (o *ViewWorkingHour) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ViewWorkingHour) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ViewWorkingHour) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ViewWorkingHour) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntries

`func (o *ViewWorkingHour) GetEntries() []ViewRelationship`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *ViewWorkingHour) GetEntriesOk() (*[]ViewRelationship, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *ViewWorkingHour) SetEntries(v []ViewRelationship)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *ViewWorkingHour) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetEntryIds

`func (o *ViewWorkingHour) GetEntryIds() []int32`

GetEntryIds returns the EntryIds field if non-nil, zero value otherwise.

### GetEntryIdsOk

`func (o *ViewWorkingHour) GetEntryIdsOk() (*[]int32, bool)`

GetEntryIdsOk returns a tuple with the EntryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryIds

`func (o *ViewWorkingHour) SetEntryIds(v []int32)`

SetEntryIds sets EntryIds field to given value.

### HasEntryIds

`func (o *ViewWorkingHour) HasEntryIds() bool`

HasEntryIds returns a boolean if a field has been set.

### GetId

`func (o *ViewWorkingHour) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewWorkingHour) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewWorkingHour) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewWorkingHour) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ViewWorkingHour) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ViewWorkingHour) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ViewWorkingHour) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ViewWorkingHour) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObject

`func (o *ViewWorkingHour) GetObject() ViewRelationship`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ViewWorkingHour) GetObjectOk() (*ViewRelationship, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ViewWorkingHour) SetObject(v ViewRelationship)`

SetObject sets Object field to given value.

### HasObject

`func (o *ViewWorkingHour) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetObjectId

`func (o *ViewWorkingHour) GetObjectId() int32`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *ViewWorkingHour) GetObjectIdOk() (*int32, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *ViewWorkingHour) SetObjectId(v int32)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *ViewWorkingHour) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetObjectType

`func (o *ViewWorkingHour) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ViewWorkingHour) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ViewWorkingHour) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *ViewWorkingHour) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ViewWorkingHour) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ViewWorkingHour) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ViewWorkingHour) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ViewWorkingHour) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


