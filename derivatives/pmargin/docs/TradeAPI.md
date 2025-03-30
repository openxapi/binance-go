# \TradeAPI

All URIs are relative to *https://papi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PmarginCreateCmConditionalOrderV1**](TradeAPI.md#PmarginCreateCmConditionalOrderV1) | **Post** /papi/v1/cm/conditional/order | New CM Conditional Order(TRADE)
[**PmarginCreateCmOrderV1**](TradeAPI.md#PmarginCreateCmOrderV1) | **Post** /papi/v1/cm/order | New CM Order(TRADE)
[**PmarginCreateMarginLoanV1**](TradeAPI.md#PmarginCreateMarginLoanV1) | **Post** /papi/v1/marginLoan | Margin Account Borrow(MARGIN)
[**PmarginCreateMarginOrderOcoV1**](TradeAPI.md#PmarginCreateMarginOrderOcoV1) | **Post** /papi/v1/margin/order/oco | Margin Account New OCO(TRADE)
[**PmarginCreateMarginOrderV1**](TradeAPI.md#PmarginCreateMarginOrderV1) | **Post** /papi/v1/margin/order | New Margin Order(TRADE)
[**PmarginCreateMarginRepayDebtV1**](TradeAPI.md#PmarginCreateMarginRepayDebtV1) | **Post** /papi/v1/margin/repay-debt | Margin Account Repay Debt(TRADE)
[**PmarginCreateRepayLoanV1**](TradeAPI.md#PmarginCreateRepayLoanV1) | **Post** /papi/v1/repayLoan | Margin Account Repay(MARGIN)
[**PmarginCreateUmConditionalOrderV1**](TradeAPI.md#PmarginCreateUmConditionalOrderV1) | **Post** /papi/v1/um/conditional/order | New UM Conditional Order (TRADE)
[**PmarginCreateUmFeeBurnV1**](TradeAPI.md#PmarginCreateUmFeeBurnV1) | **Post** /papi/v1/um/feeBurn | Toggle BNB Burn On UM Futures Trade (TRADE)
[**PmarginCreateUmOrderV1**](TradeAPI.md#PmarginCreateUmOrderV1) | **Post** /papi/v1/um/order | New UM Order (TRADE)
[**PmarginDeleteCmAllOpenOrdersV1**](TradeAPI.md#PmarginDeleteCmAllOpenOrdersV1) | **Delete** /papi/v1/cm/allOpenOrders | Cancel All CM Open Orders(TRADE)
[**PmarginDeleteCmConditionalAllOpenOrdersV1**](TradeAPI.md#PmarginDeleteCmConditionalAllOpenOrdersV1) | **Delete** /papi/v1/cm/conditional/allOpenOrders | Cancel All CM Open Conditional Orders(TRADE)
[**PmarginDeleteCmConditionalOrderV1**](TradeAPI.md#PmarginDeleteCmConditionalOrderV1) | **Delete** /papi/v1/cm/conditional/order | Cancel CM Conditional Order(TRADE)
[**PmarginDeleteCmOrderV1**](TradeAPI.md#PmarginDeleteCmOrderV1) | **Delete** /papi/v1/cm/order | Cancel CM Order(TRADE)
[**PmarginDeleteMarginAllOpenOrdersV1**](TradeAPI.md#PmarginDeleteMarginAllOpenOrdersV1) | **Delete** /papi/v1/margin/allOpenOrders | Cancel Margin Account All Open Orders on a Symbol(TRADE)
[**PmarginDeleteMarginOrderListV1**](TradeAPI.md#PmarginDeleteMarginOrderListV1) | **Delete** /papi/v1/margin/orderList | Cancel Margin Account OCO Orders(TRADE)
[**PmarginDeleteMarginOrderV1**](TradeAPI.md#PmarginDeleteMarginOrderV1) | **Delete** /papi/v1/margin/order | Cancel Margin Account Order(TRADE)
[**PmarginDeleteUmAllOpenOrdersV1**](TradeAPI.md#PmarginDeleteUmAllOpenOrdersV1) | **Delete** /papi/v1/um/allOpenOrders | Cancel All UM Open Orders(TRADE)
[**PmarginDeleteUmConditionalAllOpenOrdersV1**](TradeAPI.md#PmarginDeleteUmConditionalAllOpenOrdersV1) | **Delete** /papi/v1/um/conditional/allOpenOrders | Cancel All UM Open Conditional Orders (TRADE)
[**PmarginDeleteUmConditionalOrderV1**](TradeAPI.md#PmarginDeleteUmConditionalOrderV1) | **Delete** /papi/v1/um/conditional/order | Cancel UM Conditional Order(TRADE)
[**PmarginDeleteUmOrderV1**](TradeAPI.md#PmarginDeleteUmOrderV1) | **Delete** /papi/v1/um/order | Cancel UM Order(TRADE)
[**PmarginGetCmAdlQuantileV1**](TradeAPI.md#PmarginGetCmAdlQuantileV1) | **Get** /papi/v1/cm/adlQuantile | CM Position ADL Quantile Estimation(USER_DATA)
[**PmarginGetCmAllOrdersV1**](TradeAPI.md#PmarginGetCmAllOrdersV1) | **Get** /papi/v1/cm/allOrders | Query All CM Orders (USER_DATA)
[**PmarginGetCmConditionalAllOrdersV1**](TradeAPI.md#PmarginGetCmConditionalAllOrdersV1) | **Get** /papi/v1/cm/conditional/allOrders | Query All CM Conditional Orders(USER_DATA)
[**PmarginGetCmConditionalOpenOrderV1**](TradeAPI.md#PmarginGetCmConditionalOpenOrderV1) | **Get** /papi/v1/cm/conditional/openOrder | Query Current CM Open Conditional Order(USER_DATA)
[**PmarginGetCmConditionalOpenOrdersV1**](TradeAPI.md#PmarginGetCmConditionalOpenOrdersV1) | **Get** /papi/v1/cm/conditional/openOrders | Query All Current CM Open Conditional Orders (USER_DATA)
[**PmarginGetCmConditionalOrderHistoryV1**](TradeAPI.md#PmarginGetCmConditionalOrderHistoryV1) | **Get** /papi/v1/cm/conditional/orderHistory | Query CM Conditional Order History(USER_DATA)
[**PmarginGetCmForceOrdersV1**](TradeAPI.md#PmarginGetCmForceOrdersV1) | **Get** /papi/v1/cm/forceOrders | Query User&#39;s CM Force Orders(USER_DATA)
[**PmarginGetCmOpenOrderV1**](TradeAPI.md#PmarginGetCmOpenOrderV1) | **Get** /papi/v1/cm/openOrder | Query Current CM Open Order (USER_DATA)
[**PmarginGetCmOpenOrdersV1**](TradeAPI.md#PmarginGetCmOpenOrdersV1) | **Get** /papi/v1/cm/openOrders | Query All Current CM Open Orders(USER_DATA)
[**PmarginGetCmOrderAmendmentV1**](TradeAPI.md#PmarginGetCmOrderAmendmentV1) | **Get** /papi/v1/cm/orderAmendment | Query CM Modify Order History(TRADE)
[**PmarginGetCmOrderV1**](TradeAPI.md#PmarginGetCmOrderV1) | **Get** /papi/v1/cm/order | Query CM Order(USER_DATA)
[**PmarginGetCmUserTradesV1**](TradeAPI.md#PmarginGetCmUserTradesV1) | **Get** /papi/v1/cm/userTrades | CM Account Trade List(USER_DATA)
[**PmarginGetMarginAllOrderListV1**](TradeAPI.md#PmarginGetMarginAllOrderListV1) | **Get** /papi/v1/margin/allOrderList | Query Margin Account&#39;s all OCO (USER_DATA)
[**PmarginGetMarginAllOrdersV1**](TradeAPI.md#PmarginGetMarginAllOrdersV1) | **Get** /papi/v1/margin/allOrders | Query All Margin Account Orders (USER_DATA)
[**PmarginGetMarginForceOrdersV1**](TradeAPI.md#PmarginGetMarginForceOrdersV1) | **Get** /papi/v1/margin/forceOrders | Query User&#39;s Margin Force Orders(USER_DATA)
[**PmarginGetMarginMyTradesV1**](TradeAPI.md#PmarginGetMarginMyTradesV1) | **Get** /papi/v1/margin/myTrades | Margin Account Trade List (USER_DATA)
[**PmarginGetMarginOpenOrderListV1**](TradeAPI.md#PmarginGetMarginOpenOrderListV1) | **Get** /papi/v1/margin/openOrderList | Query Margin Account&#39;s Open OCO (USER_DATA)
[**PmarginGetMarginOpenOrdersV1**](TradeAPI.md#PmarginGetMarginOpenOrdersV1) | **Get** /papi/v1/margin/openOrders | Query Current Margin Open Order (USER_DATA)
[**PmarginGetMarginOrderListV1**](TradeAPI.md#PmarginGetMarginOrderListV1) | **Get** /papi/v1/margin/orderList | Query Margin Account&#39;s OCO (USER_DATA)
[**PmarginGetMarginOrderV1**](TradeAPI.md#PmarginGetMarginOrderV1) | **Get** /papi/v1/margin/order | Query Margin Account Order (USER_DATA)
[**PmarginGetUmAdlQuantileV1**](TradeAPI.md#PmarginGetUmAdlQuantileV1) | **Get** /papi/v1/um/adlQuantile | UM Position ADL Quantile Estimation(USER_DATA)
[**PmarginGetUmAllOrdersV1**](TradeAPI.md#PmarginGetUmAllOrdersV1) | **Get** /papi/v1/um/allOrders | Query All UM Orders(USER_DATA)
[**PmarginGetUmConditionalAllOrdersV1**](TradeAPI.md#PmarginGetUmConditionalAllOrdersV1) | **Get** /papi/v1/um/conditional/allOrders | Query All UM Conditional Orders(USER_DATA)
[**PmarginGetUmConditionalOpenOrderV1**](TradeAPI.md#PmarginGetUmConditionalOpenOrderV1) | **Get** /papi/v1/um/conditional/openOrder | Query Current UM Open Conditional Order(USER_DATA)
[**PmarginGetUmConditionalOpenOrdersV1**](TradeAPI.md#PmarginGetUmConditionalOpenOrdersV1) | **Get** /papi/v1/um/conditional/openOrders | Query All Current UM Open Conditional Orders(USER_DATA)
[**PmarginGetUmConditionalOrderHistoryV1**](TradeAPI.md#PmarginGetUmConditionalOrderHistoryV1) | **Get** /papi/v1/um/conditional/orderHistory | Query UM Conditional Order History(USER_DATA)
[**PmarginGetUmFeeBurnV1**](TradeAPI.md#PmarginGetUmFeeBurnV1) | **Get** /papi/v1/um/feeBurn | Get UM Futures BNB Burn Status (USER_DATA)
[**PmarginGetUmForceOrdersV1**](TradeAPI.md#PmarginGetUmForceOrdersV1) | **Get** /papi/v1/um/forceOrders | Query User&#39;s UM Force Orders (USER_DATA)
[**PmarginGetUmOpenOrderV1**](TradeAPI.md#PmarginGetUmOpenOrderV1) | **Get** /papi/v1/um/openOrder | Query Current UM Open Order(USER_DATA)
[**PmarginGetUmOpenOrdersV1**](TradeAPI.md#PmarginGetUmOpenOrdersV1) | **Get** /papi/v1/um/openOrders | Query All Current UM Open Orders(USER_DATA)
[**PmarginGetUmOrderAmendmentV1**](TradeAPI.md#PmarginGetUmOrderAmendmentV1) | **Get** /papi/v1/um/orderAmendment | Query UM Modify Order History(TRADE)
[**PmarginGetUmOrderV1**](TradeAPI.md#PmarginGetUmOrderV1) | **Get** /papi/v1/um/order | Query UM Order (USER_DATA)
[**PmarginGetUmUserTradesV1**](TradeAPI.md#PmarginGetUmUserTradesV1) | **Get** /papi/v1/um/userTrades | UM Account Trade List(USER_DATA)
[**PmarginUpdateCmOrderV1**](TradeAPI.md#PmarginUpdateCmOrderV1) | **Put** /papi/v1/cm/order | Modify CM Order(TRADE)
[**PmarginUpdateUmOrderV1**](TradeAPI.md#PmarginUpdateUmOrderV1) | **Put** /papi/v1/um/order | Modify UM Order(TRADE)



## PmarginCreateCmConditionalOrderV1

> PmarginCreateCmConditionalOrderV1Resp PmarginCreateCmConditionalOrderV1(ctx).Side(side).StrategyType(strategyType).Symbol(symbol).Timestamp(timestamp).ActivationPrice(activationPrice).CallbackRate(callbackRate).NewClientStrategyId(newClientStrategyId).PositionSide(positionSide).Price(price).PriceProtect(priceProtect).Quantity(quantity).RecvWindow(recvWindow).ReduceOnly(reduceOnly).StopPrice(stopPrice).TimeInForce(timeInForce).WorkingType(workingType).Execute()

New CM Conditional Order(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	side := "side_example" // string |  (default to "")
	strategyType := "strategyType_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	activationPrice := "activationPrice_example" // string |  (optional) (default to "")
	callbackRate := "callbackRate_example" // string |  (optional) (default to "")
	newClientStrategyId := "newClientStrategyId_example" // string |  (optional) (default to "")
	positionSide := "positionSide_example" // string |  (optional) (default to "")
	price := "price_example" // string |  (optional) (default to "")
	priceProtect := "priceProtect_example" // string |  (optional) (default to "")
	quantity := "quantity_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	reduceOnly := "reduceOnly_example" // string |  (optional) (default to "")
	stopPrice := "stopPrice_example" // string |  (optional) (default to "")
	timeInForce := "timeInForce_example" // string |  (optional) (default to "")
	workingType := "workingType_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginCreateCmConditionalOrderV1(context.Background()).Side(side).StrategyType(strategyType).Symbol(symbol).Timestamp(timestamp).ActivationPrice(activationPrice).CallbackRate(callbackRate).NewClientStrategyId(newClientStrategyId).PositionSide(positionSide).Price(price).PriceProtect(priceProtect).Quantity(quantity).RecvWindow(recvWindow).ReduceOnly(reduceOnly).StopPrice(stopPrice).TimeInForce(timeInForce).WorkingType(workingType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginCreateCmConditionalOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateCmConditionalOrderV1`: PmarginCreateCmConditionalOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginCreateCmConditionalOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginCreateCmConditionalOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **side** | **string** |  | [default to &quot;&quot;]
 **strategyType** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **activationPrice** | **string** |  | [default to &quot;&quot;]
 **callbackRate** | **string** |  | [default to &quot;&quot;]
 **newClientStrategyId** | **string** |  | [default to &quot;&quot;]
 **positionSide** | **string** |  | [default to &quot;&quot;]
 **price** | **string** |  | [default to &quot;&quot;]
 **priceProtect** | **string** |  | [default to &quot;&quot;]
 **quantity** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **reduceOnly** | **string** |  | [default to &quot;&quot;]
 **stopPrice** | **string** |  | [default to &quot;&quot;]
 **timeInForce** | **string** |  | [default to &quot;&quot;]
 **workingType** | **string** |  | [default to &quot;&quot;]

### Return type

[**PmarginCreateCmConditionalOrderV1Resp**](PmarginCreateCmConditionalOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginCreateCmOrderV1

> PmarginCreateCmOrderV1Resp PmarginCreateCmOrderV1(ctx).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).PositionSide(positionSide).Price(price).PriceMatch(priceMatch).Quantity(quantity).RecvWindow(recvWindow).ReduceOnly(reduceOnly).TimeInForce(timeInForce).Execute()

New CM Order(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	side := "side_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	type_ := "type__example" // string |  (default to "")
	newClientOrderId := "newClientOrderId_example" // string |  (optional) (default to "")
	newOrderRespType := "newOrderRespType_example" // string |  (optional) (default to "")
	positionSide := "positionSide_example" // string |  (optional) (default to "")
	price := "price_example" // string |  (optional) (default to "")
	priceMatch := "priceMatch_example" // string |  (optional) (default to "")
	quantity := "quantity_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	reduceOnly := "reduceOnly_example" // string |  (optional) (default to "")
	timeInForce := "timeInForce_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginCreateCmOrderV1(context.Background()).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).PositionSide(positionSide).Price(price).PriceMatch(priceMatch).Quantity(quantity).RecvWindow(recvWindow).ReduceOnly(reduceOnly).TimeInForce(timeInForce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginCreateCmOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateCmOrderV1`: PmarginCreateCmOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginCreateCmOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginCreateCmOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **side** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **type_** | **string** |  | [default to &quot;&quot;]
 **newClientOrderId** | **string** |  | [default to &quot;&quot;]
 **newOrderRespType** | **string** |  | [default to &quot;&quot;]
 **positionSide** | **string** |  | [default to &quot;&quot;]
 **price** | **string** |  | [default to &quot;&quot;]
 **priceMatch** | **string** |  | [default to &quot;&quot;]
 **quantity** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **reduceOnly** | **string** |  | [default to &quot;&quot;]
 **timeInForce** | **string** |  | [default to &quot;&quot;]

### Return type

[**PmarginCreateCmOrderV1Resp**](PmarginCreateCmOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginCreateMarginLoanV1

> PmarginCreateMarginLoanV1Resp PmarginCreateMarginLoanV1(ctx).Amount(amount).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Margin Account Borrow(MARGIN)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	amount := "amount_example" // string |  (default to "")
	asset := "asset_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginCreateMarginLoanV1(context.Background()).Amount(amount).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginCreateMarginLoanV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateMarginLoanV1`: PmarginCreateMarginLoanV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginCreateMarginLoanV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginCreateMarginLoanV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **asset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginCreateMarginLoanV1Resp**](PmarginCreateMarginLoanV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginCreateMarginOrderOcoV1

> PmarginCreateMarginOrderOcoV1Resp PmarginCreateMarginOrderOcoV1(ctx).Price(price).Quantity(quantity).Side(side).StopPrice(stopPrice).Symbol(symbol).Timestamp(timestamp).LimitClientOrderId(limitClientOrderId).LimitIcebergQty(limitIcebergQty).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).RecvWindow(recvWindow).SideEffectType(sideEffectType).StopClientOrderId(stopClientOrderId).StopIcebergQty(stopIcebergQty).StopLimitPrice(stopLimitPrice).StopLimitTimeInForce(stopLimitTimeInForce).Execute()

Margin Account New OCO(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
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
	listClientOrderId := "listClientOrderId_example" // string |  (optional) (default to "")
	newOrderRespType := "newOrderRespType_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	sideEffectType := "sideEffectType_example" // string |  (optional) (default to "")
	stopClientOrderId := "stopClientOrderId_example" // string |  (optional) (default to "")
	stopIcebergQty := "stopIcebergQty_example" // string |  (optional) (default to "")
	stopLimitPrice := "stopLimitPrice_example" // string |  (optional) (default to "")
	stopLimitTimeInForce := "stopLimitTimeInForce_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginCreateMarginOrderOcoV1(context.Background()).Price(price).Quantity(quantity).Side(side).StopPrice(stopPrice).Symbol(symbol).Timestamp(timestamp).LimitClientOrderId(limitClientOrderId).LimitIcebergQty(limitIcebergQty).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).RecvWindow(recvWindow).SideEffectType(sideEffectType).StopClientOrderId(stopClientOrderId).StopIcebergQty(stopIcebergQty).StopLimitPrice(stopLimitPrice).StopLimitTimeInForce(stopLimitTimeInForce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginCreateMarginOrderOcoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateMarginOrderOcoV1`: PmarginCreateMarginOrderOcoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginCreateMarginOrderOcoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginCreateMarginOrderOcoV1Request struct via the builder pattern


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
 **listClientOrderId** | **string** |  | [default to &quot;&quot;]
 **newOrderRespType** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **sideEffectType** | **string** |  | [default to &quot;&quot;]
 **stopClientOrderId** | **string** |  | [default to &quot;&quot;]
 **stopIcebergQty** | **string** |  | [default to &quot;&quot;]
 **stopLimitPrice** | **string** |  | [default to &quot;&quot;]
 **stopLimitTimeInForce** | **string** |  | [default to &quot;&quot;]

### Return type

[**PmarginCreateMarginOrderOcoV1Resp**](PmarginCreateMarginOrderOcoV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginCreateMarginOrderV1

> PmarginCreateMarginOrderV1Resp PmarginCreateMarginOrderV1(ctx).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).AutoRepayAtCancel(autoRepayAtCancel).IcebergQty(icebergQty).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).Price(price).Quantity(quantity).QuoteOrderQty(quoteOrderQty).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).SideEffectType(sideEffectType).StopPrice(stopPrice).TimeInForce(timeInForce).Execute()

New Margin Order(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	side := "side_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	type_ := "type__example" // string |  (default to "")
	autoRepayAtCancel := true // bool |  (optional)
	icebergQty := "icebergQty_example" // string |  (optional) (default to "")
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
	resp, r, err := apiClient.TradeAPI.PmarginCreateMarginOrderV1(context.Background()).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).AutoRepayAtCancel(autoRepayAtCancel).IcebergQty(icebergQty).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).Price(price).Quantity(quantity).QuoteOrderQty(quoteOrderQty).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).SideEffectType(sideEffectType).StopPrice(stopPrice).TimeInForce(timeInForce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginCreateMarginOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateMarginOrderV1`: PmarginCreateMarginOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginCreateMarginOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginCreateMarginOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **side** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **type_** | **string** |  | [default to &quot;&quot;]
 **autoRepayAtCancel** | **bool** |  | 
 **icebergQty** | **string** |  | [default to &quot;&quot;]
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

[**PmarginCreateMarginOrderV1Resp**](PmarginCreateMarginOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginCreateMarginRepayDebtV1

> PmarginCreateMarginRepayDebtV1Resp PmarginCreateMarginRepayDebtV1(ctx).Asset(asset).Timestamp(timestamp).Amount(amount).RecvWindow(recvWindow).SpecifyRepayAssets(specifyRepayAssets).Execute()

Margin Account Repay Debt(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	asset := "asset_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	amount := "amount_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	specifyRepayAssets := "specifyRepayAssets_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginCreateMarginRepayDebtV1(context.Background()).Asset(asset).Timestamp(timestamp).Amount(amount).RecvWindow(recvWindow).SpecifyRepayAssets(specifyRepayAssets).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginCreateMarginRepayDebtV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateMarginRepayDebtV1`: PmarginCreateMarginRepayDebtV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginCreateMarginRepayDebtV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginCreateMarginRepayDebtV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **amount** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **specifyRepayAssets** | **string** |  | [default to &quot;&quot;]

### Return type

[**PmarginCreateMarginRepayDebtV1Resp**](PmarginCreateMarginRepayDebtV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginCreateRepayLoanV1

> PmarginCreateRepayLoanV1Resp PmarginCreateRepayLoanV1(ctx).Amount(amount).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Margin Account Repay(MARGIN)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	amount := "amount_example" // string |  (default to "")
	asset := "asset_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginCreateRepayLoanV1(context.Background()).Amount(amount).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginCreateRepayLoanV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateRepayLoanV1`: PmarginCreateRepayLoanV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginCreateRepayLoanV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginCreateRepayLoanV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **asset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginCreateRepayLoanV1Resp**](PmarginCreateRepayLoanV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginCreateUmConditionalOrderV1

> PmarginCreateUmConditionalOrderV1Resp PmarginCreateUmConditionalOrderV1(ctx).Side(side).StrategyType(strategyType).Symbol(symbol).Timestamp(timestamp).ActivationPrice(activationPrice).CallbackRate(callbackRate).GoodTillDate(goodTillDate).NewClientStrategyId(newClientStrategyId).PositionSide(positionSide).Price(price).PriceMatch(priceMatch).PriceProtect(priceProtect).Quantity(quantity).RecvWindow(recvWindow).ReduceOnly(reduceOnly).SelfTradePreventionMode(selfTradePreventionMode).StopPrice(stopPrice).TimeInForce(timeInForce).WorkingType(workingType).Execute()

New UM Conditional Order (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	side := "side_example" // string |  (default to "")
	strategyType := "strategyType_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	activationPrice := "activationPrice_example" // string |  (optional) (default to "")
	callbackRate := "callbackRate_example" // string |  (optional) (default to "")
	goodTillDate := int64(789) // int64 |  (optional)
	newClientStrategyId := "newClientStrategyId_example" // string |  (optional) (default to "")
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
	resp, r, err := apiClient.TradeAPI.PmarginCreateUmConditionalOrderV1(context.Background()).Side(side).StrategyType(strategyType).Symbol(symbol).Timestamp(timestamp).ActivationPrice(activationPrice).CallbackRate(callbackRate).GoodTillDate(goodTillDate).NewClientStrategyId(newClientStrategyId).PositionSide(positionSide).Price(price).PriceMatch(priceMatch).PriceProtect(priceProtect).Quantity(quantity).RecvWindow(recvWindow).ReduceOnly(reduceOnly).SelfTradePreventionMode(selfTradePreventionMode).StopPrice(stopPrice).TimeInForce(timeInForce).WorkingType(workingType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginCreateUmConditionalOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateUmConditionalOrderV1`: PmarginCreateUmConditionalOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginCreateUmConditionalOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginCreateUmConditionalOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **side** | **string** |  | [default to &quot;&quot;]
 **strategyType** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **activationPrice** | **string** |  | [default to &quot;&quot;]
 **callbackRate** | **string** |  | [default to &quot;&quot;]
 **goodTillDate** | **int64** |  | 
 **newClientStrategyId** | **string** |  | [default to &quot;&quot;]
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

[**PmarginCreateUmConditionalOrderV1Resp**](PmarginCreateUmConditionalOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginCreateUmFeeBurnV1

> PmarginCreateUmFeeBurnV1Resp PmarginCreateUmFeeBurnV1(ctx).FeeBurn(feeBurn).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Toggle BNB Burn On UM Futures Trade (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	feeBurn := "feeBurn_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginCreateUmFeeBurnV1(context.Background()).FeeBurn(feeBurn).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginCreateUmFeeBurnV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateUmFeeBurnV1`: PmarginCreateUmFeeBurnV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginCreateUmFeeBurnV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginCreateUmFeeBurnV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **feeBurn** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginCreateUmFeeBurnV1Resp**](PmarginCreateUmFeeBurnV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginCreateUmOrderV1

> PmarginCreateUmOrderV1Resp PmarginCreateUmOrderV1(ctx).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).GoodTillDate(goodTillDate).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).PositionSide(positionSide).Price(price).PriceMatch(priceMatch).Quantity(quantity).RecvWindow(recvWindow).ReduceOnly(reduceOnly).SelfTradePreventionMode(selfTradePreventionMode).TimeInForce(timeInForce).Execute()

New UM Order (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	side := "side_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	type_ := "type__example" // string |  (default to "")
	goodTillDate := int64(789) // int64 |  (optional)
	newClientOrderId := "newClientOrderId_example" // string |  (optional) (default to "")
	newOrderRespType := "newOrderRespType_example" // string |  (optional) (default to "")
	positionSide := "positionSide_example" // string |  (optional) (default to "")
	price := "price_example" // string |  (optional) (default to "")
	priceMatch := "priceMatch_example" // string |  (optional) (default to "")
	quantity := "quantity_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	reduceOnly := "reduceOnly_example" // string |  (optional) (default to "")
	selfTradePreventionMode := "selfTradePreventionMode_example" // string |  (optional) (default to "")
	timeInForce := "timeInForce_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginCreateUmOrderV1(context.Background()).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).GoodTillDate(goodTillDate).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).PositionSide(positionSide).Price(price).PriceMatch(priceMatch).Quantity(quantity).RecvWindow(recvWindow).ReduceOnly(reduceOnly).SelfTradePreventionMode(selfTradePreventionMode).TimeInForce(timeInForce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginCreateUmOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateUmOrderV1`: PmarginCreateUmOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginCreateUmOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginCreateUmOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **side** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **type_** | **string** |  | [default to &quot;&quot;]
 **goodTillDate** | **int64** |  | 
 **newClientOrderId** | **string** |  | [default to &quot;&quot;]
 **newOrderRespType** | **string** |  | [default to &quot;&quot;]
 **positionSide** | **string** |  | [default to &quot;&quot;]
 **price** | **string** |  | [default to &quot;&quot;]
 **priceMatch** | **string** |  | [default to &quot;&quot;]
 **quantity** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **reduceOnly** | **string** |  | [default to &quot;&quot;]
 **selfTradePreventionMode** | **string** |  | [default to &quot;&quot;]
 **timeInForce** | **string** |  | [default to &quot;&quot;]

### Return type

[**PmarginCreateUmOrderV1Resp**](PmarginCreateUmOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginDeleteCmAllOpenOrdersV1

> PmarginDeleteCmAllOpenOrdersV1Resp PmarginDeleteCmAllOpenOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Cancel All CM Open Orders(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginDeleteCmAllOpenOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginDeleteCmAllOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginDeleteCmAllOpenOrdersV1`: PmarginDeleteCmAllOpenOrdersV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginDeleteCmAllOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginDeleteCmAllOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginDeleteCmAllOpenOrdersV1Resp**](PmarginDeleteCmAllOpenOrdersV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginDeleteCmConditionalAllOpenOrdersV1

> PmarginDeleteCmConditionalAllOpenOrdersV1Resp PmarginDeleteCmConditionalAllOpenOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Cancel All CM Open Conditional Orders(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginDeleteCmConditionalAllOpenOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginDeleteCmConditionalAllOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginDeleteCmConditionalAllOpenOrdersV1`: PmarginDeleteCmConditionalAllOpenOrdersV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginDeleteCmConditionalAllOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginDeleteCmConditionalAllOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginDeleteCmConditionalAllOpenOrdersV1Resp**](PmarginDeleteCmConditionalAllOpenOrdersV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginDeleteCmConditionalOrderV1

> PmarginDeleteCmConditionalOrderV1Resp PmarginDeleteCmConditionalOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()

Cancel CM Conditional Order(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	strategyId := int64(789) // int64 |  (optional)
	newClientStrategyId := "newClientStrategyId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginDeleteCmConditionalOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginDeleteCmConditionalOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginDeleteCmConditionalOrderV1`: PmarginDeleteCmConditionalOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginDeleteCmConditionalOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginDeleteCmConditionalOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **strategyId** | **int64** |  | 
 **newClientStrategyId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginDeleteCmConditionalOrderV1Resp**](PmarginDeleteCmConditionalOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginDeleteCmOrderV1

> PmarginDeleteCmOrderV1Resp PmarginDeleteCmOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Cancel CM Order(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginDeleteCmOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginDeleteCmOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginDeleteCmOrderV1`: PmarginDeleteCmOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginDeleteCmOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginDeleteCmOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginDeleteCmOrderV1Resp**](PmarginDeleteCmOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginDeleteMarginAllOpenOrdersV1

> []PmarginDeleteMarginAllOpenOrdersV1RespInner PmarginDeleteMarginAllOpenOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Cancel Margin Account All Open Orders on a Symbol(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginDeleteMarginAllOpenOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginDeleteMarginAllOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginDeleteMarginAllOpenOrdersV1`: []PmarginDeleteMarginAllOpenOrdersV1RespInner
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginDeleteMarginAllOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginDeleteMarginAllOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]PmarginDeleteMarginAllOpenOrdersV1RespInner**](PmarginDeleteMarginAllOpenOrdersV1RespInner.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginDeleteMarginOrderListV1

> PmarginDeleteMarginOrderListV1Resp PmarginDeleteMarginOrderListV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()

Cancel Margin Account OCO Orders(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	orderListId := int64(789) // int64 | Either `orderListId` or `listClientOrderId` must be provided (optional)
	listClientOrderId := "listClientOrderId_example" // string | Either `orderListId` or `listClientOrderId` must be provided (optional) (default to "")
	newClientOrderId := "newClientOrderId_example" // string | Used to uniquely identify this cancel. Automatically generated by default (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginDeleteMarginOrderListV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginDeleteMarginOrderListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginDeleteMarginOrderListV1`: PmarginDeleteMarginOrderListV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginDeleteMarginOrderListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginDeleteMarginOrderListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderListId** | **int64** | Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided | 
 **listClientOrderId** | **string** | Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided | [default to &quot;&quot;]
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**PmarginDeleteMarginOrderListV1Resp**](PmarginDeleteMarginOrderListV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginDeleteMarginOrderV1

> PmarginDeleteMarginOrderV1Resp PmarginDeleteMarginOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()

Cancel Margin Account Order(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	newClientOrderId := "newClientOrderId_example" // string | Used to uniquely identify this cancel. Automatically generated by default. (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginDeleteMarginOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginDeleteMarginOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginDeleteMarginOrderV1`: PmarginDeleteMarginOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginDeleteMarginOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginDeleteMarginOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default. | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**PmarginDeleteMarginOrderV1Resp**](PmarginDeleteMarginOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginDeleteUmAllOpenOrdersV1

> PmarginDeleteUmAllOpenOrdersV1Resp PmarginDeleteUmAllOpenOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Cancel All UM Open Orders(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginDeleteUmAllOpenOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginDeleteUmAllOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginDeleteUmAllOpenOrdersV1`: PmarginDeleteUmAllOpenOrdersV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginDeleteUmAllOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginDeleteUmAllOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginDeleteUmAllOpenOrdersV1Resp**](PmarginDeleteUmAllOpenOrdersV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginDeleteUmConditionalAllOpenOrdersV1

> PmarginDeleteUmConditionalAllOpenOrdersV1Resp PmarginDeleteUmConditionalAllOpenOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Cancel All UM Open Conditional Orders (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginDeleteUmConditionalAllOpenOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginDeleteUmConditionalAllOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginDeleteUmConditionalAllOpenOrdersV1`: PmarginDeleteUmConditionalAllOpenOrdersV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginDeleteUmConditionalAllOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginDeleteUmConditionalAllOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginDeleteUmConditionalAllOpenOrdersV1Resp**](PmarginDeleteUmConditionalAllOpenOrdersV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginDeleteUmConditionalOrderV1

> PmarginDeleteUmConditionalOrderV1Resp PmarginDeleteUmConditionalOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()

Cancel UM Conditional Order(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	strategyId := int64(789) // int64 |  (optional)
	newClientStrategyId := "newClientStrategyId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginDeleteUmConditionalOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginDeleteUmConditionalOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginDeleteUmConditionalOrderV1`: PmarginDeleteUmConditionalOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginDeleteUmConditionalOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginDeleteUmConditionalOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **strategyId** | **int64** |  | 
 **newClientStrategyId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginDeleteUmConditionalOrderV1Resp**](PmarginDeleteUmConditionalOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginDeleteUmOrderV1

> PmarginDeleteUmOrderV1Resp PmarginDeleteUmOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Cancel UM Order(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginDeleteUmOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginDeleteUmOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginDeleteUmOrderV1`: PmarginDeleteUmOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginDeleteUmOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginDeleteUmOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginDeleteUmOrderV1Resp**](PmarginDeleteUmOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetCmAdlQuantileV1

> []PmarginGetCmAdlQuantileV1RespItem PmarginGetCmAdlQuantileV1(ctx).Execute()

CM Position ADL Quantile Estimation(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginGetCmAdlQuantileV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetCmAdlQuantileV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmAdlQuantileV1`: []PmarginGetCmAdlQuantileV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetCmAdlQuantileV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetCmAdlQuantileV1Request struct via the builder pattern


### Return type

[**[]PmarginGetCmAdlQuantileV1RespItem**](PmarginGetCmAdlQuantileV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetCmAllOrdersV1

> []PmarginGetCmAllOrdersV1RespItem PmarginGetCmAllOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).Pair(pair).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query All CM Orders (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	pair := "pair_example" // string |  (optional) (default to "")
	orderId := int64(789) // int64 |  (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 50; max 100. (optional) (default to 50)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginGetCmAllOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).Pair(pair).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetCmAllOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmAllOrdersV1`: []PmarginGetCmAllOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetCmAllOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetCmAllOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **pair** | **string** |  | [default to &quot;&quot;]
 **orderId** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 50; max 100. | [default to 50]
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginGetCmAllOrdersV1RespItem**](PmarginGetCmAllOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetCmConditionalAllOrdersV1

> []PmarginGetCmConditionalAllOrdersV1RespItem PmarginGetCmConditionalAllOrdersV1(ctx).Timestamp(timestamp).Symbol(symbol).StrategyId(strategyId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query All CM Conditional Orders(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	strategyId := int64(789) // int64 |  (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 500; max 1000. (optional) (default to 500)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginGetCmConditionalAllOrdersV1(context.Background()).Timestamp(timestamp).Symbol(symbol).StrategyId(strategyId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetCmConditionalAllOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmConditionalAllOrdersV1`: []PmarginGetCmConditionalAllOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetCmConditionalAllOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetCmConditionalAllOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **strategyId** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 1000. | [default to 500]
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginGetCmConditionalAllOrdersV1RespItem**](PmarginGetCmConditionalAllOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetCmConditionalOpenOrderV1

> PmarginGetCmConditionalOpenOrderV1Resp PmarginGetCmConditionalOpenOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()

Query Current CM Open Conditional Order(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	strategyId := int64(789) // int64 |  (optional)
	newClientStrategyId := "newClientStrategyId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginGetCmConditionalOpenOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetCmConditionalOpenOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmConditionalOpenOrderV1`: PmarginGetCmConditionalOpenOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetCmConditionalOpenOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetCmConditionalOpenOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **strategyId** | **int64** |  | 
 **newClientStrategyId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetCmConditionalOpenOrderV1Resp**](PmarginGetCmConditionalOpenOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetCmConditionalOpenOrdersV1

> []PmarginGetCmConditionalOpenOrdersV1RespItem PmarginGetCmConditionalOpenOrdersV1(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query All Current CM Open Conditional Orders (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginGetCmConditionalOpenOrdersV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetCmConditionalOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmConditionalOpenOrdersV1`: []PmarginGetCmConditionalOpenOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetCmConditionalOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetCmConditionalOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginGetCmConditionalOpenOrdersV1RespItem**](PmarginGetCmConditionalOpenOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetCmConditionalOrderHistoryV1

> PmarginGetCmConditionalOrderHistoryV1Resp PmarginGetCmConditionalOrderHistoryV1(ctx).Symbol(symbol).Timestamp(timestamp).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()

Query CM Conditional Order History(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	strategyId := int64(789) // int64 |  (optional)
	newClientStrategyId := "newClientStrategyId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginGetCmConditionalOrderHistoryV1(context.Background()).Symbol(symbol).Timestamp(timestamp).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetCmConditionalOrderHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmConditionalOrderHistoryV1`: PmarginGetCmConditionalOrderHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetCmConditionalOrderHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetCmConditionalOrderHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **strategyId** | **int64** |  | 
 **newClientStrategyId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetCmConditionalOrderHistoryV1Resp**](PmarginGetCmConditionalOrderHistoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetCmForceOrdersV1

> []PmarginGetCmForceOrdersV1RespItem PmarginGetCmForceOrdersV1(ctx).Timestamp(timestamp).Symbol(symbol).AutoCloseType(autoCloseType).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query User's CM Force Orders(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	autoCloseType := "autoCloseType_example" // string | &#34;LIQUIDATION&#34; for liquidation orders, &#34;ADL&#34; for ADL orders. (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 50; max 100. (optional) (default to 50)
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginGetCmForceOrdersV1(context.Background()).Timestamp(timestamp).Symbol(symbol).AutoCloseType(autoCloseType).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetCmForceOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmForceOrdersV1`: []PmarginGetCmForceOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetCmForceOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetCmForceOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **autoCloseType** | **string** | &amp;#34;LIQUIDATION&amp;#34; for liquidation orders, &amp;#34;ADL&amp;#34; for ADL orders. | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 50; max 100. | [default to 50]
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]PmarginGetCmForceOrdersV1RespItem**](PmarginGetCmForceOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetCmOpenOrderV1

> PmarginGetCmOpenOrderV1Resp PmarginGetCmOpenOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Current CM Open Order (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginGetCmOpenOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetCmOpenOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmOpenOrderV1`: PmarginGetCmOpenOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetCmOpenOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetCmOpenOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetCmOpenOrderV1Resp**](PmarginGetCmOpenOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetCmOpenOrdersV1

> []PmarginGetCmOpenOrdersV1RespItem PmarginGetCmOpenOrdersV1(ctx).Timestamp(timestamp).Symbol(symbol).Pair(pair).RecvWindow(recvWindow).Execute()

Query All Current CM Open Orders(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	pair := "pair_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginGetCmOpenOrdersV1(context.Background()).Timestamp(timestamp).Symbol(symbol).Pair(pair).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetCmOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmOpenOrdersV1`: []PmarginGetCmOpenOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetCmOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetCmOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **pair** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginGetCmOpenOrdersV1RespItem**](PmarginGetCmOpenOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetCmOrderAmendmentV1

> []PmarginGetCmOrderAmendmentV1RespItem PmarginGetCmOrderAmendmentV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query CM Modify Order History(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 | Timestamp in ms to get modification history from INCLUSIVE (optional)
	endTime := int64(789) // int64 | Timestamp in ms to get modification history until INCLUSIVE (optional)
	limit := int32(56) // int32 | Default 50, max 100 (optional) (default to 50)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginGetCmOrderAmendmentV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetCmOrderAmendmentV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmOrderAmendmentV1`: []PmarginGetCmOrderAmendmentV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetCmOrderAmendmentV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetCmOrderAmendmentV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** | Timestamp in ms to get modification history from INCLUSIVE | 
 **endTime** | **int64** | Timestamp in ms to get modification history until INCLUSIVE | 
 **limit** | **int32** | Default 50, max 100 | [default to 50]
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginGetCmOrderAmendmentV1RespItem**](PmarginGetCmOrderAmendmentV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetCmOrderV1

> PmarginGetCmOrderV1Resp PmarginGetCmOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query CM Order(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginGetCmOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetCmOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmOrderV1`: PmarginGetCmOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetCmOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetCmOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetCmOrderV1Resp**](PmarginGetCmOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetCmUserTradesV1

> []PmarginGetCmUserTradesV1RespItem PmarginGetCmUserTradesV1(ctx).Timestamp(timestamp).Symbol(symbol).Pair(pair).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()

CM Account Trade List(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	pair := "pair_example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	fromId := int64(789) // int64 | Trade id to fetch from. Default gets most recent trades. (optional)
	limit := int32(56) // int32 | Default 50; max 1000. (optional) (default to 50)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginGetCmUserTradesV1(context.Background()).Timestamp(timestamp).Symbol(symbol).Pair(pair).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetCmUserTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmUserTradesV1`: []PmarginGetCmUserTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetCmUserTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetCmUserTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **pair** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **fromId** | **int64** | Trade id to fetch from. Default gets most recent trades. | 
 **limit** | **int32** | Default 50; max 1000. | [default to 50]
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginGetCmUserTradesV1RespItem**](PmarginGetCmUserTradesV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetMarginAllOrderListV1

> []PmarginGetMarginAllOrderListV1RespItem PmarginGetMarginAllOrderListV1(ctx).Timestamp(timestamp).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query Margin Account's all OCO (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	fromId := int64(789) // int64 | If supplied, neither startTime or endTime can be provided (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 500; max 500. (optional) (default to 500)
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginGetMarginAllOrderListV1(context.Background()).Timestamp(timestamp).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetMarginAllOrderListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetMarginAllOrderListV1`: []PmarginGetMarginAllOrderListV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetMarginAllOrderListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetMarginAllOrderListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **fromId** | **int64** | If supplied, neither startTime or endTime can be provided | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 500. | [default to 500]
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]PmarginGetMarginAllOrderListV1RespItem**](PmarginGetMarginAllOrderListV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetMarginAllOrdersV1

> []PmarginGetMarginAllOrdersV1RespItem PmarginGetMarginAllOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query All Margin Account Orders (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 |  (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 500; max 500. (optional) (default to 500)
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginGetMarginAllOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetMarginAllOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetMarginAllOrdersV1`: []PmarginGetMarginAllOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetMarginAllOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetMarginAllOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 500. | [default to 500]
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]PmarginGetMarginAllOrdersV1RespItem**](PmarginGetMarginAllOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetMarginForceOrdersV1

> PmarginGetMarginForceOrdersV1Resp PmarginGetMarginForceOrdersV1(ctx).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Query User's Margin Force Orders(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int64(789) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(789) // int64 | Default:10 Max:100 (optional)
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginGetMarginForceOrdersV1(context.Background()).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetMarginForceOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetMarginForceOrdersV1`: PmarginGetMarginForceOrdersV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetMarginForceOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetMarginForceOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10 Max:100 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**PmarginGetMarginForceOrdersV1Resp**](PmarginGetMarginForceOrdersV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetMarginMyTradesV1

> []PmarginGetMarginMyTradesV1RespItem PmarginGetMarginMyTradesV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()

Margin Account Trade List (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 |  (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	fromId := int64(789) // int64 | TradeId to fetch from. Default gets most recent trades. (optional)
	limit := int32(56) // int32 | Default 500; max 1000. (optional) (default to 500)
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginGetMarginMyTradesV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetMarginMyTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetMarginMyTradesV1`: []PmarginGetMarginMyTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetMarginMyTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetMarginMyTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **fromId** | **int64** | TradeId to fetch from. Default gets most recent trades. | 
 **limit** | **int32** | Default 500; max 1000. | [default to 500]
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]PmarginGetMarginMyTradesV1RespItem**](PmarginGetMarginMyTradesV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetMarginOpenOrderListV1

> []PmarginGetMarginOpenOrderListV1RespItem PmarginGetMarginOpenOrderListV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Margin Account's Open OCO (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginGetMarginOpenOrderListV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetMarginOpenOrderListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetMarginOpenOrderListV1`: []PmarginGetMarginOpenOrderListV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetMarginOpenOrderListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetMarginOpenOrderListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]PmarginGetMarginOpenOrderListV1RespItem**](PmarginGetMarginOpenOrderListV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetMarginOpenOrdersV1

> []PmarginGetMarginOpenOrdersV1RespItem PmarginGetMarginOpenOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Current Margin Open Order (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginGetMarginOpenOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetMarginOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetMarginOpenOrdersV1`: []PmarginGetMarginOpenOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetMarginOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetMarginOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]PmarginGetMarginOpenOrdersV1RespItem**](PmarginGetMarginOpenOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetMarginOrderListV1

> PmarginGetMarginOrderListV1Resp PmarginGetMarginOrderListV1(ctx).Timestamp(timestamp).OrderListId(orderListId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Margin Account's OCO (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	orderListId := int64(789) // int64 | Either orderListId or origClientOrderId must be provided (optional)
	origClientOrderId := "origClientOrderId_example" // string | Either orderListId or origClientOrderId must be provided (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginGetMarginOrderListV1(context.Background()).Timestamp(timestamp).OrderListId(orderListId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetMarginOrderListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetMarginOrderListV1`: PmarginGetMarginOrderListV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetMarginOrderListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetMarginOrderListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **orderListId** | **int64** | Either orderListId or origClientOrderId must be provided | 
 **origClientOrderId** | **string** | Either orderListId or origClientOrderId must be provided | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**PmarginGetMarginOrderListV1Resp**](PmarginGetMarginOrderListV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetMarginOrderV1

> PmarginGetMarginOrderV1Resp PmarginGetMarginOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Margin Account Order (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginGetMarginOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetMarginOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetMarginOrderV1`: PmarginGetMarginOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetMarginOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetMarginOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**PmarginGetMarginOrderV1Resp**](PmarginGetMarginOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmAdlQuantileV1

> []PmarginGetUmAdlQuantileV1RespItem PmarginGetUmAdlQuantileV1(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

UM Position ADL Quantile Estimation(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginGetUmAdlQuantileV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetUmAdlQuantileV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmAdlQuantileV1`: []PmarginGetUmAdlQuantileV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetUmAdlQuantileV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmAdlQuantileV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginGetUmAdlQuantileV1RespItem**](PmarginGetUmAdlQuantileV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmAllOrdersV1

> []PmarginGetUmAllOrdersV1RespItem PmarginGetUmAllOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query All UM Orders(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
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
	resp, r, err := apiClient.TradeAPI.PmarginGetUmAllOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetUmAllOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmAllOrdersV1`: []PmarginGetUmAllOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetUmAllOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmAllOrdersV1Request struct via the builder pattern


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

[**[]PmarginGetUmAllOrdersV1RespItem**](PmarginGetUmAllOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmConditionalAllOrdersV1

> []PmarginGetUmConditionalAllOrdersV1RespItem PmarginGetUmConditionalAllOrdersV1(ctx).Timestamp(timestamp).Symbol(symbol).StrategyId(strategyId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query All UM Conditional Orders(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	strategyId := int64(789) // int64 |  (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 500; max 1000. (optional) (default to 500)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginGetUmConditionalAllOrdersV1(context.Background()).Timestamp(timestamp).Symbol(symbol).StrategyId(strategyId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetUmConditionalAllOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmConditionalAllOrdersV1`: []PmarginGetUmConditionalAllOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetUmConditionalAllOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmConditionalAllOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **strategyId** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 1000. | [default to 500]
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginGetUmConditionalAllOrdersV1RespItem**](PmarginGetUmConditionalAllOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmConditionalOpenOrderV1

> PmarginGetUmConditionalOpenOrderV1Resp PmarginGetUmConditionalOpenOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()

Query Current UM Open Conditional Order(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	strategyId := int64(789) // int64 |  (optional)
	newClientStrategyId := "newClientStrategyId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginGetUmConditionalOpenOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetUmConditionalOpenOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmConditionalOpenOrderV1`: PmarginGetUmConditionalOpenOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetUmConditionalOpenOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmConditionalOpenOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **strategyId** | **int64** |  | 
 **newClientStrategyId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetUmConditionalOpenOrderV1Resp**](PmarginGetUmConditionalOpenOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmConditionalOpenOrdersV1

> []PmarginGetUmConditionalOpenOrdersV1RespItem PmarginGetUmConditionalOpenOrdersV1(ctx).Execute()

Query All Current UM Open Conditional Orders(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginGetUmConditionalOpenOrdersV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetUmConditionalOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmConditionalOpenOrdersV1`: []PmarginGetUmConditionalOpenOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetUmConditionalOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmConditionalOpenOrdersV1Request struct via the builder pattern


### Return type

[**[]PmarginGetUmConditionalOpenOrdersV1RespItem**](PmarginGetUmConditionalOpenOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmConditionalOrderHistoryV1

> PmarginGetUmConditionalOrderHistoryV1Resp PmarginGetUmConditionalOrderHistoryV1(ctx).Symbol(symbol).Timestamp(timestamp).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()

Query UM Conditional Order History(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	strategyId := int64(789) // int64 |  (optional)
	newClientStrategyId := "newClientStrategyId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginGetUmConditionalOrderHistoryV1(context.Background()).Symbol(symbol).Timestamp(timestamp).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetUmConditionalOrderHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmConditionalOrderHistoryV1`: PmarginGetUmConditionalOrderHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetUmConditionalOrderHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmConditionalOrderHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **strategyId** | **int64** |  | 
 **newClientStrategyId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetUmConditionalOrderHistoryV1Resp**](PmarginGetUmConditionalOrderHistoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmFeeBurnV1

> PmarginGetUmFeeBurnV1Resp PmarginGetUmFeeBurnV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get UM Futures BNB Burn Status (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginGetUmFeeBurnV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetUmFeeBurnV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmFeeBurnV1`: PmarginGetUmFeeBurnV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetUmFeeBurnV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmFeeBurnV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetUmFeeBurnV1Resp**](PmarginGetUmFeeBurnV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmForceOrdersV1

> []PmarginGetUmForceOrdersV1RespItem PmarginGetUmForceOrdersV1(ctx).Timestamp(timestamp).Symbol(symbol).AutoCloseType(autoCloseType).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query User's UM Force Orders (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	autoCloseType := "autoCloseType_example" // string | `LIQUIDATION` for liquidation orders, `ADL` for ADL orders. (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 50; max 100. (optional) (default to 50)
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginGetUmForceOrdersV1(context.Background()).Timestamp(timestamp).Symbol(symbol).AutoCloseType(autoCloseType).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetUmForceOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmForceOrdersV1`: []PmarginGetUmForceOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetUmForceOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmForceOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **autoCloseType** | **string** | &#x60;LIQUIDATION&#x60; for liquidation orders, &#x60;ADL&#x60; for ADL orders. | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 50; max 100. | [default to 50]
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]PmarginGetUmForceOrdersV1RespItem**](PmarginGetUmForceOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmOpenOrderV1

> PmarginGetUmOpenOrderV1Resp PmarginGetUmOpenOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Current UM Open Order(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginGetUmOpenOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetUmOpenOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmOpenOrderV1`: PmarginGetUmOpenOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetUmOpenOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmOpenOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetUmOpenOrderV1Resp**](PmarginGetUmOpenOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmOpenOrdersV1

> []PmarginGetUmOpenOrdersV1RespItem PmarginGetUmOpenOrdersV1(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query All Current UM Open Orders(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginGetUmOpenOrdersV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetUmOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmOpenOrdersV1`: []PmarginGetUmOpenOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetUmOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginGetUmOpenOrdersV1RespItem**](PmarginGetUmOpenOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmOrderAmendmentV1

> []PmarginGetUmOrderAmendmentV1RespItem PmarginGetUmOrderAmendmentV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query UM Modify Order History(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 | Timestamp in ms to get modification history from INCLUSIVE (optional)
	endTime := int64(789) // int64 | Timestamp in ms to get modification history until INCLUSIVE (optional)
	limit := int32(56) // int32 | Default 500, max 1000 (optional) (default to 500)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginGetUmOrderAmendmentV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetUmOrderAmendmentV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmOrderAmendmentV1`: []PmarginGetUmOrderAmendmentV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetUmOrderAmendmentV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmOrderAmendmentV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** | Timestamp in ms to get modification history from INCLUSIVE | 
 **endTime** | **int64** | Timestamp in ms to get modification history until INCLUSIVE | 
 **limit** | **int32** | Default 500, max 1000 | [default to 500]
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginGetUmOrderAmendmentV1RespItem**](PmarginGetUmOrderAmendmentV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmOrderV1

> PmarginGetUmOrderV1Resp PmarginGetUmOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query UM Order (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginGetUmOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetUmOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmOrderV1`: PmarginGetUmOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetUmOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetUmOrderV1Resp**](PmarginGetUmOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmUserTradesV1

> []PmarginGetUmUserTradesV1RespItem PmarginGetUmUserTradesV1(ctx).Symbol(symbol).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()

UM Account Trade List(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	fromId := int64(789) // int64 | Trade id to fetch from. Default gets most recent trades. (optional)
	limit := int32(56) // int32 | Default 500; max 1000. (optional) (default to 500)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.PmarginGetUmUserTradesV1(context.Background()).Symbol(symbol).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginGetUmUserTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmUserTradesV1`: []PmarginGetUmUserTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginGetUmUserTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmUserTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **fromId** | **int64** | Trade id to fetch from. Default gets most recent trades. | 
 **limit** | **int32** | Default 500; max 1000. | [default to 500]
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginGetUmUserTradesV1RespItem**](PmarginGetUmUserTradesV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginUpdateCmOrderV1

> PmarginUpdateCmOrderV1Resp PmarginUpdateCmOrderV1(ctx).Price(price).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).PriceMatch(priceMatch).RecvWindow(recvWindow).Execute()

Modify CM Order(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
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
	resp, r, err := apiClient.TradeAPI.PmarginUpdateCmOrderV1(context.Background()).Price(price).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).PriceMatch(priceMatch).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginUpdateCmOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginUpdateCmOrderV1`: PmarginUpdateCmOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginUpdateCmOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginUpdateCmOrderV1Request struct via the builder pattern


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

[**PmarginUpdateCmOrderV1Resp**](PmarginUpdateCmOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginUpdateUmOrderV1

> PmarginUpdateUmOrderV1Resp PmarginUpdateUmOrderV1(ctx).Price(price).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).PriceMatch(priceMatch).RecvWindow(recvWindow).Execute()

Modify UM Order(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
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
	resp, r, err := apiClient.TradeAPI.PmarginUpdateUmOrderV1(context.Background()).Price(price).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).PriceMatch(priceMatch).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.PmarginUpdateUmOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginUpdateUmOrderV1`: PmarginUpdateUmOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.PmarginUpdateUmOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginUpdateUmOrderV1Request struct via the builder pattern


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

[**PmarginUpdateUmOrderV1Resp**](PmarginUpdateUmOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

