# CustomfieldCustomFieldsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customfields** | Pointer to [**[]ViewCustomField**](ViewCustomField.md) |  | [optional] 
**Included** | Pointer to [**CalendarEventsResponseIncluded**](CalendarEventsResponseIncluded.md) |  | [optional] 
**Meta** | Pointer to [**ViewMeta**](ViewMeta.md) |  | [optional] 

## Methods

### NewCustomfieldCustomFieldsResponse

`func NewCustomfieldCustomFieldsResponse() *CustomfieldCustomFieldsResponse`

NewCustomfieldCustomFieldsResponse instantiates a new CustomfieldCustomFieldsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomfieldCustomFieldsResponseWithDefaults

`func NewCustomfieldCustomFieldsResponseWithDefaults() *CustomfieldCustomFieldsResponse`

NewCustomfieldCustomFieldsResponseWithDefaults instantiates a new CustomfieldCustomFieldsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomfields

`func (o *CustomfieldCustomFieldsResponse) GetCustomfields() []ViewCustomField`

GetCustomfields returns the Customfields field if non-nil, zero value otherwise.

### GetCustomfieldsOk

`func (o *CustomfieldCustomFieldsResponse) GetCustomfieldsOk() (*[]ViewCustomField, bool)`

GetCustomfieldsOk returns a tuple with the Customfields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomfields

`func (o *CustomfieldCustomFieldsResponse) SetCustomfields(v []ViewCustomField)`

SetCustomfields sets Customfields field to given value.

### HasCustomfields

`func (o *CustomfieldCustomFieldsResponse) HasCustomfields() bool`

HasCustomfields returns a boolean if a field has been set.

### GetIncluded

`func (o *CustomfieldCustomFieldsResponse) GetIncluded() CalendarEventsResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *CustomfieldCustomFieldsResponse) GetIncludedOk() (*CalendarEventsResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *CustomfieldCustomFieldsResponse) SetIncluded(v CalendarEventsResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *CustomfieldCustomFieldsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *CustomfieldCustomFieldsResponse) GetMeta() ViewMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CustomfieldCustomFieldsResponse) GetMetaOk() (*ViewMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CustomfieldCustomFieldsResponse) SetMeta(v ViewMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CustomfieldCustomFieldsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


