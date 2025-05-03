# \TradingAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SpotCreateOrderCancelReplaceV3**](TradingAPI.md#SpotCreateOrderCancelReplaceV3) | **Post** /api/v3/order/cancelReplace | Cancel an Existing Order and Send a New Order (TRADE)
[**SpotCreateOrderTestV3**](TradingAPI.md#SpotCreateOrderTestV3) | **Post** /api/v3/order/test | Test new order (TRADE)
[**SpotCreateOrderV3**](TradingAPI.md#SpotCreateOrderV3) | **Post** /api/v3/order | New order (TRADE)
[**SpotCreateSorOrderTestV3**](TradingAPI.md#SpotCreateSorOrderTestV3) | **Post** /api/v3/sor/order/test | Test new order using SOR (TRADE)
[**SpotDeleteOpenOrdersV3**](TradingAPI.md#SpotDeleteOpenOrdersV3) | **Delete** /api/v3/openOrders | Cancel All Open Orders on a Symbol (TRADE)



## SpotCreateOrderCancelReplaceV3

> SpotCreateOrderCancelReplaceV3Resp SpotCreateOrderCancelReplaceV3(ctx).CancelReplaceMode(cancelReplaceMode).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).CancelNewClientOrderId(cancelNewClientOrderId).CancelOrderId(cancelOrderId).CancelOrigClientOrderId(cancelOrigClientOrderId).CancelRestrictions(cancelRestrictions).IcebergQty(icebergQty).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).OrderRateLimitExceededMode(orderRateLimitExceededMode).Price(price).Quantity(quantity).QuoteOrderQty(quoteOrderQty).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).StopPrice(stopPrice).StrategyId(strategyId).StrategyType(strategyType).TimeInForce(timeInForce).TrailingDelta(trailingDelta).Execute()

Cancel an Existing Order and Send a New Order (TRADE)



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
	cancelReplaceMode := "cancelReplaceMode_example" // string |  (default to "")
	side := "side_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	type_ := "type__example" // string |  (default to "")
	cancelNewClientOrderId := "cancelNewClientOrderId_example" // string |  (optional) (default to "")
	cancelOrderId := int64(789) // int64 |  (optional)
	cancelOrigClientOrderId := "cancelOrigClientOrderId_example" // string |  (optional) (default to "")
	cancelRestrictions := "cancelRestrictions_example" // string |  (optional) (default to "")
	icebergQty := "icebergQty_example" // string |  (optional) (default to "")
	newClientOrderId := "newClientOrderId_example" // string |  (optional) (default to "")
	newOrderRespType := "newOrderRespType_example" // string |  (optional) (default to "")
	orderRateLimitExceededMode := "orderRateLimitExceededMode_example" // string |  (optional) (default to "")
	price := "price_example" // string |  (optional) (default to "")
	quantity := "quantity_example" // string |  (optional) (default to "")
	quoteOrderQty := "quoteOrderQty_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	selfTradePreventionMode := "selfTradePreventionMode_example" // string |  (optional) (default to "")
	stopPrice := "stopPrice_example" // string |  (optional) (default to "")
	strategyId := int64(789) // int64 |  (optional)
	strategyType := int32(56) // int32 |  (optional)
	timeInForce := "timeInForce_example" // string |  (optional) (default to "")
	trailingDelta := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAPI.SpotCreateOrderCancelReplaceV3(context.Background()).CancelReplaceMode(cancelReplaceMode).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).CancelNewClientOrderId(cancelNewClientOrderId).CancelOrderId(cancelOrderId).CancelOrigClientOrderId(cancelOrigClientOrderId).CancelRestrictions(cancelRestrictions).IcebergQty(icebergQty).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).OrderRateLimitExceededMode(orderRateLimitExceededMode).Price(price).Quantity(quantity).QuoteOrderQty(quoteOrderQty).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).StopPrice(stopPrice).StrategyId(strategyId).StrategyType(strategyType).TimeInForce(timeInForce).TrailingDelta(trailingDelta).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAPI.SpotCreateOrderCancelReplaceV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotCreateOrderCancelReplaceV3`: SpotCreateOrderCancelReplaceV3Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAPI.SpotCreateOrderCancelReplaceV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotCreateOrderCancelReplaceV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cancelReplaceMode** | **string** |  | [default to &quot;&quot;]
 **side** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **type_** | **string** |  | [default to &quot;&quot;]
 **cancelNewClientOrderId** | **string** |  | [default to &quot;&quot;]
 **cancelOrderId** | **int64** |  | 
 **cancelOrigClientOrderId** | **string** |  | [default to &quot;&quot;]
 **cancelRestrictions** | **string** |  | [default to &quot;&quot;]
 **icebergQty** | **string** |  | [default to &quot;&quot;]
 **newClientOrderId** | **string** |  | [default to &quot;&quot;]
 **newOrderRespType** | **string** |  | [default to &quot;&quot;]
 **orderRateLimitExceededMode** | **string** |  | [default to &quot;&quot;]
 **price** | **string** |  | [default to &quot;&quot;]
 **quantity** | **string** |  | [default to &quot;&quot;]
 **quoteOrderQty** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **selfTradePreventionMode** | **string** |  | [default to &quot;&quot;]
 **stopPrice** | **string** |  | [default to &quot;&quot;]
 **strategyId** | **int64** |  | 
 **strategyType** | **int32** |  | 
 **timeInForce** | **string** |  | [default to &quot;&quot;]
 **trailingDelta** | **int64** |  | 

### Return type

[**SpotCreateOrderCancelReplaceV3Resp**](SpotCreateOrderCancelReplaceV3Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpotCreateOrderTestV3

> SpotCreateOrderTestV3Resp SpotCreateOrderTestV3(ctx).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).ComputeCommissionRates(computeCommissionRates).IcebergQty(icebergQty).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).Price(price).Quantity(quantity).QuoteOrderQty(quoteOrderQty).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).StopPrice(stopPrice).StrategyId(strategyId).StrategyType(strategyType).TimeInForce(timeInForce).TrailingDelta(trailingDelta).Execute()

Test new order (TRADE)



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
	side := "side_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	type_ := "type__example" // string |  (default to "")
	computeCommissionRates := true // bool |  (optional)
	icebergQty := "icebergQty_example" // string |  (optional) (default to "")
	newClientOrderId := "newClientOrderId_example" // string |  (optional) (default to "")
	newOrderRespType := "newOrderRespType_example" // string |  (optional) (default to "")
	price := "price_example" // string |  (optional) (default to "")
	quantity := "quantity_example" // string |  (optional) (default to "")
	quoteOrderQty := "quoteOrderQty_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	selfTradePreventionMode := "selfTradePreventionMode_example" // string |  (optional) (default to "")
	stopPrice := "stopPrice_example" // string |  (optional) (default to "")
	strategyId := int64(789) // int64 |  (optional)
	strategyType := int32(56) // int32 |  (optional)
	timeInForce := "timeInForce_example" // string |  (optional) (default to "")
	trailingDelta := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAPI.SpotCreateOrderTestV3(context.Background()).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).ComputeCommissionRates(computeCommissionRates).IcebergQty(icebergQty).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).Price(price).Quantity(quantity).QuoteOrderQty(quoteOrderQty).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).StopPrice(stopPrice).StrategyId(strategyId).StrategyType(strategyType).TimeInForce(timeInForce).TrailingDelta(trailingDelta).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAPI.SpotCreateOrderTestV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotCreateOrderTestV3`: SpotCreateOrderTestV3Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAPI.SpotCreateOrderTestV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotCreateOrderTestV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **side** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **type_** | **string** |  | [default to &quot;&quot;]
 **computeCommissionRates** | **bool** |  | 
 **icebergQty** | **string** |  | [default to &quot;&quot;]
 **newClientOrderId** | **string** |  | [default to &quot;&quot;]
 **newOrderRespType** | **string** |  | [default to &quot;&quot;]
 **price** | **string** |  | [default to &quot;&quot;]
 **quantity** | **string** |  | [default to &quot;&quot;]
 **quoteOrderQty** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **selfTradePreventionMode** | **string** |  | [default to &quot;&quot;]
 **stopPrice** | **string** |  | [default to &quot;&quot;]
 **strategyId** | **int64** |  | 
 **strategyType** | **int32** |  | 
 **timeInForce** | **string** |  | [default to &quot;&quot;]
 **trailingDelta** | **int64** |  | 

### Return type

[**SpotCreateOrderTestV3Resp**](SpotCreateOrderTestV3Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpotCreateOrderV3

> SpotCreateOrderV3Resp SpotCreateOrderV3(ctx).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).IcebergQty(icebergQty).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).Price(price).Quantity(quantity).QuoteOrderQty(quoteOrderQty).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).StopPrice(stopPrice).StrategyId(strategyId).StrategyType(strategyType).TimeInForce(timeInForce).TrailingDelta(trailingDelta).Execute()

New order (TRADE)



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
	side := "side_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	type_ := "type__example" // string |  (default to "")
	icebergQty := "icebergQty_example" // string |  (optional) (default to "")
	newClientOrderId := "newClientOrderId_example" // string |  (optional) (default to "")
	newOrderRespType := "newOrderRespType_example" // string |  (optional) (default to "")
	price := "price_example" // string |  (optional) (default to "")
	quantity := "quantity_example" // string |  (optional) (default to "")
	quoteOrderQty := "quoteOrderQty_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	selfTradePreventionMode := "selfTradePreventionMode_example" // string |  (optional) (default to "")
	stopPrice := "stopPrice_example" // string |  (optional) (default to "")
	strategyId := int64(789) // int64 |  (optional)
	strategyType := int32(56) // int32 |  (optional)
	timeInForce := "timeInForce_example" // string |  (optional) (default to "")
	trailingDelta := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAPI.SpotCreateOrderV3(context.Background()).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).IcebergQty(icebergQty).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).Price(price).Quantity(quantity).QuoteOrderQty(quoteOrderQty).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).StopPrice(stopPrice).StrategyId(strategyId).StrategyType(strategyType).TimeInForce(timeInForce).TrailingDelta(trailingDelta).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAPI.SpotCreateOrderV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotCreateOrderV3`: SpotCreateOrderV3Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAPI.SpotCreateOrderV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotCreateOrderV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **side** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **type_** | **string** |  | [default to &quot;&quot;]
 **icebergQty** | **string** |  | [default to &quot;&quot;]
 **newClientOrderId** | **string** |  | [default to &quot;&quot;]
 **newOrderRespType** | **string** |  | [default to &quot;&quot;]
 **price** | **string** |  | [default to &quot;&quot;]
 **quantity** | **string** |  | [default to &quot;&quot;]
 **quoteOrderQty** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **selfTradePreventionMode** | **string** |  | [default to &quot;&quot;]
 **stopPrice** | **string** |  | [default to &quot;&quot;]
 **strategyId** | **int64** |  | 
 **strategyType** | **int32** |  | 
 **timeInForce** | **string** |  | [default to &quot;&quot;]
 **trailingDelta** | **int64** |  | 

### Return type

[**SpotCreateOrderV3Resp**](SpotCreateOrderV3Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpotCreateSorOrderTestV3

> SpotCreateSorOrderTestV3Resp SpotCreateSorOrderTestV3(ctx).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).ComputeCommissionRates(computeCommissionRates).IcebergQty(icebergQty).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).Price(price).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).StrategyId(strategyId).StrategyType(strategyType).TimeInForce(timeInForce).Execute()

Test new order using SOR (TRADE)



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
	quantity := "quantity_example" // string |  (default to "")
	side := "side_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	type_ := "type__example" // string |  (default to "")
	computeCommissionRates := true // bool |  (optional)
	icebergQty := "icebergQty_example" // string |  (optional) (default to "")
	newClientOrderId := "newClientOrderId_example" // string |  (optional) (default to "")
	newOrderRespType := "newOrderRespType_example" // string |  (optional) (default to "")
	price := "price_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	selfTradePreventionMode := "selfTradePreventionMode_example" // string |  (optional) (default to "")
	strategyId := int64(789) // int64 |  (optional)
	strategyType := int32(56) // int32 |  (optional)
	timeInForce := "timeInForce_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAPI.SpotCreateSorOrderTestV3(context.Background()).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).ComputeCommissionRates(computeCommissionRates).IcebergQty(icebergQty).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).Price(price).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).StrategyId(strategyId).StrategyType(strategyType).TimeInForce(timeInForce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAPI.SpotCreateSorOrderTestV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotCreateSorOrderTestV3`: SpotCreateSorOrderTestV3Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAPI.SpotCreateSorOrderTestV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotCreateSorOrderTestV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quantity** | **string** |  | [default to &quot;&quot;]
 **side** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **type_** | **string** |  | [default to &quot;&quot;]
 **computeCommissionRates** | **bool** |  | 
 **icebergQty** | **string** |  | [default to &quot;&quot;]
 **newClientOrderId** | **string** |  | [default to &quot;&quot;]
 **newOrderRespType** | **string** |  | [default to &quot;&quot;]
 **price** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **selfTradePreventionMode** | **string** |  | [default to &quot;&quot;]
 **strategyId** | **int64** |  | 
 **strategyType** | **int32** |  | 
 **timeInForce** | **string** |  | [default to &quot;&quot;]

### Return type

[**SpotCreateSorOrderTestV3Resp**](SpotCreateSorOrderTestV3Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpotDeleteOpenOrdersV3

> [][]SpotDeleteOpenOrdersV3RespInner SpotDeleteOpenOrdersV3(ctx).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Cancel All Open Orders on a Symbol (TRADE)



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
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAPI.SpotDeleteOpenOrdersV3(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAPI.SpotDeleteOpenOrdersV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotDeleteOpenOrdersV3`: [][]SpotDeleteOpenOrdersV3RespInner
	fmt.Fprintf(os.Stdout, "Response from `TradingAPI.SpotDeleteOpenOrdersV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotDeleteOpenOrdersV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[][]SpotDeleteOpenOrdersV3RespInner**](array.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

