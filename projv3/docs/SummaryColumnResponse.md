# SummaryColumnResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**[]SummaryColumnDataResponse**](SummaryColumnDataResponse.md) |  | [optional] 

## Methods

### NewSummaryColumnResponse

`func NewSummaryColumnResponse() *SummaryColumnResponse`

NewSummaryColumnResponse instantiates a new SummaryColumnResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryColumnResponseWithDefaults

`func NewSummaryColumnResponseWithDefaults() *SummaryColumnResponse`

NewSummaryColumnResponseWithDefaults instantiates a new SummaryColumnResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *SummaryColumnResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SummaryColumnResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SummaryColumnResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *SummaryColumnResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetData

`func (o *SummaryColumnResponse) GetData() []SummaryColumnDataResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SummaryColumnResponse) GetDataOk() (*[]SummaryColumnDataResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SummaryColumnResponse) SetData(v []SummaryColumnDataResponse)`

SetData sets Data field to given value.

### HasData

`func (o *SummaryColumnResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


