# \CopyTradingAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCopyTradingFuturesLeadSymbolV1**](CopyTradingAPI.md#GetCopyTradingFuturesLeadSymbolV1) | **Get** /sapi/v1/copyTrading/futures/leadSymbol | Get Futures Lead Trading Symbol Whitelist(USER_DATA)
[**GetCopyTradingFuturesUserStatusV1**](CopyTradingAPI.md#GetCopyTradingFuturesUserStatusV1) | **Get** /sapi/v1/copyTrading/futures/userStatus | Get Futures Lead Trader Status(TRADE)



## GetCopyTradingFuturesLeadSymbolV1

> GetCopyTradingFuturesLeadSymbolV1Resp GetCopyTradingFuturesLeadSymbolV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Futures Lead Trading Symbol Whitelist(USER_DATA)



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
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CopyTradingAPI.GetCopyTradingFuturesLeadSymbolV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CopyTradingAPI.GetCopyTradingFuturesLeadSymbolV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCopyTradingFuturesLeadSymbolV1`: GetCopyTradingFuturesLeadSymbolV1Resp
	fmt.Fprintf(os.Stdout, "Response from `CopyTradingAPI.GetCopyTradingFuturesLeadSymbolV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCopyTradingFuturesLeadSymbolV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetCopyTradingFuturesLeadSymbolV1Resp**](GetCopyTradingFuturesLeadSymbolV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCopyTradingFuturesUserStatusV1

> GetCopyTradingFuturesUserStatusV1Resp GetCopyTradingFuturesUserStatusV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Futures Lead Trader Status(TRADE)



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
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CopyTradingAPI.GetCopyTradingFuturesUserStatusV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CopyTradingAPI.GetCopyTradingFuturesUserStatusV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCopyTradingFuturesUserStatusV1`: GetCopyTradingFuturesUserStatusV1Resp
	fmt.Fprintf(os.Stdout, "Response from `CopyTradingAPI.GetCopyTradingFuturesUserStatusV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCopyTradingFuturesUserStatusV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetCopyTradingFuturesUserStatusV1Resp**](GetCopyTradingFuturesUserStatusV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

