# CommentFullComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**DateDeleted** | Pointer to **string** |  | [optional] 
**DateLastEdited** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**DeletedBy** | Pointer to **int32** |  | [optional] 
**DeletedByUserId** | Pointer to **int32** |  | [optional] 
**FileIds** | Pointer to **[]int32** |  | [optional] 
**Files** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Installation** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**InstallationId** | Pointer to **int32** |  | [optional] 
**LastEditedBy** | Pointer to **int32** |  | [optional] 
**LastEditedByUserId** | Pointer to **int32** |  | [optional] 
**Object** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**ObjectId** | Pointer to **int32** |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**PeopleNotifiedCount** | Pointer to **int32** |  | [optional] 
**PostedBy** | Pointer to **int32** |  | [optional] 
**PostedByUserId** | Pointer to **int32** |  | [optional] 
**PostedDateTime** | Pointer to **string** |  | [optional] 
**Project** | Pointer to [**ViewRelationship**](ViewRelationship.md) |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**Reactions** | Pointer to [**[]ViewReaction**](ViewReaction.md) |  | [optional] 
**ReactionsCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewCommentFullComment

`func NewCommentFullComment() *CommentFullComment`

NewCommentFullComment instantiates a new CommentFullComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentFullCommentWithDefaults

`func NewCommentFullCommentWithDefaults() *CommentFullComment`

NewCommentFullCommentWithDefaults instantiates a new CommentFullComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *CommentFullComment) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CommentFullComment) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CommentFullComment) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *CommentFullComment) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetContentType

`func (o *CommentFullComment) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *CommentFullComment) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *CommentFullComment) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *CommentFullComment) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetDateDeleted

`func (o *CommentFullComment) GetDateDeleted() string`

GetDateDeleted returns the DateDeleted field if non-nil, zero value otherwise.

### GetDateDeletedOk

`func (o *CommentFullComment) GetDateDeletedOk() (*string, bool)`

GetDateDeletedOk returns a tuple with the DateDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeleted

`func (o *CommentFullComment) SetDateDeleted(v string)`

SetDateDeleted sets DateDeleted field to given value.

### HasDateDeleted

`func (o *CommentFullComment) HasDateDeleted() bool`

HasDateDeleted returns a boolean if a field has been set.

### GetDateLastEdited

`func (o *CommentFullComment) GetDateLastEdited() string`

GetDateLastEdited returns the DateLastEdited field if non-nil, zero value otherwise.

### GetDateLastEditedOk

`func (o *CommentFullComment) GetDateLastEditedOk() (*string, bool)`

GetDateLastEditedOk returns a tuple with the DateLastEdited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateLastEdited

`func (o *CommentFullComment) SetDateLastEdited(v string)`

SetDateLastEdited sets DateLastEdited field to given value.

### HasDateLastEdited

`func (o *CommentFullComment) HasDateLastEdited() bool`

HasDateLastEdited returns a boolean if a field has been set.

### GetDeleted

`func (o *CommentFullComment) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *CommentFullComment) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *CommentFullComment) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *CommentFullComment) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDeletedBy

`func (o *CommentFullComment) GetDeletedBy() int32`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *CommentFullComment) GetDeletedByOk() (*int32, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *CommentFullComment) SetDeletedBy(v int32)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *CommentFullComment) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetDeletedByUserId

`func (o *CommentFullComment) GetDeletedByUserId() int32`

GetDeletedByUserId returns the DeletedByUserId field if non-nil, zero value otherwise.

### GetDeletedByUserIdOk

`func (o *CommentFullComment) GetDeletedByUserIdOk() (*int32, bool)`

GetDeletedByUserIdOk returns a tuple with the DeletedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedByUserId

`func (o *CommentFullComment) SetDeletedByUserId(v int32)`

SetDeletedByUserId sets DeletedByUserId field to given value.

### HasDeletedByUserId

`func (o *CommentFullComment) HasDeletedByUserId() bool`

HasDeletedByUserId returns a boolean if a field has been set.

### GetFileIds

`func (o *CommentFullComment) GetFileIds() []int32`

GetFileIds returns the FileIds field if non-nil, zero value otherwise.

### GetFileIdsOk

`func (o *CommentFullComment) GetFileIdsOk() (*[]int32, bool)`

GetFileIdsOk returns a tuple with the FileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileIds

`func (o *CommentFullComment) SetFileIds(v []int32)`

SetFileIds sets FileIds field to given value.

### HasFileIds

`func (o *CommentFullComment) HasFileIds() bool`

HasFileIds returns a boolean if a field has been set.

### GetFiles

`func (o *CommentFullComment) GetFiles() []ViewRelationship`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *CommentFullComment) GetFilesOk() (*[]ViewRelationship, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *CommentFullComment) SetFiles(v []ViewRelationship)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *CommentFullComment) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetId

`func (o *CommentFullComment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommentFullComment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommentFullComment) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CommentFullComment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstallation

`func (o *CommentFullComment) GetInstallation() ViewRelationship`

GetInstallation returns the Installation field if non-nil, zero value otherwise.

### GetInstallationOk

`func (o *CommentFullComment) GetInstallationOk() (*ViewRelationship, bool)`

GetInstallationOk returns a tuple with the Installation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallation

`func (o *CommentFullComment) SetInstallation(v ViewRelationship)`

SetInstallation sets Installation field to given value.

### HasInstallation

`func (o *CommentFullComment) HasInstallation() bool`

HasInstallation returns a boolean if a field has been set.

### GetInstallationId

`func (o *CommentFullComment) GetInstallationId() int32`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *CommentFullComment) GetInstallationIdOk() (*int32, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *CommentFullComment) SetInstallationId(v int32)`

SetInstallationId sets InstallationId field to given value.

### HasInstallationId

`func (o *CommentFullComment) HasInstallationId() bool`

HasInstallationId returns a boolean if a field has been set.

### GetLastEditedBy

`func (o *CommentFullComment) GetLastEditedBy() int32`

GetLastEditedBy returns the LastEditedBy field if non-nil, zero value otherwise.

### GetLastEditedByOk

`func (o *CommentFullComment) GetLastEditedByOk() (*int32, bool)`

GetLastEditedByOk returns a tuple with the LastEditedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditedBy

`func (o *CommentFullComment) SetLastEditedBy(v int32)`

SetLastEditedBy sets LastEditedBy field to given value.

### HasLastEditedBy

`func (o *CommentFullComment) HasLastEditedBy() bool`

HasLastEditedBy returns a boolean if a field has been set.

### GetLastEditedByUserId

`func (o *CommentFullComment) GetLastEditedByUserId() int32`

GetLastEditedByUserId returns the LastEditedByUserId field if non-nil, zero value otherwise.

### GetLastEditedByUserIdOk

`func (o *CommentFullComment) GetLastEditedByUserIdOk() (*int32, bool)`

GetLastEditedByUserIdOk returns a tuple with the LastEditedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditedByUserId

`func (o *CommentFullComment) SetLastEditedByUserId(v int32)`

SetLastEditedByUserId sets LastEditedByUserId field to given value.

### HasLastEditedByUserId

`func (o *CommentFullComment) HasLastEditedByUserId() bool`

HasLastEditedByUserId returns a boolean if a field has been set.

### GetObject

`func (o *CommentFullComment) GetObject() ViewRelationship`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CommentFullComment) GetObjectOk() (*ViewRelationship, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CommentFullComment) SetObject(v ViewRelationship)`

SetObject sets Object field to given value.

### HasObject

`func (o *CommentFullComment) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetObjectId

`func (o *CommentFullComment) GetObjectId() int32`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *CommentFullComment) GetObjectIdOk() (*int32, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *CommentFullComment) SetObjectId(v int32)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *CommentFullComment) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetObjectType

`func (o *CommentFullComment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CommentFullComment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CommentFullComment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *CommentFullComment) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetPeopleNotifiedCount

`func (o *CommentFullComment) GetPeopleNotifiedCount() int32`

GetPeopleNotifiedCount returns the PeopleNotifiedCount field if non-nil, zero value otherwise.

### GetPeopleNotifiedCountOk

`func (o *CommentFullComment) GetPeopleNotifiedCountOk() (*int32, bool)`

GetPeopleNotifiedCountOk returns a tuple with the PeopleNotifiedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeopleNotifiedCount

`func (o *CommentFullComment) SetPeopleNotifiedCount(v int32)`

SetPeopleNotifiedCount sets PeopleNotifiedCount field to given value.

### HasPeopleNotifiedCount

`func (o *CommentFullComment) HasPeopleNotifiedCount() bool`

HasPeopleNotifiedCount returns a boolean if a field has been set.

### GetPostedBy

`func (o *CommentFullComment) GetPostedBy() int32`

GetPostedBy returns the PostedBy field if non-nil, zero value otherwise.

### GetPostedByOk

`func (o *CommentFullComment) GetPostedByOk() (*int32, bool)`

GetPostedByOk returns a tuple with the PostedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedBy

`func (o *CommentFullComment) SetPostedBy(v int32)`

SetPostedBy sets PostedBy field to given value.

### HasPostedBy

`func (o *CommentFullComment) HasPostedBy() bool`

HasPostedBy returns a boolean if a field has been set.

### GetPostedByUserId

`func (o *CommentFullComment) GetPostedByUserId() int32`

GetPostedByUserId returns the PostedByUserId field if non-nil, zero value otherwise.

### GetPostedByUserIdOk

`func (o *CommentFullComment) GetPostedByUserIdOk() (*int32, bool)`

GetPostedByUserIdOk returns a tuple with the PostedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedByUserId

`func (o *CommentFullComment) SetPostedByUserId(v int32)`

SetPostedByUserId sets PostedByUserId field to given value.

### HasPostedByUserId

`func (o *CommentFullComment) HasPostedByUserId() bool`

HasPostedByUserId returns a boolean if a field has been set.

### GetPostedDateTime

`func (o *CommentFullComment) GetPostedDateTime() string`

GetPostedDateTime returns the PostedDateTime field if non-nil, zero value otherwise.

### GetPostedDateTimeOk

`func (o *CommentFullComment) GetPostedDateTimeOk() (*string, bool)`

GetPostedDateTimeOk returns a tuple with the PostedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedDateTime

`func (o *CommentFullComment) SetPostedDateTime(v string)`

SetPostedDateTime sets PostedDateTime field to given value.

### HasPostedDateTime

`func (o *CommentFullComment) HasPostedDateTime() bool`

HasPostedDateTime returns a boolean if a field has been set.

### GetProject

`func (o *CommentFullComment) GetProject() ViewRelationship`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CommentFullComment) GetProjectOk() (*ViewRelationship, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CommentFullComment) SetProject(v ViewRelationship)`

SetProject sets Project field to given value.

### HasProject

`func (o *CommentFullComment) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectId

`func (o *CommentFullComment) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CommentFullComment) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CommentFullComment) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CommentFullComment) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetReactions

`func (o *CommentFullComment) GetReactions() []ViewReaction`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *CommentFullComment) GetReactionsOk() (*[]ViewReaction, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *CommentFullComment) SetReactions(v []ViewReaction)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *CommentFullComment) HasReactions() bool`

HasReactions returns a boolean if a field has been set.

### GetReactionsCount

`func (o *CommentFullComment) GetReactionsCount() int32`

GetReactionsCount returns the ReactionsCount field if non-nil, zero value otherwise.

### GetReactionsCountOk

`func (o *CommentFullComment) GetReactionsCountOk() (*int32, bool)`

GetReactionsCountOk returns a tuple with the ReactionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionsCount

`func (o *CommentFullComment) SetReactionsCount(v int32)`

SetReactionsCount sets ReactionsCount field to given value.

### HasReactionsCount

`func (o *CommentFullComment) HasReactionsCount() bool`

HasReactionsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


