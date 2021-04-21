# CustomfieldResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Customfield** | Pointer to [**ViewCustomField**](ViewCustomField.md) |  | [optional] 

## Methods

### NewCustomfieldResponse

`func NewCustomfieldResponse() *CustomfieldResponse`

NewCustomfieldResponse instantiates a new CustomfieldResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomfieldResponseWithDefaults

`func NewCustomfieldResponseWithDefaults() *CustomfieldResponse`

NewCustomfieldResponseWithDefaults instantiates a new CustomfieldResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *CustomfieldResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CustomfieldResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CustomfieldResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CustomfieldResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCustomfield

`func (o *CustomfieldResponse) GetCustomfield() ViewCustomField`

GetCustomfield returns the Customfield field if non-nil, zero value otherwise.

### GetCustomfieldOk

`func (o *CustomfieldResponse) GetCustomfieldOk() (*ViewCustomField, bool)`

GetCustomfieldOk returns a tuple with the Customfield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomfield

`func (o *CustomfieldResponse) SetCustomfield(v ViewCustomField)`

SetCustomfield sets Customfield field to given value.

### HasCustomfield

`func (o *CustomfieldResponse) HasCustomfield() bool`

HasCustomfield returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


