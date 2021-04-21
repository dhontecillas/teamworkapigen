# PeopleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Included** | Pointer to [**PeopleMultiResponseIncluded**](PeopleMultiResponseIncluded.md) |  | [optional] 
**Person** | Pointer to [**ViewUser**](ViewUser.md) |  | [optional] 

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

### GetIncluded

`func (o *PeopleResponse) GetIncluded() PeopleMultiResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *PeopleResponse) GetIncludedOk() (*PeopleMultiResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *PeopleResponse) SetIncluded(v PeopleMultiResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *PeopleResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetPerson

`func (o *PeopleResponse) GetPerson() ViewUser`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *PeopleResponse) GetPersonOk() (*ViewUser, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *PeopleResponse) SetPerson(v ViewUser)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *PeopleResponse) HasPerson() bool`

HasPerson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


