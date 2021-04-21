# CompanyResponseIncluded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Countries** | Pointer to [**map[string]ViewCountry**](ViewCountry.md) |  | [optional] 
**Industries** | Pointer to [**map[string]ViewIndustry**](ViewIndustry.md) |  | [optional] 
**Tags** | Pointer to [**map[string]ViewTag**](ViewTag.md) |  | [optional] 

## Methods

### NewCompanyResponseIncluded

`func NewCompanyResponseIncluded() *CompanyResponseIncluded`

NewCompanyResponseIncluded instantiates a new CompanyResponseIncluded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyResponseIncludedWithDefaults

`func NewCompanyResponseIncludedWithDefaults() *CompanyResponseIncluded`

NewCompanyResponseIncludedWithDefaults instantiates a new CompanyResponseIncluded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountries

`func (o *CompanyResponseIncluded) GetCountries() map[string]ViewCountry`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *CompanyResponseIncluded) GetCountriesOk() (*map[string]ViewCountry, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *CompanyResponseIncluded) SetCountries(v map[string]ViewCountry)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *CompanyResponseIncluded) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetIndustries

`func (o *CompanyResponseIncluded) GetIndustries() map[string]ViewIndustry`

GetIndustries returns the Industries field if non-nil, zero value otherwise.

### GetIndustriesOk

`func (o *CompanyResponseIncluded) GetIndustriesOk() (*map[string]ViewIndustry, bool)`

GetIndustriesOk returns a tuple with the Industries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustries

`func (o *CompanyResponseIncluded) SetIndustries(v map[string]ViewIndustry)`

SetIndustries sets Industries field to given value.

### HasIndustries

`func (o *CompanyResponseIncluded) HasIndustries() bool`

HasIndustries returns a boolean if a field has been set.

### GetTags

`func (o *CompanyResponseIncluded) GetTags() map[string]ViewTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CompanyResponseIncluded) GetTagsOk() (*map[string]ViewTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CompanyResponseIncluded) SetTags(v map[string]ViewTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CompanyResponseIncluded) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


