# \SpotAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrderListOcoV3**](SpotAPI.md#CreateOrderListOcoV3) | **Post** /api/v3/orderList/oco | New Order list - OCO (TRADE)
[**CreateOrderListOtoV3**](SpotAPI.md#CreateOrderListOtoV3) | **Post** /api/v3/orderList/oto | New Order list - OTO (TRADE)
[**CreateOrderListOtocoV3**](SpotAPI.md#CreateOrderListOtocoV3) | **Post** /api/v3/orderList/otoco | New Order list - OTOCO (TRADE)
[**CreateOrderOcoV3**](SpotAPI.md#CreateOrderOcoV3) | **Post** /api/v3/order/oco | New OCO - Deprecated (TRADE)
[**CreateSorOrderV3**](SpotAPI.md#CreateSorOrderV3) | **Post** /api/v3/sor/order | New order using SOR (TRADE)
[**CreateUserDataStreamV3**](SpotAPI.md#CreateUserDataStreamV3) | **Post** /api/v3/userDataStream | Start user data stream (USER_STREAM)
[**DeleteOrderListV3**](SpotAPI.md#DeleteOrderListV3) | **Delete** /api/v3/orderList | Cancel Order list (TRADE)
[**DeleteOrderV3**](SpotAPI.md#DeleteOrderV3) | **Delete** /api/v3/order | Cancel order (TRADE)
[**DeleteUserDataStreamV3**](SpotAPI.md#DeleteUserDataStreamV3) | **Delete** /api/v3/userDataStream | Close user data stream (USER_STREAM)
[**GetAccountCommissionV3**](SpotAPI.md#GetAccountCommissionV3) | **Get** /api/v3/account/commission | Query Commission Rates (USER_DATA)
[**GetAccountV3**](SpotAPI.md#GetAccountV3) | **Get** /api/v3/account | Account information (USER_DATA)
[**GetAllOrderListV3**](SpotAPI.md#GetAllOrderListV3) | **Get** /api/v3/allOrderList | Query all Order lists (USER_DATA)
[**GetAllOrdersV3**](SpotAPI.md#GetAllOrdersV3) | **Get** /api/v3/allOrders | All orders (USER_DATA)
[**GetAvgPriceV3**](SpotAPI.md#GetAvgPriceV3) | **Get** /api/v3/avgPrice | Current average price
[**GetDepthV3**](SpotAPI.md#GetDepthV3) | **Get** /api/v3/depth | Order book
[**GetHistoricalTradesV3**](SpotAPI.md#GetHistoricalTradesV3) | **Get** /api/v3/historicalTrades | Old trade lookup
[**GetMyAllocationsV3**](SpotAPI.md#GetMyAllocationsV3) | **Get** /api/v3/myAllocations | Query Allocations (USER_DATA)
[**GetMyPreventedMatchesV3**](SpotAPI.md#GetMyPreventedMatchesV3) | **Get** /api/v3/myPreventedMatches | Query Prevented Matches (USER_DATA)
[**GetMyTradesV3**](SpotAPI.md#GetMyTradesV3) | **Get** /api/v3/myTrades | Account trade list (USER_DATA)
[**GetOpenOrderListV3**](SpotAPI.md#GetOpenOrderListV3) | **Get** /api/v3/openOrderList | Query Open Order lists (USER_DATA)
[**GetOpenOrdersV3**](SpotAPI.md#GetOpenOrdersV3) | **Get** /api/v3/openOrders | Current open orders (USER_DATA)
[**GetOrderListV3**](SpotAPI.md#GetOrderListV3) | **Get** /api/v3/orderList | Query Order list (USER_DATA)
[**GetOrderV3**](SpotAPI.md#GetOrderV3) | **Get** /api/v3/order | Query order (USER_DATA)
[**GetPingV3**](SpotAPI.md#GetPingV3) | **Get** /api/v3/ping | Test connectivity
[**GetRateLimitOrderV3**](SpotAPI.md#GetRateLimitOrderV3) | **Get** /api/v3/rateLimit/order | Query Unfilled Order Count (USER_DATA)
[**GetTimeV3**](SpotAPI.md#GetTimeV3) | **Get** /api/v3/time | Check server time
[**GetTradesV3**](SpotAPI.md#GetTradesV3) | **Get** /api/v3/trades | Recent trades list
[**UpdateUserDataStreamV3**](SpotAPI.md#UpdateUserDataStreamV3) | **Put** /api/v3/userDataStream | Keepalive user data stream (USER_STREAM)



## CreateOrderListOcoV3

> CreateOrderListOcoV3Resp CreateOrderListOcoV3(ctx).AboveType(aboveType).BelowType(belowType).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).AboveClientOrderId(aboveClientOrderId).AboveIcebergQty(aboveIcebergQty).AbovePrice(abovePrice).AboveStopPrice(aboveStopPrice).AboveStrategyId(aboveStrategyId).AboveStrategyType(aboveStrategyType).AboveTimeInForce(aboveTimeInForce).AboveTrailingDelta(aboveTrailingDelta).BelowClientOrderId(belowClientOrderId).BelowIcebergQty(belowIcebergQty).BelowPrice(belowPrice).BelowStopPrice(belowStopPrice).BelowStrategyId(belowStrategyId).BelowStrategyType(belowStrategyType).BelowTimeInForce(belowTimeInForce).BelowTrailingDelta(belowTrailingDelta).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).Execute()

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
	resp, r, err := apiClient.SpotAPI.CreateOrderListOcoV3(context.Background()).AboveType(aboveType).BelowType(belowType).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).AboveClientOrderId(aboveClientOrderId).AboveIcebergQty(aboveIcebergQty).AbovePrice(abovePrice).AboveStopPrice(aboveStopPrice).AboveStrategyId(aboveStrategyId).AboveStrategyType(aboveStrategyType).AboveTimeInForce(aboveTimeInForce).AboveTrailingDelta(aboveTrailingDelta).BelowClientOrderId(belowClientOrderId).BelowIcebergQty(belowIcebergQty).BelowPrice(belowPrice).BelowStopPrice(belowStopPrice).BelowStrategyId(belowStrategyId).BelowStrategyType(belowStrategyType).BelowTimeInForce(belowTimeInForce).BelowTrailingDelta(belowTrailingDelta).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotAPI.CreateOrderListOcoV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrderListOcoV3`: CreateOrderListOcoV3Resp
	fmt.Fprintf(os.Stdout, "Response from `SpotAPI.CreateOrderListOcoV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderListOcoV3Request struct via the builder pattern


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

[**CreateOrderListOcoV3Resp**](CreateOrderListOcoV3Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrderListOtoV3

> CreateOrderListOtoV3Resp CreateOrderListOtoV3(ctx).PendingQuantity(pendingQuantity).PendingSide(pendingSide).PendingType(pendingType).Symbol(symbol).Timestamp(timestamp).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).WorkingSide(workingSide).WorkingType(workingType).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).PendingClientOrderId(pendingClientOrderId).PendingIcebergQty(pendingIcebergQty).PendingPrice(pendingPrice).PendingStopPrice(pendingStopPrice).PendingStrategyId(pendingStrategyId).PendingStrategyType(pendingStrategyType).PendingTimeInForce(pendingTimeInForce).PendingTrailingDelta(pendingTrailingDelta).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingStrategyId(workingStrategyId).WorkingStrategyType(workingStrategyType).WorkingTimeInForce(workingTimeInForce).Execute()

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
	resp, r, err := apiClient.SpotAPI.CreateOrderListOtoV3(context.Background()).PendingQuantity(pendingQuantity).PendingSide(pendingSide).PendingType(pendingType).Symbol(symbol).Timestamp(timestamp).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).WorkingSide(workingSide).WorkingType(workingType).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).PendingClientOrderId(pendingClientOrderId).PendingIcebergQty(pendingIcebergQty).PendingPrice(pendingPrice).PendingStopPrice(pendingStopPrice).PendingStrategyId(pendingStrategyId).PendingStrategyType(pendingStrategyType).PendingTimeInForce(pendingTimeInForce).PendingTrailingDelta(pendingTrailingDelta).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingStrategyId(workingStrategyId).WorkingStrategyType(workingStrategyType).WorkingTimeInForce(workingTimeInForce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotAPI.CreateOrderListOtoV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrderListOtoV3`: CreateOrderListOtoV3Resp
	fmt.Fprintf(os.Stdout, "Response from `SpotAPI.CreateOrderListOtoV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderListOtoV3Request struct via the builder pattern


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

[**CreateOrderListOtoV3Resp**](CreateOrderListOtoV3Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrderListOtocoV3

> CreateOrderListOtocoV3Resp CreateOrderListOtocoV3(ctx).PendingAboveType(pendingAboveType).PendingQuantity(pendingQuantity).PendingSide(pendingSide).Symbol(symbol).Timestamp(timestamp).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).WorkingSide(workingSide).WorkingType(workingType).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).PendingAboveClientOrderId(pendingAboveClientOrderId).PendingAboveIcebergQty(pendingAboveIcebergQty).PendingAbovePrice(pendingAbovePrice).PendingAboveStopPrice(pendingAboveStopPrice).PendingAboveStrategyId(pendingAboveStrategyId).PendingAboveStrategyType(pendingAboveStrategyType).PendingAboveTimeInForce(pendingAboveTimeInForce).PendingAboveTrailingDelta(pendingAboveTrailingDelta).PendingBelowClientOrderId(pendingBelowClientOrderId).PendingBelowIcebergQty(pendingBelowIcebergQty).PendingBelowPrice(pendingBelowPrice).PendingBelowStopPrice(pendingBelowStopPrice).PendingBelowStrategyId(pendingBelowStrategyId).PendingBelowStrategyType(pendingBelowStrategyType).PendingBelowTimeInForce(pendingBelowTimeInForce).PendingBelowTrailingDelta(pendingBelowTrailingDelta).PendingBelowType(pendingBelowType).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingStrategyId(workingStrategyId).WorkingStrategyType(workingStrategyType).WorkingTimeInForce(workingTimeInForce).Execute()

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
	resp, r, err := apiClient.SpotAPI.CreateOrderListOtocoV3(context.Background()).PendingAboveType(pendingAboveType).PendingQuantity(pendingQuantity).PendingSide(pendingSide).Symbol(symbol).Timestamp(timestamp).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).WorkingSide(workingSide).WorkingType(workingType).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).PendingAboveClientOrderId(pendingAboveClientOrderId).PendingAboveIcebergQty(pendingAboveIcebergQty).PendingAbovePrice(pendingAbovePrice).PendingAboveStopPrice(pendingAboveStopPrice).PendingAboveStrategyId(pendingAboveStrategyId).PendingAboveStrategyType(pendingAboveStrategyType).PendingAboveTimeInForce(pendingAboveTimeInForce).PendingAboveTrailingDelta(pendingAboveTrailingDelta).PendingBelowClientOrderId(pendingBelowClientOrderId).PendingBelowIcebergQty(pendingBelowIcebergQty).PendingBelowPrice(pendingBelowPrice).PendingBelowStopPrice(pendingBelowStopPrice).PendingBelowStrategyId(pendingBelowStrategyId).PendingBelowStrategyType(pendingBelowStrategyType).PendingBelowTimeInForce(pendingBelowTimeInForce).PendingBelowTrailingDelta(pendingBelowTrailingDelta).PendingBelowType(pendingBelowType).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingStrategyId(workingStrategyId).WorkingStrategyType(workingStrategyType).WorkingTimeInForce(workingTimeInForce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotAPI.CreateOrderListOtocoV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrderListOtocoV3`: CreateOrderListOtocoV3Resp
	fmt.Fprintf(os.Stdout, "Response from `SpotAPI.CreateOrderListOtocoV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderListOtocoV3Request struct via the builder pattern


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

[**CreateOrderListOtocoV3Resp**](CreateOrderListOtocoV3Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrderOcoV3

> CreateOrderOcoV3Resp CreateOrderOcoV3(ctx).Price(price).Quantity(quantity).Side(side).StopPrice(stopPrice).Symbol(symbol).Timestamp(timestamp).LimitClientOrderId(limitClientOrderId).LimitIcebergQty(limitIcebergQty).LimitStrategyId(limitStrategyId).LimitStrategyType(limitStrategyType).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).StopClientOrderId(stopClientOrderId).StopIcebergQty(stopIcebergQty).StopLimitPrice(stopLimitPrice).StopLimitTimeInForce(stopLimitTimeInForce).StopStrategyId(stopStrategyId).StopStrategyType(stopStrategyType).TrailingDelta(trailingDelta).Execute()

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
	resp, r, err := apiClient.SpotAPI.CreateOrderOcoV3(context.Background()).Price(price).Quantity(quantity).Side(side).StopPrice(stopPrice).Symbol(symbol).Timestamp(timestamp).LimitClientOrderId(limitClientOrderId).LimitIcebergQty(limitIcebergQty).LimitStrategyId(limitStrategyId).LimitStrategyType(limitStrategyType).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).StopClientOrderId(stopClientOrderId).StopIcebergQty(stopIcebergQty).StopLimitPrice(stopLimitPrice).StopLimitTimeInForce(stopLimitTimeInForce).StopStrategyId(stopStrategyId).StopStrategyType(stopStrategyType).TrailingDelta(trailingDelta).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotAPI.CreateOrderOcoV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrderOcoV3`: CreateOrderOcoV3Resp
	fmt.Fprintf(os.Stdout, "Response from `SpotAPI.CreateOrderOcoV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderOcoV3Request struct via the builder pattern


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

[**CreateOrderOcoV3Resp**](CreateOrderOcoV3Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSorOrderV3

> CreateSorOrderV3Resp CreateSorOrderV3(ctx).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).IcebergQty(icebergQty).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).Price(price).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).StrategyId(strategyId).StrategyType(strategyType).TimeInForce(timeInForce).Execute()

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
	resp, r, err := apiClient.SpotAPI.CreateSorOrderV3(context.Background()).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).IcebergQty(icebergQty).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).Price(price).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).StrategyId(strategyId).StrategyType(strategyType).TimeInForce(timeInForce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotAPI.CreateSorOrderV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSorOrderV3`: CreateSorOrderV3Resp
	fmt.Fprintf(os.Stdout, "Response from `SpotAPI.CreateSorOrderV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSorOrderV3Request struct via the builder pattern


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

[**CreateSorOrderV3Resp**](CreateSorOrderV3Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserDataStreamV3

> CreateUserDataStreamV3Resp CreateUserDataStreamV3(ctx).Execute()

Start user data stream (USER_STREAM)



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpotAPI.CreateUserDataStreamV3(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotAPI.CreateUserDataStreamV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUserDataStreamV3`: CreateUserDataStreamV3Resp
	fmt.Fprintf(os.Stdout, "Response from `SpotAPI.CreateUserDataStreamV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserDataStreamV3Request struct via the builder pattern


### Return type

[**CreateUserDataStreamV3Resp**](CreateUserDataStreamV3Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrderListV3

> DeleteOrderListV3Resp DeleteOrderListV3(ctx).Symbol(symbol).Timestamp(timestamp).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()

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
	resp, r, err := apiClient.SpotAPI.DeleteOrderListV3(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotAPI.DeleteOrderListV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrderListV3`: DeleteOrderListV3Resp
	fmt.Fprintf(os.Stdout, "Response from `SpotAPI.DeleteOrderListV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrderListV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderListId** | **int64** | Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided | 
 **listClientOrderId** | **string** | Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided | [default to &quot;&quot;]
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**DeleteOrderListV3Resp**](DeleteOrderListV3Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrderV3

> DeleteOrderV3Resp DeleteOrderV3(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).CancelRestrictions(cancelRestrictions).RecvWindow(recvWindow).Execute()

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
	resp, r, err := apiClient.SpotAPI.DeleteOrderV3(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).CancelRestrictions(cancelRestrictions).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotAPI.DeleteOrderV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrderV3`: DeleteOrderV3Resp
	fmt.Fprintf(os.Stdout, "Response from `SpotAPI.DeleteOrderV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrderV3Request struct via the builder pattern


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

[**DeleteOrderV3Resp**](DeleteOrderV3Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserDataStreamV3

> map[string]interface{} DeleteUserDataStreamV3(ctx).ListenKey(listenKey).Execute()

Close user data stream (USER_STREAM)



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
	listenKey := "listenKey_example" // string |  (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpotAPI.DeleteUserDataStreamV3(context.Background()).ListenKey(listenKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotAPI.DeleteUserDataStreamV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUserDataStreamV3`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SpotAPI.DeleteUserDataStreamV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserDataStreamV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listenKey** | **string** |  | [default to &quot;&quot;]

### Return type

**map[string]interface{}**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountCommissionV3

> GetAccountCommissionV3Resp GetAccountCommissionV3(ctx).Symbol(symbol).Execute()

Query Commission Rates (USER_DATA)



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpotAPI.GetAccountCommissionV3(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotAPI.GetAccountCommissionV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountCommissionV3`: GetAccountCommissionV3Resp
	fmt.Fprintf(os.Stdout, "Response from `SpotAPI.GetAccountCommissionV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountCommissionV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**GetAccountCommissionV3Resp**](GetAccountCommissionV3Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountV3

> GetAccountV3Resp GetAccountV3(ctx).Timestamp(timestamp).OmitZeroBalances(omitZeroBalances).RecvWindow(recvWindow).Execute()

Account information (USER_DATA)



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
	omitZeroBalances := true // bool | When set to `true`, emits only the non-zero balances of an account. <br/>Default value: `false` (optional)
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpotAPI.GetAccountV3(context.Background()).Timestamp(timestamp).OmitZeroBalances(omitZeroBalances).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotAPI.GetAccountV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountV3`: GetAccountV3Resp
	fmt.Fprintf(os.Stdout, "Response from `SpotAPI.GetAccountV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **omitZeroBalances** | **bool** | When set to &#x60;true&#x60;, emits only the non-zero balances of an account. &lt;br/&gt;Default value: &#x60;false&#x60; | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**GetAccountV3Resp**](GetAccountV3Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllOrderListV3

> []GetAllOrderListV3RespItem GetAllOrderListV3(ctx).Timestamp(timestamp).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

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
	resp, r, err := apiClient.SpotAPI.GetAllOrderListV3(context.Background()).Timestamp(timestamp).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotAPI.GetAllOrderListV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllOrderListV3`: []GetAllOrderListV3RespItem
	fmt.Fprintf(os.Stdout, "Response from `SpotAPI.GetAllOrderListV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllOrderListV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **fromId** | **int64** | If supplied, neither &#x60;startTime&#x60; or &#x60;endTime&#x60; can be provided | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default Value: 500; Max Value: 1000 | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]GetAllOrderListV3RespItem**](GetAllOrderListV3RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllOrdersV3

> []GetAllOrdersV3RespItem GetAllOrdersV3(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

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
	resp, r, err := apiClient.SpotAPI.GetAllOrdersV3(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotAPI.GetAllOrdersV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllOrdersV3`: []GetAllOrdersV3RespItem
	fmt.Fprintf(os.Stdout, "Response from `SpotAPI.GetAllOrdersV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllOrdersV3Request struct via the builder pattern


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

[**[]GetAllOrdersV3RespItem**](GetAllOrdersV3RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvgPriceV3

> GetAvgPriceV3Resp GetAvgPriceV3(ctx).Symbol(symbol).Execute()

Current average price



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpotAPI.GetAvgPriceV3(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotAPI.GetAvgPriceV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvgPriceV3`: GetAvgPriceV3Resp
	fmt.Fprintf(os.Stdout, "Response from `SpotAPI.GetAvgPriceV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAvgPriceV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**GetAvgPriceV3Resp**](GetAvgPriceV3Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDepthV3

> GetDepthV3Resp GetDepthV3(ctx).Symbol(symbol).Limit(limit).Execute()

Order book

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
	limit := int32(56) // int32 | Default 100; max 5000. <br/> If limit &gt; 5000. then the response will truncate to 5000. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpotAPI.GetDepthV3(context.Background()).Symbol(symbol).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotAPI.GetDepthV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDepthV3`: GetDepthV3Resp
	fmt.Fprintf(os.Stdout, "Response from `SpotAPI.GetDepthV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDepthV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **limit** | **int32** | Default 100; max 5000. &lt;br/&gt; If limit &amp;gt; 5000. then the response will truncate to 5000. | [default to 100]

### Return type

[**GetDepthV3Resp**](GetDepthV3Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoricalTradesV3

> []GetHistoricalTradesV3RespItem GetHistoricalTradesV3(ctx).Symbol(symbol).Limit(limit).FromId(fromId).Execute()

Old trade lookup



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
	limit := int32(56) // int32 | Default 500; max 1000. (optional) (default to 500)
	fromId := int64(789) // int64 | TradeId to fetch from. Default gets most recent trades. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpotAPI.GetHistoricalTradesV3(context.Background()).Symbol(symbol).Limit(limit).FromId(fromId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotAPI.GetHistoricalTradesV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricalTradesV3`: []GetHistoricalTradesV3RespItem
	fmt.Fprintf(os.Stdout, "Response from `SpotAPI.GetHistoricalTradesV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricalTradesV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **limit** | **int32** | Default 500; max 1000. | [default to 500]
 **fromId** | **int64** | TradeId to fetch from. Default gets most recent trades. | 

### Return type

[**[]GetHistoricalTradesV3RespItem**](GetHistoricalTradesV3RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyAllocationsV3

> []GetMyAllocationsV3RespItem GetMyAllocationsV3(ctx).Symbol(symbol).StartTime(startTime).EndTime(endTime).FromAllocationId(fromAllocationId).Limit(limit).OrderId(orderId).RecvWindow(recvWindow).Timestamp(timestamp).Execute()

Query Allocations (USER_DATA)



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
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	fromAllocationId := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 | Default 500;Max 1000 (optional) (default to 500)
	orderId := int64(789) // int64 |  (optional)
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000`. (optional)
	timestamp := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpotAPI.GetMyAllocationsV3(context.Background()).Symbol(symbol).StartTime(startTime).EndTime(endTime).FromAllocationId(fromAllocationId).Limit(limit).OrderId(orderId).RecvWindow(recvWindow).Timestamp(timestamp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotAPI.GetMyAllocationsV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyAllocationsV3`: []GetMyAllocationsV3RespItem
	fmt.Fprintf(os.Stdout, "Response from `SpotAPI.GetMyAllocationsV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyAllocationsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **fromAllocationId** | **int32** |  | 
 **limit** | **int32** | Default 500;Max 1000 | [default to 500]
 **orderId** | **int64** |  | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60;. | 
 **timestamp** | **int64** |  | 

### Return type

[**[]GetMyAllocationsV3RespItem**](GetMyAllocationsV3RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyPreventedMatchesV3

> []GetMyPreventedMatchesV3RespItem GetMyPreventedMatchesV3(ctx).Symbol(symbol).Timestamp(timestamp).PreventedMatchId(preventedMatchId).OrderId(orderId).FromPreventedMatchId(fromPreventedMatchId).Limit(limit).RecvWindow(recvWindow).Execute()

Query Prevented Matches (USER_DATA)



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
	preventedMatchId := int64(789) // int64 |  (optional)
	orderId := int64(789) // int64 |  (optional)
	fromPreventedMatchId := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default: `500`; Max: `1000` (optional) (default to 500)
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpotAPI.GetMyPreventedMatchesV3(context.Background()).Symbol(symbol).Timestamp(timestamp).PreventedMatchId(preventedMatchId).OrderId(orderId).FromPreventedMatchId(fromPreventedMatchId).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotAPI.GetMyPreventedMatchesV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyPreventedMatchesV3`: []GetMyPreventedMatchesV3RespItem
	fmt.Fprintf(os.Stdout, "Response from `SpotAPI.GetMyPreventedMatchesV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyPreventedMatchesV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **preventedMatchId** | **int64** |  | 
 **orderId** | **int64** |  | 
 **fromPreventedMatchId** | **int64** |  | 
 **limit** | **int32** | Default: &#x60;500&#x60;; Max: &#x60;1000&#x60; | [default to 500]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]GetMyPreventedMatchesV3RespItem**](GetMyPreventedMatchesV3RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyTradesV3

> []GetMyTradesV3RespItem GetMyTradesV3(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()

Account trade list (USER_DATA)



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
	orderId := int64(789) // int64 | This can only be used in combination with `symbol`. (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	fromId := int64(789) // int64 | TradeId to fetch from. Default gets most recent trades. (optional)
	limit := int32(56) // int32 | Default 500; max 1000. (optional) (default to 500)
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpotAPI.GetMyTradesV3(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotAPI.GetMyTradesV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyTradesV3`: []GetMyTradesV3RespItem
	fmt.Fprintf(os.Stdout, "Response from `SpotAPI.GetMyTradesV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyTradesV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** | This can only be used in combination with &#x60;symbol&#x60;. | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **fromId** | **int64** | TradeId to fetch from. Default gets most recent trades. | 
 **limit** | **int32** | Default 500; max 1000. | [default to 500]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]GetMyTradesV3RespItem**](GetMyTradesV3RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOpenOrderListV3

> []GetOpenOrderListV3RespItem GetOpenOrderListV3(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

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
	resp, r, err := apiClient.SpotAPI.GetOpenOrderListV3(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotAPI.GetOpenOrderListV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOpenOrderListV3`: []GetOpenOrderListV3RespItem
	fmt.Fprintf(os.Stdout, "Response from `SpotAPI.GetOpenOrderListV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOpenOrderListV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]GetOpenOrderListV3RespItem**](GetOpenOrderListV3RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOpenOrdersV3

> []GetOpenOrdersV3RespItem GetOpenOrdersV3(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

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
	resp, r, err := apiClient.SpotAPI.GetOpenOrdersV3(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotAPI.GetOpenOrdersV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOpenOrdersV3`: []GetOpenOrdersV3RespItem
	fmt.Fprintf(os.Stdout, "Response from `SpotAPI.GetOpenOrdersV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOpenOrdersV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]GetOpenOrdersV3RespItem**](GetOpenOrdersV3RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderListV3

> GetOrderListV3Resp GetOrderListV3(ctx).Timestamp(timestamp).OrderListId(orderListId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

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
	resp, r, err := apiClient.SpotAPI.GetOrderListV3(context.Background()).Timestamp(timestamp).OrderListId(orderListId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotAPI.GetOrderListV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrderListV3`: GetOrderListV3Resp
	fmt.Fprintf(os.Stdout, "Response from `SpotAPI.GetOrderListV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderListV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **orderListId** | **int64** | Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided | 
 **origClientOrderId** | **string** | Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**GetOrderListV3Resp**](GetOrderListV3Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderV3

> GetOrderV3Resp GetOrderV3(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

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
	resp, r, err := apiClient.SpotAPI.GetOrderV3(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotAPI.GetOrderV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrderV3`: GetOrderV3Resp
	fmt.Fprintf(os.Stdout, "Response from `SpotAPI.GetOrderV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**GetOrderV3Resp**](GetOrderV3Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPingV3

> map[string]interface{} GetPingV3(ctx).Execute()

Test connectivity



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpotAPI.GetPingV3(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotAPI.GetPingV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPingV3`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SpotAPI.GetPingV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPingV3Request struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRateLimitOrderV3

> []GetRateLimitOrderV3RespItem GetRateLimitOrderV3(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Unfilled Order Count (USER_DATA)



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
	resp, r, err := apiClient.SpotAPI.GetRateLimitOrderV3(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotAPI.GetRateLimitOrderV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRateLimitOrderV3`: []GetRateLimitOrderV3RespItem
	fmt.Fprintf(os.Stdout, "Response from `SpotAPI.GetRateLimitOrderV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRateLimitOrderV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]GetRateLimitOrderV3RespItem**](GetRateLimitOrderV3RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeV3

> GetTimeV3Resp GetTimeV3(ctx).Execute()

Check server time



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpotAPI.GetTimeV3(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotAPI.GetTimeV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeV3`: GetTimeV3Resp
	fmt.Fprintf(os.Stdout, "Response from `SpotAPI.GetTimeV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeV3Request struct via the builder pattern


### Return type

[**GetTimeV3Resp**](GetTimeV3Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradesV3

> []GetTradesV3RespItem GetTradesV3(ctx).Symbol(symbol).Limit(limit).Execute()

Recent trades list



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
	limit := int32(56) // int32 | Default 500; max 1000. (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpotAPI.GetTradesV3(context.Background()).Symbol(symbol).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotAPI.GetTradesV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradesV3`: []GetTradesV3RespItem
	fmt.Fprintf(os.Stdout, "Response from `SpotAPI.GetTradesV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradesV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **limit** | **int32** | Default 500; max 1000. | [default to 500]

### Return type

[**[]GetTradesV3RespItem**](GetTradesV3RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserDataStreamV3

> map[string]interface{} UpdateUserDataStreamV3(ctx).ListenKey(listenKey).Execute()

Keepalive user data stream (USER_STREAM)



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
	listenKey := "listenKey_example" // string |  (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpotAPI.UpdateUserDataStreamV3(context.Background()).ListenKey(listenKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpotAPI.UpdateUserDataStreamV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserDataStreamV3`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SpotAPI.UpdateUserDataStreamV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserDataStreamV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listenKey** | **string** |  | [default to &quot;&quot;]

### Return type

**map[string]interface{}**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

