# PeopleAuthData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apikey** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **string** | Timestamp is not a time.Time because we need a specific format in this case. If there are more cases like these it is a good idea to create a new type with a MarshalJSON method. | [optional] 
**Token** | Pointer to **string** |  | [optional] 

## Methods

### NewPeopleAuthData

`func NewPeopleAuthData() *PeopleAuthData`

NewPeopleAuthData instantiates a new PeopleAuthData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeopleAuthDataWithDefaults

`func NewPeopleAuthDataWithDefaults() *PeopleAuthData`

NewPeopleAuthDataWithDefaults instantiates a new PeopleAuthData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApikey

`func (o *PeopleAuthData) GetApikey() string`

GetApikey returns the Apikey field if non-nil, zero value otherwise.

### GetApikeyOk

`func (o *PeopleAuthData) GetApikeyOk() (*string, bool)`

GetApikeyOk returns a tuple with the Apikey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApikey

`func (o *PeopleAuthData) SetApikey(v string)`

SetApikey sets Apikey field to given value.

### HasApikey

`func (o *PeopleAuthData) HasApikey() bool`

HasApikey returns a boolean if a field has been set.

### GetTimestamp

`func (o *PeopleAuthData) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PeopleAuthData) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PeopleAuthData) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PeopleAuthData) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetToken

`func (o *PeopleAuthData) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PeopleAuthData) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PeopleAuthData) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PeopleAuthData) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


