# InlineResponse20023Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **string** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**CreatedByUserFirstname** | Pointer to **string** |  | [optional] 
**CreatedByUserId** | Pointer to **string** |  | [optional] 
**CreatedByUserLastname** | Pointer to **string** |  | [optional] 
**CurrencyCode** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **string** |  | [optional] 
**DateUpdated** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayDate** | Pointer to **string** |  | [optional] 
**EditedByUserFirstname** | Pointer to **string** |  | [optional] 
**EditedByUserLastname** | Pointer to **string** |  | [optional] 
**Expenses** | Pointer to [**[]InlineResponse20023InvoiceExpenses**](InlineResponse20023InvoiceExpenses.md) |  | [optional] 
**ExportedByUserFirstname** | Pointer to **string** |  | [optional] 
**ExportedByUserId** | Pointer to **string** |  | [optional] 
**ExportedByUserLastname** | Pointer to **string** |  | [optional] 
**ExportedDate** | Pointer to **string** |  | [optional] 
**FixedCost** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LineItems** | Pointer to [**[]InlineResponse20023InvoiceLineItems**](InlineResponse20023InvoiceLineItems.md) |  | [optional] 
**Number** | Pointer to **string** |  | [optional] 
**PoNumber** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**ProjectName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TotalCost** | Pointer to **string** |  | [optional] 
**TotalTime** | Pointer to **string** |  | [optional] 
**TotalTimeDecimal** | Pointer to **string** |  | [optional] 
**UpdateByUserId** | Pointer to **string** |  | [optional] 

## Methods

### NewInlineResponse20023Invoice

`func NewInlineResponse20023Invoice() *InlineResponse20023Invoice`

NewInlineResponse20023Invoice instantiates a new InlineResponse20023Invoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20023InvoiceWithDefaults

`func NewInlineResponse20023InvoiceWithDefaults() *InlineResponse20023Invoice`

NewInlineResponse20023InvoiceWithDefaults instantiates a new InlineResponse20023Invoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *InlineResponse20023Invoice) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *InlineResponse20023Invoice) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *InlineResponse20023Invoice) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *InlineResponse20023Invoice) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCompanyName

`func (o *InlineResponse20023Invoice) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *InlineResponse20023Invoice) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *InlineResponse20023Invoice) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *InlineResponse20023Invoice) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCreatedByUserFirstname

`func (o *InlineResponse20023Invoice) GetCreatedByUserFirstname() string`

GetCreatedByUserFirstname returns the CreatedByUserFirstname field if non-nil, zero value otherwise.

### GetCreatedByUserFirstnameOk

`func (o *InlineResponse20023Invoice) GetCreatedByUserFirstnameOk() (*string, bool)`

GetCreatedByUserFirstnameOk returns a tuple with the CreatedByUserFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserFirstname

`func (o *InlineResponse20023Invoice) SetCreatedByUserFirstname(v string)`

SetCreatedByUserFirstname sets CreatedByUserFirstname field to given value.

### HasCreatedByUserFirstname

`func (o *InlineResponse20023Invoice) HasCreatedByUserFirstname() bool`

HasCreatedByUserFirstname returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *InlineResponse20023Invoice) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *InlineResponse20023Invoice) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *InlineResponse20023Invoice) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *InlineResponse20023Invoice) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetCreatedByUserLastname

`func (o *InlineResponse20023Invoice) GetCreatedByUserLastname() string`

GetCreatedByUserLastname returns the CreatedByUserLastname field if non-nil, zero value otherwise.

### GetCreatedByUserLastnameOk

`func (o *InlineResponse20023Invoice) GetCreatedByUserLastnameOk() (*string, bool)`

GetCreatedByUserLastnameOk returns a tuple with the CreatedByUserLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserLastname

`func (o *InlineResponse20023Invoice) SetCreatedByUserLastname(v string)`

SetCreatedByUserLastname sets CreatedByUserLastname field to given value.

### HasCreatedByUserLastname

`func (o *InlineResponse20023Invoice) HasCreatedByUserLastname() bool`

HasCreatedByUserLastname returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *InlineResponse20023Invoice) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *InlineResponse20023Invoice) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *InlineResponse20023Invoice) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *InlineResponse20023Invoice) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetDateCreated

`func (o *InlineResponse20023Invoice) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *InlineResponse20023Invoice) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *InlineResponse20023Invoice) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *InlineResponse20023Invoice) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateUpdated

`func (o *InlineResponse20023Invoice) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *InlineResponse20023Invoice) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *InlineResponse20023Invoice) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *InlineResponse20023Invoice) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse20023Invoice) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse20023Invoice) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse20023Invoice) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse20023Invoice) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayDate

`func (o *InlineResponse20023Invoice) GetDisplayDate() string`

GetDisplayDate returns the DisplayDate field if non-nil, zero value otherwise.

### GetDisplayDateOk

`func (o *InlineResponse20023Invoice) GetDisplayDateOk() (*string, bool)`

GetDisplayDateOk returns a tuple with the DisplayDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayDate

`func (o *InlineResponse20023Invoice) SetDisplayDate(v string)`

SetDisplayDate sets DisplayDate field to given value.

### HasDisplayDate

`func (o *InlineResponse20023Invoice) HasDisplayDate() bool`

HasDisplayDate returns a boolean if a field has been set.

### GetEditedByUserFirstname

`func (o *InlineResponse20023Invoice) GetEditedByUserFirstname() string`

GetEditedByUserFirstname returns the EditedByUserFirstname field if non-nil, zero value otherwise.

### GetEditedByUserFirstnameOk

`func (o *InlineResponse20023Invoice) GetEditedByUserFirstnameOk() (*string, bool)`

GetEditedByUserFirstnameOk returns a tuple with the EditedByUserFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedByUserFirstname

`func (o *InlineResponse20023Invoice) SetEditedByUserFirstname(v string)`

SetEditedByUserFirstname sets EditedByUserFirstname field to given value.

### HasEditedByUserFirstname

`func (o *InlineResponse20023Invoice) HasEditedByUserFirstname() bool`

HasEditedByUserFirstname returns a boolean if a field has been set.

### GetEditedByUserLastname

`func (o *InlineResponse20023Invoice) GetEditedByUserLastname() string`

GetEditedByUserLastname returns the EditedByUserLastname field if non-nil, zero value otherwise.

### GetEditedByUserLastnameOk

`func (o *InlineResponse20023Invoice) GetEditedByUserLastnameOk() (*string, bool)`

GetEditedByUserLastnameOk returns a tuple with the EditedByUserLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedByUserLastname

`func (o *InlineResponse20023Invoice) SetEditedByUserLastname(v string)`

SetEditedByUserLastname sets EditedByUserLastname field to given value.

### HasEditedByUserLastname

`func (o *InlineResponse20023Invoice) HasEditedByUserLastname() bool`

HasEditedByUserLastname returns a boolean if a field has been set.

### GetExpenses

`func (o *InlineResponse20023Invoice) GetExpenses() []InlineResponse20023InvoiceExpenses`

GetExpenses returns the Expenses field if non-nil, zero value otherwise.

### GetExpensesOk

`func (o *InlineResponse20023Invoice) GetExpensesOk() (*[]InlineResponse20023InvoiceExpenses, bool)`

GetExpensesOk returns a tuple with the Expenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenses

`func (o *InlineResponse20023Invoice) SetExpenses(v []InlineResponse20023InvoiceExpenses)`

SetExpenses sets Expenses field to given value.

### HasExpenses

`func (o *InlineResponse20023Invoice) HasExpenses() bool`

HasExpenses returns a boolean if a field has been set.

### GetExportedByUserFirstname

`func (o *InlineResponse20023Invoice) GetExportedByUserFirstname() string`

GetExportedByUserFirstname returns the ExportedByUserFirstname field if non-nil, zero value otherwise.

### GetExportedByUserFirstnameOk

`func (o *InlineResponse20023Invoice) GetExportedByUserFirstnameOk() (*string, bool)`

GetExportedByUserFirstnameOk returns a tuple with the ExportedByUserFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedByUserFirstname

`func (o *InlineResponse20023Invoice) SetExportedByUserFirstname(v string)`

SetExportedByUserFirstname sets ExportedByUserFirstname field to given value.

### HasExportedByUserFirstname

`func (o *InlineResponse20023Invoice) HasExportedByUserFirstname() bool`

HasExportedByUserFirstname returns a boolean if a field has been set.

### GetExportedByUserId

`func (o *InlineResponse20023Invoice) GetExportedByUserId() string`

GetExportedByUserId returns the ExportedByUserId field if non-nil, zero value otherwise.

### GetExportedByUserIdOk

`func (o *InlineResponse20023Invoice) GetExportedByUserIdOk() (*string, bool)`

GetExportedByUserIdOk returns a tuple with the ExportedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedByUserId

`func (o *InlineResponse20023Invoice) SetExportedByUserId(v string)`

SetExportedByUserId sets ExportedByUserId field to given value.

### HasExportedByUserId

`func (o *InlineResponse20023Invoice) HasExportedByUserId() bool`

HasExportedByUserId returns a boolean if a field has been set.

### GetExportedByUserLastname

`func (o *InlineResponse20023Invoice) GetExportedByUserLastname() string`

GetExportedByUserLastname returns the ExportedByUserLastname field if non-nil, zero value otherwise.

### GetExportedByUserLastnameOk

`func (o *InlineResponse20023Invoice) GetExportedByUserLastnameOk() (*string, bool)`

GetExportedByUserLastnameOk returns a tuple with the ExportedByUserLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedByUserLastname

`func (o *InlineResponse20023Invoice) SetExportedByUserLastname(v string)`

SetExportedByUserLastname sets ExportedByUserLastname field to given value.

### HasExportedByUserLastname

`func (o *InlineResponse20023Invoice) HasExportedByUserLastname() bool`

HasExportedByUserLastname returns a boolean if a field has been set.

### GetExportedDate

`func (o *InlineResponse20023Invoice) GetExportedDate() string`

GetExportedDate returns the ExportedDate field if non-nil, zero value otherwise.

### GetExportedDateOk

`func (o *InlineResponse20023Invoice) GetExportedDateOk() (*string, bool)`

GetExportedDateOk returns a tuple with the ExportedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedDate

`func (o *InlineResponse20023Invoice) SetExportedDate(v string)`

SetExportedDate sets ExportedDate field to given value.

### HasExportedDate

`func (o *InlineResponse20023Invoice) HasExportedDate() bool`

HasExportedDate returns a boolean if a field has been set.

### GetFixedCost

`func (o *InlineResponse20023Invoice) GetFixedCost() string`

GetFixedCost returns the FixedCost field if non-nil, zero value otherwise.

### GetFixedCostOk

`func (o *InlineResponse20023Invoice) GetFixedCostOk() (*string, bool)`

GetFixedCostOk returns a tuple with the FixedCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedCost

`func (o *InlineResponse20023Invoice) SetFixedCost(v string)`

SetFixedCost sets FixedCost field to given value.

### HasFixedCost

`func (o *InlineResponse20023Invoice) HasFixedCost() bool`

HasFixedCost returns a boolean if a field has been set.

### GetId

`func (o *InlineResponse20023Invoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20023Invoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20023Invoice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20023Invoice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLineItems

`func (o *InlineResponse20023Invoice) GetLineItems() []InlineResponse20023InvoiceLineItems`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *InlineResponse20023Invoice) GetLineItemsOk() (*[]InlineResponse20023InvoiceLineItems, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *InlineResponse20023Invoice) SetLineItems(v []InlineResponse20023InvoiceLineItems)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *InlineResponse20023Invoice) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetNumber

`func (o *InlineResponse20023Invoice) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InlineResponse20023Invoice) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InlineResponse20023Invoice) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *InlineResponse20023Invoice) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetPoNumber

`func (o *InlineResponse20023Invoice) GetPoNumber() string`

GetPoNumber returns the PoNumber field if non-nil, zero value otherwise.

### GetPoNumberOk

`func (o *InlineResponse20023Invoice) GetPoNumberOk() (*string, bool)`

GetPoNumberOk returns a tuple with the PoNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoNumber

`func (o *InlineResponse20023Invoice) SetPoNumber(v string)`

SetPoNumber sets PoNumber field to given value.

### HasPoNumber

`func (o *InlineResponse20023Invoice) HasPoNumber() bool`

HasPoNumber returns a boolean if a field has been set.

### GetProjectId

`func (o *InlineResponse20023Invoice) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *InlineResponse20023Invoice) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *InlineResponse20023Invoice) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *InlineResponse20023Invoice) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProjectName

`func (o *InlineResponse20023Invoice) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *InlineResponse20023Invoice) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *InlineResponse20023Invoice) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *InlineResponse20023Invoice) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20023Invoice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20023Invoice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20023Invoice) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20023Invoice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalCost

`func (o *InlineResponse20023Invoice) GetTotalCost() string`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *InlineResponse20023Invoice) GetTotalCostOk() (*string, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *InlineResponse20023Invoice) SetTotalCost(v string)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *InlineResponse20023Invoice) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### GetTotalTime

`func (o *InlineResponse20023Invoice) GetTotalTime() string`

GetTotalTime returns the TotalTime field if non-nil, zero value otherwise.

### GetTotalTimeOk

`func (o *InlineResponse20023Invoice) GetTotalTimeOk() (*string, bool)`

GetTotalTimeOk returns a tuple with the TotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTime

`func (o *InlineResponse20023Invoice) SetTotalTime(v string)`

SetTotalTime sets TotalTime field to given value.

### HasTotalTime

`func (o *InlineResponse20023Invoice) HasTotalTime() bool`

HasTotalTime returns a boolean if a field has been set.

### GetTotalTimeDecimal

`func (o *InlineResponse20023Invoice) GetTotalTimeDecimal() string`

GetTotalTimeDecimal returns the TotalTimeDecimal field if non-nil, zero value otherwise.

### GetTotalTimeDecimalOk

`func (o *InlineResponse20023Invoice) GetTotalTimeDecimalOk() (*string, bool)`

GetTotalTimeDecimalOk returns a tuple with the TotalTimeDecimal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTimeDecimal

`func (o *InlineResponse20023Invoice) SetTotalTimeDecimal(v string)`

SetTotalTimeDecimal sets TotalTimeDecimal field to given value.

### HasTotalTimeDecimal

`func (o *InlineResponse20023Invoice) HasTotalTimeDecimal() bool`

HasTotalTimeDecimal returns a boolean if a field has been set.

### GetUpdateByUserId

`func (o *InlineResponse20023Invoice) GetUpdateByUserId() string`

GetUpdateByUserId returns the UpdateByUserId field if non-nil, zero value otherwise.

### GetUpdateByUserIdOk

`func (o *InlineResponse20023Invoice) GetUpdateByUserIdOk() (*string, bool)`

GetUpdateByUserIdOk returns a tuple with the UpdateByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateByUserId

`func (o *InlineResponse20023Invoice) SetUpdateByUserId(v string)`

SetUpdateByUserId sets UpdateByUserId field to given value.

### HasUpdateByUserId

`func (o *InlineResponse20023Invoice) HasUpdateByUserId() bool`

HasUpdateByUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


