# PeopleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Companies** | Pointer to [**[]PeopleCompanyResponse**](PeopleCompanyResponse.md) |  | [optional] 
**Letters** | Pointer to **[]string** |  | [optional] 
**People** | Pointer to [**[]PeoplePersonOfPeople**](PeoplePersonOfPeople.md) |  | [optional] 

## Methods

### NewPeopleResponse

`func NewPeopleResponse() *PeopleResponse`

NewPeopleResponse instantiates a new PeopleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeopleResponseWithDefaults

`func NewPeopleResponseWithDefaults() *PeopleResponse`

NewPeopleResponseWithDefaults instantiates a new PeopleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanies

`func (o *PeopleResponse) GetCompanies() []PeopleCompanyResponse`

GetCompanies returns the Companies field if non-nil, zero value otherwise.

### GetCompaniesOk

`func (o *PeopleResponse) GetCompaniesOk() (*[]PeopleCompanyResponse, bool)`

GetCompaniesOk returns a tuple with the Companies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanies

`func (o *PeopleResponse) SetCompanies(v []PeopleCompanyResponse)`

SetCompanies sets Companies field to given value.

### HasCompanies

`func (o *PeopleResponse) HasCompanies() bool`

HasCompanies returns a boolean if a field has been set.

### GetLetters

`func (o *PeopleResponse) GetLetters() []string`

GetLetters returns the Letters field if non-nil, zero value otherwise.

### GetLettersOk

`func (o *PeopleResponse) GetLettersOk() (*[]string, bool)`

GetLettersOk returns a tuple with the Letters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLetters

`func (o *PeopleResponse) SetLetters(v []string)`

SetLetters sets Letters field to given value.

### HasLetters

`func (o *PeopleResponse) HasLetters() bool`

HasLetters returns a boolean if a field has been set.

### GetPeople

`func (o *PeopleResponse) GetPeople() []PeoplePersonOfPeople`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *PeopleResponse) GetPeopleOk() (*[]PeoplePersonOfPeople, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *PeopleResponse) SetPeople(v []PeoplePersonOfPeople)`

SetPeople sets People field to given value.

### HasPeople

`func (o *PeopleResponse) HasPeople() bool`

HasPeople returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


