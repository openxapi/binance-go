# \ExchangeLinkAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExchangelinkGetBrokerSubAccountFuturesSummaryV3**](ExchangeLinkAPI.md#ExchangelinkGetBrokerSubAccountFuturesSummaryV3) | **Get** /sapi/v3/broker/subAccount/futuresSummary | Query Sub Account Futures Asset info (V3)



## ExchangelinkGetBrokerSubAccountFuturesSummaryV3

> ExchangelinkGetBrokerSubAccountFuturesSummaryV3Resp ExchangelinkGetBrokerSubAccountFuturesSummaryV3(ctx).FuturesType(futuresType).Timestamp(timestamp).SubAccountId(subAccountId).Page(page).Size(size).RecvWindow(recvWindow).Execute()

Query Sub Account Futures Asset info (V3)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	futuresType := int32(56) // int32 | 1:USD Margined Futures, 2:COIN Margined Futures
	timestamp := int64(789) // int64 | 
	subAccountId := "subAccountId_example" // string |  (optional) (default to "")
	page := int64(789) // int64 | default 1 (optional) (default to 1)
	size := int64(789) // int64 | default 10, max 20 (optional) (default to 10)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExchangeLinkAPI.ExchangelinkGetBrokerSubAccountFuturesSummaryV3(context.Background()).FuturesType(futuresType).Timestamp(timestamp).SubAccountId(subAccountId).Page(page).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExchangeLinkAPI.ExchangelinkGetBrokerSubAccountFuturesSummaryV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExchangelinkGetBrokerSubAccountFuturesSummaryV3`: ExchangelinkGetBrokerSubAccountFuturesSummaryV3Resp
	fmt.Fprintf(os.Stdout, "Response from `ExchangeLinkAPI.ExchangelinkGetBrokerSubAccountFuturesSummaryV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExchangelinkGetBrokerSubAccountFuturesSummaryV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **futuresType** | **int32** | 1:USD Margined Futures, 2:COIN Margined Futures | 
 **timestamp** | **int64** |  | 
 **subAccountId** | **string** |  | [default to &quot;&quot;]
 **page** | **int64** | default 1 | [default to 1]
 **size** | **int64** | default 10, max 20 | [default to 10]
 **recvWindow** | **int64** |  | 

### Return type

[**ExchangelinkGetBrokerSubAccountFuturesSummaryV3Resp**](ExchangelinkGetBrokerSubAccountFuturesSummaryV3Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

