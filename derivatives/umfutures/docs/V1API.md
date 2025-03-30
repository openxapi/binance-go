# \V1API

All URIs are relative to *https://fapi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UfuturesCreateBatchOrdersV1**](V1API.md#UfuturesCreateBatchOrdersV1) | **Post** /fapi/v1/batchOrders | Place Multiple Orders(TRADE)
[**UfuturesCreateConvertAcceptQuoteV1**](V1API.md#UfuturesCreateConvertAcceptQuoteV1) | **Post** /fapi/v1/convert/acceptQuote | Accept the offered quote (USER_DATA)
[**UfuturesCreateConvertGetQuoteV1**](V1API.md#UfuturesCreateConvertGetQuoteV1) | **Post** /fapi/v1/convert/getQuote | Send Quote Request(USER_DATA)
[**UfuturesCreateCountdownCancelAllV1**](V1API.md#UfuturesCreateCountdownCancelAllV1) | **Post** /fapi/v1/countdownCancelAll | Auto-Cancel All Open Orders (TRADE)
[**UfuturesCreateFeeBurnV1**](V1API.md#UfuturesCreateFeeBurnV1) | **Post** /fapi/v1/feeBurn | Toggle BNB Burn On Futures Trade (TRADE)
[**UfuturesCreateLeverageV1**](V1API.md#UfuturesCreateLeverageV1) | **Post** /fapi/v1/leverage | Change Initial Leverage(TRADE)
[**UfuturesCreateListenKeyV1**](V1API.md#UfuturesCreateListenKeyV1) | **Post** /fapi/v1/listenKey | Start User Data Stream (USER_STREAM)
[**UfuturesCreateMarginTypeV1**](V1API.md#UfuturesCreateMarginTypeV1) | **Post** /fapi/v1/marginType | Change Margin Type(TRADE)
[**UfuturesCreateMultiAssetsMarginV1**](V1API.md#UfuturesCreateMultiAssetsMarginV1) | **Post** /fapi/v1/multiAssetsMargin | Change Multi-Assets Mode (TRADE)
[**UfuturesCreateOrderTestV1**](V1API.md#UfuturesCreateOrderTestV1) | **Post** /fapi/v1/order/test | Test Order(TRADE)
[**UfuturesCreateOrderV1**](V1API.md#UfuturesCreateOrderV1) | **Post** /fapi/v1/order | New Order(TRADE)
[**UfuturesCreatePositionMarginV1**](V1API.md#UfuturesCreatePositionMarginV1) | **Post** /fapi/v1/positionMargin | Modify Isolated Position Margin(TRADE)
[**UfuturesCreatePositionSideDualV1**](V1API.md#UfuturesCreatePositionSideDualV1) | **Post** /fapi/v1/positionSide/dual | Change Position Mode(TRADE)
[**UfuturesDeleteAllOpenOrdersV1**](V1API.md#UfuturesDeleteAllOpenOrdersV1) | **Delete** /fapi/v1/allOpenOrders | Cancel All Open Orders (TRADE)
[**UfuturesDeleteBatchOrdersV1**](V1API.md#UfuturesDeleteBatchOrdersV1) | **Delete** /fapi/v1/batchOrders | Cancel Multiple Orders (TRADE)
[**UfuturesDeleteListenKeyV1**](V1API.md#UfuturesDeleteListenKeyV1) | **Delete** /fapi/v1/listenKey | Close User Data Stream (USER_STREAM)
[**UfuturesDeleteOrderV1**](V1API.md#UfuturesDeleteOrderV1) | **Delete** /fapi/v1/order | Cancel Order (TRADE)
[**UfuturesGetAccountConfigV1**](V1API.md#UfuturesGetAccountConfigV1) | **Get** /fapi/v1/accountConfig | Futures Account Configuration(USER_DATA)
[**UfuturesGetAdlQuantileV1**](V1API.md#UfuturesGetAdlQuantileV1) | **Get** /fapi/v1/adlQuantile | Position ADL Quantile Estimation(USER_DATA)
[**UfuturesGetAggTradesV1**](V1API.md#UfuturesGetAggTradesV1) | **Get** /fapi/v1/aggTrades | Compressed/Aggregate Trades List
[**UfuturesGetAllOrdersV1**](V1API.md#UfuturesGetAllOrdersV1) | **Get** /fapi/v1/allOrders | All Orders (USER_DATA)
[**UfuturesGetApiTradingStatusV1**](V1API.md#UfuturesGetApiTradingStatusV1) | **Get** /fapi/v1/apiTradingStatus | Futures Trading Quantitative Rules Indicators (USER_DATA)
[**UfuturesGetAssetIndexV1**](V1API.md#UfuturesGetAssetIndexV1) | **Get** /fapi/v1/assetIndex | Multi-Assets Mode Asset Index
[**UfuturesGetCommissionRateV1**](V1API.md#UfuturesGetCommissionRateV1) | **Get** /fapi/v1/commissionRate | User Commission Rate (USER_DATA)
[**UfuturesGetConstituentsV1**](V1API.md#UfuturesGetConstituentsV1) | **Get** /fapi/v1/constituents | Query Index Price Constituents
[**UfuturesGetContinuousKlinesV1**](V1API.md#UfuturesGetContinuousKlinesV1) | **Get** /fapi/v1/continuousKlines | Continuous Contract Kline/Candlestick Data
[**UfuturesGetConvertExchangeInfoV1**](V1API.md#UfuturesGetConvertExchangeInfoV1) | **Get** /fapi/v1/convert/exchangeInfo | List All Convert Pairs
[**UfuturesGetConvertOrderStatusV1**](V1API.md#UfuturesGetConvertOrderStatusV1) | **Get** /fapi/v1/convert/orderStatus | Order status(USER_DATA)
[**UfuturesGetDepthV1**](V1API.md#UfuturesGetDepthV1) | **Get** /fapi/v1/depth | Order Book
[**UfuturesGetExchangeInfoV1**](V1API.md#UfuturesGetExchangeInfoV1) | **Get** /fapi/v1/exchangeInfo | Exchange Information
[**UfuturesGetFeeBurnV1**](V1API.md#UfuturesGetFeeBurnV1) | **Get** /fapi/v1/feeBurn | Get BNB Burn Status (USER_DATA)
[**UfuturesGetForceOrdersV1**](V1API.md#UfuturesGetForceOrdersV1) | **Get** /fapi/v1/forceOrders | User&#39;s Force Orders (USER_DATA)
[**UfuturesGetFundingInfoV1**](V1API.md#UfuturesGetFundingInfoV1) | **Get** /fapi/v1/fundingInfo | Get Funding Rate Info
[**UfuturesGetFundingRateV1**](V1API.md#UfuturesGetFundingRateV1) | **Get** /fapi/v1/fundingRate | Get Funding Rate History
[**UfuturesGetHistoricalTradesV1**](V1API.md#UfuturesGetHistoricalTradesV1) | **Get** /fapi/v1/historicalTrades | Old Trades Lookup (MARKET_DATA)
[**UfuturesGetIncomeAsynIdV1**](V1API.md#UfuturesGetIncomeAsynIdV1) | **Get** /fapi/v1/income/asyn/id | Get Futures Transaction History Download Link by Id (USER_DATA)
[**UfuturesGetIncomeAsynV1**](V1API.md#UfuturesGetIncomeAsynV1) | **Get** /fapi/v1/income/asyn | Get Download Id For Futures Transaction History(USER_DATA)
[**UfuturesGetIncomeV1**](V1API.md#UfuturesGetIncomeV1) | **Get** /fapi/v1/income | Get Income History (USER_DATA)
[**UfuturesGetIndexInfoV1**](V1API.md#UfuturesGetIndexInfoV1) | **Get** /fapi/v1/indexInfo | Composite Index Symbol Information
[**UfuturesGetIndexPriceKlinesV1**](V1API.md#UfuturesGetIndexPriceKlinesV1) | **Get** /fapi/v1/indexPriceKlines | Index Price Kline/Candlestick Data
[**UfuturesGetKlinesV1**](V1API.md#UfuturesGetKlinesV1) | **Get** /fapi/v1/klines | Kline/Candlestick Data
[**UfuturesGetLeverageBracketV1**](V1API.md#UfuturesGetLeverageBracketV1) | **Get** /fapi/v1/leverageBracket | Notional and Leverage Brackets (USER_DATA)
[**UfuturesGetMarkPriceKlinesV1**](V1API.md#UfuturesGetMarkPriceKlinesV1) | **Get** /fapi/v1/markPriceKlines | Mark Price Kline/Candlestick Data
[**UfuturesGetMultiAssetsMarginV1**](V1API.md#UfuturesGetMultiAssetsMarginV1) | **Get** /fapi/v1/multiAssetsMargin | Get Current Multi-Assets Mode (USER_DATA)
[**UfuturesGetOpenInterestV1**](V1API.md#UfuturesGetOpenInterestV1) | **Get** /fapi/v1/openInterest | Open Interest
[**UfuturesGetOpenOrderV1**](V1API.md#UfuturesGetOpenOrderV1) | **Get** /fapi/v1/openOrder | Query Current Open Order (USER_DATA)
[**UfuturesGetOpenOrdersV1**](V1API.md#UfuturesGetOpenOrdersV1) | **Get** /fapi/v1/openOrders | Current All Open Orders (USER_DATA)
[**UfuturesGetOrderAmendmentV1**](V1API.md#UfuturesGetOrderAmendmentV1) | **Get** /fapi/v1/orderAmendment | Get Order Modify History (USER_DATA)
[**UfuturesGetOrderAsynIdV1**](V1API.md#UfuturesGetOrderAsynIdV1) | **Get** /fapi/v1/order/asyn/id | Get Futures Order History Download Link by Id (USER_DATA)
[**UfuturesGetOrderAsynV1**](V1API.md#UfuturesGetOrderAsynV1) | **Get** /fapi/v1/order/asyn | Get Download Id For Futures Order History (USER_DATA)
[**UfuturesGetOrderV1**](V1API.md#UfuturesGetOrderV1) | **Get** /fapi/v1/order | Query Order (USER_DATA)
[**UfuturesGetPingV1**](V1API.md#UfuturesGetPingV1) | **Get** /fapi/v1/ping | Test Connectivity
[**UfuturesGetPmAccountInfoV1**](V1API.md#UfuturesGetPmAccountInfoV1) | **Get** /fapi/v1/pmAccountInfo | Classic Portfolio Margin Account Information (USER_DATA)
[**UfuturesGetPositionMarginHistoryV1**](V1API.md#UfuturesGetPositionMarginHistoryV1) | **Get** /fapi/v1/positionMargin/history | Get Position Margin Change History (TRADE)
[**UfuturesGetPositionSideDualV1**](V1API.md#UfuturesGetPositionSideDualV1) | **Get** /fapi/v1/positionSide/dual | Get Current Position Mode(USER_DATA)
[**UfuturesGetPremiumIndexKlinesV1**](V1API.md#UfuturesGetPremiumIndexKlinesV1) | **Get** /fapi/v1/premiumIndexKlines | Premium index Kline Data
[**UfuturesGetPremiumIndexV1**](V1API.md#UfuturesGetPremiumIndexV1) | **Get** /fapi/v1/premiumIndex | Mark Price
[**UfuturesGetRateLimitOrderV1**](V1API.md#UfuturesGetRateLimitOrderV1) | **Get** /fapi/v1/rateLimit/order | Query User Rate Limit (USER_DATA)
[**UfuturesGetSymbolConfigV1**](V1API.md#UfuturesGetSymbolConfigV1) | **Get** /fapi/v1/symbolConfig | Symbol Configuration(USER_DATA)
[**UfuturesGetTicker24hrV1**](V1API.md#UfuturesGetTicker24hrV1) | **Get** /fapi/v1/ticker/24hr | 24hr Ticker Price Change Statistics
[**UfuturesGetTickerBookTickerV1**](V1API.md#UfuturesGetTickerBookTickerV1) | **Get** /fapi/v1/ticker/bookTicker | Symbol Order Book Ticker
[**UfuturesGetTickerPriceV1**](V1API.md#UfuturesGetTickerPriceV1) | **Get** /fapi/v1/ticker/price | Symbol Price Ticker
[**UfuturesGetTimeV1**](V1API.md#UfuturesGetTimeV1) | **Get** /fapi/v1/time | Check Server Time
[**UfuturesGetTradeAsynIdV1**](V1API.md#UfuturesGetTradeAsynIdV1) | **Get** /fapi/v1/trade/asyn/id | Get Futures Trade Download Link by Id(USER_DATA)
[**UfuturesGetTradeAsynV1**](V1API.md#UfuturesGetTradeAsynV1) | **Get** /fapi/v1/trade/asyn | Get Download Id For Futures Trade History (USER_DATA)
[**UfuturesGetTradesV1**](V1API.md#UfuturesGetTradesV1) | **Get** /fapi/v1/trades | Recent Trades List
[**UfuturesGetUserTradesV1**](V1API.md#UfuturesGetUserTradesV1) | **Get** /fapi/v1/userTrades | Account Trade List (USER_DATA)
[**UfuturesUpdateBatchOrdersV1**](V1API.md#UfuturesUpdateBatchOrdersV1) | **Put** /fapi/v1/batchOrders | Modify Multiple Orders(TRADE)
[**UfuturesUpdateListenKeyV1**](V1API.md#UfuturesUpdateListenKeyV1) | **Put** /fapi/v1/listenKey | Keepalive User Data Stream (USER_STREAM)
[**UfuturesUpdateOrderV1**](V1API.md#UfuturesUpdateOrderV1) | **Put** /fapi/v1/order | Modify Order (TRADE)



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
	resp, r, err := apiClient.V1API.UfuturesCreateBatchOrdersV1(context.Background()).BatchOrders(batchOrders).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesCreateBatchOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesCreateBatchOrdersV1`: []UfuturesCreateBatchOrdersV1RespInner
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesCreateBatchOrdersV1`: %v\n", resp)
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


## UfuturesCreateConvertAcceptQuoteV1

> UfuturesCreateConvertAcceptQuoteV1Resp UfuturesCreateConvertAcceptQuoteV1(ctx).QuoteId(quoteId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Accept the offered quote (USER_DATA)



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
	quoteId := "quoteId_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesCreateConvertAcceptQuoteV1(context.Background()).QuoteId(quoteId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesCreateConvertAcceptQuoteV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesCreateConvertAcceptQuoteV1`: UfuturesCreateConvertAcceptQuoteV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesCreateConvertAcceptQuoteV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesCreateConvertAcceptQuoteV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quoteId** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesCreateConvertAcceptQuoteV1Resp**](UfuturesCreateConvertAcceptQuoteV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesCreateConvertGetQuoteV1

> UfuturesCreateConvertGetQuoteV1Resp UfuturesCreateConvertGetQuoteV1(ctx).FromAsset(fromAsset).Timestamp(timestamp).ToAsset(toAsset).FromAmount(fromAmount).RecvWindow(recvWindow).ToAmount(toAmount).ValidTime(validTime).Execute()

Send Quote Request(USER_DATA)



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
	fromAsset := "fromAsset_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	toAsset := "toAsset_example" // string |  (default to "")
	fromAmount := "fromAmount_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	toAmount := "toAmount_example" // string |  (optional) (default to "")
	validTime := "validTime_example" // string |  (optional) (default to "10")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesCreateConvertGetQuoteV1(context.Background()).FromAsset(fromAsset).Timestamp(timestamp).ToAsset(toAsset).FromAmount(fromAmount).RecvWindow(recvWindow).ToAmount(toAmount).ValidTime(validTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesCreateConvertGetQuoteV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesCreateConvertGetQuoteV1`: UfuturesCreateConvertGetQuoteV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesCreateConvertGetQuoteV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesCreateConvertGetQuoteV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromAsset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **toAsset** | **string** |  | [default to &quot;&quot;]
 **fromAmount** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **toAmount** | **string** |  | [default to &quot;&quot;]
 **validTime** | **string** |  | [default to &quot;10&quot;]

### Return type

[**UfuturesCreateConvertGetQuoteV1Resp**](UfuturesCreateConvertGetQuoteV1Resp.md)

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
	resp, r, err := apiClient.V1API.UfuturesCreateCountdownCancelAllV1(context.Background()).UfuturesCreateCountdownCancelAllV1Req(ufuturesCreateCountdownCancelAllV1Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesCreateCountdownCancelAllV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesCreateCountdownCancelAllV1`: UfuturesCreateCountdownCancelAllV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesCreateCountdownCancelAllV1`: %v\n", resp)
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


## UfuturesCreateFeeBurnV1

> UfuturesCreateFeeBurnV1Resp UfuturesCreateFeeBurnV1(ctx).FeeBurn(feeBurn).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Toggle BNB Burn On Futures Trade (TRADE)



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
	feeBurn := "feeBurn_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesCreateFeeBurnV1(context.Background()).FeeBurn(feeBurn).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesCreateFeeBurnV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesCreateFeeBurnV1`: UfuturesCreateFeeBurnV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesCreateFeeBurnV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesCreateFeeBurnV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **feeBurn** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesCreateFeeBurnV1Resp**](UfuturesCreateFeeBurnV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
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
	resp, r, err := apiClient.V1API.UfuturesCreateLeverageV1(context.Background()).Leverage(leverage).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesCreateLeverageV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesCreateLeverageV1`: UfuturesCreateLeverageV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesCreateLeverageV1`: %v\n", resp)
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


## UfuturesCreateListenKeyV1

> UfuturesCreateListenKeyV1Resp UfuturesCreateListenKeyV1(ctx).Execute()

Start User Data Stream (USER_STREAM)



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesCreateListenKeyV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesCreateListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesCreateListenKeyV1`: UfuturesCreateListenKeyV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesCreateListenKeyV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesCreateListenKeyV1Request struct via the builder pattern


### Return type

[**UfuturesCreateListenKeyV1Resp**](UfuturesCreateListenKeyV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
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
	resp, r, err := apiClient.V1API.UfuturesCreateMarginTypeV1(context.Background()).MarginType(marginType).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesCreateMarginTypeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesCreateMarginTypeV1`: UfuturesCreateMarginTypeV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesCreateMarginTypeV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.UfuturesCreateMultiAssetsMarginV1(context.Background()).MultiAssetsMargin(multiAssetsMargin).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesCreateMultiAssetsMarginV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesCreateMultiAssetsMarginV1`: UfuturesCreateMultiAssetsMarginV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesCreateMultiAssetsMarginV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.UfuturesCreateOrderTestV1(context.Background()).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).ActivationPrice(activationPrice).CallbackRate(callbackRate).ClosePosition(closePosition).GoodTillDate(goodTillDate).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).PositionSide(positionSide).Price(price).PriceMatch(priceMatch).PriceProtect(priceProtect).Quantity(quantity).RecvWindow(recvWindow).ReduceOnly(reduceOnly).SelfTradePreventionMode(selfTradePreventionMode).StopPrice(stopPrice).TimeInForce(timeInForce).WorkingType(workingType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesCreateOrderTestV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesCreateOrderTestV1`: UfuturesCreateOrderTestV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesCreateOrderTestV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.UfuturesCreateOrderV1(context.Background()).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).ActivationPrice(activationPrice).CallbackRate(callbackRate).ClosePosition(closePosition).GoodTillDate(goodTillDate).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).PositionSide(positionSide).Price(price).PriceMatch(priceMatch).PriceProtect(priceProtect).Quantity(quantity).RecvWindow(recvWindow).ReduceOnly(reduceOnly).SelfTradePreventionMode(selfTradePreventionMode).StopPrice(stopPrice).TimeInForce(timeInForce).WorkingType(workingType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesCreateOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesCreateOrderV1`: UfuturesCreateOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesCreateOrderV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.UfuturesCreatePositionMarginV1(context.Background()).Amount(amount).Symbol(symbol).Timestamp(timestamp).Type_(type_).PositionSide(positionSide).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesCreatePositionMarginV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesCreatePositionMarginV1`: UfuturesCreatePositionMarginV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesCreatePositionMarginV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.UfuturesCreatePositionSideDualV1(context.Background()).DualSidePosition(dualSidePosition).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesCreatePositionSideDualV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesCreatePositionSideDualV1`: UfuturesCreatePositionSideDualV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesCreatePositionSideDualV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.UfuturesDeleteAllOpenOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesDeleteAllOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesDeleteAllOpenOrdersV1`: UfuturesDeleteAllOpenOrdersV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesDeleteAllOpenOrdersV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.UfuturesDeleteBatchOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderIdList(orderIdList).OrigClientOrderIdList(origClientOrderIdList).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesDeleteBatchOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesDeleteBatchOrdersV1`: []UfuturesDeleteBatchOrdersV1RespInner
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesDeleteBatchOrdersV1`: %v\n", resp)
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


## UfuturesDeleteListenKeyV1

> map[string]interface{} UfuturesDeleteListenKeyV1(ctx).Execute()

Close User Data Stream (USER_STREAM)



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesDeleteListenKeyV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesDeleteListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesDeleteListenKeyV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesDeleteListenKeyV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesDeleteListenKeyV1Request struct via the builder pattern


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
	resp, r, err := apiClient.V1API.UfuturesDeleteOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesDeleteOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesDeleteOrderV1`: UfuturesDeleteOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesDeleteOrderV1`: %v\n", resp)
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


## UfuturesGetAccountConfigV1

> UfuturesGetAccountConfigV1Resp UfuturesGetAccountConfigV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Futures Account Configuration(USER_DATA)



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
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetAccountConfigV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetAccountConfigV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetAccountConfigV1`: UfuturesGetAccountConfigV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetAccountConfigV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetAccountConfigV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetAccountConfigV1Resp**](UfuturesGetAccountConfigV1Resp.md)

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
	resp, r, err := apiClient.V1API.UfuturesGetAdlQuantileV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetAdlQuantileV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetAdlQuantileV1`: []UfuturesGetAdlQuantileV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetAdlQuantileV1`: %v\n", resp)
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


## UfuturesGetAggTradesV1

> []UfuturesGetAggTradesV1RespItem UfuturesGetAggTradesV1(ctx).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Compressed/Aggregate Trades List



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
	fromId := int64(789) // int64 | ID to get aggregate trades from INCLUSIVE. (optional)
	startTime := int64(789) // int64 | Timestamp in ms to get aggregate trades from INCLUSIVE. (optional)
	endTime := int64(789) // int64 | Timestamp in ms to get aggregate trades until INCLUSIVE. (optional)
	limit := int32(56) // int32 | Default 500; max 1000. (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetAggTradesV1(context.Background()).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetAggTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetAggTradesV1`: []UfuturesGetAggTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetAggTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetAggTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **fromId** | **int64** | ID to get aggregate trades from INCLUSIVE. | 
 **startTime** | **int64** | Timestamp in ms to get aggregate trades from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get aggregate trades until INCLUSIVE. | 
 **limit** | **int32** | Default 500; max 1000. | [default to 500]

### Return type

[**[]UfuturesGetAggTradesV1RespItem**](UfuturesGetAggTradesV1RespItem.md)

### Authorization

No authorization required

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
	resp, r, err := apiClient.V1API.UfuturesGetAllOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetAllOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetAllOrdersV1`: []UfuturesGetAllOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetAllOrdersV1`: %v\n", resp)
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


## UfuturesGetApiTradingStatusV1

> UfuturesGetApiTradingStatusV1Resp UfuturesGetApiTradingStatusV1(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Futures Trading Quantitative Rules Indicators (USER_DATA)



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
	resp, r, err := apiClient.V1API.UfuturesGetApiTradingStatusV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetApiTradingStatusV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetApiTradingStatusV1`: UfuturesGetApiTradingStatusV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetApiTradingStatusV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetApiTradingStatusV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetApiTradingStatusV1Resp**](UfuturesGetApiTradingStatusV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetAssetIndexV1

> UfuturesGetAssetIndexV1Resp UfuturesGetAssetIndexV1(ctx).Symbol(symbol).Execute()

Multi-Assets Mode Asset Index



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
	symbol := "symbol_example" // string | Asset pair (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetAssetIndexV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetAssetIndexV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetAssetIndexV1`: UfuturesGetAssetIndexV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetAssetIndexV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetAssetIndexV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Asset pair | [default to &quot;&quot;]

### Return type

[**UfuturesGetAssetIndexV1Resp**](UfuturesGetAssetIndexV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetCommissionRateV1

> UfuturesGetCommissionRateV1Resp UfuturesGetCommissionRateV1(ctx).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

User Commission Rate (USER_DATA)



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
	resp, r, err := apiClient.V1API.UfuturesGetCommissionRateV1(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetCommissionRateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetCommissionRateV1`: UfuturesGetCommissionRateV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetCommissionRateV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetCommissionRateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetCommissionRateV1Resp**](UfuturesGetCommissionRateV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetConstituentsV1

> UfuturesGetConstituentsV1Resp UfuturesGetConstituentsV1(ctx).Symbol(symbol).Execute()

Query Index Price Constituents



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetConstituentsV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetConstituentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetConstituentsV1`: UfuturesGetConstituentsV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetConstituentsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetConstituentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**UfuturesGetConstituentsV1Resp**](UfuturesGetConstituentsV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetContinuousKlinesV1

> [][]UfuturesGetContinuousKlinesV1RespInnerInner UfuturesGetContinuousKlinesV1(ctx).Pair(pair).ContractType(contractType).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Continuous Contract Kline/Candlestick Data



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
	pair := "pair_example" // string |  (default to "")
	contractType := "contractType_example" // string |  (default to "")
	interval := "interval_example" // string |  (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 500; max 1500. (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetContinuousKlinesV1(context.Background()).Pair(pair).ContractType(contractType).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetContinuousKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetContinuousKlinesV1`: [][]UfuturesGetContinuousKlinesV1RespInnerInner
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetContinuousKlinesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetContinuousKlinesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pair** | **string** |  | [default to &quot;&quot;]
 **contractType** | **string** |  | [default to &quot;&quot;]
 **interval** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 1500. | [default to 500]

### Return type

[**[][]UfuturesGetContinuousKlinesV1RespInnerInner**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetConvertExchangeInfoV1

> []UfuturesGetConvertExchangeInfoV1RespItem UfuturesGetConvertExchangeInfoV1(ctx).FromAsset(fromAsset).ToAsset(toAsset).Execute()

List All Convert Pairs



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
	fromAsset := "fromAsset_example" // string | User spends coin (optional) (default to "")
	toAsset := "toAsset_example" // string | User receives coin (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetConvertExchangeInfoV1(context.Background()).FromAsset(fromAsset).ToAsset(toAsset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetConvertExchangeInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetConvertExchangeInfoV1`: []UfuturesGetConvertExchangeInfoV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetConvertExchangeInfoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetConvertExchangeInfoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromAsset** | **string** | User spends coin | [default to &quot;&quot;]
 **toAsset** | **string** | User receives coin | [default to &quot;&quot;]

### Return type

[**[]UfuturesGetConvertExchangeInfoV1RespItem**](UfuturesGetConvertExchangeInfoV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetConvertOrderStatusV1

> UfuturesGetConvertOrderStatusV1Resp UfuturesGetConvertOrderStatusV1(ctx).OrderId(orderId).QuoteId(quoteId).Execute()

Order status(USER_DATA)



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
	orderId := "orderId_example" // string | Either orderId or quoteId is required (optional) (default to "")
	quoteId := "quoteId_example" // string | Either orderId or quoteId is required (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetConvertOrderStatusV1(context.Background()).OrderId(orderId).QuoteId(quoteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetConvertOrderStatusV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetConvertOrderStatusV1`: UfuturesGetConvertOrderStatusV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetConvertOrderStatusV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetConvertOrderStatusV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **string** | Either orderId or quoteId is required | [default to &quot;&quot;]
 **quoteId** | **string** | Either orderId or quoteId is required | [default to &quot;&quot;]

### Return type

[**UfuturesGetConvertOrderStatusV1Resp**](UfuturesGetConvertOrderStatusV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetDepthV1

> UfuturesGetDepthV1Resp UfuturesGetDepthV1(ctx).Symbol(symbol).Limit(limit).Execute()

Order Book



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
	limit := int32(56) // int32 | Default 500; Valid limits:[5, 10, 20, 50, 100, 500, 1000] (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetDepthV1(context.Background()).Symbol(symbol).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetDepthV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetDepthV1`: UfuturesGetDepthV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetDepthV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetDepthV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **limit** | **int32** | Default 500; Valid limits:[5, 10, 20, 50, 100, 500, 1000] | [default to 500]

### Return type

[**UfuturesGetDepthV1Resp**](UfuturesGetDepthV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetExchangeInfoV1

> UfuturesGetExchangeInfoV1Resp UfuturesGetExchangeInfoV1(ctx).Execute()

Exchange Information



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetExchangeInfoV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetExchangeInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetExchangeInfoV1`: UfuturesGetExchangeInfoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetExchangeInfoV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetExchangeInfoV1Request struct via the builder pattern


### Return type

[**UfuturesGetExchangeInfoV1Resp**](UfuturesGetExchangeInfoV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetFeeBurnV1

> UfuturesGetFeeBurnV1Resp UfuturesGetFeeBurnV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get BNB Burn Status (USER_DATA)



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
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetFeeBurnV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetFeeBurnV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetFeeBurnV1`: UfuturesGetFeeBurnV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetFeeBurnV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetFeeBurnV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetFeeBurnV1Resp**](UfuturesGetFeeBurnV1Resp.md)

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
	resp, r, err := apiClient.V1API.UfuturesGetForceOrdersV1(context.Background()).Timestamp(timestamp).Symbol(symbol).AutoCloseType(autoCloseType).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetForceOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetForceOrdersV1`: []UfuturesGetForceOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetForceOrdersV1`: %v\n", resp)
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


## UfuturesGetFundingInfoV1

> []UfuturesGetFundingInfoV1RespItem UfuturesGetFundingInfoV1(ctx).Execute()

Get Funding Rate Info



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetFundingInfoV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetFundingInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetFundingInfoV1`: []UfuturesGetFundingInfoV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetFundingInfoV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetFundingInfoV1Request struct via the builder pattern


### Return type

[**[]UfuturesGetFundingInfoV1RespItem**](UfuturesGetFundingInfoV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetFundingRateV1

> []UfuturesGetFundingRateV1RespItem UfuturesGetFundingRateV1(ctx).Symbol(symbol).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Get Funding Rate History



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
	symbol := "symbol_example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 | Timestamp in ms to get funding rate from INCLUSIVE. (optional)
	endTime := int64(789) // int64 | Timestamp in ms to get funding rate  until INCLUSIVE. (optional)
	limit := int32(56) // int32 | Default 100; max 1000 (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetFundingRateV1(context.Background()).Symbol(symbol).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetFundingRateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetFundingRateV1`: []UfuturesGetFundingRateV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetFundingRateV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetFundingRateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** | Timestamp in ms to get funding rate from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get funding rate  until INCLUSIVE. | 
 **limit** | **int32** | Default 100; max 1000 | [default to 100]

### Return type

[**[]UfuturesGetFundingRateV1RespItem**](UfuturesGetFundingRateV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetHistoricalTradesV1

> []UfuturesGetHistoricalTradesV1RespItem UfuturesGetHistoricalTradesV1(ctx).Symbol(symbol).Limit(limit).FromId(fromId).Execute()

Old Trades Lookup (MARKET_DATA)



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
	limit := int32(56) // int32 | Default 100; max 500. (optional) (default to 100)
	fromId := int64(789) // int64 | TradeId to fetch from. Default gets most recent trades. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetHistoricalTradesV1(context.Background()).Symbol(symbol).Limit(limit).FromId(fromId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetHistoricalTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetHistoricalTradesV1`: []UfuturesGetHistoricalTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetHistoricalTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetHistoricalTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **limit** | **int32** | Default 100; max 500. | [default to 100]
 **fromId** | **int64** | TradeId to fetch from. Default gets most recent trades. | 

### Return type

[**[]UfuturesGetHistoricalTradesV1RespItem**](UfuturesGetHistoricalTradesV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetIncomeAsynIdV1

> UfuturesGetIncomeAsynIdV1Resp UfuturesGetIncomeAsynIdV1(ctx).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Futures Transaction History Download Link by Id (USER_DATA)



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
	downloadId := "downloadId_example" // string | get by download id api (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetIncomeAsynIdV1(context.Background()).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetIncomeAsynIdV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetIncomeAsynIdV1`: UfuturesGetIncomeAsynIdV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetIncomeAsynIdV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetIncomeAsynIdV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadId** | **string** | get by download id api | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetIncomeAsynIdV1Resp**](UfuturesGetIncomeAsynIdV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetIncomeAsynV1

> UfuturesGetIncomeAsynV1Resp UfuturesGetIncomeAsynV1(ctx).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Download Id For Futures Transaction History(USER_DATA)



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
	startTime := int64(789) // int64 | Timestamp in ms
	endTime := int64(789) // int64 | Timestamp in ms
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetIncomeAsynV1(context.Background()).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetIncomeAsynV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetIncomeAsynV1`: UfuturesGetIncomeAsynV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetIncomeAsynV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetIncomeAsynV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** | Timestamp in ms | 
 **endTime** | **int64** | Timestamp in ms | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetIncomeAsynV1Resp**](UfuturesGetIncomeAsynV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetIncomeV1

> []UfuturesGetIncomeV1RespItem UfuturesGetIncomeV1(ctx).Timestamp(timestamp).Symbol(symbol).IncomeType(incomeType).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Get Income History (USER_DATA)



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
	incomeType := "incomeType_example" // string | TRANSFER, WELCOME_BONUS, REALIZED_PNL, FUNDING_FEE, COMMISSION, INSURANCE_CLEAR, REFERRAL_KICKBACK, COMMISSION_REBATE, API_REBATE, CONTEST_REWARD, CROSS_COLLATERAL_TRANSFER, OPTIONS_PREMIUM_FEE, OPTIONS_SETTLE_PROFIT, INTERNAL_TRANSFER, AUTO_EXCHANGE, DELIVERED_SETTELMENT, COIN_SWAP_DEPOSIT, COIN_SWAP_WITHDRAW, POSITION_LIMIT_INCREASE_FEE, STRATEGY_UMFUTURES_TRANSFER，FEE_RETURN，BFUSD_REWARD (optional) (default to "")
	startTime := int64(789) // int64 | Timestamp in ms to get funding from INCLUSIVE. (optional)
	endTime := int64(789) // int64 | Timestamp in ms to get funding until INCLUSIVE. (optional)
	page := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 | Default 100; max 1000 (optional) (default to 100)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetIncomeV1(context.Background()).Timestamp(timestamp).Symbol(symbol).IncomeType(incomeType).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetIncomeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetIncomeV1`: []UfuturesGetIncomeV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetIncomeV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetIncomeV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **incomeType** | **string** | TRANSFER, WELCOME_BONUS, REALIZED_PNL, FUNDING_FEE, COMMISSION, INSURANCE_CLEAR, REFERRAL_KICKBACK, COMMISSION_REBATE, API_REBATE, CONTEST_REWARD, CROSS_COLLATERAL_TRANSFER, OPTIONS_PREMIUM_FEE, OPTIONS_SETTLE_PROFIT, INTERNAL_TRANSFER, AUTO_EXCHANGE, DELIVERED_SETTELMENT, COIN_SWAP_DEPOSIT, COIN_SWAP_WITHDRAW, POSITION_LIMIT_INCREASE_FEE, STRATEGY_UMFUTURES_TRANSFER，FEE_RETURN，BFUSD_REWARD | [default to &quot;&quot;]
 **startTime** | **int64** | Timestamp in ms to get funding from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get funding until INCLUSIVE. | 
 **page** | **int32** |  | 
 **limit** | **int32** | Default 100; max 1000 | [default to 100]
 **recvWindow** | **int64** |  | 

### Return type

[**[]UfuturesGetIncomeV1RespItem**](UfuturesGetIncomeV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetIndexInfoV1

> []UfuturesGetIndexInfoV1RespItem UfuturesGetIndexInfoV1(ctx).Symbol(symbol).Execute()

Composite Index Symbol Information



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
	symbol := "symbol_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetIndexInfoV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetIndexInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetIndexInfoV1`: []UfuturesGetIndexInfoV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetIndexInfoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetIndexInfoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]UfuturesGetIndexInfoV1RespItem**](UfuturesGetIndexInfoV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetIndexPriceKlinesV1

> [][]UfuturesGetContinuousKlinesV1RespInnerInner UfuturesGetIndexPriceKlinesV1(ctx).Pair(pair).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Index Price Kline/Candlestick Data



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
	pair := "pair_example" // string |  (default to "")
	interval := "interval_example" // string |  (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 500; max 1500. (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetIndexPriceKlinesV1(context.Background()).Pair(pair).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetIndexPriceKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetIndexPriceKlinesV1`: [][]UfuturesGetContinuousKlinesV1RespInnerInner
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetIndexPriceKlinesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetIndexPriceKlinesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pair** | **string** |  | [default to &quot;&quot;]
 **interval** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 1500. | [default to 500]

### Return type

[**[][]UfuturesGetContinuousKlinesV1RespInnerInner**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetKlinesV1

> [][]UfuturesGetContinuousKlinesV1RespInnerInner UfuturesGetKlinesV1(ctx).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Kline/Candlestick Data



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
	interval := "interval_example" // string |  (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 500; max 1500. (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetKlinesV1(context.Background()).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetKlinesV1`: [][]UfuturesGetContinuousKlinesV1RespInnerInner
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetKlinesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetKlinesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **interval** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 1500. | [default to 500]

### Return type

[**[][]UfuturesGetContinuousKlinesV1RespInnerInner**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetLeverageBracketV1

> UfuturesGetLeverageBracketV1Resp UfuturesGetLeverageBracketV1(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Notional and Leverage Brackets (USER_DATA)



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
	resp, r, err := apiClient.V1API.UfuturesGetLeverageBracketV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetLeverageBracketV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetLeverageBracketV1`: UfuturesGetLeverageBracketV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetLeverageBracketV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetLeverageBracketV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetLeverageBracketV1Resp**](UfuturesGetLeverageBracketV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetMarkPriceKlinesV1

> [][]UfuturesGetContinuousKlinesV1RespInnerInner UfuturesGetMarkPriceKlinesV1(ctx).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Mark Price Kline/Candlestick Data



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
	interval := "interval_example" // string |  (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 500; max 1500. (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetMarkPriceKlinesV1(context.Background()).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetMarkPriceKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetMarkPriceKlinesV1`: [][]UfuturesGetContinuousKlinesV1RespInnerInner
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetMarkPriceKlinesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetMarkPriceKlinesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **interval** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 1500. | [default to 500]

### Return type

[**[][]UfuturesGetContinuousKlinesV1RespInnerInner**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetMultiAssetsMarginV1

> UfuturesGetMultiAssetsMarginV1Resp UfuturesGetMultiAssetsMarginV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Current Multi-Assets Mode (USER_DATA)



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
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetMultiAssetsMarginV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetMultiAssetsMarginV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetMultiAssetsMarginV1`: UfuturesGetMultiAssetsMarginV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetMultiAssetsMarginV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetMultiAssetsMarginV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetMultiAssetsMarginV1Resp**](UfuturesGetMultiAssetsMarginV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetOpenInterestV1

> UfuturesGetOpenInterestV1Resp UfuturesGetOpenInterestV1(ctx).Symbol(symbol).Execute()

Open Interest



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetOpenInterestV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetOpenInterestV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetOpenInterestV1`: UfuturesGetOpenInterestV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetOpenInterestV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetOpenInterestV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**UfuturesGetOpenInterestV1Resp**](UfuturesGetOpenInterestV1Resp.md)

### Authorization

No authorization required

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
	resp, r, err := apiClient.V1API.UfuturesGetOpenOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetOpenOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetOpenOrderV1`: UfuturesGetOpenOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetOpenOrderV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.UfuturesGetOpenOrdersV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetOpenOrdersV1`: []UfuturesGetOpenOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetOpenOrdersV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.UfuturesGetOrderAmendmentV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetOrderAmendmentV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetOrderAmendmentV1`: []UfuturesGetOrderAmendmentV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetOrderAmendmentV1`: %v\n", resp)
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


## UfuturesGetOrderAsynIdV1

> UfuturesGetOrderAsynIdV1Resp UfuturesGetOrderAsynIdV1(ctx).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Futures Order History Download Link by Id (USER_DATA)



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
	downloadId := "downloadId_example" // string | get by download id api (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetOrderAsynIdV1(context.Background()).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetOrderAsynIdV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetOrderAsynIdV1`: UfuturesGetOrderAsynIdV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetOrderAsynIdV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetOrderAsynIdV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadId** | **string** | get by download id api | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetOrderAsynIdV1Resp**](UfuturesGetOrderAsynIdV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetOrderAsynV1

> UfuturesGetOrderAsynV1Resp UfuturesGetOrderAsynV1(ctx).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Download Id For Futures Order History (USER_DATA)



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
	startTime := int64(789) // int64 | Timestamp in ms
	endTime := int64(789) // int64 | Timestamp in ms
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetOrderAsynV1(context.Background()).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetOrderAsynV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetOrderAsynV1`: UfuturesGetOrderAsynV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetOrderAsynV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetOrderAsynV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** | Timestamp in ms | 
 **endTime** | **int64** | Timestamp in ms | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetOrderAsynV1Resp**](UfuturesGetOrderAsynV1Resp.md)

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
	resp, r, err := apiClient.V1API.UfuturesGetOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetOrderV1`: UfuturesGetOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetOrderV1`: %v\n", resp)
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


## UfuturesGetPingV1

> map[string]interface{} UfuturesGetPingV1(ctx).Execute()

Test Connectivity



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetPingV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetPingV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetPingV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetPingV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetPingV1Request struct via the builder pattern


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


## UfuturesGetPmAccountInfoV1

> UfuturesGetPmAccountInfoV1Resp UfuturesGetPmAccountInfoV1(ctx).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Classic Portfolio Margin Account Information (USER_DATA)



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
	asset := "asset_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetPmAccountInfoV1(context.Background()).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetPmAccountInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetPmAccountInfoV1`: UfuturesGetPmAccountInfoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetPmAccountInfoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetPmAccountInfoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetPmAccountInfoV1Resp**](UfuturesGetPmAccountInfoV1Resp.md)

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
	resp, r, err := apiClient.V1API.UfuturesGetPositionMarginHistoryV1(context.Background()).Symbol(symbol).Timestamp(timestamp).Type_(type_).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetPositionMarginHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetPositionMarginHistoryV1`: []UfuturesGetPositionMarginHistoryV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetPositionMarginHistoryV1`: %v\n", resp)
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


## UfuturesGetPositionSideDualV1

> UfuturesGetPositionSideDualV1Resp UfuturesGetPositionSideDualV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Current Position Mode(USER_DATA)



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
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetPositionSideDualV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetPositionSideDualV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetPositionSideDualV1`: UfuturesGetPositionSideDualV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetPositionSideDualV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetPositionSideDualV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetPositionSideDualV1Resp**](UfuturesGetPositionSideDualV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetPremiumIndexKlinesV1

> [][]UfuturesGetContinuousKlinesV1RespInnerInner UfuturesGetPremiumIndexKlinesV1(ctx).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Premium index Kline Data



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
	interval := "interval_example" // string |  (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 500; max 1500. (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetPremiumIndexKlinesV1(context.Background()).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetPremiumIndexKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetPremiumIndexKlinesV1`: [][]UfuturesGetContinuousKlinesV1RespInnerInner
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetPremiumIndexKlinesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetPremiumIndexKlinesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **interval** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 1500. | [default to 500]

### Return type

[**[][]UfuturesGetContinuousKlinesV1RespInnerInner**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetPremiumIndexV1

> UfuturesGetPremiumIndexV1Resp UfuturesGetPremiumIndexV1(ctx).Symbol(symbol).Execute()

Mark Price



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
	symbol := "symbol_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetPremiumIndexV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetPremiumIndexV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetPremiumIndexV1`: UfuturesGetPremiumIndexV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetPremiumIndexV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetPremiumIndexV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**UfuturesGetPremiumIndexV1Resp**](UfuturesGetPremiumIndexV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetRateLimitOrderV1

> []UfuturesGetRateLimitOrderV1RespItem UfuturesGetRateLimitOrderV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query User Rate Limit (USER_DATA)



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
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetRateLimitOrderV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetRateLimitOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetRateLimitOrderV1`: []UfuturesGetRateLimitOrderV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetRateLimitOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetRateLimitOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]UfuturesGetRateLimitOrderV1RespItem**](UfuturesGetRateLimitOrderV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetSymbolConfigV1

> []UfuturesGetSymbolConfigV1RespItem UfuturesGetSymbolConfigV1(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Symbol Configuration(USER_DATA)



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
	resp, r, err := apiClient.V1API.UfuturesGetSymbolConfigV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetSymbolConfigV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetSymbolConfigV1`: []UfuturesGetSymbolConfigV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetSymbolConfigV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetSymbolConfigV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]UfuturesGetSymbolConfigV1RespItem**](UfuturesGetSymbolConfigV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetTicker24hrV1

> UfuturesGetTicker24hrV1Resp UfuturesGetTicker24hrV1(ctx).Symbol(symbol).Execute()

24hr Ticker Price Change Statistics



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
	symbol := "symbol_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetTicker24hrV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetTicker24hrV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetTicker24hrV1`: UfuturesGetTicker24hrV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetTicker24hrV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetTicker24hrV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**UfuturesGetTicker24hrV1Resp**](UfuturesGetTicker24hrV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetTickerBookTickerV1

> UfuturesGetTickerBookTickerV1Resp UfuturesGetTickerBookTickerV1(ctx).Symbol(symbol).Execute()

Symbol Order Book Ticker



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
	symbol := "symbol_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetTickerBookTickerV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetTickerBookTickerV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetTickerBookTickerV1`: UfuturesGetTickerBookTickerV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetTickerBookTickerV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetTickerBookTickerV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**UfuturesGetTickerBookTickerV1Resp**](UfuturesGetTickerBookTickerV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetTickerPriceV1

> UfuturesGetTickerPriceV1Resp UfuturesGetTickerPriceV1(ctx).Symbol(symbol).Execute()

Symbol Price Ticker



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
	symbol := "symbol_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetTickerPriceV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetTickerPriceV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetTickerPriceV1`: UfuturesGetTickerPriceV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetTickerPriceV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetTickerPriceV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**UfuturesGetTickerPriceV1Resp**](UfuturesGetTickerPriceV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetTimeV1

> UfuturesGetTimeV1Resp UfuturesGetTimeV1(ctx).Execute()

Check Server Time



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetTimeV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetTimeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetTimeV1`: UfuturesGetTimeV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetTimeV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetTimeV1Request struct via the builder pattern


### Return type

[**UfuturesGetTimeV1Resp**](UfuturesGetTimeV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetTradeAsynIdV1

> UfuturesGetTradeAsynIdV1Resp UfuturesGetTradeAsynIdV1(ctx).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Futures Trade Download Link by Id(USER_DATA)



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
	downloadId := "downloadId_example" // string | get by download id api (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetTradeAsynIdV1(context.Background()).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetTradeAsynIdV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetTradeAsynIdV1`: UfuturesGetTradeAsynIdV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetTradeAsynIdV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetTradeAsynIdV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadId** | **string** | get by download id api | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetTradeAsynIdV1Resp**](UfuturesGetTradeAsynIdV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetTradeAsynV1

> UfuturesGetTradeAsynV1Resp UfuturesGetTradeAsynV1(ctx).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Download Id For Futures Trade History (USER_DATA)



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
	startTime := int64(789) // int64 | Timestamp in ms
	endTime := int64(789) // int64 | Timestamp in ms
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetTradeAsynV1(context.Background()).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetTradeAsynV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetTradeAsynV1`: UfuturesGetTradeAsynV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetTradeAsynV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetTradeAsynV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** | Timestamp in ms | 
 **endTime** | **int64** | Timestamp in ms | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetTradeAsynV1Resp**](UfuturesGetTradeAsynV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetTradesV1

> []UfuturesGetTradesV1RespItem UfuturesGetTradesV1(ctx).Symbol(symbol).Limit(limit).Execute()

Recent Trades List



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
	limit := int32(56) // int32 | Default 500; max 1000. (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesGetTradesV1(context.Background()).Symbol(symbol).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetTradesV1`: []UfuturesGetTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **limit** | **int32** | Default 500; max 1000. | [default to 500]

### Return type

[**[]UfuturesGetTradesV1RespItem**](UfuturesGetTradesV1RespItem.md)

### Authorization

No authorization required

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
	resp, r, err := apiClient.V1API.UfuturesGetUserTradesV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesGetUserTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetUserTradesV1`: []UfuturesGetUserTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesGetUserTradesV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.UfuturesUpdateBatchOrdersV1(context.Background()).BatchOrders(batchOrders).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesUpdateBatchOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesUpdateBatchOrdersV1`: []UfuturesUpdateBatchOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesUpdateBatchOrdersV1`: %v\n", resp)
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


## UfuturesUpdateListenKeyV1

> UfuturesUpdateListenKeyV1Resp UfuturesUpdateListenKeyV1(ctx).Execute()

Keepalive User Data Stream (USER_STREAM)



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UfuturesUpdateListenKeyV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesUpdateListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesUpdateListenKeyV1`: UfuturesUpdateListenKeyV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesUpdateListenKeyV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesUpdateListenKeyV1Request struct via the builder pattern


### Return type

[**UfuturesUpdateListenKeyV1Resp**](UfuturesUpdateListenKeyV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
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
	resp, r, err := apiClient.V1API.UfuturesUpdateOrderV1(context.Background()).Price(price).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).PriceMatch(priceMatch).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UfuturesUpdateOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesUpdateOrderV1`: UfuturesUpdateOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.UfuturesUpdateOrderV1`: %v\n", resp)
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

