# \WorkloadApi

All URIs are relative to *https://yoursite.teamwork.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETWorkloadJson**](WorkloadApi.md#GETWorkloadJson) | **Get** /workload.json | Get Workload



## GETWorkloadJson

> InlineResponse200120 GETWorkloadJson(ctx).StartDate(startDate).EndDate(endDate).IncludeTasksWithoutDates(includeTasksWithoutDates).DistributeEstimatedTimeToAssignees(distributeEstimatedTimeToAssignees).ProrataEstimatedTime(prorataEstimatedTime).Page(page).PageSize(pageSize).SortBy(sortBy).TagIds(tagIds).MatchAllTags(matchAllTags).ExcludeTagIds(excludeTagIds).MatchAllExcludedTags(matchAllExcludedTags).OnlyUntaggedTasks(onlyUntaggedTasks).IncludeCapacity(includeCapacity).ShowOnlyUsersWithRemainingCapacity(showOnlyUsersWithRemainingCapacity).IncludeAllUsers(includeAllUsers).UserIdsToInclude(userIdsToInclude).UserCompanyIds(userCompanyIds).UserSortBy(userSortBy).SearchUserName(searchUserName).Execute()

Get Workload



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    startDate := "startDate_example" // string | Start date in the format of yyyymmdd
    endDate := "endDate_example" // string | End date in the format of yyyymmdd
    includeTasksWithoutDates := true // bool | Include tasks without Due dates (optional) (default to false)
    distributeEstimatedTimeToAssignees := true // bool |  (optional) (default to false)
    prorataEstimatedTime := "prorataEstimatedTime_example" // string |  (optional)
    page := int32(56) // int32 | The Page number to view (optional)
    pageSize := int32(56) // int32 | The number of results per page (optional)
    sortBy := "sortBy_example" // string | Sort by 'users', 'company' or 'project'  (optional)
    tagIds := int32(56) // int32 |  (optional)
    matchAllTags := true // bool |  (optional)
    excludeTagIds := int32(56) // int32 |  (optional)
    matchAllExcludedTags := true // bool |  (optional)
    onlyUntaggedTasks := true // bool |  (optional)
    includeCapacity := true // bool | When sorting by user, returns a new capacity field for each user that represents the percentage of workload compared to the available time (optional)
    showOnlyUsersWithRemainingCapacity := true // bool | Only return users where capacity is bellow 100%, so we can detect the users that can take more tasks (optional)
    includeAllUsers := true // bool | Side load all other users that we can assign a task in the installation (optional)
    userIdsToInclude := "userIdsToInclude_example" // string | Always include this users capacity, even if not matching the current filter (optional) (default to "list of numeric numbers")
    userCompanyIds := "userCompanyIds_example" // string | Filter by users' company IDs (optional) (default to "List of numeric numbers")
    userSortBy := "userSortBy_example" // string | After sorting by user, this allows to sort the users by sub-categories (optional) (default to "Options 'name', 'company' and 'capacity'")
    searchUserName := "searchUserName_example" // string | Filter by the user names (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadApi.GETWorkloadJson(context.Background()).StartDate(startDate).EndDate(endDate).IncludeTasksWithoutDates(includeTasksWithoutDates).DistributeEstimatedTimeToAssignees(distributeEstimatedTimeToAssignees).ProrataEstimatedTime(prorataEstimatedTime).Page(page).PageSize(pageSize).SortBy(sortBy).TagIds(tagIds).MatchAllTags(matchAllTags).ExcludeTagIds(excludeTagIds).MatchAllExcludedTags(matchAllExcludedTags).OnlyUntaggedTasks(onlyUntaggedTasks).IncludeCapacity(includeCapacity).ShowOnlyUsersWithRemainingCapacity(showOnlyUsersWithRemainingCapacity).IncludeAllUsers(includeAllUsers).UserIdsToInclude(userIdsToInclude).UserCompanyIds(userCompanyIds).UserSortBy(userSortBy).SearchUserName(searchUserName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadApi.GETWorkloadJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETWorkloadJson`: InlineResponse200120
    fmt.Fprintf(os.Stdout, "Response from `WorkloadApi.GETWorkloadJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETWorkloadJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | Start date in the format of yyyymmdd | 
 **endDate** | **string** | End date in the format of yyyymmdd | 
 **includeTasksWithoutDates** | **bool** | Include tasks without Due dates | [default to false]
 **distributeEstimatedTimeToAssignees** | **bool** |  | [default to false]
 **prorataEstimatedTime** | **string** |  | 
 **page** | **int32** | The Page number to view | 
 **pageSize** | **int32** | The number of results per page | 
 **sortBy** | **string** | Sort by &#39;users&#39;, &#39;company&#39; or &#39;project&#39;  | 
 **tagIds** | **int32** |  | 
 **matchAllTags** | **bool** |  | 
 **excludeTagIds** | **int32** |  | 
 **matchAllExcludedTags** | **bool** |  | 
 **onlyUntaggedTasks** | **bool** |  | 
 **includeCapacity** | **bool** | When sorting by user, returns a new capacity field for each user that represents the percentage of workload compared to the available time | 
 **showOnlyUsersWithRemainingCapacity** | **bool** | Only return users where capacity is bellow 100%, so we can detect the users that can take more tasks | 
 **includeAllUsers** | **bool** | Side load all other users that we can assign a task in the installation | 
 **userIdsToInclude** | **string** | Always include this users capacity, even if not matching the current filter | [default to &quot;list of numeric numbers&quot;]
 **userCompanyIds** | **string** | Filter by users&#39; company IDs | [default to &quot;List of numeric numbers&quot;]
 **userSortBy** | **string** | After sorting by user, this allows to sort the users by sub-categories | [default to &quot;Options &#39;name&#39;, &#39;company&#39; and &#39;capacity&#39;&quot;]
 **searchUserName** | **string** | Filter by the user names | 

### Return type

[**InlineResponse200120**](InlineResponse200120.md)

### Authorization

[YOUR_API_KEY](../README.md#YOUR_API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

