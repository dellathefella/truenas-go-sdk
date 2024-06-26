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

// ActivedirectoryApiService ActivedirectoryApi service
type ActivedirectoryApiService service

type ApiActivedirectoryChangeTrustAccountPwGetRequest struct {
	ctx        context.Context
	ApiService *ActivedirectoryApiService
}

func (r ApiActivedirectoryChangeTrustAccountPwGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ActivedirectoryChangeTrustAccountPwGetExecute(r)
}

/*
ActivedirectoryChangeTrustAccountPwGet Method for ActivedirectoryChangeTrustAccountPwGet

Force an update of the AD machine account password. This can be used to
refresh the Kerberos principals in the server's system keytab.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiActivedirectoryChangeTrustAccountPwGetRequest
*/
func (a *ActivedirectoryApiService) ActivedirectoryChangeTrustAccountPwGet(ctx context.Context) ApiActivedirectoryChangeTrustAccountPwGetRequest {
	return ApiActivedirectoryChangeTrustAccountPwGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ActivedirectoryApiService) ActivedirectoryChangeTrustAccountPwGetExecute(r ApiActivedirectoryChangeTrustAccountPwGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActivedirectoryApiService.ActivedirectoryChangeTrustAccountPwGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/activedirectory/change_trust_account_pw"

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

type ApiActivedirectoryDomainInfoPostRequest struct {
	ctx        context.Context
	ApiService *ActivedirectoryApiService
	body       *string
}

func (r ApiActivedirectoryDomainInfoPostRequest) Body(body string) ApiActivedirectoryDomainInfoPostRequest {
	r.body = &body
	return r
}

func (r ApiActivedirectoryDomainInfoPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ActivedirectoryDomainInfoPostExecute(r)
}

/*
ActivedirectoryDomainInfoPost Method for ActivedirectoryDomainInfoPost

Returns the following information about the currently joined domain:

`LDAP server` IP address of current LDAP server to which TrueNAS is connected.

`LDAP server name` DNS name of LDAP server to which TrueNAS is connected

`Realm` Kerberos realm

`LDAP port`

`Server time` timestamp.

`KDC server` Kerberos KDC to which TrueNAS is connected

`Server time offset` current time offset from DC.

`Last machine account password change`. timestamp

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiActivedirectoryDomainInfoPostRequest
*/
func (a *ActivedirectoryApiService) ActivedirectoryDomainInfoPost(ctx context.Context) ApiActivedirectoryDomainInfoPostRequest {
	return ApiActivedirectoryDomainInfoPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ActivedirectoryApiService) ActivedirectoryDomainInfoPostExecute(r ApiActivedirectoryDomainInfoPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActivedirectoryApiService.ActivedirectoryDomainInfoPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/activedirectory/domain_info"

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
	localVarPostBody = r.body
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

type ApiActivedirectoryGetRequest struct {
	ctx        context.Context
	ApiService *ActivedirectoryApiService
}

func (r ApiActivedirectoryGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ActivedirectoryGetExecute(r)
}

/*
ActivedirectoryGet Method for ActivedirectoryGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiActivedirectoryGetRequest
*/
func (a *ActivedirectoryApiService) ActivedirectoryGet(ctx context.Context) ApiActivedirectoryGetRequest {
	return ApiActivedirectoryGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ActivedirectoryApiService) ActivedirectoryGetExecute(r ApiActivedirectoryGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActivedirectoryApiService.ActivedirectoryGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/activedirectory"

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

type ApiActivedirectoryGetSpnListGetRequest struct {
	ctx        context.Context
	ApiService *ActivedirectoryApiService
}

func (r ApiActivedirectoryGetSpnListGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ActivedirectoryGetSpnListGetExecute(r)
}

/*
ActivedirectoryGetSpnListGet Method for ActivedirectoryGetSpnListGet

Return list of kerberos SPN entries registered for the server's Active
Directory computer account. This may not reflect the state of the
server's current kerberos keytab.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiActivedirectoryGetSpnListGetRequest
*/
func (a *ActivedirectoryApiService) ActivedirectoryGetSpnListGet(ctx context.Context) ApiActivedirectoryGetSpnListGetRequest {
	return ApiActivedirectoryGetSpnListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ActivedirectoryApiService) ActivedirectoryGetSpnListGetExecute(r ApiActivedirectoryGetSpnListGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActivedirectoryApiService.ActivedirectoryGetSpnListGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/activedirectory/get_spn_list"

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

type ApiActivedirectoryGetStateGetRequest struct {
	ctx        context.Context
	ApiService *ActivedirectoryApiService
}

func (r ApiActivedirectoryGetStateGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ActivedirectoryGetStateGetExecute(r)
}

/*
ActivedirectoryGetStateGet Method for ActivedirectoryGetStateGet

Wrapper function for 'directoryservices.get_state'. Returns only the state of the
Active Directory service.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiActivedirectoryGetStateGetRequest
*/
func (a *ActivedirectoryApiService) ActivedirectoryGetStateGet(ctx context.Context) ApiActivedirectoryGetStateGetRequest {
	return ApiActivedirectoryGetStateGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ActivedirectoryApiService) ActivedirectoryGetStateGetExecute(r ApiActivedirectoryGetStateGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActivedirectoryApiService.ActivedirectoryGetStateGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/activedirectory/get_state"

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

type ApiActivedirectoryLeavePostRequest struct {
	ctx                   context.Context
	ApiService            *ActivedirectoryApiService
	activedirectoryLeave0 *ActivedirectoryLeave0
}

func (r ApiActivedirectoryLeavePostRequest) ActivedirectoryLeave0(activedirectoryLeave0 ActivedirectoryLeave0) ApiActivedirectoryLeavePostRequest {
	r.activedirectoryLeave0 = &activedirectoryLeave0
	return r
}

func (r ApiActivedirectoryLeavePostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ActivedirectoryLeavePostExecute(r)
}

/*
ActivedirectoryLeavePost Method for ActivedirectoryLeavePost

Leave Active Directory domain. This will remove computer
object from AD and clear relevant configuration data from
the NAS.
This requires credentials for appropriately-privileged user.
Credentials are used to obtain a kerberos ticket, which is
used to perform the actual removal from the domain.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiActivedirectoryLeavePostRequest
*/
func (a *ActivedirectoryApiService) ActivedirectoryLeavePost(ctx context.Context) ApiActivedirectoryLeavePostRequest {
	return ApiActivedirectoryLeavePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ActivedirectoryApiService) ActivedirectoryLeavePostExecute(r ApiActivedirectoryLeavePostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActivedirectoryApiService.ActivedirectoryLeavePost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/activedirectory/leave"

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
	localVarPostBody = r.activedirectoryLeave0
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

type ApiActivedirectoryNssInfoChoicesGetRequest struct {
	ctx        context.Context
	ApiService *ActivedirectoryApiService
}

func (r ApiActivedirectoryNssInfoChoicesGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ActivedirectoryNssInfoChoicesGetExecute(r)
}

/*
ActivedirectoryNssInfoChoicesGet Method for ActivedirectoryNssInfoChoicesGet

Returns list of available LDAP schema choices.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiActivedirectoryNssInfoChoicesGetRequest
*/
func (a *ActivedirectoryApiService) ActivedirectoryNssInfoChoicesGet(ctx context.Context) ApiActivedirectoryNssInfoChoicesGetRequest {
	return ApiActivedirectoryNssInfoChoicesGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ActivedirectoryApiService) ActivedirectoryNssInfoChoicesGetExecute(r ApiActivedirectoryNssInfoChoicesGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActivedirectoryApiService.ActivedirectoryNssInfoChoicesGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/activedirectory/nss_info_choices"

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

type ApiActivedirectoryPutRequest struct {
	ctx                    context.Context
	ApiService             *ActivedirectoryApiService
	activedirectoryUpdate0 *ActivedirectoryUpdate0
}

func (r ApiActivedirectoryPutRequest) ActivedirectoryUpdate0(activedirectoryUpdate0 ActivedirectoryUpdate0) ApiActivedirectoryPutRequest {
	r.activedirectoryUpdate0 = &activedirectoryUpdate0
	return r
}

func (r ApiActivedirectoryPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.ActivedirectoryPutExecute(r)
}

/*
ActivedirectoryPut Method for ActivedirectoryPut

Update active directory configuration.
`domainname` full DNS domain name of the Active Directory domain.

`bindname` username used to perform the intial domain join.

`bindpw` password used to perform the initial domain join. User-
provided credentials are used to obtain a kerberos ticket, which
is used to perform the actual domain join.

`verbose_logging` increase logging during the domain join process.

`use_default_domain` controls whether domain users and groups have
the pre-windows 2000 domain name prepended to the user account. When
enabled, the user appears as "administrator" rather than
"EXAMPLEdministrator"

`allow_trusted_doms` enable support for trusted domains. If this
parameter is enabled, then separate idmap backends _must_ be configured
for each trusted domain, and the idmap cache should be cleared.

`allow_dns_updates` during the domain join process, automatically
generate DNS entries in the AD domain for the NAS. If this is disabled,
then a domain administrator must manually add appropriate DNS entries
for the NAS. This parameter is recommended for TrueNAS HA servers.

`disable_freenas_cache` disables active caching of AD users and groups.
When disabled, only users cached in winbind's internal cache are
visible in GUI dropdowns. Disabling active caching is recommended
in environments with a large amount of users.

`site` AD site of which the NAS is a member. This parameter is auto-
detected during the domain join process. If no AD site is configured
for the subnet in which the NAS is configured, then this parameter
appears as 'Default-First-Site-Name'. Auto-detection is only performed
during the initial domain join.

`kerberos_realm` in which the server is located. This parameter is
automatically populated during the initial domain join. If the NAS has
an AD site configured and that site has multiple kerberos servers, then
the kerberos realm is automatically updated with a site-specific
configuration to use those servers. Auto-detection is only performed
during initial domain join.

`kerberos_principal` kerberos principal to use for AD-related
operations outside of Samba. After intial domain join, this field is
updated with the kerberos principal associated with the AD machine
account for the NAS.

`nss_info` controls how Winbind retrieves Name Service Information to
construct a user's home directory and login shell. This parameter
is only effective if the Active Directory Domain Controller supports
the Microsoft Services for Unix (SFU) LDAP schema.

`timeout` timeout value for winbind-related operations. This value may
need to be increased in  environments with high latencies for
communications with domain controllers or a large number of domain
controllers. Lowering the value may cause status checks to fail.

`dns_timeout` timeout value for DNS queries during the initial domain
join. This value is also set as the NETWORK_TIMEOUT in the ldap config
file.

`createcomputer` Active Directory Organizational Unit in which new
computer accounts are created.

The OU string is read from top to bottom without RDNs. Slashes ("/")
are used as delimiters, like `Computers/Servers/NAS`. The backslash
("\") is used to escape characters but not as a separator. Backslashes
are interpreted at multiple levels and might require doubling or even
quadrupling to take effect.

When this field is blank, new computer accounts are created in the
Active Directory default OU.

The Active Directory service is started after a configuration
update if the service was initially disabled, and the updated
configuration sets `enable` to `True`. The Active Directory
service is stopped if `enable` is changed to `False`. If the
configuration is updated, but the initial `enable` state is `True`, and
remains unchanged, then the samba server is only restarted.

During the domain join, a kerberos keytab for the newly-created AD
machine account is generated. It is used for all future
LDAP / AD interaction and the user-provided credentials are removed.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiActivedirectoryPutRequest
*/
func (a *ActivedirectoryApiService) ActivedirectoryPut(ctx context.Context) ApiActivedirectoryPutRequest {
	return ApiActivedirectoryPutRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ActivedirectoryApiService) ActivedirectoryPutExecute(r ApiActivedirectoryPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActivedirectoryApiService.ActivedirectoryPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/activedirectory"

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
	localVarPostBody = r.activedirectoryUpdate0
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

type ApiActivedirectoryStartedGetRequest struct {
	ctx        context.Context
	ApiService *ActivedirectoryApiService
}

func (r ApiActivedirectoryStartedGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ActivedirectoryStartedGetExecute(r)
}

/*
ActivedirectoryStartedGet Method for ActivedirectoryStartedGet

Issue a no-effect command to our DC. This checks if our secure channel connection to our
domain controller is still alive. It has much less impact than wbinfo -t.
Default winbind request timeout is 60 seconds, and can be adjusted by the smb4.conf parameter
'winbind request timeout ='

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiActivedirectoryStartedGetRequest
*/
func (a *ActivedirectoryApiService) ActivedirectoryStartedGet(ctx context.Context) ApiActivedirectoryStartedGetRequest {
	return ApiActivedirectoryStartedGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ActivedirectoryApiService) ActivedirectoryStartedGetExecute(r ApiActivedirectoryStartedGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActivedirectoryApiService.ActivedirectoryStartedGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/activedirectory/started"

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
