# \MarketDataAPI

All URIs are relative to *https://fapi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UmfuturesGetAggTradesV1**](MarketDataAPI.md#UmfuturesGetAggTradesV1) | **Get** /fapi/v1/aggTrades | Compressed/Aggregate Trades List
[**UmfuturesGetAssetIndexV1**](MarketDataAPI.md#UmfuturesGetAssetIndexV1) | **Get** /fapi/v1/assetIndex | Multi-Assets Mode Asset Index
[**UmfuturesGetConstituentsV1**](MarketDataAPI.md#UmfuturesGetConstituentsV1) | **Get** /fapi/v1/constituents | Query Index Price Constituents
[**UmfuturesGetContinuousKlinesV1**](MarketDataAPI.md#UmfuturesGetContinuousKlinesV1) | **Get** /fapi/v1/continuousKlines | Continuous Contract Kline/Candlestick Data
[**UmfuturesGetDepthV1**](MarketDataAPI.md#UmfuturesGetDepthV1) | **Get** /fapi/v1/depth | Order Book
[**UmfuturesGetExchangeInfoV1**](MarketDataAPI.md#UmfuturesGetExchangeInfoV1) | **Get** /fapi/v1/exchangeInfo | Exchange Information
[**UmfuturesGetFundingInfoV1**](MarketDataAPI.md#UmfuturesGetFundingInfoV1) | **Get** /fapi/v1/fundingInfo | Get Funding Rate Info
[**UmfuturesGetFundingRateV1**](MarketDataAPI.md#UmfuturesGetFundingRateV1) | **Get** /fapi/v1/fundingRate | Get Funding Rate History
[**UmfuturesGetFuturesDataBasis**](MarketDataAPI.md#UmfuturesGetFuturesDataBasis) | **Get** /futures/data/basis | Basis
[**UmfuturesGetFuturesDataDeliveryPrice**](MarketDataAPI.md#UmfuturesGetFuturesDataDeliveryPrice) | **Get** /futures/data/delivery-price | Quarterly Contract Settlement Price
[**UmfuturesGetFuturesDataGlobalLongShortAccountRatio**](MarketDataAPI.md#UmfuturesGetFuturesDataGlobalLongShortAccountRatio) | **Get** /futures/data/globalLongShortAccountRatio | Long/Short Ratio
[**UmfuturesGetFuturesDataOpenInterestHist**](MarketDataAPI.md#UmfuturesGetFuturesDataOpenInterestHist) | **Get** /futures/data/openInterestHist | Open Interest Statistics
[**UmfuturesGetFuturesDataTakerlongshortRatio**](MarketDataAPI.md#UmfuturesGetFuturesDataTakerlongshortRatio) | **Get** /futures/data/takerlongshortRatio | Taker Buy/Sell Volume
[**UmfuturesGetFuturesDataTopLongShortAccountRatio**](MarketDataAPI.md#UmfuturesGetFuturesDataTopLongShortAccountRatio) | **Get** /futures/data/topLongShortAccountRatio | Top Trader Long/Short Ratio (Accounts)
[**UmfuturesGetFuturesDataTopLongShortPositionRatio**](MarketDataAPI.md#UmfuturesGetFuturesDataTopLongShortPositionRatio) | **Get** /futures/data/topLongShortPositionRatio | Top Trader Long/Short Ratio (Positions)
[**UmfuturesGetHistoricalTradesV1**](MarketDataAPI.md#UmfuturesGetHistoricalTradesV1) | **Get** /fapi/v1/historicalTrades | Old Trades Lookup (MARKET_DATA)
[**UmfuturesGetIndexInfoV1**](MarketDataAPI.md#UmfuturesGetIndexInfoV1) | **Get** /fapi/v1/indexInfo | Composite Index Symbol Information
[**UmfuturesGetIndexPriceKlinesV1**](MarketDataAPI.md#UmfuturesGetIndexPriceKlinesV1) | **Get** /fapi/v1/indexPriceKlines | Index Price Kline/Candlestick Data
[**UmfuturesGetKlinesV1**](MarketDataAPI.md#UmfuturesGetKlinesV1) | **Get** /fapi/v1/klines | Kline/Candlestick Data
[**UmfuturesGetMarkPriceKlinesV1**](MarketDataAPI.md#UmfuturesGetMarkPriceKlinesV1) | **Get** /fapi/v1/markPriceKlines | Mark Price Kline/Candlestick Data
[**UmfuturesGetOpenInterestV1**](MarketDataAPI.md#UmfuturesGetOpenInterestV1) | **Get** /fapi/v1/openInterest | Open Interest
[**UmfuturesGetPingV1**](MarketDataAPI.md#UmfuturesGetPingV1) | **Get** /fapi/v1/ping | Test Connectivity
[**UmfuturesGetPremiumIndexKlinesV1**](MarketDataAPI.md#UmfuturesGetPremiumIndexKlinesV1) | **Get** /fapi/v1/premiumIndexKlines | Premium index Kline Data
[**UmfuturesGetPremiumIndexV1**](MarketDataAPI.md#UmfuturesGetPremiumIndexV1) | **Get** /fapi/v1/premiumIndex | Mark Price
[**UmfuturesGetTicker24hrV1**](MarketDataAPI.md#UmfuturesGetTicker24hrV1) | **Get** /fapi/v1/ticker/24hr | 24hr Ticker Price Change Statistics
[**UmfuturesGetTickerBookTickerV1**](MarketDataAPI.md#UmfuturesGetTickerBookTickerV1) | **Get** /fapi/v1/ticker/bookTicker | Symbol Order Book Ticker
[**UmfuturesGetTickerPriceV1**](MarketDataAPI.md#UmfuturesGetTickerPriceV1) | **Get** /fapi/v1/ticker/price | Symbol Price Ticker
[**UmfuturesGetTickerPriceV2**](MarketDataAPI.md#UmfuturesGetTickerPriceV2) | **Get** /fapi/v2/ticker/price | Symbol Price Ticker V2
[**UmfuturesGetTimeV1**](MarketDataAPI.md#UmfuturesGetTimeV1) | **Get** /fapi/v1/time | Check Server Time
[**UmfuturesGetTradesV1**](MarketDataAPI.md#UmfuturesGetTradesV1) | **Get** /fapi/v1/trades | Recent Trades List



## UmfuturesGetAggTradesV1

> []UmfuturesGetAggTradesV1RespItem UmfuturesGetAggTradesV1(ctx).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.UmfuturesGetAggTradesV1(context.Background()).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UmfuturesGetAggTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetAggTradesV1`: []UmfuturesGetAggTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UmfuturesGetAggTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetAggTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **fromId** | **int64** | ID to get aggregate trades from INCLUSIVE. | 
 **startTime** | **int64** | Timestamp in ms to get aggregate trades from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get aggregate trades until INCLUSIVE. | 
 **limit** | **int32** | Default 500; max 1000. | [default to 500]

### Return type

[**[]UmfuturesGetAggTradesV1RespItem**](UmfuturesGetAggTradesV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetAssetIndexV1

> UmfuturesGetAssetIndexV1Resp UmfuturesGetAssetIndexV1(ctx).Symbol(symbol).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.UmfuturesGetAssetIndexV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UmfuturesGetAssetIndexV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetAssetIndexV1`: UmfuturesGetAssetIndexV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UmfuturesGetAssetIndexV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetAssetIndexV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Asset pair | [default to &quot;&quot;]

### Return type

[**UmfuturesGetAssetIndexV1Resp**](UmfuturesGetAssetIndexV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetConstituentsV1

> UmfuturesGetConstituentsV1Resp UmfuturesGetConstituentsV1(ctx).Symbol(symbol).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.UmfuturesGetConstituentsV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UmfuturesGetConstituentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetConstituentsV1`: UmfuturesGetConstituentsV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UmfuturesGetConstituentsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetConstituentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**UmfuturesGetConstituentsV1Resp**](UmfuturesGetConstituentsV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetContinuousKlinesV1

> [][]UmfuturesGetContinuousKlinesV1RespInnerInner UmfuturesGetContinuousKlinesV1(ctx).Pair(pair).ContractType(contractType).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.UmfuturesGetContinuousKlinesV1(context.Background()).Pair(pair).ContractType(contractType).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UmfuturesGetContinuousKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetContinuousKlinesV1`: [][]UmfuturesGetContinuousKlinesV1RespInnerInner
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UmfuturesGetContinuousKlinesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetContinuousKlinesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pair** | **string** |  | [default to &quot;&quot;]
 **contractType** | **string** |  | [default to &quot;&quot;]
 **interval** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 1500. | [default to 500]

### Return type

[**[][]UmfuturesGetContinuousKlinesV1RespInnerInner**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetDepthV1

> UmfuturesGetDepthV1Resp UmfuturesGetDepthV1(ctx).Symbol(symbol).Limit(limit).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.UmfuturesGetDepthV1(context.Background()).Symbol(symbol).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UmfuturesGetDepthV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetDepthV1`: UmfuturesGetDepthV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UmfuturesGetDepthV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetDepthV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **limit** | **int32** | Default 500; Valid limits:[5, 10, 20, 50, 100, 500, 1000] | [default to 500]

### Return type

[**UmfuturesGetDepthV1Resp**](UmfuturesGetDepthV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetExchangeInfoV1

> UmfuturesGetExchangeInfoV1Resp UmfuturesGetExchangeInfoV1(ctx).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.UmfuturesGetExchangeInfoV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UmfuturesGetExchangeInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetExchangeInfoV1`: UmfuturesGetExchangeInfoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UmfuturesGetExchangeInfoV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetExchangeInfoV1Request struct via the builder pattern


### Return type

[**UmfuturesGetExchangeInfoV1Resp**](UmfuturesGetExchangeInfoV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetFundingInfoV1

> []UmfuturesGetFundingInfoV1RespItem UmfuturesGetFundingInfoV1(ctx).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.UmfuturesGetFundingInfoV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UmfuturesGetFundingInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetFundingInfoV1`: []UmfuturesGetFundingInfoV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UmfuturesGetFundingInfoV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetFundingInfoV1Request struct via the builder pattern


### Return type

[**[]UmfuturesGetFundingInfoV1RespItem**](UmfuturesGetFundingInfoV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetFundingRateV1

> []UmfuturesGetFundingRateV1RespItem UmfuturesGetFundingRateV1(ctx).Symbol(symbol).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.UmfuturesGetFundingRateV1(context.Background()).Symbol(symbol).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UmfuturesGetFundingRateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetFundingRateV1`: []UmfuturesGetFundingRateV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UmfuturesGetFundingRateV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetFundingRateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** | Timestamp in ms to get funding rate from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get funding rate  until INCLUSIVE. | 
 **limit** | **int32** | Default 100; max 1000 | [default to 100]

### Return type

[**[]UmfuturesGetFundingRateV1RespItem**](UmfuturesGetFundingRateV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetFuturesDataBasis

> []UmfuturesGetFuturesDataBasisRespItem UmfuturesGetFuturesDataBasis(ctx).Pair(pair).ContractType(contractType).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

Basis



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
	pair := "pair_example" // string | BTCUSDT (default to "")
	contractType := "contractType_example" // string | CURRENT_QUARTER, NEXT_QUARTER, PERPETUAL (default to "")
	period := "period_example" // string | &#34;5m&#34;,&#34;15m&#34;,&#34;30m&#34;,&#34;1h&#34;,&#34;2h&#34;,&#34;4h&#34;,&#34;6h&#34;,&#34;12h&#34;,&#34;1d&#34; (default to "")
	limit := int64(789) // int64 | Default 30,Max 500 (default to 30)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UmfuturesGetFuturesDataBasis(context.Background()).Pair(pair).ContractType(contractType).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UmfuturesGetFuturesDataBasis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetFuturesDataBasis`: []UmfuturesGetFuturesDataBasisRespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UmfuturesGetFuturesDataBasis`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetFuturesDataBasisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pair** | **string** | BTCUSDT | [default to &quot;&quot;]
 **contractType** | **string** | CURRENT_QUARTER, NEXT_QUARTER, PERPETUAL | [default to &quot;&quot;]
 **period** | **string** | &amp;#34;5m&amp;#34;,&amp;#34;15m&amp;#34;,&amp;#34;30m&amp;#34;,&amp;#34;1h&amp;#34;,&amp;#34;2h&amp;#34;,&amp;#34;4h&amp;#34;,&amp;#34;6h&amp;#34;,&amp;#34;12h&amp;#34;,&amp;#34;1d&amp;#34; | [default to &quot;&quot;]
 **limit** | **int64** | Default 30,Max 500 | [default to 30]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**[]UmfuturesGetFuturesDataBasisRespItem**](UmfuturesGetFuturesDataBasisRespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetFuturesDataDeliveryPrice

> []UmfuturesGetFuturesDataDeliveryPriceRespItem UmfuturesGetFuturesDataDeliveryPrice(ctx).Pair(pair).Execute()

Quarterly Contract Settlement Price



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
	pair := "pair_example" // string | e.g BTCUSDT (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UmfuturesGetFuturesDataDeliveryPrice(context.Background()).Pair(pair).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UmfuturesGetFuturesDataDeliveryPrice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetFuturesDataDeliveryPrice`: []UmfuturesGetFuturesDataDeliveryPriceRespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UmfuturesGetFuturesDataDeliveryPrice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetFuturesDataDeliveryPriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pair** | **string** | e.g BTCUSDT | [default to &quot;&quot;]

### Return type

[**[]UmfuturesGetFuturesDataDeliveryPriceRespItem**](UmfuturesGetFuturesDataDeliveryPriceRespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetFuturesDataGlobalLongShortAccountRatio

> []UmfuturesGetFuturesDataGlobalLongShortAccountRatioRespItem UmfuturesGetFuturesDataGlobalLongShortAccountRatio(ctx).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

Long/Short Ratio



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
	period := "period_example" // string | &#34;5m&#34;,&#34;15m&#34;,&#34;30m&#34;,&#34;1h&#34;,&#34;2h&#34;,&#34;4h&#34;,&#34;6h&#34;,&#34;12h&#34;,&#34;1d&#34; (default to "")
	limit := int64(789) // int64 | default 30, max 500 (optional) (default to 30)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UmfuturesGetFuturesDataGlobalLongShortAccountRatio(context.Background()).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UmfuturesGetFuturesDataGlobalLongShortAccountRatio``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetFuturesDataGlobalLongShortAccountRatio`: []UmfuturesGetFuturesDataGlobalLongShortAccountRatioRespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UmfuturesGetFuturesDataGlobalLongShortAccountRatio`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetFuturesDataGlobalLongShortAccountRatioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **period** | **string** | &amp;#34;5m&amp;#34;,&amp;#34;15m&amp;#34;,&amp;#34;30m&amp;#34;,&amp;#34;1h&amp;#34;,&amp;#34;2h&amp;#34;,&amp;#34;4h&amp;#34;,&amp;#34;6h&amp;#34;,&amp;#34;12h&amp;#34;,&amp;#34;1d&amp;#34; | [default to &quot;&quot;]
 **limit** | **int64** | default 30, max 500 | [default to 30]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**[]UmfuturesGetFuturesDataGlobalLongShortAccountRatioRespItem**](UmfuturesGetFuturesDataGlobalLongShortAccountRatioRespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetFuturesDataOpenInterestHist

> []UmfuturesGetFuturesDataOpenInterestHistRespItem UmfuturesGetFuturesDataOpenInterestHist(ctx).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

Open Interest Statistics



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
	period := "period_example" // string | &#34;5m&#34;,&#34;15m&#34;,&#34;30m&#34;,&#34;1h&#34;,&#34;2h&#34;,&#34;4h&#34;,&#34;6h&#34;,&#34;12h&#34;,&#34;1d&#34; (default to "")
	limit := int64(789) // int64 | default 30, max 500 (optional) (default to 30)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UmfuturesGetFuturesDataOpenInterestHist(context.Background()).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UmfuturesGetFuturesDataOpenInterestHist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetFuturesDataOpenInterestHist`: []UmfuturesGetFuturesDataOpenInterestHistRespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UmfuturesGetFuturesDataOpenInterestHist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetFuturesDataOpenInterestHistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **period** | **string** | &amp;#34;5m&amp;#34;,&amp;#34;15m&amp;#34;,&amp;#34;30m&amp;#34;,&amp;#34;1h&amp;#34;,&amp;#34;2h&amp;#34;,&amp;#34;4h&amp;#34;,&amp;#34;6h&amp;#34;,&amp;#34;12h&amp;#34;,&amp;#34;1d&amp;#34; | [default to &quot;&quot;]
 **limit** | **int64** | default 30, max 500 | [default to 30]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**[]UmfuturesGetFuturesDataOpenInterestHistRespItem**](UmfuturesGetFuturesDataOpenInterestHistRespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetFuturesDataTakerlongshortRatio

> []UmfuturesGetFuturesDataTakerlongshortRatioRespItem UmfuturesGetFuturesDataTakerlongshortRatio(ctx).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

Taker Buy/Sell Volume



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
	period := "period_example" // string | &#34;5m&#34;,&#34;15m&#34;,&#34;30m&#34;,&#34;1h&#34;,&#34;2h&#34;,&#34;4h&#34;,&#34;6h&#34;,&#34;12h&#34;,&#34;1d&#34; (default to "")
	limit := int64(789) // int64 | default 30, max 500 (optional) (default to 30)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UmfuturesGetFuturesDataTakerlongshortRatio(context.Background()).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UmfuturesGetFuturesDataTakerlongshortRatio``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetFuturesDataTakerlongshortRatio`: []UmfuturesGetFuturesDataTakerlongshortRatioRespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UmfuturesGetFuturesDataTakerlongshortRatio`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetFuturesDataTakerlongshortRatioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **period** | **string** | &amp;#34;5m&amp;#34;,&amp;#34;15m&amp;#34;,&amp;#34;30m&amp;#34;,&amp;#34;1h&amp;#34;,&amp;#34;2h&amp;#34;,&amp;#34;4h&amp;#34;,&amp;#34;6h&amp;#34;,&amp;#34;12h&amp;#34;,&amp;#34;1d&amp;#34; | [default to &quot;&quot;]
 **limit** | **int64** | default 30, max 500 | [default to 30]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**[]UmfuturesGetFuturesDataTakerlongshortRatioRespItem**](UmfuturesGetFuturesDataTakerlongshortRatioRespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetFuturesDataTopLongShortAccountRatio

> []UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem UmfuturesGetFuturesDataTopLongShortAccountRatio(ctx).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

Top Trader Long/Short Ratio (Accounts)



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
	period := "period_example" // string | &#34;5m&#34;,&#34;15m&#34;,&#34;30m&#34;,&#34;1h&#34;,&#34;2h&#34;,&#34;4h&#34;,&#34;6h&#34;,&#34;12h&#34;,&#34;1d&#34; (default to "")
	limit := int64(789) // int64 | default 30, max 500 (optional) (default to 30)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UmfuturesGetFuturesDataTopLongShortAccountRatio(context.Background()).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UmfuturesGetFuturesDataTopLongShortAccountRatio``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetFuturesDataTopLongShortAccountRatio`: []UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UmfuturesGetFuturesDataTopLongShortAccountRatio`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetFuturesDataTopLongShortAccountRatioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **period** | **string** | &amp;#34;5m&amp;#34;,&amp;#34;15m&amp;#34;,&amp;#34;30m&amp;#34;,&amp;#34;1h&amp;#34;,&amp;#34;2h&amp;#34;,&amp;#34;4h&amp;#34;,&amp;#34;6h&amp;#34;,&amp;#34;12h&amp;#34;,&amp;#34;1d&amp;#34; | [default to &quot;&quot;]
 **limit** | **int64** | default 30, max 500 | [default to 30]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**[]UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem**](UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetFuturesDataTopLongShortPositionRatio

> []UmfuturesGetFuturesDataTopLongShortPositionRatioRespItem UmfuturesGetFuturesDataTopLongShortPositionRatio(ctx).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

Top Trader Long/Short Ratio (Positions)



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
	period := "period_example" // string | &#34;5m&#34;,&#34;15m&#34;,&#34;30m&#34;,&#34;1h&#34;,&#34;2h&#34;,&#34;4h&#34;,&#34;6h&#34;,&#34;12h&#34;,&#34;1d&#34; (default to "")
	limit := int64(789) // int64 | default 30, max 500 (optional) (default to 30)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UmfuturesGetFuturesDataTopLongShortPositionRatio(context.Background()).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UmfuturesGetFuturesDataTopLongShortPositionRatio``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetFuturesDataTopLongShortPositionRatio`: []UmfuturesGetFuturesDataTopLongShortPositionRatioRespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UmfuturesGetFuturesDataTopLongShortPositionRatio`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetFuturesDataTopLongShortPositionRatioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **period** | **string** | &amp;#34;5m&amp;#34;,&amp;#34;15m&amp;#34;,&amp;#34;30m&amp;#34;,&amp;#34;1h&amp;#34;,&amp;#34;2h&amp;#34;,&amp;#34;4h&amp;#34;,&amp;#34;6h&amp;#34;,&amp;#34;12h&amp;#34;,&amp;#34;1d&amp;#34; | [default to &quot;&quot;]
 **limit** | **int64** | default 30, max 500 | [default to 30]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**[]UmfuturesGetFuturesDataTopLongShortPositionRatioRespItem**](UmfuturesGetFuturesDataTopLongShortPositionRatioRespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetHistoricalTradesV1

> []UmfuturesGetHistoricalTradesV1RespItem UmfuturesGetHistoricalTradesV1(ctx).Symbol(symbol).Limit(limit).FromId(fromId).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.UmfuturesGetHistoricalTradesV1(context.Background()).Symbol(symbol).Limit(limit).FromId(fromId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UmfuturesGetHistoricalTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetHistoricalTradesV1`: []UmfuturesGetHistoricalTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UmfuturesGetHistoricalTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetHistoricalTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **limit** | **int32** | Default 100; max 500. | [default to 100]
 **fromId** | **int64** | TradeId to fetch from. Default gets most recent trades. | 

### Return type

[**[]UmfuturesGetHistoricalTradesV1RespItem**](UmfuturesGetHistoricalTradesV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetIndexInfoV1

> []UmfuturesGetIndexInfoV1RespItem UmfuturesGetIndexInfoV1(ctx).Symbol(symbol).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.UmfuturesGetIndexInfoV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UmfuturesGetIndexInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetIndexInfoV1`: []UmfuturesGetIndexInfoV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UmfuturesGetIndexInfoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetIndexInfoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]UmfuturesGetIndexInfoV1RespItem**](UmfuturesGetIndexInfoV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetIndexPriceKlinesV1

> [][]UmfuturesGetContinuousKlinesV1RespInnerInner UmfuturesGetIndexPriceKlinesV1(ctx).Pair(pair).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.UmfuturesGetIndexPriceKlinesV1(context.Background()).Pair(pair).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UmfuturesGetIndexPriceKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetIndexPriceKlinesV1`: [][]UmfuturesGetContinuousKlinesV1RespInnerInner
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UmfuturesGetIndexPriceKlinesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetIndexPriceKlinesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pair** | **string** |  | [default to &quot;&quot;]
 **interval** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 1500. | [default to 500]

### Return type

[**[][]UmfuturesGetContinuousKlinesV1RespInnerInner**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetKlinesV1

> [][]UmfuturesGetContinuousKlinesV1RespInnerInner UmfuturesGetKlinesV1(ctx).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.UmfuturesGetKlinesV1(context.Background()).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UmfuturesGetKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetKlinesV1`: [][]UmfuturesGetContinuousKlinesV1RespInnerInner
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UmfuturesGetKlinesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetKlinesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **interval** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 1500. | [default to 500]

### Return type

[**[][]UmfuturesGetContinuousKlinesV1RespInnerInner**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetMarkPriceKlinesV1

> [][]UmfuturesGetContinuousKlinesV1RespInnerInner UmfuturesGetMarkPriceKlinesV1(ctx).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.UmfuturesGetMarkPriceKlinesV1(context.Background()).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UmfuturesGetMarkPriceKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetMarkPriceKlinesV1`: [][]UmfuturesGetContinuousKlinesV1RespInnerInner
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UmfuturesGetMarkPriceKlinesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetMarkPriceKlinesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **interval** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 1500. | [default to 500]

### Return type

[**[][]UmfuturesGetContinuousKlinesV1RespInnerInner**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetOpenInterestV1

> UmfuturesGetOpenInterestV1Resp UmfuturesGetOpenInterestV1(ctx).Symbol(symbol).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.UmfuturesGetOpenInterestV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UmfuturesGetOpenInterestV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetOpenInterestV1`: UmfuturesGetOpenInterestV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UmfuturesGetOpenInterestV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetOpenInterestV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**UmfuturesGetOpenInterestV1Resp**](UmfuturesGetOpenInterestV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetPingV1

> map[string]interface{} UmfuturesGetPingV1(ctx).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.UmfuturesGetPingV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UmfuturesGetPingV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetPingV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UmfuturesGetPingV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetPingV1Request struct via the builder pattern


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


## UmfuturesGetPremiumIndexKlinesV1

> [][]UmfuturesGetContinuousKlinesV1RespInnerInner UmfuturesGetPremiumIndexKlinesV1(ctx).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.UmfuturesGetPremiumIndexKlinesV1(context.Background()).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UmfuturesGetPremiumIndexKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetPremiumIndexKlinesV1`: [][]UmfuturesGetContinuousKlinesV1RespInnerInner
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UmfuturesGetPremiumIndexKlinesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetPremiumIndexKlinesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **interval** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 1500. | [default to 500]

### Return type

[**[][]UmfuturesGetContinuousKlinesV1RespInnerInner**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetPremiumIndexV1

> UmfuturesGetPremiumIndexV1Resp UmfuturesGetPremiumIndexV1(ctx).Symbol(symbol).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.UmfuturesGetPremiumIndexV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UmfuturesGetPremiumIndexV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetPremiumIndexV1`: UmfuturesGetPremiumIndexV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UmfuturesGetPremiumIndexV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetPremiumIndexV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**UmfuturesGetPremiumIndexV1Resp**](UmfuturesGetPremiumIndexV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetTicker24hrV1

> UmfuturesGetTicker24hrV1Resp UmfuturesGetTicker24hrV1(ctx).Symbol(symbol).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.UmfuturesGetTicker24hrV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UmfuturesGetTicker24hrV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetTicker24hrV1`: UmfuturesGetTicker24hrV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UmfuturesGetTicker24hrV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetTicker24hrV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**UmfuturesGetTicker24hrV1Resp**](UmfuturesGetTicker24hrV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetTickerBookTickerV1

> UmfuturesGetTickerBookTickerV1Resp UmfuturesGetTickerBookTickerV1(ctx).Symbol(symbol).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.UmfuturesGetTickerBookTickerV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UmfuturesGetTickerBookTickerV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetTickerBookTickerV1`: UmfuturesGetTickerBookTickerV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UmfuturesGetTickerBookTickerV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetTickerBookTickerV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**UmfuturesGetTickerBookTickerV1Resp**](UmfuturesGetTickerBookTickerV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetTickerPriceV1

> UmfuturesGetTickerPriceV1Resp UmfuturesGetTickerPriceV1(ctx).Symbol(symbol).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.UmfuturesGetTickerPriceV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UmfuturesGetTickerPriceV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetTickerPriceV1`: UmfuturesGetTickerPriceV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UmfuturesGetTickerPriceV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetTickerPriceV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**UmfuturesGetTickerPriceV1Resp**](UmfuturesGetTickerPriceV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetTickerPriceV2

> UmfuturesGetTickerPriceV2Resp UmfuturesGetTickerPriceV2(ctx).Symbol(symbol).Execute()

Symbol Price Ticker V2



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
	resp, r, err := apiClient.MarketDataAPI.UmfuturesGetTickerPriceV2(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UmfuturesGetTickerPriceV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetTickerPriceV2`: UmfuturesGetTickerPriceV2Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UmfuturesGetTickerPriceV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetTickerPriceV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**UmfuturesGetTickerPriceV2Resp**](UmfuturesGetTickerPriceV2Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetTimeV1

> UmfuturesGetTimeV1Resp UmfuturesGetTimeV1(ctx).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.UmfuturesGetTimeV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UmfuturesGetTimeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetTimeV1`: UmfuturesGetTimeV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UmfuturesGetTimeV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetTimeV1Request struct via the builder pattern


### Return type

[**UmfuturesGetTimeV1Resp**](UmfuturesGetTimeV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetTradesV1

> []UmfuturesGetTradesV1RespItem UmfuturesGetTradesV1(ctx).Symbol(symbol).Limit(limit).Execute()

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
	resp, r, err := apiClient.MarketDataAPI.UmfuturesGetTradesV1(context.Background()).Symbol(symbol).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UmfuturesGetTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetTradesV1`: []UmfuturesGetTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UmfuturesGetTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **limit** | **int32** | Default 500; max 1000. | [default to 500]

### Return type

[**[]UmfuturesGetTradesV1RespItem**](UmfuturesGetTradesV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

