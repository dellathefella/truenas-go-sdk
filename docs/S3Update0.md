# S3Update0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** |  | [optional] 
**Bindip** | Pointer to **string** |  | [optional] 
**Bindport** | Pointer to **int32** |  | [optional] 
**Browser** | Pointer to **bool** |  | [optional] 
**Certificate** | Pointer to **NullableInt32** |  | [optional] 
**ConsoleBindport** | Pointer to **int32** |  | [optional] 
**SecretKey** | Pointer to **string** |  | [optional] 
**StoragePath** | Pointer to **string** |  | [optional] 
**TlsServerUri** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewS3Update0

`func NewS3Update0() *S3Update0`

NewS3Update0 instantiates a new S3Update0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3Update0WithDefaults

`func NewS3Update0WithDefaults() *S3Update0`

NewS3Update0WithDefaults instantiates a new S3Update0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *S3Update0) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *S3Update0) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *S3Update0) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *S3Update0) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetBindip

`func (o *S3Update0) GetBindip() string`

GetBindip returns the Bindip field if non-nil, zero value otherwise.

### GetBindipOk

`func (o *S3Update0) GetBindipOk() (*string, bool)`

GetBindipOk returns a tuple with the Bindip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindip

`func (o *S3Update0) SetBindip(v string)`

SetBindip sets Bindip field to given value.

### HasBindip

`func (o *S3Update0) HasBindip() bool`

HasBindip returns a boolean if a field has been set.

### GetBindport

`func (o *S3Update0) GetBindport() int32`

GetBindport returns the Bindport field if non-nil, zero value otherwise.

### GetBindportOk

`func (o *S3Update0) GetBindportOk() (*int32, bool)`

GetBindportOk returns a tuple with the Bindport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindport

`func (o *S3Update0) SetBindport(v int32)`

SetBindport sets Bindport field to given value.

### HasBindport

`func (o *S3Update0) HasBindport() bool`

HasBindport returns a boolean if a field has been set.

### GetBrowser

`func (o *S3Update0) GetBrowser() bool`

GetBrowser returns the Browser field if non-nil, zero value otherwise.

### GetBrowserOk

`func (o *S3Update0) GetBrowserOk() (*bool, bool)`

GetBrowserOk returns a tuple with the Browser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowser

`func (o *S3Update0) SetBrowser(v bool)`

SetBrowser sets Browser field to given value.

### HasBrowser

`func (o *S3Update0) HasBrowser() bool`

HasBrowser returns a boolean if a field has been set.

### GetCertificate

`func (o *S3Update0) GetCertificate() int32`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *S3Update0) GetCertificateOk() (*int32, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *S3Update0) SetCertificate(v int32)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *S3Update0) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *S3Update0) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *S3Update0) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetConsoleBindport

`func (o *S3Update0) GetConsoleBindport() int32`

GetConsoleBindport returns the ConsoleBindport field if non-nil, zero value otherwise.

### GetConsoleBindportOk

`func (o *S3Update0) GetConsoleBindportOk() (*int32, bool)`

GetConsoleBindportOk returns a tuple with the ConsoleBindport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleBindport

`func (o *S3Update0) SetConsoleBindport(v int32)`

SetConsoleBindport sets ConsoleBindport field to given value.

### HasConsoleBindport

`func (o *S3Update0) HasConsoleBindport() bool`

HasConsoleBindport returns a boolean if a field has been set.

### GetSecretKey

`func (o *S3Update0) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *S3Update0) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *S3Update0) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *S3Update0) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetStoragePath

`func (o *S3Update0) GetStoragePath() string`

GetStoragePath returns the StoragePath field if non-nil, zero value otherwise.

### GetStoragePathOk

`func (o *S3Update0) GetStoragePathOk() (*string, bool)`

GetStoragePathOk returns a tuple with the StoragePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePath

`func (o *S3Update0) SetStoragePath(v string)`

SetStoragePath sets StoragePath field to given value.

### HasStoragePath

`func (o *S3Update0) HasStoragePath() bool`

HasStoragePath returns a boolean if a field has been set.

### GetTlsServerUri

`func (o *S3Update0) GetTlsServerUri() string`

GetTlsServerUri returns the TlsServerUri field if non-nil, zero value otherwise.

### GetTlsServerUriOk

`func (o *S3Update0) GetTlsServerUriOk() (*string, bool)`

GetTlsServerUriOk returns a tuple with the TlsServerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsServerUri

`func (o *S3Update0) SetTlsServerUri(v string)`

SetTlsServerUri sets TlsServerUri field to given value.

### HasTlsServerUri

`func (o *S3Update0) HasTlsServerUri() bool`

HasTlsServerUri returns a boolean if a field has been set.

### SetTlsServerUriNil

`func (o *S3Update0) SetTlsServerUriNil(b bool)`

 SetTlsServerUriNil sets the value for TlsServerUri to be an explicit nil

### UnsetTlsServerUri
`func (o *S3Update0) UnsetTlsServerUri()`

UnsetTlsServerUri ensures that no value is present for TlsServerUri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


