# SummaryTimeCounterResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Estimates** | Pointer to [**SummaryTimeCounterEstimateResponse**](SummaryTimeCounterEstimateResponse.md) |  | [optional] 
**Totals** | Pointer to [**SummaryTimeCounterTotalResponse**](SummaryTimeCounterTotalResponse.md) |  | [optional] 

## Methods

### NewSummaryTimeCounterResponse

`func NewSummaryTimeCounterResponse() *SummaryTimeCounterResponse`

NewSummaryTimeCounterResponse instantiates a new SummaryTimeCounterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryTimeCounterResponseWithDefaults

`func NewSummaryTimeCounterResponseWithDefaults() *SummaryTimeCounterResponse`

NewSummaryTimeCounterResponseWithDefaults instantiates a new SummaryTimeCounterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEstimates

`func (o *SummaryTimeCounterResponse) GetEstimates() SummaryTimeCounterEstimateResponse`

GetEstimates returns the Estimates field if non-nil, zero value otherwise.

### GetEstimatesOk

`func (o *SummaryTimeCounterResponse) GetEstimatesOk() (*SummaryTimeCounterEstimateResponse, bool)`

GetEstimatesOk returns a tuple with the Estimates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimates

`func (o *SummaryTimeCounterResponse) SetEstimates(v SummaryTimeCounterEstimateResponse)`

SetEstimates sets Estimates field to given value.

### HasEstimates

`func (o *SummaryTimeCounterResponse) HasEstimates() bool`

HasEstimates returns a boolean if a field has been set.

### GetTotals

`func (o *SummaryTimeCounterResponse) GetTotals() SummaryTimeCounterTotalResponse`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *SummaryTimeCounterResponse) GetTotalsOk() (*SummaryTimeCounterTotalResponse, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *SummaryTimeCounterResponse) SetTotals(v SummaryTimeCounterTotalResponse)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *SummaryTimeCounterResponse) HasTotals() bool`

HasTotals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


