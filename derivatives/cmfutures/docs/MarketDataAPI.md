# \MarketDataAPI

All URIs are relative to *https://dapi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CmfuturesGetAggTradesV1**](MarketDataAPI.md#CmfuturesGetAggTradesV1) | **Get** /dapi/v1/aggTrades | Compressed/Aggregate Trades List
[**CmfuturesGetConstituentsV1**](MarketDataAPI.md#CmfuturesGetConstituentsV1) | **Get** /dapi/v1/constituents | Query Index Price Constituents
[**CmfuturesGetContinuousKlinesV1**](MarketDataAPI.md#CmfuturesGetContinuousKlinesV1) | **Get** /dapi/v1/continuousKlines | Continuous Contract Kline/Candlestick Data
[**CmfuturesGetDepthV1**](MarketDataAPI.md#CmfuturesGetDepthV1) | **Get** /dapi/v1/depth | Order Book
[**CmfuturesGetExchangeInfoV1**](MarketDataAPI.md#CmfuturesGetExchangeInfoV1) | **Get** /dapi/v1/exchangeInfo | Exchange Information
[**CmfuturesGetFundingInfoV1**](MarketDataAPI.md#CmfuturesGetFundingInfoV1) | **Get** /dapi/v1/fundingInfo | Get Funding Rate Info
[**CmfuturesGetFundingRateV1**](MarketDataAPI.md#CmfuturesGetFundingRateV1) | **Get** /dapi/v1/fundingRate | Get Funding Rate History of Perpetual Futures
[**CmfuturesGetFuturesDataBasis**](MarketDataAPI.md#CmfuturesGetFuturesDataBasis) | **Get** /futures/data/basis | Basis
[**CmfuturesGetFuturesDataGlobalLongShortAccountRatio**](MarketDataAPI.md#CmfuturesGetFuturesDataGlobalLongShortAccountRatio) | **Get** /futures/data/globalLongShortAccountRatio | Long/Short Ratio
[**CmfuturesGetFuturesDataOpenInterestHist**](MarketDataAPI.md#CmfuturesGetFuturesDataOpenInterestHist) | **Get** /futures/data/openInterestHist | Open Interest Statistics
[**CmfuturesGetFuturesDataTakerBuySellVol**](MarketDataAPI.md#CmfuturesGetFuturesDataTakerBuySellVol) | **Get** /futures/data/takerBuySellVol | Taker Buy/Sell Volume
[**CmfuturesGetFuturesDataTopLongShortAccountRatio**](MarketDataAPI.md#CmfuturesGetFuturesDataTopLongShortAccountRatio) | **Get** /futures/data/topLongShortAccountRatio | Top Trader Long/Short Ratio (Accounts)
[**CmfuturesGetFuturesDataTopLongShortPositionRatio**](MarketDataAPI.md#CmfuturesGetFuturesDataTopLongShortPositionRatio) | **Get** /futures/data/topLongShortPositionRatio | Top Trader Long/Short Ratio (Positions)
[**CmfuturesGetHistoricalTradesV1**](MarketDataAPI.md#CmfuturesGetHistoricalTradesV1) | **Get** /dapi/v1/historicalTrades | Old Trades Lookup(MARKET_DATA)
[**CmfuturesGetIndexPriceKlinesV1**](MarketDataAPI.md#CmfuturesGetIndexPriceKlinesV1) | **Get** /dapi/v1/indexPriceKlines | Index Price Kline/Candlestick Data
[**CmfuturesGetKlinesV1**](MarketDataAPI.md#CmfuturesGetKlinesV1) | **Get** /dapi/v1/klines | Kline/Candlestick Data
[**CmfuturesGetMarkPriceKlinesV1**](MarketDataAPI.md#CmfuturesGetMarkPriceKlinesV1) | **Get** /dapi/v1/markPriceKlines | Mark Price Kline/Candlestick Data
[**CmfuturesGetOpenInterestV1**](MarketDataAPI.md#CmfuturesGetOpenInterestV1) | **Get** /dapi/v1/openInterest | Open Interest
[**CmfuturesGetPingV1**](MarketDataAPI.md#CmfuturesGetPingV1) | **Get** /dapi/v1/ping | Test Connectivity
[**CmfuturesGetPremiumIndexKlinesV1**](MarketDataAPI.md#CmfuturesGetPremiumIndexKlinesV1) | **Get** /dapi/v1/premiumIndexKlines | Premium index Kline Data
[**CmfuturesGetPremiumIndexV1**](MarketDataAPI.md#CmfuturesGetPremiumIndexV1) | **Get** /dapi/v1/premiumIndex | Index Price and Mark Price
[**CmfuturesGetTicker24hrV1**](MarketDataAPI.md#CmfuturesGetTicker24hrV1) | **Get** /dapi/v1/ticker/24hr | 24hr Ticker Price Change Statistics
[**CmfuturesGetTickerBookTickerV1**](MarketDataAPI.md#CmfuturesGetTickerBookTickerV1) | **Get** /dapi/v1/ticker/bookTicker | Symbol Order Book Ticker
[**CmfuturesGetTickerPriceV1**](MarketDataAPI.md#CmfuturesGetTickerPriceV1) | **Get** /dapi/v1/ticker/price | Symbol Price Ticker
[**CmfuturesGetTimeV1**](MarketDataAPI.md#CmfuturesGetTimeV1) | **Get** /dapi/v1/time | Check Server time
[**CmfuturesGetTradesV1**](MarketDataAPI.md#CmfuturesGetTradesV1) | **Get** /dapi/v1/trades | Recent Trades List



## CmfuturesGetAggTradesV1

> []CmfuturesGetAggTradesV1RespItem CmfuturesGetAggTradesV1(ctx).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.CmfuturesGetAggTradesV1(context.Background()).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CmfuturesGetAggTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CmfuturesGetAggTradesV1`: []CmfuturesGetAggTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CmfuturesGetAggTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCmfuturesGetAggTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **fromId** | **int64** | ID to get aggregate trades from INCLUSIVE. | 
 **startTime** | **int64** | Timestamp in ms to get aggregate trades from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get aggregate trades until INCLUSIVE. | 
 **limit** | **int32** | Default 500; max 1000. | [default to 500]

### Return type

[**[]CmfuturesGetAggTradesV1RespItem**](CmfuturesGetAggTradesV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CmfuturesGetConstituentsV1

> CmfuturesGetConstituentsV1Resp CmfuturesGetConstituentsV1(ctx).Symbol(symbol).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.CmfuturesGetConstituentsV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CmfuturesGetConstituentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CmfuturesGetConstituentsV1`: CmfuturesGetConstituentsV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CmfuturesGetConstituentsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCmfuturesGetConstituentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**CmfuturesGetConstituentsV1Resp**](CmfuturesGetConstituentsV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CmfuturesGetContinuousKlinesV1

> [][]CmfuturesGetContinuousKlinesV1RespInnerInner CmfuturesGetContinuousKlinesV1(ctx).Pair(pair).ContractType(contractType).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.CmfuturesGetContinuousKlinesV1(context.Background()).Pair(pair).ContractType(contractType).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CmfuturesGetContinuousKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CmfuturesGetContinuousKlinesV1`: [][]CmfuturesGetContinuousKlinesV1RespInnerInner
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CmfuturesGetContinuousKlinesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCmfuturesGetContinuousKlinesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pair** | **string** |  | [default to &quot;&quot;]
 **contractType** | **string** |  | [default to &quot;&quot;]
 **interval** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 1500. | [default to 500]

### Return type

[**[][]CmfuturesGetContinuousKlinesV1RespInnerInner**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CmfuturesGetDepthV1

> CmfuturesGetDepthV1Resp CmfuturesGetDepthV1(ctx).Symbol(symbol).Limit(limit).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.CmfuturesGetDepthV1(context.Background()).Symbol(symbol).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CmfuturesGetDepthV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CmfuturesGetDepthV1`: CmfuturesGetDepthV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CmfuturesGetDepthV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCmfuturesGetDepthV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **limit** | **int32** | Default 500; Valid limits:[5, 10, 20, 50, 100, 500, 1000] | [default to 500]

### Return type

[**CmfuturesGetDepthV1Resp**](CmfuturesGetDepthV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CmfuturesGetExchangeInfoV1

> CmfuturesGetExchangeInfoV1Resp CmfuturesGetExchangeInfoV1(ctx).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.CmfuturesGetExchangeInfoV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CmfuturesGetExchangeInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CmfuturesGetExchangeInfoV1`: CmfuturesGetExchangeInfoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CmfuturesGetExchangeInfoV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCmfuturesGetExchangeInfoV1Request struct via the builder pattern


### Return type

[**CmfuturesGetExchangeInfoV1Resp**](CmfuturesGetExchangeInfoV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CmfuturesGetFundingInfoV1

> []CmfuturesGetFundingInfoV1RespItem CmfuturesGetFundingInfoV1(ctx).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.CmfuturesGetFundingInfoV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CmfuturesGetFundingInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CmfuturesGetFundingInfoV1`: []CmfuturesGetFundingInfoV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CmfuturesGetFundingInfoV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCmfuturesGetFundingInfoV1Request struct via the builder pattern


### Return type

[**[]CmfuturesGetFundingInfoV1RespItem**](CmfuturesGetFundingInfoV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CmfuturesGetFundingRateV1

> []CmfuturesGetFundingRateV1RespItem CmfuturesGetFundingRateV1(ctx).Symbol(symbol).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.CmfuturesGetFundingRateV1(context.Background()).Symbol(symbol).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CmfuturesGetFundingRateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CmfuturesGetFundingRateV1`: []CmfuturesGetFundingRateV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CmfuturesGetFundingRateV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCmfuturesGetFundingRateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** | Timestamp in ms to get funding rate from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get funding rate  until INCLUSIVE. | 
 **limit** | **int32** | Default 100; max 1000 | [default to 100]

### Return type

[**[]CmfuturesGetFundingRateV1RespItem**](CmfuturesGetFundingRateV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CmfuturesGetFuturesDataBasis

> []CmfuturesGetFuturesDataBasisRespItem CmfuturesGetFuturesDataBasis(ctx).Pair(pair).ContractType(contractType).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.CmfuturesGetFuturesDataBasis(context.Background()).Pair(pair).ContractType(contractType).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CmfuturesGetFuturesDataBasis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CmfuturesGetFuturesDataBasis`: []CmfuturesGetFuturesDataBasisRespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CmfuturesGetFuturesDataBasis`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCmfuturesGetFuturesDataBasisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pair** | **string** | BTCUSD | [default to &quot;&quot;]
 **contractType** | **string** | CURRENT_QUARTER, NEXT_QUARTER, PERPETUAL | [default to &quot;&quot;]
 **period** | **string** | &amp;#34;5m&amp;#34;,&amp;#34;15m&amp;#34;,&amp;#34;30m&amp;#34;,&amp;#34;1h&amp;#34;,&amp;#34;2h&amp;#34;,&amp;#34;4h&amp;#34;,&amp;#34;6h&amp;#34;,&amp;#34;12h&amp;#34;,&amp;#34;1d&amp;#34; | [default to &quot;&quot;]
 **limit** | **int64** | Default 30,Max 500 | [default to 30]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**[]CmfuturesGetFuturesDataBasisRespItem**](CmfuturesGetFuturesDataBasisRespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CmfuturesGetFuturesDataGlobalLongShortAccountRatio

> []CmfuturesGetFuturesDataGlobalLongShortAccountRatioRespItem CmfuturesGetFuturesDataGlobalLongShortAccountRatio(ctx).Pair(pair).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.CmfuturesGetFuturesDataGlobalLongShortAccountRatio(context.Background()).Pair(pair).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CmfuturesGetFuturesDataGlobalLongShortAccountRatio``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CmfuturesGetFuturesDataGlobalLongShortAccountRatio`: []CmfuturesGetFuturesDataGlobalLongShortAccountRatioRespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CmfuturesGetFuturesDataGlobalLongShortAccountRatio`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCmfuturesGetFuturesDataGlobalLongShortAccountRatioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pair** | **string** | BTCUSD | [default to &quot;&quot;]
 **period** | **string** | &amp;#34;5m&amp;#34;,&amp;#34;15m&amp;#34;,&amp;#34;30m&amp;#34;,&amp;#34;1h&amp;#34;,&amp;#34;2h&amp;#34;,&amp;#34;4h&amp;#34;,&amp;#34;6h&amp;#34;,&amp;#34;12h&amp;#34;,&amp;#34;1d&amp;#34; | [default to &quot;&quot;]
 **limit** | **int64** | Default 30,Max 500 | [default to 30]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**[]CmfuturesGetFuturesDataGlobalLongShortAccountRatioRespItem**](CmfuturesGetFuturesDataGlobalLongShortAccountRatioRespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CmfuturesGetFuturesDataOpenInterestHist

> []CmfuturesGetFuturesDataOpenInterestHistRespItem CmfuturesGetFuturesDataOpenInterestHist(ctx).Pair(pair).ContractType(contractType).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.CmfuturesGetFuturesDataOpenInterestHist(context.Background()).Pair(pair).ContractType(contractType).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CmfuturesGetFuturesDataOpenInterestHist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CmfuturesGetFuturesDataOpenInterestHist`: []CmfuturesGetFuturesDataOpenInterestHistRespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CmfuturesGetFuturesDataOpenInterestHist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCmfuturesGetFuturesDataOpenInterestHistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pair** | **string** | BTCUSD | [default to &quot;&quot;]
 **contractType** | **string** | ALL, CURRENT_QUARTER, NEXT_QUARTER, PERPETUAL | [default to &quot;&quot;]
 **period** | **string** | &amp;#34;5m&amp;#34;,&amp;#34;15m&amp;#34;,&amp;#34;30m&amp;#34;,&amp;#34;1h&amp;#34;,&amp;#34;2h&amp;#34;,&amp;#34;4h&amp;#34;,&amp;#34;6h&amp;#34;,&amp;#34;12h&amp;#34;,&amp;#34;1d&amp;#34; | [default to &quot;&quot;]
 **limit** | **int64** | Default 30,Max 500 | [default to 30]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**[]CmfuturesGetFuturesDataOpenInterestHistRespItem**](CmfuturesGetFuturesDataOpenInterestHistRespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CmfuturesGetFuturesDataTakerBuySellVol

> []CmfuturesGetFuturesDataTakerBuySellVolRespItem CmfuturesGetFuturesDataTakerBuySellVol(ctx).Pair(pair).ContractType(contractType).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.CmfuturesGetFuturesDataTakerBuySellVol(context.Background()).Pair(pair).ContractType(contractType).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CmfuturesGetFuturesDataTakerBuySellVol``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CmfuturesGetFuturesDataTakerBuySellVol`: []CmfuturesGetFuturesDataTakerBuySellVolRespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CmfuturesGetFuturesDataTakerBuySellVol`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCmfuturesGetFuturesDataTakerBuySellVolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pair** | **string** | BTCUSD | [default to &quot;&quot;]
 **contractType** | **string** | ALL, CURRENT_QUARTER, NEXT_QUARTER, PERPETUAL | [default to &quot;&quot;]
 **period** | **string** | &amp;#34;5m&amp;#34;,&amp;#34;15m&amp;#34;,&amp;#34;30m&amp;#34;,&amp;#34;1h&amp;#34;,&amp;#34;2h&amp;#34;,&amp;#34;4h&amp;#34;,&amp;#34;6h&amp;#34;,&amp;#34;12h&amp;#34;,&amp;#34;1d&amp;#34; | [default to &quot;&quot;]
 **limit** | **int64** | Default 30,Max 500 | [default to 30]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**[]CmfuturesGetFuturesDataTakerBuySellVolRespItem**](CmfuturesGetFuturesDataTakerBuySellVolRespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CmfuturesGetFuturesDataTopLongShortAccountRatio

> []CmfuturesGetFuturesDataTopLongShortAccountRatioRespItem CmfuturesGetFuturesDataTopLongShortAccountRatio(ctx).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.CmfuturesGetFuturesDataTopLongShortAccountRatio(context.Background()).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CmfuturesGetFuturesDataTopLongShortAccountRatio``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CmfuturesGetFuturesDataTopLongShortAccountRatio`: []CmfuturesGetFuturesDataTopLongShortAccountRatioRespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CmfuturesGetFuturesDataTopLongShortAccountRatio`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCmfuturesGetFuturesDataTopLongShortAccountRatioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **period** | **string** | &amp;#34;5m&amp;#34;,&amp;#34;15m&amp;#34;,&amp;#34;30m&amp;#34;,&amp;#34;1h&amp;#34;,&amp;#34;2h&amp;#34;,&amp;#34;4h&amp;#34;,&amp;#34;6h&amp;#34;,&amp;#34;12h&amp;#34;,&amp;#34;1d&amp;#34; | [default to &quot;&quot;]
 **limit** | **int64** | default 30, max 500 | [default to 30]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**[]CmfuturesGetFuturesDataTopLongShortAccountRatioRespItem**](CmfuturesGetFuturesDataTopLongShortAccountRatioRespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CmfuturesGetFuturesDataTopLongShortPositionRatio

> []CmfuturesGetFuturesDataTopLongShortPositionRatioRespItem CmfuturesGetFuturesDataTopLongShortPositionRatio(ctx).Pair(pair).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.CmfuturesGetFuturesDataTopLongShortPositionRatio(context.Background()).Pair(pair).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CmfuturesGetFuturesDataTopLongShortPositionRatio``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CmfuturesGetFuturesDataTopLongShortPositionRatio`: []CmfuturesGetFuturesDataTopLongShortPositionRatioRespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CmfuturesGetFuturesDataTopLongShortPositionRatio`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCmfuturesGetFuturesDataTopLongShortPositionRatioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pair** | **string** | BTCUSD | [default to &quot;&quot;]
 **period** | **string** | &amp;#34;5m&amp;#34;,&amp;#34;15m&amp;#34;,&amp;#34;30m&amp;#34;,&amp;#34;1h&amp;#34;,&amp;#34;2h&amp;#34;,&amp;#34;4h&amp;#34;,&amp;#34;6h&amp;#34;,&amp;#34;12h&amp;#34;,&amp;#34;1d&amp;#34; | [default to &quot;&quot;]
 **limit** | **int64** | Default 30,Max 500 | [default to 30]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**[]CmfuturesGetFuturesDataTopLongShortPositionRatioRespItem**](CmfuturesGetFuturesDataTopLongShortPositionRatioRespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CmfuturesGetHistoricalTradesV1

> []CmfuturesGetHistoricalTradesV1RespItem CmfuturesGetHistoricalTradesV1(ctx).Symbol(symbol).Limit(limit).FromId(fromId).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.CmfuturesGetHistoricalTradesV1(context.Background()).Symbol(symbol).Limit(limit).FromId(fromId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CmfuturesGetHistoricalTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CmfuturesGetHistoricalTradesV1`: []CmfuturesGetHistoricalTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CmfuturesGetHistoricalTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCmfuturesGetHistoricalTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **limit** | **int32** | Default 100; max 500. | [default to 100]
 **fromId** | **int64** | TradeId to fetch from. Default gets most recent trades. | 

### Return type

[**[]CmfuturesGetHistoricalTradesV1RespItem**](CmfuturesGetHistoricalTradesV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CmfuturesGetIndexPriceKlinesV1

> [][]CmfuturesGetContinuousKlinesV1RespInnerInner CmfuturesGetIndexPriceKlinesV1(ctx).Pair(pair).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.CmfuturesGetIndexPriceKlinesV1(context.Background()).Pair(pair).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CmfuturesGetIndexPriceKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CmfuturesGetIndexPriceKlinesV1`: [][]CmfuturesGetContinuousKlinesV1RespInnerInner
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CmfuturesGetIndexPriceKlinesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCmfuturesGetIndexPriceKlinesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pair** | **string** |  | [default to &quot;&quot;]
 **interval** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 1500. | [default to 500]

### Return type

[**[][]CmfuturesGetContinuousKlinesV1RespInnerInner**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CmfuturesGetKlinesV1

> [][]CmfuturesGetContinuousKlinesV1RespInnerInner CmfuturesGetKlinesV1(ctx).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.CmfuturesGetKlinesV1(context.Background()).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CmfuturesGetKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CmfuturesGetKlinesV1`: [][]CmfuturesGetContinuousKlinesV1RespInnerInner
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CmfuturesGetKlinesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCmfuturesGetKlinesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **interval** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 1500. | [default to 500]

### Return type

[**[][]CmfuturesGetContinuousKlinesV1RespInnerInner**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CmfuturesGetMarkPriceKlinesV1

> [][]CmfuturesGetContinuousKlinesV1RespInnerInner CmfuturesGetMarkPriceKlinesV1(ctx).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.CmfuturesGetMarkPriceKlinesV1(context.Background()).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CmfuturesGetMarkPriceKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CmfuturesGetMarkPriceKlinesV1`: [][]CmfuturesGetContinuousKlinesV1RespInnerInner
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CmfuturesGetMarkPriceKlinesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCmfuturesGetMarkPriceKlinesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **interval** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 1500. | [default to 500]

### Return type

[**[][]CmfuturesGetContinuousKlinesV1RespInnerInner**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CmfuturesGetOpenInterestV1

> CmfuturesGetOpenInterestV1Resp CmfuturesGetOpenInterestV1(ctx).Symbol(symbol).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.CmfuturesGetOpenInterestV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CmfuturesGetOpenInterestV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CmfuturesGetOpenInterestV1`: CmfuturesGetOpenInterestV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CmfuturesGetOpenInterestV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCmfuturesGetOpenInterestV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**CmfuturesGetOpenInterestV1Resp**](CmfuturesGetOpenInterestV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CmfuturesGetPingV1

> map[string]interface{} CmfuturesGetPingV1(ctx).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.CmfuturesGetPingV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CmfuturesGetPingV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CmfuturesGetPingV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CmfuturesGetPingV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCmfuturesGetPingV1Request struct via the builder pattern


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


## CmfuturesGetPremiumIndexKlinesV1

> [][]CmfuturesGetContinuousKlinesV1RespInnerInner CmfuturesGetPremiumIndexKlinesV1(ctx).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.CmfuturesGetPremiumIndexKlinesV1(context.Background()).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CmfuturesGetPremiumIndexKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CmfuturesGetPremiumIndexKlinesV1`: [][]CmfuturesGetContinuousKlinesV1RespInnerInner
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CmfuturesGetPremiumIndexKlinesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCmfuturesGetPremiumIndexKlinesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **interval** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 1500. | [default to 500]

### Return type

[**[][]CmfuturesGetContinuousKlinesV1RespInnerInner**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CmfuturesGetPremiumIndexV1

> []CmfuturesGetPremiumIndexV1RespItem CmfuturesGetPremiumIndexV1(ctx).Symbol(symbol).Pair(pair).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.CmfuturesGetPremiumIndexV1(context.Background()).Symbol(symbol).Pair(pair).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CmfuturesGetPremiumIndexV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CmfuturesGetPremiumIndexV1`: []CmfuturesGetPremiumIndexV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CmfuturesGetPremiumIndexV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCmfuturesGetPremiumIndexV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **pair** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]CmfuturesGetPremiumIndexV1RespItem**](CmfuturesGetPremiumIndexV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CmfuturesGetTicker24hrV1

> []CmfuturesGetTicker24hrV1RespItem CmfuturesGetTicker24hrV1(ctx).Symbol(symbol).Pair(pair).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.CmfuturesGetTicker24hrV1(context.Background()).Symbol(symbol).Pair(pair).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CmfuturesGetTicker24hrV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CmfuturesGetTicker24hrV1`: []CmfuturesGetTicker24hrV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CmfuturesGetTicker24hrV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCmfuturesGetTicker24hrV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **pair** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]CmfuturesGetTicker24hrV1RespItem**](CmfuturesGetTicker24hrV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CmfuturesGetTickerBookTickerV1

> []CmfuturesGetTickerBookTickerV1RespItem CmfuturesGetTickerBookTickerV1(ctx).Symbol(symbol).Pair(pair).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.CmfuturesGetTickerBookTickerV1(context.Background()).Symbol(symbol).Pair(pair).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CmfuturesGetTickerBookTickerV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CmfuturesGetTickerBookTickerV1`: []CmfuturesGetTickerBookTickerV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CmfuturesGetTickerBookTickerV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCmfuturesGetTickerBookTickerV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **pair** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]CmfuturesGetTickerBookTickerV1RespItem**](CmfuturesGetTickerBookTickerV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CmfuturesGetTickerPriceV1

> []CmfuturesGetTickerPriceV1RespItem CmfuturesGetTickerPriceV1(ctx).Symbol(symbol).Pair(pair).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.CmfuturesGetTickerPriceV1(context.Background()).Symbol(symbol).Pair(pair).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CmfuturesGetTickerPriceV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CmfuturesGetTickerPriceV1`: []CmfuturesGetTickerPriceV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CmfuturesGetTickerPriceV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCmfuturesGetTickerPriceV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **pair** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]CmfuturesGetTickerPriceV1RespItem**](CmfuturesGetTickerPriceV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CmfuturesGetTimeV1

> CmfuturesGetTimeV1Resp CmfuturesGetTimeV1(ctx).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.CmfuturesGetTimeV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CmfuturesGetTimeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CmfuturesGetTimeV1`: CmfuturesGetTimeV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CmfuturesGetTimeV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCmfuturesGetTimeV1Request struct via the builder pattern


### Return type

[**CmfuturesGetTimeV1Resp**](CmfuturesGetTimeV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CmfuturesGetTradesV1

> []CmfuturesGetTradesV1RespItem CmfuturesGetTradesV1(ctx).Symbol(symbol).Limit(limit).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.CmfuturesGetTradesV1(context.Background()).Symbol(symbol).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.CmfuturesGetTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CmfuturesGetTradesV1`: []CmfuturesGetTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.CmfuturesGetTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCmfuturesGetTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **limit** | **int32** | Default 500; max 1000. | [default to 500]

### Return type

[**[]CmfuturesGetTradesV1RespItem**](CmfuturesGetTradesV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

