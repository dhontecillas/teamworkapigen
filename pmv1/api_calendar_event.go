/*
 * Teamwork Projects API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pmv1

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// CalendarEventApiService CalendarEventApi service
type CalendarEventApiService service

type ApiDELETECalendareventsIdJsonRequest struct {
	ctx _context.Context
	ApiService *CalendarEventApiService
	id string
	body *InlineObject6
}

func (r ApiDELETECalendareventsIdJsonRequest) Body(body InlineObject6) ApiDELETECalendareventsIdJsonRequest {
	r.body = &body
	return r
}

func (r ApiDELETECalendareventsIdJsonRequest) Execute() (InlineResponse2001, *_nethttp.Response, error) {
	return r.ApiService.DELETECalendareventsIdJsonExecute(r)
}

/*
 * DELETECalendareventsIdJson Delete Event
 * Delete a calendar event.

---
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiDELETECalendareventsIdJsonRequest
 */
func (a *CalendarEventApiService) DELETECalendareventsIdJson(ctx _context.Context, id string) ApiDELETECalendareventsIdJsonRequest {
	return ApiDELETECalendareventsIdJsonRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse2001
 */
func (a *CalendarEventApiService) DELETECalendareventsIdJsonExecute(r ApiDELETECalendareventsIdJsonRequest) (InlineResponse2001, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse2001
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarEventApiService.DELETECalendareventsIdJson")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/calendarevents/{id}.json"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDELETEEventtypesIdJsonRequest struct {
	ctx _context.Context
	ApiService *CalendarEventApiService
	id string
}


func (r ApiDELETEEventtypesIdJsonRequest) Execute() (InlineResponse2001, *_nethttp.Response, error) {
	return r.ApiService.DELETEEventtypesIdJsonExecute(r)
}

/*
 * DELETEEventtypesIdJson Delete Event Type
 * Delete a calendar event type.

---
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiDELETEEventtypesIdJsonRequest
 */
func (a *CalendarEventApiService) DELETEEventtypesIdJson(ctx _context.Context, id string) ApiDELETEEventtypesIdJsonRequest {
	return ApiDELETEEventtypesIdJsonRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse2001
 */
func (a *CalendarEventApiService) DELETEEventtypesIdJsonExecute(r ApiDELETEEventtypesIdJsonRequest) (InlineResponse2001, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse2001
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarEventApiService.DELETEEventtypesIdJson")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/eventtypes/{id}.json"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGETCalendareventsIdJsonRequest struct {
	ctx _context.Context
	ApiService *CalendarEventApiService
	id string
}


func (r ApiGETCalendareventsIdJsonRequest) Execute() (InlineResponse2005, *_nethttp.Response, error) {
	return r.ApiService.GETCalendareventsIdJsonExecute(r)
}

/*
 * GETCalendareventsIdJson Get an Event
 * This call will return all calendar events that the current user can see in the provided date range.

---
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiGETCalendareventsIdJsonRequest
 */
func (a *CalendarEventApiService) GETCalendareventsIdJson(ctx _context.Context, id string) ApiGETCalendareventsIdJsonRequest {
	return ApiGETCalendareventsIdJsonRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse2005
 */
func (a *CalendarEventApiService) GETCalendareventsIdJsonExecute(r ApiGETCalendareventsIdJsonRequest) (InlineResponse2005, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse2005
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarEventApiService.GETCalendareventsIdJson")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/calendarevents/{id}.json"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGETCalendareventsJsonRequest struct {
	ctx _context.Context
	ApiService *CalendarEventApiService
	startdate *string
	endDate *string
	showDeleted *bool
	updatedAfterDate *string
	eventTypeId *int32
	page *int32
	userId *float32
	attendingOnly *bool
}

func (r ApiGETCalendareventsJsonRequest) Startdate(startdate string) ApiGETCalendareventsJsonRequest {
	r.startdate = &startdate
	return r
}
func (r ApiGETCalendareventsJsonRequest) EndDate(endDate string) ApiGETCalendareventsJsonRequest {
	r.endDate = &endDate
	return r
}
func (r ApiGETCalendareventsJsonRequest) ShowDeleted(showDeleted bool) ApiGETCalendareventsJsonRequest {
	r.showDeleted = &showDeleted
	return r
}
func (r ApiGETCalendareventsJsonRequest) UpdatedAfterDate(updatedAfterDate string) ApiGETCalendareventsJsonRequest {
	r.updatedAfterDate = &updatedAfterDate
	return r
}
func (r ApiGETCalendareventsJsonRequest) EventTypeId(eventTypeId int32) ApiGETCalendareventsJsonRequest {
	r.eventTypeId = &eventTypeId
	return r
}
func (r ApiGETCalendareventsJsonRequest) Page(page int32) ApiGETCalendareventsJsonRequest {
	r.page = &page
	return r
}
func (r ApiGETCalendareventsJsonRequest) UserId(userId float32) ApiGETCalendareventsJsonRequest {
	r.userId = &userId
	return r
}
func (r ApiGETCalendareventsJsonRequest) AttendingOnly(attendingOnly bool) ApiGETCalendareventsJsonRequest {
	r.attendingOnly = &attendingOnly
	return r
}

func (r ApiGETCalendareventsJsonRequest) Execute() (InlineResponse2003, *_nethttp.Response, error) {
	return r.ApiService.GETCalendareventsJsonExecute(r)
}

/*
 * GETCalendareventsJson Get Events
 * This call will return all calendar events that the current user can see in the provided date range.

<h3>Notes</h3>
<h4>Pagination</h4>

By default we'll send back 250 records per page. Inspect the Response Headers for the following values:

- X-Page (The current page being returned)
- X-Pages (The total number of pages available)
- X-Records (The total number of items available
- You can request a specific page by calling the same API call and adding the parameter &page=n where n is the page you want. e.g: &page=2 for page 2, &page=5 for page 5.

**Please note:**

The start date and end date are required parameters here. If you do not pass in an end date, it will automatically set it as tomorrow's date. 

---
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGETCalendareventsJsonRequest
 */
func (a *CalendarEventApiService) GETCalendareventsJson(ctx _context.Context) ApiGETCalendareventsJsonRequest {
	return ApiGETCalendareventsJsonRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse2003
 */
func (a *CalendarEventApiService) GETCalendareventsJsonExecute(r ApiGETCalendareventsJsonRequest) (InlineResponse2003, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse2003
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarEventApiService.GETCalendareventsJson")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/calendarevents.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startdate == nil {
		return localVarReturnValue, nil, reportError("startdate is required and must be specified")
	}
	if r.endDate == nil {
		return localVarReturnValue, nil, reportError("endDate is required and must be specified")
	}

	localVarQueryParams.Add("startdate", parameterToString(*r.startdate, ""))
	localVarQueryParams.Add("endDate", parameterToString(*r.endDate, ""))
	if r.showDeleted != nil {
		localVarQueryParams.Add("showDeleted", parameterToString(*r.showDeleted, ""))
	}
	if r.updatedAfterDate != nil {
		localVarQueryParams.Add("updatedAfterDate", parameterToString(*r.updatedAfterDate, ""))
	}
	if r.eventTypeId != nil {
		localVarQueryParams.Add("eventTypeId", parameterToString(*r.eventTypeId, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.userId != nil {
		localVarQueryParams.Add("userId", parameterToString(*r.userId, ""))
	}
	if r.attendingOnly != nil {
		localVarQueryParams.Add("attendingOnly", parameterToString(*r.attendingOnly, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGETCalendareventtypesJsonRequest struct {
	ctx _context.Context
	ApiService *CalendarEventApiService
}


func (r ApiGETCalendareventtypesJsonRequest) Execute() (InlineResponse2006, *_nethttp.Response, error) {
	return r.ApiService.GETCalendareventtypesJsonExecute(r)
}

/*
 * GETCalendareventtypesJson Get Event Types
 * This call will return all calendar events that the current user can see in the provided date range.

---
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGETCalendareventtypesJsonRequest
 */
func (a *CalendarEventApiService) GETCalendareventtypesJson(ctx _context.Context) ApiGETCalendareventtypesJsonRequest {
	return ApiGETCalendareventtypesJsonRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse2006
 */
func (a *CalendarEventApiService) GETCalendareventtypesJsonExecute(r ApiGETCalendareventtypesJsonRequest) (InlineResponse2006, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse2006
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarEventApiService.GETCalendareventtypesJson")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/calendareventtypes.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPOSTCalendareventsJsonRequest struct {
	ctx _context.Context
	ApiService *CalendarEventApiService
	body *InlineObject4
}

func (r ApiPOSTCalendareventsJsonRequest) Body(body InlineObject4) ApiPOSTCalendareventsJsonRequest {
	r.body = &body
	return r
}

func (r ApiPOSTCalendareventsJsonRequest) Execute() (InlineResponse201, *_nethttp.Response, error) {
	return r.ApiService.POSTCalendareventsJsonExecute(r)
}

/*
 * POSTCalendareventsJson Create a Recurring Event
 * Create a new recurring calendar event.

Adding the section below to the body of the request will make it a recurring event:
"repeat":
```json
{ "frequency":"monthly", "month-type":"monthday", "ends":true, "endDate":"yyyy-mm-dd" }
```
<h4>Notes</h4>
'frequency' can be one of the following:

- weekly

- weekdays

- monthly
- yearly
- every2weeks
- every3weeks
- every4weeks
- every2months
- every3months
- every4months
- every6months

---
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPOSTCalendareventsJsonRequest
 */
func (a *CalendarEventApiService) POSTCalendareventsJson(ctx _context.Context) ApiPOSTCalendareventsJsonRequest {
	return ApiPOSTCalendareventsJsonRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse201
 */
func (a *CalendarEventApiService) POSTCalendareventsJsonExecute(r ApiPOSTCalendareventsJsonRequest) (InlineResponse201, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse201
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarEventApiService.POSTCalendareventsJson")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/calendarevents.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPOSTEventtypesJsonRequest struct {
	ctx _context.Context
	ApiService *CalendarEventApiService
	body *InlineObject12
}

func (r ApiPOSTEventtypesJsonRequest) Body(body InlineObject12) ApiPOSTEventtypesJsonRequest {
	r.body = &body
	return r
}

func (r ApiPOSTEventtypesJsonRequest) Execute() (InlineResponse201, *_nethttp.Response, error) {
	return r.ApiService.POSTEventtypesJsonExecute(r)
}

/*
 * POSTEventtypesJson Create an Event Type
 * Create an event type.

---
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPOSTEventtypesJsonRequest
 */
func (a *CalendarEventApiService) POSTEventtypesJson(ctx _context.Context) ApiPOSTEventtypesJsonRequest {
	return ApiPOSTEventtypesJsonRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse201
 */
func (a *CalendarEventApiService) POSTEventtypesJsonExecute(r ApiPOSTEventtypesJsonRequest) (InlineResponse201, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse201
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarEventApiService.POSTEventtypesJson")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/eventtypes.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPUTCalendareventsIdJsonRequest struct {
	ctx _context.Context
	ApiService *CalendarEventApiService
	id string
	body *InlineObject5
}

func (r ApiPUTCalendareventsIdJsonRequest) Body(body InlineObject5) ApiPUTCalendareventsIdJsonRequest {
	r.body = &body
	return r
}

func (r ApiPUTCalendareventsIdJsonRequest) Execute() (InlineResponse2001, *_nethttp.Response, error) {
	return r.ApiService.PUTCalendareventsIdJsonExecute(r)
}

/*
 * PUTCalendareventsIdJson Edit an Event
 * ---
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiPUTCalendareventsIdJsonRequest
 */
func (a *CalendarEventApiService) PUTCalendareventsIdJson(ctx _context.Context, id string) ApiPUTCalendareventsIdJsonRequest {
	return ApiPUTCalendareventsIdJsonRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse2001
 */
func (a *CalendarEventApiService) PUTCalendareventsIdJsonExecute(r ApiPUTCalendareventsIdJsonRequest) (InlineResponse2001, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse2001
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarEventApiService.PUTCalendareventsIdJson")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/calendarevents/{id}.json"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPUTEventtypesIdJsonRequest struct {
	ctx _context.Context
	ApiService *CalendarEventApiService
	id string
	body *InlineObject13
}

func (r ApiPUTEventtypesIdJsonRequest) Body(body InlineObject13) ApiPUTEventtypesIdJsonRequest {
	r.body = &body
	return r
}

func (r ApiPUTEventtypesIdJsonRequest) Execute() (InlineResponse2001, *_nethttp.Response, error) {
	return r.ApiService.PUTEventtypesIdJsonExecute(r)
}

/*
 * PUTEventtypesIdJson Edit an Event Type
 * Edit a calendar event type.

---
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiPUTEventtypesIdJsonRequest
 */
func (a *CalendarEventApiService) PUTEventtypesIdJson(ctx _context.Context, id string) ApiPUTEventtypesIdJsonRequest {
	return ApiPUTEventtypesIdJsonRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse2001
 */
func (a *CalendarEventApiService) PUTEventtypesIdJsonExecute(r ApiPUTEventtypesIdJsonRequest) (InlineResponse2001, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse2001
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarEventApiService.PUTEventtypesIdJson")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/eventtypes/{id}.json"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
