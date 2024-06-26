# \AuthApi

All URIs are relative to *http://10.234.27.190/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthCheckPasswordPost**](AuthApi.md#AuthCheckPasswordPost) | **Post** /auth/check_password | 
[**AuthCheckUserPost**](AuthApi.md#AuthCheckUserPost) | **Post** /auth/check_user | 
[**AuthGenerateTokenPost**](AuthApi.md#AuthGenerateTokenPost) | **Post** /auth/generate_token | 
[**AuthSessionsGet**](AuthApi.md#AuthSessionsGet) | **Get** /auth/sessions | 
[**AuthTwoFactorAuthGet**](AuthApi.md#AuthTwoFactorAuthGet) | **Get** /auth/two_factor_auth | 



## AuthCheckPasswordPost

> AuthCheckPasswordPost(ctx).AuthCheckPassword(authCheckPassword).Execute()





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
    authCheckPassword := *openapiclient.NewAuthCheckPassword() // AuthCheckPassword |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.AuthCheckPasswordPost(context.Background()).AuthCheckPassword(authCheckPassword).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthCheckPasswordPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthCheckPasswordPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authCheckPassword** | [**AuthCheckPassword**](AuthCheckPassword.md) |  | 

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


## AuthCheckUserPost

> AuthCheckUserPost(ctx).AuthCheckUser(authCheckUser).Execute()





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
    authCheckUser := *openapiclient.NewAuthCheckUser() // AuthCheckUser |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.AuthCheckUserPost(context.Background()).AuthCheckUser(authCheckUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthCheckUserPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthCheckUserPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authCheckUser** | [**AuthCheckUser**](AuthCheckUser.md) |  | 

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


## AuthGenerateTokenPost

> AuthGenerateTokenPost(ctx).AuthGenerateToken(authGenerateToken).Execute()





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
    authGenerateToken := *openapiclient.NewAuthGenerateToken() // AuthGenerateToken |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.AuthGenerateTokenPost(context.Background()).AuthGenerateToken(authGenerateToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthGenerateTokenPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthGenerateTokenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authGenerateToken** | [**AuthGenerateToken**](AuthGenerateToken.md) |  | 

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


## AuthSessionsGet

> AuthSessionsGet(ctx).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()





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
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    count := true // bool |  (optional)
    sort := "sort_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.AuthSessionsGet(context.Background()).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthSessionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthSessionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **count** | **bool** |  | 
 **sort** | **string** |  | 

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


## AuthTwoFactorAuthGet

> AuthTwoFactorAuthGet(ctx).Execute()





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
    resp, r, err := apiClient.AuthApi.AuthTwoFactorAuthGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthTwoFactorAuthGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTwoFactorAuthGetRequest struct via the builder pattern


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

