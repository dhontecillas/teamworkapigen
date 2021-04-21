# PeopleMultiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Included** | Pointer to [**PeopleMultiResponseIncluded**](PeopleMultiResponseIncluded.md) |  | [optional] 
**Meta** | Pointer to [**ViewMeta**](ViewMeta.md) |  | [optional] 
**People** | Pointer to [**[]ViewUser**](ViewUser.md) |  | [optional] 

## Methods

### NewPeopleMultiResponse

`func NewPeopleMultiResponse() *PeopleMultiResponse`

NewPeopleMultiResponse instantiates a new PeopleMultiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeopleMultiResponseWithDefaults

`func NewPeopleMultiResponseWithDefaults() *PeopleMultiResponse`

NewPeopleMultiResponseWithDefaults instantiates a new PeopleMultiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncluded

`func (o *PeopleMultiResponse) GetIncluded() PeopleMultiResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *PeopleMultiResponse) GetIncludedOk() (*PeopleMultiResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *PeopleMultiResponse) SetIncluded(v PeopleMultiResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *PeopleMultiResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *PeopleMultiResponse) GetMeta() ViewMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PeopleMultiResponse) GetMetaOk() (*ViewMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PeopleMultiResponse) SetMeta(v ViewMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PeopleMultiResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetPeople

`func (o *PeopleMultiResponse) GetPeople() []ViewUser`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *PeopleMultiResponse) GetPeopleOk() (*[]ViewUser, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *PeopleMultiResponse) SetPeople(v []ViewUser)`

SetPeople sets People field to given value.

### HasPeople

`func (o *PeopleMultiResponse) HasPeople() bool`

HasPeople returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


