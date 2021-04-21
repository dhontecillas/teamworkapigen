# InlineResponse200120

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**STATUS** | Pointer to **string** |  | [optional] 
**Workload** | Pointer to [**[]InlineResponse200120Workload**](InlineResponse200120Workload.md) |  | [optional] 

## Methods

### NewInlineResponse200120

`func NewInlineResponse200120() *InlineResponse200120`

NewInlineResponse200120 instantiates a new InlineResponse200120 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200120WithDefaults

`func NewInlineResponse200120WithDefaults() *InlineResponse200120`

NewInlineResponse200120WithDefaults instantiates a new InlineResponse200120 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSTATUS

`func (o *InlineResponse200120) GetSTATUS() string`

GetSTATUS returns the STATUS field if non-nil, zero value otherwise.

### GetSTATUSOk

`func (o *InlineResponse200120) GetSTATUSOk() (*string, bool)`

GetSTATUSOk returns a tuple with the STATUS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS

`func (o *InlineResponse200120) SetSTATUS(v string)`

SetSTATUS sets STATUS field to given value.

### HasSTATUS

`func (o *InlineResponse200120) HasSTATUS() bool`

HasSTATUS returns a boolean if a field has been set.

### GetWorkload

`func (o *InlineResponse200120) GetWorkload() []InlineResponse200120Workload`

GetWorkload returns the Workload field if non-nil, zero value otherwise.

### GetWorkloadOk

`func (o *InlineResponse200120) GetWorkloadOk() (*[]InlineResponse200120Workload, bool)`

GetWorkloadOk returns a tuple with the Workload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkload

`func (o *InlineResponse200120) SetWorkload(v []InlineResponse200120Workload)`

SetWorkload sets Workload field to given value.

### HasWorkload

`func (o *InlineResponse200120) HasWorkload() bool`

HasWorkload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


