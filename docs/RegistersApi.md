# \RegistersApi

All URIs are relative to *http://101.42.88.184/cns-backend/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRegister**](RegistersApi.md#GetRegister) | **Get** /registers/{commit_hash} | get register info by admin
[**GetRegisterOrder**](RegistersApi.md#GetRegisterOrder) | **Get** /registers/order/{commit_hash} | get register order
[**MakeRegisterOrder**](RegistersApi.md#MakeRegisterOrder) | **Post** /registers/order/{commit_hash} | make register order
[**RefreshRegisterOrderUrl**](RegistersApi.md#RefreshRegisterOrderUrl) | **Put** /registers/order/refresh-url/{commit_hash} | refresh register order url
[**Register**](RegistersApi.md#Register) | **Post** /registers | register by admin



## GetRegister

> ServicesRegisterByAdminResp GetRegister(ctx, commitHash).Execute()

get register info by admin



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
    commitHash := "commitHash_example" // string | commit hash

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistersApi.GetRegister(context.Background(), commitHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistersApi.GetRegister``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegister`: ServicesRegisterByAdminResp
    fmt.Fprintf(os.Stdout, "Response from `RegistersApi.GetRegister`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commitHash** | **string** | commit hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegisterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServicesRegisterByAdminResp**](ServicesRegisterByAdminResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegisterOrder

> ModelsRegisterCore GetRegisterOrder(ctx, commitHash).Execute()

get register order



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
    commitHash := "commitHash_example" // string | commit hash

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistersApi.GetRegisterOrder(context.Background(), commitHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistersApi.GetRegisterOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegisterOrder`: ModelsRegisterCore
    fmt.Fprintf(os.Stdout, "Response from `RegistersApi.GetRegisterOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commitHash** | **string** | commit hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegisterOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsRegisterCore**](ModelsRegisterCore.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MakeRegisterOrder

> ServicesMakeRegisterOrderResp MakeRegisterOrder(ctx, commitHash).MakeRegisterOrderRequest(makeRegisterOrderRequest).Execute()

make register order



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
    commitHash := "commitHash_example" // string | commit hash
    makeRegisterOrderRequest := *openapiclient.NewServicesMakeRegisterOrderReq("Description_example", "TradeType_example") // ServicesMakeRegisterOrderReq | make register order request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistersApi.MakeRegisterOrder(context.Background(), commitHash).MakeRegisterOrderRequest(makeRegisterOrderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistersApi.MakeRegisterOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MakeRegisterOrder`: ServicesMakeRegisterOrderResp
    fmt.Fprintf(os.Stdout, "Response from `RegistersApi.MakeRegisterOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commitHash** | **string** | commit hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiMakeRegisterOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **makeRegisterOrderRequest** | [**ServicesMakeRegisterOrderReq**](ServicesMakeRegisterOrderReq.md) | make register order request | 

### Return type

[**ServicesMakeRegisterOrderResp**](ServicesMakeRegisterOrderResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshRegisterOrderUrl

> ServicesMakeRegisterOrderResp RefreshRegisterOrderUrl(ctx, commitHash).Execute()

refresh register order url



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
    commitHash := "commitHash_example" // string | commit hash

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistersApi.RefreshRegisterOrderUrl(context.Background(), commitHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistersApi.RefreshRegisterOrderUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshRegisterOrderUrl`: ServicesMakeRegisterOrderResp
    fmt.Fprintf(os.Stdout, "Response from `RegistersApi.RefreshRegisterOrderUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commitHash** | **string** | commit hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshRegisterOrderUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServicesMakeRegisterOrderResp**](ServicesMakeRegisterOrderResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Register

> ServicesRegisterByAdminResp Register(ctx).XApiKey(xApiKey).Commit(commit).Execute()

register by admin



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
    xApiKey := "xApiKey_example" // string | api key
    commit := *openapiclient.NewModelsCommitCore("CommitHash_example", []string{"Data_example"}, int32(123), "Name_example", "Owner_example", "Resolver_example", false, "Secret_example") // ModelsCommitCore | regitser request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistersApi.Register(context.Background()).XApiKey(xApiKey).Commit(commit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistersApi.Register``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Register`: ServicesRegisterByAdminResp
    fmt.Fprintf(os.Stdout, "Response from `RegistersApi.Register`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** | api key | 
 **commit** | [**ModelsCommitCore**](ModelsCommitCore.md) | regitser request | 

### Return type

[**ServicesRegisterByAdminResp**](ServicesRegisterByAdminResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

