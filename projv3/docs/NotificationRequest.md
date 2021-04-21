# NotificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notification** | Pointer to [**NotificationProjectBudgetNotification**](NotificationProjectBudgetNotification.md) |  | [optional] 

## Methods

### NewNotificationRequest

`func NewNotificationRequest() *NotificationRequest`

NewNotificationRequest instantiates a new NotificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRequestWithDefaults

`func NewNotificationRequestWithDefaults() *NotificationRequest`

NewNotificationRequestWithDefaults instantiates a new NotificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotification

`func (o *NotificationRequest) GetNotification() NotificationProjectBudgetNotification`

GetNotification returns the Notification field if non-nil, zero value otherwise.

### GetNotificationOk

`func (o *NotificationRequest) GetNotificationOk() (*NotificationProjectBudgetNotification, bool)`

GetNotificationOk returns a tuple with the Notification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotification

`func (o *NotificationRequest) SetNotification(v NotificationProjectBudgetNotification)`

SetNotification sets Notification field to given value.

### HasNotification

`func (o *NotificationRequest) HasNotification() bool`

HasNotification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


