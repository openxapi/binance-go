# \MarketDataAPI

All URIs are relative to *https://fapi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UfuturesGetAggTradesV1**](MarketDataAPI.md#UfuturesGetAggTradesV1) | **Get** /fapi/v1/aggTrades | Compressed/Aggregate Trades List
[**UfuturesGetAssetIndexV1**](MarketDataAPI.md#UfuturesGetAssetIndexV1) | **Get** /fapi/v1/assetIndex | Multi-Assets Mode Asset Index
[**UfuturesGetConstituentsV1**](MarketDataAPI.md#UfuturesGetConstituentsV1) | **Get** /fapi/v1/constituents | Query Index Price Constituents
[**UfuturesGetContinuousKlinesV1**](MarketDataAPI.md#UfuturesGetContinuousKlinesV1) | **Get** /fapi/v1/continuousKlines | Continuous Contract Kline/Candlestick Data
[**UfuturesGetDepthV1**](MarketDataAPI.md#UfuturesGetDepthV1) | **Get** /fapi/v1/depth | Order Book
[**UfuturesGetExchangeInfoV1**](MarketDataAPI.md#UfuturesGetExchangeInfoV1) | **Get** /fapi/v1/exchangeInfo | Exchange Information
[**UfuturesGetFundingInfoV1**](MarketDataAPI.md#UfuturesGetFundingInfoV1) | **Get** /fapi/v1/fundingInfo | Get Funding Rate Info
[**UfuturesGetFundingRateV1**](MarketDataAPI.md#UfuturesGetFundingRateV1) | **Get** /fapi/v1/fundingRate | Get Funding Rate History
[**UfuturesGetFuturesDataBasis**](MarketDataAPI.md#UfuturesGetFuturesDataBasis) | **Get** /futures/data/basis | Basis
[**UfuturesGetFuturesDataDeliveryPrice**](MarketDataAPI.md#UfuturesGetFuturesDataDeliveryPrice) | **Get** /futures/data/delivery-price | Quarterly Contract Settlement Price
[**UfuturesGetFuturesDataGlobalLongShortAccountRatio**](MarketDataAPI.md#UfuturesGetFuturesDataGlobalLongShortAccountRatio) | **Get** /futures/data/globalLongShortAccountRatio | Long/Short Ratio
[**UfuturesGetFuturesDataOpenInterestHist**](MarketDataAPI.md#UfuturesGetFuturesDataOpenInterestHist) | **Get** /futures/data/openInterestHist | Open Interest Statistics
[**UfuturesGetFuturesDataTakerlongshortRatio**](MarketDataAPI.md#UfuturesGetFuturesDataTakerlongshortRatio) | **Get** /futures/data/takerlongshortRatio | Taker Buy/Sell Volume
[**UfuturesGetFuturesDataTopLongShortAccountRatio**](MarketDataAPI.md#UfuturesGetFuturesDataTopLongShortAccountRatio) | **Get** /futures/data/topLongShortAccountRatio | Top Trader Long/Short Ratio (Accounts)
[**UfuturesGetFuturesDataTopLongShortPositionRatio**](MarketDataAPI.md#UfuturesGetFuturesDataTopLongShortPositionRatio) | **Get** /futures/data/topLongShortPositionRatio | Top Trader Long/Short Ratio (Positions)
[**UfuturesGetHistoricalTradesV1**](MarketDataAPI.md#UfuturesGetHistoricalTradesV1) | **Get** /fapi/v1/historicalTrades | Old Trades Lookup (MARKET_DATA)
[**UfuturesGetIndexInfoV1**](MarketDataAPI.md#UfuturesGetIndexInfoV1) | **Get** /fapi/v1/indexInfo | Composite Index Symbol Information
[**UfuturesGetIndexPriceKlinesV1**](MarketDataAPI.md#UfuturesGetIndexPriceKlinesV1) | **Get** /fapi/v1/indexPriceKlines | Index Price Kline/Candlestick Data
[**UfuturesGetKlinesV1**](MarketDataAPI.md#UfuturesGetKlinesV1) | **Get** /fapi/v1/klines | Kline/Candlestick Data
[**UfuturesGetMarkPriceKlinesV1**](MarketDataAPI.md#UfuturesGetMarkPriceKlinesV1) | **Get** /fapi/v1/markPriceKlines | Mark Price Kline/Candlestick Data
[**UfuturesGetOpenInterestV1**](MarketDataAPI.md#UfuturesGetOpenInterestV1) | **Get** /fapi/v1/openInterest | Open Interest
[**UfuturesGetPingV1**](MarketDataAPI.md#UfuturesGetPingV1) | **Get** /fapi/v1/ping | Test Connectivity
[**UfuturesGetPremiumIndexKlinesV1**](MarketDataAPI.md#UfuturesGetPremiumIndexKlinesV1) | **Get** /fapi/v1/premiumIndexKlines | Premium index Kline Data
[**UfuturesGetPremiumIndexV1**](MarketDataAPI.md#UfuturesGetPremiumIndexV1) | **Get** /fapi/v1/premiumIndex | Mark Price
[**UfuturesGetTicker24hrV1**](MarketDataAPI.md#UfuturesGetTicker24hrV1) | **Get** /fapi/v1/ticker/24hr | 24hr Ticker Price Change Statistics
[**UfuturesGetTickerBookTickerV1**](MarketDataAPI.md#UfuturesGetTickerBookTickerV1) | **Get** /fapi/v1/ticker/bookTicker | Symbol Order Book Ticker
[**UfuturesGetTickerPriceV1**](MarketDataAPI.md#UfuturesGetTickerPriceV1) | **Get** /fapi/v1/ticker/price | Symbol Price Ticker
[**UfuturesGetTickerPriceV2**](MarketDataAPI.md#UfuturesGetTickerPriceV2) | **Get** /fapi/v2/ticker/price | Symbol Price Ticker V2
[**UfuturesGetTimeV1**](MarketDataAPI.md#UfuturesGetTimeV1) | **Get** /fapi/v1/time | Check Server Time
[**UfuturesGetTradesV1**](MarketDataAPI.md#UfuturesGetTradesV1) | **Get** /fapi/v1/trades | Recent Trades List



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
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	fromId := int64(789) // int64 | ID to get aggregate trades from INCLUSIVE. (optional)
	startTime := int64(789) // int64 | Timestamp in ms to get aggregate trades from INCLUSIVE. (optional)
	endTime := int64(789) // int64 | Timestamp in ms to get aggregate trades until INCLUSIVE. (optional)
	limit := int32(56) // int32 | Default 500; max 1000. (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UfuturesGetAggTradesV1(context.Background()).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UfuturesGetAggTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetAggTradesV1`: []UfuturesGetAggTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UfuturesGetAggTradesV1`: %v\n", resp)
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
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {
	symbol := "symbol_example" // string | Asset pair (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UfuturesGetAssetIndexV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UfuturesGetAssetIndexV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetAssetIndexV1`: UfuturesGetAssetIndexV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UfuturesGetAssetIndexV1`: %v\n", resp)
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
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UfuturesGetConstituentsV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UfuturesGetConstituentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetConstituentsV1`: UfuturesGetConstituentsV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UfuturesGetConstituentsV1`: %v\n", resp)
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

> [][]UfuturesGetContinuousKlinesV1200ResponseInnerInner UfuturesGetContinuousKlinesV1(ctx).Pair(pair).ContractType(contractType).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Continuous Contract Kline/Candlestick Data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/ufutures"
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
	resp, r, err := apiClient.MarketDataAPI.UfuturesGetContinuousKlinesV1(context.Background()).Pair(pair).ContractType(contractType).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UfuturesGetContinuousKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetContinuousKlinesV1`: [][]UfuturesGetContinuousKlinesV1200ResponseInnerInner
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UfuturesGetContinuousKlinesV1`: %v\n", resp)
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

[**[][]UfuturesGetContinuousKlinesV1200ResponseInnerInner**](array.md)

### Authorization

No authorization required

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
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	limit := int32(56) // int32 | Default 500; Valid limits:[5, 10, 20, 50, 100, 500, 1000] (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UfuturesGetDepthV1(context.Background()).Symbol(symbol).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UfuturesGetDepthV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetDepthV1`: UfuturesGetDepthV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UfuturesGetDepthV1`: %v\n", resp)
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
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UfuturesGetExchangeInfoV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UfuturesGetExchangeInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetExchangeInfoV1`: UfuturesGetExchangeInfoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UfuturesGetExchangeInfoV1`: %v\n", resp)
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
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UfuturesGetFundingInfoV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UfuturesGetFundingInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetFundingInfoV1`: []UfuturesGetFundingInfoV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UfuturesGetFundingInfoV1`: %v\n", resp)
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
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {
	symbol := "symbol_example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 | Timestamp in ms to get funding rate from INCLUSIVE. (optional)
	endTime := int64(789) // int64 | Timestamp in ms to get funding rate  until INCLUSIVE. (optional)
	limit := int32(56) // int32 | Default 100; max 1000 (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UfuturesGetFundingRateV1(context.Background()).Symbol(symbol).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UfuturesGetFundingRateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetFundingRateV1`: []UfuturesGetFundingRateV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UfuturesGetFundingRateV1`: %v\n", resp)
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


## UfuturesGetFuturesDataBasis

> []UfuturesGetFuturesDataBasisRespItem UfuturesGetFuturesDataBasis(ctx).Pair(pair).ContractType(contractType).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

Basis



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/ufutures"
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
	resp, r, err := apiClient.MarketDataAPI.UfuturesGetFuturesDataBasis(context.Background()).Pair(pair).ContractType(contractType).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UfuturesGetFuturesDataBasis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetFuturesDataBasis`: []UfuturesGetFuturesDataBasisRespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UfuturesGetFuturesDataBasis`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetFuturesDataBasisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pair** | **string** | BTCUSDT | [default to &quot;&quot;]
 **contractType** | **string** | CURRENT_QUARTER, NEXT_QUARTER, PERPETUAL | [default to &quot;&quot;]
 **period** | **string** | &amp;#34;5m&amp;#34;,&amp;#34;15m&amp;#34;,&amp;#34;30m&amp;#34;,&amp;#34;1h&amp;#34;,&amp;#34;2h&amp;#34;,&amp;#34;4h&amp;#34;,&amp;#34;6h&amp;#34;,&amp;#34;12h&amp;#34;,&amp;#34;1d&amp;#34; | [default to &quot;&quot;]
 **limit** | **int64** | Default 30,Max 500 | [default to 30]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**[]UfuturesGetFuturesDataBasisRespItem**](UfuturesGetFuturesDataBasisRespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetFuturesDataDeliveryPrice

> []UfuturesGetFuturesDataDeliveryPriceRespItem UfuturesGetFuturesDataDeliveryPrice(ctx).Pair(pair).Execute()

Quarterly Contract Settlement Price



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {
	pair := "pair_example" // string | e.g BTCUSDT (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UfuturesGetFuturesDataDeliveryPrice(context.Background()).Pair(pair).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UfuturesGetFuturesDataDeliveryPrice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetFuturesDataDeliveryPrice`: []UfuturesGetFuturesDataDeliveryPriceRespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UfuturesGetFuturesDataDeliveryPrice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetFuturesDataDeliveryPriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pair** | **string** | e.g BTCUSDT | [default to &quot;&quot;]

### Return type

[**[]UfuturesGetFuturesDataDeliveryPriceRespItem**](UfuturesGetFuturesDataDeliveryPriceRespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetFuturesDataGlobalLongShortAccountRatio

> []UfuturesGetFuturesDataGlobalLongShortAccountRatioRespItem UfuturesGetFuturesDataGlobalLongShortAccountRatio(ctx).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

Long/Short Ratio



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	period := "period_example" // string | &#34;5m&#34;,&#34;15m&#34;,&#34;30m&#34;,&#34;1h&#34;,&#34;2h&#34;,&#34;4h&#34;,&#34;6h&#34;,&#34;12h&#34;,&#34;1d&#34; (default to "")
	limit := int64(789) // int64 | default 30, max 500 (optional) (default to 30)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UfuturesGetFuturesDataGlobalLongShortAccountRatio(context.Background()).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UfuturesGetFuturesDataGlobalLongShortAccountRatio``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetFuturesDataGlobalLongShortAccountRatio`: []UfuturesGetFuturesDataGlobalLongShortAccountRatioRespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UfuturesGetFuturesDataGlobalLongShortAccountRatio`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetFuturesDataGlobalLongShortAccountRatioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **period** | **string** | &amp;#34;5m&amp;#34;,&amp;#34;15m&amp;#34;,&amp;#34;30m&amp;#34;,&amp;#34;1h&amp;#34;,&amp;#34;2h&amp;#34;,&amp;#34;4h&amp;#34;,&amp;#34;6h&amp;#34;,&amp;#34;12h&amp;#34;,&amp;#34;1d&amp;#34; | [default to &quot;&quot;]
 **limit** | **int64** | default 30, max 500 | [default to 30]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**[]UfuturesGetFuturesDataGlobalLongShortAccountRatioRespItem**](UfuturesGetFuturesDataGlobalLongShortAccountRatioRespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetFuturesDataOpenInterestHist

> []UfuturesGetFuturesDataOpenInterestHistRespItem UfuturesGetFuturesDataOpenInterestHist(ctx).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

Open Interest Statistics



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	period := "period_example" // string | &#34;5m&#34;,&#34;15m&#34;,&#34;30m&#34;,&#34;1h&#34;,&#34;2h&#34;,&#34;4h&#34;,&#34;6h&#34;,&#34;12h&#34;,&#34;1d&#34; (default to "")
	limit := int64(789) // int64 | default 30, max 500 (optional) (default to 30)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UfuturesGetFuturesDataOpenInterestHist(context.Background()).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UfuturesGetFuturesDataOpenInterestHist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetFuturesDataOpenInterestHist`: []UfuturesGetFuturesDataOpenInterestHistRespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UfuturesGetFuturesDataOpenInterestHist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetFuturesDataOpenInterestHistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **period** | **string** | &amp;#34;5m&amp;#34;,&amp;#34;15m&amp;#34;,&amp;#34;30m&amp;#34;,&amp;#34;1h&amp;#34;,&amp;#34;2h&amp;#34;,&amp;#34;4h&amp;#34;,&amp;#34;6h&amp;#34;,&amp;#34;12h&amp;#34;,&amp;#34;1d&amp;#34; | [default to &quot;&quot;]
 **limit** | **int64** | default 30, max 500 | [default to 30]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**[]UfuturesGetFuturesDataOpenInterestHistRespItem**](UfuturesGetFuturesDataOpenInterestHistRespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetFuturesDataTakerlongshortRatio

> []UfuturesGetFuturesDataTakerlongshortRatioRespItem UfuturesGetFuturesDataTakerlongshortRatio(ctx).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

Taker Buy/Sell Volume



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	period := "period_example" // string | &#34;5m&#34;,&#34;15m&#34;,&#34;30m&#34;,&#34;1h&#34;,&#34;2h&#34;,&#34;4h&#34;,&#34;6h&#34;,&#34;12h&#34;,&#34;1d&#34; (default to "")
	limit := int64(789) // int64 | default 30, max 500 (optional) (default to 30)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UfuturesGetFuturesDataTakerlongshortRatio(context.Background()).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UfuturesGetFuturesDataTakerlongshortRatio``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetFuturesDataTakerlongshortRatio`: []UfuturesGetFuturesDataTakerlongshortRatioRespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UfuturesGetFuturesDataTakerlongshortRatio`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetFuturesDataTakerlongshortRatioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **period** | **string** | &amp;#34;5m&amp;#34;,&amp;#34;15m&amp;#34;,&amp;#34;30m&amp;#34;,&amp;#34;1h&amp;#34;,&amp;#34;2h&amp;#34;,&amp;#34;4h&amp;#34;,&amp;#34;6h&amp;#34;,&amp;#34;12h&amp;#34;,&amp;#34;1d&amp;#34; | [default to &quot;&quot;]
 **limit** | **int64** | default 30, max 500 | [default to 30]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**[]UfuturesGetFuturesDataTakerlongshortRatioRespItem**](UfuturesGetFuturesDataTakerlongshortRatioRespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetFuturesDataTopLongShortAccountRatio

> []UfuturesGetFuturesDataTopLongShortAccountRatioRespItem UfuturesGetFuturesDataTopLongShortAccountRatio(ctx).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

Top Trader Long/Short Ratio (Accounts)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	period := "period_example" // string | &#34;5m&#34;,&#34;15m&#34;,&#34;30m&#34;,&#34;1h&#34;,&#34;2h&#34;,&#34;4h&#34;,&#34;6h&#34;,&#34;12h&#34;,&#34;1d&#34; (default to "")
	limit := int64(789) // int64 | default 30, max 500 (optional) (default to 30)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UfuturesGetFuturesDataTopLongShortAccountRatio(context.Background()).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UfuturesGetFuturesDataTopLongShortAccountRatio``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetFuturesDataTopLongShortAccountRatio`: []UfuturesGetFuturesDataTopLongShortAccountRatioRespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UfuturesGetFuturesDataTopLongShortAccountRatio`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetFuturesDataTopLongShortAccountRatioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **period** | **string** | &amp;#34;5m&amp;#34;,&amp;#34;15m&amp;#34;,&amp;#34;30m&amp;#34;,&amp;#34;1h&amp;#34;,&amp;#34;2h&amp;#34;,&amp;#34;4h&amp;#34;,&amp;#34;6h&amp;#34;,&amp;#34;12h&amp;#34;,&amp;#34;1d&amp;#34; | [default to &quot;&quot;]
 **limit** | **int64** | default 30, max 500 | [default to 30]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**[]UfuturesGetFuturesDataTopLongShortAccountRatioRespItem**](UfuturesGetFuturesDataTopLongShortAccountRatioRespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetFuturesDataTopLongShortPositionRatio

> []UfuturesGetFuturesDataTopLongShortPositionRatioRespItem UfuturesGetFuturesDataTopLongShortPositionRatio(ctx).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

Top Trader Long/Short Ratio (Positions)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	period := "period_example" // string | &#34;5m&#34;,&#34;15m&#34;,&#34;30m&#34;,&#34;1h&#34;,&#34;2h&#34;,&#34;4h&#34;,&#34;6h&#34;,&#34;12h&#34;,&#34;1d&#34; (default to "")
	limit := int64(789) // int64 | default 30, max 500 (optional) (default to 30)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UfuturesGetFuturesDataTopLongShortPositionRatio(context.Background()).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UfuturesGetFuturesDataTopLongShortPositionRatio``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetFuturesDataTopLongShortPositionRatio`: []UfuturesGetFuturesDataTopLongShortPositionRatioRespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UfuturesGetFuturesDataTopLongShortPositionRatio`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetFuturesDataTopLongShortPositionRatioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **period** | **string** | &amp;#34;5m&amp;#34;,&amp;#34;15m&amp;#34;,&amp;#34;30m&amp;#34;,&amp;#34;1h&amp;#34;,&amp;#34;2h&amp;#34;,&amp;#34;4h&amp;#34;,&amp;#34;6h&amp;#34;,&amp;#34;12h&amp;#34;,&amp;#34;1d&amp;#34; | [default to &quot;&quot;]
 **limit** | **int64** | default 30, max 500 | [default to 30]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**[]UfuturesGetFuturesDataTopLongShortPositionRatioRespItem**](UfuturesGetFuturesDataTopLongShortPositionRatioRespItem.md)

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
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	limit := int32(56) // int32 | Default 100; max 500. (optional) (default to 100)
	fromId := int64(789) // int64 | TradeId to fetch from. Default gets most recent trades. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UfuturesGetHistoricalTradesV1(context.Background()).Symbol(symbol).Limit(limit).FromId(fromId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UfuturesGetHistoricalTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetHistoricalTradesV1`: []UfuturesGetHistoricalTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UfuturesGetHistoricalTradesV1`: %v\n", resp)
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
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {
	symbol := "symbol_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UfuturesGetIndexInfoV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UfuturesGetIndexInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetIndexInfoV1`: []UfuturesGetIndexInfoV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UfuturesGetIndexInfoV1`: %v\n", resp)
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

> [][]UfuturesGetContinuousKlinesV1200ResponseInnerInner UfuturesGetIndexPriceKlinesV1(ctx).Pair(pair).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Index Price Kline/Candlestick Data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {
	pair := "pair_example" // string |  (default to "")
	interval := "interval_example" // string |  (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 500; max 1500. (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UfuturesGetIndexPriceKlinesV1(context.Background()).Pair(pair).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UfuturesGetIndexPriceKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetIndexPriceKlinesV1`: [][]UfuturesGetContinuousKlinesV1200ResponseInnerInner
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UfuturesGetIndexPriceKlinesV1`: %v\n", resp)
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

[**[][]UfuturesGetContinuousKlinesV1200ResponseInnerInner**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetKlinesV1

> [][]UfuturesGetContinuousKlinesV1200ResponseInnerInner UfuturesGetKlinesV1(ctx).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Kline/Candlestick Data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	interval := "interval_example" // string |  (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 500; max 1500. (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UfuturesGetKlinesV1(context.Background()).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UfuturesGetKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetKlinesV1`: [][]UfuturesGetContinuousKlinesV1200ResponseInnerInner
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UfuturesGetKlinesV1`: %v\n", resp)
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

[**[][]UfuturesGetContinuousKlinesV1200ResponseInnerInner**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetMarkPriceKlinesV1

> [][]UfuturesGetContinuousKlinesV1200ResponseInnerInner UfuturesGetMarkPriceKlinesV1(ctx).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Mark Price Kline/Candlestick Data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	interval := "interval_example" // string |  (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 500; max 1500. (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UfuturesGetMarkPriceKlinesV1(context.Background()).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UfuturesGetMarkPriceKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetMarkPriceKlinesV1`: [][]UfuturesGetContinuousKlinesV1200ResponseInnerInner
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UfuturesGetMarkPriceKlinesV1`: %v\n", resp)
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

[**[][]UfuturesGetContinuousKlinesV1200ResponseInnerInner**](array.md)

### Authorization

No authorization required

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
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UfuturesGetOpenInterestV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UfuturesGetOpenInterestV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetOpenInterestV1`: UfuturesGetOpenInterestV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UfuturesGetOpenInterestV1`: %v\n", resp)
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
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UfuturesGetPingV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UfuturesGetPingV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetPingV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UfuturesGetPingV1`: %v\n", resp)
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


## UfuturesGetPremiumIndexKlinesV1

> [][]UfuturesGetContinuousKlinesV1200ResponseInnerInner UfuturesGetPremiumIndexKlinesV1(ctx).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Premium index Kline Data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	interval := "interval_example" // string |  (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 500; max 1500. (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UfuturesGetPremiumIndexKlinesV1(context.Background()).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UfuturesGetPremiumIndexKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetPremiumIndexKlinesV1`: [][]UfuturesGetContinuousKlinesV1200ResponseInnerInner
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UfuturesGetPremiumIndexKlinesV1`: %v\n", resp)
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

[**[][]UfuturesGetContinuousKlinesV1200ResponseInnerInner**](array.md)

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
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {
	symbol := "symbol_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UfuturesGetPremiumIndexV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UfuturesGetPremiumIndexV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetPremiumIndexV1`: UfuturesGetPremiumIndexV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UfuturesGetPremiumIndexV1`: %v\n", resp)
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
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {
	symbol := "symbol_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UfuturesGetTicker24hrV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UfuturesGetTicker24hrV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetTicker24hrV1`: UfuturesGetTicker24hrV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UfuturesGetTicker24hrV1`: %v\n", resp)
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
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {
	symbol := "symbol_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UfuturesGetTickerBookTickerV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UfuturesGetTickerBookTickerV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetTickerBookTickerV1`: UfuturesGetTickerBookTickerV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UfuturesGetTickerBookTickerV1`: %v\n", resp)
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
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {
	symbol := "symbol_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UfuturesGetTickerPriceV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UfuturesGetTickerPriceV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetTickerPriceV1`: UfuturesGetTickerPriceV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UfuturesGetTickerPriceV1`: %v\n", resp)
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


## UfuturesGetTickerPriceV2

> UfuturesGetTickerPriceV2Resp UfuturesGetTickerPriceV2(ctx).Symbol(symbol).Execute()

Symbol Price Ticker V2



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {
	symbol := "symbol_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UfuturesGetTickerPriceV2(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UfuturesGetTickerPriceV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetTickerPriceV2`: UfuturesGetTickerPriceV2Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UfuturesGetTickerPriceV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetTickerPriceV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**UfuturesGetTickerPriceV2Resp**](UfuturesGetTickerPriceV2Resp.md)

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
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UfuturesGetTimeV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UfuturesGetTimeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetTimeV1`: UfuturesGetTimeV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UfuturesGetTimeV1`: %v\n", resp)
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
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	limit := int32(56) // int32 | Default 500; max 1000. (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.UfuturesGetTradesV1(context.Background()).Symbol(symbol).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.UfuturesGetTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetTradesV1`: []UfuturesGetTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.UfuturesGetTradesV1`: %v\n", resp)
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

