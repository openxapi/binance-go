# \V1API

All URIs are relative to *https://dapi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CfuturesCreateBatchOrdersV1**](V1API.md#CfuturesCreateBatchOrdersV1) | **Post** /dapi/v1/batchOrders | Place Multiple Orders(TRADE)
[**CfuturesCreateCountdownCancelAllV1**](V1API.md#CfuturesCreateCountdownCancelAllV1) | **Post** /dapi/v1/countdownCancelAll | Auto-Cancel All Open Orders (TRADE)
[**CfuturesCreateLeverageV1**](V1API.md#CfuturesCreateLeverageV1) | **Post** /dapi/v1/leverage | Change Initial Leverage (TRADE)
[**CfuturesCreateListenKeyV1**](V1API.md#CfuturesCreateListenKeyV1) | **Post** /dapi/v1/listenKey | Start User Data Stream (USER_STREAM)
[**CfuturesCreateMarginTypeV1**](V1API.md#CfuturesCreateMarginTypeV1) | **Post** /dapi/v1/marginType | Change Margin Type (TRADE)
[**CfuturesCreateOrderV1**](V1API.md#CfuturesCreateOrderV1) | **Post** /dapi/v1/order | New Order (TRADE)
[**CfuturesCreatePositionMarginV1**](V1API.md#CfuturesCreatePositionMarginV1) | **Post** /dapi/v1/positionMargin | Modify Isolated Position Margin(TRADE)
[**CfuturesCreatePositionSideDualV1**](V1API.md#CfuturesCreatePositionSideDualV1) | **Post** /dapi/v1/positionSide/dual | Change Position Mode(TRADE)
[**CfuturesDeleteAllOpenOrdersV1**](V1API.md#CfuturesDeleteAllOpenOrdersV1) | **Delete** /dapi/v1/allOpenOrders | Cancel All Open Orders(TRADE)
[**CfuturesDeleteBatchOrdersV1**](V1API.md#CfuturesDeleteBatchOrdersV1) | **Delete** /dapi/v1/batchOrders | Cancel Multiple Orders(TRADE)
[**CfuturesDeleteListenKeyV1**](V1API.md#CfuturesDeleteListenKeyV1) | **Delete** /dapi/v1/listenKey | Close User Data Stream(USER_STREAM)
[**CfuturesDeleteOrderV1**](V1API.md#CfuturesDeleteOrderV1) | **Delete** /dapi/v1/order | Cancel Order (TRADE)
[**CfuturesGetAccountV1**](V1API.md#CfuturesGetAccountV1) | **Get** /dapi/v1/account | Account Information (USER_DATA)
[**CfuturesGetAdlQuantileV1**](V1API.md#CfuturesGetAdlQuantileV1) | **Get** /dapi/v1/adlQuantile | Position ADL Quantile Estimation(USER_DATA)
[**CfuturesGetAggTradesV1**](V1API.md#CfuturesGetAggTradesV1) | **Get** /dapi/v1/aggTrades | Compressed/Aggregate Trades List
[**CfuturesGetAllOrdersV1**](V1API.md#CfuturesGetAllOrdersV1) | **Get** /dapi/v1/allOrders | All Orders (USER_DATA)
[**CfuturesGetBalanceV1**](V1API.md#CfuturesGetBalanceV1) | **Get** /dapi/v1/balance | Futures Account Balance (USER_DATA)
[**CfuturesGetCommissionRateV1**](V1API.md#CfuturesGetCommissionRateV1) | **Get** /dapi/v1/commissionRate | User Commission Rate (USER_DATA)
[**CfuturesGetConstituentsV1**](V1API.md#CfuturesGetConstituentsV1) | **Get** /dapi/v1/constituents | Query Index Price Constituents
[**CfuturesGetContinuousKlinesV1**](V1API.md#CfuturesGetContinuousKlinesV1) | **Get** /dapi/v1/continuousKlines | Continuous Contract Kline/Candlestick Data
[**CfuturesGetDepthV1**](V1API.md#CfuturesGetDepthV1) | **Get** /dapi/v1/depth | Order Book
[**CfuturesGetExchangeInfoV1**](V1API.md#CfuturesGetExchangeInfoV1) | **Get** /dapi/v1/exchangeInfo | Exchange Information
[**CfuturesGetForceOrdersV1**](V1API.md#CfuturesGetForceOrdersV1) | **Get** /dapi/v1/forceOrders | User&#39;s Force Orders(USER_DATA)
[**CfuturesGetFundingInfoV1**](V1API.md#CfuturesGetFundingInfoV1) | **Get** /dapi/v1/fundingInfo | Get Funding Rate Info
[**CfuturesGetFundingRateV1**](V1API.md#CfuturesGetFundingRateV1) | **Get** /dapi/v1/fundingRate | Get Funding Rate History of Perpetual Futures
[**CfuturesGetHistoricalTradesV1**](V1API.md#CfuturesGetHistoricalTradesV1) | **Get** /dapi/v1/historicalTrades | Old Trades Lookup(MARKET_DATA)
[**CfuturesGetIncomeAsynIdV1**](V1API.md#CfuturesGetIncomeAsynIdV1) | **Get** /dapi/v1/income/asyn/id | Get Futures Transaction History Download Link by Id (USER_DATA)
[**CfuturesGetIncomeAsynV1**](V1API.md#CfuturesGetIncomeAsynV1) | **Get** /dapi/v1/income/asyn | Get Download Id For Futures Transaction History(USER_DATA)
[**CfuturesGetIncomeV1**](V1API.md#CfuturesGetIncomeV1) | **Get** /dapi/v1/income | Get Income History(USER_DATA)
[**CfuturesGetIndexPriceKlinesV1**](V1API.md#CfuturesGetIndexPriceKlinesV1) | **Get** /dapi/v1/indexPriceKlines | Index Price Kline/Candlestick Data
[**CfuturesGetKlinesV1**](V1API.md#CfuturesGetKlinesV1) | **Get** /dapi/v1/klines | Kline/Candlestick Data
[**CfuturesGetLeverageBracketV1**](V1API.md#CfuturesGetLeverageBracketV1) | **Get** /dapi/v1/leverageBracket | Notional Bracket for Pair(USER_DATA)
[**CfuturesGetMarkPriceKlinesV1**](V1API.md#CfuturesGetMarkPriceKlinesV1) | **Get** /dapi/v1/markPriceKlines | Mark Price Kline/Candlestick Data
[**CfuturesGetOpenInterestV1**](V1API.md#CfuturesGetOpenInterestV1) | **Get** /dapi/v1/openInterest | Open Interest
[**CfuturesGetOpenOrderV1**](V1API.md#CfuturesGetOpenOrderV1) | **Get** /dapi/v1/openOrder | Query Current Open Order(USER_DATA)
[**CfuturesGetOpenOrdersV1**](V1API.md#CfuturesGetOpenOrdersV1) | **Get** /dapi/v1/openOrders | Current All Open Orders (USER_DATA)
[**CfuturesGetOrderAmendmentV1**](V1API.md#CfuturesGetOrderAmendmentV1) | **Get** /dapi/v1/orderAmendment | Get Order Modify History (USER_DATA)
[**CfuturesGetOrderAsynIdV1**](V1API.md#CfuturesGetOrderAsynIdV1) | **Get** /dapi/v1/order/asyn/id | Get Futures Order History Download Link by Id (USER_DATA)
[**CfuturesGetOrderAsynV1**](V1API.md#CfuturesGetOrderAsynV1) | **Get** /dapi/v1/order/asyn | Get Download Id For Futures Order History (USER_DATA)
[**CfuturesGetOrderV1**](V1API.md#CfuturesGetOrderV1) | **Get** /dapi/v1/order | Query Order (USER_DATA)
[**CfuturesGetPingV1**](V1API.md#CfuturesGetPingV1) | **Get** /dapi/v1/ping | Test Connectivity
[**CfuturesGetPmAccountInfoV1**](V1API.md#CfuturesGetPmAccountInfoV1) | **Get** /dapi/v1/pmAccountInfo | Classic Portfolio Margin Account Information (USER_DATA)
[**CfuturesGetPositionMarginHistoryV1**](V1API.md#CfuturesGetPositionMarginHistoryV1) | **Get** /dapi/v1/positionMargin/history | Get Position Margin Change History(TRADE)
[**CfuturesGetPositionRiskV1**](V1API.md#CfuturesGetPositionRiskV1) | **Get** /dapi/v1/positionRisk | Position Information(USER_DATA)
[**CfuturesGetPositionSideDualV1**](V1API.md#CfuturesGetPositionSideDualV1) | **Get** /dapi/v1/positionSide/dual | Get Current Position Mode(USER_DATA)
[**CfuturesGetPremiumIndexKlinesV1**](V1API.md#CfuturesGetPremiumIndexKlinesV1) | **Get** /dapi/v1/premiumIndexKlines | Premium index Kline Data
[**CfuturesGetPremiumIndexV1**](V1API.md#CfuturesGetPremiumIndexV1) | **Get** /dapi/v1/premiumIndex | Index Price and Mark Price
[**CfuturesGetTicker24hrV1**](V1API.md#CfuturesGetTicker24hrV1) | **Get** /dapi/v1/ticker/24hr | 24hr Ticker Price Change Statistics
[**CfuturesGetTickerBookTickerV1**](V1API.md#CfuturesGetTickerBookTickerV1) | **Get** /dapi/v1/ticker/bookTicker | Symbol Order Book Ticker
[**CfuturesGetTickerPriceV1**](V1API.md#CfuturesGetTickerPriceV1) | **Get** /dapi/v1/ticker/price | Symbol Price Ticker
[**CfuturesGetTimeV1**](V1API.md#CfuturesGetTimeV1) | **Get** /dapi/v1/time | Check Server time
[**CfuturesGetTradeAsynIdV1**](V1API.md#CfuturesGetTradeAsynIdV1) | **Get** /dapi/v1/trade/asyn/id | Get Futures Trade Download Link by Id(USER_DATA)
[**CfuturesGetTradeAsynV1**](V1API.md#CfuturesGetTradeAsynV1) | **Get** /dapi/v1/trade/asyn | Get Download Id For Futures Trade History (USER_DATA)
[**CfuturesGetTradesV1**](V1API.md#CfuturesGetTradesV1) | **Get** /dapi/v1/trades | Recent Trades List
[**CfuturesGetUserTradesV1**](V1API.md#CfuturesGetUserTradesV1) | **Get** /dapi/v1/userTrades | Account Trade List (USER_DATA)
[**CfuturesUpdateBatchOrdersV1**](V1API.md#CfuturesUpdateBatchOrdersV1) | **Put** /dapi/v1/batchOrders | Modify Multiple Orders(TRADE)
[**CfuturesUpdateListenKeyV1**](V1API.md#CfuturesUpdateListenKeyV1) | **Put** /dapi/v1/listenKey | Keepalive User Data Stream (USER_STREAM)
[**CfuturesUpdateOrderV1**](V1API.md#CfuturesUpdateOrderV1) | **Put** /dapi/v1/order | Modify Order (TRADE)



## CfuturesCreateBatchOrdersV1

> []CfuturesCreateBatchOrdersV1RespInner CfuturesCreateBatchOrdersV1(ctx).BatchOrders(batchOrders).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Place Multiple Orders(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	batchOrders := []openapiclient.CfuturesCreateBatchOrderV1ReqBatchOrdersItem{*openapiclient.NewCfuturesCreateBatchOrderV1ReqBatchOrdersItem("Side_example", "Symbol_example", "Type_example")} // []CfuturesCreateBatchOrderV1ReqBatchOrdersItem | 
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesCreateBatchOrdersV1(context.Background()).BatchOrders(batchOrders).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesCreateBatchOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesCreateBatchOrdersV1`: []CfuturesCreateBatchOrdersV1RespInner
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesCreateBatchOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesCreateBatchOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchOrders** | [**[]CfuturesCreateBatchOrderV1ReqBatchOrdersItem**](CfuturesCreateBatchOrderV1ReqBatchOrdersItem.md) |  | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]CfuturesCreateBatchOrdersV1RespInner**](CfuturesCreateBatchOrdersV1RespInner.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesCreateCountdownCancelAllV1

> CfuturesCreateCountdownCancelAllV1Resp CfuturesCreateCountdownCancelAllV1(ctx).CountdownTime(countdownTime).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Auto-Cancel All Open Orders (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	countdownTime := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesCreateCountdownCancelAllV1(context.Background()).CountdownTime(countdownTime).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesCreateCountdownCancelAllV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesCreateCountdownCancelAllV1`: CfuturesCreateCountdownCancelAllV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesCreateCountdownCancelAllV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesCreateCountdownCancelAllV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countdownTime** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CfuturesCreateCountdownCancelAllV1Resp**](CfuturesCreateCountdownCancelAllV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesCreateLeverageV1

> CfuturesCreateLeverageV1Resp CfuturesCreateLeverageV1(ctx).Leverage(leverage).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Change Initial Leverage (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	leverage := int32(56) // int32 | 
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesCreateLeverageV1(context.Background()).Leverage(leverage).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesCreateLeverageV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesCreateLeverageV1`: CfuturesCreateLeverageV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesCreateLeverageV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesCreateLeverageV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **leverage** | **int32** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CfuturesCreateLeverageV1Resp**](CfuturesCreateLeverageV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesCreateListenKeyV1

> CfuturesCreateListenKeyV1Resp CfuturesCreateListenKeyV1(ctx).Execute()

Start User Data Stream (USER_STREAM)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesCreateListenKeyV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesCreateListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesCreateListenKeyV1`: CfuturesCreateListenKeyV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesCreateListenKeyV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesCreateListenKeyV1Request struct via the builder pattern


### Return type

[**CfuturesCreateListenKeyV1Resp**](CfuturesCreateListenKeyV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesCreateMarginTypeV1

> CfuturesCreateMarginTypeV1Resp CfuturesCreateMarginTypeV1(ctx).MarginType(marginType).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Change Margin Type (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	marginType := "marginType_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesCreateMarginTypeV1(context.Background()).MarginType(marginType).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesCreateMarginTypeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesCreateMarginTypeV1`: CfuturesCreateMarginTypeV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesCreateMarginTypeV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesCreateMarginTypeV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marginType** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CfuturesCreateMarginTypeV1Resp**](CfuturesCreateMarginTypeV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesCreateOrderV1

> CfuturesCreateOrderV1Resp CfuturesCreateOrderV1(ctx).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).ActivationPrice(activationPrice).CallbackRate(callbackRate).ClosePosition(closePosition).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).PositionSide(positionSide).Price(price).PriceMatch(priceMatch).PriceProtect(priceProtect).Quantity(quantity).RecvWindow(recvWindow).ReduceOnly(reduceOnly).SelfTradePreventionMode(selfTradePreventionMode).StopPrice(stopPrice).TimeInForce(timeInForce).WorkingType(workingType).Execute()

New Order (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	side := "side_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	type_ := "type__example" // string |  (default to "")
	activationPrice := "activationPrice_example" // string |  (optional) (default to "")
	callbackRate := "callbackRate_example" // string |  (optional) (default to "")
	closePosition := "closePosition_example" // string |  (optional) (default to "")
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
	resp, r, err := apiClient.V1API.CfuturesCreateOrderV1(context.Background()).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).ActivationPrice(activationPrice).CallbackRate(callbackRate).ClosePosition(closePosition).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).PositionSide(positionSide).Price(price).PriceMatch(priceMatch).PriceProtect(priceProtect).Quantity(quantity).RecvWindow(recvWindow).ReduceOnly(reduceOnly).SelfTradePreventionMode(selfTradePreventionMode).StopPrice(stopPrice).TimeInForce(timeInForce).WorkingType(workingType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesCreateOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesCreateOrderV1`: CfuturesCreateOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesCreateOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesCreateOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **side** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **type_** | **string** |  | [default to &quot;&quot;]
 **activationPrice** | **string** |  | [default to &quot;&quot;]
 **callbackRate** | **string** |  | [default to &quot;&quot;]
 **closePosition** | **string** |  | [default to &quot;&quot;]
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

[**CfuturesCreateOrderV1Resp**](CfuturesCreateOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesCreatePositionMarginV1

> CfuturesCreatePositionMarginV1Resp CfuturesCreatePositionMarginV1(ctx).Amount(amount).Symbol(symbol).Timestamp(timestamp).Type_(type_).PositionSide(positionSide).RecvWindow(recvWindow).Execute()

Modify Isolated Position Margin(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
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
	resp, r, err := apiClient.V1API.CfuturesCreatePositionMarginV1(context.Background()).Amount(amount).Symbol(symbol).Timestamp(timestamp).Type_(type_).PositionSide(positionSide).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesCreatePositionMarginV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesCreatePositionMarginV1`: CfuturesCreatePositionMarginV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesCreatePositionMarginV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesCreatePositionMarginV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **type_** | **int32** |  | 
 **positionSide** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**CfuturesCreatePositionMarginV1Resp**](CfuturesCreatePositionMarginV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesCreatePositionSideDualV1

> CfuturesCreatePositionSideDualV1Resp CfuturesCreatePositionSideDualV1(ctx).DualSidePosition(dualSidePosition).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Change Position Mode(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	dualSidePosition := "dualSidePosition_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesCreatePositionSideDualV1(context.Background()).DualSidePosition(dualSidePosition).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesCreatePositionSideDualV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesCreatePositionSideDualV1`: CfuturesCreatePositionSideDualV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesCreatePositionSideDualV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesCreatePositionSideDualV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dualSidePosition** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CfuturesCreatePositionSideDualV1Resp**](CfuturesCreatePositionSideDualV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesDeleteAllOpenOrdersV1

> CfuturesDeleteAllOpenOrdersV1Resp CfuturesDeleteAllOpenOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Cancel All Open Orders(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesDeleteAllOpenOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesDeleteAllOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesDeleteAllOpenOrdersV1`: CfuturesDeleteAllOpenOrdersV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesDeleteAllOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesDeleteAllOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CfuturesDeleteAllOpenOrdersV1Resp**](CfuturesDeleteAllOpenOrdersV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesDeleteBatchOrdersV1

> []CfuturesDeleteBatchOrdersV1RespInner CfuturesDeleteBatchOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderIdList(orderIdList).OrigClientOrderIdList(origClientOrderIdList).RecvWindow(recvWindow).Execute()

Cancel Multiple Orders(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	orderIdList := []int64{int64(123)} // []int64 | max length 10 <br/> e.g. [1234567,2345678] (optional)
	origClientOrderIdList := []string{"Inner_example"} // []string | max length 10<br/> e.g. [&#34;my_id_1&#34;,&#34;my_id_2&#34;], encode the double quotes. No space after comma. (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesDeleteBatchOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderIdList(orderIdList).OrigClientOrderIdList(origClientOrderIdList).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesDeleteBatchOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesDeleteBatchOrdersV1`: []CfuturesDeleteBatchOrdersV1RespInner
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesDeleteBatchOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesDeleteBatchOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderIdList** | **[]int64** | max length 10 &lt;br/&gt; e.g. [1234567,2345678] | 
 **origClientOrderIdList** | **[]string** | max length 10&lt;br/&gt; e.g. [&amp;#34;my_id_1&amp;#34;,&amp;#34;my_id_2&amp;#34;], encode the double quotes. No space after comma. | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]CfuturesDeleteBatchOrdersV1RespInner**](CfuturesDeleteBatchOrdersV1RespInner.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesDeleteListenKeyV1

> map[string]interface{} CfuturesDeleteListenKeyV1(ctx).Execute()

Close User Data Stream(USER_STREAM)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesDeleteListenKeyV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesDeleteListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesDeleteListenKeyV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesDeleteListenKeyV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesDeleteListenKeyV1Request struct via the builder pattern


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


## CfuturesDeleteOrderV1

> CfuturesDeleteOrderV1Resp CfuturesDeleteOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Cancel Order (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesDeleteOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesDeleteOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesDeleteOrderV1`: CfuturesDeleteOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesDeleteOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesDeleteOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**CfuturesDeleteOrderV1Resp**](CfuturesDeleteOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetAccountV1

> CfuturesGetAccountV1Resp CfuturesGetAccountV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Account Information (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetAccountV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetAccountV1`: CfuturesGetAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CfuturesGetAccountV1Resp**](CfuturesGetAccountV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetAdlQuantileV1

> []CfuturesGetAdlQuantileV1RespItem CfuturesGetAdlQuantileV1(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Position ADL Quantile Estimation(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetAdlQuantileV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetAdlQuantileV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetAdlQuantileV1`: []CfuturesGetAdlQuantileV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetAdlQuantileV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetAdlQuantileV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]CfuturesGetAdlQuantileV1RespItem**](CfuturesGetAdlQuantileV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetAggTradesV1

> []CfuturesGetAggTradesV1RespItem CfuturesGetAggTradesV1(ctx).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Compressed/Aggregate Trades List



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	fromId := int64(789) // int64 | ID to get aggregate trades from INCLUSIVE. (optional)
	startTime := int64(789) // int64 | Timestamp in ms to get aggregate trades from INCLUSIVE. (optional)
	endTime := int64(789) // int64 | Timestamp in ms to get aggregate trades until INCLUSIVE. (optional)
	limit := int32(56) // int32 | Default 500; max 1000. (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetAggTradesV1(context.Background()).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetAggTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetAggTradesV1`: []CfuturesGetAggTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetAggTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetAggTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **fromId** | **int64** | ID to get aggregate trades from INCLUSIVE. | 
 **startTime** | **int64** | Timestamp in ms to get aggregate trades from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get aggregate trades until INCLUSIVE. | 
 **limit** | **int32** | Default 500; max 1000. | [default to 500]

### Return type

[**[]CfuturesGetAggTradesV1RespItem**](CfuturesGetAggTradesV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetAllOrdersV1

> []CfuturesGetAllOrdersV1RespItem CfuturesGetAllOrdersV1(ctx).Timestamp(timestamp).Symbol(symbol).Pair(pair).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

All Orders (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	pair := "pair_example" // string |  (optional) (default to "")
	orderId := int64(789) // int64 |  (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 50; max 100. (optional) (default to 50)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetAllOrdersV1(context.Background()).Timestamp(timestamp).Symbol(symbol).Pair(pair).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetAllOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetAllOrdersV1`: []CfuturesGetAllOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetAllOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetAllOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **pair** | **string** |  | [default to &quot;&quot;]
 **orderId** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 50; max 100. | [default to 50]
 **recvWindow** | **int64** |  | 

### Return type

[**[]CfuturesGetAllOrdersV1RespItem**](CfuturesGetAllOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetBalanceV1

> []CfuturesGetBalanceV1RespItem CfuturesGetBalanceV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Futures Account Balance (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetBalanceV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetBalanceV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetBalanceV1`: []CfuturesGetBalanceV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetBalanceV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetBalanceV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]CfuturesGetBalanceV1RespItem**](CfuturesGetBalanceV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetCommissionRateV1

> CfuturesGetCommissionRateV1Resp CfuturesGetCommissionRateV1(ctx).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

User Commission Rate (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetCommissionRateV1(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetCommissionRateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetCommissionRateV1`: CfuturesGetCommissionRateV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetCommissionRateV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetCommissionRateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CfuturesGetCommissionRateV1Resp**](CfuturesGetCommissionRateV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetConstituentsV1

> CfuturesGetConstituentsV1Resp CfuturesGetConstituentsV1(ctx).Symbol(symbol).Execute()

Query Index Price Constituents



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetConstituentsV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetConstituentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetConstituentsV1`: CfuturesGetConstituentsV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetConstituentsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetConstituentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**CfuturesGetConstituentsV1Resp**](CfuturesGetConstituentsV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetContinuousKlinesV1

> [][]CfuturesGetContinuousKlinesV1RespInnerInner CfuturesGetContinuousKlinesV1(ctx).Pair(pair).ContractType(contractType).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Continuous Contract Kline/Candlestick Data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	pair := "pair_example" // string |  (default to "")
	contractType := "contractType_example" // string |  (default to "")
	interval := "interval_example" // string |  (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 500; max 1500. (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetContinuousKlinesV1(context.Background()).Pair(pair).ContractType(contractType).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetContinuousKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetContinuousKlinesV1`: [][]CfuturesGetContinuousKlinesV1RespInnerInner
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetContinuousKlinesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetContinuousKlinesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pair** | **string** |  | [default to &quot;&quot;]
 **contractType** | **string** |  | [default to &quot;&quot;]
 **interval** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 1500. | [default to 500]

### Return type

[**[][]CfuturesGetContinuousKlinesV1RespInnerInner**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetDepthV1

> CfuturesGetDepthV1Resp CfuturesGetDepthV1(ctx).Symbol(symbol).Limit(limit).Execute()

Order Book



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	limit := int32(56) // int32 | Default 500; Valid limits:[5, 10, 20, 50, 100, 500, 1000] (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetDepthV1(context.Background()).Symbol(symbol).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetDepthV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetDepthV1`: CfuturesGetDepthV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetDepthV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetDepthV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **limit** | **int32** | Default 500; Valid limits:[5, 10, 20, 50, 100, 500, 1000] | [default to 500]

### Return type

[**CfuturesGetDepthV1Resp**](CfuturesGetDepthV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetExchangeInfoV1

> CfuturesGetExchangeInfoV1Resp CfuturesGetExchangeInfoV1(ctx).Execute()

Exchange Information



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetExchangeInfoV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetExchangeInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetExchangeInfoV1`: CfuturesGetExchangeInfoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetExchangeInfoV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetExchangeInfoV1Request struct via the builder pattern


### Return type

[**CfuturesGetExchangeInfoV1Resp**](CfuturesGetExchangeInfoV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetForceOrdersV1

> []CfuturesGetForceOrdersV1RespItem CfuturesGetForceOrdersV1(ctx).Timestamp(timestamp).Symbol(symbol).AutoCloseType(autoCloseType).RecvWindow(recvWindow).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

User's Force Orders(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	autoCloseType := "autoCloseType_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetForceOrdersV1(context.Background()).Timestamp(timestamp).Symbol(symbol).AutoCloseType(autoCloseType).RecvWindow(recvWindow).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetForceOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetForceOrdersV1`: []CfuturesGetForceOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetForceOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetForceOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **autoCloseType** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **limit** | **int32** |  | [default to 50]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**[]CfuturesGetForceOrdersV1RespItem**](CfuturesGetForceOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetFundingInfoV1

> []CfuturesGetFundingInfoV1RespItem CfuturesGetFundingInfoV1(ctx).Execute()

Get Funding Rate Info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetFundingInfoV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetFundingInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetFundingInfoV1`: []CfuturesGetFundingInfoV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetFundingInfoV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetFundingInfoV1Request struct via the builder pattern


### Return type

[**[]CfuturesGetFundingInfoV1RespItem**](CfuturesGetFundingInfoV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetFundingRateV1

> []CfuturesGetFundingRateV1RespItem CfuturesGetFundingRateV1(ctx).Symbol(symbol).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Get Funding Rate History of Perpetual Futures



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	startTime := int64(789) // int64 | Timestamp in ms to get funding rate from INCLUSIVE. (optional)
	endTime := int64(789) // int64 | Timestamp in ms to get funding rate  until INCLUSIVE. (optional)
	limit := int32(56) // int32 | Default 100; max 1000 (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetFundingRateV1(context.Background()).Symbol(symbol).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetFundingRateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetFundingRateV1`: []CfuturesGetFundingRateV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetFundingRateV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetFundingRateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** | Timestamp in ms to get funding rate from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get funding rate  until INCLUSIVE. | 
 **limit** | **int32** | Default 100; max 1000 | [default to 100]

### Return type

[**[]CfuturesGetFundingRateV1RespItem**](CfuturesGetFundingRateV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetHistoricalTradesV1

> []CfuturesGetHistoricalTradesV1RespItem CfuturesGetHistoricalTradesV1(ctx).Symbol(symbol).Limit(limit).FromId(fromId).Execute()

Old Trades Lookup(MARKET_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	limit := int32(56) // int32 | Default 100; max 500. (optional) (default to 100)
	fromId := int64(789) // int64 | TradeId to fetch from. Default gets most recent trades. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetHistoricalTradesV1(context.Background()).Symbol(symbol).Limit(limit).FromId(fromId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetHistoricalTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetHistoricalTradesV1`: []CfuturesGetHistoricalTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetHistoricalTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetHistoricalTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **limit** | **int32** | Default 100; max 500. | [default to 100]
 **fromId** | **int64** | TradeId to fetch from. Default gets most recent trades. | 

### Return type

[**[]CfuturesGetHistoricalTradesV1RespItem**](CfuturesGetHistoricalTradesV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetIncomeAsynIdV1

> CfuturesGetIncomeAsynIdV1Resp CfuturesGetIncomeAsynIdV1(ctx).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Futures Transaction History Download Link by Id (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	downloadId := "downloadId_example" // string | get by download id api (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetIncomeAsynIdV1(context.Background()).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetIncomeAsynIdV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetIncomeAsynIdV1`: CfuturesGetIncomeAsynIdV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetIncomeAsynIdV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetIncomeAsynIdV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadId** | **string** | get by download id api | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CfuturesGetIncomeAsynIdV1Resp**](CfuturesGetIncomeAsynIdV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetIncomeAsynV1

> CfuturesGetIncomeAsynV1Resp CfuturesGetIncomeAsynV1(ctx).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Download Id For Futures Transaction History(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	startTime := int64(789) // int64 | Timestamp in ms
	endTime := int64(789) // int64 | Timestamp in ms
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetIncomeAsynV1(context.Background()).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetIncomeAsynV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetIncomeAsynV1`: CfuturesGetIncomeAsynV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetIncomeAsynV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetIncomeAsynV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** | Timestamp in ms | 
 **endTime** | **int64** | Timestamp in ms | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CfuturesGetIncomeAsynV1Resp**](CfuturesGetIncomeAsynV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetIncomeV1

> []CfuturesGetIncomeV1RespItem CfuturesGetIncomeV1(ctx).Timestamp(timestamp).Symbol(symbol).IncomeType(incomeType).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Get Income History(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	incomeType := "incomeType_example" // string | &#34;TRANSFER&#34;,&#34;WELCOME_BONUS&#34;, &#34;FUNDING_FEE&#34;, &#34;REALIZED_PNL&#34;, &#34;COMMISSION&#34;, &#34;INSURANCE_CLEAR&#34;, and &#34;DELIVERED_SETTELMENT&#34; (optional) (default to "")
	startTime := int64(789) // int64 | Timestamp in ms to get funding from INCLUSIVE. (optional)
	endTime := int64(789) // int64 | Timestamp in ms to get funding until INCLUSIVE. (optional)
	page := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 | Default 100; max 1000 (optional) (default to 100)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetIncomeV1(context.Background()).Timestamp(timestamp).Symbol(symbol).IncomeType(incomeType).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetIncomeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetIncomeV1`: []CfuturesGetIncomeV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetIncomeV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetIncomeV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **incomeType** | **string** | &amp;#34;TRANSFER&amp;#34;,&amp;#34;WELCOME_BONUS&amp;#34;, &amp;#34;FUNDING_FEE&amp;#34;, &amp;#34;REALIZED_PNL&amp;#34;, &amp;#34;COMMISSION&amp;#34;, &amp;#34;INSURANCE_CLEAR&amp;#34;, and &amp;#34;DELIVERED_SETTELMENT&amp;#34; | [default to &quot;&quot;]
 **startTime** | **int64** | Timestamp in ms to get funding from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get funding until INCLUSIVE. | 
 **page** | **int32** |  | 
 **limit** | **int32** | Default 100; max 1000 | [default to 100]
 **recvWindow** | **int64** |  | 

### Return type

[**[]CfuturesGetIncomeV1RespItem**](CfuturesGetIncomeV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetIndexPriceKlinesV1

> [][]CfuturesGetContinuousKlinesV1RespInnerInner CfuturesGetIndexPriceKlinesV1(ctx).Pair(pair).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Index Price Kline/Candlestick Data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	pair := "pair_example" // string |  (default to "")
	interval := "interval_example" // string |  (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 500; max 1500. (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetIndexPriceKlinesV1(context.Background()).Pair(pair).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetIndexPriceKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetIndexPriceKlinesV1`: [][]CfuturesGetContinuousKlinesV1RespInnerInner
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetIndexPriceKlinesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetIndexPriceKlinesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pair** | **string** |  | [default to &quot;&quot;]
 **interval** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 1500. | [default to 500]

### Return type

[**[][]CfuturesGetContinuousKlinesV1RespInnerInner**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetKlinesV1

> [][]CfuturesGetContinuousKlinesV1RespInnerInner CfuturesGetKlinesV1(ctx).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Kline/Candlestick Data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	interval := "interval_example" // string |  (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 500; max 1500. (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetKlinesV1(context.Background()).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetKlinesV1`: [][]CfuturesGetContinuousKlinesV1RespInnerInner
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetKlinesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetKlinesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **interval** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 1500. | [default to 500]

### Return type

[**[][]CfuturesGetContinuousKlinesV1RespInnerInner**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetLeverageBracketV1

> []CfuturesGetLeverageBracketV1RespItem CfuturesGetLeverageBracketV1(ctx).Timestamp(timestamp).Pair(pair).RecvWindow(recvWindow).Execute()

Notional Bracket for Pair(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	pair := "pair_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetLeverageBracketV1(context.Background()).Timestamp(timestamp).Pair(pair).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetLeverageBracketV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetLeverageBracketV1`: []CfuturesGetLeverageBracketV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetLeverageBracketV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetLeverageBracketV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **pair** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]CfuturesGetLeverageBracketV1RespItem**](CfuturesGetLeverageBracketV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetMarkPriceKlinesV1

> [][]CfuturesGetContinuousKlinesV1RespInnerInner CfuturesGetMarkPriceKlinesV1(ctx).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Mark Price Kline/Candlestick Data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	interval := "interval_example" // string |  (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 500; max 1500. (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetMarkPriceKlinesV1(context.Background()).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetMarkPriceKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetMarkPriceKlinesV1`: [][]CfuturesGetContinuousKlinesV1RespInnerInner
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetMarkPriceKlinesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetMarkPriceKlinesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **interval** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 1500. | [default to 500]

### Return type

[**[][]CfuturesGetContinuousKlinesV1RespInnerInner**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetOpenInterestV1

> CfuturesGetOpenInterestV1Resp CfuturesGetOpenInterestV1(ctx).Symbol(symbol).Execute()

Open Interest



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetOpenInterestV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetOpenInterestV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetOpenInterestV1`: CfuturesGetOpenInterestV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetOpenInterestV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetOpenInterestV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**CfuturesGetOpenInterestV1Resp**](CfuturesGetOpenInterestV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetOpenOrderV1

> CfuturesGetOpenOrderV1Resp CfuturesGetOpenOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Current Open Order(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetOpenOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetOpenOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetOpenOrderV1`: CfuturesGetOpenOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetOpenOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetOpenOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**CfuturesGetOpenOrderV1Resp**](CfuturesGetOpenOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetOpenOrdersV1

> []CfuturesGetOpenOrdersV1RespItem CfuturesGetOpenOrdersV1(ctx).Timestamp(timestamp).Symbol(symbol).Pair(pair).RecvWindow(recvWindow).Execute()

Current All Open Orders (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	pair := "pair_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetOpenOrdersV1(context.Background()).Timestamp(timestamp).Symbol(symbol).Pair(pair).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetOpenOrdersV1`: []CfuturesGetOpenOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **pair** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]CfuturesGetOpenOrdersV1RespItem**](CfuturesGetOpenOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetOrderAmendmentV1

> []CfuturesGetOrderAmendmentV1RespItem CfuturesGetOrderAmendmentV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Get Order Modify History (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
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
	resp, r, err := apiClient.V1API.CfuturesGetOrderAmendmentV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetOrderAmendmentV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetOrderAmendmentV1`: []CfuturesGetOrderAmendmentV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetOrderAmendmentV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetOrderAmendmentV1Request struct via the builder pattern


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

[**[]CfuturesGetOrderAmendmentV1RespItem**](CfuturesGetOrderAmendmentV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetOrderAsynIdV1

> CfuturesGetOrderAsynIdV1Resp CfuturesGetOrderAsynIdV1(ctx).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Futures Order History Download Link by Id (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	downloadId := "downloadId_example" // string | get by download id api (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetOrderAsynIdV1(context.Background()).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetOrderAsynIdV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetOrderAsynIdV1`: CfuturesGetOrderAsynIdV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetOrderAsynIdV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetOrderAsynIdV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadId** | **string** | get by download id api | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CfuturesGetOrderAsynIdV1Resp**](CfuturesGetOrderAsynIdV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetOrderAsynV1

> CfuturesGetOrderAsynV1Resp CfuturesGetOrderAsynV1(ctx).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Download Id For Futures Order History (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	startTime := int64(789) // int64 | Timestamp in ms
	endTime := int64(789) // int64 | Timestamp in ms
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetOrderAsynV1(context.Background()).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetOrderAsynV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetOrderAsynV1`: CfuturesGetOrderAsynV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetOrderAsynV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetOrderAsynV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** | Timestamp in ms | 
 **endTime** | **int64** | Timestamp in ms | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CfuturesGetOrderAsynV1Resp**](CfuturesGetOrderAsynV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetOrderV1

> CfuturesGetOrderV1Resp CfuturesGetOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Order (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetOrderV1`: CfuturesGetOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**CfuturesGetOrderV1Resp**](CfuturesGetOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetPingV1

> map[string]interface{} CfuturesGetPingV1(ctx).Execute()

Test Connectivity



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetPingV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetPingV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetPingV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetPingV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetPingV1Request struct via the builder pattern


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


## CfuturesGetPmAccountInfoV1

> CfuturesGetPmAccountInfoV1Resp CfuturesGetPmAccountInfoV1(ctx).Asset(asset).RecvWindow(recvWindow).Execute()

Classic Portfolio Margin Account Information (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	asset := "asset_example" // string |  (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetPmAccountInfoV1(context.Background()).Asset(asset).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetPmAccountInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetPmAccountInfoV1`: CfuturesGetPmAccountInfoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetPmAccountInfoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetPmAccountInfoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**CfuturesGetPmAccountInfoV1Resp**](CfuturesGetPmAccountInfoV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetPositionMarginHistoryV1

> []CfuturesGetPositionMarginHistoryV1RespItem CfuturesGetPositionMarginHistoryV1(ctx).Symbol(symbol).Timestamp(timestamp).Type_(type_).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Get Position Margin Change History(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	type_ := int32(56) // int32 | 1: Add position margin,2: Reduce position margin (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default: 50 (optional) (default to 50)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetPositionMarginHistoryV1(context.Background()).Symbol(symbol).Timestamp(timestamp).Type_(type_).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetPositionMarginHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetPositionMarginHistoryV1`: []CfuturesGetPositionMarginHistoryV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetPositionMarginHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetPositionMarginHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **type_** | **int32** | 1: Add position margin,2: Reduce position margin | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default: 50 | [default to 50]
 **recvWindow** | **int64** |  | 

### Return type

[**[]CfuturesGetPositionMarginHistoryV1RespItem**](CfuturesGetPositionMarginHistoryV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetPositionRiskV1

> []CfuturesGetPositionRiskV1RespItem CfuturesGetPositionRiskV1(ctx).Timestamp(timestamp).MarginAsset(marginAsset).Pair(pair).RecvWindow(recvWindow).Execute()

Position Information(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	marginAsset := "marginAsset_example" // string |  (optional) (default to "")
	pair := "pair_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetPositionRiskV1(context.Background()).Timestamp(timestamp).MarginAsset(marginAsset).Pair(pair).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetPositionRiskV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetPositionRiskV1`: []CfuturesGetPositionRiskV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetPositionRiskV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetPositionRiskV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **marginAsset** | **string** |  | [default to &quot;&quot;]
 **pair** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]CfuturesGetPositionRiskV1RespItem**](CfuturesGetPositionRiskV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetPositionSideDualV1

> CfuturesGetPositionSideDualV1Resp CfuturesGetPositionSideDualV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Current Position Mode(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetPositionSideDualV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetPositionSideDualV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetPositionSideDualV1`: CfuturesGetPositionSideDualV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetPositionSideDualV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetPositionSideDualV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CfuturesGetPositionSideDualV1Resp**](CfuturesGetPositionSideDualV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetPremiumIndexKlinesV1

> [][]CfuturesGetContinuousKlinesV1RespInnerInner CfuturesGetPremiumIndexKlinesV1(ctx).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Premium index Kline Data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	interval := "interval_example" // string |  (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 500; max 1500. (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetPremiumIndexKlinesV1(context.Background()).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetPremiumIndexKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetPremiumIndexKlinesV1`: [][]CfuturesGetContinuousKlinesV1RespInnerInner
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetPremiumIndexKlinesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetPremiumIndexKlinesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **interval** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 1500. | [default to 500]

### Return type

[**[][]CfuturesGetContinuousKlinesV1RespInnerInner**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetPremiumIndexV1

> []CfuturesGetPremiumIndexV1RespItem CfuturesGetPremiumIndexV1(ctx).Symbol(symbol).Pair(pair).Execute()

Index Price and Mark Price



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	symbol := "symbol_example" // string |  (optional) (default to "")
	pair := "pair_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetPremiumIndexV1(context.Background()).Symbol(symbol).Pair(pair).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetPremiumIndexV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetPremiumIndexV1`: []CfuturesGetPremiumIndexV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetPremiumIndexV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetPremiumIndexV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **pair** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]CfuturesGetPremiumIndexV1RespItem**](CfuturesGetPremiumIndexV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetTicker24hrV1

> []CfuturesGetTicker24hrV1RespItem CfuturesGetTicker24hrV1(ctx).Symbol(symbol).Pair(pair).Execute()

24hr Ticker Price Change Statistics



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	symbol := "symbol_example" // string |  (optional) (default to "")
	pair := "pair_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetTicker24hrV1(context.Background()).Symbol(symbol).Pair(pair).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetTicker24hrV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetTicker24hrV1`: []CfuturesGetTicker24hrV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetTicker24hrV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetTicker24hrV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **pair** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]CfuturesGetTicker24hrV1RespItem**](CfuturesGetTicker24hrV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetTickerBookTickerV1

> []CfuturesGetTickerBookTickerV1RespItem CfuturesGetTickerBookTickerV1(ctx).Symbol(symbol).Pair(pair).Execute()

Symbol Order Book Ticker



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	symbol := "symbol_example" // string |  (optional) (default to "")
	pair := "pair_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetTickerBookTickerV1(context.Background()).Symbol(symbol).Pair(pair).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetTickerBookTickerV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetTickerBookTickerV1`: []CfuturesGetTickerBookTickerV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetTickerBookTickerV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetTickerBookTickerV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **pair** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]CfuturesGetTickerBookTickerV1RespItem**](CfuturesGetTickerBookTickerV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetTickerPriceV1

> []CfuturesGetTickerPriceV1RespItem CfuturesGetTickerPriceV1(ctx).Symbol(symbol).Pair(pair).Execute()

Symbol Price Ticker



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	symbol := "symbol_example" // string |  (optional) (default to "")
	pair := "pair_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetTickerPriceV1(context.Background()).Symbol(symbol).Pair(pair).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetTickerPriceV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetTickerPriceV1`: []CfuturesGetTickerPriceV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetTickerPriceV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetTickerPriceV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **pair** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]CfuturesGetTickerPriceV1RespItem**](CfuturesGetTickerPriceV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetTimeV1

> CfuturesGetTimeV1Resp CfuturesGetTimeV1(ctx).Execute()

Check Server time



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetTimeV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetTimeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetTimeV1`: CfuturesGetTimeV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetTimeV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetTimeV1Request struct via the builder pattern


### Return type

[**CfuturesGetTimeV1Resp**](CfuturesGetTimeV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetTradeAsynIdV1

> CfuturesGetTradeAsynIdV1Resp CfuturesGetTradeAsynIdV1(ctx).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Futures Trade Download Link by Id(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	downloadId := "downloadId_example" // string | get by download id api (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetTradeAsynIdV1(context.Background()).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetTradeAsynIdV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetTradeAsynIdV1`: CfuturesGetTradeAsynIdV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetTradeAsynIdV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetTradeAsynIdV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadId** | **string** | get by download id api | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CfuturesGetTradeAsynIdV1Resp**](CfuturesGetTradeAsynIdV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetTradeAsynV1

> CfuturesGetTradeAsynV1Resp CfuturesGetTradeAsynV1(ctx).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Download Id For Futures Trade History (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	startTime := int64(789) // int64 | Timestamp in ms
	endTime := int64(789) // int64 | Timestamp in ms
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetTradeAsynV1(context.Background()).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetTradeAsynV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetTradeAsynV1`: CfuturesGetTradeAsynV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetTradeAsynV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetTradeAsynV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** | Timestamp in ms | 
 **endTime** | **int64** | Timestamp in ms | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CfuturesGetTradeAsynV1Resp**](CfuturesGetTradeAsynV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetTradesV1

> []CfuturesGetTradesV1RespItem CfuturesGetTradesV1(ctx).Symbol(symbol).Limit(limit).Execute()

Recent Trades List



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	limit := int32(56) // int32 | Default 500; max 1000. (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetTradesV1(context.Background()).Symbol(symbol).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetTradesV1`: []CfuturesGetTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **limit** | **int32** | Default 500; max 1000. | [default to 500]

### Return type

[**[]CfuturesGetTradesV1RespItem**](CfuturesGetTradesV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetUserTradesV1

> []CfuturesGetUserTradesV1RespItem CfuturesGetUserTradesV1(ctx).Timestamp(timestamp).Symbol(symbol).Pair(pair).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()

Account Trade List (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	pair := "pair_example" // string |  (optional) (default to "")
	orderId := "orderId_example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	fromId := int64(789) // int64 | Trade id to fetch from. Default gets most recent trades. (optional)
	limit := int32(56) // int32 | Default 50; max 1000 (optional) (default to 50)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesGetUserTradesV1(context.Background()).Timestamp(timestamp).Symbol(symbol).Pair(pair).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesGetUserTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetUserTradesV1`: []CfuturesGetUserTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesGetUserTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetUserTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **pair** | **string** |  | [default to &quot;&quot;]
 **orderId** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **fromId** | **int64** | Trade id to fetch from. Default gets most recent trades. | 
 **limit** | **int32** | Default 50; max 1000 | [default to 50]
 **recvWindow** | **int64** |  | 

### Return type

[**[]CfuturesGetUserTradesV1RespItem**](CfuturesGetUserTradesV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesUpdateBatchOrdersV1

> []CfuturesUpdateBatchOrdersV1RespInner CfuturesUpdateBatchOrdersV1(ctx).BatchOrders(batchOrders).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Modify Multiple Orders(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	batchOrders := map[string]interface{}{ ... } // map[string]interface{} | 
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesUpdateBatchOrdersV1(context.Background()).BatchOrders(batchOrders).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesUpdateBatchOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesUpdateBatchOrdersV1`: []CfuturesUpdateBatchOrdersV1RespInner
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesUpdateBatchOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesUpdateBatchOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchOrders** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]CfuturesUpdateBatchOrdersV1RespInner**](CfuturesUpdateBatchOrdersV1RespInner.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesUpdateListenKeyV1

> map[string]interface{} CfuturesUpdateListenKeyV1(ctx).Execute()

Keepalive User Data Stream (USER_STREAM)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesUpdateListenKeyV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesUpdateListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesUpdateListenKeyV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesUpdateListenKeyV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesUpdateListenKeyV1Request struct via the builder pattern


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


## CfuturesUpdateOrderV1

> CfuturesUpdateOrderV1Resp CfuturesUpdateOrderV1(ctx).Side(side).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).Price(price).PriceMatch(priceMatch).Quantity(quantity).RecvWindow(recvWindow).Execute()

Modify Order (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	side := "side_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	price := "price_example" // string |  (optional) (default to "")
	priceMatch := "priceMatch_example" // string |  (optional) (default to "")
	quantity := "quantity_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CfuturesUpdateOrderV1(context.Background()).Side(side).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).Price(price).PriceMatch(priceMatch).Quantity(quantity).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CfuturesUpdateOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesUpdateOrderV1`: CfuturesUpdateOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.CfuturesUpdateOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesUpdateOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **side** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **price** | **string** |  | [default to &quot;&quot;]
 **priceMatch** | **string** |  | [default to &quot;&quot;]
 **quantity** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**CfuturesUpdateOrderV1Resp**](CfuturesUpdateOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

