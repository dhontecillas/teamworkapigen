# MilestoneResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Included** | Pointer to [**MilestoneMilestonesResponseIncluded**](MilestoneMilestonesResponseIncluded.md) |  | [optional] 
**Milestone** | Pointer to [**ViewMilestone**](ViewMilestone.md) |  | [optional] 

## Methods

### NewMilestoneResponse

`func NewMilestoneResponse() *MilestoneResponse`

NewMilestoneResponse instantiates a new MilestoneResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMilestoneResponseWithDefaults

`func NewMilestoneResponseWithDefaults() *MilestoneResponse`

NewMilestoneResponseWithDefaults instantiates a new MilestoneResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncluded

`func (o *MilestoneResponse) GetIncluded() MilestoneMilestonesResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *MilestoneResponse) GetIncludedOk() (*MilestoneMilestonesResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *MilestoneResponse) SetIncluded(v MilestoneMilestonesResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *MilestoneResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMilestone

`func (o *MilestoneResponse) GetMilestone() ViewMilestone`

GetMilestone returns the Milestone field if non-nil, zero value otherwise.

### GetMilestoneOk

`func (o *MilestoneResponse) GetMilestoneOk() (*ViewMilestone, bool)`

GetMilestoneOk returns a tuple with the Milestone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestone

`func (o *MilestoneResponse) SetMilestone(v ViewMilestone)`

SetMilestone sets Milestone field to given value.

### HasMilestone

`func (o *MilestoneResponse) HasMilestone() bool`

HasMilestone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


