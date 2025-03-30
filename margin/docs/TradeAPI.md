# \TradeAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarginCreateMarginApiKeyV1**](TradeAPI.md#MarginCreateMarginApiKeyV1) | **Post** /sapi/v1/margin/apiKey | Create Special Key(Low-Latency Trading)(TRADE)
[**MarginCreateMarginExchangeSmallLiabilityV1**](TradeAPI.md#MarginCreateMarginExchangeSmallLiabilityV1) | **Post** /sapi/v1/margin/exchange-small-liability | Small Liability Exchange (MARGIN)
[**MarginCreateMarginManualLiquidationV1**](TradeAPI.md#MarginCreateMarginManualLiquidationV1) | **Post** /sapi/v1/margin/manual-liquidation | Margin Manual Liquidation(MARGIN)
[**MarginCreateMarginOrderOcoV1**](TradeAPI.md#MarginCreateMarginOrderOcoV1) | **Post** /sapi/v1/margin/order/oco | Margin Account New OCO (TRADE)
[**MarginCreateMarginOrderOtoV1**](TradeAPI.md#MarginCreateMarginOrderOtoV1) | **Post** /sapi/v1/margin/order/oto | Margin Account New OTO (TRADE)
[**MarginCreateMarginOrderOtocoV1**](TradeAPI.md#MarginCreateMarginOrderOtocoV1) | **Post** /sapi/v1/margin/order/otoco | Margin Account New OTOCO (TRADE)
[**MarginCreateMarginOrderV1**](TradeAPI.md#MarginCreateMarginOrderV1) | **Post** /sapi/v1/margin/order | Margin Account New Order (TRADE)
[**MarginDeleteMarginApiKeyV1**](TradeAPI.md#MarginDeleteMarginApiKeyV1) | **Delete** /sapi/v1/margin/apiKey | Delete Special Key(Low-Latency Trading)(TRADE)
[**MarginDeleteMarginOpenOrdersV1**](TradeAPI.md#MarginDeleteMarginOpenOrdersV1) | **Delete** /sapi/v1/margin/openOrders | Margin Account Cancel all Open Orders on a Symbol (TRADE)
[**MarginDeleteMarginOrderListV1**](TradeAPI.md#MarginDeleteMarginOrderListV1) | **Delete** /sapi/v1/margin/orderList | Margin Account Cancel OCO (TRADE)
[**MarginDeleteMarginOrderV1**](TradeAPI.md#MarginDeleteMarginOrderV1) | **Delete** /sapi/v1/margin/order | Margin Account Cancel Order (TRADE)
[**MarginGetMarginAllOrderListV1**](TradeAPI.md#MarginGetMarginAllOrderListV1) | **Get** /sapi/v1/margin/allOrderList | Query Margin Account&#39;s all OCO (USER_DATA)
[**MarginGetMarginAllOrdersV1**](TradeAPI.md#MarginGetMarginAllOrdersV1) | **Get** /sapi/v1/margin/allOrders | Query Margin Account&#39;s All Orders (USER_DATA)
[**MarginGetMarginApiKeyListV1**](TradeAPI.md#MarginGetMarginApiKeyListV1) | **Get** /sapi/v1/margin/api-key-list | Query Special key List(Low Latency Trading)(TRADE)
[**MarginGetMarginApiKeyV1**](TradeAPI.md#MarginGetMarginApiKeyV1) | **Get** /sapi/v1/margin/apiKey | Query Special key(Low Latency Trading)(TRADE)
[**MarginGetMarginExchangeSmallLiabilityHistoryV1**](TradeAPI.md#MarginGetMarginExchangeSmallLiabilityHistoryV1) | **Get** /sapi/v1/margin/exchange-small-liability-history | Get Small Liability Exchange History (USER_DATA)
[**MarginGetMarginExchangeSmallLiabilityV1**](TradeAPI.md#MarginGetMarginExchangeSmallLiabilityV1) | **Get** /sapi/v1/margin/exchange-small-liability | Get Small Liability Exchange Coin List (USER_DATA)
[**MarginGetMarginForceLiquidationRecV1**](TradeAPI.md#MarginGetMarginForceLiquidationRecV1) | **Get** /sapi/v1/margin/forceLiquidationRec | Get Force Liquidation Record (USER_DATA)
[**MarginGetMarginMyTradesV1**](TradeAPI.md#MarginGetMarginMyTradesV1) | **Get** /sapi/v1/margin/myTrades | Query Margin Account&#39;s Trade List (USER_DATA)
[**MarginGetMarginOpenOrderListV1**](TradeAPI.md#MarginGetMarginOpenOrderListV1) | **Get** /sapi/v1/margin/openOrderList | Query Margin Account&#39;s Open OCO (USER_DATA)
[**MarginGetMarginOpenOrdersV1**](TradeAPI.md#MarginGetMarginOpenOrdersV1) | **Get** /sapi/v1/margin/openOrders | Query Margin Account&#39;s Open Orders (USER_DATA)
[**MarginGetMarginOrderListV1**](TradeAPI.md#MarginGetMarginOrderListV1) | **Get** /sapi/v1/margin/orderList | Query Margin Account&#39;s OCO (USER_DATA)
[**MarginGetMarginOrderV1**](TradeAPI.md#MarginGetMarginOrderV1) | **Get** /sapi/v1/margin/order | Query Margin Account&#39;s Order (USER_DATA)
[**MarginGetMarginRateLimitOrderV1**](TradeAPI.md#MarginGetMarginRateLimitOrderV1) | **Get** /sapi/v1/margin/rateLimit/order | Query Current Margin Order Count Usage (TRADE)
[**MarginUpdateMarginApiKeyIpV1**](TradeAPI.md#MarginUpdateMarginApiKeyIpV1) | **Put** /sapi/v1/margin/apiKey/ip | Edit ip for Special Key(Low-Latency Trading)(TRADE)



## MarginCreateMarginApiKeyV1

> MarginCreateMarginApiKeyV1Resp MarginCreateMarginApiKeyV1(ctx).ApiName(apiName).Timestamp(timestamp).Ip(ip).PermissionMode(permissionMode).PublicKey(publicKey).RecvWindow(recvWindow).Symbol(symbol).Execute()

Create Special Key(Low-Latency Trading)(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/margin"
)

func main() {
	apiName := "apiName_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	ip := "ip_example" // string |  (optional) (default to "")
	permissionMode := "permissionMode_example" // string |  (optional) (default to "")
	publicKey := "publicKey_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	symbol := "symbol_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.MarginCreateMarginApiKeyV1(context.Background()).ApiName(apiName).Timestamp(timestamp).Ip(ip).PermissionMode(permissionMode).PublicKey(publicKey).RecvWindow(recvWindow).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.MarginCreateMarginApiKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginCreateMarginApiKeyV1`: MarginCreateMarginApiKeyV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.MarginCreateMarginApiKeyV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginCreateMarginApiKeyV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiName** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **ip** | **string** |  | [default to &quot;&quot;]
 **permissionMode** | **string** |  | [default to &quot;&quot;]
 **publicKey** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**MarginCreateMarginApiKeyV1Resp**](MarginCreateMarginApiKeyV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginCreateMarginExchangeSmallLiabilityV1

> MarginCreateMarginExchangeSmallLiabilityV1Resp MarginCreateMarginExchangeSmallLiabilityV1(ctx).AssetNames(assetNames).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Small Liability Exchange (MARGIN)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/margin"
)

func main() {
	assetNames := []string{"Inner_example"} // []string | 
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.MarginCreateMarginExchangeSmallLiabilityV1(context.Background()).AssetNames(assetNames).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.MarginCreateMarginExchangeSmallLiabilityV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginCreateMarginExchangeSmallLiabilityV1`: MarginCreateMarginExchangeSmallLiabilityV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.MarginCreateMarginExchangeSmallLiabilityV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginCreateMarginExchangeSmallLiabilityV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetNames** | **[]string** |  | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**MarginCreateMarginExchangeSmallLiabilityV1Resp**](MarginCreateMarginExchangeSmallLiabilityV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginCreateMarginManualLiquidationV1

> MarginCreateMarginManualLiquidationV1Resp MarginCreateMarginManualLiquidationV1(ctx).Timestamp(timestamp).Type_(type_).RecvWindow(recvWindow).Symbol(symbol).Execute()

Margin Manual Liquidation(MARGIN)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/margin"
)

func main() {
	timestamp := int64(789) // int64 | 
	type_ := "type__example" // string |  (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	symbol := "symbol_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.MarginCreateMarginManualLiquidationV1(context.Background()).Timestamp(timestamp).Type_(type_).RecvWindow(recvWindow).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.MarginCreateMarginManualLiquidationV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginCreateMarginManualLiquidationV1`: MarginCreateMarginManualLiquidationV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.MarginCreateMarginManualLiquidationV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginCreateMarginManualLiquidationV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **type_** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**MarginCreateMarginManualLiquidationV1Resp**](MarginCreateMarginManualLiquidationV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginCreateMarginOrderOcoV1

> MarginCreateMarginOrderOcoV1Resp MarginCreateMarginOrderOcoV1(ctx).Price(price).Quantity(quantity).Side(side).StopPrice(stopPrice).Symbol(symbol).Timestamp(timestamp).AutoRepayAtCancel(autoRepayAtCancel).IsIsolated(isIsolated).LimitClientOrderId(limitClientOrderId).LimitIcebergQty(limitIcebergQty).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).SideEffectType(sideEffectType).StopClientOrderId(stopClientOrderId).StopIcebergQty(stopIcebergQty).StopLimitPrice(stopLimitPrice).StopLimitTimeInForce(stopLimitTimeInForce).Execute()

Margin Account New OCO (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/margin"
)

func main() {
	price := "price_example" // string |  (default to "")
	quantity := "quantity_example" // string |  (default to "")
	side := "side_example" // string |  (default to "")
	stopPrice := "stopPrice_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	autoRepayAtCancel := true // bool |  (optional)
	isIsolated := "isIsolated_example" // string |  (optional) (default to "")
	limitClientOrderId := "limitClientOrderId_example" // string |  (optional) (default to "")
	limitIcebergQty := "limitIcebergQty_example" // string |  (optional) (default to "")
	listClientOrderId := "listClientOrderId_example" // string |  (optional) (default to "")
	newOrderRespType := "newOrderRespType_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	selfTradePreventionMode := "selfTradePreventionMode_example" // string |  (optional) (default to "")
	sideEffectType := "sideEffectType_example" // string |  (optional) (default to "")
	stopClientOrderId := "stopClientOrderId_example" // string |  (optional) (default to "")
	stopIcebergQty := "stopIcebergQty_example" // string |  (optional) (default to "")
	stopLimitPrice := "stopLimitPrice_example" // string |  (optional) (default to "")
	stopLimitTimeInForce := "stopLimitTimeInForce_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.MarginCreateMarginOrderOcoV1(context.Background()).Price(price).Quantity(quantity).Side(side).StopPrice(stopPrice).Symbol(symbol).Timestamp(timestamp).AutoRepayAtCancel(autoRepayAtCancel).IsIsolated(isIsolated).LimitClientOrderId(limitClientOrderId).LimitIcebergQty(limitIcebergQty).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).SideEffectType(sideEffectType).StopClientOrderId(stopClientOrderId).StopIcebergQty(stopIcebergQty).StopLimitPrice(stopLimitPrice).StopLimitTimeInForce(stopLimitTimeInForce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.MarginCreateMarginOrderOcoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginCreateMarginOrderOcoV1`: MarginCreateMarginOrderOcoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.MarginCreateMarginOrderOcoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginCreateMarginOrderOcoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **price** | **string** |  | [default to &quot;&quot;]
 **quantity** | **string** |  | [default to &quot;&quot;]
 **side** | **string** |  | [default to &quot;&quot;]
 **stopPrice** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **autoRepayAtCancel** | **bool** |  | 
 **isIsolated** | **string** |  | [default to &quot;&quot;]
 **limitClientOrderId** | **string** |  | [default to &quot;&quot;]
 **limitIcebergQty** | **string** |  | [default to &quot;&quot;]
 **listClientOrderId** | **string** |  | [default to &quot;&quot;]
 **newOrderRespType** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **selfTradePreventionMode** | **string** |  | [default to &quot;&quot;]
 **sideEffectType** | **string** |  | [default to &quot;&quot;]
 **stopClientOrderId** | **string** |  | [default to &quot;&quot;]
 **stopIcebergQty** | **string** |  | [default to &quot;&quot;]
 **stopLimitPrice** | **string** |  | [default to &quot;&quot;]
 **stopLimitTimeInForce** | **string** |  | [default to &quot;&quot;]

### Return type

[**MarginCreateMarginOrderOcoV1Resp**](MarginCreateMarginOrderOcoV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginCreateMarginOrderOtoV1

> MarginCreateMarginOrderOtoV1Resp MarginCreateMarginOrderOtoV1(ctx).PendingQuantity(pendingQuantity).PendingSide(pendingSide).PendingType(pendingType).Symbol(symbol).WorkingIcebergQty(workingIcebergQty).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).WorkingSide(workingSide).WorkingType(workingType).AutoRepayAtCancel(autoRepayAtCancel).IsIsolated(isIsolated).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).PendingClientOrderId(pendingClientOrderId).PendingIcebergQty(pendingIcebergQty).PendingPrice(pendingPrice).PendingStopPrice(pendingStopPrice).PendingTimeInForce(pendingTimeInForce).PendingTrailingDelta(pendingTrailingDelta).SelfTradePreventionMode(selfTradePreventionMode).SideEffectType(sideEffectType).WorkingClientOrderId(workingClientOrderId).WorkingTimeInForce(workingTimeInForce).Execute()

Margin Account New OTO (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/margin"
)

func main() {
	pendingQuantity := "pendingQuantity_example" // string |  (default to "")
	pendingSide := "pendingSide_example" // string |  (default to "")
	pendingType := "pendingType_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	workingIcebergQty := "workingIcebergQty_example" // string |  (default to "")
	workingPrice := "workingPrice_example" // string |  (default to "")
	workingQuantity := "workingQuantity_example" // string |  (default to "")
	workingSide := "workingSide_example" // string |  (default to "")
	workingType := "workingType_example" // string |  (default to "")
	autoRepayAtCancel := true // bool |  (optional)
	isIsolated := "isIsolated_example" // string |  (optional) (default to "")
	listClientOrderId := "listClientOrderId_example" // string |  (optional) (default to "")
	newOrderRespType := "newOrderRespType_example" // string |  (optional) (default to "")
	pendingClientOrderId := "pendingClientOrderId_example" // string |  (optional) (default to "")
	pendingIcebergQty := "pendingIcebergQty_example" // string |  (optional) (default to "")
	pendingPrice := "pendingPrice_example" // string |  (optional) (default to "")
	pendingStopPrice := "pendingStopPrice_example" // string |  (optional) (default to "")
	pendingTimeInForce := "pendingTimeInForce_example" // string |  (optional) (default to "")
	pendingTrailingDelta := "pendingTrailingDelta_example" // string |  (optional) (default to "")
	selfTradePreventionMode := "selfTradePreventionMode_example" // string |  (optional) (default to "")
	sideEffectType := "sideEffectType_example" // string |  (optional) (default to "")
	workingClientOrderId := "workingClientOrderId_example" // string |  (optional) (default to "")
	workingTimeInForce := "workingTimeInForce_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.MarginCreateMarginOrderOtoV1(context.Background()).PendingQuantity(pendingQuantity).PendingSide(pendingSide).PendingType(pendingType).Symbol(symbol).WorkingIcebergQty(workingIcebergQty).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).WorkingSide(workingSide).WorkingType(workingType).AutoRepayAtCancel(autoRepayAtCancel).IsIsolated(isIsolated).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).PendingClientOrderId(pendingClientOrderId).PendingIcebergQty(pendingIcebergQty).PendingPrice(pendingPrice).PendingStopPrice(pendingStopPrice).PendingTimeInForce(pendingTimeInForce).PendingTrailingDelta(pendingTrailingDelta).SelfTradePreventionMode(selfTradePreventionMode).SideEffectType(sideEffectType).WorkingClientOrderId(workingClientOrderId).WorkingTimeInForce(workingTimeInForce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.MarginCreateMarginOrderOtoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginCreateMarginOrderOtoV1`: MarginCreateMarginOrderOtoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.MarginCreateMarginOrderOtoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginCreateMarginOrderOtoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pendingQuantity** | **string** |  | [default to &quot;&quot;]
 **pendingSide** | **string** |  | [default to &quot;&quot;]
 **pendingType** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **workingIcebergQty** | **string** |  | [default to &quot;&quot;]
 **workingPrice** | **string** |  | [default to &quot;&quot;]
 **workingQuantity** | **string** |  | [default to &quot;&quot;]
 **workingSide** | **string** |  | [default to &quot;&quot;]
 **workingType** | **string** |  | [default to &quot;&quot;]
 **autoRepayAtCancel** | **bool** |  | 
 **isIsolated** | **string** |  | [default to &quot;&quot;]
 **listClientOrderId** | **string** |  | [default to &quot;&quot;]
 **newOrderRespType** | **string** |  | [default to &quot;&quot;]
 **pendingClientOrderId** | **string** |  | [default to &quot;&quot;]
 **pendingIcebergQty** | **string** |  | [default to &quot;&quot;]
 **pendingPrice** | **string** |  | [default to &quot;&quot;]
 **pendingStopPrice** | **string** |  | [default to &quot;&quot;]
 **pendingTimeInForce** | **string** |  | [default to &quot;&quot;]
 **pendingTrailingDelta** | **string** |  | [default to &quot;&quot;]
 **selfTradePreventionMode** | **string** |  | [default to &quot;&quot;]
 **sideEffectType** | **string** |  | [default to &quot;&quot;]
 **workingClientOrderId** | **string** |  | [default to &quot;&quot;]
 **workingTimeInForce** | **string** |  | [default to &quot;&quot;]

### Return type

[**MarginCreateMarginOrderOtoV1Resp**](MarginCreateMarginOrderOtoV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginCreateMarginOrderOtocoV1

> MarginCreateMarginOrderOtocoV1Resp MarginCreateMarginOrderOtocoV1(ctx).PendingAboveType(pendingAboveType).PendingQuantity(pendingQuantity).PendingSide(pendingSide).Symbol(symbol).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).WorkingSide(workingSide).WorkingType(workingType).AutoRepayAtCancel(autoRepayAtCancel).IsIsolated(isIsolated).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).PendingAboveClientOrderId(pendingAboveClientOrderId).PendingAboveIcebergQty(pendingAboveIcebergQty).PendingAbovePrice(pendingAbovePrice).PendingAboveStopPrice(pendingAboveStopPrice).PendingAboveTimeInForce(pendingAboveTimeInForce).PendingAboveTrailingDelta(pendingAboveTrailingDelta).PendingBelowClientOrderId(pendingBelowClientOrderId).PendingBelowIcebergQty(pendingBelowIcebergQty).PendingBelowPrice(pendingBelowPrice).PendingBelowStopPrice(pendingBelowStopPrice).PendingBelowTimeInForce(pendingBelowTimeInForce).PendingBelowTrailingDelta(pendingBelowTrailingDelta).PendingBelowType(pendingBelowType).SelfTradePreventionMode(selfTradePreventionMode).SideEffectType(sideEffectType).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingTimeInForce(workingTimeInForce).Execute()

Margin Account New OTOCO (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/margin"
)

func main() {
	pendingAboveType := "pendingAboveType_example" // string |  (default to "")
	pendingQuantity := "pendingQuantity_example" // string |  (default to "")
	pendingSide := "pendingSide_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	workingPrice := "workingPrice_example" // string |  (default to "")
	workingQuantity := "workingQuantity_example" // string |  (default to "")
	workingSide := "workingSide_example" // string |  (default to "")
	workingType := "workingType_example" // string |  (default to "")
	autoRepayAtCancel := true // bool |  (optional)
	isIsolated := "isIsolated_example" // string |  (optional) (default to "")
	listClientOrderId := "listClientOrderId_example" // string |  (optional) (default to "")
	newOrderRespType := "newOrderRespType_example" // string |  (optional) (default to "")
	pendingAboveClientOrderId := "pendingAboveClientOrderId_example" // string |  (optional) (default to "")
	pendingAboveIcebergQty := "pendingAboveIcebergQty_example" // string |  (optional) (default to "")
	pendingAbovePrice := "pendingAbovePrice_example" // string |  (optional) (default to "")
	pendingAboveStopPrice := "pendingAboveStopPrice_example" // string |  (optional) (default to "")
	pendingAboveTimeInForce := "pendingAboveTimeInForce_example" // string |  (optional) (default to "")
	pendingAboveTrailingDelta := "pendingAboveTrailingDelta_example" // string |  (optional) (default to "")
	pendingBelowClientOrderId := "pendingBelowClientOrderId_example" // string |  (optional) (default to "")
	pendingBelowIcebergQty := "pendingBelowIcebergQty_example" // string |  (optional) (default to "")
	pendingBelowPrice := "pendingBelowPrice_example" // string |  (optional) (default to "")
	pendingBelowStopPrice := "pendingBelowStopPrice_example" // string |  (optional) (default to "")
	pendingBelowTimeInForce := "pendingBelowTimeInForce_example" // string |  (optional) (default to "")
	pendingBelowTrailingDelta := "pendingBelowTrailingDelta_example" // string |  (optional) (default to "")
	pendingBelowType := "pendingBelowType_example" // string |  (optional) (default to "")
	selfTradePreventionMode := "selfTradePreventionMode_example" // string |  (optional) (default to "")
	sideEffectType := "sideEffectType_example" // string |  (optional) (default to "")
	workingClientOrderId := "workingClientOrderId_example" // string |  (optional) (default to "")
	workingIcebergQty := "workingIcebergQty_example" // string |  (optional) (default to "")
	workingTimeInForce := "workingTimeInForce_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.MarginCreateMarginOrderOtocoV1(context.Background()).PendingAboveType(pendingAboveType).PendingQuantity(pendingQuantity).PendingSide(pendingSide).Symbol(symbol).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).WorkingSide(workingSide).WorkingType(workingType).AutoRepayAtCancel(autoRepayAtCancel).IsIsolated(isIsolated).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).PendingAboveClientOrderId(pendingAboveClientOrderId).PendingAboveIcebergQty(pendingAboveIcebergQty).PendingAbovePrice(pendingAbovePrice).PendingAboveStopPrice(pendingAboveStopPrice).PendingAboveTimeInForce(pendingAboveTimeInForce).PendingAboveTrailingDelta(pendingAboveTrailingDelta).PendingBelowClientOrderId(pendingBelowClientOrderId).PendingBelowIcebergQty(pendingBelowIcebergQty).PendingBelowPrice(pendingBelowPrice).PendingBelowStopPrice(pendingBelowStopPrice).PendingBelowTimeInForce(pendingBelowTimeInForce).PendingBelowTrailingDelta(pendingBelowTrailingDelta).PendingBelowType(pendingBelowType).SelfTradePreventionMode(selfTradePreventionMode).SideEffectType(sideEffectType).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingTimeInForce(workingTimeInForce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.MarginCreateMarginOrderOtocoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginCreateMarginOrderOtocoV1`: MarginCreateMarginOrderOtocoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.MarginCreateMarginOrderOtocoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginCreateMarginOrderOtocoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pendingAboveType** | **string** |  | [default to &quot;&quot;]
 **pendingQuantity** | **string** |  | [default to &quot;&quot;]
 **pendingSide** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **workingPrice** | **string** |  | [default to &quot;&quot;]
 **workingQuantity** | **string** |  | [default to &quot;&quot;]
 **workingSide** | **string** |  | [default to &quot;&quot;]
 **workingType** | **string** |  | [default to &quot;&quot;]
 **autoRepayAtCancel** | **bool** |  | 
 **isIsolated** | **string** |  | [default to &quot;&quot;]
 **listClientOrderId** | **string** |  | [default to &quot;&quot;]
 **newOrderRespType** | **string** |  | [default to &quot;&quot;]
 **pendingAboveClientOrderId** | **string** |  | [default to &quot;&quot;]
 **pendingAboveIcebergQty** | **string** |  | [default to &quot;&quot;]
 **pendingAbovePrice** | **string** |  | [default to &quot;&quot;]
 **pendingAboveStopPrice** | **string** |  | [default to &quot;&quot;]
 **pendingAboveTimeInForce** | **string** |  | [default to &quot;&quot;]
 **pendingAboveTrailingDelta** | **string** |  | [default to &quot;&quot;]
 **pendingBelowClientOrderId** | **string** |  | [default to &quot;&quot;]
 **pendingBelowIcebergQty** | **string** |  | [default to &quot;&quot;]
 **pendingBelowPrice** | **string** |  | [default to &quot;&quot;]
 **pendingBelowStopPrice** | **string** |  | [default to &quot;&quot;]
 **pendingBelowTimeInForce** | **string** |  | [default to &quot;&quot;]
 **pendingBelowTrailingDelta** | **string** |  | [default to &quot;&quot;]
 **pendingBelowType** | **string** |  | [default to &quot;&quot;]
 **selfTradePreventionMode** | **string** |  | [default to &quot;&quot;]
 **sideEffectType** | **string** |  | [default to &quot;&quot;]
 **workingClientOrderId** | **string** |  | [default to &quot;&quot;]
 **workingIcebergQty** | **string** |  | [default to &quot;&quot;]
 **workingTimeInForce** | **string** |  | [default to &quot;&quot;]

### Return type

[**MarginCreateMarginOrderOtocoV1Resp**](MarginCreateMarginOrderOtocoV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginCreateMarginOrderV1

> MarginCreateMarginOrderV1Resp MarginCreateMarginOrderV1(ctx).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).AutoRepayAtCancel(autoRepayAtCancel).IcebergQty(icebergQty).IsIsolated(isIsolated).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).Price(price).Quantity(quantity).QuoteOrderQty(quoteOrderQty).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).SideEffectType(sideEffectType).StopPrice(stopPrice).TimeInForce(timeInForce).Execute()

Margin Account New Order (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/margin"
)

func main() {
	side := "side_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	type_ := "type__example" // string |  (default to "")
	autoRepayAtCancel := true // bool |  (optional)
	icebergQty := "icebergQty_example" // string |  (optional) (default to "")
	isIsolated := "isIsolated_example" // string |  (optional) (default to "")
	newClientOrderId := "newClientOrderId_example" // string |  (optional) (default to "")
	newOrderRespType := "newOrderRespType_example" // string |  (optional) (default to "")
	price := "price_example" // string |  (optional) (default to "")
	quantity := "quantity_example" // string |  (optional) (default to "")
	quoteOrderQty := "quoteOrderQty_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	selfTradePreventionMode := "selfTradePreventionMode_example" // string |  (optional) (default to "")
	sideEffectType := "sideEffectType_example" // string |  (optional) (default to "")
	stopPrice := "stopPrice_example" // string |  (optional) (default to "")
	timeInForce := "timeInForce_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.MarginCreateMarginOrderV1(context.Background()).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).AutoRepayAtCancel(autoRepayAtCancel).IcebergQty(icebergQty).IsIsolated(isIsolated).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).Price(price).Quantity(quantity).QuoteOrderQty(quoteOrderQty).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).SideEffectType(sideEffectType).StopPrice(stopPrice).TimeInForce(timeInForce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.MarginCreateMarginOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginCreateMarginOrderV1`: MarginCreateMarginOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.MarginCreateMarginOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginCreateMarginOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **side** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **type_** | **string** |  | [default to &quot;&quot;]
 **autoRepayAtCancel** | **bool** |  | 
 **icebergQty** | **string** |  | [default to &quot;&quot;]
 **isIsolated** | **string** |  | [default to &quot;&quot;]
 **newClientOrderId** | **string** |  | [default to &quot;&quot;]
 **newOrderRespType** | **string** |  | [default to &quot;&quot;]
 **price** | **string** |  | [default to &quot;&quot;]
 **quantity** | **string** |  | [default to &quot;&quot;]
 **quoteOrderQty** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **selfTradePreventionMode** | **string** |  | [default to &quot;&quot;]
 **sideEffectType** | **string** |  | [default to &quot;&quot;]
 **stopPrice** | **string** |  | [default to &quot;&quot;]
 **timeInForce** | **string** |  | [default to &quot;&quot;]

### Return type

[**MarginCreateMarginOrderV1Resp**](MarginCreateMarginOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginDeleteMarginApiKeyV1

> map[string]interface{} MarginDeleteMarginApiKeyV1(ctx).Timestamp(timestamp).ApiKey(apiKey).ApiName(apiName).Symbol(symbol).RecvWindow(recvWindow).Execute()

Delete Special Key(Low-Latency Trading)(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/margin"
)

func main() {
	timestamp := int64(789) // int64 | 
	apiKey := "apiKey_example" // string |  (optional) (default to "")
	apiName := "apiName_example" // string |  (optional) (default to "")
	symbol := "symbol_example" // string | isolated margin pair (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.MarginDeleteMarginApiKeyV1(context.Background()).Timestamp(timestamp).ApiKey(apiKey).ApiName(apiName).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.MarginDeleteMarginApiKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginDeleteMarginApiKeyV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.MarginDeleteMarginApiKeyV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginDeleteMarginApiKeyV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **apiKey** | **string** |  | [default to &quot;&quot;]
 **apiName** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** | isolated margin pair | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

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


## MarginDeleteMarginOpenOrdersV1

> []MarginDeleteMarginOpenOrdersV1RespItem MarginDeleteMarginOpenOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).RecvWindow(recvWindow).Execute()

Margin Account Cancel all Open Orders on a Symbol (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/margin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	isIsolated := "isIsolated_example" // string | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;，default &#34;FALSE&#34; (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.MarginDeleteMarginOpenOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.MarginDeleteMarginOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginDeleteMarginOpenOrdersV1`: []MarginDeleteMarginOpenOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.MarginDeleteMarginOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginDeleteMarginOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **isIsolated** | **string** | for isolated margin or not, &amp;#34;TRUE&amp;#34;, &amp;#34;FALSE&amp;#34;，default &amp;#34;FALSE&amp;#34; | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]MarginDeleteMarginOpenOrdersV1RespItem**](MarginDeleteMarginOpenOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginDeleteMarginOrderListV1

> MarginDeleteMarginOrderListV1Resp MarginDeleteMarginOrderListV1(ctx).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()

Margin Account Cancel OCO (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/margin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	isIsolated := "isIsolated_example" // string | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;，default &#34;FALSE&#34; (optional) (default to "")
	orderListId := int64(789) // int64 | Either `orderListId` or `listClientOrderId` must be provided (optional)
	listClientOrderId := "listClientOrderId_example" // string | Either `orderListId` or `listClientOrderId` must be provided (optional) (default to "")
	newClientOrderId := "newClientOrderId_example" // string | Used to uniquely identify this cancel. Automatically generated by default (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.MarginDeleteMarginOrderListV1(context.Background()).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.MarginDeleteMarginOrderListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginDeleteMarginOrderListV1`: MarginDeleteMarginOrderListV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.MarginDeleteMarginOrderListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginDeleteMarginOrderListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **isIsolated** | **string** | for isolated margin or not, &amp;#34;TRUE&amp;#34;, &amp;#34;FALSE&amp;#34;，default &amp;#34;FALSE&amp;#34; | [default to &quot;&quot;]
 **orderListId** | **int64** | Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided | 
 **listClientOrderId** | **string** | Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided | [default to &quot;&quot;]
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**MarginDeleteMarginOrderListV1Resp**](MarginDeleteMarginOrderListV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginDeleteMarginOrderV1

> MarginDeleteMarginOrderV1Resp MarginDeleteMarginOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()

Margin Account Cancel Order (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/margin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	isIsolated := "isIsolated_example" // string | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;，default &#34;FALSE&#34; (optional) (default to "")
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	newClientOrderId := "newClientOrderId_example" // string | Used to uniquely identify this cancel. Automatically generated by default. (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.MarginDeleteMarginOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.MarginDeleteMarginOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginDeleteMarginOrderV1`: MarginDeleteMarginOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.MarginDeleteMarginOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginDeleteMarginOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **isIsolated** | **string** | for isolated margin or not, &amp;#34;TRUE&amp;#34;, &amp;#34;FALSE&amp;#34;，default &amp;#34;FALSE&amp;#34; | [default to &quot;&quot;]
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default. | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**MarginDeleteMarginOrderV1Resp**](MarginDeleteMarginOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginAllOrderListV1

> []MarginGetMarginAllOrderListV1RespItem MarginGetMarginAllOrderListV1(ctx).Timestamp(timestamp).IsIsolated(isIsolated).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query Margin Account's all OCO (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/margin"
)

func main() {
	timestamp := int64(789) // int64 | 
	isIsolated := "isIsolated_example" // string | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;，default &#34;FALSE&#34; (optional) (default to "")
	symbol := "symbol_example" // string | mandatory for isolated margin, not supported for cross margin (optional) (default to "")
	fromId := int64(789) // int64 | If supplied, neither `startTime` or `endTime` can be provided (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default Value: 500; Max Value: 1000 (optional)
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.MarginGetMarginAllOrderListV1(context.Background()).Timestamp(timestamp).IsIsolated(isIsolated).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.MarginGetMarginAllOrderListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginAllOrderListV1`: []MarginGetMarginAllOrderListV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.MarginGetMarginAllOrderListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginAllOrderListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **isIsolated** | **string** | for isolated margin or not, &amp;#34;TRUE&amp;#34;, &amp;#34;FALSE&amp;#34;，default &amp;#34;FALSE&amp;#34; | [default to &quot;&quot;]
 **symbol** | **string** | mandatory for isolated margin, not supported for cross margin | [default to &quot;&quot;]
 **fromId** | **int64** | If supplied, neither &#x60;startTime&#x60; or &#x60;endTime&#x60; can be provided | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default Value: 500; Max Value: 1000 | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]MarginGetMarginAllOrderListV1RespItem**](MarginGetMarginAllOrderListV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginAllOrdersV1

> []MarginGetMarginAllOrdersV1RespItem MarginGetMarginAllOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query Margin Account's All Orders (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/margin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	isIsolated := "isIsolated_example" // string | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;，default &#34;FALSE&#34; (optional) (default to "")
	orderId := int64(789) // int64 |  (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 500; max 500. (optional) (default to 500)
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.MarginGetMarginAllOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.MarginGetMarginAllOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginAllOrdersV1`: []MarginGetMarginAllOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.MarginGetMarginAllOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginAllOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **isIsolated** | **string** | for isolated margin or not, &amp;#34;TRUE&amp;#34;, &amp;#34;FALSE&amp;#34;，default &amp;#34;FALSE&amp;#34; | [default to &quot;&quot;]
 **orderId** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 500. | [default to 500]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]MarginGetMarginAllOrdersV1RespItem**](MarginGetMarginAllOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginApiKeyListV1

> []MarginGetMarginApiKeyListV1RespItem MarginGetMarginApiKeyListV1(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query Special key List(Low Latency Trading)(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/margin"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string | isolated margin pair (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.MarginGetMarginApiKeyListV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.MarginGetMarginApiKeyListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginApiKeyListV1`: []MarginGetMarginApiKeyListV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.MarginGetMarginApiKeyListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginApiKeyListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** | isolated margin pair | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]MarginGetMarginApiKeyListV1RespItem**](MarginGetMarginApiKeyListV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginApiKeyV1

> MarginGetMarginApiKeyV1Resp MarginGetMarginApiKeyV1(ctx).ApiKey(apiKey).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query Special key(Low Latency Trading)(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/margin"
)

func main() {
	apiKey := "apiKey_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string | isolated margin pair (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.MarginGetMarginApiKeyV1(context.Background()).ApiKey(apiKey).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.MarginGetMarginApiKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginApiKeyV1`: MarginGetMarginApiKeyV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.MarginGetMarginApiKeyV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginApiKeyV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **symbol** | **string** | isolated margin pair | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**MarginGetMarginApiKeyV1Resp**](MarginGetMarginApiKeyV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginExchangeSmallLiabilityHistoryV1

> MarginGetMarginExchangeSmallLiabilityHistoryV1Resp MarginGetMarginExchangeSmallLiabilityHistoryV1(ctx).Current(current).Size(size).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

Get Small Liability Exchange History (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/margin"
)

func main() {
	current := int32(56) // int32 | Currently querying page. Start from 1. Default:1
	size := int32(56) // int32 | Default:10, Max:100
	timestamp := int64(789) // int64 | 
	startTime := int64(789) // int64 | Default: 30 days from current timestamp (optional)
	endTime := int64(789) // int64 | Default: present timestamp (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.MarginGetMarginExchangeSmallLiabilityHistoryV1(context.Background()).Current(current).Size(size).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.MarginGetMarginExchangeSmallLiabilityHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginExchangeSmallLiabilityHistoryV1`: MarginGetMarginExchangeSmallLiabilityHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.MarginGetMarginExchangeSmallLiabilityHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginExchangeSmallLiabilityHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **current** | **int32** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int32** | Default:10, Max:100 | 
 **timestamp** | **int64** |  | 
 **startTime** | **int64** | Default: 30 days from current timestamp | 
 **endTime** | **int64** | Default: present timestamp | 
 **recvWindow** | **int64** |  | 

### Return type

[**MarginGetMarginExchangeSmallLiabilityHistoryV1Resp**](MarginGetMarginExchangeSmallLiabilityHistoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginExchangeSmallLiabilityV1

> []MarginGetMarginExchangeSmallLiabilityV1RespItem MarginGetMarginExchangeSmallLiabilityV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Small Liability Exchange Coin List (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/margin"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.MarginGetMarginExchangeSmallLiabilityV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.MarginGetMarginExchangeSmallLiabilityV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginExchangeSmallLiabilityV1`: []MarginGetMarginExchangeSmallLiabilityV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.MarginGetMarginExchangeSmallLiabilityV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginExchangeSmallLiabilityV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]MarginGetMarginExchangeSmallLiabilityV1RespItem**](MarginGetMarginExchangeSmallLiabilityV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginForceLiquidationRecV1

> MarginGetMarginForceLiquidationRecV1Resp MarginGetMarginForceLiquidationRecV1(ctx).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).IsolatedSymbol(isolatedSymbol).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Force Liquidation Record (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/margin"
)

func main() {
	timestamp := int64(789) // int64 | 
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	isolatedSymbol := "isolatedSymbol_example" // string |  (optional) (default to "")
	current := int64(789) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(789) // int64 | Default:10 Max:100 (optional)
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.MarginGetMarginForceLiquidationRecV1(context.Background()).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).IsolatedSymbol(isolatedSymbol).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.MarginGetMarginForceLiquidationRecV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginForceLiquidationRecV1`: MarginGetMarginForceLiquidationRecV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.MarginGetMarginForceLiquidationRecV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginForceLiquidationRecV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **isolatedSymbol** | **string** |  | [default to &quot;&quot;]
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10 Max:100 | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**MarginGetMarginForceLiquidationRecV1Resp**](MarginGetMarginForceLiquidationRecV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginMyTradesV1

> []MarginGetMarginMyTradesV1RespItem MarginGetMarginMyTradesV1(ctx).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()

Query Margin Account's Trade List (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/margin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	isIsolated := "isIsolated_example" // string | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;，default &#34;FALSE&#34; (optional) (default to "")
	orderId := int64(789) // int64 |  (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	fromId := int64(789) // int64 | TradeId to fetch from. Default gets most recent trades. (optional)
	limit := int32(56) // int32 | Default 500; max 1000. (optional) (default to 500)
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.MarginGetMarginMyTradesV1(context.Background()).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.MarginGetMarginMyTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginMyTradesV1`: []MarginGetMarginMyTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.MarginGetMarginMyTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginMyTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **isIsolated** | **string** | for isolated margin or not, &amp;#34;TRUE&amp;#34;, &amp;#34;FALSE&amp;#34;，default &amp;#34;FALSE&amp;#34; | [default to &quot;&quot;]
 **orderId** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **fromId** | **int64** | TradeId to fetch from. Default gets most recent trades. | 
 **limit** | **int32** | Default 500; max 1000. | [default to 500]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]MarginGetMarginMyTradesV1RespItem**](MarginGetMarginMyTradesV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginOpenOrderListV1

> []MarginGetMarginOpenOrderListV1RespItem MarginGetMarginOpenOrderListV1(ctx).Timestamp(timestamp).IsIsolated(isIsolated).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query Margin Account's Open OCO (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/margin"
)

func main() {
	timestamp := int64(789) // int64 | 
	isIsolated := "isIsolated_example" // string | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;，default &#34;FALSE&#34; (optional) (default to "")
	symbol := "symbol_example" // string | mandatory for isolated margin, not supported for cross margin (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.MarginGetMarginOpenOrderListV1(context.Background()).Timestamp(timestamp).IsIsolated(isIsolated).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.MarginGetMarginOpenOrderListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginOpenOrderListV1`: []MarginGetMarginOpenOrderListV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.MarginGetMarginOpenOrderListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginOpenOrderListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **isIsolated** | **string** | for isolated margin or not, &amp;#34;TRUE&amp;#34;, &amp;#34;FALSE&amp;#34;，default &amp;#34;FALSE&amp;#34; | [default to &quot;&quot;]
 **symbol** | **string** | mandatory for isolated margin, not supported for cross margin | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]MarginGetMarginOpenOrderListV1RespItem**](MarginGetMarginOpenOrderListV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginOpenOrdersV1

> []MarginGetMarginOpenOrdersV1RespItem MarginGetMarginOpenOrdersV1(ctx).Timestamp(timestamp).Symbol(symbol).IsIsolated(isIsolated).RecvWindow(recvWindow).Execute()

Query Margin Account's Open Orders (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/margin"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	isIsolated := "isIsolated_example" // string | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;，default &#34;FALSE&#34; (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.MarginGetMarginOpenOrdersV1(context.Background()).Timestamp(timestamp).Symbol(symbol).IsIsolated(isIsolated).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.MarginGetMarginOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginOpenOrdersV1`: []MarginGetMarginOpenOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.MarginGetMarginOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **isIsolated** | **string** | for isolated margin or not, &amp;#34;TRUE&amp;#34;, &amp;#34;FALSE&amp;#34;，default &amp;#34;FALSE&amp;#34; | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]MarginGetMarginOpenOrdersV1RespItem**](MarginGetMarginOpenOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginOrderListV1

> MarginGetMarginOrderListV1Resp MarginGetMarginOrderListV1(ctx).Timestamp(timestamp).IsIsolated(isIsolated).Symbol(symbol).OrderListId(orderListId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Margin Account's OCO (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/margin"
)

func main() {
	timestamp := int64(789) // int64 | 
	isIsolated := "isIsolated_example" // string | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;，default &#34;FALSE&#34; (optional) (default to "")
	symbol := "symbol_example" // string | mandatory for isolated margin, not supported for cross margin (optional) (default to "")
	orderListId := int64(789) // int64 | Either `orderListId` or `origClientOrderId` must be provided (optional)
	origClientOrderId := "origClientOrderId_example" // string | Either `orderListId` or `origClientOrderId` must be provided (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.MarginGetMarginOrderListV1(context.Background()).Timestamp(timestamp).IsIsolated(isIsolated).Symbol(symbol).OrderListId(orderListId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.MarginGetMarginOrderListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginOrderListV1`: MarginGetMarginOrderListV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.MarginGetMarginOrderListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginOrderListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **isIsolated** | **string** | for isolated margin or not, &amp;#34;TRUE&amp;#34;, &amp;#34;FALSE&amp;#34;，default &amp;#34;FALSE&amp;#34; | [default to &quot;&quot;]
 **symbol** | **string** | mandatory for isolated margin, not supported for cross margin | [default to &quot;&quot;]
 **orderListId** | **int64** | Either &#x60;orderListId&#x60; or &#x60;origClientOrderId&#x60; must be provided | 
 **origClientOrderId** | **string** | Either &#x60;orderListId&#x60; or &#x60;origClientOrderId&#x60; must be provided | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**MarginGetMarginOrderListV1Resp**](MarginGetMarginOrderListV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginOrderV1

> MarginGetMarginOrderV1Resp MarginGetMarginOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Margin Account's Order (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/margin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	isIsolated := "isIsolated_example" // string | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;，default &#34;FALSE&#34; (optional) (default to "")
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.MarginGetMarginOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.MarginGetMarginOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginOrderV1`: MarginGetMarginOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.MarginGetMarginOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **isIsolated** | **string** | for isolated margin or not, &amp;#34;TRUE&amp;#34;, &amp;#34;FALSE&amp;#34;，default &amp;#34;FALSE&amp;#34; | [default to &quot;&quot;]
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**MarginGetMarginOrderV1Resp**](MarginGetMarginOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginRateLimitOrderV1

> []MarginGetMarginRateLimitOrderV1RespItem MarginGetMarginRateLimitOrderV1(ctx).Timestamp(timestamp).IsIsolated(isIsolated).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query Current Margin Order Count Usage (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/margin"
)

func main() {
	timestamp := int64(789) // int64 | 
	isIsolated := "isIsolated_example" // string | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;，default &#34;FALSE&#34; (optional) (default to "")
	symbol := "symbol_example" // string | isolated symbol, mandatory for isolated margin (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.MarginGetMarginRateLimitOrderV1(context.Background()).Timestamp(timestamp).IsIsolated(isIsolated).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.MarginGetMarginRateLimitOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginRateLimitOrderV1`: []MarginGetMarginRateLimitOrderV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.MarginGetMarginRateLimitOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginRateLimitOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **isIsolated** | **string** | for isolated margin or not, &amp;#34;TRUE&amp;#34;, &amp;#34;FALSE&amp;#34;，default &amp;#34;FALSE&amp;#34; | [default to &quot;&quot;]
 **symbol** | **string** | isolated symbol, mandatory for isolated margin | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]MarginGetMarginRateLimitOrderV1RespItem**](MarginGetMarginRateLimitOrderV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginUpdateMarginApiKeyIpV1

> map[string]interface{} MarginUpdateMarginApiKeyIpV1(ctx).ApiKey(apiKey).Ip(ip).Timestamp(timestamp).RecvWindow(recvWindow).Symbol(symbol).Execute()

Edit ip for Special Key(Low-Latency Trading)(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/margin"
)

func main() {
	apiKey := "apiKey_example" // string |  (default to "")
	ip := "ip_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)
	symbol := "symbol_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.MarginUpdateMarginApiKeyIpV1(context.Background()).ApiKey(apiKey).Ip(ip).Timestamp(timestamp).RecvWindow(recvWindow).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.MarginUpdateMarginApiKeyIpV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginUpdateMarginApiKeyIpV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.MarginUpdateMarginApiKeyIpV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginUpdateMarginApiKeyIpV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** |  | [default to &quot;&quot;]
 **ip** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]

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

