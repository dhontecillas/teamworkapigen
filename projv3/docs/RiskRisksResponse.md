# RiskRisksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Included** | Pointer to [**ActivityActivitiesResponseIncluded**](ActivityActivitiesResponseIncluded.md) |  | [optional] 
**Meta** | Pointer to [**ViewMeta**](ViewMeta.md) |  | [optional] 
**Risks** | Pointer to [**[]RiskRisk**](RiskRisk.md) |  | [optional] 

## Methods

### NewRiskRisksResponse

`func NewRiskRisksResponse() *RiskRisksResponse`

NewRiskRisksResponse instantiates a new RiskRisksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskRisksResponseWithDefaults

`func NewRiskRisksResponseWithDefaults() *RiskRisksResponse`

NewRiskRisksResponseWithDefaults instantiates a new RiskRisksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncluded

`func (o *RiskRisksResponse) GetIncluded() ActivityActivitiesResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *RiskRisksResponse) GetIncludedOk() (*ActivityActivitiesResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *RiskRisksResponse) SetIncluded(v ActivityActivitiesResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *RiskRisksResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *RiskRisksResponse) GetMeta() ViewMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RiskRisksResponse) GetMetaOk() (*ViewMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RiskRisksResponse) SetMeta(v ViewMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *RiskRisksResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetRisks

`func (o *RiskRisksResponse) GetRisks() []RiskRisk`

GetRisks returns the Risks field if non-nil, zero value otherwise.

### GetRisksOk

`func (o *RiskRisksResponse) GetRisksOk() (*[]RiskRisk, bool)`

GetRisksOk returns a tuple with the Risks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisks

`func (o *RiskRisksResponse) SetRisks(v []RiskRisk)`

SetRisks sets Risks field to given value.

### HasRisks

`func (o *RiskRisksResponse) HasRisks() bool`

HasRisks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


