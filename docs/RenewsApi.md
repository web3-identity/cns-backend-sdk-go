# \RenewsApi

All URIs are relative to *http://101.42.88.184/cns-backend/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRenew**](RenewsApi.md#GetRenew) | **Get** /renews/{id} | get renew order by admin
[**GetRenewOrder**](RenewsApi.md#GetRenewOrder) | **Get** /renews/order/{id} | get renew order
[**MakeRenewOrder**](RenewsApi.md#MakeRenewOrder) | **Post** /renews/order | make renew order
[**RefreshRenewOrderUrl**](RenewsApi.md#RefreshRenewOrderUrl) | **Put** /renews/order/refresh-url/{id} | refresh renew order url
[**Renew**](RenewsApi.md#Renew) | **Post** /renews | renew by admin



## GetRenew

> ServicesRnewByAdminResp GetRenew(ctx, id).Execute()

get renew order by admin



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
    id := float32(8.14) // float32 | id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RenewsApi.GetRenew(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RenewsApi.GetRenew``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRenew`: ServicesRnewByAdminResp
    fmt.Fprintf(os.Stdout, "Response from `RenewsApi.GetRenew`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRenewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServicesRnewByAdminResp**](ServicesRnewByAdminResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRenewOrder

> ModelsRenewCore GetRenewOrder(ctx, id).Execute()

get renew order



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
    id := float32(8.14) // float32 | id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RenewsApi.GetRenewOrder(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RenewsApi.GetRenewOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRenewOrder`: ModelsRenewCore
    fmt.Fprintf(os.Stdout, "Response from `RenewsApi.GetRenewOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRenewOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsRenewCore**](ModelsRenewCore.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MakeRenewOrder

> ServicesMakeRenewOrderResp MakeRenewOrder(ctx).MakeRenewOrderRequest(makeRenewOrderRequest).Execute()

make renew order



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
    makeRenewOrderRequest := *openapiclient.NewServicesMakeRenewOrderReq("CnsName_example", "Description_example", int32(123), "TradeType_example", int32(123)) // ServicesMakeRenewOrderReq | make renew order request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RenewsApi.MakeRenewOrder(context.Background()).MakeRenewOrderRequest(makeRenewOrderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RenewsApi.MakeRenewOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MakeRenewOrder`: ServicesMakeRenewOrderResp
    fmt.Fprintf(os.Stdout, "Response from `RenewsApi.MakeRenewOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMakeRenewOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **makeRenewOrderRequest** | [**ServicesMakeRenewOrderReq**](ServicesMakeRenewOrderReq.md) | make renew order request | 

### Return type

[**ServicesMakeRenewOrderResp**](ServicesMakeRenewOrderResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshRenewOrderUrl

> ServicesMakeRenewOrderResp RefreshRenewOrderUrl(ctx, id).Execute()

refresh renew order url



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
    id := float32(8.14) // float32 | id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RenewsApi.RefreshRenewOrderUrl(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RenewsApi.RefreshRenewOrderUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshRenewOrderUrl`: ServicesMakeRenewOrderResp
    fmt.Fprintf(os.Stdout, "Response from `RenewsApi.RefreshRenewOrderUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshRenewOrderUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServicesMakeRenewOrderResp**](ServicesMakeRenewOrderResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Renew

> ServicesRnewByAdminResp Renew(ctx).MakeRenewOrderRequest(makeRenewOrderRequest).Execute()

renew by admin



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
    makeRenewOrderRequest := *openapiclient.NewModelsRenewOrderArgs("CnsName_example", int32(123), int32(123)) // ModelsRenewOrderArgs | renew request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RenewsApi.Renew(context.Background()).MakeRenewOrderRequest(makeRenewOrderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RenewsApi.Renew``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Renew`: ServicesRnewByAdminResp
    fmt.Fprintf(os.Stdout, "Response from `RenewsApi.Renew`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRenewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **makeRenewOrderRequest** | [**ModelsRenewOrderArgs**](ModelsRenewOrderArgs.md) | renew request | 

### Return type

[**ServicesRnewByAdminResp**](ServicesRnewByAdminResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

