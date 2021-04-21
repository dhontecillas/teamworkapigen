# PeoplePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | project only permissions used by /projects/:id/people/:id endpoint | [optional] 
**AddFiles** | Pointer to **bool** |  | [optional] 
**AddLinks** | Pointer to **bool** |  | [optional] 
**AddMessages** | Pointer to **bool** |  | [optional] 
**AddMilestones** | Pointer to **bool** |  | [optional] 
**AddNotebooks** | Pointer to **bool** |  | [optional] 
**AddPeopleToProject** | Pointer to **bool** |  | [optional] 
**AddProjectUpdate** | Pointer to **bool** |  | [optional] 
**AddRisks** | Pointer to **bool** |  | [optional] 
**AddTaskLists** | Pointer to **bool** |  | [optional] 
**AddTasks** | Pointer to **bool** |  | [optional] 
**AddTime** | Pointer to **bool** |  | [optional] 
**Administrator** | Pointer to **bool** | used when retrieving a specific person | [optional] 
**CanAccess** | Pointer to **bool** |  | [optional] 
**CanAccessBox** | Pointer to **bool** |  | [optional] 
**CanAccessCalendar** | Pointer to **bool** |  | [optional] 
**CanAccessDropbox** | Pointer to **bool** |  | [optional] 
**CanAccessGoogleDocs** | Pointer to **bool** |  | [optional] 
**CanAccessInvoiceTracking** | Pointer to **bool** |  | [optional] 
**CanAccessOneDrive** | Pointer to **bool** |  | [optional] 
**CanAccessOneDriveBusiness** | Pointer to **bool** |  | [optional] 
**CanAccessPortfolio** | Pointer to **bool** |  | [optional] 
**CanAccessTemplates** | Pointer to **bool** |  | [optional] 
**CanAddForms** | Pointer to **bool** |  | [optional] 
**CanAddProjects** | Pointer to **bool** |  | [optional] 
**CanBeAssignedToTasksAndMilestones** | Pointer to **bool** |  | [optional] 
**CanManageCustomFields** | Pointer to **bool** |  | [optional] 
**CanManagePeople** | Pointer to **bool** |  | [optional] 
**CanManagePortfolio** | Pointer to **bool** |  | [optional] 
**CanManageProjectBudget** | Pointer to **bool** |  | [optional] 
**CanManageProjectTemplates** | Pointer to **bool** |  | [optional] 
**CanManageRates** | Pointer to **bool** |  | [optional] 
**CanManageReports** | Pointer to **bool** |  | [optional] 
**CanManageSchedule** | Pointer to **bool** |  | [optional] 
**CanReceiveEmail** | Pointer to **bool** |  | [optional] 
**CanViewForms** | Pointer to **bool** |  | [optional] 
**CanViewProjectBudget** | Pointer to **bool** |  | [optional] 
**CanViewProjectMembers** | Pointer to **bool** |  | [optional] 
**CanViewProjectTemplates** | Pointer to **bool** |  | [optional] 
**CanViewRates** | Pointer to **bool** |  | [optional] 
**CanViewReports** | Pointer to **bool** |  | [optional] 
**CanViewSchedule** | Pointer to **bool** |  | [optional] 
**EditAllTasks** | Pointer to **bool** |  | [optional] 
**HasAccessToNewProjects** | Pointer to **bool** |  | [optional] 
**IsArchived** | Pointer to **bool** |  | [optional] 
**IsObserving** | Pointer to **bool** |  | [optional] 
**ManageCustomFields** | Pointer to **bool** |  | [optional] 
**NotifyDefaults** | Pointer to [**PeopleNotifyDefaults**](PeopleNotifyDefaults.md) |  | [optional] 
**ProjectAdministrator** | Pointer to **bool** |  | [optional] 
**SetPrivacy** | Pointer to **bool** |  | [optional] 
**ViewAllTimeLogs** | Pointer to **bool** |  | [optional] 
**ViewEstimatedTime** | Pointer to **bool** |  | [optional] 
**ViewInvoices** | Pointer to **bool** |  | [optional] 
**ViewLinks** | Pointer to **bool** |  | [optional] 
**ViewMessagesAndFiles** | Pointer to **bool** |  | [optional] 
**ViewNotebooks** | Pointer to **bool** |  | [optional] 
**ViewProjectUpdate** | Pointer to **bool** |  | [optional] 
**ViewRiskRegister** | Pointer to **bool** |  | [optional] 
**ViewTasksAndMilestones** | Pointer to **bool** |  | [optional] 
**ViewTime** | Pointer to **bool** |  | [optional] 

## Methods

### NewPeoplePermissions

`func NewPeoplePermissions() *PeoplePermissions`

NewPeoplePermissions instantiates a new PeoplePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeoplePermissionsWithDefaults

`func NewPeoplePermissionsWithDefaults() *PeoplePermissions`

NewPeoplePermissionsWithDefaults instantiates a new PeoplePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *PeoplePermissions) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PeoplePermissions) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PeoplePermissions) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PeoplePermissions) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAddFiles

`func (o *PeoplePermissions) GetAddFiles() bool`

GetAddFiles returns the AddFiles field if non-nil, zero value otherwise.

### GetAddFilesOk

`func (o *PeoplePermissions) GetAddFilesOk() (*bool, bool)`

GetAddFilesOk returns a tuple with the AddFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddFiles

`func (o *PeoplePermissions) SetAddFiles(v bool)`

SetAddFiles sets AddFiles field to given value.

### HasAddFiles

`func (o *PeoplePermissions) HasAddFiles() bool`

HasAddFiles returns a boolean if a field has been set.

### GetAddLinks

`func (o *PeoplePermissions) GetAddLinks() bool`

GetAddLinks returns the AddLinks field if non-nil, zero value otherwise.

### GetAddLinksOk

`func (o *PeoplePermissions) GetAddLinksOk() (*bool, bool)`

GetAddLinksOk returns a tuple with the AddLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddLinks

`func (o *PeoplePermissions) SetAddLinks(v bool)`

SetAddLinks sets AddLinks field to given value.

### HasAddLinks

`func (o *PeoplePermissions) HasAddLinks() bool`

HasAddLinks returns a boolean if a field has been set.

### GetAddMessages

`func (o *PeoplePermissions) GetAddMessages() bool`

GetAddMessages returns the AddMessages field if non-nil, zero value otherwise.

### GetAddMessagesOk

`func (o *PeoplePermissions) GetAddMessagesOk() (*bool, bool)`

GetAddMessagesOk returns a tuple with the AddMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddMessages

`func (o *PeoplePermissions) SetAddMessages(v bool)`

SetAddMessages sets AddMessages field to given value.

### HasAddMessages

`func (o *PeoplePermissions) HasAddMessages() bool`

HasAddMessages returns a boolean if a field has been set.

### GetAddMilestones

`func (o *PeoplePermissions) GetAddMilestones() bool`

GetAddMilestones returns the AddMilestones field if non-nil, zero value otherwise.

### GetAddMilestonesOk

`func (o *PeoplePermissions) GetAddMilestonesOk() (*bool, bool)`

GetAddMilestonesOk returns a tuple with the AddMilestones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddMilestones

`func (o *PeoplePermissions) SetAddMilestones(v bool)`

SetAddMilestones sets AddMilestones field to given value.

### HasAddMilestones

`func (o *PeoplePermissions) HasAddMilestones() bool`

HasAddMilestones returns a boolean if a field has been set.

### GetAddNotebooks

`func (o *PeoplePermissions) GetAddNotebooks() bool`

GetAddNotebooks returns the AddNotebooks field if non-nil, zero value otherwise.

### GetAddNotebooksOk

`func (o *PeoplePermissions) GetAddNotebooksOk() (*bool, bool)`

GetAddNotebooksOk returns a tuple with the AddNotebooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddNotebooks

`func (o *PeoplePermissions) SetAddNotebooks(v bool)`

SetAddNotebooks sets AddNotebooks field to given value.

### HasAddNotebooks

`func (o *PeoplePermissions) HasAddNotebooks() bool`

HasAddNotebooks returns a boolean if a field has been set.

### GetAddPeopleToProject

`func (o *PeoplePermissions) GetAddPeopleToProject() bool`

GetAddPeopleToProject returns the AddPeopleToProject field if non-nil, zero value otherwise.

### GetAddPeopleToProjectOk

`func (o *PeoplePermissions) GetAddPeopleToProjectOk() (*bool, bool)`

GetAddPeopleToProjectOk returns a tuple with the AddPeopleToProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddPeopleToProject

`func (o *PeoplePermissions) SetAddPeopleToProject(v bool)`

SetAddPeopleToProject sets AddPeopleToProject field to given value.

### HasAddPeopleToProject

`func (o *PeoplePermissions) HasAddPeopleToProject() bool`

HasAddPeopleToProject returns a boolean if a field has been set.

### GetAddProjectUpdate

`func (o *PeoplePermissions) GetAddProjectUpdate() bool`

GetAddProjectUpdate returns the AddProjectUpdate field if non-nil, zero value otherwise.

### GetAddProjectUpdateOk

`func (o *PeoplePermissions) GetAddProjectUpdateOk() (*bool, bool)`

GetAddProjectUpdateOk returns a tuple with the AddProjectUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddProjectUpdate

`func (o *PeoplePermissions) SetAddProjectUpdate(v bool)`

SetAddProjectUpdate sets AddProjectUpdate field to given value.

### HasAddProjectUpdate

`func (o *PeoplePermissions) HasAddProjectUpdate() bool`

HasAddProjectUpdate returns a boolean if a field has been set.

### GetAddRisks

`func (o *PeoplePermissions) GetAddRisks() bool`

GetAddRisks returns the AddRisks field if non-nil, zero value otherwise.

### GetAddRisksOk

`func (o *PeoplePermissions) GetAddRisksOk() (*bool, bool)`

GetAddRisksOk returns a tuple with the AddRisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddRisks

`func (o *PeoplePermissions) SetAddRisks(v bool)`

SetAddRisks sets AddRisks field to given value.

### HasAddRisks

`func (o *PeoplePermissions) HasAddRisks() bool`

HasAddRisks returns a boolean if a field has been set.

### GetAddTaskLists

`func (o *PeoplePermissions) GetAddTaskLists() bool`

GetAddTaskLists returns the AddTaskLists field if non-nil, zero value otherwise.

### GetAddTaskListsOk

`func (o *PeoplePermissions) GetAddTaskListsOk() (*bool, bool)`

GetAddTaskListsOk returns a tuple with the AddTaskLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTaskLists

`func (o *PeoplePermissions) SetAddTaskLists(v bool)`

SetAddTaskLists sets AddTaskLists field to given value.

### HasAddTaskLists

`func (o *PeoplePermissions) HasAddTaskLists() bool`

HasAddTaskLists returns a boolean if a field has been set.

### GetAddTasks

`func (o *PeoplePermissions) GetAddTasks() bool`

GetAddTasks returns the AddTasks field if non-nil, zero value otherwise.

### GetAddTasksOk

`func (o *PeoplePermissions) GetAddTasksOk() (*bool, bool)`

GetAddTasksOk returns a tuple with the AddTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTasks

`func (o *PeoplePermissions) SetAddTasks(v bool)`

SetAddTasks sets AddTasks field to given value.

### HasAddTasks

`func (o *PeoplePermissions) HasAddTasks() bool`

HasAddTasks returns a boolean if a field has been set.

### GetAddTime

`func (o *PeoplePermissions) GetAddTime() bool`

GetAddTime returns the AddTime field if non-nil, zero value otherwise.

### GetAddTimeOk

`func (o *PeoplePermissions) GetAddTimeOk() (*bool, bool)`

GetAddTimeOk returns a tuple with the AddTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTime

`func (o *PeoplePermissions) SetAddTime(v bool)`

SetAddTime sets AddTime field to given value.

### HasAddTime

`func (o *PeoplePermissions) HasAddTime() bool`

HasAddTime returns a boolean if a field has been set.

### GetAdministrator

`func (o *PeoplePermissions) GetAdministrator() bool`

GetAdministrator returns the Administrator field if non-nil, zero value otherwise.

### GetAdministratorOk

`func (o *PeoplePermissions) GetAdministratorOk() (*bool, bool)`

GetAdministratorOk returns a tuple with the Administrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrator

`func (o *PeoplePermissions) SetAdministrator(v bool)`

SetAdministrator sets Administrator field to given value.

### HasAdministrator

`func (o *PeoplePermissions) HasAdministrator() bool`

HasAdministrator returns a boolean if a field has been set.

### GetCanAccess

`func (o *PeoplePermissions) GetCanAccess() bool`

GetCanAccess returns the CanAccess field if non-nil, zero value otherwise.

### GetCanAccessOk

`func (o *PeoplePermissions) GetCanAccessOk() (*bool, bool)`

GetCanAccessOk returns a tuple with the CanAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAccess

`func (o *PeoplePermissions) SetCanAccess(v bool)`

SetCanAccess sets CanAccess field to given value.

### HasCanAccess

`func (o *PeoplePermissions) HasCanAccess() bool`

HasCanAccess returns a boolean if a field has been set.

### GetCanAccessBox

`func (o *PeoplePermissions) GetCanAccessBox() bool`

GetCanAccessBox returns the CanAccessBox field if non-nil, zero value otherwise.

### GetCanAccessBoxOk

`func (o *PeoplePermissions) GetCanAccessBoxOk() (*bool, bool)`

GetCanAccessBoxOk returns a tuple with the CanAccessBox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAccessBox

`func (o *PeoplePermissions) SetCanAccessBox(v bool)`

SetCanAccessBox sets CanAccessBox field to given value.

### HasCanAccessBox

`func (o *PeoplePermissions) HasCanAccessBox() bool`

HasCanAccessBox returns a boolean if a field has been set.

### GetCanAccessCalendar

`func (o *PeoplePermissions) GetCanAccessCalendar() bool`

GetCanAccessCalendar returns the CanAccessCalendar field if non-nil, zero value otherwise.

### GetCanAccessCalendarOk

`func (o *PeoplePermissions) GetCanAccessCalendarOk() (*bool, bool)`

GetCanAccessCalendarOk returns a tuple with the CanAccessCalendar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAccessCalendar

`func (o *PeoplePermissions) SetCanAccessCalendar(v bool)`

SetCanAccessCalendar sets CanAccessCalendar field to given value.

### HasCanAccessCalendar

`func (o *PeoplePermissions) HasCanAccessCalendar() bool`

HasCanAccessCalendar returns a boolean if a field has been set.

### GetCanAccessDropbox

`func (o *PeoplePermissions) GetCanAccessDropbox() bool`

GetCanAccessDropbox returns the CanAccessDropbox field if non-nil, zero value otherwise.

### GetCanAccessDropboxOk

`func (o *PeoplePermissions) GetCanAccessDropboxOk() (*bool, bool)`

GetCanAccessDropboxOk returns a tuple with the CanAccessDropbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAccessDropbox

`func (o *PeoplePermissions) SetCanAccessDropbox(v bool)`

SetCanAccessDropbox sets CanAccessDropbox field to given value.

### HasCanAccessDropbox

`func (o *PeoplePermissions) HasCanAccessDropbox() bool`

HasCanAccessDropbox returns a boolean if a field has been set.

### GetCanAccessGoogleDocs

`func (o *PeoplePermissions) GetCanAccessGoogleDocs() bool`

GetCanAccessGoogleDocs returns the CanAccessGoogleDocs field if non-nil, zero value otherwise.

### GetCanAccessGoogleDocsOk

`func (o *PeoplePermissions) GetCanAccessGoogleDocsOk() (*bool, bool)`

GetCanAccessGoogleDocsOk returns a tuple with the CanAccessGoogleDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAccessGoogleDocs

`func (o *PeoplePermissions) SetCanAccessGoogleDocs(v bool)`

SetCanAccessGoogleDocs sets CanAccessGoogleDocs field to given value.

### HasCanAccessGoogleDocs

`func (o *PeoplePermissions) HasCanAccessGoogleDocs() bool`

HasCanAccessGoogleDocs returns a boolean if a field has been set.

### GetCanAccessInvoiceTracking

`func (o *PeoplePermissions) GetCanAccessInvoiceTracking() bool`

GetCanAccessInvoiceTracking returns the CanAccessInvoiceTracking field if non-nil, zero value otherwise.

### GetCanAccessInvoiceTrackingOk

`func (o *PeoplePermissions) GetCanAccessInvoiceTrackingOk() (*bool, bool)`

GetCanAccessInvoiceTrackingOk returns a tuple with the CanAccessInvoiceTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAccessInvoiceTracking

`func (o *PeoplePermissions) SetCanAccessInvoiceTracking(v bool)`

SetCanAccessInvoiceTracking sets CanAccessInvoiceTracking field to given value.

### HasCanAccessInvoiceTracking

`func (o *PeoplePermissions) HasCanAccessInvoiceTracking() bool`

HasCanAccessInvoiceTracking returns a boolean if a field has been set.

### GetCanAccessOneDrive

`func (o *PeoplePermissions) GetCanAccessOneDrive() bool`

GetCanAccessOneDrive returns the CanAccessOneDrive field if non-nil, zero value otherwise.

### GetCanAccessOneDriveOk

`func (o *PeoplePermissions) GetCanAccessOneDriveOk() (*bool, bool)`

GetCanAccessOneDriveOk returns a tuple with the CanAccessOneDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAccessOneDrive

`func (o *PeoplePermissions) SetCanAccessOneDrive(v bool)`

SetCanAccessOneDrive sets CanAccessOneDrive field to given value.

### HasCanAccessOneDrive

`func (o *PeoplePermissions) HasCanAccessOneDrive() bool`

HasCanAccessOneDrive returns a boolean if a field has been set.

### GetCanAccessOneDriveBusiness

`func (o *PeoplePermissions) GetCanAccessOneDriveBusiness() bool`

GetCanAccessOneDriveBusiness returns the CanAccessOneDriveBusiness field if non-nil, zero value otherwise.

### GetCanAccessOneDriveBusinessOk

`func (o *PeoplePermissions) GetCanAccessOneDriveBusinessOk() (*bool, bool)`

GetCanAccessOneDriveBusinessOk returns a tuple with the CanAccessOneDriveBusiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAccessOneDriveBusiness

`func (o *PeoplePermissions) SetCanAccessOneDriveBusiness(v bool)`

SetCanAccessOneDriveBusiness sets CanAccessOneDriveBusiness field to given value.

### HasCanAccessOneDriveBusiness

`func (o *PeoplePermissions) HasCanAccessOneDriveBusiness() bool`

HasCanAccessOneDriveBusiness returns a boolean if a field has been set.

### GetCanAccessPortfolio

`func (o *PeoplePermissions) GetCanAccessPortfolio() bool`

GetCanAccessPortfolio returns the CanAccessPortfolio field if non-nil, zero value otherwise.

### GetCanAccessPortfolioOk

`func (o *PeoplePermissions) GetCanAccessPortfolioOk() (*bool, bool)`

GetCanAccessPortfolioOk returns a tuple with the CanAccessPortfolio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAccessPortfolio

`func (o *PeoplePermissions) SetCanAccessPortfolio(v bool)`

SetCanAccessPortfolio sets CanAccessPortfolio field to given value.

### HasCanAccessPortfolio

`func (o *PeoplePermissions) HasCanAccessPortfolio() bool`

HasCanAccessPortfolio returns a boolean if a field has been set.

### GetCanAccessTemplates

`func (o *PeoplePermissions) GetCanAccessTemplates() bool`

GetCanAccessTemplates returns the CanAccessTemplates field if non-nil, zero value otherwise.

### GetCanAccessTemplatesOk

`func (o *PeoplePermissions) GetCanAccessTemplatesOk() (*bool, bool)`

GetCanAccessTemplatesOk returns a tuple with the CanAccessTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAccessTemplates

`func (o *PeoplePermissions) SetCanAccessTemplates(v bool)`

SetCanAccessTemplates sets CanAccessTemplates field to given value.

### HasCanAccessTemplates

`func (o *PeoplePermissions) HasCanAccessTemplates() bool`

HasCanAccessTemplates returns a boolean if a field has been set.

### GetCanAddForms

`func (o *PeoplePermissions) GetCanAddForms() bool`

GetCanAddForms returns the CanAddForms field if non-nil, zero value otherwise.

### GetCanAddFormsOk

`func (o *PeoplePermissions) GetCanAddFormsOk() (*bool, bool)`

GetCanAddFormsOk returns a tuple with the CanAddForms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAddForms

`func (o *PeoplePermissions) SetCanAddForms(v bool)`

SetCanAddForms sets CanAddForms field to given value.

### HasCanAddForms

`func (o *PeoplePermissions) HasCanAddForms() bool`

HasCanAddForms returns a boolean if a field has been set.

### GetCanAddProjects

`func (o *PeoplePermissions) GetCanAddProjects() bool`

GetCanAddProjects returns the CanAddProjects field if non-nil, zero value otherwise.

### GetCanAddProjectsOk

`func (o *PeoplePermissions) GetCanAddProjectsOk() (*bool, bool)`

GetCanAddProjectsOk returns a tuple with the CanAddProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAddProjects

`func (o *PeoplePermissions) SetCanAddProjects(v bool)`

SetCanAddProjects sets CanAddProjects field to given value.

### HasCanAddProjects

`func (o *PeoplePermissions) HasCanAddProjects() bool`

HasCanAddProjects returns a boolean if a field has been set.

### GetCanBeAssignedToTasksAndMilestones

`func (o *PeoplePermissions) GetCanBeAssignedToTasksAndMilestones() bool`

GetCanBeAssignedToTasksAndMilestones returns the CanBeAssignedToTasksAndMilestones field if non-nil, zero value otherwise.

### GetCanBeAssignedToTasksAndMilestonesOk

`func (o *PeoplePermissions) GetCanBeAssignedToTasksAndMilestonesOk() (*bool, bool)`

GetCanBeAssignedToTasksAndMilestonesOk returns a tuple with the CanBeAssignedToTasksAndMilestones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanBeAssignedToTasksAndMilestones

`func (o *PeoplePermissions) SetCanBeAssignedToTasksAndMilestones(v bool)`

SetCanBeAssignedToTasksAndMilestones sets CanBeAssignedToTasksAndMilestones field to given value.

### HasCanBeAssignedToTasksAndMilestones

`func (o *PeoplePermissions) HasCanBeAssignedToTasksAndMilestones() bool`

HasCanBeAssignedToTasksAndMilestones returns a boolean if a field has been set.

### GetCanManageCustomFields

`func (o *PeoplePermissions) GetCanManageCustomFields() bool`

GetCanManageCustomFields returns the CanManageCustomFields field if non-nil, zero value otherwise.

### GetCanManageCustomFieldsOk

`func (o *PeoplePermissions) GetCanManageCustomFieldsOk() (*bool, bool)`

GetCanManageCustomFieldsOk returns a tuple with the CanManageCustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageCustomFields

`func (o *PeoplePermissions) SetCanManageCustomFields(v bool)`

SetCanManageCustomFields sets CanManageCustomFields field to given value.

### HasCanManageCustomFields

`func (o *PeoplePermissions) HasCanManageCustomFields() bool`

HasCanManageCustomFields returns a boolean if a field has been set.

### GetCanManagePeople

`func (o *PeoplePermissions) GetCanManagePeople() bool`

GetCanManagePeople returns the CanManagePeople field if non-nil, zero value otherwise.

### GetCanManagePeopleOk

`func (o *PeoplePermissions) GetCanManagePeopleOk() (*bool, bool)`

GetCanManagePeopleOk returns a tuple with the CanManagePeople field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManagePeople

`func (o *PeoplePermissions) SetCanManagePeople(v bool)`

SetCanManagePeople sets CanManagePeople field to given value.

### HasCanManagePeople

`func (o *PeoplePermissions) HasCanManagePeople() bool`

HasCanManagePeople returns a boolean if a field has been set.

### GetCanManagePortfolio

`func (o *PeoplePermissions) GetCanManagePortfolio() bool`

GetCanManagePortfolio returns the CanManagePortfolio field if non-nil, zero value otherwise.

### GetCanManagePortfolioOk

`func (o *PeoplePermissions) GetCanManagePortfolioOk() (*bool, bool)`

GetCanManagePortfolioOk returns a tuple with the CanManagePortfolio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManagePortfolio

`func (o *PeoplePermissions) SetCanManagePortfolio(v bool)`

SetCanManagePortfolio sets CanManagePortfolio field to given value.

### HasCanManagePortfolio

`func (o *PeoplePermissions) HasCanManagePortfolio() bool`

HasCanManagePortfolio returns a boolean if a field has been set.

### GetCanManageProjectBudget

`func (o *PeoplePermissions) GetCanManageProjectBudget() bool`

GetCanManageProjectBudget returns the CanManageProjectBudget field if non-nil, zero value otherwise.

### GetCanManageProjectBudgetOk

`func (o *PeoplePermissions) GetCanManageProjectBudgetOk() (*bool, bool)`

GetCanManageProjectBudgetOk returns a tuple with the CanManageProjectBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageProjectBudget

`func (o *PeoplePermissions) SetCanManageProjectBudget(v bool)`

SetCanManageProjectBudget sets CanManageProjectBudget field to given value.

### HasCanManageProjectBudget

`func (o *PeoplePermissions) HasCanManageProjectBudget() bool`

HasCanManageProjectBudget returns a boolean if a field has been set.

### GetCanManageProjectTemplates

`func (o *PeoplePermissions) GetCanManageProjectTemplates() bool`

GetCanManageProjectTemplates returns the CanManageProjectTemplates field if non-nil, zero value otherwise.

### GetCanManageProjectTemplatesOk

`func (o *PeoplePermissions) GetCanManageProjectTemplatesOk() (*bool, bool)`

GetCanManageProjectTemplatesOk returns a tuple with the CanManageProjectTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageProjectTemplates

`func (o *PeoplePermissions) SetCanManageProjectTemplates(v bool)`

SetCanManageProjectTemplates sets CanManageProjectTemplates field to given value.

### HasCanManageProjectTemplates

`func (o *PeoplePermissions) HasCanManageProjectTemplates() bool`

HasCanManageProjectTemplates returns a boolean if a field has been set.

### GetCanManageRates

`func (o *PeoplePermissions) GetCanManageRates() bool`

GetCanManageRates returns the CanManageRates field if non-nil, zero value otherwise.

### GetCanManageRatesOk

`func (o *PeoplePermissions) GetCanManageRatesOk() (*bool, bool)`

GetCanManageRatesOk returns a tuple with the CanManageRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageRates

`func (o *PeoplePermissions) SetCanManageRates(v bool)`

SetCanManageRates sets CanManageRates field to given value.

### HasCanManageRates

`func (o *PeoplePermissions) HasCanManageRates() bool`

HasCanManageRates returns a boolean if a field has been set.

### GetCanManageReports

`func (o *PeoplePermissions) GetCanManageReports() bool`

GetCanManageReports returns the CanManageReports field if non-nil, zero value otherwise.

### GetCanManageReportsOk

`func (o *PeoplePermissions) GetCanManageReportsOk() (*bool, bool)`

GetCanManageReportsOk returns a tuple with the CanManageReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageReports

`func (o *PeoplePermissions) SetCanManageReports(v bool)`

SetCanManageReports sets CanManageReports field to given value.

### HasCanManageReports

`func (o *PeoplePermissions) HasCanManageReports() bool`

HasCanManageReports returns a boolean if a field has been set.

### GetCanManageSchedule

`func (o *PeoplePermissions) GetCanManageSchedule() bool`

GetCanManageSchedule returns the CanManageSchedule field if non-nil, zero value otherwise.

### GetCanManageScheduleOk

`func (o *PeoplePermissions) GetCanManageScheduleOk() (*bool, bool)`

GetCanManageScheduleOk returns a tuple with the CanManageSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageSchedule

`func (o *PeoplePermissions) SetCanManageSchedule(v bool)`

SetCanManageSchedule sets CanManageSchedule field to given value.

### HasCanManageSchedule

`func (o *PeoplePermissions) HasCanManageSchedule() bool`

HasCanManageSchedule returns a boolean if a field has been set.

### GetCanReceiveEmail

`func (o *PeoplePermissions) GetCanReceiveEmail() bool`

GetCanReceiveEmail returns the CanReceiveEmail field if non-nil, zero value otherwise.

### GetCanReceiveEmailOk

`func (o *PeoplePermissions) GetCanReceiveEmailOk() (*bool, bool)`

GetCanReceiveEmailOk returns a tuple with the CanReceiveEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanReceiveEmail

`func (o *PeoplePermissions) SetCanReceiveEmail(v bool)`

SetCanReceiveEmail sets CanReceiveEmail field to given value.

### HasCanReceiveEmail

`func (o *PeoplePermissions) HasCanReceiveEmail() bool`

HasCanReceiveEmail returns a boolean if a field has been set.

### GetCanViewForms

`func (o *PeoplePermissions) GetCanViewForms() bool`

GetCanViewForms returns the CanViewForms field if non-nil, zero value otherwise.

### GetCanViewFormsOk

`func (o *PeoplePermissions) GetCanViewFormsOk() (*bool, bool)`

GetCanViewFormsOk returns a tuple with the CanViewForms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewForms

`func (o *PeoplePermissions) SetCanViewForms(v bool)`

SetCanViewForms sets CanViewForms field to given value.

### HasCanViewForms

`func (o *PeoplePermissions) HasCanViewForms() bool`

HasCanViewForms returns a boolean if a field has been set.

### GetCanViewProjectBudget

`func (o *PeoplePermissions) GetCanViewProjectBudget() bool`

GetCanViewProjectBudget returns the CanViewProjectBudget field if non-nil, zero value otherwise.

### GetCanViewProjectBudgetOk

`func (o *PeoplePermissions) GetCanViewProjectBudgetOk() (*bool, bool)`

GetCanViewProjectBudgetOk returns a tuple with the CanViewProjectBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewProjectBudget

`func (o *PeoplePermissions) SetCanViewProjectBudget(v bool)`

SetCanViewProjectBudget sets CanViewProjectBudget field to given value.

### HasCanViewProjectBudget

`func (o *PeoplePermissions) HasCanViewProjectBudget() bool`

HasCanViewProjectBudget returns a boolean if a field has been set.

### GetCanViewProjectMembers

`func (o *PeoplePermissions) GetCanViewProjectMembers() bool`

GetCanViewProjectMembers returns the CanViewProjectMembers field if non-nil, zero value otherwise.

### GetCanViewProjectMembersOk

`func (o *PeoplePermissions) GetCanViewProjectMembersOk() (*bool, bool)`

GetCanViewProjectMembersOk returns a tuple with the CanViewProjectMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewProjectMembers

`func (o *PeoplePermissions) SetCanViewProjectMembers(v bool)`

SetCanViewProjectMembers sets CanViewProjectMembers field to given value.

### HasCanViewProjectMembers

`func (o *PeoplePermissions) HasCanViewProjectMembers() bool`

HasCanViewProjectMembers returns a boolean if a field has been set.

### GetCanViewProjectTemplates

`func (o *PeoplePermissions) GetCanViewProjectTemplates() bool`

GetCanViewProjectTemplates returns the CanViewProjectTemplates field if non-nil, zero value otherwise.

### GetCanViewProjectTemplatesOk

`func (o *PeoplePermissions) GetCanViewProjectTemplatesOk() (*bool, bool)`

GetCanViewProjectTemplatesOk returns a tuple with the CanViewProjectTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewProjectTemplates

`func (o *PeoplePermissions) SetCanViewProjectTemplates(v bool)`

SetCanViewProjectTemplates sets CanViewProjectTemplates field to given value.

### HasCanViewProjectTemplates

`func (o *PeoplePermissions) HasCanViewProjectTemplates() bool`

HasCanViewProjectTemplates returns a boolean if a field has been set.

### GetCanViewRates

`func (o *PeoplePermissions) GetCanViewRates() bool`

GetCanViewRates returns the CanViewRates field if non-nil, zero value otherwise.

### GetCanViewRatesOk

`func (o *PeoplePermissions) GetCanViewRatesOk() (*bool, bool)`

GetCanViewRatesOk returns a tuple with the CanViewRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewRates

`func (o *PeoplePermissions) SetCanViewRates(v bool)`

SetCanViewRates sets CanViewRates field to given value.

### HasCanViewRates

`func (o *PeoplePermissions) HasCanViewRates() bool`

HasCanViewRates returns a boolean if a field has been set.

### GetCanViewReports

`func (o *PeoplePermissions) GetCanViewReports() bool`

GetCanViewReports returns the CanViewReports field if non-nil, zero value otherwise.

### GetCanViewReportsOk

`func (o *PeoplePermissions) GetCanViewReportsOk() (*bool, bool)`

GetCanViewReportsOk returns a tuple with the CanViewReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewReports

`func (o *PeoplePermissions) SetCanViewReports(v bool)`

SetCanViewReports sets CanViewReports field to given value.

### HasCanViewReports

`func (o *PeoplePermissions) HasCanViewReports() bool`

HasCanViewReports returns a boolean if a field has been set.

### GetCanViewSchedule

`func (o *PeoplePermissions) GetCanViewSchedule() bool`

GetCanViewSchedule returns the CanViewSchedule field if non-nil, zero value otherwise.

### GetCanViewScheduleOk

`func (o *PeoplePermissions) GetCanViewScheduleOk() (*bool, bool)`

GetCanViewScheduleOk returns a tuple with the CanViewSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewSchedule

`func (o *PeoplePermissions) SetCanViewSchedule(v bool)`

SetCanViewSchedule sets CanViewSchedule field to given value.

### HasCanViewSchedule

`func (o *PeoplePermissions) HasCanViewSchedule() bool`

HasCanViewSchedule returns a boolean if a field has been set.

### GetEditAllTasks

`func (o *PeoplePermissions) GetEditAllTasks() bool`

GetEditAllTasks returns the EditAllTasks field if non-nil, zero value otherwise.

### GetEditAllTasksOk

`func (o *PeoplePermissions) GetEditAllTasksOk() (*bool, bool)`

GetEditAllTasksOk returns a tuple with the EditAllTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditAllTasks

`func (o *PeoplePermissions) SetEditAllTasks(v bool)`

SetEditAllTasks sets EditAllTasks field to given value.

### HasEditAllTasks

`func (o *PeoplePermissions) HasEditAllTasks() bool`

HasEditAllTasks returns a boolean if a field has been set.

### GetHasAccessToNewProjects

`func (o *PeoplePermissions) GetHasAccessToNewProjects() bool`

GetHasAccessToNewProjects returns the HasAccessToNewProjects field if non-nil, zero value otherwise.

### GetHasAccessToNewProjectsOk

`func (o *PeoplePermissions) GetHasAccessToNewProjectsOk() (*bool, bool)`

GetHasAccessToNewProjectsOk returns a tuple with the HasAccessToNewProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccessToNewProjects

`func (o *PeoplePermissions) SetHasAccessToNewProjects(v bool)`

SetHasAccessToNewProjects sets HasAccessToNewProjects field to given value.

### HasHasAccessToNewProjects

`func (o *PeoplePermissions) HasHasAccessToNewProjects() bool`

HasHasAccessToNewProjects returns a boolean if a field has been set.

### GetIsArchived

`func (o *PeoplePermissions) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *PeoplePermissions) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *PeoplePermissions) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *PeoplePermissions) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### GetIsObserving

`func (o *PeoplePermissions) GetIsObserving() bool`

GetIsObserving returns the IsObserving field if non-nil, zero value otherwise.

### GetIsObservingOk

`func (o *PeoplePermissions) GetIsObservingOk() (*bool, bool)`

GetIsObservingOk returns a tuple with the IsObserving field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsObserving

`func (o *PeoplePermissions) SetIsObserving(v bool)`

SetIsObserving sets IsObserving field to given value.

### HasIsObserving

`func (o *PeoplePermissions) HasIsObserving() bool`

HasIsObserving returns a boolean if a field has been set.

### GetManageCustomFields

`func (o *PeoplePermissions) GetManageCustomFields() bool`

GetManageCustomFields returns the ManageCustomFields field if non-nil, zero value otherwise.

### GetManageCustomFieldsOk

`func (o *PeoplePermissions) GetManageCustomFieldsOk() (*bool, bool)`

GetManageCustomFieldsOk returns a tuple with the ManageCustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageCustomFields

`func (o *PeoplePermissions) SetManageCustomFields(v bool)`

SetManageCustomFields sets ManageCustomFields field to given value.

### HasManageCustomFields

`func (o *PeoplePermissions) HasManageCustomFields() bool`

HasManageCustomFields returns a boolean if a field has been set.

### GetNotifyDefaults

`func (o *PeoplePermissions) GetNotifyDefaults() PeopleNotifyDefaults`

GetNotifyDefaults returns the NotifyDefaults field if non-nil, zero value otherwise.

### GetNotifyDefaultsOk

`func (o *PeoplePermissions) GetNotifyDefaultsOk() (*PeopleNotifyDefaults, bool)`

GetNotifyDefaultsOk returns a tuple with the NotifyDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyDefaults

`func (o *PeoplePermissions) SetNotifyDefaults(v PeopleNotifyDefaults)`

SetNotifyDefaults sets NotifyDefaults field to given value.

### HasNotifyDefaults

`func (o *PeoplePermissions) HasNotifyDefaults() bool`

HasNotifyDefaults returns a boolean if a field has been set.

### GetProjectAdministrator

`func (o *PeoplePermissions) GetProjectAdministrator() bool`

GetProjectAdministrator returns the ProjectAdministrator field if non-nil, zero value otherwise.

### GetProjectAdministratorOk

`func (o *PeoplePermissions) GetProjectAdministratorOk() (*bool, bool)`

GetProjectAdministratorOk returns a tuple with the ProjectAdministrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectAdministrator

`func (o *PeoplePermissions) SetProjectAdministrator(v bool)`

SetProjectAdministrator sets ProjectAdministrator field to given value.

### HasProjectAdministrator

`func (o *PeoplePermissions) HasProjectAdministrator() bool`

HasProjectAdministrator returns a boolean if a field has been set.

### GetSetPrivacy

`func (o *PeoplePermissions) GetSetPrivacy() bool`

GetSetPrivacy returns the SetPrivacy field if non-nil, zero value otherwise.

### GetSetPrivacyOk

`func (o *PeoplePermissions) GetSetPrivacyOk() (*bool, bool)`

GetSetPrivacyOk returns a tuple with the SetPrivacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetPrivacy

`func (o *PeoplePermissions) SetSetPrivacy(v bool)`

SetSetPrivacy sets SetPrivacy field to given value.

### HasSetPrivacy

`func (o *PeoplePermissions) HasSetPrivacy() bool`

HasSetPrivacy returns a boolean if a field has been set.

### GetViewAllTimeLogs

`func (o *PeoplePermissions) GetViewAllTimeLogs() bool`

GetViewAllTimeLogs returns the ViewAllTimeLogs field if non-nil, zero value otherwise.

### GetViewAllTimeLogsOk

`func (o *PeoplePermissions) GetViewAllTimeLogsOk() (*bool, bool)`

GetViewAllTimeLogsOk returns a tuple with the ViewAllTimeLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewAllTimeLogs

`func (o *PeoplePermissions) SetViewAllTimeLogs(v bool)`

SetViewAllTimeLogs sets ViewAllTimeLogs field to given value.

### HasViewAllTimeLogs

`func (o *PeoplePermissions) HasViewAllTimeLogs() bool`

HasViewAllTimeLogs returns a boolean if a field has been set.

### GetViewEstimatedTime

`func (o *PeoplePermissions) GetViewEstimatedTime() bool`

GetViewEstimatedTime returns the ViewEstimatedTime field if non-nil, zero value otherwise.

### GetViewEstimatedTimeOk

`func (o *PeoplePermissions) GetViewEstimatedTimeOk() (*bool, bool)`

GetViewEstimatedTimeOk returns a tuple with the ViewEstimatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewEstimatedTime

`func (o *PeoplePermissions) SetViewEstimatedTime(v bool)`

SetViewEstimatedTime sets ViewEstimatedTime field to given value.

### HasViewEstimatedTime

`func (o *PeoplePermissions) HasViewEstimatedTime() bool`

HasViewEstimatedTime returns a boolean if a field has been set.

### GetViewInvoices

`func (o *PeoplePermissions) GetViewInvoices() bool`

GetViewInvoices returns the ViewInvoices field if non-nil, zero value otherwise.

### GetViewInvoicesOk

`func (o *PeoplePermissions) GetViewInvoicesOk() (*bool, bool)`

GetViewInvoicesOk returns a tuple with the ViewInvoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewInvoices

`func (o *PeoplePermissions) SetViewInvoices(v bool)`

SetViewInvoices sets ViewInvoices field to given value.

### HasViewInvoices

`func (o *PeoplePermissions) HasViewInvoices() bool`

HasViewInvoices returns a boolean if a field has been set.

### GetViewLinks

`func (o *PeoplePermissions) GetViewLinks() bool`

GetViewLinks returns the ViewLinks field if non-nil, zero value otherwise.

### GetViewLinksOk

`func (o *PeoplePermissions) GetViewLinksOk() (*bool, bool)`

GetViewLinksOk returns a tuple with the ViewLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewLinks

`func (o *PeoplePermissions) SetViewLinks(v bool)`

SetViewLinks sets ViewLinks field to given value.

### HasViewLinks

`func (o *PeoplePermissions) HasViewLinks() bool`

HasViewLinks returns a boolean if a field has been set.

### GetViewMessagesAndFiles

`func (o *PeoplePermissions) GetViewMessagesAndFiles() bool`

GetViewMessagesAndFiles returns the ViewMessagesAndFiles field if non-nil, zero value otherwise.

### GetViewMessagesAndFilesOk

`func (o *PeoplePermissions) GetViewMessagesAndFilesOk() (*bool, bool)`

GetViewMessagesAndFilesOk returns a tuple with the ViewMessagesAndFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewMessagesAndFiles

`func (o *PeoplePermissions) SetViewMessagesAndFiles(v bool)`

SetViewMessagesAndFiles sets ViewMessagesAndFiles field to given value.

### HasViewMessagesAndFiles

`func (o *PeoplePermissions) HasViewMessagesAndFiles() bool`

HasViewMessagesAndFiles returns a boolean if a field has been set.

### GetViewNotebooks

`func (o *PeoplePermissions) GetViewNotebooks() bool`

GetViewNotebooks returns the ViewNotebooks field if non-nil, zero value otherwise.

### GetViewNotebooksOk

`func (o *PeoplePermissions) GetViewNotebooksOk() (*bool, bool)`

GetViewNotebooksOk returns a tuple with the ViewNotebooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewNotebooks

`func (o *PeoplePermissions) SetViewNotebooks(v bool)`

SetViewNotebooks sets ViewNotebooks field to given value.

### HasViewNotebooks

`func (o *PeoplePermissions) HasViewNotebooks() bool`

HasViewNotebooks returns a boolean if a field has been set.

### GetViewProjectUpdate

`func (o *PeoplePermissions) GetViewProjectUpdate() bool`

GetViewProjectUpdate returns the ViewProjectUpdate field if non-nil, zero value otherwise.

### GetViewProjectUpdateOk

`func (o *PeoplePermissions) GetViewProjectUpdateOk() (*bool, bool)`

GetViewProjectUpdateOk returns a tuple with the ViewProjectUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewProjectUpdate

`func (o *PeoplePermissions) SetViewProjectUpdate(v bool)`

SetViewProjectUpdate sets ViewProjectUpdate field to given value.

### HasViewProjectUpdate

`func (o *PeoplePermissions) HasViewProjectUpdate() bool`

HasViewProjectUpdate returns a boolean if a field has been set.

### GetViewRiskRegister

`func (o *PeoplePermissions) GetViewRiskRegister() bool`

GetViewRiskRegister returns the ViewRiskRegister field if non-nil, zero value otherwise.

### GetViewRiskRegisterOk

`func (o *PeoplePermissions) GetViewRiskRegisterOk() (*bool, bool)`

GetViewRiskRegisterOk returns a tuple with the ViewRiskRegister field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRiskRegister

`func (o *PeoplePermissions) SetViewRiskRegister(v bool)`

SetViewRiskRegister sets ViewRiskRegister field to given value.

### HasViewRiskRegister

`func (o *PeoplePermissions) HasViewRiskRegister() bool`

HasViewRiskRegister returns a boolean if a field has been set.

### GetViewTasksAndMilestones

`func (o *PeoplePermissions) GetViewTasksAndMilestones() bool`

GetViewTasksAndMilestones returns the ViewTasksAndMilestones field if non-nil, zero value otherwise.

### GetViewTasksAndMilestonesOk

`func (o *PeoplePermissions) GetViewTasksAndMilestonesOk() (*bool, bool)`

GetViewTasksAndMilestonesOk returns a tuple with the ViewTasksAndMilestones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewTasksAndMilestones

`func (o *PeoplePermissions) SetViewTasksAndMilestones(v bool)`

SetViewTasksAndMilestones sets ViewTasksAndMilestones field to given value.

### HasViewTasksAndMilestones

`func (o *PeoplePermissions) HasViewTasksAndMilestones() bool`

HasViewTasksAndMilestones returns a boolean if a field has been set.

### GetViewTime

`func (o *PeoplePermissions) GetViewTime() bool`

GetViewTime returns the ViewTime field if non-nil, zero value otherwise.

### GetViewTimeOk

`func (o *PeoplePermissions) GetViewTimeOk() (*bool, bool)`

GetViewTimeOk returns a tuple with the ViewTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewTime

`func (o *PeoplePermissions) SetViewTime(v bool)`

SetViewTime sets ViewTime field to given value.

### HasViewTime

`func (o *PeoplePermissions) HasViewTime() bool`

HasViewTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


