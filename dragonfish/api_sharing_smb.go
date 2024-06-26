/*
TrueNAS RESTful API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dragonfish

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// SharingSmbApiService SharingSmbApi service
type SharingSmbApiService service

type ApiSharingSmbGetRequest struct {
	ctx        context.Context
	ApiService *SharingSmbApiService
	limit      *int32
	offset     *int32
	count      *bool
	sort       *string
}

func (r ApiSharingSmbGetRequest) Limit(limit int32) ApiSharingSmbGetRequest {
	r.limit = &limit
	return r
}

func (r ApiSharingSmbGetRequest) Offset(offset int32) ApiSharingSmbGetRequest {
	r.offset = &offset
	return r
}

func (r ApiSharingSmbGetRequest) Count(count bool) ApiSharingSmbGetRequest {
	r.count = &count
	return r
}

func (r ApiSharingSmbGetRequest) Sort(sort string) ApiSharingSmbGetRequest {
	r.sort = &sort
	return r
}

func (r ApiSharingSmbGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.SharingSmbGetExecute(r)
}

/*
SharingSmbGet Method for SharingSmbGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSharingSmbGetRequest
*/
func (a *SharingSmbApiService) SharingSmbGet(ctx context.Context) ApiSharingSmbGetRequest {
	return ApiSharingSmbGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *SharingSmbApiService) SharingSmbGetExecute(r ApiSharingSmbGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharingSmbApiService.SharingSmbGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sharing/smb"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiSharingSmbIdIdDeleteRequest struct {
	ctx        context.Context
	ApiService *SharingSmbApiService
	id         int32
}

func (r ApiSharingSmbIdIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.SharingSmbIdIdDeleteExecute(r)
}

/*
SharingSmbIdIdDelete Method for SharingSmbIdIdDelete

Delete SMB Share of `id`. This will forcibly disconnect SMB clients
that are accessing the share.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiSharingSmbIdIdDeleteRequest
*/
func (a *SharingSmbApiService) SharingSmbIdIdDelete(ctx context.Context, id int32) ApiSharingSmbIdIdDeleteRequest {
	return ApiSharingSmbIdIdDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *SharingSmbApiService) SharingSmbIdIdDeleteExecute(r ApiSharingSmbIdIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharingSmbApiService.SharingSmbIdIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sharing/smb/id/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiSharingSmbIdIdGetRequest struct {
	ctx        context.Context
	ApiService *SharingSmbApiService
	id         int32
	limit      *int32
	offset     *int32
	count      *bool
	sort       *string
}

func (r ApiSharingSmbIdIdGetRequest) Limit(limit int32) ApiSharingSmbIdIdGetRequest {
	r.limit = &limit
	return r
}

func (r ApiSharingSmbIdIdGetRequest) Offset(offset int32) ApiSharingSmbIdIdGetRequest {
	r.offset = &offset
	return r
}

func (r ApiSharingSmbIdIdGetRequest) Count(count bool) ApiSharingSmbIdIdGetRequest {
	r.count = &count
	return r
}

func (r ApiSharingSmbIdIdGetRequest) Sort(sort string) ApiSharingSmbIdIdGetRequest {
	r.sort = &sort
	return r
}

func (r ApiSharingSmbIdIdGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.SharingSmbIdIdGetExecute(r)
}

/*
SharingSmbIdIdGet Method for SharingSmbIdIdGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiSharingSmbIdIdGetRequest
*/
func (a *SharingSmbApiService) SharingSmbIdIdGet(ctx context.Context, id int32) ApiSharingSmbIdIdGetRequest {
	return ApiSharingSmbIdIdGetRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *SharingSmbApiService) SharingSmbIdIdGetExecute(r ApiSharingSmbIdIdGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharingSmbApiService.SharingSmbIdIdGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sharing/smb/id/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiSharingSmbIdIdPutRequest struct {
	ctx               context.Context
	ApiService        *SharingSmbApiService
	id                int32
	sharingSmbUpdate1 *SharingSmbUpdate1
}

func (r ApiSharingSmbIdIdPutRequest) SharingSmbUpdate1(sharingSmbUpdate1 SharingSmbUpdate1) ApiSharingSmbIdIdPutRequest {
	r.sharingSmbUpdate1 = &sharingSmbUpdate1
	return r
}

func (r ApiSharingSmbIdIdPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.SharingSmbIdIdPutExecute(r)
}

/*
SharingSmbIdIdPut Method for SharingSmbIdIdPut

Update SMB Share of `id`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiSharingSmbIdIdPutRequest
*/
func (a *SharingSmbApiService) SharingSmbIdIdPut(ctx context.Context, id int32) ApiSharingSmbIdIdPutRequest {
	return ApiSharingSmbIdIdPutRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *SharingSmbApiService) SharingSmbIdIdPutExecute(r ApiSharingSmbIdIdPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharingSmbApiService.SharingSmbIdIdPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sharing/smb/id/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.sharingSmbUpdate1
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiSharingSmbPostRequest struct {
	ctx               context.Context
	ApiService        *SharingSmbApiService
	sharingSmbCreate0 *SharingSmbCreate0
}

func (r ApiSharingSmbPostRequest) SharingSmbCreate0(sharingSmbCreate0 SharingSmbCreate0) ApiSharingSmbPostRequest {
	r.sharingSmbCreate0 = &sharingSmbCreate0
	return r
}

func (r ApiSharingSmbPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.SharingSmbPostExecute(r)
}

/*
SharingSmbPost Method for SharingSmbPost

Create a SMB Share.

`purpose` applies common configuration presets depending on intended purpose.

`timemachine` when set, enables Time Machine backups for this share.

`ro` when enabled, prohibits write access to the share.

`guestok` when enabled, allows access to this share without a password.

`hostsallow` is a list of hostnames / IP addresses which have access to this share.

`hostsdeny` is a list of hostnames / IP addresses which are not allowed access to this share. If a handful
of hostnames are to be only allowed access, `hostsdeny` can be passed "ALL" which means that it will deny
access to ALL hostnames except for the ones which have been listed in `hostsallow`.

`acl` enables support for storing the SMB Security Descriptor as a Filesystem ACL.

`streams` enables support for storing alternate datastreams as filesystem extended attributes.

`fsrvp` enables support for the filesystem remote VSS protocol. This allows clients to create
ZFS snapshots through RPC.

`shadowcopy` enables support for the volume shadow copy service.

`auxsmbconf` is a string of additional smb4.conf parameters not covered by the system's API.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSharingSmbPostRequest
*/
func (a *SharingSmbApiService) SharingSmbPost(ctx context.Context) ApiSharingSmbPostRequest {
	return ApiSharingSmbPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *SharingSmbApiService) SharingSmbPostExecute(r ApiSharingSmbPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharingSmbApiService.SharingSmbPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sharing/smb"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.sharingSmbCreate0
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiSharingSmbPresetsGetRequest struct {
	ctx        context.Context
	ApiService *SharingSmbApiService
}

func (r ApiSharingSmbPresetsGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.SharingSmbPresetsGetExecute(r)
}

/*
SharingSmbPresetsGet Method for SharingSmbPresetsGet

Retrieve pre-defined configuration sets for specific use-cases. These parameter
combinations are often non-obvious, but beneficial in these scenarios.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSharingSmbPresetsGetRequest
*/
func (a *SharingSmbApiService) SharingSmbPresetsGet(ctx context.Context) ApiSharingSmbPresetsGetRequest {
	return ApiSharingSmbPresetsGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *SharingSmbApiService) SharingSmbPresetsGetExecute(r ApiSharingSmbPresetsGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharingSmbApiService.SharingSmbPresetsGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sharing/smb/presets"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
