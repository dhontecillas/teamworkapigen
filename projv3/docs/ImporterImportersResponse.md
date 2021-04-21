# ImporterImportersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Importers** | Pointer to [**[]ImporterImporter**](ImporterImporter.md) |  | [optional] 
**IsImporting** | Pointer to **bool** |  | [optional] 

## Methods

### NewImporterImportersResponse

`func NewImporterImportersResponse() *ImporterImportersResponse`

NewImporterImportersResponse instantiates a new ImporterImportersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImporterImportersResponseWithDefaults

`func NewImporterImportersResponseWithDefaults() *ImporterImportersResponse`

NewImporterImportersResponseWithDefaults instantiates a new ImporterImportersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImporters

`func (o *ImporterImportersResponse) GetImporters() []ImporterImporter`

GetImporters returns the Importers field if non-nil, zero value otherwise.

### GetImportersOk

`func (o *ImporterImportersResponse) GetImportersOk() (*[]ImporterImporter, bool)`

GetImportersOk returns a tuple with the Importers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImporters

`func (o *ImporterImportersResponse) SetImporters(v []ImporterImporter)`

SetImporters sets Importers field to given value.

### HasImporters

`func (o *ImporterImportersResponse) HasImporters() bool`

HasImporters returns a boolean if a field has been set.

### GetIsImporting

`func (o *ImporterImportersResponse) GetIsImporting() bool`

GetIsImporting returns the IsImporting field if non-nil, zero value otherwise.

### GetIsImportingOk

`func (o *ImporterImportersResponse) GetIsImportingOk() (*bool, bool)`

GetIsImportingOk returns a tuple with the IsImporting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsImporting

`func (o *ImporterImportersResponse) SetIsImporting(v bool)`

SetIsImporting sets IsImporting field to given value.

### HasIsImporting

`func (o *ImporterImportersResponse) HasIsImporting() bool`

HasIsImporting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


