# \AuthTwofactorApi

All URIs are relative to *http://10.234.27.190/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthTwofactorGet**](AuthTwofactorApi.md#AuthTwofactorGet) | **Get** /auth/twofactor | 
[**AuthTwofactorProvisioningUriGet**](AuthTwofactorApi.md#AuthTwofactorProvisioningUriGet) | **Get** /auth/twofactor/provisioning_uri | 
[**AuthTwofactorPut**](AuthTwofactorApi.md#AuthTwofactorPut) | **Put** /auth/twofactor | 
[**AuthTwofactorRenewSecretGet**](AuthTwofactorApi.md#AuthTwofactorRenewSecretGet) | **Get** /auth/twofactor/renew_secret | 
[**AuthTwofactorVerifyPost**](AuthTwofactorApi.md#AuthTwofactorVerifyPost) | **Post** /auth/twofactor/verify | 



## AuthTwofactorGet

> AuthTwofactorGet(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthTwofactorApi.AuthTwofactorGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthTwofactorApi.AuthTwofactorGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTwofactorGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTwofactorProvisioningUriGet

> AuthTwofactorProvisioningUriGet(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthTwofactorApi.AuthTwofactorProvisioningUriGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthTwofactorApi.AuthTwofactorProvisioningUriGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTwofactorProvisioningUriGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTwofactorPut

> AuthTwofactorPut(ctx).AuthTwofactorUpdate0(authTwofactorUpdate0).Execute()





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
    authTwofactorUpdate0 := *openapiclient.NewAuthTwofactorUpdate0() // AuthTwofactorUpdate0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthTwofactorApi.AuthTwofactorPut(context.Background()).AuthTwofactorUpdate0(authTwofactorUpdate0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthTwofactorApi.AuthTwofactorPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthTwofactorPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authTwofactorUpdate0** | [**AuthTwofactorUpdate0**](AuthTwofactorUpdate0.md) |  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTwofactorRenewSecretGet

> AuthTwofactorRenewSecretGet(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthTwofactorApi.AuthTwofactorRenewSecretGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthTwofactorApi.AuthTwofactorRenewSecretGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTwofactorRenewSecretGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTwofactorVerifyPost

> AuthTwofactorVerifyPost(ctx).Body(body).Execute()





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
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthTwofactorApi.AuthTwofactorVerifyPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthTwofactorApi.AuthTwofactorVerifyPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthTwofactorVerifyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

