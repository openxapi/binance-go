# \MarketDataAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FuturesdataGetFuturesHistDataLinkV1**](MarketDataAPI.md#FuturesdataGetFuturesHistDataLinkV1) | **Get** /sapi/v1/futures/histDataLink | Get Future TickLevel Orderbook Historical Data Download Link(USER_DATA)



## FuturesdataGetFuturesHistDataLinkV1

> FuturesdataGetFuturesHistDataLinkV1Resp FuturesdataGetFuturesHistDataLinkV1(ctx).Symbol(symbol).DataType(dataType).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Future TickLevel Orderbook Historical Data Download Link(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/futuresdata"
)

func main() {
	symbol := "symbol_example" // string | symbol name, e.g. BTCUSDT or BTCUSD_PERP ｜ (default to "")
	dataType := "dataType_example" // string | `T_DEPTH` for ticklevel orderbook data, `S_DEPTH` for orderbook snapshot data (default to "")
	startTime := int64(789) // int64 | 
	endTime := int64(789) // int64 | 
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.FuturesdataGetFuturesHistDataLinkV1(context.Background()).Symbol(symbol).DataType(dataType).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.FuturesdataGetFuturesHistDataLinkV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FuturesdataGetFuturesHistDataLinkV1`: FuturesdataGetFuturesHistDataLinkV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.FuturesdataGetFuturesHistDataLinkV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFuturesdataGetFuturesHistDataLinkV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | symbol name, e.g. BTCUSDT or BTCUSD_PERP ｜ | [default to &quot;&quot;]
 **dataType** | **string** | &#x60;T_DEPTH&#x60; for ticklevel orderbook data, &#x60;S_DEPTH&#x60; for orderbook snapshot data | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**FuturesdataGetFuturesHistDataLinkV1Resp**](FuturesdataGetFuturesHistDataLinkV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

