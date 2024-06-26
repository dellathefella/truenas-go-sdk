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
)

// WebdavApiService WebdavApi service
type WebdavApiService service

type ApiWebdavCertChoicesGetRequest struct {
	ctx        context.Context
	ApiService *WebdavApiService
}

func (r ApiWebdavCertChoicesGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebdavCertChoicesGetExecute(r)
}

/*
WebdavCertChoicesGet Method for WebdavCertChoicesGet

Permitted certificate choices for webdav service.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebdavCertChoicesGetRequest
*/
func (a *WebdavApiService) WebdavCertChoicesGet(ctx context.Context) ApiWebdavCertChoicesGetRequest {
	return ApiWebdavCertChoicesGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *WebdavApiService) WebdavCertChoicesGetExecute(r ApiWebdavCertChoicesGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebdavApiService.WebdavCertChoicesGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webdav/cert_choices"

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

type ApiWebdavGetRequest struct {
	ctx        context.Context
	ApiService *WebdavApiService
}

func (r ApiWebdavGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebdavGetExecute(r)
}

/*
WebdavGet Method for WebdavGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebdavGetRequest
*/
func (a *WebdavApiService) WebdavGet(ctx context.Context) ApiWebdavGetRequest {
	return ApiWebdavGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *WebdavApiService) WebdavGetExecute(r ApiWebdavGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebdavApiService.WebdavGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webdav"

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

type ApiWebdavPutRequest struct {
	ctx           context.Context
	ApiService    *WebdavApiService
	webdavUpdate0 *WebdavUpdate0
}

func (r ApiWebdavPutRequest) WebdavUpdate0(webdavUpdate0 WebdavUpdate0) ApiWebdavPutRequest {
	r.webdavUpdate0 = &webdavUpdate0
	return r
}

func (r ApiWebdavPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebdavPutExecute(r)
}

/*
WebdavPut Method for WebdavPut

Update Webdav Service Configuration.

`protocol` specifies which protocol should be used for connecting to Webdav Serivce. Value of "HTTPHTTPS"
allows both HTTP and HTTPS connections to the share.

`certssl` is a valid id of a certificate configured in the system. This is required if HTTPS connection is
desired with Webdave Service.

There are 3 types of Authentication supported with Webdav:
1) NONE      -   No authentication is required
2) BASIC     -   Password is sent over the network as plaintext
3) DIGEST    -   Hash of the password is sent over the network

`htauth` should be one of the valid types described above.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebdavPutRequest
*/
func (a *WebdavApiService) WebdavPut(ctx context.Context) ApiWebdavPutRequest {
	return ApiWebdavPutRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *WebdavApiService) WebdavPutExecute(r ApiWebdavPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebdavApiService.WebdavPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webdav"

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
	localVarPostBody = r.webdavUpdate0
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
