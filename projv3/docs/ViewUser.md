# ViewUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarUrl** | Pointer to **string** |  | [optional] 
**Company** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**CompanyId** | Pointer to **int32** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**IsAdmin** | Pointer to **bool** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**LengthOfDay** | Pointer to **float32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**WorkingHour** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**WorkingHoursId** | Pointer to **int32** |  | [optional] 

## Methods

### NewViewUser

`func NewViewUser() *ViewUser`

NewViewUser instantiates a new ViewUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewUserWithDefaults

`func NewViewUserWithDefaults() *ViewUser`

NewViewUserWithDefaults instantiates a new ViewUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarUrl

`func (o *ViewUser) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *ViewUser) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *ViewUser) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *ViewUser) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetCompany

`func (o *ViewUser) GetCompany() ViewRelationship`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ViewUser) GetCompanyOk() (*ViewRelationship, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ViewUser) SetCompany(v ViewRelationship)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ViewUser) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCompanyId

`func (o *ViewUser) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ViewUser) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ViewUser) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *ViewUser) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetEmail

`func (o *ViewUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ViewUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ViewUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ViewUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *ViewUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ViewUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ViewUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ViewUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetId

`func (o *ViewUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewUser) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsAdmin

`func (o *ViewUser) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *ViewUser) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *ViewUser) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *ViewUser) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetLastName

`func (o *ViewUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ViewUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ViewUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ViewUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLengthOfDay

`func (o *ViewUser) GetLengthOfDay() float32`

GetLengthOfDay returns the LengthOfDay field if non-nil, zero value otherwise.

### GetLengthOfDayOk

`func (o *ViewUser) GetLengthOfDayOk() (*float32, bool)`

GetLengthOfDayOk returns a tuple with the LengthOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthOfDay

`func (o *ViewUser) SetLengthOfDay(v float32)`

SetLengthOfDay sets LengthOfDay field to given value.

### HasLengthOfDay

`func (o *ViewUser) HasLengthOfDay() bool`

HasLengthOfDay returns a boolean if a field has been set.

### GetTitle

`func (o *ViewUser) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ViewUser) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ViewUser) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ViewUser) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *ViewUser) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ViewUser) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ViewUser) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ViewUser) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWorkingHour

`func (o *ViewUser) GetWorkingHour() ViewRelationship`

GetWorkingHour returns the WorkingHour field if non-nil, zero value otherwise.

### GetWorkingHourOk

`func (o *ViewUser) GetWorkingHourOk() (*ViewRelationship, bool)`

GetWorkingHourOk returns a tuple with the WorkingHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingHour

`func (o *ViewUser) SetWorkingHour(v ViewRelationship)`

SetWorkingHour sets WorkingHour field to given value.

### HasWorkingHour

`func (o *ViewUser) HasWorkingHour() bool`

HasWorkingHour returns a boolean if a field has been set.

### GetWorkingHoursId

`func (o *ViewUser) GetWorkingHoursId() int32`

GetWorkingHoursId returns the WorkingHoursId field if non-nil, zero value otherwise.

### GetWorkingHoursIdOk

`func (o *ViewUser) GetWorkingHoursIdOk() (*int32, bool)`

GetWorkingHoursIdOk returns a tuple with the WorkingHoursId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingHoursId

`func (o *ViewUser) SetWorkingHoursId(v int32)`

SetWorkingHoursId sets WorkingHoursId field to given value.

### HasWorkingHoursId

`func (o *ViewUser) HasWorkingHoursId() bool`

HasWorkingHoursId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


