# \TradeAPI

All URIs are relative to *https://fapi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UfuturesCreateBatchOrdersV1**](TradeAPI.md#UfuturesCreateBatchOrdersV1) | **Post** /fapi/v1/batchOrders | Place Multiple Orders(TRADE)
[**UfuturesCreateCountdownCancelAllV1**](TradeAPI.md#UfuturesCreateCountdownCancelAllV1) | **Post** /fapi/v1/countdownCancelAll | Auto-Cancel All Open Orders (TRADE)
[**UfuturesCreateLeverageV1**](TradeAPI.md#UfuturesCreateLeverageV1) | **Post** /fapi/v1/leverage | Change Initial Leverage(TRADE)
[**UfuturesCreateMarginTypeV1**](TradeAPI.md#UfuturesCreateMarginTypeV1) | **Post** /fapi/v1/marginType | Change Margin Type(TRADE)
[**UfuturesCreateMultiAssetsMarginV1**](TradeAPI.md#UfuturesCreateMultiAssetsMarginV1) | **Post** /fapi/v1/multiAssetsMargin | Change Multi-Assets Mode (TRADE)
[**UfuturesCreateOrderTestV1**](TradeAPI.md#UfuturesCreateOrderTestV1) | **Post** /fapi/v1/order/test | Test Order(TRADE)
[**UfuturesCreateOrderV1**](TradeAPI.md#UfuturesCreateOrderV1) | **Post** /fapi/v1/order | New Order(TRADE)
[**UfuturesCreatePositionMarginV1**](TradeAPI.md#UfuturesCreatePositionMarginV1) | **Post** /fapi/v1/positionMargin | Modify Isolated Position Margin(TRADE)
[**UfuturesCreatePositionSideDualV1**](TradeAPI.md#UfuturesCreatePositionSideDualV1) | **Post** /fapi/v1/positionSide/dual | Change Position Mode(TRADE)
[**UfuturesDeleteAllOpenOrdersV1**](TradeAPI.md#UfuturesDeleteAllOpenOrdersV1) | **Delete** /fapi/v1/allOpenOrders | Cancel All Open Orders (TRADE)
[**UfuturesDeleteBatchOrdersV1**](TradeAPI.md#UfuturesDeleteBatchOrdersV1) | **Delete** /fapi/v1/batchOrders | Cancel Multiple Orders (TRADE)
[**UfuturesDeleteOrderV1**](TradeAPI.md#UfuturesDeleteOrderV1) | **Delete** /fapi/v1/order | Cancel Order (TRADE)
[**UfuturesGetAdlQuantileV1**](TradeAPI.md#UfuturesGetAdlQuantileV1) | **Get** /fapi/v1/adlQuantile | Position ADL Quantile Estimation(USER_DATA)
[**UfuturesGetAllOrdersV1**](TradeAPI.md#UfuturesGetAllOrdersV1) | **Get** /fapi/v1/allOrders | All Orders (USER_DATA)
[**UfuturesGetForceOrdersV1**](TradeAPI.md#UfuturesGetForceOrdersV1) | **Get** /fapi/v1/forceOrders | User&#39;s Force Orders (USER_DATA)
[**UfuturesGetOpenOrderV1**](TradeAPI.md#UfuturesGetOpenOrderV1) | **Get** /fapi/v1/openOrder | Query Current Open Order (USER_DATA)
[**UfuturesGetOpenOrdersV1**](TradeAPI.md#UfuturesGetOpenOrdersV1) | **Get** /fapi/v1/openOrders | Current All Open Orders (USER_DATA)
[**UfuturesGetOrderAmendmentV1**](TradeAPI.md#UfuturesGetOrderAmendmentV1) | **Get** /fapi/v1/orderAmendment | Get Order Modify History (USER_DATA)
[**UfuturesGetOrderV1**](TradeAPI.md#UfuturesGetOrderV1) | **Get** /fapi/v1/order | Query Order (USER_DATA)
[**UfuturesGetPositionMarginHistoryV1**](TradeAPI.md#UfuturesGetPositionMarginHistoryV1) | **Get** /fapi/v1/positionMargin/history | Get Position Margin Change History (TRADE)
[**UfuturesGetPositionRiskV2**](TradeAPI.md#UfuturesGetPositionRiskV2) | **Get** /fapi/v2/positionRisk | Position Information V2 (USER_DATA)
[**UfuturesGetPositionRiskV3**](TradeAPI.md#UfuturesGetPositionRiskV3) | **Get** /fapi/v3/positionRisk | Position Information V3 (USER_DATA)
[**UfuturesGetUserTradesV1**](TradeAPI.md#UfuturesGetUserTradesV1) | **Get** /fapi/v1/userTrades | Account Trade List (USER_DATA)
[**UfuturesUpdateBatchOrdersV1**](TradeAPI.md#UfuturesUpdateBatchOrdersV1) | **Put** /fapi/v1/batchOrders | Modify Multiple Orders(TRADE)
[**UfuturesUpdateOrderV1**](TradeAPI.md#UfuturesUpdateOrderV1) | **Put** /fapi/v1/order | Modify Order (TRADE)



## UfuturesCreateBatchOrdersV1

> []UfuturesCreateBatchOrdersV1RespInner UfuturesCreateBatchOrdersV1(ctx).BatchOrders(batchOrders).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Place Multiple Orders(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func main() {
	batchOrders := []openapiclient.UfuturesCreateBatchOrdersV1ReqBatchOrdersItem{*openapiclient.NewUfuturesCreateBatchOrdersV1ReqBatchOrdersItem("Quantity_example", "Side_example", "Symbol_example", "Type_example")} // []UfuturesCreateBatchOrdersV1ReqBatchOrdersItem | 
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.UfuturesCreateBatchOrdersV1(context.Background()).BatchOrders(batchOrders).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.UfuturesCreateBatchOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesCreateBatchOrdersV1`: []UfuturesCreateBatchOrdersV1RespInner
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.UfuturesCreateBatchOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesCreateBatchOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchOrders** | [**[]UfuturesCreateBatchOrdersV1ReqBatchOrdersItem**](UfuturesCreateBatchOrdersV1ReqBatchOrdersItem.md) |  | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]UfuturesCreateBatchOrdersV1RespInner**](UfuturesCreateBatchOrdersV1RespInner.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesCreateCountdownCancelAllV1

> UfuturesCreateCountdownCancelAllV1Resp UfuturesCreateCountdownCancelAllV1(ctx).UfuturesCreateCountdownCancelAllV1Req(ufuturesCreateCountdownCancelAllV1Req).Execute()

Auto-Cancel All Open Orders (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func main() {
	ufuturesCreateCountdownCancelAllV1Req := *openapiclient.NewUfuturesCreateCountdownCancelAllV1Req(int64(123), "Symbol_example", int64(123)) // UfuturesCreateCountdownCancelAllV1Req |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.UfuturesCreateCountdownCancelAllV1(context.Background()).UfuturesCreateCountdownCancelAllV1Req(ufuturesCreateCountdownCancelAllV1Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.UfuturesCreateCountdownCancelAllV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesCreateCountdownCancelAllV1`: UfuturesCreateCountdownCancelAllV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.UfuturesCreateCountdownCancelAllV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesCreateCountdownCancelAllV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ufuturesCreateCountdownCancelAllV1Req** | [**UfuturesCreateCountdownCancelAllV1Req**](UfuturesCreateCountdownCancelAllV1Req.md) |  | 

### Return type

[**UfuturesCreateCountdownCancelAllV1Resp**](UfuturesCreateCountdownCancelAllV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesCreateLeverageV1

> UfuturesCreateLeverageV1Resp UfuturesCreateLeverageV1(ctx).Leverage(leverage).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Change Initial Leverage(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func main() {
	leverage := int32(56) // int32 | 
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.UfuturesCreateLeverageV1(context.Background()).Leverage(leverage).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.UfuturesCreateLeverageV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesCreateLeverageV1`: UfuturesCreateLeverageV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.UfuturesCreateLeverageV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesCreateLeverageV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **leverage** | **int32** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesCreateLeverageV1Resp**](UfuturesCreateLeverageV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesCreateMarginTypeV1

> UfuturesCreateMarginTypeV1Resp UfuturesCreateMarginTypeV1(ctx).MarginType(marginType).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Change Margin Type(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func main() {
	marginType := "marginType_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.UfuturesCreateMarginTypeV1(context.Background()).MarginType(marginType).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.UfuturesCreateMarginTypeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesCreateMarginTypeV1`: UfuturesCreateMarginTypeV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.UfuturesCreateMarginTypeV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesCreateMarginTypeV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marginType** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesCreateMarginTypeV1Resp**](UfuturesCreateMarginTypeV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesCreateMultiAssetsMarginV1

> UfuturesCreateMultiAssetsMarginV1Resp UfuturesCreateMultiAssetsMarginV1(ctx).MultiAssetsMargin(multiAssetsMargin).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Change Multi-Assets Mode (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func main() {
	multiAssetsMargin := "multiAssetsMargin_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.UfuturesCreateMultiAssetsMarginV1(context.Background()).MultiAssetsMargin(multiAssetsMargin).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.UfuturesCreateMultiAssetsMarginV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesCreateMultiAssetsMarginV1`: UfuturesCreateMultiAssetsMarginV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.UfuturesCreateMultiAssetsMarginV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesCreateMultiAssetsMarginV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **multiAssetsMargin** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesCreateMultiAssetsMarginV1Resp**](UfuturesCreateMultiAssetsMarginV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesCreateOrderTestV1

> UfuturesCreateOrderTestV1Resp UfuturesCreateOrderTestV1(ctx).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).ActivationPrice(activationPrice).CallbackRate(callbackRate).ClosePosition(closePosition).GoodTillDate(goodTillDate).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).PositionSide(positionSide).Price(price).PriceMatch(priceMatch).PriceProtect(priceProtect).Quantity(quantity).RecvWindow(recvWindow).ReduceOnly(reduceOnly).SelfTradePreventionMode(selfTradePreventionMode).StopPrice(stopPrice).TimeInForce(timeInForce).WorkingType(workingType).Execute()

Test Order(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func main() {
	side := "side_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	type_ := "type__example" // string |  (default to "")
	activationPrice := "activationPrice_example" // string |  (optional) (default to "")
	callbackRate := "callbackRate_example" // string |  (optional) (default to "")
	closePosition := "closePosition_example" // string |  (optional) (default to "")
	goodTillDate := int64(789) // int64 |  (optional)
	newClientOrderId := "newClientOrderId_example" // string |  (optional) (default to "")
	newOrderRespType := "newOrderRespType_example" // string |  (optional) (default to "")
	positionSide := "positionSide_example" // string |  (optional) (default to "")
	price := "price_example" // string |  (optional) (default to "")
	priceMatch := "priceMatch_example" // string |  (optional) (default to "")
	priceProtect := "priceProtect_example" // string |  (optional) (default to "")
	quantity := "quantity_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	reduceOnly := "reduceOnly_example" // string |  (optional) (default to "")
	selfTradePreventionMode := "selfTradePreventionMode_example" // string |  (optional) (default to "")
	stopPrice := "stopPrice_example" // string |  (optional) (default to "")
	timeInForce := "timeInForce_example" // string |  (optional) (default to "")
	workingType := "workingType_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.UfuturesCreateOrderTestV1(context.Background()).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).ActivationPrice(activationPrice).CallbackRate(callbackRate).ClosePosition(closePosition).GoodTillDate(goodTillDate).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).PositionSide(positionSide).Price(price).PriceMatch(priceMatch).PriceProtect(priceProtect).Quantity(quantity).RecvWindow(recvWindow).ReduceOnly(reduceOnly).SelfTradePreventionMode(selfTradePreventionMode).StopPrice(stopPrice).TimeInForce(timeInForce).WorkingType(workingType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.UfuturesCreateOrderTestV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesCreateOrderTestV1`: UfuturesCreateOrderTestV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.UfuturesCreateOrderTestV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesCreateOrderTestV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **side** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **type_** | **string** |  | [default to &quot;&quot;]
 **activationPrice** | **string** |  | [default to &quot;&quot;]
 **callbackRate** | **string** |  | [default to &quot;&quot;]
 **closePosition** | **string** |  | [default to &quot;&quot;]
 **goodTillDate** | **int64** |  | 
 **newClientOrderId** | **string** |  | [default to &quot;&quot;]
 **newOrderRespType** | **string** |  | [default to &quot;&quot;]
 **positionSide** | **string** |  | [default to &quot;&quot;]
 **price** | **string** |  | [default to &quot;&quot;]
 **priceMatch** | **string** |  | [default to &quot;&quot;]
 **priceProtect** | **string** |  | [default to &quot;&quot;]
 **quantity** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **reduceOnly** | **string** |  | [default to &quot;&quot;]
 **selfTradePreventionMode** | **string** |  | [default to &quot;&quot;]
 **stopPrice** | **string** |  | [default to &quot;&quot;]
 **timeInForce** | **string** |  | [default to &quot;&quot;]
 **workingType** | **string** |  | [default to &quot;&quot;]

### Return type

[**UfuturesCreateOrderTestV1Resp**](UfuturesCreateOrderTestV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesCreateOrderV1

> UfuturesCreateOrderV1Resp UfuturesCreateOrderV1(ctx).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).ActivationPrice(activationPrice).CallbackRate(callbackRate).ClosePosition(closePosition).GoodTillDate(goodTillDate).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).PositionSide(positionSide).Price(price).PriceMatch(priceMatch).PriceProtect(priceProtect).Quantity(quantity).RecvWindow(recvWindow).ReduceOnly(reduceOnly).SelfTradePreventionMode(selfTradePreventionMode).StopPrice(stopPrice).TimeInForce(timeInForce).WorkingType(workingType).Execute()

New Order(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func main() {
	side := "side_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	type_ := "type__example" // string |  (default to "")
	activationPrice := "activationPrice_example" // string |  (optional) (default to "")
	callbackRate := "callbackRate_example" // string |  (optional) (default to "")
	closePosition := "closePosition_example" // string |  (optional) (default to "")
	goodTillDate := int64(789) // int64 |  (optional)
	newClientOrderId := "newClientOrderId_example" // string |  (optional) (default to "")
	newOrderRespType := "newOrderRespType_example" // string |  (optional) (default to "")
	positionSide := "positionSide_example" // string |  (optional) (default to "")
	price := "price_example" // string |  (optional) (default to "")
	priceMatch := "priceMatch_example" // string |  (optional) (default to "")
	priceProtect := "priceProtect_example" // string |  (optional) (default to "")
	quantity := "quantity_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	reduceOnly := "reduceOnly_example" // string |  (optional) (default to "")
	selfTradePreventionMode := "selfTradePreventionMode_example" // string |  (optional) (default to "")
	stopPrice := "stopPrice_example" // string |  (optional) (default to "")
	timeInForce := "timeInForce_example" // string |  (optional) (default to "")
	workingType := "workingType_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.UfuturesCreateOrderV1(context.Background()).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).ActivationPrice(activationPrice).CallbackRate(callbackRate).ClosePosition(closePosition).GoodTillDate(goodTillDate).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).PositionSide(positionSide).Price(price).PriceMatch(priceMatch).PriceProtect(priceProtect).Quantity(quantity).RecvWindow(recvWindow).ReduceOnly(reduceOnly).SelfTradePreventionMode(selfTradePreventionMode).StopPrice(stopPrice).TimeInForce(timeInForce).WorkingType(workingType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.UfuturesCreateOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesCreateOrderV1`: UfuturesCreateOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.UfuturesCreateOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesCreateOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **side** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **type_** | **string** |  | [default to &quot;&quot;]
 **activationPrice** | **string** |  | [default to &quot;&quot;]
 **callbackRate** | **string** |  | [default to &quot;&quot;]
 **closePosition** | **string** |  | [default to &quot;&quot;]
 **goodTillDate** | **int64** |  | 
 **newClientOrderId** | **string** |  | [default to &quot;&quot;]
 **newOrderRespType** | **string** |  | [default to &quot;&quot;]
 **positionSide** | **string** |  | [default to &quot;&quot;]
 **price** | **string** |  | [default to &quot;&quot;]
 **priceMatch** | **string** |  | [default to &quot;&quot;]
 **priceProtect** | **string** |  | [default to &quot;&quot;]
 **quantity** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **reduceOnly** | **string** |  | [default to &quot;&quot;]
 **selfTradePreventionMode** | **string** |  | [default to &quot;&quot;]
 **stopPrice** | **string** |  | [default to &quot;&quot;]
 **timeInForce** | **string** |  | [default to &quot;&quot;]
 **workingType** | **string** |  | [default to &quot;&quot;]

### Return type

[**UfuturesCreateOrderV1Resp**](UfuturesCreateOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesCreatePositionMarginV1

> UfuturesCreatePositionMarginV1Resp UfuturesCreatePositionMarginV1(ctx).Amount(amount).Symbol(symbol).Timestamp(timestamp).Type_(type_).PositionSide(positionSide).RecvWindow(recvWindow).Execute()

Modify Isolated Position Margin(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func main() {
	amount := "amount_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	type_ := int32(56) // int32 | 
	positionSide := "positionSide_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.UfuturesCreatePositionMarginV1(context.Background()).Amount(amount).Symbol(symbol).Timestamp(timestamp).Type_(type_).PositionSide(positionSide).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.UfuturesCreatePositionMarginV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesCreatePositionMarginV1`: UfuturesCreatePositionMarginV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.UfuturesCreatePositionMarginV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesCreatePositionMarginV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **type_** | **int32** |  | 
 **positionSide** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesCreatePositionMarginV1Resp**](UfuturesCreatePositionMarginV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesCreatePositionSideDualV1

> UfuturesCreatePositionSideDualV1Resp UfuturesCreatePositionSideDualV1(ctx).DualSidePosition(dualSidePosition).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Change Position Mode(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func main() {
	dualSidePosition := "dualSidePosition_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.UfuturesCreatePositionSideDualV1(context.Background()).DualSidePosition(dualSidePosition).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.UfuturesCreatePositionSideDualV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesCreatePositionSideDualV1`: UfuturesCreatePositionSideDualV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.UfuturesCreatePositionSideDualV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesCreatePositionSideDualV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dualSidePosition** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesCreatePositionSideDualV1Resp**](UfuturesCreatePositionSideDualV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesDeleteAllOpenOrdersV1

> UfuturesDeleteAllOpenOrdersV1Resp UfuturesDeleteAllOpenOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Cancel All Open Orders (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.UfuturesDeleteAllOpenOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.UfuturesDeleteAllOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesDeleteAllOpenOrdersV1`: UfuturesDeleteAllOpenOrdersV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.UfuturesDeleteAllOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesDeleteAllOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesDeleteAllOpenOrdersV1Resp**](UfuturesDeleteAllOpenOrdersV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesDeleteBatchOrdersV1

> []UfuturesDeleteBatchOrdersV1RespInner UfuturesDeleteBatchOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderIdList(orderIdList).OrigClientOrderIdList(origClientOrderIdList).RecvWindow(recvWindow).Execute()

Cancel Multiple Orders (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	orderIdList := []int64{int64(123)} // []int64 | max length 10 <br/> e.g. [1234567,2345678] (optional)
	origClientOrderIdList := []string{"Inner_example"} // []string | max length 10<br/> e.g. [&#34;my_id_1&#34;,&#34;my_id_2&#34;], encode the double quotes. No space after comma. (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.UfuturesDeleteBatchOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderIdList(orderIdList).OrigClientOrderIdList(origClientOrderIdList).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.UfuturesDeleteBatchOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesDeleteBatchOrdersV1`: []UfuturesDeleteBatchOrdersV1RespInner
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.UfuturesDeleteBatchOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesDeleteBatchOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderIdList** | **[]int64** | max length 10 &lt;br/&gt; e.g. [1234567,2345678] | 
 **origClientOrderIdList** | **[]string** | max length 10&lt;br/&gt; e.g. [&amp;#34;my_id_1&amp;#34;,&amp;#34;my_id_2&amp;#34;], encode the double quotes. No space after comma. | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]UfuturesDeleteBatchOrdersV1RespInner**](UfuturesDeleteBatchOrdersV1RespInner.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesDeleteOrderV1

> UfuturesDeleteOrderV1Resp UfuturesDeleteOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Cancel Order (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.UfuturesDeleteOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.UfuturesDeleteOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesDeleteOrderV1`: UfuturesDeleteOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.UfuturesDeleteOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesDeleteOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesDeleteOrderV1Resp**](UfuturesDeleteOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetAdlQuantileV1

> []UfuturesGetAdlQuantileV1RespItem UfuturesGetAdlQuantileV1(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Position ADL Quantile Estimation(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.UfuturesGetAdlQuantileV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.UfuturesGetAdlQuantileV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetAdlQuantileV1`: []UfuturesGetAdlQuantileV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.UfuturesGetAdlQuantileV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetAdlQuantileV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]UfuturesGetAdlQuantileV1RespItem**](UfuturesGetAdlQuantileV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetAllOrdersV1

> []UfuturesGetAllOrdersV1RespItem UfuturesGetAllOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

All Orders (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 |  (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 500; max 1000. (optional) (default to 500)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.UfuturesGetAllOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.UfuturesGetAllOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetAllOrdersV1`: []UfuturesGetAllOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.UfuturesGetAllOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetAllOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 1000. | [default to 500]
 **recvWindow** | **int64** |  | 

### Return type

[**[]UfuturesGetAllOrdersV1RespItem**](UfuturesGetAllOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetForceOrdersV1

> []UfuturesGetForceOrdersV1RespItem UfuturesGetForceOrdersV1(ctx).Timestamp(timestamp).Symbol(symbol).AutoCloseType(autoCloseType).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

User's Force Orders (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	autoCloseType := "autoCloseType_example" // string | &#34;LIQUIDATION&#34; for liquidation orders, &#34;ADL&#34; for ADL orders. (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 50; max 100. (optional) (default to 50)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.UfuturesGetForceOrdersV1(context.Background()).Timestamp(timestamp).Symbol(symbol).AutoCloseType(autoCloseType).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.UfuturesGetForceOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetForceOrdersV1`: []UfuturesGetForceOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.UfuturesGetForceOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetForceOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **autoCloseType** | **string** | &amp;#34;LIQUIDATION&amp;#34; for liquidation orders, &amp;#34;ADL&amp;#34; for ADL orders. | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 50; max 100. | [default to 50]
 **recvWindow** | **int64** |  | 

### Return type

[**[]UfuturesGetForceOrdersV1RespItem**](UfuturesGetForceOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetOpenOrderV1

> UfuturesGetOpenOrderV1Resp UfuturesGetOpenOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Current Open Order (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.UfuturesGetOpenOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.UfuturesGetOpenOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetOpenOrderV1`: UfuturesGetOpenOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.UfuturesGetOpenOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetOpenOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetOpenOrderV1Resp**](UfuturesGetOpenOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetOpenOrdersV1

> []UfuturesGetOpenOrdersV1RespItem UfuturesGetOpenOrdersV1(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Current All Open Orders (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.UfuturesGetOpenOrdersV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.UfuturesGetOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetOpenOrdersV1`: []UfuturesGetOpenOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.UfuturesGetOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]UfuturesGetOpenOrdersV1RespItem**](UfuturesGetOpenOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetOrderAmendmentV1

> []UfuturesGetOrderAmendmentV1RespItem UfuturesGetOrderAmendmentV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Get Order Modify History (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 | Timestamp in ms to get modification history from INCLUSIVE (optional)
	endTime := int64(789) // int64 | Timestamp in ms to get modification history until INCLUSIVE (optional)
	limit := int32(56) // int32 | Default 50; max 100 (optional) (default to 50)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.UfuturesGetOrderAmendmentV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.UfuturesGetOrderAmendmentV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetOrderAmendmentV1`: []UfuturesGetOrderAmendmentV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.UfuturesGetOrderAmendmentV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetOrderAmendmentV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** | Timestamp in ms to get modification history from INCLUSIVE | 
 **endTime** | **int64** | Timestamp in ms to get modification history until INCLUSIVE | 
 **limit** | **int32** | Default 50; max 100 | [default to 50]
 **recvWindow** | **int64** |  | 

### Return type

[**[]UfuturesGetOrderAmendmentV1RespItem**](UfuturesGetOrderAmendmentV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetOrderV1

> UfuturesGetOrderV1Resp UfuturesGetOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Order (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.UfuturesGetOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.UfuturesGetOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetOrderV1`: UfuturesGetOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.UfuturesGetOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetOrderV1Resp**](UfuturesGetOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetPositionMarginHistoryV1

> []UfuturesGetPositionMarginHistoryV1RespItem UfuturesGetPositionMarginHistoryV1(ctx).Symbol(symbol).Timestamp(timestamp).Type_(type_).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Get Position Margin Change History (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	type_ := int32(56) // int32 | 1: Add position margin，2: Reduce position margin (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 | Default current time if not pass (optional)
	limit := int32(56) // int32 | Default: 500 (optional) (default to 500)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.UfuturesGetPositionMarginHistoryV1(context.Background()).Symbol(symbol).Timestamp(timestamp).Type_(type_).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.UfuturesGetPositionMarginHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetPositionMarginHistoryV1`: []UfuturesGetPositionMarginHistoryV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.UfuturesGetPositionMarginHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetPositionMarginHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **type_** | **int32** | 1: Add position margin，2: Reduce position margin | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** | Default current time if not pass | 
 **limit** | **int32** | Default: 500 | [default to 500]
 **recvWindow** | **int64** |  | 

### Return type

[**[]UfuturesGetPositionMarginHistoryV1RespItem**](UfuturesGetPositionMarginHistoryV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetPositionRiskV2

> []UfuturesGetPositionRiskV2RespItem UfuturesGetPositionRiskV2(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Position Information V2 (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.UfuturesGetPositionRiskV2(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.UfuturesGetPositionRiskV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetPositionRiskV2`: []UfuturesGetPositionRiskV2RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.UfuturesGetPositionRiskV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetPositionRiskV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]UfuturesGetPositionRiskV2RespItem**](UfuturesGetPositionRiskV2RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetPositionRiskV3

> []UfuturesGetPositionRiskV3RespItem UfuturesGetPositionRiskV3(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Position Information V3 (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.UfuturesGetPositionRiskV3(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.UfuturesGetPositionRiskV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetPositionRiskV3`: []UfuturesGetPositionRiskV3RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.UfuturesGetPositionRiskV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetPositionRiskV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]UfuturesGetPositionRiskV3RespItem**](UfuturesGetPositionRiskV3RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetUserTradesV1

> []UfuturesGetUserTradesV1RespItem UfuturesGetUserTradesV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()

Account Trade List (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 | This can only be used in combination with `symbol` (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	fromId := int64(789) // int64 | Trade id to fetch from. Default gets most recent trades. (optional)
	limit := int32(56) // int32 | Default 500; max 1000. (optional) (default to 500)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.UfuturesGetUserTradesV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.UfuturesGetUserTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetUserTradesV1`: []UfuturesGetUserTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.UfuturesGetUserTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetUserTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** | This can only be used in combination with &#x60;symbol&#x60; | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **fromId** | **int64** | Trade id to fetch from. Default gets most recent trades. | 
 **limit** | **int32** | Default 500; max 1000. | [default to 500]
 **recvWindow** | **int64** |  | 

### Return type

[**[]UfuturesGetUserTradesV1RespItem**](UfuturesGetUserTradesV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesUpdateBatchOrdersV1

> []UfuturesUpdateBatchOrdersV1RespItem UfuturesUpdateBatchOrdersV1(ctx).BatchOrders(batchOrders).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Modify Multiple Orders(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func main() {
	batchOrders := []openapiclient.UfuturesUpdateBatchOrdersV1ReqBatchOrdersItem{*openapiclient.NewUfuturesUpdateBatchOrdersV1ReqBatchOrdersItem("Price_example", "Quantity_example", "Side_example", "Symbol_example")} // []UfuturesUpdateBatchOrdersV1ReqBatchOrdersItem | 
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.UfuturesUpdateBatchOrdersV1(context.Background()).BatchOrders(batchOrders).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.UfuturesUpdateBatchOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesUpdateBatchOrdersV1`: []UfuturesUpdateBatchOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.UfuturesUpdateBatchOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesUpdateBatchOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchOrders** | [**[]UfuturesUpdateBatchOrdersV1ReqBatchOrdersItem**](UfuturesUpdateBatchOrdersV1ReqBatchOrdersItem.md) |  | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]UfuturesUpdateBatchOrdersV1RespItem**](UfuturesUpdateBatchOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesUpdateOrderV1

> UfuturesUpdateOrderV1Resp UfuturesUpdateOrderV1(ctx).Price(price).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).PriceMatch(priceMatch).RecvWindow(recvWindow).Execute()

Modify Order (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func main() {
	price := "price_example" // string |  (default to "")
	quantity := "quantity_example" // string |  (default to "")
	side := "side_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	priceMatch := "priceMatch_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.UfuturesUpdateOrderV1(context.Background()).Price(price).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).PriceMatch(priceMatch).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.UfuturesUpdateOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesUpdateOrderV1`: UfuturesUpdateOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.UfuturesUpdateOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesUpdateOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **price** | **string** |  | [default to &quot;&quot;]
 **quantity** | **string** |  | [default to &quot;&quot;]
 **side** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **priceMatch** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesUpdateOrderV1Resp**](UfuturesUpdateOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

