# \TradingAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SpotCreateOrderCancelReplaceV3**](TradingAPI.md#SpotCreateOrderCancelReplaceV3) | **Post** /api/v3/order/cancelReplace | Cancel an Existing Order and Send a New Order (TRADE)
[**SpotCreateOrderListOcoV3**](TradingAPI.md#SpotCreateOrderListOcoV3) | **Post** /api/v3/orderList/oco | New Order list - OCO (TRADE)
[**SpotCreateOrderListOtoV3**](TradingAPI.md#SpotCreateOrderListOtoV3) | **Post** /api/v3/orderList/oto | New Order list - OTO (TRADE)
[**SpotCreateOrderListOtocoV3**](TradingAPI.md#SpotCreateOrderListOtocoV3) | **Post** /api/v3/orderList/otoco | New Order list - OTOCO (TRADE)
[**SpotCreateOrderOcoV3**](TradingAPI.md#SpotCreateOrderOcoV3) | **Post** /api/v3/order/oco | New OCO - Deprecated (TRADE)
[**SpotCreateOrderTestV3**](TradingAPI.md#SpotCreateOrderTestV3) | **Post** /api/v3/order/test | Test new order (TRADE)
[**SpotCreateOrderV3**](TradingAPI.md#SpotCreateOrderV3) | **Post** /api/v3/order | New order (TRADE)
[**SpotCreateSorOrderTestV3**](TradingAPI.md#SpotCreateSorOrderTestV3) | **Post** /api/v3/sor/order/test | Test new order using SOR (TRADE)
[**SpotCreateSorOrderV3**](TradingAPI.md#SpotCreateSorOrderV3) | **Post** /api/v3/sor/order | New order using SOR (TRADE)
[**SpotDeleteOpenOrdersV3**](TradingAPI.md#SpotDeleteOpenOrdersV3) | **Delete** /api/v3/openOrders | Cancel All Open Orders on a Symbol (TRADE)
[**SpotDeleteOrderListV3**](TradingAPI.md#SpotDeleteOrderListV3) | **Delete** /api/v3/orderList | Cancel Order list (TRADE)
[**SpotDeleteOrderV3**](TradingAPI.md#SpotDeleteOrderV3) | **Delete** /api/v3/order | Cancel order (TRADE)
[**SpotGetAllOrderListV3**](TradingAPI.md#SpotGetAllOrderListV3) | **Get** /api/v3/allOrderList | Query all Order lists (USER_DATA)
[**SpotGetAllOrdersV3**](TradingAPI.md#SpotGetAllOrdersV3) | **Get** /api/v3/allOrders | All orders (USER_DATA)
[**SpotGetOpenOrderListV3**](TradingAPI.md#SpotGetOpenOrderListV3) | **Get** /api/v3/openOrderList | Query Open Order lists (USER_DATA)
[**SpotGetOpenOrdersV3**](TradingAPI.md#SpotGetOpenOrdersV3) | **Get** /api/v3/openOrders | Current open orders (USER_DATA)
[**SpotGetOrderListV3**](TradingAPI.md#SpotGetOrderListV3) | **Get** /api/v3/orderList | Query Order list (USER_DATA)
[**SpotGetOrderV3**](TradingAPI.md#SpotGetOrderV3) | **Get** /api/v3/order | Query order (USER_DATA)



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


## SpotCreateOrderListOcoV3

> SpotCreateOrderListOcoV3Resp SpotCreateOrderListOcoV3(ctx).AboveType(aboveType).BelowType(belowType).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).AboveClientOrderId(aboveClientOrderId).AboveIcebergQty(aboveIcebergQty).AbovePrice(abovePrice).AboveStopPrice(aboveStopPrice).AboveStrategyId(aboveStrategyId).AboveStrategyType(aboveStrategyType).AboveTimeInForce(aboveTimeInForce).AboveTrailingDelta(aboveTrailingDelta).BelowClientOrderId(belowClientOrderId).BelowIcebergQty(belowIcebergQty).BelowPrice(belowPrice).BelowStopPrice(belowStopPrice).BelowStrategyId(belowStrategyId).BelowStrategyType(belowStrategyType).BelowTimeInForce(belowTimeInForce).BelowTrailingDelta(belowTrailingDelta).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).Execute()

New Order list - OCO (TRADE)



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
	aboveType := "aboveType_example" // string |  (default to "")
	belowType := "belowType_example" // string |  (default to "")
	quantity := "quantity_example" // string |  (default to "")
	side := "side_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	aboveClientOrderId := "aboveClientOrderId_example" // string |  (optional) (default to "")
	aboveIcebergQty := int64(789) // int64 |  (optional)
	abovePrice := "abovePrice_example" // string |  (optional) (default to "")
	aboveStopPrice := "aboveStopPrice_example" // string |  (optional) (default to "")
	aboveStrategyId := int64(789) // int64 |  (optional)
	aboveStrategyType := int32(56) // int32 |  (optional)
	aboveTimeInForce := "aboveTimeInForce_example" // string |  (optional) (default to "")
	aboveTrailingDelta := int64(789) // int64 |  (optional)
	belowClientOrderId := "belowClientOrderId_example" // string |  (optional) (default to "")
	belowIcebergQty := int64(789) // int64 |  (optional)
	belowPrice := "belowPrice_example" // string |  (optional) (default to "")
	belowStopPrice := "belowStopPrice_example" // string |  (optional) (default to "")
	belowStrategyId := int64(789) // int64 |  (optional)
	belowStrategyType := int32(56) // int32 |  (optional)
	belowTimeInForce := "belowTimeInForce_example" // string |  (optional) (default to "")
	belowTrailingDelta := int64(789) // int64 |  (optional)
	listClientOrderId := "listClientOrderId_example" // string |  (optional) (default to "")
	newOrderRespType := "newOrderRespType_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	selfTradePreventionMode := "selfTradePreventionMode_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAPI.SpotCreateOrderListOcoV3(context.Background()).AboveType(aboveType).BelowType(belowType).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).AboveClientOrderId(aboveClientOrderId).AboveIcebergQty(aboveIcebergQty).AbovePrice(abovePrice).AboveStopPrice(aboveStopPrice).AboveStrategyId(aboveStrategyId).AboveStrategyType(aboveStrategyType).AboveTimeInForce(aboveTimeInForce).AboveTrailingDelta(aboveTrailingDelta).BelowClientOrderId(belowClientOrderId).BelowIcebergQty(belowIcebergQty).BelowPrice(belowPrice).BelowStopPrice(belowStopPrice).BelowStrategyId(belowStrategyId).BelowStrategyType(belowStrategyType).BelowTimeInForce(belowTimeInForce).BelowTrailingDelta(belowTrailingDelta).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAPI.SpotCreateOrderListOcoV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotCreateOrderListOcoV3`: SpotCreateOrderListOcoV3Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAPI.SpotCreateOrderListOcoV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotCreateOrderListOcoV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aboveType** | **string** |  | [default to &quot;&quot;]
 **belowType** | **string** |  | [default to &quot;&quot;]
 **quantity** | **string** |  | [default to &quot;&quot;]
 **side** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **aboveClientOrderId** | **string** |  | [default to &quot;&quot;]
 **aboveIcebergQty** | **int64** |  | 
 **abovePrice** | **string** |  | [default to &quot;&quot;]
 **aboveStopPrice** | **string** |  | [default to &quot;&quot;]
 **aboveStrategyId** | **int64** |  | 
 **aboveStrategyType** | **int32** |  | 
 **aboveTimeInForce** | **string** |  | [default to &quot;&quot;]
 **aboveTrailingDelta** | **int64** |  | 
 **belowClientOrderId** | **string** |  | [default to &quot;&quot;]
 **belowIcebergQty** | **int64** |  | 
 **belowPrice** | **string** |  | [default to &quot;&quot;]
 **belowStopPrice** | **string** |  | [default to &quot;&quot;]
 **belowStrategyId** | **int64** |  | 
 **belowStrategyType** | **int32** |  | 
 **belowTimeInForce** | **string** |  | [default to &quot;&quot;]
 **belowTrailingDelta** | **int64** |  | 
 **listClientOrderId** | **string** |  | [default to &quot;&quot;]
 **newOrderRespType** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **selfTradePreventionMode** | **string** |  | [default to &quot;&quot;]

### Return type

[**SpotCreateOrderListOcoV3Resp**](SpotCreateOrderListOcoV3Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpotCreateOrderListOtoV3

> SpotCreateOrderListOtoV3Resp SpotCreateOrderListOtoV3(ctx).PendingQuantity(pendingQuantity).PendingSide(pendingSide).PendingType(pendingType).Symbol(symbol).Timestamp(timestamp).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).WorkingSide(workingSide).WorkingType(workingType).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).PendingClientOrderId(pendingClientOrderId).PendingIcebergQty(pendingIcebergQty).PendingPrice(pendingPrice).PendingStopPrice(pendingStopPrice).PendingStrategyId(pendingStrategyId).PendingStrategyType(pendingStrategyType).PendingTimeInForce(pendingTimeInForce).PendingTrailingDelta(pendingTrailingDelta).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingStrategyId(workingStrategyId).WorkingStrategyType(workingStrategyType).WorkingTimeInForce(workingTimeInForce).Execute()

New Order list - OTO (TRADE)



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
	pendingQuantity := "pendingQuantity_example" // string |  (default to "")
	pendingSide := "pendingSide_example" // string |  (default to "")
	pendingType := "pendingType_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	workingPrice := "workingPrice_example" // string |  (default to "")
	workingQuantity := "workingQuantity_example" // string |  (default to "")
	workingSide := "workingSide_example" // string |  (default to "")
	workingType := "workingType_example" // string |  (default to "")
	listClientOrderId := "listClientOrderId_example" // string |  (optional) (default to "")
	newOrderRespType := "newOrderRespType_example" // string |  (optional) (default to "")
	pendingClientOrderId := "pendingClientOrderId_example" // string |  (optional) (default to "")
	pendingIcebergQty := "pendingIcebergQty_example" // string |  (optional) (default to "")
	pendingPrice := "pendingPrice_example" // string |  (optional) (default to "")
	pendingStopPrice := "pendingStopPrice_example" // string |  (optional) (default to "")
	pendingStrategyId := int64(789) // int64 |  (optional)
	pendingStrategyType := int32(56) // int32 |  (optional)
	pendingTimeInForce := "pendingTimeInForce_example" // string |  (optional) (default to "")
	pendingTrailingDelta := "pendingTrailingDelta_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	selfTradePreventionMode := "selfTradePreventionMode_example" // string |  (optional) (default to "")
	workingClientOrderId := "workingClientOrderId_example" // string |  (optional) (default to "")
	workingIcebergQty := "workingIcebergQty_example" // string |  (optional) (default to "")
	workingStrategyId := int64(789) // int64 |  (optional)
	workingStrategyType := int32(56) // int32 |  (optional)
	workingTimeInForce := "workingTimeInForce_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAPI.SpotCreateOrderListOtoV3(context.Background()).PendingQuantity(pendingQuantity).PendingSide(pendingSide).PendingType(pendingType).Symbol(symbol).Timestamp(timestamp).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).WorkingSide(workingSide).WorkingType(workingType).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).PendingClientOrderId(pendingClientOrderId).PendingIcebergQty(pendingIcebergQty).PendingPrice(pendingPrice).PendingStopPrice(pendingStopPrice).PendingStrategyId(pendingStrategyId).PendingStrategyType(pendingStrategyType).PendingTimeInForce(pendingTimeInForce).PendingTrailingDelta(pendingTrailingDelta).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingStrategyId(workingStrategyId).WorkingStrategyType(workingStrategyType).WorkingTimeInForce(workingTimeInForce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAPI.SpotCreateOrderListOtoV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotCreateOrderListOtoV3`: SpotCreateOrderListOtoV3Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAPI.SpotCreateOrderListOtoV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotCreateOrderListOtoV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pendingQuantity** | **string** |  | [default to &quot;&quot;]
 **pendingSide** | **string** |  | [default to &quot;&quot;]
 **pendingType** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **workingPrice** | **string** |  | [default to &quot;&quot;]
 **workingQuantity** | **string** |  | [default to &quot;&quot;]
 **workingSide** | **string** |  | [default to &quot;&quot;]
 **workingType** | **string** |  | [default to &quot;&quot;]
 **listClientOrderId** | **string** |  | [default to &quot;&quot;]
 **newOrderRespType** | **string** |  | [default to &quot;&quot;]
 **pendingClientOrderId** | **string** |  | [default to &quot;&quot;]
 **pendingIcebergQty** | **string** |  | [default to &quot;&quot;]
 **pendingPrice** | **string** |  | [default to &quot;&quot;]
 **pendingStopPrice** | **string** |  | [default to &quot;&quot;]
 **pendingStrategyId** | **int64** |  | 
 **pendingStrategyType** | **int32** |  | 
 **pendingTimeInForce** | **string** |  | [default to &quot;&quot;]
 **pendingTrailingDelta** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **selfTradePreventionMode** | **string** |  | [default to &quot;&quot;]
 **workingClientOrderId** | **string** |  | [default to &quot;&quot;]
 **workingIcebergQty** | **string** |  | [default to &quot;&quot;]
 **workingStrategyId** | **int64** |  | 
 **workingStrategyType** | **int32** |  | 
 **workingTimeInForce** | **string** |  | [default to &quot;&quot;]

### Return type

[**SpotCreateOrderListOtoV3Resp**](SpotCreateOrderListOtoV3Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpotCreateOrderListOtocoV3

> SpotCreateOrderListOtocoV3Resp SpotCreateOrderListOtocoV3(ctx).PendingAboveType(pendingAboveType).PendingQuantity(pendingQuantity).PendingSide(pendingSide).Symbol(symbol).Timestamp(timestamp).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).WorkingSide(workingSide).WorkingType(workingType).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).PendingAboveClientOrderId(pendingAboveClientOrderId).PendingAboveIcebergQty(pendingAboveIcebergQty).PendingAbovePrice(pendingAbovePrice).PendingAboveStopPrice(pendingAboveStopPrice).PendingAboveStrategyId(pendingAboveStrategyId).PendingAboveStrategyType(pendingAboveStrategyType).PendingAboveTimeInForce(pendingAboveTimeInForce).PendingAboveTrailingDelta(pendingAboveTrailingDelta).PendingBelowClientOrderId(pendingBelowClientOrderId).PendingBelowIcebergQty(pendingBelowIcebergQty).PendingBelowPrice(pendingBelowPrice).PendingBelowStopPrice(pendingBelowStopPrice).PendingBelowStrategyId(pendingBelowStrategyId).PendingBelowStrategyType(pendingBelowStrategyType).PendingBelowTimeInForce(pendingBelowTimeInForce).PendingBelowTrailingDelta(pendingBelowTrailingDelta).PendingBelowType(pendingBelowType).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingStrategyId(workingStrategyId).WorkingStrategyType(workingStrategyType).WorkingTimeInForce(workingTimeInForce).Execute()

New Order list - OTOCO (TRADE)



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
	pendingAboveType := "pendingAboveType_example" // string |  (default to "")
	pendingQuantity := "pendingQuantity_example" // string |  (default to "")
	pendingSide := "pendingSide_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	workingPrice := "workingPrice_example" // string |  (default to "")
	workingQuantity := "workingQuantity_example" // string |  (default to "")
	workingSide := "workingSide_example" // string |  (default to "")
	workingType := "workingType_example" // string |  (default to "")
	listClientOrderId := "listClientOrderId_example" // string |  (optional) (default to "")
	newOrderRespType := "newOrderRespType_example" // string |  (optional) (default to "")
	pendingAboveClientOrderId := "pendingAboveClientOrderId_example" // string |  (optional) (default to "")
	pendingAboveIcebergQty := "pendingAboveIcebergQty_example" // string |  (optional) (default to "")
	pendingAbovePrice := "pendingAbovePrice_example" // string |  (optional) (default to "")
	pendingAboveStopPrice := "pendingAboveStopPrice_example" // string |  (optional) (default to "")
	pendingAboveStrategyId := int64(789) // int64 |  (optional)
	pendingAboveStrategyType := int32(56) // int32 |  (optional)
	pendingAboveTimeInForce := "pendingAboveTimeInForce_example" // string |  (optional) (default to "")
	pendingAboveTrailingDelta := "pendingAboveTrailingDelta_example" // string |  (optional) (default to "")
	pendingBelowClientOrderId := "pendingBelowClientOrderId_example" // string |  (optional) (default to "")
	pendingBelowIcebergQty := "pendingBelowIcebergQty_example" // string |  (optional) (default to "")
	pendingBelowPrice := "pendingBelowPrice_example" // string |  (optional) (default to "")
	pendingBelowStopPrice := "pendingBelowStopPrice_example" // string |  (optional) (default to "")
	pendingBelowStrategyId := int64(789) // int64 |  (optional)
	pendingBelowStrategyType := int32(56) // int32 |  (optional)
	pendingBelowTimeInForce := "pendingBelowTimeInForce_example" // string |  (optional) (default to "")
	pendingBelowTrailingDelta := "pendingBelowTrailingDelta_example" // string |  (optional) (default to "")
	pendingBelowType := "pendingBelowType_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	selfTradePreventionMode := "selfTradePreventionMode_example" // string |  (optional) (default to "")
	workingClientOrderId := "workingClientOrderId_example" // string |  (optional) (default to "")
	workingIcebergQty := "workingIcebergQty_example" // string |  (optional) (default to "")
	workingStrategyId := int64(789) // int64 |  (optional)
	workingStrategyType := int32(56) // int32 |  (optional)
	workingTimeInForce := "workingTimeInForce_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAPI.SpotCreateOrderListOtocoV3(context.Background()).PendingAboveType(pendingAboveType).PendingQuantity(pendingQuantity).PendingSide(pendingSide).Symbol(symbol).Timestamp(timestamp).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).WorkingSide(workingSide).WorkingType(workingType).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).PendingAboveClientOrderId(pendingAboveClientOrderId).PendingAboveIcebergQty(pendingAboveIcebergQty).PendingAbovePrice(pendingAbovePrice).PendingAboveStopPrice(pendingAboveStopPrice).PendingAboveStrategyId(pendingAboveStrategyId).PendingAboveStrategyType(pendingAboveStrategyType).PendingAboveTimeInForce(pendingAboveTimeInForce).PendingAboveTrailingDelta(pendingAboveTrailingDelta).PendingBelowClientOrderId(pendingBelowClientOrderId).PendingBelowIcebergQty(pendingBelowIcebergQty).PendingBelowPrice(pendingBelowPrice).PendingBelowStopPrice(pendingBelowStopPrice).PendingBelowStrategyId(pendingBelowStrategyId).PendingBelowStrategyType(pendingBelowStrategyType).PendingBelowTimeInForce(pendingBelowTimeInForce).PendingBelowTrailingDelta(pendingBelowTrailingDelta).PendingBelowType(pendingBelowType).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingStrategyId(workingStrategyId).WorkingStrategyType(workingStrategyType).WorkingTimeInForce(workingTimeInForce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAPI.SpotCreateOrderListOtocoV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotCreateOrderListOtocoV3`: SpotCreateOrderListOtocoV3Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAPI.SpotCreateOrderListOtocoV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotCreateOrderListOtocoV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pendingAboveType** | **string** |  | [default to &quot;&quot;]
 **pendingQuantity** | **string** |  | [default to &quot;&quot;]
 **pendingSide** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **workingPrice** | **string** |  | [default to &quot;&quot;]
 **workingQuantity** | **string** |  | [default to &quot;&quot;]
 **workingSide** | **string** |  | [default to &quot;&quot;]
 **workingType** | **string** |  | [default to &quot;&quot;]
 **listClientOrderId** | **string** |  | [default to &quot;&quot;]
 **newOrderRespType** | **string** |  | [default to &quot;&quot;]
 **pendingAboveClientOrderId** | **string** |  | [default to &quot;&quot;]
 **pendingAboveIcebergQty** | **string** |  | [default to &quot;&quot;]
 **pendingAbovePrice** | **string** |  | [default to &quot;&quot;]
 **pendingAboveStopPrice** | **string** |  | [default to &quot;&quot;]
 **pendingAboveStrategyId** | **int64** |  | 
 **pendingAboveStrategyType** | **int32** |  | 
 **pendingAboveTimeInForce** | **string** |  | [default to &quot;&quot;]
 **pendingAboveTrailingDelta** | **string** |  | [default to &quot;&quot;]
 **pendingBelowClientOrderId** | **string** |  | [default to &quot;&quot;]
 **pendingBelowIcebergQty** | **string** |  | [default to &quot;&quot;]
 **pendingBelowPrice** | **string** |  | [default to &quot;&quot;]
 **pendingBelowStopPrice** | **string** |  | [default to &quot;&quot;]
 **pendingBelowStrategyId** | **int64** |  | 
 **pendingBelowStrategyType** | **int32** |  | 
 **pendingBelowTimeInForce** | **string** |  | [default to &quot;&quot;]
 **pendingBelowTrailingDelta** | **string** |  | [default to &quot;&quot;]
 **pendingBelowType** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **selfTradePreventionMode** | **string** |  | [default to &quot;&quot;]
 **workingClientOrderId** | **string** |  | [default to &quot;&quot;]
 **workingIcebergQty** | **string** |  | [default to &quot;&quot;]
 **workingStrategyId** | **int64** |  | 
 **workingStrategyType** | **int32** |  | 
 **workingTimeInForce** | **string** |  | [default to &quot;&quot;]

### Return type

[**SpotCreateOrderListOtocoV3Resp**](SpotCreateOrderListOtocoV3Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpotCreateOrderOcoV3

> SpotCreateOrderOcoV3Resp SpotCreateOrderOcoV3(ctx).Price(price).Quantity(quantity).Side(side).StopPrice(stopPrice).Symbol(symbol).Timestamp(timestamp).LimitClientOrderId(limitClientOrderId).LimitIcebergQty(limitIcebergQty).LimitStrategyId(limitStrategyId).LimitStrategyType(limitStrategyType).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).StopClientOrderId(stopClientOrderId).StopIcebergQty(stopIcebergQty).StopLimitPrice(stopLimitPrice).StopLimitTimeInForce(stopLimitTimeInForce).StopStrategyId(stopStrategyId).StopStrategyType(stopStrategyType).TrailingDelta(trailingDelta).Execute()

New OCO - Deprecated (TRADE)



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
	price := "price_example" // string |  (default to "")
	quantity := "quantity_example" // string |  (default to "")
	side := "side_example" // string |  (default to "")
	stopPrice := "stopPrice_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	limitClientOrderId := "limitClientOrderId_example" // string |  (optional) (default to "")
	limitIcebergQty := "limitIcebergQty_example" // string |  (optional) (default to "")
	limitStrategyId := int64(789) // int64 |  (optional)
	limitStrategyType := int32(56) // int32 |  (optional)
	listClientOrderId := "listClientOrderId_example" // string |  (optional) (default to "")
	newOrderRespType := "newOrderRespType_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	selfTradePreventionMode := "selfTradePreventionMode_example" // string |  (optional) (default to "")
	stopClientOrderId := "stopClientOrderId_example" // string |  (optional) (default to "")
	stopIcebergQty := "stopIcebergQty_example" // string |  (optional) (default to "")
	stopLimitPrice := "stopLimitPrice_example" // string |  (optional) (default to "")
	stopLimitTimeInForce := "stopLimitTimeInForce_example" // string |  (optional) (default to "")
	stopStrategyId := int64(789) // int64 |  (optional)
	stopStrategyType := int32(56) // int32 |  (optional)
	trailingDelta := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAPI.SpotCreateOrderOcoV3(context.Background()).Price(price).Quantity(quantity).Side(side).StopPrice(stopPrice).Symbol(symbol).Timestamp(timestamp).LimitClientOrderId(limitClientOrderId).LimitIcebergQty(limitIcebergQty).LimitStrategyId(limitStrategyId).LimitStrategyType(limitStrategyType).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).StopClientOrderId(stopClientOrderId).StopIcebergQty(stopIcebergQty).StopLimitPrice(stopLimitPrice).StopLimitTimeInForce(stopLimitTimeInForce).StopStrategyId(stopStrategyId).StopStrategyType(stopStrategyType).TrailingDelta(trailingDelta).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAPI.SpotCreateOrderOcoV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotCreateOrderOcoV3`: SpotCreateOrderOcoV3Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAPI.SpotCreateOrderOcoV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotCreateOrderOcoV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **price** | **string** |  | [default to &quot;&quot;]
 **quantity** | **string** |  | [default to &quot;&quot;]
 **side** | **string** |  | [default to &quot;&quot;]
 **stopPrice** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **limitClientOrderId** | **string** |  | [default to &quot;&quot;]
 **limitIcebergQty** | **string** |  | [default to &quot;&quot;]
 **limitStrategyId** | **int64** |  | 
 **limitStrategyType** | **int32** |  | 
 **listClientOrderId** | **string** |  | [default to &quot;&quot;]
 **newOrderRespType** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **selfTradePreventionMode** | **string** |  | [default to &quot;&quot;]
 **stopClientOrderId** | **string** |  | [default to &quot;&quot;]
 **stopIcebergQty** | **string** |  | [default to &quot;&quot;]
 **stopLimitPrice** | **string** |  | [default to &quot;&quot;]
 **stopLimitTimeInForce** | **string** |  | [default to &quot;&quot;]
 **stopStrategyId** | **int64** |  | 
 **stopStrategyType** | **int32** |  | 
 **trailingDelta** | **int64** |  | 

### Return type

[**SpotCreateOrderOcoV3Resp**](SpotCreateOrderOcoV3Resp.md)

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


## SpotCreateSorOrderV3

> SpotCreateSorOrderV3Resp SpotCreateSorOrderV3(ctx).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).IcebergQty(icebergQty).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).Price(price).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).StrategyId(strategyId).StrategyType(strategyType).TimeInForce(timeInForce).Execute()

New order using SOR (TRADE)



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
	resp, r, err := apiClient.TradingAPI.SpotCreateSorOrderV3(context.Background()).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).IcebergQty(icebergQty).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).Price(price).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).StrategyId(strategyId).StrategyType(strategyType).TimeInForce(timeInForce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAPI.SpotCreateSorOrderV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotCreateSorOrderV3`: SpotCreateSorOrderV3Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAPI.SpotCreateSorOrderV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotCreateSorOrderV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quantity** | **string** |  | [default to &quot;&quot;]
 **side** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **type_** | **string** |  | [default to &quot;&quot;]
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

[**SpotCreateSorOrderV3Resp**](SpotCreateSorOrderV3Resp.md)

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


## SpotDeleteOrderListV3

> SpotDeleteOrderListV3Resp SpotDeleteOrderListV3(ctx).Symbol(symbol).Timestamp(timestamp).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()

Cancel Order list (TRADE)



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
	orderListId := int64(789) // int64 | Either `orderListId` or `listClientOrderId` must be provided (optional)
	listClientOrderId := "listClientOrderId_example" // string | Either `orderListId` or `listClientOrderId` must be provided (optional) (default to "")
	newClientOrderId := "newClientOrderId_example" // string | Used to uniquely identify this cancel. Automatically generated by default (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAPI.SpotDeleteOrderListV3(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAPI.SpotDeleteOrderListV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotDeleteOrderListV3`: SpotDeleteOrderListV3Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAPI.SpotDeleteOrderListV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotDeleteOrderListV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderListId** | **int64** | Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided | 
 **listClientOrderId** | **string** | Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided | [default to &quot;&quot;]
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**SpotDeleteOrderListV3Resp**](SpotDeleteOrderListV3Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpotDeleteOrderV3

> SpotDeleteOrderV3Resp SpotDeleteOrderV3(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).CancelRestrictions(cancelRestrictions).RecvWindow(recvWindow).Execute()

Cancel order (TRADE)



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
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	newClientOrderId := "newClientOrderId_example" // string | Used to uniquely identify this cancel. Automatically generated by default. (optional) (default to "")
	cancelRestrictions := "cancelRestrictions_example" // string | Supported values: <br/>`ONLY_NEW` - Cancel will succeed if the order status is `NEW`.<br/> `ONLY_PARTIALLY_FILLED ` - Cancel will succeed if order status is `PARTIALLY_FILLED`. (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAPI.SpotDeleteOrderV3(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).CancelRestrictions(cancelRestrictions).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAPI.SpotDeleteOrderV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotDeleteOrderV3`: SpotDeleteOrderV3Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAPI.SpotDeleteOrderV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotDeleteOrderV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default. | [default to &quot;&quot;]
 **cancelRestrictions** | **string** | Supported values: &lt;br/&gt;&#x60;ONLY_NEW&#x60; - Cancel will succeed if the order status is &#x60;NEW&#x60;.&lt;br/&gt; &#x60;ONLY_PARTIALLY_FILLED &#x60; - Cancel will succeed if order status is &#x60;PARTIALLY_FILLED&#x60;. | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60;. | 

### Return type

[**SpotDeleteOrderV3Resp**](SpotDeleteOrderV3Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpotGetAllOrderListV3

> []SpotGetAllOrderListV3RespItem SpotGetAllOrderListV3(ctx).Timestamp(timestamp).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query all Order lists (USER_DATA)



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
	fromId := int64(789) // int64 | If supplied, neither `startTime` or `endTime` can be provided (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default Value: 500; Max Value: 1000 (optional)
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAPI.SpotGetAllOrderListV3(context.Background()).Timestamp(timestamp).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAPI.SpotGetAllOrderListV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotGetAllOrderListV3`: []SpotGetAllOrderListV3RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradingAPI.SpotGetAllOrderListV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotGetAllOrderListV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **fromId** | **int64** | If supplied, neither &#x60;startTime&#x60; or &#x60;endTime&#x60; can be provided | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default Value: 500; Max Value: 1000 | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]SpotGetAllOrderListV3RespItem**](SpotGetAllOrderListV3RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpotGetAllOrdersV3

> []SpotGetAllOrdersV3RespItem SpotGetAllOrdersV3(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

All orders (USER_DATA)



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
	orderId := int64(789) // int64 |  (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 500; max 1000. (optional) (default to 500)
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAPI.SpotGetAllOrdersV3(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAPI.SpotGetAllOrdersV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotGetAllOrdersV3`: []SpotGetAllOrdersV3RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradingAPI.SpotGetAllOrdersV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotGetAllOrdersV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 1000. | [default to 500]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]SpotGetAllOrdersV3RespItem**](SpotGetAllOrdersV3RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpotGetOpenOrderListV3

> []SpotGetOpenOrderListV3RespItem SpotGetOpenOrderListV3(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Open Order lists (USER_DATA)

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
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAPI.SpotGetOpenOrderListV3(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAPI.SpotGetOpenOrderListV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotGetOpenOrderListV3`: []SpotGetOpenOrderListV3RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradingAPI.SpotGetOpenOrderListV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotGetOpenOrderListV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]SpotGetOpenOrderListV3RespItem**](SpotGetOpenOrderListV3RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpotGetOpenOrdersV3

> []SpotGetOpenOrdersV3RespItem SpotGetOpenOrdersV3(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Current open orders (USER_DATA)



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
	symbol := "symbol_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAPI.SpotGetOpenOrdersV3(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAPI.SpotGetOpenOrdersV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotGetOpenOrdersV3`: []SpotGetOpenOrdersV3RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradingAPI.SpotGetOpenOrdersV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotGetOpenOrdersV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]SpotGetOpenOrdersV3RespItem**](SpotGetOpenOrdersV3RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpotGetOrderListV3

> SpotGetOrderListV3Resp SpotGetOrderListV3(ctx).Timestamp(timestamp).OrderListId(orderListId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Order list (USER_DATA)



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
	orderListId := int64(789) // int64 | Either `orderListId` or `listClientOrderId` must be provided (optional)
	origClientOrderId := "origClientOrderId_example" // string | Either `orderListId` or `listClientOrderId` must be provided (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAPI.SpotGetOrderListV3(context.Background()).Timestamp(timestamp).OrderListId(orderListId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAPI.SpotGetOrderListV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotGetOrderListV3`: SpotGetOrderListV3Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAPI.SpotGetOrderListV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotGetOrderListV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **orderListId** | **int64** | Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided | 
 **origClientOrderId** | **string** | Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**SpotGetOrderListV3Resp**](SpotGetOrderListV3Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpotGetOrderV3

> SpotGetOrderV3Resp SpotGetOrderV3(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query order (USER_DATA)



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
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAPI.SpotGetOrderV3(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAPI.SpotGetOrderV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotGetOrderV3`: SpotGetOrderV3Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAPI.SpotGetOrderV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotGetOrderV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**SpotGetOrderV3Resp**](SpotGetOrderV3Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

