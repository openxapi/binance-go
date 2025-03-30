# \V1API

All URIs are relative to *https://eapi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OptionsCreateBatchOrdersV1**](V1API.md#OptionsCreateBatchOrdersV1) | **Post** /eapi/v1/batchOrders | Place Multiple Orders(TRADE)
[**OptionsCreateBlockOrderCreateV1**](V1API.md#OptionsCreateBlockOrderCreateV1) | **Post** /eapi/v1/block/order/create | New Block Trade Order (TRADE)
[**OptionsCreateBlockOrderExecuteV1**](V1API.md#OptionsCreateBlockOrderExecuteV1) | **Post** /eapi/v1/block/order/execute | Accept Block Trade Order (TRADE)
[**OptionsCreateCountdownCancelAllHeartBeatV1**](V1API.md#OptionsCreateCountdownCancelAllHeartBeatV1) | **Post** /eapi/v1/countdownCancelAllHeartBeat | Auto-Cancel All Open Orders (Kill-Switch) Heartbeat (TRADE)
[**OptionsCreateCountdownCancelAllV1**](V1API.md#OptionsCreateCountdownCancelAllV1) | **Post** /eapi/v1/countdownCancelAll | Set Auto-Cancel All Open Orders (Kill-Switch) Config (TRADE)
[**OptionsCreateListenKeyV1**](V1API.md#OptionsCreateListenKeyV1) | **Post** /eapi/v1/listenKey | Start User Data Stream (USER_STREAM)
[**OptionsCreateMmpResetV1**](V1API.md#OptionsCreateMmpResetV1) | **Post** /eapi/v1/mmpReset | Reset Market Maker Protection Config (TRADE)
[**OptionsCreateMmpSetV1**](V1API.md#OptionsCreateMmpSetV1) | **Post** /eapi/v1/mmpSet | Set Market Maker Protection Config (TRADE)
[**OptionsCreateOrderV1**](V1API.md#OptionsCreateOrderV1) | **Post** /eapi/v1/order | New Order (TRADE)
[**OptionsDeleteAllOpenOrdersByUnderlyingV1**](V1API.md#OptionsDeleteAllOpenOrdersByUnderlyingV1) | **Delete** /eapi/v1/allOpenOrdersByUnderlying | Cancel All Option Orders By Underlying (TRADE)
[**OptionsDeleteAllOpenOrdersV1**](V1API.md#OptionsDeleteAllOpenOrdersV1) | **Delete** /eapi/v1/allOpenOrders | Cancel all Option orders on specific symbol (TRADE)
[**OptionsDeleteBatchOrdersV1**](V1API.md#OptionsDeleteBatchOrdersV1) | **Delete** /eapi/v1/batchOrders | Cancel Multiple Option Orders (TRADE)
[**OptionsDeleteBlockOrderCreateV1**](V1API.md#OptionsDeleteBlockOrderCreateV1) | **Delete** /eapi/v1/block/order/create | Cancel Block Trade Order (TRADE)
[**OptionsDeleteListenKeyV1**](V1API.md#OptionsDeleteListenKeyV1) | **Delete** /eapi/v1/listenKey | Close User Data Stream (USER_STREAM)
[**OptionsDeleteOrderV1**](V1API.md#OptionsDeleteOrderV1) | **Delete** /eapi/v1/order | Cancel Option Order (TRADE)
[**OptionsGetAccountV1**](V1API.md#OptionsGetAccountV1) | **Get** /eapi/v1/account | Option Account Information(TRADE)
[**OptionsGetBillV1**](V1API.md#OptionsGetBillV1) | **Get** /eapi/v1/bill | Account Funding Flow (USER_DATA)
[**OptionsGetBlockOrderExecuteV1**](V1API.md#OptionsGetBlockOrderExecuteV1) | **Get** /eapi/v1/block/order/execute | Query Block Trade Details (USER_DATA)
[**OptionsGetBlockOrderOrdersV1**](V1API.md#OptionsGetBlockOrderOrdersV1) | **Get** /eapi/v1/block/order/orders | Query Block Trade Order (TRADE)
[**OptionsGetBlockTradesV1**](V1API.md#OptionsGetBlockTradesV1) | **Get** /eapi/v1/blockTrades | Recent Block Trades List
[**OptionsGetBlockUserTradesV1**](V1API.md#OptionsGetBlockUserTradesV1) | **Get** /eapi/v1/block/user-trades | Account Block Trade List (USER_DATA)
[**OptionsGetCountdownCancelAllV1**](V1API.md#OptionsGetCountdownCancelAllV1) | **Get** /eapi/v1/countdownCancelAll | Get Auto-Cancel All Open Orders (Kill-Switch) Config (TRADE)
[**OptionsGetDepthV1**](V1API.md#OptionsGetDepthV1) | **Get** /eapi/v1/depth | Order Book
[**OptionsGetExchangeInfoV1**](V1API.md#OptionsGetExchangeInfoV1) | **Get** /eapi/v1/exchangeInfo | Exchange Information
[**OptionsGetExerciseHistoryV1**](V1API.md#OptionsGetExerciseHistoryV1) | **Get** /eapi/v1/exerciseHistory | Historical Exercise Records
[**OptionsGetExerciseRecordV1**](V1API.md#OptionsGetExerciseRecordV1) | **Get** /eapi/v1/exerciseRecord | User Exercise Record (USER_DATA)
[**OptionsGetHistoricalTradesV1**](V1API.md#OptionsGetHistoricalTradesV1) | **Get** /eapi/v1/historicalTrades | Old Trades Lookup (MARKET_DATA)
[**OptionsGetHistoryOrdersV1**](V1API.md#OptionsGetHistoryOrdersV1) | **Get** /eapi/v1/historyOrders | Query Option Order History (TRADE)
[**OptionsGetIncomeAsynIdV1**](V1API.md#OptionsGetIncomeAsynIdV1) | **Get** /eapi/v1/income/asyn/id | Get Option Transaction History Download Link by Id (USER_DATA)
[**OptionsGetIncomeAsynV1**](V1API.md#OptionsGetIncomeAsynV1) | **Get** /eapi/v1/income/asyn | Get Download Id For Option Transaction History (USER_DATA)
[**OptionsGetIndexV1**](V1API.md#OptionsGetIndexV1) | **Get** /eapi/v1/index | Symbol Price Ticker
[**OptionsGetKlinesV1**](V1API.md#OptionsGetKlinesV1) | **Get** /eapi/v1/klines | Kline/Candlestick Data
[**OptionsGetMarginAccountV1**](V1API.md#OptionsGetMarginAccountV1) | **Get** /eapi/v1/marginAccount | Option Margin Account Information (USER_DATA)
[**OptionsGetMarkV1**](V1API.md#OptionsGetMarkV1) | **Get** /eapi/v1/mark | Option Mark Price
[**OptionsGetMmpV1**](V1API.md#OptionsGetMmpV1) | **Get** /eapi/v1/mmp | Get Market Maker Protection Config (TRADE)
[**OptionsGetOpenInterestV1**](V1API.md#OptionsGetOpenInterestV1) | **Get** /eapi/v1/openInterest | Open Interest
[**OptionsGetOpenOrdersV1**](V1API.md#OptionsGetOpenOrdersV1) | **Get** /eapi/v1/openOrders | Query Current Open Option Orders (USER_DATA)
[**OptionsGetOrderV1**](V1API.md#OptionsGetOrderV1) | **Get** /eapi/v1/order | Query Single Order (TRADE)
[**OptionsGetPingV1**](V1API.md#OptionsGetPingV1) | **Get** /eapi/v1/ping | Test Connectivity
[**OptionsGetPositionV1**](V1API.md#OptionsGetPositionV1) | **Get** /eapi/v1/position | Option Position Information (USER_DATA)
[**OptionsGetTickerV1**](V1API.md#OptionsGetTickerV1) | **Get** /eapi/v1/ticker | 24hr Ticker Price Change Statistics
[**OptionsGetTimeV1**](V1API.md#OptionsGetTimeV1) | **Get** /eapi/v1/time | Check Server Time
[**OptionsGetTradesV1**](V1API.md#OptionsGetTradesV1) | **Get** /eapi/v1/trades | Recent Trades List
[**OptionsGetUserTradesV1**](V1API.md#OptionsGetUserTradesV1) | **Get** /eapi/v1/userTrades | Account Trade List (USER_DATA)
[**OptionsUpdateBlockOrderCreateV1**](V1API.md#OptionsUpdateBlockOrderCreateV1) | **Put** /eapi/v1/block/order/create | Extend Block Trade Order (TRADE)
[**OptionsUpdateListenKeyV1**](V1API.md#OptionsUpdateListenKeyV1) | **Put** /eapi/v1/listenKey | Keepalive User Data Stream (USER_STREAM)



## OptionsCreateBatchOrdersV1

> []OptionsCreateBatchOrdersV1RespInner OptionsCreateBatchOrdersV1(ctx).Orders(orders).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Place Multiple Orders(TRADE)



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
	orders := []openapiclient.OptionsCreateBatchOrdersV1ReqOrdersItem{*openapiclient.NewOptionsCreateBatchOrdersV1ReqOrdersItem("Quantity_example", "Side_example", "Symbol_example", "Type_example")} // []OptionsCreateBatchOrdersV1ReqOrdersItem | 
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.OptionsCreateBatchOrdersV1(context.Background()).Orders(orders).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsCreateBatchOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsCreateBatchOrdersV1`: []OptionsCreateBatchOrdersV1RespInner
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsCreateBatchOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsCreateBatchOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orders** | [**[]OptionsCreateBatchOrdersV1ReqOrdersItem**](OptionsCreateBatchOrdersV1ReqOrdersItem.md) |  | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]OptionsCreateBatchOrdersV1RespInner**](OptionsCreateBatchOrdersV1RespInner.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsCreateBlockOrderCreateV1

> OptionsCreateBlockOrderCreateV1Resp OptionsCreateBlockOrderCreateV1(ctx).Legs(legs).Liquidity(liquidity).Price(price).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

New Block Trade Order (TRADE)



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
	legs := []string{"Inner_example"} // []string | 
	liquidity := "liquidity_example" // string |  (default to "")
	price := "price_example" // string |  (default to "")
	quantity := "quantity_example" // string |  (default to "")
	side := "side_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int32(56) // int32 | 
	recvWindow := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.OptionsCreateBlockOrderCreateV1(context.Background()).Legs(legs).Liquidity(liquidity).Price(price).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsCreateBlockOrderCreateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsCreateBlockOrderCreateV1`: OptionsCreateBlockOrderCreateV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsCreateBlockOrderCreateV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsCreateBlockOrderCreateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **legs** | **[]string** |  | 
 **liquidity** | **string** |  | [default to &quot;&quot;]
 **price** | **string** |  | [default to &quot;&quot;]
 **quantity** | **string** |  | [default to &quot;&quot;]
 **side** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int32** |  | 
 **recvWindow** | **int32** |  | 

### Return type

[**OptionsCreateBlockOrderCreateV1Resp**](OptionsCreateBlockOrderCreateV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsCreateBlockOrderExecuteV1

> OptionsCreateBlockOrderExecuteV1Resp OptionsCreateBlockOrderExecuteV1(ctx).BlockOrderMatchingKey(blockOrderMatchingKey).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Accept Block Trade Order (TRADE)



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
	blockOrderMatchingKey := "blockOrderMatchingKey_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.OptionsCreateBlockOrderExecuteV1(context.Background()).BlockOrderMatchingKey(blockOrderMatchingKey).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsCreateBlockOrderExecuteV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsCreateBlockOrderExecuteV1`: OptionsCreateBlockOrderExecuteV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsCreateBlockOrderExecuteV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsCreateBlockOrderExecuteV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockOrderMatchingKey** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**OptionsCreateBlockOrderExecuteV1Resp**](OptionsCreateBlockOrderExecuteV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsCreateCountdownCancelAllHeartBeatV1

> OptionsCreateCountdownCancelAllHeartBeatV1Resp OptionsCreateCountdownCancelAllHeartBeatV1(ctx).Timestamp(timestamp).Underlyings(underlyings).RecvWindow(recvWindow).Execute()

Auto-Cancel All Open Orders (Kill-Switch) Heartbeat (TRADE)



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
	timestamp := int64(789) // int64 | 
	underlyings := "underlyings_example" // string |  (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.OptionsCreateCountdownCancelAllHeartBeatV1(context.Background()).Timestamp(timestamp).Underlyings(underlyings).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsCreateCountdownCancelAllHeartBeatV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsCreateCountdownCancelAllHeartBeatV1`: OptionsCreateCountdownCancelAllHeartBeatV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsCreateCountdownCancelAllHeartBeatV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsCreateCountdownCancelAllHeartBeatV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **underlyings** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**OptionsCreateCountdownCancelAllHeartBeatV1Resp**](OptionsCreateCountdownCancelAllHeartBeatV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsCreateCountdownCancelAllV1

> OptionsCreateCountdownCancelAllV1Resp OptionsCreateCountdownCancelAllV1(ctx).CountdownTime(countdownTime).Timestamp(timestamp).Underlying(underlying).RecvWindow(recvWindow).Execute()

Set Auto-Cancel All Open Orders (Kill-Switch) Config (TRADE)



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
	countdownTime := int64(789) // int64 | 
	timestamp := int64(789) // int64 | 
	underlying := "underlying_example" // string |  (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.OptionsCreateCountdownCancelAllV1(context.Background()).CountdownTime(countdownTime).Timestamp(timestamp).Underlying(underlying).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsCreateCountdownCancelAllV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsCreateCountdownCancelAllV1`: OptionsCreateCountdownCancelAllV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsCreateCountdownCancelAllV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsCreateCountdownCancelAllV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countdownTime** | **int64** |  | 
 **timestamp** | **int64** |  | 
 **underlying** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**OptionsCreateCountdownCancelAllV1Resp**](OptionsCreateCountdownCancelAllV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsCreateListenKeyV1

> OptionsCreateListenKeyV1Resp OptionsCreateListenKeyV1(ctx).Execute()

Start User Data Stream (USER_STREAM)



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
	resp, r, err := apiClient.V1API.OptionsCreateListenKeyV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsCreateListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsCreateListenKeyV1`: OptionsCreateListenKeyV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsCreateListenKeyV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOptionsCreateListenKeyV1Request struct via the builder pattern


### Return type

[**OptionsCreateListenKeyV1Resp**](OptionsCreateListenKeyV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsCreateMmpResetV1

> OptionsCreateMmpResetV1Resp OptionsCreateMmpResetV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Underlying(underlying).Execute()

Reset Market Maker Protection Config (TRADE)



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
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)
	underlying := "underlying_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.OptionsCreateMmpResetV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Underlying(underlying).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsCreateMmpResetV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsCreateMmpResetV1`: OptionsCreateMmpResetV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsCreateMmpResetV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsCreateMmpResetV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 
 **underlying** | **string** |  | [default to &quot;&quot;]

### Return type

[**OptionsCreateMmpResetV1Resp**](OptionsCreateMmpResetV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsCreateMmpSetV1

> OptionsCreateMmpSetV1Resp OptionsCreateMmpSetV1(ctx).Timestamp(timestamp).DeltaLimit(deltaLimit).FrozenTimeInMilliseconds(frozenTimeInMilliseconds).QtyLimit(qtyLimit).RecvWindow(recvWindow).Underlying(underlying).WindowTimeInMilliseconds(windowTimeInMilliseconds).Execute()

Set Market Maker Protection Config (TRADE)



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
	timestamp := int64(789) // int64 | 
	deltaLimit := "deltaLimit_example" // string |  (optional) (default to "")
	frozenTimeInMilliseconds := int64(789) // int64 |  (optional)
	qtyLimit := "qtyLimit_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	underlying := "underlying_example" // string |  (optional) (default to "")
	windowTimeInMilliseconds := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.OptionsCreateMmpSetV1(context.Background()).Timestamp(timestamp).DeltaLimit(deltaLimit).FrozenTimeInMilliseconds(frozenTimeInMilliseconds).QtyLimit(qtyLimit).RecvWindow(recvWindow).Underlying(underlying).WindowTimeInMilliseconds(windowTimeInMilliseconds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsCreateMmpSetV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsCreateMmpSetV1`: OptionsCreateMmpSetV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsCreateMmpSetV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsCreateMmpSetV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **deltaLimit** | **string** |  | [default to &quot;&quot;]
 **frozenTimeInMilliseconds** | **int64** |  | 
 **qtyLimit** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **underlying** | **string** |  | [default to &quot;&quot;]
 **windowTimeInMilliseconds** | **int64** |  | 

### Return type

[**OptionsCreateMmpSetV1Resp**](OptionsCreateMmpSetV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsCreateOrderV1

> OptionsCreateOrderV1Resp OptionsCreateOrderV1(ctx).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).ClientOrderId(clientOrderId).IsMmp(isMmp).NewOrderRespType(newOrderRespType).PostOnly(postOnly).Price(price).RecvWindow(recvWindow).ReduceOnly(reduceOnly).TimeInForce(timeInForce).Execute()

New Order (TRADE)



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
	quantity := "quantity_example" // string |  (default to "")
	side := "side_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	type_ := "type__example" // string |  (default to "")
	clientOrderId := "clientOrderId_example" // string |  (optional) (default to "")
	isMmp := true // bool |  (optional)
	newOrderRespType := "newOrderRespType_example" // string |  (optional) (default to "")
	postOnly := true // bool |  (optional) (default to false)
	price := "price_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	reduceOnly := true // bool |  (optional) (default to false)
	timeInForce := "timeInForce_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.OptionsCreateOrderV1(context.Background()).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).ClientOrderId(clientOrderId).IsMmp(isMmp).NewOrderRespType(newOrderRespType).PostOnly(postOnly).Price(price).RecvWindow(recvWindow).ReduceOnly(reduceOnly).TimeInForce(timeInForce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsCreateOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsCreateOrderV1`: OptionsCreateOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsCreateOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsCreateOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quantity** | **string** |  | [default to &quot;&quot;]
 **side** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **type_** | **string** |  | [default to &quot;&quot;]
 **clientOrderId** | **string** |  | [default to &quot;&quot;]
 **isMmp** | **bool** |  | 
 **newOrderRespType** | **string** |  | [default to &quot;&quot;]
 **postOnly** | **bool** |  | [default to false]
 **price** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **reduceOnly** | **bool** |  | [default to false]
 **timeInForce** | **string** |  | [default to &quot;&quot;]

### Return type

[**OptionsCreateOrderV1Resp**](OptionsCreateOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsDeleteAllOpenOrdersByUnderlyingV1

> OptionsDeleteAllOpenOrdersByUnderlyingV1Resp OptionsDeleteAllOpenOrdersByUnderlyingV1(ctx).Underlying(underlying).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Cancel All Option Orders By Underlying (TRADE)



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
	underlying := "underlying_example" // string | Option underlying, e.g BTCUSDT (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.OptionsDeleteAllOpenOrdersByUnderlyingV1(context.Background()).Underlying(underlying).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsDeleteAllOpenOrdersByUnderlyingV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsDeleteAllOpenOrdersByUnderlyingV1`: OptionsDeleteAllOpenOrdersByUnderlyingV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsDeleteAllOpenOrdersByUnderlyingV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsDeleteAllOpenOrdersByUnderlyingV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **underlying** | **string** | Option underlying, e.g BTCUSDT | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**OptionsDeleteAllOpenOrdersByUnderlyingV1Resp**](OptionsDeleteAllOpenOrdersByUnderlyingV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsDeleteAllOpenOrdersV1

> OptionsDeleteAllOpenOrdersV1Resp OptionsDeleteAllOpenOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Cancel all Option orders on specific symbol (TRADE)



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
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.OptionsDeleteAllOpenOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsDeleteAllOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsDeleteAllOpenOrdersV1`: OptionsDeleteAllOpenOrdersV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsDeleteAllOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsDeleteAllOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**OptionsDeleteAllOpenOrdersV1Resp**](OptionsDeleteAllOpenOrdersV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsDeleteBatchOrdersV1

> []OptionsDeleteBatchOrdersV1RespInner OptionsDeleteBatchOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderIds(orderIds).ClientOrderIds(clientOrderIds).RecvWindow(recvWindow).Execute()

Cancel Multiple Option Orders (TRADE)



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
	timestamp := int64(789) // int64 | 
	orderIds := []int64{int64(123)} // []int64 | Order ID, e.g [4611875134427365377,4611875134427365378] (optional)
	clientOrderIds := []string{"Inner_example"} // []string | User-defined order ID, e.g [&#34;my_id_1&#34;,&#34;my_id_2&#34;] (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.OptionsDeleteBatchOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderIds(orderIds).ClientOrderIds(clientOrderIds).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsDeleteBatchOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsDeleteBatchOrdersV1`: []OptionsDeleteBatchOrdersV1RespInner
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsDeleteBatchOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsDeleteBatchOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderIds** | **[]int64** | Order ID, e.g [4611875134427365377,4611875134427365378] | 
 **clientOrderIds** | **[]string** | User-defined order ID, e.g [&amp;#34;my_id_1&amp;#34;,&amp;#34;my_id_2&amp;#34;] | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]OptionsDeleteBatchOrdersV1RespInner**](OptionsDeleteBatchOrdersV1RespInner.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsDeleteBlockOrderCreateV1

> map[string]interface{} OptionsDeleteBlockOrderCreateV1(ctx).BlockOrderMatchingKey(blockOrderMatchingKey).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Cancel Block Trade Order (TRADE)



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
	blockOrderMatchingKey := "blockOrderMatchingKey_example" // string |  (default to "")
	timestamp := int32(56) // int32 | 
	recvWindow := int32(56) // int32 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.OptionsDeleteBlockOrderCreateV1(context.Background()).BlockOrderMatchingKey(blockOrderMatchingKey).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsDeleteBlockOrderCreateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsDeleteBlockOrderCreateV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsDeleteBlockOrderCreateV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsDeleteBlockOrderCreateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockOrderMatchingKey** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int32** |  | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 

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


## OptionsDeleteListenKeyV1

> map[string]interface{} OptionsDeleteListenKeyV1(ctx).Execute()

Close User Data Stream (USER_STREAM)



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
	resp, r, err := apiClient.V1API.OptionsDeleteListenKeyV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsDeleteListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsDeleteListenKeyV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsDeleteListenKeyV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOptionsDeleteListenKeyV1Request struct via the builder pattern


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


## OptionsDeleteOrderV1

> OptionsDeleteOrderV1Resp OptionsDeleteOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).ClientOrderId(clientOrderId).RecvWindow(recvWindow).Execute()

Cancel Option Order (TRADE)



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
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 | Order ID, e.g 4611875134427365377 (optional)
	clientOrderId := "clientOrderId_example" // string | User-defined order ID, e.g 10000 (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.OptionsDeleteOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).ClientOrderId(clientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsDeleteOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsDeleteOrderV1`: OptionsDeleteOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsDeleteOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsDeleteOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** | Order ID, e.g 4611875134427365377 | 
 **clientOrderId** | **string** | User-defined order ID, e.g 10000 | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**OptionsDeleteOrderV1Resp**](OptionsDeleteOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetAccountV1

> OptionsGetAccountV1Resp OptionsGetAccountV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Option Account Information(TRADE)



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
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.OptionsGetAccountV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsGetAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetAccountV1`: OptionsGetAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsGetAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**OptionsGetAccountV1Resp**](OptionsGetAccountV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetBillV1

> []OptionsGetBillV1RespItem OptionsGetBillV1(ctx).Currency(currency).Timestamp(timestamp).RecordId(recordId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Account Funding Flow (USER_DATA)



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
	currency := "currency_example" // string | Asset type, only support USDT  as of now (default to "")
	timestamp := int64(789) // int64 | 
	recordId := int64(789) // int64 | Return the recordId and subsequent data, the latest data is returned by default, e.g 100000 (optional)
	startTime := int64(789) // int64 | Start Time, e.g 1593511200000 (optional)
	endTime := int64(789) // int64 | End Time, e.g 1593512200000 (optional)
	limit := int32(56) // int32 | Number of result sets returned Default:100 Max:1000 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.OptionsGetBillV1(context.Background()).Currency(currency).Timestamp(timestamp).RecordId(recordId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsGetBillV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetBillV1`: []OptionsGetBillV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsGetBillV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetBillV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | **string** | Asset type, only support USDT  as of now | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recordId** | **int64** | Return the recordId and subsequent data, the latest data is returned by default, e.g 100000 | 
 **startTime** | **int64** | Start Time, e.g 1593511200000 | 
 **endTime** | **int64** | End Time, e.g 1593512200000 | 
 **limit** | **int32** | Number of result sets returned Default:100 Max:1000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]OptionsGetBillV1RespItem**](OptionsGetBillV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetBlockOrderExecuteV1

> OptionsGetBlockOrderExecuteV1Resp OptionsGetBlockOrderExecuteV1(ctx).BlockOrderMatchingKey(blockOrderMatchingKey).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Block Trade Details (USER_DATA)



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
	blockOrderMatchingKey := "blockOrderMatchingKey_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.OptionsGetBlockOrderExecuteV1(context.Background()).BlockOrderMatchingKey(blockOrderMatchingKey).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsGetBlockOrderExecuteV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetBlockOrderExecuteV1`: OptionsGetBlockOrderExecuteV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsGetBlockOrderExecuteV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetBlockOrderExecuteV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockOrderMatchingKey** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**OptionsGetBlockOrderExecuteV1Resp**](OptionsGetBlockOrderExecuteV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetBlockOrderOrdersV1

> []OptionsGetBlockOrderOrdersV1RespItem OptionsGetBlockOrderOrdersV1(ctx).Timestamp(timestamp).BlockOrderMatchingKey(blockOrderMatchingKey).EndTime(endTime).StartTime(startTime).Underlying(underlying).RecvWindow(recvWindow).Execute()

Query Block Trade Order (TRADE)



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
	timestamp := int64(789) // int64 | 
	blockOrderMatchingKey := "blockOrderMatchingKey_example" // string | If specified, returns the specific block trade associated with the blockOrderMatchingKey (optional) (default to "")
	endTime := int64(789) // int64 |  (optional)
	startTime := int64(789) // int64 |  (optional)
	underlying := "underlying_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.OptionsGetBlockOrderOrdersV1(context.Background()).Timestamp(timestamp).BlockOrderMatchingKey(blockOrderMatchingKey).EndTime(endTime).StartTime(startTime).Underlying(underlying).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsGetBlockOrderOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetBlockOrderOrdersV1`: []OptionsGetBlockOrderOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsGetBlockOrderOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetBlockOrderOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **blockOrderMatchingKey** | **string** | If specified, returns the specific block trade associated with the blockOrderMatchingKey | [default to &quot;&quot;]
 **endTime** | **int64** |  | 
 **startTime** | **int64** |  | 
 **underlying** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]OptionsGetBlockOrderOrdersV1RespItem**](OptionsGetBlockOrderOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	resp, r, err := apiClient.V1API.OptionsGetBlockTradesV1(context.Background()).Symbol(symbol).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsGetBlockTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetBlockTradesV1`: []OptionsGetBlockTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsGetBlockTradesV1`: %v\n", resp)
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


## OptionsGetBlockUserTradesV1

> []OptionsGetBlockUserTradesV1RespItem OptionsGetBlockUserTradesV1(ctx).Timestamp(timestamp).EndTime(endTime).StartTime(startTime).Underlying(underlying).RecvWindow(recvWindow).Execute()

Account Block Trade List (USER_DATA)



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
	timestamp := int64(789) // int64 | 
	endTime := int64(789) // int64 |  (optional)
	startTime := int64(789) // int64 |  (optional)
	underlying := "underlying_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.OptionsGetBlockUserTradesV1(context.Background()).Timestamp(timestamp).EndTime(endTime).StartTime(startTime).Underlying(underlying).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsGetBlockUserTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetBlockUserTradesV1`: []OptionsGetBlockUserTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsGetBlockUserTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetBlockUserTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **endTime** | **int64** |  | 
 **startTime** | **int64** |  | 
 **underlying** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]OptionsGetBlockUserTradesV1RespItem**](OptionsGetBlockUserTradesV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetCountdownCancelAllV1

> OptionsGetCountdownCancelAllV1Resp OptionsGetCountdownCancelAllV1(ctx).Timestamp(timestamp).Underlying(underlying).RecvWindow(recvWindow).Execute()

Get Auto-Cancel All Open Orders (Kill-Switch) Config (TRADE)



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
	timestamp := int64(789) // int64 | 
	underlying := "underlying_example" // string | Option underlying, e.g BTCUSDT (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.OptionsGetCountdownCancelAllV1(context.Background()).Timestamp(timestamp).Underlying(underlying).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsGetCountdownCancelAllV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetCountdownCancelAllV1`: OptionsGetCountdownCancelAllV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsGetCountdownCancelAllV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetCountdownCancelAllV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **underlying** | **string** | Option underlying, e.g BTCUSDT | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**OptionsGetCountdownCancelAllV1Resp**](OptionsGetCountdownCancelAllV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

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
	resp, r, err := apiClient.V1API.OptionsGetDepthV1(context.Background()).Symbol(symbol).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsGetDepthV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetDepthV1`: OptionsGetDepthV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsGetDepthV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.OptionsGetExchangeInfoV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsGetExchangeInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetExchangeInfoV1`: OptionsGetExchangeInfoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsGetExchangeInfoV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.OptionsGetExerciseHistoryV1(context.Background()).Underlying(underlying).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsGetExerciseHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetExerciseHistoryV1`: []OptionsGetExerciseHistoryV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsGetExerciseHistoryV1`: %v\n", resp)
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


## OptionsGetExerciseRecordV1

> []OptionsGetExerciseRecordV1RespItem OptionsGetExerciseRecordV1(ctx).Timestamp(timestamp).Symbol(symbol).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

User Exercise Record (USER_DATA)



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
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string | Option trading pair, e.g BTC-200730-9000-C (optional) (default to "")
	startTime := int64(789) // int64 | startTime (optional)
	endTime := int64(789) // int64 | endTime (optional)
	limit := int32(56) // int32 | default 1000, max 1000 (optional) (default to 1000)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.OptionsGetExerciseRecordV1(context.Background()).Timestamp(timestamp).Symbol(symbol).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsGetExerciseRecordV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetExerciseRecordV1`: []OptionsGetExerciseRecordV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsGetExerciseRecordV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetExerciseRecordV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | [default to &quot;&quot;]
 **startTime** | **int64** | startTime | 
 **endTime** | **int64** | endTime | 
 **limit** | **int32** | default 1000, max 1000 | [default to 1000]
 **recvWindow** | **int64** |  | 

### Return type

[**[]OptionsGetExerciseRecordV1RespItem**](OptionsGetExerciseRecordV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

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
	resp, r, err := apiClient.V1API.OptionsGetHistoricalTradesV1(context.Background()).Symbol(symbol).FromId(fromId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsGetHistoricalTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetHistoricalTradesV1`: []OptionsGetHistoricalTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsGetHistoricalTradesV1`: %v\n", resp)
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


## OptionsGetHistoryOrdersV1

> []OptionsGetHistoryOrdersV1RespItem OptionsGetHistoryOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query Option Order History (TRADE)



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
	symbol := "symbol_example" // string | Option trading pair (default to "")
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 | Returns the orderId and subsequent orders, the most recent order is returned by default (optional)
	startTime := int64(789) // int64 | Start Time, e.g 1593511200000 (optional)
	endTime := int64(789) // int64 | End Time, e.g 1593512200000 (optional)
	limit := int32(56) // int32 | Number of result sets returned Default:100 Max:1000 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.OptionsGetHistoryOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsGetHistoryOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetHistoryOrdersV1`: []OptionsGetHistoryOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsGetHistoryOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetHistoryOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** | Returns the orderId and subsequent orders, the most recent order is returned by default | 
 **startTime** | **int64** | Start Time, e.g 1593511200000 | 
 **endTime** | **int64** | End Time, e.g 1593512200000 | 
 **limit** | **int32** | Number of result sets returned Default:100 Max:1000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]OptionsGetHistoryOrdersV1RespItem**](OptionsGetHistoryOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetIncomeAsynIdV1

> OptionsGetIncomeAsynIdV1Resp OptionsGetIncomeAsynIdV1(ctx).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Option Transaction History Download Link by Id (USER_DATA)



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
	downloadId := "downloadId_example" // string | get by download id api (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.OptionsGetIncomeAsynIdV1(context.Background()).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsGetIncomeAsynIdV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetIncomeAsynIdV1`: OptionsGetIncomeAsynIdV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsGetIncomeAsynIdV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetIncomeAsynIdV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadId** | **string** | get by download id api | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**OptionsGetIncomeAsynIdV1Resp**](OptionsGetIncomeAsynIdV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetIncomeAsynV1

> OptionsGetIncomeAsynV1Resp OptionsGetIncomeAsynV1(ctx).Execute()

Get Download Id For Option Transaction History (USER_DATA)



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
	resp, r, err := apiClient.V1API.OptionsGetIncomeAsynV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsGetIncomeAsynV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetIncomeAsynV1`: OptionsGetIncomeAsynV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsGetIncomeAsynV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetIncomeAsynV1Request struct via the builder pattern


### Return type

[**OptionsGetIncomeAsynV1Resp**](OptionsGetIncomeAsynV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

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
	resp, r, err := apiClient.V1API.OptionsGetIndexV1(context.Background()).Underlying(underlying).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsGetIndexV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetIndexV1`: OptionsGetIndexV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsGetIndexV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.OptionsGetKlinesV1(context.Background()).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsGetKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetKlinesV1`: []OptionsGetKlinesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsGetKlinesV1`: %v\n", resp)
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


## OptionsGetMarginAccountV1

> OptionsGetMarginAccountV1Resp OptionsGetMarginAccountV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Option Margin Account Information (USER_DATA)



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
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.OptionsGetMarginAccountV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsGetMarginAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetMarginAccountV1`: OptionsGetMarginAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsGetMarginAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetMarginAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**OptionsGetMarginAccountV1Resp**](OptionsGetMarginAccountV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

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
	resp, r, err := apiClient.V1API.OptionsGetMarkV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsGetMarkV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetMarkV1`: []OptionsGetMarkV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsGetMarkV1`: %v\n", resp)
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


## OptionsGetMmpV1

> OptionsGetMmpV1Resp OptionsGetMmpV1(ctx).Timestamp(timestamp).Underlying(underlying).RecvWindow(recvWindow).Execute()

Get Market Maker Protection Config (TRADE)



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
	timestamp := int64(789) // int64 | 
	underlying := "underlying_example" // string | underlying, e.g BTCUSDT (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.OptionsGetMmpV1(context.Background()).Timestamp(timestamp).Underlying(underlying).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsGetMmpV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetMmpV1`: OptionsGetMmpV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsGetMmpV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetMmpV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **underlying** | **string** | underlying, e.g BTCUSDT | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**OptionsGetMmpV1Resp**](OptionsGetMmpV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

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
	resp, r, err := apiClient.V1API.OptionsGetOpenInterestV1(context.Background()).UnderlyingAsset(underlyingAsset).Expiration(expiration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsGetOpenInterestV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetOpenInterestV1`: []OptionsGetOpenInterestV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsGetOpenInterestV1`: %v\n", resp)
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


## OptionsGetOpenOrdersV1

> []OptionsGetOpenOrdersV1RespItem OptionsGetOpenOrdersV1(ctx).Timestamp(timestamp).Symbol(symbol).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query Current Open Option Orders (USER_DATA)



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
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string | return all orders if don&#39;t pass, Option trading pair, e.g BTC-200730-9000-C, (optional) (default to "")
	orderId := int64(789) // int64 | Returns the orderId and subsequent orders, the most recent order is returned by default (optional)
	startTime := int64(789) // int64 | Start Time (optional)
	endTime := int64(789) // int64 | End Time (optional)
	limit := int32(56) // int32 | Number of result sets returned Default:100 Max:1000 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.OptionsGetOpenOrdersV1(context.Background()).Timestamp(timestamp).Symbol(symbol).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsGetOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetOpenOrdersV1`: []OptionsGetOpenOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsGetOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** | return all orders if don&amp;#39;t pass, Option trading pair, e.g BTC-200730-9000-C, | [default to &quot;&quot;]
 **orderId** | **int64** | Returns the orderId and subsequent orders, the most recent order is returned by default | 
 **startTime** | **int64** | Start Time | 
 **endTime** | **int64** | End Time | 
 **limit** | **int32** | Number of result sets returned Default:100 Max:1000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]OptionsGetOpenOrdersV1RespItem**](OptionsGetOpenOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetOrderV1

> OptionsGetOrderV1Resp OptionsGetOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).ClientOrderId(clientOrderId).RecvWindow(recvWindow).Execute()

Query Single Order (TRADE)



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
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 | Order id (optional)
	clientOrderId := "clientOrderId_example" // string | User-defined order ID cannot be repeated in pending orders (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.OptionsGetOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).ClientOrderId(clientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsGetOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetOrderV1`: OptionsGetOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsGetOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** | Order id | 
 **clientOrderId** | **string** | User-defined order ID cannot be repeated in pending orders | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**OptionsGetOrderV1Resp**](OptionsGetOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

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
	resp, r, err := apiClient.V1API.OptionsGetPingV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsGetPingV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetPingV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsGetPingV1`: %v\n", resp)
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


## OptionsGetPositionV1

> []OptionsGetPositionV1RespItem OptionsGetPositionV1(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Option Position Information (USER_DATA)



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
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string | Option trading pair, e.g BTC-200730-9000-C (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.OptionsGetPositionV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsGetPositionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetPositionV1`: []OptionsGetPositionV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsGetPositionV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetPositionV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]OptionsGetPositionV1RespItem**](OptionsGetPositionV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

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
	resp, r, err := apiClient.V1API.OptionsGetTickerV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsGetTickerV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetTickerV1`: []OptionsGetTickerV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsGetTickerV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.OptionsGetTimeV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsGetTimeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetTimeV1`: OptionsGetTimeV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsGetTimeV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.OptionsGetTradesV1(context.Background()).Symbol(symbol).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsGetTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetTradesV1`: []OptionsGetTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsGetTradesV1`: %v\n", resp)
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


## OptionsGetUserTradesV1

> []OptionsGetUserTradesV1RespItem OptionsGetUserTradesV1(ctx).Timestamp(timestamp).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Account Trade List (USER_DATA)



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
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string | Option symbol, e.g BTC-200730-9000-C (optional) (default to "")
	fromId := int64(789) // int64 | Trade id to fetch from. Default gets most recent trades, e.g 4611875134427365376 (optional)
	startTime := int64(789) // int64 | Start time, e.g 1593511200000 (optional)
	endTime := int64(789) // int64 | End time, e.g 1593512200000 (optional)
	limit := int32(56) // int32 | Default 100; max 1000 (optional) (default to 100)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.OptionsGetUserTradesV1(context.Background()).Timestamp(timestamp).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsGetUserTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetUserTradesV1`: []OptionsGetUserTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsGetUserTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetUserTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** | Option symbol, e.g BTC-200730-9000-C | [default to &quot;&quot;]
 **fromId** | **int64** | Trade id to fetch from. Default gets most recent trades, e.g 4611875134427365376 | 
 **startTime** | **int64** | Start time, e.g 1593511200000 | 
 **endTime** | **int64** | End time, e.g 1593512200000 | 
 **limit** | **int32** | Default 100; max 1000 | [default to 100]
 **recvWindow** | **int64** |  | 

### Return type

[**[]OptionsGetUserTradesV1RespItem**](OptionsGetUserTradesV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsUpdateBlockOrderCreateV1

> OptionsUpdateBlockOrderCreateV1Resp OptionsUpdateBlockOrderCreateV1(ctx).BlockOrderMatchingKey(blockOrderMatchingKey).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Extend Block Trade Order (TRADE)



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
	blockOrderMatchingKey := "blockOrderMatchingKey_example" // string |  (default to "")
	timestamp := int32(56) // int32 | 
	recvWindow := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.OptionsUpdateBlockOrderCreateV1(context.Background()).BlockOrderMatchingKey(blockOrderMatchingKey).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsUpdateBlockOrderCreateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsUpdateBlockOrderCreateV1`: OptionsUpdateBlockOrderCreateV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsUpdateBlockOrderCreateV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsUpdateBlockOrderCreateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockOrderMatchingKey** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int32** |  | 
 **recvWindow** | **int32** |  | 

### Return type

[**OptionsUpdateBlockOrderCreateV1Resp**](OptionsUpdateBlockOrderCreateV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsUpdateListenKeyV1

> map[string]interface{} OptionsUpdateListenKeyV1(ctx).Execute()

Keepalive User Data Stream (USER_STREAM)



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
	resp, r, err := apiClient.V1API.OptionsUpdateListenKeyV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OptionsUpdateListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsUpdateListenKeyV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.OptionsUpdateListenKeyV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOptionsUpdateListenKeyV1Request struct via the builder pattern


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

