# \MarketDataAPI

All URIs are relative to *https://eapi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OptionsGetBlockTradesV1**](MarketDataAPI.md#OptionsGetBlockTradesV1) | **Get** /eapi/v1/blockTrades | Recent Block Trades List
[**OptionsGetDepthV1**](MarketDataAPI.md#OptionsGetDepthV1) | **Get** /eapi/v1/depth | Order Book
[**OptionsGetExchangeInfoV1**](MarketDataAPI.md#OptionsGetExchangeInfoV1) | **Get** /eapi/v1/exchangeInfo | Exchange Information
[**OptionsGetExerciseHistoryV1**](MarketDataAPI.md#OptionsGetExerciseHistoryV1) | **Get** /eapi/v1/exerciseHistory | Historical Exercise Records
[**OptionsGetHistoricalTradesV1**](MarketDataAPI.md#OptionsGetHistoricalTradesV1) | **Get** /eapi/v1/historicalTrades | Old Trades Lookup (MARKET_DATA)
[**OptionsGetIndexV1**](MarketDataAPI.md#OptionsGetIndexV1) | **Get** /eapi/v1/index | Symbol Price Ticker
[**OptionsGetKlinesV1**](MarketDataAPI.md#OptionsGetKlinesV1) | **Get** /eapi/v1/klines | Kline/Candlestick Data
[**OptionsGetMarkV1**](MarketDataAPI.md#OptionsGetMarkV1) | **Get** /eapi/v1/mark | Option Mark Price
[**OptionsGetOpenInterestV1**](MarketDataAPI.md#OptionsGetOpenInterestV1) | **Get** /eapi/v1/openInterest | Open Interest
[**OptionsGetPingV1**](MarketDataAPI.md#OptionsGetPingV1) | **Get** /eapi/v1/ping | Test Connectivity
[**OptionsGetTickerV1**](MarketDataAPI.md#OptionsGetTickerV1) | **Get** /eapi/v1/ticker | 24hr Ticker Price Change Statistics
[**OptionsGetTimeV1**](MarketDataAPI.md#OptionsGetTimeV1) | **Get** /eapi/v1/time | Check Server Time
[**OptionsGetTradesV1**](MarketDataAPI.md#OptionsGetTradesV1) | **Get** /eapi/v1/trades | Recent Trades List



## OptionsGetBlockTradesV1

> []OptionsGetBlockTradesV1RespItem OptionsGetBlockTradesV1(ctx).Symbol(symbol).Limit(limit).Execute()

Recent Block Trades List



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/options"
)

func main() {
	symbol := "symbol_example" // string | Option trading pair, e.g. BTC-200730-9000-C (optional) (default to "")
	limit := int32(56) // int32 | Number of records; Default: 100 and Max: 500 (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.OptionsGetBlockTradesV1(context.Background()).Symbol(symbol).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.OptionsGetBlockTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetBlockTradesV1`: []OptionsGetBlockTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.OptionsGetBlockTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetBlockTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair, e.g. BTC-200730-9000-C | [default to &quot;&quot;]
 **limit** | **int32** | Number of records; Default: 100 and Max: 500 | [default to 100]

### Return type

[**[]OptionsGetBlockTradesV1RespItem**](OptionsGetBlockTradesV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetDepthV1

> OptionsGetDepthV1Resp OptionsGetDepthV1(ctx).Symbol(symbol).Limit(limit).Execute()

Order Book



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/options"
)

func main() {
	symbol := "symbol_example" // string | Option trading pair, e.g BTC-200730-9000-C (default to "")
	limit := int32(56) // int32 | Default:100 Max:1000.Optional value:[10, 20, 50, 100, 500, 1000] (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.OptionsGetDepthV1(context.Background()).Symbol(symbol).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.OptionsGetDepthV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetDepthV1`: OptionsGetDepthV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.OptionsGetDepthV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetDepthV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | [default to &quot;&quot;]
 **limit** | **int32** | Default:100 Max:1000.Optional value:[10, 20, 50, 100, 500, 1000] | 

### Return type

[**OptionsGetDepthV1Resp**](OptionsGetDepthV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetExchangeInfoV1

> OptionsGetExchangeInfoV1Resp OptionsGetExchangeInfoV1(ctx).Execute()

Exchange Information



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/options"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.OptionsGetExchangeInfoV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.OptionsGetExchangeInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetExchangeInfoV1`: OptionsGetExchangeInfoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.OptionsGetExchangeInfoV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetExchangeInfoV1Request struct via the builder pattern


### Return type

[**OptionsGetExchangeInfoV1Resp**](OptionsGetExchangeInfoV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetExerciseHistoryV1

> []OptionsGetExerciseHistoryV1RespItem OptionsGetExerciseHistoryV1(ctx).Underlying(underlying).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Historical Exercise Records



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/options"
)

func main() {
	underlying := "underlying_example" // string | Underlying index like BTCUSDT (optional) (default to "")
	startTime := int64(789) // int64 | Start Time (optional)
	endTime := int64(789) // int64 | End Time (optional)
	limit := int32(56) // int32 | Number of records Default:100 Max:100 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.OptionsGetExerciseHistoryV1(context.Background()).Underlying(underlying).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.OptionsGetExerciseHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetExerciseHistoryV1`: []OptionsGetExerciseHistoryV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.OptionsGetExerciseHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetExerciseHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **underlying** | **string** | Underlying index like BTCUSDT | [default to &quot;&quot;]
 **startTime** | **int64** | Start Time | 
 **endTime** | **int64** | End Time | 
 **limit** | **int32** | Number of records Default:100 Max:100 | 

### Return type

[**[]OptionsGetExerciseHistoryV1RespItem**](OptionsGetExerciseHistoryV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetHistoricalTradesV1

> []OptionsGetHistoricalTradesV1RespItem OptionsGetHistoricalTradesV1(ctx).Symbol(symbol).FromId(fromId).Limit(limit).Execute()

Old Trades Lookup (MARKET_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/options"
)

func main() {
	symbol := "symbol_example" // string | Option trading pair, e.g BTC-200730-9000-C (default to "")
	fromId := int64(789) // int64 | The UniqueId ID from which to return. The latest deal record is returned by default (optional)
	limit := int32(56) // int32 | Number of records Default:100 Max:500 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.OptionsGetHistoricalTradesV1(context.Background()).Symbol(symbol).FromId(fromId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.OptionsGetHistoricalTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetHistoricalTradesV1`: []OptionsGetHistoricalTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.OptionsGetHistoricalTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetHistoricalTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | [default to &quot;&quot;]
 **fromId** | **int64** | The UniqueId ID from which to return. The latest deal record is returned by default | 
 **limit** | **int32** | Number of records Default:100 Max:500 | 

### Return type

[**[]OptionsGetHistoricalTradesV1RespItem**](OptionsGetHistoricalTradesV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetIndexV1

> OptionsGetIndexV1Resp OptionsGetIndexV1(ctx).Underlying(underlying).Execute()

Symbol Price Ticker



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/options"
)

func main() {
	underlying := "underlying_example" // string | Spot pair（Option contract underlying asset, e.g BTCUSDT) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.OptionsGetIndexV1(context.Background()).Underlying(underlying).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.OptionsGetIndexV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetIndexV1`: OptionsGetIndexV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.OptionsGetIndexV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetIndexV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **underlying** | **string** | Spot pair（Option contract underlying asset, e.g BTCUSDT) | [default to &quot;&quot;]

### Return type

[**OptionsGetIndexV1Resp**](OptionsGetIndexV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetKlinesV1

> []OptionsGetKlinesV1RespItem OptionsGetKlinesV1(ctx).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Kline/Candlestick Data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/options"
)

func main() {
	symbol := "symbol_example" // string | Option trading pair, e.g BTC-200730-9000-C (default to "")
	interval := "interval_example" // string | Time interval (default to "")
	startTime := int64(789) // int64 | Start Time  1592317127349 (optional)
	endTime := int64(789) // int64 | End Time (optional)
	limit := int32(56) // int32 | Number of records Default:500 Max:1500 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.OptionsGetKlinesV1(context.Background()).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.OptionsGetKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetKlinesV1`: []OptionsGetKlinesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.OptionsGetKlinesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetKlinesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | [default to &quot;&quot;]
 **interval** | **string** | Time interval | [default to &quot;&quot;]
 **startTime** | **int64** | Start Time  1592317127349 | 
 **endTime** | **int64** | End Time | 
 **limit** | **int32** | Number of records Default:500 Max:1500 | 

### Return type

[**[]OptionsGetKlinesV1RespItem**](OptionsGetKlinesV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetMarkV1

> []OptionsGetMarkV1RespItem OptionsGetMarkV1(ctx).Symbol(symbol).Execute()

Option Mark Price



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/options"
)

func main() {
	symbol := "symbol_example" // string | Option trading pair, e.g BTC-200730-9000-C (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.OptionsGetMarkV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.OptionsGetMarkV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetMarkV1`: []OptionsGetMarkV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.OptionsGetMarkV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetMarkV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | [default to &quot;&quot;]

### Return type

[**[]OptionsGetMarkV1RespItem**](OptionsGetMarkV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetOpenInterestV1

> []OptionsGetOpenInterestV1RespItem OptionsGetOpenInterestV1(ctx).UnderlyingAsset(underlyingAsset).Expiration(expiration).Execute()

Open Interest



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/options"
)

func main() {
	underlyingAsset := "underlyingAsset_example" // string | underlying asset, e.g ETH/BTC (default to "")
	expiration := "expiration_example" // string | expiration date, e.g 221225 (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.OptionsGetOpenInterestV1(context.Background()).UnderlyingAsset(underlyingAsset).Expiration(expiration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.OptionsGetOpenInterestV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetOpenInterestV1`: []OptionsGetOpenInterestV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.OptionsGetOpenInterestV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetOpenInterestV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **underlyingAsset** | **string** | underlying asset, e.g ETH/BTC | [default to &quot;&quot;]
 **expiration** | **string** | expiration date, e.g 221225 | [default to &quot;&quot;]

### Return type

[**[]OptionsGetOpenInterestV1RespItem**](OptionsGetOpenInterestV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetPingV1

> map[string]interface{} OptionsGetPingV1(ctx).Execute()

Test Connectivity



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/options"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.OptionsGetPingV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.OptionsGetPingV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetPingV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.OptionsGetPingV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetPingV1Request struct via the builder pattern


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


## OptionsGetTickerV1

> []OptionsGetTickerV1RespItem OptionsGetTickerV1(ctx).Symbol(symbol).Execute()

24hr Ticker Price Change Statistics



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/options"
)

func main() {
	symbol := "symbol_example" // string | Option trading pair, e.g BTC-200730-9000-C (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.OptionsGetTickerV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.OptionsGetTickerV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetTickerV1`: []OptionsGetTickerV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.OptionsGetTickerV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetTickerV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | [default to &quot;&quot;]

### Return type

[**[]OptionsGetTickerV1RespItem**](OptionsGetTickerV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetTimeV1

> OptionsGetTimeV1Resp OptionsGetTimeV1(ctx).Execute()

Check Server Time



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/options"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.OptionsGetTimeV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.OptionsGetTimeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetTimeV1`: OptionsGetTimeV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.OptionsGetTimeV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetTimeV1Request struct via the builder pattern


### Return type

[**OptionsGetTimeV1Resp**](OptionsGetTimeV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetTradesV1

> []OptionsGetTradesV1RespItem OptionsGetTradesV1(ctx).Symbol(symbol).Limit(limit).Execute()

Recent Trades List



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/options"
)

func main() {
	symbol := "symbol_example" // string | Option trading pair, e.g BTC-200730-9000-C (default to "")
	limit := int32(56) // int32 | Number of records Default:100 Max:500 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.OptionsGetTradesV1(context.Background()).Symbol(symbol).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.OptionsGetTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetTradesV1`: []OptionsGetTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.OptionsGetTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | [default to &quot;&quot;]
 **limit** | **int32** | Number of records Default:100 Max:500 | 

### Return type

[**[]OptionsGetTradesV1RespItem**](OptionsGetTradesV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

