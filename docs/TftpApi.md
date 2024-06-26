# \TftpApi

All URIs are relative to *http://10.234.27.190/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TftpGet**](TftpApi.md#TftpGet) | **Get** /tftp | 
[**TftpPut**](TftpApi.md#TftpPut) | **Put** /tftp | 



## TftpGet

> TftpGet(ctx).Execute()



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
    resp, r, err := apiClient.TftpApi.TftpGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TftpApi.TftpGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTftpGetRequest struct via the builder pattern


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


## TftpPut

> TftpPut(ctx).TftpUpdate0(tftpUpdate0).Execute()





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
    tftpUpdate0 := *openapiclient.NewTftpUpdate0() // TftpUpdate0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TftpApi.TftpPut(context.Background()).TftpUpdate0(tftpUpdate0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TftpApi.TftpPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTftpPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tftpUpdate0** | [**TftpUpdate0**](TftpUpdate0.md) |  | 

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

