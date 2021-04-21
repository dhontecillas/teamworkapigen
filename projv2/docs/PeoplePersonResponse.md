# PeoplePersonResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**STATUS** | Pointer to **string** |  | [optional] 
**Person** | Pointer to [**PeopleSpecificPerson**](PeopleSpecificPerson.md) |  | [optional] 

## Methods

### NewPeoplePersonResponse

`func NewPeoplePersonResponse() *PeoplePersonResponse`

NewPeoplePersonResponse instantiates a new PeoplePersonResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeoplePersonResponseWithDefaults

`func NewPeoplePersonResponseWithDefaults() *PeoplePersonResponse`

NewPeoplePersonResponseWithDefaults instantiates a new PeoplePersonResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSTATUS

`func (o *PeoplePersonResponse) GetSTATUS() string`

GetSTATUS returns the STATUS field if non-nil, zero value otherwise.

### GetSTATUSOk

`func (o *PeoplePersonResponse) GetSTATUSOk() (*string, bool)`

GetSTATUSOk returns a tuple with the STATUS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS

`func (o *PeoplePersonResponse) SetSTATUS(v string)`

SetSTATUS sets STATUS field to given value.

### HasSTATUS

`func (o *PeoplePersonResponse) HasSTATUS() bool`

HasSTATUS returns a boolean if a field has been set.

### GetPerson

`func (o *PeoplePersonResponse) GetPerson() PeopleSpecificPerson`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *PeoplePersonResponse) GetPersonOk() (*PeopleSpecificPerson, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *PeoplePersonResponse) SetPerson(v PeopleSpecificPerson)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *PeoplePersonResponse) HasPerson() bool`

HasPerson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


