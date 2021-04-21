# SummaryColumnDataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cards** | Pointer to [**SummaryColumnCardResponse**](SummaryColumnCardResponse.md) |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**EstimatedTime** | Pointer to [**SummaryColumnEstimatedResponse**](SummaryColumnEstimatedResponse.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewSummaryColumnDataResponse

`func NewSummaryColumnDataResponse() *SummaryColumnDataResponse`

NewSummaryColumnDataResponse instantiates a new SummaryColumnDataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryColumnDataResponseWithDefaults

`func NewSummaryColumnDataResponseWithDefaults() *SummaryColumnDataResponse`

NewSummaryColumnDataResponseWithDefaults instantiates a new SummaryColumnDataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCards

`func (o *SummaryColumnDataResponse) GetCards() SummaryColumnCardResponse`

GetCards returns the Cards field if non-nil, zero value otherwise.

### GetCardsOk

`func (o *SummaryColumnDataResponse) GetCardsOk() (*SummaryColumnCardResponse, bool)`

GetCardsOk returns a tuple with the Cards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCards

`func (o *SummaryColumnDataResponse) SetCards(v SummaryColumnCardResponse)`

SetCards sets Cards field to given value.

### HasCards

`func (o *SummaryColumnDataResponse) HasCards() bool`

HasCards returns a boolean if a field has been set.

### GetColor

`func (o *SummaryColumnDataResponse) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *SummaryColumnDataResponse) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *SummaryColumnDataResponse) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *SummaryColumnDataResponse) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetEstimatedTime

`func (o *SummaryColumnDataResponse) GetEstimatedTime() SummaryColumnEstimatedResponse`

GetEstimatedTime returns the EstimatedTime field if non-nil, zero value otherwise.

### GetEstimatedTimeOk

`func (o *SummaryColumnDataResponse) GetEstimatedTimeOk() (*SummaryColumnEstimatedResponse, bool)`

GetEstimatedTimeOk returns a tuple with the EstimatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTime

`func (o *SummaryColumnDataResponse) SetEstimatedTime(v SummaryColumnEstimatedResponse)`

SetEstimatedTime sets EstimatedTime field to given value.

### HasEstimatedTime

`func (o *SummaryColumnDataResponse) HasEstimatedTime() bool`

HasEstimatedTime returns a boolean if a field has been set.

### GetId

`func (o *SummaryColumnDataResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SummaryColumnDataResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SummaryColumnDataResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SummaryColumnDataResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SummaryColumnDataResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SummaryColumnDataResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SummaryColumnDataResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SummaryColumnDataResponse) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


