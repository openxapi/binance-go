# \BinancePayHistoryAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPayTransactionsV1**](BinancePayHistoryAPI.md#GetPayTransactionsV1) | **Get** /sapi/v1/pay/transactions | Get Pay Trade History



## GetPayTransactionsV1

> GetPayTransactionsV1Resp GetPayTransactionsV1(ctx).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Get Pay Trade History



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
	limit := int32(56) // int32 | default 100, max 100 (optional) (default to 100)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinancePayHistoryAPI.GetPayTransactionsV1(context.Background()).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinancePayHistoryAPI.GetPayTransactionsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPayTransactionsV1`: GetPayTransactionsV1Resp
	fmt.Fprintf(os.Stdout, "Response from `BinancePayHistoryAPI.GetPayTransactionsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPayTransactionsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | default 100, max 100 | [default to 100]
 **recvWindow** | **int64** |  | 

### Return type

[**GetPayTransactionsV1Resp**](GetPayTransactionsV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

