# ViewFormToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanonicalURL** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**DeletedBy** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**Expires** | Pointer to **bool** |  | [optional] 
**ExpiryDate** | Pointer to **string** |  | [optional] 
**Form** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Installation** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewViewFormToken

`func NewViewFormToken() *ViewFormToken`

NewViewFormToken instantiates a new ViewFormToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewFormTokenWithDefaults

`func NewViewFormTokenWithDefaults() *ViewFormToken`

NewViewFormTokenWithDefaults instantiates a new ViewFormToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanonicalURL

`func (o *ViewFormToken) GetCanonicalURL() string`

GetCanonicalURL returns the CanonicalURL field if non-nil, zero value otherwise.

### GetCanonicalURLOk

`func (o *ViewFormToken) GetCanonicalURLOk() (*string, bool)`

GetCanonicalURLOk returns a tuple with the CanonicalURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalURL

`func (o *ViewFormToken) SetCanonicalURL(v string)`

SetCanonicalURL sets CanonicalURL field to given value.

### HasCanonicalURL

`func (o *ViewFormToken) HasCanonicalURL() bool`

HasCanonicalURL returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ViewFormToken) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ViewFormToken) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ViewFormToken) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ViewFormToken) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ViewFormToken) GetCreatedBy() ViewRelationship`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ViewFormToken) GetCreatedByOk() (*ViewRelationship, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ViewFormToken) SetCreatedBy(v ViewRelationship)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ViewFormToken) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ViewFormToken) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ViewFormToken) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ViewFormToken) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ViewFormToken) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeletedBy

`func (o *ViewFormToken) GetDeletedBy() ViewRelationship`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *ViewFormToken) GetDeletedByOk() (*ViewRelationship, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *ViewFormToken) SetDeletedBy(v ViewRelationship)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *ViewFormToken) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetExpires

`func (o *ViewFormToken) GetExpires() bool`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *ViewFormToken) GetExpiresOk() (*bool, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *ViewFormToken) SetExpires(v bool)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *ViewFormToken) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetExpiryDate

`func (o *ViewFormToken) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *ViewFormToken) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *ViewFormToken) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *ViewFormToken) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetForm

`func (o *ViewFormToken) GetForm() ViewRelationship`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *ViewFormToken) GetFormOk() (*ViewRelationship, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *ViewFormToken) SetForm(v ViewRelationship)`

SetForm sets Form field to given value.

### HasForm

`func (o *ViewFormToken) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetId

`func (o *ViewFormToken) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewFormToken) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewFormToken) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewFormToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstallation

`func (o *ViewFormToken) GetInstallation() ViewRelationship`

GetInstallation returns the Installation field if non-nil, zero value otherwise.

### GetInstallationOk

`func (o *ViewFormToken) GetInstallationOk() (*ViewRelationship, bool)`

GetInstallationOk returns a tuple with the Installation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallation

`func (o *ViewFormToken) SetInstallation(v ViewRelationship)`

SetInstallation sets Installation field to given value.

### HasInstallation

`func (o *ViewFormToken) HasInstallation() bool`

HasInstallation returns a boolean if a field has been set.

### GetState

`func (o *ViewFormToken) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ViewFormToken) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ViewFormToken) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ViewFormToken) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ViewFormToken) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ViewFormToken) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ViewFormToken) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ViewFormToken) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ViewFormToken) GetUpdatedBy() ViewRelationship`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ViewFormToken) GetUpdatedByOk() (*ViewRelationship, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ViewFormToken) SetUpdatedBy(v ViewRelationship)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ViewFormToken) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetValue

`func (o *ViewFormToken) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ViewFormToken) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ViewFormToken) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ViewFormToken) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


