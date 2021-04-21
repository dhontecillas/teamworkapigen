# ColumnResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Column** | Pointer to [**ViewBoardColumn**](ViewBoardColumn.md) |  | [optional] 
**Included** | Pointer to [**ColumnResponseIncluded**](ColumnResponseIncluded.md) |  | [optional] 
**Meta** | Pointer to [**ColumnBoardColumnMeta**](ColumnBoardColumnMeta.md) |  | [optional] 

## Methods

### NewColumnResponse

`func NewColumnResponse() *ColumnResponse`

NewColumnResponse instantiates a new ColumnResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColumnResponseWithDefaults

`func NewColumnResponseWithDefaults() *ColumnResponse`

NewColumnResponseWithDefaults instantiates a new ColumnResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumn

`func (o *ColumnResponse) GetColumn() ViewBoardColumn`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *ColumnResponse) GetColumnOk() (*ViewBoardColumn, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *ColumnResponse) SetColumn(v ViewBoardColumn)`

SetColumn sets Column field to given value.

### HasColumn

`func (o *ColumnResponse) HasColumn() bool`

HasColumn returns a boolean if a field has been set.

### GetIncluded

`func (o *ColumnResponse) GetIncluded() ColumnResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *ColumnResponse) GetIncludedOk() (*ColumnResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *ColumnResponse) SetIncluded(v ColumnResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *ColumnResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *ColumnResponse) GetMeta() ColumnBoardColumnMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ColumnResponse) GetMetaOk() (*ColumnBoardColumnMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ColumnResponse) SetMeta(v ColumnBoardColumnMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ColumnResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


