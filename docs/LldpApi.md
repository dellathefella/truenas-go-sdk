# \LldpApi

All URIs are relative to *http://10.234.27.190/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LldpCountryChoicesGet**](LldpApi.md#LldpCountryChoicesGet) | **Get** /lldp/country_choices | 
[**LldpGet**](LldpApi.md#LldpGet) | **Get** /lldp | 
[**LldpPut**](LldpApi.md#LldpPut) | **Put** /lldp | 



## LldpCountryChoicesGet

> LldpCountryChoicesGet(ctx).Execute()





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
    resp, r, err := apiClient.LldpApi.LldpCountryChoicesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LldpApi.LldpCountryChoicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLldpCountryChoicesGetRequest struct via the builder pattern


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


## LldpGet

> LldpGet(ctx).Execute()



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
    resp, r, err := apiClient.LldpApi.LldpGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LldpApi.LldpGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLldpGetRequest struct via the builder pattern


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


## LldpPut

> LldpPut(ctx).LldpUpdate0(lldpUpdate0).Execute()





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
    lldpUpdate0 := *openapiclient.NewLldpUpdate0() // LldpUpdate0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LldpApi.LldpPut(context.Background()).LldpUpdate0(lldpUpdate0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LldpApi.LldpPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLldpPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lldpUpdate0** | [**LldpUpdate0**](LldpUpdate0.md) |  | 

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

