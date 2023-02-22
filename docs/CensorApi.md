# \CensorApi

All URIs are relative to *http://101.42.88.184/cns-backend/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CensorText**](CensorApi.md#CensorText) | **Post** /censor/text | censor text if compliance



## CensorText

> ServicesCensorTextResp CensorText(ctx).CensorTextReq(censorTextReq).Execute()

censor text if compliance



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
    censorTextReq := *openapiclient.NewServicesCensorTextReq("Content_example") // ServicesCensorTextReq | text for censor

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CensorApi.CensorText(context.Background()).CensorTextReq(censorTextReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CensorApi.CensorText``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CensorText`: ServicesCensorTextResp
    fmt.Fprintf(os.Stdout, "Response from `CensorApi.CensorText`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCensorTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **censorTextReq** | [**ServicesCensorTextReq**](ServicesCensorTextReq.md) | text for censor | 

### Return type

[**ServicesCensorTextResp**](ServicesCensorTextResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

