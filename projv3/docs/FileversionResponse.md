# FileversionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fileversion** | Pointer to [**ViewFileversion**](ViewFileversion.md) |  | [optional] 
**Included** | Pointer to [**FileversionResponseIncluded**](FileversionResponseIncluded.md) |  | [optional] 

## Methods

### NewFileversionResponse

`func NewFileversionResponse() *FileversionResponse`

NewFileversionResponse instantiates a new FileversionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileversionResponseWithDefaults

`func NewFileversionResponseWithDefaults() *FileversionResponse`

NewFileversionResponseWithDefaults instantiates a new FileversionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileversion

`func (o *FileversionResponse) GetFileversion() ViewFileversion`

GetFileversion returns the Fileversion field if non-nil, zero value otherwise.

### GetFileversionOk

`func (o *FileversionResponse) GetFileversionOk() (*ViewFileversion, bool)`

GetFileversionOk returns a tuple with the Fileversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileversion

`func (o *FileversionResponse) SetFileversion(v ViewFileversion)`

SetFileversion sets Fileversion field to given value.

### HasFileversion

`func (o *FileversionResponse) HasFileversion() bool`

HasFileversion returns a boolean if a field has been set.

### GetIncluded

`func (o *FileversionResponse) GetIncluded() FileversionResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *FileversionResponse) GetIncludedOk() (*FileversionResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *FileversionResponse) SetIncluded(v FileversionResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *FileversionResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


