# \CommitsApi

All URIs are relative to *http://101.42.88.184/cns-backend/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCommit**](CommitsApi.md#GetCommit) | **Get** /commits/{commit_hash} | get commit
[**MakeCommits**](CommitsApi.md#MakeCommits) | **Post** /commits | make commit
[**QueryCommit**](CommitsApi.md#QueryCommit) | **Get** /commits | query commit



## GetCommit

> ModelsCommitCore GetCommit(ctx, commitHash).Execute()

get commit



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
    resp, r, err := apiClient.CommitsApi.GetCommit(context.Background(), commitHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.GetCommit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommit`: ModelsCommitCore
    fmt.Fprintf(os.Stdout, "Response from `CommitsApi.GetCommit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commitHash** | **string** | commit hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsCommitCore**](ModelsCommitCore.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MakeCommits

> ServicesMakeCommitResp MakeCommits(ctx).MakeCommitReq(makeCommitReq).Execute()

make commit



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
    makeCommitReq := *openapiclient.NewModelsCommitCore("CommitHash_example", []string{"Data_example"}, int32(123), "Name_example", "Owner_example", "Resolver_example", false, "Secret_example") // ModelsCommitCore | make commit request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CommitsApi.MakeCommits(context.Background()).MakeCommitReq(makeCommitReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.MakeCommits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MakeCommits`: ServicesMakeCommitResp
    fmt.Fprintf(os.Stdout, "Response from `CommitsApi.MakeCommits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMakeCommitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **makeCommitReq** | [**ModelsCommitCore**](ModelsCommitCore.md) | make commit request | 

### Return type

[**ServicesMakeCommitResp**](ServicesMakeCommitResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryCommit

> []ModelsCommitCore QueryCommit(ctx).OrderState(orderState).Owner(owner).Page(page).PageSize(pageSize).Execute()

query commit



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
    orderState := "orderState_example" // string |  (optional)
    owner := "owner_example" // string |  (optional)
    page := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CommitsApi.QueryCommit(context.Background()).OrderState(orderState).Owner(owner).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommitsApi.QueryCommit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryCommit`: []ModelsCommitCore
    fmt.Fprintf(os.Stdout, "Response from `CommitsApi.QueryCommit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderState** | **string** |  | 
 **owner** | **string** |  | 
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 

### Return type

[**[]ModelsCommitCore**](ModelsCommitCore.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

