# \ConvertAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConvertAcceptQuoteV1**](ConvertAPI.md#CreateConvertAcceptQuoteV1) | **Post** /sapi/v1/convert/acceptQuote | Accept Quote (TRADE)
[**CreateConvertGetQuoteV1**](ConvertAPI.md#CreateConvertGetQuoteV1) | **Post** /sapi/v1/convert/getQuote | Send Quote Request(USER_DATA)
[**CreateConvertLimitCancelOrderV1**](ConvertAPI.md#CreateConvertLimitCancelOrderV1) | **Post** /sapi/v1/convert/limit/cancelOrder | Cancel limit order (USER_DATA)
[**CreateConvertLimitPlaceOrderV1**](ConvertAPI.md#CreateConvertLimitPlaceOrderV1) | **Post** /sapi/v1/convert/limit/placeOrder | Place limit order (USER_DATA)
[**CreateConvertLimitQueryOpenOrdersV1**](ConvertAPI.md#CreateConvertLimitQueryOpenOrdersV1) | **Post** /sapi/v1/convert/limit/queryOpenOrders | Query limit open orders (USER_DATA)
[**GetConvertAssetInfoV1**](ConvertAPI.md#GetConvertAssetInfoV1) | **Get** /sapi/v1/convert/assetInfo | Query order quantity precision per asset(USER_DATA)
[**GetConvertExchangeInfoV1**](ConvertAPI.md#GetConvertExchangeInfoV1) | **Get** /sapi/v1/convert/exchangeInfo | List All Convert Pairs
[**GetConvertOrderStatusV1**](ConvertAPI.md#GetConvertOrderStatusV1) | **Get** /sapi/v1/convert/orderStatus | Order status(USER_DATA)
[**GetConvertTradeFlowV1**](ConvertAPI.md#GetConvertTradeFlowV1) | **Get** /sapi/v1/convert/tradeFlow | Get Convert Trade History(USER_DATA)



## CreateConvertAcceptQuoteV1

> CreateConvertAcceptQuoteV1Resp CreateConvertAcceptQuoteV1(ctx).QuoteId(quoteId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Accept Quote (TRADE)



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
	quoteId := "quoteId_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConvertAPI.CreateConvertAcceptQuoteV1(context.Background()).QuoteId(quoteId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConvertAPI.CreateConvertAcceptQuoteV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConvertAcceptQuoteV1`: CreateConvertAcceptQuoteV1Resp
	fmt.Fprintf(os.Stdout, "Response from `ConvertAPI.CreateConvertAcceptQuoteV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConvertAcceptQuoteV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quoteId** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateConvertAcceptQuoteV1Resp**](CreateConvertAcceptQuoteV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConvertGetQuoteV1

> CreateConvertGetQuoteV1Resp CreateConvertGetQuoteV1(ctx).FromAsset(fromAsset).Timestamp(timestamp).ToAsset(toAsset).FromAmount(fromAmount).RecvWindow(recvWindow).ToAmount(toAmount).ValidTime(validTime).WalletType(walletType).Execute()

Send Quote Request(USER_DATA)



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
	resp, r, err := apiClient.ConvertAPI.CreateConvertGetQuoteV1(context.Background()).FromAsset(fromAsset).Timestamp(timestamp).ToAsset(toAsset).FromAmount(fromAmount).RecvWindow(recvWindow).ToAmount(toAmount).ValidTime(validTime).WalletType(walletType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConvertAPI.CreateConvertGetQuoteV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConvertGetQuoteV1`: CreateConvertGetQuoteV1Resp
	fmt.Fprintf(os.Stdout, "Response from `ConvertAPI.CreateConvertGetQuoteV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConvertGetQuoteV1Request struct via the builder pattern


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

[**CreateConvertGetQuoteV1Resp**](CreateConvertGetQuoteV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConvertLimitCancelOrderV1

> CreateConvertLimitCancelOrderV1Resp CreateConvertLimitCancelOrderV1(ctx).OrderId(orderId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Cancel limit order (USER_DATA)



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
	orderId := int64(789) // int64 | 
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConvertAPI.CreateConvertLimitCancelOrderV1(context.Background()).OrderId(orderId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConvertAPI.CreateConvertLimitCancelOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConvertLimitCancelOrderV1`: CreateConvertLimitCancelOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `ConvertAPI.CreateConvertLimitCancelOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConvertLimitCancelOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **int64** |  | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateConvertLimitCancelOrderV1Resp**](CreateConvertLimitCancelOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConvertLimitPlaceOrderV1

> CreateConvertLimitPlaceOrderV1Resp CreateConvertLimitPlaceOrderV1(ctx).BaseAsset(baseAsset).ExpiredType(expiredType).LimitPrice(limitPrice).QuoteAsset(quoteAsset).Side(side).Timestamp(timestamp).BaseAmount(baseAmount).QuoteAmount(quoteAmount).RecvWindow(recvWindow).WalletType(walletType).Execute()

Place limit order (USER_DATA)



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
	resp, r, err := apiClient.ConvertAPI.CreateConvertLimitPlaceOrderV1(context.Background()).BaseAsset(baseAsset).ExpiredType(expiredType).LimitPrice(limitPrice).QuoteAsset(quoteAsset).Side(side).Timestamp(timestamp).BaseAmount(baseAmount).QuoteAmount(quoteAmount).RecvWindow(recvWindow).WalletType(walletType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConvertAPI.CreateConvertLimitPlaceOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConvertLimitPlaceOrderV1`: CreateConvertLimitPlaceOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `ConvertAPI.CreateConvertLimitPlaceOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConvertLimitPlaceOrderV1Request struct via the builder pattern


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

[**CreateConvertLimitPlaceOrderV1Resp**](CreateConvertLimitPlaceOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConvertLimitQueryOpenOrdersV1

> CreateConvertLimitQueryOpenOrdersV1Resp CreateConvertLimitQueryOpenOrdersV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query limit open orders (USER_DATA)



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
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConvertAPI.CreateConvertLimitQueryOpenOrdersV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConvertAPI.CreateConvertLimitQueryOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConvertLimitQueryOpenOrdersV1`: CreateConvertLimitQueryOpenOrdersV1Resp
	fmt.Fprintf(os.Stdout, "Response from `ConvertAPI.CreateConvertLimitQueryOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConvertLimitQueryOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateConvertLimitQueryOpenOrdersV1Resp**](CreateConvertLimitQueryOpenOrdersV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConvertAssetInfoV1

> []GetConvertAssetInfoV1RespItem GetConvertAssetInfoV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query order quantity precision per asset(USER_DATA)



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
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConvertAPI.GetConvertAssetInfoV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConvertAPI.GetConvertAssetInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConvertAssetInfoV1`: []GetConvertAssetInfoV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `ConvertAPI.GetConvertAssetInfoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConvertAssetInfoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]GetConvertAssetInfoV1RespItem**](GetConvertAssetInfoV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConvertExchangeInfoV1

> []GetConvertExchangeInfoV1RespItem GetConvertExchangeInfoV1(ctx).FromAsset(fromAsset).ToAsset(toAsset).Execute()

List All Convert Pairs



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
	fromAsset := "fromAsset_example" // string | User spends coin (optional) (default to "")
	toAsset := "toAsset_example" // string | User receives coin (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConvertAPI.GetConvertExchangeInfoV1(context.Background()).FromAsset(fromAsset).ToAsset(toAsset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConvertAPI.GetConvertExchangeInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConvertExchangeInfoV1`: []GetConvertExchangeInfoV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `ConvertAPI.GetConvertExchangeInfoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConvertExchangeInfoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromAsset** | **string** | User spends coin | [default to &quot;&quot;]
 **toAsset** | **string** | User receives coin | [default to &quot;&quot;]

### Return type

[**[]GetConvertExchangeInfoV1RespItem**](GetConvertExchangeInfoV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConvertOrderStatusV1

> GetConvertOrderStatusV1Resp GetConvertOrderStatusV1(ctx).OrderId(orderId).QuoteId(quoteId).Execute()

Order status(USER_DATA)



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
	orderId := "orderId_example" // string | Either orderId or quoteId is required (optional) (default to "")
	quoteId := "quoteId_example" // string | Either orderId or quoteId is required (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConvertAPI.GetConvertOrderStatusV1(context.Background()).OrderId(orderId).QuoteId(quoteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConvertAPI.GetConvertOrderStatusV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConvertOrderStatusV1`: GetConvertOrderStatusV1Resp
	fmt.Fprintf(os.Stdout, "Response from `ConvertAPI.GetConvertOrderStatusV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConvertOrderStatusV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **string** | Either orderId or quoteId is required | [default to &quot;&quot;]
 **quoteId** | **string** | Either orderId or quoteId is required | [default to &quot;&quot;]

### Return type

[**GetConvertOrderStatusV1Resp**](GetConvertOrderStatusV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConvertTradeFlowV1

> GetConvertTradeFlowV1Resp GetConvertTradeFlowV1(ctx).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).Limit(limit).RecvWindow(recvWindow).Execute()

Get Convert Trade History(USER_DATA)



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
	startTime := int64(789) // int64 | 
	endTime := int64(789) // int64 | 
	timestamp := int64(789) // int64 | 
	limit := int32(56) // int32 | Default 100, Max 1000 (optional) (default to 100)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConvertAPI.GetConvertTradeFlowV1(context.Background()).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConvertAPI.GetConvertTradeFlowV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConvertTradeFlowV1`: GetConvertTradeFlowV1Resp
	fmt.Fprintf(os.Stdout, "Response from `ConvertAPI.GetConvertTradeFlowV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConvertTradeFlowV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **timestamp** | **int64** |  | 
 **limit** | **int32** | Default 100, Max 1000 | [default to 100]
 **recvWindow** | **int64** |  | 

### Return type

[**GetConvertTradeFlowV1Resp**](GetConvertTradeFlowV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

