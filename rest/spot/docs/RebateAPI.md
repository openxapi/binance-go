# \RebateAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRebateTaxQueryV1**](RebateAPI.md#GetRebateTaxQueryV1) | **Get** /sapi/v1/rebate/taxQuery | Get Spot Rebate History Records (USER_DATA)



## GetRebateTaxQueryV1

> GetRebateTaxQueryV1Resp GetRebateTaxQueryV1(ctx).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Page(page).RecvWindow(recvWindow).Execute()

Get Spot Rebate History Records (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	page := int32(56) // int32 | Default 1 (optional) (default to 1)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RebateAPI.GetRebateTaxQueryV1(context.Background()).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Page(page).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RebateAPI.GetRebateTaxQueryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRebateTaxQueryV1`: GetRebateTaxQueryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `RebateAPI.GetRebateTaxQueryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRebateTaxQueryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **page** | **int32** | Default 1 | [default to 1]
 **recvWindow** | **int64** |  | 

### Return type

[**GetRebateTaxQueryV1Resp**](GetRebateTaxQueryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

