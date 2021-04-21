# PeopleCompanyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**People** | Pointer to [**[]PeoplePersonOfPeople**](PeoplePersonOfPeople.md) |  | [optional] 

## Methods

### NewPeopleCompanyResponse

`func NewPeopleCompanyResponse() *PeopleCompanyResponse`

NewPeopleCompanyResponse instantiates a new PeopleCompanyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeopleCompanyResponseWithDefaults

`func NewPeopleCompanyResponseWithDefaults() *PeopleCompanyResponse`

NewPeopleCompanyResponseWithDefaults instantiates a new PeopleCompanyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PeopleCompanyResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PeopleCompanyResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PeopleCompanyResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PeopleCompanyResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PeopleCompanyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PeopleCompanyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PeopleCompanyResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PeopleCompanyResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPeople

`func (o *PeopleCompanyResponse) GetPeople() []PeoplePersonOfPeople`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *PeopleCompanyResponse) GetPeopleOk() (*[]PeoplePersonOfPeople, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *PeopleCompanyResponse) SetPeople(v []PeoplePersonOfPeople)`

SetPeople sets People field to given value.

### HasPeople

`func (o *PeopleCompanyResponse) HasPeople() bool`

HasPeople returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


