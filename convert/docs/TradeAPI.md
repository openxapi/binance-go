# \TradeAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConvertCreateConvertAcceptQuoteV1**](TradeAPI.md#ConvertCreateConvertAcceptQuoteV1) | **Post** /sapi/v1/convert/acceptQuote | Accept Quote (TRADE)
[**ConvertCreateConvertGetQuoteV1**](TradeAPI.md#ConvertCreateConvertGetQuoteV1) | **Post** /sapi/v1/convert/getQuote | Send Quote Request(USER_DATA)
[**ConvertCreateConvertLimitCancelOrderV1**](TradeAPI.md#ConvertCreateConvertLimitCancelOrderV1) | **Post** /sapi/v1/convert/limit/cancelOrder | Cancel limit order (USER_DATA)
[**ConvertCreateConvertLimitPlaceOrderV1**](TradeAPI.md#ConvertCreateConvertLimitPlaceOrderV1) | **Post** /sapi/v1/convert/limit/placeOrder | Place limit order (USER_DATA)
[**ConvertCreateConvertLimitQueryOpenOrdersV1**](TradeAPI.md#ConvertCreateConvertLimitQueryOpenOrdersV1) | **Post** /sapi/v1/convert/limit/queryOpenOrders | Query limit open orders (USER_DATA)
[**ConvertGetConvertOrderStatusV1**](TradeAPI.md#ConvertGetConvertOrderStatusV1) | **Get** /sapi/v1/convert/orderStatus | Order status(USER_DATA)
[**ConvertGetConvertTradeFlowV1**](TradeAPI.md#ConvertGetConvertTradeFlowV1) | **Get** /sapi/v1/convert/tradeFlow | Get Convert Trade History(USER_DATA)



## ConvertCreateConvertAcceptQuoteV1

> ConvertCreateConvertAcceptQuoteV1Resp ConvertCreateConvertAcceptQuoteV1(ctx).QuoteId(quoteId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Accept Quote (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/convert"
)

func main() {
	quoteId := "quoteId_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.ConvertCreateConvertAcceptQuoteV1(context.Background()).QuoteId(quoteId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.ConvertCreateConvertAcceptQuoteV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConvertCreateConvertAcceptQuoteV1`: ConvertCreateConvertAcceptQuoteV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.ConvertCreateConvertAcceptQuoteV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConvertCreateConvertAcceptQuoteV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quoteId** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**ConvertCreateConvertAcceptQuoteV1Resp**](ConvertCreateConvertAcceptQuoteV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConvertCreateConvertGetQuoteV1

> ConvertCreateConvertGetQuoteV1Resp ConvertCreateConvertGetQuoteV1(ctx).FromAsset(fromAsset).Timestamp(timestamp).ToAsset(toAsset).FromAmount(fromAmount).RecvWindow(recvWindow).ToAmount(toAmount).ValidTime(validTime).WalletType(walletType).Execute()

Send Quote Request(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/convert"
)

func main() {
	fromAsset := "fromAsset_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	toAsset := "toAsset_example" // string |  (default to "")
	fromAmount := "fromAmount_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	toAmount := "toAmount_example" // string |  (optional) (default to "")
	validTime := "validTime_example" // string |  (optional) (default to "10")
	walletType := "walletType_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.ConvertCreateConvertGetQuoteV1(context.Background()).FromAsset(fromAsset).Timestamp(timestamp).ToAsset(toAsset).FromAmount(fromAmount).RecvWindow(recvWindow).ToAmount(toAmount).ValidTime(validTime).WalletType(walletType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.ConvertCreateConvertGetQuoteV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConvertCreateConvertGetQuoteV1`: ConvertCreateConvertGetQuoteV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.ConvertCreateConvertGetQuoteV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConvertCreateConvertGetQuoteV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromAsset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **toAsset** | **string** |  | [default to &quot;&quot;]
 **fromAmount** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **toAmount** | **string** |  | [default to &quot;&quot;]
 **validTime** | **string** |  | [default to &quot;10&quot;]
 **walletType** | **string** |  | [default to &quot;&quot;]

### Return type

[**ConvertCreateConvertGetQuoteV1Resp**](ConvertCreateConvertGetQuoteV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConvertCreateConvertLimitCancelOrderV1

> ConvertCreateConvertLimitCancelOrderV1Resp ConvertCreateConvertLimitCancelOrderV1(ctx).OrderId(orderId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Cancel limit order (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/convert"
)

func main() {
	orderId := int64(789) // int64 | 
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.ConvertCreateConvertLimitCancelOrderV1(context.Background()).OrderId(orderId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.ConvertCreateConvertLimitCancelOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConvertCreateConvertLimitCancelOrderV1`: ConvertCreateConvertLimitCancelOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.ConvertCreateConvertLimitCancelOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConvertCreateConvertLimitCancelOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **int64** |  | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**ConvertCreateConvertLimitCancelOrderV1Resp**](ConvertCreateConvertLimitCancelOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConvertCreateConvertLimitPlaceOrderV1

> ConvertCreateConvertLimitPlaceOrderV1Resp ConvertCreateConvertLimitPlaceOrderV1(ctx).BaseAsset(baseAsset).ExpiredType(expiredType).LimitPrice(limitPrice).QuoteAsset(quoteAsset).Side(side).Timestamp(timestamp).BaseAmount(baseAmount).QuoteAmount(quoteAmount).RecvWindow(recvWindow).WalletType(walletType).Execute()

Place limit order (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/convert"
)

func main() {
	baseAsset := "baseAsset_example" // string |  (default to "")
	expiredType := "expiredType_example" // string |  (default to "")
	limitPrice := "limitPrice_example" // string |  (default to "")
	quoteAsset := "quoteAsset_example" // string |  (default to "")
	side := "side_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	baseAmount := "baseAmount_example" // string |  (optional) (default to "")
	quoteAmount := "quoteAmount_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	walletType := "walletType_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.ConvertCreateConvertLimitPlaceOrderV1(context.Background()).BaseAsset(baseAsset).ExpiredType(expiredType).LimitPrice(limitPrice).QuoteAsset(quoteAsset).Side(side).Timestamp(timestamp).BaseAmount(baseAmount).QuoteAmount(quoteAmount).RecvWindow(recvWindow).WalletType(walletType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.ConvertCreateConvertLimitPlaceOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConvertCreateConvertLimitPlaceOrderV1`: ConvertCreateConvertLimitPlaceOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.ConvertCreateConvertLimitPlaceOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConvertCreateConvertLimitPlaceOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseAsset** | **string** |  | [default to &quot;&quot;]
 **expiredType** | **string** |  | [default to &quot;&quot;]
 **limitPrice** | **string** |  | [default to &quot;&quot;]
 **quoteAsset** | **string** |  | [default to &quot;&quot;]
 **side** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **baseAmount** | **string** |  | [default to &quot;&quot;]
 **quoteAmount** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **walletType** | **string** |  | [default to &quot;&quot;]

### Return type

[**ConvertCreateConvertLimitPlaceOrderV1Resp**](ConvertCreateConvertLimitPlaceOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConvertCreateConvertLimitQueryOpenOrdersV1

> ConvertCreateConvertLimitQueryOpenOrdersV1Resp ConvertCreateConvertLimitQueryOpenOrdersV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query limit open orders (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/convert"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.ConvertCreateConvertLimitQueryOpenOrdersV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.ConvertCreateConvertLimitQueryOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConvertCreateConvertLimitQueryOpenOrdersV1`: ConvertCreateConvertLimitQueryOpenOrdersV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.ConvertCreateConvertLimitQueryOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConvertCreateConvertLimitQueryOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**ConvertCreateConvertLimitQueryOpenOrdersV1Resp**](ConvertCreateConvertLimitQueryOpenOrdersV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConvertGetConvertOrderStatusV1

> ConvertGetConvertOrderStatusV1Resp ConvertGetConvertOrderStatusV1(ctx).OrderId(orderId).QuoteId(quoteId).Execute()

Order status(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/convert"
)

func main() {
	orderId := "orderId_example" // string | Either orderId or quoteId is required (optional) (default to "")
	quoteId := "quoteId_example" // string | Either orderId or quoteId is required (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.ConvertGetConvertOrderStatusV1(context.Background()).OrderId(orderId).QuoteId(quoteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.ConvertGetConvertOrderStatusV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConvertGetConvertOrderStatusV1`: ConvertGetConvertOrderStatusV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.ConvertGetConvertOrderStatusV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConvertGetConvertOrderStatusV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **string** | Either orderId or quoteId is required | [default to &quot;&quot;]
 **quoteId** | **string** | Either orderId or quoteId is required | [default to &quot;&quot;]

### Return type

[**ConvertGetConvertOrderStatusV1Resp**](ConvertGetConvertOrderStatusV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConvertGetConvertTradeFlowV1

> ConvertGetConvertTradeFlowV1Resp ConvertGetConvertTradeFlowV1(ctx).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).Limit(limit).RecvWindow(recvWindow).Execute()

Get Convert Trade History(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/convert"
)

func main() {
	startTime := int64(789) // int64 | 
	endTime := int64(789) // int64 | 
	timestamp := int64(789) // int64 | 
	limit := int32(56) // int32 | Default 100, Max 1000 (optional) (default to 100)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.ConvertGetConvertTradeFlowV1(context.Background()).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.ConvertGetConvertTradeFlowV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConvertGetConvertTradeFlowV1`: ConvertGetConvertTradeFlowV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.ConvertGetConvertTradeFlowV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConvertGetConvertTradeFlowV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **timestamp** | **int64** |  | 
 **limit** | **int32** | Default 100, Max 1000 | [default to 100]
 **recvWindow** | **int64** |  | 

### Return type

[**ConvertGetConvertTradeFlowV1Resp**](ConvertGetConvertTradeFlowV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

