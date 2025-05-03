# \MarketDataAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarginGetMarginAvailableInventoryV1**](MarketDataAPI.md#MarginGetMarginAvailableInventoryV1) | **Get** /sapi/v1/margin/available-inventory | Query Margin Available Inventory(USER_DATA)
[**SpotGetAggTradesV3**](MarketDataAPI.md#SpotGetAggTradesV3) | **Get** /api/v3/aggTrades | Compressed/Aggregate trades list
[**SpotGetKlinesV3**](MarketDataAPI.md#SpotGetKlinesV3) | **Get** /api/v3/klines | Kline/Candlestick data
[**SpotGetTicker24hrV3**](MarketDataAPI.md#SpotGetTicker24hrV3) | **Get** /api/v3/ticker/24hr | 24hr ticker price change statistics
[**SpotGetTickerBookTickerV3**](MarketDataAPI.md#SpotGetTickerBookTickerV3) | **Get** /api/v3/ticker/bookTicker | Symbol order book ticker
[**SpotGetTickerPriceV3**](MarketDataAPI.md#SpotGetTickerPriceV3) | **Get** /api/v3/ticker/price | Symbol price ticker
[**SpotGetTickerTradingDayV3**](MarketDataAPI.md#SpotGetTickerTradingDayV3) | **Get** /api/v3/ticker/tradingDay | Trading Day Ticker
[**SpotGetTickerV3**](MarketDataAPI.md#SpotGetTickerV3) | **Get** /api/v3/ticker | Rolling window price change statistics
[**SpotGetUiKlinesV3**](MarketDataAPI.md#SpotGetUiKlinesV3) | **Get** /api/v3/uiKlines | UIKlines



## MarginGetMarginAvailableInventoryV1

> MarginGetMarginAvailableInventoryV1Resp MarginGetMarginAvailableInventoryV1(ctx).Type_(type_).Execute()

Query Margin Available Inventory(USER_DATA)



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
	type_ := "type__example" // string | MARGIN,ISOLATED (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.MarginGetMarginAvailableInventoryV1(context.Background()).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.MarginGetMarginAvailableInventoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginAvailableInventoryV1`: MarginGetMarginAvailableInventoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.MarginGetMarginAvailableInventoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginAvailableInventoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | MARGIN,ISOLATED | [default to &quot;&quot;]

### Return type

[**MarginGetMarginAvailableInventoryV1Resp**](MarginGetMarginAvailableInventoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpotGetAggTradesV3

> []SpotGetAggTradesV3RespItem SpotGetAggTradesV3(ctx).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Compressed/Aggregate trades list



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
	fromId := int64(789) // int64 | ID to get aggregate trades from INCLUSIVE. (optional)
	startTime := int64(789) // int64 | Timestamp in ms to get aggregate trades from INCLUSIVE. (optional)
	endTime := int64(789) // int64 | Timestamp in ms to get aggregate trades until INCLUSIVE. (optional)
	limit := int32(56) // int32 | Default 500; max 1000. (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.SpotGetAggTradesV3(context.Background()).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.SpotGetAggTradesV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotGetAggTradesV3`: []SpotGetAggTradesV3RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.SpotGetAggTradesV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotGetAggTradesV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **fromId** | **int64** | ID to get aggregate trades from INCLUSIVE. | 
 **startTime** | **int64** | Timestamp in ms to get aggregate trades from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get aggregate trades until INCLUSIVE. | 
 **limit** | **int32** | Default 500; max 1000. | [default to 500]

### Return type

[**[]SpotGetAggTradesV3RespItem**](SpotGetAggTradesV3RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpotGetKlinesV3

> [][]SpotGetKlinesV3200ResponseInnerInner SpotGetKlinesV3(ctx).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).TimeZone(timeZone).Limit(limit).Execute()

Kline/Candlestick data



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
	interval := "interval_example" // string |  (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	timeZone := "timeZone_example" // string | Default: 0 (UTC) (optional) (default to "0")
	limit := int32(56) // int32 | Default 500; max 1000. (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.SpotGetKlinesV3(context.Background()).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).TimeZone(timeZone).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.SpotGetKlinesV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotGetKlinesV3`: [][]SpotGetKlinesV3200ResponseInnerInner
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.SpotGetKlinesV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotGetKlinesV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **interval** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **timeZone** | **string** | Default: 0 (UTC) | [default to &quot;0&quot;]
 **limit** | **int32** | Default 500; max 1000. | [default to 500]

### Return type

[**[][]SpotGetKlinesV3200ResponseInnerInner**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpotGetTicker24hrV3

> SpotGetTicker24hrV3Resp SpotGetTicker24hrV3(ctx).Symbol(symbol).Symbols(symbols).Type_(type_).Execute()

24hr ticker price change statistics



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
	symbol := "symbol_example" // string | Parameter symbol and symbols cannot be used in combination. <br/> If neither parameter is sent, tickers for all symbols will be returned in an array. <br/><br/>          Examples of accepted format for the symbols parameter:          [&#34;BTCUSDT&#34;,&#34;BNBUSDT&#34;] <br/>          or <br/>          %5B%22BTCUSDT%22,%22BNBUSDT%22%5D (optional) (default to "")
	symbols := "symbols_example" // string | Parameter symbol and symbols cannot be used in combination. <br/> If neither parameter is sent, tickers for all symbols will be returned in an array. <br/><br/>          Examples of accepted format for the symbols parameter:          [&#34;BTCUSDT&#34;,&#34;BNBUSDT&#34;] <br/>          or <br/>          %5B%22BTCUSDT%22,%22BNBUSDT%22%5D (optional) (default to "")
	type_ := "type__example" // string | Supported values: `FULL` or `MINI`. <br/>If none provided, the default is `FULL` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.SpotGetTicker24hrV3(context.Background()).Symbol(symbol).Symbols(symbols).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.SpotGetTicker24hrV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotGetTicker24hrV3`: SpotGetTicker24hrV3Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.SpotGetTicker24hrV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotGetTicker24hrV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Parameter symbol and symbols cannot be used in combination. &lt;br/&gt; If neither parameter is sent, tickers for all symbols will be returned in an array. &lt;br/&gt;&lt;br/&gt;          Examples of accepted format for the symbols parameter:          [&amp;#34;BTCUSDT&amp;#34;,&amp;#34;BNBUSDT&amp;#34;] &lt;br/&gt;          or &lt;br/&gt;          %5B%22BTCUSDT%22,%22BNBUSDT%22%5D | [default to &quot;&quot;]
 **symbols** | **string** | Parameter symbol and symbols cannot be used in combination. &lt;br/&gt; If neither parameter is sent, tickers for all symbols will be returned in an array. &lt;br/&gt;&lt;br/&gt;          Examples of accepted format for the symbols parameter:          [&amp;#34;BTCUSDT&amp;#34;,&amp;#34;BNBUSDT&amp;#34;] &lt;br/&gt;          or &lt;br/&gt;          %5B%22BTCUSDT%22,%22BNBUSDT%22%5D | [default to &quot;&quot;]
 **type_** | **string** | Supported values: &#x60;FULL&#x60; or &#x60;MINI&#x60;. &lt;br/&gt;If none provided, the default is &#x60;FULL&#x60; | [default to &quot;&quot;]

### Return type

[**SpotGetTicker24hrV3Resp**](SpotGetTicker24hrV3Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpotGetTickerBookTickerV3

> SpotGetTickerBookTickerV3Resp SpotGetTickerBookTickerV3(ctx).Symbol(symbol).Symbols(symbols).Execute()

Symbol order book ticker



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
	symbol := "symbol_example" // string | Parameter symbol and symbols cannot be used in combination. <br/> If neither parameter is sent, bookTickers for all symbols will be returned in an array.          <br/><br/>         Examples of accepted format for the symbols parameter:          [&#34;BTCUSDT&#34;,&#34;BNBUSDT&#34;] <br/>          or <br/>          %5B%22BTCUSDT%22,%22BNBUSDT%22%5D (optional) (default to "")
	symbols := "symbols_example" // string | Parameter symbol and symbols cannot be used in combination. <br/> If neither parameter is sent, bookTickers for all symbols will be returned in an array.          <br/><br/>         Examples of accepted format for the symbols parameter:          [&#34;BTCUSDT&#34;,&#34;BNBUSDT&#34;] <br/>          or <br/>          %5B%22BTCUSDT%22,%22BNBUSDT%22%5D (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.SpotGetTickerBookTickerV3(context.Background()).Symbol(symbol).Symbols(symbols).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.SpotGetTickerBookTickerV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotGetTickerBookTickerV3`: SpotGetTickerBookTickerV3Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.SpotGetTickerBookTickerV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotGetTickerBookTickerV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Parameter symbol and symbols cannot be used in combination. &lt;br/&gt; If neither parameter is sent, bookTickers for all symbols will be returned in an array.          &lt;br/&gt;&lt;br/&gt;         Examples of accepted format for the symbols parameter:          [&amp;#34;BTCUSDT&amp;#34;,&amp;#34;BNBUSDT&amp;#34;] &lt;br/&gt;          or &lt;br/&gt;          %5B%22BTCUSDT%22,%22BNBUSDT%22%5D | [default to &quot;&quot;]
 **symbols** | **string** | Parameter symbol and symbols cannot be used in combination. &lt;br/&gt; If neither parameter is sent, bookTickers for all symbols will be returned in an array.          &lt;br/&gt;&lt;br/&gt;         Examples of accepted format for the symbols parameter:          [&amp;#34;BTCUSDT&amp;#34;,&amp;#34;BNBUSDT&amp;#34;] &lt;br/&gt;          or &lt;br/&gt;          %5B%22BTCUSDT%22,%22BNBUSDT%22%5D | [default to &quot;&quot;]

### Return type

[**SpotGetTickerBookTickerV3Resp**](SpotGetTickerBookTickerV3Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpotGetTickerPriceV3

> SpotGetTickerPriceV3Resp SpotGetTickerPriceV3(ctx).Symbol(symbol).Symbols(symbols).Execute()

Symbol price ticker



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
	symbol := "symbol_example" // string | Parameter symbol and symbols cannot be used in combination. <br/> If neither parameter is sent, prices for all symbols will be returned in an array. <br/><br/>         Examples of accepted format for the symbols parameter:          [&#34;BTCUSDT&#34;,&#34;BNBUSDT&#34;] <br/>          or <br/>          %5B%22BTCUSDT%22,%22BNBUSDT%22%5D (optional) (default to "")
	symbols := "symbols_example" // string | Parameter symbol and symbols cannot be used in combination. <br/> If neither parameter is sent, prices for all symbols will be returned in an array. <br/><br/>         Examples of accepted format for the symbols parameter:          [&#34;BTCUSDT&#34;,&#34;BNBUSDT&#34;] <br/>          or <br/>          %5B%22BTCUSDT%22,%22BNBUSDT%22%5D (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.SpotGetTickerPriceV3(context.Background()).Symbol(symbol).Symbols(symbols).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.SpotGetTickerPriceV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotGetTickerPriceV3`: SpotGetTickerPriceV3Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.SpotGetTickerPriceV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotGetTickerPriceV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Parameter symbol and symbols cannot be used in combination. &lt;br/&gt; If neither parameter is sent, prices for all symbols will be returned in an array. &lt;br/&gt;&lt;br/&gt;         Examples of accepted format for the symbols parameter:          [&amp;#34;BTCUSDT&amp;#34;,&amp;#34;BNBUSDT&amp;#34;] &lt;br/&gt;          or &lt;br/&gt;          %5B%22BTCUSDT%22,%22BNBUSDT%22%5D | [default to &quot;&quot;]
 **symbols** | **string** | Parameter symbol and symbols cannot be used in combination. &lt;br/&gt; If neither parameter is sent, prices for all symbols will be returned in an array. &lt;br/&gt;&lt;br/&gt;         Examples of accepted format for the symbols parameter:          [&amp;#34;BTCUSDT&amp;#34;,&amp;#34;BNBUSDT&amp;#34;] &lt;br/&gt;          or &lt;br/&gt;          %5B%22BTCUSDT%22,%22BNBUSDT%22%5D | [default to &quot;&quot;]

### Return type

[**SpotGetTickerPriceV3Resp**](SpotGetTickerPriceV3Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpotGetTickerTradingDayV3

> SpotGetTickerTradingDayV3Resp SpotGetTickerTradingDayV3(ctx).Symbol(symbol).Symbols(symbols).TimeZone(timeZone).Type_(type_).Execute()

Trading Day Ticker



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
	symbol := "symbol_example" // string | Either `symbol` or `symbols` must be provided <br/><br/> Examples of accepted format for the `symbols` parameter: <br/> [&#34;BTCUSDT&#34;,&#34;BNBUSDT&#34;] <br/>or <br/>%5B%22BTCUSDT%22,%22BNBUSDT%22%5D <br/><br/> The maximum number of `symbols` allowed in a request is 100. (default to "")
	symbols := "symbols_example" // string | Either `symbol` or `symbols` must be provided <br/><br/> Examples of accepted format for the `symbols` parameter: <br/> [&#34;BTCUSDT&#34;,&#34;BNBUSDT&#34;] <br/>or <br/>%5B%22BTCUSDT%22,%22BNBUSDT%22%5D <br/><br/> The maximum number of `symbols` allowed in a request is 100. (default to "")
	timeZone := "timeZone_example" // string | Default: 0 (UTC) (optional) (default to "0")
	type_ := "type__example" // string | Supported values: `FULL` or `MINI`. <br/>If none provided, the default is `FULL` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.SpotGetTickerTradingDayV3(context.Background()).Symbol(symbol).Symbols(symbols).TimeZone(timeZone).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.SpotGetTickerTradingDayV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotGetTickerTradingDayV3`: SpotGetTickerTradingDayV3Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.SpotGetTickerTradingDayV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotGetTickerTradingDayV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Either &#x60;symbol&#x60; or &#x60;symbols&#x60; must be provided &lt;br/&gt;&lt;br/&gt; Examples of accepted format for the &#x60;symbols&#x60; parameter: &lt;br/&gt; [&amp;#34;BTCUSDT&amp;#34;,&amp;#34;BNBUSDT&amp;#34;] &lt;br/&gt;or &lt;br/&gt;%5B%22BTCUSDT%22,%22BNBUSDT%22%5D &lt;br/&gt;&lt;br/&gt; The maximum number of &#x60;symbols&#x60; allowed in a request is 100. | [default to &quot;&quot;]
 **symbols** | **string** | Either &#x60;symbol&#x60; or &#x60;symbols&#x60; must be provided &lt;br/&gt;&lt;br/&gt; Examples of accepted format for the &#x60;symbols&#x60; parameter: &lt;br/&gt; [&amp;#34;BTCUSDT&amp;#34;,&amp;#34;BNBUSDT&amp;#34;] &lt;br/&gt;or &lt;br/&gt;%5B%22BTCUSDT%22,%22BNBUSDT%22%5D &lt;br/&gt;&lt;br/&gt; The maximum number of &#x60;symbols&#x60; allowed in a request is 100. | [default to &quot;&quot;]
 **timeZone** | **string** | Default: 0 (UTC) | [default to &quot;0&quot;]
 **type_** | **string** | Supported values: &#x60;FULL&#x60; or &#x60;MINI&#x60;. &lt;br/&gt;If none provided, the default is &#x60;FULL&#x60; | [default to &quot;&quot;]

### Return type

[**SpotGetTickerTradingDayV3Resp**](SpotGetTickerTradingDayV3Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpotGetTickerV3

> SpotGetTickerV3Resp SpotGetTickerV3(ctx).Symbol(symbol).Symbols(symbols).WindowSize(windowSize).Type_(type_).Execute()

Rolling window price change statistics



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
	symbol := "symbol_example" // string | Either `symbol` or `symbols` must be provided <br/><br/> Examples of accepted format for the `symbols` parameter: <br/> [&#34;BTCUSDT&#34;,&#34;BNBUSDT&#34;] <br/>or <br/>%5B%22BTCUSDT%22,%22BNBUSDT%22%5D <br/><br/> The maximum number of `symbols` allowed in a request is 100. (default to "")
	symbols := "symbols_example" // string | Either `symbol` or `symbols` must be provided <br/><br/> Examples of accepted format for the `symbols` parameter: <br/> [&#34;BTCUSDT&#34;,&#34;BNBUSDT&#34;] <br/>or <br/>%5B%22BTCUSDT%22,%22BNBUSDT%22%5D <br/><br/> The maximum number of `symbols` allowed in a request is 100. (default to "")
	windowSize := "windowSize_example" // string | Defaults to `1d` if no parameter provided <br/> Supported `windowSize` values: <br/> `1m`,`2m`....`59m` for minutes <br/> `1h`, `2h`....`23h` - for hours <br/> `1d`...`7d` - for days <br/><br/> Units cannot be combined (e.g. `1d2h` is not allowed) (optional) (default to "")
	type_ := "type__example" // string | Supported values: `FULL` or `MINI`. <br/>If none provided, the default is `FULL` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.SpotGetTickerV3(context.Background()).Symbol(symbol).Symbols(symbols).WindowSize(windowSize).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.SpotGetTickerV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotGetTickerV3`: SpotGetTickerV3Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.SpotGetTickerV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotGetTickerV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Either &#x60;symbol&#x60; or &#x60;symbols&#x60; must be provided &lt;br/&gt;&lt;br/&gt; Examples of accepted format for the &#x60;symbols&#x60; parameter: &lt;br/&gt; [&amp;#34;BTCUSDT&amp;#34;,&amp;#34;BNBUSDT&amp;#34;] &lt;br/&gt;or &lt;br/&gt;%5B%22BTCUSDT%22,%22BNBUSDT%22%5D &lt;br/&gt;&lt;br/&gt; The maximum number of &#x60;symbols&#x60; allowed in a request is 100. | [default to &quot;&quot;]
 **symbols** | **string** | Either &#x60;symbol&#x60; or &#x60;symbols&#x60; must be provided &lt;br/&gt;&lt;br/&gt; Examples of accepted format for the &#x60;symbols&#x60; parameter: &lt;br/&gt; [&amp;#34;BTCUSDT&amp;#34;,&amp;#34;BNBUSDT&amp;#34;] &lt;br/&gt;or &lt;br/&gt;%5B%22BTCUSDT%22,%22BNBUSDT%22%5D &lt;br/&gt;&lt;br/&gt; The maximum number of &#x60;symbols&#x60; allowed in a request is 100. | [default to &quot;&quot;]
 **windowSize** | **string** | Defaults to &#x60;1d&#x60; if no parameter provided &lt;br/&gt; Supported &#x60;windowSize&#x60; values: &lt;br/&gt; &#x60;1m&#x60;,&#x60;2m&#x60;....&#x60;59m&#x60; for minutes &lt;br/&gt; &#x60;1h&#x60;, &#x60;2h&#x60;....&#x60;23h&#x60; - for hours &lt;br/&gt; &#x60;1d&#x60;...&#x60;7d&#x60; - for days &lt;br/&gt;&lt;br/&gt; Units cannot be combined (e.g. &#x60;1d2h&#x60; is not allowed) | [default to &quot;&quot;]
 **type_** | **string** | Supported values: &#x60;FULL&#x60; or &#x60;MINI&#x60;. &lt;br/&gt;If none provided, the default is &#x60;FULL&#x60; | [default to &quot;&quot;]

### Return type

[**SpotGetTickerV3Resp**](SpotGetTickerV3Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpotGetUiKlinesV3

> [][]SpotGetKlinesV3200ResponseInnerInner SpotGetUiKlinesV3(ctx).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).TimeZone(timeZone).Limit(limit).Execute()

UIKlines



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
	interval := "interval_example" // string | See <a href=\"/docs/binance-spot-api-docs/rest-api/market-data-endpoints#kline-intervals\">`klines`</a> (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	timeZone := "timeZone_example" // string | Default: 0 (UTC) (optional) (default to "0")
	limit := int32(56) // int32 | Default 500; max 1000. (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.SpotGetUiKlinesV3(context.Background()).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).TimeZone(timeZone).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.SpotGetUiKlinesV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotGetUiKlinesV3`: [][]SpotGetKlinesV3200ResponseInnerInner
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.SpotGetUiKlinesV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotGetUiKlinesV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **interval** | **string** | See &lt;a href&#x3D;\&quot;/docs/binance-spot-api-docs/rest-api/market-data-endpoints#kline-intervals\&quot;&gt;&#x60;klines&#x60;&lt;/a&gt; | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **timeZone** | **string** | Default: 0 (UTC) | [default to &quot;0&quot;]
 **limit** | **int32** | Default 500; max 1000. | [default to 500]

### Return type

[**[][]SpotGetKlinesV3200ResponseInnerInner**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

