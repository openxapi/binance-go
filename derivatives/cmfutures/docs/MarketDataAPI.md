# \MarketDataAPI

All URIs are relative to *https://dapi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CfuturesGetAggTradesV1**](MarketDataAPI.md#CfuturesGetAggTradesV1) | **Get** /dapi/v1/aggTrades | Compressed/Aggregate Trades List
[**CfuturesGetConstituentsV1**](MarketDataAPI.md#CfuturesGetConstituentsV1) | **Get** /dapi/v1/constituents | Query Index Price Constituents
[**CfuturesGetContinuousKlinesV1**](MarketDataAPI.md#CfuturesGetContinuousKlinesV1) | **Get** /dapi/v1/continuousKlines | Continuous Contract Kline/Candlestick Data
[**CfuturesGetDepthV1**](MarketDataAPI.md#CfuturesGetDepthV1) | **Get** /dapi/v1/depth | Order Book
[**CfuturesGetExchangeInfoV1**](MarketDataAPI.md#CfuturesGetExchangeInfoV1) | **Get** /dapi/v1/exchangeInfo | Exchange Information
[**CfuturesGetFundingInfoV1**](MarketDataAPI.md#CfuturesGetFundingInfoV1) | **Get** /dapi/v1/fundingInfo | Get Funding Rate Info
[**CfuturesGetFundingRateV1**](MarketDataAPI.md#CfuturesGetFundingRateV1) | **Get** /dapi/v1/fundingRate | Get Funding Rate History of Perpetual Futures
[**CfuturesGetFuturesDataBasis**](MarketDataAPI.md#CfuturesGetFuturesDataBasis) | **Get** /futures/data/basis | Basis
[**CfuturesGetFuturesDataGlobalLongShortAccountRatio**](MarketDataAPI.md#CfuturesGetFuturesDataGlobalLongShortAccountRatio) | **Get** /futures/data/globalLongShortAccountRatio | Long/Short Ratio
[**CfuturesGetFuturesDataOpenInterestHist**](MarketDataAPI.md#CfuturesGetFuturesDataOpenInterestHist) | **Get** /futures/data/openInterestHist | Open Interest Statistics
[**CfuturesGetFuturesDataTakerBuySellVol**](MarketDataAPI.md#CfuturesGetFuturesDataTakerBuySellVol) | **Get** /futures/data/takerBuySellVol | Taker Buy/Sell Volume
[**CfuturesGetFuturesDataTopLongShortAccountRatio**](MarketDataAPI.md#CfuturesGetFuturesDataTopLongShortAccountRatio) | **Get** /futures/data/topLongShortAccountRatio | Top Trader Long/Short Ratio (Accounts)
[**CfuturesGetFuturesDataTopLongShortPositionRatio**](MarketDataAPI.md#CfuturesGetFuturesDataTopLongShortPositionRatio) | **Get** /futures/data/topLongShortPositionRatio | Top Trader Long/Short Ratio (Positions)
[**CfuturesGetHistoricalTradesV1**](MarketDataAPI.md#CfuturesGetHistoricalTradesV1) | **Get** /dapi/v1/historicalTrades | Old Trades Lookup(MARKET_DATA)
[**CfuturesGetIndexPriceKlinesV1**](MarketDataAPI.md#CfuturesGetIndexPriceKlinesV1) | **Get** /dapi/v1/indexPriceKlines | Index Price Kline/Candlestick Data
[**CfuturesGetKlinesV1**](MarketDataAPI.md#CfuturesGetKlinesV1) | **Get** /dapi/v1/klines | Kline/Candlestick Data
[**CfuturesGetMarkPriceKlinesV1**](MarketDataAPI.md#CfuturesGetMarkPriceKlinesV1) | **Get** /dapi/v1/markPriceKlines | Mark Price Kline/Candlestick Data
[**CfuturesGetOpenInterestV1**](MarketDataAPI.md#CfuturesGetOpenInterestV1) | **Get** /dapi/v1/openInterest | Open Interest
[**CfuturesGetPingV1**](MarketDataAPI.md#CfuturesGetPingV1) | **Get** /dapi/v1/ping | Test Connectivity
[**CfuturesGetPremiumIndexKlinesV1**](MarketDataAPI.md#CfuturesGetPremiumIndexKlinesV1) | **Get** /dapi/v1/premiumIndexKlines | Premium index Kline Data
[**CfuturesGetPremiumIndexV1**](MarketDataAPI.md#CfuturesGetPremiumIndexV1) | **Get** /dapi/v1/premiumIndex | Index Price and Mark Price
[**CfuturesGetTicker24hrV1**](MarketDataAPI.md#CfuturesGetTicker24hrV1) | **Get** /dapi/v1/ticker/24hr | 24hr Ticker Price Change Statistics
[**CfuturesGetTickerBookTickerV1**](MarketDataAPI.md#CfuturesGetTickerBookTickerV1) | **Get** /dapi/v1/ticker/bookTicker | Symbol Order Book Ticker
[**CfuturesGetTickerPriceV1**](MarketDataAPI.md#CfuturesGetTickerPriceV1) | **Get** /dapi/v1/ticker/price | Symbol Price Ticker
[**CfuturesGetTimeV1**](MarketDataAPI.md#CfuturesGetTimeV1) | **Get** /dapi/v1/time | Check Server time
[**CfuturesGetTradesV1**](MarketDataAPI.md#CfuturesGetTradesV1) | **Get** /dapi/v1/trades | Recent Trades List



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
	resp, r, err := apiClient.MarketDataAPI.CfuturesGetAggTradesV1(context.Background()).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CfuturesGetAggTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetAggTradesV1`: []CfuturesGetAggTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CfuturesGetAggTradesV1`: %v\n", resp)
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
	resp, r, err := apiClient.MarketDataAPI.CfuturesGetConstituentsV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CfuturesGetConstituentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetConstituentsV1`: CfuturesGetConstituentsV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CfuturesGetConstituentsV1`: %v\n", resp)
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
	resp, r, err := apiClient.MarketDataAPI.CfuturesGetContinuousKlinesV1(context.Background()).Pair(pair).ContractType(contractType).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CfuturesGetContinuousKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetContinuousKlinesV1`: [][]CfuturesGetContinuousKlinesV1RespInnerInner
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CfuturesGetContinuousKlinesV1`: %v\n", resp)
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
	resp, r, err := apiClient.MarketDataAPI.CfuturesGetDepthV1(context.Background()).Symbol(symbol).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CfuturesGetDepthV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetDepthV1`: CfuturesGetDepthV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CfuturesGetDepthV1`: %v\n", resp)
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
	resp, r, err := apiClient.MarketDataAPI.CfuturesGetExchangeInfoV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CfuturesGetExchangeInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetExchangeInfoV1`: CfuturesGetExchangeInfoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CfuturesGetExchangeInfoV1`: %v\n", resp)
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
	resp, r, err := apiClient.MarketDataAPI.CfuturesGetFundingInfoV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CfuturesGetFundingInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetFundingInfoV1`: []CfuturesGetFundingInfoV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CfuturesGetFundingInfoV1`: %v\n", resp)
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
	resp, r, err := apiClient.MarketDataAPI.CfuturesGetFundingRateV1(context.Background()).Symbol(symbol).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CfuturesGetFundingRateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetFundingRateV1`: []CfuturesGetFundingRateV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CfuturesGetFundingRateV1`: %v\n", resp)
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


## CfuturesGetFuturesDataBasis

> []CfuturesGetFuturesDataBasisRespItem CfuturesGetFuturesDataBasis(ctx).Pair(pair).ContractType(contractType).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

Basis



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
	pair := "pair_example" // string | BTCUSD (default to "")
	contractType := "contractType_example" // string | CURRENT_QUARTER, NEXT_QUARTER, PERPETUAL (default to "")
	period := "period_example" // string | &#34;5m&#34;,&#34;15m&#34;,&#34;30m&#34;,&#34;1h&#34;,&#34;2h&#34;,&#34;4h&#34;,&#34;6h&#34;,&#34;12h&#34;,&#34;1d&#34; (default to "")
	limit := int64(789) // int64 | Default 30,Max 500 (optional) (default to 30)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.CfuturesGetFuturesDataBasis(context.Background()).Pair(pair).ContractType(contractType).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CfuturesGetFuturesDataBasis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetFuturesDataBasis`: []CfuturesGetFuturesDataBasisRespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CfuturesGetFuturesDataBasis`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetFuturesDataBasisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pair** | **string** | BTCUSD | [default to &quot;&quot;]
 **contractType** | **string** | CURRENT_QUARTER, NEXT_QUARTER, PERPETUAL | [default to &quot;&quot;]
 **period** | **string** | &amp;#34;5m&amp;#34;,&amp;#34;15m&amp;#34;,&amp;#34;30m&amp;#34;,&amp;#34;1h&amp;#34;,&amp;#34;2h&amp;#34;,&amp;#34;4h&amp;#34;,&amp;#34;6h&amp;#34;,&amp;#34;12h&amp;#34;,&amp;#34;1d&amp;#34; | [default to &quot;&quot;]
 **limit** | **int64** | Default 30,Max 500 | [default to 30]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**[]CfuturesGetFuturesDataBasisRespItem**](CfuturesGetFuturesDataBasisRespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetFuturesDataGlobalLongShortAccountRatio

> []CfuturesGetFuturesDataGlobalLongShortAccountRatioRespItem CfuturesGetFuturesDataGlobalLongShortAccountRatio(ctx).Pair(pair).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

Long/Short Ratio



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
	pair := "pair_example" // string | BTCUSD (default to "")
	period := "period_example" // string | &#34;5m&#34;,&#34;15m&#34;,&#34;30m&#34;,&#34;1h&#34;,&#34;2h&#34;,&#34;4h&#34;,&#34;6h&#34;,&#34;12h&#34;,&#34;1d&#34; (default to "")
	limit := int64(789) // int64 | Default 30,Max 500 (optional) (default to 30)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.CfuturesGetFuturesDataGlobalLongShortAccountRatio(context.Background()).Pair(pair).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CfuturesGetFuturesDataGlobalLongShortAccountRatio``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetFuturesDataGlobalLongShortAccountRatio`: []CfuturesGetFuturesDataGlobalLongShortAccountRatioRespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CfuturesGetFuturesDataGlobalLongShortAccountRatio`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetFuturesDataGlobalLongShortAccountRatioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pair** | **string** | BTCUSD | [default to &quot;&quot;]
 **period** | **string** | &amp;#34;5m&amp;#34;,&amp;#34;15m&amp;#34;,&amp;#34;30m&amp;#34;,&amp;#34;1h&amp;#34;,&amp;#34;2h&amp;#34;,&amp;#34;4h&amp;#34;,&amp;#34;6h&amp;#34;,&amp;#34;12h&amp;#34;,&amp;#34;1d&amp;#34; | [default to &quot;&quot;]
 **limit** | **int64** | Default 30,Max 500 | [default to 30]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**[]CfuturesGetFuturesDataGlobalLongShortAccountRatioRespItem**](CfuturesGetFuturesDataGlobalLongShortAccountRatioRespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetFuturesDataOpenInterestHist

> []CfuturesGetFuturesDataOpenInterestHistRespItem CfuturesGetFuturesDataOpenInterestHist(ctx).Pair(pair).ContractType(contractType).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

Open Interest Statistics



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
	pair := "pair_example" // string | BTCUSD (default to "")
	contractType := "contractType_example" // string | ALL, CURRENT_QUARTER, NEXT_QUARTER, PERPETUAL (default to "")
	period := "period_example" // string | &#34;5m&#34;,&#34;15m&#34;,&#34;30m&#34;,&#34;1h&#34;,&#34;2h&#34;,&#34;4h&#34;,&#34;6h&#34;,&#34;12h&#34;,&#34;1d&#34; (default to "")
	limit := int64(789) // int64 | Default 30,Max 500 (optional) (default to 30)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.CfuturesGetFuturesDataOpenInterestHist(context.Background()).Pair(pair).ContractType(contractType).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CfuturesGetFuturesDataOpenInterestHist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetFuturesDataOpenInterestHist`: []CfuturesGetFuturesDataOpenInterestHistRespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CfuturesGetFuturesDataOpenInterestHist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetFuturesDataOpenInterestHistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pair** | **string** | BTCUSD | [default to &quot;&quot;]
 **contractType** | **string** | ALL, CURRENT_QUARTER, NEXT_QUARTER, PERPETUAL | [default to &quot;&quot;]
 **period** | **string** | &amp;#34;5m&amp;#34;,&amp;#34;15m&amp;#34;,&amp;#34;30m&amp;#34;,&amp;#34;1h&amp;#34;,&amp;#34;2h&amp;#34;,&amp;#34;4h&amp;#34;,&amp;#34;6h&amp;#34;,&amp;#34;12h&amp;#34;,&amp;#34;1d&amp;#34; | [default to &quot;&quot;]
 **limit** | **int64** | Default 30,Max 500 | [default to 30]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**[]CfuturesGetFuturesDataOpenInterestHistRespItem**](CfuturesGetFuturesDataOpenInterestHistRespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetFuturesDataTakerBuySellVol

> []CfuturesGetFuturesDataTakerBuySellVolRespItem CfuturesGetFuturesDataTakerBuySellVol(ctx).Pair(pair).ContractType(contractType).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

Taker Buy/Sell Volume



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
	pair := "pair_example" // string | BTCUSD (default to "")
	contractType := "contractType_example" // string | ALL, CURRENT_QUARTER, NEXT_QUARTER, PERPETUAL (default to "")
	period := "period_example" // string | &#34;5m&#34;,&#34;15m&#34;,&#34;30m&#34;,&#34;1h&#34;,&#34;2h&#34;,&#34;4h&#34;,&#34;6h&#34;,&#34;12h&#34;,&#34;1d&#34; (default to "")
	limit := int64(789) // int64 | Default 30,Max 500 (optional) (default to 30)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.CfuturesGetFuturesDataTakerBuySellVol(context.Background()).Pair(pair).ContractType(contractType).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CfuturesGetFuturesDataTakerBuySellVol``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetFuturesDataTakerBuySellVol`: []CfuturesGetFuturesDataTakerBuySellVolRespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CfuturesGetFuturesDataTakerBuySellVol`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetFuturesDataTakerBuySellVolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pair** | **string** | BTCUSD | [default to &quot;&quot;]
 **contractType** | **string** | ALL, CURRENT_QUARTER, NEXT_QUARTER, PERPETUAL | [default to &quot;&quot;]
 **period** | **string** | &amp;#34;5m&amp;#34;,&amp;#34;15m&amp;#34;,&amp;#34;30m&amp;#34;,&amp;#34;1h&amp;#34;,&amp;#34;2h&amp;#34;,&amp;#34;4h&amp;#34;,&amp;#34;6h&amp;#34;,&amp;#34;12h&amp;#34;,&amp;#34;1d&amp;#34; | [default to &quot;&quot;]
 **limit** | **int64** | Default 30,Max 500 | [default to 30]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**[]CfuturesGetFuturesDataTakerBuySellVolRespItem**](CfuturesGetFuturesDataTakerBuySellVolRespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetFuturesDataTopLongShortAccountRatio

> []CfuturesGetFuturesDataTopLongShortAccountRatioRespItem CfuturesGetFuturesDataTopLongShortAccountRatio(ctx).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

Top Trader Long/Short Ratio (Accounts)



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
	period := "period_example" // string | &#34;5m&#34;,&#34;15m&#34;,&#34;30m&#34;,&#34;1h&#34;,&#34;2h&#34;,&#34;4h&#34;,&#34;6h&#34;,&#34;12h&#34;,&#34;1d&#34; (default to "")
	limit := int64(789) // int64 | default 30, max 500 (optional) (default to 30)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.CfuturesGetFuturesDataTopLongShortAccountRatio(context.Background()).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CfuturesGetFuturesDataTopLongShortAccountRatio``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetFuturesDataTopLongShortAccountRatio`: []CfuturesGetFuturesDataTopLongShortAccountRatioRespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CfuturesGetFuturesDataTopLongShortAccountRatio`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetFuturesDataTopLongShortAccountRatioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **period** | **string** | &amp;#34;5m&amp;#34;,&amp;#34;15m&amp;#34;,&amp;#34;30m&amp;#34;,&amp;#34;1h&amp;#34;,&amp;#34;2h&amp;#34;,&amp;#34;4h&amp;#34;,&amp;#34;6h&amp;#34;,&amp;#34;12h&amp;#34;,&amp;#34;1d&amp;#34; | [default to &quot;&quot;]
 **limit** | **int64** | default 30, max 500 | [default to 30]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**[]CfuturesGetFuturesDataTopLongShortAccountRatioRespItem**](CfuturesGetFuturesDataTopLongShortAccountRatioRespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetFuturesDataTopLongShortPositionRatio

> []CfuturesGetFuturesDataTopLongShortPositionRatioRespItem CfuturesGetFuturesDataTopLongShortPositionRatio(ctx).Pair(pair).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

Top Trader Long/Short Ratio (Positions)



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
	pair := "pair_example" // string | BTCUSD (default to "")
	period := "period_example" // string | &#34;5m&#34;,&#34;15m&#34;,&#34;30m&#34;,&#34;1h&#34;,&#34;2h&#34;,&#34;4h&#34;,&#34;6h&#34;,&#34;12h&#34;,&#34;1d&#34; (default to "")
	limit := int64(789) // int64 | Default 30,Max 500 (optional) (default to 30)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.CfuturesGetFuturesDataTopLongShortPositionRatio(context.Background()).Pair(pair).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CfuturesGetFuturesDataTopLongShortPositionRatio``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetFuturesDataTopLongShortPositionRatio`: []CfuturesGetFuturesDataTopLongShortPositionRatioRespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CfuturesGetFuturesDataTopLongShortPositionRatio`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetFuturesDataTopLongShortPositionRatioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pair** | **string** | BTCUSD | [default to &quot;&quot;]
 **period** | **string** | &amp;#34;5m&amp;#34;,&amp;#34;15m&amp;#34;,&amp;#34;30m&amp;#34;,&amp;#34;1h&amp;#34;,&amp;#34;2h&amp;#34;,&amp;#34;4h&amp;#34;,&amp;#34;6h&amp;#34;,&amp;#34;12h&amp;#34;,&amp;#34;1d&amp;#34; | [default to &quot;&quot;]
 **limit** | **int64** | Default 30,Max 500 | [default to 30]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**[]CfuturesGetFuturesDataTopLongShortPositionRatioRespItem**](CfuturesGetFuturesDataTopLongShortPositionRatioRespItem.md)

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
	resp, r, err := apiClient.MarketDataAPI.CfuturesGetHistoricalTradesV1(context.Background()).Symbol(symbol).Limit(limit).FromId(fromId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CfuturesGetHistoricalTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetHistoricalTradesV1`: []CfuturesGetHistoricalTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CfuturesGetHistoricalTradesV1`: %v\n", resp)
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
	resp, r, err := apiClient.MarketDataAPI.CfuturesGetIndexPriceKlinesV1(context.Background()).Pair(pair).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CfuturesGetIndexPriceKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetIndexPriceKlinesV1`: [][]CfuturesGetContinuousKlinesV1RespInnerInner
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CfuturesGetIndexPriceKlinesV1`: %v\n", resp)
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
	resp, r, err := apiClient.MarketDataAPI.CfuturesGetKlinesV1(context.Background()).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CfuturesGetKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetKlinesV1`: [][]CfuturesGetContinuousKlinesV1RespInnerInner
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CfuturesGetKlinesV1`: %v\n", resp)
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
	resp, r, err := apiClient.MarketDataAPI.CfuturesGetMarkPriceKlinesV1(context.Background()).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CfuturesGetMarkPriceKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetMarkPriceKlinesV1`: [][]CfuturesGetContinuousKlinesV1RespInnerInner
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CfuturesGetMarkPriceKlinesV1`: %v\n", resp)
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
	resp, r, err := apiClient.MarketDataAPI.CfuturesGetOpenInterestV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CfuturesGetOpenInterestV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetOpenInterestV1`: CfuturesGetOpenInterestV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CfuturesGetOpenInterestV1`: %v\n", resp)
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
	resp, r, err := apiClient.MarketDataAPI.CfuturesGetPingV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CfuturesGetPingV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetPingV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CfuturesGetPingV1`: %v\n", resp)
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
	resp, r, err := apiClient.MarketDataAPI.CfuturesGetPremiumIndexKlinesV1(context.Background()).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CfuturesGetPremiumIndexKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetPremiumIndexKlinesV1`: [][]CfuturesGetContinuousKlinesV1RespInnerInner
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CfuturesGetPremiumIndexKlinesV1`: %v\n", resp)
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
	resp, r, err := apiClient.MarketDataAPI.CfuturesGetPremiumIndexV1(context.Background()).Symbol(symbol).Pair(pair).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CfuturesGetPremiumIndexV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetPremiumIndexV1`: []CfuturesGetPremiumIndexV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CfuturesGetPremiumIndexV1`: %v\n", resp)
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
	resp, r, err := apiClient.MarketDataAPI.CfuturesGetTicker24hrV1(context.Background()).Symbol(symbol).Pair(pair).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CfuturesGetTicker24hrV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetTicker24hrV1`: []CfuturesGetTicker24hrV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CfuturesGetTicker24hrV1`: %v\n", resp)
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
	resp, r, err := apiClient.MarketDataAPI.CfuturesGetTickerBookTickerV1(context.Background()).Symbol(symbol).Pair(pair).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CfuturesGetTickerBookTickerV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetTickerBookTickerV1`: []CfuturesGetTickerBookTickerV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CfuturesGetTickerBookTickerV1`: %v\n", resp)
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
	resp, r, err := apiClient.MarketDataAPI.CfuturesGetTickerPriceV1(context.Background()).Symbol(symbol).Pair(pair).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CfuturesGetTickerPriceV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetTickerPriceV1`: []CfuturesGetTickerPriceV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CfuturesGetTickerPriceV1`: %v\n", resp)
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
	resp, r, err := apiClient.MarketDataAPI.CfuturesGetTimeV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CfuturesGetTimeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetTimeV1`: CfuturesGetTimeV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CfuturesGetTimeV1`: %v\n", resp)
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
	resp, r, err := apiClient.MarketDataAPI.CfuturesGetTradesV1(context.Background()).Symbol(symbol).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CfuturesGetTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetTradesV1`: []CfuturesGetTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CfuturesGetTradesV1`: %v\n", resp)
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

