# \V1API

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopytradingGetCopyTradingFuturesLeadSymbolV1**](V1API.md#CopytradingGetCopyTradingFuturesLeadSymbolV1) | **Get** /sapi/v1/copyTrading/futures/leadSymbol | Get Futures Lead Trading Symbol Whitelist(USER_DATA)
[**CopytradingGetCopyTradingFuturesUserStatusV1**](V1API.md#CopytradingGetCopyTradingFuturesUserStatusV1) | **Get** /sapi/v1/copyTrading/futures/userStatus | Get Futures Lead Trader Status(TRADE)



## CopytradingGetCopyTradingFuturesLeadSymbolV1

> CopytradingGetCopyTradingFuturesLeadSymbolV1Resp CopytradingGetCopyTradingFuturesLeadSymbolV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Futures Lead Trading Symbol Whitelist(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/copytrading"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CopytradingGetCopyTradingFuturesLeadSymbolV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CopytradingGetCopyTradingFuturesLeadSymbolV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CopytradingGetCopyTradingFuturesLeadSymbolV1`: CopytradingGetCopyTradingFuturesLeadSymbolV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.CopytradingGetCopyTradingFuturesLeadSymbolV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCopytradingGetCopyTradingFuturesLeadSymbolV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CopytradingGetCopyTradingFuturesLeadSymbolV1Resp**](CopytradingGetCopyTradingFuturesLeadSymbolV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CopytradingGetCopyTradingFuturesUserStatusV1

> CopytradingGetCopyTradingFuturesUserStatusV1Resp CopytradingGetCopyTradingFuturesUserStatusV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Futures Lead Trader Status(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/copytrading"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CopytradingGetCopyTradingFuturesUserStatusV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CopytradingGetCopyTradingFuturesUserStatusV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CopytradingGetCopyTradingFuturesUserStatusV1`: CopytradingGetCopyTradingFuturesUserStatusV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.CopytradingGetCopyTradingFuturesUserStatusV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCopytradingGetCopyTradingFuturesUserStatusV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CopytradingGetCopyTradingFuturesUserStatusV1Resp**](CopytradingGetCopyTradingFuturesUserStatusV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

